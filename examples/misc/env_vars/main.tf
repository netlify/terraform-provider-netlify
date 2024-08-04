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

data "netlify_team" "current" {
  slug = "ramon-test-1"
}

data "netlify_site" "platform_test" {
  team_slug = data.netlify_team.current.slug
  name      = "platform-test-1"
}

resource "netlify_environment_variable" "woof" {
  team_id = data.netlify_team.current.id
  site_id = data.netlify_site.platform_test.id
  key     = "WOOF"
  values = [
    {
      value   = "dogs are here",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "woof2" {
  site_id = data.netlify_site.platform_test.id
  key     = "WOOF2"
  values = [
    {
      value   = "dogs are here",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "meow" {
  team_id = data.netlify_team.current.id
  site_id = data.netlify_site.platform_test.id
  key     = "TEST_MEOW"
  values = [
    {
      value   = "roflmaocopter",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "secret_meow" {
  team_id = data.netlify_team.current.id
  site_id = data.netlify_site.platform_test.id
  key     = "SECRET_TEST_MEOW"
  secret_values = [
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
  team_id = data.netlify_team.current.id
  key     = "TEST_MEOW"
  values = [
    {
      value   = "global roflmaocopter",
      context = "all",
    }
  ]
}
