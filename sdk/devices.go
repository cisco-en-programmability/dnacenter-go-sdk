package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DevicesService service

type GetPlannedAccessPointsForBuildingQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //limit
	Offset float64 `url:"offset,omitempty"` //offset
	Radios bool    `url:"radios,omitempty"` //inlcude planned radio details
}
type GetDeviceDetailQueryParams struct {
	Timestamp  string `url:"timestamp,omitempty"`  //Epoch time(in milliseconds) when the device data is required
	SearchBy   string `url:"searchBy,omitempty"`   //MAC Address or Device Name value or UUID of the network device
	IDentifier string `url:"identifier,omitempty"` //One of keywords : macAddress or uuid or nwDeviceName
}
type GetDeviceEnrichmentDetailsHeaderParams struct {
	EntityType  string `url:"entity_type,omitempty"`  //Expects type string. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
	EntityValue string `url:"entity_value,omitempty"` //Expects type string. Contains the actual value for the entity type that has been defined
}
type DevicesQueryParams struct {
	DeviceRole string  `url:"deviceRole,omitempty"` //The device role (One of CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, AP)
	SiteID     string  `url:"siteId,omitempty"`     //Assurance site UUID value
	Health     string  `url:"health,omitempty"`     //The device overall health (One of POOR, FAIR, GOOD)
	StartTime  float64 `url:"startTime,omitempty"`  //UTC epoch time in milliseconds
	EndTime    float64 `url:"endTime,omitempty"`    //UTC epoch time in miliseconds
	Limit      float64 `url:"limit,omitempty"`      //Max number of device entries in the response (default to 50.  Max at 1000)
	Offset     float64 `url:"offset,omitempty"`     //The offset of the first device in the returned data
}
type GetPlannedAccessPointsForFloorQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //limit
	Offset float64 `url:"offset,omitempty"` //offset
	Radios bool    `url:"radios,omitempty"` //inlcude planned radio details
}
type GetAllInterfacesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //offset
	Limit  float64 `url:"limit,omitempty"`  //limit
}
type GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams struct {
	Name string `url:"name,omitempty"` //Interface name
}
type UpdateInterfaceDetailsQueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type ClearMacAddressTableQueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type GetDeviceListQueryParams struct {
	Hostname                  []string `url:"hostname,omitempty"`                   //hostname
	ManagementIPAddress       []string `url:"managementIpAddress,omitempty"`        //managementIpAddress
	MacAddress                []string `url:"macAddress,omitempty"`                 //macAddress
	LocationName              []string `url:"locationName,omitempty"`               //locationName
	SerialNumber              []string `url:"serialNumber,omitempty"`               //serialNumber
	Location                  []string `url:"location,omitempty"`                   //location
	Family                    []string `url:"family,omitempty"`                     //family
	Type                      []string `url:"type,omitempty"`                       //type
	Series                    []string `url:"series,omitempty"`                     //series
	CollectionStatus          []string `url:"collectionStatus,omitempty"`           //collectionStatus
	CollectionInterval        []string `url:"collectionInterval,omitempty"`         //collectionInterval
	NotSyncedForMinutes       []string `url:"notSyncedForMinutes,omitempty"`        //notSyncedForMinutes
	ErrorCode                 []string `url:"errorCode,omitempty"`                  //errorCode
	ErrorDescription          []string `url:"errorDescription,omitempty"`           //errorDescription
	SoftwareVersion           []string `url:"softwareVersion,omitempty"`            //softwareVersion
	SoftwareType              []string `url:"softwareType,omitempty"`               //softwareType
	PlatformID                []string `url:"platformId,omitempty"`                 //platformId
	Role                      []string `url:"role,omitempty"`                       //role
	ReachabilityStatus        []string `url:"reachabilityStatus,omitempty"`         //reachabilityStatus
	UpTime                    []string `url:"upTime,omitempty"`                     //upTime
	AssociatedWlcIP           []string `url:"associatedWlcIp,omitempty"`            //associatedWlcIp
	Licensename               []string `url:"license.name,omitempty"`               //licenseName
	Licensetype               []string `url:"license.type,omitempty"`               //licenseType
	Licensestatus             []string `url:"license.status,omitempty"`             //licenseStatus
	Modulename                []string `url:"module+name,omitempty"`                //moduleName
	Moduleequpimenttype       []string `url:"module+equpimenttype,omitempty"`       //moduleEqupimentType
	Moduleservicestate        []string `url:"module+servicestate,omitempty"`        //moduleServiceState
	Modulevendorequipmenttype []string `url:"module+vendorequipmenttype,omitempty"` //moduleVendorEquipmentType
	Modulepartnumber          []string `url:"module+partnumber,omitempty"`          //modulePartNumber
	Moduleoperationstatecode  []string `url:"module+operationstatecode,omitempty"`  //moduleOperationStateCode
	ID                        string   `url:"id,omitempty"`                         //Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
	DeviceSupportLevel        string   `url:"deviceSupportLevel,omitempty"`         //deviceSupportLevel
	Offset                    float64  `url:"offset,omitempty"`                     //offset >= 1 [X gives results from Xth device onwards]
	Limit                     float64  `url:"limit,omitempty"`                      //1 <= limit <= 500 [max. no. of devices to be returned in the result]
}
type GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams struct {
	VrfName                   string  `url:"vrfName,omitempty"`                   //vrfName
	ManagementIPAddress       string  `url:"managementIpAddress,omitempty"`       //managementIpAddress
	Hostname                  string  `url:"hostname,omitempty"`                  //hostname
	MacAddress                string  `url:"macAddress,omitempty"`                //macAddress
	Family                    string  `url:"family,omitempty"`                    //family
	CollectionStatus          string  `url:"collectionStatus,omitempty"`          //collectionStatus
	CollectionInterval        string  `url:"collectionInterval,omitempty"`        //collectionInterval
	SoftwareVersion           string  `url:"softwareVersion,omitempty"`           //softwareVersion
	SoftwareType              string  `url:"softwareType,omitempty"`              //softwareType
	ReachabilityStatus        string  `url:"reachabilityStatus,omitempty"`        //reachabilityStatus
	ReachabilityFailureReason string  `url:"reachabilityFailureReason,omitempty"` //reachabilityFailureReason
	ErrorCode                 string  `url:"errorCode,omitempty"`                 //errorCode
	PlatformID                string  `url:"platformId,omitempty"`                //platformId
	Series                    string  `url:"series,omitempty"`                    //series
	Type                      string  `url:"type,omitempty"`                      //type
	SerialNumber              string  `url:"serialNumber,omitempty"`              //serialNumber
	UpTime                    string  `url:"upTime,omitempty"`                    //upTime
	Role                      string  `url:"role,omitempty"`                      //role
	RoleSource                string  `url:"roleSource,omitempty"`                //roleSource
	AssociatedWlcIP           string  `url:"associatedWlcIp,omitempty"`           //associatedWlcIp
	Offset                    float64 `url:"offset,omitempty"`                    //offset
	Limit                     float64 `url:"limit,omitempty"`                     //limit
}
type GetFunctionalCapabilityForDevicesQueryParams struct {
	DeviceID     string   `url:"deviceId,omitempty"`     //Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
	FunctionName []string `url:"functionName,omitempty"` //functionName
}
type InventoryInsightDeviceLinkMismatchApIQueryParams struct {
	Offset   string `url:"offset,omitempty"`   //Row Number.  Default value is 1
	Limit    string `url:"limit,omitempty"`    //Default value is 500
	Category string `url:"category,omitempty"` //Links mismatch category.  Value can be speed-duplex or vlan.
	SortBy   string `url:"sortBy,omitempty"`   //Sort By
	Order    string `url:"order,omitempty"`    //Order.  Value can be asc or desc.  Default value is asc
}
type ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESQueryParams struct {
	Offset string `url:"offset,omitempty"` //Row Number.  Default value is 1
	Limit  string `url:"limit,omitempty"`  //Default value is 500
	SortBy string `url:"sortBy,omitempty"` //Sort By
	Order  string `url:"order,omitempty"`  //Order
}
type GetModulesQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	Limit                    string   `url:"limit,omitempty"`                    //limit
	Offset                   string   `url:"offset,omitempty"`                   //offset
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type GetModuleCountQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type SyncDevicesQueryParams struct {
	ForceSync bool `url:"forceSync,omitempty"` //forceSync
}
type RegisterDeviceForWsaQueryParams struct {
	SerialNumber string `url:"serialNumber,omitempty"` //Serial number of the device
	Macaddress   string `url:"macaddress,omitempty"`   //Mac addres of the device
}
type ReturnPowerSupplyFanDetailsForTheGivenDeviceQueryParams struct {
	Type string `url:"type,omitempty"` //Type value should be PowerSupply or Fan
}
type ReturnsPoeInterfaceDetailsForTheDeviceQueryParams struct {
	InterfaceNameList string `url:"interfaceNameList,omitempty"` //comma seperated interface names
}
type DeleteDeviceByIDQueryParams struct {
	CleanConfig bool `url:"cleanConfig,omitempty"` //cleanConfig
}
type GetDeviceInterfaceVLANsQueryParams struct {
	InterfaceType string `url:"interfaceType,omitempty"` //Vlan assocaited with sub-interface
}

