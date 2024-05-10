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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/models"
	"github.com/netlify/terraform-provider-netlify/internal/plumbing/operations"
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
	ID          types.String        `tfsdk:"id"`
	LastUpdated types.String        `tfsdk:"last_updated"`
	Name        types.String        `tfsdk:"name"`
	AccountID   types.String        `tfsdk:"account_id"`
	AccountSlug types.String        `tfsdk:"account_slug"`
	DnsServers  types.List          `tfsdk:"dns_servers"`
	Domain      *NetlifyDomainModel `tfsdk:"domain"`
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
			"account_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"account_slug": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"dns_servers": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			"domain": schema.SingleNestedBlock{
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

	dnsZone, err := r.data.client.Operations.CreateDNSZone(
		operations.
			NewCreateDNSZoneParams().
			WithDNSZoneParams(&models.DNSZoneSetup{
				AccountSlug: plan.AccountSlug.ValueString(),
				Name:        plan.Name.ValueString(),
			}),
		r.data.authInfo,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Netlify DNS zone",
			fmt.Sprintf(
				"Could not create Netlify DNS zone %q (account slug: %q): %q",
				plan.Name.ValueString(),
				plan.AccountSlug.ValueString(),
				err.Error(),
			),
		)
		return
	}
	plan.ID = types.StringValue(dnsZone.Payload.ID)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	plan.AccountID = types.StringValue(dnsZone.Payload.AccountID)
	dnsServers := make([]types.String, len(dnsZone.Payload.DNSServers))
	for i, dnsServer := range dnsZone.Payload.DNSServers {
		dnsServers[i] = types.StringValue(dnsServer)
	}
	var diags diag.Diagnostics
	plan.DnsServers, diags = types.ListValueFrom(ctx, types.StringType, dnsServers)
	resp.Diagnostics.Append(diags...)
	if dnsZone.Payload.Domain == nil {
		plan.Domain = nil
	} else {
		plan.Domain = &NetlifyDomainModel{
			ID:           types.StringValue(dnsZone.Payload.Domain.ID),
			Name:         types.StringValue(dnsZone.Payload.Domain.Name),
			RegisteredAt: types.StringValue(dnsZone.Payload.Domain.RegisteredAt),
			ExpiresAt:    types.StringValue(dnsZone.Payload.Domain.ExpiresAt),
			RenewalPrice: types.StringValue(dnsZone.Payload.Domain.RenewalPrice),
			AutoRenew:    types.BoolValue(dnsZone.Payload.Domain.AutoRenew),
			AutoRenewAt:  types.StringValue(dnsZone.Payload.Domain.AutoRenewAt),
		}
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

	dnsZone, err := r.data.client.Operations.GetDNSZone(
		operations.
			NewGetDNSZoneParams().
			WithZoneID(state.ID.ValueString()),
		r.data.authInfo,
	)
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
	state.Name = types.StringValue(dnsZone.Payload.Name)
	state.AccountID = types.StringValue(dnsZone.Payload.AccountID)
	state.AccountSlug = types.StringValue(dnsZone.Payload.AccountSlug)
	dnsServers := make([]types.String, len(dnsZone.Payload.DNSServers))
	for i, dnsServer := range dnsZone.Payload.DNSServers {
		dnsServers[i] = types.StringValue(dnsServer)
	}
	var diags diag.Diagnostics
	state.DnsServers, diags = types.ListValueFrom(ctx, types.StringType, dnsServers)
	resp.Diagnostics.Append(diags...)
	if dnsZone.Payload.Domain == nil {
		state.Domain = nil
	} else {
		state.Domain = &NetlifyDomainModel{
			ID:           types.StringValue(dnsZone.Payload.Domain.ID),
			Name:         types.StringValue(dnsZone.Payload.Domain.Name),
			RegisteredAt: types.StringValue(dnsZone.Payload.Domain.RegisteredAt),
			ExpiresAt:    types.StringValue(dnsZone.Payload.Domain.ExpiresAt),
			RenewalPrice: types.StringValue(dnsZone.Payload.Domain.RenewalPrice),
			AutoRenew:    types.BoolValue(dnsZone.Payload.Domain.AutoRenew),
			AutoRenewAt:  types.StringValue(dnsZone.Payload.Domain.AutoRenewAt),
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *dnsZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
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

	_, err := r.data.client.Operations.DeleteDNSZone(
		operations.
			NewDeleteDNSZoneParams().
			WithZoneID(state.ID.ValueString()),
		r.data.authInfo,
	)
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