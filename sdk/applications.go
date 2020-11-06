package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ApplicationsService is the service to communicate with the Applications API endpoint
type ApplicationsService service

// ApplicationsResponse is the ApplicationsResponse definition
type ApplicationsResponse struct {
	Response []struct {
		ApplicationServerLatency string `json:"applicationServerLatency,omitempty"` //
		AverageThroughput        int    `json:"averageThroughput,omitempty"`        //
		BusinessRelevance        string `json:"businessRelevance,omitempty"`        //
		ClientNetworkLatency     string `json:"clientNetworkLatency,omitempty"`     //
		Health                   string `json:"health,omitempty"`                   //
		Jitter                   string `json:"jitter,omitempty"`                   //
		Name                     string `json:"name,omitempty"`                     //
		NetworkLatency           string `json:"networkLatency,omitempty"`           //
		PacketLossPercent        string `json:"packetLossPercent,omitempty"`        //
		ServerNetworkLatency     string `json:"serverNetworkLatency,omitempty"`     //
		TrafficClass             string `json:"trafficClass,omitempty"`             //
		UsageBytes               int    `json:"usageBytes,omitempty"`               //
	} `json:"response,omitempty"` //
	TotalCount int    `json:"totalCount,omitempty"` //
	Version    string `json:"version,omitempty"`    //
}

// ApplicationsQueryParams defines the query parameters for this request
type ApplicationsQueryParams struct {
	SiteID            string `url:"siteId,omitempty"`            // Assurance site UUID value (Cannot be submitted together with deviceId and clientMac)
	DeviceID          string `url:"deviceId,omitempty"`          // Assurance device UUID value (Cannot be submitted together with siteId and clientMac)
	MacAddress        string `url:"macAddress,omitempty"`        // Client device's MAC address (Cannot be submitted together with siteId and deviceId)
	StartTime         int    `url:"startTime,omitempty"`         // Starting epoch time in milliseconds of time window
	EndTime           int    `url:"endTime,omitempty"`           // Ending epoch time in milliseconds of time window
	ApplicationHealth string `url:"applicationHealth,omitempty"` // Application health category (POOR, FAIR, or GOOD.  Optionally use with siteId only)
	Offset            int    `url:"offset,omitempty"`            // The offset of the first application in the returned data (optionally used with siteId only)
	Limit             int    `url:"limit,omitempty"`             // The max number of application entries in returned data [1, 1000] (optionally used with siteId only)
}

// Applications applications
/* Intent API to get a list of applications for a specific site, a device, or a client device's MAC address.
@param siteID Assurance site UUID value (Cannot be submitted together with deviceId and clientMac)
@param deviceID Assurance device UUID value (Cannot be submitted together with siteId and clientMac)
@param macAddress Client device's MAC address (Cannot be submitted together with siteId and deviceId)
@param startTime Starting epoch time in milliseconds of time window
@param endTime Ending epoch time in milliseconds of time window
@param applicationHealth Application health category (POOR, FAIR, or GOOD.  Optionally use with siteId only)
@param offset The offset of the first application in the returned data (optionally used with siteId only)
@param limit The max number of application entries in returned data [1, 1000] (optionally used with siteId only)
*/
func (s *ApplicationsService) Applications(applicationsQueryParams *ApplicationsQueryParams) (*ApplicationsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/application-health"

	queryString, _ := query.Values(applicationsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&ApplicationsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*ApplicationsResponse)
	return result, response, err
}
