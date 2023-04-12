package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LicensesService service

type DeviceCountDetails2QueryParams struct {
	DeviceType         string `url:"device_type,omitempty"`          //Type of device
	RegistrationStatus string `url:"registration_status,omitempty"`  //Smart license registration status of device
	DnaLevel           string `url:"dna_level,omitempty"`            //Device Cisco DNA License Level
	VirtualAccountName string `url:"virtual_account_name,omitempty"` //Virtual account name
	SmartAccountID     string `url:"smart_account_id,omitempty"`     //Smart account id
}
type DeviceLicenseSummary2QueryParams struct {
	PageNumber         float64 `url:"page_number,omitempty"`          //Page number of response
	Order              string  `url:"order,omitempty"`                //Sorting order
	SortBy             string  `url:"sort_by,omitempty"`              //Sort result by field
	DnaLevel           string  `url:"dna_level,omitempty"`            //Device Cisco DNA license level
	DeviceType         string  `url:"device_type,omitempty"`          //Type of device
	Limit              float64 `url:"limit,omitempty"`                //Limit
	RegistrationStatus string  `url:"registration_status,omitempty"`  //Smart license registration status of device
	VirtualAccountName string  `url:"virtual_account_name,omitempty"` //Name of virtual account
	SmartAccountID     float64 `url:"smart_account_id,omitempty"`     //Id of smart account
	DeviceUUID         string  `url:"device_uuid,omitempty"`          //Id of device
}
type LicenseTermDetails2QueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}
type LicenseUsageDetails2QueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}

