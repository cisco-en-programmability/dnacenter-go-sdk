package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type HealthAndPerformanceService service

type RetrievesAllTheValidationSetsQueryParams struct {
	View string `url:"view,omitempty"` //When the query parameter `view=DETAIL` is passed, all validation sets and associated validations will be returned. When the query parameter `view=DEFAULT` is passed, only validation sets metadata will be returned.
}
type RetrievesTheListOfValidationWorkflowsQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Workflows started after the given time (as milliseconds since UNIX epoch).
	EndTime   float64 `url:"endTime,omitempty"`   //Workflows started before the given time (as milliseconds since UNIX epoch).
	RunStatus string  `url:"runStatus,omitempty"` //Execution status of the workflow. If the workflow is successfully submitted, runStatus is `PENDING`. If the workflow execution has started, runStatus is `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus is `COMPLETED`. If the workflow execution fails while running validations, runStatus is `FAILED`.
	Offset    float64 `url:"offset,omitempty"`    //The first record to show for this page; the first record is numbered 1.
	Limit     float64 `url:"limit,omitempty"`     //The number of records to show for this page.
}
type RetrievesTheCountOfValidationWorkflowsQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Workflows started after the given time (as milliseconds since UNIX epoch).
	EndTime   float64 `url:"endTime,omitempty"`   //Workflows started before the given time (as milliseconds since UNIX epoch).
	RunStatus string  `url:"runStatus,omitempty"` //Execution status of the workflow. If the workflow is successfully submitted, runStatus is `PENDING`. If the workflow execution has started, runStatus is `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus is `COMPLETED`. If the workflow execution fails while running validations, runStatus is `FAILED`.
}
type SystemHealthAPIQueryParams struct {
	Summary   bool    `url:"summary,omitempty"`   //Fetch the latest high severity event
	Domain    string  `url:"domain,omitempty"`    //Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Subdomain string  `url:"subdomain,omitempty"` //Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Limit     float64 `url:"limit,omitempty"`     //limit
	Offset    float64 `url:"offset,omitempty"`    //offset
}
type SystemHealthCountAPIQueryParams struct {
	Domain    string `url:"domain,omitempty"`    //Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Subdomain string `url:"subdomain,omitempty"` //Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
}
type SystemPerformanceAPIQueryParams struct {
	Kpi       string  `url:"kpi,omitempty"`       //Valid values: cpu,memory,network
	Function  string  `url:"function,omitempty"`  //Valid values: sum,average,max
	StartTime float64 `url:"startTime,omitempty"` //This is the epoch start time in milliseconds from which performance indicator need to be fetched
	EndTime   float64 `url:"endTime,omitempty"`   //This is the epoch end time in milliseconds upto which performance indicator need to be fetched
}
type SystemPerformanceHistoricalAPIQueryParams struct {
	Kpi       string  `url:"kpi,omitempty"`       //Fetch historical data for this kpi. Valid values: cpu,memory,network
	StartTime float64 `url:"startTime,omitempty"` //This is the epoch start time in milliseconds from which performance indicator need to be fetched
	EndTime   float64 `url:"endTime,omitempty"`   //This is the epoch end time in milliseconds upto which performance indicator need to be fetched
}

