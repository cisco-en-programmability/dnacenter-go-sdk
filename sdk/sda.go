package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SdaService service

type ReadListOfFabricSitesWithTheirHealthSummaryQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit     float64 `url:"limit,omitempty"`     //Maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
	ID        string  `url:"id,omitempty"`        //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	Attribute string  `url:"attribute,omitempty"` //The list of FabricSite health attributes. Please refer to ```fabricSiteAttributes``` section in the Open API specification document mentioned in the description.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. A maximum of 3 views can be queried at a time per request.  Please refer to ```fabricSiteViews``` section in the Open API specification document mentioned in the description.
}
type ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadFabricSiteCountQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ID        string  `url:"id,omitempty"`        //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
}
type ReadFabricSiteCountHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadFabricSitesWithHealthSummaryFromIDQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Attribute string  `url:"attribute,omitempty"` //The list of FabricSite health attributes. Please refer to ```fabricSiteAttributes``` section in the Open API specification document mentioned in the description.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. A maximum of 3 views can be queried at a time per request.  Please refer to ```fabricSiteViews``` section in the Open API specification document mentioned in the description.
}
type ReadFabricSitesWithHealthSummaryFromIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TrendInterval string  `url:"trendInterval,omitempty"` //The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Order         string  `url:"order,omitempty"`         //The sort order of the field ascending or descending.
	Attribute     string  `url:"attribute,omitempty"`     //The interested fields in the request. For valid attributes, verify the documentation.
}
type TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadFabricEntitySummaryQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type ReadFabricEntitySummaryHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadListOfTransitNetworksWithTheirHealthSummaryQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit     float64 `url:"limit,omitempty"`     //Maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
	ID        string  `url:"id,omitempty"`        //The list of transit entity ids. (Ex "1551156a-bc97-3c63-aeda-8a6d3765b5b9") Examples: id=1551156a-bc97-3c63-aeda-8a6d3765b5b9 (single entity uuid requested) id=1551156a-bc97-3c63-aeda-8a6d3765b5b9&id=4aa20652-237c-4625-b2b4-fd7e82b6a81e (multiple entity uuids with '&' separator)
	Attribute string  `url:"attribute,omitempty"` //The interested fields in the request. For valid attributes, verify the documentation.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites.
}
type ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadTransitNetworksCountQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ID        string  `url:"id,omitempty"`        //The list of transit entity ids. (Ex "1551156a-bc97-3c63-aeda-8a6d3765b5b9") Examples: id=1551156a-bc97-3c63-aeda-8a6d3765b5b9 (single entity uuid requested) id=1551156a-bc97-3c63-aeda-8a6d3765b5b9&id=4aa20652-237c-4625-b2b4-fd7e82b6a81e (multiple entity uuids with '&' separator)
}
type ReadTransitNetworksCountHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadTransitNetworkWithItsHealthSummaryFromIDQueryParams struct {
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Attribute string  `url:"attribute,omitempty"` //The interested fields in the request. For valid attributes, verify the documentation.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with sites.
}
type ReadTransitNetworkWithItsHealthSummaryFromIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeQueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TrendInterval string  `url:"trendInterval,omitempty"` //The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Order         string  `url:"order,omitempty"`         //The sort order of the field ascending or descending.
	Attribute     string  `url:"attribute,omitempty"`     //The interested fields in the request. For valid attributes, verify the documentation.
}
type TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadListOfVirtualNetworksWithTheirHealthSummaryQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit     float64 `url:"limit,omitempty"`     //Maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
	ID        string  `url:"id,omitempty"`        //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	VnLayer   string  `url:"vnLayer,omitempty"`   //VN Layer information covering Layer 3 or Layer 2 VNs.
	Attribute string  `url:"attribute,omitempty"` //The interested fields in the request. For valid attributes, verify the documentation.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with virtual networks.
}
type ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadVirtualNetworksCountQueryParams struct {
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ID        string  `url:"id,omitempty"`        //The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
	VnLayer   string  `url:"vnLayer,omitempty"`   //VN Layer information covering Layer 3 or Layer 2 VNs.
}
type ReadVirtualNetworksCountHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type ReadVirtualNetworkWithItsHealthSummaryFromIDQueryParams struct {
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Attribute string  `url:"attribute,omitempty"` //The interested fields in the request. For valid attributes, verify the documentation.
	View      string  `url:"view,omitempty"`      //The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with virtual networks.
}
type ReadVirtualNetworkWithItsHealthSummaryFromIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeQueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	TrendInterval string  `url:"trendInterval,omitempty"` //The time window to aggregate the metrics. Interval can be 5 minutes or 10 minutes or 1 hour or 1 day or 7 days
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	Order         string  `url:"order,omitempty"`         //The sort order of the field ascending or descending.
	Attribute     string  `url:"attribute,omitempty"`     //The interested fields in the request. For valid attributes, verify the documentation.
}
type TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetDefaultAuthenticationProfileFromSdaFabricQueryParams struct {
	SiteNameHierarchy        string `url:"siteNameHierarchy,omitempty"`        //siteNameHierarchy
	AuthenticateTemplateName string `url:"authenticateTemplateName,omitempty"` //authenticateTemplateName
}
type DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //siteNameHierarchy
}
type GetBorderDeviceDetailFromSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteBorderDeviceFromSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteControlPlaneDeviceInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetControlPlaneDeviceFromSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetDeviceInfoFromSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetDeviceRoleInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //Device Management IP Address
}
type DeleteEdgeDeviceFromSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetEdgeDeviceFromSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type GetSiteFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Hierarchy
}
type DeleteSiteFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Hierarchy
}
type DeletePortAssignmentForAccessPointInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type GetPortAssignmentForAccessPointInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type DeletePortAssignmentForUserDeviceInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type GetPortAssignmentForUserDeviceInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
	InterfaceName             string `url:"interfaceName,omitempty"`             //interfaceName
}
type GetMulticastDetailsFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //fabric site name hierarchy
}
type DeleteMulticastFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //siteNameHierarchy
}
type DeleteProvisionedWiredDeviceQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //Valid IP address of the device currently provisioned in a fabric site
}
type GetProvisionedWiredDeviceQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteTransitPeerNetworkQueryParams struct {
	TransitPeerNetworkName string `url:"transitPeerNetworkName,omitempty"` //Transit Peer Network Name
}
type GetTransitPeerNetworkInfoQueryParams struct {
	TransitPeerNetworkName string `url:"transitPeerNetworkName,omitempty"` //Transit or Peer Network Name
}
type DeleteVnFromSdaFabricQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
}
type GetVnFromSdaFabricQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
}
type GetVirtualNetworkSummaryQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Complete fabric siteNameHierarchy Path
}
type GetIPPoolFromSdaVirtualNetworkQueryParams struct {
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	IPPoolName         string `url:"ipPoolName,omitempty"`         //ipPoolName. Note: Use vlanName as a value for this parameter if same ip pool is assigned to multiple virtual networks (e.g.. ipPoolName=vlan1021)
}
type DeleteIPPoolFromSdaVirtualNetworkQueryParams struct {
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	IPPoolName         string `url:"ipPoolName,omitempty"`         //ipPoolName
}
type GetAnycastGatewaysQueryParams struct {
	ID                 string  `url:"id,omitempty"`                 //ID of the anycast gateway.
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric the anycast gateway is assigned to.
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the virtual network associated with the anycast gateways.
	IPPoolName         string  `url:"ipPoolName,omitempty"`         //Name of the IP pool associated with the anycast gateways.
	VLANName           string  `url:"vlanName,omitempty"`           //VLAN name of the anycast gateways.
	VLANID             float64 `url:"vlanId,omitempty"`             //VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetAnycastGatewayCountQueryParams struct {
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric the anycast gateway is assigned to.
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the virtual network associated with the anycast gateways.
	IPPoolName         string  `url:"ipPoolName,omitempty"`         //Name of the IP pool associated with the anycast gateways.
	VLANName           string  `url:"vlanName,omitempty"`           //VLAN name of the anycast gateways.
	VLANID             float64 `url:"vlanId,omitempty"`             //VLAN ID of the anycast gateways. The allowed range for vlanId is [2-4093] except for reserved VLANs [1002-1005], 2046, and 4094.
}
type GetAuthenticationProfilesQueryParams struct {
	FabricID                      string  `url:"fabricId,omitempty"`                      //ID of the fabric the authentication profile is assigned to.
	AuthenticationProfileName     string  `url:"authenticationProfileName,omitempty"`     //Return only the authentication profiles with this specified name. Note that 'No Authentication' is not a valid option for this parameter.
	IsGlobalAuthenticationProfile bool    `url:"isGlobalAuthenticationProfile,omitempty"` //Set to true to return only global authentication profiles, or set to false to hide them. isGlobalAuthenticationProfile must not be true when fabricId is provided.
	Offset                        float64 `url:"offset,omitempty"`                        //Starting record for pagination.
	Limit                         float64 `url:"limit,omitempty"`                         //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeleteExtranetPoliciesQueryParams struct {
	ExtranetPolicyName string `url:"extranetPolicyName,omitempty"` //Name of the extranet policy.
}
type GetExtranetPoliciesQueryParams struct {
	ExtranetPolicyName string  `url:"extranetPolicyName,omitempty"` //Name of the extranet policy.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetFabricDevicesQueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	DeviceRoles     string  `url:"deviceRoles,omitempty"`     //Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeleteFabricDevicesQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	DeviceRoles     string `url:"deviceRoles,omitempty"`     //Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE].
}
type GetFabricDevicesCountQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	DeviceRoles     string `url:"deviceRoles,omitempty"`     //Device roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].
}
type DeleteFabricDeviceLayer2HandoffsQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer2HandoffsQueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetFabricDevicesLayer2HandoffsCountQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type DeleteFabricDeviceLayer3HandoffsWithIPTransitQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetFabricDevicesLayer3HandoffsWithIPTransitCountQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricDevicesLayer3HandoffsWithSdaTransitCountQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric this device belongs to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the fabric device.
}
type GetFabricSitesQueryParams struct {
	ID     string  `url:"id,omitempty"`     //ID of the fabric site.
	SiteID string  `url:"siteId,omitempty"` //ID of the network hierarchy associated with the fabric site.
	Offset float64 `url:"offset,omitempty"` //Starting record for pagination.
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetFabricZonesQueryParams struct {
	ID     string  `url:"id,omitempty"`     //ID of the fabric zone.
	SiteID string  `url:"siteId,omitempty"` //ID of the network hierarchy associated with the fabric zone.
	Offset float64 `url:"offset,omitempty"` //Starting record for pagination.
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeleteLayer2VirtualNetworksQueryParams struct {
	FabricID                           string  `url:"fabricId,omitempty"`                           //ID of the fabric the layer 2 virtual network is assigned to.
	VLANName                           string  `url:"vlanName,omitempty"`                           //The vlan name of the layer 2 virtual network.
	VLANID                             float64 `url:"vlanId,omitempty"`                             //The vlan ID of the layer 2 virtual network.
	TrafficType                        string  `url:"trafficType,omitempty"`                        //The traffic type of the layer 2 virtual network.
	AssociatedLayer3VirtualNetworkName string  `url:"associatedLayer3VirtualNetworkName,omitempty"` //Name of the associated layer 3 virtual network.
}
type GetLayer2VirtualNetworksQueryParams struct {
	ID                                 string  `url:"id,omitempty"`                                 //ID of the layer 2 virtual network.
	FabricID                           string  `url:"fabricId,omitempty"`                           //ID of the fabric the layer 2 virtual network is assigned to.
	VLANName                           string  `url:"vlanName,omitempty"`                           //The vlan name of the layer 2 virtual network.
	VLANID                             float64 `url:"vlanId,omitempty"`                             //The vlan ID of the layer 2 virtual network.
	TrafficType                        string  `url:"trafficType,omitempty"`                        //The traffic type of the layer 2 virtual network.
	AssociatedLayer3VirtualNetworkName string  `url:"associatedLayer3VirtualNetworkName,omitempty"` //Name of the associated layer 3 virtual network.
	Offset                             float64 `url:"offset,omitempty"`                             //Starting record for pagination.
	Limit                              float64 `url:"limit,omitempty"`                              //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetLayer2VirtualNetworkCountQueryParams struct {
	FabricID                           string  `url:"fabricId,omitempty"`                           //ID of the fabric the layer 2 virtual network is assigned to.
	VLANName                           string  `url:"vlanName,omitempty"`                           //The vlan name of the layer 2 virtual network.
	VLANID                             float64 `url:"vlanId,omitempty"`                             //The vlan ID of the layer 2 virtual network.
	TrafficType                        string  `url:"trafficType,omitempty"`                        //The traffic type of the layer 2 virtual network.
	AssociatedLayer3VirtualNetworkName string  `url:"associatedLayer3VirtualNetworkName,omitempty"` //Name of the associated layer 3 virtual network.
}
type GetLayer3VirtualNetworksQueryParams struct {
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the layer 3 virtual network.
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric the layer 3 virtual network is assigned to.
	AnchoredSiteID     string  `url:"anchoredSiteId,omitempty"`     //Fabric ID of the fabric site the layer 3 virtual network is anchored at.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeleteLayer3VirtualNetworksQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //Name of the layer 3 virtual network.
}
type GetLayer3VirtualNetworksCountQueryParams struct {
	FabricID       string `url:"fabricId,omitempty"`       //ID of the fabric the layer 3 virtual network is assigned to.
	AnchoredSiteID string `url:"anchoredSiteId,omitempty"` //Fabric ID of the fabric site the layer 3 virtual network is anchored at.
}
type GetMulticastQueryParams struct {
	FabricID string  `url:"fabricId,omitempty"` //ID of the fabric site where multicast is configured.
	Offset   float64 `url:"offset,omitempty"`   //Starting record for pagination.
	Limit    float64 `url:"limit,omitempty"`    //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetMulticastVirtualNetworksQueryParams struct {
	FabricID           string  `url:"fabricId,omitempty"`           //ID of the fabric site where multicast is configured.
	VirtualNetworkName string  `url:"virtualNetworkName,omitempty"` //Name of the virtual network associated to the multicast configuration.
	Offset             float64 `url:"offset,omitempty"`             //Starting record for pagination.
	Limit              float64 `url:"limit,omitempty"`              //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetMulticastVirtualNetworkCountQueryParams struct {
	FabricID string `url:"fabricId,omitempty"` //ID of the fabric site the multicast configuration is associated with.
}
type GetPendingFabricEventsQueryParams struct {
	FabricID string  `url:"fabricId,omitempty"` //ID of the fabric.
	Offset   float64 `url:"offset,omitempty"`   //Starting record for pagination.
	Limit    float64 `url:"limit,omitempty"`    //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetPortAssignmentsQueryParams struct {
	FabricID        string  `url:"fabricId,omitempty"`        //ID of the fabric the device is assigned to.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Network device ID of the port assignment.
	InterfaceName   string  `url:"interfaceName,omitempty"`   //Interface name of the port assignment.
	DataVLANName    string  `url:"dataVlanName,omitempty"`    //Data VLAN name of the port assignment.
	VoiceVLANName   string  `url:"voiceVlanName,omitempty"`   //Voice VLAN name of the port assignment.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeletePortAssignmentsQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric the device is assigned to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the port assignment.
	InterfaceName   string `url:"interfaceName,omitempty"`   //Interface name of the port assignment.
	DataVLANName    string `url:"dataVlanName,omitempty"`    //Data VLAN name of the port assignment.
	VoiceVLANName   string `url:"voiceVlanName,omitempty"`   //Voice VLAN name of the port assignment.
}
type GetPortAssignmentCountQueryParams struct {
	FabricID        string `url:"fabricId,omitempty"`        //ID of the fabric the device is assigned to.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Network device ID of the port assignment.
	InterfaceName   string `url:"interfaceName,omitempty"`   //Interface name of the port assignment.
	DataVLANName    string `url:"dataVlanName,omitempty"`    //Data VLAN name of the port assignment.
	VoiceVLANName   string `url:"voiceVlanName,omitempty"`   //Voice VLAN name of the port assignment.
}
type GetPortChannelsQueryParams struct {
	FabricID            string  `url:"fabricId,omitempty"`            //ID of the fabric the device is assigned to.
	NetworkDeviceID     string  `url:"networkDeviceId,omitempty"`     //ID of the network device.
	PortChannelName     string  `url:"portChannelName,omitempty"`     //Name of the port channel.
	ConnectedDeviceType string  `url:"connectedDeviceType,omitempty"` //Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
	Offset              float64 `url:"offset,omitempty"`              //Starting record for pagination.
	Limit               float64 `url:"limit,omitempty"`               //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type DeletePortChannelsQueryParams struct {
	FabricID            string `url:"fabricId,omitempty"`            //ID of the fabric the device is assigned to.
	NetworkDeviceID     string `url:"networkDeviceId,omitempty"`     //ID of the network device.
	PortChannelName     string `url:"portChannelName,omitempty"`     //Name of the port channel.
	PortChannelIDs      string `url:"portChannelIds,omitempty"`      //IDs of the port channels to be selectively deleted(Maximum number of IDs this parameter could consume is 10).
	ConnectedDeviceType string `url:"connectedDeviceType,omitempty"` //Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
}
type GetPortChannelCountQueryParams struct {
	FabricID            string `url:"fabricId,omitempty"`            //ID of the fabric the device is assigned to.
	NetworkDeviceID     string `url:"networkDeviceId,omitempty"`     //ID of the network device.
	PortChannelName     string `url:"portChannelName,omitempty"`     //Name of the port channel.
	ConnectedDeviceType string `url:"connectedDeviceType,omitempty"` //Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
}
type DeleteProvisionedDevicesQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //ID of the network device.
	SiteID          string `url:"siteId,omitempty"`          //ID of the site hierarchy.
	CleanUpConfig   bool   `url:"cleanUpConfig,omitempty"`   //Enable/disable configuration cleanup for the device(s). Defaults to true.
}
type GetProvisionedDevicesQueryParams struct {
	ID              string  `url:"id,omitempty"`              //ID of the provisioned device.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //ID of the network device.
	SiteID          string  `url:"siteId,omitempty"`          //ID of the site hierarchy.
	Offset          float64 `url:"offset,omitempty"`          //Starting record for pagination.
	Limit           float64 `url:"limit,omitempty"`           //Maximum number of devices to return. The maximum number of objects supported in a single request is 500.
}
type GetProvisionedDevicesCountQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //ID of the site hierarchy.
}
type DeleteProvisionedDeviceByIDQueryParams struct {
	CleanUpConfig bool `url:"cleanUpConfig,omitempty"` //Enable/disable configuration cleanup for the particular device. Defaults to true.
}
type GetTransitNetworksQueryParams struct {
	ID     string  `url:"id,omitempty"`     //ID of the transit network.
	Name   string  `url:"name,omitempty"`   //Name of the transit network.
	Type   string  `url:"type,omitempty"`   //Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].
	Offset float64 `url:"offset,omitempty"` //Starting record for pagination.
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return. The maximum number of objects supported in a single request is 500.
}
type GetTransitNetworksCountQueryParams struct {
	Type string `url:"type,omitempty"` //Type of the transit network. Allowed values are [IP_BASED_TRANSIT, SDA_LISP_PUB_SUB_TRANSIT, SDA_LISP_BGP_TRANSIT].
}
type DeleteVirtualNetworkWithScalableGroupsQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}
type GetVirtualNetworkWithScalableGroupsQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}

