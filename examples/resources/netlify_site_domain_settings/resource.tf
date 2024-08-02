resource "netlify_site_domain_settings" "blog" {
  site_id                      = data.netlify_site.blog.id
  custom_domain                = "blog.example.com"
  domain_aliases               = ["blog-alias.example.com"]
  branch_deploy_custom_domain  = "blog-branch.example.com"
  deploy_preview_custom_domain = "blog-dp.example.com"
}
