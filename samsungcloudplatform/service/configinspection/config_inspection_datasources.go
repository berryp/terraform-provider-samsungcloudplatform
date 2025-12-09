package configinspection

import (
	"context"
	"time"

	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/client"
	"github.com/SamsungSDSCloud/terraform-provider-samsungcloudplatform/v3/samsungcloudplatform/common"
	configinspection "github.com/SamsungSDSCloud/terraform-sdk-samsungcloudplatform/v3/library/config-inspection"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	uuid "github.com/satori/go.uuid"
)

func init() {
	samsungcloudplatform.RegisterDataSource("Config Inspection", "samsungcloudplatform_config_inspections", DatasourceConfigInspections())
	samsungcloudplatform.RegisterDataSource("Config Inspection", "samsungcloudplatform_config_inspection", DatasourceConfigInspection())
}

func DatasourceConfigInspection() *schema.Resource {
	return &schema.Resource{
		ReadContext: getDetail,
		Schema: map[string]*schema.Schema{
			// Payload
			common.ToSnakeCase("DiagnosisId"): {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Diagnosis ID",
			},
			// Response
			common.ToSnakeCase("AuthKeyResponse"): {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("AuthKeyCreateDt"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auth Key Start Date",
						},
						common.ToSnakeCase("AuthKeyExpireDt"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auth Key Expiration Date",
						},
						common.ToSnakeCase("AuthKeyId"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "start date",
						},
						common.ToSnakeCase("AuthKeyState"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auth Key State",
						},
						common.ToSnakeCase("UserId"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User id",
						},
					},
				},
			},
			common.ToSnakeCase("ScheduleResponse"): {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("DiagnosisId"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis ID",
						},
						common.ToSnakeCase("DiagnosisStartTimePattern"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Start Time (5-minute interval, 00~23 hours, 00~55 minutes)",
						},
						common.ToSnakeCase("FrequencyType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Schedule Type (Monthly, Weekly, Daily)",
						},
						common.ToSnakeCase("FrequencyValue"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Schedule (0131, MondaySunday, everyDay)",
						},
						common.ToSnakeCase("UseDiagnosisCheckTypeBP"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Use Best Practice Checklist",
						},
						common.ToSnakeCase("UseDiagnosisCheckTypeSSI"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Use SSI Checklist",
						},
					},
				},
			},
			common.ToSnakeCase("SummaryResponse"): {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("CspType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CSP Type (SCP, AWS, Azure)",
						},
						common.ToSnakeCase("DiagnosisAccount"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis account",
						},
						common.ToSnakeCase("DiagnosisCheckType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BP or SSI",
						},
						common.ToSnakeCase("DiagnosisId"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis ID",
						},
						common.ToSnakeCase("DiagnosisName"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Name",
						},
						common.ToSnakeCase("DiagnosisType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Type(Console, SSI)",
						},
						common.ToSnakeCase("ErrorState"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error State",
						},
						common.ToSnakeCase("PlanType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Plan Type (STANDARD, MONTHLY)",
						},
						common.ToSnakeCase("RecentDiagnosisDt"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Recent Diagnosis Date",
						},
						common.ToSnakeCase("RecentDiagnosisState"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Recent Diagnosis State",
						},
						common.ToSnakeCase("CreatedDt"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Start Date",
						},
					},
				},
			},
		},
		Description: "Get diagnosis result details",
	}
}

func DatasourceConfigInspections() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceCIList,
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("CspType"): {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"SCP", "AWS", "Azure"}, false),
				},
				Optional:    true,
				Description: "CSP Type (SCP, AWS, Azure)",
			},
			common.ToSnakeCase("DiagnosisAccountId"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Diagnosis Account",
			},
			common.ToSnakeCase("DiagnosisName"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Diagnosis Name",
			},
			common.ToSnakeCase("EndDate"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "end date",
			},
			common.ToSnakeCase("RecentDiagnosisState"): {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Recent Diagnosis State",
			},
			common.ToSnakeCase("StartDate"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "start date",
			},
			common.ToSnakeCase("Page"): {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Request page number",
			},
			common.ToSnakeCase("Size"): {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     20,
				Description: "Number of items in the Page",
			},
			common.ToSnakeCase("Sort"): {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Sort condition",
			},
			common.ToSnakeCase("Contents"): {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("CSPType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CSP Type (SCP, AWS, Azure)",
						},
						common.ToSnakeCase("DiagnosisAccount"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis account",
						},
						common.ToSnakeCase("DiagnosisCheckType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BP or SSI",
						},
						common.ToSnakeCase("DiagnosisID"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis ID",
						},
						common.ToSnakeCase("DiagnosisName"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Name",
						},
						common.ToSnakeCase("DiagnosisType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Type(Console, SSI)",
						},
						common.ToSnakeCase("ErrorState"): {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Error State",
						},
						common.ToSnakeCase("PlanType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Plan Type (STANDARD, MONTHLY)",
						},
						common.ToSnakeCase("RecentDiagnosisDt"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Recent Diagnosis Date",
						},
						common.ToSnakeCase("RecentDiagnosisState"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Recent Diagnosis State",
						},
						common.ToSnakeCase("CreatedDt"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Start Date",
						},
					},
				},
			},
			common.ToSnakeCase("TotalCount"): {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
		Description: "Get List of Config Inspection",
	}
}

