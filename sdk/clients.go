package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ClientsService is the service to communicate with the Clients API endpoint
type ClientsService service

// GetClientDetailResponse is the getClientDetailResponse definition
type GetClientDetailResponse struct {
	ConnectionInfo GetClientDetailResponseConnectionInfo `json:"connectionInfo,omitempty"` //
	Detail         GetClientDetailResponseDetail         `json:"detail,omitempty"`         //
	Topology       GetClientDetailResponseTopology       `json:"topology,omitempty"`       //
}

// GetClientDetailResponseConnectionInfo is the getClientDetailResponseConnectionInfo definition
type GetClientDetailResponseConnectionInfo struct {
	Band          string `json:"band,omitempty"`          //
	Channel       string `json:"channel,omitempty"`       //
	ChannelWidth  string `json:"channelWidth,omitempty"`  //
	HostType      string `json:"hostType,omitempty"`      //
	NwDeviceMac   string `json:"nwDeviceMac,omitempty"`   //
	NwDeviceName  string `json:"nwDeviceName,omitempty"`  //
	Protocol      string `json:"protocol,omitempty"`      //
	SpatialStream string `json:"spatialStream,omitempty"` //
	Timestamp     int    `json:"timestamp,omitempty"`     //
	Uapsd         string `json:"uapsd,omitempty"`         //
	Wmm           string `json:"wmm,omitempty"`           //
}

// GetClientDetailResponseDetail is the getClientDetailResponseDetail definition
type GetClientDetailResponseDetail struct {
	ApGroup          string                                     `json:"apGroup,omitempty"`          //
	AuthType         string                                     `json:"authType,omitempty"`         //
	AvgRssi          string                                     `json:"avgRssi,omitempty"`          //
	AvgSnr           string                                     `json:"avgSnr,omitempty"`           //
	Channel          string                                     `json:"channel,omitempty"`          //
	ClientConnection string                                     `json:"clientConnection,omitempty"` //
	ClientType       string                                     `json:"clientType,omitempty"`       //
	ConnectedDevice  []string                                   `json:"connectedDevice,omitempty"`  //
	ConnectionStatus string                                     `json:"connectionStatus,omitempty"` //
	DataRate         string                                     `json:"dataRate,omitempty"`         //
	DNSFailure       string                                     `json:"dnsFailure,omitempty"`       //
	DNSSuccess       string                                     `json:"dnsSuccess,omitempty"`       //
	Frequency        string                                     `json:"frequency,omitempty"`        //
	HealthScore      []GetClientDetailResponseDetailHealthScore `json:"healthScore,omitempty"`      //
	HostIPV4         string                                     `json:"hostIpV4,omitempty"`         //
	HostIPV6         []string                                   `json:"hostIpV6,omitempty"`         //
	HostMac          string                                     `json:"hostMac,omitempty"`          //
	HostName         string                                     `json:"hostName,omitempty"`         //
	HostOs           string                                     `json:"hostOs,omitempty"`           //
	HostType         string                                     `json:"hostType,omitempty"`         //
	HostVersion      string                                     `json:"hostVersion,omitempty"`      //
	ID               string                                     `json:"id,omitempty"`               //
	IosCapable       bool                                       `json:"iosCapable,omitempty"`       //
	IssueCount       int                                        `json:"issueCount,omitempty"`       //
	LastUpdated      int                                        `json:"lastUpdated,omitempty"`      //
	Location         string                                     `json:"location,omitempty"`         //
	Onboarding       GetClientDetailResponseDetailOnboarding    `json:"onboarding,omitempty"`       //
	OnboardingTime   string                                     `json:"onboardingTime,omitempty"`   //
	Port             string                                     `json:"port,omitempty"`             //
	Rssi             string                                     `json:"rssi,omitempty"`             //
	RxBytes          string                                     `json:"rxBytes,omitempty"`          //
	Snr              string                                     `json:"snr,omitempty"`              //
	SSID             string                                     `json:"ssid,omitempty"`             //
	SubType          string                                     `json:"subType,omitempty"`          //
	TxBytes          string                                     `json:"txBytes,omitempty"`          //
	UserID           string                                     `json:"userId,omitempty"`           //
	VLANID           string                                     `json:"vlanId,omitempty"`           //
	Vnid             string                                     `json:"vnid,omitempty"`             //
}