type ResponseDevicesGetPlannedAccessPointsForBuilding struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForBuildingResponse `json:"response,omitempty"` //
	Version  *int                                                        `json:"version,omitempty"`  // Version
	Total    *int                                                        `json:"total,omitempty"`    // Total
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponse struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributes `json:"attributes,omitempty"` //
	Location   *ResponseDevicesGetPlannedAccessPointsForBuildingResponseLocation   `json:"location,omitempty"`   // Location
	Position   *ResponseDevicesGetPlannedAccessPointsForBuildingResponsePosition   `json:"position,omitempty"`   //
	RadioCount *int                                                                `json:"radioCount,omitempty"` // Radio Count
	Radios     *[]ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadios   `json:"radios,omitempty"`     //
	IsSensor   *bool                                                               `json:"isSensor,omitempty"`   // Is Sensor
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributes struct {
	ID            *int                                                                          `json:"id,omitempty"`            // Id
	InstanceUUID  string                                                                        `json:"instanceUuid,omitempty"`  // Instance Uuid
	Name          string                                                                        `json:"name,omitempty"`          // Name
	TypeString    string                                                                        `json:"typeString,omitempty"`    // Type String
	Domain        string                                                                        `json:"domain,omitempty"`        // Domain
	HeirarchyName string                                                                        `json:"heirarchyName,omitempty"` // Heirarchy Name
	Source        string                                                                        `json:"source,omitempty"`        // Source
	CreateDate    *int                                                                          `json:"createDate,omitempty"`    // Create Date
	Macaddress    *ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributesMacaddress `json:"macaddress,omitempty"`    // Macaddress
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributesMacaddress interface{}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseLocation interface{}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponsePosition struct {
	X *float64 `json:"x,omitempty"` // X
	Y *float64 `json:"y,omitempty"` // Y
	Z *float64 `json:"z,omitempty"` // Z
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributes `json:"attributes,omitempty"` //
	Antenna    *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAntenna    `json:"antenna,omitempty"`    //
	IsSensor   *bool                                                                     `json:"isSensor,omitempty"`   // Is Sensor
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributes struct {
	ID            *int                                                                                   `json:"id,omitempty"`            // Id
	InstanceUUID  string                                                                                 `json:"instanceUuid,omitempty"`  // Instance Uuid
	SlotID        *int                                                                                   `json:"slotId,omitempty"`        // Slot Id
	IfTypeString  string                                                                                 `json:"ifTypeString,omitempty"`  // If Type String
	IfTypeSubband string                                                                                 `json:"ifTypeSubband,omitempty"` // If Type Subband
	Channel       *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributesChannel       `json:"channel,omitempty"`       // Channel
	ChannelString *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributesChannelString `json:"channelString,omitempty"` // Channel String
	IfMode        string                                                                                 `json:"ifMode,omitempty"`        // If Mode
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributesChannel interface{}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributesChannelString interface{}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAntenna struct {
	Name           string   `json:"name,omitempty"`           // Name
	Type           string   `json:"type,omitempty"`           // Type
	Mode           string   `json:"mode,omitempty"`           // Mode
	AzimuthAngle   *float64 `json:"azimuthAngle,omitempty"`   // Azimuth Angle
	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation Angle
	Gain           *float64 `json:"gain,omitempty"`           // Gain
}
type ResponseDevicesGetDeviceDetail struct {
	Response *ResponseDevicesGetDeviceDetailResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetDeviceDetailResponse struct {
	HALastResetReason      string `json:"HALastResetReason,omitempty"`      // H A Last Reset Reason
	ManagementIPAddr       string `json:"managementIpAddr,omitempty"`       // Management Ip Addr
	HAPrimaryPowerStatus   string `json:"HAPrimaryPowerStatus,omitempty"`   // H A Primary Power Status
	RedundancyMode         string `json:"redundancyMode,omitempty"`         // Redundancy Mode
	CommunicationState     string `json:"communicationState,omitempty"`     // Communication State
	NwDeviceName           string `json:"nwDeviceName,omitempty"`           // Nw Device Name
	RedundancyUnit         string `json:"redundancyUnit,omitempty"`         // Redundancy Unit
	PlatformID             string `json:"platformId,omitempty"`             // Platform Id
	RedundancyPeerState    string `json:"redundancyPeerState,omitempty"`    // Redundancy Peer State
	NwDeviceID             string `json:"nwDeviceId,omitempty"`             // Nw Device Id
	RedundancyState        string `json:"redundancyState,omitempty"`        // Redundancy State
	NwDeviceRole           string `json:"nwDeviceRole,omitempty"`           // Nw Device Role
	NwDeviceFamily         string `json:"nwDeviceFamily,omitempty"`         // Nw Device Family
	MacAddress             string `json:"macAddress,omitempty"`             // Mac Address
	CollectionStatus       string `json:"collectionStatus,omitempty"`       // Collection Status
	DeviceSeries           string `json:"deviceSeries,omitempty"`           // Device Series
	OsType                 string `json:"osType,omitempty"`                 // Os Type
	ClientCount            string `json:"clientCount,omitempty"`            // Client Count
	HASecondaryPowerStatus string `json:"HASecondaryPowerStatus,omitempty"` // H A Secondary Power Status
	SoftwareVersion        string `json:"softwareVersion,omitempty"`        // Software Version
	NwDeviceType           string `json:"nwDeviceType,omitempty"`           // Nw Device Type
	OverallHealth          *int   `json:"overallHealth,omitempty"`          // Overall Health
	MemoryScore            *int   `json:"memoryScore,omitempty"`            // Memory Score
	CPUScore               *int   `json:"cpuScore,omitempty"`               // Cpu Score
	NoiseScore             *int   `json:"noiseScore,omitempty"`             // Noise Score
	UtilizationScore       *int   `json:"utilizationScore,omitempty"`       // Utilization Score
	AirQualityScore        *int   `json:"airQualityScore,omitempty"`        // Air Quality Score
	InterferenceScore      *int   `json:"interferenceScore,omitempty"`      // Interference Score
	WqeScore               *int   `json:"wqeScore,omitempty"`               // Wqe Score
	FreeMbufScore          *int   `json:"freeMbufScore,omitempty"`          // Free Mbuf Score
	PacketPoolScore        *int   `json:"packetPoolScore,omitempty"`        // Packet Pool Score
	FreeTimerScore         *int   `json:"freeTimerScore,omitempty"`         // Free Timer Score
	Memory                 string `json:"memory,omitempty"`                 // Memory
	CPU                    string `json:"cpu,omitempty"`                    // Cpu
	Noise                  string `json:"noise,omitempty"`                  // Noise
	Utilization            string `json:"utilization,omitempty"`            // Utilization
	AirQuality             string `json:"airQuality,omitempty"`             // Air Quality
	Interference           string `json:"interference,omitempty"`           // Interference
	Wqe                    string `json:"wqe,omitempty"`                    // Wqe
	FreeMbuf               string `json:"freeMbuf,omitempty"`               // Free Mbuf
	PacketPool             string `json:"packetPool,omitempty"`             // Packet Pool
	FreeTimer              string `json:"freeTimer,omitempty"`              // Free Timer
	Location               string `json:"location,omitempty"`               // Location
	Timestamp              string `json:"timestamp,omitempty"`              // Timestamp
}
type ResponseDevicesGetDeviceEnrichmentDetails []ResponseItemDevicesGetDeviceEnrichmentDetails // Array of ResponseDevicesGetDeviceEnrichmentDetails
type ResponseItemDevicesGetDeviceEnrichmentDetails struct {
	DeviceDetails *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails struct {
	Family                    string                                                                        `json:"family,omitempty"`                    // Family
	Type                      string                                                                        `json:"type,omitempty"`                      // Type
	Location                  *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation           `json:"location,omitempty"`                  // Location
	ErrorCode                 string                                                                        `json:"errorCode,omitempty"`                 // Error Code
	MacAddress                string                                                                        `json:"macAddress,omitempty"`                // Mac Address
	Role                      string                                                                        `json:"role,omitempty"`                      // Role
	ApManagerInterfaceIP      string                                                                        `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                                        `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              string                                                                        `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                                        `json:"collectionStatus,omitempty"`          // Collection Status
	InterfaceCount            string                                                                        `json:"interfaceCount,omitempty"`            // Interface Count
	LineCardCount             string                                                                        `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                string                                                                        `json:"lineCardId,omitempty"`                // Line Card Id
	ManagementIPAddress       string                                                                        `json:"managementIpAddress,omitempty"`       // Management Ip Address
	MemorySize                string                                                                        `json:"memorySize,omitempty"`                // Memory Size
	PlatformID                string                                                                        `json:"platformId,omitempty"`                // Platform Id
	ReachabilityFailureReason string                                                                        `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                                        `json:"reachabilityStatus,omitempty"`        // Reachability Status
	SNMPContact               string                                                                        `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                                        `json:"snmpLocation,omitempty"`              // Snmp Location
	TunnelUDPPort             *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort      `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	WaasDeviceMode            *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	Series                    string                                                                        `json:"series,omitempty"`                    // Series
	InventoryStatusDetail     string                                                                        `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	CollectionInterval        string                                                                        `json:"collectionInterval,omitempty"`        // Collection Interval
	SerialNumber              string                                                                        `json:"serialNumber,omitempty"`              // Serial Number
	SoftwareVersion           string                                                                        `json:"softwareVersion,omitempty"`           // Software Version
	RoleSource                string                                                                        `json:"roleSource,omitempty"`                // Role Source
	Hostname                  string                                                                        `json:"hostname,omitempty"`                  // Hostname
	UpTime                    string                                                                        `json:"upTime,omitempty"`                    // Up Time
	LastUpdateTime            *int                                                                          `json:"lastUpdateTime,omitempty"`            // Last Update Time
	ErrorDescription          string                                                                        `json:"errorDescription,omitempty"`          // Error Description
	LocationName              *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName       `json:"locationName,omitempty"`              // Location Name
	TagCount                  string                                                                        `json:"tagCount,omitempty"`                  // Tag Count
	LastUpdated               string                                                                        `json:"lastUpdated,omitempty"`               // Last Updated
	InstanceUUID              string                                                                        `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                                        `json:"id,omitempty"`                        // Id
	NeighborTopology          *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes struct {
	Role            string                                                                                          `json:"role,omitempty"`            // Role
	Name            string                                                                                          `json:"name,omitempty"`            // Name
	ID              string                                                                                          `json:"id,omitempty"`              // Id
	Description     string                                                                                          `json:"description,omitempty"`     // Description
	DeviceType      string                                                                                          `json:"deviceType,omitempty"`      // Device Type
	PlatformID      string                                                                                          `json:"platformId,omitempty"`      // Platform Id
	Family          string                                                                                          `json:"family,omitempty"`          // Family
	IP              string                                                                                          `json:"ip,omitempty"`              // Ip
	SoftwareVersion string                                                                                          `json:"softwareVersion,omitempty"` // Software Version
	UserID          *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID          `json:"userId,omitempty"`          // User Id
	NodeType        string                                                                                          `json:"nodeType,omitempty"`        // Node Type
	RadioFrequency  *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency  `json:"radioFrequency,omitempty"`  // Radio Frequency
	Clients         *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients         `json:"clients,omitempty"`         // Clients
	Count           *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount           `json:"count,omitempty"`           // Count
	HealthScore     *int                                                                                            `json:"healthScore,omitempty"`     // Health Score
	Level           *float64                                                                                        `json:"level,omitempty"`           // Level
	FabricGroup     *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup     `json:"fabricGroup,omitempty"`     // Fabric Group
	ConnectedDevice *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice `json:"connectedDevice,omitempty"` // Connected Device
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks struct {
	Source          string                                                                                          `json:"source,omitempty"`          // Source
	LinkStatus      string                                                                                          `json:"linkStatus,omitempty"`      // Link Status
	Label           *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel         `json:"label,omitempty"`           // Label
	Target          string                                                                                          `json:"target,omitempty"`          // Target
	ID              *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID              `json:"id,omitempty"`              // Id
	PortUtilization *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Port Utilization
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseDevicesDevices struct {
	Version    string                            `json:"version,omitempty"`    // Version
	TotalCount *int                              `json:"totalCount,omitempty"` // Total Count
	Response   *[]ResponseDevicesDevicesResponse `json:"response,omitempty"`   //
}
type ResponseDevicesDevicesResponse struct {
	Name                       string                                            `json:"name,omitempty"`                       // Name
	Model                      string                                            `json:"model,omitempty"`                      // Model
	OsVersion                  string                                            `json:"osVersion,omitempty"`                  // Os Version
	IPAddress                  string                                            `json:"ipAddress,omitempty"`                  // Ip Address
	OverallHealth              *int                                              `json:"overallHealth,omitempty"`              // Overall Health
	IssueCount                 *float64                                          `json:"issueCount,omitempty"`                 // Issue Count
	Location                   string                                            `json:"location,omitempty"`                   // Location
	DeviceFamily               string                                            `json:"deviceFamily,omitempty"`               // Device Family
	DeviceType                 string                                            `json:"deviceType,omitempty"`                 // Device Type
	MacAddress                 string                                            `json:"macAddress,omitempty"`                 // Mac Address
	InterfaceLinkErrHealth     *int                                              `json:"interfaceLinkErrHealth,omitempty"`     // Interface Link Err Health
	CPUUlitilization           *int                                              `json:"cpuUlitilization,omitempty"`           // Cpu Ulitilization
	CPUHealth                  *int                                              `json:"cpuHealth,omitempty"`                  // Cpu Health
	MemoryUtilizationHealth    *int                                              `json:"memoryUtilizationHealth,omitempty"`    // Memory Utilization Health
	MemoryUtilization          *int                                              `json:"memoryUtilization,omitempty"`          // Memory Utilization
	InterDeviceLinkAvailHealth *int                                              `json:"interDeviceLinkAvailHealth,omitempty"` // Inter Device Link Avail Health
	ReachabilityHealth         string                                            `json:"reachabilityHealth,omitempty"`         // Reachability Health
	ClientCount                *ResponseDevicesDevicesResponseClientCount        `json:"clientCount,omitempty"`                //
	InterferenceHealth         *ResponseDevicesDevicesResponseInterferenceHealth `json:"interferenceHealth,omitempty"`         //
	NoiseHealth                *ResponseDevicesDevicesResponseNoiseHealth        `json:"noiseHealth,omitempty"`                //
	AirQualityHealth           *ResponseDevicesDevicesResponseAirQualityHealth   `json:"airQualityHealth,omitempty"`           //
	UtilizationHealth          *ResponseDevicesDevicesResponseUtilizationHealth  `json:"utilizationHealth,omitempty"`          //
}
type ResponseDevicesDevicesResponseClientCount struct {
	Radio0 *float64 `json:"radio0,omitempty"` // Radio0
	Radio1 *int     `json:"radio1,omitempty"` // Radio1
	Ghz24  *float64 `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int     `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseInterferenceHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseNoiseHealth struct {
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseAirQualityHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesDevicesResponseUtilizationHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0
	Radio1 *int `json:"radio1,omitempty"` // Radio1
	Ghz24  *int `json:"Ghz24,omitempty"`  // Ghz24
	Ghz50  *int `json:"Ghz50,omitempty"`  // Ghz50
}
type ResponseDevicesGetPlannedAccessPointsForFloor struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForFloorResponse `json:"response,omitempty"` //
	Version  *int                                                     `json:"version,omitempty"`  // Version
	Total    *int                                                     `json:"total,omitempty"`    // Total
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponse struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes `json:"attributes,omitempty"` //
	Location   *ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation   `json:"location,omitempty"`   // Location
	Position   *ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition   `json:"position,omitempty"`   //
	RadioCount *int                                                             `json:"radioCount,omitempty"` // Radio Count
	Radios     *[]ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios   `json:"radios,omitempty"`     //
	IsSensor   *bool                                                            `json:"isSensor,omitempty"`   // Is Sensor
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes struct {
	ID            *int                                                                       `json:"id,omitempty"`            // Id
	InstanceUUID  string                                                                     `json:"instanceUuid,omitempty"`  // Instance Uuid
	Name          string                                                                     `json:"name,omitempty"`          // Name
	TypeString    string                                                                     `json:"typeString,omitempty"`    // Type String
	Domain        string                                                                     `json:"domain,omitempty"`        // Domain
	HeirarchyName string                                                                     `json:"heirarchyName,omitempty"` // Heirarchy Name
	Source        string                                                                     `json:"source,omitempty"`        // Source
	CreateDate    *int                                                                       `json:"createDate,omitempty"`    // Create Date
	Macaddress    *ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributesMacaddress `json:"macaddress,omitempty"`    // Macaddress
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributesMacaddress interface{}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation interface{}
type ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition struct {
	X *float64 `json:"x,omitempty"` // X
	Y *float64 `json:"y,omitempty"` // Y
	Z *float64 `json:"z,omitempty"` // Z
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes `json:"attributes,omitempty"` //
	Antenna    *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna    `json:"antenna,omitempty"`    //
	IsSensor   *bool                                                                  `json:"isSensor,omitempty"`   // Is Sensor
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes struct {
	ID            *int                                                                                `json:"id,omitempty"`            // Id
	InstanceUUID  string                                                                              `json:"instanceUuid,omitempty"`  // Instance Uuid
	SlotID        *int                                                                                `json:"slotId,omitempty"`        // Slot Id
	IfTypeString  string                                                                              `json:"ifTypeString,omitempty"`  // If Type String
	IfTypeSubband string                                                                              `json:"ifTypeSubband,omitempty"` // If Type Subband
	Channel       *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributesChannel       `json:"channel,omitempty"`       // Channel
	ChannelString *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributesChannelString `json:"channelString,omitempty"` // Channel String
	IfMode        string                                                                              `json:"ifMode,omitempty"`        // If Mode
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributesChannel interface{}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributesChannelString interface{}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna struct {
	Name           string   `json:"name,omitempty"`           // Name
	Type           string   `json:"type,omitempty"`           // Type
	Mode           string   `json:"mode,omitempty"`           // Mode
	AzimuthAngle   *float64 `json:"azimuthAngle,omitempty"`   // Azimuth Angle
	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation Angle
	Gain           *float64 `json:"gain,omitempty"`           // Gain
}
type ResponseDevicesGetAllInterfaces struct {
	Response *[]ResponseDevicesGetAllInterfacesResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseDevicesGetAllInterfacesResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetDeviceInterfaceCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceByIP struct {
	Response *[]ResponseDevicesGetInterfaceByIPResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceByIPResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetIsisInterfaces struct {
	Response *[]ResponseDevicesGetIsisInterfacesResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetIsisInterfacesResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetInterfaceInfoByID struct {
	Response *[]ResponseDevicesGetInterfaceInfoByIDResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceInfoByIDResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetDeviceInterfaceCount2 struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName struct {
	Response *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRange struct {
	Response *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetOspfInterfaces struct {
	Response *[]ResponseDevicesGetOspfInterfacesResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetOspfInterfacesResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesGetInterfaceByID struct {
	Response *ResponseDevicesGetInterfaceByIDResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesGetInterfaceByIDResponse struct {
	AdminStatus                 string `json:"adminStatus,omitempty"`                 //
	ClassName                   string `json:"className,omitempty"`                   //
	Description                 string `json:"description,omitempty"`                 //
	DeviceID                    string `json:"deviceId,omitempty"`                    //
	Duplex                      string `json:"duplex,omitempty"`                      //
	ID                          string `json:"id,omitempty"`                          //
	IfIndex                     string `json:"ifIndex,omitempty"`                     //
	InstanceTenantID            string `json:"instanceTenantId,omitempty"`            //
	InstanceUUID                string `json:"instanceUuid,omitempty"`                //
	InterfaceType               string `json:"interfaceType,omitempty"`               //
	IPv4Address                 string `json:"ipv4Address,omitempty"`                 //
	IPv4Mask                    string `json:"ipv4Mask,omitempty"`                    //
	IsisSupport                 string `json:"isisSupport,omitempty"`                 //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	MappedPhysicalInterfaceID   string `json:"mappedPhysicalInterfaceId,omitempty"`   //
	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` //
	MediaType                   string `json:"mediaType,omitempty"`                   //
	NativeVLANID                string `json:"nativeVlanId,omitempty"`                //
	OspfSupport                 string `json:"ospfSupport,omitempty"`                 //
	Pid                         string `json:"pid,omitempty"`                         //
	PortMode                    string `json:"portMode,omitempty"`                    //
	PortName                    string `json:"portName,omitempty"`                    //
	PortType                    string `json:"portType,omitempty"`                    //
	SerialNo                    string `json:"serialNo,omitempty"`                    //
	Series                      string `json:"series,omitempty"`                      //
	Speed                       string `json:"speed,omitempty"`                       //
	Status                      string `json:"status,omitempty"`                      //
	VLANID                      string `json:"vlanId,omitempty"`                      //
	VoiceVLAN                   string `json:"voiceVlan,omitempty"`                   //
}
type ResponseDevicesUpdateInterfaceDetails struct {
	Response *ResponseDevicesUpdateInterfaceDetailsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseDevicesUpdateInterfaceDetailsResponse struct {
	Type       string                                                   `json:"type,omitempty"`       // Type
	Properties *ResponseDevicesUpdateInterfaceDetailsResponseProperties `json:"properties,omitempty"` //
	Required   []string                                                 `json:"required,omitempty"`   // Required
}
type ResponseDevicesUpdateInterfaceDetailsResponseProperties struct {
	TaskID *ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID `json:"taskId,omitempty"` //
	URL    *ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL    `json:"url,omitempty"`    //
}
type ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL struct {
	Type string `json:"type,omitempty"` // Type
}

type ResponseDevicesLegitOperationsForInterface struct {
	Response *ResponseDevicesLegitOperationsForInterfaceResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseDevicesLegitOperationsForInterfaceResponse struct {
	InterfaceUUID string                                                          `json:"interfaceUuid,omitempty"` // Interface Uuid
	Properties    *[]ResponseDevicesLegitOperationsForInterfaceResponseProperties `json:"properties,omitempty"`    //
	Operations    *[]ResponseDevicesLegitOperationsForInterfaceResponseOperations `json:"operations,omitempty"`    //
}
type ResponseDevicesLegitOperationsForInterfaceResponseProperties struct {
	Name          string                                                                     `json:"name,omitempty"`          // Name
	Applicable    *bool                                                                      `json:"applicable,omitempty"`    // Applicable
	FailureReason *ResponseDevicesLegitOperationsForInterfaceResponsePropertiesFailureReason `json:"failureReason,omitempty"` // Failure Reason
}
type ResponseDevicesLegitOperationsForInterfaceResponsePropertiesFailureReason interface{}
type ResponseDevicesLegitOperationsForInterfaceResponseOperations struct {
	Name          string                                                                     `json:"name,omitempty"`          // Name
	Applicable    *bool                                                                      `json:"applicable,omitempty"`    // Applicable
	FailureReason *ResponseDevicesLegitOperationsForInterfaceResponseOperationsFailureReason `json:"failureReason,omitempty"` // Failure Reason
}
type ResponseDevicesLegitOperationsForInterfaceResponseOperationsFailureReason interface{}
type ResponseDevicesClearMacAddressTable struct {
	Response *ResponseDevicesClearMacAddressTableResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesClearMacAddressTableResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseDevicesGetDeviceList struct {
	Response *[]ResponseDevicesGetDeviceListResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetDeviceListResponse struct {
	ReachabilityFailureReason string                                                    `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                    `json:"reachabilityStatus,omitempty"`        // Reachability Status
	Series                    string                                                    `json:"series,omitempty"`                    // Series
	SNMPContact               string                                                    `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                    `json:"snmpLocation,omitempty"`              // Snmp Location
	TagCount                  string                                                    `json:"tagCount,omitempty"`                  // Tag Count
	TunnelUDPPort             *ResponseDevicesGetDeviceListResponseTunnelUDPPort        `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	UptimeSeconds             *int                                                      `json:"uptimeSeconds,omitempty"`             // Uptime Seconds
	WaasDeviceMode            *ResponseDevicesGetDeviceListResponseWaasDeviceMode       `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	SerialNumber              string                                                    `json:"serialNumber,omitempty"`              // Serial Number
	LastUpdateTime            *int                                                      `json:"lastUpdateTime,omitempty"`            // Last Update Time
	MacAddress                string                                                    `json:"macAddress,omitempty"`                // Mac Address
	UpTime                    string                                                    `json:"upTime,omitempty"`                    // Up Time
	DeviceSupportLevel        string                                                    `json:"deviceSupportLevel,omitempty"`        // Device Support Level
	Hostname                  string                                                    `json:"hostname,omitempty"`                  // Hostname
	Type                      string                                                    `json:"type,omitempty"`                      // Type
	MemorySize                string                                                    `json:"memorySize,omitempty"`                // Memory Size
	Family                    string                                                    `json:"family,omitempty"`                    // Family
	ErrorCode                 string                                                    `json:"errorCode,omitempty"`                 // Error Code
	SoftwareType              string                                                    `json:"softwareType,omitempty"`              // Software Type
	SoftwareVersion           string                                                    `json:"softwareVersion,omitempty"`           // Software Version
	Description               string                                                    `json:"description,omitempty"`               // Description
	RoleSource                string                                                    `json:"roleSource,omitempty"`                // Role Source
	Location                  *ResponseDevicesGetDeviceListResponseLocation             `json:"location,omitempty"`                  // Location
	Role                      string                                                    `json:"role,omitempty"`                      // Role
	CollectionInterval        string                                                    `json:"collectionInterval,omitempty"`        // Collection Interval
	InventoryStatusDetail     string                                                    `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	ApEthernetMacAddress      *ResponseDevicesGetDeviceListResponseApEthernetMacAddress `json:"apEthernetMacAddress,omitempty"`      // Ap Ethernet Mac Address
	ApManagerInterfaceIP      string                                                    `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                    `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              string                                                    `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                    `json:"collectionStatus,omitempty"`          // Collection Status
	ErrorDescription          string                                                    `json:"errorDescription,omitempty"`          // Error Description
	InterfaceCount            string                                                    `json:"interfaceCount,omitempty"`            // Interface Count
	LastUpdated               string                                                    `json:"lastUpdated,omitempty"`               // Last Updated
	LineCardCount             string                                                    `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                string                                                    `json:"lineCardId,omitempty"`                // Line Card Id
	LocationName              *ResponseDevicesGetDeviceListResponseLocationName         `json:"locationName,omitempty"`              // Location Name
	ManagedAtleastOnce        *bool                                                     `json:"managedAtleastOnce,omitempty"`        // Managed Atleast Once
	ManagementIPAddress       string                                                    `json:"managementIpAddress,omitempty"`       // Management Ip Address
	PlatformID                string                                                    `json:"platformId,omitempty"`                // Platform Id
	ManagementState           string                                                    `json:"managementState,omitempty"`           // Management State
	InstanceTenantID          string                                                    `json:"instanceTenantId,omitempty"`          // Instance Tenant Id
	InstanceUUID              string                                                    `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                    `json:"id,omitempty"`                        // Id
}
type ResponseDevicesGetDeviceListResponseTunnelUDPPort interface{}
type ResponseDevicesGetDeviceListResponseWaasDeviceMode interface{}
type ResponseDevicesGetDeviceListResponseLocation interface{}
type ResponseDevicesGetDeviceListResponseApEthernetMacAddress interface{}
type ResponseDevicesGetDeviceListResponseLocationName interface{}
type ResponseDevicesAddDevice2 struct {
	Response *ResponseDevicesAddDevice2Response `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}
type ResponseDevicesAddDevice2Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesSyncDevices2 struct {
	Response *ResponseDevicesSyncDevices2Response `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}
type ResponseDevicesSyncDevices2Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}
type ResponseDevicesUpdateDeviceRole struct {
	Response *ResponseDevicesUpdateDeviceRoleResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesUpdateDeviceRoleResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetPollingIntervalForAllDevices struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceConfigForAllDevices struct {
	Response *[]ResponseDevicesGetDeviceConfigForAllDevicesResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceConfigForAllDevicesResponse struct {
	AttributeInfo   *ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo `json:"attributeInfo,omitempty"`   //
	CdpNeighbors    string                                                            `json:"cdpNeighbors,omitempty"`    //
	HealthMonitor   string                                                            `json:"healthMonitor,omitempty"`   //
	ID              string                                                            `json:"id,omitempty"`              //
	IntfDescription string                                                            `json:"intfDescription,omitempty"` //
	Inventory       string                                                            `json:"inventory,omitempty"`       //
	IPIntfBrief     string                                                            `json:"ipIntfBrief,omitempty"`     //
	MacAddressTable string                                                            `json:"macAddressTable,omitempty"` //
	RunningConfig   string                                                            `json:"runningConfig,omitempty"`   //
	SNMP            string                                                            `json:"snmp,omitempty"`            //
	Version         string                                                            `json:"version,omitempty"`         //
}
type ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceConfigCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceCount2 struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesExportDeviceList struct {
	Response *ResponseDevicesExportDeviceListResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesExportDeviceListResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetFunctionalCapabilityForDevices struct {
	Response *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponse struct {
	AttributeInfo        *ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo          `json:"attributeInfo,omitempty"`        //
	DeviceID             string                                                                          `json:"deviceId,omitempty"`             //
	FunctionalCapability *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability `json:"functionalCapability,omitempty"` //
	ID                   string                                                                          `json:"id,omitempty"`                   //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability struct {
	AttributeInfo   *ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo     `json:"attributeInfo,omitempty"`   //
	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails `json:"functionDetails,omitempty"` //
	FunctionName    string                                                                                         `json:"functionName,omitempty"`    //
	FunctionOpState string                                                                                         `json:"functionOpState,omitempty"` //
	ID              string                                                                                         `json:"id,omitempty"`              //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` //
	ID            string                                                                                                    `json:"id,omitempty"`            //
	PropertyName  string                                                                                                    `json:"propertyName,omitempty"`  //
	StringValue   string                                                                                                    `json:"stringValue,omitempty"`   //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByID struct {
	Response *ResponseDevicesGetFunctionalCapabilityByIDResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDevicesGetFunctionalCapabilityByIDResponse struct {
	AttributeInfo   *ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo     `json:"attributeInfo,omitempty"`   //
	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails `json:"functionDetails,omitempty"` //
	FunctionName    string                                                               `json:"functionName,omitempty"`    //
	FunctionOpState string                                                               `json:"functionOpState,omitempty"` //
	ID              string                                                               `json:"id,omitempty"`              //
}
type ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` //
	ID            string                                                                          `json:"id,omitempty"`            //
	PropertyName  string                                                                          `json:"propertyName,omitempty"`  //
	StringValue   string                                                                          `json:"stringValue,omitempty"`   //
}
type ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo interface{}
type ResponseDevicesInventoryInsightDeviceLinkMismatchApI struct {
	Response *[]ResponseDevicesInventoryInsightDeviceLinkMismatchAPIResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Api version
}
type ResponseDevicesInventoryInsightDeviceLinkMismatchAPIResponse struct {
	EndPortAllowedVLANIDs   string   `json:"endPortAllowedVlanIds,omitempty"`   // End port allowed vlan ids
	EndPortNativeVLANID     string   `json:"endPortNativeVlanId,omitempty"`     // End port native vlan id
	StartPortAllowedVLANIDs string   `json:"startPortAllowedVlanIds,omitempty"` // Start port allowed vlan ids
	StartPortNativeVLANID   string   `json:"startPortNativeVlanId,omitempty"`   // Start port native vlan id
	LinkStatus              string   `json:"linkStatus,omitempty"`              // Link status
	EndDeviceHostName       string   `json:"endDeviceHostName,omitempty"`       // End device hostname
	EndDeviceID             string   `json:"endDeviceId,omitempty"`             // End device id
	EndDeviceIPAddress      string   `json:"endDeviceIpAddress,omitempty"`      // End device ip address
	EndPortAddress          string   `json:"endPortAddress,omitempty"`          // End port address
	EndPortDuplex           string   `json:"endPortDuplex,omitempty"`           // End port duplex
	EndPortID               string   `json:"endPortId,omitempty"`               // End port id
	EndPortMask             string   `json:"endPortMask,omitempty"`             // End port mask
	EndPortName             string   `json:"endPortName,omitempty"`             // End port name
	EndPortPepID            string   `json:"endPortPepId,omitempty"`            // End port pep id
	EndPortSpeed            string   `json:"endPortSpeed,omitempty"`            // End port speed
	StartDeviceHostName     string   `json:"startDeviceHostName,omitempty"`     // Start device hostname
	StartDeviceID           string   `json:"startDeviceId,omitempty"`           // Start device id
	StartDeviceIPAddress    string   `json:"startDeviceIpAddress,omitempty"`    // Start device ip address
	StartPortAddress        string   `json:"startPortAddress,omitempty"`        // Start port address
	StartPortDuplex         string   `json:"startPortDuplex,omitempty"`         // Start port duplex
	StartPortID             string   `json:"startPortId,omitempty"`             // Start port id
	StartPortMask           string   `json:"startPortMask,omitempty"`           // Start port mask
	StartPortName           string   `json:"startPortName,omitempty"`           // Start port name
	StartPortPepID          string   `json:"startPortPepId,omitempty"`          // Start port pep id
	StartPortSpeed          string   `json:"startPortSpeed,omitempty"`          // Start port speed
	LastUpdated             string   `json:"lastUpdated,omitempty"`             // Last updated
	NumUpdates              *float64 `json:"numUpdates,omitempty"`              // Number updates
	AvgUpdateFrequency      *float64 `json:"avgUpdateFrequency,omitempty"`      // Average update frequency
	Type                    string   `json:"type,omitempty"`                    // Type
	InstanceUUID            string   `json:"instanceUuid,omitempty"`            // Unique instance id
	InstanceTenantID        string   `json:"instanceTenantId,omitempty"`        // Instance tenant id
}
type ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES struct {
	Response *[]ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESResponse `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  // Version
}
type ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESResponse struct {
	ID                  string `json:"id,omitempty"`                  // Id
	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Management Ip Address
	Hostname            string `json:"hostname,omitempty"`            // Hostname
	Type                string `json:"type,omitempty"`                // Type
	Family              string `json:"family,omitempty"`              // Family
	LastUpdated         string `json:"lastUpdated,omitempty"`         // Last Updated
	UpTime              string `json:"upTime,omitempty"`              // Up Time
	ReachabilityStatus  string `json:"reachabilityStatus,omitempty"`  // Reachability Status
}
type ResponseDevicesGetNetworkDeviceByIP struct {
	Response *ResponseDevicesGetNetworkDeviceByIPResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}
type ResponseDevicesGetNetworkDeviceByIPResponse struct {
	ApManagerInterfaceIP      string `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIP           string `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string `json:"bootDateTime,omitempty"`              //
	CollectionInterval        string `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string `json:"errorCode,omitempty"`                 //
	ErrorDescription          string `json:"errorDescription,omitempty"`          //
	Family                    string `json:"family,omitempty"`                    //
	Hostname                  string `json:"hostname,omitempty"`                  //
	ID                        string `json:"id,omitempty"`                        //
	InstanceTenantID          string `json:"instanceTenantId,omitempty"`          //
	InstanceUUID              string `json:"instanceUuid,omitempty"`              //
	InterfaceCount            string `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            *int   `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string `json:"lastUpdated,omitempty"`               //
	LineCardCount             string `json:"lineCardCount,omitempty"`             //
	LineCardID                string `json:"lineCardId,omitempty"`                //
	Location                  string `json:"location,omitempty"`                  //
	LocationName              string `json:"locationName,omitempty"`              //
	MacAddress                string `json:"macAddress,omitempty"`                //
	ManagementIPAddress       string `json:"managementIpAddress,omitempty"`       //
	MemorySize                string `json:"memorySize,omitempty"`                //
	PlatformID                string `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        //
	Role                      string `json:"role,omitempty"`                      //
	RoleSource                string `json:"roleSource,omitempty"`                //
	SerialNumber              string `json:"serialNumber,omitempty"`              //
	Series                    string `json:"series,omitempty"`                    //
	SNMPContact               string `json:"snmpContact,omitempty"`               //
	SNMPLocation              string `json:"snmpLocation,omitempty"`              //
	SoftwareType              string `json:"softwareType,omitempty"`              //
	SoftwareVersion           string `json:"softwareVersion,omitempty"`           //
	TagCount                  string `json:"tagCount,omitempty"`                  //
	TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             //
	Type                      string `json:"type,omitempty"`                      //
	UpTime                    string `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string `json:"waasDeviceMode,omitempty"`            //
}
type ResponseDevicesGetModules struct {
	Response *[]ResponseDevicesGetModulesResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}
type ResponseDevicesGetModulesResponse struct {
	AssemblyNumber           string                                          `json:"assemblyNumber,omitempty"`           //
	AssemblyRevision         string                                          `json:"assemblyRevision,omitempty"`         //
	AttributeInfo            *ResponseDevicesGetModulesResponseAttributeInfo `json:"attributeInfo,omitempty"`            //
	ContainmentEntity        string                                          `json:"containmentEntity,omitempty"`        //
	Description              string                                          `json:"description,omitempty"`              //
	EntityPhysicalIndex      string                                          `json:"entityPhysicalIndex,omitempty"`      //
	ID                       string                                          `json:"id,omitempty"`                       //
	IsFieldReplaceable       string                                          `json:"isFieldReplaceable,omitempty"`       //
	IsReportingAlarmsAllowed string                                          `json:"isReportingAlarmsAllowed,omitempty"` //
	Manufacturer             string                                          `json:"manufacturer,omitempty"`             //
	ModuleIndex              *int                                            `json:"moduleIndex,omitempty"`              //
	Name                     string                                          `json:"name,omitempty"`                     //
	OperationalStateCode     string                                          `json:"operationalStateCode,omitempty"`     //
	PartNumber               string                                          `json:"partNumber,omitempty"`               //
	SerialNumber             string                                          `json:"serialNumber,omitempty"`             //
	VendorEquipmentType      string                                          `json:"vendorEquipmentType,omitempty"`      //
}
type ResponseDevicesGetModulesResponseAttributeInfo interface{}
type ResponseDevicesGetModuleCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetModuleInfoByID struct {
	Response *ResponseDevicesGetModuleInfoByIDResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  //
}
type ResponseDevicesGetModuleInfoByIDResponse struct {
	AssemblyNumber           string                                                 `json:"assemblyNumber,omitempty"`           //
	AssemblyRevision         string                                                 `json:"assemblyRevision,omitempty"`         //
	AttributeInfo            *ResponseDevicesGetModuleInfoByIDResponseAttributeInfo `json:"attributeInfo,omitempty"`            //
	ContainmentEntity        string                                                 `json:"containmentEntity,omitempty"`        //
	Description              string                                                 `json:"description,omitempty"`              //
	EntityPhysicalIndex      string                                                 `json:"entityPhysicalIndex,omitempty"`      //
	ID                       string                                                 `json:"id,omitempty"`                       //
	IsFieldReplaceable       string                                                 `json:"isFieldReplaceable,omitempty"`       //
	IsReportingAlarmsAllowed string                                                 `json:"isReportingAlarmsAllowed,omitempty"` //
	Manufacturer             string                                                 `json:"manufacturer,omitempty"`             //
	ModuleIndex              *int                                                   `json:"moduleIndex,omitempty"`              //
	Name                     string                                                 `json:"name,omitempty"`                     //
	OperationalStateCode     string                                                 `json:"operationalStateCode,omitempty"`     //
	PartNumber               string                                                 `json:"partNumber,omitempty"`               //
	SerialNumber             string                                                 `json:"serialNumber,omitempty"`             //
	VendorEquipmentType      string                                                 `json:"vendorEquipmentType,omitempty"`      //
}
type ResponseDevicesGetModuleInfoByIDResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceBySerialNumber struct {
	Response *ResponseDevicesGetDeviceBySerialNumberResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceBySerialNumberResponse struct {
	ApManagerInterfaceIP      string `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIP           string `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string `json:"bootDateTime,omitempty"`              //
	CollectionInterval        string `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string `json:"errorCode,omitempty"`                 //
	ErrorDescription          string `json:"errorDescription,omitempty"`          //
	Family                    string `json:"family,omitempty"`                    //
	Hostname                  string `json:"hostname,omitempty"`                  //
	ID                        string `json:"id,omitempty"`                        //
	InstanceTenantID          string `json:"instanceTenantId,omitempty"`          //
	InstanceUUID              string `json:"instanceUuid,omitempty"`              //
	InterfaceCount            string `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            *int   `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string `json:"lastUpdated,omitempty"`               //
	LineCardCount             string `json:"lineCardCount,omitempty"`             //
	LineCardID                string `json:"lineCardId,omitempty"`                //
	Location                  string `json:"location,omitempty"`                  //
	LocationName              string `json:"locationName,omitempty"`              //
	MacAddress                string `json:"macAddress,omitempty"`                //
	ManagementIPAddress       string `json:"managementIpAddress,omitempty"`       //
	MemorySize                string `json:"memorySize,omitempty"`                //
	PlatformID                string `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        //
	Role                      string `json:"role,omitempty"`                      //
	RoleSource                string `json:"roleSource,omitempty"`                //
	SerialNumber              string `json:"serialNumber,omitempty"`              //
	Series                    string `json:"series,omitempty"`                    //
	SNMPContact               string `json:"snmpContact,omitempty"`               //
	SNMPLocation              string `json:"snmpLocation,omitempty"`              //
	SoftwareType              string `json:"softwareType,omitempty"`              //
	SoftwareVersion           string `json:"softwareVersion,omitempty"`           //
	TagCount                  string `json:"tagCount,omitempty"`                  //
	TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             //
	Type                      string `json:"type,omitempty"`                      //
	UpTime                    string `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string `json:"waasDeviceMode,omitempty"`            //
}
type ResponseDevicesSyncDevices struct {
	Response *ResponseDevicesSyncDevicesResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  //
}
type ResponseDevicesSyncDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesRegisterDeviceForWsa struct {
	Response *ResponseDevicesRegisterDeviceForWsaResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}
type ResponseDevicesRegisterDeviceForWsaResponse struct {
	MacAddress   string `json:"macAddress,omitempty"`   //
	ModelNumber  string `json:"modelNumber,omitempty"`  //
	Name         string `json:"name,omitempty"`         //
	SerialNumber string `json:"serialNumber,omitempty"` //
	TenantID     string `json:"tenantId,omitempty"`     //
}
type ResponseDevicesGetChassisDetailsForDevice struct {
	Response *[]ResponseDevicesGetChassisDetailsForDeviceResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDevicesGetChassisDetailsForDeviceResponse struct {
	AssemblyNumber           string `json:"assemblyNumber,omitempty"`           // Assembly Number of the chassis
	AssemblyRevision         string `json:"assemblyRevision,omitempty"`         // Assembly Revision of the chassis
	ContainmentEntity        string `json:"containmentEntity,omitempty"`        // Containment Entity of the chassis
	Description              string `json:"description,omitempty"`              // Description of the chassis
	EntityPhysicalIndex      string `json:"entityPhysicalIndex,omitempty"`      // Entity Physical Index of the chassis
	HardwareVersion          string `json:"hardwareVersion,omitempty"`          // Hardware Version of the chassis
	InstanceUUID             string `json:"instanceUuid,omitempty"`             // ID of the chassis
	IsFieldReplaceable       string `json:"isFieldReplaceable,omitempty"`       // To mention if field is replaceable
	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed
	Manufacturer             string `json:"manufacturer,omitempty"`             // Manufacturer of the chassis
	Name                     string `json:"name,omitempty"`                     // Name of the chassis
	PartNumber               string `json:"partNumber,omitempty"`               // Part Number of the chassis
	SerialNumber             string `json:"serialNumber,omitempty"`             // Serial Number of the chassis
	VendorEquipmentType      string `json:"vendorEquipmentType,omitempty"`      // Vendor Equipment Type of the chassis
}
type ResponseDevicesGetStackDetailsForDevice struct {
	Response *ResponseDevicesGetStackDetailsForDeviceResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseDevicesGetStackDetailsForDeviceResponse struct {
	DeviceID        string                                                            `json:"deviceId,omitempty"`        // Device ID
	StackPortInfo   *[]ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo   `json:"stackPortInfo,omitempty"`   //
	StackSwitchInfo *[]ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo `json:"stackSwitchInfo,omitempty"` //
	SvlSwitchInfo   *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo   `json:"svlSwitchInfo,omitempty"`   //
}
type ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo struct {
	IsSynchOk               string `json:"isSynchOk,omitempty"`               // If link partner sends valid protocol message
	LinkActive              *bool  `json:"linkActive,omitempty"`              // If stack port is in same state as link partner
	LinkOk                  *bool  `json:"linkOk,omitempty"`                  // If link is stable
	Name                    string `json:"name,omitempty"`                    // Name of the stack port
	NeighborPort            string `json:"neighborPort,omitempty"`            // Neighbor's member number and stack port number
	NrLinkOkChanges         *int   `json:"nrLinkOkChanges,omitempty"`         // Relative stability of the link
	StackCableLengthInfo    string `json:"stackCableLengthInfo,omitempty"`    // Cable length
	StackPortOperStatusInfo string `json:"stackPortOperStatusInfo,omitempty"` // Port opearation status
	SwitchPort              string `json:"switchPort,omitempty"`              // Member number and stack port number
}
type ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo struct {
	EntPhysicalIndex  string `json:"entPhysicalIndex,omitempty"`  //
	HwPriority        *int   `json:"hwPriority,omitempty"`        // Hardware priority of the switch
	MacAddress        string `json:"macAddress,omitempty"`        // Mac address of the switch
	NumNextReload     *int   `json:"numNextReload,omitempty"`     // Stack member number to be used in next reload
	PlatformID        string `json:"platformId,omitempty"`        // Platform Id
	Role              string `json:"role,omitempty"`              // Function of the switch
	SerialNumber      string `json:"serialNumber,omitempty"`      // Serial number
	SoftwareImage     string `json:"softwareImage,omitempty"`     // Software image type running on the switch
	StackMemberNumber *int   `json:"stackMemberNumber,omitempty"` // Switch member number
	State             string `json:"state,omitempty"`             // Current state of the switch
	SwitchPriority    *int   `json:"switchPriority,omitempty"`    // Priority of the switch
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo struct {
	DadProtocol              string                                                                       `json:"dadProtocol,omitempty"`              // Stackwise virtual dual active detection config
	DadRecoveryReloadEnabled *bool                                                                        `json:"dadRecoveryReloadEnabled,omitempty"` // If dad recovery reload enabled
	DomainNumber             *int                                                                         `json:"domainNumber,omitempty"`             // Stackwise virtual switch domain number
	InDadRecoveryMode        *bool                                                                        `json:"inDadRecoveryMode,omitempty"`        // If in dad recovery mode
	SwVirtualStatus          string                                                                       `json:"swVirtualStatus,omitempty"`          // Stackwise virtual status
	SwitchMembers            *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers `json:"switchMembers,omitempty"`            //
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers struct {
	Bandwidth            string                                                                                           `json:"bandwidth,omitempty"`            // Bandwidth
	SvlMemberEndPoints   *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints   `json:"svlMemberEndPoints,omitempty"`   //
	SvlMemberNumber      *int                                                                                             `json:"svlMemberNumber,omitempty"`      // Switch member number
	SvlMemberPepSettings *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings `json:"svlMemberPepSettings,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints struct {
	SvlMemberEndPointPorts *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts `json:"svlMemberEndPointPorts,omitempty"` //
	SvlNumber              *int                                                                                                                 `json:"svlNumber,omitempty"`              // Stackwise virtual link number
	SvlStatus              string                                                                                                               `json:"svlStatus,omitempty"`              // Stackwise virtual status
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts struct {
	SvlProtocolStatus string `json:"svlProtocolStatus,omitempty"` // Stackwise virtual protocol status
	SwLocalInterface  string `json:"swLocalInterface,omitempty"`  // Stackwise virtual local interface
	SwRemoteInterface string `json:"swRemoteInterface,omitempty"` // Stackwise virtual remote interface
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings struct {
	DadEnabled       *bool  `json:"dadEnabled,omitempty"`       // If dadInterface is configured for dual active detection
	DadInterfaceName string `json:"dadInterfaceName,omitempty"` // Interface for dual active detection
}
type ResponseDevicesReturnPowerSupplyFanDetailsForTheGivenDevice struct {
	Response *[]ResponseDevicesReturnPowerSupplyFanDetailsForTheGivenDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseDevicesReturnPowerSupplyFanDetailsForTheGivenDeviceResponse struct {
	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational State Code
	ProductID            string `json:"productId,omitempty"`            // Product Id
	SerialNumber         string `json:"serialNumber,omitempty"`         // Serial Number
	VendorEquipmentType  string `json:"vendorEquipmentType,omitempty"`  // Vendor Equipment Type
	Description          string `json:"description,omitempty"`          // Description
	InstanceUUID         string `json:"instanceUuid,omitempty"`         // Instance Uuid
	Name                 string `json:"name,omitempty"`                 // Name
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice struct {
	Version  string                                                           `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse `json:"response,omitempty"` //
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse struct {
	AdminStatus    string `json:"adminStatus,omitempty"`    // Admin Status
	OperStatus     string `json:"operStatus,omitempty"`     // Oper Status
	InterfaceName  string `json:"interfaceName,omitempty"`  // Interface Name
	MaxPortPower   string `json:"maxPortPower,omitempty"`   // Max Port Power
	AllocatedPower string `json:"allocatedPower,omitempty"` // Allocated Power
	PortPowerDrawn string `json:"portPowerDrawn,omitempty"` // Port Power Drawn
}
type ResponseDevicesGetConnectedDeviceDetail struct {
	Response *ResponseDevicesGetConnectedDeviceDetailResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetConnectedDeviceDetailResponse struct {
	NeighborDevice string   `json:"neighborDevice,omitempty"` // Neighbor Device
	NeighborPort   string   `json:"neighborPort,omitempty"`   // Neighbor Port
	Capabilities   []string `json:"capabilities,omitempty"`   // Capabilities
}
type ResponseDevicesGetLinecardDetails struct {
	Response *[]ResponseDevicesGetLinecardDetailsResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetLinecardDetailsResponse struct {
	Serialno string `json:"serialno,omitempty"` // Serialno
	Partno   string `json:"partno,omitempty"`   // Partno
	Switchno string `json:"switchno,omitempty"` // Switchno
	Slotno   string `json:"slotno,omitempty"`   // Slotno
}
type ResponseDevicesPoeDetails struct {
	Response *ResponseDevicesPoeDetailsResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  // Version
}
type ResponseDevicesPoeDetailsResponse struct {
	PowerAllocated string `json:"powerAllocated,omitempty"` // Power Allocated
	PowerConsumed  string `json:"powerConsumed,omitempty"`  // Power Consumed
	PowerRemaining string `json:"powerRemaining,omitempty"` // Power Remaining
}
type ResponseDevicesGetSupervisorCardDetail struct {
	Response *[]ResponseDevicesGetSupervisorCardDetailResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetSupervisorCardDetailResponse struct {
	Serialno string `json:"serialno,omitempty"` // Serialno
	Partno   string `json:"partno,omitempty"`   // Partno
	Switchno string `json:"switchno,omitempty"` // Switchno
	Slotno   string `json:"slotno,omitempty"`   // Slotno
}
type ResponseDevicesGetDeviceByID struct {
	Response *ResponseDevicesGetDeviceByIDResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceByIDResponse struct {
	ApManagerInterfaceIP      string `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIP           string `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string `json:"bootDateTime,omitempty"`              //
	CollectionInterval        string `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string `json:"errorCode,omitempty"`                 //
	ErrorDescription          string `json:"errorDescription,omitempty"`          //
	Family                    string `json:"family,omitempty"`                    //
	Hostname                  string `json:"hostname,omitempty"`                  //
	ID                        string `json:"id,omitempty"`                        //
	InstanceTenantID          string `json:"instanceTenantId,omitempty"`          //
	InstanceUUID              string `json:"instanceUuid,omitempty"`              //
	InterfaceCount            string `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            *int   `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string `json:"lastUpdated,omitempty"`               //
	LineCardCount             string `json:"lineCardCount,omitempty"`             //
	LineCardID                string `json:"lineCardId,omitempty"`                //
	Location                  string `json:"location,omitempty"`                  //
	LocationName              string `json:"locationName,omitempty"`              //
	MacAddress                string `json:"macAddress,omitempty"`                //
	ManagementIPAddress       string `json:"managementIpAddress,omitempty"`       //
	MemorySize                string `json:"memorySize,omitempty"`                //
	PlatformID                string `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        //
	Role                      string `json:"role,omitempty"`                      //
	RoleSource                string `json:"roleSource,omitempty"`                //
	SerialNumber              string `json:"serialNumber,omitempty"`              //
	Series                    string `json:"series,omitempty"`                    //
	SNMPContact               string `json:"snmpContact,omitempty"`               //
	SNMPLocation              string `json:"snmpLocation,omitempty"`              //
	SoftwareType              string `json:"softwareType,omitempty"`              //
	SoftwareVersion           string `json:"softwareVersion,omitempty"`           //
	TagCount                  string `json:"tagCount,omitempty"`                  //
	TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             //
	Type                      string `json:"type,omitempty"`                      //
	UpTime                    string `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string `json:"waasDeviceMode,omitempty"`            //
}
type ResponseDevicesDeleteDeviceByID struct {
	Response *ResponseDevicesDeleteDeviceByIDResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesDeleteDeviceByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDevicesGetDeviceSummary struct {
	Response *ResponseDevicesGetDeviceSummaryResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceSummaryResponse struct {
	ID         string `json:"id,omitempty"`         //
	Role       string `json:"role,omitempty"`       //
	RoleSource string `json:"roleSource,omitempty"` //
}
type ResponseDevicesGetPollingIntervalByID struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetOrganizationListForMeraki struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceInterfaceVLANs struct {
	Response *[]ResponseDevicesGetDeviceInterfaceVLANsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDevicesGetDeviceInterfaceVLANsResponse struct {
	InterfaceName  string `json:"interfaceName,omitempty"`  //
	IPAddress      string `json:"ipAddress,omitempty"`      //
	Mask           *int   `json:"mask,omitempty"`           //
	NetworkAddress string `json:"networkAddress,omitempty"` //
	NumberOfIPs    *int   `json:"numberOfIPs,omitempty"`    //
	Prefix         string `json:"prefix,omitempty"`         //
	VLANNumber     *int   `json:"vlanNumber,omitempty"`     //
	VLANType       string `json:"vlanType,omitempty"`       //
}
type ResponseDevicesGetWirelessLanControllerDetailsByID struct {
	Response *ResponseDevicesGetWirelessLanControllerDetailsByIDResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetWirelessLanControllerDetailsByIDResponse struct {
	AdminEnabledPorts        *[]int `json:"adminEnabledPorts,omitempty"`        //
	ApGroupName              string `json:"apGroupName,omitempty"`              //
	DeviceID                 string `json:"deviceId,omitempty"`                 //
	EthMacAddress            string `json:"ethMacAddress,omitempty"`            //
	FlexGroupName            string `json:"flexGroupName,omitempty"`            //
	ID                       string `json:"id,omitempty"`                       //
	InstanceTenantID         string `json:"instanceTenantId,omitempty"`         //
	InstanceUUID             string `json:"instanceUuid,omitempty"`             //
	LagModeEnabled           *bool  `json:"lagModeEnabled,omitempty"`           //
	NetconfEnabled           *bool  `json:"netconfEnabled,omitempty"`           //
	WirelessLicenseInfo      string `json:"wirelessLicenseInfo,omitempty"`      //
	WirelessPackageInstalled *bool  `json:"wirelessPackageInstalled,omitempty"` //
}
type ResponseDevicesGetDeviceConfigByID struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDevicesGetNetworkDeviceByPaginationRange struct {
	Response *[]ResponseDevicesGetNetworkDeviceByPaginationRangeResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseDevicesGetNetworkDeviceByPaginationRangeResponse struct {
	ApManagerInterfaceIP      string `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIP           string `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string `json:"bootDateTime,omitempty"`              //
	CollectionInterval        string `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string `json:"errorCode,omitempty"`                 //
	ErrorDescription          string `json:"errorDescription,omitempty"`          //
	Family                    string `json:"family,omitempty"`                    //
	Hostname                  string `json:"hostname,omitempty"`                  //
	ID                        string `json:"id,omitempty"`                        //
	InstanceTenantID          string `json:"instanceTenantId,omitempty"`          //
	InstanceUUID              string `json:"instanceUuid,omitempty"`              //
	InterfaceCount            string `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            *int   `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string `json:"lastUpdated,omitempty"`               //
	LineCardCount             string `json:"lineCardCount,omitempty"`             //
	LineCardID                string `json:"lineCardId,omitempty"`                //
	Location                  string `json:"location,omitempty"`                  //
	LocationName              string `json:"locationName,omitempty"`              //
	MacAddress                string `json:"macAddress,omitempty"`                //
	ManagementIPAddress       string `json:"managementIpAddress,omitempty"`       //
	MemorySize                string `json:"memorySize,omitempty"`                //
	PlatformID                string `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        //
	Role                      string `json:"role,omitempty"`                      //
	RoleSource                string `json:"roleSource,omitempty"`                //
	SerialNumber              string `json:"serialNumber,omitempty"`              //
	Series                    string `json:"series,omitempty"`                    //
	SNMPContact               string `json:"snmpContact,omitempty"`               //
	SNMPLocation              string `json:"snmpLocation,omitempty"`              //
	SoftwareType              string `json:"softwareType,omitempty"`              //
	SoftwareVersion           string `json:"softwareVersion,omitempty"`           //
	TagCount                  string `json:"tagCount,omitempty"`                  //
	TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             //
	Type                      string `json:"type,omitempty"`                      //
	UpTime                    string `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string `json:"waasDeviceMode,omitempty"`            //
}
type RequestDevicesUpdateInterfaceDetails struct {
	Description string `json:"description,omitempty"` // Description
	AdminStatus string `json:"adminStatus,omitempty"` // Admin Status
	VLANID      *int   `json:"vlanId,omitempty"`      // Vlan Id
	VoiceVLANID *int   `json:"voiceVlanId,omitempty"` // Voice Vlan Id
}
type RequestDevicesClearMacAddressTable struct {
	Operation string                                     `json:"operation,omitempty"` // Operation
	Payload   *RequestDevicesClearMacAddressTablePayload `json:"payload,omitempty"`   // Payload
}
type RequestDevicesClearMacAddressTablePayload interface{}
type RequestDevicesAddDevice2 struct {
	CliTransport            string                                             `json:"cliTransport,omitempty"`            //
	ComputeDevice           *bool                                              `json:"computeDevice,omitempty"`           //
	EnablePassword          string                                             `json:"enablePassword,omitempty"`          //
	ExtendedDiscoveryInfo   string                                             `json:"extendedDiscoveryInfo,omitempty"`   //
	HTTPPassword            string                                             `json:"httpPassword,omitempty"`            //
	HTTPPort                string                                             `json:"httpPort,omitempty"`                //
	HTTPSecure              *bool                                              `json:"httpSecure,omitempty"`              //
	HTTPUserName            string                                             `json:"httpUserName,omitempty"`            //
	IPAddress               []string                                           `json:"ipAddress,omitempty"`               //
	MerakiOrgID             []string                                           `json:"merakiOrgId,omitempty"`             //
	NetconfPort             string                                             `json:"netconfPort,omitempty"`             //
	Password                string                                             `json:"password,omitempty"`                //
	SerialNumber            string                                             `json:"serialNumber,omitempty"`            //
	SNMPAuthPassphrase      string                                             `json:"snmpAuthPassphrase,omitempty"`      //
	SNMPAuthProtocol        string                                             `json:"snmpAuthProtocol,omitempty"`        //
	SNMPMode                string                                             `json:"snmpMode,omitempty"`                //
	SNMPPrivPassphrase      string                                             `json:"snmpPrivPassphrase,omitempty"`      //
	SNMPPrivProtocol        string                                             `json:"snmpPrivProtocol,omitempty"`        //
	SNMPROCommunity         string                                             `json:"snmpROCommunity,omitempty"`         //
	SNMPRWCommunity         string                                             `json:"snmpRWCommunity,omitempty"`         //
	SNMPRetry               *int                                               `json:"snmpRetry,omitempty"`               //
	SNMPTimeout             *int                                               `json:"snmpTimeout,omitempty"`             //
	SNMPUserName            string                                             `json:"snmpUserName,omitempty"`            //
	SNMPVersion             string                                             `json:"snmpVersion,omitempty"`             //
	Type                    string                                             `json:"type,omitempty"`                    //
	UpdateMgmtIPaddressList *[]RequestDevicesAddDevice2UpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //
	UserName                string                                             `json:"userName,omitempty"`                //
}
type RequestDevicesAddDevice2UpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` //
	NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   //
}
type RequestDevicesSyncDevices2 struct {
	CliTransport            string                                               `json:"cliTransport,omitempty"`            //
	ComputeDevice           *bool                                                `json:"computeDevice,omitempty"`           //
	EnablePassword          string                                               `json:"enablePassword,omitempty"`          //
	ExtendedDiscoveryInfo   string                                               `json:"extendedDiscoveryInfo,omitempty"`   //
	HTTPPassword            string                                               `json:"httpPassword,omitempty"`            //
	HTTPPort                string                                               `json:"httpPort,omitempty"`                //
	HTTPSecure              *bool                                                `json:"httpSecure,omitempty"`              //
	HTTPUserName            string                                               `json:"httpUserName,omitempty"`            //
	IPAddress               []string                                             `json:"ipAddress,omitempty"`               //
	MerakiOrgID             []string                                             `json:"merakiOrgId,omitempty"`             //
	NetconfPort             string                                               `json:"netconfPort,omitempty"`             //
	Password                string                                               `json:"password,omitempty"`                //
	SerialNumber            string                                               `json:"serialNumber,omitempty"`            //
	SNMPAuthPassphrase      string                                               `json:"snmpAuthPassphrase,omitempty"`      //
	SNMPAuthProtocol        string                                               `json:"snmpAuthProtocol,omitempty"`        //
	SNMPMode                string                                               `json:"snmpMode,omitempty"`                //
	SNMPPrivPassphrase      string                                               `json:"snmpPrivPassphrase,omitempty"`      //
	SNMPPrivProtocol        string                                               `json:"snmpPrivProtocol,omitempty"`        //
	SNMPROCommunity         string                                               `json:"snmpROCommunity,omitempty"`         //
	SNMPRWCommunity         string                                               `json:"snmpRWCommunity,omitempty"`         //
	SNMPRetry               *int                                                 `json:"snmpRetry,omitempty"`               //
	SNMPTimeout             *int                                                 `json:"snmpTimeout,omitempty"`             //
	SNMPUserName            string                                               `json:"snmpUserName,omitempty"`            //
	SNMPVersion             string                                               `json:"snmpVersion,omitempty"`             //
	Type                    string                                               `json:"type,omitempty"`                    //
	UpdateMgmtIPaddressList *[]RequestDevicesSyncDevices2UpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //
	UserName                string                                               `json:"userName,omitempty"`                //
}
type RequestDevicesSyncDevices2UpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` //
	NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   //
}
type RequestDevicesUpdateDeviceRole struct {
	ID         string `json:"id,omitempty"`         //
	Role       string `json:"role,omitempty"`       //
	RoleSource string `json:"roleSource,omitempty"` //
}
type RequestDevicesExportDeviceList struct {
	DeviceUUIDs   []string `json:"deviceUuids,omitempty"`   //
	ID            string   `json:"id,omitempty"`            //
	OperationEnum string   `json:"operationEnum,omitempty"` //
	Parameters    []string `json:"parameters,omitempty"`    //
	Password      string   `json:"password,omitempty"`      //
}
type RequestDevicesSyncDevices []RequestItemDevicesSyncDevices // Array of RequestDevicesSyncDevices
type RequestItemDevicesSyncDevices interface{}

//GetPlannedAccessPointsForBuilding Get Planned Access Points for Building - b699-9b85-4e3b-acdd
/* Provides a list of Planned Access Points for the Building it is requested for


@param buildingID buildingId path parameter. Building Id

@param GetPlannedAccessPointsForBuildingQueryParams Filtering parameter
*/
func (s *DevicesService) GetPlannedAccessPointsForBuilding(buildingID string, GetPlannedAccessPointsForBuildingQueryParams *GetPlannedAccessPointsForBuildingQueryParams) (*ResponseDevicesGetPlannedAccessPointsForBuilding, *resty.Response, error) {
	path := "/dna/intent/api/v1/buildings/{buildingId}/planned-access-points"
	path = strings.Replace(path, "{buildingId}", fmt.Sprintf("%v", buildingID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForBuildingQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForBuilding{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForBuilding")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForBuilding)
	return result, response, err

}

//GetDeviceDetail Get Device Detail - 89b2-fb14-4f5b-b09b
/* Returns detailed Network Device information retrieved by Mac Address, Device Name or UUID for any given point of time.


@param GetDeviceDetailQueryParams Filtering parameter
*/
func (s *DevicesService) GetDeviceDetail(GetDeviceDetailQueryParams *GetDeviceDetailQueryParams) (*ResponseDevicesGetDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-detail"

	queryString, _ := query.Values(GetDeviceDetailQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceDetail")
	}

	result := response.Result().(*ResponseDevicesGetDeviceDetail)
	return result, response, err

}

//GetDeviceEnrichmentDetails Get Device Enrichment Details - e0b5-599b-4f29-97b7
/* Enriches a given network device context (device id or device Mac Address or device management IP address) with details about the device and neighbor topology


@param GetDeviceEnrichmentDetailsHeaderParams Custom header parameters
*/
func (s *DevicesService) GetDeviceEnrichmentDetails(GetDeviceEnrichmentDetailsHeaderParams *GetDeviceEnrichmentDetailsHeaderParams) (*ResponseDevicesGetDeviceEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDeviceEnrichmentDetailsHeaderParams != nil {

		if GetDeviceEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetDeviceEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetDeviceEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetDeviceEnrichmentDetailsHeaderParams.EntityValue)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetDeviceEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceEnrichmentDetails")
	}

	result := response.Result().(*ResponseDevicesGetDeviceEnrichmentDetails)
	return result, response, err

}

//Devices Devices - a2b4-79a0-4529-8dca
/* Intent API for accessing DNA Assurance Device object for generating reports, creating dashboards or creating additional value added services.


@param DevicesQueryParams Filtering parameter
*/
func (s *DevicesService) Devices(DevicesQueryParams *DevicesQueryParams) (*ResponseDevicesDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-health"

	queryString, _ := query.Values(DevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation Devices")
	}

	result := response.Result().(*ResponseDevicesDevices)
	return result, response, err

}

//GetPlannedAccessPointsForFloor Get Planned Access Points for Floor - 6780-6977-4589-9a54
/* Provides a list of Planned Access Points for the Floor it is requested for


@param floorID floorId path parameter. Floor Id

@param GetPlannedAccessPointsForFloorQueryParams Filtering parameter
*/
func (s *DevicesService) GetPlannedAccessPointsForFloor(floorID string, GetPlannedAccessPointsForFloorQueryParams *GetPlannedAccessPointsForFloorQueryParams) (*ResponseDevicesGetPlannedAccessPointsForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForFloorQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForFloor{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForFloor")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForFloor)
	return result, response, err

}

//GetAllInterfaces Get all interfaces - f594-7a4c-439a-8bf0
/* Returns all available interfaces. This endpoint can return a maximum of 500 interfaces


@param GetAllInterfacesQueryParams Filtering parameter
*/
func (s *DevicesService) GetAllInterfaces(GetAllInterfacesQueryParams *GetAllInterfacesQueryParams) (*ResponseDevicesGetAllInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface"

	queryString, _ := query.Values(GetAllInterfacesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAllInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetAllInterfaces)
	return result, response, err

}

//GetDeviceInterfaceCount Get Device Interface Count - 3d92-3b18-4dc9-a4ca
/* Returns the count of interfaces for all devices


 */
func (s *DevicesService) GetDeviceInterfaceCount() (*ResponseDevicesGetDeviceInterfaceCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCount")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCount)
	return result, response, err

}

//GetInterfaceByIP Get Interface by IP - cd84-69e6-47ca-ab0e
/* Returns list of interfaces for specified device management IP address


@param ipAddress ipAddress path parameter. IP address of the interface

*/
func (s *DevicesService) GetInterfaceByIP(ipAddress string) (*ResponseDevicesGetInterfaceByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInterfaceByIp")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByIP)
	return result, response, err

}

//GetIsisInterfaces Get ISIS interfaces - 84ad-8b0e-42ca-b48a
/* Returns the interfaces that has ISIS enabled


 */
func (s *DevicesService) GetIsisInterfaces() (*ResponseDevicesGetIsisInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/isis"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetIsisInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIsisInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetIsisInterfaces)
	return result, response, err

}

//GetInterfaceInfoByID Get Interface info by Id - ba9d-c85b-4b8a-9a17
/* Returns list of interfaces by specified device


@param deviceID deviceId path parameter. Device ID

*/
func (s *DevicesService) GetInterfaceInfoByID(deviceID string) (*ResponseDevicesGetInterfaceInfoByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInterfaceInfoById")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceInfoByID)
	return result, response, err

}

//GetDeviceInterfaceCount2 Get Device Interface count - 5b86-3922-4cd8-8ea7
/* Returns the interface count for the given device


@param deviceID deviceId path parameter. Device ID

*/
func (s *DevicesService) GetDeviceInterfaceCount2(deviceID string) (*ResponseDevicesGetDeviceInterfaceCount2, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/count"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCount2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCount2")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCount2)
	return result, response, err

}

//GetInterfaceDetailsByDeviceIDAndInterfaceName Get Interface details by device Id and interface name - 4eb5-6a61-4cc9-a2d2
/* Returns interface by specified device Id and interface name


@param deviceID deviceId path parameter. Device ID

@param GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams Filtering parameter
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID string, GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams) (*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/interface-name"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInterfaceDetailsByDeviceIdAndInterfaceName")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName)
	return result, response, err

}

//GetDeviceInterfacesBySpecifiedRange Get Device Interfaces by specified range - 349c-8884-43b8-9a58
/* Returns the list of interfaces for the device for the specified range


@param deviceID deviceId path parameter. Device ID

@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return

*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRange(deviceID string, startIndex int, recordsToReturn int) (*ResponseDevicesGetDeviceInterfacesBySpecifiedRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfacesBySpecifiedRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfacesBySpecifiedRange")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfacesBySpecifiedRange)
	return result, response, err

}

//GetOspfInterfaces Get OSPF interfaces - 70ad-3976-49e9-b4d3
/* Returns the interfaces that has OSPF enabled


 */
func (s *DevicesService) GetOspfInterfaces() (*ResponseDevicesGetOspfInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ospf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOspfInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOspfInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetOspfInterfaces)
	return result, response, err

}

//GetInterfaceByID Get Interface by Id - b888-792d-43ba-ba46
/* Returns the interface for the given interface ID


@param id id path parameter. Interface ID

*/
func (s *DevicesService) GetInterfaceByID(id string) (*ResponseDevicesGetInterfaceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetInterfaceById")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByID)
	return result, response, err

}

//LegitOperationsForInterface Legit operations for interface - 87a3-3a52-46ea-a40e
/* Get list of all properties & operations valid for an interface.


@param interfaceUUID interfaceUuid path parameter. Interface ID

*/
func (s *DevicesService) LegitOperationsForInterface(interfaceUUID string) (*ResponseDevicesLegitOperationsForInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/legit-operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesLegitOperationsForInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LegitOperationsForInterface")
	}

	result := response.Result().(*ResponseDevicesLegitOperationsForInterface)
	return result, response, err

}

//GetDeviceList Get Device list - 20b1-9b52-464b-8972
/* Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc. You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-device?hostname=myhost.*&managementIpAddress=192.25.18..*
If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and ignores the other request parameters. You can also specify offset & limit to get the required list.


@param GetDeviceListQueryParams Filtering parameter
*/
func (s *DevicesService) GetDeviceList(GetDeviceListQueryParams *GetDeviceListQueryParams) (*ResponseDevicesGetDeviceList, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	queryString, _ := query.Values(GetDeviceListQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceList")
	}

	result := response.Result().(*ResponseDevicesGetDeviceList)
	return result, response, err

}

//GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute Get Device Values that match fully or partially an Attribute - ffa7-48cc-44e9-a437
/* Returns the list of values of the first given required parameter. You can use the .* in any value to conduct a wildcard search. For example, to get all the devices with the management IP address starting with 10.10. , issue the following request: GET /dna/inten/api/v1/network-device/autocomplete?managementIpAddress=10.10..* It will return the device management IP addresses that match fully or partially the provided attribute. {[10.10.1.1, 10.10.20.2, …]}.


@param GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams Filtering parameter
*/
func (s *DevicesService) GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams *GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams) (*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/autocomplete"

	queryString, _ := query.Values(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute")
	}

	result := response.Result().(*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute)
	return result, response, err

}

//GetPollingIntervalForAllDevices Get Polling Interval for all devices - 38bd-0b88-4b89-a785
/* Returns polling interval of all devices


 */
func (s *DevicesService) GetPollingIntervalForAllDevices() (*ResponseDevicesGetPollingIntervalForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/collection-schedule/global"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalForAllDevices")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalForAllDevices)
	return result, response, err

}

//GetDeviceConfigForAllDevices Get Device Config for all devices - b7bc-aa08-4e2b-90d0
/* Returns the config for all devices


 */
func (s *DevicesService) GetDeviceConfigForAllDevices() (*ResponseDevicesGetDeviceConfigForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigForAllDevices")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigForAllDevices)
	return result, response, err

}

//GetDeviceConfigCount Get Device Config Count - 888f-585c-49b8-8441
/* Returns the count of device configs


 */
func (s *DevicesService) GetDeviceConfigCount() (*ResponseDevicesGetDeviceConfigCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigCount")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigCount)
	return result, response, err

}

//GetDeviceCount2 Get Device Count - 5db2-1b8e-43fa-b7d8
/* Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and location name


 */
func (s *DevicesService) GetDeviceCount2() (*ResponseDevicesGetDeviceCount2, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceCount2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCount2")
	}

	result := response.Result().(*ResponseDevicesGetDeviceCount2)
	return result, response, err

}

//GetFunctionalCapabilityForDevices Get Functional Capability for devices - c3b3-c9ef-4e6b-8a09
/* Returns the functional-capability for given devices


@param GetFunctionalCapabilityForDevicesQueryParams Filtering parameter
*/
func (s *DevicesService) GetFunctionalCapabilityForDevices(GetFunctionalCapabilityForDevicesQueryParams *GetFunctionalCapabilityForDevicesQueryParams) (*ResponseDevicesGetFunctionalCapabilityForDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability"

	queryString, _ := query.Values(GetFunctionalCapabilityForDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetFunctionalCapabilityForDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityForDevices")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityForDevices)
	return result, response, err

}

//GetFunctionalCapabilityByID Get Functional Capability by Id - 81bb-4804-405a-8d2f
/* Returns functional capability with given Id


@param id id path parameter. Functional Capability UUID

*/
func (s *DevicesService) GetFunctionalCapabilityByID(id string) (*ResponseDevicesGetFunctionalCapabilityByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetFunctionalCapabilityByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityById")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityByID)
	return result, response, err

}

//InventoryInsightDeviceLinkMismatchApI Inventory Insight Device Link Mismatch API - 5792-59d8-4208-8190
/* Find all devices with link mismatch (speed /  vlan)


@param siteID siteId path parameter.
@param InventoryInsightDeviceLinkMismatchAPIQueryParams Filtering parameter
*/
func (s *DevicesService) InventoryInsightDeviceLinkMismatchApI(siteID string, InventoryInsightDeviceLinkMismatchAPIQueryParams *InventoryInsightDeviceLinkMismatchApIQueryParams) (*ResponseDevicesInventoryInsightDeviceLinkMismatchApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/insight/{siteId}/device-link"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(InventoryInsightDeviceLinkMismatchAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesInventoryInsightDeviceLinkMismatchApI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation InventoryInsightDeviceLinkMismatchApI")
	}

	result := response.Result().(*ResponseDevicesInventoryInsightDeviceLinkMismatchApI)
	return result, response, err

}

//ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES Returns devices added to Cisco DNA center with snmp v3 DES. - afba-7a69-4d38-8de1
/* Returns devices added to Cisco DNA center with snmp v3 DES, where siteId is mandatory & accepts offset, limit, sortby, order which are optional.


@param siteID siteId path parameter.
@param ReturnsDevicesAddedToCiscoDNACenterWithSnmpV3DESQueryParams Filtering parameter
*/
func (s *DevicesService) ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES(siteID string, ReturnsDevicesAddedToCiscoDNACenterWithSnmpV3DESQueryParams *ReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DESQueryParams) (*ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/insight/{siteId}/insecure-connection"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(ReturnsDevicesAddedToCiscoDNACenterWithSnmpV3DESQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ReturnsDevicesAddedToCiscoDnaCenterWithSnmpV3DES")
	}

	result := response.Result().(*ResponseDevicesReturnsDevicesAddedToCiscoDnaCenterWithSNMPV3DES)
	return result, response, err

}

//GetNetworkDeviceByIP Get Network Device by IP - d0a4-b881-45aa-bb51
/* Returns the network device by specified IP address


@param ipAddress ipAddress path parameter. Device IP address

*/
func (s *DevicesService) GetNetworkDeviceByIP(ipAddress string) (*ResponseDevicesGetNetworkDeviceByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByIp")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByIP)
	return result, response, err

}

//GetModules Get Modules - eb82-49e3-4f69-b0f1
/* Returns modules by specified device id


@param GetModulesQueryParams Filtering parameter
*/
func (s *DevicesService) GetModules(GetModulesQueryParams *GetModulesQueryParams) (*ResponseDevicesGetModules, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module"

	queryString, _ := query.Values(GetModulesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetModules")
	}

	result := response.Result().(*ResponseDevicesGetModules)
	return result, response, err

}

//GetModuleCount Get Module count - 8db9-3974-4649-a782
/* Returns Module Count


@param GetModuleCountQueryParams Filtering parameter
*/
func (s *DevicesService) GetModuleCount(GetModuleCountQueryParams *GetModuleCountQueryParams) (*ResponseDevicesGetModuleCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/count"

	queryString, _ := query.Values(GetModuleCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModuleCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetModuleCount")
	}

	result := response.Result().(*ResponseDevicesGetModuleCount)
	return result, response, err

}

//GetModuleInfoByID Get Module Info by Id - 0db7-da74-4c0b-83d8
/* Returns Module info by 'module id'


@param id id path parameter. Module id

*/
func (s *DevicesService) GetModuleInfoByID(id string) (*ResponseDevicesGetModuleInfoByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetModuleInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetModuleInfoById")
	}

	result := response.Result().(*ResponseDevicesGetModuleInfoByID)
	return result, response, err

}

//GetDeviceBySerialNumber Get Device by Serial number - d888-ab6d-4d59-a8c1
/* Returns the network device with given serial number


@param serialNumber serialNumber path parameter. Device serial number

*/
func (s *DevicesService) GetDeviceBySerialNumber(serialNumber string) (*ResponseDevicesGetDeviceBySerialNumber, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{serialNumber}", fmt.Sprintf("%v", serialNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceBySerialNumber{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceBySerialNumber")
	}

	result := response.Result().(*ResponseDevicesGetDeviceBySerialNumber)
	return result, response, err

}

//RegisterDeviceForWsa Register device for WSA - c980-9b67-44f8-a502
/* Registers a device for WSA notification


@param RegisterDeviceForWSAQueryParams Filtering parameter
*/
func (s *DevicesService) RegisterDeviceForWsa(RegisterDeviceForWSAQueryParams *RegisterDeviceForWsaQueryParams) (*ResponseDevicesRegisterDeviceForWsa, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/tenantinfo/macaddress"

	queryString, _ := query.Values(RegisterDeviceForWSAQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRegisterDeviceForWsa{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RegisterDeviceForWsa")
	}

	result := response.Result().(*ResponseDevicesRegisterDeviceForWsa)
	return result, response, err

}

//GetChassisDetailsForDevice Get Chassis Details for Device - 0486-9b26-49ab-b579
/* Returns chassis details for given device ID


@param deviceID deviceId path parameter. Device ID

*/
func (s *DevicesService) GetChassisDetailsForDevice(deviceID string) (*ResponseDevicesGetChassisDetailsForDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/chassis"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetChassisDetailsForDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetChassisDetailsForDevice")
	}

	result := response.Result().(*ResponseDevicesGetChassisDetailsForDevice)
	return result, response, err

}

//GetStackDetailsForDevice Get Stack Details for Device - 78a7-7ab0-4d5a-8a10
/* Retrieves complete stack details for given device ID


@param deviceID deviceId path parameter. Device ID

*/
func (s *DevicesService) GetStackDetailsForDevice(deviceID string) (*ResponseDevicesGetStackDetailsForDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/stack"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetStackDetailsForDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetStackDetailsForDevice")
	}

	result := response.Result().(*ResponseDevicesGetStackDetailsForDevice)
	return result, response, err

}

//ReturnPowerSupplyFanDetailsForTheGivenDevice Return PowerSupply/ Fan details for the Given device - 20b1-9b52-464b-897a
/* Return PowerSupply/ Fan details for the Given device


@param deviceUUID deviceUuid path parameter.
@param ReturnPowerSupplyFanDetailsForTheGivenDeviceQueryParams Filtering parameter
*/
func (s *DevicesService) ReturnPowerSupplyFanDetailsForTheGivenDevice(deviceUUID string, ReturnPowerSupplyFanDetailsForTheGivenDeviceQueryParams *ReturnPowerSupplyFanDetailsForTheGivenDeviceQueryParams) (*ResponseDevicesReturnPowerSupplyFanDetailsForTheGivenDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/equipment"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ReturnPowerSupplyFanDetailsForTheGivenDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesReturnPowerSupplyFanDetailsForTheGivenDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ReturnPowerSupplyFanDetailsForTheGivenDevice")
	}

	result := response.Result().(*ResponseDevicesReturnPowerSupplyFanDetailsForTheGivenDevice)
	return result, response, err

}

//ReturnsPoeInterfaceDetailsForTheDevice Returns POE interface details for the device. - 20b5-48af-42da-a337
/* Returns POE interface details for the device, where deviceuuid is mandatory & accepts comma seperated interface names which is optional and returns information for that particular interfaces where(operStatus = operationalStatus)


@param deviceUUID deviceUuid path parameter. uuid of the device

@param ReturnsPOEInterfaceDetailsForTheDeviceQueryParams Filtering parameter
*/
func (s *DevicesService) ReturnsPoeInterfaceDetailsForTheDevice(deviceUUID string, ReturnsPOEInterfaceDetailsForTheDeviceQueryParams *ReturnsPoeInterfaceDetailsForTheDeviceQueryParams) (*ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/poe-detail"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ReturnsPOEInterfaceDetailsForTheDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ReturnsPoeInterfaceDetailsForTheDevice")
	}

	result := response.Result().(*ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice)
	return result, response, err

}

//GetConnectedDeviceDetail Get connected device detail - a8aa-ca21-4c09-8388
/* Get connected device detail for given deviceUuid and interfaceUuid


@param deviceUUID deviceUuid path parameter. instanceuuid of Device

@param interfaceUUID interfaceUuid path parameter. instanceuuid of interface

*/
func (s *DevicesService) GetConnectedDeviceDetail(deviceUUID string, interfaceUUID string) (*ResponseDevicesGetConnectedDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/{interfaceUuid}/neighbor"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetConnectedDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectedDeviceDetail")
	}

	result := response.Result().(*ResponseDevicesGetConnectedDeviceDetail)
	return result, response, err

}

