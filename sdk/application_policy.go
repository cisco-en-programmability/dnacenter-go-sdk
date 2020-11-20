package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ApplicationPolicyService is the service to communicate with the ApplicationPolicy API endpoint
type ApplicationPolicyService service

// CreateApplicationRequest is the createApplicationRequest definition
type CreateApplicationRequest struct {
	ApplicationSet      CreateApplicationRequestApplicationSet        `json:"applicationSet,omitempty"`      //
	Name                string                                        `json:"name,omitempty"`                //
	NetworkApplications []CreateApplicationRequestNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     []CreateApplicationRequestNetworkIDentity     `json:"networkIdentity,omitempty"`     //
}

// CreateApplicationRequestApplicationSet is the createApplicationRequestApplicationSet definition
type CreateApplicationRequestApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` //
}

// CreateApplicationRequestNetworkApplications is the createApplicationRequestNetworkApplications definition
type CreateApplicationRequestNetworkApplications struct {
	AppProtocol        string `json:"appProtocol,omitempty"`        //
	ApplicationSubType string `json:"applicationSubType,omitempty"` //
	ApplicationType    string `json:"applicationType,omitempty"`    //
	CategoryID         string `json:"categoryId,omitempty"`         //
	DisplayName        string `json:"displayName,omitempty"`        //
	Dscp               string `json:"dscp,omitempty"`               //
	EngineID           string `json:"engineId,omitempty"`           //
	HelpString         string `json:"helpString,omitempty"`         //
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     //
	LongDescription    string `json:"longDescription,omitempty"`    //
	Name               string `json:"name,omitempty"`               //
	Popularity         int    `json:"popularity,omitempty"`         //
	Rank               int    `json:"rank,omitempty"`               //
	ServerName         string `json:"serverName,omitempty"`         //
	TrafficClass       string `json:"trafficClass,omitempty"`       //
	URL                string `json:"url,omitempty"`                //
}

// CreateApplicationRequestNetworkIDentity is the createApplicationRequestNetworkIDentity definition
type CreateApplicationRequestNetworkIDentity struct {
	DisplayName string `json:"displayName,omitempty"` //
	LowerPort   int    `json:"lowerPort,omitempty"`   //
	Ports       string `json:"ports,omitempty"`       //
	Protocol    string `json:"protocol,omitempty"`    //
	UpperPort   int    `json:"upperPort,omitempty"`   //
}

// CreateApplicationSetRequest is the createApplicationSetRequest definition
type CreateApplicationSetRequest struct {
	Name string `json:"name,omitempty"` //
}

// EditApplicationRequest is the editApplicationRequest definition
type EditApplicationRequest struct {
	ApplicationSet      EditApplicationRequestApplicationSet        `json:"applicationSet,omitempty"`      //
	ID                  string                                      `json:"id,omitempty"`                  //
	Name                string                                      `json:"name,omitempty"`                //
	NetworkApplications []EditApplicationRequestNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     []EditApplicationRequestNetworkIDentity     `json:"networkIdentity,omitempty"`     //
}

// EditApplicationRequestApplicationSet is the editApplicationRequestApplicationSet definition
type EditApplicationRequestApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` //
}

// EditApplicationRequestNetworkApplications is the editApplicationRequestNetworkApplications definition
type EditApplicationRequestNetworkApplications struct {
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
	Popularity         int    `json:"popularity,omitempty"`         //
	Rank               int    `json:"rank,omitempty"`               //
	ServerName         string `json:"serverName,omitempty"`         //
	TrafficClass       string `json:"trafficClass,omitempty"`       //
	URL                string `json:"url,omitempty"`                //
}

// EditApplicationRequestNetworkIDentity is the editApplicationRequestNetworkIDentity definition
type EditApplicationRequestNetworkIDentity struct {
	DisplayName string `json:"displayName,omitempty"` //
	ID          string `json:"id,omitempty"`          //
	LowerPort   int    `json:"lowerPort,omitempty"`   //
	Ports       string `json:"ports,omitempty"`       //
	Protocol    string `json:"protocol,omitempty"`    //
	UpperPort   int    `json:"upperPort,omitempty"`   //
}

