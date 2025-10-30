package configinspection

import (
	"context"

	sdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
	configinspection "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/config-inspection"
)

type Client struct {
	config    *sdk.Configuration
	sdkClient *configinspection.APIClient
}

func NewClient(config *sdk.Configuration) *Client {
	return &Client{
		config:    config,
		sdkClient: configinspection.NewAPIClient(config),
	}
}

func (client *Client) GetConfigInspectionProductList(ctx context.Context, request *configinspection.ConfigInspectionOpenApiControllerApiConfigInspectionListOpts) (configinspection.ConfigInspectionListResponse, int, error) {
	result, c, err := client.sdkClient.ConfigInspectionOpenApiControllerApi.ConfigInspectionList(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) GetConfigInspectionProductDetail(ctx context.Context, diagnosisId string) (configinspection.DiagnosisObjectDetailResponse, int, error) {
	result, c, err := client.sdkClient.ConfigInspectionOpenApiControllerApi.GetDiagnosisObjectDetail(ctx, client.config.ProjectId, diagnosisId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) ConfigInspectionProductCreate(ctx context.Context, diagnosisObjectListRequest configinspection.DiagnosisObjectListRequest, tags map[string]interface{}) (configinspection.DiagnosisCreateResponse, int, error) {

	diagnosisObjectListRequest.Tags = client.sdkClient.ToTagRequestList(tags)
	result, c, err := client.sdkClient.ConfigInspectionOpenApiControllerApi.CreateDiagnosisObject(ctx, client.config.ProjectId, diagnosisObjectListRequest)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) ConfigInspectionProductDelete(ctx context.Context, diagnosisId string) (configinspection.TerminateResponse, error) {
	result, _, err := client.sdkClient.ConfigInspectionOpenApiControllerApi.TerminateDiagnosisObject(ctx, client.config.ProjectId, diagnosisId)

	return result, err
}

func (client *Client) GetConfigInspectionList(ctx context.Context, request *configinspection.ConfigInspectionDashboardOpenApiControllerApiGetDiagnosisResultListOpts) (configinspection.PageResponseDiagnosisResultResponse, int, error) {
	result, c, err := client.sdkClient.ConfigInspectionDashboardOpenApiControllerApi.GetDiagnosisResultList(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) GetConfigInspectionDetail(ctx context.Context, DiagnosisId string, DiagnosisRequestSequence string) (configinspection.DiagnosisResultDetailResponse, int, error) {
	result, c, err := client.sdkClient.ConfigInspectionDashboardOpenApiControllerApi.GetDiagnosisResultDetail(ctx, client.config.ProjectId, DiagnosisId, DiagnosisRequestSequence)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) RequestNewConfigInspection(ctx context.Context, request configinspection.DiagnosisRequest) (configinspection.CheckResponse, int, error) {
	result, c, err := client.sdkClient.ConfigInspectionOpenApiControllerApi.DiagnosisRequest(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}
