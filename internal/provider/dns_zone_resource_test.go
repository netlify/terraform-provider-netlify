package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDnsZone(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_dns_zone" "example" {
  name      = "nf-terraform-test-2.com"
  team_slug = "netlify-terraform-test"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_dns_zone.example", "id"),
				resource.TestCheckResourceAttr("netlify_dns_zone.example", "team_id", "66ae34e11a567e9092e3850f"),
			),
		},
		{
			ResourceName:            "netlify_dns_zone.example",
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"last_updated"},
		},

		{
			Config: `resource "netlify_dns_zone" "example" {
	name      = "nf-terraform-test-2-updated.com"
	team_slug = "netlify-terraform-test"
}`,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction("netlify_dns_zone.example", plancheck.ResourceActionReplace),
				},
			},
		},
	}, testAccDnsZoneDestroy)
}

func testAccDnsZoneDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_dns_zone.example"]; ok {
			key, _, err := testAccProvider.client.DNSZonesAPI.GetDnsZone(context.Background(), v.Primary.Attributes["id"]).Execute()
			if err != nil {
				//lint:ignore nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("DNS zone still exists: %s", key.Id)
		}
	}
	return fmt.Errorf("not found in testAccDnsZoneDestroy check destroy")
}
