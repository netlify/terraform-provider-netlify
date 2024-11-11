package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSiteMetricsSettings(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_metrics_settings" "example" {
  site_id          = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  site_analytics   = false
  real_user_metrics = false
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_metrics_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_metrics_settings.example", "site_analytics", "false"),
				resource.TestCheckResourceAttr("netlify_site_metrics_settings.example", "real_user_metrics", "false"),
			),
		},
		{
			Config: `resource "netlify_site_metrics_settings" "example" {
  site_id          = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  site_analytics   = false # on purpose due to 10 minutes delay
  real_user_metrics = true
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_metrics_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_metrics_settings.example", "site_analytics", "false"),
				resource.TestCheckResourceAttr("netlify_site_metrics_settings.example", "real_user_metrics", "true"),
			),
		},
		{
			ResourceName:                         "netlify_site_metrics_settings.example",
			ImportState:                          true,
			ImportStateId:                        "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}
