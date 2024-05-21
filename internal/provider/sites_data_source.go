package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &sitesDataSource{}
	_ datasource.DataSourceWithConfigure = &sitesDataSource{}
)

func NewSitesDataSource() datasource.DataSource {
	return &sitesDataSource{}
}

type sitesDataSource struct {
	data NetlifyProviderData
}

type sitesDataSourceModel struct {
	AccountSlug types.String     `tfsdk:"account_slug"`
	Sites       []sitesSiteModel `tfsdk:"sites"`
}

type sitesSiteModel struct {
	ID            types.String   `tfsdk:"id"`
	AccountSlug   types.String   `tfsdk:"account_slug"`
	Name          types.String   `tfsdk:"name"`
	CustomDomain  types.String   `tfsdk:"custom_domain"`
	DomainAliases []types.String `tfsdk:"domain_aliases"`
}

func (d *sitesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *sitesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sites"
}

func (d *sitesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_slug": schema.StringAttribute{
				Required: true,
			},
			"sites": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"account_slug": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
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
				},
			},
		},
	}
}

func (d *sitesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config sitesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sites, _, err := d.data.client.SitesAPI.ListSitesForAccount(ctx, config.AccountSlug.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not list Netlify sites in account %q: %q", config.AccountSlug.ValueString(), err.Error()))
		return
	}
	config.Sites = make([]sitesSiteModel, len(sites))
	for i, site := range sites {
		config.Sites[i] = sitesSiteModel{
			ID:            types.StringValue(site.Id),
			AccountSlug:   types.StringValue(site.AccountSlug),
			Name:          types.StringValue(site.Name),
			CustomDomain:  types.StringValue(site.CustomDomain),
			DomainAliases: make([]types.String, len(site.DomainAliases)),
		}
		for j, alias := range site.DomainAliases {
			config.Sites[i].DomainAliases[j] = types.StringValue(alias)
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
