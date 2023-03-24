package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type HealthAndPerformanceService service

type SystemHealthApIQueryParams struct {
	Summary   bool    `url:"summary,omitempty"`   //Fetch the latest high severity event
	Domain    string  `url:"domain,omitempty"`    //Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Subdomain string  `url:"subdomain,omitempty"` //Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Limit     float64 `url:"limit,omitempty"`     //limit
	Offset    float64 `url:"offset,omitempty"`    //offset
}
type SystemHealthCountApIQueryParams struct {
	Domain    string `url:"domain,omitempty"`    //Fetch system events with this domain. Possible values of domain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
	Subdomain string `url:"subdomain,omitempty"` //Fetch system events with this subdomain. Possible values of subdomain are listed here : /dna/platform/app/consumer-portal/developer-toolkit/events
}
type SystemPerformanceApIQueryParams struct {
	Kpi       string  `url:"kpi,omitempty"`       //Valid values: cpu,memory,network
	Function  string  `url:"function,omitempty"`  //Valid values: sum,average,max
	StartTime float64 `url:"startTime,omitempty"` //This is the epoch start time in milliseconds from which performance indicator need to be fetched
	EndTime   float64 `url:"endTime,omitempty"`   //This is the epoch end time in milliseconds upto which performance indicator need to be fetched
}
type SystemPerformanceHistoricalApIQueryParams struct {
	Kpi       string  `url:"kpi,omitempty"`       //Fetch historical data for this kpi. Valid values: cpu,memory,network
	StartTime float64 `url:"startTime,omitempty"` //This is the epoch start time in milliseconds from which performance indicator need to be fetched
	EndTime   float64 `url:"endTime,omitempty"`   //This is the epoch end time in milliseconds upto which performance indicator need to be fetched
}

type ResponseHealthAndPerformanceSystemHealthApI struct {
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
type ResponseHealthAndPerformanceSystemHealthCountApI struct {
	Count *float64 `json:"count,omitempty"` // Count of the events
}
type ResponseHealthAndPerformanceSystemPerformanceApI struct {
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
type ResponseHealthAndPerformanceSystemPerformanceHistoricalApI struct {
	HostName string                                                          `json:"hostName,omitempty"` // Hostname
	Version  string                                                          `json:"version,omitempty"`  // Version of the API
	Kpis     *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpis `json:"kpis,omitempty"`     //
}
type ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpis struct {
	Legends *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisLegends `json:"legends,omitempty"` //
	Data    *ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisData    `json:"data,omitempty"`    //
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

//SystemHealthApI System Health API - 6085-eb1b-4f48-9740
/* This API retrieves the latest system events


@param SystemHealthAPIQueryParams Filtering parameter
*/
func (s *HealthAndPerformanceService) SystemHealthApI(SystemHealthAPIQueryParams *SystemHealthApIQueryParams) (*ResponseHealthAndPerformanceSystemHealthApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/health"

	queryString, _ := query.Values(SystemHealthAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemHealthApI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SystemHealthApI")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemHealthApI)
	return result, response, err

}

//SystemHealthCountApI System Health Count API - 5289-0891-4729-8714
/* This API gives the count of the latest system events


@param SystemHealthCountAPIQueryParams Filtering parameter
*/
func (s *HealthAndPerformanceService) SystemHealthCountApI(SystemHealthCountAPIQueryParams *SystemHealthCountApIQueryParams) (*ResponseHealthAndPerformanceSystemHealthCountApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/health/count"

	queryString, _ := query.Values(SystemHealthCountAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemHealthCountApI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SystemHealthCountApI")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemHealthCountApI)
	return result, response, err

}

//SystemPerformanceApI System Performance API - f2a9-5b4d-48eb-a4f8
/* This API gives the aggregated performance indicators. The data can be retrieved for the last 3 months.


@param SystemPerformanceAPIQueryParams Filtering parameter
*/
func (s *HealthAndPerformanceService) SystemPerformanceApI(SystemPerformanceAPIQueryParams *SystemPerformanceApIQueryParams) (*ResponseHealthAndPerformanceSystemPerformanceApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/performance"

	queryString, _ := query.Values(SystemPerformanceAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemPerformanceApI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SystemPerformanceApI")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemPerformanceApI)
	return result, response, err

}

//SystemPerformanceHistoricalApI System Performance Historical API - 879b-ea1e-4389-83d7
/* This API retrieves the historical performance indicators . The data can be retrieved for the last 3 months.


@param SystemPerformanceHistoricalAPIQueryParams Filtering parameter
*/
func (s *HealthAndPerformanceService) SystemPerformanceHistoricalApI(SystemPerformanceHistoricalAPIQueryParams *SystemPerformanceHistoricalApIQueryParams) (*ResponseHealthAndPerformanceSystemPerformanceHistoricalApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/diagnostics/system/performance/history"

	queryString, _ := query.Values(SystemPerformanceHistoricalAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseHealthAndPerformanceSystemPerformanceHistoricalApI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SystemPerformanceHistoricalApI")
	}

	result := response.Result().(*ResponseHealthAndPerformanceSystemPerformanceHistoricalApI)
	return result, response, err

}
