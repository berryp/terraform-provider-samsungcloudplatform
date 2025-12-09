package kubernetes

import (
	"context"

	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uuid "github.com/satori/go.uuid"
)

func init() {
	samsungcloudplatform.RegisterResource("Kubernetes", "samsungcloudplatform_kubernetes_admin_kubeconfig", ResourceKubernetesAdminKubeconfig())
}

func ResourceKubernetesAdminKubeconfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAdminKubeconfigCreate, //데이터 조회 함수
		ReadContext:   resourceAdminKubeconfigRead,   //실제로는 조회하지 않음 -> state 확인 가능
		DeleteContext: resourceAdminKubeconfigDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("KubernetesEngineId"): {Type: schema.TypeString, Required: true, ForceNew: true, Description: "Engine Id"},
			common.ToSnakeCase("kubeconfigType"):     {Type: schema.TypeString, Required: true, ForceNew: true, Description: "kubeconfig Type"},
			common.ToSnakeCase("KubeConfig"):         {Type: schema.TypeString, Computed: true, ForceNew: true, Description: "Administrator KubeConfig"},
		},
		Description: "Provides Kubernetes Engine Administrator Kubeconfig",
	}
}

func resourceAdminKubeconfigCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	engineId := data.Get("kubernetes_engine_id").(string)
	kubeconfigType := data.Get("kubeconfig_type").(string)

	response, _, err := inst.Client.KubernetesEngine.GetKubeConfig(ctx, engineId, kubeconfigType)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(uuid.NewV4().String())
	data.Set(common.ToSnakeCase("KubeConfig"), response)

	return nil
}

func resourceAdminKubeconfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// state 에 이미 저장된 값만 반환 → 아무 작업도 하지 않음
	// (필요하면 여기서 캐시가 없을 때만 재호출하도록 로직을 넣을 수 있음)
	return nil
}

func resourceAdminKubeconfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// SCP 에서는 kubeconfig 삭제 API 가 없으므로, Terraform state만 삭제
	d.SetId("")
	return nil
}
