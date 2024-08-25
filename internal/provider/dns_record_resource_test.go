package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var zoneId = "66afdbce3cf2b4f0fab520d9"

func TestAccDnsRecordA(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: fmt.Sprintf(`resource "netlify_dns_record" "example" {
	type = "A"
	zone_id = "%s"
	hostname = "testacc.examplepetstore.com"
	value = "10.0.0.0"
}`, zoneId),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_dns_record.example", "type", "A"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "zone_id", zoneId),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "hostname", "testacc.examplepetstore.com"),
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
			Config: fmt.Sprintf(`resource "netlify_dns_record" "example" {
	type = "A"
	zone_id = "%s"
	hostname = "testacc.examplepetstore.com"
	value = "10.0.0.1"
}`, zoneId),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_dns_record.example", "type", "A"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "zone_id", zoneId),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "hostname", "testacc.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_dns_record.example", "value", "10.0.0.1"),
			),
		},
	}, testAccDnsRecordCheckDestroy("testacc.examplepetstore.com"))
}

func testAccDnsRecordCheckDestroy(hostname string) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		records, _, err := testAccProvider.client.DNSZonesAPI.GetDnsRecords(context.Background(), zoneId).Execute()
		if err != nil {
			return err
		}
		for _, record := range records {
			if record.Hostname == hostname {
				return fmt.Errorf("DNS record still exists")
			}
		}
		return nil
	}
}
