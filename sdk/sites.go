package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SitesService is the service to communicate with the Sites API endpoint
type SitesService service

// AssignDeviceToSiteRequest is the assignDeviceToSiteRequest definition
type AssignDeviceToSiteRequest struct {
	Device []AssignDeviceToSiteRequestDevice `json:"device,omitempty"` //
}

// AssignDeviceToSiteRequestDevice is the assignDeviceToSiteRequestDevice definition
type AssignDeviceToSiteRequestDevice struct {
	IP string `json:"ip,omitempty"` //
}

// CreateSiteRequest is the createSiteRequest definition
type CreateSiteRequest struct {
	Site CreateSiteRequestSite `json:"site,omitempty"` //
	Type string                `json:"type,omitempty"` //
}

// CreateSiteRequestSite is the createSiteRequestSite definition
type CreateSiteRequestSite struct {
	Area     CreateSiteRequestSiteArea     `json:"area,omitempty"`     //
	Building CreateSiteRequestSiteBuilding `json:"building,omitempty"` //
	Floor    CreateSiteRequestSiteFloor    `json:"floor,omitempty"`    //
}

// CreateSiteRequestSiteArea is the createSiteRequestSiteArea definition
type CreateSiteRequestSiteArea struct {
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// CreateSiteRequestSiteBuilding is the createSiteRequestSiteBuilding definition
type CreateSiteRequestSiteBuilding struct {
	Address    string `json:"address,omitempty"`    //
	Latitude   int    `json:"latitude,omitempty"`   //
	Longitude  int    `json:"longitude,omitempty"`  //
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// CreateSiteRequestSiteFloor is the createSiteRequestSiteFloor definition
type CreateSiteRequestSiteFloor struct {
	Height     int    `json:"height,omitempty"`     //
	Length     int    `json:"length,omitempty"`     //
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
	RfModel    string `json:"rfModel,omitempty"`    //
	Width      int    `json:"width,omitempty"`      //
}

// UpdateSiteRequest is the updateSiteRequest definition
type UpdateSiteRequest struct {
	Site UpdateSiteRequestSite `json:"site,omitempty"` //
	Type string                `json:"type,omitempty"` //
}

// UpdateSiteRequestSite is the updateSiteRequestSite definition
type UpdateSiteRequestSite struct {
	Area     UpdateSiteRequestSiteArea     `json:"area,omitempty"`     //
	Building UpdateSiteRequestSiteBuilding `json:"building,omitempty"` //
	Floor    UpdateSiteRequestSiteFloor    `json:"floor,omitempty"`    //
}

// UpdateSiteRequestSiteArea is the updateSiteRequestSiteArea definition
type UpdateSiteRequestSiteArea struct {
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// UpdateSiteRequestSiteBuilding is the updateSiteRequestSiteBuilding definition
type UpdateSiteRequestSiteBuilding struct {
	Address    string `json:"address,omitempty"`    //
	Latitude   int    `json:"latitude,omitempty"`   //
	Longitude  int    `json:"longitude,omitempty"`  //
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// UpdateSiteRequestSiteFloor is the updateSiteRequestSiteFloor definition
type UpdateSiteRequestSiteFloor struct {
	Height  int    `json:"height,omitempty"`  //
	Length  int    `json:"length,omitempty"`  //
	Name    string `json:"name,omitempty"`    //
	RfModel string `json:"rfModel,omitempty"` //
	Width   int    `json:"width,omitempty"`   //
}

// AssignDeviceToSiteResponse is the assignDeviceToSiteResponse definition
type AssignDeviceToSiteResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateSiteResponse is the createSiteResponse definition
type CreateSiteResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteSiteResponse is the deleteSiteResponse definition
type DeleteSiteResponse struct {
	Message string `json:"message,omitempty"` //
	Status  string `json:"status,omitempty"`  //
}

// GetMembershipResponse is the getMembershipResponse definition
type GetMembershipResponse struct {
	Device []GetMembershipResponseDevice `json:"device,omitempty"` //
	Site   GetMembershipResponseSite     `json:"site,omitempty"`   //
}

// GetMembershipResponseDevice is the getMembershipResponseDevice definition
type GetMembershipResponseDevice struct {
	Response []string `json:"response,omitempty"` //
	SiteID   string   `json:"siteId,omitempty"`   //
	Version  string   `json:"version,omitempty"`  //
}

// GetMembershipResponseDeviceResponse is the getMembershipResponseDeviceResponse definition
type GetMembershipResponseDeviceResponse []string

// GetMembershipResponseSite is the getMembershipResponseSite definition
type GetMembershipResponseSite struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetMembershipResponseSiteResponse is the getMembershipResponseSiteResponse definition
type GetMembershipResponseSiteResponse []string

// GetSiteCountResponse is the getSiteCountResponse definition
type GetSiteCountResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetSiteHealthResponse is the getSiteHealthResponse definition
type GetSiteHealthResponse struct {
	Response []GetSiteHealthResponseResponse `json:"response,omitempty"` //
}

// GetSiteHealthResponseResponse is the getSiteHealthResponseResponse definition
type GetSiteHealthResponseResponse struct {
	AccessGoodCount                    int                                                 `json:"accessGoodCount,omitempty"`                    //
	AccessTotalCount                   int                                                 `json:"accessTotalCount,omitempty"`                   //
	ApplicationBytesTotalCount         int                                                 `json:"applicationBytesTotalCount,omitempty"`         //
	ApplicationGoodCount               int                                                 `json:"applicationGoodCount,omitempty"`               //
	ApplicationHealth                  int                                                 `json:"applicationHealth,omitempty"`                  //
	ApplicationHealthStats             GetSiteHealthResponseResponseApplicationHealthStats `json:"applicationHealthStats,omitempty"`             //
	ApplicationTotalCount              int                                                 `json:"applicationTotalCount,omitempty"`              //
	ClientHealthWired                  int                                                 `json:"clientHealthWired,omitempty"`                  //
	ClientHealthWireless               int                                                 `json:"clientHealthWireless,omitempty"`               //
	CoreGoodCount                      int                                                 `json:"coreGoodCount,omitempty"`                      //
	CoreTotalCount                     int                                                 `json:"coreTotalCount,omitempty"`                     //
	DistributionGoodCount              int                                                 `json:"distributionGoodCount,omitempty"`              //
	DistributionTotalCount             int                                                 `json:"distributionTotalCount,omitempty"`             //
	DnacInfo                           int                                                 `json:"dnacInfo,omitempty"`                           //
	HealthyClientsPercentage           int                                                 `json:"healthyClientsPercentage,omitempty"`           //
	HealthyNetworkDevicePercentage     int                                                 `json:"healthyNetworkDevicePercentage,omitempty"`     //
	Latitude                           float32                                             `json:"latitude,omitempty"`                           //
	Longitude                          float32                                             `json:"longitude,omitempty"`                          //
	NetworkHealthAccess                int                                                 `json:"networkHealthAccess,omitempty"`                //
	NetworkHealthAverage               int                                                 `json:"networkHealthAverage,omitempty"`               //
	NetworkHealthCore                  int                                                 `json:"networkHealthCore,omitempty"`                  //
	NetworkHealthDistribution          int                                                 `json:"networkHealthDistribution,omitempty"`          //
	NetworkHealthOthers                int                                                 `json:"networkHealthOthers,omitempty"`                //
	NetworkHealthRouter                int                                                 `json:"networkHealthRouter,omitempty"`                //
	NetworkHealthWireless              int                                                 `json:"networkHealthWireless,omitempty"`              //
	NumberOfClients                    int                                                 `json:"numberOfClients,omitempty"`                    //
	NumberOfNetworkDevice              int                                                 `json:"numberOfNetworkDevice,omitempty"`              //
	NumberOfWiredClients               int                                                 `json:"numberOfWiredClients,omitempty"`               //
	NumberOfWirelessClients            int                                                 `json:"numberOfWirelessClients,omitempty"`            //
	OverallGoodDevices                 int                                                 `json:"overallGoodDevices,omitempty"`                 //
	ParentSiteID                       string                                              `json:"parentSiteId,omitempty"`                       //
	ParentSiteName                     string                                              `json:"parentSiteName,omitempty"`                     //
	RouterGoodCount                    int                                                 `json:"routerGoodCount,omitempty"`                    //
	RouterTotalCount                   int                                                 `json:"routerTotalCount,omitempty"`                   //
	SiteID                             string                                              `json:"siteId,omitempty"`                             //
	SiteName                           string                                              `json:"siteName,omitempty"`                           //
	SiteType                           string                                              `json:"siteType,omitempty"`                           //
	TotalNumberOfActiveWirelessClients int                                                 `json:"totalNumberOfActiveWirelessClients,omitempty"` //
	TotalNumberOfConnectedWiredClients int                                                 `json:"totalNumberOfConnectedWiredClients,omitempty"` //
	WiredGoodClients                   int                                                 `json:"wiredGoodClients,omitempty"`                   //
	WirelessDeviceGoodCount            int                                                 `json:"wirelessDeviceGoodCount,omitempty"`            //
	WirelessDeviceTotalCount           int                                                 `json:"wirelessDeviceTotalCount,omitempty"`           //
	WirelessGoodClients                int                                                 `json:"wirelessGoodClients,omitempty"`                //
}

// GetSiteHealthResponseResponseApplicationHealthStats is the getSiteHealthResponseResponseApplicationHealthStats definition
type GetSiteHealthResponseResponseApplicationHealthStats struct {
	AppTotalCount              int                                                                           `json:"appTotalCount,omitempty"`              //
	BusinessIrrelevantAppCount GetSiteHealthResponseResponseApplicationHealthStatsBusinessIrrelevantAppCount `json:"businessIrrelevantAppCount,omitempty"` //
	BusinessRelevantAppCount   GetSiteHealthResponseResponseApplicationHealthStatsBusinessRelevantAppCount   `json:"businessRelevantAppCount,omitempty"`   //
	DefaultHealthAppCount      GetSiteHealthResponseResponseApplicationHealthStatsDefaultHealthAppCount      `json:"defaultHealthAppCount,omitempty"`      //
}

// GetSiteHealthResponseResponseApplicationHealthStatsBusinessIrrelevantAppCount is the getSiteHealthResponseResponseApplicationHealthStatsBusinessIrrelevantAppCount definition
type GetSiteHealthResponseResponseApplicationHealthStatsBusinessIrrelevantAppCount struct {
	Fair int `json:"fair,omitempty"` //
	Good int `json:"good,omitempty"` //
	Poor int `json:"poor,omitempty"` //
}

// GetSiteHealthResponseResponseApplicationHealthStatsBusinessRelevantAppCount is the getSiteHealthResponseResponseApplicationHealthStatsBusinessRelevantAppCount definition
type GetSiteHealthResponseResponseApplicationHealthStatsBusinessRelevantAppCount struct {
	Fair int `json:"fair,omitempty"` //
	Good int `json:"good,omitempty"` //
	Poor int `json:"poor,omitempty"` //
}

// GetSiteHealthResponseResponseApplicationHealthStatsDefaultHealthAppCount is the getSiteHealthResponseResponseApplicationHealthStatsDefaultHealthAppCount definition
type GetSiteHealthResponseResponseApplicationHealthStatsDefaultHealthAppCount struct {
	Fair int `json:"fair,omitempty"` //
	Good int `json:"good,omitempty"` //
	Poor int `json:"poor,omitempty"` //
}

// GetSiteResponse is the getSiteResponse definition
type GetSiteResponse struct {
	Response []GetSiteResponseResponse `json:"response,omitempty"` //
}

// GetSiteResponseResponse is the getSiteResponseResponse definition
type GetSiteResponseResponse struct {
	AdditionalInfo    []string `json:"additionalInfo,omitempty"`    //
	ID                string   `json:"id,omitempty"`                //
	InstanceTenantID  string   `json:"instanceTenantId,omitempty"`  //
	Name              string   `json:"name,omitempty"`              //
	ParentID          string   `json:"parentId,omitempty"`          //
	SiteHierarchy     string   `json:"siteHierarchy,omitempty"`     //
	SiteNameHierarchy string   `json:"siteNameHierarchy,omitempty"` //
}

// GetSiteResponseResponseAdditionalInfo is the getSiteResponseResponseAdditionalInfo definition
type GetSiteResponseResponseAdditionalInfo []string

// UpdateSiteResponse is the updateSiteResponse definition
type UpdateSiteResponse struct {
	Response UpdateSiteResponseResponse `json:"response,omitempty"` //
	Result   string                     `json:"result,omitempty"`   //
	Status   string                     `json:"status,omitempty"`   //
}

// UpdateSiteResponseResponse is the updateSiteResponseResponse definition
type UpdateSiteResponseResponse struct {
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

// UpdateSiteResponseResponseOperationIDList is the updateSiteResponseResponseOperationIDList definition
type UpdateSiteResponseResponseOperationIDList []string

// AssignDeviceToSite assignDeviceToSite
/* Assigns list of devices to a site
@param __runsync Enable this parameter to execute the API and return a response synchronously
@param __persistbapioutput Persist bapi sync response
@param __runsynctimeout During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
@param siteID Site id to which site the device to assign
*/
func (s *SitesService) AssignDeviceToSite(siteID string, assignDeviceToSiteRequest *AssignDeviceToSiteRequest) (*AssignDeviceToSiteResponse, *resty.Response, error) {

	path := "/dna/system/api/v1/site/{siteId}/device"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

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
/* Delete site with area/building/floor by siteId.
@param siteID Site id to which site details to be deleted.
*/
func (s *SitesService) DeleteSite(siteID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/site/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

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
	SerialNumber string `url:"serialNumber,omitempty"` // Device serial number
}

// GetMembership getMembership
/* Getting the site children details and device details.
@param siteID Site id to retrieve device associated with the site.
@param offset offset/starting row
@param limit Number of sites to be retrieved
@param deviceFamily Device family name
@param serialNumber Device serial number
*/
func (s *SitesService) GetMembership(siteID string, getMembershipQueryParams *GetMembershipQueryParams) (*GetMembershipResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/membership/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

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

	path := "/dna/intent/api/v1/site/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

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
