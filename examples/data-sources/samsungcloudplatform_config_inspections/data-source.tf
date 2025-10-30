
data "samsungcloudplatform_config_inspections" "config_inspection" {
}

output "output_config_inspection" {
  value = {
    contents    = data.samsungcloudplatform_config_inspections.config_inspection.contents
    total_count = data.samsungcloudplatform_config_inspections.config_inspection.total_count
  }
}
