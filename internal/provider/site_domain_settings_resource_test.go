package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSiteDomainSettings(t *testing.T) {
	// TODO: change domain and domain aliases, and wait for the certificate to update
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_domain_settings" "example" {
  site_id                      = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  custom_domain                = "tf-test-1.examplepetstore.com"
  domain_aliases               = ["tf-test-1-alias.examplepetstore.com"]
  branch_deploy_custom_domain  = "tf-test-12-branch.examplepetstore.com"
  deploy_preview_custom_domain = "tf-test-12-dp.examplepetstore.com"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "custom_domain", "tf-test-1.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.#", "1"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.0", "tf-test-1-alias.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "branch_deploy_custom_domain", "tf-test-12-branch.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "deploy_preview_custom_domain", "tf-test-12-dp.examplepetstore.com"),
			),
		},
		{
			Config: `resource "netlify_site_domain_settings" "example" {
  site_id                      = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  custom_domain                = "tf-test-1.examplepetstore.com"
  domain_aliases               = ["tf-test-1-alias.examplepetstore.com"]
  branch_deploy_custom_domain  = "tf-test-1-branch.examplepetstore.com"
  deploy_preview_custom_domain = "tf-test-1-dp.examplepetstore.com"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "custom_domain", "tf-test-1.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.#", "1"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.0", "tf-test-1-alias.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "branch_deploy_custom_domain", "tf-test-1-branch.examplepetstore.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "deploy_preview_custom_domain", "tf-test-1-dp.examplepetstore.com"),
			),
		},
		{
			ResourceName:                         "netlify_site_domain_settings.example",
			ImportState:                          true,
			ImportStateId:                        "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}
