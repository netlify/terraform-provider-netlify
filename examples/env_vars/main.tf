terraform {
  required_providers {
    netlify = {
      source = "registry.terraform.io/netlify/netlify"
    }
  }
  required_version = ">= 1.6.0"
}

provider "netlify" {}

resource "netlify_environment_variable" "woof" {
  account_id = "65aaff3a32bdb71d9c4861d5"
  site_id    = "120a8a22-c806-4deb-b152-6e71b7ae3a16"
  key        = "WOOF"
  value = [
    {
      value   = "dogs are here",
      context = "all",
    }
  ]
}

resource "netlify_environment_variable" "meow" {
  account_id = "65aaff3a32bdb71d9c4861d5"
  site_id    = "120a8a22-c806-4deb-b152-6e71b7ae3a16"
  key        = "TEST_MEOW"
  value = [
    {
      value   = "roflmaocopter",
      context = "all",
    }
  ]
}

resource "netlify_secret_environment_variable" "meow" {
  account_id = "65aaff3a32bdb71d9c4861d5"
  site_id    = "120a8a22-c806-4deb-b152-6e71b7ae3a16"
  key        = "SECRET_TEST_MEOW"
  value = [
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
  account_id = "65aaff3a32bdb71d9c4861d5"
  key        = "TEST_MEOW"
  value = [
    {
      value   = "global roflmaocopter",
      context = "all",
    }
  ]
}
