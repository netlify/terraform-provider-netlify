resource "netlify_site_build_settings" "blog" {
  site_id                      = data.netlify_site.blog.id
  build_command                = "npm run build"
  publish_directory            = "dist"
  production_branch            = "main"
  branch_deploy_branches       = ["preview", "staging"]
  prevent_non_git_prod_deploys = true
}
