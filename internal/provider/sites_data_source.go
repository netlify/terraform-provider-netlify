package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
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
	TeamSlug types.String     `tfsdk:"team_slug"`
	Sites    []sitesSiteModel `tfsdk:"sites"`
}

type sitesSiteModel struct {
	ID            types.String   `tfsdk:"id"`
	TeamSlug      types.String   `tfsdk:"team_slug"`
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
			"team_slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Required if a default team was not configured in the provider configuration.",
			},
			"sites": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"team_slug": schema.StringAttribute{
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

	teamSlug := d.data.teamSlugOrDefault(config.TeamSlug)
	if teamSlug == nil {
		resp.Diagnostics.AddError(
			"Missing team slug",
			"Team slug is required for reading Netlify sites. Please provide a team slug in the plan or configure a default team in the provider configuration.",
		)
		return
	}

	r := d.data.client.SitesAPI.
		ListSitesForAccount(ctx, *teamSlug).
		PerPage(100)
	sites := make([]netlifyapi.Site, 0)
	var page int64 = 1
	for {
		items, _, err := r.Page(page).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify team", fmt.Sprintf("Could not list Netlify sites in team %q: %q", *teamSlug, err.Error()))
			return
		}
		if len(items) == 0 {
			break
		}
		sites = append(sites, items...)
		page++
	}
	config.Sites = make([]sitesSiteModel, len(sites))
	for i, site := range sites {
		config.Sites[i] = sitesSiteModel{
			ID:            types.StringValue(site.Id),
			TeamSlug:      types.StringValue(site.AccountSlug),
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
