resource "netlify_site_collaboration_settings" "blog" {
  site_id                           = data.netlify_site.blog.id
  netlify_drawer_in_deploy_previews = true
  netlify_drawer_in_branch_deploys  = true
  netlify_heads_up_display          = true
}