//GetLinecardDetails Get Linecard details - 46a1-4b02-48fb-8fbf
/* Get line card detail for a given deviceuuid.  Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device

*/
func (s *DevicesService) GetLinecardDetails(deviceUUID string) (*ResponseDevicesGetLinecardDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/line-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetLinecardDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetLinecardDetails")
	}

	result := response.Result().(*ResponseDevicesGetLinecardDetails)
	return result, response, err

}

//PoeDetails POE details  - 8ba6-7932-4ed9-abae
/* Returns POE details for device.


@param deviceUUID deviceUuid path parameter. uuid of the device

*/
func (s *DevicesService) PoeDetails(deviceUUID string) (*ResponseDevicesPoeDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/poe"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesPoeDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PoeDetails")
	}

	result := response.Result().(*ResponseDevicesPoeDetails)
	return result, response, err

}

//GetSupervisorCardDetail Get Supervisor card detail - 88aa-1b52-4a38-bf97
/* Get supervisor card detail for a given deviceuuid. Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device

*/
func (s *DevicesService) GetSupervisorCardDetail(deviceUUID string) (*ResponseDevicesGetSupervisorCardDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/supervisor-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetSupervisorCardDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSupervisorCardDetail")
	}

	result := response.Result().(*ResponseDevicesGetSupervisorCardDetail)
	return result, response, err

}

