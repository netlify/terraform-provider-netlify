package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
	"github.com/netlify/terraform-provider-netlify/internal/provider/netlify_validators"
)

var (
	_ resource.Resource                = &logDrainResource{}
	_ resource.ResourceWithConfigure   = &logDrainResource{}
	_ resource.ResourceWithImportState = &logDrainResource{}
)

func NewLogDrainResource() resource.Resource {
	return &logDrainResource{}
}

type logDrainResource struct {
	data NetlifyProviderData
}

type logDrainResourceModel struct {
	ID            types.String                `tfsdk:"id"`
	SiteID        types.String                `tfsdk:"site_id"`
	LastUpdated   types.String                `tfsdk:"last_updated"`
	Destination   types.String                `tfsdk:"destination"`
	Format        types.String                `tfsdk:"format"`
	LogTypes      []types.String              `tfsdk:"log_types"`
	ExcludePII    types.Bool                  `tfsdk:"exclude_pii"`
	ServiceConfig *logDrainServiceConfigModel `tfsdk:"service_config"`
}

type logDrainServiceConfigModel struct {
	URL                  types.String            `tfsdk:"url"`
	Tags                 map[string]types.String `tfsdk:"tags"`
	IntegrationName      types.String            `tfsdk:"integration_name"`
	AuthHeader           types.String            `tfsdk:"authorization_header"`
	BucketName           types.String            `tfsdk:"bucket_name"`
	BucketRegion         types.String            `tfsdk:"bucket_region"`
	Path                 types.String            `tfsdk:"path"`
	VerificationFilename types.String            `tfsdk:"verification_filename"`
}

func (r *logDrainResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_log_drain"
}

