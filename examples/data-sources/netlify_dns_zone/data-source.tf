# Read-only data source for a Netlify DNS zone.
data "netlify_dns_zone" "example" {
  name = "example.com"
}
