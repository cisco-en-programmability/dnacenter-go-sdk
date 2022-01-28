package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DiscoveryService service

type GetDiscoveryJobsByIPQueryParams struct {
	Offset    int    `url:"offset,omitempty"`    //offset
	Limit     int    `url:"limit,omitempty"`     //limit
	IPAddress string `url:"ipAddress,omitempty"` //ipAddress
	Name      string `url:"name,omitempty"`      //name
}
type GetListOfDiscoveriesByDiscoveryIDQueryParams struct {
	Offset    int    `url:"offset,omitempty"`    //offset
	Limit     int    `url:"limit,omitempty"`     //limit
	IPAddress string `url:"ipAddress,omitempty"` //ipAddress
}
type GetDiscoveredNetworkDevicesByDiscoveryIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //taskId
}
type GetDevicesDiscoveredByIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //taskId
}
type GetDiscoveredDevicesByRangeQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //taskId
}
type GetNetworkDevicesFromDiscoveryQueryParams struct {
	TaskID        string `url:"taskId,omitempty"`        //taskId
	SortBy        string `url:"sortBy,omitempty"`        //sortBy
	SortOrder     string `url:"sortOrder,omitempty"`     //sortOrder
	IPAddress     string `url:"ipAddress,omitempty"`     //ipAddress
	PingStatus    string `url:"pingStatus,omitempty"`    //pingStatus
	SNMPStatus    string `url:"snmpStatus,omitempty"`    //snmpStatus
	CliStatus     string `url:"cliStatus,omitempty"`     //cliStatus
	NetconfStatus string `url:"netconfStatus,omitempty"` //netconfStatus
	HTTPStatus    string `url:"httpStatus,omitempty"`    //httpStatus
}
type GetGlobalCredentialsQueryParams struct {
	CredentialSubType string `url:"credentialSubType,omitempty"` //Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3 / HTTP_WRITE / HTTP_READ / NETCONF
	SortBy            string `url:"sortBy,omitempty"`            //sortBy
	Order             string `url:"order,omitempty"`             //order
}

