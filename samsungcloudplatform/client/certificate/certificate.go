package certificate

import (
	"context"
	sdk "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/client"
	"github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/certificate"
)

type Client struct {
	config    *sdk.Configuration
	sdkClient *certificate.APIClient
}

func NewClient(config *sdk.Configuration) *Client {
	return &Client{
		config:    config,
		sdkClient: certificate.NewAPIClient(config),
	}
}

func (client *Client) GetCertificatesList(ctx context.Context, request *certificate.CertControlApiListCertificatesOpts) (certificate.PageResponseV2CertificateListResponse, int, error) {
	result, c, err := client.sdkClient.CertControlApi.ListCertificates(ctx, client.config.ProjectId, request)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) GetCertificateInfo(ctx context.Context, CertificateId string) (certificate.CertificateResponse, int, error) {
	result, c, err := client.sdkClient.CertControlApi.DetailCertificate(ctx, client.config.ProjectId, CertificateId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) CreateCertificate(ctx context.Context, request CreatePrdCertificateRequest) (certificate.CertificateResponse, int, error) {

	recipients := make([]certificate.Recipient, 0)
	for _, recipient := range request.Recipients {
		recipients = append(recipients, certificate.Recipient{
			Email:    recipient.Email,
			UserId:   recipient.UserId,
			UserName: recipient.UserName,
		})
	}

	result, c, err := client.sdkClient.CertControlApi.CreateCertificate(ctx, client.config.ProjectId, certificate.CreatePrdCertificateRequest{
		CertificateChain:  request.CertificateChain,
		CertificateName:   request.CertificateName,
		PrivateKey:        request.PrivateKey,
		PublicCertificate: request.PublicCertificate,
		Recipients:        recipients,
		Tags:              client.sdkClient.ToTagRequestList(request.Tags),
	}, nil)

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) CreateSelfSignCertificate(ctx context.Context, request CreateDevCertificateRequest) (certificate.CertificateResponse, int, error) {

	recipients := make([]certificate.Recipient, 0)
	for _, recipient := range request.Recipients {
		recipients = append(recipients, certificate.Recipient{
			Email:    recipient.Email,
			UserId:   recipient.UserId,
			UserName: recipient.UserName,
		})
	}

	result, c, err := client.sdkClient.CertControlApi.CreateSelfSignedCertificate(ctx, client.config.ProjectId, certificate.CreateDevCertificateRequest{
		CertificateExpirationDate: request.CertificateExpirationDate,
		CertificateName:           request.CertificateName,
		CertificateStartDate:      request.CertificateStartDate,
		CommonName:                request.CommonName,
		OrganizationName:          request.OrganizationName,
		Recipients:                recipients,
		Tags:                      client.sdkClient.ToTagRequestList(request.Tags),
	}, nil)

	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}

func (client *Client) DeleteCertificate(ctx context.Context, CertificateId string) (certificate.StatusResponse, int, error) {
	result, c, err := client.sdkClient.CertControlApi.DeleteCertificate(ctx, client.config.ProjectId, CertificateId)
	var statusCode int
	if c != nil {
		statusCode = c.StatusCode
	}
	return result, statusCode, err
}
