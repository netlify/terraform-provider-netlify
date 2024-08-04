terraform {
  required_providers {
    netlify = {
      source = "registry.terraform.io/netlify/netlify"
    }
  }
  required_version = ">= 1.6.0"
}

# `token` comes from NETLIFY_API_TOKEN, but can be specified with a Terraform variable
provider "netlify" {
  default_team_slug = "ramon-test-1"
}

data "netlify_team" "current" {}

data "netlify_site" "platform_test" {
  name = "platform-test-1"
}

data "netlify_sites" "mine" {}

data "netlify_sites" "testing" {
  team_slug = "netlify-testing"
}

output "sites" {
  value = [
    for site in data.netlify_sites.testing.sites : site
    if site.custom_domain != ""
  ]
}
