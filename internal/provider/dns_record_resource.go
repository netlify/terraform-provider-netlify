package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/models"
	"github.com/netlify/terraform-provider-netlify/internal/plumbing/operations"
)

var (
	_ resource.Resource                = &dnsRecordResource{}
	_ resource.ResourceWithConfigure   = &dnsRecordResource{}
	_ resource.ResourceWithImportState = &dnsRecordResource{}
)

var NewDnsRecordResource = func(recordType string) func() resource.Resource {
	return func() resource.Resource {
		return &dnsRecordResource{
			recordType: recordType,
		}
	}
}

type dnsRecordResource struct {
	data       NetlifyProviderData
	recordType string
}

type dnsRecordResourceModel struct {
	ZoneID      types.String `tfsdk:"zone_id"`
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Hostname    types.String `tfsdk:"hostname"`
	Value       types.String `tfsdk:"value"`
	TTL         types.Int64  `tfsdk:"ttl"`
	Priority    types.Int64  `tfsdk:"priority"`
	Flag        types.Int64  `tfsdk:"flag"`
	Tag         types.String `tfsdk:"tag"`
}

func (r *dnsRecordResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = fmt.Sprintf("%s_dns_%s_record", req.ProviderTypeName, strings.ToLower(r.recordType))
}

func (r *dnsRecordResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *dnsRecordResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"hostname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"value": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(3600),
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"flag": schema.Int64Attribute{
				Computed: true,
			},
			"tag": schema.StringAttribute{
				Computed: true,
			},
			"priority": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
	if r.recordType == "CAA" {
		resp.Schema.Attributes["flag"] = schema.Int64Attribute{
			Required: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.RequiresReplace(),
			},
		}
		resp.Schema.Attributes["tag"] = schema.StringAttribute{
			Required: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplace(),
			},
		}
	}
	if r.recordType == "MX" {
		resp.Schema.Attributes["priority"] = schema.Int64Attribute{
			Required: true,
			PlanModifiers: []planmodifier.Int64{
				int64planmodifier.RequiresReplace(),
			},
		}
	}
}

func (r *dnsRecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan dnsRecordResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dnsRecordCreate := models.DNSRecordCreate{
		Type:     r.recordType,
		Hostname: plan.Hostname.ValueString(),
		Value:    plan.Value.ValueString(),
		TTL:      plan.TTL.ValueInt64(),
	}
	if r.recordType == "CAA" {
		dnsRecordCreate.Flag = plan.Flag.ValueInt64()
		dnsRecordCreate.Tag = plan.Tag.ValueString()
	}
	if r.recordType == "MX" {
		dnsRecordCreate.Priority = plan.Priority.ValueInt64()
	}
	dnsRecord, err := r.data.client.Operations.CreateDNSRecord(
		operations.
			NewCreateDNSRecordParams().
			WithZoneID(plan.ZoneID.ValueString()).
			WithDNSRecord(&dnsRecordCreate),
		r.data.authInfo,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Netlify DNS record",
			fmt.Sprintf(
				"Could not create Netlify DNS record %q (zone ID: %q): %q",
				plan.Hostname.ValueString(),
				plan.ZoneID.ValueString(),
				err.Error(),
			),
		)
		return
	}
	plan.ID = types.StringValue(dnsRecord.Payload.ID)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	if r.recordType != "CAA" {
		plan.Flag = types.Int64Null()
		plan.Tag = types.StringNull()
	}
	if r.recordType != "MX" {
		plan.Priority = types.Int64Null()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *dnsRecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state dnsRecordResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	dnsRecord, err := r.data.client.Operations.GetIndividualDNSRecord(
		operations.NewGetIndividualDNSRecordParams().
			WithZoneID(state.ZoneID.ValueString()).
			WithDNSRecordID(state.ID.ValueString()),
		r.data.authInfo,
	)
	if err != nil || dnsRecord.Payload.Type != r.recordType {
		resp.Diagnostics.AddError(
			"Error reading Netlify DNS record",
			fmt.Sprintf(
				"Could not read Netlify DNS record %q (zone ID: %q): %q",
				state.Hostname.ValueString(),
				state.ZoneID.ValueString(),
				err.Error(),
			),
		)
		return
	}
	state.Hostname = types.StringValue(dnsRecord.Payload.Hostname)
	state.Value = types.StringValue(dnsRecord.Payload.Value)
	state.TTL = types.Int64Value(dnsRecord.Payload.TTL)
	if r.recordType == "CAA" {
		state.Flag = types.Int64Value(dnsRecord.Payload.Flag)
		state.Tag = types.StringValue(dnsRecord.Payload.Tag)
	} else {
		state.Flag = types.Int64Null()
		state.Tag = types.StringNull()
	}
	if r.recordType == "MX" {
		state.Priority = types.Int64Value(dnsRecord.Payload.Priority)
	} else {
		state.Priority = types.Int64Null()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *dnsRecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update not supported for Netlify DNS records",
		"Update is not supported for Netlify DNS records at this time.",
	)
}

func (r *dnsRecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state dnsRecordResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.client.Operations.DeleteDNSRecord(
		operations.NewDeleteDNSRecordParams().
			WithZoneID(state.ZoneID.ValueString()).
			WithDNSRecordID(state.ID.ValueString()),
		r.data.authInfo,
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Netlify DNS record",
			fmt.Sprintf(
				"Could not delete Netlify DNS record %q (zone ID: %q): %q",
				state.Hostname.ValueString(),
				state.ZoneID.ValueString(),
				err.Error(),
			),
		)
		return
	}
}

func (r *dnsRecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ":")

	errorMessage := fmt.Sprintf("Expected import identifier in the formats: zone_id,record_id. Got: %q", req.ID)

	if len(idParts) == 2 {
		if idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("zone_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	} else {
		resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
		return
	}
}
