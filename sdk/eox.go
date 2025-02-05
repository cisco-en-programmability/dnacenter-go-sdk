package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EoXService service

type GetEoXStatusForAllDevicesQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page, the first record is numbered 1
}

type ResponseEoXGetEoXStatusForAllDevices struct {
	Response *[]ResponseEoXGetEoXStatusForAllDevicesResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version of the response
}
type ResponseEoXGetEoXStatusForAllDevicesResponse struct {
	DeviceID     string                                                 `json:"deviceId,omitempty"`     // Device instance UUID
	AlertCount   *int                                                   `json:"alertCount,omitempty"`   // Number of EoX alerts on the network device
	Summary      *[]ResponseEoXGetEoXStatusForAllDevicesResponseSummary `json:"summary,omitempty"`      //
	ScanStatus   string                                                 `json:"scanStatus,omitempty"`   // Status of the scan performed on the network device
	Comments     []string                                               `json:"comments,omitempty"`     // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure.
	LastScanTime *int                                                   `json:"lastScanTime,omitempty"` // Time at which the network device was scanned. The representation is unix time.
}
type ResponseEoXGetEoXStatusForAllDevicesResponseSummary struct {
	EoXType string `json:"eoxType,omitempty"` // Type of EoX Alert
}
type ResponseEoXGetEoXDetailsPerDevice struct {
	Response *ResponseEoXGetEoXDetailsPerDeviceResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version of the response
}
type ResponseEoXGetEoXDetailsPerDeviceResponse struct {
	DeviceID     string                                                 `json:"deviceId,omitempty"`     // Device instance UUID
	AlertCount   *int                                                   `json:"alertCount,omitempty"`   // Number of EoX alerts on the network device
	EoXDetails   *[]ResponseEoXGetEoXDetailsPerDeviceResponseEoXDetails `json:"EoXDetails,omitempty"`   //
	ScanStatus   string                                                 `json:"scanStatus,omitempty"`   // Status of the scan performed on the network device
	Comments     []string                                               `json:"comments,omitempty"`     // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure.
	LastScanTime *int                                                   `json:"lastScanTime,omitempty"` // Time at which the network device was scanned. The representation is unix time.
}
type ResponseEoXGetEoXDetailsPerDeviceResponseEoXDetails struct {
	Name                                              string `json:"name,omitempty"`                                              // Name of the EoX alert. Every EoX announcement has a unique name. ie:- EOL13873
	BulletinHeadline                                  string `json:"bulletinHeadline,omitempty"`                                  // Title of the EoX bulletin
	BulletinName                                      string `json:"bulletinName,omitempty"`                                      // Name of the EoX bulletin
	BulletinNumber                                    string `json:"bulletinNumber,omitempty"`                                    // Identifier of the EoX bulletin. Usually the same as name.
	BulletinURL                                       string `json:"bulletinURL,omitempty"`                                       // URL where the EoX bulletin is posted
	EndOfHardwareNewServiceAttachmentDate             string `json:"endOfHardwareNewServiceAttachmentDate,omitempty"`             // For equipment and software that is not covered by a service-and-support contract, this is the last date to order a new service-and-support contract or add the equipment and/or software to an existing service-and-support contract
	EndOfHardwareServiceContractRenewalDate           string `json:"endOfHardwareServiceContractRenewalDate,omitempty"`           // The last date to extend or renew a service contract for the product
	EndOfLastHardwareShipDate                         string `json:"endOfLastHardwareShipDate,omitempty"`                         // The last-possible ship date that can be requested of Cisco and/or its contract manufacturers
	EndOfLifeExternalAnnouncementDate                 string `json:"endOfLifeExternalAnnouncementDate,omitempty"`                 // The date the document that announces the end-of-sale and end-of-life of a product is distributed to the general public
	EndOfSignatureReleasesDate                        string `json:"endOfSignatureReleasesDate,omitempty"`                        // The date after which there will be no more signature update release for the product
	EndOfSoftwareVulnerabilityOrSecuritySupportDate   string `json:"endOfSoftwareVulnerabilityOrSecuritySupportDate,omitempty"`   // The last date that Cisco Engineering may release bug fixes for Vulnerability or Security issues for the product. This will be populated for software alerts only.
	EndOfSoftwareVulnerabilityOrSecuritySupportDateHw string `json:"endOfSoftwareVulnerabilityOrSecuritySupportDateHw,omitempty"` // The last date that Cisco Engineering may release bug fixes for Vulnerability or Security issues for the product. This will be populated for hardware or module alerts only.
	EndOfSaleDate                                     string `json:"endOfSaleDate,omitempty"`                                     // The last date to order the product through Cisco point-of-sale mechanisms
	EndOfLifeDate                                     string `json:"endOfLifeDate,omitempty"`                                     // The last date to receive applicable service and support for the product as entitled by active service contracts or by warranty terms and conditions. This will be populated for software alerts only.
	LastDateOfSupport                                 string `json:"lastDateOfSupport,omitempty"`                                 // The last date to receive applicable service and support for the product as entitled by active service contracts or by warranty terms and conditions. This will be populated for hardware and module alerts only.
	EndOfSoftwareMaintenanceReleasesDate              string `json:"endOfSoftwareMaintenanceReleasesDate,omitempty"`              // The last date that Cisco Engineering may release any final software maintenance releases or bug fixes for the product
	EoXAlertType                                      string `json:"EoXAlertType,omitempty"`                                      // Type of EoX alert
	EoXPhysicalType                                   string `json:"EoXPhysicalType,omitempty"`                                   // The type of part for EoX alert. eg:- Power Supply, Chassis, Fan etc.
	BulletinPID                                       string `json:"bulletinPID,omitempty"`                                       // The part number for the EoX alert. eg:- PWR-C1-1100WAC
}
type ResponseEoXGetEoXSummary struct {
	Response *ResponseEoXGetEoXSummaryResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  // Version of the response
}
type ResponseEoXGetEoXSummaryResponse struct {
	HardwareCount *int `json:"hardwareCount,omitempty"` // Number of hardware EoX alerts detected on the network
	SoftwareCount *int `json:"softwareCount,omitempty"` // Number of software EoX alerts detected on the network
	ModuleCount   *int `json:"moduleCount,omitempty"`   // Number of module EoX alerts detected on the network
	TotalCount    *int `json:"totalCount,omitempty"`    // Total number of EoX alerts detected on the network. This is the sum of hardwareCount, softwareCount and moduleCount.
}

