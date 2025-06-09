package dnac

import (
	"fmt"
	"net/http"
	"strings"

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
type RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsQueryParams struct {
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` must be a site UUID of a building. The list of buildings can be fetched using API `GET /dna/intent/api/v1/sites?type=building`. Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TestID            string  `url:"testId,omitempty"`            //Unique identifier of the ThousandEyes test. Examples: `testId=2043918` (filter for single testId) `testId=2043918&testId=129440` (filter for multiple testIds)
	TestName          string  `url:"testName,omitempty"`          //Name of the ThousandEyes test. This supports `*` wildcard, and filtering is case-insensitve. Examples: `testName=Cisco Webex` (exact match) `testName=Microsoft*` (starts with given string)
	TestType          string  `url:"testType,omitempty"`          //Type of the ThousandEyes test. Please note that Catalyst Center supports only a subset of all possible ThousandEyes test types.
	AgentID           string  `url:"agentId,omitempty"`           //Unique identifier of the ThousandEyes agent. Examples: `agentId=199345` (filter for single agentId) `agentId=1993458&agentId=499387` (filter for multiple agentIds)
	NetworkDeviceName string  `url:"networkDeviceName,omitempty"` //Name of the network device as per the inventory. This supports `*` wildcard, and filtering is case-insensitve.
	Attribute         string  `url:"attribute,omitempty"`         //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Examples: `attribute=testName` (single attribute requested) `attribute=testId&attribute=testName&attribute=averageLatency` (multiple attributes requested) . For valid attributes, verify the documentation.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy            string  `url:"sortBy,omitempty"`            //Attribute name by which the results should be sorted
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
}
type RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalCountOfThousandEyesTestResultsQueryParams struct {
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` must be a site UUID of a building. The list of buildings can be fetched using API `GET /dna/intent/api/v1/sites?type=building`. Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TestID            string  `url:"testId,omitempty"`            //Unique identifier of the ThousandEyes test. Examples: `testId=2043918` (filter for single testId) `testId=2043918&testId=129440` (filter for multiple testIds)
	TestName          string  `url:"testName,omitempty"`          //Name of the ThousandEyes test. This supports `*` wildcard, and filtering is case-insensitve. Examples: `testName=Cisco Webex` (exact match) `testName=Microsoft*` (starts with given string)
	TestType          string  `url:"testType,omitempty"`          //Type of the ThousandEyes test. Please note that Catalyst Center supports only a subset of all possible ThousandEyes test types.
	AgentID           string  `url:"agentId,omitempty"`           //Unique identifier of the ThousandEyes agent. Examples: `agentId=199345` (filter for single agentId) `agentId=1993458&agentId=499387` (filter for multiple agentIds)
	NetworkDeviceName string  `url:"networkDeviceName,omitempty"` //Name of the network device as per the inventory. This supports `*` wildcard, and filtering is case-insensitve.
}
type RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeQueryParams struct {
	SiteID            string  `url:"siteId,omitempty"`            //The site UUID without the top level hierarchy. `siteId` must be a site UUID of a building. The list of buildings can be fetched using API `GET /dna/intent/api/v1/sites?type=building`. Examples: `siteId=buildingUuid` (single siteId requested) `siteId=buildingUuid1&siteId=buildingUuid2` (multiple siteId requested)
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TrendInterval     string  `url:"trendInterval,omitempty"`     //The time interval to aggregate the metrics. Recommendation: |Time duration |Recommended `trendInterval`| |--------------|---------------------------| |Up to 6 hr    | `5MIN`                    | |6 hr to 2 days| `1HR`                     | |More than 2 days| `3HR`                     |
	TestID            string  `url:"testId,omitempty"`            //Unique identifier of the ThousandEyes test. Examples: `testId=2043918` (filter for single testId) `testId=2043918&testId=129440` (filter for multiple testIds)
	TestName          string  `url:"testName,omitempty"`          //Name of the ThousandEyes test. This supports `*` wildcard, and filtering is case-insensitve. Examples: `testName=Cisco Webex` (exact match) `testName=Microsoft*` (starts with given string)
	TestType          string  `url:"testType,omitempty"`          //Type of the ThousandEyes test. Please note that Catalyst Center supports only a subset of all possible ThousandEyes test types.
	AgentID           string  `url:"agentId,omitempty"`           //Unique identifier of the ThousandEyes agent. Examples: `agentId=199345` (filter for single agentId) `agentId=1993458&agentId=499387` (filter for multiple agentIds)
	NetworkDeviceName string  `url:"networkDeviceName,omitempty"` //Name of the network device as per the inventory. This supports `*` wildcard, and filtering is case-insensitve.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
}
type TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
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
	Page     *ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPage       `json:"page,omitempty"`     //
	Version  string                                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsResponse struct {
	ID                       string   `json:"id,omitempty"`                       // Id
	ApplicationName          string   `json:"applicationName,omitempty"`          // Application Name
	BusinessRelevance        string   `json:"businessRelevance,omitempty"`        // Business Relevance
	SiteID                   string   `json:"siteId,omitempty"`                   // Site Id
	ExporterIPAddress        string   `json:"exporterIpAddress,omitempty"`        // Exporter Ip Address
	ExporterNetworkDeviceID  string   `json:"exporterNetworkDeviceId,omitempty"`  // Exporter Network Device Id
	HealthScore              *int     `json:"healthScore,omitempty"`              // Health Score
	Usage                    *float64 `json:"usage,omitempty"`                    // Usage
	Throughput               *float64 `json:"throughput,omitempty"`               // Throughput
	PacketLossPercent        *float64 `json:"packetLossPercent,omitempty"`        // Packet Loss Percent
	NetworkLatency           *float64 `json:"networkLatency,omitempty"`           // Network Latency
	ApplicationServerLatency *float64 `json:"applicationServerLatency,omitempty"` // Application Server Latency
	ClientNetworkLatency     *float64 `json:"clientNetworkLatency,omitempty"`     // Client Network Latency
	ServerNetworkLatency     *float64 `json:"serverNetworkLatency,omitempty"`     // Server Network Latency
	TrafficClass             string   `json:"trafficClass,omitempty"`             // Traffic Class
	Jitter                   *float64 `json:"jitter,omitempty"`                   // Jitter
	SSID                     string   `json:"ssid,omitempty"`                     // Ssid
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPage struct {
	Limit  *int                                                                                                      `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                      `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseApplicationsRetrievesTheListOfNetworkApplicationsAlongWithExperienceAndHealthMetricsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFiltering struct {
	Response *ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringResponse `json:"response,omitempty"` //
	Version  string                                                                                           `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesTheTotalCountOfNetworkApplicationsByApplyingBasicFilteringResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics struct {
	Response *ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsResponse `json:"response,omitempty"` //
	Page     *ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPage     `json:"page,omitempty"`     //
	Version  string                                                                                                       `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsResponse struct {
	Attributes          *[]ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPage struct {
	Limit  *int                                                                                                             `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                             `json:"offset,omitempty"` // Offset
	Cursor string                                                                                                           `json:"cursor,omitempty"` // Cursor
	SortBy *[]ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications struct {
	Response *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponse `json:"response,omitempty"` //
	Page     *ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage       `json:"page,omitempty"`     //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponse struct {
	Timestamp           *int                                                                                                         `json:"timestamp,omitempty"`           // Timestamp
	Attributes          *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Groups              *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroups              `json:"groups,omitempty"`              //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    string `json:"value,omitempty"`    // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroups struct {
	ID                  string                                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    string `json:"value,omitempty"`    // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication struct {
	Response *[]ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationResponse `json:"response,omitempty"` //
	Page     *ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationPage       `json:"page,omitempty"`     //
	Version  string                                                                                       `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationResponse struct {
	Timestamp           *int                                                                                                            `json:"timestamp,omitempty"`           // Timestamp
	Attributes          *[]ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics struct {
	Response *[]ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponse `json:"response,omitempty"` //
	Page     *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsPage       `json:"page,omitempty"`     //
	Version  string                                                                                          `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponse struct {
	ID                        string                                                                                                     `json:"id,omitempty"`                        // Id
	TestID                    string                                                                                                     `json:"testId,omitempty"`                    // Test Id
	TestName                  string                                                                                                     `json:"testName,omitempty"`                  // Test Name
	TestType                  string                                                                                                     `json:"testType,omitempty"`                  // Test Type
	AgentID                   string                                                                                                     `json:"agentId,omitempty"`                   // Agent Id
	AgentName                 string                                                                                                     `json:"agentName,omitempty"`                 // Agent Name
	NetworkDeviceName         string                                                                                                     `json:"networkDeviceName,omitempty"`         // Network Device Name
	NetworkDeviceType         string                                                                                                     `json:"networkDeviceType,omitempty"`         // Network Device Type
	SiteID                    string                                                                                                     `json:"siteId,omitempty"`                    // Site Id
	SiteName                  string                                                                                                     `json:"siteName,omitempty"`                  // Site Name
	TestInterval              *int                                                                                                       `json:"testInterval,omitempty"`              // Test Interval
	TestTarget                string                                                                                                     `json:"testTarget,omitempty"`                // Test Target
	SampleTime                *int                                                                                                       `json:"sampleTime,omitempty"`                // Sample Time
	AveragePacketLoss         *float64                                                                                                   `json:"averagePacketLoss,omitempty"`         // Average Packet Loss
	LatestPacketLoss          *float64                                                                                                   `json:"latestPacketLoss,omitempty"`          // Latest Packet Loss
	MaxPacketLoss             *float64                                                                                                   `json:"maxPacketLoss,omitempty"`             // Max Packet Loss
	AverageJitter             *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseAverageJitter `json:"averageJitter,omitempty"`             // Average Jitter
	LatestJitter              *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseLatestJitter  `json:"latestJitter,omitempty"`              // Latest Jitter
	MaxJitter                 *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseMaxJitter     `json:"maxJitter,omitempty"`                 // Max Jitter
	AverageLatency            *int                                                                                                       `json:"averageLatency,omitempty"`            // Average Latency
	LatestLatency             *int                                                                                                       `json:"latestLatency,omitempty"`             // Latest Latency
	MaxLatency                *int                                                                                                       `json:"maxLatency,omitempty"`                // Max Latency
	AverageResponseTime       *int                                                                                                       `json:"averageResponseTime,omitempty"`       // Average Response Time
	LatestResponseTime        *int                                                                                                       `json:"latestResponseTime,omitempty"`        // Latest Response Time
	MaxResponseTime           *int                                                                                                       `json:"maxResponseTime,omitempty"`           // Max Response Time
	AverageMos                *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseAverageMos    `json:"averageMos,omitempty"`                // Average Mos
	LatestMos                 *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseLatestMos     `json:"latestMos,omitempty"`                 // Latest Mos
	MinMos                    *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseMinMos        `json:"minMos,omitempty"`                    // Min Mos
	AveragePdv                *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseAveragePdv    `json:"averagePdv,omitempty"`                // Average Pdv
	LatestPdv                 *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseLatestPdv     `json:"latestPdv,omitempty"`                 // Latest Pdv
	MaxPdv                    *ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseMaxPdv        `json:"maxPdv,omitempty"`                    // Max Pdv
	TotalAlerts               *float64                                                                                                   `json:"totalAlerts,omitempty"`               // Total Alerts
	TotalActiveAlerts         *int                                                                                                       `json:"totalActiveAlerts,omitempty"`         // Total Active Alerts
	TotalSamplingTests        *int                                                                                                       `json:"totalSamplingTests,omitempty"`        // Total Sampling Tests
	TotalFailureSamplingTests *int                                                                                                       `json:"totalFailureSamplingTests,omitempty"` // Total Failure Sampling Tests
	TotalErrorsSamplingTests  *int                                                                                                       `json:"totalErrorsSamplingTests,omitempty"`  // Total Errors Sampling Tests
}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseAverageJitter interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseLatestJitter interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseMaxJitter interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseAverageMos interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseLatestMos interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseMinMos interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseAveragePdv interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseLatestPdv interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsResponseMaxPdv interface{}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsPage struct {
	Limit  *int                                                                                              `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                              `json:"offset,omitempty"` // Offset
	Count  *int                                                                                              `json:"count,omitempty"`  // Count
	SortBy *[]ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseApplicationsRetrievesTheTotalCountOfThousandEyesTestResults struct {
	Response *ResponseApplicationsRetrievesTheTotalCountOfThousandEyesTestResultsResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseApplicationsRetrievesTheTotalCountOfThousandEyesTestResultsResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange struct {
	Response *[]ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeResponse `json:"response,omitempty"` //
	Page     *ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangePage       `json:"page,omitempty"`     //
	Version  string                                                                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeResponse struct {
	Timestamp  *int                                                                                                            `json:"timestamp,omitempty"`  // Timestamp
	Attributes *[]ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangePage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Offset        *int   `json:"offset,omitempty"`        // Offset
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
	Count         *int   `json:"count,omitempty"`         // Count
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
type RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics struct {
	StartTime           *int                                                                                                                     `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                                     `json:"endTime,omitempty"`             // End Time
	SiteIDs             []string                                                                                                                 `json:"siteIds,omitempty"`             // Site Ids
	Attributes          []string                                                                                                                 `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPage                  `json:"page,omitempty"`                //
}
type RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPage struct {
	Limit  *int                                                                                                            `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                            `json:"offset,omitempty"` // Offset
	Cursor string                                                                                                          `json:"cursor,omitempty"` // Cursor
	SortBy *[]RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications struct {
	StartTime           *int                                                                                                `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                `json:"endTime,omitempty"`             // End Time
	SiteIDs             []string                                                                                            `json:"siteIds,omitempty"`             // Site Ids
	TrendInterval       string                                                                                              `json:"trendInterval,omitempty"`       // Trend Interval
	GroupBy             []string                                                                                            `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                                            `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage                  `json:"page,omitempty"`                //
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication struct {
	StartTime           *int                                                                                                   `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                   `json:"endTime,omitempty"`             // End Time
	SiteIDs             []string                                                                                               `json:"siteIds,omitempty"`             // Site Ids
	TrendInterval       string                                                                                                 `json:"trendInterval,omitempty"`       // Trend Interval
	Attributes          []string                                                                                               `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationPage                  `json:"page,omitempty"`                //
}
type RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
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

//RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics Retrieves the list of ThousandEyes test results along with related metrics - 38bf-4bbf-4aba-9879
/* Retrieves the list of ThousandEyes test results along with related metrics. If `startTime` and `endTime` are not provided, the API defaults to the last 24 hours.
Please note that `siteId` filter (if used) should be using only site UUIDs of buildings.

  For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-thousandEyesTestResults-1.0.0-resolved.yaml


@param RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams Custom header parameters
@param RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-thousand-eyes-test-results-along-with-related-metrics
*/
func (s *ApplicationsService) RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics(RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams *RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams, RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsQueryParams *RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsQueryParams) (*ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics, *resty.Response, error) {
	path := "/dna/data/api/v1/thousandEyesTestResults"

	queryString, _ := query.Values(RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams != nil {

		if RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics(RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsHeaderParams, RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetricsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheListOfThousandEyesTestResultsAlongWithRelatedMetrics)
	return result, response, err

}

//RetrievesTheTotalCountOfThousandEyesTestResults Retrieves the total count of ThousandEyes test results - 58a4-0889-4a6b-9467
/* Retrieves the total count of ThousandEyes test results for the given filters. If `startTime` and `endTime` are not provided, the API defaults to the last 24 hours.  For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-thousandEyesTestResults-1.0.0-resolved.yaml


@param RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams Custom header parameters
@param RetrievesTheTotalCountOfThousandEyesTestResultsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-count-of-thousand-eyes-test-results
*/
func (s *ApplicationsService) RetrievesTheTotalCountOfThousandEyesTestResults(RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams *RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams, RetrievesTheTotalCountOfThousandEyesTestResultsQueryParams *RetrievesTheTotalCountOfThousandEyesTestResultsQueryParams) (*ResponseApplicationsRetrievesTheTotalCountOfThousandEyesTestResults, *resty.Response, error) {
	path := "/dna/data/api/v1/thousandEyesTestResults/count"

	queryString, _ := query.Values(RetrievesTheTotalCountOfThousandEyesTestResultsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams != nil {

		if RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsRetrievesTheTotalCountOfThousandEyesTestResults{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalCountOfThousandEyesTestResults(RetrievesTheTotalCountOfThousandEyesTestResultsHeaderParams, RetrievesTheTotalCountOfThousandEyesTestResultsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalCountOfThousandEyesTestResults")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheTotalCountOfThousandEyesTestResults)
	return result, response, err

}

//TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange The trend analytics data for ThousandEyes test results in the specified time range - 13ae-b974-4fe8-a677
/* Get trend time series for ThousandEyes test results.
The data will be grouped based on the specified trend time interval. If `startTime` and `endTime` are not provided, the API defaults to the last 24 hours.
By default the number of records returned will be 100 and the records will be sorted by time in ascending (`asc`) order .  For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-thousandEyesTestResults-1.0.0-resolved.yaml


@param TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams Custom header parameters
@param TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-thousand-eyes-test-results-in-the-specified-time-range
*/
func (s *ApplicationsService) TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange(TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams *TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeQueryParams *TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeQueryParams) (*ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/thousandEyesTestResults/trendAnalytics"

	queryString, _ := query.Values(TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams != nil {

		if TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange(TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRangeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseApplicationsTheTrendAnalyticsDataForThousandEyesTestResultsInTheSpecifiedTimeRange)
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

//RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics Retrieves summary analytics data related to network applications along with health metrics. - a7bb-fa3c-4c69-9457
/* Retrieves summary analytics data related to network applications while applying complex filtering, aggregate functions, and grouping.  This API facilitates obtaining consolidated insights into the performance and status of the network applications. If startTime and endTime are not provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.1-resolved.yaml


@param RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-summary-analytics-data-related-to-network-applications-along-with-health-metrics
*/
func (s *ApplicationsService) RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics(requestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics *RequestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics, RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams *RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams) (*ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams != nil {

		if RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics).
		SetResult(&ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics(requestApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics, RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetricsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics")
	}

	result := response.Result().(*ResponseApplicationsRetrievesSummaryAnalyticsDataRelatedToNetworkApplicationsAlongWithHealthMetrics)
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

//RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication Retrieves the Trend analytics related to specific network application. - 96be-dab7-4959-afc1
/* Retrieves the trend analytics of applications experience data to specific network application for the specified time range. The data will be grouped based on the given trend time interval. This API facilitates obtaining consolidated insights into the performance and status of the network applications over the specified start and end time. If startTime and endTime are not provided, the API defaults to the last 24 hours.`siteId` and `trendInterval` are mandatory. `siteId` must be a site UUID of a building.For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.1-resolved.yaml


@param id id path parameter. id is the network application name.

@param RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-trend-analytics-related-to-specific-network-application
*/
func (s *ApplicationsService) RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication(id string, requestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication *RequestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication, RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams *RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams) (*ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication, *resty.Response, error) {
	path := "/dna/data/api/v1/networkApplications/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams != nil {

		if RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication).
		SetResult(&ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication(id, requestApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication, RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplicationHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication")
	}

	result := response.Result().(*ResponseApplicationsRetrievesTheTrendAnalyticsRelatedToSpecificNetworkApplication)
	return result, response, err

}
