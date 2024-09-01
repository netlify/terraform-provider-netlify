package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccDataTeam(t *testing.T) {
	accTest(t, []resource.TestStep{
		{
			Config: `data "netlify_team" "example_by_id" {
	id = "66ae34e11a567e9092e3850f"
}

data "netlify_team" "example_by_slug" {
	id = "netlify-terraform-test"
}`,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("data.netlify_team.example_by_id", "slug", "netlify-terraform-test"),
				resource.TestCheckResourceAttr("data.netlify_team.example_by_id", "name", "Netlify Terraform Test"),
				resource.TestCheckResourceAttr("data.netlify_team.example_by_slug", "id", "66ae34e11a567e9092e3850f"),
			),
		},
	}, func(s *terraform.State) error { return nil })
}
