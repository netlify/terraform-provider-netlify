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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ resource.Resource                = &siteDomainSettingsResource{}
	_ resource.ResourceWithConfigure   = &siteDomainSettingsResource{}
	_ resource.ResourceWithImportState = &siteDomainSettingsResource{}
)

func NewSiteDomainSettingsResource() resource.Resource {
	return &siteDomainSettingsResource{}
}

type siteDomainSettingsResource struct {
	data NetlifyProviderData
}

type siteDomainSettingsResourceModel struct {
	SiteID      types.String `tfsdk:"site_id"`
	LastUpdated types.String `tfsdk:"last_updated"`

	CustomDomain              types.String   `tfsdk:"custom_domain"`
	DomainAliases             []types.String `tfsdk:"domain_aliases"`
	BranchDeployCustomDomain  types.String   `tfsdk:"branch_deploy_custom_domain"`
	DeployPreviewCustomDomain types.String   `tfsdk:"deploy_preview_custom_domain"`
}

func (r *siteDomainSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site_domain_settings"
}

func (r *siteDomainSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *siteDomainSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

func (r *siteDomainSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan siteDomainSettingsResourceModel
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

func (r *siteDomainSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state siteDomainSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	site, _, err := r.data.client.SitesAPI.GetSite(ctx, state.SiteID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading site domain settings",
			fmt.Sprintf("Could not read site domain settings for site %q: %q", state.SiteID.ValueString(), err.Error()),
		)
		return
	}

	if site.CustomDomain != "" {
		state.CustomDomain = types.StringValue(site.CustomDomain)
	} else {
		state.CustomDomain = types.StringNull()
	}
	state.DomainAliases = make([]types.String, len(site.DomainAliases))
	for i, domainAlias := range site.DomainAliases {
		state.DomainAliases[i] = types.StringValue(domainAlias)
	}
	if site.BranchDeployCustomDomain != "" {
		state.BranchDeployCustomDomain = types.StringValue(site.BranchDeployCustomDomain)
	} else {
		state.BranchDeployCustomDomain = types.StringNull()
	}

	if site.DeployPreviewCustomDomain != "" {
		state.DeployPreviewCustomDomain = types.StringValue(site.DeployPreviewCustomDomain)
	} else {
		state.DeployPreviewCustomDomain = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *siteDomainSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan siteDomainSettingsResourceModel
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

func (r *siteDomainSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state siteDomainSettingsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.AddWarning("Site domain settings are now unmanaged.", "Site domain settings are now unmanaged. The site will continue to deploy with the last settings.")
}

func (r *siteDomainSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("site_id"), req, resp)
}

func (r *siteDomainSettingsResource) write(ctx context.Context, plan *siteDomainSettingsResourceModel, diagnostics *diag.Diagnostics) {
	domainAliases := make([]string, len(plan.DomainAliases))
	for i, domainAlias := range plan.DomainAliases {
		domainAliases[i] = domainAlias.ValueString()
	}
	site := netlifyapi.PartialSite{
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
			"Error updating site domain settings",
			fmt.Sprintf("Could not update site domain settings for site %q: %q", plan.SiteID.ValueString(), err.Error()),
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}
