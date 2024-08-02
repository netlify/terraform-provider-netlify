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
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
	"github.com/netlify/terraform-provider-netlify/internal/provider/netlify_validators"
)

var (
	_ resource.Resource                = &environmentVariableResource{}
	_ resource.ResourceWithConfigure   = &environmentVariableResource{}
	_ resource.ResourceWithImportState = &environmentVariableResource{}
)

func NewEnvironmentVariableResource() resource.Resource {
	return &environmentVariableResource{}
}

type environmentVariableResource struct {
	data NetlifyProviderData
}

type environmentVariableResourceModel struct {
	TeamID       types.String                    `tfsdk:"team_id"`
	SiteID       types.String                    `tfsdk:"site_id"`
	LastUpdated  types.String                    `tfsdk:"last_updated"`
	Key          types.String                    `tfsdk:"key"`
	Scopes       []types.String                  `tfsdk:"scopes"`
	Values       []environmentVariableValueModel `tfsdk:"values"`
	SecretValues []environmentVariableValueModel `tfsdk:"secret_values"`
}

type environmentVariableValueModel struct {
	Value            types.String `tfsdk:"value"`
	Context          types.String `tfsdk:"context"`
	ContextParameter types.String `tfsdk:"context_parameter"`
}

func (r *environmentVariableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment_variable"
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
		Description:         "Environment variables for Netlify sites",
		MarkdownDescription: "Environment variables for Netlify sites. [Read more](https://docs.netlify.com/environment-variables/overview/)",
		Attributes: map[string]schema.Attribute{
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"team_id": schema.StringAttribute{
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
				Description: "One or more of builds, functions, runtime, and post-processing",
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
			"values": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO: confirm it's OK that we aren't tracking the ID of value items
						"value": schema.StringAttribute{
							Required: true,
						},
						"context": schema.StringAttribute{
							Required:    true,
							Description: "One of all, dev, branch-deploy, deploy-preview, production, or branch",
							Validators: []validator.String{
								stringvalidator.OneOf("all", "dev", "branch-deploy", "deploy-preview", "production", "branch"),
							},
						},
						"context_parameter": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								netlify_validators.EnvironmentVariableContextParameterValidator{
									ContextPathExpression: path.MatchRelative().AtParent().AtName("context"),
								},
							},
						},
					},
				},
				// TODO: validate that values don't overlap
			},
			"secret_values": schema.SetNestedAttribute{
				Optional: true,
				Validators: []validator.Set{
					setvalidator.ExactlyOneOf(path.MatchRoot("value")),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						// TODO: confirm it's OK that we aren't tracking the ID of value items
						"value": schema.StringAttribute{
							Required:  true,
							Sensitive: true,
						},
						"context": schema.StringAttribute{
							Required:    true,
							Description: "One of all, dev, branch-deploy, deploy-preview, production, or branch",
							Validators: []validator.String{
								stringvalidator.OneOf("all", "dev", "branch-deploy", "deploy-preview", "production", "branch"),
							},
						},
						"context_parameter": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(""),
							Validators: []validator.String{
								netlify_validators.EnvironmentVariableContextParameterValidator{
									ContextPathExpression: path.MatchRelative().AtParent().AtName("context"),
								},
							},
						},
					},
				},
				// TODO: validate that values don't overlap
			},
		},
	}
}

