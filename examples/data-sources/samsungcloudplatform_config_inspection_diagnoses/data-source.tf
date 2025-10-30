
data "samsungcloudplatform_config_inspection_diagnoses" "my_diagnosis_list" {

}

output "output_config_inspection_diagnoses" {
  value = {
    contents : data.samsungcloudplatform_config_inspection_diagnoses.my_diagnosis_list.contents,
    total_count : data.samsungcloudplatform_config_inspection_diagnoses.my_diagnosis_list.total_count,
  }
}
