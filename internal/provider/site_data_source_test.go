package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataSite(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `data "netlify_site" "example_by_id" {
	id = "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"
}

data "netlify_site" "example_by_name" {
	team_slug = "netlify-terraform-test"
	name      = "tf-test-1"
}`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.netlify_site.example_by_id", "name", "tf-test-1"),
				resource.TestCheckResourceAttr("data.netlify_site.example_by_name", "id", "5b407d6d-9385-4e7a-a4c4-8efc11ea3c26"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}
