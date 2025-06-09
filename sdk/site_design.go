package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SiteDesignService service

type GetSiteAssignedNetworkDevicesQueryParams struct {
	SiteID string  `url:"siteId,omitempty"` //Site Id. It must be area Id or building Id or floor Id.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type GetSiteAssignedNetworkDevicesCountQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site Id. It must be area Id or building Id or floor Id.
}
type GetSiteNotAssignedNetworkDevicesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type RetrievesTheListOfNetworkProfilesForSitesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
	Type   string  `url:"type,omitempty"`   //Filter responses to only include profiles of a given type
}
type RetrievesTheCountOfNetworkProfilesForSitesQueryParams struct {
	Type string `url:"type,omitempty"` //Filter the response to only count profiles of a given type
}
type RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //The id or ids of the network profile, retrievable from /dna/intent/api/v1/sites.. A list of profile ids can be passed as a queryParameter in two ways:  1. a comma-separated string ( siteId=388a23e9-4739-4be7-a0aa-cc5a95d158dd,2726dc60-3a12-451e-947a-d972ebf58743), or... 2. as separate query parameters with the same name ( siteId=388a23e9-4739-4be7-a0aa-cc5a95d158dd&siteId=2726dc60-3a12-451e-947a-d972ebf58743
}
type GetSitesQueryParams struct {
	Name           string  `url:"name,omitempty"`            //Site name.
	NameHierarchy  string  `url:"nameHierarchy,omitempty"`   //Site name hierarchy.
	Type           string  `url:"type,omitempty"`            //Site type.
	UnitsOfMeasure string  `url:"_unitsOfMeasure,omitempty"` //Floor units of measure
	Offset         float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1.
	Limit          float64 `url:"limit,omitempty"`           //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type GetSitesCountQueryParams struct {
	Name string `url:"name,omitempty"` //Site name.
}
type RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type GetAccessPointsPositionsV2QueryParams struct {
	Name       string  `url:"name,omitempty"`       //Access Point name.
	MacAddress string  `url:"macAddress,omitempty"` //Access Point mac address.
	Type       string  `url:"type,omitempty"`       //Access Point type.
	Model      string  `url:"model,omitempty"`      //Access Point model.
	Offset     float64 `url:"offset,omitempty"`     //The first record to show for this page; the first record is numbered 1. Minimum: 1
	Limit      float64 `url:"limit,omitempty"`      //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type GetAccessPointsPositionsCountV2QueryParams struct {
	Name       string `url:"name,omitempty"`       //Access Point name.
	MacAddress string `url:"macAddress,omitempty"` //Access Point mac address.
	Type       string `url:"type,omitempty"`       //Access Point type.
	Model      string `url:"model,omitempty"`      //Access Point model.
}
type GetPlannedAccessPointsPositionsV2QueryParams struct {
	Name       string  `url:"name,omitempty"`       //Planned Access Point name.
	MacAddress string  `url:"macAddress,omitempty"` //Planned Access Point mac address.
	Type       string  `url:"type,omitempty"`       //Planned Access Point type.
	Offset     float64 `url:"offset,omitempty"`     //The first record to show for this page; the first record is numbered 1. Minimum: 1
	Limit      float64 `url:"limit,omitempty"`      //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}
type GetPlannedAccessPointsPositionsCountV2QueryParams struct {
	Name       string `url:"name,omitempty"`       //Planned Access Point name.
	MacAddress string `url:"macAddress,omitempty"` //Planned Access Point mac address.
	Type       string `url:"type,omitempty"`       //Planned Access Point type.
}
type GetsAFloorV2QueryParams struct {
	UnitsOfMeasure string `url:"_unitsOfMeasure,omitempty"` //Floor units of measure
}

type ResponseSiteDesignCreatesAnArea struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesAnAreaResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesAnAreaResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesAnArea struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesAnAreaResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesAnAreaResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignDeletesAnArea struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesAnAreaResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesAnAreaResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsAnArea struct {
	Response *ResponseSiteDesignGetsAnAreaResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetsAnAreaResponse struct {
	ID              string `json:"id,omitempty"`              // Aread Id. Read only.
	SiteHierarchyID string `json:"siteHierarchyId,omitempty"` // Site Hierarchical Id. Read Only. Can be used to add the access groups using the API POST:/dna/system/api/v1/accessGroups, this value should be used to populate the srcResourceId field of the request payload.
	Name            string `json:"name,omitempty"`            // Area name
	NameHierarchy   string `json:"nameHierarchy,omitempty"`   // Area hierarchical name. Read only.
	ParentID        string `json:"parentId,omitempty"`        // Parent Id
	Type            string `json:"type,omitempty"`            // Site Type.
}
type ResponseSiteDesignAssignNetworkDevicesToASite struct {
	Version  string                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignNetworkDevicesToASiteResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignNetworkDevicesToASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSiteAssignedNetworkDevices struct {
	Response *[]ResponseSiteDesignGetSiteAssignedNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesResponse struct {
	DeviceID          string `json:"deviceId,omitempty"`          // Site assigned network device Id.
	SiteID            string `json:"siteId,omitempty"`            // Site Id where device has been assigned.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site name hierarchy
	SiteType          string `json:"siteType,omitempty"`          // Type of the site where device has been assigned.
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesCount struct {
	Response *ResponseSiteDesignGetSiteAssignedNetworkDevicesCountResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // The version of the response
}
type ResponseSiteDesignGetSiteAssignedNetworkDevicesCountResponse struct {
	Count *int `json:"count,omitempty"` // The total number of records related to the resource
}
type ResponseSiteDesignGetDeviceControllabilitySettings struct {
	Response *ResponseSiteDesignGetDeviceControllabilitySettingsResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetDeviceControllabilitySettingsResponse struct {
	AutocorrectTelemetryConfig *bool `json:"autocorrectTelemetryConfig,omitempty"` // If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
	DeviceControllability      *bool `json:"deviceControllability,omitempty"`      // If it is true, device controllability is enabled. If it is false, device controllability is disabled.
}
type ResponseSiteDesignUpdateDeviceControllabilitySettings struct {
	Version  string                                                         `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdateDeviceControllabilitySettingsResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdateDeviceControllabilitySettingsResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevices struct {
	Response *ResponseSiteDesignGetSiteNotAssignedNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesResponse struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Network device Ids.
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCount struct {
	Response *ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // The version of the response
}
type ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCountResponse struct {
	Count *int `json:"count,omitempty"` // The total number of records related to the resource
}
type ResponseSiteDesignUnassignNetworkDevicesFromSites struct {
	Version  string                                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUnassignNetworkDevicesFromSitesResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUnassignNetworkDevicesFromSitesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSiteAssignedNetworkDevice struct {
	Response *ResponseSiteDesignGetSiteAssignedNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseSiteDesignGetSiteAssignedNetworkDeviceResponse struct {
	DeviceID          string `json:"deviceId,omitempty"`          // Site assigned network device Id.
	SiteID            string `json:"siteId,omitempty"`            // Site Id where device has been assigned.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site name hierarchy
	SiteType          string `json:"siteType,omitempty"`          // Type of the site where device has been assigned.
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSites struct {
	Response *[]ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesResponse struct {
	ID   string `json:"id,omitempty"`   // The ID of this network profile.
	Name string `json:"name,omitempty"` // The name of the network profile.
	Type string `json:"type,omitempty"` // Type
}
type ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSites struct {
	Response *ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignDeletesANetworkProfileForSites struct {
	Version  string                                                    `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesANetworkProfileForSitesResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesANetworkProfileForSitesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignRetrieveANetworkProfileForSitesByID struct {
	Response *ResponseSiteDesignRetrieveANetworkProfileForSitesByIDResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignRetrieveANetworkProfileForSitesByIDResponse struct {
	ID   string `json:"id,omitempty"`   // The ID of this network profile.
	Name string `json:"name,omitempty"` // The name of the network profile.
	Type string `json:"type,omitempty"` // Type
}
type ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSite struct {
	Version  string                                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo struct {
	Response *[]ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse `json:"response,omitempty"` //
	Version  string                                                                                             `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse struct {
	ID string `json:"id,omitempty"` // Id
}
type ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSites struct {
	Version  string                                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSites struct {
	Version  string                                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo struct {
	Response *ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse `json:"response,omitempty"` //
	Version  string                                                                                            `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromASite struct {
	Version  string                                                               `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUnassignsANetworkProfileForSitesFromASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignAssociate struct {
	Version  string                               `json:"version,omitempty"`  // Version
	Response *ResponseSiteDesignAssociateResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignAssociateResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSiteDesignDisassociate struct {
	Version  string                                  `json:"version,omitempty"`  // Version
	Response *ResponseSiteDesignDisassociateResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignDisassociateResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseSiteDesignGetSites struct {
	Response *[]ResponseSiteDesignGetSitesResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetSitesResponse struct {
	NameHierarchy   string   `json:"nameHierarchy,omitempty"`   // Site hierarchical name. Read only. Example: Global/USA/San Jose/Building1
	Name            string   `json:"name,omitempty"`            // Site name.
	Latitude        *float64 `json:"latitude,omitempty"`        // Building Latitude. Example: 37.403712
	Longitude       *float64 `json:"longitude,omitempty"`       // Building Longitude. Example: -121.971063
	Address         string   `json:"address,omitempty"`         // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country         string   `json:"country,omitempty"`         // Country name for the building.
	FloorNumber     *int     `json:"floorNumber,omitempty"`     // Floor number
	RfModel         string   `json:"rfModel,omitempty"`         // Floor RF Model
	Width           *float64 `json:"width,omitempty"`           // Floor width. Example : 100.5
	Length          *float64 `json:"length,omitempty"`          // Floor length. Example : 110.3
	Height          *float64 `json:"height,omitempty"`          // Floor height. Example : 10.1
	UnitsOfMeasure  string   `json:"unitsOfMeasure,omitempty"`  // Floor unit of measure
	Type            string   `json:"type,omitempty"`            // Type
	ID              string   `json:"id,omitempty"`              // Site Id. Read only.
	ParentID        string   `json:"parentId,omitempty"`        // Parent Id. Read only
	SiteHierarchyID string   `json:"siteHierarchyId,omitempty"` // Site Hierarchical Id. Read only. Can be used to add the access groups using the API POST:/dna/system/api/v1/accessGroups, this value should be used to populate the srcResourceId field of the request payload.
}
type ResponseSiteDesignCreateSites struct {
	Version  string                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreateSitesResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignCreateSitesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetSitesCount []ResponseItemSiteDesignGetSitesCount // Array of ResponseSiteDesignGetSitesCount
type ResponseItemSiteDesignGetSitesCount struct {
	Response *ResponseItemSiteDesignGetSitesCountResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // The version of the response
}
type ResponseItemSiteDesignGetSitesCountResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned struct {
	Response *[]ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedResponse `json:"response,omitempty"` //
	Version  string                                                                                        `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedResponse struct {
	ID string `json:"id,omitempty"` // Id
}
type ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned struct {
	Response *ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedResponse `json:"response,omitempty"` //
	Version  string                                                                                `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignCreatesABuildingV2 struct {
	Version  string                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesABuildingV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesABuildingV2 struct {
	Version  string                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesABuildingV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignDeletesABuildingV2 struct {
	Version  string                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesABuildingV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsABuildingV2 struct {
	Response *ResponseSiteDesignGetsABuildingV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignGetsABuildingV2Response struct {
	ParentID        string   `json:"parentId,omitempty"`        // Parent Id
	Name            string   `json:"name,omitempty"`            // Building name
	Latitude        *float64 `json:"latitude,omitempty"`        // Building Latitude. Example: 37.403712
	Longitude       *float64 `json:"longitude,omitempty"`       // Building Longitude. Example: -121.971063
	Address         string   `json:"address,omitempty"`         // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country         string   `json:"country,omitempty"`         // Country name
	Type            string   `json:"type,omitempty"`            // Example: building
	ID              string   `json:"id,omitempty"`              // Building Id. Read only
	NameHierarchy   string   `json:"nameHierarchy,omitempty"`   // Building hierarchical name. Read only
	SiteHierarchyID string   `json:"siteHierarchyId,omitempty"` // Site Hierarchical Id. Read only. Can be used to add the access groups using the API POST:/dna/system/api/v1/accessGroups, this value should be used to populate the srcResourceId field of the request payload.
}
type ResponseSiteDesignCreatesAFloorV2 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesAFloorV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesFloorSettingsV2 struct {
	Version  string                                            `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesFloorSettingsV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesFloorSettingsV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetFloorSettingsV2 struct {
	Response *ResponseSiteDesignGetFloorSettingsV2Response `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetFloorSettingsV2Response struct {
	UnitsOfMeasure string `json:"unitsOfMeasure,omitempty"` // Floor units of measure.
}
type ResponseSiteDesignGetAccessPointsPositionsV2 struct {
	Response *[]ResponseSiteDesignGetAccessPointsPositionsV2Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetAccessPointsPositionsV2Response struct {
	ID         string                                                        `json:"id,omitempty"`         // Access Point Id (readonly)
	MacAddress string                                                        `json:"macAddress,omitempty"` // Access Point MAC address (readonly)
	Model      string                                                        `json:"model,omitempty"`      // Access Point model (readonly)
	Name       string                                                        `json:"name,omitempty"`       // Access Point Name (readonly)
	Type       string                                                        `json:"type,omitempty"`       // Access Point type (readonly)
	Position   *ResponseSiteDesignGetAccessPointsPositionsV2ResponsePosition `json:"position,omitempty"`   //
	Radios     *[]ResponseSiteDesignGetAccessPointsPositionsV2ResponseRadios `json:"radios,omitempty"`     //
}
type ResponseSiteDesignGetAccessPointsPositionsV2ResponsePosition struct {
	X *float64 `json:"x,omitempty"` // Access Point X coordinate in feet
	Y *float64 `json:"y,omitempty"` // Access Point Y coordinate in feet
	Z *float64 `json:"z,omitempty"` // Access Point Z coordinate in feet
}
type ResponseSiteDesignGetAccessPointsPositionsV2ResponseRadios struct {
	ID      string                                                             `json:"id,omitempty"`      // Radio Id (readonly)
	Bands   *[]float64                                                         `json:"bands,omitempty"`   // Radio frequencies in GHz (readonly). Radio frequencies are expected to be 2.4, 5, and 6. MinItems: 1; MaxItems: 3
	Channel *int                                                               `json:"channel,omitempty"` // Channel to be used by the Access Point (readonly). The value gets updated only every 24 hours
	TxPower *int                                                               `json:"txPower,omitempty"` // Transmit power for the channel in Decibel milliwatts (dBm) (readonly). The value gets updated only every 24 hours
	Antenna *ResponseSiteDesignGetAccessPointsPositionsV2ResponseRadiosAntenna `json:"antenna,omitempty"` //
}
type ResponseSiteDesignGetAccessPointsPositionsV2ResponseRadiosAntenna struct {
	Name      string `json:"name,omitempty"`      // Antenna type for this Access Point. Use `/dna/intent/api/v1/maps/supported-access-points` to find supported Antennas for a particualr Access Point model.
	Azimuth   *int   `json:"azimuth,omitempty"`   // Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
	Elevation *int   `json:"elevation,omitempty"` // Elevation of the antenna. The elevation range is from -90 through 90
}
type ResponseSiteDesignEditTheAccessPointsPositionsV2 struct {
	Version  string                                                    `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignEditTheAccessPointsPositionsV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignEditTheAccessPointsPositionsV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetAccessPointsPositionsCountV2 struct {
	Response *ResponseSiteDesignGetAccessPointsPositionsCountV2Response `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetAccessPointsPositionsCountV2Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsV2 struct {
	Response *[]ResponseSiteDesignGetPlannedAccessPointsPositionsV2Response `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsV2Response struct {
	Name       string                                                               `json:"name,omitempty"`       // Planned Access Point Name
	MacAddress string                                                               `json:"macAddress,omitempty"` // Planned Access Point MAC address
	Type       string                                                               `json:"type,omitempty"`       // Planned Access Point type. Use `dna/intent/api/v1/maps/supported-access-points` to find the supported models
	Position   *ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponsePosition `json:"position,omitempty"`   //
	Radios     *[]ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponseRadios `json:"radios,omitempty"`     //
	ID         string                                                               `json:"id,omitempty"`         // Planned Access Point Id (readonly)
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponsePosition struct {
	X *float64 `json:"x,omitempty"` // Planned Access Point X coordinate in feet
	Y *float64 `json:"y,omitempty"` // Planned Access Point Y coordinate in feet
	Z *float64 `json:"z,omitempty"` // Planned Access Point Z coordinate in feet
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponseRadios struct {
	Bands   *[]float64                                                                `json:"bands,omitempty"`   // Radio frequencies in GHz (readonly). Radio frequencies are expected to be 2.4, 5, and 6. MinItems: 1; MaxItems: 3
	Channel *int                                                                      `json:"channel,omitempty"` // Channel to be used by the Planned Access Point. In the context of a Planned Access Point, the channel have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
	TxPower *int                                                                      `json:"txPower,omitempty"` // Transmit power for the channel in Decibel milliwatts (dBm). In the context of a Planned Access Point, the txPower have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
	Antenna *ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponseRadiosAntenna `json:"antenna,omitempty"` //
	ID      string                                                                    `json:"id,omitempty"`      // Radio Id (readonly)
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsV2ResponseRadiosAntenna struct {
	Name      string `json:"name,omitempty"`      // Antenna type for this Planned Access Point. Use `/dna/intent/api/v1/maps/supported-access-points` to find supported Antennas for a particualr Planned Access Point type
	Azimuth   *int   `json:"azimuth,omitempty"`   // Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
	Elevation *int   `json:"elevation,omitempty"` // Elevation of the antenna. The elevation range is from -90 through 90
}
type ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 struct {
	Version  string                                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignAddPlannedAccessPointsPositionsV2 struct {
	Version  string                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignAddPlannedAccessPointsPositionsV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignAddPlannedAccessPointsPositionsV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignEditPlannedAccessPointsPositionsV2 struct {
	Version  string                                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignEditPlannedAccessPointsPositionsV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignEditPlannedAccessPointsPositionsV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2 struct {
	Response *ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2Response `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2Response struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSiteDesignDeletePlannedAccessPointsPositionV2 struct {
	Version  string                                                         `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletePlannedAccessPointsPositionV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletePlannedAccessPointsPositionV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesAFloorV2 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesAFloorV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsAFloorV2 struct {
	Response *ResponseSiteDesignGetsAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignGetsAFloorV2Response struct {
	ParentID        string   `json:"parentId,omitempty"`        // Parent Id.
	Name            string   `json:"name,omitempty"`            // Floor name
	FloorNumber     *int     `json:"floorNumber,omitempty"`     // Floor number
	RfModel         string   `json:"rfModel,omitempty"`         // RF Model
	Width           *float64 `json:"width,omitempty"`           // Floor width. Example : 100.5
	Length          *float64 `json:"length,omitempty"`          // Floor length. Example : 110.3
	Height          *float64 `json:"height,omitempty"`          // Floor height. Example : 10.1
	UnitsOfMeasure  string   `json:"unitsOfMeasure,omitempty"`  // Units Of Measure
	Type            string   `json:"type,omitempty"`            // Example : floor
	ID              string   `json:"id,omitempty"`              // Floor Id. Read only.
	NameHierarchy   string   `json:"nameHierarchy,omitempty"`   // Floor hierarchical name. Read only.
	SiteHierarchyID string   `json:"siteHierarchyId,omitempty"` // Site Hierarchical Id. Read only. Can be used to add the access groups using the API POST:/dna/system/api/v1/accessGroups, this value should be used to populate the srcResourceId field of the request payload.
}
type ResponseSiteDesignDeletesAFloorV2 struct {
	Version  string                                     `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesAFloorV2Response `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesAFloorV2Response struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type RequestSiteDesignCreatesAnArea struct {
	Name     string `json:"name,omitempty"`     // Area name
	ParentID string `json:"parentId,omitempty"` // Parent Id
}
type RequestSiteDesignUpdatesAnArea struct {
	Name     string `json:"name,omitempty"`     // Area name
	ParentID string `json:"parentId,omitempty"` // Parent Id
}
type RequestSiteDesignAssignNetworkDevicesToASite struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Unassigned network devices, ranging from a minimum of 1 to a maximum of 100.
	SiteID    string   `json:"siteId,omitempty"`    // This must be building Id or floor Id. Access points, Sensors are assigned to floor. Remaining network devices are assigned to building. Site Id can be retrieved using '/intent/api/v1/sites'.
}
type RequestSiteDesignUpdateDeviceControllabilitySettings struct {
	AutocorrectTelemetryConfig *bool `json:"autocorrectTelemetryConfig,omitempty"` // If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
	DeviceControllability      *bool `json:"deviceControllability,omitempty"`      // If it is true, device controllability is enabled. If it is false, device controllability is disabled.
}
type RequestSiteDesignUnassignNetworkDevicesFromSites struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Network device Ids, ranging from a minimum of 1 to a maximum of 100.
}
type RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSite struct {
	ID string `json:"id,omitempty"` // Id
}
type RequestSiteDesignAssignANetworkProfileForSitesToAListOfSites struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type RequestSiteDesignCreateSites []RequestItemSiteDesignCreateSites // Array of RequestSiteDesignCreateSites
type RequestItemSiteDesignCreateSites struct {
	ParentNameHierarchy string   `json:"parentNameHierarchy,omitempty"` // Parent hierarchical name. Example: Global/USA/San Jose/Building1
	Name                string   `json:"name,omitempty"`                // Site name.
	Latitude            *float64 `json:"latitude,omitempty"`            // Building Latitude. Example: 37.403712
	Longitude           *float64 `json:"longitude,omitempty"`           // Building Longitude. Example: -121.971063
	Address             string   `json:"address,omitempty"`             // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country             string   `json:"country,omitempty"`             // Country name. Required for building.
	FloorNumber         *int     `json:"floorNumber,omitempty"`         // Floor number. Required for floor.
	RfModel             string   `json:"rfModel,omitempty"`             // Floor RF Model. Required for floor.
	Width               *float64 `json:"width,omitempty"`               // Floor width. Required for floor. Example : 100.5
	Length              *float64 `json:"length,omitempty"`              // Floor length. Required for floor. Example : 110.3
	Height              *float64 `json:"height,omitempty"`              // Floor height. Required for floor. Example : 10.1
	UnitsOfMeasure      string   `json:"unitsOfMeasure,omitempty"`      // Floor unit of measure. Required for floor.
	Type                string   `json:"type,omitempty"`                // Type
}
type RequestSiteDesignCreatesABuildingV2 struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States. Please note that if only the address is provided when creating a building, the UI will not display the geo-location on the map. To ensure the location is rendered, you must also provide the latitude and longitude. If a building has been created without these coordinates and you wish to display its geo-location on the map later, you can edit the building details via the UI to include the latitude and longitude. This limitation will be resolved in a future release.
	Country   string   `json:"country,omitempty"`   // Country name
}
type RequestSiteDesignUpdatesABuildingV2 struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States. Please note that if only the address is provided when creating a building, the UI will not display the geo-location on the map. To ensure the location is rendered, you must also provide the latitude and longitude. If a building has been created without these coordinates and you wish to display its geo-location on the map later, you can edit the building details via the UI to include the latitude and longitude. This limitation will be resolved in a future release.
	Country   string   `json:"country,omitempty"`   // Country name
}
type RequestSiteDesignCreatesAFloorV2 struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
}
type RequestSiteDesignUpdatesFloorSettingsV2 struct {
	UnitsOfMeasure string `json:"unitsOfMeasure,omitempty"` // Floor units of measure
}
type RequestSiteDesignEditTheAccessPointsPositionsV2 []RequestItemSiteDesignEditTheAccessPointsPositionsV2 // Array of RequestSiteDesignEditTheAccessPointsPositionsV2
type RequestItemSiteDesignEditTheAccessPointsPositionsV2 struct {
	ID       string                                                       `json:"id,omitempty"`       // Access Point Id
	Position *RequestItemSiteDesignEditTheAccessPointsPositionsV2Position `json:"position,omitempty"` //
	Radios   *[]RequestItemSiteDesignEditTheAccessPointsPositionsV2Radios `json:"radios,omitempty"`   //
}
type RequestItemSiteDesignEditTheAccessPointsPositionsV2Position struct {
	X *float64 `json:"x,omitempty"` // Access Point X coordinate in feet
	Y *float64 `json:"y,omitempty"` // Access Point Y coordinate in feet
	Z *float64 `json:"z,omitempty"` // Access Point Z coordinate in feet
}
type RequestItemSiteDesignEditTheAccessPointsPositionsV2Radios struct {
	ID      string                                                            `json:"id,omitempty"`      // Radio Id
	Antenna *RequestItemSiteDesignEditTheAccessPointsPositionsV2RadiosAntenna `json:"antenna,omitempty"` //
}
type RequestItemSiteDesignEditTheAccessPointsPositionsV2RadiosAntenna struct {
	Name      string `json:"name,omitempty"`      // Antenna type for this Access Point. Use `/dna/intent/api/v1/maps/supported-access-points` to find supported Antennas for a particualr Access Point model
	Azimuth   *int   `json:"azimuth,omitempty"`   // Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
	Elevation *int   `json:"elevation,omitempty"` // Elevation of the antenna. The elevation range is from -90 through 90
}
type RequestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 []RequestItemSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 // Array of RequestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2
type RequestItemSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 struct {
	PlannedAccessPointID string `json:"plannedAccessPointId,omitempty"` // Planned Access Point Id
	AccessPointID        string `json:"accessPointId,omitempty"`        // Operational Access Point Id
}
type RequestSiteDesignAddPlannedAccessPointsPositionsV2 []RequestItemSiteDesignAddPlannedAccessPointsPositionsV2 // Array of RequestSiteDesignAddPlannedAccessPointsPositionsV2
type RequestItemSiteDesignAddPlannedAccessPointsPositionsV2 struct {
	Name       string                                                          `json:"name,omitempty"`       // Planned Access Point Name
	MacAddress string                                                          `json:"macAddress,omitempty"` // Planned Access Point MAC address
	Type       string                                                          `json:"type,omitempty"`       // Planned Access Point type. Use `dna/intent/api/v1/maps/supported-access-points` to find the supported models
	Position   *RequestItemSiteDesignAddPlannedAccessPointsPositionsV2Position `json:"position,omitempty"`   //
	Radios     *[]RequestItemSiteDesignAddPlannedAccessPointsPositionsV2Radios `json:"radios,omitempty"`     //
}
type RequestItemSiteDesignAddPlannedAccessPointsPositionsV2Position struct {
	X *float64 `json:"x,omitempty"` // Planned Access Point X coordinate in feet
	Y *float64 `json:"y,omitempty"` // Planned Access Point Y coordinate in feet
	Z *float64 `json:"z,omitempty"` // Planned Access Point Z coordinate in feet
}
type RequestItemSiteDesignAddPlannedAccessPointsPositionsV2Radios struct {
	Bands   *[]float64                                                           `json:"bands,omitempty"`   // Radio frequencies in GHz. Radio frequencies are expected to be 2.4, 5, and 6. MinItems: 1; MaxItems: 3
	Channel *int                                                                 `json:"channel,omitempty"` // Channel to be used by the Planned Access Point. In the context of a Planned Access Point, the channel have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
	TxPower *int                                                                 `json:"txPower,omitempty"` // Transmit power for the channel in Decibel milliwatts (dBm). In the context of a Planned Access Point, the txPower have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
	Antenna *RequestItemSiteDesignAddPlannedAccessPointsPositionsV2RadiosAntenna `json:"antenna,omitempty"` //
}
type RequestItemSiteDesignAddPlannedAccessPointsPositionsV2RadiosAntenna struct {
	Name      string `json:"name,omitempty"`      // Antenna type for this Planned Access Point. Use `/dna/intent/api/v1/maps/supported-access-points` to find supported Antennas for a particualr Planned Access Point type
	Azimuth   *int   `json:"azimuth,omitempty"`   // Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
	Elevation *int   `json:"elevation,omitempty"` // Elevation of the antenna. The elevation range is from -90 through 90
}
type RequestSiteDesignEditPlannedAccessPointsPositionsV2 []RequestItemSiteDesignEditPlannedAccessPointsPositionsV2 // Array of RequestSiteDesignEditPlannedAccessPointsPositionsV2
type RequestItemSiteDesignEditPlannedAccessPointsPositionsV2 struct {
	Name       string                                                           `json:"name,omitempty"`       // Planned Access Point Name
	MacAddress string                                                           `json:"macAddress,omitempty"` // Planned Access Point MAC address
	Type       string                                                           `json:"type,omitempty"`       // Planned Access Point type. Use `dna/intent/api/v1/maps/supported-access-points` to find the supported models
	Position   *RequestItemSiteDesignEditPlannedAccessPointsPositionsV2Position `json:"position,omitempty"`   //
	Radios     *[]RequestItemSiteDesignEditPlannedAccessPointsPositionsV2Radios `json:"radios,omitempty"`     //
	ID         string                                                           `json:"id,omitempty"`         // Planned Access Point Id
}
type RequestItemSiteDesignEditPlannedAccessPointsPositionsV2Position struct {
	X *float64 `json:"x,omitempty"` // Planned Access Point X coordinate in feet
	Y *float64 `json:"y,omitempty"` // Planned Access Point Y coordinate in feet
	Z *float64 `json:"z,omitempty"` // Planned Access Point Z coordinate in feet
}
type RequestItemSiteDesignEditPlannedAccessPointsPositionsV2Radios struct {
	Channel *int                                                                  `json:"channel,omitempty"` // Channel to be used by the Planned Access Point. In the context of a Planned Access Point, the channel have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
	TxPower *int                                                                  `json:"txPower,omitempty"` // Transmit power for the channel in Decibel milliwatts (dBm). In the context of a Planned Access Point, the txPower have no bearing on what the real Access Point will actually be, they are just used for Maps visualization feature set
	Antenna *RequestItemSiteDesignEditPlannedAccessPointsPositionsV2RadiosAntenna `json:"antenna,omitempty"` //
	ID      string                                                                `json:"id,omitempty"`      // Radio Id
}
type RequestItemSiteDesignEditPlannedAccessPointsPositionsV2RadiosAntenna struct {
	Name      string `json:"name,omitempty"`      // Antenna type for this Planned Access Point. Use `/dna/intent/api/v1/maps/supported-access-points` to find supported Antennas for a particualr Planned Access Point type
	Azimuth   *int   `json:"azimuth,omitempty"`   // Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
	Elevation *int   `json:"elevation,omitempty"` // Elevation of the antenna. The elevation range is from -90 through 90
}
type RequestSiteDesignUpdatesAFloorV2 struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
}

//GetsAnArea Gets an area - d6af-ab3e-43bb-a73c
/* Gets an area in the network hierarchy.


@param id id path parameter. Area Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-an-area
*/
func (s *SiteDesignService) GetsAnArea(id string) (*ResponseSiteDesignGetsAnArea, *resty.Response, error) {
	path := "/dna/intent/api/v1/areas/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetsAnArea{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAnArea(id)
		}
		return nil, response, fmt.Errorf("error with operation GetsAnArea")
	}

	result := response.Result().(*ResponseSiteDesignGetsAnArea)
	return result, response, err

}

//GetSiteAssignedNetworkDevices Get site assigned network devices - 0ea1-4875-4219-995d
/* Get all site assigned network devices. The items in the list are arranged in an order that corresponds with their internal identifiers.


@param GetSiteAssignedNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-assigned-network-devices
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevices(GetSiteAssignedNetworkDevicesQueryParams *GetSiteAssignedNetworkDevicesQueryParams) (*ResponseSiteDesignGetSiteAssignedNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/assignedToSite"

	queryString, _ := query.Values(GetSiteAssignedNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSiteAssignedNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAssignedNetworkDevices(GetSiteAssignedNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAssignedNetworkDevices")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteAssignedNetworkDevices)
	return result, response, err

}

//GetSiteAssignedNetworkDevicesCount Get site assigned network devices count - fd93-c911-48ba-9386
/* Get all network devices count under the given site in the network hierarchy.


@param GetSiteAssignedNetworkDevicesCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-assigned-network-devices-count
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevicesCount(GetSiteAssignedNetworkDevicesCountQueryParams *GetSiteAssignedNetworkDevicesCountQueryParams) (*ResponseSiteDesignGetSiteAssignedNetworkDevicesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/assignedToSite/count"

	queryString, _ := query.Values(GetSiteAssignedNetworkDevicesCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSiteAssignedNetworkDevicesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAssignedNetworkDevicesCount(GetSiteAssignedNetworkDevicesCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAssignedNetworkDevicesCount")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteAssignedNetworkDevicesCount)
	return result, response, err

}

//GetDeviceControllabilitySettings Get device controllability settings - 06b2-4916-486b-aeec
/* Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs to manage devices. Changes are made on network devices during discovery, when adding a device to Inventory, or when assigning a device to a site. If changes are made to any settings that are under the scope of this process, these changes are applied to the network devices during the Provision and Update Telemetry Settings operations, even if Device Controllability is disabled. The following device settings will be enabled as part of Device Controllability when devices are discovered. SNMP Credentials. NETCONF Credentials. Subsequent to discovery, devices will be added to Inventory. The following device settings will be enabled when devices are added to inventory. Cisco TrustSec (CTS) Credentials. The following device settings will be enabled when devices are assigned to a site. Some of these settings can be defined at a site level under Design > Network Settings > Telemetry & Wireless. Wired Endpoint Data Collection Enablement. Controller Certificates. SNMP Trap Server Definitions. Syslog Server Definitions. Application Visibility. Application QoS Policy. Wireless Service Assurance (WSA). Wireless Telemetry. DTLS Ciphersuite. AP Impersonation. If Device Controllability is disabled, Catalyst Center does not configure any of the preceding credentials or settings on devices during discovery, at runtime, or during site assignment. However, the telemetry settings and related configuration are pushed when the device is provisioned or when the update Telemetry Settings action is performed. Catalyst Center identifies and automatically corrects the following telemetry configuration issues on the device. SWIM certificate issue. IOS WLC NA certificate issue. PKCS12 certificate issue. IOS telemetry configuration issu



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-controllability-settings
*/
func (s *SiteDesignService) GetDeviceControllabilitySettings() (*ResponseSiteDesignGetDeviceControllabilitySettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deviceControllability/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetDeviceControllabilitySettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceControllabilitySettings()
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceControllabilitySettings")
	}

	result := response.Result().(*ResponseSiteDesignGetDeviceControllabilitySettings)
	return result, response, err

}

//GetSiteNotAssignedNetworkDevices Get site not assigned network devices - cd89-78de-4109-8f0d
/* Get network devices that are not assigned to any site.


@param GetSiteNotAssignedNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-not-assigned-network-devices
*/
func (s *SiteDesignService) GetSiteNotAssignedNetworkDevices(GetSiteNotAssignedNetworkDevicesQueryParams *GetSiteNotAssignedNetworkDevicesQueryParams) (*ResponseSiteDesignGetSiteNotAssignedNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/notAssignedToSite"

	queryString, _ := query.Values(GetSiteNotAssignedNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSiteNotAssignedNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteNotAssignedNetworkDevices(GetSiteNotAssignedNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteNotAssignedNetworkDevices")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteNotAssignedNetworkDevices)
	return result, response, err

}

//GetSiteNotAssignedNetworkDevicesCount Get site not assigned network devices count - b28e-881d-4cda-b06c
/* Get network devices count that are not assigned to any site.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-not-assigned-network-devices-count
*/
func (s *SiteDesignService) GetSiteNotAssignedNetworkDevicesCount() (*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/notAssignedToSite/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteNotAssignedNetworkDevicesCount()
		}
		return nil, response, fmt.Errorf("error with operation GetSiteNotAssignedNetworkDevicesCount")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteNotAssignedNetworkDevicesCount)
	return result, response, err

}

//GetSiteAssignedNetworkDevice Get site assigned network device - f08f-3b31-4bda-9c96
/* Get site assigned network device. The items in the list are arranged in an order that corresponds with their internal identifiers.


@param id id path parameter. Network Device Id.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-assigned-network-device
*/
func (s *SiteDesignService) GetSiteAssignedNetworkDevice(id string) (*ResponseSiteDesignGetSiteAssignedNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/{id}/assignedToSite"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetSiteAssignedNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteAssignedNetworkDevice(id)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteAssignedNetworkDevice")
	}

	result := response.Result().(*ResponseSiteDesignGetSiteAssignedNetworkDevice)
	return result, response, err

}

//RetrievesTheListOfNetworkProfilesForSites Retrieves the list of network profiles for sites - a78d-8918-4898-9cf2
/* Retrieves the list of network profiles for sites.


@param RetrievesTheListOfNetworkProfilesForSitesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-profiles-for-sites
*/
func (s *SiteDesignService) RetrievesTheListOfNetworkProfilesForSites(RetrievesTheListOfNetworkProfilesForSitesQueryParams *RetrievesTheListOfNetworkProfilesForSitesQueryParams) (*ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites"

	queryString, _ := query.Values(RetrievesTheListOfNetworkProfilesForSitesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSites{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkProfilesForSites(RetrievesTheListOfNetworkProfilesForSitesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkProfilesForSites")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSites)
	return result, response, err

}

//RetrievesTheCountOfNetworkProfilesForSites Retrieves the count of network profiles for sites - 57a7-c9d2-4f4a-b000
/* Retrieves the count of network profiles for sites


@param RetrievesTheCountOfNetworkProfilesForSitesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-network-profiles-for-sites
*/
func (s *SiteDesignService) RetrievesTheCountOfNetworkProfilesForSites(RetrievesTheCountOfNetworkProfilesForSitesQueryParams *RetrievesTheCountOfNetworkProfilesForSitesQueryParams) (*ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/count"

	queryString, _ := query.Values(RetrievesTheCountOfNetworkProfilesForSitesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSites{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfNetworkProfilesForSites(RetrievesTheCountOfNetworkProfilesForSitesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfNetworkProfilesForSites")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSites)
	return result, response, err

}

//RetrieveANetworkProfileForSitesByID Retrieve a network profile for sites by id - 87a8-eb79-4f28-acc4
/* Retrieves a network profile for sites by id.


@param id id path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-network-profile-for-sites-by-id
*/
func (s *SiteDesignService) RetrieveANetworkProfileForSitesByID(id string) (*ResponseSiteDesignRetrieveANetworkProfileForSitesByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignRetrieveANetworkProfileForSitesByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveANetworkProfileForSitesByID(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveANetworkProfileForSitesById")
	}

	result := response.Result().(*ResponseSiteDesignRetrieveANetworkProfileForSitesByID)
	return result, response, err

}

//RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo Retrieves the list of sites that the given network profile for sites is assigned to - 0f84-ba73-4429-accd
/* Retrieves the list of sites that the given network profile for sites is assigned to.
The list includes the sites the profile has been directly assigned to, as well as child sites that have inherited the profile.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`

@param RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-sites-that-the-given-network-profile-for-sites-is-assigned-to
*/
func (s *SiteDesignService) RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(profileID string, RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams *RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams) (*ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	queryString, _ := query.Values(RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(profileID, RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo)
	return result, response, err

}

//RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo Retrieves the count of sites that the given network profile for sites is assigned to - 53ba-29dd-42eb-bd19
/* Retrieves the count of sites that the given network profile for sites is assigned to.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-sites-that-the-given-network-profile-for-sites-is-assigned-to
*/
func (s *SiteDesignService) RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(profileID string) (*ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/count"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo(profileID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheCountOfSitesThatTheGivenNetworkProfileForSitesIsAssignedTo)
	return result, response, err

}

//GetSites Get sites - 4e8a-49c3-4b49-b291
/* Get sites.


@param GetSitesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sites
*/
func (s *SiteDesignService) GetSites(GetSitesQueryParams *GetSitesQueryParams) (*ResponseSiteDesignGetSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites"

	queryString, _ := query.Values(GetSitesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSites{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSites(GetSitesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSites")
	}

	result := response.Result().(*ResponseSiteDesignGetSites)
	return result, response, err

}

//GetSitesCount Get sites count - 0fbf-482e-446a-835f
/* Get sites count.


@param GetSitesCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sites-count
*/
func (s *SiteDesignService) GetSitesCount(GetSitesCountQueryParams *GetSitesCountQueryParams) (*ResponseSiteDesignGetSitesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/count"

	queryString, _ := query.Values(GetSitesCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetSitesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSitesCount(GetSitesCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSitesCount")
	}

	result := response.Result().(*ResponseSiteDesignGetSitesCount)
	return result, response, err

}

//RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned Retrieves the list of network profiles that the given site has been assigned - b0b4-5962-49d9-8d6b
/* Retrieves the list of profiles that the given site has been assigned.  These profiles may either be directly assigned to this site, or were assigned to a parent site and have been inherited.
These assigments can be modified via the `/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments` resources.


@param siteID siteId path parameter. The `id` of the site, retrievable from `/dna/intent/api/v1/sites`

@param RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-profiles-that-the-given-site-has-been-assigned
*/
func (s *SiteDesignService) RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned(siteID string, RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams *RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams) (*ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/profileAssignments"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned(siteID, RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssigned)
	return result, response, err

}

//RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned Retrieves the count of profiles that the given site has been assigned - 28be-9a3f-4688-a2d4
/* Retrieves the count of profiles that the given site has been assigned.  These profiles may either be directly assigned to this site, or were assigned to a parent site and have been inherited.


@param siteID siteId path parameter. The `id` of the site, retrievable from `/dna/intent/api/v1/sites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-profiles-that-the-given-site-has-been-assigned
*/
func (s *SiteDesignService) RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned(siteID string) (*ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/profileAssignments/count"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned(siteID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned")
	}

	result := response.Result().(*ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssigned)
	return result, response, err

}

//GetsABuildingV2 Gets a building - e293-295e-4a78-bf64
/* Gets a building in the network hierarchy.


@param id id path parameter. Building Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-building-v2
*/
func (s *SiteDesignService) GetsABuildingV2(id string) (*ResponseSiteDesignGetsABuildingV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetsABuildingV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsABuildingV2(id)
		}
		return nil, response, fmt.Errorf("error with operation GetsABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignGetsABuildingV2)
	return result, response, err

}

//GetFloorSettingsV2 Get floor settings - f697-a95f-4469-958f
/* Gets UI user preference for floor unit system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-floor-settings-v2
*/
func (s *SiteDesignService) GetFloorSettingsV2() (*ResponseSiteDesignGetFloorSettingsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetFloorSettingsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFloorSettingsV2()
		}
		return nil, response, fmt.Errorf("error with operation GetFloorSettingsV2")
	}

	result := response.Result().(*ResponseSiteDesignGetFloorSettingsV2)
	return result, response, err

}

//GetAccessPointsPositionsV2 Get Access Points positions - fd8b-0aab-4a49-8a71
/* Retrieve all Access Points positions assigned for a specific floor


@param floorID floorId path parameter. Floor Id

@param GetAccessPointsPositionsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-points-positions-v2
*/
func (s *SiteDesignService) GetAccessPointsPositionsV2(floorID string, GetAccessPointsPositionsV2QueryParams *GetAccessPointsPositionsV2QueryParams) (*ResponseSiteDesignGetAccessPointsPositionsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/accessPointPositions"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetAccessPointsPositionsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetAccessPointsPositionsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointsPositionsV2(floorID, GetAccessPointsPositionsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointsPositionsV2")
	}

	result := response.Result().(*ResponseSiteDesignGetAccessPointsPositionsV2)
	return result, response, err

}

//GetAccessPointsPositionsCountV2 Get Access Points positions count - ccab-58ba-432a-88b3
/* Retrieve Access Points positions count assigned for a specific floor


@param floorID floorId path parameter. Floor Id

@param GetAccessPointsPositionsCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-points-positions-count-v2
*/
func (s *SiteDesignService) GetAccessPointsPositionsCountV2(floorID string, GetAccessPointsPositionsCountV2QueryParams *GetAccessPointsPositionsCountV2QueryParams) (*ResponseSiteDesignGetAccessPointsPositionsCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/accessPointPositions/count"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetAccessPointsPositionsCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetAccessPointsPositionsCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointsPositionsCountV2(floorID, GetAccessPointsPositionsCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointsPositionsCountV2")
	}

	result := response.Result().(*ResponseSiteDesignGetAccessPointsPositionsCountV2)
	return result, response, err

}

//GetPlannedAccessPointsPositionsV2 Get Planned Access Points Positions - 7888-2b7d-431b-8e32
/* Retrieve all Planned Access Points Positions designated for a specific floor


@param floorID floorId path parameter. Floor Id

@param GetPlannedAccessPointsPositionsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-positions-v2
*/
func (s *SiteDesignService) GetPlannedAccessPointsPositionsV2(floorID string, GetPlannedAccessPointsPositionsV2QueryParams *GetPlannedAccessPointsPositionsV2QueryParams) (*ResponseSiteDesignGetPlannedAccessPointsPositionsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/plannedAccessPointPositions"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsPositionsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetPlannedAccessPointsPositionsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsPositionsV2(floorID, GetPlannedAccessPointsPositionsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsPositionsV2")
	}

	result := response.Result().(*ResponseSiteDesignGetPlannedAccessPointsPositionsV2)
	return result, response, err

}

//GetPlannedAccessPointsPositionsCountV2 Get Planned Access Points Positions count - 63b0-c84b-47a9-beda
/* Retrieve all Planned Access Points Positions count designated for a specific floor


@param floorID floorId path parameter. Floor Id

@param GetPlannedAccessPointsPositionsCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-planned-access-points-positions-count-v2
*/
func (s *SiteDesignService) GetPlannedAccessPointsPositionsCountV2(floorID string, GetPlannedAccessPointsPositionsCountV2QueryParams *GetPlannedAccessPointsPositionsCountV2QueryParams) (*ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/plannedAccessPointPositions/count"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	queryString, _ := query.Values(GetPlannedAccessPointsPositionsCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPlannedAccessPointsPositionsCountV2(floorID, GetPlannedAccessPointsPositionsCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPlannedAccessPointsPositionsCountV2")
	}

	result := response.Result().(*ResponseSiteDesignGetPlannedAccessPointsPositionsCountV2)
	return result, response, err

}

//GetsAFloorV2 Gets a floor - ff92-2958-4bba-9288
/* Gets a floor in the network hierarchy.


@param id id path parameter. Floor Id

@param GetsAFloorV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-floor-v2
*/
func (s *SiteDesignService) GetsAFloorV2(id string, GetsAFloorV2QueryParams *GetsAFloorV2QueryParams) (*ResponseSiteDesignGetsAFloorV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetsAFloorV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetsAFloorV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAFloorV2(id, GetsAFloorV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignGetsAFloorV2)
	return result, response, err

}

//CreatesAnArea Creates an area - a8bc-d9dc-43ea-a7e3
/* Creates an area in the network hierarchy.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-an-area
*/
func (s *SiteDesignService) CreatesAnArea(requestSiteDesignCreatesAnArea *RequestSiteDesignCreatesAnArea) (*ResponseSiteDesignCreatesAnArea, *resty.Response, error) {
	path := "/dna/intent/api/v1/areas"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesAnArea).
		SetResult(&ResponseSiteDesignCreatesAnArea{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAnArea(requestSiteDesignCreatesAnArea)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAnArea")
	}

	result := response.Result().(*ResponseSiteDesignCreatesAnArea)
	return result, response, err

}

//AssignNetworkDevicesToASite Assign network devices to a site - a1b4-8949-4a6a-a6ac
/* Assign unprovisioned network devices to a site. Along with that it can also be used to assign unprovisioned network devices to a different site. If device controllability is enabled, it will be triggered once device assigned to site successfully. Device Controllability can be enabled/disabled using `/dna/intent/api/v1/networkDevices/deviceControllability/settings`.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-network-devices-to-a-site
*/
func (s *SiteDesignService) AssignNetworkDevicesToASite(requestSiteDesignAssignNetworkDevicesToASite *RequestSiteDesignAssignNetworkDevicesToASite) (*ResponseSiteDesignAssignNetworkDevicesToASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/assignToSite/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignNetworkDevicesToASite).
		SetResult(&ResponseSiteDesignAssignNetworkDevicesToASite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignNetworkDevicesToASite(requestSiteDesignAssignNetworkDevicesToASite)
		}

		return nil, response, fmt.Errorf("error with operation AssignNetworkDevicesToASite")
	}

	result := response.Result().(*ResponseSiteDesignAssignNetworkDevicesToASite)
	return result, response, err

}

//UnassignNetworkDevicesFromSites Unassign network devices from sites - 08a6-1a87-44eb-8606
/* Unassign unprovisioned network devices from their site. If device controllability is enabled, it will be triggered once device unassigned from site successfully. Device Controllability can be enabled/disabled using `/dna/intent/api/v1/networkDevices/deviceControllability/settings`.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassign-network-devices-from-sites
*/
func (s *SiteDesignService) UnassignNetworkDevicesFromSites(requestSiteDesignUnassignNetworkDevicesFromSites *RequestSiteDesignUnassignNetworkDevicesFromSites) (*ResponseSiteDesignUnassignNetworkDevicesFromSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/unassignFromSite/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUnassignNetworkDevicesFromSites).
		SetResult(&ResponseSiteDesignUnassignNetworkDevicesFromSites{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignNetworkDevicesFromSites(requestSiteDesignUnassignNetworkDevicesFromSites)
		}

		return nil, response, fmt.Errorf("error with operation UnassignNetworkDevicesFromSites")
	}

	result := response.Result().(*ResponseSiteDesignUnassignNetworkDevicesFromSites)
	return result, response, err

}

//AssignANetworkProfileForSitesToTheGivenSite Assign a network profile for sites to the given site - 40ba-89da-420a-a21b
/* Assigns a given network profile for sites to a given site. Also assigns the profile to child sites.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-a-network-profile-for-sites-to-the-given-site
*/
func (s *SiteDesignService) AssignANetworkProfileForSitesToTheGivenSite(profileID string, requestSiteDesignAssignANetworkProfileForSitesToTheGivenSite *RequestSiteDesignAssignANetworkProfileForSitesToTheGivenSite) (*ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignANetworkProfileForSitesToTheGivenSite).
		SetResult(&ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignANetworkProfileForSitesToTheGivenSite(profileID, requestSiteDesignAssignANetworkProfileForSitesToTheGivenSite)
		}

		return nil, response, fmt.Errorf("error with operation AssignANetworkProfileForSitesToTheGivenSite")
	}

	result := response.Result().(*ResponseSiteDesignAssignANetworkProfileForSitesToTheGivenSite)
	return result, response, err

}

//AssignANetworkProfileForSitesToAListOfSites Assign a network profile for sites to a list of sites - 6ab6-e992-451b-8bc9
/* Assign a network profile for sites to a list of sites. Also assigns the profile to child sites.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-a-network-profile-for-sites-to-a-list-of-sites
*/
func (s *SiteDesignService) AssignANetworkProfileForSitesToAListOfSites(profileID string, requestSiteDesignAssignANetworkProfileForSitesToAListOfSites *RequestSiteDesignAssignANetworkProfileForSitesToAListOfSites) (*ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/bulk"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignANetworkProfileForSitesToAListOfSites).
		SetResult(&ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSites{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignANetworkProfileForSitesToAListOfSites(profileID, requestSiteDesignAssignANetworkProfileForSitesToAListOfSites)
		}

		return nil, response, fmt.Errorf("error with operation AssignANetworkProfileForSitesToAListOfSites")
	}

	result := response.Result().(*ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSites)
	return result, response, err

}

//Associate Associate - 308e-195d-403a-bbd4
/* Associate Site to a Network Profile


@param networkProfileID networkProfileId path parameter. Network-Profile Id to be associated

@param siteID siteId path parameter. Site Id to be associated


Documentation Link: https://developer.cisco.com/docs/dna-center/#!associate
*/
func (s *SiteDesignService) Associate(networkProfileID string, siteID string) (*ResponseSiteDesignAssociate, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkprofile/{networkProfileId}/site/{siteId}"
	path = strings.Replace(path, "{networkProfileId}", fmt.Sprintf("%v", networkProfileID), -1)
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignAssociate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.Associate(networkProfileID, siteID)
		}

		return nil, response, fmt.Errorf("error with operation Associate")
	}

	result := response.Result().(*ResponseSiteDesignAssociate)
	return result, response, err

}

//CreateSites Create sites - efac-69a1-4c2a-9d5e
/* Create area/building/floor together in bulk. If site already exist, then that will be ignored. Sites in the request payload need not to be ordered.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sites
*/
func (s *SiteDesignService) CreateSites(requestSiteDesignCreateSites *RequestSiteDesignCreateSites) (*ResponseSiteDesignCreateSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/bulk"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreateSites).
		SetResult(&ResponseSiteDesignCreateSites{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSites(requestSiteDesignCreateSites)
		}

		return nil, response, fmt.Errorf("error with operation CreateSites")
	}

	result := response.Result().(*ResponseSiteDesignCreateSites)
	return result, response, err

}

//CreatesABuildingV2 Creates a building - 73ae-3922-466b-adb1
/* Creates a building in the network hierarchy under area.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-building-v2
*/
func (s *SiteDesignService) CreatesABuildingV2(requestSiteDesignCreatesABuildingV2 *RequestSiteDesignCreatesABuildingV2) (*ResponseSiteDesignCreatesABuildingV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesABuildingV2).
		SetResult(&ResponseSiteDesignCreatesABuildingV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesABuildingV2(requestSiteDesignCreatesABuildingV2)
		}

		return nil, response, fmt.Errorf("error with operation CreatesABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignCreatesABuildingV2)
	return result, response, err

}

//CreatesAFloorV2 Creates a floor - 8882-b8fb-450a-8528
/* Create a floor in the network hierarchy under building.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-floor-v2
*/
func (s *SiteDesignService) CreatesAFloorV2(requestSiteDesignCreatesAFloorV2 *RequestSiteDesignCreatesAFloorV2) (*ResponseSiteDesignCreatesAFloorV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesAFloorV2).
		SetResult(&ResponseSiteDesignCreatesAFloorV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAFloorV2(requestSiteDesignCreatesAFloorV2)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignCreatesAFloorV2)
	return result, response, err

}

//EditTheAccessPointsPositionsV2 Edit the Access Points Positions - dba0-5a1b-495b-9520
/* Position or reposition the Access Points on the map.


@param floorID floorId path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!edit-the-access-points-positions-v2
*/
func (s *SiteDesignService) EditTheAccessPointsPositionsV2(floorID string, requestSiteDesignEditTheAccessPointsPositionsV2 *RequestSiteDesignEditTheAccessPointsPositionsV2) (*ResponseSiteDesignEditTheAccessPointsPositionsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/accessPointPositions/bulkChange"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignEditTheAccessPointsPositionsV2).
		SetResult(&ResponseSiteDesignEditTheAccessPointsPositionsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditTheAccessPointsPositionsV2(floorID, requestSiteDesignEditTheAccessPointsPositionsV2)
		}

		return nil, response, fmt.Errorf("error with operation EditTheAccessPointsPositionsV2")
	}

	result := response.Result().(*ResponseSiteDesignEditTheAccessPointsPositionsV2)
	return result, response, err

}

//AssignPlannedAccessPointsToOperationsOnesV2 Assign Planned Access Points to operations ones - 0c9d-f92d-4aab-b8c1
/* Assign Planned Access Points to operations ones.


@param floorID floorId path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-planned-access-points-to-operations-ones-v2
*/
func (s *SiteDesignService) AssignPlannedAccessPointsToOperationsOnesV2(floorID string, requestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 *RequestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2) (*ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/plannedAccessPointPositions/assignAccessPointPositions"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2).
		SetResult(&ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignPlannedAccessPointsToOperationsOnesV2(floorID, requestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2)
		}

		return nil, response, fmt.Errorf("error with operation AssignPlannedAccessPointsToOperationsOnesV2")
	}

	result := response.Result().(*ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2)
	return result, response, err

}