//GetDeviceByID Get Device by ID - 8fa8-eb40-4a4a-8d96
/* Returns the network device details for the given device ID


@param id id path parameter. Device ID

*/
func (s *DevicesService) GetDeviceByID(id string) (*ResponseDevicesGetDeviceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceById")
	}

	result := response.Result().(*ResponseDevicesGetDeviceByID)
	return result, response, err

}

//GetDeviceSummary Get Device Summary - 819f-9aa5-4fea-b7bf
/* Returns brief summary of device info such as hostname, management IP address for the given device Id


@param id id path parameter. Device ID

*/
func (s *DevicesService) GetDeviceSummary(id string) (*ResponseDevicesGetDeviceSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/brief"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceSummary")
	}

	result := response.Result().(*ResponseDevicesGetDeviceSummary)
	return result, response, err

}

//GetPollingIntervalByID Get Polling Interval by Id - 8291-8a1b-4d28-9c5c
/* Returns polling interval by device id


@param id id path parameter. Device ID

*/
func (s *DevicesService) GetPollingIntervalByID(id string) (*ResponseDevicesGetPollingIntervalByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/collection-schedule"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalById")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalByID)
	return result, response, err

}

//GetOrganizationListForMeraki Get Organization list for Meraki - 84b3-7ae5-4c59-ab28
/* Returns list of organizations for meraki dashboard


@param id id path parameter.
*/
func (s *DevicesService) GetOrganizationListForMeraki(id string) (*ResponseDevicesGetOrganizationListForMeraki, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/meraki-organization"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationListForMeraki{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetOrganizationListForMeraki")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationListForMeraki)
	return result, response, err

}

