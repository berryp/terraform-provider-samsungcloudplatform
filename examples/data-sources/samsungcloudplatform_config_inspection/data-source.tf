data "samsungcloudplatform_config_inspection" "my_config_inspection_detail" {
  diagnosis_id = "DIA-XXXXXXXXXX"
}

output "output_config_inspection_detail" {
  value = data.samsungcloudplatform_config_inspection.my_config_inspection_detail
}
