package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// CommandRunnerService is the service to communicate with the CommandRunner API endpoint
type CommandRunnerService service

// CommandRunnerDTO is the CommandRunnerDTO definition
type CommandRunnerDTO struct {
	Commands    []string `json:"commands,omitempty"`    //
	Description string   `json:"description,omitempty"` //
	DeviceUuids []string `json:"deviceUuids,omitempty"` //
	Name        string   `json:"name,omitempty"`        //
	Timeout     int      `json:"timeout,omitempty"`     //
}

// LegitCliKeyResult is the LegitCliKeyResult definition
type LegitCliKeyResult struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// Response is the Response definition
type Response struct {
	TaskId string `json:"taskId,omitempty"` //
	Url    string `json:"url,omitempty"`    //
}

// TaskIdResult is the TaskIdResult definition
type TaskIdResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetAllKeywordsOfCLIsAcceptedByCommandRunner getAllKeywordsOfCLIsAcceptedByCommandRunner
/* Get valid keywords
 */
func (s *CommandRunnerService) GetAllKeywordsOfCLIsAcceptedByCommandRunner() (*LegitCliKeyResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device-poller/cli/legit-reads"

	response, err := RestyClient.R().
		SetResult(&LegitCliKeyResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*LegitCliKeyResult)
	return result, response, err

}

// RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration
/* Submit request for read-only CLIs
 */
func (s *CommandRunnerService) RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest *RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/network-device-poller/cli/read-request"

	response, err := RestyClient.R().
		SetBody(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}
