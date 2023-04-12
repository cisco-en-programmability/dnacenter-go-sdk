package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplicationsService service

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
		return nil, response, fmt.Errorf("error with operation Applications")
	}

	result := response.Result().(*ResponseApplicationsApplications)
	return result, response, err

}
