package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSiteBuildSettings(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_build_settings" "example" {
  site_id                = "49137d35-1470-4db1-810f-c185b8381cd3"
  build_command          = "npm run build && true"
  publish_directory      = "dist/dist"
  production_branch      = "preview"
  branch_deploy_branches = ["staging"]
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "site_id", "49137d35-1470-4db1-810f-c185b8381cd3"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "build_command", "npm run build && true"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "publish_directory", "dist/dist"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "production_branch", "preview"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "branch_deploy_branches.#", "1"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "branch_deploy_branches.0", "staging"),
			),
		},
		{
			Config: `resource "netlify_site_build_settings" "example" {
  site_id                = "49137d35-1470-4db1-810f-c185b8381cd3"
  build_command          = "npm run build"
  publish_directory      = "dist"
  production_branch      = "main"
  branch_deploy_branches = ["preview", "staging"]
  waf_policy_id          = netlify_waf_policy.example.id
}

resource "netlify_waf_policy" "example" {
  team_id     = "66ae34e11a567e9092e3850f"
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = true,
      overall_threshold = 5
    }
  ]
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "site_id", "49137d35-1470-4db1-810f-c185b8381cd3"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "build_command", "npm run build"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "publish_directory", "dist"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "production_branch", "main"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "branch_deploy_branches.#", "2"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "branch_deploy_branches.0", "preview"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "branch_deploy_branches.1", "staging"),
				resource.TestCheckResourceAttrSet("netlify_site_build_settings.example", "waf_policy_id"),
			),
		},
		{
			ResourceName:                         "netlify_site_build_settings.example",
			ImportState:                          true,
			ImportStateId:                        "49137d35-1470-4db1-810f-c185b8381cd3",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}