type ResponseLicensesDeviceCountDetails2 struct {
	Response *int   `json:"response,omitempty"` // Total number of managed device
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseLicensesDeviceLicenseSummary2 struct {
	Response *[]ResponseLicensesDeviceLicenseSummary2Response `json:"response,omitempty"` //
}
type ResponseLicensesDeviceLicenseSummary2Response struct {
	AuthorizationStatus              string `json:"authorization_status,omitempty"`                  // Smart license authorization status of device
	LastUpdatedTime                  string `json:"last_updated_time,omitempty"`                     // Time when license information was collected from device
	IsPerformanceAllowed             *bool  `json:"is_performance_allowed,omitempty"`                // Is performance license available
	SleAuthCode                      string `json:"sle_auth_code,omitempty"`                         // SLE Authcode installed or not installed
	ThroughputLevel                  string `json:"throughput_level,omitempty"`                      // Throughput level on device
	HsecStatus                       string `json:"hsec_status,omitempty"`                           // HSEC status of the device
	DeviceUUID                       string `json:"device_uuid,omitempty"`                           // Id of device
	Site                             string `json:"site,omitempty"`                                  // Site of device
	TotalAccessPointCount            *int   `json:"total_access_point_count,omitempty"`              // Total number of Access Points connected
	Model                            string `json:"model,omitempty"`                                 // Model of device
	IsWirelessCapable                *bool  `json:"is_wireless_capable,omitempty"`                   // Is device wireless capable
	RegistrationStatus               string `json:"registration_status,omitempty"`                   // Smart license registration status of device
	SleState                         string `json:"sle_state,omitempty"`                             // SLE state on device
	PerformanceLicense               string `json:"performance_license,omitempty"`                   // Is performance license enabled
	LicenseMode                      string `json:"license_mode,omitempty"`                          // Mode of license
	IsLicenseExpired                 *bool  `json:"is_license_expired,omitempty"`                    // Is device license expired
	SoftwareVersion                  string `json:"software_version,omitempty"`                      // Software image version of device
	ReservationStatus                string `json:"reservation_status,omitempty"`                    // Smart license reservation status
	IsWireless                       *bool  `json:"is_wireless,omitempty"`                           // Is device wireless controller
	NetworkLicense                   string `json:"network_license,omitempty"`                       // Device Network license level
	EvaluationLicenseExpiry          string `json:"evaluation_license_expiry,omitempty"`             // Evaluation period expiry date
	WirelessCapableNetworkLicense    string `json:"wireless_capable_network_license,omitempty"`      // Wireless Cisco Network license value
	DeviceName                       string `json:"device_name,omitempty"`                           // Name of device
	DeviceType                       string `json:"device_type,omitempty"`                           // Type of device
	DnaLevel                         string `json:"dna_level,omitempty"`                             // Device Cisco DNA license level
	VirtualAccountName               string `json:"virtual_account_name,omitempty"`                  // Name of virtual account
	LastSuccessfulRumUsageUploadTime string `json:"last_successful_rum_usage_upload_time,omitempty"` // Last successful rum usage upload time
	IPAddress                        string `json:"ip_address,omitempty"`                            // IP address of device
	WirelessCapableDnaLicense        string `json:"wireless_capable_dna_license,omitempty"`          // Wireless Cisco DNA license value
	MacAddress                       string `json:"mac_address,omitempty"`                           // MAC address of device
	CustomerTag1                     string `json:"customer_tag1,omitempty"`                         // Customer Tag1 set on device
	CustomerTag2                     string `json:"customer_tag2,omitempty"`                         // Customer Tag2 set on device
	CustomerTag3                     string `json:"customer_tag3,omitempty"`                         // Customer Tag3 set on device
	CustomerTag4                     string `json:"customer_tag4,omitempty"`                         // Customer Tag4 set on device
	SmartAccountName                 string `json:"smart_account_name,omitempty"`                    // Name of smart account
}
type ResponseLicensesDeviceLicenseDetails2 struct {
	Response *[]ResponseLicensesDeviceLicenseDetails2Response `json:"response,omitempty"` //
}
type ResponseLicensesDeviceLicenseDetails2Response struct {
	DeviceUUID              string                                                         `json:"device_uuid,omitempty"`               // Id of device
	Site                    string                                                         `json:"site,omitempty"`                      // Site of device
	Model                   string                                                         `json:"model,omitempty"`                     // Model of device
	LicenseMode             string                                                         `json:"license_mode,omitempty"`              // Mode of license
	IsLicenseExpired        *bool                                                          `json:"is_license_expired,omitempty"`        // Is device license expired
	SoftwareVersion         string                                                         `json:"software_version,omitempty"`          // Software image version of device
	NetworkLicense          string                                                         `json:"network_license,omitempty"`           // Device network license level
	EvaluationLicenseExpiry string                                                         `json:"evaluation_license_expiry,omitempty"` // Evaluation period expiry date
	DeviceName              string                                                         `json:"device_name,omitempty"`               // Name of device
	DeviceType              string                                                         `json:"device_type,omitempty"`               // Type of device
	DnaLevel                string                                                         `json:"dna_level,omitempty"`                 // Device Cisco DNA license level
	VirtualAccountName      string                                                         `json:"virtual_account_name,omitempty"`      // Name of virtual account
	IPAddress               string                                                         `json:"ip_address,omitempty"`                // IP address of device
	MacAddress              string                                                         `json:"mac_address,omitempty"`               // MAC address of device
	SntcStatus              string                                                         `json:"sntc_status,omitempty"`               // Valid if device is covered under license contract and invalid if not covered, otherwise unknown.
	FeatureLicense          []string                                                       `json:"feature_license,omitempty"`           // Name of feature licenses
	HasSupCards             *bool                                                          `json:"has_sup_cards,omitempty"`             // Whether device has supervisor cards
	Udi                     string                                                         `json:"udi,omitempty"`                       // Unique Device Identifier
	StackedDevices          *[]ResponseLicensesDeviceLicenseDetails2ResponseStackedDevices `json:"stacked_devices,omitempty"`           //
	IsStackedDevice         *bool                                                          `json:"is_stacked_device,omitempty"`         // Is Stacked Device
	AccessPoints            *[]ResponseLicensesDeviceLicenseDetails2ResponseAccessPoints   `json:"access_points,omitempty"`             //
	ChassisDetails          *ResponseLicensesDeviceLicenseDetails2ResponseChassisDetails   `json:"chassis_details,omitempty"`           //
}
type ResponseLicensesDeviceLicenseDetails2ResponseStackedDevices struct {
	MacAddress   string `json:"mac_address,omitempty"`   // Stack mac address
	ID           string `json:"id,omitempty"`            // Id
	Role         string `json:"role,omitempty"`          // Chassis role
	SerialNumber string `json:"serial_number,omitempty"` // Chassis serial number
}
type ResponseLicensesDeviceLicenseDetails2ResponseAccessPoints struct {
	ApType string `json:"ap_type,omitempty"` // Type of access point
	Count  string `json:"count,omitempty"`   // Number of access point
}
type ResponseLicensesDeviceLicenseDetails2ResponseChassisDetails struct {
	BoardSerialNumber string                                                                        `json:"board_serial_number,omitempty"` // Board serial number
	Modules           *[]ResponseLicensesDeviceLicenseDetails2ResponseChassisDetailsModules         `json:"modules,omitempty"`             //
	SupervisorCards   *[]ResponseLicensesDeviceLicenseDetails2ResponseChassisDetailsSupervisorCards `json:"supervisor_cards,omitempty"`    //
	Port              *int                                                                          `json:"port,omitempty"`                // Number of port
}
type ResponseLicensesDeviceLicenseDetails2ResponseChassisDetailsModules struct {
	ModuleType   string `json:"module_type,omitempty"`   // Type of module
	ModuleName   string `json:"module_name,omitempty"`   // Name of module
	SerialNumber string `json:"serial_number,omitempty"` // Serial number of module
	ID           string `json:"id,omitempty"`            // Id of module
}
type ResponseLicensesDeviceLicenseDetails2ResponseChassisDetailsSupervisorCards struct {
	SerialNumber       string `json:"serial_number,omitempty"`        // Serial number of supervisor card
	SupervisorCardType string `json:"supervisor_card_type,omitempty"` // Type of supervisor card
	Status             string `json:"status,omitempty"`               // Status of supervisor card
}
type ResponseLicensesDeviceDeregistration2 struct {
	Response *ResponseLicensesDeviceDeregistration2Response `json:"response,omitempty"` //
}
type ResponseLicensesDeviceDeregistration2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesDeviceRegistration2 struct {
	Response *ResponseLicensesDeviceRegistration2Response `json:"response,omitempty"` //
}
type ResponseLicensesDeviceRegistration2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesChangeVirtualAccount2 struct {
	Response *ResponseLicensesChangeVirtualAccount2Response `json:"response,omitempty"` //
}
type ResponseLicensesChangeVirtualAccount2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesVirtualAccountDetails2 struct {
	SmartAccountID        string                                                         `json:"smart_account_id,omitempty"`        // Id of smart account
	SmartAccountName      string                                                         `json:"smart_account_name,omitempty"`      // Name of smart account
	VirtualAccountDetails *[]ResponseLicensesVirtualAccountDetails2VirtualAccountDetails `json:"virtual_account_details,omitempty"` //
}
type ResponseLicensesVirtualAccountDetails2VirtualAccountDetails struct {
	VirtualAccountID   string `json:"virtual_account_id,omitempty"`   // Id of virtual account
	VirtualAccountName string `json:"virtual_account_name,omitempty"` // Name of virtual account
}
type ResponseLicensesSmartAccountDetails struct {
	Response *[]ResponseLicensesSmartAccountDetailsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseLicensesSmartAccountDetailsResponse struct {
	Name                 string `json:"name,omitempty"`                    // Name of smart account
	ID                   string `json:"id,omitempty"`                      // Id of smart account
	Domain               string `json:"domain,omitempty"`                  // Domain of smart account
	IsActiveSmartAccount *bool  `json:"is_active_smart_account,omitempty"` // Is active smart account
}
type ResponseLicensesLicenseTermDetails2 struct {
	LicenseDetails *[]ResponseLicensesLicenseTermDetails2LicenseDetails `json:"license_details,omitempty"` //
}
type ResponseLicensesLicenseTermDetails2LicenseDetails struct {
	Model                    string `json:"model,omitempty"`                       // Model of device
	VirtualAccountName       string `json:"virtual_account_name,omitempty"`        // Name of virtual account
	LicenseTermStartDate     string `json:"license_term_start_date,omitempty"`     // Start date of license term
	LicenseTermEndDate       string `json:"license_term_end_date,omitempty"`       // End date of license term
	DnaLevel                 string `json:"dna_level,omitempty"`                   // Cisco DNA license level
	PurchasedDnaLicenseCount string `json:"purchased_dna_license_count,omitempty"` // Number of purchased DNA licenses
	IsLicenseExpired         string `json:"is_license_expired,omitempty"`          // Is license expired
}
type ResponseLicensesLicenseUsageDetails2 struct {
	PurchasedDnaLicense     *ResponseLicensesLicenseUsageDetails2PurchasedDnaLicense     `json:"purchased_dna_license,omitempty"`     //
	PurchasedNetworkLicense *ResponseLicensesLicenseUsageDetails2PurchasedNetworkLicense `json:"purchased_network_license,omitempty"` //
	UsedDnaLicense          *ResponseLicensesLicenseUsageDetails2UsedDnaLicense          `json:"used_dna_license,omitempty"`          //
	UsedNetworkLicense      *ResponseLicensesLicenseUsageDetails2UsedNetworkLicense      `json:"used_network_license,omitempty"`      //
}
type ResponseLicensesLicenseUsageDetails2PurchasedDnaLicense struct {
	TotalLicenseCount  *int                                                                         `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetails2PurchasedDnaLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetails2PurchasedDnaLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetails2PurchasedNetworkLicense struct {
	TotalLicenseCount  *int                                                                             `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetails2PurchasedNetworkLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetails2PurchasedNetworkLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetails2UsedDnaLicense struct {
	TotalLicenseCount  *int                                                                    `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetails2UsedDnaLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetails2UsedDnaLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetails2UsedNetworkLicense struct {
	TotalLicenseCount  *int                                                                        `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetails2UsedNetworkLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetails2UsedNetworkLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type RequestLicensesDeviceDeregistration2 struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesDeviceRegistration2 struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesChangeVirtualAccount2 struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}

//DeviceCountDetails2 Device Count Details - 949a-6983-4c38-9d59
/* Get total number of managed device(s).


@param DeviceCountDetails2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-count-details2
*/
func (s *LicensesService) DeviceCountDetails2(DeviceCountDetails2QueryParams *DeviceCountDetails2QueryParams) (*ResponseLicensesDeviceCountDetails2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/count"

	queryString, _ := query.Values(DeviceCountDetails2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesDeviceCountDetails2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceCountDetails2")
	}

	result := response.Result().(*ResponseLicensesDeviceCountDetails2)
	return result, response, err

}

//DeviceLicenseSummary2 Device License Summary - c491-4937-44fa-a216
/* Show license summary of device(s).


@param DeviceLicenseSummary2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-license-summary2
*/
func (s *LicensesService) DeviceLicenseSummary2(DeviceLicenseSummary2QueryParams *DeviceLicenseSummary2QueryParams) (*ResponseLicensesDeviceLicenseSummary2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/summary"

	queryString, _ := query.Values(DeviceLicenseSummary2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesDeviceLicenseSummary2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceLicenseSummary2")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseSummary2)
	return result, response, err

}

//DeviceLicenseDetails2 Device License Details - dca1-1bc2-4e7b-8c5d
/* Get detailed license information of a device.


@param deviceuuid device_uuid path parameter. Id of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-license-details2
*/
func (s *LicensesService) DeviceLicenseDetails2(deviceuuid string) (*ResponseLicensesDeviceLicenseDetails2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/{device_uuid}/details"
	path = strings.Replace(path, "{device_uuid}", fmt.Sprintf("%v", deviceuuid), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesDeviceLicenseDetails2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceLicenseDetails2")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseDetails2)
	return result, response, err

}

//VirtualAccountDetails2 Virtual Account Details - b681-09f6-4bb9-b3b3
/* Get virtual account details of a smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account


Documentation Link: https://developer.cisco.com/docs/dna-center/#!virtual-account-details2
*/
func (s *LicensesService) VirtualAccountDetails2(smartaccountTypeID string) (*ResponseLicensesVirtualAccountDetails2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/{smart_account_id}/virtualAccounts"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesVirtualAccountDetails2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation VirtualAccountDetails2")
	}

	result := response.Result().(*ResponseLicensesVirtualAccountDetails2)
	return result, response, err

}

//SmartAccountDetails Smart Account Details - c181-b8db-4c89-adf0
/* Get detail of all smart accounts.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!smart-account-details
*/
func (s *LicensesService) SmartAccountDetails() (*ResponseLicensesSmartAccountDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccounts"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesSmartAccountDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SmartAccountDetails")
	}

	result := response.Result().(*ResponseLicensesSmartAccountDetails)
	return result, response, err

}

//LicenseTermDetails2 License Term Details - c8a9-cb5e-4078-a16f
/* Get license term details.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseTermDetails2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!license-term-details2
*/
func (s *LicensesService) LicenseTermDetails2(smartaccountTypeID string, virtualaccountname string, LicenseTermDetails2QueryParams *LicenseTermDetails2QueryParams) (*ResponseLicensesLicenseTermDetails2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/term/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	queryString, _ := query.Values(LicenseTermDetails2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesLicenseTermDetails2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LicenseTermDetails2")
	}

	result := response.Result().(*ResponseLicensesLicenseTermDetails2)
	return result, response, err

}

//LicenseUsageDetails2 License Usage Details - 418a-6b43-4e29-bfe5
/* Get count of purchased and in use Cisco DNA and Network licenses.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseUsageDetails2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!license-usage-details2
*/
func (s *LicensesService) LicenseUsageDetails2(smartaccountTypeID string, virtualaccountname string, LicenseUsageDetails2QueryParams *LicenseUsageDetails2QueryParams) (*ResponseLicensesLicenseUsageDetails2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/usage/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	queryString, _ := query.Values(LicenseUsageDetails2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesLicenseUsageDetails2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LicenseUsageDetails2")
	}

	result := response.Result().(*ResponseLicensesLicenseUsageDetails2)
	return result, response, err

}

//ChangeVirtualAccount2 Change Virtual Account - ba89-1873-4089-9a37
/* Transfer device(s) from one virtual account to another within same smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of target virtual account


Documentation Link: https://developer.cisco.com/docs/dna-center/#!change-virtual-account2
*/
func (s *LicensesService) ChangeVirtualAccount2(smartaccountTypeID string, virtualaccountname string, requestLicensesChangeVirtualAccount2 *RequestLicensesChangeVirtualAccount2) (*ResponseLicensesChangeVirtualAccount2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}/device/transfer"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesChangeVirtualAccount2).
		SetResult(&ResponseLicensesChangeVirtualAccount2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ChangeVirtualAccount2")
	}

	result := response.Result().(*ResponseLicensesChangeVirtualAccount2)
	return result, response, err

}

