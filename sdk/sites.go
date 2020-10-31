package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SitesService is the service to communicate with the Sites API endpoint
type SitesService service

// AssignDeviceToSiteRequest is the AssignDeviceToSiteRequest definition
type AssignDeviceToSiteRequest struct {
	Device []Device `json:"device,omitempty"` //
}

// CreateSiteRequest is the CreateSiteRequest definition
type CreateSiteRequest struct {
	Site Site   `json:"site,omitempty"` //
	Type string `json:"type,omitempty"` //
}

// SitesDevice is the Device definition
type SitesDevice struct {
	IP string `json:"ip,omitempty"` //
}

// Site is the Site definition
type Site struct {
	Area     Area     `json:"area,omitempty"`     //
	Building Building `json:"building,omitempty"` //
	Floor    Floor    `json:"floor,omitempty"`    //
}

// UpdateSiteRequest is the UpdateSiteRequest definition
type UpdateSiteRequest struct {
	Site Site   `json:"site,omitempty"` //
	Type string `json:"type,omitempty"` //
}

// ApplicationHealthStats is the ApplicationHealthStats definition
type ApplicationHealthStats struct {
	AppTotalCount              int                        `json:"appTotalCount,omitempty"`              //
	BusinessIrrelevantAppCount BusinessIrrelevantAppCount `json:"businessIrrelevantAppCount,omitempty"` //
	BusinessRelevantAppCount   BusinessRelevantAppCount   `json:"businessRelevantAppCount,omitempty"`   //
	DefaultHealthAppCount      DefaultHealthAppCount      `json:"defaultHealthAppCount,omitempty"`      //
}

// AssignDeviceToSiteResponse is the AssignDeviceToSiteResponse definition
type AssignDeviceToSiteResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// BusinessIrrelevantAppCount is the BusinessIrrelevantAppCount definition
type BusinessIrrelevantAppCount struct {
	Fair int `json:"fair,omitempty"` //
	Good int `json:"good,omitempty"` //
	Poor int `json:"poor,omitempty"` //
}

// BusinessRelevantAppCount is the BusinessRelevantAppCount definition
type BusinessRelevantAppCount struct {
	Fair int `json:"fair,omitempty"` //
	Good int `json:"good,omitempty"` //
	Poor int `json:"poor,omitempty"` //
}

// CreateSiteResponse is the CreateSiteResponse definition
type CreateSiteResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DefaultHealthAppCount is the DefaultHealthAppCount definition
type DefaultHealthAppCount struct {
	Fair int `json:"fair,omitempty"` //
	Good int `json:"good,omitempty"` //
	Poor int `json:"poor,omitempty"` //
}

// DeleteSiteResponse is the DeleteSiteResponse definition
type DeleteSiteResponse struct {
	Message string `json:"message,omitempty"` //
	Status  string `json:"status,omitempty"`  //
}

// SiteDevice is the Device definition
type SiteDevice struct {
	Response []string `json:"response,omitempty"` //
	SiteID   string   `json:"siteId,omitempty"`   //
	Version  string   `json:"version,omitempty"`  //
}

// GetMembershipResponse is the GetMembershipResponse definition
type GetMembershipResponse struct {
	Device []SiteDevice `json:"device,omitempty"` //
	Site   Site         `json:"site,omitempty"`   //
}

// GetSiteCountResponse is the GetSiteCountResponse definition
type GetSiteCountResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetSiteHealthResponse is the GetSiteHealthResponse definition
type GetSiteHealthResponse struct {
	Response []Response `json:"response,omitempty"` //
}

// GetSiteResponse is the GetSiteResponse definition
type GetSiteResponse struct {
	Response []SiteResponse `json:"response,omitempty"` //
}

// SiteResponse is the Response definition
type SiteResponse struct {
	Data             string   `json:"data,omitempty"`             //
	EndTime          string   `json:"endTime,omitempty"`          //
	ID               string   `json:"id,omitempty"`               //
	InstanceTenantID string   `json:"instanceTenantId,omitempty"` //
	IsError          string   `json:"isError,omitempty"`          //
	OperationIDList  []string `json:"operationIdList,omitempty"`  //
	Progress         string   `json:"progress,omitempty"`         //
	RootID           string   `json:"rootId,omitempty"`           //
	ServiceType      string   `json:"serviceType,omitempty"`      //
	StartTime        string   `json:"startTime,omitempty"`        //
	Version          string   `json:"version,omitempty"`          //
}

// SitesResponse is the Site definition
type SitesResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// UpdateSiteResponse is the UpdateSiteResponse definition
type UpdateSiteResponse struct {
	Response Response `json:"response,omitempty"` //
	Result   string   `json:"result,omitempty"`   //
	Status   string   `json:"status,omitempty"`   //
}

