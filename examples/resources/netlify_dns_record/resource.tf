resource "netlify_dns_record" "www" {
  type     = "A"
  zone_id  = netlify_dns_zone.example.id
  hostname = "www.example.com"
  value    = "198.18.0.50"
}

resource "netlify_dns_record" "calendar" {
  type     = "CNAME"
  zone_id  = netlify_dns_zone.example.id
  hostname = "calendar.example.com"
  value    = "ghs.googlehosted.com."
}

resource "netlify_dns_record" "mx" {
  type     = "MX"
  zone_id  = netlify_dns_zone.example.id
  hostname = "example.com"
  value    = "smtp.google.com"
  priority = 1
}
