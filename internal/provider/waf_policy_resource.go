package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netlify/terraform-provider-netlify/internal/netlifyapi"
)

var (
	_ resource.Resource                = &wafPolicyResource{}
	_ resource.ResourceWithConfigure   = &wafPolicyResource{}
	_ resource.ResourceWithImportState = &wafPolicyResource{}
)

func NewWafPolicyResource() resource.Resource {
	return &wafPolicyResource{}
}

type wafPolicyResource struct {
	data NetlifyProviderData
}

type wafPolicyResourceModel struct {
	ID          types.String      `tfsdk:"id"`
	TeamID      types.String      `tfsdk:"team_id"`
	LastUpdated types.String      `tfsdk:"last_updated"`
	Name        types.String      `tfsdk:"name"`
	Description types.String      `tfsdk:"description"`
	RuleSets    []wafRuleSetModel `tfsdk:"rule_sets"`
}

type wafRuleSetModel struct {
	ManagedID          types.String                    `tfsdk:"managed_id"`
	ExcludedPatterns   []types.String                  `tfsdk:"excluded_patterns"`
	PassiveMode        types.Bool                      `tfsdk:"passive_mode"`
	OverallThreshold   types.Int64                     `tfsdk:"overall_threshold"`
	CategoryThresholds map[string]types.Int64          `tfsdk:"category_thresholds"`
	RuleOverrides      map[string]wafRuleOverrideModel `tfsdk:"rule_overrides"`
}

type wafRuleOverrideModel struct {
	Action types.String `tfsdk:"action"`
}

func (r *wafPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_waf_policy"
}

func (r *wafPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// TODO: verify the required and optional properties.
func (r *wafPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Netlify Web Application Firewall (WAF) policy",
		MarkdownDescription: "Netlify Web Application Firewall (WAF) policy. [Read more](https://docs.netlify.com/security/secure-access-to-sites/web-application-firewall/)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"team_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Required: true,
			},
			"rule_sets": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"managed_id": schema.StringAttribute{
							Required:    true,
							Description: "The managed ID of the rule set. Currently, only crs-basic is supported.",
							Validators: []validator.String{
								stringvalidator.OneOf("crs-basic"),
							},
						},
						"excluded_patterns": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"passive_mode": schema.BoolAttribute{
							Required: true,
						},
						"overall_threshold": schema.Int64Attribute{
							Required:    true,
							Description: "Recommended default value is 5",
						},
						"category_thresholds": schema.MapAttribute{
							Optional:    true,
							Computed:    true,
							ElementType: types.Int64Type,
							Description: "Thresholds for each category, e.g. fixation, injection-generic, injection-java, injection-php, lfi, protocol, rce, reputation-scanner, rfi, sqli, ssrf, xss",
							Default:     mapdefault.StaticValue(types.MapValueMust(types.Int64Type, map[string]attr.Value{})),
							Validators: []validator.Map{
								mapvalidator.KeysAre(stringvalidator.OneOf(
									"fixation",
									"injection-generic",
									"injection-java",
									"injection-php",
									"lfi",
									"protocol",
									"rce",
									"reputation-scanner",
									"rfi",
									"sqli",
									"ssrf",
									"xss",
								)),
							},
						},
						"rule_overrides": schema.MapNestedAttribute{
							Optional: true,
							Computed: true,
							Default:  mapdefault.StaticValue(types.MapValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{"action": types.StringType}}, map[string]attr.Value{})),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Required:    true,
										Description: "log_only or none",
										Validators: []validator.String{
											stringvalidator.OneOf("log_only", "none"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *wafPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan wafPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policyId, _, err := r.data.client.WAFPoliciesAPI.CreateWafPolicy(ctx, plan.TeamID.ValueString()).WafPolicy(r.serializeWafPolicy(&plan)).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating WAF policy",
			fmt.Sprintf("Could not create WAF policy (team ID %q): %q", plan.TeamID.ValueString(), err.Error()),
		)
		return
	}
	plan.ID = types.StringValue(policyId.PolicyId)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wafPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state wafPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, _, err := r.data.client.WAFPoliciesAPI.GetWafPolicy(ctx, state.TeamID.ValueString(), state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading WAF policy",
			fmt.Sprintf(
				"Could not read WAF policy %q (team ID %q): %q",
				state.ID.ValueString(),
				state.TeamID.ValueString(),
				err.Error(),
			),
		)
		return
	}

	state.Name = types.StringValue(policy.Name)
	state.Description = types.StringValue(policy.Description)
	state.RuleSets = make([]wafRuleSetModel, len(policy.RuleSets))
	for i, ruleSet := range policy.RuleSets {
		state.RuleSets[i].ManagedID = types.StringPointerValue(ruleSet.ManagedId)
		state.RuleSets[i].PassiveMode = types.BoolValue(ruleSet.PassiveMode != nil && *ruleSet.PassiveMode)
		state.RuleSets[i].OverallThreshold = types.Int64Value(ruleSet.OverallThreshold)
		state.RuleSets[i].CategoryThresholds = make(map[string]types.Int64)
		for k, v := range ruleSet.CategoryThresholds {
			state.RuleSets[i].CategoryThresholds[k] = types.Int64Value(v)
		}
		state.RuleSets[i].RuleOverrides = make(map[string]wafRuleOverrideModel)
		for k, v := range ruleSet.RuleOverrides {
			state.RuleSets[i].RuleOverrides[k] = wafRuleOverrideModel{
				Action: types.StringValue(v.Action),
			}
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wafPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan wafPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.client.WAFPoliciesAPI.UpdateWafPolicy(ctx, plan.TeamID.ValueString(), plan.ID.ValueString()).WafPolicy(r.serializeWafPolicy(&plan)).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating WAF policy",
			fmt.Sprintf("Could not update WAF policy %q (team ID %q): %q", plan.ID.ValueString(), plan.TeamID.ValueString(), err.Error()),
		)
		return
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *wafPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state wafPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.client.WAFPoliciesAPI.DeleteWafPolicy(ctx, state.TeamID.ValueString(), state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting WAF policy",
			fmt.Sprintf("Could not delete WAF policy %q (team ID %q): %q", state.ID.ValueString(), state.TeamID.ValueString(), err.Error()),
		)
		return
	}
}

