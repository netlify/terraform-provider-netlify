terraform {
  required_providers {
    netlify = {
      source = "registry.terraform.io/netlify/netlify"
    }
  }
  required_version = ">= 1.6.0"
}

# `token` comes from NETLIFY_API_TOKEN, but can be specified with a Terraform variable
provider "netlify" {}

data "netlify_site" "platform_test" {
  account_slug = "ramon-test-1"
  name         = "platform-test-1"
}

resource "netlify_site_build_settings" "platform_test" {
  site_id           = data.netlify_site.platform_test.id
  build_command     = "npm run build"
  publish_directory = "dist"
}
