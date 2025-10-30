package configinspection

import (
	"context"
	"fmt"
	"time"

	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	configinspection "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/config-inspection"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func init() {
	samsungcloudplatform.RegisterResource("samsungcloudplatform_config_inspection", Configinspections())
}

func Configinspections() *schema.Resource {
	return &schema.Resource{
		CreateContext: datasourceCICreate,
		ReadContext:   datasourceCIRead,
		DeleteContext: datasourceCIDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("CspType"): {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"SCP", "AWS", "Azure"}, false),
				Description:  "CSP Type (SCP, AWS, Azure)",
			},
			common.ToSnakeCase("DiagnosisCheckType"): {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"BP", "SSI"}, false),
				Description:  "BP or SSI",
			},
			common.ToSnakeCase("DiagnosisType"): {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Console", "SSI"}, false),
				Description:  "Diagnosis Type(Console, SSI)",
			},
			common.ToSnakeCase("PlanType"): {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"STANDARD", "MONTHLY"}, false),
				Description:  "Plan Type (STANDARD, MONTHLY)",
			},
			common.ToSnakeCase("DiagnosisObjectRequestList"): {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("DiagnosisAccountId"): {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Diagnosis Account ID",
						},
						common.ToSnakeCase("DiagnosisId"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Diagnosis ID",
						},
						common.ToSnakeCase("DiagnosisName"): {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Diagnosis Name",
						},
					},
				},
			},
			common.ToSnakeCase("Tags"): {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			common.ToSnakeCase("ScheduleRequest"): {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("DiagnosisStartTimePattern"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Diagnosis Start Time (5-minute interval, 00~23 hours, 00~55 minutes)",
						},
						common.ToSnakeCase("FrequencyType"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Diagnosis Schedule Type (Monthly, Weekly, Daily)",
						},
						common.ToSnakeCase("FrequencyValue"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Diagnosis Schedule (0131, MondaySunday, everyDay)",
						},
						common.ToSnakeCase("UseDiagnosisCheckTypeBP"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Use Best Practice Checklist",
						},
						common.ToSnakeCase("UseDiagnosisCheckTypeSSI"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Use SSI Checklist",
						},
					},
				},
			},
			common.ToSnakeCase("AuthKeyRequest"): {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("AuthKeyId"): {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Auth Key ID",
						},
					},
				},
			},
		},
		Description: "Save Config Inspection",
	}
}

func datasourceCIRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) (diagnostics diag.Diagnostics) {
	inst := meta.(*client.Instance)

	info, _, err := inst.Client.ConfigInspection.GetConfigInspectionProductDetail(ctx, rd.Id())

	if err != nil {
		rd.SetId("")
		if common.IsDeleted(err) {
			return nil
		}

		return diag.FromErr(err)
	}

	rd.SetId(rd.Id())
	fmt.Printf("datasourceCIDetail: %v\n", info)
	return nil
}

func datasourceCICreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) (diagnostics diag.Diagnostics) {
	cspType := rd.Get("csp_type").(string)
	diagnosisCheckType := rd.Get("diagnosis_check_type").(string)
	diagnosisType := rd.Get("diagnosis_type").(string)
	planType := rd.Get("plan_type").(string)
	diagnosisObjectRequestList := rd.Get("diagnosis_object_request_list").(common.HclListObject)
	tags := rd.Get("tags").(map[string]interface{})
	scheduleRequest := rd.Get("schedule_request").(common.HclListObject)[0].(map[string]interface{})
	authKeyRequest := rd.Get("auth_key_request").(common.HclListObject)[0].(map[string]interface{})
	requestList := convertRequests(diagnosisObjectRequestList)
	var err error = nil
	defer func() {
		if err != nil {
			diagnostics = diag.FromErr(err)
		}
	}()

	inst := meta.(*client.Instance)

	diagnosisObjectListRequest := configinspection.DiagnosisObjectListRequest{
		AuthKeyRequest: &configinspection.AuthKeyRequest{
			AuthKeyId: authKeyRequest["auth_key_id"].(string),
		},
		CspType:                    cspType,
		DiagnosisCheckType:         diagnosisCheckType,
		DiagnosisType:              diagnosisType,
		PlanType:                   planType,
		DiagnosisObjectRequestList: requestList,
		ScheduleRequest: &configinspection.DiagnosisScheduleRequest{
			DiagnosisStartTimePattern: scheduleRequest["diagnosis_start_time_pattern"].(string),
			FrequencyType:             scheduleRequest["frequency_type"].(string),
			FrequencyValue:            scheduleRequest["frequency_value"].(string),
			UseDiagnosisCheckTypeBP:   scheduleRequest["use_diagnosis_check_type_bp"].(string),
			UseDiagnosisCheckTypeSSI:  scheduleRequest["use_diagnosis_check_type_ssi"].(string),
		},
	}
	result, _, err := inst.Client.ConfigInspection.ConfigInspectionProductCreate(ctx, diagnosisObjectListRequest, tags)

	if err != nil {
		return
	}

	fmt.Printf("datasourceCICreate result: %v\n", result)

	// wait for server state
	time.Sleep(10 * time.Second)
	rd.SetId(result.DiagnosisId)

	if err != nil {
		return
	}

	return datasourceCIRead(ctx, rd, meta)
}

func datasourceCIDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) (diagnostics diag.Diagnostics) {
	var err error = nil
	defer func() {
		if err != nil {
			diagnostics = diag.FromErr(err)
		}
	}()

	inst := meta.(*client.Instance)

	result, err := inst.Client.ConfigInspection.ConfigInspectionProductDelete(ctx, rd.Id())

	fmt.Printf("datasourceCIDelete result: %v\n", result)

	if err != nil {
		return
	}

	time.Sleep(10 * time.Second)

	if err != nil {
		return
	}

	return datasourceCIRead(ctx, rd, meta)
}

func convertRequests(schedules common.HclListObject) []configinspection.DiagnosisObjectRequest {
	var result []configinspection.DiagnosisObjectRequest
	for _, itemObject := range schedules {
		item := itemObject.(common.HclKeyValueObject)
		var info configinspection.DiagnosisObjectRequest
		if v, ok := item["diagnosis_id"]; ok {
			info.DiagnosisId = v.(string)
		}
		if v, ok := item["diagnosis_account_id"]; ok {
			info.DiagnosisAccountId = v.(string)
		}
		if v, ok := item["diagnosis_name"]; ok {
			info.DiagnosisName = v.(string)
		}
		result = append(result, info)
	}
	return result
}