type ResponseSdaReadListOfFabricSitesWithTheirHealthSummary struct {
	Response *[]ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryResponse `json:"response,omitempty"` //
}
type ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	GoodHealthPercentage *int `json:"goodHealthPercentage,omitempty"` // Good Health Percentage

	GoodHealthDeviceCount *int `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count

	TotalDeviceCount *int `json:"totalDeviceCount,omitempty"` // Total Device Count

	PoorHealthDeviceCount *int `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count

	FairHealthDeviceCount *int `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count

	AssociatedL2VnCount *int `json:"associatedL2VnCount,omitempty"` // Associated L2 Vn Count

	AssociatedL3VnCount *int `json:"associatedL3VnCount,omitempty"` // Associated L3 Vn Count

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol

	ConnectivityGoodHealthPercentage *int `json:"connectivityGoodHealthPercentage,omitempty"` // Connectivity Good Health Percentage

	ConnectivityTotalHealthDeviceCount *int `json:"connectivityTotalHealthDeviceCount,omitempty"` // Connectivity Total Health Device Count

	ConnectivityGoodHealthDeviceCount *int `json:"connectivityGoodHealthDeviceCount,omitempty"` // Connectivity Good Health Device Count

	ConnectivityPoorHealthDeviceCount *int `json:"connectivityPoorHealthDeviceCount,omitempty"` // Connectivity Poor Health Device Count

	ConnectivityFairHealthDeviceCount *int `json:"connectivityFairHealthDeviceCount,omitempty"` // Connectivity Fair Health Device Count

	InfraGoodHealthPercentage *int `json:"infraGoodHealthPercentage,omitempty"` // Infra Good Health Percentage

	InfraTotalHealthDeviceCount *int `json:"infraTotalHealthDeviceCount,omitempty"` // Infra Total Health Device Count

	InfraGoodHealthDeviceCount *int `json:"infraGoodHealthDeviceCount,omitempty"` // Infra Good Health Device Count

	InfraFairHealthDeviceCount *float64 `json:"infraFairHealthDeviceCount,omitempty"` // Infra Fair Health Device Count

	InfraPoorHealthDeviceCount *float64 `json:"infraPoorHealthDeviceCount,omitempty"` // Infra Poor Health Device Count

	ControlPlaneGoodHealthPercentage *int `json:"controlPlaneGoodHealthPercentage,omitempty"` // Control Plane Good Health Percentage

	ControlPlaneTotalHealthDeviceCount *int `json:"controlPlaneTotalHealthDeviceCount,omitempty"` // Control Plane Total Health Device Count

	ControlPlaneGoodHealthDeviceCount *int `json:"controlPlaneGoodHealthDeviceCount,omitempty"` // Control Plane Good Health Device Count

	ControlPlanePoorHealthDeviceCount *float64 `json:"controlPlanePoorHealthDeviceCount,omitempty"` // Control Plane Poor Health Device Count

	ControlPlaneFairHealthDeviceCount *int `json:"controlPlaneFairHealthDeviceCount,omitempty"` // Control Plane Fair Health Device Count

	PubsubInfraVnGoodHealthPercentage *int `json:"pubsubInfraVnGoodHealthPercentage,omitempty"` // Pubsub Infra Vn Good Health Percentage

	PubsubInfraVnTotalHealthDeviceCount *int `json:"pubsubInfraVnTotalHealthDeviceCount,omitempty"` // Pubsub Infra Vn Total Health Device Count

	PubsubInfraVnGoodHealthDeviceCount *int `json:"pubsubInfraVnGoodHealthDeviceCount,omitempty"` // Pubsub Infra Vn Good Health Device Count

	PubsubInfraVnPoorHealthDeviceCount *int `json:"pubsubInfraVnPoorHealthDeviceCount,omitempty"` // Pubsub Infra Vn Poor Health Device Count

	PubsubInfraVnFairHealthDeviceCount *int `json:"pubsubInfraVnFairHealthDeviceCount,omitempty"` // Pubsub Infra Vn Fair Health Device Count

	BgpEvpnGoodHealthPercentage *ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryResponseBgpEvpnGoodHealthPercentage `json:"bgpEvpnGoodHealthPercentage,omitempty"` // Bgp Evpn Good Health Percentage

	BgpEvpnTotalHealthDeviceCount *float64 `json:"bgpEvpnTotalHealthDeviceCount,omitempty"` // Bgp Evpn Total Health Device Count

	BgpEvpnGoodHealthDeviceCount *float64 `json:"bgpEvpnGoodHealthDeviceCount,omitempty"` // Bgp Evpn Good Health Device Count

	BgpEvpnpoorHealthDeviceCount *float64 `json:"bgpEvpnPoorHealthDeviceCount,omitempty"` // Bgp Evpn Poor Health Device Count

	BgpEvpnFairHealthDeviceCount *float64 `json:"bgpEvpnFairHealthDeviceCount,omitempty"` // Bgp Evpn Fair Health Device Count

	CtsEnvDataDownloadGoodHealthPercentage *int `json:"ctsEnvDataDownloadGoodHealthPercentage,omitempty"` // Cts Env Data Download Good Health Percentage

	CtsEnvDataDownloadTotalHealthDeviceCount *int `json:"ctsEnvDataDownloadTotalHealthDeviceCount,omitempty"` // Cts Env Data Download Total Health Device Count

	CtsEnvDataDownloadGoodHealthDeviceCount *int `json:"ctsEnvDataDownloadGoodHealthDeviceCount,omitempty"` // Cts Env Data Download Good Health Device Count

	CtsEnvDataDownloadPoorHealthDeviceCount *int `json:"ctsEnvDataDownloadPoorHealthDeviceCount,omitempty"` // Cts Env Data Download Poor Health Device Count

	CtsEnvDataDownloadFairHealthDeviceCount *float64 `json:"ctsEnvDataDownloadFairHealthDeviceCount,omitempty"` // Cts Env Data Download Fair Health Device Count

	AAAStatusGoodHealthPercentage *int `json:"aaaStatusGoodHealthPercentage,omitempty"` // Aaa Status Good Health Percentage

	AAAStatusTotalHealthDeviceCount *int `json:"aaaStatusTotalHealthDeviceCount,omitempty"` // Aaa Status Total Health Device Count

	AAAStatusGoodHealthDeviceCount *int `json:"aaaStatusGoodHealthDeviceCount,omitempty"` // Aaa Status Good Health Device Count

	AAAStatusPoorHealthDeviceCount *int `json:"aaaStatusPoorHealthDeviceCount,omitempty"` // Aaa Status Poor Health Device Count

	AAAStatusFairHealthDeviceCount *int `json:"aaaStatusFairHealthDeviceCount,omitempty"` // Aaa Status Fair Health Device Count

	PortChannelGoodHealthPercentage *int `json:"portChannelGoodHealthPercentage,omitempty"` // Port Channel Good Health Percentage

	PortChannelTotalHealthDeviceCount *int `json:"portChannelTotalHealthDeviceCount,omitempty"` // Port Channel Total Health Device Count

	PortChannelGoodHealthDeviceCount *int `json:"portChannelGoodHealthDeviceCount,omitempty"` // Port Channel Good Health Device Count

	PortChannelPoorHealthDeviceCount *int `json:"portChannelPoorHealthDeviceCount,omitempty"` // Port Channel Poor Health Device Count

	PortChannelFairHealthDeviceCount *int `json:"portChannelFairHealthDeviceCount,omitempty"` // Port Channel Fair Health Device Count

	PeerScoreGoodHealthPercentage *ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryResponsePeerScoreGoodHealthPercentage `json:"peerScoreGoodHealthPercentage,omitempty"` // Peer Score Good Health Percentage

	PeerScoreTotalHealthDeviceCount *float64 `json:"peerScoreTotalHealthDeviceCount,omitempty"` // Peer Score Total Health Device Count

	PeerScoreGoodHealthDeviceCount *float64 `json:"peerScoreGoodHealthDeviceCount,omitempty"` // Peer Score Good Health Device Count

	PeerScorePoorHealthDeviceCount *float64 `json:"peerScorePoorHealthDeviceCount,omitempty"` // Peer Score Poor Health Device Count

	PeerScoreFairHealthDeviceCount *float64 `json:"peerScoreFairHealthDeviceCount,omitempty"` // Peer Score Fair Health Device Count

	LispSessionGoodHealthPercentage *int `json:"lispSessionGoodHealthPercentage,omitempty"` // Lisp Session Good Health Percentage

	LispSessionTotalHealthDeviceCount *int `json:"lispSessionTotalHealthDeviceCount,omitempty"` // Lisp Session Total Health Device Count

	LispSessionGoodHealthDeviceCount *int `json:"lispSessionGoodHealthDeviceCount,omitempty"` // Lisp Session Good Health Device Count

	LispSessionPoorHealthDeviceCount *int `json:"lispSessionPoorHealthDeviceCount,omitempty"` // Lisp Session Poor Health Device Count

	LispSessionFairHealthDeviceCount *float64 `json:"lispSessionFairHealthDeviceCount,omitempty"` // Lisp Session Fair Health Device Count

	BorderToControlPlaneGoodHealthPercentage *int `json:"borderToControlPlaneGoodHealthPercentage,omitempty"` // Border To Control Plane Good Health Percentage

	BorderToControlPlaneTotalHealthDeviceCount *int `json:"borderToControlPlaneTotalHealthDeviceCount,omitempty"` // Border To Control Plane Total Health Device Count

	BorderToControlPlaneGoodHealthDeviceCount *int `json:"borderToControlPlaneGoodHealthDeviceCount,omitempty"` // Border To Control Plane Good Health Device Count

	BorderToControlPlanePoorHealthDeviceCount *float64 `json:"borderToControlPlanePoorHealthDeviceCount,omitempty"` // Border To Control Plane Poor Health Device Count

	BorderToControlPlaneFairHealthDeviceCount *int `json:"borderToControlPlaneFairHealthDeviceCount,omitempty"` // Border To Control Plane Fair Health Device Count

	BgpBgpSiteGoodHealthPercentage *int `json:"bgpBgpSiteGoodHealthPercentage,omitempty"` // Bgp Bgp Site Good Health Percentage

	BgpBgpSiteTotalHealthDeviceCount *int `json:"bgpBgpSiteTotalHealthDeviceCount,omitempty"` // Bgp Bgp Site Total Health Device Count

	BgpBgpSiteGoodHealthDeviceCount *int `json:"bgpBgpSiteGoodHealthDeviceCount,omitempty"` // Bgp Bgp Site Good Health Device Count

	BgpBgpSitePoorHealthDeviceCount *float64 `json:"bgpBgpSitePoorHealthDeviceCount,omitempty"` // Bgp Bgp Site Poor Health Device Count

	BgpBgpSiteFairHealthDeviceCount *int `json:"bgpBgpSiteFairHealthDeviceCount,omitempty"` // Bgp Bgp Site Fair Health Device Count

	BgpPubsubSiteGoodHealthPercentage *int `json:"bgpPubsubSiteGoodHealthPercentage,omitempty"` // Bgp Pubsub Site Good Health Percentage

	BgpPubsubSiteTotalHealthDeviceCount *int `json:"bgpPubsubSiteTotalHealthDeviceCount,omitempty"` // Bgp Pubsub Site Total Health Device Count

	BgpPubsubSiteGoodHealthDeviceCount *int `json:"bgpPubsubSiteGoodHealthDeviceCount,omitempty"` // Bgp Pubsub Site Good Health Device Count

	BgpPubsubSitePoorHealthDeviceCount *float64 `json:"bgpPubsubSitePoorHealthDeviceCount,omitempty"` // Bgp Pubsub Site Poor Health Device Count

	BgpPubsubSiteFairHealthDeviceCount *int `json:"bgpPubsubSiteFairHealthDeviceCount,omitempty"` // Bgp Pubsub Site Fair Health Device Count

	BgpPeerInfraVnScoreGoodHealthPercentage *int `json:"bgpPeerInfraVnScoreGoodHealthPercentage,omitempty"` // Bgp Peer Infra Vn Score Good Health Percentage

	BgpPeerInfraVnTotalHealthDeviceCount *int `json:"bgpPeerInfraVnTotalHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Total Health Device Count

	BgpPeerInfraVnGoodHealthDeviceCount *int `json:"bgpPeerInfraVnGoodHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Good Health Device Count

	BgpPeerInfraVnPoorHealthDeviceCount *int `json:"bgpPeerInfraVnPoorHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Poor Health Device Count

	BgpPeerInfraVnFairHealthDeviceCount *float64 `json:"bgpPeerInfraVnFairHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Fair Health Device Count
}
type ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryResponseBgpEvpnGoodHealthPercentage interface{}
type ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryResponsePeerScoreGoodHealthPercentage interface{}
type ResponseSdaReadFabricSiteCount struct {
	Response *ResponseSdaReadFabricSiteCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSdaReadFabricSiteCountResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSdaReadFabricSitesWithHealthSummaryFromID struct {
	Response *ResponseSdaReadFabricSitesWithHealthSummaryFromIDResponse `json:"response,omitempty"` //
}
type ResponseSdaReadFabricSitesWithHealthSummaryFromIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	GoodHealthPercentage *int `json:"goodHealthPercentage,omitempty"` // Good Health Percentage

	GoodHealthDeviceCount *int `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count

	TotalDeviceCount *int `json:"totalDeviceCount,omitempty"` // Total Device Count

	PoorHealthDeviceCount *int `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count

	FairHealthDeviceCount *int `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count

	AssociatedL2VnCount *int `json:"associatedL2VnCount,omitempty"` // Associated L2 Vn Count

	AssociatedL3VnCount *int `json:"associatedL3VnCount,omitempty"` // Associated L3 Vn Count

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol

	ConnectivityGoodHealthPercentage *int `json:"connectivityGoodHealthPercentage,omitempty"` // Connectivity Good Health Percentage

	ConnectivityTotalHealthDeviceCount *int `json:"connectivityTotalHealthDeviceCount,omitempty"` // Connectivity Total Health Device Count

	ConnectivityGoodHealthDeviceCount *int `json:"connectivityGoodHealthDeviceCount,omitempty"` // Connectivity Good Health Device Count

	ConnectivityPoorHealthDeviceCount *int `json:"connectivityPoorHealthDeviceCount,omitempty"` // Connectivity Poor Health Device Count

	ConnectivityFairHealthDeviceCount *int `json:"connectivityFairHealthDeviceCount,omitempty"` // Connectivity Fair Health Device Count

	InfraGoodHealthPercentage *int `json:"infraGoodHealthPercentage,omitempty"` // Infra Good Health Percentage

	InfraTotalHealthDeviceCount *int `json:"infraTotalHealthDeviceCount,omitempty"` // Infra Total Health Device Count

	InfraGoodHealthDeviceCount *int `json:"infraGoodHealthDeviceCount,omitempty"` // Infra Good Health Device Count

	InfraFairHealthDeviceCount *float64 `json:"infraFairHealthDeviceCount,omitempty"` // Infra Fair Health Device Count

	InfraPoorHealthDeviceCount *float64 `json:"infraPoorHealthDeviceCount,omitempty"` // Infra Poor Health Device Count

	ControlPlaneGoodHealthPercentage *int `json:"controlPlaneGoodHealthPercentage,omitempty"` // Control Plane Good Health Percentage

	ControlPlaneTotalHealthDeviceCount *int `json:"controlPlaneTotalHealthDeviceCount,omitempty"` // Control Plane Total Health Device Count

	ControlPlaneGoodHealthDeviceCount *int `json:"controlPlaneGoodHealthDeviceCount,omitempty"` // Control Plane Good Health Device Count

	ControlPlanePoorHealthDeviceCount *float64 `json:"controlPlanePoorHealthDeviceCount,omitempty"` // Control Plane Poor Health Device Count

	ControlPlaneFairHealthDeviceCount *int `json:"controlPlaneFairHealthDeviceCount,omitempty"` // Control Plane Fair Health Device Count

	PubsubInfraVnGoodHealthPercentage *int `json:"pubsubInfraVnGoodHealthPercentage,omitempty"` // Pubsub Infra Vn Good Health Percentage

	PubsubInfraVnTotalHealthDeviceCount *int `json:"pubsubInfraVnTotalHealthDeviceCount,omitempty"` // Pubsub Infra Vn Total Health Device Count

	PubsubInfraVnGoodHealthDeviceCount *int `json:"pubsubInfraVnGoodHealthDeviceCount,omitempty"` // Pubsub Infra Vn Good Health Device Count

	PubsubInfraVnPoorHealthDeviceCount *int `json:"pubsubInfraVnPoorHealthDeviceCount,omitempty"` // Pubsub Infra Vn Poor Health Device Count

	PubsubInfraVnFairHealthDeviceCount *int `json:"pubsubInfraVnFairHealthDeviceCount,omitempty"` // Pubsub Infra Vn Fair Health Device Count

	BgpEvpnGoodHealthPercentage *ResponseSdaReadFabricSitesWithHealthSummaryFromIDResponseBgpEvpnGoodHealthPercentage `json:"bgpEvpnGoodHealthPercentage,omitempty"` // Bgp Evpn Good Health Percentage

	BgpEvpnTotalHealthDeviceCount *float64 `json:"bgpEvpnTotalHealthDeviceCount,omitempty"` // Bgp Evpn Total Health Device Count

	BgpEvpnGoodHealthDeviceCount *float64 `json:"bgpEvpnGoodHealthDeviceCount,omitempty"` // Bgp Evpn Good Health Device Count

	BgpEvpnpoorHealthDeviceCount *float64 `json:"bgpEvpnPoorHealthDeviceCount,omitempty"` // Bgp Evpn Poor Health Device Count

	BgpEvpnFairHealthDeviceCount *float64 `json:"bgpEvpnFairHealthDeviceCount,omitempty"` // Bgp Evpn Fair Health Device Count

	CtsEnvDataDownloadGoodHealthPercentage *int `json:"ctsEnvDataDownloadGoodHealthPercentage,omitempty"` // Cts Env Data Download Good Health Percentage

	CtsEnvDataDownloadTotalHealthDeviceCount *int `json:"ctsEnvDataDownloadTotalHealthDeviceCount,omitempty"` // Cts Env Data Download Total Health Device Count

	CtsEnvDataDownloadGoodHealthDeviceCount *int `json:"ctsEnvDataDownloadGoodHealthDeviceCount,omitempty"` // Cts Env Data Download Good Health Device Count

	CtsEnvDataDownloadPoorHealthDeviceCount *int `json:"ctsEnvDataDownloadPoorHealthDeviceCount,omitempty"` // Cts Env Data Download Poor Health Device Count

	CtsEnvDataDownloadFairHealthDeviceCount *float64 `json:"ctsEnvDataDownloadFairHealthDeviceCount,omitempty"` // Cts Env Data Download Fair Health Device Count

	AAAStatusGoodHealthPercentage *int `json:"aaaStatusGoodHealthPercentage,omitempty"` // Aaa Status Good Health Percentage

	AAAStatusTotalHealthDeviceCount *int `json:"aaaStatusTotalHealthDeviceCount,omitempty"` // Aaa Status Total Health Device Count

	AAAStatusGoodHealthDeviceCount *int `json:"aaaStatusGoodHealthDeviceCount,omitempty"` // Aaa Status Good Health Device Count

	AAAStatusPoorHealthDeviceCount *int `json:"aaaStatusPoorHealthDeviceCount,omitempty"` // Aaa Status Poor Health Device Count

	AAAStatusFairHealthDeviceCount *int `json:"aaaStatusFairHealthDeviceCount,omitempty"` // Aaa Status Fair Health Device Count

	PortChannelGoodHealthPercentage *int `json:"portChannelGoodHealthPercentage,omitempty"` // Port Channel Good Health Percentage

	PortChannelTotalHealthDeviceCount *int `json:"portChannelTotalHealthDeviceCount,omitempty"` // Port Channel Total Health Device Count

	PortChannelGoodHealthDeviceCount *int `json:"portChannelGoodHealthDeviceCount,omitempty"` // Port Channel Good Health Device Count

	PortChannelPoorHealthDeviceCount *int `json:"portChannelPoorHealthDeviceCount,omitempty"` // Port Channel Poor Health Device Count

	PortChannelFairHealthDeviceCount *int `json:"portChannelFairHealthDeviceCount,omitempty"` // Port Channel Fair Health Device Count

	PeerScoreGoodHealthPercentage *ResponseSdaReadFabricSitesWithHealthSummaryFromIDResponsePeerScoreGoodHealthPercentage `json:"peerScoreGoodHealthPercentage,omitempty"` // Peer Score Good Health Percentage

	PeerScoreTotalHealthDeviceCount *float64 `json:"peerScoreTotalHealthDeviceCount,omitempty"` // Peer Score Total Health Device Count

	PeerScoreGoodHealthDeviceCount *float64 `json:"peerScoreGoodHealthDeviceCount,omitempty"` // Peer Score Good Health Device Count

	PeerScorePoorHealthDeviceCount *float64 `json:"peerScorePoorHealthDeviceCount,omitempty"` // Peer Score Poor Health Device Count

	PeerScoreFairHealthDeviceCount *float64 `json:"peerScoreFairHealthDeviceCount,omitempty"` // Peer Score Fair Health Device Count

	LispSessionGoodHealthPercentage *int `json:"lispSessionGoodHealthPercentage,omitempty"` // Lisp Session Good Health Percentage

	LispSessionTotalHealthDeviceCount *int `json:"lispSessionTotalHealthDeviceCount,omitempty"` // Lisp Session Total Health Device Count

	LispSessionGoodHealthDeviceCount *int `json:"lispSessionGoodHealthDeviceCount,omitempty"` // Lisp Session Good Health Device Count

	LispSessionPoorHealthDeviceCount *int `json:"lispSessionPoorHealthDeviceCount,omitempty"` // Lisp Session Poor Health Device Count

	LispSessionFairHealthDeviceCount *float64 `json:"lispSessionFairHealthDeviceCount,omitempty"` // Lisp Session Fair Health Device Count

	BorderToControlPlaneGoodHealthPercentage *int `json:"borderToControlPlaneGoodHealthPercentage,omitempty"` // Border To Control Plane Good Health Percentage

	BorderToControlPlaneTotalHealthDeviceCount *int `json:"borderToControlPlaneTotalHealthDeviceCount,omitempty"` // Border To Control Plane Total Health Device Count

	BorderToControlPlaneGoodHealthDeviceCount *int `json:"borderToControlPlaneGoodHealthDeviceCount,omitempty"` // Border To Control Plane Good Health Device Count

	BorderToControlPlanePoorHealthDeviceCount *float64 `json:"borderToControlPlanePoorHealthDeviceCount,omitempty"` // Border To Control Plane Poor Health Device Count

	BorderToControlPlaneFairHealthDeviceCount *int `json:"borderToControlPlaneFairHealthDeviceCount,omitempty"` // Border To Control Plane Fair Health Device Count

	BgpBgpSiteGoodHealthPercentage *int `json:"bgpBgpSiteGoodHealthPercentage,omitempty"` // Bgp Bgp Site Good Health Percentage

	BgpBgpSiteTotalHealthDeviceCount *int `json:"bgpBgpSiteTotalHealthDeviceCount,omitempty"` // Bgp Bgp Site Total Health Device Count

	BgpBgpSiteGoodHealthDeviceCount *int `json:"bgpBgpSiteGoodHealthDeviceCount,omitempty"` // Bgp Bgp Site Good Health Device Count

	BgpBgpSitePoorHealthDeviceCount *float64 `json:"bgpBgpSitePoorHealthDeviceCount,omitempty"` // Bgp Bgp Site Poor Health Device Count

	BgpBgpSiteFairHealthDeviceCount *int `json:"bgpBgpSiteFairHealthDeviceCount,omitempty"` // Bgp Bgp Site Fair Health Device Count

	BgpPubsubSiteGoodHealthPercentage *int `json:"bgpPubsubSiteGoodHealthPercentage,omitempty"` // Bgp Pubsub Site Good Health Percentage

	BgpPubsubSiteTotalHealthDeviceCount *int `json:"bgpPubsubSiteTotalHealthDeviceCount,omitempty"` // Bgp Pubsub Site Total Health Device Count

	BgpPubsubSiteGoodHealthDeviceCount *int `json:"bgpPubsubSiteGoodHealthDeviceCount,omitempty"` // Bgp Pubsub Site Good Health Device Count

	BgpPubsubSitePoorHealthDeviceCount *float64 `json:"bgpPubsubSitePoorHealthDeviceCount,omitempty"` // Bgp Pubsub Site Poor Health Device Count

	BgpPubsubSiteFairHealthDeviceCount *int `json:"bgpPubsubSiteFairHealthDeviceCount,omitempty"` // Bgp Pubsub Site Fair Health Device Count

	BgpPeerInfraVnScoreGoodHealthPercentage *int `json:"bgpPeerInfraVnScoreGoodHealthPercentage,omitempty"` // Bgp Peer Infra Vn Score Good Health Percentage

	BgpPeerInfraVnTotalHealthDeviceCount *int `json:"bgpPeerInfraVnTotalHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Total Health Device Count

	BgpPeerInfraVnGoodHealthDeviceCount *int `json:"bgpPeerInfraVnGoodHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Good Health Device Count

	BgpPeerInfraVnPoorHealthDeviceCount *int `json:"bgpPeerInfraVnPoorHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Poor Health Device Count

	BgpPeerInfraVnFairHealthDeviceCount *float64 `json:"bgpPeerInfraVnFairHealthDeviceCount,omitempty"` // Bgp Peer Infra Vn Fair Health Device Count
}
type ResponseSdaReadFabricSitesWithHealthSummaryFromIDResponseBgpEvpnGoodHealthPercentage interface{}
type ResponseSdaReadFabricSitesWithHealthSummaryFromIDResponsePeerScoreGoodHealthPercentage interface{}
type ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange struct {
	Response *[]ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeResponse `json:"response,omitempty"` //

	Page *ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangePage `json:"page,omitempty"` //

	Version *int `json:"version,omitempty"` // Version
}
type ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value *int `json:"value,omitempty"` // Value
}
type ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *float64 `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSdaReadFabricEntitySummary struct {
	Response *ResponseSdaReadFabricEntitySummaryResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSdaReadFabricEntitySummaryResponse struct {
	ProtocolSummaries *[]ResponseSdaReadFabricEntitySummaryResponseProtocolSummaries `json:"protocolSummaries,omitempty"` //
}
type ResponseSdaReadFabricEntitySummaryResponseProtocolSummaries struct {
	FabricSiteGoodHealthCount *int `json:"fabricSiteGoodHealthCount,omitempty"` // Fabric Site Good Health Count

	FabricSiteCount *int `json:"fabricSiteCount,omitempty"` // Fabric Site Count

	FabricSiteGoodHealthPercentage *int `json:"fabricSiteGoodHealthPercentage,omitempty"` // Fabric Site Good Health Percentage

	FabricSiteNoHealthCount *int `json:"fabricSiteNoHealthCount,omitempty"` // Fabric Site No Health Count

	FabricSitePoorHealthCount *int `json:"fabricSitePoorHealthCount,omitempty"` // Fabric Site Poor Health Count

	FabricSiteFairHealthCount *float64 `json:"fabricSiteFairHealthCount,omitempty"` // Fabric Site Fair Health Count

	L3VnGoodHealthCount *int `json:"l3VnGoodHealthCount,omitempty"` // L3 Vn Good Health Count

	L3VnCount *int `json:"l3VnCount,omitempty"` // L3 Vn Count

	L3VnGoodHealthPercentage *int `json:"l3VnGoodHealthPercentage,omitempty"` // L3 Vn Good Health Percentage

	L3VnNoHealthCount *int `json:"l3VnNoHealthCount,omitempty"` // L3 Vn No Health Count

	L3VnFairHealthCount *float64 `json:"l3VnFairHealthCount,omitempty"` // L3 Vn Fair Health Count

	L3VnPoorHealthCount *int `json:"l3VnPoorHealthCount,omitempty"` // L3 Vn Poor Health Count

	L2VnGoodHealthCount *int `json:"l2VnGoodHealthCount,omitempty"` // L2 Vn Good Health Count

	L2VnCount *int `json:"l2VnCount,omitempty"` // L2 Vn Count

	L2VnGoodHealthPercentage *int `json:"l2VnGoodHealthPercentage,omitempty"` // L2 Vn Good Health Percentage

	L2VnNoHealthCount *int `json:"l2VnNoHealthCount,omitempty"` // L2 Vn No Health Count

	L2VnPoorHealthCount *int `json:"l2VnPoorHealthCount,omitempty"` // L2 Vn Poor Health Count

	L2VnFairHealthCount *float64 `json:"l2VnFairHealthCount,omitempty"` // L2 Vn Fair Health Count

	TransitNetworkGoodHealthCount *int `json:"transitNetworkGoodHealthCount,omitempty"` // Transit Network Good Health Count

	TransitNetworkCount *int `json:"transitNetworkCount,omitempty"` // Transit Network Count

	TransitNetworkGoodHealthPercentage *int `json:"transitNetworkGoodHealthPercentage,omitempty"` // Transit Network Good Health Percentage

	TransitNetworkNoHealthCount *int `json:"transitNetworkNoHealthCount,omitempty"` // Transit Network No Health Count

	TransitNetworkPoorHealthCount *int `json:"transitNetworkPoorHealthCount,omitempty"` // Transit Network Poor Health Count

	TransitNetworkFairHealthCount *float64 `json:"transitNetworkFairHealthCount,omitempty"` // Transit Network Fair Health Count

	IPTransitNetworkCount *int `json:"ipTransitNetworkCount,omitempty"` // Ip Transit Network Count

	FabricDeviceCount *int `json:"fabricDeviceCount,omitempty"` // Fabric Device Count

	P1IssueCount *int `json:"p1IssueCount,omitempty"` // P1 Issue Count

	P2IssueCount *int `json:"p2IssueCount,omitempty"` // P2 Issue Count

	P3IssueCount *int `json:"p3IssueCount,omitempty"` // P3 Issue Count

	NetworkSegmentProtocol string `json:"networkSegmentProtocol,omitempty"` // Network Segment Protocol
}
type ResponseSdaReadListOfTransitNetworksWithTheirHealthSummary struct {
	Response *[]ResponseSdaReadListOfTransitNetworksWithTheirHealthSummaryResponse `json:"response,omitempty"` //
}
type ResponseSdaReadListOfTransitNetworksWithTheirHealthSummaryResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	ControlPlaneCount *int `json:"controlPlaneCount,omitempty"` // Control Plane Count

	TransitType string `json:"transitType,omitempty"` // Transit Type

	FabricSitesCount *int `json:"fabricSitesCount,omitempty"` // Fabric Sites Count

	GoodHealthPercentage *int `json:"goodHealthPercentage,omitempty"` // Good Health Percentage

	GoodHealthDeviceCount *int `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count

	TotalDeviceCount *int `json:"totalDeviceCount,omitempty"` // Total Device Count

	PoorHealthDeviceCount *int `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count

	FairHealthDeviceCount *int `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count

	TransitControlPlaneHealthPercentage *int `json:"transitControlPlaneHealthPercentage,omitempty"` // Transit Control Plane Health Percentage

	TransitControlPlaneTotalDeviceCount *int `json:"transitControlPlaneTotalDeviceCount,omitempty"` // Transit Control Plane Total Device Count

	TransitControlPlaneGoodHealthDeviceCount *int `json:"transitControlPlaneGoodHealthDeviceCount,omitempty"` // Transit Control Plane Good Health Device Count

	TransitControlPlanePoorHealthDeviceCount *int `json:"transitControlPlanePoorHealthDeviceCount,omitempty"` // Transit Control Plane Poor Health Device Count

	TransitControlPlaneFairHealthDeviceCount *int `json:"transitControlPlaneFairHealthDeviceCount,omitempty"` // Transit Control Plane Fair Health Device Count

	TransitServicesHealthPercentage *int `json:"transitServicesHealthPercentage,omitempty"` // Transit Services Health Percentage

	TransitServicesTotalDeviceCount *int `json:"transitServicesTotalDeviceCount,omitempty"` // Transit Services Total Device Count

	TransitServicesGoodHealthDeviceCount *int `json:"transitServicesGoodHealthDeviceCount,omitempty"` // Transit Services Good Health Device Count

	TransitServicesPoorHealthDeviceCount *float64 `json:"transitServicesPoorHealthDeviceCount,omitempty"` // Transit Services Poor Health Device Count

	TransitServicesFairHealthDeviceCount *int `json:"transitServicesFairHealthDeviceCount,omitempty"` // Transit Services Fair Health Device Count

	PubsubTransitHealthPercentage *int `json:"pubsubTransitHealthPercentage,omitempty"` // Pubsub Transit Health Percentage

	PubsubTransitTotalDeviceCount *int `json:"pubsubTransitTotalDeviceCount,omitempty"` // Pubsub Transit Total Device Count

	PubsubTransitGoodHealthDeviceCount *int `json:"pubsubTransitGoodHealthDeviceCount,omitempty"` // Pubsub Transit Good Health Device Count

	PubsubTransitPoorHealthDeviceCount *int `json:"pubsubTransitPoorHealthDeviceCount,omitempty"` // Pubsub Transit Poor Health Device Count

	PubsubTransitFairHealthDeviceCount *int `json:"pubsubTransitFairHealthDeviceCount,omitempty"` // Pubsub Transit Fair Health Device Count

	LispTransitHealthPercentage *int `json:"lispTransitHealthPercentage,omitempty"` // Lisp Transit Health Percentage

	LispTransitTotalDeviceCount *int `json:"lispTransitTotalDeviceCount,omitempty"` // Lisp Transit Total Device Count

	LispTransitGoodHealthDeviceCount *int `json:"lispTransitGoodHealthDeviceCount,omitempty"` // Lisp Transit Good Health Device Count

	LispTransitPoorHealthDeviceCount *float64 `json:"lispTransitPoorHealthDeviceCount,omitempty"` // Lisp Transit Poor Health Device Count

	LispTransitFairHealthDeviceCount *int `json:"lispTransitFairHealthDeviceCount,omitempty"` // Lisp Transit Fair Health Device Count

	InternetAvailTransitHealthPercentage *int `json:"internetAvailTransitHealthPercentage,omitempty"` // Internet Avail Transit Health Percentage

	InternetAvailTransitTotalDeviceCount *int `json:"internetAvailTransitTotalDeviceCount,omitempty"` // Internet Avail Transit Total Device Count

	InternetAvailTransitGoodHealthDeviceCount *int `json:"internetAvailTransitGoodHealthDeviceCount,omitempty"` // Internet Avail Transit Good Health Device Count

	InternetAvailTransitPoorHealthDeviceCount *int `json:"internetAvailTransitPoorHealthDeviceCount,omitempty"` // Internet Avail Transit Poor Health Device Count

	InternetAvailTransitFairHealthDeviceCount *int `json:"internetAvailTransitFairHealthDeviceCount,omitempty"` // Internet Avail Transit Fair Health Device Count

	BgpTCPHealthPercentage *int `json:"bgpTcpHealthPercentage,omitempty"` // Bgp Tcp Health Percentage

	BgpTCPTotalDeviceCount *int `json:"bgpTcpTotalDeviceCount,omitempty"` // Bgp Tcp Total Device Count

	BgpTCPGoodHealthDeviceCount *int `json:"bgpTcpGoodHealthDeviceCount,omitempty"` // Bgp Tcp Good Health Device Count

	BgpTCPPoorHealthDeviceCount *int `json:"bgpTcpPoorHealthDeviceCount,omitempty"` // Bgp Tcp Poor Health Device Count

	BgpTCPFairHealthDeviceCount *int `json:"bgpTcpFairHealthDeviceCount,omitempty"` // Bgp Tcp Fair Health Device Count
}
type ResponseSdaReadTransitNetworksCount struct {
	Response *[]ResponseSdaReadTransitNetworksCountResponse `json:"response,omitempty"` //
}
type ResponseSdaReadTransitNetworksCountResponse struct {
	ErrorCode *int `json:"errorCode,omitempty"` // Error Code

	Message string `json:"message,omitempty"` // Message

	Detail string `json:"detail,omitempty"` // Detail
}
type ResponseSdaReadTransitNetworkWithItsHealthSummaryFromID struct {
	Response *ResponseSdaReadTransitNetworkWithItsHealthSummaryFromIDResponse `json:"response,omitempty"` //
}
type ResponseSdaReadTransitNetworkWithItsHealthSummaryFromIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	ControlPlaneCount *int `json:"controlPlaneCount,omitempty"` // Control Plane Count

	TransitType string `json:"transitType,omitempty"` // Transit Type

	FabricSitesCount *int `json:"fabricSitesCount,omitempty"` // Fabric Sites Count

	GoodHealthPercentage *int `json:"goodHealthPercentage,omitempty"` // Good Health Percentage

	GoodHealthDeviceCount *int `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count

	TotalDeviceCount *int `json:"totalDeviceCount,omitempty"` // Total Device Count

	PoorHealthDeviceCount *int `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count

	FairHealthDeviceCount *int `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count

	TransitControlPlaneHealthPercentage *int `json:"transitControlPlaneHealthPercentage,omitempty"` // Transit Control Plane Health Percentage

	TransitControlPlaneTotalDeviceCount *int `json:"transitControlPlaneTotalDeviceCount,omitempty"` // Transit Control Plane Total Device Count

	TransitControlPlaneGoodHealthDeviceCount *int `json:"transitControlPlaneGoodHealthDeviceCount,omitempty"` // Transit Control Plane Good Health Device Count

	TransitControlPlanePoorHealthDeviceCount *int `json:"transitControlPlanePoorHealthDeviceCount,omitempty"` // Transit Control Plane Poor Health Device Count

	TransitControlPlaneFairHealthDeviceCount *int `json:"transitControlPlaneFairHealthDeviceCount,omitempty"` // Transit Control Plane Fair Health Device Count

	TransitServicesHealthPercentage *int `json:"transitServicesHealthPercentage,omitempty"` // Transit Services Health Percentage

	TransitServicesTotalDeviceCount *int `json:"transitServicesTotalDeviceCount,omitempty"` // Transit Services Total Device Count

	TransitServicesGoodHealthDeviceCount *int `json:"transitServicesGoodHealthDeviceCount,omitempty"` // Transit Services Good Health Device Count

	TransitServicesPoorHealthDeviceCount *float64 `json:"transitServicesPoorHealthDeviceCount,omitempty"` // Transit Services Poor Health Device Count

	TransitServicesFairHealthDeviceCount *int `json:"transitServicesFairHealthDeviceCount,omitempty"` // Transit Services Fair Health Device Count

	PubsubTransitHealthPercentage *int `json:"pubsubTransitHealthPercentage,omitempty"` // Pubsub Transit Health Percentage

	PubsubTransitTotalDeviceCount *int `json:"pubsubTransitTotalDeviceCount,omitempty"` // Pubsub Transit Total Device Count

	PubsubTransitGoodHealthDeviceCount *int `json:"pubsubTransitGoodHealthDeviceCount,omitempty"` // Pubsub Transit Good Health Device Count

	PubsubTransitPoorHealthDeviceCount *int `json:"pubsubTransitPoorHealthDeviceCount,omitempty"` // Pubsub Transit Poor Health Device Count

	PubsubTransitFairHealthDeviceCount *int `json:"pubsubTransitFairHealthDeviceCount,omitempty"` // Pubsub Transit Fair Health Device Count

	LispTransitHealthPercentage *int `json:"lispTransitHealthPercentage,omitempty"` // Lisp Transit Health Percentage

	LispTransitTotalDeviceCount *int `json:"lispTransitTotalDeviceCount,omitempty"` // Lisp Transit Total Device Count

	LispTransitGoodHealthDeviceCount *int `json:"lispTransitGoodHealthDeviceCount,omitempty"` // Lisp Transit Good Health Device Count

	LispTransitPoorHealthDeviceCount *float64 `json:"lispTransitPoorHealthDeviceCount,omitempty"` // Lisp Transit Poor Health Device Count

	LispTransitFairHealthDeviceCount *int `json:"lispTransitFairHealthDeviceCount,omitempty"` // Lisp Transit Fair Health Device Count

	InternetAvailTransitHealthPercentage *int `json:"internetAvailTransitHealthPercentage,omitempty"` // Internet Avail Transit Health Percentage

	InternetAvailTransitTotalDeviceCount *int `json:"internetAvailTransitTotalDeviceCount,omitempty"` // Internet Avail Transit Total Device Count

	InternetAvailTransitGoodHealthDeviceCount *int `json:"internetAvailTransitGoodHealthDeviceCount,omitempty"` // Internet Avail Transit Good Health Device Count

	InternetAvailTransitPoorHealthDeviceCount *int `json:"internetAvailTransitPoorHealthDeviceCount,omitempty"` // Internet Avail Transit Poor Health Device Count

	InternetAvailTransitFairHealthDeviceCount *int `json:"internetAvailTransitFairHealthDeviceCount,omitempty"` // Internet Avail Transit Fair Health Device Count

	BgpTCPHealthPercentage *int `json:"bgpTcpHealthPercentage,omitempty"` // Bgp Tcp Health Percentage

	BgpTCPTotalDeviceCount *int `json:"bgpTcpTotalDeviceCount,omitempty"` // Bgp Tcp Total Device Count

	BgpTCPGoodHealthDeviceCount *int `json:"bgpTcpGoodHealthDeviceCount,omitempty"` // Bgp Tcp Good Health Device Count

	BgpTCPPoorHealthDeviceCount *int `json:"bgpTcpPoorHealthDeviceCount,omitempty"` // Bgp Tcp Poor Health Device Count

	BgpTCPFairHealthDeviceCount *int `json:"bgpTcpFairHealthDeviceCount,omitempty"` // Bgp Tcp Fair Health Device Count
}
type ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange struct {
	Response *[]ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeResponse `json:"response,omitempty"` //

	Page *ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangePage `json:"page,omitempty"` //

	Version *int `json:"version,omitempty"` // Version
}
type ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value *int `json:"value,omitempty"` // Value
}
type ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *float64 `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummary struct {
	Response *[]ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponse `json:"response,omitempty"` //
}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol

	VLAN *ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseVLAN `json:"vlan,omitempty"` // Vlan

	Vnid *int `json:"vnid,omitempty"` // Vnid

	Layer *ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseLayer `json:"layer,omitempty"` // Layer

	TotalFabricSites *int `json:"totalFabricSites,omitempty"` // Total Fabric Sites

	AssociatedL3Vn *ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseAssociatedL3Vn `json:"associatedL3Vn,omitempty"` // Associated L3 Vn

	TotalEndpoints *float64 `json:"totalEndpoints,omitempty"` // Total Endpoints

	GoodHealthPercentage *int `json:"goodHealthPercentage,omitempty"` // Good Health Percentage

	TotalDeviceCount *int `json:"totalDeviceCount,omitempty"` // Total Device Count

	GoodHealthDeviceCount *int `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count

	FairHealthDeviceCount *int `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count

	PoorHealthDeviceCount *int `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count

	NoHealthDeviceCount *int `json:"noHealthDeviceCount,omitempty"` // No Health Device Count

	VnFabricControlPlaneGoodHealthPercentage *int `json:"vnFabricControlPlaneGoodHealthPercentage,omitempty"` // Vn Fabric Control Plane Good Health Percentage

	VnFabricControlPlaneTotalDeviceCount *int `json:"vnFabricControlPlaneTotalDeviceCount,omitempty"` // Vn Fabric Control Plane Total Device Count

	VnFabricControlPlaneGoodHealthDeviceCount *int `json:"vnFabricControlPlaneGoodHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Good Health Device Count

	VnFabricControlPlanePoorHealthDeviceCount *int `json:"vnFabricControlPlanePoorHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Poor Health Device Count

	VnFabricControlPlaneFairHealthDeviceCount *float64 `json:"vnFabricControlPlaneFairHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Fair Health Device Count

	VnFabricControlPlaneNoHealthDeviceCount *float64 `json:"vnFabricControlPlaneNoHealthDeviceCount,omitempty"` // Vn Fabric Control Plane No Health Device Count

	VnServicesHealthPercentage *int `json:"vnServicesHealthPercentage,omitempty"` // Vn Services Health Percentage

	VnServicesTotalDeviceCount *int `json:"vnServicesTotalDeviceCount,omitempty"` // Vn Services Total Device Count

	VnServicesGoodHealthDeviceCount *int `json:"vnServicesGoodHealthDeviceCount,omitempty"` // Vn Services Good Health Device Count

	VnServicesPoorHealthDeviceCount *int `json:"vnServicesPoorHealthDeviceCount,omitempty"` // Vn Services Poor Health Device Count

	VnServicesFairHealthDeviceCount *float64 `json:"vnServicesFairHealthDeviceCount,omitempty"` // Vn Services Fair Health Device Count

	VnServicesNoHealthDeviceCount *float64 `json:"vnServicesNoHealthDeviceCount,omitempty"` // Vn Services No Health Device Count

	VnExitHealthPercentage *int `json:"vnExitHealthPercentage,omitempty"` // Vn Exit Health Percentage

	VnExitTotalDeviceCount *int `json:"vnExitTotalDeviceCount,omitempty"` // Vn Exit Total Device Count

	VnExitGoodHealthDeviceCount *int `json:"vnExitGoodHealthDeviceCount,omitempty"` // Vn Exit Good Health Device Count

	VnExitPoorHealthDeviceCount *int `json:"vnExitPoorHealthDeviceCount,omitempty"` // Vn Exit Poor Health Device Count

	VnExitFairHealthDeviceCount *float64 `json:"vnExitFairHealthDeviceCount,omitempty"` // Vn Exit Fair Health Device Count

	VnExitNoHealthDeviceCount *float64 `json:"vnExitNoHealthDeviceCount,omitempty"` // Vn Exit No Health Device Count

	VnStatusHealthPercentage *ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseVnStatusHealthPercentage `json:"vnStatusHealthPercentage,omitempty"` // Vn Status Health Percentage

	VnStatusTotalDeviceCount *float64 `json:"vnStatusTotalDeviceCount,omitempty"` // Vn Status Total Device Count

	VnStatusGoodHealthDeviceCount *float64 `json:"vnStatusGoodHealthDeviceCount,omitempty"` // Vn Status Good Health Device Count

	VnStatusPoorHealthDeviceCount *float64 `json:"vnStatusPoorHealthDeviceCount,omitempty"` // Vn Status Poor Health Device Count

	VnStatusFairHealthDeviceCount *float64 `json:"vnStatusFairHealthDeviceCount,omitempty"` // Vn Status Fair Health Device Count

	VnStatusNoHealthDeviceCount *float64 `json:"vnStatusNoHealthDeviceCount,omitempty"` // Vn Status No Health Device Count

	PubsubSessionGoodHealthPercentage *int `json:"pubsubSessionGoodHealthPercentage,omitempty"` // Pubsub Session Good Health Percentage

	PubsubSessionTotalDeviceCount *int `json:"pubsubSessionTotalDeviceCount,omitempty"` // Pubsub Session Total Device Count

	PubsubSessionGoodHealthDeviceCount *int `json:"pubsubSessionGoodHealthDeviceCount,omitempty"` // Pubsub Session Good Health Device Count

	PubsubSessionPoorHealthDeviceCount *int `json:"pubsubSessionPoorHealthDeviceCount,omitempty"` // Pubsub Session Poor Health Device Count

	PubsubSessionFairHealthDeviceCount *float64 `json:"pubsubSessionFairHealthDeviceCount,omitempty"` // Pubsub Session Fair Health Device Count

	PubsubSessionNoHealthDeviceCount *float64 `json:"pubsubSessionNoHealthDeviceCount,omitempty"` // Pubsub Session No Health Device Count

	MultiCastGoodHealthPercentage *ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseMultiCastGoodHealthPercentage `json:"multiCastGoodHealthPercentage,omitempty"` // Multi Cast Good Health Percentage

	MultiCastTotalDeviceCount *float64 `json:"multiCastTotalDeviceCount,omitempty"` // Multi Cast Total Device Count

	MultiCastGoodHealthDeviceCount *float64 `json:"multiCastGoodHealthDeviceCount,omitempty"` // Multi Cast Good Health Device Count

	MultiCastPoorHealthDeviceCount *float64 `json:"multiCastPoorHealthDeviceCount,omitempty"` // Multi Cast Poor Health Device Count

	MultiCastFairHealthDeviceCount *float64 `json:"multiCastFairHealthDeviceCount,omitempty"` // Multi Cast Fair Health Device Count

	MultiCastNoHealthDeviceCount *float64 `json:"multiCastNoHealthDeviceCount,omitempty"` // Multi Cast No Health Device Count

	InternetAvailGoodHealthPercentage *int `json:"internetAvailGoodHealthPercentage,omitempty"` // Internet Avail Good Health Percentage

	InternetAvailTotalDeviceCount *int `json:"internetAvailTotalDeviceCount,omitempty"` // Internet Avail Total Device Count

	InternetAvailGoodHealthDeviceCount *int `json:"internetAvailGoodHealthDeviceCount,omitempty"` // Internet Avail Good Health Device Count

	InternetAvailPoorHealthDeviceCount *int `json:"internetAvailPoorHealthDeviceCount,omitempty"` // Internet Avail Poor Health Device Count

	InternetAvailFairHealthDeviceCount *int `json:"internetAvailFairHealthDeviceCount,omitempty"` // Internet Avail Fair Health Device Count

	InternetAvailNoHealthDeviceCount *int `json:"internetAvailNoHealthDeviceCount,omitempty"` // Internet Avail No Health Device Count

	BgpPeerGoodHealthPercentage *int `json:"bgpPeerGoodHealthPercentage,omitempty"` // Bgp Peer Good Health Percentage

	BgpPeerTotalDeviceCount *int `json:"bgpPeerTotalDeviceCount,omitempty"` // Bgp Peer Total Device Count

	BgpPeerGoodHealthDeviceCount *int `json:"bgpPeerGoodHealthDeviceCount,omitempty"` // Bgp Peer Good Health Device Count

	BgpPeerPoorHealthDeviceCount *int `json:"bgpPeerPoorHealthDeviceCount,omitempty"` // Bgp Peer Poor Health Device Count

	BgpPeerFairHealthDeviceCount *float64 `json:"bgpPeerFairHealthDeviceCount,omitempty"` // Bgp Peer Fair Health Device Count

	BgpPeerNoHealthDeviceCount *float64 `json:"bgpPeerNoHealthDeviceCount,omitempty"` // Bgp Peer No Health Device Count

	VniGoodHealthPercentage *ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseVniGoodHealthPercentage `json:"vniGoodHealthPercentage,omitempty"` // Vni Good Health Percentage

	VniTotalDeviceCount *float64 `json:"vniTotalDeviceCount,omitempty"` // Vni Total Device Count

	VniGoodHealthDeviceCount *float64 `json:"vniGoodHealthDeviceCount,omitempty"` // Vni Good Health Device Count

	VniPoorHealthDeviceCount *float64 `json:"vniPoorHealthDeviceCount,omitempty"` // Vni Poor Health Device Count

	VniFairHealthDeviceCount *float64 `json:"vniFairHealthDeviceCount,omitempty"` // Vni Fair Health Device Count

	VniNoHealthDeviceCount *float64 `json:"vniNoHealthDeviceCount,omitempty"` // Vni No Health Device Count
}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseVLAN interface{}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseLayer interface{}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseAssociatedL3Vn interface{}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseVnStatusHealthPercentage interface{}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseMultiCastGoodHealthPercentage interface{}
type ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryResponseVniGoodHealthPercentage interface{}
type ResponseSdaReadVirtualNetworksCount struct {
	Response *ResponseSdaReadVirtualNetworksCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSdaReadVirtualNetworksCountResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromID struct {
	Response *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponse `json:"response,omitempty"` //
}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	Name string `json:"name,omitempty"` // Name

	NetworkProtocol string `json:"networkProtocol,omitempty"` // Network Protocol

	VLAN *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseVLAN `json:"vlan,omitempty"` // Vlan

	Vnid *int `json:"vnid,omitempty"` // Vnid

	Layer *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseLayer `json:"layer,omitempty"` // Layer

	TotalFabricSites *int `json:"totalFabricSites,omitempty"` // Total Fabric Sites

	AssociatedL3Vn *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseAssociatedL3Vn `json:"associatedL3Vn,omitempty"` // Associated L3 Vn

	TotalEndpoints *float64 `json:"totalEndpoints,omitempty"` // Total Endpoints

	GoodHealthPercentage *int `json:"goodHealthPercentage,omitempty"` // Good Health Percentage

	TotalDeviceCount *int `json:"totalDeviceCount,omitempty"` // Total Device Count

	GoodHealthDeviceCount *int `json:"goodHealthDeviceCount,omitempty"` // Good Health Device Count

	FairHealthDeviceCount *int `json:"fairHealthDeviceCount,omitempty"` // Fair Health Device Count

	PoorHealthDeviceCount *int `json:"poorHealthDeviceCount,omitempty"` // Poor Health Device Count

	NoHealthDeviceCount *int `json:"noHealthDeviceCount,omitempty"` // No Health Device Count

	VnFabricControlPlaneGoodHealthPercentage *int `json:"vnFabricControlPlaneGoodHealthPercentage,omitempty"` // Vn Fabric Control Plane Good Health Percentage

	VnFabricControlPlaneTotalDeviceCount *int `json:"vnFabricControlPlaneTotalDeviceCount,omitempty"` // Vn Fabric Control Plane Total Device Count

	VnFabricControlPlaneGoodHealthDeviceCount *int `json:"vnFabricControlPlaneGoodHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Good Health Device Count

	VnFabricControlPlanePoorHealthDeviceCount *int `json:"vnFabricControlPlanePoorHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Poor Health Device Count

	VnFabricControlPlaneFairHealthDeviceCount *float64 `json:"vnFabricControlPlaneFairHealthDeviceCount,omitempty"` // Vn Fabric Control Plane Fair Health Device Count

	VnFabricControlPlaneNoHealthDeviceCount *float64 `json:"vnFabricControlPlaneNoHealthDeviceCount,omitempty"` // Vn Fabric Control Plane No Health Device Count

	VnServicesHealthPercentage *int `json:"vnServicesHealthPercentage,omitempty"` // Vn Services Health Percentage

	VnServicesTotalDeviceCount *int `json:"vnServicesTotalDeviceCount,omitempty"` // Vn Services Total Device Count

	VnServicesGoodHealthDeviceCount *int `json:"vnServicesGoodHealthDeviceCount,omitempty"` // Vn Services Good Health Device Count

	VnServicesPoorHealthDeviceCount *int `json:"vnServicesPoorHealthDeviceCount,omitempty"` // Vn Services Poor Health Device Count

	VnServicesFairHealthDeviceCount *float64 `json:"vnServicesFairHealthDeviceCount,omitempty"` // Vn Services Fair Health Device Count

	VnServicesNoHealthDeviceCount *float64 `json:"vnServicesNoHealthDeviceCount,omitempty"` // Vn Services No Health Device Count

	VnExitHealthPercentage *int `json:"vnExitHealthPercentage,omitempty"` // Vn Exit Health Percentage

	VnExitTotalDeviceCount *int `json:"vnExitTotalDeviceCount,omitempty"` // Vn Exit Total Device Count

	VnExitGoodHealthDeviceCount *int `json:"vnExitGoodHealthDeviceCount,omitempty"` // Vn Exit Good Health Device Count

	VnExitPoorHealthDeviceCount *int `json:"vnExitPoorHealthDeviceCount,omitempty"` // Vn Exit Poor Health Device Count

	VnExitFairHealthDeviceCount *float64 `json:"vnExitFairHealthDeviceCount,omitempty"` // Vn Exit Fair Health Device Count

	VnExitNoHealthDeviceCount *float64 `json:"vnExitNoHealthDeviceCount,omitempty"` // Vn Exit No Health Device Count

	VnStatusHealthPercentage *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseVnStatusHealthPercentage `json:"vnStatusHealthPercentage,omitempty"` // Vn Status Health Percentage

	VnStatusTotalDeviceCount *float64 `json:"vnStatusTotalDeviceCount,omitempty"` // Vn Status Total Device Count

	VnStatusGoodHealthDeviceCount *float64 `json:"vnStatusGoodHealthDeviceCount,omitempty"` // Vn Status Good Health Device Count

	VnStatusPoorHealthDeviceCount *float64 `json:"vnStatusPoorHealthDeviceCount,omitempty"` // Vn Status Poor Health Device Count

	VnStatusFairHealthDeviceCount *float64 `json:"vnStatusFairHealthDeviceCount,omitempty"` // Vn Status Fair Health Device Count

	VnStatusNoHealthDeviceCount *float64 `json:"vnStatusNoHealthDeviceCount,omitempty"` // Vn Status No Health Device Count

	PubsubSessionGoodHealthPercentage *int `json:"pubsubSessionGoodHealthPercentage,omitempty"` // Pubsub Session Good Health Percentage

	PubsubSessionTotalDeviceCount *int `json:"pubsubSessionTotalDeviceCount,omitempty"` // Pubsub Session Total Device Count

	PubsubSessionGoodHealthDeviceCount *int `json:"pubsubSessionGoodHealthDeviceCount,omitempty"` // Pubsub Session Good Health Device Count

	PubsubSessionPoorHealthDeviceCount *int `json:"pubsubSessionPoorHealthDeviceCount,omitempty"` // Pubsub Session Poor Health Device Count

	PubsubSessionFairHealthDeviceCount *float64 `json:"pubsubSessionFairHealthDeviceCount,omitempty"` // Pubsub Session Fair Health Device Count

	PubsubSessionNoHealthDeviceCount *float64 `json:"pubsubSessionNoHealthDeviceCount,omitempty"` // Pubsub Session No Health Device Count

	MultiCastGoodHealthPercentage *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseMultiCastGoodHealthPercentage `json:"multiCastGoodHealthPercentage,omitempty"` // Multi Cast Good Health Percentage

	MultiCastTotalDeviceCount *float64 `json:"multiCastTotalDeviceCount,omitempty"` // Multi Cast Total Device Count

	MultiCastGoodHealthDeviceCount *float64 `json:"multiCastGoodHealthDeviceCount,omitempty"` // Multi Cast Good Health Device Count

	MultiCastPoorHealthDeviceCount *float64 `json:"multiCastPoorHealthDeviceCount,omitempty"` // Multi Cast Poor Health Device Count

	MultiCastFairHealthDeviceCount *float64 `json:"multiCastFairHealthDeviceCount,omitempty"` // Multi Cast Fair Health Device Count

	MultiCastNoHealthDeviceCount *float64 `json:"multiCastNoHealthDeviceCount,omitempty"` // Multi Cast No Health Device Count

	InternetAvailGoodHealthPercentage *int `json:"internetAvailGoodHealthPercentage,omitempty"` // Internet Availability Good Health Percentage

	InternetAvailTotalDeviceCount *int `json:"internetAvailTotalDeviceCount,omitempty"` // Internet Availability Total Device Count

	InternetAvailGoodHealthDeviceCount *int `json:"internetAvailGoodHealthDeviceCount,omitempty"` // Internet Availability Good Health Device Count

	InternetAvailPoorHealthDeviceCount *int `json:"internetAvailPoorHealthDeviceCount,omitempty"` // Internet Availability Poor Health Device Count

	InternetAvailFairHealthDeviceCount *int `json:"internetAvailFairHealthDeviceCount,omitempty"` // Internet Availability Fair Health Device Count

	InternetAvailNoHealthDeviceCount *int `json:"internetAvailNoHealthDeviceCount,omitempty"` // Internet Availability No Health Device Count

	BgpPeerGoodHealthPercentage *int `json:"bgpPeerGoodHealthPercentage,omitempty"` // Bgp Peer Good Health Percentage

	BgpPeerTotalDeviceCount *int `json:"bgpPeerTotalDeviceCount,omitempty"` // Bgp Peer Total Device Count

	BgpPeerGoodHealthDeviceCount *int `json:"bgpPeerGoodHealthDeviceCount,omitempty"` // Bgp Peer Good Health Device Count

	BgpPeerPoorHealthDeviceCount *int `json:"bgpPeerPoorHealthDeviceCount,omitempty"` // Bgp Peer Poor Health Device Count

	BgpPeerFairHealthDeviceCount *float64 `json:"bgpPeerFairHealthDeviceCount,omitempty"` // Bgp Peer Fair Health Device Count

	BgpPeerNoHealthDeviceCount *float64 `json:"bgpPeerNoHealthDeviceCount,omitempty"` // Bgp Peer No Health Device Count

	VniGoodHealthPercentage *ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseVniGoodHealthPercentage `json:"vniGoodHealthPercentage,omitempty"` // Vni Good Health Percentage

	VniTotalDeviceCount *float64 `json:"vniTotalDeviceCount,omitempty"` // Vni Total Device Count

	VniGoodHealthDeviceCount *float64 `json:"vniGoodHealthDeviceCount,omitempty"` // Vni Good Health Device Count

	VniPoorHealthDeviceCount *float64 `json:"vniPoorHealthDeviceCount,omitempty"` // Vni Poor Health Device Count

	VniFairHealthDeviceCount *float64 `json:"vniFairHealthDeviceCount,omitempty"` // Vni Fair Health Device Count

	VniNoHealthDeviceCount *float64 `json:"vniNoHealthDeviceCount,omitempty"` // Vni No Health Device Count
}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseVLAN interface{}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseLayer interface{}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseAssociatedL3Vn interface{}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseVnStatusHealthPercentage interface{}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseMultiCastGoodHealthPercentage interface{}
type ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromIDResponseVniGoodHealthPercentage interface{}
type ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange struct {
	Response *[]ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeResponse `json:"response,omitempty"` //

	Page *ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangePage `json:"page,omitempty"` //

	Version *int `json:"version,omitempty"` // Version
}
type ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Attributes *[]ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeResponseAttributes `json:"attributes,omitempty"` //
}
type ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeResponseAttributes struct {
	Name string `json:"name,omitempty"` // Name

	Value *int `json:"value,omitempty"` // Value
}
type ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *float64 `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name

	AuthenticationOrder string `json:"authenticationOrder,omitempty"` // Authentication Order

	Dot1XToMabFallbackTimeout string `json:"dot1xToMabFallbackTimeout,omitempty"` // Dot1x To Mab Fallback Timeout

	WakeOnLan *bool `json:"wakeOnLan,omitempty"` // Wake On Lan

	NumberOfHosts string `json:"numberOfHosts,omitempty"` // Number Of Hosts

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Authenticate Template info reterieved successfully in sda fabric site

	ExecutionID string `json:"executionId,omitempty"` // Execution Id
}
type ResponseSdaAddDefaultAuthenticationTemplateInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaAddBorderDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabric struct {
	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Description

	Payload *ResponseSdaGetBorderDeviceDetailFromSdaFabricPayload `json:"payload,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayload struct {
	ID string `json:"id,omitempty"` // Id

	InstanceID *int `json:"instanceId,omitempty"` // Instance Id

	AuthEntityID *int `json:"authEntityId,omitempty"` // Auth Entity Id

	DisplayName string `json:"displayName,omitempty"` // Display Name

	AuthEntityClass *int `json:"authEntityClass,omitempty"` // Auth Entity Class

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id

	DeployPending string `json:"deployPending,omitempty"` // Deploy Pending

	InstanceVersion *int `json:"instanceVersion,omitempty"` // Instance Version

	CreateTime *int `json:"createTime,omitempty"` // Create Time

	Deployed *bool `json:"deployed,omitempty"` // Deployed

	IsSeeded *bool `json:"isSeeded,omitempty"` // Is Seeded

	IsStale *bool `json:"isStale,omitempty"` // Is Stale

	LastUpdateTime *int `json:"lastUpdateTime,omitempty"` // Last Update Time

	Name string `json:"name,omitempty"` // Name

	Namespace string `json:"namespace,omitempty"` // Namespace

	ProvisioningState string `json:"provisioningState,omitempty"` // Provisioning State

	ResourceVersion *int `json:"resourceVersion,omitempty"` // Resource Version

	TargetIDList *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadTargetIDList `json:"targetIdList,omitempty"` // Target Id List

	Type string `json:"type,omitempty"` // Type

	CfsChangeInfo *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadCfsChangeInfo `json:"cfsChangeInfo,omitempty"` // Cfs Change Info

	CustomProvisions *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadCustomProvisions `json:"customProvisions,omitempty"` // Custom Provisions

	Configs *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadConfigs `json:"configs,omitempty"` // Configs

	ManagedSites *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadManagedSites `json:"managedSites,omitempty"` // Managed Sites

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	Roles []string `json:"roles,omitempty"` // Roles

	SaveWanConnectivityDetailsOnly *bool `json:"saveWanConnectivityDetailsOnly,omitempty"` // Save Wan Connectivity Details Only

	SiteID string `json:"siteId,omitempty"` // Site Id

	AkcSettingsCfs *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadAkcSettingsCfs `json:"akcSettingsCfs,omitempty"` // Akc Settings Cfs

	DeviceInterfaceInfo *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceInterfaceInfo `json:"deviceInterfaceInfo,omitempty"` // Device Interface Info

	DeviceSettings *ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettings `json:"deviceSettings,omitempty"` //

	NetworkWidesettings *ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettings `json:"networkWideSettings,omitempty"` //

	OtherDevice *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadOtherDevice `json:"otherDevice,omitempty"` // Other Device

	TransitNetworks *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadTransitNetworks `json:"transitNetworks,omitempty"` //

	VirtualNetwork *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadVirtualNetwork `json:"virtualNetwork,omitempty"` // Virtual Network

	WLAN *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadWLAN `json:"wlan,omitempty"` // Wlan
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadTargetIDList interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadCfsChangeInfo interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadCustomProvisions interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadConfigs interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadManagedSites interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadAkcSettingsCfs interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceInterfaceInfo interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettings struct {
	ID string `json:"id,omitempty"` // Id

	InstanceID *int `json:"instanceId,omitempty"` // Instance Id

	DisplayName string `json:"displayName,omitempty"` // Display Name

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id

	DeployPending string `json:"deployPending,omitempty"` // Deploy Pending

	InstanceVersion *int `json:"instanceVersion,omitempty"` // Instance Version

	ConnectedTo *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsConnectedTo `json:"connectedTo,omitempty"` // Connected To

	CPU *float64 `json:"cpu,omitempty"` // Cpu

	DhcpEnabled *bool `json:"dhcpEnabled,omitempty"` // Dhcp Enabled

	ExternalConnectivityIPPool string `json:"externalConnectivityIpPool,omitempty"` // External Connectivity Ip Pool

	ExternalDomainRoutingProtocol string `json:"externalDomainRoutingProtocol,omitempty"` // External Domain Routing Protocol

	InternalDomainProtocolNumber string `json:"internalDomainProtocolNumber,omitempty"` // Internal Domain Protocol Number

	Memory *float64 `json:"memory,omitempty"` // Memory

	NodeType []string `json:"nodeType,omitempty"` // Node Type

	Storage *float64 `json:"storage,omitempty"` // Storage

	ExtConnectivitySettings *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettings `json:"extConnectivitySettings,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsConnectedTo interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettings struct {
	ID string `json:"id,omitempty"` // Id

	InstanceID *int `json:"instanceId,omitempty"` // Instance Id

	DisplayName string `json:"displayName,omitempty"` // Display Name

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id

	DeployPending string `json:"deployPending,omitempty"` // Deploy Pending

	InstanceVersion *int `json:"instanceVersion,omitempty"` // Instance Version

	ExternalDomainProtocolNumber string `json:"externalDomainProtocolNumber,omitempty"` // External Domain Protocol Number

	InterfaceUUID string `json:"interfaceUuid,omitempty"` // Interface Uuid

	PolicyPropagationEnabled *bool `json:"policyPropagationEnabled,omitempty"` // Policy Propagation Enabled

	PolicySgtTag *float64 `json:"policySgtTag,omitempty"` // Policy Sgt Tag

	L2Handoff *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL2Handoff `json:"l2Handoff,omitempty"` // L2 Handoff

	L3Handoff *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL2Handoff interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3Handoff struct {
	ID string `json:"id,omitempty"` // Id

	InstanceID *int `json:"instanceId,omitempty"` // Instance Id

	DisplayName string `json:"displayName,omitempty"` // Display Name

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id

	DeployPending string `json:"deployPending,omitempty"` // Deploy Pending

	InstanceVersion *float64 `json:"instanceVersion,omitempty"` // Instance Version

	LocalIPAddress string `json:"localIpAddress,omitempty"` // Local Ip Address

	RemoteIPAddress string `json:"remoteIpAddress,omitempty"` // Remote Ip Address

	VLANID *int `json:"vlanId,omitempty"` // Vlan Id

	VirtualNetwork *ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettings struct {
	ID string `json:"id,omitempty"` // Id

	InstanceID *int `json:"instanceId,omitempty"` // Instance Id

	DisplayName string `json:"displayName,omitempty"` // Display Name

	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id

	DeployPending string `json:"deployPending,omitempty"` // Deploy Pending

	InstanceVersion *int `json:"instanceVersion,omitempty"` // Instance Version

	AAA *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsAAA `json:"aaa,omitempty"` // Aaa

	Cmx *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsCmx `json:"cmx,omitempty"` // Cmx

	Dhcp *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDhcp `json:"dhcp,omitempty"` //

	DNS *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDNS `json:"dns,omitempty"` //

	Ldap *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsLdap `json:"ldap,omitempty"` // Ldap

	NativeVLAN *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsNativeVLAN `json:"nativeVlan,omitempty"` // Native Vlan

	Netflow *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsNetflow `json:"netflow,omitempty"` // Netflow

	Ntp *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsNtp `json:"ntp,omitempty"` // Ntp

	SNMP *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsSNMP `json:"snmp,omitempty"` // Snmp

	Syslogs *[]ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsSyslogs `json:"syslogs,omitempty"` // Syslogs
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsAAA interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsCmx interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDhcp struct {
	ID string `json:"id,omitempty"` // Id

	IPAddress *ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDhcpIPAddress `json:"ipAddress,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDhcpIPAddress struct {
	ID string `json:"id,omitempty"` // Id

	PaddedAddress string `json:"paddedAddress,omitempty"` // Padded Address

	AddressType string `json:"addressType,omitempty"` // Address Type

	Address string `json:"address,omitempty"` // Address
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDNS struct {
	ID string `json:"id,omitempty"` // Id

	DomainName string `json:"domainName,omitempty"` // Domain Name

	IP *ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDNSIP `json:"ip,omitempty"` //
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsDNSIP struct {
	ID string `json:"id,omitempty"` // Id

	PaddedAddress string `json:"paddedAddress,omitempty"` // Padded Address

	AddressType string `json:"addressType,omitempty"` // Address Type

	Address string `json:"address,omitempty"` // Address
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsLdap interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsNativeVLAN interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsNetflow interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsNtp interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsSNMP interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadNetworkWidesettingsSyslogs interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadOtherDevice interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadTransitNetworks struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadVirtualNetwork interface{}
type ResponseSdaGetBorderDeviceDetailFromSdaFabricPayloadWLAN interface{}
type ResponseSdaDeleteBorderDeviceFromSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeleteControlPlaneDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetControlPlaneDeviceFromSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully

	DeviceName string `json:"deviceName,omitempty"` // Device Name

	Roles string `json:"roles,omitempty"` // Assigned roles

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	RouteDistributionProtocol string `json:"routeDistributionProtocol,omitempty"` // routeDistributionProtocol

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Control plane device info retrieved successfully in sda fabric
}
type ResponseSdaAddControlPlaneDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetDeviceInfoFromSdaFabric struct {
	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Description

	Name string `json:"name,omitempty"` // Name

	Roles []string `json:"roles,omitempty"` // Roles

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy
}
type ResponseSdaGetDeviceRoleInSdaFabric struct {
	Roles []string `json:"roles,omitempty"` // Assigned device roles. Possible roles are [Edge Node, Control Plane, Border Node, Extended Node, Wireless Controller, Transit Control Plane]

	Status string `json:"status,omitempty"` // Status indicates if API failed or passed.

	Description string `json:"description,omitempty"` // Device role successfully retrieved from sda fabric.
}
type ResponseSdaAddEdgeDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeleteEdgeDeviceFromSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetEdgeDeviceFromSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully

	DeviceName string `json:"deviceName,omitempty"` // Device Name

	Roles string `json:"roles,omitempty"` // Assigned roles

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	FabricSiteNameHierarchy string `json:"fabricSiteNameHierarchy,omitempty"` // Fabric Site Name Hierarchy

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Edge device info retrieved successfully in sda fabric
}
type ResponseSdaGetSiteFromSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	FabricName string `json:"fabricName,omitempty"` // Fabric Name

	FabricType string `json:"fabricType,omitempty"` // Fabric Type

	FabricDomainType string `json:"fabricDomainType,omitempty"` // Fabric Domain Type

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Fabric Site info successfully retrieved from sda fabric
}
type ResponseSdaDeleteSiteFromSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaAddSiteInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaAddPortAssignmentForAccessPointInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetPortAssignmentForAccessPointInSdaFabric struct {
	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Description

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	DataIPAddressPoolName string `json:"dataIpAddressPoolName,omitempty"` // Data Ip Address Pool Name

	VoiceIPAddressPoolName string `json:"voiceIpAddressPoolName,omitempty"` // Voice Ip Address Pool Name

	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group Name

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
}
type ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaAddPortAssignmentForUserDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric struct {
	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Description

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	DataIPAddressPoolName string `json:"dataIpAddressPoolName,omitempty"` // Data Ip Address Pool Name

	VoiceIPAddressPoolName string `json:"voiceIpAddressPoolName,omitempty"` // Voice Ip Address Pool Name

	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group Name

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
}
type ResponseSdaAddMulticastInSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetMulticastDetailsFromSdaFabric struct {
	MulticastMethod string `json:"multicastMethod,omitempty"` // Multicast Method

	MulticastType string `json:"multicastType,omitempty"` // Multicast Type

	MulticastVnInfo *[]ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfo `json:"multicastVnInfo,omitempty"` //

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // multicast configuration info retrieved successfully from sda fabric
}
type ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfo struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site

	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name, that is reserved to Fabric Site

	InternalRpIPAddress []string `json:"internalRpIpAddress,omitempty"` // InternalRpIpAddress

	ExternalRpIPAddress string `json:"externalRpIpAddress,omitempty"` // ExternalRpIpAddress

	SsmInfo *[]ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfoSsmInfo `json:"ssmInfo,omitempty"` //
}
type ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfoSsmInfo struct {
	SsmGroupRange string `json:"ssmGroupRange,omitempty"` // SSM group range

	SsmWildcardMask string `json:"ssmWildcardMask,omitempty"` // SSM Wildcard Mask
}
type ResponseSdaDeleteMulticastFromSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeleteProvisionedWiredDevice struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaReProvisionWiredDevice struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaProvisionWiredDevice struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetProvisionedWiredDevice struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the device to be provisioned

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy for device location(only building / floor level)

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Wired Provisioned device detail retrieved successfully
}
type ResponseSdaDeleteTransitPeerNetwork struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetTransitPeerNetworkInfo struct {
	TransitPeerNetworkName string `json:"transitPeerNetworkName,omitempty"` // Transit Peer Network Name

	TransitPeerNetworkType string `json:"transitPeerNetworkType,omitempty"` // Transit Peer Network Type

	IPTransitSettings *ResponseSdaGetTransitPeerNetworkInfoIPTransitSettings `json:"ipTransitSettings,omitempty"` //

	SdaTransitSettings *ResponseSdaGetTransitPeerNetworkInfoSdaTransitSettings `json:"sdaTransitSettings,omitempty"` //

	Status string `json:"status,omitempty"` // status

	Description string `json:"description,omitempty"` // Transit Peer network info retrieved successfully

	TransitPeerNetworkID string `json:"transitPeerNetworkId,omitempty"` // Transit Peer Network Id
}
type ResponseSdaGetTransitPeerNetworkInfoIPTransitSettings struct {
	RoutingProtocolName string `json:"routingProtocolName,omitempty"` // Routing Protocol Name

	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number
}
type ResponseSdaGetTransitPeerNetworkInfoSdaTransitSettings struct {
	TransitControlPlaneSettings *[]ResponseSdaGetTransitPeerNetworkInfoSdaTransitSettingsTransitControlPlaneSettings `json:"transitControlPlaneSettings,omitempty"` //
}
type ResponseSdaGetTransitPeerNetworkInfoSdaTransitSettingsTransitControlPlaneSettings struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
}
type ResponseSdaAddTransitPeerNetwork struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeleteVnFromSdaFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetVnFromSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name

	FabricName string `json:"fabricName,omitempty"` // Fabric Name

	IsInfraVn *bool `json:"isInfraVN,omitempty"` // Infra VN

	IsDefaultVn *bool `json:"isDefaultVN,omitempty"` // Default VN

	VirtualNetworkContextID string `json:"virtualNetworkContextId,omitempty"` // Virtual Network Context Id

	VirtualNetworkID string `json:"virtualNetworkId,omitempty"` // Virtual Network Id

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Virtual Network info retrieved successfully from SDA Fabric

	ExecutionID string `json:"executionId,omitempty"` // Execution Id
}
type ResponseSdaAddVnInFabric struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetVirtualNetworkSummary struct {
	VirtualNetworkCount *int `json:"virtualNetworkCount,omitempty"` // Virtual Networks Count

	VirtualNetworkSummary *[]ResponseSdaGetVirtualNetworkSummaryVirtualNetworkSummary `json:"virtualNetworkSummary,omitempty"` //

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Virtual Network summary retrieved successfully from SDA Fabric

	ExecutionID string `json:"executionId,omitempty"` // Execution Id
}
type ResponseSdaGetVirtualNetworkSummaryVirtualNetworkSummary struct {
	VirtualNetworkContextID string `json:"virtualNetworkContextId,omitempty"` // Virtual Network Context Id

	VirtualNetworkID string `json:"virtualNetworkId,omitempty"` // Virtual Network Id

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name

	Layer3Instance *int `json:"layer3Instance,omitempty"` // layer3 Instance

	VirtualNetworkStatus string `json:"virtualNetworkStatus,omitempty"` // Virtual Network Status
}
type ResponseSdaGetIPPoolFromSdaVirtualNetwork struct {
	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Description

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name

	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name

	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` // Authentication Policy Name

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type

	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group Name

	IsL2FloodingEnabled *bool `json:"isL2FloodingEnabled,omitempty"` // Is L2 Flooding Enabled

	IsThisCriticalPool *bool `json:"isThisCriticalPool,omitempty"` // Is This Critical Pool
}
type ResponseSdaDeleteIPPoolFromSdaVirtualNetwork struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaAddIPPoolInSdaVirtualNetwork struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaUpdateAnycastGateways struct {
	Response *ResponseSdaUpdateAnycastGatewaysResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateAnycastGatewaysResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaAddAnycastGateways struct {
	Response *ResponseSdaAddAnycastGatewaysResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddAnycastGatewaysResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetAnycastGateways struct {
	Response *[]ResponseSdaGetAnycastGatewaysResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetAnycastGatewaysResponse struct {
	ID string `json:"id,omitempty"` // ID of the anycast gateway.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this anycast gateway is assigned to.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the anycast gateway.

	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool associated with the anycast gateway.

	TCPMssAdjustment *int `json:"tcpMssAdjustment,omitempty"` // TCP maximum segment size adjustment.

	VLANName string `json:"vlanName,omitempty"` // Name of the VLAN of the anycast gateway.

	VLANID *int `json:"vlanId,omitempty"` // ID of the VLAN of the anycast gateway.

	TrafficType string `json:"trafficType,omitempty"` // The type of traffic the anycast gateway serves.

	PoolType string `json:"poolType,omitempty"` // The pool type of the anycast gateway (applicable only to INFRA_VN).

	SecurityGroupName string `json:"securityGroupName,omitempty"` // Name of the associated Security Group (not applicable to INFRA_VN).

	IsCriticalPool *bool `json:"isCriticalPool,omitempty"` // Enable/disable critical VLAN (not applicable to INFRA_VN).

	IsLayer2FloodingEnabled *bool `json:"isLayer2FloodingEnabled,omitempty"` // Enable/disable layer 2 flooding (not applicable to INFRA_VN).

	IsWirelessPool *bool `json:"isWirelessPool,omitempty"` // Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).

	IsIPDirectedBroadcast *bool `json:"isIpDirectedBroadcast,omitempty"` // Enable/disable IP-directed broadcast (not applicable to INFRA_VN).

	IsIntraSubnetRoutingEnabled *bool `json:"isIntraSubnetRoutingEnabled,omitempty"` // Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN).

	IsMultipleIPToMacAddresses *bool `json:"isMultipleIpToMacAddresses,omitempty"` // Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).

	IsSupplicantBasedExtendedNodeOnboarding *bool `json:"isSupplicantBasedExtendedNodeOnboarding,omitempty"` // Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN).

	IsGroupBasedPolicyEnforcementEnabled *bool `json:"isGroupBasedPolicyEnforcementEnabled,omitempty"` // Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN).
}
type ResponseSdaGetAnycastGatewayCount struct {
	Response *ResponseSdaGetAnycastGatewayCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetAnycastGatewayCountResponse struct {
	Count *int `json:"count,omitempty"` // The number of anycast gateways.
}
type ResponseSdaDeleteAnycastGatewayByID struct {
	Response *ResponseSdaDeleteAnycastGatewayByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteAnycastGatewayByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetAuthenticationProfiles struct {
	Response *[]ResponseSdaGetAuthenticationProfilesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetAuthenticationProfilesResponse struct {
	ID string `json:"id,omitempty"` // ID of the authentication profile.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this authentication profile is assigned to. (This property is not applicable to global authentication profiles and will not be present in such cases.)

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // The default host authentication template.

	AuthenticationOrder string `json:"authenticationOrder,omitempty"` // First authentication method.

	Dot1XToMabFallbackTimeout *int `json:"dot1xToMabFallbackTimeout,omitempty"` // 802.1x Timeout.

	WakeOnLan *bool `json:"wakeOnLan,omitempty"` // Wake on LAN.

	NumberOfHosts string `json:"numberOfHosts,omitempty"` // Number of Hosts.

	IsBpduGuardEnabled *bool `json:"isBpduGuardEnabled,omitempty"` // Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication".

	PreAuthACL *ResponseSdaGetAuthenticationProfilesResponsePreAuthACL `json:"preAuthAcl,omitempty"` //
}
type ResponseSdaGetAuthenticationProfilesResponsePreAuthACL struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable/disable Pre-Authentication ACL.

	ImplicitAction string `json:"implicitAction,omitempty"` // Implicit behaviour unless overridden.

	Description string `json:"description,omitempty"` // Description of this Pre-Authentication ACL.

	AccessContracts *[]ResponseSdaGetAuthenticationProfilesResponsePreAuthACLAccessContracts `json:"accessContracts,omitempty"` //
}
type ResponseSdaGetAuthenticationProfilesResponsePreAuthACLAccessContracts struct {
	Action string `json:"action,omitempty"` // Contract behaviour.

	Protocol string `json:"protocol,omitempty"` // Protocol for the access contract.

	Port string `json:"port,omitempty"` // Port for the access contract.
}
type ResponseSdaUpdateAuthenticationProfile struct {
	Response *ResponseSdaUpdateAuthenticationProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateAuthenticationProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaDeleteExtranetPolicies struct {
	Response *ResponseSdaDeleteExtranetPoliciesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteExtranetPoliciesResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaUpdateExtranetPolicy struct {
	Response *ResponseSdaUpdateExtranetPolicyResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateExtranetPolicyResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaAddExtranetPolicy struct {
	Response *ResponseSdaAddExtranetPolicyResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddExtranetPolicyResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetExtranetPolicies struct {
	Response *[]ResponseSdaGetExtranetPoliciesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetExtranetPoliciesResponse struct {
	ID string `json:"id,omitempty"` // ID of the extranet policy.

	ExtranetPolicyName string `json:"extranetPolicyName,omitempty"` // Name of the extranet policy.

	FabricIDs []string `json:"fabricIds,omitempty"` // IDs of the fabric sites associated with this extranet policy.

	ProviderVirtualNetworkName string `json:"providerVirtualNetworkName,omitempty"` // Name of the provider virtual network.

	SubscriberVirtualNetworkNames []string `json:"subscriberVirtualNetworkNames,omitempty"` // Name of the subscriber virtual network names.
}
type ResponseSdaGetExtranetPolicyCount struct {
	Response *ResponseSdaGetExtranetPolicyCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetExtranetPolicyCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of extranet policies.
}
type ResponseSdaDeleteExtranetPolicyByID struct {
	Response *ResponseSdaDeleteExtranetPolicyByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteExtranetPolicyByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetFabricDevices struct {
	Response *[]ResponseSdaGetFabricDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesResponse struct {
	ID string `json:"id,omitempty"` // ID of the fabric device.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric of this fabric device.

	DeviceRoles []string `json:"deviceRoles,omitempty"` // List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE, EXTENDED_NODE].

	BorderDeviceSettings *ResponseSdaGetFabricDevicesResponseBorderDeviceSettings `json:"borderDeviceSettings,omitempty"` //
}
type ResponseSdaGetFabricDevicesResponseBorderDeviceSettings struct {
	BorderTypes []string `json:"borderTypes,omitempty"` // List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].

	Layer3Settings *ResponseSdaGetFabricDevicesResponseBorderDeviceSettingsLayer3Settings `json:"layer3Settings,omitempty"` //
}
type ResponseSdaGetFabricDevicesResponseBorderDeviceSettingsLayer3Settings struct {
	LocalAutonomousSystemNumber string `json:"localAutonomousSystemNumber,omitempty"` // BGP Local autonomous system number of the fabric border device.

	IsDefaultExit *bool `json:"isDefaultExit,omitempty"` // Is default exit value of the fabric border device.

	ImportExternalRoutes *bool `json:"importExternalRoutes,omitempty"` // Import external routes value of the fabric border device.

	BorderPriority *int `json:"borderPriority,omitempty"` // Border priority of the fabric border device.  A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5.

	PrependAutonomousSystemCount *int `json:"prependAutonomousSystemCount,omitempty"` // Prepend autonomous system count of the fabric border device.
}
type ResponseSdaUpdateFabricDevices struct {
	Response *ResponseSdaUpdateFabricDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaUpdateFabricDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaDeleteFabricDevices struct {
	Response *ResponseSdaDeleteFabricDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaAddFabricDevices struct {
	Response *ResponseSdaAddFabricDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaAddFabricDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricDevicesCount struct {
	Response *ResponseSdaGetFabricDevicesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of fabric devices.
}
type ResponseSdaDeleteFabricDeviceLayer2Handoffs struct {
	Response *ResponseSdaDeleteFabricDeviceLayer2HandoffsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffsResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer2Handoffs struct {
	Response *[]ResponseSdaGetFabricDevicesLayer2HandoffsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsResponse struct {
	ID string `json:"id,omitempty"` // ID of the layer 2 handoff of a fabric device.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the layer 2 handoff.

	InternalVLANID *int `json:"internalVlanId,omitempty"` // VLAN number associated with this fabric.

	ExternalVLANID *int `json:"externalVlanId,omitempty"` // External VLAN number into which the fabric is extended.
}
type ResponseSdaAddFabricDevicesLayer2Handoffs struct {
	Response *ResponseSdaAddFabricDevicesLayer2HandoffsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaAddFabricDevicesLayer2HandoffsResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsCount struct {
	Response *ResponseSdaGetFabricDevicesLayer2HandoffsCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesLayer2HandoffsCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of fabric device layer 2 handoffs.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffByID struct {
	Response *ResponseSdaDeleteFabricDeviceLayer2HandoffByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer2HandoffByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransit struct {
	Response *ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransitResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit struct {
	Response *ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransitResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransit struct {
	Response *ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransitResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransit struct {
	Response *[]ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitResponse struct {
	ID string `json:"id,omitempty"` // ID of the fabric device layer 3 handoff ip transit.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to.

	TransitNetworkID string `json:"transitNetworkId,omitempty"` // ID of the transit network of the layer 3 handoff ip transit.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the layer 3 handoff ip transit.

	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External connectivity ip pool is used by Catalyst Center to allocate IP address for the connection between the border node and peer.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with this fabric site.

	VLANID *int `json:"vlanId,omitempty"` // VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).

	TCPMssAdjustment *int `json:"tcpMssAdjustment,omitempty"` // TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.

	LocalIPAddress string `json:"localIpAddress,omitempty"` // Local ipv4 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.

	RemoteIPAddress string `json:"remoteIpAddress,omitempty"` // Remote ipv4 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.

	LocalIPv6Address string `json:"localIpv6Address,omitempty"` // Local ipv6 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.

	RemoteIPv6Address string `json:"remoteIpv6Address,omitempty"` // Remote ipv6 address for the selected virtual network. IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if an external connectivity ip pool name is present.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCount struct {
	Response *ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of fabric device layer 3 handoffs with IP transit.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByID struct {
	Response *ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit struct {
	Response *ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransitResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransit struct {
	Response *[]ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to.

	TransitNetworkID string `json:"transitNetworkId,omitempty"` // ID of the transit network of the layer 3 handoff sda transit.

	AffinityIDPrime *int `json:"affinityIdPrime,omitempty"` // Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.

	AffinityIDDecider *int `json:"affinityIdDecider,omitempty"` // Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.

	ConnectedToInternet *bool `json:"connectedToInternet,omitempty"` // True value for this allows associated site to provide internet access to other sites through sd-access.

	IsMulticastOverTransitEnabled *bool `json:"isMulticastOverTransitEnabled,omitempty"` // True value for this configures native multicast over multiple sites that are connected to an sd-access transit.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransit struct {
	Response *ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransitResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransit struct {
	Response *ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransitResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCount struct {
	Response *ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of fabric device layer 3 handoffs with sda transit.
}
type ResponseSdaDeleteFabricDeviceByID struct {
	Response *ResponseSdaDeleteFabricDeviceByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteFabricDeviceByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetFabricSites struct {
	Response *[]ResponseSdaGetFabricSitesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricSitesResponse struct {
	ID string `json:"id,omitempty"` // ID of the fabric site.

	SiteID string `json:"siteId,omitempty"` // ID of the network hierarchy.

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.

	IsPubSubEnabled *bool `json:"isPubSubEnabled,omitempty"` // Specifies whether this fabric site will use pub/sub for control nodes.
}
type ResponseSdaAddFabricSite struct {
	Response *ResponseSdaAddFabricSiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddFabricSiteResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaUpdateFabricSite struct {
	Response *ResponseSdaUpdateFabricSiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateFabricSiteResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetFabricSiteCount struct {
	Response *ResponseSdaGetFabricSiteCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricSiteCountResponse struct {
	Count *int `json:"count,omitempty"` // The number of fabric sites.
}
type ResponseSdaDeleteFabricSiteByID struct {
	Response *ResponseSdaDeleteFabricSiteByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteFabricSiteByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetFabricZones struct {
	Response *[]ResponseSdaGetFabricZonesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricZonesResponse struct {
	ID string `json:"id,omitempty"` // ID of the fabric zone.

	SiteID string `json:"siteId,omitempty"` // ID of the network hierarchy.

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
}
type ResponseSdaUpdateFabricZone struct {
	Response *ResponseSdaUpdateFabricZoneResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateFabricZoneResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaAddFabricZone struct {
	Response *ResponseSdaAddFabricZoneResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddFabricZoneResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetFabricZoneCount struct {
	Response *ResponseSdaGetFabricZoneCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetFabricZoneCountResponse struct {
	Count *int `json:"count,omitempty"` // The number of fabric zones.
}
type ResponseSdaDeleteFabricZoneByID struct {
	Response *ResponseSdaDeleteFabricZoneByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteFabricZoneByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaAddLayer2VirtualNetworks struct {
	Response *ResponseSdaAddLayer2VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddLayer2VirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaDeleteLayer2VirtualNetworks struct {
	Response *ResponseSdaDeleteLayer2VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteLayer2VirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetLayer2VirtualNetworks struct {
	Response *[]ResponseSdaGetLayer2VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetLayer2VirtualNetworksResponse struct {
	ID string `json:"id,omitempty"` // ID of the layer 2 virtual network.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this layer 2 virtual network is assigned to.

	VLANName string `json:"vlanName,omitempty"` // Name of the VLAN of the layer 2 virtual network.

	VLANID *int `json:"vlanId,omitempty"` // ID of the VLAN of the layer 2 virtual network.

	TrafficType string `json:"trafficType,omitempty"` // The type of traffic that is served.

	IsFabricEnabledWireless *bool `json:"isFabricEnabledWireless,omitempty"` // Set to true to enable wireless.

	IsMultipleIPToMacAddresses *bool `json:"isMultipleIpToMacAddresses,omitempty"` // Set to true to enable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine). This field will only be present on layer 2 virtual networks associated with a layer 3 virtual network.

	AssociatedLayer3VirtualNetworkName string `json:"associatedLayer3VirtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring.
}
type ResponseSdaUpdateLayer2VirtualNetworks struct {
	Response *ResponseSdaUpdateLayer2VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateLayer2VirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetLayer2VirtualNetworkCount struct {
	Response *ResponseSdaGetLayer2VirtualNetworkCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetLayer2VirtualNetworkCountResponse struct {
	Count *int `json:"count,omitempty"` // The number of layer 2 virtual networks
}
type ResponseSdaDeleteLayer2VirtualNetworkByID struct {
	Response *ResponseSdaDeleteLayer2VirtualNetworkByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteLayer2VirtualNetworkByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaAddLayer3VirtualNetworks struct {
	Response *ResponseSdaAddLayer3VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddLayer3VirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetLayer3VirtualNetworks struct {
	Response *[]ResponseSdaGetLayer3VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetLayer3VirtualNetworksResponse struct {
	ID string `json:"id,omitempty"` // ID of the layer 3 virtual network.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network.

	FabricIDs []string `json:"fabricIds,omitempty"` // IDs of the fabrics this layer 3 virtual network is assigned to.

	AnchoredSiteID string `json:"anchoredSiteId,omitempty"` // Fabric ID of the fabric site this layer 3 virtual network is anchored at.
}
type ResponseSdaDeleteLayer3VirtualNetworks struct {
	Response *ResponseSdaDeleteLayer3VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteLayer3VirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaUpdateLayer3VirtualNetworks struct {
	Response *ResponseSdaUpdateLayer3VirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateLayer3VirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetLayer3VirtualNetworksCount struct {
	Response *ResponseSdaGetLayer3VirtualNetworksCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetLayer3VirtualNetworksCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of layer 3 virtual networks.
}
type ResponseSdaDeleteLayer3VirtualNetworkByID struct {
	Response *ResponseSdaDeleteLayer3VirtualNetworkByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteLayer3VirtualNetworkByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaUpdateMulticast struct {
	Response *ResponseSdaUpdateMulticastResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateMulticastResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetMulticast struct {
	Response *[]ResponseSdaGetMulticastResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetMulticastResponse struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric site.

	ReplicationMode string `json:"replicationMode,omitempty"` // Replication Mode deployed in the fabric site.
}
type ResponseSdaAddMulticastVirtualNetworks struct {
	Response *ResponseSdaAddMulticastVirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddMulticastVirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetMulticastVirtualNetworks struct {
	Response *[]ResponseSdaGetMulticastVirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetMulticastVirtualNetworksResponse struct {
	ID string `json:"id,omitempty"` // ID of the multicast configuration.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric site.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the virtual network.

	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP Pool.

	IPv4SsmRanges []string `json:"ipv4SsmRanges,omitempty"` // IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.

	MulticastRPs *[]ResponseSdaGetMulticastVirtualNetworksResponseMulticastRPs `json:"multicastRPs,omitempty"` //
}
type ResponseSdaGetMulticastVirtualNetworksResponseMulticastRPs struct {
	RpDeviceLocation string `json:"rpDeviceLocation,omitempty"` // Device location of the RP.

	IPv4Address string `json:"ipv4Address,omitempty"` // IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.

	IPv6Address string `json:"ipv6Address,omitempty"` // IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.

	IsDefaultV4RP *bool `json:"isDefaultV4RP,omitempty"` // Specifies whether it is a default IPv4 RP.

	IsDefaultV6RP *bool `json:"isDefaultV6RP,omitempty"` // Specifies whether it is a default IPv6 RP.

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.

	IPv4AsmRanges []string `json:"ipv4AsmRanges,omitempty"` // IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.

	IPv6AsmRanges []string `json:"ipv6AsmRanges,omitempty"` // IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
}
type ResponseSdaUpdateMulticastVirtualNetworks struct {
	Response *ResponseSdaUpdateMulticastVirtualNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdateMulticastVirtualNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetMulticastVirtualNetworkCount struct {
	Response *ResponseSdaGetMulticastVirtualNetworkCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetMulticastVirtualNetworkCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of multicast configurations.
}
type ResponseSdaDeleteMulticastVirtualNetworkByID struct {
	Response *ResponseSdaDeleteMulticastVirtualNetworkByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteMulticastVirtualNetworkByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetPendingFabricEvents struct {
	Response *[]ResponseSdaGetPendingFabricEventsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetPendingFabricEventsResponse struct {
	ID string `json:"id,omitempty"` // ID of the pending fabric event.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric.

	Detail string `json:"detail,omitempty"` // Detail of the pending fabric event.
}
type ResponseSdaApplyPendingFabricEvents struct {
	Response *ResponseSdaApplyPendingFabricEventsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaApplyPendingFabricEventsResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaAddPortAssignments struct {
	Response *ResponseSdaAddPortAssignmentsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddPortAssignmentsResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetPortAssignments struct {
	Response *[]ResponseSdaGetPortAssignmentsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetPortAssignmentsResponse struct {
	ID string `json:"id,omitempty"` // ID of the port assignment.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric the device is assigned to.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the port assignment.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the port assignment.

	ConnectedDeviceType string `json:"connectedDeviceType,omitempty"` // Connected device type of the port assignment.

	DataVLANName string `json:"dataVlanName,omitempty"` // Data VLAN name of the port assignment.

	VoiceVLANName string `json:"voiceVlanName,omitempty"` // Voice VLAN name of the port assignment.

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate template name of the port assignment.

	SecurityGroupName string `json:"securityGroupName,omitempty"` // Security group name of the port assignment.

	InterfaceDescription string `json:"interfaceDescription,omitempty"` // Interface description of the port assignment.
}
type ResponseSdaUpdatePortAssignments struct {
	Response *ResponseSdaUpdatePortAssignmentsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdatePortAssignmentsResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaDeletePortAssignments struct {
	Response *ResponseSdaDeletePortAssignmentsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeletePortAssignmentsResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetPortAssignmentCount struct {
	Response *ResponseSdaGetPortAssignmentCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetPortAssignmentCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of port assignments.
}
type ResponseSdaDeletePortAssignmentByID struct {
	Response *ResponseSdaDeletePortAssignmentByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeletePortAssignmentByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetPortChannels struct {
	Response *[]ResponseSdaGetPortChannelsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetPortChannelsResponse struct {
	ID string `json:"id,omitempty"` // ID of the port channel.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric the device is assigned to.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device.

	PortChannelName string `json:"portChannelName,omitempty"` // Name of the port channel.

	InterfaceNames []string `json:"interfaceNames,omitempty"` // Interface names of this port channel.

	ConnectedDeviceType string `json:"connectedDeviceType,omitempty"` // Connected device type of the port channel.

	Protocol string `json:"protocol,omitempty"` // Protocol of the port channel.

	Description string `json:"description,omitempty"` // Description of the port channel.
}
type ResponseSdaAddPortChannels struct {
	Response *ResponseSdaAddPortChannelsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaAddPortChannelsResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaUpdatePortChannels struct {
	Response *ResponseSdaUpdatePortChannelsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaUpdatePortChannelsResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaDeletePortChannels struct {
	Response *ResponseSdaDeletePortChannelsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeletePortChannelsResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetPortChannelCount struct {
	Response *ResponseSdaGetPortChannelCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetPortChannelCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of port channels.
}
type ResponseSdaDeletePortChannelByID struct {
	Response *ResponseSdaDeletePortChannelByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeletePortChannelByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaDeleteProvisionedDevices struct {
	Response *ResponseSdaDeleteProvisionedDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteProvisionedDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaProvisionDevices struct {
	Response *ResponseSdaProvisionDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaProvisionDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetProvisionedDevices struct {
	Response *[]ResponseSdaGetProvisionedDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetProvisionedDevicesResponse struct {
	ID string `json:"id,omitempty"` // ID of the provisioned device.

	SiteID string `json:"siteId,omitempty"` // ID of the site this device is provisioned to.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device.
}
type ResponseSdaReProvisionDevices struct {
	Response *ResponseSdaReProvisionDevicesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaReProvisionDevicesResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaGetProvisionedDevicesCount struct {
	Response *ResponseSdaGetProvisionedDevicesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetProvisionedDevicesCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of provisioned devices.
}
type ResponseSdaDeleteProvisionedDeviceByID struct {
	Response *ResponseSdaDeleteProvisionedDeviceByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // version number.
}
type ResponseSdaDeleteProvisionedDeviceByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // id of the submitted task.

	URL string `json:"url,omitempty"` // url for querying the task's status.
}
type ResponseSdaUpdateTransitNetworks struct {
	Response *ResponseSdaUpdateTransitNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaUpdateTransitNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetTransitNetworks struct {
	Response *[]ResponseSdaGetTransitNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetTransitNetworksResponse struct {
	ID string `json:"id,omitempty"` // ID of the transit network.

	Name string `json:"name,omitempty"` // Name of the transit network.

	Type string `json:"type,omitempty"` // Type of the transit network.

	IPTransitSettings *ResponseSdaGetTransitNetworksResponseIPTransitSettings `json:"ipTransitSettings,omitempty"` //

	SdaTransitSettings *ResponseSdaGetTransitNetworksResponseSdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type ResponseSdaGetTransitNetworksResponseIPTransitSettings struct {
	RoutingProtocolName string `json:"routingProtocolName,omitempty"` // Routing Protocol Name of the IP transit network.

	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number of the IP transit network. Allowed range is [1 to 4294967295].
}
type ResponseSdaGetTransitNetworksResponseSdaTransitSettings struct {
	IsMulticastOverTransitEnabled *bool `json:"isMulticastOverTransitEnabled,omitempty"` // This indicates that multicast is enabled over SD-Access Transit. This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.

	ControlPlaneNetworkDeviceIDs []string `json:"controlPlaneNetworkDeviceIds,omitempty"` // List of network device IDs that are used as control plane nodes.
}
type ResponseSdaAddTransitNetworks struct {
	Response *ResponseSdaAddTransitNetworksResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaAddTransitNetworksResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaGetTransitNetworksCount struct {
	Response *ResponseSdaGetTransitNetworksCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaGetTransitNetworksCountResponse struct {
	Count *int `json:"count,omitempty"` // Number of transit networks.
}
type ResponseSdaDeleteTransitNetworkByID struct {
	Response *ResponseSdaDeleteTransitNetworkByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version number.
}
type ResponseSdaDeleteTransitNetworkByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // ID of the task.

	URL string `json:"url,omitempty"` // Task status lookup url.
}
type ResponseSdaAddVirtualNetworkWithScalableGroups struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaDeleteVirtualNetworkWithScalableGroups struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type ResponseSdaGetVirtualNetworkWithScalableGroups struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name to be assigned at global level

	IsGuestVirtualNetwork *bool `json:"isGuestVirtualNetwork,omitempty"` // Guest Virtual Network

	ScalableGroupNames []string `json:"scalableGroupNames,omitempty"` // Scalable Group Names

	VManageVpnID string `json:"vManageVpnId,omitempty"` // vManage vpn id for SD-WAN

	VirtualNetworkContextID string `json:"virtualNetworkContextId,omitempty"` // Virtual Network Context Id for Global Virtual Network

	Status string `json:"status,omitempty"` // Status

	Description string `json:"description,omitempty"` // Virtual network info retrieved successfully

	ExecutionID string `json:"executionId,omitempty"` // Execution Id
}
type ResponseSdaUpdateVirtualNetworkWithScalableGroups struct {
	Status string `json:"status,omitempty"` // represents return status of API. status=success when API completed successfully, status=failed when API failed and has not completed the user request, status=pending when API execution is still in progression and user needs to track its further progress via taskId field.

	Description string `json:"description,omitempty"` // provides detailed information for API success or failure.

	TaskID string `json:"taskId,omitempty"` // Catalyst Center taskId that carried out the API execution. It will be provided if task was generated by API. For failed status, taskId may not be available

	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // /dna/intent/api/v1/task/<taskId> , allows API progression via taskId for async API

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // /dna/intent/api/v1/dnacaap/management/execution-status/<executionId>

	ExecutionID string `json:"executionId,omitempty"` // uuid for API execution status
}
type RequestSdaAddDefaultAuthenticationTemplateInSdaFabric []RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric // Array of RequestSdaAddDefaultAuthenticationTemplateInSDAFabric
type RequestItemSdaAddDefaultAuthenticationTemplateInSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
}
type RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric []RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric // Array of RequestSdaUpdateDefaultAuthenticationProfileInSDAFabric
type RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name

	AuthenticationOrder string `json:"authenticationOrder,omitempty"` // Authentication Order

	Dot1XToMabFallbackTimeout string `json:"dot1xToMabFallbackTimeout,omitempty"` // Dot1x To MabFallback Timeout( Allowed range is [3-120])

	WakeOnLan *bool `json:"wakeOnLan,omitempty"` // Wake On Lan

	NumberOfHosts string `json:"numberOfHosts,omitempty"` // Number Of Hosts
}
type RequestSdaAddBorderDeviceInSdaFabric []RequestItemSdaAddBorderDeviceInSdaFabric // Array of RequestSdaAddBorderDeviceInSDAFabric
type RequestItemSdaAddBorderDeviceInSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the provisioned Device

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy of provisioned Device(site should be part of Fabric Site)

	DeviceRole []string `json:"deviceRole,omitempty"` // Supported Device Roles in SD-Access fabric. Allowed roles are "Border_Node","Control_Plane_Node","Edge_Node". E.g. ["Border_Node"] or ["Border_Node", "Control_Plane_Node"] or ["Border_Node", "Control_Plane_Node","Edge_Node"]

	RouteDistributionProtocol string `json:"routeDistributionProtocol,omitempty"` // Route Distribution Protocol for Control Plane Device. Allowed values are "LISP_BGP" or "LISP_PUB_SUB". Default value is "LISP_BGP"

	ExternalDomainRoutingProtocolName string `json:"externalDomainRoutingProtocolName,omitempty"` // External Domain Routing Protocol Name

	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External Connectivity IpPool Name

	InternalAutonomouSystemNumber string `json:"internalAutonomouSystemNumber,omitempty"` // Internal Autonomous System Number

	BorderPriority string `json:"borderPriority,omitempty"` // Border priority associated with a given device. Allowed range for Border Priority is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.

	BorderSessionType string `json:"borderSessionType,omitempty"` // Border Session Type

	ConnectedToInternet *bool `json:"connectedToInternet,omitempty"` // Connected to Internet

	SdaTransitNetworkName string `json:"sdaTransitNetworkName,omitempty"` // SD-Access Transit Network Name

	BorderWithExternalConnectivity *bool `json:"borderWithExternalConnectivity,omitempty"` // Border With External Connectivity (Note: True for transit and False for non-transit border)

	ExternalConnectivitySettings *[]RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettings `json:"externalConnectivitySettings,omitempty"` //
}
type RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettings struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	InterfaceDescription string `json:"interfaceDescription,omitempty"` // Interface Description

	ExternalAutonomouSystemNumber string `json:"externalAutonomouSystemNumber,omitempty"` // External Autonomous System Number peer (e.g.,1-65535)

	L3Handoff *[]RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"` //

	L2Handoff *[]RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff `json:"l2Handoff,omitempty"` //
}
type RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff struct {
	VirtualNetwork *RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"` //
}
type RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site

	VLANID string `json:"vlanId,omitempty"` // Vlan Id (e.g.,2-4096 except for reserved VLANs (1002-1005, 2046, 4095))
}
type RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site

	VLANName string `json:"vlanName,omitempty"` // Vlan Name of L2 Handoff
}
type RequestSdaAddControlPlaneDeviceInSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // siteNameHierarchy of the Provisioned Device(site should be part of Fabric Site)

	RouteDistributionProtocol string `json:"routeDistributionProtocol,omitempty"` // Route Distribution Protocol for Control Plane Device. Allowed values are "LISP_BGP" or "LISP_PUB_SUB". Default value is "LISP_BGP"
}
type RequestSdaAddEdgeDeviceInSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Device which is provisioned successfully

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // siteNameHierarchy of the Provisioned Device(site should be part of Fabric Site)
}
type RequestSdaAddSiteInSdaFabric struct {
	FabricName string `json:"fabricName,omitempty"` // Warning - Starting DNA Center 2.2.3.5 release, this field has been deprecated. SD-Access Fabric does not need it anymore.  It will be removed in future DNA Center releases.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Existing site name hierarchy available at global level. For Example "Global/Chicago/Building21/Floor1"

	FabricType string `json:"fabricType,omitempty"` // Type of SD-Access Fabric. Allowed values are "FABRIC_SITE" or "FABRIC_ZONE".  Default value is "FABRIC_SITE".
}
type RequestSdaAddPortAssignmentForAccessPointInSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the edge device

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name of the edge device

	DataIPAddressPoolName string `json:"dataIpAddressPoolName,omitempty"` // Ip Pool Name, that is assigned to INFRA_VN

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate TemplateName associated to Fabric Site

	InterfaceDescription string `json:"interfaceDescription,omitempty"` // Details or note of interface port assignment
}
type RequestSdaAddPortAssignmentForUserDeviceInSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Complete Path of SD-Access Fabric Site.

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the Edge Node Device.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name on the Edge Node Device.

	InterfaceNames []string `json:"interfaceNames,omitempty"` // List of Interface Names on the Edge Node Device. E.g.["GigabitEthernet1/0/3","GigabitEthernet1/0/4"]

	DataIPAddressPoolName string `json:"dataIpAddressPoolName,omitempty"` // Ip Pool Name, that is assigned to virtual network with traffic type as DATA(can't be empty if voiceIpAddressPoolName is empty)

	VoiceIPAddressPoolName string `json:"voiceIpAddressPoolName,omitempty"` // Ip Pool Name, that is assigned to virtual network with traffic type as VOICE(can't be empty if dataIpAddressPoolName is empty)

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate TemplateName associated with siteNameHierarchy

	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group name associated with VN

	InterfaceDescription string `json:"interfaceDescription,omitempty"` // User defined text message for port assignment
}
type RequestSdaAddMulticastInSdaFabric struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Full path of sda Fabric Site

	MulticastMethod string `json:"multicastMethod,omitempty"` // Multicast Method

	MulticastType string `json:"multicastType,omitempty"` // Multicast Type

	MulticastVnInfo *[]RequestSdaAddMulticastInSdaFabricMulticastVnInfo `json:"multicastVnInfo,omitempty"` //
}
type RequestSdaAddMulticastInSdaFabricMulticastVnInfo struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site

	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name, that is reserved to Fabric Site

	InternalRpIPAddress []string `json:"internalRpIpAddress,omitempty"` // InternalRpIpAddress, required if multicastType is asm_with_internal_rp

	ExternalRpIPAddress string `json:"externalRpIpAddress,omitempty"` // ExternalRpIpAddress, required if multicastType is asm_with_external_rp

	SsmInfo *[]RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo `json:"ssmInfo,omitempty"` //
}
type RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo struct {
	SsmGroupRange string `json:"ssmGroupRange,omitempty"` // Valid SSM group range ip address(e.g., 230.0.0.0)

	SsmWildcardMask string `json:"ssmWildcardMask,omitempty"` // Valid SSM Wildcard Mask ip address(e.g.,0.255.255.255)
}
type RequestSdaReProvisionWiredDevice struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the device to be re-provisioned

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // siteNameHierarchy of the provisioned device
}
type RequestSdaProvisionWiredDevice struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Management Ip Address of the device to be provisioned

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy for device location(only building / floor level)
}
type RequestSdaAddTransitPeerNetwork struct {
	TransitPeerNetworkName string `json:"transitPeerNetworkName,omitempty"` // Transit Peer Network Name

	TransitPeerNetworkType string `json:"transitPeerNetworkType,omitempty"` // Transit Peer Network Type

	IPTransitSettings *RequestSdaAddTransitPeerNetworkIPTransitSettings `json:"ipTransitSettings,omitempty"` //

	SdaTransitSettings *RequestSdaAddTransitPeerNetworkSdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type RequestSdaAddTransitPeerNetworkIPTransitSettings struct {
	RoutingProtocolName string `json:"routingProtocolName,omitempty"` // Routing Protocol Name

	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number
}
type RequestSdaAddTransitPeerNetworkSdaTransitSettings struct {
	TransitControlPlaneSettings *[]RequestSdaAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings `json:"transitControlPlaneSettings,omitempty"` //
}
type RequestSdaAddTransitPeerNetworkSdaTransitSettingsTransitControlPlaneSettings struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy where device is provisioned

	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address of provisioned device
}
type RequestSdaAddVnInFabric struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is created at Global level

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site
}
type RequestSdaAddIPPoolInSdaVirtualNetwork struct {
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Path of sda Fabric Site

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name, that is associated to Fabric Site

	IsLayer2Only *bool `json:"isLayer2Only,omitempty"` // Layer2 Only enablement flag and default value is False

	IPPoolName string `json:"ipPoolName,omitempty"` // Ip Pool Name, that is reserved to Fabric Site (Required for L3 and INFRA_VN)

	VLANID string `json:"vlanId,omitempty"` // vlan Id(applicable for L3 , L2 and  INFRA_VN)

	VLANName string `json:"vlanName,omitempty"` // Vlan name represent the segment name, if empty, vlanName would be auto generated by API

	AutoGenerateVLANName *bool `json:"autoGenerateVlanName,omitempty"` // It will auto generate vlanName, if vlanName is empty(applicable for L3  and INFRA_VN)

	TrafficType string `json:"trafficType,omitempty"` // Traffic type(applicable for L3  and L2)

	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group Name(applicable for L3)

	IsL2FloodingEnabled *bool `json:"isL2FloodingEnabled,omitempty"` // Layer2 flooding enablement flag(applicable for L3 , L2 and always true for L2 and default value is False )

	IsThisCriticalPool *bool `json:"isThisCriticalPool,omitempty"` // Critical pool enablement flag(applicable for L3 and default value is False )

	IsWirelessPool *bool `json:"isWirelessPool,omitempty"` // Wireless Pool enablement flag(applicable for L3  and L2 and default value is False )

	IsIPDirectedBroadcast *bool `json:"isIpDirectedBroadcast,omitempty"` // Ip Directed Broadcast enablement flag(applicable for L3 and default value is False )

	IsCommonPool *bool `json:"isCommonPool,omitempty"` // Common Pool enablement flag(applicable for L3 and L2 and default value is False )

	IsBridgeModeVm *bool `json:"isBridgeModeVm,omitempty"` // Bridge Mode Vm enablement flag (applicable for L3 and L2 and default value is False )

	PoolType string `json:"poolType,omitempty"` // Pool Type (applicable for INFRA_VN)
}
type RequestSdaUpdateAnycastGateways []RequestItemSdaUpdateAnycastGateways // Array of RequestSdaUpdateAnycastGateways
type RequestItemSdaUpdateAnycastGateways struct {
	ID string `json:"id,omitempty"` // ID of the anycast gateway (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this anycast gateway is assigned to. Updating anycast gateways on fabric zones is not allowed--instead, update the corresponding anycast gateway on the fabric site and the updates will be applied on all applicable fabric zones (updating this field is not allowed).

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the anycast gateway (updating this field is not allowed).

	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool associated with the anycast gateway (updating this field is not allowed).

	TCPMssAdjustment *int `json:"tcpMssAdjustment,omitempty"` // TCP maximum segment size adjustment.

	VLANName string `json:"vlanName,omitempty"` // Name of the VLAN of the anycast gateway (updating this field is not allowed).

	VLANID *int `json:"vlanId,omitempty"` // ID of the VLAN of the anycast gateway (updating this field is not allowed).

	TrafficType string `json:"trafficType,omitempty"` // The type of traffic the anycast gateway serves.

	PoolType string `json:"poolType,omitempty"` // The pool type of the anycast gateway (required for & applicable only to INFRA_VN; updating this field is not allowed).

	SecurityGroupName string `json:"securityGroupName,omitempty"` // Name of the associated Security Group (not applicable to INFRA_VN).

	IsCriticalPool *bool `json:"isCriticalPool,omitempty"` // Enable/disable critical VLAN (not applicable to INFRA_VN; updating this field is not allowed).

	IsLayer2FloodingEnabled *bool `json:"isLayer2FloodingEnabled,omitempty"` // Enable/disable layer 2 flooding (not applicable to INFRA_VN).

	IsWirelessPool *bool `json:"isWirelessPool,omitempty"` // Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).

	IsIPDirectedBroadcast *bool `json:"isIpDirectedBroadcast,omitempty"` // Enable/disable IP-directed broadcast (not applicable to INFRA_VN).

	IsIntraSubnetRoutingEnabled *bool `json:"isIntraSubnetRoutingEnabled,omitempty"` // Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN; updating this field is not allowed).

	IsMultipleIPToMacAddresses *bool `json:"isMultipleIpToMacAddresses,omitempty"` // Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).

	IsSupplicantBasedExtendedNodeOnboarding *bool `json:"isSupplicantBasedExtendedNodeOnboarding,omitempty"` // Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN requests; must not be null when poolType is EXTENDED_NODE).

	IsGroupBasedPolicyEnforcementEnabled *bool `json:"isGroupBasedPolicyEnforcementEnabled,omitempty"` // Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN; defaults to false).
}
type RequestSdaAddAnycastGateways []RequestItemSdaAddAnycastGateways // Array of RequestSdaAddAnycastGateways
type RequestItemSdaAddAnycastGateways struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this anycast gateway is to be assigned to.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it.

	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP pool associated with the anycast gateway.

	TCPMssAdjustment *int `json:"tcpMssAdjustment,omitempty"` // TCP maximum segment size adjustment.

	VLANName string `json:"vlanName,omitempty"` // Name of the VLAN of the anycast gateway.

	VLANID *int `json:"vlanId,omitempty"` // ID of the VLAN of the anycast gateway. allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, 2046, and 4094. if deploying an anycast gateway on a fabric zone, this vlanId must match the vlanId of the corresponding anycast gateway on the fabric site.

	TrafficType string `json:"trafficType,omitempty"` // The type of traffic the anycast gateway serves.

	PoolType string `json:"poolType,omitempty"` // The pool type of the anycast gateway (required for & applicable only to INFRA_VN).

	SecurityGroupName string `json:"securityGroupName,omitempty"` // Name of the associated Security Group (not applicable to INFRA_VN).

	IsCriticalPool *bool `json:"isCriticalPool,omitempty"` // Enable/disable critical VLAN. if true, autoGenerateVlanName must also be true. (isCriticalPool is not applicable to INFRA_VN).

	IsLayer2FloodingEnabled *bool `json:"isLayer2FloodingEnabled,omitempty"` // Enable/disable layer 2 flooding (not applicable to INFRA_VN).

	IsWirelessPool *bool `json:"isWirelessPool,omitempty"` // Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).

	IsIPDirectedBroadcast *bool `json:"isIpDirectedBroadcast,omitempty"` // Enable/disable IP-directed broadcast (not applicable to INFRA_VN).

	IsIntraSubnetRoutingEnabled *bool `json:"isIntraSubnetRoutingEnabled,omitempty"` // Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN).

	IsMultipleIPToMacAddresses *bool `json:"isMultipleIpToMacAddresses,omitempty"` // Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).

	IsSupplicantBasedExtendedNodeOnboarding *bool `json:"isSupplicantBasedExtendedNodeOnboarding,omitempty"` // Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN).

	IsGroupBasedPolicyEnforcementEnabled *bool `json:"isGroupBasedPolicyEnforcementEnabled,omitempty"` // Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN; defaults to false).

	AutoGenerateVLANName *bool `json:"autoGenerateVlanName,omitempty"` // This field cannot be true when vlanName is provided. the vlanName will be generated as "{ipPoolGroupV4Cidr}-{virtualNetworkName}" for non-critical VLANs. for critical VLANs with DATA trafficType, vlanName will be "CRITICAL_VLAN". for critical VLANs with VOICE trafficType, vlanName will be "VOICE_VLAN".
}
type RequestSdaUpdateAuthenticationProfile []RequestItemSdaUpdateAuthenticationProfile // Array of RequestSdaUpdateAuthenticationProfile
type RequestItemSdaUpdateAuthenticationProfile struct {
	ID string `json:"id,omitempty"` // ID of the authentication profile (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this authentication profile is assigned to (updating this field is not allowed). To update a global authentication profile, either remove this property or set its value to null.

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // The default host authentication template (updating this field is not allowed).

	AuthenticationOrder string `json:"authenticationOrder,omitempty"` // First authentication method.

	Dot1XToMabFallbackTimeout *int `json:"dot1xToMabFallbackTimeout,omitempty"` // 802.1x Timeout.

	WakeOnLan *bool `json:"wakeOnLan,omitempty"` // Wake on LAN.

	NumberOfHosts string `json:"numberOfHosts,omitempty"` // Number of Hosts.

	IsBpduGuardEnabled *bool `json:"isBpduGuardEnabled,omitempty"` // Enable/disable BPDU Guard. Only applicable when authenticationProfileName is set to "Closed Authentication" (defaults to true).

	PreAuthACL *RequestItemSdaUpdateAuthenticationProfilePreAuthACL `json:"preAuthAcl,omitempty"` //
}
type RequestItemSdaUpdateAuthenticationProfilePreAuthACL struct {
	Enabled *bool `json:"enabled,omitempty"` // Enable/disable Pre-Authentication ACL.

	ImplicitAction string `json:"implicitAction,omitempty"` // Implicit behaviour unless overridden (defaults to "DENY").

	Description string `json:"description,omitempty"` // Description of this Pre-Authentication ACL.

	AccessContracts *[]RequestItemSdaUpdateAuthenticationProfilePreAuthACLAccessContracts `json:"accessContracts,omitempty"` //
}
type RequestItemSdaUpdateAuthenticationProfilePreAuthACLAccessContracts struct {
	Action string `json:"action,omitempty"` // Contract behaviour.

	Protocol string `json:"protocol,omitempty"` // Protocol for the access contract. "TCP" and "TCP_UDP" are only allowed when the contract port is "domain".

	Port string `json:"port,omitempty"` // Port for the access contract. The port can only be used once in the Access Contract list.
}
type RequestSdaUpdateExtranetPolicy []RequestItemSdaUpdateExtranetPolicy // Array of RequestSdaUpdateExtranetPolicy
type RequestItemSdaUpdateExtranetPolicy struct {
	ID string `json:"id,omitempty"` // ID of the existing extranet policy (updating this field is not allowed).

	ExtranetPolicyName string `json:"extranetPolicyName,omitempty"` // Name of the existing extranet policy (updating this field is not allowed).

	FabricIDs []string `json:"fabricIds,omitempty"` // IDs of the fabric sites associated with this extranet policy.

	ProviderVirtualNetworkName string `json:"providerVirtualNetworkName,omitempty"` // Name of the existing provider virtual network (updating this field is not allowed).

	SubscriberVirtualNetworkNames []string `json:"subscriberVirtualNetworkNames,omitempty"` // Name of the subscriber virtual networks.
}
type RequestSdaAddExtranetPolicy []RequestItemSdaAddExtranetPolicy // Array of RequestSdaAddExtranetPolicy
type RequestItemSdaAddExtranetPolicy struct {
	ExtranetPolicyName string `json:"extranetPolicyName,omitempty"` // Name of the extranet policy to be created.

	FabricIDs []string `json:"fabricIds,omitempty"` // IDs of the fabric sites to be associated with this extranet policy.

	ProviderVirtualNetworkName string `json:"providerVirtualNetworkName,omitempty"` // Name of the existing provider virtual network.

	SubscriberVirtualNetworkNames []string `json:"subscriberVirtualNetworkNames,omitempty"` // Name of the subscriber virtual networks.
}
type RequestSdaUpdateFabricDevices []RequestItemSdaUpdateFabricDevices // Array of RequestSdaUpdateFabricDevices
type RequestItemSdaUpdateFabricDevices struct {
	ID string `json:"id,omitempty"` // ID of the fabric device. (updating this field is not allowed).

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device. (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric of this fabric device. (updating this field is not allowed).

	DeviceRoles []string `json:"deviceRoles,omitempty"` // List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE]. (updating this field is not allowed).

	BorderDeviceSettings *RequestItemSdaUpdateFabricDevicesBorderDeviceSettings `json:"borderDeviceSettings,omitempty"` //
}
type RequestItemSdaUpdateFabricDevicesBorderDeviceSettings struct {
	BorderTypes []string `json:"borderTypes,omitempty"` // List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].

	Layer3Settings *RequestItemSdaUpdateFabricDevicesBorderDeviceSettingsLayer3Settings `json:"layer3Settings,omitempty"` //
}
type RequestItemSdaUpdateFabricDevicesBorderDeviceSettingsLayer3Settings struct {
	LocalAutonomousSystemNumber string `json:"localAutonomousSystemNumber,omitempty"` // BGP Local autonomous system number of the fabric border device. Allowed range is [1 to 4294967295]. (updating this field is not allowed).

	IsDefaultExit *bool `json:"isDefaultExit,omitempty"` // Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes. (updating this field is not allowed).

	ImportExternalRoutes *bool `json:"importExternalRoutes,omitempty"` // Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane. (updating this field is not allowed).

	BorderPriority *int `json:"borderPriority,omitempty"` // Border priority of the fabric border device. Allowed range is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.

	PrependAutonomousSystemCount *int `json:"prependAutonomousSystemCount,omitempty"` // Prepend autonomous system count of the fabric border device. Allowed range is [1 to 10].
}
type RequestSdaAddFabricDevices []RequestItemSdaAddFabricDevices // Array of RequestSdaAddFabricDevices
type RequestItemSdaAddFabricDevices struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric of this fabric device.

	DeviceRoles []string `json:"deviceRoles,omitempty"` // List of the roles of the fabric device. Allowed values are [CONTROL_PLANE_NODE, EDGE_NODE, BORDER_NODE, WIRELESS_CONTROLLER_NODE].

	BorderDeviceSettings *RequestItemSdaAddFabricDevicesBorderDeviceSettings `json:"borderDeviceSettings,omitempty"` //
}
type RequestItemSdaAddFabricDevicesBorderDeviceSettings struct {
	BorderTypes []string `json:"borderTypes,omitempty"` // List of the border types of the fabric device. Allowed values are [LAYER_2, LAYER_3].

	Layer3Settings *RequestItemSdaAddFabricDevicesBorderDeviceSettingsLayer3Settings `json:"layer3Settings,omitempty"` //
}
type RequestItemSdaAddFabricDevicesBorderDeviceSettingsLayer3Settings struct {
	LocalAutonomousSystemNumber string `json:"localAutonomousSystemNumber,omitempty"` // BGP Local autonomous system number of the fabric border device. Allowed range is [1 to 4294967295].

	IsDefaultExit *bool `json:"isDefaultExit,omitempty"` // Set this to make the fabric border device the gateway of last resort for this site. Any unknown traffic will be sent to this fabric border device from edge nodes.

	ImportExternalRoutes *bool `json:"importExternalRoutes,omitempty"` // Set this to import external routes from other routing protocols (such as BGP) to the fabric control plane.

	BorderPriority *int `json:"borderPriority,omitempty"` // Border priority of the fabric border device. Allowed range is [1-9]. A lower value indicates higher priority. E.g., a priority of 1 takes precedence over 5. Default priority would be set to 10.

	PrependAutonomousSystemCount *int `json:"prependAutonomousSystemCount,omitempty"` // Prepend autonomous system count of the fabric border device. Allowed range is [1 to 10].
}
type RequestSdaAddFabricDevicesLayer2Handoffs []RequestItemSdaAddFabricDevicesLayer2Handoffs // Array of RequestSdaAddFabricDevicesLayer2Handoffs
type RequestItemSdaAddFabricDevicesLayer2Handoffs struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the layer 2 handoff. E.g., GigabitEthernet1/0/4

	InternalVLANID *int `json:"internalVlanId,omitempty"` // VLAN number associated with this fabric. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).

	ExternalVLANID *int `json:"externalVlanId,omitempty"` // External VLAN number into which the fabric must be extended. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).
}
type RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransit []RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransit // Array of RequestSdaAddFabricDevicesLayer3HandoffsWithIpTransit
type RequestItemSdaAddFabricDevicesLayer3HandoffsWithIPTransit struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to.

	TransitNetworkID string `json:"transitNetworkId,omitempty"` // ID of the transit network of the layer 3 handoff ip transit.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the layer 3 handoff ip transit.

	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with this fabric site.

	VLANID *int `json:"vlanId,omitempty"` // VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network.  Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094).

	TCPMssAdjustment *int `json:"tcpMssAdjustment,omitempty"` // TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.

	LocalIPAddress string `json:"localIpAddress,omitempty"` // Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.

	RemoteIPAddress string `json:"remoteIpAddress,omitempty"` // Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.

	LocalIPv6Address string `json:"localIpv6Address,omitempty"` // Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.

	RemoteIPv6Address string `json:"remoteIpv6Address,omitempty"` // Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). Not applicable if you have already provided an external connectivity ip pool name.
}
type RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit []RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit // Array of RequestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransit
type RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit struct {
	ID string `json:"id,omitempty"` // ID of the fabric device layer 3 handoff ip transit. (updating this field is not allowed).

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device. (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to. (updating this field is not allowed).

	TransitNetworkID string `json:"transitNetworkId,omitempty"` // ID of the transit network of the layer 3 handoff ip transit. (updating this field is not allowed).

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the layer 3 handoff ip transit. (updating this field is not allowed).

	ExternalConnectivityIPPoolName string `json:"externalConnectivityIpPoolName,omitempty"` // External connectivity ip pool will be used by Catalyst Center to allocate IP address for the connection between the border node and peer. (updating this field is not allowed).

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with this fabric site. (updating this field is not allowed).

	VLANID *int `json:"vlanId,omitempty"` // VLAN number for the Switch Virtual Interface (SVI) used to establish BGP peering with the external domain for the virtual network. Allowed VLAN range is 2-4094 except for reserved vlans (1, 1002-1005, 2046, 4094). (updating this field is not allowed).

	TCPMssAdjustment *int `json:"tcpMssAdjustment,omitempty"` // TCP maximum segment size (mss) value for the layer 3 handoff. Allowed range is [500-1440]. TCP MSS Adjustment value is applicable for the TCP sessions over both IPv4 and IPv6.

	LocalIPAddress string `json:"localIpAddress,omitempty"` // Local ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). (updating this field is not allowed).

	RemoteIPAddress string `json:"remoteIpAddress,omitempty"` // Remote ipv4 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). (updating this field is not allowed).

	LocalIPv6Address string `json:"localIpv6Address,omitempty"` // Local ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). If the value has already been set, it cannot be updated.

	RemoteIPv6Address string `json:"remoteIpv6Address,omitempty"` // Remote ipv6 address for the selected virtual network. Enter the IP addresses and subnet mask in the CIDR notation (IP address/prefix-length). If the value has already been set, it cannot be updated.
}
type RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit []RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit // Array of RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit
type RequestItemSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device. (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to. (updating this field is not allowed).

	TransitNetworkID string `json:"transitNetworkId,omitempty"` // ID of the transit network of the layer 3 handoff sda transit. (updating this field is not allowed).

	AffinityIDPrime *int `json:"affinityIdPrime,omitempty"` // Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.

	AffinityIDDecider *int `json:"affinityIdDecider,omitempty"` // Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.

	ConnectedToInternet *bool `json:"connectedToInternet,omitempty"` // Set this true to allow associated site to provide internet access to other sites through sd-access.

	IsMulticastOverTransitEnabled *bool `json:"isMulticastOverTransitEnabled,omitempty"` // Set this true to configure native multicast over multiple sites that are connected to an sd-access transit.
}
type RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransit []RequestItemSdaAddFabricDevicesLayer3HandoffsWithSdaTransit // Array of RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransit
type RequestItemSdaAddFabricDevicesLayer3HandoffsWithSdaTransit struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the fabric device.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this device is assigned to.

	TransitNetworkID string `json:"transitNetworkId,omitempty"` // ID of the transit network of the layer 3 handoff sda transit.

	AffinityIDPrime *int `json:"affinityIdPrime,omitempty"` // Affinity id prime value of the border node. It supersedes the border priority to determine border node preference. Allowed range is [0-2147483647]. The lower the relative value of affinity id prime, the higher the preference for a destination border node.

	AffinityIDDecider *int `json:"affinityIdDecider,omitempty"` // Affinity id decider value of the border node. When the affinity id prime value is the same on multiple devices, the affinity id decider value is used as a tiebreaker. Allowed range is [0-2147483647]. The lower the relative value of affinity id decider, the higher the preference for a destination border node.

	ConnectedToInternet *bool `json:"connectedToInternet,omitempty"` // Set this true to allow associated site to provide internet access to other sites through sd-access.

	IsMulticastOverTransitEnabled *bool `json:"isMulticastOverTransitEnabled,omitempty"` // Set this true to configure native multicast over multiple sites that are connected to an sd-access transit.
}
type RequestSdaAddFabricSite []RequestItemSdaAddFabricSite // Array of RequestSdaAddFabricSite
type RequestItemSdaAddFabricSite struct {
	SiteID string `json:"siteId,omitempty"` // ID of the network hierarchy.

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.

	IsPubSubEnabled *bool `json:"isPubSubEnabled,omitempty"` // Specifies whether this fabric site will use pub/sub for control nodes.
}
type RequestSdaUpdateFabricSite []RequestItemSdaUpdateFabricSite // Array of RequestSdaUpdateFabricSite
type RequestItemSdaUpdateFabricSite struct {
	ID string `json:"id,omitempty"` // ID of the fabric site (updating this field is not allowed).

	SiteID string `json:"siteId,omitempty"` // ID of the network hierarchy (updating this field is not allowed).

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.

	IsPubSubEnabled *bool `json:"isPubSubEnabled,omitempty"` // Specifies whether this fabric site will use pub/sub for control nodes.
}
type RequestSdaUpdateFabricZone []RequestItemSdaUpdateFabricZone // Array of RequestSdaUpdateFabricZone
type RequestItemSdaUpdateFabricZone struct {
	ID string `json:"id,omitempty"` // ID of the fabric zone (updating this field is not allowed).

	SiteID string `json:"siteId,omitempty"` // ID of the network hierarchy (updating this field is not allowed).

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
}
type RequestSdaAddFabricZone []RequestItemSdaAddFabricZone // Array of RequestSdaAddFabricZone
type RequestItemSdaAddFabricZone struct {
	SiteID string `json:"siteId,omitempty"` // ID of the network hierarchy.

	AuthenticationProfileName string `json:"authenticationProfileName,omitempty"` // Authentication profile used for this fabric.
}
type RequestSdaAddLayer2VirtualNetworks []RequestItemSdaAddLayer2VirtualNetworks // Array of RequestSdaAddLayer2VirtualNetworks
type RequestItemSdaAddLayer2VirtualNetworks struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this layer 2 virtual network is to be assigned to.

	VLANName string `json:"vlanName,omitempty"` // Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens.

	VLANID *int `json:"vlanId,omitempty"` // ID of the VLAN of the layer 2 virtual network. Allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, and 2046. If deploying on a fabric zone, this vlanId must match the vlanId of the corresponding layer 2 virtual network on the fabric site.

	TrafficType string `json:"trafficType,omitempty"` // The type of traffic that is served.

	IsFabricEnabledWireless *bool `json:"isFabricEnabledWireless,omitempty"` // Set to true to enable wireless. Default is false.

	IsMultipleIPToMacAddresses *bool `json:"isMultipleIpToMacAddresses,omitempty"` // Set to true to enable multiple IP-to-MAC addresses (Wireless Bridged-Network Virtual Machine). This field defaults to false when associated with a layer 3 virtual network and cannot be used when not associated with a layer 3 virtual network.

	AssociatedLayer3VirtualNetworkName string `json:"associatedLayer3VirtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. The layer 3 virtual network must have already been added to the fabric before association. This field must either be present in all payload elements or none.
}
type RequestSdaUpdateLayer2VirtualNetworks []RequestItemSdaUpdateLayer2VirtualNetworks // Array of RequestSdaUpdateLayer2VirtualNetworks
type RequestItemSdaUpdateLayer2VirtualNetworks struct {
	ID string `json:"id,omitempty"` // ID of the layer 2 virtual network (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric this layer 2 virtual network is assigned to (updating this field is not allowed).

	VLANName string `json:"vlanName,omitempty"` // Name of the VLAN of the layer 2 virtual network. Must contain only alphanumeric characters, underscores, and hyphens (updating this field is not allowed).

	VLANID *int `json:"vlanId,omitempty"` // ID of the VLAN of the layer 2 virtual network (updating this field is not allowed).

	TrafficType string `json:"trafficType,omitempty"` // The type of traffic that is served.

	IsFabricEnabledWireless *bool `json:"isFabricEnabledWireless,omitempty"` // Set to true to enable wireless.

	IsMultipleIPToMacAddresses *bool `json:"isMultipleIpToMacAddresses,omitempty"` // Set to true to enable multiple IP-to-MAC addresses (Wireless Bridged-Network Virtual Machine). This field defaults to false when associated with a layer 3 virtual network and cannot be used when not associated with a layer 3 virtual network.

	AssociatedLayer3VirtualNetworkName string `json:"associatedLayer3VirtualNetworkName,omitempty"` // Name of the layer 3 virtual network associated with the layer 2 virtual network. This field is provided to support requests related to virtual network anchoring. This field must either be present in all payload elements or none (updating this field is not allowed).
}
type RequestSdaAddLayer3VirtualNetworks []RequestItemSdaAddLayer3VirtualNetworks // Array of RequestSdaAddLayer3VirtualNetworks
type RequestItemSdaAddLayer3VirtualNetworks struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network.

	FabricIDs []string `json:"fabricIds,omitempty"` // IDs of the fabrics this layer 3 virtual network is to be assigned to.

	AnchoredSiteID string `json:"anchoredSiteId,omitempty"` // Fabric ID of the fabric site this layer 3 virtual network is to be anchored at.
}
type RequestSdaUpdateLayer3VirtualNetworks []RequestItemSdaUpdateLayer3VirtualNetworks // Array of RequestSdaUpdateLayer3VirtualNetworks
type RequestItemSdaUpdateLayer3VirtualNetworks struct {
	ID string `json:"id,omitempty"` // ID of the layer 3 virtual network (updating this field is not allowed).

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the layer 3 virtual network (updating this field is not allowed).

	FabricIDs []string `json:"fabricIds,omitempty"` // IDs of the fabrics this layer 3 virtual network is assigned to.

	AnchoredSiteID string `json:"anchoredSiteId,omitempty"` // Fabric ID of the fabric site this layer 3 virtual network is anchored at.
}
type RequestSdaUpdateMulticast []RequestItemSdaUpdateMulticast // Array of RequestSdaUpdateMulticast
type RequestItemSdaUpdateMulticast struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric site (updating this field is not allowed).

	ReplicationMode string `json:"replicationMode,omitempty"` // Replication Mode deployed in the fabric site.
}
type RequestSdaAddMulticastVirtualNetworks []RequestItemSdaAddMulticastVirtualNetworks // Array of RequestSdaAddMulticastVirtualNetworks
type RequestItemSdaAddMulticastVirtualNetworks struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric site this multicast configuration is associated with.

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with the fabric site.

	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP Pool associated with the fabric site.

	IPv4SsmRanges []string `json:"ipv4SsmRanges,omitempty"` // IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.

	MulticastRPs *[]RequestItemSdaAddMulticastVirtualNetworksMulticastRPs `json:"multicastRPs,omitempty"` //
}
type RequestItemSdaAddMulticastVirtualNetworksMulticastRPs struct {
	RpDeviceLocation string `json:"rpDeviceLocation,omitempty"` // Device location of the RP.

	IPv4Address string `json:"ipv4Address,omitempty"` // IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.

	IPv6Address string `json:"ipv6Address,omitempty"` // IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.

	IsDefaultV4RP *bool `json:"isDefaultV4RP,omitempty"` // Specifies whether it is a default IPv4 RP.

	IsDefaultV6RP *bool `json:"isDefaultV6RP,omitempty"` // Specifies whether it is a default IPv6 RP.

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.

	IPv4AsmRanges []string `json:"ipv4AsmRanges,omitempty"` // IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.

	IPv6AsmRanges []string `json:"ipv6AsmRanges,omitempty"` // IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
}
type RequestSdaUpdateMulticastVirtualNetworks []RequestItemSdaUpdateMulticastVirtualNetworks // Array of RequestSdaUpdateMulticastVirtualNetworks
type RequestItemSdaUpdateMulticastVirtualNetworks struct {
	ID string `json:"id,omitempty"` // ID of the multicast configuration (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric site this multicast configuration is associated with (updating this field is not allowed).

	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Name of the virtual network associated with the fabric site (updating this field is not allowed).

	IPPoolName string `json:"ipPoolName,omitempty"` // Name of the IP Pool associated with the fabric site (updating this field is not allowed).

	IPv4SsmRanges []string `json:"ipv4SsmRanges,omitempty"` // IPv4 Source Specific Multicast (SSM) ranges. Allowed ranges are from 225.0.0.0/8 to 239.0.0.0/8. SSM ranges should not conflict with ranges provided for ASM multicast.

	MulticastRPs *[]RequestItemSdaUpdateMulticastVirtualNetworksMulticastRPs `json:"multicastRPs,omitempty"` //
}
type RequestItemSdaUpdateMulticastVirtualNetworksMulticastRPs struct {
	RpDeviceLocation string `json:"rpDeviceLocation,omitempty"` // Device location of the RP.

	IPv4Address string `json:"ipv4Address,omitempty"` // IPv4 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests.

	IPv6Address string `json:"ipv6Address,omitempty"` // IPv6 address of the RP. For external RP configuration, exactly one of ipv4Address or ipv6Address must be provided. For fabric RP, this address is allocated by SDA and should not be provided during RP creation request and SDA allocated address should be retained in subsequent requests. ipv6Address can only be provided for virtual networks with dual stack (IPv4 + IPv6) multicast pool.

	IsDefaultV4RP *bool `json:"isDefaultV4RP,omitempty"` // Specifies whether it is a default IPv4 RP.

	IsDefaultV6RP *bool `json:"isDefaultV6RP,omitempty"` // Specifies whether it is a default IPv6 RP.

	NetworkDeviceIDs []string `json:"networkDeviceIds,omitempty"` // IDs of the network devices. This is a required field for fabric RPs. There can be maximum of two fabric RPs for a fabric site and these are shared across all multicast virtual networks. For configuring two fabric RPs in a fabric site both devices must have border roles. Only one RP can be configured in scenarios where a fabric edge device is used as RP or a dual stack multicast pool is used.

	IPv4AsmRanges []string `json:"ipv4AsmRanges,omitempty"` // IPv4 Any Source Multicast ranges. Comma seperated list of IPv4 multicast group ranges that will be served by a given Multicast RP. Only IPv4 ranges can be provided. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.

	IPv6AsmRanges []string `json:"ipv6AsmRanges,omitempty"` // IPv6 Any Source Multicast ranges. Comma seperated list of IPv6 multicast group ranges that will be served by a given Multicast RP. Only IPv6 ranges can be provided. IPv6 ranges can only be provided for dual stack multicast pool. For fabric RP, both IPv4 and IPv6 ranges can be provided together. For external RP, IPv4 ranges should be provided for IPv4 external RP and IPv6 ranges should be provided for IPv6 external RP.
}
type RequestSdaApplyPendingFabricEvents []RequestItemSdaApplyPendingFabricEvents // Array of RequestSdaApplyPendingFabricEvents
type RequestItemSdaApplyPendingFabricEvents struct {
	ID string `json:"id,omitempty"` // ID of the pending fabric event to be applied.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric.
}
type RequestSdaAddPortAssignments []RequestItemSdaAddPortAssignments // Array of RequestSdaAddPortAssignments
type RequestItemSdaAddPortAssignments struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric the device is assigned to.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the port assignment.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the port assignment.

	ConnectedDeviceType string `json:"connectedDeviceType,omitempty"` // Connected device type of the port assignment.

	DataVLANName string `json:"dataVlanName,omitempty"` // Data VLAN name of the port assignment.

	VoiceVLANName string `json:"voiceVlanName,omitempty"` // Voice VLAN name of the port assignment.

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate template name of the port assignment.

	SecurityGroupName string `json:"securityGroupName,omitempty"` // Security group name of the port assignment.

	InterfaceDescription string `json:"interfaceDescription,omitempty"` // Interface description of the port assignment.
}
type RequestSdaUpdatePortAssignments []RequestItemSdaUpdatePortAssignments // Array of RequestSdaUpdatePortAssignments
type RequestItemSdaUpdatePortAssignments struct {
	ID string `json:"id,omitempty"` // ID of the port assignment.

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric the device is assigned to (updating this filed is not allowed).

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the port assignment (updating this field is not allowed).

	InterfaceName string `json:"interfaceName,omitempty"` // Interface name of the port assignment (updating this field is not allowed).

	ConnectedDeviceType string `json:"connectedDeviceType,omitempty"` // Connected device type of the port assignment (updating this field is not allowed).

	DataVLANName string `json:"dataVlanName,omitempty"` // Data VLAN name of the port assignment.

	VoiceVLANName string `json:"voiceVlanName,omitempty"` // Voice VLAN name of the port assignment.

	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate template name of the port assignment.

	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable group name of the port assignment.

	InterfaceDescription string `json:"interfaceDescription,omitempty"` // Interface description of the port assignment.
}
type RequestSdaAddPortChannels []RequestItemSdaAddPortChannels // Array of RequestSdaAddPortChannels
type RequestItemSdaAddPortChannels struct {
	FabricID string `json:"fabricId,omitempty"` // ID of the fabric the device is assigned to.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device.

	InterfaceNames []string `json:"interfaceNames,omitempty"` // Interface names for this port channel (Maximum 16 ports for LACP protocol, Maximum 8 ports for PAGP and ON protocol).

	ConnectedDeviceType string `json:"connectedDeviceType,omitempty"` // Connected device type of the port channel.

	Protocol string `json:"protocol,omitempty"` // Protocol of the port channel (only PAGP is allowed if connectedDeviceType is EXTENDED_NODE).

	Description string `json:"description,omitempty"` // Description of the port channel.
}
type RequestSdaUpdatePortChannels []RequestItemSdaUpdatePortChannels // Array of RequestSdaUpdatePortChannels
type RequestItemSdaUpdatePortChannels struct {
	ID string `json:"id,omitempty"` // ID of the port channel (updating this field is not allowed).

	FabricID string `json:"fabricId,omitempty"` // ID of the fabric the device is assigned to (updating this field is not allowed).

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device (updating this field is not allowed).

	PortChannelName string `json:"portChannelName,omitempty"` // Name of the port channel (updating this field is not allowed).

	InterfaceNames []string `json:"interfaceNames,omitempty"` // Interface names for this port channel (Maximum 16 ports for LACP protocol, Maximum 8 ports for PAGP and ON protocol).

	ConnectedDeviceType string `json:"connectedDeviceType,omitempty"` // Connected device type of the port channel.

	Protocol string `json:"protocol,omitempty"` // Protocol of the port channel (updating this field is not allowed).

	Description string `json:"description,omitempty"` // Description of the port channel.
}
type RequestSdaProvisionDevices []RequestItemSdaProvisionDevices // Array of RequestSdaProvisionDevices
type RequestItemSdaProvisionDevices struct {
	SiteID string `json:"siteId,omitempty"` // ID of the site this network device needs to be provisioned.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of network device to be provisioned.
}
type RequestSdaReProvisionDevices []RequestItemSdaReProvisionDevices // Array of RequestSdaReProvisionDevices
type RequestItemSdaReProvisionDevices struct {
	ID string `json:"id,omitempty"` // ID of the provisioned device.

	SiteID string `json:"siteId,omitempty"` // ID of the site this device is already provisioned to. (updating this field is not allowed).

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // ID of the network device to be re-provisioned. (updating this field is not allowed).
}
type RequestSdaUpdateTransitNetworks []RequestItemSdaUpdateTransitNetworks // Array of RequestSdaUpdateTransitNetworks
type RequestItemSdaUpdateTransitNetworks struct {
	ID string `json:"id,omitempty"` // ID of the transit network (updating this field is not allowed).

	Name string `json:"name,omitempty"` // Name of the transit network (updating this field is not allowed).

	Type string `json:"type,omitempty"` // Type of the transit network (updating this field is not allowed).

	IPTransitSettings *RequestItemSdaUpdateTransitNetworksIPTransitSettings `json:"ipTransitSettings,omitempty"` //

	SdaTransitSettings *RequestItemSdaUpdateTransitNetworksSdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type RequestItemSdaUpdateTransitNetworksIPTransitSettings struct {
	RoutingProtocolName string `json:"routingProtocolName,omitempty"` // Routing Protocol Name of the IP transit network (updating this field is not allowed).

	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous System Number of the IP transit network. Allowed range is [1 to 4294967295] (updating this field is not allowed).
}
type RequestItemSdaUpdateTransitNetworksSdaTransitSettings struct {
	IsMulticastOverTransitEnabled *bool `json:"isMulticastOverTransitEnabled,omitempty"` // Set this to true to enable multicast over SD-Access transit. This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.

	ControlPlaneNetworkDeviceIDs []string `json:"controlPlaneNetworkDeviceIds,omitempty"` // List of network device IDs that will be used as control plane nodes. Maximum 2 network device IDs can be provided for transit of type SDA_LISP_BGP_TRANSIT and maximum 4 network device IDs can be provided for transit of type SDA_LISP_PUB_SUB_TRANSIT.
}
type RequestSdaAddTransitNetworks []RequestItemSdaAddTransitNetworks // Array of RequestSdaAddTransitNetworks
type RequestItemSdaAddTransitNetworks struct {
	Name string `json:"name,omitempty"` // Name of the transit network.

	Type string `json:"type,omitempty"` // Type of the transit network.

	IPTransitSettings *RequestItemSdaAddTransitNetworksIPTransitSettings `json:"ipTransitSettings,omitempty"` //

	SdaTransitSettings *RequestItemSdaAddTransitNetworksSdaTransitSettings `json:"sdaTransitSettings,omitempty"` //
}
type RequestItemSdaAddTransitNetworksIPTransitSettings struct {
	RoutingProtocolName string `json:"routingProtocolName,omitempty"` // Routing protocol name of the IP transit network.

	AutonomousSystemNumber string `json:"autonomousSystemNumber,omitempty"` // Autonomous system number of the IP transit network. Allowed range is [1 to 4294967295].
}
type RequestItemSdaAddTransitNetworksSdaTransitSettings struct {
	IsMulticastOverTransitEnabled *bool `json:"isMulticastOverTransitEnabled,omitempty"` // Set this to true to enable multicast over SD-Access transit.  This supports Native Multicast over SD-Access Transit. This is only applicable for transit of type SDA_LISP_PUB_SUB_TRANSIT.

	ControlPlaneNetworkDeviceIDs []string `json:"controlPlaneNetworkDeviceIds,omitempty"` // List of network device IDs that will be used as control plane nodes. Maximum 2 network device IDs can be provided for transit of type SDA_LISP_BGP_TRANSIT and maximum 4 network device IDs can be provided for transit of type SDA_LISP_PUB_SUB_TRANSIT.
}
type RequestSdaAddVirtualNetworkWithScalableGroups struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name to be assigned at global level

	IsGuestVirtualNetwork *bool `json:"isGuestVirtualNetwork,omitempty"` // Guest Virtual Network enablement flag, default value is False.

	ScalableGroupNames []string `json:"scalableGroupNames,omitempty"` // Scalable Group to be associated to virtual network

	VManageVpnID string `json:"vManageVpnId,omitempty"` // vManage vpn id for SD-WAN
}
type RequestSdaUpdateVirtualNetworkWithScalableGroups struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name to be assigned global level

	IsGuestVirtualNetwork *bool `json:"isGuestVirtualNetwork,omitempty"` // Indicates whether to set this as guest virtual network or not, default value is False.

	ScalableGroupNames []string `json:"scalableGroupNames,omitempty"` // Scalable Group Name to be associated to virtual network

	VManageVpnID string `json:"vManageVpnId,omitempty"` // vManage vpn id for SD-WAN
}

//ReadListOfFabricSitesWithTheirHealthSummary Read list of Fabric Sites with their health summary - 1c9a-8a15-4848-aa79
/* Get a paginated list of Fabric sites Networks with health summary.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-fabricSiteHealthSummaries-1.0.1-resolved.yaml


@param ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams Custom header parameters
@param ReadListOfFabricSitesWithTheirHealthSummaryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-list-of-fabric-sites-with-their-health-summary
*/
func (s *SdaService) ReadListOfFabricSitesWithTheirHealthSummary(ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams *ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams, ReadListOfFabricSitesWithTheirHealthSummaryQueryParams *ReadListOfFabricSitesWithTheirHealthSummaryQueryParams) (*ResponseSdaReadListOfFabricSitesWithTheirHealthSummary, *resty.Response, error) {
	path := "/dna/data/api/v1/fabricSiteHealthSummaries"

	queryString, _ := query.Values(ReadListOfFabricSitesWithTheirHealthSummaryQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams != nil {

		if ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadListOfFabricSitesWithTheirHealthSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadListOfFabricSitesWithTheirHealthSummary(ReadListOfFabricSitesWithTheirHealthSummaryHeaderParams, ReadListOfFabricSitesWithTheirHealthSummaryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadListOfFabricSitesWithTheirHealthSummary")
	}

	result := response.Result().(*ResponseSdaReadListOfFabricSitesWithTheirHealthSummary)
	return result, response, err

}

//ReadFabricSiteCount Read fabric site count - b99f-5a81-44fb-b518
/* Get a count of Fabric sites. Use available query parameters to get the count of a subset of fabric sites.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-fabricSiteHealthSummaries-1.0.1-resolved.yaml


@param ReadFabricSiteCountHeaderParams Custom header parameters
@param ReadFabricSiteCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-fabric-site-count
*/
func (s *SdaService) ReadFabricSiteCount(ReadFabricSiteCountHeaderParams *ReadFabricSiteCountHeaderParams, ReadFabricSiteCountQueryParams *ReadFabricSiteCountQueryParams) (*ResponseSdaReadFabricSiteCount, *resty.Response, error) {
	path := "/dna/data/api/v1/fabricSiteHealthSummaries/count"

	queryString, _ := query.Values(ReadFabricSiteCountQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadFabricSiteCountHeaderParams != nil {

		if ReadFabricSiteCountHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadFabricSiteCountHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadFabricSiteCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadFabricSiteCount(ReadFabricSiteCountHeaderParams, ReadFabricSiteCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadFabricSiteCount")
	}

	result := response.Result().(*ResponseSdaReadFabricSiteCount)
	return result, response, err

}

//ReadFabricSitesWithHealthSummaryFromID Read Fabric Sites with health summary from id - f196-7bdb-4fcb-8bf5
/* Get Fabric site health summary for a specific fabric site by providing the unique fabric site id in the url path.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-fabricSiteHealthSummaries-1.0.1-resolved.yaml


@param id id path parameter. unique fabric site id

@param ReadFabricSitesWithHealthSummaryFromIdHeaderParams Custom header parameters
@param ReadFabricSitesWithHealthSummaryFromIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-fabric-sites-with-health-summary-from-id
*/
func (s *SdaService) ReadFabricSitesWithHealthSummaryFromID(id string, ReadFabricSitesWithHealthSummaryFromIdHeaderParams *ReadFabricSitesWithHealthSummaryFromIDHeaderParams, ReadFabricSitesWithHealthSummaryFromIdQueryParams *ReadFabricSitesWithHealthSummaryFromIDQueryParams) (*ResponseSdaReadFabricSitesWithHealthSummaryFromID, *resty.Response, error) {
	path := "/dna/data/api/v1/fabricSiteHealthSummaries/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(ReadFabricSitesWithHealthSummaryFromIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadFabricSitesWithHealthSummaryFromIdHeaderParams != nil {

		if ReadFabricSitesWithHealthSummaryFromIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadFabricSitesWithHealthSummaryFromIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadFabricSitesWithHealthSummaryFromID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadFabricSitesWithHealthSummaryFromID(id, ReadFabricSitesWithHealthSummaryFromIdHeaderParams, ReadFabricSitesWithHealthSummaryFromIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadFabricSitesWithHealthSummaryFromId")
	}

	result := response.Result().(*ResponseSdaReadFabricSitesWithHealthSummaryFromID)
	return result, response, err

}

//TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange The Trend analytics data for a fabric site in the specified time range - 798b-bbc6-4839-85ff
/* Get health time series for a specific Fabric Site by providing the unique Fabric site id in the url path. The data will be grouped based on the specified trend time interval. If startTime and endTime are not provided, the API defaults to the last 24 hours.
By default: the number of records returned will be 500. the records will be sorted in time ascending (`asc`) order
ex: id:93a25378-7740-4e20-8d90-0060ad9a1be0
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-fabricSiteHealthSummaries-1.0.1-resolved.yaml


@param id id path parameter. unique fabric site id

@param TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams Custom header parameters
@param TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-a-fabric-site-in-the-specified-time-range
*/
func (s *SdaService) TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange(id string, TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams *TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams *TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams) (*ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/fabricSiteHealthSummaries/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams != nil {

		if TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange(id, TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRangeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseSdaTheTrendAnalyticsDataForAFabricSiteInTheSpecifiedTimeRange)
	return result, response, err

}

//ReadFabricEntitySummary Read Fabric entity summary - c9b6-8967-4feb-8337
/* Read Fabric summary for overall deployment. Get an aggregated summary of all fabric entities in a deployment including the entity health.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-fabricSummary-1.0.1-oas3-resolved.yaml


@param ReadFabricEntitySummaryHeaderParams Custom header parameters
@param ReadFabricEntitySummaryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-fabric-entity-summary
*/
func (s *SdaService) ReadFabricEntitySummary(ReadFabricEntitySummaryHeaderParams *ReadFabricEntitySummaryHeaderParams, ReadFabricEntitySummaryQueryParams *ReadFabricEntitySummaryQueryParams) (*ResponseSdaReadFabricEntitySummary, *resty.Response, error) {
	path := "/dna/data/api/v1/fabricSummary"

	queryString, _ := query.Values(ReadFabricEntitySummaryQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadFabricEntitySummaryHeaderParams != nil {

		if ReadFabricEntitySummaryHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadFabricEntitySummaryHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadFabricEntitySummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadFabricEntitySummary(ReadFabricEntitySummaryHeaderParams, ReadFabricEntitySummaryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadFabricEntitySummary")
	}

	result := response.Result().(*ResponseSdaReadFabricEntitySummary)
	return result, response, err

}

//ReadListOfTransitNetworksWithTheirHealthSummary Read list of Transit Networks with their health summary - 4b91-1aae-43f8-9ba8
/* Get a paginated list of Transit Networks with health summary.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-transitNetworkHealthSummaries-1.0.1-resolved.yaml


@param ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams Custom header parameters
@param ReadListOfTransitNetworksWithTheirHealthSummaryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-list-of-transit-networks-with-their-health-summary
*/
func (s *SdaService) ReadListOfTransitNetworksWithTheirHealthSummary(ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams *ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams, ReadListOfTransitNetworksWithTheirHealthSummaryQueryParams *ReadListOfTransitNetworksWithTheirHealthSummaryQueryParams) (*ResponseSdaReadListOfTransitNetworksWithTheirHealthSummary, *resty.Response, error) {
	path := "/dna/data/api/v1/transitNetworkHealthSummaries"

	queryString, _ := query.Values(ReadListOfTransitNetworksWithTheirHealthSummaryQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams != nil {

		if ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadListOfTransitNetworksWithTheirHealthSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadListOfTransitNetworksWithTheirHealthSummary(ReadListOfTransitNetworksWithTheirHealthSummaryHeaderParams, ReadListOfTransitNetworksWithTheirHealthSummaryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadListOfTransitNetworksWithTheirHealthSummary")
	}

	result := response.Result().(*ResponseSdaReadListOfTransitNetworksWithTheirHealthSummary)
	return result, response, err

}

//ReadTransitNetworksCount Read Transit Networks count - 63ac-f8e3-4a28-aee6
/* Get a count of transit networks. Use available query parameters to get the count of a subset of transit networks.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-transitNetworkHealthSummaries-1.0.1-resolved.yaml


@param ReadTransitNetworksCountHeaderParams Custom header parameters
@param ReadTransitNetworksCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-transit-networks-count
*/
func (s *SdaService) ReadTransitNetworksCount(ReadTransitNetworksCountHeaderParams *ReadTransitNetworksCountHeaderParams, ReadTransitNetworksCountQueryParams *ReadTransitNetworksCountQueryParams) (*ResponseSdaReadTransitNetworksCount, *resty.Response, error) {
	path := "/dna/data/api/v1/transitNetworkHealthSummaries/count"

	queryString, _ := query.Values(ReadTransitNetworksCountQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadTransitNetworksCountHeaderParams != nil {

		if ReadTransitNetworksCountHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadTransitNetworksCountHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadTransitNetworksCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadTransitNetworksCount(ReadTransitNetworksCountHeaderParams, ReadTransitNetworksCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadTransitNetworksCount")
	}

	result := response.Result().(*ResponseSdaReadTransitNetworksCount)
	return result, response, err

}

//ReadTransitNetworkWithItsHealthSummaryFromID Read transit network with its health summary from id - 278c-0877-4ba9-919b
/* Get health summary for a specific transit Network by providing the unique transit networks id in the url path.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-transitNetworkHealthSummaries-1.0.1-resolved.yaml


@param id id path parameter. The unique transit network id, Ex “1551156a-bc97-3c63-aeda-8a6d3765b5b9"

@param ReadTransitNetworkWithItsHealthSummaryFromIdHeaderParams Custom header parameters
@param ReadTransitNetworkWithItsHealthSummaryFromIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-transit-network-with-its-health-summary-from-id
*/
func (s *SdaService) ReadTransitNetworkWithItsHealthSummaryFromID(id string, ReadTransitNetworkWithItsHealthSummaryFromIdHeaderParams *ReadTransitNetworkWithItsHealthSummaryFromIDHeaderParams, ReadTransitNetworkWithItsHealthSummaryFromIdQueryParams *ReadTransitNetworkWithItsHealthSummaryFromIDQueryParams) (*ResponseSdaReadTransitNetworkWithItsHealthSummaryFromID, *resty.Response, error) {
	path := "/dna/data/api/v1/transitNetworkHealthSummaries/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(ReadTransitNetworkWithItsHealthSummaryFromIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadTransitNetworkWithItsHealthSummaryFromIdHeaderParams != nil {

		if ReadTransitNetworkWithItsHealthSummaryFromIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadTransitNetworkWithItsHealthSummaryFromIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadTransitNetworkWithItsHealthSummaryFromID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadTransitNetworkWithItsHealthSummaryFromID(id, ReadTransitNetworkWithItsHealthSummaryFromIdHeaderParams, ReadTransitNetworkWithItsHealthSummaryFromIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadTransitNetworkWithItsHealthSummaryFromId")
	}

	result := response.Result().(*ResponseSdaReadTransitNetworkWithItsHealthSummaryFromID)
	return result, response, err

}

//TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange The Trend analytics data for a transit network in the specified time range - 10a2-0bed-4cda-8fd3
/* Get health time series for a specific Transit Network by providing the unique Transit Network id in the url path. The data will be grouped based on the specified trend time interval. If startTime and endTime are not provided, the API defaults to the last 24 hours.
By default: the number of records returned will be 500. the records will be sorted in time ascending (`asc`) order
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-transitNetworkHealthSummaries-1.0.1-resolved.yaml


@param id id path parameter. The unique transit network id, Ex “1551156a-bc97-3c63-aeda-8a6d3765b5b9"

@param TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams Custom header parameters
@param TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-a-transit-network-in-the-specified-time-range
*/
func (s *SdaService) TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange(id string, TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams *TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeQueryParams *TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeQueryParams) (*ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/transitNetworkHealthSummaries/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams != nil {

		if TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange(id, TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRangeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseSdaTheTrendAnalyticsDataForATransitNetworkInTheSpecifiedTimeRange)
	return result, response, err

}

//ReadListOfVirtualNetworksWithTheirHealthSummary Read list of Virtual Networks with their health summary - 95a0-580c-422a-bd8b
/* Get a paginated list of Virtual Networks with health summary. Layer 2 Virtual Networks are only included in health reporting for EVPN protocol deployments. The special Layer 3 VN called ‘INFRA_VN’ is also not included for user access through Assurance virtualNetworkHealthSummaries APIS. Please find INFRA_VN related health metrics under /data/api/v1/fabricSiteHealthSummaries (Ex: attributes ‘pubsubInfraVnGoodHealthPercentage’ and ‘bgpPeerInfraVnScoreGoodHealthPercentage’).
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-virtualNetworkHealthSummaries-1.0.1-resolved.yaml


@param ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams Custom header parameters
@param ReadListOfVirtualNetworksWithTheirHealthSummaryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-list-of-virtual-networks-with-their-health-summary
*/
func (s *SdaService) ReadListOfVirtualNetworksWithTheirHealthSummary(ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams *ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams, ReadListOfVirtualNetworksWithTheirHealthSummaryQueryParams *ReadListOfVirtualNetworksWithTheirHealthSummaryQueryParams) (*ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummary, *resty.Response, error) {
	path := "/dna/data/api/v1/virtualNetworkHealthSummaries"

	queryString, _ := query.Values(ReadListOfVirtualNetworksWithTheirHealthSummaryQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams != nil {

		if ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadListOfVirtualNetworksWithTheirHealthSummary(ReadListOfVirtualNetworksWithTheirHealthSummaryHeaderParams, ReadListOfVirtualNetworksWithTheirHealthSummaryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadListOfVirtualNetworksWithTheirHealthSummary")
	}

	result := response.Result().(*ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummary)
	return result, response, err

}

//ReadVirtualNetworksCount Read Virtual Networks count - a892-d9f5-4669-8725
/* Get a count of virtual networks. Use available query parameters to get the count of a subset of virtual networks. Layer 2 Virtual Networks are only included for EVPN protocol deployments. The special Layer 3 VN called ‘INFRA_VN’ is also not included.
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-virtualNetworkHealthSummaries-1.0.1-resolved.yaml


@param ReadVirtualNetworksCountHeaderParams Custom header parameters
@param ReadVirtualNetworksCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-virtual-networks-count
*/
func (s *SdaService) ReadVirtualNetworksCount(ReadVirtualNetworksCountHeaderParams *ReadVirtualNetworksCountHeaderParams, ReadVirtualNetworksCountQueryParams *ReadVirtualNetworksCountQueryParams) (*ResponseSdaReadVirtualNetworksCount, *resty.Response, error) {
	path := "/dna/data/api/v1/virtualNetworkHealthSummaries/count"

	queryString, _ := query.Values(ReadVirtualNetworksCountQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadVirtualNetworksCountHeaderParams != nil {

		if ReadVirtualNetworksCountHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadVirtualNetworksCountHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadVirtualNetworksCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadVirtualNetworksCount(ReadVirtualNetworksCountHeaderParams, ReadVirtualNetworksCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadVirtualNetworksCount")
	}

	result := response.Result().(*ResponseSdaReadVirtualNetworksCount)
	return result, response, err

}

//ReadVirtualNetworkWithItsHealthSummaryFromID Read virtual network with its health summary from id - 3883-1b45-4be9-9fef
/* Get health summary for a specific Virtual Network by providing the unique virtual networks id in the url path. L2 Virtual Networks are only included in health reporting for EVPN protocol deployments. The special Layer 3 VN called ‘INFRA_VN’ is also not included for user access through Assurance virtualNetworkHealthSummaries APIS. Please find INFRA_VN related health metrics under /data/api/v1/fabricSiteHealthSummaries (Ex: attributes ‘pubsubInfraVnGoodHealthPercentage’ and ‘bgpPeerInfraVnScoreGoodHealthPercentage’).
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-virtualNetworkHealthSummaries-1.0.1-resolved.yaml


@param id id path parameter. unique virtual networks id

@param ReadVirtualNetworkWithItsHealthSummaryFromIdHeaderParams Custom header parameters
@param ReadVirtualNetworkWithItsHealthSummaryFromIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!read-virtual-network-with-its-health-summary-from-id
*/
func (s *SdaService) ReadVirtualNetworkWithItsHealthSummaryFromID(id string, ReadVirtualNetworkWithItsHealthSummaryFromIdHeaderParams *ReadVirtualNetworkWithItsHealthSummaryFromIDHeaderParams, ReadVirtualNetworkWithItsHealthSummaryFromIdQueryParams *ReadVirtualNetworkWithItsHealthSummaryFromIDQueryParams) (*ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromID, *resty.Response, error) {
	path := "/dna/data/api/v1/virtualNetworkHealthSummaries/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(ReadVirtualNetworkWithItsHealthSummaryFromIdQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ReadVirtualNetworkWithItsHealthSummaryFromIdHeaderParams != nil {

		if ReadVirtualNetworkWithItsHealthSummaryFromIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ReadVirtualNetworkWithItsHealthSummaryFromIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReadVirtualNetworkWithItsHealthSummaryFromID(id, ReadVirtualNetworkWithItsHealthSummaryFromIdHeaderParams, ReadVirtualNetworkWithItsHealthSummaryFromIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReadVirtualNetworkWithItsHealthSummaryFromId")
	}

	result := response.Result().(*ResponseSdaReadVirtualNetworkWithItsHealthSummaryFromID)
	return result, response, err

}

//TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange The Trend analytics data for a virtual network in the specified time range - b0a9-69dd-4eda-bb56
/* Get health time series for a specific Virtual Network by providing the unique Virtual Network id in the url path. Layer 2 Virtual Networks are only included in health reporting for EVPN protocol deployments. The special Layer 3 VN called ‘INFRA_VN’ is also not included for user access through Assurance virtualNetworkHealthSummaries APIS.
The data will be grouped based on the specified trend time interval. If startTime and endTime are not provided, the API defaults to the last 24 hours.
By default: the number of records returned will be 500. the records will be sorted in time ascending (`asc`) order
For EVPN , {id} is a combination of VN:FabrisiteId. ex: L2VN1:93a25378-7740-4e20-8d90-0060ad9a1be0
This API provides the latest health data until the given `endTime`. If data is not ready for the provided endTime, the request will fail with error code `400 Bad Request`, and the error message will indicate the recommended endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we are not a real time system. When `endTime` is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-virtualNetworkHealthSummaries-1.0.1-resolved.yaml


@param id id path parameter. unique virtual network id

@param TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams Custom header parameters
@param TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!the-trend-analytics-data-for-a-virtual-network-in-the-specified-time-range
*/
func (s *SdaService) TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange(id string, TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams *TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeQueryParams *TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeQueryParams) (*ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange, *resty.Response, error) {
	path := "/dna/data/api/v1/virtualNetworkHealthSummaries/{id}/trendAnalytics"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams != nil {

		if TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange(id, TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeHeaderParams, TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRangeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation TheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange")
	}

	result := response.Result().(*ResponseSdaTheTrendAnalyticsDataForAVirtualNetworkInTheSpecifiedTimeRange)
	return result, response, err

}

//GetDefaultAuthenticationProfileFromSdaFabric Get default authentication profile from SDA Fabric - 8b90-8a4e-4c5a-9a23
/* Get default authentication profile from SDA Fabric


@param GetDefaultAuthenticationProfileFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-default-authentication-profile-from-sda-fabric
*/
func (s *SdaService) GetDefaultAuthenticationProfileFromSdaFabric(GetDefaultAuthenticationProfileFromSDAFabricQueryParams *GetDefaultAuthenticationProfileFromSdaFabricQueryParams) (*ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	queryString, _ := query.Values(GetDefaultAuthenticationProfileFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDefaultAuthenticationProfileFromSdaFabric(GetDefaultAuthenticationProfileFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDefaultAuthenticationProfileFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric)
	return result, response, err

}

//GetBorderDeviceDetailFromSdaFabric Get border device detail from SDA Fabric - 98a3-9bf4-485a-9871
/* Get border device detail from SDA Fabric


@param GetBorderDeviceDetailFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-border-device-detail-from-sda-fabric
*/
func (s *SdaService) GetBorderDeviceDetailFromSdaFabric(GetBorderDeviceDetailFromSDAFabricQueryParams *GetBorderDeviceDetailFromSdaFabricQueryParams) (*ResponseSdaGetBorderDeviceDetailFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(GetBorderDeviceDetailFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetBorderDeviceDetailFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBorderDeviceDetailFromSdaFabric(GetBorderDeviceDetailFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetBorderDeviceDetailFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetBorderDeviceDetailFromSdaFabric)
	return result, response, err

}

//GetControlPlaneDeviceFromSdaFabric Get control plane device from SDA Fabric - aba4-991d-4e9b-8747
/* Get control plane device from SDA Fabric


@param GetControlPlaneDeviceFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-control-plane-device-from-sda-fabric
*/
func (s *SdaService) GetControlPlaneDeviceFromSdaFabric(GetControlPlaneDeviceFromSDAFabricQueryParams *GetControlPlaneDeviceFromSdaFabricQueryParams) (*ResponseSdaGetControlPlaneDeviceFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	queryString, _ := query.Values(GetControlPlaneDeviceFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetControlPlaneDeviceFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetControlPlaneDeviceFromSdaFabric(GetControlPlaneDeviceFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetControlPlaneDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetControlPlaneDeviceFromSdaFabric)
	return result, response, err

}

//GetDeviceInfoFromSdaFabric Get device info from SDA Fabric - 1385-18e1-4069-ab5f
/* Get device info from SDA Fabric


@param GetDeviceInfoFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-info-from-sda-fabric
*/
func (s *SdaService) GetDeviceInfoFromSdaFabric(GetDeviceInfoFromSDAFabricQueryParams *GetDeviceInfoFromSdaFabricQueryParams) (*ResponseSdaGetDeviceInfoFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/device"

	queryString, _ := query.Values(GetDeviceInfoFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetDeviceInfoFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceInfoFromSdaFabric(GetDeviceInfoFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceInfoFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetDeviceInfoFromSdaFabric)
	return result, response, err

}

//GetDeviceRoleInSdaFabric Get device role in SDA Fabric - 8a92-d87c-416a-8e83
/* Get device role in SDA Fabric


@param GetDeviceRoleInSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-role-in-sda-fabric
*/
func (s *SdaService) GetDeviceRoleInSdaFabric(GetDeviceRoleInSDAFabricQueryParams *GetDeviceRoleInSdaFabricQueryParams) (*ResponseSdaGetDeviceRoleInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/device/role"

	queryString, _ := query.Values(GetDeviceRoleInSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetDeviceRoleInSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceRoleInSdaFabric(GetDeviceRoleInSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceRoleInSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetDeviceRoleInSdaFabric)
	return result, response, err

}

//GetEdgeDeviceFromSdaFabric Get edge device from SDA Fabric - 7683-f90b-4efa-b090
/* Get edge device from SDA Fabric


@param GetEdgeDeviceFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-edge-device-from-sda-fabric
*/
func (s *SdaService) GetEdgeDeviceFromSdaFabric(GetEdgeDeviceFromSDAFabricQueryParams *GetEdgeDeviceFromSdaFabricQueryParams) (*ResponseSdaGetEdgeDeviceFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/edge-device"

	queryString, _ := query.Values(GetEdgeDeviceFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetEdgeDeviceFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEdgeDeviceFromSdaFabric(GetEdgeDeviceFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEdgeDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetEdgeDeviceFromSdaFabric)
	return result, response, err

}

//GetSiteFromSdaFabric Get Site from SDA Fabric - 80b7-f8e6-406a-8701
/* Get Site info from SDA Fabric


@param GetSiteFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-site-from-sda-fabric
*/
func (s *SdaService) GetSiteFromSdaFabric(GetSiteFromSDAFabricQueryParams *GetSiteFromSdaFabricQueryParams) (*ResponseSdaGetSiteFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric-site"

	queryString, _ := query.Values(GetSiteFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetSiteFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSiteFromSdaFabric(GetSiteFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSiteFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetSiteFromSdaFabric)
	return result, response, err

}

//GetPortAssignmentForAccessPointInSdaFabric Get Port assignment for access point in SDA Fabric - 5097-f8d4-45f9-8f51
/* Get Port assignment for access point in SDA Fabric


@param GetPortAssignmentForAccessPointInSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignment-for-access-point-in-sda-fabric
*/
func (s *SdaService) GetPortAssignmentForAccessPointInSdaFabric(GetPortAssignmentForAccessPointInSDAFabricQueryParams *GetPortAssignmentForAccessPointInSdaFabricQueryParams) (*ResponseSdaGetPortAssignmentForAccessPointInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	queryString, _ := query.Values(GetPortAssignmentForAccessPointInSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentForAccessPointInSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentForAccessPointInSdaFabric(GetPortAssignmentForAccessPointInSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentForAccessPointInSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentForAccessPointInSdaFabric)
	return result, response, err

}

//GetPortAssignmentForUserDeviceInSdaFabric Get Port assignment for user device in SDA Fabric - a4a1-e8ed-41cb-9653
/* Get Port assignment for user device in SDA Fabric.


@param GetPortAssignmentForUserDeviceInSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignment-for-user-device-in-sda-fabric
*/
func (s *SdaService) GetPortAssignmentForUserDeviceInSdaFabric(GetPortAssignmentForUserDeviceInSDAFabricQueryParams *GetPortAssignmentForUserDeviceInSdaFabricQueryParams) (*ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	queryString, _ := query.Values(GetPortAssignmentForUserDeviceInSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentForUserDeviceInSdaFabric(GetPortAssignmentForUserDeviceInSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentForUserDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric)
	return result, response, err

}

//GetMulticastDetailsFromSdaFabric Get multicast details from SDA fabric - c286-f98b-47ba-9ab4
/* Get multicast details from SDA fabric


@param GetMulticastDetailsFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-details-from-sda-fabric
*/
func (s *SdaService) GetMulticastDetailsFromSdaFabric(GetMulticastDetailsFromSDAFabricQueryParams *GetMulticastDetailsFromSdaFabricQueryParams) (*ResponseSdaGetMulticastDetailsFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/multicast"

	queryString, _ := query.Values(GetMulticastDetailsFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastDetailsFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastDetailsFromSdaFabric(GetMulticastDetailsFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastDetailsFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetMulticastDetailsFromSdaFabric)
	return result, response, err

}

//GetProvisionedWiredDevice Get Provisioned Wired Device - dfbf-2ae2-42ca-a449
/* Get Provisioned Wired Device


@param GetProvisionedWiredDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioned-wired-device
*/
func (s *SdaService) GetProvisionedWiredDevice(GetProvisionedWiredDeviceQueryParams *GetProvisionedWiredDeviceQueryParams) (*ResponseSdaGetProvisionedWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/provision-device"

	queryString, _ := query.Values(GetProvisionedWiredDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetProvisionedWiredDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisionedWiredDevice(GetProvisionedWiredDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProvisionedWiredDevice")
	}

	result := response.Result().(*ResponseSdaGetProvisionedWiredDevice)
	return result, response, err

}

//GetTransitPeerNetworkInfo Get Transit Peer Network Info - 16a1-bb5d-48cb-873d
/* Get Transit Peer Network Info from SD-Access


@param GetTransitPeerNetworkInfoQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-transit-peer-network-info
*/
func (s *SdaService) GetTransitPeerNetworkInfo(GetTransitPeerNetworkInfoQueryParams *GetTransitPeerNetworkInfoQueryParams) (*ResponseSdaGetTransitPeerNetworkInfo, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/transit-peer-network"

	queryString, _ := query.Values(GetTransitPeerNetworkInfoQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetTransitPeerNetworkInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTransitPeerNetworkInfo(GetTransitPeerNetworkInfoQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTransitPeerNetworkInfo")
	}

	result := response.Result().(*ResponseSdaGetTransitPeerNetworkInfo)
	return result, response, err

}

//GetVnFromSdaFabric Get VN from SDA Fabric - 2eb1-fa1e-49ca-a2b4
/* Get virtual network (VN) from SDA Fabric


@param GetVNFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-vn-from-sda-fabric
*/
func (s *SdaService) GetVnFromSdaFabric(GetVNFromSDAFabricQueryParams *GetVnFromSdaFabricQueryParams) (*ResponseSdaGetVnFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	queryString, _ := query.Values(GetVNFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetVnFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVnFromSdaFabric(GetVNFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetVnFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetVnFromSdaFabric)
	return result, response, err

}

//GetVirtualNetworkSummary Get Virtual Network Summary - 6fa0-f8d5-4d29-857a
/* Get Virtual Network Summary


@param GetVirtualNetworkSummaryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-virtual-network-summary
*/
func (s *SdaService) GetVirtualNetworkSummary(GetVirtualNetworkSummaryQueryParams *GetVirtualNetworkSummaryQueryParams) (*ResponseSdaGetVirtualNetworkSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network/summary"

	queryString, _ := query.Values(GetVirtualNetworkSummaryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetVirtualNetworkSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVirtualNetworkSummary(GetVirtualNetworkSummaryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworkSummary")
	}

	result := response.Result().(*ResponseSdaGetVirtualNetworkSummary)
	return result, response, err

}

//GetIPPoolFromSdaVirtualNetwork Get IP Pool from SDA Virtual Network - fa92-19bf-45c8-b43b
/* Get IP Pool from SDA Virtual Network


@param GetIPPoolFromSDAVirtualNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ip-pool-from-sda-virtual-network
*/
func (s *SdaService) GetIPPoolFromSdaVirtualNetwork(GetIPPoolFromSDAVirtualNetworkQueryParams *GetIPPoolFromSdaVirtualNetworkQueryParams) (*ResponseSdaGetIPPoolFromSdaVirtualNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	queryString, _ := query.Values(GetIPPoolFromSDAVirtualNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetIPPoolFromSdaVirtualNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetIPPoolFromSdaVirtualNetwork(GetIPPoolFromSDAVirtualNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetIpPoolFromSdaVirtualNetwork")
	}

	result := response.Result().(*ResponseSdaGetIPPoolFromSdaVirtualNetwork)
	return result, response, err

}

//GetAnycastGateways Get anycast gateways - 5cb3-f980-670e-770a
/* Returns a list of anycast gateways that match the provided query parameters.


@param GetAnycastGatewaysQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anycast-gateways
*/
func (s *SdaService) GetAnycastGateways(GetAnycastGatewaysQueryParams *GetAnycastGatewaysQueryParams) (*ResponseSdaGetAnycastGateways, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways"

	queryString, _ := query.Values(GetAnycastGatewaysQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetAnycastGateways{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnycastGateways(GetAnycastGatewaysQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAnycastGateways")
	}

	result := response.Result().(*ResponseSdaGetAnycastGateways)
	return result, response, err

}

//GetAnycastGatewayCount Get anycast gateway count - e504-152d-3f53-4d07
/* Returns the count of anycast gateways that match the provided query parameters.


@param GetAnycastGatewayCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anycast-gateway-count
*/
func (s *SdaService) GetAnycastGatewayCount(GetAnycastGatewayCountQueryParams *GetAnycastGatewayCountQueryParams) (*ResponseSdaGetAnycastGatewayCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways/count"

	queryString, _ := query.Values(GetAnycastGatewayCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetAnycastGatewayCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnycastGatewayCount(GetAnycastGatewayCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAnycastGatewayCount")
	}

	result := response.Result().(*ResponseSdaGetAnycastGatewayCount)
	return result, response, err

}

//GetAuthenticationProfiles Get authentication profiles - 9eb7-1a2d-44c8-82aa
/* Returns a list of authentication profiles that match the provided query parameters.


@param GetAuthenticationProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-authentication-profiles
*/
func (s *SdaService) GetAuthenticationProfiles(GetAuthenticationProfilesQueryParams *GetAuthenticationProfilesQueryParams) (*ResponseSdaGetAuthenticationProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/authenticationProfiles"

	queryString, _ := query.Values(GetAuthenticationProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetAuthenticationProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAuthenticationProfiles(GetAuthenticationProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAuthenticationProfiles")
	}

	result := response.Result().(*ResponseSdaGetAuthenticationProfiles)
	return result, response, err

}

//GetExtranetPolicies Get extranet policies - 3f85-3834-4b1b-bbcb
/* Returns a list of extranet policies that match the provided query parameters.


@param GetExtranetPoliciesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-extranet-policies
*/
func (s *SdaService) GetExtranetPolicies(GetExtranetPoliciesQueryParams *GetExtranetPoliciesQueryParams) (*ResponseSdaGetExtranetPolicies, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	queryString, _ := query.Values(GetExtranetPoliciesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetExtranetPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExtranetPolicies(GetExtranetPoliciesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetExtranetPolicies")
	}

	result := response.Result().(*ResponseSdaGetExtranetPolicies)
	return result, response, err

}

//GetExtranetPolicyCount Get extranet policy count - 35a7-4975-447a-a6b8
/* Returns the count of extranet policies that match the provided query parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-extranet-policy-count
*/
func (s *SdaService) GetExtranetPolicyCount() (*ResponseSdaGetExtranetPolicyCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetExtranetPolicyCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetExtranetPolicyCount()
		}
		return nil, response, fmt.Errorf("error with operation GetExtranetPolicyCount")
	}

	result := response.Result().(*ResponseSdaGetExtranetPolicyCount)
	return result, response, err

}

//GetFabricDevices Get fabric devices - e680-7a97-47db-99e5
/* Returns a list of fabric devices that match the provided query parameters.


@param GetFabricDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices
*/
func (s *SdaService) GetFabricDevices(GetFabricDevicesQueryParams *GetFabricDevicesQueryParams) (*ResponseSdaGetFabricDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices"

	queryString, _ := query.Values(GetFabricDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevices(GetFabricDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevices")
	}

	result := response.Result().(*ResponseSdaGetFabricDevices)
	return result, response, err

}

//GetFabricDevicesCount Get fabric devices count - 9ba6-7b73-44b9-bb42
/* Returns the count of fabric devices that match the provided query parameters.


@param GetFabricDevicesCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-count
*/
func (s *SdaService) GetFabricDevicesCount(GetFabricDevicesCountQueryParams *GetFabricDevicesCountQueryParams) (*ResponseSdaGetFabricDevicesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/count"

	queryString, _ := query.Values(GetFabricDevicesCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesCount(GetFabricDevicesCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesCount")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesCount)
	return result, response, err

}

//GetFabricDevicesLayer2Handoffs Get fabric devices layer 2 handoffs - b7af-eb15-4409-86a4
/* Returns a list of layer 2 handoffs of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer2HandoffsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer2-handoffs
*/
func (s *SdaService) GetFabricDevicesLayer2Handoffs(GetFabricDevicesLayer2HandoffsQueryParams *GetFabricDevicesLayer2HandoffsQueryParams) (*ResponseSdaGetFabricDevicesLayer2Handoffs, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"

	queryString, _ := query.Values(GetFabricDevicesLayer2HandoffsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer2Handoffs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer2Handoffs(GetFabricDevicesLayer2HandoffsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer2Handoffs")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer2Handoffs)
	return result, response, err

}

//GetFabricDevicesLayer2HandoffsCount Get fabric devices layer 2 handoffs count - 019c-791b-48f9-b1d9
/* Returns the count of layer 2 handoffs of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer2HandoffsCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer2-handoffs-count
*/
func (s *SdaService) GetFabricDevicesLayer2HandoffsCount(GetFabricDevicesLayer2HandoffsCountQueryParams *GetFabricDevicesLayer2HandoffsCountQueryParams) (*ResponseSdaGetFabricDevicesLayer2HandoffsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs/count"

	queryString, _ := query.Values(GetFabricDevicesLayer2HandoffsCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer2HandoffsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer2HandoffsCount(GetFabricDevicesLayer2HandoffsCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer2HandoffsCount")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer2HandoffsCount)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithIPTransit Get fabric devices layer 3 handoffs with ip transit - cbb9-daa0-43a9-913b
/* Returns a list of layer 3 handoffs with ip transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithIpTransitQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-ip-transit
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithIPTransit(GetFabricDevicesLayer3HandoffsWithIpTransitQueryParams *GetFabricDevicesLayer3HandoffsWithIPTransitQueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransit, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithIpTransitQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransit{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithIPTransit(GetFabricDevicesLayer3HandoffsWithIpTransitQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithIpTransit")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransit)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithIPTransitCount Get fabric devices layer 3 handoffs with ip transit count - bb90-4a31-4378-9125
/* Returns the count of layer 3 handoffs with ip transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithIpTransitCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-ip-transit-count
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithIPTransitCount(GetFabricDevicesLayer3HandoffsWithIpTransitCountQueryParams *GetFabricDevicesLayer3HandoffsWithIPTransitCountQueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits/count"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithIpTransitCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithIPTransitCount(GetFabricDevicesLayer3HandoffsWithIpTransitCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithIpTransitCount")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithIPTransitCount)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithSdaTransit Get fabric devices layer 3 handoffs with sda transit - 0d8e-d8dd-458b-9dc1
/* Returns a list of layer 3 handoffs with sda transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-sda-transit
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithSdaTransit(GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams *GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransit, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransit{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithSdaTransit(GetFabricDevicesLayer3HandoffsWithSdaTransitQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithSdaTransit")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransit)
	return result, response, err

}

//GetFabricDevicesLayer3HandoffsWithSdaTransitCount Get fabric devices layer 3 handoffs with sda transit count - bd89-6aca-46cb-8f65
/* Returns the count of layer 3 handoffs with sda transit of fabric devices that match the provided query parameters.


@param GetFabricDevicesLayer3HandoffsWithSdaTransitCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-devices-layer3-handoffs-with-sda-transit-count
*/
func (s *SdaService) GetFabricDevicesLayer3HandoffsWithSdaTransitCount(GetFabricDevicesLayer3HandoffsWithSdaTransitCountQueryParams *GetFabricDevicesLayer3HandoffsWithSdaTransitCountQueryParams) (*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits/count"

	queryString, _ := query.Values(GetFabricDevicesLayer3HandoffsWithSdaTransitCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricDevicesLayer3HandoffsWithSdaTransitCount(GetFabricDevicesLayer3HandoffsWithSdaTransitCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricDevicesLayer3HandoffsWithSdaTransitCount")
	}

	result := response.Result().(*ResponseSdaGetFabricDevicesLayer3HandoffsWithSdaTransitCount)
	return result, response, err

}

//GetFabricSites Get fabric sites - b78b-fa87-49a9-b804
/* Returns a list of fabric sites that match the provided query parameters.


@param GetFabricSitesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-sites
*/
func (s *SdaService) GetFabricSites(GetFabricSitesQueryParams *GetFabricSitesQueryParams) (*ResponseSdaGetFabricSites, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites"

	queryString, _ := query.Values(GetFabricSitesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricSites{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricSites(GetFabricSitesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricSites")
	}

	result := response.Result().(*ResponseSdaGetFabricSites)
	return result, response, err

}

//GetFabricSiteCount Get fabric site count - 109a-0907-4a9b-82ef
/* Returns the count of fabric sites that match the provided query parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-site-count
*/
func (s *SdaService) GetFabricSiteCount() (*ResponseSdaGetFabricSiteCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetFabricSiteCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricSiteCount()
		}
		return nil, response, fmt.Errorf("error with operation GetFabricSiteCount")
	}

	result := response.Result().(*ResponseSdaGetFabricSiteCount)
	return result, response, err

}

//GetFabricZones Get fabric zones - d0bc-0b5c-4fdb-839a
/* Returns a list of fabric zones that match the provided query parameters.


@param GetFabricZonesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-zones
*/
func (s *SdaService) GetFabricZones(GetFabricZonesQueryParams *GetFabricZonesQueryParams) (*ResponseSdaGetFabricZones, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones"

	queryString, _ := query.Values(GetFabricZonesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetFabricZones{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricZones(GetFabricZonesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFabricZones")
	}

	result := response.Result().(*ResponseSdaGetFabricZones)
	return result, response, err

}

//GetFabricZoneCount Get fabric zone count - 15a2-da20-4758-bc78
/* Returns the count of fabric zones that match the provided query parameters.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-fabric-zone-count
*/
func (s *SdaService) GetFabricZoneCount() (*ResponseSdaGetFabricZoneCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetFabricZoneCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFabricZoneCount()
		}
		return nil, response, fmt.Errorf("error with operation GetFabricZoneCount")
	}

	result := response.Result().(*ResponseSdaGetFabricZoneCount)
	return result, response, err

}

//GetLayer2VirtualNetworks Get layer 2 virtual networks - 659a-ab00-4c69-a663
/* Returns a list of layer 2 virtual networks that match the provided query parameters.


@param GetLayer2VirtualNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer2-virtual-networks
*/
func (s *SdaService) GetLayer2VirtualNetworks(GetLayer2VirtualNetworksQueryParams *GetLayer2VirtualNetworksQueryParams) (*ResponseSdaGetLayer2VirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	queryString, _ := query.Values(GetLayer2VirtualNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer2VirtualNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer2VirtualNetworks(GetLayer2VirtualNetworksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer2VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaGetLayer2VirtualNetworks)
	return result, response, err

}

//GetLayer2VirtualNetworkCount Get layer 2 virtual network count - 5c9f-0a6e-445b-b743
/* Returns the count of layer 2 virtual networks that match the provided query parameters.


@param GetLayer2VirtualNetworkCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer2-virtual-network-count
*/
func (s *SdaService) GetLayer2VirtualNetworkCount(GetLayer2VirtualNetworkCountQueryParams *GetLayer2VirtualNetworkCountQueryParams) (*ResponseSdaGetLayer2VirtualNetworkCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks/count"

	queryString, _ := query.Values(GetLayer2VirtualNetworkCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer2VirtualNetworkCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer2VirtualNetworkCount(GetLayer2VirtualNetworkCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer2VirtualNetworkCount")
	}

	result := response.Result().(*ResponseSdaGetLayer2VirtualNetworkCount)
	return result, response, err

}

//GetLayer3VirtualNetworks Get layer 3 virtual networks - 2892-e9d4-4b68-b538
/* Returns a list of layer 3 virtual networks that match the provided query parameters.


@param GetLayer3VirtualNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer3-virtual-networks
*/
func (s *SdaService) GetLayer3VirtualNetworks(GetLayer3VirtualNetworksQueryParams *GetLayer3VirtualNetworksQueryParams) (*ResponseSdaGetLayer3VirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	queryString, _ := query.Values(GetLayer3VirtualNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer3VirtualNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer3VirtualNetworks(GetLayer3VirtualNetworksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer3VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaGetLayer3VirtualNetworks)
	return result, response, err

}

//GetLayer3VirtualNetworksCount Get layer 3 virtual networks count - 87af-99e9-493b-9f3d
/* Returns the count of layer 3 virtual networks that match the provided query parameters.


@param GetLayer3VirtualNetworksCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-layer3-virtual-networks-count
*/
func (s *SdaService) GetLayer3VirtualNetworksCount(GetLayer3VirtualNetworksCountQueryParams *GetLayer3VirtualNetworksCountQueryParams) (*ResponseSdaGetLayer3VirtualNetworksCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks/count"

	queryString, _ := query.Values(GetLayer3VirtualNetworksCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetLayer3VirtualNetworksCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetLayer3VirtualNetworksCount(GetLayer3VirtualNetworksCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetLayer3VirtualNetworksCount")
	}

	result := response.Result().(*ResponseSdaGetLayer3VirtualNetworksCount)
	return result, response, err

}

//GetMulticast Get multicast - b48d-e933-4e5b-988a
/* Returns a list of multicast configurations at a fabric site level that match the provided query parameters.


@param GetMulticastQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast
*/
func (s *SdaService) GetMulticast(GetMulticastQueryParams *GetMulticastQueryParams) (*ResponseSdaGetMulticast, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast"

	queryString, _ := query.Values(GetMulticastQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticast{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticast(GetMulticastQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticast")
	}

	result := response.Result().(*ResponseSdaGetMulticast)
	return result, response, err

}

//GetMulticastVirtualNetworks Get multicast virtual networks - 048b-698b-4048-b3aa
/* Returns a list of multicast configurations for virtual networks that match the provided query parameters.


@param GetMulticastVirtualNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-virtual-networks
*/
func (s *SdaService) GetMulticastVirtualNetworks(GetMulticastVirtualNetworksQueryParams *GetMulticastVirtualNetworksQueryParams) (*ResponseSdaGetMulticastVirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks"

	queryString, _ := query.Values(GetMulticastVirtualNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastVirtualNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastVirtualNetworks(GetMulticastVirtualNetworksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastVirtualNetworks")
	}

	result := response.Result().(*ResponseSdaGetMulticastVirtualNetworks)
	return result, response, err

}

//GetMulticastVirtualNetworkCount Get multicast virtual network count - 7cbb-0b86-4a39-98ab
/* Returns the count of multicast configurations associated to virtual networks that match the provided query parameters.


@param GetMulticastVirtualNetworkCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-multicast-virtual-network-count
*/
func (s *SdaService) GetMulticastVirtualNetworkCount(GetMulticastVirtualNetworkCountQueryParams *GetMulticastVirtualNetworkCountQueryParams) (*ResponseSdaGetMulticastVirtualNetworkCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks/count"

	queryString, _ := query.Values(GetMulticastVirtualNetworkCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetMulticastVirtualNetworkCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMulticastVirtualNetworkCount(GetMulticastVirtualNetworkCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMulticastVirtualNetworkCount")
	}

	result := response.Result().(*ResponseSdaGetMulticastVirtualNetworkCount)
	return result, response, err

}

//GetPendingFabricEvents Get pending fabric events - 46aa-996a-4eda-a8fc
/* Returns a list of pending fabric events that match the provided query parameters.


@param GetPendingFabricEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-pending-fabric-events
*/
func (s *SdaService) GetPendingFabricEvents(GetPendingFabricEventsQueryParams *GetPendingFabricEventsQueryParams) (*ResponseSdaGetPendingFabricEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/pendingFabricEvents"

	queryString, _ := query.Values(GetPendingFabricEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPendingFabricEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPendingFabricEvents(GetPendingFabricEventsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPendingFabricEvents")
	}

	result := response.Result().(*ResponseSdaGetPendingFabricEvents)
	return result, response, err

}

//GetPortAssignments Get port assignments - c199-09a2-4619-a140
/* Returns a list of port assignments that match the provided query parameters.


@param GetPortAssignmentsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignments
*/
func (s *SdaService) GetPortAssignments(GetPortAssignmentsQueryParams *GetPortAssignmentsQueryParams) (*ResponseSdaGetPortAssignments, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments"

	queryString, _ := query.Values(GetPortAssignmentsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignments{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignments(GetPortAssignmentsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignments")
	}

	result := response.Result().(*ResponseSdaGetPortAssignments)
	return result, response, err

}

//GetPortAssignmentCount Get port assignment count - 4587-0827-4f1b-a2d4
/* Returns the count of port assignments that match the provided query parameters.


@param GetPortAssignmentCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-assignment-count
*/
func (s *SdaService) GetPortAssignmentCount(GetPortAssignmentCountQueryParams *GetPortAssignmentCountQueryParams) (*ResponseSdaGetPortAssignmentCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments/count"

	queryString, _ := query.Values(GetPortAssignmentCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortAssignmentCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortAssignmentCount(GetPortAssignmentCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentCount")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentCount)
	return result, response, err

}

//GetPortChannels Get port channels - dea6-fbe3-4469-8d79
/* Returns a list of port channels that match the provided query parameters.


@param GetPortChannelsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-channels
*/
func (s *SdaService) GetPortChannels(GetPortChannelsQueryParams *GetPortChannelsQueryParams) (*ResponseSdaGetPortChannels, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels"

	queryString, _ := query.Values(GetPortChannelsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortChannels{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortChannels(GetPortChannelsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortChannels")
	}

	result := response.Result().(*ResponseSdaGetPortChannels)
	return result, response, err

}

//GetPortChannelCount Get port channel count - 7ebb-88ff-4c2b-989c
/* Returns the count of port channels that match the provided query parameters.


@param GetPortChannelCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-port-channel-count
*/
func (s *SdaService) GetPortChannelCount(GetPortChannelCountQueryParams *GetPortChannelCountQueryParams) (*ResponseSdaGetPortChannelCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels/count"

	queryString, _ := query.Values(GetPortChannelCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetPortChannelCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPortChannelCount(GetPortChannelCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPortChannelCount")
	}

	result := response.Result().(*ResponseSdaGetPortChannelCount)
	return result, response, err

}

//GetProvisionedDevices Get provisioned devices - 99b3-ba27-4fe9-9e6b
/* Returns the list of provisioned devices based on query parameters.


@param GetProvisionedDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioned-devices
*/
func (s *SdaService) GetProvisionedDevices(GetProvisionedDevicesQueryParams *GetProvisionedDevicesQueryParams) (*ResponseSdaGetProvisionedDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices"

	queryString, _ := query.Values(GetProvisionedDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetProvisionedDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisionedDevices(GetProvisionedDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProvisionedDevices")
	}

	result := response.Result().(*ResponseSdaGetProvisionedDevices)
	return result, response, err

}

//GetProvisionedDevicesCount Get Provisioned Devices count - e0b3-195e-4678-aeb4
/* Returns the count of provisioned devices based on query parameters.


@param GetProvisionedDevicesCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-provisioned-devices-count
*/
func (s *SdaService) GetProvisionedDevicesCount(GetProvisionedDevicesCountQueryParams *GetProvisionedDevicesCountQueryParams) (*ResponseSdaGetProvisionedDevicesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices/count"

	queryString, _ := query.Values(GetProvisionedDevicesCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetProvisionedDevicesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProvisionedDevicesCount(GetProvisionedDevicesCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProvisionedDevicesCount")
	}

	result := response.Result().(*ResponseSdaGetProvisionedDevicesCount)
	return result, response, err

}

//GetTransitNetworks Get transit networks - e395-fae2-4f0a-b11e
/* Returns a list of transit networks that match the provided query parameters.


@param GetTransitNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-transit-networks
*/
func (s *SdaService) GetTransitNetworks(GetTransitNetworksQueryParams *GetTransitNetworksQueryParams) (*ResponseSdaGetTransitNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks"

	queryString, _ := query.Values(GetTransitNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetTransitNetworks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTransitNetworks(GetTransitNetworksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTransitNetworks")
	}

	result := response.Result().(*ResponseSdaGetTransitNetworks)
	return result, response, err

}

//GetTransitNetworksCount Get transit networks count - 9397-d838-446b-b716
/* Returns the count of transit networks that match the provided query parameters.


@param GetTransitNetworksCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-transit-networks-count
*/
func (s *SdaService) GetTransitNetworksCount(GetTransitNetworksCountQueryParams *GetTransitNetworksCountQueryParams) (*ResponseSdaGetTransitNetworksCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks/count"

	queryString, _ := query.Values(GetTransitNetworksCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetTransitNetworksCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTransitNetworksCount(GetTransitNetworksCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTransitNetworksCount")
	}

	result := response.Result().(*ResponseSdaGetTransitNetworksCount)
	return result, response, err

}

//GetVirtualNetworkWithScalableGroups Get virtual network with scalable groups - ec8a-1ab5-4eba-bca7
/* Get virtual network with scalable groups


@param GetVirtualNetworkWithScalableGroupsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-virtual-network-with-scalable-groups
*/
func (s *SdaService) GetVirtualNetworkWithScalableGroups(GetVirtualNetworkWithScalableGroupsQueryParams *GetVirtualNetworkWithScalableGroupsQueryParams) (*ResponseSdaGetVirtualNetworkWithScalableGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/virtual-network"

	queryString, _ := query.Values(GetVirtualNetworkWithScalableGroupsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetVirtualNetworkWithScalableGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetVirtualNetworkWithScalableGroups(GetVirtualNetworkWithScalableGroupsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaGetVirtualNetworkWithScalableGroups)
	return result, response, err

}

//AddDefaultAuthenticationTemplateInSdaFabric Add default authentication template in SDA Fabric - bca3-39d8-44c8-a3c0
/* Add default authentication template in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-default-authentication-template-in-sda-fabric
*/
func (s *SdaService) AddDefaultAuthenticationTemplateInSdaFabric(requestSdaAddDefaultAuthenticationTemplateInSDAFabric *RequestSdaAddDefaultAuthenticationTemplateInSdaFabric) (*ResponseSdaAddDefaultAuthenticationTemplateInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddDefaultAuthenticationTemplateInSDAFabric).
		SetResult(&ResponseSdaAddDefaultAuthenticationTemplateInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddDefaultAuthenticationTemplateInSdaFabric(requestSdaAddDefaultAuthenticationTemplateInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddDefaultAuthenticationTemplateInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddDefaultAuthenticationTemplateInSdaFabric)
	return result, response, err

}

//AddBorderDeviceInSdaFabric Add border device in SDA Fabric - bead-7b34-43b9-96a7
/* Add border device in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-border-device-in-sda-fabric
*/
func (s *SdaService) AddBorderDeviceInSdaFabric(requestSdaAddBorderDeviceInSDAFabric *RequestSdaAddBorderDeviceInSdaFabric) (*ResponseSdaAddBorderDeviceInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddBorderDeviceInSDAFabric).
		SetResult(&ResponseSdaAddBorderDeviceInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddBorderDeviceInSdaFabric(requestSdaAddBorderDeviceInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddBorderDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddBorderDeviceInSdaFabric)
	return result, response, err

}

//AddControlPlaneDeviceInSdaFabric Add control plane device in SDA Fabric - dd85-c910-4248-9a3f
/* Add control plane device in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-control-plane-device-in-sda-fabric
*/
func (s *SdaService) AddControlPlaneDeviceInSdaFabric(requestSdaAddControlPlaneDeviceInSDAFabric *RequestSdaAddControlPlaneDeviceInSdaFabric) (*ResponseSdaAddControlPlaneDeviceInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddControlPlaneDeviceInSDAFabric).
		SetResult(&ResponseSdaAddControlPlaneDeviceInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddControlPlaneDeviceInSdaFabric(requestSdaAddControlPlaneDeviceInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddControlPlaneDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddControlPlaneDeviceInSdaFabric)
	return result, response, err

}

//AddEdgeDeviceInSdaFabric Add edge device in SDA Fabric - 87a8-ba44-4ce9-bc59
/* Add edge device in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-edge-device-in-sda-fabric
*/
func (s *SdaService) AddEdgeDeviceInSdaFabric(requestSdaAddEdgeDeviceInSDAFabric *RequestSdaAddEdgeDeviceInSdaFabric) (*ResponseSdaAddEdgeDeviceInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/edge-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddEdgeDeviceInSDAFabric).
		SetResult(&ResponseSdaAddEdgeDeviceInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddEdgeDeviceInSdaFabric(requestSdaAddEdgeDeviceInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddEdgeDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddEdgeDeviceInSdaFabric)
	return result, response, err

}

//AddSiteInSdaFabric Add Site in SDA Fabric - d2b4-d9d0-4a4b-884c
/* Add Site in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-site-in-sda-fabric
*/
func (s *SdaService) AddSiteInSdaFabric(requestSdaAddSiteInSDAFabric *RequestSdaAddSiteInSdaFabric) (*ResponseSdaAddSiteInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric-site"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddSiteInSDAFabric).
		SetResult(&ResponseSdaAddSiteInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddSiteInSdaFabric(requestSdaAddSiteInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddSiteInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddSiteInSdaFabric)
	return result, response, err

}

//AddPortAssignmentForAccessPointInSdaFabric Add Port assignment for access point in SDA Fabric - c2a4-3ad2-4098-baa7
/* Add Port assignment for access point in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-assignment-for-access-point-in-sda-fabric
*/
func (s *SdaService) AddPortAssignmentForAccessPointInSdaFabric(requestSdaAddPortAssignmentForAccessPointInSDAFabric *RequestSdaAddPortAssignmentForAccessPointInSdaFabric) (*ResponseSdaAddPortAssignmentForAccessPointInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortAssignmentForAccessPointInSDAFabric).
		SetResult(&ResponseSdaAddPortAssignmentForAccessPointInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortAssignmentForAccessPointInSdaFabric(requestSdaAddPortAssignmentForAccessPointInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddPortAssignmentForAccessPointInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentForAccessPointInSdaFabric)
	return result, response, err

}

//AddPortAssignmentForUserDeviceInSdaFabric Add Port assignment for user device in SDA Fabric - 9582-ab82-4ce8-b29d
/* Add Port assignment for user device in SDA Fabric.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-assignment-for-user-device-in-sda-fabric
*/
func (s *SdaService) AddPortAssignmentForUserDeviceInSdaFabric(requestSdaAddPortAssignmentForUserDeviceInSDAFabric *RequestSdaAddPortAssignmentForUserDeviceInSdaFabric) (*ResponseSdaAddPortAssignmentForUserDeviceInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortAssignmentForUserDeviceInSDAFabric).
		SetResult(&ResponseSdaAddPortAssignmentForUserDeviceInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortAssignmentForUserDeviceInSdaFabric(requestSdaAddPortAssignmentForUserDeviceInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddPortAssignmentForUserDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentForUserDeviceInSdaFabric)
	return result, response, err

}

//AddMulticastInSdaFabric Add multicast in SDA fabric - ff85-3826-472a-98fb
/* Add multicast in SDA fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-multicast-in-sda-fabric
*/
func (s *SdaService) AddMulticastInSdaFabric(requestSdaAddMulticastInSDAFabric *RequestSdaAddMulticastInSdaFabric) (*ResponseSdaAddMulticastInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/multicast"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddMulticastInSDAFabric).
		SetResult(&ResponseSdaAddMulticastInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddMulticastInSdaFabric(requestSdaAddMulticastInSDAFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddMulticastInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddMulticastInSdaFabric)
	return result, response, err

}

//ProvisionWiredDevice Provision Wired Device - cf9a-5843-45fa-9399
/* Provision Wired Device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision-wired-device
*/
func (s *SdaService) ProvisionWiredDevice(requestSdaProvisionWiredDevice *RequestSdaProvisionWiredDevice) (*ResponseSdaProvisionWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/provision-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaProvisionWiredDevice).
		SetResult(&ResponseSdaProvisionWiredDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionWiredDevice(requestSdaProvisionWiredDevice)
		}

		return nil, response, fmt.Errorf("error with operation ProvisionWiredDevice")
	}

	result := response.Result().(*ResponseSdaProvisionWiredDevice)
	return result, response, err

}

//AddTransitPeerNetwork Add Transit Peer Network - 6db9-292d-4f28-a26b
/* Add Transit Peer Network in SD-Access



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-transit-peer-network
*/
func (s *SdaService) AddTransitPeerNetwork(requestSdaAddTransitPeerNetwork *RequestSdaAddTransitPeerNetwork) (*ResponseSdaAddTransitPeerNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/transit-peer-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddTransitPeerNetwork).
		SetResult(&ResponseSdaAddTransitPeerNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddTransitPeerNetwork(requestSdaAddTransitPeerNetwork)
		}

		return nil, response, fmt.Errorf("error with operation AddTransitPeerNetwork")
	}

	result := response.Result().(*ResponseSdaAddTransitPeerNetwork)
	return result, response, err

}

//AddVnInFabric Add VN in fabric - 518c-59cd-441a-a9fc
/* Add virtual network (VN) in SDA Fabric



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-vn-in-fabric
*/
func (s *SdaService) AddVnInFabric(requestSdaAddVNInFabric *RequestSdaAddVnInFabric) (*ResponseSdaAddVnInFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddVNInFabric).
		SetResult(&ResponseSdaAddVnInFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddVnInFabric(requestSdaAddVNInFabric)
		}

		return nil, response, fmt.Errorf("error with operation AddVnInFabric")
	}

	result := response.Result().(*ResponseSdaAddVnInFabric)
	return result, response, err

}

//AddIPPoolInSdaVirtualNetwork Add IP Pool in SDA Virtual Network - 2085-79ea-4ed9-8f4f
/* Add IP Pool in SDA Virtual Network



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-ip-pool-in-sda-virtual-network
*/
func (s *SdaService) AddIPPoolInSdaVirtualNetwork(requestSdaAddIPPoolInSDAVirtualNetwork *RequestSdaAddIPPoolInSdaVirtualNetwork) (*ResponseSdaAddIPPoolInSdaVirtualNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddIPPoolInSDAVirtualNetwork).
		SetResult(&ResponseSdaAddIPPoolInSdaVirtualNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddIPPoolInSdaVirtualNetwork(requestSdaAddIPPoolInSDAVirtualNetwork)
		}

		return nil, response, fmt.Errorf("error with operation AddIpPoolInSdaVirtualNetwork")
	}

	result := response.Result().(*ResponseSdaAddIPPoolInSdaVirtualNetwork)
	return result, response, err

}

//AddAnycastGateways Add anycast gateways - a6fa-2ce6-46ac-7ac5
/* Adds anycast gateways based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-anycast-gateways
*/
func (s *SdaService) AddAnycastGateways(requestSdaAddAnycastGateways *RequestSdaAddAnycastGateways) (*ResponseSdaAddAnycastGateways, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddAnycastGateways).
		SetResult(&ResponseSdaAddAnycastGateways{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddAnycastGateways(requestSdaAddAnycastGateways)
		}

		return nil, response, fmt.Errorf("error with operation AddAnycastGateways")
	}

	result := response.Result().(*ResponseSdaAddAnycastGateways)
	return result, response, err

}

//AddExtranetPolicy Add extranet policy - 1282-78e3-45e8-aae7
/* Adds an extranet policy based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-extranet-policy
*/
func (s *SdaService) AddExtranetPolicy(requestSdaAddExtranetPolicy *RequestSdaAddExtranetPolicy) (*ResponseSdaAddExtranetPolicy, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddExtranetPolicy).
		SetResult(&ResponseSdaAddExtranetPolicy{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddExtranetPolicy(requestSdaAddExtranetPolicy)
		}

		return nil, response, fmt.Errorf("error with operation AddExtranetPolicy")
	}

	result := response.Result().(*ResponseSdaAddExtranetPolicy)
	return result, response, err

}

//AddFabricDevices Add fabric devices - 698a-f8d2-4fd8-b091
/* Adds fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices
*/
func (s *SdaService) AddFabricDevices(requestSdaAddFabricDevices *RequestSdaAddFabricDevices) (*ResponseSdaAddFabricDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevices).
		SetResult(&ResponseSdaAddFabricDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevices(requestSdaAddFabricDevices)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevices")
	}

	result := response.Result().(*ResponseSdaAddFabricDevices)
	return result, response, err

}

//AddFabricDevicesLayer2Handoffs Add fabric devices layer 2 handoffs - 60ae-1ba0-4ebb-9ae1
/* Adds layer 2 handoffs in fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-layer2-handoffs
*/
func (s *SdaService) AddFabricDevicesLayer2Handoffs(requestSdaAddFabricDevicesLayer2Handoffs *RequestSdaAddFabricDevicesLayer2Handoffs) (*ResponseSdaAddFabricDevicesLayer2Handoffs, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesLayer2Handoffs).
		SetResult(&ResponseSdaAddFabricDevicesLayer2Handoffs{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesLayer2Handoffs(requestSdaAddFabricDevicesLayer2Handoffs)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesLayer2Handoffs")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesLayer2Handoffs)
	return result, response, err

}

//AddFabricDevicesLayer3HandoffsWithIPTransit Add fabric devices layer 3 handoffs with ip transit - 248b-d8b2-4c69-b397
/* Adds layer 3 handoffs with ip transit in fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-layer3-handoffs-with-ip-transit
*/
func (s *SdaService) AddFabricDevicesLayer3HandoffsWithIPTransit(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransit *RequestSdaAddFabricDevicesLayer3HandoffsWithIPTransit) (*ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransit, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransit).
		SetResult(&ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransit{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesLayer3HandoffsWithIPTransit(requestSdaAddFabricDevicesLayer3HandoffsWithIpTransit)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesLayer3HandoffsWithIpTransit")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesLayer3HandoffsWithIPTransit)
	return result, response, err

}

//AddFabricDevicesLayer3HandoffsWithSdaTransit Add fabric devices layer 3 handoffs with sda transit - 61a1-aa25-40fa-8312
/* Adds layer 3 handoffs with sda transit in fabric devices based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-devices-layer3-handoffs-with-sda-transit
*/
func (s *SdaService) AddFabricDevicesLayer3HandoffsWithSdaTransit(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransit *RequestSdaAddFabricDevicesLayer3HandoffsWithSdaTransit) (*ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransit, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransit).
		SetResult(&ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransit{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricDevicesLayer3HandoffsWithSdaTransit(requestSdaAddFabricDevicesLayer3HandoffsWithSdaTransit)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricDevicesLayer3HandoffsWithSdaTransit")
	}

	result := response.Result().(*ResponseSdaAddFabricDevicesLayer3HandoffsWithSdaTransit)
	return result, response, err

}

//AddFabricSite Add fabric site - 83af-8a6c-4b58-8f99
/* Adds a fabric site based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-site
*/
func (s *SdaService) AddFabricSite(requestSdaAddFabricSite *RequestSdaAddFabricSite) (*ResponseSdaAddFabricSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricSite).
		SetResult(&ResponseSdaAddFabricSite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricSite(requestSdaAddFabricSite)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricSite")
	}

	result := response.Result().(*ResponseSdaAddFabricSite)
	return result, response, err

}

//AddFabricZone Add fabric zone - 658c-083b-4cb9-92aa
/* Adds a fabric zone based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-fabric-zone
*/
func (s *SdaService) AddFabricZone(requestSdaAddFabricZone *RequestSdaAddFabricZone) (*ResponseSdaAddFabricZone, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabricZone).
		SetResult(&ResponseSdaAddFabricZone{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddFabricZone(requestSdaAddFabricZone)
		}

		return nil, response, fmt.Errorf("error with operation AddFabricZone")
	}

	result := response.Result().(*ResponseSdaAddFabricZone)
	return result, response, err

}

//AddLayer2VirtualNetworks Add layer 2 virtual networks - f8bd-49ed-4f5a-9aec
/* Adds layer 2 virtual networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-layer2-virtual-networks
*/
func (s *SdaService) AddLayer2VirtualNetworks(requestSdaAddLayer2VirtualNetworks *RequestSdaAddLayer2VirtualNetworks) (*ResponseSdaAddLayer2VirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddLayer2VirtualNetworks).
		SetResult(&ResponseSdaAddLayer2VirtualNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddLayer2VirtualNetworks(requestSdaAddLayer2VirtualNetworks)
		}

		return nil, response, fmt.Errorf("error with operation AddLayer2VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaAddLayer2VirtualNetworks)
	return result, response, err

}

//AddLayer3VirtualNetworks Add layer 3 virtual networks - aba8-d8b2-482a-8b53
/* Adds layer 3 virtual networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-layer3-virtual-networks
*/
func (s *SdaService) AddLayer3VirtualNetworks(requestSdaAddLayer3VirtualNetworks *RequestSdaAddLayer3VirtualNetworks) (*ResponseSdaAddLayer3VirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddLayer3VirtualNetworks).
		SetResult(&ResponseSdaAddLayer3VirtualNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddLayer3VirtualNetworks(requestSdaAddLayer3VirtualNetworks)
		}

		return nil, response, fmt.Errorf("error with operation AddLayer3VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaAddLayer3VirtualNetworks)
	return result, response, err

}

//AddMulticastVirtualNetworks Add multicast virtual networks - 0284-eac5-4a3b-abfd
/* Adds multicast for virtual networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-multicast-virtual-networks
*/
func (s *SdaService) AddMulticastVirtualNetworks(requestSdaAddMulticastVirtualNetworks *RequestSdaAddMulticastVirtualNetworks) (*ResponseSdaAddMulticastVirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddMulticastVirtualNetworks).
		SetResult(&ResponseSdaAddMulticastVirtualNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddMulticastVirtualNetworks(requestSdaAddMulticastVirtualNetworks)
		}

		return nil, response, fmt.Errorf("error with operation AddMulticastVirtualNetworks")
	}

	result := response.Result().(*ResponseSdaAddMulticastVirtualNetworks)
	return result, response, err

}

//ApplyPendingFabricEvents Apply pending fabric events - 7b9b-1887-4688-abc7
/* Applies pending fabric events based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!apply-pending-fabric-events
*/
func (s *SdaService) ApplyPendingFabricEvents(requestSdaApplyPendingFabricEvents *RequestSdaApplyPendingFabricEvents) (*ResponseSdaApplyPendingFabricEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/pendingFabricEvents/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaApplyPendingFabricEvents).
		SetResult(&ResponseSdaApplyPendingFabricEvents{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApplyPendingFabricEvents(requestSdaApplyPendingFabricEvents)
		}

		return nil, response, fmt.Errorf("error with operation ApplyPendingFabricEvents")
	}

	result := response.Result().(*ResponseSdaApplyPendingFabricEvents)
	return result, response, err

}

//AddPortAssignments Add port assignments - d8bb-0923-498a-9f63
/* Adds port assignments based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-assignments
*/
func (s *SdaService) AddPortAssignments(requestSdaAddPortAssignments *RequestSdaAddPortAssignments) (*ResponseSdaAddPortAssignments, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortAssignments).
		SetResult(&ResponseSdaAddPortAssignments{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortAssignments(requestSdaAddPortAssignments)
		}

		return nil, response, fmt.Errorf("error with operation AddPortAssignments")
	}

	result := response.Result().(*ResponseSdaAddPortAssignments)
	return result, response, err

}

//AddPortChannels Add port channels - 2ba0-7a63-43db-843c
/* Adds port channels based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-port-channels
*/
func (s *SdaService) AddPortChannels(requestSdaAddPortChannels *RequestSdaAddPortChannels) (*ResponseSdaAddPortChannels, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddPortChannels).
		SetResult(&ResponseSdaAddPortChannels{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddPortChannels(requestSdaAddPortChannels)
		}

		return nil, response, fmt.Errorf("error with operation AddPortChannels")
	}

	result := response.Result().(*ResponseSdaAddPortChannels)
	return result, response, err

}

//ProvisionDevices Provision devices - 9c83-f9d6-4bda-b655
/* Provisions network devices to respective Sites based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision-devices
*/
func (s *SdaService) ProvisionDevices(requestSdaProvisionDevices *RequestSdaProvisionDevices) (*ResponseSdaProvisionDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaProvisionDevices).
		SetResult(&ResponseSdaProvisionDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionDevices(requestSdaProvisionDevices)
		}

		return nil, response, fmt.Errorf("error with operation ProvisionDevices")
	}

	result := response.Result().(*ResponseSdaProvisionDevices)
	return result, response, err

}

//AddTransitNetworks Add transit networks - bd82-db95-41fb-83b9
/* Adds transit networks based on user input.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-transit-networks
*/
func (s *SdaService) AddTransitNetworks(requestSdaAddTransitNetworks *RequestSdaAddTransitNetworks) (*ResponseSdaAddTransitNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddTransitNetworks).
		SetResult(&ResponseSdaAddTransitNetworks{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddTransitNetworks(requestSdaAddTransitNetworks)
		}

		return nil, response, fmt.Errorf("error with operation AddTransitNetworks")
	}

	result := response.Result().(*ResponseSdaAddTransitNetworks)
	return result, response, err

}

//AddVirtualNetworkWithScalableGroups Add virtual network with scalable groups - e3a8-5b19-406a-9f4e
/* Add virtual network with scalable groups at global level



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-virtual-network-with-scalable-groups
*/
func (s *SdaService) AddVirtualNetworkWithScalableGroups(requestSdaAddVirtualNetworkWithScalableGroups *RequestSdaAddVirtualNetworkWithScalableGroups) (*ResponseSdaAddVirtualNetworkWithScalableGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddVirtualNetworkWithScalableGroups).
		SetResult(&ResponseSdaAddVirtualNetworkWithScalableGroups{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddVirtualNetworkWithScalableGroups(requestSdaAddVirtualNetworkWithScalableGroups)
		}

		return nil, response, fmt.Errorf("error with operation AddVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaAddVirtualNetworkWithScalableGroups)
	return result, response, err

}

//UpdateDefaultAuthenticationProfileInSdaFabric Update default authentication profile in SDA Fabric - 8984-ea77-44d9-8a54
/* Update default authentication profile in SDA Fabric


 */
func (s *SdaService) UpdateDefaultAuthenticationProfileInSdaFabric(requestSdaUpdateDefaultAuthenticationProfileInSDAFabric *RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric) (*ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateDefaultAuthenticationProfileInSDAFabric).
		SetResult(&ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabric{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDefaultAuthenticationProfileInSdaFabric(requestSdaUpdateDefaultAuthenticationProfileInSDAFabric)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDefaultAuthenticationProfileInSdaFabric")
	}

	result := response.Result().(*ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabric)
	return result, response, err

}

//ReProvisionWiredDevice Re-Provision Wired Device - 4e95-c9a2-41ab-8889
/* Re-Provision Wired Device


 */
func (s *SdaService) ReProvisionWiredDevice(requestSdaReProvisionWiredDevice *RequestSdaReProvisionWiredDevice) (*ResponseSdaReProvisionWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/provision-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaReProvisionWiredDevice).
		SetResult(&ResponseSdaReProvisionWiredDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReProvisionWiredDevice(requestSdaReProvisionWiredDevice)
		}
		return nil, response, fmt.Errorf("error with operation ReProvisionWiredDevice")
	}

	result := response.Result().(*ResponseSdaReProvisionWiredDevice)
	return result, response, err

}

//UpdateAnycastGateways Update anycast gateways - 0dde-2905-a7b5-4e8d
/* Updates anycast gateways based on user input.


 */
func (s *SdaService) UpdateAnycastGateways(requestSdaUpdateAnycastGateways *RequestSdaUpdateAnycastGateways) (*ResponseSdaUpdateAnycastGateways, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/anycastGateways"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateAnycastGateways).
		SetResult(&ResponseSdaUpdateAnycastGateways{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAnycastGateways(requestSdaUpdateAnycastGateways)
		}
		return nil, response, fmt.Errorf("error with operation UpdateAnycastGateways")
	}

	result := response.Result().(*ResponseSdaUpdateAnycastGateways)
	return result, response, err

}

//UpdateAuthenticationProfile Update authentication profile - 8296-3890-4978-bda1
/* Updates an authentication profile based on user input.


 */
func (s *SdaService) UpdateAuthenticationProfile(requestSdaUpdateAuthenticationProfile *RequestSdaUpdateAuthenticationProfile) (*ResponseSdaUpdateAuthenticationProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/authenticationProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateAuthenticationProfile).
		SetResult(&ResponseSdaUpdateAuthenticationProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAuthenticationProfile(requestSdaUpdateAuthenticationProfile)
		}
		return nil, response, fmt.Errorf("error with operation UpdateAuthenticationProfile")
	}

	result := response.Result().(*ResponseSdaUpdateAuthenticationProfile)
	return result, response, err

}

//UpdateExtranetPolicy Update extranet policy - 899d-9aab-4b0a-8d5a
/* Updates an extranet policy based on user input.


 */
func (s *SdaService) UpdateExtranetPolicy(requestSdaUpdateExtranetPolicy *RequestSdaUpdateExtranetPolicy) (*ResponseSdaUpdateExtranetPolicy, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateExtranetPolicy).
		SetResult(&ResponseSdaUpdateExtranetPolicy{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateExtranetPolicy(requestSdaUpdateExtranetPolicy)
		}
		return nil, response, fmt.Errorf("error with operation UpdateExtranetPolicy")
	}

	result := response.Result().(*ResponseSdaUpdateExtranetPolicy)
	return result, response, err

}

//UpdateFabricDevices Update fabric devices - ceb9-2a9a-409b-8066
/* Updates fabric devices based on user input.


 */
func (s *SdaService) UpdateFabricDevices(requestSdaUpdateFabricDevices *RequestSdaUpdateFabricDevices) (*ResponseSdaUpdateFabricDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricDevices).
		SetResult(&ResponseSdaUpdateFabricDevices{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricDevices(requestSdaUpdateFabricDevices)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricDevices")
	}

	result := response.Result().(*ResponseSdaUpdateFabricDevices)
	return result, response, err

}

//UpdateFabricDevicesLayer3HandoffsWithIPTransit Update fabric devices layer 3 handoffs with ip transit - ff8b-cba9-4829-8ca6
/* Updates layer 3 handoffs with ip transit of fabric devices based on user input.


 */
func (s *SdaService) UpdateFabricDevicesLayer3HandoffsWithIPTransit(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransit *RequestSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit) (*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransit).
		SetResult(&ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricDevicesLayer3HandoffsWithIPTransit(requestSdaUpdateFabricDevicesLayer3HandoffsWithIpTransit)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricDevicesLayer3HandoffsWithIpTransit")
	}

	result := response.Result().(*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithIPTransit)
	return result, response, err

}

//UpdateFabricDevicesLayer3HandoffsWithSdaTransit Update fabric devices layer 3 handoffs with sda transit - b6a1-3a87-435a-a5ed
/* Updates layer 3 handoffs with sda transit of fabric devices based on user input.


 */
func (s *SdaService) UpdateFabricDevicesLayer3HandoffsWithSdaTransit(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit *RequestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit) (*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit).
		SetResult(&ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricDevicesLayer3HandoffsWithSdaTransit(requestSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricDevicesLayer3HandoffsWithSdaTransit")
	}

	result := response.Result().(*ResponseSdaUpdateFabricDevicesLayer3HandoffsWithSdaTransit)
	return result, response, err

}

//UpdateFabricSite Update fabric site - b683-b984-430b-8c74
/* Updates a fabric site based on user input.


 */
func (s *SdaService) UpdateFabricSite(requestSdaUpdateFabricSite *RequestSdaUpdateFabricSite) (*ResponseSdaUpdateFabricSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricSites"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricSite).
		SetResult(&ResponseSdaUpdateFabricSite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricSite(requestSdaUpdateFabricSite)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricSite")
	}

	result := response.Result().(*ResponseSdaUpdateFabricSite)
	return result, response, err

}

//UpdateFabricZone Update fabric zone - f39c-1a54-4c2a-a790
/* Updates a fabric zone based on user input.


 */
func (s *SdaService) UpdateFabricZone(requestSdaUpdateFabricZone *RequestSdaUpdateFabricZone) (*ResponseSdaUpdateFabricZone, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabricZones"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateFabricZone).
		SetResult(&ResponseSdaUpdateFabricZone{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateFabricZone(requestSdaUpdateFabricZone)
		}
		return nil, response, fmt.Errorf("error with operation UpdateFabricZone")
	}

	result := response.Result().(*ResponseSdaUpdateFabricZone)
	return result, response, err

}

//UpdateLayer2VirtualNetworks Update layer 2 virtual networks - 18ba-d85c-4dcb-bc31
/* Updates layer 2 virtual networks based on user input.


 */
func (s *SdaService) UpdateLayer2VirtualNetworks(requestSdaUpdateLayer2VirtualNetworks *RequestSdaUpdateLayer2VirtualNetworks) (*ResponseSdaUpdateLayer2VirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateLayer2VirtualNetworks).
		SetResult(&ResponseSdaUpdateLayer2VirtualNetworks{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateLayer2VirtualNetworks(requestSdaUpdateLayer2VirtualNetworks)
		}
		return nil, response, fmt.Errorf("error with operation UpdateLayer2VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaUpdateLayer2VirtualNetworks)
	return result, response, err

}

//UpdateLayer3VirtualNetworks Update layer 3 virtual networks - c995-fbe5-465b-be45
/* Updates layer 3 virtual networks based on user input.


 */
func (s *SdaService) UpdateLayer3VirtualNetworks(requestSdaUpdateLayer3VirtualNetworks *RequestSdaUpdateLayer3VirtualNetworks) (*ResponseSdaUpdateLayer3VirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateLayer3VirtualNetworks).
		SetResult(&ResponseSdaUpdateLayer3VirtualNetworks{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateLayer3VirtualNetworks(requestSdaUpdateLayer3VirtualNetworks)
		}
		return nil, response, fmt.Errorf("error with operation UpdateLayer3VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaUpdateLayer3VirtualNetworks)
	return result, response, err

}

//UpdateMulticast Update multicast - 9cb0-68b9-4e6b-8c0c
/* Updates a multicast configuration at a fabric level based on user input.


 */
func (s *SdaService) UpdateMulticast(requestSdaUpdateMulticast *RequestSdaUpdateMulticast) (*ResponseSdaUpdateMulticast, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateMulticast).
		SetResult(&ResponseSdaUpdateMulticast{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateMulticast(requestSdaUpdateMulticast)
		}
		return nil, response, fmt.Errorf("error with operation UpdateMulticast")
	}

	result := response.Result().(*ResponseSdaUpdateMulticast)
	return result, response, err

}

//UpdateMulticastVirtualNetworks Update multicast virtual networks - 6c8a-cbbf-4a19-b48b
/* Updates multicast configurations for virtual networks based on user input.


 */
func (s *SdaService) UpdateMulticastVirtualNetworks(requestSdaUpdateMulticastVirtualNetworks *RequestSdaUpdateMulticastVirtualNetworks) (*ResponseSdaUpdateMulticastVirtualNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateMulticastVirtualNetworks).
		SetResult(&ResponseSdaUpdateMulticastVirtualNetworks{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateMulticastVirtualNetworks(requestSdaUpdateMulticastVirtualNetworks)
		}
		return nil, response, fmt.Errorf("error with operation UpdateMulticastVirtualNetworks")
	}

	result := response.Result().(*ResponseSdaUpdateMulticastVirtualNetworks)
	return result, response, err

}

//UpdatePortAssignments Update port assignments - eab7-6a0e-469a-b21a
/* Updates port assignments based on user input.


 */
func (s *SdaService) UpdatePortAssignments(requestSdaUpdatePortAssignments *RequestSdaUpdatePortAssignments) (*ResponseSdaUpdatePortAssignments, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portAssignments"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdatePortAssignments).
		SetResult(&ResponseSdaUpdatePortAssignments{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePortAssignments(requestSdaUpdatePortAssignments)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePortAssignments")
	}

	result := response.Result().(*ResponseSdaUpdatePortAssignments)
	return result, response, err

}

//UpdatePortChannels Update port channels - fc9d-7a51-472b-ba5e
/* Updates port channels based on user input.


 */
func (s *SdaService) UpdatePortChannels(requestSdaUpdatePortChannels *RequestSdaUpdatePortChannels) (*ResponseSdaUpdatePortChannels, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/portChannels"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdatePortChannels).
		SetResult(&ResponseSdaUpdatePortChannels{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePortChannels(requestSdaUpdatePortChannels)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePortChannels")
	}

	result := response.Result().(*ResponseSdaUpdatePortChannels)
	return result, response, err

}

//ReProvisionDevices Re-provision devices - 898a-c9a2-4ee8-8e7f
/* Re-provisions network devices to the site based on the user input.


 */
func (s *SdaService) ReProvisionDevices(requestSdaReProvisionDevices *RequestSdaReProvisionDevices) (*ResponseSdaReProvisionDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/provisionDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaReProvisionDevices).
		SetResult(&ResponseSdaReProvisionDevices{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReProvisionDevices(requestSdaReProvisionDevices)
		}
		return nil, response, fmt.Errorf("error with operation ReProvisionDevices")
	}

	result := response.Result().(*ResponseSdaReProvisionDevices)
	return result, response, err

}

//UpdateTransitNetworks Update transit networks - 6fa2-3824-49f8-a0d7
/* Updates transit networks based on user input.


 */
func (s *SdaService) UpdateTransitNetworks(requestSdaUpdateTransitNetworks *RequestSdaUpdateTransitNetworks) (*ResponseSdaUpdateTransitNetworks, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/transitNetworks"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateTransitNetworks).
		SetResult(&ResponseSdaUpdateTransitNetworks{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTransitNetworks(requestSdaUpdateTransitNetworks)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTransitNetworks")
	}

	result := response.Result().(*ResponseSdaUpdateTransitNetworks)
	return result, response, err

}

//UpdateVirtualNetworkWithScalableGroups Update virtual network with scalable groups - c48b-2904-49bb-875f
/* Update virtual network with scalable groups


 */
func (s *SdaService) UpdateVirtualNetworkWithScalableGroups(requestSdaUpdateVirtualNetworkWithScalableGroups *RequestSdaUpdateVirtualNetworkWithScalableGroups) (*ResponseSdaUpdateVirtualNetworkWithScalableGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaUpdateVirtualNetworkWithScalableGroups).
		SetResult(&ResponseSdaUpdateVirtualNetworkWithScalableGroups{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateVirtualNetworkWithScalableGroups(requestSdaUpdateVirtualNetworkWithScalableGroups)
		}
		return nil, response, fmt.Errorf("error with operation UpdateVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaUpdateVirtualNetworkWithScalableGroups)
	return result, response, err

}

//DeleteDefaultAuthenticationProfileFromSdaFabric Delete default authentication profile from SDA Fabric - 3ebc-da3e-4acb-afb7
/* Delete default authentication profile in SDA Fabric


@param DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-default-authentication-profile-from-sda-fabric
*/
func (s *SdaService) DeleteDefaultAuthenticationProfileFromSdaFabric(DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams *DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams) (*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric, *resty.Response, error) {
	//DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams *DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	queryString, _ := query.Values(DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDefaultAuthenticationProfileFromSdaFabric(DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDefaultAuthenticationProfileFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric)
	return result, response, err

}

//DeleteBorderDeviceFromSdaFabric Delete border device from SDA Fabric - cb81-b935-40ba-aab0
/* Delete border device from SDA Fabric


@param DeleteBorderDeviceFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-border-device-from-sda-fabric
*/
func (s *SdaService) DeleteBorderDeviceFromSdaFabric(DeleteBorderDeviceFromSDAFabricQueryParams *DeleteBorderDeviceFromSdaFabricQueryParams) (*ResponseSdaDeleteBorderDeviceFromSdaFabric, *resty.Response, error) {
	//DeleteBorderDeviceFromSDAFabricQueryParams *DeleteBorderDeviceFromSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(DeleteBorderDeviceFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteBorderDeviceFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteBorderDeviceFromSdaFabric(DeleteBorderDeviceFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteBorderDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteBorderDeviceFromSdaFabric)
	return result, response, err

}

//DeleteControlPlaneDeviceInSdaFabric Delete control plane device in SDA Fabric - f6bd-6bf6-4e68-90be
/* Delete control plane device in SDA Fabric


@param DeleteControlPlaneDeviceInSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-control-plane-device-in-sda-fabric
*/
func (s *SdaService) DeleteControlPlaneDeviceInSdaFabric(DeleteControlPlaneDeviceInSDAFabricQueryParams *DeleteControlPlaneDeviceInSdaFabricQueryParams) (*ResponseSdaDeleteControlPlaneDeviceInSdaFabric, *resty.Response, error) {
	//DeleteControlPlaneDeviceInSDAFabricQueryParams *DeleteControlPlaneDeviceInSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	queryString, _ := query.Values(DeleteControlPlaneDeviceInSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteControlPlaneDeviceInSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteControlPlaneDeviceInSdaFabric(DeleteControlPlaneDeviceInSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteControlPlaneDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteControlPlaneDeviceInSdaFabric)
	return result, response, err

}

//DeleteEdgeDeviceFromSdaFabric Delete edge device from SDA Fabric - 1fb8-f9f2-4c99-8133
/* Delete edge device from SDA Fabric.


@param DeleteEdgeDeviceFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-edge-device-from-sda-fabric
*/
func (s *SdaService) DeleteEdgeDeviceFromSdaFabric(DeleteEdgeDeviceFromSDAFabricQueryParams *DeleteEdgeDeviceFromSdaFabricQueryParams) (*ResponseSdaDeleteEdgeDeviceFromSdaFabric, *resty.Response, error) {
	//DeleteEdgeDeviceFromSDAFabricQueryParams *DeleteEdgeDeviceFromSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/edge-device"

	queryString, _ := query.Values(DeleteEdgeDeviceFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteEdgeDeviceFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteEdgeDeviceFromSdaFabric(DeleteEdgeDeviceFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteEdgeDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteEdgeDeviceFromSdaFabric)
	return result, response, err

}

//DeleteSiteFromSdaFabric Delete Site from SDA Fabric - 5086-4acf-4ad8-b54d
/* Delete Site from SDA Fabric


@param DeleteSiteFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-site-from-sda-fabric
*/
func (s *SdaService) DeleteSiteFromSdaFabric(DeleteSiteFromSDAFabricQueryParams *DeleteSiteFromSdaFabricQueryParams) (*ResponseSdaDeleteSiteFromSdaFabric, *resty.Response, error) {
	//DeleteSiteFromSDAFabricQueryParams *DeleteSiteFromSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/fabric-site"

	queryString, _ := query.Values(DeleteSiteFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteSiteFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSiteFromSdaFabric(DeleteSiteFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSiteFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteSiteFromSdaFabric)
	return result, response, err

}

//DeletePortAssignmentForAccessPointInSdaFabric Delete Port assignment for access point in SDA Fabric - 0787-4a4c-4c9a-abd9
/* Delete Port assignment for access point in SDA Fabric


@param DeletePortAssignmentForAccessPointInSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignment-for-access-point-in-sda-fabric
*/
func (s *SdaService) DeletePortAssignmentForAccessPointInSdaFabric(DeletePortAssignmentForAccessPointInSDAFabricQueryParams *DeletePortAssignmentForAccessPointInSdaFabricQueryParams) (*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric, *resty.Response, error) {
	//DeletePortAssignmentForAccessPointInSDAFabricQueryParams *DeletePortAssignmentForAccessPointInSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	queryString, _ := query.Values(DeletePortAssignmentForAccessPointInSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentForAccessPointInSdaFabric(DeletePortAssignmentForAccessPointInSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentForAccessPointInSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric)
	return result, response, err

}

//DeletePortAssignmentForUserDeviceInSdaFabric Delete Port assignment for user device in SDA Fabric - cba5-b8b1-4edb-81f4
/* Delete Port assignment for user device in SDA Fabric.


@param DeletePortAssignmentForUserDeviceInSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignment-for-user-device-in-sda-fabric
*/
func (s *SdaService) DeletePortAssignmentForUserDeviceInSdaFabric(DeletePortAssignmentForUserDeviceInSDAFabricQueryParams *DeletePortAssignmentForUserDeviceInSdaFabricQueryParams) (*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric, *resty.Response, error) {
	//DeletePortAssignmentForUserDeviceInSDAFabricQueryParams *DeletePortAssignmentForUserDeviceInSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	queryString, _ := query.Values(DeletePortAssignmentForUserDeviceInSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentForUserDeviceInSdaFabric(DeletePortAssignmentForUserDeviceInSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentForUserDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric)
	return result, response, err

}

//DeleteMulticastFromSdaFabric Delete multicast from SDA fabric - 2bb0-0be5-45cb-bc99
/* Delete multicast from SDA fabric


@param DeleteMulticastFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-multicast-from-sda-fabric
*/
func (s *SdaService) DeleteMulticastFromSdaFabric(DeleteMulticastFromSDAFabricQueryParams *DeleteMulticastFromSdaFabricQueryParams) (*ResponseSdaDeleteMulticastFromSdaFabric, *resty.Response, error) {
	//DeleteMulticastFromSDAFabricQueryParams *DeleteMulticastFromSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/multicast"

	queryString, _ := query.Values(DeleteMulticastFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteMulticastFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteMulticastFromSdaFabric(DeleteMulticastFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteMulticastFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteMulticastFromSdaFabric)
	return result, response, err

}

//DeleteProvisionedWiredDevice Delete provisioned Wired Device - e495-b94e-463b-ae04
/* Delete provisioned Wired Device


@param DeleteProvisionedWiredDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-provisioned-wired-device
*/
func (s *SdaService) DeleteProvisionedWiredDevice(DeleteProvisionedWiredDeviceQueryParams *DeleteProvisionedWiredDeviceQueryParams) (*ResponseSdaDeleteProvisionedWiredDevice, *resty.Response, error) {
	//DeleteProvisionedWiredDeviceQueryParams *DeleteProvisionedWiredDeviceQueryParams
	path := "/dna/intent/api/v1/business/sda/provision-device"

	queryString, _ := query.Values(DeleteProvisionedWiredDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteProvisionedWiredDevice{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteProvisionedWiredDevice(DeleteProvisionedWiredDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedWiredDevice")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedWiredDevice)
	return result, response, err

}

//DeleteTransitPeerNetwork Delete Transit Peer Network - d0aa-fa69-4f4b-9d7b
/* Delete Transit Peer Network from SD-Access


@param DeleteTransitPeerNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-transit-peer-network
*/
func (s *SdaService) DeleteTransitPeerNetwork(DeleteTransitPeerNetworkQueryParams *DeleteTransitPeerNetworkQueryParams) (*ResponseSdaDeleteTransitPeerNetwork, *resty.Response, error) {
	//DeleteTransitPeerNetworkQueryParams *DeleteTransitPeerNetworkQueryParams
	path := "/dna/intent/api/v1/business/sda/transit-peer-network"

	queryString, _ := query.Values(DeleteTransitPeerNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteTransitPeerNetwork{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTransitPeerNetwork(DeleteTransitPeerNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTransitPeerNetwork")
	}

	result := response.Result().(*ResponseSdaDeleteTransitPeerNetwork)
	return result, response, err

}

//DeleteVnFromSdaFabric Delete VN from SDA Fabric - c78c-9ad2-45bb-9657
/* Delete virtual network (VN) from SDA Fabric


@param DeleteVNFromSDAFabricQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-vn-from-sda-fabric
*/
func (s *SdaService) DeleteVnFromSdaFabric(DeleteVNFromSDAFabricQueryParams *DeleteVnFromSdaFabricQueryParams) (*ResponseSdaDeleteVnFromSdaFabric, *resty.Response, error) {
	//DeleteVNFromSDAFabricQueryParams *DeleteVnFromSdaFabricQueryParams
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	queryString, _ := query.Values(DeleteVNFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteVnFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteVnFromSdaFabric(DeleteVNFromSDAFabricQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteVnFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteVnFromSdaFabric)
	return result, response, err

}

//DeleteIPPoolFromSdaVirtualNetwork Delete IP Pool from SDA Virtual Network - 549e-4aff-42bb-b52a
/* Delete IP Pool from SDA Virtual Network


@param DeleteIPPoolFromSDAVirtualNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ip-pool-from-sda-virtual-network
*/
func (s *SdaService) DeleteIPPoolFromSdaVirtualNetwork(DeleteIPPoolFromSDAVirtualNetworkQueryParams *DeleteIPPoolFromSdaVirtualNetworkQueryParams) (*ResponseSdaDeleteIPPoolFromSdaVirtualNetwork, *resty.Response, error) {
	//DeleteIPPoolFromSDAVirtualNetworkQueryParams *DeleteIPPoolFromSdaVirtualNetworkQueryParams
	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	queryString, _ := query.Values(DeleteIPPoolFromSDAVirtualNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteIPPoolFromSdaVirtualNetwork{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteIPPoolFromSdaVirtualNetwork(DeleteIPPoolFromSDAVirtualNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteIpPoolFromSdaVirtualNetwork")
	}

	result := response.Result().(*ResponseSdaDeleteIPPoolFromSdaVirtualNetwork)
	return result, response, err

}

//DeleteAnycastGatewayByID Delete anycast gateway by id - 4bfa-d25a-ce07-99f3
/* Deletes an anycast gateway based on id.


@param id id path parameter. ID of the anycast gateway.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-anycast-gateway-by-id
*/
func (s *SdaService) DeleteAnycastGatewayByID(id string) (*ResponseSdaDeleteAnycastGatewayByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/anycastGateways/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteAnycastGatewayByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAnycastGatewayByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAnycastGatewayById")
	}

	result := response.Result().(*ResponseSdaDeleteAnycastGatewayByID)
	return result, response, err

}

//DeleteExtranetPolicies Delete extranet policies - 908a-8bbf-4aeb-9382
/* Deletes extranet policies based on user input.


@param DeleteExtranetPoliciesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-extranet-policies
*/
func (s *SdaService) DeleteExtranetPolicies(DeleteExtranetPoliciesQueryParams *DeleteExtranetPoliciesQueryParams) (*ResponseSdaDeleteExtranetPolicies, *resty.Response, error) {
	//DeleteExtranetPoliciesQueryParams *DeleteExtranetPoliciesQueryParams
	path := "/dna/intent/api/v1/sda/extranetPolicies"

	queryString, _ := query.Values(DeleteExtranetPoliciesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteExtranetPolicies{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteExtranetPolicies(DeleteExtranetPoliciesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteExtranetPolicies")
	}

	result := response.Result().(*ResponseSdaDeleteExtranetPolicies)
	return result, response, err

}

//DeleteExtranetPolicyByID Delete extranet policy by id - 45a7-eb82-446a-b812
/* Deletes an extranet policy based on id.


@param id id path parameter. ID of the extranet policy.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-extranet-policy-by-id
*/
func (s *SdaService) DeleteExtranetPolicyByID(id string) (*ResponseSdaDeleteExtranetPolicyByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/extranetPolicies/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteExtranetPolicyByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteExtranetPolicyByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteExtranetPolicyById")
	}

	result := response.Result().(*ResponseSdaDeleteExtranetPolicyByID)
	return result, response, err

}

//DeleteFabricDevices Delete fabric devices - 8db3-88ed-4018-810a
/* Deletes fabric devices based on user input.


@param DeleteFabricDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-devices
*/
func (s *SdaService) DeleteFabricDevices(DeleteFabricDevicesQueryParams *DeleteFabricDevicesQueryParams) (*ResponseSdaDeleteFabricDevices, *resty.Response, error) {
	//DeleteFabricDevicesQueryParams *DeleteFabricDevicesQueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices"

	queryString, _ := query.Values(DeleteFabricDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDevices{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDevices(DeleteFabricDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDevices")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDevices)
	return result, response, err

}

//DeleteFabricDeviceLayer2Handoffs Delete fabric device layer 2 handoffs - aea7-8b07-48a9-955d
/* Deletes layer 2 handoffs of a fabric device based on user input.


@param DeleteFabricDeviceLayer2HandoffsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer2-handoffs
*/
func (s *SdaService) DeleteFabricDeviceLayer2Handoffs(DeleteFabricDeviceLayer2HandoffsQueryParams *DeleteFabricDeviceLayer2HandoffsQueryParams) (*ResponseSdaDeleteFabricDeviceLayer2Handoffs, *resty.Response, error) {
	//DeleteFabricDeviceLayer2HandoffsQueryParams *DeleteFabricDeviceLayer2HandoffsQueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs"

	queryString, _ := query.Values(DeleteFabricDeviceLayer2HandoffsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDeviceLayer2Handoffs{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer2Handoffs(DeleteFabricDeviceLayer2HandoffsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer2Handoffs")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer2Handoffs)
	return result, response, err

}

//DeleteFabricDeviceLayer2HandoffByID Delete fabric device layer 2 handoff by id - 61ab-38c3-4948-b928
/* Deletes a layer 2 handoff of a fabric device based on id.


@param id id path parameter. ID of the layer 2 handoff of a fabric device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer2-handoff-by-id
*/
func (s *SdaService) DeleteFabricDeviceLayer2HandoffByID(id string) (*ResponseSdaDeleteFabricDeviceLayer2HandoffByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricDevices/layer2Handoffs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricDeviceLayer2HandoffByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer2HandoffByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer2HandoffById")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer2HandoffByID)
	return result, response, err

}

//DeleteFabricDeviceLayer3HandoffsWithIPTransit Delete fabric device layer 3 handoffs with ip transit - 8a87-9bad-45db-8f8f
/* Deletes layer 3 handoffs with ip transit of a fabric device based on user input.


@param DeleteFabricDeviceLayer3HandoffsWithIpTransitQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer3-handoffs-with-ip-transit
*/
func (s *SdaService) DeleteFabricDeviceLayer3HandoffsWithIPTransit(DeleteFabricDeviceLayer3HandoffsWithIpTransitQueryParams *DeleteFabricDeviceLayer3HandoffsWithIPTransitQueryParams) (*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransit, *resty.Response, error) {
	//DeleteFabricDeviceLayer3HandoffsWithIpTransitQueryParams *DeleteFabricDeviceLayer3HandoffsWithIPTransitQueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits"

	queryString, _ := query.Values(DeleteFabricDeviceLayer3HandoffsWithIpTransitQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransit{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer3HandoffsWithIPTransit(DeleteFabricDeviceLayer3HandoffsWithIpTransitQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer3HandoffsWithIpTransit")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithIPTransit)
	return result, response, err

}

//DeleteFabricDeviceLayer3HandoffWithIPTransitByID Delete fabric device layer 3 handoff with ip transit by id - d396-7b01-4d1b-845b
/* Deletes a layer 3 handoff with ip transit of a fabric device by id.


@param id id path parameter. ID of the layer 3 handoff with ip transit of a fabric device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer3-handoff-with-ip-transit-by-id
*/
func (s *SdaService) DeleteFabricDeviceLayer3HandoffWithIPTransitByID(id string) (*ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/ipTransits/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer3HandoffWithIPTransitByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer3HandoffWithIpTransitById")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer3HandoffWithIPTransitByID)
	return result, response, err

}

//DeleteFabricDeviceLayer3HandoffsWithSdaTransit Delete fabric device layer 3 handoffs with sda transit - 8d8c-b8b6-4e9a-a432
/* Deletes layer 3 handoffs with sda transit of a fabric device based on user input.


@param DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-layer3-handoffs-with-sda-transit
*/
func (s *SdaService) DeleteFabricDeviceLayer3HandoffsWithSdaTransit(DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams *DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams) (*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransit, *resty.Response, error) {
	//DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams *DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams
	path := "/dna/intent/api/v1/sda/fabricDevices/layer3Handoffs/sdaTransits"

	queryString, _ := query.Values(DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransit{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceLayer3HandoffsWithSdaTransit(DeleteFabricDeviceLayer3HandoffsWithSdaTransitQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceLayer3HandoffsWithSdaTransit")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceLayer3HandoffsWithSdaTransit)
	return result, response, err

}

//DeleteFabricDeviceByID Delete fabric device by id - 7593-7bef-4a68-b011
/* Deletes a fabric device based on id.


@param id id path parameter. ID of the fabric device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-device-by-id
*/
func (s *SdaService) DeleteFabricDeviceByID(id string) (*ResponseSdaDeleteFabricDeviceByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricDeviceByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricDeviceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricDeviceById")
	}

	result := response.Result().(*ResponseSdaDeleteFabricDeviceByID)
	return result, response, err

}

//DeleteFabricSiteByID Delete fabric site by id - aea4-2a3c-4799-8f73
/* Deletes a fabric site based on id.


@param id id path parameter. ID of the fabric site.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-site-by-id
*/
func (s *SdaService) DeleteFabricSiteByID(id string) (*ResponseSdaDeleteFabricSiteByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricSites/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricSiteByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricSiteByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricSiteById")
	}

	result := response.Result().(*ResponseSdaDeleteFabricSiteByID)
	return result, response, err

}

//DeleteFabricZoneByID Delete fabric zone by id - 78be-d947-4eb8-8fc3
/* Deletes a fabric zone based on id.


@param id id path parameter. ID of the fabric zone.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-fabric-zone-by-id
*/
func (s *SdaService) DeleteFabricZoneByID(id string) (*ResponseSdaDeleteFabricZoneByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/fabricZones/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteFabricZoneByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteFabricZoneByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteFabricZoneById")
	}

	result := response.Result().(*ResponseSdaDeleteFabricZoneByID)
	return result, response, err

}

//DeleteLayer2VirtualNetworks Delete layer 2 virtual networks - d9a4-09cb-4b08-be45
/* Deletes layer 2 virtual networks based on user input.


@param DeleteLayer2VirtualNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer2-virtual-networks
*/
func (s *SdaService) DeleteLayer2VirtualNetworks(DeleteLayer2VirtualNetworksQueryParams *DeleteLayer2VirtualNetworksQueryParams) (*ResponseSdaDeleteLayer2VirtualNetworks, *resty.Response, error) {
	//DeleteLayer2VirtualNetworksQueryParams *DeleteLayer2VirtualNetworksQueryParams
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks"

	queryString, _ := query.Values(DeleteLayer2VirtualNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteLayer2VirtualNetworks{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer2VirtualNetworks(DeleteLayer2VirtualNetworksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer2VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaDeleteLayer2VirtualNetworks)
	return result, response, err

}

//DeleteLayer2VirtualNetworkByID Delete layer 2 virtual network by id - b081-c850-4ab9-86f6
/* Deletes a layer 2 virtual network based on id.


@param id id path parameter. ID of the layer 2 virtual network.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer2-virtual-network-by-id
*/
func (s *SdaService) DeleteLayer2VirtualNetworkByID(id string) (*ResponseSdaDeleteLayer2VirtualNetworkByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/layer2VirtualNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteLayer2VirtualNetworkByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer2VirtualNetworkByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer2VirtualNetworkById")
	}

	result := response.Result().(*ResponseSdaDeleteLayer2VirtualNetworkByID)
	return result, response, err

}

//DeleteLayer3VirtualNetworks Delete layer 3 virtual networks - 49bf-69ec-4a8a-a473
/* Deletes layer 3 virtual networks based on user input.


@param DeleteLayer3VirtualNetworksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer3-virtual-networks
*/
func (s *SdaService) DeleteLayer3VirtualNetworks(DeleteLayer3VirtualNetworksQueryParams *DeleteLayer3VirtualNetworksQueryParams) (*ResponseSdaDeleteLayer3VirtualNetworks, *resty.Response, error) {
	//DeleteLayer3VirtualNetworksQueryParams *DeleteLayer3VirtualNetworksQueryParams
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks"

	queryString, _ := query.Values(DeleteLayer3VirtualNetworksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteLayer3VirtualNetworks{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer3VirtualNetworks(DeleteLayer3VirtualNetworksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer3VirtualNetworks")
	}

	result := response.Result().(*ResponseSdaDeleteLayer3VirtualNetworks)
	return result, response, err

}

//DeleteLayer3VirtualNetworkByID Delete layer 3 virtual network by id - 4cb3-2a8c-4e38-b2ea
/* Deletes a layer 3 virtual network based on id.


@param id id path parameter. ID of the layer 3 virtual network.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-layer3-virtual-network-by-id
*/
func (s *SdaService) DeleteLayer3VirtualNetworkByID(id string) (*ResponseSdaDeleteLayer3VirtualNetworkByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/layer3VirtualNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteLayer3VirtualNetworkByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteLayer3VirtualNetworkByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteLayer3VirtualNetworkById")
	}

	result := response.Result().(*ResponseSdaDeleteLayer3VirtualNetworkByID)
	return result, response, err

}

//DeleteMulticastVirtualNetworkByID Delete multicast virtual network by id - 0b91-0b73-4649-a1d2
/* Deletes a multicast configuration for a virtual network based on id.


@param id id path parameter. ID of the multicast configuration.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-multicast-virtual-network-by-id
*/
func (s *SdaService) DeleteMulticastVirtualNetworkByID(id string) (*ResponseSdaDeleteMulticastVirtualNetworkByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/multicast/virtualNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteMulticastVirtualNetworkByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteMulticastVirtualNetworkByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteMulticastVirtualNetworkById")
	}

	result := response.Result().(*ResponseSdaDeleteMulticastVirtualNetworkByID)
	return result, response, err

}

//DeletePortAssignments Delete port assignments - bd9b-8a54-494b-a3f1
/* Deletes port assignments based on user input.


@param DeletePortAssignmentsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignments
*/
func (s *SdaService) DeletePortAssignments(DeletePortAssignmentsQueryParams *DeletePortAssignmentsQueryParams) (*ResponseSdaDeletePortAssignments, *resty.Response, error) {
	//DeletePortAssignmentsQueryParams *DeletePortAssignmentsQueryParams
	path := "/dna/intent/api/v1/sda/portAssignments"

	queryString, _ := query.Values(DeletePortAssignmentsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortAssignments{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignments(DeletePortAssignmentsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignments")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignments)
	return result, response, err

}

//DeletePortAssignmentByID Delete port assignment by id - fdbe-aa08-422b-9fa1
/* Deletes a port assignment based on id.


@param id id path parameter. ID of the port assignment.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-assignment-by-id
*/
func (s *SdaService) DeletePortAssignmentByID(id string) (*ResponseSdaDeletePortAssignmentByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/portAssignments/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeletePortAssignmentByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortAssignmentByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentById")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentByID)
	return result, response, err

}

//DeletePortChannels Delete port channels - ffb2-e803-4c7b-94ad
/* Deletes port channels based on user input.


@param DeletePortChannelsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-channels
*/
func (s *SdaService) DeletePortChannels(DeletePortChannelsQueryParams *DeletePortChannelsQueryParams) (*ResponseSdaDeletePortChannels, *resty.Response, error) {
	//DeletePortChannelsQueryParams *DeletePortChannelsQueryParams
	path := "/dna/intent/api/v1/sda/portChannels"

	queryString, _ := query.Values(DeletePortChannelsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletePortChannels{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortChannels(DeletePortChannelsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortChannels")
	}

	result := response.Result().(*ResponseSdaDeletePortChannels)
	return result, response, err

}

//DeletePortChannelByID Delete port channel by id - 55ab-6978-47fa-b9d8
/* Deletes a port channel based on id.


@param id id path parameter. ID of the port channel.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-port-channel-by-id
*/
func (s *SdaService) DeletePortChannelByID(id string) (*ResponseSdaDeletePortChannelByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/portChannels/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeletePortChannelByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePortChannelByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePortChannelById")
	}

	result := response.Result().(*ResponseSdaDeletePortChannelByID)
	return result, response, err

}

//DeleteProvisionedDevices Delete provisioned devices - 559a-8ac2-4729-a509
/* Delete provisioned devices based on query parameters.


@param DeleteProvisionedDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-provisioned-devices
*/
func (s *SdaService) DeleteProvisionedDevices(DeleteProvisionedDevicesQueryParams *DeleteProvisionedDevicesQueryParams) (*ResponseSdaDeleteProvisionedDevices, *resty.Response, error) {
	//DeleteProvisionedDevicesQueryParams *DeleteProvisionedDevicesQueryParams
	path := "/dna/intent/api/v1/sda/provisionDevices"

	queryString, _ := query.Values(DeleteProvisionedDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteProvisionedDevices{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteProvisionedDevices(DeleteProvisionedDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedDevices")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedDevices)
	return result, response, err

}

//DeleteProvisionedDeviceByID Delete provisioned device by Id - 8bb4-88f0-4f58-9856
/* Deletes provisioned device based on Id.


@param id id path parameter. ID of the provisioned device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-provisioned-device-by-id
*/
func (s *SdaService) DeleteProvisionedDeviceByID(id string) (*ResponseSdaDeleteProvisionedDeviceByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/provisionDevices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteProvisionedDeviceByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteProvisionedDeviceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedDeviceById")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedDeviceByID)
	return result, response, err

}

//DeleteTransitNetworkByID Delete transit network by id - 91bd-2956-4359-a935
/* Deletes a transit network based on id.


@param id id path parameter. ID of the transit network.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-transit-network-by-id
*/
func (s *SdaService) DeleteTransitNetworkByID(id string) (*ResponseSdaDeleteTransitNetworkByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/sda/transitNetworks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaDeleteTransitNetworkByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTransitNetworkByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTransitNetworkById")
	}

	result := response.Result().(*ResponseSdaDeleteTransitNetworkByID)
	return result, response, err

}

//DeleteVirtualNetworkWithScalableGroups Delete virtual network with scalable groups - c8b6-0bc3-4808-8d56
/* Delete virtual network with scalable groups


@param DeleteVirtualNetworkWithScalableGroupsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-virtual-network-with-scalable-groups
*/
func (s *SdaService) DeleteVirtualNetworkWithScalableGroups(DeleteVirtualNetworkWithScalableGroupsQueryParams *DeleteVirtualNetworkWithScalableGroupsQueryParams) (*ResponseSdaDeleteVirtualNetworkWithScalableGroups, *resty.Response, error) {
	//DeleteVirtualNetworkWithScalableGroupsQueryParams *DeleteVirtualNetworkWithScalableGroupsQueryParams
	path := "/dna/intent/api/v1/virtual-network"

	queryString, _ := query.Values(DeleteVirtualNetworkWithScalableGroupsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteVirtualNetworkWithScalableGroups{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteVirtualNetworkWithScalableGroups(DeleteVirtualNetworkWithScalableGroupsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaDeleteVirtualNetworkWithScalableGroups)
	return result, response, err

}
