package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
	"github.com/netlify/terraform-provider-netlify/internal/provider/netlify_validators"
)

var (
	_ resource.Resource                = &siteDeploySettingsResource{}
	_ resource.ResourceWithConfigure   = &siteDeploySettingsResource{}
	_ resource.ResourceWithImportState = &siteDeploySettingsResource{}
)

func NewSiteDeploySettingsResource() resource.Resource {
	return &siteDeploySettingsResource{}
}

type siteDeploySettingsResource struct {
	data NetlifyProviderData
}

type siteDeploySettingsResourceModel struct {
	SiteID          types.String `tfsdk:"site_id"`
	LastUpdated     types.String `tfsdk:"last_updated"`
	FunctionsRegion types.String `tfsdk:"functions_region"`

	ProductionBranch              types.String   `tfsdk:"production_branch"`
	BranchDeployAllBranches       types.Bool     `tfsdk:"branch_deploy_all_branches"`
	BranchDeployBranches          []types.String `tfsdk:"branch_deploy_branches"`
	DeployPreviews                types.Bool     `tfsdk:"deploy_previews"`
	NetlifyDrawerInDeployPreviews types.Bool     `tfsdk:"netlify_drawer_in_deploy_previews"`
	NetlifyDrawerInBranchDeploys  types.Bool     `tfsdk:"netlify_drawer_in_branch_deploys"`
	NetlifyHeadsUpDisplay         types.Bool     `tfsdk:"netlify_heads_up_display"`

	CustomDomain              types.String   `tfsdk:"custom_domain"`
	DomainAliases             []types.String `tfsdk:"domain_aliases"`
	BranchDeployCustomDomain  types.String   `tfsdk:"branch_deploy_custom_domain"`
	DeployPreviewCustomDomain types.String   `tfsdk:"deploy_preview_custom_domain"`
}

func (r *siteDeploySettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_deploy_settings"
}

func (r *siteDeploySettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *siteDeploySettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	emptyList, diags := types.ListValue(types.StringType, []attr.Value{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
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
			"functions_region": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("us-east-2"),
			},
			"production_branch": schema.StringAttribute{
				Required: true,
			},
			"branch_deploy_all_branches": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"branch_deploy_branches": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Default:     listdefault.StaticValue(emptyList),
				Validators: []validator.List{
					netlify_validators.SiteDeploySettingsDeployBranchesValidator{
						AllBranchesPathExpression: path.MatchRoot("branch_deploy_all_branches"),
					},
				},
			},
			"deploy_previews": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"netlify_drawer_in_deploy_previews": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"netlify_drawer_in_branch_deploys": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"netlify_heads_up_display": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"custom_domain": schema.StringAttribute{
				Optional: true,
			},
			"domain_aliases": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Default:     listdefault.StaticValue(emptyList),
			},
			"branch_deploy_custom_domain": schema.StringAttribute{
				Optional: true,
			},
			"deploy_preview_custom_domain": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *siteDeploySettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteDeploySettingsResourceModel
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

func (r *siteDeploySettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteDeploySettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, _, err := r.data.client.SitesAPI.GetSite(ctx, state.SiteID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading site deploy settings",
			fmt.Sprintf("Could not read site deploy settings for site %q: %q", state.SiteID.ValueString(), err.Error()),
		)
		return
	}

	state.FunctionsRegion = types.StringPointerValue(site.FunctionsRegion)
	state.ProductionBranch = types.StringPointerValue(site.BuildSettings.RepoBranch)
	allowedBranchedLen := len(site.BuildSettings.AllowedBranches)
	state.BranchDeployAllBranches = types.BoolValue(allowedBranchedLen == 0)
	if allowedBranchedLen == 0 {
		state.BranchDeployBranches = make([]types.String, 0)
	} else {
		state.BranchDeployBranches = make([]types.String, 0, allowedBranchedLen-1)
		for _, branch := range site.BuildSettings.AllowedBranches {
			if branch != *site.BuildSettings.RepoBranch {
				state.BranchDeployBranches = append(state.BranchDeployBranches, types.StringValue(branch))
			}
		}
	}
	if site.BuildSettings.SkipPrs == nil {
		state.DeployPreviews = types.BoolValue(true)
	} else {
		state.DeployPreviews = types.BoolValue(!*site.BuildSettings.SkipPrs)
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
	state.CustomDomain = types.StringValue(site.CustomDomain)
	state.DomainAliases = make([]types.String, len(site.DomainAliases))
	for i, domainAlias := range site.DomainAliases {
		state.DomainAliases[i] = types.StringValue(domainAlias)
	}
	state.BranchDeployCustomDomain = types.StringValue(site.BranchDeployCustomDomain)
	state.DeployPreviewCustomDomain = types.StringValue(site.DeployPreviewCustomDomain)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteDeploySettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan siteDeploySettingsResourceModel
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

func (r *siteDeploySettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state siteDeploySettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning("Site deploy settings are now unmanaged.", "Site deploy settings are now unmanaged. The site will continue to deploy with the last settings.")
}

func (r *siteDeploySettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("site_id"), req, resp)
}

func (r *siteDeploySettingsResource) write(ctx context.Context, plan *siteDeploySettingsResourceModel, diagnostics *diag.Diagnostics) {
	allowedBranches := make([]string, 0, len(plan.BranchDeployBranches)+1)
	if !plan.BranchDeployAllBranches.ValueBool() {
		allowedBranches = append(allowedBranches, plan.ProductionBranch.ValueString())
		for _, branch := range plan.BranchDeployBranches {
			allowedBranches = append(allowedBranches, branch.ValueString())
		}
	}
	skipPrs := !plan.DeployPreviews.ValueBool()
	cdpContexts := make([]string, 0, 2)
	if plan.NetlifyDrawerInDeployPreviews.ValueBool() {
		cdpContexts = append(cdpContexts, "deploy-preview")
	}
	if plan.NetlifyDrawerInBranchDeploys.ValueBool() {
		cdpContexts = append(cdpContexts, "branch-deploy")
	}
	domainAliases := make([]string, len(plan.DomainAliases))
	for i, domainAlias := range plan.DomainAliases {
		domainAliases[i] = domainAlias.ValueString()
	}
	site := netlifyapi.PartialSite{
		FunctionsRegion: plan.FunctionsRegion.ValueStringPointer(),
		BuildSettings: &netlifyapi.Repo{
			RepoBranch:      plan.ProductionBranch.ValueStringPointer(),
			AllowedBranches: allowedBranches,
			SkipPrs:         &skipPrs,
		},
		CdpEnabledContexts:        cdpContexts,
		HudEnabled:                plan.NetlifyHeadsUpDisplay.ValueBoolPointer(),
		CustomDomain:              plan.CustomDomain.ValueStringPointer(),
		DomainAliases:             domainAliases,
		BranchDeployCustomDomain:  plan.BranchDeployCustomDomain.ValueStringPointer(),
		DeployPreviewCustomDomain: plan.DeployPreviewCustomDomain.ValueStringPointer(),
	}

	_, _, err := r.data.client.SitesAPI.
		UpdateSite(ctx, plan.SiteID.ValueString()).
		PartialSite(site).
		Execute()
	if err != nil {
		diagnostics.AddError(
			"Error updating site deploy settings",
			fmt.Sprintf("Could not update site deploy settings for site %q: %q", plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