//GetDeviceInterfaceVLANs Get Device Interface VLANs - 288d-f949-4f2a-9746
/* Returns Device Interface VLANs


@param id id path parameter.
@param GetDeviceInterfaceVLANsQueryParams Filtering parameter
*/
func (s *DevicesService) GetDeviceInterfaceVLANs(id string, GetDeviceInterfaceVLANsQueryParams *GetDeviceInterfaceVLANsQueryParams) (*ResponseDevicesGetDeviceInterfaceVLANs, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/vlan"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDeviceInterfaceVLANsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceInterfaceVLANs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceVlans")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceVLANs)
	return result, response, err

}

//GetWirelessLanControllerDetailsByID Get wireless lan controller details by Id - f682-6a8e-41bb-a242
/* Returns the wireless lan controller info with given device ID


@param id id path parameter. Device ID

*/
func (s *DevicesService) GetWirelessLanControllerDetailsByID(id string) (*ResponseDevicesGetWirelessLanControllerDetailsByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/wireless-info"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetWirelessLanControllerDetailsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetWirelessLanControllerDetailsById")
	}

	result := response.Result().(*ResponseDevicesGetWirelessLanControllerDetailsByID)
	return result, response, err

}

//GetDeviceConfigByID Get Device Config by Id - 84b3-3a9e-480a-bcaf
/* Returns the device config by specified device ID


@param networkDeviceID networkDeviceId path parameter.
*/
func (s *DevicesService) GetDeviceConfigByID(networkDeviceID string) (*ResponseDevicesGetDeviceConfigByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{networkDeviceId}/config"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigById")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigByID)
	return result, response, err

}