func (r *logDrainResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *logDrainResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Netlify log drain",
		MarkdownDescription: "Netlify log drain. [Read more](https://docs.netlify.com/monitor-sites/log-drains/)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"destination": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "One of datadog, newrelic, logflare, s3, splunkcloud, http, axiom, or azure",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"datadog",
						"newrelic",
						"logflare",
						"s3",
						"splunkcloud",
						"http",
						"axiom",
						"azure",
					),
					// TODO: These are partial validations, need to be completed
					netlify_validators.RequiredIfEquals("datadog", path.MatchRoot("service_config").AtName("url")),
					netlify_validators.ForbiddenIfEquals( // Must be part of the URL
						"datadog",
						path.MatchRoot("service_config").AtName("tags"),
						path.MatchRoot("service_config").AtName("authorization_header"),
					),
					netlify_validators.RequiredIfEquals(
						"newrelic",
						path.MatchRoot("service_config").AtName("url"),
						path.MatchRoot("service_config").AtName("tags"),
					),
					netlify_validators.RequiredIfEquals("http", path.MatchRoot("service_config").AtName("url")),
					netlify_validators.RequiredIfEquals(
						"s3",
						path.MatchRoot("service_config").AtName("bucket_name"),
						path.MatchRoot("service_config").AtName("bucket_region"),
						path.MatchRoot("service_config").AtName("path"),
						path.MatchRoot("service_config").AtName("verification_filename"),
					),
				},
			},
			"format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "json or ndjson",
				Default:     stringdefault.StaticString("json"),
			},
			"log_types": schema.SetAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "One or more of user_traffic, functions, edge_functions, waf, and deploys",
				Validators: []validator.Set{
					setvalidator.ValueStringsAre(
						stringvalidator.OneOf("user_traffic", "functions", "edge_functions", "waf", "deploys"),
					),
				},
			},
			"exclude_pii": schema.BoolAttribute{
				Required: true,
			},
			"service_config": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
					},
					"tags": schema.MapAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"integration_name": schema.StringAttribute{
						Optional: true,
					},
					"authorization_header": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
					},
					"bucket_name": schema.StringAttribute{
						Optional: true,
					},
					"bucket_region": schema.StringAttribute{
						Optional: true,
					},
					"path": schema.StringAttribute{
						Optional: true,
					},
					"verification_filename": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r *logDrainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan logDrainResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var logTypes string
	logTypesArr := make([]string, len(plan.LogTypes))
	for i, logType := range plan.LogTypes {
		logTypesArr[i] = logType.ValueString()
	}
	if logTypesBytes, err := json.Marshal(logTypesArr); err != nil {
		resp.Diagnostics.AddError(
			"Error creating log drain",
			fmt.Sprintf("Could not marshal log types: %q", err.Error()),
		)
		return
	} else {
		logTypes = string(logTypesBytes)
	}

	var serviceConfig netlifyapi.LogDrainServiceConfig
	if plan.ServiceConfig == nil {
		plan.ServiceConfig = &logDrainServiceConfigModel{}
	}
	serviceConfig.Url = plan.ServiceConfig.URL.ValueStringPointer()
	if len(plan.ServiceConfig.Tags) > 0 {
		serviceConfig.Tags = make(map[string]string)
		for k, v := range plan.ServiceConfig.Tags {
			serviceConfig.Tags[k] = v.ValueString()
		}
	}
	serviceConfig.IntegrationName = plan.ServiceConfig.IntegrationName.ValueStringPointer()
	serviceConfig.AuthorizationHeader = plan.ServiceConfig.AuthHeader.ValueStringPointer()
	serviceConfig.BucketName = plan.ServiceConfig.BucketName.ValueStringPointer()
	serviceConfig.BucketRegion = plan.ServiceConfig.BucketRegion.ValueStringPointer()
	serviceConfig.Path = plan.ServiceConfig.Path.ValueStringPointer()
	serviceConfig.VerificationFilename = plan.ServiceConfig.VerificationFilename.ValueStringPointer()

	drain, _, err := r.data.client.LogDrainsAPI.
		LogDrainsCreate(ctx, plan.SiteID.ValueString()).
		LogDrain(netlifyapi.LogDrain{
			SiteId:        plan.SiteID.ValueString(),
			Destination:   plan.Destination.ValueString(),
			Format:        plan.Format.ValueString(),
			LogTypes:      logTypes,
			ExcludePii:    plan.ExcludePII.ValueBoolPointer(),
			ServiceConfig: serviceConfig,
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating log drain",
			fmt.Sprintf("Could not create log drain: %q", err.Error()),
		)
		return
	}

	plan.ID = types.StringValue(drain.Id)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *logDrainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state logDrainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	drain, _, err := r.data.client.LogDrainsAPI.LogDrainsShow(ctx, state.ID.ValueString(), state.SiteID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading log drain",
			fmt.Sprintf(
				"Could not read log drain %q (site ID: %q): %q",
				state.ID.ValueString(),
				state.SiteID.ValueString(),
				err.Error(),
			),
		)
		return
	}
	state.Destination = types.StringValue(drain.Destination)
	state.Format = types.StringValue(drain.Format)
	var logTypesArr []string
	err = json.Unmarshal([]byte(drain.LogTypes), &logTypesArr)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading log drain",
			fmt.Sprintf(
				"Could not read log drain %q (site ID: %q): %q",
				state.ID.ValueString(),
				state.SiteID.ValueString(),
				err.Error(),
			),
		)
		return
	}
	state.LogTypes = make([]types.String, len(logTypesArr))
	for i, logType := range logTypesArr {
		state.LogTypes[i] = types.StringValue(logType)
	}
	if drain.ExcludePii == nil {
		state.ExcludePII = types.BoolValue(false)
	} else {
		state.ExcludePII = types.BoolValue(*drain.ExcludePii)
	}
	state.ServiceConfig = &logDrainServiceConfigModel{}
	state.ServiceConfig.URL = types.StringPointerValue(drain.ServiceConfig.Url)
	if len(drain.ServiceConfig.Tags) > 0 {
		state.ServiceConfig.Tags = make(map[string]types.String)
		for k, v := range drain.ServiceConfig.Tags {
			state.ServiceConfig.Tags[k] = types.StringValue(v)
		}
	}
	state.ServiceConfig.IntegrationName = types.StringPointerValue(drain.ServiceConfig.IntegrationName)
	state.ServiceConfig.AuthHeader = types.StringPointerValue(drain.ServiceConfig.AuthorizationHeader)
	state.ServiceConfig.BucketName = types.StringPointerValue(drain.ServiceConfig.BucketName)
	state.ServiceConfig.BucketRegion = types.StringPointerValue(drain.ServiceConfig.BucketRegion)
	state.ServiceConfig.Path = types.StringPointerValue(drain.ServiceConfig.Path)
	state.ServiceConfig.VerificationFilename = types.StringPointerValue(drain.ServiceConfig.VerificationFilename)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *logDrainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan logDrainResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var logTypes string
	logTypesArr := make([]string, len(plan.LogTypes))
	for i, logType := range plan.LogTypes {
		logTypesArr[i] = logType.ValueString()
	}
	if logTypesBytes, err := json.Marshal(logTypesArr); err != nil {
		resp.Diagnostics.AddError(
			"Error updating log drain",
			fmt.Sprintf("Could not marshal log types: %q", err.Error()),
		)
		return
	} else {
		logTypes = string(logTypesBytes)
	}

	var serviceConfig netlifyapi.LogDrainServiceConfig
	if plan.ServiceConfig == nil {
		plan.ServiceConfig = &logDrainServiceConfigModel{}
	}
	serviceConfig.Url = plan.ServiceConfig.URL.ValueStringPointer()
	if len(plan.ServiceConfig.Tags) > 0 {
		serviceConfig.Tags = make(map[string]string)
		for k, v := range plan.ServiceConfig.Tags {
			serviceConfig.Tags[k] = v.ValueString()
		}
	}
	serviceConfig.IntegrationName = plan.ServiceConfig.IntegrationName.ValueStringPointer()
	serviceConfig.AuthorizationHeader = plan.ServiceConfig.AuthHeader.ValueStringPointer()
	serviceConfig.BucketName = plan.ServiceConfig.BucketName.ValueStringPointer()
	serviceConfig.BucketRegion = plan.ServiceConfig.BucketRegion.ValueStringPointer()
	serviceConfig.Path = plan.ServiceConfig.Path.ValueStringPointer()
	serviceConfig.VerificationFilename = plan.ServiceConfig.VerificationFilename.ValueStringPointer()

	drain, _, err := r.data.client.LogDrainsAPI.
		LogDrainsUpdate(ctx, plan.ID.ValueString(), plan.SiteID.ValueString()).
		LogDrain(netlifyapi.LogDrain{
			Id:            plan.ID.ValueString(),
			SiteId:        plan.SiteID.ValueString(),
			Destination:   plan.Destination.ValueString(),
			Format:        plan.Format.ValueString(),
			LogTypes:      logTypes,
			ExcludePii:    plan.ExcludePII.ValueBoolPointer(),
			ServiceConfig: serviceConfig,
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating log drain",
			fmt.Sprintf("Could not update log drain %q (site ID: %q): %q", plan.ID.ValueString(), plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	plan.ID = types.StringValue(drain.Id)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *logDrainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state logDrainResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.client.LogDrainsAPI.LogDrainsDestroy(ctx, state.ID.ValueString(), state.SiteID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting log drain",
			fmt.Sprintf(
				"Could not delete log drain %q (site ID: %q): %q",
				state.ID.ValueString(),
				state.SiteID.ValueString(),
				err.Error(),
			),
		)
		return
	}
}

func (r *logDrainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ":")

	errorMessage := fmt.Sprintf("Expected import identifier in the formats: site_id,log_drain_id. Got: %q", req.ID)

	if len(idParts) == 2 {
		if idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	} else {
		resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
		return
	}
}