// GetClientDetailResponseDetailConnectedDevice is the getClientDetailResponseDetailConnectedDevice definition
type GetClientDetailResponseDetailConnectedDevice []string

// GetClientDetailResponseDetailHealthScore is the getClientDetailResponseDetailHealthScore definition
type GetClientDetailResponseDetailHealthScore struct {
	HealthType string `json:"healthType,omitempty"` //
	Reason     string `json:"reason,omitempty"`     //
	Score      int    `json:"score,omitempty"`      //
}

// GetClientDetailResponseDetailHostIPV6 is the getClientDetailResponseDetailHostIPV6 definition
type GetClientDetailResponseDetailHostIPV6 []string

// GetClientDetailResponseDetailOnboarding is the getClientDetailResponseDetailOnboarding definition
type GetClientDetailResponseDetailOnboarding struct {
	AAARootcauseList     []string `json:"aaaRootcauseList,omitempty"`     //
	AAAServerIP          string   `json:"aaaServerIp,omitempty"`          //
	AssocDoneTime        string   `json:"assocDoneTime,omitempty"`        //
	AssocRootcauseList   []string `json:"assocRootcauseList,omitempty"`   //
	AuthDoneTime         string   `json:"authDoneTime,omitempty"`         //
	AverageAssocDuration string   `json:"averageAssocDuration,omitempty"` //
	AverageAuthDuration  string   `json:"averageAuthDuration,omitempty"`  //
	AverageDhcpDuration  string   `json:"averageDhcpDuration,omitempty"`  //
	AverageRunDuration   string   `json:"averageRunDuration,omitempty"`   //
	DhcpDoneTime         string   `json:"dhcpDoneTime,omitempty"`         //
	DhcpRootcauseList    []string `json:"dhcpRootcauseList,omitempty"`    //
	DhcpServerIP         string   `json:"dhcpServerIp,omitempty"`         //
	MaxAssocDuration     string   `json:"maxAssocDuration,omitempty"`     //
	MaxAuthDuration      string   `json:"maxAuthDuration,omitempty"`      //
	MaxDhcpDuration      string   `json:"maxDhcpDuration,omitempty"`      //
	MaxRunDuration       string   `json:"maxRunDuration,omitempty"`       //
	OtherRootcauseList   []string `json:"otherRootcauseList,omitempty"`   //
}

// GetClientDetailResponseDetailOnboardingAAARootcauseList is the getClientDetailResponseDetailOnboardingAAARootcauseList definition
type GetClientDetailResponseDetailOnboardingAAARootcauseList []string

// GetClientDetailResponseDetailOnboardingAssocRootcauseList is the getClientDetailResponseDetailOnboardingAssocRootcauseList definition
type GetClientDetailResponseDetailOnboardingAssocRootcauseList []string

// GetClientDetailResponseDetailOnboardingDhcpRootcauseList is the getClientDetailResponseDetailOnboardingDhcpRootcauseList definition
type GetClientDetailResponseDetailOnboardingDhcpRootcauseList []string

// GetClientDetailResponseDetailOnboardingOtherRootcauseList is the getClientDetailResponseDetailOnboardingOtherRootcauseList definition
type GetClientDetailResponseDetailOnboardingOtherRootcauseList []string

// GetClientDetailResponseTopology is the getClientDetailResponseTopology definition
type GetClientDetailResponseTopology struct {
	Links []GetClientDetailResponseTopologyLinks `json:"links,omitempty"` //
	Nodes []GetClientDetailResponseTopologyNodes `json:"nodes,omitempty"` //
}

