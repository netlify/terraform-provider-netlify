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
  site_id                = data.netlify_site.platform_test.id
  build_command          = "npm run build"
  publish_directory      = "dist"
  production_branch      = "main"
  branch_deploy_branches = ["meow", "woof"]
}

resource "netlify_site_domain_settings" "platform_test" {
  site_id                      = data.netlify_site.platform_test.id
  custom_domain                = "platform-test.example-tf-test-test.com"
  domain_aliases               = ["meow.example-tf-test-test.com"]
  branch_deploy_custom_domain  = "branch.example-tf-test-test.com"
  deploy_preview_custom_domain = "dp.example-tf-test-test.com"
}

resource "netlify_site_collaboration_settings" "platform_test" {
  site_id                           = data.netlify_site.platform_test.id
  netlify_drawer_in_deploy_previews = true
  netlify_drawer_in_branch_deploys  = true
  netlify_heads_up_display          = true
}