// AssignDeviceToSite assignDeviceToSite
/* Assigns list of devices to a site
@param __runsync Enable this parameter to execute the API and return a response synchronously
@param __persistbapioutput Persist bapi sync response
@param __runsynctimeout During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
@param siteID Site id to which site the device to assign
*/
func (s *SitesService) AssignDeviceToSite(siteID string, assignDeviceToSiteRequest *AssignDeviceToSiteRequest) (*AssignDeviceToSiteResponse, *resty.Response, error) {

	path := "/dna/system/api/v1/site/{siteID}/device"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(assignDeviceToSiteRequest).
		SetResult(&AssignDeviceToSiteResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AssignDeviceToSiteResponse)
	return result, response, err

}

// CreateSite createSite
/* Creates site with area/building/floor with specified hierarchy.
@param __runsync Enable this parameter to execute the API and return a response synchronously
@param __timeout During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
@param __persistbapioutput Persist bapi sync response
*/
func (s *SitesService) CreateSite(createSiteRequest *CreateSiteRequest) (*CreateSiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/site"

	response, err := RestyClient.R().
		SetBody(createSiteRequest).
		SetResult(&CreateSiteResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateSiteResponse)
	return result, response, err

}

// DeleteSite deleteSite
/* Delete site with area/building/floor by siteID.
@param siteID Site id to which site details to be deleted.
*/
func (s *SitesService) DeleteSite(siteID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/site/{siteID}"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetMembershipQueryParams defines the query parameters for this request
type GetMembershipQueryParams struct {
	Offset       string `url:"offset,omitempty"`       // offset/starting row
	Limit        string `url:"limit,omitempty"`        // Number of sites to be retrieved
	DeviceFamily string `url:"deviceFamily,omitempty"` // Device family name
	SerialNumber string `url:"serialNumber,omitempty"` // Device serial int
}

// GetMembership getMembership
/* Getting the site children details and device details.
@param siteID Site id to retrieve device associated with the site.
@param offset offset/starting row
@param limit Number of sites to be retrieved
@param deviceFamily Device family name
@param serialNumber Device serial int
*/
func (s *SitesService) GetMembership(siteID string, getMembershipQueryParams *GetMembershipQueryParams) (*GetMembershipResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/membership/{siteID}"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(getMembershipQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetMembershipResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetMembershipResponse)
	return result, response, err

}

// GetSiteQueryParams defines the query parameters for this request
type GetSiteQueryParams struct {
	Name   string `url:"name,omitempty"`   // siteNameHierarchy (ex: global/groupName)
	SiteID string `url:"siteId,omitempty"` // Site id to which site details to retrieve.
	Type   string `url:"type,omitempty"`   // type (ex: area, building, floor)
	Offset string `url:"offset,omitempty"` // offset/starting row
	Limit  string `url:"limit,omitempty"`  // Number of sites to be retrieved
}

// GetSite getSite
/* Get site with area/building/floor with specified hierarchy.
@param name siteNameHierarchy (ex: global/groupName)
@param siteID Site id to which site details to retrieve.
@param type type (ex: area, building, floor)
@param offset offset/starting row
@param limit Number of sites to be retrieved
*/
func (s *SitesService) GetSite(getSiteQueryParams *GetSiteQueryParams) (*GetSiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/site"

	queryString, _ := query.Values(getSiteQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetSiteResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSiteResponse)
	return result, response, err

}

// GetSiteCountQueryParams defines the query parameters for this request
type GetSiteCountQueryParams struct {
	SiteID string `url:"siteId,omitempty"` // Site id to retrieve site count.
}

// GetSiteCount getSiteCount
/* API to get site count
@param siteID Site id to retrieve site count.
*/
func (s *SitesService) GetSiteCount(getSiteCountQueryParams *GetSiteCountQueryParams) (*GetSiteCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/site/count"

	queryString, _ := query.Values(getSiteCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetSiteCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSiteCountResponse)
	return result, response, err

}

// GetSiteHealthQueryParams defines the query parameters for this request
type GetSiteHealthQueryParams struct {
	Timestamp string `url:"timestamp,omitempty"` // Epoch time(in milliseconds) when the Site Hierarchy data is required
}

// GetSiteHealth getSiteHealth
/* Returns Overall Health information for all sites
@param timestamp Epoch time(in milliseconds) when the Site Hierarchy data is required
*/
func (s *SitesService) GetSiteHealth(getSiteHealthQueryParams *GetSiteHealthQueryParams) (*GetSiteHealthResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/site-health"

	queryString, _ := query.Values(getSiteHealthQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetSiteHealthResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSiteHealthResponse)
	return result, response, err

}

// UpdateSite updateSite
/* Update site area/building/floor with specified hierarchy and new values
@param __runsync Enable this parameter to execute the API and return a response synchronously
@param __timeout During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
@param __persistbapioutput Persist bapi sync response
@param siteID Site id to which site details to be updated.
*/
func (s *SitesService) UpdateSite(siteID string, updateSiteRequest *UpdateSiteRequest) (*UpdateSiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/site/{siteID}"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(updateSiteRequest).
		SetResult(&UpdateSiteResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateSiteResponse)
	return result, response, err

}
