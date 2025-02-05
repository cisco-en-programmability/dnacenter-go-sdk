package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplicationsService service

type RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams struct {
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy            string  `url:"sortBy,omitempty"`            //A field within the response to sort by.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` is mandatory. `siteId` must be a site UUID of a building. (Ex."buildingUuid") Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	SSID              string  `url:"ssid,omitempty"`              //In the context of a network application, SSID refers to the name of the wireless network to which the client connects. Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	ApplicationName   string  `url:"applicationName,omitempty"`   //Name of the application for which the experience data is intended. Examples: `applicationName=webex` (single applicationName requested) `applicationName=webex&applicationName=teams` (multiple applicationName requested)
	BusinessRelevance string  `url:"businessRelevance,omitempty"` //The application can be chosen to be categorized as business-relevant, irrelevant, or default (neutral). By doing so, the assurance application prioritizes the monitoring and analysis of business-relevant data, ensuring critical insights are captured. Applications marked as irrelevant or default are selectively excluded from certain data sets, streamlining focus on what's most important for business outcomes.
	Attribute         string  `url:"attribute,omitempty"`         //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Supported attributes are applicationName, siteId, exporterIpAddress, exporterNetworkDeviceId, healthScore, businessRelevance, usage, throughput, packetLossPercent, networkLatency, applicationServerLatency, clientNetworkLatency, serverNetworkLatency, trafficClass, jitter, ssid Examples: `attribute=healthScore` (single attribute requested) `attribute=healthScore&attribute=ssid&attribute=jitter` (multiple attribute requested)
}
type RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams struct {
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` is mandatory. `siteId` must be a site UUID of a building. (Ex."buildingUuid") Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	SSID              string  `url:"ssid,omitempty"`              //In the context of a network application, SSID refers to the name of the wireless network to which the client connects. Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	ApplicationName   string  `url:"applicationName,omitempty"`   //Name of the application for which the experience data is intended. Examples: `applicationName=webex` (single applicationName requested) `applicationName=webex&applicationName=teams` (multiple applicationName requested)
	BusinessRelevance string  `url:"businessRelevance,omitempty"` //The application can be chosen to be categorized as business-relevant, irrelevant, or default (neutral). By doing so, the assurance application prioritizes the monitoring and analysis of business-relevant data, ensuring critical insights are captured. Applications marked as irrelevant or default are selectively excluded from certain data sets, streamlining focus on what's most important for business outcomes.
}
type RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ApplicationsQueryParams struct {
	SiteID            string  `url:"siteId,omitempty"`            //Assurance site UUID value (Cannot be submitted together with deviceId and clientMac)
	DeviceID          string  `url:"deviceId,omitempty"`          //Assurance device UUID value (Cannot be submitted together with siteId and clientMac)
	MacAddress        string  `url:"macAddress,omitempty"`        //Client device's MAC address (Cannot be submitted together with siteId and deviceId)
	StartTime         float64 `url:"startTime,omitempty"`         //Starting epoch time in milliseconds of time window
	EndTime           float64 `url:"endTime,omitempty"`           //Ending epoch time in milliseconds of time window
	ApplicationHealth string  `url:"applicationHealth,omitempty"` //Application health category (POOR, FAIR, or GOOD.  Optionally use with siteId only)
	Offset            float64 `url:"offset,omitempty"`            //The offset of the first application in the returned data (optionally used with siteId only)
	Limit             float64 `url:"limit,omitempty"`             //The max number of application entries in returned data [1, 1000] (optionally used with siteId only)
	ApplicationName   string  `url:"applicationName,omitempty"`   //The name of the application to get information on
}

type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics struct {
	Response *[]ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsResponse `json:"response,omitempty"` //

	Page *ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsResponse struct {
	ID string `json:"id,omitempty"` // Id

	ApplicationName string `json:"applicationName,omitempty"` // Application Name

	BusinessRelevance string `json:"businessRelevance,omitempty"` // Business Relevance

	SiteID string `json:"siteId,omitempty"` // Site Id

	ExporterIPAddress string `json:"exporterIpAddress,omitempty"` // Exporter Ip Address

	ExporterNetworkDeviceID string `json:"exporterNetworkDeviceId,omitempty"` // Exporter Network Device Id

	HealthScore *int `json:"healthScore,omitempty"` // Health Score

	Usage *float64 `json:"usage,omitempty"` // Usage

	Throughput *float64 `json:"throughput,omitempty"` // Throughput

	PacketLossPercent *float64 `json:"packetLossPercent,omitempty"` // Packet Loss Percent

	NetworkLatency *float64 `json:"networkLatency,omitempty"` // Network Latency

	ApplicationServerLatency *float64 `json:"applicationServerLatency,omitempty"` // Application Server Latency

	ClientNetworkLatency *float64 `json:"clientNetworkLatency,omitempty"` // Client Network Latency

	ServerNetworkLatency *float64 `json:"serverNetworkLatency,omitempty"` // Server Network Latency

	TrafficClass string `json:"trafficClass,omitempty"` // Traffic Class

	Jitter *float64 `json:"jitter,omitempty"` // Jitter

	SSID string `json:"ssid,omitempty"` // Ssid
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering struct {
	Response *ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications struct {
	Response *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponse `json:"response,omitempty"` //

	Page *ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Groups *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroups `json:"groups,omitempty"` //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Cursor string `json:"cursor,omitempty"` // Cursor

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseApplicationsApplications struct {
	Version    string                                      `json:"version,omitempty"`    // API version
	TotalCount *int                                        `json:"totalCount,omitempty"` // Count of items in response
	Response   *[]ResponseApplicationsApplicationsResponse `json:"response,omitempty"`   //
}
type ResponseApplicationsApplicationsResponse struct {
	Name                     string                                                            `json:"name,omitempty"`                     // Application name
	Health                   *int                                                              `json:"health,omitempty"`                   // Health score
	BusinessRelevance        string                                                            `json:"businessRelevance,omitempty"`        // Application's business relevance
	TrafficClass             string                                                            `json:"trafficClass,omitempty"`             // Application's traffic class
	UsageBytes               *int                                                              `json:"usageBytes,omitempty"`               // Usage amount in bytes
	AverageThroughput        *float64                                                          `json:"averageThroughput,omitempty"`        // Average throughput of application
	PacketLossPercent        *ResponseApplicationsApplicationsResponsePacketLossPercent        `json:"packetLossPercent,omitempty"`        // Packet loss percentage for application
	NetworkLatency           *ResponseApplicationsApplicationsResponseNetworkLatency           `json:"networkLatency,omitempty"`           // Network latency for application
	Jitter                   *ResponseApplicationsApplicationsResponseJitter                   `json:"jitter,omitempty"`                   // Jitter for application
	ApplicationServerLatency *ResponseApplicationsApplicationsResponseApplicationServerLatency `json:"applicationServerLatency,omitempty"` // Latency of application server
	ClientNetworkLatency     *ResponseApplicationsApplicationsResponseClientNetworkLatency     `json:"clientNetworkLatency,omitempty"`     // Latency of client network
	ServerNetworkLatency     *ResponseApplicationsApplicationsResponseServerNetworkLatency     `json:"serverNetworkLatency,omitempty"`     // Latency of server network
	ExporterIPAddress        string                                                            `json:"exporterIpAddress,omitempty"`        // Ip address of exporter device
	ExporterName             string                                                            `json:"exporterName,omitempty"`             // Name of exporter device
	ExporterUUID             string                                                            `json:"exporterUUID,omitempty"`             // UUID of exporter device
	ExporterFamily           string                                                            `json:"exporterFamily,omitempty"`           // Devices family of exporter device
	ClientName               string                                                            `json:"clientName,omitempty"`               // Endpoint client name
	ClientIP                 string                                                            `json:"clientIp,omitempty"`                 // Endpoint client ip
	Location                 string                                                            `json:"location,omitempty"`                 // Site location
	OperatingSystem          string                                                            `json:"operatingSystem,omitempty"`          // Endpoint's operating system
	DeviceType               string                                                            `json:"deviceType,omitempty"`               // Type of device
	ClientMacAddress         string                                                            `json:"clientMacAddress,omitempty"`         // Endpoint mac address
	IssueID                  string                                                            `json:"issueId,omitempty"`                  // Id number of issue
	IssueName                string                                                            `json:"issueName,omitempty"`                // Name of issue
	Application              string                                                            `json:"application,omitempty"`              // Issue reltaed application
	Severity                 string                                                            `json:"severity,omitempty"`                 // Issue severity
	Summary                  string                                                            `json:"summary,omitempty"`                  // Issue summary
	RootCause                string                                                            `json:"rootCause,omitempty"`                // Issue's root cause
	Timestamp                *int                                                              `json:"timestamp,omitempty"`                // Issue's timestamp
	Occurrences              *int                                                              `json:"occurrences,omitempty"`              // Issue occurrences
	Priority                 string                                                            `json:"priority,omitempty"`                 // Issue priority
}
type ResponseApplicationsApplicationsResponsePacketLossPercent interface{}
type ResponseApplicationsApplicationsResponseNetworkLatency interface{}
type ResponseApplicationsApplicationsResponseJitter interface{}
type ResponseApplicationsApplicationsResponseApplicationServerLatency interface{}
type ResponseApplicationsApplicationsResponseClientNetworkLatency interface{}
type ResponseApplicationsApplicationsResponseServerNetworkLatency interface{}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage `json:"page,omitempty"` //
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value *int `json:"value,omitempty"` // Value
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Cursor string `json:"cursor,omitempty"` // Cursor

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}

//RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics Retrieves the list of network applications along with experience and health metrics - 0d8c-6945-49f8-845d
/* Retrieves the list of network applications along with experience and health metrics. If startTime and endTime are not provided, the API defaults to the last 24 hours. `siteId` is mandatory. `siteId` must be a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml


@param RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams Custom header parameters
@param RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-applications-along-with-experience-and-health-metrics
*/
func (s *ApplicationsService) RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams *RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams, RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams *RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams) (*ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications"

	queryString, _ := query.Values(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams != nil {

		if RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics(RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsHeaderParams, RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetrics)
	return result, response, err

}

//RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering Retrieves the total count of network applications by applying basic filtering - 90b0-5a40-422a-8446
/* Retrieves the number of network applications by applying basic filtering. If startTime and endTime are not provided, the API defaults to the last 24 hours. `siteId` is mandatory. `siteId` must be a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml


@param RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams Custom header parameters
@param RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-count-of-network-applications-by-applying-basic-filtering
*/
func (s *ApplicationsService) RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams *RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams, RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams *RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams) (*ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications/count"

	queryString, _ := query.Values(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams != nil {

		if RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering(RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringHeaderParams, RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering)
	return result, response, err

}

//Applications Applications - 2db5-8a1f-4fea-9242
/* Intent API to get a list of applications for a specific site, a device, or a client device's MAC address. For a combination of a specific application with site and/or device the API gets list of issues/devices/endpoints.


@param ApplicationsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!applications
*/
func (s *ApplicationsService) Applications(ApplicationsQueryParams *ApplicationsQueryParams) (*ResponseApplicationsApplications, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-health"

	queryString, _ := query.Values(ApplicationsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsApplications{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Applications(ApplicationsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Applications")
	}

	result := response.Result().(*ResponseApplicationsApplications)
	return result, response, err

}

//RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications Retrieves the Trend analytics data related to network applications. - 22b6-58d5-41ba-ac1a
/* Retrieves the trend analytics of applications experience data for the specified time range. The data will be grouped based on the given trend time interval. This API facilitates obtaining consolidated insights into the performance and status of the network applications over the specified start and end time. If startTime and endTime are not provided, the API defaults to the last 24 hours. `siteId` and `trendInterval` are mandatory. `siteId` must be a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml


@param RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-trend-analytics-data-related-to-network-applications
*/
func (s *ApplicationsService) RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications *RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications, RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams *RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams) (*ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams != nil {

		if RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications).
		SetResult(&ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications(requestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications, RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications)
	return result, response, err

}
