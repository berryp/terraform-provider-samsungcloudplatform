package certificate

import (
	"context"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	"github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/certificate"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_certificate", DatasourceCertificate())
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_certificates", DatasourceCertificates())
}

func DatasourceCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceCertificateRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("CertificateId"):             {Type: schema.TypeString, Optional: true, Description: "Certificate id"},
			common.ToSnakeCase("CertificateName"):           {Type: schema.TypeString, Optional: true, Description: "Certificate Name"},
			common.ToSnakeCase("CommonName"):                {Type: schema.TypeString, Optional: true, Description: "Common Name"},
			common.ToSnakeCase("PurposeType"):               {Type: schema.TypeString, Optional: true, Description: "Purpose Type"},
			common.ToSnakeCase("CertificateStartDate"):      {Type: schema.TypeString, Optional: true, Description: "Start Date"},
			common.ToSnakeCase("CertificateExpirationDate"): {Type: schema.TypeString, Optional: true, Description: "Expiration Date"},
			common.ToSnakeCase("CreatedDt"):                 {Type: schema.TypeString, Optional: true, Description: "Creation date"},
			common.ToSnakeCase("CreatedBy"):                 {Type: schema.TypeString, Optional: true, Description: "Created By"},
			common.ToSnakeCase("ModifiedDt"):                {Type: schema.TypeString, Optional: true, Description: "Modified date"},
			common.ToSnakeCase("ModifiedBy"):                {Type: schema.TypeString, Optional: true, Description: "Modified By"},
			common.ToSnakeCase("CertificateState"):          {Type: schema.TypeString, Optional: true, Description: "Certificate State"},
			common.ToSnakeCase("PrivateKey"):                {Type: schema.TypeString, Optional: true, Description: "Certificate Private Key"},
			common.ToSnakeCase("PublicCertificate"):         {Type: schema.TypeString, Optional: true, Description: "Certificate Public(Body)"},
			common.ToSnakeCase("CertificateChain"):          {Type: schema.TypeString, Optional: true, Description: "Certificate Chain"},
			common.ToSnakeCase("CertificateType"):           {Type: schema.TypeString, Optional: true, Description: "Certificate Type"},
			common.ToSnakeCase("CertificateVersion"):        {Type: schema.TypeString, Optional: true, Description: "Certificate Version"},
			common.ToSnakeCase("CreatedByName"):             {Type: schema.TypeString, Optional: true, Description: "Certificate Created By Name"},
			common.ToSnakeCase("CreatedByEmail"):            {Type: schema.TypeString, Optional: true, Description: "Certificate Created By Email"},
			common.ToSnakeCase("OrganizationName"):          {Type: schema.TypeString, Optional: true, Description: "Organization Name"},
			common.ToSnakeCase("ProjectId"):                 {Type: schema.TypeString, Optional: true, Description: "Project Id"},
			common.ToSnakeCase("KeyBitSize"):                {Type: schema.TypeInt, Optional: true, Description: "Key Bit Size"},
			common.ToSnakeCase("UsedResourceCount"):         {Type: schema.TypeInt, Optional: true, Description: "Used Resource Count"},
		},
		Description: "Provides a Certificate",
	}
}

func datasourceCertificateRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	certificateId := rd.Get(common.ToSnakeCase("CertificateId")).(string)

	info, _, err := inst.Client.Certificate.GetCertificateInfo(ctx, certificateId)

	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.SetId(uuid.NewV4().String())
	rd.Set(common.ToSnakeCase("CertificateId"), info.CertificateId)
	rd.Set(common.ToSnakeCase("CertificateName"), info.CertificateName)
	rd.Set(common.ToSnakeCase("CommonName"), info.CommonName)
	rd.Set(common.ToSnakeCase("PurposeType"), info.PurposeType)
	rd.Set(common.ToSnakeCase("CertificateStartDate"), info.CertificateStartDate.Format(time.RFC3339))
	rd.Set(common.ToSnakeCase("CertificateExpirationDate"), info.CertificateExpirationDate.Format(time.RFC3339))
	rd.Set(common.ToSnakeCase("CreatedDt"), info.CreatedDt.Format(time.RFC3339))
	rd.Set(common.ToSnakeCase("CreatedBy"), info.CreatedBy)
	rd.Set(common.ToSnakeCase("ModifiedDt"), info.ModifiedDt.Format(time.RFC3339))
	rd.Set(common.ToSnakeCase("ModifiedBy"), info.ModifiedBy)
	rd.Set(common.ToSnakeCase("CertificateState"), info.CertificateState)
	rd.Set(common.ToSnakeCase("PrivateKey"), info.PrivateKey)
	rd.Set(common.ToSnakeCase("PublicCertificate"), info.PublicCertificate)
	rd.Set(common.ToSnakeCase("CertificateChain"), info.CertificateChain)
	rd.Set(common.ToSnakeCase("CertificateType"), info.CertificateType)
	rd.Set(common.ToSnakeCase("CertificateVersion"), info.CertificateVersion)
	rd.Set(common.ToSnakeCase("CreatedByName"), info.CreatedByName)
	rd.Set(common.ToSnakeCase("CreatedByEmail"), info.CreatedByEmail)
	rd.Set(common.ToSnakeCase("OrganizationName"), info.OrganizationName)
	rd.Set(common.ToSnakeCase("ProjectId"), info.ProjectId)
	rd.Set(common.ToSnakeCase("KeyBitSize"), info.KeyBitSize)
	rd.Set(common.ToSnakeCase("UsedResourceCount"), info.UsedResourceCount)

	return nil

}

func DatasourceCertificates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCertificateList, //데이터 조회 함수
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{ //스키마 정의
			common.ToSnakeCase("CertificateName"): {Type: schema.TypeString, Computed: true, Description: "Certificate Name"},
			common.ToSnakeCase("CommonName"):      {Type: schema.TypeString, Computed: true, Description: "Common Name"},
			"page":                                {Type: schema.TypeInt, Optional: true, Default: 0, Description: "Page start number from which to get the list"},
			"size":                                {Type: schema.TypeInt, Optional: true, Default: 20, Description: "Size to get list"},
			"sort":                                {Type: schema.TypeString, Optional: true, Description: "Sort"},
			"contents":                            {Type: schema.TypeList, Computed: true, Description: "Certificate list", Elem: DatasourceCertificateElem()},
			"total_count":                         {Type: schema.TypeInt, Computed: true, Description: "Total list size"},
		},
		Description: "Provides Certificate List",
	}
}

func DatasourceCertificateElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("CertificateId"):             {Type: schema.TypeString, Computed: true, Description: "Certificate Id"},
			common.ToSnakeCase("CertificateName"):           {Type: schema.TypeString, Computed: true, Description: "Certificate Name"},
			common.ToSnakeCase("CommonName"):                {Type: schema.TypeString, Computed: true, Description: "Common Name"},
			common.ToSnakeCase("PurposeType"):               {Type: schema.TypeString, Computed: true, Description: "Purpose Type"},
			common.ToSnakeCase("CertificateStartDate"):      {Type: schema.TypeString, Computed: true, Description: "Certificate start Date"},
			common.ToSnakeCase("CertificateExpirationDate"): {Type: schema.TypeString, Computed: true, Description: "Certificate Expiration Date"},
			common.ToSnakeCase("CreatedDt"):                 {Type: schema.TypeString, Computed: true, Description: "Creation date"},
			common.ToSnakeCase("CreatedBy"):                 {Type: schema.TypeString, Computed: true, Description: "Created By"},
			common.ToSnakeCase("ModifiedDt"):                {Type: schema.TypeString, Computed: true, Description: "Modified date"},
			common.ToSnakeCase("ModifiedBy"):                {Type: schema.TypeString, Computed: true, Description: "Modified By"},
			common.ToSnakeCase("CertificateState"):          {Type: schema.TypeString, Computed: true, Description: "Certificate State"},
		},
	}
}
func dataSourceCertificateList(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	responses, _, err := inst.Client.Certificate.GetCertificatesList(ctx, &certificate.CertControlApiListCertificatesOpts{
		CertificateName: common.GetKeyString(rd, common.ToSnakeCase("CertificateName")),
		CommonName:      common.GetKeyString(rd, common.ToSnakeCase("CommonName")),
		Page:            optional.NewInt32(0),
		Size:            optional.NewInt32(10000),
		Sort:            optional.NewInterface([]string{"createdDt:desc"}),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	contents := common.ConvertStructToMaps(responses.Contents)

	rd.SetId(uuid.NewV4().String())
	rd.Set("contents", contents)
	rd.Set("total_count", responses.TotalCount)

	return nil
}
