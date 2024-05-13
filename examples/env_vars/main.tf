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

data "netlify_account" "current" {
  slug = "ramon-test-1"
}

data "netlify_site" "platform_test" {
  account_slug = data.netlify_account.current.slug
  name         = "platform-test-1"
}

resource "netlify_environment_variable" "woof" {
  account_id = data.netlify_account.current.id
  site_id    = data.netlify_site.platform_test.id
  key        = "WOOF"
  value = [
    {
      value   = "dogs are here",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "meow" {
  account_id = data.netlify_account.current.id
  site_id    = data.netlify_site.platform_test.id
  key        = "TEST_MEOW"
  value = [
    {
      value   = "roflmaocopter",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "secret_meow" {
  account_id = data.netlify_account.current.id
  site_id    = data.netlify_site.platform_test.id
  key        = "SECRET_TEST_MEOW"
  secret_value = [
    {
      value   = "secret roflmaocopter",
      context = "production",
    },
    {
      value   = "secret roflmaocopter",
      context = "deploy-preview",
    }
  ]
}

resource "netlify_environment_variable" "global_meow" {
  account_id = data.netlify_account.current.id
  key        = "TEST_MEOW"
  value = [
    {
      value   = "global roflmaocopter",
      context = "all",
    }
  ]
}
