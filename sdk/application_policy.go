package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ApplicationPolicyService is the service to communicate with the ApplicationPolicy API endpoint
type ApplicationPolicyService service

// ApplicationSet is the ApplicationSet definition
type ApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` //
}

// CreateApplicationRequest is the CreateApplicationRequest definition
type CreateApplicationRequest struct {
	ApplicationSet      ApplicationSet        `json:"applicationSet,omitempty"`      //
	Name                string                `json:"name,omitempty"`                //
	NetworkApplications []NetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     []NetworkIdentity     `json:"networkIdentity,omitempty"`     //
}

// CreateApplicationSetRequest is the CreateApplicationSetRequest definition
type CreateApplicationSetRequest struct {
	Name string `json:"name,omitempty"` //
}

// EditApplicationRequest is the EditApplicationRequest definition
type EditApplicationRequest struct {
	ApplicationSet      ApplicationSet        `json:"applicationSet,omitempty"`      //
	ID                  string                `json:"id,omitempty"`                  //
	Name                string                `json:"name,omitempty"`                //
	NetworkApplications []NetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     []NetworkIdentity     `json:"networkIdentity,omitempty"`     //
}

// NetworkApplications is the NetworkApplications definition
type NetworkApplications struct {
	AppProtocol        string `json:"appProtocol,omitempty"`        //
	ApplicationSubType string `json:"applicationSubType,omitempty"` //
	ApplicationType    string `json:"applicationType,omitempty"`    //
	CategoryID         string `json:"categoryId,omitempty"`         //
	DisplayName        string `json:"displayName,omitempty"`        //
	Dscp               string `json:"dscp,omitempty"`               //
	EngineID           string `json:"engineId,omitempty"`           //
	HelpString         string `json:"helpString,omitempty"`         //
	ID                 string `json:"id,omitempty"`                 //
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     //
	LongDescription    string `json:"longDescription,omitempty"`    //
	Name               string `json:"name,omitempty"`               //
	Popularity         string `json:"popularity,omitempty"`         //
	Rank               string `json:"rank,omitempty"`               //
	ServerName         string `json:"serverName,omitempty"`         //
	TrafficClass       string `json:"trafficClass,omitempty"`       //
	URL                string `json:"url,omitempty"`                //
}

// NetworkIdentity is the NetworkIdentity definition
type NetworkIdentity struct {
	DisplayName string `json:"displayName,omitempty"` //
	ID          string `json:"id,omitempty"`          //
	LowerPort   string `json:"lowerPort,omitempty"`   //
	Ports       string `json:"ports,omitempty"`       //
	Protocol    string `json:"protocol,omitempty"`    //
	UpperPort   string `json:"upperPort,omitempty"`   //
}

// CreateApplicationResponse is the CreateApplicationResponse definition
type CreateApplicationResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// CreateApplicationSetResponse is the CreateApplicationSetResponse definition
type CreateApplicationSetResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DeleteApplicationResponse is the DeleteApplicationResponse definition
type DeleteApplicationResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DeleteApplicationSetResponse is the DeleteApplicationSetResponse definition
type DeleteApplicationSetResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// EditApplicationResponse is the EditApplicationResponse definition
type EditApplicationResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetApplicationSetsCountResponse is the GetApplicationSetsCountResponse definition
type GetApplicationSetsCountResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetApplicationSetsResponse is the GetApplicationSetsResponse definition
type GetApplicationSetsResponse struct {
	Response []Response `json:"response,omitempty"` //
}

// GetApplicationsCountResponse is the GetApplicationsCountResponse definition
type GetApplicationsCountResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetApplicationsResponse is the GetApplicationsResponse definition
type GetApplicationsResponse struct {
	ApplicationSet      ApplicationSet        `json:"applicationSet,omitempty"`      //
	ID                  string                `json:"id,omitempty"`                  //
	Name                string                `json:"name,omitempty"`                //
	NetworkApplications []NetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     []NetworkIdentity     `json:"networkIdentity,omitempty"`     //
}

// IDentitySource is the IdentitySource definition
type IDentitySource struct {
	ID   string `json:"id,omitempty"`   //
	Type string `json:"type,omitempty"` //
}

// Response is the Response definition
type Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// CreateApplication createApplication
/* Create new Custom application
 */
func (s *ApplicationPolicyService) CreateApplication(createApplicationRequest *CreateApplicationRequest) (*CreateApplicationResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	response, err := RestyClient.R().
		SetBody(createApplicationRequest).
		SetResult(&CreateApplicationResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateApplicationResponse)
	return result, response, err

}

// CreateApplicationSet createApplicationSet
/* Create new custom application-set/s
 */
func (s *ApplicationPolicyService) CreateApplicationSet(createApplicationSetRequest *CreateApplicationSetRequest) (*CreateApplicationSetResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/application-policy-application-set"

	response, err := RestyClient.R().
		SetBody(createApplicationSetRequest).
		SetResult(&CreateApplicationSetResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateApplicationSetResponse)
	return result, response, err

}

// DeleteApplicationQueryParams defines the query parameters for this request
type DeleteApplicationQueryParams struct {
	ID string `url:"id,omitempty"` // Application's Id
}

// DeleteApplication deleteApplication
/* Delete existing application by its id
@param id Application's ID
*/
func (s *ApplicationPolicyService) DeleteApplication(deleteApplicationQueryParams *DeleteApplicationQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(deleteApplicationQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteApplicationSetQueryParams defines the query parameters for this request
type DeleteApplicationSetQueryParams struct {
	ID string `url:"id,omitempty"` //
}

// DeleteApplicationSet deleteApplicationSet
/* Delete existing application-set by it's id
@param id
*/
func (s *ApplicationPolicyService) DeleteApplicationSet(deleteApplicationSetQueryParams *DeleteApplicationSetQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(deleteApplicationSetQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// EditApplication editApplication
/* Edit the attributes of an existing application
 */
func (s *ApplicationPolicyService) EditApplication(editApplicationRequest *EditApplicationRequest) (*EditApplicationResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	response, err := RestyClient.R().
		SetBody(editApplicationRequest).
		SetResult(&EditApplicationResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*EditApplicationResponse)
	return result, response, err

}

// GetApplicationSetsQueryParams defines the query parameters for this request
type GetApplicationSetsQueryParams struct {
	Offset int    `url:"offset,omitempty"` //
	Limit  int    `url:"limit,omitempty"`  //
	Name   string `url:"name,omitempty"`   //
}

// GetApplicationSets getApplicationSets
/* Get appllication-sets by offset/limit or by name
@param offset
@param limit
@param name
*/
func (s *ApplicationPolicyService) GetApplicationSets(getApplicationSetsQueryParams *GetApplicationSetsQueryParams) (*GetApplicationSetsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(getApplicationSetsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetApplicationSetsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetApplicationSetsResponse)
	return result, response, err

}

// GetApplicationSetsCount getApplicationSetsCount
/* Get the int of existing application-sets
 */
func (s *ApplicationPolicyService) GetApplicationSetsCount() (*GetApplicationSetsCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/application-policy-application-set-count"

	response, err := RestyClient.R().
		SetResult(&GetApplicationSetsCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetApplicationSetsCountResponse)
	return result, response, err

}

// GetApplicationsQueryParams defines the query parameters for this request
type GetApplicationsQueryParams struct {
	Offset int    `url:"offset,omitempty"` // The offset of the first application to be returned
	Limit  int    `url:"limit,omitempty"`  // The maximum int of applications to be returned
	Name   string `url:"name,omitempty"`   // Application's name
}

// GetApplications getApplications
/* Get applications by offset/limit or by name
@param offset The offset of the first application to be returned
@param limit The maximum int of applications to be returned
@param name Application's name
*/
func (s *ApplicationPolicyService) GetApplications(getApplicationsQueryParams *GetApplicationsQueryParams) (*GetApplicationsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(getApplicationsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetApplicationsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetApplicationsResponse)
	return result, response, err

}

// GetApplicationsCount getApplicationsCount
/* Get the int of all existing applications
 */
func (s *ApplicationPolicyService) GetApplicationsCount() (*GetApplicationsCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications-count"

	response, err := RestyClient.R().
		SetResult(&GetApplicationsCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetApplicationsCountResponse)
	return result, response, err

}
