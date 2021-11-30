package dnac

import (
	"github.com/go-resty/resty/v2"
)

// AuthenticationService is the service to communicate with the Authentication API endpoint
type AuthenticationService service

// AuthenticationAPIResponse is the AuthenticationAPIResponse definition
type AuthenticationAPIResponse struct {
	Token string `json:"Token,omitempty"` //
}

// AuthenticationAPI authenticationAPI
/* API to obtain an access token. The token obtained using this API is required to be set as value to the X-Auth-Token HTTP Header for all API calls to Cisco DNA Center.
@param Content-Type Request body content type
@param Authorization String composed of “Basic”, followed by a space, followed by the Base64 encoding of “username:password”, NOT including the quotes. For example “Basic YWRtaW46TWFnbGV2MTIz”, where YWRtaW46TWFnbGV2MTIz is the Base 64 encoding.
*/
func (s *AuthenticationService) AuthenticationAPI(username string, password string) (*AuthenticationAPIResponse, *resty.Response, error) {

	path := "/dna/system/api/v1/auth/token"

	response, err := s.client.R().
		SetBasicAuth(username, password).
		SetResult(&AuthenticationAPIResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AuthenticationAPIResponse)
	return result, response, err

}
