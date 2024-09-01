package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataDnsZone(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `data "netlify_dns_zone" "example" {
	name = "examplepetstore.com"
}`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.netlify_dns_zone.example", "id", "66afdbce3cf2b4f0fab520d9"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}
