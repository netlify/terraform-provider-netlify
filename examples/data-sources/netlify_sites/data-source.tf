# List all sites in a team, by the team's slug
data "netlify_sites" "team" {
  team_slug = "my-team-slug"
}
