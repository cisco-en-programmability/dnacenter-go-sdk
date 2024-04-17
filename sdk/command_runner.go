package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type CommandRunnerService service

type ResponseCommandRunnerGetAllKeywordsOfCliSAcceptedByCommandRunner struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}
type ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration struct {
	Response *ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  //
}
type ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfigurationResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration struct {
	Commands    []string `json:"commands,omitempty"`    //
	Description string   `json:"description,omitempty"` //
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` //
	Name        string   `json:"name,omitempty"`        //
	Timeout     *int     `json:"timeout,omitempty"`     //
}

//GetAllKeywordsOfCliSAcceptedByCommandRunner Get all keywords of CLIs accepted by command runner - 33bb-2b9d-4019-9e14
/* Get valid keywords



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-keywords-of-c-l-is-accepted-by-command-runner
*/
func (s *CommandRunnerService) GetAllKeywordsOfCliSAcceptedByCommandRunner() (*ResponseCommandRunnerGetAllKeywordsOfCliSAcceptedByCommandRunner, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-poller/cli/legit-reads"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseCommandRunnerGetAllKeywordsOfCliSAcceptedByCommandRunner{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllKeywordsOfCliSAcceptedByCommandRunner()
		}
		return nil, response, fmt.Errorf("error with operation GetAllKeywordsOfCliSAcceptedByCommandRunner")
	}

	result := response.Result().(*ResponseCommandRunnerGetAllKeywordsOfCliSAcceptedByCommandRunner)
	return result, response, err

}

//RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration Run read-only commands on devices to get their real-time configuration - d6b8-ca77-4739-adf4
/* Submit request for read-only CLIs



Documentation Link: https://developer.cisco.com/docs/dna-center/#!run-read-only-commands-on-devices-to-get-their-real-time-configuration
*/
func (s *CommandRunnerService) RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(requestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration *RequestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration) (*ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-poller/cli/read-request"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration).
		SetResult(&ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration(requestCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration)
		}

		return nil, response, fmt.Errorf("error with operation RunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration")
	}

	result := response.Result().(*ResponseCommandRunnerRunReadOnlyCommandsOnDevicesToGetTheirRealTimeConfiguration)
	return result, response, err

}
