package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LanAutomationService service

type LanAutomationLogQueryParams struct {
	Offset int `url:"offset,omitempty"` //Starting index of the LAN Automation session. Minimum value is 1.
	Limit  int `url:"limit,omitempty"`  //Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
}
type LanAutomationStatusQueryParams struct {
	Offset int `url:"offset,omitempty"` //Starting index of the LAN Automation session. Minimum value is 1.
	Limit  int `url:"limit,omitempty"`  //Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
}

type ResponseLanAutomationLanAutomation2 struct {
	Response *ResponseLanAutomationLanAutomation2Response `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomation2Response struct {
	ErrorCode string `json:"errorCode,omitempty"` // Error code value.
	Message   string `json:"message,omitempty"`   // Description of the error code.
	Detail    string `json:"detail,omitempty"`    // Detailed information of the error code.
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
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO and WARNING.
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
	LogLevel  string `json:"logLevel,omitempty"`  // Supported levels are ERROR, INFO and WARNING.
	TimeStamp string `json:"timeStamp,omitempty"` // Time at which the log message is created.
	Record    string `json:"record,omitempty"`    // Detailed log message.
	DeviceID  string `json:"deviceId,omitempty"`  // Device serial number for which the log message is associated.
}
type ResponseLanAutomationLanAutomationStatus struct {
	Response *[]ResponseLanAutomationLanAutomationStatusResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusResponse struct {
	ID                                string                                                                  `json:"id,omitempty"`                                // LAN Automation session id.
	DiscoveredDeviceSiteNameHierarchy string                                                                  `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                                  `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address.
	IPPoolList                        *[]ResponseLanAutomationLanAutomationStatusResponseIPPoolList           `json:"ipPoolList,omitempty"`                        //
	PrimaryDeviceInterfaceNames       []string                                                                `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	Status                            string                                                                  `json:"status,omitempty"`                            // Status of the LAN Automation session along with the number of discovered devices.
	Action                            string                                                                  `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation session.
	CreationTime                      string                                                                  `json:"creationTime,omitempty"`                      // LAN Automation session creation time.
	MulticastEnabled                  *bool                                                                   `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not.
	PeerDeviceManagmentIPAddress      string                                                                  `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address.
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
	RedistributeIsisToBgp             *bool                                                                   `json:"redistributeIsisToBgp,omitempty"`             // Shows whether advertise LAN Automation summary route into BGP is enabled or not.
}
type ResponseLanAutomationLanAutomationStatusResponseIPPoolList struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device.
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device.
	State              string   `json:"state,omitempty"`              // State of the device (Added to inventory/Deleted from inventory).
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // List of IP address used by the device.
}
type ResponseLanAutomationLanAutomationStatusByID struct {
	Response *[]ResponseLanAutomationLanAutomationStatusByIDResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusByIDResponse struct {
	ID                                string                                                                      `json:"id,omitempty"`                                // LAN Automation session id.
	DiscoveredDeviceSiteNameHierarchy string                                                                      `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                                                      `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address.
	IPPoolList                        *[]ResponseLanAutomationLanAutomationStatusByIDResponseIPPoolList           `json:"ipPoolList,omitempty"`                        //
	PrimaryDeviceInterfaceNames       []string                                                                    `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	Status                            string                                                                      `json:"status,omitempty"`                            // Status of the LAN Automation session along with the number of discovered devices.
	Action                            string                                                                      `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation session.
	CreationTime                      string                                                                      `json:"creationTime,omitempty"`                      // LAN Automation session creation time.
	MulticastEnabled                  *bool                                                                       `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not.
	PeerDeviceManagmentIPAddress      string                                                                      `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address.
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
	RedistributeIsisToBgp             *bool                                                                       `json:"redistributeIsisToBgp,omitempty"`             // Shows whether advertise LAN Automation summary route into BGP is enabled or not.
}
type ResponseLanAutomationLanAutomationStatusByIDResponseIPPoolList struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}
type ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device.
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device.
	State              string   `json:"state,omitempty"`              // State of the device (Added to inventory/Deleted from inventory).
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // List of IP address used by the device.
}
type ResponseLanAutomationLanAutomation struct {
	Response *ResponseLanAutomationLanAutomationResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationResponse struct {
	ErrorCode string `json:"errorCode,omitempty"` // Error code value.
	Message   string `json:"message,omitempty"`   // Description of the error code.
	Detail    string `json:"detail,omitempty"`    // Detailed information of the error code.
}
type RequestLanAutomationLanAutomation2 []RequestItemLanAutomationLanAutomation2 // Array of RequestLanAutomationLANAutomation2
type RequestItemLanAutomationLanAutomation2 struct {
	DiscoveredDeviceSiteNameHierarchy string                                           `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name.
	PrimaryDeviceManagmentIPAddress   string                                           `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed management IP address.
	PeerDeviceManagmentIPAddress      string                                           `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed management IP address.
	PrimaryDeviceInterfaceNames       []string                                         `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed via which the discovered devices are connected.
	IPPools                           *[]RequestItemLanAutomationLanAutomation2IPPools `json:"ipPools,omitempty"`                           //
	MulitcastEnabled                  *bool                                            `json:"mulitcastEnabled,omitempty"`                  // To enable underlay native multicast.
	HostNamePrefix                    string                                           `json:"hostNamePrefix,omitempty"`                    // Host name prefix which shall be assigned to the discovered device.
	HostNameFileID                    string                                           `json:"hostNameFileId,omitempty"`                    // Use /dna/intent/api/v1/file/namespace/nw_orch api to get the file id for the already uploaded file in nw_orch namespace.
	IsisDomainPwd                     string                                           `json:"isisDomainPwd,omitempty"`                     // IS-IS domain password in plain text.
	RedistributeIsisToBgp             *bool                                            `json:"redistributeIsisToBgp,omitempty"`             // Advertise LAN Automation summary route into BGP.
}
type RequestItemLanAutomationLanAutomation2IPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool.
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
}

//LanAutomationSessionCount LAN Automation Session Count - b08b-6b11-4669-a12b
/* Invoke this API to get the total count of LAN Automation sessions.


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
		return nil, response, fmt.Errorf("error with operation LanAutomationSessionCount")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationSessionCount)
	return result, response, err

}

//LanAutomationLog LAN Automation Log  - 93a9-68c2-480a-85d1
/* Invoke this API to get the LAN Automation session logs.


@param LANAutomationLogQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation LanAutomationLog")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLog)
	return result, response, err

}

//LanAutomationLogByID LAN Automation Log by Id - 55b5-eb50-440a-a123
/* Invoke this API to get the LAN Automation session logs based on the given LAN Automation session id.


@param id id path parameter. LAN Automation session identifier.

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
		return nil, response, fmt.Errorf("error with operation LanAutomationLogById")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationLogByID)
	return result, response, err

}

//LanAutomationStatus LAN Automation Status - a4ab-087e-4ed9-a3bb
/* Invoke this API to get the LAN Automation session status.


@param LANAutomationStatusQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation LanAutomationStatus")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStatus)
	return result, response, err

}

//LanAutomationStatusByID LAN Automation Status by Id - 5b99-8b6e-47b8-9882
/* Invoke this API to get the LAN Automation session status based on the given Lan Automation session id.


@param id id path parameter. LAN Automation session identifier.

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
		return nil, response, fmt.Errorf("error with operation LanAutomationStatusById")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStatusByID)
	return result, response, err

}

//LanAutomation2 LAN Automation - 9795-f927-469a-a6d2
/* Invoke this API to start LAN Automation for the given site.


 */
func (s *LanAutomationService) LanAutomation2(requestLanAutomationLANAutomation2 *RequestLanAutomationLanAutomation2) (*ResponseLanAutomationLanAutomation2, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLanAutomationLANAutomation2).
		SetResult(&ResponseLanAutomationLanAutomation2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LanAutomation2")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomation2)
	return result, response, err

}

//LanAutomation LAN Automation - e6a0-da69-4adb-8929
/* Invoke this API to stop LAN Automation for the given site.


@param id id path parameter. LAN Automation id can be obtained from /dna/intent/api/v1/lan-automation/status.

*/
func (s *LanAutomationService) LanAutomation(id string) (*ResponseLanAutomationLanAutomation, *resty.Response, error) {
	path := "/dna/intent/api/v1/lan-automation/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLanAutomationLanAutomation{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LanAutomation")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomation)
	return result, response, err

}
