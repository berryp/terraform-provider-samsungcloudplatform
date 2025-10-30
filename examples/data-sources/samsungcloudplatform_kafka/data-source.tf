data "samsungcloudplatform_kafka" "my_scp_kafka" {
  kafka_cluster_id = "SERVICE-XXXXX"
}

output "output_my_scp_kafka" {
  value = data.samsungcloudplatform_kafka.my_scp_kafka
}