//AddPlannedAccessPointsPositionsV2 Add Planned Access Points Positions - b689-cbbc-4938-a234
/* Add Planned Access Points Positions on the map.


@param floorID floorId path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-planned-access-points-positions-v2
*/
func (s *SiteDesignService) AddPlannedAccessPointsPositionsV2(floorID string, requestSiteDesignAddPlannedAccessPointsPositionsV2 *RequestSiteDesignAddPlannedAccessPointsPositionsV2) (*ResponseSiteDesignAddPlannedAccessPointsPositionsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/plannedAccessPointPositions/bulk"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignAddPlannedAccessPointsPositionsV2).
		SetResult(&ResponseSiteDesignAddPlannedAccessPointsPositionsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPlannedAccessPointsPositionsV2(floorID, requestSiteDesignAddPlannedAccessPointsPositionsV2)
		}

		return nil, response, fmt.Errorf("error with operation AddPlannedAccessPointsPositionsV2")
	}

	result := response.Result().(*ResponseSiteDesignAddPlannedAccessPointsPositionsV2)
	return result, response, err

}

//EditPlannedAccessPointsPositionsV2 Edit Planned Access Points Positions - 189f-8bc9-4d88-a44b
/* Edit Planned Access Points Positions on the map.


@param floorID floorId path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!edit-planned-access-points-positions-v2
*/
func (s *SiteDesignService) EditPlannedAccessPointsPositionsV2(floorID string, requestSiteDesignEditPlannedAccessPointsPositionsV2 *RequestSiteDesignEditPlannedAccessPointsPositionsV2) (*ResponseSiteDesignEditPlannedAccessPointsPositionsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{floorId}/plannedAccessPointPositions/bulkChange"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignEditPlannedAccessPointsPositionsV2).
		SetResult(&ResponseSiteDesignEditPlannedAccessPointsPositionsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditPlannedAccessPointsPositionsV2(floorID, requestSiteDesignEditPlannedAccessPointsPositionsV2)
		}

		return nil, response, fmt.Errorf("error with operation EditPlannedAccessPointsPositionsV2")
	}

	result := response.Result().(*ResponseSiteDesignEditPlannedAccessPointsPositionsV2)
	return result, response, err

}

