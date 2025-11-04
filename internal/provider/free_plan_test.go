package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccFreeEnvVar(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_environment_variable" "site_level" {
  team_id = "66e98216e3fe031846dc998a"
  site_id = "fbba82b0-f1e9-4e92-9203-eefc62857545"
  key     = "TEST_SITE_LEVEL"
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}
`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level", "team_id", "66e98216e3fe031846dc998a"),
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level", "site_id", "fbba82b0-f1e9-4e92-9203-eefc62857545"),
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level", "key", "TEST_SITE_LEVEL"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}

func TestAccFreeSecretEnvVar(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_environment_variable" "site_level_secret" {
  team_id = "66e98216e3fe031846dc998a"
  site_id = "fbba82b0-f1e9-4e92-9203-eefc62857545"
  key     = "TEST_SITE_LEVEL_SECRET"
	scopes  = ["functions", "builds", "runtime"]
  secret_values = [
    {
      value   = "ill-never-tell",
      context = "production",
    }
  ]
}
`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level_secret", "team_id", "66e98216e3fe031846dc998a"),
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level_secret", "site_id", "fbba82b0-f1e9-4e92-9203-eefc62857545"),
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level_secret", "key", "TEST_SITE_LEVEL_SECRET"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}

func TestAccFreeSiteBuildSettings(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_build_settings" "example" {
  site_id                = "3a3c2d0f-4e59-46c0-ae6a-4bdc0d49aca9"
  build_command          = "npm run build"
  publish_directory      = "dist"
  production_branch      = "main"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "site_id", "3a3c2d0f-4e59-46c0-ae6a-4bdc0d49aca9"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "build_command", "npm run build"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "publish_directory", "dist"),
				resource.TestCheckResourceAttr("netlify_site_build_settings.example", "production_branch", "main"),
			),
		},
		{
			ResourceName:                         "netlify_site_build_settings.example",
			ImportState:                          true,
			ImportStateId:                        "3a3c2d0f-4e59-46c0-ae6a-4bdc0d49aca9",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}
