variable "broker_nodes" {
  type = list(object({
    broker_node_name = string
    nat_public_ip_id = string
    availability_zone_name = string
  }))
  default = [
    {
      broker_node_name = "brokernode-01"
      nat_public_ip_id = null
      availability_zone_name = null
    },
    {
      broker_node_name = "brokernode-02"
      nat_public_ip_id = null
      availability_zone_name = null
    },
    {
      broker_node_name = "brokernode-03"
      nat_public_ip_id = null
      availability_zone_name = null
    },
  ]
}

variable "zookeeper_nodes" {
  type = list(object({
    zookeeper_node_name = string
    nat_public_ip_id = string
    availability_zone_name = string
  }))
  default = [
    {
      zookeeper_node_name = "zookeepernode-01"
      nat_public_ip_id = null
      availability_zone_name = null
    },
    {
      zookeeper_node_name = "zookeepernode-02"
      nat_public_ip_id = null
      availability_zone_name = null
    },
    {
      zookeeper_node_name = "zookeepernode-03"
      nat_public_ip_id = null
      availability_zone_name = null
    },
  ]
}
