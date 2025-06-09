package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LanAutomationService service

type LanAutomationLogQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //Starting index of the LAN Automation session. Minimum value is 1.
	Limit  float64 `url:"limit,omitempty"`  //Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
}
type LanAutomationLogsForIndividualDevicesQueryParams struct {
	LogLevel string `url:"logLevel,omitempty"` //Supported levels are ERROR, INFO, WARNING, TRACE, CONFIG and ALL. Specifying ALL will display device specific logs with the exception of CONFIG logs. In order to view CONFIG logs along with the remaining logs, please leave the query parameter blank.
}
type LanAutomationStatusQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //Starting index of the LAN Automation session. Minimum value is 1.
	Limit  float64 `url:"limit,omitempty"`  //Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
}
type LanAutomationDeviceUpdateQueryParams struct {
	Feature string `url:"feature,omitempty"` //Feature ID for the update. Supported feature IDs include: LOOPBACK0_IPADDRESS_UPDATE, HOSTNAME_UPDATE, LINK_ADD, and LINK_DELETE.
}
type GetPortChannelsQueryParams struct {
	Device1ManagementIPAddress string `url:"device1ManagementIPAddress,omitempty"` //The management IP address of the device1.
	Device1UUID                string `url:"device1Uuid,omitempty"`                //Unique identifier for the network device1
	Device2ManagementIPAddress string `url:"device2ManagementIPAddress,omitempty"` //The management IP address of the device2.
	Device2UUID                string `url:"device2Uuid,omitempty"`                //Unique identifier for the network device2
	Offset                     int    `url:"offset,omitempty"`                     //Starting record for pagination.
	Limit                      int    `url:"limit,omitempty"`                      //Maximum number of Port Channel to return.
}

