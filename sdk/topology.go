package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// TopologyService is the service to communicate with the Topology API endpoint
type TopologyService service

// GetL3TopologyDetailsResponse is the getL3TopologyDetailsResponse definition
type GetL3TopologyDetailsResponse struct {
	Response GetL3TopologyDetailsResponseResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

// GetL3TopologyDetailsResponseResponse is the getL3TopologyDetailsResponseResponse definition
type GetL3TopologyDetailsResponseResponse struct {
	ID    string                                      `json:"id,omitempty"`    //
	Links []GetL3TopologyDetailsResponseResponseLinks `json:"links,omitempty"` //
	Nodes []GetL3TopologyDetailsResponseResponseNodes `json:"nodes,omitempty"` //
}

// GetL3TopologyDetailsResponseResponseLinks is the getL3TopologyDetailsResponseResponseLinks definition
type GetL3TopologyDetailsResponseResponseLinks struct {
	AdditionalInfo       string `json:"additionalInfo,omitempty"`       //
	EndPortID            string `json:"endPortID,omitempty"`            //
	EndPortIPv4Address   string `json:"endPortIpv4Address,omitempty"`   //
	EndPortIPv4Mask      string `json:"endPortIpv4Mask,omitempty"`      //
	EndPortName          string `json:"endPortName,omitempty"`          //
	EndPortSpeed         string `json:"endPortSpeed,omitempty"`         //
	GreyOut              bool   `json:"greyOut,omitempty"`              //
	ID                   string `json:"id,omitempty"`                   //
	LinkStatus           string `json:"linkStatus,omitempty"`           //
	Source               string `json:"source,omitempty"`               //
	StartPortID          string `json:"startPortID,omitempty"`          //
	StartPortIPv4Address string `json:"startPortIpv4Address,omitempty"` //
	StartPortIPv4Mask    string `json:"startPortIpv4Mask,omitempty"`    //
	StartPortName        string `json:"startPortName,omitempty"`        //
	StartPortSpeed       string `json:"startPortSpeed,omitempty"`       //
	Tag                  string `json:"tag,omitempty"`                  //
	Target               string `json:"target,omitempty"`               //
}

// GetL3TopologyDetailsResponseResponseNodes is the getL3TopologyDetailsResponseResponseNodes definition
type GetL3TopologyDetailsResponseResponseNodes struct {
	ACLApplied      bool                                                 `json:"aclApplied,omitempty"`      //
	AdditionalInfo  string                                               `json:"additionalInfo,omitempty"`  //
	CustomParam     GetL3TopologyDetailsResponseResponseNodesCustomParam `json:"customParam,omitempty"`     //
	DataPathID      string                                               `json:"dataPathId,omitempty"`      //
	DeviceType      string                                               `json:"deviceType,omitempty"`      //
	Family          string                                               `json:"family,omitempty"`          //
	Fixed           bool                                                 `json:"fixed,omitempty"`           //
	GreyOut         bool                                                 `json:"greyOut,omitempty"`         //
	ID              string                                               `json:"id,omitempty"`              //
	IP              string                                               `json:"ip,omitempty"`              //
	Label           string                                               `json:"label,omitempty"`           //
	NetworkType     string                                               `json:"networkType,omitempty"`     //
	NodeType        string                                               `json:"nodeType,omitempty"`        //
	Order           int                                                  `json:"order,omitempty"`           //
	OsType          string                                               `json:"osType,omitempty"`          //
	PlatformID      string                                               `json:"platformId,omitempty"`      //
	Role            string                                               `json:"role,omitempty"`            //
	RoleSource      string                                               `json:"roleSource,omitempty"`      //
	SoftwareVersion string                                               `json:"softwareVersion,omitempty"` //
	Tags            []string                                             `json:"tags,omitempty"`            //
	UpperNode       string                                               `json:"upperNode,omitempty"`       //
	UserID          string                                               `json:"userId,omitempty"`          //
	VLANID          string                                               `json:"vlanId,omitempty"`          //
	X               int                                                  `json:"x,omitempty"`               //
	Y               int                                                  `json:"y,omitempty"`               //
}

// GetL3TopologyDetailsResponseResponseNodesCustomParam is the getL3TopologyDetailsResponseResponseNodesCustomParam definition
type GetL3TopologyDetailsResponseResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            int    `json:"x,omitempty"`            //
	Y            int    `json:"y,omitempty"`            //
}

