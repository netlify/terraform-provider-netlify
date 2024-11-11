package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ resource.Resource                = &siteMetricsSettingsResource{}
	_ resource.ResourceWithConfigure   = &siteMetricsSettingsResource{}
	_ resource.ResourceWithImportState = &siteMetricsSettingsResource{}
)

func NewSiteMetricsSettingsResource() resource.Resource {
	return &siteMetricsSettingsResource{}
}

type siteMetricsSettingsResource struct {
	data NetlifyProviderData
}

type siteMetricsSettingsResourceModel struct {
	SiteID      types.String `tfsdk:"site_id"`
	LastUpdated types.String `tfsdk:"last_updated"`

	SiteAnalytics   types.Bool `tfsdk:"site_analytics"`
	RealUserMetrics types.Bool `tfsdk:"real_user_metrics"`
}

func (r *siteMetricsSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_metrics_settings"
}

func (r *siteMetricsSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *siteMetricsSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"site_analytics": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable site analytics. Warning: This might incur a cost on certain plans. Note: You must wait 10 minutes before disabling after enabling.",
			},
			"real_user_metrics": schema.BoolAttribute{
				Optional:    true,
				Description: "Enable real user metrics. Warning: This might incur a cost on certain plans.",
			},
		},
	}
}

func (r *siteMetricsSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteMetricsSettingsResourceModel
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

func (r *siteMetricsSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteMetricsSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, _, err := r.data.client.SitesAPI.GetSite(ctx, state.SiteID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading site metrics settings",
			fmt.Sprintf("Could not read site metrics settings for site %q: %q", state.SiteID.ValueString(), err.Error()),
		)
		return
	}

	if site.AnalyticsInstanceId != nil && *site.AnalyticsInstanceId != "" {
		state.SiteAnalytics = types.BoolValue(true)
	} else {
		state.SiteAnalytics = types.BoolValue(false)
	}
	if site.RumEnabled != nil && *site.RumEnabled {
		state.RealUserMetrics = types.BoolValue(true)
	} else {
		state.RealUserMetrics = types.BoolValue(false)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteMetricsSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan siteMetricsSettingsResourceModel
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

func (r *siteMetricsSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state siteMetricsSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning("Site metrics settings are now unmanaged.", "Site metrics settings are now unmanaged. The site will continue to deploy with the last settings.")
}

func (r *siteMetricsSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("site_id"), req, resp)
}

func (r *siteMetricsSettingsResource) write(ctx context.Context, plan *siteMetricsSettingsResourceModel, diagnostics *diag.Diagnostics) {
	psite := netlifyapi.PartialSite{
		RumEnabled: plan.RealUserMetrics.ValueBoolPointer(),
	}

	_, _, err := r.data.client.SitesAPI.
		UpdateSite(ctx, plan.SiteID.ValueString()).
		PartialSite(psite).
		Execute()
	if err != nil {
		diagnostics.AddError(
			"Error updating site metrics settings",
			fmt.Sprintf("Could not update site metrics settings for site %q: %q", plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	site, _, err := r.data.client.SitesAPI.GetSite(ctx, plan.SiteID.ValueString()).Execute()
	if err != nil {
		diagnostics.AddError(
			"Error reading site metrics settings",
			fmt.Sprintf("Could not read site metrics settings for site %q: %q", plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	if !plan.SiteAnalytics.IsNull() {
		analyticsEnabled := false
		if site.AnalyticsInstanceId != nil && *site.AnalyticsInstanceId != "" {
			analyticsEnabled = true
		}

		if !analyticsEnabled && plan.SiteAnalytics.ValueBool() {
			_, _, err := r.data.client.AnalyticsAPI.EnableAnalytics(ctx, plan.SiteID.ValueString()).Execute()
			if err != nil {
				diagnostics.AddError(
					"Error enabling site analytics",
					fmt.Sprintf("Could not enable site analytics for site %q: %q", plan.SiteID.ValueString(), err.Error()),
				)
				return
			}
		}
		if analyticsEnabled && !plan.SiteAnalytics.ValueBool() {
			_, err := r.data.client.AnalyticsAPI.DisableAnalytics(ctx, *site.AnalyticsInstanceId, plan.SiteID.ValueString()).Execute()
			if err != nil {
				diagnostics.AddError(
					"Error disabling site analytics",
					fmt.Sprintf("Could not disable site analytics for site %q: %q", plan.SiteID.ValueString(), err.Error()),
				)
				return
			}
		}
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
