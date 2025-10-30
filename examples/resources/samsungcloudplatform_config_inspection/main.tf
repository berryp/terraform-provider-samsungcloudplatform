
resource "samsungcloudplatform_config_inspection" "my_inspections_001" {
  auth_key_request {
    auth_key_id = var.auth_key_request.auth_key_id
  }
  csp_type             = var.csp_type
  diagnosis_check_type = var.diagnosis_check_type
  diagnosis_type       = var.diagnosis_type
  plan_type            = var.plan_type
  schedule_request {
    diagnosis_start_time_pattern = var.schedule_request.diagnosis_start_time_pattern
    frequency_type               = var.schedule_request.frequency_type
    frequency_value              = var.schedule_request.frequency_value
    use_diagnosis_check_type_bp  = var.schedule_request.use_diagnosis_check_type_bp
    use_diagnosis_check_type_ssi = var.schedule_request.use_diagnosis_check_type_ssi
  }
  dynamic "diagnosis_object_request_list" {
    for_each = var.diagnosis_object_request_list
    content {
      diagnosis_account_id = diagnosis_object_request_list.value.diagnosis_account_id
      diagnosis_id         = diagnosis_object_request_list.value.diagnosis_id
      diagnosis_name       = diagnosis_object_request_list.value.diagnosis_name
    }
  }
  tags = var.tags
}
