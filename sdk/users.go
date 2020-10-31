package dnac

import (
	"github.com/go-resty/resty/v2"
)

// UsersService is the service to communicate with the Users API endpoint
type UsersService service

// GetUserEnrichmentDetailsResponse is the GetUserEnrichmentDetailsResponse definition
type GetUserEnrichmentDetailsResponse struct {
	ConnectedDevice []ConnectedDevice `json:"connectedDevice,omitempty"` //
	UserDetails     UserDetails       `json:"userDetails,omitempty"`     //
}

// NeighborTopology is the NeighborTopology definition
// type UsersNeighborTopology struct {
// 	Detail    string `json:"detail,omitempty"`    //
// 	ErrorCode int    `json:"errorCode,omitempty"` //
// 	Message   string `json:"message,omitempty"`   //
// }

// UsersOnboarding is the Onboarding definition
type UsersOnboarding struct {
	AaaServerIP          string `json:"aaaServerIp,omitempty"`          //
	AverageAssocDuration string `json:"averageAssocDuration,omitempty"` //
	AverageAuthDuration  string `json:"averageAuthDuration,omitempty"`  //
	AverageDhcpDuration  string `json:"averageDhcpDuration,omitempty"`  //
	AverageRunDuration   string `json:"averageRunDuration,omitempty"`   //
	DhcpServerIP         string `json:"dhcpServerIp,omitempty"`         //
	MaxAssocDuration     string `json:"maxAssocDuration,omitempty"`     //
	MaxAuthDuration      string `json:"maxAuthDuration,omitempty"`      //
	MaxDhcpDuration      string `json:"maxDhcpDuration,omitempty"`      //
	MaxRunDuration       string `json:"maxRunDuration,omitempty"`       //
}

// GetUserEnrichmentDetails getUserEnrichmentDetails
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user and devices that the user is connected to
@param entity_type User enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
@param entity_value Contains the actual value for the entity type that has been defined
*/
func (s *UsersService) GetUserEnrichmentDetails() (*GetUserEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/user-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetUserEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetUserEnrichmentDetailsResponse)
	return result, response, err

}
