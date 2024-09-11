package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type LicensesService service

type DeviceCountDetailsQueryParams struct {
	DeviceType         string `url:"device_type,omitempty"`          //Type of device
	RegistrationStatus string `url:"registration_status,omitempty"`  //Smart license registration status of device
	DnaLevel           string `url:"dna_level,omitempty"`            //Device Cisco Catalyst Center License Level
	VirtualAccountName string `url:"virtual_account_name,omitempty"` //Virtual account name
	SmartAccountID     string `url:"smart_account_id,omitempty"`     //Smart account id
}
type DeviceLicenseSummaryQueryParams struct {
	PageNumber         float64 `url:"page_number,omitempty"`          //Page number of response
	Order              string  `url:"order,omitempty"`                //Sorting order
	SortBy             string  `url:"sort_by,omitempty"`              //Sort result by field
	DnaLevel           string  `url:"dna_level,omitempty"`            //Device Cisco Catalyst Center license level. The valid values are Advantage, Essentials
	DeviceType         string  `url:"device_type,omitempty"`          //Type of device. The valid values are Routers, Switches and Hubs, Wireless Controller
	Limit              float64 `url:"limit,omitempty"`                //Limit
	RegistrationStatus string  `url:"registration_status,omitempty"`  //Smart license registration status of device. The valid values are Unknown, NA, Unregistered, Registered, Registration_expired, Reservation_in_progress, Registered_slr, Registered_plr, Registered_satellite
	VirtualAccountName string  `url:"virtual_account_name,omitempty"` //Name of virtual account
	SmartAccountID     string  `url:"smart_account_id,omitempty"`     //Id of smart account
	DeviceUUID         string  `url:"device_uuid,omitempty"`          //Id of device
}
type LicenseTermDetailsQueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}
type LicenseUsageDetailsQueryParams struct {
	DeviceType string `url:"device_type,omitempty"` //Type of device like router, switch, wireless or ise
}

