# Import a team-level environment variable using the team ID and the environment variable key
terraform import netlify_environment_variable.name 6600abcdef1234567890abcd:ASTRO_DATABASE_FILE

# Import a site-level environment variable using the team ID, the site ID, and the environment variable key
terraform import netlify_environment_variable.name 6600abcdef1234567890abcd:12345667-0000-0000-0000-abcdef012345:ASTRO_DATABASE_FILE
