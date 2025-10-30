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
	samsungcloudplatform.RegisterResource("samsungcloudplatform_certificate", Certificate())
}

func Certificate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertificateCreate,
		ReadContext:   resourceCertificateRead,
		UpdateContext: resourceCertificateUpdate,
		DeleteContext: resourceCertificateDelete,
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
			"chain": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Certificate Chain",
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Private Key",
			},
			"body": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Public Certificate",
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

func makeRecipients(rd *schema.ResourceData) ([]certificate.Recipient, error) {
	recipientsList := rd.Get("recipients").([]interface{})
	recipients := make([]certificate.Recipient, len(recipientsList))
	for i, recipient := range recipientsList {
		if recipient == nil {
			continue
		}
		r := recipient.(map[string]interface{})
		if t, ok := r["user_id"]; ok {
			recipients[i].UserId = t.(string)
		}
		if t, ok := r["user_name"]; ok {
			recipients[i].UserName = t.(string)
		}
		if t, ok := r["email"]; ok {
			recipients[i].Email = t.(string)
		}
	}
	return recipients, nil
}

func resourceCertificateRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

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

func resourceCertificateUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("Update function is not implemented")
}

func resourceCertificateCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {

	// Get values from schema
	certificateName := rd.Get("name").(string)
	certificateKey := rd.Get("key").(string)
	certificateBody := rd.Get("body").(string)
	CertificateChain := rd.Get("chain").(string)
	recipients, err := makeRecipients(rd)
	tags := rd.Get("tags").(map[string]interface{})

	inst := meta.(*client.Instance)

	response, _, err := inst.Client.Certificate.CreateCertificate(ctx, certificate.CreatePrdCertificateRequest{
		CertificateChain:  CertificateChain,
		CertificateName:   certificateName,
		PrivateKey:        certificateKey,
		PublicCertificate: certificateBody,
		Recipients:        recipients,
		Tags:              tags,
	})
	if err != nil {
		return diag.FromErr(err)
	}
	rd.SetId(response.CertificateId)

	// Refresh
	return resourceCertificateRead(ctx, rd, meta)
}

func resourceCertificateDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func waitForCertificateStatus(ctx context.Context, scpClient *client.SCPClient, CertificateId string, pendingStates []string, targetStates []string, errorOnNotFound bool) error {
	return client.WaitForStatus(ctx, scpClient, pendingStates, targetStates, func() (interface{}, string, error) {
		info, c, err := scpClient.Certificate.GetCertificateInfo(ctx, CertificateId)
		if err != nil {
			if c == 404 && !errorOnNotFound {
				return "", "DELETED", nil
			}
			if c == 403 && !errorOnNotFound {
				return "", "DELETED", nil
			}
			return nil, "", err
		}
		return info, info.CertificateState, nil
	})
}