//DeviceDeregistration2 Device Deregistration - 8c82-dad4-49ba-b8eb
/* Deregister device(s) from CSSM(Cisco Smart Software Manager).


 */
func (s *LicensesService) DeviceDeregistration2(requestLicensesDeviceDeregistration2 *RequestLicensesDeviceDeregistration2) (*ResponseLicensesDeviceDeregistration2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/virtualAccount/deregister"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesDeviceDeregistration2).
		SetResult(&ResponseLicensesDeviceDeregistration2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceDeregistration2")
	}

	result := response.Result().(*ResponseLicensesDeviceDeregistration2)
	return result, response, err

}

//DeviceRegistration2 Device Registration - a08b-eae5-47fb-95e3
/* Register device(s) in CSSM(Cisco Smart Software Manager).


@param virtualaccountname virtual_account_name path parameter. Name of virtual account

*/
func (s *LicensesService) DeviceRegistration2(virtualaccountname string, requestLicensesDeviceRegistration2 *RequestLicensesDeviceRegistration2) (*ResponseLicensesDeviceRegistration2, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/virtualAccount/{virtual_account_name}/register"
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesDeviceRegistration2).
		SetResult(&ResponseLicensesDeviceRegistration2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceRegistration2")
	}

	result := response.Result().(*ResponseLicensesDeviceRegistration2)
	return result, response, err

}
