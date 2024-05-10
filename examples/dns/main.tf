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

# data "netlify_account" "current" {
#   slug = "ramon-test-1"
# }

resource "netlify_dns_zone" "example" {
  account_slug = "ramon-test-1" // data.netlify_account.current.slug
  name         = "example-tf-test-test.com"
  lifecycle {
    prevent_destroy = true
  }
}

resource "netlify_dns_a_record" "cat" {
  zone_id  = netlify_dns_zone.example.id
  hostname = "cat.example-tf-test-test.com"
  value    = "10.0.0.15"
}

resource "netlify_dns_cname_record" "dog" {
  zone_id  = netlify_dns_zone.example.id
  hostname = "dog.example-tf-test-test.com"
  value    = "cat.example-tf-test-test.com"
  ttl      = 60
}

resource "netlify_dns_txt_record" "bird" {
  zone_id  = netlify_dns_zone.example.id
  hostname = "bird.example-tf-test-test.com"
  value    = "hello world"
}

resource "netlify_dns_mx_record" "fish" {
  zone_id  = netlify_dns_zone.example.id
  hostname = "fish.example-tf-test-test.com"
  value    = "mail.example-tf-test-test.com"
  priority = 10
}
