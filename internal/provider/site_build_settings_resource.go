package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ resource.Resource                = &siteBuildSettingsResource{}
	_ resource.ResourceWithConfigure   = &siteBuildSettingsResource{}
	_ resource.ResourceWithImportState = &siteBuildSettingsResource{}
)

func NewSiteBuildSettingsResource() resource.Resource {
	return &siteBuildSettingsResource{}
}

type siteBuildSettingsResource struct {
	data NetlifyProviderData
}

type siteBuildSettingsResourceModel struct {
	SiteID             types.String `tfsdk:"site_id"`
	LastUpdated        types.String `tfsdk:"last_updated"`
	BaseDirectory      types.String `tfsdk:"base_directory"`
	PackageDirectory   types.String `tfsdk:"package_directory"`
	BuildCommand       types.String `tfsdk:"build_command"`
	PublishDirectory   types.String `tfsdk:"publish_directory"`
	FunctionsDirectory types.String `tfsdk:"functions_directory"`
	StopBuilds         types.Bool   `tfsdk:"stop_builds"`
	// Runtime            types.String `tfsdk:"runtime"`             // ?!?! is this plugins.package?

	BuildImage types.String `tfsdk:"build_image"`
	// NodeJSVersion types.String `tfsdk:"node_js_version"` // versions.node.active / default: versions.node.active or versions.node.default

	PrettyURLs types.Bool `tfsdk:"pretty_urls"`
}

func (r *siteBuildSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_build_settings"
}

func (r *siteBuildSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *siteBuildSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"site_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			// "runtime": schema.StringAttribute{
			// 	Optional: true,
			// },
			"base_directory": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"package_directory": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"build_command": schema.StringAttribute{
				Required: true,
			},
			"publish_directory": schema.StringAttribute{
				Required: true,
			},
			"functions_directory": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"stop_builds": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"build_image": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// "node_js_version": schema.StringAttribute{
			// 	Optional: true,
			// 	Computed: true,
			// 	PlanModifiers: []planmodifier.String{
			// 		stringplanmodifier.UseStateForUnknown(),
			// 	},
			// },
			"pretty_urls": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
		},
	}
}

func (r *siteBuildSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteBuildSettingsResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.write(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteBuildSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteBuildSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.read(ctx, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteBuildSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan siteBuildSettingsResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.write(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteBuildSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state siteBuildSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning("Site build settings are now unmanaged.", "Site build settings are now unmanaged. The site will continue to build with the last settings.")
}

func (r *siteBuildSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("site_id"), req, resp)
}

func (r *siteBuildSettingsResource) read(ctx context.Context, state *siteBuildSettingsResourceModel, diagnostics *diag.Diagnostics) {
	site, _, err := r.data.client.SitesAPI.GetSite(ctx, state.SiteID.ValueString()).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error reading site build settings",
			fmt.Sprintf("Could not read site build settings for site %q: %q", state.SiteID.ValueString(), err.Error()),
		)
		return
	}

	state.BaseDirectory = types.StringPointerValue(site.BuildSettings.Base)
	state.PackageDirectory = types.StringPointerValue(site.BuildSettings.PackagePath)
	state.BuildCommand = types.StringPointerValue(site.BuildSettings.Cmd)
	state.PublishDirectory = types.StringPointerValue(site.BuildSettings.Dir)
	state.FunctionsDirectory = types.StringPointerValue(site.BuildSettings.FunctionsDir)
	state.StopBuilds = types.BoolPointerValue(site.BuildSettings.StopBuilds)
	state.BuildImage = types.StringValue(site.BuildImage)
	state.PrettyURLs = types.BoolPointerValue(site.ProcessingSettings.Html.PrettyUrls)
}

func (r *siteBuildSettingsResource) write(ctx context.Context, plan *siteBuildSettingsResourceModel, diagnostics *diag.Diagnostics) {
	var curState siteBuildSettingsResourceModel
	curState.SiteID = plan.SiteID
	r.read(ctx, &curState, diagnostics)
	if diagnostics.HasError() {
		return
	}

	site := netlifyapi.PartialSite{
		BuildSettings: &netlifyapi.Repo{
			Base:         plan.BaseDirectory.ValueStringPointer(),
			PackagePath:  plan.PackageDirectory.ValueStringPointer(),
			Cmd:          plan.BuildCommand.ValueStringPointer(),
			Dir:          plan.PublishDirectory.ValueStringPointer(),
			FunctionsDir: plan.FunctionsDirectory.ValueStringPointer(),
			StopBuilds:   plan.StopBuilds.ValueBoolPointer(),
		},
		ProcessingSettings: &netlifyapi.SiteProcessingSettings{
			Html: &netlifyapi.SiteProcessingSettingsHtml{
				PrettyUrls: plan.PrettyURLs.ValueBoolPointer(),
			},
		},
	}

	if plan.BuildImage.IsUnknown() || plan.BuildImage.IsNull() {
		site.BuildImage = curState.BuildImage.ValueStringPointer()
	} else {
		site.BuildImage = plan.BuildImage.ValueStringPointer()
	}

	_, _, err := r.data.client.SitesAPI.
		UpdateSite(ctx, plan.SiteID.ValueString()).
		PartialSite(site).
		Execute()
	if err != nil {
		diagnostics.AddError(
			"Error updating site build settings",
			fmt.Sprintf("Could not update site build settings for site %q: %q", plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
