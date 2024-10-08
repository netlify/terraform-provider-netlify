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

data "netlify_team" "team" {
  slug = "netlify-terraform-test"
}

resource "netlify_team_firewall_traffic_rules" "team" {
  team_id = data.netlify_team.team.id
  published = {
    default_action = "allow"
    ip_restrictions = [
      {
        description = "Meow"
        addresses   = ["173.54.6.0/30"]
      },
      {
        description = "bot network"
        addresses = [
          "90.12.4.1/32",
          "90.12.4.2/32",
          "90.12.5.5/32",
        ]
      }
    ]
    geo_exceptions = [
      {
        description = "brazil"
        countries   = ["BR"]
      }
    ]
  }
  unpublished = {
    default_action = "deny"
    ip_exceptions = [
      {
        description = "Allow my IP"
        addresses   = ["71.105.184.66/32"]
      }
    ]
  }
}
