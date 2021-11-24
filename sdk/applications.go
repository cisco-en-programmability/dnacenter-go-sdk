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
}

type ResponseApplicationsApplications struct {
	Version    string                                      `json:"version,omitempty"`    // Version
	TotalCount *int                                        `json:"totalCount,omitempty"` // Total Count
	Response   *[]ResponseApplicationsApplicationsResponse `json:"response,omitempty"`   //
}
type ResponseApplicationsApplicationsResponse struct {
	Name                     string                                                            `json:"name,omitempty"`                     // Name
	Health                   *ResponseApplicationsApplicationsResponseHealth                   `json:"health,omitempty"`                   // Health
	BusinessRelevance        string                                                            `json:"businessRelevance,omitempty"`        // Business Relevance
	TrafficClass             string                                                            `json:"trafficClass,omitempty"`             // Traffic Class
	UsageBytes               *int                                                              `json:"usageBytes,omitempty"`               // Usage Bytes
	AverageThroughput        *float64                                                          `json:"averageThroughput,omitempty"`        // Average Throughput
	PacketLossPercent        *ResponseApplicationsApplicationsResponsePacketLossPercent        `json:"packetLossPercent,omitempty"`        // Packet Loss Percent
	NetworkLatency           *ResponseApplicationsApplicationsResponseNetworkLatency           `json:"networkLatency,omitempty"`           // Network Latency
	Jitter                   *ResponseApplicationsApplicationsResponseJitter                   `json:"jitter,omitempty"`                   // Jitter
	ApplicationServerLatency *ResponseApplicationsApplicationsResponseApplicationServerLatency `json:"applicationServerLatency,omitempty"` // Application Server Latency
	ClientNetworkLatency     *ResponseApplicationsApplicationsResponseClientNetworkLatency     `json:"clientNetworkLatency,omitempty"`     // Client Network Latency
	ServerNetworkLatency     *ResponseApplicationsApplicationsResponseServerNetworkLatency     `json:"serverNetworkLatency,omitempty"`     // Server Network Latency
}
type ResponseApplicationsApplicationsResponseHealth interface{}
type ResponseApplicationsApplicationsResponsePacketLossPercent interface{}
type ResponseApplicationsApplicationsResponseNetworkLatency interface{}
type ResponseApplicationsApplicationsResponseJitter interface{}
type ResponseApplicationsApplicationsResponseApplicationServerLatency interface{}
type ResponseApplicationsApplicationsResponseClientNetworkLatency interface{}
type ResponseApplicationsApplicationsResponseServerNetworkLatency interface{}

//Applications Applications - 2db5-8a1f-4fea-9242
/* Intent API to get a list of applications for a specific site, a device, or a client device's MAC address.


@param ApplicationsQueryParams Filtering parameter
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
