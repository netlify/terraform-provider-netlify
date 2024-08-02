variable "netlify_api_token" {
  type = string
}

terraform {
  required_providers {
    netlify = {
      source = "registry.terraform.io/netlify/netlify"
    }
  }
}

provider "netlify" {
  token = var.netlify_api_token
}

data "netlify_team" "team" {
  slug = "your-team-slug"
}

data "netlify_site" "blog" {
  team_slug = data.netlify_team.team.slug
  name      = "blog"
}

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
