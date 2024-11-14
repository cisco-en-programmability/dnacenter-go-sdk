package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type NetworkSettingsService service

type AssignDeviceCredentialToSiteHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Persist bapi sync response
}
type GetDeviceCredentialDetailsQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site id to retrieve the credential details associated with the site.
}
type GetGlobalPoolQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //Offset/starting row. Indexed from 1. Default value of 1.
	Limit  float64 `url:"limit,omitempty"`  //Number of Global Pools to be retrieved. Default is 25 if not specified.
}
type GetNetworkQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site id to get the network settings associated with the site.
}
type CreateNetworkHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetReserveIPSubpoolQueryParams struct {
	SiteID                string  `url:"siteId,omitempty"`                //site id of site from which to retrieve associated reserve pools. Either siteId (per site queries) or ignoreInheritedGroups must be used. They can also be used together.
	Offset                float64 `url:"offset,omitempty"`                //offset/starting row. Indexed from 1.
	Limit                 float64 `url:"limit,omitempty"`                 //Number of reserve pools to be retrieved. Default is 25 if not specified. Maximum allowed limit is 500.
	IgnoreInheritedGroups string  `url:"ignoreInheritedGroups,omitempty"` //Ignores pools inherited from parent site. Either siteId or ignoreInheritedGroups must be passed. They can also be used together.
	PoolUsage             string  `url:"poolUsage,omitempty"`             //Can take values empty, partially-full or empty-partially-full
	GroupName             string  `url:"groupName,omitempty"`             //Name of the group
}
type UpdateReserveIPSubpoolQueryParams struct {
	ID string `url:"id,omitempty"` //Id of subpool group
}
type RetrieveAAASettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveBannerSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type GetDeviceCredentialSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveDHCPSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveDNSSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveImageDistributionSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveNTPSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveTelemetrySettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type RetrieveTimeZoneSettingsForASiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when `false`, `null` values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
}
type GetNetworkV2QueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site Id to get the network settings associated with the site.
}

