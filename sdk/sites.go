package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SitesService service

type AssignDevicesToSiteHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
	Runsynctimeout    string `url:"__runsynctimeout,omitempty"`    //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
}
type GetMembershipQueryParams struct {
	Offset       int    `url:"offset,omitempty"`       //offset/starting row
	Limit        int    `url:"limit,omitempty"`        //Number of sites to be retrieved
	DeviceFamily string `url:"deviceFamily,omitempty"` //Device family name
	SerialNumber string `url:"serialNumber,omitempty"` //Device serial number
}
type CreateSiteHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetSiteQueryParams struct {
	Name   string `url:"name,omitempty"`   //siteNameHierarchy (ex: global/groupName)
	SiteID string `url:"siteId,omitempty"` //Site id to which site details to retrieve.
	Type   string `url:"type,omitempty"`   //type (ex: area, building, floor)
	Offset int    `url:"offset,omitempty"` //offset/starting row. The default value is 1
	Limit  int    `url:"limit,omitempty"`  //Number of sites to be retrieved. The default value is 500
}
type GetSiteHealthQueryParams struct {
	Timestamp string  `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Site Hierarchy data is required
	SiteType  string  `url:"siteType,omitempty"`  //Type of the site to return.  AREA or BUILDING.  Default to AREA
	Offset    float64 `url:"offset,omitempty"`    //The offset value, starting from 1, of the first returned site entry.  Default is 1.
	Limit     float64 `url:"limit,omitempty"`     //The max number of sites in the returned data set.  Default is 25, and max at 50
}
type GetSiteCountQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site id to retrieve site count.
}
type UpdateSiteHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}

type ResponseSitesAssignDevicesToSite struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesGetMembership struct {
	Site   *ResponseSitesGetMembershipSite     `json:"site,omitempty"`   //
	Device *[]ResponseSitesGetMembershipDevice `json:"device,omitempty"` //
}
type ResponseSitesGetMembershipSite struct {
	Response *[]ResponseSitesGetMembershipSiteResponse `json:"response,omitempty"` // Response
	Version  string                                    `json:"version,omitempty"`  // Version
}
type ResponseSitesGetMembershipSiteResponse interface{}
type ResponseSitesGetMembershipDevice struct {
	Response *[]ResponseSitesGetMembershipDeviceResponse `json:"response,omitempty"` // Response
	Version  string                                      `json:"version,omitempty"`  // Version
	SiteID   string                                      `json:"siteId,omitempty"`   // Site Id
}
type ResponseSitesGetMembershipDeviceResponse interface{}
type ResponseSitesCreateSite struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesGetSite struct {
	Response *[]ResponseSitesGetSiteResponse `json:"response,omitempty"` //
}
type ResponseSitesGetSiteResponse struct {
	ParentID          string                                       `json:"parentId,omitempty"`          // Parent Id
	Name              string                                       `json:"name,omitempty"`              // Name
	AdditionalInfo    []ResponseSitesGetSiteResponseAdditionalInfo `json:"additionalInfo,omitempty"`    //
	SiteHierarchy     string                                       `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteNameHierarchy string                                       `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	InstanceTenantID  string                                       `json:"instanceTenantId,omitempty"`  // Instance Tenant Id
	ID                string                                       `json:"id,omitempty"`                // Id
}
type ResponseSitesGetSiteResponseAdditionalInfo struct {
	Namespace  string                                               `json:"nameSpace,omitempty"`  //
	Attributes ResponseSitesGetSiteResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}

//type ResponseSitesGetSiteResponseAdditionalInfoAttributes map[string]string

type ResponseSitesGetSiteResponseAdditionalInfoAttributes struct {
	Country              string `json:"country,omitempty"`              //
	Address              string `json:"address,omitempty"`              //
	Latitude             string `json:"latitude,omitempty"`             //
	AddressInheritedFrom string `json:"addressInheritedFrom,omitempty"` //
	Type                 string `json:"type,omitempty"`                 //
	Longitude            string `json:"longitude,omitempty"`            //
	OffsetX              string `json:"offsetX,omitempty"`              //
	OffsetY              string `json:"offsetY,omitempty"`              //
	Length               string `json:"length,omitempty"`               //
	Width                string `json:"width,omitempty"`                //
	Height               string `json:"height,omitempty"`               //
	RfModel              string `json:"rfModel,omitempty"`              //
	FloorIndex           string `json:"floorIndex,omitempty"`           //
}

// Area
type ResponseSitesGetArea struct {
	Response *[]ResponseSitesGetAreaResponse `json:"response,omitempty"` //
}
type ResponseSitesGetAreaResponse struct {
	ParentID          string                                       `json:"parentId,omitempty"`          // Parent Id
	Name              string                                       `json:"name,omitempty"`              // Name
	AdditionalInfo    []ResponseSitesGetAreaResponseAdditionalInfo `json:"additionalInfo,omitempty"`    //
	SiteHierarchy     string                                       `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteNameHierarchy string                                       `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	InstanceTenantID  string                                       `json:"instanceTenantId,omitempty"`  // Instance Tenant Id
	ID                string                                       `json:"id,omitempty"`
	ParentName        string                                       `json:"parent_name,omitempty"` // Id
}
type ResponseSitesGetAreaResponseAdditionalInfo struct {
	Namespace  string                                               `json:"nameSpace,omitempty"`  //
	Attributes ResponseSitesGetAreaResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}

//type ResponseSitesGetSiteResponseAdditionalInfoAttributes map[string]string

type ResponseSitesGetAreaResponseAdditionalInfoAttributes struct {
	Name                 string `json:"name,omitempty"` //
	ParentName           string `json:"parent_name,omitempty"`
	AddressInheritedFrom string `json:"addressinheritedfrom,omitempty"` //
	Type                 string `json:"type,omitempty"`                 //
}

// Floor
type ResponseSitesGetFloor struct {
	Response *[]ResponseSitesGetFloorResponse `json:"response,omitempty"` //
}
type ResponseSitesGetFloorResponse struct {
	ParentID          string                                        `json:"parentId,omitempty"`          // Parent Id
	Name              string                                        `json:"name,omitempty"`              // Name
	AdditionalInfo    []ResponseSitesGetFloorResponseAdditionalInfo `json:"additionalInfo,omitempty"`    //
	SiteHierarchy     string                                        `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteNameHierarchy string                                        `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
	InstanceTenantID  string                                        `json:"instanceTenantId,omitempty"`  // Instance Tenant Id
	ID                string                                        `json:"id,omitempty"`
	ParentName        string                                        `json:"parent_name,omitempty"` // Id
}
type ResponseSitesGetFloorResponseAdditionalInfo struct {
	Namespace  string                                                `json:"nameSpace,omitempty"`  //
	Attributes ResponseSitesGetFloorResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}

type ResponseSitesGetFloorResponseAdditionalInfoAttributes struct {
	FloorIndex string `json:"floorIndex,omitempty"` //
	Height     string `json:"height,omitempty"`
	Length     string `json:"length,omitempty"`      //
	Name       string `json:"name,omitempty"`        //
	ParentName string `json:"parent_name,omitempty"` //
	RfModel    string `json:"rfmodel,omitempty"`     //
	Width      string `json:"width,omitempty"`       //
}

type ResponseSitesGetSiteHealth struct {
	Response *[]ResponseSitesGetSiteHealthResponse `json:"response,omitempty"` //
}
type ResponseSitesGetSiteHealthResponse struct {
	SiteName                           string                                                                `json:"siteName,omitempty"`                           // Site Name
	SiteID                             string                                                                `json:"siteId,omitempty"`                             // Site Id
	ParentSiteID                       string                                                                `json:"parentSiteId,omitempty"`                       // Parent Site Id
	ParentSiteName                     string                                                                `json:"parentSiteName,omitempty"`                     // Parent Site Name
	SiteType                           string                                                                `json:"siteType,omitempty"`                           // Site Type
	Latitude                           *float64                                                              `json:"latitude,omitempty"`                           // Latitude
	Longitude                          *float64                                                              `json:"longitude,omitempty"`                          // Longitude
	HealthyNetworkDevicePercentage     *ResponseSitesGetSiteHealthResponseHealthyNetworkDevicePercentage     `json:"healthyNetworkDevicePercentage,omitempty"`     // Healthy Network Device Percentage
	HealthyClientsPercentage           *ResponseSitesGetSiteHealthResponseHealthyClientsPercentage           `json:"healthyClientsPercentage,omitempty"`           // Healthy Clients Percentage
	ClientHealthWired                  *ResponseSitesGetSiteHealthResponseClientHealthWired                  `json:"clientHealthWired,omitempty"`                  // Client Health Wired
	ClientHealthWireless               *ResponseSitesGetSiteHealthResponseClientHealthWireless               `json:"clientHealthWireless,omitempty"`               // Client Health Wireless
	NumberOfClients                    *ResponseSitesGetSiteHealthResponseNumberOfClients                    `json:"numberOfClients,omitempty"`                    // Number Of Clients
	NumberOfNetworkDevice              *ResponseSitesGetSiteHealthResponseNumberOfNetworkDevice              `json:"numberOfNetworkDevice,omitempty"`              // Number Of Network Device
	NetworkHealthAverage               *ResponseSitesGetSiteHealthResponseNetworkHealthAverage               `json:"networkHealthAverage,omitempty"`               // Network Health Average
	NetworkHealthAccess                *ResponseSitesGetSiteHealthResponseNetworkHealthAccess                `json:"networkHealthAccess,omitempty"`                // Network Health Access
	NetworkHealthCore                  *ResponseSitesGetSiteHealthResponseNetworkHealthCore                  `json:"networkHealthCore,omitempty"`                  // Network Health Core
	NetworkHealthDistribution          *ResponseSitesGetSiteHealthResponseNetworkHealthDistribution          `json:"networkHealthDistribution,omitempty"`          // Network Health Distribution
	NetworkHealthRouter                *ResponseSitesGetSiteHealthResponseNetworkHealthRouter                `json:"networkHealthRouter,omitempty"`                // Network Health Router
	NetworkHealthWireless              *ResponseSitesGetSiteHealthResponseNetworkHealthWireless              `json:"networkHealthWireless,omitempty"`              // Network Health Wireless
	NetworkHealthOthers                *ResponseSitesGetSiteHealthResponseNetworkHealthOthers                `json:"networkHealthOthers,omitempty"`                // Network Health Others
	NumberOfWiredClients               *ResponseSitesGetSiteHealthResponseNumberOfWiredClients               `json:"numberOfWiredClients,omitempty"`               // Number Of Wired Clients
	NumberOfWirelessClients            *ResponseSitesGetSiteHealthResponseNumberOfWirelessClients            `json:"numberOfWirelessClients,omitempty"`            // Number Of Wireless Clients
	TotalNumberOfConnectedWiredClients *ResponseSitesGetSiteHealthResponseTotalNumberOfConnectedWiredClients `json:"totalNumberOfConnectedWiredClients,omitempty"` // Total Number Of Connected Wired Clients
	TotalNumberOfActiveWirelessClients *ResponseSitesGetSiteHealthResponseTotalNumberOfActiveWirelessClients `json:"totalNumberOfActiveWirelessClients,omitempty"` // Total Number Of Active Wireless Clients
	WiredGoodClients                   *ResponseSitesGetSiteHealthResponseWiredGoodClients                   `json:"wiredGoodClients,omitempty"`                   // Wired Good Clients
	WirelessGoodClients                *ResponseSitesGetSiteHealthResponseWirelessGoodClients                `json:"wirelessGoodClients,omitempty"`                // Wireless Good Clients
	OverallGoodDevices                 *ResponseSitesGetSiteHealthResponseOverallGoodDevices                 `json:"overallGoodDevices,omitempty"`                 // Overall Good Devices
	AccessGoodCount                    *ResponseSitesGetSiteHealthResponseAccessGoodCount                    `json:"accessGoodCount,omitempty"`                    // Access Good Count
	AccessTotalCount                   *ResponseSitesGetSiteHealthResponseAccessTotalCount                   `json:"accessTotalCount,omitempty"`                   // Access Total Count
	CoreGoodCount                      *ResponseSitesGetSiteHealthResponseCoreGoodCount                      `json:"coreGoodCount,omitempty"`                      // Core Good Count
	CoreTotalCount                     *ResponseSitesGetSiteHealthResponseCoreTotalCount                     `json:"coreTotalCount,omitempty"`                     // Core Total Count
	DistributionGoodCount              *ResponseSitesGetSiteHealthResponseDistributionGoodCount              `json:"distributionGoodCount,omitempty"`              // Distribution Good Count
	DistributionTotalCount             *ResponseSitesGetSiteHealthResponseDistributionTotalCount             `json:"distributionTotalCount,omitempty"`             // Distribution Total Count
	RouterGoodCount                    *ResponseSitesGetSiteHealthResponseRouterGoodCount                    `json:"routerGoodCount,omitempty"`                    // Router Good Count
	RouterTotalCount                   *ResponseSitesGetSiteHealthResponseRouterTotalCount                   `json:"routerTotalCount,omitempty"`                   // Router Total Count
	WirelessDeviceGoodCount            *ResponseSitesGetSiteHealthResponseWirelessDeviceGoodCount            `json:"wirelessDeviceGoodCount,omitempty"`            // Wireless Device Good Count
	WirelessDeviceTotalCount           *ResponseSitesGetSiteHealthResponseWirelessDeviceTotalCount           `json:"wirelessDeviceTotalCount,omitempty"`           // Wireless Device Total Count
	ApplicationHealth                  *ResponseSitesGetSiteHealthResponseApplicationHealth                  `json:"applicationHealth,omitempty"`                  // Application Health
	ApplicationGoodCount               *ResponseSitesGetSiteHealthResponseApplicationGoodCount               `json:"applicationGoodCount,omitempty"`               // Application Good Count
	ApplicationTotalCount              *ResponseSitesGetSiteHealthResponseApplicationTotalCount              `json:"applicationTotalCount,omitempty"`              // Application Total Count
	ApplicationBytesTotalCount         *ResponseSitesGetSiteHealthResponseApplicationBytesTotalCount         `json:"applicationBytesTotalCount,omitempty"`         // Application Bytes Total Count
	DnacInfo                           *ResponseSitesGetSiteHealthResponseDnacInfo                           `json:"dnacInfo,omitempty"`                           // Dnac Info
	ApplicationHealthStats             *ResponseSitesGetSiteHealthResponseApplicationHealthStats             `json:"applicationHealthStats,omitempty"`             //
}
type ResponseSitesGetSiteHealthResponseHealthyNetworkDevicePercentage interface{}
type ResponseSitesGetSiteHealthResponseHealthyClientsPercentage interface{}
type ResponseSitesGetSiteHealthResponseClientHealthWired interface{}
type ResponseSitesGetSiteHealthResponseClientHealthWireless interface{}
type ResponseSitesGetSiteHealthResponseNumberOfClients interface{}
type ResponseSitesGetSiteHealthResponseNumberOfNetworkDevice interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthAverage interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthAccess interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthCore interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthDistribution interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthRouter interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthWireless interface{}
type ResponseSitesGetSiteHealthResponseNetworkHealthOthers interface{}
type ResponseSitesGetSiteHealthResponseNumberOfWiredClients interface{}
type ResponseSitesGetSiteHealthResponseNumberOfWirelessClients interface{}
type ResponseSitesGetSiteHealthResponseTotalNumberOfConnectedWiredClients interface{}
type ResponseSitesGetSiteHealthResponseTotalNumberOfActiveWirelessClients interface{}
type ResponseSitesGetSiteHealthResponseWiredGoodClients interface{}
type ResponseSitesGetSiteHealthResponseWirelessGoodClients interface{}
type ResponseSitesGetSiteHealthResponseOverallGoodDevices interface{}
type ResponseSitesGetSiteHealthResponseAccessGoodCount interface{}
type ResponseSitesGetSiteHealthResponseAccessTotalCount interface{}
type ResponseSitesGetSiteHealthResponseCoreGoodCount interface{}
type ResponseSitesGetSiteHealthResponseCoreTotalCount interface{}
type ResponseSitesGetSiteHealthResponseDistributionGoodCount interface{}
type ResponseSitesGetSiteHealthResponseDistributionTotalCount interface{}
type ResponseSitesGetSiteHealthResponseRouterGoodCount interface{}
type ResponseSitesGetSiteHealthResponseRouterTotalCount interface{}
type ResponseSitesGetSiteHealthResponseWirelessDeviceGoodCount interface{}
type ResponseSitesGetSiteHealthResponseWirelessDeviceTotalCount interface{}
type ResponseSitesGetSiteHealthResponseApplicationHealth interface{}
type ResponseSitesGetSiteHealthResponseApplicationGoodCount interface{}
type ResponseSitesGetSiteHealthResponseApplicationTotalCount interface{}
type ResponseSitesGetSiteHealthResponseApplicationBytesTotalCount interface{}
type ResponseSitesGetSiteHealthResponseDnacInfo interface{}
type ResponseSitesGetSiteHealthResponseApplicationHealthStats struct {
	AppTotalCount              *float64                                                                            `json:"appTotalCount,omitempty"`              // App Total Count
	BusinessRelevantAppCount   *ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessRelevantAppCount   `json:"businessRelevantAppCount,omitempty"`   //
	BusinessIrrelevantAppCount *ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessIrrelevantAppCount `json:"businessIrrelevantAppCount,omitempty"` //
	DefaultHealthAppCount      *ResponseSitesGetSiteHealthResponseApplicationHealthStatsDefaultHealthAppCount      `json:"defaultHealthAppCount,omitempty"`      //
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessRelevantAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor
	Fair *float64 `json:"fair,omitempty"` // Fair
	Good *float64 `json:"good,omitempty"` // Good
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessIrrelevantAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor
	Fair *float64 `json:"fair,omitempty"` // Fair
	Good *float64 `json:"good,omitempty"` // Good
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStatsDefaultHealthAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor
	Fair *float64 `json:"fair,omitempty"` // Fair
	Good *float64 `json:"good,omitempty"` // Good
}
type ResponseSitesGetSiteCount struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseSitesUpdateSite struct {
	Result   string                           `json:"result,omitempty"`   // Result
	Response *ResponseSitesUpdateSiteResponse `json:"response,omitempty"` //
	Status   string                           `json:"status,omitempty"`   // Status
}
type ResponseSitesUpdateSiteResponse struct {
	EndTime          string   `json:"endTime,omitempty"`          // End Time
	Version          string   `json:"version,omitempty"`          // Version
	StartTime        string   `json:"startTime,omitempty"`        // Start Time
	Progress         string   `json:"progress,omitempty"`         // Progress
	Data             string   `json:"data,omitempty"`             // Data
	ServiceType      string   `json:"serviceType,omitempty"`      // Service Type
	OperationIDList  []string `json:"operationIdList,omitempty"`  // Operation Id List
	IsError          string   `json:"isError,omitempty"`          // Is Error
	RootID           string   `json:"rootId,omitempty"`           // Root Id
	InstanceTenantID string   `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string   `json:"id,omitempty"`               // Id
}
type ResponseSitesDeleteSite struct {
	ExecutionID        string `json:"executionId,omitempty"`        // executionId
	ExecutionStatusURL string `json:"executionStatusURL,omitempty"` // executionStatusURL
	Message            string `json:"message,omitempty"`            // Message
}
type RequestSitesAssignDevicesToSite struct {
	Device *[]RequestSitesAssignDevicesToSiteDevice `json:"device,omitempty"` //
}
type RequestSitesAssignDevicesToSiteDevice struct {
	IP string `json:"ip,omitempty"` // Device ip (eg: 10.104.240.64)
}
type RequestSitesCreateSite struct {
	Type string                      `json:"type,omitempty"` // Type of site to create (eg: area, building, floor)
	Site *RequestSitesCreateSiteSite `json:"site,omitempty"` //
}
type RequestSitesCreateSiteSite struct {
	Area     *RequestSitesCreateSiteSiteArea     `json:"area,omitempty"`     //
	Building *RequestSitesCreateSiteSiteBuilding `json:"building,omitempty"` //
	Floor    *RequestSitesCreateSiteSiteFloor    `json:"floor,omitempty"`    //
}
type RequestSitesCreateSiteSiteArea struct {
	Name       string `json:"name,omitempty"`       // Name of the area (eg: Area1)
	ParentName string `json:"parentName,omitempty"` // Parent name of the area to be created
}
type RequestSitesCreateSiteSiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Name of the building (eg: building1)
	Address    string   `json:"address,omitempty"`    // Address of the building to be created
	ParentName string   `json:"parentName,omitempty"` // Parent name of building to be created
	Latitude   *float64 `json:"latitude,omitempty"`   // Latitude coordinate of the building (eg:37.338)
	Longitude  *float64 `json:"longitude,omitempty"`  // Longitude coordinate of the building (eg:-121.832)
	Country    string   `json:"country,omitempty"`    // Country (eg:United States)
}
type RequestSitesCreateSiteSiteFloor struct {
	Name        string   `json:"name,omitempty"`        // Name of the floor (eg:floor-1)
	ParentName  string   `json:"parentName,omitempty"`  // Parent name of the floor to be created
	RfModel     string   `json:"rfModel,omitempty"`     // Type of floor (eg: Cubes And Walled Offices0
	Width       *float64 `json:"width,omitempty"`       // Width of the floor. Unit of measure is ft. (eg: 100)
	Length      *float64 `json:"length,omitempty"`      // Length of the floor. Unit of measure is ft. (eg: 100)
	Height      *float64 `json:"height,omitempty"`      // Height of the floor. Unit of measure is ft. (eg: 15)
	FloorNumber *float64 `json:"floorNumber,omitempty"` // Floor number. (eg: 5)
}
type RequestSitesUpdateSite struct {
	Type string                      `json:"type,omitempty"` // Type
	Site *RequestSitesUpdateSiteSite `json:"site,omitempty"` //
}
type RequestSitesUpdateSiteSite struct {
	Area     *RequestSitesUpdateSiteSiteArea     `json:"area,omitempty"`     //
	Building *RequestSitesUpdateSiteSiteBuilding `json:"building,omitempty"` //
	Floor    *RequestSitesUpdateSiteSiteFloor    `json:"floor,omitempty"`    //
}
type RequestSitesUpdateSiteSiteArea struct {
	Name       string `json:"name,omitempty"`       // Name
	ParentName string `json:"parentName,omitempty"` // Parent Name
}
type RequestSitesUpdateSiteSiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Name
	Address    string   `json:"address,omitempty"`    // Address
	ParentName string   `json:"parentName,omitempty"` // Parent Name
	Latitude   *float64 `json:"latitude,omitempty"`   // Latitude
	Longitude  *float64 `json:"longitude,omitempty"`  // Longitude
}
type RequestSitesUpdateSiteSiteFloor struct {
	Name    string   `json:"name,omitempty"`    // Name
	RfModel string   `json:"rfModel,omitempty"` // Rf Model
	Width   *float64 `json:"width,omitempty"`   // Width
	Length  *float64 `json:"length,omitempty"`  // Length
	Height  *float64 `json:"height,omitempty"`  // Height
}

