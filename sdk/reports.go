package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ReportsService service

type GetListOfScheduledReportsQueryParams struct {
	ViewGroupID string `url:"viewGroupId,omitempty"` //viewGroupId of viewgroup for report
	ViewID      string `url:"viewId,omitempty"`      //viewId of view for report
}

type ResponseReportsCreateOrScheduleAReport struct {
	Tags              []string                                            `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                              `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseReportsCreateOrScheduleAReportDeliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                                `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseReportsCreateOrScheduleAReportExecutions `json:"executions,omitempty"`        //
	Name              string                                              `json:"name,omitempty"`              // report name
	ReportID          string                                              `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                               `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseReportsCreateOrScheduleAReportSchedule     `json:"schedule,omitempty"`          //
	View              *ResponseReportsCreateOrScheduleAReportView         `json:"view,omitempty"`              //
	ViewGroupID       string                                              `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                              `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseReportsCreateOrScheduleAReportDeliveries interface{}
type ResponseReportsCreateOrScheduleAReportExecutions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseReportsCreateOrScheduleAReportSchedule interface{}
type ResponseReportsCreateOrScheduleAReportView struct {
	FieldGroups *[]ResponseReportsCreateOrScheduleAReportViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsCreateOrScheduleAReportViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseReportsCreateOrScheduleAReportViewFormat        `json:"format,omitempty"`      //
	Name        string                                                   `json:"name,omitempty"`        // view name
	ViewID      string                                                   `json:"viewId,omitempty"`      // view Id
	Description string                                                   `json:"description,omitempty"` // view description
	ViewInfo    string                                                   `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseReportsCreateOrScheduleAReportViewFieldGroups struct {
	FieldGroupDisplayName string                                                         `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                         `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseReportsCreateOrScheduleAReportViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type ResponseReportsCreateOrScheduleAReportViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseReportsCreateOrScheduleAReportViewFilters struct {
	DisplayName string                                                  `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                  `json:"name,omitempty"`        // filter name
	Type        string                                                  `json:"type,omitempty"`        // filter type
	Value       *ResponseReportsCreateOrScheduleAReportViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type.
}
type ResponseReportsCreateOrScheduleAReportViewFiltersValue interface{}
type ResponseReportsCreateOrScheduleAReportViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
}
type ResponseReportsGetListOfScheduledReports []ResponseItemReportsGetListOfScheduledReports // Array of ResponseReportsGetListOfScheduledReports
type ResponseItemReportsGetListOfScheduledReports struct {
	Tags              []string                                                  `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                                    `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseItemReportsGetListOfScheduledReportsDeliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                                      `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseItemReportsGetListOfScheduledReportsExecutions `json:"executions,omitempty"`        //
	Name              string                                                    `json:"name,omitempty"`              // report name
	ReportID          string                                                    `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                                     `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseItemReportsGetListOfScheduledReportsSchedule     `json:"schedule,omitempty"`          //
	View              *ResponseItemReportsGetListOfScheduledReportsView         `json:"view,omitempty"`              //
	ViewGroupID       string                                                    `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                                    `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseItemReportsGetListOfScheduledReportsDeliveries interface{}
type ResponseItemReportsGetListOfScheduledReportsExecutions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseItemReportsGetListOfScheduledReportsSchedule interface{}
type ResponseItemReportsGetListOfScheduledReportsView struct {
	FieldGroups *[]ResponseItemReportsGetListOfScheduledReportsViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseItemReportsGetListOfScheduledReportsViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseItemReportsGetListOfScheduledReportsViewFormat        `json:"format,omitempty"`      //
	Name        string                                                         `json:"name,omitempty"`        // view name
	ViewID      string                                                         `json:"viewId,omitempty"`      // view Id
	Description string                                                         `json:"description,omitempty"` // view description
	ViewInfo    string                                                         `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseItemReportsGetListOfScheduledReportsViewFieldGroups struct {
	FieldGroupDisplayName string                                                               `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                               `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseItemReportsGetListOfScheduledReportsViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type ResponseItemReportsGetListOfScheduledReportsViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseItemReportsGetListOfScheduledReportsViewFilters struct {
	DisplayName string                                                        `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                        `json:"name,omitempty"`        // filter name
	Type        string                                                        `json:"type,omitempty"`        // filter type
	Value       *ResponseItemReportsGetListOfScheduledReportsViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type.
}
type ResponseItemReportsGetListOfScheduledReportsViewFiltersValue interface{}
type ResponseItemReportsGetListOfScheduledReportsViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
	Default    *bool  `json:"default,omitempty"`    // true, if the format type is considered default
}
type ResponseReportsGetAScheduledReport struct {
	Tags              []string                                        `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                          `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseReportsGetAScheduledReportDeliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                            `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseReportsGetAScheduledReportExecutions `json:"executions,omitempty"`        //
	Name              string                                          `json:"name,omitempty"`              // report name
	ReportID          string                                          `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                           `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseReportsGetAScheduledReportSchedule     `json:"schedule,omitempty"`          //
	View              *ResponseReportsGetAScheduledReportView         `json:"view,omitempty"`              //
	ViewGroupID       string                                          `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                          `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseReportsGetAScheduledReportDeliveries interface{}
type ResponseReportsGetAScheduledReportExecutions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseReportsGetAScheduledReportSchedule interface{}
type ResponseReportsGetAScheduledReportView struct {
	FieldGroups *[]ResponseReportsGetAScheduledReportViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsGetAScheduledReportViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseReportsGetAScheduledReportViewFormat        `json:"format,omitempty"`      //
	Name        string                                               `json:"name,omitempty"`        // view name
	ViewID      string                                               `json:"viewId,omitempty"`      // view Id
	Description string                                               `json:"description,omitempty"` // view description
	ViewInfo    string                                               `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseReportsGetAScheduledReportViewFieldGroups struct {
	FieldGroupDisplayName string                                                     `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                     `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseReportsGetAScheduledReportViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type ResponseReportsGetAScheduledReportViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseReportsGetAScheduledReportViewFilters struct {
	DisplayName string                                              `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                              `json:"name,omitempty"`        // filter name
	Type        string                                              `json:"type,omitempty"`        // filter type
	Value       *ResponseReportsGetAScheduledReportViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type.
}
type ResponseReportsGetAScheduledReportViewFiltersValue interface{}
type ResponseReportsGetAScheduledReportViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
	Default    *bool  `json:"default,omitempty"`    // true, if the format type is considered default
}
type ResponseReportsDeleteAScheduledReport struct {
	Message string `json:"message,omitempty"` // Response message
	Status  *int   `json:"status,omitempty"`  // Response Status
}
type ResponseReportsGetAllExecutionDetailsForAGivenReport struct {
	Tags              []string                                                          `json:"tags,omitempty"`              // array of tags for report
	DataCategory      string                                                            `json:"dataCategory,omitempty"`      // data category of the report
	Deliveries        *[]ResponseReportsGetAllExecutionDetailsForAGivenReportDeliveries `json:"deliveries,omitempty"`        // Array of available delivery channels
	ExecutionCount    *int                                                              `json:"executionCount,omitempty"`    // Total number of report executions
	Executions        *[]ResponseReportsGetAllExecutionDetailsForAGivenReportExecutions `json:"executions,omitempty"`        //
	Name              string                                                            `json:"name,omitempty"`              // report dataset name
	ReportID          string                                                            `json:"reportId,omitempty"`          // report Id
	ReportWasExecuted *bool                                                             `json:"reportWasExecuted,omitempty"` // true if atleast one execution has started
	Schedule          *ResponseReportsGetAllExecutionDetailsForAGivenReportSchedule     `json:"schedule,omitempty"`          //
	View              *ResponseReportsGetAllExecutionDetailsForAGivenReportView         `json:"view,omitempty"`              //
	ViewGroupID       string                                                            `json:"viewGroupId,omitempty"`       // viewGroupId of the viewgroup for the report
	ViewGroupVersion  string                                                            `json:"viewGroupVersion,omitempty"`  // version of viewgroup for the report
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportDeliveries interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportExecutions struct {
	EndTime       *int     `json:"endTime,omitempty"`       // Report execution pipeline end time
	Errors        []string `json:"errors,omitempty"`        //
	ExecutionID   string   `json:"executionId,omitempty"`   // Report execution Id.
	ProcessStatus string   `json:"processStatus,omitempty"` // Report execution status
	RequestStatus string   `json:"requestStatus,omitempty"` // Report execution acceptance status from scheduler
	StartTime     *int     `json:"startTime,omitempty"`     // Report execution pipeline start time
	Warnings      []string `json:"warnings,omitempty"`      //
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportSchedule interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportView struct {
	FieldGroups *[]ResponseReportsGetAllExecutionDetailsForAGivenReportViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsGetAllExecutionDetailsForAGivenReportViewFilters     `json:"filters,omitempty"`     //
	Format      *ResponseReportsGetAllExecutionDetailsForAGivenReportViewFormat        `json:"format,omitempty"`      //
	Name        string                                                                 `json:"name,omitempty"`        // view name
	ViewID      string                                                                 `json:"viewId,omitempty"`      // view Id
	Description string                                                                 `json:"description,omitempty"` // view description
	ViewInfo    string                                                                 `json:"viewInfo,omitempty"`    // view filters info
}
type ResponseReportsGetAllExecutionDetailsForAGivenReportViewFieldGroups interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportViewFilters interface{}
type ResponseReportsGetAllExecutionDetailsForAGivenReportViewFormat interface{}
type ResponseReportsGetAllViewGroups []ResponseItemReportsGetAllViewGroups // Array of ResponseReportsGetAllViewGroups
type ResponseItemReportsGetAllViewGroups struct {
	Category    string `json:"category,omitempty"`    // category of the view group
	Description string `json:"description,omitempty"` // view group description
	Name        string `json:"name,omitempty"`        // name of view group
	ViewGroupID string `json:"viewGroupId,omitempty"` // id of viewgroup
}
type ResponseReportsGetViewsForAGivenViewGroup struct {
	ViewGroupID string                                            `json:"viewGroupId,omitempty"` // viewgroup Id
	Views       *[]ResponseReportsGetViewsForAGivenViewGroupViews `json:"views,omitempty"`       //
}
type ResponseReportsGetViewsForAGivenViewGroupViews struct {
	Description string `json:"description,omitempty"` //
	ViewID      string `json:"viewId,omitempty"`      // Unique id for a view within viewgroup
	ViewName    string `json:"viewName,omitempty"`    // view name
}
type ResponseReportsGetViewDetailsForAGivenViewGroupView struct {
	Deliveries  *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewDeliveries  `json:"deliveries,omitempty"`  //
	Description string                                                            `json:"description,omitempty"` // view description
	FieldGroups *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewFilters     `json:"filters,omitempty"`     //
	Formats     *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewFormats     `json:"formats,omitempty"`     //
	Schedules   *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewSchedules   `json:"schedules,omitempty"`   //
	ViewID      string                                                            `json:"viewId,omitempty"`      // Unique view Id
	ViewInfo    string                                                            `json:"viewInfo,omitempty"`    // view filters info
	ViewName    string                                                            `json:"viewName,omitempty"`    // view name
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewDeliveries struct {
	Type    string `json:"type,omitempty"`    // delivery type
	Default *bool  `json:"default,omitempty"` // true, if the delivery type is considered default
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFieldGroups struct {
	FieldGroupDisplayName string                                                                  `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                                  `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewFieldGroupsFields `json:"fields,omitempty"`                //
	TableID               string                                                                  `json:"tableId,omitempty"`               // Table Id of the corresponding table mapped to fieldgroup
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFilters struct {
	AdditionalInfo *ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersAdditionalInfo `json:"additionalInfo,omitempty"` // Additional info for managing filter options
	CacheFilter    *bool                                                                     `json:"cacheFilter,omitempty"`    //
	DataType       string                                                                    `json:"dataType,omitempty"`       // data type of filter value
	DisplayName    string                                                                    `json:"displayName,omitempty"`    // filter label/displayname
	FilterSource   *ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersFilterSource   `json:"filterSource,omitempty"`   //
	Name           string                                                                    `json:"name,omitempty"`           // filter name
	Required       *bool                                                                     `json:"required,omitempty"`       // true if the filter is required
	TimeOptions    *[]ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersTimeOptions  `json:"timeOptions,omitempty"`    //
	Type           string                                                                    `json:"type,omitempty"`           // filter type. Used to handle filter value selection by the client for report configuration.
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersAdditionalInfo interface{}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersFilterSource struct {
	DataSource       *ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersFilterSourceDataSource `json:"dataSource,omitempty"`       //
	DisplayValuePath string                                                                            `json:"displayValuePath,omitempty"` // JSONPath of the label of filter option from the filter option as root
	RootPath         string                                                                            `json:"rootPath,omitempty"`         // JSONPath of the filter options array in the API response
	ValuePath        string                                                                            `json:"valuePath,omitempty"`        // JSONPath of the value of filter option from the filter option as root
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersFilterSourceDataSource interface{}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFiltersTimeOptions struct {
	Info     string `json:"info,omitempty"`     // Time range option description
	MaxValue *int   `json:"maxValue,omitempty"` // Maximum number of hours allowed for the time range option. (Client Validation)
	MinValue *int   `json:"minValue,omitempty"` // Minimum number of hours allowed for the time range option. (Client Validation)
	Name     string `json:"name,omitempty"`     // Time range option label
	Value    string `json:"value,omitempty"`    // Time range option value
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFormats struct {
	Format   string                                                              `json:"format,omitempty"`   // format type
	Name     string                                                              `json:"name,omitempty"`     // format name
	Default  *bool                                                               `json:"default,omitempty"`  // true, if the format type is considered default
	Template *ResponseReportsGetViewDetailsForAGivenViewGroupViewFormatsTemplate `json:"template,omitempty"` //
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewFormatsTemplate struct {
	JsTemplateID string `json:"jsTemplateId,omitempty"` // TemplateId of template
}
type ResponseReportsGetViewDetailsForAGivenViewGroupViewSchedules struct {
	Type    string `json:"type,omitempty"`    // schedule type
	Default *bool  `json:"default,omitempty"` // true, if the schedule type is default
}
type RequestReportsCreateOrScheduleAReport struct {
	Tags             []string                                           `json:"tags,omitempty"`             // array of tags for report
	Deliveries       *[]RequestReportsCreateOrScheduleAReportDeliveries `json:"deliveries,omitempty"`       // Array of available delivery channels
	Name             string                                             `json:"name,omitempty"`             // report name
	Schedule         *RequestReportsCreateOrScheduleAReportSchedule     `json:"schedule,omitempty"`         //
	View             *RequestReportsCreateOrScheduleAReportView         `json:"view,omitempty"`             //
	ViewGroupID      string                                             `json:"viewGroupId,omitempty"`      // viewGroupId of the viewgroup for the report
	ViewGroupVersion string                                             `json:"viewGroupVersion,omitempty"` // version of viewgroup for the report
}
type RequestReportsCreateOrScheduleAReportDeliveries interface{}
type RequestReportsCreateOrScheduleAReportSchedule interface{}
type RequestReportsCreateOrScheduleAReportView struct {
	FieldGroups *[]RequestReportsCreateOrScheduleAReportViewFieldGroups `json:"fieldGroups,omitempty"` //
	Filters     *[]RequestReportsCreateOrScheduleAReportViewFilters     `json:"filters,omitempty"`     //
	Format      *RequestReportsCreateOrScheduleAReportViewFormat        `json:"format,omitempty"`      //
	Name        string                                                  `json:"name,omitempty"`        // view name
	ViewID      string                                                  `json:"viewId,omitempty"`      // view Id
}
type RequestReportsCreateOrScheduleAReportViewFieldGroups struct {
	FieldGroupDisplayName string                                                        `json:"fieldGroupDisplayName,omitempty"` // Field group label/displayname for user
	FieldGroupName        string                                                        `json:"fieldGroupName,omitempty"`        // Field group name
	Fields                *[]RequestReportsCreateOrScheduleAReportViewFieldGroupsFields `json:"fields,omitempty"`                //
}
type RequestReportsCreateOrScheduleAReportViewFieldGroupsFields struct {
	DisplayName string `json:"displayName,omitempty"` // field label/displayname
	Name        string `json:"name,omitempty"`        // field name
}
type RequestReportsCreateOrScheduleAReportViewFilters struct {
	DisplayName string                                                 `json:"displayName,omitempty"` // filter label/displayname
	Name        string                                                 `json:"name,omitempty"`        // filter name
	Type        string                                                 `json:"type,omitempty"`        // filter type
	Value       *RequestReportsCreateOrScheduleAReportViewFiltersValue `json:"value,omitempty"`       // value of filter. data type is based on the filter type. Use the filter definitions from the view to fetch the options for a filter.
}
type RequestReportsCreateOrScheduleAReportViewFiltersValue interface{}
type RequestReportsCreateOrScheduleAReportViewFormat struct {
	FormatType string `json:"formatType,omitempty"` // format type of report
	Name       string `json:"name,omitempty"`       // format name of report
}

//GetListOfScheduledReports Get list of scheduled reports - 2ab4-b80d-49ca-ae42
/* Get list of scheduled report configurations.


@param GetListOfScheduledReportsQueryParams Filtering parameter
*/
func (s *ReportsService) GetListOfScheduledReports(GetListOfScheduledReportsQueryParams *GetListOfScheduledReportsQueryParams) (*ResponseReportsGetListOfScheduledReports, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports"

	queryString, _ := query.Values(GetListOfScheduledReportsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseReportsGetListOfScheduledReports{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetListOfScheduledReports")
	}

	result := response.Result().(*ResponseReportsGetListOfScheduledReports)
	return result, response, err

}

//GetAScheduledReport Get a scheduled report - b79a-3910-4e18-9251
/* Get scheduled report configuration by reportId


@param reportID reportId path parameter. reportId of report

*/
func (s *ReportsService) GetAScheduledReport(reportID string) (*ResponseReportsGetAScheduledReport, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAScheduledReport{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAScheduledReport")
	}

	result := response.Result().(*ResponseReportsGetAScheduledReport)
	return result, response, err

}

//GetAllExecutionDetailsForAGivenReport Get all execution details for a given report - 91b9-d830-4679-a273
/* Get details of all executions for a given report


@param reportID reportId path parameter. reportId of report

*/
func (s *ReportsService) GetAllExecutionDetailsForAGivenReport(reportID string) (*ResponseReportsGetAllExecutionDetailsForAGivenReport, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}/executions"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAllExecutionDetailsForAGivenReport{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllExecutionDetailsForAGivenReport")
	}

	result := response.Result().(*ResponseReportsGetAllExecutionDetailsForAGivenReport)
	return result, response, err

}

//DownloadReportContent Download report content - d6bb-ebd7-4a48-87bd
/* Returns report content. Save the response to a file by converting the response data as a blob and setting the file format available from content-disposition response header.


@param reportID reportId path parameter. reportId of report

@param executionID executionId path parameter. executionId of report execution

*/
func (s *ReportsService) DownloadReportContent(reportID string, executionID string) (FileDownload, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}/executions/{executionId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	return fdownload, response, err

}

//GetAllViewGroups Get all view groups - 2f90-4a35-44ab-b1c9
/* Gives a list of summary of all view groups.


 */
func (s *ReportsService) GetAllViewGroups() (*ResponseReportsGetAllViewGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/view-groups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetAllViewGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllViewGroups")
	}

	result := response.Result().(*ResponseReportsGetAllViewGroups)
	return result, response, err

}

//GetViewsForAGivenViewGroup Get views for a given view group - 03b6-aa2b-4dda-a555
/* Gives a list of summary of all views in a viewgroup. Use "Get all view groups" API to get the viewGroupIds (required as a query param for this API) for available viewgroups.


@param viewGroupID viewGroupId path parameter. viewGroupId of viewgroup.

*/
func (s *ReportsService) GetViewsForAGivenViewGroup(viewGroupID string) (*ResponseReportsGetViewsForAGivenViewGroup, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/view-groups/{viewGroupId}"
	path = strings.Replace(path, "{viewGroupId}", fmt.Sprintf("%v", viewGroupID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetViewsForAGivenViewGroup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetViewsForAGivenViewGroup")
	}

	result := response.Result().(*ResponseReportsGetViewsForAGivenViewGroup)
	return result, response, err

}

//GetViewDetailsForAGivenViewGroupView Get view details for a given view group & view - 1d9a-ba2f-4f89-ae51
/* Gives complete information of the view that is required to configure a report. Use "Get views for a given view group" API to get the viewIds  (required as a query param for this API) for available views.


@param viewGroupID viewGroupId path parameter. viewGroupId of viewgroup

@param viewID viewId path parameter. view id of view

*/
func (s *ReportsService) GetViewDetailsForAGivenViewGroupView(viewGroupID string, viewID string) (*ResponseReportsGetViewDetailsForAGivenViewGroupView, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/view-groups/{viewGroupId}/views/{viewId}"
	path = strings.Replace(path, "{viewGroupId}", fmt.Sprintf("%v", viewGroupID), -1)
	path = strings.Replace(path, "{viewId}", fmt.Sprintf("%v", viewID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsGetViewDetailsForAGivenViewGroupView{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetViewDetailsForAGivenViewGroupView")
	}

	result := response.Result().(*ResponseReportsGetViewDetailsForAGivenViewGroupView)
	return result, response, err

}

//CreateOrScheduleAReport Create or Schedule a report - 8abf-291a-42aa-8860
/* Create/Schedule a report configuration. Use "Get view details for a given view group & view" API to get the metadata required to configure a report.


 */
func (s *ReportsService) CreateOrScheduleAReport(requestReportsCreateOrScheduleAReport *RequestReportsCreateOrScheduleAReport) (*ResponseReportsCreateOrScheduleAReport, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestReportsCreateOrScheduleAReport).
		SetResult(&ResponseReportsCreateOrScheduleAReport{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrScheduleAReport")
	}

	result := response.Result().(*ResponseReportsCreateOrScheduleAReport)
	return result, response, err

}

//DeleteAScheduledReport Delete a scheduled report - 239c-6921-4f9b-b12e
/* Delete a scheduled report configuration. Deletes the report executions also.


@param reportID reportId path parameter. reportId of report

*/
func (s *ReportsService) DeleteAScheduledReport(reportID string) (*ResponseReportsDeleteAScheduledReport, *resty.Response, error) {
	path := "/dna/intent/api/v1/data/reports/{reportId}"
	path = strings.Replace(path, "{reportId}", fmt.Sprintf("%v", reportID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseReportsDeleteAScheduledReport{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteAScheduledReport")
	}

	result := response.Result().(*ResponseReportsDeleteAScheduledReport)
	return result, response, err

}
