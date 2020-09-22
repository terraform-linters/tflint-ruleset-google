resource "google_app_engine_domain_mapping" "domain_mapping" {
  domain_name       = "verified-domain.com"
  override_strategy = "DEFAULT"

  ssl_settings {
    ssl_management_type = "AUTOMATIC"
  }
}
