package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// TopologyService is the service to communicate with the Topology API endpoint
type TopologyService service

// CustomParam is the CustomParam definition
type CustomParam struct {
	ID           string `json:"id,omitempty"`           //
	Label        string `json:"label,omitempty"`        //
	ParentNodeID string `json:"parentNodeId,omitempty"` //
	X            int    `json:"x,omitempty"`            //
	Y            int    `json:"y,omitempty"`            //
}

// GetOverallNetworkHealthResponse is the GetOverallNetworkHealthResponse definition
type GetOverallNetworkHealthResponse struct {
	HealthDistirubution []struct {
		BadCount        int      `json:"badCount,omitempty"`        //
		BadPercentage   int      `json:"badPercentage,omitempty"`   //
		Category        string   `json:"category,omitempty"`        //
		FairCount       int      `json:"fairCount,omitempty"`       //
		FairPercentage  int      `json:"fairPercentage,omitempty"`  //
		GoodCount       int      `json:"goodCount,omitempty"`       //
		GoodPercentage  int      `json:"goodPercentage,omitempty"`  //
		HealthScore     int      `json:"healthScore,omitempty"`     //
		KpiMetrics      []string `json:"kpiMetrics,omitempty"`      //
		TotalCount      int      `json:"totalCount,omitempty"`      //
		UnmonCount      int      `json:"unmonCount,omitempty"`      //
		UnmonPercentage int      `json:"unmonPercentage,omitempty"` //
	} `json:"healthDistirubution,omitempty"` //
	LatestHealthScore         int    `json:"latestHealthScore,omitempty"`         //
	LatestMeasuredByEntity    string `json:"latestMeasuredByEntity,omitempty"`    //
	MeasuredBy                string `json:"measuredBy,omitempty"`                //
	MonitoredDevices          int    `json:"monitoredDevices,omitempty"`          //
	MonitoredHealthyDevices   int    `json:"monitoredHealthyDevices,omitempty"`   //
	MonitoredUnHealthyDevices int    `json:"monitoredUnHealthyDevices,omitempty"` //
	Response                  []struct {
		BadCount     int    `json:"badCount,omitempty"`     //
		Entity       string `json:"entity,omitempty"`       //
		FairCount    int    `json:"fairCount,omitempty"`    //
		GoodCount    int    `json:"goodCount,omitempty"`    //
		HealthScore  int    `json:"healthScore,omitempty"`  //
		Time         string `json:"time,omitempty"`         //
		TimeinMillis int    `json:"timeinMillis,omitempty"` //
		TotalCount   int    `json:"totalCount,omitempty"`   //
		UnmonCount   int    `json:"unmonCount,omitempty"`   //
	} `json:"response,omitempty"` //
	UnMonitoredDevices int    `json:"unMonitoredDevices,omitempty"` //
	Version            string `json:"version,omitempty"`            //
}

// HealthDistirubution is the HealthDistirubution definition
type HealthDistirubution struct {
	BadCount        int      `json:"badCount,omitempty"`        //
	BadPercentage   int      `json:"badPercentage,omitempty"`   //
	Category        string   `json:"category,omitempty"`        //
	FairCount       int      `json:"fairCount,omitempty"`       //
	FairPercentage  int      `json:"fairPercentage,omitempty"`  //
	GoodCount       int      `json:"goodCount,omitempty"`       //
	GoodPercentage  int      `json:"goodPercentage,omitempty"`  //
	HealthScore     int      `json:"healthScore,omitempty"`     //
	KpiMetrics      []string `json:"kpiMetrics,omitempty"`      //
	TotalCount      int      `json:"totalCount,omitempty"`      //
	UnmonCount      int      `json:"unmonCount,omitempty"`      //
	UnmonPercentage int      `json:"unmonPercentage,omitempty"` //
}

