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
	DiscoveryDevices                  *[]ResponseLanAutomationLanAutomationStatusResponseDiscoveryDevices     `json:"discoveryDevices,omitempty"`                  // Specific devices that will be LAN Automated in this session. Any other device discovered via DHCP will be attempted for a reset and reload to bring it back to the PnP agent state at the end of the LAN Automation process before process completion. The maximum supported devices that can be provided for a session is 50.
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
type ResponseLanAutomationLanAutomationStatusResponseDiscoveryDevices interface{}
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
	DiscoveryDevices                  *[]ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveryDevices     `json:"discoveryDevices,omitempty"`                  // Specific devices that will be LAN Automated in this session. Any other device discovered via DHCP will be attempted for a reset and reload to bring it back to the PnP agent state at the end of the LAN Automation process before process completion. The maximum supported devices that can be provided for a session is 50.
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
type ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveryDevices interface{}
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
	IPPoolName                           string `json:"ipPoolName,omitempty"`                           // Name of the IP LAN Pool, required for Link Add should be from discovery site of source and destination device.
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

//LanAutomationStartV2 LAN Automation Start V2 - 51ba-8921-46da-9bec
/* Invoke V2 LAN Automation Start API, which supports optional auto-stop processing feature based on the provided timeout or a specific device list, or both. The stop processing will be executed automatically when either of the cases is satisfied, without specifically calling the stop API. The V2 API behaves similarly to  if no timeout or device list is provided, and the user needs to call the stop API for LAN Automation stop processing. With the V2 API, the user can also specify the level up to which the devices can be LAN automated.



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
			return s.LanAutomationStop(id)
		}
		return nil, response, fmt.Errorf("error with operation LanAutomationStop")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStop)
	return result, response, err

}
