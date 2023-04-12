package dnac

import (
	"fmt"
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
	Offset int `url:"offset,omitempty"` //offset/starting row
	Limit  int `url:"limit,omitempty"`  //No of Global Pools to be retrieved
}
type GetNetworkQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site id to get the network settings associated with the site.
}
type CreateNetworkHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetReserveIPSubpoolQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //site id to get the reserve ip associated with the site
	Offset int    `url:"offset,omitempty"` //offset/starting row
	Limit  int    `url:"limit,omitempty"`  //No of Global Pools to be retrieved
}
type UpdateReserveIPSubpoolQueryParams struct {
	ID string `url:"id,omitempty"` //Id of subpool to be associated with the site
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
	IPPoolName            string                                                     `json:"ipPoolName,omitempty"`            // Ip Pool Name
	DhcpServerIPs         []string                                                   `json:"dhcpServerIps,omitempty"`         // Dhcp Server Ips
	Gateways              []string                                                   `json:"gateways,omitempty"`              // Gateways
	CreateTime            string                                                     `json:"createTime,omitempty"`            // Create Time
	LastUpdateTime        string                                                     `json:"lastUpdateTime,omitempty"`        // Last Update Time
	TotalIPAddressCount   string                                                     `json:"totalIpAddressCount,omitempty"`   // Total Ip Address Count
	UsedIPAddressCount    string                                                     `json:"usedIpAddressCount,omitempty"`    // Used Ip Address Count
	ParentUUID            string                                                     `json:"parentUuid,omitempty"`            // Parent Uuid
	Owner                 string                                                     `json:"owner,omitempty"`                 // Owner
	Shared                string                                                     `json:"shared,omitempty"`                // Shared
	Overlapping           string                                                     `json:"overlapping,omitempty"`           // Overlapping
	ConfigureExternalDhcp string                                                     `json:"configureExternalDhcp,omitempty"` // Configure External Dhcp
	UsedPercentage        string                                                     `json:"usedPercentage,omitempty"`        // Used Percentage
	ClientOptions         *ResponseNetworkSettingsGetGlobalPoolResponseClientOptions `json:"clientOptions,omitempty"`         // Client Options
	DNSServerIPs          []string                                                   `json:"dnsServerIps,omitempty"`          // Dns Server Ips
	Context               *[]ResponseNetworkSettingsGetGlobalPoolResponseContext     `json:"context,omitempty"`               //
	IPv6                  string                                                     `json:"ipv6,omitempty"`                  // Ipv6
	ID                    string                                                     `json:"id,omitempty"`                    // Id
	IPPoolCidr            string                                                     `json:"ipPoolCidr,omitempty"`            // Ip Pool Cidr
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
	Version  *int                                                        `json:"version,omitempty"`  // Version
}
type ResponseNetworkSettingsGetServiceProviderDetailsResponse struct {
	InstanceType       string                                                           `json:"instanceType,omitempty"`       // Instance Type
	InstanceUUID       string                                                           `json:"instanceUuid,omitempty"`       // Instance Uuid
	Namespace          string                                                           `json:"namespace,omitempty"`          // Namespace
	Type               string                                                           `json:"type,omitempty"`               // Type
	Key                string                                                           `json:"key,omitempty"`                // Key
	Version            string                                                           `json:"version,omitempty"`            // Version
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
type ResponseNetworkSettingsDeleteSpProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
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
	IPv6AddressSpace *bool    `json:"ipv6AddressSpace,omitempty"` // If the value is false only ipv4 input are required, otherwise both ipv6 and ipv4 are required
	IPv4GlobalPool   string   `json:"ipv4GlobalPool,omitempty"`   // IP v4 Global pool address with cidr, example: 175.175.0.0/16
	IPv4Prefix       *bool    `json:"ipv4Prefix,omitempty"`       // IPv4 prefix value is true, the ip4 prefix length input field is enabled , if it is false ipv4 total Host input is enable
	IPv4PrefixLength *int     `json:"ipv4PrefixLength,omitempty"` // The ipv4 prefix length is required when ipv4prefix value is true.
	IPv4Subnet       string   `json:"ipv4Subnet,omitempty"`       // IPv4 Subnet address, example: 175.175.0.0
	IPv4GateWay      string   `json:"ipv4GateWay,omitempty"`      // Gateway ip address details, example: 175.175.0.1
	IPv4DhcpServers  []string `json:"ipv4DhcpServers,omitempty"`  // IPv4 input for dhcp server ip example: 1.1.1.1
	IPv4DNSServers   []string `json:"ipv4DnsServers,omitempty"`   // IPv4 input for dns server ip example: 4.4.4.4
	IPv6GlobalPool   string   `json:"ipv6GlobalPool,omitempty"`   // IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64
	IPv6Prefix       *bool    `json:"ipv6Prefix,omitempty"`       // Ipv6 prefix value is true, the ip6 prefix length input field is enabled , if it is false ipv6 total Host input is enable
	IPv6PrefixLength *int     `json:"ipv6PrefixLength,omitempty"` // IPv6 prefix length is required when the ipv6prefix value is true
	IPv6Subnet       string   `json:"ipv6Subnet,omitempty"`       // IPv6 Subnet address, example :2001:db8:85a3:0:100::
	IPv6GateWay      string   `json:"ipv6GateWay,omitempty"`      // Gateway ip address details, example: 2001:db8:85a3:0:100::1
	IPv6DhcpServers  []string `json:"ipv6DhcpServers,omitempty"`  // IPv6 format dhcp server as input example : 2001:db8::1234
	IPv6DNSServers   []string `json:"ipv6DnsServers,omitempty"`   // IPv6 format dns server input example: 2001:db8::1234
	IPv4TotalHost    *int     `json:"ipv4TotalHost,omitempty"`    // IPv4 total host is required when ipv4prefix value is false.
	IPv6TotalHost    *int     `json:"ipv6TotalHost,omitempty"`    // IPv6 total host is required when ipv6prefix value is false.
	SLAacSupport     *bool    `json:"slaacSupport,omitempty"`     // Slaac Support
}
type RequestNetworkSettingsUpdateReserveIPSubpool struct {
	Name             string   `json:"name,omitempty"`             // Name of the reserve ip sub pool
	IPv6AddressSpace *bool    `json:"ipv6AddressSpace,omitempty"` // If the value is false only ipv4 input are required, otherwise both ipv6 and ipv4 are required
	IPv4DhcpServers  []string `json:"ipv4DhcpServers,omitempty"`  // IPv4 input for dhcp server ip example: 1.1.1.1
	IPv4DNSServers   []string `json:"ipv4DnsServers,omitempty"`   // IPv4 input for dns server ip  example: 4.4.4.4
	IPv6GlobalPool   string   `json:"ipv6GlobalPool,omitempty"`   // IP v6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64
	IPv6Prefix       *bool    `json:"ipv6Prefix,omitempty"`       // IPv6 prefix value is true, the ip6 prefix length input field is enabled , if it is false  ipv6 total Host input is enable
	IPv6PrefixLength *int     `json:"ipv6PrefixLength,omitempty"` // IPv6 prefix length is required when the ipv6prefix value is true
	IPv6Subnet       string   `json:"ipv6Subnet,omitempty"`       // IPv6 Subnet address, example :2001:db8:85a3:0:100::
	IPv6GateWay      string   `json:"ipv6GateWay,omitempty"`      // Gateway ip address details, example: 2001:db8:85a3:0:100::1
	IPv6DhcpServers  []string `json:"ipv6DhcpServers,omitempty"`  // IPv6 format dhcp server as input example : 2001:db8::1234
	IPv6DNSServers   []string `json:"ipv6DnsServers,omitempty"`   // IPv6 format dns server input example: 2001:db8::1234
	IPv6TotalHost    *int     `json:"ipv6TotalHost,omitempty"`    // IPv6 total host is required when ipv6prefix value is false.
	SLAacSupport     *bool    `json:"slaacSupport,omitempty"`     // Slaac Support
	IPv4GateWay      string   `json:"ipv4GateWay,omitempty"`      // Ipv4 Gate Way
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
/* API to get device credential details.


@param GetDeviceCredentialDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-credential-details
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
		return nil, response, fmt.Errorf("error with operation GetDeviceCredentialDetails")
	}

	result := response.Result().(*ResponseNetworkSettingsGetDeviceCredentialDetails)
	return result, response, err

}

//GetGlobalPool Get Global Pool - c0bc-a856-43c8-b58d
/* API to get global pool.


@param GetGlobalPoolQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-global-pool
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
		return nil, response, fmt.Errorf("error with operation GetGlobalPool")
	}

	result := response.Result().(*ResponseNetworkSettingsGetGlobalPool)
	return result, response, err

}

//GetNetwork Get Network - 38b7-eb13-449b-9471
/* API to get  DHCP and DNS center server details.


@param GetNetworkQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network
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
		return nil, response, fmt.Errorf("error with operation GetNetwork")
	}

	result := response.Result().(*ResponseNetworkSettingsGetNetwork)
	return result, response, err

}

//GetReserveIPSubpool Get Reserve IP Subpool - 4586-0917-4fab-87e2
/* API to get the ip subpool info.


@param GetReserveIPSubpoolQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-reserve-ip-subpool
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
		return nil, response, fmt.Errorf("error with operation GetReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsGetReserveIPSubpool)
	return result, response, err

}

//GetServiceProviderDetails Get Service provider Details - 7084-7bdc-4d89-a437
/* API to get service provider details (QoS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-service-provider-details
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
		return nil, response, fmt.Errorf("error with operation GetServiceProviderDetails")
	}

	result := response.Result().(*ResponseNetworkSettingsGetServiceProviderDetails)
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
		return nil, response, fmt.Errorf("error with operation GetServiceProviderDetailsV2")
	}

	result := response.Result().(*ResponseNetworkSettingsGetServiceProviderDetailsV2)
	return result, response, err

}

//AssignDeviceCredentialToSite Assign Device Credential To Site. - 4da9-1a54-4e29-842d
/* Assign Device Credential to a site.


@param siteID siteId path parameter. site id to assign credential.

@param AssignDeviceCredentialToSiteHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-device-credential-to-site
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
		return nil, response, fmt.Errorf("error with operation AssignDeviceCredentialToSite")
	}

	result := response.Result().(*ResponseNetworkSettingsAssignDeviceCredentialToSite)
	return result, response, err

}

//CreateDeviceCredentials Create Device Credentials - fbb9-5b37-484a-9fce
/* API to create device credentials.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-device-credentials
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
		return nil, response, fmt.Errorf("error with operation CreateDeviceCredentials")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateDeviceCredentials)
	return result, response, err

}

//CreateGlobalPool Create Global Pool - f793-192a-43da-bed9
/* API to create global pool.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-global-pool
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
		return nil, response, fmt.Errorf("error with operation CreateGlobalPool")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateGlobalPool)
	return result, response, err

}

//CreateNetwork Create Network - be89-2bd8-4a78-865a
/* API to create a network for DHCP,  Syslog, SNMP, NTP, Network AAA, Client and EndPoint AAA, and/or DNS center server settings.


@param siteID siteId path parameter. Site id to which site details to associate with the network settings.

@param CreateNetworkHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-network
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
		return nil, response, fmt.Errorf("error with operation CreateNetwork")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateNetwork)
	return result, response, err

}

//ReserveIPSubpool Reserve IP Subpool - 429f-aa81-4d3b-960a
/* API to reserve an ip subpool from the global pool


@param siteID siteId path parameter. Site id to reserve the ip sub pool.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!reserve-ip-subpool
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
		return nil, response, fmt.Errorf("error with operation ReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsReserveIPSubpool)
	return result, response, err

}

//CreateSpProfile Create SP Profile - a39a-1a21-4deb-b781
/* API to create Service Provider Profile(QOS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-s-p-profile
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
		return nil, response, fmt.Errorf("error with operation CreateSpProfile")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateSpProfile)
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
		return nil, response, fmt.Errorf("error with operation CreateNetworkV2")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateNetworkV2)
	return result, response, err

}

//CreateSpProfileV2 Create SP Profile V2 - b6b1-1aa9-4b4b-99e0
/* API to create Service Provider Profile(QOS).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-s-p-profile-v2
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
		return nil, response, fmt.Errorf("error with operation CreateSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsCreateSpProfileV2)
	return result, response, err

}

//UpdateDeviceCredentials Update Device Credentials - 4f94-7a1c-4fc8-84f6
/* API to update device credentials.


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
		return nil, response, fmt.Errorf("error with operation UpdateSpProfile")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateSpProfile)
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
		return nil, response, fmt.Errorf("error with operation UpdateSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsUpdateSpProfileV2)
	return result, response, err

}

//DeleteDeviceCredential Delete Device Credential - 259e-ab30-4598-8958
/* Delete device credential.


@param id id path parameter. global credential id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-device-credential
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
		return nil, response, fmt.Errorf("error with operation DeleteDeviceCredential")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteDeviceCredential)
	return result, response, err

}

//DeleteGlobalIPPool Delete Global IP Pool - 1eaa-8b21-48ab-81de
/* API to delete global IP pool.


@param id id path parameter. global pool id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-global-ip-pool
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
		return nil, response, fmt.Errorf("error with operation DeleteGlobalIpPool")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteGlobalIPPool)
	return result, response, err

}

//ReleaseReserveIPSubpool Release Reserve IP Subpool - 85b2-89e3-4489-9dc1
/* API to delete the reserved ip subpool


@param id id path parameter. Id of reserve ip subpool to be deleted.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!release-reserve-ip-subpool
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
		return nil, response, fmt.Errorf("error with operation ReleaseReserveIpSubpool")
	}

	result := response.Result().(*ResponseNetworkSettingsReleaseReserveIPSubpool)
	return result, response, err

}

//DeleteSpProfile Delete SP Profile - 4ca2-db11-43eb-b5d7
/* API to delete Service Provider Profile (QoS).


@param spProfileName spProfileName path parameter. sp profile name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-s-p-profile
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
		return nil, response, fmt.Errorf("error with operation DeleteSpProfile")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteSpProfile)
	return result, response, err

}

//DeleteSpProfileV2 Delete SP Profile V2 - 5ea0-cb0e-4ed9-9bec
/* API to delete Service Provider Profile (QoS).


@param spProfileName spProfileName path parameter. sp profile name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-s-p-profile-v2
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
		return nil, response, fmt.Errorf("error with operation DeleteSpProfileV2")
	}

	result := response.Result().(*ResponseNetworkSettingsDeleteSpProfileV2)
	return result, response, err

}
