package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/models"
	"github.com/netlify/terraform-provider-netlify/internal/plumbing/operations"
)

var (
	_ resource.Resource                = &environmentVariableResource{}
	_ resource.ResourceWithConfigure   = &environmentVariableResource{}
	_ resource.ResourceWithImportState = &environmentVariableResource{}
)

func NewEnvironmentVariableResource(is_secret bool) resource.Resource {
	return &environmentVariableResource{
		is_secret: is_secret,
	}
}

type environmentVariableResource struct {
	data      NetlifyProviderData
	is_secret bool
}

type environmentVariableResourceModel struct {
	AccountID   types.String                    `tfsdk:"account_id"`
	SiteID      types.String                    `tfsdk:"site_id"`
	LastUpdated types.String                    `tfsdk:"last_updated"`
	Key         types.String                    `tfsdk:"key"`
	Scopes      []types.String                  `tfsdk:"scopes"`
	Value       []environmentVariableValueModel `tfsdk:"value"`
}

type environmentVariableValueModel struct {
	Value            types.String `tfsdk:"value"`
	Context          types.String `tfsdk:"context"`
	ContextParameter types.String `tfsdk:"context_parameter"`
}

func (r *environmentVariableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	if r.is_secret {
		resp.TypeName = req.ProviderTypeName + "_secret_environment_variable"
	} else {
		resp.TypeName = req.ProviderTypeName + "_environment_variable"
	}
}

func (r *environmentVariableResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(NetlifyProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected NetlifyProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.data = data
}

func (r *environmentVariableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"key": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"scopes": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("builds", "functions", "runtime", "post-processing"),
					),
				},
				Default: setdefault.StaticValue(types.SetValueMust(types.StringType, []attr.Value{
					types.StringValue("builds"),
					types.StringValue("functions"),
					types.StringValue("runtime"),
					types.StringValue("post-processing"),
				})),
			},
			"value": schema.SetNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO: confirm it's OK that we aren't tracking the ID of value items
						"value": schema.StringAttribute{
							Required:  true,
							Sensitive: r.is_secret,
						},
						"context": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("all", "dev", "branch-deploy", "deploy-preview", "production", "branch"),
							},
						},
						"context_parameter": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(""),
						},
					},
					// TODO: validate that context_parameter is not empty iff context is "branch"
				},
				// TODO: validate that values don't overlap
			},
		},
	}
}

