package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type SecurityAdvisoriesService service

type ResponseSecurityAdvisoriesGetAdvisoriesList struct {
	Response *[]ResponseSecurityAdvisoriesGetAdvisoriesListResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesListResponse struct {
	AdvisoryID                string   `json:"advisoryId,omitempty"`                // Advisory Id
	DeviceCount               *int     `json:"deviceCount,omitempty"`               // Device Count
	HiddenDeviceCount         *float64 `json:"hiddenDeviceCount,omitempty"`         // Hidden Device Count
	Cves                      []string `json:"cves,omitempty"`                      // Cves
	PublicationURL            string   `json:"publicationUrl,omitempty"`            // Publication Url
	Sir                       string   `json:"sir,omitempty"`                       // Sir
	DetectionType             string   `json:"detectionType,omitempty"`             // Detection Type
	DefaultDetectionType      string   `json:"defaultDetectionType,omitempty"`      // Default Detection Type
	DefaultConfigMatchPattern string   `json:"defaultConfigMatchPattern,omitempty"` // Default Config Match Pattern
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummary struct {
	Response *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponse struct {
	NA            *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseNA            `json:"NA,omitempty"`            //
	INFORMATIONAL *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseINFORMATIONAL `json:"INFORMATIONAL,omitempty"` //
	LOW           *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseLOW           `json:"LOW,omitempty"`           //
	MEDIUM        *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseMEDIUM        `json:"MEDIUM,omitempty"`        //
	HIGH          *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseHIGH          `json:"HIGH,omitempty"`          //
	CRITICaL      *ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseCRITICaL      `json:"CRITICAL,omitempty"`      //
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseNA struct {
	CONFIG  *float64 `json:"CONFIG,omitempty"`  // C O N F I G
	VERSION *float64 `json:"VERSION,omitempty"` // V E R S I O N
	TOTAL   *float64 `json:"TOTAL,omitempty"`   // T O T A L
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseINFORMATIONAL struct {
	CONFIG  *float64 `json:"CONFIG,omitempty"`  // C O N F I G
	VERSION *float64 `json:"VERSION,omitempty"` // V E R S I O N
	TOTAL   *float64 `json:"TOTAL,omitempty"`   // T O T A L
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseLOW struct {
	CONFIG  *float64 `json:"CONFIG,omitempty"`  // C O N F I G
	VERSION *float64 `json:"VERSION,omitempty"` // V E R S I O N
	TOTAL   *float64 `json:"TOTAL,omitempty"`   // T O T A L
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseMEDIUM struct {
	CONFIG  *float64 `json:"CONFIG,omitempty"`  // C O N F I G
	VERSION *int     `json:"VERSION,omitempty"` // V E R S I O N
	TOTAL   *int     `json:"TOTAL,omitempty"`   // T O T A L
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseHIGH struct {
	CONFIG  *float64 `json:"CONFIG,omitempty"`  // C O N F I G
	VERSION *int     `json:"VERSION,omitempty"` // V E R S I O N
	TOTAL   *int     `json:"TOTAL,omitempty"`   // T O T A L
}
type ResponseSecurityAdvisoriesGetAdvisoriesSummaryResponseCRITICaL struct {
	CONFIG  *float64 `json:"CONFIG,omitempty"`  // C O N F I G
	VERSION *int     `json:"VERSION,omitempty"` // V E R S I O N
	TOTAL   *int     `json:"TOTAL,omitempty"`   // T O T A L
}
type ResponseSecurityAdvisoriesGetDevicesPerAdvisory struct {
	Response []string `json:"response,omitempty"` // Response
	Version  string   `json:"version,omitempty"`  // Version
}
type ResponseSecurityAdvisoriesGetAdvisoryIDsPerDevice struct {
	Response *[]ResponseSecurityAdvisoriesGetAdvisoryIDsPerDeviceResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseSecurityAdvisoriesGetAdvisoryIDsPerDeviceResponse struct {
	DeviceID    string   `json:"deviceId,omitempty"`    // Device Id
	AdvisoryIDs []string `json:"advisoryIds,omitempty"` // Advisory Ids
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDevice struct {
	Response *[]ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseSecurityAdvisoriesGetAdvisoriesPerDeviceResponse struct {
	AdvisoryID                string   `json:"advisoryId,omitempty"`                // Advisory Id
	DeviceCount               *int     `json:"deviceCount,omitempty"`               // Device Count
	HiddenDeviceCount         *float64 `json:"hiddenDeviceCount,omitempty"`         // Hidden Device Count
	Cves                      []string `json:"cves,omitempty"`                      // Cves
	PublicationURL            string   `json:"publicationUrl,omitempty"`            // Publication Url
	Sir                       string   `json:"sir,omitempty"`                       // Sir
	DetectionType             string   `json:"detectionType,omitempty"`             // Detection Type
	DefaultDetectionType      string   `json:"defaultDetectionType,omitempty"`      // Default Detection Type
	DefaultConfigMatchPattern string   `json:"defaultConfigMatchPattern,omitempty"` // Default Config Match Pattern
}

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

//GetAdvisoryIDsPerDevice Get Advisory IDs Per Device - e295-09d0-420b-8cc4
/* Retrieves list of advisory IDs for a device


@param deviceID deviceId path parameter. Device instance UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-advisory-ids-per-device
*/
func (s *SecurityAdvisoriesService) GetAdvisoryIDsPerDevice(deviceID string) (*ResponseSecurityAdvisoriesGetAdvisoryIDsPerDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/security-advisory/device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSecurityAdvisoriesGetAdvisoryIDsPerDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAdvisoryIDsPerDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetAdvisoryIdsPerDevice")
	}

	result := response.Result().(*ResponseSecurityAdvisoriesGetAdvisoryIDsPerDevice)
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
