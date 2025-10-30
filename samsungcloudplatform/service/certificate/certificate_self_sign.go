package certificate

import (
	"context"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client/certificate"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	tfTags "github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/service/tag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	samsungcloudplatform.RegisterResource("samsungcloudplatform_certificate_self_sign", CertificateSelfSign())
}

func CertificateSelfSign() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertificateSelfSignCreate,
		ReadContext:   resourceCertificateSelfSignRead,
		UpdateContext: resourceCertificateSelfSignUpdate,
		DeleteContext: resourceCertificateSelfSignDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Certificate Name",
			},
			"common_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Common Name",
			},
			"organization_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Organization Name",
			},
			"start_date": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Certificate Start Date",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Certificate Expiration Date",
			},
			"recipients": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "recipients",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tags": tfTags.TagsSchema(),
		},
		Description: "Provides a Certificate resource.",
	}
}

func resourceCertificateSelfSignRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	inst := meta.(*client.Instance)

	dcInfo, _, err := inst.Client.Certificate.GetCertificateInfo(ctx, rd.Id())
	if err != nil {
		rd.SetId("")
		if common.IsDeleted(err) {
			return nil
		}

		return diag.FromErr(err)
	}
	rd.Set("name", dcInfo.CertificateName)

	return nil
}

func resourceCertificateSelfSignUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("Update function is not implemented")
}

func resourceCertificateSelfSignCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// Get values from schema
	certificateName := rd.Get("name").(string)
	commonName := rd.Get("common_name").(string)
	organizationName := rd.Get("organization_name").(string)
	startDate := rd.Get("start_date").(string)
	expirationDate := rd.Get("expiration_date").(string)
	recipients, err := makeRecipients(rd)
	tags := rd.Get("tags").(map[string]interface{})

	inst := meta.(*client.Instance)

	response, _, err := inst.Client.Certificate.CreateSelfSignCertificate(ctx, certificate.CreateDevCertificateRequest{
		CertificateExpirationDate: expirationDate,
		CertificateName:           certificateName,
		CertificateStartDate:      startDate,
		CommonName:                commonName,
		OrganizationName:          organizationName,
		Recipients:                recipients,
		Tags:                      tags,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	rd.SetId(response.CertificateId)

	// Refresh
	return resourceCertificateRead(ctx, rd, meta)
}

func resourceCertificateSelfSignDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	_, _, err := inst.Client.Certificate.DeleteCertificate(ctx, rd.Id())
	if err != nil && !common.IsDeleted(err) {
		return diag.FromErr(err)
	}

	err = waitForCertificateStatus(ctx, inst.Client, rd.Id(), []string{}, []string{"DELETED", "FREE"}, false)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
