package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type UserandRolesService service

type GetRolesAPIHeaderParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //Expects type string. The source that invokes this API. The value of this header must be set to "external".
}
type GetUsersAPIQueryParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //The source that invokes this API. The value of this query parameter must be set to "external".
	AuthSource   string `url:"authSource,omitempty"`   //The source that authenticates the user. The value of this query parameter can be set to "internal" or "external". If not provided, then all users will be returned in the response.
}
type GetExternalAuthenticationServersAPIQueryParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //The source that invokes this API. The value of this query parameter must be set to "external".
}

type ResponseUserandRolesAddRoleAPI struct {
	Response *ResponseUserandRolesAddRoleAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesAddRoleAPIResponse struct {
	RoleID string `json:"roleId,omitempty"` // Role Id

	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesUpdateRoleAPI struct {
	Response *ResponseUserandRolesUpdateRoleAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesUpdateRoleAPIResponse struct {
	RoleID string `json:"roleId,omitempty"` // Role Id

	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetPermissionsAPI struct {
	Response *ResponseUserandRolesGetPermissionsAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetPermissionsAPIResponse struct {
	ResourceTypes *[]ResponseUserandRolesGetPermissionsAPIResponseResourceTypes `json:"resource-types,omitempty"` //
}
type ResponseUserandRolesGetPermissionsAPIResponseResourceTypes struct {
	Type string `json:"type,omitempty"` // Type

	DisplayName string `json:"displayName,omitempty"` // Display Name

	Description string `json:"description,omitempty"` // Description

	DefaultPermission string `json:"defaultPermission,omitempty"` // Default permission
}
type ResponseUserandRolesDeleteRoleAPI struct {
	Response *ResponseUserandRolesDeleteRoleAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesDeleteRoleAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetRolesAPI struct {
	Response *ResponseUserandRolesGetRolesAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetRolesAPIResponse struct {
	Roles *[]ResponseUserandRolesGetRolesAPIResponseRoles `json:"roles,omitempty"` //
}
type ResponseUserandRolesGetRolesAPIResponseRoles struct {
	ResourceTypes *[]ResponseUserandRolesGetRolesAPIResponseRolesResourceTypes `json:"resourceTypes,omitempty"` //

	Meta *ResponseUserandRolesGetRolesAPIResponseRolesMeta `json:"meta,omitempty"` //

	RoleID string `json:"roleId,omitempty"` // Role Id

	Name string `json:"name,omitempty"` // Role name

	Description string `json:"description,omitempty"` // Description

	Type string `json:"type,omitempty"` // Role type, possible values are: "DEFAULT", "SYSTEM", "CUSTOM"
}
type ResponseUserandRolesGetRolesAPIResponseRolesResourceTypes struct {
	Operations []string `json:"operations,omitempty"` // Operations

	Type string `json:"type,omitempty"` // Type
}
type ResponseUserandRolesGetRolesAPIResponseRolesMeta struct {
	CreatedBy string `json:"createdBy,omitempty"` // The user that creates the resource type

	Created string `json:"created,omitempty"` // The timestamp that the resource type was created

	LastModified string `json:"lastModified,omitempty"` // The latestest timestamp that the resource type was updated
}
type ResponseUserandRolesGetUsersAPI struct {
	Response *ResponseUserandRolesGetUsersAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetUsersAPIResponse struct {
	Users *[]ResponseUserandRolesGetUsersAPIResponseUsers `json:"users,omitempty"` //
}
type ResponseUserandRolesGetUsersAPIResponseUsers struct {
	FirstName string `json:"firstName,omitempty"` // First Name

	LastName string `json:"lastName,omitempty"` // Last Name

	AuthSource string `json:"authSource,omitempty"` // Authentiction source, internal or external

	PassphraseUpdateTime string `json:"passphraseUpdateTime,omitempty"` // Passphrase Update Time

	RoleList []string `json:"roleList,omitempty"` // A list of role ids

	UserID string `json:"userId,omitempty"` // User Id

	Email string `json:"email,omitempty"` // Email

	Username string `json:"username,omitempty"` // Username
}
type ResponseUserandRolesAddUserAPI struct {
	Response *ResponseUserandRolesAddUserAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesAddUserAPIResponse struct {
	Message string `json:"message,omitempty"` // Message

	UserID string `json:"userId,omitempty"` // User Id
}
type ResponseUserandRolesUpdateUserAPI struct {
	Response *ResponseUserandRolesUpdateUserAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesUpdateUserAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesDeleteUserAPI struct {
	Response *ResponseUserandRolesDeleteUserAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesDeleteUserAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetExternalAuthenticationSettingAPI struct {
	Response *ResponseUserandRolesGetExternalAuthenticationSettingAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationSettingAPIResponse struct {
	ExternalAuthenticationFlag *[]ResponseUserandRolesGetExternalAuthenticationSettingAPIResponseExternalAuthenticationFlag `json:"external-authentication-flag,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationSettingAPIResponseExternalAuthenticationFlag struct {
	Enabled *bool `json:"enabled,omitempty"` // External Authentication is enabled/disabled.
}
type ResponseUserandRolesManageExternalAuthenticationSettingAPI struct {
	Response *ResponseUserandRolesManageExternalAuthenticationSettingAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesManageExternalAuthenticationSettingAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetExternalAuthenticationServersAPI struct {
	Response *ResponseUserandRolesGetExternalAuthenticationServersAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIResponse struct {
	AAAServers *[]ResponseUserandRolesGetExternalAuthenticationServersAPIResponseAAAServers `json:"aaa-servers,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIResponseAAAServers struct {
	AccountingPort *int `json:"accountingPort,omitempty"` // RADIUS server accounting requests port

	Retries *int `json:"retries,omitempty"` // Retries

	Protocol string `json:"protocol,omitempty"` // Protocol

	SocketTimeout *int `json:"socketTimeout,omitempty"` // Timeout in seconds

	ServerIP string `json:"serverIp,omitempty"` // Server Ip

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared Secret

	ServerID string `json:"serverId,omitempty"` // Server Id

	AuthenticationPort *int `json:"authenticationPort,omitempty"` // RADIUS server authorization requests port

	AAAAttribute string `json:"aaaAttribute,omitempty"` // Aaa Attribute

	Role string `json:"role,omitempty"` // Role of AAA server, primary or secondary server
}
type ResponseUserandRolesAddAndUpdateAAAAttributeAPI struct {
	Response *ResponseUserandRolesAddAndUpdateAAAAttributeAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesAddAndUpdateAAAAttributeAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesDeleteAAAAttributeAPI struct {
	Response *ResponseUserandRolesDeleteAAAAttributeAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesDeleteAAAAttributeAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetAAAAttributeAPI struct {
	Response *ResponseUserandRolesGetAAAAttributeAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetAAAAttributeAPIResponse struct {
	AAAAttributes *[]ResponseUserandRolesGetAAAAttributeAPIResponseAAAAttributes `json:"aaa-attributes,omitempty"` //
}
type ResponseUserandRolesGetAAAAttributeAPIResponseAAAAttributes struct {
	AttributeName string `json:"attributeName,omitempty"` // Value of the custom AAA attribute name
}
type RequestUserandRolesAddRoleAPI struct {
	Role string `json:"role,omitempty"` // Name of the role

	Description string `json:"description,omitempty"` // Description of role

	ResourceTypes *[]RequestUserandRolesAddRoleAPIResourceTypes `json:"resourceTypes,omitempty"` //
}
type RequestUserandRolesAddRoleAPIResourceTypes struct {
	Type string `json:"type,omitempty"` // Name of the application in the System

	Operations []string `json:"operations,omitempty"` // List of operations allowed for the application. Possible values are "gRead", "gCreate", "gUpdate", "gRemove", or some combination of these.
}
type RequestUserandRolesUpdateRoleAPI struct {
	RoleID string `json:"roleId,omitempty"` // Id of the role

	Description string `json:"description,omitempty"` // Description of the role

	ResourceTypes *[]RequestUserandRolesUpdateRoleAPIResourceTypes `json:"resourceTypes,omitempty"` //
}
type RequestUserandRolesUpdateRoleAPIResourceTypes struct {
	Type string `json:"type,omitempty"` // Name of application in the System

	Operations []string `json:"operations,omitempty"` // List of operations allowed for the application. Possible values are "gRead", "gCreate", "gUpdate", "gRemove", or some combination of these.
}
type RequestUserandRolesAddUserAPI struct {
	FirstName string `json:"firstName,omitempty"` // First Name

	LastName string `json:"lastName,omitempty"` // Last Name

	Username string `json:"username,omitempty"` // Username

	Password string `json:"password,omitempty"` // Password

	Email string `json:"email,omitempty"` // Email

	RoleList []string `json:"roleList,omitempty"` // Role id list
}
type RequestUserandRolesUpdateUserAPI struct {
	FirstName string `json:"firstName,omitempty"` // firstName should be set if the original value is not empty

	LastName string `json:"lastName,omitempty"` // lastName should be set if the original value is not empty

	Email string `json:"email,omitempty"` // email should be set if the original value is not empty

	Username string `json:"username,omitempty"` // Username

	UserID string `json:"userId,omitempty"` // User Id

	RoleList []string `json:"roleList,omitempty"` // Role id list
}
type RequestUserandRolesManageExternalAuthenticationSettingAPI struct {
	Enable *bool `json:"enable,omitempty"` // Enable/disable External Authentication.
}
type RequestUserandRolesAddAndUpdateAAAAttributeAPI struct {
	AttributeName string `json:"attributeName,omitempty"` // name of the custom AAA attribute.
}

//GetPermissionsAPI Get permissions API - 8a9c-6885-455b-a2db
/* Get permissions for a role in the system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-permissions-api
*/
func (s *UserandRolesService) GetPermissionsAPI() (*ResponseUserandRolesGetPermissionsAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/role/permissions"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetPermissionsAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPermissionsAPI()
		}
		return nil, response, fmt.Errorf("error with operation GetPermissionsApi")
	}

	result := response.Result().(*ResponseUserandRolesGetPermissionsAPI)
	return result, response, err

}

//GetRolesAPI Get roles API - 7c86-da3f-4b08-8593
/* Get all roles in the system


@param GetRolesAPIHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-roles-api
*/
func (s *UserandRolesService) GetRolesAPI(GetRolesAPIHeaderParams *GetRolesAPIHeaderParams) (*ResponseUserandRolesGetRolesAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/roles"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetRolesAPIHeaderParams != nil {

		if GetRolesAPIHeaderParams.InvokeSource != "" {
			clientRequest = clientRequest.SetHeader("invokeSource", GetRolesAPIHeaderParams.InvokeSource)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseUserandRolesGetRolesAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRolesAPI(GetRolesAPIHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetRolesApi")
	}

	result := response.Result().(*ResponseUserandRolesGetRolesAPI)
	return result, response, err

}

//GetUsersAPI Get users API - 918c-89fa-4a98-a528
/* Get all users in the system


@param GetUsersAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-users-api
*/
func (s *UserandRolesService) GetUsersAPI(GetUsersAPIQueryParams *GetUsersAPIQueryParams) (*ResponseUserandRolesGetUsersAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	queryString, _ := query.Values(GetUsersAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserandRolesGetUsersAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetUsersAPI(GetUsersAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetUsersApi")
	}

	result := response.Result().(*ResponseUserandRolesGetUsersAPI)
	return result, response, err

}

//GetExternalAuthenticationSettingAPI Get External Authentication Setting API - e0a8-aa75-49cb-815c
/* Get the External Authentication setting.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-external-authentication-setting-api
*/
func (s *UserandRolesService) GetExternalAuthenticationSettingAPI() (*ResponseUserandRolesGetExternalAuthenticationSettingAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-authentication"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetExternalAuthenticationSettingAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExternalAuthenticationSettingAPI()
		}
		return nil, response, fmt.Errorf("error with operation GetExternalAuthenticationSettingApi")
	}

	result := response.Result().(*ResponseUserandRolesGetExternalAuthenticationSettingAPI)
	return result, response, err

}

//GetExternalAuthenticationServersAPI Get external authentication servers API - 9dbd-0b01-4758-bde4
/* Get external users authentication servers.


@param GetExternalAuthenticationServersAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-external-authentication-servers-api
*/
func (s *UserandRolesService) GetExternalAuthenticationServersAPI(GetExternalAuthenticationServersAPIQueryParams *GetExternalAuthenticationServersAPIQueryParams) (*ResponseUserandRolesGetExternalAuthenticationServersAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers"

	queryString, _ := query.Values(GetExternalAuthenticationServersAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserandRolesGetExternalAuthenticationServersAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExternalAuthenticationServersAPI(GetExternalAuthenticationServersAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetExternalAuthenticationServersApi")
	}

	result := response.Result().(*ResponseUserandRolesGetExternalAuthenticationServersAPI)
	return result, response, err

}

//GetAAAAttributeAPI Get AAA Attribute API - 2eb5-ea84-4d29-bf8b
/* Get the current value of the custom AAA attribute.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-a-a-a-attribute-api
*/
func (s *UserandRolesService) GetAAAAttributeAPI() (*ResponseUserandRolesGetAAAAttributeAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers/aaa-attribute"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetAAAAttributeAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAAAAttributeAPI()
		}
		return nil, response, fmt.Errorf("error with operation GetAAAAttributeApi")
	}

	result := response.Result().(*ResponseUserandRolesGetAAAAttributeAPI)
	return result, response, err

}

//AddRoleAPI Add role API - b697-0a1e-46a9-b542
/* Add a new role in the system



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-role-api
*/
func (s *UserandRolesService) AddRoleAPI(requestUserandRolesAddRoleAPI *RequestUserandRolesAddRoleAPI) (*ResponseUserandRolesAddRoleAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/role"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddRoleAPI).
		SetResult(&ResponseUserandRolesAddRoleAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddRoleAPI(requestUserandRolesAddRoleAPI)
		}

		return nil, response, fmt.Errorf("error with operation AddRoleApi")
	}

	result := response.Result().(*ResponseUserandRolesAddRoleAPI)
	return result, response, err

}

//AddUserAPI Add user API - 6c9a-09c4-4a39-9e2b
/* Add a new user in the system



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-user-api
*/
func (s *UserandRolesService) AddUserAPI(requestUserandRolesAddUserAPI *RequestUserandRolesAddUserAPI) (*ResponseUserandRolesAddUserAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddUserAPI).
		SetResult(&ResponseUserandRolesAddUserAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUserAPI(requestUserandRolesAddUserAPI)
		}

		return nil, response, fmt.Errorf("error with operation AddUserApi")
	}

	result := response.Result().(*ResponseUserandRolesAddUserAPI)
	return result, response, err

}

//ManageExternalAuthenticationSettingAPI Manage External Authentication Setting API - e09c-1806-48da-bb40
/* Enable or disable external authentication in the System.
Please find the Administrator Guide for your particular release from the list linked below and follow the steps required to enable external authentication before trying to do so from this API.
https://www.cisco.com/c/en/us/support/cloud-systems-management/dna-center/products-maintenance-guides-list.html



Documentation Link: https://developer.cisco.com/docs/dna-center/#!manage-external-authentication-setting-api
*/
func (s *UserandRolesService) ManageExternalAuthenticationSettingAPI(requestUserandRolesManageExternalAuthenticationSettingAPI *RequestUserandRolesManageExternalAuthenticationSettingAPI) (*ResponseUserandRolesManageExternalAuthenticationSettingAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-authentication"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesManageExternalAuthenticationSettingAPI).
		SetResult(&ResponseUserandRolesManageExternalAuthenticationSettingAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ManageExternalAuthenticationSettingAPI(requestUserandRolesManageExternalAuthenticationSettingAPI)
		}

		return nil, response, fmt.Errorf("error with operation ManageExternalAuthenticationSettingApi")
	}

	result := response.Result().(*ResponseUserandRolesManageExternalAuthenticationSettingAPI)
	return result, response, err

}

//AddAndUpdateAAAAttributeAPI Add and Update AAA Attribute API - 808a-0aa0-491b-891f
/* Add or update the custom AAA attribute for external authentication. Note that if you decide not to set the custom AAA attribute, a default AAA attribute will be used for authentication based on the protocol supported by your server. For TACACS servers it will be "cisco-av-pair" and for RADIUS servers it will be "Cisco-AVPair".



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-and-update-a-a-a-attribute-api
*/
func (s *UserandRolesService) AddAndUpdateAAAAttributeAPI(requestUserandRolesAddAndUpdateAAAAttributeAPI *RequestUserandRolesAddAndUpdateAAAAttributeAPI) (*ResponseUserandRolesAddAndUpdateAAAAttributeAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers/aaa-attribute"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddAndUpdateAAAAttributeAPI).
		SetResult(&ResponseUserandRolesAddAndUpdateAAAAttributeAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAndUpdateAAAAttributeAPI(requestUserandRolesAddAndUpdateAAAAttributeAPI)
		}

		return nil, response, fmt.Errorf("error with operation AddAndUpdateAAAAttributeApi")
	}

	result := response.Result().(*ResponseUserandRolesAddAndUpdateAAAAttributeAPI)
	return result, response, err

}

//UpdateRoleAPI Update role API - 539c-ea73-400b-bf20
/* Update a role in the system


 */
func (s *UserandRolesService) UpdateRoleAPI(requestUserandRolesUpdateRoleAPI *RequestUserandRolesUpdateRoleAPI) (*ResponseUserandRolesUpdateRoleAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/role"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesUpdateRoleAPI).
		SetResult(&ResponseUserandRolesUpdateRoleAPI{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRoleAPI(requestUserandRolesUpdateRoleAPI)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRoleApi")
	}

	result := response.Result().(*ResponseUserandRolesUpdateRoleAPI)
	return result, response, err

}

//UpdateUserAPI Update user API - f596-6adc-492b-a2ff
/* Update a user in the system


 */
func (s *UserandRolesService) UpdateUserAPI(requestUserandRolesUpdateUserAPI *RequestUserandRolesUpdateUserAPI) (*ResponseUserandRolesUpdateUserAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesUpdateUserAPI).
		SetResult(&ResponseUserandRolesUpdateUserAPI{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateUserAPI(requestUserandRolesUpdateUserAPI)
		}
		return nil, response, fmt.Errorf("error with operation UpdateUserApi")
	}

	result := response.Result().(*ResponseUserandRolesUpdateUserAPI)
	return result, response, err

}

//DeleteRoleAPI Delete role API - d3b9-8bdc-472b-b236
/* Delete a role in the system


@param roleID roleId path parameter. The Id of the role to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-role-api
*/
func (s *UserandRolesService) DeleteRoleAPI(roleID string) (*ResponseUserandRolesDeleteRoleAPI, *resty.Response, error) {
	//roleID string
	path := "/dna/system/api/v1/role/{roleId}"
	path = strings.Replace(path, "{roleId}", fmt.Sprintf("%v", roleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesDeleteRoleAPI{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteRoleAPI(roleID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteRoleApi")
	}

	result := response.Result().(*ResponseUserandRolesDeleteRoleAPI)
	return result, response, err

}

//DeleteUserAPI Delete user API - 69b4-ba37-4aca-8e86
/* Delete a user in the system


@param userID userId path parameter. The id of the user to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-user-api
*/
func (s *UserandRolesService) DeleteUserAPI(userID string) (*ResponseUserandRolesDeleteUserAPI, *resty.Response, error) {
	//userID string
	path := "/dna/system/api/v1/user/{userId}"
	path = strings.Replace(path, "{userId}", fmt.Sprintf("%v", userID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesDeleteUserAPI{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteUserAPI(userID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteUserApi")
	}

	result := response.Result().(*ResponseUserandRolesDeleteUserAPI)
	return result, response, err

}

//DeleteAAAAttributeAPI Delete AAA Attribute API - d99e-c8df-4f1b-98fe
/* Delete the custom AAA attribute that was added. Note that by deleting the AAA attribute, a default AAA attribute will be used for authentication based on the protocol supported by your server. For TACACS servers it will be "cisco-av-pair" and for RADIUS servers it will be "Cisco-AVPair".



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-a-a-attribute-api
*/
func (s *UserandRolesService) DeleteAAAAttributeAPI() (*ResponseUserandRolesDeleteAAAAttributeAPI, *resty.Response, error) {
	//
	path := "/dna/system/api/v1/users/external-servers/aaa-attribute"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesDeleteAAAAttributeAPI{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAAAAttributeAPI()
		}
		return nil, response, fmt.Errorf("error with operation DeleteAAAAttributeApi")
	}

	result := response.Result().(*ResponseUserandRolesDeleteAAAAttributeAPI)
	return result, response, err

}
