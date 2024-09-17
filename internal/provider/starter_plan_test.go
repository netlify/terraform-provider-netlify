package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccStarterEnvVar(t *testing.T) {
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
