package dnac

import (
	"github.com/go-resty/resty/v2"
)

// CommandRunnerService is the service to communicate with the CommandRunner API endpoint
type CommandRunnerService service

// CommandRunnerDTO is the CommandRunnerDTO definition
type CommandRunnerDTO struct {
	Commands    []string `json:"commands,omitempty"`    //
	Description string   `json:"description,omitempty"` //
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` //
	Name        string   `json:"name,omitempty"`        //
	Timeout     int      `json:"timeout,omitempty"`     //
}

// LegitCliKeyResult is the LegitCliKeyResult definition
type LegitCliKeyResult struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// TaskIDResult is the TaskIDResult definition
type TaskIDResult struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
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

// // RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration
// /* Submit request for read-only CLIs
//  */
// func (s *CommandRunnerService) RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest *RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/network-device-poller/cli/read-request"

// 	response, err := RestyClient.R().
// 		SetBody(runReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }
