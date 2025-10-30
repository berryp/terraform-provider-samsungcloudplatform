package kafka

import (
	"context"
	sdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
	"github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/kafka"
)

type Client struct {
	config    *sdk.Configuration
	sdkClient *kafka.APIClient
}

func NewClient(config *sdk.Configuration) *Client {
	return &Client{
		config:    config,
		sdkClient: kafka.NewAPIClient(config),
	}
}

func (client *Client) DetailKafkaCluster(ctx context.Context, kafkaClusterId string) (kafka.KafkaClusterDetailResponse, int, error) {
	result, c, err := client.sdkClient.KafkaSearchApi.DetailKafkaCluster(ctx, client.config.ProjectId, kafkaClusterId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) ListKafkaClusters(ctx context.Context, request *kafka.KafkaSearchApiListKafkaClustersOpts) (kafka.ListResponseKafkaClusterListItemResponse, int, error) {
	result, c, err := client.sdkClient.KafkaSearchApi.ListKafkaClusters(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) CreateKafkaCluster(ctx context.Context, request kafka.KafkaClusterCreateRequest, tags map[string]interface{}) (kafka.AsyncResponse, int, error) {
	request.Tags = client.sdkClient.ToTagRequestList(tags)
	result, c, err := client.sdkClient.KafkaProvisioningApi.CreateKafkaCluster(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) DeleteKafkaCluster(ctx context.Context, kafkaClusterId string) (kafka.AsyncResponse, int, error) {
	result, c, err := client.sdkClient.KafkaOperationManagementApi.DeleteKafkaCluster(ctx, client.config.ProjectId, kafkaClusterId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}
