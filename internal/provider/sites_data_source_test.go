package provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSites(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `data "netlify_sites" "example" {
	team_slug = "netlify-terraform-test"
}`,
			Check: resource.ComposeTestCheckFunc(
				func(s *terraform.State) error {
					for _, m := range s.Modules {
						if v, ok := m.Resources["data.netlify_sites.example"]; ok {
							for k, v := range v.Primary.Attributes {
								if strings.HasPrefix(k, "sites.") && strings.HasSuffix(k, ".name") && v == "tf-test-1" {
									return nil
								}
							}
							return fmt.Errorf("not found in sites list at TestAccDataSites test step")
						}
					}
					return fmt.Errorf("sites list not found in TestAccDataSites test step")
				},
			),
		},
	}, func(s *terraform.State) error { return nil })
}
