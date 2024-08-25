package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var testAccProvider = &NetlifyProvider{version: "0.0.0-test"}

func accTest(t *testing.T, steps []resource.TestStep, checkDestroy resource.TestCheckFunc) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: map[string]func() (tfprotov6.ProviderServer, error){
			"netlify": providerserver.NewProtocol6WithError(testAccProvider),
		},
		CheckDestroy: checkDestroy,
		Steps:        steps,
	})
}

func testAccPreCheck(t *testing.T) {
	v := os.Getenv("NETLIFY_API_TOKEN")
	if v == "" {
		t.Fatal("NETLIFY_API_TOKEN must be set for acceptance tests")
	}
}