func (r *wafPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ":")

	errorMessage := fmt.Sprintf("Expected import identifier in the formats: team_id:waf_policy_id. Got: %q", req.ID)

	if len(idParts) == 2 {
		if idParts[0] == "" || idParts[1] == "" {
			resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
			return
		}
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("team_id"), idParts[0])...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
	} else {
		resp.Diagnostics.AddError("Unexpected Import Identifier", errorMessage)
		return
	}
}

func (r *wafPolicyResource) serializeWafPolicy(plan *wafPolicyResourceModel) netlifyapi.WafPolicy {
	policy := netlifyapi.WafPolicy{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueString(),
		RuleSets:    make([]netlifyapi.WafPolicyRuleSetsInner, len(plan.RuleSets)),
	}
	for i, ruleSet := range plan.RuleSets {
		policy.RuleSets[i].ManagedId = ruleSet.ManagedID.ValueStringPointer()
		policy.RuleSets[i].PassiveMode = ruleSet.PassiveMode.ValueBoolPointer()
		policy.RuleSets[i].OverallThreshold = ruleSet.OverallThreshold.ValueInt64()
		policy.RuleSets[i].CategoryThresholds = make(map[string]int64)
		for k, v := range ruleSet.CategoryThresholds {
			policy.RuleSets[i].CategoryThresholds[k] = v.ValueInt64()
		}
		policy.RuleSets[i].RuleOverrides = make(map[string]netlifyapi.WafPolicyRuleOverride)
		for k, v := range ruleSet.RuleOverrides {
			policy.RuleSets[i].RuleOverrides[k] = netlifyapi.WafPolicyRuleOverride{
				Action: v.Action.ValueString(),
			}
		}
	}
	return policy
}
