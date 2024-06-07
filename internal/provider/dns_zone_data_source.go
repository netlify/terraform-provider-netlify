package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ datasource.DataSource              = &dnsZoneDataSource{}
	_ datasource.DataSourceWithConfigure = &dnsZoneDataSource{}
)

func NewDnsZoneDataSource() datasource.DataSource {
	return &dnsZoneDataSource{}
}

type dnsZoneDataSource struct {
	data NetlifyProviderData
}

type dnsZoneDataSourceModel struct {
	ID          types.String        `tfsdk:"id"`
	Name        types.String        `tfsdk:"name"`
	AccountID   types.String        `tfsdk:"account_id"`
	AccountSlug types.String        `tfsdk:"account_slug"`
	DnsServers  types.List          `tfsdk:"dns_servers"`
	Domain      *netlifyDomainModel `tfsdk:"domain"`
	Records     []dnsRecordModel    `tfsdk:"records"`
}

type dnsRecordModel struct {
	ID       types.String `tfsdk:"id"`
	Type     types.String `tfsdk:"type"`
	Hostname types.String `tfsdk:"hostname"`
	Value    types.String `tfsdk:"value"`
	TTL      types.Int64  `tfsdk:"ttl"`
	Priority types.Int64  `tfsdk:"priority"`
	Flag     types.Int64  `tfsdk:"flag"`
	Tag      types.String `tfsdk:"tag"`
}

func (d *dnsZoneDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *dnsZoneDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone"
}

func (d *dnsZoneDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.AtLeastOneOf(path.MatchRoot("name")),
				},
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Computed: true,
			},
			"account_slug": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"registered_at": schema.StringAttribute{
						Computed: true,
					},
					"expires_at": schema.StringAttribute{
						Computed: true,
					},
					"renewal_price": schema.StringAttribute{
						Computed: true,
					},
					"auto_renew": schema.BoolAttribute{
						Computed: true,
					},
					"auto_renew_at": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"dns_servers": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"records": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed: true,
						},
						"hostname": schema.StringAttribute{
							Computed: true,
						},
						"value": schema.StringAttribute{
							Computed: true,
						},
						"ttl": schema.Int64Attribute{
							Computed: true,
						},
						"priority": schema.Int64Attribute{
							Computed: true,
						},
						"flag": schema.Int64Attribute{
							Computed: true,
						},
						"tag": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *dnsZoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config dnsZoneDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var zone *netlifyapi.DnsZone
	if !config.ID.IsUnknown() && !config.ID.IsNull() {
		var err error
		zone, _, err = d.data.client.DNSZonesAPI.GetDnsZone(ctx, config.ID.ValueString()).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify DNS zone", fmt.Sprintf("Could not read Netlify DNS zone ID %q: %q", config.ID.ValueString(), err.Error()))
			return
		}
	} else {
		zones, _, err := d.data.client.DNSZonesAPI.GetDnsZones(ctx).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Error reading Netlify DNS zone", fmt.Sprintf("Could not list Netlify DNS zones: %q", err.Error()))
			return
		}
		nameString := config.Name.ValueString()
		for _, z := range zones {
			if z.Name == nameString {
				zon := z
				zone = &zon
				break
			}
		}
		if zone == nil {
			resp.Diagnostics.AddError("Error reading Netlify DNS zone", fmt.Sprintf("Could not find Netlify DNS zone with name %q", nameString))
			return
		}
	}

	records, _, err := d.data.client.DNSZonesAPI.GetDnsRecords(ctx, zone.Id).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Netlify DNS records",
			fmt.Sprintf("Could not read Netlify DNS records for zone %q: %q", zone.Id, err.Error()),
		)
		return
	}

	config.ID = types.StringValue(zone.Id)
	config.Name = types.StringValue(zone.Name)
	config.AccountID = types.StringValue(zone.AccountId)
	config.AccountSlug = types.StringValue(zone.AccountSlug)
	dnsServers := make([]types.String, len(zone.DnsServers))
	for i, dnsServer := range zone.DnsServers {
		dnsServers[i] = types.StringValue(dnsServer)
	}
	var diags diag.Diagnostics
	config.DnsServers, diags = types.ListValueFrom(ctx, types.StringType, dnsServers)
	resp.Diagnostics.Append(diags...)
	if zone.Domain == nil {
		config.Domain = nil
	} else {
		config.Domain = &netlifyDomainModel{
			ID:           types.StringValue(zone.Domain.Id),
			Name:         types.StringValue(zone.Domain.Name),
			RegisteredAt: types.StringValue(zone.Domain.RegisteredAt.Format(time.RFC3339)),
			ExpiresAt:    types.StringValue(zone.Domain.ExpiresAt.Format(time.RFC3339)),
			RenewalPrice: types.StringValue(zone.Domain.RenewalPrice),
			AutoRenew:    types.BoolValue(zone.Domain.AutoRenew),
			AutoRenewAt:  types.StringValue(zone.Domain.AutoRenewAt.Format(time.RFC3339)),
		}
	}

	config.Records = make([]dnsRecordModel, len(records))
	for i, record := range records {
		config.Records[i] = dnsRecordModel{
			ID:       types.StringValue(record.Id),
			Type:     types.StringValue(record.Type),
			Hostname: types.StringValue(record.Hostname),
			Value:    types.StringValue(record.Value),
			TTL:      types.Int64Value(record.Ttl),
		}
		if record.Type == "CAA" {
			config.Records[i].Flag = types.Int64Value(record.Flag)
			config.Records[i].Tag = types.StringValue(record.Tag)
		} else {
			config.Records[i].Flag = types.Int64Null()
			config.Records[i].Tag = types.StringNull()
		}
		if record.Type == "MX" {
			config.Records[i].Priority = types.Int64Value(record.Priority)
		} else {
			config.Records[i].Priority = types.Int64Null()
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