//GetEoXStatusForAllDevices Get EoX Status For All Devices - 3281-fa04-49ba-87d9
/* Retrieves EoX status for all devices in the network


@param GetEoXStatusForAllDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eox-status-for-all-devices
*/
func (s *EoXService) GetEoXStatusForAllDevices(GetEoXStatusForAllDevicesQueryParams *GetEoXStatusForAllDevicesQueryParams) (*ResponseEoXGetEoXStatusForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/device"

	queryString, _ := query.Values(GetEoXStatusForAllDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEoXGetEoXStatusForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoXStatusForAllDevices(GetEoXStatusForAllDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEoXStatusForAllDevices")
	}

	result := response.Result().(*ResponseEoXGetEoXStatusForAllDevices)
	return result, response, err

}

//GetEoXDetailsPerDevice Get EoX Details Per Device - dc80-099e-4d59-986d
/* Retrieves EoX details for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eox-details-per-device
*/
func (s *EoXService) GetEoXDetailsPerDevice(deviceID string) (*ResponseEoXGetEoXDetailsPerDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEoXGetEoXDetailsPerDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoXDetailsPerDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetEoXDetailsPerDevice")
	}

	result := response.Result().(*ResponseEoXGetEoXDetailsPerDevice)
	return result, response, err

}

//GetEoXSummary Get EoX Summary - f0b2-7a23-4fea-96fc
/* Retrieves EoX summary for all devices in the network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eox-summary
*/
func (s *EoXService) GetEoXSummary() (*ResponseEoXGetEoXSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/summary"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEoXGetEoXSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoXSummary()
		}
		return nil, response, fmt.Errorf("error with operation GetEoXSummary")
	}

	result := response.Result().(*ResponseEoXGetEoXSummary)
	return result, response, err

}
