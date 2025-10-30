resource "samsungcloudplatform_transit_gateway_peering" "tgw_peering" {
  requester_transit_gateway_id  = var.requester_transit_gateway_id
  approver_transit_gateway_id   = var.approver_transit_gateway_id
  requester_project_id          = var.requester_project_id
  approver_project_id           = var.approver_project_id
  transit_gateway_peering_description = var.transit_gateway_peering_description
}
