resource "samsungcloudplatform_kubernetes_admin_kubeconfig" "admin_kubeconfig" {
  kubernetes_engine_id =  data.terraform_remote_state.engine.outputs.id
  kubeconfig_type = var.kubeconfig_type
}
