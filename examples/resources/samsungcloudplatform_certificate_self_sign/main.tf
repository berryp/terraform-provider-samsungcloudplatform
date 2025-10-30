resource "samsungcloudplatform_certificate_self_sign" "my_selfSign" {
  name = var.name
  common_name = "sds"
  organization_name = "test"
  start_date = "2025-07-16"
  expiration_date = "2025-07-19"
  recipients {
  }
}

