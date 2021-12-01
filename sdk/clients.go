package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ClientsService service

type GetClientDetailQueryParams struct {
	Timestamp  string `url:"timestamp,omitempty"`  //Epoch time(in milliseconds) when the Client health data is required
	MacAddress string `url:"macAddress,omitempty"` //MAC Address of the client
}
type GetClientEnrichmentDetailsHeaderParams struct {
	EntityType    string `url:"entity_type,omitempty"`   //Expects type string. Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
	EntityValue   string `url:"entity_value,omitempty"`  //Expects type string. Contains the actual value for the entity type that has been defined
	IssueCategory string `url:"issueCategory,omitempty"` //Expects type string. The category of the DNA event based on which the underlying issues need to be fetched
}
type GetOverallClientHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Client health data is required
}
type ClientProximityQueryParams struct {
	Username       string  `url:"username,omitempty"`        //Wireless client username for which proximity information is required
	NumberDays     float64 `url:"number_days,omitempty"`     //Number of days to track proximity until current date. Defaults and maximum up to 14 days.
	TimeResolution float64 `url:"time_resolution,omitempty"` //Time interval (in minutes) to measure proximity. Defaults to 15 minutes with a minimum 5 minutes.
}

type ResponseClientsGetClientDetail struct {
	Detail         *ResponseClientsGetClientDetailDetail         `json:"detail,omitempty"`         //
	ConnectionInfo *ResponseClientsGetClientDetailConnectionInfo `json:"connectionInfo,omitempty"` //
	Topology       *ResponseClientsGetClientDetailTopology       `json:"topology,omitempty"`       //
}
type ResponseClientsGetClientDetailDetail struct {
	ID               string                                                 `json:"id,omitempty"`               // Id
	ConnectionStatus string                                                 `json:"connectionStatus,omitempty"` // Connection Status
	HostType         string                                                 `json:"hostType,omitempty"`         // Host Type
	UserID           *ResponseClientsGetClientDetailDetailUserID            `json:"userId,omitempty"`           // User Id
	HostName         string                                                 `json:"hostName,omitempty"`         // Host Name
	HostOs           *ResponseClientsGetClientDetailDetailHostOs            `json:"hostOs,omitempty"`           // Host Os
	HostVersion      *ResponseClientsGetClientDetailDetailHostVersion       `json:"hostVersion,omitempty"`      // Host Version
	SubType          string                                                 `json:"subType,omitempty"`          // Sub Type
	LastUpdated      *int                                                   `json:"lastUpdated,omitempty"`      // Last Updated
	HealthScore      *[]ResponseClientsGetClientDetailDetailHealthScore     `json:"healthScore,omitempty"`      //
	HostMac          string                                                 `json:"hostMac,omitempty"`          // Host Mac
	HostIPV4         string                                                 `json:"hostIpV4,omitempty"`         // Host Ip V4
	HostIPV6         []string                                               `json:"hostIpV6,omitempty"`         // Host Ip V6
	AuthType         string                                                 `json:"authType,omitempty"`         // Auth Type
	VLANID           string                                                 `json:"vlanId,omitempty"`           // Vlan Id
	Vnid             string                                                 `json:"vnid,omitempty"`             // Vnid
	SSID             string                                                 `json:"ssid,omitempty"`             // Ssid
	Frequency        string                                                 `json:"frequency,omitempty"`        // Frequency
	Channel          string                                                 `json:"channel,omitempty"`          // Channel
	ApGroup          *ResponseClientsGetClientDetailDetailApGroup           `json:"apGroup,omitempty"`          // Ap Group
	Location         *ResponseClientsGetClientDetailDetailLocation          `json:"location,omitempty"`         // Location
	ClientConnection string                                                 `json:"clientConnection,omitempty"` // Client Connection
	ConnectedDevice  *[]ResponseClientsGetClientDetailDetailConnectedDevice `json:"connectedDevice,omitempty"`  // Connected Device
	IssueCount       *float64                                               `json:"issueCount,omitempty"`       // Issue Count
	Rssi             string                                                 `json:"rssi,omitempty"`             // Rssi
	AvgRssi          *ResponseClientsGetClientDetailDetailAvgRssi           `json:"avgRssi,omitempty"`          // Avg Rssi
	Snr              string                                                 `json:"snr,omitempty"`              // Snr
	AvgSnr           *ResponseClientsGetClientDetailDetailAvgSnr            `json:"avgSnr,omitempty"`           // Avg Snr
	DataRate         string                                                 `json:"dataRate,omitempty"`         // Data Rate
	TxBytes          string                                                 `json:"txBytes,omitempty"`          // Tx Bytes
	RxBytes          string                                                 `json:"rxBytes,omitempty"`          // Rx Bytes
	DNSSuccess       *ResponseClientsGetClientDetailDetailDNSSuccess        `json:"dnsSuccess,omitempty"`       // Dns Success
	DNSFailure       *ResponseClientsGetClientDetailDetailDNSFailure        `json:"dnsFailure,omitempty"`       // Dns Failure
	Onboarding       *ResponseClientsGetClientDetailDetailOnboarding        `json:"onboarding,omitempty"`       //
	ClientType       string                                                 `json:"clientType,omitempty"`       // Client Type
	OnboardingTime   *ResponseClientsGetClientDetailDetailOnboardingTime    `json:"onboardingTime,omitempty"`   // Onboarding Time
	Port             *ResponseClientsGetClientDetailDetailPort              `json:"port,omitempty"`             // Port
	IosCapable       *bool                                                  `json:"iosCapable,omitempty"`       // Ios Capable
}
type ResponseClientsGetClientDetailDetailUserID interface{}
type ResponseClientsGetClientDetailDetailHostOs interface{}
type ResponseClientsGetClientDetailDetailHostVersion interface{}
type ResponseClientsGetClientDetailDetailHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Health Type
	Reason     string `json:"reason,omitempty"`     // Reason
	Score      *int   `json:"score,omitempty"`      // Score
}
type ResponseClientsGetClientDetailDetailApGroup interface{}
type ResponseClientsGetClientDetailDetailLocation interface{}
type ResponseClientsGetClientDetailDetailConnectedDevice interface{}
type ResponseClientsGetClientDetailDetailAvgRssi interface{}
type ResponseClientsGetClientDetailDetailAvgSnr interface{}
type ResponseClientsGetClientDetailDetailDNSSuccess interface{}
type ResponseClientsGetClientDetailDetailDNSFailure interface{}
type ResponseClientsGetClientDetailDetailOnboarding struct {
	AverageRunDuration   *ResponseClientsGetClientDetailDetailOnboardingAverageRunDuration   `json:"averageRunDuration,omitempty"`   // Average Run Duration
	MaxRunDuration       *ResponseClientsGetClientDetailDetailOnboardingMaxRunDuration       `json:"maxRunDuration,omitempty"`       // Max Run Duration
	AverageAssocDuration *ResponseClientsGetClientDetailDetailOnboardingAverageAssocDuration `json:"averageAssocDuration,omitempty"` // Average Assoc Duration
	MaxAssocDuration     *ResponseClientsGetClientDetailDetailOnboardingMaxAssocDuration     `json:"maxAssocDuration,omitempty"`     // Max Assoc Duration
	AverageAuthDuration  *ResponseClientsGetClientDetailDetailOnboardingAverageAuthDuration  `json:"averageAuthDuration,omitempty"`  // Average Auth Duration
	MaxAuthDuration      *ResponseClientsGetClientDetailDetailOnboardingMaxAuthDuration      `json:"maxAuthDuration,omitempty"`      // Max Auth Duration
	AverageDhcpDuration  *ResponseClientsGetClientDetailDetailOnboardingAverageDhcpDuration  `json:"averageDhcpDuration,omitempty"`  // Average Dhcp Duration
	MaxDhcpDuration      *ResponseClientsGetClientDetailDetailOnboardingMaxDhcpDuration      `json:"maxDhcpDuration,omitempty"`      // Max Dhcp Duration
	AAAServerIP          string                                                              `json:"aaaServerIp,omitempty"`          // Aaa Server Ip
	DhcpServerIP         *ResponseClientsGetClientDetailDetailOnboardingDhcpServerIP         `json:"dhcpServerIp,omitempty"`         // Dhcp Server Ip
	AuthDoneTime         *ResponseClientsGetClientDetailDetailOnboardingAuthDoneTime         `json:"authDoneTime,omitempty"`         // Auth Done Time
	AssocDoneTime        *ResponseClientsGetClientDetailDetailOnboardingAssocDoneTime        `json:"assocDoneTime,omitempty"`        // Assoc Done Time
	DhcpDoneTime         *ResponseClientsGetClientDetailDetailOnboardingDhcpDoneTime         `json:"dhcpDoneTime,omitempty"`         // Dhcp Done Time
	AssocRootcauseList   *[]ResponseClientsGetClientDetailDetailOnboardingAssocRootcauseList `json:"assocRootcauseList,omitempty"`   // Assoc Rootcause List
	AAARootcauseList     *[]ResponseClientsGetClientDetailDetailOnboardingAAARootcauseList   `json:"aaaRootcauseList,omitempty"`     // Aaa Rootcause List
	DhcpRootcauseList    *[]ResponseClientsGetClientDetailDetailOnboardingDhcpRootcauseList  `json:"dhcpRootcauseList,omitempty"`    // Dhcp Rootcause List
	OtherRootcauseList   *[]ResponseClientsGetClientDetailDetailOnboardingOtherRootcauseList `json:"otherRootcauseList,omitempty"`   // Other Rootcause List
}
type ResponseClientsGetClientDetailDetailOnboardingAverageRunDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingMaxRunDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingAverageAssocDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingMaxAssocDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingAverageAuthDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingMaxAuthDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingAverageDhcpDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingMaxDhcpDuration interface{}
type ResponseClientsGetClientDetailDetailOnboardingDhcpServerIP interface{}
type ResponseClientsGetClientDetailDetailOnboardingAuthDoneTime interface{}
type ResponseClientsGetClientDetailDetailOnboardingAssocDoneTime interface{}
type ResponseClientsGetClientDetailDetailOnboardingDhcpDoneTime interface{}
type ResponseClientsGetClientDetailDetailOnboardingAssocRootcauseList interface{}
type ResponseClientsGetClientDetailDetailOnboardingAAARootcauseList interface{}
type ResponseClientsGetClientDetailDetailOnboardingDhcpRootcauseList interface{}
type ResponseClientsGetClientDetailDetailOnboardingOtherRootcauseList interface{}
type ResponseClientsGetClientDetailDetailOnboardingTime interface{}
type ResponseClientsGetClientDetailDetailPort interface{}
type ResponseClientsGetClientDetailConnectionInfo struct {
	HostType      string `json:"hostType,omitempty"`      // Host Type
	NwDeviceName  string `json:"nwDeviceName,omitempty"`  // Nw Device Name
	NwDeviceMac   string `json:"nwDeviceMac,omitempty"`   // Nw Device Mac
	Protocol      string `json:"protocol,omitempty"`      // Protocol
	Band          string `json:"band,omitempty"`          // Band
	SpatialStream string `json:"spatialStream,omitempty"` // Spatial Stream
	Channel       string `json:"channel,omitempty"`       // Channel
	ChannelWidth  string `json:"channelWidth,omitempty"`  // Channel Width
	Wmm           string `json:"wmm,omitempty"`           // Wmm
	Uapsd         string `json:"uapsd,omitempty"`         // Uapsd
	Timestamp     *int   `json:"timestamp,omitempty"`     // Timestamp
}
type ResponseClientsGetClientDetailTopology struct {
	Nodes *[]ResponseClientsGetClientDetailTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseClientsGetClientDetailTopologyLinks `json:"links,omitempty"` //
}
type ResponseClientsGetClientDetailTopologyNodes struct {
	Role            string                                                      `json:"role,omitempty"`            // Role
	Name            string                                                      `json:"name,omitempty"`            // Name
	ID              string                                                      `json:"id,omitempty"`              // Id
	Description     string                                                      `json:"description,omitempty"`     // Description
	DeviceType      string                                                      `json:"deviceType,omitempty"`      // Device Type
	PlatformID      *ResponseClientsGetClientDetailTopologyNodesPlatformID      `json:"platformId,omitempty"`      // Platform Id
	Family          *ResponseClientsGetClientDetailTopologyNodesFamily          `json:"family,omitempty"`          // Family
	IP              string                                                      `json:"ip,omitempty"`              // Ip
	SoftwareVersion *ResponseClientsGetClientDetailTopologyNodesSoftwareVersion `json:"softwareVersion,omitempty"` // Software Version
	UserID          *ResponseClientsGetClientDetailTopologyNodesUserID          `json:"userId,omitempty"`          // User Id
	NodeType        string                                                      `json:"nodeType,omitempty"`        // Node Type
	RadioFrequency  *ResponseClientsGetClientDetailTopologyNodesRadioFrequency  `json:"radioFrequency,omitempty"`  // Radio Frequency
	Clients         *ResponseClientsGetClientDetailTopologyNodesClients         `json:"clients,omitempty"`         // Clients
	Count           *ResponseClientsGetClientDetailTopologyNodesCount           `json:"count,omitempty"`           // Count
	HealthScore     *int                                                        `json:"healthScore,omitempty"`     // Health Score
	Level           *float64                                                    `json:"level,omitempty"`           // Level
	FabricGroup     *ResponseClientsGetClientDetailTopologyNodesFabricGroup     `json:"fabricGroup,omitempty"`     // Fabric Group
	ConnectedDevice *ResponseClientsGetClientDetailTopologyNodesConnectedDevice `json:"connectedDevice,omitempty"` // Connected Device
}
type ResponseClientsGetClientDetailTopologyNodesPlatformID interface{}
type ResponseClientsGetClientDetailTopologyNodesFamily interface{}
type ResponseClientsGetClientDetailTopologyNodesSoftwareVersion interface{}
type ResponseClientsGetClientDetailTopologyNodesUserID interface{}
type ResponseClientsGetClientDetailTopologyNodesRadioFrequency interface{}
type ResponseClientsGetClientDetailTopologyNodesClients interface{}
type ResponseClientsGetClientDetailTopologyNodesCount interface{}
type ResponseClientsGetClientDetailTopologyNodesFabricGroup interface{}
type ResponseClientsGetClientDetailTopologyNodesConnectedDevice interface{}
type ResponseClientsGetClientDetailTopologyLinks struct {
	Source          string                                                      `json:"source,omitempty"`          // Source
	LinkStatus      string                                                      `json:"linkStatus,omitempty"`      // Link Status
	Label           []string                                                    `json:"label,omitempty"`           // Label
	Target          string                                                      `json:"target,omitempty"`          // Target
	ID              *ResponseClientsGetClientDetailTopologyLinksID              `json:"id,omitempty"`              // Id
	PortUtilization *ResponseClientsGetClientDetailTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Port Utilization
}
type ResponseClientsGetClientDetailTopologyLinksID interface{}
type ResponseClientsGetClientDetailTopologyLinksPortUtilization interface{}
type ResponseClientsGetClientEnrichmentDetails []ResponseItemClientsGetClientEnrichmentDetails // Array of ResponseClientsGetClientEnrichmentDetails
type ResponseItemClientsGetClientEnrichmentDetails struct {
	UserDetails     *ResponseItemClientsGetClientEnrichmentDetailsUserDetails       `json:"userDetails,omitempty"`     //
	ConnectedDevice *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDevice `json:"connectedDevice,omitempty"` //
	IssueDetails    *ResponseItemClientsGetClientEnrichmentDetailsIssueDetails      `json:"issueDetails,omitempty"`    //
}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetails struct {
	ID               string                                                                     `json:"id,omitempty"`               // Id
	ConnectionStatus string                                                                     `json:"connectionStatus,omitempty"` // Connection Status
	HostType         string                                                                     `json:"hostType,omitempty"`         // Host Type
	UserID           string                                                                     `json:"userId,omitempty"`           // User Id
	HostName         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostName          `json:"hostName,omitempty"`         // Host Name
	HostOs           *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostOs            `json:"hostOs,omitempty"`           // Host Os
	HostVersion      *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostVersion       `json:"hostVersion,omitempty"`      // Host Version
	SubType          *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSubType           `json:"subType,omitempty"`          // Sub Type
	LastUpdated      *int                                                                       `json:"lastUpdated,omitempty"`      // Last Updated
	HealthScore      *[]ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHealthScore     `json:"healthScore,omitempty"`      //
	HostMac          string                                                                     `json:"hostMac,omitempty"`          // Host Mac
	HostIPV4         string                                                                     `json:"hostIpV4,omitempty"`         // Host Ip V4
	HostIPV6         *[]ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostIPV6        `json:"hostIpV6,omitempty"`         // Host Ip V6
	AuthType         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsAuthType          `json:"authType,omitempty"`         // Auth Type
	VLANID           string                                                                     `json:"vlanId,omitempty"`           // Vlan Id
	SSID             *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSSID              `json:"ssid,omitempty"`             // Ssid
	Location         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsLocation          `json:"location,omitempty"`         // Location
	ClientConnection string                                                                     `json:"clientConnection,omitempty"` // Client Connection
	ConnectedDevice  *[]ResponseItemClientsGetClientEnrichmentDetailsUserDetailsConnectedDevice `json:"connectedDevice,omitempty"`  // Connected Device
	IssueCount       *float64                                                                   `json:"issueCount,omitempty"`       // Issue Count
	Rssi             *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsRssi              `json:"rssi,omitempty"`             // Rssi
	Snr              *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSnr               `json:"snr,omitempty"`              // Snr
	DataRate         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsDataRate          `json:"dataRate,omitempty"`         // Data Rate
	Port             *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsPort              `json:"port,omitempty"`             // Port
}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostName interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostOs interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostVersion interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSubType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Health Type
	Reason     string `json:"reason,omitempty"`     // Reason
	Score      *int   `json:"score,omitempty"`      // Score
}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostIPV6 interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsAuthType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSSID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsLocation interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsConnectedDevice interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsRssi interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSnr interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsDataRate interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsPort interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDevice struct {
	DeviceDetails *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetails struct {
	Family                    string                                                                                       `json:"family,omitempty"`                    // Family
	Type                      string                                                                                       `json:"type,omitempty"`                      // Type
	Location                  *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocation           `json:"location,omitempty"`                  // Location
	ErrorCode                 string                                                                                       `json:"errorCode,omitempty"`                 // Error Code
	MacAddress                string                                                                                       `json:"macAddress,omitempty"`                // Mac Address
	Role                      string                                                                                       `json:"role,omitempty"`                      // Role
	ApManagerInterfaceIP      string                                                                                       `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                                                       `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsBootDateTime       `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                                                       `json:"collectionStatus,omitempty"`          // Collection Status
	InterfaceCount            *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsInterfaceCount     `json:"interfaceCount,omitempty"`            // Interface Count
	LineCardCount             *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardCount      `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardID         `json:"lineCardId,omitempty"`                // Line Card Id
	ManagementIPAddress       string                                                                                       `json:"managementIpAddress,omitempty"`       // Management Ip Address
	MemorySize                string                                                                                       `json:"memorySize,omitempty"`                // Memory Size
	PlatformID                string                                                                                       `json:"platformId,omitempty"`                // Platform Id
	ReachabilityFailureReason string                                                                                       `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                                                       `json:"reachabilityStatus,omitempty"`        // Reachability Status
	SNMPContact               string                                                                                       `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                                                       `json:"snmpLocation,omitempty"`              // Snmp Location
	TunnelUDPPort             string                                                                                       `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	WaasDeviceMode            *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	Series                    string                                                                                       `json:"series,omitempty"`                    // Series
	InventoryStatusDetail     string                                                                                       `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	CollectionInterval        string                                                                                       `json:"collectionInterval,omitempty"`        // Collection Interval
	SerialNumber              string                                                                                       `json:"serialNumber,omitempty"`              // Serial Number
	SoftwareVersion           string                                                                                       `json:"softwareVersion,omitempty"`           // Software Version
	RoleSource                string                                                                                       `json:"roleSource,omitempty"`                // Role Source
	Hostname                  string                                                                                       `json:"hostname,omitempty"`                  // Hostname
	UpTime                    string                                                                                       `json:"upTime,omitempty"`                    // Up Time
	LastUpdateTime            *int                                                                                         `json:"lastUpdateTime,omitempty"`            // Last Update Time
	ErrorDescription          *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription   `json:"errorDescription,omitempty"`          // Error Description
	LocationName              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName       `json:"locationName,omitempty"`              // Location Name
	TagCount                  string                                                                                       `json:"tagCount,omitempty"`                  // Tag Count
	LastUpdated               string                                                                                       `json:"lastUpdated,omitempty"`               // Last Updated
	InstanceUUID              string                                                                                       `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                                                       `json:"id,omitempty"`                        // Id
	NeighborTopology          *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
	Cisco360View              string                                                                                       `json:"cisco360view,omitempty"`              // Cisco360view
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocation interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsBootDateTime interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsInterfaceCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodes struct {
	Role            string                                                                                                         `json:"role,omitempty"`            // Role
	Name            string                                                                                                         `json:"name,omitempty"`            // Name
	ID              string                                                                                                         `json:"id,omitempty"`              // Id
	Description     string                                                                                                         `json:"description,omitempty"`     // Description
	DeviceType      *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType      `json:"deviceType,omitempty"`      // Device Type
	PlatformID      *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID      `json:"platformId,omitempty"`      // Platform Id
	Family          *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily          `json:"family,omitempty"`          // Family
	IP              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP              `json:"ip,omitempty"`              // Ip
	SoftwareVersion *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion `json:"softwareVersion,omitempty"` // Software Version
	UserID          *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID          `json:"userId,omitempty"`          // User Id
	NodeType        *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType        `json:"nodeType,omitempty"`        // Node Type
	RadioFrequency  *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency  `json:"radioFrequency,omitempty"`  // Radio Frequency
	Clients         *float64                                                                                                       `json:"clients,omitempty"`         // Clients
	Count           *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount           `json:"count,omitempty"`           // Count
	HealthScore     *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore     `json:"healthScore,omitempty"`     // Health Score
	Level           *float64                                                                                                       `json:"level,omitempty"`           // Level
	FabricGroup     *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup     `json:"fabricGroup,omitempty"`     // Fabric Group
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinks struct {
	Source          string                                                                                                         `json:"source,omitempty"`          // Source
	LinkStatus      string                                                                                                         `json:"linkStatus,omitempty"`      // Link Status
	Label           *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel         `json:"label,omitempty"`           // Label
	Target          string                                                                                                         `json:"target,omitempty"`          // Target
	ID              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksID              `json:"id,omitempty"`              // Id
	PortUtilization *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Port Utilization
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetails struct {
	Issue *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssue `json:"issue,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssue struct {
	IssueID          string                                                                            `json:"issueId,omitempty"`          // Issue Id
	IssueSource      string                                                                            `json:"issueSource,omitempty"`      // Issue Source
	IssueCategory    string                                                                            `json:"issueCategory,omitempty"`    // Issue Category
	IssueName        string                                                                            `json:"issueName,omitempty"`        // Issue Name
	IssueDescription string                                                                            `json:"issueDescription,omitempty"` // Issue Description
	IssueEntity      string                                                                            `json:"issueEntity,omitempty"`      // Issue Entity
	IssueEntityValue string                                                                            `json:"issueEntityValue,omitempty"` // Issue Entity Value
	IssueSeverity    string                                                                            `json:"issueSeverity,omitempty"`    // Issue Severity
	IssuePriority    string                                                                            `json:"issuePriority,omitempty"`    // Issue Priority
	IssueSummary     string                                                                            `json:"issueSummary,omitempty"`     // Issue Summary
	IssueTimestamp   *int                                                                              `json:"issueTimestamp,omitempty"`   // Issue Timestamp
	SuggestedActions *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
	ImpactedHosts    *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    //
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActions struct {
	Message string                                                                                 `json:"message,omitempty"` // Message
	Steps   *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps interface{}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHosts struct {
	HostType           string                                                                               `json:"hostType,omitempty"`           // Host Type
	HostName           string                                                                               `json:"hostName,omitempty"`           // Host Name
	HostOs             string                                                                               `json:"hostOs,omitempty"`             // Host Os
	SSID               string                                                                               `json:"ssid,omitempty"`               // Ssid
	ConnectedInterface string                                                                               `json:"connectedInterface,omitempty"` // Connected Interface
	MacAddress         string                                                                               `json:"macAddress,omitempty"`         // Mac Address
	FailedAttempts     *int                                                                                 `json:"failedAttempts,omitempty"`     // Failed Attempts
	Location           *ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocation `json:"location,omitempty"`           //
	Timestamp          *int                                                                                 `json:"timestamp,omitempty"`          // Timestamp
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocation struct {
	SiteID      string                                                                                            `json:"siteId,omitempty"`      // Site Id
	SiteType    string                                                                                            `json:"siteType,omitempty"`    // Site Type
	Area        string                                                                                            `json:"area,omitempty"`        // Area
	Building    string                                                                                            `json:"building,omitempty"`    // Building
	Floor       *ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationFloor         `json:"floor,omitempty"`       // Floor
	ApsImpacted *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationApsImpacted `json:"apsImpacted,omitempty"` // Aps Impacted
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationFloor interface{}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationApsImpacted interface{}
type ResponseClientsGetOverallClientHealth struct {
	Response *[]ResponseClientsGetOverallClientHealthResponse `json:"response,omitempty"` //
}
type ResponseClientsGetOverallClientHealthResponse struct {
	SiteID      string                                                      `json:"siteId,omitempty"`      // Site Id
	ScoreDetail *[]ResponseClientsGetOverallClientHealthResponseScoreDetail `json:"scoreDetail,omitempty"` //
}
type ResponseClientsGetOverallClientHealthResponseScoreDetail struct {
	ScoreCategory     *ResponseClientsGetOverallClientHealthResponseScoreDetailScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreValue        *int                                                                   `json:"scoreValue,omitempty"`        // Score Value
	ClientCount       *int                                                                   `json:"clientCount,omitempty"`       // Client Count
	ClientUniqueCount *int                                                                   `json:"clientUniqueCount,omitempty"` // Client Unique Count
	Starttime         *int                                                                   `json:"starttime,omitempty"`         // Starttime
	Endtime           *int                                                                   `json:"endtime,omitempty"`           // Endtime
	ScoreList         *[]ResponseClientsGetOverallClientHealthResponseScoreDetailScoreList   `json:"scoreList,omitempty"`         //
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Score Category
	Value         string `json:"value,omitempty"`         // Value
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreList struct {
	ScoreCategory     *ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreCategory `json:"scoreCategory,omitempty"`     //
	ScoreValue        *int                                                                            `json:"scoreValue,omitempty"`        // Score Value
	ClientCount       *int                                                                            `json:"clientCount,omitempty"`       // Client Count
	ClientUniqueCount *float64                                                                        `json:"clientUniqueCount,omitempty"` // Client Unique Count
	Starttime         *int                                                                            `json:"starttime,omitempty"`         // Starttime
	Endtime           *int                                                                            `json:"endtime,omitempty"`           // Endtime
	ScoreList         *[]ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreList   `json:"scoreList,omitempty"`         //
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Score Category
	Value         string `json:"value,omitempty"`         // Value
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreList struct {
	ScoreCategory     *ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreListScoreCategory     `json:"scoreCategory,omitempty"`     //
	ScoreValue        *int                                                                                         `json:"scoreValue,omitempty"`        // Score Value
	ClientCount       *int                                                                                         `json:"clientCount,omitempty"`       // Client Count
	ClientUniqueCount *ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreListClientUniqueCount `json:"clientUniqueCount,omitempty"` // Client Unique Count
	Starttime         *int                                                                                         `json:"starttime,omitempty"`         // Starttime
	Endtime           *int                                                                                         `json:"endtime,omitempty"`           // Endtime
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreListScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Score Category
	Value         string `json:"value,omitempty"`         // Value
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreListClientUniqueCount interface{}
type ResponseClientsClientProximity struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}

//GetClientDetail Get Client Detail - e2ad-ba79-43ba-b3e9
/* Returns detailed Client information retrieved by Mac Address for any given point of time.


@param GetClientDetailQueryParams Filtering parameter
*/
func (s *ClientsService) GetClientDetail(GetClientDetailQueryParams *GetClientDetailQueryParams) (*ResponseClientsGetClientDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-detail"

	queryString, _ := query.Values(GetClientDetailQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetClientDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetClientDetail")
	}

	result := response.Result().(*ResponseClientsGetClientDetail)
	return result, response, err

}

//GetClientEnrichmentDetails Get Client Enrichment Details - b199-685d-4d08-9a67
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user, the devices that the user is connected to and the assurance issues that the user is impacted by


@param GetClientEnrichmentDetailsHeaderParams Custom header parameters
*/
func (s *ClientsService) GetClientEnrichmentDetails(GetClientEnrichmentDetailsHeaderParams *GetClientEnrichmentDetailsHeaderParams) (*ResponseClientsGetClientEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetClientEnrichmentDetailsHeaderParams != nil {

		if GetClientEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetClientEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetClientEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetClientEnrichmentDetailsHeaderParams.EntityValue)
		}

		if GetClientEnrichmentDetailsHeaderParams.IssueCategory != "" {
			clientRequest = clientRequest.SetHeader("issueCategory", GetClientEnrichmentDetailsHeaderParams.IssueCategory)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseClientsGetClientEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetClientEnrichmentDetails")
	}

	result := response.Result().(*ResponseClientsGetClientEnrichmentDetails)
	return result, response, err

}

//GetOverallClientHealth Get Overall Client Health - 149a-a93b-4ddb-80dd
/* Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time


@param GetOverallClientHealthQueryParams Filtering parameter
*/
func (s *ClientsService) GetOverallClientHealth(GetOverallClientHealthQueryParams *GetOverallClientHealthQueryParams) (*ResponseClientsGetOverallClientHealth, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-health"

	queryString, _ := query.Values(GetOverallClientHealthQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetOverallClientHealth{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOverallClientHealth")
	}

	result := response.Result().(*ResponseClientsGetOverallClientHealth)
	return result, response, err

}

//ClientProximity Client Proximity - 4497-ebe2-4c88-84a1
/* This intent API will provide client proximity information for a specific wireless user. Proximity is defined as presence on the same floor at the same time as the specified wireless user. The Proximity workflow requires the subscription to the following event (via the Event Notification workflow) prior to making this API call: NETWORK-CLIENTS-3-506 Client Proximity Report.


@param ClientProximityQueryParams Filtering parameter
*/
func (s *ClientsService) ClientProximity(ClientProximityQueryParams *ClientProximityQueryParams) (*ResponseClientsClientProximity, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-proximity"

	queryString, _ := query.Values(ClientProximityQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsClientProximity{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ClientProximity")
	}

	result := response.Result().(*ResponseClientsClientProximity)
	return result, response, err

}
