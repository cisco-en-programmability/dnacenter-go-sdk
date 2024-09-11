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
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type GetSiteAssignedNetworkDevicesCountQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site Id. It must be area Id or building Id or floor Id.
}
type GetSiteNotAssignedNetworkDevicesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type RetrievesTheListOfNetworkProfilesForSitesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
	Type   string  `url:"type,omitempty"`   //Filter responses to only include profiles of a given type
}
type RetrievesTheCountOfNetworkProfilesForSitesQueryParams struct {
	Type string `url:"type,omitempty"` //Filter the response to only count profiles of a given type
}
type RetrievesTheListOfSitesThatTheGivenNetworkProfileForSitesIsAssignedToQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //The 'id' of the site, retrievable from 'GET /intent/api/v1/sites'
}
type GetSitesQueryParams struct {
	Name           string `url:"name,omitempty"`            //Site name.
	NameHierarchy  string `url:"nameHierarchy,omitempty"`   //Site name hierarchy.
	Type           string `url:"type,omitempty"`            //Site type.
	UnitsOfMeasure string `url:"_unitsOfMeasure,omitempty"` //Floor units of measure
	Offset         int    `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1.
	Limit          int    `url:"limit,omitempty"`           //The number of records to show for this page.
}
type GetSitesCountQueryParams struct {
	Name string `url:"name,omitempty"` //Site name.
}
type RetrievesTheListOfNetworkProfilesThatTheGivenSiteHasBeenAssignedQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type GetsAFloorQueryParams struct {
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
	ID            string `json:"id,omitempty"`            // Aread Id. Read only.
	Name          string `json:"name,omitempty"`          // Area name
	NameHierarchy string `json:"nameHierarchy,omitempty"` // Area hierarchical name. Read only.
	ParentID      string `json:"parentId,omitempty"`      // Parent Id
	Type          string `json:"type,omitempty"`          // Site Type.
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
	NameHierarchy  string   `json:"nameHierarchy,omitempty"`  // Site hierarchical name. Read only. Example: Global/USA/San Jose/Building1
	Name           string   `json:"name,omitempty"`           // Site name.
	Latitude       *float64 `json:"latitude,omitempty"`       // Building Latitude. Example: 37.403712
	Longitude      *float64 `json:"longitude,omitempty"`      // Building Longitude. Example: -121.971063
	Address        string   `json:"address,omitempty"`        // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country        string   `json:"country,omitempty"`        // Country name for the building.
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // Floor RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Floor unit of measure
	Type           string   `json:"type,omitempty"`           // Type
	ID             string   `json:"id,omitempty"`             // Site Id. Read only.
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id. Read only
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
type ResponseSiteDesignCreatesABuilding struct {
	Version  string                                      `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesABuildingResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesABuildingResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesABuilding struct {
	Version  string                                      `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesABuildingResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesABuildingResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignDeletesABuilding struct {
	Version  string                                      `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesABuildingResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesABuildingResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsABuilding struct {
	Response *ResponseSiteDesignGetsABuildingResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignGetsABuildingResponse struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country   string   `json:"country,omitempty"`   // Country name
	Type      string   `json:"type,omitempty"`      // Example: building
}
type ResponseSiteDesignCreatesAFloor struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignCreatesAFloorResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignCreatesAFloorResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignUpdatesFloorSettings struct {
	Version  string                                          `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesFloorSettingsResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesFloorSettingsResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetFloorSettings struct {
	Response *ResponseSiteDesignGetFloorSettingsResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version
}
type ResponseSiteDesignGetFloorSettingsResponse struct {
	UnitsOfMeasure string `json:"unitsOfMeasure,omitempty"` // Floor units of measure.
}
type ResponseSiteDesignUpdatesAFloor struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignUpdatesAFloorResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignUpdatesAFloorResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSiteDesignGetsAFloor struct {
	Response *ResponseSiteDesignGetsAFloorResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignGetsAFloorResponse struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id.
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
	Type           string   `json:"type,omitempty"`           // Example : floor
	ID             string   `json:"id,omitempty"`             // Floor Id. Read only.
	NameHierarchy  string   `json:"nameHierarchy,omitempty"`  // Floor hierarchical name. Read only.
}
type ResponseSiteDesignDeletesAFloor struct {
	Version  string                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSiteDesignDeletesAFloorResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignDeletesAFloorResponse struct {
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
	DeviceIDs []string `json:"deviceIds,omitempty"` // Unassigned network devices.
	SiteID    string   `json:"siteId,omitempty"`    // This must be building Id or floor Id. Access points, Sensors are assigned to floor. Remaining network devices are assigned to building. Site Id can be retrieved using '/intent/api/v1/sites'.
}
type RequestSiteDesignUpdateDeviceControllabilitySettings struct {
	AutocorrectTelemetryConfig *bool `json:"autocorrectTelemetryConfig,omitempty"` // If it is true, autocorrect telemetry config is enabled. If it is false, autocorrect telemetry config is disabled. The autocorrect telemetry config feature is supported only when device controllability is enabled.
	DeviceControllability      *bool `json:"deviceControllability,omitempty"`      // If it is true, device controllability is enabled. If it is false, device controllability is disabled.
}
type RequestSiteDesignUnassignNetworkDevicesFromSites struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // Network device Ids.
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
type RequestSiteDesignCreatesABuilding struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country   string   `json:"country,omitempty"`   // Country name
}
type RequestSiteDesignUpdatesABuilding struct {
	ParentID  string   `json:"parentId,omitempty"`  // Parent Id
	Name      string   `json:"name,omitempty"`      // Building name
	Latitude  *float64 `json:"latitude,omitempty"`  // Building Latitude. Example: 37.403712
	Longitude *float64 `json:"longitude,omitempty"` // Building Longitude. Example: -121.971063
	Address   string   `json:"address,omitempty"`   // Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
	Country   string   `json:"country,omitempty"`   // Country name
}
type RequestSiteDesignCreatesAFloor struct {
	ParentID       string   `json:"parentId,omitempty"`       // Parent Id
	Name           string   `json:"name,omitempty"`           // Floor name
	FloorNumber    *int     `json:"floorNumber,omitempty"`    // Floor number
	RfModel        string   `json:"rfModel,omitempty"`        // RF Model
	Width          *float64 `json:"width,omitempty"`          // Floor width. Example : 100.5
	Length         *float64 `json:"length,omitempty"`         // Floor length. Example : 110.3
	Height         *float64 `json:"height,omitempty"`         // Floor height. Example : 10.1
	UnitsOfMeasure string   `json:"unitsOfMeasure,omitempty"` // Units Of Measure
}
type RequestSiteDesignUpdatesFloorSettings struct {
	UnitsOfMeasure string `json:"unitsOfMeasure,omitempty"` // Floor units of measure
}
type RequestSiteDesignUpdatesAFloor struct {
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


@param id id path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'


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


@param profileID profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'

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


@param profileID profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'


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
These assigments can be modified via the '/dna/intent/api/v1/networkProfilesForSites/{profileId}/siteAssignments' resources.


@param siteID siteId path parameter. The 'id' of the site, retrievable from '/dna/intent/api/v1/sites'

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


@param siteID siteId path parameter. The 'id' of the site, retrievable from '/dna/intent/api/v1/sites'


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

//GetsABuilding Gets a building - e293-295e-4a78-bf64
/* Gets a building in the network hierarchy.


@param id id path parameter. Building Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-building
*/
func (s *SiteDesignService) GetsABuilding(id string) (*ResponseSiteDesignGetsABuilding, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetsABuilding{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsABuilding(id)
		}
		return nil, response, fmt.Errorf("error with operation GetsABuilding")
	}

	result := response.Result().(*ResponseSiteDesignGetsABuilding)
	return result, response, err

}

//GetFloorSettings Get floor settings - f697-a95f-4469-958f
/* Gets UI user preference for floor unit system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-floor-settings
*/
func (s *SiteDesignService) GetFloorSettings() (*ResponseSiteDesignGetFloorSettings, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignGetFloorSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFloorSettings()
		}
		return nil, response, fmt.Errorf("error with operation GetFloorSettings")
	}

	result := response.Result().(*ResponseSiteDesignGetFloorSettings)
	return result, response, err

}

//GetsAFloor Gets a floor - ff92-2958-4bba-9288
/* Gets a floor in the network hierarchy.


@param id id path parameter. Floor Id

@param GetsAFloorQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-floor
*/
func (s *SiteDesignService) GetsAFloor(id string, GetsAFloorQueryParams *GetsAFloorQueryParams) (*ResponseSiteDesignGetsAFloor, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetsAFloorQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetsAFloor{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAFloor(id, GetsAFloorQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsAFloor")
	}

	result := response.Result().(*ResponseSiteDesignGetsAFloor)
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
/* Assign unprovisioned network devices to a site. Along with that it can also be used to assign unprovisioned network devices to a different site. If device controllability is enabled, it will be triggered once device assigned to site successfully. Device Controllability can be enabled/disabled using '/dna/intent/api/v1/networkDevices/deviceControllability/settings'.



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
/* Unassign unprovisioned network devices from their site. If device controllability is enabled, it will be triggered once device unassigned from site successfully. Device Controllability can be enabled/disabled using '/dna/intent/api/v1/networkDevices/deviceControllability/settings'.



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


@param profileID profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'


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


@param profileID profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'


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

//CreatesABuilding Creates a building - 73ae-3922-466b-adb1
/* Creates a building in the network hierarchy under area.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-building
*/
func (s *SiteDesignService) CreatesABuilding(requestSiteDesignCreatesABuilding *RequestSiteDesignCreatesABuilding) (*ResponseSiteDesignCreatesABuilding, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesABuilding).
		SetResult(&ResponseSiteDesignCreatesABuilding{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesABuilding(requestSiteDesignCreatesABuilding)
		}

		return nil, response, fmt.Errorf("error with operation CreatesABuilding")
	}

	result := response.Result().(*ResponseSiteDesignCreatesABuilding)
	return result, response, err

}

//CreatesAFloor Creates a floor - 8882-b8fb-450a-8528
/* Create a floor in the network hierarchy under building.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-floor
*/
func (s *SiteDesignService) CreatesAFloor(requestSiteDesignCreatesAFloor *RequestSiteDesignCreatesAFloor) (*ResponseSiteDesignCreatesAFloor, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreatesAFloor).
		SetResult(&ResponseSiteDesignCreatesAFloor{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAFloor(requestSiteDesignCreatesAFloor)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAFloor")
	}

	result := response.Result().(*ResponseSiteDesignCreatesAFloor)
	return result, response, err

}

//UploadsFloorImage Uploads floor image - fca4-5804-4758-98d2
/* Uploads floor image.


@param id id path parameter. Floor Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!uploads-floor-image
*/
func (s *SiteDesignService) UploadsFloorImage(id string) (*resty.Response, error) {
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
			return s.UploadsFloorImage(id)
		}

		return response, fmt.Errorf("error with operation UploadsFloorImage")
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
/* Device Controllability is a system-level process on Catalyst Center that enforces state synchronization for some device-layer features. Its purpose is to aid in the deployment of required network settings that Catalyst Center needs to manage devices. Changes are made on network devices  during discovery, when adding a device to Inventory, or when assigning a device to a site. If changes  are made to any settings that are under the scope of this process, these changes are applied to the  network devices during the Provision and Update Telemetry Settings operations, even if Device  Controllability is disabled. The following device settings will be enabled as part of  Device Controllability when devices are discovered.

  SNMP Credentials.
  NETCONF Credentials.

Subsequent to discovery, devices will be added to Inventory. The following device settings will be  enabled when devices are added to inventory.

  Cisco TrustSec (CTS) Credentials.

The following device settings will be enabled when devices are assigned to a site. Some of these  settings can be defined at a site level under Design > Network Settings > Telemetry & Wireless.

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

If Device Controllability is disabled, Catalyst Center does not configure any of the preceding  credentials or settings on devices during discovery, at runtime, or during site assignment. However,  the telemetry settings and related configuration are pushed when the device is provisioned or when the  update Telemetry Settings action is performed.
Catalyst Center identifies and automatically corrects the following telemetry configuration issues on  the device.

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

//UpdatesABuilding Updates a building - 0280-1b95-4d78-845d
/* Updates a building in the network hierarchy.


@param id id path parameter. Building Id

*/
func (s *SiteDesignService) UpdatesABuilding(id string, requestSiteDesignUpdatesABuilding *RequestSiteDesignUpdatesABuilding) (*ResponseSiteDesignUpdatesABuilding, *resty.Response, error) {
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesABuilding).
		SetResult(&ResponseSiteDesignUpdatesABuilding{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesABuilding(id, requestSiteDesignUpdatesABuilding)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesABuilding")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesABuilding)
	return result, response, err

}

//UpdatesFloorSettings Updates floor settings - 16b8-59a5-4b38-8cb7
/* Updates UI user preference for floor unit system. Unit sytem change will effect for all floors across all sites.


 */
func (s *SiteDesignService) UpdatesFloorSettings(requestSiteDesignUpdatesFloorSettings *RequestSiteDesignUpdatesFloorSettings) (*ResponseSiteDesignUpdatesFloorSettings, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesFloorSettings).
		SetResult(&ResponseSiteDesignUpdatesFloorSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesFloorSettings(requestSiteDesignUpdatesFloorSettings)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesFloorSettings")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesFloorSettings)
	return result, response, err

}

//UpdatesAFloor Updates a floor - ee9f-c9f2-4a3a-b7a5
/* Updates a floor in the network hierarchy.


@param id id path parameter. Floor Id

*/
func (s *SiteDesignService) UpdatesAFloor(id string, requestSiteDesignUpdatesAFloor *RequestSiteDesignUpdatesAFloor) (*ResponseSiteDesignUpdatesAFloor, *resty.Response, error) {
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdatesAFloor).
		SetResult(&ResponseSiteDesignUpdatesAFloor{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatesAFloor(id, requestSiteDesignUpdatesAFloor)
		}
		return nil, response, fmt.Errorf("error with operation UpdatesAFloor")
	}

	result := response.Result().(*ResponseSiteDesignUpdatesAFloor)
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
			return s.DeletesAnArea(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesAnArea")
	}

	result := response.Result().(*ResponseSiteDesignDeletesAnArea)
	return result, response, err

}

//DeletesANetworkProfileForSites Deletes a network profile for sites - a48b-b959-493b-93d1
/* Deletes a network profile for sites.


@param id id path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'


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
			return s.DeletesANetworkProfileForSites(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesANetworkProfileForSites")
	}

	result := response.Result().(*ResponseSiteDesignDeletesANetworkProfileForSites)
	return result, response, err

}

//UnassignsANetworkProfileForSitesFromMultipleSites Unassigns a network profile for sites from multiple sites - ec90-f973-435b-a6f5
/* Unassigns a given network profile for sites from multiple sites. The profile must be removed from the containing building first if this site is a floor.


@param profileID profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'

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
			return s.UnassignsANetworkProfileForSitesFromMultipleSites(profileID, UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UnassignsANetworkProfileForSitesFromMultipleSites")
	}

	result := response.Result().(*ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSites)
	return result, response, err

}

//UnassignsANetworkProfileForSitesFromASite Unassigns a network profile for sites from a site - 8c94-0a8e-450b-98cc
/* Unassigns a given network profile for sites from a site. The profile must be removed from parent sites first, otherwise this operation will not ulimately  unassign the profile.


@param profileID profileId path parameter. The 'id' of the network profile, retrievable from 'GET /intent/api/v1/networkProfilesForSites'

@param id id path parameter. The 'id' of the site, retrievable from 'GET /intent/api/v1/networkProfilesForSites/{id}/siteAssignments'


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
			return s.UnassignsANetworkProfileForSitesFromASite(profileID, id)
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
			return s.Disassociate(networkProfileID, siteID)
		}
		return nil, response, fmt.Errorf("error with operation Disassociate")
	}

	result := response.Result().(*ResponseSiteDesignDisassociate)
	return result, response, err

}

//DeletesABuilding Deletes a building - 45ae-a9e4-4008-b0b6
/* Deletes building in the network hierarchy. This operations fails if there are any floors for this building, or if there are any devices assigned to this building.


@param id id path parameter. Building ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-building
*/
func (s *SiteDesignService) DeletesABuilding(id string) (*ResponseSiteDesignDeletesABuilding, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/buildings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesABuilding{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesABuilding(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesABuilding")
	}

	result := response.Result().(*ResponseSiteDesignDeletesABuilding)
	return result, response, err

}

//DeletesAFloor Deletes a floor - 6cb4-884b-47db-a808
/* Deletes a floor from the network hierarchy. This operations fails if there are any devices assigned to this floor.


@param id id path parameter. Floor ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-a-floor
*/
func (s *SiteDesignService) DeletesAFloor(id string) (*ResponseSiteDesignDeletesAFloor, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/floors/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignDeletesAFloor{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesAFloor(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletesAFloor")
	}

	result := response.Result().(*ResponseSiteDesignDeletesAFloor)
	return result, response, err

}
