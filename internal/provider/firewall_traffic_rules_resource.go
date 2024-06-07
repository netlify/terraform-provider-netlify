package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
	"github.com/netlify/terraform-provider-netlify/internal/provider/netlify_planmodifiers"
	"github.com/netlify/terraform-provider-netlify/internal/provider/netlify_validators"
)

var (
	_ resource.Resource                = &firewallTrafficRulesResource{}
	_ resource.ResourceWithConfigure   = &firewallTrafficRulesResource{}
	_ resource.ResourceWithImportState = &firewallTrafficRulesResource{}
)

func NewFirewallTrafficRulesResource(accountLevel bool) func() resource.Resource {
	return func() resource.Resource {
		return &firewallTrafficRulesResource{
			accountLevel: accountLevel,
		}
	}
}

type firewallTrafficRulesResource struct {
	data         NetlifyProviderData
	accountLevel bool
}

type firewallTrafficRulesResourceModel struct {
	SiteID      types.String            `tfsdk:"site_id"`
	AccountID   types.String            `tfsdk:"account_id"`
	LastUpdated types.String            `tfsdk:"last_updated"`
	Published   *firewallTrafficRuleSet `tfsdk:"published"`
	Unpublished *firewallTrafficRuleSet `tfsdk:"unpublished"`
}

type firewallTrafficRuleSet struct {
	DefaultAction   types.String             `tfsdk:"default_action"`
	IPRestrictions  []ipFirewallTrafficRule  `tfsdk:"ip_restrictions"`
	GeoRestrictions []geoFirewallTrafficRule `tfsdk:"geo_restrictions"`
	IPExceptions    []ipFirewallTrafficRule  `tfsdk:"ip_exceptions"`
	GeoExceptions   []geoFirewallTrafficRule `tfsdk:"geo_exceptions"`
}

type ipFirewallTrafficRule struct {
	Addresses   []types.String `tfsdk:"addresses"`
	Description types.String   `tfsdk:"description"`
}

type geoFirewallTrafficRule struct {
	Countries   []types.String `tfsdk:"countries"`
	Subregions  []types.String `tfsdk:"subregions"`
	Description types.String   `tfsdk:"description"`
}

func (r *firewallTrafficRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	if r.accountLevel {
		resp.TypeName = req.ProviderTypeName + "_account_firewall_traffic_rules"
	} else {
		resp.TypeName = req.ProviderTypeName + "_site_firewall_traffic_rules"
	}
}

func (r *firewallTrafficRulesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *firewallTrafficRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	ipFirewallTrafficRuleSchema := schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"addresses": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"description": schema.StringAttribute{
				Required: true,
			},
		},
	}
	geoFirewallTrafficRuleSchema := schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"countries": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"subregions": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"description": schema.StringAttribute{
				Required: true,
			},
		},
	}

	firewallTrafficRuleSetSchema := schema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]schema.Attribute{
			"default_action": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("allow", "deny"),
					netlify_validators.ForbiddenIfEquals(
						"deny",
						path.MatchRelative().AtParent().AtName("ip_restrictions"),
						path.MatchRelative().AtParent().AtName("geo_restrictions"),
					),
				},
			},
			"ip_restrictions": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: ipFirewallTrafficRuleSchema,
			},
			"geo_restrictions": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: geoFirewallTrafficRuleSchema,
			},
			"ip_exceptions": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: ipFirewallTrafficRuleSchema,
			},
			"geo_exceptions": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: geoFirewallTrafficRuleSchema,
			},
		},
	}

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"site_id": schema.StringAttribute{
				Required: !r.accountLevel,
				Computed: r.accountLevel,
				PlanModifiers: []planmodifier.String{
					netlify_planmodifiers.UseNullForUnknown(),
				},
			},
			"account_id": schema.StringAttribute{
				Required: r.accountLevel,
				Computed: !r.accountLevel,
				PlanModifiers: []planmodifier.String{
					netlify_planmodifiers.UseNullForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"published":   firewallTrafficRuleSetSchema,
			"unpublished": firewallTrafficRuleSetSchema,
		},
	}
}

