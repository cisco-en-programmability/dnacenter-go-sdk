package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DevicesService is the service to communicate with the Devices API endpoint
type DevicesService service

// AddDeviceRequest is the AddDeviceRequest definition
type AddDeviceRequest struct {
	CLITransport            string   `json:"cliTransport,omitempty"`          //
	ComputeDevice           bool     `json:"computeDevice,omitempty"`         //
	EnablePassword          string   `json:"enablePassword,omitempty"`        //
	ExtendedDiscoveryInfo   string   `json:"extendedDiscoveryInfo,omitempty"` //
	HTTPPassword            string   `json:"httpPassword,omitempty"`          //
	HTTPPort                string   `json:"httpPort,omitempty"`              //
	HTTPSecure              bool     `json:"httpSecure,omitempty"`            //
	HTTPUserName            string   `json:"httpUserName,omitempty"`          //
	IPAddress               []string `json:"ipAddress,omitempty"`             //
	MerakiOrgID             []string `json:"merakiOrgId,omitempty"`           //
	NetconfPort             string   `json:"netconfPort,omitempty"`           //
	Password                string   `json:"password,omitempty"`              //
	SerialNumber            string   `json:"serialNumber,omitempty"`          //
	SNMPAuthPassphrase      string   `json:"snmpAuthPassphrase,omitempty"`    //
	SNMPAuthProtocol        string   `json:"snmpAuthProtocol,omitempty"`      //
	SNMPMode                string   `json:"snmpMode,omitempty"`              //
	SNMPPrivPassphrase      string   `json:"snmpPrivPassphrase,omitempty"`    //
	SNMPPrivProtocol        string   `json:"snmpPrivProtocol,omitempty"`      //
	SNMPROCommunity         string   `json:"snmpROCommunity,omitempty"`       //
	SNMPRWCommunity         string   `json:"snmpRWCommunity,omitempty"`       //
	SNMPRetry               int      `json:"snmpRetry,omitempty"`             //
	SNMPTimeout             int      `json:"snmpTimeout,omitempty"`           //
	SNMPUserName            string   `json:"snmpUserName,omitempty"`          //
	SNMPVersion             string   `json:"snmpVersion,omitempty"`           //
	Type                    string   `json:"type,omitempty"`                  //
	UpdateMgmtIPaddressList []struct {
		ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` //
		NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   //
	} `json:"updateMgmtIPaddressList,omitempty"` //
	UserName string `json:"userName,omitempty"` //
}

// ExportDeviceListRequest is the ExportDeviceListRequest definition
type ExportDeviceListRequest struct {
	DeviceUUIDs   []string `json:"deviceUuids,omitempty"`   //
	ID            string   `json:"id,omitempty"`            //
	OperationEnum string   `json:"operationEnum,omitempty"` //
	Parameters    []string `json:"parameters,omitempty"`    //
	Password      string   `json:"password,omitempty"`      //
}

// SyncDevicesRequest is the SyncDevicesRequest definition
type SyncDevicesRequest struct {
	CLITransport            string   `json:"cliTransport,omitempty"`          //
	ComputeDevice           bool     `json:"computeDevice,omitempty"`         //
	EnablePassword          string   `json:"enablePassword,omitempty"`        //
	ExtendedDiscoveryInfo   string   `json:"extendedDiscoveryInfo,omitempty"` //
	HTTPPassword            string   `json:"httpPassword,omitempty"`          //
	HTTPPort                string   `json:"httpPort,omitempty"`              //
	HTTPSecure              bool     `json:"httpSecure,omitempty"`            //
	HTTPUserName            string   `json:"httpUserName,omitempty"`          //
	IPAddress               []string `json:"ipAddress,omitempty"`             //
	MerakiOrgID             []string `json:"merakiOrgId,omitempty"`           //
	NetconfPort             string   `json:"netconfPort,omitempty"`           //
	Password                string   `json:"password,omitempty"`              //
	SerialNumber            string   `json:"serialNumber,omitempty"`          //
	SNMPAuthPassphrase      string   `json:"snmpAuthPassphrase,omitempty"`    //
	SNMPAuthProtocol        string   `json:"snmpAuthProtocol,omitempty"`      //
	SNMPMode                string   `json:"snmpMode,omitempty"`              //
	SNMPPrivPassphrase      string   `json:"snmpPrivPassphrase,omitempty"`    //
	SNMPPrivProtocol        string   `json:"snmpPrivProtocol,omitempty"`      //
	SNMPROCommunity         string   `json:"snmpROCommunity,omitempty"`       //
	SNMPRWCommunity         string   `json:"snmpRWCommunity,omitempty"`       //
	SNMPRetry               int      `json:"snmpRetry,omitempty"`             //
	SNMPTimeout             int      `json:"snmpTimeout,omitempty"`           //
	SNMPUserName            string   `json:"snmpUserName,omitempty"`          //
	SNMPVersion             string   `json:"snmpVersion,omitempty"`           //
	Type                    string   `json:"type,omitempty"`                  //
	UpdateMgmtIPaddressList []struct {
		ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` //
		NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   //
	} `json:"updateMgmtIPaddressList,omitempty"` //
	UserName string `json:"userName,omitempty"` //
}

// SyncNetworkDevicesRequest is the SyncNetworkDevicesRequest definition
type SyncNetworkDevicesRequest []string

// UpdateDeviceRoleRequest is the UpdateDeviceRoleRequest definition
type UpdateDeviceRoleRequest struct {
	ID         string `json:"id,omitempty"`         //
	Role       string `json:"role,omitempty"`       //
	RoleSource string `json:"roleSource,omitempty"` //
}

