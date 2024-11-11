package provider

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var _ provider.Provider = &NetlifyProvider{}

type NetlifyProvider struct {
	version string
	client  *netlifyapi.APIClient
}

type NetlifyProviderModel struct {
	Endpoint        types.String `tfsdk:"endpoint"`
	Token           types.String `tfsdk:"token"`
	DefaultTeamID   types.String `tfsdk:"default_team_id"`
	DefaultTeamSlug types.String `tfsdk:"default_team_slug"`
}

type NetlifyProviderData struct {
	client          *netlifyapi.APIClient
	defaultTeamId   *string
	defaultTeamSlug *string
}

func (d *NetlifyProviderData) teamIdOrDefault(value types.String) *string {
	if value.IsUnknown() || value.IsNull() {
		return d.defaultTeamId
	}
	return value.ValueStringPointer()
}

func (d *NetlifyProviderData) teamSlugOrDefault(value types.String) *string {
	if value.IsUnknown() || value.IsNull() {
		return d.defaultTeamSlug
	}
	return value.ValueStringPointer()
}

func (p *NetlifyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netlify"
	resp.Version = p.version
}

func (p *NetlifyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Defaults to: https://api.netlify.com",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "Read: https://docs.netlify.com/api/get-started/#authentication , will use the `NETLIFY_API_TOKEN` environment variable if not set.",
				Optional:            true,
				Sensitive:           true,
			},
			"default_team_id": schema.StringAttribute{
				MarkdownDescription: "The default team ID to use for resources that require a team ID or a team slug. Warning: Changing this value may not trigger recreation of resources.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.ConflictsWith(path.MatchRoot("default_team_slug"))},
			},
			"default_team_slug": schema.StringAttribute{
				MarkdownDescription: "The default team slug to use for resources that require a team ID or a team slug. Warning: Changing this value may not trigger recreation of resources.",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.ConflictsWith(path.MatchRoot("default_team_id"))},
			},
		},
	}
}

func (p *NetlifyProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config NetlifyProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Netlify API Endpoint",
			"The provider cannot create the Netlify API client as there is an unknown configuration value for the Netlify API endpoint. "+
				"Either use the default value, target apply the source of the value first, set the value statically in the configuration, or use the NETLIFY_API_ENDPOINT environment variable.",
		)
	}
	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown Netlify API Authentication Token",
			"The provider cannot create the Netlify API client as there is an unknown configuration value for the Netlify API authentication token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the NETLIFY_API_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, noEndpointEnvVar := os.LookupEnv("NETLIFY_API_ENDPOINT")
	if !noEndpointEnvVar {
		endpoint = "https://api.netlify.com"
	}
	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}
	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Missing Netlify API Endpoint",
			"Please use the default value, set the NETLIFY_API_ENDPOINT environment variable, or provide a value for the 'endpoint' configuration attribute.",
		)
	}

	url, err := url.Parse(endpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid Netlify API Endpoint",
			"Please provide a valid URL for the 'endpoint' configuration attribute.",
		)
	}
	if url.Scheme == "" {
		url.Scheme = "https"
	}

	token := os.Getenv("NETLIFY_API_TOKEN")
	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}
	if token == "" {
		resp.Diagnostics.AddError(
			"Missing Netlify API Authentication Token",
			"Please set the NETLIFY_API_TOKEN environment variable, or provide a value for the 'token' configuration attribute.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var data NetlifyProviderData

	apiConfig := netlifyapi.NewConfiguration()
	apiConfig.Servers[0].URL = url.String()
	apiConfig.UserAgent = "Terraform Provider/" + p.version
	apiConfig.AddDefaultHeader("Authorization", "Bearer "+token)
	// TODO: Add debug/trace logging to the API client, perhaps by overriding the behavior of apiConfig.Debug
	p.client = netlifyapi.NewAPIClient(apiConfig)
	data.client = p.client

	var account *netlifyapi.Account

	if !config.DefaultTeamID.IsNull() {
		var err error
		account, _, err = data.client.AccountsAPI.GetAccount(ctx, config.DefaultTeamID.ValueString()).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading Netlify team",
				fmt.Sprintf("Could not read Netlify team ID %q: %q", config.DefaultTeamID.ValueString(), err.Error()),
			)
			return
		}
	} else if !config.DefaultTeamSlug.IsNull() {
		accounts, _, err := data.client.AccountsAPI.ListAccountsForUser(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify team", fmt.Sprintf("Could not list Netlify teams: %q", err.Error()))
			return
		}
		slugString := config.DefaultTeamSlug.ValueString()
		for _, a := range accounts {
			if a.Slug == slugString {
				acc := a
				account = &acc
				break
			}
		}
		if account == nil {
			resp.Diagnostics.AddError("Error reading Netlify team", fmt.Sprintf("Could not find Netlify team with slug %q", slugString))
			return
		}
	}

	if account != nil {
		data.defaultTeamId = &account.Id
		data.defaultTeamSlug = &account.Slug
	}

	resp.DataSourceData = data
	resp.ResourceData = data
}

func (p *NetlifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDeployKeyResource,
		NewDnsRecordResource,
		NewDnsZoneResource,
		NewEnvironmentVariableResource,
		NewFirewallTrafficRulesResource(true),
		NewFirewallTrafficRulesResource(false),
		NewLogDrainResource,
		NewSiteBuildSettingsResource,
		NewSiteCollaborationSettingsResource,
		NewSiteDomainSettingsResource,
		NewSiteMetricsSettingsResource,
	}
}

func (p *NetlifyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewDnsZoneDataSource,
		NewSiteDataSource,
		NewSitesDataSource,
		NewTeamDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NetlifyProvider{
			version: version,
		}
	}
}
