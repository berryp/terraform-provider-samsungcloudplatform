data "terraform_remote_state" "engine" {
  backend = "local"

  config = {
    path = "../samsungcloudplatform_kubernetes_engine/terraform.tfstate"
  }
}

variable "kubeconfig_type" {
  default = "private"
  description = "private/public"
}