// AddDeviceResponse is the AddDeviceResponse definition
type AddDeviceResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteDeviceByIDResponse is the DeleteDeviceByIdResponse definition
type DeleteDeviceByIDResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DevicesResponse is the DevicesResponse definition
type DevicesResponse struct {
	Response []struct {
		AirQualityHealth struct {
			Ghz24  int `json:"Ghz24,omitempty"`  //
			Ghz50  int `json:"Ghz50,omitempty"`  //
			Radio0 int `json:"radio0,omitempty"` //
			Radio1 int `json:"radio1,omitempty"` //
		} `json:"airQualityHealth,omitempty"` //
		CLIentCount struct {
			Ghz24  int `json:"Ghz24,omitempty"`  //
			Ghz50  int `json:"Ghz50,omitempty"`  //
			Radio0 int `json:"radio0,omitempty"` //
			Radio1 int `json:"radio1,omitempty"` //
		} `json:"clientCount,omitempty"` //
		CPUHealth                  int    `json:"cpuHealth,omitempty"`                  //
		CPUUlitilization           int    `json:"cpuUlitilization,omitempty"`           //
		DeviceFamily               string `json:"deviceFamily,omitempty"`               //
		DeviceType                 string `json:"deviceType,omitempty"`                 //
		InterDeviceLinkAvailHealth int    `json:"interDeviceLinkAvailHealth,omitempty"` //
		InterfaceLinkErrHealth     int    `json:"interfaceLinkErrHealth,omitempty"`     //
		InterferenceHealth         struct {
			Ghz24  int `json:"Ghz24,omitempty"`  //
			Ghz50  int `json:"Ghz50,omitempty"`  //
			Radio0 int `json:"radio0,omitempty"` //
			Radio1 int `json:"radio1,omitempty"` //
		} `json:"interferenceHealth,omitempty"` //
		IPAddress               string `json:"ipAddress,omitempty"`               //
		IssueCount              int    `json:"issueCount,omitempty"`              //
		Location                string `json:"location,omitempty"`                //
		MacAddress              string `json:"macAddress,omitempty"`              //
		MemoryUtilization       int    `json:"memoryUtilization,omitempty"`       //
		MemoryUtilizationHealth int    `json:"memoryUtilizationHealth,omitempty"` //
		Model                   string `json:"model,omitempty"`                   //
		Name                    string `json:"name,omitempty"`                    //
		NoiseHealth             struct {
			Ghz50  int `json:"Ghz50,omitempty"`  //
			Radio1 int `json:"radio1,omitempty"` //
		} `json:"noiseHealth,omitempty"` //
		OsVersion          string `json:"osVersion,omitempty"`          //
		OverallHealth      int    `json:"overallHealth,omitempty"`      //
		ReachabilityHealth string `json:"reachabilityHealth,omitempty"` //
		UtilizationHealth  struct {
			Ghz24  int `json:"Ghz24,omitempty"`  //
			Ghz50  int `json:"Ghz50,omitempty"`  //
			Radio0 int `json:"radio0,omitempty"` //
			Radio1 int `json:"radio1,omitempty"` //
		} `json:"utilizationHealth,omitempty"` //
	} `json:"response,omitempty"` //
	TotalCount int    `json:"totalCount,omitempty"` //
	Version    string `json:"version,omitempty"`    //
}

