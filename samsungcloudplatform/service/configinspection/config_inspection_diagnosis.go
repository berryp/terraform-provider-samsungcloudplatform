package configinspection

import (
	"context"

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
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_config_inspection_diagnosis_request", DatasourceDianosisRequest())
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_config_inspection_diagnoses", DatasourceDianoses())
	samsungcloudplatform.RegisterDataSource("samsungcloudplatform_config_inspection_diagnosis", DatasourceDianosis())
}

// List diagnosis
func DatasourceDianoses() *schema.Resource {
	return &schema.Resource{
		ReadContext: getDianosisList,
		Schema: map[string]*schema.Schema{
			// Input parameters
			common.ToSnakeCase("DiagnosisName"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Diagnosis Name",
			},
			common.ToSnakeCase("DiagnosisState"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "State of diagnosis",
			},
			common.ToSnakeCase("EndDate"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "End Date",
			},
			common.ToSnakeCase("StartDate"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Start Date",
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

			// Output parameters
			common.ToSnakeCase("Contents"): {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("Check"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Check",
						},
						common.ToSnakeCase("CspType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CSP Type (SCP, AWS, Azure)",
						},
						common.ToSnakeCase("DiagnosisAccountId"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis Account",
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
						common.ToSnakeCase("DiagnosisRequestSequence"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "KAFKA Request Number",
						},
						common.ToSnakeCase("DiagnosisResult"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Result of diagnosis",
						},
						common.ToSnakeCase("DiagnosisTotalCnt"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of diagnostic items",
						},
						common.ToSnakeCase("Error"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Error",
						},
						common.ToSnakeCase("Fail"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Fail",
						},
						common.ToSnakeCase("Na"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "N/A",
						},
						common.ToSnakeCase("Pass"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Pass",
						},
						common.ToSnakeCase("ProceedDate"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Processing date",
						},
						common.ToSnakeCase("Total"): {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total",
						},
					},
				},
				Description: "List of search items",
			},
			common.ToSnakeCase("TotalCount"): {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of hits",
			},
		},
		Description: "Get List of Config Inspection",
	}
}

// Diagnosis detail
func DatasourceDianosis() *schema.Resource {
	return &schema.Resource{
		ReadContext: getDiagnosisDetail,
		Schema: map[string]*schema.Schema{
			// Payload
			common.ToSnakeCase("DiagnosisId"): {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DiagnosisId is obtained through Config Inspection List",
			},
			common.ToSnakeCase("DiagnosisRequestSequence"): {
				Type:        schema.TypeString,
				Required:    true,
				Description: "KAFKA Request Sequence is obtained through Report Diagnosis result list",
			},

			// Response
			common.ToSnakeCase("ChecklistName"): {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Check list name",
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
			common.ToSnakeCase("DiagnosisName"): {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Diagnosis Name",
			},
			common.ToSnakeCase("ProceedDate"): {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Processed date",
			},
			common.ToSnakeCase("ResultDetailList"): {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						common.ToSnakeCase("ActionGuide"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Action guide",
						},
						common.ToSnakeCase("Changed"): {
							Type:     schema.TypeBool,
							Computed: true,
						},
						common.ToSnakeCase("DiagnosisCheckType"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "BP or SSI",
						},
						common.ToSnakeCase("DiagnosisCriteria"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis criteria",
						},
						common.ToSnakeCase("DiagnosisItem"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis item",
						},
						common.ToSnakeCase("DiagnosisLayer"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis layer",
						},
						common.ToSnakeCase("DiagnosisMethod"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Diagnosis method",
						},
						common.ToSnakeCase("DiagnosisResult"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Result of diagnosis",
						},
						common.ToSnakeCase("ResultContents"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Result contents",
						},
						common.ToSnakeCase("SubCategory"): {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sub-category",
						},
					},
				},
			},
			common.ToSnakeCase("TotalCount"): {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total count",
			},
		},
		Description: "Get diagnosis result details",
	}
}

