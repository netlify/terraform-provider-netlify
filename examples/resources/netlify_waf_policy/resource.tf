resource "netlify_waf_policy" "example" {
  team_id     = data.netlify_team.team.id
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = true,
      overall_threshold = 5,
      category_thresholds = {
        "fixation" = 8,
      },
      rule_overrides = {
        "920100" = {
          action = "log_only"
        }
      }
    }
  ]
}

# To use this policy in a site, use the netlify_site_build_settings resource:

resource "netlify_site_build_settings" "example" {
  # other attributes...
  waf_policy_id = netlify_waf_policy.example.id
}

# To dynamically define the rule overrides, you can query netlify_managed_waf_rules to get the rule IDs:

data "netlify_managed_waf_rules" "example" {
  team_id = "6600abcdef1234567890abcd"
}

resource "netlify_waf_policy" "example" {
  team_id     = "66ae34e11a567e9092e3850f"
  name        = "Terraform Policy"
  description = "This is a test policy through Terraform"
  rule_sets = [
    {
      managed_id        = "crs-basic",
      passive_mode      = true,
      overall_threshold = 5,
      rule_overrides = {
        for rule in data.netlify_managed_waf_rules.example.rule_sets["crs-basic"].rules : rule.id => {
          action = "log_only"
        } if rule.category == "rce"
      }
    }
  ]
}
