data "samsungcloudplatform_region" "region" {
  filter {
    name = "location"
    values = ["KR-WEST-1"]
  }
}

data "samsungcloudplatform_standard_image" "kafka_3_8_0_image" {
  service_group = "EXTENSION"
  service       = "Apache Kafka"
  region        = data.samsungcloudplatform_region.region.location

  filter {
    name   = "image_name"
    values = ["Apache Kafka 3.8.0"]
  }
}

resource "samsungcloudplatform_kafka" "demo_db" {
  kafka_cluster_name = "kafkacluster"

  image_id = data.samsungcloudplatform_standard_image.kafka_3_8_0_image.id
  service_zone_id = data.samsungcloudplatform_region.region.id

  subnet_id = "SUBNET-XXXXX"
  security_group_ids = ["FIREWALL_SECURITY_GROUP-XXXXX"]

  nat_enabled = false

  broker_sasl_account = "brokersasl"
  broker_sasl_password = "UserPa$$w0rd"
  broker_port = 9091

  zookeeper_sasl_account = "zookeepersasl"
  zookeeper_sasl_password = "UserPa$$w0rd"
  zookeeper_port = 2180

  broker_server_type = "kaS1v2m4"
  dynamic "broker_nodes" {
    for_each = var.broker_nodes
    content {
      broker_node_name = broker_nodes.value.broker_node_name
      nat_public_ip_id = broker_nodes.value.nat_public_ip_id
      availability_zone_name = broker_nodes.value.availability_zone_name
    }
  }
  broker_block_storages {
    block_storage_type = "SSD"
    block_storage_size = 11
  }

  zookeeper_server_type = "kaS1v2m4"
  dynamic "zookeeper_nodes" {
    for_each = var.zookeeper_nodes
    content {
      zookeeper_node_name = zookeeper_nodes.value.zookeeper_node_name
      nat_public_ip_id = zookeeper_nodes.value.nat_public_ip_id
    }
  }
  zookeeper_block_storages {
    block_storage_type = "SSD"
    block_storage_size = 12
  }

  akhq_enabled = true
  akhq_account = "akhqaccount"
  akhq_password = "UserPa$$w0rd"
  akhq_node {
    akhq_node_name = "akhqnode"
    nat_public_ip_id = null
    akhq_availability_zone_name = null
  }

  availability_zone_config {
    availability_zone_deployment_type = "DESIGNATED"
    availability_zone_name = "AZ1"
  }

  timezone = "Asia/Seoul"
}
