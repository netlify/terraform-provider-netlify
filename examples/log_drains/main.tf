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

data "netlify_site" "test4" {
  account_slug = "netlify-testing"
  name         = "ramontest4"
}

resource "netlify_log_drain" "ramontest4" {
  site_id     = data.netlify_site.test4.id
  destination = "http"
  log_types   = ["user_traffic", "deploys", "edge_functions", "functions"]
  format      = "ndjson"
  exclude_pii = true
  service_config = {
    # https://webhook.site/#!/view/524008a8-bbdd-418e-8238-dd988b8d7d54
    url = "https://webhook.site/524008a8-bbdd-418e-8238-dd988b8d7d54"
  }
}
