package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LicensesService service

type DeviceCountDetailsQueryParams struct {
	DeviceType         string `url:"device_type,omitempty"`          //Type of device
	RegistrationStatus string `url:"registration_status,omitempty"`  //Smart license registration status of device
	DnaLevel           string `url:"dna_level,omitempty"`            //Device Cisco DNA license level
	VirtualAccountName string `url:"virtual_account_name,omitempty"` //Name of virtual account
	SmartAccountID     string `url:"smart_account_id,omitempty"`     //Id of smart account
}
type DeviceLicenseSummaryQueryParams struct {
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
type LicenseTermDetailsQueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}
type LicenseUsageDetailsQueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}

type ResponseLicensesDeviceCountDetails struct {
	Response *int   `json:"response,omitempty"` // Total number of managed device
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseLicensesDeviceLicenseSummary struct {
	Version  string                                          `json:"version,omitempty"`  // Version
	Response *[]ResponseLicensesDeviceLicenseSummaryResponse `json:"response,omitempty"` //
}
type ResponseLicensesDeviceLicenseSummaryResponse struct {
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
type ResponseLicensesDeviceLicenseDetails struct {
	DeviceUUID              string                                              `json:"device_uuid,omitempty"`               // Id of device
	Site                    string                                              `json:"site,omitempty"`                      // Site of device
	Model                   string                                              `json:"model,omitempty"`                     // Model of device
	LicenseMode             string                                              `json:"license_mode,omitempty"`              // Mode of license
	IsLicenseExpired        *bool                                               `json:"is_license_expired,omitempty"`        // Is device license expired
	SoftwareVersion         string                                              `json:"software_version,omitempty"`          // Software image version of device
	NetworkLicense          string                                              `json:"network_license,omitempty"`           // Device network license level
	EvaluationLicenseExpiry string                                              `json:"evaluation_license_expiry,omitempty"` // Evaluation period expiry date
	DeviceName              string                                              `json:"device_name,omitempty"`               // Name of device
	DeviceType              string                                              `json:"device_type,omitempty"`               // Type of device
	DnaLevel                string                                              `json:"dna_level,omitempty"`                 // Device Cisco DNA license level
	VirtualAccountName      string                                              `json:"virtual_account_name,omitempty"`      // Name of virtual account
	IPAddress               string                                              `json:"ip_address,omitempty"`                // IP address of device
	MacAddress              string                                              `json:"mac_address,omitempty"`               // MAC address of device
	SntcStatus              string                                              `json:"sntc_status,omitempty"`               // Valid if device is covered under license contract and invalid if not covered, otherwise unknown.
	FeatureLicense          []string                                            `json:"feature_license,omitempty"`           // Name of feature licenses
	HasSupCards             *bool                                               `json:"has_sup_cards,omitempty"`             // Whether device has supervisor cards
	Udi                     string                                              `json:"udi,omitempty"`                       // Unique Device Identifier
	StackedDevices          *[][]string                                         `json:"stacked_devices,omitempty"`           // Stacked device details
	IsStackedDevice         *bool                                               `json:"is_stacked_device,omitempty"`         // Is Stacked Device
	AccessPoints            *[][]string                                         `json:"access_points,omitempty"`             // Access point details
	ChassisDetails          *ResponseLicensesDeviceLicenseDetailsChassisDetails `json:"chassis_details,omitempty"`           //
}
type ResponseLicensesDeviceLicenseDetailsChassisDetails struct {
	BoardSerialNumber string      `json:"board_serial_number,omitempty"` // Board serial number
	Modules           *[][]string `json:"modules,omitempty"`             // Module details
	SupervisorCards   *[][]string `json:"supervisor_cards,omitempty"`    // Supervisor card details
	Port              *int        `json:"port,omitempty"`                // Number of port
}
type ResponseLicensesDeviceDeregistration struct {
	Response *ResponseLicensesDeviceDeregistrationResponse `json:"response,omitempty"` //
}
type ResponseLicensesDeviceDeregistrationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesDeviceRegistration struct {
	Response *ResponseLicensesDeviceRegistrationResponse `json:"response,omitempty"` //
}
type ResponseLicensesDeviceRegistrationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesChangeVirtualAccount struct {
	Response *ResponseLicensesChangeVirtualAccountResponse `json:"response,omitempty"` //
}
type ResponseLicensesChangeVirtualAccountResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id of process
	URL    string `json:"url,omitempty"`    // Task URL of process
}
type ResponseLicensesVirtualAccountDetails struct {
	SmartAccountID        string                                                        `json:"smart_account_id,omitempty"`        // Id of smart account
	SmartAccountName      string                                                        `json:"smart_account_name,omitempty"`      // Name of smart account
	VirtualAccountDetails *[]ResponseLicensesVirtualAccountDetailsVirtualAccountDetails `json:"virtual_account_details,omitempty"` //
}
type ResponseLicensesVirtualAccountDetailsVirtualAccountDetails struct {
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
type ResponseLicensesLicenseTermDetails struct {
	LicenseDetails *[]ResponseLicensesLicenseTermDetailsLicenseDetails `json:"license_details,omitempty"` //
}
type ResponseLicensesLicenseTermDetailsLicenseDetails struct {
	Model                    string `json:"model,omitempty"`                       // Model of device
	VirtualAccountName       string `json:"virtual_account_name,omitempty"`        // Name of virtual account
	LicenseTermStartDate     string `json:"license_term_start_date,omitempty"`     // Start date of license term
	LicenseTermEndDate       string `json:"license_term_end_date,omitempty"`       // End date of license term
	DnaLevel                 string `json:"dna_level,omitempty"`                   // Cisco DNA license level
	PurchasedDnaLicenseCount string `json:"purchased_dna_license_count,omitempty"` // Number of purchased DNA licenses
	IsLicenseExpired         string `json:"is_license_expired,omitempty"`          // Is license expired
}
type ResponseLicensesLicenseUsageDetails struct {
	PurchasedDnaLicense     *ResponseLicensesLicenseUsageDetailsPurchasedDnaLicense     `json:"purchased_dna_license,omitempty"`     //
	PurchasedNetworkLicense *ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicense `json:"purchased_network_license,omitempty"` //
	UsedDnaLicense          *ResponseLicensesLicenseUsageDetailsUsedDnaLicense          `json:"used_dna_license,omitempty"`          //
	UsedNetworkLicense      *ResponseLicensesLicenseUsageDetailsUsedNetworkLicense      `json:"used_network_license,omitempty"`      //
}
type ResponseLicensesLicenseUsageDetailsPurchasedDnaLicense struct {
	TotalLicenseCount  *int                                                                        `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsPurchasedDnaLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsPurchasedDnaLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicense struct {
	TotalLicenseCount  *int                                                                            `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsUsedDnaLicense struct {
	TotalLicenseCount  *int                                                                   `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsUsedDnaLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsUsedDnaLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsUsedNetworkLicense struct {
	TotalLicenseCount  *int                                                                       `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsUsedNetworkLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsUsedNetworkLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type RequestLicensesDeviceDeregistration struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesDeviceRegistration struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}
type RequestLicensesChangeVirtualAccount struct {
	DeviceUUIDs []string `json:"device_uuids,omitempty"` // Comma separated device ids
}

//DeviceCountDetails Device Count Details - 6581-a84e-477a-b2f9
/* Get total number of managed device(s).


@param DeviceCountDetailsQueryParams Filtering parameter
*/
func (s *LicensesService) DeviceCountDetails(DeviceCountDetailsQueryParams *DeviceCountDetailsQueryParams) (*ResponseLicensesDeviceCountDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/count"

	queryString, _ := query.Values(DeviceCountDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesDeviceCountDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceCountDetails")
	}

	result := response.Result().(*ResponseLicensesDeviceCountDetails)
	return result, response, err

}

//DeviceLicenseSummary Device License Summary - c491-4937-44fa-a216
/* Show license summary of device(s).


@param DeviceLicenseSummaryQueryParams Filtering parameter
*/
func (s *LicensesService) DeviceLicenseSummary(DeviceLicenseSummaryQueryParams *DeviceLicenseSummaryQueryParams) (*ResponseLicensesDeviceLicenseSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/summary"

	queryString, _ := query.Values(DeviceLicenseSummaryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesDeviceLicenseSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceLicenseSummary")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseSummary)
	return result, response, err

}

//DeviceLicenseDetails Device License Details - 2693-a968-4c3b-86cf
/* Get detailed license information of a device.


@param deviceuuid device_uuid path parameter. Id of device

*/
func (s *LicensesService) DeviceLicenseDetails(deviceuuid string) (*ResponseLicensesDeviceLicenseDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/device/{device_uuid}/details"
	path = strings.Replace(path, "{device_uuid}", fmt.Sprintf("%v", deviceuuid), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesDeviceLicenseDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceLicenseDetails")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseDetails)
	return result, response, err

}

//VirtualAccountDetails Virtual Account Details - b681-09f6-4bb9-b3b3
/* Get virtual account details of a smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

*/
func (s *LicensesService) VirtualAccountDetails(smartaccountTypeID string) (*ResponseLicensesVirtualAccountDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/{smart_account_id}/virtualAccounts"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesVirtualAccountDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation VirtualAccountDetails")
	}

	result := response.Result().(*ResponseLicensesVirtualAccountDetails)
	return result, response, err

}

//SmartAccountDetails Smart Account Details - c181-b8db-4c89-adf0
/* Get detail of all smart accounts.


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

//LicenseTermDetails License Term Details - c8a9-cb5e-4078-a16f
/* Get license term details.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseTermDetailsQueryParams Filtering parameter
*/
func (s *LicensesService) LicenseTermDetails(smartaccountTypeID string, virtualaccountname string, LicenseTermDetailsQueryParams *LicenseTermDetailsQueryParams) (*ResponseLicensesLicenseTermDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/term/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	queryString, _ := query.Values(LicenseTermDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesLicenseTermDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LicenseTermDetails")
	}

	result := response.Result().(*ResponseLicensesLicenseTermDetails)
	return result, response, err

}

//LicenseUsageDetails License Usage Details - f191-0b23-40fb-89d3
/* Get count of purchased and in use DNA and Network licenses.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license usage detail for all virtual accounts.

@param LicenseUsageDetailsQueryParams Filtering parameter
*/
func (s *LicensesService) LicenseUsageDetails(smartaccountTypeID string, virtualaccountname string, LicenseUsageDetailsQueryParams *LicenseUsageDetailsQueryParams) (*ResponseLicensesLicenseUsageDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/usage/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	queryString, _ := query.Values(LicenseUsageDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseLicensesLicenseUsageDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation LicenseUsageDetails")
	}

	result := response.Result().(*ResponseLicensesLicenseUsageDetails)
	return result, response, err

}

//ChangeVirtualAccount Change Virtual Account - ba89-1873-4089-9a37
/* Transfer device(s) from one virtual account to another within same smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of target virtual account

*/
func (s *LicensesService) ChangeVirtualAccount(smartaccountTypeID string, virtualaccountname string, requestLicensesChangeVirtualAccount *RequestLicensesChangeVirtualAccount) (*ResponseLicensesChangeVirtualAccount, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/{smart_account_id}/virtualAccount/{virtual_account_name}/device/transfer"
	path = strings.Replace(path, "{smart_account_id}", fmt.Sprintf("%v", smartaccountTypeID), -1)
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesChangeVirtualAccount).
		SetResult(&ResponseLicensesChangeVirtualAccount{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ChangeVirtualAccount")
	}

	result := response.Result().(*ResponseLicensesChangeVirtualAccount)
	return result, response, err

}

//DeviceDeregistration Device Deregistration - f1bf-1b00-4f8a-a7d2
/* Deregister device(s) from CSSM(Cisco Smart Software Manager).


 */
func (s *LicensesService) DeviceDeregistration(requestLicensesDeviceDeregistration *RequestLicensesDeviceDeregistration) (*ResponseLicensesDeviceDeregistration, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/virtualAccount/deregister"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesDeviceDeregistration).
		SetResult(&ResponseLicensesDeviceDeregistration{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceDeregistration")
	}

	result := response.Result().(*ResponseLicensesDeviceDeregistration)
	return result, response, err

}

//DeviceRegistration Device Registration - a097-5922-439b-a63a
/* Register device(s) in CSSM(Cisco Smart Software Manager).


@param virtualaccountname virtual_account_name path parameter. Name of virtual account

*/
func (s *LicensesService) DeviceRegistration(virtualaccountname string, requestLicensesDeviceRegistration *RequestLicensesDeviceRegistration) (*ResponseLicensesDeviceRegistration, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenses/smartAccount/virtualAccount/{virtual_account_name}/register"
	path = strings.Replace(path, "{virtual_account_name}", fmt.Sprintf("%v", virtualaccountname), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesDeviceRegistration).
		SetResult(&ResponseLicensesDeviceRegistration{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeviceRegistration")
	}

	result := response.Result().(*ResponseLicensesDeviceRegistration)
	return result, response, err

}
