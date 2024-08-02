resource "netlify_log_drain" "blog" {
  site_id     = data.netlify_site.blog.id
  destination = "http"
  log_types   = ["user_traffic", "deploys", "edge_functions", "functions"]
  format      = "ndjson"
  exclude_pii = true
  service_config = {
    url = "https://destinationurl/"
  }
}
