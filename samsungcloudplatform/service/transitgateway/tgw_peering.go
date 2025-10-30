package transitgateway

import (
	"context"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	tfTags "github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/service/tag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	samsungcloudplatform.RegisterResource("samsungcloudplatform_transit_gateway_peering", ResourceTransitGatewayPeering())
}

func ResourceTransitGatewayPeering() *schema.Resource {
	return &schema.Resource{
		//CRUD
		CreateContext: resourceTransitGatewayPeeringCreate,
		ReadContext:   resourceTransitGatewayPeeringRead,
		UpdateContext: resourceTransitGatewayPeeringUpdate,
		DeleteContext: resourceTransitGatewayPeeringDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"requester_transit_gateway_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Requester TGW ID",
				ValidateFunc: validation.StringLenBetween(3, 60),
			},
			"approver_transit_gateway_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Approver TGW ID",
				ValidateFunc: validation.StringLenBetween(3, 60),
			},
			"requester_project_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Requester Project ID",
				ValidateFunc: validation.StringLenBetween(3, 60),
			},
			"approver_project_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Approver Project ID",
				ValidateFunc: validation.StringLenBetween(3, 60),
			},
			"transit_gateway_peering_description": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Transit Gateway Peering Description",
				ValidateFunc: validation.StringLenBetween(0, 50),
			},
			"transit_gateway_peering_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Transit Gateway Peering State",
			},
			"tags": tfTags.TagsSchema(),
		},
		Description: "Provides a Transit Gateway Peering resource.",
	}
}

func resourceTransitGatewayPeeringCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	requesterTgwId := rd.Get("requester_transit_gateway_id").(string)
	approverTgwId := rd.Get("approver_transit_gateway_id").(string)
	requesterProjectId := rd.Get("requester_project_id").(string)
	approverProjectId := rd.Get("approver_project_id").(string)
	tgwPeeringDescription := rd.Get("transit_gateway_peering_description").(string)
	tags := rd.Get("tags").(map[string]interface{})

	inst := meta.(*client.Instance)

	response, _, err := inst.Client.TransitGateway.CreateTransitGatewayPeering(ctx, requesterTgwId, approverTgwId, requesterProjectId, approverProjectId, tgwPeeringDescription, tags)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(response.TransitGatewayPeeringId)
	return resourceTransitGatewayPeeringRead(ctx, rd, meta)

}
func resourceTransitGatewayPeeringRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	tgwInfo, _, err := inst.Client.TransitGateway.GetTransitGatewayPeeringInfo(ctx, rd.Id())
	if err != nil {
		rd.SetId("")
		if common.IsDeleted(err) {
			return nil
		}
		return diag.FromErr(err)
	}

	rd.Set("requester_transit_gateway_id", tgwInfo.RequesterTransitGatewayId)
	rd.Set("approver_transit_gateway_id", tgwInfo.ApproverTransitGatewayId)
	rd.Set("requester_project_id", tgwInfo.RequesterProjectId)
	rd.Set("approver_project_id", tgwInfo.ApproverProjectId)
	rd.Set("transit_gateway_peering_description", tgwInfo.TransitGatewayPeeringDescription)
	rd.Set("transit_gateway_peering_state", tgwInfo.TransitGatewayPeeringState)

	tfTags.SetTags(ctx, rd, meta, rd.Id())

	return nil
}

func resourceTransitGatewayPeeringUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	if rd.HasChanges("transit_gateway_peering_description") {
		_, _, err := inst.Client.TransitGateway.UpdateTransitGatewayPeeringDescription(ctx, rd.Id(), rd.Get("transit_gateway_peering_description").(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	err := tfTags.UpdateTags(ctx, rd, meta, rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceTransitGatewayPeeringRead(ctx, rd, meta)
}

func resourceTransitGatewayPeeringDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	state := rd.Get("transit_gateway_peering_state").(string)
	if state == "REQUESTING" {
		if _, err := inst.Client.TransitGateway.CancelTransitGatewayPeering(ctx, rd.Id()); err != nil {
			return diag.FromErr(err)
		}
	}

	err := inst.Client.TransitGateway.DeleteTransitGatewayPeering(ctx, rd.Id())
	if err != nil && !common.IsDeleted(err) {
		return diag.FromErr(err)
	}

	err = waitTransitGatewayPeering(ctx, inst.Client, rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func waitTransitGatewayPeering(ctx context.Context, scpClient *client.SCPClient, peeringId string) error {
	return client.WaitForStatus(ctx, scpClient, []string{"DELETING"}, []string{"DELETED"}, func() (interface{}, string, error) {
		return scpClient.TransitGateway.GetTransitGatewayPeeringForDelete(ctx, peeringId)
	})
}
