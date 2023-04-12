package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SystemSettingsService service

type GetAuthenticationAndPolicyServersQueryParams struct {
	IsIseEnabled bool   `url:"isIseEnabled,omitempty"` //Valid values are : true, false
	State        string `url:"state,omitempty"`        //Valid values are: INPROGRESS, ACTIVE, DELETED, RBAC-FAILURE, FAILED
	Role         string `url:"role,omitempty"`         //Authentication and Policy Server Role (Example: primary, secondary)
}

type ResponseSystemSettingsGetAuthenticationAndPolicyServers struct {
	Response *[]ResponseSystemSettingsGetAuthenticationAndPolicyServersResponse `json:"response,omitempty"` //
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersResponse struct {
	IPAddress            string                                                                         `json:"ipAddress,omitempty"`            // IP address of authentication and policy server
	SharedSecret         string                                                                         `json:"sharedSecret,omitempty"`         // Shared secret between devices and authentication and policy server
	Protocol             string                                                                         `json:"protocol,omitempty"`             // Type of protocol for authentication and policy server
	Role                 string                                                                         `json:"role,omitempty"`                 // Role of authentication and policy server (Example: primary, secondary)
	Port                 *int                                                                           `json:"port,omitempty"`                 // Port of TACACS server (Default: 49)
	AuthenticationPort   string                                                                         `json:"authenticationPort,omitempty"`   // Authentication port of RADIUS server (Default: 1812)
	AccountingPort       string                                                                         `json:"accountingPort,omitempty"`       // Accounting port of RADIUS server (Default: 1813)
	Retries              *int                                                                           `json:"retries,omitempty"`              // Number of communication retries between devices and authentication and policy server (Default: 3)
	TimeoutSeconds       *int                                                                           `json:"timeoutSeconds,omitempty"`       // Number of seconds before timing out between devices and authentication and policy server (Default: 4 seconds)
	IsIseEnabled         *bool                                                                          `json:"isIseEnabled,omitempty"`         // If server type is ISE, value will be true otherwise false
	InstanceUUID         string                                                                         `json:"instanceUuid,omitempty"`         // Internal record identifier
	State                string                                                                         `json:"state,omitempty"`                // State of authentication and policy server
	CiscoIseDtos         *[]ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtos `json:"ciscoIseDtos,omitempty"`         //
	EncryptionScheme     string                                                                         `json:"encryptionScheme,omitempty"`     // Type of encryption scheme for additional security
	MessageKey           string                                                                         `json:"messageKey,omitempty"`           // Message key used to encrypt shared secret
	EncryptionKey        string                                                                         `json:"encryptionKey,omitempty"`        // Encryption key used to encrypt shared secret
	UseDnacCertForPxgrid *bool                                                                          `json:"useDnacCertForPxgrid,omitempty"` // Use DNAC Certificate For Pxgrid
	IseEnabled           *bool                                                                          `json:"iseEnabled,omitempty"`           // If server type is ISE, value will be true otherwise false
	PxgridEnabled        *bool                                                                          `json:"pxgridEnabled,omitempty"`        // If pxgrid enabled, value will be true otherwise false
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtos struct {
	SubscriberName             string                                                                                                 `json:"subscriberName,omitempty"`             // Subscriber name of the ISE server (Example: pxgrid_client_1662589467)
	Description                string                                                                                                 `json:"description,omitempty"`                // Description about the ISE server
	Password                   string                                                                                                 `json:"password,omitempty"`                   // For security reasons the value will always be empty
	UserName                   string                                                                                                 `json:"userName,omitempty"`                   // User name of the ISE server
	Fqdn                       string                                                                                                 `json:"fqdn,omitempty"`                       // Fully-qualified domain name of the ISE server (Example: xi-62.my.com)
	IPAddress                  string                                                                                                 `json:"ipAddress,omitempty"`                  // IP Address of the ISE server
	TrustState                 string                                                                                                 `json:"trustState,omitempty"`                 // Trust State between DNAC and the ISE server
	InstanceUUID               string                                                                                                 `json:"instanceUuid,omitempty"`               // Internal record identifier
	SSHkey                     string                                                                                                 `json:"sshkey,omitempty"`                     // For security reasons the value will always be empty
	Type                       string                                                                                                 `json:"type,omitempty"`                       // Type (Example: ISE)
	FailureReason              string                                                                                                 `json:"failureReason,omitempty"`              // Reason for integration failure between DNAC and the ISE server
	Role                       string                                                                                                 `json:"role,omitempty"`                       // Role of the ISE server
	ExternalCiscoIseIPAddrDtos *ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtosExternalCiscoIseIPAddrDtos `json:"externalCiscoIseIpAddrDtos,omitempty"` //
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtosExternalCiscoIseIPAddrDtos struct {
	Type                        string                                                                                                                              `json:"type,omitempty"`                        // Type
	ExternalCiscoIseIPAddresses *[]ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses `json:"externalCiscoIseIpAddresses,omitempty"` //
}
type ResponseSystemSettingsGetAuthenticationAndPolicyServersResponseCiscoIseDtosExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses struct {
	ExternalIPAddress string `json:"externalIpAddress,omitempty"` // External IP Address
}
type ResponseSystemSettingsCustomPromptSupportGetAPI struct {
	Response *ResponseSystemSettingsCustomPromptSupportGETAPIResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptSupportGETAPIResponse struct {
	CustomUsernamePrompt  string `json:"customUsernamePrompt,omitempty"`  // Custom Username
	CustomPasswordPrompt  string `json:"customPasswordPrompt,omitempty"`  // Custom Password
	DefaultUsernamePrompt string `json:"defaultUsernamePrompt,omitempty"` // Default Username
	DefaultPasswordPrompt string `json:"defaultPasswordPrompt,omitempty"` // Default Password
}
type ResponseSystemSettingsCustomPromptPostAPI struct {
	Response *ResponseSystemSettingsCustomPromptPOSTAPIResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptPOSTAPIResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type RequestSystemSettingsCustomPromptPostAPI struct {
	UsernamePrompt string `json:"usernamePrompt,omitempty"` // Username Prompt
	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password Prompt
}

//GetAuthenticationAndPolicyServers Get Authentication and Policy Servers - a4b4-c849-4be8-b362
/* API to get Authentication and Policy Servers


@param GetAuthenticationAndPolicyServersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-authentication-and-policy-servers
*/
func (s *SystemSettingsService) GetAuthenticationAndPolicyServers(GetAuthenticationAndPolicyServersQueryParams *GetAuthenticationAndPolicyServersQueryParams) (*ResponseSystemSettingsGetAuthenticationAndPolicyServers, *resty.Response, error) {
	path := "/dna/intent/api/v1/authentication-policy-servers"

	queryString, _ := query.Values(GetAuthenticationAndPolicyServersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSystemSettingsGetAuthenticationAndPolicyServers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuthenticationAndPolicyServers")
	}

	result := response.Result().(*ResponseSystemSettingsGetAuthenticationAndPolicyServers)
	return result, response, err

}

//CustomPromptSupportGetAPI Custom-prompt support GET API - 2aa8-f90e-4ebb-a629
/* Returns supported custom prompts by Cisco DNA Center



Documentation Link: https://developer.cisco.com/docs/dna-center/#!custom-prompt-support-get-api
*/
func (s *SystemSettingsService) CustomPromptSupportGetAPI() (*ResponseSystemSettingsCustomPromptSupportGetAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsCustomPromptSupportGetAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CustomPromptSupportGetApi")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptSupportGetAPI)
	return result, response, err

}

//CustomPromptPostAPI Custom Prompt POST API - f4b9-1a8a-4718-aa97
/* Save custom prompt added by user in Cisco DNA Center. API will always override the existing prompts. User should provide all the custom prompt in case of any update



Documentation Link: https://developer.cisco.com/docs/dna-center/#!custom-prompt-p-o-s-t-api
*/
func (s *SystemSettingsService) CustomPromptPostAPI(requestSystemSettingsCustomPromptPOSTAPI *RequestSystemSettingsCustomPromptPostAPI) (*ResponseSystemSettingsCustomPromptPostAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsCustomPromptPOSTAPI).
		SetResult(&ResponseSystemSettingsCustomPromptPostAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CustomPromptPostApi")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptPostAPI)
	return result, response, err

}
