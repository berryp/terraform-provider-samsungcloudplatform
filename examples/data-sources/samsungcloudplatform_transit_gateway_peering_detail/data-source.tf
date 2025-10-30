data "samsungcloudplatform_transit_gateway_peerings" "tgw_peerings" {
}

data "samsungcloudplatform_transit_gateway_peering_detail" "tgw_peering" {
  transit_gateway_peering_id = data.samsungcloudplatform_transit_gateway_peerings.tgw_peerings.contents[0].transit_gateway_peering_id
}

output "detail" {
  value = data.samsungcloudplatform_transit_gateway_peering_detail.tgw_peering
}
