package provider

import (
	"context"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var _ provider.Provider = &NetlifyProvider{}

type NetlifyProvider struct {
	version string
}

type NetlifyProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	Token    types.String `tfsdk:"token"`
}

type NetlifyProviderData struct {
	client *netlifyapi.APIClient
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
				MarkdownDescription: "Read: https://docs.netlify.com/api/get-started/",
				Optional:            true,
				Sensitive:           true,
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
	apiConfig.UserAgent = "Terraform Provider"
	apiConfig.AddDefaultHeader("Authorization", "Bearer "+token)
	// TODO: Add debug/trace logging to the API client, perhaps by overriding the behavior of apiConfig.Debug
	data.client = netlifyapi.NewAPIClient(apiConfig)

	resp.DataSourceData = data
	resp.ResourceData = data
}

func (p *NetlifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewDnsRecordResource,
		NewDnsZoneResource,
		NewEnvironmentVariableResource,
		NewFirewallTrafficRulesResource(true),
		NewFirewallTrafficRulesResource(false),
		NewLogDrainResource,
		NewSiteBuildSettingsResource,
		NewSiteCollaborationSettingsResource,
		NewSiteDomainSettingsResource,
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
