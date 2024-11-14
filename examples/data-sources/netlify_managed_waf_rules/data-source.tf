# Read-only definitions of all managed WAF rules available in Netlify.
# The team ID is required to query the rules.
data "netlify_managed_waf_rules" "example" {
  team_id = "6600abcdef1234567890abcd"
}
