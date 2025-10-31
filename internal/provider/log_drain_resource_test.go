package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccLogDrain(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_log_drain" "example" {
  site_id     = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  destination = "http"
  log_types   = ["deploys"]
  format      = "ndjson"
  exclude_pii = false
  service_config = {
    url = "https://tools-httpstatus.pickup-services.com/200"
  }
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttrSet("netlify_log_drain.example", "id"),
				resource.TestCheckResourceAttr("netlify_log_drain.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_log_drain.example", "destination", "http"),
				resource.TestCheckResourceAttr("netlify_log_drain.example", "exclude_pii", "false"),
			),
		},
		{
			ResourceName: "netlify_log_drain.example",
			ImportState:  true,
			ImportStateIdFunc: func(s *terraform.State) (string, error) {
				for _, m := range s.Modules {
					if v, ok := m.Resources["netlify_log_drain.example"]; ok {
						return fmt.Sprintf("%s:%s", v.Primary.Attributes["site_id"], v.Primary.Attributes["id"]), nil
					}
				}
				return "", fmt.Errorf("not found in TestAccLogDrain import test step")
			},
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{"last_updated"},
		},
		{
			Config: `resource "netlify_log_drain" "example" {
  site_id     = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  destination = "http"
  log_types   = ["deploys"]
  format      = "ndjson"
  exclude_pii = true
  service_config = {
    url = "https://tools-httpstatus.pickup-services.com/200"
  }
}`,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction("netlify_log_drain.example", plancheck.ResourceActionUpdate),
				},
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_log_drain.example", "exclude_pii", "true"),
			),
		},
	}, testAccLogDrainDestroy)
}

func testAccLogDrainDestroy(s *terraform.State) error {
	for _, m := range s.Modules {
		if v, ok := m.Resources["netlify_log_drain.example"]; ok {
			key, _, err := testAccProvider.client.LogDrainsAPI.LogDrainsIndex(context.Background(), v.Primary.Attributes["site_id"]).Execute()
			if err != nil {
				//lint:ignore nilerr we expect an error to know it was not found
				return nil
			}
			return fmt.Errorf("Log drain still exists: %s", key.Id)
		}
	}
	return fmt.Errorf("not found in testAccLogDrainDestroy check destroy")
}
