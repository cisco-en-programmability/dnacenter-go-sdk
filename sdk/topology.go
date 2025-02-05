package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TopologyService service

type GetOverallNetworkHealthQueryParams struct {
	Timestamp float64 `url:"timestamp,omitempty"` //UTC timestamp of network health data in milliseconds
}
type GetPhysicalTopologyQueryParams struct {
	NodeType string `url:"nodeType,omitempty"` //nodeType
}

type ResponseTopologyGetOverallNetworkHealth struct {
	Version string `json:"version,omitempty"` // This output's version string

	Response *[]ResponseTopologyGetOverallNetworkHealthResponse `json:"response,omitempty"` //

	MeasuredBy string `json:"measuredBy,omitempty"` // Overall network health measure by 'global'

	LatestMeasuredByEntity string `json:"latestMeasuredByEntity,omitempty"` // Latest measured by entity

	LatestHealthScore *int `json:"latestHealthScore,omitempty"` // Latest health score value

	MonitoredDevices *int `json:"monitoredDevices,omitempty"` // Number of monitored devices

	MonitoredHealthyDevices *int `json:"monitoredHealthyDevices,omitempty"` // Number of healthy devices

	MonitoredUnHealthyDevices *int `json:"monitoredUnHealthyDevices,omitempty"` // Number of unhealthy devices

	UnMonitoredDevices *int `json:"unMonitoredDevices,omitempty"` // Number of un-monitored devices

	NoHealthDevices *int `json:"noHealthDevices,omitempty"` // Number of un-monitored devices

	TotalDevices *int `json:"totalDevices,omitempty"` // Total number of devices

	MonitoredPoorHealthDevices *int `json:"monitoredPoorHealthDevices,omitempty"` // Number of poor health devices

	MonitoredFairHealthDevices *int `json:"monitoredFairHealthDevices,omitempty"` // Number of fair health devices

	HealthContributingDevices *int `json:"healthContributingDevices,omitempty"` // Number of health contributing devices

	HealthDistirubution *[]ResponseTopologyGetOverallNetworkHealthHealthDistirubution `json:"healthDistirubution,omitempty"` //
}
type ResponseTopologyGetOverallNetworkHealthResponse struct {
	Time string `json:"time,omitempty"` // Date-time string

	HealthScore *int `json:"healthScore,omitempty"` // Health score

	TotalCount *int `json:"totalCount,omitempty"` // Total health count

	GoodCount *int `json:"goodCount,omitempty"` // Total good health count

	NoHealthCount *int `json:"noHealthCount,omitempty"` // Total no health count

	UnmonCount *int `json:"unmonCount,omitempty"` // Total no health count

	FairCount *int `json:"fairCount,omitempty"` // Total fair health count

	BadCount *int `json:"badCount,omitempty"` // Total bad health count

	MaintenanceModeCount *int `json:"maintenanceModeCount,omitempty"` // Total maintenance mode count

	Entity string `json:"entity,omitempty"` // Entity of the health data

	TimeinMillis *int `json:"timeinMillis,omitempty"` // UTC time value of property 'time' in milliseconds
}
type ResponseTopologyGetOverallNetworkHealthHealthDistirubution struct {
	Category string `json:"category,omitempty"` // Device category in this health data

	TotalCount *int `json:"totalCount,omitempty"` // Total device count

	HealthScore *int `json:"healthScore,omitempty"` // Health score

	GoodPercentage *float64 `json:"goodPercentage,omitempty"` // Good health percent

	BadPercentage *float64 `json:"badPercentage,omitempty"` // Poor health percent

	FairPercentage *float64 `json:"fairPercentage,omitempty"` // Fair health percent

	NoHealthPercentage *float64 `json:"noHealthPercentage,omitempty"` // No health percent

	UnmonPercentage *float64 `json:"unmonPercentage,omitempty"` // No health percent

	GoodCount *float64 `json:"goodCount,omitempty"` // Good health count

	BadCount *float64 `json:"badCount,omitempty"` // Poor health count

	FairCount *float64 `json:"fairCount,omitempty"` // Fair health count

	NoHealthCount *float64 `json:"noHealthCount,omitempty"` // No health count

	UnmonCount *float64 `json:"unmonCount,omitempty"` // No health count

	ThirdPartyDeviceCount *float64 `json:"thirdPartyDeviceCount,omitempty"` // Third party device count

	KpiMetrics *[]ResponseTopologyGetOverallNetworkHealthHealthDistirubutionKpiMetrics `json:"kpiMetrics,omitempty"` //
}
type ResponseTopologyGetOverallNetworkHealthHealthDistirubutionKpiMetrics struct {
	Key string `json:"key,omitempty"` // Health key

	Value string `json:"value,omitempty"` // Health value
}
type ResponseTopologyGetTopologyDetails struct {
	Response *ResponseTopologyGetTopologyDetailsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTopologyGetTopologyDetailsResponse struct {
	ID string `json:"id,omitempty"` // [Deprecated]

	Links *[]ResponseTopologyGetTopologyDetailsResponseLinks `json:"links,omitempty"` //

	Nodes *[]ResponseTopologyGetTopologyDetailsResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetTopologyDetailsResponseLinks struct {
	AdditionalInfo *ResponseTopologyGetTopologyDetailsResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the link

	EndPortID string `json:"endPortID,omitempty"` // Device port ID corresponding to the end device

	EndPortIPv4Address string `json:"endPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to the end device

	EndPortIPv4Mask string `json:"endPortIpv4Mask,omitempty"` // Interface port IPv4 mask corresponding to the end device

	EndPortName string `json:"endPortName,omitempty"` // Interface port name corresponding to the end device

	EndPortSpeed string `json:"endPortSpeed,omitempty"` // Interface port speed corresponding to end device

	GreyOut *bool `json:"greyOut,omitempty"` // Indicates if the link is greyed out

	ID string `json:"id,omitempty"` // Id of the link

	LinkStatus string `json:"linkStatus,omitempty"` // Indicates whether link is up or down

	Source string `json:"source,omitempty"` // Device ID corresponding to the source device

	StartPortID string `json:"startPortID,omitempty"` // Device port ID corresponding to start device

	StartPortIPv4Address string `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start device

	StartPortIPv4Mask string `json:"startPortIpv4Mask,omitempty"` // Interface port IPv4 mask corresponding to start device

	StartPortName string `json:"startPortName,omitempty"` // Interface port name corresponding to start device

	StartPortSpeed string `json:"startPortSpeed,omitempty"` // Interface port speed corresponding to start device

	Tag string `json:"tag,omitempty"` // [Deprecated]

	Target string `json:"target,omitempty"` // Device ID corresponding to the target device
}
type ResponseTopologyGetTopologyDetailsResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetTopologyDetailsResponseNodes struct {
	ACLApplied *bool `json:"aclApplied,omitempty"` // Indicates if the Access Control List (ACL) is applied on the device

	AdditionalInfo *ResponseTopologyGetTopologyDetailsResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the node

	CustomParam *ResponseTopologyGetTopologyDetailsResponseNodesCustomParam `json:"customParam,omitempty"` //

	ConnectedDeviceID string `json:"connectedDeviceId,omitempty"` // ID of the connected device when the nodeType is HOST

	DataPathID string `json:"dataPathId,omitempty"` // ID of the path between devices

	DeviceType string `json:"deviceType,omitempty"` // Type of the device.

	DeviceSeries string `json:"deviceSeries,omitempty"` // The series of the device

	Family string `json:"family,omitempty"` // The product family of the device

	Fixed *bool `json:"fixed,omitempty"` // Boolean value indicating whether the position is fixed or will use auto layout

	GreyOut *bool `json:"greyOut,omitempty"` // Boolean value indicating whether the node is active for the topology view.

	ID string `json:"id,omitempty"` // Unique identifier for the device

	IP string `json:"ip,omitempty"` // IP address of the device

	Label string `json:"label,omitempty"` // Label of the node, typically the hostname of the device

	NetworkType string `json:"networkType,omitempty"` // Type of the network

	NodeType string `json:"nodeType,omitempty"` // Type of the node can be 'device' or 'HOST'

	Order *int `json:"order,omitempty"` // Device order by link number

	OsType string `json:"osType,omitempty"` // OS type of the device

	PlatformID string `json:"platformId,omitempty"` // Platform description of the device

	Role string `json:"role,omitempty"` // Role of the device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the role is assigned manually or automatically

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device OS version

	Tags []string `json:"tags,omitempty"` // [Deprecated]

	UpperNode string `json:"upperNode,omitempty"` // ID of the start node

	UserID string `json:"userId,omitempty"` // ID of the host

	VLANID string `json:"vlanId,omitempty"` // VLAN ID

	X *int `json:"x,omitempty"` // [Deprecated] Please refer to customParam.x

	Y *int `json:"y,omitempty"` // [Deprecated] Please refer to customerParam.y
}
type ResponseTopologyGetTopologyDetailsResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetTopologyDetailsResponseNodesCustomParam struct {
	ID string `json:"id,omitempty"` // [Deprecated] Please refer to nodes.id

	Label string `json:"label,omitempty"` // Label of the node

	ParentNodeID string `json:"parentNodeId,omitempty"` // Id of the parent node

	X *int `json:"x,omitempty"` // X coordinate for this node in the topology view

	Y *int `json:"y,omitempty"` // Y coordinate for this node in the topology view
}
type ResponseTopologyGetL3TopologyDetails struct {
	Response *ResponseTopologyGetL3TopologyDetailsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTopologyGetL3TopologyDetailsResponse struct {
	ID string `json:"id,omitempty"` // [Deprecated]

	Links *[]ResponseTopologyGetL3TopologyDetailsResponseLinks `json:"links,omitempty"` //

	Nodes *[]ResponseTopologyGetL3TopologyDetailsResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetL3TopologyDetailsResponseLinks struct {
	AdditionalInfo *ResponseTopologyGetL3TopologyDetailsResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the link

	EndPortID string `json:"endPortID,omitempty"` // Device port ID corresponding to the end device

	EndPortIPv4Address string `json:"endPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to the end device

	EndPortIPv4Mask string `json:"endPortIpv4Mask,omitempty"` // Interface port IPv4 mask corresponding to the end device

	EndPortName string `json:"endPortName,omitempty"` // Interface port name corresponding to the end device

	EndPortSpeed string `json:"endPortSpeed,omitempty"` // Interface port speed corresponding to end device

	GreyOut *bool `json:"greyOut,omitempty"` // Indicates if the link is greyed out

	ID string `json:"id,omitempty"` // Id of the link

	LinkStatus string `json:"linkStatus,omitempty"` // Indicates whether link is up or down

	Source string `json:"source,omitempty"` // Device ID corresponding to the source device

	StartPortID string `json:"startPortID,omitempty"` // Device port ID corresponding to start device

	StartPortIPv4Address string `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start device

	StartPortIPv4Mask string `json:"startPortIpv4Mask,omitempty"` // Interface port IPv4 mask corresponding to start device

	StartPortName string `json:"startPortName,omitempty"` // Interface port name corresponding to start device

	StartPortSpeed string `json:"startPortSpeed,omitempty"` // Interface port speed corresponding to start device

	Tag string `json:"tag,omitempty"` // [Deprecated]

	Target string `json:"target,omitempty"` // Device ID corresponding to the target device
}
type ResponseTopologyGetL3TopologyDetailsResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetL3TopologyDetailsResponseNodes struct {
	ACLApplied *bool `json:"aclApplied,omitempty"` // Indicates if the Access Control List (ACL) is applied on the device

	AdditionalInfo *ResponseTopologyGetL3TopologyDetailsResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the node

	CustomParam *ResponseTopologyGetL3TopologyDetailsResponseNodesCustomParam `json:"customParam,omitempty"` //

	ConnectedDeviceID string `json:"connectedDeviceId,omitempty"` // ID of the connected device when the nodeType is HOST

	DataPathID string `json:"dataPathId,omitempty"` // ID of the path between devices

	DeviceType string `json:"deviceType,omitempty"` // Type of the device.

	DeviceSeries string `json:"deviceSeries,omitempty"` // The series of the device

	Family string `json:"family,omitempty"` // The product family of the device

	Fixed *bool `json:"fixed,omitempty"` // Boolean value indicating whether the position is fixed or will use auto layout

	GreyOut *bool `json:"greyOut,omitempty"` // Boolean value indicating whether the node is active for the topology view.

	ID string `json:"id,omitempty"` // Unique identifier for the device

	IP string `json:"ip,omitempty"` // IP address of the device

	Label string `json:"label,omitempty"` // Label of the node, typically the hostname of the device

	NetworkType string `json:"networkType,omitempty"` // Type of the network

	NodeType string `json:"nodeType,omitempty"` // Type of the node can be 'device' or 'HOST'

	Order *int `json:"order,omitempty"` // Device order by link number

	OsType string `json:"osType,omitempty"` // OS type of the device

	PlatformID string `json:"platformId,omitempty"` // Platform description of the device

	Role string `json:"role,omitempty"` // Role of the device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the role is assigned manually or automatically

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device OS version

	Tags []string `json:"tags,omitempty"` // [Deprecated]

	UpperNode string `json:"upperNode,omitempty"` // ID of the start node

	UserID string `json:"userId,omitempty"` // ID of the host

	VLANID string `json:"vlanId,omitempty"` // VLAN ID

	X *int `json:"x,omitempty"` // [Deprecated] Please refer to customParam.x

	Y *int `json:"y,omitempty"` // [Deprecated] Please refer to customerParam.y
}
type ResponseTopologyGetL3TopologyDetailsResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetL3TopologyDetailsResponseNodesCustomParam struct {
	ID string `json:"id,omitempty"` // [Deprecated] Please refer to nodes.id

	Label string `json:"label,omitempty"` // Label of the node

	ParentNodeID string `json:"parentNodeId,omitempty"` // Id of the parent node

	X *int `json:"x,omitempty"` // X coordinate for this node in the topology view

	Y *int `json:"y,omitempty"` // Y coordinate for this node in the topology view
}
type ResponseTopologyGetPhysicalTopology struct {
	Response *ResponseTopologyGetPhysicalTopologyResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTopologyGetPhysicalTopologyResponse struct {
	ID string `json:"id,omitempty"` // [Deprecated]

	Links *[]ResponseTopologyGetPhysicalTopologyResponseLinks `json:"links,omitempty"` //

	Nodes *[]ResponseTopologyGetPhysicalTopologyResponseNodes `json:"nodes,omitempty"` //
}
type ResponseTopologyGetPhysicalTopologyResponseLinks struct {
	AdditionalInfo *ResponseTopologyGetPhysicalTopologyResponseLinksAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the link

	EndPortID string `json:"endPortID,omitempty"` // Device port ID corresponding to the end device

	EndPortIPv4Address string `json:"endPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to the end device

	EndPortIPv4Mask string `json:"endPortIpv4Mask,omitempty"` // Interface port IPv4 mask corresponding to the end device

	EndPortName string `json:"endPortName,omitempty"` // Interface port name corresponding to the end device

	EndPortSpeed string `json:"endPortSpeed,omitempty"` // Interface port speed corresponding to end device

	GreyOut *bool `json:"greyOut,omitempty"` // Indicates if the link is greyed out

	ID string `json:"id,omitempty"` // Id of the link

	LinkStatus string `json:"linkStatus,omitempty"` // Indicates whether link is up or down

	Source string `json:"source,omitempty"` // Device ID corresponding to the source device

	StartPortID string `json:"startPortID,omitempty"` // Device port ID corresponding to start device

	StartPortIPv4Address string `json:"startPortIpv4Address,omitempty"` // Interface port IPv4 address corresponding to start device

	StartPortIPv4Mask string `json:"startPortIpv4Mask,omitempty"` // Interface port IPv4 mask corresponding to start device

	StartPortName string `json:"startPortName,omitempty"` // Interface port name corresponding to start device

	StartPortSpeed string `json:"startPortSpeed,omitempty"` // Interface port speed corresponding to start device

	Tag string `json:"tag,omitempty"` // [Deprecated]

	Target string `json:"target,omitempty"` // Device ID corresponding to the target device
}
type ResponseTopologyGetPhysicalTopologyResponseLinksAdditionalInfo interface{}
type ResponseTopologyGetPhysicalTopologyResponseNodes struct {
	ACLApplied *bool `json:"aclApplied,omitempty"` // Indicates if the Access Control List (ACL) is applied on the device

	AdditionalInfo *ResponseTopologyGetPhysicalTopologyResponseNodesAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the node

	CustomParam *ResponseTopologyGetPhysicalTopologyResponseNodesCustomParam `json:"customParam,omitempty"` //

	ConnectedDeviceID string `json:"connectedDeviceId,omitempty"` // ID of the connected device when the nodeType is HOST

	DataPathID string `json:"dataPathId,omitempty"` // ID of the path between devices

	DeviceType string `json:"deviceType,omitempty"` // Type of the device.

	DeviceSeries string `json:"deviceSeries,omitempty"` // The series of the device

	Family string `json:"family,omitempty"` // The product family of the device

	Fixed *bool `json:"fixed,omitempty"` // Boolean value indicating whether the position is fixed or will use auto layout

	GreyOut *bool `json:"greyOut,omitempty"` // Boolean value indicating whether the node is active for the topology view.

	ID string `json:"id,omitempty"` // Unique identifier for the device

	IP string `json:"ip,omitempty"` // IP address of the device

	Label string `json:"label,omitempty"` // Label of the node, typically the hostname of the device

	NetworkType string `json:"networkType,omitempty"` // Type of the network

	NodeType string `json:"nodeType,omitempty"` // Type of the node can be 'device' or 'HOST'

	Order *int `json:"order,omitempty"` // Device order by link number

	OsType string `json:"osType,omitempty"` // OS type of the device

	PlatformID string `json:"platformId,omitempty"` // Platform description of the device

	Role string `json:"role,omitempty"` // Role of the device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the role is assigned manually or automatically

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device OS version

	Tags []string `json:"tags,omitempty"` // [Deprecated]

	UpperNode string `json:"upperNode,omitempty"` // ID of the start node

	UserID string `json:"userId,omitempty"` // ID of the host

	VLANID string `json:"vlanId,omitempty"` // VLAN ID

	X *int `json:"x,omitempty"` // [Deprecated] Please refer to customParam.x

	Y *int `json:"y,omitempty"` // [Deprecated] Please refer to customerParam.y
}
type ResponseTopologyGetPhysicalTopologyResponseNodesAdditionalInfo interface{}
type ResponseTopologyGetPhysicalTopologyResponseNodesCustomParam struct {
	ID string `json:"id,omitempty"` // [Deprecated] Please refer to nodes.id

	Label string `json:"label,omitempty"` // Label of the node

	ParentNodeID string `json:"parentNodeId,omitempty"` // Id of the parent node

	X *int `json:"x,omitempty"` // X coordinate for this node in the topology view

	Y *int `json:"y,omitempty"` // Y coordinate for this node in the topology view
}
type ResponseTopologyGetSiteTopology struct {
	Response *ResponseTopologyGetSiteTopologyResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseTopologyGetSiteTopologyResponse struct {
	Sites *[]ResponseTopologyGetSiteTopologyResponseSites `json:"sites,omitempty"` //
}
type ResponseTopologyGetSiteTopologyResponseSites struct {
	DisplayName string `json:"displayName,omitempty"` // Group id of the site

	GroupNameHierarchy string `json:"groupNameHierarchy,omitempty"` // Hierarchy of the site names from the root site to the current site. Each site name is separated by a '/'. Eg. 'Global/Site1/Building1/Floor1'

	ID string `json:"id,omitempty"` // Unique identifier of the site

	Latitude string `json:"latitude,omitempty"` // Latitude of the site

	LocationAddress string `json:"locationAddress,omitempty"` // Address of the site

	LocationCountry string `json:"locationCountry,omitempty"` // Country corresponding to the address of the site

	LocationType string `json:"locationType,omitempty"` // Type of site, eg. 'building', 'area' or 'floor'

	Longitude string `json:"longitude,omitempty"` // Longitude of the site

	Name string `json:"name,omitempty"` // Name of the site

	ParentID string `json:"parentId,omitempty"` // Unique identifier of the parent site
}
type ResponseTopologyGetVLANDetails struct {
	Response []string `json:"response,omitempty"` // Lists of all available VLAN names

	Version string `json:"version,omitempty"` //
}

//GetOverallNetworkHealth Get Overall Network Health - 7997-6a34-4409-bfbb
/* Returns Overall Network Health information by Device category (Access, Distribution, Core, Router, Wireless) for any given point of time


@param GetOverallNetworkHealthQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-overall-network-health
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOverallNetworkHealth(GetOverallNetworkHealthQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetOverallNetworkHealth")
	}

	result := response.Result().(*ResponseTopologyGetOverallNetworkHealth)
	return result, response, err

}

//GetTopologyDetails Get topology details - b9b4-8ac8-463a-8aba
/* Returns Layer 2 network topology by specified VLAN ID


@param vlanID vlanID path parameter. Vlan Name for e.g Vlan1, Vlan23 etc


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-topology-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopologyDetails(vlanID)
		}
		return nil, response, fmt.Errorf("error with operation GetTopologyDetails")
	}

	result := response.Result().(*ResponseTopologyGetTopologyDetails)
	return result, response, err

}

//GetL3TopologyDetails Get L3 Topology Details - c2b5-fb76-4d88-8375
/* Returns the Layer 3 network topology by routing protocol


@param topologyType topologyType path parameter. Type of topology(OSPF,ISIS,etc)


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-l3-topology-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetL3TopologyDetails(topologyType)
		}
		return nil, response, fmt.Errorf("error with operation GetL3TopologyDetails")
	}

	result := response.Result().(*ResponseTopologyGetL3TopologyDetails)
	return result, response, err

}

//GetPhysicalTopology Get Physical Topology - b2b8-cb91-459a-a58f
/* Returns the raw physical topology by specified criteria of nodeType


@param GetPhysicalTopologyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-physical-topology
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPhysicalTopology(GetPhysicalTopologyQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPhysicalTopology")
	}

	result := response.Result().(*ResponseTopologyGetPhysicalTopology)
	return result, response, err

}

//GetSiteTopology Get Site Topology - 9ba1-4a9e-441b-8a60
/* Returns site topology



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-topology
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteTopology()
		}
		return nil, response, fmt.Errorf("error with operation GetSiteTopology")
	}

	result := response.Result().(*ResponseTopologyGetSiteTopology)
	return result, response, err

}

//GetVLANDetails Get VLAN details - 6284-db46-49aa-8d31
/* Returns the list of VLAN names that are involved in a loop as identified by the Spanning Tree Protocol



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-vlan-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVLANDetails()
		}
		return nil, response, fmt.Errorf("error with operation GetVlanDetails")
	}

	result := response.Result().(*ResponseTopologyGetVLANDetails)
	return result, response, err

}
