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
	_ datasource.DataSource              = &accountDataSource{}
	_ datasource.DataSourceWithConfigure = &accountDataSource{}
)

func NewAccountDataSource() datasource.DataSource {
	return &accountDataSource{}
}

type accountDataSource struct {
	data NetlifyProviderData
}

type accountDataSourceModel struct {
	ID   types.String `tfsdk:"id"`
	Slug types.String `tfsdk:"slug"`
	Name types.String `tfsdk:"name"`
}

func (d *accountDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *accountDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (d *accountDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"slug": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *accountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config accountDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if (config.ID.IsUnknown() || config.ID.IsNull()) && (config.Slug.IsUnknown() || config.Slug.IsNull()) {
		resp.Diagnostics.AddError("Error reading Netlify account", "Either id or slug must be specified for an account search")
		return
	}

	var account *models.AccountMembership
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		accountOk, err := d.data.client.Operations.GetAccount(
			operations.NewGetAccountParams().WithAccountID(config.ID.ValueString()),
			d.data.authInfo,
		)
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not read Netlify account ID %q: %q",
				config.ID.ValueString(), err.Error()))
			return
		}
		account = accountOk.Payload
	} else {
		accountsOk, err := d.data.client.Operations.ListAccountsForUser(
			operations.NewListAccountsForUserParams(),
			d.data.authInfo,
		)
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not list Netlify accounts: %q", err.Error()))
			return
		}
		slugString := config.Slug.ValueString()
		for _, a := range accountsOk.Payload {
			if a.Slug == slugString {
				account = a
				break
			}
		}
		if account == nil {
			resp.Diagnostics.AddError("Error reading Netlify account", fmt.Sprintf("Could not find Netlify account with slug %q", slugString))
			return
		}
	}

	config.ID = types.StringValue(account.ID)
	config.Slug = types.StringValue(account.Slug)
	config.Name = types.StringValue(account.Name)

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