// GetL3TopologyDetailsResponseResponseNodesTags is the getL3TopologyDetailsResponseResponseNodesTags definition
type GetL3TopologyDetailsResponseResponseNodesTags []string

// GetOverallNetworkHealthResponse is the getOverallNetworkHealthResponse definition
type GetOverallNetworkHealthResponse struct {
	HealthDistirubution       []GetOverallNetworkHealthResponseHealthDistirubution `json:"healthDistirubution,omitempty"`       //
	LatestHealthScore         int                                                  `json:"latestHealthScore,omitempty"`         //
	LatestMeasuredByEntity    string                                               `json:"latestMeasuredByEntity,omitempty"`    //
	MeasuredBy                string                                               `json:"measuredBy,omitempty"`                //
	MonitoredDevices          int                                                  `json:"monitoredDevices,omitempty"`          //
	MonitoredHealthyDevices   int                                                  `json:"monitoredHealthyDevices,omitempty"`   //
	MonitoredUnHealthyDevices int                                                  `json:"monitoredUnHealthyDevices,omitempty"` //
	Response                  []GetOverallNetworkHealthResponseResponse            `json:"response,omitempty"`                  //
	UnMonitoredDevices        float64                                              `json:"unMonitoredDevices,omitempty"`        //
	Version                   string                                               `json:"version,omitempty"`                   //
}

// GetOverallNetworkHealthResponseHealthDistirubution is the getOverallNetworkHealthResponseHealthDistirubution definition
type GetOverallNetworkHealthResponseHealthDistirubution struct {
	BadCount        float64  `json:"badCount,omitempty"`        //
	BadPercentage   int      `json:"badPercentage,omitempty"`   //
	Category        string   `json:"category,omitempty"`        //
	FairCount       float64  `json:"fairCount,omitempty"`       //
	FairPercentage  int      `json:"fairPercentage,omitempty"`  //
	GoodCount       int      `json:"goodCount,omitempty"`       //
	GoodPercentage  int      `json:"goodPercentage,omitempty"`  //
	HealthScore     int      `json:"healthScore,omitempty"`     //
	KpiMetrics      []string `json:"kpiMetrics,omitempty"`      //
	TotalCount      int      `json:"totalCount,omitempty"`      //
	UnmonCount      float64  `json:"unmonCount,omitempty"`      //
	UnmonPercentage int      `json:"unmonPercentage,omitempty"` //
}

// GetOverallNetworkHealthResponseHealthDistirubutionKpiMetrics is the getOverallNetworkHealthResponseHealthDistirubutionKpiMetrics definition
type GetOverallNetworkHealthResponseHealthDistirubutionKpiMetrics []string

// GetOverallNetworkHealthResponseResponse is the getOverallNetworkHealthResponseResponse definition
type GetOverallNetworkHealthResponseResponse struct {
	BadCount     float64 `json:"badCount,omitempty"`     //
	Entity       string  `json:"entity,omitempty"`       //
	FairCount    int     `json:"fairCount,omitempty"`    //
	GoodCount    int     `json:"goodCount,omitempty"`    //
	HealthScore  int     `json:"healthScore,omitempty"`  //
	Time         string  `json:"time,omitempty"`         //
	TimeinMillis int     `json:"timeinMillis,omitempty"` //
	TotalCount   int     `json:"totalCount,omitempty"`   //
	UnmonCount   float64 `json:"unmonCount,omitempty"`   //
}

// GetPhysicalTopologyResponse is the getPhysicalTopologyResponse definition
type GetPhysicalTopologyResponse struct {
	Response GetPhysicalTopologyResponseResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  //
}

// GetPhysicalTopologyResponseResponse is the getPhysicalTopologyResponseResponse definition
type GetPhysicalTopologyResponseResponse struct {
	ID    string                                     `json:"id,omitempty"`    //
	Links []GetPhysicalTopologyResponseResponseLinks `json:"links,omitempty"` //
	Nodes []GetPhysicalTopologyResponseResponseNodes `json:"nodes,omitempty"` //
}

