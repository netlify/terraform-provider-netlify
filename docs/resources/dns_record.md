---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netlify_dns_record Resource - netlify"
subcategory: ""
description: |-
  Netlify DNS record. Read more https://docs.netlify.com/domains-https/netlify-dns/
---

# netlify_dns_record (Resource)

Netlify DNS record. [Read more](https://docs.netlify.com/domains-https/netlify-dns/)

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `hostname` (String) The hostname for the DNS record. For example, `www.example.com`.
- `type` (String) One of A, AAAA, ALIAS, CAA, CNAME, MX, NS, SPF, or TXT
- `value` (String)
- `zone_id` (String)

### Optional

- `flag` (Number)
- `priority` (Number)
- `tag` (String)
- `ttl` (Number)

### Read-Only

- `id` (String) The ID of this resource.
- `last_updated` (String)

## Import

Import is supported using the following syntax:

```shell
# Import a DNS record by its zone ID and its record ID
terraform import netlify_dns_record.www_example 6600abcdef1234567890abcd:6600abcdef1234567890abcd
```
