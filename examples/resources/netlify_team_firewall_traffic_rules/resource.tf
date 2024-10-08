resource "netlify_team_firewall_traffic_rules" "team" {
  team_id = data.netlify_team.team.id
  published = {
    default_action = "allow"
    ip_restrictions = [
      {
        description = "bot network"
        addresses = [
          "192.0.2.0/24",
          "198.51.100.0/24",
        ]
      }
    ]
    geo_exceptions = [
      {
        description = "brazil"
        countries   = ["BR"]
      }
    ]
  }
  unpublished = {
    default_action = "deny"
    ip_exceptions = [
      {
        description = "Allow the VPN IP"
        addresses   = ["203.0.113.65/32"]
      }
    ]
  }
}
