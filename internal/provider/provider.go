package provider

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/go-openapi/runtime"
	openApiClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netlify/terraform-provider-netlify/internal/plumbing"
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
	client   *plumbing.Netlify
	authInfo runtime.ClientAuthInfoWriter
}

func (p *NetlifyProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "netlify"
	resp.Version = p.version
}

func (p *NetlifyProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Defaults to: https://api.netlify.com/api/v1/",
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
		endpoint = "https://api.netlify.com/api/v1/"
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

	data.client = plumbing.New(openApiClient.NewWithClient(
		url.Host, url.Path, []string{url.Scheme},
		&http.Client{
			Transport: &loggingTransport{
				Transport: cleanhttp.DefaultClient().Transport,
			},
		}), strfmt.Default)
	data.authInfo = runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		err := r.SetHeaderParam("User-Agent", "Terraform")
		if err != nil {
			return err
		}
		err = r.SetHeaderParam("Authorization", "Bearer "+token)
		if err != nil {
			return err
		}
		return nil
	})

	resp.DataSourceData = data
	resp.ResourceData = data
}

func (p *NetlifyProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource { return NewEnvironmentVariableResource(false) },
		func() resource.Resource { return NewEnvironmentVariableResource(true) },
	}
}

func (p *NetlifyProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAccountDataSource,
		NewSiteDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &NetlifyProvider{
			version: version,
		}
	}
}

type loggingTransport struct {
	Transport http.RoundTripper
}

func (lt *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()

	requestDump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		tflog.Trace(ctx, "Netlify API Request", map[string]any{"request": string(requestDump)})
	}

	resp, err := lt.Transport.RoundTrip(req)
	if err != nil {
		tflog.Error(ctx, "Netlify API Request failed", map[string]any{"error": err.Error()})
		return nil, err
	}

	responseDump, err := httputil.DumpResponse(resp, true)
	if err == nil {
		tflog.Trace(ctx, "Netlify API Response", map[string]any{"response": string(responseDump)})
	}

	return resp, nil
}
