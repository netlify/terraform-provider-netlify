package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSite(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site" "example" {
  name = "test-netlify-terraform-provider"
  team_slug = "jenslanghammer"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_site.example", "id"),
				// resource.TestCheckResourceAttr("netlify_site.example", "custom_domain", "test-netlify-terraform-provider.netlify.app"),
			),
		},
	}, testAccSiteDestroy)
}

func testAccSiteDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_site.example"]; ok {
			key, _, err := testAccProvider.client.SitesAPI.GetSite(context.Background(), v.Primary.Attributes["id"]).Execute()
			if err != nil {
				//lint:ignore nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("Site still exists: %s", key.Id)
		}
	}
	return fmt.Errorf("not found in testAccSiteDestroy check destroy")
}