func (r *firewallTrafficRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan firewallTrafficRulesResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.write(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *firewallTrafficRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state firewallTrafficRulesResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var config *netlifyapi.SiteFirewallConfig

	if r.accountLevel {
		var err error
		config, _, err = r.data.client.AccountsAPI.
			GetAccountFirewallRuleSet(ctx, state.AccountID.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading account firewall rule set",
				fmt.Sprintf(
					"Could not read account firewall rule set %q: %q",
					state.AccountID.ValueString(),
					err.Error(),
				),
			)
			return
		}
	} else {
		var err error
		config, _, err = r.data.client.SitesAPI.
			GetSiteFirewallRuleSet(ctx, state.SiteID.ValueString()).
			Id(state.SiteID.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading site firewall rule set",
				fmt.Sprintf(
					"Could not read site firewall rule set %q: %q",
					state.SiteID.ValueString(),
					err.Error(),
				),
			)
			return
		}
	}

	published := r.deserializeRuleSet(config.Published)
	state.Published = &published
	unpublished := r.deserializeRuleSet(config.Unpublished)
	state.Unpublished = &unpublished

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *firewallTrafficRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan firewallTrafficRulesResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.write(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *firewallTrafficRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state firewallTrafficRulesResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.accountLevel {
		_, err := r.data.client.AccountsAPI.
			DeleteAccountFirewallRuleSet(ctx, state.AccountID.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting account firewall rule set",
				fmt.Sprintf(
					"Could not delete account firewall rule set %q: %q",
					state.AccountID.ValueString(),
					err.Error(),
				),
			)
			return
		}
	} else {
		_, err := r.data.client.SitesAPI.
			DeleteSiteFirewallRuleSet(ctx, state.SiteID.ValueString()).
			Id(state.SiteID.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error deleting site firewall rule set",
				fmt.Sprintf(
					"Could not delete site firewall rule set %q: %q",
					state.SiteID.ValueString(),
					err.Error(),
				),
			)
			return
		}
	}
}

func (r *firewallTrafficRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if r.accountLevel {
		resource.ImportStatePassthroughID(ctx, path.Root("account_id"), req, resp)
	} else {
		resource.ImportStatePassthroughID(ctx, path.Root("site_id"), req, resp)
	}
}

