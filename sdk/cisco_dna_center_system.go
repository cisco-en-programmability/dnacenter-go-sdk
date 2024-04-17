package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type CiscoDnaCenterSystemService service

type ResponseCiscoDnaCenterSystemAuthorizeDevice struct {
	JSONResponse      *ResponseCiscoDnaCenterSystemAuthorizeDeviceJSONResponse `json:"jsonResponse,omitempty"`      //
	Message           string                                                   `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                 `json:"statusCode,omitempty"`        // Status Code
	JSONArrayResponse []string                                                 `json:"jsonArrayResponse,omitempty"` // Json Array Response
}
type ResponseCiscoDnaCenterSystemAuthorizeDeviceJSONResponse struct {
	Empty *bool `json:"empty,omitempty"` // Empty
}
type RequestCiscoDnaCenterSystemAuthorizeDevice struct {
	DeviceIDList []string `json:"deviceIdList,omitempty"` // Device Id List
}

//AuthorizeDevice Authorize Device - 2897-4ae4-4ae9-a1dc
/* Authorizes one of more devices. A device can only be authorized if Authorization is set in Device Settings.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!authorize-device
*/
func (s *CiscoDnaCenterSystemService) AuthorizeDevice(requestCiscoDnaCenterSystemAuthorizeDevice *RequestCiscoDnaCenterSystemAuthorizeDevice) (*ResponseCiscoDnaCenterSystemAuthorizeDevice, *resty.Response, error) {
	path := "/api/v1/onboarding/pnp-device/authorize"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestCiscoDnaCenterSystemAuthorizeDevice).
		SetResult(&ResponseCiscoDnaCenterSystemAuthorizeDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AuthorizeDevice(requestCiscoDnaCenterSystemAuthorizeDevice)
		}

		return nil, response, fmt.Errorf("error with operation AuthorizeDevice")
	}

	result := response.Result().(*ResponseCiscoDnaCenterSystemAuthorizeDevice)
	return result, response, err

}