// GetPhysicalTopologyResponseResponseLinks is the getPhysicalTopologyResponseResponseLinks definition
type GetPhysicalTopologyResponseResponseLinks struct {
	AdditionalInfo       string `json:"additionalInfo,omitempty"`       //
	EndPortID            string `json:"endPortID,omitempty"`            //
	EndPortIPv4Address   string `json:"endPortIpv4Address,omitempty"`   //
	EndPortIPv4Mask      string `json:"endPortIpv4Mask,omitempty"`      //
	EndPortName          string `json:"endPortName,omitempty"`          //
	EndPortSpeed         string `json:"endPortSpeed,omitempty"`         //
	GreyOut              bool   `json:"greyOut,omitempty"`              //
	ID                   string `json:"id,omitempty"`                   //
	LinkStatus           string `json:"linkStatus,omitempty"`           //
	Source               string `json:"source,omitempty"`               //
	StartPortID          string `json:"startPortID,omitempty"`          //
	StartPortIPv4Address string `json:"startPortIpv4Address,omitempty"` //
	StartPortIPv4Mask    string `json:"startPortIpv4Mask,omitempty"`    //
	StartPortName        string `json:"startPortName,omitempty"`        //
	StartPortSpeed       string `json:"startPortSpeed,omitempty"`       //
	Tag                  string `json:"tag,omitempty"`                  //
	Target               string `json:"target,omitempty"`               //
}

// GetPhysicalTopologyResponseResponseNodes is the getPhysicalTopologyResponseResponseNodes definition
type GetPhysicalTopologyResponseResponseNodes struct {
	ACLApplied      bool                                                `json:"aclApplied,omitempty"`      //
	AdditionalInfo  string                                              `json:"additionalInfo,omitempty"`  //
	CustomParam     GetPhysicalTopologyResponseResponseNodesCustomParam `json:"customParam,omitempty"`     //
	DataPathID      string                                              `json:"dataPathId,omitempty"`      //
	DeviceType      string                                              `json:"deviceType,omitempty"`      //
	Family          string                                              `json:"family,omitempty"`          //
	Fixed           bool                                                `json:"fixed,omitempty"`           //
	GreyOut         bool                                                `json:"greyOut,omitempty"`         //
	ID              string                                              `json:"id,omitempty"`              //
	IP              string                                              `json:"ip,omitempty"`              //
	Label           string                                              `json:"label,omitempty"`           //
	NetworkType     string                                              `json:"networkType,omitempty"`     //
	NodeType        string                                              `json:"nodeType,omitempty"`        //
	Order           int                                                 `json:"order,omitempty"`           //
	OsType          string                                              `json:"osType,omitempty"`          //
	PlatformID      string                                              `json:"platformId,omitempty"`      //
	Role            string                                              `json:"role,omitempty"`            //
	RoleSource      string                                              `json:"roleSource,omitempty"`      //
	SoftwareVersion string                                              `json:"softwareVersion,omitempty"` //
	Tags            []string                                            `json:"tags,omitempty"`            //
	UpperNode       string                                              `json:"upperNode,omitempty"`       //
	UserID          string                                              `json:"userId,omitempty"`          //
	VLANID          string                                              `json:"vlanId,omitempty"`          //
	X               int                                                 `json:"x,omitempty"`               //
	Y               int                                                 `json:"y,omitempty"`               //
}

// GetPhysicalTopologyResponseResponseNodesCustomParam is the getPhysicalTopologyResponseResponseNodesCustomParam definition
type GetPhysicalTopologyResponseResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            int    `json:"x,omitempty"`            //
	Y            int    `json:"y,omitempty"`            //
}

// GetPhysicalTopologyResponseResponseNodesTags is the getPhysicalTopologyResponseResponseNodesTags definition
type GetPhysicalTopologyResponseResponseNodesTags []string

// GetSiteTopologyResponse is the getSiteTopologyResponse definition
type GetSiteTopologyResponse struct {
	Response GetSiteTopologyResponseResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}

// GetSiteTopologyResponseResponse is the getSiteTopologyResponseResponse definition
type GetSiteTopologyResponseResponse struct {
	Sites []GetSiteTopologyResponseResponseSites `json:"sites,omitempty"` //
}

