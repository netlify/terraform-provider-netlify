# Looking up a site by its team slug and site name
data "netlify_site" "blog" {
  team_slug = "my-team-slug"
  name      = "Blog"
}

# Looking up a blog by its ID
data "netlify_site" "blog" {
  id = "12345667-0000-0000-0000-abcdef012345"
}
