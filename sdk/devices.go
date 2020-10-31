package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DevicesService is the service to communicate with the Devices API endpoint
type DevicesService service

// ExportDeviceDTO is the ExportDeviceDTO definition
type ExportDeviceDTO struct {
	DeviceUUIDs   []string `json:"deviceUuids,omitempty"`   //
	ID            string   `json:"id,omitempty"`            //
	OperationEnum string   `json:"operationEnum,omitempty"` //
	Parameters    []string `json:"parameters,omitempty"`    //
	Password      string   `json:"password,omitempty"`      //
}

// InventoryDeviceInfo is the InventoryDeviceInfo definition
type InventoryDeviceInfo struct {
	CliTransport            string                    `json:"cliTransport,omitempty"`            //
	ComputeDevice           bool                      `json:"computeDevice,omitempty"`           //
	EnablePassword          string                    `json:"enablePassword,omitempty"`          //
	ExtendedDiscoveryInfo   string                    `json:"extendedDiscoveryInfo,omitempty"`   //
	HTTPPassword            string                    `json:"httpPassword,omitempty"`            //
	HTTPPort                string                    `json:"httpPort,omitempty"`                //
	HTTPSecure              bool                      `json:"httpSecure,omitempty"`              //
	HTTPUserName            string                    `json:"httpUserName,omitempty"`            //
	IPAddress               []string                  `json:"ipAddress,omitempty"`               //
	MerakiOrgID             []string                  `json:"merakiOrgId,omitempty"`             //
	NetconfPort             string                    `json:"netconfPort,omitempty"`             //
	Password                string                    `json:"password,omitempty"`                //
	SerialNumber            string                    `json:"serialNumber,omitempty"`            //
	SNMPAuthPassphrase      string                    `json:"snmpAuthPassphrase,omitempty"`      //
	SNMPAuthProtocol        string                    `json:"snmpAuthProtocol,omitempty"`        //
	SNMPMode                string                    `json:"snmpMode,omitempty"`                //
	SNMPPrivPassphrase      string                    `json:"snmpPrivPassphrase,omitempty"`      //
	SNMPPrivProtocol        string                    `json:"snmpPrivProtocol,omitempty"`        //
	SNMPROCommunity         string                    `json:"snmpROCommunity,omitempty"`         //
	SNMPRWCommunity         string                    `json:"snmpRWCommunity,omitempty"`         //
	SNMPRetry               int                       `json:"snmpRetry,omitempty"`               //
	SNMPTimeout             int                       `json:"snmpTimeout,omitempty"`             //
	SNMPUserName            string                    `json:"snmpUserName,omitempty"`            //
	SNMPVersion             string                    `json:"snmpVersion,omitempty"`             //
	Type                    string                    `json:"type,omitempty"`                    //
	UpdateMgmtIPaddressList []UpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //
	UserName                string                    `json:"userName,omitempty"`                //
}

// NetworkDeviceBriefNIO is the NetworkDeviceBriefNIO definition
type NetworkDeviceBriefNIO struct {
	ID         string `json:"id,omitempty"`         //
	Role       string `json:"role,omitempty"`       //
	RoleSource string `json:"roleSource,omitempty"` //
}

