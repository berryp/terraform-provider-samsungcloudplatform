package transitgateway

import (
	"context"
	sdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
	transitgateway2 "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/transit-gateway2"
	"github.com/antihax/optional"
)

type Client struct {
	config    *sdk.Configuration
	sdkClient *transitgateway2.APIClient
}

func NewClient(config *sdk.Configuration) *Client {
	return &Client{
		config:    config,
		sdkClient: transitgateway2.NewAPIClient(config),
	}
}

// Transit Gateway --------------------------->

func (client *Client) GetTransitGatewayInfo(ctx context.Context, transitGatewayId string) (transitgateway2.TransitGatewayDetailResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayOpenApiControllerApi.DetailTransitGateway(ctx, client.config.ProjectId, transitGatewayId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) GetTransitGatewayList(ctx context.Context, request *transitgateway2.TransitGatewayOpenApiControllerApiListTransitGatewaysOpts) (transitgateway2.ListResponseTransitGatewayListItemResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayOpenApiControllerApi.ListTransitGateways(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) CreateTransitGateway(ctx context.Context, bandwidthGbps int32, serviceZoneId string, name string, uplinkEnabled bool, description string) (transitgateway2.AsyncResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayOpenApiV3ControllerApi.CreateTransitGateway1(ctx, client.config.ProjectId, transitgateway2.TransitGatewayCreateV3Request{
		BandwidthGbps:             bandwidthGbps,
		ServiceZoneId:             serviceZoneId,
		TransitGatewayName:        name,
		UplinkEnabled:             &uplinkEnabled,
		TransitGatewayDescription: description,
	})
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) UpdateTransitGatewayUplinkEnable(ctx context.Context, transitGatewayId string, uplinkEnabled bool) (transitgateway2.TransitGatewayDetailResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayOpenApiControllerApi.UpdateTransitGatewayUplinkEnabled(ctx, client.config.ProjectId, transitGatewayId, transitgateway2.TransitGatewayUplinkUpdateRequest{
		UplinkEnabled: &uplinkEnabled,
	})
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) DeleteTransitGateway(ctx context.Context, transitGatewayId string) error {
	_, _, err := client.sdkClient.TransitGatewayOpenApiControllerApi.DeleteTransitGateway(ctx, client.config.ProjectId, transitGatewayId)
	return err
}

// <-------------------------------

//  Transit Gateway - VPC Connections  -------------------------->

func (client *Client) GetTransitGatewayConnectionInfo(ctx context.Context, transitGatewayConnectionId string) (transitgateway2.TransitGatewayConnectionDetailResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.DetailTransitGatewayConnection(ctx, client.config.ProjectId, transitGatewayConnectionId)

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err

}

func (client *Client) GetTransitGatewayConnectionList(ctx context.Context, request *transitgateway2.TransitGatewayConnectionOpenApiControllerApiListTransitGatewayConnectionsOpts) (transitgateway2.ListResponseTransitGatewayConnectionListItemResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.ListTransitGatewayConnections(ctx, client.config.ProjectId, request)

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}

	return result, statusCode, err
}

func (client *Client) CreateTransitGatewayConnection(ctx context.Context, transitGatewayId string, vpcId string, requesterProjectId string, approverProjectId string, description string, firewallEnabled bool, firewallLoggable bool, connectionType string, tags map[string]interface{}) (transitgateway2.TransitGatewayConnectionApprovalResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.CreateTransitGatewayConnection(ctx, client.config.ProjectId, transitgateway2.TransitGatewayConnectionCreateRequest{
		RequesterProjectId:                  requesterProjectId,
		RequesterTransitGatewayId:           transitGatewayId,
		ApproverProjectId:                   approverProjectId,
		ApproverVpcId:                       vpcId,
		TransitGatewayConnectionDescription: description,
		FirewallEnabled:                     &firewallEnabled,
		FirewallLoggable:                    &firewallLoggable,
		TransitGatewayConnectionType:        connectionType,
		Tags:                                client.sdkClient.ToTagRequestList(tags),
	})

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) ApproveTransitGatewayConnection(ctx context.Context, transitGatewayConnectionId string) (transitgateway2.TransitGatewayConnectionApprovalResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.ApproveTransitGatewayConnection(ctx, client.config.ProjectId, transitGatewayConnectionId)

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) CancelTransitGatewayConnection(ctx context.Context, transitGatewayConnectionId string) (transitgateway2.TransitGatewayConnectionApprovalResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.CancelTransitGatewayConnection(ctx, client.config.ProjectId, transitGatewayConnectionId)

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) UpdateTransitGatewayConnectionDescription(ctx context.Context, transitGatewayConnectionId string, description string) (transitgateway2.TransitGatewayConnectionDetailResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.UpdateTransitGatewayConnectionDescription(ctx, client.config.ProjectId, transitGatewayConnectionId, transitgateway2.TransitGatewayConnectionDescriptionUpdateRequest{
		TransitGatewayConnectionDescription: description,
	})

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) DeleteTransitGatewayConnection(ctx context.Context, transitGatewayConnectionId string) error {

	_, _, err := client.sdkClient.TransitGatewayConnectionOpenApiControllerApi.DeleteTransitGatewayConnection(ctx, client.config.ProjectId, transitGatewayConnectionId)

	return err

}

// <-------------------------------

//  Transit Gateway Peering  -------------------------->

func (client *Client) GetTransitGatewayPeeringList(ctx context.Context, request *transitgateway2.TransitGatewayPeeringOpenApiControllerApiListTransitGatewayPeeringOpts) (transitgateway2.ListResponseTransitGatewayPeeringListItemResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.ListTransitGatewayPeering(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) GetTransitGatewayPeeringListByTransitGateway(ctx context.Context, transitGatewayId string) (transitgateway2.ListResponseTransitGatewayPeeringListItemResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.ListTransitGatewayPeeringByTransitGatewayId(ctx, client.config.ProjectId, transitGatewayId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) CreateTransitGatewayPeering(ctx context.Context, requesterTgwId string, approverTgwId string, requesterProjectId string, approverProjectId string, description string, tags map[string]interface{}) (transitgateway2.TransitGatewayPeeringResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.CreateTransitGatewayPeering(ctx, client.config.ProjectId, transitgateway2.TransitGatewayPeeringCreateRequest{
		RequesterProjectId:               requesterProjectId,
		RequesterTransitGatewayId:        requesterTgwId,
		ApproverProjectId:                approverProjectId,
		ApproverTransitGatewayId:         approverTgwId,
		TransitGatewayPeeringDescription: description,
		Tags:                             client.sdkClient.ToTagRequestList(tags),
	})

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) GetTransitGatewayPeeringInfo(ctx context.Context, transitGatewayPeeringId string) (transitgateway2.TransitGatewayPeeringDetailResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.DetailTransitGatewayPeering(ctx, client.config.ProjectId, transitGatewayPeeringId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) UpdateTransitGatewayPeeringDescription(ctx context.Context, transitGatewayPeeringId string, description string) (transitgateway2.TransitGatewayPeeringDetailResponse, int, error) {
	result, c, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.UpdateTransitGatewayPeeringDescription(ctx, client.config.ProjectId, transitGatewayPeeringId, transitgateway2.TransitGatewayPeeringDescriptionUpdateRequest{
		TransitGatewayPeeringDescription: description,
	})

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) DeleteTransitGatewayPeering(ctx context.Context, transitGatewayPeeringId string) error {
	_, _, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.DeleteTransitGatewayPeering(ctx, client.config.ProjectId, transitGatewayPeeringId)
	return err
}

func (client *Client) GetTransitGatewayPeeringForDelete(ctx context.Context, peeringId string) (transitgateway2.TransitGatewayPeeringListItemResponse, string, error) {
	result, _, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.ListTransitGatewayPeering(ctx, client.config.ProjectId, &transitgateway2.TransitGatewayPeeringOpenApiControllerApiListTransitGatewayPeeringOpts{
		Size: optional.NewInt32(1000),
		Page: optional.NewInt32(0),
	})
	if err != nil {
		return transitgateway2.TransitGatewayPeeringListItemResponse{}, "", err
	}
	for _, peeringInfo := range result.Contents {
		if peeringInfo.TransitGatewayPeeringId == peeringId {
			return peeringInfo, peeringInfo.TransitGatewayPeeringState, nil
		}
	}
	return transitgateway2.TransitGatewayPeeringListItemResponse{}, "DELETED", nil
}

func (client *Client) CancelTransitGatewayPeering(ctx context.Context, transitGatewayPeeringId string) (transitgateway2.TransitGatewayPeeringResponse, error) {
	result, _, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.CancelTransitGatewayPeering(ctx, client.config.ProjectId, transitGatewayPeeringId)
	return result, err
}

func (client *Client) ApproveTransitGatewayPeering(ctx context.Context, transitGatewayPeeringId string) (transitgateway2.TransitGatewayPeeringResponse, error) {
	result, _, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.ApproveTransitGatewayPeering(ctx, client.config.ProjectId, transitGatewayPeeringId)
	return result, err
}

func (client *Client) RejectTransitGatewayPeering(ctx context.Context, transitGatewayPeeringId string) (transitgateway2.TransitGatewayPeeringResponse, error) {
	result, _, err := client.sdkClient.TransitGatewayPeeringOpenApiControllerApi.RejectTransitGatewayPeering(ctx, client.config.ProjectId, transitGatewayPeeringId)
	return result, err
}