type ResponseLicensesRetrieveLicenseSetting struct {
	Response *ResponseLicensesRetrieveLicenseSettingResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // API version
}
type ResponseLicensesRetrieveLicenseSettingResponse struct {
	DefaultSmartAccountID            string `json:"defaultSmartAccountId,omitempty"`            // Default smart account id
	AutoRegistrationVirtualAccountID string `json:"autoRegistrationVirtualAccountId,omitempty"` // Virtual account id
}
type ResponseLicensesUpdateLicenseSetting struct {
	Response *ResponseLicensesUpdateLicenseSettingResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // API version
}
type ResponseLicensesUpdateLicenseSettingResponse struct {
	DefaultSmartAccountID            string `json:"defaultSmartAccountId,omitempty"`            // Default smart account id
	AutoRegistrationVirtualAccountID string `json:"autoRegistrationVirtualAccountId,omitempty"` // Virtual account id
}
type ResponseLicensesDeviceCountDetails struct {
	Response *int   `json:"response,omitempty"` // Total number of managed device
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseLicensesDeviceLicenseSummary struct {
	Response *[]ResponseLicensesDeviceLicenseSummaryResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
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
	DnaLevel                         string `json:"dna_level,omitempty"`                             // Device Cisco Catalyst Center license level
	VirtualAccountName               string `json:"virtual_account_name,omitempty"`                  // Name of virtual account
	LastSuccessfulRumUsageUploadTime string `json:"last_successful_rum_usage_upload_time,omitempty"` // Last successful rum usage upload time
	IPAddress                        string `json:"ip_address,omitempty"`                            // IP address of device
	WirelessCapableDnaLicense        string `json:"wireless_capable_dna_license,omitempty"`          // Wireless Cisco Catalyst Center license value
	MacAddress                       string `json:"mac_address,omitempty"`                           // MAC address of device
	CustomerTag1                     string `json:"customer_tag1,omitempty"`                         // Customer Tag1 set on device
	CustomerTag2                     string `json:"customer_tag2,omitempty"`                         // Customer Tag2 set on device
	CustomerTag3                     string `json:"customer_tag3,omitempty"`                         // Customer Tag3 set on device
	CustomerTag4                     string `json:"customer_tag4,omitempty"`                         // Customer Tag4 set on device
	SmartAccountName                 string `json:"smart_account_name,omitempty"`                    // Name of smart account
}
type ResponseLicensesDeviceLicenseDetails struct {
	DeviceUUID              string                                                `json:"device_uuid,omitempty"`               // Id of device
	Site                    string                                                `json:"site,omitempty"`                      // Site of device
	Model                   string                                                `json:"model,omitempty"`                     // Model of device
	LicenseMode             string                                                `json:"license_mode,omitempty"`              // Mode of license
	IsLicenseExpired        *bool                                                 `json:"is_license_expired,omitempty"`        // Is device license expired
	SoftwareVersion         string                                                `json:"software_version,omitempty"`          // Software image version of device
	NetworkLicense          string                                                `json:"network_license,omitempty"`           // Device network license level
	EvaluationLicenseExpiry string                                                `json:"evaluation_license_expiry,omitempty"` // Evaluation period expiry date
	DeviceName              string                                                `json:"device_name,omitempty"`               // Name of device
	DeviceType              string                                                `json:"device_type,omitempty"`               // Type of device
	DnaLevel                string                                                `json:"dna_level,omitempty"`                 // Device Cisco Catalyst Center license level
	VirtualAccountName      string                                                `json:"virtual_account_name,omitempty"`      // Name of virtual account
	IPAddress               string                                                `json:"ip_address,omitempty"`                // IP address of device
	MacAddress              string                                                `json:"mac_address,omitempty"`               // MAC address of device
	SntcStatus              string                                                `json:"sntc_status,omitempty"`               // Valid if device is covered under license contract and invalid if not covered, otherwise unknown.
	FeatureLicense          []string                                              `json:"feature_license,omitempty"`           // Name of feature licenses
	HasSupCards             *bool                                                 `json:"has_sup_cards,omitempty"`             // Whether device has supervisor cards
	Udi                     string                                                `json:"udi,omitempty"`                       // Unique Device Identifier
	StackedDevices          *[]ResponseLicensesDeviceLicenseDetailsStackedDevices `json:"stacked_devices,omitempty"`           //
	IsStackedDevice         *bool                                                 `json:"is_stacked_device,omitempty"`         // Is Stacked Device
	AccessPoints            *[]ResponseLicensesDeviceLicenseDetailsAccessPoints   `json:"access_points,omitempty"`             //
	ChassisDetails          *ResponseLicensesDeviceLicenseDetailsChassisDetails   `json:"chassis_details,omitempty"`           //
}
type ResponseLicensesDeviceLicenseDetailsStackedDevices struct {
	MacAddress   string `json:"mac_address,omitempty"`   // Stack mac address
	ID           *int   `json:"id,omitempty"`            // Id
	Role         string `json:"role,omitempty"`          // Chassis role
	SerialNumber string `json:"serial_number,omitempty"` // Chassis serial number
}
type ResponseLicensesDeviceLicenseDetailsAccessPoints struct {
	ApType string `json:"ap_type,omitempty"` // Type of access point
	Count  string `json:"count,omitempty"`   // Number of access point
}
type ResponseLicensesDeviceLicenseDetailsChassisDetails struct {
	BoardSerialNumber string                                                               `json:"board_serial_number,omitempty"` // Board serial number
	Modules           *[]ResponseLicensesDeviceLicenseDetailsChassisDetailsModules         `json:"modules,omitempty"`             //
	SupervisorCards   *[]ResponseLicensesDeviceLicenseDetailsChassisDetailsSupervisorCards `json:"supervisor_cards,omitempty"`    //
	Port              *int                                                                 `json:"port,omitempty"`                // Number of port
}
type ResponseLicensesDeviceLicenseDetailsChassisDetailsModules struct {
	ModuleType   string `json:"module_type,omitempty"`   // Type of module
	ModuleName   string `json:"module_name,omitempty"`   // Name of module
	SerialNumber string `json:"serial_number,omitempty"` // Serial number of module
	ID           *int   `json:"id,omitempty"`            // Id of module
}
type ResponseLicensesDeviceLicenseDetailsChassisDetailsSupervisorCards struct {
	SerialNumber       string `json:"serial_number,omitempty"`        // Serial number of supervisor card
	SupervisorCardType string `json:"supervisor_card_type,omitempty"` // Type of supervisor card
	Status             string `json:"status,omitempty"`               // Status of supervisor card
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
	DnaLevel                 string `json:"dna_level,omitempty"`                   // Cisco Catalyst Center license level
	PurchasedDnaLicenseCount string `json:"purchased_dna_license_count,omitempty"` // Number of purchased Catalyst Center licenses
	IsLicenseExpired         string `json:"is_license_expired,omitempty"`          // Is license expired
}
type ResponseLicensesLicenseUsageDetails struct {
	PurchasedDnaLicense     *ResponseLicensesLicenseUsageDetailsPurchasedDnaLicense     `json:"purchased_dna_license,omitempty"`     //
	PurchasedNetworkLicense *ResponseLicensesLicenseUsageDetailsPurchasedNetworkLicense `json:"purchased_network_license,omitempty"` //
	UsedDnaLicense          *ResponseLicensesLicenseUsageDetailsUsedDnaLicense          `json:"used_dna_license,omitempty"`          //
	UsedNetworkLicense      *ResponseLicensesLicenseUsageDetailsUsedNetworkLicense      `json:"used_network_license,omitempty"`      //
	PurchasedIseLicense     *ResponseLicensesLicenseUsageDetailsPurchasedIseLicense     `json:"purchased_ise_license,omitempty"`     //
	UsedIseLicense          *ResponseLicensesLicenseUsageDetailsUsedIseLicense          `json:"used_ise_license,omitempty"`          //
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
type ResponseLicensesLicenseUsageDetailsPurchasedIseLicense struct {
	TotalLicenseCount  *int                                                                        `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsPurchasedIseLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsPurchasedIseLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type ResponseLicensesLicenseUsageDetailsUsedIseLicense struct {
	TotalLicenseCount  *int                                                                   `json:"total_license_count,omitempty"`   // Total number of licenses
	LicenseCountByType *[]ResponseLicensesLicenseUsageDetailsUsedIseLicenseLicenseCountByType `json:"license_count_by_type,omitempty"` //
}
type ResponseLicensesLicenseUsageDetailsUsedIseLicenseLicenseCountByType struct {
	LicenseType  string `json:"license_type,omitempty"`  // Type of license
	LicenseCount *int   `json:"license_count,omitempty"` // Number of licenses
}
type RequestLicensesUpdateLicenseSetting struct {
	DefaultSmartAccountID            string `json:"defaultSmartAccountId,omitempty"`            // Default smart account id
	AutoRegistrationVirtualAccountID string `json:"autoRegistrationVirtualAccountId,omitempty"` // Virtual account id
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

//RetrieveLicenseSetting Retrieve license setting - c489-d9a3-4c09-84c7
/* Retrieves license setting Default smart account id and virtual account id for auto registration of devices for smart license flow. If default smart account is not configured, 'defaultSmartAccountId' is 'null'. Similarly, if auto registration of devices for smart license flow is not enabled, 'autoRegistrationVirtualAccountId' is 'null'. For smart proxy connection mode, 'autoRegistrationVirtualAccountId' is always 'null'.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-license-setting
*/
func (s *LicensesService) RetrieveLicenseSetting() (*ResponseLicensesRetrieveLicenseSetting, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenseSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseLicensesRetrieveLicenseSetting{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveLicenseSetting()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveLicenseSetting")
	}

	result := response.Result().(*ResponseLicensesRetrieveLicenseSetting)
	return result, response, err

}

//DeviceCountDetails Device Count Details - 949a-6983-4c38-9d59
/* Get total number of managed device(s).


@param DeviceCountDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-count-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceCountDetails(DeviceCountDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeviceCountDetails")
	}

	result := response.Result().(*ResponseLicensesDeviceCountDetails)
	return result, response, err

}

//DeviceLicenseSummary Device License Summary - 529e-2ae0-4a9b-89b1
/* Show license summary of device(s).


@param DeviceLicenseSummaryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-license-summary
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceLicenseSummary(DeviceLicenseSummaryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeviceLicenseSummary")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseSummary)
	return result, response, err

}

//DeviceLicenseDetails Device License Details - dca1-1bc2-4e7b-8c5d
/* Get detailed license information of a device.


@param deviceuuid device_uuid path parameter. Id of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-license-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceLicenseDetails(deviceuuid)
		}
		return nil, response, fmt.Errorf("error with operation DeviceLicenseDetails")
	}

	result := response.Result().(*ResponseLicensesDeviceLicenseDetails)
	return result, response, err

}

//VirtualAccountDetails Virtual Account Details - 829e-daf5-4d49-a581
/* Get virtual account details of a smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account


Documentation Link: https://developer.cisco.com/docs/dna-center/#!virtual-account-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.VirtualAccountDetails(smartaccountTypeID)
		}
		return nil, response, fmt.Errorf("error with operation VirtualAccountDetails")
	}

	result := response.Result().(*ResponseLicensesVirtualAccountDetails)
	return result, response, err

}

//SmartAccountDetails Smart Account Details - c181-b8db-4c89-adf0
/* Retrieve details of all smart accounts.



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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SmartAccountDetails()
		}
		return nil, response, fmt.Errorf("error with operation SmartAccountDetails")
	}

	result := response.Result().(*ResponseLicensesSmartAccountDetails)
	return result, response, err

}

//LicenseTermDetails License Term Details - e891-4b73-4ad9-b414
/* Get license term details.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseTermDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!license-term-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LicenseTermDetails(smartaccountTypeID, virtualaccountname, LicenseTermDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LicenseTermDetails")
	}

	result := response.Result().(*ResponseLicensesLicenseTermDetails)
	return result, response, err

}

//LicenseUsageDetails License Usage Details - 418a-6b43-4e29-bfe5
/* Get count of purchased and in use Cisco Catalyst Center and Network licenses.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.

@param LicenseUsageDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!license-usage-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LicenseUsageDetails(smartaccountTypeID, virtualaccountname, LicenseUsageDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation LicenseUsageDetails")
	}

	result := response.Result().(*ResponseLicensesLicenseUsageDetails)
	return result, response, err

}

//ChangeVirtualAccount Change Virtual Account - bea7-4a0b-4778-8c89
/* Transfer device(s) from one virtual account to another within same smart account.


@param smartaccountTypeID smart_account_id path parameter. Id of smart account

@param virtualaccountname virtual_account_name path parameter. Name of target virtual account


Documentation Link: https://developer.cisco.com/docs/dna-center/#!change-virtual-account
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ChangeVirtualAccount(smartaccountTypeID, virtualaccountname, requestLicensesChangeVirtualAccount)
		}

		return nil, response, fmt.Errorf("error with operation ChangeVirtualAccount")
	}

	result := response.Result().(*ResponseLicensesChangeVirtualAccount)
	return result, response, err

}

//UpdateLicenseSetting Update license setting - 97ae-8980-475a-961e
/* Update license setting Configure default smart account id  and/or virtual account id for auto registration of devices for smart license flow. Virtual account should be part of default smart account. Default smart account id cannot be set to 'null'. Auto registration of devices for smart license flow is applicable only for direct or on-prem SSM connection mode.


 */
func (s *LicensesService) UpdateLicenseSetting(requestLicensesUpdateLicenseSetting *RequestLicensesUpdateLicenseSetting) (*ResponseLicensesUpdateLicenseSetting, *resty.Response, error) {
	path := "/dna/intent/api/v1/licenseSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestLicensesUpdateLicenseSetting).
		SetResult(&ResponseLicensesUpdateLicenseSetting{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateLicenseSetting(requestLicensesUpdateLicenseSetting)
		}
		return nil, response, fmt.Errorf("error with operation UpdateLicenseSetting")
	}

	result := response.Result().(*ResponseLicensesUpdateLicenseSetting)
	return result, response, err

}

//DeviceDeregistration Device Deregistration - 8c82-dad4-49ba-b8eb
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceDeregistration(requestLicensesDeviceDeregistration)
		}
		return nil, response, fmt.Errorf("error with operation DeviceDeregistration")
	}

	result := response.Result().(*ResponseLicensesDeviceDeregistration)
	return result, response, err

}

//DeviceRegistration Device Registration - a08b-eae5-47fb-95e3
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceRegistration(virtualaccountname, requestLicensesDeviceRegistration)
		}
		return nil, response, fmt.Errorf("error with operation DeviceRegistration")
	}

	result := response.Result().(*ResponseLicensesDeviceRegistration)
	return result, response, err

}
