package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/models"
	"github.com/netlify/terraform-provider-netlify/internal/plumbing/operations"
)

var (
	_ datasource.DataSource              = &siteDataSource{}
	_ datasource.DataSourceWithConfigure = &siteDataSource{}
)

func NewSiteDataSource() datasource.DataSource {
	return &siteDataSource{}
}

type siteDataSource struct {
	data NetlifyProviderData
}

type NetlifySiteModel struct {
	ID            types.String   `tfsdk:"id"`
	AccountSlug   types.String   `tfsdk:"account_slug"`
	Name          types.String   `tfsdk:"name"`
	CustomDomain  types.String   `tfsdk:"custom_domain"`
	DomainAliases []types.String `tfsdk:"domain_aliases"`
}

func (d *siteDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(NetlifyProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected NetlifyProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.data = data
}

func (d *siteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_site"
}

func (d *siteDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"account_slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"custom_domain": schema.StringAttribute{
				Computed: true,
			},
			"domain_aliases": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (d *siteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config NetlifySiteModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if (config.ID.IsUnknown() || config.ID.IsNull()) &&
		(config.AccountSlug.IsUnknown() || config.AccountSlug.IsNull() || config.Name.IsUnknown() || config.Name.IsNull()) {
		resp.Diagnostics.AddError("Error reading Netlify site", "Either id, or account slug and site name, must be specified for a site search")
		return
	}

	var site *models.Site
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		siteOk, err := d.data.client.Operations.GetSite(
			operations.NewGetSiteParams().WithSiteID(config.ID.ValueString()),
			d.data.authInfo,
		)
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify site", fmt.Sprintf("Could not read Netlify site ID %q: %q",
				config.ID.ValueString(), err.Error()))
			return
		}
		site = siteOk.Payload
	} else {
		sitesOk, err := d.data.client.Operations.ListSitesForAccount(
			operations.NewListSitesForAccountParams().WithAccountSlug(config.AccountSlug.ValueString()),
			d.data.authInfo,
		)
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not list Netlify sites in account %q: %q", config.AccountSlug.ValueString(), err.Error()))
			return
		}
		nameString := config.Name.ValueString()
		for _, a := range sitesOk.Payload {
			if a.Name == nameString {
				site = a
				break
			}
		}
		if site == nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not find Netlify site with name %q in account %q", nameString, config.AccountSlug.ValueString()))
			return
		}
	}

	config.ID = types.StringValue(site.ID)
	config.AccountSlug = types.StringValue(site.AccountSlug)
	config.Name = types.StringValue(site.Name)
	config.CustomDomain = types.StringValue(site.CustomDomain)
	config.DomainAliases = make([]types.String, len(site.DomainAliases))
	for i, alias := range site.DomainAliases {
		config.DomainAliases[i] = types.StringValue(alias)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