func (r *environmentVariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan environmentVariableResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	scopes := make([]string, len(plan.Scopes))
	for i, scope := range plan.Scopes {
		scopes[i] = scope.ValueString()
	}
	var values []netlifyapi.EnvVarValue
	var isSecret bool
	if plan.SecretValues != nil && len(plan.SecretValues) > 0 {
		values = serializeValues(plan.SecretValues)
		isSecret = true
	} else {
		values = serializeValues(plan.Values)
		isSecret = false
	}
	createEnvVars := r.data.client.EnvironmentVariablesAPI.
		CreateEnvVars(ctx, plan.TeamID.ValueString()).
		EnvVar([]netlifyapi.EnvVar{
			{
				Key:      plan.Key.ValueString(),
				Scopes:   scopes,
				Values:   values,
				IsSecret: &isSecret,
			},
		})
	if plan.SiteID.ValueString() != "" {
		createEnvVars = createEnvVars.SiteId(plan.SiteID.ValueString())
	}
	envVars, _, err := createEnvVars.Execute()
	if err != nil || len(envVars) == 0 {
		resp.Diagnostics.AddError(
			"Error creating Netlify environment variable",
			fmt.Sprintf(
				"Could not create Netlify environment variable order ID %q (team ID: %q, site ID: %q, secret: %v): %q",
				plan.Key.ValueString(),
				plan.TeamID.ValueString(),
				plan.SiteID.ValueString(),
				isSecret,
				err.Error(),
			),
		)
		return
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *environmentVariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state environmentVariableResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	getEnvVar := r.data.client.EnvironmentVariablesAPI.GetEnvVar(ctx, state.TeamID.ValueString(), state.Key.ValueString())
	if state.SiteID.ValueString() != "" {
		getEnvVar = getEnvVar.SiteId(state.SiteID.ValueString())
	}
	envVar, _, err := getEnvVar.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Netlify environment variable",
			fmt.Sprintf(
				"Could not read Netlify environment variable order ID %q (team ID: %q, site ID: %q): %q",
				state.Key.ValueString(),
				state.TeamID.ValueString(),
				state.SiteID.ValueString(),
				err.Error(),
			),
		)
		return
	}

	state.Scopes = make([]types.String, len(envVar.Scopes))
	for i, scope := range envVar.Scopes {
		state.Scopes[i] = types.StringValue(strings.ReplaceAll(strings.ReplaceAll(scope, " ", "-"), "_", "-"))
	}
	if !*envVar.IsSecret {
		state.Values = parseValues(envVar.Values)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *environmentVariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan environmentVariableResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	scopes := make([]string, len(plan.Scopes))
	for i, scope := range plan.Scopes {
		scopes[i] = scope.ValueString()
	}
	var values []netlifyapi.EnvVarValue
	var isSecret bool
	if plan.SecretValues != nil && len(plan.SecretValues) > 0 {
		values = serializeValues(plan.SecretValues)
		isSecret = true
	} else {
		values = serializeValues(plan.Values)
		isSecret = false
	}
	updateEnvVar := r.data.client.EnvironmentVariablesAPI.
		UpdateEnvVar(ctx, plan.TeamID.ValueString(), plan.Key.ValueString()).
		Key(plan.Key.ValueString()).UpdateEnvVarRequest(netlifyapi.UpdateEnvVarRequest{
		Scopes:   scopes,
		Values:   values,
		IsSecret: &isSecret,
	})
	if plan.SiteID.ValueString() != "" {
		updateEnvVar = updateEnvVar.SiteId(plan.SiteID.ValueString())
	}
	_, _, err := updateEnvVar.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Netlify environment variable",
			fmt.Sprintf(
				"Could not update Netlify environment variable order ID %q (team ID: %q, site ID: %q, secret: %v): %q",
				plan.Key.ValueString(),
				plan.TeamID.ValueString(),
				plan.SiteID.ValueString(),
				isSecret,
				err.Error(),
			),
		)
		return
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *environmentVariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state environmentVariableResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	deleteEnvVar := r.data.client.EnvironmentVariablesAPI.
		DeleteEnvVar(ctx, state.TeamID.ValueString(), state.Key.ValueString())
	if state.SiteID.ValueString() != "" {
		deleteEnvVar = deleteEnvVar.SiteId(state.SiteID.ValueString())
	}
	_, err := deleteEnvVar.Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Netlify environment variable",
			fmt.Sprintf(
				"Could not delete Netlify environment variable order ID %q (team ID: %q, site ID: %q): %q",
				state.Key.ValueString(),
				state.TeamID.ValueString(),
				state.SiteID.ValueString(),
				err.Error(),
			),
		)
		return
	}
}

func (r *environmentVariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ":")

	errorMessage := fmt.Sprintf("Expected import identifier with one of these formats: team_id,key or team_id,site_id,key. Got: %q", req.ID)

	if len(idParts) == 2 {
		if idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("team_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), idParts[1])...)
	} else if len(idParts) == 3 {
		if idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("team_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[1])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key"), idParts[2])...)
	} else {
		resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
		return
	}
}

func serializeValues(values []environmentVariableValueModel) []netlifyapi.EnvVarValue {
	envVarValues := make([]netlifyapi.EnvVarValue, len(values))
	for i, value := range values {
		envVarValues[i] = netlifyapi.EnvVarValue{
			Value:            value.Value.ValueStringPointer(),
			Context:          value.Context.ValueStringPointer(),
			ContextParameter: value.ContextParameter.ValueStringPointer(),
		}
		if envVarValues[i].ContextParameter != nil && *envVarValues[i].ContextParameter == "" {
			envVarValues[i].ContextParameter = nil
		}
	}
	return envVarValues
}

func parseValues(values []netlifyapi.EnvVarValue) []environmentVariableValueModel {
	envVarValues := make([]environmentVariableValueModel, len(values))
	for i, value := range values {
		envVarValues[i] = environmentVariableValueModel{
			Value:            types.StringPointerValue(value.Value),
			Context:          types.StringPointerValue(value.Context),
			ContextParameter: types.StringPointerValue(value.ContextParameter),
		}
		if envVarValues[i].ContextParameter.ValueString() == "" {
			envVarValues[i].ContextParameter = types.StringPointerValue(nil)
		}
	}
	return envVarValues
}
