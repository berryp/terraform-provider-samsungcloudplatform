package transitgateway

import (
	"context"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	transitgateway2 "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/transit-gateway2"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_transit_gateway_peerings", DataSourceTransitGatewayPeerings())
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_tgw_peerings_by_tgw", DataSourceTgwPeeringsByTgw())
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_transit_gateway_peering_detail", DataSourceTransitGatewayPeeringDetail())
}

func DataSourceTransitGatewayPeerings() *schema.Resource {
	return &schema.Resource{
		ReadContext: tgwPeeringdataSourceList,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("ApproverTransitGatewayId"):  {Type: schema.TypeString, Optional: true, Description: "Approver Transit Gateway ID"},
			common.ToSnakeCase("RequesterTransitGatewayId"): {Type: schema.TypeString, Optional: true, Description: "Requester Transit Gateway ID"},
			common.ToSnakeCase("TransitGatewayPeeringName"): {Type: schema.TypeString, Optional: true, Description: "Transit Gateway Peering Name"},
			common.ToSnakeCase("CreatedBy"):                 {Type: schema.TypeString, Optional: true, Description: "User ID who create the resources"},
			common.ToSnakeCase("Page"):                      {Type: schema.TypeInt, Optional: true, Default: 0, Description: "Page number"},
			common.ToSnakeCase("Size"):                      {Type: schema.TypeInt, Optional: true, Default: 20, Description: "List size per a page"},
			"contents":                                      {Type: schema.TypeList, Optional: true, Description: "List of TGW Peerings", Elem: tgwPeeringDataSourceElem()},
			"total_count":                                   {Type: schema.TypeInt, Computed: true, Description: "Total Count of TGW Peerings"},
		},
		Description: "provides Lists of TGW Peering",
	}
}

func tgwPeeringdataSourceList(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	responses, _, err := inst.Client.TransitGateway.GetTransitGatewayPeeringList(ctx, &transitgateway2.TransitGatewayPeeringOpenApiControllerApiListTransitGatewayPeeringOpts{
		ApproverTransitGatewayId:  optional.NewString(rd.Get("approver_transit_gateway_id").(string)),
		RequesterTransitGatewayId: optional.NewString(rd.Get("requester_transit_gateway_id").(string)),
		TransitGatewayPeeringName: optional.NewString(rd.Get("transit_gateway_peering_name").(string)),
		CreatedBy:                 optional.NewString(rd.Get("created_by").(string)),
		Page:                      optional.NewInt32((int32)(rd.Get("page").(int))),
		Size:                      optional.NewInt32((int32)(rd.Get("size").(int))),
		Sort:                      optional.Interface{},
	})
	if err != nil {
		diag.FromErr(err)
	}

	contents := common.ConvertStructToMaps(responses.Contents)

	rd.SetId(uuid.NewV4().String())
	rd.Set("contents", contents)
	rd.Set("total_count", responses.TotalCount)

	return nil
}

func tgwPeeringDataSourceElem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("TransitGatewayPeeringId"):          {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Id"},
			common.ToSnakeCase("TransitGatewayPeeringName"):        {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Name"},
			common.ToSnakeCase("TransitGatewayPeeringDescription"): {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Description"},
			common.ToSnakeCase("TransitGatewayPeeringState"):       {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering State"},
			common.ToSnakeCase("RequesterTransitGatewayId"):        {Type: schema.TypeString, Computed: true, Description: "Requester Transit Gateway Id"},
			common.ToSnakeCase("RequesterProjectId"):               {Type: schema.TypeString, Computed: true, Description: "Requester Project Id"},
			common.ToSnakeCase("ApproverTransitGatewayId"):         {Type: schema.TypeString, Computed: true, Description: "Approver Transit Gateway Id"},
			common.ToSnakeCase("ApproverProjectId"):                {Type: schema.TypeString, Computed: true, Description: "Approver Project Id"},
			common.ToSnakeCase("Automated"):                        {Type: schema.TypeBool, Computed: true, Description: "Is Automated"},
			common.ToSnakeCase("ProjectId"):                        {Type: schema.TypeString, Computed: true, Description: "Project Id"},
			common.ToSnakeCase("CompletedDt"):                      {Type: schema.TypeString, Computed: true, Description: "Completed Date"},
			common.ToSnakeCase("CreatedBy"):                        {Type: schema.TypeString, Computed: true, Description: "Created By"},
			common.ToSnakeCase("CreatedDt"):                        {Type: schema.TypeString, Computed: true, Description: "Created Date"},
			common.ToSnakeCase("ModifiedBy"):                       {Type: schema.TypeString, Computed: true, Description: "Modified By"},
			common.ToSnakeCase("ModifiedDt"):                       {Type: schema.TypeString, Computed: true, Description: "Modified Date"},
		},
	}

}

func DataSourceTgwPeeringsByTgw() *schema.Resource {
	return &schema.Resource{
		ReadContext: tgwPeeringdataSourceListByTgw,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("TransitGatewayId"): {Type: schema.TypeString, Required: true, Description: "Transit Gateway ID"},
			"contents":                             {Type: schema.TypeList, Optional: true, Description: "List of TGW Peerings", Elem: tgwPeeringDataSourceElem()},
			"total_count":                          {Type: schema.TypeInt, Computed: true, Description: "Total Count of TGW Peerings"},
		},
		Description: "provides List of TGW Peerings By Transit Gateway",
	}
}

func tgwPeeringdataSourceListByTgw(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	responses, _, err := inst.Client.TransitGateway.GetTransitGatewayPeeringListByTransitGateway(ctx, rd.Get("transit_gateway_id").(string))

	if err != nil {
		return diag.FromErr(err)
	}

	contents := common.ConvertStructToMaps(responses.Contents)
	rd.SetId(uuid.NewV4().String())
	rd.Set("contents", contents)
	rd.Set("total_count", responses.TotalCount)

	return nil
}

