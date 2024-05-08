terraform {
  required_providers {
    netlify = {
      source = "registry.terraform.io/netlify/netlify"
    }
  }
  required_version = ">= 1.6.0"
}

provider "netlify" {}

# resource "netlify_environment_variable" "woof" {
#   account_id = "ramon-test-1"
#   site_id    = "platform-test-1"
#   key        = "WOOF"
#   value = [
#     {
#       value   = "dogs are here",
#       context = "all",
#     }
#   ]
# }

resource "netlify_environment_variable" "meow" {
  account_id = "65aaff3a32bdb71d9c4861d5"
  site_id    = "120a8a22-c806-4deb-b152-6e71b7ae3a16"
  key        = "TEST_MEOW"
  value = [
    {
      value   = "roflmao",
      context = "all",
    }
  ]
}
