resource "netlify_dns_zone" "example" {
  team_slug = data.netlify_team.team.slug
  name      = "example.com"
  lifecycle {
    prevent_destroy = true
  }
}