//GetNetworkDeviceByPaginationRange Get Network Device by pagination range - f495-48c5-4be8-a3e2
/* Returns the list of network devices for the given pagination range


@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return

*/
func (s *DevicesService) GetNetworkDeviceByPaginationRange(startIndex int, recordsToReturn int) (*ResponseDevicesGetNetworkDeviceByPaginationRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByPaginationRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByPaginationRange")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByPaginationRange)
	return result, response, err

}

//ClearMacAddressTable Clear Mac-Address table - 24be-a97f-43f9-bc65
/* Clear mac-address on an individual port. In request body, operation needs to be specified as 'ClearMacAddress'. In the future more possible operations will be added to this API


@param interfaceUUID interfaceUuid path parameter. Interface Id

@param ClearMacAddressTableQueryParams Filtering parameter
*/
func (s *DevicesService) ClearMacAddressTable(interfaceUUID string, requestDevicesClearMacAddressTable *RequestDevicesClearMacAddressTable, ClearMacAddressTableQueryParams *ClearMacAddressTableQueryParams) (*ResponseDevicesClearMacAddressTable, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(ClearMacAddressTableQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesClearMacAddressTable).
		SetResult(&ResponseDevicesClearMacAddressTable{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ClearMacAddressTable")
	}

	result := response.Result().(*ResponseDevicesClearMacAddressTable)
	return result, response, err

}

//AddDevice2 Add Device - 4bb2-2af0-46fa-8f08
/* Adds the device with given credential


 */
func (s *DevicesService) AddDevice2(requestDevicesAddDevice2 *RequestDevicesAddDevice2) (*ResponseDevicesAddDevice2, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddDevice2).
		SetResult(&ResponseDevicesAddDevice2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddDevice2")
	}

	result := response.Result().(*ResponseDevicesAddDevice2)
	return result, response, err

}

