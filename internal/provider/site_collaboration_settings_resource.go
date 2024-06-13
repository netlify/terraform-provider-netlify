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
	_ resource.Resource                = &siteCollaborationSettingsResource{}
	_ resource.ResourceWithConfigure   = &siteCollaborationSettingsResource{}
	_ resource.ResourceWithImportState = &siteCollaborationSettingsResource{}
)

func NewSiteCollaborationSettingsResource() resource.Resource {
	return &siteCollaborationSettingsResource{}
}

type siteCollaborationSettingsResource struct {
	data NetlifyProviderData
}

type siteCollaborationSettingsResourceModel struct {
	SiteID      types.String `tfsdk:"site_id"`
	LastUpdated types.String `tfsdk:"last_updated"`

	NetlifyDrawerInDeployPreviews types.Bool `tfsdk:"netlify_drawer_in_deploy_previews"`
	NetlifyDrawerInBranchDeploys  types.Bool `tfsdk:"netlify_drawer_in_branch_deploys"`
	NetlifyHeadsUpDisplay         types.Bool `tfsdk:"netlify_heads_up_display"`
}

func (r *siteCollaborationSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_collaboration_settings"
}

func (r *siteCollaborationSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *siteCollaborationSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"netlify_drawer_in_deploy_previews": schema.BoolAttribute{
				Required: true,
			},
			"netlify_drawer_in_branch_deploys": schema.BoolAttribute{
				Required: true,
			},
			"netlify_heads_up_display": schema.BoolAttribute{
				Required: true,
			},
		},
	}
}

func (r *siteCollaborationSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteCollaborationSettingsResourceModel
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

func (r *siteCollaborationSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteCollaborationSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, _, err := r.data.client.SitesAPI.GetSite(ctx, state.SiteID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading site collaboration settings",
			fmt.Sprintf("Could not read site collaboration settings for site %q: %q", state.SiteID.ValueString(), err.Error()),
		)
		return
	}

	state.NetlifyDrawerInDeployPreviews = types.BoolValue(false)
	state.NetlifyDrawerInBranchDeploys = types.BoolValue(false)
	for _, context := range site.CdpEnabledContexts {
		if context == "deploy-preview" {
			state.NetlifyDrawerInDeployPreviews = types.BoolValue(true)
		} else if context == "branch-deploy" {
			state.NetlifyDrawerInBranchDeploys = types.BoolValue(true)
		}
	}
	state.NetlifyHeadsUpDisplay = types.BoolPointerValue(site.HudEnabled)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteCollaborationSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan siteCollaborationSettingsResourceModel
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

func (r *siteCollaborationSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state siteCollaborationSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning("Site collaboration settings are now unmanaged.", "Site collaboration settings are now unmanaged. The site will continue to deploy with the last settings.")
}

func (r *siteCollaborationSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("site_id"), req, resp)
}

func (r *siteCollaborationSettingsResource) write(ctx context.Context, plan *siteCollaborationSettingsResourceModel, diagnostics *diag.Diagnostics) {
	cdpContexts := make([]string, 0, 2)
	if plan.NetlifyDrawerInDeployPreviews.ValueBool() {
		cdpContexts = append(cdpContexts, "deploy-preview")
	}
	if plan.NetlifyDrawerInBranchDeploys.ValueBool() {
		cdpContexts = append(cdpContexts, "branch-deploy")
	}
	site := netlifyapi.PartialSite{
		CdpEnabledContexts: cdpContexts,
		HudEnabled:         plan.NetlifyHeadsUpDisplay.ValueBoolPointer(),
	}

	_, _, err := r.data.client.SitesAPI.
		UpdateSite(ctx, plan.SiteID.ValueString()).
		PartialSite(site).
		Execute()
	if err != nil {
		diagnostics.AddError(
			"Error updating site collaboration settings",
			fmt.Sprintf("Could not update site collaboration settings for site %q: %q", plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
