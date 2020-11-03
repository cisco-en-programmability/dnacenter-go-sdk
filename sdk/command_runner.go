package dnac

import (
	"github.com/go-resty/resty/v2"
)

// CommandRunnerService is the service to communicate with the CommandRunner API endpoint
type CommandRunnerService service

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest is the RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest struct {
	Commands    []string `json:"commands,omitempty"`    //
	Description string   `json:"description,omitempty"` //
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` //
	Name        string   `json:"name,omitempty"`        //
	Timeout     int      `json:"timeout,omitempty"`     //
}

// GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse is the GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse definition
type GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse is the RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetAllKeywordsOfCLIsAcceptedByCommandRunner getAllKeywordsOfCLIsAcceptedByCommandRunner
/* Get valid keywords
 */
func (s *CommandRunnerService) GetAllKeywordsOfCLIsAcceptedByCommandRunner() (*GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device-poller/cli/legit-reads"

	response, err := RestyClient.R().
		SetResult(&GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse)
	return result, response, err
}

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration
/* Submit request for read-only CLIs
 */
func (s *CommandRunnerService) RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest *RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest) (*RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device-poller/cli/read-request"

	response, err := RestyClient.R().
		SetBody(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest).
		SetResult(&RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse)
	return result, response, err
}
