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
  token = "nfp_C92ECMDqhbZpjtbhGmzGT2GQnFJSA4zvcc22"
}

data "netlify_site" "test4" {
  account_slug = "netlify-testing"
  name         = "ramontest4"
}

# TODO: add an example for netlify_account_firewall_traffic_rules

resource "netlify_site_firewall_traffic_rules" "ramontest4" {
  site_id = data.netlify_site.test4.id
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