type ResponseNetworkSettingsAssignDeviceCredentialToSite struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsCreateDeviceCredentials struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsUpdateDeviceCredentials struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetDeviceCredentialDetails struct {
	SNMPV3      *[]ResponseNetworkSettingsGetDeviceCredentialDetailsSNMPV3      `json:"snmp_v3,omitempty"`       //
	HTTPRead    *[]ResponseNetworkSettingsGetDeviceCredentialDetailsHTTPRead    `json:"http_read,omitempty"`     //
	HTTPWrite   *[]ResponseNetworkSettingsGetDeviceCredentialDetailsHTTPWrite   `json:"http_write,omitempty"`    //
	SNMPV2Write *[]ResponseNetworkSettingsGetDeviceCredentialDetailsSNMPV2Write `json:"snmp_v2_write,omitempty"` //
	SNMPV2Read  *[]ResponseNetworkSettingsGetDeviceCredentialDetailsSNMPV2Read  `json:"snmp_v2_read,omitempty"`  //
	Cli         *[]ResponseNetworkSettingsGetDeviceCredentialDetailsCli         `json:"cli,omitempty"`           //
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsSNMPV3 struct {
	Username         string `json:"username,omitempty"`         // Username
	AuthPassword     string `json:"authPassword,omitempty"`     // Auth Password
	AuthType         string `json:"authType,omitempty"`         // Auth Type
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  // Privacy Password
	PrivacyType      string `json:"privacyType,omitempty"`      // Privacy Type
	SNMPMode         string `json:"snmpMode,omitempty"`         // Snmp Mode
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsHTTPRead struct {
	Secure           string `json:"secure,omitempty"`           // Secure
	Username         string `json:"username,omitempty"`         // Username
	Password         string `json:"password,omitempty"`         // Password
	Port             string `json:"port,omitempty"`             // Port
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsHTTPWrite struct {
	Secure           string `json:"secure,omitempty"`           // Secure
	Username         string `json:"username,omitempty"`         // Username
	Password         string `json:"password,omitempty"`         // Password
	Port             string `json:"port,omitempty"`             // Port
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsSNMPV2Write struct {
	WriteCommunity   string `json:"writeCommunity,omitempty"`   // Write Community
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsSNMPV2Read struct {
	ReadCommunity    string `json:"readCommunity,omitempty"`    // Read Community
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsGetDeviceCredentialDetailsCli struct {
	Username         string `json:"username,omitempty"`         // Username
	EnablePassword   string `json:"enablePassword,omitempty"`   // Enable Password
	Password         string `json:"password,omitempty"`         // Password
	Comments         string `json:"comments,omitempty"`         // Comments
	Description      string `json:"description,omitempty"`      // Description
	CredentialType   string `json:"credentialType,omitempty"`   // Credential Type
	InstanceUUID     string `json:"instanceUuid,omitempty"`     // Instance Uuid
	InstanceTenantID string `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	ID               string `json:"id,omitempty"`               // Id
}
type ResponseNetworkSettingsDeleteDeviceCredential struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetGlobalPool struct {
	Response *[]ResponseNetworkSettingsGetGlobalPoolResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetGlobalPoolResponse struct {
	IPPoolName                    string                                                     `json:"ipPoolName,omitempty"`                    // Ip Pool Name
	DhcpServerIPs                 []string                                                   `json:"dhcpServerIps,omitempty"`                 // Dhcp Server Ips
	Gateways                      []string                                                   `json:"gateways,omitempty"`                      // Gateways
	CreateTime                    *int                                                       `json:"createTime,omitempty"`                    // Create Time
	LastUpdateTime                *int                                                       `json:"lastUpdateTime,omitempty"`                // Last Update Time
	TotalIPAddressCount           *int                                                       `json:"totalIpAddressCount,omitempty"`           // Total Ip Address Count
	UsedIPAddressCount            *int                                                       `json:"usedIpAddressCount,omitempty"`            // Used Ip Address Count
	ParentUUID                    string                                                     `json:"parentUuid,omitempty"`                    // Parent Uuid
	Owner                         string                                                     `json:"owner,omitempty"`                         // Owner
	Shared                        *bool                                                      `json:"shared,omitempty"`                        // Shared
	Overlapping                   *bool                                                      `json:"overlapping,omitempty"`                   // Overlapping
	ConfigureExternalDhcp         *bool                                                      `json:"configureExternalDhcp,omitempty"`         // Configure External Dhcp
	UsedPercentage                string                                                     `json:"usedPercentage,omitempty"`                // Used Percentage
	ClientOptions                 *ResponseNetworkSettingsGetGlobalPoolResponseClientOptions `json:"clientOptions,omitempty"`                 // Client Options
	IPPoolType                    string                                                     `json:"ipPoolType,omitempty"`                    // Ip Pool Type
	UnavailableIPAddressCount     *float64                                                   `json:"unavailableIpAddressCount,omitempty"`     // Unavailable Ip Address Count
	AvailableIPAddressCount       *float64                                                   `json:"availableIpAddressCount,omitempty"`       // Available Ip Address Count
	TotalAssignableIPAddressCount *int                                                       `json:"totalAssignableIpAddressCount,omitempty"` // Total Assignable Ip Address Count
	DNSServerIPs                  []string                                                   `json:"dnsServerIps,omitempty"`                  // Dns Server Ips
	HasSubpools                   *bool                                                      `json:"hasSubpools,omitempty"`                   // Has Subpools
	DefaultAssignedIPAddressCount *int                                                       `json:"defaultAssignedIpAddressCount,omitempty"` // Default Assigned Ip Address Count
	Context                       *[]ResponseNetworkSettingsGetGlobalPoolResponseContext     `json:"context,omitempty"`                       //
	IPv6                          *bool                                                      `json:"ipv6,omitempty"`                          // Ipv6
	ID                            string                                                     `json:"id,omitempty"`                            // Id
	IPPoolCidr                    string                                                     `json:"ipPoolCidr,omitempty"`                    // Ip Pool Cidr
}
type ResponseNetworkSettingsGetGlobalPoolResponseClientOptions interface{}
type ResponseNetworkSettingsGetGlobalPoolResponseContext struct {
	Owner        string `json:"owner,omitempty"`        // Owner
	ContextKey   string `json:"contextKey,omitempty"`   // Context Key
	ContextValue string `json:"contextValue,omitempty"` // Context Value
}
type ResponseNetworkSettingsUpdateGlobalPool struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsCreateGlobalPool struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsDeleteGlobalIPPool struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetNetwork struct {
	Response *[]ResponseNetworkSettingsGetNetworkResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetNetworkResponse struct {
	InstanceType       string                                            `json:"instanceType,omitempty"`       // Instance Type
	InstanceUUID       string                                            `json:"instanceUuid,omitempty"`       // Instance Uuid
	Namespace          string                                            `json:"namespace,omitempty"`          // Namespace
	Type               string                                            `json:"type,omitempty"`               // Type
	Key                string                                            `json:"key,omitempty"`                // Key
	Version            *int                                              `json:"version,omitempty"`            // Version
	Value              *[]ResponseNetworkSettingsGetNetworkResponseValue `json:"value,omitempty"`              //
	GroupUUID          string                                            `json:"groupUuid,omitempty"`          // Group Uuid
	InheritedGroupUUID string                                            `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid
	InheritedGroupName string                                            `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsGetNetworkResponseValue struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // Ip Addresses
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configure Dnac I P
}
type ResponseNetworkSettingsCreateNetwork struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsUpdateNetwork struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetReserveIPSubpool struct {
	Response *[]ResponseNetworkSettingsGetReserveIPSubpoolResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetReserveIPSubpoolResponse struct {
	ID            string                                                       `json:"id,omitempty"`            // Id
	GroupName     string                                                       `json:"groupName,omitempty"`     // Group Name
	IPPools       *[]ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPools `json:"ipPools,omitempty"`       //
	SiteID        string                                                       `json:"siteId,omitempty"`        // Site Id
	SiteHierarchy string                                                       `json:"siteHierarchy,omitempty"` // Site Hierarchy
	Type          string                                                       `json:"type,omitempty"`          // Type
	GroupOwner    string                                                       `json:"groupOwner,omitempty"`    // Group Owner
}
type ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPools struct {
	IPPoolName            string                                                                    `json:"ipPoolName,omitempty"`            // Ip Pool Name
	DhcpServerIPs         *[]ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsDhcpServerIPs `json:"dhcpServerIps,omitempty"`         // Dhcp Server Ips
	Gateways              []string                                                                  `json:"gateways,omitempty"`              // Gateways
	CreateTime            *int                                                                      `json:"createTime,omitempty"`            // Create Time
	LastUpdateTime        *int                                                                      `json:"lastUpdateTime,omitempty"`        // Last Update Time
	TotalIPAddressCount   *int                                                                      `json:"totalIpAddressCount,omitempty"`   // Total Ip Address Count
	UsedIPAddressCount    *int                                                                      `json:"usedIpAddressCount,omitempty"`    // Used Ip Address Count
	ParentUUID            string                                                                    `json:"parentUuid,omitempty"`            // Parent Uuid
	Owner                 string                                                                    `json:"owner,omitempty"`                 // Owner
	Shared                *bool                                                                     `json:"shared,omitempty"`                // Shared
	Overlapping           *bool                                                                     `json:"overlapping,omitempty"`           // Overlapping
	ConfigureExternalDhcp *bool                                                                     `json:"configureExternalDhcp,omitempty"` // Configure External Dhcp
	UsedPercentage        string                                                                    `json:"usedPercentage,omitempty"`        // Used Percentage
	ClientOptions         *ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsClientOptions   `json:"clientOptions,omitempty"`         // Client Options
	GroupUUID             string                                                                    `json:"groupUuid,omitempty"`             // Group Uuid
	DNSServerIPs          *[]ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsDNSServerIPs  `json:"dnsServerIps,omitempty"`          // Dns Server Ips
	Context               *[]ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsContext       `json:"context,omitempty"`               //
	IPv6                  *bool                                                                     `json:"ipv6,omitempty"`                  // Ipv6
	ID                    string                                                                    `json:"id,omitempty"`                    // Id
	IPPoolCidr            string                                                                    `json:"ipPoolCidr,omitempty"`            // Ip Pool Cidr
}
type ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsDhcpServerIPs interface{}
type ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsClientOptions interface{}
type ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsDNSServerIPs interface{}
type ResponseNetworkSettingsGetReserveIPSubpoolResponseIPPoolsContext struct {
	Owner        string `json:"owner,omitempty"`        // Owner
	ContextKey   string `json:"contextKey,omitempty"`   // Context Key
	ContextValue string `json:"contextValue,omitempty"` // Context Value
}
type ResponseNetworkSettingsReleaseReserveIPSubpool struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsReserveIPSubpool struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsUpdateReserveIPSubpool struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsGetServiceProviderDetails struct {
	Response *[]ResponseNetworkSettingsGetServiceProviderDetailsResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetServiceProviderDetailsResponse struct {
	InstanceType       string                                                           `json:"instanceType,omitempty"`       // Instance Type
	InstanceUUID       string                                                           `json:"instanceUuid,omitempty"`       // Instance Uuid
	Namespace          string                                                           `json:"namespace,omitempty"`          // Namespace
	Type               string                                                           `json:"type,omitempty"`               // Type
	Key                string                                                           `json:"key,omitempty"`                // Key
	Version            *int                                                             `json:"version,omitempty"`            // Version
	Value              *[]ResponseNetworkSettingsGetServiceProviderDetailsResponseValue `json:"value,omitempty"`              //
	GroupUUID          string                                                           `json:"groupUuid,omitempty"`          // Group Uuid
	InheritedGroupUUID string                                                           `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid
	InheritedGroupName string                                                           `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsGetServiceProviderDetailsResponseValue struct {
	WanProvider    string `json:"wanProvider,omitempty"`    // Wan Provider
	SpProfileName  string `json:"spProfileName,omitempty"`  // Sp Profile Name
	SLAProfileName string `json:"slaProfileName,omitempty"` // Sla Profile Name
}
type ResponseNetworkSettingsCreateSpProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsUpdateSpProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsSyncNetworkDevicesCredential struct {
	Version  string                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSyncNetworkDevicesCredentialResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSyncNetworkDevicesCredentialResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsSetAAASettingsForASite struct {
	Version  string                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetAAASettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetAAASettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveAAASettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveAAASettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteResponse struct {
	AAANetwork *ResponseNetworkSettingsRetrieveAAASettingsForASiteResponseAAANetwork `json:"aaaNetwork,omitempty"` //
	AAAClient  *ResponseNetworkSettingsRetrieveAAASettingsForASiteResponseAAAClient  `json:"aaaClient,omitempty"`  //
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteResponseAAANetwork struct {
	ServerType        string `json:"serverType,omitempty"`        // Server Type
	Protocol          string `json:"protocol,omitempty"`          // Protocol
	Pan               string `json:"pan,omitempty"`               // Administration Node. Required for ISE.
	PrimaryServerIP   string `json:"primaryServerIp,omitempty"`   // The server to use as a primary.
	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.
	SharedSecret      string `json:"sharedSecret,omitempty"`      // Shared Secret
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveAAASettingsForASiteResponseAAAClient struct {
	ServerType        string `json:"serverType,omitempty"`        // Server Type
	Protocol          string `json:"protocol,omitempty"`          // Protocol
	Pan               string `json:"pan,omitempty"`               // Administration Node. Required for ISE.
	PrimaryServerIP   string `json:"primaryServerIp,omitempty"`   // The server to use as a primary.
	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.
	SharedSecret      string `json:"sharedSecret,omitempty"`      // Shared Secret
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveBannerSettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveBannerSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveBannerSettingsForASiteResponse struct {
	Banner *ResponseNetworkSettingsRetrieveBannerSettingsForASiteResponseBanner `json:"banner,omitempty"` //
}
type ResponseNetworkSettingsRetrieveBannerSettingsForASiteResponseBanner struct {
	Type              string `json:"type,omitempty"`              // Type
	Message           string `json:"message,omitempty"`           // Custom message that appears when logging into routers, switches, and hubs. Required for custom type.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsSetBannerSettingsForASite struct {
	Version  string                                                    `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetBannerSettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetBannerSettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASite struct {
	Response *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponse struct {
	CliCredentialsID          *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseCliCredentialsID          `json:"cliCredentialsId,omitempty"`          //
	SNMPv2CReadCredentialsID  *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv2CReadCredentialsID  `json:"snmpv2cReadCredentialsId,omitempty"`  //
	SNMPv2CWriteCredentialsID *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv2CWriteCredentialsID `json:"snmpv2cWriteCredentialsId,omitempty"` //
	SNMPv3CredentialsID       *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv3CredentialsID       `json:"snmpv3CredentialsId,omitempty"`       //
	HTTPReadCredentialsID     *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseHTTPReadCredentialsID     `json:"httpReadCredentialsId,omitempty"`     //
	HTTPWriteCredentialsID    *ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseHTTPWriteCredentialsID    `json:"httpWriteCredentialsId,omitempty"`    //
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseCliCredentialsID struct {
	CredentialsID     string `json:"credentialsId,omitempty"`     // The `id` of the credentials.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv2CReadCredentialsID struct {
	CredentialsID     string `json:"credentialsId,omitempty"`     // The `id` of the credentials.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv2CWriteCredentialsID struct {
	CredentialsID     string `json:"credentialsId,omitempty"`     // The `id` of the credentials.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseSNMPv3CredentialsID struct {
	CredentialsID     string `json:"credentialsId,omitempty"`     // The `id` of the credentials.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseHTTPReadCredentialsID struct {
	CredentialsID     string `json:"credentialsId,omitempty"`     // The `id` of the credentials.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteResponseHTTPWriteCredentialsID struct {
	CredentialsID     string `json:"credentialsId,omitempty"`     // The `id` of the credentials.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name
}
type ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASite struct {
	Version  string                                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatus struct {
	Response *ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  //
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponse struct {
	Cli         *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseCli         `json:"cli,omitempty"`         //
	SNMPV2Read  *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV2Read  `json:"snmpV2Read,omitempty"`  //
	SNMPV2Write *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV2Write `json:"snmpV2Write,omitempty"` //
	SNMPV3      *[]ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV3      `json:"snmpV3,omitempty"`      //
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseCli struct {
	DeviceCount *int   `json:"deviceCount,omitempty"` // Device count
	Status      string `json:"status,omitempty"`      // Sync status
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV2Read struct {
	DeviceCount *int   `json:"deviceCount,omitempty"` // Device count
	Status      string `json:"status,omitempty"`      // Sync status
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV2Write struct {
	DeviceCount *int   `json:"deviceCount,omitempty"` // Device count
	Status      string `json:"status,omitempty"`      // Sync status
}
type ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatusResponseSNMPV3 struct {
	DeviceCount *int   `json:"deviceCount,omitempty"` // Device count
	Status      string `json:"status,omitempty"`      // Sync status
}
type ResponseNetworkSettingsSetDhcpSettingsForASite struct {
	Version  string                                                  `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetDhcpSettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetDhcpSettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveDHCPSettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveDHCPSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveDHCPSettingsForASiteResponse struct {
	Dhcp *ResponseNetworkSettingsRetrieveDHCPSettingsForASiteResponseDhcp `json:"dhcp,omitempty"` //
}
type ResponseNetworkSettingsRetrieveDHCPSettingsForASiteResponseDhcp struct {
	Servers           []string `json:"servers,omitempty"`           // DHCP servers for managing client device networking configuration.
	InheritedSiteID   string   `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string   `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveDNSSettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveDNSSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveDNSSettingsForASiteResponse struct {
	DNS *ResponseNetworkSettingsRetrieveDNSSettingsForASiteResponseDNS `json:"dns,omitempty"` //
}
type ResponseNetworkSettingsRetrieveDNSSettingsForASiteResponseDNS struct {
	DomainName        string   `json:"domainName,omitempty"`        // Network's domain name.
	DNSServers        []string `json:"dnsServers,omitempty"`        // DNS servers for hostname resolution.
	InheritedSiteID   string   `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string   `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsSetDNSSettingsForASite struct {
	Version  string                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetDNSSettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetDNSSettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsSetImageDistributionSettingsForASite struct {
	Version  string                                                               `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetImageDistributionSettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetImageDistributionSettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveImageDistributionSettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                                    `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteResponse struct {
	ImageDistribution *ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteResponseImageDistribution `json:"imageDistribution,omitempty"` //
}
type ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteResponseImageDistribution struct {
	Servers           []string `json:"servers,omitempty"`           // This field holds an array of unique identifiers representing image distribution servers. SFTP servers to act as image distribution servers. A distributed SWIM architecture, using suitably located SFTP servers, can help support large-scale device software image upgrades and conserve WAN bandwidth.
	InheritedSiteID   string   `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string   `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsSetNTPSettingsForASite struct {
	Version  string                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetNTPSettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetNTPSettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveNTPSettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveNTPSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveNTPSettingsForASiteResponse struct {
	Ntp *ResponseNetworkSettingsRetrieveNTPSettingsForASiteResponseNtp `json:"ntp,omitempty"` //
}
type ResponseNetworkSettingsRetrieveNTPSettingsForASiteResponseNtp struct {
	Servers           []string `json:"servers,omitempty"`           // NTP servers to facilitate system clock synchronization for your network.
	InheritedSiteID   string   `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string   `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponse struct {
	WiredDataCollection   *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseWiredDataCollection   `json:"wiredDataCollection,omitempty"`   //
	WirelessTelemetry     *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseWirelessTelemetry     `json:"wirelessTelemetry,omitempty"`     //
	SNMPTraps             *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseSNMPTraps             `json:"snmpTraps,omitempty"`             //
	Syslogs               *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseSyslogs               `json:"syslogs,omitempty"`               //
	ApplicationVisibility *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseApplicationVisibility `json:"applicationVisibility,omitempty"` //
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseWiredDataCollection struct {
	EnableWiredDataCollectio *bool  `json:"enableWiredDataCollectio,omitempty"` // Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
	InheritedSiteID          string `json:"inheritedSiteId,omitempty"`          // Inherited Site Id
	InheritedSiteName        string `json:"inheritedSiteName,omitempty"`        // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseWirelessTelemetry struct {
	EnableWirelessTelemetry *bool  `json:"enableWirelessTelemetry,omitempty"` // Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
	InheritedSiteID         string `json:"inheritedSiteId,omitempty"`         // Inherited Site Id
	InheritedSiteName       string `json:"inheritedSiteName,omitempty"`       // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseSNMPTraps struct {
	UseBuiltinTrapServer *bool    `json:"useBuiltinTrapServer,omitempty"` // Enable this server as a destination server for SNMP traps and messages from your network
	ExternalTrapServers  []string `json:"externalTrapServers,omitempty"`  // External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
	InheritedSiteID      string   `json:"inheritedSiteId,omitempty"`      // Inherited Site Id
	InheritedSiteName    string   `json:"inheritedSiteName,omitempty"`    // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseSyslogs struct {
	UseBuiltinSyslogServer *bool    `json:"useBuiltinSyslogServer,omitempty"` // Enable this server as a destination server for syslog messages.
	ExternalSyslogServers  []string `json:"externalSyslogServers,omitempty"`  // External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
	InheritedSiteID        string   `json:"inheritedSiteId,omitempty"`        // Inherited Site Id
	InheritedSiteName      string   `json:"inheritedSiteName,omitempty"`      // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseApplicationVisibility struct {
	Collector                  *ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseApplicationVisibilityCollector `json:"collector,omitempty"`                  //
	EnableOnWiredAccessDevices *bool                                                                                           `json:"enableOnWiredAccessDevices,omitempty"` // Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
	InheritedSiteID            string                                                                                          `json:"inheritedSiteId,omitempty"`            // Inherited Site Id
	InheritedSiteName          string                                                                                          `json:"inheritedSiteName,omitempty"`          // Inherited Site Name
}
type ResponseNetworkSettingsRetrieveTelemetrySettingsForASiteResponseApplicationVisibilityCollector struct {
	CollectorType string `json:"collectorType,omitempty"` // Collector Type
	Address       string `json:"address,omitempty"`       // IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"
	Port          *int   `json:"port,omitempty"`          // Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
}
type ResponseNetworkSettingsSetTelemetrySettingsForASite struct {
	Version  string                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetTelemetrySettingsForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetTelemetrySettingsForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsSetTimeZoneForASite struct {
	Version  string                                              `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsSetTimeZoneForASiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsSetTimeZoneForASiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsRetrieveTimeZoneSettingsForASite struct {
	Response *ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteResponse struct {
	TimeZone *ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteResponseTimeZone `json:"timeZone,omitempty"` //
}
type ResponseNetworkSettingsRetrieveTimeZoneSettingsForASiteResponseTimeZone struct {
	IDentifier        string `json:"identifier,omitempty"`        // Time zone that corresponds to the site's physical location. The site time zone is used when scheduling device provisioning and updates. Example : GMT
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id.
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name.
}
type ResponseNetworkSettingsDeleteSpProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite struct {
	Version  string                                                                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteResponse `json:"response,omitempty"` //
}
type ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSiteResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseNetworkSettingsAssignDeviceCredentialToSiteV2 struct {
	Response *ResponseNetworkSettingsAssignDeviceCredentialToSiteV2Response `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsAssignDeviceCredentialToSiteV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseNetworkSettingsGetNetworkV2 struct {
	Response *[]ResponseNetworkSettingsGetNetworkV2Response `json:"response,omitempty"` //
}
type ResponseNetworkSettingsGetNetworkV2Response struct {
	InstanceType       string   `json:"instanceType,omitempty"`       // Instance Type
	InstanceUUID       string   `json:"instanceUuid,omitempty"`       // Instance Uuid
	Namespace          string   `json:"namespace,omitempty"`          // Namespace
	Type               string   `json:"type,omitempty"`               // Type
	Key                string   `json:"key,omitempty"`                // Key
	Version            *int     `json:"version,omitempty"`            // Version
	Value              []string `json:"value,omitempty"`              // Value
	GroupUUID          string   `json:"groupUuid,omitempty"`          // Group Uuid
	InheritedGroupUUID string   `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid
	InheritedGroupName string   `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsCreateNetworkV2 struct {
	Response *ResponseNetworkSettingsCreateNetworkV2Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsCreateNetworkV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseNetworkSettingsUpdateNetworkV2 struct {
	Response *ResponseNetworkSettingsUpdateNetworkV2Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsUpdateNetworkV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseNetworkSettingsCreateSpProfileV2 struct {
	Response *ResponseNetworkSettingsCreateSpProfileV2Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsCreateSpProfileV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseNetworkSettingsUpdateSpProfileV2 struct {
	Response *ResponseNetworkSettingsUpdateSpProfileV2Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsUpdateSpProfileV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseNetworkSettingsGetServiceProviderDetailsV2 struct {
	Response *[]ResponseNetworkSettingsGetServiceProviderDetailsV2Response `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetServiceProviderDetailsV2Response struct {
	InstanceType       string                                                             `json:"instanceType,omitempty"`       // Instance Type
	InstanceUUID       string                                                             `json:"instanceUuid,omitempty"`       // Instance Uuid
	Namespace          string                                                             `json:"namespace,omitempty"`          // Namespace
	Type               string                                                             `json:"type,omitempty"`               // Type
	Key                string                                                             `json:"key,omitempty"`                // Key
	Version            *int                                                               `json:"version,omitempty"`            // Version
	Value              *[]ResponseNetworkSettingsGetServiceProviderDetailsV2ResponseValue `json:"value,omitempty"`              //
	GroupUUID          string                                                             `json:"groupUuid,omitempty"`          // Group Uuid
	InheritedGroupUUID string                                                             `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid
	InheritedGroupName string                                                             `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseNetworkSettingsGetServiceProviderDetailsV2ResponseValue struct {
	WanProvider    string `json:"wanProvider,omitempty"`    // Wan Provider
	SpProfileName  string `json:"spProfileName,omitempty"`  // Sp Profile Name
	SLAProfileName string `json:"slaProfileName,omitempty"` // Sla Profile Name
}
type ResponseNetworkSettingsDeleteSpProfileV2 struct {
	Response *ResponseNetworkSettingsDeleteSpProfileV2Response `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsDeleteSpProfileV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type RequestNetworkSettingsAssignDeviceCredentialToSite struct {
	CliID         string `json:"cliId,omitempty"`         // Cli Id
	SNMPV2ReadID  string `json:"snmpV2ReadId,omitempty"`  // Snmp V2 Read Id
	SNMPV2WriteID string `json:"snmpV2WriteId,omitempty"` // Snmp V2 Write Id
	HTTPRead      string `json:"httpRead,omitempty"`      // Http Read
	HTTPWrite     string `json:"httpWrite,omitempty"`     // Http Write
	SNMPV3ID      string `json:"snmpV3Id,omitempty"`      // Snmp V3 Id
}
type RequestNetworkSettingsCreateDeviceCredentials struct {
	Settings *RequestNetworkSettingsCreateDeviceCredentialsSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateDeviceCredentialsSettings struct {
	CliCredential *[]RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential `json:"cliCredential,omitempty"` //
	SNMPV2CRead   *[]RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2CWrite  *[]RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite  `json:"snmpV2cWrite,omitempty"`  //
	SNMPV3        *[]RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3        `json:"snmpV3,omitempty"`        //
	HTTPSRead     *[]RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    *[]RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite    `json:"httpsWrite,omitempty"`    //
}
type RequestNetworkSettingsCreateDeviceCredentialsSettingsCliCredential struct {
	Description    string `json:"description,omitempty"`    // Name or description for CLI credential
	Username       string `json:"username,omitempty"`       // User name for CLI credential
	Password       string `json:"password,omitempty"`       // Password for CLI credential
	EnablePassword string `json:"enablePassword,omitempty"` // Enable password for CLI credential
}
type RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CRead struct {
	Description   string `json:"description,omitempty"`   // Description for snmp v2 read
	ReadCommunity string `json:"readCommunity,omitempty"` // Ready community for snmp v2 read credential
}
type RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV2CWrite struct {
	Description    string `json:"description,omitempty"`    // Description for snmp v2 write
	WriteCommunity string `json:"writeCommunity,omitempty"` // Write community for snmp v2 write credential
}
type RequestNetworkSettingsCreateDeviceCredentialsSettingsSNMPV3 struct {
	Description     string `json:"description,omitempty"`     // Name or description for SNMPV3 credential
	Username        string `json:"username,omitempty"`        // User name for SNMPv3 credential
	PrivacyType     string `json:"privacyType,omitempty"`     // Privacy type for snmpv3 credential
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy password for snmpv3 credential
	AuthType        string `json:"authType,omitempty"`        // Authentication type for snmpv3 credential
	AuthPassword    string `json:"authPassword,omitempty"`    // Authentication password for snmpv3 credential
	SNMPMode        string `json:"snmpMode,omitempty"`        // Mode for snmpv3 credential
}
type RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSRead struct {
	Name     string   `json:"name,omitempty"`     // Name or description of http read credential
	Username string   `json:"username,omitempty"` // User name of the http read credential
	Password string   `json:"password,omitempty"` // Password for http read credential
	Port     *float64 `json:"port,omitempty"`     // Port for http read credential
}
type RequestNetworkSettingsCreateDeviceCredentialsSettingsHTTPSWrite struct {
	Name     string   `json:"name,omitempty"`     // Name or description of http write credential
	Username string   `json:"username,omitempty"` // User name of the http write credential
	Password string   `json:"password,omitempty"` // Password for http write credential
	Port     *float64 `json:"port,omitempty"`     // Port for http write credential
}
type RequestNetworkSettingsUpdateDeviceCredentials struct {
	Settings *RequestNetworkSettingsUpdateDeviceCredentialsSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettings struct {
	CliCredential *RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential `json:"cliCredential,omitempty"` //
	SNMPV2CRead   *RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2CWrite  *RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite  `json:"snmpV2cWrite,omitempty"`  //
	SNMPV3        *RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3        `json:"snmpV3,omitempty"`        //
	HTTPSRead     *RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    *RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite    `json:"httpsWrite,omitempty"`    //
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettingsCliCredential struct {
	Description    string `json:"description,omitempty"`    // Description
	Username       string `json:"username,omitempty"`       // Username
	Password       string `json:"password,omitempty"`       // Password
	EnablePassword string `json:"enablePassword,omitempty"` // Enable Password
	ID             string `json:"id,omitempty"`             // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CRead struct {
	Description   string `json:"description,omitempty"`   // Description
	ReadCommunity string `json:"readCommunity,omitempty"` // Read Community
	ID            string `json:"id,omitempty"`            // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV2CWrite struct {
	Description    string `json:"description,omitempty"`    // Description
	WriteCommunity string `json:"writeCommunity,omitempty"` // Write Community
	ID             string `json:"id,omitempty"`             // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettingsSNMPV3 struct {
	AuthPassword    string `json:"authPassword,omitempty"`    // Auth Password
	AuthType        string `json:"authType,omitempty"`        // Auth Type
	SNMPMode        string `json:"snmpMode,omitempty"`        // Snmp Mode
	PrivacyPassword string `json:"privacyPassword,omitempty"` // Privacy Password
	PrivacyType     string `json:"privacyType,omitempty"`     // Privacy Type
	Username        string `json:"username,omitempty"`        // Username
	Description     string `json:"description,omitempty"`     // Description
	ID              string `json:"id,omitempty"`              // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSRead struct {
	Name     string `json:"name,omitempty"`     // Name
	Username string `json:"username,omitempty"` // Username
	Password string `json:"password,omitempty"` // Password
	Port     string `json:"port,omitempty"`     // Port
	ID       string `json:"id,omitempty"`       // Id
}
type RequestNetworkSettingsUpdateDeviceCredentialsSettingsHTTPSWrite struct {
	Name     string `json:"name,omitempty"`     // Name
	Username string `json:"username,omitempty"` // Username
	Password string `json:"password,omitempty"` // Password
	Port     string `json:"port,omitempty"`     // Port
	ID       string `json:"id,omitempty"`       // Id
}
type RequestNetworkSettingsUpdateGlobalPool struct {
	Settings *RequestNetworkSettingsUpdateGlobalPoolSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateGlobalPoolSettings struct {
	IPpool *[]RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool `json:"ippool,omitempty"` //
}
type RequestNetworkSettingsUpdateGlobalPoolSettingsIPpool struct {
	IPPoolName    string   `json:"ipPoolName,omitempty"`    // Ip Pool Name
	Gateway       string   `json:"gateway,omitempty"`       // Gateway
	DhcpServerIPs []string `json:"dhcpServerIps,omitempty"` // Dhcp Server Ips
	DNSServerIPs  []string `json:"dnsServerIps,omitempty"`  // Dns Server Ips
	ID            string   `json:"id,omitempty"`            // Id
}
type RequestNetworkSettingsCreateGlobalPool struct {
	Settings *RequestNetworkSettingsCreateGlobalPoolSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateGlobalPoolSettings struct {
	IPpool *[]RequestNetworkSettingsCreateGlobalPoolSettingsIPpool `json:"ippool,omitempty"` //
}
type RequestNetworkSettingsCreateGlobalPoolSettingsIPpool struct {
	IPPoolName     string   `json:"ipPoolName,omitempty"`     // Ip Pool Name
	Type           string   `json:"type,omitempty"`           // Type
	IPPoolCidr     string   `json:"ipPoolCidr,omitempty"`     // Ip Pool Cidr
	Gateway        string   `json:"gateway,omitempty"`        // Gateway
	DhcpServerIPs  []string `json:"dhcpServerIps,omitempty"`  // Dhcp Server Ips
	DNSServerIPs   []string `json:"dnsServerIps,omitempty"`   // Dns Server Ips
	IPAddressSpace string   `json:"IpAddressSpace,omitempty"` // Ip Address Space
}
type RequestNetworkSettingsCreateNetwork struct {
	Settings *RequestNetworkSettingsCreateNetworkSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkSettings struct {
	DhcpServer           []string                                                         `json:"dhcpServer,omitempty"`            // DHCP Server IP (eg: 1.1.1.1)
	DNSServer            *RequestNetworkSettingsCreateNetworkSettingsDNSServer            `json:"dnsServer,omitempty"`             //
	SyslogServer         *RequestNetworkSettingsCreateNetworkSettingsSyslogServer         `json:"syslogServer,omitempty"`          //
	SNMPServer           *RequestNetworkSettingsCreateNetworkSettingsSNMPServer           `json:"snmpServer,omitempty"`            //
	Netflowcollector     *RequestNetworkSettingsCreateNetworkSettingsNetflowcollector     `json:"netflowcollector,omitempty"`      //
	NtpServer            []string                                                         `json:"ntpServer,omitempty"`             // IP address for NTP server (eg: 1.1.1.2)
	Timezone             string                                                           `json:"timezone,omitempty"`              // Input for time zone (eg: Africa/Abidjan)
	MessageOfTheday      *RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday      `json:"messageOfTheday,omitempty"`       //
	NetworkAAA           *RequestNetworkSettingsCreateNetworkSettingsNetworkAAA           `json:"network_aaa,omitempty"`           //
	ClientAndEndpointAAA *RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkSettingsDNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         // Domain Name of DHCP (eg; cisco)
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   // Primary IP Address for DHCP (eg: 2.2.2.2)
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsCreateNetworkSettingsSyslogServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for syslog server (eg: 4.4.4.4)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsCreateNetworkSettingsSNMPServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for SNMP Server (eg: 4.4.4.1)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsCreateNetworkSettingsNetflowcollector struct {
	IPAddress string   `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)
	Port      *float64 `json:"port,omitempty"`      // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsCreateNetworkSettingsMessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        // Massage for Banner message (eg; Good day)
	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsCreateNetworkSettingsNetworkAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type for AAA Network (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for AAA and ISE server (eg: 1.1.1.1)
	Network      string `json:"network,omitempty"`      // IP Address for AAA or ISE server (eg: 2.2.2.2)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsCreateNetworkSettingsClientAndEndpointAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type AAA or ISE server (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for ISE serve (eg: 1.1.1.4)
	Network      string `json:"network,omitempty"`      // IP address for AAA or ISE server (eg: 2.2.2.1)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsUpdateNetwork struct {
	Settings *RequestNetworkSettingsUpdateNetworkSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkSettings struct {
	DhcpServer           []string                                                         `json:"dhcpServer,omitempty"`            // DHCP Server IP (eg: 1.1.1.1)
	DNSServer            *RequestNetworkSettingsUpdateNetworkSettingsDNSServer            `json:"dnsServer,omitempty"`             //
	SyslogServer         *RequestNetworkSettingsUpdateNetworkSettingsSyslogServer         `json:"syslogServer,omitempty"`          //
	SNMPServer           *RequestNetworkSettingsUpdateNetworkSettingsSNMPServer           `json:"snmpServer,omitempty"`            //
	Netflowcollector     *RequestNetworkSettingsUpdateNetworkSettingsNetflowcollector     `json:"netflowcollector,omitempty"`      //
	NtpServer            []string                                                         `json:"ntpServer,omitempty"`             // IP address for NTP server (eg: 1.1.1.2)
	Timezone             string                                                           `json:"timezone,omitempty"`              // Input for time zone (eg: Africa/Abidjan)
	MessageOfTheday      *RequestNetworkSettingsUpdateNetworkSettingsMessageOfTheday      `json:"messageOfTheday,omitempty"`       //
	NetworkAAA           *RequestNetworkSettingsUpdateNetworkSettingsNetworkAAA           `json:"network_aaa,omitempty"`           //
	ClientAndEndpointAAA *RequestNetworkSettingsUpdateNetworkSettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkSettingsDNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         // Domain Name of DHCP (eg; cisco)
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   // Primary IP Address for DHCP (eg: 2.2.2.2)
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsUpdateNetworkSettingsSyslogServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for syslog server (eg: 4.4.4.4)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkSettingsSNMPServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for SNMP Server (eg: 4.4.4.1)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkSettingsNetflowcollector struct {
	IPAddress string   `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)
	Port      *float64 `json:"port,omitempty"`      // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsUpdateNetworkSettingsMessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        // Massage for Banner message (eg; Good day)
	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsUpdateNetworkSettingsNetworkAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type for AAA Network (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for AAA and ISE server (eg: 1.1.1.1)
	Network      string `json:"network,omitempty"`      // IP Address for AAA or ISE server (eg: 2.2.2.2)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsUpdateNetworkSettingsClientAndEndpointAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type AAA or ISE server (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for ISE serve (eg: 1.1.1.4)
	Network      string `json:"network,omitempty"`      // IP address for AAA or ISE server (eg: 2.2.2.1)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsReserveIPSubpool struct {
	Name             string   `json:"name,omitempty"`             // Name of the reserve ip sub pool
	Type             string   `json:"type,omitempty"`             // Type of the reserve ip sub pool
	IPv6AddressSpace *bool    `json:"ipv6AddressSpace,omitempty"` // If the value is omitted or false only ipv4 input are required, otherwise both ipv6 and ipv4 are required
	IPv4GlobalPool   string   `json:"ipv4GlobalPool,omitempty"`   // IP v4 Global pool address with cidr, example: 175.175.0.0/16
	IPv4Prefix       *bool    `json:"ipv4Prefix,omitempty"`       // IPv4 prefix value is true, the ip4 prefix length input field is enabled , if it is false ipv4 total Host input is enable
	IPv4PrefixLength *int     `json:"ipv4PrefixLength,omitempty"` // The ipv4 prefix length is required when ipv4prefix value is true.
	IPv4Subnet       string   `json:"ipv4Subnet,omitempty"`       // IPv4 Subnet address, example: 175.175.0.0. Either ipv4Subnet or ipv4TotalHost needs to be passed if creating IPv4 subpool.
	IPv4GateWay      string   `json:"ipv4GateWay,omitempty"`      // Gateway ip address details, example: 175.175.0.1
	IPv4DhcpServers  []string `json:"ipv4DhcpServers,omitempty"`  // IPv4 input for dhcp server ip example: ["1.1.1.1"]
	IPv4DNSServers   []string `json:"ipv4DnsServers,omitempty"`   // IPv4 input for dns server ip example: ["4.4.4.4"]
	IPv6GlobalPool   string   `json:"ipv6GlobalPool,omitempty"`   // IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64
	IPv6Prefix       *bool    `json:"ipv6Prefix,omitempty"`       // Ipv6 prefix value is true, the ip6 prefix length input field is enabled , if it is false ipv6 total Host input is enable
	IPv6PrefixLength *int     `json:"ipv6PrefixLength,omitempty"` // IPv6 prefix length is required when the ipv6prefix value is true
	IPv6Subnet       string   `json:"ipv6Subnet,omitempty"`       // IPv6 Subnet address, example :2001:db8:85a3:0:100::. Either ipv6Subnet or ipv6TotalHost needs to be passed if creating IPv6 subpool.
	IPv6GateWay      string   `json:"ipv6GateWay,omitempty"`      // Gateway ip address details, example: 2001:db8:85a3:0:100::1
	IPv6DhcpServers  []string `json:"ipv6DhcpServers,omitempty"`  // IPv6 format dhcp server as input example : ["2001:db8::1234"]
	IPv6DNSServers   []string `json:"ipv6DnsServers,omitempty"`   // IPv6 format dns server input example: ["2001:db8::1234"]
	IPv4TotalHost    *int     `json:"ipv4TotalHost,omitempty"`    // IPv4 total host is required when ipv4prefix value is false.
	IPv6TotalHost    *int     `json:"ipv6TotalHost,omitempty"`    // IPv6 total host is required when ipv6prefix value is false.
	SLAacSupport     *bool    `json:"slaacSupport,omitempty"`     // Slaac Support
}
type RequestNetworkSettingsUpdateReserveIPSubpool struct {
	Name             string   `json:"name,omitempty"`             // Name of the reserve ip sub pool
	IPv6AddressSpace *bool    `json:"ipv6AddressSpace,omitempty"` // If the value is false only ipv4 input are required. NOTE if value is false then any existing ipv6 subpool in the group will be removed.
	IPv4DhcpServers  []string `json:"ipv4DhcpServers,omitempty"`  // IPv4 input for dhcp server ip example: ["1.1.1.1"]
	IPv4DNSServers   []string `json:"ipv4DnsServers,omitempty"`   // IPv4 input for dns server ip  example: ["4.4.4.4"]
	IPv6GlobalPool   string   `json:"ipv6GlobalPool,omitempty"`   // IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64
	IPv6Prefix       *bool    `json:"ipv6Prefix,omitempty"`       // Ipv6 prefix value is true, the ip6 prefix length input field is enabled, if it is false ipv6 total Host input is enable
	IPv6PrefixLength *int     `json:"ipv6PrefixLength,omitempty"` // IPv6 prefix length is required when the ipv6prefix value is true
	IPv6Subnet       string   `json:"ipv6Subnet,omitempty"`       // IPv6 Subnet address, example :2001:db8:85a3:0:100::.
	IPv6TotalHost    *int     `json:"ipv6TotalHost,omitempty"`    // Size of pool in terms of number of IPs. IPv6 total host is required when ipv6prefix value is false.
	IPv6GateWay      string   `json:"ipv6GateWay,omitempty"`      // Gateway ip address details, example: 2001:db8:85a3:0:100::1
	IPv6DhcpServers  []string `json:"ipv6DhcpServers,omitempty"`  // IPv6 format dhcp server as input example : ["2001:db8::1234"]
	IPv6DNSServers   []string `json:"ipv6DnsServers,omitempty"`   // IPv6 format dns server input example: ["2001:db8::1234"]
	SLAacSupport     *bool    `json:"slaacSupport,omitempty"`     // Slaac Support
	IPv4GateWay      string   `json:"ipv4GateWay,omitempty"`      // Gateway ip address details, example: 175.175.0.1
}
type RequestNetworkSettingsCreateSpProfile struct {
	Settings *RequestNetworkSettingsCreateSpProfileSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileSettings struct {
	Qos *[]RequestNetworkSettingsCreateSpProfileSettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileSettingsQos struct {
	ProfileName string `json:"profileName,omitempty"` // Profile Name
	Model       string `json:"model,omitempty"`       // Model
	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider
}
type RequestNetworkSettingsUpdateSpProfile struct {
	Settings *RequestNetworkSettingsUpdateSpProfileSettings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileSettings struct {
	Qos *[]RequestNetworkSettingsUpdateSpProfileSettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileSettingsQos struct {
	ProfileName    string `json:"profileName,omitempty"`    // Profile Name
	Model          string `json:"model,omitempty"`          // Model
	WanProvider    string `json:"wanProvider,omitempty"`    // Wan Provider
	OldProfileName string `json:"oldProfileName,omitempty"` // Old Profile Name
}
type RequestNetworkSettingsSyncNetworkDevicesCredential struct {
	DeviceCredentialID string `json:"deviceCredentialId,omitempty"` // It must be cli/snmpV2Read/snmpV2Write/snmpV3 Id.
	SiteID             string `json:"siteId,omitempty"`             // Site Id.
}
type RequestNetworkSettingsSetAAASettingsForASite struct {
	AAANetwork *RequestNetworkSettingsSetAAASettingsForASiteAAANetwork `json:"aaaNetwork,omitempty"` //
	AAAClient  *RequestNetworkSettingsSetAAASettingsForASiteAAAClient  `json:"aaaClient,omitempty"`  //
}
type RequestNetworkSettingsSetAAASettingsForASiteAAANetwork struct {
	ServerType        string `json:"serverType,omitempty"`        // Server Type
	Protocol          string `json:"protocol,omitempty"`          // Protocol
	Pan               string `json:"pan,omitempty"`               // Administration Node. Required for ISE.
	PrimaryServerIP   string `json:"primaryServerIp,omitempty"`   // The server to use as a primary.
	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.
	SharedSecret      string `json:"sharedSecret,omitempty"`      // Shared Secret
}
type RequestNetworkSettingsSetAAASettingsForASiteAAAClient struct {
	ServerType        string `json:"serverType,omitempty"`        // Server Type
	Protocol          string `json:"protocol,omitempty"`          // Protocol
	Pan               string `json:"pan,omitempty"`               // Administration Node.  Required for ISE.
	PrimaryServerIP   string `json:"primaryServerIp,omitempty"`   // The server to use as a primary.
	SecondaryServerIP string `json:"secondaryServerIp,omitempty"` // The server to use as a secondary.
	SharedSecret      string `json:"sharedSecret,omitempty"`      // Shared Secret
}
type RequestNetworkSettingsSetBannerSettingsForASite struct {
	Banner *RequestNetworkSettingsSetBannerSettingsForASiteBanner `json:"banner,omitempty"` //
}
type RequestNetworkSettingsSetBannerSettingsForASiteBanner struct {
	Type    string `json:"type,omitempty"`    // Type
	Message string `json:"message,omitempty"` // Custom message that appears when logging into routers, switches, and hubs. Required for custom type.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASite struct {
	CliCredentialsID          *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteCliCredentialsID          `json:"cliCredentialsId,omitempty"`          //
	SNMPv2CReadCredentialsID  *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CReadCredentialsID  `json:"snmpv2cReadCredentialsId,omitempty"`  //
	SNMPv2CWriteCredentialsID *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CWriteCredentialsID `json:"snmpv2cWriteCredentialsId,omitempty"` //
	SNMPv3CredentialsID       *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv3CredentialsID       `json:"snmpv3CredentialsId,omitempty"`       //
	HTTPReadCredentialsID     *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPReadCredentialsID     `json:"httpReadCredentialsId,omitempty"`     //
	HTTPWriteCredentialsID    *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPWriteCredentialsID    `json:"httpWriteCredentialsId,omitempty"`    //
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteCliCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CReadCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv2CWriteCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteSNMPv3CredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPReadCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsUpdateDeviceCredentialSettingsForASiteHTTPWriteCredentialsID struct {
	CredentialsID string `json:"credentialsId,omitempty"` // The `id` of the credentials.
}
type RequestNetworkSettingsSetDhcpSettingsForASite struct {
	Dhcp *RequestNetworkSettingsSetDhcpSettingsForASiteDhcp `json:"dhcp,omitempty"` //
}
type RequestNetworkSettingsSetDhcpSettingsForASiteDhcp struct {
	Servers []string `json:"servers,omitempty"` // DHCP servers for managing client device networking configuration. Max:10
}
type RequestNetworkSettingsSetDNSSettingsForASite struct {
	DNS *RequestNetworkSettingsSetDNSSettingsForASiteDNS `json:"dns,omitempty"` //
}
type RequestNetworkSettingsSetDNSSettingsForASiteDNS struct {
	DomainName string   `json:"domainName,omitempty"` // Network's domain name. Example : myCompnay.com
	DNSServers []string `json:"dnsServers,omitempty"` // DNS servers for hostname resolution.
}
type RequestNetworkSettingsSetImageDistributionSettingsForASite struct {
	ImageDistribution *RequestNetworkSettingsSetImageDistributionSettingsForASiteImageDistribution `json:"imageDistribution,omitempty"` //
}
type RequestNetworkSettingsSetImageDistributionSettingsForASiteImageDistribution struct {
	Servers []string `json:"servers,omitempty"` // This field holds an array of unique identifiers representing image distribution servers. Use ‘/intent/api/v1/images/distributionServerSettings’ to find the Image distribution server Id. Max:2. Use SFTP servers to act as image distribution servers. A distributed SWIM architecture, using suitably located SFTP servers, can help support large-scale device software image upgrades and conserve WAN bandwidth.
}
type RequestNetworkSettingsSetNTPSettingsForASite struct {
	Ntp *RequestNetworkSettingsSetNTPSettingsForASiteNtp `json:"ntp,omitempty"` //
}
type RequestNetworkSettingsSetNTPSettingsForASiteNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP servers to facilitate system clock synchronization for your network. Max:10
}
type RequestNetworkSettingsSetTelemetrySettingsForASite struct {
	WiredDataCollection   *RequestNetworkSettingsSetTelemetrySettingsForASiteWiredDataCollection   `json:"wiredDataCollection,omitempty"`   //
	WirelessTelemetry     *RequestNetworkSettingsSetTelemetrySettingsForASiteWirelessTelemetry     `json:"wirelessTelemetry,omitempty"`     //
	SNMPTraps             *RequestNetworkSettingsSetTelemetrySettingsForASiteSNMPTraps             `json:"snmpTraps,omitempty"`             //
	Syslogs               *RequestNetworkSettingsSetTelemetrySettingsForASiteSyslogs               `json:"syslogs,omitempty"`               //
	ApplicationVisibility *RequestNetworkSettingsSetTelemetrySettingsForASiteApplicationVisibility `json:"applicationVisibility,omitempty"` //
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteWiredDataCollection struct {
	EnableWiredDataCollectio *bool `json:"enableWiredDataCollectio,omitempty"` // Track the presence, location, and movement of wired endpoints in the network. Traffic received from endpoints is used to extract and store their identity information (MAC address and IP address). Other features, such as IEEE 802.1X, web authentication, Cisco Security Groups (formerly TrustSec), SD-Access, and Assurance, depend on this identity information to operate properly. Wired Endpoint Data Collection enables Device Tracking policies on devices assigned to the Access role in Inventory.
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteWirelessTelemetry struct {
	EnableWirelessTelemetry *bool `json:"enableWirelessTelemetry,omitempty"` // Enables Streaming Telemetry on your wireless controllers in order to determine the health of your wireless controller, access points and wireless clients.
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteSNMPTraps struct {
	UseBuiltinTrapServer *bool    `json:"useBuiltinTrapServer,omitempty"` // Enable this server as a destination server for SNMP traps and messages from your network
	ExternalTrapServers  []string `json:"externalTrapServers,omitempty"`  // External SNMP trap servers. Example: ["250.162.252.170","2001:db8:3c4d:15::1a2f:1a2b"]
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteSyslogs struct {
	UseBuiltinSyslogServer *bool    `json:"useBuiltinSyslogServer,omitempty"` // Enable this server as a destination server for syslog messages.
	ExternalSyslogServers  []string `json:"externalSyslogServers,omitempty"`  // External syslog servers. Example: ["250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"]
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteApplicationVisibility struct {
	Collector                  *RequestNetworkSettingsSetTelemetrySettingsForASiteApplicationVisibilityCollector `json:"collector,omitempty"`                  //
	EnableOnWiredAccessDevices *bool                                                                             `json:"enableOnWiredAccessDevices,omitempty"` // Enable Netflow Application Telemetry and Controller Based Application Recognition (CBAR) by default upon network device site assignment for wired access devices.
}
type RequestNetworkSettingsSetTelemetrySettingsForASiteApplicationVisibilityCollector struct {
	CollectorType string `json:"collectorType,omitempty"` // Collector Type
	Address       string `json:"address,omitempty"`       // IP Address. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional. Examples: "250.162.252.170", "2001:db8:3c4d:15::1a2f:1a2b"
	Port          *int   `json:"port,omitempty"`          // Min:1; Max: 65535. If collection type is 'TelemetryBrokerOrUDPDirector', this field value is mandatory otherwise it is optional.
}
type RequestNetworkSettingsSetTimeZoneForASite struct {
	TimeZone *RequestNetworkSettingsSetTimeZoneForASiteTimeZone `json:"timeZone,omitempty"` //
}
type RequestNetworkSettingsSetTimeZoneForASiteTimeZone struct {
	IDentifier string `json:"identifier,omitempty"` // Time zone that corresponds to the site's physical location. The site time zone is used when scheduling device provisioning and updates. Example: GMT
}
type RequestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite struct {
	DeviceIDs []string `json:"deviceIds,omitempty"` // The list of device Ids to perform the provisioning against
}
type RequestNetworkSettingsAssignDeviceCredentialToSiteV2 struct {
	CliID         string `json:"cliId,omitempty"`         // CLI Credential Id
	SNMPV2ReadID  string `json:"snmpV2ReadId,omitempty"`  // SNMPv2c Read Credential Id
	SNMPV2WriteID string `json:"snmpV2WriteId,omitempty"` // SNMPv2c Write Credential Id
	SNMPV3ID      string `json:"snmpV3Id,omitempty"`      // SNMPv3 Credential Id
	HTTPRead      string `json:"httpRead,omitempty"`      // HTTP(S) Read Credential Id
	HTTPWrite     string `json:"httpWrite,omitempty"`     // HTTP(S) Write Credential Id
}
type RequestNetworkSettingsCreateNetworkV2 struct {
	Settings *RequestNetworkSettingsCreateNetworkV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkV2Settings struct {
	DhcpServer           []string                                                           `json:"dhcpServer,omitempty"`            // DHCP Server IP (eg: 1.1.1.1)
	DNSServer            *RequestNetworkSettingsCreateNetworkV2SettingsDNSServer            `json:"dnsServer,omitempty"`             //
	SyslogServer         *RequestNetworkSettingsCreateNetworkV2SettingsSyslogServer         `json:"syslogServer,omitempty"`          //
	SNMPServer           *RequestNetworkSettingsCreateNetworkV2SettingsSNMPServer           `json:"snmpServer,omitempty"`            //
	Netflowcollector     *RequestNetworkSettingsCreateNetworkV2SettingsNetflowcollector     `json:"netflowcollector,omitempty"`      //
	NtpServer            []string                                                           `json:"ntpServer,omitempty"`             // IP address for NTP server (eg: 1.1.1.2)
	Timezone             string                                                             `json:"timezone,omitempty"`              // Input for time zone (eg: Africa/Abidjan)
	MessageOfTheday      *RequestNetworkSettingsCreateNetworkV2SettingsMessageOfTheday      `json:"messageOfTheday,omitempty"`       //
	NetworkAAA           *RequestNetworkSettingsCreateNetworkV2SettingsNetworkAAA           `json:"network_aaa,omitempty"`           //
	ClientAndEndpointAAA *RequestNetworkSettingsCreateNetworkV2SettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsCreateNetworkV2SettingsDNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         // Domain Name of DHCP (eg; cisco)
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   // Primary IP Address for DHCP (eg: 2.2.2.2)
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsCreateNetworkV2SettingsSyslogServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for syslog server (eg: 4.4.4.4)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsCreateNetworkV2SettingsSNMPServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for SNMP Server (eg: 4.4.4.1)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsCreateNetworkV2SettingsNetflowcollector struct {
	IPAddress string   `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)
	Port      *float64 `json:"port,omitempty"`      // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsCreateNetworkV2SettingsMessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        // Massage for Banner message (eg; Good day)
	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsCreateNetworkV2SettingsNetworkAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type for AAA Network (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for AAA and ISE server (eg: 1.1.1.1)
	Network      string `json:"network,omitempty"`      // IP Address for AAA or ISE server (eg: 2.2.2.2)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsCreateNetworkV2SettingsClientAndEndpointAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type AAA or ISE server (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for ISE serve (eg: 1.1.1.4)
	Network      string `json:"network,omitempty"`      // IP address for AAA or ISE server (eg: 2.2.2.1)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsUpdateNetworkV2 struct {
	Settings *RequestNetworkSettingsUpdateNetworkV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkV2Settings struct {
	DhcpServer           []string                                                           `json:"dhcpServer,omitempty"`            // DHCP Server IP (eg: 1.1.1.1)
	DNSServer            *RequestNetworkSettingsUpdateNetworkV2SettingsDNSServer            `json:"dnsServer,omitempty"`             //
	SyslogServer         *RequestNetworkSettingsUpdateNetworkV2SettingsSyslogServer         `json:"syslogServer,omitempty"`          //
	SNMPServer           *RequestNetworkSettingsUpdateNetworkV2SettingsSNMPServer           `json:"snmpServer,omitempty"`            //
	Netflowcollector     *RequestNetworkSettingsUpdateNetworkV2SettingsNetflowcollector     `json:"netflowcollector,omitempty"`      //
	NtpServer            []string                                                           `json:"ntpServer,omitempty"`             // IP address for NTP server (eg: 1.1.1.2)
	Timezone             string                                                             `json:"timezone,omitempty"`              // Input for time zone (eg: Africa/Abidjan)
	MessageOfTheday      *RequestNetworkSettingsUpdateNetworkV2SettingsMessageOfTheday      `json:"messageOfTheday,omitempty"`       //
	NetworkAAA           *RequestNetworkSettingsUpdateNetworkV2SettingsNetworkAAA           `json:"network_aaa,omitempty"`           //
	ClientAndEndpointAAA *RequestNetworkSettingsUpdateNetworkV2SettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
}
type RequestNetworkSettingsUpdateNetworkV2SettingsDNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         // Domain Name of DHCP (eg; cisco)
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   // Primary IP Address for DHCP (eg: 2.2.2.2)
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` // Secondary IP Address for DHCP (eg: 3.3.3.3)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsSyslogServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for syslog server (eg: 4.4.4.4)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for syslog server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsSNMPServer struct {
	IPAddresses     []string `json:"ipAddresses,omitempty"`     // IP Address for SNMP Server (eg: 4.4.4.1)
	ConfigureDnacIP *bool    `json:"configureDnacIP,omitempty"` // Configuration DNAC IP for SNMP Server (eg: true)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsNetflowcollector struct {
	IPAddress string   `json:"ipAddress,omitempty"` // IP Address for NetFlow collector (eg: 3.3.3.1)
	Port      *float64 `json:"port,omitempty"`      // Port for NetFlow Collector (eg; 443)
}
type RequestNetworkSettingsUpdateNetworkV2SettingsMessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        // Massage for Banner message (eg; Good day)
	RetainExistingBanner string `json:"retainExistingBanner,omitempty"` // Retain existing Banner Message (eg: "true" or "false")
}
type RequestNetworkSettingsUpdateNetworkV2SettingsNetworkAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type for AAA Network (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for AAA and ISE server (eg: 1.1.1.1)
	Network      string `json:"network,omitempty"`      // IP Address for AAA or ISE server (eg: 2.2.2.2)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE Server
}
type RequestNetworkSettingsUpdateNetworkV2SettingsClientAndEndpointAAA struct {
	Servers      string `json:"servers,omitempty"`      // Server type AAA or ISE server (eg: AAA)
	IPAddress    string `json:"ipAddress,omitempty"`    // IP address for ISE serve (eg: 1.1.1.4)
	Network      string `json:"network,omitempty"`      // IP address for AAA or ISE server (eg: 2.2.2.1)
	Protocol     string `json:"protocol,omitempty"`     // Protocol for AAA or ISE serve (eg: RADIUS)
	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret for ISE server
}
type RequestNetworkSettingsCreateSpProfileV2 struct {
	Settings *RequestNetworkSettingsCreateSpProfileV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileV2Settings struct {
	Qos *[]RequestNetworkSettingsCreateSpProfileV2SettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsCreateSpProfileV2SettingsQos struct {
	ProfileName string `json:"profileName,omitempty"` // Profile Name
	Model       string `json:"model,omitempty"`       // Model
	WanProvider string `json:"wanProvider,omitempty"` // Wan Provider
}
type RequestNetworkSettingsUpdateSpProfileV2 struct {
	Settings *RequestNetworkSettingsUpdateSpProfileV2Settings `json:"settings,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileV2Settings struct {
	Qos *[]RequestNetworkSettingsUpdateSpProfileV2SettingsQos `json:"qos,omitempty"` //
}
type RequestNetworkSettingsUpdateSpProfileV2SettingsQos struct {
	ProfileName    string `json:"profileName,omitempty"`    // Profile Name
	Model          string `json:"model,omitempty"`          // Model
	WanProvider    string `json:"wanProvider,omitempty"`    // Wan Provider
	OldProfileName string `json:"oldProfileName,omitempty"` // Old Profile Name
}

//GetDeviceCredentialDetails Get Device Credential Details - 899f-08e7-401b-82dd
/* API to get device credential details. This API has been deprecated and will not be available in a Cisco DNA Center release after August 1st 2024 23:59:59 GMT. Please refer new Intent API : Get All Global Credentials V2


@param GetDeviceCredentialDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-credential-details-v1
*/
func (s *NetworkSettingsService) GetDeviceCredentialDetails(GetDeviceCredentialDetailsQueryParams *GetDeviceCredentialDetailsQueryParams) (*ResponseNetworkSettingsGetDeviceCredentialDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-credential"

	queryString, _ := query.Values(GetDeviceCredentialDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetDeviceCredentialDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCredentialDetails(GetDeviceCredentialDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCredentialDetails")
	}

	result := response.Result().(*ResponseNetworkSettingsGetDeviceCredentialDetails)
	return result, response, err

}

//GetGlobalPool Get Global Pool - c0bc-a856-43c8-b58d
/* API to get the global pool.


@param GetGlobalPoolQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-global-pool-v1
*/
func (s *NetworkSettingsService) GetGlobalPool(GetGlobalPoolQueryParams *GetGlobalPoolQueryParams) (*ResponseNetworkSettingsGetGlobalPool, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-pool"

	queryString, _ := query.Values(GetGlobalPoolQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetGlobalPool{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetGlobalPool(GetGlobalPoolQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetGlobalPool")
	}

	result := response.Result().(*ResponseNetworkSettingsGetGlobalPool)
	return result, response, err

}

//GetNetwork Get Network - 38b7-eb13-449b-9471
/* API to get  DHCP and DNS center server details.


@param GetNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-v1
*/
func (s *NetworkSettingsService) GetNetwork(GetNetworkQueryParams *GetNetworkQueryParams) (*ResponseNetworkSettingsGetNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/network"

	queryString, _ := query.Values(GetNetworkQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetNetwork{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetwork(GetNetworkQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetwork")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetwork)
	return result, response, err

}

//GetReserveIPSubpool Get Reserve IP Subpool - 4586-0917-4fab-87e2
/* API to get the ip subpool info.


@param GetReserveIPSubpoolQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-reserve-ip-subpool-v1
*/
func (s *NetworkSettingsService) GetReserveIPSubpool(GetReserveIPSubpoolQueryParams *GetReserveIPSubpoolQueryParams) (*ResponseNetworkSettingsGetReserveIPSubpool, *resty.Response, error) {
	path := "/dna/intent/api/v1/reserve-ip-subpool"

	queryString, _ := query.Values(GetReserveIPSubpoolQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetReserveIPSubpool{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetReserveIPSubpool(GetReserveIPSubpoolQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsGetReserveIPSubpool)
	return result, response, err

}

//GetServiceProviderDetails Get Service provider Details - 7084-7bdc-4d89-a437
/* API to get service provider details (QoS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-service-provider-details-v1
*/
func (s *NetworkSettingsService) GetServiceProviderDetails() (*ResponseNetworkSettingsGetServiceProviderDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsGetServiceProviderDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetServiceProviderDetails()
		}
		return nil, response, fmt.Errorf("error with operation GetServiceProviderDetails")
	}

	result := response.Result().(*ResponseNetworkSettingsGetServiceProviderDetails)
	return result, response, err

}

//RetrieveAAASettingsForASite Retrieve AAA settings for a site - 3c99-79ea-4ab9-bd33
/* Retrieve AAA settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveAAASettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-a-a-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveAAASettingsForASite(id string, RetrieveAAASettingsForASiteQueryParams *RetrieveAAASettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveAAASettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/aaaSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveAAASettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveAAASettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAAASettingsForASite(id, RetrieveAAASettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAAASettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveAAASettingsForASite)
	return result, response, err

}

//RetrieveBannerSettingsForASite Retrieve banner settings for a site - 2a9f-3b2f-4cda-8390
/* Retrieve banner settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveBannerSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-banner-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveBannerSettingsForASite(id string, RetrieveBannerSettingsForASiteQueryParams *RetrieveBannerSettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveBannerSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/bannerSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveBannerSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveBannerSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveBannerSettingsForASite(id, RetrieveBannerSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveBannerSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveBannerSettingsForASite)
	return result, response, err

}

//GetDeviceCredentialSettingsForASite Get device credential settings for a site - bebf-c9fc-4d3a-be03
/* Gets device credential settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the credential is unset, and that no credential of that type will be used for the site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

@param GetDeviceCredentialSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-credential-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) GetDeviceCredentialSettingsForASite(id string, GetDeviceCredentialSettingsForASiteQueryParams *GetDeviceCredentialSettingsForASiteQueryParams) (*ResponseNetworkSettingsGetDeviceCredentialSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/deviceCredentials"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetDeviceCredentialSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetDeviceCredentialSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceCredentialSettingsForASite(id, GetDeviceCredentialSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceCredentialSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsGetDeviceCredentialSettingsForASite)
	return result, response, err

}

//GetNetworkDevicesCredentialsSyncStatus Get network devices credentials sync status - 0f9e-d8a6-45eb-9590
/* Get network devices credentials sync status at a given site.


@param id id path parameter. Site Id.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-devices-credentials-sync-status-v1
*/
func (s *NetworkSettingsService) GetNetworkDevicesCredentialsSyncStatus(id string) (*ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/deviceCredentials/status"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDevicesCredentialsSyncStatus(id)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDevicesCredentialsSyncStatus")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetworkDevicesCredentialsSyncStatus)
	return result, response, err

}

//RetrieveDHCPSettingsForASite Retrieve DHCP settings for a site - cfbb-ca8d-4529-a94b
/* Retrieve DHCP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveDHCPSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-d-h-c-p-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveDHCPSettingsForASite(id string, RetrieveDHCPSettingsForASiteQueryParams *RetrieveDHCPSettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveDHCPSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dhcpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveDHCPSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveDHCPSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveDHCPSettingsForASite(id, RetrieveDHCPSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveDHCPSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveDHCPSettingsForASite)
	return result, response, err

}

//RetrieveDNSSettingsForASite Retrieve DNS settings for a site - d7a4-0932-41d9-bcf8
/* Retrieve DNS settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveDNSSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-d-n-s-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveDNSSettingsForASite(id string, RetrieveDNSSettingsForASiteQueryParams *RetrieveDNSSettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveDNSSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dnsSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveDNSSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveDNSSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveDNSSettingsForASite(id, RetrieveDNSSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveDNSSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveDNSSettingsForASite)
	return result, response, err

}

//RetrieveImageDistributionSettingsForASite Retrieve image distribution settings for a site - d2ad-d9bc-4bcb-9fed
/* Retrieve image distribution settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveImageDistributionSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-image-distribution-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveImageDistributionSettingsForASite(id string, RetrieveImageDistributionSettingsForASiteQueryParams *RetrieveImageDistributionSettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveImageDistributionSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/imageDistributionSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveImageDistributionSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveImageDistributionSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveImageDistributionSettingsForASite(id, RetrieveImageDistributionSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveImageDistributionSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveImageDistributionSettingsForASite)
	return result, response, err

}

//RetrieveNTPSettingsForASite Retrieve NTP settings for a site - beae-2bf1-4cdb-8f60
/* Retrieve NTP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveNTPSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-n-t-p-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveNTPSettingsForASite(id string, RetrieveNTPSettingsForASiteQueryParams *RetrieveNTPSettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveNTPSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/ntpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveNTPSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveNTPSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveNTPSettingsForASite(id, RetrieveNTPSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveNTPSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveNTPSettingsForASite)
	return result, response, err

}

//RetrieveTelemetrySettingsForASite Retrieve Telemetry settings for a site - 11a7-cbc7-4a9a-bac3
/* Retrieves telemetry settings for the given site. `null` values indicate that the setting will be inherited from the parent site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

@param RetrieveTelemetrySettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-telemetry-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveTelemetrySettingsForASite(id string, RetrieveTelemetrySettingsForASiteQueryParams *RetrieveTelemetrySettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveTelemetrySettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/telemetrySettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveTelemetrySettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveTelemetrySettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTelemetrySettingsForASite(id, RetrieveTelemetrySettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTelemetrySettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveTelemetrySettingsForASite)
	return result, response, err

}

//RetrieveTimeZoneSettingsForASite Retrieve time zone settings for a site - 5ba6-0966-4768-9ae7
/* Retrieve time zone settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the setting is unset at a site.


@param id id path parameter. Site Id

@param RetrieveTimeZoneSettingsForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-time-zone-settings-for-a-site-v1
*/
func (s *NetworkSettingsService) RetrieveTimeZoneSettingsForASite(id string, RetrieveTimeZoneSettingsForASiteQueryParams *RetrieveTimeZoneSettingsForASiteQueryParams) (*ResponseNetworkSettingsRetrieveTimeZoneSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/timeZoneSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveTimeZoneSettingsForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsRetrieveTimeZoneSettingsForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTimeZoneSettingsForASite(id, RetrieveTimeZoneSettingsForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTimeZoneSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsRetrieveTimeZoneSettingsForASite)
	return result, response, err

}

//GetNetworkV2 Get Network V2 - fdbd-3b7b-4048-bbfd
/* API to get SNMP, NTP, Network AAA, Client and Endpoint AAA, and/or DNS center server settings.


@param GetNetworkV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-v2
*/
func (s *NetworkSettingsService) GetNetworkV2(GetNetworkV2QueryParams *GetNetworkV2QueryParams) (*ResponseNetworkSettingsGetNetworkV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/network"

	queryString, _ := query.Values(GetNetworkV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseNetworkSettingsGetNetworkV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkV2(GetNetworkV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetworkV2)
	return result, response, err

}

//GetServiceProviderDetailsV2 Get Service Provider Details V2 - c89e-0ac4-4279-8591
/* API to get Service Provider details (QoS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-service-provider-details-v2
*/
func (s *NetworkSettingsService) GetServiceProviderDetailsV2() (*ResponseNetworkSettingsGetServiceProviderDetailsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsGetServiceProviderDetailsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetServiceProviderDetailsV2()
		}
		return nil, response, fmt.Errorf("error with operation GetServiceProviderDetailsV2")
	}

	result := response.Result().(*ResponseNetworkSettingsGetServiceProviderDetailsV2)
	return result, response, err

}

//AssignDeviceCredentialToSite Assign Device Credential To Site - 4da9-1a54-4e29-842d
/* Assign Device Credential to a site.


@param siteID siteId path parameter. site id to assign credential.

@param AssignDeviceCredentialToSiteHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-device-credential-to-site-v1
*/
func (s *NetworkSettingsService) AssignDeviceCredentialToSite(siteID string, requestNetworkSettingsAssignDeviceCredentialToSite *RequestNetworkSettingsAssignDeviceCredentialToSite, AssignDeviceCredentialToSiteHeaderParams *AssignDeviceCredentialToSiteHeaderParams) (*ResponseNetworkSettingsAssignDeviceCredentialToSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/credential-to-site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if AssignDeviceCredentialToSiteHeaderParams != nil {

		if AssignDeviceCredentialToSiteHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", AssignDeviceCredentialToSiteHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestNetworkSettingsAssignDeviceCredentialToSite).
		SetResult(&ResponseNetworkSettingsAssignDeviceCredentialToSite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignDeviceCredentialToSite(siteID, requestNetworkSettingsAssignDeviceCredentialToSite, AssignDeviceCredentialToSiteHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation AssignDeviceCredentialToSite")
	}

	result := response.Result().(*ResponseNetworkSettingsAssignDeviceCredentialToSite)
	return result, response, err

}

//CreateDeviceCredentials Create Device Credentials - fbb9-5b37-484a-9fce
/* API to create device credentials. This API has been deprecated and will not be available in a Cisco DNA Center release after August 1st 2024 23:59:59 GMT. Please refer new Intent API : Create Global Credentials V2



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-credentials-v1
*/
func (s *NetworkSettingsService) CreateDeviceCredentials(requestNetworkSettingsCreateDeviceCredentials *RequestNetworkSettingsCreateDeviceCredentials) (*ResponseNetworkSettingsCreateDeviceCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateDeviceCredentials).
		SetResult(&ResponseNetworkSettingsCreateDeviceCredentials{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateDeviceCredentials(requestNetworkSettingsCreateDeviceCredentials)
		}

		return nil, response, fmt.Errorf("error with operation CreateDeviceCredentials")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateDeviceCredentials)
	return result, response, err

}

//CreateGlobalPool Create Global Pool - f793-192a-43da-bed9
/* API to create global pool.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-global-pool-v1
*/
func (s *NetworkSettingsService) CreateGlobalPool(requestNetworkSettingsCreateGlobalPool *RequestNetworkSettingsCreateGlobalPool) (*ResponseNetworkSettingsCreateGlobalPool, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-pool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateGlobalPool).
		SetResult(&ResponseNetworkSettingsCreateGlobalPool{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateGlobalPool(requestNetworkSettingsCreateGlobalPool)
		}

		return nil, response, fmt.Errorf("error with operation CreateGlobalPool")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateGlobalPool)
	return result, response, err

}

//CreateNetwork Create Network - be89-2bd8-4a78-865a
/* API to create a network for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and EndPoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site id to which site details to associate with the network settings.

@param CreateNetworkHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-v1
*/
func (s *NetworkSettingsService) CreateNetwork(siteID string, requestNetworkSettingsCreateNetwork *RequestNetworkSettingsCreateNetwork, CreateNetworkHeaderParams *CreateNetworkHeaderParams) (*ResponseNetworkSettingsCreateNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateNetworkHeaderParams != nil {

		if CreateNetworkHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", CreateNetworkHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestNetworkSettingsCreateNetwork).
		SetResult(&ResponseNetworkSettingsCreateNetwork{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateNetwork(siteID, requestNetworkSettingsCreateNetwork, CreateNetworkHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreateNetwork")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateNetwork)
	return result, response, err

}

//ReserveIPSubpool Reserve IP Subpool - 429f-aa81-4d3b-960a
/* API to reserve an ip subpool from the global pool


@param siteID siteId path parameter. Site id to reserve the ip sub pool.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!reserve-ip-subpool-v1
*/
func (s *NetworkSettingsService) ReserveIPSubpool(siteID string, requestNetworkSettingsReserveIPSubpool *RequestNetworkSettingsReserveIPSubpool) (*ResponseNetworkSettingsReserveIPSubpool, *resty.Response, error) {
	path := "/dna/intent/api/v1/reserve-ip-subpool/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsReserveIPSubpool).
		SetResult(&ResponseNetworkSettingsReserveIPSubpool{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReserveIPSubpool(siteID, requestNetworkSettingsReserveIPSubpool)
		}

		return nil, response, fmt.Errorf("error with operation ReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsReserveIPSubpool)
	return result, response, err

}

//CreateSpProfile Create SP Profile - a39a-1a21-4deb-b781
/* API to create Service Provider Profile(QOS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sp-profile-v1
*/
func (s *NetworkSettingsService) CreateSpProfile(requestNetworkSettingsCreateSPProfile *RequestNetworkSettingsCreateSpProfile) (*ResponseNetworkSettingsCreateSpProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateSPProfile).
		SetResult(&ResponseNetworkSettingsCreateSpProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSpProfile(requestNetworkSettingsCreateSPProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateSpProfile")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateSpProfile)
	return result, response, err

}

//SyncNetworkDevicesCredential Sync network devices credential - 5ca0-2a36-4d68-92e7
/* When sync is triggered at a site with the credential that are associated to the same site, network devices in impacted sites (child sites which are inheriting the credential) get managed in inventory with the associated site credential. Credential gets configured on network devices before these get managed in inventory. Please make a note that cli credential wouldn't be configured on AAA authenticated devices but they just get managed with the associated site cli credential.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!sync-network-devices-credential-v1
*/
func (s *NetworkSettingsService) SyncNetworkDevicesCredential(requestNetworkSettingsSyncNetworkDevicesCredential *RequestNetworkSettingsSyncNetworkDevicesCredential) (*ResponseNetworkSettingsSyncNetworkDevicesCredential, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/deviceCredentials/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSyncNetworkDevicesCredential).
		SetResult(&ResponseNetworkSettingsSyncNetworkDevicesCredential{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.SyncNetworkDevicesCredential(requestNetworkSettingsSyncNetworkDevicesCredential)
		}

		return nil, response, fmt.Errorf("error with operation SyncNetworkDevicesCredential")
	}

	result := response.Result().(*ResponseNetworkSettingsSyncNetworkDevicesCredential)
	return result, response, err

}

//UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite Update a device(s) telemetry settings to conform to the telemetry settings for its site - 14bf-7997-432b-94a1
/* Update a device(s) telemetry settings to conform to the telemetry settings for its site.  One Task is created to track the update, for more granular status tracking, split your devices into multiple requests.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-a-devices-telemetry-settings-to-conform-to-the-telemetry-settings-for-its-site-v1
*/
func (s *NetworkSettingsService) UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite(requestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite *RequestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite) (*ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/telemetrySettings/apply"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite).
		SetResult(&ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite(requestNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite)
		}

		return nil, response, fmt.Errorf("error with operation UpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateADevicesTelemetrySettingsToConformToTheTelemetrySettingsForItsSite)
	return result, response, err

}

//AssignDeviceCredentialToSiteV2 Assign Device Credential To Site V2 - 0eb2-8bc5-4b99-8d6c
/* API to assign Device Credential to a site.


@param siteID siteId path parameter. Site Id to assign credential.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-device-credential-to-site-v2
*/
func (s *NetworkSettingsService) AssignDeviceCredentialToSiteV2(siteID string, requestNetworkSettingsAssignDeviceCredentialToSiteV2 *RequestNetworkSettingsAssignDeviceCredentialToSiteV2) (*ResponseNetworkSettingsAssignDeviceCredentialToSiteV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/credential-to-site/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsAssignDeviceCredentialToSiteV2).
		SetResult(&ResponseNetworkSettingsAssignDeviceCredentialToSiteV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignDeviceCredentialToSiteV2(siteID, requestNetworkSettingsAssignDeviceCredentialToSiteV2)
		}

		return nil, response, fmt.Errorf("error with operation AssignDeviceCredentialToSiteV2")
	}

	result := response.Result().(*ResponseNetworkSettingsAssignDeviceCredentialToSiteV2)
	return result, response, err

}

//CreateNetworkV2 Create Network V2 - 4696-fb19-48da-ac38
/* API to create network settings for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and Endpoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site Id to which site details to associate with the network settings.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network-v2
*/
func (s *NetworkSettingsService) CreateNetworkV2(siteID string, requestNetworkSettingsCreateNetworkV2 *RequestNetworkSettingsCreateNetworkV2) (*ResponseNetworkSettingsCreateNetworkV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateNetworkV2).
		SetResult(&ResponseNetworkSettingsCreateNetworkV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateNetworkV2(siteID, requestNetworkSettingsCreateNetworkV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateNetworkV2)
	return result, response, err

}

//CreateSpProfileV2 Create SP Profile V2 - b6b1-1aa9-4b4b-99e0
/* API to create Service Provider Profile(QOS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sp-profile-v2
*/
func (s *NetworkSettingsService) CreateSpProfileV2(requestNetworkSettingsCreateSPProfileV2 *RequestNetworkSettingsCreateSpProfileV2) (*ResponseNetworkSettingsCreateSpProfileV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsCreateSPProfileV2).
		SetResult(&ResponseNetworkSettingsCreateSpProfileV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSpProfileV2(requestNetworkSettingsCreateSPProfileV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateSpProfileV2)
	return result, response, err

}

//UpdateDeviceCredentials Update Device Credentials - 4f94-7a1c-4fc8-84f6
/* API to update device credentials. This API has been deprecated and will not be available in a Cisco DNA Center release after August 1st 2024 23:59:59 GMT. Please refer new Intent API : Update Global Credentials V2


 */
func (s *NetworkSettingsService) UpdateDeviceCredentials(requestNetworkSettingsUpdateDeviceCredentials *RequestNetworkSettingsUpdateDeviceCredentials) (*ResponseNetworkSettingsUpdateDeviceCredentials, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-credential"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateDeviceCredentials).
		SetResult(&ResponseNetworkSettingsUpdateDeviceCredentials{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceCredentials(requestNetworkSettingsUpdateDeviceCredentials)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCredentials")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateDeviceCredentials)
	return result, response, err

}

//UpdateGlobalPool Update Global Pool - 03b4-c8b4-4919-b964
/* API to update global pool


 */
func (s *NetworkSettingsService) UpdateGlobalPool(requestNetworkSettingsUpdateGlobalPool *RequestNetworkSettingsUpdateGlobalPool) (*ResponseNetworkSettingsUpdateGlobalPool, *resty.Response, error) {
	path := "/dna/intent/api/v1/global-pool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateGlobalPool).
		SetResult(&ResponseNetworkSettingsUpdateGlobalPool{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateGlobalPool(requestNetworkSettingsUpdateGlobalPool)
		}
		return nil, response, fmt.Errorf("error with operation UpdateGlobalPool")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateGlobalPool)
	return result, response, err

}

//UpdateNetwork Update Network - 698b-fbb4-4dcb-9fca
/* API to update network settings for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and EndPoint AAA, and/or DNS server settings.


@param siteID siteId path parameter. Site id to update the network settings which is associated with the site

*/
func (s *NetworkSettingsService) UpdateNetwork(siteID string, requestNetworkSettingsUpdateNetwork *RequestNetworkSettingsUpdateNetwork) (*ResponseNetworkSettingsUpdateNetwork, *resty.Response, error) {
	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateNetwork).
		SetResult(&ResponseNetworkSettingsUpdateNetwork{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateNetwork(siteID, requestNetworkSettingsUpdateNetwork)
		}
		return nil, response, fmt.Errorf("error with operation UpdateNetwork")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateNetwork)
	return result, response, err

}

//UpdateReserveIPSubpool Update Reserve IP Subpool - 6992-d8ec-42cb-88f1
/* API to update ip subpool from the global pool


@param siteID siteId path parameter. Site id of site to update sub pool.

@param UpdateReserveIPSubpoolQueryParams Filtering parameter
*/
func (s *NetworkSettingsService) UpdateReserveIPSubpool(siteID string, requestNetworkSettingsUpdateReserveIPSubpool *RequestNetworkSettingsUpdateReserveIPSubpool, UpdateReserveIPSubpoolQueryParams *UpdateReserveIPSubpoolQueryParams) (*ResponseNetworkSettingsUpdateReserveIPSubpool, *resty.Response, error) {
	path := "/dna/intent/api/v1/reserve-ip-subpool/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(UpdateReserveIPSubpoolQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestNetworkSettingsUpdateReserveIPSubpool).
		SetResult(&ResponseNetworkSettingsUpdateReserveIPSubpool{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateReserveIPSubpool(siteID, requestNetworkSettingsUpdateReserveIPSubpool, UpdateReserveIPSubpoolQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation UpdateReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateReserveIPSubpool)
	return result, response, err

}

//UpdateSpProfile Update SP Profile - 5087-daae-4cc9-8566
/* API to update Service Provider Profile (QoS).


 */
func (s *NetworkSettingsService) UpdateSpProfile(requestNetworkSettingsUpdateSPProfile *RequestNetworkSettingsUpdateSpProfile) (*ResponseNetworkSettingsUpdateSpProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateSPProfile).
		SetResult(&ResponseNetworkSettingsUpdateSpProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSpProfile(requestNetworkSettingsUpdateSPProfile)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSpProfile")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateSpProfile)
	return result, response, err

}

//SetAAASettingsForASite Set AAA settings for a site - 3582-ca30-4718-a064
/* Set AAA settings for a site; `null` values indicate that the settings will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetAAASettingsForASite(id string, requestNetworkSettingsSetAAASettingsForASite *RequestNetworkSettingsSetAAASettingsForASite) (*ResponseNetworkSettingsSetAAASettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/aaaSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetAAASettingsForASite).
		SetResult(&ResponseNetworkSettingsSetAAASettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetAAASettingsForASite(id, requestNetworkSettingsSetAAASettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetAAASettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetAAASettingsForASite)
	return result, response, err

}

//SetBannerSettingsForASite Set banner settings for a site - 0aae-aa56-44f9-95a7
/* Set banner settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetBannerSettingsForASite(id string, requestNetworkSettingsSetBannerSettingsForASite *RequestNetworkSettingsSetBannerSettingsForASite) (*ResponseNetworkSettingsSetBannerSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/bannerSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetBannerSettingsForASite).
		SetResult(&ResponseNetworkSettingsSetBannerSettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetBannerSettingsForASite(id, requestNetworkSettingsSetBannerSettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetBannerSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetBannerSettingsForASite)
	return result, response, err

}

//UpdateDeviceCredentialSettingsForASite Update device credential settings for a site. - 5aa0-6949-4dfa-bec5
/* Updates device credential settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the credential is unset, and that no credential of that type will be used for the site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

*/
func (s *NetworkSettingsService) UpdateDeviceCredentialSettingsForASite(id string, requestNetworkSettingsUpdateDeviceCredentialSettingsForASite *RequestNetworkSettingsUpdateDeviceCredentialSettingsForASite) (*ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/deviceCredentials"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateDeviceCredentialSettingsForASite).
		SetResult(&ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateDeviceCredentialSettingsForASite(id, requestNetworkSettingsUpdateDeviceCredentialSettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation UpdateDeviceCredentialSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateDeviceCredentialSettingsForASite)
	return result, response, err

}

//SetDhcpSettingsForASite Set dhcp settings for a site - c1ac-194d-40d9-8ae4
/* Set DHCP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetDhcpSettingsForASite(id string, requestNetworkSettingsSetDhcpSettingsForASite *RequestNetworkSettingsSetDhcpSettingsForASite) (*ResponseNetworkSettingsSetDhcpSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dhcpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetDhcpSettingsForASite).
		SetResult(&ResponseNetworkSettingsSetDhcpSettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetDhcpSettingsForASite(id, requestNetworkSettingsSetDhcpSettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetDhcpSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetDhcpSettingsForASite)
	return result, response, err

}

//SetDNSSettingsForASite Set DNS settings for a site - 9892-798e-4ed8-a40b
/* Set DNS settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetDNSSettingsForASite(id string, requestNetworkSettingsSetDNSSettingsForASite *RequestNetworkSettingsSetDNSSettingsForASite) (*ResponseNetworkSettingsSetDNSSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/dnsSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetDNSSettingsForASite).
		SetResult(&ResponseNetworkSettingsSetDNSSettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetDNSSettingsForASite(id, requestNetworkSettingsSetDNSSettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetDNSSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetDNSSettingsForASite)
	return result, response, err

}

//SetImageDistributionSettingsForASite Set image distribution settings for a site - 8ab7-3889-45f8-9e5d
/* Set image distribution settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetImageDistributionSettingsForASite(id string, requestNetworkSettingsSetImageDistributionSettingsForASite *RequestNetworkSettingsSetImageDistributionSettingsForASite) (*ResponseNetworkSettingsSetImageDistributionSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/imageDistributionSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetImageDistributionSettingsForASite).
		SetResult(&ResponseNetworkSettingsSetImageDistributionSettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetImageDistributionSettingsForASite(id, requestNetworkSettingsSetImageDistributionSettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetImageDistributionSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetImageDistributionSettingsForASite)
	return result, response, err

}

//SetNTPSettingsForASite Set NTP settings for a site - 9d80-8815-42a9-b006
/* Set NTP settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetNTPSettingsForASite(id string, requestNetworkSettingsSetNTPSettingsForASite *RequestNetworkSettingsSetNTPSettingsForASite) (*ResponseNetworkSettingsSetNTPSettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/ntpSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetNTPSettingsForASite).
		SetResult(&ResponseNetworkSettingsSetNTPSettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetNTPSettingsForASite(id, requestNetworkSettingsSetNTPSettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetNTPSettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetNTPSettingsForASite)
	return result, response, err

}

//SetTelemetrySettingsForASite Set Telemetry settings for a site - a5a1-6835-40ab-8d2f
/* Sets telemetry settings for the given site; `null` values indicate that the setting will be inherited from the parent site.


@param id id path parameter. Site Id, retrievable from the `id` attribute in `/dna/intent/api/v1/sites`

*/
func (s *NetworkSettingsService) SetTelemetrySettingsForASite(id string, requestNetworkSettingsSetTelemetrySettingsForASite *RequestNetworkSettingsSetTelemetrySettingsForASite) (*ResponseNetworkSettingsSetTelemetrySettingsForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/telemetrySettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetTelemetrySettingsForASite).
		SetResult(&ResponseNetworkSettingsSetTelemetrySettingsForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetTelemetrySettingsForASite(id, requestNetworkSettingsSetTelemetrySettingsForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetTelemetrySettingsForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetTelemetrySettingsForASite)
	return result, response, err

}

//SetTimeZoneForASite Set time zone for a site - 16a7-b874-4b19-88d0
/* Set time zone settings for a site; `null` values indicate that the setting will be inherited from the parent site; empty objects (`{}`) indicate that the settings is unset.


@param id id path parameter. Site Id

*/
func (s *NetworkSettingsService) SetTimeZoneForASite(id string, requestNetworkSettingsSetTimeZoneForASite *RequestNetworkSettingsSetTimeZoneForASite) (*ResponseNetworkSettingsSetTimeZoneForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{id}/timeZoneSettings"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsSetTimeZoneForASite).
		SetResult(&ResponseNetworkSettingsSetTimeZoneForASite{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SetTimeZoneForASite(id, requestNetworkSettingsSetTimeZoneForASite)
		}
		return nil, response, fmt.Errorf("error with operation SetTimeZoneForASite")
	}

	result := response.Result().(*ResponseNetworkSettingsSetTimeZoneForASite)
	return result, response, err

}

//UpdateNetworkV2 Update Network V2 - ac85-8bd9-4c78-a705
/* API to update network settings for DHCP, Syslog, SNMP, NTP, Network AAA, Client and Endpoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site Id to update the network settings which is associated with the site

*/
func (s *NetworkSettingsService) UpdateNetworkV2(siteID string, requestNetworkSettingsUpdateNetworkV2 *RequestNetworkSettingsUpdateNetworkV2) (*ResponseNetworkSettingsUpdateNetworkV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/network/{siteId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateNetworkV2).
		SetResult(&ResponseNetworkSettingsUpdateNetworkV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateNetworkV2(siteID, requestNetworkSettingsUpdateNetworkV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdateNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateNetworkV2)
	return result, response, err

}

//UpdateSpProfileV2 Update SP Profile V2 - 7fa6-78e5-455a-94a6
/* API to update Service Provider Profile (QoS).


 */
func (s *NetworkSettingsService) UpdateSpProfileV2(requestNetworkSettingsUpdateSPProfileV2 *RequestNetworkSettingsUpdateSpProfileV2) (*ResponseNetworkSettingsUpdateSpProfileV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/service-provider"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestNetworkSettingsUpdateSPProfileV2).
		SetResult(&ResponseNetworkSettingsUpdateSpProfileV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSpProfileV2(requestNetworkSettingsUpdateSPProfileV2)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateSpProfileV2)
	return result, response, err

}

//DeleteDeviceCredential Delete Device Credential - 259e-ab30-4598-8958
/* Delete device credential. This API has been deprecated and will not be available in a Cisco DNA Center release after August 1st 2024 23:59:59 GMT. Please refer new Intent API : Delete Global Credentials V2


@param id id path parameter. global credential id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-credential-v1
*/
func (s *NetworkSettingsService) DeleteDeviceCredential(id string) (*ResponseNetworkSettingsDeleteDeviceCredential, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/device-credential/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteDeviceCredential{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDeviceCredential(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDeviceCredential")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteDeviceCredential)
	return result, response, err

}

//DeleteGlobalIPPool Delete Global IP Pool - 1eaa-8b21-48ab-81de
/* API to delete global IP pool.


@param id id path parameter. global pool id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-global-ip-pool-v1
*/
func (s *NetworkSettingsService) DeleteGlobalIPPool(id string) (*ResponseNetworkSettingsDeleteGlobalIPPool, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/global-pool/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteGlobalIPPool{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteGlobalIPPool(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteGlobalIpPool")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteGlobalIPPool)
	return result, response, err

}

//ReleaseReserveIPSubpool Release Reserve IP Subpool - 85b2-89e3-4489-9dc1
/* API to delete the reserved ip subpool


@param id id path parameter. Id of reserve ip subpool to be deleted.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!release-reserve-ip-subpool-v1
*/
func (s *NetworkSettingsService) ReleaseReserveIPSubpool(id string) (*ResponseNetworkSettingsReleaseReserveIPSubpool, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/reserve-ip-subpool/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsReleaseReserveIPSubpool{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReleaseReserveIPSubpool(id)
		}
		return nil, response, fmt.Errorf("error with operation ReleaseReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsReleaseReserveIPSubpool)
	return result, response, err

}

//DeleteSpProfile Delete SP Profile - 4ca2-db11-43eb-b5d7
/* API to delete Service Provider Profile (QoS).


@param spProfileName spProfileName path parameter. sp profile name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sp-profile-v1
*/
func (s *NetworkSettingsService) DeleteSpProfile(spProfileName string) (*ResponseNetworkSettingsDeleteSpProfile, *resty.Response, error) {
	//spProfileName string
	path := "/dna/intent/api/v1/sp-profile/{spProfileName}"
	path = strings.Replace(path, "{spProfileName}", fmt.Sprintf("%v", spProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteSpProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSpProfile(spProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSpProfile")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteSpProfile)
	return result, response, err

}

//DeleteSpProfileV2 Delete SP Profile V2 - 5ea0-cb0e-4ed9-9bec
/* API to delete Service Provider Profile (QoS).


@param spProfileName spProfileName path parameter. SP profile name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sp-profile-v2
*/
func (s *NetworkSettingsService) DeleteSpProfileV2(spProfileName string) (*ResponseNetworkSettingsDeleteSpProfileV2, *resty.Response, error) {
	//spProfileName string
	path := "/dna/intent/api/v2/sp-profile/{spProfileName}"
	path = strings.Replace(path, "{spProfileName}", fmt.Sprintf("%v", spProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseNetworkSettingsDeleteSpProfileV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSpProfileV2(spProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteSpProfileV2)
	return result, response, err

}
