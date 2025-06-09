package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IssuesService service

type GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams struct {
	StartTime              float64 `url:"startTime,omitempty"`              //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                float64 `url:"endTime,omitempty"`                //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                  float64 `url:"limit,omitempty"`                  //Maximum number of issues to return
	Offset                 float64 `url:"offset,omitempty"`                 //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                 string  `url:"sortBy,omitempty"`                 //
	Order                  string  `url:"order,omitempty"`                  //The sort order of the field ascending or descending.
	IsGlobal               bool    `url:"isGlobal,omitempty"`               //Global issues are those issues which impacts across many devices, sites. They are also displayed on Issue Dashboard in Catalyst Center UI. Non-Global issues are displayed only on Client 360 or Device 360 pages. If this flag is 'true', only global issues are returned. If it if 'false', all issues are returned.
	Priority               string  `url:"priority,omitempty"`               //Priority of the issue. Supports single priority and multiple priorities Examples: priority=P1 (single priority requested) priority=P1&priority=P2&priority=P3 (multiple priorities requested)
	Severity               string  `url:"severity,omitempty"`               //Severity of the issue. Supports single severity and multiple severities. Examples: severity=high (single severity requested) severity=high&severity=medium (multiple severities requested)
	Status                 string  `url:"status,omitempty"`                 //Status of the issue. Supports single status and multiple statuses. Examples: status=active (single status requested) status=active&status=resolved (multiple statuses requested)
	EntityType             string  `url:"entityType,omitempty"`             //Entity type of the issue. Supports single entity type and multiple entity types. Examples: entityType=networkDevice (single entity type requested) entityType=network device&entityType=client (multiple entity types requested)
	Category               string  `url:"category,omitempty"`               //Categories of the issue. Supports single category and multiple categories. Examples: category=availability (single status requested) category=availability&category=onboarding (multiple categories requested)
	DeviceType             string  `url:"deviceType,omitempty"`             //Device Type of the device to which this issue belongs to. Supports single device type and multiple device types. Examples: deviceType=wireless controller (single device type requested) deviceType=wireless controller&deviceType=core (multiple device types requested)
	Name                   string  `url:"name,omitempty"`                   //The name of the issue Examples: name=ap_down (single issue name requested) name=ap_down&name=wlc_monitor (multiple issue names requested) Issue names can be retrieved using the API - /data/api/v1/assuranceIssueConfigurations
	IssueID                string  `url:"issueId,omitempty"`                //UUID of the issue Examples: issueId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single issue id requested) issueId=e52aecfe-b142-4287-a587-11a16ba6dd26&issueId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple issue ids requested)
	EntityID               string  `url:"entityId,omitempty"`               //Id of the entity for which this issue belongs to. For example, it      could be mac address of AP or UUID of Sensor   example: 68:ca:e4:79:3f:20 4de02167-901b-43cf-8822-cffd3caa286f Examples: entityId=68:ca:e4:79:3f:20 (single entity id requested) entityId=68:ca:e4:79:3f:20&entityId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple entity ids requested)
	UpdatedBy              string  `url:"updatedBy,omitempty"`              //The user who last updated this issue. Examples: updatedBy=admin (single updatedBy requested) updatedBy=admin&updatedBy=john (multiple updatedBy requested)
	SiteHierarchy          string  `url:"siteHierarchy,omitempty"`          //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID        string  `url:"siteHierarchyId,omitempty"`        //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteName               string  `url:"siteName,omitempty"`               //The name of the site. (Ex. `FloorName`) This field supports wildcard asterisk (*) character search support. E.g. *San*, *San, San* Examples: `?siteName=building1` (single siteName requested) `?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)
	SiteID                 string  `url:"siteId,omitempty"`                 //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	FabricSiteID           string  `url:"fabricSiteId,omitempty"`           //The UUID of the fabric site. (Ex. "flooruuid") Examples: fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26,864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	FabricVnName           string  `url:"fabricVnName,omitempty"`           //The name of the fabric virtual network Examples: fabricVnName=name1 (single fabric virtual network name requested) fabricVnName=name1&fabricVnName=name2&fabricVnName=name3 (multiple fabric virtual network names requested)
	FabricTransitSiteID    string  `url:"fabricTransitSiteId,omitempty"`    //The UUID of the fabric transit site. (Ex. "flooruuid") Examples: fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26&fabricTransitSiteId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	NetworkDeviceID        string  `url:"networkDeviceId,omitempty"`        //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress string  `url:"networkDeviceIpAddress,omitempty"` //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	MacAddress             string  `url:"macAddress,omitempty"`             //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	View                   string  `url:"view,omitempty"`                   //The name of the View. Each view represents a specific data set. Please refer to the `IssuesView` Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned. | View Name | Included Attributes | | --- | --- | | `update` | updatedTime, updatedBy | | `site` | siteName, siteHierarchy, siteId, siteHierarchyId | Examples: `view=update` (single view requested) `view=update&view=site` (multiple views requested)
	Attribute              string  `url:"attribute,omitempty"`              //List of attributes related to the issue. If these are provided, then only those attributes will be part of response along with the default attributes. Please refer to the `IssuesResponseAttribute` Model for supported attributes. Examples: `attribute=deviceType` (single attribute requested) `attribute=deviceType&attribute=updatedBy` (multiple attributes requested)
	AiDriven               bool    `url:"aiDriven,omitempty"`               //Flag whether the issue is AI driven issue
	FabricDriven           bool    `url:"fabricDriven,omitempty"`           //Flag whether the issue is related to a Fabric site, a virtual network or a transit.
	FabricSiteDriven       bool    `url:"fabricSiteDriven,omitempty"`       //Flag whether the issue is Fabric site driven issue
	FabricVnDriven         bool    `url:"fabricVnDriven,omitempty"`         //Flag whether the issue is Fabric Virtual Network driven issue
	FabricTransitDriven    bool    `url:"fabricTransitDriven,omitempty"`    //Flag whether the issue is Fabric Transit driven issue
}
type GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams struct {
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams struct {
	StartTime              float64 `url:"startTime,omitempty"`              //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                float64 `url:"endTime,omitempty"`                //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	IsGlobal               bool    `url:"isGlobal,omitempty"`               //Global issues are those issues which impacts across many devices, sites. They are also displayed on Issue Dashboard in Catalyst Center UI. Non-Global issues are displayed only on Client 360 or Device 360 pages. If this flag is 'true', only global issues are returned. If it if 'false', all issues are returned.
	Priority               string  `url:"priority,omitempty"`               //Priority of the issue. Supports single priority and multiple priorities Examples: priority=P1 (single priority requested) priority=P1&priority=P2&priority=P3 (multiple priorities requested)
	Severity               string  `url:"severity,omitempty"`               //Severity of the issue. Supports single severity and multiple severities. Examples: severity=high (single severity requested) severity=high&severity=medium (multiple severities requested)
	Status                 string  `url:"status,omitempty"`                 //Status of the issue. Supports single status and multiple statuses. Examples: status=active (single status requested) status=active&status=resolved (multiple statuses requested)
	EntityType             string  `url:"entityType,omitempty"`             //Entity type of the issue. Supports single entity type and multiple entity types. Examples: entityType=networkDevice (single entity type requested) entityType=network device&entityType=client (multiple entity types requested)
	Category               string  `url:"category,omitempty"`               //Categories of the issue. Supports single category and multiple categories. Examples: category=availability (single status requested) category=availability&category=onboarding (multiple categories requested)
	DeviceType             string  `url:"deviceType,omitempty"`             //Device Type of the device to which this issue belongs to. Supports single device type and multiple device types. Examples: deviceType=wireless controller (single device type requested) deviceType=wireless controller&deviceType=core (multiple device types requested)
	Name                   string  `url:"name,omitempty"`                   //The name of the issue Examples: name=ap_down (single issue name requested) name=ap_down&name=wlc_monitor (multiple issue names requested) Issue names can be retrieved using the API - /data/api/v1/assuranceIssueConfigurations
	IssueID                string  `url:"issueId,omitempty"`                //UUID of the issue Examples: issueId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single issue id requested) issueId=e52aecfe-b142-4287-a587-11a16ba6dd26&issueId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple issue ids requested)
	EntityID               string  `url:"entityId,omitempty"`               //Id of the entity for which this issue belongs to. For example, it      could be mac address of AP or UUID of Sensor   example: 68:ca:e4:79:3f:20 4de02167-901b-43cf-8822-cffd3caa286f Examples: entityId=68:ca:e4:79:3f:20 (single entity id requested) entityId=68:ca:e4:79:3f:20&entityId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple entity ids requested)
	UpdatedBy              string  `url:"updatedBy,omitempty"`              //The user who last updated this issue. Examples: updatedBy=admin (single updatedBy requested) updatedBy=admin&updatedBy=john (multiple updatedBy requested)
	SiteHierarchy          string  `url:"siteHierarchy,omitempty"`          //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID        string  `url:"siteHierarchyId,omitempty"`        //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteName               string  `url:"siteName,omitempty"`               //The name of the site. (Ex. `FloorName`) This field supports wildcard asterisk (*) character search support. E.g. *San*, *San, San* Examples: `?siteName=building1` (single siteName requested) `?siteName=building1&siteName=building2&siteName=building3` (multiple siteNames requested)
	SiteID                 string  `url:"siteId,omitempty"`                 //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	FabricSiteID           string  `url:"fabricSiteId,omitempty"`           //The UUID of the fabric site. (Ex. "flooruuid") Examples: fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26,864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	FabricVnName           string  `url:"fabricVnName,omitempty"`           //The name of the fabric virtual network Examples: fabricVnName=name1 (single fabric virtual network name requested) fabricVnName=name1&fabricVnName=name2&fabricVnName=name3 (multiple fabric virtual network names requested)
	FabricTransitSiteID    string  `url:"fabricTransitSiteId,omitempty"`    //The UUID of the fabric transit site. (Ex. "flooruuid") Examples: fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26 (single id requested) fabricTransitSiteId=e52aecfe-b142-4287-a587-11a16ba6dd26&fabricTransitSiteId=864d0421-02c0-43a6-9c52-81cad45f66d8 (multiple ids requested)
	NetworkDeviceID        string  `url:"networkDeviceId,omitempty"`        //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress string  `url:"networkDeviceIpAddress,omitempty"` //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	MacAddress             string  `url:"macAddress,omitempty"`             //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	AiDriven               bool    `url:"aiDriven,omitempty"`               //Flag whether the issue is AI driven issue
	FabricDriven           bool    `url:"fabricDriven,omitempty"`           //Flag whether the issue is related to a Fabric site, a virtual network or a transit.
	FabricSiteDriven       bool    `url:"fabricSiteDriven,omitempty"`       //Flag whether the issue is Fabric site driven issue
	FabricVnDriven         bool    `url:"fabricVnDriven,omitempty"`         //Flag whether the issue is Fabric Virtual Network driven issue
	FabricTransitDriven    bool    `url:"fabricTransitDriven,omitempty"`    //Flag whether the issue is Fabric Transit driven issue
}
type GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetSummaryAnalyticsDataOfIssuesHeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTopNAnalyticsDataOfIssuesHeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTrendAnalyticsDataOfIssuesHeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue display name need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue display name is returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDQueryParams struct {
	View      string `url:"view,omitempty"`      //The name of the View. Each view represents a specific data set. Please refer to the `IssuesView` Model for supported views. View is predefined set of attributes supported by the API. Only the attributes related to the given view will be part of the API response along with default attributes. If multiple views are provided, then response will contain attributes from all those views. If no views are specified, all attributes will be returned. | View Name | Included Attributes | | --- | --- | | `update` | updatedTime, updatedBy | | `site` | siteName, siteHierarchy, siteId, siteHierarchyId | Examples: `view=update` (single view requested) `view=update&view=site` (multiple views requested)
	Attribute string `url:"attribute,omitempty"` //List of attributes related to the issue. If these are provided, then only those attributes will be part of response along with the default attributes. Please refer to the `IssuesResponseAttribute` Model for supported attributes. Examples: `attribute=deviceType` (single attribute requested) `attribute=deviceType&attribute=updatedBy` (multiple attributes requested)
}
type GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDHeaderParams struct {
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type IgnoreTheGivenListOfIssuesHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type ResolveTheGivenListsOfIssuesHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams struct {
	ContentType    string `url:"Content-Type,omitempty"`    //Expects type string. Request body content type
	AcceptLanguage string `url:"Accept-Language,omitempty"` //Expects type string. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are - 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
	XCaLLERID      string `url:"X-CALLER-ID,omitempty"`     //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type CreatesANewUserDefinedIssueDefinitionsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams struct {
	ID        string  `url:"id,omitempty"`        //The custom issue definition identifier and unique identifier across the profile.Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=19ca-4170-8375-b694e251101c-6bef213c (multiple Id request in the query param)
	ProfileID string  `url:"profileId,omitempty"` //The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
	Name      string  `url:"name,omitempty"`      //The list of UDI issue names
	Priority  string  `url:"priority,omitempty"`  //The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
	IsEnabled bool    `url:"isEnabled,omitempty"` //The enable status of the custom issue definition, either true or false.
	Severity  float64 `url:"severity,omitempty"`  //The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
	Facility  string  `url:"facility,omitempty"`  //The syslog facility name
	Mnemonic  string  `url:"mnemonic,omitempty"`  //The syslog mnemonic name
	Limit     float64 `url:"limit,omitempty"`     //The maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
}
type GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams struct {
	ID        string  `url:"id,omitempty"`        //The custom issue definition identifier and unique identifier across the profile. Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=19ca-4170-8375-b694e251101c-6bef213c (multiple Id request in the query param)
	ProfileID string  `url:"profileId,omitempty"` //The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
	Name      string  `url:"name,omitempty"`      //The list of UDI issue names. (Ex."TestUdiIssues")
	Priority  string  `url:"priority,omitempty"`  //The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
	IsEnabled bool    `url:"isEnabled,omitempty"` //The enable status of the custom issue definition, either true or false.
	Severity  float64 `url:"severity,omitempty"`  //The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
	Facility  string  `url:"facility,omitempty"`  //The syslog facility name
	Mnemonic  string  `url:"mnemonic,omitempty"`  //The syslog mnemonic name
}
type GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
}
type GetIssueEnrichmentDetailsHeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Issue enrichment details can be fetched based on either Issue ID or Client MAC address. This parameter value must either be issue_id/mac_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. For the enrichment details to be made available as part of the API response, this header must be set to true. This header must be explicitly passed when called from client applications outside Catalyst Center
}
type IssuesQueryParams struct {
	StartTime   float64 `url:"startTime,omitempty"`   //Starting epoch time in milliseconds of query time window
	EndTime     float64 `url:"endTime,omitempty"`     //Ending epoch time in milliseconds of query time window
	SiteID      string  `url:"siteId,omitempty"`      //Assurance UUID value of the site in the issue content
	DeviceID    string  `url:"deviceId,omitempty"`    //Assurance UUID value of the device in the issue content
	MacAddress  string  `url:"macAddress,omitempty"`  //Client's device MAC address of the issue (format xx:xx:xx:xx:xx:xx)
	Priority    string  `url:"priority,omitempty"`    //The issue's priority value: P1, P2, P3, or P4 (case insensitive) (Use only when macAddress and deviceId are not provided)
	IssueStatus string  `url:"issueStatus,omitempty"` //The issue's status value: ACTIVE, IGNORED, RESOLVED (case insensitive)
	AiDriven    string  `url:"aiDriven,omitempty"`    //The issue's AI driven value: YES or NO (case insensitive) (Use only when macAddress and deviceId are not provided)
}
type ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams struct {
	DeviceType   string  `url:"deviceType,omitempty"`   //These are the device families/types supported for system issue definitions. If no input is made on device type, all device types are considered.
	ProfileID    string  `url:"profileId,omitempty"`    //The profile identier to fetch the profile associated issue defintions. The default is `global`. Please refer Network design profiles documentation for more details.
	ID           string  `url:"id,omitempty"`           //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	Name         string  `url:"name,omitempty"`         //The list of system defined issue names. (Ex."BGP_Down") Examples: name=BGP_Down (single entity uuid requested) name=BGP_Down&name=BGP_Flap (multiple issue names separated by & operator)
	Priority     string  `url:"priority,omitempty"`     //Issue priority, possible values are P1, P2, P3, P4. `P1`: A critical issue that needs immediate attention and can have a wide impact on network operations. `P2`: A major issue that can potentially impact multiple devices or clients. `P3`: A minor issue that has a localized or minimal impact. `P4`: A warning issue that may not be an immediate problem but addressing it can optimize the network performance.
	IssueEnabled bool    `url:"issueEnabled,omitempty"` //The enablement status of the issue definition, either true or false.
	Attribute    string  `url:"attribute,omitempty"`    //These are the attributes supported in system issue definitions response. By default, all properties are sent in response.
	Offset       float64 `url:"offset,omitempty"`       //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit        float64 `url:"limit,omitempty"`        //Maximum number of records to return
	SortBy       string  `url:"sortBy,omitempty"`       //A field within the response to sort by.
	Order        string  `url:"order,omitempty"`        //The sort order of the field ascending or descending.
}
type ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams struct {
	DeviceType   string `url:"deviceType,omitempty"`   //These are the device families/types supported for system issue definitions. If no input is made on device type, all device types are considered.
	ProfileID    string `url:"profileId,omitempty"`    //The profile identier to fetch the profile associated issue defintions. The default is `global`. Please refer Network design profiles documentation for more details.
	ID           string `url:"id,omitempty"`           //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	Name         string `url:"name,omitempty"`         //The list of system defined issue names. (Ex."BGP_Down") Examples: name=BGP_Down (single entity uuid requested) name=BGP_Down&name=BGP_Flap (multiple issue names separated by & operator)
	Priority     string `url:"priority,omitempty"`     //Issue priority, possible values are P1, P2, P3, P4. `P1`: A critical issue that needs immediate attention and can have a wide impact on network operations. `P2`: A major issue that can potentially impact multiple devices or clients. `P3`: A minor issue that has a localized or minimal impact. `P4`: A warning issue that may not be an immediate problem but addressing it can optimize the network performance.
	IssueEnabled bool   `url:"issueEnabled,omitempty"` //The enablement status of the issue definition, either true or false.
}
type GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetIssueTriggerDefinitionForGivenIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}