func DataSourceTransitGatewayPeeringDetail() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceTransitGatewayPeeringDetailRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("TransitGatewayPeeringId"):          {Type: schema.TypeString, Required: true, Description: "Transit Gateway Peering Id"},
			common.ToSnakeCase("TransitGatewayPeeringName"):        {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Name"},
			common.ToSnakeCase("TransitGatewayPeeringDescription"): {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering Description"},
			common.ToSnakeCase("TransitGatewayPeeringState"):       {Type: schema.TypeString, Computed: true, Description: "Transit Gateway Peering State"},
			common.ToSnakeCase("RequesterTransitGatewayId"):        {Type: schema.TypeString, Computed: true, Description: "Requester Transit Gateway Id"},
			common.ToSnakeCase("RequesterProjectId"):               {Type: schema.TypeString, Computed: true, Description: "Requester Project Id"},
			common.ToSnakeCase("ApproverTransitGatewayId"):         {Type: schema.TypeString, Computed: true, Description: "Approver Transit Gateway Id"},
			common.ToSnakeCase("ApproverProjectId"):                {Type: schema.TypeString, Computed: true, Description: "Approver Project Id"},
			common.ToSnakeCase("RequestedBy"):                      {Type: schema.TypeString, Computed: true, Description: "Requested By"},
			common.ToSnakeCase("RequestedDt"):                      {Type: schema.TypeString, Computed: true, Description: "Requested Date"},
			common.ToSnakeCase("ApprovedBy"):                       {Type: schema.TypeString, Computed: true, Description: "Approved By"},
			common.ToSnakeCase("ApprovedDt"):                       {Type: schema.TypeString, Computed: true, Description: "Approved Date"},
			common.ToSnakeCase("Automated"):                        {Type: schema.TypeBool, Computed: true, Description: "Automated"},
			common.ToSnakeCase("ProjectId"):                        {Type: schema.TypeString, Computed: true, Description: "Project Id"},
			common.ToSnakeCase("ServiceZoneId"):                    {Type: schema.TypeString, Computed: true, Description: "Service Zone Id"},
			common.ToSnakeCase("CreatedBy"):                        {Type: schema.TypeString, Computed: true, Description: "Created By"},
			common.ToSnakeCase("CreatedDt"):                        {Type: schema.TypeString, Computed: true, Description: "Created Date"},
			common.ToSnakeCase("ModifiedBy"):                       {Type: schema.TypeString, Computed: true, Description: "Modified By"},
			common.ToSnakeCase("ModifiedDt"):                       {Type: schema.TypeString, Computed: true, Description: "Modified Date"},
		},
		Description: "Provides a Transit Gateway Peering detail.",
	}
}

func resourceTransitGatewayPeeringDetailRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	tgwPeeringInfo, _, err := inst.Client.TransitGateway.GetTransitGatewayPeeringInfo(ctx, rd.Get(common.ToSnakeCase("TransitGatewayPeeringId")).(string))
	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.SetId(uuid.NewV4().String())
	rd.Set(common.ToSnakeCase("TransitGatewayPeeringId"), tgwPeeringInfo.TransitGatewayPeeringId)
	rd.Set(common.ToSnakeCase("TransitGatewayPeeringName"), tgwPeeringInfo.TransitGatewayPeeringName)
	rd.Set(common.ToSnakeCase("TransitGatewayPeeringDescription"), tgwPeeringInfo.TransitGatewayPeeringDescription)
	rd.Set(common.ToSnakeCase("TransitGatewayPeeringState"), tgwPeeringInfo.TransitGatewayPeeringState)
	rd.Set(common.ToSnakeCase("RequesterTransitGatewayId"), tgwPeeringInfo.RequesterTransitGatewayId)
	rd.Set(common.ToSnakeCase("RequesterProjectId"), tgwPeeringInfo.RequesterProjectId)
	rd.Set(common.ToSnakeCase("ApproverTransitGatewayId"), tgwPeeringInfo.ApproverTransitGatewayId)
	rd.Set(common.ToSnakeCase("ApproverProjectId"), tgwPeeringInfo.ApproverProjectId)
	rd.Set(common.ToSnakeCase("RequestedBy"), tgwPeeringInfo.RequestedBy)
	rd.Set(common.ToSnakeCase("RequestedDt"), time.Time.String(tgwPeeringInfo.RequestedDt))
	rd.Set(common.ToSnakeCase("ApprovedBy"), tgwPeeringInfo.ApprovedBy)
	rd.Set(common.ToSnakeCase("ApprovedDt"), time.Time.String(tgwPeeringInfo.ApprovedDt))
	rd.Set(common.ToSnakeCase("Automated"), tgwPeeringInfo.Automated)
	rd.Set(common.ToSnakeCase("ProjectId"), tgwPeeringInfo.ProjectId)
	rd.Set(common.ToSnakeCase("ServiceZoneId"), tgwPeeringInfo.ServiceZoneId)
	rd.Set(common.ToSnakeCase("CreatedBy"), tgwPeeringInfo.CreatedBy)
	rd.Set(common.ToSnakeCase("CreatedDt"), time.Time.String(tgwPeeringInfo.CreatedDt))
	rd.Set(common.ToSnakeCase("ModifiedBy"), tgwPeeringInfo.ModifiedBy)
	rd.Set(common.ToSnakeCase("ModifiedDt"), time.Time.String(tgwPeeringInfo.ModifiedDt))

	return nil
}
