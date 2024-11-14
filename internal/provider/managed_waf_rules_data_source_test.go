package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataManagedWafRules(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `data "netlify_managed_waf_rules" "example" {
  team_id = "66ae34e11a567e9092e3850f"
}`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "rule_sets.crs-basic.definition.id", "crs-basic"),
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "rule_sets.crs-basic.rules.0.id", "913100"),
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "rule_sets.crs-basic.rules.0.description", "Found User-Agent associated with security scanner"),
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "rule_sets.crs-basic.rules.0.phase", "0"),
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "rule_sets.crs-basic.rules.0.category", "reputation-scanner"),
				resource.TestCheckResourceAttr("data.netlify_managed_waf_rules.example", "rule_sets.crs-basic.rules.0.severity", "critical"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}
