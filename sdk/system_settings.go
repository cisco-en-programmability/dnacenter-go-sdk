package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SystemSettingsService service

type ResponseSystemSettingsCustomPromptSupportGETApI struct {
	Response *ResponseSystemSettingsCustomPromptSupportGETAPIResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptSupportGETAPIResponse struct {
	CustomUsernamePrompt  string `json:"customUsernamePrompt,omitempty"`  // Custom Username
	CustomPasswordPrompt  string `json:"customPasswordPrompt,omitempty"`  // Custom Password
	DefaultUsernamePrompt string `json:"defaultUsernamePrompt,omitempty"` // Default Username
	DefaultPasswordPrompt string `json:"defaultPasswordPrompt,omitempty"` // Default Password
}
type ResponseSystemSettingsCustomPromptPOSTApI struct {
	Response *ResponseSystemSettingsCustomPromptPOSTAPIResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptPOSTAPIResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type RequestSystemSettingsCustomPromptPOSTApI struct {
	UsernamePrompt string `json:"usernamePrompt,omitempty"` // Username Prompt
	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password Prompt
}

//CustomPromptSupportGETApI Custom-prompt support GET API - 2aa8-f90e-4ebb-a629
/* Returns supported custom prompts by DNAC


 */
func (s *SystemSettingsService) CustomPromptSupportGETApI() (*ResponseSystemSettingsCustomPromptSupportGETApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsCustomPromptSupportGETApI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CustomPromptSupportGETApI")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptSupportGETApI)
	return result, response, err

}

//CustomPromptPOSTApI Custom Prompt POST API - f4b9-1a8a-4718-aa97
/* Save custom prompt added by user in DNAC . API will always override the existing prompts. User should provide all the custom prompt in case of any update


 */
func (s *SystemSettingsService) CustomPromptPOSTApI(requestSystemSettingsCustomPromptPOSTAPI *RequestSystemSettingsCustomPromptPOSTApI) (*ResponseSystemSettingsCustomPromptPOSTApI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsCustomPromptPOSTAPI).
		SetResult(&ResponseSystemSettingsCustomPromptPOSTApI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CustomPromptPOSTApI")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptPOSTApI)
	return result, response, err

}