// GetClientDetailResponseTopologyLinks is the getClientDetailResponseTopologyLinks definition
type GetClientDetailResponseTopologyLinks struct {
	ID              string   `json:"id,omitempty"`              //
	Label           []string `json:"label,omitempty"`           //
	LinkStatus      string   `json:"linkStatus,omitempty"`      //
	PortUtilization string   `json:"portUtilization,omitempty"` //
	Source          string   `json:"source,omitempty"`          //
	Target          string   `json:"target,omitempty"`          //
}

// GetClientDetailResponseTopologyLinksLabel is the getClientDetailResponseTopologyLinksLabel definition
type GetClientDetailResponseTopologyLinksLabel []string

// GetClientDetailResponseTopologyNodes is the getClientDetailResponseTopologyNodes definition
type GetClientDetailResponseTopologyNodes struct {
	Clients         string  `json:"clients,omitempty"`         //
	ConnectedDevice string  `json:"connectedDevice,omitempty"` //
	Count           string  `json:"count,omitempty"`           //
	Description     string  `json:"description,omitempty"`     //
	DeviceType      string  `json:"deviceType,omitempty"`      //
	FabricGroup     string  `json:"fabricGroup,omitempty"`     //
	Family          string  `json:"family,omitempty"`          //
	HealthScore     int     `json:"healthScore,omitempty"`     //
	ID              string  `json:"id,omitempty"`              //
	IP              string  `json:"ip,omitempty"`              //
	Level           float64 `json:"level,omitempty"`           //
	Name            string  `json:"name,omitempty"`            //
	NodeType        string  `json:"nodeType,omitempty"`        //
	PlatformID      string  `json:"platformId,omitempty"`      //
	RadioFrequency  string  `json:"radioFrequency,omitempty"`  //
	Role            string  `json:"role,omitempty"`            //
	SoftwareVersion string  `json:"softwareVersion,omitempty"` //
	UserID          string  `json:"userId,omitempty"`          //
}

// GetClientEnrichmentDetailsResponse is the getClientEnrichmentDetailsResponse definition
type GetClientEnrichmentDetailsResponse struct {
	ConnectedDevice []GetClientEnrichmentDetailsResponseConnectedDevice `json:"connectedDevice,omitempty"` //
	IssueDetails    GetClientEnrichmentDetailsResponseIssueDetails      `json:"issueDetails,omitempty"`    //
	UserDetails     GetClientEnrichmentDetailsResponseUserDetails       `json:"userDetails,omitempty"`     //
}

// GetClientEnrichmentDetailsResponseConnectedDevice is the getClientEnrichmentDetailsResponseConnectedDevice definition
type GetClientEnrichmentDetailsResponseConnectedDevice struct {
	DeviceDetails GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetails `json:"deviceDetails,omitempty"` //
}

// GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetails is the getClientEnrichmentDetailsResponseConnectedDeviceDeviceDetails definition
type GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetails struct {
	ApManagerInterfaceIP      string                                                                           `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIP           string                                                                           `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string                                                                           `json:"bootDateTime,omitempty"`              //
	Cisco360view              string                                                                           `json:"cisco360view,omitempty"`              //
	CollectionInterval        string                                                                           `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string                                                                           `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string                                                                           `json:"errorCode,omitempty"`                 //
	ErrorDescription          string                                                                           `json:"errorDescription,omitempty"`          //
	Family                    string                                                                           `json:"family,omitempty"`                    //
	Hostname                  string                                                                           `json:"hostname,omitempty"`                  //
	ID                        string                                                                           `json:"id,omitempty"`                        //
	InstanceUUID              string                                                                           `json:"instanceUuid,omitempty"`              //
	InterfaceCount            int                                                                              `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string                                                                           `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            int                                                                              `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string                                                                           `json:"lastUpdated,omitempty"`               //
	LineCardCount             int                                                                              `json:"lineCardCount,omitempty"`             //
	LineCardID                string                                                                           `json:"lineCardId,omitempty"`                //
	Location                  string                                                                           `json:"location,omitempty"`                  //
	LocationName              string                                                                           `json:"locationName,omitempty"`              //
	MacAddress                string                                                                           `json:"macAddress,omitempty"`                //
	ManagementIPAddress       string                                                                           `json:"managementIpAddress,omitempty"`       //
	MemorySize                string                                                                           `json:"memorySize,omitempty"`                //
	NeighborTopology          []GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
	PlatformID                string                                                                           `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string                                                                           `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string                                                                           `json:"reachabilityStatus,omitempty"`        //
	Role                      string                                                                           `json:"role,omitempty"`                      //
	RoleSource                string                                                                           `json:"roleSource,omitempty"`                //
	SerialNumber              string                                                                           `json:"serialNumber,omitempty"`              //
	Series                    string                                                                           `json:"series,omitempty"`                    //
	SNMPContact               string                                                                           `json:"snmpContact,omitempty"`               //
	SNMPLocation              string                                                                           `json:"snmpLocation,omitempty"`              //
	SoftwareVersion           string                                                                           `json:"softwareVersion,omitempty"`           //
	TagCount                  int                                                                              `json:"tagCount,omitempty"`                  //
	TunnelUDPPort             string                                                                           `json:"tunnelUdpPort,omitempty"`             //
	Type                      string                                                                           `json:"type,omitempty"`                      //
	UpTime                    int                                                                              `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string                                                                           `json:"waasDeviceMode,omitempty"`            //
}

// GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopology is the getClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopology definition
type GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopology struct {
	Links []GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
	Nodes []GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //
}

// GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinks is the getClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinks definition
type GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinks struct {
	ID              string   `json:"id,omitempty"`              //
	Label           []string `json:"label,omitempty"`           //
	LinkStatus      string   `json:"linkStatus,omitempty"`      //
	PortUtilization string   `json:"portUtilization,omitempty"` //
	Source          string   `json:"source,omitempty"`          //
	Target          string   `json:"target,omitempty"`          //
}

// GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel is the getClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel definition
type GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel []string

// GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyNodes is the getClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyNodes definition
type GetClientEnrichmentDetailsResponseConnectedDeviceDeviceDetailsNeighborTopologyNodes struct {
	Clients         float64 `json:"clients,omitempty"`         //
	Count           string  `json:"count,omitempty"`           //
	Description     string  `json:"description,omitempty"`     //
	DeviceType      string  `json:"deviceType,omitempty"`      //
	FabricGroup     string  `json:"fabricGroup,omitempty"`     //
	Family          string  `json:"family,omitempty"`          //
	HealthScore     string  `json:"healthScore,omitempty"`     //
	ID              string  `json:"id,omitempty"`              //
	IP              string  `json:"ip,omitempty"`              //
	Level           float64 `json:"level,omitempty"`           //
	Name            string  `json:"name,omitempty"`            //
	NodeType        string  `json:"nodeType,omitempty"`        //
	PlatformID      string  `json:"platformId,omitempty"`      //
	RadioFrequency  string  `json:"radioFrequency,omitempty"`  //
	Role            string  `json:"role,omitempty"`            //
	SoftwareVersion string  `json:"softwareVersion,omitempty"` //
	UserID          string  `json:"userId,omitempty"`          //
}

// GetClientEnrichmentDetailsResponseIssueDetails is the getClientEnrichmentDetailsResponseIssueDetails definition
type GetClientEnrichmentDetailsResponseIssueDetails struct {
	Issue []GetClientEnrichmentDetailsResponseIssueDetailsIssue `json:"issue,omitempty"` //
}

