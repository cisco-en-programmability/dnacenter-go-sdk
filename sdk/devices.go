package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DevicesService service

type RetrievesTheListOfAAAServicesForGivenParametersQueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                 float64 `url:"limit,omitempty"`                 //Maximum number of records to return
	Offset                float64 `url:"offset,omitempty"`                //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                string  `url:"sortBy,omitempty"`                //Field name on which sorting needs to be done
	Order                 string  `url:"order,omitempty"`                 //The sort order of the field ascending or descending.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the AAA Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	SiteHierarchy         string  `url:"siteHierarchy,omitempty"`         //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID                string  `url:"siteId,omitempty"`                //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheListOfAAAServicesForGivenParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfAAAServicesForGivenParametersQueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the AAA Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	DeviceSiteHierarchy   string  `url:"deviceSiteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryAssuranceEventsQueryParams struct {
	DeviceFamily      string  `url:"deviceFamily,omitempty"`      //Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing `Routers` along with `Wireless Client` or `Unified AP` is not supported. Examples: `deviceFamily=Switches and Hubs` (single deviceFamily requested) `deviceFamily=Switches and Hubs&deviceFamily=Routers` (multiple deviceFamily requested)
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time minus 24 hours.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to current time.
	MessageType       string  `url:"messageType,omitempty"`       //Message type for the event. Examples: `messageType=Syslog` (single messageType requested) `messageType=Trap&messageType=Syslog` (multiple messageType requested)
	Severity          float64 `url:"severity,omitempty"`          //Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and `Wired Client` events. | Value | Severity    | | ----- | ----------- | | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        | Examples: `severity=0` (single severity requested) `severity=0&severity=1` (multiple severity requested)
	SiteID            string  `url:"siteId,omitempty"`            //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single siteId requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple siteId requested)
	SiteHierarchyID   string  `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyId requested)
	NetworkDeviceName string  `url:"networkDeviceName,omitempty"` //Network device name. This parameter is applicable for network device related families. This field supports wildcard (`*`) character-based search. Ex: `*Branch*` or `Branch*` or `*Branch` Examples: `networkDeviceName=Branch-3-Gateway` (single networkDeviceName requested) `networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch` (multiple networkDeviceName requested)
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceId with & separator)
	ApMac             string  `url:"apMac,omitempty"`             //MAC address of the access point. This parameter is applicable for `Unified AP` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*50:0F*` or `50:0F*` or `*50:0F` Examples: `apMac=50:0F:80:0F:F7:E0` (single apMac requested) `apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0` (multiple apMac requested)
	ClientMac         string  `url:"clientMac,omitempty"`         //MAC address of the client. This parameter is applicable for `Wired Client` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*66:2B*` or `66:2B*` or `*66:2B` Examples: `clientMac=66:2B:B8:D2:01:56` (single clientMac requested) `clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89` (multiple clientMac requested)
	Attribute         string  `url:"attribute,omitempty"`         //The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes (`id`, `name`, `timestamp`, `details`, `messageType`, `siteHierarchyId`, `siteHierarchy`, `deviceFamily`, `networkDeviceId`, `networkDeviceName`, `managementIpAddress`) would be part of the response.  Examples:  `attribute=name` (single attribute requested) `attribute=name&attribute=networkDeviceName` (multiple attribute requested)
	View              string  `url:"view,omitempty"`              //The list of events views. Please refer to `EventViews` for the supported list  Examples:  `view=network` (single view requested) `view=network&view=ap` (multiple view requested)
	Offset            float64 `url:"offset,omitempty"`            //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	SortBy            string  `url:"sortBy,omitempty"`            //A field within the response to sort by.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
}
type QueryAssuranceEventsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountTheNumberOfEventsQueryParams struct {
	DeviceFamily      string `url:"deviceFamily,omitempty"`      //Device family. Please note that multiple families across network device type and client type is not allowed. For example, choosing `Routers` along with `Wireless Client` or `Unified AP` is not supported. Examples: `deviceFamily=Switches and Hubs` (single deviceFamily requested) `deviceFamily=Switches and Hubs&deviceFamily=Routers` (multiple deviceFamily requested)
	StartTime         string `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time minus 24 hours.
	EndTime           string `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to current time.
	MessageType       string `url:"messageType,omitempty"`       //Message type for the event. Examples: `messageType=Syslog` (single messageType requested) `messageType=Trap&messageType=Syslog` (multiple messageType requested)
	Severity          string `url:"severity,omitempty"`          //Severity of the event between 0 and 6. This is applicable only for events related to network devices (other than AP) and `Wired Client` events. | Value | Severity    | | ----- | ----------- | | 0     | Emergency   | | 1     | Alert       | | 2     | Critical    | | 3     | Error       | | 4     | Warning     | | 5     | Notice      | | 6     | Info        | Examples: `severity=0` (single severity requested) `severity=0&severity=1` (multiple severity requested)
	SiteID            string `url:"siteId,omitempty"`            //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single siteId requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple siteId requested)
	SiteHierarchyID   string `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyId requested)
	NetworkDeviceName string `url:"networkDeviceName,omitempty"` //Network device name. This parameter is applicable for network device related families. This field supports wildcard (`*`) character-based search. Ex: `*Branch*` or `Branch*` or `*Branch` Examples: `networkDeviceName=Branch-3-Gateway` (single networkDeviceName requested) `networkDeviceName=Branch-3-Gateway&networkDeviceName=Branch-3-Switch` (multiple networkDeviceName requested)
	NetworkDeviceID   string `url:"networkDeviceId,omitempty"`   //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceId requested)
	ApMac             string `url:"apMac,omitempty"`             //MAC address of the access point. This parameter is applicable for `Unified AP` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*50:0F*` or `50:0F*` or `*50:0F` Examples: `apMac=50:0F:80:0F:F7:E0` (single apMac requested) `apMac=50:0F:80:0F:F7:E0&apMac=18:80:90:AB:7E:A0` (multiple apMac requested)
	ClientMac         string `url:"clientMac,omitempty"`         //MAC address of the client. This parameter is applicable for `Wired Client` and `Wireless Client` events. This field supports wildcard (`*`) character-based search. Ex: `*66:2B*` or `66:2B*` or `*66:2B` Examples: `clientMac=66:2B:B8:D2:01:56` (single clientMac requested) `clientMac=66:2B:B8:D2:01:56&clientMac=DC:A6:32:F5:5A:89` (multiple clientMac requested)
}
type CountTheNumberOfEventsHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryAssuranceEventsWithFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountTheNumberOfEventsWithFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetDetailsOfASingleAssuranceEventQueryParams struct {
	Attribute string `url:"attribute,omitempty"` //The list of attributes that needs to be included in the response. If this parameter is not provided, then basic attributes (`id`, `name`, `timestamp`, `details`, `messageType`, `siteHierarchyId`, `siteHierarchy`, `deviceFamily`, `networkDeviceId`, `networkDeviceName`, `managementIpAddress`) would be part of the response.  Examples:  `attribute=name` (single attribute requested) `attribute=name&attribute=networkDeviceName` (multiple attribute requested)
	View      string `url:"view,omitempty"`      //The list of events views. Please refer to `EventViews` for the supported list  Examples:  `view=network` (single view requested) `view=network&view=ap` (multiple view requested)
}
type GetDetailsOfASingleAssuranceEventHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDHCPServicesForGivenParametersQueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                 float64 `url:"limit,omitempty"`                 //Maximum number of records to return
	Offset                float64 `url:"offset,omitempty"`                //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                string  `url:"sortBy,omitempty"`                //Field name on which sorting needs to be done
	Order                 string  `url:"order,omitempty"`                 //The sort order of the field ascending or descending.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DHCP Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	DeviceSiteHierarchy   string  `url:"deviceSiteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DHCP Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceName            string  `url:"deviceName,omitempty"`            //Name of the device. This parameter supports wildcard (`*`) character -based search. Example: `wnbu-sjc*` or `*wnbu-sjc*` or `*wnbu-sjc` Examples: deviceName=wnbu-sjc24.cisco.com (single device name is requested) deviceName=wnbu-sjc24.cisco.com&deviceName=wnbu-sjc22.cisco.com (multiple device names are requested)
	DeviceSiteHierarchy   string  `url:"deviceSiteHierarchy,omitempty"`   //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?deviceSiteHierarchy=Global/AreaName/BuildingName/FloorName&deviceSiteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
}
type RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDNSServicesForGivenParametersQueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                 float64 `url:"limit,omitempty"`                 //Maximum number of records to return
	Offset                float64 `url:"offset,omitempty"`                //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                string  `url:"sortBy,omitempty"`                //Field name on which sorting needs to be done
	Order                 string  `url:"order,omitempty"`                 //The sort order of the field ascending or descending.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DNS Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
	SSID                  string  `url:"ssid,omitempty"`                  //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the field contains the (`*`) character, please use the /query API for search. Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
}
type RetrievesTheListOfDNSServicesForGivenParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDNSServicesForGivenParametersQueryParams struct {
	StartTime             float64 `url:"startTime,omitempty"`             //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime               float64 `url:"endTime,omitempty"`               //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ServerIP              string  `url:"serverIp,omitempty"`              //IP Address of the DNS Server. This parameter supports wildcard (`*`) character -based search. Example: `10.76.81.*` or `*56.78*` or `*50.28` Examples: serverIp=10.42.3.31 (single IP Address is requested) serverIp=10.42.3.31&serverIp=name2&fabricVnName=name3 (multiple IP Addresses are requested)
	DeviceID              string  `url:"deviceId,omitempty"`              //The device UUID.   Examples:  `deviceId=6bef213c-19ca-4170-8375-b694e251101c` (single deviceId is requested)  `deviceId=6bef213c-19ca-4170-8375-b694e251101c&deviceId=32219612-819e-4b5e-a96b-cf22aca13dd9 (multiple networkDeviceIds with & separator)
	DeviceSiteHierarchyID string  `url:"deviceSiteHierarchyId,omitempty"` //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?deviceSiteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&deviceSiteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceSiteID          string  `url:"deviceSiteId,omitempty"`          //The UUID of the site. (Ex. `flooruuid`) Examples: `?deviceSiteIds=id1` (single id requested) `?deviceSiteIds=id1&deviceSiteIds=id2&siteId=id3` (multiple ids requested)
	SSID                  string  `url:"ssid,omitempty"`                  //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the field contains the (`*`) character, please use the /query API for search. Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
}
type RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesQueryParams struct {
	StartTime               float64 `url:"startTime,omitempty"`               //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                 float64 `url:"endTime,omitempty"`                 //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                   float64 `url:"limit,omitempty"`                   //Maximum number of records to return
	Offset                  float64 `url:"offset,omitempty"`                  //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                  string  `url:"sortBy,omitempty"`                  //A field within the response to sort by.
	Order                   string  `url:"order,omitempty"`                   //The sort order of the field ascending or descending.
	SiteHierarchy           string  `url:"siteHierarchy,omitempty"`           //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID         string  `url:"siteHierarchyId,omitempty"`         //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID                  string  `url:"siteId,omitempty"`                  //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	View                    string  `url:"view,omitempty"`                    //Views which are supported by this API. Each view represents a specific         data set.           ### Response data provided by each view:             1. **configuration**          [id,name,adminStatus,description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,         portChannelId,portMode,         portType,speed,timestamp,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]             2. **statistics**          [id,name,rxDiscards,rxError,rxRate,rxUtilization,txDiscards,txError,txRate,txUtilization,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]            3. **stackPort**          [id,name,peerStackMember,peerStackPort,stackPortType,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]                   4. **poE**              [id, name,rxDiscards,rxError,rxRate,rxUtilization,txDiscards,txError,txRate,txUtilization,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId]              When this query parameter is not added by default all configuration attributes will be         available in the response.     **[configuration,statistics,stackPort]**
	Attribute               string  `url:"attribute,omitempty"`               //The following list of attributes can be provided in the attribute field          [id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId,poeAdminStatus,poeOperStatus,chassisId,moduleId,pdClassSignal,pdClassSpare,pdDeviceType,pdDeviceModel,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn,pdConnectedDeviceList,poeOperPriority,fastPoEEnabled,perpetualPoEEnabled,policingPoEEnabled,upoePlusEnabled,fourPairEnabled,poeDataTimestamp,pdLocation,pdDeviceName,pdConnectedSwitch,connectedSwitchUuid,ieeeCompliant,connectedSwitchType]          If length of attribute list is too long, please use 'views' param instead.          Examples:          attributes=name (single attribute requested)          attributes=name&description&duplexOper (multiple attributes with comma separator)
	NetworkDeviceID         string  `url:"networkDeviceId,omitempty"`         //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress  string  `url:"networkDeviceIpAddress,omitempty"`  //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	NetworkDeviceMacAddress string  `url:"networkDeviceMacAddress,omitempty"` //The list of Network Device MAC Address. (Ex. `64:f6:9d:07:9a:00`) This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `networkDeviceMacAddress=64:f6:9d:07:9a:00` `networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77` (multiple networkDevice MAC addresses with & separator)
	InterfaceID             string  `url:"interfaceId,omitempty"`             //The list of Interface Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `interfaceId=6bef213c-19ca-4170-8375-b694e251101c` (single interface uuid ) `interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple Interface uuid with & separator)
	InterfaceName           string  `url:"interfaceName,omitempty"`           //The list of Interface name (Ex. `GigabitEthernet1/0/1`) This field supports wildcard (`*`) character-based search. Ex: `*1/0/1*` or `1/0/1*` or `*1/0/1` Examples: `interfaceNames=GigabitEthernet1/0/1` (single interface name) `interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1` (multiple interface names with & separator)
}
type GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams struct {
	StartTime               float64 `url:"startTime,omitempty"`               //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                 float64 `url:"endTime,omitempty"`                 //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	SiteHierarchy           string  `url:"siteHierarchy,omitempty"`           //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID         string  `url:"siteHierarchyId,omitempty"`         //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID                  string  `url:"siteId,omitempty"`                  //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	NetworkDeviceID         string  `url:"networkDeviceId,omitempty"`         //The list of Network Device Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c` (single networkDeviceId requested) `networkDeviceId=6bef213c-19ca-4170-8375-b694e251101c&networkDeviceId=32219612-819e-4b5e-a96b-cf22aca13dd9&networkDeviceId=2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple networkDeviceIds with & separator)
	NetworkDeviceIPAddress  string  `url:"networkDeviceIpAddress,omitempty"`  //The list of Network Device management IP Address. (Ex. `121.1.1.10`) This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `networkDeviceIpAddress=121.1.1.10` `networkDeviceIpAddress=121.1.1.10&networkDeviceIpAddress=172.20.1.10&networkDeviceIpAddress=10.10.20.10` (multiple networkDevice IP Address with & separator)
	NetworkDeviceMacAddress string  `url:"networkDeviceMacAddress,omitempty"` //The list of Network Device MAC Address. (Ex. `64:f6:9d:07:9a:00`) This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `networkDeviceMacAddress=64:f6:9d:07:9a:00` `networkDeviceMacAddress=64:f6:9d:07:9a:00&networkDeviceMacAddress=70:56:9d:07:ac:77` (multiple networkDevice MAC addresses with & separator)
	InterfaceID             string  `url:"interfaceId,omitempty"`             //The list of Interface Uuids. (Ex. `6bef213c-19ca-4170-8375-b694e251101c`) Examples: `interfaceId=6bef213c-19ca-4170-8375-b694e251101c` (single interface uuid ) `interfaceId=6bef213c-19ca-4170-8375-b694e251101c&32219612-819e-4b5e-a96b-cf22aca13dd9&2541e9a7-b80d-4955-8aa2-79b233318ba0` (multiple Interface uuid with & separator)
	InterfaceName           string  `url:"interfaceName,omitempty"`           //The list of Interface name (Ex. `GigabitEthernet1/0/1`) This field supports wildcard (`*`) character-based search. Ex: `*1/0/1*` or `1/0/1*` or `*1/0/1` Examples: `interfaceNames=GigabitEthernet1/0/1` (single interface name) `interfaceNames=GigabitEthernet1/0/1&GigabitEthernet2/0/1&GigabitEthernet3/0/1` (multiple interface names with & separator)
}
type GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //Interface data model views
	Attribute string  `url:"attribute,omitempty"` //The following list of attributes can be provided in the attribute field          [id,adminStatus, description,duplexConfig,duplexOper,interfaceIfIndex,interfaceType,ipv4Address,ipv6AddressList,isL3Interface,isWan,macAddress,mediaType,name,operStatus,peerStackMember,peerStackPort, portChannelId,portMode, portType,rxDiscards,rxError,rxRate,rxUtilization,speed,stackPortType,timestamp,txDiscards,txError,txRate,txUtilization,vlanId,networkDeviceId,networkDeviceIpAddress,networkDeviceMacAddress,siteName,siteHierarchy,siteHierarchyId,poeAdminStatus,poeOperStatus,chassisId,moduleId,pdClassSignal,pdClassSpare,pdDeviceType,pdDeviceModel,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn,pdConnectedDeviceList,poeOperPriority,fastPoEEnabled,perpetualPoEEnabled,policingPoEEnabled,upoePlusEnabled,fourPairEnabled,poeDataTimestamp,pdLocation,pdDeviceName,pdConnectedSwitch,connectedSwitchUuid,ieeeCompliant,connectedSwitchType]          If length of attribute list is too long, please use 'views' param instead.          Examples:          attributes=name (single attribute requested)          attributes=name&description&duplexOper (multiple attributes with comma separator)
}
type GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams struct {
	StartTime           float64 `url:"startTime,omitempty"`           //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime             float64 `url:"endTime,omitempty"`             //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit               float64 `url:"limit,omitempty"`               //Maximum number of records to return
	Offset              float64 `url:"offset,omitempty"`              //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy              string  `url:"sortBy,omitempty"`              //A field within the response to sort by.
	Order               string  `url:"order,omitempty"`               //The sort order of the field ascending or descending.
	SiteHierarchy       string  `url:"siteHierarchy,omitempty"`       //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID     string  `url:"siteHierarchyId,omitempty"`     //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID              string  `url:"siteId,omitempty"`              //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	ID                  string  `url:"id,omitempty"`                  //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	ManagementIPAddress string  `url:"managementIpAddress,omitempty"` //The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10") This field supports wildcard (`*`) character-based search.  Ex: `*1.1*` or `1.1*` or `*1.1` Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
	MacAddress          string  `url:"macAddress,omitempty"`          //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	Family              string  `url:"family,omitempty"`              //The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
	Type                string  `url:"type,omitempty"`                //The list of network device type This field supports wildcard (`*`) character-based search. Ex: `*9407R*` or `*9407R` or `9407R*` Examples: type=SwitchesCisco Catalyst 9407R Switch (single network device types ) type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
	Role                string  `url:"role,omitempty"`                //The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
	SerialNumber        string  `url:"serialNumber,omitempty"`        //The list of network device serial numbers. This field supports wildcard (`*`) character-based search.  Ex: `*MS1SV*` or `MS1SV*` or `*MS1SV` Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
	MaintenanceMode     bool    `url:"maintenanceMode,omitempty"`     //The device maintenanceMode status true or false
	SoftwareVersion     string  `url:"softwareVersion,omitempty"`     //The list of network device software version This field supports wildcard (`*`) character-based search. Ex: `*17.8*` or `*17.8` or `17.8*` Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
	HealthScore         string  `url:"healthScore,omitempty"`         //The list of entity health score categories Examples: healthScore=good, healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
	View                string  `url:"view,omitempty"`                //The List of Network Device model views. Please refer to ```NetworkDeviceView``` section in the Open API specification document mentioned in the description.
	Attribute           string  `url:"attribute,omitempty"`           //The List of Network Device model attributes. Please refer to ```NetworkDeviceAttribute``` section in the Open API specification document mentioned in the description.
	FabricSiteID        string  `url:"fabricSiteId,omitempty"`        //The fabric site Id or list to fabric site Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?fabricSiteId=fabricSiteUuid)  ?fabricSiteId=fabricSiteUuid1&fabricSiteId=fabricSiteUuid2 (multiple fabricSiteIds requested)
	L2Vn                string  `url:"l2Vn,omitempty"`                //The L2 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l2Vn=virtualNetworkId  ?l2Vn=virtualNetworkId1&l2Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	L3Vn                string  `url:"l3Vn,omitempty"`                //The L3 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l3Vn=virtualNetworkId  ?l3Vn=virtualNetworkId1&l3Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	TransitNetworkID    string  `url:"transitNetworkId,omitempty"`    //The Transit Network Id or list to Transit Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?transitNetworkId=transitNetworkId  ?transitNetworkId=transitNetworkuuid1&transitNetworkId=transitNetworkuuid1 (multiple transitNetworkIds requested
	FabricRole          string  `url:"fabricRole,omitempty"`          //The list of fabric device role. Examples: fabricRole=BORDER, fabricRole=BORDER&fabricRole=EDGE (multiple fabric device roles with & separator)  Available values : BORDER, EDGE, MAP-SERVER, LEAF, SPINE, TRANSIT-CP, EXTENDED-NODE, WLC, UNIFIED-AP
}
type GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams struct {
	StartTime           float64 `url:"startTime,omitempty"`           //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime             float64 `url:"endTime,omitempty"`             //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ID                  string  `url:"id,omitempty"`                  //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	SiteHierarchy       string  `url:"siteHierarchy,omitempty"`       //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (*) character search support. E.g. */San*, */San, /San* Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID     string  `url:"siteHierarchyId,omitempty"`     //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (*) character search support. E.g. `*uuid*, *uuid, uuid* Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	SiteID              string  `url:"siteId,omitempty"`              //The UUID of the site. (Ex. `flooruuid`) This field supports wildcard asterisk (*) character search support. E.g.*flooruuid*, *flooruuid, flooruuid* Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	ManagementIPAddress string  `url:"managementIpAddress,omitempty"` //The list of entity management IP Address. It can be either Ipv4 or Ipv6 address or combination of both(Ex. "121.1.1.10") This field supports wildcard (`*`) character-based search.  Ex: `*1.1*` or `1.1*` or `*1.1` Examples: managementIpAddresses=121.1.1.10 managementIpAddresses=121.1.1.10&managementIpAddresses=172.20.1.10&managementIpAddresses=200:10&=managementIpAddresses172.20.3.4 (multiple entity IP Address with & separator)
	MacAddress          string  `url:"macAddress,omitempty"`          //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	Family              string  `url:"family,omitempty"`              //The list of network device family names Examples:family=Switches and Hubs (single network device family name )family=Switches and Hubs&family=Router&family=Wireless Controller (multiple Network device family names with & separator). This field is not case sensitive.
	Type                string  `url:"type,omitempty"`                //The list of network device type This field supports wildcard (`*`) character-based search. Ex: `*9407R*` or `*9407R` or `9407R*`Examples:type=SwitchesCisco Catalyst 9407R Switch (single network device types )type=Cisco Catalyst 38xx stack-able ethernet switch&type=Cisco 3945 Integrated Services Router G2 (multiple Network device types with & separator)
	Role                string  `url:"role,omitempty"`                //The list of network device role. Examples:role=CORE, role=CORE&role=ACCESS&role=ROUTER (multiple Network device roles with & separator). This field is not case sensitive.
	SerialNumber        string  `url:"serialNumber,omitempty"`        //The list of network device serial numbers. This field supports wildcard (`*`) character-based search.  Ex: `*MS1SV*` or `MS1SV*` or `*MS1SV` Examples: serialNumber=9FUFMS1SVAX serialNumber=9FUFMS1SVAX&FCW2333Q0BY&FJC240617JX(multiple Network device serial number with & separator)
	MaintenanceMode     bool    `url:"maintenanceMode,omitempty"`     //The device maintenanceMode status true or false
	SoftwareVersion     string  `url:"softwareVersion,omitempty"`     //The list of network device software version This field supports wildcard (`*`) character-based search. Ex: `*17.8*` or `*17.8` or `17.8*` Examples: softwareVersion=2.3.4.0 (single network device software version ) softwareVersion=17.9.3.23&softwareVersion=17.7.1.2&softwareVersion=*.17.7 (multiple Network device software versions with & separator)
	HealthScore         string  `url:"healthScore,omitempty"`         //The list of entity health score categories Examples:healthScore=good,healthScore=good&healthScore=fair (multiple entity healthscore values with & separator). This field is not case sensitive.
	FabricSiteID        string  `url:"fabricSiteId,omitempty"`        //The fabric site Id or list to fabric site Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?fabricSiteId=fabricSiteUuid)  ?fabricSiteId=fabricSiteUuid1&fabricSiteId=fabricSiteUuid2 (multiple fabricSiteIds requested)
	L2Vn                string  `url:"l2Vn,omitempty"`                //The L2 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l2Vn=virtualNetworkId  ?l2Vn=virtualNetworkId1&l2Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	L3Vn                string  `url:"l3Vn,omitempty"`                //The L3 Virtual Network Id or list to Virtual Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?l3Vn=virtualNetworkId  ?l3Vn=virtualNetworkId1&l3Vn=virtualNetworkId2 (multiple virtualNetworkId's requested)
	TransitNetworkID    string  `url:"transitNetworkId,omitempty"`    //The Transit Network Id or list to Transit Network Ids to filter the data  This field supports wildcard asterisk (*) character search support. E.g. *uuid*, *uuid, uuid*  Examples:  `?transitNetworkId=transitNetworkId  ?transitNetworkId=transitNetworkuuid1&transitNetworkId=transitNetworkuuid1 (multiple transitNetworkIds requested)
	FabricRole          string  `url:"fabricRole,omitempty"`          //The list of fabric device role. Examples: fabricRole=BORDER, fabricRole=BORDER&fabricRole=EDGE (multiple fabric device roles with & separator)  Available values : BORDER, EDGE, MAP-SERVER, LEAF, SPINE, TRANSIT-CP, EXTENDED-NODE, WLC, UNIFIED-AP
}
type GetTheDeviceDataForTheGivenDeviceIDUUIDQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //The List of Network Device model views. Please refer to ```NetworkDeviceView``` section in the Open API specification document mentioned in the description.
	Attribute string  `url:"attribute,omitempty"` //The List of Network Device model attributes. Please refer to ```NetworkDeviceAttribute``` section in the Open API specification document mentioned in the description.
}
type GetPlannedAccessPointsForBuildingQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	Offset float64 `url:"offset,omitempty"` //The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
	Radios bool    `url:"radios,omitempty"` //Whether to include the planned radio details of the planned access points
}
type GetDeviceDetailQueryParams struct {
	Timestamp  float64 `url:"timestamp,omitempty"`  //UTC timestamp of device data in milliseconds
	IDentifier string  `url:"identifier,omitempty"` //One of "macAddress", "nwDeviceName", "uuid" (case insensitive)
	SearchBy   string  `url:"searchBy,omitempty"`   //MAC Address, device name, or UUID of the network device
}
type GetDeviceEnrichmentDetailsHeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. For the enrichment details to be made available as part of the API response, this header must be set to true. This header must be explicitly passed when called from client applications outside Catalyst Center
}
type DevicesQueryParams struct {
	DeviceRole string  `url:"deviceRole,omitempty"` //CORE, ACCESS, DISTRIBUTION, ROUTER, WLC, or AP (case insensitive)
	SiteID     string  `url:"siteId,omitempty"`     //DNAC site UUID
	Health     string  `url:"health,omitempty"`     //DNAC health catagory: POOR, FAIR, or GOOD (case insensitive)
	StartTime  float64 `url:"startTime,omitempty"`  //UTC epoch time in milliseconds
	EndTime    float64 `url:"endTime,omitempty"`    //UTC epoch time in milliseconds
	Limit      float64 `url:"limit,omitempty"`      //Max number of device entries in the response (default to 50. Max at 500)
	Offset     float64 `url:"offset,omitempty"`     //The offset of the first device in the returned data (Mutiple of 'limit' + 1)
}
type GetPlannedAccessPointsForFloorQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	Offset float64 `url:"offset,omitempty"` //The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
	Radios bool    `url:"radios,omitempty"` //Whether to include the planned radio details of the planned access points
}
type GetAllHealthScoreDefinitionsForGivenFiltersQueryParams struct {
	DeviceType              string  `url:"deviceType,omitempty"`              //These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
	ID                      string  `url:"id,omitempty"`                      //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	IncludeForOverallHealth bool    `url:"includeForOverallHealth,omitempty"` //The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
	Attribute               string  `url:"attribute,omitempty"`               //These are the attributes supported in health score definitions response. By default, all properties are sent in response.
	Offset                  float64 `url:"offset,omitempty"`                  //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Limit                   float64 `url:"limit,omitempty"`                   //Maximum number of records to return
}
type GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type UpdateHealthScoreDefinitionsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams struct {
	DeviceType              string `url:"deviceType,omitempty"`              //These are the device families supported for health score definitions. If no input is made on device family, all device families are considered.
	ID                      string `url:"id,omitempty"`                      //The definition identifier. Examples: id=015d9cba-4f53-4087-8317-7e49e5ffef46 (single entity id request) id=015d9cba-4f53-4087-8317-7e49e5ffef46&id=015d9cba-4f53-4087-8317-7e49e5ffef47 (multiple ids in the query param)
	IncludeForOverallHealth bool   `url:"includeForOverallHealth,omitempty"` //The inclusion status of the issue definition, either true or false. true indicates that particular health metric is included in overall health computation, otherwise false. By default it's set to true.
}
type GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetHealthScoreDefinitionForTheGivenIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetAllInterfacesQueryParams struct {
	Offset         int    `url:"offset,omitempty"`         //Offset
	Limit          int    `url:"limit,omitempty"`          //The number of records to show for this page. Min: 1, Max: 500
	LastInputTime  string `url:"lastInputTime,omitempty"`  //Last Input Time
	LastOutputTime string `url:"lastOutputTime,omitempty"` //Last Output Time
}
type GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams struct {
	Name string `url:"name,omitempty"` //Interface name
}
type UpdateInterfaceDetailsQueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type ClearMacAddressTableQueryParams struct {
	DeploymentMode string `url:"deploymentMode,omitempty"` //Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
}
type GetDeviceListQueryParams struct {
	Hostname                  []string `url:"hostname,omitempty"`                   //hostname
	ManagementIPAddress       []string `url:"managementIpAddress,omitempty"`        //managementIpAddress
	MacAddress                []string `url:"macAddress,omitempty"`                 //macAddress
	LocationName              []string `url:"locationName,omitempty"`               //locationName
	SerialNumber              []string `url:"serialNumber,omitempty"`               //serialNumber
	Location                  []string `url:"location,omitempty"`                   //location
	Family                    []string `url:"family,omitempty"`                     //family
	Type                      []string `url:"type,omitempty"`                       //type
	Series                    []string `url:"series,omitempty"`                     //series
	CollectionStatus          []string `url:"collectionStatus,omitempty"`           //collectionStatus
	CollectionInterval        []string `url:"collectionInterval,omitempty"`         //collectionInterval
	NotSyncedForMinutes       []string `url:"notSyncedForMinutes,omitempty"`        //notSyncedForMinutes
	ErrorCode                 []string `url:"errorCode,omitempty"`                  //errorCode
	ErrorDescription          []string `url:"errorDescription,omitempty"`           //errorDescription
	SoftwareVersion           []string `url:"softwareVersion,omitempty"`            //softwareVersion
	SoftwareType              []string `url:"softwareType,omitempty"`               //softwareType
	PlatformID                []string `url:"platformId,omitempty"`                 //platformId
	Role                      []string `url:"role,omitempty"`                       //role
	ReachabilityStatus        []string `url:"reachabilityStatus,omitempty"`         //reachabilityStatus
	UpTime                    []string `url:"upTime,omitempty"`                     //upTime
	AssociatedWlcIP           []string `url:"associatedWlcIp,omitempty"`            //associatedWlcIp
	Licensename               []string `url:"license.name,omitempty"`               //licenseName
	Licensetype               []string `url:"license.type,omitempty"`               //licenseType
	Licensestatus             []string `url:"license.status,omitempty"`             //licenseStatus
	Modulename                []string `url:"module+name,omitempty"`                //moduleName
	Moduleequpimenttype       []string `url:"module+equpimenttype,omitempty"`       //moduleEqupimentType
	Moduleservicestate        []string `url:"module+servicestate,omitempty"`        //moduleServiceState
	Modulevendorequipmenttype []string `url:"module+vendorequipmenttype,omitempty"` //moduleVendorEquipmentType
	Modulepartnumber          []string `url:"module+partnumber,omitempty"`          //modulePartNumber
	Moduleoperationstatecode  []string `url:"module+operationstatecode,omitempty"`  //moduleOperationStateCode
	ID                        string   `url:"id,omitempty"`                         //Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
	DeviceSupportLevel        string   `url:"deviceSupportLevel,omitempty"`         //deviceSupportLevel
	Offset                    int      `url:"offset,omitempty"`                     //offset >= 1 [X gives results from Xth device onwards]
	Limit                     int      `url:"limit,omitempty"`                      //The number of records to show for this page. Min: 1, Max: 500
}
type GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams struct {
	VrfName                   string `url:"vrfName,omitempty"`                   //vrfName
	ManagementIPAddress       string `url:"managementIpAddress,omitempty"`       //managementIpAddress
	Hostname                  string `url:"hostname,omitempty"`                  //hostname
	MacAddress                string `url:"macAddress,omitempty"`                //macAddress
	Family                    string `url:"family,omitempty"`                    //family
	CollectionStatus          string `url:"collectionStatus,omitempty"`          //collectionStatus
	CollectionInterval        string `url:"collectionInterval,omitempty"`        //collectionInterval
	SoftwareVersion           string `url:"softwareVersion,omitempty"`           //softwareVersion
	SoftwareType              string `url:"softwareType,omitempty"`              //softwareType
	ReachabilityStatus        string `url:"reachabilityStatus,omitempty"`        //reachabilityStatus
	ReachabilityFailureReason string `url:"reachabilityFailureReason,omitempty"` //reachabilityFailureReason
	ErrorCode                 string `url:"errorCode,omitempty"`                 //errorCode
	PlatformID                string `url:"platformId,omitempty"`                //platformId
	Series                    string `url:"series,omitempty"`                    //series
	Type                      string `url:"type,omitempty"`                      //type
	SerialNumber              string `url:"serialNumber,omitempty"`              //serialNumber
	UpTime                    string `url:"upTime,omitempty"`                    //upTime
	Role                      string `url:"role,omitempty"`                      //role
	RoleSource                string `url:"roleSource,omitempty"`                //roleSource
	AssociatedWlcIP           string `url:"associatedWlcIp,omitempty"`           //associatedWlcIp
	Offset                    int    `url:"offset,omitempty"`                    //offset
	Limit                     int    `url:"limit,omitempty"`                     //The number of records to show for this page. Min: 1, Max: 500
}
type GetDeviceCountKnowYourNetworkQueryParams struct {
	Hostname            []string `url:"hostname,omitempty"`            //hostname
	ManagementIPAddress []string `url:"managementIpAddress,omitempty"` //managementIpAddress
	MacAddress          []string `url:"macAddress,omitempty"`          //macAddress
	LocationName        []string `url:"locationName,omitempty"`        //locationName
}
type GetFunctionalCapabilityForDevicesQueryParams struct {
	DeviceID     string   `url:"deviceId,omitempty"`     //Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
	FunctionName []string `url:"functionName,omitempty"` //functionName
}
type InventoryInsightDeviceLinkMismatchAPIQueryParams struct {
	Offset   int    `url:"offset,omitempty"`   //Row Number.  Default value is 1
	Limit    int    `url:"limit,omitempty"`    //The number of records to show for this page. Min: 1, Max: 500
	Category string `url:"category,omitempty"` //Links mismatch category.  Value can be speed-duplex or vlan.
	SortBy   string `url:"sortBy,omitempty"`   //Sort By
	Order    string `url:"order,omitempty"`    //Order.  Value can be asc or desc.  Default value is asc
}
type GetModulesQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	Limit                    int      `url:"limit,omitempty"`                    //The number of records to show for this page. Min: 1, Max: 500
	Offset                   int      `url:"offset,omitempty"`                   //offset
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type GetModuleCountQueryParams struct {
	DeviceID                 string   `url:"deviceId,omitempty"`                 //deviceId
	NameList                 []string `url:"nameList,omitempty"`                 //nameList
	VendorEquipmentTypeList  []string `url:"vendorEquipmentTypeList,omitempty"`  //vendorEquipmentTypeList
	PartNumberList           []string `url:"partNumberList,omitempty"`           //partNumberList
	OperationalStateCodeList []string `url:"operationalStateCodeList,omitempty"` //operationalStateCodeList
}
type SyncDevicesQueryParams struct {
	ForceSync bool `url:"forceSync,omitempty"` //forceSync
}
type GetDevicesRegisteredForWsaNotificationQueryParams struct {
	SerialNumber string `url:"serialNumber,omitempty"` //Serial number of the device
	Macaddress   string `url:"macaddress,omitempty"`   //Mac addres of the device
}
type GetAllUserDefinedFieldsQueryParams struct {
	ID   string `url:"id,omitempty"`   //Comma-seperated id(s) used for search/filtering
	Name string `url:"name,omitempty"` //Comma-seperated name(s) used for search/filtering
}
type RemoveUserDefinedFieldFromDeviceQueryParams struct {
	Name string `url:"name,omitempty"` //Name of UDF to be removed
}
type GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams struct {
	Type string `url:"type,omitempty"` //Type value can be PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other, SFP. If no type is mentioned, All equipments are fetched for the device.
}
type ReturnsPoeInterfaceDetailsForTheDeviceQueryParams struct {
	InterfaceNameList string `url:"interfaceNameList,omitempty"` //comma seperated interface names
}
type DeleteDeviceByIDQueryParams struct {
	CleanConfig bool `url:"cleanConfig,omitempty"` //Selecting the clean up configuration option will attempt to remove device settings that were configured during the addition of the device to the inventory and site assignment. Please note that this operation is different from deprovisioning. It does not remove configurations that were pushed during device provisioning.
}
type GetDeviceInterfaceVLANsQueryParams struct {
	InterfaceType string `url:"interfaceType,omitempty"` //Vlan associated with sub-interface. If no interfaceType mentioned it will return all types of Vlan interfaces. If interfaceType is selected but not specified then it will take default value.
}
type RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams struct {
	NetworkDeviceIDs string `url:"networkDeviceIds,omitempty"` //List of network device ids.
	Status           string `url:"status,omitempty"`           //The status of the maintenance schedule. Possible values are: UPCOMING, IN_PROGRESS, COMPLETED, FAILED. Refer features for more details.
	Limit            string `url:"limit,omitempty"`            //The number of records to show for this page. Min: 1, Max: 500
	Offset           string `url:"offset,omitempty"`           //The first record to show for this page; the first record is numbered 1.
	SortBy           string `url:"sortBy,omitempty"`           //A property within the response to sort by.
	Order            string `url:"order,omitempty"`            //Whether ascending or descending order should be used to sort the response.
}
type RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams struct {
	NetworkDeviceIDs string `url:"networkDeviceIds,omitempty"` //List of network device ids.
	Status           string `url:"status,omitempty"`           //The status of the maintenance schedule. Possible values are: UPCOMING, IN_PROGRESS, COMPLETED, FAILED. Refer features for more details.
}
type RetrieveNetworkDevicesQueryParams struct {
	ID                 string `url:"id,omitempty"`                 //Network device Id
	ManagementAddress  string `url:"managementAddress,omitempty"`  //Management address of the network device
	SerialNumber       string `url:"serialNumber,omitempty"`       //Serial number of the network device
	Family             string `url:"family,omitempty"`             //Product family of the network device. For example, Switches, Routers, etc.
	StackDevice        string `url:"stackDevice,omitempty"`        //Flag indicating if the device is a stack device
	Role               string `url:"role,omitempty"`               //Role assigned to the network device. Available values : BORDER_ROUTER, CORE, DISTRIBUTION, ACCESS, UNKNOWN
	Status             string `url:"status,omitempty"`             //Inventory related status of the network device. Available values : MANAGED, SYNC_NOT_STARTED, SYNC_INIT_FAILED, SYNC_PRECHECK_FAILED, SYNC_IN_PROGRESS, SYNC_INTERNAL_ERROR, SYNC_DISABLED, DELETING_DEVICE, UNDER_MAINTENANCE, QUARANTINED, UNASSOCIATED, UNREACHABLE, UNKNOWN. Refer features for more details.
	ReachabilityStatus string `url:"reachabilityStatus,omitempty"` //Reachability status of the network device. Available values : REACHABLE, ONLY_PING_REACHABLE, UNREACHABLE, UNKNOWN. Refer features for more details.
	ManagementState    string `url:"managementState,omitempty"`    //The status of the network device's manageability. Available statuses are MANAGED, UNDER_MAINTENANCE, NEVER_MANAGED. Refer features for more details.
	Views              string `url:"views,omitempty"`              //The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views. Refer features for more details. Available values : BASIC, RESYNC, USER_DEFINED_FIELDS.
	Limit              string `url:"limit,omitempty"`              //The number of records to show for this page. Min: 1, Max: 500
	Offset             string `url:"offset,omitempty"`             //The first record to show for this page; the first record is numbered 1.
	SortBy             string `url:"sortBy,omitempty"`             //A property within the response to sort by. Available values : id, managementAddress, dnsResolvedManagementIpAddress, hostname, macAddress, type, family, series, platformids, softwareType, softwareVersion, vendor, bootTime, role, roleSource, apEthernetMacAddress, apManagerInterfaceIpAddress, apWlcIpAddress, deviceSupportLevel, reachabilityFailureReason, resyncStartTime, resyncEndTime, resyncReasons, pendingResyncRequestCount, pendingResyncRequestReasons, resyncIntervalSource, resyncIntervalMinutes
	Order              string `url:"order,omitempty"`              //Whether ascending or descending order should be used to sort the response.
}
type CountTheNumberOfNetworkDevicesQueryParams struct {
	ID                 string `url:"id,omitempty"`                 //Network device Id
	ManagementAddress  string `url:"managementAddress,omitempty"`  //Management address of the network device
	SerialNumber       string `url:"serialNumber,omitempty"`       //Serial number of the network device
	Family             string `url:"family,omitempty"`             //Product family of the network device. For example, Switches, Routers, etc.
	StackDevice        string `url:"stackDevice,omitempty"`        //Flag indicating if the device is a stack device
	Role               string `url:"role,omitempty"`               //Role assigned to the network device. Available values : BORDER_ROUTER, CORE, DISTRIBUTION, ACCESS, UNKNOWN
	Status             string `url:"status,omitempty"`             //Inventory related status of the network device. Available values : MANAGED, SYNC_NOT_STARTED, SYNC_INIT_FAILED, SYNC_PRECHECK_FAILED, SYNC_IN_PROGRESS, SYNC_INTERNAL_ERROR, SYNC_DISABLED, DELETING_DEVICE, UNDER_MAINTENANCE, QUARANTINED, UNASSOCIATED, UNREACHABLE, UNKNOWN. Refer features for more details.
	ReachabilityStatus string `url:"reachabilityStatus,omitempty"` //Reachability status of the network device. Available values : REACHABLE, ONLY_PING_REACHABLE, UNREACHABLE, UNKNOWN. Refer features for more details.
	ManagementState    string `url:"managementState,omitempty"`    //The status of the network device's manageability. Available values : MANAGED, UNDER_MAINTENANCE, NEVER_MANAGED. Refer features for more details.
}
type GetDetailsOfASingleNetworkDeviceQueryParams struct {
	Views string `url:"views,omitempty"` //The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views. Available values : BASIC, RESYNC, USER_DEFINED_FIELDS
}
type GetAllowedMacAddressQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The offset of the first item in the collection to return.
	Limit  float64 `url:"limit,omitempty"`  //The maximum number of entries to return. If the value exceeds the total count, then the maximum entries will be returned.
}

type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParameters struct {
	Response *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersPage       `json:"page,omitempty"`     //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersResponse struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	EapTransactions           *int   `json:"eapTransactions,omitempty"`           // Eap Transactions
	EapFailedTransactions     *int   `json:"eapFailedTransactions,omitempty"`     // Eap Failed Transactions
	EapSuccessfulTransactions *int   `json:"eapSuccessfulTransactions,omitempty"` // Eap Successful Transactions
	MabTransactions           *int   `json:"mabTransactions,omitempty"`           // Mab Transactions
	MabFailedTransactions     *int   `json:"mabFailedTransactions,omitempty"`     // Mab Failed Transactions
	MabSuccessfulTransactions *int   `json:"mabSuccessfulTransactions,omitempty"` // Mab Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	EapLatency                *int   `json:"eapLatency,omitempty"`                // Eap Latency
	MabLatency                *int   `json:"mabLatency,omitempty"`                // Mab Latency
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersPage struct {
	Limit  *int                                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                                        `json:"offset,omitempty"` // Offset
	Count  *int                                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenParametersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParameters struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersResponse `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParametersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters struct {
	Response *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersResponse struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	EapTransactions           *int   `json:"eapTransactions,omitempty"`           // Eap Transactions
	EapFailedTransactions     *int   `json:"eapFailedTransactions,omitempty"`     // Eap Failed Transactions
	EapSuccessfulTransactions *int   `json:"eapSuccessfulTransactions,omitempty"` // Eap Successful Transactions
	MabTransactions           *int   `json:"mabTransactions,omitempty"`           // Mab Transactions
	MabFailedTransactions     *int   `json:"mabFailedTransactions,omitempty"`     // Mab Failed Transactions
	MabSuccessfulTransactions *int   `json:"mabSuccessfulTransactions,omitempty"` // Mab Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	EapLatency                *int   `json:"eapLatency,omitempty"`                // Eap Latency
	MabLatency                *int   `json:"mabLatency,omitempty"`                // Mab Latency
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                 `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                 `json:"offset,omitempty"` // Offset
	Count  *int                                                                                 `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Version
	Response *ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage     `json:"page,omitempty"`     //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse struct {
	Groups              *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroups struct {
	ID                  string                                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                        `json:"offset,omitempty"` // Offset
	Count  string                                                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                 `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse struct {
	ID                  string                                                                                                    `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                     `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                     `json:"offset,omitempty"` // Offset
	Count  string                                                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponse struct {
	Timestamp           *int                                                                                                       `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroups struct {
	ID                  string                                                                                                           `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService struct {
	Response *ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceResponse `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceResponse struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	EapTransactions           *int   `json:"eapTransactions,omitempty"`           // Eap Transactions
	EapFailedTransactions     *int   `json:"eapFailedTransactions,omitempty"`     // Eap Failed Transactions
	EapSuccessfulTransactions *int   `json:"eapSuccessfulTransactions,omitempty"` // Eap Successful Transactions
	MabTransactions           *int   `json:"mabTransactions,omitempty"`           // Mab Transactions
	MabFailedTransactions     *int   `json:"mabFailedTransactions,omitempty"`     // Mab Failed Transactions
	MabSuccessfulTransactions *int   `json:"mabSuccessfulTransactions,omitempty"` // Mab Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	EapLatency                *int   `json:"eapLatency,omitempty"`                // Eap Latency
	MabLatency                *int   `json:"mabLatency,omitempty"`                // Mab Latency
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService struct {
	Version  string                                                                                      `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServicePage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Groups *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseGroups `json:"groups,omitempty"` //

	Attributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *int `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *int `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServicePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesQueryAssuranceEvents struct {
	Response *[]ResponseDevicesQueryAssuranceEventsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
	Page     *ResponseDevicesQueryAssuranceEventsPage       `json:"page,omitempty"`     //
}
type ResponseDevicesQueryAssuranceEventsResponse struct {
	OldRadioChannelWidth         string                                                           `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                           `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                           `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                             `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                         `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                            `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                           `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                           `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                           `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                           `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                           `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                           `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                           `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                           `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                           `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                             `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                           `json:"details,omitempty"`                      // Details
	ID                           string                                                           `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                           `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                           `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                           `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                           `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                           `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                           `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                           `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                           `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                           `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                           `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                           `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                           `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                           `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                           `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                           `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                           `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                           `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesQueryAssuranceEventsResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                           `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                           `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                           `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                           `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                             `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                           `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                           `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                           `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesQueryAssuranceEventsResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                             `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                             `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                             `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                           `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                             `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                           `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesQueryAssuranceEventsResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                           `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                           `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                           `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                             `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                           `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                           `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                           `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                           `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                           `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                             `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                           `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                           `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                           `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                           `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesQueryAssuranceEventsResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                           `json:"username,omitempty"`                     // Username
}
type ResponseDevicesQueryAssuranceEventsResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesQueryAssuranceEventsResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesQueryAssuranceEventsResponseChildEvents struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonCode        string `json:"subReasonCode,omitempty"`        // Sub Reason Code
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesQueryAssuranceEventsResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesQueryAssuranceEventsPage struct {
	Limit  *int                                             `json:"limit,omitempty"`  // Limit
	Offset *int                                             `json:"offset,omitempty"` // Offset
	Count  *int                                             `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesQueryAssuranceEventsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesQueryAssuranceEventsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesCountTheNumberOfEvents struct {
	Response *ResponseDevicesCountTheNumberOfEventsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesCountTheNumberOfEventsResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesQueryAssuranceEventsWithFilters struct {
	Response *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
	Page     *ResponseDevicesQueryAssuranceEventsWithFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponse struct {
	OldRadioChannelWidth         string                                                                      `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                                      `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                                      `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                                        `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                                    `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                                       `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                                      `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                                      `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                                      `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                                      `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                                      `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                                      `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                                      `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                                      `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                                      `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                                        `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                                      `json:"details,omitempty"`                      // Details
	ID                           string                                                                      `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                                      `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                                      `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                                      `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                                      `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                                      `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                                      `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                                      `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                                      `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                                      `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                                      `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                                      `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                                      `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                                      `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                                      `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                                      `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                                      `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                                      `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                                      `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                                      `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                                      `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                                      `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                                        `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                                      `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                                      `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                                      `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                                        `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                                        `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                                        `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                                      `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                                        `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                                      `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                                      `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                                      `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                                      `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                                        `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                                      `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                                      `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                                      `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                                      `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                                      `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                                        `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                                      `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                                      `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                                      `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                                      `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesQueryAssuranceEventsWithFiltersResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                                      `json:"username,omitempty"`                     // Username
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseChildEvents struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonCode        string `json:"subReasonCode,omitempty"`        // Sub Reason Code
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesQueryAssuranceEventsWithFiltersResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesQueryAssuranceEventsWithFiltersPage struct {
	Limit  *int                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                        `json:"offset,omitempty"` // Offset
	Count  *int                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesQueryAssuranceEventsWithFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesQueryAssuranceEventsWithFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesCountTheNumberOfEventsWithFilters struct {
	Response *ResponseDevicesCountTheNumberOfEventsWithFiltersResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesCountTheNumberOfEventsWithFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetDetailsOfASingleAssuranceEvent struct {
	Response *ResponseDevicesGetDetailsOfASingleAssuranceEventResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponse struct {
	OldRadioChannelWidth         string                                                                        `json:"oldRadioChannelWidth,omitempty"`         // Old Radio Channel Width
	ClientMac                    string                                                                        `json:"clientMac,omitempty"`                    // Client Mac
	SwitchNumber                 string                                                                        `json:"switchNumber,omitempty"`                 // Switch Number
	AssocRssi                    *int                                                                          `json:"assocRssi,omitempty"`                    // Assoc Rssi
	AffectedClients              []string                                                                      `json:"affectedClients,omitempty"`              // Affected Clients
	IsPrivateMac                 *bool                                                                         `json:"isPrivateMac,omitempty"`                 // Is Private Mac
	Frequency                    string                                                                        `json:"frequency,omitempty"`                    // Frequency
	ApRole                       string                                                                        `json:"apRole,omitempty"`                       // Ap Role
	ReplacingDeviceSerialNumber  string                                                                        `json:"replacingDeviceSerialNumber,omitempty"`  // Replacing Device Serial Number
	MessageType                  string                                                                        `json:"messageType,omitempty"`                  // Message Type
	FailureCategory              string                                                                        `json:"failureCategory,omitempty"`              // Failure Category
	ApSwitchName                 string                                                                        `json:"apSwitchName,omitempty"`                 // Ap Switch Name
	ApSwitchID                   string                                                                        `json:"apSwitchId,omitempty"`                   // Ap Switch Id
	RadioChannelUtilization      string                                                                        `json:"radioChannelUtilization,omitempty"`      // Radio Channel Utilization
	Mnemonic                     string                                                                        `json:"mnemonic,omitempty"`                     // Mnemonic
	RadioChannelSlot             *int                                                                          `json:"radioChannelSlot,omitempty"`             // Radio Channel Slot
	Details                      string                                                                        `json:"details,omitempty"`                      // Details
	ID                           string                                                                        `json:"id,omitempty"`                           // Id
	LastApDisconnectReason       string                                                                        `json:"lastApDisconnectReason,omitempty"`       // Last Ap Disconnect Reason
	NetworkDeviceName            string                                                                        `json:"networkDeviceName,omitempty"`            // Network Device Name
	IDentifier                   string                                                                        `json:"identifier,omitempty"`                   // Identifier
	ReasonDescription            string                                                                        `json:"reasonDescription,omitempty"`            // Reason Description
	VLANID                       string                                                                        `json:"vlanId,omitempty"`                       // Vlan Id
	UdnID                        string                                                                        `json:"udnId,omitempty"`                        // Udn Id
	AuditSessionID               string                                                                        `json:"auditSessionId,omitempty"`               // Audit Session Id
	ApMac                        string                                                                        `json:"apMac,omitempty"`                        // Ap Mac
	DeviceFamily                 string                                                                        `json:"deviceFamily,omitempty"`                 // Device Family
	RadioNoise                   string                                                                        `json:"radioNoise,omitempty"`                   // Radio Noise
	WlcName                      string                                                                        `json:"wlcName,omitempty"`                      // Wlc Name
	ApRadioOperationState        string                                                                        `json:"apRadioOperationState,omitempty"`        // Ap Radio Operation State
	Name                         string                                                                        `json:"name,omitempty"`                         // Name
	FailureIPAddress             string                                                                        `json:"failureIpAddress,omitempty"`             // Failure Ip Address
	NewRadioChannelList          string                                                                        `json:"newRadioChannelList,omitempty"`          // New Radio Channel List
	Duid                         string                                                                        `json:"duid,omitempty"`                         // Duid
	RoamType                     string                                                                        `json:"roamType,omitempty"`                     // Roam Type
	CandidateAPs                 *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseCandidateAPs       `json:"candidateAPs,omitempty"`                 //
	ReplacedDeviceSerialNumber   string                                                                        `json:"replacedDeviceSerialNumber,omitempty"`   // Replaced Device Serial Number
	OldRadioChannelList          string                                                                        `json:"oldRadioChannelList,omitempty"`          // Old Radio Channel List
	SSID                         string                                                                        `json:"ssid,omitempty"`                         // Ssid
	SubReasonDescription         string                                                                        `json:"subReasonDescription,omitempty"`         // Sub Reason Description
	WirelessClientEventEndTime   *int                                                                          `json:"wirelessClientEventEndTime,omitempty"`   // Wireless Client Event End Time
	IPv4                         string                                                                        `json:"ipv4,omitempty"`                         // Ipv4
	WlcID                        string                                                                        `json:"wlcId,omitempty"`                        // Wlc Id
	IPv6                         string                                                                        `json:"ipv6,omitempty"`                         // Ipv6
	MissingResponseAPs           *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseMissingResponseAPs `json:"missingResponseAPs,omitempty"`           //
	Timestamp                    *int                                                                          `json:"timestamp,omitempty"`                    // Timestamp
	Severity                     *int                                                                          `json:"severity,omitempty"`                     // Severity
	CurrentRadioPowerLevel       *int                                                                          `json:"currentRadioPowerLevel,omitempty"`       // Current Radio Power Level
	NewRadioChannelWidth         string                                                                        `json:"newRadioChannelWidth,omitempty"`         // New Radio Channel Width
	AssocSnr                     *int                                                                          `json:"assocSnr,omitempty"`                     // Assoc Snr
	AuthServerIP                 string                                                                        `json:"authServerIp,omitempty"`                 // Auth Server Ip
	ChildEvents                  *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseChildEvents        `json:"childEvents,omitempty"`                  //
	ConnectedInterfaceName       string                                                                        `json:"connectedInterfaceName,omitempty"`       // Connected Interface Name
	DhcpServerIP                 string                                                                        `json:"dhcpServerIp,omitempty"`                 // Dhcp Server Ip
	ManagementIPAddress          string                                                                        `json:"managementIpAddress,omitempty"`          // Management Ip Address
	PreviousRadioPowerLevel      *int                                                                          `json:"previousRadioPowerLevel,omitempty"`      // Previous Radio Power Level
	ResultStatus                 string                                                                        `json:"resultStatus,omitempty"`                 // Result Status
	RadioInterference            string                                                                        `json:"radioInterference,omitempty"`            // Radio Interference
	NetworkDeviceID              string                                                                        `json:"networkDeviceId,omitempty"`              // Network Device Id
	SiteHierarchy                string                                                                        `json:"siteHierarchy,omitempty"`                // Site Hierarchy
	EventStatus                  string                                                                        `json:"eventStatus,omitempty"`                  // Event Status
	WirelessClientEventStartTime *int                                                                          `json:"wirelessClientEventStartTime,omitempty"` // Wireless Client Event Start Time
	SiteHierarchyID              string                                                                        `json:"siteHierarchyId,omitempty"`              // Site Hierarchy Id
	UdnName                      string                                                                        `json:"udnName,omitempty"`                      // Udn Name
	Facility                     string                                                                        `json:"facility,omitempty"`                     // Facility
	LastApResetType              string                                                                        `json:"lastApResetType,omitempty"`              // Last Ap Reset Type
	InvalidIeAPs                 *[]ResponseDevicesGetDetailsOfASingleAssuranceEventResponseInvalidIeAPs       `json:"invalidIeAPs,omitempty"`                 //
	Username                     string                                                                        `json:"username,omitempty"`                     // Username
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseCandidateAPs struct {
	APID   string `json:"apId,omitempty"`   // Ap Id
	ApName string `json:"apName,omitempty"` // Ap Name
	ApMac  string `json:"apMac,omitempty"`  // Ap Mac
	Bssid  string `json:"bssid,omitempty"`  // Bssid
	Rssi   *int   `json:"rssi,omitempty"`   // Rssi
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseMissingResponseAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseChildEvents struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonCode        string `json:"subReasonCode,omitempty"`        // Sub Reason Code
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesGetDetailsOfASingleAssuranceEventResponseInvalidIeAPs struct {
	APID      string `json:"apId,omitempty"`      // Ap Id
	ApName    string `json:"apName,omitempty"`    // Ap Name
	ApMac     string `json:"apMac,omitempty"`     // Ap Mac
	Bssid     string `json:"bssid,omitempty"`     // Bssid
	Type      string `json:"type,omitempty"`      // Type
	FrameType string `json:"frameType,omitempty"` // Frame Type
	Ies       string `json:"ies,omitempty"`       // Ies
}
type ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent struct {
	Response *[]ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEventResponse struct {
	ID                   string `json:"id,omitempty"`                   // Id
	Name                 string `json:"name,omitempty"`                 // Name
	Timestamp            *int   `json:"timestamp,omitempty"`            // Timestamp
	WirelessEventType    *int   `json:"wirelessEventType,omitempty"`    // Wireless Event Type
	Details              string `json:"details,omitempty"`              // Details
	ReasonCode           string `json:"reasonCode,omitempty"`           // Reason Code
	SubreasonCode        string `json:"subreasonCode,omitempty"`        // Subreason Code
	ResultStatus         string `json:"resultStatus,omitempty"`         // Result Status
	ReasonDescription    string `json:"reasonDescription,omitempty"`    // Reason Description
	SubReasonDescription string `json:"subReasonDescription,omitempty"` // Sub Reason Description
	FailureCategory      string `json:"failureCategory,omitempty"`      // Failure Category
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParameters struct {
	Response *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersPage       `json:"page,omitempty"`     //
	Version  string                                                                     `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersResponse struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	DiscoverOfferLatency      *int   `json:"discoverOfferLatency,omitempty"`      // Discover Offer Latency
	RequestAcknowledgeLatency *int   `json:"requestAcknowledgeLatency,omitempty"` // Request Acknowledge Latency
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersPage struct {
	Limit  *int                                                                         `json:"limit,omitempty"`  // Limit
	Offset *int                                                                         `json:"offset,omitempty"` // Offset
	Count  *int                                                                         `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParametersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParameters struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersResponse `json:"response,omitempty"` //
	Version  string                                                                          `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParametersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters struct {
	Response *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersResponse struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	DiscoverOfferLatency      *int   `json:"discoverOfferLatency,omitempty"`      // Discover Offer Latency
	RequestAcknowledgeLatency *int   `json:"requestAcknowledgeLatency,omitempty"` // Request Acknowledge Latency
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                  `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                  `json:"offset,omitempty"` // Offset
	Count  *int                                                                                  `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                   `json:"version,omitempty"`  // Version
	Response *ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage     `json:"page,omitempty"`     //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponse struct {
	Groups              *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroups struct {
	ID                  string                                                                                                              `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                         `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                         `json:"offset,omitempty"` // Offset
	Count  string                                                                                       `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponse struct {
	ID                  string                                                                                                     `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                      `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                      `json:"offset,omitempty"` // Offset
	Count  string                                                                                    `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                   `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponse struct {
	Timestamp           *int                                                                                                        `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroups struct {
	ID                  string                                                                                                            `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService struct {
	Response *ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceResponse `json:"response,omitempty"` //
	Version  string                                                                                     `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceResponse struct {
	ID                        string `json:"id,omitempty"`                        // Id
	ServerIP                  string `json:"serverIp,omitempty"`                  // Server Ip
	DeviceID                  string `json:"deviceId,omitempty"`                  // Device Id
	DeviceName                string `json:"deviceName,omitempty"`                // Device Name
	DeviceFamily              string `json:"deviceFamily,omitempty"`              // Device Family
	DeviceSiteHierarchy       string `json:"deviceSiteHierarchy,omitempty"`       // Device Site Hierarchy
	DeviceSiteID              string `json:"deviceSiteId,omitempty"`              // Device Site Id
	DeviceSiteHierarchyID     string `json:"deviceSiteHierarchyId,omitempty"`     // Device Site Hierarchy Id
	Transactions              *int   `json:"transactions,omitempty"`              // Transactions
	FailedTransactions        *int   `json:"failedTransactions,omitempty"`        // Failed Transactions
	SuccessfulTransactions    *int   `json:"successfulTransactions,omitempty"`    // Successful Transactions
	Latency                   *int   `json:"latency,omitempty"`                   // Latency
	DiscoverOfferLatency      *int   `json:"discoverOfferLatency,omitempty"`      // Discover Offer Latency
	RequestAcknowledgeLatency *int   `json:"requestAcknowledgeLatency,omitempty"` // Request Acknowledge Latency
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService struct {
	Version  string                                                                                       `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServicePage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponse struct {
	Timestamp           *int                                                                                                            `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseGroups struct {
	ID                  string                                                                                                                `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServicePage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParameters struct {
	Response *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersPage       `json:"page,omitempty"`     //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersResponse struct {
	ID                     string                                                                            `json:"id,omitempty"`                     // Id
	ServerIP               string                                                                            `json:"serverIp,omitempty"`               // Server Ip
	DeviceID               string                                                                            `json:"deviceId,omitempty"`               // Device Id
	DeviceName             string                                                                            `json:"deviceName,omitempty"`             // Device Name
	DeviceFamily           string                                                                            `json:"deviceFamily,omitempty"`           // Device Family
	DeviceSiteHierarchy    string                                                                            `json:"deviceSiteHierarchy,omitempty"`    // Device Site Hierarchy
	DeviceSiteID           string                                                                            `json:"deviceSiteId,omitempty"`           // Device Site Id
	DeviceSiteHierarchyID  string                                                                            `json:"deviceSiteHierarchyId,omitempty"`  // Device Site Hierarchy Id
	Transactions           *int                                                                              `json:"transactions,omitempty"`           // Transactions
	FailedTransactions     *int                                                                              `json:"failedTransactions,omitempty"`     // Failed Transactions
	Failures               *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersResponseFailures `json:"failures,omitempty"`               //
	SuccessfulTransactions *int                                                                              `json:"successfulTransactions,omitempty"` // Successful Transactions
	Latency                *int                                                                              `json:"latency,omitempty"`                // Latency
	SSID                   string                                                                            `json:"ssid,omitempty"`                   // Ssid
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersResponseFailures struct {
	FailureResponseCode *int   `json:"failureResponseCode,omitempty"` // Failure Response Code
	FailureDescription  string `json:"failureDescription,omitempty"`  // Failure Description
	FailedTransactions  *int   `json:"failedTransactions,omitempty"`  // Failed Transactions
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersPage struct {
	Limit  *int                                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                                        `json:"offset,omitempty"` // Offset
	Count  *int                                                                        `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenParametersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParameters struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersResponse `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParametersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters struct {
	Response *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersResponse struct {
	ID                     string                                                                                     `json:"id,omitempty"`                     // Id
	ServerIP               string                                                                                     `json:"serverIp,omitempty"`               // Server Ip
	DeviceID               string                                                                                     `json:"deviceId,omitempty"`               // Device Id
	DeviceName             string                                                                                     `json:"deviceName,omitempty"`             // Device Name
	DeviceFamily           string                                                                                     `json:"deviceFamily,omitempty"`           // Device Family
	DeviceSiteHierarchy    string                                                                                     `json:"deviceSiteHierarchy,omitempty"`    // Device Site Hierarchy
	DeviceSiteID           string                                                                                     `json:"deviceSiteId,omitempty"`           // Device Site Id
	DeviceSiteHierarchyID  string                                                                                     `json:"deviceSiteHierarchyId,omitempty"`  // Device Site Hierarchy Id
	Transactions           *int                                                                                       `json:"transactions,omitempty"`           // Transactions
	FailedTransactions     *int                                                                                       `json:"failedTransactions,omitempty"`     // Failed Transactions
	Failures               *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersResponseFailures `json:"failures,omitempty"`               //
	SuccessfulTransactions *int                                                                                       `json:"successfulTransactions,omitempty"` // Successful Transactions
	Latency                *int                                                                                       `json:"latency,omitempty"`                // Latency
	SSID                   string                                                                                     `json:"ssid,omitempty"`                   // Ssid
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersResponseFailures struct {
	FailureResponseCode *int   `json:"failureResponseCode,omitempty"` // Failure Response Code
	FailureDescription  string `json:"failureDescription,omitempty"`  // Failure Description
	FailedTransactions  *int   `json:"failedTransactions,omitempty"`  // Failed Transactions
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                 `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                 `json:"offset,omitempty"` // Offset
	Count  *int                                                                                 `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters struct {
	Response *ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Version
	Response *ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage     `json:"page,omitempty"`     //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponse struct {
	Groups              *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroups struct {
	ID                  string                                                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                        `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                        `json:"offset,omitempty"` // Offset
	Count  string                                                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                 `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponse struct {
	ID                  string                                                                                                    `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit  *int                                                                                     `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                     `json:"offset,omitempty"` // Offset
	Count  string                                                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Order    string `json:"order,omitempty"`    // Order
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponse struct {
	Timestamp           *int                                                                                                       `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroups struct {
	ID                  string                                                                                                           `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService struct {
	Response *ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceResponse `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceResponse struct {
	ID                     string                                                                                              `json:"id,omitempty"`                     // Id
	ServerIP               string                                                                                              `json:"serverIp,omitempty"`               // Server Ip
	DeviceID               string                                                                                              `json:"deviceId,omitempty"`               // Device Id
	DeviceName             string                                                                                              `json:"deviceName,omitempty"`             // Device Name
	DeviceFamily           string                                                                                              `json:"deviceFamily,omitempty"`           // Device Family
	DeviceSiteHierarchy    string                                                                                              `json:"deviceSiteHierarchy,omitempty"`    // Device Site Hierarchy
	DeviceSiteID           string                                                                                              `json:"deviceSiteId,omitempty"`           // Device Site Id
	DeviceSiteHierarchyID  string                                                                                              `json:"deviceSiteHierarchyId,omitempty"`  // Device Site Hierarchy Id
	Transactions           *int                                                                                                `json:"transactions,omitempty"`           // Transactions
	FailedTransactions     *int                                                                                                `json:"failedTransactions,omitempty"`     // Failed Transactions
	Failures               *[]ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceResponseFailures `json:"failures,omitempty"`               //
	SuccessfulTransactions *int                                                                                                `json:"successfulTransactions,omitempty"` // Successful Transactions
	Latency                *int                                                                                                `json:"latency,omitempty"`                // Latency
	SSID                   string                                                                                              `json:"ssid,omitempty"`                   // Ssid
}
type ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceResponseFailures struct {
	FailureResponseCode *int   `json:"failureResponseCode,omitempty"` // Failure Response Code
	FailureDescription  string `json:"failureDescription,omitempty"`  // Failure Description
	FailedTransactions  *int   `json:"failedTransactions,omitempty"`  // Failed Transactions
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService struct {
	Version  string                                                                                      `json:"version,omitempty"`  // Version
	Response *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage       `json:"page,omitempty"`     //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponse struct {
	Timestamp           *int                                                                                                           `json:"timestamp,omitempty"`           // Timestamp
	Groups              *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroups              `json:"groups,omitempty"`              //
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroups struct {
	ID                  string                                                                                                               `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices struct {
	Response *[]ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesPage       `json:"page,omitempty"`     //
	Version  string                                                                                     `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesResponse struct {
	ID                      string   `json:"id,omitempty"`                      // Id
	AdminStatus             string   `json:"adminStatus,omitempty"`             // Admin Status
	Description             string   `json:"description,omitempty"`             // Description
	DuplexConfig            string   `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string   `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int     `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string   `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string   `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool    `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool    `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string   `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string   `json:"mediaType,omitempty"`               // Media Type
	Name                    string   `json:"name,omitempty"`                    // Name
	OperStatus              string   `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int     `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string   `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string   `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string   `json:"portMode,omitempty"`                // Port Mode
	PortType                string   `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64 `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int     `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64 `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64 `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string   `json:"speed,omitempty"`                   // Speed
	StackPortType           string   `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int     `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64 `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int     `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64 `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64 `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string   `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string   `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string   `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string   `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteHierarchy           string   `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string   `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	PoeAdminStatus          string   `json:"poeAdminStatus,omitempty"`          // Poe Admin Status
	PoeOperStatus           string   `json:"poeOperStatus,omitempty"`           // Poe Oper Status
	ChassisID               *int     `json:"chassisId,omitempty"`               // Chassis Id
	ModuleID                *int     `json:"moduleId,omitempty"`                // Module Id
	PdClassSignal           string   `json:"pdClassSignal,omitempty"`           // Pd Class Signal
	PdClassSpare            string   `json:"pdClassSpare,omitempty"`            // Pd Class Spare
	PdDeviceType            string   `json:"pdDeviceType,omitempty"`            // Pd Device Type
	PdDeviceModel           string   `json:"pdDeviceModel,omitempty"`           // Pd Device Model
	PdPowerAdminMaxInWatt   string   `json:"pdPowerAdminMaxInWatt,omitempty"`   // Pd Power Admin Max In Watt
	PdPowerBudgetInWatt     string   `json:"pdPowerBudgetInWatt,omitempty"`     // Pd Power Budget In Watt
	PdPowerConsumedInWatt   string   `json:"pdPowerConsumedInWatt,omitempty"`   // Pd Power Consumed In Watt
	PdPowerRemainingInWatt  string   `json:"pdPowerRemainingInWatt,omitempty"`  // Pd Power Remaining In Watt
	PdMaxPowerDrawn         string   `json:"pdMaxPowerDrawn,omitempty"`         // Pd Max Power Drawn
	PdConnectedDeviceList   []string `json:"pdConnectedDeviceList,omitempty"`   // Pd Connected Device List
	PoeOperPriority         string   `json:"poeOperPriority,omitempty"`         // Poe Oper Priority
	FastPoEEnabled          *bool    `json:"fastPoEEnabled,omitempty"`          // Fast Po E Enabled
	PerpetualPoEEnabled     *bool    `json:"perpetualPoEEnabled,omitempty"`     // Perpetual Po E Enabled
	PolicingPoEEnabled      *bool    `json:"policingPoEEnabled,omitempty"`      // Policing Po E Enabled
	UpoePlusEnabled         *bool    `json:"upoePlusEnabled,omitempty"`         // Upoe Plus Enabled
	FourPairEnabled         *bool    `json:"fourPairEnabled,omitempty"`         // Four Pair Enabled
	PoeDataTimestamp        *int     `json:"poeDataTimestamp,omitempty"`        // Poe Data Timestamp
	PdLocation              string   `json:"pdLocation,omitempty"`              // Pd Location
	PdDeviceName            string   `json:"pdDeviceName,omitempty"`            // Pd Device Name
	PdConnectedSwitch       string   `json:"pdConnectedSwitch,omitempty"`       // Pd Connected Switch
	ConnectedSwitchUUID     string   `json:"connectedSwitchUuid,omitempty"`     // Connected Switch Uuid
	IeeeCompliant           *bool    `json:"ieeeCompliant,omitempty"`           // Ieee Compliant
	ConnectedSwitchType     string   `json:"connectedSwitchType,omitempty"`     // Connected Switch Type
	SiteName                string   `json:"siteName,omitempty"`
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesPage struct {
	Limit  *int                                                                                         `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                         `json:"offset,omitempty"` // Offset
	Count  *int                                                                                         `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount struct {
	Response *ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountResponse `json:"response,omitempty"` //
	Version  string                                                                                                                                                           `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	Response *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage       `json:"page,omitempty"`     //
	Version  string                                                                                                                          `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse struct {
	ID                      string                                                                                                                                             `json:"id,omitempty"`                      // Id
	AdminStatus             string                                                                                                                                             `json:"adminStatus,omitempty"`             // Admin Status
	Description             string                                                                                                                                             `json:"description,omitempty"`             // Description
	DuplexConfig            string                                                                                                                                             `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string                                                                                                                                             `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int                                                                                                                                               `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string                                                                                                                                             `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string                                                                                                                                             `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string                                                                                                                                           `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool                                                                                                                                              `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool                                                                                                                                              `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string                                                                                                                                             `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string                                                                                                                                             `json:"mediaType,omitempty"`               // Media Type
	Name                    string                                                                                                                                             `json:"name,omitempty"`                    // Name
	OperStatus              string                                                                                                                                             `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int                                                                                                                                               `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string                                                                                                                                             `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string                                                                                                                                             `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string                                                                                                                                             `json:"portMode,omitempty"`                // Port Mode
	PortType                string                                                                                                                                             `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64                                                                                                                                           `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int                                                                                                                                               `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64                                                                                                                                           `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64                                                                                                                                           `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string                                                                                                                                             `json:"speed,omitempty"`                   // Speed
	StackPortType           string                                                                                                                                             `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int                                                                                                                                               `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64                                                                                                                                           `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int                                                                                                                                               `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64                                                                                                                                           `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64                                                                                                                                           `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string                                                                                                                                             `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string                                                                                                                                             `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string                                                                                                                                             `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string                                                                                                                                             `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteName                string                                                                                                                                             `json:"siteName,omitempty"`                // Site Name
	SiteHierarchy           string                                                                                                                                             `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string                                                                                                                                             `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	AggregateAttributes     *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`     //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes struct {
	Name   string                                                                                                                                                   `json:"name,omitempty"`   // Name
	Values *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributesValues `json:"values,omitempty"` //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributesValues struct {
	Key   string   `json:"key,omitempty"`   // Key
	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int                                                                                                                              `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                              `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                                              `json:"count,omitempty"`  // Count
	SortBy *[]ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices struct {
	Response *ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData struct {
	Response *ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataResponse `json:"response,omitempty"` //
	Version  string                                                                                                        `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataResponse struct {
	ID                      string   `json:"id,omitempty"`                      // Id
	AdminStatus             string   `json:"adminStatus,omitempty"`             // Admin Status
	Description             string   `json:"description,omitempty"`             // Description
	DuplexConfig            string   `json:"duplexConfig,omitempty"`            // Duplex Config
	DuplexOper              string   `json:"duplexOper,omitempty"`              // Duplex Oper
	InterfaceIfIndex        *int     `json:"interfaceIfIndex,omitempty"`        // Interface If Index
	InterfaceType           string   `json:"interfaceType,omitempty"`           // Interface Type
	IPv4Address             string   `json:"ipv4Address,omitempty"`             // Ipv4 Address
	IPv6AddressList         []string `json:"ipv6AddressList,omitempty"`         // Ipv6 Address List
	IsL3Interface           *bool    `json:"isL3Interface,omitempty"`           // Is L3 Interface
	IsWan                   *bool    `json:"isWan,omitempty"`                   // Is Wan
	MacAddr                 string   `json:"macAddr,omitempty"`                 // Mac Addr
	MediaType               string   `json:"mediaType,omitempty"`               // Media Type
	Name                    string   `json:"name,omitempty"`                    // Name
	OperStatus              string   `json:"operStatus,omitempty"`              // Oper Status
	PeerStackMember         *int     `json:"peerStackMember,omitempty"`         // Peer Stack Member
	PeerStackPort           string   `json:"peerStackPort,omitempty"`           // Peer Stack Port
	PortChannelID           string   `json:"portChannelId,omitempty"`           // Port Channel Id
	PortMode                string   `json:"portMode,omitempty"`                // Port Mode
	PortType                string   `json:"portType,omitempty"`                // Port Type
	RxDiscards              *float64 `json:"rxDiscards,omitempty"`              // Rx Discards
	RxError                 *int     `json:"rxError,omitempty"`                 // Rx Error
	RxRate                  *float64 `json:"rxRate,omitempty"`                  // Rx Rate
	RxUtilization           *float64 `json:"rxUtilization,omitempty"`           // Rx Utilization
	Speed                   string   `json:"speed,omitempty"`                   // Speed
	StackPortType           string   `json:"stackPortType,omitempty"`           // Stack Port Type
	Timestamp               *int     `json:"timestamp,omitempty"`               // Timestamp
	TxDiscards              *float64 `json:"txDiscards,omitempty"`              // Tx Discards
	TxError                 *int     `json:"txError,omitempty"`                 // Tx Error
	TxRate                  *float64 `json:"txRate,omitempty"`                  // Tx Rate
	TxUtilization           *float64 `json:"txUtilization,omitempty"`           // Tx Utilization
	VLANID                  string   `json:"vlanId,omitempty"`                  // Vlan Id
	NetworkDeviceID         string   `json:"networkDeviceId,omitempty"`         // Network Device Id
	NetworkDeviceIPAddress  string   `json:"networkDeviceIpAddress,omitempty"`  // Network Device Ip Address
	NetworkDeviceMacAddress string   `json:"networkDeviceMacAddress,omitempty"` // Network Device Mac Address
	SiteHierarchy           string   `json:"siteHierarchy,omitempty"`           // Site Hierarchy
	SiteHierarchyID         string   `json:"siteHierarchyId,omitempty"`         // Site Hierarchy Id
	PoeAdminStatus          string   `json:"poeAdminStatus,omitempty"`          // Poe Admin Status
	PoeOperStatus           string   `json:"poeOperStatus,omitempty"`           // Poe Oper Status
	ChassisID               *int     `json:"chassisId,omitempty"`               // Chassis Id
	ModuleID                *int     `json:"moduleId,omitempty"`                // Module Id
	PdClassSignal           string   `json:"pdClassSignal,omitempty"`           // Pd Class Signal
	PdClassSpare            string   `json:"pdClassSpare,omitempty"`            // Pd Class Spare
	PdDeviceType            string   `json:"pdDeviceType,omitempty"`            // Pd Device Type
	PdDeviceModel           string   `json:"pdDeviceModel,omitempty"`           // Pd Device Model
	PdPowerAdminMaxInWatt   string   `json:"pdPowerAdminMaxInWatt,omitempty"`   // Pd Power Admin Max In Watt
	PdPowerBudgetInWatt     string   `json:"pdPowerBudgetInWatt,omitempty"`     // Pd Power Budget In Watt
	PdPowerConsumedInWatt   string   `json:"pdPowerConsumedInWatt,omitempty"`   // Pd Power Consumed In Watt
	PdPowerRemainingInWatt  string   `json:"pdPowerRemainingInWatt,omitempty"`  // Pd Power Remaining In Watt
	PdMaxPowerDrawn         string   `json:"pdMaxPowerDrawn,omitempty"`         // Pd Max Power Drawn
	PdConnectedDeviceList   []string `json:"pdConnectedDeviceList,omitempty"`   // Pd Connected Device List
	PoeOperPriority         string   `json:"poeOperPriority,omitempty"`         // Poe Oper Priority
	FastPoEEnabled          *bool    `json:"fastPoEEnabled,omitempty"`          // Fast Po E Enabled
	PerpetualPoEEnabled     *bool    `json:"perpetualPoEEnabled,omitempty"`     // Perpetual Po E Enabled
	PolicingPoEEnabled      *bool    `json:"policingPoEEnabled,omitempty"`      // Policing Po E Enabled
	UpoePlusEnabled         *bool    `json:"upoePlusEnabled,omitempty"`         // Upoe Plus Enabled
	FourPairEnabled         *bool    `json:"fourPairEnabled,omitempty"`         // Four Pair Enabled
	PoeDataTimestamp        *int     `json:"poeDataTimestamp,omitempty"`        // Poe Data Timestamp
	PdLocation              string   `json:"pdLocation,omitempty"`              // Pd Location
	PdDeviceName            string   `json:"pdDeviceName,omitempty"`            // Pd Device Name
	PdConnectedSwitch       string   `json:"pdConnectedSwitch,omitempty"`       // Pd Connected Switch
	ConnectedSwitchUUID     string   `json:"connectedSwitchUuid,omitempty"`     // Connected Switch Uuid
	IeeeCompliant           *bool    `json:"ieeeCompliant,omitempty"`           // Ieee Compliant
	ConnectedSwitchType     string   `json:"connectedSwitchType,omitempty"`     // Connected Switch Type
	SiteName                string   `json:"siteName,omitempty"`
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange struct {
	Response       *[]ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponse `json:"response,omitempty"`       //
	TimestampOrder string                                                                                 `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponse struct {
	Timestamp           *int                                                                                                      `json:"timestamp,omitempty"`           // Timestamp
	Attributes          *[]ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters struct {
	Response *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersPage       `json:"page,omitempty"`     //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponse struct {
	ID                         string                                                                                                    `json:"id,omitempty"`                         // Id
	Name                       string                                                                                                    `json:"name,omitempty"`                       // Name
	ManagementIPAddress        string                                                                                                    `json:"managementIpAddress,omitempty"`        // Management Ip Address
	PlatformID                 string                                                                                                    `json:"platformId,omitempty"`                 // Platform Id
	DeviceFamily               string                                                                                                    `json:"deviceFamily,omitempty"`               // Device Family
	SerialNumber               string                                                                                                    `json:"serialNumber,omitempty"`               // Serial Number
	MacAddress                 string                                                                                                    `json:"macAddress,omitempty"`                 // Mac Address
	DeviceSeries               string                                                                                                    `json:"deviceSeries,omitempty"`               // Device Series
	SoftwareVersion            string                                                                                                    `json:"softwareVersion,omitempty"`            // Software Version
	ProductVendor              string                                                                                                    `json:"productVendor,omitempty"`              // Product Vendor
	DeviceRole                 string                                                                                                    `json:"deviceRole,omitempty"`                 // Device Role
	DeviceType                 string                                                                                                    `json:"deviceType,omitempty"`                 // Device Type
	CommunicationState         string                                                                                                    `json:"communicationState,omitempty"`         // Communication State
	CollectionStatus           string                                                                                                    `json:"collectionStatus,omitempty"`           // Collection Status
	HaStatus                   string                                                                                                    `json:"haStatus,omitempty"`                   // Ha Status
	LastBootTime               *int                                                                                                      `json:"lastBootTime,omitempty"`               // Last Boot Time
	SiteHierarchyID            string                                                                                                    `json:"siteHierarchyId,omitempty"`            // Site Hierarchy Id
	SiteHierarchy              string                                                                                                    `json:"siteHierarchy,omitempty"`              // Site Hierarchy
	SiteID                     string                                                                                                    `json:"siteId,omitempty"`                     // Site Id
	DeviceGroupHierarchyID     string                                                                                                    `json:"deviceGroupHierarchyId,omitempty"`     // Device Group Hierarchy Id
	TagNames                   []string                                                                                                  `json:"tagNames,omitempty"`                   // Tag Names
	StackType                  string                                                                                                    `json:"stackType,omitempty"`                  // Stack Type
	OsType                     string                                                                                                    `json:"osType,omitempty"`                     // Os Type
	RingStatus                 *bool                                                                                                     `json:"ringStatus,omitempty"`                 // Ring Status
	MaintenanceModeEnabled     *bool                                                                                                     `json:"maintenanceModeEnabled,omitempty"`     // Maintenance Mode Enabled
	UpTime                     *int                                                                                                      `json:"upTime,omitempty"`                     // Up Time
	IPv4Address                string                                                                                                    `json:"ipv4Address,omitempty"`                // Ipv4 Address
	IPv6Address                string                                                                                                    `json:"ipv6Address,omitempty"`                // Ipv6 Address
	RedundancyMode             string                                                                                                    `json:"redundancyMode,omitempty"`             // Redundancy Mode
	FeatureFlagList            []string                                                                                                  `json:"featureFlagList,omitempty"`            // Feature Flag List
	HaLastResetReason          string                                                                                                    `json:"haLastResetReason,omitempty"`          // Ha Last Reset Reason
	RedundancyPeerStateDerived string                                                                                                    `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived
	RedundancyPeerState        string                                                                                                    `json:"redundancyPeerState,omitempty"`        // Redundancy Peer State
	RedundancyStateDerived     string                                                                                                    `json:"redundancyStateDerived,omitempty"`     // Redundancy State Derived
	RedundancyState            string                                                                                                    `json:"redundancyState,omitempty"`            // Redundancy State
	WiredClientCount           *int                                                                                                      `json:"wiredClientCount,omitempty"`           // Wired Client Count
	WirelessClientCount        *int                                                                                                      `json:"wirelessClientCount,omitempty"`        // Wireless Client Count
	PortCount                  *int                                                                                                      `json:"portCount,omitempty"`                  // Port Count
	PhysicalPortCount          *int                                                                                                      `json:"physicalPortCount,omitempty"`          // Physical Port Count
	VirtualPortCount           *int                                                                                                      `json:"virtualPortCount,omitempty"`           // Virtual Port Count
	ClientCount                *int                                                                                                      `json:"clientCount,omitempty"`                // Client Count
	ApDetails                  *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetails             `json:"apDetails,omitempty"`                  //
	MetricsDetails             *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseMetricsDetails        `json:"metricsDetails,omitempty"`             //
	FabricDetails              *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricDetails         `json:"fabricDetails,omitempty"`              //
	SwitchPoeDetails           *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseSwitchPoeDetails      `json:"switchPoeDetails,omitempty"`           //
	FabricMetricsDetails       *ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricMetricsDetails  `json:"fabricMetricsDetails,omitempty"`       //
	AggregateAttributes        *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`        //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetails struct {
	ConnectedWlcName     string                                                                                                `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                                                `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                                                `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                                                `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                                                `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                                                `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                                                `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                                                `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                                                  `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                                                 `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                                                  `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                                                `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                                                 `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                                                `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                                                `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                                                `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                                                `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                                                `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                                                `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                                                `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                                                `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                                                `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                                                `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                                                `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                                                `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                                                `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"`                 // Overall Health Score
	CPUUtilization                     *float64 `json:"cpuUtilization,omitempty"`                     // Cpu Utilization
	CPUScore                           *int     `json:"cpuScore,omitempty"`                           // Cpu Score
	MemoryUtilization                  *float64 `json:"memoryUtilization,omitempty"`                  // Memory Utilization
	MemoryScore                        *int     `json:"memoryScore,omitempty"`                        // Memory Score
	AvgTemperature                     *float64 `json:"avgTemperature,omitempty"`                     // Avg Temperature
	MaxTemperature                     *float64 `json:"maxTemperature,omitempty"`                     // Max Temperature
	DiscardScore                       *int     `json:"discardScore,omitempty"`                       // Discard Score
	DiscardInterfaces                  []string `json:"discardInterfaces,omitempty"`                  // Discard Interfaces
	ErrorScore                         *int     `json:"errorScore,omitempty"`                         // Error Score
	ErrorInterfaces                    []string `json:"errorInterfaces,omitempty"`                    // Error Interfaces
	InterDeviceLinkScore               *int     `json:"interDeviceLinkScore,omitempty"`               // Inter Device Link Score
	InterDeviceConnectedDownInterfaces []string `json:"interDeviceConnectedDownInterfaces,omitempty"` // Inter Device Connected Down Interfaces
	LinkUtilizationScore               *int     `json:"linkUtilizationScore,omitempty"`               // Link Utilization Score
	HighLinkUtilizationInterfaces      []string `json:"highLinkUtilizationInterfaces,omitempty"`      // High Link Utilization Interfaces
	FreeTimerScore                     *int     `json:"freeTimerScore,omitempty"`                     // Free Timer Score
	FreeTimer                          *float64 `json:"freeTimer,omitempty"`                          // Free Timer
	PacketPoolScore                    *int     `json:"packetPoolScore,omitempty"`                    // Packet Pool Score
	PacketPool                         *int     `json:"packetPool,omitempty"`                         // Packet Pool
	FreeMemoryBufferScore              *int     `json:"freeMemoryBufferScore,omitempty"`              // Free Memory Buffer Score
	FreeMemoryBuffer                   *float64 `json:"freeMemoryBuffer,omitempty"`                   // Free Memory Buffer
	WqePoolScore                       *int     `json:"wqePoolScore,omitempty"`                       // Wqe Pool Score
	WqePool                            *int     `json:"wqePool,omitempty"`                            // Wqe Pool
	ApCount                            *int     `json:"apCount,omitempty"`                            // Ap Count
	NoiseScore                         *int     `json:"noiseScore,omitempty"`                         // Noise Score
	UtilizationScore                   *int     `json:"utilizationScore,omitempty"`                   // Utilization Score
	InterferenceScore                  *int     `json:"interferenceScore,omitempty"`                  // Interference Score
	AirQualityScore                    *int     `json:"airQualityScore,omitempty"`                    // Air Quality Score
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricDetails struct {
	FabricRole      []string `json:"fabricRole,omitempty"`      // Fabric Role
	FabricSiteName  string   `json:"fabricSiteName,omitempty"`  // Fabric Site Name
	TransitFabrics  []string `json:"transitFabrics,omitempty"`  // Transit Fabrics
	L2Vns           []string `json:"l2Vns,omitempty"`           // L2 Vns
	L3Vns           []string `json:"l3Vns,omitempty"`           // L3 Vns
	FabricSiteID    string   `json:"fabricSiteId,omitempty"`    // Fabric Site Id
	NetworkProtocol string   `json:"networkProtocol,omitempty"` // Network Protocol
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseSwitchPoeDetails struct {
	PortCount            *int                                                                                                                `json:"portCount,omitempty"`            // Port Count
	UsedPortCount        *int                                                                                                                `json:"usedPortCount,omitempty"`        // Used Port Count
	FreePortCount        *int                                                                                                                `json:"freePortCount,omitempty"`        // Free Port Count
	PowerConsumed        *float64                                                                                                            `json:"powerConsumed,omitempty"`        // Power Consumed
	PoePowerConsumed     *int                                                                                                                `json:"poePowerConsumed,omitempty"`     // Poe Power Consumed
	SystemPowerConsumed  *float64                                                                                                            `json:"systemPowerConsumed,omitempty"`  // System Power Consumed
	PowerBudget          *int                                                                                                                `json:"powerBudget,omitempty"`          // Power Budget
	PoePowerAllocated    *float64                                                                                                            `json:"poePowerAllocated,omitempty"`    // Poe Power Allocated
	SystemPowerAllocated *int                                                                                                                `json:"systemPowerAllocated,omitempty"` // System Power Allocated
	PowerRemaining       *float64                                                                                                            `json:"powerRemaining,omitempty"`       // Power Remaining
	PoeVersion           string                                                                                                              `json:"poeVersion,omitempty"`           // Poe Version
	ChassisCount         *int                                                                                                                `json:"chassisCount,omitempty"`         // Chassis Count
	ModuleCount          *int                                                                                                                `json:"moduleCount,omitempty"`          // Module Count
	ModuleDetails        *[]ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseSwitchPoeDetailsModuleDetails `json:"moduleDetails,omitempty"`        //
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseSwitchPoeDetailsModuleDetails struct {
	ModuleID                   string   `json:"moduleId,omitempty"`                   // Module Id
	ChassisID                  string   `json:"chassisId,omitempty"`                  // Chassis Id
	ModulePortCount            *int     `json:"modulePortCount,omitempty"`            // Module Port Count
	ModuleUsedPortCount        *int     `json:"moduleUsedPortCount,omitempty"`        // Module Used Port Count
	ModuleFreePortCount        *int     `json:"moduleFreePortCount,omitempty"`        // Module Free Port Count
	ModulePowerConsumed        *float64 `json:"modulePowerConsumed,omitempty"`        // Module Power Consumed
	ModulePoePowerConsumed     *int     `json:"modulePoePowerConsumed,omitempty"`     // Module Poe Power Consumed
	ModuleSystemPowerConsumed  *float64 `json:"moduleSystemPowerConsumed,omitempty"`  // Module System Power Consumed
	ModulePowerBudget          *int     `json:"modulePowerBudget,omitempty"`          // Module Power Budget
	ModulePoePowerAllocated    *float64 `json:"modulePoePowerAllocated,omitempty"`    // Module Poe Power Allocated
	ModuleSystemPowerAllocated *int     `json:"moduleSystemPowerAllocated,omitempty"` // Module System Power Allocated
	ModulePowerRemaining       *float64 `json:"modulePowerRemaining,omitempty"`       // Module Power Remaining
	InterfacePowerMax          *int     `json:"interfacePowerMax,omitempty"`          // Interface Power Max
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseFabricMetricsDetails struct {
	OverallFabricScore       *int `json:"overallFabricScore,omitempty"`       // Overall Fabric Score
	FabricTransitScore       *int `json:"fabricTransitScore,omitempty"`       // Fabric Transit Score
	FabricSiteScore          *int `json:"fabricSiteScore,omitempty"`          // Fabric Site Score
	FabricVnScore            *int `json:"fabricVnScore,omitempty"`            // Fabric Vn Score
	FabsiteFcpScore          *int `json:"fabsiteFcpScore,omitempty"`          // Fabsite Fcp Score
	FabsiteInfraScore        *int `json:"fabsiteInfraScore,omitempty"`        // Fabsite Infra Score
	FabsiteFsconnScore       *int `json:"fabsiteFsconnScore,omitempty"`       // Fabsite Fsconn Score
	VnExitScore              *int `json:"vnExitScore,omitempty"`              // Vn Exit Score
	VnFcpScore               *int `json:"vnFcpScore,omitempty"`               // Vn Fcp Score
	VnStatusScore            *int `json:"vnStatusScore,omitempty"`            // Vn Status Score
	VnServiceScore           *int `json:"vnServiceScore,omitempty"`           // Vn Service Score
	TransitControlPlaneScore *int `json:"transitControlPlaneScore,omitempty"` // Transit Control Plane Score
	TransitServicesScore     *int `json:"transitServicesScore,omitempty"`     // Transit Services Score
	TCPConnScore             *int `json:"tcpConnScore,omitempty"`             // Tcp Conn Score
	BgpBgpSiteScore          *int `json:"bgpBgpSiteScore,omitempty"`          // Bgp Bgp Site Score
	VniStatusScore           *int `json:"vniStatusScore,omitempty"`           // Vni Status Score
	PubsubTransitConnScore   *int `json:"pubsubTransitConnScore,omitempty"`   // Pubsub Transit Conn Score
	BgpPeerInfraVnScore      *int `json:"bgpPeerInfraVnScore,omitempty"`      // Bgp Peer Infra Vn Score
	InternetAvailScore       *int `json:"internetAvailScore,omitempty"`       // Internet Avail Score
	BgpEvpnScore             *int `json:"bgpEvpnScore,omitempty"`             // Bgp Evpn Score
	LispTransitConnScore     *int `json:"lispTransitConnScore,omitempty"`     // Lisp Transit Conn Score
	CtsEnvDataDownloadScore  *int `json:"ctsEnvDataDownloadScore,omitempty"`  // Cts Env Data Download Score
	PubsubInfraVnScore       *int `json:"pubsubInfraVnScore,omitempty"`       // Pubsub Infra Vn Score
	PeerScore                *int `json:"peerScore,omitempty"`                // Peer Score
	BgpPeerScore             *int `json:"bgpPeerScore,omitempty"`             // Bgp Peer Score
	RemoteInternetAvailScore *int `json:"remoteInternetAvailScore,omitempty"` // Remote Internet Avail Score
	BgpTCPScore              *int `json:"bgpTcpScore,omitempty"`              // Bgp Tcp Score
	PubsubSessionScore       *int `json:"pubsubSessionScore,omitempty"`       // Pubsub Session Score
	AAAStatusScore           *int `json:"aaaStatusScore,omitempty"`           // Aaa Status Score
	LispCpConnScore          *int `json:"lispCpConnScore,omitempty"`          // Lisp Cp Conn Score
	BgpPubsubSiteScore       *int `json:"bgpPubsubSiteScore,omitempty"`       // Bgp Pubsub Site Score
	McastScore               *int `json:"mcastScore,omitempty"`               // Mcast Score
	PortChannelScore         *int `json:"portChannelScore,omitempty"`         // Port Channel Score
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersPage struct {
	Limit  *int   `json:"limit,omitempty"`  // Limit
	Offset *int   `json:"offset,omitempty"` // Offset
	Count  *int   `json:"count,omitempty"`  // Count
	SortBy string `json:"sortBy,omitempty"` // Sort By
	Order  string `json:"order,omitempty"`  // Order
}
type ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters struct {
	Response *ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	Response *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse `json:"response,omitempty"` //
	Page     *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage       `json:"page,omitempty"`     //
	Version  string                                                                                                       `json:"version,omitempty"`  // Version
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Management Ip Address

	PlatformID string `json:"platformId,omitempty"` // Platform Id

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number

	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	DeviceSeries string `json:"deviceSeries,omitempty"` // Device Series

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software Version

	ProductVendor string `json:"productVendor,omitempty"` // Product Vendor

	DeviceRole string `json:"deviceRole,omitempty"` // Device Role

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	CommunicationState string `json:"communicationState,omitempty"` // Communication State

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection Status

	HaStatus string `json:"haStatus,omitempty"` // Ha Status

	LastBootTime *int `json:"lastBootTime,omitempty"` // Last Boot Time

	SiteHierarchyID string `json:"siteHierarchyId,omitempty"` // Site Hierarchy Id

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy

	SiteID string `json:"siteId,omitempty"` // Site Id

	DeviceGroupHierarchyID string `json:"deviceGroupHierarchyId,omitempty"` // Device Group Hierarchy Id

	TagNames []string `json:"tagNames,omitempty"` // Tag Names

	StackType string `json:"stackType,omitempty"` // Stack Type

	OsType string `json:"osType,omitempty"` // Os Type

	RingStatus *bool `json:"ringStatus,omitempty"` // Ring Status

	MaintenanceModeEnabled *bool `json:"maintenanceModeEnabled,omitempty"` // Maintenance Mode Enabled

	UpTime *int `json:"upTime,omitempty"` // Up Time

	IPv4Address string `json:"ipv4Address,omitempty"` // Ipv4 Address

	IPv6Address string `json:"ipv6Address,omitempty"` // Ipv6 Address

	RedundancyMode string `json:"redundancyMode,omitempty"` // Redundancy Mode

	FeatureFlagList []string `json:"featureFlagList,omitempty"` // Feature Flag List

	HaLastResetReason string `json:"haLastResetReason,omitempty"` // Ha Last Reset Reason

	RedundancyPeerStateDerived string `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived

	RedundancyPeerState string `json:"redundancyPeerState,omitempty"` // Redundancy Peer State

	RedundancyStateDerived string `json:"redundancyStateDerived,omitempty"` // Redundancy State Derived

	RedundancyState string `json:"redundancyState,omitempty"` // Redundancy State

	WiredClientCount *int `json:"wiredClientCount,omitempty"` // Wired Client Count

	WirelessClientCount *int `json:"wirelessClientCount,omitempty"` // Wireless Client Count

	PortCount *int `json:"portCount,omitempty"` // Port Count

	PhysicalPortCount *int `json:"physicalPortCount,omitempty"` // Physical Port Count

	VirtualPortCount *int `json:"virtualPortCount,omitempty"` // Virtual Port Count

	ClientCount *int `json:"clientCount,omitempty"` // Client Count

	ApDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetails `json:"apDetails,omitempty"` //

	MetricsDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseMetricsDetails `json:"metricsDetails,omitempty"` //

	FabricDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricDetails `json:"fabricDetails,omitempty"` //

	SwitchPoeDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseSwitchPoeDetails `json:"switchPoeDetails,omitempty"` //

	FabricMetricsDetails *ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricMetricsDetails `json:"fabricMetricsDetails,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetails struct {
	ConnectedWlcName     string                                                                                                                      `json:"connectedWlcName,omitempty"`     // Connected Wlc Name
	PolicyTagName        string                                                                                                                      `json:"policyTagName,omitempty"`        // Policy Tag Name
	ApOperationalState   string                                                                                                                      `json:"apOperationalState,omitempty"`   // Ap Operational State
	PowerSaveMode        string                                                                                                                      `json:"powerSaveMode,omitempty"`        // Power Save Mode
	OperationalMode      string                                                                                                                      `json:"operationalMode,omitempty"`      // Operational Mode
	ResetReason          string                                                                                                                      `json:"resetReason,omitempty"`          // Reset Reason
	Protocol             string                                                                                                                      `json:"protocol,omitempty"`             // Protocol
	PowerMode            string                                                                                                                      `json:"powerMode,omitempty"`            // Power Mode
	ConnectedTime        *int                                                                                                                        `json:"connectedTime,omitempty"`        // Connected Time
	LedFlashEnabled      *bool                                                                                                                       `json:"ledFlashEnabled,omitempty"`      // Led Flash Enabled
	LedFlashSeconds      *int                                                                                                                        `json:"ledFlashSeconds,omitempty"`      // Led Flash Seconds
	SubMode              string                                                                                                                      `json:"subMode,omitempty"`              // Sub Mode
	HomeApEnabled        *bool                                                                                                                       `json:"homeApEnabled,omitempty"`        // Home Ap Enabled
	PowerType            string                                                                                                                      `json:"powerType,omitempty"`            // Power Type
	ApType               string                                                                                                                      `json:"apType,omitempty"`               // Ap Type
	AdminState           string                                                                                                                      `json:"adminState,omitempty"`           // Admin State
	IcapCapability       string                                                                                                                      `json:"icapCapability,omitempty"`       // Icap Capability
	RegulatoryDomain     string                                                                                                                      `json:"regulatoryDomain,omitempty"`     // Regulatory Domain
	EthernetMac          string                                                                                                                      `json:"ethernetMac,omitempty"`          // Ethernet Mac
	RfTagName            string                                                                                                                      `json:"rfTagName,omitempty"`            // Rf Tag Name
	SiteTagName          string                                                                                                                      `json:"siteTagName,omitempty"`          // Site Tag Name
	PowerSaveModeCapable string                                                                                                                      `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable
	PowerProfile         string                                                                                                                      `json:"powerProfile,omitempty"`         // Power Profile
	FlexGroup            string                                                                                                                      `json:"flexGroup,omitempty"`            // Flex Group
	PowerCalendarProfile string                                                                                                                      `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile
	ApGroup              string                                                                                                                      `json:"apGroup,omitempty"`              // Ap Group
	Radios               *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetailsRadios `json:"radios,omitempty"`               //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseApDetailsRadios struct {
	ID           string   `json:"id,omitempty"`           // Id
	Band         string   `json:"band,omitempty"`         // Band
	Noise        *int     `json:"noise,omitempty"`        // Noise
	AirQuality   *float64 `json:"airQuality,omitempty"`   // Air Quality
	Interference *float64 `json:"interference,omitempty"` // Interference
	TrafficUtil  *int     `json:"trafficUtil,omitempty"`  // Traffic Util
	Utilization  *float64 `json:"utilization,omitempty"`  // Utilization
	ClientCount  *int     `json:"clientCount,omitempty"`  // Client Count
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseMetricsDetails struct {
	OverallHealthScore                 *int     `json:"overallHealthScore,omitempty"`                 // Overall Health Score
	CPUUtilization                     *float64 `json:"cpuUtilization,omitempty"`                     // Cpu Utilization
	CPUScore                           *int     `json:"cpuScore,omitempty"`                           // Cpu Score
	MemoryUtilization                  *float64 `json:"memoryUtilization,omitempty"`                  // Memory Utilization
	MemoryScore                        *int     `json:"memoryScore,omitempty"`                        // Memory Score
	AvgTemperature                     *float64 `json:"avgTemperature,omitempty"`                     // Avg Temperature
	MaxTemperature                     *float64 `json:"maxTemperature,omitempty"`                     // Max Temperature
	DiscardScore                       *int     `json:"discardScore,omitempty"`                       // Discard Score
	DiscardInterfaces                  []string `json:"discardInterfaces,omitempty"`                  // Discard Interfaces
	ErrorScore                         *int     `json:"errorScore,omitempty"`                         // Error Score
	ErrorInterfaces                    []string `json:"errorInterfaces,omitempty"`                    // Error Interfaces
	InterDeviceLinkScore               *int     `json:"interDeviceLinkScore,omitempty"`               // Inter Device Link Score
	InterDeviceConnectedDownInterfaces []string `json:"interDeviceConnectedDownInterfaces,omitempty"` // Inter Device Connected Down Interfaces
	LinkUtilizationScore               *int     `json:"linkUtilizationScore,omitempty"`               // Link Utilization Score
	HighLinkUtilizationInterfaces      []string `json:"highLinkUtilizationInterfaces,omitempty"`      // High Link Utilization Interfaces
	FreeTimerScore                     *int     `json:"freeTimerScore,omitempty"`                     // Free Timer Score
	FreeTimer                          *float64 `json:"freeTimer,omitempty"`                          // Free Timer
	PacketPoolScore                    *int     `json:"packetPoolScore,omitempty"`                    // Packet Pool Score
	PacketPool                         *int     `json:"packetPool,omitempty"`                         // Packet Pool
	FreeMemoryBufferScore              *int     `json:"freeMemoryBufferScore,omitempty"`              // Free Memory Buffer Score
	FreeMemoryBuffer                   *float64 `json:"freeMemoryBuffer,omitempty"`                   // Free Memory Buffer
	WqePoolScore                       *int     `json:"wqePoolScore,omitempty"`                       // Wqe Pool Score
	WqePool                            *int     `json:"wqePool,omitempty"`                            // Wqe Pool
	ApCount                            *int     `json:"apCount,omitempty"`                            // Ap Count
	NoiseScore                         *int     `json:"noiseScore,omitempty"`                         // Noise Score
	UtilizationScore                   *int     `json:"utilizationScore,omitempty"`                   // Utilization Score
	InterferenceScore                  *int     `json:"interferenceScore,omitempty"`                  // Interference Score
	AirQualityScore                    *int     `json:"airQualityScore,omitempty"`                    // Air Quality Score
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricDetails struct {
	FabricRole []string `json:"fabricRole,omitempty"` // Fabric Role

	FabricSiteName string `json:"fabricSiteName,omitempty"` // Fabric Site Name

	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics

	L2Vns []string `json:"l2Vns,omitempty"` // L2 Vns

	L3Vns []string `json:"l3Vns,omitempty"` // L3 Vns

	FabricSiteID string `json:"fabricSiteId,omitempty"` // Fabric Site Id

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseSwitchPoeDetails struct {
	PortCount *int `json:"portCount,omitempty"` // Port Count

	UsedPortCount *int `json:"usedPortCount,omitempty"` // Used Port Count

	FreePortCount *int `json:"freePortCount,omitempty"` // Free Port Count

	PowerConsumed *float64 `json:"powerConsumed,omitempty"` // Power Consumed

	PoePowerConsumed *int `json:"poePowerConsumed,omitempty"` // Poe Power Consumed

	SystemPowerConsumed *float64 `json:"systemPowerConsumed,omitempty"` // System Power Consumed

	PowerBudget *int `json:"powerBudget,omitempty"` // Power Budget

	PoePowerAllocated *float64 `json:"poePowerAllocated,omitempty"` // Poe Power Allocated

	SystemPowerAllocated *int `json:"systemPowerAllocated,omitempty"` // System Power Allocated

	PowerRemaining *float64 `json:"powerRemaining,omitempty"` // Power Remaining

	PoeVersion string `json:"poeVersion,omitempty"` // Poe Version

	ChassisCount *int `json:"chassisCount,omitempty"` // Chassis Count

	ModuleCount *int `json:"moduleCount,omitempty"` // Module Count

	ModuleDetails *[]ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseSwitchPoeDetailsModuleDetails `json:"moduleDetails,omitempty"` //
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseSwitchPoeDetailsModuleDetails struct {
	ModuleID string `json:"moduleId,omitempty"` // Module Id

	ChassisID string `json:"chassisId,omitempty"` // Chassis Id

	ModulePortCount *int `json:"modulePortCount,omitempty"` // Module Port Count

	ModuleUsedPortCount *int `json:"moduleUsedPortCount,omitempty"` // Module Used Port Count

	ModuleFreePortCount *int `json:"moduleFreePortCount,omitempty"` // Module Free Port Count

	ModulePowerConsumed *float64 `json:"modulePowerConsumed,omitempty"` // Module Power Consumed

	ModulePoePowerConsumed *int `json:"modulePoePowerConsumed,omitempty"` // Module Poe Power Consumed

	ModuleSystemPowerConsumed *float64 `json:"moduleSystemPowerConsumed,omitempty"` // Module System Power Consumed

	ModulePowerBudget *int `json:"modulePowerBudget,omitempty"` // Module Power Budget

	ModulePoePowerAllocated *float64 `json:"modulePoePowerAllocated,omitempty"` // Module Poe Power Allocated

	ModuleSystemPowerAllocated *int `json:"moduleSystemPowerAllocated,omitempty"` // Module System Power Allocated

	ModulePowerRemaining *float64 `json:"modulePowerRemaining,omitempty"` // Module Power Remaining

	InterfacePowerMax *int `json:"interfacePowerMax,omitempty"` // Interface Power Max
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseFabricMetricsDetails struct {
	OverallFabricScore *int `json:"overallFabricScore,omitempty"` // Overall Fabric Score

	FabricTransitScore *int `json:"fabricTransitScore,omitempty"` // Fabric Transit Score

	FabricSiteScore *int `json:"fabricSiteScore,omitempty"` // Fabric Site Score

	FabricVnScore *int `json:"fabricVnScore,omitempty"` // Fabric Vn Score

	FabsiteFcpScore *int `json:"fabsiteFcpScore,omitempty"` // Fabsite Fcp Score

	FabsiteInfraScore *int `json:"fabsiteInfraScore,omitempty"` // Fabsite Infra Score

	FabsiteFsconnScore *int `json:"fabsiteFsconnScore,omitempty"` // Fabsite Fsconn Score

	VnExitScore *int `json:"vnExitScore,omitempty"` // Vn Exit Score

	VnFcpScore *int `json:"vnFcpScore,omitempty"` // Vn Fcp Score

	VnStatusScore *int `json:"vnStatusScore,omitempty"` // Vn Status Score

	VnServiceScore *int `json:"vnServiceScore,omitempty"` // Vn Service Score

	TransitControlPlaneScore *int `json:"transitControlPlaneScore,omitempty"` // Transit Control Plane Score

	TransitServicesScore *int `json:"transitServicesScore,omitempty"` // Transit Services Score

	TCPConnScore *int `json:"tcpConnScore,omitempty"` // Tcp Conn Score

	BgpBgpSiteScore *int `json:"bgpBgpSiteScore,omitempty"` // Bgp Bgp Site Score

	VniStatusScore *int `json:"vniStatusScore,omitempty"` // Vni Status Score

	PubsubTransitConnScore *int `json:"pubsubTransitConnScore,omitempty"` // Pubsub Transit Conn Score

	BgpPeerInfraVnScore *int `json:"bgpPeerInfraVnScore,omitempty"` // Bgp Peer Infra Vn Score

	InternetAvailScore *int `json:"internetAvailScore,omitempty"` // Internet Avail Score

	BgpEvpnScore *int `json:"bgpEvpnScore,omitempty"` // Bgp Evpn Score

	LispTransitConnScore *int `json:"lispTransitConnScore,omitempty"` // Lisp Transit Conn Score

	CtsEnvDataDownloadScore *int `json:"ctsEnvDataDownloadScore,omitempty"` // Cts Env Data Download Score

	PubsubInfraVnScore *int `json:"pubsubInfraVnScore,omitempty"` // Pubsub Infra Vn Score

	PeerScore *int `json:"peerScore,omitempty"` // Peer Score

	BgpPeerScore *int `json:"bgpPeerScore,omitempty"` // Bgp Peer Score

	RemoteInternetAvailScore *int `json:"remoteInternetAvailScore,omitempty"` // Remote Internet Avail Score

	BgpTCPScore *int `json:"bgpTcpScore,omitempty"` // Bgp Tcp Score

	PubsubSessionScore *int `json:"pubsubSessionScore,omitempty"` // Pubsub Session Score

	AAAStatusScore *int `json:"aaaStatusScore,omitempty"` // Aaa Status Score

	LispCpConnScore *int `json:"lispCpConnScore,omitempty"` // Lisp Cp Conn Score

	BgpPubsubSiteScore *int `json:"bgpPubsubSiteScore,omitempty"` // Bgp Pubsub Site Score

	McastScore *int `json:"mcastScore,omitempty"` // Mcast Score

	PortChannelScore *int `json:"portChannelScore,omitempty"` // Port Channel Score
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy string `json:"sortBy,omitempty"` // Sort By

	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	Response *ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices struct {
	Response *ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponse `json:"response,omitempty"` //

	Page *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage `json:"page,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponse struct {
	Attributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAttributes `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` // Aggregate Attributes

	Groups *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroups `json:"groups,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAttributes interface{}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseAggregateAttributes interface{}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices struct {
	Response *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesResponse `json:"response,omitempty"` //

	Page *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesResponse struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy *[]ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order

	Function string `json:"function,omitempty"` // Function
}
type ResponseDevicesGetsTheTrendAnalyticsData struct {
	Response *[]ResponseDevicesGetsTheTrendAnalyticsDataResponse `json:"response,omitempty"` //

	Page *ResponseDevicesGetsTheTrendAnalyticsDataPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponse struct {
	Timestamp *float64 `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Groups *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseGroups `json:"groups,omitempty"` //
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesGetsTheTrendAnalyticsDataPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID struct {
	Response *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Management Ip Address

	PlatformID string `json:"platformId,omitempty"` // Platform Id

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number

	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	DeviceSeries string `json:"deviceSeries,omitempty"` // Device Series

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software Version

	ProductVendor string `json:"productVendor,omitempty"` // Product Vendor

	DeviceRole string `json:"deviceRole,omitempty"` // Device Role

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	CommunicationState string `json:"communicationState,omitempty"` // Communication State

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection Status

	HaStatus string `json:"haStatus,omitempty"` // Ha Status

	LastBootTime *int `json:"lastBootTime,omitempty"` // Last Boot Time

	SiteHierarchyID string `json:"siteHierarchyId,omitempty"` // Site Hierarchy Id

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy

	SiteID string `json:"siteId,omitempty"` // Site Id

	DeviceGroupHierarchyID string `json:"deviceGroupHierarchyId,omitempty"` // Device Group Hierarchy Id

	TagNames []string `json:"tagNames,omitempty"` // Tag Names

	StackType string `json:"stackType,omitempty"` // Stack Type

	OsType string `json:"osType,omitempty"` // Os Type

	RingStatus *bool `json:"ringStatus,omitempty"` // Ring Status

	MaintenanceModeEnabled *bool `json:"maintenanceModeEnabled,omitempty"` // Maintenance Mode Enabled

	UpTime *int `json:"upTime,omitempty"` // Up Time

	IPv4Address string `json:"ipv4Address,omitempty"` // Ipv4 Address

	IPv6Address string `json:"ipv6Address,omitempty"` // Ipv6 Address

	RedundancyMode string `json:"redundancyMode,omitempty"` // Redundancy Mode

	FeatureFlagList []string `json:"featureFlagList,omitempty"` // Feature Flag List

	HaLastResetReason string `json:"haLastResetReason,omitempty"` // Ha Last Reset Reason

	RedundancyPeerStateDerived string `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived

	RedundancyPeerState string `json:"redundancyPeerState,omitempty"` // Redundancy Peer State

	RedundancyStateDerived string `json:"redundancyStateDerived,omitempty"` // Redundancy State Derived

	RedundancyState string `json:"redundancyState,omitempty"` // Redundancy State

	WiredClientCount *int `json:"wiredClientCount,omitempty"` // Wired Client Count

	WirelessClientCount *int `json:"wirelessClientCount,omitempty"` // Wireless Client Count

	PortCount *int `json:"portCount,omitempty"` // Port Count

	PhysicalPortCount *int `json:"physicalPortCount,omitempty"` // Physical Port Count

	VirtualPortCount *int `json:"virtualPortCount,omitempty"` // Virtual Port Count

	ClientCount *int `json:"clientCount,omitempty"` // Client Count

	ApDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetails `json:"apDetails,omitempty"` //

	MetricsDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseMetricsDetails `json:"metricsDetails,omitempty"` //

	FabricDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricDetails `json:"fabricDetails,omitempty"` //

	SwitchPoeDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseSwitchPoeDetails `json:"switchPoeDetails,omitempty"` //

	FabricMetricsDetails *ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricMetricsDetails `json:"fabricMetricsDetails,omitempty"` //

	AggregateAttributes *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetails struct {
	ConnectedWlcName string `json:"connectedWlcName,omitempty"` // Connected Wlc Name

	PolicyTagName string `json:"policyTagName,omitempty"` // Policy Tag Name

	ApOperationalState string `json:"apOperationalState,omitempty"` // Ap Operational State

	PowerSaveMode string `json:"powerSaveMode,omitempty"` // Power Save Mode

	OperationalMode string `json:"operationalMode,omitempty"` // Operational Mode

	ResetReason string `json:"resetReason,omitempty"` // Reset Reason

	Protocol string `json:"protocol,omitempty"` // Protocol

	PowerMode string `json:"powerMode,omitempty"` // Power Mode

	ConnectedTime *int `json:"connectedTime,omitempty"` // Connected Time

	LedFlashEnabled *bool `json:"ledFlashEnabled,omitempty"` // Led Flash Enabled

	LedFlashSeconds *int `json:"ledFlashSeconds,omitempty"` // Led Flash Seconds

	SubMode string `json:"subMode,omitempty"` // Sub Mode

	HomeApEnabled *bool `json:"homeApEnabled,omitempty"` // Home Ap Enabled

	PowerType string `json:"powerType,omitempty"` // Power Type

	ApType string `json:"apType,omitempty"` // Ap Type

	AdminState string `json:"adminState,omitempty"` // Admin State

	IcapCapability string `json:"icapCapability,omitempty"` // Icap Capability

	RegulatoryDomain string `json:"regulatoryDomain,omitempty"` // Regulatory Domain

	EthernetMac string `json:"ethernetMac,omitempty"` // Ethernet Mac

	RfTagName string `json:"rfTagName,omitempty"` // Rf Tag Name

	SiteTagName string `json:"siteTagName,omitempty"` // Site Tag Name

	PowerSaveModeCapable string `json:"powerSaveModeCapable,omitempty"` // Power Save Mode Capable

	PowerProfile string `json:"powerProfile,omitempty"` // Power Profile

	FlexGroup string `json:"flexGroup,omitempty"` // Flex Group

	PowerCalendarProfile string `json:"powerCalendarProfile,omitempty"` // Power Calendar Profile

	ApGroup string `json:"apGroup,omitempty"` // Ap Group

	Radios *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetailsRadios `json:"radios,omitempty"` //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseApDetailsRadios struct {
	ID string `json:"id,omitempty"` // Id

	Band string `json:"band,omitempty"` // Band

	Noise *int `json:"noise,omitempty"` // Noise

	AirQuality *float64 `json:"airQuality,omitempty"` // Air Quality

	Interference *float64 `json:"interference,omitempty"` // Interference

	TrafficUtil *int `json:"trafficUtil,omitempty"` // Traffic Util

	Utilization *float64 `json:"utilization,omitempty"` // Utilization

	ClientCount *int `json:"clientCount,omitempty"` // Client Count
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseMetricsDetails struct {
	OverallHealthScore *int `json:"overallHealthScore,omitempty"` // Overall Health Score

	CPUUtilization *float64 `json:"cpuUtilization,omitempty"` // Cpu Utilization

	CPUScore *int `json:"cpuScore,omitempty"` // Cpu Score

	MemoryUtilization *float64 `json:"memoryUtilization,omitempty"` // Memory Utilization

	MemoryScore *int `json:"memoryScore,omitempty"` // Memory Score

	AvgTemperature *float64 `json:"avgTemperature,omitempty"` // Avg Temperature

	MaxTemperature *float64 `json:"maxTemperature,omitempty"` // Max Temperature

	DiscardScore *int `json:"discardScore,omitempty"` // Discard Score

	DiscardInterfaces []string `json:"discardInterfaces,omitempty"` // Discard Interfaces

	ErrorScore *int `json:"errorScore,omitempty"` // Error Score

	ErrorInterfaces []string `json:"errorInterfaces,omitempty"` // Error Interfaces

	InterDeviceLinkScore *int `json:"interDeviceLinkScore,omitempty"` // Inter Device Link Score

	InterDeviceConnectedDownInterfaces []string `json:"interDeviceConnectedDownInterfaces,omitempty"` // Inter Device Connected Down Interfaces

	LinkUtilizationScore *int `json:"linkUtilizationScore,omitempty"` // Link Utilization Score

	HighLinkUtilizationInterfaces []string `json:"highLinkUtilizationInterfaces,omitempty"` // High Link Utilization Interfaces

	FreeTimerScore *int `json:"freeTimerScore,omitempty"` // Free Timer Score

	FreeTimer *float64 `json:"freeTimer,omitempty"` // Free Timer

	PacketPoolScore *int `json:"packetPoolScore,omitempty"` // Packet Pool Score

	PacketPool *int `json:"packetPool,omitempty"` // Packet Pool

	FreeMemoryBufferScore *int `json:"freeMemoryBufferScore,omitempty"` // Free Memory Buffer Score

	FreeMemoryBuffer *float64 `json:"freeMemoryBuffer,omitempty"` // Free Memory Buffer

	WqePoolScore *int `json:"wqePoolScore,omitempty"` // Wqe Pool Score

	WqePool *int `json:"wqePool,omitempty"` // Wqe Pool

	ApCount *int `json:"apCount,omitempty"` // Ap Count

	NoiseScore *int `json:"noiseScore,omitempty"` // Noise Score

	UtilizationScore *int `json:"utilizationScore,omitempty"` // Utilization Score

	InterferenceScore *int `json:"interferenceScore,omitempty"` // Interference Score

	AirQualityScore *int `json:"airQualityScore,omitempty"` // Air Quality Score
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricDetails struct {
	FabricRole []string `json:"fabricRole,omitempty"` // Fabric Role

	FabricSiteName string `json:"fabricSiteName,omitempty"` // Fabric Site Name

	TransitFabrics []string `json:"transitFabrics,omitempty"` // Transit Fabrics

	L2Vns []string `json:"l2Vns,omitempty"` // L2 Vns

	L3Vns []string `json:"l3Vns,omitempty"` // L3 Vns

	FabricSiteID string `json:"fabricSiteId,omitempty"` // Fabric Site Id

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseSwitchPoeDetails struct {
	PortCount *int `json:"portCount,omitempty"` // Port Count

	UsedPortCount *int `json:"usedPortCount,omitempty"` // Used Port Count

	FreePortCount *int `json:"freePortCount,omitempty"` // Free Port Count

	PowerConsumed *float64 `json:"powerConsumed,omitempty"` // Power Consumed

	PoePowerConsumed *int `json:"poePowerConsumed,omitempty"` // Poe Power Consumed

	SystemPowerConsumed *float64 `json:"systemPowerConsumed,omitempty"` // System Power Consumed

	PowerBudget *int `json:"powerBudget,omitempty"` // Power Budget

	PoePowerAllocated *float64 `json:"poePowerAllocated,omitempty"` // Poe Power Allocated

	SystemPowerAllocated *int `json:"systemPowerAllocated,omitempty"` // System Power Allocated

	PowerRemaining *float64 `json:"powerRemaining,omitempty"` // Power Remaining

	PoeVersion string `json:"poeVersion,omitempty"` // Poe Version

	ChassisCount *int `json:"chassisCount,omitempty"` // Chassis Count

	ModuleCount *int `json:"moduleCount,omitempty"` // Module Count

	ModuleDetails *[]ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseSwitchPoeDetailsModuleDetails `json:"moduleDetails,omitempty"` //
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseSwitchPoeDetailsModuleDetails struct {
	ModuleID string `json:"moduleId,omitempty"` // Module Id

	ChassisID string `json:"chassisId,omitempty"` // Chassis Id

	ModulePortCount *int `json:"modulePortCount,omitempty"` // Module Port Count

	ModuleUsedPortCount *int `json:"moduleUsedPortCount,omitempty"` // Module Used Port Count

	ModuleFreePortCount *int `json:"moduleFreePortCount,omitempty"` // Module Free Port Count

	ModulePowerConsumed *float64 `json:"modulePowerConsumed,omitempty"` // Module Power Consumed

	ModulePoePowerConsumed *int `json:"modulePoePowerConsumed,omitempty"` // Module Poe Power Consumed

	ModuleSystemPowerConsumed *float64 `json:"moduleSystemPowerConsumed,omitempty"` // Module System Power Consumed

	ModulePowerBudget *int `json:"modulePowerBudget,omitempty"` // Module Power Budget

	ModulePoePowerAllocated *float64 `json:"modulePoePowerAllocated,omitempty"` // Module Poe Power Allocated

	ModuleSystemPowerAllocated *int `json:"moduleSystemPowerAllocated,omitempty"` // Module System Power Allocated

	ModulePowerRemaining *float64 `json:"modulePowerRemaining,omitempty"` // Module Power Remaining

	InterfacePowerMax *int `json:"interfacePowerMax,omitempty"` // Interface Power Max
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseFabricMetricsDetails struct {
	OverallFabricScore *int `json:"overallFabricScore,omitempty"` // Overall Fabric Score

	FabricTransitScore *int `json:"fabricTransitScore,omitempty"` // Fabric Transit Score

	FabricSiteScore *int `json:"fabricSiteScore,omitempty"` // Fabric Site Score

	FabricVnScore *int `json:"fabricVnScore,omitempty"` // Fabric Vn Score

	FabsiteFcpScore *int `json:"fabsiteFcpScore,omitempty"` // Fabsite Fcp Score

	FabsiteInfraScore *int `json:"fabsiteInfraScore,omitempty"` // Fabsite Infra Score

	FabsiteFsconnScore *int `json:"fabsiteFsconnScore,omitempty"` // Fabsite Fsconn Score

	VnExitScore *int `json:"vnExitScore,omitempty"` // Vn Exit Score

	VnFcpScore *int `json:"vnFcpScore,omitempty"` // Vn Fcp Score

	VnStatusScore *int `json:"vnStatusScore,omitempty"` // Vn Status Score

	VnServiceScore *int `json:"vnServiceScore,omitempty"` // Vn Service Score

	TransitControlPlaneScore *int `json:"transitControlPlaneScore,omitempty"` // Transit Control Plane Score

	TransitServicesScore *int `json:"transitServicesScore,omitempty"` // Transit Services Score

	TCPConnScore *int `json:"tcpConnScore,omitempty"` // Tcp Conn Score

	BgpBgpSiteScore *int `json:"bgpBgpSiteScore,omitempty"` // Bgp Bgp Site Score

	VniStatusScore *int `json:"vniStatusScore,omitempty"` // Vni Status Score

	PubsubTransitConnScore *int `json:"pubsubTransitConnScore,omitempty"` // Pubsub Transit Conn Score

	BgpPeerInfraVnScore *int `json:"bgpPeerInfraVnScore,omitempty"` // Bgp Peer Infra Vn Score

	InternetAvailScore *int `json:"internetAvailScore,omitempty"` // Internet Avail Score

	BgpEvpnScore *int `json:"bgpEvpnScore,omitempty"` // Bgp Evpn Score

	LispTransitConnScore *int `json:"lispTransitConnScore,omitempty"` // Lisp Transit Conn Score

	CtsEnvDataDownloadScore *int `json:"ctsEnvDataDownloadScore,omitempty"` // Cts Env Data Download Score

	PubsubInfraVnScore *int `json:"pubsubInfraVnScore,omitempty"` // Pubsub Infra Vn Score

	PeerScore *int `json:"peerScore,omitempty"` // Peer Score

	BgpPeerScore *int `json:"bgpPeerScore,omitempty"` // Bgp Peer Score

	RemoteInternetAvailScore *int `json:"remoteInternetAvailScore,omitempty"` // Remote Internet Avail Score

	BgpTCPScore *int `json:"bgpTcpScore,omitempty"` // Bgp Tcp Score

	PubsubSessionScore *int `json:"pubsubSessionScore,omitempty"` // Pubsub Session Score

	AAAStatusScore *int `json:"aaaStatusScore,omitempty"` // Aaa Status Score

	LispCpConnScore *int `json:"lispCpConnScore,omitempty"` // Lisp Cp Conn Score

	BgpPubsubSiteScore *int `json:"bgpPubsubSiteScore,omitempty"` // Bgp Pubsub Site Score

	McastScore *int `json:"mcastScore,omitempty"` // Mcast Score

	PortChannelScore *int `json:"portChannelScore,omitempty"` // Port Channel Score
}
type ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUIDResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange struct {
	Response *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponse `json:"response,omitempty"` //

	Page *ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponse struct {
	Timestamp *float64 `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Groups *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroups `json:"groups,omitempty"` //
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroups struct {
	ID string `json:"id,omitempty"` // Id

	Attributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAttributes `json:"attributes,omitempty"` //

	AggregateAttributes *[]ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value string `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeResponseGroupsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Value *float64 `json:"value,omitempty"` // Value
}
type ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type ResponseDevicesGetPlannedAccessPointsForBuilding struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForBuildingResponse `json:"response,omitempty"` //

	Version *int `json:"version,omitempty"` // Version of the api response model

	Total *int `json:"total,omitempty"` // Total number of the planned access points
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponse struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributes `json:"attributes,omitempty"` //

	Location *ResponseDevicesGetPlannedAccessPointsForBuildingResponseLocation `json:"location,omitempty"` //

	Position *ResponseDevicesGetPlannedAccessPointsForBuildingResponsePosition `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadios `json:"radios,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if the planned access point is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseAttributes struct {
	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseLocation struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponsePosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributes `json:"attributes,omitempty"` //

	Antenna *ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAntenna `json:"antenna,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAttributes struct {
	ID *int `json:"id,omitempty"` // Id of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band type of the radio

	Channel *float64 `json:"channel,omitempty"` // Channel in which the radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type ResponseDevicesGetPlannedAccessPointsForBuildingResponseRadiosAntenna struct {
	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna
}
type ResponseDevicesGetDeviceDetail struct {
	Response *ResponseDevicesGetDeviceDetailResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetDeviceDetailResponse struct {
	NoiseScore *int `json:"noiseScore,omitempty"` // Device (AP) WIFI signal noise health score

	PolicyTagName string `json:"policyTagName,omitempty"` // Device (AP) policy tag

	InterferenceScore *int `json:"interferenceScore,omitempty"` // Device (AP) WIFI signal interference health score

	OpState string `json:"opState,omitempty"` // Operation state of device (AP)

	PowerSaveMode string `json:"powerSaveMode,omitempty"` // Device power save mode

	Mode string `json:"mode,omitempty"` // Device mode (AP)

	ResetReason string `json:"resetReason,omitempty"` // Device reset reason

	NwDeviceRole string `json:"nwDeviceRole,omitempty"` // Device role

	Protocol string `json:"protocol,omitempty"` // Protocol code

	PowerMode string `json:"powerMode,omitempty"` // Device's power mode

	ConnectedTime string `json:"connectedTime,omitempty"` // UTC timestamp

	RingStatus *bool `json:"ringStatus,omitempty"` // Device's ring status

	LedFlashSeconds string `json:"ledFlashSeconds,omitempty"` // LED flash seconds

	IPAddrManagementIPAddr string `json:"ip_addr_managementIpAddr,omitempty"` // Device's management IP address

	StackType string `json:"stackType,omitempty"` // Device stack type (applicable for stackable devices)

	SubMode string `json:"subMode,omitempty"` // Device submode

	SerialNumber string `json:"serialNumber,omitempty"` // Device serial number

	NwDeviceName string `json:"nwDeviceName,omitempty"` // Device name

	DeviceGroupHierarchyID string `json:"deviceGroupHierarchyId,omitempty"` // Device group site hierarchy UUID

	CPU *float64 `json:"cpu,omitempty"` // Device CPU utilization

	Utilization string `json:"utilization,omitempty"` // Device utilization

	NwDeviceID string `json:"nwDeviceId,omitempty"` // Device's UUID

	SiteHierarchyGraphID string `json:"siteHierarchyGraphId,omitempty"` // Site hierarchy UUID in which device is assigned to

	NwDeviceFamily string `json:"nwDeviceFamily,omitempty"` // Device faimly string

	MacAddress string `json:"macAddress,omitempty"` // Device MAC address

	HomeApEnabled string `json:"homeApEnabled,omitempty"` // Home Ap Enabled

	DeviceSeries string `json:"deviceSeries,omitempty"` // Device series string

	CollectionStatus string `json:"collectionStatus,omitempty"` // Device's telemetry data collection status for DNAC

	UtilizationScore *int `json:"utilizationScore,omitempty"` // Device utilization health score

	MaintenanceMode *bool `json:"maintenanceMode,omitempty"` // Whether device is in maintenance mode

	Interference string `json:"interference,omitempty"` // Device (AP) WIFI signal interference

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device's software version string

	TagIDList *[]ResponseDevicesGetDeviceDetailResponseTagIDList `json:"tagIdList,omitempty"` // Tag ID List

	PowerType string `json:"powerType,omitempty"` // Device (AP) power type

	OverallHealth *int `json:"overallHealth,omitempty"` // Device's overall health score

	ManagementIPAddr string `json:"managementIpAddr,omitempty"` // Management IP address of the device

	Memory string `json:"memory,omitempty"` // Device memory utilization

	CommunicationState string `json:"communicationState,omitempty"` // Device communication state

	ApType string `json:"apType,omitempty"` // Ap Type

	AdminState string `json:"adminState,omitempty"` // Device (AP) admin state

	Noise string `json:"noise,omitempty"` // Device (AP) WIFI signal noise

	IcapCapability string `json:"icapCapability,omitempty"` // Device (AP) ICAP capability bit values

	RegulatoryDomain string `json:"regulatoryDomain,omitempty"` // Device (AP) WIFI domain

	EthernetMac string `json:"ethernetMac,omitempty"` // Device (AP) ethernet MAC address

	NwDeviceType string `json:"nwDeviceType,omitempty"` // Device type

	AirQuality string `json:"airQuality,omitempty"` // Device (AP) WIFI air quality

	RfTagName string `json:"rfTagName,omitempty"` // Device (AP) RF tag name

	SiteTagName string `json:"siteTagName,omitempty"` // Device (AP) site tag name

	PlatformID string `json:"platformId,omitempty"` // Device's platform ID

	UpTime string `json:"upTime,omitempty"` // Device up time

	MemoryScore *int `json:"memoryScore,omitempty"` // Device's memory usage score

	PowerSaveModeCapable string `json:"powerSaveModeCapable,omitempty"` // Device (AP) power save mode capability

	PowerProfile string `json:"powerProfile,omitempty"` // Device (AP) power profile name

	AirQualityScore *int `json:"airQualityScore,omitempty"` // Device (AP) air quality health score

	Location string `json:"location,omitempty"` // Device's site hierarchy UUID

	FlexGroup string `json:"flexGroup,omitempty"` // Deivce (A) flexconnect group

	LastBootTime *float64 `json:"lastBootTime,omitempty"` // Device's last boot UTC timestamp

	PowerCalendarProfile string `json:"powerCalendarProfile,omitempty"` // Device (AP) power calendar profile name

	ConnectivityStatus *int `json:"connectivityStatus,omitempty"` // Device connectivity status

	LedFlashEnabled string `json:"ledFlashEnabled,omitempty"` // Device (AP) LED flash

	CPUScore *int `json:"cpuScore,omitempty"` // Device's CPU usage score

	AvgTemperature *float64 `json:"avgTemperature,omitempty"` // Device's average temperature

	MaxTemperature *float64 `json:"maxTemperature,omitempty"` // Device's max temperature

	HaStatus string `json:"haStatus,omitempty"` // Device's HA status

	OsType string `json:"osType,omitempty"` // Device's OS type

	Timestamp *int `json:"timestamp,omitempty"` // UTC timestamp of the device health data

	ApGroup string `json:"apGroup,omitempty"` // Device (AP) AP group

	RedundancyMode string `json:"redundancyMode,omitempty"` // Device redundancy mode

	FeatureFlagList []string `json:"featureFlagList,omitempty"` // List of device feature capabilities

	FreeMbufScore *int `json:"freeMbufScore,omitempty"` // Free memory buffer health score

	HALastResetReason string `json:"HALastResetReason,omitempty"` // Last HA reset reason

	WqeScore *int `json:"wqeScore,omitempty"` // WQE health score

	RedundancyPeerStateDerived string `json:"redundancyPeerStateDerived,omitempty"` // Redundancy Peer State Derived

	FreeTimerScore *int `json:"freeTimerScore,omitempty"` // Free Timer Score

	RedundancyPeerState string `json:"redundancyPeerState,omitempty"` // Redundancy Peer State

	RedundancyStateDerived string `json:"redundancyStateDerived,omitempty"` // Derived redundancy state

	RedundancyState string `json:"redundancyState,omitempty"` // Redundancy state

	PacketPoolScore *int `json:"packetPoolScore,omitempty"` // Device packet pool health score

	FreeTimer *float64 `json:"freeTimer,omitempty"` // Free timer of the device

	PacketPool *float64 `json:"packetPool,omitempty"` // Packet pool of the device

	Wqe *float64 `json:"wqe,omitempty"` // WQE of the device

	FreeMbuf *float64 `json:"freeMbuf,omitempty"` // Free memory buffer of the device
}
type ResponseDevicesGetDeviceDetailResponseTagIDList interface{}
type ResponseDevicesGetDeviceEnrichmentDetails []ResponseItemDevicesGetDeviceEnrichmentDetails // Array of ResponseDevicesGetDeviceEnrichmentDetails
type ResponseItemDevicesGetDeviceEnrichmentDetails struct {
	DeviceDetails *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetails struct {
	Family string `json:"family,omitempty"` // Device Family

	Type string `json:"type,omitempty"` // Device Type

	Location *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation `json:"location,omitempty"` // Device location - Site hierarchy

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	MacAddress string `json:"macAddress,omitempty"` // Device MAC address

	Role string `json:"role,omitempty"` // Device role

	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated WLC IP address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device's last boot UTC timestamp

	CollectionStatus string `json:"collectionStatus,omitempty"` // Device's telemetry data collection status for DNAC

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // Device Management Ip Address

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Device's platform ID

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability Status of the Device(Reachable/Unreachable)

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	TunnelUDPPort *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	WaasDeviceMode *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode `json:"waasDeviceMode,omitempty"` // WAAS device mode

	Series string `json:"series,omitempty"` // Device Series

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	SerialNumber string `json:"serialNumber,omitempty"` // Device Serial Number

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Device Software Version

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	Hostname string `json:"hostname,omitempty"` // Device Hostname

	UpTime string `json:"upTime,omitempty"` // Device's uptime

	LastUpdateTime *int `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	LocationName *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	ID string `json:"id,omitempty"` // Device's UUID

	NeighborTopology *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocation interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsTunnelUDPPort interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsWaasDeviceMode interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsLocationName interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //

	Links *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodes struct {
	Role string `json:"role,omitempty"` // Role of the Node

	Name string `json:"name,omitempty"` // Hostname of the Node

	ID string `json:"id,omitempty"` // Id of the Node

	Description string `json:"description,omitempty"` // Description of the Node

	DeviceType string `json:"deviceType,omitempty"` // Device type of the node, like switch, AP, WCL,GateWay

	PlatformID string `json:"platformId,omitempty"` // Type of platform

	Family string `json:"family,omitempty"` // Device Family of the Node

	IP string `json:"ip,omitempty"` // IP Address of the Node

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software Version of the Node

	UserID *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID `json:"userId,omitempty"` // User Id of the Node

	NodeType string `json:"nodeType,omitempty"` // Type of the Node

	RadioFrequency *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency `json:"radioFrequency,omitempty"` // Frequency of wireless radio channel

	Clients *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients `json:"clients,omitempty"` // Number of clients

	Count *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount `json:"count,omitempty"` // The number of group nodes (for ap sepecifically)

	HealthScore *int `json:"healthScore,omitempty"` // The total health score of the node

	Level *float64 `json:"level,omitempty"` // The level index to be used by UI widget (starts from 0)

	FabricGroup *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup `json:"fabricGroup,omitempty"` // Fabric device group name

	ConnectedDevice *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice `json:"connectedDevice,omitempty"` // The connected device to show the connected switch to wlc
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesClients interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyNodesConnectedDevice interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinks struct {
	Source string `json:"source,omitempty"` // Edge line starting node

	LinkStatus string `json:"linkStatus,omitempty"` // The status of the link (up/down)

	Label *[]ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel `json:"label,omitempty"` // The details of the edge

	Target string `json:"target,omitempty"` // End node of the edge line

	ID *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID `json:"id,omitempty"` // Id of the node

	PortUtilization *ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Number of clients
}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemDevicesGetDeviceEnrichmentDetailsDeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseDevicesDevices struct {
	Version string `json:"version,omitempty"` // Response data's version string

	TotalCount *int `json:"totalCount,omitempty"` // Total number of devices

	Response *[]ResponseDevicesDevicesResponse `json:"response,omitempty"` //
}
type ResponseDevicesDevicesResponse struct {
	DeviceType string `json:"deviceType,omitempty"` // Device type

	CPUUtilization *float64 `json:"cpuUtilization,omitempty"` // Device's CPU utilization

	OverallHealth *int `json:"overallHealth,omitempty"` // Overall health score

	UtilizationHealth *ResponseDevicesDevicesResponseUtilizationHealth `json:"utilizationHealth,omitempty"` //

	AirQualityHealth *ResponseDevicesDevicesResponseAirQualityHealth `json:"airQualityHealth,omitempty"` //

	IPAddress string `json:"ipAddress,omitempty"` // Management IP address of the device

	CPUHealth *int `json:"cpuHealth,omitempty"` // Device CPU health score

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device family

	IssueCount *int `json:"issueCount,omitempty"` // Number of issues

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the device

	NoiseHealth *ResponseDevicesDevicesResponseNoiseHealth `json:"noiseHealth,omitempty"` //

	OsVersion string `json:"osVersion,omitempty"` // Device OS version string

	Name string `json:"name,omitempty"` // Device name

	InterfaceLinkErrHealth *int `json:"interfaceLinkErrHealth,omitempty"` // Device (AP) error health score

	MemoryUtilization *float64 `json:"memoryUtilization,omitempty"` // Device memory utilization

	InterDeviceLinkAvailHealth *int `json:"interDeviceLinkAvailHealth,omitempty"` // Device connectivity status

	InterferenceHealth *ResponseDevicesDevicesResponseInterferenceHealth `json:"interferenceHealth,omitempty"` //

	Model string `json:"model,omitempty"` // Device model string

	Location string `json:"location,omitempty"` // Site location in which this device is assigned to

	ReachabilityHealth string `json:"reachabilityHealth,omitempty"` // Device reachability in the network

	Band *ResponseDevicesDevicesResponseBand `json:"band,omitempty"` //

	MemoryUtilizationHealth *int `json:"memoryUtilizationHealth,omitempty"` // Device memory utilization health score

	ClientCount *ResponseDevicesDevicesResponseClientCount `json:"clientCount,omitempty"` //

	AvgTemperature *float64 `json:"avgTemperature,omitempty"` // Average device (switch) temperature

	MaxTemperature *float64 `json:"maxTemperature,omitempty"` // Max device (switch) temperature

	InterDeviceLinkAvailFabric *int `json:"interDeviceLinkAvailFabric,omitempty"` // Device uplink health

	ApCount *int `json:"apCount,omitempty"` // Number of AP count

	FreeTimerScore *int `json:"freeTimerScore,omitempty"` // Device free timer health score

	FreeTimer *float64 `json:"freeTimer,omitempty"` // Device free timer

	PacketPoolHealth *int `json:"packetPoolHealth,omitempty"` // Device packet pool

	PacketPool *int `json:"packetPool,omitempty"` // Device packet pool

	FreeMemoryBufferHealth *int `json:"freeMemoryBufferHealth,omitempty"` // Device free memory buffer health

	FreeMemoryBuffer *float64 `json:"freeMemoryBuffer,omitempty"` // Device free memory

	WqePoolsHealth *int `json:"wqePoolsHealth,omitempty"` // Device WQE pool health

	WqePools *float64 `json:"wqePools,omitempty"` // Device WQE pool

	WanLinkUtilization *float64 `json:"wanLinkUtilization,omitempty"` // WLAN link utilization

	CPUUlitilization *float64 `json:"cpuUlitilization,omitempty"` // Device's CPU utilization

	UUID string `json:"uuid,omitempty"` // Device UUID
}
type ResponseDevicesDevicesResponseUtilizationHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesResponseAirQualityHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesResponseNoiseHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesResponseInterferenceHealth struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesDevicesResponseBand struct {
	Radio0 string `json:"radio0,omitempty"` // Radio0

	Radio1 string `json:"radio1,omitempty"` // Radio1

	Radio2 string `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3
}
type ResponseDevicesDevicesResponseClientCount struct {
	Radio0 *int `json:"radio0,omitempty"` // Radio0

	Radio1 *int `json:"radio1,omitempty"` // Radio1

	Radio2 *int `json:"radio2,omitempty"` // Radio2

	Radio3 *int `json:"radio3,omitempty"` // Radio3

	Ghz24 *int `json:"Ghz24,omitempty"` // Ghz24

	Ghz50 *int `json:"Ghz50,omitempty"` // Ghz50
}
type ResponseDevicesUpdatePlannedAccessPointForFloor struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseDevicesUpdatePlannedAccessPointForFloorResponse `json:"response,omitempty"` //
}
type ResponseDevicesUpdatePlannedAccessPointForFloorResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesCreatePlannedAccessPointForFloor struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseDevicesCreatePlannedAccessPointForFloorResponse `json:"response,omitempty"` //
}
type ResponseDevicesCreatePlannedAccessPointForFloorResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesGetPlannedAccessPointsForFloor struct {
	Response *[]ResponseDevicesGetPlannedAccessPointsForFloorResponse `json:"response,omitempty"` //

	Version *int `json:"version,omitempty"` // Version of the api response model

	Total *int `json:"total,omitempty"` // Total number of the planned access points
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponse struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes `json:"attributes,omitempty"` //

	Location *ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation `json:"location,omitempty"` //

	Position *ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios `json:"radios,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if the planned access point is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseAttributes struct {
	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseLocation struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponsePosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadios struct {
	Attributes *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes `json:"attributes,omitempty"` //

	Antenna *ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna `json:"antenna,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAttributes struct {
	ID *int `json:"id,omitempty"` // Id of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band type of the radio

	Channel *float64 `json:"channel,omitempty"` // Channel in which the radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type ResponseDevicesGetPlannedAccessPointsForFloorResponseRadiosAntenna struct {
	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna
}
type ResponseDevicesDeletePlannedAccessPointForFloor struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseDevicesDeletePlannedAccessPointForFloorResponse `json:"response,omitempty"` //
}
type ResponseDevicesDeletePlannedAccessPointForFloorResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters struct {
	Response *[]ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetAllHealthScoreDefinitionsForGivenFiltersResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesUpdateHealthScoreDefinitions struct {
	Response *[]ResponseDevicesUpdateHealthScoreDefinitionsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateHealthScoreDefinitionsResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters struct {
	Response *ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseDevicesGetHealthScoreDefinitionForTheGivenID struct {
	Response *[]ResponseDevicesGetHealthScoreDefinitionForTheGivenIDResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetHealthScoreDefinitionForTheGivenIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID struct {
	Response *ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateHealthScoreDefinitionForTheGivenIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	DisplayName string `json:"displayName,omitempty"` // Display Name

	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family

	Description string `json:"description,omitempty"` // Description

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	DefinitionStatus string `json:"definitionStatus,omitempty"` // Definition Status

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold

	LastModified string `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesGetAllInterfaces struct {
	Response *[]ResponseDevicesGetAllInterfacesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetAllInterfacesResponse struct {
	Addresses *[]ResponseDevicesGetAllInterfacesResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetAllInterfacesResponseAddresses struct {
	Address *ResponseDevicesGetAllInterfacesResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetAllInterfacesResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetAllInterfacesResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetAllInterfacesResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetAllInterfacesResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetAllInterfacesResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfaceCountForMultipleDevices struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceByIP struct {
	Response *[]ResponseDevicesGetInterfaceByIPResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceByIPResponse struct {
	Addresses *[]ResponseDevicesGetInterfaceByIPResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetInterfaceByIPResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceByIPResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceByIPResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceByIPResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetIsisInterfaces struct {
	Response *[]ResponseDevicesGetIsisInterfacesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetIsisInterfacesResponse struct {
	Addresses *[]ResponseDevicesGetIsisInterfacesResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetIsisInterfacesResponseAddresses struct {
	Address *ResponseDevicesGetIsisInterfacesResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetIsisInterfacesResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetIsisInterfacesResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetInterfaceInfoByID struct {
	Response *[]ResponseDevicesGetInterfaceInfoByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceInfoByIDResponse struct {
	Addresses *[]ResponseDevicesGetInterfaceInfoByIDResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceInfoByIDResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfaceCount struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName struct {
	Response *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponse struct {
	Addresses *[]ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface

	Poweroverethernet string `json:"poweroverethernet,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	NetworkdeviceID string `json:"networkdevice_id,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedComputeElement string `json:"managedComputeElement,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedNetworkElement string `json:"managedNetworkElement,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedNetworkElementURL string `json:"managedNetworkElementUrl,omitempty"` // This is internal attribute.  Not to be used.  Deprecated

	ManagedComputeElementURL string `json:"managedComputeElementUrl,omitempty"` // This is internal attribute.  Not to be used.  Deprecated
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceNameResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRange struct {
	Response *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponse struct {
	Addresses *[]ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddresses struct {
	Address *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetDeviceInterfacesBySpecifiedRangeResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetOspfInterfaces struct {
	Response *[]ResponseDevicesGetOspfInterfacesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetOspfInterfacesResponse struct {
	Addresses *[]ResponseDevicesGetOspfInterfacesResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetOspfInterfacesResponseAddresses struct {
	Address *ResponseDevicesGetOspfInterfacesResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetOspfInterfacesResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetOspfInterfacesResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesGetInterfaceByID struct {
	Response *ResponseDevicesGetInterfaceByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetInterfaceByIDResponse struct {
	Addresses *[]ResponseDevicesGetInterfaceByIDResponseAddresses `json:"addresses,omitempty"` //

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	ClassName string `json:"className,omitempty"` // Classifies the port as switch port ,loopback interface etc.

	Description string `json:"description,omitempty"` // Description for the Interface

	Name string `json:"name,omitempty"` // Name for the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	Duplex string `json:"duplex,omitempty"` // Interface duplex as AutoNegotiate or FullDuplex

	ID string `json:"id,omitempty"` // ID of the Interface

	IfIndex string `json:"ifIndex,omitempty"` // Interface index

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the Interface

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the Interface

	InterfaceType string `json:"interfaceType,omitempty"` // Interface type as Physical or Virtual

	IPv4Address string `json:"ipv4Address,omitempty"` // IPV4 Address of the device

	IPv4Mask string `json:"ipv4Mask,omitempty"` // IPV4 Mask of the device

	IsisSupport string `json:"isisSupport,omitempty"` // Flag for ISIS enabled / disabled

	LastOutgoingPacketTime *float64 `json:"lastOutgoingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was sent from this interface

	LastIncomingPacketTime *float64 `json:"lastIncomingPacketTime,omitempty"` // Time, in milliseconds since UNIX epoch, when the last packet was received on this interface

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the device interface info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of interface

	MappedPhysicalInterfaceID string `json:"mappedPhysicalInterfaceId,omitempty"` // ID of physical interface mapped with the virtual interface of WLC

	MappedPhysicalInterfaceName string `json:"mappedPhysicalInterfaceName,omitempty"` // Physical interface name mapped with the virtual interface of WLC

	MediaType string `json:"mediaType,omitempty"` // Media Type of the interface

	Mtu string `json:"mtu,omitempty"` // MTU Information of Interface

	NativeVLANID string `json:"nativeVlanId,omitempty"` // Vlan to receive untagged frames on trunk port

	OspfSupport string `json:"ospfSupport,omitempty"` // Flag for OSPF enabled / disabled

	Pid string `json:"pid,omitempty"` // Platform ID of the device

	PortMode string `json:"portMode,omitempty"` // Port mode as access, trunk, routed

	PortName string `json:"portName,omitempty"` // Interface name

	PortType string `json:"portType,omitempty"` // Port type as Ethernet Port / Ethernet SVI / Ethernet Sub Interface

	SerialNo string `json:"serialNo,omitempty"` // Serial number of the device

	Series string `json:"series,omitempty"` // Series of the device

	Speed string `json:"speed,omitempty"` // Speed of the interface

	Status string `json:"status,omitempty"` // Interface status as Down / Up

	VLANID string `json:"vlanId,omitempty"` // Vlan ID of interface

	VoiceVLAN string `json:"voiceVlan,omitempty"` // Vlan information of the interface
}
type ResponseDevicesGetInterfaceByIDResponseAddresses struct {
	Address *ResponseDevicesGetInterfaceByIDResponseAddressesAddress `json:"address,omitempty"` //

	Type string `json:"type,omitempty"` // Type of the interface. For e.g. IPv4, IPv6 (with unicast, multicast, anycast, etc.)
}
type ResponseDevicesGetInterfaceByIDResponseAddressesAddress struct {
	IPAddress *ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPAddress `json:"ipAddress,omitempty"` //

	IPMask *ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPMask `json:"ipMask,omitempty"` //

	IsInverseMask *bool `json:"isInverseMask,omitempty"` // Inverse Mask of the IP address is enabled or not
}
type ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPAddress struct {
	Address string `json:"address,omitempty"` // IP address of the interface
}
type ResponseDevicesGetInterfaceByIDResponseAddressesAddressIPMask struct {
	Address string `json:"address,omitempty"` // IP Mask of the interface
}
type ResponseDevicesUpdateInterfaceDetails struct {
	Response *ResponseDevicesUpdateInterfaceDetailsResponse `json:"response,omitempty"` //

	Version *ResponseDevicesUpdateInterfaceDetailsVersion `json:"version,omitempty"` //
}
type ResponseDevicesUpdateInterfaceDetailsResponse struct {
	Type string `json:"type,omitempty"` // Type

	Properties *ResponseDevicesUpdateInterfaceDetailsResponseProperties `json:"properties,omitempty"` //

	Required []string `json:"required,omitempty"` // Required
}
type ResponseDevicesUpdateInterfaceDetailsResponseProperties struct {
	TaskID *ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID `json:"taskId,omitempty"` //

	URL *ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL `json:"url,omitempty"` //
}
type ResponseDevicesUpdateInterfaceDetailsResponsePropertiesTaskID struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesUpdateInterfaceDetailsResponsePropertiesURL struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesUpdateInterfaceDetailsVersion struct {
	Type string `json:"type,omitempty"` // Type
}
type ResponseDevicesLegitOperationsForInterface struct {
	Response *ResponseDevicesLegitOperationsForInterfaceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesLegitOperationsForInterfaceResponse struct {
	InterfaceUUID string `json:"interfaceUuid,omitempty"` // Id of the Interface

	Properties *[]ResponseDevicesLegitOperationsForInterfaceResponseProperties `json:"properties,omitempty"` //

	Operations *[]ResponseDevicesLegitOperationsForInterfaceResponseOperations `json:"operations,omitempty"` //
}
type ResponseDevicesLegitOperationsForInterfaceResponseProperties struct {
	Name string `json:"name,omitempty"` // Name of the Property

	Applicable string `json:"applicable,omitempty"` // Checks if property is applicable to interface

	FailureReason string `json:"failureReason,omitempty"` // Failure reason of the Property
}
type ResponseDevicesLegitOperationsForInterfaceResponseOperations struct {
	Name string `json:"name,omitempty"` // Name of the Operation

	Applicable string `json:"applicable,omitempty"` // Checks if operation is applicable to interface

	FailureReason string `json:"failureReason,omitempty"` // Failure reason of the Operation
}
type ResponseDevicesClearMacAddressTable struct {
	Response *ResponseDevicesClearMacAddressTableResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesClearMacAddressTableResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetDeviceList struct {
	Response *[]ResponseDevicesGetDeviceListResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetDeviceListResponse struct {
	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	Hostname string `json:"hostname,omitempty"` // Device name

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	Description string `json:"description,omitempty"` // System description

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	SyncRequestedByApp string `json:"syncRequestedByApp,omitempty"` // Applications which requested for the resync of network device

	LastManagedResyncReasons string `json:"lastManagedResyncReasons,omitempty"` // Reasons for last successful sync

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	ID string `json:"id,omitempty"` // Instance Uuid of the device
}
type ResponseDevicesAddDeviceKnowYourNetwork struct {
	Response *ResponseDevicesAddDeviceKnowYourNetworkResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesAddDeviceKnowYourNetworkResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesUpdateDeviceDetails struct {
	Response *ResponseDevicesUpdateDeviceDetailsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateDeviceDetailsResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute struct {
	Response []string `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateDeviceRole struct {
	Response *ResponseDevicesUpdateDeviceRoleResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesUpdateDeviceRoleResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetPollingIntervalForAllDevices struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceConfigForAllDevices struct {
	Response *[]ResponseDevicesGetDeviceConfigForAllDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceConfigForAllDevicesResponse struct {
	AttributeInfo *ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo `json:"attributeInfo,omitempty"` //

	CdpNeighbors string `json:"cdpNeighbors,omitempty"` //

	HealthMonitor string `json:"healthMonitor,omitempty"` //

	ID string `json:"id,omitempty"` //

	IntfDescription string `json:"intfDescription,omitempty"` //

	Inventory string `json:"inventory,omitempty"` //

	IPIntfBrief string `json:"ipIntfBrief,omitempty"` //

	MacAddressTable string `json:"macAddressTable,omitempty"` //

	RunningConfig string `json:"runningConfig,omitempty"` //

	SNMP string `json:"snmp,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceConfigForAllDevicesResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceConfigCount struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceCountKnowYourNetwork struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesExportDeviceList struct {
	Response *ResponseDevicesExportDeviceListResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesExportDeviceListResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetFunctionalCapabilityForDevices struct {
	Response *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponse struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	DeviceID string `json:"deviceId,omitempty"` // Device Id of the device

	FunctionalCapability *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability `json:"functionalCapability,omitempty"` //

	ID string `json:"id,omitempty"` // Deprecated
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapability struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails `json:"functionDetails,omitempty"` //

	FunctionName string `json:"functionName,omitempty"` // Name of the function

	FunctionOpState string `json:"functionOpState,omitempty"` // Operational state of the function

	ID string `json:"id,omitempty"` // Id of the function
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ID string `json:"id,omitempty"` // Deprecated

	PropertyName string `json:"propertyName,omitempty"` // Property Name of the function

	StringValue string `json:"stringValue,omitempty"` // Value for the property
}
type ResponseDevicesGetFunctionalCapabilityForDevicesResponseFunctionalCapabilityFunctionDetailsAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByID struct {
	Response *ResponseDevicesGetFunctionalCapabilityByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetFunctionalCapabilityByIDResponse struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	FunctionDetails *[]ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails `json:"functionDetails,omitempty"` //

	FunctionName string `json:"functionName,omitempty"` // Name of the function

	FunctionOpState string `json:"functionOpState,omitempty"` // Operational state of the function

	ID string `json:"id,omitempty"` // Id of the function
}
type ResponseDevicesGetFunctionalCapabilityByIDResponseAttributeInfo interface{}
type ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetails struct {
	AttributeInfo *ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ID string `json:"id,omitempty"` // Deprecated

	PropertyName string `json:"propertyName,omitempty"` // Property Name of the function

	StringValue string `json:"stringValue,omitempty"` // Value for the property
}
type ResponseDevicesGetFunctionalCapabilityByIDResponseFunctionDetailsAttributeInfo interface{}
type ResponseDevicesInventoryInsightDeviceLinkMismatchAPI struct {
	Response *[]ResponseDevicesInventoryInsightDeviceLinkMismatchAPIResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Api version
}
type ResponseDevicesInventoryInsightDeviceLinkMismatchAPIResponse struct {
	EndPortAllowedVLANIDs string `json:"endPortAllowedVlanIds,omitempty"` // End port allowed vlan ids

	EndPortNativeVLANID string `json:"endPortNativeVlanId,omitempty"` // End port native vlan id

	StartPortAllowedVLANIDs string `json:"startPortAllowedVlanIds,omitempty"` // Start port allowed vlan ids

	StartPortNativeVLANID string `json:"startPortNativeVlanId,omitempty"` // Start port native vlan id

	LinkStatus string `json:"linkStatus,omitempty"` // Link status

	EndDeviceHostName string `json:"endDeviceHostName,omitempty"` // End device hostname

	EndDeviceID string `json:"endDeviceId,omitempty"` // End device id

	EndDeviceIPAddress string `json:"endDeviceIpAddress,omitempty"` // End device ip address

	EndPortAddress string `json:"endPortAddress,omitempty"` // End port address

	EndPortDuplex string `json:"endPortDuplex,omitempty"` // End port duplex

	EndPortID string `json:"endPortId,omitempty"` // End port id

	EndPortMask string `json:"endPortMask,omitempty"` // End port mask

	EndPortName string `json:"endPortName,omitempty"` // End port name

	EndPortPepID string `json:"endPortPepId,omitempty"` // End port pep id

	EndPortSpeed string `json:"endPortSpeed,omitempty"` // End port speed

	StartDeviceHostName string `json:"startDeviceHostName,omitempty"` // Start device hostname

	StartDeviceID string `json:"startDeviceId,omitempty"` // Start device id

	StartDeviceIPAddress string `json:"startDeviceIpAddress,omitempty"` // Start device ip address

	StartPortAddress string `json:"startPortAddress,omitempty"` // Start port address

	StartPortDuplex string `json:"startPortDuplex,omitempty"` // Start port duplex

	StartPortID string `json:"startPortId,omitempty"` // Start port id

	StartPortMask string `json:"startPortMask,omitempty"` // Start port mask

	StartPortName string `json:"startPortName,omitempty"` // Start port name

	StartPortPepID string `json:"startPortPepId,omitempty"` // Start port pep id

	StartPortSpeed string `json:"startPortSpeed,omitempty"` // Start port speed

	LastUpdated string `json:"lastUpdated,omitempty"` // Last updated

	NumUpdates *float64 `json:"numUpdates,omitempty"` // Number updates

	AvgUpdateFrequency *float64 `json:"avgUpdateFrequency,omitempty"` // Average update frequency

	Type string `json:"type,omitempty"` // Type

	InstanceUUID string `json:"instanceUuid,omitempty"` // Unique instance id

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance tenant id
}
type ResponseDevicesGetNetworkDeviceByIP struct {
	Response *ResponseDevicesGetNetworkDeviceByIPResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetNetworkDeviceByIPResponse struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesGetModules struct {
	Response *[]ResponseDevicesGetModulesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetModulesResponse struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"` // Assembly Number of the module

	AssemblyRevision string `json:"assemblyRevision,omitempty"` // Assembly Revision of the module

	AttributeInfo *ResponseDevicesGetModulesResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ContainmentEntity string `json:"containmentEntity,omitempty"` // Containment Entity of the module

	Description string `json:"description,omitempty"` // Description of the module

	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"` // Entity Physical Index of the module

	ID string `json:"id,omitempty"` // ID of the module

	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"` // To mention if field is replaceable

	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer of the module

	ModuleIndex *int `json:"moduleIndex,omitempty"` // Index of the module

	Name string `json:"name,omitempty"` // Name of the module

	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational state of the module

	PartNumber string `json:"partNumber,omitempty"` // Part number of the module

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of the module

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor euipment type of the module
}
type ResponseDevicesGetModulesResponseAttributeInfo interface{}
type ResponseDevicesGetModuleCount struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetModuleInfoByID struct {
	Response *ResponseDevicesGetModuleInfoByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetModuleInfoByIDResponse struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"` // Assembly number of the module

	AssemblyRevision string `json:"assemblyRevision,omitempty"` // Assembly revision of the module

	AttributeInfo *ResponseDevicesGetModuleInfoByIDResponseAttributeInfo `json:"attributeInfo,omitempty"` // Deprecated

	ContainmentEntity string `json:"containmentEntity,omitempty"` // Containment entity of the module

	Description string `json:"description,omitempty"` // Description of the module

	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"` // Entity physical index of the module

	ID string `json:"id,omitempty"` // Id of the module

	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"` // To mention if field is replaceable

	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer of the module

	ModuleIndex *int `json:"moduleIndex,omitempty"` // Index of the module

	Name string `json:"name,omitempty"` // Name of the module

	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational state of the module

	PartNumber string `json:"partNumber,omitempty"` // Part number of the module

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of the modules

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor equipment type of the module
}
type ResponseDevicesGetModuleInfoByIDResponseAttributeInfo interface{}
type ResponseDevicesGetDeviceBySerialNumber struct {
	Response *ResponseDevicesGetDeviceBySerialNumberResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceBySerialNumberResponse struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesSyncDevices struct {
	Response *ResponseDevicesSyncDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesSyncDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetDevicesRegisteredForWsaNotification struct {
	Response *ResponseDevicesGetDevicesRegisteredForWsaNotificationResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDevicesRegisteredForWsaNotificationResponse struct {
	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ModelNumber string `json:"modelNumber,omitempty"` // Model number of the device

	Name string `json:"name,omitempty"` // Name of the device

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the device

	TenantID string `json:"tenantId,omitempty"` // Tenant Id of the device
}
type ResponseDevicesGetAllUserDefinedFields struct {
	Response *[]ResponseDevicesGetAllUserDefinedFieldsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetAllUserDefinedFieldsResponse struct {
	ID string `json:"id,omitempty"` // DeviceId of the Device

	Name string `json:"name,omitempty"` // UDF name

	Description string `json:"description,omitempty"` // Description for UDF
}
type ResponseDevicesCreateUserDefinedField struct {
	Response *ResponseDevicesCreateUserDefinedFieldResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesCreateUserDefinedFieldResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesUpdateUserDefinedField struct {
	Response *ResponseDevicesUpdateUserDefinedFieldResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateUserDefinedFieldResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesDeleteUserDefinedField struct {
	Response *ResponseDevicesDeleteUserDefinedFieldResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesDeleteUserDefinedFieldResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetChassisDetailsForDevice struct {
	Response *[]ResponseDevicesGetChassisDetailsForDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetChassisDetailsForDeviceResponse struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"` // Assembly Number of the chassis

	AssemblyRevision string `json:"assemblyRevision,omitempty"` // Assembly Revision of the chassis

	ContainmentEntity string `json:"containmentEntity,omitempty"` // Containment Entity of the chassis

	Description string `json:"description,omitempty"` // Description of the chassis

	EntityPhysicalIndex string `json:"entityPhysicalIndex,omitempty"` // Entity Physical Index of the chassis

	HardwareVersion string `json:"hardwareVersion,omitempty"` // Hardware Version of the chassis

	InstanceUUID string `json:"instanceUuid,omitempty"` // ID of the chassis

	IsFieldReplaceable string `json:"isFieldReplaceable,omitempty"` // To mention if field is replaceable

	IsReportingAlarmsAllowed string `json:"isReportingAlarmsAllowed,omitempty"` // To mention if reporting alarms are allowed

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer of the chassis

	Name string `json:"name,omitempty"` // Name of the chassis

	PartNumber string `json:"partNumber,omitempty"` // Part Number of the chassis

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the chassis

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor Equipment Type of the chassis
}
type ResponseDevicesGetStackDetailsForDevice struct {
	Response *ResponseDevicesGetStackDetailsForDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceResponse struct {
	DeviceID string `json:"deviceId,omitempty"` // Device ID

	StackPortInfo *[]ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo `json:"stackPortInfo,omitempty"` //

	StackSwitchInfo *[]ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo `json:"stackSwitchInfo,omitempty"` //

	SvlSwitchInfo *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo `json:"svlSwitchInfo,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceResponseStackPortInfo struct {
	IsSynchOk string `json:"isSynchOk,omitempty"` // If link partner sends valid protocol message

	LinkActive *bool `json:"linkActive,omitempty"` // If stack port is in same state as link partner

	LinkOk *bool `json:"linkOk,omitempty"` // If link is stable

	Name string `json:"name,omitempty"` // Name of the stack port

	NeighborPort string `json:"neighborPort,omitempty"` // Neighbor's member number and stack port number

	NrLinkOkChanges *int `json:"nrLinkOkChanges,omitempty"` // Relative stability of the link

	StackCableLengthInfo string `json:"stackCableLengthInfo,omitempty"` // Cable length

	StackPortOperStatusInfo string `json:"stackPortOperStatusInfo,omitempty"` // Port opearation status

	SwitchPort string `json:"switchPort,omitempty"` // Member number and stack port number
}
type ResponseDevicesGetStackDetailsForDeviceResponseStackSwitchInfo struct {
	EntPhysicalIndex string `json:"entPhysicalIndex,omitempty"` //

	HwPriority *int `json:"hwPriority,omitempty"` // Hardware priority of the switch

	MacAddress string `json:"macAddress,omitempty"` // Mac address of the switch

	NumNextReload *int `json:"numNextReload,omitempty"` // Stack member number to be used in next reload

	PlatformID string `json:"platformId,omitempty"` // Platform Id

	Role string `json:"role,omitempty"` // Function of the switch

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number

	SoftwareImage string `json:"softwareImage,omitempty"` // Software image type running on the switch

	StackMemberNumber *int `json:"stackMemberNumber,omitempty"` // Switch member number

	State string `json:"state,omitempty"` // Current state of the switch

	SwitchPriority *int `json:"switchPriority,omitempty"` // Priority of the switch
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfo struct {
	DadProtocol string `json:"dadProtocol,omitempty"` // Stackwise virtual dual active detection config

	DadRecoveryReloadEnabled *bool `json:"dadRecoveryReloadEnabled,omitempty"` // If dad recovery reload enabled

	DomainNumber *int `json:"domainNumber,omitempty"` // Stackwise virtual switch domain number

	InDadRecoveryMode *bool `json:"inDadRecoveryMode,omitempty"` // If in dad recovery mode

	SwVirtualStatus string `json:"swVirtualStatus,omitempty"` // Stackwise virtual status

	SwitchMembers *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers `json:"switchMembers,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembers struct {
	Bandwidth string `json:"bandwidth,omitempty"` // Bandwidth

	SvlMemberEndPoints *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints `json:"svlMemberEndPoints,omitempty"` //

	SvlMemberNumber *int `json:"svlMemberNumber,omitempty"` // Switch member number

	SvlMemberPepSettings *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings `json:"svlMemberPepSettings,omitempty"` //
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPoints struct {
	SvlMemberEndPointPorts *[]ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts `json:"svlMemberEndPointPorts,omitempty"` //

	SvlNumber *int `json:"svlNumber,omitempty"` // Stackwise virtual link number

	SvlStatus string `json:"svlStatus,omitempty"` // Stackwise virtual status
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberEndPointsSvlMemberEndPointPorts struct {
	SvlProtocolStatus string `json:"svlProtocolStatus,omitempty"` // Stackwise virtual protocol status

	SwLocalInterface string `json:"swLocalInterface,omitempty"` // Stackwise virtual local interface

	SwRemoteInterface string `json:"swRemoteInterface,omitempty"` // Stackwise virtual remote interface
}
type ResponseDevicesGetStackDetailsForDeviceResponseSvlSwitchInfoSwitchMembersSvlMemberPepSettings struct {
	DadEnabled *bool `json:"dadEnabled,omitempty"` // If dadInterface is configured for dual active detection

	DadInterfaceName string `json:"dadInterfaceName,omitempty"` // Interface for dual active detection
}
type ResponseDevicesRemoveUserDefinedFieldFromDevice struct {
	Response *ResponseDevicesRemoveUserDefinedFieldFromDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesRemoveUserDefinedFieldFromDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesAddUserDefinedFieldToDevice struct {
	Response *ResponseDevicesAddUserDefinedFieldToDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesAddUserDefinedFieldToDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice struct {
	Response *[]ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDeviceResponse struct {
	OperationalStateCode string `json:"operationalStateCode,omitempty"` // Operational State Code

	ProductID string `json:"productId,omitempty"` // Product Id

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number

	VendorEquipmentType string `json:"vendorEquipmentType,omitempty"` // Vendor Equipment Type

	Description string `json:"description,omitempty"` // Description

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Name string `json:"name,omitempty"` // Name

	Manufacturer string `json:"manufacturer,omitempty"` // Manufacturer
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice struct {
	Version string `json:"version,omitempty"` // Version

	Response *[]ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse `json:"response,omitempty"` //
}
type ResponseDevicesReturnsPoeInterfaceDetailsForTheDeviceResponse struct {
	AdminStatus string `json:"adminStatus,omitempty"` // Administration Status. Possible values: AUTO, STATIC, NEVER

	OperStatus string `json:"operStatus,omitempty"` // Operational Status. Possible values: ON, OFF, FAULTY, POWER_DENY

	InterfaceName string `json:"interfaceName,omitempty"` // Name of the interface

	MaxPortPower string `json:"maxPortPower,omitempty"` // Maximum power (in Watts) that port can hold

	AllocatedPower string `json:"allocatedPower,omitempty"` // Power (in Watts) allocated for a given interface

	PortPowerDrawn string `json:"portPowerDrawn,omitempty"` // Power (in Watts) that the port has drawn so far
}
type ResponseDevicesGetConnectedDeviceDetail struct {
	Response *ResponseDevicesGetConnectedDeviceDetailResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetConnectedDeviceDetailResponse struct {
	NeighborDevice string `json:"neighborDevice,omitempty"` // Info about the devices connected to the interface

	NeighborPort string `json:"neighborPort,omitempty"` // Info about the connected interface

	Capabilities []string `json:"capabilities,omitempty"` // Info about capabilities of the connected device
}
type ResponseDevicesGetLinecardDetails struct {
	Response *[]ResponseDevicesGetLinecardDetailsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetLinecardDetailsResponse struct {
	Serialno string `json:"serialno,omitempty"` // Serial number of the line card

	Partno string `json:"partno,omitempty"` // Part number of the line card

	Switchno string `json:"switchno,omitempty"` // Switch number of the line card

	Slotno string `json:"slotno,omitempty"` // Slot number of line card
}
type ResponseDevicesPoeDetails struct {
	Response *ResponseDevicesPoeDetailsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesPoeDetailsResponse struct {
	PowerAllocated string `json:"powerAllocated,omitempty"` // Total power available on the switch on all interfaces combined in Watts

	PowerConsumed string `json:"powerConsumed,omitempty"` // Total power being currently drawn by all interfaces combined in Watts

	PowerRemaining string `json:"powerRemaining,omitempty"` // Total power remaining in Watts (powerConsumed - powerAllocated)
}
type ResponseDevicesGetSupervisorCardDetail struct {
	Response *[]ResponseDevicesGetSupervisorCardDetailResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetSupervisorCardDetailResponse struct {
	Serialno string `json:"serialno,omitempty"` // Serial number of the supervisor card

	Partno string `json:"partno,omitempty"` // Part number of the supervisor card

	Switchno string `json:"switchno,omitempty"` // Switch number of the supervisor card

	Slotno string `json:"slotno,omitempty"` // Slot number of supervisor card
}
type ResponseDevicesUpdateDeviceManagementAddress struct {
	Response *ResponseDevicesUpdateDeviceManagementAddressResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesUpdateDeviceManagementAddressResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseDevicesGetDeviceByID struct {
	Response *ResponseDevicesGetDeviceByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceByIDResponse struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesDeleteDeviceByID struct {
	Response *ResponseDevicesDeleteDeviceByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesDeleteDeviceByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseDevicesGetDeviceSummary struct {
	Response *ResponseDevicesGetDeviceSummaryResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceSummaryResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto
}
type ResponseDevicesGetPollingIntervalByID struct {
	Response *int `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetOrganizationListForMeraki struct {
	Response []string `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceVLANs struct {
	Response *[]ResponseDevicesGetDeviceInterfaceVLANsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceVLANsResponse struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface name

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	Mask *int `json:"mask,omitempty"` // Mask IP

	NetworkAddress string `json:"networkAddress,omitempty"` // Network addresses

	NumberOfIPs *int `json:"numberOfIPs,omitempty"` // Number of Ip addresses

	Prefix string `json:"prefix,omitempty"` // Prefix associated with the IP address

	VLANNumber *int `json:"vlanNumber,omitempty"` // Vlan Number

	VLANType string `json:"vlanType,omitempty"` // [Deprecated] Description of the interface VLAN
}
type ResponseDevicesGetWirelessLanControllerDetailsByID struct {
	AdminEnabledPorts *[]int `json:"adminEnabledPorts,omitempty"` // Admin Enabled Ports of the Device

	ApGroupName string `json:"apGroupName,omitempty"` // Name of the AP Group that Access point assigned

	DeviceID string `json:"deviceId,omitempty"` // Device Id

	EthMacAddress string `json:"ethMacAddress,omitempty"` // Ethernet MacAddress of the Device

	FlexGroupName string `json:"flexGroupName,omitempty"` // Name of the Flex Group that Access point assigned

	ID string `json:"id,omitempty"` // Id of the Device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // TenantId of the Device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance UUID of the Device

	LagModeEnabled *bool `json:"lagModeEnabled,omitempty"` // LagMode status of the Device

	NetconfEnabled *bool `json:"netconfEnabled,omitempty"` // Netconf Status of the Device

	WirelessLicenseInfo string `json:"wirelessLicenseInfo,omitempty"` // License type of Wireless Device

	WirelessPackageInstalled *bool `json:"wirelessPackageInstalled,omitempty"` // Status of the Wireless Package on the Device
}
type ResponseDevicesGetDeviceConfigByID struct {
	Response string `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetNetworkDeviceByPaginationRange struct {
	Response *[]ResponseDevicesGetNetworkDeviceByPaginationRangeResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseDevicesGetNetworkDeviceByPaginationRangeResponse struct {
	ApManagerInterfaceIP string `json:"apManagerInterfaceIp,omitempty"` // IP address of WLC on AP manager interface

	AssociatedWlcIP string `json:"associatedWlcIp,omitempty"` // Associated Wlc Ip address of the AP device

	BootDateTime string `json:"bootDateTime,omitempty"` // Device boot time

	CollectionInterval string `json:"collectionInterval,omitempty"` // Re sync Interval of the device

	CollectionStatus string `json:"collectionStatus,omitempty"` // Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress

	ErrorCode string `json:"errorCode,omitempty"` // Inventory status error code

	ErrorDescription string `json:"errorDescription,omitempty"` // Inventory status description

	Family string `json:"family,omitempty"` // Family of device as switch, router, wireless lan controller, accesspoints

	Hostname string `json:"hostname,omitempty"` // Device name

	ID string `json:"id,omitempty"` // Instance Uuid of the device

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id of the device

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the device

	InterfaceCount string `json:"interfaceCount,omitempty"` // Number of interfaces on the device

	InventoryStatusDetail string `json:"inventoryStatusDetail,omitempty"` // Status detail of inventory sync

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Time in epoch when the network device info last got updated

	LastUpdated string `json:"lastUpdated,omitempty"` // Time when the network device info last got updated

	LineCardCount string `json:"lineCardCount,omitempty"` // Number of linecards on the device

	LineCardID string `json:"lineCardId,omitempty"` // IDs of linecards of the device

	Location string `json:"location,omitempty"` // [Deprecated] Location ID that is associated with the device

	LocationName string `json:"locationName,omitempty"` // [Deprecated] Name of the associated location

	MacAddress string `json:"macAddress,omitempty"` // MAC address of device

	ManagementIPAddress string `json:"managementIpAddress,omitempty"` // IP address of the device

	MemorySize string `json:"memorySize,omitempty"` // Processor memory size

	PlatformID string `json:"platformId,omitempty"` // Platform ID of device

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Failure reason for unreachable devices

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Device reachability status as Reachable / Unreachable

	Role string `json:"role,omitempty"` // Role of device as access, distribution, border router, core

	RoleSource string `json:"roleSource,omitempty"` // Role source as manual / auto

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number of device

	Series string `json:"series,omitempty"` // Device series

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact on device

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location on device

	SoftwareType string `json:"softwareType,omitempty"` // Software type on the device

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Software version on the device

	TagCount string `json:"tagCount,omitempty"` // Number of tags associated with the device

	TunnelUDPPort string `json:"tunnelUdpPort,omitempty"` // Mobility protocol port is stored as tunneludpport for WLC

	Type string `json:"type,omitempty"` // Type of device as switch, router, wireless lan controller, accesspoints

	UpTime string `json:"upTime,omitempty"` // Time that shows for how long the device has been up

	WaasDeviceMode string `json:"waasDeviceMode,omitempty"` // WAAS device mode

	DNSResolvedManagementAddress string `json:"dnsResolvedManagementAddress,omitempty"` // Specifies the resolved ip address of dns name

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // AccessPoint Ethernet MacAddress of AP device

	Vendor string `json:"vendor,omitempty"` // Vendor details

	ReasonsForPendingSyncRequests string `json:"reasonsForPendingSyncRequests,omitempty"` // Reasons for pending sync requests , if any

	PendingSyncRequestsCount string `json:"pendingSyncRequestsCount,omitempty"` // Count of pending sync requests , if any

	ReasonsForDeviceResync string `json:"reasonsForDeviceResync,omitempty"` // Reason for last/ongoing sync

	LastDeviceResyncStartTime string `json:"lastDeviceResyncStartTime,omitempty"` // Start time for last/ongoing sync

	UptimeSeconds *float64 `json:"uptimeSeconds,omitempty"` // Uptime in Seconds

	ManagedAtleastOnce *bool `json:"managedAtleastOnce,omitempty"` // Indicates if device went into Managed state atleast once

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // Support level of the device

	ManagementState string `json:"managementState,omitempty"` // Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.

	Description string `json:"description,omitempty"` // System description
}
type ResponseDevicesCreateMaintenanceScheduleForNetworkDevices struct {
	Response *ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesCreateMaintenanceScheduleForNetworkDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices struct {
	Response *[]ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponse struct {
	ID string `json:"id,omitempty"` // Id of the schedule maintenance window

	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceSchedule struct {
	StartID string `json:"startId,omitempty"` // Activity id of start schedule of the maintenance window. To check the status of the start schedule, use GET /dna/intent/api/v1/activities/{id}. startId remains same for every occurrence of recurrence instance.

	EndID string `json:"endId,omitempty"` // Activity id of end schedule of the maintenance window. To check the status of the end schedule, use GET /dna/intent/api/v1/activities/{id}. endId remains same for every occurrence of recurrence instance.

	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //

	Status string `json:"status,omitempty"` // The status of the maintenance schedule.
}
type ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevicesResponseMaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindows struct {
	Response *ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsResponse struct {
	Count *int `json:"count,omitempty"` // Count of scheduled maintenance windows
}
type ResponseDevicesUpdatesTheMaintenanceScheduleInformation struct {
	Response *ResponseDevicesUpdatesTheMaintenanceScheduleInformationResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesUpdatesTheMaintenanceScheduleInformationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformation struct {
	Response *ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponse struct {
	ID string `json:"id,omitempty"` // Id of the schedule maintenance window

	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponseMaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponseMaintenanceSchedule struct {
	StartID string `json:"startId,omitempty"` // Activity id of start schedule of the maintenance window. To check the status of the start schedule, use GET /intent/api/v1/activities/{id}. startId remains same for every occurrence of recurrence instance.

	EndID string `json:"endId,omitempty"` // Activity id of end schedule of the maintenance window. To check the status of the end schedule, use GET /intent/api/v1/activities/{id}. endId remains same for every occurrence of recurrence instance.

	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponseMaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //

	Status string `json:"status,omitempty"` // The status of the maintenance schedule.
}
type ResponseDevicesRetrievesTheMaintenanceScheduleInformationResponseMaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type ResponseDevicesDeleteMaintenanceSchedule struct {
	Response *ResponseDevicesDeleteMaintenanceScheduleResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesDeleteMaintenanceScheduleResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesRetrieveNetworkDevices struct {
	Response *[]ResponseDevicesRetrieveNetworkDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesRetrieveNetworkDevicesResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	DNSResolvedManagementIPAddress string `json:"dnsResolvedManagementIpAddress,omitempty"` // DNS-resolved management IP address of the network device

	Hostname string `json:"hostname,omitempty"` // Hostname of the network device

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the network device

	SerialNumbers []string `json:"serialNumbers,omitempty"` // Serial number of the network device. In case of stack device, there will be multiple serial numbers

	Type string `json:"type,omitempty"` // Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.

	Family string `json:"family,omitempty"` // Product family of the network device. For example, Switches, Routers, etc

	Series string `json:"series,omitempty"` // The model range or series of the network device

	Status string `json:"status,omitempty"` // Inventory related status of the network device. Refer features for more details

	PlatformIDs []string `json:"platformIds,omitempty"` // Platform identifier of the network device

	SoftwareType string `json:"softwareType,omitempty"` // Type of software running on the network device. For example, IOS-XE, etc.

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Version of the software running on the network device

	Vendor string `json:"vendor,omitempty"` // Vendor of the network device

	StackDevice *bool `json:"stackDevice,omitempty"` // Flag indicating if the network device is a stack device

	BootTime *float64 `json:"bootTime,omitempty"` // The time at which the network device was last rebooted or powered on represented as epoch in milliseconds

	Role string `json:"role,omitempty"` // Role assigned to the network device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // Ethernet MAC address of the AP network device

	ApManagerInterfaceIPAddress string `json:"apManagerInterfaceIpAddress,omitempty"` // Management IP address of the AP network device

	ApWlcIPAddress string `json:"apWlcIpAddress,omitempty"` // Management IP address of the WLC on which AP is associated to

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // The level of support Catalyst Center provides for the network device.

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location of the network device

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact of the network device

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability status of the network device. Refer features for more details

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for reachability failure. This message that provides more information about the reachability failure.

	ManagementState string `json:"managementState,omitempty"` // The status of the network device's manageability. Refer features for more details.

	LastSuccessfulResyncReasons []string `json:"lastSuccessfulResyncReasons,omitempty"` // List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncStartTime *float64 `json:"resyncStartTime,omitempty"` // Start time for the last/ongoing resync represented as epoch in milliseconds

	ResyncEndTime *float64 `json:"resyncEndTime,omitempty"` // End time for the last resync represented as epoch in milliseconds

	ResyncReasons []string `json:"resyncReasons,omitempty"` // List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncRequestedByApps []string `json:"resyncRequestedByApps,omitempty"` // List of applications that requested the last/ongoing resync on the device

	PendingResyncRequestCount *int `json:"pendingResyncRequestCount,omitempty"` // Number of pending resync requests for the device

	PendingResyncRequestReasons []string `json:"pendingResyncRequestReasons,omitempty"` // List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncIntervalSource string `json:"resyncIntervalSource,omitempty"` // Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.

	ResyncIntervalMinutes *int `json:"resyncIntervalMinutes,omitempty"` // The duration in minutes between the periodic resync attempts for the device

	ErrorCode string `json:"errorCode,omitempty"` // Error code indicating the reason for the last resync failure

	ErrorDescription string `json:"errorDescription,omitempty"` // Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.

	UserDefinedFields *ResponseDevicesRetrieveNetworkDevicesResponseUserDefinedFields `json:"userDefinedFields,omitempty"` // Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
}
type ResponseDevicesRetrieveNetworkDevicesResponseUserDefinedFields interface{}
type ResponseDevicesCountTheNumberOfNetworkDevices struct {
	Response *ResponseDevicesCountTheNumberOfNetworkDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseDevicesCountTheNumberOfNetworkDevicesResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanup struct {
	Response *ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanupResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanup struct {
	Response *ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanupResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesQueryNetworkDevicesWithFilters struct {
	Response *[]ResponseDevicesQueryNetworkDevicesWithFiltersResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesQueryNetworkDevicesWithFiltersResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	DNSResolvedManagementIPAddress string `json:"dnsResolvedManagementIpAddress,omitempty"` // DNS-resolved management IP address of the network device

	Hostname string `json:"hostname,omitempty"` // Hostname of the network device

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the network device

	SerialNumbers []string `json:"serialNumbers,omitempty"` // Serial number of the network device. In case of stack device, there will be multiple serial numbers

	Type string `json:"type,omitempty"` // Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.

	Family string `json:"family,omitempty"` // Product family of the network device. For example, Switches, Routers, etc

	Series string `json:"series,omitempty"` // The model range or series of the network device

	Status string `json:"status,omitempty"` // Inventory related status of the network device. Refer features for more details

	PlatformIDs []string `json:"platformIds,omitempty"` // Platform identifier of the network device

	SoftwareType string `json:"softwareType,omitempty"` // Type of software running on the network device. For example, IOS-XE, etc.

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Version of the software running on the network device

	Vendor string `json:"vendor,omitempty"` // Vendor of the network device

	StackDevice *bool `json:"stackDevice,omitempty"` // Flag indicating if the network device is a stack device

	BootTime *float64 `json:"bootTime,omitempty"` // The time at which the network device was last rebooted or powered on represented as epoch in milliseconds

	Role string `json:"role,omitempty"` // Role assigned to the network device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // Ethernet MAC address of the AP network device

	ApManagerInterfaceIPAddress string `json:"apManagerInterfaceIpAddress,omitempty"` // Management IP address of the AP network device

	ApWlcIPAddress string `json:"apWlcIpAddress,omitempty"` // Management IP address of the WLC on which AP is associated to

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // The level of support Catalyst Center provides for the network device.

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location of the network device

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact of the network device

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability status of the network device. Refer features for more details

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for reachability failure. This message that provides more information about the reachability failure.

	ManagementState string `json:"managementState,omitempty"` // The status of the network device's manageability. Refer features for more details.

	LastSuccessfulResyncReasons []string `json:"lastSuccessfulResyncReasons,omitempty"` // List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncStartTime *float64 `json:"resyncStartTime,omitempty"` // Start time for the last/ongoing resync represented as epoch in milliseconds

	ResyncEndTime *float64 `json:"resyncEndTime,omitempty"` // End time for the last resync represented as epoch in milliseconds

	ResyncReasons []string `json:"resyncReasons,omitempty"` // List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncRequestedByApps []string `json:"resyncRequestedByApps,omitempty"` // List of applications that requested the last/ongoing resync on the device

	PendingResyncRequestCount *int `json:"pendingResyncRequestCount,omitempty"` // Number of pending resync requests for the device

	PendingResyncRequestReasons []string `json:"pendingResyncRequestReasons,omitempty"` // List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncIntervalSource string `json:"resyncIntervalSource,omitempty"` // Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.

	ResyncIntervalMinutes *int `json:"resyncIntervalMinutes,omitempty"` // The duration in minutes between the periodic resync attempts for the device

	ErrorCode string `json:"errorCode,omitempty"` // Error code indicating the reason for the last resync failure

	ErrorDescription string `json:"errorDescription,omitempty"` // Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.

	UserDefinedFields *ResponseDevicesQueryNetworkDevicesWithFiltersResponseUserDefinedFields `json:"userDefinedFields,omitempty"` // Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
}
type ResponseDevicesQueryNetworkDevicesWithFiltersResponseUserDefinedFields interface{}
type ResponseDevicesCountTheNumberOfNetworkDevicesWithFilters struct {
	Response *ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseDevicesUpdateGlobalResyncInterval struct {
	Response *ResponseDevicesUpdateGlobalResyncIntervalResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesUpdateGlobalResyncIntervalResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesOverrideResyncInterval struct {
	Response *ResponseDevicesOverrideResyncIntervalResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesOverrideResyncIntervalResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesGetDetailsOfASingleNetworkDevice struct {
	Response *ResponseDevicesGetDetailsOfASingleNetworkDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number of the response
}
type ResponseDevicesGetDetailsOfASingleNetworkDeviceResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier of the network device

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	DNSResolvedManagementIPAddress string `json:"dnsResolvedManagementIpAddress,omitempty"` // DNS-resolved management IP address of the network device

	Hostname string `json:"hostname,omitempty"` // Hostname of the network device

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the network device

	SerialNumbers []string `json:"serialNumbers,omitempty"` // Serial number of the network device. In case of stack device, there will be multiple serial numbers

	Type string `json:"type,omitempty"` // Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.

	Family string `json:"family,omitempty"` // Product family of the network device. For example, Switches, Routers, etc

	Series string `json:"series,omitempty"` // The model range or series of the network device

	Status string `json:"status,omitempty"` // Inventory related status of the network device. Refer features for more details

	PlatformIDs []string `json:"platformIds,omitempty"` // Platform identifier of the network device

	SoftwareType string `json:"softwareType,omitempty"` // Type of software running on the network device. For example, IOS-XE, etc.

	SoftwareVersion string `json:"softwareVersion,omitempty"` // Version of the software running on the network device

	Vendor string `json:"vendor,omitempty"` // Vendor of the network device

	StackDevice *bool `json:"stackDevice,omitempty"` // Flag indicating if the network device is a stack device

	BootTime *float64 `json:"bootTime,omitempty"` // The time at which the network device was last rebooted or powered on represented as epoch in milliseconds

	Role string `json:"role,omitempty"` // Role assigned to the network device

	RoleSource string `json:"roleSource,omitempty"` // Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.

	ApEthernetMacAddress string `json:"apEthernetMacAddress,omitempty"` // Ethernet MAC address of the AP network device

	ApManagerInterfaceIPAddress string `json:"apManagerInterfaceIpAddress,omitempty"` // Management IP address of the AP network device

	ApWlcIPAddress string `json:"apWlcIpAddress,omitempty"` // Management IP address of the WLC on which AP is associated to

	DeviceSupportLevel string `json:"deviceSupportLevel,omitempty"` // The level of support Catalyst Center provides for the network device.

	SNMPLocation string `json:"snmpLocation,omitempty"` // SNMP location of the network device

	SNMPContact string `json:"snmpContact,omitempty"` // SNMP contact of the network device

	ReachabilityStatus string `json:"reachabilityStatus,omitempty"` // Reachability status of the network device. Refer features for more details

	ReachabilityFailureReason string `json:"reachabilityFailureReason,omitempty"` // Reason for reachability failure. This message that provides more information about the reachability failure.

	ManagementState string `json:"managementState,omitempty"` // The status of the network device's manageability. Refer features for more details.

	LastSuccessfulResyncReasons []string `json:"lastSuccessfulResyncReasons,omitempty"` // List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncStartTime *float64 `json:"resyncStartTime,omitempty"` // Start time for the last/ongoing resync represented as epoch in milliseconds

	ResyncEndTime *float64 `json:"resyncEndTime,omitempty"` // End time for the last resync represented as epoch in milliseconds

	ResyncReasons []string `json:"resyncReasons,omitempty"` // List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncRequestedByApps []string `json:"resyncRequestedByApps,omitempty"` // List of applications that requested the last/ongoing resync on the device

	PendingResyncRequestCount *int `json:"pendingResyncRequestCount,omitempty"` // Number of pending resync requests for the device

	PendingResyncRequestReasons []string `json:"pendingResyncRequestReasons,omitempty"` // List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC

	ResyncIntervalSource string `json:"resyncIntervalSource,omitempty"` // Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.

	ResyncIntervalMinutes *int `json:"resyncIntervalMinutes,omitempty"` // The duration in minutes between the periodic resync attempts for the device

	ErrorCode string `json:"errorCode,omitempty"` // Error code indicating the reason for the last resync failure

	ErrorDescription string `json:"errorDescription,omitempty"` // Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.

	UserDefinedFields *ResponseDevicesGetDetailsOfASingleNetworkDeviceResponseUserDefinedFields `json:"userDefinedFields,omitempty"` // Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
}
type ResponseDevicesGetDetailsOfASingleNetworkDeviceResponseUserDefinedFields interface{}
type ResponseDevicesUpdateResyncIntervalForTheNetworkDevice struct {
	Response *ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseDevicesUpdateResyncIntervalForTheNetworkDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseDevicesGetResyncIntervalForTheNetworkDevice struct {
	Response *ResponseDevicesGetResyncIntervalForTheNetworkDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetResyncIntervalForTheNetworkDeviceResponse struct {
	Interval *int `json:"interval,omitempty"` // Resync interval of the device
}
type ResponseDevicesRogueAdditionalDetails struct {
	Response *[]ResponseDevicesRogueAdditionalDetailsResponse `json:"response,omitempty"` //

	TotalCount *int `json:"totalCount,omitempty"` // Total Count

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesRogueAdditionalDetailsResponse struct {
	MacAddress string `json:"macAddress,omitempty"` // MAC Address of the Rogue BSSID

	MldMacAddress string `json:"mldMacAddress,omitempty"` // MLD MAC Address of the Rogue BSSID, this is applicable only for Wi-Fi 7 Rogues

	UpdatedTime *int `json:"updatedTime,omitempty"` // Last time when the Rogue is seen in the network

	CreatedTime *int `json:"createdTime,omitempty"` // First time when the Rogue is seen in the network

	ThreatType string `json:"threatType,omitempty"` // Type of the Rogue Threat

	ThreatLevel string `json:"threatLevel,omitempty"` // Level of the Rogue Threat

	ApName string `json:"apName,omitempty"` // Detecting AP Name

	DetectingApMac string `json:"detectingAPMac,omitempty"` // MAC Address of the Detecting AP

	SSID string `json:"ssid,omitempty"` // Rogue SSID

	Containment string `json:"containment,omitempty"` // Containment Status of the Rogue

	RadioType string `json:"radioType,omitempty"` // Radio Type on which Rogue is detected

	ControllerIP string `json:"controllerIp,omitempty"` // IP Address of the Controller detecting this Rogue

	ControllerName string `json:"controllerName,omitempty"` // Name of the Controller detecting this Rogue

	ChannelNumber string `json:"channelNumber,omitempty"` // Channel Number on which the Rogue is detected

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Hierarchy of the Rogue

	Encryption string `json:"encryption,omitempty"` // Security status of the Rogue SSID

	SwitchIP string `json:"switchIp,omitempty"` // IP Address of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type

	SwitchName string `json:"switchName,omitempty"` // Name of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type

	PortDescription string `json:"portDescription,omitempty"` // Port information of the Switch on which the Rogue is connected. This will be filled only in case of Rogue on Wire Threat Type
}
type ResponseDevicesRogueAdditionalDetailCount struct {
	Response *int `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesStartWirelessRogueApContainment struct {
	Response *ResponseDevicesStartWirelessRogueApContainmentResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesStartWirelessRogueApContainmentResponse struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	InitiatedOnWlcIP string `json:"initiatedOnWlcIp,omitempty"` // Initiated On Wlc Ip

	TaskID string `json:"taskId,omitempty"` // Task Id

	TaskType string `json:"taskType,omitempty"` // Task Type

	InitiatedOnBssid []string `json:"initiatedOnBssid,omitempty"` // Initiated On Bssid
}
type ResponseDevicesWirelessRogueApContainmentStatus struct {
	Response *[]ResponseDevicesWirelessRogueApContainmentStatusResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesWirelessRogueApContainmentStatusResponse struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	Classification string `json:"classification,omitempty"` // Classification

	ContainmentStatus string `json:"containmentStatus,omitempty"` // Containment Status

	ContainedByWlcIP []string `json:"containedByWlcIp,omitempty"` // Contained By Wlc Ip

	LastSeen *int `json:"lastSeen,omitempty"` // Last Seen

	StrongestDetectingWlcIP string `json:"strongestDetectingWlcIp,omitempty"` // Strongest Detecting Wlc Ip

	LastTaskDetail *ResponseDevicesWirelessRogueApContainmentStatusResponseLastTaskDetail `json:"lastTaskDetail,omitempty"` //

	BssidContainmentStatus *[]ResponseDevicesWirelessRogueApContainmentStatusResponseBssidContainmentStatus `json:"bssidContainmentStatus,omitempty"` //
}
type ResponseDevicesWirelessRogueApContainmentStatusResponseLastTaskDetail struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	TaskType string `json:"taskType,omitempty"` // Task Type

	TaskState string `json:"taskState,omitempty"` // Task State

	TaskStartTime *int `json:"taskStartTime,omitempty"` // Task Start Time

	InitiatedOnWlcIP string `json:"initiatedOnWlcIp,omitempty"` // Initiated On Wlc Ip

	InitiatedOnBssid []string `json:"initiatedOnBssid,omitempty"` // Initiated On Bssid
}
type ResponseDevicesWirelessRogueApContainmentStatusResponseBssidContainmentStatus struct {
	Bssid string `json:"bssid,omitempty"` // Bssid

	SSID string `json:"ssid,omitempty"` // Ssid

	RadioType string `json:"radioType,omitempty"` // Radio Type

	ContainmentStatus string `json:"containmentStatus,omitempty"` // Containment Status

	ContainedByWlcIP string `json:"containedByWlcIp,omitempty"` // Contained By Wlc Ip

	IsAdhoc *bool `json:"isAdhoc,omitempty"` // Is Adhoc
}
type ResponseDevicesStopWirelessRogueApContainment struct {
	Response *ResponseDevicesStopWirelessRogueApContainmentResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesStopWirelessRogueApContainmentResponse struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	InitiatedOnWlcIP string `json:"initiatedOnWlcIp,omitempty"` // Initiated On Wlc Ip

	TaskID string `json:"taskId,omitempty"` // Task Id

	TaskType string `json:"taskType,omitempty"` // Task Type

	InitiatedOnBssid []string `json:"initiatedOnBssid,omitempty"` // Initiated On Bssid
}
type ResponseDevicesThreatDetails struct {
	Response *[]ResponseDevicesThreatDetailsResponse `json:"response,omitempty"` //

	TotalCount *int `json:"totalCount,omitempty"` // Total Count

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesThreatDetailsResponse struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	UpdatedTime *int `json:"updatedTime,omitempty"` // Updated Time

	Vendor string `json:"vendor,omitempty"` // Vendor

	ThreatType string `json:"threatType,omitempty"` // Threat Type

	ThreatLevel string `json:"threatLevel,omitempty"` // Threat Level

	ApName string `json:"apName,omitempty"` // Ap Name

	DetectingApMac string `json:"detectingAPMac,omitempty"` // Detecting A P Mac

	SiteID string `json:"siteId,omitempty"` // Site Id

	Rssi string `json:"rssi,omitempty"` // Rssi

	SSID string `json:"ssid,omitempty"` // Ssid

	Containment string `json:"containment,omitempty"` // Containment

	State string `json:"state,omitempty"` // State

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type ResponseDevicesThreatDetailCount struct {
	Response *int `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesGetThreatLevels struct {
	Response *[]ResponseDevicesGetThreatLevelsResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetThreatLevelsResponse struct {
	Name string `json:"name,omitempty"` // Name

	Value *int `json:"value,omitempty"` // Value
}
type ResponseDevicesAddAllowedMacAddress struct {
	Response string `json:"response,omitempty"` // Response

	Error *ResponseDevicesAddAllowedMacAddressError `json:"error,omitempty"` // Error
}
type ResponseDevicesAddAllowedMacAddressError interface{}
type ResponseDevicesGetAllowedMacAddress []ResponseItemDevicesGetAllowedMacAddress // Array of ResponseDevicesGetAllowedMacAddress
type ResponseItemDevicesGetAllowedMacAddress struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Category *int `json:"category,omitempty"` // Category

	LastModified *int `json:"lastModified,omitempty"` // Last Modified
}
type ResponseDevicesGetAllowedMacAddressCount struct {
	Response *int `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesRemoveAllowedMacAddress struct {
	Response string `json:"response,omitempty"` // Response

	Error *ResponseDevicesRemoveAllowedMacAddressError `json:"error,omitempty"` // Error
}
type ResponseDevicesRemoveAllowedMacAddressError interface{}
type ResponseDevicesThreatSummary struct {
	Response *[]ResponseDevicesThreatSummaryResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseDevicesThreatSummaryResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	ThreatData *[]ResponseDevicesThreatSummaryResponseThreatData `json:"threatData,omitempty"` //
}
type ResponseDevicesThreatSummaryResponseThreatData struct {
	ThreatType string `json:"threatType,omitempty"` // Threat Type

	ThreatLevel string `json:"threatLevel,omitempty"` // Threat Level

	ThreatCount *int `json:"threatCount,omitempty"` // Threat Count
}
type ResponseDevicesGetThreatTypes struct {
	Response *[]ResponseDevicesGetThreatTypesResponse `json:"response,omitempty"` //
}
type ResponseDevicesGetThreatTypesResponse struct {
	Value *int `json:"value,omitempty"` // Value

	Name string `json:"name,omitempty"` // Name

	Label string `json:"label,omitempty"` // Label

	IsCustom *bool `json:"isCustom,omitempty"` // Is Custom

	IsDeleted *bool `json:"isDeleted,omitempty"` // Is Deleted
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2 struct {
	Version string `json:"version,omitempty"` // Version

	TotalCount *float64 `json:"totalCount,omitempty"` // The total count

	Response *[]ResponseDevicesGetDeviceInterfaceStatsInfoV2Response `json:"response,omitempty"` //

	Page *ResponseDevicesGetDeviceInterfaceStatsInfoV2Page `json:"page,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2Response struct {
	ID string `json:"id,omitempty"` // Interface Instance Id

	Values *ResponseDevicesGetDeviceInterfaceStatsInfoV2ResponseValues `json:"values,omitempty"` //
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2ResponseValues struct {
	AdminStatus string `json:"adminStatus,omitempty"` // The desired state of the interface

	DeviceID string `json:"deviceId,omitempty"` // Device Id

	DuplexConfig string `json:"duplexConfig,omitempty"` // Interface duplex config status

	DuplexOper string `json:"duplexOper,omitempty"` // Interface duplex operational status

	InterfaceID string `json:"interfaceId,omitempty"` // Interface ifIndex

	InterfaceType string `json:"interfaceType,omitempty"` // Physical or Virtual type

	InstanceID string `json:"instanceId,omitempty"` // Interface InstanceId

	IPv4Address string `json:"ipv4Address,omitempty"` // Interface IPV4 Address

	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` // List of interface IPV6 Address

	IsL3Interface string `json:"isL3Interface,omitempty"` // Interface is L3 or not

	IsWan string `json:"isWan,omitempty"` // nterface is WAN link or not

	MacAddr string `json:"macAddr,omitempty"` // Interface MAC Address

	MediaType string `json:"mediaType,omitempty"` // Interface media type

	Name string `json:"name,omitempty"` // Name of the interface

	OperStatus string `json:"operStatus,omitempty"` // Interface operational status

	PeerStackMember string `json:"peerStackMember,omitempty"` // Interface peer stack member Id

	PeerStackPort string `json:"peerStackPort,omitempty"` // Interface peer stack member port

	PortChannelID string `json:"portChannelId,omitempty"` // Interface Port-Channel Id

	PortMode string `json:"portMode,omitempty"` // Interface Port Mode

	PortType string `json:"portType,omitempty"` // Interface ifType

	Description string `json:"description,omitempty"` // Interface description

	RxDiscards string `json:"rxDiscards,omitempty"` // Rx Discards in %

	RxError string `json:"rxError,omitempty"` // Rx Errors in %

	RxRate string `json:"rxRate,omitempty"` // Rx rate in bps

	RxUtilization string `json:"rxUtilization,omitempty"` // Rx Utilization in %

	Speed string `json:"speed,omitempty"` // Speed of the Interface in kbps

	StackPortType string `json:"stackPortType,omitempty"` // Interface stack port type. SVL or DAD

	Timestamp string `json:"timestamp,omitempty"` // Interface stats collected timestamp

	TxDiscards string `json:"txDiscards,omitempty"` // Tx Discards in %

	TxError string `json:"txError,omitempty"` // Tx Errors in %

	TxRate string `json:"txRate,omitempty"` // Tx Rate in bps

	TxUtilization string `json:"txUtilization,omitempty"` // Tx  Utilization in %

	VLANID string `json:"vlanId,omitempty"` // Interface VLAN Id
}
type ResponseDevicesGetDeviceInterfaceStatsInfoV2Page struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *float64 `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Page *RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServicePage `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceFiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServicePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesQueryAssuranceEventsWithFilters struct {
	DeviceFamily []string `json:"deviceFamily,omitempty"` // Device Family

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Views []string `json:"views,omitempty"` // Views

	Filters *[]RequestDevicesQueryAssuranceEventsWithFiltersFilters `json:"filters,omitempty"` //

	Page *RequestDevicesQueryAssuranceEventsWithFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesQueryAssuranceEventsWithFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesQueryAssuranceEventsWithFiltersPage struct {
	Offset *int `json:"offset,omitempty"` // Offset

	Limit *int `json:"limit,omitempty"` // Limit

	SortBy *[]RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesQueryAssuranceEventsWithFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesCountTheNumberOfEventsWithFilters struct {
	DeviceFamily []string `json:"deviceFamily,omitempty"` // Device Family

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesCountTheNumberOfEventsWithFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesCountTheNumberOfEventsWithFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Page *RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServicePage `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceFiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServicePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Page *RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value []string `json:"value,omitempty"` // Value
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersFiltersValue interface{}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersFiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage `json:"page,omitempty"` //
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceFiltersValue interface{}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServicePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Views []string `json:"views,omitempty"` // Views

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage `json:"page,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue `json:"value,omitempty"` // Value

	Filters *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters `json:"filters,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersValue interface{}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFiltersFiltersValue interface{}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Views []string `json:"views,omitempty"` // Views

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPage `json:"page,omitempty"` //
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesFiltersValue interface{}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevicesPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendIntervalInMinutes *int `json:"trendIntervalInMinutes,omitempty"` // Trend Interval In Minutes

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRangeAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Views []string `json:"views,omitempty"` // Views

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage `json:"page,omitempty"` //
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy string `json:"sortBy,omitempty"` // Sort By
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters `json:"filters,omitempty"` //

	Views []string `json:"views,omitempty"`

	Attributes []string `json:"attributes,omitempty"`

	AggregateAttributes []RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes `json:"aggregateAttributes,omitempty"`

	Page *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage `json:"page,omitempty"`
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPage struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	SortBy *[]RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsPageSortBy struct {
	Name  string `json:"name,omitempty"`
	Order string `json:"order,omitempty"`
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`
	Function string `json:"function,omitempty"`
}
type RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage `json:"page,omitempty"` //
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TopN *int `json:"topN,omitempty"` // Top N

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Attributes *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesAttributes `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPage `json:"page,omitempty"` //
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesAttributes interface{}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	SortBy *[]RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevicesPageSortBy struct {
	Name string `json:"name,omitempty"` // Name

	Order string `json:"order,omitempty"` // Order

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesGetsTheTrendAnalyticsData struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendInterval string `json:"trendInterval,omitempty"` // Trend Interval

	GroupBy *[]RequestDevicesGetsTheTrendAnalyticsDataGroupBy `json:"groupBy,omitempty"` // Group By

	Attributes []string `json:"attributes,omitempty"` // Attributes

	Filters *[]RequestDevicesGetsTheTrendAnalyticsDataFilters `json:"filters,omitempty"` //

	AggregateAttributes *[]RequestDevicesGetsTheTrendAnalyticsDataAggregateAttributes `json:"aggregateAttributes,omitempty"` // Aggregate Attributes

	Page *RequestDevicesGetsTheTrendAnalyticsDataPage `json:"page,omitempty"` //
}
type RequestDevicesGetsTheTrendAnalyticsDataGroupBy interface{}
type RequestDevicesGetsTheTrendAnalyticsDataFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value string `json:"value,omitempty"` // Value
}
type RequestDevicesGetsTheTrendAnalyticsDataAggregateAttributes interface{}
type RequestDevicesGetsTheTrendAnalyticsDataPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	TrendIntervalInMinutes *int `json:"trendIntervalInMinutes,omitempty"` // Trend Interval In Minutes

	GroupBy []string `json:"groupBy,omitempty"` // Group By

	Filters *[]RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters `json:"filters,omitempty"` //

	Attributes []string `json:"attributes,omitempty"` // Attributes

	AggregateAttributes *[]RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes `json:"aggregateAttributes,omitempty"` //

	Page *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage `json:"page,omitempty"` //
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	LogicalOperator string `json:"logicalOperator,omitempty"` // Logical Operator

	Value *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue `json:"value,omitempty"` // Value

	Filters []string `json:"filters,omitempty"` // Filters
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeFiltersValue interface{}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangeAggregateAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Function string `json:"function,omitempty"` // Function
}
type RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRangePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestDevicesUpdatePlannedAccessPointForFloor struct {
	Attributes *RequestDevicesUpdatePlannedAccessPointForFloorAttributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Indicates that PAP is a sensor

	Location *RequestDevicesUpdatePlannedAccessPointForFloorLocation `json:"location,omitempty"` //

	Position *RequestDevicesUpdatePlannedAccessPointForFloorPosition `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]RequestDevicesUpdatePlannedAccessPointForFloorRadios `json:"radios,omitempty"` //
}
type RequestDevicesUpdatePlannedAccessPointForFloorAttributes struct {
	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point
}
type RequestDevicesUpdatePlannedAccessPointForFloorLocation struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type RequestDevicesUpdatePlannedAccessPointForFloorPosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type RequestDevicesUpdatePlannedAccessPointForFloorRadios struct {
	Antenna *RequestDevicesUpdatePlannedAccessPointForFloorRadiosAntenna `json:"antenna,omitempty"` //

	Attributes *RequestDevicesUpdatePlannedAccessPointForFloorRadiosAttributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type RequestDevicesUpdatePlannedAccessPointForFloorRadiosAntenna struct {
	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio
}
type RequestDevicesUpdatePlannedAccessPointForFloorRadiosAttributes struct {
	Channel *float64 `json:"channel,omitempty"` // Channel in which this radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	ID *int `json:"id,omitempty"` // Id of the radio

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type RequestDevicesCreatePlannedAccessPointForFloor struct {
	Attributes *RequestDevicesCreatePlannedAccessPointForFloorAttributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Indicates that PAP is a sensor

	Location *RequestDevicesCreatePlannedAccessPointForFloorLocation `json:"location,omitempty"` //

	Position *RequestDevicesCreatePlannedAccessPointForFloorPosition `json:"position,omitempty"` //

	RadioCount *int `json:"radioCount,omitempty"` // Number of radios of the planned access point

	Radios *[]RequestDevicesCreatePlannedAccessPointForFloorRadios `json:"radios,omitempty"` //
}
type RequestDevicesCreatePlannedAccessPointForFloorAttributes struct {
	CreateDate *float64 `json:"createDate,omitempty"` // Created date of the planned access point

	Domain string `json:"domain,omitempty"` // Service domain to which the planned access point belongs

	HeirarchyName string `json:"heirarchyName,omitempty"` // Hierarchy name of the planned access point

	ID *float64 `json:"id,omitempty"` // Unique id of the planned access point

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance uuid of the planned access point

	MacAddress string `json:"macAddress,omitempty"` // MAC address of the planned access point

	Name string `json:"name,omitempty"` // Display name of the planned access point

	Source string `json:"source,omitempty"` // Source of the data used to create the planned access point

	TypeString string `json:"typeString,omitempty"` // Type string representation of the planned access point
}
type RequestDevicesCreatePlannedAccessPointForFloorLocation struct {
	Altitude *float64 `json:"altitude,omitempty"` // Altitude of the planned access point's location

	Lattitude *float64 `json:"lattitude,omitempty"` // Latitude of the planned access point's location

	Longtitude *float64 `json:"longtitude,omitempty"` // Longitude of the planned access point's location
}
type RequestDevicesCreatePlannedAccessPointForFloorPosition struct {
	X *float64 `json:"x,omitempty"` // x-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Y *float64 `json:"y,omitempty"` // y-coordinate of the planned access point on the map, 0,0 point being the top-left corner

	Z *float64 `json:"z,omitempty"` // z-coordinate, or height, of the planned access point on the map
}
type RequestDevicesCreatePlannedAccessPointForFloorRadios struct {
	Antenna *RequestDevicesCreatePlannedAccessPointForFloorRadiosAntenna `json:"antenna,omitempty"` //

	Attributes *RequestDevicesCreatePlannedAccessPointForFloorRadiosAttributes `json:"attributes,omitempty"` //

	IsSensor *bool `json:"isSensor,omitempty"` // Determines if it is sensor or not
}
type RequestDevicesCreatePlannedAccessPointForFloorRadiosAntenna struct {
	AzimuthAngle *float64 `json:"azimuthAngle,omitempty"` // Azimuth angle of the antenna

	ElevationAngle *float64 `json:"elevationAngle,omitempty"` // Elevation angle of the antenna

	Gain *float64 `json:"gain,omitempty"` // Gain of the antenna

	Mode string `json:"mode,omitempty"` // Mode of the antenna associated with this radio

	Name string `json:"name,omitempty"` // Name of the antenna

	Type string `json:"type,omitempty"` // Type of the antenna associated with this radio
}
type RequestDevicesCreatePlannedAccessPointForFloorRadiosAttributes struct {
	Channel *float64 `json:"channel,omitempty"` // Channel in which this radio operates

	ChannelString string `json:"channelString,omitempty"` // Channel string representation

	ID *int `json:"id,omitempty"` // Id of the radio

	IfMode string `json:"ifMode,omitempty"` // IF mode of the radio

	IfTypeString string `json:"ifTypeString,omitempty"` // String representation of native band

	IfTypeSubband string `json:"ifTypeSubband,omitempty"` // Sub band of the radio

	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid of the radio

	SlotID *float64 `json:"slotId,omitempty"` // Slot number in which the radio resides in the parent access point

	TxPowerLevel *float64 `json:"txPowerLevel,omitempty"` // Tx Power at which this radio operates (in dBm)
}
type RequestDevicesUpdateHealthScoreDefinitions []RequestItemDevicesUpdateHealthScoreDefinitions // Array of RequestDevicesUpdateHealthScoreDefinitions
type RequestItemDevicesUpdateHealthScoreDefinitions struct {
	ID string `json:"id,omitempty"` // Id

	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Threshold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
}
type RequestDevicesUpdateHealthScoreDefinitionForTheGivenID struct {
	IncludeForOverallHealth *bool `json:"includeForOverallHealth,omitempty"` // Include For Overall Health

	ThresholdValue *float64 `json:"thresholdValue,omitempty"` // Thresehold Value

	SynchronizeToIssueThreshold *bool `json:"synchronizeToIssueThreshold,omitempty"` // Synchronize To Issue Threshold
}
type RequestDevicesUpdateInterfaceDetails struct {
	Description string `json:"description,omitempty"` // Description for the Interface

	AdminStatus string `json:"adminStatus,omitempty"` // Admin status as ('UP'/'DOWN')

	VLANID *int `json:"vlanId,omitempty"` // VLAN Id to be Updated

	VoiceVLANID *int `json:"voiceVlanId,omitempty"` // Voice Vlan Id to be Updated
}
type RequestDevicesClearMacAddressTable struct {
	Operation string `json:"operation,omitempty"` // Operation needs to be specified as 'ClearMacAddress'.

	Payload *RequestDevicesClearMacAddressTablePayload `json:"payload,omitempty"` // Payload is not applicable
}
type RequestDevicesClearMacAddressTablePayload interface{}
type RequestDevicesAddDeviceKnowYourNetwork struct {
	CliTransport string `json:"cliTransport,omitempty"` // CLI transport. Supported values: telnet, ssh. Required if type is NETWORK_DEVICE.

	ComputeDevice *bool `json:"computeDevice,omitempty"` // Compute Device or not. Options are true / false.

	EnablePassword string `json:"enablePassword,omitempty"` // CLI enable password of the device. Required if device is configured to use enable password.

	ExtendedDiscoveryInfo string `json:"extendedDiscoveryInfo,omitempty"` // This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.

	HTTPPassword string `json:"httpPassword,omitempty"` // HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE.

	HTTPPort string `json:"httpPort,omitempty"` // HTTP port of the device. Required if type is COMPUTE_DEVICE.

	HTTPSecure *bool `json:"httpSecure,omitempty"` // Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP. Default is true.

	HTTPUserName string `json:"httpUserName,omitempty"` // HTTP Username of the device. Required if type is COMPUTE_DEVICE.

	IPAddress []string `json:"ipAddress,omitempty"` // IP Address of the device. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.

	MerakiOrgID []string `json:"merakiOrgId,omitempty"` // Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.

	NetconfPort string `json:"netconfPort,omitempty"` // Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided. Netconf port is required for eWLC.

	Password string `json:"password,omitempty"` // CLI Password of the device. Required if type is NETWORK_DEVICE.

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.

	SNMPAuthPassphrase string `json:"snmpAuthPassphrase,omitempty"` // SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv.

	SNMPAuthProtocol string `json:"snmpAuthProtocol,omitempty"` // SNMPv3 auth protocol. Supported values: sha, md5. Required if snmpMode is authNoPriv or authPriv.

	SNMPMode string `json:"snmpMode,omitempty"` // SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3.

	SNMPPrivPassphrase string `json:"snmpPrivPassphrase,omitempty"` // SNMPv3 priv passphrase. Required if snmpMode is authPriv.

	SNMPPrivProtocol string `json:"snmpPrivProtocol,omitempty"` // SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv.

	SNMPROCommunity string `json:"snmpROCommunity,omitempty"` // SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.

	SNMPRWCommunity string `json:"snmpRWCommunity,omitempty"` // SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required.

	SNMPRetry *int `json:"snmpRetry,omitempty"` // SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.

	SNMPTimeout *int `json:"snmpTimeout,omitempty"` // SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.

	SNMPUserName string `json:"snmpUserName,omitempty"` // SNMPV3 user name of the device. Required if snmpVersion is v3.

	SNMPVersion string `json:"snmpVersion,omitempty"` // SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE.

	Type string `json:"type,omitempty"` // Type of device being added. Default is NETWORK_DEVICE.

	UserName string `json:"userName,omitempty"` // CLI user name of the device. Required if type is NETWORK_DEVICE.
}
type RequestDevicesUpdateDeviceDetails struct {
	CliTransport string `json:"cliTransport,omitempty"` // CLI transport. Supported values: telnet, ssh. Use NO!$DATA!$ if no change is required. Required if type is NETWORK_DEVICE.

	ComputeDevice *bool `json:"computeDevice,omitempty"` // Compute Device or not. Options are true / false.

	EnablePassword string `json:"enablePassword,omitempty"` // CLI enable password of the device. Required if device is configured to use enable password. Use NO!$DATA!$ if no change is required.

	ExtendedDiscoveryInfo string `json:"extendedDiscoveryInfo,omitempty"` // This field holds that info as whether to add device with canned data or not. Supported values: DISCOVER_WITH_CANNED_DATA.

	HTTPPassword string `json:"httpPassword,omitempty"` // HTTP password of the device / API key for Meraki Dashboard. Required if type is MERAKI_DASHBOARD or COMPUTE_DEVICE. Use NO!$DATA!$ if no change is required.

	HTTPPort string `json:"httpPort,omitempty"` // HTTP port of the device. Required if type is COMPUTE_DEVICE.

	HTTPSecure *bool `json:"httpSecure,omitempty"` // Flag to select HTTP / HTTPS protocol. Options are true / false. true for HTTPS and false for HTTP.

	HTTPUserName string `json:"httpUserName,omitempty"` // HTTP Username of the device. Required if type is COMPUTE_DEVICE. Use NO!$DATA!$ if no change is required.

	IPAddress []string `json:"ipAddress,omitempty"` // IP Address of the device. Required. Use 'api.meraki.com' for Meraki Dashboard.

	MerakiOrgID []string `json:"merakiOrgId,omitempty"` // Selected Meraki organization for which the devices needs to be imported. Required if type is MERAKI_DASHBOARD.

	NetconfPort string `json:"netconfPort,omitempty"` // Netconf Port of the device. cliTransport must be 'ssh' if netconf is provided. Netconf port is required for eWLC.

	Password string `json:"password,omitempty"` // CLI Password of the device. Required if type is NETWORK_DEVICE. Use NO!$DATA!$ if no change is required.

	SerialNumber string `json:"serialNumber,omitempty"` // Serial Number of the Device. Required if extendedDiscoveryInfo is 'DISCOVER_WITH_CANNED_DATA'.

	SNMPAuthPassphrase string `json:"snmpAuthPassphrase,omitempty"` // SNMPv3 auth passphrase of the device. Required if snmpMode is authNoPriv or authPriv. Use NO!$DATA!$ if no change is required.

	SNMPAuthProtocol string `json:"snmpAuthProtocol,omitempty"` // SNMPv3 auth protocol. Supported values: sha, md5.  Required if snmpMode is authNoPriv or authPriv. Use NODATACHANGE if no change is required.

	SNMPMode string `json:"snmpMode,omitempty"` // SNMPv3 mode. Supported values: noAuthnoPriv, authNoPriv, authPriv. Required if snmpVersion is v3. Use NODATACHANGE if no change is required.

	SNMPPrivPassphrase string `json:"snmpPrivPassphrase,omitempty"` // SNMPv3 priv passphrase. Required if snmpMode is authPriv. Use NO!$DATA!$ if no change is required.

	SNMPPrivProtocol string `json:"snmpPrivProtocol,omitempty"` // SNMPv3 priv protocol. Supported values: AES128. Required if snmpMode is authPriv. Use NODATACHANGE if no change is required.

	SNMPROCommunity string `json:"snmpROCommunity,omitempty"` // SNMP Read Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required. Use NO!$DATA!$ if no change is required.

	SNMPRWCommunity string `json:"snmpRWCommunity,omitempty"` // SNMP Write Community of the device. If snmpVersion is v2, at least one of snmpROCommunity and snmpRWCommunity is required. Use NO!$DATA!$ if no change is required.

	SNMPRetry *int `json:"snmpRetry,omitempty"` // SNMP retry count. Max value supported is 3. Default is Global SNMP retry (if exists) or 3.

	SNMPTimeout *int `json:"snmpTimeout,omitempty"` // SNMP timeout in seconds. Max value supported is 300. Default is Global SNMP timeout (if exists) or 5.

	SNMPUserName string `json:"snmpUserName,omitempty"` // SNMPV3 user name of the device. Required if snmpVersion is v3. Use NO!$DATA!$ if no change is required.

	SNMPVersion string `json:"snmpVersion,omitempty"` // SNMP version. Values supported: v2, v3. Required if type is NETWORK_DEVICE, COMPUTE_DEVICE or THIRD_PARTY_DEVICE. Use NODATACHANGE if no change is required.

	Type string `json:"type,omitempty"` // Type of device being edited. Default is NETWORK_DEVICE.

	UpdateMgmtIPaddressList *[]RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList `json:"updateMgmtIPaddressList,omitempty"` //

	UserName string `json:"userName,omitempty"` // CLI user name of the device. Required if type is NETWORK_DEVICE. Use NO!$DATA!$ if no change is required.
}
type RequestDevicesUpdateDeviceDetailsUpdateMgmtIPaddressList struct {
	ExistMgmtIPAddress string `json:"existMgmtIpAddress,omitempty"` // existMgmtIpAddress IP Address of the device.

	NewMgmtIPAddress string `json:"newMgmtIpAddress,omitempty"` // New IP Address to be Updated.
}
type RequestDevicesUpdateDeviceRole struct {
	ID string `json:"id,omitempty"` // DeviceId of the Device

	Role string `json:"role,omitempty"` // Role of device as ACCESS, CORE, DISTRIBUTION, BORDER ROUTER

	RoleSource string `json:"roleSource,omitempty"` // Role source as MANUAL / AUTO
}
type RequestDevicesExportDeviceList struct {
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` // List of device uuids

	OperationEnum string `json:"operationEnum,omitempty"` // 0 to export Device Credential Details Or 1 to export Device Details

	Parameters []string `json:"parameters,omitempty"` // List of device parameters that needs to be exported to file

	Password string `json:"password,omitempty"` // Password is required when the operationEnum value is 0
}
type RequestDevicesSyncDevices []string // Array of RequestDevicesSyncDevices
type RequestDevicesCreateUserDefinedField struct {
	Name string `json:"name,omitempty"` // Name of UDF

	Description string `json:"description,omitempty"` // Description of UDF
}
type RequestDevicesUpdateUserDefinedField struct {
	Name string `json:"name,omitempty"` // Name of UDF

	Description string `json:"description,omitempty"` // Description of UDF
}
type RequestDevicesAddUserDefinedFieldToDevice []RequestItemDevicesAddUserDefinedFieldToDevice // Array of RequestDevicesAddUserDefinedFieldToDevice
type RequestItemDevicesAddUserDefinedFieldToDevice struct {
	Name string `json:"name,omitempty"` // Name of the User Defined Field

	Value string `json:"value,omitempty"` // Value of the User Defined Field that will be assigned to the device
}
type RequestDevicesUpdateDeviceManagementAddress struct {
	NewIP string `json:"newIP,omitempty"` // New IP Address of the device to be Updated
}
type RequestDevicesCreateMaintenanceScheduleForNetworkDevices struct {
	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceSchedule struct {
	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //
}
type RequestDevicesCreateMaintenanceScheduleForNetworkDevicesMaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type RequestDevicesUpdatesTheMaintenanceScheduleInformation struct {
	Description string `json:"description,omitempty"` // A brief narrative describing the maintenance schedule.

	MaintenanceSchedule *RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceSchedule `json:"maintenanceSchedule,omitempty"` //

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // List of network device ids. This field is applicable only during creation of schedules; for updates, it is read-only.
}
type RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceSchedule struct {
	StartTime *float64 `json:"startTime,omitempty"` // Start time indicates the beginning of the maintenance window in Unix epoch time in milliseconds.

	EndTime *float64 `json:"endTime,omitempty"` // End time indicates the ending of the maintenance window in Unix epoch time in milliseconds.

	Recurrence *RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceScheduleRecurrence `json:"recurrence,omitempty"` //
}
type RequestDevicesUpdatesTheMaintenanceScheduleInformationMaintenanceScheduleRecurrence struct {
	Interval *int `json:"interval,omitempty"` // Interval for recurrence in days. The interval must be longer than the duration of the schedules. The maximum allowed interval is 365 days.

	RecurrenceEndTime *float64 `json:"recurrenceEndTime,omitempty"` // The end date for the recurrence in Unix epoch time in milliseconds. Recurrence end time should be greater than maintenance end date/time.
}
type RequestDevicesDeleteNetworkDeviceWithConfigurationCleanup struct {
	ID string `json:"id,omitempty"` // The unique identifier of the network device to be deleted
}
type RequestDevicesDeleteANetworkDeviceWithoutConfigurationCleanup struct {
	ID string `json:"id,omitempty"` // The unique identifier of the network device to be deleted
}
type RequestDevicesQueryNetworkDevicesWithFilters struct {
	Filter *RequestDevicesQueryNetworkDevicesWithFiltersFilter `json:"filter,omitempty"` //

	Views []string `json:"views,omitempty"` // The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views.

	Page *RequestDevicesQueryNetworkDevicesWithFiltersPage `json:"page,omitempty"` //
}
type RequestDevicesQueryNetworkDevicesWithFiltersFilter struct {
	LogicalOperator string `json:"logicalOperator,omitempty"` // The logical operator to use for combining the filter criteria. If not provided, the default value is AND.

	Filters *[]RequestDevicesQueryNetworkDevicesWithFiltersFilterFilters `json:"filters,omitempty"` //
}
type RequestDevicesQueryNetworkDevicesWithFiltersFilterFilters struct {
	Key string `json:"key,omitempty"` // The key to filter by

	Operator string `json:"operator,omitempty"` // The operator to use for filtering the values

	Value *RequestDevicesQueryNetworkDevicesWithFiltersFilterFiltersValue `json:"value,omitempty"` // Value to filter by. For `in` operator, the value should be a list of values.
}
type RequestDevicesQueryNetworkDevicesWithFiltersFilterFiltersValue interface{}
type RequestDevicesQueryNetworkDevicesWithFiltersPage struct {
	SortBy *RequestDevicesQueryNetworkDevicesWithFiltersPageSortBy `json:"sortBy,omitempty"` //

	Limit *int `json:"limit,omitempty"` // The number of records to show for this page. Min: 1, Max: 500

	Offset *int `json:"offset,omitempty"` // The first record to show for this page; the first record is numbered 1.
}
type RequestDevicesQueryNetworkDevicesWithFiltersPageSortBy struct {
	Name string `json:"name,omitempty"` // The field to sort by. Default is hostname.

	Order string `json:"order,omitempty"` // The order to sort by.
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFilters struct {
	Filter *RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilter `json:"filter,omitempty"` //
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilter struct {
	LogicalOperator string `json:"logicalOperator,omitempty"` // The logical operator to use for combining the filter criteria. If not provided, the default value is AND.

	Filters *[]RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFilters `json:"filters,omitempty"` //
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFilters struct {
	Key string `json:"key,omitempty"` // The key to filter by

	Operator string `json:"operator,omitempty"` // The operator to use for filtering the values

	Value *RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersValue `json:"value,omitempty"` // Value to filter by. For `in` operator, the value should be a list of values.
}
type RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersValue interface{}
type RequestDevicesUpdateGlobalResyncInterval struct {
	Interval *int `json:"interval,omitempty"` // Resync Interval should be between 25 to 1440 minutes
}
type RequestDevicesUpdateResyncIntervalForTheNetworkDevice struct {
	Interval *int `json:"interval,omitempty"` // Resync interval in minutes. To disable periodic resync, set interval as `0`. To use global settings, set interval as `null`.
}
type RequestDevicesRogueAdditionalDetails struct {
	Offset *float64 `json:"offset,omitempty"` // The offset of the first item in the collection to return. Default value is 1

	Limit *float64 `json:"limit,omitempty"` // The maximum number of entries to return. Default value is 1000

	StartTime *float64 `json:"startTime,omitempty"` // This is the epoch start time in milliseconds from which data need to be fetched. Default value is 24 hours earlier to endTime

	EndTime *float64 `json:"endTime,omitempty"` // This is the epoch end time in milliseconds upto which data need to be fetched. Default value is current time

	SiteID []string `json:"siteId,omitempty"` // Filter Rogues by location. Site IDs information can be fetched from "Get Site" API

	ThreatLevel []string `json:"threatLevel,omitempty"` // Filter Rogues by Threat Level. Threat Level information can be fetched from "Get Threat Levels" API

	ThreatType []string `json:"threatType,omitempty"` // Filter Rogues by Threat Type. Threat Type information can be fetched from "Get Threat Types" API
}
type RequestDevicesRogueAdditionalDetailCount struct {
	StartTime *float64 `json:"startTime,omitempty"` // This is the epoch start time in milliseconds from which data need to be fetched. Default value is 24 hours earlier to endTime

	EndTime *float64 `json:"endTime,omitempty"` // This is the epoch end time in milliseconds upto which data need to be fetched. Default value is current time

	SiteID []string `json:"siteId,omitempty"` // Filter Rogues by location. Site IDs information can be fetched from "Get Site" API

	ThreatLevel []string `json:"threatLevel,omitempty"` // This information can be fetched from "Get Threat Levels" API

	ThreatType []string `json:"threatType,omitempty"` // This information can be fetched from "Get Threat Types" API
}
type RequestDevicesStartWirelessRogueApContainment struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type
}
type RequestDevicesStopWirelessRogueApContainment struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Type *int `json:"type,omitempty"` // Type

	WlcIP string `json:"wlcIp,omitempty"` // Wlc Ip
}
type RequestDevicesThreatDetails struct {
	Offset *int `json:"offset,omitempty"` // Offset

	Limit *int `json:"limit,omitempty"` // Limit

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteID []string `json:"siteId,omitempty"` // Site Id

	ThreatLevel []string `json:"threatLevel,omitempty"` // Threat Level

	ThreatType []string `json:"threatType,omitempty"` // Threat Type

	IsNewThreat *bool `json:"isNewThreat,omitempty"` // Is New Threat
}
type RequestDevicesThreatDetailCount struct {
	Offset *int `json:"offset,omitempty"` // Offset

	Limit *int `json:"limit,omitempty"` // Limit

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteID []string `json:"siteId,omitempty"` // Site Id

	ThreatLevel []string `json:"threatLevel,omitempty"` // Threat Level

	ThreatType []string `json:"threatType,omitempty"` // Threat Type

	IsNewThreat *bool `json:"isNewThreat,omitempty"` // Is New Threat
}
type RequestDevicesAddAllowedMacAddress []RequestItemDevicesAddAllowedMacAddress // Array of RequestDevicesAddAllowedMacAddress
type RequestItemDevicesAddAllowedMacAddress struct {
	MacAddress string `json:"macAddress,omitempty"` // Mac Address

	Category *int `json:"category,omitempty"` // Category
}
type RequestDevicesThreatSummary struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	SiteID []string `json:"siteId,omitempty"` // Site Id

	ThreatLevel []string `json:"threatLevel,omitempty"` // Threat Level

	ThreatType []string `json:"threatType,omitempty"` // Threat Type
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2 struct {
	StartTime *int `json:"startTime,omitempty"` // UTC epoch timestamp in milliseconds

	EndTime *int `json:"endTime,omitempty"` // UTC epoch timestamp in milliseconds

	Query *RequestDevicesGetDeviceInterfaceStatsInfoV2Query `json:"query,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2Query struct {
	Fields *[]RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFields `json:"fields,omitempty"` // Required field names, default ALL

	Filters *[]RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFilters `json:"filters,omitempty"` //

	Page *RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPage `json:"page,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFields interface{}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryFilters struct {
	Key string `json:"key,omitempty"` // Name of the field that the filter should be applied to

	Operator string `json:"operator,omitempty"` // Supported operators are eq,in,like

	Value string `json:"value,omitempty"` // Value of the field
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPage struct {
	Limit *int `json:"limit,omitempty"` // Number of records, Max is 1000

	Offset *float64 `json:"offset,omitempty"` // Record offset value, default 0

	OrderBy *[]RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPageOrderBy `json:"orderBy,omitempty"` //
}
type RequestDevicesGetDeviceInterfaceStatsInfoV2QueryPageOrderBy struct {
	Name string `json:"name,omitempty"` // Name of the field used to sort

	Order string `json:"order,omitempty"` // Possible values asc, des
}

//RetrievesTheListOfAAAServicesForGivenParameters Retrieves the list of AAA Services for given parameters. - b990-0822-4c08-a780
/* Retrieves the list of AAA Services and offers basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheListOfAAAServicesForGivenParametersHeaderParams Custom header parameters
@param RetrievesTheListOfAAAServicesForGivenParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-a-a-a-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheListOfAAAServicesForGivenParameters(RetrievesTheListOfAAAServicesForGivenParametersHeaderParams *RetrievesTheListOfAAAServicesForGivenParametersHeaderParams, RetrievesTheListOfAAAServicesForGivenParametersQueryParams *RetrievesTheListOfAAAServicesForGivenParametersQueryParams) (*ResponseDevicesRetrievesTheListOfAAAServicesForGivenParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices"

	queryString, _ := query.Values(RetrievesTheListOfAAAServicesForGivenParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfAAAServicesForGivenParametersHeaderParams != nil {

		if RetrievesTheListOfAAAServicesForGivenParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfAAAServicesForGivenParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheListOfAAAServicesForGivenParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfAAAServicesForGivenParameters(RetrievesTheListOfAAAServicesForGivenParametersHeaderParams, RetrievesTheListOfAAAServicesForGivenParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfAAAServicesForGivenParameters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfAAAServicesForGivenParameters)
	return result, response, err

}

//RetrievesTheTotalNumberOfAAAServicesForGivenParameters Retrieves the total number of AAA Services for given parameters. - c393-f961-4939-b53c
/* Retrieves the total number of AAA Services for given parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfAAAServicesForGivenParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-a-a-a-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfAAAServicesForGivenParameters(RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams *RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams, RetrievesTheTotalNumberOfAAAServicesForGivenParametersQueryParams *RetrievesTheTotalNumberOfAAAServicesForGivenParametersQueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfAAAServicesForGivenParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams != nil {

		if RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfAAAServicesForGivenParameters(RetrievesTheTotalNumberOfAAAServicesForGivenParametersHeaderParams, RetrievesTheTotalNumberOfAAAServicesForGivenParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfAAAServicesForGivenParameters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenParameters)
	return result, response, err

}

//RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService Retrieves the details of a specific AAA Service matching the id of the Service. - 35a1-3ae4-4cbb-ae6f
/* Retrieves the details of the AAA Service matching the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the AAA Service. It is the combination of AAA Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceHeaderParams Custom header parameters
@param RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-details-of-a-specific-a-a-a-service-matching-the-id-of-the-service
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService(id string, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceHeaderParams *RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceHeaderParams, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceQueryParams *RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheServiceQueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceHeaderParams != nil {

		if RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService(id, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceHeaderParams, RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheServiceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDetailsOfASpecificAAAServiceMatchingTheIdOfTheService")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheDetailsOfASpecificAAAServiceMatchingTheIDOfTheService)
	return result, response, err

}

//QueryAssuranceEvents Query assurance events - 15a5-3b2c-4908-8ba3
/* Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param QueryAssuranceEventsHeaderParams Custom header parameters
@param QueryAssuranceEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-assurance-events
*/
func (s *DevicesService) QueryAssuranceEvents(QueryAssuranceEventsHeaderParams *QueryAssuranceEventsHeaderParams, QueryAssuranceEventsQueryParams *QueryAssuranceEventsQueryParams) (*ResponseDevicesQueryAssuranceEvents, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents"

	queryString, _ := query.Values(QueryAssuranceEventsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryAssuranceEventsHeaderParams != nil {

		if QueryAssuranceEventsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryAssuranceEventsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesQueryAssuranceEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAssuranceEvents(QueryAssuranceEventsHeaderParams, QueryAssuranceEventsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation QueryAssuranceEvents")
	}

	result := response.Result().(*ResponseDevicesQueryAssuranceEvents)
	return result, response, err

}

//CountTheNumberOfEvents Count the number of events - 349f-a9d8-4a6a-b951
/* API to fetch the count of assurance events that match the filter criteria. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param CountTheNumberOfEventsHeaderParams Custom header parameters
@param CountTheNumberOfEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-events
*/
func (s *DevicesService) CountTheNumberOfEvents(CountTheNumberOfEventsHeaderParams *CountTheNumberOfEventsHeaderParams, CountTheNumberOfEventsQueryParams *CountTheNumberOfEventsQueryParams) (*ResponseDevicesCountTheNumberOfEvents, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/count"

	queryString, _ := query.Values(CountTheNumberOfEventsQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountTheNumberOfEventsHeaderParams != nil {

		if CountTheNumberOfEventsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountTheNumberOfEventsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesCountTheNumberOfEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfEvents(CountTheNumberOfEventsHeaderParams, CountTheNumberOfEventsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountTheNumberOfEvents")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfEvents)
	return result, response, err

}

//GetDetailsOfASingleAssuranceEvent Get details of a single assurance event - 039e-2909-449a-8f51
/* API to fetch the details of an assurance event using event `id`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param id id path parameter. Unique identifier for the event

@param GetDetailsOfASingleAssuranceEventHeaderParams Custom header parameters
@param GetDetailsOfASingleAssuranceEventQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-assurance-event
*/
func (s *DevicesService) GetDetailsOfASingleAssuranceEvent(id string, GetDetailsOfASingleAssuranceEventHeaderParams *GetDetailsOfASingleAssuranceEventHeaderParams, GetDetailsOfASingleAssuranceEventQueryParams *GetDetailsOfASingleAssuranceEventQueryParams) (*ResponseDevicesGetDetailsOfASingleAssuranceEvent, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDetailsOfASingleAssuranceEventQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDetailsOfASingleAssuranceEventHeaderParams != nil {

		if GetDetailsOfASingleAssuranceEventHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetDetailsOfASingleAssuranceEventHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDetailsOfASingleAssuranceEvent{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleAssuranceEvent(id, GetDetailsOfASingleAssuranceEventHeaderParams, GetDetailsOfASingleAssuranceEventQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleAssuranceEvent")
	}

	result := response.Result().(*ResponseDevicesGetDetailsOfASingleAssuranceEvent)
	return result, response, err

}

//GetListOfChildEventsForTheGivenWirelessClientEvent Get list of child events for the given wireless client event - d78f-7acc-4a88-b616
/* Wireless client event could have child events and this API can be used to fetch the same using parent event `id` as the input. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param id id path parameter. Unique identifier for the event

@param GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-child-events-for-the-given-wireless-client-event
*/
func (s *DevicesService) GetListOfChildEventsForTheGivenWirelessClientEvent(id string, GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams *GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams) (*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/{id}/childEvents"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams != nil {

		if GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfChildEventsForTheGivenWirelessClientEvent(id, GetListOfChildEventsForTheGivenWirelessClientEventHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfChildEventsForTheGivenWirelessClientEvent")
	}

	result := response.Result().(*ResponseDevicesGetListOfChildEventsForTheGivenWirelessClientEvent)
	return result, response, err

}

//RetrievesTheListOfDHCPServicesForGivenParameters Retrieves the list of DHCP Services for given parameters. - bfa0-ebff-418a-b093
/* Retrieves the list of DHCP Services and offers basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams Custom header parameters
@param RetrievesTheListOfDHCPServicesForGivenParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-h-c-p-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheListOfDHCPServicesForGivenParameters(RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams *RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams, RetrievesTheListOfDHCPServicesForGivenParametersQueryParams *RetrievesTheListOfDHCPServicesForGivenParametersQueryParams) (*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices"

	queryString, _ := query.Values(RetrievesTheListOfDHCPServicesForGivenParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams != nil {

		if RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDHCPServicesForGivenParameters(RetrievesTheListOfDHCPServicesForGivenParametersHeaderParams, RetrievesTheListOfDHCPServicesForGivenParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDHCPServicesForGivenParameters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenParameters)
	return result, response, err

}

//RetrievesTheTotalNumberOfDHCPServicesForGivenParameters Retrieves the total number of DHCP Services for given parameters. - 8eaf-6891-4319-9f95
/* Retrieves the total number of DHCP Services for given parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-h-c-p-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDHCPServicesForGivenParameters(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams *RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams, RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams *RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams != nil {

		if RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDHCPServicesForGivenParameters(RetrievesTheTotalNumberOfDHCPServicesForGivenParametersHeaderParams, RetrievesTheTotalNumberOfDHCPServicesForGivenParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDHCPServicesForGivenParameters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenParameters)
	return result, response, err

}

//RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService Retrieves the details of a specific DHCP Service matching the id of the Service. - 3287-8874-4319-8db1
/* Retrieves the details of the DHCP Service matching the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DHCP Service. It is the combination of DHCP Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceHeaderParams Custom header parameters
@param RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-details-of-a-specific-d-h-c-p-service-matching-the-id-of-the-service
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService(id string, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceHeaderParams *RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceHeaderParams, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceQueryParams *RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheServiceQueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceHeaderParams != nil {

		if RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService(id, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceHeaderParams, RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheServiceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIdOfTheService")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheDetailsOfASpecificDHCPServiceMatchingTheIDOfTheService)
	return result, response, err

}

//RetrievesTheListOfDNSServicesForGivenParameters Retrieves the list of DNS Services for given parameters. - 0bbd-2bd5-4f9b-9a57
/* Retrieves the list of DNS Services and offers basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDNSServicesForGivenParametersHeaderParams Custom header parameters
@param RetrievesTheListOfDNSServicesForGivenParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-n-s-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheListOfDNSServicesForGivenParameters(RetrievesTheListOfDNSServicesForGivenParametersHeaderParams *RetrievesTheListOfDNSServicesForGivenParametersHeaderParams, RetrievesTheListOfDNSServicesForGivenParametersQueryParams *RetrievesTheListOfDNSServicesForGivenParametersQueryParams) (*ResponseDevicesRetrievesTheListOfDNSServicesForGivenParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices"

	queryString, _ := query.Values(RetrievesTheListOfDNSServicesForGivenParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDNSServicesForGivenParametersHeaderParams != nil {

		if RetrievesTheListOfDNSServicesForGivenParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDNSServicesForGivenParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheListOfDNSServicesForGivenParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDNSServicesForGivenParameters(RetrievesTheListOfDNSServicesForGivenParametersHeaderParams, RetrievesTheListOfDNSServicesForGivenParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDNSServicesForGivenParameters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDNSServicesForGivenParameters)
	return result, response, err

}

//RetrievesTheTotalNumberOfDNSServicesForGivenParameters Retrieves the total number of DNS Services for given parameters. - 4385-991e-43a9-9561
/* Retrieves the total number of DNS Services for given parameters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfDNSServicesForGivenParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-n-s-services-for-given-parameters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDNSServicesForGivenParameters(RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams *RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams, RetrievesTheTotalNumberOfDNSServicesForGivenParametersQueryParams *RetrievesTheTotalNumberOfDNSServicesForGivenParametersQueryParams) (*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfDNSServicesForGivenParametersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams != nil {

		if RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDNSServicesForGivenParameters(RetrievesTheTotalNumberOfDNSServicesForGivenParametersHeaderParams, RetrievesTheTotalNumberOfDNSServicesForGivenParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDNSServicesForGivenParameters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenParameters)
	return result, response, err

}

//RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService Retrieves the details of a specific DNS Service matching the id of the Service. - 84ab-b9c3-498a-b6a7
/* Retrieves the details of the DNS Service matching the given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DNS Service. It is the combination of DNS Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceHeaderParams Custom header parameters
@param RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-details-of-a-specific-d-n-s-service-matching-the-id-of-the-service
*/
func (s *DevicesService) RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService(id string, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceHeaderParams *RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceHeaderParams, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceQueryParams *RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheServiceQueryParams) (*ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceHeaderParams != nil {

		if RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService(id, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceHeaderParams, RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheServiceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDetailsOfASpecificDNSServiceMatchingTheIdOfTheService")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheDetailsOfASpecificDNSServiceMatchingTheIDOfTheService)
	return result, response, err

}

//GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices Gets interfaces along with statistics and poe data from all network devices. - 9898-9b5a-445b-884f
/* Retrieves the list of the interfaces from all network devices based on the provided query parameters. The latest interfaces data in the specified start and end time range will be returned. When there is no start and end time specified returns the latest available data.

The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.


The supported sorting options are:
name, adminStatus, description, duplexConfig, duplexOper, interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus, portChannelId, portMode, portType,speed, vlanId



This API can paginate up to 500,000 records, please narrow matching results with additional filters beyond that value. The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.

 The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper,interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus,portChannelId, portMode, portType,speed, vlanId,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-interfaces-along-with-statistics-and-poe-data-from-all-network-devices
*/
func (s *DevicesService) GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesQueryParams *GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesQueryParams) (*ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces"

	queryString, _ := query.Values(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices(GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetsInterfacesAlongWithStatisticsAndPoeDataFromAllNetworkDevices)
	return result, response, err

}

//GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount Gets the total Network device interface counts in the specified time range. When there is no start and end time specified returns the latest interfaces total count. - 40ab-799f-465a-82f4
/* Gets the total Network device interface counts. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-network-device-interface-counts-in-the-specified-time-range-when-there-is-no-start-and-end-time-specified-returns-the-latest-interfaces-total-count
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams *GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/count"

	queryString, _ := query.Values(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount(GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNetworkDeviceInterfaceCountsInTheSpecifiedTimeRangeWhenThereIsNoStartAndEndTimeSpecifiedReturnsTheLatestInterfacesTotalCount)
	return result, response, err

}

//GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData Get the interface data for the given interface id (instance Uuid) along with the statistics and poe data - c08d-d95c-4c7b-8283
/* Returns the interface data for the given interface instance Uuid along with the statistics data. The latest interface data in the specified start and end time range will be returned. When there is no start and end time specified returns the latest available data for the given interface Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param id id path parameter. The interface Uuid

@param GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-interface-data-for-the-given-interface-idinstance-uuid-along-with-the-statistics-and-poe-data
*/
func (s *DevicesService) GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData(id string, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataQueryParams *GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeDataQueryParams) (*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData(id, GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeDataQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheInterfaceDataForTheGivenInterfaceIdinstanceUuidAlongWithTheStatisticsAndPoeData")
	}

	result := response.Result().(*ResponseDevicesGetTheInterfaceDataForTheGivenInterfaceIDinstanceUUIDAlongWithTheStatisticsAndPoeData)
	return result, response, err

}

//GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters Gets the Network Device details based on the provided query parameters. - c8b4-f894-4c3a-932f
/* Gets the Network Device details based on the provided query parameters.  When there is no start and end time specified returns the latest device details. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-network-device-details-based-on-the-provided-query-parameters
*/
func (s *DevicesService) GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams *GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams) (*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices"

	queryString, _ := query.Values(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters(GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters")
	}

	result := response.Result().(*ResponseDevicesGetsTheNetworkDeviceDetailsBasedOnTheProvidedQueryParameters)
	return result, response, err

}

//GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters Gets the total Network device counts based on the provided query parameters. - f0a6-e96b-44fb-a549
/* Gets the total Network device counts. When there is no start and end time specified returns the latest interfaces total count. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-network-device-counts-based-on-the-provided-query-parameters
*/
func (s *DevicesService) GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams *GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams) (*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/count"

	queryString, _ := query.Values(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters(GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParametersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNetworkDeviceCountsBasedOnTheProvidedQueryParameters)
	return result, response, err

}

//GetTheDeviceDataForTheGivenDeviceIDUUID Get the device data for the given device id (Uuid) - 5a93-1957-475b-95b3
/* Returns the device data for the given device Uuid in the specified start and end time range. When there is no start and end time specified returns the latest available data for the given Id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param id id path parameter. The device Uuid

@param GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-device-data-for-the-given-device-id-uuid
*/
func (s *DevicesService) GetTheDeviceDataForTheGivenDeviceIDUUID(id string, GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams *GetTheDeviceDataForTheGivenDeviceIDUUIDQueryParams) (*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDeviceDataForTheGivenDeviceIDUUID(id, GetTheDeviceDataForTheGivenDeviceIdUuidQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDeviceDataForTheGivenDeviceIdUuid")
	}

	result := response.Result().(*ResponseDevicesGetTheDeviceDataForTheGivenDeviceIDUUID)
	return result, response, err

}

//GetPlannedAccessPointsForBuilding Get Planned Access Points for Building - b699-9b85-4e3b-acdd
/* Provides a list of Planned Access Points for the Building it is requested for


@param buildingID buildingId path parameter. The instance UUID of the building hierarchy element

@param GetPlannedAccessPointsForBuildingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-for-building
*/
func (s *DevicesService) GetPlannedAccessPointsForBuilding(buildingID string, GetPlannedAccessPointsForBuildingQueryParams *GetPlannedAccessPointsForBuildingQueryParams) (*ResponseDevicesGetPlannedAccessPointsForBuilding, *resty.Response, error) {
	path := "/dna/intent/api/v1/buildings/{buildingId}/planned-access-points"
	path = strings.Replace(path, "{buildingId}", fmt.Sprintf("%v", buildingID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForBuildingQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForBuilding{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsForBuilding(buildingID, GetPlannedAccessPointsForBuildingQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForBuilding")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForBuilding)
	return result, response, err

}

//GetDeviceDetail Get Device Detail - ca98-fac4-4b08-895c
/* Returns detailed Network Device information retrieved by Mac Address, Device Name or UUID for any given point of time.


@param GetDeviceDetailQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-detail
*/
func (s *DevicesService) GetDeviceDetail(GetDeviceDetailQueryParams *GetDeviceDetailQueryParams) (*ResponseDevicesGetDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-detail"

	queryString, _ := query.Values(GetDeviceDetailQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDetail(GetDeviceDetailQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDetail")
	}

	result := response.Result().(*ResponseDevicesGetDeviceDetail)
	return result, response, err

}

//GetDeviceEnrichmentDetails Get Device Enrichment Details - e0b5-599b-4f29-97b7
/* Enriches a given network device context (device id or device Mac Address or device management IP address) with details about the device and neighbor topology


@param GetDeviceEnrichmentDetailsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-enrichment-details
*/
func (s *DevicesService) GetDeviceEnrichmentDetails(GetDeviceEnrichmentDetailsHeaderParams *GetDeviceEnrichmentDetailsHeaderParams) (*ResponseDevicesGetDeviceEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetDeviceEnrichmentDetailsHeaderParams != nil {

		if GetDeviceEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetDeviceEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetDeviceEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetDeviceEnrichmentDetailsHeaderParams.EntityValue)
		}

		if GetDeviceEnrichmentDetailsHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetDeviceEnrichmentDetailsHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetDeviceEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceEnrichmentDetails(GetDeviceEnrichmentDetailsHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceEnrichmentDetails")
	}

	result := response.Result().(*ResponseDevicesGetDeviceEnrichmentDetails)
	return result, response, err

}

//Devices Devices - 3ab2-bb64-4cca-81ee
/* Intent API for accessing DNA Assurance Device object for generating reports, creating dashboards or creating additional value added services.


@param DevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!devices
*/
func (s *DevicesService) Devices(DevicesQueryParams *DevicesQueryParams) (*ResponseDevicesDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-health"

	queryString, _ := query.Values(DevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Devices(DevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Devices")
	}

	result := response.Result().(*ResponseDevicesDevices)
	return result, response, err

}

//GetPlannedAccessPointsForFloor Get Planned Access Points for Floor - 6780-6977-4589-9a54
/* Provides a list of Planned Access Points for the Floor it is requested for


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

@param GetPlannedAccessPointsForFloorQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-for-floor
*/
func (s *DevicesService) GetPlannedAccessPointsForFloor(floorID string, GetPlannedAccessPointsForFloorQueryParams *GetPlannedAccessPointsForFloorQueryParams) (*ResponseDevicesGetPlannedAccessPointsForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsForFloorQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetPlannedAccessPointsForFloor{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsForFloor(floorID, GetPlannedAccessPointsForFloorQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsForFloor")
	}

	result := response.Result().(*ResponseDevicesGetPlannedAccessPointsForFloor)
	return result, response, err

}

//GetAllHealthScoreDefinitionsForGivenFilters Get all health score definitions for given filters. - 9bb6-ea87-4ffb-b492
/* Get all health score defintions.
Supported filters are id, name and overall health include status. A health score definition can be different across device type. So, deviceType in the query param is important and default is all device types.
By default all supported attributes are listed in response. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams Custom header parameters
@param GetAllHealthScoreDefinitionsForGivenFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-health-score-definitions-for-given-filters
*/
func (s *DevicesService) GetAllHealthScoreDefinitionsForGivenFilters(GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams *GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersQueryParams *GetAllHealthScoreDefinitionsForGivenFiltersQueryParams) (*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions"

	queryString, _ := query.Values(GetAllHealthScoreDefinitionsForGivenFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams != nil {

		if GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllHealthScoreDefinitionsForGivenFilters(GetAllHealthScoreDefinitionsForGivenFiltersHeaderParams, GetAllHealthScoreDefinitionsForGivenFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllHealthScoreDefinitionsForGivenFilters")
	}

	result := response.Result().(*ResponseDevicesGetAllHealthScoreDefinitionsForGivenFilters)
	return result, response, err

}

//GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters Get the count of health score definitions based on provided filters. - 49aa-bb2c-46ca-b58a
/* Get the count of health score definitions based on provided filters. Supported filters are id, name and overall health include status. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams Custom header parameters
@param GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-count-of-health-score-definitions-based-on-provided-filters
*/
func (s *DevicesService) GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams *GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams) (*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/count"

	queryString, _ := query.Values(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams != nil {

		if GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters(GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersHeaderParams, GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFiltersQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters")
	}

	result := response.Result().(*ResponseDevicesGetTheCountOfHealthScoreDefinitionsBasedOnProvidedFilters)
	return result, response, err

}

//GetHealthScoreDefinitionForTheGivenID Get health score definition for the given id. - 99b5-d81a-4408-94c3
/* Get health score defintion for the given id. Definition includes all properties from HealthScoreDefinition schema by default. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Health score definition id.

@param GetHealthScoreDefinitionForTheGivenIdHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-health-score-definition-for-the-given-id
*/
func (s *DevicesService) GetHealthScoreDefinitionForTheGivenID(id string, GetHealthScoreDefinitionForTheGivenIdHeaderParams *GetHealthScoreDefinitionForTheGivenIDHeaderParams) (*ResponseDevicesGetHealthScoreDefinitionForTheGivenID, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetHealthScoreDefinitionForTheGivenIdHeaderParams != nil {

		if GetHealthScoreDefinitionForTheGivenIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetHealthScoreDefinitionForTheGivenIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseDevicesGetHealthScoreDefinitionForTheGivenID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetHealthScoreDefinitionForTheGivenID(id, GetHealthScoreDefinitionForTheGivenIdHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetHealthScoreDefinitionForTheGivenId")
	}

	result := response.Result().(*ResponseDevicesGetHealthScoreDefinitionForTheGivenID)
	return result, response, err

}

//GetAllInterfaces Get all interfaces - f594-7a4c-439a-8bf0
/* Returns all available interfaces. This endpoint can return a maximum of 500 interfaces


@param GetAllInterfacesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-interfaces
*/
func (s *DevicesService) GetAllInterfaces(GetAllInterfacesQueryParams *GetAllInterfacesQueryParams) (*ResponseDevicesGetAllInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface"

	queryString, _ := query.Values(GetAllInterfacesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllInterfaces(GetAllInterfacesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetAllInterfaces)
	return result, response, err

}

//GetDeviceInterfaceCountForMultipleDevices Get Device Interface Count for Multiple Devices - 3d92-3b18-4dc9-a4ca
/* Returns the count of interfaces for all devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-count-for-multiple-devices
*/
func (s *DevicesService) GetDeviceInterfaceCountForMultipleDevices() (*ResponseDevicesGetDeviceInterfaceCountForMultipleDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCountForMultipleDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceCountForMultipleDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCountForMultipleDevices")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCountForMultipleDevices)
	return result, response, err

}

//GetInterfaceByIP Get Interface by IP - cd84-69e6-47ca-ab0e
/* Returns list of interfaces for specified device management IP address


@param ipAddress ipAddress path parameter. IP address of the interface


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-ip
*/
func (s *DevicesService) GetInterfaceByIP(ipAddress string) (*ResponseDevicesGetInterfaceByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByIP(ipAddress)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceByIp")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByIP)
	return result, response, err

}

//GetIsisInterfaces Get ISIS interfaces - 84ad-8b0e-42ca-b48a
/* Returns the interfaces that has ISIS enabled



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-isis-interfaces
*/
func (s *DevicesService) GetIsisInterfaces() (*ResponseDevicesGetIsisInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/isis"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetIsisInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIsisInterfaces()
		}
		return nil, response, fmt.Errorf("error with operation GetIsisInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetIsisInterfaces)
	return result, response, err

}

//GetInterfaceInfoByID Get Interface info by Id - ba9d-c85b-4b8a-9a17
/* Returns list of interfaces by specified device


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-info-by-id
*/
func (s *DevicesService) GetInterfaceInfoByID(deviceID string) (*ResponseDevicesGetInterfaceInfoByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceInfoByID(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceInfoById")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceInfoByID)
	return result, response, err

}

//GetDeviceInterfaceCount Get Device Interface count - 5b86-3922-4cd8-8ea7
/* Returns the interface count for the given device


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-count
*/
func (s *DevicesService) GetDeviceInterfaceCount(deviceID string) (*ResponseDevicesGetDeviceInterfaceCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/count"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfaceCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceCount(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceCount")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceCount)
	return result, response, err

}

//GetInterfaceDetailsByDeviceIDAndInterfaceName Get Interface details by device Id and interface name - 4eb5-6a61-4cc9-a2d2
/* Returns interface by specified device Id and interface name


@param deviceID deviceId path parameter. Device ID

@param GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-details-by-device-id-and-interface-name
*/
func (s *DevicesService) GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID string, GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams *GetInterfaceDetailsByDeviceIDAndInterfaceNameQueryParams) (*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/interface-name"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceDetailsByDeviceIDAndInterfaceName(deviceID, GetInterfaceDetailsByDeviceIdAndInterfaceNameQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceDetailsByDeviceIdAndInterfaceName")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceDetailsByDeviceIDAndInterfaceName)
	return result, response, err

}

//GetDeviceInterfacesBySpecifiedRange Get Device Interfaces by specified range - 349c-8884-43b8-9a58
/* Returns the list of interfaces for the device for the specified range


@param deviceID deviceId path parameter. Device ID

@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interfaces-by-specified-range
*/
func (s *DevicesService) GetDeviceInterfacesBySpecifiedRange(deviceID string, startIndex int, recordsToReturn int) (*ResponseDevicesGetDeviceInterfacesBySpecifiedRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/network-device/{deviceId}/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceInterfacesBySpecifiedRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfacesBySpecifiedRange(deviceID, startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfacesBySpecifiedRange")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfacesBySpecifiedRange)
	return result, response, err

}

//GetOspfInterfaces Get OSPF interfaces - 70ad-3976-49e9-b4d3
/* Returns the interfaces that has OSPF enabled



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ospf-interfaces
*/
func (s *DevicesService) GetOspfInterfaces() (*ResponseDevicesGetOspfInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/ospf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOspfInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOspfInterfaces()
		}
		return nil, response, fmt.Errorf("error with operation GetOspfInterfaces")
	}

	result := response.Result().(*ResponseDevicesGetOspfInterfaces)
	return result, response, err

}

//GetInterfaceByID Get Interface by Id - b888-792d-43ba-ba46
/* Returns the interface for the given interface ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-id
*/
func (s *DevicesService) GetInterfaceByID(id string) (*ResponseDevicesGetInterfaceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetInterfaceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceById")
	}

	result := response.Result().(*ResponseDevicesGetInterfaceByID)
	return result, response, err

}

//LegitOperationsForInterface Legit operations for interface - 87a3-3a52-46ea-a40e
/* Get list of all properties & operations valid for an interface.


@param interfaceUUID interfaceUuid path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!legit-operations-for-interface
*/
func (s *DevicesService) LegitOperationsForInterface(interfaceUUID string) (*ResponseDevicesLegitOperationsForInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/legit-operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesLegitOperationsForInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.LegitOperationsForInterface(interfaceUUID)
		}
		return nil, response, fmt.Errorf("error with operation LegitOperationsForInterface")
	}

	result := response.Result().(*ResponseDevicesLegitOperationsForInterface)
	return result, response, err

}

//GetDeviceList Get Device list - 20b1-9b52-464b-8972
/* Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc. You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-device?hostname=myhost.*&managementIpAddress=192.25.18..*
If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and ignores the other request parameters. You can also specify offset & limit to get the required list.


@param GetDeviceListQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-list
*/
func (s *DevicesService) GetDeviceList(GetDeviceListQueryParams *GetDeviceListQueryParams) (*ResponseDevicesGetDeviceList, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	queryString, _ := query.Values(GetDeviceListQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceList(GetDeviceListQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceList")
	}

	result := response.Result().(*ResponseDevicesGetDeviceList)
	return result, response, err

}

//GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute Get Device Values that match fully or partially an Attribute - ffa7-48cc-44e9-a437
/* Returns the list of values of the first given required parameter. You can use the .* in any value to conduct a wildcard search. For example, to get all the devices with the management IP address starting with 10.10. , issue the following request: GET /dna/inten/api/v1/network-device/autocomplete?managementIpAddress=10.10..* It will return the device management IP addresses that match fully or partially the provided attribute. {[10.10.1.1, 10.10.20.2, …]}.


@param GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-values-that-match-fully-or-partially-an-attribute
*/
func (s *DevicesService) GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams *GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams) (*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/autocomplete"

	queryString, _ := query.Values(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute(GetDeviceValuesThatMatchFullyOrPartiallyAnAttributeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceValuesThatMatchFullyOrPartiallyAnAttribute")
	}

	result := response.Result().(*ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute)
	return result, response, err

}

//GetPollingIntervalForAllDevices Get Polling Interval for all devices - 38bd-0b88-4b89-a785
/* Returns polling interval of all devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-polling-interval-for-all-devices
*/
func (s *DevicesService) GetPollingIntervalForAllDevices() (*ResponseDevicesGetPollingIntervalForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/collection-schedule/global"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPollingIntervalForAllDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalForAllDevices")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalForAllDevices)
	return result, response, err

}

//GetDeviceConfigForAllDevices Get Device Config for all devices - b7bc-aa08-4e2b-90d0
/* Returns the config for all devices. This API has been deprecated and will not be available in a Cisco Catalyst Center release after Nov 1st 2024 23:59:59 GMT.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-for-all-devices
*/
func (s *DevicesService) GetDeviceConfigForAllDevices() (*ResponseDevicesGetDeviceConfigForAllDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigForAllDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigForAllDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigForAllDevices")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigForAllDevices)
	return result, response, err

}

//GetDeviceConfigCount Get Device Config Count - 888f-585c-49b8-8441
/* Returns the count of device configs



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-count
*/
func (s *DevicesService) GetDeviceConfigCount() (*ResponseDevicesGetDeviceConfigCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/config/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigCount()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigCount")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigCount)
	return result, response, err

}

//GetDeviceCountKnowYourNetwork Get Device Count - 5db2-1b8e-43fa-b7d8
/* Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and location name


@param GetDeviceCountKnowYourNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-count-know-your-network
*/
func (s *DevicesService) GetDeviceCountKnowYourNetwork(GetDeviceCountKnowYourNetworkQueryParams *GetDeviceCountKnowYourNetworkQueryParams) (*ResponseDevicesGetDeviceCountKnowYourNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/count"

	queryString, _ := query.Values(GetDeviceCountKnowYourNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceCountKnowYourNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCountKnowYourNetwork(GetDeviceCountKnowYourNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCountKnowYourNetwork")
	}

	result := response.Result().(*ResponseDevicesGetDeviceCountKnowYourNetwork)
	return result, response, err

}

//GetFunctionalCapabilityForDevices Get Functional Capability for devices - c3b3-c9ef-4e6b-8a09
/* Returns the functional-capability for given devices


@param GetFunctionalCapabilityForDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-functional-capability-for-devices
*/
func (s *DevicesService) GetFunctionalCapabilityForDevices(GetFunctionalCapabilityForDevicesQueryParams *GetFunctionalCapabilityForDevicesQueryParams) (*ResponseDevicesGetFunctionalCapabilityForDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability"

	queryString, _ := query.Values(GetFunctionalCapabilityForDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetFunctionalCapabilityForDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFunctionalCapabilityForDevices(GetFunctionalCapabilityForDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityForDevices")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityForDevices)
	return result, response, err

}

//GetFunctionalCapabilityByID Get Functional Capability by Id - 81bb-4804-405a-8d2f
/* Returns functional capability with given Id


@param id id path parameter. Functional Capability UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-functional-capability-by-id
*/
func (s *DevicesService) GetFunctionalCapabilityByID(id string) (*ResponseDevicesGetFunctionalCapabilityByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/functional-capability/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetFunctionalCapabilityByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFunctionalCapabilityByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetFunctionalCapabilityById")
	}

	result := response.Result().(*ResponseDevicesGetFunctionalCapabilityByID)
	return result, response, err

}

//InventoryInsightDeviceLinkMismatchAPI Inventory Insight Device Link Mismatch API - 5792-59d8-4208-8190
/* Find all devices with link mismatch (speed /  vlan)


@param siteID siteId path parameter.
@param InventoryInsightDeviceLinkMismatchAPIQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!inventory-insight-device-link-mismatch-api
*/
func (s *DevicesService) InventoryInsightDeviceLinkMismatchAPI(siteID string, InventoryInsightDeviceLinkMismatchAPIQueryParams *InventoryInsightDeviceLinkMismatchAPIQueryParams) (*ResponseDevicesInventoryInsightDeviceLinkMismatchAPI, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/insight/{siteId}/device-link"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(InventoryInsightDeviceLinkMismatchAPIQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesInventoryInsightDeviceLinkMismatchAPI{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.InventoryInsightDeviceLinkMismatchAPI(siteID, InventoryInsightDeviceLinkMismatchAPIQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation InventoryInsightDeviceLinkMismatchApi")
	}

	result := response.Result().(*ResponseDevicesInventoryInsightDeviceLinkMismatchAPI)
	return result, response, err

}

//GetNetworkDeviceByIP Get Network Device by IP - d0a4-b881-45aa-bb51
/* Returns the network device by specified IP address


@param ipAddress ipAddress path parameter. Device IP address


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-by-ip
*/
func (s *DevicesService) GetNetworkDeviceByIP(ipAddress string) (*ResponseDevicesGetNetworkDeviceByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/ip-address/{ipAddress}"
	path = strings.Replace(path, "{ipAddress}", fmt.Sprintf("%v", ipAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceByIP(ipAddress)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByIp")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByIP)
	return result, response, err

}

//GetModules Get Modules - eb82-49e3-4f69-b0f1
/* Returns modules by specified device id


@param GetModulesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-modules
*/
func (s *DevicesService) GetModules(GetModulesQueryParams *GetModulesQueryParams) (*ResponseDevicesGetModules, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module"

	queryString, _ := query.Values(GetModulesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModules(GetModulesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetModules")
	}

	result := response.Result().(*ResponseDevicesGetModules)
	return result, response, err

}

//GetModuleCount Get Module count - 8db9-3974-4649-a782
/* Returns Module Count


@param GetModuleCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-module-count
*/
func (s *DevicesService) GetModuleCount(GetModuleCountQueryParams *GetModuleCountQueryParams) (*ResponseDevicesGetModuleCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/count"

	queryString, _ := query.Values(GetModuleCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetModuleCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModuleCount(GetModuleCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetModuleCount")
	}

	result := response.Result().(*ResponseDevicesGetModuleCount)
	return result, response, err

}

//GetModuleInfoByID Get Module Info by Id - 0db7-da74-4c0b-83d8
/* Returns Module info by 'module id'


@param id id path parameter. Module id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-module-info-by-id
*/
func (s *DevicesService) GetModuleInfoByID(id string) (*ResponseDevicesGetModuleInfoByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/module/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetModuleInfoByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetModuleInfoByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetModuleInfoById")
	}

	result := response.Result().(*ResponseDevicesGetModuleInfoByID)
	return result, response, err

}

//GetDeviceBySerialNumber Get Device by Serial number - d888-ab6d-4d59-a8c1
/* Returns the network device with given serial number


@param serialNumber serialNumber path parameter. Device serial number


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-serial-number
*/
func (s *DevicesService) GetDeviceBySerialNumber(serialNumber string) (*ResponseDevicesGetDeviceBySerialNumber, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/serial-number/{serialNumber}"
	path = strings.Replace(path, "{serialNumber}", fmt.Sprintf("%v", serialNumber), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceBySerialNumber{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceBySerialNumber(serialNumber)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceBySerialNumber")
	}

	result := response.Result().(*ResponseDevicesGetDeviceBySerialNumber)
	return result, response, err

}

//GetDevicesRegisteredForWsaNotification Get Devices registered for WSA Notification - c980-9b67-44f8-a502
/* It fetches devices which are registered to receive WSA notifications. The device serial number and/or MAC address are required to be provided as query parameters.


@param GetDevicesRegisteredForWSANotificationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-devices-registered-for-wsa-notification
*/
func (s *DevicesService) GetDevicesRegisteredForWsaNotification(GetDevicesRegisteredForWSANotificationQueryParams *GetDevicesRegisteredForWsaNotificationQueryParams) (*ResponseDevicesGetDevicesRegisteredForWsaNotification, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/tenantinfo/macaddress"

	queryString, _ := query.Values(GetDevicesRegisteredForWSANotificationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDevicesRegisteredForWsaNotification{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDevicesRegisteredForWsaNotification(GetDevicesRegisteredForWSANotificationQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDevicesRegisteredForWsaNotification")
	}

	result := response.Result().(*ResponseDevicesGetDevicesRegisteredForWsaNotification)
	return result, response, err

}

//GetAllUserDefinedFields Get All User-Defined-Fields - 058d-2a92-4899-b7bb
/* Gets existing global User Defined Fields. If no input is given, it fetches ALL the Global UDFs. Filter/search is supported by UDF Id(s) or UDF name(s) or both.


@param GetAllUserDefinedFieldsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-user-defined-fields
*/
func (s *DevicesService) GetAllUserDefinedFields(GetAllUserDefinedFieldsQueryParams *GetAllUserDefinedFieldsQueryParams) (*ResponseDevicesGetAllUserDefinedFields, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field"

	queryString, _ := query.Values(GetAllUserDefinedFieldsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllUserDefinedFields{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllUserDefinedFields(GetAllUserDefinedFieldsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllUserDefinedFields")
	}

	result := response.Result().(*ResponseDevicesGetAllUserDefinedFields)
	return result, response, err

}

//GetChassisDetailsForDevice Get Chassis Details for Device - 0486-9b26-49ab-b579
/* Returns chassis details for given device ID


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-chassis-details-for-device
*/
func (s *DevicesService) GetChassisDetailsForDevice(deviceID string) (*ResponseDevicesGetChassisDetailsForDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/chassis"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetChassisDetailsForDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetChassisDetailsForDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetChassisDetailsForDevice")
	}

	result := response.Result().(*ResponseDevicesGetChassisDetailsForDevice)
	return result, response, err

}

//GetStackDetailsForDevice Get Stack Details for Device - 78a7-7ab0-4d5a-8a10
/* Retrieves complete stack details for given device ID


@param deviceID deviceId path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-stack-details-for-device
*/
func (s *DevicesService) GetStackDetailsForDevice(deviceID string) (*ResponseDevicesGetStackDetailsForDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/stack"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetStackDetailsForDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetStackDetailsForDevice(deviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetStackDetailsForDevice")
	}

	result := response.Result().(*ResponseDevicesGetStackDetailsForDevice)
	return result, response, err

}

//GetTheDetailsOfPhysicalComponentsOfTheGivenDevice Get the Details of Physical Components of the Given Device. - 20b1-9b52-464b-897a
/* Return all types of equipment details like PowerSupply, Fan, Chassis, Backplane, Module, PROCESSOR, Other and SFP for the Given device.


@param deviceUUID deviceUuid path parameter.
@param GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-details-of-physical-components-of-the-given-device
*/
func (s *DevicesService) GetTheDetailsOfPhysicalComponentsOfTheGivenDevice(deviceUUID string, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams *GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams) (*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/equipment"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheDetailsOfPhysicalComponentsOfTheGivenDevice(deviceUUID, GetTheDetailsOfPhysicalComponentsOfTheGivenDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTheDetailsOfPhysicalComponentsOfTheGivenDevice")
	}

	result := response.Result().(*ResponseDevicesGetTheDetailsOfPhysicalComponentsOfTheGivenDevice)
	return result, response, err

}

//ReturnsPoeInterfaceDetailsForTheDevice Returns POE interface details for the device. - 20b5-48af-42da-a337
/* Returns POE interface details for the device, where deviceuuid is mandatory & accepts comma seperated interface names which is optional and returns information for that particular interfaces where(operStatus = operationalStatus)


@param deviceUUID deviceUuid path parameter. uuid of the device

@param ReturnsPOEInterfaceDetailsForTheDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-poe-interface-details-for-the-device
*/
func (s *DevicesService) ReturnsPoeInterfaceDetailsForTheDevice(deviceUUID string, ReturnsPOEInterfaceDetailsForTheDeviceQueryParams *ReturnsPoeInterfaceDetailsForTheDeviceQueryParams) (*ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/poe-detail"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ReturnsPOEInterfaceDetailsForTheDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsPoeInterfaceDetailsForTheDevice(deviceUUID, ReturnsPOEInterfaceDetailsForTheDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsPoeInterfaceDetailsForTheDevice")
	}

	result := response.Result().(*ResponseDevicesReturnsPoeInterfaceDetailsForTheDevice)
	return result, response, err

}

//GetConnectedDeviceDetail Get connected device detail - a8aa-ca21-4c09-8388
/* Get connected device detail for given deviceUuid and interfaceUuid


@param deviceUUID deviceUuid path parameter. instanceuuid of Device

@param interfaceUUID interfaceUuid path parameter. instanceuuid of interface


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-connected-device-detail
*/
func (s *DevicesService) GetConnectedDeviceDetail(deviceUUID string, interfaceUUID string) (*ResponseDevicesGetConnectedDeviceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/interface/{interfaceUuid}/neighbor"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetConnectedDeviceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConnectedDeviceDetail(deviceUUID, interfaceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetConnectedDeviceDetail")
	}

	result := response.Result().(*ResponseDevicesGetConnectedDeviceDetail)
	return result, response, err

}

//GetLinecardDetails Get Linecard details - 46a1-4b02-48fb-8fbf
/* Get line card detail for a given deviceuuid.  Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-linecard-details
*/
func (s *DevicesService) GetLinecardDetails(deviceUUID string) (*ResponseDevicesGetLinecardDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/line-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetLinecardDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLinecardDetails(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetLinecardDetails")
	}

	result := response.Result().(*ResponseDevicesGetLinecardDetails)
	return result, response, err

}

//PoeDetails POE details  - 8ba6-7932-4ed9-abae
/* Returns POE details for device.


@param deviceUUID deviceUuid path parameter. UUID of the device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!poe-details
*/
func (s *DevicesService) PoeDetails(deviceUUID string) (*ResponseDevicesPoeDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/poe"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesPoeDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.PoeDetails(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation PoeDetails")
	}

	result := response.Result().(*ResponseDevicesPoeDetails)
	return result, response, err

}

//GetSupervisorCardDetail Get Supervisor card detail - 88aa-1b52-4a38-bf97
/* Get supervisor card detail for a given deviceuuid. Response will contain serial no, part no, switch no and slot no.


@param deviceUUID deviceUuid path parameter. instanceuuid of device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-supervisor-card-detail
*/
func (s *DevicesService) GetSupervisorCardDetail(deviceUUID string) (*ResponseDevicesGetSupervisorCardDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceUuid}/supervisor-card"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetSupervisorCardDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSupervisorCardDetail(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation GetSupervisorCardDetail")
	}

	result := response.Result().(*ResponseDevicesGetSupervisorCardDetail)
	return result, response, err

}

//GetDeviceByID Get Device by ID - 8fa8-eb40-4a4a-8d96
/* Returns the network device details for the given device ID


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-by-id
*/
func (s *DevicesService) GetDeviceByID(id string) (*ResponseDevicesGetDeviceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceById")
	}

	result := response.Result().(*ResponseDevicesGetDeviceByID)
	return result, response, err

}

//GetDeviceSummary Get Device Summary - 819f-9aa5-4fea-b7bf
/* Returns brief summary of device info for the given device Id


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-summary
*/
func (s *DevicesService) GetDeviceSummary(id string) (*ResponseDevicesGetDeviceSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/brief"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceSummary(id)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceSummary")
	}

	result := response.Result().(*ResponseDevicesGetDeviceSummary)
	return result, response, err

}

//GetPollingIntervalByID Get Polling Interval by Id - 8291-8a1b-4d28-9c5c
/* Returns polling interval by device id


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-polling-interval-by-id
*/
func (s *DevicesService) GetPollingIntervalByID(id string) (*ResponseDevicesGetPollingIntervalByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/collection-schedule"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetPollingIntervalByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPollingIntervalByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetPollingIntervalById")
	}

	result := response.Result().(*ResponseDevicesGetPollingIntervalByID)
	return result, response, err

}

//GetOrganizationListForMeraki Get Organization list for Meraki - 84b3-7ae5-4c59-ab28
/* Returns list of organizations for meraki dashboard


@param id id path parameter. Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-organization-list-for-meraki
*/
func (s *DevicesService) GetOrganizationListForMeraki(id string) (*ResponseDevicesGetOrganizationListForMeraki, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/meraki-organization"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetOrganizationListForMeraki{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOrganizationListForMeraki(id)
		}
		return nil, response, fmt.Errorf("error with operation GetOrganizationListForMeraki")
	}

	result := response.Result().(*ResponseDevicesGetOrganizationListForMeraki)
	return result, response, err

}

//GetDeviceInterfaceVLANs Get Device Interface VLANs - 288d-f949-4f2a-9746
/* Returns Device Interface VLANs. If parameter value is null or empty, it won't return any value in response.


@param id id path parameter.
@param GetDeviceInterfaceVLANsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-vlans
*/
func (s *DevicesService) GetDeviceInterfaceVLANs(id string, GetDeviceInterfaceVLANsQueryParams *GetDeviceInterfaceVLANsQueryParams) (*ResponseDevicesGetDeviceInterfaceVLANs, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/vlan"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDeviceInterfaceVLANsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDeviceInterfaceVLANs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceVLANs(id, GetDeviceInterfaceVLANsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceVlans")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceVLANs)
	return result, response, err

}

//GetWirelessLanControllerDetailsByID Get wireless lan controller details by Id - f682-6a8e-41bb-a242
/* Returns the wireless lan controller info with given device ID


@param id id path parameter. Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-lan-controller-details-by-id
*/
func (s *DevicesService) GetWirelessLanControllerDetailsByID(id string) (*ResponseDevicesGetWirelessLanControllerDetailsByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{id}/wireless-info"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetWirelessLanControllerDetailsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessLanControllerDetailsByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessLanControllerDetailsById")
	}

	result := response.Result().(*ResponseDevicesGetWirelessLanControllerDetailsByID)
	return result, response, err

}

//GetDeviceConfigByID Get Device Config by Id - 84b3-3a9e-480a-bcaf
/* Returns the device config by specified device ID


@param networkDeviceID networkDeviceId path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-config-by-id
*/
func (s *DevicesService) GetDeviceConfigByID(networkDeviceID string) (*ResponseDevicesGetDeviceConfigByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{networkDeviceId}/config"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetDeviceConfigByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceConfigByID(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceConfigById")
	}

	result := response.Result().(*ResponseDevicesGetDeviceConfigByID)
	return result, response, err

}

//GetNetworkDeviceByPaginationRange Get Network Device by pagination range - f495-48c5-4be8-a3e2
/* Returns the list of network devices for the given pagination range. The maximum number of records that can be retrieved is 500


@param startIndex startIndex path parameter. Start index [>=1]

@param recordsToReturn recordsToReturn path parameter. Number of records to return [1<= recordsToReturn <= 500]


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-by-pagination-range
*/
func (s *DevicesService) GetNetworkDeviceByPaginationRange(startIndex int, recordsToReturn int) (*ResponseDevicesGetNetworkDeviceByPaginationRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetNetworkDeviceByPaginationRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceByPaginationRange(startIndex, recordsToReturn)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceByPaginationRange")
	}

	result := response.Result().(*ResponseDevicesGetNetworkDeviceByPaginationRange)
	return result, response, err

}

//RetrieveScheduledMaintenanceWindowsForNetworkDevices Retrieve scheduled maintenance windows for network devices - 7d9e-198c-4a9b-8971
/* This API retrieves a list of scheduled maintenance windows for network devices based on filter parameters. Each maintenance window is composed of a start schedule and end schedule, both of which have unique identifiers(`startId` and `endId`). These identifiers can be used to fetch the status of the start schedule and end schedule using the `GET /dna/intent/api/v1/activities/{id}` API. Completed maintenance schedules are automatically removed from the system after two weeks.


@param RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-scheduled-maintenance-windows-for-network-devices
*/
func (s *DevicesService) RetrieveScheduledMaintenanceWindowsForNetworkDevices(RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams *RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams) (*ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules"

	queryString, _ := query.Values(RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveScheduledMaintenanceWindowsForNetworkDevices(RetrieveScheduledMaintenanceWindowsForNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveScheduledMaintenanceWindowsForNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesRetrieveScheduledMaintenanceWindowsForNetworkDevices)
	return result, response, err

}

//RetrieveTheTotalNumberOfScheduledMaintenanceWindows Retrieve the total number of scheduled maintenance windows - 6981-69d1-44e8-9284
/* Retrieve the total count of all scheduled maintenance windows for network devices.


@param RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-total-number-of-scheduled-maintenance-windows
*/
func (s *DevicesService) RetrieveTheTotalNumberOfScheduledMaintenanceWindows(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams *RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams) (*ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindows, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/count"

	queryString, _ := query.Values(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindows{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheTotalNumberOfScheduledMaintenanceWindows(RetrieveTheTotalNumberOfScheduledMaintenanceWindowsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheTotalNumberOfScheduledMaintenanceWindows")
	}

	result := response.Result().(*ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindows)
	return result, response, err

}

//RetrievesTheMaintenanceScheduleInformation Retrieves the maintenance schedule information. - 5fb8-487f-4248-9f71
/* API to retrieve the maintenance schedule information for the given id.


@param id id path parameter. Unique identifier for the maintenance schedule


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-maintenance-schedule-information
*/
func (s *DevicesService) RetrievesTheMaintenanceScheduleInformation(id string) (*ResponseDevicesRetrievesTheMaintenanceScheduleInformation, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesRetrievesTheMaintenanceScheduleInformation{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheMaintenanceScheduleInformation(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheMaintenanceScheduleInformation")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheMaintenanceScheduleInformation)
	return result, response, err

}

//RetrieveNetworkDevices Retrieve network devices - 9e97-b8b2-4539-86a3
/* API to fetch the list of network devices using basic filters. Use the `/dna/intent/api/v1/networkDevices/query` API for advanced filtering. Refer features for more details.


@param RetrieveNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-network-devices
*/
func (s *DevicesService) RetrieveNetworkDevices(RetrieveNetworkDevicesQueryParams *RetrieveNetworkDevicesQueryParams) (*ResponseDevicesRetrieveNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices"

	queryString, _ := query.Values(RetrieveNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRetrieveNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveNetworkDevices(RetrieveNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesRetrieveNetworkDevices)
	return result, response, err

}

//CountTheNumberOfNetworkDevices Count the number of network devices - b988-e961-493b-8e72
/* API to fetch the count of network devices using basic filters. Use the `/dna/intent/api/v1/networkDevices/query/count` API if you need advanced filtering.


@param CountTheNumberOfNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-network-devices
*/
func (s *DevicesService) CountTheNumberOfNetworkDevices(CountTheNumberOfNetworkDevicesQueryParams *CountTheNumberOfNetworkDevicesQueryParams) (*ResponseDevicesCountTheNumberOfNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/count"

	queryString, _ := query.Values(CountTheNumberOfNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesCountTheNumberOfNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfNetworkDevices(CountTheNumberOfNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountTheNumberOfNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfNetworkDevices)
	return result, response, err

}

//GetDetailsOfASingleNetworkDevice Get details of a single network device - 289a-88ec-4e49-a477
/* API to fetch the details of network device using the `id`. Use the `/dna/intent/api/v1/networkDevices/query` API for advanced filtering. The API supports views to fetch only the required fields. Refer features for more details.


@param id id path parameter. Unique identifier for the network device

@param GetDetailsOfASingleNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-network-device
*/
func (s *DevicesService) GetDetailsOfASingleNetworkDevice(id string, GetDetailsOfASingleNetworkDeviceQueryParams *GetDetailsOfASingleNetworkDeviceQueryParams) (*ResponseDevicesGetDetailsOfASingleNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDetailsOfASingleNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetDetailsOfASingleNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleNetworkDevice(id, GetDetailsOfASingleNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleNetworkDevice")
	}

	result := response.Result().(*ResponseDevicesGetDetailsOfASingleNetworkDevice)
	return result, response, err

}

//GetResyncIntervalForTheNetworkDevice Get resync interval for the network device - 4783-7a87-4aea-91e6
/* Fetch the reysnc interval for the given network device id.


@param id id path parameter. The id of the network device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-resync-interval-for-the-network-device
*/
func (s *DevicesService) GetResyncIntervalForTheNetworkDevice(id string) (*ResponseDevicesGetResyncIntervalForTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/resyncIntervalSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetResyncIntervalForTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetResyncIntervalForTheNetworkDevice(id)
		}
		return nil, response, fmt.Errorf("error with operation GetResyncIntervalForTheNetworkDevice")
	}

	result := response.Result().(*ResponseDevicesGetResyncIntervalForTheNetworkDevice)
	return result, response, err

}

//WirelessRogueApContainmentStatus Wireless Rogue AP Containment Status - a1ab-f9ae-4c38-9286
/* Intent API to check the wireless rogue access point containment status. The response includes all the details like containment status, contained by WLC, containment status of each BSSID etc. This API also includes the information of strongest detecting WLC for this rogue access point.


@param macAddress macAddress path parameter. MAC Address of the Wireless Rogue AP


Documentation Link: https://developer.cisco.com/docs/dna-center/#!wireless-rogue-ap-containment-status
*/
func (s *DevicesService) WirelessRogueApContainmentStatus(macAddress string) (*ResponseDevicesWirelessRogueApContainmentStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/wireless-containment/status/{macAddress}"
	path = strings.Replace(path, "{macAddress}", fmt.Sprintf("%v", macAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesWirelessRogueApContainmentStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.WirelessRogueApContainmentStatus(macAddress)
		}
		return nil, response, fmt.Errorf("error with operation WirelessRogueApContainmentStatus")
	}

	result := response.Result().(*ResponseDevicesWirelessRogueApContainmentStatus)
	return result, response, err

}

//GetThreatLevels Get Threat Levels - 64ba-bad4-4aa9-b493
/* Intent API to fetch all threat levels defined.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-threat-levels
*/
func (s *DevicesService) GetThreatLevels() (*ResponseDevicesGetThreatLevels, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/level"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetThreatLevels{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetThreatLevels()
		}
		return nil, response, fmt.Errorf("error with operation GetThreatLevels")
	}

	result := response.Result().(*ResponseDevicesGetThreatLevels)
	return result, response, err

}

//GetAllowedMacAddress Get Allowed Mac Address - 18ae-3ab0-447a-872f
/* Intent API to fetch all the allowed mac addresses in the system.


@param GetAllowedMacAddressQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-allowed-mac-address
*/
func (s *DevicesService) GetAllowedMacAddress(GetAllowedMacAddressQueryParams *GetAllowedMacAddressQueryParams) (*ResponseDevicesGetAllowedMacAddress, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list"

	queryString, _ := query.Values(GetAllowedMacAddressQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesGetAllowedMacAddress{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllowedMacAddress(GetAllowedMacAddressQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllowedMacAddress")
	}

	result := response.Result().(*ResponseDevicesGetAllowedMacAddress)
	return result, response, err

}

//GetAllowedMacAddressCount Get Allowed Mac Address Count - d4a1-e8c8-410a-b009
/* Intent API to fetch the count of allowed mac addresses in the system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-allowed-mac-address-count
*/
func (s *DevicesService) GetAllowedMacAddressCount() (*ResponseDevicesGetAllowedMacAddressCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetAllowedMacAddressCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllowedMacAddressCount()
		}
		return nil, response, fmt.Errorf("error with operation GetAllowedMacAddressCount")
	}

	result := response.Result().(*ResponseDevicesGetAllowedMacAddressCount)
	return result, response, err

}

//GetThreatTypes Get Threat Types - 519a-9b70-45c8-8b82
/* Intent API to fetch all threat types defined.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-threat-types
*/
func (s *DevicesService) GetThreatTypes() (*ResponseDevicesGetThreatTypes, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/type"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesGetThreatTypes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetThreatTypes()
		}
		return nil, response, fmt.Errorf("error with operation GetThreatTypes")
	}

	result := response.Result().(*ResponseDevicesGetThreatTypes)
	return result, response, err

}

//RetrievesTheListOfAAAServicesForGivenSetOfComplexFilters Retrieves the list of AAA Services for given set of complex filters. - 55b1-ab25-40f8-ac6b
/* Retrieves the list of AAA Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheListOfAAAServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters *RequestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters, RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams *RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfAAAServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters, RetrievesTheListOfAAAServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfAAAServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfAAAServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters Retrieves the total number of AAA Services for given set of complex filters. - f894-482b-4d28-852f
/* Retrieves the total number of AAA Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters *RequestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters, RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams *RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters, RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfAAAServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters Get summary analytics data of AAA Services for given set of complex filters. - 0aaa-aa9b-4d39-90ec
/* Gets the summary analytics data related to AAA Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters *RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams *GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters Get Top N analytics data of AAA Services for given set of complex filters. - 69ad-ebf6-4ffa-89d4
/* Gets the Top N analytics data related to AAA Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters *RequestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams *GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetTopNAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters Get trend analytics data of AAA Services for given set of complex filters. - 00a7-4ab0-4b3b-8e31
/* Gets the trend analytics data related to AAA Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-a-a-a-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters *RequestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams *GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters, GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService Get trend analytics data for a given AAA Service matching the id of the Service. - f595-b9d5-413b-95fa
/* Gets the trend analytics data related to a particular AAA Service matching the id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AAAServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the AAA Service. It is the combination of AAA Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-for-a-given-a-a-a-service-matching-the-id-of-the-service
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService(id string, requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheService *RequestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService, GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceHeaderParams *GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheServiceHeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService, *resty.Response, error) {
	path := "/dna/data/api/v1/aaaServices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceHeaderParams != nil {

		if GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheService).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService(id, requestDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheService, GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheServiceHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIdOfTheService")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataForAGivenAAAServiceMatchingTheIDOfTheService)
	return result, response, err

}

//QueryAssuranceEventsWithFilters Query assurance events with filters - c5b7-0a69-4409-9b5b
/* Returns the list of events discovered by Catalyst Center, determined by the complex filters. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param QueryAssuranceEventsWithFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-assurance-events-with-filters
*/
func (s *DevicesService) QueryAssuranceEventsWithFilters(requestDevicesQueryAssuranceEventsWithFilters *RequestDevicesQueryAssuranceEventsWithFilters, QueryAssuranceEventsWithFiltersHeaderParams *QueryAssuranceEventsWithFiltersHeaderParams) (*ResponseDevicesQueryAssuranceEventsWithFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryAssuranceEventsWithFiltersHeaderParams != nil {

		if QueryAssuranceEventsWithFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryAssuranceEventsWithFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesQueryAssuranceEventsWithFilters).
		SetResult(&ResponseDevicesQueryAssuranceEventsWithFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryAssuranceEventsWithFilters(requestDevicesQueryAssuranceEventsWithFilters, QueryAssuranceEventsWithFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation QueryAssuranceEventsWithFilters")
	}

	result := response.Result().(*ResponseDevicesQueryAssuranceEventsWithFilters)
	return result, response, err

}

//CountTheNumberOfEventsWithFilters Count the number of events with filters - d685-3aeb-4878-a8fd
/* API to fetch the count of assurance events for the given complex query. Please refer to the 'API Support Documentation' section to understand which fields are supported. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceEvents-1.0.0-resolved.yaml


@param CountTheNumberOfEventsWithFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-events-with-filters
*/
func (s *DevicesService) CountTheNumberOfEventsWithFilters(requestDevicesCountTheNumberOfEventsWithFilters *RequestDevicesCountTheNumberOfEventsWithFilters, CountTheNumberOfEventsWithFiltersHeaderParams *CountTheNumberOfEventsWithFiltersHeaderParams) (*ResponseDevicesCountTheNumberOfEventsWithFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceEvents/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountTheNumberOfEventsWithFiltersHeaderParams != nil {

		if CountTheNumberOfEventsWithFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountTheNumberOfEventsWithFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesCountTheNumberOfEventsWithFilters).
		SetResult(&ResponseDevicesCountTheNumberOfEventsWithFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfEventsWithFilters(requestDevicesCountTheNumberOfEventsWithFilters, CountTheNumberOfEventsWithFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CountTheNumberOfEventsWithFilters")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfEventsWithFilters)
	return result, response, err

}

//RetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters Retrieves the list of DHCP Services for given set of complex filters. - 9b95-faaf-467b-8b13
/* Retrieves the list of DHCP Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters *RequestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters, RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams *RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters, RetrievesTheListOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDHCPServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters Retrieves the total number of DHCP Services for given set of complex filters. - 4b94-993c-459b-96a2
/* Retrieves the total number of DHCP Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters *RequestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters, RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams *RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters, RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDHCPServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters Get summary analytics data of DHCP Services for given set of complex filters. - 84b8-3b7d-405a-95e3
/* Gets the summary analytics data related to DHCP Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters *RequestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams *GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetSummaryAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters Get Top N analytics data of DHCP Services for given set of complex filters. - c2b2-0a95-46ca-a4f1
/* Gets the Top N analytics data related to DHCP Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters *RequestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams *GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetTopNAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters Get trend analytics data of DHCP Services for given set of complex filters. - 73ad-bb60-447b-8f79
/* Gets the trend analytics data related to DHCP Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-d-h-c-p-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters *RequestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams *GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters, GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataOfDHCPServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService Get trend analytics data for a given DHCP Service matching the id of the Service. - c28b-d8c9-4e59-bf4a
/* Gets the trend analytics data related to a particular DHCP Service matching the id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DHCPServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DHCP Service. It is the combination of DHCP Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-for-a-given-d-h-c-p-service-matching-the-id-of-the-service
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService(id string, requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheService *RequestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService, GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceHeaderParams *GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheServiceHeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService, *resty.Response, error) {
	path := "/dna/data/api/v1/dhcpServices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceHeaderParams != nil {

		if GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheService).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService(id, requestDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheService, GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheServiceHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIdOfTheService")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataForAGivenDHCPServiceMatchingTheIDOfTheService)
	return result, response, err

}

//RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters Retrieves the list of DNS Services for given set of complex filters. - ccbf-bb30-4b29-b2c4
/* Retrieves the list of DNS Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters *RequestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters, RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams *RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters, RetrievesTheListOfDNSServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfDNSServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheListOfDNSServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters Retrieves the total number of DNS Services for given set of complex filters. - 8bb8-abac-469a-b511
/* Retrieves the total number of DNS Services and offers complex filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters *RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters, RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams *RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters(requestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters, RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters Get summary analytics data of DNS Services for given set of complex filters. - 6993-98d1-456a-a499
/* Gets the summary analytics data related to DNS Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-summary-analytics-data-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters *RequestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams *GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetSummaryAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters Get Top N analytics data of DNS Services for given set of complex filters. - 90a0-19ad-4668-88c8
/* Gets the Top N analytics data related to DNS Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-top-n-analytics-data-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters *RequestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams *GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetTopNAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters Get trend analytics data of DNS Services for given set of complex filters. - feb3-7b93-4b0b-a74d
/* Gets the trend analytics data related to DNS Services based on given filters and group by field. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-of-d-n-s-services-for-given-set-of-complex-filters
*/
func (s *DevicesService) GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters *RequestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams *GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams) (*ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams != nil {

		if GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters(requestDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters, GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataOfDNSServicesForGivenSetOfComplexFilters)
	return result, response, err

}

//GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService Get trend analytics data for a given DNS Service matching the id of the Service. - 6099-29a8-404a-bf56
/* Gets the trend analytics data related to a particular DNS Service matching the id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml


@param id id path parameter. Unique id of the DNS Service. It is the combination of DNS Server IP (`serverIp`) and Device UUID (`deviceId`) separated by underscore (`_`). Example: If `serverIp` is `10.76.81.33` and `deviceId` is `6bef213c-19ca-4170-8375-b694e251101c`, then the `id` would be `10.76.81.33_6bef213c-19ca-4170-8375-b694e251101c`

@param GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trend-analytics-data-for-a-given-d-n-s-service-matching-the-id-of-the-service
*/
func (s *DevicesService) GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService(id string, requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheService *RequestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService, GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceHeaderParams *GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheServiceHeaderParams) (*ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService, *resty.Response, error) {
	path := "/dna/data/api/v1/dnsServices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceHeaderParams != nil {

		if GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheService).
		SetResult(&ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService(id, requestDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheService, GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheServiceHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIdOfTheService")
	}

	result := response.Result().(*ResponseDevicesGetTrendAnalyticsDataForAGivenDNSServiceMatchingTheIDOfTheService)
	return result, response, err

}

//GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions - 45b8-ba96-4daa-843c
/* Gets the list of interfaces across the Network Devices based on the provided complex filters and aggregation functions
The elements are grouped and sorted by deviceUuid first, and are then sorted by the given sort field, or by the default value: name.
The supported sorting options are: name, adminStatus, description, duplexConfig, duplexOper, interfaceIfIndex,interfaceType, macAddress,mediaType, operStatus, portChannelId, portMode, portType,speed, vlanId,pdPowerAdminMaxInWatt,pdPowerBudgetInWatt,pdPowerConsumedInWatt,pdPowerRemainingInWatt,pdMaxPowerDrawn. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-list-of-interfaces-across-the-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions *RequestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions) (*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions).
		SetResult(&ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions")
	}

	result := response.Result().(*ResponseDevicesGetsTheListOfInterfacesAcrossTheNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
	return result, response, err

}

//TheTotalInterfacesCountAcrossTheNetworkDevices The Total interfaces count across the Network devices. - a0bb-1bed-4529-98b1
/* Gets the total number of interfaces across the Network devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-total-interfaces-count-across-the-network-devices
*/
func (s *DevicesService) TheTotalInterfacesCountAcrossTheNetworkDevices(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices *RequestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices) (*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices).
		SetResult(&ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTotalInterfacesCountAcrossTheNetworkDevices(requestDevicesTheTotalInterfacesCountAcrossTheNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation TheTotalInterfacesCountAcrossTheNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesTheTotalInterfacesCountAcrossTheNetworkDevices)
	return result, response, err

}

//TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange The Trend analytcis data for the interfaces in the specified time range - ed96-48ef-4a98-a0a7
/* The Trend analytcis data for the interface, identified by its instanceUuid, in the specified time range. The data is grouped based on the trend time Interval, other input parameters like attributes and aggregate attributes. The default time interval range is 3 hours when start and endTime is not provided.
The field trendIntervalInMinutes is requiered and either the attributes or the aggregateAttributes fields is also required.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-interfaces-2.0.0-resolved.yaml


@param id id path parameter. The interface instance Uuid


Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytcis-data-for-the-interfaces-in-the-specified-time-range
*/
func (s *DevicesService) TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange(id string, requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange *RequestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange) (*ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/interfaces/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange).
		SetResult(&ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange(id, requestDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange)
		}

		return nil, response, fmt.Errorf("error with operation TheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseDevicesTheTrendAnalytcisDataForTheInterfacesInTheSpecifiedTimeRange)
	return result, response, err

}

//GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions Gets the list of Network Devices based on the provided complex filters and aggregation functions. - e794-1a90-428b-b583
/* Gets the list of Network Devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-list-of-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions *RequestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions) (*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions).
		SetResult(&ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions")
	}

	result := response.Result().(*ResponseDevicesGetsTheListOfNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
	return result, response, err

}

//GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions Gets the total number Network Devices based on the provided complex filters and aggregation functions. - 278f-1a5c-40ab-b65a
/* Gets the total number Network Devices based on the provided complex filters and aggregation functions. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-total-number-network-devices-based-on-the-provided-complex-filters-and-aggregation-functions
*/
func (s *DevicesService) GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions *RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions) (*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions).
		SetResult(&ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions(requestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions")
	}

	result := response.Result().(*ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctions)
	return result, response, err

}

//GetsTheSummaryAnalyticsDataRelatedToNetworkDevices Gets the summary analytics data related to network devices. - 15be-c9ed-4cba-8f91
/* Gets the summary analytics data related to network devices based on the provided input data. This endpoint helps to obtain the consolidated insights into the performance and status of the monitored network devices. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-summary-analytics-data-related-to-network-devices
*/
func (s *DevicesService) GetsTheSummaryAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices *RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices) (*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/summaryAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices).
		SetResult(&ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheSummaryAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheSummaryAnalyticsDataRelatedToNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevices)
	return result, response, err

}

//GetsTheTopNAnalyticsDataRelatedToNetworkDevices Gets the Top-N analytics data related to network devices. - 5b94-cae7-4acb-a8fe
/* Gets the Top N analytics data related to network devices based on the provided input data. This endpoint is valuable to obtain the top-performing or most impacted network devices. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml
The required properties for this API are topN and groupBy



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-top-n-analytics-data-related-to-network-devices
*/
func (s *DevicesService) GetsTheTopNAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices *RequestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices) (*ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/topNAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices).
		SetResult(&ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTopNAnalyticsDataRelatedToNetworkDevices(requestDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTopNAnalyticsDataRelatedToNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesGetsTheTopNAnalyticsDataRelatedToNetworkDevices)
	return result, response, err

}

//GetsTheTrendAnalyticsData Gets the Trend analytics data. - 0c93-595e-451b-910e
/* Gets the Trend analytics Network device data for the given time range. The data will be grouped based on the given trend time Interval. The required property for this API is `trendInterval`. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml



Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-trend-analytics-data
*/
func (s *DevicesService) GetsTheTrendAnalyticsData(requestDevicesGetsTheTrendAnalyticsData *RequestDevicesGetsTheTrendAnalyticsData) (*ResponseDevicesGetsTheTrendAnalyticsData, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/trendAnalytics"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetsTheTrendAnalyticsData).
		SetResult(&ResponseDevicesGetsTheTrendAnalyticsData{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTrendAnalyticsData(requestDevicesGetsTheTrendAnalyticsData)
		}

		return nil, response, fmt.Errorf("error with operation GetsTheTrendAnalyticsData")
	}

	result := response.Result().(*ResponseDevicesGetsTheTrendAnalyticsData)
	return result, response, err

}

//TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange The Trend analytics data for the network Device in the specified time range - 00ba-7a81-431b-93e7
/* The Trend analytics data for the network Device in the specified time range. The data is grouped based on the trend time Interval, other input parameters like attribute and aggregate attributes. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-2.0.1-resolved.yaml


@param id id path parameter. The device Uuid


Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-the-network-device-in-the-specified-time-range
*/
func (s *DevicesService) TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(id string, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange *RequestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange) (*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/networkDevices/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange).
		SetResult(&ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange(id, requestDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange)
		}

		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseDevicesTheTrendAnalyticsDataForTheNetworkDeviceInTheSpecifiedTimeRange)
	return result, response, err

}

//CreatePlannedAccessPointForFloor Create Planned Access Point for Floor - 7eaa-8b15-454a-8c1d
/* Allows creation of a new planned access point on an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch any existing planned access points for the floor.  The payload to create a planned access point is in the same format, albeit a single object instead of a list, of that API.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-planned-access-point-for-floor
*/
func (s *DevicesService) CreatePlannedAccessPointForFloor(floorID string, requestDevicesCreatePlannedAccessPointForFloor *RequestDevicesCreatePlannedAccessPointForFloor) (*ResponseDevicesCreatePlannedAccessPointForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreatePlannedAccessPointForFloor).
		SetResult(&ResponseDevicesCreatePlannedAccessPointForFloor{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatePlannedAccessPointForFloor(floorID, requestDevicesCreatePlannedAccessPointForFloor)
		}

		return nil, response, fmt.Errorf("error with operation CreatePlannedAccessPointForFloor")
	}

	result := response.Result().(*ResponseDevicesCreatePlannedAccessPointForFloor)
	return result, response, err

}

//UpdateHealthScoreDefinitions Update health score definitions. - 1aab-193d-40bb-9d2f
/* Update health thresholds, include status of overall health status for each metric.
And also to synchronize with global profile issue thresholds of the definition for given metric. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param UpdateHealthScoreDefinitionsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-health-score-definitions
*/
func (s *DevicesService) UpdateHealthScoreDefinitions(requestDevicesUpdateHealthScoreDefinitions *RequestDevicesUpdateHealthScoreDefinitions, UpdateHealthScoreDefinitionsHeaderParams *UpdateHealthScoreDefinitionsHeaderParams) (*ResponseDevicesUpdateHealthScoreDefinitions, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/bulkUpdate"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UpdateHealthScoreDefinitionsHeaderParams != nil {

		if UpdateHealthScoreDefinitionsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", UpdateHealthScoreDefinitionsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestDevicesUpdateHealthScoreDefinitions).
		SetResult(&ResponseDevicesUpdateHealthScoreDefinitions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHealthScoreDefinitions(requestDevicesUpdateHealthScoreDefinitions, UpdateHealthScoreDefinitionsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation UpdateHealthScoreDefinitions")
	}

	result := response.Result().(*ResponseDevicesUpdateHealthScoreDefinitions)
	return result, response, err

}

//ClearMacAddressTable Clear Mac-Address table - 24be-a97f-43f9-bc65
/* Clear mac-address on an individual port. In request body, operation needs to be specified as 'ClearMacAddress'. In the future more possible operations will be added to this API


@param interfaceUUID interfaceUuid path parameter. Interface Id

@param ClearMacAddressTableQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!clear-mac-address-table
*/
func (s *DevicesService) ClearMacAddressTable(interfaceUUID string, requestDevicesClearMacAddressTable *RequestDevicesClearMacAddressTable, ClearMacAddressTableQueryParams *ClearMacAddressTableQueryParams) (*ResponseDevicesClearMacAddressTable, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}/operation"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(ClearMacAddressTableQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesClearMacAddressTable).
		SetResult(&ResponseDevicesClearMacAddressTable{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClearMacAddressTable(interfaceUUID, requestDevicesClearMacAddressTable, ClearMacAddressTableQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ClearMacAddressTable")
	}

	result := response.Result().(*ResponseDevicesClearMacAddressTable)
	return result, response, err

}

//AddDeviceKnowYourNetwork Add Device - 4bb2-2af0-46fa-8f08
/* Adds the device with given credential



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-device-know-your-network
*/
func (s *DevicesService) AddDeviceKnowYourNetwork(requestDevicesAddDeviceKnowYourNetwork *RequestDevicesAddDeviceKnowYourNetwork) (*ResponseDevicesAddDeviceKnowYourNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddDeviceKnowYourNetwork).
		SetResult(&ResponseDevicesAddDeviceKnowYourNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddDeviceKnowYourNetwork(requestDevicesAddDeviceKnowYourNetwork)
		}

		return nil, response, fmt.Errorf("error with operation AddDeviceKnowYourNetwork")
	}

	result := response.Result().(*ResponseDevicesAddDeviceKnowYourNetwork)
	return result, response, err

}

//ExportDeviceList Export Device list - cd98-780f-4888-a66d
/* Exports the selected network device to a file



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-list
*/
func (s *DevicesService) ExportDeviceList(requestDevicesExportDeviceList *RequestDevicesExportDeviceList) (*ResponseDevicesExportDeviceList, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/file"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesExportDeviceList).
		SetResult(&ResponseDevicesExportDeviceList{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportDeviceList(requestDevicesExportDeviceList)
		}

		return nil, response, fmt.Errorf("error with operation ExportDeviceList")
	}

	result := response.Result().(*ResponseDevicesExportDeviceList)
	return result, response, err

}

//CreateUserDefinedField Create User-Defined-Field - 0a9c-18e7-4caa-8b07
/* Creates a new global User Defined Field, which can be assigned to devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-user-defined-field
*/
func (s *DevicesService) CreateUserDefinedField(requestDevicesCreateUserDefinedField *RequestDevicesCreateUserDefinedField) (*ResponseDevicesCreateUserDefinedField, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateUserDefinedField).
		SetResult(&ResponseDevicesCreateUserDefinedField{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateUserDefinedField(requestDevicesCreateUserDefinedField)
		}

		return nil, response, fmt.Errorf("error with operation CreateUserDefinedField")
	}

	result := response.Result().(*ResponseDevicesCreateUserDefinedField)
	return result, response, err

}

//CreateMaintenanceScheduleForNetworkDevices Create maintenance schedule for network devices - 0a8a-6859-4dca-b635
/* API to create maintenance schedule for network devices. The state of network device can be queried using API `GET /dna/intent/api/v1/networkDevices`. The `managementState` attribute of the network device will be updated to `UNDER_MAINTENANCE` when the maintenance window starts.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-maintenance-schedule-for-network-devices
*/
func (s *DevicesService) CreateMaintenanceScheduleForNetworkDevices(requestDevicesCreateMaintenanceScheduleForNetworkDevices *RequestDevicesCreateMaintenanceScheduleForNetworkDevices) (*ResponseDevicesCreateMaintenanceScheduleForNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCreateMaintenanceScheduleForNetworkDevices).
		SetResult(&ResponseDevicesCreateMaintenanceScheduleForNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateMaintenanceScheduleForNetworkDevices(requestDevicesCreateMaintenanceScheduleForNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation CreateMaintenanceScheduleForNetworkDevices")
	}

	result := response.Result().(*ResponseDevicesCreateMaintenanceScheduleForNetworkDevices)
	return result, response, err

}

//DeleteNetworkDeviceWithConfigurationCleanup Delete network device with configuration cleanup - 5080-5a46-4bbb-8ec0
/* This API endpoint facilitates the deletion of a network device after performing configuration cleanup on the device.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-network-device-with-configuration-cleanup
*/
func (s *DevicesService) DeleteNetworkDeviceWithConfigurationCleanup(requestDevicesDeleteNetworkDeviceWithConfigurationCleanup *RequestDevicesDeleteNetworkDeviceWithConfigurationCleanup) (*ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanup, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deleteWithCleanup"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesDeleteNetworkDeviceWithConfigurationCleanup).
		SetResult(&ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteNetworkDeviceWithConfigurationCleanup(requestDevicesDeleteNetworkDeviceWithConfigurationCleanup)
		}

		return nil, response, fmt.Errorf("error with operation DeleteNetworkDeviceWithConfigurationCleanup")
	}

	result := response.Result().(*ResponseDevicesDeleteNetworkDeviceWithConfigurationCleanup)
	return result, response, err

}

//DeleteANetworkDeviceWithoutConfigurationCleanup Delete a network device without configuration cleanup - 9a88-7984-4ba8-ab1c
/* This API endpoint facilitates the deletion of a network device without performing configuration cleanup on the device.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-network-device-without-configuration-cleanup
*/
func (s *DevicesService) DeleteANetworkDeviceWithoutConfigurationCleanup(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanup *RequestDevicesDeleteANetworkDeviceWithoutConfigurationCleanup) (*ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanup, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deleteWithoutCleanup"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanup).
		SetResult(&ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteANetworkDeviceWithoutConfigurationCleanup(requestDevicesDeleteANetworkDeviceWithoutConfigurationCleanup)
		}

		return nil, response, fmt.Errorf("error with operation DeleteANetworkDeviceWithoutConfigurationCleanup")
	}

	result := response.Result().(*ResponseDevicesDeleteANetworkDeviceWithoutConfigurationCleanup)
	return result, response, err

}

//QueryNetworkDevicesWithFilters Query network devices with filters - 0caa-2abb-490a-8e69
/* Returns the list of network devices, determined by the filters. It is possible to filter the network devices based on various parameters, such as device type, device role, software version, etc.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-network-devices-with-filters
*/
func (s *DevicesService) QueryNetworkDevicesWithFilters(requestDevicesQueryNetworkDevicesWithFilters *RequestDevicesQueryNetworkDevicesWithFilters) (*ResponseDevicesQueryNetworkDevicesWithFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesQueryNetworkDevicesWithFilters).
		SetResult(&ResponseDevicesQueryNetworkDevicesWithFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryNetworkDevicesWithFilters(requestDevicesQueryNetworkDevicesWithFilters)
		}

		return nil, response, fmt.Errorf("error with operation QueryNetworkDevicesWithFilters")
	}

	result := response.Result().(*ResponseDevicesQueryNetworkDevicesWithFilters)
	return result, response, err

}

//CountTheNumberOfNetworkDevicesWithFilters Count the number of network devices with filters - 83a1-08b8-471b-9820
/* API to fetch the count of network devices for the given filter query.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-the-number-of-network-devices-with-filters
*/
func (s *DevicesService) CountTheNumberOfNetworkDevicesWithFilters(requestDevicesCountTheNumberOfNetworkDevicesWithFilters *RequestDevicesCountTheNumberOfNetworkDevicesWithFilters) (*ResponseDevicesCountTheNumberOfNetworkDevicesWithFilters, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesCountTheNumberOfNetworkDevicesWithFilters).
		SetResult(&ResponseDevicesCountTheNumberOfNetworkDevicesWithFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountTheNumberOfNetworkDevicesWithFilters(requestDevicesCountTheNumberOfNetworkDevicesWithFilters)
		}

		return nil, response, fmt.Errorf("error with operation CountTheNumberOfNetworkDevicesWithFilters")
	}

	result := response.Result().(*ResponseDevicesCountTheNumberOfNetworkDevicesWithFilters)
	return result, response, err

}

//OverrideResyncInterval Override resync interval - 42ac-59bd-41db-a4fe
/* Overrides the global resync interval on all network devices. This essentially removes device specific intervals if set.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!override-resync-interval
*/
func (s *DevicesService) OverrideResyncInterval() (*ResponseDevicesOverrideResyncInterval, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/resyncIntervalSettings/override"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesOverrideResyncInterval{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.OverrideResyncInterval()
		}

		return nil, response, fmt.Errorf("error with operation OverrideResyncInterval")
	}

	result := response.Result().(*ResponseDevicesOverrideResyncInterval)
	return result, response, err

}

//RogueAdditionalDetails Rogue Additional Details - 659c-e9bd-403a-8de6
/* This API provides additional information of the rogue threats with details at BSSID level. The additional information includes Switch Port details in case of Rogue on Wire, first time when the rogue is seen in the network etc.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!rogue-additional-details
*/
func (s *DevicesService) RogueAdditionalDetails(requestDevicesRogueAdditionalDetails *RequestDevicesRogueAdditionalDetails) (*ResponseDevicesRogueAdditionalDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/additional/details"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRogueAdditionalDetails).
		SetResult(&ResponseDevicesRogueAdditionalDetails{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RogueAdditionalDetails(requestDevicesRogueAdditionalDetails)
		}

		return nil, response, fmt.Errorf("error with operation RogueAdditionalDetails")
	}

	result := response.Result().(*ResponseDevicesRogueAdditionalDetails)
	return result, response, err

}

//RogueAdditionalDetailCount Rogue Additional Detail Count - 4ca7-59be-4b99-9041
/* This API returns the count for the Rogue Additional Details.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!rogue-additional-detail-count
*/
func (s *DevicesService) RogueAdditionalDetailCount(requestDevicesRogueAdditionalDetailCount *RequestDevicesRogueAdditionalDetailCount) (*ResponseDevicesRogueAdditionalDetailCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/additional/details/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesRogueAdditionalDetailCount).
		SetResult(&ResponseDevicesRogueAdditionalDetailCount{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RogueAdditionalDetailCount(requestDevicesRogueAdditionalDetailCount)
		}

		return nil, response, fmt.Errorf("error with operation RogueAdditionalDetailCount")
	}

	result := response.Result().(*ResponseDevicesRogueAdditionalDetailCount)
	return result, response, err

}

//StartWirelessRogueApContainment Start Wireless Rogue AP Containment - 6998-5b93-4218-aea5
/* Intent API to start the wireless rogue access point containment. This API will initiate the containment operation on the strongest detecting WLC for the given Rogue AP. This is a resource intensive operation which has legal implications since the rogue access point on whom it is triggered, might be a valid neighbor access point.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!start-wireless-rogue-ap-containment
*/
func (s *DevicesService) StartWirelessRogueApContainment(requestDevicesStartWirelessRogueAPContainment *RequestDevicesStartWirelessRogueApContainment) (*ResponseDevicesStartWirelessRogueApContainment, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/wireless-containment/start"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesStartWirelessRogueAPContainment).
		SetResult(&ResponseDevicesStartWirelessRogueApContainment{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.StartWirelessRogueApContainment(requestDevicesStartWirelessRogueAPContainment)
		}

		return nil, response, fmt.Errorf("error with operation StartWirelessRogueApContainment")
	}

	result := response.Result().(*ResponseDevicesStartWirelessRogueApContainment)
	return result, response, err

}

//StopWirelessRogueApContainment Stop Wireless Rogue AP Containment - b692-6b1c-4d0a-b3fb
/* Intent API to stop the wireless rogue access point containment. This API will stop the containment through single WLC. The response includes the details like WLC and BSSID on which the stop containment has been initiated.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!stop-wireless-rogue-ap-containment
*/
func (s *DevicesService) StopWirelessRogueApContainment(requestDevicesStopWirelessRogueAPContainment *RequestDevicesStopWirelessRogueApContainment) (*ResponseDevicesStopWirelessRogueApContainment, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/rogue/wireless-containment/stop"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesStopWirelessRogueAPContainment).
		SetResult(&ResponseDevicesStopWirelessRogueApContainment{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.StopWirelessRogueApContainment(requestDevicesStopWirelessRogueAPContainment)
		}

		return nil, response, fmt.Errorf("error with operation StopWirelessRogueApContainment")
	}

	result := response.Result().(*ResponseDevicesStopWirelessRogueApContainment)
	return result, response, err

}

//ThreatDetails Threat Details - f6bf-c880-435a-ae2a
/* The details for the Rogue and aWIPS threats



Documentation Link: https://developer.cisco.com/docs/dna-center/#!threat-details
*/
func (s *DevicesService) ThreatDetails(requestDevicesThreatDetails *RequestDevicesThreatDetails) (*ResponseDevicesThreatDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/details"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesThreatDetails).
		SetResult(&ResponseDevicesThreatDetails{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ThreatDetails(requestDevicesThreatDetails)
		}

		return nil, response, fmt.Errorf("error with operation ThreatDetails")
	}

	result := response.Result().(*ResponseDevicesThreatDetails)
	return result, response, err

}

//ThreatDetailCount Threat Detail Count - eb8c-2a83-45aa-871f
/* The details count for the Rogue and aWIPS threats



Documentation Link: https://developer.cisco.com/docs/dna-center/#!threat-detail-count
*/
func (s *DevicesService) ThreatDetailCount(requestDevicesThreatDetailCount *RequestDevicesThreatDetailCount) (*ResponseDevicesThreatDetailCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/details/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesThreatDetailCount).
		SetResult(&ResponseDevicesThreatDetailCount{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ThreatDetailCount(requestDevicesThreatDetailCount)
		}

		return nil, response, fmt.Errorf("error with operation ThreatDetailCount")
	}

	result := response.Result().(*ResponseDevicesThreatDetailCount)
	return result, response, err

}

//AddAllowedMacAddress Add Allowed Mac Address - b6a0-887d-4fe9-9d5f
/* Intent API to add the threat mac address to allowed list.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-allowed-mac-address
*/
func (s *DevicesService) AddAllowedMacAddress(requestDevicesAddAllowedMacAddress *RequestDevicesAddAllowedMacAddress) (*ResponseDevicesAddAllowedMacAddress, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddAllowedMacAddress).
		SetResult(&ResponseDevicesAddAllowedMacAddress{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAllowedMacAddress(requestDevicesAddAllowedMacAddress)
		}

		return nil, response, fmt.Errorf("error with operation AddAllowedMacAddress")
	}

	result := response.Result().(*ResponseDevicesAddAllowedMacAddress)
	return result, response, err

}

//ThreatSummary Threat Summary - 3b98-98f0-4cfb-b74b
/* The Threat Summary for the Rogues and aWIPS



Documentation Link: https://developer.cisco.com/docs/dna-center/#!threat-summary
*/
func (s *DevicesService) ThreatSummary(requestDevicesThreatSummary *RequestDevicesThreatSummary) (*ResponseDevicesThreatSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/security/threats/summary"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesThreatSummary).
		SetResult(&ResponseDevicesThreatSummary{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ThreatSummary(requestDevicesThreatSummary)
		}

		return nil, response, fmt.Errorf("error with operation ThreatSummary")
	}

	result := response.Result().(*ResponseDevicesThreatSummary)
	return result, response, err

}

//GetDeviceInterfaceStatsInfoV2 Get Device Interface Stats Info - 76bb-5957-49ab-8a3b
/* This API returns the Interface Stats for the given Device Id. Please refer to the Feature tab for the Request Body usage and the API filtering support.


@param deviceID deviceId path parameter. Network Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-interface-stats-info-v2
*/
func (s *DevicesService) GetDeviceInterfaceStatsInfoV2(deviceID string, requestDevicesGetDeviceInterfaceStatsInfoV2 *RequestDevicesGetDeviceInterfaceStatsInfoV2) (*ResponseDevicesGetDeviceInterfaceStatsInfoV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/networkDevices/{deviceId}/interfaces/query"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesGetDeviceInterfaceStatsInfoV2).
		SetResult(&ResponseDevicesGetDeviceInterfaceStatsInfoV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInterfaceStatsInfoV2(deviceID, requestDevicesGetDeviceInterfaceStatsInfoV2)
		}

		return nil, response, fmt.Errorf("error with operation GetDeviceInterfaceStatsInfoV2")
	}

	result := response.Result().(*ResponseDevicesGetDeviceInterfaceStatsInfoV2)
	return result, response, err

}

//UpdatePlannedAccessPointForFloor Update Planned Access Point for Floor - 399f-596d-4f69-a080
/* Allows updating a planned access point on an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch the existing planned access points for the floor.  The payload to update a planned access point is in the same format, albeit a single object instead of a list, of that API.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

*/
func (s *DevicesService) UpdatePlannedAccessPointForFloor(floorID string, requestDevicesUpdatePlannedAccessPointForFloor *RequestDevicesUpdatePlannedAccessPointForFloor) (*ResponseDevicesUpdatePlannedAccessPointForFloor, *resty.Response, error) {
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdatePlannedAccessPointForFloor).
		SetResult(&ResponseDevicesUpdatePlannedAccessPointForFloor{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePlannedAccessPointForFloor(floorID, requestDevicesUpdatePlannedAccessPointForFloor)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePlannedAccessPointForFloor")
	}

	result := response.Result().(*ResponseDevicesUpdatePlannedAccessPointForFloor)
	return result, response, err

}

//UpdateHealthScoreDefinitionForTheGivenID Update health score definition for the given id. - f295-190f-4f08-bbe0
/* Update health threshold, include status of overall health status.
And also to synchronize with global profile issue thresholds of the definition for given id. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-issueAndHealthDefinitions-1.0.0-resolved.yaml


@param id id path parameter. Health score definition id.

*/
func (s *DevicesService) UpdateHealthScoreDefinitionForTheGivenID(id string, requestDevicesUpdateHealthScoreDefinitionForTheGivenId *RequestDevicesUpdateHealthScoreDefinitionForTheGivenID) (*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID, *resty.Response, error) {
	path := "/dna/intent/api/v1/healthScoreDefinitions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateHealthScoreDefinitionForTheGivenId).
		SetResult(&ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateHealthScoreDefinitionForTheGivenID(id, requestDevicesUpdateHealthScoreDefinitionForTheGivenId)
		}
		return nil, response, fmt.Errorf("error with operation UpdateHealthScoreDefinitionForTheGivenId")
	}

	result := response.Result().(*ResponseDevicesUpdateHealthScoreDefinitionForTheGivenID)
	return result, response, err

}

//UpdateInterfaceDetails Update Interface details - 868b-5a60-4be8-a2d7
/* Add/Update Interface description, VLAN membership, Voice VLAN and change Interface admin status ('UP'/'DOWN') from Request body.


@param interfaceUUID interfaceUuid path parameter. Interface ID

@param UpdateInterfaceDetailsQueryParams Filtering parameter
*/
func (s *DevicesService) UpdateInterfaceDetails(interfaceUUID string, requestDevicesUpdateInterfaceDetails *RequestDevicesUpdateInterfaceDetails, UpdateInterfaceDetailsQueryParams *UpdateInterfaceDetailsQueryParams) (*ResponseDevicesUpdateInterfaceDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/interface/{interfaceUuid}"
	path = strings.Replace(path, "{interfaceUuid}", fmt.Sprintf("%v", interfaceUUID), -1)

	queryString, _ := query.Values(UpdateInterfaceDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesUpdateInterfaceDetails).
		SetResult(&ResponseDevicesUpdateInterfaceDetails{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateInterfaceDetails(interfaceUUID, requestDevicesUpdateInterfaceDetails, UpdateInterfaceDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateInterfaceDetails")
	}

	result := response.Result().(*ResponseDevicesUpdateInterfaceDetails)
	return result, response, err

}

//UpdateDeviceDetails Update Device Details - aeb9-eb67-460b-92df
/* Update the credentials, management IP address of a given device (or a set of devices) in Catalyst Center and trigger an inventory sync.


 */
func (s *DevicesService) UpdateDeviceDetails(requestDevicesUpdateDeviceDetails *RequestDevicesUpdateDeviceDetails) (*ResponseDevicesUpdateDeviceDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceDetails).
		SetResult(&ResponseDevicesUpdateDeviceDetails{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceDetails(requestDevicesUpdateDeviceDetails)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceDetails")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceDetails)
	return result, response, err

}

//UpdateDeviceRole Update Device role - b985-5ad5-4ae9-8156
/* Updates the role of the device as access, core, distribution, border router


 */
func (s *DevicesService) UpdateDeviceRole(requestDevicesUpdateDeviceRole *RequestDevicesUpdateDeviceRole) (*ResponseDevicesUpdateDeviceRole, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/brief"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceRole).
		SetResult(&ResponseDevicesUpdateDeviceRole{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceRole(requestDevicesUpdateDeviceRole)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceRole")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceRole)
	return result, response, err

}

//SyncDevices Sync Devices - 3b9e-f967-4429-be4c
/* Synchronizes the devices. If forceSync param is false (default) then the sync would run in normal priority thread. If forceSync param is true then the sync would run in high priority thread if available, else the sync will fail. Result can be seen in the child task of each device


@param SyncDevicesQueryParams Filtering parameter
*/
func (s *DevicesService) SyncDevices(requestDevicesSyncDevices *RequestDevicesSyncDevices, SyncDevicesQueryParams *SyncDevicesQueryParams) (*ResponseDevicesSyncDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/sync"

	queryString, _ := query.Values(SyncDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestDevicesSyncDevices).
		SetResult(&ResponseDevicesSyncDevices{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SyncDevices(requestDevicesSyncDevices, SyncDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SyncDevices")
	}

	result := response.Result().(*ResponseDevicesSyncDevices)
	return result, response, err

}

//UpdateUserDefinedField Update User-Defined-Field - aa8c-ea8f-41aa-a346
/* Updates an existing global User Defined Field, using it's id.


@param id id path parameter. UDF id

*/
func (s *DevicesService) UpdateUserDefinedField(id string, requestDevicesUpdateUserDefinedField *RequestDevicesUpdateUserDefinedField) (*ResponseDevicesUpdateUserDefinedField, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/user-defined-field/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateUserDefinedField).
		SetResult(&ResponseDevicesUpdateUserDefinedField{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateUserDefinedField(id, requestDevicesUpdateUserDefinedField)
		}
		return nil, response, fmt.Errorf("error with operation UpdateUserDefinedField")
	}

	result := response.Result().(*ResponseDevicesUpdateUserDefinedField)
	return result, response, err

}

//AddUserDefinedFieldToDevice Add User-Defined-Field to device - d3af-395c-4669-adaf
/* Assigns an existing Global User-Defined-Field to a device. If the UDF is already assigned to the specific device, then it updates the device UDF value accordingly. Please note that the assigning UDF 'name' must be an existing global UDF. Otherwise error shall be shown.


@param deviceID deviceId path parameter. UUID of device to which UDF has to be added

*/
func (s *DevicesService) AddUserDefinedFieldToDevice(deviceID string, requestDevicesAddUserDefinedFieldToDevice *RequestDevicesAddUserDefinedFieldToDevice) (*ResponseDevicesAddUserDefinedFieldToDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceId}/user-defined-field"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesAddUserDefinedFieldToDevice).
		SetResult(&ResponseDevicesAddUserDefinedFieldToDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUserDefinedFieldToDevice(deviceID, requestDevicesAddUserDefinedFieldToDevice)
		}
		return nil, response, fmt.Errorf("error with operation AddUserDefinedFieldToDevice")
	}

	result := response.Result().(*ResponseDevicesAddUserDefinedFieldToDevice)
	return result, response, err

}

//UpdateDeviceManagementAddress Update Device Management Address - af93-b807-4feb-a985
/* This is a simple PUT API to edit the management IP Address of the device.


@param deviceid deviceid path parameter. The UUID of the device whose management IP address is to be updated.

*/
func (s *DevicesService) UpdateDeviceManagementAddress(deviceid string, requestDevicesUpdateDeviceManagementAddress *RequestDevicesUpdateDeviceManagementAddress) (*ResponseDevicesUpdateDeviceManagementAddress, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device/{deviceid}/management-address"
	path = strings.Replace(path, "{deviceid}", fmt.Sprintf("%v", deviceid), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateDeviceManagementAddress).
		SetResult(&ResponseDevicesUpdateDeviceManagementAddress{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceManagementAddress(deviceid, requestDevicesUpdateDeviceManagementAddress)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceManagementAddress")
	}

	result := response.Result().(*ResponseDevicesUpdateDeviceManagementAddress)
	return result, response, err

}

//UpdatesTheMaintenanceScheduleInformation Updates the maintenance schedule information - 8cb8-5ae3-4a2b-8a08
/* API to update the maintenance schedule for the network devices. The `maintenanceSchedule` can be updated only if the `status` value is `UPCOMING` or `IN_PROGRESS`. User can exit `IN_PROGRESS` maintenance window by setting the `endTime` to -1. This will update the endTime to the current time and exit the maintenance window immediately. When exiting the maintenance window, only the endTime will be updated while other parameters remain read-only.


@param id id path parameter. Unique identifier for the maintenance schedule

*/
func (s *DevicesService) UpdatesTheMaintenanceScheduleInformation(id string, requestDevicesUpdatesTheMaintenanceScheduleInformation *RequestDevicesUpdatesTheMaintenanceScheduleInformation) (*ResponseDevicesUpdatesTheMaintenanceScheduleInformation, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdatesTheMaintenanceScheduleInformation).
		SetResult(&ResponseDevicesUpdatesTheMaintenanceScheduleInformation{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesTheMaintenanceScheduleInformation(id, requestDevicesUpdatesTheMaintenanceScheduleInformation)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesTheMaintenanceScheduleInformation")
	}

	result := response.Result().(*ResponseDevicesUpdatesTheMaintenanceScheduleInformation)
	return result, response, err

}

//UpdateGlobalResyncInterval Update global resync interval - 25b5-39b4-4609-9e2a
/* Updates the resync interval (in minutes) globally for devices which do not have custom resync interval. To override this setting for all network devices refer to [/networkDevices/resyncIntervalSettings/override]


 */
func (s *DevicesService) UpdateGlobalResyncInterval(requestDevicesUpdateGlobalResyncInterval *RequestDevicesUpdateGlobalResyncInterval) (*ResponseDevicesUpdateGlobalResyncInterval, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/resyncIntervalSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateGlobalResyncInterval).
		SetResult(&ResponseDevicesUpdateGlobalResyncInterval{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalResyncInterval(requestDevicesUpdateGlobalResyncInterval)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalResyncInterval")
	}

	result := response.Result().(*ResponseDevicesUpdateGlobalResyncInterval)
	return result, response, err

}

//UpdateResyncIntervalForTheNetworkDevice Update resync interval for the network device - 92a0-db6c-428a-92d9
/* Update the resync interval (in minutes) for the given network device id.
To disable periodic resync, set interval as `0`.
To use global settings, set interval as `null`.


@param id id path parameter. The id of the network device.

*/
func (s *DevicesService) UpdateResyncIntervalForTheNetworkDevice(id string, requestDevicesUpdateResyncIntervalForTheNetworkDevice *RequestDevicesUpdateResyncIntervalForTheNetworkDevice) (*ResponseDevicesUpdateResyncIntervalForTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/resyncIntervalSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDevicesUpdateResyncIntervalForTheNetworkDevice).
		SetResult(&ResponseDevicesUpdateResyncIntervalForTheNetworkDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateResyncIntervalForTheNetworkDevice(id, requestDevicesUpdateResyncIntervalForTheNetworkDevice)
		}
		return nil, response, fmt.Errorf("error with operation UpdateResyncIntervalForTheNetworkDevice")
	}

	result := response.Result().(*ResponseDevicesUpdateResyncIntervalForTheNetworkDevice)
	return result, response, err

}

//DeletePlannedAccessPointForFloor Delete Planned Access Point for Floor - 6dad-1aac-4b3a-9e67
/* Allow to delete a planned access point from an existing floor map including its planned radio and antenna details.  Use the Get variant of this API to fetch the existing planned access points for the floor.  The instanceUUID listed in each of the planned access point attributes acts as the path param input to this API to delete that specific instance.


@param floorID floorId path parameter. The instance UUID of the floor hierarchy element

@param plannedAccessPointUUID plannedAccessPointUuid path parameter. The instance UUID of the planned access point to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-planned-access-point-for-floor
*/
func (s *DevicesService) DeletePlannedAccessPointForFloor(floorID string, plannedAccessPointUUID string) (*ResponseDevicesDeletePlannedAccessPointForFloor, *resty.Response, error) {
	//floorID string,plannedAccessPointUUID string
	path := "/dna/intent/api/v1/floors/{floorId}/planned-access-points/{plannedAccessPointUuid}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)
	path = strings.Replace(path, "{plannedAccessPointUuid}", fmt.Sprintf("%v", plannedAccessPointUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeletePlannedAccessPointForFloor{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePlannedAccessPointForFloor(floorID, plannedAccessPointUUID)
		}
		return nil, response, fmt.Errorf("error with operation DeletePlannedAccessPointForFloor")
	}

	result := response.Result().(*ResponseDevicesDeletePlannedAccessPointForFloor)
	return result, response, err

}

//DeleteUserDefinedField Delete User-Defined-Field - 78a3-c8b1-4799-892e
/* Deletes an existing Global User-Defined-Field using it's id.


@param id id path parameter. UDF id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-user-defined-field
*/
func (s *DevicesService) DeleteUserDefinedField(id string) (*ResponseDevicesDeleteUserDefinedField, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/network-device/user-defined-field/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeleteUserDefinedField{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteUserDefinedField(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteUserDefinedField")
	}

	result := response.Result().(*ResponseDevicesDeleteUserDefinedField)
	return result, response, err

}

//RemoveUserDefinedFieldFromDevice Remove User-Defined-Field from device - 8c9f-d9e8-4cab-bf96
/* Remove a User-Defined-Field from device. Name of UDF has to be passed as the query parameter. Please note that Global UDF will not be deleted by this operation.


@param deviceID deviceId path parameter. UUID of device from which UDF has to be removed

@param RemoveUserDefinedFieldFromDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-user-defined-field-from-device
*/
func (s *DevicesService) RemoveUserDefinedFieldFromDevice(deviceID string, RemoveUserDefinedFieldFromDeviceQueryParams *RemoveUserDefinedFieldFromDeviceQueryParams) (*ResponseDevicesRemoveUserDefinedFieldFromDevice, *resty.Response, error) {
	//deviceID string,RemoveUserDefinedFieldFromDeviceQueryParams *RemoveUserDefinedFieldFromDeviceQueryParams
	path := "/dna/intent/api/v1/network-device/{deviceId}/user-defined-field"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	queryString, _ := query.Values(RemoveUserDefinedFieldFromDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesRemoveUserDefinedFieldFromDevice{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveUserDefinedFieldFromDevice(deviceID, RemoveUserDefinedFieldFromDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RemoveUserDefinedFieldFromDevice")
	}

	result := response.Result().(*ResponseDevicesRemoveUserDefinedFieldFromDevice)
	return result, response, err

}

//DeleteDeviceByID Delete Device by Id - 1c89-4b58-48ea-b214
/* This API allows any network device that is not currently provisioned to be removed from the inventory. Important: Devices currently provisioned cannot be deleted. To delete a provisioned device, the device must be first deprovisioned.


@param id id path parameter. Device ID

@param DeleteDeviceByIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-by-id
*/
func (s *DevicesService) DeleteDeviceByID(id string, DeleteDeviceByIdQueryParams *DeleteDeviceByIDQueryParams) (*ResponseDevicesDeleteDeviceByID, *resty.Response, error) {
	//id string,DeleteDeviceByIdQueryParams *DeleteDeviceByIDQueryParams
	path := "/dna/intent/api/v1/network-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(DeleteDeviceByIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDevicesDeleteDeviceByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDeviceByID(id, DeleteDeviceByIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDeviceById")
	}

	result := response.Result().(*ResponseDevicesDeleteDeviceByID)
	return result, response, err

}

//DeleteMaintenanceSchedule Delete maintenance schedule. - f5b4-5b47-402a-a14a
/* API to delete maintenance schedule by id. Deletion is allowed if the maintenance window is in the `UPCOMING`, `COMPLETED`, or `FAILED` state. Deletion of maintenance schedule is not allowed if the maintenance window is currently `IN_PROGRESS`. To delete the maintenance schedule while it is `IN_PROGRESS`, first exit the current maintenance window using `PUT /dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}` API, and then proceed to delete the maintenance schedule.


@param id id path parameter. Unique identifier for the maintenance schedule


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-maintenance-schedule
*/
func (s *DevicesService) DeleteMaintenanceSchedule(id string) (*ResponseDevicesDeleteMaintenanceSchedule, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/networkDeviceMaintenanceSchedules/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesDeleteMaintenanceSchedule{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteMaintenanceSchedule(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteMaintenanceSchedule")
	}

	result := response.Result().(*ResponseDevicesDeleteMaintenanceSchedule)
	return result, response, err

}

//RemoveAllowedMacAddress Remove Allowed Mac Address - c8ac-a91b-4c5a-9b5c
/* Intent API to remove the threat mac address from allowed list.


@param macAddress macAddress path parameter. Threat mac address which needs to be removed from the allowed list. Multiple mac addresses will be removed if provided as comma separated values (example: 00:2A:10:51:22:43,00:2A:10:51:22:44). Note: In one request, maximum 100 mac addresses can be removed.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-allowed-mac-address
*/
func (s *DevicesService) RemoveAllowedMacAddress(macAddress string) (*ResponseDevicesRemoveAllowedMacAddress, *resty.Response, error) {
	//macAddress string
	path := "/dna/intent/api/v1/security/threats/rogue/allowed-list/{macAddress}"
	path = strings.Replace(path, "{macAddress}", fmt.Sprintf("%v", macAddress), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDevicesRemoveAllowedMacAddress{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveAllowedMacAddress(macAddress)
		}
		return nil, response, fmt.Errorf("error with operation RemoveAllowedMacAddress")
	}

	result := response.Result().(*ResponseDevicesRemoveAllowedMacAddress)
	return result, response, err

}
