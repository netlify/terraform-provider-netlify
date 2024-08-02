# Site-level environment variable, note that both team_id and site_id are specified
resource "netlify_environment_variable" "astro_database_file" {
  team_id = data.netlify_team.team.id
  site_id = data.netlify_site.blog.id
  key     = "ASTRO_DATABASE_FILE"
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}

# Team-level environment variable, note that only team_id is specified
# Not supported on all Netlify plans
resource "netlify_environment_variable" "astro_database_file" {
  team_id = data.netlify_team.team.id
  key     = "ASTRO_DATABASE_FILE"
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}

# Secret environment variable
# Not supported on all Netlify plans
resource "netlify_environment_variable" "astro_studio_app_token" {
  team_id = data.netlify_team.team.id
  site_id = data.netlify_site.blog.id
  key     = "ASTRO_STUDIO_APP_TOKEN"
  secret_values = [
    {
      value   = "token-here",
      context = "all",
    }
  ]
}

# Values that differ by context
resource "netlify_environment_variable" "astro_studio_app_token" {
  team_id = data.netlify_team.team.id
  site_id = data.netlify_site.blog.id
  key     = "ASTRO_STUDIO_APP_TOKEN"
  secret_values = [
    {
      value   = "token-here",
      context = "production",
    },
    {
      value   = "non-prod-token-here",
      context = "deploy-preview",
    }
  ]
}

# A variable that's only available in some scopes, e.g. in builds
resource "netlify_environment_variable" "astro_database_file" {
  team_id = data.netlify_team.team.id
  site_id = data.netlify_site.blog.id
  key     = "ASTRO_DATABASE_FILE"
  scopes  = ["builds"]
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}