//UploadsFloorImageV2 Uploads floor image - fca4-5804-4758-98d2
/* Uploads floor image.


@param id id path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!uploads-floor-image-v2
*/
func (s *SiteDesignService) UploadsFloorImageV2(id string) (*resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}/uploadImage"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

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
			return s.UploadsFloorImageV2(id)
		}

		return response, fmt.Errorf("error with operation UploadsFloorImageV2")
	}

	return response, err

}

//UpdatesAnArea Updates an area - fab7-a965-4599-885f
/* Updates an area in the network hierarchy.


@param id id path parameter. Area Id

*/
func (s *SiteDesignService) UpdatesAnArea(id string, requestSiteDesignUpdatesAnArea *RequestSiteDesignUpdatesAnArea) (*ResponseSiteDesignUpdatesAnArea, *resty.Response, error) {
	path := "/dna/intent/api/v1/areas/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesAnArea).
		SetResult(&ResponseSiteDesignUpdatesAnArea{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAnArea(id, requestSiteDesignUpdatesAnArea)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAnArea")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesAnArea)
	return result, response, err

}

//UpdateDeviceControllabilitySettings Update device controllability settings - f58f-e8d6-4b88-912c
/* Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs to manage devices. Changes are made on network devices during discovery, when adding a device to Inventory, or when assigning a device to a site. If changes are made to any settings that are under the scope of this process, these changes are applied to the network devices during the Provision and Update Telemetry Settings operations, even if Device Controllability is disabled. The following device settings will be enabled as part of Device Controllability when devices are discovered.


  SNMP Credentials.
  NETCONF Credentials.

Subsequent to discovery, devices will be added to Inventory. The following device settings will be enabled when devices are added to inventory.


  Cisco TrustSec (CTS) Credentials.

The following device settings will be enabled when devices are assigned to a site. Some of these settings can be defined at a site level under Design > Network Settings > Telemetry & Wireless.


  Wired Endpoint Data Collection Enablement.
  Controller Certificates.
  SNMP Trap Server Definitions.
  Syslog Server Definitions.
  Application Visibility.
  Application QoS Policy.
  Wireless Service Assurance (WSA).
  Wireless Telemetry.
  DTLS Ciphersuite.
  AP Impersonation.

If Device Controllability is disabled, Catalyst Center does not configure any of the preceding credentials or settings on devices during discovery, at runtime, or during site assignment. However, the telemetry settings and related configuration are pushed when the device is provisioned or when the update Telemetry Settings action is performed.
Catalyst Center identifies and automatically corrects the following telemetry configuration issues on the device.


  SWIM certificate issue.
  IOS WLC NA certificate issue.
  PKCS12 certificate issue.
  IOS telemetry configuration issue.

The autocorrect telemetry config feature is supported only when Device Controllability is enabled.


*/
func (s *SiteDesignService) UpdateDeviceControllabilitySettings(requestSiteDesignUpdateDeviceControllabilitySettings *RequestSiteDesignUpdateDeviceControllabilitySettings) (*ResponseSiteDesignUpdateDeviceControllabilitySettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDevices/deviceControllability/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdateDeviceControllabilitySettings).
		SetResult(&ResponseSiteDesignUpdateDeviceControllabilitySettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceControllabilitySettings(requestSiteDesignUpdateDeviceControllabilitySettings)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceControllabilitySettings")
	}

	result := response.Result().(*ResponseSiteDesignUpdateDeviceControllabilitySettings)
	return result, response, err

}

//UpdatesABuildingV2 Updates a building - 0280-1b95-4d78-845d
/* Updates a building in the network hierarchy.


@param id id path parameter. Building Id

*/
func (s *SiteDesignService) UpdatesABuildingV2(id string, requestSiteDesignUpdatesABuildingV2 *RequestSiteDesignUpdatesABuildingV2) (*ResponseSiteDesignUpdatesABuildingV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesABuildingV2).
		SetResult(&ResponseSiteDesignUpdatesABuildingV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesABuildingV2(id, requestSiteDesignUpdatesABuildingV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesABuildingV2)
	return result, response, err

}

//UpdatesFloorSettingsV2 Updates floor settings - 16b8-59a5-4b38-8cb7
/* Updates UI user preference for floor unit system. Unit sytem change will effect for all floors across all sites.


 */
func (s *SiteDesignService) UpdatesFloorSettingsV2(requestSiteDesignUpdatesFloorSettingsV2 *RequestSiteDesignUpdatesFloorSettingsV2) (*ResponseSiteDesignUpdatesFloorSettingsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesFloorSettingsV2).
		SetResult(&ResponseSiteDesignUpdatesFloorSettingsV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesFloorSettingsV2(requestSiteDesignUpdatesFloorSettingsV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesFloorSettingsV2")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesFloorSettingsV2)
	return result, response, err

}

//UpdatesAFloorV2 Updates a floor - ee9f-c9f2-4a3a-b7a5
/* Updates a floor in the network hierarchy.


@param id id path parameter. Floor Id

*/
func (s *SiteDesignService) UpdatesAFloorV2(id string, requestSiteDesignUpdatesAFloorV2 *RequestSiteDesignUpdatesAFloorV2) (*ResponseSiteDesignUpdatesAFloorV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesAFloorV2).
		SetResult(&ResponseSiteDesignUpdatesAFloorV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAFloorV2(id, requestSiteDesignUpdatesAFloorV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesAFloorV2)
	return result, response, err

}

//DeletesAnArea Deletes an area - 2ba4-c8f4-4ec8-b80e
/* Deletes an area in the network hierarchy. This operations fails if there are any child areas or buildings for this area.


@param id id path parameter. Area ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-an-area
*/
func (s *SiteDesignService) DeletesAnArea(id string) (*ResponseSiteDesignDeletesAnArea, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/areas/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesAnArea{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAnArea(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesAnArea")
	}

	result := response.Result().(*ResponseSiteDesignDeletesAnArea)
	return result, response, err

}

//DeletesANetworkProfileForSites Deletes a network profile for sites - a48b-b959-493b-93d1
/* Deletes a network profile for sites.


@param id id path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-network-profile-for-sites
*/
func (s *SiteDesignService) DeletesANetworkProfileForSites(id string) (*ResponseSiteDesignDeletesANetworkProfileForSites, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/networkProfilesForSites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesANetworkProfileForSites{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesANetworkProfileForSites(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesANetworkProfileForSites")
	}

	result := response.Result().(*ResponseSiteDesignDeletesANetworkProfileForSites)
	return result, response, err

}

//UnassignsANetworkProfileForSitesFromMultipleSites Unassigns a network profile for sites from multiple sites - ec90-f973-435b-a6f5
/* Unassigns a given network profile for sites from multiple sites. The profile must be removed from the containing building first if this site is a floor.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`

@param UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassigns-a-network-profile-for-sites-from-multiple-sites
*/
func (s *SiteDesignService) UnassignsANetworkProfileForSitesFromMultipleSites(profileID string, UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams *UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams) (*ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSites, *resty.Response, error) {
	//profileID string,UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams *UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/bulk"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	queryString, _ := query.Values(UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSites{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignsANetworkProfileForSitesFromMultipleSites(
				profileID, UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UnassignsANetworkProfileForSitesFromMultipleSites")
	}

	result := response.Result().(*ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSites)
	return result, response, err

}

//UnassignsANetworkProfileForSitesFromASite Unassigns a network profile for sites from a site - 8c94-0a8e-450b-98cc
/* Unassigns a given network profile for sites from a site. The profile must be removed from parent sites first, otherwise this operation will not ulimately  unassign the profile.


@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`

@param id id path parameter. The `id` of the site, retrievable from `GET /intent/api/v1/networkProfilesForSites/{id}/siteAssignments`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassigns-a-network-profile-for-sites-from-a-site
*/
func (s *SiteDesignService) UnassignsANetworkProfileForSitesFromASite(profileID string, id string) (*ResponseSiteDesignUnassignsANetworkProfileForSitesFromASite, *resty.Response, error) {
	//profileID string,id string
	path := "/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments/{id}"
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignUnassignsANetworkProfileForSitesFromASite{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignsANetworkProfileForSitesFromASite(
				profileID, id)
		}
		return nil, response, fmt.Errorf("error with operation UnassignsANetworkProfileForSitesFromASite")
	}

	result := response.Result().(*ResponseSiteDesignUnassignsANetworkProfileForSitesFromASite)
	return result, response, err

}

//Disassociate Disassociate - e687-58d2-4b19-b5c6
/* Disassociate a Site from a Network Profile


@param networkProfileID networkProfileId path parameter. Network-Profile Id to be associated

@param siteID siteId path parameter. Site Id to be associated


Documentation Link: https://developer.cisco.com/docs/dna-center/#!disassociate
*/
func (s *SiteDesignService) Disassociate(networkProfileID string, siteID string) (*ResponseSiteDesignDisassociate, *resty.Response, error) {
	//networkProfileID string,siteID string
	path := "/dna/intent/api/v1/networkprofile/{networkProfileId}/site/{siteId}"
	path = strings.Replace(path, "{networkProfileId}", fmt.Sprintf("%v", networkProfileID), -1)
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDisassociate{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Disassociate(
				networkProfileID, siteID)
		}
		return nil, response, fmt.Errorf("error with operation Disassociate")
	}

	result := response.Result().(*ResponseSiteDesignDisassociate)
	return result, response, err

}

//DeletesABuildingV2 Deletes a building - 45ae-a9e4-4008-b0b6
/* Deletes building in the network hierarchy. This operations fails if there are any floors for this building, or if there are any devices assigned to this building.


@param id id path parameter. Building ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-building-v2
*/
func (s *SiteDesignService) DeletesABuildingV2(id string) (*ResponseSiteDesignDeletesABuildingV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesABuildingV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesABuildingV2(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesABuildingV2")
	}

	result := response.Result().(*ResponseSiteDesignDeletesABuildingV2)
	return result, response, err

}

//DeletePlannedAccessPointsPositionV2 Delete Planned Access Points Position - fda5-0ba1-462a-ae58
/* Delete specified Planned Access Points Position designated for a specific floor.


@param floorID floorId path parameter. Floor Id

@param id id path parameter. Planned Access Point Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-planned-access-points-position-v2
*/
func (s *SiteDesignService) DeletePlannedAccessPointsPositionV2(floorID string, id string) (*ResponseSiteDesignDeletePlannedAccessPointsPositionV2, *resty.Response, error) {
	//floorID string,id string
	path := "/dna/intent/api/v2/floors/{floorId}/plannedAccessPointPositions/{id}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletePlannedAccessPointsPositionV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePlannedAccessPointsPositionV2(
				floorID, id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePlannedAccessPointsPositionV2")
	}

	result := response.Result().(*ResponseSiteDesignDeletePlannedAccessPointsPositionV2)
	return result, response, err

}

//DeletesAFloorV2 Deletes a floor - 6cb4-884b-47db-a808
/* Deletes a floor from the network hierarchy. This operations fails if there are any devices assigned to this floor.


@param id id path parameter. Floor ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-floor-v2
*/
func (s *SiteDesignService) DeletesAFloorV2(id string) (*ResponseSiteDesignDeletesAFloorV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesAFloorV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAFloorV2(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesAFloorV2")
	}

	result := response.Result().(*ResponseSiteDesignDeletesAFloorV2)
	return result, response, err

}
