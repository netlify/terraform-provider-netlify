---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "netlify_team Data Source - netlify"
subcategory: ""
description: |-
  
---

# netlify_team (Data Source)



## Example Usage

```terraform
# Looking up a team by its slug
data "netlify_team" "team" {
  slug = "my-team-slug"
}

# Looking up a team by its ID
data "netlify_team" "team" {
  id = "6600abcdef1234567890abcd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) ID or slug are required if a default team was not configured in the provider configuration.
- `slug` (String) ID or slug are required if a default team was not configured in the provider configuration.

### Read-Only

- `name` (String)