// GetClientEnrichmentDetailsResponseIssueDetailsIssue is the getClientEnrichmentDetailsResponseIssueDetailsIssue definition
type GetClientEnrichmentDetailsResponseIssueDetailsIssue struct {
	ImpactedHosts    []GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    //
	IssueCategory    string                                                                `json:"issueCategory,omitempty"`    //
	IssueDescription string                                                                `json:"issueDescription,omitempty"` //
	IssueEntity      string                                                                `json:"issueEntity,omitempty"`      //
	IssueEntityValue string                                                                `json:"issueEntityValue,omitempty"` //
	IssueID          string                                                                `json:"issueId,omitempty"`          //
	IssueName        string                                                                `json:"issueName,omitempty"`        //
	IssuePriority    string                                                                `json:"issuePriority,omitempty"`    //
	IssueSeverity    string                                                                `json:"issueSeverity,omitempty"`    //
	IssueSource      string                                                                `json:"issueSource,omitempty"`      //
	IssueSummary     string                                                                `json:"issueSummary,omitempty"`     //
	IssueTimestamp   int                                                                   `json:"issueTimestamp,omitempty"`   //
	SuggestedActions []GetClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
}

// GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts is the getClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts definition
type GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts struct {
	ConnectedInterface string                                                                   `json:"connectedInterface,omitempty"` //
	FailedAttempts     int                                                                      `json:"failedAttempts,omitempty"`     //
	HostName           string                                                                   `json:"hostName,omitempty"`           //
	HostOs             string                                                                   `json:"hostOs,omitempty"`             //
	HostType           string                                                                   `json:"hostType,omitempty"`           //
	Location           GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocation `json:"location,omitempty"`           //
	MacAddress         string                                                                   `json:"macAddress,omitempty"`         //
	SSID               string                                                                   `json:"ssid,omitempty"`               //
	Timestamp          int                                                                      `json:"timestamp,omitempty"`          //
}

// GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocation is the getClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocation definition
type GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocation struct {
	ApsImpacted []string `json:"apsImpacted,omitempty"` //
	Area        string   `json:"area,omitempty"`        //
	Building    string   `json:"building,omitempty"`    //
	Floor       string   `json:"floor,omitempty"`       //
	SiteID      string   `json:"siteId,omitempty"`      //
	SiteType    string   `json:"siteType,omitempty"`    //
}

// GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocationApsImpacted is the getClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocationApsImpacted definition
type GetClientEnrichmentDetailsResponseIssueDetailsIssueImpactedHostsLocationApsImpacted []string

// GetClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions is the getClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions definition
type GetClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions struct {
	Message string   `json:"message,omitempty"` //
	Steps   []string `json:"steps,omitempty"`   //
}

// GetClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActionsSteps is the getClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActionsSteps definition
type GetClientEnrichmentDetailsResponseIssueDetailsIssueSuggestedActionsSteps []string

// GetClientEnrichmentDetailsResponseUserDetails is the getClientEnrichmentDetailsResponseUserDetails definition
type GetClientEnrichmentDetailsResponseUserDetails struct {
	AuthType         string                                                     `json:"authType,omitempty"`         //
	ClientConnection string                                                     `json:"clientConnection,omitempty"` //
	ConnectedDevice  []string                                                   `json:"connectedDevice,omitempty"`  //
	ConnectionStatus string                                                     `json:"connectionStatus,omitempty"` //
	DataRate         string                                                     `json:"dataRate,omitempty"`         //
	HealthScore      []GetClientEnrichmentDetailsResponseUserDetailsHealthScore `json:"healthScore,omitempty"`      //
	HostIPV4         string                                                     `json:"hostIpV4,omitempty"`         //
	HostIPV6         []string                                                   `json:"hostIpV6,omitempty"`         //
	HostMac          string                                                     `json:"hostMac,omitempty"`          //
	HostName         string                                                     `json:"hostName,omitempty"`         //
	HostOs           string                                                     `json:"hostOs,omitempty"`           //
	HostType         string                                                     `json:"hostType,omitempty"`         //
	HostVersion      string                                                     `json:"hostVersion,omitempty"`      //
	ID               string                                                     `json:"id,omitempty"`               //
	IssueCount       int                                                        `json:"issueCount,omitempty"`       //
	LastUpdated      int                                                        `json:"lastUpdated,omitempty"`      //
	Location         string                                                     `json:"location,omitempty"`         //
	Port             string                                                     `json:"port,omitempty"`             //
	Rssi             string                                                     `json:"rssi,omitempty"`             //
	Snr              string                                                     `json:"snr,omitempty"`              //
	SSID             string                                                     `json:"ssid,omitempty"`             //
	SubType          string                                                     `json:"subType,omitempty"`          //
	UserID           string                                                     `json:"userId,omitempty"`           //
	VLANID           string                                                     `json:"vlanId,omitempty"`           //
}

