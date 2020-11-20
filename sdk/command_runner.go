package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// CommandRunnerService is the service to communicate with the CommandRunner API endpoint
type CommandRunnerService service

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest is the runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest struct {
	Commands    []string `json:"commands,omitempty"`    //
	Description string   `json:"description,omitempty"` //
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` //
	Name        string   `json:"name,omitempty"`        //
	Timeout     int      `json:"timeout,omitempty"`     //
}

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequestCommands is the runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequestCommands definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequestCommands []string

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequestDeviceUUIDs is the runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequestDeviceUUIDs definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequestDeviceUUIDs []string

// GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse is the getAllKeywordsOfCLIsAcceptedByCommandRunnerResponse definition
type GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponseResponse is the getAllKeywordsOfCLIsAcceptedByCommandRunnerResponseResponse definition
type GetAllKeywordsOfCLIsAcceptedByCommandRunnerResponseResponse []string

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse is the runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse struct {
	Response RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponseResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  //
}

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponseResponse is the runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponseResponse definition
type RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getAllKeywordsOfCLIsAcceptedByCommandRunner")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration")
	}

	result := response.Result().(*RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse)
	return result, response, err
}
