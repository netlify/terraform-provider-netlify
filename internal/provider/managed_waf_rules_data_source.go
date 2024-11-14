package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &managedWafRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &managedWafRulesDataSource{}
)

func NewManagedWafRulesDataSource() datasource.DataSource {
	return &managedWafRulesDataSource{}
}

type managedWafRulesDataSource struct {
	data NetlifyProviderData
}

type managedWafRulesDataSourceModel struct {
	TeamID   types.String                 `tfsdk:"team_id"`
	RuleSets map[string]managedWafRuleSet `tfsdk:"rule_sets"`
}

type managedWafRuleSet struct {
	Definition managedWafRuleSetDefinition `tfsdk:"definition"`
	Rules      []managedWafRule            `tfsdk:"rules"`
}

type managedWafRuleSetDefinition struct {
	ID      types.String `tfsdk:"id"`
	Type    types.String `tfsdk:"type"`
	Version types.String `tfsdk:"version"`
}

type managedWafRule struct {
	ID          types.String `tfsdk:"id"`
	Description types.String `tfsdk:"description"`
	Phase       types.String `tfsdk:"phase"`
	Category    types.String `tfsdk:"category"`
	Severity    types.String `tfsdk:"severity"`
}

func (d *managedWafRulesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *managedWafRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_managed_waf_rules"
}

func (d *managedWafRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Netlify managed WAF rule sets. This should be used when defining a WAF policy (netlify_waf_policy).",
		Attributes: map[string]schema.Attribute{
			"team_id": schema.StringAttribute{
				Required: true,
			},
			"rule_sets": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"definition": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
								"version": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"rules": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Computed: true,
									},
									"description": schema.StringAttribute{
										Computed: true,
									},
									"phase": schema.StringAttribute{
										Computed: true,
									},
									"category": schema.StringAttribute{
										Computed: true,
									},
									"severity": schema.StringAttribute{
										Computed:    true,
										Description: "notice, warning, error, critical",
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

func (d *managedWafRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config managedWafRulesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, _, err := d.data.client.WAFManagedRulesAPI.GetManagedWafRules(ctx, config.TeamID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Netlify WAF managed rules",
			fmt.Sprintf("Could not read Netlify WAF managed rules for team %q: %q", config.TeamID.ValueString(), err.Error()),
		)
		return
	}
	config.RuleSets = make(map[string]managedWafRuleSet)
	for name, ruleSet := range data.RuleSets {
		config.RuleSets[name] = managedWafRuleSet{
			Definition: managedWafRuleSetDefinition{
				ID:      types.StringPointerValue(ruleSet.Definition.Id),
				Type:    types.StringPointerValue(ruleSet.Definition.Type),
				Version: types.StringPointerValue(ruleSet.Definition.Version),
			},
			Rules: make([]managedWafRule, len(ruleSet.Rules)),
		}
		for i, rule := range ruleSet.Rules {
			config.RuleSets[name].Rules[i] = managedWafRule{
				ID:          types.StringPointerValue(rule.Id),
				Description: types.StringPointerValue(rule.Description),
				Phase:       types.StringPointerValue(rule.Phase),
				Category:    types.StringPointerValue(rule.Category),
				Severity:    types.StringPointerValue(rule.Severity),
			}
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
}