// GetSiteTopologyResponseResponseSites is the getSiteTopologyResponseResponseSites definition
type GetSiteTopologyResponseResponseSites struct {
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

// GetTopologyDetailsResponse is the getTopologyDetailsResponse definition
type GetTopologyDetailsResponse struct {
	Response GetTopologyDetailsResponseResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}

// GetTopologyDetailsResponseResponse is the getTopologyDetailsResponseResponse definition
type GetTopologyDetailsResponseResponse struct {
	ID    string                                    `json:"id,omitempty"`    //
	Links []GetTopologyDetailsResponseResponseLinks `json:"links,omitempty"` //
	Nodes []GetTopologyDetailsResponseResponseNodes `json:"nodes,omitempty"` //
}

// GetTopologyDetailsResponseResponseLinks is the getTopologyDetailsResponseResponseLinks definition
type GetTopologyDetailsResponseResponseLinks struct {
	AdditionalInfo       string `json:"additionalInfo,omitempty"`       //
	EndPortID            string `json:"endPortID,omitempty"`            //
	EndPortIPv4Address   string `json:"endPortIpv4Address,omitempty"`   //
	EndPortIPv4Mask      string `json:"endPortIpv4Mask,omitempty"`      //
	EndPortName          string `json:"endPortName,omitempty"`          //
	EndPortSpeed         string `json:"endPortSpeed,omitempty"`         //
	GreyOut              bool   `json:"greyOut,omitempty"`              //
	ID                   string `json:"id,omitempty"`                   //
	LinkStatus           string `json:"linkStatus,omitempty"`           //
	Source               string `json:"source,omitempty"`               //
	StartPortID          string `json:"startPortID,omitempty"`          //
	StartPortIPv4Address string `json:"startPortIpv4Address,omitempty"` //
	StartPortIPv4Mask    string `json:"startPortIpv4Mask,omitempty"`    //
	StartPortName        string `json:"startPortName,omitempty"`        //
	StartPortSpeed       string `json:"startPortSpeed,omitempty"`       //
	Tag                  string `json:"tag,omitempty"`                  //
	Target               string `json:"target,omitempty"`               //
}

// GetTopologyDetailsResponseResponseNodes is the getTopologyDetailsResponseResponseNodes definition
type GetTopologyDetailsResponseResponseNodes struct {
	ACLApplied      bool                                               `json:"aclApplied,omitempty"`      //
	AdditionalInfo  string                                             `json:"additionalInfo,omitempty"`  //
	CustomParam     GetTopologyDetailsResponseResponseNodesCustomParam `json:"customParam,omitempty"`     //
	DataPathID      string                                             `json:"dataPathId,omitempty"`      //
	DeviceType      string                                             `json:"deviceType,omitempty"`      //
	Family          string                                             `json:"family,omitempty"`          //
	Fixed           bool                                               `json:"fixed,omitempty"`           //
	GreyOut         bool                                               `json:"greyOut,omitempty"`         //
	ID              string                                             `json:"id,omitempty"`              //
	IP              string                                             `json:"ip,omitempty"`              //
	Label           string                                             `json:"label,omitempty"`           //
	NetworkType     string                                             `json:"networkType,omitempty"`     //
	NodeType        string                                             `json:"nodeType,omitempty"`        //
	Order           int                                                `json:"order,omitempty"`           //
	OsType          string                                             `json:"osType,omitempty"`          //
	PlatformID      string                                             `json:"platformId,omitempty"`      //
	Role            string                                             `json:"role,omitempty"`            //
	RoleSource      string                                             `json:"roleSource,omitempty"`      //
	SoftwareVersion string                                             `json:"softwareVersion,omitempty"` //
	Tags            []string                                           `json:"tags,omitempty"`            //
	UpperNode       string                                             `json:"upperNode,omitempty"`       //
	UserID          string                                             `json:"userId,omitempty"`          //
	VLANID          string                                             `json:"vlanId,omitempty"`          //
	X               int                                                `json:"x,omitempty"`               //
	Y               int                                                `json:"y,omitempty"`               //
}

// GetTopologyDetailsResponseResponseNodesCustomParam is the getTopologyDetailsResponseResponseNodesCustomParam definition
type GetTopologyDetailsResponseResponseNodesCustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            int    `json:"x,omitempty"`            //
	Y            int    `json:"y,omitempty"`            //
}

