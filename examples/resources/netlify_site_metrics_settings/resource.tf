resource "netlify_site_metrics_settings" "blog" {
  site_id           = data.netlify_site.blog.id
  site_analytics    = true
  real_user_metrics = true
}
