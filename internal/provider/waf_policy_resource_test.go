package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccWafPolicy(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_waf_policy" "example" {
  team_id     = "66ae34e11a567e9092e3850f"
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = true,
      overall_threshold = 5,
      category_thresholds = {
        "fixation" = 8,
      },
      rule_overrides = {
        "920100" = {
          action = "log_only"
        }
      }
    }
  ]
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_waf_policy.example", "id"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "name", "Terraform Policy"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "description", "This is a test policy through Terraform"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.managed_id", "crs-basic"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.passive_mode", "true"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.overall_threshold", "5"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.category_thresholds.fixation", "8"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.rule_overrides.920100.action", "log_only"),
			),
		},
		{
			ResourceName: "netlify_waf_policy.example",
			ImportState:  true,
			ImportStateIdFunc: func(s *terraform.State) (string, error) {
				for _, m := range s.Modules {
					if v, ok := m.Resources["netlify_waf_policy.example"]; ok {
						return fmt.Sprintf("%s:%s", v.Primary.Attributes["team_id"], v.Primary.Attributes["id"]), nil
					}
				}
				return "", fmt.Errorf("not found in TestAccWafPolicy import test step")
			},
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"last_updated"},
		},
		{
			Config: `resource "netlify_waf_policy" "example" {
  team_id     = "66ae34e11a567e9092e3850f"
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = false,
      overall_threshold = 6,
      category_thresholds = {
        "fixation" = 9,
      },
      rule_overrides = {
        "920100" = {
          action = "log_only"
        }
        "920170" = {
          action = "none"
        }
      }
    }
  ]
}`,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction("netlify_waf_policy.example", plancheck.ResourceActionUpdate),
				},
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_waf_policy.example", "id"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "name", "Terraform Policy"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "description", "This is a test policy through Terraform"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.managed_id", "crs-basic"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.passive_mode", "false"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.overall_threshold", "6"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.category_thresholds.fixation", "9"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.rule_overrides.920100.action", "log_only"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.rule_overrides.920170.action", "none"),
			),
		},
	}, testAccWafPolicyDestroy)
}

func TestAccWafPolicyOnlyRequired(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_waf_policy" "example" {
  team_id     = "66ae34e11a567e9092e3850f"
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = true,
      overall_threshold = 5
    }
  ]
}`,
		},
	}, testAccWafPolicyDestroy)
}

func TestAccWafPolicyOverrideByCategory(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_waf_policy" "example" {
  team_id     = "66ae34e11a567e9092e3850f"
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = true,
      overall_threshold = 5,
      rule_overrides = {
        for rule in local.crs_basic_rce_rules : rule.id => {
          action = "log_only"
        }
      }
    }
  ]
}

locals {
  crs_basic_rce_rules = flatten([
    for rule_set_key, rule_set in data.netlify_managed_waf_rules.example.rule_sets :
    [
      for rule in rule_set.rules :
      rule if rule_set_key == "crs-basic" && rule.category == "rce"
    ]
  ])
}

data "netlify_managed_waf_rules" "example" {
  team_id = "66ae34e11a567e9092e3850f"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_waf_policy.example", "id"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "name", "Terraform Policy"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "description", "This is a test policy through Terraform"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.managed_id", "crs-basic"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.passive_mode", "true"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.overall_threshold", "5"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.rule_overrides.932175.action", "log_only"),
				resource.TestCheckResourceAttr("netlify_waf_policy.example", "rule_sets.0.rule_overrides.932180.action", "log_only"),
				resource.TestCheckNoResourceAttr("netlify_waf_policy.example", "rule_sets.0.rule_overrides.920500"),
			),
		},
	}, testAccWafPolicyDestroy)
}

func testAccWafPolicyDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_waf_policy.example"]; ok {
			policy, _, err := testAccProvider.client.WAFPoliciesAPI.GetWafPolicy(context.Background(), v.Primary.Attributes["team_id"], v.Primary.Attributes["id"]).Execute()
			if err != nil {
				//lint:ignore nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("WAF policy still exists: %s (team ID: %s)", *policy.Id, v.Primary.Attributes["team_id"])
		}
	}
	return fmt.Errorf("not found in testAccWafPolicyDestroy check destroy")
}