// CreateApplicationResponse is the createApplicationResponse definition
type CreateApplicationResponse struct {
	Response CreateApplicationResponseResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  //
}

// CreateApplicationResponseResponse is the createApplicationResponseResponse definition
type CreateApplicationResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// CreateApplicationSetResponse is the createApplicationSetResponse definition
type CreateApplicationSetResponse struct {
	Response CreateApplicationSetResponseResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

// CreateApplicationSetResponseResponse is the createApplicationSetResponseResponse definition
type CreateApplicationSetResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeleteApplicationResponse is the deleteApplicationResponse definition
type DeleteApplicationResponse struct {
	Response DeleteApplicationResponseResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  //
}

// DeleteApplicationResponseResponse is the deleteApplicationResponseResponse definition
type DeleteApplicationResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeleteApplicationSetResponse is the deleteApplicationSetResponse definition
type DeleteApplicationSetResponse struct {
	Response DeleteApplicationSetResponseResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

// DeleteApplicationSetResponseResponse is the deleteApplicationSetResponseResponse definition
type DeleteApplicationSetResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// EditApplicationResponse is the editApplicationResponse definition
type EditApplicationResponse struct {
	Response EditApplicationResponseResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}

// EditApplicationResponseResponse is the editApplicationResponseResponse definition
type EditApplicationResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// GetApplicationSetsCountResponse is the getApplicationSetsCountResponse definition
type GetApplicationSetsCountResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetApplicationSetsResponse is the getApplicationSetsResponse definition
type GetApplicationSetsResponse struct {
	Response []GetApplicationSetsResponseResponse `json:"response,omitempty"` //
}

// GetApplicationSetsResponseResponse is the getApplicationSetsResponseResponse definition
type GetApplicationSetsResponseResponse struct {
	ID             string                                           `json:"id,omitempty"`             //
	IDentitySource GetApplicationSetsResponseResponseIDentitySource `json:"identitySource,omitempty"` //
	Name           string                                           `json:"name,omitempty"`           //
}

// GetApplicationSetsResponseResponseIDentitySource is the getApplicationSetsResponseResponseIDentitySource definition
type GetApplicationSetsResponseResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   //
	Type string `json:"type,omitempty"` //
}

// GetApplicationsCountResponse is the getApplicationsCountResponse definition
type GetApplicationsCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetApplicationsResponse is the getApplicationsResponse definition
type GetApplicationsResponse struct {
	Response []GetApplicationsResponseResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  //
}

// GetApplicationsResponseResponse is the getApplicationsResponseResponse definition
type GetApplicationsResponseResponse struct {
	ApplicationSet      GetApplicationsResponseApplicationSet        `json:"applicationSet,omitempty"`      //
	ID                  string                                       `json:"id,omitempty"`                  //
	Name                string                                       `json:"name,omitempty"`                //
	NetworkApplications []GetApplicationsResponseNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     []GetApplicationsResponseNetworkIDentity     `json:"networkIdentity,omitempty"`     //
}

// GetApplicationsResponseApplicationSet is the getApplicationsResponseApplicationSet definition
type GetApplicationsResponseApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` //
}

// GetApplicationsResponseNetworkApplications is the getApplicationsResponseNetworkApplications definition
type GetApplicationsResponseNetworkApplications struct {
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
	Popularity         int    `json:"popularity,omitempty"`         //
	Rank               int    `json:"rank,omitempty"`               //
	ServerName         string `json:"serverName,omitempty"`         //
	TrafficClass       string `json:"trafficClass,omitempty"`       //
	URL                string `json:"url,omitempty"`                //
}

// GetApplicationsResponseNetworkIDentity is the getApplicationsResponseNetworkIDentity definition
type GetApplicationsResponseNetworkIDentity struct {
	DisplayName string `json:"displayName,omitempty"` //
	ID          string `json:"id,omitempty"`          //
	LowerPort   int    `json:"lowerPort,omitempty"`   //
	Ports       string `json:"ports,omitempty"`       //
	Protocol    string `json:"protocol,omitempty"`    //
	UpperPort   int    `json:"upperPort,omitempty"`   //
}

// CreateApplication createApplication
/* Create new Custom application
 */
func (s *ApplicationPolicyService) CreateApplication(createApplicationRequest *[]CreateApplicationRequest) (*CreateApplicationResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	response, err := RestyClient.R().
		SetBody(createApplicationRequest).
		SetResult(&CreateApplicationResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createApplication")
	}

	result := response.Result().(*CreateApplicationResponse)
	return result, response, err
}

// CreateApplicationSet createApplicationSet
/* Create new custom application-set/s
 */
func (s *ApplicationPolicyService) CreateApplicationSet(createApplicationSetRequest *[]CreateApplicationSetRequest) (*CreateApplicationSetResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/application-policy-application-set"

	response, err := RestyClient.R().
		SetBody(createApplicationSetRequest).
		SetResult(&CreateApplicationSetResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createApplicationSet")
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
@param id Application's Id
*/
func (s *ApplicationPolicyService) DeleteApplication(deleteApplicationQueryParams *DeleteApplicationQueryParams) (*DeleteApplicationResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(deleteApplicationQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeleteApplicationResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteApplication")
	}

	result := response.Result().(*DeleteApplicationResponse)
	return result, response, err
}

// DeleteApplicationSetQueryParams defines the query parameters for this request
type DeleteApplicationSetQueryParams struct {
	ID string `url:"id,omitempty"` //
}

// DeleteApplicationSet deleteApplicationSet
/* Delete existing application-set by it's id
@param id
*/
func (s *ApplicationPolicyService) DeleteApplicationSet(deleteApplicationSetQueryParams *DeleteApplicationSetQueryParams) (*DeleteApplicationSetResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(deleteApplicationSetQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeleteApplicationSetResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteApplicationSet")
	}

	result := response.Result().(*DeleteApplicationSetResponse)
	return result, response, err
}

// EditApplication editApplication
/* Edit the attributes of an existing application
 */
func (s *ApplicationPolicyService) EditApplication(editApplicationRequest *[]EditApplicationRequest) (*EditApplicationResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/applications"

	response, err := RestyClient.R().
		SetBody(editApplicationRequest).
		SetResult(&EditApplicationResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation editApplication")
	}

	result := response.Result().(*EditApplicationResponse)
	return result, response, err
}

// GetApplicationSetsQueryParams defines the query parameters for this request
type GetApplicationSetsQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //
	Limit  float64 `url:"limit,omitempty"`  //
	Name   string  `url:"name,omitempty"`   //
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getApplicationSets")
	}

	result := response.Result().(*GetApplicationSetsResponse)
	return result, response, err
}

// GetApplicationSetsCount getApplicationSetsCount
/* Get the number of existing application-sets
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getApplicationSetsCount")
	}

	result := response.Result().(*GetApplicationSetsCountResponse)
	return result, response, err
}

// GetApplicationsQueryParams defines the query parameters for this request
type GetApplicationsQueryParams struct {
	Offset float64 `url:"offset,omitempty"` // The offset of the first application to be returned
	Limit  float64 `url:"limit,omitempty"`  // The maximum number of applications to be returned
	Name   string  `url:"name,omitempty"`   // Application's name
}

// GetApplications getApplications
/* Get applications by offset/limit or by name
@param offset The offset of the first application to be returned
@param limit The maximum number of applications to be returned
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getApplications")
	}

	result := response.Result().(*GetApplicationsResponse)
	return result, response, err
}

// GetApplicationsCount getApplicationsCount
/* Get the number of all existing applications
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getApplicationsCount")
	}

	result := response.Result().(*GetApplicationsCountResponse)
	return result, response, err
}
