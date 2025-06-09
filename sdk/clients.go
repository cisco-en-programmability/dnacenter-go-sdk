package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ClientsService service

type RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams struct {
	StartTime                  float64 `url:"startTime,omitempty"`                  //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                    float64 `url:"endTime,omitempty"`                    //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit                      float64 `url:"limit,omitempty"`                      //Maximum number of records to return
	Offset                     float64 `url:"offset,omitempty"`                     //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy                     string  `url:"sortBy,omitempty"`                     //A field within the response to sort by.
	Order                      string  `url:"order,omitempty"`                      //The sort order of the field ascending or descending.
	Type                       string  `url:"type,omitempty"`                       //The client device type whether client is connected to network through Wired or Wireless medium.
	OsType                     string  `url:"osType,omitempty"`                     //Client device operating system type. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*iOS*` or `iOS*` or `*iOS` Examples: `osType=iOS` (single osType requested) `osType=iOS&osType=Android` (multiple osType requested)
	OsVersion                  string  `url:"osVersion,omitempty"`                  //Client device operating system version This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*14.3*` or `14.3*` or `*14.3` Examples: `osVersion=14.3` (single osVersion requested) `osVersion=14.3&osVersion=10.1` (multiple osVersion requested)
	SiteHierarchy              string  `url:"siteHierarchy,omitempty"`              //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. "Global/AreaName/BuildingName/FloorName") This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*BuildingName*` or `BuildingName*` or `*BuildingName` Examples: `siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `siteHierarchy=Global/AreaName/BuildingName1/FloorName1&siteHierarchy=Global/AreaName/BuildingName1/FloorName2` (multiple siteHierarchy requested)
	SiteHierarchyID            string  `url:"siteHierarchyId,omitempty"`            //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. "globalUuid/areaUuid/buildingUuid/floorUuid") This field supports wildcard (`*`) character-based search.  Ex: `*buildingUuid*` or `buildingUuid*` or `*buildingUuid` Examples: `siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid` (single siteHierarchyId requested) `siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid1&siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid2` (multiple siteHierarchyId requested)
	SiteID                     string  `url:"siteId,omitempty"`                     //The site UUID without the top level hierarchy. (Ex."floorUuid") Examples: `siteId=floorUuid` (single siteId requested) `siteId=floorUuid1&siteId=floorUuid2` (multiple siteId requested)
	IPv4Address                string  `url:"ipv4Address,omitempty"`                //IPv4 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `ipv4Address=1.1.1.1` (single ipv4Address requested) `ipv4Address=1.1.1.1&ipv4Address=2.2.2.2` (multiple ipv4Address requested)
	IPv6Address                string  `url:"ipv6Address,omitempty"`                //IPv6 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search. Ex: `*2001:db8*` or `2001:db8*` or `*2001:db8` Examples: `ipv6Address=2001:db8:0:0:0:0:2:1` (single ipv6Address requested) `ipv6Address=2001:db8:0:0:0:0:2:1&ipv6Address=2001:db8:85a3:8d3:1319:8a2e:370:7348` (multiple ipv6Address requested)
	MacAddress                 string  `url:"macAddress,omitempty"`                 //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	WlcName                    string  `url:"wlcName,omitempty"`                    //Wireless Controller name that reports the wireless client. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*wlc-25*` or `wlc-25*` or `*wlc-25` Examples: `wlcName=wlc-25` (single wlcName requested) `wlcName=wlc-25&wlc-34` (multiple wlcName requested)
	ConnectedNetworkDeviceName string  `url:"connectedNetworkDeviceName,omitempty"` //Name of the neighbor network device that client is connected to. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*ap-25*` or `ap-25*` or `*ap-25` Examples: `connectedNetworkDeviceName=ap-25` (single connectedNetworkDeviceName requested) `connectedNetworkDeviceName=ap-25&ap-34` (multiple connectedNetworkDeviceName requested)
	SSID                       string  `url:"ssid,omitempty"`                       //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	Band                       string  `url:"band,omitempty"`                       //WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz - GHz Examples: `band=5GHZ` (single band requested) `band=2.4GHZ&band=6GHZ` (multiple band requested)
	View                       string  `url:"view,omitempty"`                       //Client related Views Refer to ClientView schema for list of views supported Examples: `view=Wireless` (single view requested) `view=WirelessHealth&view=WirelessTraffic` (multiple view requested)
	Attribute                  string  `url:"attribute,omitempty"`                  //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Refer to ClientAttribute schema for list of attributes supported Examples: `attribute=band` (single attribute requested) `attribute=band&attribute=ssid&attribute=overallScore` (multiple attribute requested)
}
type RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams struct {
	StartTime                  float64 `url:"startTime,omitempty"`                  //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime                    float64 `url:"endTime,omitempty"`                    //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Type                       string  `url:"type,omitempty"`                       //The client device type whether client is connected to network through Wired or Wireless medium.
	OsType                     string  `url:"osType,omitempty"`                     //Client device operating system type. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*iOS*` or `iOS*` or `*iOS` Examples: `osType=iOS` (single osType requested) `osType=iOS&osType=Android` (multiple osType requested)
	OsVersion                  string  `url:"osVersion,omitempty"`                  //Client device operating system version This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*14.3*` or `14.3*` or `*14.3` Examples: `osVersion=14.3` (single osVersion requested) `osVersion=14.3&osVersion=10.1` (multiple osVersion requested)
	SiteHierarchy              string  `url:"siteHierarchy,omitempty"`              //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. "Global/AreaName/BuildingName/FloorName") This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search.  Ex: `*BuildingName*` or `BuildingName*` or `*BuildingName` Examples: `siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `siteHierarchy=Global/AreaName/BuildingName1/FloorName1&siteHierarchy=Global/AreaName/BuildingName1/FloorName2` (multiple siteHierarchy requested)
	SiteHierarchyID            string  `url:"siteHierarchyId,omitempty"`            //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. "globalUuid/areaUuid/buildingUuid/floorUuid") This field supports wildcard (`*`) character-based search.  Ex: `*buildingUuid*` or `buildingUuid*` or `*buildingUuid` Examples: `siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid` (single siteHierarchyId requested) `siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid1&siteHierarchyId=globalUuid/areaUuid/buildingUuid1/floorUuid2` (multiple siteHierarchyId requested)
	SiteID                     string  `url:"siteId,omitempty"`                     //The site UUID without the top level hierarchy. (Ex."floorUuid") Examples: `siteId=floorUuid` (single siteId requested) `siteId=floorUuid1&siteId=floorUuid2` (multiple siteId requested)
	IPv4Address                string  `url:"ipv4Address,omitempty"`                //IPv4 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search. Ex: `*1.1*` or `1.1*` or `*1.1` Examples: `ipv4Address=1.1.1.1` (single ipv4Address requested) `ipv4Address=1.1.1.1&ipv4Address=2.2.2.2` (multiple ipv4Address requested)
	IPv6Address                string  `url:"ipv6Address,omitempty"`                //IPv6 Address of the network entity either network device or client This field supports wildcard (`*`) character-based search. Ex: `*2001:db8*` or `2001:db8*` or `*2001:db8` Examples: `ipv6Address=2001:db8:0:0:0:0:2:1` (single ipv6Address requested) `ipv6Address=2001:db8:0:0:0:0:2:1&ipv6Address=2001:db8:85a3:8d3:1319:8a2e:370:7348` (multiple ipv6Address requested)
	MacAddress                 string  `url:"macAddress,omitempty"`                 //The macAddress of the network device or client This field supports wildcard (`*`) character-based search. Ex: `*AB:AB:AB*` or `AB:AB:AB*` or `*AB:AB:AB` Examples: `macAddress=AB:AB:AB:CD:CD:CD` (single macAddress requested) `macAddress=AB:AB:AB:CD:CD:DC&macAddress=AB:AB:AB:CD:CD:FE` (multiple macAddress requested)
	WlcName                    string  `url:"wlcName,omitempty"`                    //Wireless Controller name that reports the wireless client. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*wlc-25*` or `wlc-25*` or `*wlc-25` Examples: `wlcName=wlc-25` (single wlcName requested) `wlcName=wlc-25&wlc-34` (multiple wlcName requested)
	ConnectedNetworkDeviceName string  `url:"connectedNetworkDeviceName,omitempty"` //Name of the neighbor network device that client is connected to. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*ap-25*` or `ap-25*` or `*ap-25` Examples: `connectedNetworkDeviceName=ap-25` (single connectedNetworkDeviceName requested) `connectedNetworkDeviceName=ap-25&ap-34` (multiple connectedNetworkDeviceName requested)
	SSID                       string  `url:"ssid,omitempty"`                       //SSID is the name of wireless network to which client connects to. It is also referred to as WLAN ID - Wireless Local Area Network Identifier. This field supports wildcard (`*`) character-based search. If the value contains the (`*`) character, please use the /query API for regex search. Ex: `*Alpha*` or `Alpha*` or `*Alpha` Examples: `ssid=Alpha` (single ssid requested) `ssid=Alpha&ssid=Guest` (multiple ssid requested)
	Band                       string  `url:"band,omitempty"`                       //WiFi frequency band that client or Access Point operates. Band value is represented in Giga Hertz - GHz Examples: `band=5GHZ` (single band requested) `band=2.4GHZ&band=6GHZ` (multiple band requested)
}
type RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificClientInformationMatchingTheMacaddressQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to current time.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	View      string  `url:"view,omitempty"`      //Client related Views Refer to ClientView schema for list of views supported Examples: `view=Wireless` (single view requested) `view=WirelessHealth&view=WirelessTraffic` (multiple view requested)
	Attribute string  `url:"attribute,omitempty"` //List of attributes related to resource that can be requested to only be part of the response along with the required attributes. Refer to ClientAttribute schema for list of attributes supported Examples: `attribute=band` (single attribute requested) `attribute=band&attribute=ssid&attribute=overallScore` (multiple attribute requested)
}
type RetrievesSpecificClientInformationMatchingTheMacaddressHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetClientsEnergyQueryParams struct {
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to one day before `endTime`.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.
	Limit             float64 `url:"limit,omitempty"`             //Maximum number of records to return
	Cursor            string  `url:"cursor,omitempty"`            //It's an opaque string field that indicates the next record in the requested collection. If no records remain, the API returns a response with a count of zero. The default value is an empty string, and the initial value must be an empty string. The cursor value is populated by the API in the response page block. If the user wants more records, the cursor in the subsequent request must be updated with the value from the previous response.
	SortBy            string  `url:"sortBy,omitempty"`            //A field within the response to sort by.
	Order             string  `url:"order,omitempty"`             //The sort order of the field ascending or descending.
	ID                string  `url:"id,omitempty"`                //The list of Mac addresses (e.g., `54:9F:C6:43:FF:80`). Examples: `id=54:9F:C6:43:FF:80` (single device requested) `id=54:9F:C6:43:FF:80&id=01:23:45:67:89:AB`
	SiteID            string  `url:"siteId,omitempty"`            //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	SiteHierarchy     string  `url:"siteHierarchy,omitempty"`     //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID   string  `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceCategory    string  `url:"deviceCategory,omitempty"`    //The list of device deviceCategories. Examples: `deviceCategory=AccessPoint` (single device family requested) `deviceCategory=AccessPoint&deviceCategory=OtherPOEDevice` (multiple device categories with comma separator)
	DeviceSubCategory string  `url:"deviceSubCategory,omitempty"` //The list of device sub categories. Examples: `deviceSubCategory=IP Phone 7821` (single sub category requested) `deviceSubCategory=IP Phone 7821&deviceSubCategory=IEEE PD`
	View              string  `url:"view,omitempty"`              //List of views. View and attribute work in union. Each view will include its attributes. For example, view device includes all the attributes related to device. Please refer to `ClientDeviceEnergyView` model for supported list of views Examples: `view=device&view=energy`
	Attribute         string  `url:"attribute,omitempty"`         //List of attributes. Please refer to `ClientDeviceEnergyAttribute` for supported list of attributes Examples: `attribute=id&attribute=energyConsumed`
}
type GetClientsEnergyHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountClientsEnergyQueryParams struct {
	StartTime         float64 `url:"startTime,omitempty"`         //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to one day before `endTime`.
	EndTime           float64 `url:"endTime,omitempty"`           //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.
	ID                string  `url:"id,omitempty"`                //The list of Mac addresses (e.g., `54:9F:C6:43:FF:80`). Examples: `id=54:9F:C6:43:FF:80` (single device requested) `id=54:9F:C6:43:FF:80&id=01:23:45:67:89:AB`
	SiteID            string  `url:"siteId,omitempty"`            //The UUID of the site. (Ex. `flooruuid`) Examples: `?siteId=id1` (single id requested) `?siteId=id1&siteId=id2&siteId=id3` (multiple ids requested)
	SiteHierarchy     string  `url:"siteHierarchy,omitempty"`     //The full hierarchical breakdown of the site tree starting from Global site name and ending with the specific site name. The Root site is named "Global" (Ex. `Global/AreaName/BuildingName/FloorName`) This field supports wildcard asterisk (`*`) character search support. E.g. `*/San*, */San, /San*` Examples: `?siteHierarchy=Global/AreaName/BuildingName/FloorName` (single siteHierarchy requested) `?siteHierarchy=Global/AreaName/BuildingName/FloorName&siteHierarchy=Global/AreaName2/BuildingName2/FloorName2` (multiple siteHierarchies requested)
	SiteHierarchyID   string  `url:"siteHierarchyId,omitempty"`   //The full hierarchy breakdown of the site tree in id form starting from Global site UUID and ending with the specific site UUID. (Ex. `globalUuid/areaUuid/buildingUuid/floorUuid`) This field supports wildcard asterisk (`*`) character search support. E.g. `*uuid*, *uuid, uuid*` Examples: `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid `(single siteHierarchyId requested) `?siteHierarchyId=globalUuid/areaUuid/buildingUuid/floorUuid&siteHierarchyId=globalUuid/areaUuid2/buildingUuid2/floorUuid2` (multiple siteHierarchyIds requested)
	DeviceCategory    string  `url:"deviceCategory,omitempty"`    //The list of device deviceCategories. Examples: `deviceCategory=AccessPoint` (single device family requested) `deviceCategory=AccessPoint&deviceCategory=OtherPOEDevice` (multiple device categories with comma separator)
	DeviceSubCategory string  `url:"deviceSubCategory,omitempty"` //The list of device sub categories. Examples: `deviceSubCategory=IP Phone 7821` (single sub category requested) `deviceSubCategory=IP Phone 7821&deviceSubCategory=IEEE PD`
}
type CountClientsEnergyHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type QueryClientsEnergyHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type CountClientsEnergyFromQueryHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetClientEnergyByIDQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `startTime` is not provided, API will default to one day before `endTime`.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive. If `endTime` is not provided, API will default to one day after `startTime`. If `startTime` is not provided either, API will default to current time.
	View      string  `url:"view,omitempty"`      //List of views. View and attribute work in union. Each view will include its attributes. For example, view device includes all the attributes related to device. Please refer to `ClientDeviceEnergyView` model for supported list of views Examples: `view=device&view=energy`
	Attribute string  `url:"attribute,omitempty"` //List of attributes. Please refer to `ClientDeviceEnergyAttribute` for supported list of attributes Examples: `attribute=id&attribute=energyConsumed`
}
type GetClientEnergyByIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetClientDetailQueryParams struct {
	MacAddress string  `url:"macAddress,omitempty"` //MAC Address of the client
	Timestamp  float64 `url:"timestamp,omitempty"`  //Epoch time(in milliseconds) when the Client health data is required
}
type GetClientEnrichmentDetailsHeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	IssueCategory     string `url:"issueCategory,omitempty"`       //Expects type string. The category of the DNA event based on which the underlying issues need to be fetched
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. For the enrichment details to be made available as part of the API response, this header must be set to true. This header must be explicitly passed when called from client applications outside Catalyst Center
}
type GetOverallClientHealthQueryParams struct {
	Timestamp float64 `url:"timestamp,omitempty"` //Epoch time(in milliseconds) when the Client health data is required
}
type ClientProximityQueryParams struct {
	Username       string  `url:"username,omitempty"`        //Wireless client username for which proximity information is required
	NumberDays     float64 `url:"number_days,omitempty"`     //Number of days to track proximity until current date. Defaults and maximum up to 14 days.
	TimeResolution float64 `url:"time_resolution,omitempty"` //Time interval (in minutes) to measure proximity. Defaults to 15 minutes with a minimum 5 minutes.
}

