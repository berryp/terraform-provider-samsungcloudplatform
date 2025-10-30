data "samsungcloudplatform_transit_gateway_peerings" "tgw_peerings" {
}

output "contents" {
  value = data.samsungcloudplatform_transit_gateway_peerings.tgw_peerings.contents
}
