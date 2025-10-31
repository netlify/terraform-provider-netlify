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
	name = "nf-terraform-test.com"
}`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.netlify_dns_zone.example", "id", "69052bba28ce689f129b1ac8"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}