// ExportDeviceListResponse is the ExportDeviceListResponse definition
type ExportDeviceListResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetAllInterfacesResponse is the GetAllInterfacesResponse definition
type GetAllInterfacesResponse struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DevicesGetDeviceByIDResponse is the GetDeviceByIDResponse definition
type DevicesGetDeviceByIDResponse struct {
	Response struct {
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
		LastUpdateTime            string `json:"lastUpdateTime,omitempty"`            //
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDeviceBySerialNumberResponse is the GetDeviceBySerialNumberResponse definition
type GetDeviceBySerialNumberResponse struct {
	Response struct {
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
		LastUpdateTime            string `json:"lastUpdateTime,omitempty"`            //
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDeviceConfigByIDResponse is the GetDeviceConfigByIdResponse definition
type GetDeviceConfigByIDResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDeviceConfigCountResponse is the GetDeviceConfigCountResponse definition
type GetDeviceConfigCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDeviceConfigForAllDevicesResponse is the GetDeviceConfigForAllDevicesResponse definition
type GetDeviceConfigForAllDevicesResponse struct {
	Response []struct {
		AttributeInfo   string `json:"attributeInfo,omitempty"`   //
		CdpNeighbors    string `json:"cdpNeighbors,omitempty"`    //
		HealthMonitor   string `json:"healthMonitor,omitempty"`   //
		ID              string `json:"id,omitempty"`              //
		IntfDescription string `json:"intfDescription,omitempty"` //
		Inventory       string `json:"inventory,omitempty"`       //
		IPIntfBrief     string `json:"ipIntfBrief,omitempty"`     //
		MacAddressTable string `json:"macAddressTable,omitempty"` //
		RunningConfig   string `json:"runningConfig,omitempty"`   //
		SNMP            string `json:"snmp,omitempty"`            //
		Version         string `json:"version,omitempty"`         //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDeviceCountResponse is the GetDeviceCountResponse definition
type GetDeviceCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDeviceDetailResponse is the GetDeviceDetailResponse definition
type GetDeviceDetailResponse struct {
	Response struct {
		HALastResetReason      string `json:"HALastResetReason,omitempty"`      //
		HAPrimaryPowerStatus   string `json:"HAPrimaryPowerStatus,omitempty"`   //
		HASecondaryPowerStatus string `json:"HASecondaryPowerStatus,omitempty"` //
		AirQuality             string `json:"airQuality,omitempty"`             //
		AirQualityScore        int    `json:"airQualityScore,omitempty"`        //
		CLIentCount            string `json:"clientCount,omitempty"`            //
		CollectionStatus       string `json:"collectionStatus,omitempty"`       //
		CommunicationState     string `json:"communicationState,omitempty"`     //
		CPU                    string `json:"cpu,omitempty"`                    //
		CPUScore               int    `json:"cpuScore,omitempty"`               //
		DeviceSeries           string `json:"deviceSeries,omitempty"`           //
		FreeMbuf               string `json:"freeMbuf,omitempty"`               //
		FreeMbufScore          int    `json:"freeMbufScore,omitempty"`          //
		FreeTimer              string `json:"freeTimer,omitempty"`              //
		FreeTimerScore         int    `json:"freeTimerScore,omitempty"`         //
		Interference           string `json:"interference,omitempty"`           //
		InterferenceScore      int    `json:"interferenceScore,omitempty"`      //
		Location               string `json:"location,omitempty"`               //
		MacAddress             string `json:"macAddress,omitempty"`             //
		ManagementIPAddr       string `json:"managementIpAddr,omitempty"`       //
		Memory                 string `json:"memory,omitempty"`                 //
		MemoryScore            int    `json:"memoryScore,omitempty"`            //
		Noise                  string `json:"noise,omitempty"`                  //
		NoiseScore             int    `json:"noiseScore,omitempty"`             //
		NwDeviceFamily         string `json:"nwDeviceFamily,omitempty"`         //
		NwDeviceID             string `json:"nwDeviceId,omitempty"`             //
		NwDeviceName           string `json:"nwDeviceName,omitempty"`           //
		NwDeviceRole           string `json:"nwDeviceRole,omitempty"`           //
		NwDeviceType           string `json:"nwDeviceType,omitempty"`           //
		OsType                 string `json:"osType,omitempty"`                 //
		OverallHealth          int    `json:"overallHealth,omitempty"`          //
		PacketPool             string `json:"packetPool,omitempty"`             //
		PacketPoolScore        int    `json:"packetPoolScore,omitempty"`        //
		PlatformID             string `json:"platformId,omitempty"`             //
		RedundancyMode         string `json:"redundancyMode,omitempty"`         //
		RedundancyPeerState    string `json:"redundancyPeerState,omitempty"`    //
		RedundancyState        string `json:"redundancyState,omitempty"`        //
		RedundancyUnit         string `json:"redundancyUnit,omitempty"`         //
		SoftwareVersion        string `json:"softwareVersion,omitempty"`        //
		Timestamp              string `json:"timestamp,omitempty"`              //
		Utilization            string `json:"utilization,omitempty"`            //
		UtilizationScore       int    `json:"utilizationScore,omitempty"`       //
		Wqe                    string `json:"wqe,omitempty"`                    //
		WqeScore               int    `json:"wqeScore,omitempty"`               //
	} `json:"response,omitempty"` //
}

// GetDeviceEnrichmentDetailsResponse is the GetDeviceEnrichmentDetailsResponse definition
type GetDeviceEnrichmentDetailsResponse struct {
	DeviceDetails struct {
		ApManagerInterfaceIP  string `json:"apManagerInterfaceIp,omitempty"`  //
		AssociatedWlcIP       string `json:"associatedWlcIp,omitempty"`       //
		BootDateTime          string `json:"bootDateTime,omitempty"`          //
		CollectionInterval    string `json:"collectionInterval,omitempty"`    //
		CollectionStatus      string `json:"collectionStatus,omitempty"`      //
		ErrorCode             string `json:"errorCode,omitempty"`             //
		ErrorDescription      string `json:"errorDescription,omitempty"`      //
		Family                string `json:"family,omitempty"`                //
		Hostname              string `json:"hostname,omitempty"`              //
		ID                    string `json:"id,omitempty"`                    //
		InstanceUUID          string `json:"instanceUuid,omitempty"`          //
		InterfaceCount        string `json:"interfaceCount,omitempty"`        //
		InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` //
		LastUpdateTime        int    `json:"lastUpdateTime,omitempty"`        //
		LastUpdated           string `json:"lastUpdated,omitempty"`           //
		LineCardCount         string `json:"lineCardCount,omitempty"`         //
		LineCardID            string `json:"lineCardId,omitempty"`            //
		Location              string `json:"location,omitempty"`              //
		LocationName          string `json:"locationName,omitempty"`          //
		MacAddress            string `json:"macAddress,omitempty"`            //
		ManagementIPAddress   string `json:"managementIpAddress,omitempty"`   //
		MemorySize            string `json:"memorySize,omitempty"`            //
		NeighborTopology      []struct {
			Links []struct {
				ID              string   `json:"id,omitempty"`              //
				Label           []string `json:"label,omitempty"`           //
				LinkStatus      string   `json:"linkStatus,omitempty"`      //
				PortUtilization string   `json:"portUtilization,omitempty"` //
				Source          string   `json:"source,omitempty"`          //
				Target          string   `json:"target,omitempty"`          //
			} `json:"links,omitempty"` //
			Nodes []struct {
				CLIents         string `json:"clients,omitempty"`         //
				ConnectedDevice string `json:"connectedDevice,omitempty"` //
				Count           string `json:"count,omitempty"`           //
				Description     string `json:"description,omitempty"`     //
				DeviceType      string `json:"deviceType,omitempty"`      //
				FabricGroup     string `json:"fabricGroup,omitempty"`     //
				Family          string `json:"family,omitempty"`          //
				HealthScore     int    `json:"healthScore,omitempty"`     //
				ID              string `json:"id,omitempty"`              //
				IP              string `json:"ip,omitempty"`              //
				Level           int    `json:"level,omitempty"`           //
				Name            string `json:"name,omitempty"`            //
				NodeType        string `json:"nodeType,omitempty"`        //
				PlatformID      string `json:"platformId,omitempty"`      //
				RadioFrequency  string `json:"radioFrequency,omitempty"`  //
				Role            string `json:"role,omitempty"`            //
				SoftwareVersion string `json:"softwareVersion,omitempty"` //
				UserID          string `json:"userId,omitempty"`          //
			} `json:"nodes,omitempty"` //
		} `json:"neighborTopology,omitempty"` //
		PlatformID                string `json:"platformId,omitempty"`                //
		ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` //
		ReachabilityStatus        string `json:"reachabilityStatus,omitempty"`        //
		Role                      string `json:"role,omitempty"`                      //
		RoleSource                string `json:"roleSource,omitempty"`                //
		SerialNumber              string `json:"serialNumber,omitempty"`              //
		Series                    string `json:"series,omitempty"`                    //
		SNMPContact               string `json:"snmpContact,omitempty"`               //
		SNMPLocation              string `json:"snmpLocation,omitempty"`              //
		SoftwareVersion           string `json:"softwareVersion,omitempty"`           //
		TagCount                  string `json:"tagCount,omitempty"`                  //
		TunnelUDPPort             string `json:"tunnelUdpPort,omitempty"`             //
		Type                      string `json:"type,omitempty"`                      //
		UpTime                    string `json:"upTime,omitempty"`                    //
		WaasDeviceMode            string `json:"waasDeviceMode,omitempty"`            //
	} `json:"deviceDetails,omitempty"` //
}

// GetDeviceInterfaceCountByDeviceIDResponse is the GetDeviceInterfaceCountByDeviceIdResponse definition
type GetDeviceInterfaceCountByDeviceIDResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDeviceInterfaceCountResponse is the GetDeviceInterfaceCountResponse definition
type GetDeviceInterfaceCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDeviceInterfaceVLANsResponse is the GetDeviceInterfaceVLANsResponse definition
type GetDeviceInterfaceVLANsResponse struct {
	Response []struct {
		InterfaceName  string `json:"interfaceName,omitempty"`  //
		IPAddress      string `json:"ipAddress,omitempty"`      //
		Mask           int    `json:"mask,omitempty"`           //
		NetworkAddress string `json:"networkAddress,omitempty"` //
		NumberOfIPs    int    `json:"numberOfIPs,omitempty"`    //
		Prefix         string `json:"prefix,omitempty"`         //
		VLANNumber     int    `json:"vlanNumber,omitempty"`     //
		VLANType       string `json:"vlanType,omitempty"`       //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDeviceInterfacesBySpecifiedRangeResponse is the GetDeviceInterfacesBySpecifiedRangeResponse definition
type GetDeviceInterfacesBySpecifiedRangeResponse struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDeviceListResponse is the GetDeviceListResponse definition
type GetDeviceListResponse struct {
	Response []struct {
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
		LastUpdateTime            string `json:"lastUpdateTime,omitempty"`            //
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDeviceSummaryResponse is the GetDeviceSummaryResponse definition
type GetDeviceSummaryResponse struct {
	Response struct {
		ID         string `json:"id,omitempty"`         //
		Role       string `json:"role,omitempty"`       //
		RoleSource string `json:"roleSource,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetFunctionalCapabilityByIDResponse is the GetFunctionalCapabilityByIdResponse definition
type GetFunctionalCapabilityByIDResponse struct {
	Response struct {
		AttributeInfo   string `json:"attributeInfo,omitempty"` //
		FunctionDetails []struct {
			AttributeInfo string `json:"attributeInfo,omitempty"` //
			ID            string `json:"id,omitempty"`            //
			PropertyName  string `json:"propertyName,omitempty"`  //
			StringValue   string `json:"stringValue,omitempty"`   //
		} `json:"functionDetails,omitempty"` //
		FunctionName    string `json:"functionName,omitempty"`    //
		FunctionOpState string `json:"functionOpState,omitempty"` //
		ID              string `json:"id,omitempty"`              //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetFunctionalCapabilityForDevicesResponse is the GetFunctionalCapabilityForDevicesResponse definition
type GetFunctionalCapabilityForDevicesResponse struct {
	Response []struct {
		AttributeInfo        string `json:"attributeInfo,omitempty"` //
		DeviceID             string `json:"deviceId,omitempty"`      //
		FunctionalCapability []struct {
			AttributeInfo   string `json:"attributeInfo,omitempty"` //
			FunctionDetails []struct {
				AttributeInfo string `json:"attributeInfo,omitempty"` //
				ID            string `json:"id,omitempty"`            //
				PropertyName  string `json:"propertyName,omitempty"`  //
				StringValue   string `json:"stringValue,omitempty"`   //
			} `json:"functionDetails,omitempty"` //
			FunctionName    string `json:"functionName,omitempty"`    //
			FunctionOpState string `json:"functionOpState,omitempty"` //
			ID              string `json:"id,omitempty"`              //
		} `json:"functionalCapability,omitempty"` //
		ID string `json:"id,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetISISInterfacesResponse is the GetISISInterfacesResponse definition
type GetISISInterfacesResponse struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetInterfaceByIPResponse is the GetInterfaceByIPResponse definition
type GetInterfaceByIPResponse struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetInterfaceByIDResponse is the GetInterfaceByIdResponse definition
type GetInterfaceByIDResponse struct {
	Response struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetInterfaceDetailsByDeviceIDAndInterfaceNameResponse is the GetInterfaceDetailsByDeviceIdAndInterfaceNameResponse definition
type GetInterfaceDetailsByDeviceIDAndInterfaceNameResponse struct {
	Response struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetInterfaceInfoByIDResponse is the GetInterfaceInfoByIdResponse definition
type GetInterfaceInfoByIDResponse struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetModuleCountResponse is the GetModuleCountResponse definition
type GetModuleCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetModuleInfoByIDResponse is the GetModuleInfoByIdResponse definition
type GetModuleInfoByIDResponse struct {
	Response struct {
		AssemblyNumber           string `json:"assemblyNumber,omitempty"`           //
		AssemblyRevision         string `json:"assemblyRevision,omitempty"`         //
		AttributeInfo            string `json:"attributeInfo,omitempty"`            //
		ContainmentEntity        string `json:"containmentEntity,omitempty"`        //
		Description              string `json:"description,omitempty"`              //
		EntityPhysicalIndex      string `json:"entityPhysicalIndex,omitempty"`      //
		ID                       string `json:"id,omitempty"`                       //
		IsFieldReplaceable       string `json:"isFieldReplaceable,omitempty"`       //
		IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` //
		Manufacturer             string `json:"manufacturer,omitempty"`             //
		ModuleIndex              int    `json:"moduleIndex,omitempty"`              //
		Name                     string `json:"name,omitempty"`                     //
		OperationalStateCode     string `json:"operationalStateCode,omitempty"`     //
		PartNumber               string `json:"partNumber,omitempty"`               //
		SerialNumber             string `json:"serialNumber,omitempty"`             //
		VendorEquipmentType      string `json:"vendorEquipmentType,omitempty"`      //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetModulesResponse is the GetModulesResponse definition
type GetModulesResponse struct {
	Response []struct {
		AssemblyNumber           string `json:"assemblyNumber,omitempty"`           //
		AssemblyRevision         string `json:"assemblyRevision,omitempty"`         //
		AttributeInfo            string `json:"attributeInfo,omitempty"`            //
		ContainmentEntity        string `json:"containmentEntity,omitempty"`        //
		Description              string `json:"description,omitempty"`              //
		EntityPhysicalIndex      string `json:"entityPhysicalIndex,omitempty"`      //
		ID                       string `json:"id,omitempty"`                       //
		IsFieldReplaceable       string `json:"isFieldReplaceable,omitempty"`       //
		IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` //
		Manufacturer             string `json:"manufacturer,omitempty"`             //
		ModuleIndex              int    `json:"moduleIndex,omitempty"`              //
		Name                     string `json:"name,omitempty"`                     //
		OperationalStateCode     string `json:"operationalStateCode,omitempty"`     //
		PartNumber               string `json:"partNumber,omitempty"`               //
		SerialNumber             string `json:"serialNumber,omitempty"`             //
		VendorEquipmentType      string `json:"vendorEquipmentType,omitempty"`      //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetNetworkDeviceByIPResponse is the GetNetworkDeviceByIPResponse definition
type GetNetworkDeviceByIPResponse struct {
	Response struct {
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
		LastUpdateTime            string `json:"lastUpdateTime,omitempty"`            //
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetNetworkDeviceByPaginationRangeResponse is the GetNetworkDeviceByPaginationRangeResponse definition
type GetNetworkDeviceByPaginationRangeResponse struct {
	Response []struct {
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
		LastUpdateTime            string `json:"lastUpdateTime,omitempty"`            //
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetOSPFInterfacesResponse is the GetOSPFInterfacesResponse definition
type GetOSPFInterfacesResponse struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetOrganizationListForMerakiResponse is the GetOrganizationListForMerakiResponse definition
type GetOrganizationListForMerakiResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetPollingIntervalByIDResponse is the GetPollingIntervalByIdResponse definition
type GetPollingIntervalByIDResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetPollingIntervalForAllDevicesResponse is the GetPollingIntervalForAllDevicesResponse definition
type GetPollingIntervalForAllDevicesResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetWirelessLanControllerDetailsByIDResponse is the GetWirelessLanControllerDetailsByIdResponse definition
type GetWirelessLanControllerDetailsByIDResponse struct {
	Response struct {
		AdminEnabledPorts        []int  `json:"adminEnabledPorts,omitempty"`        //
		ApGroupName              string `json:"apGroupName,omitempty"`              //
		DeviceID                 string `json:"deviceId,omitempty"`                 //
		EthMacAddress            string `json:"ethMacAddress,omitempty"`            //
		FlexGroupName            string `json:"flexGroupName,omitempty"`            //
		ID                       string `json:"id,omitempty"`                       //
		InstanceTenantID         string `json:"instanceTenantId,omitempty"`         //
		InstanceUUID             string `json:"instanceUuid,omitempty"`             //
		LagModeEnabled           bool   `json:"lagModeEnabled,omitempty"`           //
		NetconfEnabled           bool   `json:"netconfEnabled,omitempty"`           //
		WirelessLicenseInfo      string `json:"wirelessLicenseInfo,omitempty"`      //
		WirelessPackageInstalled bool   `json:"wirelessPackageInstalled,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// RegisterDeviceForWSAResponse is the RegisterDeviceForWSAResponse definition
type RegisterDeviceForWSAResponse struct {
	Response struct {
		MacAddress   string `json:"macAddress,omitempty"`   //
		ModelNumber  string `json:"modelNumber,omitempty"`  //
		Name         string `json:"name,omitempty"`         //
		SerialNumber string `json:"serialNumber,omitempty"` //
		TenantID     string `json:"tenantId,omitempty"`     //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// SyncDevicesResponse is the SyncDevicesResponse definition
type SyncDevicesResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// SyncNetworkDevicesResponse is the SyncNetworkDevicesResponse definition
type SyncNetworkDevicesResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateDeviceRoleResponse is the UpdateDeviceRoleResponse definition
type UpdateDeviceRoleResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// AddDevice addDevice
/* Adds the device with given credential
 */
func (s *DevicesService) AddDevice(addDeviceRequest *AddDeviceRequest) (*AddDeviceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device"

	response, err := RestyClient.R().
		SetBody(addDeviceRequest).
		SetResult(&AddDeviceResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*AddDeviceResponse)
	return result, response, err
}

// DeleteDeviceByIDQueryParams defines the query parameters for this request
type DeleteDeviceByIDQueryParams struct {
	IsForceDelete bool `url:"isForceDelete,omitempty"` // isForceDelete
}

// DeleteDeviceByID deleteDeviceById
/* Deletes the network device for the given Id
@param id Device ID
@param isForceDelete isForceDelete
*/
func (s *DevicesService) DeleteDeviceByID(id string, deleteDeviceByIDQueryParams *DeleteDeviceByIDQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(deleteDeviceByIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DevicesQueryParams defines the query parameters for this request
type DevicesQueryParams struct {
	DeviceRole string `url:"deviceRole,omitempty"` // The device role (One of CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, AP)
	SiteID     string `url:"siteId,omitempty"`     // Assurance site UUID value
	Health     string `url:"health,omitempty"`     // The device overall health (One of POOR, FAIR, GOOD)
	StartTime  int    `url:"startTime,omitempty"`  // UTC epoch time in milliseconds
	EndTime    int    `url:"endTime,omitempty"`    // UTC epoch time in miliseconds
	Limit      int    `url:"limit,omitempty"`      // Max number of device entries in the response (default to 50.  Max at 1000)
	Offset     int    `url:"offset,omitempty"`     // The offset of the first device in the returned data
}

// Devices devices
/* Intent API for accessing DNA Assurance Device object for generating reports, creating dashboards or creating additional value added services.
@param deviceRole The device role (One of CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, AP)
@param siteID Assurance site UUID value
@param health The device overall health (One of POOR, FAIR, GOOD)
@param startTime UTC epoch time in milliseconds
@param endTime UTC epoch time in miliseconds
@param limit Max number of device entries in the response (default to 50.  Max at 1000)
@param offset The offset of the first device in the returned data
*/
func (s *DevicesService) Devices(devicesQueryParams *DevicesQueryParams) (*DevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-health"

	queryString, _ := query.Values(devicesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DevicesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*DevicesResponse)
	return result, response, err
}

// ExportDeviceList exportDeviceList
/* Exports the selected network device to a file
 */
func (s *DevicesService) ExportDeviceList(exportDeviceListRequest *ExportDeviceListRequest) (*ExportDeviceListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/file"

	response, err := RestyClient.R().
		SetBody(exportDeviceListRequest).
		SetResult(&ExportDeviceListResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*ExportDeviceListResponse)
	return result, response, err
}

// GetAllInterfacesQueryParams defines the query parameters for this request
type GetAllInterfacesQueryParams struct {
	Offset int `url:"offset,omitempty"` // offset
	Limit  int `url:"limit,omitempty"`  // limit
}

// GetAllInterfaces getAllInterfaces
/* Returns all available interfaces. This endpoint can return a maximum of 500 interfaces
@param offset offset
@param limit limit
*/
func (s *DevicesService) GetAllInterfaces(getAllInterfacesQueryParams *GetAllInterfacesQueryParams) (*GetAllInterfacesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface"

	queryString, _ := query.Values(getAllInterfacesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetAllInterfacesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetAllInterfacesResponse)
	return result, response, err
}

// GetDeviceByID getDeviceByID
/* Returns the network device details for the given device ID
@param id Device ID
*/
func (s *DevicesService) GetDeviceByID(id string) (*DevicesGetDeviceByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DevicesGetDeviceByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*DevicesGetDeviceByIDResponse)
	return result, response, err
}

// GetDeviceBySerialNumber getDeviceBySerialNumber
/* Returns the network device with given serial number
@param serialNumber Device serial number
*/
func (s *DevicesService) GetDeviceBySerialNumber(serialNumber string) (*GetDeviceBySerialNumberResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{"+"serialNumber"+"}", fmt.Sprintf("%v", serialNumber), -1)

	response, err := RestyClient.R().
		SetResult(&GetDeviceBySerialNumberResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceBySerialNumberResponse)
	return result, response, err
}

// GetDeviceConfigByID getDeviceConfigById
/* Returns the device config by specified device ID
@param networkDeviceID networkDeviceId
*/
func (s *DevicesService) GetDeviceConfigByID(networkDeviceID string) (*GetDeviceConfigByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{networkDeviceId}/config"
	path = strings.Replace(path, "{"+"networkDeviceId"+"}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := RestyClient.R().
		SetResult(&GetDeviceConfigByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceConfigByIDResponse)
	return result, response, err
}

// GetDeviceConfigCount getDeviceConfigCount
/* Returns the count of device configs
 */
func (s *DevicesService) GetDeviceConfigCount() (*GetDeviceConfigCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/config/count"

	response, err := RestyClient.R().
		SetResult(&GetDeviceConfigCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceConfigCountResponse)
	return result, response, err
}

// GetDeviceConfigForAllDevices getDeviceConfigForAllDevices
/* Returns the config for all devices
 */
func (s *DevicesService) GetDeviceConfigForAllDevices() (*GetDeviceConfigForAllDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/config"

	response, err := RestyClient.R().
		SetResult(&GetDeviceConfigForAllDevicesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceConfigForAllDevicesResponse)
	return result, response, err
}

// GetDeviceCount getDeviceCount
/* Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and location name
 */
func (s *DevicesService) GetDeviceCount() (*GetDeviceCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/count"

	response, err := RestyClient.R().
		SetResult(&GetDeviceCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceCountResponse)
	return result, response, err
}

// GetDeviceDetailQueryParams defines the query parameters for this request
type GetDeviceDetailQueryParams struct {
	Timestamp  string `url:"timestamp,omitempty"`  // Epoch time(in milliseconds) when the device data is required
	SearchBy   string `url:"searchBy,omitempty"`   // MAC Address or Device Name value or UUID of the network device
	IDentifier string `url:"identifier,omitempty"` // One of keywords : macAddress or uuid or nwDeviceName
}

// GetDeviceDetail getDeviceDetail
/* Returns detailed Network Device information retrieved by Mac Address, Device Name or UUID for any given point of time.
@param timestamp Epoch time(in milliseconds) when the device data is required
@param searchBy MAC Address or Device Name value or UUID of the network device
@param identifier One of keywords : macAddress or uuid or nwDeviceName
*/
func (s *DevicesService) GetDeviceDetail(getDeviceDetailQueryParams *GetDeviceDetailQueryParams) (*GetDeviceDetailResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-detail"

	queryString, _ := query.Values(getDeviceDetailQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceDetailResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceDetailResponse)
	return result, response, err
}

// GetDeviceEnrichmentDetails getDeviceEnrichmentDetails
/* Enriches a given network device context (device id or device Mac Address or device management IP address) with details about the device and neighbor topology
@param entity_type Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
@param entity_value Contains the actual value for the entity type that has been defined
*/
func (s *DevicesService) GetDeviceEnrichmentDetails() (*GetDeviceEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetDeviceEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceEnrichmentDetailsResponse)
	return result, response, err
}

// GetDeviceInterfaceCount getDeviceInterfaceCount
/* Returns the count of interfaces for all devices
 */
func (s *DevicesService) GetDeviceInterfaceCount() (*GetDeviceInterfaceCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/count"

	response, err := RestyClient.R().
		SetResult(&GetDeviceInterfaceCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceInterfaceCountResponse)
	return result, response, err
}

// GetDeviceInterfaceCountByDeviceID getDeviceInterfaceCountByDeviceId
/* Returns the interface count for the given device
@param deviceID Device ID
*/
func (s *DevicesService) GetDeviceInterfaceCountByDeviceID(deviceID string) (*GetDeviceInterfaceCountByDeviceIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/count"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)

	response, err := RestyClient.R().
		SetResult(&GetDeviceInterfaceCountByDeviceIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceInterfaceCountByDeviceIDResponse)
	return result, response, err
}

// GetDeviceInterfaceVLANsQueryParams defines the query parameters for this request
type GetDeviceInterfaceVLANsQueryParams struct {
	InterfaceType string `url:"interfaceType,omitempty"` // Vlan assocaited with sub-interface
}

// GetDeviceInterfaceVLANs getDeviceInterfaceVLANs
/* Returns Device Interface VLANs
@param id deviceUUID
@param interfaceType Vlan assocaited with sub-interface
*/
func (s *DevicesService) GetDeviceInterfaceVLANs(id string, getDeviceInterfaceVLANsQueryParams *GetDeviceInterfaceVLANsQueryParams) (*GetDeviceInterfaceVLANsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/vlan"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getDeviceInterfaceVLANsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceInterfaceVLANsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceInterfaceVLANsResponse)
	return result, response, err
}

// GetDeviceInterfacesBySpecifiedRange getDeviceInterfacesBySpecifiedRange
/* Returns the list of interfaces for the device for the specified range
@param deviceID Device ID
@param startIndex Start index
@param recordsToReturn Number of records to return
*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRange(deviceID string, startIndex int, recordsToReturn int) (*GetDeviceInterfacesBySpecifiedRangeResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := RestyClient.R().
		SetResult(&GetDeviceInterfacesBySpecifiedRangeResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceInterfacesBySpecifiedRangeResponse)
	return result, response, err
}

// GetDeviceListQueryParams defines the query parameters for this request
type GetDeviceListQueryParams struct {
	Hostname                []string `url:"hostname,omitempty"`                   // hostname
	ManagementIPAddress     []string `url:"managementIpAddress,omitempty"`        // managementIpAddress
	MacAddress              []string `url:"macAddress,omitempty"`                 // macAddress
	LocationName            []string `url:"locationName,omitempty"`               // locationName
	SerialNumber            []string `url:"serialNumber,omitempty"`               // serialNumber
	Location                []string `url:"location,omitempty"`                   // location
	Family                  []string `url:"family,omitempty"`                     // family
	Type                    []string `url:"type,omitempty"`                       // type
	Series                  []string `url:"series,omitempty"`                     // series
	CollectionStatus        []string `url:"collectionStatus,omitempty"`           // collectionStatus
	CollectionInterval      []string `url:"collectionInterval,omitempty"`         // collectionInterval
	NotSyncedForMinutes     []string `url:"notSyncedForMinutes,omitempty"`        // notSyncedForMinutes
	ErrorCode               []string `url:"errorCode,omitempty"`                  // errorCode
	ErrorDescription        []string `url:"errorDescription,omitempty"`           // errorDescription
	SoftwareVersion         []string `url:"softwareVersion,omitempty"`            // softwareVersion
	SoftwareType            []string `url:"softwareType,omitempty"`               // softwareType
	PlatformID              []string `url:"platformId,omitempty"`                 // platformId
	Role                    []string `url:"role,omitempty"`                       // role
	ReachabilityStatus      []string `url:"reachabilityStatus,omitempty"`         // reachabilityStatus
	UpTime                  []string `url:"upTime,omitempty"`                     // upTime
	AssociatedWlcIP         []string `url:"associatedWlcIp,omitempty"`            // associatedWlcIp
	Licensame               []string `url:"license.name,omitempty"`               // licenseName
	Licensype               []string `url:"license.type,omitempty"`               // licenseType
	Licenstatus             []string `url:"license.status,omitempty"`             // licenseStatus
	Modulame                []string `url:"module+name,omitempty"`                // moduleName
	Modulqupimenttype       []string `url:"module+equpimenttype,omitempty"`       // moduleEqupimentType
	Modulervicestate        []string `url:"module+servicestate,omitempty"`        // moduleServiceState
	Modulendorequipmenttype []string `url:"module+vendorequipmenttype,omitempty"` // moduleVendorEquipmentType
	Modulartnumber          []string `url:"module+partnumber,omitempty"`          // modulePartNumber
	Modulperationstatecode  []string `url:"module+operationstatecode,omitempty"`  // moduleOperationStateCode
	ID                      string   `url:"id,omitempty"`                         // Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
}

// GetDeviceList getDeviceList
/* Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc. You can use the .* in any value to conduct a wildcard search.
For example, to find all hostnames beginning with myhost in the IP address range 192.25.18.n, issue the following request:
GET /dna/intent/api/v1/network-device?hostname=myhost.*&managementIpAddress=192.25.18..*

If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and ignores the other request parameters.
@param hostname hostname
@param managementIPAddress managementIpAddress
@param macAddress macAddress
@param locationName locationName
@param serialNumber serialNumber
@param location location
@param family family
@param type type
@param series series
@param collectionStatus collectionStatus
@param collectionInterval collectionInterval
@param notSyncedForMinutes notSyncedForMinutes
@param errorCode errorCode
@param errorDescription errorDescription
@param softwareVersion softwareVersion
@param softwareType softwareType
@param platformID platformId
@param role role
@param reachabilityStatus reachabilityStatus
@param upTime upTime
@param associatedWlcIP associatedWlcIp
@param license.name licenseName
@param license.type licenseType
@param license.status licenseStatus
@param module+name moduleName
@param module+equpimenttype moduleEqupimentType
@param module+servicestate moduleServiceState
@param module+vendorequipmenttype moduleVendorEquipmentType
@param module+partnumber modulePartNumber
@param module+operationstatecode moduleOperationStateCode
@param id Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
*/
func (s *DevicesService) GetDeviceList(getDeviceListQueryParams *GetDeviceListQueryParams) (*GetDeviceListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device"

	queryString, _ := query.Values(getDeviceListQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceListResponse)
	return result, response, err
}

// GetDeviceSummary getDeviceSummary
/* Returns brief summary of device info such as hostname, management IP address for the given device Id
@param id Device ID
*/
func (s *DevicesService) GetDeviceSummary(id string) (*GetDeviceSummaryResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/brief"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetDeviceSummaryResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDeviceSummaryResponse)
	return result, response, err
}

// GetFunctionalCapabilityByID getFunctionalCapabilityById
/* Returns functional capability with given Id
@param id Functional Capability UUID
*/
func (s *DevicesService) GetFunctionalCapabilityByID(id string) (*GetFunctionalCapabilityByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/functional-capability/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetFunctionalCapabilityByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetFunctionalCapabilityByIDResponse)
	return result, response, err
}

// GetFunctionalCapabilityForDevicesQueryParams defines the query parameters for this request
type GetFunctionalCapabilityForDevicesQueryParams struct {
	DeviceID     string   `url:"deviceId,omitempty"`     // Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
	FunctionName []string `url:"functionName,omitempty"` // functionName
}

// GetFunctionalCapabilityForDevices getFunctionalCapabilityForDevices
/* Returns the functional-capability for given devices
@param deviceID Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
@param functionName functionName
*/
func (s *DevicesService) GetFunctionalCapabilityForDevices(getFunctionalCapabilityForDevicesQueryParams *GetFunctionalCapabilityForDevicesQueryParams) (*GetFunctionalCapabilityForDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/functional-capability"

	queryString, _ := query.Values(getFunctionalCapabilityForDevicesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetFunctionalCapabilityForDevicesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetFunctionalCapabilityForDevicesResponse)
	return result, response, err
}

// GetISISInterfaces getISISInterfaces
/* Returns the interfaces that has ISIS enabled
 */
func (s *DevicesService) GetISISInterfaces() (*GetISISInterfacesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/isis"

	response, err := RestyClient.R().
		SetResult(&GetISISInterfacesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetISISInterfacesResponse)
	return result, response, err
}

// GetInterfaceByIP getInterfaceByIP
/* Returns list of interfaces for specified device management IP address
@param ipAddress IP address of the interface
*/
func (s *DevicesService) GetInterfaceByIP(ipAddress string) (*GetInterfaceByIPResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := RestyClient.R().
		SetResult(&GetInterfaceByIPResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetInterfaceByIPResponse)
	return result, response, err
}

// GetInterfaceByID getInterfaceById
/* Returns the interface for the given interface ID
@param id Interface ID
*/
func (s *DevicesService) GetInterfaceByID(id string) (*GetInterfaceByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetInterfaceByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetInterfaceByIDResponse)
	return result, response, err
}

// GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams defines the query parameters for this request
type GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams struct {
	Name string `url:"name,omitempty"` // Interface name
}

// GetInterfaceDetailsByDeviceIDAndInterfaceName getInterfaceDetailsByDeviceIdAndInterfaceName
/* Returns interface by specified device Id and interface name
@param deviceID Device ID
@param name Interface name
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID string, getInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams) (*GetInterfaceDetailsByDeviceIDAndInterfaceNameResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/interface-name"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(getInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetInterfaceDetailsByDeviceIDAndInterfaceNameResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetInterfaceDetailsByDeviceIDAndInterfaceNameResponse)
	return result, response, err
}

// GetInterfaceInfoByID getInterfaceInfoById
/* Returns list of interfaces by specified device
@param deviceID Device ID
*/
func (s *DevicesService) GetInterfaceInfoByID(deviceID string) (*GetInterfaceInfoByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceId}"
	path = strings.Replace(path, "{"+"deviceId"+"}", fmt.Sprintf("%v", deviceID), -1)

	response, err := RestyClient.R().
		SetResult(&GetInterfaceInfoByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetInterfaceInfoByIDResponse)
	return result, response, err
}

// GetModuleCountQueryParams defines the query parameters for this request
type GetModuleCountQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 // deviceId
	NameList                 []string `url:"nameList,omitempty"`                 // nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  // vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           // partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` // operationalStateCodeList
}

// GetModuleCount getModuleCount
/* Returns Module Count
@param deviceID deviceId
@param nameList nameList
@param vendorEquipmentTypeList vendorEquipmentTypeList
@param partNumberList partNumberList
@param operationalStateCodeList operationalStateCodeList
*/
func (s *DevicesService) GetModuleCount(getModuleCountQueryParams *GetModuleCountQueryParams) (*GetModuleCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/module/count"

	queryString, _ := query.Values(getModuleCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetModuleCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetModuleCountResponse)
	return result, response, err
}

// GetModuleInfoByID getModuleInfoById
/* Returns Module info by id
@param id id
*/
func (s *DevicesService) GetModuleInfoByID(id string) (*GetModuleInfoByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/module/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetModuleInfoByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetModuleInfoByIDResponse)
	return result, response, err
}

// GetModulesQueryParams defines the query parameters for this request
type GetModulesQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 // deviceId
	Limit                    string   `url:"limit,omitempty"`                    // limit
	Offset                   string   `url:"offset,omitempty"`                   // offset
	NameList                 []string `url:"nameList,omitempty"`                 // nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  // vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           // partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` // operationalStateCodeList
}

// GetModules getModules
/* Returns modules by specified device id
@param deviceID deviceId
@param limit limit
@param offset offset
@param nameList nameList
@param vendorEquipmentTypeList vendorEquipmentTypeList
@param partNumberList partNumberList
@param operationalStateCodeList operationalStateCodeList
*/
func (s *DevicesService) GetModules(getModulesQueryParams *GetModulesQueryParams) (*GetModulesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/module"

	queryString, _ := query.Values(getModulesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetModulesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetModulesResponse)
	return result, response, err
}

// GetNetworkDeviceByIP getNetworkDeviceByIP
/* Returns the network device by specified IP address
@param ipAddress Device IP address
*/
func (s *DevicesService) GetNetworkDeviceByIP(ipAddress string) (*GetNetworkDeviceByIPResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := RestyClient.R().
		SetResult(&GetNetworkDeviceByIPResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetNetworkDeviceByIPResponse)
	return result, response, err
}

// GetNetworkDeviceByPaginationRange getNetworkDeviceByPaginationRange
/* Returns the list of network devices for the given pagination range
@param startIndex Start index
@param recordsToReturn Number of records to return
*/
func (s *DevicesService) GetNetworkDeviceByPaginationRange(startIndex int, recordsToReturn int) (*GetNetworkDeviceByPaginationRangeResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := RestyClient.R().
		SetResult(&GetNetworkDeviceByPaginationRangeResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetNetworkDeviceByPaginationRangeResponse)
	return result, response, err
}

// GetOSPFInterfaces getOSPFInterfaces
/* Returns the interfaces that has OSPF enabled
 */
func (s *DevicesService) GetOSPFInterfaces() (*GetOSPFInterfacesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/ospf"

	response, err := RestyClient.R().
		SetResult(&GetOSPFInterfacesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetOSPFInterfacesResponse)
	return result, response, err
}

// GetOrganizationListForMeraki getOrganizationListForMeraki
/* Returns list of organizations for meraki dashboard
@param id id
*/
func (s *DevicesService) GetOrganizationListForMeraki(id string) (*GetOrganizationListForMerakiResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/meraki-organization"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetOrganizationListForMerakiResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetOrganizationListForMerakiResponse)
	return result, response, err
}

// GetPollingIntervalByID getPollingIntervalById
/* Returns polling interval by device id
@param id Device ID
*/
func (s *DevicesService) GetPollingIntervalByID(id string) (*GetPollingIntervalByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/collection-schedule"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetPollingIntervalByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetPollingIntervalByIDResponse)
	return result, response, err
}

// GetPollingIntervalForAllDevices getPollingIntervalForAllDevices
/* Returns polling interval of all devices
 */
func (s *DevicesService) GetPollingIntervalForAllDevices() (*GetPollingIntervalForAllDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/collection-schedule/global"

	response, err := RestyClient.R().
		SetResult(&GetPollingIntervalForAllDevicesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetPollingIntervalForAllDevicesResponse)
	return result, response, err
}

// GetWirelessLanControllerDetailsByID getWirelessLanControllerDetailsById
/* Returns the wireless lan controller info with given device ID
@param id Device ID
*/
func (s *DevicesService) GetWirelessLanControllerDetailsByID(id string) (*GetWirelessLanControllerDetailsByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/wireless-info"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetWirelessLanControllerDetailsByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetWirelessLanControllerDetailsByIDResponse)
	return result, response, err
}

// RegisterDeviceForWSAQueryParams defines the query parameters for this request
type RegisterDeviceForWSAQueryParams struct {
	SerialNumber string `url:"serialNumber,omitempty"` // Serial number of the device
	Macaddress   string `url:"macaddress,omitempty"`   // Mac addres of the device
}

// RegisterDeviceForWSA registerDeviceForWSA
/* Registers a device for WSA notification
@param serialNumber Serial number of the device
@param macaddress Mac addres of the device
*/
func (s *DevicesService) RegisterDeviceForWSA(registerDeviceForWSAQueryParams *RegisterDeviceForWSAQueryParams) (*RegisterDeviceForWSAResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/tenantinfo/macaddress"

	queryString, _ := query.Values(registerDeviceForWSAQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&RegisterDeviceForWSAResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*RegisterDeviceForWSAResponse)
	return result, response, err
}

// RetrievesAllNetworkDevicesQueryParams defines the query parameters for this request
type RetrievesAllNetworkDevicesQueryParams struct {
	VrfName                   string `url:"vrfName,omitempty"`                   // vrfName
	ManagementIPAddress       string `url:"managementIpAddress,omitempty"`       // managementIpAddress
	Hostname                  string `url:"hostname,omitempty"`                  // hostname
	MacAddress                string `url:"macAddress,omitempty"`                // macAddress
	Family                    string `url:"family,omitempty"`                    // family
	CollectionStatus          string `url:"collectionStatus,omitempty"`          // collectionStatus
	CollectionInterval        string `url:"collectionInterval,omitempty"`        // collectionInterval
	SoftwareVersion           string `url:"softwareVersion,omitempty"`           // softwareVersion
	SoftwareType              string `url:"softwareType,omitempty"`              // softwareType
	ReachabilityStatus        string `url:"reachabilityStatus,omitempty"`        // reachabilityStatus
	ReachabilityFailureReason string `url:"reachabilityFailureReason,omitempty"` // reachabilityFailureReason
	ErrorCode                 string `url:"errorCode,omitempty"`                 // errorCode
	PlatformID                string `url:"platformId,omitempty"`                // platformId
	Series                    string `url:"series,omitempty"`                    // series
	Type                      string `url:"type,omitempty"`                      // type
	SerialNumber              string `url:"serialNumber,omitempty"`              // serialNumber
	UpTime                    string `url:"upTime,omitempty"`                    // upTime
	Role                      string `url:"role,omitempty"`                      // role
	RoleSource                string `url:"roleSource,omitempty"`                // roleSource
	AssociatedWlcIP           string `url:"associatedWlcIp,omitempty"`           // associatedWlcIp
	Offset                    string `url:"offset,omitempty"`                    // offset
	Limit                     string `url:"limit,omitempty"`                     // limit
}

// RetrievesAllNetworkDevices retrievesAllNetworkDevices
/* Gets the list of first 500 network devices sorted lexicographically based on host name. It can be filtered using management IP address, mac address, hostname and location name. If id param is provided, it will be returning the list of network-devices for the given id's and other request params will be ignored. In case of autocomplete request, returns the list of specified attributes.
@param vrfName vrfName
@param managementIPAddress managementIpAddress
@param hostname hostname
@param macAddress macAddress
@param family family
@param collectionStatus collectionStatus
@param collectionInterval collectionInterval
@param softwareVersion softwareVersion
@param softwareType softwareType
@param reachabilityStatus reachabilityStatus
@param reachabilityFailureReason reachabilityFailureReason
@param errorCode errorCode
@param platformID platformId
@param series series
@param type type
@param serialNumber serialNumber
@param upTime upTime
@param role role
@param roleSource roleSource
@param associatedWlcIP associatedWlcIp
@param offset offset
@param limit limit
*/
func (s *DevicesService) RetrievesAllNetworkDevices(retrievesAllNetworkDevicesQueryParams *RetrievesAllNetworkDevicesQueryParams) (string, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/autocomplete"

	queryString, _ := query.Values(retrievesAllNetworkDevicesQueryParams)

	var operationResult string
	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&operationResult).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return "", nil, err
	}
	return operationResult, response, err
}

// SyncDevices syncDevices
/* Sync the devices provided as input
 */
func (s *DevicesService) SyncDevices(syncDevicesRequest *SyncDevicesRequest) (*SyncDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device"

	response, err := RestyClient.R().
		SetBody(syncDevicesRequest).
		SetResult(&SyncDevicesResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*SyncDevicesResponse)
	return result, response, err
}

// SyncNetworkDevicesQueryParams defines the query parameters for this request
type SyncNetworkDevicesQueryParams struct {
	ForceSync bool `url:"forceSync,omitempty"` // forceSync
}

// SyncNetworkDevices syncNetworkDevices
/* Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result can be seen in the child task of each device
@param forceSync forceSync
*/
func (s *DevicesService) SyncNetworkDevices(syncNetworkDevicesQueryParams *SyncNetworkDevicesQueryParams, syncNetworkDevicesRequest *SyncNetworkDevicesRequest) (*SyncNetworkDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/sync"

	queryString, _ := query.Values(syncNetworkDevicesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(syncNetworkDevicesRequest).
		SetResult(&SyncNetworkDevicesResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*SyncNetworkDevicesResponse)
	return result, response, err
}

// UpdateDeviceRole updateDeviceRole
/* Updates the role of the device as access, core, distribution, border router
 */
func (s *DevicesService) UpdateDeviceRole(updateDeviceRoleRequest *UpdateDeviceRoleRequest) (*UpdateDeviceRoleResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/brief"

	response, err := RestyClient.R().
		SetBody(updateDeviceRoleRequest).
		SetResult(&UpdateDeviceRoleResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateDeviceRoleResponse)
	return result, response, err
}