type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities struct {
	Response *[]ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponse `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesPage       `json:"page,omitempty"`     //
	Version  string                                                                                                   `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponse struct {
	ID                     string                                                                                                                       `json:"id,omitempty"`                     // Id
	MacAddress             string                                                                                                                       `json:"macAddress,omitempty"`             // Mac Address
	Type                   string                                                                                                                       `json:"type,omitempty"`                   // Type
	Name                   string                                                                                                                       `json:"name,omitempty"`                   // Name
	UserID                 string                                                                                                                       `json:"userId,omitempty"`                 // User Id
	Username               string                                                                                                                       `json:"username,omitempty"`               // Username
	IPv4Address            string                                                                                                                       `json:"ipv4Address,omitempty"`            // Ipv4 Address
	IPv6Addresses          []string                                                                                                                     `json:"ipv6Addresses,omitempty"`          // Ipv6 Addresses
	Vendor                 string                                                                                                                       `json:"vendor,omitempty"`                 // Vendor
	OsType                 string                                                                                                                       `json:"osType,omitempty"`                 // Os Type
	OsVersion              string                                                                                                                       `json:"osVersion,omitempty"`              // Os Version
	FormFactor             string                                                                                                                       `json:"formFactor,omitempty"`             // Form Factor
	SiteHierarchy          string                                                                                                                       `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                                                                                       `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteID                 string                                                                                                                       `json:"siteId,omitempty"`                 // Site Id
	LastUpdatedTime        *int                                                                                                                         `json:"lastUpdatedTime,omitempty"`        // Last Updated Time
	ConnectionStatus       string                                                                                                                       `json:"connectionStatus,omitempty"`       // Connection Status
	Tracked                string                                                                                                                       `json:"tracked,omitempty"`                // Tracked
	IsPrivateMacAddress    *bool                                                                                                                        `json:"isPrivateMacAddress,omitempty"`    // Is Private Mac Address
	Health                 *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseHealth                 `json:"health,omitempty"`                 //
	Traffic                *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseTraffic                `json:"traffic,omitempty"`                //
	ConnectedNetworkDevice *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseConnectedNetworkDevice `json:"connectedNetworkDevice,omitempty"` //
	Connection             *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseConnection             `json:"connection,omitempty"`             //
	Onboarding             *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseOnboarding             `json:"onboarding,omitempty"`             //
	Latency                *ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseLatency                `json:"latency,omitempty"`                //
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseHealth struct {
	OverallScore                 *int  `json:"overallScore,omitempty"`                 // Overall Score
	OnboardingScore              *int  `json:"onboardingScore,omitempty"`              // Onboarding Score
	ConnectedScore               *int  `json:"connectedScore,omitempty"`               // Connected Score
	LinkErrorPercentageThreshold *int  `json:"linkErrorPercentageThreshold,omitempty"` // Link Error Percentage Threshold
	IsLinkErrorIncluded          *bool `json:"isLinkErrorIncluded,omitempty"`          // Is Link Error Included
	RssiThreshold                *int  `json:"rssiThreshold,omitempty"`                // Rssi Threshold
	SnrThreshold                 *int  `json:"snrThreshold,omitempty"`                 // Snr Threshold
	IsRssiIncluded               *bool `json:"isRssiIncluded,omitempty"`               // Is Rssi Included
	IsSnrIncluded                *bool `json:"isSnrIncluded,omitempty"`                // Is Snr Included
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseTraffic struct {
	TxBytes               *int     `json:"txBytes,omitempty"`               // Tx Bytes
	RxBytes               *int     `json:"rxBytes,omitempty"`               // Rx Bytes
	Usage                 *int     `json:"usage,omitempty"`                 // Usage
	RxPackets             *int     `json:"rxPackets,omitempty"`             // Rx Packets
	TxPackets             *int     `json:"txPackets,omitempty"`             // Tx Packets
	RxRate                *float64 `json:"rxRate,omitempty"`                // Rx Rate
	TxRate                *int     `json:"txRate,omitempty"`                // Tx Rate
	RxLinkErrorPercentage *float64 `json:"rxLinkErrorPercentage,omitempty"` // Rx Link Error Percentage
	TxLinkErrorPercentage *float64 `json:"txLinkErrorPercentage,omitempty"` // Tx Link Error Percentage
	RxRetries             *int     `json:"rxRetries,omitempty"`             // Rx Retries
	RxRetryPercentage     *float64 `json:"rxRetryPercentage,omitempty"`     // Rx Retry Percentage
	TxDrops               *int     `json:"txDrops,omitempty"`               // Tx Drops
	TxDropPercentage      *int     `json:"txDropPercentage,omitempty"`      // Tx Drop Percentage
	DNSRequestCount       *int     `json:"dnsRequestCount,omitempty"`       // Dns Request Count
	DNSResponseCount      *int     `json:"dnsResponseCount,omitempty"`      // Dns Response Count
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseConnectedNetworkDevice struct {
	ConnectedNetworkDeviceID           string `json:"connectedNetworkDeviceId,omitempty"`           // Connected Network Device Id
	ConnectedNetworkDeviceName         string `json:"connectedNetworkDeviceName,omitempty"`         // Connected Network Device Name
	ConnectedNetworkDeviceManagementIP string `json:"connectedNetworkDeviceManagementIp,omitempty"` // Connected Network Device Management Ip
	ConnectedNetworkDeviceMac          string `json:"connectedNetworkDeviceMac,omitempty"`          // Connected Network Device Mac
	ConnectedNetworkDeviceType         string `json:"connectedNetworkDeviceType,omitempty"`         // Connected Network Device Type
	InterfaceName                      string `json:"interfaceName,omitempty"`                      // Interface Name
	InterfaceSpeed                     *int   `json:"interfaceSpeed,omitempty"`                     // Interface Speed
	DuplexMode                         string `json:"duplexMode,omitempty"`                         // Duplex Mode
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseConnection struct {
	VLANID                string `json:"vlanId,omitempty"`                // Vlan Id
	SessionDuration       *int   `json:"sessionDuration,omitempty"`       // Session Duration
	VnID                  string `json:"vnId,omitempty"`                  // Vn Id
	L2Vn                  string `json:"l2Vn,omitempty"`                  // L2 Vn
	L3Vn                  string `json:"l3Vn,omitempty"`                  // L3 Vn
	SecurityGroupTag      string `json:"securityGroupTag,omitempty"`      // Security Group Tag
	LinkSpeed             *int   `json:"linkSpeed,omitempty"`             // Link Speed
	BridgeVMMode          string `json:"bridgeVMMode,omitempty"`          // Bridge V M Mode
	Band                  string `json:"band,omitempty"`                  // Band
	SSID                  string `json:"ssid,omitempty"`                  // Ssid
	AuthType              string `json:"authType,omitempty"`              // Auth Type
	WlcName               string `json:"wlcName,omitempty"`               // Wlc Name
	WlcID                 string `json:"wlcId,omitempty"`                 // Wlc Id
	ApMac                 string `json:"apMac,omitempty"`                 // Ap Mac
	ApEthernetMac         string `json:"apEthernetMac,omitempty"`         // Ap Ethernet Mac
	ApMode                string `json:"apMode,omitempty"`                // Ap Mode
	RadioID               *int   `json:"radioId,omitempty"`               // Radio Id
	Channel               string `json:"channel,omitempty"`               // Channel
	ChannelWidth          string `json:"channelWidth,omitempty"`          // Channel Width
	Protocol              string `json:"protocol,omitempty"`              // Protocol
	ProtocolCapability    string `json:"protocolCapability,omitempty"`    // Protocol Capability
	UpnID                 string `json:"upnId,omitempty"`                 // Upn Id
	UpnName               string `json:"upnName,omitempty"`               // Upn Name
	UpnOwner              string `json:"upnOwner,omitempty"`              // Upn Owner
	UpnDuid               string `json:"upnDuid,omitempty"`               // Upn Duid
	Rssi                  *int   `json:"rssi,omitempty"`                  // Rssi
	Snr                   *int   `json:"snr,omitempty"`                   // Snr
	DataRate              *int   `json:"dataRate,omitempty"`              // Data Rate
	IsIosAnalyticsCapable *bool  `json:"isIosAnalyticsCapable,omitempty"` // Is Ios Analytics Capable
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseOnboarding struct {
	AvgRunDuration         *int   `json:"avgRunDuration,omitempty"`         // Avg Run Duration
	MaxRunDuration         *int   `json:"maxRunDuration,omitempty"`         // Max Run Duration
	AvgAssocDuration       *int   `json:"avgAssocDuration,omitempty"`       // Avg Assoc Duration
	MaxAssocDuration       *int   `json:"maxAssocDuration,omitempty"`       // Max Assoc Duration
	AvgAuthDuration        *int   `json:"avgAuthDuration,omitempty"`        // Avg Auth Duration
	MaxAuthDuration        *int   `json:"maxAuthDuration,omitempty"`        // Max Auth Duration
	AvgDhcpDuration        *int   `json:"avgDhcpDuration,omitempty"`        // Avg Dhcp Duration
	MaxDhcpDuration        *int   `json:"maxDhcpDuration,omitempty"`        // Max Dhcp Duration
	MaxRoamingDuration     *int   `json:"maxRoamingDuration,omitempty"`     // Max Roaming Duration
	AAAServerIP            string `json:"aaaServerIp,omitempty"`            // Aaa Server Ip
	DhcpServerIP           string `json:"dhcpServerIp,omitempty"`           // Dhcp Server Ip
	OnboardingTime         *int   `json:"onboardingTime,omitempty"`         // Onboarding Time
	AuthDoneTime           *int   `json:"authDoneTime,omitempty"`           // Auth Done Time
	AssocDoneTime          *int   `json:"assocDoneTime,omitempty"`          // Assoc Done Time
	DhcpDoneTime           *int   `json:"dhcpDoneTime,omitempty"`           // Dhcp Done Time
	RoamingTime            *int   `json:"roamingTime,omitempty"`            // Roaming Time
	FailedRoamingCount     *int   `json:"failedRoamingCount,omitempty"`     // Failed Roaming Count
	SuccessfulRoamingCount *int   `json:"successfulRoamingCount,omitempty"` // Successful Roaming Count
	TotalRoamingAttempts   *int   `json:"totalRoamingAttempts,omitempty"`   // Total Roaming Attempts
	AssocFailureReason     string `json:"assocFailureReason,omitempty"`     // Assoc Failure Reason
	AAAFailureReason       string `json:"aaaFailureReason,omitempty"`       // Aaa Failure Reason
	DhcpFailureReason      string `json:"dhcpFailureReason,omitempty"`      // Dhcp Failure Reason
	OtherFailureReason     string `json:"otherFailureReason,omitempty"`     // Other Failure Reason
	LatestFailureReason    string `json:"latestFailureReason,omitempty"`    // Latest Failure Reason
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesResponseLatency struct {
	Video      *int `json:"video,omitempty"`      // Video
	Voice      *int `json:"voice,omitempty"`      // Voice
	BestEffort *int `json:"bestEffort,omitempty"` // Best Effort
	Background *int `json:"background,omitempty"` // Background
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesPage struct {
	Limit  *int                                                                                                       `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                       `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                       `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFiltering struct {
	Response *ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringResponse `json:"response,omitempty"` //
	Version  string                                                                          `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFilteringResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes struct {
	Response *[]ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponse `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage       `json:"page,omitempty"`     //
	Version  string                                                                                                            `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponse struct {
	ID                     string                                                                                                                                `json:"id,omitempty"`                     // Id
	MacAddress             string                                                                                                                                `json:"macAddress,omitempty"`             // Mac Address
	Type                   string                                                                                                                                `json:"type,omitempty"`                   // Type
	Name                   string                                                                                                                                `json:"name,omitempty"`                   // Name
	UserID                 string                                                                                                                                `json:"userId,omitempty"`                 // User Id
	Username               string                                                                                                                                `json:"username,omitempty"`               // Username
	IPv4Address            string                                                                                                                                `json:"ipv4Address,omitempty"`            // Ipv4 Address
	IPv6Addresses          []string                                                                                                                              `json:"ipv6Addresses,omitempty"`          // Ipv6 Addresses
	Vendor                 string                                                                                                                                `json:"vendor,omitempty"`                 // Vendor
	OsType                 string                                                                                                                                `json:"osType,omitempty"`                 // Os Type
	OsVersion              string                                                                                                                                `json:"osVersion,omitempty"`              // Os Version
	FormFactor             string                                                                                                                                `json:"formFactor,omitempty"`             // Form Factor
	SiteHierarchy          string                                                                                                                                `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                                                                                                `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteID                 string                                                                                                                                `json:"siteId,omitempty"`                 // Site Id
	LastUpdatedTime        *int                                                                                                                                  `json:"lastUpdatedTime,omitempty"`        // Last Updated Time
	ConnectionStatus       string                                                                                                                                `json:"connectionStatus,omitempty"`       // Connection Status
	Tracked                string                                                                                                                                `json:"tracked,omitempty"`                // Tracked
	IsPrivateMacAddress    *bool                                                                                                                                 `json:"isPrivateMacAddress,omitempty"`    // Is Private Mac Address
	Health                 *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseHealth                 `json:"health,omitempty"`                 //
	Traffic                *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseTraffic                `json:"traffic,omitempty"`                //
	ConnectedNetworkDevice *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseConnectedNetworkDevice `json:"connectedNetworkDevice,omitempty"` //
	Connection             *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseConnection             `json:"connection,omitempty"`             //
	Onboarding             *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseOnboarding             `json:"onboarding,omitempty"`             //
	Latency                *ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseLatency                `json:"latency,omitempty"`                //
	AggregateAttributes    *[]ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseAggregateAttributes  `json:"aggregateAttributes,omitempty"`    //
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseHealth struct {
	OverallScore                 *int  `json:"overallScore,omitempty"`                 // Overall Score
	OnboardingScore              *int  `json:"onboardingScore,omitempty"`              // Onboarding Score
	ConnectedScore               *int  `json:"connectedScore,omitempty"`               // Connected Score
	LinkErrorPercentageThreshold *int  `json:"linkErrorPercentageThreshold,omitempty"` // Link Error Percentage Threshold
	IsLinkErrorIncluded          *bool `json:"isLinkErrorIncluded,omitempty"`          // Is Link Error Included
	RssiThreshold                *int  `json:"rssiThreshold,omitempty"`                // Rssi Threshold
	SnrThreshold                 *int  `json:"snrThreshold,omitempty"`                 // Snr Threshold
	IsRssiIncluded               *bool `json:"isRssiIncluded,omitempty"`               // Is Rssi Included
	IsSnrIncluded                *bool `json:"isSnrIncluded,omitempty"`                // Is Snr Included
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseTraffic struct {
	TxBytes               *int     `json:"txBytes,omitempty"`               // Tx Bytes
	RxBytes               *int     `json:"rxBytes,omitempty"`               // Rx Bytes
	Usage                 *int     `json:"usage,omitempty"`                 // Usage
	RxPackets             *int     `json:"rxPackets,omitempty"`             // Rx Packets
	TxPackets             *int     `json:"txPackets,omitempty"`             // Tx Packets
	RxRate                *float64 `json:"rxRate,omitempty"`                // Rx Rate
	TxRate                *int     `json:"txRate,omitempty"`                // Tx Rate
	RxLinkErrorPercentage *float64 `json:"rxLinkErrorPercentage,omitempty"` // Rx Link Error Percentage
	TxLinkErrorPercentage *float64 `json:"txLinkErrorPercentage,omitempty"` // Tx Link Error Percentage
	RxRetries             *int     `json:"rxRetries,omitempty"`             // Rx Retries
	RxRetryPercentage     *float64 `json:"rxRetryPercentage,omitempty"`     // Rx Retry Percentage
	TxDrops               *int     `json:"txDrops,omitempty"`               // Tx Drops
	TxDropPercentage      *int     `json:"txDropPercentage,omitempty"`      // Tx Drop Percentage
	DNSRequestCount       *int     `json:"dnsRequestCount,omitempty"`       // Dns Request Count
	DNSResponseCount      *int     `json:"dnsResponseCount,omitempty"`      // Dns Response Count
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseConnectedNetworkDevice struct {
	ConnectedNetworkDeviceID           string `json:"connectedNetworkDeviceId,omitempty"`           // Connected Network Device Id
	ConnectedNetworkDeviceName         string `json:"connectedNetworkDeviceName,omitempty"`         // Connected Network Device Name
	ConnectedNetworkDeviceManagementIP string `json:"connectedNetworkDeviceManagementIp,omitempty"` // Connected Network Device Management Ip
	ConnectedNetworkDeviceMac          string `json:"connectedNetworkDeviceMac,omitempty"`          // Connected Network Device Mac
	ConnectedNetworkDeviceType         string `json:"connectedNetworkDeviceType,omitempty"`         // Connected Network Device Type
	InterfaceName                      string `json:"interfaceName,omitempty"`                      // Interface Name
	InterfaceSpeed                     *int   `json:"interfaceSpeed,omitempty"`                     // Interface Speed
	DuplexMode                         string `json:"duplexMode,omitempty"`                         // Duplex Mode
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseConnection struct {
	VLANID                string `json:"vlanId,omitempty"`                // Vlan Id
	SessionDuration       *int   `json:"sessionDuration,omitempty"`       // Session Duration
	VnID                  string `json:"vnId,omitempty"`                  // Vn Id
	L2Vn                  string `json:"l2Vn,omitempty"`                  // L2 Vn
	L3Vn                  string `json:"l3Vn,omitempty"`                  // L3 Vn
	SecurityGroupTag      string `json:"securityGroupTag,omitempty"`      // Security Group Tag
	LinkSpeed             *int   `json:"linkSpeed,omitempty"`             // Link Speed
	BridgeVMMode          string `json:"bridgeVMMode,omitempty"`          // Bridge V M Mode
	Band                  string `json:"band,omitempty"`                  // Band
	SSID                  string `json:"ssid,omitempty"`                  // Ssid
	AuthType              string `json:"authType,omitempty"`              // Auth Type
	WlcName               string `json:"wlcName,omitempty"`               // Wlc Name
	WlcID                 string `json:"wlcId,omitempty"`                 // Wlc Id
	ApMac                 string `json:"apMac,omitempty"`                 // Ap Mac
	ApEthernetMac         string `json:"apEthernetMac,omitempty"`         // Ap Ethernet Mac
	ApMode                string `json:"apMode,omitempty"`                // Ap Mode
	RadioID               *int   `json:"radioId,omitempty"`               // Radio Id
	Channel               string `json:"channel,omitempty"`               // Channel
	ChannelWidth          string `json:"channelWidth,omitempty"`          // Channel Width
	Protocol              string `json:"protocol,omitempty"`              // Protocol
	ProtocolCapability    string `json:"protocolCapability,omitempty"`    // Protocol Capability
	UpnID                 string `json:"upnId,omitempty"`                 // Upn Id
	UpnName               string `json:"upnName,omitempty"`               // Upn Name
	UpnOwner              string `json:"upnOwner,omitempty"`              // Upn Owner
	UpnDuid               string `json:"upnDuid,omitempty"`               // Upn Duid
	Rssi                  *int   `json:"rssi,omitempty"`                  // Rssi
	Snr                   *int   `json:"snr,omitempty"`                   // Snr
	DataRate              *int   `json:"dataRate,omitempty"`              // Data Rate
	IsIosAnalyticsCapable *bool  `json:"isIosAnalyticsCapable,omitempty"` // Is Ios Analytics Capable
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseOnboarding struct {
	AvgRunDuration         *int   `json:"avgRunDuration,omitempty"`         // Avg Run Duration
	MaxRunDuration         *int   `json:"maxRunDuration,omitempty"`         // Max Run Duration
	AvgAssocDuration       *int   `json:"avgAssocDuration,omitempty"`       // Avg Assoc Duration
	MaxAssocDuration       *int   `json:"maxAssocDuration,omitempty"`       // Max Assoc Duration
	AvgAuthDuration        *int   `json:"avgAuthDuration,omitempty"`        // Avg Auth Duration
	MaxAuthDuration        *int   `json:"maxAuthDuration,omitempty"`        // Max Auth Duration
	AvgDhcpDuration        *int   `json:"avgDhcpDuration,omitempty"`        // Avg Dhcp Duration
	MaxDhcpDuration        *int   `json:"maxDhcpDuration,omitempty"`        // Max Dhcp Duration
	MaxRoamingDuration     *int   `json:"maxRoamingDuration,omitempty"`     // Max Roaming Duration
	AAAServerIP            string `json:"aaaServerIp,omitempty"`            // Aaa Server Ip
	DhcpServerIP           string `json:"dhcpServerIp,omitempty"`           // Dhcp Server Ip
	OnboardingTime         *int   `json:"onboardingTime,omitempty"`         // Onboarding Time
	AuthDoneTime           *int   `json:"authDoneTime,omitempty"`           // Auth Done Time
	AssocDoneTime          *int   `json:"assocDoneTime,omitempty"`          // Assoc Done Time
	DhcpDoneTime           *int   `json:"dhcpDoneTime,omitempty"`           // Dhcp Done Time
	RoamingTime            *int   `json:"roamingTime,omitempty"`            // Roaming Time
	FailedRoamingCount     *int   `json:"failedRoamingCount,omitempty"`     // Failed Roaming Count
	SuccessfulRoamingCount *int   `json:"successfulRoamingCount,omitempty"` // Successful Roaming Count
	TotalRoamingAttempts   *int   `json:"totalRoamingAttempts,omitempty"`   // Total Roaming Attempts
	AssocFailureReason     string `json:"assocFailureReason,omitempty"`     // Assoc Failure Reason
	AAAFailureReason       string `json:"aaaFailureReason,omitempty"`       // Aaa Failure Reason
	DhcpFailureReason      string `json:"dhcpFailureReason,omitempty"`      // Dhcp Failure Reason
	OtherFailureReason     string `json:"otherFailureReason,omitempty"`     // Other Failure Reason
	LatestFailureReason    string `json:"latestFailureReason,omitempty"`    // Latest Failure Reason
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseLatency struct {
	Video      *int `json:"video,omitempty"`      // Video
	Voice      *int `json:"voice,omitempty"`      // Voice
	BestEffort *int `json:"bestEffort,omitempty"` // Best Effort
	Background *int `json:"background,omitempty"` // Background
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage struct {
	Limit  *int                                                                                                                `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                                `json:"offset,omitempty"` // Offset
	Count  *int                                                                                                                `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFilters struct {
	Response *ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClients struct {
	Response *ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponse `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsPage     `json:"page,omitempty"`     //
	Version  string                                                                `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponse struct {
	Groups *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponseGroups `json:"groups,omitempty"` //
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponseGroups struct {
	ID                  string                                                                                           `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsPage struct {
	Limit  *int                                                                      `json:"limit,omitempty"`  // Limit
	Cursor string                                                                    `json:"cursor,omitempty"` // Cursor
	Count  *int                                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClientsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClients struct {
	Response *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponse `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPage       `json:"page,omitempty"`     //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponse struct {
	ID                  string                                                                                     `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponseAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponseAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponseAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsResponseAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPage struct {
	Limit  *int                                                                      `json:"limit,omitempty"`  // Limit
	Cursor string                                                                    `json:"cursor,omitempty"` // Cursor
	Count  *int                                                                      `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClients struct {
	Response *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponse `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsPage       `json:"page,omitempty"`     //
	Version  string                                                                   `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponse struct {
	Timestamp *int                                                                           `json:"timestamp,omitempty"` // Timestamp
	Groups    *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponseGroups `json:"groups,omitempty"`    //
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponseGroups struct {
	ID                  string                                                                                            `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClientsPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddress struct {
	Response *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponse `json:"response,omitempty"` //
	Version  string                                                                          `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponse struct {
	ID                     string                                                                                                `json:"id,omitempty"`                     // Id
	MacAddress             string                                                                                                `json:"macAddress,omitempty"`             // Mac Address
	Type                   string                                                                                                `json:"type,omitempty"`                   // Type
	Name                   string                                                                                                `json:"name,omitempty"`                   // Name
	UserID                 string                                                                                                `json:"userId,omitempty"`                 // User Id
	Username               string                                                                                                `json:"username,omitempty"`               // Username
	IPv4Address            string                                                                                                `json:"ipv4Address,omitempty"`            // Ipv4 Address
	IPv6Addresses          []string                                                                                              `json:"ipv6Addresses,omitempty"`          // Ipv6 Addresses
	Vendor                 string                                                                                                `json:"vendor,omitempty"`                 // Vendor
	OsType                 string                                                                                                `json:"osType,omitempty"`                 // Os Type
	OsVersion              string                                                                                                `json:"osVersion,omitempty"`              // Os Version
	FormFactor             string                                                                                                `json:"formFactor,omitempty"`             // Form Factor
	SiteHierarchy          string                                                                                                `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                                                                `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	SiteID                 string                                                                                                `json:"siteId,omitempty"`                 // Site Id
	LastUpdatedTime        *int                                                                                                  `json:"lastUpdatedTime,omitempty"`        // Last Updated Time
	ConnectionStatus       string                                                                                                `json:"connectionStatus,omitempty"`       // Connection Status
	Tracked                string                                                                                                `json:"tracked,omitempty"`                // Tracked
	IsPrivateMacAddress    *bool                                                                                                 `json:"isPrivateMacAddress,omitempty"`    // Is Private Mac Address
	Health                 *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseHealth                 `json:"health,omitempty"`                 //
	Traffic                *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseTraffic                `json:"traffic,omitempty"`                //
	ConnectedNetworkDevice *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseConnectedNetworkDevice `json:"connectedNetworkDevice,omitempty"` //
	Connection             *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseConnection             `json:"connection,omitempty"`             //
	Onboarding             *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseOnboarding             `json:"onboarding,omitempty"`             //
	Latency                *ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseLatency                `json:"latency,omitempty"`                //
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseHealth struct {
	OverallScore                 *int  `json:"overallScore,omitempty"`                 // Overall Score
	OnboardingScore              *int  `json:"onboardingScore,omitempty"`              // Onboarding Score
	ConnectedScore               *int  `json:"connectedScore,omitempty"`               // Connected Score
	LinkErrorPercentageThreshold *int  `json:"linkErrorPercentageThreshold,omitempty"` // Link Error Percentage Threshold
	IsLinkErrorIncluded          *bool `json:"isLinkErrorIncluded,omitempty"`          // Is Link Error Included
	RssiThreshold                *int  `json:"rssiThreshold,omitempty"`                // Rssi Threshold
	SnrThreshold                 *int  `json:"snrThreshold,omitempty"`                 // Snr Threshold
	IsRssiIncluded               *bool `json:"isRssiIncluded,omitempty"`               // Is Rssi Included
	IsSnrIncluded                *bool `json:"isSnrIncluded,omitempty"`                // Is Snr Included
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseTraffic struct {
	TxBytes               *int     `json:"txBytes,omitempty"`               // Tx Bytes
	RxBytes               *int     `json:"rxBytes,omitempty"`               // Rx Bytes
	Usage                 *int     `json:"usage,omitempty"`                 // Usage
	RxPackets             *int     `json:"rxPackets,omitempty"`             // Rx Packets
	TxPackets             *int     `json:"txPackets,omitempty"`             // Tx Packets
	RxRate                *float64 `json:"rxRate,omitempty"`                // Rx Rate
	TxRate                *int     `json:"txRate,omitempty"`                // Tx Rate
	RxLinkErrorPercentage *float64 `json:"rxLinkErrorPercentage,omitempty"` // Rx Link Error Percentage
	TxLinkErrorPercentage *float64 `json:"txLinkErrorPercentage,omitempty"` // Tx Link Error Percentage
	RxRetries             *int     `json:"rxRetries,omitempty"`             // Rx Retries
	RxRetryPercentage     *float64 `json:"rxRetryPercentage,omitempty"`     // Rx Retry Percentage
	TxDrops               *int     `json:"txDrops,omitempty"`               // Tx Drops
	TxDropPercentage      *int     `json:"txDropPercentage,omitempty"`      // Tx Drop Percentage
	DNSRequestCount       *int     `json:"dnsRequestCount,omitempty"`       // Dns Request Count
	DNSResponseCount      *int     `json:"dnsResponseCount,omitempty"`      // Dns Response Count
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseConnectedNetworkDevice struct {
	ConnectedNetworkDeviceID           string `json:"connectedNetworkDeviceId,omitempty"`           // Connected Network Device Id
	ConnectedNetworkDeviceName         string `json:"connectedNetworkDeviceName,omitempty"`         // Connected Network Device Name
	ConnectedNetworkDeviceManagementIP string `json:"connectedNetworkDeviceManagementIp,omitempty"` // Connected Network Device Management Ip
	ConnectedNetworkDeviceMac          string `json:"connectedNetworkDeviceMac,omitempty"`          // Connected Network Device Mac
	ConnectedNetworkDeviceType         string `json:"connectedNetworkDeviceType,omitempty"`         // Connected Network Device Type
	InterfaceName                      string `json:"interfaceName,omitempty"`                      // Interface Name
	InterfaceSpeed                     *int   `json:"interfaceSpeed,omitempty"`                     // Interface Speed
	DuplexMode                         string `json:"duplexMode,omitempty"`                         // Duplex Mode
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseConnection struct {
	VLANID                string `json:"vlanId,omitempty"`                // Vlan Id
	SessionDuration       *int   `json:"sessionDuration,omitempty"`       // Session Duration
	VnID                  string `json:"vnId,omitempty"`                  // Vn Id
	L2Vn                  string `json:"l2Vn,omitempty"`                  // L2 Vn
	L3Vn                  string `json:"l3Vn,omitempty"`                  // L3 Vn
	SecurityGroupTag      string `json:"securityGroupTag,omitempty"`      // Security Group Tag
	LinkSpeed             *int   `json:"linkSpeed,omitempty"`             // Link Speed
	BridgeVMMode          string `json:"bridgeVMMode,omitempty"`          // Bridge V M Mode
	Band                  string `json:"band,omitempty"`                  // Band
	SSID                  string `json:"ssid,omitempty"`                  // Ssid
	AuthType              string `json:"authType,omitempty"`              // Auth Type
	WlcName               string `json:"wlcName,omitempty"`               // Wlc Name
	WlcID                 string `json:"wlcId,omitempty"`                 // Wlc Id
	ApMac                 string `json:"apMac,omitempty"`                 // Ap Mac
	ApEthernetMac         string `json:"apEthernetMac,omitempty"`         // Ap Ethernet Mac
	ApMode                string `json:"apMode,omitempty"`                // Ap Mode
	RadioID               *int   `json:"radioId,omitempty"`               // Radio Id
	Channel               string `json:"channel,omitempty"`               // Channel
	ChannelWidth          string `json:"channelWidth,omitempty"`          // Channel Width
	Protocol              string `json:"protocol,omitempty"`              // Protocol
	ProtocolCapability    string `json:"protocolCapability,omitempty"`    // Protocol Capability
	UpnID                 string `json:"upnId,omitempty"`                 // Upn Id
	UpnName               string `json:"upnName,omitempty"`               // Upn Name
	UpnOwner              string `json:"upnOwner,omitempty"`              // Upn Owner
	UpnDuid               string `json:"upnDuid,omitempty"`               // Upn Duid
	Rssi                  *int   `json:"rssi,omitempty"`                  // Rssi
	Snr                   *int   `json:"snr,omitempty"`                   // Snr
	DataRate              *int   `json:"dataRate,omitempty"`              // Data Rate
	IsIosAnalyticsCapable *bool  `json:"isIosAnalyticsCapable,omitempty"` // Is Ios Analytics Capable
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseOnboarding struct {
	AvgRunDuration         *int   `json:"avgRunDuration,omitempty"`         // Avg Run Duration
	MaxRunDuration         *int   `json:"maxRunDuration,omitempty"`         // Max Run Duration
	AvgAssocDuration       *int   `json:"avgAssocDuration,omitempty"`       // Avg Assoc Duration
	MaxAssocDuration       *int   `json:"maxAssocDuration,omitempty"`       // Max Assoc Duration
	AvgAuthDuration        *int   `json:"avgAuthDuration,omitempty"`        // Avg Auth Duration
	MaxAuthDuration        *int   `json:"maxAuthDuration,omitempty"`        // Max Auth Duration
	AvgDhcpDuration        *int   `json:"avgDhcpDuration,omitempty"`        // Avg Dhcp Duration
	MaxDhcpDuration        *int   `json:"maxDhcpDuration,omitempty"`        // Max Dhcp Duration
	MaxRoamingDuration     *int   `json:"maxRoamingDuration,omitempty"`     // Max Roaming Duration
	AAAServerIP            string `json:"aaaServerIp,omitempty"`            // Aaa Server Ip
	DhcpServerIP           string `json:"dhcpServerIp,omitempty"`           // Dhcp Server Ip
	OnboardingTime         *int   `json:"onboardingTime,omitempty"`         // Onboarding Time
	AuthDoneTime           *int   `json:"authDoneTime,omitempty"`           // Auth Done Time
	AssocDoneTime          *int   `json:"assocDoneTime,omitempty"`          // Assoc Done Time
	DhcpDoneTime           *int   `json:"dhcpDoneTime,omitempty"`           // Dhcp Done Time
	RoamingTime            *int   `json:"roamingTime,omitempty"`            // Roaming Time
	FailedRoamingCount     *int   `json:"failedRoamingCount,omitempty"`     // Failed Roaming Count
	SuccessfulRoamingCount *int   `json:"successfulRoamingCount,omitempty"` // Successful Roaming Count
	TotalRoamingAttempts   *int   `json:"totalRoamingAttempts,omitempty"`   // Total Roaming Attempts
	AssocFailureReason     string `json:"assocFailureReason,omitempty"`     // Assoc Failure Reason
	AAAFailureReason       string `json:"aaaFailureReason,omitempty"`       // Aaa Failure Reason
	DhcpFailureReason      string `json:"dhcpFailureReason,omitempty"`      // Dhcp Failure Reason
	OtherFailureReason     string `json:"otherFailureReason,omitempty"`     // Other Failure Reason
	LatestFailureReason    string `json:"latestFailureReason,omitempty"`    // Latest Failure Reason
}
type ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddressResponseLatency struct {
	Video      *int `json:"video,omitempty"`      // Video
	Voice      *int `json:"voice,omitempty"`      // Voice
	BestEffort *int `json:"bestEffort,omitempty"` // Best Effort
	Background *int `json:"background,omitempty"` // Background
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime struct {
	Response *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponse `json:"response,omitempty"` //
	Page     *ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimePage       `json:"page,omitempty"`     //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponse struct {
	Timestamp *int                                                                                         `json:"timestamp,omitempty"` // Timestamp
	Groups    *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponseGroups `json:"groups,omitempty"`    //
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponseGroups struct {
	ID                  string                                                                                                          `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value *int   `json:"value,omitempty"` // Value
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeResponseGroupsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
	Value    *int   `json:"value,omitempty"`    // Value
}
type ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimePage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	Count         *int   `json:"count,omitempty"`         // Count
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseClientsGetClientsEnergy struct {
	Response *[]ResponseClientsGetClientsEnergyResponse `json:"response,omitempty"` //
	Page     *ResponseClientsGetClientsEnergyPage       `json:"page,omitempty"`     //
	Version  string                                     `json:"version,omitempty"`  // Version
}
type ResponseClientsGetClientsEnergyResponse struct {
	ID                     string   `json:"id,omitempty"`                     // Id
	DeviceName             string   `json:"deviceName,omitempty"`             // Device Name
	DeviceCategory         string   `json:"deviceCategory,omitempty"`         // Device Category
	DeviceSubCategory      string   `json:"deviceSubCategory,omitempty"`      // Device Sub Category
	SiteID                 string   `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchy          string   `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string   `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	EnergyConsumed         *float64 `json:"energyConsumed,omitempty"`         // Energy Consumed
	EstimatedCost          *float64 `json:"estimatedCost,omitempty"`          // Estimated Cost
	EstimatedEmission      *float64 `json:"estimatedEmission,omitempty"`      // Estimated Emission
	CarbonIntensity        *float64 `json:"carbonIntensity,omitempty"`        // Carbon Intensity
	ConnectedDeviceName    string   `json:"connectedDeviceName,omitempty"`    // Connected Device Name
	ConnectedInterfaceName string   `json:"connectedInterfaceName,omitempty"` // Connected Interface Name
}
type ResponseClientsGetClientsEnergyPage struct {
	Limit  *int                                         `json:"limit,omitempty"`  // Limit
	Cursor string                                       `json:"cursor,omitempty"` // Cursor
	Count  *int                                         `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsGetClientsEnergyPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsGetClientsEnergyPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseClientsCountClientsEnergy struct {
	Response *ResponseClientsCountClientsEnergyResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version
}
type ResponseClientsCountClientsEnergyResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseClientsQueryClientsEnergy struct {
	Response *[]ResponseClientsQueryClientsEnergyResponse `json:"response,omitempty"` //
	Page     *ResponseClientsQueryClientsEnergyPage       `json:"page,omitempty"`     //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseClientsQueryClientsEnergyResponse struct {
	ID                     string                                                          `json:"id,omitempty"`                     // Id
	DeviceName             string                                                          `json:"deviceName,omitempty"`             // Device Name
	DeviceCategory         string                                                          `json:"deviceCategory,omitempty"`         // Device Category
	DeviceSubCategory      string                                                          `json:"deviceSubCategory,omitempty"`      // Device Sub Category
	SiteID                 string                                                          `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchy          string                                                          `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string                                                          `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	EnergyConsumed         *float64                                                        `json:"energyConsumed,omitempty"`         // Energy Consumed
	EstimatedCost          *float64                                                        `json:"estimatedCost,omitempty"`          // Estimated Cost
	EstimatedEmission      *float64                                                        `json:"estimatedEmission,omitempty"`      // Estimated Emission
	CarbonIntensity        *float64                                                        `json:"carbonIntensity,omitempty"`        // Carbon Intensity
	ConnectedDeviceName    string                                                          `json:"connectedDeviceName,omitempty"`    // Connected Device Name
	ConnectedInterfaceName string                                                          `json:"connectedInterfaceName,omitempty"` // Connected Interface Name
	AggregateAttributes    *[]ResponseClientsQueryClientsEnergyResponseAggregateAttributes `json:"aggregateAttributes,omitempty"`    //
}
type ResponseClientsQueryClientsEnergyResponseAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseClientsQueryClientsEnergyPage struct {
	Limit  *int                                           `json:"limit,omitempty"`  // Limit
	Cursor string                                         `json:"cursor,omitempty"` // Cursor
	Count  *int                                           `json:"count,omitempty"`  // Count
	SortBy *[]ResponseClientsQueryClientsEnergyPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseClientsQueryClientsEnergyPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type ResponseClientsCountClientsEnergyFromQuery struct {
	Response *ResponseClientsCountClientsEnergyFromQueryResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseClientsCountClientsEnergyFromQueryResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseClientsGetClientEnergyByID struct {
	Response *ResponseClientsGetClientEnergyByIDResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version
}
type ResponseClientsGetClientEnergyByIDResponse struct {
	ID                     string   `json:"id,omitempty"`                     // Id
	DeviceName             string   `json:"deviceName,omitempty"`             // Device Name
	DeviceCategory         string   `json:"deviceCategory,omitempty"`         // Device Category
	DeviceSubCategory      string   `json:"deviceSubCategory,omitempty"`      // Device Sub Category
	SiteID                 string   `json:"siteId,omitempty"`                 // Site Id
	SiteHierarchy          string   `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	SiteHierarchyID        string   `json:"siteHierarchyId,omitempty"`        // Site Hierarchy Id
	EnergyConsumed         *float64 `json:"energyConsumed,omitempty"`         // Energy Consumed
	EstimatedCost          *float64 `json:"estimatedCost,omitempty"`          // Estimated Cost
	EstimatedEmission      *float64 `json:"estimatedEmission,omitempty"`      // Estimated Emission
	CarbonIntensity        *float64 `json:"carbonIntensity,omitempty"`        // Carbon Intensity
	ConnectedDeviceName    string   `json:"connectedDeviceName,omitempty"`    // Connected Device Name
	ConnectedInterfaceName string   `json:"connectedInterfaceName,omitempty"` // Connected Interface Name
}
type ResponseClientsGetClientDetail struct {
	Detail         *ResponseClientsGetClientDetailDetail         `json:"detail,omitempty"`         //
	ConnectionInfo *ResponseClientsGetClientDetailConnectionInfo `json:"connectionInfo,omitempty"` //
	Topology       *ResponseClientsGetClientDetailTopology       `json:"topology,omitempty"`       //
}
type ResponseClientsGetClientDetailDetail struct {
	ID                           string                                                 `json:"id,omitempty"`                           // Unique identifier representing a specific host
	ConnectionStatus             string                                                 `json:"connectionStatus,omitempty"`             // The client is connected, connecting, or disconnected
	Tracked                      string                                                 `json:"tracked,omitempty"`                      // Tracking status of this host
	HostType                     string                                                 `json:"hostType,omitempty"`                     // WIRED or WIRELESS
	UserID                       string                                                 `json:"userId,omitempty"`                       // The user ID of this host
	Duid                         string                                                 `json:"duid,omitempty"`                         // Device UID for MAC
	IDentifier                   string                                                 `json:"identifier,omitempty"`                   // The host's unique identifier, which is populated by and in order of userId, hostName, hostIpV4, hostIpV6, or hostMac
	HostName                     string                                                 `json:"hostName,omitempty"`                     // The hostname of the host
	HostOs                       string                                                 `json:"hostOs,omitempty"`                       // The OS of host
	HostVersion                  string                                                 `json:"hostVersion,omitempty"`                  // The version of OS of host
	SubType                      string                                                 `json:"subType,omitempty"`                      // The device type of host
	FirmwareVersion              string                                                 `json:"firmwareVersion,omitempty"`              // The firmware version of the host
	DeviceVendor                 string                                                 `json:"deviceVendor,omitempty"`                 // The device vendor string
	DeviceForm                   string                                                 `json:"deviceForm,omitempty"`                   // The device form of the host (e.g. Phone/Tablet)
	SalesCode                    string                                                 `json:"salesCode,omitempty"`                    // The Sales Code of the host
	CountryCode                  string                                                 `json:"countryCode,omitempty"`                  // The country code of the host
	LastUpdated                  *int                                                   `json:"lastUpdated,omitempty"`                  // Epoch/Unix time in milliseconds
	HealthScore                  *[]ResponseClientsGetClientDetailDetailHealthScore     `json:"healthScore,omitempty"`                  //
	HostMac                      string                                                 `json:"hostMac,omitempty"`                      // MAC address of the interface
	HostIPV4                     string                                                 `json:"hostIpV4,omitempty"`                     // IPv4 address of the interface
	HostIPV6                     []string                                               `json:"hostIpV6,omitempty"`                     // List of IPv6 addresses
	AuthType                     string                                                 `json:"authType,omitempty"`                     // Authentication mechanism of the client
	VLANID                       *int                                                   `json:"vlanId,omitempty"`                       // VLAN ID for the host
	L3VirtualNetwork             string                                                 `json:"l3VirtualNetwork,omitempty"`             // Comma separated Level 3 virtual network names
	L2VirtualNetwork             string                                                 `json:"l2VirtualNetwork,omitempty"`             // Comma separated Level 2 virtual network names
	Vnid                         *int                                                   `json:"vnid,omitempty"`                         // VNID of the host
	UpnID                        string                                                 `json:"upnId,omitempty"`                        // Registered UPN ID of the host
	UpnName                      string                                                 `json:"upnName,omitempty"`                      // Registered UPN name of the host
	SSID                         string                                                 `json:"ssid,omitempty"`                         // WLAN SSID to which the client is connected
	Frequency                    string                                                 `json:"frequency,omitempty"`                    // Frequency band to which the client is connected
	Channel                      string                                                 `json:"channel,omitempty"`                      // Channel to which the client is connected
	ApGroup                      string                                                 `json:"apGroup,omitempty"`                      // AP group to which the client belongs
	Sgt                          string                                                 `json:"sgt,omitempty"`                          // Security group tag
	Location                     string                                                 `json:"location,omitempty"`                     // Site location of client
	ClientConnection             string                                                 `json:"clientConnection,omitempty"`             // AP/Switch to which the client is connected
	ConnectedDevice              *[]ResponseClientsGetClientDetailDetailConnectedDevice `json:"connectedDevice,omitempty"`              //
	IssueCount                   *int                                                   `json:"issueCount,omitempty"`                   // Issue count for a device
	Rssi                         string                                                 `json:"rssi,omitempty"`                         // Min RSSI value for the client
	RssiThreshold                string                                                 `json:"rssiThreshold,omitempty"`                // RSSI threshold
	RssiIsInclude                string                                                 `json:"rssiIsInclude,omitempty"`                // RSSI include/exclude flag
	AvgRssi                      string                                                 `json:"avgRssi,omitempty"`                      // Average RSSI value for the client
	Snr                          string                                                 `json:"snr,omitempty"`                          // Min signal to noise ratio for the client
	SnrThreshold                 string                                                 `json:"snrThreshold,omitempty"`                 // SNR threshold
	SnrIsInclude                 string                                                 `json:"snrIsInclude,omitempty"`                 // SNR include/exclude flag
	AvgSnr                       string                                                 `json:"avgSnr,omitempty"`                       // Average signal to noise ratio for a client
	DataRate                     string                                                 `json:"dataRate,omitempty"`                     // MCS data rates for a client
	TxBytes                      string                                                 `json:"txBytes,omitempty"`                      // total transmitted bytes for a client
	RxBytes                      string                                                 `json:"rxBytes,omitempty"`                      // Total received bytes for a client
	DNSResponse                  string                                                 `json:"dnsResponse,omitempty"`                  // DNS response attempts for a client
	DNSRequest                   string                                                 `json:"dnsRequest,omitempty"`                   // DNS request attempts for a client
	Onboarding                   *ResponseClientsGetClientDetailDetailOnboarding        `json:"onboarding,omitempty"`                   //
	ClientType                   string                                                 `json:"clientType,omitempty"`                   // OLD or NEW
	OnboardingTime               *int                                                   `json:"onboardingTime,omitempty"`               // Epoch/Unix time in milliseconds
	Port                         string                                                 `json:"port,omitempty"`                         // switch port for client
	IosCapable                   *bool                                                  `json:"iosCapable,omitempty"`                   // IOS Capable
	Usage                        *float64                                               `json:"usage,omitempty"`                        // Usage of txBytes and rxBytes of client
	LinkSpeed                    *float64                                               `json:"linkSpeed,omitempty"`                    // The speed of wired client
	LinkThreshold                string                                                 `json:"linkThreshold,omitempty"`                // Link error threshold of wired client
	RemoteEndDuplexMode          string                                                 `json:"remoteEndDuplexMode,omitempty"`          // The remote end duplex mode of wired client
	TxLinkError                  *float64                                               `json:"txLinkError,omitempty"`                  // The error of tx link
	RxLinkError                  *float64                                               `json:"rxLinkError,omitempty"`                  // The error of rx link
	TxRate                       *float64                                               `json:"txRate,omitempty"`                       // The rate of tx
	RxRate                       *float64                                               `json:"rxRate,omitempty"`                       // The rate of rx
	RxRetryPct                   string                                                 `json:"rxRetryPct,omitempty"`                   // The retry count as percentage wrt to total rx packets
	VersionTime                  *int                                                   `json:"versionTime,omitempty"`                  // The metric modification time of the new version
	Dot11Protocol                string                                                 `json:"dot11Protocol,omitempty"`                // Description of dot11 protocol
	SlotID                       *int                                                   `json:"slotId,omitempty"`                       // Slot ID of AP which client is connected
	Dot11ProtocolCapability      string                                                 `json:"dot11ProtocolCapability,omitempty"`      // description of dot11 protocol capability
	PrivateMac                   *bool                                                  `json:"privateMac,omitempty"`                   // Private Mac
	DhcpServerIP                 string                                                 `json:"dhcpServerIp,omitempty"`                 // The DHCP server IP
	AAAServerIP                  string                                                 `json:"aaaServerIp,omitempty"`                  // The AAA server IP
	AAAServerTransaction         *int                                                   `json:"aaaServerTransaction,omitempty"`         // The number of AAA server transactions
	AAAServerFailedTransaction   *int                                                   `json:"aaaServerFailedTransaction,omitempty"`   // The number of failed AAA server transactions
	AAAServerSuccessTransaction  *int                                                   `json:"aaaServerSuccessTransaction,omitempty"`  // The number of successful AAA server transactions
	AAAServerLatency             *float64                                               `json:"aaaServerLatency,omitempty"`             // The AAA server latency
	AAAServerMABLatency          *float64                                               `json:"aaaServerMABLatency,omitempty"`          // The AAA server MAB latency
	AAAServerEApLatency          *float64                                               `json:"aaaServerEAPLatency,omitempty"`          // The AAA server EAP latency
	DhcpServerTransaction        *int                                                   `json:"dhcpServerTransaction,omitempty"`        // The number of DHCP server transactions
	DhcpServerFailedTransaction  *int                                                   `json:"dhcpServerFailedTransaction,omitempty"`  // The number of failed DHCP server transactions
	DhcpServerSuccessTransaction *int                                                   `json:"dhcpServerSuccessTransaction,omitempty"` // The number of successful DHCP server transactions
	DhcpServerLatency            *float64                                               `json:"dhcpServerLatency,omitempty"`            // The DHCP server latency
	DhcpServerDOLatency          *float64                                               `json:"dhcpServerDOLatency,omitempty"`          // The DHCP server DO latency
	DhcpServerRALatency          *float64                                               `json:"dhcpServerRALatency,omitempty"`          // The DHCP RA latency
	MaxRoamingDuration           string                                                 `json:"maxRoamingDuration,omitempty"`           // Max roaming duration for a client
	UpnOwner                     string                                                 `json:"upnOwner,omitempty"`                     // Owner of registered UPN name
	ConnectedUpn                 string                                                 `json:"connectedUpn,omitempty"`                 // connected UPN ID
	ConnectedUpnOwner            string                                                 `json:"connectedUpnOwner,omitempty"`            // Connected UPN owner
	ConnectedUpnID               string                                                 `json:"connectedUpnId,omitempty"`               // Connected UPN ID
	IsGuestUPNEndpoint           *bool                                                  `json:"isGuestUPNEndpoint,omitempty"`           // Whether it is guest UPN endpoint
	WlcName                      string                                                 `json:"wlcName,omitempty"`                      // The name of the connected wireless controller
	WlcUUID                      string                                                 `json:"wlcUuid,omitempty"`                      // UUID of the WLC the client connected to
	SessionDuration              string                                                 `json:"sessionDuration,omitempty"`              // Time duration the session took from run time to delete time
	IntelCapable                 *bool                                                  `json:"intelCapable,omitempty"`                 // Whether support Intel device analytics
	HwModel                      string                                                 `json:"hwModel,omitempty"`                      // Hardware model
	PowerType                    string                                                 `json:"powerType,omitempty"`                    // AC/DC voltage
	ModelName                    string                                                 `json:"modelName,omitempty"`                    // System model
	BridgeVMMode                 string                                                 `json:"bridgeVMMode,omitempty"`                 // Bridge VM mode
	DhcpNakIP                    string                                                 `json:"dhcpNakIp,omitempty"`                    // DHCP NAK IP
	DhcpDeclineIP                string                                                 `json:"dhcpDeclineIp,omitempty"`                // DHCP decline IP
	PortDescription              string                                                 `json:"portDescription,omitempty"`              // Port desctiption of wired client
	LatencyVoice                 *float64                                               `json:"latencyVoice,omitempty"`                 // Voice latency
	LatencyVideo                 *float64                                               `json:"latencyVideo,omitempty"`                 // Video latency
	LatencyBg                    *float64                                               `json:"latencyBg,omitempty"`                    // Background latency
	LatencyBe                    *float64                                               `json:"latencyBe,omitempty"`                    // Best-effort latency
	TrustScore                   string                                                 `json:"trustScore,omitempty"`                   // Trust score of Client received from EndPoint Analytics
	TrustDetails                 string                                                 `json:"trustDetails,omitempty"`                 // Trust details explaining reason for corresponding Trust score
}
type ResponseClientsGetClientDetailDetailHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Type of device health
	Reason     string `json:"reason,omitempty"`     // Reason for the health score value
	Score      *int   `json:"score,omitempty"`      // health score of client device in the range of 1 to 10. Value 0 for a client represents an IDLE client
}
type ResponseClientsGetClientDetailDetailConnectedDevice struct {
	Type      string `json:"type,omitempty"`       // Type of device (AP or SWITCH)
	Name      string `json:"name,omitempty"`       // Name of the device
	Mac       string `json:"mac,omitempty"`        // MAC address of the access point
	ID        string `json:"id,omitempty"`         // Unique identifier of the device
	IPaddress string `json:"ip address,omitempty"` // Management IP address of the connected device.  (deprecated soon in favor of 'mgmtIp')
	MgmtIP    string `json:"mgmtIp,omitempty"`     // The IP address of the connected device
	Band      string `json:"band,omitempty"`       // Band of the AP
	Mode      string `json:"mode,omitempty"`       // The mode of the access point
}
type ResponseClientsGetClientDetailDetailOnboarding struct {
	AverageRunDuration   string   `json:"averageRunDuration,omitempty"`   // Average run Duration for a client
	MaxRunDuration       string   `json:"maxRunDuration,omitempty"`       // Max run duration for a client
	AverageAssocDuration string   `json:"averageAssocDuration,omitempty"` // Average association duration for a client
	MaxAssocDuration     string   `json:"maxAssocDuration,omitempty"`     // Max association duration for a client
	AverageAuthDuration  string   `json:"averageAuthDuration,omitempty"`  // Average auth duration for a client
	MaxAuthDuration      string   `json:"maxAuthDuration,omitempty"`      // Max auth duration for a client
	AverageDhcpDuration  string   `json:"averageDhcpDuration,omitempty"`  // Average DHCP duration for a client
	MaxDhcpDuration      string   `json:"maxDhcpDuration,omitempty"`      // Max DHCP duration for a client
	AAAServerIP          string   `json:"aaaServerIp,omitempty"`          // AAA server IP for a client
	DhcpServerIP         string   `json:"dhcpServerIp,omitempty"`         // DHCP server IP for a client
	AuthDoneTime         *int     `json:"authDoneTime,omitempty"`         // Epoch/Unix time in milliseconds
	AssocDoneTime        *int     `json:"assocDoneTime,omitempty"`        // Epoch/Unix time in milliseconds
	DhcpDoneTime         *int     `json:"dhcpDoneTime,omitempty"`         // Epoch/Unix time in milliseconds
	AssocRootcauseList   []string `json:"assocRootcauseList,omitempty"`   // Root cause list of ASSOC failure category
	AAARootcauseList     []string `json:"aaaRootcauseList,omitempty"`     // Root cause list of AAA failure category
	DhcpRootcauseList    []string `json:"dhcpRootcauseList,omitempty"`    // Root cause list of DHCP failure category
	OtherRootcauseList   []string `json:"otherRootcauseList,omitempty"`   // Root cause list of other failure category
	LatestRootCauseList  []string `json:"latestRootCauseList,omitempty"`  // Root cause list of latest root cause category
}
type ResponseClientsGetClientDetailConnectionInfo struct {
	HostType      string `json:"hostType,omitempty"`      // Host Type - WIRELESS or WIRED
	NwDeviceName  string `json:"nwDeviceName,omitempty"`  // Name of the network device it is connected to. In case of wireless, it would be an AccessPoint
	NwDeviceMac   string `json:"nwDeviceMac,omitempty"`   // Device MAC address
	Protocol      string `json:"protocol,omitempty"`      // Connection Protocol used. This information is present for wireless hosts only
	Band          string `json:"band,omitempty"`          // The band at which the host is connected. This information is present for wireless hosts only
	SpatialStream string `json:"spatialStream,omitempty"` // The spatial stream of host. This information is present for wireless hosts only
	Channel       string `json:"channel,omitempty"`       // The channel used by the host. This information is present for wireless hosts only
	ChannelWidth  string `json:"channelWidth,omitempty"`  // The channel width used by the host. This information is present for wireless hosts only
	Wmm           string `json:"wmm,omitempty"`           // The wmm of the host. This information is present for wireless hosts only
	Uapsd         string `json:"uapsd,omitempty"`         // The UAPSD of the host. This information is present for wireless hosts only
	Timestamp     *int   `json:"timestamp,omitempty"`     // Epoch/Unix time in milliseconds
}
type ResponseClientsGetClientDetailTopology struct {
	Nodes *[]ResponseClientsGetClientDetailTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseClientsGetClientDetailTopologyLinks `json:"links,omitempty"` //
}
type ResponseClientsGetClientDetailTopologyNodes struct {
	Role            string   `json:"role,omitempty"`            // Device role
	Name            string   `json:"name,omitempty"`            // Device name
	ID              string   `json:"id,omitempty"`              // User ID, hostname, IP address, or MAC address
	Description     string   `json:"description,omitempty"`     // Description of the topology node
	DeviceType      string   `json:"deviceType,omitempty"`      // Device type
	PlatformID      string   `json:"platformId,omitempty"`      // Device platform ID
	Family          string   `json:"family,omitempty"`          // Device family
	IP              string   `json:"ip,omitempty"`              // Device IP
	IPv6            []string `json:"ipv6,omitempty"`            // Device IPv6
	SoftwareVersion string   `json:"softwareVersion,omitempty"` // Device software version
	UserID          string   `json:"userId,omitempty"`          // User ID
	NodeType        string   `json:"nodeType,omitempty"`        // Node type
	RadioFrequency  string   `json:"radioFrequency,omitempty"`  // Radio frequency
	Clients         *float64 `json:"clients,omitempty"`         // Number of clients
	Count           *int     `json:"count,omitempty"`           // Count
	HealthScore     *float64 `json:"healthScore,omitempty"`     // Device health score
	Level           *float64 `json:"level,omitempty"`           // Level in the topology
	FabricGroup     string   `json:"fabricGroup,omitempty"`     // Fabric Group
	FabricRole      []string `json:"fabricRole,omitempty"`      // Fabric Role
	ConnectedDevice string   `json:"connectedDevice,omitempty"` // Connected Device
	StackType       string   `json:"stackType,omitempty"`       // Stack Type
}
type ResponseClientsGetClientDetailTopologyLinks struct {
	Source              string                                                         `json:"source,omitempty"`              // Edge line starting node
	LinkStatus          string                                                         `json:"linkStatus,omitempty"`          // Link status of the link
	SourceLinkStatus    string                                                         `json:"sourceLinkStatus,omitempty"`    // The status of the link
	TargetLinkStatus    string                                                         `json:"targetLinkStatus,omitempty"`    // The status of the link
	Label               []string                                                       `json:"label,omitempty"`               // The details of the edge
	Target              string                                                         `json:"target,omitempty"`              // End node of the edge line
	ID                  string                                                         `json:"id,omitempty"`                  // Identifier of the node
	PortUtilization     *float64                                                       `json:"portUtilization,omitempty"`     // Port utilization
	SourceInterfaceName string                                                         `json:"sourceInterfaceName,omitempty"` // The interface name of the source
	TargetInterfaceName string                                                         `json:"targetInterfaceName,omitempty"` // The interface name of the target
	SourceDuplexInfo    string                                                         `json:"sourceDuplexInfo,omitempty"`    // The duplex info of the source
	TargetDuplexInfo    string                                                         `json:"targetDuplexInfo,omitempty"`    // The duplex info of the target
	SourcePortMode      string                                                         `json:"sourcePortMode,omitempty"`      // The port mode of the source
	TargetPortMode      string                                                         `json:"targetPortMode,omitempty"`      // The port mode of the target
	SourceAdminStatus   string                                                         `json:"sourceAdminStatus,omitempty"`   // The admin status of the source
	TargetAdminStatus   string                                                         `json:"targetAdminStatus,omitempty"`   // The admin status of the target
	ApRadioAdminStatus  string                                                         `json:"apRadioAdminStatus,omitempty"`  // The admin status of the radio
	ApRadioOperStatus   string                                                         `json:"apRadioOperStatus,omitempty"`   // The oper status of the radio
	SourcePortVLANInfo  string                                                         `json:"sourcePortVLANInfo,omitempty"`  // List of VLANs configured on the source port
	TargetPortVLANInfo  string                                                         `json:"targetPortVLANInfo,omitempty"`  // List of VLANs configured on the target port
	InterfaceDetails    *[]ResponseClientsGetClientDetailTopologyLinksInterfaceDetails `json:"interfaceDetails,omitempty"`    //
}
type ResponseClientsGetClientDetailTopologyLinksInterfaceDetails struct {
	ClientMacAddress       string `json:"clientMacAddress,omitempty"`       // The MAC address of the client device
	ConnectedDeviceIntName string `json:"connectedDeviceIntName,omitempty"` // The interface name of the network device
	Duplex                 string `json:"duplex,omitempty"`                 // The duplex info of the network device interface
	PortMode               string `json:"portMode,omitempty"`               // The port mode info of network device interface
	AdminStatus            string `json:"adminStatus,omitempty"`            // The admin status of network device interface
}
type ResponseClientsGetClientEnrichmentDetails []ResponseItemClientsGetClientEnrichmentDetails // Array of ResponseClientsGetClientEnrichmentDetails
type ResponseItemClientsGetClientEnrichmentDetails struct {
	UserDetails     *ResponseItemClientsGetClientEnrichmentDetailsUserDetails       `json:"userDetails,omitempty"`     //
	ConnectedDevice *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDevice `json:"connectedDevice,omitempty"` //
	IssueDetails    *ResponseItemClientsGetClientEnrichmentDetailsIssueDetails      `json:"issueDetails,omitempty"`    //
}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetails struct {
	ID               string                                                                     `json:"id,omitempty"`               // Id
	ConnectionStatus string                                                                     `json:"connectionStatus,omitempty"` // Connection Status
	HostType         string                                                                     `json:"hostType,omitempty"`         // Host Type
	UserID           string                                                                     `json:"userId,omitempty"`           // User Id
	HostName         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostName          `json:"hostName,omitempty"`         // Host Name
	HostOs           *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostOs            `json:"hostOs,omitempty"`           // Host Os
	HostVersion      *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostVersion       `json:"hostVersion,omitempty"`      // Host Version
	SubType          *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSubType           `json:"subType,omitempty"`          // Sub Type
	LastUpdated      *int                                                                       `json:"lastUpdated,omitempty"`      // Last Updated
	HealthScore      *[]ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHealthScore     `json:"healthScore,omitempty"`      //
	HostMac          string                                                                     `json:"hostMac,omitempty"`          // Host Mac
	HostIPV4         string                                                                     `json:"hostIpV4,omitempty"`         // Host Ip V4
	HostIPV6         *[]ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostIPV6        `json:"hostIpV6,omitempty"`         // Host Ip V6
	AuthType         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsAuthType          `json:"authType,omitempty"`         // Auth Type
	VLANID           string                                                                     `json:"vlanId,omitempty"`           // Vlan Id
	SSID             *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSSID              `json:"ssid,omitempty"`             // Ssid
	Location         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsLocation          `json:"location,omitempty"`         // Location
	ClientConnection string                                                                     `json:"clientConnection,omitempty"` // Client Connection
	ConnectedDevice  *[]ResponseItemClientsGetClientEnrichmentDetailsUserDetailsConnectedDevice `json:"connectedDevice,omitempty"`  // Connected Device
	IssueCount       *float64                                                                   `json:"issueCount,omitempty"`       // Issue Count
	Rssi             *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsRssi              `json:"rssi,omitempty"`             // Rssi
	Snr              *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSnr               `json:"snr,omitempty"`              // Snr
	DataRate         *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsDataRate          `json:"dataRate,omitempty"`         // Data Rate
	Port             *ResponseItemClientsGetClientEnrichmentDetailsUserDetailsPort              `json:"port,omitempty"`             // Port
}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostName interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostOs interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostVersion interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSubType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Health Type
	Reason     string `json:"reason,omitempty"`     // Reason
	Score      *int   `json:"score,omitempty"`      // Score
}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsHostIPV6 interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsAuthType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSSID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsLocation interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsConnectedDevice interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsRssi interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsSnr interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsDataRate interface{}
type ResponseItemClientsGetClientEnrichmentDetailsUserDetailsPort interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDevice struct {
	DeviceDetails *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetails struct {
	Family                    string                                                                                       `json:"family,omitempty"`                    // Family
	Type                      string                                                                                       `json:"type,omitempty"`                      // Type
	Location                  *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocation           `json:"location,omitempty"`                  // Location
	ErrorCode                 string                                                                                       `json:"errorCode,omitempty"`                 // Error Code
	MacAddress                string                                                                                       `json:"macAddress,omitempty"`                // Mac Address
	Role                      string                                                                                       `json:"role,omitempty"`                      // Role
	ApManagerInterfaceIP      string                                                                                       `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                                                       `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsBootDateTime       `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                                                       `json:"collectionStatus,omitempty"`          // Collection Status
	InterfaceCount            *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsInterfaceCount     `json:"interfaceCount,omitempty"`            // Interface Count
	LineCardCount             *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardCount      `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardID         `json:"lineCardId,omitempty"`                // Line Card Id
	ManagementIPAddress       string                                                                                       `json:"managementIpAddress,omitempty"`       // Management Ip Address
	MemorySize                string                                                                                       `json:"memorySize,omitempty"`                // Memory Size
	PlatformID                string                                                                                       `json:"platformId,omitempty"`                // Platform Id
	ReachabilityFailureReason string                                                                                       `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                                                       `json:"reachabilityStatus,omitempty"`        // Reachability Status
	SNMPContact               string                                                                                       `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                                                       `json:"snmpLocation,omitempty"`              // Snmp Location
	TunnelUDPPort             string                                                                                       `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	WaasDeviceMode            *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	Series                    string                                                                                       `json:"series,omitempty"`                    // Series
	InventoryStatusDetail     string                                                                                       `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	CollectionInterval        string                                                                                       `json:"collectionInterval,omitempty"`        // Collection Interval
	SerialNumber              string                                                                                       `json:"serialNumber,omitempty"`              // Serial Number
	SoftwareVersion           string                                                                                       `json:"softwareVersion,omitempty"`           // Software Version
	RoleSource                string                                                                                       `json:"roleSource,omitempty"`                // Role Source
	Hostname                  string                                                                                       `json:"hostname,omitempty"`                  // Hostname
	UpTime                    string                                                                                       `json:"upTime,omitempty"`                    // Up Time
	LastUpdateTime            *int                                                                                         `json:"lastUpdateTime,omitempty"`            // Last Update Time
	ErrorDescription          *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription   `json:"errorDescription,omitempty"`          // Error Description
	LocationName              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName       `json:"locationName,omitempty"`              // Location Name
	TagCount                  string                                                                                       `json:"tagCount,omitempty"`                  // Tag Count
	LastUpdated               string                                                                                       `json:"lastUpdated,omitempty"`               // Last Updated
	InstanceUUID              string                                                                                       `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                                                       `json:"id,omitempty"`                        // Id
	NeighborTopology          *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
	Cisco360View              string                                                                                       `json:"cisco360view,omitempty"`              // Cisco360view
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocation interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsBootDateTime interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsInterfaceCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLineCardID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology struct {
	Nodes *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodes `json:"nodes,omitempty"` //
	Links *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinks `json:"links,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodes struct {
	Role            string                                                                                                         `json:"role,omitempty"`            // Role
	Name            string                                                                                                         `json:"name,omitempty"`            // Name
	ID              string                                                                                                         `json:"id,omitempty"`              // Id
	Description     string                                                                                                         `json:"description,omitempty"`     // Description
	DeviceType      *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType      `json:"deviceType,omitempty"`      // Device Type
	PlatformID      *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID      `json:"platformId,omitempty"`      // Platform Id
	Family          *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily          `json:"family,omitempty"`          // Family
	IP              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP              `json:"ip,omitempty"`              // Ip
	SoftwareVersion *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion `json:"softwareVersion,omitempty"` // Software Version
	UserID          *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID          `json:"userId,omitempty"`          // User Id
	NodeType        *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType        `json:"nodeType,omitempty"`        // Node Type
	RadioFrequency  *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency  `json:"radioFrequency,omitempty"`  // Radio Frequency
	Clients         *float64                                                                                                       `json:"clients,omitempty"`         // Clients
	Count           *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount           `json:"count,omitempty"`           // Count
	HealthScore     *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore     `json:"healthScore,omitempty"`     // Health Score
	Level           *float64                                                                                                       `json:"level,omitempty"`           // Level
	FabricGroup     *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup     `json:"fabricGroup,omitempty"`     // Fabric Group
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinks struct {
	Source          string                                                                                                         `json:"source,omitempty"`          // Source
	LinkStatus      string                                                                                                         `json:"linkStatus,omitempty"`      // Link Status
	Label           *[]ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel         `json:"label,omitempty"`           // Label
	Target          string                                                                                                         `json:"target,omitempty"`          // Target
	ID              *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksID              `json:"id,omitempty"`              // Id
	PortUtilization *ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization `json:"portUtilization,omitempty"` // Port Utilization
}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksID interface{}
type ResponseItemClientsGetClientEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization interface{}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetails struct {
	Issue *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssue `json:"issue,omitempty"` //
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssue struct {
	IssueID          string                                                                            `json:"issueId,omitempty"`          // Issue Id
	IssueSource      string                                                                            `json:"issueSource,omitempty"`      // Issue Source
	IssueCategory    string                                                                            `json:"issueCategory,omitempty"`    // Issue Category
	IssueName        string                                                                            `json:"issueName,omitempty"`        // Issue Name
	IssueDescription string                                                                            `json:"issueDescription,omitempty"` // Issue Description
	IssueEntity      string                                                                            `json:"issueEntity,omitempty"`      // Issue Entity
	IssueEntityValue string                                                                            `json:"issueEntityValue,omitempty"` // Issue Entity Value
	IssueSeverity    string                                                                            `json:"issueSeverity,omitempty"`    // Issue Severity
	IssuePriority    string                                                                            `json:"issuePriority,omitempty"`    // Issue Priority
	IssueSummary     string                                                                            `json:"issueSummary,omitempty"`     // Issue Summary
	IssueTimestamp   *int                                                                              `json:"issueTimestamp,omitempty"`   // Issue Timestamp
	SuggestedActions *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
	ImpactedHosts    *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    //
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActions struct {
	Message string                                                                                 `json:"message,omitempty"` // Message
	Steps   *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps interface{}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHosts struct {
	HostType           string                                                                               `json:"hostType,omitempty"`           // Host Type
	HostName           string                                                                               `json:"hostName,omitempty"`           // Host Name
	HostOs             string                                                                               `json:"hostOs,omitempty"`             // Host Os
	SSID               string                                                                               `json:"ssid,omitempty"`               // Ssid
	ConnectedInterface string                                                                               `json:"connectedInterface,omitempty"` // Connected Interface
	MacAddress         string                                                                               `json:"macAddress,omitempty"`         // Mac Address
	FailedAttempts     *int                                                                                 `json:"failedAttempts,omitempty"`     // Failed Attempts
	Location           *ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocation `json:"location,omitempty"`           //
	Timestamp          *int                                                                                 `json:"timestamp,omitempty"`          // Timestamp
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocation struct {
	SiteID      string                                                                                            `json:"siteId,omitempty"`      // Site Id
	SiteType    string                                                                                            `json:"siteType,omitempty"`    // Site Type
	Area        string                                                                                            `json:"area,omitempty"`        // Area
	Building    string                                                                                            `json:"building,omitempty"`    // Building
	Floor       *ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationFloor         `json:"floor,omitempty"`       // Floor
	ApsImpacted *[]ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationApsImpacted `json:"apsImpacted,omitempty"` // Aps Impacted
}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationFloor interface{}
type ResponseItemClientsGetClientEnrichmentDetailsIssueDetailsIssueImpactedHostsLocationApsImpacted interface{}
type ResponseClientsGetOverallClientHealth struct {
	Version  string                                           `json:"version,omitempty"`  // Response output version
	Response *[]ResponseClientsGetOverallClientHealthResponse `json:"response,omitempty"` //
}
type ResponseClientsGetOverallClientHealthResponse struct {
	SiteID      string                                                      `json:"siteId,omitempty"`      // Site UUID or 'global'
	ScoreDetail *[]ResponseClientsGetOverallClientHealthResponseScoreDetail `json:"scoreDetail,omitempty"` //
}
type ResponseClientsGetOverallClientHealthResponseScoreDetail struct {
	ScoreCategory                  *ResponseClientsGetOverallClientHealthResponseScoreDetailScoreCategory `json:"scoreCategory,omitempty"`                  //
	ScoreValue                     *float64                                                               `json:"scoreValue,omitempty"`                     // Percentage of GOOD health score in the category.  (-1 means not applicable for the category)
	ClientCount                    *int                                                                   `json:"clientCount,omitempty"`                    // Total client count
	ClientUniqueCount              *int                                                                   `json:"clientUniqueCount,omitempty"`              // Total unique client count
	MaintenanceAffectedClientCount *int                                                                   `json:"maintenanceAffectedClientCount,omitempty"` // Total client count affected by maintenance
	RandomMacCount                 *int                                                                   `json:"randomMacCount,omitempty"`                 // Total client count with random MAC count
	DuidCount                      *int                                                                   `json:"duidCount,omitempty"`                      // Device UUID count
	Starttime                      *int                                                                   `json:"starttime,omitempty"`                      // UTC timestamp of data start time
	Endtime                        *int                                                                   `json:"endtime,omitempty"`                        // UTC timestamp of data end time
	ConnectedToUdnCount            *int                                                                   `json:"connectedToUdnCount,omitempty"`            // Total connected to UDN count
	UnconnectedToUdnCount          *int                                                                   `json:"unconnectedToUdnCount,omitempty"`          // Total unconnected to UDN count
	ScoreList                      *[]ResponseClientsGetOverallClientHealthResponseScoreDetailScoreList   `json:"scoreList,omitempty"`                      //
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Health score category
	Value         string `json:"value,omitempty"`         // Health score category value
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreList struct {
	ScoreCategory                  *ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreCategory `json:"scoreCategory,omitempty"`                  //
	ScoreValue                     *float64                                                                        `json:"scoreValue,omitempty"`                     // Percentage of GOOD health score in the category.  (-1 means not applicable for the category)
	ClientCount                    *int                                                                            `json:"clientCount,omitempty"`                    // Total client count
	ClientUniqueCount              *int                                                                            `json:"clientUniqueCount,omitempty"`              // Total unique client count
	MaintenanceAffectedClientCount *int                                                                            `json:"maintenanceAffectedClientCount,omitempty"` // Total client count affected by maintenance
	RandomMacCount                 *int                                                                            `json:"randomMacCount,omitempty"`                 // Total client count with random MAC count
	DuidCount                      *int                                                                            `json:"duidCount,omitempty"`                      // Device UUID count
	Starttime                      *int                                                                            `json:"starttime,omitempty"`                      // UTC timestamp of data start time
	Endtime                        *int                                                                            `json:"endtime,omitempty"`                        // UTC timestamp of data end time
	ConnectedToUdnCount            *int                                                                            `json:"connectedToUdnCount,omitempty"`            // Total connected to UDN count
	UnconnectedToUdnCount          *int                                                                            `json:"unconnectedToUdnCount,omitempty"`          // Total unconnected to UDN count
}
type ResponseClientsGetOverallClientHealthResponseScoreDetailScoreListScoreCategory struct {
	ScoreCategory string `json:"scoreCategory,omitempty"` // Category of the overall health score
	Value         string `json:"value,omitempty"`         // Health score category value
}
type ResponseClientsClientProximity struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes struct {
	StartTime           *int                                                                                                                        `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                                                        `json:"endTime,omitempty"`             // End Time
	Views               []string                                                                                                                    `json:"views,omitempty"`               // Views
	Attributes          []string                                                                                                                    `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPage struct {
	Limit  *int                                                                                                               `json:"limit,omitempty"`  // Limit
	Offset *int                                                                                                               `json:"offset,omitempty"` // Offset
	SortBy *[]RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFilters struct {
	StartTime *int                                                                        `json:"startTime,omitempty"` // Start Time
	EndTime   *int                                                                        `json:"endTime,omitempty"`   // End Time
	Filters   *[]RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersFilters `json:"filters,omitempty"`   //
}
type RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFiltersFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClients struct {
	StartTime           *int                                                                              `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                              `json:"endTime,omitempty"`             // End Time
	GroupBy             []string                                                                          `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                          `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsPage                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsPage struct {
	Limit  *int                                                                     `json:"limit,omitempty"`  // Limit
	Cursor string                                                                   `json:"cursor,omitempty"` // Cursor
	SortBy *[]RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsRetrievesSummaryAnalyticsDataRelatedToClientsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClients struct {
	StartTime           *int                                                                              `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                              `json:"endTime,omitempty"`             // End Time
	TopN                *int                                                                              `json:"topN,omitempty"`                // Top N
	GroupBy             []string                                                                          `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                          `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPage                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPage struct {
	Limit  *int                                                                     `json:"limit,omitempty"`  // Limit
	Cursor string                                                                   `json:"cursor,omitempty"` // Cursor
	SortBy *[]RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClientsPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClients struct {
	StartTime           *int                                                                               `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                               `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                                             `json:"trendInterval,omitempty"`       // Trend Interval
	GroupBy             []string                                                                           `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                           `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsPage                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClientsPage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime struct {
	StartTime           *int                                                                                             `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                                             `json:"endTime,omitempty"`             // End Time
	TrendInterval       string                                                                                           `json:"trendInterval,omitempty"`       // Trend Interval
	GroupBy             []string                                                                                         `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                                                         `json:"attributes,omitempty"`          // Attributes
	Filters             *[]RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeFilters             `json:"filters,omitempty"`             //
	AggregateAttributes *[]RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimePage                  `json:"page,omitempty"`                //
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeFilters struct {
	Key      string `json:"key,omitempty"`      // Key
	Operator string `json:"operator,omitempty"` // Operator
	Value    *int   `json:"value,omitempty"`    // Value
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTimePage struct {
	Limit         *int   `json:"limit,omitempty"`         // Limit
	Cursor        string `json:"cursor,omitempty"`        // Cursor
	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestClientsQueryClientsEnergy struct {
	StartTime           *int                                                   `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                   `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestClientsQueryClientsEnergyFilters             `json:"filters,omitempty"`             //
	Views               []string                                               `json:"views,omitempty"`               // Views
	Attributes          []string                                               `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestClientsQueryClientsEnergyAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsQueryClientsEnergyPage                  `json:"page,omitempty"`                //
}
type RequestClientsQueryClientsEnergyFilters struct {
	LogicalOperator string                                            `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestClientsQueryClientsEnergyFiltersFilters `json:"filters,omitempty"`         //
}
type RequestClientsQueryClientsEnergyFiltersFilters struct {
	Key      string   `json:"key,omitempty"`      // Key
	Operator string   `json:"operator,omitempty"` // Operator
	Value    []string `json:"value,omitempty"`    // Value
}
type RequestClientsQueryClientsEnergyAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsQueryClientsEnergyPage struct {
	Limit  *int                                          `json:"limit,omitempty"`  // Limit
	Cursor string                                        `json:"cursor,omitempty"` // Cursor
	SortBy *[]RequestClientsQueryClientsEnergyPageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsQueryClientsEnergyPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsCountClientsEnergyFromQuery struct {
	StartTime           *int                                                            `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                            `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestClientsCountClientsEnergyFromQueryFilters             `json:"filters,omitempty"`             //
	Views               []string                                                        `json:"views,omitempty"`               // Views
	Attributes          []string                                                        `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestClientsCountClientsEnergyFromQueryAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestClientsCountClientsEnergyFromQueryPage                  `json:"page,omitempty"`                //
}
type RequestClientsCountClientsEnergyFromQueryFilters struct {
	LogicalOperator string                                                     `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestClientsCountClientsEnergyFromQueryFiltersFilters `json:"filters,omitempty"`         //
}
type RequestClientsCountClientsEnergyFromQueryFiltersFilters struct {
	Key      string   `json:"key,omitempty"`      // Key
	Operator string   `json:"operator,omitempty"` // Operator
	Value    []string `json:"value,omitempty"`    // Value
}
type RequestClientsCountClientsEnergyFromQueryAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestClientsCountClientsEnergyFromQueryPage struct {
	Limit  *int                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                   `json:"offset,omitempty"` // Offset
	SortBy *[]RequestClientsCountClientsEnergyFromQueryPageSortBy `json:"sortBy,omitempty"` //
}
type RequestClientsCountClientsEnergyFromQueryPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}

//RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities Retrieves the list of clients, while also offering basic filtering and sorting capabilities. - ecb7-ab7e-47eb-8793
/* Retrieves the list of clients, while also offering basic filtering and sorting capabilities. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams Custom header parameters
@param RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-clients-while-also-offering-basic-filtering-and-sorting-capabilities
*/
func (s *ClientsService) RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams *RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams, RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams *RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams) (*ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities, *resty.Response, error) {
	path := "/dna/data/api/v1/clients"

	queryString, _ := query.Values(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams != nil {

		if RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities(RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesHeaderParams, RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilitiesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities")
	}

	result := response.Result().(*ResponseClientsRetrievesTheListOfClientsWhileAlsoOfferingBasicFilteringAndSortingCapabilities)
	return result, response, err

}

//RetrievesTheTotalCountOfClientsByApplyingBasicFiltering Retrieves the total count of clients by applying basic filtering - a486-4bef-4cab-9548
/* Retrieves the number of clients by applying basic filtering. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams Custom header parameters
@param RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-count-of-clients-by-applying-basic-filtering
*/
func (s *ClientsService) RetrievesTheTotalCountOfClientsByApplyingBasicFiltering(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams *RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams, RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams *RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams) (*ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFiltering, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/count"

	queryString, _ := query.Values(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams != nil {

		if RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFiltering{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalCountOfClientsByApplyingBasicFiltering(RetrievesTheTotalCountOfClientsByApplyingBasicFilteringHeaderParams, RetrievesTheTotalCountOfClientsByApplyingBasicFilteringQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalCountOfClientsByApplyingBasicFiltering")
	}

	result := response.Result().(*ResponseClientsRetrievesTheTotalCountOfClientsByApplyingBasicFiltering)
	return result, response, err

}

//RetrievesSpecificClientInformationMatchingTheMacaddress Retrieves specific client information matching the MAC address. - 829f-8b65-4779-9069
/* Retrieves specific client information matching the MAC address. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param id id path parameter. id is the client mac address. It can be specified is any notational conventions 01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive

@param RetrievesSpecificClientInformationMatchingTheMACAddressHeaderParams Custom header parameters
@param RetrievesSpecificClientInformationMatchingTheMACAddressQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-client-information-matching-the-macaddress
*/
func (s *ClientsService) RetrievesSpecificClientInformationMatchingTheMacaddress(id string, RetrievesSpecificClientInformationMatchingTheMACAddressHeaderParams *RetrievesSpecificClientInformationMatchingTheMacaddressHeaderParams, RetrievesSpecificClientInformationMatchingTheMACAddressQueryParams *RetrievesSpecificClientInformationMatchingTheMacaddressQueryParams) (*ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddress, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesSpecificClientInformationMatchingTheMACAddressQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificClientInformationMatchingTheMACAddressHeaderParams != nil {

		if RetrievesSpecificClientInformationMatchingTheMACAddressHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificClientInformationMatchingTheMACAddressHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddress{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificClientInformationMatchingTheMacaddress(id, RetrievesSpecificClientInformationMatchingTheMACAddressHeaderParams, RetrievesSpecificClientInformationMatchingTheMACAddressQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesSpecificClientInformationMatchingTheMacaddress")
	}

	result := response.Result().(*ResponseClientsRetrievesSpecificClientInformationMatchingTheMacaddress)
	return result, response, err

}

//GetClientsEnergy Get clients energy - a2a3-4b20-4fe8-9bfd
/* Retrieves a list of client devices with energy data based on the specified query parameters. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param GetClientsEnergyHeaderParams Custom header parameters
@param GetClientsEnergyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-clients-energy
*/
func (s *ClientsService) GetClientsEnergy(GetClientsEnergyHeaderParams *GetClientsEnergyHeaderParams, GetClientsEnergyQueryParams *GetClientsEnergyQueryParams) (*ResponseClientsGetClientsEnergy, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/clients"

	queryString, _ := query.Values(GetClientsEnergyQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetClientsEnergyHeaderParams != nil {

		if GetClientsEnergyHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetClientsEnergyHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetClientsEnergy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetClientsEnergy(GetClientsEnergyHeaderParams, GetClientsEnergyQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetClientsEnergy")
	}

	result := response.Result().(*ResponseClientsGetClientsEnergy)
	return result, response, err

}

//CountClientsEnergy Count clients energy - b28b-2962-403b-b688
/* Retrieves the total count of client devices that provide energy data, filtered according to the specified query parameters. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param CountClientsEnergyHeaderParams Custom header parameters
@param CountClientsEnergyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-clients-energy
*/
func (s *ClientsService) CountClientsEnergy(CountClientsEnergyHeaderParams *CountClientsEnergyHeaderParams, CountClientsEnergyQueryParams *CountClientsEnergyQueryParams) (*ResponseClientsCountClientsEnergy, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/clients/count"

	queryString, _ := query.Values(CountClientsEnergyQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountClientsEnergyHeaderParams != nil {

		if CountClientsEnergyHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountClientsEnergyHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsCountClientsEnergy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountClientsEnergy(CountClientsEnergyHeaderParams, CountClientsEnergyQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountClientsEnergy")
	}

	result := response.Result().(*ResponseClientsCountClientsEnergy)
	return result, response, err

}

//GetClientEnergyByID Get client energy by ID - b69f-fad9-45fa-94f4
/* Retrieves client device energy data for a specified time range based on the client ID. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param id id path parameter. Mac address of a client device (e.g., 54:9F:C6:43:FF:80). It can be specified is any notational conventions

  01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive.

@param GetClientEnergyByIDHeaderParams Custom header parameters
@param GetClientEnergyByIDQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-client-energy-by-id
*/
func (s *ClientsService) GetClientEnergyByID(id string, GetClientEnergyByIDHeaderParams *GetClientEnergyByIDHeaderParams, GetClientEnergyByIDQueryParams *GetClientEnergyByIDQueryParams) (*ResponseClientsGetClientEnergyByID, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/clients/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetClientEnergyByIDQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetClientEnergyByIDHeaderParams != nil {

		if GetClientEnergyByIDHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetClientEnergyByIDHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetClientEnergyByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetClientEnergyByID(id, GetClientEnergyByIDHeaderParams, GetClientEnergyByIDQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetClientEnergyById")
	}

	result := response.Result().(*ResponseClientsGetClientEnergyByID)
	return result, response, err

}

//GetClientDetail Get Client Detail - 1980-1996-4389-9d65
/* Returns detailed Client information retrieved by Mac Address for any given point of time.


@param GetClientDetailQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-client-detail
*/
func (s *ClientsService) GetClientDetail(GetClientDetailQueryParams *GetClientDetailQueryParams) (*ResponseClientsGetClientDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-detail"

	queryString, _ := query.Values(GetClientDetailQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetClientDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetClientDetail(GetClientDetailQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetClientDetail")
	}

	result := response.Result().(*ResponseClientsGetClientDetail)
	return result, response, err

}

//GetClientEnrichmentDetails Get Client Enrichment Details - b199-685d-4d08-9a67
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user, the devices that the user is connected to and the assurance issues that the user is impacted by


@param GetClientEnrichmentDetailsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-client-enrichment-details
*/
func (s *ClientsService) GetClientEnrichmentDetails(GetClientEnrichmentDetailsHeaderParams *GetClientEnrichmentDetailsHeaderParams) (*ResponseClientsGetClientEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetClientEnrichmentDetailsHeaderParams != nil {

		if GetClientEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetClientEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetClientEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetClientEnrichmentDetailsHeaderParams.EntityValue)
		}

		if GetClientEnrichmentDetailsHeaderParams.IssueCategory != "" {
			clientRequest = clientRequest.SetHeader("issueCategory", GetClientEnrichmentDetailsHeaderParams.IssueCategory)
		}

		if GetClientEnrichmentDetailsHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetClientEnrichmentDetailsHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseClientsGetClientEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetClientEnrichmentDetails(GetClientEnrichmentDetailsHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetClientEnrichmentDetails")
	}

	result := response.Result().(*ResponseClientsGetClientEnrichmentDetails)
	return result, response, err

}

//GetOverallClientHealth Get Overall Client Health - 3f9f-d80e-4df9-863c
/* Returns Overall Client Health information by Client type (Wired and Wireless) for any given point of time


@param GetOverallClientHealthQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-overall-client-health
*/
func (s *ClientsService) GetOverallClientHealth(GetOverallClientHealthQueryParams *GetOverallClientHealthQueryParams) (*ResponseClientsGetOverallClientHealth, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-health"

	queryString, _ := query.Values(GetOverallClientHealthQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsGetOverallClientHealth{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetOverallClientHealth(GetOverallClientHealthQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetOverallClientHealth")
	}

	result := response.Result().(*ResponseClientsGetOverallClientHealth)
	return result, response, err

}

//ClientProximity Client Proximity - 4497-ebe2-4c88-84a1
/* This intent API will provide client proximity information for a specific wireless user. Proximity is defined as presence on the same floor at the same time as the specified wireless user. The Proximity workflow requires the subscription to the following event (via the Event Notification workflow) prior to making this API call: NETWORK-CLIENTS-3-506 Client Proximity Report.


@param ClientProximityQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!client-proximity
*/
func (s *ClientsService) ClientProximity(ClientProximityQueryParams *ClientProximityQueryParams) (*ResponseClientsClientProximity, *resty.Response, error) {
	path := "/dna/intent/api/v1/client-proximity"

	queryString, _ := query.Values(ClientProximityQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseClientsClientProximity{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ClientProximity(ClientProximityQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ClientProximity")
	}

	result := response.Result().(*ResponseClientsClientProximity)
	return result, response, err

}

//RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes Retrieves the list of clients by applying complex filters while also supporting aggregate attributes. - e982-db36-48c9-a651
/* Retrieves the list of clients by applying complex filters while also supporting aggregate attributes. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-clients-by-applying-complex-filters-while-also-supporting-aggregate-attributes
*/
func (s *ClientsService) RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes *RequestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes, RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams *RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams) (*ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams != nil {

		if RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes).
		SetResult(&ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes(requestClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes, RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributesHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes")
	}

	result := response.Result().(*ResponseClientsRetrievesTheListOfClientsByApplyingComplexFiltersWhileAlsoSupportingAggregateAttributes)
	return result, response, err

}

//RetrievesTheNumberOfClientsByApplyingComplexFilters Retrieves the number of clients by applying complex filters. - b596-a8d5-40ba-9c5c
/* Retrieves the number of clients by applying complex filters. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-number-of-clients-by-applying-complex-filters
*/
func (s *ClientsService) RetrievesTheNumberOfClientsByApplyingComplexFilters(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFilters *RequestClientsRetrievesTheNumberOfClientsByApplyingComplexFilters, RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams *RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams) (*ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFilters, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams != nil {

		if RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFilters).
		SetResult(&ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFilters{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheNumberOfClientsByApplyingComplexFilters(requestClientsRetrievesTheNumberOfClientsByApplyingComplexFilters, RetrievesTheNumberOfClientsByApplyingComplexFiltersHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheNumberOfClientsByApplyingComplexFilters")
	}

	result := response.Result().(*ResponseClientsRetrievesTheNumberOfClientsByApplyingComplexFilters)
	return result, response, err

}

//RetrievesSummaryAnalyticsDataRelatedToClients Retrieves summary analytics data related to clients. - d7a5-8ae6-4ecb-aab9
/* Retrieves summary analytics data related to clients while applying complex filtering, aggregate functions, and grouping. This API facilitates obtaining consolidated insights into the performance and status of the clients. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-summary-analytics-data-related-to-clients
*/
func (s *ClientsService) RetrievesSummaryAnalyticsDataRelatedToClients(requestClientsRetrievesSummaryAnalyticsDataRelatedToClients *RequestClientsRetrievesSummaryAnalyticsDataRelatedToClients, RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams *RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams) (*ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClients, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams != nil {

		if RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesSummaryAnalyticsDataRelatedToClients).
		SetResult(&ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClients{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSummaryAnalyticsDataRelatedToClients(requestClientsRetrievesSummaryAnalyticsDataRelatedToClients, RetrievesSummaryAnalyticsDataRelatedToClientsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSummaryAnalyticsDataRelatedToClients")
	}

	result := response.Result().(*ResponseClientsRetrievesSummaryAnalyticsDataRelatedToClients)
	return result, response, err

}

//RetrievesTheTopNAnalyticsDataRelatedToClients Retrieves the Top-N analytics data related to clients. - dd98-6b95-401a-90e7
/* Retrieves the top N analytics data related to clients based on the provided input data. This API facilitates obtaining insights into the top-performing or most impacted clients. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-top-n-analytics-data-related-to-clients
*/
func (s *ClientsService) RetrievesTheTopNAnalyticsDataRelatedToClients(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClients *RequestClientsRetrievesTheTopNAnalyticsDataRelatedToClients, RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams *RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams) (*ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClients, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/topNAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams != nil {

		if RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClients).
		SetResult(&ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClients{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTopNAnalyticsDataRelatedToClients(requestClientsRetrievesTheTopNAnalyticsDataRelatedToClients, RetrievesTheTopNAnalyticsDataRelatedToClientsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTopNAnalyticsDataRelatedToClients")
	}

	result := response.Result().(*ResponseClientsRetrievesTheTopNAnalyticsDataRelatedToClients)
	return result, response, err

}

//RetrievesTheTrendAnalyticsDataRelatedToClients Retrieves the Trend analytics data related to clients. - 5d8e-eb0c-4988-af1f
/* Retrieves the trend analytics of client data for the specified time range. The data will be grouped based on the given trend time interval. This API facilitates obtaining consolidated insights into the performance and status of the clients over the specified start and end time. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-trend-analytics-data-related-to-clients
*/
func (s *ClientsService) RetrievesTheTrendAnalyticsDataRelatedToClients(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClients *RequestClientsRetrievesTheTrendAnalyticsDataRelatedToClients, RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams *RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams) (*ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClients, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams != nil {

		if RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClients).
		SetResult(&ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClients{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTrendAnalyticsDataRelatedToClients(requestClientsRetrievesTheTrendAnalyticsDataRelatedToClients, RetrievesTheTrendAnalyticsDataRelatedToClientsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheTrendAnalyticsDataRelatedToClients")
	}

	result := response.Result().(*ResponseClientsRetrievesTheTrendAnalyticsDataRelatedToClients)
	return result, response, err

}

//RetrievesSpecificClientInformationOverASpecifiedPeriodOfTime Retrieves specific client information over a specified period of time. - d086-9874-4ffa-a996
/* Retrieves the time series information of a specific client by applying complex filters, aggregate functions, and grouping. The data will be grouped based on the specified trend time interval. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-clients1-1.0.0-resolved.yaml


@param id id path parameter. id is the client mac address. It can be specified in one of the notational conventions 01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive

@param RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-client-information-over-a-specified-period-of-time
*/
func (s *ClientsService) RetrievesSpecificClientInformationOverASpecifiedPeriodOfTime(id string, requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime *RequestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime, RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams *RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams) (*ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime, *resty.Response, error) {
	path := "/dna/data/api/v1/clients/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams != nil {

		if RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime).
		SetResult(&ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificClientInformationOverASpecifiedPeriodOfTime(id, requestClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime, RetrievesSpecificClientInformationOverASpecifiedPeriodOfTimeHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSpecificClientInformationOverASpecifiedPeriodOfTime")
	}

	result := response.Result().(*ResponseClientsRetrievesSpecificClientInformationOverASpecifiedPeriodOfTime)
	return result, response, err

}

//QueryClientsEnergy Query clients energy - 3d9c-da44-4a6a-951e
/* Retrieves a list of client devices along with their energy data for a specified time range, based on the filters provided in the request body. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param QueryClientsEnergyHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-clients-energy
*/
func (s *ClientsService) QueryClientsEnergy(requestClientsQueryClientsEnergy *RequestClientsQueryClientsEnergy, QueryClientsEnergyHeaderParams *QueryClientsEnergyHeaderParams) (*ResponseClientsQueryClientsEnergy, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/clients/query"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if QueryClientsEnergyHeaderParams != nil {

		if QueryClientsEnergyHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", QueryClientsEnergyHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsQueryClientsEnergy).
		SetResult(&ResponseClientsQueryClientsEnergy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryClientsEnergy(requestClientsQueryClientsEnergy, QueryClientsEnergyHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation QueryClientsEnergy")
	}

	result := response.Result().(*ResponseClientsQueryClientsEnergy)
	return result, response, err

}

//CountClientsEnergyFromQuery Count clients energy from query - 5498-59c5-4508-97da
/* Retrieves the total count of client devices based on the specified complex filters. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param CountClientsEnergyFromQueryHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-clients-energy-from-query
*/
func (s *ClientsService) CountClientsEnergyFromQuery(requestClientsCountClientsEnergyFromQuery *RequestClientsCountClientsEnergyFromQuery, CountClientsEnergyFromQueryHeaderParams *CountClientsEnergyFromQueryHeaderParams) (*ResponseClientsCountClientsEnergyFromQuery, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/clients/query/count"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CountClientsEnergyFromQueryHeaderParams != nil {

		if CountClientsEnergyFromQueryHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", CountClientsEnergyFromQueryHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestClientsCountClientsEnergyFromQuery).
		SetResult(&ResponseClientsCountClientsEnergyFromQuery{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountClientsEnergyFromQuery(requestClientsCountClientsEnergyFromQuery, CountClientsEnergyFromQueryHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CountClientsEnergyFromQuery")
	}

	result := response.Result().(*ResponseClientsCountClientsEnergyFromQuery)
	return result, response, err

}