//ExportDeviceList Export Device list - cd98-780f-4888-a66d
/* Exports the selected network device to a file


 */
func (s *DevicesService) ExportDeviceList(requestDevicesExportDeviceList *RequestDevicesExportDeviceList) (*ResponseDevicesExportDeviceList, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/file"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesExportDeviceList).
		SetResult(&ResponseDevicesExportDeviceList{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ExportDeviceList")
	}

	result := response.Result().(*ResponseDevicesExportDeviceList)
	return result, response, err

}

//UpdateInterfaceDetails Update Interface details - 868b-5a60-4be8-a2d7
/* Add/Update Interface description, VLAN membership, Voice VLAN and change Interface admin status ('UP'/'DOWN') from Request body.


@param interfaceUUID interfaceUuid path parameter. Interface ID

@param UpdateInterfaceDetailsQueryParams Filtering parameter
*/
func (s *DevicesService) UpdateInterfaceDetails(interfaceUUID string, requestDevicesUpdateInterfaceDetails *RequestDevicesUpdateInterfaceDetails, UpdateInterfaceDetailsQueryParams *UpdateInterfaceDetailsQueryParams) (*ResponseDevicesUpdateInterfaceDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(UpdateInterfaceDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesUpdateInterfaceDetails).
		SetResult(&ResponseDevicesUpdateInterfaceDetails{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateInterfaceDetails")
	}

	result := response.Result().(*ResponseDevicesUpdateInterfaceDetails)
	return result, response, err

}

//SyncDevices2 Sync Devices - aeb9-eb67-460b-92df
/* Sync the devices provided as input


 */
func (s *DevicesService) SyncDevices2(requestDevicesSyncDevices2 *RequestDevicesSyncDevices2) (*ResponseDevicesSyncDevices2, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesSyncDevices2).
		SetResult(&ResponseDevicesSyncDevices2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SyncDevices2")
	}

	result := response.Result().(*ResponseDevicesSyncDevices2)
	return result, response, err

}

//UpdateDeviceRole Update Device role - b985-5ad5-4ae9-8156
/* Updates the role of the device as access, core, distribution, border router


 */
func (s *DevicesService) UpdateDeviceRole(requestDevicesUpdateDeviceRole *RequestDevicesUpdateDeviceRole) (*ResponseDevicesUpdateDeviceRole, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/brief"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceRole).
		SetResult(&ResponseDevicesUpdateDeviceRole{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDeviceRole")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceRole)
	return result, response, err

}

//SyncDevices Sync Devices - 3b9e-f967-4429-be4c
/* Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result can be seen in the child task of each device


@param SyncDevicesQueryParams Filtering parameter
*/
func (s *DevicesService) SyncDevices(requestDevicesSyncDevices *RequestDevicesSyncDevices, SyncDevicesQueryParams *SyncDevicesQueryParams) (*ResponseDevicesSyncDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/sync"

	queryString, _ := query.Values(SyncDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesSyncDevices).
		SetResult(&ResponseDevicesSyncDevices{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SyncDevices")
	}

	result := response.Result().(*ResponseDevicesSyncDevices)
	return result, response, err

}

//DeleteDeviceByID Delete Device by Id - 1c89-4b58-48ea-b214
/* Deletes the network device for the given Id


@param id id path parameter. Device ID

@param DeleteDeviceByIdQueryParams Filtering parameter
*/
func (s *DevicesService) DeleteDeviceByID(id string, DeleteDeviceByIdQueryParams *DeleteDeviceByIDQueryParams) (*ResponseDevicesDeleteDeviceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(DeleteDeviceByIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDeleteDeviceByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceById")
	}

	result := response.Result().(*ResponseDevicesDeleteDeviceByID)
	return result, response, err

}