//GetMembership Get Membership - eba6-6905-4e08-a60e
/* Getting the site children details and device details.


@param siteID siteId path parameter. Site id to retrieve device associated with the site.

@param GetMembershipQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-membership
*/
func (s *SitesService) GetMembership(siteID string, GetMembershipQueryParams *GetMembershipQueryParams) (*ResponseSitesGetMembership, *resty.Response, error) {
	path := "/dna/intent/api/v1/membership/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(GetMembershipQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetMembership{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMembership(siteID, GetMembershipQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMembership")
	}

	result := response.Result().(*ResponseSitesGetMembership)
	return result, response, err

}

//GetSite Get Site - 6fb4-ab36-43fa-a80f
/* Get site using siteNameHierarchy/siteId/type ,return all sites if these parameters are not given as input.


@param GetSiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site
*/
func (s *SitesService) GetSite(GetSiteQueryParams *GetSiteQueryParams) (*ResponseSitesGetSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/site"

	queryString, _ := query.Values(GetSiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSite(GetSiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSite")
	}

	result := response.Result().(*ResponseSitesGetSite)
	return result, response, err

}

//GetSiteHealth Get Site Health - 15b7-aa0c-4dda-8e85
/* Returns Overall Health information for all sites


@param GetSiteHealthQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-health
*/
func (s *SitesService) GetSiteHealth(GetSiteHealthQueryParams *GetSiteHealthQueryParams) (*ResponseSitesGetSiteHealth, *resty.Response, error) {
	path := "/dna/intent/api/v1/site-health"

	queryString, _ := query.Values(GetSiteHealthQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteHealth{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteHealth(GetSiteHealthQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteHealth")
	}

	result := response.Result().(*ResponseSitesGetSiteHealth)
	return result, response, err

}

//GetSiteCount Get Site Count - b0b7-eabc-4f4b-9b28
/* API to get site count


@param GetSiteCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-count
*/
func (s *SitesService) GetSiteCount(GetSiteCountQueryParams *GetSiteCountQueryParams) (*ResponseSitesGetSiteCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/site/count"

	queryString, _ := query.Values(GetSiteCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteCount(GetSiteCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteCount")
	}

	result := response.Result().(*ResponseSitesGetSiteCount)
	return result, response, err

}

//AssignDevicesToSite Assign Devices To Site - eeb1-68eb-4198-8e07
/* Assigns unassigned devices to a site. This API does not move assigned devices to other sites.


@param siteID siteId path parameter. Site id to which site the device to assign

@param AssignDevicesToSiteHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-devices-to-site
*/
func (s *SitesService) AssignDevicesToSite(siteID string, requestSitesAssignDevicesToSite *RequestSitesAssignDevicesToSite, AssignDevicesToSiteHeaderParams *AssignDevicesToSiteHeaderParams) (*ResponseSitesAssignDevicesToSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/assign-device-to-site/{siteId}/device"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if AssignDevicesToSiteHeaderParams != nil {

		if AssignDevicesToSiteHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", AssignDevicesToSiteHeaderParams.Runsync)
		}

		if AssignDevicesToSiteHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", AssignDevicesToSiteHeaderParams.Persistbapioutput)
		}

		if AssignDevicesToSiteHeaderParams.Runsynctimeout != "" {
			clientRequest = clientRequest.SetHeader("__runsynctimeout", AssignDevicesToSiteHeaderParams.Runsynctimeout)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesAssignDevicesToSite).
		SetResult(&ResponseSitesAssignDevicesToSite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignDevicesToSite(siteID, requestSitesAssignDevicesToSite, AssignDevicesToSiteHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation AssignDevicesToSite")
	}

	result := response.Result().(*ResponseSitesAssignDevicesToSite)
	return result, response, err

}

//CreateSite Create Site - 50b5-89fd-4c7a-930a
/* Creates site with area/building/floor with specified hierarchy.


@param CreateSiteHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-site
*/
func (s *SitesService) CreateSite(requestSitesCreateSite *RequestSitesCreateSite, CreateSiteHeaderParams *CreateSiteHeaderParams) (*ResponseSitesCreateSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/site"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateSiteHeaderParams != nil {

		if CreateSiteHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", CreateSiteHeaderParams.Runsync)
		}

		if CreateSiteHeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", CreateSiteHeaderParams.Timeout)
		}

		if CreateSiteHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", CreateSiteHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesCreateSite).
		SetResult(&ResponseSitesCreateSite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSite(requestSitesCreateSite, CreateSiteHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreateSite")
	}

	result := response.Result().(*ResponseSitesCreateSite)
	return result, response, err

}

//UpdateSite Update Site - eeb7-eb4b-4bd8-a1dd
/* Update site area/building/floor with specified hierarchy and new values


@param siteID siteId path parameter. Site id to which site details to be updated.

@param UpdateSiteHeaderParams Custom header parameters
*/
func (s *SitesService) UpdateSite(siteID string, requestSitesUpdateSite *RequestSitesUpdateSite, UpdateSiteHeaderParams *UpdateSiteHeaderParams) (*ResponseSitesUpdateSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateSiteHeaderParams != nil {

		if UpdateSiteHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", UpdateSiteHeaderParams.Runsync)
		}

		if UpdateSiteHeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", UpdateSiteHeaderParams.Timeout)
		}

		if UpdateSiteHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", UpdateSiteHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesUpdateSite).
		SetResult(&ResponseSitesUpdateSite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSite(siteID, requestSitesUpdateSite, UpdateSiteHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSite")
	}

	result := response.Result().(*ResponseSitesUpdateSite)
	return result, response, err

}

//DeleteSite Delete Site - f083-cb13-484a-8fae
/* Delete site with area/building/floor by siteId.


@param siteID siteId path parameter. Site id to which site details to be deleted.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-site
*/
func (s *SitesService) DeleteSite(siteID string) (*ResponseSitesDeleteSite, *resty.Response, error) {
	//siteID string
	path := "/dna/intent/api/v1/site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesDeleteSite{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSite(siteID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSite")
	}

	result := response.Result().(*ResponseSitesDeleteSite)
	return result, response, err

}