func datasourceCIList(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	opts := &configinspection.ConfigInspectionOpenApiControllerApiConfigInspectionListOpts{}
	if v, ok := rd.GetOk("diagnosis_account_id"); ok {
		opts.DiagnosisAccountId = optional.NewString(v.(string))
	}

	if v, ok := rd.GetOk("diagnosis_name"); ok {
		opts.DiagnosisName = optional.NewString(v.(string))
	}
	if v, ok := rd.GetOk("end_date"); ok {
		opts.EndDate = optional.NewString(v.(string))
	}
	if v, ok := rd.GetOk("start_date"); ok {
		opts.StartDate = optional.NewString(v.(string))
	}
	if v, ok := rd.GetOk("page"); ok {
		opts.Page = optional.NewInt32(int32(v.(int)))
	}
	if v, ok := rd.GetOk("size"); ok {
		opts.Size = optional.NewInt32(int32(v.(int)))
	}
	if v, ok := rd.GetOk("csp_type"); ok {
		cspType := v.([]interface{})
		var cspTypeStrings []string
		for _, val := range cspType {
			cspTypeStrings = append(cspTypeStrings, val.(string))
		}
		opts.CspType = optional.NewInterface(cspTypeStrings)
	}
	if v, ok := rd.GetOk("recent_diagnosis_state"); ok {
		recentDiagnosisState := v.([]interface{})
		var recentDiagnosisStateStrings []string
		for _, val := range recentDiagnosisState {
			recentDiagnosisStateStrings = append(recentDiagnosisStateStrings, val.(string))
		}
		opts.RecentDiagnosisState = optional.NewInterface(recentDiagnosisStateStrings)
	}
	if v, ok := rd.GetOk("sort"); ok {
		sortValues := v.([]interface{})
		var sortStrings []string
		for _, val := range sortValues {
			sortStrings = append(sortStrings, val.(string))
		}
		opts.Sort = optional.NewInterface(sortStrings)
	}
	info, _, err := inst.Client.ConfigInspection.GetConfigInspectionProductList(ctx, opts)

	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	contents := common.ConvertStructToMaps(info.Contents)

	rd.SetId(uuid.NewV4().String())
	rd.Set("contents", contents)
	rd.Set("total_count", info.TotalCount)
	return nil
}

func getDetail(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	diagnosisId := rd.Get(common.ToSnakeCase("DiagnosisId")).(string)

	info, _, err := inst.Client.ConfigInspection.GetConfigInspectionProductDetail(ctx, diagnosisId)

	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	if info.AuthKeyResponse != nil {
		authKeyResponse := []interface{}{
			map[string]interface{}{
				common.ToSnakeCase("AuthKeyCreateDt"): info.AuthKeyResponse.AuthKeyCreateDt.Format(time.RFC3339),
				common.ToSnakeCase("AuthKeyExpireDt"): info.AuthKeyResponse.AuthKeyExpireDt.Format(time.RFC3339),
				common.ToSnakeCase("AuthKeyId"):       info.AuthKeyResponse.AuthKeyId,
				common.ToSnakeCase("AuthKeyState"):    info.AuthKeyResponse.AuthKeyState,
				common.ToSnakeCase("UserId"):          info.AuthKeyResponse.UserId,
			},
		}
		rd.Set(common.ToSnakeCase("AuthKeyResponse"), authKeyResponse)
	}

	if info.ScheduleResponse != nil {
		scheduleResponse := []interface{}{
			map[string]interface{}{
				common.ToSnakeCase("DiagnosisId"):               info.ScheduleResponse.DiagnosisId,
				common.ToSnakeCase("DiagnosisStartTimePattern"): info.ScheduleResponse.DiagnosisStartTimePattern,
				common.ToSnakeCase("FrequencyType"):             info.ScheduleResponse.FrequencyType,
				common.ToSnakeCase("FrequencyValue"):            info.ScheduleResponse.FrequencyValue,
				common.ToSnakeCase("UseDiagnosisCheckTypeBP"):   info.ScheduleResponse.UseDiagnosisCheckTypeBP,
				common.ToSnakeCase("UseDiagnosisCheckTypeSSI"):  info.ScheduleResponse.UseDiagnosisCheckTypeSSI,
			},
		}
		rd.Set(common.ToSnakeCase("ScheduleResponse"), scheduleResponse)
	}

	if info.SummaryResponse != nil {
		summaryResponse := []interface{}{
			map[string]interface{}{
				common.ToSnakeCase("CspType"):              info.SummaryResponse.CspType,
				common.ToSnakeCase("DiagnosisAccount"):     info.SummaryResponse.DiagnosisAccount,
				common.ToSnakeCase("DiagnosisCheckType"):   info.SummaryResponse.DiagnosisCheckType,
				common.ToSnakeCase("DiagnosisId"):          info.SummaryResponse.DiagnosisId,
				common.ToSnakeCase("DiagnosisName"):        info.SummaryResponse.DiagnosisName,
				common.ToSnakeCase("DiagnosisType"):        info.SummaryResponse.DiagnosisType,
				common.ToSnakeCase("ErrorState"):           info.SummaryResponse.ErrorState,
				common.ToSnakeCase("PlanType"):             info.SummaryResponse.PlanType,
				common.ToSnakeCase("RecentDiagnosisDt"):    info.SummaryResponse.RecentDiagnosisDt.Format(time.RFC3339),
				common.ToSnakeCase("RecentDiagnosisState"): info.SummaryResponse.RecentDiagnosisState,
				common.ToSnakeCase("CreatedDt"):            info.SummaryResponse.CreatedDt.Format(time.RFC3339),
			},
		}
		rd.Set(common.ToSnakeCase("SummaryResponse"), summaryResponse)
	}

	rd.SetId(uuid.NewV4().String())

	return nil
}