// UpdateMgmtIPaddressList is the UpdateMgmtIPaddressList definition
type UpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` //
	NewMgmtIPAddress   string `json:"newMgmtIpAddress,omitempty"`   //
}

// AirQualityHealth is the AirQualityHealth definition
type AirQualityHealth struct {
	Ghz24  int `json:"Ghz24,omitempty"`  //
	Ghz50  int `json:"Ghz50,omitempty"`  //
	Radio0 int `json:"radio0,omitempty"` //
	Radio1 int `json:"radio1,omitempty"` //
}

// ClientCount is the ClientCount definition
type ClientCount struct {
	Ghz24  int `json:"Ghz24,omitempty"`  //
	Ghz50  int `json:"Ghz50,omitempty"`  //
	Radio0 int `json:"radio0,omitempty"` //
	Radio1 int `json:"radio1,omitempty"` //
}

// DeviceIfListResult is the DeviceIfListResult definition
type DeviceIfListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// DeviceIfResult is the DeviceIfResult definition
type DeviceIfResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DevicesResponse is the DevicesResponse definition
type DevicesResponse struct {
	Response   []Response `json:"response,omitempty"`   //
	TotalCount int        `json:"totalCount,omitempty"` //
	Version    string     `json:"version,omitempty"`    //
}

// FunctionDetails is the FunctionDetails definition
type FunctionDetails struct {
	AttributeInfo string `json:"attributeInfo,omitempty"` //
	ID            string `json:"id,omitempty"`            //
	PropertyName  string `json:"propertyName,omitempty"`  //
	StringValue   string `json:"stringValue,omitempty"`   //
}

// FunctionalCapability is the FunctionalCapability definition
type FunctionalCapability struct {
	AttributeInfo   string            `json:"attributeInfo,omitempty"`   //
	FunctionDetails []FunctionDetails `json:"functionDetails,omitempty"` //
	FunctionName    string            `json:"functionName,omitempty"`    //
	FunctionOpState string            `json:"functionOpState,omitempty"` //
	ID              string            `json:"id,omitempty"`              //
}

// FunctionalCapabilityListResult is the FunctionalCapabilityListResult definition
type FunctionalCapabilityListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// FunctionalCapabilityResult is the FunctionalCapabilityResult definition
type FunctionalCapabilityResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetDeviceDetailResponse is the GetDeviceDetailResponse definition
type GetDeviceDetailResponse struct {
	Response Response `json:"response,omitempty"` //
}

// GetDeviceEnrichmentDetailsResponse is the GetDeviceEnrichmentDetailsResponse definition
type GetDeviceEnrichmentDetailsResponse struct {
	DeviceDetails DeviceDetails `json:"deviceDetails,omitempty"` //
}

// InterferenceHealth is the InterferenceHealth definition
type InterferenceHealth struct {
	Ghz24  int `json:"Ghz24,omitempty"`  //
	Ghz50  int `json:"Ghz50,omitempty"`  //
	Radio0 int `json:"radio0,omitempty"` //
	Radio1 int `json:"radio1,omitempty"` //
}

// ModuleListResult is the ModuleListResult definition
type ModuleListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// ModuleResult is the ModuleResult definition
type ModuleResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// NetworkDeviceBriefNIOResult is the NetworkDeviceBriefNIOResult definition
type NetworkDeviceBriefNIOResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// NetworkDeviceListResult is the NetworkDeviceListResult definition
type NetworkDeviceListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// NetworkDeviceResult is the NetworkDeviceResult definition
type NetworkDeviceResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// NoiseHealth is the NoiseHealth definition
type NoiseHealth struct {
	Ghz50  int `json:"Ghz50,omitempty"`  //
	Radio1 int `json:"radio1,omitempty"` //
}

// RawCliInfoNIOListResult is the RawCliInfoNIOListResult definition
type RawCliInfoNIOListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// RegisterNetworkDeviceResult is the RegisterNetworkDeviceResult definition
type RegisterNetworkDeviceResult struct {
	Response RegisterNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                        `json:"version,omitempty"`  //
}

// RegisterNetworkDeviceResponse is the Response definition
type RegisterNetworkDeviceResponse struct {
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
}

// SuccessResult is the SuccessResult definition
type SuccessResult struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// SuccessResultList is the SuccessResultList definition
type SuccessResultList struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// UtilizationHealth is the UtilizationHealth definition
type UtilizationHealth struct {
	Ghz24  int `json:"Ghz24,omitempty"`  //
	Ghz50  int `json:"Ghz50,omitempty"`  //
	Radio0 int `json:"radio0,omitempty"` //
	Radio1 int `json:"radio1,omitempty"` //
}

// VlanListResult is the VlanListResult definition
type VlanListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// WirelessInfoResult is the WirelessInfoResult definition
type WirelessInfoResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// AddDevice addDevice
/* Adds the device with given credential
 */
// func (s *DevicesService) AddDevice(addDeviceRequest *AddDeviceRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/network-device"

// 	response, err := RestyClient.R().
// 		SetBody(addDeviceRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// DeleteDeviceByIDQueryParams defines the query parameters for this request
type DeleteDeviceByIDQueryParams struct {
	IsForceDelete bool `url:"isForceDelete,omitempty"` // isForceDelete
}

// DeleteDeviceByID deleteDeviceById
/* Deletes the network device for the given ID
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
	Limit      int    `url:"limit,omitempty"`      // Max int of device entries in the response (default to 50.  Max at 1000)
	Offset     int    `url:"offset,omitempty"`     // The offset of the first device in the returned data
}

// Devices devices
/* Intent API for accessing DNA Assurance Device object for generating reports, creating dashboards or creating additional value added services.
@param deviceRole The device role (One of CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, AP)
@param siteID Assurance site UUID value
@param health The device overall health (One of POOR, FAIR, GOOD)
@param startTime UTC epoch time in milliseconds
@param endTime UTC epoch time in miliseconds
@param limit Max int of device entries in the response (default to 50.  Max at 1000)
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
// func (s *DevicesService) ExportDeviceList(exportDeviceListRequest *ExportDeviceListRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/network-device/file"

// 	response, err := RestyClient.R().
// 		SetBody(exportDeviceListRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

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
func (s *DevicesService) GetAllInterfaces(getAllInterfacesQueryParams *GetAllInterfacesQueryParams) (*DeviceIfListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface"

	queryString, _ := query.Values(getAllInterfacesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeviceIfListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfListResult)
	return result, response, err

}

// GetDeviceByID getDeviceByID
/* Returns the network device details for the given device ID
@param id Device ID
*/
func (s *DevicesService) GetDeviceByID(id string) (*NetworkDeviceResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&NetworkDeviceResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceResult)
	return result, response, err

}

// GetDeviceBySerialNumber getDeviceBySerialNumber
/* Returns the network device with given serial int
@param serialNumber Device serial int
*/
func (s *DevicesService) GetDeviceBySerialNumber(serialNumber string) (*NetworkDeviceResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/serial-int/{serialNumber}"
	path = strings.Replace(path, "{"+"serialNumber"+"}", fmt.Sprintf("%v", serialNumber), -1)

	response, err := RestyClient.R().
		SetResult(&NetworkDeviceResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceResult)
	return result, response, err

}

// GetDeviceConfigByID getDeviceConfigById
/* Returns the device config by specified device ID
@param networkDeviceID networkDeviceId
*/
func (s *DevicesService) GetDeviceConfigByID(networkDeviceID string) (*SuccessResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{networkDeviceID}/config"
	path = strings.Replace(path, "{"+"networkDeviceID"+"}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := RestyClient.R().
		SetResult(&SuccessResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*SuccessResult)
	return result, response, err

}

// GetDeviceConfigCount getDeviceConfigCount
/* Returns the count of device configs
 */
func (s *DevicesService) GetDeviceConfigCount() (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/config/count"

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetDeviceConfigForAllDevices getDeviceConfigForAllDevices
/* Returns the config for all devices
 */
func (s *DevicesService) GetDeviceConfigForAllDevices() (*RawCliInfoNIOListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/config"

	response, err := RestyClient.R().
		SetResult(&RawCliInfoNIOListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*RawCliInfoNIOListResult)
	return result, response, err

}

// GetDeviceCount getDeviceCount
/* Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and location name
 */
func (s *DevicesService) GetDeviceCount() (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/count"

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
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
func (s *DevicesService) GetDeviceInterfaceCount() (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/count"

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetDeviceInterfaceCountByDeviceID getDeviceInterfaceCountByDeviceId
/* Returns the interface count for the given device
@param deviceID Device ID
*/
func (s *DevicesService) GetDeviceInterfaceCountByDeviceID(deviceID string) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceID}/count"
	path = strings.Replace(path, "{"+"deviceID"+"}", fmt.Sprintf("%v", deviceID), -1)

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
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
func (s *DevicesService) GetDeviceInterfaceVLANs(id string, getDeviceInterfaceVLANsQueryParams *GetDeviceInterfaceVLANsQueryParams) (*VlanListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/vlan"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getDeviceInterfaceVLANsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&VlanListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*VlanListResult)
	return result, response, err

}

// GetDeviceInterfacesBySpecifiedRange getDeviceInterfacesBySpecifiedRange
/* Returns the list of interfaces for the device for the specified range
@param deviceID Device ID
@param startIndex Start index
@param recordsToReturn Number of records to return
*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRange(deviceID string, startIndex int, recordsToReturn int) (*DeviceIfListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceID}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"deviceID"+"}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := RestyClient.R().
		SetResult(&DeviceIfListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfListResult)
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
	Modulartint             []string `url:"module+partint,omitempty"`             // modulePartNumber
	Modulperationstatecode  []string `url:"module+operationstatecode,omitempty"`  // moduleOperationStateCode
	ID                      string   `url:"id,omitempty"`                         // Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
}

// GetDeviceList getDeviceList
/* Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc. You can use the .* in any value to conduct a wildcard search.
For example, to find all hostnames beginning with myhost in the IP address range 192.25.18.n, issue the following request:
GET /dna/intent/api/v1/network-device?hostname=myhost.*&managementIPAddress=192.25.18..*

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
@param module+partint modulePartNumber
@param module+operationstatecode moduleOperationStateCode
@param id Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
*/
func (s *DevicesService) GetDeviceList(getDeviceListQueryParams *GetDeviceListQueryParams) (*NetworkDeviceListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device"

	queryString, _ := query.Values(getDeviceListQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&NetworkDeviceListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceListResult)
	return result, response, err

}

// GetDeviceSummary getDeviceSummary
/* Returns brief summary of device info such as hostname, management IP address for the given device ID
@param id Device ID
*/
func (s *DevicesService) GetDeviceSummary(id string) (*NetworkDeviceBriefNIOResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/brief"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&NetworkDeviceBriefNIOResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceBriefNIOResult)
	return result, response, err

}

// GetFunctionalCapabilityByID getFunctionalCapabilityById
/* Returns functional capability with given ID
@param id Functional Capability UUID
*/
func (s *DevicesService) GetFunctionalCapabilityByID(id string) (*FunctionalCapabilityResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/functional-capability/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&FunctionalCapabilityResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*FunctionalCapabilityResult)
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
func (s *DevicesService) GetFunctionalCapabilityForDevices(getFunctionalCapabilityForDevicesQueryParams *GetFunctionalCapabilityForDevicesQueryParams) (*FunctionalCapabilityListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/functional-capability"

	queryString, _ := query.Values(getFunctionalCapabilityForDevicesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&FunctionalCapabilityListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*FunctionalCapabilityListResult)
	return result, response, err

}

// GetISISInterfaces getISISInterfaces
/* Returns the interfaces that has ISIS enabled
 */
func (s *DevicesService) GetISISInterfaces() (*DeviceIfListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/isis"

	response, err := RestyClient.R().
		SetResult(&DeviceIfListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfListResult)
	return result, response, err

}

// GetInterfaceByIP getInterfaceByIP
/* Returns list of interfaces for specified device management IP address
@param ipAddress IP address of the interface
*/
func (s *DevicesService) GetInterfaceByIP(ipAddress string) (*DeviceIfListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := RestyClient.R().
		SetResult(&DeviceIfListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfListResult)
	return result, response, err

}

// GetInterfaceByID getInterfaceById
/* Returns the interface for the given interface ID
@param id Interface ID
*/
func (s *DevicesService) GetInterfaceByID(id string) (*DeviceIfResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DeviceIfResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfResult)
	return result, response, err

}

// GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams defines the query parameters for this request
type GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams struct {
	Name string `url:"name,omitempty"` // Interface name
}

// GetInterfaceDetailsByDeviceIDAndInterfaceName getInterfaceDetailsByDeviceIdAndInterfaceName
/* Returns interface by specified device ID and interface name
@param deviceID Device ID
@param name Interface name
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID string, getInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams) (*DeviceIfResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceID}/interface-name"
	path = strings.Replace(path, "{"+"deviceID"+"}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(getInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeviceIfResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfResult)
	return result, response, err

}

// GetInterfaceInfoByID getInterfaceInfoById
/* Returns list of interfaces by specified device
@param deviceID Device ID
*/
func (s *DevicesService) GetInterfaceInfoByID(deviceID string) (*DeviceIfListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/network-device/{deviceID}"
	path = strings.Replace(path, "{"+"deviceID"+"}", fmt.Sprintf("%v", deviceID), -1)

	response, err := RestyClient.R().
		SetResult(&DeviceIfListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfListResult)
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
func (s *DevicesService) GetModuleCount(getModuleCountQueryParams *GetModuleCountQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/module/count"

	queryString, _ := query.Values(getModuleCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetModuleInfoByID getModuleInfoById
/* Returns Module info by id
@param id id
*/
func (s *DevicesService) GetModuleInfoByID(id string) (*ModuleResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/module/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&ModuleResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*ModuleResult)
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
func (s *DevicesService) GetModules(getModulesQueryParams *GetModulesQueryParams) (*ModuleListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/module"

	queryString, _ := query.Values(getModulesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&ModuleListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*ModuleListResult)
	return result, response, err

}

// GetNetworkDeviceByIP getNetworkDeviceByIP
/* Returns the network device by specified IP address
@param ipAddress Device IP address
*/
func (s *DevicesService) GetNetworkDeviceByIP(ipAddress string) (*NetworkDeviceResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{"+"ipAddress"+"}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := RestyClient.R().
		SetResult(&NetworkDeviceResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceResult)
	return result, response, err

}

// GetNetworkDeviceByPaginationRange getNetworkDeviceByPaginationRange
/* Returns the list of network devices for the given pagination range
@param startIndex Start index
@param recordsToReturn Number of records to return
*/
func (s *DevicesService) GetNetworkDeviceByPaginationRange(startIndex int, recordsToReturn int) (*NetworkDeviceListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := RestyClient.R().
		SetResult(&NetworkDeviceListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceListResult)
	return result, response, err

}

// GetOSPFInterfaces getOSPFInterfaces
/* Returns the interfaces that has OSPF enabled
 */
func (s *DevicesService) GetOSPFInterfaces() (*DeviceIfListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/interface/ospf"

	response, err := RestyClient.R().
		SetResult(&DeviceIfListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceIfListResult)
	return result, response, err

}

// GetOrganizationListForMeraki getOrganizationListForMeraki
/* Returns list of organizations for meraki dashboard
@param id id
*/
func (s *DevicesService) GetOrganizationListForMeraki(id string) (*SuccessResultList, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/meraki-organization"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&SuccessResultList{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*SuccessResultList)
	return result, response, err

}

// GetPollingIntervalByID getPollingIntervalById
/* Returns polling interval by device id
@param id Device ID
*/
func (s *DevicesService) GetPollingIntervalByID(id string) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/collection-schedule"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetPollingIntervalForAllDevices getPollingIntervalForAllDevices
/* Returns polling interval of all devices
 */
func (s *DevicesService) GetPollingIntervalForAllDevices() (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/collection-schedule/global"

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetWirelessLanControllerDetailsByID getWirelessLanControllerDetailsById
/* Returns the wireless lan controller info with given device ID
@param id Device ID
*/
func (s *DevicesService) GetWirelessLanControllerDetailsByID(id string) (*WirelessInfoResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/{id}/wireless-info"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&WirelessInfoResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*WirelessInfoResult)
	return result, response, err

}

// RegisterDeviceForWSAQueryParams defines the query parameters for this request
type RegisterDeviceForWSAQueryParams struct {
	SerialNumber string `url:"serialNumber,omitempty"` // Serial int of the device
	Macaddress   string `url:"macaddress,omitempty"`   // Mac addres of the device
}

// RegisterDeviceForWSA registerDeviceForWSA
/* Registers a device for WSA notification
@param serialNumber Serial int of the device
@param macaddress Mac addres of the device
*/
func (s *DevicesService) RegisterDeviceForWSA(registerDeviceForWSAQueryParams *RegisterDeviceForWSAQueryParams) (*RegisterNetworkDeviceResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/tenantinfo/macaddress"

	queryString, _ := query.Values(registerDeviceForWSAQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&RegisterNetworkDeviceResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*RegisterNetworkDeviceResult)
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
func (s *DevicesService) RetrievesAllNetworkDevices(retrievesAllNetworkDevicesQueryParams *RetrievesAllNetworkDevicesQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/network-device/autocomplete"

	queryString, _ := query.Values(retrievesAllNetworkDevicesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// SyncDevices syncDevices
/* Sync the devices provided as input
 */
// func (s *DevicesService) SyncDevices(syncDevicesRequest *SyncDevicesRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/network-device"

// 	response, err := RestyClient.R().
// 		SetBody(syncDevicesRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// SyncNetworkDevicesQueryParams defines the query parameters for this request
type SyncNetworkDevicesQueryParams struct {
	ForceSync bool `url:"forceSync,omitempty"` // forceSync
}

// SyncNetworkDevices syncNetworkDevices
/* Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result can be seen in the child task of each device
@param forceSync forceSync
*/
// func (s *DevicesService) SyncNetworkDevices(syncNetworkDevicesQueryParams *SyncNetworkDevicesQueryParams, syncNetworkDevicesRequest *SyncNetworkDevicesRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/network-device/sync"

// 	queryString, _ := query.Values(syncNetworkDevicesQueryParams)

// 	response, err := RestyClient.R().
// 		SetQueryString(queryString.Encode()).
// 		SetBody(syncNetworkDevicesRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

//}

// UpdateDeviceRole updateDeviceRole
/* Updates the role of the device as access, core, distribution, border router
 */
// func (s *DevicesService) UpdateDeviceRole(updateDeviceRoleRequest *UpdateDeviceRoleRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/network-device/brief"

// 	response, err := RestyClient.R().
// 		SetBody(updateDeviceRoleRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }
