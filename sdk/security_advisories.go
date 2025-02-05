package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SecurityAdvisoriesService service

type ResponseSecurityAdvisoriesGetAdvisoriesList struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesListResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoriesListResponse struct {
	AdvisoryID string `json:"advisoryId,omitempty"` // Advisory Id

	DeviceCount *int `json:"deviceCount,omitempty"` // Device Count

	HiddenDeviceCount *float64 `json:"hiddenDeviceCount,omitempty"` // Hidden Device Count

	Cves []string `json:"cves,omitempty"` // Cves

	PublicationURL string `json:"publicationUrl,omitempty"` // Publication Url

	Sir string `json:"sir,omitempty"` // Sir

	DetectionType string `json:"detectionType,omitempty"` // Detection Type

	DefaultDetectionType string `json:"defaultDetectionType,omitempty"` // Default Detection Type

	DefaultConfigMatchPattern string `json:"defaultConfigMatchPattern,omitempty"` // Default Config Match Pattern

	FixedVersions *ResponseSecurityAdvisoriesGetAdvisoriesListResponseFixedVersions `json:"fixedVersions,omitempty"` //
}
type ResponseSecurityAdvisoriesGetAdvisoriesListResponseFixedVersions struct {
	Version string `json:"version,omitempty"` // version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummary struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponse struct {
	INFORMATIONAL *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseINFORMATIONAL `json:"INFORMATIONAL,omitempty"` //
	LOW           *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseLOW           `json:"LOW,omitempty"`           //
	MEDIUM        *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseMEDIUM        `json:"MEDIUM,omitempty"`        //
	HIGH          *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseHIGH          `json:"HIGH,omitempty"`          //
	CRITICaL      *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseCRITICaL      `json:"CRITICAL,omitempty"`      //
	NA            *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseNA            `json:"NA,omitempty"`            //
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseINFORMATIONAL struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseLOW struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseMEDIUM struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseHIGH struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseCRITICaL struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseNA struct {
	CONFIG       *int `json:"CONFIG,omitempty"`        // Number of advisories matched using default config
	CUSTOMCONFIG *int `json:"CUSTOM_CONFIG,omitempty"` // Number of advisories matched using user provided config
	VERSION      *int `json:"VERSION,omitempty"`       // Number of advisories matched using software version
	TOTAL        *int `json:"TOTAL,omitempty"`         // Sum of Config, Custom Config and Version
}
type ResponseSecurityAdvisoriesGetDevicesPerAdvisory struct {
	Response []string `json:"response,omitempty"` // List of device IDs vulnerable to the advisory
	Version  string   `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoryDeviceDetail struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoryDeviceDetailResponse struct {
	DeviceID            string   `json:"deviceId,omitempty"`            // Network device ID
	AdvisoryIDs         []string `json:"advisoryIds,omitempty"`         // Advisories detected on the network device
	HiddenAdvisoryCount *int     `json:"hiddenAdvisoryCount,omitempty"` // Number of advisories detected on the network device that were suppressed by the user
	ScanMode            string   `json:"scanMode,omitempty"`            // Criteria on which the network device was scanned
	ScanStatus          string   `json:"scanStatus,omitempty"`          // Status of the scan performed on the network device
	Comments            string   `json:"comments,omitempty"`            // More details about the scan status. Ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime        *int     `json:"lastScanTime,omitempty"`        // Time at which the network device was scanned. The representation is unix time.
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDevice struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version of the response
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceResponse struct {
	AdvisoryID                string                                                                 `json:"advisoryId,omitempty"`                // Id of the advisory
	DeviceCount               *int                                                                   `json:"deviceCount,omitempty"`               // Number of devices vulnerable to the advisory
	HiddenDeviceCount         *int                                                                   `json:"hiddenDeviceCount,omitempty"`         // Number of devices vulnerable to the advisory but were suppressed by the user
	Cves                      []string                                                               `json:"cves,omitempty"`                      // CVE (Common Vulnerabilities and Exposures) IDs of the advisory
	PublicationURL            string                                                                 `json:"publicationUrl,omitempty"`            // CISCO publication URL for the advisory
	Sir                       string                                                                 `json:"sir,omitempty"`                       // Security Impact Rating of the advisory
	DetectionType             string                                                                 `json:"detectionType,omitempty"`             // Criteria for advisory detection
	DefaultDetectionType      string                                                                 `json:"defaultDetectionType,omitempty"`      // Original criteria for advisory detection
	DefaultConfigMatchPattern string                                                                 `json:"defaultConfigMatchPattern,omitempty"` // Regular expression used by the system to detect the advisory
	FixedVersions             *ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceResponseFixedVersions `json:"fixedVersions,omitempty"`             // Map where each key is a vulnerable version and the value is a list of versions in which the advisory has been fixed
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceResponseFixedVersions interface{}

//GetAdvisoriesList Get Advisories List - 4295-0bf8-4939-ac35
/* Retrieves list of advisories on the network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisories-list
*/
func (s *SecurityAdvisoriesService) GetAdvisoriesList() (*ResponseSecurityAdvisoriesGetAdvisoriesList, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/advisory"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoriesList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoriesList()
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoriesList")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoriesList)
	return result, response, err

}

//GetAdvisoriesSummary Get Advisories Summary - 3ebf-898d-482b-9207
/* Retrieves summary of advisories on the network.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisories-summary
*/
func (s *SecurityAdvisoriesService) GetAdvisoriesSummary() (*ResponseSecurityAdvisoriesGetAdvisoriesSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/advisory/aggregate"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoriesSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoriesSummary()
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoriesSummary")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoriesSummary)
	return result, response, err

}

//GetDevicesPerAdvisory Get Devices Per Advisory - f49c-4ae0-43fa-8352
/* Retrieves list of devices for an advisory


@param advisoryID advisoryId path parameter. Advisory ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-per-advisory
*/
func (s *SecurityAdvisoriesService) GetDevicesPerAdvisory(advisoryID string) (*ResponseSecurityAdvisoriesGetDevicesPerAdvisory, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/advisory/{advisoryId}/device"
	path = strings.Replace(path, "{advisoryId}", fmt.Sprintf("%v", advisoryID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetDevicesPerAdvisory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesPerAdvisory(advisoryID)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesPerAdvisory")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetDevicesPerAdvisory)
	return result, response, err

}

//GetAdvisoryDeviceDetail Get Advisory Device Detail - e295-09d0-420b-8cc4
/* Retrieves advisory device details for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisory-device-detail
*/
func (s *SecurityAdvisoriesService) GetAdvisoryDeviceDetail(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoryDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoryDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoryDeviceDetail(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoryDeviceDetail")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoryDeviceDetail)
	return result, response, err

}

//GetAdvisoriesPerDevice Get Advisories Per Device - 42a6-c9a1-4ea9-b002
/* Retrieves list of advisories for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisories-per-device
*/
func (s *SecurityAdvisoriesService) GetAdvisoriesPerDevice(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoriesPerDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/device/{deviceId}/advisory"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoriesPerDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoriesPerDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoriesPerDevice")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoriesPerDevice)
	return result, response, err

}
