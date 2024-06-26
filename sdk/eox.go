package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type EoXService service

type ResponseEoXGetEoXStatusForAllDevices struct {
	Response *[]ResponseEoXGetEoXStatusForAllDevicesResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseEoXGetEoXStatusForAllDevicesResponse struct {
	DeviceID     string                                                 `json:"deviceId,omitempty"`     // Device Id
	AlertCount   *int                                                   `json:"alertCount,omitempty"`   // Alert Count
	Summary      *[]ResponseEoXGetEoXStatusForAllDevicesResponseSummary `json:"summary,omitempty"`      //
	ScanStatus   string                                                 `json:"scanStatus,omitempty"`   // Scan Status
	LastScanTime *int                                                   `json:"lastScanTime,omitempty"` // Last Scan Time
}
type ResponseEoXGetEoXStatusForAllDevicesResponseSummary struct {
	EoxType string `json:"eoxType,omitempty"` // Eox Type
}
type ResponseEoXGetEoXDetailsPerDevice struct {
	Response *ResponseEoXGetEoXDetailsPerDeviceResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version
}
type ResponseEoXGetEoXDetailsPerDeviceResponse struct {
	DeviceID     string                                                 `json:"deviceId,omitempty"`     // Device Id
	AlertCount   *int                                                   `json:"alertCount,omitempty"`   // Alert Count
	EoxDetails   *[]ResponseEoXGetEoXDetailsPerDeviceResponseEoxDetails `json:"eoxDetails,omitempty"`   //
	ScanStatus   string                                                 `json:"scanStatus,omitempty"`   // Scan Status
	Comments     *[]ResponseEoXGetEoXDetailsPerDeviceResponseComments   `json:"comments,omitempty"`     // Comments
	LastScanTime *int                                                   `json:"lastScanTime,omitempty"` // Last Scan Time
}
type ResponseEoXGetEoXDetailsPerDeviceResponseEoxDetails struct {
	BulletinHeadline                                  string `json:"bulletinHeadline,omitempty"`                                  // Bulletin Headline
	BulletinNumber                                    string `json:"bulletinNumber,omitempty"`                                    // Bulletin Number
	BulletinURL                                       string `json:"bulletinURL,omitempty"`                                       // Bulletin U R L
	EndOfHardwareNewServiceAttachmentDate             *int   `json:"endOfHardwareNewServiceAttachmentDate,omitempty"`             // End Of Hardware New Service Attachment Date
	EndOfHardwareServiceContractRenewalDate           *int   `json:"endOfHardwareServiceContractRenewalDate,omitempty"`           // End Of Hardware Service Contract Renewal Date
	EndOfLastHardwareShipDate                         *int   `json:"endOfLastHardwareShipDate,omitempty"`                         // End Of Last Hardware Ship Date
	EndOfLifeDate                                     *int   `json:"endOfLifeDate,omitempty"`                                     // End Of Life Date
	EndOfLifeExternalAnnouncementDate                 *int   `json:"endOfLifeExternalAnnouncementDate,omitempty"`                 // End Of Life External Announcement Date
	EndOfSaleDate                                     *int   `json:"endOfSaleDate,omitempty"`                                     // End Of Sale Date
	EndOfSignatureReleasesDate                        *int   `json:"endOfSignatureReleasesDate,omitempty"`                        // End Of Signature Releases Date
	EndOfSoftwareVulnerabilityOrSecuritySupportDate   *int   `json:"endOfSoftwareVulnerabilityOrSecuritySupportDate,omitempty"`   // End Of Software Vulnerability Or Security Support Date
	EndOfSoftwareVulnerabilityOrSecuritySupportDateHw *int   `json:"endOfSoftwareVulnerabilityOrSecuritySupportDateHw,omitempty"` // End Of Software Vulnerability Or Security Support Date Hw
	EndOfSoftwareMaintenanceReleasesDate              *int   `json:"endOfSoftwareMaintenanceReleasesDate,omitempty"`              // End Of Software Maintenance Releases Date
	EoxAlertType                                      string `json:"eoxAlertType,omitempty"`                                      // Eox Alert Type
	LastDateOfSupport                                 *int   `json:"lastDateOfSupport,omitempty"`                                 // Last Date Of Support
	Name                                              string `json:"name,omitempty"`                                              // Name
}
type ResponseEoXGetEoXDetailsPerDeviceResponseComments interface{}
type ResponseEoXGetEoXSummary struct {
	Response *ResponseEoXGetEoXSummaryResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  // Version
}
type ResponseEoXGetEoXSummaryResponse struct {
	HardwareCount *int `json:"hardwareCount,omitempty"` // Hardware Count
	SoftwareCount *int `json:"softwareCount,omitempty"` // Software Count
	ModuleCount   *int `json:"moduleCount,omitempty"`   // Module Count
	TotalCount    *int `json:"totalCount,omitempty"`    // Total Count
}

//GetEoXStatusForAllDevices Get EoX Status For All Devices - 3281-fa04-49ba-87d9
/* Retrieves EoX status for all devices in the network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-eo-x-status-for-all-devices
*/
func (s *EoXService) GetEoXStatusForAllDevices() (*ResponseEoXGetEoXStatusForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/eox-status/device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEoXGetEoXStatusForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEoXStatusForAllDevices()
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
