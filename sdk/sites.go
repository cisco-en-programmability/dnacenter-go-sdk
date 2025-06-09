package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SitesService service

type GetSitesEnergyQueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to one day before `endTime`.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return
	Offset          float64 `url:"offset,omitempty"`          //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy          string  `url:"sortBy,omitempty"`          //A field within the response to sort by.
	Order           string  `url:"order,omitempty"`           //The sort order of the field ascending or descending.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteName        string  `url:"siteName,omitempty"`        //The name of the site. (Ex. `FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*San*, *San, San*` Examples: `?siteName=building1` (single siteName requested) `?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	DeviceCategory  string  `url:"deviceCategory,omitempty"`  //The list of device categories. Note that this filter specifies which devices will be included when calculating energy consumption values, rather than specifying the list of returned sites. Examples: `deviceCategory=Switch` (single device category requested) `deviceCategory=Switch&deviceCategory=Router` (multiple device categories with comma separator)
	SiteID          string  `url:"siteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	Views           string  `url:"views,omitempty"`           //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view: 1. **Site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude] 2. **Energy** [energyConsumed, estimatedCost, estimatedEmission, carbonIntensity, numberOfDevices] When this query parameter is not added the default summaries are: **[site,energy]** Examples: views=site (single view requested) views=site,energy (multiple views requested)
	Attribute       string  `url:"attribute,omitempty"`       //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, energyConsumed, estimatedCost, estimatedEmission, carbonIntensity, numberOfDevices] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=energyConsumed (multiple attributes requested)
	TaskID          string  `url:"taskId,omitempty"`          //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type GetSitesEnergyHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountSitesEnergyQueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to one day before `endTime`.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteName        string  `url:"siteName,omitempty"`        //The name of the site. (Ex. `FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*San*, *San, San*` Examples: `?siteName=building1` (single siteName requested) `?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	DeviceCategory  string  `url:"deviceCategory,omitempty"`  //The list of device categories. Note that this filter specifies which devices will be included when calculating energy consumption values, rather than specifying the list of returned sites. Examples: `deviceCategory=Switch` (single device category requested) `deviceCategory=Switch&deviceCategory=Router` (multiple device categories with comma separator)
	SiteID          string  `url:"siteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	TaskID          string  `url:"taskId,omitempty"`          //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type CountSitesEnergyHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type SubmitRequestToQuerySitesEnergyQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type SubmitRequestToQuerySitesEnergyHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QuerySitesEnergyForTheGivenTaskIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type CountSitesEnergyForTheGivenTaskIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type CountSitesEnergyForTheGivenTaskIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type SubmitRequestToCountSitesEnergyFromQueryQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type GetSiteEnergyByIDQueryParams struct {
	StartTime      float64 `url:"startTime,omitempty"`      //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to one day before `endTime`.
	EndTime        float64 `url:"endTime,omitempty"`        //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.
	Views          string  `url:"views,omitempty"`          //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view: 1. **Site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude] 2. **Energy** [energyConsumed, estimatedCost, estimatedEmission, carbonIntensity, numberOfDevices] When this query parameter is not added the default summaries are: **[site,energy]** Examples: views=site (single view requested) views=site,energy (multiple views requested)
	Attribute      string  `url:"attribute,omitempty"`      //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, energyConsumed, estimatedCost, estimatedEmission, carbonIntensity, numberOfDevices] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=energyConsumed (multiple attributes requested)
	DeviceCategory string  `url:"deviceCategory,omitempty"` //The list of device categories. Note that this filter specifies which devices will be included when calculating energy consumption values, rather than specifying the list of returned sites. Examples: `deviceCategory=Switch` (single device category requested) `deviceCategory=Switch&deviceCategory=Router` (multiple device categories with comma separator)
	TaskID         string  `url:"taskId,omitempty"`         //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type GetSiteEnergyByIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadListOfSiteHealthSummariesQueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return
	Offset          float64 `url:"offset,omitempty"`          //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy          string  `url:"sortBy,omitempty"`          //A field within the response to sort by.
	Order           string  `url:"order,omitempty"`           //The sort order of the field ascending or descending.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	View            string  `url:"view,omitempty"`            //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view: 1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude] 2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage] 3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage] 4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] When this query parameter is not added the default summaries are: **[site,client,network,issue]** Examples: view=client (single view requested) view=client&view=network&view=issue (multiple views requested)
	Attribute       string  `url:"attribute,omitempty"`       //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
}
type ReadListOfSiteHealthSummariesHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadSiteCountQueryParams struct {
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
}
type ReadSiteCountHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadAnAggregatedSummaryOfSiteHealthDataQueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	View            string  `url:"view,omitempty"`            //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view: 1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude] 2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage] 3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage] 4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] When this query parameter is not added the default summaries are: **[site,client,network,issue]** Examples: view=client (single view requested) view=client&view=network&view=issue (multiple views requested)
	Attribute       string  `url:"attribute,omitempty"`       //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
}
type ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
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

type QueryAnAggregatedSummaryOfSiteHealthDataQueryParams struct {
	SiteHierarchy   string `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
}
type ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	ID              string  `url:"id,omitempty"`              //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	TrendInterval   string  `url:"trendInterval,omitempty"`   //The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return
	Offset          float64 `url:"offset,omitempty"`          //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	TimeSortOrder   string  `url:"timeSortOrder,omitempty"`   //The sort order of a time sorted API response.
	Attribute       string  `url:"attribute,omitempty"`       //Supported Analytics Attributes: [networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] attribute=networkDeviceCount (single attribute requested) attribute=networkDeviceCount&attribute=clientCount (multiple attributes requested)
	TaskID          string  `url:"taskId,omitempty"`          //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadSiteHealthSummaryDataBySiteIDQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites. ### Response data proviced by each view: 1. **site** [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude] 2. **network** [id, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage] 3. **client** [id, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage] 4. **issue** [id, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] When this query parameter is not added the default summaries are: **[site,client,network,issue]** Examples: view=client (single view requested) view=client&view=network&view=issue (multiple views requested)
	Attribute string  `url:"attribute,omitempty"` //Supported Attributes: [id, siteHierarchy, siteHierarchyId, siteType, latitude, longitude, networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] If length of attribute list is too long, please use 'view' param instead. Examples: attribute=siteHierarchy (single attribute requested) attribute=siteHierarchy&attribute=clientCount (multiple attributes requested)
}
type ReadSiteHealthSummaryDataBySiteIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TrendInterval string  `url:"trendInterval,omitempty"` //The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	TimeSortOrder string  `url:"timeSortOrder,omitempty"` //The sort order of a time sorted API response.
	Attribute     string  `url:"attribute,omitempty"`     //Supported Analytics Attributes: [networkDeviceCount, networkDeviceGoodHealthCount,wirelessDeviceCount, wirelessDeviceGoodHealthCount, accessDeviceCount, accessDeviceGoodHealthCount, coreDeviceCount, coreDeviceGoodHealthCount, distributionDeviceCount, distributionDeviceGoodHealthCount, routerDeviceCount, routerDeviceGoodHealthCount, apDeviceCount, apDeviceGoodHealthCount, wlcDeviceCount, wlcDeviceGoodHealthCount, switchDeviceCount, switchDeviceGoodHealthCount, networkDeviceGoodHealthPercentage, accessDeviceGoodHealthPercentage, coreDeviceGoodHealthPercentage, distributionDeviceGoodHealthPercentage, routerDeviceGoodHealthPercentage, apDeviceGoodHealthPercentage, wlcDeviceGoodHealthPercentage, switchDeviceGoodHealthPercentage, wirelessDeviceGoodHealthPercentage, clientCount, clientGoodHealthCount, wiredClientCount, wirelessClientCount, wiredClientGoodHealthCount, wirelessClientGoodHealthCount, clientGoodHealthPercentage, wiredClientGoodHealthPercentage, wirelessClientGoodHealthPercentage, clientDataUsage, p1IssueCount, p2IssueCount, p3IssueCount, p4IssueCount, issueCount] attribute=networkDeviceCount (single attribute requested) attribute=networkDeviceCount&attribute=clientCount (multiple attributes requested)
	TaskID        string  `url:"taskId,omitempty"`        //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams struct {
	TaskID          string  `url:"taskId,omitempty"`          //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID          string  `url:"siteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
	SSID            string  `url:"ssid,omitempty"`            //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	Band            string  `url:"band,omitempty"`            //WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz - GHz Examples: `band=5` (single band requested) `band=2.4&band=6` (multiple band requested)
	FailureCategory string  `url:"failureCategory,omitempty"` //Category of failure when a client fails to meet the threshold. Examples: `failureCategory=AUTH` (single failure category requested) `failureCategory=AUTH&failureCategory=DHCP` (multiple failure categories requested)
	FailureReason   string  `url:"failureReason,omitempty"`   //Reason for failure when a client fails to meet the threshold. Examples: `failureReason=MOBILITY_FAILURE` (single ssid requested) `failureReason=REASON_IPLEARN_CONNECT_TIMEOUT&failureReason=ST_EAP_TIMEOUT`   (multiple ssid requested)
	View            string  `url:"view,omitempty"`            //<p>The name of the View. Each view represents a specific data set. Please refer to the <code>SiteAnalyticsView</code> Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned.</p><table><thead><tr><th>View Name</th><th>Included Attributes</th></tr></thead><tbody><tr><td><code>coverage</code></td><td>coverageAverage, coverageSuccessPercentage, coverageSuccessCount, coverageTotalCount, coverageFailureCount, coverageClientCount, coverageImpactedEntities, coverageFailureImpactedEntities, coverageFailureMetrics</td></tr><tr><td><code>onboardingAttempts</code></td><td>onboardingAttemptsSuccessPercentage, onboardingAttemptsSuccessCount, onboardingAttemptsTotalCount, onboardingAttemptsFailureCount, onboardingAttemptsClientCount, onboardingAttemptsImpactedEntities, onboardingAttemptsFailureImpactedEntities, onboardingAttemptsFailureMetrics</td></tr><tr><td><code>onboardingDuration</code></td><td>onboardingDurationAverage, onboardingDurationSuccessPercentage, onboardingDurationSuccessCount, onboardingDurationTotalCount, onboardingDurationFailureCount, onboardingDurationClientCount, onboardingDurationImpactedEntities, onboardingDurationFailureImpactedEntities, onboardingDurationFailureMetrics</td></tr><tr><td><code>roamingAttempts</code></td><td>roamingAttemptsSuccessPercentage, roamingAttemptsSuccessCount, roamingAttemptsTotalCount, roamingAttemptsFailureCount, roamingAttemptsClientCount, roamingAttemptsImpactedEntities, roamingAttemptsFailureImpactedEntities, roamingAttemptsFailureMetrics</td></tr><tr><td><code>roamingDuration</code></td><td>roamingDurationAverage, roamingDurationSuccessPercentage, roamingDurationSuccessCount, roamingDurationTotalCount, roamingDurationFailureCount, roamingDurationClientCount, roamingDurationImpactedEntities, roamingDurationFailureImpactedEntities, roamingDurationFailureMetrics</td></tr><tr><td><code>connectionSpeed</code></td><td>connectionSpeedAverage, connectionSpeedSuccessPercentage, connectionSpeedSuccessCount, connectionSpeedTotalCount, connectionSpeedFailureCount, connectionSpeedClientCount, connectionSpeedImpactedEntities, connectionSpeedFailureImpactedEntities, connectionSpeedFailureMetrics</td></tr></tbody></table><p>Examples: <code>view=connectionSpeed</code> (single view requested) <code>view=roamingDuration&amp;view=roamingAttempts</code> (multiple views requested)       </p>
	Attribute       string  `url:"attribute,omitempty"`       //List of attributes related to site analytics. If these are provided, then only those attributes will be part of response along with the default attributes. Examples: `attribute=coverageAverage` (single attribute requested) `attribute=coverageFailureMetrics&attribute=coverageTotalCount` (multiple attributes requested)
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return
	Offset          float64 `url:"offset,omitempty"`          //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy          string  `url:"sortBy,omitempty"`          //Field name on which sorting needs to be done
	Order           string  `url:"order,omitempty"`           //The sort order of the field ascending or descending.
}
type GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersQueryParams struct {
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy   string  `url:"siteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID string  `url:"siteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID          string  `url:"siteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	SiteType        string  `url:"siteType,omitempty"`        //The type of the site. A site can be an area, building, or floor. Default when not provided will be `[floor,building,area]` Examples: `?siteType=area` (single siteType requested) `?siteType=area&siteType=building&siteType=floor` (multiple siteTypes requested)
}
type GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSiteAnalyticsSummaryDataForTheGivenTaskIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type GetSiteAnalyticsSummaryDataForTheGivenTaskIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSiteAnalyticsTrendDataForTheGivenTaskIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
}
type GetSiteAnalyticsTrendDataForTheGivenTaskIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSiteAnalyticsForOneSiteQueryParams struct {
	TaskID          string  `url:"taskId,omitempty"`          //used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
	StartTime       float64 `url:"startTime,omitempty"`       //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime         float64 `url:"endTime,omitempty"`         //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SSID            string  `url:"ssid,omitempty"`            //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	Band            string  `url:"band,omitempty"`            //WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz - GHz Examples: `band=5` (single band requested) `band=2.4&band=6` (multiple band requested)
	FailureCategory string  `url:"failureCategory,omitempty"` //Category of failure when a client fails to meet the threshold. Examples: `failureCategory=AUTH` (single failure category requested) `failureCategory=AUTH&failureCategory=DHCP` (multiple failure categories requested)
	FailureReason   string  `url:"failureReason,omitempty"`   //Reason for failure when a client fails to meet the threshold. Examples: `failureReason=MOBILITY_FAILURE` (single ssid requested) `failureReason=REASON_IPLEARN_CONNECT_TIMEOUT&failureReason=ST_EAP_TIMEOUT`   (multiple ssid requested)
	View            string  `url:"view,omitempty"`            //<p>The name of the View. Each view represents a specific data set. Please refer to the <code>SiteAnalyticsView</code> Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned.</p><table><thead><tr><th>View Name</th><th>Included Attributes</th></tr></thead><tbody><tr><td><code>coverage</code></td><td>coverageAverage, coverageSuccessPercentage, coverageSuccessCount, coverageTotalCount, coverageFailureCount, coverageClientCount, coverageImpactedEntities, coverageFailureImpactedEntities, coverageFailureMetrics</td></tr><tr><td><code>onboardingAttempts</code></td><td>onboardingAttemptsSuccessPercentage, onboardingAttemptsSuccessCount, onboardingAttemptsTotalCount, onboardingAttemptsFailureCount, onboardingAttemptsClientCount, onboardingAttemptsImpactedEntities, onboardingAttemptsFailureImpactedEntities, onboardingAttemptsFailureMetrics</td></tr><tr><td><code>onboardingDuration</code></td><td>onboardingDurationAverage, onboardingDurationSuccessPercentage, onboardingDurationSuccessCount, onboardingDurationTotalCount, onboardingDurationFailureCount, onboardingDurationClientCount, onboardingDurationImpactedEntities, onboardingDurationFailureImpactedEntities, onboardingDurationFailureMetrics</td></tr><tr><td><code>roamingAttempts</code></td><td>roamingAttemptsSuccessPercentage, roamingAttemptsSuccessCount, roamingAttemptsTotalCount, roamingAttemptsFailureCount, roamingAttemptsClientCount, roamingAttemptsImpactedEntities, roamingAttemptsFailureImpactedEntities, roamingAttemptsFailureMetrics</td></tr><tr><td><code>roamingDuration</code></td><td>roamingDurationAverage, roamingDurationSuccessPercentage, roamingDurationSuccessCount, roamingDurationTotalCount, roamingDurationFailureCount, roamingDurationClientCount, roamingDurationImpactedEntities, roamingDurationFailureImpactedEntities, roamingDurationFailureMetrics</td></tr><tr><td><code>connectionSpeed</code></td><td>connectionSpeedAverage, connectionSpeedSuccessPercentage, connectionSpeedSuccessCount, connectionSpeedTotalCount, connectionSpeedFailureCount, connectionSpeedClientCount, connectionSpeedImpactedEntities, connectionSpeedFailureImpactedEntities, connectionSpeedFailureMetrics</td></tr></tbody></table><p>Examples: <code>view=connectionSpeed</code> (single view requested) <code>view=roamingDuration&amp;view=roamingAttempts</code> (multiple views requested)       </p>
	Attribute       string  `url:"attribute,omitempty"`       //List of attributes related to site analytics. If these are provided, then only those attributes will be part of response along with the default attributes. Examples: `attribute=coverageAverage` (single attribute requested) `attribute=coverageFailureMetrics&attribute=coverageTotalCount` (multiple attributes requested)
}
type GetSiteAnalyticsForOneSiteHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type AssignDevicesToSiteHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetMembershipQueryParams struct {
	Offset       float64 `url:"offset,omitempty"`       //offset/starting row
	Limit        float64 `url:"limit,omitempty"`        //Number of sites to be retrieved
	DeviceFamily string  `url:"deviceFamily,omitempty"` //Device family name
	SerialNumber string  `url:"serialNumber,omitempty"` //Device serial number
}
type CreateSiteHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetSiteQueryParams struct {
	Name   string `url:"name,omitempty"`   //Site name hierarchy (E.g Global/USA/CA)
	SiteID string `url:"siteId,omitempty"` //Site Id
	Type   string `url:"type,omitempty"`   //Site type (Ex: area, building, floor)
	Offset int    `url:"offset,omitempty"` //Offset/starting index for pagination. Indexed from 1.
	Limit  int    `url:"limit,omitempty"`  //Number of sites to be listed
}
type GetSiteHealthQueryParams struct {
	SiteType  string  `url:"siteType,omitempty"`  //site type: AREA or BUILDING (case insensitive)
	Offset    float64 `url:"offset,omitempty"`    //Offset of the first returned data set entry (Multiple of 'limit' + 1)
	Limit     float64 `url:"limit,omitempty"`     //Max number of data entries in the returned data set [1,50].  Default is 25
	Timestamp float64 `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Site Hierarchy data is required
}
type GetDevicesThatAreAssignedToASiteQueryParams struct {
	Offset     string `url:"offset,omitempty"`     //Offset/starting index for pagination
	Limit      string `url:"limit,omitempty"`      //Number of devices to be listed. Default and max supported value is 500
	MemberType string `url:"memberType,omitempty"` //Member type (This API only supports the 'networkdevice' type)
	Level      string `url:"level,omitempty"`      //Depth of site hierarchy to be considered to list the devices. If the provided value is -1, devices for all child sites will be listed.
}
type GetSiteCountQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site instance UUID
}
type UpdateSiteHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetSiteV2QueryParams struct {
	GroupNameHierarchy string `url:"groupNameHierarchy,omitempty"` //Site name hierarchy (E.g. Global/USA/CA)
	ID                 string `url:"id,omitempty"`                 //Site Id
	Type               string `url:"type,omitempty"`               //Site type (Acceptable values: area, building, floor)
	Offset             string `url:"offset,omitempty"`             //Offset/starting index for pagination
	Limit              string `url:"limit,omitempty"`              //Number of sites to be listed. Default and max supported value is 500
}
type GetSiteCountV2QueryParams struct {
	ID string `url:"id,omitempty"` //Site instance UUID
}

type ResponseSitesGetSitesEnergy struct {
	Response *[]ResponseSitesGetSitesEnergyResponse `json:"response,omitempty"` //
	Page     *ResponseSitesGetSitesEnergyPage       `json:"page,omitempty"`     //
	Version  string                                 `json:"version,omitempty"`  // Version
}
type ResponseSitesGetSitesEnergyResponse struct {
	ID                string   `json:"id,omitempty"`                // Id
	SiteName          string   `json:"siteName,omitempty"`          // Site Name
	SiteHierarchy     string   `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteHierarchyID   string   `json:"siteHierarchyId,omitempty"`   // Site Hierarchy Id
	SiteType          string   `json:"siteType,omitempty"`          // Site Type
	Latitude          *float64 `json:"latitude,omitempty"`          // Latitude
	Longitude         *float64 `json:"longitude,omitempty"`         // Longitude
	DeviceCategories  []string `json:"deviceCategories,omitempty"`  // Device Categories
	EnergyConsumed    *float64 `json:"energyConsumed,omitempty"`    // Energy Consumed
	EstimatedCost     *float64 `json:"estimatedCost,omitempty"`     // Estimated Cost
	EstimatedEmission *float64 `json:"estimatedEmission,omitempty"` // Estimated Emission
	CarbonIntensity   *float64 `json:"carbonIntensity,omitempty"`   // Carbon Intensity
	NumberOfDevices   *int     `json:"numberOfDevices,omitempty"`   // Number Of Devices
}
type ResponseSitesGetSitesEnergyPage struct {
	Limit  *int                                     `json:"limit,omitempty"`  // Limit
	Offset *int                                     `json:"offset,omitempty"` // Offset
	Count  *int                                     `json:"count,omitempty"`  // Count
	SortBy *[]ResponseSitesGetSitesEnergyPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseSitesGetSitesEnergyPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseSitesCountSitesEnergy struct {
	Response *ResponseSitesCountSitesEnergyResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version
}
type ResponseSitesCountSitesEnergyResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesSubmitRequestToQuerySitesEnergy struct {
	Response *ResponseSitesSubmitRequestToQuerySitesEnergyResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseSitesSubmitRequestToQuerySitesEnergyResponse struct {
	TaskLocation string `json:"taskLocation,omitempty"` // Task Location
	TaskID       string `json:"taskId,omitempty"`       // Task Id
}
type ResponseSitesQuerySitesEnergyForTheGivenTaskID struct {
	Response *[]ResponseSitesQuerySitesEnergyForTheGivenTaskIDResponse `json:"response,omitempty"` //
	Page     *ResponseSitesQuerySitesEnergyForTheGivenTaskIDPage       `json:"page,omitempty"`     //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseSitesQuerySitesEnergyForTheGivenTaskIDResponse struct {
	ID                  string                                                                       `json:"id,omitempty"`                  // Id
	SiteName            string                                                                       `json:"siteName,omitempty"`            // Site Name
	SiteHierarchy       string                                                                       `json:"siteHierarchy,omitempty"`       // Site Hierarchy
	SiteHierarchyID     string                                                                       `json:"siteHierarchyId,omitempty"`     // Site Hierarchy Id
	SiteType            string                                                                       `json:"siteType,omitempty"`            // Site Type
	Latitude            *float64                                                                     `json:"latitude,omitempty"`            // Latitude
	Longitude           *float64                                                                     `json:"longitude,omitempty"`           // Longitude
	DeviceCategories    []string                                                                     `json:"deviceCategories,omitempty"`    // Device Categories
	EnergyConsumed      *float64                                                                     `json:"energyConsumed,omitempty"`      // Energy Consumed
	EstimatedCost       *float64                                                                     `json:"estimatedCost,omitempty"`       // Estimated Cost
	EstimatedEmission   *float64                                                                     `json:"estimatedEmission,omitempty"`   // Estimated Emission
	CarbonIntensity     *float64                                                                     `json:"carbonIntensity,omitempty"`     // Carbon Intensity
	NumberOfDevices     *int                                                                         `json:"numberOfDevices,omitempty"`     // Number Of Devices
	AggregateAttributes *[]ResponseSitesQuerySitesEnergyForTheGivenTaskIDResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseSitesQuerySitesEnergyForTheGivenTaskIDResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseSitesQuerySitesEnergyForTheGivenTaskIDPage struct {
	Limit  *int                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                        `json:"offset,omitempty"` // Offset
	Count  *int                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseSitesQuerySitesEnergyForTheGivenTaskIDPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseSitesQuerySitesEnergyForTheGivenTaskIDPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type ResponseSitesCountSitesEnergyForTheGivenTaskID struct {
	Response *ResponseSitesCountSitesEnergyForTheGivenTaskIDResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseSitesCountSitesEnergyForTheGivenTaskIDResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesSubmitRequestToCountSitesEnergyFromQuery struct {
	Response *ResponseSitesSubmitRequestToCountSitesEnergyFromQueryResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseSitesSubmitRequestToCountSitesEnergyFromQueryResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesGetSiteEnergyByID struct {
	Response *ResponseSitesGetSiteEnergyByIDResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  // Version
}
type ResponseSitesGetSiteEnergyByIDResponse struct {
	ID                string   `json:"id,omitempty"`                // Id
	SiteName          string   `json:"siteName,omitempty"`          // Site Name
	SiteHierarchy     string   `json:"siteHierarchy,omitempty"`     // Site Hierarchy
	SiteHierarchyID   string   `json:"siteHierarchyId,omitempty"`   // Site Hierarchy Id
	SiteType          string   `json:"siteType,omitempty"`          // Site Type
	Latitude          *float64 `json:"latitude,omitempty"`          // Latitude
	Longitude         *float64 `json:"longitude,omitempty"`         // Longitude
	DeviceCategories  []string `json:"deviceCategories,omitempty"`  // Device Categories
	EnergyConsumed    *float64 `json:"energyConsumed,omitempty"`    // Energy Consumed
	EstimatedCost     *float64 `json:"estimatedCost,omitempty"`     // Estimated Cost
	EstimatedEmission *float64 `json:"estimatedEmission,omitempty"` // Estimated Emission
	CarbonIntensity   *float64 `json:"carbonIntensity,omitempty"`   // Carbon Intensity
	NumberOfDevices   *int     `json:"numberOfDevices,omitempty"`   // Number Of Devices
}
type ResponseSitesReadListOfSiteHealthSummaries struct {
	Response *[]ResponseSitesReadListOfSiteHealthSummariesResponse `json:"response,omitempty"` //
	Page     *ResponseSitesReadListOfSiteHealthSummariesPage       `json:"page,omitempty"`     //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseSitesReadListOfSiteHealthSummariesResponse struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesReadListOfSiteHealthSummariesPage struct {
	Limit  *int                                                    `json:"limit,omitempty"`  // Limit
	Offset *int                                                    `json:"offset,omitempty"` // Offset
	Count  *int                                                    `json:"count,omitempty"`  // Count
	SortBy *[]ResponseSitesReadListOfSiteHealthSummariesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseSitesReadListOfSiteHealthSummariesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseSitesReadSiteCount struct {
	Response *ResponseSitesReadSiteCountResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version
}
type ResponseSitesReadSiteCountResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesReadAnAggregatedSummaryOfSiteHealthData struct {
	Response *ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // Version
}
type ResponseSitesReadAnAggregatedSummaryOfSiteHealthDataResponse struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesQueryAnAggregatedSummaryOfSiteHealthData struct {
	Response *ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseSitesQueryAnAggregatedSummaryOfSiteHealthDataResponse struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork struct {
	Response *[]ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkResponse `json:"response,omitempty"` //
	Page     *ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkPage       `json:"page,omitempty"`     //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkResponse struct {
	Timestamp  *int                                                                                     `json:"timestamp,omitempty"`  // Timestamp
	Attributes *[]ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkResponseAttributes struct {
	Name  string   `json:"name,omitempty"`  // Name
	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Offset        *int   `json:"offset,omitempty"`        // Offset
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSitesReadSiteHealthSummaryDataBySiteID struct {
	Response *ResponseSitesReadSiteHealthSummaryDataBySiteIDResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseSitesReadSiteHealthSummaryDataBySiteIDResponse struct {
	ID                                     string   `json:"id,omitempty"`                                     // Id
	SiteHierarchy                          string   `json:"siteHierarchy,omitempty"`                          // Site Hierarchy
	SiteHierarchyID                        string   `json:"siteHierarchyId,omitempty"`                        // Site Hierarchy Id
	SiteType                               string   `json:"siteType,omitempty"`                               // Site Type
	Latitude                               *float64 `json:"latitude,omitempty"`                               // Latitude
	Longitude                              *float64 `json:"longitude,omitempty"`                              // Longitude
	NetworkDeviceGoodHealthPercentage      *int     `json:"networkDeviceGoodHealthPercentage,omitempty"`      // Network Device Good Health Percentage
	NetworkDeviceGoodHealthCount           *int     `json:"networkDeviceGoodHealthCount,omitempty"`           // Network Device Good Health Count
	ClientGoodHealthCount                  *int     `json:"clientGoodHealthCount,omitempty"`                  // Client Good Health Count
	ClientGoodHealthPercentage             *int     `json:"clientGoodHealthPercentage,omitempty"`             // Client Good Health Percentage
	WiredClientGoodHealthPercentage        *int     `json:"wiredClientGoodHealthPercentage,omitempty"`        // Wired Client Good Health Percentage
	WirelessClientGoodHealthPercentage     *int     `json:"wirelessClientGoodHealthPercentage,omitempty"`     // Wireless Client Good Health Percentage
	ClientCount                            *int     `json:"clientCount,omitempty"`                            // Client Count
	WiredClientCount                       *int     `json:"wiredClientCount,omitempty"`                       // Wired Client Count
	WirelessClientCount                    *int     `json:"wirelessClientCount,omitempty"`                    // Wireless Client Count
	WiredClientGoodHealthCount             *int     `json:"wiredClientGoodHealthCount,omitempty"`             // Wired Client Good Health Count
	WirelessClientGoodHealthCount          *int     `json:"wirelessClientGoodHealthCount,omitempty"`          // Wireless Client Good Health Count
	NetworkDeviceCount                     *int     `json:"networkDeviceCount,omitempty"`                     // Network Device Count
	AccessDeviceCount                      *int     `json:"accessDeviceCount,omitempty"`                      // Access Device Count
	AccessDeviceGoodHealthCount            *int     `json:"accessDeviceGoodHealthCount,omitempty"`            // Access Device Good Health Count
	CoreDeviceCount                        *int     `json:"coreDeviceCount,omitempty"`                        // Core Device Count
	CoreDeviceGoodHealthCount              *int     `json:"coreDeviceGoodHealthCount,omitempty"`              // Core Device Good Health Count
	DistributionDeviceCount                *int     `json:"distributionDeviceCount,omitempty"`                // Distribution Device Count
	DistributionDeviceGoodHealthCount      *int     `json:"distributionDeviceGoodHealthCount,omitempty"`      // Distribution Device Good Health Count
	RouterDeviceCount                      *int     `json:"routerDeviceCount,omitempty"`                      // Router Device Count
	RouterDeviceGoodHealthCount            *int     `json:"routerDeviceGoodHealthCount,omitempty"`            // Router Device Good Health Count
	WirelessDeviceCount                    *int     `json:"wirelessDeviceCount,omitempty"`                    // Wireless Device Count
	WirelessDeviceGoodHealthCount          *int     `json:"wirelessDeviceGoodHealthCount,omitempty"`          // Wireless Device Good Health Count
	ApDeviceCount                          *int     `json:"apDeviceCount,omitempty"`                          // Ap Device Count
	ApDeviceGoodHealthCount                *int     `json:"apDeviceGoodHealthCount,omitempty"`                // Ap Device Good Health Count
	WlcDeviceCount                         *int     `json:"wlcDeviceCount,omitempty"`                         // Wlc Device Count
	WlcDeviceGoodHealthCount               *int     `json:"wlcDeviceGoodHealthCount,omitempty"`               // Wlc Device Good Health Count
	SwitchDeviceCount                      *int     `json:"switchDeviceCount,omitempty"`                      // Switch Device Count
	SwitchDeviceGoodHealthCount            *int     `json:"switchDeviceGoodHealthCount,omitempty"`            // Switch Device Good Health Count
	AccessDeviceGoodHealthPercentage       *int     `json:"accessDeviceGoodHealthPercentage,omitempty"`       // Access Device Good Health Percentage
	CoreDeviceGoodHealthPercentage         *int     `json:"coreDeviceGoodHealthPercentage,omitempty"`         // Core Device Good Health Percentage
	DistributionDeviceGoodHealthPercentage *int     `json:"distributionDeviceGoodHealthPercentage,omitempty"` // Distribution Device Good Health Percentage
	RouterDeviceGoodHealthPercentage       *int     `json:"routerDeviceGoodHealthPercentage,omitempty"`       // Router Device Good Health Percentage
	ApDeviceGoodHealthPercentage           *int     `json:"apDeviceGoodHealthPercentage,omitempty"`           // Ap Device Good Health Percentage
	WlcDeviceGoodHealthPercentage          *int     `json:"wlcDeviceGoodHealthPercentage,omitempty"`          // Wlc Device Good Health Percentage
	SwitchDeviceGoodHealthPercentage       *int     `json:"switchDeviceGoodHealthPercentage,omitempty"`       // Switch Device Good Health Percentage
	WirelessDeviceGoodHealthPercentage     *int     `json:"wirelessDeviceGoodHealthPercentage,omitempty"`     // Wireless Device Good Health Percentage
	ClientDataUsage                        *float64 `json:"clientDataUsage,omitempty"`                        // Client Data Usage
	P1IssueCount                           *int     `json:"p1IssueCount,omitempty"`                           // P1 Issue Count
	P2IssueCount                           *int     `json:"p2IssueCount,omitempty"`                           // P2 Issue Count
	P3IssueCount                           *int     `json:"p3IssueCount,omitempty"`                           // P3 Issue Count
	P4IssueCount                           *int     `json:"p4IssueCount,omitempty"`                           // P4 Issue Count
	IssueCount                             *int     `json:"issueCount,omitempty"`                             // Issue Count
}
type ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetwork struct {
	Response *[]ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkResponse `json:"response,omitempty"` //
	Page     *ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkPage       `json:"page,omitempty"`     //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkResponse struct {
	Timestamp  *int                                                                                  `json:"timestamp,omitempty"`  // Timestamp
	Attributes *[]ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkResponseAttributes struct {
	Name  string   `json:"name,omitempty"`  // Name
	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetworkPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Offset        *int   `json:"offset,omitempty"`        // Offset
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters struct {
	Response *[]ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponse `json:"response,omitempty"` //
	Page     *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersPage       `json:"page,omitempty"`     //
	Version  string                                                                                           `json:"version,omitempty"`  // Version
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponse struct {
	ID                                        string                                                                                                                                  `json:"id,omitempty"`                                        // Id
	SiteID                                    string                                                                                                                                  `json:"siteId,omitempty"`                                    // Site Id
	SiteHierarchyID                           string                                                                                                                                  `json:"siteHierarchyId,omitempty"`                           // Site Hierarchy Id
	SiteHierarchy                             string                                                                                                                                  `json:"siteHierarchy,omitempty"`                             // Site Hierarchy
	SiteType                                  string                                                                                                                                  `json:"siteType,omitempty"`                                  // Site Type
	ApCount                                   *int                                                                                                                                    `json:"apCount,omitempty"`                                   // Ap Count
	CoverageAverage                           *int                                                                                                                                    `json:"coverageAverage,omitempty"`                           // Coverage Average
	CoverageSuccessPercentage                 *int                                                                                                                                    `json:"coverageSuccessPercentage,omitempty"`                 // Coverage Success Percentage
	CoverageSuccessCount                      *int                                                                                                                                    `json:"coverageSuccessCount,omitempty"`                      // Coverage Success Count
	CoverageTotalCount                        *int                                                                                                                                    `json:"coverageTotalCount,omitempty"`                        // Coverage Total Count
	CoverageFailureCount                      *int                                                                                                                                    `json:"coverageFailureCount,omitempty"`                      // Coverage Failure Count
	CoverageClientCount                       *int                                                                                                                                    `json:"coverageClientCount,omitempty"`                       // Coverage Client Count
	CoverageImpactedEntities                  *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageImpactedEntities                  `json:"coverageImpactedEntities,omitempty"`                  //
	CoverageFailureImpactedEntities           *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageFailureImpactedEntities           `json:"coverageFailureImpactedEntities,omitempty"`           //
	CoverageFailureMetrics                    *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageFailureMetrics                    `json:"coverageFailureMetrics,omitempty"`                    //
	OnboardingAttemptsSuccessPercentage       *int                                                                                                                                    `json:"onboardingAttemptsSuccessPercentage,omitempty"`       // Onboarding Attempts Success Percentage
	OnboardingAttemptsSuccessCount            *int                                                                                                                                    `json:"onboardingAttemptsSuccessCount,omitempty"`            // Onboarding Attempts Success Count
	OnboardingAttemptsTotalCount              *int                                                                                                                                    `json:"onboardingAttemptsTotalCount,omitempty"`              // Onboarding Attempts Total Count
	OnboardingAttemptsFailureCount            *int                                                                                                                                    `json:"onboardingAttemptsFailureCount,omitempty"`            // Onboarding Attempts Failure Count
	OnboardingAttemptsClientCount             *int                                                                                                                                    `json:"onboardingAttemptsClientCount,omitempty"`             // Onboarding Attempts Client Count
	OnboardingAttemptsImpactedEntities        *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsImpactedEntities        `json:"onboardingAttemptsImpactedEntities,omitempty"`        //
	OnboardingAttemptsFailureImpactedEntities *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsFailureImpactedEntities `json:"onboardingAttemptsFailureImpactedEntities,omitempty"` //
	OnboardingAttemptsFailureMetrics          *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsFailureMetrics          `json:"onboardingAttemptsFailureMetrics,omitempty"`          //
	OnboardingDurationAverage                 *int                                                                                                                                    `json:"onboardingDurationAverage,omitempty"`                 // Onboarding Duration Average
	OnboardingDurationSuccessPercentage       *int                                                                                                                                    `json:"onboardingDurationSuccessPercentage,omitempty"`       // Onboarding Duration Success Percentage
	OnboardingDurationSuccessCount            *int                                                                                                                                    `json:"onboardingDurationSuccessCount,omitempty"`            // Onboarding Duration Success Count
	OnboardingDurationTotalCount              *int                                                                                                                                    `json:"onboardingDurationTotalCount,omitempty"`              // Onboarding Duration Total Count
	OnboardingDurationFailureCount            *int                                                                                                                                    `json:"onboardingDurationFailureCount,omitempty"`            // Onboarding Duration Failure Count
	OnboardingDurationClientCount             *int                                                                                                                                    `json:"onboardingDurationClientCount,omitempty"`             // Onboarding Duration Client Count
	OnboardingDurationImpactedEntities        *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationImpactedEntities        `json:"onboardingDurationImpactedEntities,omitempty"`        //
	OnboardingDurationFailureImpactedEntities *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationFailureImpactedEntities `json:"onboardingDurationFailureImpactedEntities,omitempty"` //
	OnboardingDurationFailureMetrics          *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationFailureMetrics          `json:"onboardingDurationFailureMetrics,omitempty"`          //
	RoamingAttemptsSuccessPercentage          *int                                                                                                                                    `json:"roamingAttemptsSuccessPercentage,omitempty"`          // Roaming Attempts Success Percentage
	RoamingAttemptsSuccessCount               *int                                                                                                                                    `json:"roamingAttemptsSuccessCount,omitempty"`               // Roaming Attempts Success Count
	RoamingAttemptsTotalCount                 *int                                                                                                                                    `json:"roamingAttemptsTotalCount,omitempty"`                 // Roaming Attempts Total Count
	RoamingAttemptsFailureCount               *int                                                                                                                                    `json:"roamingAttemptsFailureCount,omitempty"`               // Roaming Attempts Failure Count
	RoamingAttemptsClientCount                *int                                                                                                                                    `json:"roamingAttemptsClientCount,omitempty"`                // Roaming Attempts Client Count
	RoamingAttemptsImpactedEntities           *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsImpactedEntities           `json:"roamingAttemptsImpactedEntities,omitempty"`           //
	RoamingAttemptsFailureImpactedEntities    *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsFailureImpactedEntities    `json:"roamingAttemptsFailureImpactedEntities,omitempty"`    //
	RoamingAttemptsFailureMetrics             *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsFailureMetrics             `json:"roamingAttemptsFailureMetrics,omitempty"`             //
	RoamingDurationAverage                    *int                                                                                                                                    `json:"roamingDurationAverage,omitempty"`                    // Roaming Duration Average
	RoamingDurationSuccessPercentage          *int                                                                                                                                    `json:"roamingDurationSuccessPercentage,omitempty"`          // Roaming Duration Success Percentage
	RoamingDurationSuccessCount               *int                                                                                                                                    `json:"roamingDurationSuccessCount,omitempty"`               // Roaming Duration Success Count
	RoamingDurationTotalCount                 *int                                                                                                                                    `json:"roamingDurationTotalCount,omitempty"`                 // Roaming Duration Total Count
	RoamingDurationFailureCount               *int                                                                                                                                    `json:"roamingDurationFailureCount,omitempty"`               // Roaming Duration Failure Count
	RoamingDurationClientCount                *int                                                                                                                                    `json:"roamingDurationClientCount,omitempty"`                // Roaming Duration Client Count
	RoamingDurationImpactedEntities           *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationImpactedEntities           `json:"roamingDurationImpactedEntities,omitempty"`           //
	RoamingDurationFailureImpactedEntities    *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationFailureImpactedEntities    `json:"roamingDurationFailureImpactedEntities,omitempty"`    //
	RoamingDurationFailureMetrics             *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationFailureMetrics             `json:"roamingDurationFailureMetrics,omitempty"`             //
	ConnectionSpeedAverage                    *int                                                                                                                                    `json:"connectionSpeedAverage,omitempty"`                    // Connection Speed Average
	ConnectionSpeedSuccessPercentage          *int                                                                                                                                    `json:"connectionSpeedSuccessPercentage,omitempty"`          // Connection Speed Success Percentage
	ConnectionSpeedSuccessCount               *int                                                                                                                                    `json:"connectionSpeedSuccessCount,omitempty"`               // Connection Speed Success Count
	ConnectionSpeedTotalCount                 *int                                                                                                                                    `json:"connectionSpeedTotalCount,omitempty"`                 // Connection Speed Total Count
	ConnectionSpeedFailureCount               *int                                                                                                                                    `json:"connectionSpeedFailureCount,omitempty"`               // Connection Speed Failure Count
	ConnectionSpeedClientCount                *int                                                                                                                                    `json:"connectionSpeedClientCount,omitempty"`                // Connection Speed Client Count
	ConnectionSpeedImpactedEntities           *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedImpactedEntities           `json:"connectionSpeedImpactedEntities,omitempty"`           //
	ConnectionSpeedFailureImpactedEntities    *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedFailureImpactedEntities    `json:"connectionSpeedFailureImpactedEntities,omitempty"`    //
	ConnectionSpeedFailureMetrics             *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedFailureMetrics             `json:"connectionSpeedFailureMetrics,omitempty"`             //
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseCoverageFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingAttemptsFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseOnboardingDurationFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingAttemptsFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseRoamingDurationFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersResponseConnectionSpeedFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersPage struct {
	Limit  *int                                                                                             `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                             `json:"offset,omitempty"` // Offset
	Count  *int                                                                                             `json:"count,omitempty"`  // Count
	SortBy *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters struct {
	Response *ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersResponse `json:"response,omitempty"` //
	Version  string                                                                                                `json:"version,omitempty"`  // Version
}
type ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters struct {
	Response *ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersResponse struct {
	TaskLocation string `json:"taskLocation,omitempty"` // Task Location
	TaskID       string `json:"taskId,omitempty"`       // Task Id
}
type ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters struct {
	Response *ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                        `json:"version,omitempty"`  // Version
}
type ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskID struct {
	Response *ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDResponse `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDResponse struct {
	Attributes *[]ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskIDResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseSitesSubmitRequestForSiteAnalyticsSummaryData struct {
	Response *ResponseSitesSubmitRequestForSiteAnalyticsSummaryDataResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseSitesSubmitRequestForSiteAnalyticsSummaryDataResponse struct {
	TaskLocation string `json:"taskLocation,omitempty"` // Task Location
	TaskID       string `json:"taskId,omitempty"`       // Task Id
}
type ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID struct {
	Version  string                                                                         `json:"version,omitempty"`  // Version
	Response *[]ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponse `json:"response,omitempty"` //
	Page     *ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDPage       `json:"page,omitempty"`     //
}
type ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponse struct {
	ID         string                                                                                   `json:"id,omitempty"`         // Id
	Attributes *[]ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDPage struct {
	Limit  *int `json:"limit,omitempty"`  // Limit
	Offset *int `json:"offset,omitempty"` // Offset
	Count  *int `json:"count,omitempty"`  // Count
}
type ResponseSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics struct {
	Response *ResponseSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsResponse `json:"response,omitempty"` //
	Version  string                                                                   `json:"version,omitempty"`  // Version
}
type ResponseSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsResponse struct {
	TaskLocation string `json:"taskLocation,omitempty"` // Task Location
	TaskID       string `json:"taskId,omitempty"`       // Task Id
}
type ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskID struct {
	Version  string                                                             `json:"version,omitempty"`  // Version
	Response *[]ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponse `json:"response,omitempty"` //
	Page     *ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDPage       `json:"page,omitempty"`     //
}
type ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponse struct {
	Timestamp  *int                                                                         `json:"timestamp,omitempty"`  // Timestamp
	Attributes *[]ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskIDPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseSitesSubmitRequestForSiteAnalyticsTrendData struct {
	Response *ResponseSitesSubmitRequestForSiteAnalyticsTrendDataResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseSitesSubmitRequestForSiteAnalyticsTrendDataResponse struct {
	TaskLocation string `json:"taskLocation,omitempty"` // Task Location
	TaskID       string `json:"taskId,omitempty"`       // Task Id
}
type ResponseSitesGetSiteAnalyticsForOneSite struct {
	Response *ResponseSitesGetSiteAnalyticsForOneSiteResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponse struct {
	ID                                        string                                                                                    `json:"id,omitempty"`                                        // Id
	SiteID                                    string                                                                                    `json:"siteId,omitempty"`                                    // Site Id
	SiteHierarchyID                           string                                                                                    `json:"siteHierarchyId,omitempty"`                           // Site Hierarchy Id
	SiteHierarchy                             string                                                                                    `json:"siteHierarchy,omitempty"`                             // Site Hierarchy
	SiteType                                  string                                                                                    `json:"siteType,omitempty"`                                  // Site Type
	ApCount                                   *int                                                                                      `json:"apCount,omitempty"`                                   // Ap Count
	CoverageAverage                           *int                                                                                      `json:"coverageAverage,omitempty"`                           // Coverage Average
	CoverageSuccessPercentage                 *int                                                                                      `json:"coverageSuccessPercentage,omitempty"`                 // Coverage Success Percentage
	CoverageSuccessCount                      *int                                                                                      `json:"coverageSuccessCount,omitempty"`                      // Coverage Success Count
	CoverageTotalCount                        *int                                                                                      `json:"coverageTotalCount,omitempty"`                        // Coverage Total Count
	CoverageFailureCount                      *int                                                                                      `json:"coverageFailureCount,omitempty"`                      // Coverage Failure Count
	CoverageClientCount                       *int                                                                                      `json:"coverageClientCount,omitempty"`                       // Coverage Client Count
	CoverageImpactedEntities                  *ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageImpactedEntities                  `json:"coverageImpactedEntities,omitempty"`                  //
	CoverageFailureImpactedEntities           *ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageFailureImpactedEntities           `json:"coverageFailureImpactedEntities,omitempty"`           //
	CoverageFailureMetrics                    *ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageFailureMetrics                    `json:"coverageFailureMetrics,omitempty"`                    //
	OnboardingAttemptsSuccessPercentage       *int                                                                                      `json:"onboardingAttemptsSuccessPercentage,omitempty"`       // Onboarding Attempts Success Percentage
	OnboardingAttemptsSuccessCount            *int                                                                                      `json:"onboardingAttemptsSuccessCount,omitempty"`            // Onboarding Attempts Success Count
	OnboardingAttemptsTotalCount              *int                                                                                      `json:"onboardingAttemptsTotalCount,omitempty"`              // Onboarding Attempts Total Count
	OnboardingAttemptsFailureCount            *int                                                                                      `json:"onboardingAttemptsFailureCount,omitempty"`            // Onboarding Attempts Failure Count
	OnboardingAttemptsClientCount             *int                                                                                      `json:"onboardingAttemptsClientCount,omitempty"`             // Onboarding Attempts Client Count
	OnboardingAttemptsImpactedEntities        *ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsImpactedEntities        `json:"onboardingAttemptsImpactedEntities,omitempty"`        //
	OnboardingAttemptsFailureImpactedEntities *ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsFailureImpactedEntities `json:"onboardingAttemptsFailureImpactedEntities,omitempty"` //
	OnboardingAttemptsFailureMetrics          *ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsFailureMetrics          `json:"onboardingAttemptsFailureMetrics,omitempty"`          //
	OnboardingDurationAverage                 *int                                                                                      `json:"onboardingDurationAverage,omitempty"`                 // Onboarding Duration Average
	OnboardingDurationSuccessPercentage       *int                                                                                      `json:"onboardingDurationSuccessPercentage,omitempty"`       // Onboarding Duration Success Percentage
	OnboardingDurationSuccessCount            *int                                                                                      `json:"onboardingDurationSuccessCount,omitempty"`            // Onboarding Duration Success Count
	OnboardingDurationTotalCount              *int                                                                                      `json:"onboardingDurationTotalCount,omitempty"`              // Onboarding Duration Total Count
	OnboardingDurationFailureCount            *int                                                                                      `json:"onboardingDurationFailureCount,omitempty"`            // Onboarding Duration Failure Count
	OnboardingDurationClientCount             *int                                                                                      `json:"onboardingDurationClientCount,omitempty"`             // Onboarding Duration Client Count
	OnboardingDurationImpactedEntities        *ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationImpactedEntities        `json:"onboardingDurationImpactedEntities,omitempty"`        //
	OnboardingDurationFailureImpactedEntities *ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationFailureImpactedEntities `json:"onboardingDurationFailureImpactedEntities,omitempty"` //
	OnboardingDurationFailureMetrics          *ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationFailureMetrics          `json:"onboardingDurationFailureMetrics,omitempty"`          //
	RoamingAttemptsSuccessPercentage          *int                                                                                      `json:"roamingAttemptsSuccessPercentage,omitempty"`          // Roaming Attempts Success Percentage
	RoamingAttemptsSuccessCount               *int                                                                                      `json:"roamingAttemptsSuccessCount,omitempty"`               // Roaming Attempts Success Count
	RoamingAttemptsTotalCount                 *int                                                                                      `json:"roamingAttemptsTotalCount,omitempty"`                 // Roaming Attempts Total Count
	RoamingAttemptsFailureCount               *int                                                                                      `json:"roamingAttemptsFailureCount,omitempty"`               // Roaming Attempts Failure Count
	RoamingAttemptsClientCount                *int                                                                                      `json:"roamingAttemptsClientCount,omitempty"`                // Roaming Attempts Client Count
	RoamingAttemptsImpactedEntities           *ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsImpactedEntities           `json:"roamingAttemptsImpactedEntities,omitempty"`           //
	RoamingAttemptsFailureImpactedEntities    *ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsFailureImpactedEntities    `json:"roamingAttemptsFailureImpactedEntities,omitempty"`    //
	RoamingAttemptsFailureMetrics             *ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsFailureMetrics             `json:"roamingAttemptsFailureMetrics,omitempty"`             //
	RoamingDurationAverage                    *int                                                                                      `json:"roamingDurationAverage,omitempty"`                    // Roaming Duration Average
	RoamingDurationSuccessPercentage          *int                                                                                      `json:"roamingDurationSuccessPercentage,omitempty"`          // Roaming Duration Success Percentage
	RoamingDurationSuccessCount               *int                                                                                      `json:"roamingDurationSuccessCount,omitempty"`               // Roaming Duration Success Count
	RoamingDurationTotalCount                 *int                                                                                      `json:"roamingDurationTotalCount,omitempty"`                 // Roaming Duration Total Count
	RoamingDurationFailureCount               *int                                                                                      `json:"roamingDurationFailureCount,omitempty"`               // Roaming Duration Failure Count
	RoamingDurationClientCount                *int                                                                                      `json:"roamingDurationClientCount,omitempty"`                // Roaming Duration Client Count
	RoamingDurationImpactedEntities           *ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationImpactedEntities           `json:"roamingDurationImpactedEntities,omitempty"`           //
	RoamingDurationFailureImpactedEntities    *ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationFailureImpactedEntities    `json:"roamingDurationFailureImpactedEntities,omitempty"`    //
	RoamingDurationFailureMetrics             *ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationFailureMetrics             `json:"roamingDurationFailureMetrics,omitempty"`             //
	ConnectionSpeedAverage                    *int                                                                                      `json:"connectionSpeedAverage,omitempty"`                    // Connection Speed Average
	ConnectionSpeedSuccessPercentage          *int                                                                                      `json:"connectionSpeedSuccessPercentage,omitempty"`          // Connection Speed Success Percentage
	ConnectionSpeedSuccessCount               *int                                                                                      `json:"connectionSpeedSuccessCount,omitempty"`               // Connection Speed Success Count
	ConnectionSpeedTotalCount                 *int                                                                                      `json:"connectionSpeedTotalCount,omitempty"`                 // Connection Speed Total Count
	ConnectionSpeedFailureCount               *int                                                                                      `json:"connectionSpeedFailureCount,omitempty"`               // Connection Speed Failure Count
	ConnectionSpeedClientCount                *int                                                                                      `json:"connectionSpeedClientCount,omitempty"`                // Connection Speed Client Count
	ConnectionSpeedImpactedEntities           *ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedImpactedEntities           `json:"connectionSpeedImpactedEntities,omitempty"`           //
	ConnectionSpeedFailureImpactedEntities    *ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedFailureImpactedEntities    `json:"connectionSpeedFailureImpactedEntities,omitempty"`    //
	ConnectionSpeedFailureMetrics             *ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedFailureMetrics             `json:"connectionSpeedFailureMetrics,omitempty"`             //
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseCoverageFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingAttemptsFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseOnboardingDurationFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingAttemptsFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseRoamingDurationFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedFailureImpactedEntities struct {
	BuildingCount *int `json:"buildingCount,omitempty"` // Building Count
	FloorCount    *int `json:"floorCount,omitempty"`    // Floor Count
	SitesCount    *int `json:"sitesCount,omitempty"`    // Sites Count
	ApCount       *int `json:"apCount,omitempty"`       // Ap Count
}
type ResponseSitesGetSiteAnalyticsForOneSiteResponseConnectionSpeedFailureMetrics struct {
	FailureApCount     *int `json:"failureApCount,omitempty"`     // Failure Ap Count
	FailureClientCount *int `json:"failureClientCount,omitempty"` // Failure Client Count
	FailurePercentage  *int `json:"failurePercentage,omitempty"`  // Failure Percentage
}
type ResponseSitesAssignDevicesToSite struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesExportMapArchive struct {
	Response *ResponseSitesExportMapArchiveResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  //
}
type ResponseSitesExportMapArchiveResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
} // # Review unknown case
type ResponseSitesImportMapArchiveCancelAnImport interface{}
type ResponseSitesImportMapArchivePerformImport interface{}
type ResponseSitesImportMapArchiveImportStatus struct {
	AuditLog *ResponseSitesImportMapArchiveImportStatusAuditLog `json:"auditLog,omitempty"` //
	Status   string                                             `json:"status,omitempty"`   //
	UUID     *ResponseSitesImportMapArchiveImportStatusUUID     `json:"uuid,omitempty"`     //
}
type ResponseSitesImportMapArchiveImportStatusAuditLog struct {
	Children                   *[]ResponseSitesImportMapArchiveImportStatusAuditLogChildren              `json:"children,omitempty"`                   //
	EntitiesCount              *[]ResponseSitesImportMapArchiveImportStatusAuditLogEntitiesCount         `json:"entitiesCount,omitempty"`              //
	EntityName                 string                                                                    `json:"entityName,omitempty"`                 //
	EntityType                 string                                                                    `json:"entityType,omitempty"`                 //
	ErrorEntitiesCount         *[]ResponseSitesImportMapArchiveImportStatusAuditLogErrorEntitiesCount    `json:"errorEntitiesCount,omitempty"`         //
	Errors                     *[]ResponseSitesImportMapArchiveImportStatusAuditLogErrors                `json:"errors,omitempty"`                     //
	Infos                      *[]ResponseSitesImportMapArchiveImportStatusAuditLogInfos                 `json:"infos,omitempty"`                      //
	MatchingEntitiesCount      *[]ResponseSitesImportMapArchiveImportStatusAuditLogMatchingEntitiesCount `json:"matchingEntitiesCount,omitempty"`      //
	SubTasksRootTaskID         string                                                                    `json:"subTasksRootTaskId,omitempty"`         //
	SuccessfullyImportedFloors []string                                                                  `json:"successfullyImportedFloors,omitempty"` //
	Warnings                   *[]ResponseSitesImportMapArchiveImportStatusAuditLogWarnings              `json:"warnings,omitempty"`                   //
}
type ResponseSitesImportMapArchiveImportStatusAuditLogChildren interface{}
type ResponseSitesImportMapArchiveImportStatusAuditLogEntitiesCount struct {
	Key *int `json:"key,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusAuditLogErrorEntitiesCount struct {
	Key *int `json:"key,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusAuditLogErrors struct {
	Message string `json:"message,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusAuditLogInfos struct {
	Message string `json:"message,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusAuditLogMatchingEntitiesCount struct {
	Key *int `json:"key,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusAuditLogWarnings struct {
	Message string `json:"message,omitempty"` //
}
type ResponseSitesImportMapArchiveImportStatusUUID struct {
	LeastSignificantBits *int `json:"leastSignificantBits,omitempty"` //
	MostSignificantBits  *int `json:"mostSignificantBits,omitempty"`  //
}
type ResponseSitesMapsSupportedAccessPoints []ResponseItemSitesMapsSupportedAccessPoints // Array of ResponseSitesMapsSupportedAccessPoints
type ResponseItemSitesMapsSupportedAccessPoints struct {
	AntennaPatterns *[]ResponseItemSitesMapsSupportedAccessPointsAntennaPatterns `json:"antennaPatterns,omitempty"` //
	ApType          string                                                       `json:"apType,omitempty"`          //
}
type ResponseItemSitesMapsSupportedAccessPointsAntennaPatterns struct {
	Band  string   `json:"band,omitempty"`  //
	Names []string `json:"names,omitempty"` //
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
type ResponseSitesGetSiteHealth struct {
	Response *[]ResponseSitesGetSiteHealthResponse `json:"response,omitempty"` //
}
type ResponseSitesGetSiteHealthResponse struct {
	SiteName                           string                                                     `json:"siteName,omitempty"`                           // Name of the site
	SiteID                             string                                                     `json:"siteId,omitempty"`                             // Site UUID
	ParentSiteID                       string                                                     `json:"parentSiteId,omitempty"`                       // The parent site's UUID of this site
	ParentSiteName                     string                                                     `json:"parentSiteName,omitempty"`                     // The parent site's name of this site
	SiteType                           string                                                     `json:"siteType,omitempty"`                           // Site type of this site
	Latitude                           *float64                                                   `json:"latitude,omitempty"`                           // Site (building) location's latitude
	Longitude                          *float64                                                   `json:"longitude,omitempty"`                          // Site (building) location's longitude
	HealthyNetworkDevicePercentage     *int                                                       `json:"healthyNetworkDevicePercentage,omitempty"`     // Network health of devices on the site
	HealthyClientsPercentage           *int                                                       `json:"healthyClientsPercentage,omitempty"`           // Client health of all clients in the site
	ClientHealthWired                  *int                                                       `json:"clientHealthWired,omitempty"`                  // Health of all wired clients in the site
	ClientHealthWireless               *int                                                       `json:"clientHealthWireless,omitempty"`               // Health of all wireless clients in the site
	NumberOfClients                    *int                                                       `json:"numberOfClients,omitempty"`                    // Total number of clients in the site
	NumberOfNetworkDevice              *int                                                       `json:"numberOfNetworkDevice,omitempty"`              // Total number of network devices in the site
	NetworkHealthAverage               *int                                                       `json:"networkHealthAverage,omitempty"`               // Average network health in the site
	NetworkHealthAccess                *int                                                       `json:"networkHealthAccess,omitempty"`                // Network health for access devices in the site
	NetworkHealthCore                  *int                                                       `json:"networkHealthCore,omitempty"`                  // Network health for core devices in the site
	NetworkHealthDistribution          *int                                                       `json:"networkHealthDistribution,omitempty"`          // Network health for distribution devices in the site
	NetworkHealthRouter                *int                                                       `json:"networkHealthRouter,omitempty"`                // Network health for router devices in the site
	NetworkHealthWireless              *int                                                       `json:"networkHealthWireless,omitempty"`              // Network health for wireless devices in the site
	NetworkHealthAP                    *int                                                       `json:"networkHealthAP,omitempty"`                    // Network health for AP devices in the site
	NetworkHealthWLC                   *int                                                       `json:"networkHealthWLC,omitempty"`                   // Network health for WLC devices in the site
	NetworkHealthSwitch                *int                                                       `json:"networkHealthSwitch,omitempty"`                // Network health for switch devices in the site
	NetworkHealthOthers                *int                                                       `json:"networkHealthOthers,omitempty"`                // Network health for other devices in the site
	NumberOfWiredClients               *int                                                       `json:"numberOfWiredClients,omitempty"`               // Number of wired clients in the site
	NumberOfWirelessClients            *int                                                       `json:"numberOfWirelessClients,omitempty"`            // Number of wireless clients in the site
	TotalNumberOfConnectedWiredClients *int                                                       `json:"totalNumberOfConnectedWiredClients,omitempty"` // Number of connected wired clients in the site
	TotalNumberOfActiveWirelessClients *int                                                       `json:"totalNumberOfActiveWirelessClients,omitempty"` // Number of active wireless clients in the site
	WiredGoodClients                   *int                                                       `json:"wiredGoodClients,omitempty"`                   // Number of GOOD health wired clients in the site
	WirelessGoodClients                *int                                                       `json:"wirelessGoodClients,omitempty"`                // Number of GOOD health wireless clients in the site
	OverallGoodDevices                 *int                                                       `json:"overallGoodDevices,omitempty"`                 // Number of GOOD health devices in the site
	AccessGoodCount                    *int                                                       `json:"accessGoodCount,omitempty"`                    // Number of GOOD health access devices in the site
	AccessTotalCount                   *int                                                       `json:"accessTotalCount,omitempty"`                   // Number of access devices in the site
	CoreGoodCount                      *int                                                       `json:"coreGoodCount,omitempty"`                      // Number of GOOD health core devices in the site
	CoreTotalCount                     *int                                                       `json:"coreTotalCount,omitempty"`                     // Number of core devices in the site
	DistributionGoodCount              *int                                                       `json:"distributionGoodCount,omitempty"`              // Number of GOOD health distribution devices in the site
	DistributionTotalCount             *int                                                       `json:"distributionTotalCount,omitempty"`             // Number of distribution devices in the site
	RouterGoodCount                    *int                                                       `json:"routerGoodCount,omitempty"`                    // Number of GOOD health router in the site
	RouterTotalCount                   *int                                                       `json:"routerTotalCount,omitempty"`                   // Number of router devices in the site
	WirelessDeviceGoodCount            *int                                                       `json:"wirelessDeviceGoodCount,omitempty"`            // Number of GOOD health wireless devices in the site
	WirelessDeviceTotalCount           *int                                                       `json:"wirelessDeviceTotalCount,omitempty"`           // Number of wireless devices in the site
	ApDeviceGoodCount                  *int                                                       `json:"apDeviceGoodCount,omitempty"`                  // Number of GOOD health AP devices in the site
	ApDeviceTotalCount                 *int                                                       `json:"apDeviceTotalCount,omitempty"`                 // Number of AP devices in the site
	WlcDeviceGoodCount                 *int                                                       `json:"wlcDeviceGoodCount,omitempty"`                 // Number of GOOD health wireless controller devices in the site
	WlcDeviceTotalCount                *int                                                       `json:"wlcDeviceTotalCount,omitempty"`                // Number of wireless controller devices in the site
	SwitchDeviceGoodCount              *int                                                       `json:"switchDeviceGoodCount,omitempty"`              // Number of GOOD health switch devices in the site
	SwitchDeviceTotalCount             *int                                                       `json:"switchDeviceTotalCount,omitempty"`             // Number of switch devices in the site
	ApplicationHealth                  *int                                                       `json:"applicationHealth,omitempty"`                  // Average application health in the site
	ApplicationHealthInfo              *[]ResponseSitesGetSiteHealthResponseApplicationHealthInfo `json:"applicationHealthInfo,omitempty"`              //
	ApplicationGoodCount               *int                                                       `json:"applicationGoodCount,omitempty"`               // Number of GOOD health applications int the site
	ApplicationTotalCount              *int                                                       `json:"applicationTotalCount,omitempty"`              // Number of applications int the site
	ApplicationBytesTotalCount         *float64                                                   `json:"applicationBytesTotalCount,omitempty"`         // Total application bytes
	DnacInfo                           *ResponseSitesGetSiteHealthResponseDnacInfo                `json:"dnacInfo,omitempty"`                           //
	Usage                              *float64                                                   `json:"usage,omitempty"`                              // Total bits used by all clients in a site
	ApplicationHealthStats             *ResponseSitesGetSiteHealthResponseApplicationHealthStats  `json:"applicationHealthStats,omitempty"`             //
}
type ResponseSitesGetSiteHealthResponseApplicationHealthInfo struct {
	TrafficClass string   `json:"trafficClass,omitempty"` // Traffic class of the application
	BytesCount   *float64 `json:"bytesCount,omitempty"`   // Byte count of the application
	HealthScore  *int     `json:"healthScore,omitempty"`  // Health score of the application
}
type ResponseSitesGetSiteHealthResponseDnacInfo struct {
	UUID   string `json:"uuid,omitempty"`   // UUID of the DNAC
	IP     string `json:"ip,omitempty"`     // IP address of the DNAC
	Status string `json:"status,omitempty"` // Status of the DNAC
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStats struct {
	AppTotalCount              *float64                                                                            `json:"appTotalCount,omitempty"`              // Total application count
	BusinessRelevantAppCount   *ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessRelevantAppCount   `json:"businessRelevantAppCount,omitempty"`   //
	BusinessIrrelevantAppCount *ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessIrrelevantAppCount `json:"businessIrrelevantAppCount,omitempty"` //
	DefaultHealthAppCount      *ResponseSitesGetSiteHealthResponseApplicationHealthStatsDefaultHealthAppCount      `json:"defaultHealthAppCount,omitempty"`      //
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessRelevantAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor business relevant application count
	Fair *float64 `json:"fair,omitempty"` // Fair business relevant application count
	Good *float64 `json:"good,omitempty"` // Good business relevant application count
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStatsBusinessIrrelevantAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor business irrelevant application count
	Fair *float64 `json:"fair,omitempty"` // Fair business irrelevant application count
	Good *float64 `json:"good,omitempty"` // Good business irrelevant application count
}
type ResponseSitesGetSiteHealthResponseApplicationHealthStatsDefaultHealthAppCount struct {
	Poor *float64 `json:"poor,omitempty"` // Poor default application count
	Fair *float64 `json:"fair,omitempty"` // Fair default application count
	Good *float64 `json:"good,omitempty"` // Good default application count
}
type ResponseSitesGetDevicesThatAreAssignedToASite struct {
	Response *[]ResponseSitesGetDevicesThatAreAssignedToASiteResponse `json:"response,omitempty"` //
}
type ResponseSitesGetDevicesThatAreAssignedToASiteResponse struct {
	InstanceUUID                  string `json:"instanceUuid,omitempty"`                  // Device UUID (E.g. 48eebb3e-b3fc-4928-a7df-1c80e216f930)
	InstanceID                    *int   `json:"instanceId,omitempty"`                    // Device Id (E.g. 230230)
	AuthEntityID                  *int   `json:"authEntityId,omitempty"`                  // Authentication Entity Id (Internal record)
	AuthEntityClass               *int   `json:"authEntityClass,omitempty"`               // Authentication entity class (Internal record)
	InstanceTenantID              string `json:"instanceTenantId,omitempty"`              // Device tenant Id (E.g. 64472cc32d3bc1658597669c)
	DeployPending                 string `json:"deployPending,omitempty"`                 // Deploy pending (Internal record)
	InstanceVersion               *int   `json:"instanceVersion,omitempty"`               // Instance version (Internal record)
	ApManagerInterfaceIP          string `json:"apManagerInterfaceIp,omitempty"`          // Access Point manager interface IP
	AssociatedWlcIP               string `json:"associatedWlcIp,omitempty"`               // Associated Wireless Controller IP
	BootDateTime                  string `json:"bootDateTime,omitempty"`                  // Device boot date and time
	CollectionInterval            string `json:"collectionInterval,omitempty"`            // Device resync interval type (E.g. Global Default)
	CollectionIntervalValue       string `json:"collectionIntervalValue,omitempty"`       // Device resync interval value
	CollectionStatus              string `json:"collectionStatus,omitempty"`              // Device inventory collection status (E.g. Managed)
	Description                   string `json:"description,omitempty"`                   // Device description
	DeviceSupportLevel            string `json:"deviceSupportLevel,omitempty"`            // Device support level (E.g. Supported)
	DNSResolvedManagementAddress  string `json:"dnsResolvedManagementAddress,omitempty"`  // DNS resolved management address
	Family                        string `json:"family,omitempty"`                        // Device family (E.g. Routers)
	Hostname                      string `json:"hostname,omitempty"`                      // Device hostname
	InterfaceCount                string `json:"interfaceCount,omitempty"`                // Device interface count
	InventoryStatusDetail         string `json:"inventoryStatusDetail,omitempty"`         // Device inventory collection status detail
	LastUpdateTime                *int   `json:"lastUpdateTime,omitempty"`                // Last update time
	LastUpdated                   string `json:"lastUpdated,omitempty"`                   // Last updated date and time
	LineCardCount                 string `json:"lineCardCount,omitempty"`                 // Line card count
	LineCardID                    string `json:"lineCardId,omitempty"`                    // Line card Id
	LastDeviceResyncStartTime     string `json:"lastDeviceResyncStartTime,omitempty"`     // Last device inventory resync start date and time
	MacAddress                    string `json:"macAddress,omitempty"`                    // MAC address
	ManagedAtleastOnce            *bool  `json:"managedAtleastOnce,omitempty"`            // If device managed atleast once, value will be true otherwise false
	ManagementIPAddress           string `json:"managementIpAddress,omitempty"`           // Management IP address
	ManagementState               string `json:"managementState,omitempty"`               // Device management state (E.g. Managed)
	MemorySize                    string `json:"memorySize,omitempty"`                    // Memory size
	PaddedMgmtIPAddress           string `json:"paddedMgmtIpAddress,omitempty"`           // Padded management IP address. Internal record
	PendingSyncRequestsCount      string `json:"pendingSyncRequestsCount,omitempty"`      // Pending sync requests count. Internal record
	PlatformID                    string `json:"platformId,omitempty"`                    // Device platform Id (E.g. CSR1000V)
	ReachabilityFailureReason     string `json:"reachabilityFailureReason,omitempty"`     // Device reachability failure reason
	ReachabilityStatus            string `json:"reachabilityStatus,omitempty"`            // Device reachability status (E.g. Reachable)
	ReasonsForDeviceResync        string `json:"reasonsForDeviceResync,omitempty"`        // Reasons for device resync (E.g. Periodic)
	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending device sync requests
	Role                          string `json:"role,omitempty"`                          // Device role (E.g. BORDER ROUTER)
	RoleSource                    string `json:"roleSource,omitempty"`                    // Device role source. Internal record
	SerialNumber                  string `json:"serialNumber,omitempty"`                  // Device serial Number
	Series                        string `json:"series,omitempty"`                        // Device series
	SNMPContact                   string `json:"snmpContact,omitempty"`                   // Device snmp contact. Internal record
	SNMPLocation                  string `json:"snmpLocation,omitempty"`                  // Device snmp location
	SoftwareType                  string `json:"softwareType,omitempty"`                  // Device software type
	SoftwareVersion               string `json:"softwareVersion,omitempty"`               // Device software version
	TagCount                      string `json:"tagCount,omitempty"`                      // Device tag Count
	Type                          string `json:"type,omitempty"`                          // Device type (E.g. Cisco Cloud Services Router 1000V)
	UpTime                        string `json:"upTime,omitempty"`                        // Device up time (E.g. 112 days, 6:09:13.86)
	UptimeSeconds                 *int   `json:"uptimeSeconds,omitempty"`                 // Device uptime in seconds
	Vendor                        string `json:"vendor,omitempty"`                        // Vendor (E.g. Cisco)
	DisplayName                   string `json:"displayName,omitempty"`                   // Device display name
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
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSitesGetSiteV2 struct {
	Response *[]ResponseSitesGetSiteV2Response `json:"response,omitempty"` //
}
type ResponseSitesGetSiteV2Response struct {
	ParentID           string                                          `json:"parentId,omitempty"`           // Parent site Instance UUID (e.g. b27181bb-211b-40ec-ba5d-2603867c3f2c)
	GroupTypeList      []string                                        `json:"groupTypeList,omitempty"`      // There are different group types like 'RBAC', 'POLICY', 'SITE', 'TAG', 'PORT', 'DEVICE_TYPE'. This API is for site, so list contains 'SITE' only
	GroupHierarchy     string                                          `json:"groupHierarchy,omitempty"`     // Site hierarchy by instance UUID (e.g. b27181bb-211b-40ec-ba5d-2603867c3f2c/576c7859-e485-4073-a46f-305f475de4c5)
	AdditionalInfo     *[]ResponseSitesGetSiteV2ResponseAdditionalInfo `json:"additionalInfo,omitempty"`     //
	GroupNameHierarchy string                                          `json:"groupNameHierarchy,omitempty"` // Site hierarchy by name (e.g. Global/USA/CA/San Jose/Building4)
	Name               string                                          `json:"name,omitempty"`               // Site name (e.g. Building4)
	InstanceTenantID   string                                          `json:"instanceTenantId,omitempty"`   // Tenant instance Id where site created (e.g. 63bf047b64ec9c1c45f9019c)
	ID                 string                                          `json:"id,omitempty"`                 // Site instance UUID (e.g. bb5122ce-4527-4af5-8718-44b746a3a3d8)
}
type ResponseSitesGetSiteV2ResponseAdditionalInfo struct {
	NameSpace  string                                                  `json:"nameSpace,omitempty"`  // Site name space. Default value is 'Location'
	Attributes *ResponseSitesGetSiteV2ResponseAdditionalInfoAttributes `json:"attributes,omitempty"` //
}
type ResponseSitesGetSiteV2ResponseAdditionalInfoAttributes struct {
	AddressInheritedFrom string `json:"addressInheritedFrom,omitempty"` // Site instance UUID from where address inherited (e.g. 576c7859-e485-4073-a46f-305f475de4c5)
	Type                 string `json:"type,omitempty"`                 // Site type
	Country              string `json:"country,omitempty"`              // Site Country (e.g. United States)
	Address              string `json:"address,omitempty"`              // Site address (e.g. 269 East Tasman Drive, San Jose, California 95134, United States)
	Latitude             string `json:"latitude,omitempty"`             // Site latitude (e.g. 37.413082)
	Longitude            string `json:"longitude,omitempty"`            // Site longitude (e.g. -121.933886)
}
type ResponseSitesGetSiteCountV2 struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type RequestSitesSubmitRequestToQuerySitesEnergy struct {
	StartTime           *int                                                              `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                              `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestSitesSubmitRequestToQuerySitesEnergyFilters             `json:"filters,omitempty"`             //
	Views               []string                                                          `json:"views,omitempty"`               // Views
	Attributes          []string                                                          `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestSitesSubmitRequestToQuerySitesEnergyAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestSitesSubmitRequestToQuerySitesEnergyPage                  `json:"page,omitempty"`                //
}
type RequestSitesSubmitRequestToQuerySitesEnergyFilters struct {
	LogicalOperator string                                                       `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestSitesSubmitRequestToQuerySitesEnergyFiltersFilters `json:"filters,omitempty"`         //
}
type RequestSitesSubmitRequestToQuerySitesEnergyFiltersFilters struct {
	Key      string   `json:"key,omitempty"`      // Key
	Operator string   `json:"operator,omitempty"` // Operator
	Value    []string `json:"value,omitempty"`    // Value
}
type RequestSitesSubmitRequestToQuerySitesEnergyAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestSitesSubmitRequestToQuerySitesEnergyPage struct {
	Limit  *int                                                     `json:"limit,omitempty"`  // Limit
	Offset *int                                                     `json:"offset,omitempty"` // Offset
	SortBy *[]RequestSitesSubmitRequestToQuerySitesEnergyPageSortBy `json:"sortBy,omitempty"` //
}
type RequestSitesSubmitRequestToQuerySitesEnergyPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type RequestSitesSubmitRequestToCountSitesEnergyFromQuery struct {
	StartTime           *int                                                                       `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                       `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestSitesSubmitRequestToCountSitesEnergyFromQueryFilters             `json:"filters,omitempty"`             //
	Views               []string                                                                   `json:"views,omitempty"`               // Views
	Attributes          []string                                                                   `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestSitesSubmitRequestToCountSitesEnergyFromQueryAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestSitesSubmitRequestToCountSitesEnergyFromQueryPage                  `json:"page,omitempty"`                //
}
type RequestSitesSubmitRequestToCountSitesEnergyFromQueryFilters struct {
	LogicalOperator string                                                                `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestSitesSubmitRequestToCountSitesEnergyFromQueryFiltersFilters `json:"filters,omitempty"`         //
}
type RequestSitesSubmitRequestToCountSitesEnergyFromQueryFiltersFilters struct {
	Key      string   `json:"key,omitempty"`      // Key
	Operator string   `json:"operator,omitempty"` // Operator
	Value    []string `json:"value,omitempty"`    // Value
}
type RequestSitesSubmitRequestToCountSitesEnergyFromQueryAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestSitesSubmitRequestToCountSitesEnergyFromQueryPage struct {
	Limit  *int                                                              `json:"limit,omitempty"`  // Limit
	Offset *int                                                              `json:"offset,omitempty"` // Offset
	SortBy *[]RequestSitesSubmitRequestToCountSitesEnergyFromQueryPageSortBy `json:"sortBy,omitempty"` //
}
type RequestSitesSubmitRequestToCountSitesEnergyFromQueryPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type RequestSitesQueryAnAggregatedSummaryOfSiteHealthData struct {
	StartTime  *int     `json:"startTime,omitempty"`  // Start Time
	EndTime    *int     `json:"endTime,omitempty"`    // End Time
	Views      []string `json:"views,omitempty"`      // Views
	Attributes []string `json:"attributes,omitempty"` // Attributes
}
type RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters struct {
	StartTime  *int                                                                                   `json:"startTime,omitempty"`  // Start Time
	EndTime    *int                                                                                   `json:"endTime,omitempty"`    // End Time
	Filters    *[]RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersFilters `json:"filters,omitempty"`    //
	Attributes []string                                                                               `json:"attributes,omitempty"` // Attributes
	Views      []string                                                                               `json:"views,omitempty"`      // Views
	Page       *RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersPage      `json:"page,omitempty"`       //
}
type RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersPage struct {
	Limit  *int                                                                                    `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                    `json:"offset,omitempty"` // Offset
	SortBy *RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters struct {
	StartTime *int                                                                                          `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                                                          `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersFilters `json:"filters,omitempty"`   //
}
type RequestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestSitesSubmitRequestForSiteAnalyticsSummaryData struct {
	StartTime  *int                                                           `json:"startTime,omitempty"`  // Start Time
	EndTime    *int                                                           `json:"endTime,omitempty"`    // End Time
	Filters    *[]RequestSitesSubmitRequestForSiteAnalyticsSummaryDataFilters `json:"filters,omitempty"`    //
	Attributes []string                                                       `json:"attributes,omitempty"` // Attributes
}
type RequestSitesSubmitRequestForSiteAnalyticsSummaryDataFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics struct {
	StartTime *int                                                                     `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                                     `json:"endTime,omitempty"`   // End Time
	TopN      *int                                                                     `json:"topN,omitempty"`      // Top N
	GroupBy   []string                                                                 `json:"groupBy,omitempty"`   // Group By
	Filters   *[]RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters `json:"filters,omitempty"`   //
}
type RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalyticsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestSitesSubmitRequestForSiteAnalyticsTrendData struct {
	StartTime     *int                                                         `json:"startTime,omitempty"`     // Start Time
	EndTime       *int                                                         `json:"endTime,omitempty"`       // End Time
	TrendInterval string                                                       `json:"trendInterval,omitempty"` // Trend Interval
	Filters       *[]RequestSitesSubmitRequestForSiteAnalyticsTrendDataFilters `json:"filters,omitempty"`       //
	Attributes    []string                                                     `json:"attributes,omitempty"`    // Attributes
	Page          *RequestSitesSubmitRequestForSiteAnalyticsTrendDataPage      `json:"page,omitempty"`          //
}
type RequestSitesSubmitRequestForSiteAnalyticsTrendDataFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestSitesSubmitRequestForSiteAnalyticsTrendDataPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestSitesAssignDevicesToSite struct {
	Device *[]RequestSitesAssignDevicesToSiteDevice `json:"device,omitempty"` //
}
type RequestSitesAssignDevicesToSiteDevice struct {
	IP string `json:"ip,omitempty"` // Device IP. It can be either IPv4 or IPv6. IPV4 e.g., 10.104.240.64. IPV6 e.g., 2001:420:284:2004:4:181:500:183
} // # Review unknown case
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
	ParentName string `json:"parentName,omitempty"` // Parent hierarchical name (Example: Global/USA/CA)
}
type RequestSitesCreateSiteSiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Name of the building (eg: building1)
	Address    string   `json:"address,omitempty"`    // Address of the building to be created
	ParentName string   `json:"parentName,omitempty"` // Parent hierarchical name (Example: Global/USA/CA/SantaClara)
	Latitude   *float64 `json:"latitude,omitempty"`   // Latitude coordinate of the building (eg:37.338)
	Longitude  *float64 `json:"longitude,omitempty"`  // Longitude coordinate of the building (eg:-121.832)
	Country    string   `json:"country,omitempty"`    // Country (eg:United States)
}
type RequestSitesCreateSiteSiteFloor struct {
	Name        string   `json:"name,omitempty"`        // Name of the floor (eg:floor-1)
	ParentName  string   `json:"parentName,omitempty"`  // Parent hierarchical name (Example: Global/USA/CA/SantaClara/Academic)
	RfModel     string   `json:"rfModel,omitempty"`     // Type of floor (eg: Cubes And Walled Offices0
	Width       *float64 `json:"width,omitempty"`       // Width of the floor. Unit of measure is ft. (eg: 100)
	Length      *float64 `json:"length,omitempty"`      // Length of the floor. Unit of measure is ft. (eg: 100)
	Height      *float64 `json:"height,omitempty"`      // Height of the floor. Unit of measure is ft. (eg: 15)
	FloorNumber *float64 `json:"floorNumber,omitempty"` // Floor number. (eg: 5)
}
type RequestSitesUpdateSite struct {
	Type string                      `json:"type,omitempty"` // Site type
	Site *RequestSitesUpdateSiteSite `json:"site,omitempty"` //
}
type RequestSitesUpdateSiteSite struct {
	Area     *RequestSitesUpdateSiteSiteArea     `json:"area,omitempty"`     //
	Building *RequestSitesUpdateSiteSiteBuilding `json:"building,omitempty"` //
	Floor    *RequestSitesUpdateSiteSiteFloor    `json:"floor,omitempty"`    //
}
type RequestSitesUpdateSiteSiteArea struct {
	Name       string `json:"name,omitempty"`       // Area name
	ParentName string `json:"parentName,omitempty"` // Parent hierarchical name (Example: Global/USA/CA)
}
type RequestSitesUpdateSiteSiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Building name
	Address    string   `json:"address,omitempty"`    // Building address (Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States)
	ParentName string   `json:"parentName,omitempty"` // Parent hierarchical name (Example: Global/USA/CA/SantaClara)
	Latitude   *float64 `json:"latitude,omitempty"`   // Building latitude (Example: 37.403712)
	Longitude  *float64 `json:"longitude,omitempty"`  // Building longitude (Example: -121.971063)
	Country    string   `json:"country,omitempty"`    // Country name. This field is mandatory for air-gapped networks (Example: United States)
}
type RequestSitesUpdateSiteSiteFloor struct {
	Name        string   `json:"name,omitempty"`        // Floor name
	RfModel     string   `json:"rfModel,omitempty"`     // RF model (Example : Cubes And Walled Offices)
	Width       *float64 `json:"width,omitempty"`       // Floor width in feet (Example: 200)
	Length      *float64 `json:"length,omitempty"`      // Floor length in feet (Example: 100)
	Height      *float64 `json:"height,omitempty"`      // Floor height in feet (Example: 10)
	FloorNumber *float64 `json:"floorNumber,omitempty"` // Floor Number (Example: 3)
	ParentName  string   `json:"parentName,omitempty"`  // Parent hierarchical name (Example: Global/USA/CA/SantaClara/Academic)
}

//GetSitesEnergy Get sites energy - 5985-7ad7-4b3a-b26a
/* Retrieves a list of sites with energy data based on the specified query parameters. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param GetSitesEnergyHeaderParams Custom header parameters
@param GetSitesEnergyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sites-energy
*/
func (s *SitesService) GetSitesEnergy(GetSitesEnergyHeaderParams *GetSitesEnergyHeaderParams, GetSitesEnergyQueryParams *GetSitesEnergyQueryParams) (*ResponseSitesGetSitesEnergy, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites"

	queryString, _ := query.Values(GetSitesEnergyQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSitesEnergyHeaderParams != nil {

		if GetSitesEnergyHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSitesEnergyHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSitesEnergy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSitesEnergy(GetSitesEnergyHeaderParams, GetSitesEnergyQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSitesEnergy")
	}

	result := response.Result().(*ResponseSitesGetSitesEnergy)
	return result, response, err

}

//CountSitesEnergy Count sites energy - b9b8-5b6f-4ae8-8e7c
/* Retrieves the total count of sites that provide energy data, filtered according to the specified query parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param CountSitesEnergyHeaderParams Custom header parameters
@param CountSitesEnergyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-sites-energy
*/
func (s *SitesService) CountSitesEnergy(CountSitesEnergyHeaderParams *CountSitesEnergyHeaderParams, CountSitesEnergyQueryParams *CountSitesEnergyQueryParams) (*ResponseSitesCountSitesEnergy, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites/count"

	queryString, _ := query.Values(CountSitesEnergyQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountSitesEnergyHeaderParams != nil {

		if CountSitesEnergyHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountSitesEnergyHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesCountSitesEnergy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountSitesEnergy(CountSitesEnergyHeaderParams, CountSitesEnergyQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountSitesEnergy")
	}

	result := response.Result().(*ResponseSitesCountSitesEnergy)
	return result, response, err

}

//QuerySitesEnergyForTheGivenTaskID Query sites energy for the given task ID - e28f-4afb-44fa-989d
/* Gets query sites energy task result for the given task ID. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param QuerySitesEnergyForTheGivenTaskIDQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-sites-energy-for-the-given-task-id
*/
func (s *SitesService) QuerySitesEnergyForTheGivenTaskID(QuerySitesEnergyForTheGivenTaskIDQueryParams *QuerySitesEnergyForTheGivenTaskIDQueryParams) (*ResponseSitesQuerySitesEnergyForTheGivenTaskID, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites/query"

	queryString, _ := query.Values(QuerySitesEnergyForTheGivenTaskIDQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesQuerySitesEnergyForTheGivenTaskID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.QuerySitesEnergyForTheGivenTaskID(QuerySitesEnergyForTheGivenTaskIDQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation QuerySitesEnergyForTheGivenTaskId")
	}

	result := response.Result().(*ResponseSitesQuerySitesEnergyForTheGivenTaskID)
	return result, response, err

}

//CountSitesEnergyForTheGivenTaskID Count sites energy for the given task ID - 75ae-3959-4c1a-b993
/* Gets count sites energy task result for the given task ID. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param CountSitesEnergyForTheGivenTaskIDHeaderParams Custom header parameters
@param CountSitesEnergyForTheGivenTaskIDQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-sites-energy-for-the-given-task-id
*/
func (s *SitesService) CountSitesEnergyForTheGivenTaskID(CountSitesEnergyForTheGivenTaskIDHeaderParams *CountSitesEnergyForTheGivenTaskIDHeaderParams, CountSitesEnergyForTheGivenTaskIDQueryParams *CountSitesEnergyForTheGivenTaskIDQueryParams) (*ResponseSitesCountSitesEnergyForTheGivenTaskID, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites/query/count"

	queryString, _ := query.Values(CountSitesEnergyForTheGivenTaskIDQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountSitesEnergyForTheGivenTaskIDHeaderParams != nil {

		if CountSitesEnergyForTheGivenTaskIDHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountSitesEnergyForTheGivenTaskIDHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesCountSitesEnergyForTheGivenTaskID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountSitesEnergyForTheGivenTaskID(CountSitesEnergyForTheGivenTaskIDHeaderParams, CountSitesEnergyForTheGivenTaskIDQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountSitesEnergyForTheGivenTaskId")
	}

	result := response.Result().(*ResponseSitesCountSitesEnergyForTheGivenTaskID)
	return result, response, err

}

//GetSiteEnergyByID Get site energy by ID - 62ad-aba7-4ec8-9429
/* Retrieve the energy summary data for a specific site based on the site ID. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param id id path parameter. The UUID of the Site. (Ex. "6bef213c-19ca-4170-8375-b694e251101c")

@param GetSiteEnergyByIDHeaderParams Custom header parameters
@param GetSiteEnergyByIDQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-energy-by-id
*/
func (s *SitesService) GetSiteEnergyByID(id string, GetSiteEnergyByIDHeaderParams *GetSiteEnergyByIDHeaderParams, GetSiteEnergyByIDQueryParams *GetSiteEnergyByIDQueryParams) (*ResponseSitesGetSiteEnergyByID, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetSiteEnergyByIDQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSiteEnergyByIDHeaderParams != nil {

		if GetSiteEnergyByIDHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSiteEnergyByIDHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteEnergyByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteEnergyByID(id, GetSiteEnergyByIDHeaderParams, GetSiteEnergyByIDQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteEnergyById")
	}

	result := response.Result().(*ResponseSitesGetSiteEnergyByID)
	return result, response, err

}

//ReadListOfSiteHealthSummaries Read list of site health summaries. - e4b7-1b5e-4099-b15b
/* Get a paginated list of site health summaries. Use the available query parameters to identify a subset of sites you want health summaries for. This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime. Valid values for `sortBy` param in this API are limited to the attributes provided in the `site` view. Default sortBy is 'siteHierarchy' in order 'asc' (ascending). For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param ReadListOfSiteHealthSummariesHeaderParams Custom header parameters
@param ReadListOfSiteHealthSummariesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-list-of-site-health-summaries
*/
func (s *SitesService) ReadListOfSiteHealthSummaries(ReadListOfSiteHealthSummariesHeaderParams *ReadListOfSiteHealthSummariesHeaderParams, ReadListOfSiteHealthSummariesQueryParams *ReadListOfSiteHealthSummariesQueryParams) (*ResponseSitesReadListOfSiteHealthSummaries, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries"

	queryString, _ := query.Values(ReadListOfSiteHealthSummariesQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadListOfSiteHealthSummariesHeaderParams != nil {

		if ReadListOfSiteHealthSummariesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadListOfSiteHealthSummariesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadListOfSiteHealthSummaries{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadListOfSiteHealthSummaries(ReadListOfSiteHealthSummariesHeaderParams, ReadListOfSiteHealthSummariesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadListOfSiteHealthSummaries")
	}

	result := response.Result().(*ResponseSitesReadListOfSiteHealthSummaries)
	return result, response, err

}

//ReadSiteCount Read site count. - b6ac-283b-4f39-b488
/* Get a count of sites. Use the available query parameters to get the count of a subset of sites. This API provides the latest data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param ReadSiteCountHeaderParams Custom header parameters
@param ReadSiteCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-site-count
*/
func (s *SitesService) ReadSiteCount(ReadSiteCountHeaderParams *ReadSiteCountHeaderParams, ReadSiteCountQueryParams *ReadSiteCountQueryParams) (*ResponseSitesReadSiteCount, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/count"

	queryString, _ := query.Values(ReadSiteCountQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadSiteCountHeaderParams != nil {

		if ReadSiteCountHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadSiteCountHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadSiteCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadSiteCount(ReadSiteCountHeaderParams, ReadSiteCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadSiteCount")
	}

	result := response.Result().(*ResponseSitesReadSiteCount)
	return result, response, err

}

//ReadAnAggregatedSummaryOfSiteHealthData Read an aggregated summary of site health data. - e2b7-8b97-4c78-a4f7
/* Get an aggregated summary of all site health or use the query params to get an aggregated summary of health for a subset of sites. This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime. Aggregated response data will NOT have unique identifier data populated. List of unique identifier data: [`id`, `siteHierarchy`, `siteHierarchyId`, `siteType`, `latitude`, `longitude`]. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams Custom header parameters
@param ReadAnAggregatedSummaryOfSiteHealthDataQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-an-aggregated-summary-of-site-health-data
*/
func (s *SitesService) ReadAnAggregatedSummaryOfSiteHealthData(ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams *ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams, ReadAnAggregatedSummaryOfSiteHealthDataQueryParams *ReadAnAggregatedSummaryOfSiteHealthDataQueryParams) (*ResponseSitesReadAnAggregatedSummaryOfSiteHealthData, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/summaryAnalytics"

	queryString, _ := query.Values(ReadAnAggregatedSummaryOfSiteHealthDataQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams != nil {

		if ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadAnAggregatedSummaryOfSiteHealthData{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadAnAggregatedSummaryOfSiteHealthData(ReadAnAggregatedSummaryOfSiteHealthDataHeaderParams, ReadAnAggregatedSummaryOfSiteHealthDataQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadAnAggregatedSummaryOfSiteHealthData")
	}

	result := response.Result().(*ResponseSitesReadAnAggregatedSummaryOfSiteHealthData)
	return result, response, err

}

//ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork Read trend analytics data for a grouping of sites in your network - 7285-a858-45aa-b953
/* Retrieves the time series information of health and issue data for sites specified by query parameters, or all sites. The data will be grouped based on the specified trend time interval. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-2.0.0-resolved.yaml


@param ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams Custom header parameters
@param ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-trend-analytics-data-for-a-grouping-of-sites-in-your-network
*/
func (s *SitesService) ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork(ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams *ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams, ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams *ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams) (*ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/trendAnalytics"

	queryString, _ := query.Values(ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams != nil {

		if ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork(ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkHeaderParams, ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork")
	}

	result := response.Result().(*ResponseSitesReadTrendAnalyticsDataForAGroupingOfSitesInYourNetwork)
	return result, response, err

}

//ReadSiteHealthSummaryDataBySiteID Read site health summary data by site id. - 48aa-094f-423b-bd33
/* Get a health summary for a specific site by providing the unique site id in the url path. This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param id id path parameter. unique site uuid

@param ReadSiteHealthSummaryDataBySiteIdHeaderParams Custom header parameters
@param ReadSiteHealthSummaryDataBySiteIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-site-health-summary-data-by-site-id
*/
func (s *SitesService) ReadSiteHealthSummaryDataBySiteID(id string, ReadSiteHealthSummaryDataBySiteIdHeaderParams *ReadSiteHealthSummaryDataBySiteIDHeaderParams, ReadSiteHealthSummaryDataBySiteIdQueryParams *ReadSiteHealthSummaryDataBySiteIDQueryParams) (*ResponseSitesReadSiteHealthSummaryDataBySiteID, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(ReadSiteHealthSummaryDataBySiteIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadSiteHealthSummaryDataBySiteIdHeaderParams != nil {

		if ReadSiteHealthSummaryDataBySiteIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadSiteHealthSummaryDataBySiteIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadSiteHealthSummaryDataBySiteID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadSiteHealthSummaryDataBySiteID(id, ReadSiteHealthSummaryDataBySiteIdHeaderParams, ReadSiteHealthSummaryDataBySiteIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadSiteHealthSummaryDataBySiteId")
	}

	result := response.Result().(*ResponseSitesReadSiteHealthSummaryDataBySiteID)
	return result, response, err

}

//ReadTrendAnalyticsDataForASpecificSiteInYourNetwork Read trend analytics data for a specific site in your network - e1b4-5aa7-47db-b848
/* Retrieves the time series information of health and issue data for a site specified by the path parameter. The data will be grouped based on the specified trend time interval. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-2.0.0-resolved.yaml


@param id id path parameter. unique site uuid

@param ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams Custom header parameters
@param ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-trend-analytics-data-for-a-specific-site-in-your-network
*/
func (s *SitesService) ReadTrendAnalyticsDataForASpecificSiteInYourNetwork(id string, ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams *ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams, ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams *ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams) (*ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetwork, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams != nil {

		if ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadTrendAnalyticsDataForASpecificSiteInYourNetwork(id, ReadTrendAnalyticsDataForASpecificSiteInYourNetworkHeaderParams, ReadTrendAnalyticsDataForASpecificSiteInYourNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadTrendAnalyticsDataForASpecificSiteInYourNetwork")
	}

	result := response.Result().(*ResponseSitesReadTrendAnalyticsDataForASpecificSiteInYourNetwork)
	return result, response, err

}

//GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters Get site analytics for the child sites of given parent site and other query parameters. - e6b7-0a0e-4709-9fd3
/* Returns site analytics for all child sites of given parent site. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams Custom header parameters
@param GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-analytics-for-the-child-sites-of-given-parent-site-and-other-query-parameters
*/
func (s *SitesService) GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters(GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams *GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams, GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams *GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams) (*ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries"

	queryString, _ := query.Values(GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams != nil {

		if GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters(GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersHeaderParams, GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters")
	}

	result := response.Result().(*ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherQueryParameters)
	return result, response, err

}

//GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters Get the total number of site analytics records available for for given set of query parameters. - cbb0-1a0d-4bab-940b
/* Returns the total number of site analytics records available for for given set of query parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams Custom header parameters
@param GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-number-of-site-analytics-records-available-for-for-given-set-of-query-parameters
*/
func (s *SitesService) GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters(GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams *GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams, GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersQueryParams *GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersQueryParams) (*ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/count"

	queryString, _ := query.Values(GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams != nil {

		if GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters(GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersHeaderParams, GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters")
	}

	result := response.Result().(*ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfQueryParameters)
	return result, response, err

}

//GetSiteAnalyticsSummaryDataForTheGivenTaskID Get site analytics summary data for the given task id - a685-a977-4cea-b6e9
/* Get site analytics summary data for the given task id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetSiteAnalyticsSummaryDataForTheGivenTaskIdHeaderParams Custom header parameters
@param GetSiteAnalyticsSummaryDataForTheGivenTaskIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-analytics-summary-data-for-the-given-task-id
*/
func (s *SitesService) GetSiteAnalyticsSummaryDataForTheGivenTaskID(GetSiteAnalyticsSummaryDataForTheGivenTaskIdHeaderParams *GetSiteAnalyticsSummaryDataForTheGivenTaskIDHeaderParams, GetSiteAnalyticsSummaryDataForTheGivenTaskIdQueryParams *GetSiteAnalyticsSummaryDataForTheGivenTaskIDQueryParams) (*ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskID, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/summaryAnalytics"

	queryString, _ := query.Values(GetSiteAnalyticsSummaryDataForTheGivenTaskIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSiteAnalyticsSummaryDataForTheGivenTaskIdHeaderParams != nil {

		if GetSiteAnalyticsSummaryDataForTheGivenTaskIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSiteAnalyticsSummaryDataForTheGivenTaskIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAnalyticsSummaryDataForTheGivenTaskID(GetSiteAnalyticsSummaryDataForTheGivenTaskIdHeaderParams, GetSiteAnalyticsSummaryDataForTheGivenTaskIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAnalyticsSummaryDataForTheGivenTaskId")
	}

	result := response.Result().(*ResponseSitesGetSiteAnalyticsSummaryDataForTheGivenTaskID)
	return result, response, err

}

//GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID Get Top N entities related to site analytics for the given task id - a29f-d81b-4279-91bb
/* Gets the topN analytics data for a given taskId. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdHeaderParams Custom header parameters
@param GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-entities-related-to-site-analytics-for-the-given-task-id
*/
func (s *SitesService) GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdHeaderParams *GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDHeaderParams, GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdQueryParams *GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDQueryParams) (*ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/topNAnalytics"

	queryString, _ := query.Values(GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdHeaderParams != nil {

		if GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID(GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdHeaderParams, GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskId")
	}

	result := response.Result().(*ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskID)
	return result, response, err

}

//GetSiteAnalyticsTrendDataForTheGivenTaskID Get site analytics trend data for the given task id - 4fac-396e-4cc8-993a
/* Gets site analytics trend data for the given task id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetSiteAnalyticsTrendDataForTheGivenTaskIdHeaderParams Custom header parameters
@param GetSiteAnalyticsTrendDataForTheGivenTaskIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-analytics-trend-data-for-the-given-task-id
*/
func (s *SitesService) GetSiteAnalyticsTrendDataForTheGivenTaskID(GetSiteAnalyticsTrendDataForTheGivenTaskIdHeaderParams *GetSiteAnalyticsTrendDataForTheGivenTaskIDHeaderParams, GetSiteAnalyticsTrendDataForTheGivenTaskIdQueryParams *GetSiteAnalyticsTrendDataForTheGivenTaskIDQueryParams) (*ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskID, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/trendAnalytics"

	queryString, _ := query.Values(GetSiteAnalyticsTrendDataForTheGivenTaskIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSiteAnalyticsTrendDataForTheGivenTaskIdHeaderParams != nil {

		if GetSiteAnalyticsTrendDataForTheGivenTaskIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSiteAnalyticsTrendDataForTheGivenTaskIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAnalyticsTrendDataForTheGivenTaskID(GetSiteAnalyticsTrendDataForTheGivenTaskIdHeaderParams, GetSiteAnalyticsTrendDataForTheGivenTaskIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAnalyticsTrendDataForTheGivenTaskId")
	}

	result := response.Result().(*ResponseSitesGetSiteAnalyticsTrendDataForTheGivenTaskID)
	return result, response, err

}

//GetSiteAnalyticsForOneSite Get site analytics for one site. - 65be-bad5-48b9-9134
/* Returns site analytics for the given site. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param id id path parameter. The Site UUID

@param GetSiteAnalyticsForOneSiteHeaderParams Custom header parameters
@param GetSiteAnalyticsForOneSiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-analytics-for-one-site
*/
func (s *SitesService) GetSiteAnalyticsForOneSite(id string, GetSiteAnalyticsForOneSiteHeaderParams *GetSiteAnalyticsForOneSiteHeaderParams, GetSiteAnalyticsForOneSiteQueryParams *GetSiteAnalyticsForOneSiteQueryParams) (*ResponseSitesGetSiteAnalyticsForOneSite, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetSiteAnalyticsForOneSiteQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSiteAnalyticsForOneSiteHeaderParams != nil {

		if GetSiteAnalyticsForOneSiteHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSiteAnalyticsForOneSiteHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteAnalyticsForOneSite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAnalyticsForOneSite(id, GetSiteAnalyticsForOneSiteHeaderParams, GetSiteAnalyticsForOneSiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAnalyticsForOneSite")
	}

	result := response.Result().(*ResponseSitesGetSiteAnalyticsForOneSite)
	return result, response, err

}

//ImportMapArchiveImportStatus Import Map Archive - Import Status - d9b0-599f-4d88-9a9a
/* Gets the status of a map archive import operation. For a map archive import that has just been initiated, will provide the result of validation of the archive and a pre-import preview of what will be performed if the import is performed.  Once an import is requested to be performed, this API will give the status of the import and upon completion a post-import summary of what was performed by the operation.


@param importContextUUID importContextUuid path parameter. The unique import context UUID given by a previous and recent call to maps/import/start API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-import-status
*/
func (s *SitesService) ImportMapArchiveImportStatus(importContextUUID string) (*ResponseSitesImportMapArchiveImportStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/import/{importContextUuid}/status"
	path = strings.Replace(path, "{importContextUuid}", fmt.Sprintf("%v", importContextUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesImportMapArchiveImportStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchiveImportStatus(importContextUUID)
		}
		return nil, response, fmt.Errorf("error with operation ImportMapArchiveImportStatus")
	}

	result := response.Result().(*ResponseSitesImportMapArchiveImportStatus)
	return result, response, err

}

//MapsSupportedAccessPoints Maps Supported Access Points - 97b4-9a04-403b-bc05
/* Gets the list of supported access point types as well as valid antenna pattern names that can be used for each.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!maps-supported-access-points
*/
func (s *SitesService) MapsSupportedAccessPoints() (*ResponseSitesMapsSupportedAccessPoints, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/supported-access-points"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesMapsSupportedAccessPoints{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.MapsSupportedAccessPoints()
		}
		return nil, response, fmt.Errorf("error with operation MapsSupportedAccessPoints")
	}

	result := response.Result().(*ResponseSitesMapsSupportedAccessPoints)
	return result, response, err

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
/* Get site(s) by site-name-hierarchy or siteId or type. List all sites if these parameters are not given as an input.


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

//GetSiteHealth Get Site Health - 2597-2a31-43c8-8729
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

//GetDevicesThatAreAssignedToASite Get devices that are assigned to a site - ae86-1be7-4d39-b0d1
/* API to get devices that are assigned to a site.


@param id id path parameter. Site Id

@param GetDevicesThatAreAssignedToASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-that-are-assigned-to-a-site
*/
func (s *SitesService) GetDevicesThatAreAssignedToASite(id string, GetDevicesThatAreAssignedToASiteQueryParams *GetDevicesThatAreAssignedToASiteQueryParams) (*ResponseSitesGetDevicesThatAreAssignedToASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/site-member/{id}/member"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDevicesThatAreAssignedToASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetDevicesThatAreAssignedToASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesThatAreAssignedToASite(id, GetDevicesThatAreAssignedToASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesThatAreAssignedToASite")
	}

	result := response.Result().(*ResponseSitesGetDevicesThatAreAssignedToASite)
	return result, response, err

}

//GetSiteCount Get Site Count - b0b7-eabc-4f4b-9b28
/* Get the site count of the specified site's sub-hierarchy (inclusive of the provided site)


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

//GetSiteV2 Get Site V2 - 6490-7aa4-40c9-ba3a
/* API to get site(s) by site-name-hierarchy or siteId or type. List all sites if these parameters  are not given as an input.


@param GetSiteV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-v2
*/
func (s *SitesService) GetSiteV2(GetSiteV2QueryParams *GetSiteV2QueryParams) (*ResponseSitesGetSiteV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/site"

	queryString, _ := query.Values(GetSiteV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteV2(GetSiteV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteV2")
	}

	result := response.Result().(*ResponseSitesGetSiteV2)
	return result, response, err

}

//GetSiteCountV2 Get Site Count V2 - 24a8-fb4c-4fbb-9a47
/* Get the site count of the specified site's sub-hierarchy (inclusive of the provided site)


@param GetSiteCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-count-v2
*/
func (s *SitesService) GetSiteCountV2(GetSiteCountV2QueryParams *GetSiteCountV2QueryParams) (*ResponseSitesGetSiteCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/site/count"

	queryString, _ := query.Values(GetSiteCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetSiteCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteCountV2(GetSiteCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteCountV2")
	}

	result := response.Result().(*ResponseSitesGetSiteCountV2)
	return result, response, err

}

//SubmitRequestToQuerySitesEnergy Submit request to query sites energy - a1af-fadd-401a-bbdf
/* Submits a request to retrieve a list of sites along with their energy data for a specified time range, based on the filters provided in the request body. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param SubmitRequestToQuerySitesEnergyHeaderParams Custom header parameters
@param SubmitRequestToQuerySitesEnergyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!submit-request-to-query-sites-energy
*/
func (s *SitesService) SubmitRequestToQuerySitesEnergy(requestSitesSubmitRequestToQuerySitesEnergy *RequestSitesSubmitRequestToQuerySitesEnergy, SubmitRequestToQuerySitesEnergyHeaderParams *SubmitRequestToQuerySitesEnergyHeaderParams, SubmitRequestToQuerySitesEnergyQueryParams *SubmitRequestToQuerySitesEnergyQueryParams) (*ResponseSitesSubmitRequestToQuerySitesEnergy, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites/query"

	queryString, _ := query.Values(SubmitRequestToQuerySitesEnergyQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if SubmitRequestToQuerySitesEnergyHeaderParams != nil {

		if SubmitRequestToQuerySitesEnergyHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", SubmitRequestToQuerySitesEnergyHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetBody(requestSitesSubmitRequestToQuerySitesEnergy).
		SetResult(&ResponseSitesSubmitRequestToQuerySitesEnergy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitRequestToQuerySitesEnergy(requestSitesSubmitRequestToQuerySitesEnergy, SubmitRequestToQuerySitesEnergyHeaderParams, SubmitRequestToQuerySitesEnergyQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation SubmitRequestToQuerySitesEnergy")
	}

	result := response.Result().(*ResponseSitesSubmitRequestToQuerySitesEnergy)
	return result, response, err

}

//SubmitRequestToCountSitesEnergyFromQuery Submit request to count sites energy from query - 03a3-496c-4c0b-86a9
/* Submits a request to retrieve the total count of sites that provide energy data, filtered according to the specified query parameters. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-sitesEnergy-1.0.1-resolved.yaml


@param SubmitRequestToCountSitesEnergyFromQueryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!submit-request-to-count-sites-energy-from-query
*/
func (s *SitesService) SubmitRequestToCountSitesEnergyFromQuery(requestSitesSubmitRequestToCountSitesEnergyFromQuery *RequestSitesSubmitRequestToCountSitesEnergyFromQuery, SubmitRequestToCountSitesEnergyFromQueryQueryParams *SubmitRequestToCountSitesEnergyFromQueryQueryParams) (*ResponseSitesSubmitRequestToCountSitesEnergyFromQuery, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/sites/query/count"

	queryString, _ := query.Values(SubmitRequestToCountSitesEnergyFromQueryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSitesSubmitRequestToCountSitesEnergyFromQuery).
		SetResult(&ResponseSitesSubmitRequestToCountSitesEnergyFromQuery{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitRequestToCountSitesEnergyFromQuery(requestSitesSubmitRequestToCountSitesEnergyFromQuery, SubmitRequestToCountSitesEnergyFromQueryQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation SubmitRequestToCountSitesEnergyFromQuery")
	}

	result := response.Result().(*ResponseSitesSubmitRequestToCountSitesEnergyFromQuery)
	return result, response, err

}

//QueryAnAggregatedSummaryOfSiteHealthData Query an aggregated summary of site health data. - 2782-ca59-4cc8-ad34
/* Query an aggregated summary of all site health This API provides the latest health data from a given `endTime` If data is not ready for the provided endTime, the request will fail, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data. This API also provides issue data. The `startTime` query param can be used to specify the beginning point of time range to retrieve the active issue counts in. When this param is not provided, the default `startTime` will be 24 hours before endTime.


 Aggregated response data will NOT have unique identifier data populated.

 List of unique identifier data: [`id`, `siteHierarchy`,
`siteHierarchyId`, `siteType`, `latitude`, `longitude`] Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-siteHealthSummaries-1.0.3-resolved.yaml


@param QueryAnAggregatedSummaryOfSiteHealthDataQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-an-aggregated-summary-of-site-health-data
*/
func (s *SitesService) QueryAnAggregatedSummaryOfSiteHealthData(requestSitesQueryAnAggregatedSummaryOfSiteHealthData *RequestSitesQueryAnAggregatedSummaryOfSiteHealthData, QueryAnAggregatedSummaryOfSiteHealthDataQueryParams *QueryAnAggregatedSummaryOfSiteHealthDataQueryParams) (*ResponseSitesQueryAnAggregatedSummaryOfSiteHealthData, *resty.Response, error) {
	path := "/dna/data/api/v1/siteHealthSummaries/summaryAnalytics"

	queryString, _ := query.Values(QueryAnAggregatedSummaryOfSiteHealthDataQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSitesQueryAnAggregatedSummaryOfSiteHealthData).
		SetResult(&ResponseSitesQueryAnAggregatedSummaryOfSiteHealthData{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAnAggregatedSummaryOfSiteHealthData(requestSitesQueryAnAggregatedSummaryOfSiteHealthData, QueryAnAggregatedSummaryOfSiteHealthDataQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation QueryAnAggregatedSummaryOfSiteHealthData")
	}

	result := response.Result().(*ResponseSitesQueryAnAggregatedSummaryOfSiteHealthData)
	return result, response, err

}

//GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters Get site analytics for the child sites of given parent site and other filters. - 31bb-a81f-467a-b596
/* Returns site analytics for all child sites of given parent site. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-analytics-for-the-child-sites-of-given-parent-site-and-other-filters
*/
func (s *SitesService) GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters(requestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters *RequestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters, GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams *GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams) (*ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams != nil {

		if GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters).
		SetResult(&ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters(requestSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters, GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters")
	}

	result := response.Result().(*ResponseSitesGetSiteAnalyticsForTheChildSitesOfGivenParentSiteAndOtherFilters)
	return result, response, err

}

//GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters Get the total number of site analytics records available for for given set of filters. - f19a-d9e4-4c19-9399
/* Returns the total number of site analytics records available for for given set of filters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml


@param GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-number-of-site-analytics-records-available-for-for-given-set-of-filters
*/
func (s *SitesService) GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters(requestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters *RequestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters, GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams *GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams) (*ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams != nil {

		if GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters).
		SetResult(&ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters(requestSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters, GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters")
	}

	result := response.Result().(*ResponseSitesGetTheTotalNumberOfSiteAnalyticsRecordsAvailableForForGivenSetOfFilters)
	return result, response, err

}

//SubmitRequestForSiteAnalyticsSummaryData Submit request for site analytics summary data - efb8-19ca-4eba-b9d1
/* Submits the task to get summary analytics data for a given site. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!submit-request-for-site-analytics-summary-data
*/
func (s *SitesService) SubmitRequestForSiteAnalyticsSummaryData(requestSitesSubmitRequestForSiteAnalyticsSummaryData *RequestSitesSubmitRequestForSiteAnalyticsSummaryData) (*ResponseSitesSubmitRequestForSiteAnalyticsSummaryData, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/summaryAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSitesSubmitRequestForSiteAnalyticsSummaryData).
		SetResult(&ResponseSitesSubmitRequestForSiteAnalyticsSummaryData{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitRequestForSiteAnalyticsSummaryData(requestSitesSubmitRequestForSiteAnalyticsSummaryData)
		}

		return nil, response, fmt.Errorf("error with operation SubmitRequestForSiteAnalyticsSummaryData")
	}

	result := response.Result().(*ResponseSitesSubmitRequestForSiteAnalyticsSummaryData)
	return result, response, err

}

//SubmitRequestForTopNEntitiesRelatedToSiteAnalytics Submit request for Top N entities related to site analytics - 5a88-0934-4699-8d1c
/* Gets the Top N entites related based on site analytics for a given kpi type. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!submit-request-for-top-n-entities-related-to-site-analytics
*/
func (s *SitesService) SubmitRequestForTopNEntitiesRelatedToSiteAnalytics(requestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics *RequestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics) (*ResponseSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/topNAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics).
		SetResult(&ResponseSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitRequestForTopNEntitiesRelatedToSiteAnalytics(requestSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics)
		}

		return nil, response, fmt.Errorf("error with operation SubmitRequestForTopNEntitiesRelatedToSiteAnalytics")
	}

	result := response.Result().(*ResponseSitesSubmitRequestForTopNEntitiesRelatedToSiteAnalytics)
	return result, response, err

}

//SubmitRequestForSiteAnalyticsTrendData Submit request for site analytics trend data - 2680-198e-477b-9226
/* Submits the task to get site analytics trend data for a given site. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!submit-request-for-site-analytics-trend-data
*/
func (s *SitesService) SubmitRequestForSiteAnalyticsTrendData(requestSitesSubmitRequestForSiteAnalyticsTrendData *RequestSitesSubmitRequestForSiteAnalyticsTrendData) (*ResponseSitesSubmitRequestForSiteAnalyticsTrendData, *resty.Response, error) {
	path := "/dna/data/api/v1/siteKpiSummaries/trendAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSitesSubmitRequestForSiteAnalyticsTrendData).
		SetResult(&ResponseSitesSubmitRequestForSiteAnalyticsTrendData{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SubmitRequestForSiteAnalyticsTrendData(requestSitesSubmitRequestForSiteAnalyticsTrendData)
		}

		return nil, response, fmt.Errorf("error with operation SubmitRequestForSiteAnalyticsTrendData")
	}

	result := response.Result().(*ResponseSitesSubmitRequestForSiteAnalyticsTrendData)
	return result, response, err

}

//AssignDevicesToSite Assign Devices To Site - 98a8-aa5e-40cb-b90b
/* Assigns unassigned devices to a site. This API does not move assigned devices to other sites.


@param siteID siteId path parameter. Site Id where device(s) needs to be assigned

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

		if AssignDevicesToSiteHeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", AssignDevicesToSiteHeaderParams.Timeout)
		}

		if AssignDevicesToSiteHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", AssignDevicesToSiteHeaderParams.Persistbapioutput)
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

//ExportMapArchive Export Map Archive - fb9f-2948-493b-bc68
/* Allows exporting a Map archive in an XML interchange format along with the associated images.


@param siteHierarchyUUID siteHierarchyUuid path parameter. The site hierarchy element UUID to export, all child elements starting at this hierarchy element will be included. Limited to a hierarchy that contains 500 or fewer maps.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-map-archive
*/
func (s *SitesService) ExportMapArchive(siteHierarchyUUID string) (*ResponseSitesExportMapArchive, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/export/{siteHierarchyUuid}"
	path = strings.Replace(path, "{siteHierarchyUuid}", fmt.Sprintf("%v", siteHierarchyUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSitesExportMapArchive{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportMapArchive(siteHierarchyUUID)
		}

		return nil, response, fmt.Errorf("error with operation ExportMapArchive")
	}

	result := response.Result().(*ResponseSitesExportMapArchive)
	return result, response, err

}

//ImportMapArchiveStartImport Import Map Archive - Start Import - b485-8b4e-4f6b-8409
/* Initiates a map archive import of a tar.gz file.  The archive must consist of one xmlDir/MapsImportExport.xml map descriptor file, and 1 or more images for the map areas nested under /images folder.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-start-import
*/
func (s *SitesService) ImportMapArchiveStartImport() (*resty.Response, error) {
	path := "/dna/intent/api/v1/maps/import/start"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchiveStartImport()
		}

		return response, fmt.Errorf("error with operation ImportMapArchiveStartImport")
	}

	return response, err

}

//ImportMapArchivePerformImport Import Map Archive - Perform Import - 5a9d-db37-4c18-9f6f
/* For a previously initatied import, approves the import to be performed, accepting that data loss may occur.  A Map import will fully replace existing Maps data for the site(s) defined in the archive. The Map Archive Import Status API /maps/import/${contextUuid}/status should always be checked to validate the pre-import validation output prior to performing the import.


@param importContextUUID importContextUuid path parameter. The unique import context UUID given by a previous call of Start Import API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-perform-import
*/
func (s *SitesService) ImportMapArchivePerformImport(importContextUUID string) (*ResponseSitesImportMapArchivePerformImport, *resty.Response, error) {
	path := "/dna/intent/api/v1/maps/import/{importContextUuid}/perform"
	path = strings.Replace(path, "{importContextUuid}", fmt.Sprintf("%v", importContextUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseSitesImportMapArchivePerformImport{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchivePerformImport(importContextUUID)
		}

		return nil, response, fmt.Errorf("error with operation ImportMapArchivePerformImport")
	}

	result := response.Result().(ResponseSitesImportMapArchivePerformImport)

	return &result, response, err

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

//ImportMapArchiveCancelAnImport Import Map Archive - Cancel an Import - 52b0-d8a3-4159-a6b7
/* Cancels a previously initatied import, allowing the system to cleanup cached resources about that import data, and ensures the import cannot accidentally be performed / approved at a later time.


@param importContextUUID importContextUuid path parameter. The unique import context UUID given by a previous call to Start Import API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-map-archive-cancel-an-import
*/
func (s *SitesService) ImportMapArchiveCancelAnImport(importContextUUID string) (*ResponseSitesImportMapArchiveCancelAnImport, *resty.Response, error) {
	//importContextUUID string
	path := "/dna/intent/api/v1/maps/import/{importContextUuid}"
	path = strings.Replace(path, "{importContextUuid}", fmt.Sprintf("%v", importContextUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		// SetResult(&ResponseSitesImportMapArchiveCancelAnImport{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportMapArchiveCancelAnImport(
				importContextUUID)
		}
		return nil, response, fmt.Errorf("error with operation ImportMapArchiveCancelAnImport")
	}

	result := response.Result().(ResponseSitesImportMapArchiveCancelAnImport)

	return &result, response, err

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
			return s.DeleteSite(
				siteID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSite")
	}

	result := response.Result().(*ResponseSitesDeleteSite)
	return result, response, err

}

func (s *SitesService) GetArea(GetSiteQueryParams *GetSiteQueryParams) (*ResponseSitesGetArea, *resty.Response, error) {
	path := "/dna/intent/api/v1/site"

	queryString, _ := query.Values(GetSiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetArea{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetArea(GetSiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSite")
	}

	result := response.Result().(*ResponseSitesGetArea)
	return result, response, err

}

func (s *SitesService) GetFloor(GetSiteQueryParams *GetSiteQueryParams) (*ResponseSitesGetFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/site"

	queryString, _ := query.Values(GetSiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSitesGetFloor{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFloor(GetSiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSite")
	}

	result := response.Result().(*ResponseSitesGetFloor)
	return result, response, err

}
