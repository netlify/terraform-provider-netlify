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
  custom_domain                = "tf-test-1.nf-terraform-test.com"
  domain_aliases               = ["tf-test-1-alias.nf-terraform-test.com"]
  branch_deploy_custom_domain  = "tf-test-12-branch.nf-terraform-test.com"
  deploy_preview_custom_domain = "tf-test-12-dp.nf-terraform-test.com"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "custom_domain", "tf-test-1.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.#", "1"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.0", "tf-test-1-alias.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "branch_deploy_custom_domain", "tf-test-12-branch.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "deploy_preview_custom_domain", "tf-test-12-dp.nf-terraform-test.com"),
			),
		},
		{
			Config: `resource "netlify_site_domain_settings" "example" {
  site_id                      = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
  custom_domain                = "tf-test-1.nf-terraform-test.com"
  domain_aliases               = ["tf-test-1-alias.nf-terraform-test.com"]
  branch_deploy_custom_domain  = "tf-test-1-branch.nf-terraform-test.com"
  deploy_preview_custom_domain = "tf-test-1-dp.nf-terraform-test.com"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "site_id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "custom_domain", "tf-test-1.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.#", "1"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "domain_aliases.0", "tf-test-1-alias.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "branch_deploy_custom_domain", "tf-test-1-branch.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.example", "deploy_preview_custom_domain", "tf-test-1-dp.nf-terraform-test.com"),
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

func TestAccSiteDomainSettingsNullFields(t *testing.T) {
	// Test case for preventing drift when optional fields are null from API
	accTest(t, []resource.TestStep{
		{
			Config: `resource "netlify_site_domain_settings" "null_fields" {
  site_id       = "605830b8-16e3-4c9e-9010-be3fef9e977b"
  custom_domain = "tf-test-null-fields.nf-terraform-test.com"
}`,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_domain_settings.null_fields", "site_id", "605830b8-16e3-4c9e-9010-be3fef9e977b"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.null_fields", "custom_domain", "tf-test-null-fields.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.null_fields", "domain_aliases.#", "0"),
				// These should be null/empty when not set
				resource.TestCheckNoResourceAttr("netlify_site_domain_settings.null_fields", "branch_deploy_custom_domain"),
				resource.TestCheckNoResourceAttr("netlify_site_domain_settings.null_fields", "deploy_preview_custom_domain"),
			),
		},
		{
			// Second step with same config should not detect changes (no drift)
			Config: `resource "netlify_site_domain_settings" "null_fields" {
  site_id       = "605830b8-16e3-4c9e-9010-be3fef9e977b"
  custom_domain = "tf-test-null-fields.nf-terraform-test.com"
}`,
			PlanOnly: true,
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("netlify_site_domain_settings.null_fields", "site_id", "605830b8-16e3-4c9e-9010-be3fef9e977b"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.null_fields", "custom_domain", "tf-test-null-fields.nf-terraform-test.com"),
				resource.TestCheckResourceAttr("netlify_site_domain_settings.null_fields", "domain_aliases.#", "0"),
			),
		},
		{
			ResourceName:                         "netlify_site_domain_settings.null_fields",
			ImportState:                          true,
			ImportStateId:                        "605830b8-16e3-4c9e-9010-be3fef9e977b",
			ImportStateVerifyIdentifierAttribute: "site_id",
			ImportStateVerify:                    true,
			ImportStateVerifyIgnore:              []string{"last_updated"},
		},
	}, func(s *terraform.State) error { return nil })
}