// Diagnosis request
func DatasourceDianosisRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: createDiagnosisRequest,
		Schema: map[string]*schema.Schema{
			common.ToSnakeCase("ProjectId"): {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Project ID",
			},
			common.ToSnakeCase("AccessKey"): {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access Key",
			},
			common.ToSnakeCase("SecretKey"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secret Key (optional)",
			},
			common.ToSnakeCase("TenantId"): {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Tenant ID (Azure, optional)",
			},
			common.ToSnakeCase("DiagnosisCheckType"): {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"BP", "SSI"}, false),
				Description:  "Diagnosis Check Type (either 'BP' or 'SSI')",
			},
			common.ToSnakeCase("DiagnosisId"): {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Diagnosis ID (obtained through Config Inspection List)",
			},
			common.ToSnakeCase("Result"): {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Result of the diagnosis request",
			},
		},
		Description: "Request Config Inspection Diagnosis",
	}
}

func createDiagnosisRequest(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	opts := configinspection.DiagnosisRequest{
		ProjectId:          rd.Get(common.ToSnakeCase("ProjectId")).(string),
		AccessKey:          rd.Get(common.ToSnakeCase("AccessKey")).(string),
		DiagnosisCheckType: rd.Get(common.ToSnakeCase("DiagnosisCheckType")).(string),
		DiagnosisId:        rd.Get(common.ToSnakeCase("DiagnosisId")).(string),
	}

	if v, ok := rd.GetOk(common.ToSnakeCase("SecretKey")); ok {
		opts.SecretKey = v.(string)
	}
	if v, ok := rd.GetOk(common.ToSnakeCase("TenantId")); ok {
		opts.TenantId = v.(string)
	}

	info, _, err := inst.Client.ConfigInspection.RequestNewConfigInspection(ctx, opts)

	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.SetId(uuid.NewV4().String())
	rd.Set("result", info.Result)

	return nil
}

func getDianosisList(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	opts := &configinspection.ConfigInspectionDashboardOpenApiControllerApiGetDiagnosisResultListOpts{}
	if v, ok := rd.GetOk("diagnosis_name"); ok {
		opts.DiagnosisName = optional.NewString(v.(string))
	}
	if v, ok := rd.GetOk("diagnosis_state"); ok {
		opts.DiagnosisState = optional.NewString(v.(string))
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
	if v, ok := rd.GetOk("sort"); ok {
		sortValues := v.([]interface{})
		var sortStrings []string
		for _, val := range sortValues {
			sortStrings = append(sortStrings, val.(string))
		}
		opts.Sort = optional.NewInterface(sortStrings)
	}

	responses, _, err := inst.Client.ConfigInspection.GetConfigInspectionList(ctx, opts)
	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	contents := common.ConvertStructToMaps(responses.Contents)

	rd.SetId(uuid.NewV4().String())
	rd.Set("contents", contents)
	rd.Set("total_count", responses.TotalCount)

	return nil
}

func getDiagnosisDetail(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)
	diagnosisId := rd.Get(common.ToSnakeCase("DiagnosisId")).(string)
	diagnosisRequestSequence := rd.Get(common.ToSnakeCase("DiagnosisRequestSequence")).(string)

	info, _, err := inst.Client.ConfigInspection.GetConfigInspectionDetail(ctx, diagnosisId, diagnosisRequestSequence)

	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.SetId(uuid.NewV4().String())
	rd.Set(common.ToSnakeCase("ChecklistName"), info.ChecklistName)
	rd.Set(common.ToSnakeCase("DiagnosisAccount"), info.DiagnosisAccount)
	rd.Set(common.ToSnakeCase("DiagnosisCheckType"), info.DiagnosisCheckType)
	rd.Set(common.ToSnakeCase("DiagnosisName"), info.DiagnosisName)
	rd.Set(common.ToSnakeCase("ProceedDate"), info.ProceedDate)
	rd.Set(common.ToSnakeCase("ResultDetailList"), common.ConvertStructToMaps(info.ResultDetailList))
	rd.Set(common.ToSnakeCase("TotalCount"), info.TotalCount)

	return nil
}