// GetClientEnrichmentDetailsResponseUserDetailsConnectedDevice is the getClientEnrichmentDetailsResponseUserDetailsConnectedDevice definition
type GetClientEnrichmentDetailsResponseUserDetailsConnectedDevice []string

// GetClientEnrichmentDetailsResponseUserDetailsHealthScore is the getClientEnrichmentDetailsResponseUserDetailsHealthScore definition
type GetClientEnrichmentDetailsResponseUserDetailsHealthScore struct {
	HealthType string `json:"healthType,omitempty"` //
	Reason     string `json:"reason,omitempty"`     //
	Score      int    `json:"score,omitempty"`      //
}

// GetClientEnrichmentDetailsResponseUserDetailsHostIPV6 is the getClientEnrichmentDetailsResponseUserDetailsHostIPV6 definition
type GetClientEnrichmentDetailsResponseUserDetailsHostIPV6 []string

// GetOverallClientHealthResponse is the getOverallClientHealthResponse definition
type GetOverallClientHealthResponse struct {
	Response []GetOverallClientHealthResponseResponse `json:"response,omitempty"` //
}

// GetOverallClientHealthResponseResponse is the getOverallClientHealthResponseResponse definition
type GetOverallClientHealthResponseResponse struct {
	ScoreDetail []GetOverallClientHealthResponseResponseScoreDetail `json:"scoreDetail,omitempty"` //
	SiteID      string                                              `json:"siteId,omitempty"`      //
}

// GetOverallClientHealthResponseResponseScoreDetail is the getOverallClientHealthResponseResponseScoreDetail definition
type GetOverallClientHealthResponseResponseScoreDetail struct {
	ClientCount       int                                                            `json:"clientCount,omitempty"`       //
	ClientUniqueCount int                                                            `json:"clientUniqueCount,omitempty"` //
	Endtime           int                                                            `json:"endtime,omitempty"`           //
	ScoreCategory     GetOverallClientHealthResponseResponseScoreDetailScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreList         []GetOverallClientHealthResponseResponseScoreDetailScoreList   `json:"scoreList,omitempty"`         //
	ScoreValue        int                                                            `json:"scoreValue,omitempty"`        //
	Starttime         int                                                            `json:"starttime,omitempty"`         //
}

// GetOverallClientHealthResponseResponseScoreDetailScoreCategory is the getOverallClientHealthResponseResponseScoreDetailScoreCategory definition
type GetOverallClientHealthResponseResponseScoreDetailScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` //
	Value         string `json:"value,omitempty"`         //
}

// GetOverallClientHealthResponseResponseScoreDetailScoreList is the getOverallClientHealthResponseResponseScoreDetailScoreList definition
type GetOverallClientHealthResponseResponseScoreDetailScoreList struct {
	ClientCount       int                                                                     `json:"clientCount,omitempty"`       //
	ClientUniqueCount int                                                                     `json:"clientUniqueCount,omitempty"` //
	Endtime           int                                                                     `json:"endtime,omitempty"`           //
	ScoreCategory     GetOverallClientHealthResponseResponseScoreDetailScoreListScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreList         []GetOverallClientHealthResponseResponseScoreDetailScoreListScoreList   `json:"scoreList,omitempty"`         //
	ScoreValue        int                                                                     `json:"scoreValue,omitempty"`        //
	Starttime         int                                                                     `json:"starttime,omitempty"`         //
}

// GetOverallClientHealthResponseResponseScoreDetailScoreListScoreCategory is the getOverallClientHealthResponseResponseScoreDetailScoreListScoreCategory definition
type GetOverallClientHealthResponseResponseScoreDetailScoreListScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` //
	Value         string `json:"value,omitempty"`         //
}