type ResponseHealthAndPerformanceRetrievesAllTheValidationSets struct {
	Response *[]ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponse `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponse struct {
	ID               string                                                                               `json:"id,omitempty"`               // Validation set id
	Name             string                                                                               `json:"name,omitempty"`             // Name of the validation set
	Description      string                                                                               `json:"description,omitempty"`      // Description of the validation set
	Version          string                                                                               `json:"version,omitempty"`          // Version of the validation set
	ValidationGroups *[]ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponseValidationGroups `json:"validationGroups,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponseValidationGroups struct {
	Name        string                                                                                          `json:"name,omitempty"`        // Name of the validation group
	ID          string                                                                                          `json:"id,omitempty"`          // Validation group id
	Description string                                                                                          `json:"description,omitempty"` // Description of the validation group
	Validations *[]ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponseValidationGroupsValidations `json:"validations,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesAllTheValidationSetsResponseValidationGroupsValidations struct {
	ID   string `json:"id,omitempty"`   // Validation id
	Name string `json:"name,omitempty"` // Name of the validation
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSet struct {
	Response *ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponse `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponse struct {
	ID               string                                                                                             `json:"id,omitempty"`               // Validation set id
	Name             string                                                                                             `json:"name,omitempty"`             // Name of the validation set
	Description      string                                                                                             `json:"description,omitempty"`      // Description of the validation set
	Version          string                                                                                             `json:"version,omitempty"`          // Version of validation set
	ValidationGroups *[]ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponseValidationGroups `json:"validationGroups,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponseValidationGroups struct {
	Name        string                                                                                                        `json:"name,omitempty"`        // Name of the validation group
	ID          string                                                                                                        `json:"id,omitempty"`          // Validation group id
	Description string                                                                                                        `json:"description,omitempty"` // Description of the validation group
	Validations *[]ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponseValidationGroupsValidations `json:"validations,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSetResponseValidationGroupsValidations struct {
	ID   string `json:"id,omitempty"`   // Validation id
	Name string `json:"name,omitempty"` // Name of the validation
}
type ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflows struct {
	Response *[]ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflowsResponse struct {
	ID               string   `json:"id,omitempty"`               // Workflow id
	Name             string   `json:"name,omitempty"`             // Workflow name
	Description      string   `json:"description,omitempty"`      // Workflow description
	RunStatus        string   `json:"runStatus,omitempty"`        // Execution status of the workflow. If the workflow is successfully submitted, runStatus will return `PENDING`. If the workflow execution has started, runStatus will return `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus will return `COMPLETED`. If the workflow execution fails while running validations, runStatus will return `FAILED`.
	SubmitTime       *int     `json:"submitTime,omitempty"`       // Workflow submit time (as milliseconds since UNIX epoch).
	StartTime        *int     `json:"startTime,omitempty"`        // Workflow start time (as milliseconds since UNIX epoch).
	EndTime          *int     `json:"endTime,omitempty"`          // Workflow finish time (as milliseconds since UNIX epoch).
	ValidationStatus string   `json:"validationStatus,omitempty"` // Overall result of execution of the validation workflow
	ValidationSetIDs []string `json:"validationSetIds,omitempty"` // List of validation set ids
}
type ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations struct {
	Response *ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsResponse `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidationsResponse struct {
	ID string `json:"id,omitempty"` // UUID of the workflow submitted for executing validations
}
type ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflows struct {
	Response *ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflowsResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetails struct {
	Response *ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // The version of the response
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponse struct {
	ID                       string                                                                                            `json:"id,omitempty"`                       // Workflow id
	Name                     string                                                                                            `json:"name,omitempty"`                     // Workflow name
	Description              string                                                                                            `json:"description,omitempty"`              // Workflow description
	RunStatus                string                                                                                            `json:"runStatus,omitempty"`                // Execution status of the workflow. If the workflow is successfully submitted, runStatus will return `PENDING`. If the workflow execution has started, runStatus will return `IN_PROGRESS`. If the workflow executed is completed with all validations executed, runStatus will return `COMPLETED`. If the workflow execution fails while running validations, runStatus will return `FAILED`.
	SubmitTime               *int                                                                                              `json:"submitTime,omitempty"`               // Workflow submit time (as milliseconds since UNIX epoch).
	ValidationSetIDs         []string                                                                                          `json:"validationSetIds,omitempty"`         // List of validation set ids
	ReleaseVersion           string                                                                                            `json:"releaseVersion,omitempty"`           // Product version
	ValidationSetsRunDetails *[]ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponseValidationSetsRunDetails `json:"validationSetsRunDetails,omitempty"` //
	ValidationStatus         string                                                                                            `json:"validationStatus,omitempty"`         // Overall result of the execution of all the validations. If any of the contained validation execution status is `CRITICAL`, this is marked as `CRITICAL`. Else, if any of the contained validation execution status is `WARNING`, this is marked as `WARNING`. Else, this is marked as `INFORMATION`.
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponseValidationSetsRunDetails struct {
	ValidationSetID      string                                                                                                                `json:"validationSetId,omitempty"`      // Validation set id
	StartTime            *int                                                                                                                  `json:"startTime,omitempty"`            // Validation set run start time (as milliseconds since UNIX epoch).
	EndTime              *int                                                                                                                  `json:"endTime,omitempty"`              // Validation set run finish time (as milliseconds since UNIX epoch).
	ValidationStatus     string                                                                                                                `json:"validationStatus,omitempty"`     // Overall result of the validation set execution. If any of the contained validation execution status is `CRITICAL`, this is marked as `CRITICAL`. Else, if any of the contained validation execution status is `WARNING`, this is marked as `WARNING`. Else, this is marked as `INFORMATION`. This is empty when the workflow is in progress.
	Version              string                                                                                                                `json:"version,omitempty"`              // Validation set version
	ValidationRunDetails *[]ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponseValidationSetsRunDetailsValidationRunDetails `json:"validationRunDetails,omitempty"` //
}
type ResponseHealthAndPerformanceRetrievesValidationWorkflowDetailsResponseValidationSetsRunDetailsValidationRunDetails struct {
	ValidationID      string `json:"validationId,omitempty"`      // Validation id
	ValidationName    string `json:"validationName,omitempty"`    // Validation name
	ValidationMessage string `json:"validationMessage,omitempty"` // Validation execution result detail message
	ValidationStatus  string `json:"validationStatus,omitempty"`  // Validation execution result status
}
type ResponseHealthAndPerformanceSystemHealthAPI struct {
	HealthEvents *[]ResponseHealthAndPerformanceSystemHealthAPIHealthEvents `json:"healthEvents,omitempty"` //
	Version      string                                                     `json:"version,omitempty"`      // API version
	HostName     string                                                     `json:"hostName,omitempty"`     // Cluster name
	Cimcaddress  []string                                                   `json:"cimcaddress,omitempty"`  // List of configured cimc addresse(s)
}
type ResponseHealthAndPerformanceSystemHealthAPIHealthEvents struct {
	Severity    string `json:"severity,omitempty"`    // Severity of the event
	Hostname    string `json:"hostname,omitempty"`    // Hostname of the event
	Instance    string `json:"instance,omitempty"`    // Instance of the event
	SubDomain   string `json:"subDomain,omitempty"`   // Sub domain of the event
	Domain      string `json:"domain,omitempty"`      // Domain of the event
	Description string `json:"description,omitempty"` // Details of the event
	State       string `json:"state,omitempty"`       // State of the event
	Timestamp   string `json:"timestamp,omitempty"`   // Time of the event occurance
	Status      string `json:"status,omitempty"`      // Event status
}
type ResponseHealthAndPerformanceSystemHealthCountAPI struct {
	Count *float64 `json:"count,omitempty"` // Count of the events
}
type ResponseHealthAndPerformanceSystemPerformanceAPI struct {
	HostName string                                                `json:"hostName,omitempty"` // Hostname
	Version  string                                                `json:"version,omitempty"`  // Version of the API
	Kpis     *ResponseHealthAndPerformanceSystemPerformanceAPIKpis `json:"kpis,omitempty"`     //
}
type ResponseHealthAndPerformanceSystemPerformanceAPIKpis struct {
	CPU           *ResponseHealthAndPerformanceSystemPerformanceAPIKpisCPU           `json:"cpu,omitempty"`             //
	Memory        *ResponseHealthAndPerformanceSystemPerformanceAPIKpisMemory        `json:"memory,omitempty"`          //
	NetworktxRate *ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworktxRate `json:"network tx_rate,omitempty"` //
	NetworkrxRate *ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworkrxRate `json:"network rx_rate,omitempty"` //
}
type ResponseHealthAndPerformanceSystemPerformanceAPIKpisCPU struct {
	Units       string `json:"units,omitempty"`       // Units for cpu usage
	Utilization string `json:"utilization,omitempty"` // cpu usage in units
}
type ResponseHealthAndPerformanceSystemPerformanceAPIKpisMemory struct {
	Units       string `json:"units,omitempty"`       // Units for memory usage
	Utilization string `json:"utilization,omitempty"` // Memory usage in units
}
type ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworktxRate struct {
	Units       string `json:"units,omitempty"`       // Units for network tx_rate
	Utilization string `json:"utilization,omitempty"` // Network tx_rate in units
}
type ResponseHealthAndPerformanceSystemPerformanceAPIKpisNetworkrxRate struct {
	Units       string `json:"units,omitempty"`       // Units for network rx_rate
	Utilization string `json:"utilization,omitempty"` // Network rx_rate in units
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPI struct {
	HostName string                                                          `json:"hostName,omitempty"` // Hostname
	Version  string                                                          `json:"version,omitempty"`  // Version of the API
	Kpis     *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpis `json:"kpis,omitempty"`     //
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpis struct {
	Legends   *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegends `json:"legends,omitempty"`   //
	Data      *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisData    `json:"data,omitempty"`      //
	CPUAvg    string                                                                 `json:"cpuAvg,omitempty"`    // CPU average utilization
	MemoryAvg string                                                                 `json:"memoryAvg,omitempty"` // Memory average utilization
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegends struct {
	CPU           *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsCPU           `json:"cpu,omitempty"`             //
	Memory        *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsMemory        `json:"memory,omitempty"`          //
	NetworktxRate *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworktxRate `json:"network tx_rate,omitempty"` //
	NetworkrxRate *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworkrxRate `json:"network rx_rate,omitempty"` //
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsCPU struct {
	Units string `json:"units,omitempty"` // Units for cpu usage
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsMemory struct {
	Units string `json:"units,omitempty"` // Units for memory usage
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworktxRate struct {
	Units string `json:"units,omitempty"` // Units for network tx_rate
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegendsNetworkrxRate struct {
	Units string `json:"units,omitempty"` // Units for network rx_rate
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisData map[string][]string // Time in 'YYYY-MM-DDT00:00:00Z' format with values for legends
type RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations struct {
	Name             string   `json:"name,omitempty"`             // Name of the workflow to run. It must be unique.
	Description      string   `json:"description,omitempty"`      // Description of the workflow to run
	ValidationSetIDs []string `json:"validationSetIds,omitempty"` // List of validation set ids
}

//RetrievesAllTheValidationSets Retrieves all the validation sets - 11bb-4b03-4059-a001
/* Retrieves all the validation sets and optionally the contained validations


@param RetrievesAllTheValidationSetsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-all-the-validation-sets-v1
*/
func (s *HealthAndPerformanceService) RetrievesAllTheValidationSets(RetrievesAllTheValidationSetsQueryParams *RetrievesAllTheValidationSetsQueryParams) (*ResponseHealthAndPerformanceRetrievesAllTheValidationSets, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationSets"

	queryString, _ := query.Values(RetrievesAllTheValidationSetsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceRetrievesAllTheValidationSets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesAllTheValidationSets(RetrievesAllTheValidationSetsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesAllTheValidationSets")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesAllTheValidationSets)
	return result, response, err

}

//RetrievesValidationDetailsForAValidationSet Retrieves validation details for a validation set - 37b7-88bd-47b9-8533
/* Retrieves validation details for the given validation set id


@param id id path parameter. Validation set id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-validation-details-for-a-validation-set-v1
*/
func (s *HealthAndPerformanceService) RetrievesValidationDetailsForAValidationSet(id string) (*ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSet, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationSets/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSet{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesValidationDetailsForAValidationSet(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesValidationDetailsForAValidationSet")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesValidationDetailsForAValidationSet)
	return result, response, err

}

//RetrievesTheListOfValidationWorkflows Retrieves the list of validation workflows - 0fab-cafd-440b-98f8
/* Retrieves the workflows that have been successfully submitted and are currently available. This is sorted by `submitTime`


@param RetrievesTheListOfValidationWorkflowsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-validation-workflows-v1
*/
func (s *HealthAndPerformanceService) RetrievesTheListOfValidationWorkflows(RetrievesTheListOfValidationWorkflowsQueryParams *RetrievesTheListOfValidationWorkflowsQueryParams) (*ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflows, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows"

	queryString, _ := query.Values(RetrievesTheListOfValidationWorkflowsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflows{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfValidationWorkflows(RetrievesTheListOfValidationWorkflowsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfValidationWorkflows")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesTheListOfValidationWorkflows)
	return result, response, err

}

//RetrievesTheCountOfValidationWorkflows Retrieves the count of validation workflows - 4f91-8ac9-44c9-baef
/* Retrieves the count of workflows that have been successfully submitted and are currently available.


@param RetrievesTheCountOfValidationWorkflowsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-validation-workflows-v1
*/
func (s *HealthAndPerformanceService) RetrievesTheCountOfValidationWorkflows(RetrievesTheCountOfValidationWorkflowsQueryParams *RetrievesTheCountOfValidationWorkflowsQueryParams) (*ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflows, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows/count"

	queryString, _ := query.Values(RetrievesTheCountOfValidationWorkflowsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflows{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfValidationWorkflows(RetrievesTheCountOfValidationWorkflowsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfValidationWorkflows")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesTheCountOfValidationWorkflows)
	return result, response, err

}

//RetrievesValidationWorkflowDetails Retrieves validation workflow details - eb8b-eaad-451a-9c09
/* Retrieves workflow details for a workflow id


@param id id path parameter. Workflow id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-validation-workflow-details-v1
*/
func (s *HealthAndPerformanceService) RetrievesValidationWorkflowDetails(id string) (*ResponseHealthAndPerformanceRetrievesValidationWorkflowDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseHealthAndPerformanceRetrievesValidationWorkflowDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesValidationWorkflowDetails(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesValidationWorkflowDetails")
	}

	result := response.Result().(*ResponseHealthAndPerformanceRetrievesValidationWorkflowDetails)
	return result, response, err

}

//SystemHealthAPI System Health API - 6085-eb1b-4f48-9740
/* This API retrieves the latest system events


@param SystemHealthAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-health-api-v1
*/
func (s *HealthAndPerformanceService) SystemHealthAPI(SystemHealthAPIQueryParams *SystemHealthAPIQueryParams) (*ResponseHealthAndPerformanceSystemHealthAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/health"

	queryString, _ := query.Values(SystemHealthAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemHealthAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemHealthAPI(SystemHealthAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemHealthApi")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemHealthAPI)
	return result, response, err

}

//SystemHealthCountAPI System Health Count API - 5289-0891-4729-8714
/* This API gives the count of the latest system events


@param SystemHealthCountAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-health-count-api-v1
*/
func (s *HealthAndPerformanceService) SystemHealthCountAPI(SystemHealthCountAPIQueryParams *SystemHealthCountAPIQueryParams) (*ResponseHealthAndPerformanceSystemHealthCountAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/health/count"

	queryString, _ := query.Values(SystemHealthCountAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemHealthCountAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemHealthCountAPI(SystemHealthCountAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemHealthCountApi")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemHealthCountAPI)
	return result, response, err

}

//SystemPerformanceAPI System Performance API - f2a9-5b4d-48eb-a4f8
/* Retrieves the aggregated metrics (total, average or maximum) of cluster key performance indicators (KPIs), such as CPU utilization, memory utilization or network rates recorded within a specified time period. The data will be available from the past 24 hours.


@param SystemPerformanceAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-performance-api-v1
*/
func (s *HealthAndPerformanceService) SystemPerformanceAPI(SystemPerformanceAPIQueryParams *SystemPerformanceAPIQueryParams) (*ResponseHealthAndPerformanceSystemPerformanceAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/performance"

	queryString, _ := query.Values(SystemPerformanceAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemPerformanceAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemPerformanceAPI(SystemPerformanceAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemPerformanceApi")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemPerformanceAPI)
	return result, response, err

}

//SystemPerformanceHistoricalAPI System Performance Historical API - 879b-ea1e-4389-83d7
/* Retrieves the average values of cluster key performance indicators (KPIs), like CPU utilization, memory utilization or network rates grouped by time intervals within a specified time range. The data will be available from the past 24 hours.


@param SystemPerformanceHistoricalAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!system-performance-historical-api-v1
*/
func (s *HealthAndPerformanceService) SystemPerformanceHistoricalAPI(SystemPerformanceHistoricalAPIQueryParams *SystemPerformanceHistoricalAPIQueryParams) (*ResponseHealthAndPerformanceSystemPerformanceHistoricalAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/performance/history"

	queryString, _ := query.Values(SystemPerformanceHistoricalAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemPerformanceHistoricalAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SystemPerformanceHistoricalAPI(SystemPerformanceHistoricalAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SystemPerformanceHistoricalApi")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemPerformanceHistoricalAPI)
	return result, response, err

}

//SubmitsTheWorkflowForExecutingValidations Submits the workflow for executing validations - 52a0-7981-41ab-81d8
/* Submits the workflow for executing the validations for the given validation specifications



Documentation Link: https://developer.cisco.com/docs/dna-center/#!submits-the-workflow-for-executing-validations-v1
*/
func (s *HealthAndPerformanceService) SubmitsTheWorkflowForExecutingValidations(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations *RequestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations) (*ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations).
		SetResult(&ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitsTheWorkflowForExecutingValidations(requestHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations)
		}

		return nil, response, fmt.Errorf("error with operation SubmitsTheWorkflowForExecutingValidations")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSubmitsTheWorkflowForExecutingValidations)
	return result, response, err

}

//DeletesAValidationWorkflow Deletes a validation workflow - 8eb3-3959-47fa-9c50
/* Deletes the workflow for the given id


@param id id path parameter. Workflow id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-validation-workflow-v1
*/
func (s *HealthAndPerformanceService) DeletesAValidationWorkflow(id string) (*resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/diagnosticValidationWorkflows/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAValidationWorkflow(
				id)
		}
		return response, fmt.Errorf("error with operation DeletesAValidationWorkflow")
	}

	return response, err

}
