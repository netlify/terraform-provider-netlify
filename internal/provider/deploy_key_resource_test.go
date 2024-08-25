package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDeployKey(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_deploy_key" "example" {}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_deploy_key.example", "id"),
				resource.TestCheckResourceAttrSet("netlify_deploy_key.example", "last_updated"),
				resource.TestCheckResourceAttrSet("netlify_deploy_key.example", "public_key"),
			),
		},
		{
			ResourceName:            "netlify_deploy_key.example",
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"last_updated"},
		},
	}, testAccDeployKeyDestroy)
}

func testAccDeployKeyDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_deploy_key.example"]; ok {
			key, _, err := testAccProvider.client.DeployKeysAPI.GetDeployKey(context.Background(), v.Primary.Attributes["id"]).Execute()
			if err != nil {
				//lint:ignore nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("Deploy key still exists: %s", key.Id)
		}
	}
	return fmt.Errorf("not found in testAccDeployKeyDestroy check destroy")
}
