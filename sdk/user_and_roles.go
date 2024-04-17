package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type UserandRolesService service

type GetRolesApIHeaderParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //Expects type string. The source that invoke this API
}
type GetUsersApIQueryParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //The source that invokes this API
}
type GetExternalAuthenticationServersApIQueryParams struct {
	InvokeSource string `url:"invokeSource,omitempty"` //The source that invokes this API
}

type ResponseUserandRolesGetPermissionsApI struct {
	Response *ResponseUserandRolesGetPermissionsAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetPermissionsAPIResponse struct {
	ResourceTypes *[]ResponseUserandRolesGetPermissionsAPIResponseResourceTypes `json:"resource-types,omitempty"` //
}
type ResponseUserandRolesGetPermissionsAPIResponseResourceTypes struct {
	Type              string `json:"type,omitempty"`              // Type
	DisplayName       string `json:"displayName,omitempty"`       // Display Name
	Description       string `json:"description,omitempty"`       // Description
	DefaultPermission string `json:"defaultPermission,omitempty"` // Default permission
}
type ResponseUserandRolesGetRolesApI struct {
	Response *ResponseUserandRolesGetRolesAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetRolesAPIResponse struct {
	Roles *[]ResponseUserandRolesGetRolesAPIResponseRoles `json:"roles,omitempty"` //
}
type ResponseUserandRolesGetRolesAPIResponseRoles struct {
	ResourceTypes *[]ResponseUserandRolesGetRolesAPIResponseRolesResourceTypes `json:"resourceTypes,omitempty"` //
	Meta          *ResponseUserandRolesGetRolesAPIResponseRolesMeta            `json:"meta,omitempty"`          //
	RoleID        string                                                       `json:"roleId,omitempty"`        // Role Id
	Name          string                                                       `json:"name,omitempty"`          // Role name
	Description   string                                                       `json:"description,omitempty"`   // Description
	Type          string                                                       `json:"type,omitempty"`          // Role type, possible values are: "DEFAULT", "SYSTEM", "CUSTOM"
}
type ResponseUserandRolesGetRolesAPIResponseRolesResourceTypes struct {
	Operations []string `json:"operations,omitempty"` // Operations
	Type       string   `json:"type,omitempty"`       // Type
}
type ResponseUserandRolesGetRolesAPIResponseRolesMeta struct {
	CreatedBy    string `json:"createdBy,omitempty"`    // The user that creates the resource type
	Created      string `json:"created,omitempty"`      // The timestamp that the resource type was created
	LastModified string `json:"lastModified,omitempty"` // The latestest timestamp that the resource type was updated
}
type ResponseUserandRolesGetUsersApI struct {
	Response *ResponseUserandRolesGetUsersAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetUsersAPIResponse struct {
	Users *[]ResponseUserandRolesGetUsersAPIResponseUsers `json:"users,omitempty"` //
}
type ResponseUserandRolesGetUsersAPIResponseUsers struct {
	FirstName            string   `json:"firstName,omitempty"`            // First Name
	LastName             string   `json:"lastName,omitempty"`             // Last Name
	AuthSource           string   `json:"authSource,omitempty"`           // Authentiction source, internal or external
	PassphraseUpdateTime string   `json:"passphraseUpdateTime,omitempty"` // Passphrase Update Time
	RoleList             []string `json:"roleList,omitempty"`             // A list of role ids
	UserID               string   `json:"userId,omitempty"`               // User Id
	Email                string   `json:"email,omitempty"`                // Email
	Username             string   `json:"username,omitempty"`             // Username
}
type ResponseUserandRolesAddUserApI struct {
	Response *ResponseUserandRolesAddUserAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesAddUserAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
	UserID  string `json:"userId,omitempty"`  // User Id
}
type ResponseUserandRolesUpdateUserApI struct {
	Response *ResponseUserandRolesUpdateUserAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesUpdateUserAPIResponse struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseUserandRolesGetExternalAuthenticationServersApI struct {
	Response *ResponseUserandRolesGetExternalAuthenticationServersAPIResponse `json:"response,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIResponse struct {
	AAAServers *[]ResponseUserandRolesGetExternalAuthenticationServersAPIResponseAAAServers `json:"aaa-servers,omitempty"` //
}
type ResponseUserandRolesGetExternalAuthenticationServersAPIResponseAAAServers struct {
	AccountingPort     *int   `json:"accountingPort,omitempty"`     // RADIUS server accounting requests port
	Retries            *int   `json:"retries,omitempty"`            // Retries
	Protocol           string `json:"protocol,omitempty"`           // Protocol
	SocketTimeout      *int   `json:"socketTimeout,omitempty"`      // Timeout in seconds
	ServerIP           string `json:"serverIp,omitempty"`           // Server Ip
	SharedSecret       string `json:"sharedSecret,omitempty"`       // Shared Secret
	ServerID           string `json:"serverId,omitempty"`           // Server Id
	AuthenticationPort *int   `json:"authenticationPort,omitempty"` // RADIUS server authorization requests port
	AAAAttribute       string `json:"aaaAttribute,omitempty"`       // Aaa Attribute
	Role               string `json:"role,omitempty"`               // Role of AAA server, primary or secondary server
}
type RequestUserandRolesAddUserApI struct {
	FirstName string   `json:"firstName,omitempty"` // First Name
	LastName  string   `json:"lastName,omitempty"`  // Last Name
	Username  string   `json:"username,omitempty"`  // Username
	Password  string   `json:"password,omitempty"`  // Password
	Email     string   `json:"email,omitempty"`     // Email
	RoleList  []string `json:"roleList,omitempty"`  // Role id list
}
type RequestUserandRolesUpdateUserApI struct {
	FirstName string   `json:"firstName,omitempty"` // firstName should be set if the original value is not empty
	LastName  string   `json:"lastName,omitempty"`  // lastName should be set if the original value is not empty
	Email     string   `json:"email,omitempty"`     // email should be set if the original value is not empty
	Username  string   `json:"username,omitempty"`  // Username
	UserID    string   `json:"userId,omitempty"`    // User Id
	RoleList  []string `json:"roleList,omitempty"`  // Role id list
}

//GetPermissionsApI Get permissions API - 8a9c-6885-455b-a2db
/* Get permissions for a role from Cisco DNA Center System



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-permissions-api
*/
func (s *UserandRolesService) GetPermissionsAPI() (*ResponseUserandRolesGetPermissionsApI, *resty.Response, error) {
	path := "/dna/system/api/v1/role/permissions"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseUserandRolesGetPermissionsApI{}).
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

	result := response.Result().(*ResponseUserandRolesGetPermissionsApI)
	return result, response, err

}

//GetRolesAPI Get roles API - 7c86-da3f-4b08-8593
/* Get all roles for the Cisco DNA Center system


@param GetRolesAPIHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-roles-api
*/
func (s *UserandRolesService) GetRolesAPI(GetRolesAPIHeaderParams *GetRolesApIHeaderParams) (*ResponseUserandRolesGetRolesApI, *resty.Response, error) {
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
		SetResult(&ResponseUserandRolesGetRolesApI{}).
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

	result := response.Result().(*ResponseUserandRolesGetRolesApI)
	return result, response, err

}

//GetUsersAPI Get users API - 918c-89fa-4a98-a528
/* Get all users for the Cisco DNA Center system


@param GetUsersAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-users-api
*/
func (s *UserandRolesService) GetUsersAPI(GetUsersAPIQueryParams *GetUsersApIQueryParams) (*ResponseUserandRolesGetUsersApI, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	queryString, _ := query.Values(GetUsersAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserandRolesGetUsersApI{}).
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

	result := response.Result().(*ResponseUserandRolesGetUsersApI)
	return result, response, err

}

//GetExternalAuthenticationServersAPI Get external authentication servers API - 9dbd-0b01-4758-bde4
/* Get external users authentication servers


@param GetExternalAuthenticationServersAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-external-authentication-servers-api
*/
func (s *UserandRolesService) GetExternalAuthenticationServersAPI(GetExternalAuthenticationServersAPIQueryParams *GetExternalAuthenticationServersApIQueryParams) (*ResponseUserandRolesGetExternalAuthenticationServersApI, *resty.Response, error) {
	path := "/dna/system/api/v1/users/external-servers"

	queryString, _ := query.Values(GetExternalAuthenticationServersAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseUserandRolesGetExternalAuthenticationServersApI{}).
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

	result := response.Result().(*ResponseUserandRolesGetExternalAuthenticationServersApI)
	return result, response, err

}

//AddUserAPI Add user API - 6c9a-09c4-4a39-9e2b
/* Add a new user for Cisco DNA Center system



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-user-api
*/
func (s *UserandRolesService) AddUserAPI(requestUserandRolesAddUserAPI *RequestUserandRolesAddUserApI) (*ResponseUserandRolesAddUserApI, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesAddUserAPI).
		SetResult(&ResponseUserandRolesAddUserApI{}).
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

	result := response.Result().(*ResponseUserandRolesAddUserApI)
	return result, response, err

}

//UpdateUserAPI Update user API - f596-6adc-492b-a2ff
/* Update a user for Cisco DNA Center system


 */
func (s *UserandRolesService) UpdateUserAPI(requestUserandRolesUpdateUserAPI *RequestUserandRolesUpdateUserApI) (*ResponseUserandRolesUpdateUserApI, *resty.Response, error) {
	path := "/dna/system/api/v1/user"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestUserandRolesUpdateUserAPI).
		SetResult(&ResponseUserandRolesUpdateUserApI{}).
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

	result := response.Result().(*ResponseUserandRolesUpdateUserApI)
	return result, response, err

}
