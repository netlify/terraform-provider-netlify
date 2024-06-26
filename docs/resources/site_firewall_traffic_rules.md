---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netlify_site_firewall_traffic_rules Resource - netlify"
subcategory: ""
description: |-
  
---

# netlify_site_firewall_traffic_rules (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `published` (Attributes) (see [below for nested schema](#nestedatt--published))
- `site_id` (String)
- `unpublished` (Attributes) (see [below for nested schema](#nestedatt--unpublished))

### Read-Only

- `last_updated` (String)
- `team_id` (String)

<a id="nestedatt--published"></a>
### Nested Schema for `published`

Required:

- `default_action` (String)

Optional:

- `geo_exceptions` (Attributes List) (see [below for nested schema](#nestedatt--published--geo_exceptions))
- `geo_restrictions` (Attributes List) (see [below for nested schema](#nestedatt--published--geo_restrictions))
- `ip_exceptions` (Attributes List) (see [below for nested schema](#nestedatt--published--ip_exceptions))
- `ip_restrictions` (Attributes List) (see [below for nested schema](#nestedatt--published--ip_restrictions))

<a id="nestedatt--published--geo_exceptions"></a>
### Nested Schema for `published.geo_exceptions`

Required:

- `countries` (List of String)
- `description` (String)

Optional:

- `subregions` (List of String)


<a id="nestedatt--published--geo_restrictions"></a>
### Nested Schema for `published.geo_restrictions`

Required:

- `countries` (List of String)
- `description` (String)

Optional:

- `subregions` (List of String)


<a id="nestedatt--published--ip_exceptions"></a>
### Nested Schema for `published.ip_exceptions`

Required:

- `addresses` (List of String)
- `description` (String)


<a id="nestedatt--published--ip_restrictions"></a>
### Nested Schema for `published.ip_restrictions`

Required:

- `addresses` (List of String)
- `description` (String)



<a id="nestedatt--unpublished"></a>
### Nested Schema for `unpublished`

Required:

- `default_action` (String)

Optional:

- `geo_exceptions` (Attributes List) (see [below for nested schema](#nestedatt--unpublished--geo_exceptions))
- `geo_restrictions` (Attributes List) (see [below for nested schema](#nestedatt--unpublished--geo_restrictions))
- `ip_exceptions` (Attributes List) (see [below for nested schema](#nestedatt--unpublished--ip_exceptions))
- `ip_restrictions` (Attributes List) (see [below for nested schema](#nestedatt--unpublished--ip_restrictions))

<a id="nestedatt--unpublished--geo_exceptions"></a>
### Nested Schema for `unpublished.geo_exceptions`

Required:

- `countries` (List of String)
- `description` (String)

Optional:

- `subregions` (List of String)


<a id="nestedatt--unpublished--geo_restrictions"></a>
### Nested Schema for `unpublished.geo_restrictions`

Required:

- `countries` (List of String)
- `description` (String)

Optional:

- `subregions` (List of String)


<a id="nestedatt--unpublished--ip_exceptions"></a>
### Nested Schema for `unpublished.ip_exceptions`

Required:

- `addresses` (List of String)
- `description` (String)


<a id="nestedatt--unpublished--ip_restrictions"></a>
### Nested Schema for `unpublished.ip_restrictions`

Required:

- `addresses` (List of String)
- `description` (String)
