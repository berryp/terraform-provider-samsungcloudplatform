data "samsungcloudplatform_config_inspection_diagnosis_request" "my_diag_request" {
  project_id           = "PROJECT-XXXXXXXXXX"
  access_key           = ""
  diagnosis_check_type = "BP"
  diagnosis_id         = "DIA-XXXXXXXXXX"
}

output "output_config_inspection_diagnosis_request" {
  value = data.samsungcloudplatform_config_inspection_diagnosis_request.my_diag_request.result
}
