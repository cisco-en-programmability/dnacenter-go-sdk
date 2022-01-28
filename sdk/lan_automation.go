package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LanAutomationService service

type LanAutomationLogQueryParams struct {
	Offset string `url:"offset,omitempty"` //Offset/starting row of the LAN Automation session from which logs are required
	Limit  string `url:"limit,omitempty"`  //Number of LAN Automations sessions to be retrieved
}
type LanAutomationStatusQueryParams struct {
	Offset string `url:"offset,omitempty"` //Offset/starting row of the LAN Automation session from which status needs to be retrieved
	Limit  string `url:"limit,omitempty"`  //Number of LAN Automations session status to be retrieved
}

type ResponseLanAutomationLanAutomationStart struct {
	Response *ResponseLanAutomationLanAutomationStartResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStartResponse struct {
	ErrorCode string `json:"errorCode,omitempty"` // ErrorCode Value
	Message   string `json:"message,omitempty"`   // Descriptionn of the error code
	Detail    string `json:"detail,omitempty"`    // Detailed  information of the error code
}
type ResponseLanAutomationLanAutomationSessionCount struct {
	Response *ResponseLanAutomationLanAutomationSessionCountResponse `json:"response,omitempty"` //
}
type ResponseLanAutomationLanAutomationSessionCountResponse struct {
	SessionCount string `json:"sessionCount,omitempty"` // Total Number of session count
}
type ResponseLanAutomationLanAutomationLog struct {
	Response *[]ResponseLanAutomationLanAutomationLogResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogResponse struct {
	NwOrchID string                                                `json:"nwOrchId,omitempty"` // Network Orchestration Identifier
	Entry    *[]ResponseLanAutomationLanAutomationLogResponseEntry `json:"entry,omitempty"`    //
}
type ResponseLanAutomationLanAutomationLogResponseEntry struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Log level and the value could be Info, Warning and Error
	TimeStamp string `json:"timeStamp,omitempty"` // The time at which the log message created
	Record    string `json:"record,omitempty"`    // Log message in detail
	DeviceID  string `json:"deviceId,omitempty"`  // The device serial number for which the log message is associated
}
type ResponseLanAutomationLanAutomationLogByID struct {
	Response *[]ResponseLanAutomationLanAutomationLogByIDResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationLogByIDResponse struct {
	NwOrchID string                                                    `json:"nwOrchId,omitempty"` // Network Orchestration Identifier
	Entry    *[]ResponseLanAutomationLanAutomationLogByIDResponseEntry `json:"entry,omitempty"`    //
}
type ResponseLanAutomationLanAutomationLogByIDResponseEntry struct {
	LogLevel  string `json:"logLevel,omitempty"`  // Log level and the value could be Info, Warning and Error
	TimeStamp string `json:"timeStamp,omitempty"` // The time at which the log message created
	Record    string `json:"record,omitempty"`    // Log message in detail
	DeviceID  string `json:"deviceId,omitempty"`  // The device serial number for which the log message is associated
}
type ResponseLanAutomationLanAutomationStatus struct {
	Response *[]ResponseLanAutomationLanAutomationStatusResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusResponse struct {
	ID                                string                                                                  `json:"id,omitempty"`                                // System generated identifier
	DiscoveredDeviceSiteNameHierarchy string                                                                  `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered sevice site name
	PrimaryDeviceManagmentIPAddress   string                                                                  `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address
	IPPoolList                        *[]ResponseLanAutomationLanAutomationStatusResponseIPPoolList           `json:"ipPoolList,omitempty"`                        //
	PrimaryDeviceInterfaceNames       []string                                                                `json:"primaryDeviceInterfaceNames,omitempty"`       // The List of interfaces on primary seed device via which the discovered devices are connected
	Status                            string                                                                  `json:"status,omitempty"`                            // Status of LAN Automation session and provides the number of discovered devices
	Action                            string                                                                  `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation Session
	CreationTime                      string                                                                  `json:"creationTime,omitempty"`                      // LAN Automation session creation time
	MulticastEnabled                  *bool                                                                   `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not
	PeerDeviceManagmentIPAddress      string                                                                  `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
}
type ResponseLanAutomationLanAutomationStatusResponseIPPoolList struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool
}
type ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device
	State              string   `json:"state,omitempty"`              // state of the device like added to inventory/ deleted from inventory
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // The list of IP address used by device
}
type ResponseLanAutomationLanAutomationStatusByID struct {
	Response *[]ResponseLanAutomationLanAutomationStatusByIDResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStatusByIDResponse struct {
	ID                                string                                                                      `json:"id,omitempty"`                                // System generated identifier
	DiscoveredDeviceSiteNameHierarchy string                                                                      `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered sevice site name
	PrimaryDeviceManagmentIPAddress   string                                                                      `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address
	IPPoolList                        *[]ResponseLanAutomationLanAutomationStatusByIDResponseIPPoolList           `json:"ipPoolList,omitempty"`                        //
	PrimaryDeviceInterfaceNames       []string                                                                    `json:"primaryDeviceInterfaceNames,omitempty"`       // The List of interfaces on primary seed device via which the discovered devices are connected
	Status                            string                                                                      `json:"status,omitempty"`                            // Status of LAN Automation session and provides the number of discovered devices
	Action                            string                                                                      `json:"action,omitempty"`                            // State (START/STOP) of the LAN Automation Session
	CreationTime                      string                                                                      `json:"creationTime,omitempty"`                      // LAN Automation session creation time
	MulticastEnabled                  *bool                                                                       `json:"multicastEnabled,omitempty"`                  // Shows whether underlay multicast is enabled or not
	PeerDeviceManagmentIPAddress      string                                                                      `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address
	DiscoveredDeviceList              *[]ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList `json:"discoveredDeviceList,omitempty"`              //
}
type ResponseLanAutomationLanAutomationStatusByIDResponseIPPoolList struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool
}
type ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList struct {
	Name               string   `json:"name,omitempty"`               // Name of the device
	SerialNumber       string   `json:"serialNumber,omitempty"`       // Serial number of the device
	State              string   `json:"state,omitempty"`              // state of the device like added to inventory/ deleted from inventory
	IPAddressInUseList []string `json:"ipAddressInUseList,omitempty"` // The list of IP address used by device
}
type ResponseLanAutomationLanAutomationStop struct {
	Response *ResponseLanAutomationLanAutomationStopResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseLanAutomationLanAutomationStopResponse struct {
	ErrorCode string `json:"errorCode,omitempty"` // ErrorCode Value
	Message   string `json:"message,omitempty"`   // Descriptionn of the error code
	Detail    string `json:"detail,omitempty"`    // Detailed information of the error code
}
type RequestLanAutomationLanAutomationStart []RequestItemLanAutomationLanAutomationStart // Array of RequestLanAutomationLANAutomationStart
type RequestItemLanAutomationLanAutomationStart struct {
	DiscoveredDeviceSiteNameHierarchy string                                               `json:"discoveredDeviceSiteNameHierarchy,omitempty"` // Discovered device site name
	PrimaryDeviceManagmentIPAddress   string                                               `json:"primaryDeviceManagmentIPAddress,omitempty"`   // Primary seed device management IP address
	PeerDeviceManagmentIPAddress      string                                               `json:"peerDeviceManagmentIPAddress,omitempty"`      // Peer seed device management IP address
	PrimaryDeviceInterfaceNames       []string                                             `json:"primaryDeviceInterfaceNames,omitempty"`       // The list of interfaces on primary seed device via which the discovered devices are connected
	IPPools                           *[]RequestItemLanAutomationLanAutomationStartIPPools `json:"ipPools,omitempty"`                           //
	MulitcastEnabled                  *bool                                                `json:"mulitcastEnabled,omitempty"`                  // To enable underlay native multicast
	HostNamePrefix                    string                                               `json:"hostNamePrefix,omitempty"`                    // Host name prefix which shall be assigned to the discovered device
	HostNameFileID                    string                                               `json:"hostNameFileId,omitempty"`                    // By using /dna/intent/api/v1/file/namespace/nw_orch api get the file id for the already uploaded file for the nw_orch namespace
	IsisDomainPwd                     string                                               `json:"isisDomainPwd,omitempty"`                     // isis domain password in plain text.
}
type RequestItemLanAutomationLanAutomationStartIPPools struct {
	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool
	IPPoolRole string `json:"ipPoolRole,omitempty"` // Role of the IP pool
}

//LanAutomationSessionCount LAN Automation Session Count - b08b-6b11-4669-a12b
/* Invoke this API to get the total count of LAN Automation sessions


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
/* Invoke this API to get the LAN Automation session logs


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
/* Invoke this API to get the  LAN Automation session logs based on the given Lan Automation session Id


@param id id path parameter. LAN Automation Session Identifier

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
/* Invoke this API to get the LAN Automation session status


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
/* Invoke this API to get the LAN Automation session status based on the given Lan Automation session Id


@param id id path parameter. LAN Automation Session Identifier

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

//LanAutomationStart LAN Automation - 9795-f927-469a-a6d2
/* Invoke this API to start LAN Automation for the given site


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
		return nil, response, fmt.Errorf("error with operation LanAutomationStart")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStart)
	return result, response, err

}

//LanAutomationStop LAN Automation - e6a0-da69-4adb-8929
/* Invoke this API to stop LAN Automation for the given site


@param id id path parameter. LAN Automation Id needs to be retrived via /dna/intent/status

*/
func (s *LanAutomationService) LanAutomationStop(id string) (*ResponseLanAutomationLanAutomationStop, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation LanAutomationStop")
	}

	result := response.Result().(*ResponseLanAutomationLanAutomationStop)
	return result, response, err

}
