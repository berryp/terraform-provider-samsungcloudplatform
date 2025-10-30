variable "auth_key_request" {
  description = ""
  type = object({
    auth_key_id = string
  })
  default = { "auth_key_id" : "ACCESSKEY-icH4wCkXqspG3xCFh4QhSc" }
}

variable "csp_type" {
  default = "SCP"
}

variable "diagnosis_check_type" {
  default = "BP"
}

variable "diagnosis_object_request_list" {
  description = ""
  type = list(object({
    diagnosis_account_id = string
    diagnosis_id         = string
    diagnosis_name       = string
  }))
  default = [{
    "diagnosis_account_id" : "PROJECT-s3Gg1NnWrGiPR0JefMj7yh",
    "diagnosis_id" : "",
    "diagnosis_name" : "sds-data-03123",
  }]
}

variable "diagnosis_type" {
  default = "SSI"
}

variable "plan_type" {
  default = "STANDARD"
}

variable "schedule_request" {
  description = ""
  type = object({
    diagnosis_start_time_pattern = string
    frequency_type               = string
    frequency_value              = string
    use_diagnosis_check_type_bp  = string
    use_diagnosis_check_type_ssi = string
  })
  default = {
    "diagnosis_start_time_pattern" : "00:00",
    "frequency_type" : "month",
    "frequency_value" : "01",
    "use_diagnosis_check_type_bp" : "y",
    "use_diagnosis_check_type_ssi" : "n",
  }
}

variable "tags" {
  type = map(string)
  default = {
    "tk01" = "tv01"
    "tk02" = "tv02"
  }
}