// GetTopologyDetailsResponseResponseNodesTags is the getTopologyDetailsResponseResponseNodesTags definition
type GetTopologyDetailsResponseResponseNodesTags []string

// GetVLANDetailsResponse is the getVLANDetailsResponse definition
type GetVLANDetailsResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetVLANDetailsResponseResponse is the getVLANDetailsResponseResponse definition
type GetVLANDetailsResponseResponse []string

// GetL3TopologyDetails getL3TopologyDetails
/* Returns the Layer 3 network topology by routing protocol
@param topologyType Type of topology(OSPF,ISIS,etc)
*/
func (s *TopologyService) GetL3TopologyDetails(topologyType string) (*GetL3TopologyDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/l3/{topologyType}"
	path = strings.Replace(path, "{"+"topologyType"+"}", fmt.Sprintf("%v", topologyType), -1)

	response, err := RestyClient.R().
		SetResult(&GetL3TopologyDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getL3TopologyDetails")
	}

	result := response.Result().(*GetL3TopologyDetailsResponse)
	return result, response, err
}

// GetOverallNetworkHealthQueryParams defines the query parameters for this request
type GetOverallNetworkHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` // Epoch time(in milliseconds) when the Network health data is required
}

// GetOverallNetworkHealth getOverallNetworkHealth
/* Returns Overall Network Health information by Device category (Access, Distribution, Core, Router, Wireless) for any given point of time
@param timestamp Epoch time(in milliseconds) when the Network health data is required
*/
func (s *TopologyService) GetOverallNetworkHealth(getOverallNetworkHealthQueryParams *GetOverallNetworkHealthQueryParams) (*GetOverallNetworkHealthResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-health"

	queryString, _ := query.Values(getOverallNetworkHealthQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetOverallNetworkHealthResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getOverallNetworkHealth")
	}

	result := response.Result().(*GetOverallNetworkHealthResponse)
	return result, response, err
}

// GetPhysicalTopologyQueryParams defines the query parameters for this request
type GetPhysicalTopologyQueryParams struct {
	NodeType string `url:"nodeType,omitempty"` // nodeType
}

// GetPhysicalTopology getPhysicalTopology
/* Returns the raw physical topology by specified criteria of nodeType
@param nodeType nodeType
*/
func (s *TopologyService) GetPhysicalTopology(getPhysicalTopologyQueryParams *GetPhysicalTopologyQueryParams) (*GetPhysicalTopologyResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/physical-topology"

	queryString, _ := query.Values(getPhysicalTopologyQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetPhysicalTopologyResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getPhysicalTopology")
	}

	result := response.Result().(*GetPhysicalTopologyResponse)
	return result, response, err
}

// GetSiteTopology getSiteTopology
/* Returns site topology
 */
func (s *TopologyService) GetSiteTopology() (*GetSiteTopologyResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/site-topology"

	response, err := RestyClient.R().
		SetResult(&GetSiteTopologyResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getSiteTopology")
	}

	result := response.Result().(*GetSiteTopologyResponse)
	return result, response, err
}

// GetTopologyDetails getTopologyDetails
/* Returns Layer 2 network topology by specified VLAN ID
@param vlanID Vlan Name for e.g Vlan1, Vlan23 etc
*/
func (s *TopologyService) GetTopologyDetails(vlanID string) (*GetTopologyDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/l2/{vlanID}"
	path = strings.Replace(path, "{"+"vlanID"+"}", fmt.Sprintf("%v", vlanID), -1)

	response, err := RestyClient.R().
		SetResult(&GetTopologyDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTopologyDetails")
	}

	result := response.Result().(*GetTopologyDetailsResponse)
	return result, response, err
}

// GetVLANDetails getVLANDetails
/* Returns the list of VLAN names
 */
func (s *TopologyService) GetVLANDetails() (*GetVLANDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/vlan/vlan-names"

	response, err := RestyClient.R().
		SetResult(&GetVLANDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getVLANDetails")
	}

	result := response.Result().(*GetVLANDetailsResponse)
	return result, response, err
}
