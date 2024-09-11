package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SystemSettingsService service

type GetAuthenticationAndPolicyServersQueryParams struct {
	IsIseEnabled bool   `url:"isIseEnabled,omitempty"` //Valid values are : true, false
	State        string `url:"state,omitempty"`        //Valid values are: ACTIVE, INACTIVE, RBAC_SUCCESS, RBAC_FAILURE, DELETED, FAILED, INPROGRESS
	Role         string `url:"role,omitempty"`         //Authentication and Policy Server Role (Example: primary, secondary)
}

type ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration struct {
	Response *ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
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
	AuthenticationPort   *int                                                                           `json:"authenticationPort,omitempty"`   // Authentication port of RADIUS server (Default: 1812)
	AccountingPort       *int                                                                           `json:"accountingPort,omitempty"`       // Accounting port of RADIUS server (Default: 1813)
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
	RbacUUID             string                                                                         `json:"rbacUuid,omitempty"`             // Internal use only
	MultiDnacEnabled     *bool                                                                          `json:"multiDnacEnabled,omitempty"`     // Internal use only
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
type ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfiguration struct {
	Response *ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationResponse `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfigurationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration struct {
	Response *ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationResponse `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}

type ResponseSystemSettingsCiscoIseServerIntegrationStatus struct {
	AAAServerSettingID  string                                                        `json:"aaaServerSettingId,omitempty"`  // Cisco ISE Server setting identifier (E.g. 867e46c9-f8f5-40b1-8de2-62f7744f75f6)
	OverallStatus       string                                                        `json:"overallStatus,omitempty"`       // Cisco ISE Server integration status
	OverallErrorMessage string                                                        `json:"overallErrorMessage,omitempty"` // Cisco ISE Server integration failure message
	Steps               *[]ResponseSystemSettingsCiscoIseServerIntegrationStatusSteps `json:"steps,omitempty"`               //
}
type ResponseSystemSettingsCiscoIseServerIntegrationStatusSteps struct {
	StepID             string `json:"stepId,omitempty"`             // Cisco ISE Server integration step identifier (E.g. 1)
	StepOrder          *int   `json:"stepOrder,omitempty"`          // Cisco ISE Server integration step order (E.g. 1)
	StepName           string `json:"stepName,omitempty"`           // Cisco ISE Server integration step name
	StepDescription    string `json:"stepDescription,omitempty"`    // Cisco ISE Server step description
	StepStatus         string `json:"stepStatus,omitempty"`         // Cisco ISE Server integration step status
	CertAcceptedByUser *bool  `json:"certAcceptedByUser,omitempty"` // If user accept Cisco ISE Server certificate, value will be true otherwise it will be false
	StepTime           *int   `json:"stepTime,omitempty"`           // Last updated epoc time  by the step (E.g. 1677745739314)
}
type ResponseSystemSettingsCustomPromptSupportGETAPI struct {
	Response *ResponseSystemSettingsCustomPromptSupportGETAPIResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptSupportGETAPIResponse struct {
	CustomUsernamePrompt  string `json:"customUsernamePrompt,omitempty"`  // Username for Custom Prompt
	CustomPasswordPrompt  string `json:"customPasswordPrompt,omitempty"`  // Password for Custom Prompt
	DefaultUsernamePrompt string `json:"defaultUsernamePrompt,omitempty"` // Default Username for Custom Prompt
	DefaultPasswordPrompt string `json:"defaultPasswordPrompt,omitempty"` // Default Password for Custom Prompt
}
type ResponseSystemSettingsCustomPromptPOSTAPI struct {
	Response *ResponseSystemSettingsCustomPromptPOSTAPIResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCustomPromptPOSTAPIResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsSetProvisioningSettings struct {
	Version  string                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSystemSettingsSetProvisioningSettingsResponse `json:"response,omitempty"` //
}
type ResponseSystemSettingsSetProvisioningSettingsResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSystemSettingsGetProvisioningSettings struct {
	Response *ResponseSystemSettingsGetProvisioningSettingsResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseSystemSettingsGetProvisioningSettingsResponse struct {
	RequireItsmApproval *bool `json:"requireItsmApproval,omitempty"` // If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
	RequirePreview      *bool `json:"requirePreview,omitempty"`      // If require preview is enabled, the device configurations must be reviewed before deploying them
}
type ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer struct {
	Response *ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServerResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServer struct {
	Response *ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerResponse `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServerResponse struct {
	Provider   string `json:"provider,omitempty"`   // Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.
	ServerName string `json:"serverName,omitempty"` // A descriptive name of this external server, used for identification purposes
	ServerURL  string `json:"serverUrl,omitempty"`  // The URL of this external server
	State      string `json:"state,omitempty"`      // State of the the external IPAM.* OK indicates success of most recent periodic communication check with external IPAM.* CRITICAL indicates failure of most recent attempt to communicate with the external IPAM.* SYNCHRONIZING indicates that the process of synchronizing the external IPAM database with the local IPAM database is running and all other IPAM processes will be blocked until the completes.* DISCONNECTED indicates the external IPAM is no longer being used.
	UserName   string `json:"userName,omitempty"`   // The external IPAM server login username
	View       string `json:"view,omitempty"`       // The view under which pools are created in the external IPAM server.
}
type ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServer struct {
	Response *ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServerResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer struct {
	Response *ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServerResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration struct {
	AuthenticationPort         *int                                                                                                  `json:"authenticationPort,omitempty"`         // Authentication port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1812
	AccountingPort             *int                                                                                                  `json:"accountingPort,omitempty"`             // Accounting port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1813
	CiscoIseDtos               *[]RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationCiscoIseDtos               `json:"ciscoIseDtos,omitempty"`               //
	IPAddress                  string                                                                                                `json:"ipAddress,omitempty"`                  // IP address of authentication and policy server (readonly)
	PxgridEnabled              *bool                                                                                                 `json:"pxgridEnabled,omitempty"`              // Value true for enable, false for disable. Default value is true
	UseDnacCertForPxgrid       *bool                                                                                                 `json:"useDnacCertForPxgrid,omitempty"`       // Value true to use DNAC certificate for Pxgrid. Default value is false
	IsIseEnabled               *bool                                                                                                 `json:"isIseEnabled,omitempty"`               // Value true for Cisco ISE Server (readonly). Default value is false
	Port                       *int                                                                                                  `json:"port,omitempty"`                       // Port of TACACS server (readonly). The range is from 1 to 65535
	Protocol                   string                                                                                                `json:"protocol,omitempty"`                   // Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS
	Retries                    string                                                                                                `json:"retries,omitempty"`                    // Number of communication retries between devices and authentication and policy server. The range is from 1 to 3
	Role                       string                                                                                                `json:"role,omitempty"`                       // Role of authentication and policy server (readonly). E.g. primary, secondary
	SharedSecret               string                                                                                                `json:"sharedSecret,omitempty"`               // Shared secret between devices and authentication and policy server (readonly)
	TimeoutSeconds             string                                                                                                `json:"timeoutSeconds,omitempty"`             // Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20
	EncryptionScheme           string                                                                                                `json:"encryptionScheme,omitempty"`           // Type of encryption scheme for additional security (readonly)
	MessageKey                 string                                                                                                `json:"messageKey,omitempty"`                 // Message key used to encrypt shared secret (readonly)
	EncryptionKey              string                                                                                                `json:"encryptionKey,omitempty"`              // Encryption key used to encrypt shared secret (readonly)
	ExternalCiscoIseIPAddrDtos *[]RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtos `json:"externalCiscoIseIpAddrDtos,omitempty"` //
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationCiscoIseDtos struct {
	Description    string `json:"description,omitempty"`    // Description about the Cisco ISE server
	Fqdn           string `json:"fqdn,omitempty"`           // Fully-qualified domain name of the Cisco ISE server (readonly). E.g. xi-62.my.com
	Password       string `json:"password,omitempty"`       // Password of the Cisco ISE server
	SSHkey         string `json:"sshkey,omitempty"`         // SSH key of the Cisco ISE server
	IPAddress      string `json:"ipAddress,omitempty"`      // IP Address of the Cisco ISE Server (readonly)
	SubscriberName string `json:"subscriberName,omitempty"` // Subscriber name of the Cisco ISE server (readonly). E.g. pxgrid_client_1662589467
	UserName       string `json:"userName,omitempty"`       // User name of the Cisco ISE server
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtos struct {
	ExternalCiscoIseIPAddresses *[]RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses `json:"externalCiscoIseIpAddresses,omitempty"` //
	Type                        string                                                                                                                           `json:"type,omitempty"`                        // Type
}
type RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses struct {
	ExternalIPAddress string `json:"externalIpAddress,omitempty"` // External IP Address
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration struct {
	AuthenticationPort         *int                                                                                                   `json:"authenticationPort,omitempty"`         // Authentication port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1812
	AccountingPort             *int                                                                                                   `json:"accountingPort,omitempty"`             // Accounting port of RADIUS server (readonly). The range is from 1 to 65535. E.g. 1813
	CiscoIseDtos               *[]RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationCiscoIseDtos               `json:"ciscoIseDtos,omitempty"`               //
	IPAddress                  string                                                                                                 `json:"ipAddress,omitempty"`                  // IP address of authentication and policy server (readonly)
	PxgridEnabled              *bool                                                                                                  `json:"pxgridEnabled,omitempty"`              // Value true for enable, false for disable. Default value is true
	UseDnacCertForPxgrid       *bool                                                                                                  `json:"useDnacCertForPxgrid,omitempty"`       // Value true to use DNAC certificate for Pxgrid. Default value is false
	IsIseEnabled               *bool                                                                                                  `json:"isIseEnabled,omitempty"`               // Value true for Cisco ISE Server (readonly). Default value is false
	Port                       *int                                                                                                   `json:"port,omitempty"`                       // Port of TACACS server (readonly). The range is from 1 to 65535
	Protocol                   string                                                                                                 `json:"protocol,omitempty"`                   // Type of protocol for authentication and policy server. If already saved with RADIUS, can update to RADIUS_TACACS. If already saved with TACACS, can update to RADIUS_TACACS
	Retries                    string                                                                                                 `json:"retries,omitempty"`                    // Number of communication retries between devices and authentication and policy server. The range is from 1 to 3
	Role                       string                                                                                                 `json:"role,omitempty"`                       // Role of authentication and policy server (readonly). E.g. primary, secondary
	SharedSecret               string                                                                                                 `json:"sharedSecret,omitempty"`               // Shared secret between devices and authentication and policy server (readonly)
	TimeoutSeconds             string                                                                                                 `json:"timeoutSeconds,omitempty"`             // Number of seconds before timing out between devices and authentication and policy server. The range is from 2 to 20
	EncryptionScheme           string                                                                                                 `json:"encryptionScheme,omitempty"`           // Type of encryption scheme for additional security (readonly)
	MessageKey                 string                                                                                                 `json:"messageKey,omitempty"`                 // Message key used to encrypt shared secret (readonly)
	EncryptionKey              string                                                                                                 `json:"encryptionKey,omitempty"`              // Encryption key used to encrypt shared secret (readonly)
	ExternalCiscoIseIPAddrDtos *[]RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtos `json:"externalCiscoIseIpAddrDtos,omitempty"` //
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationCiscoIseDtos struct {
	Description    string `json:"description,omitempty"`    // Description about the Cisco ISE server
	Fqdn           string `json:"fqdn,omitempty"`           // Fully-qualified domain name of the Cisco ISE server (readonly). E.g. xi-62.my.com
	Password       string `json:"password,omitempty"`       // Password of the Cisco ISE server
	SSHkey         string `json:"sshkey,omitempty"`         // SSH key of the Cisco ISE server
	IPAddress      string `json:"ipAddress,omitempty"`      // IP Address of the Cisco ISE Server (readonly)
	SubscriberName string `json:"subscriberName,omitempty"` // Subscriber name of the Cisco ISE server (readonly). E.g. pxgrid_client_1662589467
	UserName       string `json:"userName,omitempty"`       // User name of the Cisco ISE server
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtos struct {
	ExternalCiscoIseIPAddresses *[]RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses `json:"externalCiscoIseIpAddresses,omitempty"` //
	Type                        string                                                                                                                            `json:"type,omitempty"`                        // Type
}
type RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfigurationExternalCiscoIseIPAddrDtosExternalCiscoIseIPAddresses struct {
	ExternalIPAddress string `json:"externalIpAddress,omitempty"` // External IP Address
}
type RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegration struct {
	IsCertAcceptedByUser *bool `json:"isCertAcceptedByUser,omitempty"` // Value true for accept, false for deny. Remove this field and send empty request payload ( {} ) to retry the failed integration
}
type RequestSystemSettingsCustomPromptPOSTAPI struct {
	UsernamePrompt string `json:"usernamePrompt,omitempty"` // Username for Custom Prompt
	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password for Custom Prompt
}
type RequestSystemSettingsSetProvisioningSettings struct {
	RequireItsmApproval *bool `json:"requireItsmApproval,omitempty"` // If require ITSM approval is enabled, the planned configurations must be submitted for ITSM approval. Also if enabled, requirePreview will default to enabled.
	RequirePreview      *bool `json:"requirePreview,omitempty"`      // If require preview is enabled, the device configurations must be reviewed before deploying them
}
type RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer struct {
	ServerName string `json:"serverName,omitempty"` // A descriptive name of this external server, used for identification purposes
	ServerURL  string `json:"serverUrl,omitempty"`  // The URL of this external server
	Password   string `json:"password,omitempty"`   // The password for the external IPAM server login username
	UserName   string `json:"userName,omitempty"`   // The external IPAM server login username
	Provider   string `json:"provider,omitempty"`   // Type of external IPAM. Can be either INFOBLOX, BLUECAT or GENERIC.
	View       string `json:"view,omitempty"`       // The view under which pools are created in the external IPAM server.
	SyncView   *bool  `json:"syncView,omitempty"`   // Synchronize the IP pools from the local IPAM to this external server
}
type RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer struct {
	ServerName string `json:"serverName,omitempty"` // A descriptive name of this external server, used for identification purposes
	ServerURL  string `json:"serverUrl,omitempty"`  // The URL of this external server
	Password   string `json:"password,omitempty"`   // The password for the external IPAM server login username
	UserName   string `json:"userName,omitempty"`   // The external IPAM server login username
	View       string `json:"view,omitempty"`       // The view under which pools are created in the external IPAM server.
	SyncView   *bool  `json:"syncView,omitempty"`   // Synchronize the IP pools from the local IPAM to this external server
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuthenticationAndPolicyServers(GetAuthenticationAndPolicyServersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuthenticationAndPolicyServers")
	}

	result := response.Result().(*ResponseSystemSettingsGetAuthenticationAndPolicyServers)
	return result, response, err

}

//CiscoIseServerIntegrationStatus Cisco ISE Server Integration Status - c1a4-f8fb-448a-8135
/* API to check Cisco ISE server integration status.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-ise-server-integration-status
*/
func (s *SystemSettingsService) CiscoIseServerIntegrationStatus() (*ResponseSystemSettingsCiscoIseServerIntegrationStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/ise-integration-status"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsCiscoIseServerIntegrationStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoIseServerIntegrationStatus()
		}
		return nil, response, fmt.Errorf("error with operation CiscoIseServerIntegrationStatus")
	}

	result := response.Result().(*ResponseSystemSettingsCiscoIseServerIntegrationStatus)
	return result, response, err

}

//CustomPromptSupportGETAPI Custom-prompt support GET API - 2aa8-f90e-4ebb-a629
/* Returns supported custom prompts by Catalyst Center



Documentation Link: https://developer.cisco.com/docs/dna-center/#!custom-prompt-support-g-e-t-api
*/
func (s *SystemSettingsService) CustomPromptSupportGETAPI() (*ResponseSystemSettingsCustomPromptSupportGETAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsCustomPromptSupportGETAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CustomPromptSupportGETAPI()
		}
		return nil, response, fmt.Errorf("error with operation CustomPromptSupportGETApi")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptSupportGETAPI)
	return result, response, err

}

//GetProvisioningSettings Get provisioning settings - b9b5-db54-4788-82e4
/* Returns provisioning settings



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioning-settings
*/
func (s *SystemSettingsService) GetProvisioningSettings() (*ResponseSystemSettingsGetProvisioningSettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/provisioningSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsGetProvisioningSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisioningSettings()
		}
		return nil, response, fmt.Errorf("error with operation GetProvisioningSettings")
	}

	result := response.Result().(*ResponseSystemSettingsGetProvisioningSettings)
	return result, response, err

}

//RetrievesConfigurationDetailsOfTheExternalIPAMServer Retrieves configuration details of the external IPAM server. - 3ebf-1bc3-4c8a-95e4
/* Retrieves configuration details of the external IPAM server.  If an external IPAM server has not been created, this resource will return a '404' response.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-configuration-details-of-the-external-ip-a-m-server
*/
func (s *SystemSettingsService) RetrievesConfigurationDetailsOfTheExternalIPAMServer() (*ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServer, *resty.Response, error) {
	path := "/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServer{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesConfigurationDetailsOfTheExternalIPAMServer()
		}
		return nil, response, fmt.Errorf("error with operation RetrievesConfigurationDetailsOfTheExternalIpAMServer")
	}

	result := response.Result().(*ResponseSystemSettingsRetrievesConfigurationDetailsOfTheExternalIPAMServer)
	return result, response, err

}

//AddAuthenticationAndPolicyServerAccessConfiguration Add Authentication and Policy Server Access Configuration - 5282-78a3-4fbb-a82c
/* API to add AAA/ISE server access configuration. Protocol can be configured as either RADIUS OR TACACS OR RADIUS_TACACS. If configuring Cisco ISE server, after configuration, use ‘Cisco ISE Server Integration Status’ Intent API to check the integration status. Based on integration status, if require use 'Accept Cisco ISE Server Certificate for Cisco ISE Server Integration' Intent API to accept the Cisco ISE certificate for Cisco ISE server integration, then use again ‘Cisco ISE Server Integration Status’ Intent API to check the integration status.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-authentication-and-policy-server-access-configuration
*/
func (s *SystemSettingsService) AddAuthenticationAndPolicyServerAccessConfiguration(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration *RequestSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration) (*ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration, *resty.Response, error) {
	path := "/dna/intent/api/v1/authentication-policy-servers"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration).
		SetResult(&ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAuthenticationAndPolicyServerAccessConfiguration(requestSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration)
		}

		return nil, response, fmt.Errorf("error with operation AddAuthenticationAndPolicyServerAccessConfiguration")
	}

	result := response.Result().(*ResponseSystemSettingsAddAuthenticationAndPolicyServerAccessConfiguration)
	return result, response, err

}

//CustomPromptPOSTAPI Custom Prompt POST API - f4b9-1a8a-4718-aa97
/* Save custom prompt added by user in Catalyst Center. API will always override the existing prompts. User should provide all the custom prompt in case of any update



Documentation Link: https://developer.cisco.com/docs/dna-center/#!custom-prompt-p-o-s-t-api
*/
func (s *SystemSettingsService) CustomPromptPOSTAPI(requestSystemSettingsCustomPromptPOSTAPI *RequestSystemSettingsCustomPromptPOSTAPI) (*ResponseSystemSettingsCustomPromptPOSTAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/custom-prompt"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsCustomPromptPOSTAPI).
		SetResult(&ResponseSystemSettingsCustomPromptPOSTAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CustomPromptPOSTAPI(requestSystemSettingsCustomPromptPOSTAPI)
		}

		return nil, response, fmt.Errorf("error with operation CustomPromptPostAPI")
	}

	result := response.Result().(*ResponseSystemSettingsCustomPromptPOSTAPI)
	return result, response, err

}

//CreatesConfigurationDetailsOfTheExternalIPAMServer Creates configuration details of the external IPAM server. - 36a4-38d6-4589-8f87
/* Creates configuration details of the external IPAM server. You should only create one external IPAM server; delete any existing external server before creating a new one.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-configuration-details-of-the-external-ip-a-m-server
*/
func (s *SystemSettingsService) CreatesConfigurationDetailsOfTheExternalIPAMServer(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer *RequestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer) (*ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer, *resty.Response, error) {
	path := "/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer).
		SetResult(&ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesConfigurationDetailsOfTheExternalIPAMServer(requestSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer)
		}

		return nil, response, fmt.Errorf("error with operation CreatesConfigurationDetailsOfTheExternalIpAMServer")
	}

	result := response.Result().(*ResponseSystemSettingsCreatesConfigurationDetailsOfTheExternalIPAMServer)
	return result, response, err

}

//EditAuthenticationAndPolicyServerAccessConfiguration Edit Authentication and Policy Server Access Configuration - e4bf-3bf1-48c9-ba11
/* API to edit AAA/ISE server access configuration. After edit, use ‘Cisco ISE Server Integration Status’ Intent API to check the integration status.


@param id id path parameter. Authentication and Policy Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.

*/
func (s *SystemSettingsService) EditAuthenticationAndPolicyServerAccessConfiguration(id string, requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration *RequestSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration) (*ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration, *resty.Response, error) {
	path := "/dna/intent/api/v1/authentication-policy-servers/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration).
		SetResult(&ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditAuthenticationAndPolicyServerAccessConfiguration(id, requestSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration)
		}
		return nil, response, fmt.Errorf("error with operation EditAuthenticationAndPolicyServerAccessConfiguration")
	}

	result := response.Result().(*ResponseSystemSettingsEditAuthenticationAndPolicyServerAccessConfiguration)
	return result, response, err

}

//AcceptCiscoIseServerCertificateForCiscoIseServerIntegration Accept Cisco ISE Server Certificate for Cisco ISE Server Integration - c8a6-aae2-48b8-b41c
/* API to accept Cisco ISE server certificate for Cisco ISE server integration. Use ‘Cisco ISE Server Integration Status’ Intent API to check the integration status. This API can be used to retry the failed integration.


@param id id path parameter. Cisco ISE Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.

*/
func (s *SystemSettingsService) AcceptCiscoIseServerCertificateForCiscoIseServerIntegration(id string, requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegration *RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegration) (*resty.Response, error) {
	path := "/dna/intent/api/v1/integrate-ise/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegration).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AcceptCiscoIseServerCertificateForCiscoIseServerIntegration(id, requestSystemSettingsAcceptCiscoISEServerCertificateForCiscoISEServerIntegration)
		}
		return response, fmt.Errorf("error with operation AcceptCiscoIseServerCertificateForCiscoIseServerIntegration")
	}

	return response, err

}

//SetProvisioningSettings Set provisioning settings - e5ab-1811-450a-bb01
/* Sets provisioning settings


 */
func (s *SystemSettingsService) SetProvisioningSettings(requestSystemSettingsSetProvisioningSettings *RequestSystemSettingsSetProvisioningSettings) (*ResponseSystemSettingsSetProvisioningSettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/provisioningSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsSetProvisioningSettings).
		SetResult(&ResponseSystemSettingsSetProvisioningSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetProvisioningSettings(requestSystemSettingsSetProvisioningSettings)
		}
		return nil, response, fmt.Errorf("error with operation SetProvisioningSettings")
	}

	result := response.Result().(*ResponseSystemSettingsSetProvisioningSettings)
	return result, response, err

}

//UpdatesConfigurationDetailsOfTheExternalIPAMServer Updates configuration details of the external IPAM server. - 5a99-0bfe-4c99-a0a4
/* Updates configuration details of the external IPAM server.


 */
func (s *SystemSettingsService) UpdatesConfigurationDetailsOfTheExternalIPAMServer(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer *RequestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer) (*ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer, *resty.Response, error) {
	path := "/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer).
		SetResult(&ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesConfigurationDetailsOfTheExternalIPAMServer(requestSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesConfigurationDetailsOfTheExternalIpAMServer")
	}

	result := response.Result().(*ResponseSystemSettingsUpdatesConfigurationDetailsOfTheExternalIPAMServer)
	return result, response, err

}

//DeleteAuthenticationAndPolicyServerAccessConfiguration Delete Authentication and Policy Server Access Configuration - 0b92-bb8a-477a-a942
/* API to delete AAA/ISE server access configuration.


@param id id path parameter. Authentication and Policy Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-authentication-and-policy-server-access-configuration
*/
func (s *SystemSettingsService) DeleteAuthenticationAndPolicyServerAccessConfiguration(id string) (*ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfiguration, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/authentication-policy-servers/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfiguration{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAuthenticationAndPolicyServerAccessConfiguration(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAuthenticationAndPolicyServerAccessConfiguration")
	}

	result := response.Result().(*ResponseSystemSettingsDeleteAuthenticationAndPolicyServerAccessConfiguration)
	return result, response, err

}

//DeletesConfigurationDetailsOfTheExternalIPAMServer Deletes configuration details of the external IPAM server. - 67b6-eb01-4688-a164
/* Deletes configuration details of the external IPAM server.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-configuration-details-of-the-external-ip-a-m-server
*/
func (s *SystemSettingsService) DeletesConfigurationDetailsOfTheExternalIPAMServer() (*ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServer, *resty.Response, error) {
	//
	path := "/intent/api/v1/ipam/serverSetting"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServer{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesConfigurationDetailsOfTheExternalIPAMServer()
		}
		return nil, response, fmt.Errorf("error with operation DeletesConfigurationDetailsOfTheExternalIpAMServer")
	}

	result := response.Result().(*ResponseSystemSettingsDeletesConfigurationDetailsOfTheExternalIPAMServer)
	return result, response, err

}