type ResponseDiscoveryDeleteAllDiscovery struct {
	Response *ResponseDiscoveryDeleteAllDiscoveryResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteAllDiscoveryResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID struct {
	Response *ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryStartDiscovery struct {
	Response *ResponseDiscoveryStartDiscoveryResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}
type ResponseDiscoveryStartDiscoveryResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetCountOfAllDiscoveryJobs struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveryJobsByIP struct {
	Response *[]ResponseDiscoveryGetDiscoveryJobsByIPResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveryJobsByIPResponse struct {
	AttributeInfo               *ResponseDiscoveryGetDiscoveryJobsByIPResponseAttributeInfo `json:"attributeInfo,omitempty"`               //
	CliStatus                   string                                                      `json:"cliStatus,omitempty"`                   //
	DiscoveryStatus             string                                                      `json:"discoveryStatus,omitempty"`             //
	EndTime                     string                                                      `json:"endTime,omitempty"`                     //
	HTTPStatus                  string                                                      `json:"httpStatus,omitempty"`                  //
	ID                          string                                                      `json:"id,omitempty"`                          //
	InventoryCollectionStatus   string                                                      `json:"inventoryCollectionStatus,omitempty"`   //
	InventoryReachabilityStatus string                                                      `json:"inventoryReachabilityStatus,omitempty"` //
	IPAddress                   string                                                      `json:"ipAddress,omitempty"`                   //
	JobStatus                   string                                                      `json:"jobStatus,omitempty"`                   //
	Name                        string                                                      `json:"name,omitempty"`                        //
	NetconfStatus               string                                                      `json:"netconfStatus,omitempty"`               //
	PingStatus                  string                                                      `json:"pingStatus,omitempty"`                  //
	SNMPStatus                  string                                                      `json:"snmpStatus,omitempty"`                  //
	StartTime                   string                                                      `json:"startTime,omitempty"`                   //
	TaskID                      string                                                      `json:"taskId,omitempty"`                      //
}
type ResponseDiscoveryGetDiscoveryJobsByIPResponseAttributeInfo interface{}
type ResponseDiscoveryDeleteDiscoveryByID struct {
	Response *ResponseDiscoveryDeleteDiscoveryByIDResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteDiscoveryByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetDiscoveryByID struct {
	Response *ResponseDiscoveryGetDiscoveryByIDResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveryByIDResponse struct {
	AttributeInfo          *ResponseDiscoveryGetDiscoveryByIDResponseAttributeInfo       `json:"attributeInfo,omitempty"`          //
	CdpLevel               *int                                                          `json:"cdpLevel,omitempty"`               //
	DeviceIDs              string                                                        `json:"deviceIds,omitempty"`              //
	DiscoveryCondition     string                                                        `json:"discoveryCondition,omitempty"`     //
	DiscoveryStatus        string                                                        `json:"discoveryStatus,omitempty"`        //
	DiscoveryType          string                                                        `json:"discoveryType,omitempty"`          //
	EnablePasswordList     string                                                        `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string                                                      `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     *ResponseDiscoveryGetDiscoveryByIDResponseHTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *ResponseDiscoveryGetDiscoveryByIDResponseHTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string                                                        `json:"id,omitempty"`                     //
	IPAddressList          string                                                        `json:"ipAddressList,omitempty"`          //
	IPFilterList           string                                                        `json:"ipFilterList,omitempty"`           //
	IsAutoCdp              *bool                                                         `json:"isAutoCdp,omitempty"`              //
	LldpLevel              *int                                                          `json:"lldpLevel,omitempty"`              //
	Name                   string                                                        `json:"name,omitempty"`                   //
	NetconfPort            string                                                        `json:"netconfPort,omitempty"`            //
	NumDevices             *int                                                          `json:"numDevices,omitempty"`             //
	ParentDiscoveryID      string                                                        `json:"parentDiscoveryId,omitempty"`      //
	PasswordList           string                                                        `json:"passwordList,omitempty"`           //
	PreferredMgmtIPMethod  string                                                        `json:"preferredMgmtIPMethod,omitempty"`  //
	ProtocolOrder          string                                                        `json:"protocolOrder,omitempty"`          //
	RetryCount             *int                                                          `json:"retryCount,omitempty"`             //
	SNMPAuthPassphrase     string                                                        `json:"snmpAuthPassphrase,omitempty"`     //
	SNMPAuthProtocol       string                                                        `json:"snmpAuthProtocol,omitempty"`       //
	SNMPMode               string                                                        `json:"snmpMode,omitempty"`               //
	SNMPPrivPassphrase     string                                                        `json:"snmpPrivPassphrase,omitempty"`     //
	SNMPPrivProtocol       string                                                        `json:"snmpPrivProtocol,omitempty"`       //
	SNMPRoCommunity        string                                                        `json:"snmpRoCommunity,omitempty"`        //
	SNMPRoCommunityDesc    string                                                        `json:"snmpRoCommunityDesc,omitempty"`    //
	SNMPRwCommunity        string                                                        `json:"snmpRwCommunity,omitempty"`        //
	SNMPRwCommunityDesc    string                                                        `json:"snmpRwCommunityDesc,omitempty"`    //
	SNMPUserName           string                                                        `json:"snmpUserName,omitempty"`           //
	TimeOut                *int                                                          `json:"timeOut,omitempty"`                //
	UpdateMgmtIP           *bool                                                         `json:"updateMgmtIp,omitempty"`           //
	UserNameList           string                                                        `json:"userNameList,omitempty"`           //
}
type ResponseDiscoveryGetDiscoveryByIDResponseAttributeInfo interface{}
type ResponseDiscoveryGetDiscoveryByIDResponseHTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type ResponseDiscoveryGetDiscoveryByIDResponseHTTPWriteCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type ResponseDiscoveryGetListOfDiscoveriesByDiscoveryID struct {
	Response *[]ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDResponse struct {
	AttributeInfo               *ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDResponseAttributeInfo `json:"attributeInfo,omitempty"`               //
	CliStatus                   string                                                                   `json:"cliStatus,omitempty"`                   //
	DiscoveryStatus             string                                                                   `json:"discoveryStatus,omitempty"`             //
	EndTime                     string                                                                   `json:"endTime,omitempty"`                     //
	HTTPStatus                  string                                                                   `json:"httpStatus,omitempty"`                  //
	ID                          string                                                                   `json:"id,omitempty"`                          //
	InventoryCollectionStatus   string                                                                   `json:"inventoryCollectionStatus,omitempty"`   //
	InventoryReachabilityStatus string                                                                   `json:"inventoryReachabilityStatus,omitempty"` //
	IPAddress                   string                                                                   `json:"ipAddress,omitempty"`                   //
	JobStatus                   string                                                                   `json:"jobStatus,omitempty"`                   //
	Name                        string                                                                   `json:"name,omitempty"`                        //
	NetconfStatus               string                                                                   `json:"netconfStatus,omitempty"`               //
	PingStatus                  string                                                                   `json:"pingStatus,omitempty"`                  //
	SNMPStatus                  string                                                                   `json:"snmpStatus,omitempty"`                  //
	StartTime                   string                                                                   `json:"startTime,omitempty"`                   //
	TaskID                      string                                                                   `json:"taskId,omitempty"`                      //
}
type ResponseDiscoveryGetListOfDiscoveriesByDiscoveryIDResponseAttributeInfo interface{}
type ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryID struct {
	Response *[]ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDResponse `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryIDResponse struct {
	AnchorWlcForAp              string `json:"anchorWlcForAp,omitempty"`              //
	AuthModelID                 string `json:"authModelId,omitempty"`                 //
	AvgUpdateFrequency          *int   `json:"avgUpdateFrequency,omitempty"`          //
	BootDateTime                string `json:"bootDateTime,omitempty"`                //
	CliStatus                   string `json:"cliStatus,omitempty"`                   //
	DuplicateDeviceID           string `json:"duplicateDeviceId,omitempty"`           //
	ErrorCode                   string `json:"errorCode,omitempty"`                   //
	ErrorDescription            string `json:"errorDescription,omitempty"`            //
	Family                      string `json:"family,omitempty"`                      //
	Hostname                    string `json:"hostname,omitempty"`                    //
	HTTPStatus                  string `json:"httpStatus,omitempty"`                  //
	ID                          string `json:"id,omitempty"`                          //
	ImageName                   string `json:"imageName,omitempty"`                   //
	IngressQueueConfig          string `json:"ingressQueueConfig,omitempty"`          //
	InterfaceCount              string `json:"interfaceCount,omitempty"`              //
	InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   //
	InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	LineCardCount               string `json:"lineCardCount,omitempty"`               //
	LineCardID                  string `json:"lineCardId,omitempty"`                  //
	Location                    string `json:"location,omitempty"`                    //
	LocationName                string `json:"locationName,omitempty"`                //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	ManagementIPAddress         string `json:"managementIpAddress,omitempty"`         //
	MemorySize                  string `json:"memorySize,omitempty"`                  //
	NetconfStatus               string `json:"netconfStatus,omitempty"`               //
	NumUpdates                  *int   `json:"numUpdates,omitempty"`                  //
	PingStatus                  string `json:"pingStatus,omitempty"`                  //
	PlatformID                  string `json:"platformId,omitempty"`                  //
	PortRange                   string `json:"portRange,omitempty"`                   //
	QosStatus                   string `json:"qosStatus,omitempty"`                   //
	ReachabilityFailureReason   string `json:"reachabilityFailureReason,omitempty"`   //
	ReachabilityStatus          string `json:"reachabilityStatus,omitempty"`          //
	Role                        string `json:"role,omitempty"`                        //
	RoleSource                  string `json:"roleSource,omitempty"`                  //
	SerialNumber                string `json:"serialNumber,omitempty"`                //
	SNMPContact                 string `json:"snmpContact,omitempty"`                 //
	SNMPLocation                string `json:"snmpLocation,omitempty"`                //
	SNMPStatus                  string `json:"snmpStatus,omitempty"`                  //
	SoftwareVersion             string `json:"softwareVersion,omitempty"`             //
	Tag                         string `json:"tag,omitempty"`                         //
	TagCount                    *int   `json:"tagCount,omitempty"`                    //
	Type                        string `json:"type,omitempty"`                        //
	UpTime                      string `json:"upTime,omitempty"`                      //
	Vendor                      string `json:"vendor,omitempty"`                      //
	WlcApDeviceStatus           string `json:"wlcApDeviceStatus,omitempty"`           //
}
type ResponseDiscoveryGetDevicesDiscoveredByID struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveredDevicesByRange struct {
	Response *[]ResponseDiscoveryGetDiscoveredDevicesByRangeResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveredDevicesByRangeResponse struct {
	AnchorWlcForAp              string `json:"anchorWlcForAp,omitempty"`              //
	AuthModelID                 string `json:"authModelId,omitempty"`                 //
	AvgUpdateFrequency          *int   `json:"avgUpdateFrequency,omitempty"`          //
	BootDateTime                string `json:"bootDateTime,omitempty"`                //
	CliStatus                   string `json:"cliStatus,omitempty"`                   //
	DuplicateDeviceID           string `json:"duplicateDeviceId,omitempty"`           //
	ErrorCode                   string `json:"errorCode,omitempty"`                   //
	ErrorDescription            string `json:"errorDescription,omitempty"`            //
	Family                      string `json:"family,omitempty"`                      //
	Hostname                    string `json:"hostname,omitempty"`                    //
	HTTPStatus                  string `json:"httpStatus,omitempty"`                  //
	ID                          string `json:"id,omitempty"`                          //
	ImageName                   string `json:"imageName,omitempty"`                   //
	IngressQueueConfig          string `json:"ingressQueueConfig,omitempty"`          //
	InterfaceCount              string `json:"interfaceCount,omitempty"`              //
	InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   //
	InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` //
	LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
	LineCardCount               string `json:"lineCardCount,omitempty"`               //
	LineCardID                  string `json:"lineCardId,omitempty"`                  //
	Location                    string `json:"location,omitempty"`                    //
	LocationName                string `json:"locationName,omitempty"`                //
	MacAddress                  string `json:"macAddress,omitempty"`                  //
	ManagementIPAddress         string `json:"managementIpAddress,omitempty"`         //
	MemorySize                  string `json:"memorySize,omitempty"`                  //
	NetconfStatus               string `json:"netconfStatus,omitempty"`               //
	NumUpdates                  *int   `json:"numUpdates,omitempty"`                  //
	PingStatus                  string `json:"pingStatus,omitempty"`                  //
	PlatformID                  string `json:"platformId,omitempty"`                  //
	PortRange                   string `json:"portRange,omitempty"`                   //
	QosStatus                   string `json:"qosStatus,omitempty"`                   //
	ReachabilityFailureReason   string `json:"reachabilityFailureReason,omitempty"`   //
	ReachabilityStatus          string `json:"reachabilityStatus,omitempty"`          //
	Role                        string `json:"role,omitempty"`                        //
	RoleSource                  string `json:"roleSource,omitempty"`                  //
	SerialNumber                string `json:"serialNumber,omitempty"`                //
	SNMPContact                 string `json:"snmpContact,omitempty"`                 //
	SNMPLocation                string `json:"snmpLocation,omitempty"`                //
	SNMPStatus                  string `json:"snmpStatus,omitempty"`                  //
	SoftwareVersion             string `json:"softwareVersion,omitempty"`             //
	Tag                         string `json:"tag,omitempty"`                         //
	TagCount                    *int   `json:"tagCount,omitempty"`                    //
	Type                        string `json:"type,omitempty"`                        //
	UpTime                      string `json:"upTime,omitempty"`                      //
	Vendor                      string `json:"vendor,omitempty"`                      //
	WlcApDeviceStatus           string `json:"wlcApDeviceStatus,omitempty"`           //
}
type ResponseDiscoveryGetNetworkDevicesFromDiscovery struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteDiscoveryBySpecifiedRange struct {
	Response *ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteDiscoveryBySpecifiedRangeResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetDiscoveriesByRange struct {
	Response *[]ResponseDiscoveryGetDiscoveriesByRangeResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetDiscoveriesByRangeResponse struct {
	AttributeInfo          *ResponseDiscoveryGetDiscoveriesByRangeResponseAttributeInfo       `json:"attributeInfo,omitempty"`          //
	CdpLevel               *int                                                               `json:"cdpLevel,omitempty"`               //
	DeviceIDs              string                                                             `json:"deviceIds,omitempty"`              //
	DiscoveryCondition     string                                                             `json:"discoveryCondition,omitempty"`     //
	DiscoveryStatus        string                                                             `json:"discoveryStatus,omitempty"`        //
	DiscoveryType          string                                                             `json:"discoveryType,omitempty"`          //
	EnablePasswordList     string                                                             `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string                                                           `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     *ResponseDiscoveryGetDiscoveriesByRangeResponseHTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *ResponseDiscoveryGetDiscoveriesByRangeResponseHTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string                                                             `json:"id,omitempty"`                     //
	IPAddressList          string                                                             `json:"ipAddressList,omitempty"`          //
	IPFilterList           string                                                             `json:"ipFilterList,omitempty"`           //
	IsAutoCdp              *bool                                                              `json:"isAutoCdp,omitempty"`              //
	LldpLevel              *int                                                               `json:"lldpLevel,omitempty"`              //
	Name                   string                                                             `json:"name,omitempty"`                   //
	NetconfPort            string                                                             `json:"netconfPort,omitempty"`            //
	NumDevices             *int                                                               `json:"numDevices,omitempty"`             //
	ParentDiscoveryID      string                                                             `json:"parentDiscoveryId,omitempty"`      //
	PasswordList           string                                                             `json:"passwordList,omitempty"`           //
	PreferredMgmtIPMethod  string                                                             `json:"preferredMgmtIPMethod,omitempty"`  //
	ProtocolOrder          string                                                             `json:"protocolOrder,omitempty"`          //
	RetryCount             *int                                                               `json:"retryCount,omitempty"`             //
	SNMPAuthPassphrase     string                                                             `json:"snmpAuthPassphrase,omitempty"`     //
	SNMPAuthProtocol       string                                                             `json:"snmpAuthProtocol,omitempty"`       //
	SNMPMode               string                                                             `json:"snmpMode,omitempty"`               //
	SNMPPrivPassphrase     string                                                             `json:"snmpPrivPassphrase,omitempty"`     //
	SNMPPrivProtocol       string                                                             `json:"snmpPrivProtocol,omitempty"`       //
	SNMPRoCommunity        string                                                             `json:"snmpRoCommunity,omitempty"`        //
	SNMPRoCommunityDesc    string                                                             `json:"snmpRoCommunityDesc,omitempty"`    //
	SNMPRwCommunity        string                                                             `json:"snmpRwCommunity,omitempty"`        //
	SNMPRwCommunityDesc    string                                                             `json:"snmpRwCommunityDesc,omitempty"`    //
	SNMPUserName           string                                                             `json:"snmpUserName,omitempty"`           //
	TimeOut                *int                                                               `json:"timeOut,omitempty"`                //
	UpdateMgmtIP           *bool                                                              `json:"updateMgmtIp,omitempty"`           //
	UserNameList           string                                                             `json:"userNameList,omitempty"`           //
}
type ResponseDiscoveryGetDiscoveriesByRangeResponseAttributeInfo interface{}
type ResponseDiscoveryGetDiscoveriesByRangeResponseHTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type ResponseDiscoveryGetDiscoveriesByRangeResponseHTTPWriteCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type ResponseDiscoveryGetGlobalCredentials struct {
	Response *[]ResponseDiscoveryGetGlobalCredentialsResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetGlobalCredentialsResponse struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
}
type ResponseDiscoveryUpdateCliCredentials struct {
	Response *ResponseDiscoveryUpdateCliCredentialsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateCliCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateCliCredentials struct {
	Response *ResponseDiscoveryCreateCliCredentialsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateCliCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateHTTPReadCredentials struct {
	Response *ResponseDiscoveryCreateHTTPReadCredentialsResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateHTTPReadCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateHTTPReadCredential struct {
	Response *ResponseDiscoveryUpdateHTTPReadCredentialResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateHTTPReadCredentialResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateHTTPWriteCredentials struct {
	Response *ResponseDiscoveryUpdateHTTPWriteCredentialsResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateHTTPWriteCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateHTTPWriteCredentials struct {
	Response *ResponseDiscoveryCreateHTTPWriteCredentialsResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateHTTPWriteCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateNetconfCredentials struct {
	Response *ResponseDiscoveryUpdateNetconfCredentialsResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateNetconfCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateNetconfCredentials struct {
	Response *ResponseDiscoveryCreateNetconfCredentialsResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateNetconfCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateSNMPReadCommunity struct {
	Response *ResponseDiscoveryUpdateSNMPReadCommunityResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateSNMPReadCommunityResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateSNMPReadCommunity struct {
	Response *ResponseDiscoveryCreateSNMPReadCommunityResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateSNMPReadCommunityResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateSNMPWriteCommunity struct {
	Response *ResponseDiscoveryCreateSNMPWriteCommunityResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateSNMPWriteCommunityResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateSNMPWriteCommunity struct {
	Response *ResponseDiscoveryUpdateSNMPWriteCommunityResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateSNMPWriteCommunityResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateSNMPv3Credentials struct {
	Response *ResponseDiscoveryUpdateSNMPv3CredentialsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateSNMPv3CredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryCreateSNMPv3Credentials struct {
	Response *ResponseDiscoveryCreateSNMPv3CredentialsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateSNMPv3CredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryDeleteGlobalCredentialsByID struct {
	Response *ResponseDiscoveryDeleteGlobalCredentialsByIDResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  //
}
type ResponseDiscoveryDeleteGlobalCredentialsByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryUpdateGlobalCredentials struct {
	Response *ResponseDiscoveryUpdateGlobalCredentialsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDiscoveryUpdateGlobalCredentialsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDiscoveryGetCredentialSubTypeByCredentialID struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetSNMPProperties struct {
	Response *[]ResponseDiscoveryGetSNMPPropertiesResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  //
}
type ResponseDiscoveryGetSNMPPropertiesResponse struct {
	ID                 string `json:"id,omitempty"`                 //
	InstanceTenantID   string `json:"instanceTenantId,omitempty"`   //
	InstanceUUID       string `json:"instanceUuid,omitempty"`       //
	IntValue           *int   `json:"intValue,omitempty"`           //
	SystemPropertyName string `json:"systemPropertyName,omitempty"` //
}
type ResponseDiscoveryCreateUpdateSNMPProperties struct {
	Response *ResponseDiscoveryCreateUpdateSNMPPropertiesResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseDiscoveryCreateUpdateSNMPPropertiesResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID struct {
	AttributeInfo          *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo       `json:"attributeInfo,omitempty"`          //
	CdpLevel               *int                                                                        `json:"cdpLevel,omitempty"`               //
	DeviceIDs              string                                                                      `json:"deviceIds,omitempty"`              //
	DiscoveryCondition     string                                                                      `json:"discoveryCondition,omitempty"`     //
	DiscoveryStatus        string                                                                      `json:"discoveryStatus,omitempty"`        //
	DiscoveryType          string                                                                      `json:"discoveryType,omitempty"`          //
	EnablePasswordList     string                                                                      `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string                                                                    `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string                                                                      `json:"id,omitempty"`                     //
	IPAddressList          string                                                                      `json:"ipAddressList,omitempty"`          //
	IPFilterList           string                                                                      `json:"ipFilterList,omitempty"`           //
	IsAutoCdp              *bool                                                                       `json:"isAutoCdp,omitempty"`              //
	LldpLevel              *int                                                                        `json:"lldpLevel,omitempty"`              //
	Name                   string                                                                      `json:"name,omitempty"`                   //
	NetconfPort            string                                                                      `json:"netconfPort,omitempty"`            //
	NumDevices             *int                                                                        `json:"numDevices,omitempty"`             //
	ParentDiscoveryID      string                                                                      `json:"parentDiscoveryId,omitempty"`      //
	PasswordList           string                                                                      `json:"passwordList,omitempty"`           //
	PreferredMgmtIPMethod  string                                                                      `json:"preferredMgmtIPMethod,omitempty"`  //
	ProtocolOrder          string                                                                      `json:"protocolOrder,omitempty"`          //
	RetryCount             *int                                                                        `json:"retryCount,omitempty"`             //
	SNMPAuthPassphrase     string                                                                      `json:"snmpAuthPassphrase,omitempty"`     //
	SNMPAuthProtocol       string                                                                      `json:"snmpAuthProtocol,omitempty"`       //
	SNMPMode               string                                                                      `json:"snmpMode,omitempty"`               //
	SNMPPrivPassphrase     string                                                                      `json:"snmpPrivPassphrase,omitempty"`     //
	SNMPPrivProtocol       string                                                                      `json:"snmpPrivProtocol,omitempty"`       //
	SNMPRoCommunity        string                                                                      `json:"snmpRoCommunity,omitempty"`        //
	SNMPRoCommunityDesc    string                                                                      `json:"snmpRoCommunityDesc,omitempty"`    //
	SNMPRwCommunity        string                                                                      `json:"snmpRwCommunity,omitempty"`        //
	SNMPRwCommunityDesc    string                                                                      `json:"snmpRwCommunityDesc,omitempty"`    //
	SNMPUserName           string                                                                      `json:"snmpUserName,omitempty"`           //
	TimeOut                *int                                                                        `json:"timeOut,omitempty"`                //
	UpdateMgmtIP           *bool                                                                       `json:"updateMgmtIp,omitempty"`           //
	UserNameList           string                                                                      `json:"userNameList,omitempty"`           //
}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDAttributeInfo interface{}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedIDHTTPWriteCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryStartDiscovery struct {
	CdpLevel               *int                                               `json:"cdpLevel,omitempty"`               // CDP level to which neighbor devices to be discovered
	DiscoveryType          string                                             `json:"discoveryType,omitempty"`          // Type of Discovery. 'SINGLE', 'RANGE', 'MULTI RANGE', 'CDP', 'LLDP'
	EnablePasswordList     []string                                           `json:"enablePasswordList,omitempty"`     // Enable Password of the devices to be discovered
	GlobalCredentialIDList []string                                           `json:"globalCredentialIdList,omitempty"` // Global Credential Ids to be used for discovery
	HTTPReadCredential     *RequestDiscoveryStartDiscoveryHTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    *RequestDiscoveryStartDiscoveryHTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	IPAddressList          string                                             `json:"ipAddressList,omitempty"`          // IP Address of devices to be discovered. Ex: '172.30.0.1' for SINGLE, CDP and LLDP; '72.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE
	IPFilterList           []string                                           `json:"ipFilterList,omitempty"`           // IP Addresses of the devices to be filtered out during discovery
	LldpLevel              *int                                               `json:"lldpLevel,omitempty"`              // LLDP level to which neighbor devices to be discovered
	Name                   string                                             `json:"name,omitempty"`                   // Name of the discovery
	NetconfPort            string                                             `json:"netconfPort,omitempty"`            // Netconf Port. It will need valid SSH credentials to work
	PasswordList           []string                                           `json:"passwordList,omitempty"`           // Password of the devices to be discovered
	PreferredMgmtIPMethod  string                                             `json:"preferredMgmtIPMethod,omitempty"`  // Preferred Management IP Method.'' or 'UseLoopBack'. Default is ''
	ProtocolOrder          string                                             `json:"protocolOrder,omitempty"`          // Order of protocol (ssh/telnet) in which device connection will be tried. Ex: 'telnet': only telnet; 'ssh,telnet': ssh with higher order than telnet
	Retry                  *int                                               `json:"retry,omitempty"`                  // Number of times to try establishing connection to device
	SNMPAuthPassphrase     string                                             `json:"snmpAuthPassphrase,omitempty"`     // Auth Pass phrase for SNMP
	SNMPAuthProtocol       string                                             `json:"snmpAuthProtocol,omitempty"`       // SNMP auth protocol. SHA' or 'MD5'
	SNMPMode               string                                             `json:"snmpMode,omitempty"`               // Mode of SNMP. 'AUTHPRIV' or 'AUTHNOPRIV' or 'NOAUTHNOPRIV'
	SNMPPrivPassphrase     string                                             `json:"snmpPrivPassphrase,omitempty"`     // Pass phrase for SNMP privacy
	SNMPPrivProtocol       string                                             `json:"snmpPrivProtocol,omitempty"`       // SNMP privacy protocol. 'DES' or 'AES128'
	SNMPROCommunity        string                                             `json:"snmpROCommunity,omitempty"`        // Snmp RO community of the devices to be discovered
	SNMPROCommunityDesc    string                                             `json:"snmpROCommunityDesc,omitempty"`    // Description for Snmp RO community
	SNMPRWCommunity        string                                             `json:"snmpRWCommunity,omitempty"`        // Snmp RW community of the devices to be discovered
	SNMPRWCommunityDesc    string                                             `json:"snmpRWCommunityDesc,omitempty"`    // Description for Snmp RW community
	SNMPUserName           string                                             `json:"snmpUserName,omitempty"`           // SNMP username of the device
	SNMPVersion            string                                             `json:"snmpVersion,omitempty"`            // Version of SNMP. v2 or v3
	Timeout                *int                                               `json:"timeout,omitempty"`                // Time to wait for device response in seconds
	UserNameList           []string                                           `json:"userNameList,omitempty"`           // Username of the devices to be discovered
}
type RequestDiscoveryStartDiscoveryHTTPReadCredential struct {
	Password string `json:"password,omitempty"` // HTTP(S) password
	Port     *int   `json:"port,omitempty"`     // HTTP(S) port
	Secure   *bool  `json:"secure,omitempty"`   // Flag for HTTPS
	Username string `json:"username,omitempty"` // HTTP(S) username
}
type RequestDiscoveryStartDiscoveryHTTPWriteCredential struct {
	Password string `json:"password,omitempty"` // HTTP(S) password
	Port     *int   `json:"port,omitempty"`     // HTTP(S) port
	Secure   *bool  `json:"secure,omitempty"`   // Flag for HTTPS
	Username string `json:"username,omitempty"` // HTTP(S) username
}
type RequestDiscoveryUpdateCliCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	EnablePassword   string `json:"enablePassword,omitempty"`   //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryCreateCliCredentials []RequestItemDiscoveryCreateCliCredentials // Array of RequestDiscoveryCreateCLICredentials
type RequestItemDiscoveryCreateCliCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	EnablePassword   string `json:"enablePassword,omitempty"`   //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryCreateHTTPReadCredentials []RequestItemDiscoveryCreateHTTPReadCredentials // Array of RequestDiscoveryCreateHTTPReadCredentials
type RequestItemDiscoveryCreateHTTPReadCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryUpdateHTTPReadCredential struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryUpdateHTTPWriteCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryCreateHTTPWriteCredentials []RequestItemDiscoveryCreateHTTPWriteCredentials // Array of RequestDiscoveryCreateHTTPWriteCredentials
type RequestItemDiscoveryCreateHTTPWriteCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             *int   `json:"port,omitempty"`             //
	Secure           *bool  `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryUpdateNetconfCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	NetconfPort      string `json:"netconfPort,omitempty"`      //
}
type RequestDiscoveryCreateNetconfCredentials []RequestItemDiscoveryCreateNetconfCredentials // Array of RequestDiscoveryCreateNetconfCredentials
type RequestItemDiscoveryCreateNetconfCredentials struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	NetconfPort      string `json:"netconfPort,omitempty"`      //
}
type RequestDiscoveryUpdateSNMPReadCommunity struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`   //
	ReadCommunity  string `json:"readCommunity,omitempty"`  // SNMP read community. NO!$DATA!$ for no value change
}
type RequestDiscoveryCreateSNMPReadCommunity []RequestItemDiscoveryCreateSNMPReadCommunity // Array of RequestDiscoveryCreateSNMPReadCommunity
type RequestItemDiscoveryCreateSNMPReadCommunity struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	ReadCommunity  string `json:"readCommunity,omitempty"`  // SNMP read community
}
type RequestDiscoveryCreateSNMPWriteCommunity []RequestItemDiscoveryCreateSNMPWriteCommunity // Array of RequestDiscoveryCreateSNMPWriteCommunity
type RequestItemDiscoveryCreateSNMPWriteCommunity struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	WriteCommunity string `json:"writeCommunity,omitempty"` // SNMP write community
}
type RequestDiscoveryUpdateSNMPWriteCommunity struct {
	Comments       string `json:"comments,omitempty"`       // Comments to identify the credential
	CredentialType string `json:"credentialType,omitempty"` // Credential type to identify the application that uses the credential
	Description    string `json:"description,omitempty"`    // Name/Description of the credential
	InstanceUUID   string `json:"instanceUuid,omitempty"`   //
	WriteCommunity string `json:"writeCommunity,omitempty"` // SNMP write community. NO!$DATA!$ for no value change
}
type RequestDiscoveryUpdateSNMPv3Credentials struct {
	AuthPassword     string `json:"authPassword,omitempty"`     //
	AuthType         string `json:"authType,omitempty"`         //
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  //
	PrivacyType      string `json:"privacyType,omitempty"`      //
	SNMPMode         string `json:"snmpMode,omitempty"`         //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryCreateSNMPv3Credentials []RequestItemDiscoveryCreateSNMPv3Credentials // Array of RequestDiscoveryCreateSNMPv3Credentials
type RequestItemDiscoveryCreateSNMPv3Credentials struct {
	AuthPassword     string `json:"authPassword,omitempty"`     //
	AuthType         string `json:"authType,omitempty"`         //
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  //
	PrivacyType      string `json:"privacyType,omitempty"`      //
	SNMPMode         string `json:"snmpMode,omitempty"`         //
	Username         string `json:"username,omitempty"`         //
}
type RequestDiscoveryUpdateGlobalCredentials struct {
	SiteUUIDs []string `json:"siteUuids,omitempty"` //
}
type RequestDiscoveryCreateUpdateSNMPProperties []RequestItemDiscoveryCreateUpdateSNMPProperties // Array of RequestDiscoveryCreateUpdateSNMPProperties
type RequestItemDiscoveryCreateUpdateSNMPProperties struct {
	ID                 string `json:"id,omitempty"`                 //
	InstanceTenantID   string `json:"instanceTenantId,omitempty"`   //
	InstanceUUID       string `json:"instanceUuid,omitempty"`       //
	IntValue           *int   `json:"intValue,omitempty"`           //
	SystemPropertyName string `json:"systemPropertyName,omitempty"` //
}

//GetCountOfAllDiscoveryJobs Get count of all discovery jobs - 069d-9823-451b-892d
/* Returns the count of all available discovery jobs


 */
func (s *DiscoveryService) GetCountOfAllDiscoveryJobs() (*ResponseDiscoveryGetCountOfAllDiscoveryJobs, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetCountOfAllDiscoveryJobs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCountOfAllDiscoveryJobs")
	}

	result := response.Result().(*ResponseDiscoveryGetCountOfAllDiscoveryJobs)
	return result, response, err

}

//GetDiscoveryJobsByIP Get Discovery jobs by IP - a496-7be6-4dfa-aa1a
/* Returns the list of discovery jobs for the given IP


@param GetDiscoveryJobsByIPQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetDiscoveryJobsByIP(GetDiscoveryJobsByIPQueryParams *GetDiscoveryJobsByIPQueryParams) (*ResponseDiscoveryGetDiscoveryJobsByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/job"

	queryString, _ := query.Values(GetDiscoveryJobsByIPQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDiscoveryJobsByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDiscoveryJobsByIp")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveryJobsByIP)
	return result, response, err

}

//GetDiscoveryByID Get Discovery by Id - 63bb-88b7-4f59-aa17
/* Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

*/
func (s *DiscoveryService) GetDiscoveryByID(id string) (*ResponseDiscoveryGetDiscoveryByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetDiscoveryByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDiscoveryById")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveryByID)
	return result, response, err

}

//GetListOfDiscoveriesByDiscoveryID Get list of discoveries by discovery Id - 9987-2a13-4d0a-9fb4
/* Returns the list of discovery jobs for the given Discovery ID. The results can be optionally filtered based on IP. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetListOfDiscoveriesByDiscoveryIdQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetListOfDiscoveriesByDiscoveryID(id string, GetListOfDiscoveriesByDiscoveryIdQueryParams *GetListOfDiscoveriesByDiscoveryIDQueryParams) (*ResponseDiscoveryGetListOfDiscoveriesByDiscoveryID, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/job"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetListOfDiscoveriesByDiscoveryIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetListOfDiscoveriesByDiscoveryID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetListOfDiscoveriesByDiscoveryId")
	}

	result := response.Result().(*ResponseDiscoveryGetListOfDiscoveriesByDiscoveryID)
	return result, response, err

}

//GetDiscoveredNetworkDevicesByDiscoveryID Get Discovered network devices by discovery Id - f6ac-994f-451b-a011
/* Returns the network devices discovered for the given Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetDiscoveredNetworkDevicesByDiscoveryIdQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetDiscoveredNetworkDevicesByDiscoveryID(id string, GetDiscoveredNetworkDevicesByDiscoveryIdQueryParams *GetDiscoveredNetworkDevicesByDiscoveryIDQueryParams) (*ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryID, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/network-device"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDiscoveredNetworkDevicesByDiscoveryIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDiscoveredNetworkDevicesByDiscoveryId")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveredNetworkDevicesByDiscoveryID)
	return result, response, err

}

//GetDevicesDiscoveredByID Get Devices discovered by Id - a696-5b45-4c9a-8663
/* Returns the count of network devices discovered in the given discovery. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetDevicesDiscoveredByIdQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetDevicesDiscoveredByID(id string, GetDevicesDiscoveredByIdQueryParams *GetDevicesDiscoveredByIDQueryParams) (*ResponseDiscoveryGetDevicesDiscoveredByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/network-device/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDevicesDiscoveredByIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDevicesDiscoveredByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDevicesDiscoveredById")
	}

	result := response.Result().(*ResponseDiscoveryGetDevicesDiscoveredByID)
	return result, response, err

}

//GetDiscoveredDevicesByRange Get Discovered devices by range - a6b7-98ab-4aca-a34e
/* Returns the network devices discovered for the given discovery and for the given range. The maximum number of records that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return

@param GetDiscoveredDevicesByRangeQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetDiscoveredDevicesByRange(id string, startIndex int, recordsToReturn int, GetDiscoveredDevicesByRangeQueryParams *GetDiscoveredDevicesByRangeQueryParams) (*ResponseDiscoveryGetDiscoveredDevicesByRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	queryString, _ := query.Values(GetDiscoveredDevicesByRangeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetDiscoveredDevicesByRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDiscoveredDevicesByRange")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveredDevicesByRange)
	return result, response, err

}

//GetNetworkDevicesFromDiscovery Get network devices from Discovery - 3d9b-99c3-4339-8a27
/* Returns the network devices from a discovery job based on given filters. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

@param GetNetworkDevicesFromDiscoveryQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetNetworkDevicesFromDiscovery(id string, GetNetworkDevicesFromDiscoveryQueryParams *GetNetworkDevicesFromDiscoveryQueryParams) (*ResponseDiscoveryGetNetworkDevicesFromDiscovery, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}/summary"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetNetworkDevicesFromDiscoveryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetNetworkDevicesFromDiscovery{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNetworkDevicesFromDiscovery")
	}

	result := response.Result().(*ResponseDiscoveryGetNetworkDevicesFromDiscovery)
	return result, response, err

}

//GetDiscoveriesByRange Get Discoveries by range - 33b7-99d0-4d0a-8907
/* Returns the discovery by specified range


@param startIndex startIndex path parameter. Start index

@param recordsToReturn recordsToReturn path parameter. Number of records to return

*/
func (s *DiscoveryService) GetDiscoveriesByRange(startIndex int, recordsToReturn int) (*ResponseDiscoveryGetDiscoveriesByRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToReturn}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetDiscoveriesByRange{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDiscoveriesByRange")
	}

	result := response.Result().(*ResponseDiscoveryGetDiscoveriesByRange)
	return result, response, err

}

//GetGlobalCredentials Get Global credentials - ff81-6b8e-4358-97eb
/* Returns global credential for the given credential sub type


@param GetGlobalCredentialsQueryParams Filtering parameter
*/
func (s *DiscoveryService) GetGlobalCredentials(GetGlobalCredentialsQueryParams *GetGlobalCredentialsQueryParams) (*ResponseDiscoveryGetGlobalCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential"

	queryString, _ := query.Values(GetGlobalCredentialsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDiscoveryGetGlobalCredentials{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGlobalCredentials")
	}

	result := response.Result().(*ResponseDiscoveryGetGlobalCredentials)
	return result, response, err

}

//GetCredentialSubTypeByCredentialID Get Credential sub type by credential Id - 58a3-699e-489b-9529
/* Returns the credential sub type for the given Id


@param id id path parameter. Global Credential ID

*/
func (s *DiscoveryService) GetCredentialSubTypeByCredentialID(id string) (*ResponseDiscoveryGetCredentialSubTypeByCredentialID, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetCredentialSubTypeByCredentialID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCredentialSubTypeByCredentialId")
	}

	result := response.Result().(*ResponseDiscoveryGetCredentialSubTypeByCredentialID)
	return result, response, err

}

//GetSNMPProperties Get SNMP properties - 4497-4ba5-435a-801d
/* Returns SNMP properties


 */
func (s *DiscoveryService) GetSNMPProperties() (*ResponseDiscoveryGetSNMPProperties, *resty.Response, error) {
	path := "/dna/intent/api/v1/snmp-property"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryGetSNMPProperties{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSnmpProperties")
	}

	result := response.Result().(*ResponseDiscoveryGetSNMPProperties)
	return result, response, err

}

//StartDiscovery Start discovery - 55b4-39dc-4239-b140
/* Initiates discovery with the given parameters


 */
func (s *DiscoveryService) StartDiscovery(requestDiscoveryStartDiscovery *RequestDiscoveryStartDiscovery) (*ResponseDiscoveryStartDiscovery, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryStartDiscovery).
		SetResult(&ResponseDiscoveryStartDiscovery{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation StartDiscovery")
	}

	result := response.Result().(*ResponseDiscoveryStartDiscovery)
	return result, response, err

}

//CreateCliCredentials Create CLI credentials - 948e-a819-4348-bc0b
/* Adds global CLI credential


 */
func (s *DiscoveryService) CreateCliCredentials(requestDiscoveryCreateCLICredentials *RequestDiscoveryCreateCliCredentials) (*ResponseDiscoveryCreateCliCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/cli"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateCLICredentials).
		SetResult(&ResponseDiscoveryCreateCliCredentials{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateCliCredentials")
	}

	result := response.Result().(*ResponseDiscoveryCreateCliCredentials)
	return result, response, err

}

//CreateHTTPReadCredentials Create HTTP read credentials - bf85-9ac6-4a0b-a19c
/* Adds HTTP read credentials


 */
func (s *DiscoveryService) CreateHTTPReadCredentials(requestDiscoveryCreateHTTPReadCredentials *RequestDiscoveryCreateHTTPReadCredentials) (*ResponseDiscoveryCreateHTTPReadCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-read"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateHTTPReadCredentials).
		SetResult(&ResponseDiscoveryCreateHTTPReadCredentials{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateHttpReadCredentials")
	}

	result := response.Result().(*ResponseDiscoveryCreateHTTPReadCredentials)
	return result, response, err

}

//CreateHTTPWriteCredentials Create HTTP write credentials - 4d9c-a8e2-431a-8a24
/* Adds global HTTP write credentials


 */
func (s *DiscoveryService) CreateHTTPWriteCredentials(requestDiscoveryCreateHTTPWriteCredentials *RequestDiscoveryCreateHTTPWriteCredentials) (*ResponseDiscoveryCreateHTTPWriteCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-write"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateHTTPWriteCredentials).
		SetResult(&ResponseDiscoveryCreateHTTPWriteCredentials{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateHttpWriteCredentials")
	}

	result := response.Result().(*ResponseDiscoveryCreateHTTPWriteCredentials)
	return result, response, err

}

//CreateNetconfCredentials Create Netconf credentials - 1792-9bc7-465b-b564
/* Adds global netconf credentials


 */
func (s *DiscoveryService) CreateNetconfCredentials(requestDiscoveryCreateNetconfCredentials *RequestDiscoveryCreateNetconfCredentials) (*ResponseDiscoveryCreateNetconfCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/netconf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateNetconfCredentials).
		SetResult(&ResponseDiscoveryCreateNetconfCredentials{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNetconfCredentials")
	}

	result := response.Result().(*ResponseDiscoveryCreateNetconfCredentials)
	return result, response, err

}

//CreateSNMPReadCommunity Create SNMP read community - 7aa3-da9d-4e09-8ef2
/* Adds global SNMP read community


 */
func (s *DiscoveryService) CreateSNMPReadCommunity(requestDiscoveryCreateSNMPReadCommunity *RequestDiscoveryCreateSNMPReadCommunity) (*ResponseDiscoveryCreateSNMPReadCommunity, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateSNMPReadCommunity).
		SetResult(&ResponseDiscoveryCreateSNMPReadCommunity{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSnmpReadCommunity")
	}

	result := response.Result().(*ResponseDiscoveryCreateSNMPReadCommunity)
	return result, response, err

}

//CreateSNMPWriteCommunity Create SNMP write community - 6bac-b8d1-4639-bdc7
/* Adds global SNMP write community


 */
func (s *DiscoveryService) CreateSNMPWriteCommunity(requestDiscoveryCreateSNMPWriteCommunity *RequestDiscoveryCreateSNMPWriteCommunity) (*ResponseDiscoveryCreateSNMPWriteCommunity, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateSNMPWriteCommunity).
		SetResult(&ResponseDiscoveryCreateSNMPWriteCommunity{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSnmpWriteCommunity")
	}

	result := response.Result().(*ResponseDiscoveryCreateSNMPWriteCommunity)
	return result, response, err

}

//CreateSNMPv3Credentials Create SNMPv3 credentials - 9796-8808-4b7b-a60d
/* Adds global SNMPv3 credentials


 */
func (s *DiscoveryService) CreateSNMPv3Credentials(requestDiscoveryCreateSNMPv3Credentials *RequestDiscoveryCreateSNMPv3Credentials) (*ResponseDiscoveryCreateSNMPv3Credentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv3"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateSNMPv3Credentials).
		SetResult(&ResponseDiscoveryCreateSNMPv3Credentials{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSnmpv3Credentials")
	}

	result := response.Result().(*ResponseDiscoveryCreateSNMPv3Credentials)
	return result, response, err

}

//CreateUpdateSNMPProperties Create/Update SNMP properties - a5ac-9977-4c6b-b541
/* Adds SNMP properties


 */
func (s *DiscoveryService) CreateUpdateSNMPProperties(requestDiscoveryCreateUpdateSNMPProperties *RequestDiscoveryCreateUpdateSNMPProperties) (*ResponseDiscoveryCreateUpdateSNMPProperties, *resty.Response, error) {
	path := "/dna/intent/api/v1/snmp-property"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryCreateUpdateSNMPProperties).
		SetResult(&ResponseDiscoveryCreateUpdateSNMPProperties{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateUpdateSnmpProperties")
	}

	result := response.Result().(*ResponseDiscoveryCreateUpdateSNMPProperties)
	return result, response, err

}

//UpdatesAnExistingDiscoveryBySpecifiedID Updates an existing discovery by specified Id - 9788-b8fc-4418-831d
/* Stops or starts an existing discovery


 */
func (s *DiscoveryService) UpdatesAnExistingDiscoveryBySpecifiedID(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedId *RequestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID) (*ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdatesAnExistingDiscoveryBySpecifiedId).
		SetResult(&ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatesAnExistingDiscoveryBySpecifiedId")
	}

	result := response.Result().(*ResponseDiscoveryUpdatesAnExistingDiscoveryBySpecifiedID)
	return result, response, err

}

//UpdateCliCredentials Update CLI credentials - fba0-d807-47eb-82e8
/* Updates global CLI credentials


 */
func (s *DiscoveryService) UpdateCliCredentials(requestDiscoveryUpdateCLICredentials *RequestDiscoveryUpdateCliCredentials) (*ResponseDiscoveryUpdateCliCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/cli"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateCLICredentials).
		SetResult(&ResponseDiscoveryUpdateCliCredentials{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateCliCredentials")
	}

	result := response.Result().(*ResponseDiscoveryUpdateCliCredentials)
	return result, response, err

}

//UpdateHTTPReadCredential Update HTTP read credential - 89b3-6b46-4999-9d81
/* Updates global HTTP Read credential


 */
func (s *DiscoveryService) UpdateHTTPReadCredential(requestDiscoveryUpdateHTTPReadCredential *RequestDiscoveryUpdateHTTPReadCredential) (*ResponseDiscoveryUpdateHTTPReadCredential, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-read"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateHTTPReadCredential).
		SetResult(&ResponseDiscoveryUpdateHTTPReadCredential{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateHttpReadCredential")
	}

	result := response.Result().(*ResponseDiscoveryUpdateHTTPReadCredential)
	return result, response, err

}

//UpdateHTTPWriteCredentials Update HTTP write credentials - b68a-6bd8-473a-9a25
/* Updates global HTTP write credentials


 */
func (s *DiscoveryService) UpdateHTTPWriteCredentials(requestDiscoveryUpdateHTTPWriteCredentials *RequestDiscoveryUpdateHTTPWriteCredentials) (*ResponseDiscoveryUpdateHTTPWriteCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/http-write"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateHTTPWriteCredentials).
		SetResult(&ResponseDiscoveryUpdateHTTPWriteCredentials{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateHttpWriteCredentials")
	}

	result := response.Result().(*ResponseDiscoveryUpdateHTTPWriteCredentials)
	return result, response, err

}

//UpdateNetconfCredentials Update Netconf credentials - c5ac-d9fa-4c1a-8abc
/* Updates global netconf credentials


 */
func (s *DiscoveryService) UpdateNetconfCredentials(requestDiscoveryUpdateNetconfCredentials *RequestDiscoveryUpdateNetconfCredentials) (*ResponseDiscoveryUpdateNetconfCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/netconf"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateNetconfCredentials).
		SetResult(&ResponseDiscoveryUpdateNetconfCredentials{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNetconfCredentials")
	}

	result := response.Result().(*ResponseDiscoveryUpdateNetconfCredentials)
	return result, response, err

}

//UpdateSNMPReadCommunity Update SNMP read community - 47a1-b84b-4e1b-8044
/* Updates global SNMP read community


 */
func (s *DiscoveryService) UpdateSNMPReadCommunity(requestDiscoveryUpdateSNMPReadCommunity *RequestDiscoveryUpdateSNMPReadCommunity) (*ResponseDiscoveryUpdateSNMPReadCommunity, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateSNMPReadCommunity).
		SetResult(&ResponseDiscoveryUpdateSNMPReadCommunity{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSnmpReadCommunity")
	}

	result := response.Result().(*ResponseDiscoveryUpdateSNMPReadCommunity)
	return result, response, err

}

//UpdateSNMPWriteCommunity Update SNMP write community - 10b0-6a6a-4f7b-b3cb
/* Updates global SNMP write community


 */
func (s *DiscoveryService) UpdateSNMPWriteCommunity(requestDiscoveryUpdateSNMPWriteCommunity *RequestDiscoveryUpdateSNMPWriteCommunity) (*ResponseDiscoveryUpdateSNMPWriteCommunity, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateSNMPWriteCommunity).
		SetResult(&ResponseDiscoveryUpdateSNMPWriteCommunity{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSnmpWriteCommunity")
	}

	result := response.Result().(*ResponseDiscoveryUpdateSNMPWriteCommunity)
	return result, response, err

}

//UpdateSNMPv3Credentials Update SNMPv3 credentials - 1da5-ebdd-434a-acfe
/* Updates global SNMPv3 credential


 */
func (s *DiscoveryService) UpdateSNMPv3Credentials(requestDiscoveryUpdateSNMPv3Credentials *RequestDiscoveryUpdateSNMPv3Credentials) (*ResponseDiscoveryUpdateSNMPv3Credentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/snmpv3"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateSNMPv3Credentials).
		SetResult(&ResponseDiscoveryUpdateSNMPv3Credentials{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSnmpv3Credentials")
	}

	result := response.Result().(*ResponseDiscoveryUpdateSNMPv3Credentials)
	return result, response, err

}

//UpdateGlobalCredentials Update global credentials - 709f-da3c-42b8-877a
/* Update global credential for network devices in site(s)


@param globalCredentialID globalCredentialId path parameter. Global credential Uuid

*/
func (s *DiscoveryService) UpdateGlobalCredentials(globalCredentialID string, requestDiscoveryUpdateGlobalCredentials *RequestDiscoveryUpdateGlobalCredentials) (*ResponseDiscoveryUpdateGlobalCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{globalCredentialId}", fmt.Sprintf("%v", globalCredentialID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDiscoveryUpdateGlobalCredentials).
		SetResult(&ResponseDiscoveryUpdateGlobalCredentials{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateGlobalCredentials")
	}

	result := response.Result().(*ResponseDiscoveryUpdateGlobalCredentials)
	return result, response, err

}

//DeleteAllDiscovery Delete all discovery - db8e-0923-4a98-8bab
/* Stops all the discoveries and removes them


 */
func (s *DiscoveryService) DeleteAllDiscovery() (*ResponseDiscoveryDeleteAllDiscovery, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteAllDiscovery{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteAllDiscovery")
	}

	result := response.Result().(*ResponseDiscoveryDeleteAllDiscovery)
	return result, response, err

}

//DeleteDiscoveryByID Delete discovery by Id - 4c8c-ab5f-435a-80f4
/* Stops the discovery for the given Discovery ID and removes it. Discovery ID can be obtained using the "Get Discoveries by range" API.


@param id id path parameter. Discovery ID

*/
func (s *DiscoveryService) DeleteDiscoveryByID(id string) (*ResponseDiscoveryDeleteDiscoveryByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteDiscoveryByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDiscoveryById")
	}

	result := response.Result().(*ResponseDiscoveryDeleteDiscoveryByID)
	return result, response, err

}

//DeleteDiscoveryBySpecifiedRange Delete discovery by specified range - c1ba-9a42-4c08-a01b
/* Stops discovery for the given range and removes them


@param startIndex startIndex path parameter. Start index

@param recordsToDelete recordsToDelete path parameter. Number of records to delete

*/
func (s *DiscoveryService) DeleteDiscoveryBySpecifiedRange(startIndex int, recordsToDelete int) (*ResponseDiscoveryDeleteDiscoveryBySpecifiedRange, *resty.Response, error) {
	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToDelete}"
	path = strings.Replace(path, "{startIndex}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{recordsToDelete}", fmt.Sprintf("%v", recordsToDelete), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteDiscoveryBySpecifiedRange{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDiscoveryBySpecifiedRange")
	}

	result := response.Result().(*ResponseDiscoveryDeleteDiscoveryBySpecifiedRange)
	return result, response, err

}

//DeleteGlobalCredentialsByID Delete global credentials by Id - f5ac-590c-4ca9-975a
/* Deletes global credential for the given ID


@param globalCredentialID globalCredentialId path parameter. ID of global-credential

*/
func (s *DiscoveryService) DeleteGlobalCredentialsByID(globalCredentialID string) (*ResponseDiscoveryDeleteGlobalCredentialsByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{globalCredentialId}", fmt.Sprintf("%v", globalCredentialID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDiscoveryDeleteGlobalCredentialsByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteGlobalCredentialsById")
	}

	result := response.Result().(*ResponseDiscoveryDeleteGlobalCredentialsByID)
	return result, response, err

}
