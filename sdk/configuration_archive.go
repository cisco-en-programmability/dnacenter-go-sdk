package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ConfigurationArchiveService service

type ResponseConfigurationArchiveExportDeviceConfigurations struct {
	Response *ResponseConfigurationArchiveExportDeviceConfigurationsResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseConfigurationArchiveExportDeviceConfigurationsResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type RequestConfigurationArchiveExportDeviceConfigurations struct {
	DeviceID []string `json:"deviceId,omitempty"` // Device Id
	Password string   `json:"password,omitempty"` // Password
}

//ExportDeviceConfigurations Export Device configurations - 51a4-0aba-4c68-ac17
/* Export Device configurations to an encrypted zip file.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-configurations
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurations(requestConfigurationArchiveExportDeviceConfigurations *RequestConfigurationArchiveExportDeviceConfigurations) (*ResponseConfigurationArchiveExportDeviceConfigurations, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-archive/cleartext"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationArchiveExportDeviceConfigurations).
		SetResult(&ResponseConfigurationArchiveExportDeviceConfigurations{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ExportDeviceConfigurations")
	}

	result := response.Result().(*ResponseConfigurationArchiveExportDeviceConfigurations)
	return result, response, err

}