type ResponseLanAutomationLanAutomationStart struct {
	Response *ResponseLanAutomationLanAutomationStartResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStartResponse struct {
	Message string `json:"message,omitempty"` // Status of the LAN Automation Start request
	ID      string `json:"id,omitempty"`      // LAN Automation Session Id
}
type ResponseLanAutomationLanAutomationSessionCount struct {
	Response *ResponseLanAutomationLanAutomationSessionCountResponse `json:"response,omitempty"` //
}
type ResponseLanAutomationLanAutomationSessionCountResponse struct {
	SessionCount string `json:"sessionCount,omitempty"` // Total number of sessions executed.
}
type ResponseLanAutomationLanAutomationLog struct {
	Response *[]ResponseLanAutomationLanAutomationLogResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogResponse struct {
	NwOrchID string                                                `json:"nwOrchId,omitempty"` // LAN Automation session identifier.
	Entry    *[]ResponseLanAutomationLanAutomationLogResponseEntry `json:"entry,omitempty"`    //
}
type ResponseLanAutomationLanAutomationLogResponseEntry struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO, WARNING, TRACE and CONFIG.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
	DeviceID  string `json:"deviceId,omitempty"`  // Device serial number for which the log message is associated.
}
type ResponseLanAutomationLanAutomationLogByID struct {
	Response *[]ResponseLanAutomationLanAutomationLogByIDResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogByIDResponse struct {
	NwOrchID string                                                    `json:"nwOrchId,omitempty"` // LAN Automation session identifier.
	Entry    *[]ResponseLanAutomationLanAutomationLogByIDResponseEntry `json:"entry,omitempty"`    //
}
type ResponseLanAutomationLanAutomationLogByIDResponseEntry struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO, WARNING, TRACE and CONFIG.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
	DeviceID  string `json:"deviceId,omitempty"`  // Device serial number for which the log message is associated.
}
type ResponseLanAutomationLanAutomationLogsForIndividualDevices struct {
	Response *[]ResponseLanAutomationLanAutomationLogsForIndividualDevicesResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogsForIndividualDevicesResponse struct {
	NwOrchID     string                                                                    `json:"nwOrchId,omitempty"`     // LAN Automation session identifier.
	Logs         *[]ResponseLanAutomationLanAutomationLogsForIndividualDevicesResponseLogs `json:"logs,omitempty"`         //
	SerialNumber string                                                                    `json:"serialNumber,omitempty"` // Device serial number for which the log messages are associated.
}
type ResponseLanAutomationLanAutomationLogsForIndividualDevicesResponseLogs struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO, WARNING, TRACE, CONFIG and ALL. Specifying ALL will display device specific logs with the exception of CONFIG logs. In order to view CONFIG logs along with the remaining logs, please leave the query parameter blank.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
}
type ResponseLanAutomationLanAutomationActiveSessions struct {
	Response *ResponseLanAutomationLanAutomationActiveSessionsResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationActiveSessionsResponse struct {
	MaxSupportedCount string   `json:"maxSupportedCount,omitempty"` // Maximum supported parallel sessions count
	ActiveSessions    string   `json:"activeSessions,omitempty"`    // Current active sessions count
	ActiveSessionIDs  []string `json:"activeSessionIds,omitempty"`  // List of Active LAN Automation IDs
}
type ResponseLanAutomationLanAutomationStatus struct {
	Response *[]ResponseLanAutomationLanAutomationStatusResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusResponse struct {
	ID                                string                                                                  `json:"id,omitempty"`                                // LAN Automation session id.
	DiscoveredDeviceSiteNameHierarchy string                                                                  `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                                  `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address.
	IPPools                           *[]ResponseLanAutomationLanAutomationStatusResponseIPPools              `json:"ipPools,omitempty"`                           //
	PrimaryDeviceInterfaceNames       []string                                                                `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	Status                            string                                                                  `json:"status,omitempty"`                            // Status of the LAN Automation session along with the number of discovered devices.
	Action                            string                                                                  `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation session.
	CreationTime                      string                                                                  `json:"creationTime,omitempty"`                      // LAN Automation session creation time.
	MulticastEnabled                  *bool                                                                   `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not.
	PeerDeviceManagmentIPAddress      string                                                                  `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address.
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
	RedistributeIsisToBgp             *bool                                                                   `json:"redistributeIsisToBgp,omitempty"`             // Shows whether advertise LAN Automation summary route into BGP is enabled or not.
	DiscoveryLevel                    *int                                                                    `json:"discoveryLevel,omitempty"`                    // Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
	DiscoveryTimeout                  *int                                                                    `json:"discoveryTimeout,omitempty"`                  // Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080].
	DiscoveryDevices                  *[]ResponseLanAutomationLanAutomationStatusResponseDiscoveryDevices     `json:"discoveryDevices,omitempty"`                  //
	HostNamePrefix                    string                                                                  `json:"hostNamePrefix,omitempty"`                    // Hostname prefix of the discovered devices
	HostNameFileID                    string                                                                  `json:"hostNameFileId,omitempty"`                    // File ID for the uploaded file in the nw_orch namespace
}
type ResponseLanAutomationLanAutomationStatusResponseIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device.
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device.
	State              string   `json:"state,omitempty"`              // State of the device (Added to inventory/Deleted from inventory).
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // List of IP address used by the device.
}
type ResponseLanAutomationLanAutomationStatusResponseDiscoveryDevices struct {
	DeviceSerialNumber        string `json:"deviceSerialNumber,omitempty"`        // Serial number of the device
	DeviceHostName            string `json:"deviceHostName,omitempty"`            // Hostname of the device
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Management IP Address of the device
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Site UUID of the device
	DeviceSiteNameHierarchy   string `json:"deviceSiteNameHierarchy,omitempty"`   // Site name hierarchy of the device
	IsDeviceDiscovered        *bool  `json:"isDeviceDiscovered,omitempty"`        // Discovery status of the device
	IsIPAllocated             *bool  `json:"isIPAllocated,omitempty"`             // IP Allocation status of the device
	IsIPAssigned              *bool  `json:"isIPAssigned,omitempty"`              // IP Assignment status of the device
	PnpDeviceID               string `json:"pnpDeviceId,omitempty"`               // PnP device ID
}
type ResponseLanAutomationLanAutomationStatusByID struct {
	Response *[]ResponseLanAutomationLanAutomationStatusByIDResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusByIDResponse struct {
	ID                                string                                                                      `json:"id,omitempty"`                                // LAN Automation session id.
	DiscoveredDeviceSiteNameHierarchy string                                                                      `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                                      `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address.
	IPPools                           *[]ResponseLanAutomationLanAutomationStatusByIDResponseIPPools              `json:"ipPools,omitempty"`                           //
	PrimaryDeviceInterfaceNames       []string                                                                    `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	Status                            string                                                                      `json:"status,omitempty"`                            // Status of the LAN Automation session along with the number of discovered devices.
	Action                            string                                                                      `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation session.
	CreationTime                      string                                                                      `json:"creationTime,omitempty"`                      // LAN Automation session creation time.
	MulticastEnabled                  *bool                                                                       `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not.
	PeerDeviceManagmentIPAddress      string                                                                      `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address.
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
	RedistributeIsisToBgp             *bool                                                                       `json:"redistributeIsisToBgp,omitempty"`             // Shows whether advertise LAN Automation summary route into BGP is enabled or not.
	DiscoveryLevel                    *int                                                                        `json:"discoveryLevel,omitempty"`                    // Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
	DiscoveryTimeout                  *int                                                                        `json:"discoveryTimeout,omitempty"`                  // Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080].
	DiscoveryDevices                  *[]ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveryDevices     `json:"discoveryDevices,omitempty"`                  //
	HostNamePrefix                    string                                                                      `json:"hostNamePrefix,omitempty"`                    // Hostname prefix of the discovered devices
	HostNameFileID                    string                                                                      `json:"hostNameFileId,omitempty"`                    // File ID for the uploaded file in the nw_orch namespace
}
type ResponseLanAutomationLanAutomationStatusByIDResponseIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device.
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device.
	State              string   `json:"state,omitempty"`              // State of the device (Added to inventory/Deleted from inventory).
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // List of IP address used by the device.
}
type ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveryDevices struct {
	DeviceSerialNumber        string `json:"deviceSerialNumber,omitempty"`        // Serial number of the device
	DeviceHostName            string `json:"deviceHostName,omitempty"`            // Hostname of the device
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Management IP Address of the device
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Site UUID of the device
	DeviceSiteNameHierarchy   string `json:"deviceSiteNameHierarchy,omitempty"`   // Site name hierarchy of the device
	IsDeviceDiscovered        *bool  `json:"isDeviceDiscovered,omitempty"`        // Discovery status of the device
	IsIPAllocated             *bool  `json:"isIPAllocated,omitempty"`             // IP Allocation status of the device
	IsIPAssigned              *bool  `json:"isIPAssigned,omitempty"`              // IP Assignment status of the device
	PnpDeviceID               string `json:"pnpDeviceId,omitempty"`               // PnP device ID
}
type ResponseLanAutomationLanAutomationDeviceUpdate struct {
	Response *ResponseLanAutomationLanAutomationDeviceUpdateResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationDeviceUpdateResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationLanAutomationStop struct {
	Response *ResponseLanAutomationLanAutomationStopResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStopResponse struct {
	ErrorCode string `json:"errorCode,omitempty"` // Error code value.
	Message   string `json:"message,omitempty"`   // Description of the error code.
	Detail    string `json:"detail,omitempty"`    // Detailed information of the error code.
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevices struct {
	Response *ResponseLanAutomationLanAutomationStopAndUpdateDevicesResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationGetPortChannels struct {
	Response *[]ResponseLanAutomationGetPortChannelsResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version number.
}
type ResponseLanAutomationGetPortChannelsResponse struct {
	ID                         string                                                            `json:"id,omitempty"`                         // Unique id of the port-channel link. Concatenation of both device's Port Channel UUID.
	Device1ManagementIPAddress string                                                            `json:"device1ManagementIPAddress,omitempty"` // IP address of device 1.
	Device1UUID                string                                                            `json:"device1Uuid,omitempty"`                // ID of device1.
	Device2ManagementIPAddress string                                                            `json:"device2ManagementIPAddress,omitempty"` // IP address of device 2.
	Device2UUID                string                                                            `json:"device2Uuid,omitempty"`                // ID of device2.
	Device1PortChannelUUID     string                                                            `json:"device1PortChannelUuid,omitempty"`     // ID of the port channel.
	Device1PortChannelNumber   *int                                                              `json:"device1PortChannelNumber,omitempty"`   // Port Channel number.
	Device2PortChannelUUID     string                                                            `json:"device2PortChannelUuid,omitempty"`     // ID of the port channel.
	Device2PortChannelNumber   *int                                                              `json:"device2PortChannelNumber,omitempty"`   // Port Channel number.
	PortChannelMembers         *[]ResponseLanAutomationGetPortChannelsResponsePortChannelMembers `json:"portChannelMembers,omitempty"`         //
	PortChannelName            string                                                            `json:"portChannelName,omitempty"`            //
}
type ResponseLanAutomationGetPortChannelsResponsePortChannelMembers struct {
	Device1InterfaceUUID string `json:"device1InterfaceUuid,omitempty"` // Device 1 interface UUID.
	Device1Interface     string `json:"device1Interface,omitempty"`     // Device 1 interface.
	Device2InterfaceUUID string `json:"device2InterfaceUuid,omitempty"` // Device 2 interface UUID.
	Device2Interface     string `json:"device2Interface,omitempty"`     // Device 2 interface name.
}
type ResponseLanAutomationCreateANewPortChannelBetweenDevices struct {
	Response *ResponseLanAutomationCreateANewPortChannelBetweenDevicesResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // version
}
type ResponseLanAutomationCreateANewPortChannelBetweenDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationGetPortChannelInformationByID struct {
	Response *[]ResponseLanAutomationGetPortChannelInformationByIDResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // Version number.
}
type ResponseLanAutomationGetPortChannelInformationByIDResponse struct {
	ID                         string                                                                          `json:"id,omitempty"`                         // Unique id of the port-channel link. Concatenation of both device's Port Channel UUID.
	Device1ManagementIPAddress string                                                                          `json:"device1ManagementIPAddress,omitempty"` // IP address of device 1.
	Device1UUID                string                                                                          `json:"device1Uuid,omitempty"`                // ID of device1.
	Device2ManagementIPAddress string                                                                          `json:"device2ManagementIPAddress,omitempty"` // IP address of device 2.
	Device2UUID                string                                                                          `json:"device2Uuid,omitempty"`                // ID of device2.
	Device1PortChannelUUID     string                                                                          `json:"device1PortChannelUuid,omitempty"`     // ID of the port channel.
	Device1PortChannelNumber   *int                                                                            `json:"device1PortChannelNumber,omitempty"`   // Port Channel number.
	Device2PortChannelUUID     string                                                                          `json:"device2PortChannelUuid,omitempty"`     // ID of the port channel.
	Device2PortChannelNumber   *int                                                                            `json:"device2PortChannelNumber,omitempty"`   // Port Channel number.
	PortChannelMembers         *[]ResponseLanAutomationGetPortChannelInformationByIDResponsePortChannelMembers `json:"portChannelMembers,omitempty"`         //
}
type ResponseLanAutomationGetPortChannelInformationByIDResponsePortChannelMembers struct {
	Device1InterfaceUUID string `json:"device1InterfaceUuid,omitempty"` // Device 1 interface UUID.
	Device1Interface     string `json:"device1Interface,omitempty"`     // Device 1 interface.
	Device2InterfaceUUID string `json:"device2InterfaceUuid,omitempty"` // Device 2 interface UUID.
	Device2Interface     string `json:"device2Interface,omitempty"`     // Device 2 interface name.
}
type ResponseLanAutomationDeletePortChannel struct {
	Response *ResponseLanAutomationDeletePortChannelResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // version
}
type ResponseLanAutomationDeletePortChannelResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationAddALanAutomatedLinkToAPortChannel struct {
	Response *ResponseLanAutomationAddALanAutomatedLinkToAPortChannelResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // version
}
type ResponseLanAutomationAddALanAutomatedLinkToAPortChannelResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationRemoveALinkFromPortChannel struct {
	Response *ResponseLanAutomationRemoveALinkFromPortChannelResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // version
}
type ResponseLanAutomationRemoveALinkFromPortChannelResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationLanAutomationStartV2 struct {
	Response *ResponseLanAutomationLanAutomationStartV2Response `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationStartV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2 struct {
	Response *ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // version
}
type ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // url to check the status of task
}
type RequestLanAutomationLanAutomationStart []RequestItemLanAutomationLanAutomationStart // Array of RequestLanAutomationLANAutomationStart
type RequestItemLanAutomationLanAutomationStart struct {
	DiscoveredDeviceSiteNameHierarchy string                                               `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                               `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed management IP address.
	PeerDeviceManagmentIPAddress      string                                               `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed management IP address.
	PrimaryDeviceInterfaceNames       []string                                             `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	IPPools                           *[]RequestItemLanAutomationLanAutomationStartIPPools `json:"ipPools,omitempty"`                           //
	MulitcastEnabled                  *bool                                                `json:"mulitcastEnabled,omitempty"`                  // To enable underlay native multicast.
	HostNamePrefix                    string                                               `json:"hostNamePrefix,omitempty"`                    // Host name prefix which shall be assigned to the discovered device.
	HostNameFileID                    string                                               `json:"hostNameFileId,omitempty"`                    // Use /dna/intent/api/v1/file/namespace/nw_orch api to get the file id for the already uploaded file in nw_orch namespace.
	IsisDomainPwd                     string                                               `json:"isisDomainPwd,omitempty"`                     // IS-IS domain password in plain text.
	RedistributeIsisToBgp             *bool                                                `json:"redistributeIsisToBgp,omitempty"`             // Advertise LAN Automation summary route into BGP.
}
type RequestItemLanAutomationLanAutomationStartIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type RequestLanAutomationLanAutomationDeviceUpdate struct {
	LoopbackUpdateDeviceList *[]RequestLanAutomationLanAutomationDeviceUpdateLoopbackUpdateDeviceList `json:"loopbackUpdateDeviceList,omitempty"` //
	LinkUpdate               *RequestLanAutomationLanAutomationDeviceUpdateLinkUpdate                 `json:"linkUpdate,omitempty"`               //
	HostnameUpdateDevices    *[]RequestLanAutomationLanAutomationDeviceUpdateHostnameUpdateDevices    `json:"hostnameUpdateDevices,omitempty"`    //
}
type RequestLanAutomationLanAutomationDeviceUpdateLoopbackUpdateDeviceList struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewLoopback0IPAddress     string `json:"newLoopback0IPAddress,omitempty"`     // New Loopback0 IP Address from LAN Pool of Device Discovery Site(Shared pool should not be used).
}
type RequestLanAutomationLanAutomationDeviceUpdateLinkUpdate struct {
	SourceDeviceManagementIPAddress      string `json:"sourceDeviceManagementIPAddress,omitempty"`      // Source Device Management IP Address
	SourceDeviceInterfaceName            string `json:"sourceDeviceInterfaceName,omitempty"`            // Source Device Interface Name
	DestinationDeviceManagementIPAddress string `json:"destinationDeviceManagementIPAddress,omitempty"` // Destination Device Management IP Address
	DestinationDeviceInterfaceName       string `json:"destinationDeviceInterfaceName,omitempty"`       // Destination Device Interface Name
	IPPoolName                           string `json:"ipPoolName,omitempty"`                           // Name of the IP LAN Pool, required for Link Add should be from discovery site of source and destination device. It is optional for Link Delete.
}
type RequestLanAutomationLanAutomationDeviceUpdateHostnameUpdateDevices struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewHostName               string `json:"newHostName,omitempty"`               // New hostname for the device
}
type RequestLanAutomationLanAutomationStopAndUpdateDevices []RequestItemLanAutomationLanAutomationStopAndUpdateDevices // Array of RequestLanAutomationLANAutomationStopAndUpdateDevices
type RequestItemLanAutomationLanAutomationStopAndUpdateDevices struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewLoopback0IPAddress     string `json:"newLoopback0IPAddress,omitempty"`     // New Loopback0 IP Address from LAN pool of Device Discovery Site.
}
type RequestLanAutomationCreateANewPortChannelBetweenDevices struct {
	Device1ManagementIPAddress string                                                                       `json:"device1ManagementIPAddress,omitempty"` // Either device1ManagementIPAddress or device1Uuid is required.
	Device1UUID                string                                                                       `json:"device1Uuid,omitempty"`                // Either device1ManagementIPAddress or device1Uuid is required.
	Device2ManagementIPAddress string                                                                       `json:"device2ManagementIPAddress,omitempty"` // Either device2ManagementIPAddress or device2Uuid is required.
	Device2UUID                string                                                                       `json:"device2Uuid,omitempty"`                // Either device2ManagementIPAddress or device2Uuid is required.
	PortChannelMembers         *[]RequestLanAutomationCreateANewPortChannelBetweenDevicesPortChannelMembers `json:"portChannelMembers,omitempty"`         //
}
type RequestLanAutomationCreateANewPortChannelBetweenDevicesPortChannelMembers struct {
	Device1InterfaceUUID string `json:"device1InterfaceUuid,omitempty"` // Either device1InterfaceUuid or device1InterfaceName is required.
	Device1Interface     string `json:"device1Interface,omitempty"`     // Either device1InterfaceUuid or device1InterfaceName is required.
	Device2InterfaceUUID string `json:"device2InterfaceUuid,omitempty"` // Either device2InterfaceUuid or device2InterfaceName is required.
	Device2Interface     string `json:"device2Interface,omitempty"`     // Either device2InterfaceUuid or device2InterfaceName is required.
}
type RequestLanAutomationAddALanAutomatedLinkToAPortChannel struct {
	PortChannelMembers *[]RequestLanAutomationAddALanAutomatedLinkToAPortChannelPortChannelMembers `json:"portChannelMembers,omitempty"` //
}
type RequestLanAutomationAddALanAutomatedLinkToAPortChannelPortChannelMembers struct {
	Device1InterfaceUUID string `json:"device1InterfaceUuid,omitempty"` // Either device1InterfaceUuid or device1InterfaceName is required.
	Device1Interface     string `json:"device1Interface,omitempty"`     // Either device1InterfaceUuid or device1InterfaceName is required.
	Device2InterfaceUUID string `json:"device2InterfaceUuid,omitempty"` // Either device2InterfaceUuid or device1InterfaceName is required.
	Device2Interface     string `json:"device2Interface,omitempty"`     // Either device2InterfaceUuid or device1InterfaceName is required.
}
type RequestLanAutomationRemoveALinkFromPortChannel struct {
	PortChannelMembers *[]RequestLanAutomationRemoveALinkFromPortChannelPortChannelMembers `json:"portChannelMembers,omitempty"` //
}
type RequestLanAutomationRemoveALinkFromPortChannelPortChannelMembers struct {
	Device1InterfaceUUID string `json:"device1InterfaceUuid,omitempty"` // Either device1InterfaceUuid or device1InterfaceName is required.
	Device1Interface     string `json:"device1Interface,omitempty"`     // Either device1InterfaceUuid or device1InterfaceName is required.
	Device2InterfaceUUID string `json:"device2InterfaceUuid,omitempty"` // Either device2InterfaceUuid or device1InterfaceName is required.
	Device2Interface     string `json:"device2Interface,omitempty"`     // Either device2InterfaceUuid or device1InterfaceName is required.
}
type RequestLanAutomationLanAutomationStartV2 []RequestItemLanAutomationLanAutomationStartV2 // Array of RequestLanAutomationLANAutomationStartV2
type RequestItemLanAutomationLanAutomationStartV2 struct {
	DiscoveredDeviceSiteNameHierarchy string                                                          `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                          `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed management IP address.
	PeerDeviceManagmentIPAddress      string                                                          `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed management IP address.
	PrimaryDeviceInterfaceNames       []string                                                        `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	IPPools                           *[]RequestItemLanAutomationLanAutomationStartV2IPPools          `json:"ipPools,omitempty"`                           //
	MulticastEnabled                  *bool                                                           `json:"multicastEnabled,omitempty"`                  // Enable underlay native multicast.
	HostNamePrefix                    string                                                          `json:"hostNamePrefix,omitempty"`                    // Host name prefix assigned to the discovered device.
	HostNameFileID                    string                                                          `json:"hostNameFileId,omitempty"`                    // Use /dna/intent/api/v1/file/namespace/nw_orch API to get the file ID for the already uploaded file in the nw_orch namespace.
	RedistributeIsisToBgp             *bool                                                           `json:"redistributeIsisToBgp,omitempty"`             // Advertise LAN Automation summary route into BGP.
	IsisDomainPwd                     string                                                          `json:"isisDomainPwd,omitempty"`                     // IS-IS domain password in plain text.
	DiscoveryLevel                    *int                                                            `json:"discoveryLevel,omitempty"`                    // Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
	DiscoveryTimeout                  *int                                                            `json:"discoveryTimeout,omitempty"`                  // Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080]. If both timeout and discovery devices list are provided, the stop processing will be attempted whichever happens earlier. Users can always use the LAN Automation delete API to force stop processing.
	DiscoveryDevices                  *[]RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices `json:"discoveryDevices,omitempty"`                  //
}
type RequestItemLanAutomationLanAutomationStartV2IPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type RequestItemLanAutomationLanAutomationStartV2DiscoveryDevices struct {
	DeviceSerialNumber        string `json:"deviceSerialNumber,omitempty"`        // Serial number of the device
	DeviceHostName            string `json:"deviceHostName,omitempty"`            // Hostname of the device
	DeviceSiteNameHierarchy   string `json:"deviceSiteNameHierarchy,omitempty"`   // Site name hierarchy for the device, must be a child site of the discoveredDeviceSiteNameHierarchy or same if it’s not area type.
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Management IP Address of the device
}
type RequestLanAutomationLanAutomationStopAndUpdateDevicesV2 []RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2 // Array of RequestLanAutomationLANAutomationStopAndUpdateDevicesV2
type RequestItemLanAutomationLanAutomationStopAndUpdateDevicesV2 struct {
	DeviceManagementIPAddress string `json:"deviceManagementIPAddress,omitempty"` // Device Management IP Address
	NewLoopback0IPAddress     string `json:"newLoopback0IPAddress,omitempty"`     // New Loopback0 IP Address from LAN pool of Device Discovery Site.
	NewHostName               string `json:"newHostName,omitempty"`               // New hostname to be assigned to the device
}

//LanAutomationSessionCount LAN Automation Session Count - b08b-6b11-4669-a12b
/* Invoke this API to get the total count of LAN Automation sessions.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-session-count
*/
func (s *LanAutomationService) LanAutomationSessionCount() (*ResponseLanAutomationLanAutomationSessionCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationSessionCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationSessionCount()
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationSessionCount")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationSessionCount)
	return result, response, err

}

//LanAutomationLog LAN Automation Log  - 93a9-68c2-480a-85d1
/* Invoke this API to get the LAN Automation session logs.


@param LANAutomationLogQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-log
*/
func (s *LanAutomationService) LanAutomationLog(LANAutomationLogQueryParams *LanAutomationLogQueryParams) (*ResponseLanAutomationLanAutomationLog, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/log"

	queryString, _ := query.Values(LANAutomationLogQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationLanAutomationLog{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationLog(LANAutomationLogQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationLog")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLog)
	return result, response, err

}

//LanAutomationLogByID LAN Automation Log by Id - 55b5-eb50-440a-a123
/* Invoke this API to get the LAN Automation session logs based on the given LAN Automation session id.


@param id id path parameter. LAN Automation session identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-log-by-id
*/
func (s *LanAutomationService) LanAutomationLogByID(id string) (*ResponseLanAutomationLanAutomationLogByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/log/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationLogByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationLogByID(id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationLogById")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLogByID)
	return result, response, err

}

//LanAutomationLogsForIndividualDevices LAN Automation Logs for Individual Devices - b2ac-5af7-45d8-8c4e
/* Invoke this API to get the LAN Automation session logs for individual devices based on the given LAN Automation session id and device serial number.


@param id id path parameter. LAN Automation session identifier.

@param serialNumber serialNumber path parameter. Device serial number.

@param LANAutomationLogsForIndividualDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-logs-for-individual-devices
*/
func (s *LanAutomationService) LanAutomationLogsForIndividualDevices(id string, serialNumber string, LANAutomationLogsForIndividualDevicesQueryParams *LanAutomationLogsForIndividualDevicesQueryParams) (*ResponseLanAutomationLanAutomationLogsForIndividualDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/log/{id}/{serialNumber}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{serialNumber}", fmt.Sprintf("%v", serialNumber), -1)

	queryString, _ := query.Values(LANAutomationLogsForIndividualDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationLanAutomationLogsForIndividualDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationLogsForIndividualDevices(id, serialNumber, LANAutomationLogsForIndividualDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationLogsForIndividualDevices")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLogsForIndividualDevices)
	return result, response, err

}

//LanAutomationActiveSessions LAN Automation Active Sessions - c1bf-69fb-4ad8-979c
/* Invoke this API to get the LAN Automation active session information



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-active-sessions
*/
func (s *LanAutomationService) LanAutomationActiveSessions() (*ResponseLanAutomationLanAutomationActiveSessions, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/sessions"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationActiveSessions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationActiveSessions()
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationActiveSessions")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationActiveSessions)
	return result, response, err

}

//LanAutomationStatus LAN Automation Status - a4ab-087e-4ed9-a3bb
/* Invoke this API to get the LAN Automation session status.


@param LANAutomationStatusQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-status
*/
func (s *LanAutomationService) LanAutomationStatus(LANAutomationStatusQueryParams *LanAutomationStatusQueryParams) (*ResponseLanAutomationLanAutomationStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/status"

	queryString, _ := query.Values(LANAutomationStatusQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationLanAutomationStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStatus(LANAutomationStatusQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStatus")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStatus)
	return result, response, err

}

//LanAutomationStatusByID LAN Automation Status by Id - 5b99-8b6e-47b8-9882
/* Invoke this API to get the LAN Automation session status based on the given Lan Automation session id.


@param id id path parameter. LAN Automation session identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-status-by-id
*/
func (s *LanAutomationService) LanAutomationStatusByID(id string) (*ResponseLanAutomationLanAutomationStatusByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/status/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationStatusByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStatusByID(id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStatusById")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStatusByID)
	return result, response, err

}

//GetPortChannels Get Port Channels - c9a9-9a58-452a-87b8
/* Returns a list of Port Channel between the LAN Automation associated devices.


@param GetPortChannelsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-channels
*/
func (s *LanAutomationService) GetPortChannels(GetPortChannelsQueryParams *GetPortChannelsQueryParams) (*ResponseLanAutomationGetPortChannels, *resty.Response, error) {
	path := "/dna/intent/api/v1/lanAutomation/portChannels"

	queryString, _ := query.Values(GetPortChannelsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLanAutomationGetPortChannels{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortChannels(GetPortChannelsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortChannels")
	}

	result := response.Result().(*ResponseLanAutomationGetPortChannels)
	return result, response, err

}

//GetPortChannelInformationByID Get Port Channel information by id - 0d8c-1a1b-4cf8-82f4
/* This API retrieves Port Channel information using its ID.


@param id id path parameter. ID of the port channel.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-channel-information-by-id
*/
func (s *LanAutomationService) GetPortChannelInformationByID(id string) (*ResponseLanAutomationGetPortChannelInformationByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/lanAutomation/portChannels/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationGetPortChannelInformationByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortChannelInformationByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetPortChannelInformationById")
	}

	result := response.Result().(*ResponseLanAutomationGetPortChannelInformationByID)
	return result, response, err

}

//LanAutomationStart LAN Automation Start - 9795-f927-469a-a6d2
/* Invoke this API to start LAN Automation for the given site.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-start
*/
func (s *LanAutomationService) LanAutomationStart(requestLanAutomationLANAutomationStart *RequestLanAutomationLanAutomationStart) (*ResponseLanAutomationLanAutomationStart, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStart).
		SetResult(&ResponseLanAutomationLanAutomationStart{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStart(requestLanAutomationLANAutomationStart)
		}

		return nil, response, fmt.Errorf("error with operation LanAutomationStart")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStart)
	return result, response, err

}

//CreateANewPortChannelBetweenDevices Create a New Port Channel between devices - 1690-29e1-475b-8733
/* This API creates a Port Channel between two LAN Automation associated devices using the PAgP protocol, with a minimum of 2 and a maximum of 8 LAN Automated interfaces in UP status.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a-new-port-channel-between-devices
*/
func (s *LanAutomationService) CreateANewPortChannelBetweenDevices(requestLanAutomationCreateANewPortChannelBetweenDevices *RequestLanAutomationCreateANewPortChannelBetweenDevices) (*ResponseLanAutomationCreateANewPortChannelBetweenDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/lanAutomation/portChannels"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationCreateANewPortChannelBetweenDevices).
		SetResult(&ResponseLanAutomationCreateANewPortChannelBetweenDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateANewPortChannelBetweenDevices(requestLanAutomationCreateANewPortChannelBetweenDevices)
		}

		return nil, response, fmt.Errorf("error with operation CreateANewPortChannelBetweenDevices")
	}

	result := response.Result().(*ResponseLanAutomationCreateANewPortChannelBetweenDevices)
	return result, response, err

}

//AddALanAutomatedLinkToAPortChannel Add a LAN Automated link to a Port Channel - f6a9-e8d7-4d6b-9ef2
/* This API adds a new LAN Automated link as a member to an existing Port Channel, provided the interface is in UP state and already LAN Automated.


@param id id path parameter. ID of the port channel.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-a-lan-automated-link-to-a-port-channel
*/
func (s *LanAutomationService) AddALanAutomatedLinkToAPortChannel(id string, requestLanAutomationAddALANAutomatedLinkToAPortChannel *RequestLanAutomationAddALanAutomatedLinkToAPortChannel) (*ResponseLanAutomationAddALanAutomatedLinkToAPortChannel, *resty.Response, error) {
	path := "/dna/intent/api/v1/lanAutomation/portChannels/{id}/addLink"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationAddALANAutomatedLinkToAPortChannel).
		SetResult(&ResponseLanAutomationAddALanAutomatedLinkToAPortChannel{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddALanAutomatedLinkToAPortChannel(id, requestLanAutomationAddALANAutomatedLinkToAPortChannel)
		}

		return nil, response, fmt.Errorf("error with operation AddALanAutomatedLinkToAPortChannel")
	}

	result := response.Result().(*ResponseLanAutomationAddALanAutomatedLinkToAPortChannel)
	return result, response, err

}

//RemoveALinkFromPortChannel Remove a link from Port Channel - 05bf-0a11-4e1a-824f
/* This API removes a member link from an existing Port Channel, reverting the link to a P2P L3 interface.


@param id id path parameter. ID of the port channel.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-a-link-from-port-channel
*/
func (s *LanAutomationService) RemoveALinkFromPortChannel(id string, requestLanAutomationRemoveALinkFromPortChannel *RequestLanAutomationRemoveALinkFromPortChannel) (*ResponseLanAutomationRemoveALinkFromPortChannel, *resty.Response, error) {
	path := "/dna/intent/api/v1/lanAutomation/portChannels/{id}/removeLink"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationRemoveALinkFromPortChannel).
		SetResult(&ResponseLanAutomationRemoveALinkFromPortChannel{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveALinkFromPortChannel(id, requestLanAutomationRemoveALinkFromPortChannel)
		}

		return nil, response, fmt.Errorf("error with operation RemoveALinkFromPortChannel")
	}

	result := response.Result().(*ResponseLanAutomationRemoveALinkFromPortChannel)
	return result, response, err

}

//LanAutomationStartV2 LAN Automation Start V2 - 51ba-8921-46da-9bec
/* Invoke V2 LAN Automation Start API, which supports optional auto-stop processing feature based on the provided timeout or a specific device list, or both. The stop processing will be executed automatically when either of the cases is satisfied, without specifically calling the stop API. The V2 API behaves similarly to V1 if no timeout or device list is provided, and the user needs to call the stop API for LAN Automation stop processing. With the V2 API, the user can also specify the level up to which the devices can be LAN automated.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-start-v2
*/
func (s *LanAutomationService) LanAutomationStartV2(requestLanAutomationLANAutomationStartV2 *RequestLanAutomationLanAutomationStartV2) (*ResponseLanAutomationLanAutomationStartV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/lan-automation"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStartV2).
		SetResult(&ResponseLanAutomationLanAutomationStartV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStartV2(requestLanAutomationLANAutomationStartV2)
		}

		return nil, response, fmt.Errorf("error with operation LanAutomationStartV2")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStartV2)
	return result, response, err

}

//LanAutomationDeviceUpdate LAN Automation Device Update - 1190-5ac3-4e88-bd5e
/* Invoke this API to perform a DAY-N update on LAN Automation-related devices. Supported features include Loopback0 IP update, hostname update, link addition, and link deletion.


@param LANAutomationDeviceUpdateQueryParams Filtering parameter
*/
func (s *LanAutomationService) LanAutomationDeviceUpdate(requestLanAutomationLANAutomationDeviceUpdate *RequestLanAutomationLanAutomationDeviceUpdate, LANAutomationDeviceUpdateQueryParams *LanAutomationDeviceUpdateQueryParams) (*ResponseLanAutomationLanAutomationDeviceUpdate, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/updateDevice"

	queryString, _ := query.Values(LANAutomationDeviceUpdateQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestLanAutomationLANAutomationDeviceUpdate).
		SetResult(&ResponseLanAutomationLanAutomationDeviceUpdate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationDeviceUpdate(requestLanAutomationLANAutomationDeviceUpdate, LANAutomationDeviceUpdateQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationDeviceUpdate")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationDeviceUpdate)
	return result, response, err

}

//LanAutomationStopAndUpdateDevices LAN Automation Stop and Update Devices - 0780-4a4c-44cb-bae8
/* Invoke this API to stop LAN Automation and Update Loopback0 IP Address of Devices, discovered in the current session


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.

*/
func (s *LanAutomationService) LanAutomationStopAndUpdateDevices(id string, requestLanAutomationLANAutomationStopAndUpdateDevices *RequestLanAutomationLanAutomationStopAndUpdateDevices) (*ResponseLanAutomationLanAutomationStopAndUpdateDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStopAndUpdateDevices).
		SetResult(&ResponseLanAutomationLanAutomationStopAndUpdateDevices{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStopAndUpdateDevices(id, requestLanAutomationLANAutomationStopAndUpdateDevices)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStopAndUpdateDevices")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStopAndUpdateDevices)
	return result, response, err

}

//LanAutomationStopAndUpdateDevicesV2 LAN Automation Stop and Update Devices V2 - 9381-ba28-42c9-9e6a
/* Invoke this API to stop LAN Automation and update device parameters such as Loopback0 IP address and/or hostname discovered in the current session.


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.

*/
func (s *LanAutomationService) LanAutomationStopAndUpdateDevicesV2(id string, requestLanAutomationLANAutomationStopAndUpdateDevicesV2 *RequestLanAutomationLanAutomationStopAndUpdateDevicesV2) (*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomationStopAndUpdateDevicesV2).
		SetResult(&ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStopAndUpdateDevicesV2(id, requestLanAutomationLANAutomationStopAndUpdateDevicesV2)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStopAndUpdateDevicesV2")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStopAndUpdateDevicesV2)
	return result, response, err

}

//LanAutomationStop LAN Automation Stop - e6a0-da69-4adb-8929
/* Invoke this API to stop LAN Automation for the given site.


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!lan-automation-stop
*/
func (s *LanAutomationService) LanAutomationStop(id string) (*ResponseLanAutomationLanAutomationStop, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomationStop{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LanAutomationStop(
				id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStop")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStop)
	return result, response, err

}

//DeletePortChannel Delete Port Channel - 9099-fb62-4c98-bf2e
/* This API deletes a Port Channel between LAN Automation associated devices using a valid Port Channel ID.


@param id id path parameter. ID of the port channel.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-channel
*/
func (s *LanAutomationService) DeletePortChannel(id string) (*ResponseLanAutomationDeletePortChannel, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/lanAutomation/portChannels/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationDeletePortChannel{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortChannel(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortChannel")
	}

	result := response.Result().(*ResponseLanAutomationDeletePortChannel)
	return result, response, err

}
