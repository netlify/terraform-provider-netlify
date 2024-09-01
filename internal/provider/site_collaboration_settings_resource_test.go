package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSiteCollaborationSettings(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_collaboration_settings" "example" {
  site_id                           = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  netlify_drawer_in_deploy_previews = true
  netlify_drawer_in_branch_deploys  = true
  netlify_heads_up_display          = true
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_collaboration_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_collaboration_settings.example", "netlify_drawer_in_deploy_previews", "true"),
				resource.TestCheckResourceAttr("netlify_site_collaboration_settings.example", "netlify_drawer_in_branch_deploys", "true"),
				resource.TestCheckResourceAttr("netlify_site_collaboration_settings.example", "netlify_heads_up_display", "true"),
			),
		},
		{
			ResourceName:                         "netlify_site_collaboration_settings.example",
			ImportState:                          true,
			ImportStateId:                        "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}
