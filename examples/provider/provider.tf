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
  # Optionally, set a default team through its ID or its slug to avoid repeating it.
  default_team_slug = "your-team-slug"
}

data "netlify_team" "team" {
  # slug coming from the default team
}

data "netlify_site" "blog" {
  # team_slug coming from the default team
  name = "blog"
}

resource "netlify_environment_variable" "astro_database_file" {
  # team_id coming from the default team
  site_id = data.netlify_site.blog.id
  key     = "ASTRO_DATABASE_FILE"
  values = [
    {
      value   = "/path/here",
      context = "all",
    }
  ]
}
