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
	samsungcloudplatform.RegisterDataSource("Kubernetes", "samsungcloudplatform_kubernetes_user_kubeconfig", DatasourceUserKubeConfig())
}

func DatasourceUserKubeConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: engineUserKubeConfig, //데이터 조회 함수
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("KubernetesEngineId"): {Type: schema.TypeString, Required: true, Description: "Engine Id"},
			common.ToSnakeCase("kubeconfigType"):     {Type: schema.TypeString, Required: true, Description: "kubeconfig Type"},
			common.ToSnakeCase("KubeConfig"):         {Type: schema.TypeString, Computed: true, Description: "KubeConfig"},
		},
		Description: "Provides Kubernetes Engine Detail",
	}
}

func engineUserKubeConfig(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	engineId := rd.Get("kubernetes_engine_id").(string)
	kubeconfigType := rd.Get("kubeconfig_type").(string)

	response, _, err := inst.Client.KubernetesEngine.GetUserKubeConfig(ctx, engineId, kubeconfigType)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(uuid.NewV4().String())

	rd.Set(common.ToSnakeCase("KubeConfig"), response)

	return nil
}
