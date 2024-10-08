package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ resource.Resource                = &dnsZoneResource{}
	_ resource.ResourceWithConfigure   = &dnsZoneResource{}
	_ resource.ResourceWithImportState = &dnsZoneResource{}
)

func NewDnsZoneResource() resource.Resource {
	return &dnsZoneResource{}
}

type dnsZoneResource struct {
	data NetlifyProviderData
}

type dnsZoneResourceModel struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Name        types.String `tfsdk:"name"`
	TeamID      types.String `tfsdk:"team_id"`
	TeamSlug    types.String `tfsdk:"team_slug"`
	DnsServers  types.List   `tfsdk:"dns_servers"`
	Domain      types.Object `tfsdk:"domain"`
}

func (r *dnsZoneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dns_zone"
}

func (r *dnsZoneResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dnsZoneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Netlify DNS zone",
		MarkdownDescription: "Netlify DNS zone. [Read more](https://docs.netlify.com/domains-https/netlify-dns/)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"team_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"team_slug": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "Required if a default team was not configured in the provider configuration.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dns_servers": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
			"domain": schema.SingleNestedAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
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
		},
	}
}

func (r *dnsZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan dnsZoneResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	teamSlug := r.data.teamSlugOrDefault(plan.TeamSlug)
	if teamSlug == nil {
		resp.Diagnostics.AddError(
			"Missing team slug",
			"Team slug is required for creating a Netlify DNS zone. Please provide a team slug in the plan or configure a default team in the provider configuration.",
		)
		return
	}

	dnsZone, _, err := r.data.client.DNSZonesAPI.
		CreateDnsZone(ctx).
		DnsZoneCreateParams(netlifyapi.DnsZoneCreateParams{
			AccountSlug: teamSlug,
			Name:        plan.Name.ValueStringPointer(),
		}).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Netlify DNS zone",
			fmt.Sprintf(
				"Could not create Netlify DNS zone %q (team slug: %q): %q",
				plan.Name.ValueString(),
				*teamSlug,
				err.Error(),
			),
		)
		return
	}
	plan.ID = types.StringValue(dnsZone.Id)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))
	plan.TeamID = types.StringValue(dnsZone.AccountId)
	plan.TeamSlug = types.StringValue(dnsZone.AccountSlug)
	dnsServers := make([]types.String, len(dnsZone.DnsServers))
	for i, dnsServer := range dnsZone.DnsServers {
		dnsServers[i] = types.StringValue(dnsServer)
	}
	var diags diag.Diagnostics
	plan.DnsServers, diags = types.ListValueFrom(ctx, types.StringType, dnsServers)
	resp.Diagnostics.Append(diags...)
	if dnsZone.Domain == nil {
		plan.Domain = types.ObjectNull(netlifyDomainModel{}.AttributeTypes())
	} else {
		plan.Domain, diags = types.ObjectValueFrom(ctx, netlifyDomainModel{}.AttributeTypes(), netlifyDomainModel{
			ID:           types.StringValue(dnsZone.Domain.Id),
			Name:         types.StringValue(dnsZone.Domain.Name),
			RegisteredAt: types.StringValue(dnsZone.Domain.RegisteredAt.Format(time.RFC3339)),
			ExpiresAt:    types.StringValue(dnsZone.Domain.ExpiresAt.Format(time.RFC3339)),
			RenewalPrice: types.StringValue(dnsZone.Domain.RenewalPrice),
			AutoRenew:    types.BoolValue(dnsZone.Domain.AutoRenew),
			AutoRenewAt:  types.StringValue(dnsZone.Domain.AutoRenewAt.Format(time.RFC3339)),
		})
		resp.Diagnostics.Append(diags...)
	}

	_, _, err = r.data.client.DNSZonesAPI.EnableDnsZoneIpv6(ctx, plan.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddWarning(
			"Error enabling IPv6 for Netlify DNS zone",
			fmt.Sprintf(
				"Could not enable IPv6 for Netlify DNS zone %q: %q",
				plan.ID.ValueString(),
				err.Error(),
			),
		)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *dnsZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dnsZoneResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dnsZone, _, err := r.data.client.DNSZonesAPI.GetDnsZone(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Netlify DNS zone",
			fmt.Sprintf(
				"Could not read Netlify DNS zone %q: %q",
				state.ID.ValueString(),
				err.Error(),
			),
		)
		return
	}
	state.Name = types.StringValue(dnsZone.Name)
	state.TeamID = types.StringValue(dnsZone.AccountId)
	state.TeamSlug = types.StringValue(dnsZone.AccountSlug)
	dnsServers := make([]types.String, len(dnsZone.DnsServers))
	for i, dnsServer := range dnsZone.DnsServers {
		dnsServers[i] = types.StringValue(dnsServer)
	}
	var diags diag.Diagnostics
	state.DnsServers, diags = types.ListValueFrom(ctx, types.StringType, dnsServers)
	resp.Diagnostics.Append(diags...)
	if dnsZone.Domain == nil {
		state.Domain = types.ObjectNull(netlifyDomainModel{}.AttributeTypes())
	} else {
		state.Domain, diags = types.ObjectValueFrom(ctx, netlifyDomainModel{}.AttributeTypes(), netlifyDomainModel{
			ID:           types.StringValue(dnsZone.Domain.Id),
			Name:         types.StringValue(dnsZone.Domain.Name),
			RegisteredAt: types.StringValue(dnsZone.Domain.RegisteredAt.Format(time.RFC3339)),
			ExpiresAt:    types.StringValue(dnsZone.Domain.ExpiresAt.Format(time.RFC3339)),
			RenewalPrice: types.StringValue(dnsZone.Domain.RenewalPrice),
			AutoRenew:    types.BoolValue(dnsZone.Domain.AutoRenew),
			AutoRenewAt:  types.StringValue(dnsZone.Domain.AutoRenewAt.Format(time.RFC3339)),
		})
		resp.Diagnostics.Append(diags...)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *dnsZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning(
		"Update not supported for Netlify DNS zones",
		"Update is not supported for Netlify DNS zones at this time.",
	)
}

func (r *dnsZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state dnsZoneResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.client.DNSZonesAPI.DeleteDnsZone(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Netlify DNS zone",
			fmt.Sprintf(
				"Could not delete Netlify DNS zone %q: %q",
				state.ID.ValueString(),
				err.Error(),
			),
		)
		return
	}
}

func (r *dnsZoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