type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork struct {
	Response *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponse `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkPage       `json:"page,omitempty"`     //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponse struct {
	IssueID                string                                                                                                `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                                `json:"name,omitempty"`                   // Name
	Description            string                                                                                                `json:"description,omitempty"`            // Description
	Summary                string                                                                                                `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                                `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                                `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                                `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                                `json:"category,omitempty"`               // Category
	EntityType             string                                                                                                `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                                `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                                  `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                                  `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                                `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                                 `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseUpdatedBy              `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseUpdatedTime            `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseNotes                  `json:"notes,omitempty"`                  // Notes
	SiteID                 *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteID                 `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteHierarchyID        `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteName               `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteHierarchy          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseUpdatedBy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseUpdatedTime interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseNotes interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteHierarchyID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteName interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSiteHierarchy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSuggestedActions struct {
	Message string                                                                                                 `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseSuggestedActionsSteps interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkPage struct {
	Limit  *int                                                                                `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                `json:"offset,omitempty"` // Offset
	Count  *int                                                                                `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork struct {
	Response *ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkResponse `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFilters struct {
	Response *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersPage       `json:"page,omitempty"`     //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponse struct {
	IssueID                string                                                                                 `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                 `json:"name,omitempty"`                   // Name
	Description            string                                                                                 `json:"description,omitempty"`            // Description
	Summary                string                                                                                 `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                 `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                 `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                 `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                 `json:"category,omitempty"`               // Category
	EntityType             string                                                                                 `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                 `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                   `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                   `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                 `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                  `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseUpdatedBy              `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseUpdatedTime            `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseNotes                  `json:"notes,omitempty"`                  // Notes
	SiteID                 *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteID                 `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteHierarchyID        `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteName               `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          *ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteHierarchy          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseUpdatedBy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseUpdatedTime interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseNotes interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteHierarchyID interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteName interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSiteHierarchy interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSuggestedActions struct {
	Message string                                                                                  `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseSuggestedActionsSteps interface{}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersPage struct {
	Limit  *int                                                                 `json:"limit,omitempty"`  // Limit
	Offset *int                                                                 `json:"offset,omitempty"` // Offset
	Count  *int                                                                 `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters struct {
	Filters *[]ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersFilters `json:"filters,omitempty"` //
}
type ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Value    string `json:"value,omitempty"`    // Value
	Operator string `json:"operator,omitempty"` // Operator
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssues struct {
	Version  string                                                 `json:"version,omitempty"`  // Version
	Response *ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponse `json:"response,omitempty"` //
	Page     *ResponseIssuesGetSummaryAnalyticsDataOfIssuesPage     `json:"page,omitempty"`     //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponse struct {
	Groups              *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseGroups struct {
	ID                  string                                                                            `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesPage struct {
	Limit  *int                                                       `json:"limit,omitempty"`  // Limit
	Offset *int                                                       `json:"offset,omitempty"` // Offset
	Count  *int                                                       `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetSummaryAnalyticsDataOfIssuesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesPageSortBy struct {
	Name     string                                                           `json:"name,omitempty"`     // Name
	Function *ResponseIssuesGetSummaryAnalyticsDataOfIssuesPageSortByFunction `json:"function,omitempty"` // Function
	Order    string                                                           `json:"order,omitempty"`    // Order
}
type ResponseIssuesGetSummaryAnalyticsDataOfIssuesPageSortByFunction interface{}
type ResponseIssuesGetTopNAnalyticsDataOfIssues struct {
	Version  string                                                `json:"version,omitempty"`  // Version
	Response *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesResponse `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTopNAnalyticsDataOfIssuesPage       `json:"page,omitempty"`     //
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesResponse struct {
	ID                  string                                                                   `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesPage struct {
	Limit  *int                                                    `json:"limit,omitempty"`  // Limit
	Offset *int                                                    `json:"offset,omitempty"` // Offset
	Count  *int                                                    `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetTopNAnalyticsDataOfIssuesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesPageSortBy struct {
	Name     string                                                        `json:"name,omitempty"`     // Name
	Function *ResponseIssuesGetTopNAnalyticsDataOfIssuesPageSortByFunction `json:"function,omitempty"` // Function
	Order    string                                                        `json:"order,omitempty"`    // Order
}
type ResponseIssuesGetTopNAnalyticsDataOfIssuesPageSortByFunction interface{}
type ResponseIssuesGetTrendAnalyticsDataOfIssues struct {
	Version  string                                                 `json:"version,omitempty"`  // Version
	Response *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesResponse `json:"response,omitempty"` //
	Page     *ResponseIssuesGetTrendAnalyticsDataOfIssuesPage       `json:"page,omitempty"`     //
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesResponse struct {
	Timestamp           *int                                                                      `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseGroups struct {
	ID                  string                                                                          `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseIssuesGetTrendAnalyticsDataOfIssuesPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID struct {
	Response *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponse `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponse struct {
	IssueID                string                                                                                                      `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                                      `json:"name,omitempty"`                   // Name
	Description            string                                                                                                      `json:"description,omitempty"`            // Description
	Summary                string                                                                                                      `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                                      `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                                      `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                                      `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                                      `json:"category,omitempty"`               // Category
	EntityType             string                                                                                                      `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                                      `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                                        `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                                        `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                                      `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                                       `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              string                                                                                                      `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *int                                                                                                        `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseNotes                  `json:"notes,omitempty"`                  // Notes
	SiteID                 *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteID                 `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteHierarchyID        `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteName               `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          *ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteHierarchy          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseNotes interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteID interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteHierarchyID interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteName interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSiteHierarchy interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSuggestedActions struct {
	Message string                                                                                                       `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseSuggestedActionsSteps interface{}
type ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesIgnoreTheGivenListOfIssues struct {
	Response *ResponseIssuesIgnoreTheGivenListOfIssuesResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseIssuesIgnoreTheGivenListOfIssuesResponse struct {
	SuccessfulIssueIDs []string `json:"successfulIssueIds,omitempty"` // Successful Issue Ids
	FailureIssueIDs    []string `json:"failureIssueIds,omitempty"`    // Failure Issue Ids
}
type ResponseIssuesResolveTheGivenListsOfIssues struct {
	Response *ResponseIssuesResolveTheGivenListsOfIssuesResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseIssuesResolveTheGivenListsOfIssuesResponse struct {
	SuccessfulIssueIDs []string `json:"successfulIssueIds,omitempty"` // Successful Issue Ids
	FailureIssueIDs    []string `json:"failureIssueIds,omitempty"`    // Failure Issue Ids
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFields struct {
	Response *ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponse `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponse struct {
	IssueID                string                                                                                   `json:"issueId,omitempty"`                // Issue Id
	Name                   string                                                                                   `json:"name,omitempty"`                   // Name
	Description            string                                                                                   `json:"description,omitempty"`            // Description
	Summary                string                                                                                   `json:"summary,omitempty"`                // Summary
	Priority               string                                                                                   `json:"priority,omitempty"`               // Priority
	Severity               string                                                                                   `json:"severity,omitempty"`               // Severity
	DeviceType             string                                                                                   `json:"deviceType,omitempty"`             // Device Type
	Category               string                                                                                   `json:"category,omitempty"`               // Category
	EntityType             string                                                                                   `json:"entityType,omitempty"`             // Entity Type
	EntityID               string                                                                                   `json:"entityId,omitempty"`               // Entity Id
	FirstOccurredTime      *int                                                                                     `json:"firstOccurredTime,omitempty"`      // First Occurred Time
	MostRecentOccurredTime *int                                                                                     `json:"mostRecentOccurredTime,omitempty"` // Most Recent Occurred Time
	Status                 string                                                                                   `json:"status,omitempty"`                 // Status
	IsGlobal               *bool                                                                                    `json:"isGlobal,omitempty"`               // Is Global
	UpdatedBy              string                                                                                   `json:"updatedBy,omitempty"`              // Updated By
	UpdatedTime            *int                                                                                     `json:"updatedTime,omitempty"`            // Updated Time
	Notes                  string                                                                                   `json:"notes,omitempty"`                  // Notes
	SiteID                 string                                                                                   `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchyID        string                                                                                   `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteName               string                                                                                   `json:"siteName,omitempty"`               // Site Name
	SiteHierarchy          string                                                                                   `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SuggestedActions       *[]ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponseSuggestedActions     `json:"suggestedActions,omitempty"`       //
	AdditionalAttributes   *[]ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponseAdditionalAttributes `json:"additionalAttributes,omitempty"`   //
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponseSuggestedActions struct {
	Message string `json:"message,omitempty"` // Message
}
type ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFieldsResponseAdditionalAttributes struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseIssuesCreatesANewUserDefinedIssueDefinitions struct {
	Response *ResponseIssuesCreatesANewUserDefinedIssueDefinitionsResponse `json:"response,omitempty"` //
}
type ResponseIssuesCreatesANewUserDefinedIssueDefinitionsResponse struct {
	ID                    string                                                               `json:"id,omitempty"`                    // Id
	Name                  string                                                               `json:"name,omitempty"`                  // Name
	Description           string                                                               `json:"description,omitempty"`           // Description
	ProfileID             string                                                               `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                               `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesCreatesANewUserDefinedIssueDefinitionsResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                               `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                 `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                 `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesCreatesANewUserDefinedIssueDefinitionsResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters struct {
	Response *[]ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersPage       `json:"page,omitempty"`     //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse struct {
	ID                    string                                                                              `json:"id,omitempty"`                    // Id
	Name                  string                                                                              `json:"name,omitempty"`                  // Name
	Description           string                                                                              `json:"description,omitempty"`           // Description
	ProfileID             string                                                                              `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                                              `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                               `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                              `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                               `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                               `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                                `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                                `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersPage struct {
	Limit  *int                                                                             `json:"limit,omitempty"`  // Limit
	Offset *int                                                                             `json:"offset,omitempty"` // Offset
	Count  *int                                                                             `json:"count,omitempty"`  // Count
	SortBy *[]ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters struct {
	Response *ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID struct {
	Response *ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponse `json:"response,omitempty"` //
}
type ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponse struct {
	ID                    string                                                                                      `json:"id,omitempty"`                    // Id
	Name                  string                                                                                      `json:"name,omitempty"`                  // Name
	Description           string                                                                                      `json:"description,omitempty"`           // Description
	ProfileID             string                                                                                      `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                                                      `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                                       `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                                      `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                                       `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                                       `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                                        `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                                        `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID struct {
	Response *ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDResponse `json:"response,omitempty"` //
}
type ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDResponse struct {
	ID                    string                                                                                   `json:"id,omitempty"`                    // Id
	Name                  string                                                                                   `json:"name,omitempty"`                  // Name
	Description           string                                                                                   `json:"description,omitempty"`           // Description
	ProfileID             string                                                                                   `json:"profileId,omitempty"`             // Profile Id
	TriggerID             string                                                                                   `json:"triggerId,omitempty"`             // Trigger Id
	Rules                 *[]ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDResponseRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                                    `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                                   `json:"priority,omitempty"`              // Priority
	IsDeletable           *bool                                                                                    `json:"isDeletable,omitempty"`           // Is Deletable
	IsNotificationEnabled *bool                                                                                    `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
	CreatedTime           *int                                                                                     `json:"createdTime,omitempty"`           // Created Time
	LastUpdatedTime       *int                                                                                     `json:"lastUpdatedTime,omitempty"`       // Last Updated Time
}
type ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDResponseRules struct {
	Type              string `json:"type,omitempty"`              // Type
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type ResponseIssuesExecuteSuggestedActionsCommands []ResponseItemIssuesExecuteSuggestedActionsCommands // Array of ResponseIssuesExecuteSuggestedActionsCommands
type ResponseItemIssuesExecuteSuggestedActionsCommands struct {
	ActionInfo       string                                                          `json:"actionInfo,omitempty"`       // Actions Info
	StepsCount       *int                                                            `json:"stepsCount,omitempty"`       // Steps Count
	EntityID         string                                                          `json:"entityId,omitempty"`         // Entity Id
	Hostname         string                                                          `json:"hostname,omitempty"`         // Hostname
	StepsDescription string                                                          `json:"stepsDescription,omitempty"` // Steps Description
	Command          string                                                          `json:"command,omitempty"`          // Command
	CommandOutput    *ResponseItemIssuesExecuteSuggestedActionsCommandsCommandOutput `json:"commandOutput,omitempty"`    // Command Output
}
type ResponseItemIssuesExecuteSuggestedActionsCommandsCommandOutput interface{}
type ResponseIssuesGetIssueEnrichmentDetails struct {
	IssueDetails *ResponseIssuesGetIssueEnrichmentDetailsIssueDetails `json:"issueDetails,omitempty"` //
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetails struct {
	Issue *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssue `json:"issue,omitempty"` //
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssue struct {
	IssueID          string                                                                      `json:"issueId,omitempty"`          // Issue Id
	IssueSource      string                                                                      `json:"issueSource,omitempty"`      // Issue Source
	IssueCategory    string                                                                      `json:"issueCategory,omitempty"`    // Issue Category
	IssueName        string                                                                      `json:"issueName,omitempty"`        // Issue Name
	IssueDescription string                                                                      `json:"issueDescription,omitempty"` // Issue Description
	IssueEntity      string                                                                      `json:"issueEntity,omitempty"`      // Issue Entity
	IssueEntityValue string                                                                      `json:"issueEntityValue,omitempty"` // Issue Entity Value
	IssueSeverity    string                                                                      `json:"issueSeverity,omitempty"`    // Issue Severity
	IssuePriority    string                                                                      `json:"issuePriority,omitempty"`    // Issue Priority
	IssueSummary     string                                                                      `json:"issueSummary,omitempty"`     // Issue Summary
	IssueTimestamp   *int                                                                        `json:"issueTimestamp,omitempty"`   // Issue Timestamp
	SuggestedActions *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
	ImpactedHosts    *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    // Impacted Hosts
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActions struct {
	Message string                                                                           `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps interface{}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueImpactedHosts interface{}
type ResponseIssuesIssues struct {
	Version    string                          `json:"version,omitempty"`    // Response body's schema version string
	TotalCount string                          `json:"totalCount,omitempty"` // Total number of issues in the query time window
	Response   *[]ResponseIssuesIssuesResponse `json:"response,omitempty"`   //
}
type ResponseIssuesIssuesResponse struct {
	IssueID             string `json:"issueId,omitempty"`               // The issue's unique identifier
	Name                string `json:"name,omitempty"`                  // The issue's display name
	SiteID              string `json:"siteId,omitempty"`                // The site UUID where the issue occurred
	DeviceID            string `json:"deviceId,omitempty"`              // The device UUID where the issue occurred
	DeviceRole          string `json:"deviceRole,omitempty"`            // The device role
	AiDriven            string `json:"aiDriven,omitempty"`              // Whether the issue is AI driven ('Yes' or 'No')
	ClientMac           string `json:"clientMac,omitempty"`             // The client MAC address related to this issue
	IssueOccurenceCount *int   `json:"issue_occurence_count,omitempty"` // Total number of instances of this issue in the query time window
	Status              string `json:"status,omitempty"`                // The status of the issue
	Priority            string `json:"priority,omitempty"`              // Priority setting of the issue
	Category            string `json:"category,omitempty"`              // Category of the issue
	LastOccurenceTime   *int   `json:"last_occurence_time,omitempty"`   // The UTC timestamp of last occurence of this issue
}
type ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFilters struct {
	Response *[]ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersResponse `json:"response,omitempty"` //
}
type ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFiltersResponse struct {
	ID                           string   `json:"id,omitempty"`                           // Id
	Name                         string   `json:"name,omitempty"`                         // Name
	DisplayName                  string   `json:"displayName,omitempty"`                  // Display Name
	Description                  string   `json:"description,omitempty"`                  // Description
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	DefaultPriority              string   `json:"defaultPriority,omitempty"`              // Default Priority
	DeviceType                   string   `json:"deviceType,omitempty"`                   // Device Type
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ProfileID                    string   `json:"profileId,omitempty"`                    // Profile Id
	DefinitionStatus             string   `json:"definitionStatus,omitempty"`             // Definition Status
	CategoryName                 string   `json:"categoryName,omitempty"`                 // Category Name
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
	LastModified                 string   `json:"lastModified,omitempty"`                 // Last Modified
}
type ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters struct {
	Response *ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseIssuesGetIssueTriggerDefinitionForGivenID struct {
	Response *ResponseIssuesGetIssueTriggerDefinitionForGivenIDResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // Version
}
type ResponseIssuesGetIssueTriggerDefinitionForGivenIDResponse struct {
	ID                           string   `json:"id,omitempty"`                           // Id
	Name                         string   `json:"name,omitempty"`                         // Name
	DisplayName                  string   `json:"displayName,omitempty"`                  // Display Name
	Description                  string   `json:"description,omitempty"`                  // Description
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	DefaultPriority              string   `json:"defaultPriority,omitempty"`              // Default Priority
	DeviceType                   string   `json:"deviceType,omitempty"`                   // Device Type
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ProfileID                    string   `json:"profileId,omitempty"`                    // Profile Id
	DefinitionStatus             string   `json:"definitionStatus,omitempty"`             // Definition Status
	CategoryName                 string   `json:"categoryName,omitempty"`                 // Category Name
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
	LastModified                 string   `json:"lastModified,omitempty"`                 // Last Modified
}
type ResponseIssuesIssueTriggerDefinitionUpdate struct {
	Response *ResponseIssuesIssueTriggerDefinitionUpdateResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseIssuesIssueTriggerDefinitionUpdateResponse struct {
	ID                           string   `json:"id,omitempty"`                           // Id
	Name                         string   `json:"name,omitempty"`                         // Name
	DisplayName                  string   `json:"displayName,omitempty"`                  // Display Name
	Description                  string   `json:"description,omitempty"`                  // Description
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	DefaultPriority              string   `json:"defaultPriority,omitempty"`              // Default Priority
	DeviceType                   string   `json:"deviceType,omitempty"`                   // Device Type
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ProfileID                    string   `json:"profileId,omitempty"`                    // Profile Id
	DefinitionStatus             string   `json:"definitionStatus,omitempty"`             // Definition Status
	CategoryName                 string   `json:"categoryName,omitempty"`                 // Category Name
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
	LastModified                 string   `json:"lastModified,omitempty"`                 // Last Modified
}
type RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters struct {
	StartTime *int                                                             `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                             `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFilters `json:"filters,omitempty"`   //
}
type RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFilters struct {
	Key             string                                                                  `json:"key,omitempty"`             // Key
	Operator        string                                                                  `json:"operator,omitempty"`        // Operator
	Value           string                                                                  `json:"value,omitempty"`           // Value
	LogicalOperator string                                                                  `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters struct {
	StartTime *int                                                                 `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                                 `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersFilters `json:"filters,omitempty"`   //
}
type RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersFilters struct {
	Key             string                                                                      `json:"key,omitempty"`             // Key
	Operator        string                                                                      `json:"operator,omitempty"`        // Operator
	Value           string                                                                      `json:"value,omitempty"`           // Value
	LogicalOperator string                                                                      `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersFiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetSummaryAnalyticsDataOfIssues struct {
	StartTime           *int                                                               `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                               `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesFilters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                           `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                           `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestIssuesGetSummaryAnalyticsDataOfIssuesPage                  `json:"page,omitempty"`                //
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesFilters struct {
	Key             string                                                        `json:"key,omitempty"`             // Key
	Operator        string                                                        `json:"operator,omitempty"`        // Operator
	Value           string                                                        `json:"value,omitempty"`           // Value
	LogicalOperator string                                                        `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesFiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesPage struct {
	Limit  *int                                                      `json:"limit,omitempty"`  // Limit
	Offset *int                                                      `json:"offset,omitempty"` // Offset
	SortBy *[]RequestIssuesGetSummaryAnalyticsDataOfIssuesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestIssuesGetSummaryAnalyticsDataOfIssuesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestIssuesGetTopNAnalyticsDataOfIssues struct {
	StartTime           *int                                                            `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                            `json:"endTime,omitempty"`             // End Time
	TopN                *int                                                            `json:"topN,omitempty"`                // Top N
	Filters             *[]RequestIssuesGetTopNAnalyticsDataOfIssuesFilters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                        `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                        `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestIssuesGetTopNAnalyticsDataOfIssuesAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestIssuesGetTopNAnalyticsDataOfIssuesPage                  `json:"page,omitempty"`                //
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesFilters struct {
	Key             string                                                     `json:"key,omitempty"`             // Key
	Operator        string                                                     `json:"operator,omitempty"`        // Operator
	Value           string                                                     `json:"value,omitempty"`           // Value
	LogicalOperator string                                                     `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestIssuesGetTopNAnalyticsDataOfIssuesFiltersFilters `json:"filters,omitempty"`         //
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    string `json:"value,omitempty"`    // Value
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesPage struct {
	Limit  *int                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                   `json:"offset,omitempty"` // Offset
	SortBy *[]RequestIssuesGetTopNAnalyticsDataOfIssuesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestIssuesGetTopNAnalyticsDataOfIssuesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestIssuesGetTrendAnalyticsDataOfIssues struct {
	StartTime           *int                                                             `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                             `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                           `json:"trendInterval,omitempty"`       // Trend Interval
	Filters             *[]RequestIssuesGetTrendAnalyticsDataOfIssuesFilters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                         `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                         `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestIssuesGetTrendAnalyticsDataOfIssuesAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestIssuesGetTrendAnalyticsDataOfIssuesPage                  `json:"page,omitempty"`                //
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Value    string `json:"value,omitempty"`    // Value
	Operator string `json:"operator,omitempty"` // Operator
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestIssuesGetTrendAnalyticsDataOfIssuesPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestIssuesIgnoreTheGivenListOfIssues struct {
	IssueIDs    []string `json:"issueIds,omitempty"`    // Issue Ids
	IgnoreHours *int     `json:"ignoreHours,omitempty"` // Ignore Hours
}
type RequestIssuesResolveTheGivenListsOfIssues struct {
	IssueIDs []string `json:"issueIds,omitempty"` // Issue Ids
}
type RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFields struct {
	Notes string `json:"notes,omitempty"` // Notes
}
type RequestIssuesCreatesANewUserDefinedIssueDefinitions struct {
	Name                  string                                                      `json:"name,omitempty"`                  // Name
	Description           string                                                      `json:"description,omitempty"`           // Description
	Rules                 *[]RequestIssuesCreatesANewUserDefinedIssueDefinitionsRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                       `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                      `json:"priority,omitempty"`              // Priority
	IsNotificationEnabled *bool                                                       `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
}
type RequestIssuesCreatesANewUserDefinedIssueDefinitionsRules struct {
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID struct {
	Name                  string                                                                          `json:"name,omitempty"`                  // Name
	Description           string                                                                          `json:"description,omitempty"`           // Description
	Rules                 *[]RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules `json:"rules,omitempty"`                 //
	IsEnabled             *bool                                                                           `json:"isEnabled,omitempty"`             // Is Enabled
	Priority              string                                                                          `json:"priority,omitempty"`              // Priority
	IsNotificationEnabled *bool                                                                           `json:"isNotificationEnabled,omitempty"` // Is Notification Enabled
}
type RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedIDRules struct {
	Severity          *int   `json:"severity,omitempty"`          // Severity
	Facility          string `json:"facility,omitempty"`          // Facility
	Mnemonic          string `json:"mnemonic,omitempty"`          // Mnemonic
	Pattern           string `json:"pattern,omitempty"`           // Pattern
	Occurrences       *int   `json:"occurrences,omitempty"`       // Occurrences
	DurationInMinutes *int   `json:"durationInMinutes,omitempty"` // Duration In Minutes
}
type RequestIssuesExecuteSuggestedActionsCommands struct {
	EntityType  string `json:"entity_type,omitempty"`  // Commands provided as part of the suggested actions for an issue can be executed based on issue id. The value here must be issue_id
	EntityValue string `json:"entity_value,omitempty"` // Contains the actual value for the entity type that has been defined
}
type RequestIssuesIssueTriggerDefinitionUpdate struct {
	SynchronizeToHealthThreshold *bool    `json:"synchronizeToHealthThreshold,omitempty"` // Synchronize To Health Threshold
	Priority                     string   `json:"priority,omitempty"`                     // Priority
	IssueEnabled                 *bool    `json:"issueEnabled,omitempty"`                 // Issue Enabled
	ThresholdValue               *float64 `json:"thresholdValue,omitempty"`               // Threshold Value
}

//GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork Get the details of issues for given set of filters - a991-6985-476b-b271
/* Returns all details of each issue along with suggested actions for given set of filters specified in query parameters. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. All string type query parameters support wildcard search (using *). For example: siteHierarchy=Global/San Jose/* returns issues under all sites whole siteHierarchy starts with "Global/San Jose/". https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams Custom header parameters
@param GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-issues-for-given-set-of-filters-know-your-network
*/
func (s *IssuesService) GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams *GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams, GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams *GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams) (*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues"

	queryString, _ := query.Values(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams != nil {

		if GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams.AcceptLanguage)
		}

		if GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork(GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams, GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork")
	}

	result := response.Result().(*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersKnowYourNetwork)
	return result, response, err

}

//GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork Get the total number of issues for given set of filters - 049b-c87d-456a-a69b
/* Returns the total number issues for given set of filters. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams Custom header parameters
@param GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-number-of-issues-for-given-set-of-filters-know-your-network
*/
func (s *IssuesService) GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams, GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams) (*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/count"

	queryString, _ := query.Values(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams != nil {

		if GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork(GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkHeaderParams, GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork")
	}

	result := response.Result().(*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFiltersKnowYourNetwork)
	return result, response, err

}

//GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID Get all the details and suggested actions of an issue for the given issue id - 82ae-1acd-4b6a-ab00
/* Returns all the details and suggested actions of an issue for the given issue id. https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param id id path parameter. The issue Uuid

@param GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams Custom header parameters
@param GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-the-details-and-suggested-actions-of-an-issue-for-the-given-issue-id
*/
func (s *IssuesService) GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID(id string, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams *GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDHeaderParams, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdQueryParams *GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIDQueryParams) (*ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams != nil {

		if GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams.AcceptLanguage)
		}

		if GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID(id, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdHeaderParams, GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueId")
	}

	result := response.Result().(*ResponseIssuesGetAllTheDetailsAndSuggestedActionsOfAnIssueForTheGivenIssueID)
	return result, response, err

}

//GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters Get all the custom issue definitions based on the given filters. - 1bb9-bb87-4efa-afd2
/* Retrieve the existing syslog-based custom issue definitions. The supported filters are id, name, profileId,  definition enable status, priority, severity, facility and mnemonic. The issue definition configurations may vary across profiles, hence specifying the profile Id in the query parameter is important and the default profile is global.


  The assurance profile definitions can be obtain via the API endpoint: /api/v1/siteprofile?namespace=assurance. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-the-custom-issue-definitions-based-on-the-given-filters
*/
func (s *IssuesService) GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams *GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams) (*ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions"

	queryString, _ := query.Values(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters")
	}

	result := response.Result().(*ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters)
	return result, response, err

}

//GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters Get the total custom issue definitions count based on the provided filters. - 9b91-2a4a-4d1a-9595
/* Get the total number of Custom issue definitions count based on the provided filters. The supported filters are id, name, profileId and definition enable status, severity, facility and mnemonic. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams Custom header parameters
@param GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-custom-issue-definitions-count-based-on-the-provided-filters
*/
func (s *IssuesService) GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams *GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams, GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams *GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams) (*ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions/count"

	queryString, _ := query.Values(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams != nil {

		if GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters(GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersHeaderParams, GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters")
	}

	result := response.Result().(*ResponseIssuesGetTheTotalCustomIssueDefinitionsCountBasedOnTheProvidedFilters)
	return result, response, err

}

//GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID Get the custom issue definition for the given custom issue definition Id. - d39f-a9d8-44b8-880d
/* Get the custom issue definition for the given custom issue definition Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param id id path parameter. Get the custom issue definition for the given custom issue definition Id.

@param GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-custom-issue-definition-for-the-given-custom-issue-definition-id
*/
func (s *IssuesService) GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID(id string, GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdHeaderParams *GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDHeaderParams) (*ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdHeaderParams != nil {

		if GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID(id, GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIdHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionId")
	}

	result := response.Result().(*ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID)
	return result, response, err

}

//GetIssueEnrichmentDetails Get Issue Enrichment Details - 8684-39bb-4e89-a6e4
/* Enriches a given network issue context (an issue id or end user’s Mac Address) with details about the issue(s), impacted hosts and suggested actions for remediation


@param GetIssueEnrichmentDetailsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-issue-enrichment-details
*/
func (s *IssuesService) GetIssueEnrichmentDetails(GetIssueEnrichmentDetailsHeaderParams *GetIssueEnrichmentDetailsHeaderParams) (*ResponseIssuesGetIssueEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/issue-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetIssueEnrichmentDetailsHeaderParams != nil {

		if GetIssueEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetIssueEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetIssueEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetIssueEnrichmentDetailsHeaderParams.EntityValue)
		}

		if GetIssueEnrichmentDetailsHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetIssueEnrichmentDetailsHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetIssueEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIssueEnrichmentDetails(GetIssueEnrichmentDetailsHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetIssueEnrichmentDetails")
	}

	result := response.Result().(*ResponseIssuesGetIssueEnrichmentDetails)
	return result, response, err

}

//Issues Issues - ecb6-7807-47c9-bc59
/* Intent API to get a list of global issues, issues for a specific device, or issue for a specific client device's MAC address.


@param IssuesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!issues
*/
func (s *IssuesService) Issues(IssuesQueryParams *IssuesQueryParams) (*ResponseIssuesIssues, *resty.Response, error) {
	path := "/dna/intent/api/v1/issues"

	queryString, _ := query.Values(IssuesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesIssues{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Issues(IssuesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Issues")
	}

	result := response.Result().(*ResponseIssuesIssues)
	return result, response, err

}

//ReturnsAllIssueTriggerDefinitionsForGivenFilters Returns all issue trigger definitions for given filters. - 199e-880b-4dc9-95c3
/* Get all system issue defintions. The supported filters are id, name, profileId and definition enable status. An issue trigger definition can be different across the profile and device type. So, `profileId` and `deviceType` in the query param is important and default is global profile and all device type. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams Custom header parameters
@param ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-all-issue-trigger-definitions-for-given-filters
*/
func (s *IssuesService) ReturnsAllIssueTriggerDefinitionsForGivenFilters(ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams *ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams, ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams *ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams) (*ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions"

	queryString, _ := query.Values(ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams != nil {

		if ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsAllIssueTriggerDefinitionsForGivenFilters(ReturnsAllIssueTriggerDefinitionsForGivenFiltersHeaderParams, ReturnsAllIssueTriggerDefinitionsForGivenFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsAllIssueTriggerDefinitionsForGivenFilters")
	}

	result := response.Result().(*ResponseIssuesReturnsAllIssueTriggerDefinitionsForGivenFilters)
	return result, response, err

}

//GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters Get the count of system defined issue definitions based on provided filters. - a7b5-4a48-4b5b-a680
/* Get the count of system defined issue definitions based on provided filters. Supported filters are id, name, profileId and definition enable status. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams Custom header parameters
@param GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-count-of-system-defined-issue-definitions-based-on-provided-filters
*/
func (s *IssuesService) GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams *GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams, GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams *GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams) (*ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions/count"

	queryString, _ := query.Values(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams != nil {

		if GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters(GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersHeaderParams, GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters")
	}

	result := response.Result().(*ResponseIssuesGetTheCountOfSystemDefinedIssueDefinitionsBasedOnProvidedFilters)
	return result, response, err

}

//GetIssueTriggerDefinitionForGivenID Get issue trigger definition for given id. - 71a4-aa5c-400a-a129
/* Get system issue defintion for the given id. Definition includes all properties from IssueTriggerDefinition schema by default. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Issue trigger definition id.

@param GetIssueTriggerDefinitionForGivenIdHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-issue-trigger-definition-for-given-id
*/
func (s *IssuesService) GetIssueTriggerDefinitionForGivenID(id string, GetIssueTriggerDefinitionForGivenIdHeaderParams *GetIssueTriggerDefinitionForGivenIDHeaderParams) (*ResponseIssuesGetIssueTriggerDefinitionForGivenID, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetIssueTriggerDefinitionForGivenIdHeaderParams != nil {

		if GetIssueTriggerDefinitionForGivenIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetIssueTriggerDefinitionForGivenIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetIssueTriggerDefinitionForGivenID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIssueTriggerDefinitionForGivenID(id, GetIssueTriggerDefinitionForGivenIdHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetIssueTriggerDefinitionForGivenId")
	}

	result := response.Result().(*ResponseIssuesGetIssueTriggerDefinitionForGivenID)
	return result, response, err

}

//GetTheDetailsOfIssuesForGivenSetOfFilters Get the details of issues for given set of filters - 82ad-186f-4848-a3dd
/* Returns all details of each issue along with suggested actions for given set of filters specified in request body. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-issues-for-given-set-of-filters
*/
func (s *IssuesService) GetTheDetailsOfIssuesForGivenSetOfFilters(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters *RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters, GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams *GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams) (*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams != nil {

		if GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams.AcceptLanguage)
		}

		if GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters).
		SetResult(&ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfIssuesForGivenSetOfFilters(requestIssuesGetTheDetailsOfIssuesForGivenSetOfFilters, GetTheDetailsOfIssuesForGivenSetOfFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfIssuesForGivenSetOfFilters")
	}

	result := response.Result().(*ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFilters)
	return result, response, err

}

//GetTheTotalNumberOfIssuesForGivenSetOfFilters Get the total number of issues for given set of filters - b3ad-493a-409b-90b4
/* Returns the total number issues for given set of filters. If there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to 24-hours ago from end time. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml


@param GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-total-number-of-issues-for-given-set-of-filters
*/
func (s *IssuesService) GetTheTotalNumberOfIssuesForGivenSetOfFilters(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters *RequestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters, GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams *GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams) (*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams != nil {

		if GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters).
		SetResult(&ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheTotalNumberOfIssuesForGivenSetOfFilters(requestIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters, GetTheTotalNumberOfIssuesForGivenSetOfFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTheTotalNumberOfIssuesForGivenSetOfFilters")
	}

	result := response.Result().(*ResponseIssuesGetTheTotalNumberOfIssuesForGivenSetOfFilters)
	return result, response, err

}

//GetSummaryAnalyticsDataOfIssues Get summary analytics data of issues - afaa-2bdf-424b-9161
/* Gets the summary analytics data related to issues based on given filters and group by field. This data can be used to find issue counts grouped by different keys. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.1-resolved.yaml


@param GetSummaryAnalyticsDataOfIssuesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-issues
*/
func (s *IssuesService) GetSummaryAnalyticsDataOfIssues(requestIssuesGetSummaryAnalyticsDataOfIssues *RequestIssuesGetSummaryAnalyticsDataOfIssues, GetSummaryAnalyticsDataOfIssuesHeaderParams *GetSummaryAnalyticsDataOfIssuesHeaderParams) (*ResponseIssuesGetSummaryAnalyticsDataOfIssues, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfIssuesHeaderParams != nil {

		if GetSummaryAnalyticsDataOfIssuesHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetSummaryAnalyticsDataOfIssuesHeaderParams.AcceptLanguage)
		}

		if GetSummaryAnalyticsDataOfIssuesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfIssuesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetSummaryAnalyticsDataOfIssues).
		SetResult(&ResponseIssuesGetSummaryAnalyticsDataOfIssues{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfIssues(requestIssuesGetSummaryAnalyticsDataOfIssues, GetSummaryAnalyticsDataOfIssuesHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfIssues")
	}

	result := response.Result().(*ResponseIssuesGetSummaryAnalyticsDataOfIssues)
	return result, response, err

}

//GetTopNAnalyticsDataOfIssues Get Top N analytics data of issues - 21a7-c91a-4f5a-b54d
/* Gets the Top N analytics data related to issues based on given filters and group by field. This data can be used to find top sites which has most issues or top device types with most issue etc,. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.1-resolved.yaml


@param GetTopNAnalyticsDataOfIssuesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-issues
*/
func (s *IssuesService) GetTopNAnalyticsDataOfIssues(requestIssuesGetTopNAnalyticsDataOfIssues *RequestIssuesGetTopNAnalyticsDataOfIssues, GetTopNAnalyticsDataOfIssuesHeaderParams *GetTopNAnalyticsDataOfIssuesHeaderParams) (*ResponseIssuesGetTopNAnalyticsDataOfIssues, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfIssuesHeaderParams != nil {

		if GetTopNAnalyticsDataOfIssuesHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTopNAnalyticsDataOfIssuesHeaderParams.AcceptLanguage)
		}

		if GetTopNAnalyticsDataOfIssuesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfIssuesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTopNAnalyticsDataOfIssues).
		SetResult(&ResponseIssuesGetTopNAnalyticsDataOfIssues{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfIssues(requestIssuesGetTopNAnalyticsDataOfIssues, GetTopNAnalyticsDataOfIssuesHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfIssues")
	}

	result := response.Result().(*ResponseIssuesGetTopNAnalyticsDataOfIssues)
	return result, response, err

}

//GetTrendAnalyticsDataOfIssues Get trend analytics data of issues - f9ae-db6a-4618-b045
/* Gets the trend analytics data related to issues based on given filters and group by field. This data can be used to find issue counts in different intervals over a period of time. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.1-resolved.yaml


@param GetTrendAnalyticsDataOfIssuesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-issues
*/
func (s *IssuesService) GetTrendAnalyticsDataOfIssues(requestIssuesGetTrendAnalyticsDataOfIssues *RequestIssuesGetTrendAnalyticsDataOfIssues, GetTrendAnalyticsDataOfIssuesHeaderParams *GetTrendAnalyticsDataOfIssuesHeaderParams) (*ResponseIssuesGetTrendAnalyticsDataOfIssues, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceIssues/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfIssuesHeaderParams != nil {

		if GetTrendAnalyticsDataOfIssuesHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", GetTrendAnalyticsDataOfIssuesHeaderParams.AcceptLanguage)
		}

		if GetTrendAnalyticsDataOfIssuesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfIssuesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesGetTrendAnalyticsDataOfIssues).
		SetResult(&ResponseIssuesGetTrendAnalyticsDataOfIssues{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfIssues(requestIssuesGetTrendAnalyticsDataOfIssues, GetTrendAnalyticsDataOfIssuesHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfIssues")
	}

	result := response.Result().(*ResponseIssuesGetTrendAnalyticsDataOfIssues)
	return result, response, err

}

//IgnoreTheGivenListOfIssues Ignore the given list of issues - 4b92-ca6b-4918-b9fd
/* Ignores the given list of issues. The response contains the list of issues which were successfully ignored as well as the issues which are failed to ignore. After this API returns success response, it may take few seconds for the issue status to be updated if the system is heavily loaded. Please use `GET /dna/data/api/v1/assuranceIssues/{id}` API to fetch the details of a particular issue and verify `updatedTime`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml


@param IgnoreTheGivenListOfIssuesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!ignore-the-given-list-of-issues
*/
func (s *IssuesService) IgnoreTheGivenListOfIssues(requestIssuesIgnoreTheGivenListOfIssues *RequestIssuesIgnoreTheGivenListOfIssues, IgnoreTheGivenListOfIssuesHeaderParams *IgnoreTheGivenListOfIssuesHeaderParams) (*ResponseIssuesIgnoreTheGivenListOfIssues, *resty.Response, error) {
	path := "/dna/intent/api/v1/assuranceIssues/ignore"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if IgnoreTheGivenListOfIssuesHeaderParams != nil {

		if IgnoreTheGivenListOfIssuesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", IgnoreTheGivenListOfIssuesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesIgnoreTheGivenListOfIssues).
		SetResult(&ResponseIssuesIgnoreTheGivenListOfIssues{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.IgnoreTheGivenListOfIssues(requestIssuesIgnoreTheGivenListOfIssues, IgnoreTheGivenListOfIssuesHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation IgnoreTheGivenListOfIssues")
	}

	result := response.Result().(*ResponseIssuesIgnoreTheGivenListOfIssues)
	return result, response, err

}

//ResolveTheGivenListsOfIssues Resolve the given lists of issues - d48f-a9ed-4929-a6dd
/* Resolves the given list of issues. The response contains the list of issues which were successfully resolved as well as the issues which are failed to resolve. After this API returns success response, it may take few seconds for the issue status to be updated if the system is heavily loaded. Please use `GET /dna/data/api/v1/assuranceIssues/{id}` API to fetch the details of a particular issue and verify `updatedTime`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml


@param ResolveTheGivenListsOfIssuesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!resolve-the-given-lists-of-issues
*/
func (s *IssuesService) ResolveTheGivenListsOfIssues(requestIssuesResolveTheGivenListsOfIssues *RequestIssuesResolveTheGivenListsOfIssues, ResolveTheGivenListsOfIssuesHeaderParams *ResolveTheGivenListsOfIssuesHeaderParams) (*ResponseIssuesResolveTheGivenListsOfIssues, *resty.Response, error) {
	path := "/dna/intent/api/v1/assuranceIssues/resolve"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ResolveTheGivenListsOfIssuesHeaderParams != nil {

		if ResolveTheGivenListsOfIssuesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ResolveTheGivenListsOfIssuesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesResolveTheGivenListsOfIssues).
		SetResult(&ResponseIssuesResolveTheGivenListsOfIssues{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ResolveTheGivenListsOfIssues(requestIssuesResolveTheGivenListsOfIssues, ResolveTheGivenListsOfIssuesHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation ResolveTheGivenListsOfIssues")
	}

	result := response.Result().(*ResponseIssuesResolveTheGivenListsOfIssues)
	return result, response, err

}

//UpdateTheGivenIssueByUpdatingSelectedFields Update the given issue by updating selected fields - b0bc-dba1-4c19-8d7c
/* Updates selected fields in the given issue. Currently the only field that can be updated is 'notes' field. After this API returns success response, it may take few seconds for the issue details to be updated if the system is heavily loaded. Please use `GET /dna/data/api/v1/assuranceIssues/{id}` API to fetch the details of a particular issue and verify `updatedTime`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesLifecycle-1.0.0-resolved.yaml


@param id id path parameter. The issue Uuid

@param UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-the-given-issue-by-updating-selected-fields
*/
func (s *IssuesService) UpdateTheGivenIssueByUpdatingSelectedFields(id string, requestIssuesUpdateTheGivenIssueByUpdatingSelectedFields *RequestIssuesUpdateTheGivenIssueByUpdatingSelectedFields, UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams *UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams) (*ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFields, *resty.Response, error) {
	path := "/dna/intent/api/v1/assuranceIssues/{id}/update"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams != nil {

		if UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams.AcceptLanguage != "" {
			clientRequest = clientRequest.SetHeader("Accept-Language", UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams.AcceptLanguage)
		}

		if UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesUpdateTheGivenIssueByUpdatingSelectedFields).
		SetResult(&ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFields{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTheGivenIssueByUpdatingSelectedFields(id, requestIssuesUpdateTheGivenIssueByUpdatingSelectedFields, UpdateTheGivenIssueByUpdatingSelectedFieldsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation UpdateTheGivenIssueByUpdatingSelectedFields")
	}

	result := response.Result().(*ResponseIssuesUpdateTheGivenIssueByUpdatingSelectedFields)
	return result, response, err

}

//CreatesANewUserDefinedIssueDefinitions Creates a new user-defined issue definitions. - 95b5-9b50-4e48-9d82
/* Create a new custom issue definition using the provided input request data. The unique identifier for this issue definition is id. Please note that the issue names cannot be duplicated. The definition is based on the syslog. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param CreatesANewUserDefinedIssueDefinitionsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-new-user-defined-issue-definitions
*/
func (s *IssuesService) CreatesANewUserDefinedIssueDefinitions(requestIssuesCreatesANewUserDefinedIssueDefinitions *RequestIssuesCreatesANewUserDefinedIssueDefinitions, CreatesANewUserDefinedIssueDefinitionsHeaderParams *CreatesANewUserDefinedIssueDefinitionsHeaderParams) (*ResponseIssuesCreatesANewUserDefinedIssueDefinitions, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreatesANewUserDefinedIssueDefinitionsHeaderParams != nil {

		if CreatesANewUserDefinedIssueDefinitionsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CreatesANewUserDefinedIssueDefinitionsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestIssuesCreatesANewUserDefinedIssueDefinitions).
		SetResult(&ResponseIssuesCreatesANewUserDefinedIssueDefinitions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesANewUserDefinedIssueDefinitions(requestIssuesCreatesANewUserDefinedIssueDefinitions, CreatesANewUserDefinedIssueDefinitionsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreatesANewUserDefinedIssueDefinitions")
	}

	result := response.Result().(*ResponseIssuesCreatesANewUserDefinedIssueDefinitions)
	return result, response, err

}

//ExecuteSuggestedActionsCommands Execute Suggested Actions Commands - cfb2-ab10-4cea-bfbb
/* This API fetches the issue details and suggested actions for an issue, given the Issue Id, executes the commands associated with the suggested actions to remediate the issue



Documentation Link: https://developer.cisco.com/docs/dna-center/#!execute-suggested-actions-commands
*/
func (s *IssuesService) ExecuteSuggestedActionsCommands(requestIssuesExecuteSuggestedActionsCommands *RequestIssuesExecuteSuggestedActionsCommands) (*ResponseIssuesExecuteSuggestedActionsCommands, *resty.Response, error) {
	path := "/dna/intent/api/v1/execute-suggested-actions-commands"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIssuesExecuteSuggestedActionsCommands).
		SetResult(&ResponseIssuesExecuteSuggestedActionsCommands{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExecuteSuggestedActionsCommands(requestIssuesExecuteSuggestedActionsCommands)
		}

		return nil, response, fmt.Errorf("error with operation ExecuteSuggestedActionsCommands")
	}

	result := response.Result().(*ResponseIssuesExecuteSuggestedActionsCommands)
	return result, response, err

}

//UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID Updates an existing custom issue definition based on the provided Id. - 8b90-3b69-4c18-90ad
/* Updates an existing custom issue definition based on the provided Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param id id path parameter. The custom issue definition Identifier

*/
func (s *IssuesService) UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID(id string, requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedId *RequestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID) (*ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID, *resty.Response, error) {
	path := "/dna/intent/api/v1/customIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	response, err = clientRequest.
		SetBody(requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedId).
		SetResult(&ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID(id, requestIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedId)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedId")
	}

	result := response.Result().(*ResponseIssuesUpdatesAnExistingCustomIssueDefinitionBasedOnTheProvidedID)
	return result, response, err

}

//IssueTriggerDefinitionUpdate Issue trigger definition update. - 099a-397b-46c8-8aa7
/* Update issue trigger threshold, priority for the given id.
Also enable or disable issue trigger for the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Issue trigger definition id.

*/
func (s *IssuesService) IssueTriggerDefinitionUpdate(id string, requestIssuesIssueTriggerDefinitionUpdate *RequestIssuesIssueTriggerDefinitionUpdate) (*ResponseIssuesIssueTriggerDefinitionUpdate, *resty.Response, error) {
	path := "/dna/intent/api/v1/systemIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIssuesIssueTriggerDefinitionUpdate).
		SetResult(&ResponseIssuesIssueTriggerDefinitionUpdate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.IssueTriggerDefinitionUpdate(id, requestIssuesIssueTriggerDefinitionUpdate)
		}
		return nil, response, fmt.Errorf("error with operation IssueTriggerDefinitionUpdate")
	}

	result := response.Result().(*ResponseIssuesIssueTriggerDefinitionUpdate)
	return result, response, err

}

//DeletesAnExistingCustomIssueDefinition Deletes an existing custom issue definition. - e38b-fa80-4c28-955f
/* Deletes an existing custom issue definition based on the Id. Only the Global profile issue has the access to delete the issue definition, so no profile id is required. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml


@param id id path parameter. The custom issue definition unique identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-an-existing-custom-issue-definition
*/
func (s *IssuesService) DeletesAnExistingCustomIssueDefinition(id string) (*resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/customIssueDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAnExistingCustomIssueDefinition(
				id)
		}
		return response, fmt.Errorf("error with operation DeletesAnExistingCustomIssueDefinition")
	}

	return response, err

}
