data "samsungcloudplatform_transit_gateway_peerings" "tgw_peerings" {
}

data "samsungcloudplatform_tgw_peerings_by_tgw" "tgw_peerings_by_tgw" {
  transit_gateway_id = data.samsungcloudplatform_transit_gateway_peerings.tgw_peerings.contents[0].requester_transit_gateway_id
}

output "contents" {
  value = data.samsungcloudplatform_tgw_peerings_by_tgw.tgw_peerings_by_tgw.contents
}