func (r *environmentVariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan environmentVariableResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	scopes := []string{}
	for _, scope := range plan.Scopes {
		scopes = append(scopes, scope.ValueString())
	}
	createEnvVarsParams := operations.
		NewCreateEnvVarsParams().
		WithAccountID(plan.AccountID.ValueString()).
		WithEnvVars([]*models.CreateEnvVarsParamsBodyItems{
			{
				Key:      plan.Key.ValueString(),
				Scopes:   scopes,
				Values:   serializeValues(plan.Value),
				IsSecret: r.is_secret,
			},
		})
	if plan.SiteID.ValueString() != "" {
		createEnvVarsParams.SetSiteID(plan.SiteID.ValueStringPointer())
	}
	envVar, err := r.data.client.Operations.CreateEnvVars(createEnvVarsParams, r.data.authInfo)
	if err != nil || len(envVar.Payload) == 0 {
		resp.Diagnostics.AddError(
			"Error creating Netlify environment variable",
			fmt.Sprintf(
				"Could not create Netlify environment variable order ID %q (account ID: %q, site ID: %q, secret: %v): %q",
				plan.Key.ValueString(),
				plan.AccountID.ValueString(),
				plan.SiteID.ValueString(),
				r.is_secret,
				err.Error(),
			),
		)
		return
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *environmentVariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state environmentVariableResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	getEnvVarParams := operations.
		NewGetEnvVarParams().
		WithAccountID(state.AccountID.ValueString()).
		WithKey(state.Key.ValueString())
	if state.SiteID.ValueString() != "" {
		getEnvVarParams.SetSiteID(state.SiteID.ValueStringPointer())
	}
	envVar, err := r.data.client.Operations.GetEnvVar(getEnvVarParams, r.data.authInfo)
	if err != nil || envVar.Payload.IsSecret != r.is_secret {
		resp.Diagnostics.AddError(
			"Error reading Netlify environment variable",
			fmt.Sprintf(
				"Could not read Netlify environment variable order ID %q (account ID: %q, site ID: %q, secret: %v): %q",
				state.Key.ValueString(),
				state.AccountID.ValueString(),
				state.SiteID.ValueString(),
				r.is_secret,
				err.Error(),
			),
		)
		return
	}

	state.Scopes = []types.String{}
	for _, scope := range envVar.Payload.Scopes {
		state.Scopes = append(state.Scopes, types.StringValue(strings.ReplaceAll(strings.ReplaceAll(scope, " ", "-"), "_", "-")))
	}
	state.Value = parseValues(envVar.Payload.Values)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *environmentVariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan environmentVariableResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	scopes := []string{}
	for _, scope := range plan.Scopes {
		scopes = append(scopes, scope.ValueString())
	}
	updateEnvVarParams := operations.
		NewUpdateEnvVarParams().
		WithAccountID(plan.AccountID.ValueString()).
		WithKey(plan.Key.ValueString()).
		WithEnvVar(&models.UpdateEnvVarParamsBody{
			Key:      plan.Key.ValueString(),
			Scopes:   scopes,
			Values:   serializeValues(plan.Value),
			IsSecret: r.is_secret,
		})
	if plan.SiteID.ValueString() != "" {
		updateEnvVarParams.SetSiteID(plan.SiteID.ValueStringPointer())
	}
	_, err := r.data.client.Operations.UpdateEnvVar(updateEnvVarParams, r.data.authInfo)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Netlify environment variable",
			fmt.Sprintf(
				"Could not update Netlify environment variable order ID %q (account ID: %q, site ID: %q, secret: %v): %q",
				plan.Key.ValueString(),
				plan.AccountID.ValueString(),
				plan.SiteID.ValueString(),
				r.is_secret,
				err.Error(),
			),
		)
		return
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *environmentVariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state environmentVariableResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteEnvVarParams := operations.
		NewDeleteEnvVarParams().
		WithAccountID(state.AccountID.ValueString()).
		WithKey(state.Key.ValueString())
	if state.SiteID.ValueString() != "" {
		deleteEnvVarParams.SetSiteID(state.SiteID.ValueStringPointer())
	}
	_, err := r.data.client.Operations.DeleteEnvVar(deleteEnvVarParams, r.data.authInfo)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Netlify environment variable",
			fmt.Sprintf(
				"Could not delete Netlify environment variable order ID %q (account ID: %q, site ID: %q, secret: %v): %q",
				state.Key.ValueString(),
				state.AccountID.ValueString(),
				state.SiteID.ValueString(),
				r.is_secret,
				err.Error(),
			),
		)
		return
	}
}

func (r *environmentVariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	errorMessage := fmt.Sprintf("Expected import identifier with one of these formats: account_id,key or account_id,site_id,key. Got: %q", req.ID)

	if len(idParts) == 2 {
		if idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), idParts[1])...)
	} else if len(idParts) == 3 {
		if idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("account_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), idParts[2])...)
	} else {
		resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
		return
	}
}

func serializeValues(values []environmentVariableValueModel) []*models.EnvVarValue {
	envVarValues := []*models.EnvVarValue{}
	for _, value := range values {
		envVarValue := &models.EnvVarValue{
			Value:            value.Value.ValueString(),
			Context:          value.Context.ValueString(),
			ContextParameter: value.ContextParameter.ValueString(),
		}
		envVarValues = append(envVarValues, envVarValue)
	}
	return envVarValues
}

func parseValues(values []*models.EnvVarValue) []environmentVariableValueModel {
	envVarValues := []environmentVariableValueModel{}
	for _, value := range values {
		envVarValues = append(envVarValues, environmentVariableValueModel{
			Value:            types.StringValue(value.Value),
			Context:          types.StringValue(value.Context),
			ContextParameter: types.StringValue(value.ContextParameter),
		})
	}
	return envVarValues
}