func (r *firewallTrafficRulesResource) write(ctx context.Context, plan *firewallTrafficRulesResourceModel, diagnostics *diag.Diagnostics) {
	var createSiteFirewallConfig netlifyapi.CreateSiteFirewallConfig

	published := r.serializeRuleSet(plan.Published)
	createSiteFirewallConfig.Published = &published
	unpublished := r.serializeRuleSet(plan.Unpublished)
	createSiteFirewallConfig.Unpublished = &unpublished

	if r.accountLevel {
		_, err := r.data.client.AccountsAPI.
			UpdateAccountFirewallRuleSet(ctx, plan.AccountID.ValueString()).
			CreateSiteFirewallConfig(createSiteFirewallConfig).
			Execute()
		if err != nil {
			diagnostics.AddError(
				"Error updating account firewall rule set",
				fmt.Sprintf(
					"Could not update account firewall rule set %q: %q",
					plan.AccountID.ValueString(),
					err.Error(),
				),
			)
			return
		}
		plan.SiteID = types.StringNull()
	} else {
		_, err := r.data.client.SitesAPI.
			UpdateSiteFirewallRuleSet(ctx, plan.SiteID.ValueString()).
			Id(plan.SiteID.ValueString()).
			CreateSiteFirewallConfig(createSiteFirewallConfig).
			Execute()
		if err != nil {
			diagnostics.AddError(
				"Error updating site firewall rule set",
				fmt.Sprintf(
					"Could not update site firewall rule set %q: %q",
					plan.SiteID.ValueString(),
					err.Error(),
				),
			)
			return
		}
		plan.AccountID = types.StringNull()
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
}

func (r *firewallTrafficRulesResource) serializeRuleSet(ruleSet *firewallTrafficRuleSet) netlifyapi.FirewallRuleSet {
	if ruleSet == nil {
		return netlifyapi.FirewallRuleSet{
			Default: "allow",
			Rules:   make([]netlifyapi.FirewallRule, 0),
		}
	}
	rules := make([]netlifyapi.FirewallRule, 0, len(ruleSet.IPRestrictions)+len(ruleSet.GeoRestrictions)+len(ruleSet.IPExceptions)+len(ruleSet.GeoExceptions))
	for _, rule := range ruleSet.IPRestrictions {
		addresses := make([]string, len(rule.Addresses))
		for i, address := range rule.Addresses {
			addresses[i] = address.ValueString()
		}
		rules = append(rules, netlifyapi.FirewallRule{
			Type:        "ip",
			Action:      "deny",
			Description: rule.Description.ValueStringPointer(),
			Data: map[string][]string{
				"addresses": addresses,
			},
		})
	}
	for _, rule := range ruleSet.GeoRestrictions {
		countries := make([]string, len(rule.Countries))
		for i, country := range rule.Countries {
			countries[i] = country.ValueString()
		}
		subregions := make([]string, len(rule.Subregions))
		for i, subregion := range rule.Subregions {
			subregions[i] = subregion.ValueString()
		}
		s := netlifyapi.FirewallRule{
			Type:        "geo",
			Action:      "deny",
			Description: rule.Description.ValueStringPointer(),
			Data: map[string][]string{
				"countries": countries,
			},
		}
		if len(subregions) > 0 {
			s.Data["subregions"] = subregions
		}
		rules = append(rules, s)
	}
	for _, rule := range ruleSet.IPExceptions {
		addresses := make([]string, len(rule.Addresses))
		for i, address := range rule.Addresses {
			addresses[i] = address.ValueString()
		}
		rules = append(rules, netlifyapi.FirewallRule{
			Type:        "ip",
			Action:      "allow",
			Description: rule.Description.ValueStringPointer(),
			Data: map[string][]string{
				"addresses": addresses,
			},
		})
	}
	for _, rule := range ruleSet.GeoExceptions {
		countries := make([]string, len(rule.Countries))
		for i, country := range rule.Countries {
			countries[i] = country.ValueString()
		}
		subregions := make([]string, len(rule.Subregions))
		for i, subregion := range rule.Subregions {
			subregions[i] = subregion.ValueString()
		}
		s := netlifyapi.FirewallRule{
			Type:        "geo",
			Action:      "allow",
			Description: rule.Description.ValueStringPointer(),
			Data: map[string][]string{
				"countries": countries,
			},
		}
		if len(subregions) > 0 {
			s.Data["subregions"] = subregions
		}
		rules = append(rules, s)
	}
	return netlifyapi.FirewallRuleSet{
		Default: ruleSet.DefaultAction.ValueString(),
		Rules:   rules,
	}
}

func (r *firewallTrafficRulesResource) deserializeRuleSet(ruleSet netlifyapi.FirewallRuleSet) firewallTrafficRuleSet {
	rs := firewallTrafficRuleSet{
		DefaultAction: types.StringValue(ruleSet.Default),
	}
	for _, rule := range ruleSet.Rules {
		if rule.Disabled != nil && *rule.Disabled {
			continue
		}
		switch rule.Type {
		case "ip":
			addresses := make([]types.String, len(rule.Data["addresses"]))
			for i, address := range rule.Data["addresses"] {
				addresses[i] = types.StringValue(address)
			}
			if rule.Action == "allow" {
				rs.IPExceptions = append(rs.IPExceptions, ipFirewallTrafficRule{
					Addresses:   addresses,
					Description: types.StringPointerValue(rule.Description),
				})
			} else {
				rs.IPRestrictions = append(rs.IPRestrictions, ipFirewallTrafficRule{
					Addresses:   addresses,
					Description: types.StringPointerValue(rule.Description),
				})
			}
		case "geo":
			countries := make([]types.String, len(rule.Data["countries"]))
			for i, country := range rule.Data["countries"] {
				countries[i] = types.StringValue(country)
			}
			var subregions []types.String
			if subregionsData, ok := rule.Data["subregions"]; ok {
				subregions = make([]types.String, len(subregionsData))
				for i, subregion := range subregionsData {
					subregions[i] = types.StringValue(subregion)
				}
			}
			if rule.Action == "allow" {
				rs.GeoExceptions = append(rs.GeoExceptions, geoFirewallTrafficRule{
					Countries:   countries,
					Subregions:  subregions,
					Description: types.StringPointerValue(rule.Description),
				})
			} else {
				rs.GeoRestrictions = append(rs.GeoRestrictions, geoFirewallTrafficRule{
					Countries:   countries,
					Subregions:  subregions,
					Description: types.StringPointerValue(rule.Description),
				})
			}
		}
	}
	return rs
}
