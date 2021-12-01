package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplicationPolicyService service

type GetApplicationSetsQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //
	Limit  float64 `url:"limit,omitempty"`  //
	Name   string  `url:"name,omitempty"`   //
}
type DeleteApplicationSetQueryParams struct {
	ID string `url:"id,omitempty"` //
}
type DeleteApplicationQueryParams struct {
	ID string `url:"id,omitempty"` //Application's Id
}
type GetApplicationsQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The offset of the first application to be returned
	Limit  float64 `url:"limit,omitempty"`  //The maximum number of applications to be returned
	Name   string  `url:"name,omitempty"`   //Application's name
}

type ResponseApplicationPolicyGetApplicationSets struct {
	Response *[]ResponseApplicationPolicyGetApplicationSetsResponse `json:"response,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationSetsResponse struct {
	ID             string                                                             `json:"id,omitempty"`             // Id
	IDentitySource *ResponseApplicationPolicyGetApplicationSetsResponseIDentitySource `json:"identitySource,omitempty"` //
	Name           string                                                             `json:"name,omitempty"`           // Name
}
type ResponseApplicationPolicyGetApplicationSetsResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type
}
type ResponseApplicationPolicyDeleteApplicationSet struct {
	Response *ResponseApplicationPolicyDeleteApplicationSetResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyCreateApplicationSet struct {
	Response *ResponseApplicationPolicyCreateApplicationSetResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationSetResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyGetApplicationSetsCount struct {
	Response string `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplication struct {
	Response *ResponseApplicationPolicyCreateApplicationResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyEditApplication struct {
	Response *ResponseApplicationPolicyEditApplicationResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyEditApplicationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyDeleteApplication struct {
	Response *ResponseApplicationPolicyDeleteApplicationResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyGetApplications []ResponseItemApplicationPolicyGetApplications // Array of ResponseApplicationPolicyGetApplications
type ResponseItemApplicationPolicyGetApplications struct {
	ID                  string                                                             `json:"id,omitempty"`                  // Id
	Name                string                                                             `json:"name,omitempty"`                // Name
	NetworkApplications *[]ResponseItemApplicationPolicyGetApplicationsNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]ResponseItemApplicationPolicyGetApplicationsNetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *ResponseItemApplicationPolicyGetApplicationsApplicationSet        `json:"applicationSet,omitempty"`      //
}
type ResponseItemApplicationPolicyGetApplicationsNetworkApplications struct {
	ID                 string `json:"id,omitempty"`                 // Id
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type ResponseItemApplicationPolicyGetApplicationsNetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // Id
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type ResponseItemApplicationPolicyGetApplicationsApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseApplicationPolicyGetApplicationsCount struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type RequestApplicationPolicyCreateApplicationSet []RequestItemApplicationPolicyCreateApplicationSet // Array of RequestApplicationPolicyCreateApplicationSet
type RequestItemApplicationPolicyCreateApplicationSet struct {
	Name string `json:"name,omitempty"` // Name
}
type RequestApplicationPolicyCreateApplication []RequestItemApplicationPolicyCreateApplication // Array of RequestApplicationPolicyCreateApplication
type RequestItemApplicationPolicyCreateApplication struct {
	Name                string                                                              `json:"name,omitempty"`                // Name
	NetworkApplications *[]RequestItemApplicationPolicyCreateApplicationNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]RequestItemApplicationPolicyCreateApplicationNetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *RequestItemApplicationPolicyCreateApplicationApplicationSet        `json:"applicationSet,omitempty"`      //
}
type RequestItemApplicationPolicyCreateApplicationNetworkApplications struct {
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type RequestItemApplicationPolicyCreateApplicationNetworkIDentity struct {
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type RequestItemApplicationPolicyCreateApplicationApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type RequestApplicationPolicyEditApplication []RequestItemApplicationPolicyEditApplication // Array of RequestApplicationPolicyEditApplication
type RequestItemApplicationPolicyEditApplication struct {
	ID                  string                                                            `json:"id,omitempty"`                  // Id
	Name                string                                                            `json:"name,omitempty"`                // Name
	NetworkApplications *[]RequestItemApplicationPolicyEditApplicationNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]RequestItemApplicationPolicyEditApplicationNetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *RequestItemApplicationPolicyEditApplicationApplicationSet        `json:"applicationSet,omitempty"`      //
}
type RequestItemApplicationPolicyEditApplicationNetworkApplications struct {
	ID                 string `json:"id,omitempty"`                 // Id
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type RequestItemApplicationPolicyEditApplicationNetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // Id
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type RequestItemApplicationPolicyEditApplicationApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}

//GetApplicationSets Get Application Sets - cb86-8b21-4289-8159
/* Get appllication-sets by offset/limit or by name


@param GetApplicationSetsQueryParams Filtering parameter
*/
func (s *ApplicationPolicyService) GetApplicationSets(GetApplicationSetsQueryParams *GetApplicationSetsQueryParams) (*ResponseApplicationPolicyGetApplicationSets, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(GetApplicationSetsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetApplicationSets")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSets)
	return result, response, err

}

//GetApplicationSetsCount Get Application Sets Count - cfa0-49a6-44bb-8a07
/* Get the number of existing application-sets


 */
func (s *ApplicationPolicyService) GetApplicationSetsCount() (*ResponseApplicationPolicyGetApplicationSetsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationSetsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsCount)
	return result, response, err

}

//GetApplications Get Applications - 8893-b834-445b-b29c
/* Get applications by offset/limit or by name


@param GetApplicationsQueryParams Filtering parameter
*/
func (s *ApplicationPolicyService) GetApplications(GetApplicationsQueryParams *GetApplicationsQueryParams) (*ResponseApplicationPolicyGetApplications, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(GetApplicationsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplications{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetApplications")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplications)
	return result, response, err

}

//GetApplicationsCount Get Applications Count - 039d-e8b1-47a9-8690
/* Get the number of all existing applications


 */
func (s *ApplicationPolicyService) GetApplicationsCount() (*ResponseApplicationPolicyGetApplicationsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetApplicationsCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsCount)
	return result, response, err

}

//CreateApplicationSet Create Application Set - 3e94-cb1b-485b-8b0e
/* Create new custom application-set/s


 */
func (s *ApplicationPolicyService) CreateApplicationSet(requestApplicationPolicyCreateApplicationSet *RequestApplicationPolicyCreateApplicationSet) (*ResponseApplicationPolicyCreateApplicationSet, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationSet).
		SetResult(&ResponseApplicationPolicyCreateApplicationSet{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateApplicationSet")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationSet)
	return result, response, err

}

//CreateApplication Create Application - fb9b-f80f-491a-9851
/* Create new Custom application


 */
func (s *ApplicationPolicyService) CreateApplication(requestApplicationPolicyCreateApplication *RequestApplicationPolicyCreateApplication) (*ResponseApplicationPolicyCreateApplication, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplication).
		SetResult(&ResponseApplicationPolicyCreateApplication{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplication)
	return result, response, err

}

//EditApplication Edit Application - 3986-6887-4439-a41d
/* Edit the attributes of an existing application


 */
func (s *ApplicationPolicyService) EditApplication(requestApplicationPolicyEditApplication *RequestApplicationPolicyEditApplication) (*ResponseApplicationPolicyEditApplication, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyEditApplication).
		SetResult(&ResponseApplicationPolicyEditApplication{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation EditApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyEditApplication)
	return result, response, err

}

//DeleteApplicationSet Delete Application Set - 70b6-f8e1-40b8-b784
/* Delete existing application-set by it's id


@param DeleteApplicationSetQueryParams Filtering parameter
*/
func (s *ApplicationPolicyService) DeleteApplicationSet(DeleteApplicationSetQueryParams *DeleteApplicationSetQueryParams) (*ResponseApplicationPolicyDeleteApplicationSet, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(DeleteApplicationSetQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyDeleteApplicationSet{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteApplicationSet")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationSet)
	return result, response, err

}

//DeleteApplication Delete Application - d49a-f9b8-4c6a-a8ea
/* Delete existing application by its id


@param DeleteApplicationQueryParams Filtering parameter
*/
func (s *ApplicationPolicyService) DeleteApplication(DeleteApplicationQueryParams *DeleteApplicationQueryParams) (*ResponseApplicationPolicyDeleteApplication, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(DeleteApplicationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyDeleteApplication{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplication)
	return result, response, err

}
