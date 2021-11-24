package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TopologyService service

type GetOverallNetworkHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Network health data is required
}
type GetPhysicalTopologyQueryParams struct {
	NodeType string `url:"nodeType,omitempty"` //nodeType
}

type ResponseTopologyGetOverallNetworkHealth struct {
	Version                   string                                                         `json:"version,omitempty"`                   // Version
	Response                  *[]ResponseTopologyGetOverallNetworkHealthResponse             `json:"response,omitempty"`                  //
	MeasuredBy                string                                                         `json:"measuredBy,omitempty"`                // Measured By
	LatestMeasuredByEntity    *ResponseTopologyGetOverallNetworkHealthLatestMeasuredByEntity `json:"latestMeasuredByEntity,omitempty"`    // Latest Measured By Entity
	LatestHealthScore         *int                                                           `json:"latestHealthScore,omitempty"`         // Latest Health Score
	MonitoredDevices          *int                                                           `json:"monitoredDevices,omitempty"`          // Monitored Devices
	MonitoredHealthyDevices   *int                                                           `json:"monitoredHealthyDevices,omitempty"`   // Monitored Healthy Devices
	MonitoredUnHealthyDevices *int                                                           `json:"monitoredUnHealthyDevices,omitempty"` // Monitored Un Healthy Devices
	UnMonitoredDevices        *float64                                                       `json:"unMonitoredDevices,omitempty"`        // Un Monitored Devices
	HealthDistirubution       *[]ResponseTopologyGetOverallNetworkHealthHealthDistirubution  `json:"healthDistirubution,omitempty"`       //
}
type ResponseTopologyGetOverallNetworkHealthResponse struct {
	Time         string                                                 `json:"time,omitempty"`         // Time
	HealthScore  *int                                                   `json:"healthScore,omitempty"`  // Health Score
	TotalCount   *int                                                   `json:"totalCount,omitempty"`   // Total Count
	GoodCount    *int                                                   `json:"goodCount,omitempty"`    // Good Count
	UnmonCount   *float64                                               `json:"unmonCount,omitempty"`   // Unmon Count
	FairCount    *int                                                   `json:"fairCount,omitempty"`    // Fair Count
	BadCount     *int                                                   `json:"badCount,omitempty"`     // Bad Count
	Entity       *ResponseTopologyGetOverallNetworkHealthResponseEntity `json:"entity,omitempty"`       // Entity
	TimeinMillis *int                                                   `json:"timeinMillis,omitempty"` // Timein Millis
}
type ResponseTopologyGetOverallNetworkHealthResponseEntity interface{}
type ResponseTopologyGetOverallNetworkHealthLatestMeasuredByEntity interface{}
type ResponseTopologyGetOverallNetworkHealthHealthDistirubution struct {
	Category        string                                                                  `json:"category,omitempty"`        // Category
	TotalCount      *int                                                                    `json:"totalCount,omitempty"`      // Total Count
	HealthScore     *int                                                                    `json:"healthScore,omitempty"`     // Health Score
	GoodPercentage  *float64                                                                `json:"goodPercentage,omitempty"`  // Good Percentage
	BadPercentage   *float64                                                                `json:"badPercentage,omitempty"`   // Bad Percentage
	FairPercentage  *float64                                                                `json:"fairPercentage,omitempty"`  // Fair Percentage
	UnmonPercentage *float64                                                                `json:"unmonPercentage,omitempty"` // Unmon Percentage
	GoodCount       *int                                                                    `json:"goodCount,omitempty"`       // Good Count
	BadCount        *float64                                                                `json:"badCount,omitempty"`        // Bad Count
	FairCount       *int                                                                    `json:"fairCount,omitempty"`       // Fair Count
	UnmonCount      *int                                                                    `json:"unmonCount,omitempty"`      // Unmon Count
	KpiMetrics      *[]ResponseTopologyGetOverallNetworkHealthHealthDistirubutionKpiMetrics `json:"kpiMetrics,omitempty"`      // Kpi Metrics
}
type ResponseTopologyGetOverallNetworkHealthHealthDistirubutionKpiMetrics struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseTopologyGetTopologyDetails struct {
	Response *ResponseTopologyGetTopologyDetailsResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseTopologyGetTopologyDetailsResponse struct {
	ID    string                                             `json:"id,omitempty"`    //
	Links *[]ResponseTopologyGetTopologyDetailsResponseLinks `json:"links,omitempty"` //
	Nodes *[]ResponseTopologyGetTopologyDetailsResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetTopologyDetailsResponseLinks struct {
	AdditionalInfo       *ResponseTopologyGetTopologyDetailsResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"`       //
	EndPortID            string                                                         `json:"endPortID,omitempty"`            //
	EndPortIPv4Address   string                                                         `json:"endPortIpv4Address,omitempty"`   //
	EndPortIPv4Mask      string                                                         `json:"endPortIpv4Mask,omitempty"`      //
	EndPortName          string                                                         `json:"endPortName,omitempty"`          //
	EndPortSpeed         string                                                         `json:"endPortSpeed,omitempty"`         //
	GreyOut              *bool                                                          `json:"greyOut,omitempty"`              //
	ID                   string                                                         `json:"id,omitempty"`                   //
	LinkStatus           string                                                         `json:"linkStatus,omitempty"`           //
	Source               string                                                         `json:"source,omitempty"`               //
	StartPortID          string                                                         `json:"startPortID,omitempty"`          //
	StartPortIPv4Address string                                                         `json:"startPortIpv4Address,omitempty"` //
	StartPortIPv4Mask    string                                                         `json:"startPortIpv4Mask,omitempty"`    //
	StartPortName        string                                                         `json:"startPortName,omitempty"`        //
	StartPortSpeed       string                                                         `json:"startPortSpeed,omitempty"`       //
	Tag                  string                                                         `json:"tag,omitempty"`                  //
	Target               string                                                         `json:"target,omitempty"`               //
}
type ResponseTopologyGetTopologyDetailsResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetTopologyDetailsResponseNodes struct {
	ACLApplied      *bool                                                          `json:"aclApplied,omitempty"`      //
	AdditionalInfo  *ResponseTopologyGetTopologyDetailsResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"`  //
	CustomParam     *ResponseTopologyGetTopologyDetailsResponseNodesCustomParam    `json:"customParam,omitempty"`     //
	DataPathID      string                                                         `json:"dataPathId,omitempty"`      //
	DeviceType      string                                                         `json:"deviceType,omitempty"`      //
	Family          string                                                         `json:"family,omitempty"`          //
	Fixed           *bool                                                          `json:"fixed,omitempty"`           //
	GreyOut         *bool                                                          `json:"greyOut,omitempty"`         //
	ID              string                                                         `json:"id,omitempty"`              //
	IP              string                                                         `json:"ip,omitempty"`              //
	Label           string                                                         `json:"label,omitempty"`           //
	NetworkType     string                                                         `json:"networkType,omitempty"`     //
	NodeType        string                                                         `json:"nodeType,omitempty"`        //
	Order           *int                                                           `json:"order,omitempty"`           //
	OsType          string                                                         `json:"osType,omitempty"`          //
	PlatformID      string                                                         `json:"platformId,omitempty"`      //
	Role            string                                                         `json:"role,omitempty"`            //
	RoleSource      string                                                         `json:"roleSource,omitempty"`      //
	SoftwareVersion string                                                         `json:"softwareVersion,omitempty"` //
	Tags            []string                                                       `json:"tags,omitempty"`            //
	UpperNode       string                                                         `json:"upperNode,omitempty"`       //
	UserID          string                                                         `json:"userId,omitempty"`          //
	VLANID          string                                                         `json:"vlanId,omitempty"`          //
	X               *int                                                           `json:"x,omitempty"`               //
	Y               *int                                                           `json:"y,omitempty"`               //
}
type ResponseTopologyGetTopologyDetailsResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetTopologyDetailsResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            *int   `json:"x,omitempty"`            //
	Y            *int   `json:"y,omitempty"`            //
}
type ResponseTopologyGetL3TopologyDetails struct {
	Response *ResponseTopologyGetL3TopologyDetailsResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}
type ResponseTopologyGetL3TopologyDetailsResponse struct {
	ID    string                                               `json:"id,omitempty"`    //
	Links *[]ResponseTopologyGetL3TopologyDetailsResponseLinks `json:"links,omitempty"` //
	Nodes *[]ResponseTopologyGetL3TopologyDetailsResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetL3TopologyDetailsResponseLinks struct {
	AdditionalInfo       *ResponseTopologyGetL3TopologyDetailsResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"`       //
	EndPortID            string                                                           `json:"endPortID,omitempty"`            //
	EndPortIPv4Address   string                                                           `json:"endPortIpv4Address,omitempty"`   //
	EndPortIPv4Mask      string                                                           `json:"endPortIpv4Mask,omitempty"`      //
	EndPortName          string                                                           `json:"endPortName,omitempty"`          //
	EndPortSpeed         string                                                           `json:"endPortSpeed,omitempty"`         //
	GreyOut              *bool                                                            `json:"greyOut,omitempty"`              //
	ID                   string                                                           `json:"id,omitempty"`                   //
	LinkStatus           string                                                           `json:"linkStatus,omitempty"`           //
	Source               string                                                           `json:"source,omitempty"`               //
	StartPortID          string                                                           `json:"startPortID,omitempty"`          //
	StartPortIPv4Address string                                                           `json:"startPortIpv4Address,omitempty"` //
	StartPortIPv4Mask    string                                                           `json:"startPortIpv4Mask,omitempty"`    //
	StartPortName        string                                                           `json:"startPortName,omitempty"`        //
	StartPortSpeed       string                                                           `json:"startPortSpeed,omitempty"`       //
	Tag                  string                                                           `json:"tag,omitempty"`                  //
	Target               string                                                           `json:"target,omitempty"`               //
}
type ResponseTopologyGetL3TopologyDetailsResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetL3TopologyDetailsResponseNodes struct {
	ACLApplied      *bool                                                            `json:"aclApplied,omitempty"`      //
	AdditionalInfo  *ResponseTopologyGetL3TopologyDetailsResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"`  //
	CustomParam     *ResponseTopologyGetL3TopologyDetailsResponseNodesCustomParam    `json:"customParam,omitempty"`     //
	DataPathID      string                                                           `json:"dataPathId,omitempty"`      //
	DeviceType      string                                                           `json:"deviceType,omitempty"`      //
	Family          string                                                           `json:"family,omitempty"`          //
	Fixed           *bool                                                            `json:"fixed,omitempty"`           //
	GreyOut         *bool                                                            `json:"greyOut,omitempty"`         //
	ID              string                                                           `json:"id,omitempty"`              //
	IP              string                                                           `json:"ip,omitempty"`              //
	Label           string                                                           `json:"label,omitempty"`           //
	NetworkType     string                                                           `json:"networkType,omitempty"`     //
	NodeType        string                                                           `json:"nodeType,omitempty"`        //
	Order           *int                                                             `json:"order,omitempty"`           //
	OsType          string                                                           `json:"osType,omitempty"`          //
	PlatformID      string                                                           `json:"platformId,omitempty"`      //
	Role            string                                                           `json:"role,omitempty"`            //
	RoleSource      string                                                           `json:"roleSource,omitempty"`      //
	SoftwareVersion string                                                           `json:"softwareVersion,omitempty"` //
	Tags            []string                                                         `json:"tags,omitempty"`            //
	UpperNode       string                                                           `json:"upperNode,omitempty"`       //
	UserID          string                                                           `json:"userId,omitempty"`          //
	VLANID          string                                                           `json:"vlanId,omitempty"`          //
	X               *int                                                             `json:"x,omitempty"`               //
	Y               *int                                                             `json:"y,omitempty"`               //
}
type ResponseTopologyGetL3TopologyDetailsResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetL3TopologyDetailsResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            *int   `json:"x,omitempty"`            //
	Y            *int   `json:"y,omitempty"`            //
}
type ResponseTopologyGetPhysicalTopology struct {
	Response *ResponseTopologyGetPhysicalTopologyResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}
type ResponseTopologyGetPhysicalTopologyResponse struct {
	ID    string                                              `json:"id,omitempty"`    //
	Links *[]ResponseTopologyGetPhysicalTopologyResponseLinks `json:"links,omitempty"` //
	Nodes *[]ResponseTopologyGetPhysicalTopologyResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetPhysicalTopologyResponseLinks struct {
	AdditionalInfo       *ResponseTopologyGetPhysicalTopologyResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"`       //
	EndPortID            string                                                          `json:"endPortID,omitempty"`            //
	EndPortIPv4Address   string                                                          `json:"endPortIpv4Address,omitempty"`   //
	EndPortIPv4Mask      string                                                          `json:"endPortIpv4Mask,omitempty"`      //
	EndPortName          string                                                          `json:"endPortName,omitempty"`          //
	EndPortSpeed         string                                                          `json:"endPortSpeed,omitempty"`         //
	GreyOut              *bool                                                           `json:"greyOut,omitempty"`              //
	ID                   string                                                          `json:"id,omitempty"`                   //
	LinkStatus           string                                                          `json:"linkStatus,omitempty"`           //
	Source               string                                                          `json:"source,omitempty"`               //
	StartPortID          string                                                          `json:"startPortID,omitempty"`          //
	StartPortIPv4Address string                                                          `json:"startPortIpv4Address,omitempty"` //
	StartPortIPv4Mask    string                                                          `json:"startPortIpv4Mask,omitempty"`    //
	StartPortName        string                                                          `json:"startPortName,omitempty"`        //
	StartPortSpeed       string                                                          `json:"startPortSpeed,omitempty"`       //
	Tag                  string                                                          `json:"tag,omitempty"`                  //
	Target               string                                                          `json:"target,omitempty"`               //
}
type ResponseTopologyGetPhysicalTopologyResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetPhysicalTopologyResponseNodes struct {
	ACLApplied      *bool                                                           `json:"aclApplied,omitempty"`      //
	AdditionalInfo  *ResponseTopologyGetPhysicalTopologyResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"`  //
	CustomParam     *ResponseTopologyGetPhysicalTopologyResponseNodesCustomParam    `json:"customParam,omitempty"`     //
	DataPathID      string                                                          `json:"dataPathId,omitempty"`      //
	DeviceType      string                                                          `json:"deviceType,omitempty"`      //
	Family          string                                                          `json:"family,omitempty"`          //
	Fixed           *bool                                                           `json:"fixed,omitempty"`           //
	GreyOut         *bool                                                           `json:"greyOut,omitempty"`         //
	ID              string                                                          `json:"id,omitempty"`              //
	IP              string                                                          `json:"ip,omitempty"`              //
	Label           string                                                          `json:"label,omitempty"`           //
	NetworkType     string                                                          `json:"networkType,omitempty"`     //
	NodeType        string                                                          `json:"nodeType,omitempty"`        //
	Order           *int                                                            `json:"order,omitempty"`           //
	OsType          string                                                          `json:"osType,omitempty"`          //
	PlatformID      string                                                          `json:"platformId,omitempty"`      //
	Role            string                                                          `json:"role,omitempty"`            //
	RoleSource      string                                                          `json:"roleSource,omitempty"`      //
	SoftwareVersion string                                                          `json:"softwareVersion,omitempty"` //
	Tags            []string                                                        `json:"tags,omitempty"`            //
	UpperNode       string                                                          `json:"upperNode,omitempty"`       //
	UserID          string                                                          `json:"userId,omitempty"`          //
	VLANID          string                                                          `json:"vlanId,omitempty"`          //
	X               *int                                                            `json:"x,omitempty"`               //
	Y               *int                                                            `json:"y,omitempty"`               //
}
type ResponseTopologyGetPhysicalTopologyResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetPhysicalTopologyResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            *int   `json:"x,omitempty"`            //
	Y            *int   `json:"y,omitempty"`            //
}
type ResponseTopologyGetSiteTopology struct {
	Response *ResponseTopologyGetSiteTopologyResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseTopologyGetSiteTopologyResponse struct {
	Sites *[]ResponseTopologyGetSiteTopologyResponseSites `json:"sites,omitempty"` //
}
type ResponseTopologyGetSiteTopologyResponseSites struct {
	DisplayName        string `json:"displayName,omitempty"`        //
	GroupNameHierarchy string `json:"groupNameHierarchy,omitempty"` //
	ID                 string `json:"id,omitempty"`                 //
	Latitude           string `json:"latitude,omitempty"`           //
	LocationAddress    string `json:"locationAddress,omitempty"`    //
	LocationCountry    string `json:"locationCountry,omitempty"`    //
	LocationType       string `json:"locationType,omitempty"`       //
	Longitude          string `json:"longitude,omitempty"`          //
	Name               string `json:"name,omitempty"`               //
	ParentID           string `json:"parentId,omitempty"`           //
}
type ResponseTopologyGetVLANDetails struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

//GetOverallNetworkHealth Get Overall Network Health - ca91-da84-401a-bba1
/* Returns Overall Network Health information by Device category (Access, Distribution, Core, Router, Wireless) for any given point of time


@param GetOverallNetworkHealthQueryParams Filtering parameter
*/
func (s *TopologyService) GetOverallNetworkHealth(GetOverallNetworkHealthQueryParams *GetOverallNetworkHealthQueryParams) (*ResponseTopologyGetOverallNetworkHealth, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-health"

	queryString, _ := query.Values(GetOverallNetworkHealthQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTopologyGetOverallNetworkHealth{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOverallNetworkHealth")
	}

	result := response.Result().(*ResponseTopologyGetOverallNetworkHealth)
	return result, response, err

}

//GetTopologyDetails Get topology details - b9b4-8ac8-463a-8aba
/* Returns Layer 2 network topology by specified VLAN ID


@param vlanID vlanID path parameter. Vlan Name for e.g Vlan1, Vlan23 etc

*/
func (s *TopologyService) GetTopologyDetails(vlanID string) (*ResponseTopologyGetTopologyDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/l2/{vlanID}"
	path = strings.Replace(path, "{vlanID}", fmt.Sprintf("%v", vlanID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetTopologyDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTopologyDetails")
	}

	result := response.Result().(*ResponseTopologyGetTopologyDetails)
	return result, response, err

}

//GetL3TopologyDetails Get L3 Topology Details - c2b5-fb76-4d88-8375
/* Returns the Layer 3 network topology by routing protocol


@param topologyType topologyType path parameter. Type of topology(OSPF,ISIS,etc)

*/
func (s *TopologyService) GetL3TopologyDetails(topologyType string) (*ResponseTopologyGetL3TopologyDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/l3/{topologyType}"
	path = strings.Replace(path, "{topologyType}", fmt.Sprintf("%v", topologyType), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetL3TopologyDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetL3TopologyDetails")
	}

	result := response.Result().(*ResponseTopologyGetL3TopologyDetails)
	return result, response, err

}

//GetPhysicalTopology Get Physical Topology - b2b8-cb91-459a-a58f
/* Returns the raw physical topology by specified criteria of nodeType


@param GetPhysicalTopologyQueryParams Filtering parameter
*/
func (s *TopologyService) GetPhysicalTopology(GetPhysicalTopologyQueryParams *GetPhysicalTopologyQueryParams) (*ResponseTopologyGetPhysicalTopology, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/physical-topology"

	queryString, _ := query.Values(GetPhysicalTopologyQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTopologyGetPhysicalTopology{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPhysicalTopology")
	}

	result := response.Result().(*ResponseTopologyGetPhysicalTopology)
	return result, response, err

}

//GetSiteTopology Get Site Topology - 9ba1-4a9e-441b-8a60
/* Returns site topology


 */
func (s *TopologyService) GetSiteTopology() (*ResponseTopologyGetSiteTopology, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/site-topology"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetSiteTopology{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSiteTopology")
	}

	result := response.Result().(*ResponseTopologyGetSiteTopology)
	return result, response, err

}

//GetVLANDetails Get VLAN details - 6284-db46-49aa-8d31
/* Returns the list of VLAN names


 */
func (s *TopologyService) GetVLANDetails() (*ResponseTopologyGetVLANDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/topology/vlan/vlan-names"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTopologyGetVLANDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVlanDetails")
	}

	result := response.Result().(*ResponseTopologyGetVLANDetails)
	return result, response, err

}
