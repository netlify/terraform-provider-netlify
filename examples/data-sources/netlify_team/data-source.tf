# Looking up a team by its slug
data "netlify_team" "team" {
  slug = "my-team-slug"
}

# Looking up a team by its ID
data "netlify_team" "team" {
  id = "6600abcdef1234567890abcd"
}
