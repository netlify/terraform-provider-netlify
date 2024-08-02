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

data "netlify_team" "current" {
  slug = "ramon-test-1"
}

data "netlify_site" "platform_test" {
  team_slug = data.netlify_team.current.slug
  name      = "platform-test-1"
}

data "netlify_sites" "all" {
  team_slug = "netlify-testing"
}

output "sites" {
  value = [
    for site in data.netlify_sites.all.sites : site
    if site.custom_domain != ""
  ]
}
