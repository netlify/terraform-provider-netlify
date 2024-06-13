package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
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

type siteModel struct {
	ID            types.String   `tfsdk:"id"`
	TeamSlug      types.String   `tfsdk:"team_slug"`
	Name          types.String   `tfsdk:"name"`
	CustomDomain  types.String   `tfsdk:"custom_domain"`
	DomainAliases []types.String `tfsdk:"domain_aliases"`
	GitDeployKey  types.String   `tfsdk:"git_deploy_key"`
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
				Validators: []validator.String{
					stringvalidator.AtLeastOneOf(path.MatchRoot("name")),
				},
			},
			"team_slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.AlsoRequires(path.MatchRoot("team_slug")),
				},
			},
			"custom_domain": schema.StringAttribute{
				Computed: true,
			},
			"domain_aliases": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"git_deploy_key": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *siteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config siteModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var site *netlifyapi.Site
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		var err error
		site, _, err = d.data.client.SitesAPI.GetSite(ctx, config.ID.ValueString()).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify site", fmt.Sprintf("Could not read Netlify site ID %q: %q",
				config.ID.ValueString(), err.Error()))
			return
		}
	} else {
		sites, _, err := d.data.client.SitesAPI.
			ListSitesForAccount(ctx, config.TeamSlug.ValueString()).
			Name(config.Name.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify team", fmt.Sprintf("Could not list Netlify sites in team %q: %q", config.TeamSlug.ValueString(), err.Error()))
			return
		}
		nameString := config.Name.ValueString()
		for _, s := range sites {
			if s.Name == nameString {
				sit := s
				site = &sit
				break
			}
		}
		if site == nil {
			resp.Diagnostics.AddError("Error reading Netlify team", fmt.Sprintf("Could not find Netlify site with name %q in team %q", nameString, config.TeamSlug.ValueString()))
			return
		}
	}

	var siteDeployKey *netlifyapi.DeployKey
	if site.BuildSettings.DeployKeyId != nil {
		var err error
		siteDeployKey, _, err = d.data.client.DeployKeysAPI.GetDeployKey(ctx, *site.BuildSettings.DeployKeyId).Execute()
		if siteDeployKey == nil || err != nil {
			resp.Diagnostics.AddError(
				"Error reading site deploy key",
				fmt.Sprintf("Could not read site deploy key for site %q: %q", site.Id, err.Error()),
			)
			return
		}
	}

	config.ID = types.StringValue(site.Id)
	config.TeamSlug = types.StringValue(site.AccountSlug)
	config.Name = types.StringValue(site.Name)
	config.CustomDomain = types.StringValue(site.CustomDomain)
	config.DomainAliases = make([]types.String, len(site.DomainAliases))
	for i, alias := range site.DomainAliases {
		config.DomainAliases[i] = types.StringValue(alias)
	}
	if siteDeployKey != nil {
		config.GitDeployKey = types.StringValue(siteDeployKey.PublicKey)
	} else {
		config.GitDeployKey = types.StringNull()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
