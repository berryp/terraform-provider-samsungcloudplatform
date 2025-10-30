package transitgateway

import (
	"context"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	samsungcloudplatform.RegisterResource("samsungcloudplatform_transit_gateway_peering_approve", ResourceTransitGatewayPeeringApprove())
	samsungcloudplatform.RegisterResource("samsungcloudplatform_transit_gateway_peering_reject", ResourceTransitGatewayPeeringReject())
	samsungcloudplatform.RegisterResource("samsungcloudplatform_transit_gateway_peering_cancel", ResourceTransitGatewayPeeringCancel())
}

func ResourceTransitGatewayPeeringApprove() *schema.Resource {
	return &schema.Resource{
		//CRUD
		CreateContext: resourceTransitGatewayPeeringApproveCreate,
		ReadContext:   resourceTransitGatewayPeeringActionRead,
		UpdateContext: resourceTransitGatewayPeeringApproveCreate,
		DeleteContext: resourceTransitGatewayPeeringActionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("TransitGatewayPeeringId"):    {Type: schema.TypeString, Required: true, Description: "Transit Gateway Peering Id"},
			common.ToSnakeCase("TransitGatewayPeeringState"): {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering State"},
		},
		Description: "Approve TGW Peering Request.",
	}
}

func resourceTransitGatewayPeeringApproveCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	transitGatewayPeeringId := rd.Get(common.ToSnakeCase("TransitGatewayPeeringId")).(string)

	inst := meta.(*client.Instance)

	result, err := inst.Client.TransitGateway.ApproveTransitGatewayPeering(ctx, transitGatewayPeeringId)
	if err != nil {
		return diag.FromErr(err)
	}
	if !(*result.Success) {
		// check when false //
	}

	err = waitTransitGatewayPeeringCreating(ctx, inst.Client, transitGatewayPeeringId)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(result.TransitGatewayPeeringId)
	return resourceTransitGatewayPeeringActionRead(ctx, rd, meta)
}

func ResourceTransitGatewayPeeringReject() *schema.Resource {
	return &schema.Resource{
		//CRUD
		CreateContext: resourceTransitGatewayPeeringRejectCreate,
		ReadContext:   resourceTransitGatewayPeeringActionRead,
		UpdateContext: resourceTransitGatewayPeeringActionUpdate,
		DeleteContext: resourceTransitGatewayPeeringActionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("TransitGatewayPeeringId"):    {Type: schema.TypeString, Required: true, Description: "Transit Gateway Peering Id"},
			common.ToSnakeCase("TransitGatewayPeeringState"): {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Id"},
		},
		Description: "Reject Peering Request.",
	}
}

func resourceTransitGatewayPeeringRejectCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	transitGatewayPeeringId := rd.Get(common.ToSnakeCase("TransitGatewayPeeringId")).(string)

	inst := meta.(*client.Instance)

	result, err := inst.Client.TransitGateway.RejectTransitGatewayPeering(ctx, transitGatewayPeeringId)
	if err != nil {
		return diag.FromErr(err)
	}

	if !(*result.Success) {
		// check when false //
	}
	rd.SetId(result.TransitGatewayPeeringId)
	return resourceTransitGatewayPeeringActionRead(ctx, rd, meta)
}

func ResourceTransitGatewayPeeringCancel() *schema.Resource {
	return &schema.Resource{
		//CRUD
		CreateContext: resourceTransitGatewayPeeringCancelCreate,
		ReadContext:   resourceTransitGatewayPeeringActionRead,
		UpdateContext: resourceTransitGatewayPeeringActionUpdate,
		DeleteContext: resourceTransitGatewayPeeringActionDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("TransitGatewayPeeringId"):    {Type: schema.TypeString, Required: true, Description: "Transit Gateway Peering Id"},
			common.ToSnakeCase("TransitGatewayPeeringState"): {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Id"},
		},
		Description: "Cancel Peering Request.",
	}
}

func resourceTransitGatewayPeeringCancelCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	transitGatewayPeeringId := rd.Get(common.ToSnakeCase("TransitGatewayPeeringId")).(string)

	inst := meta.(*client.Instance)

	result, err := inst.Client.TransitGateway.CancelTransitGatewayPeering(ctx, transitGatewayPeeringId)
	if err != nil {
		return diag.FromErr(err)
	}

	if !(*result.Success) {
		// check when false //
	}
	rd.SetId(result.TransitGatewayPeeringId)
	return resourceTransitGatewayPeeringActionRead(ctx, rd, meta)
}

func resourceTransitGatewayPeeringActionRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	peeringInfo, _, err := inst.Client.TransitGateway.GetTransitGatewayPeeringInfo(ctx, rd.Id())
	if err != nil {
		rd.SetId("")
		if common.IsDeleted(err) {
			return nil
		}
		return diag.FromErr(err)
	}
	rd.Set(common.ToSnakeCase("TransitGatewayPeeringState"), peeringInfo.TransitGatewayPeeringState)

	return nil
}

func resourceTransitGatewayPeeringActionUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("Update function is not supported!")
}

func resourceTransitGatewayPeeringActionDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rd.SetId("")
	return nil
}

func waitTransitGatewayPeeringCreating(ctx context.Context, scpClient *client.SCPClient, peeringId string) error {
	return client.WaitForStatus(ctx, scpClient, []string{"CREATING"}, []string{"ACTIVE"}, func() (interface{}, string, error) {
		info, _, err := scpClient.TransitGateway.GetTransitGatewayPeeringInfo(ctx, peeringId)
		if err != nil {
			return nil, "", err
		}
		return info, info.TransitGatewayPeeringState, nil
	})
}
