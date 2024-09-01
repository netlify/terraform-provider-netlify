package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccEnvVar(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_environment_variable" "site_level" {
  team_id = "66ae34e11a567e9092e3850f"
  site_id = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  key     = "TEST_SITE_LEVEL"
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "team_level" {
  team_id = "66ae34e11a567e9092e3850f"
  key     = "TEST_TEAM_LEVEL"
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}
`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_environment_variable.site_level", "key", "TEST_SITE_LEVEL"),
				resource.TestCheckResourceAttr("netlify_environment_variable.team_level", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckNoResourceAttr("netlify_environment_variable.team_level", "site_id"),
				resource.TestCheckResourceAttr("netlify_environment_variable.team_level", "key", "TEST_TEAM_LEVEL"),
			),
		},
		{
			ResourceName: "netlify_environment_variable.site_level",
			ImportState:  true,
			ImportStateIdFunc: func(s *terraform.State) (string, error) {
				for _, m := range s.Modules {
					if v, ok := m.Resources["netlify_environment_variable.site_level"]; ok {
						return fmt.Sprintf("%s:%s:%s", v.Primary.Attributes["team_id"], v.Primary.Attributes["site_id"], v.Primary.Attributes["key"]), nil
					}
				}
				return "", fmt.Errorf("not found in TestAccEnvVar import test step")
			},
			ImportStateVerifyIdentifierAttribute: "key",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
		{
			ResourceName: "netlify_environment_variable.team_level",
			ImportState:  true,
			ImportStateIdFunc: func(s *terraform.State) (string, error) {
				for _, m := range s.Modules {
					if v, ok := m.Resources["netlify_environment_variable.team_level"]; ok {
						return fmt.Sprintf("%s:%s", v.Primary.Attributes["team_id"], v.Primary.Attributes["key"]), nil
					}
				}
				return "", fmt.Errorf("not found in TestAccEnvVar import test step")
			},
			ImportStateVerifyIdentifierAttribute: "key",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}

func TestAccEnvVarComplex(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_environment_variable" "a" {
  team_id = "66ae34e11a567e9092e3850f"
  site_id = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  key     = "C_A"
  scopes  = ["builds"]
  secret_values = [
    {
      value   = "token-here",
      context = "production",
    },
    {
      value   = "non-prod-token-here",
      context = "deploy-preview",
    }
  ]
}

resource "netlify_environment_variable" "b" {
  team_id = "66ae34e11a567e9092e3850f"
  site_id = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  key     = "C_B"
  values = [
    {
      value             = "staging"
      context           = "branch"
      context_parameter = "staging"
    },
    {
      value             = "onboarding"
      context           = "branch"
      context_parameter = "onboarding"
    },
    {
      value   = "production"
      context = "production"
    },
    {
      value   = "branch-deploy"
      context = "branch-deploy"
    },
    {
      value   = "deploy-preview"
      context = "deploy-preview"
    }
  ]
}
`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_environment_variable.a", "team_id", "66ae34e11a567e9092e3850f"),
				resource.TestCheckResourceAttr("netlify_environment_variable.b", "team_id", "66ae34e11a567e9092e3850f"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}
