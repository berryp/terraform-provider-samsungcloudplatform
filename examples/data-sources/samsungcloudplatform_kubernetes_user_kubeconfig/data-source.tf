data "samsungcloudplatform_kubernetes_user_kubeconfig" "kube_config" {
  kubernetes_engine_id = "HSCLUSTER-XXXXXXXXXXXXX"
  kubeconfig_type = "private"
}

output "output_scp_kubernetes_user_kubeconfig" {
  value = data.samsungcloudplatform_kubernetes_user_kubeconfig.kube_config.kube_config
}
