package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccTeamFirewallTrafficRules(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_team_firewall_traffic_rules" "example" {
  team_id = "66ae34e11a567e9092e3850f"
  published = {
    default_action = "allow"
    ip_restrictions = [
      {
        description = "bot network"
        addresses = [
          "192.0.2.0/24",
          "198.51.100.0/24",
        ]
      }
    ]
    geo_exceptions = [
      {
        description = "brazil"
        countries   = ["BR"]
      }
    ]
  }
  unpublished = {
    default_action = "deny"
    ip_exceptions = [
      {
        description = "Allow the VPN IP"
        addresses   = ["203.0.113.65/32"]
      }
    ]
  }
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_team_firewall_traffic_rules.example", "team_id", "66ae34e11a567e9092e3850f"),
			),
		},
		{
			ResourceName:                         "netlify_team_firewall_traffic_rules.example",
			ImportState:                          true,
			ImportStateId:                        "66ae34e11a567e9092e3850f",
			ImportStateVerifyIdentifierAttribute: "team_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, testAccTeamFirewallTrafficRulesDestroy)
}

func testAccTeamFirewallTrafficRulesDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_team_firewall_traffic_rules.example"]; ok {
			key, _, err := testAccProvider.client.AccountsAPI.GetAccountFirewallRuleSet(context.Background(), v.Primary.Attributes["team_id"]).Execute()
			if err != nil {
				//lint:ignore nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("Team firewall traffic rules still exist: %s", key.Id)
		}
	}
	return fmt.Errorf("not found in testAccTeamFirewallTrafficRulesDestroy check destroy")
}

func TestAccSiteFirewallTrafficRules(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_firewall_traffic_rules" "example" {
  site_id = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  published = {
    default_action = "allow"
    ip_restrictions = [
      {
        description = "bot network"
        addresses = [
          "192.0.2.0/24",
          "198.51.100.0/24",
        ]
      }
    ]
    geo_exceptions = [
      {
        description = "brazil"
        countries   = ["BR"]
      }
    ]
  }
  unpublished = {
    default_action = "deny"
    ip_exceptions = [
      {
        description = "Allow the VPN IP"
        addresses   = ["203.0.113.65/32"]
      }
    ]
  }
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_firewall_traffic_rules.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
			),
		},
		{
			ResourceName:                         "netlify_site_firewall_traffic_rules.example",
			ImportState:                          true,
			ImportStateId:                        "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, testAccSiteFirewallTrafficRulesDestroy)
}

func testAccSiteFirewallTrafficRulesDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_site_firewall_traffic_rules.example"]; ok {
			key, _, err := testAccProvider.client.SitesAPI.GetSiteFirewallRuleSet(context.Background(), v.Primary.Attributes["site_id"]).Execute()
			if err != nil {
				//nolint nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("Site firewall traffic rules still exist: %s", key.Id)
		}
	}
	return fmt.Errorf("not found in testAccSiteFirewallTrafficRulesDestroy check destroy")
}