// Sites is the Sites definition
type Sites struct {
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

// SiteResult is the SiteResult definition
type SiteResult struct {
	Response struct {
		Sites []struct {
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
		} `json:"sites,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// TopologyResult is the TopologyResult definition
type TopologyResult struct {
	Response struct {
		ID    string `json:"id,omitempty"` //
		Links []struct {
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
		} `json:"links,omitempty"` //
		Nodes []struct {
			ACLApplied     bool   `json:"aclApplied,omitempty"`     //
			AdditionalInfo string `json:"additionalInfo,omitempty"` //
			CustomParam    struct {
				ID           string `json:"id,omitempty"`           //
				Label        string `json:"label,omitempty"`        //
				ParentNodeID string `json:"parentNodeId,omitempty"` //
				X            int    `json:"x,omitempty"`            //
				Y            int    `json:"y,omitempty"`            //
			} `json:"customParam,omitempty"` //
			DataPathID      string   `json:"dataPathId,omitempty"`      //
			DeviceType      string   `json:"deviceType,omitempty"`      //
			Family          string   `json:"family,omitempty"`          //
			Fixed           bool     `json:"fixed,omitempty"`           //
			GreyOut         bool     `json:"greyOut,omitempty"`         //
			ID              string   `json:"id,omitempty"`              //
			IP              string   `json:"ip,omitempty"`              //
			Label           string   `json:"label,omitempty"`           //
			NetworkType     string   `json:"networkType,omitempty"`     //
			NodeType        string   `json:"nodeType,omitempty"`        //
			Order           int      `json:"order,omitempty"`           //
			OsType          string   `json:"osType,omitempty"`          //
			PlatformID      string   `json:"platformId,omitempty"`      //
			Role            string   `json:"role,omitempty"`            //
			RoleSource      string   `json:"roleSource,omitempty"`      //
			SoftwareVersion string   `json:"softwareVersion,omitempty"` //
			Tags            []string `json:"tags,omitempty"`            //
			UpperNode       string   `json:"upperNode,omitempty"`       //
			UserID          string   `json:"userId,omitempty"`          //
			VlanID          string   `json:"vlanId,omitempty"`          //
			X               int      `json:"x,omitempty"`               //
			Y               int      `json:"y,omitempty"`               //
		} `json:"nodes,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// VlanNamesResult is the VlanNamesResult definition
type VlanNamesResult struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetL3TopologyDetails getL3TopologyDetails
/* Returns the Layer 3 network topology by routing protocol
@param topologyType Type of topology(OSPF,ISIS,etc)
*/
func (s *TopologyService) GetL3TopologyDetails(topologyType string) (*TopologyResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/l3/{topologyType}"
	path = strings.Replace(path, "{"+"topologyType"+"}", fmt.Sprintf("%v", topologyType), -1)

	response, err := RestyClient.R().
		SetResult(&TopologyResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TopologyResult)
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
func (s *TopologyService) GetPhysicalTopology(getPhysicalTopologyQueryParams *GetPhysicalTopologyQueryParams) (*TopologyResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/physical-topology"

	queryString, _ := query.Values(getPhysicalTopologyQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&TopologyResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TopologyResult)
	return result, response, err

}

// GetSiteTopology getSiteTopology
/* Returns site topology
 */
func (s *TopologyService) GetSiteTopology() (*SiteResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/site-topology"

	response, err := RestyClient.R().
		SetResult(&SiteResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*SiteResult)
	return result, response, err

}

// GetTopologyDetails getTopologyDetails
/* Returns Layer 2 network topology by specified VLAN ID
@param vlanID Vlan Name for e.g Vlan1, Vlan23 etc
*/
func (s *TopologyService) GetTopologyDetails(vlanID string) (*TopologyResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/l2/{vlanID}"
	path = strings.Replace(path, "{"+"vlanID"+"}", fmt.Sprintf("%v", vlanID), -1)

	response, err := RestyClient.R().
		SetResult(&TopologyResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TopologyResult)
	return result, response, err

}

// GetVLANDetails getVLANDetails
/* Returns the list of VLAN names
 */
func (s *TopologyService) GetVLANDetails() (*VlanNamesResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/topology/vlan/vlan-names"

	response, err := RestyClient.R().
		SetResult(&VlanNamesResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*VlanNamesResult)
	return result, response, err

}
