
data "samsungcloudplatform_config_inspection_diagnosis" "my_diag_detail" {
  diagnosis_request_sequence = "SCPCIS-XXXXXXXXXX"
  diagnosis_id               = "DIA-XXXXXXXXXX"
}

output "output_config_inspection_diagnosis_detail" {
  value = data.samsungcloudplatform_config_inspection_diagnosis.my_diag_detail
}
