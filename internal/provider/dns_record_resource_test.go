package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDnsRecordA(t *testing.T) {
	var zoneId = "69052bba28ce689f129b1ac8"
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_dns_record" "example" {
	type = "A"
	zone_id = "69052bba28ce689f129b1ac8"
	hostname = "testacc.nf-terraform-test.com"
	value = "10.0.0.0"
}`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_dns_record.example", "type", "A"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "zone_id", zoneId),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "hostname", "testacc.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "value", "10.0.0.0"),
			),
		},
		{
			ResourceName: "netlify_dns_record.example",
			ImportState:  true,
			ImportStateIdFunc: func(s *terraform.State) (string, error) {
				for _, m := range s.Modules {
					if v, ok := m.Resources["netlify_dns_record.example"]; ok {
						return fmt.Sprintf("%s:%s", v.Primary.Attributes["zone_id"], v.Primary.Attributes["id"]), nil
					}
				}
				return "", fmt.Errorf("not found in TestAccDnsRecordA import test step")
			},
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"last_updated"},
		},
		{
			Config: `resource "netlify_dns_record" "example" {
	type = "A"
	zone_id = "69052bba28ce689f129b1ac8"
	hostname = "testacc.nf-terraform-test.com"
	value = "10.0.0.1"
}`,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction("netlify_dns_record.example", plancheck.ResourceActionReplace),
				},
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_dns_record.example", "type", "A"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "zone_id", zoneId),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "hostname", "testacc.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "value", "10.0.0.1"),
			),
		},
	}, testAccDnsRecordCheckDestroy)
}

func testAccDnsRecordCheckDestroy(s *terraform.State) error {
	records, _, err := testAccProvider.client.DNSZonesAPI.GetDnsRecords(context.Background(), "69052bba28ce689f129b1ac8").Execute()
	if err != nil {
		return err
	}
	for _, record := range records {
		if record.Hostname == "testacc.nf-terraform-test.com" {
			return fmt.Errorf("DNS record still exists")
		}
	}
	return nil
}