// GetOverallClientHealthResponseResponseScoreDetailScoreListScoreList is the getOverallClientHealthResponseResponseScoreDetailScoreListScoreList definition
type GetOverallClientHealthResponseResponseScoreDetailScoreListScoreList struct {
	ClientCount       int                                                                              `json:"clientCount,omitempty"`       //
	ClientUniqueCount int                                                                              `json:"clientUniqueCount,omitempty"` //
	Endtime           int                                                                              `json:"endtime,omitempty"`           //
	ScoreCategory     GetOverallClientHealthResponseResponseScoreDetailScoreListScoreListScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreValue        int                                                                              `json:"scoreValue,omitempty"`        //
	Starttime         int                                                                              `json:"starttime,omitempty"`         //
}

// GetOverallClientHealthResponseResponseScoreDetailScoreListScoreListScoreCategory is the getOverallClientHealthResponseResponseScoreDetailScoreListScoreListScoreCategory definition
type GetOverallClientHealthResponseResponseScoreDetailScoreListScoreListScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` //
	Value         string `json:"value,omitempty"`         //
}

// GetClientDetailQueryParams defines the query parameters for this request
type GetClientDetailQueryParams struct {
	Timestamp  string `url:"timestamp,omitempty"`  // Epoch time(in milliseconds) when the Client health data is required
	MacAddress string `url:"macAddress,omitempty"` // MAC Address of the client
}

// GetClientDetail getClientDetail
/* Returns detailed Client information retrieved by Mac Address for any given point of time.
@param timestamp Epoch time(in milliseconds) when the Client health data is required
@param macAddress MAC Address of the client
*/
func (s *ClientsService) GetClientDetail(getClientDetailQueryParams *GetClientDetailQueryParams) (*GetClientDetailResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-detail"

	queryString, _ := query.Values(getClientDetailQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetClientDetailResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getClientDetail")
	}

	result := response.Result().(*GetClientDetailResponse)
	return result, response, err
}

// GetClientEnrichmentDetails getClientEnrichmentDetails
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user, the devices that the user is connected to and the assurance issues that the user is impacted by
@param entity_type Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
@param entity_value Contains the actual value for the entity type that has been defined
@param issueCategory The category of the DNA event based on which the underlying issues need to be fetched
*/
func (s *ClientsService) GetClientEnrichmentDetails() (*GetClientEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetClientEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getClientEnrichmentDetails")
	}

	result := response.Result().(*GetClientEnrichmentDetailsResponse)
	return result, response, err
}

// GetOverallClientHealthQueryParams defines the query parameters for this request
type GetOverallClientHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` // Epoch time(in milliseconds) when the Client health data is required
}

// GetOverallClientHealth getOverallClientHealth
/* Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time
@param timestamp Epoch time(in milliseconds) when the Client health data is required
*/
func (s *ClientsService) GetOverallClientHealth(getOverallClientHealthQueryParams *GetOverallClientHealthQueryParams) (*GetOverallClientHealthResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/client-health"

	queryString, _ := query.Values(getOverallClientHealthQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetOverallClientHealthResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getOverallClientHealth")
	}

	result := response.Result().(*GetOverallClientHealthResponse)
	return result, response, err
}
