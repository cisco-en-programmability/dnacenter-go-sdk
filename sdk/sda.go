package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SdaService service

type GetDefaultAuthenticationProfileFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //siteNameHierarchy
}
type DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //siteNameHierarchy
}
type GetsBorderDeviceDetailFromSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type DeletesBorderDeviceFromSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type DeleteControlPlaneDeviceInSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type GetControlPlaneDeviceFromSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type GetDeviceInfoFromSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type GetDeviceRoleInSdaFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //Device Management IP Address
}
type DeleteEdgeDeviceFromSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type GetEdgeDeviceFromSdaFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device IP Address
}
type DeleteSdaFabricQueryParams struct {
	FabricName string `url:"fabricName,omitempty"` //Fabric Name
}
type GetSdaFabricInfoQueryParams struct {
	FabricName string `url:"fabricName,omitempty"` //Fabric Name
}
type GetSiteFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Hierarchy
}
type DeleteSiteFromSdaFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Hierarchy
}
type DeletePortAssignmentForAccessPointInSdaFabricQueryParams struct {
	DeviceIP      string `url:"device-ip,omitempty"`     //device-ip
	InterfaceName string `url:"interfaceName,omitempty"` //interfaceName
}
type GetPortAssignmentForAccessPointInSdaFabricQueryParams struct {
	DeviceIP      string `url:"device-ip,omitempty"`     //device-ip
	InterfaceName string `url:"interfaceName,omitempty"` //interfaceName
}
type DeletePortAssignmentForUserDeviceInSdaFabricQueryParams struct {
	DeviceIP      string `url:"device-ip,omitempty"`     //device-ip
	InterfaceName string `url:"interfaceName,omitempty"` //interfaceName
}
type GetPortAssignmentForUserDeviceInSdaFabricQueryParams struct {
	DeviceIP      string `url:"device-ip,omitempty"`     //device-ip
	InterfaceName string `url:"interfaceName,omitempty"` //interfaceName
}
type GetMulticastDetailsFromSdaFabricQueryParams struct {
	FabricSiteNameHierarchy string `url:"fabricSiteNameHierarchy,omitempty"` //fabricSiteNameHierarchy
}
type DeleteMulticastFromSdaFabricQueryParams struct {
	FabricSiteNameHierarchy string `url:"fabricSiteNameHierarchy,omitempty"` //fabricSiteNameHierarchy
}
type DeleteProvisionedWiredDeviceQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //Valid IP address of the device currently provisioned in a fabric site
}
type GetProvisionedWiredDeviceQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` //deviceManagementIpAddress
}
type DeleteVnFromSdaFabricQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
}
type GetVnFromSdaFabricQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  //siteNameHierarchy
}
type GetIPPoolFromSdaVirtualNetworkQueryParams struct {
	IPPoolName         string `url:"ipPoolName,omitempty"`         //ipPoolName
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}
type DeleteIPPoolFromSdaVirtualNetworkQueryParams struct {
	IPPoolName         string `url:"ipPoolName,omitempty"`         //ipPoolName
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}
type DeleteVirtualNetworkWithScalableGroupsQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}
type GetVirtualNetworkWithScalableGroupsQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` //virtualNetworkName
}

type ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric struct {
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        // Site Name Hierarchy
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
	AuthenticateTemplateID   string `json:"authenticateTemplateId,omitempty"`   // Authenticate Template Id
}
type ResponseSdaAddDefaultAuthenticationProfileInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaAddsBorderDeviceInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabric struct {
	Status      string                                                 `json:"status,omitempty"`      // Status
	Description string                                                 `json:"description,omitempty"` // Description
	Payload     *ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayload `json:"payload,omitempty"`     //
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayload struct {
	ID                             string                                                                      `json:"id,omitempty"`                             // Id
	InstanceID                     *int                                                                        `json:"instanceId,omitempty"`                     // Instance Id
	AuthEntityID                   *int                                                                        `json:"authEntityId,omitempty"`                   // Auth Entity Id
	DisplayName                    string                                                                      `json:"displayName,omitempty"`                    // Display Name
	AuthEntityClass                *int                                                                        `json:"authEntityClass,omitempty"`                // Auth Entity Class
	InstanceTenantID               string                                                                      `json:"instanceTenantId,omitempty"`               // Instance Tenant Id
	DeployPending                  string                                                                      `json:"deployPending,omitempty"`                  // Deploy Pending
	InstanceVersion                *int                                                                        `json:"instanceVersion,omitempty"`                // Instance Version
	CreateTime                     *int                                                                        `json:"createTime,omitempty"`                     // Create Time
	Deployed                       *bool                                                                       `json:"deployed,omitempty"`                       // Deployed
	IsSeeded                       *bool                                                                       `json:"isSeeded,omitempty"`                       // Is Seeded
	IsStale                        *bool                                                                       `json:"isStale,omitempty"`                        // Is Stale
	LastUpdateTime                 *int                                                                        `json:"lastUpdateTime,omitempty"`                 // Last Update Time
	Name                           string                                                                      `json:"name,omitempty"`                           // Name
	Namespace                      string                                                                      `json:"namespace,omitempty"`                      // Namespace
	ProvisioningState              string                                                                      `json:"provisioningState,omitempty"`              // Provisioning State
	ResourceVersion                *int                                                                        `json:"resourceVersion,omitempty"`                // Resource Version
	TargetIDList                   *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadTargetIDList        `json:"targetIdList,omitempty"`                   // Target Id List
	Type                           string                                                                      `json:"type,omitempty"`                           // Type
	CfsChangeInfo                  *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadCfsChangeInfo       `json:"cfsChangeInfo,omitempty"`                  // Cfs Change Info
	CustomProvisions               *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadCustomProvisions    `json:"customProvisions,omitempty"`               // Custom Provisions
	Configs                        *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadConfigs             `json:"configs,omitempty"`                        // Configs
	ManagedSites                   *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadManagedSites        `json:"managedSites,omitempty"`                   // Managed Sites
	NetworkDeviceID                string                                                                      `json:"networkDeviceId,omitempty"`                // Network Device Id
	Roles                          []string                                                                    `json:"roles,omitempty"`                          // Roles
	SaveWanConnectivityDetailsOnly *bool                                                                       `json:"saveWanConnectivityDetailsOnly,omitempty"` // Save Wan Connectivity Details Only
	SiteID                         string                                                                      `json:"siteId,omitempty"`                         // Site Id
	AkcSettingsCfs                 *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadAkcSettingsCfs      `json:"akcSettingsCfs,omitempty"`                 // Akc Settings Cfs
	DeviceInterfaceInfo            *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceInterfaceInfo `json:"deviceInterfaceInfo,omitempty"`            // Device Interface Info
	DeviceSettings                 *ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettings        `json:"deviceSettings,omitempty"`                 //
	NetworkWideSettings            *ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettings   `json:"networkWideSettings,omitempty"`            //
	OtherDevice                    *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadOtherDevice         `json:"otherDevice,omitempty"`                    // Other Device
	TransitNetworks                *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadTransitNetworks     `json:"transitNetworks,omitempty"`                //
	VirtualNetwork                 *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadVirtualNetwork      `json:"virtualNetwork,omitempty"`                 // Virtual Network
	WLAN                           *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadWLAN                `json:"wlan,omitempty"`                           // Wlan
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadTargetIDList interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadCfsChangeInfo interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadCustomProvisions interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadConfigs interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadManagedSites interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadAkcSettingsCfs interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceInterfaceInfo interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettings struct {
	ID                            string                                                                                        `json:"id,omitempty"`                            // Id
	InstanceID                    *int                                                                                          `json:"instanceId,omitempty"`                    // Instance Id
	DisplayName                   string                                                                                        `json:"displayName,omitempty"`                   // Display Name
	InstanceTenantID              string                                                                                        `json:"instanceTenantId,omitempty"`              // Instance Tenant Id
	DeployPending                 string                                                                                        `json:"deployPending,omitempty"`                 // Deploy Pending
	InstanceVersion               *int                                                                                          `json:"instanceVersion,omitempty"`               // Instance Version
	ConnectedTo                   *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsConnectedTo             `json:"connectedTo,omitempty"`                   // Connected To
	CPU                           *float64                                                                                      `json:"cpu,omitempty"`                           // Cpu
	DhcpEnabled                   *bool                                                                                         `json:"dhcpEnabled,omitempty"`                   // Dhcp Enabled
	ExternalConnectivityIPPool    string                                                                                        `json:"externalConnectivityIpPool,omitempty"`    // External Connectivity Ip Pool
	ExternalDomainRoutingProtocol string                                                                                        `json:"externalDomainRoutingProtocol,omitempty"` // External Domain Routing Protocol
	InternalDomainProtocolNumber  string                                                                                        `json:"internalDomainProtocolNumber,omitempty"`  // Internal Domain Protocol Number
	Memory                        *float64                                                                                      `json:"memory,omitempty"`                        // Memory
	NodeType                      []string                                                                                      `json:"nodeType,omitempty"`                      // Node Type
	Storage                       *float64                                                                                      `json:"storage,omitempty"`                       // Storage
	ExtConnectivitySettings       *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettings `json:"extConnectivitySettings,omitempty"`       //
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsConnectedTo interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettings struct {
	ID                           string                                                                                                 `json:"id,omitempty"`                           // Id
	InstanceID                   *int                                                                                                   `json:"instanceId,omitempty"`                   // Instance Id
	DisplayName                  string                                                                                                 `json:"displayName,omitempty"`                  // Display Name
	InstanceTenantID             string                                                                                                 `json:"instanceTenantId,omitempty"`             // Instance Tenant Id
	DeployPending                string                                                                                                 `json:"deployPending,omitempty"`                // Deploy Pending
	InstanceVersion              *int                                                                                                   `json:"instanceVersion,omitempty"`              // Instance Version
	ExternalDomainProtocolNumber string                                                                                                 `json:"externalDomainProtocolNumber,omitempty"` // External Domain Protocol Number
	InterfaceUUID                string                                                                                                 `json:"interfaceUuid,omitempty"`                // Interface Uuid
	PolicyPropagationEnabled     *bool                                                                                                  `json:"policyPropagationEnabled,omitempty"`     // Policy Propagation Enabled
	PolicySgtTag                 *float64                                                                                               `json:"policySgtTag,omitempty"`                 // Policy Sgt Tag
	L2Handoff                    *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL2Handoff `json:"l2Handoff,omitempty"`                    // L2 Handoff
	L3Handoff                    *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"`                    //
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL2Handoff interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3Handoff struct {
	ID               string                                                                                                             `json:"id,omitempty"`               // Id
	InstanceID       *int                                                                                                               `json:"instanceId,omitempty"`       // Instance Id
	DisplayName      string                                                                                                             `json:"displayName,omitempty"`      // Display Name
	InstanceTenantID string                                                                                                             `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	DeployPending    string                                                                                                             `json:"deployPending,omitempty"`    // Deploy Pending
	InstanceVersion  *float64                                                                                                           `json:"instanceVersion,omitempty"`  // Instance Version
	LocalIPAddress   string                                                                                                             `json:"localIpAddress,omitempty"`   // Local Ip Address
	RemoteIPAddress  string                                                                                                             `json:"remoteIpAddress,omitempty"`  // Remote Ip Address
	VLANID           *int                                                                                                               `json:"vlanId,omitempty"`           // Vlan Id
	VirtualNetwork   *ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"`   //
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettings struct {
	ID               string                                                                                `json:"id,omitempty"`               // Id
	InstanceID       *int                                                                                  `json:"instanceId,omitempty"`       // Instance Id
	DisplayName      string                                                                                `json:"displayName,omitempty"`      // Display Name
	InstanceTenantID string                                                                                `json:"instanceTenantId,omitempty"` // Instance Tenant Id
	DeployPending    string                                                                                `json:"deployPending,omitempty"`    // Deploy Pending
	InstanceVersion  *int                                                                                  `json:"instanceVersion,omitempty"`  // Instance Version
	AAA              *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsAAA        `json:"aaa,omitempty"`              // Aaa
	Cmx              *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsCmx        `json:"cmx,omitempty"`              // Cmx
	Dhcp             *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcp       `json:"dhcp,omitempty"`             //
	DNS              *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNS        `json:"dns,omitempty"`              //
	Ldap             *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsLdap       `json:"ldap,omitempty"`             // Ldap
	NativeVLAN       *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNativeVLAN `json:"nativeVlan,omitempty"`       // Native Vlan
	Netflow          *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNetflow    `json:"netflow,omitempty"`          // Netflow
	Ntp              *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNtp        `json:"ntp,omitempty"`              // Ntp
	SNMP             *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSNMP       `json:"snmp,omitempty"`             // Snmp
	Syslogs          *[]ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSyslogs    `json:"syslogs,omitempty"`          // Syslogs
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsAAA interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsCmx interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcp struct {
	ID        string                                                                                 `json:"id,omitempty"`        // Id
	IPAddress *ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcpIPAddress `json:"ipAddress,omitempty"` //
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDhcpIPAddress struct {
	ID            string `json:"id,omitempty"`            // Id
	PaddedAddress string `json:"paddedAddress,omitempty"` // Padded Address
	AddressType   string `json:"addressType,omitempty"`   // Address Type
	Address       string `json:"address,omitempty"`       // Address
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNS struct {
	ID         string                                                                         `json:"id,omitempty"`         // Id
	DomainName string                                                                         `json:"domainName,omitempty"` // Domain Name
	IP         *ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNSIP `json:"ip,omitempty"`         //
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsDNSIP struct {
	ID            string `json:"id,omitempty"`            // Id
	PaddedAddress string `json:"paddedAddress,omitempty"` // Padded Address
	AddressType   string `json:"addressType,omitempty"`   // Address Type
	Address       string `json:"address,omitempty"`       // Address
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsLdap interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNativeVLAN interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNetflow interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsNtp interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSNMP interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadNetworkWideSettingsSyslogs interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadOtherDevice interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadTransitNetworks struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadVirtualNetwork interface{}
type ResponseSdaGetsBorderDeviceDetailFromSdaFabricPayloadWLAN interface{}
type ResponseSdaDeletesBorderDeviceFromSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaDeleteControlPlaneDeviceInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetControlPlaneDeviceFromSdaFabric struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaAddControlPlaneDeviceInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetDeviceInfoFromSdaFabric struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaGetDeviceRoleInSdaFabric struct {
	Response *ResponseSdaGetDeviceRoleInSdaFabricResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseSdaGetDeviceRoleInSdaFabricResponse struct {
	Status      string   `json:"status,omitempty"`      // Status
	Description string   `json:"description,omitempty"` // Description
	Roles       []string `json:"roles,omitempty"`       // Roles
}
type ResponseSdaAddEdgeDeviceInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaDeleteEdgeDeviceFromSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetEdgeDeviceFromSdaFabric struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaDeleteSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetSdaFabricInfo struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaAddFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetSiteFromSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaDeleteSiteFromSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaAddSiteInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetSdaFabricCount struct {
	Response *ResponseSdaGetSdaFabricCountResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  // Version
}
type ResponseSdaGetSdaFabricCountResponse struct {
	Status      string `json:"status,omitempty"`      // Status
	Description string `json:"description,omitempty"` // Description
	FabricCount string `json:"fabricCount,omitempty"` // Fabric Count
}
type ResponseSdaAddPortAssignmentForAccessPointInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetPortAssignmentForAccessPointInSdaFabric struct {
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Description
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Data Ip Address Pool Name
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    // Voice Ip Address Pool Name
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         // Scalable Group Name
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
}
type ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaAddPortAssignmentForUserDeviceInSdaFabric struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric struct {
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Description
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Data Ip Address Pool Name
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    // Voice Ip Address Pool Name
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         // Scalable Group Name
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
}
type ResponseSdaAddMulticastInSdaFabric struct {
	Status        string `json:"status,omitempty"`        // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description   string `json:"description,omitempty"`   // gives the api status description and reason for its success or failure.
	TaskID        string `json:"taskId,omitempty"`        // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID   string `json:"executionId,omitempty"`   // /dna/intent/api/v1/execution-status/<executionId>
}
type ResponseSdaGetMulticastDetailsFromSdaFabric struct {
	FabricSiteNameHierarchy string                                                      `json:"fabricSiteNameHierarchy,omitempty"` // fabricSiteNameHierarchy
	MulticastMethod         string                                                      `json:"multicastMethod,omitempty"`         // multicast methods
	MuticastType            string                                                      `json:"muticastType,omitempty"`            // muticast type
	MulticastVnInfo         *ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfo `json:"multicastVnInfo,omitempty"`         //
}
type ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfo struct {
	VirtualNetworkName  string                                                             `json:"virtualNetworkName,omitempty"`  // virtual network name in fabricSiteNameHierarchy
	IPPoolName          string                                                             `json:"ipPoolName,omitempty"`          // Ip Pool Name in fabricSiteNameHierarchy
	ExternalRpIPAddress string                                                             `json:"externalRpIpAddress,omitempty"` // External Rp Ip Address, required for muticastType=asm_with_external_rp
	SsmInfo             *ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfoSsmInfo `json:"ssmInfo,omitempty"`             // Source-specific multicast information, required if muticastType=ssm
	SsmGroupRange       string                                                             `json:"ssmGroupRange,omitempty"`       // valid ssm group range ip address
	SsmWildcardMask     string                                                             `json:"ssmWildcardMask,omitempty"`     // valid ssm Wildcard Mask ip address
}
type ResponseSdaGetMulticastDetailsFromSdaFabricMulticastVnInfoSsmInfo interface{}
type ResponseSdaDeleteMulticastFromSdaFabric struct {
	Status        string `json:"status,omitempty"`        // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description   string `json:"description,omitempty"`   // gives the api status description and reason for its success or failure.
	TaskID        string `json:"taskId,omitempty"`        // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID   string `json:"executionId,omitempty"`   // /dna/intent/api/v1/execution-status/<executionId>
}
type ResponseSdaDeleteProvisionedWiredDevice struct {
	Status        string `json:"status,omitempty"`        // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description   string `json:"description,omitempty"`   // gives the api status description and reason for its success or failure.
	TaskID        string `json:"taskId,omitempty"`        // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID   string `json:"executionId,omitempty"`   // /dna/intent/api/v1/execution-status/<executionId>
}
type ResponseSdaProvisionWiredDevice struct {
	Status        string `json:"status,omitempty"`        // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description   string `json:"description,omitempty"`   // gives the api status description and reason for its success or failure.
	TaskID        string `json:"taskId,omitempty"`        // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID   string `json:"executionId,omitempty"`   // /dna/intent/api/v1/execution-status/<executionId>
}
type ResponseSdaGetProvisionedWiredDevice struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	Status                    string `json:"status,omitempty"`                    // Status
	Description               string `json:"description,omitempty"`               // Description
	TaskID                    string `json:"taskId,omitempty"`                    // Task Id
	TaskStatusURL             string `json:"taskStatusUrl,omitempty"`             // Task Status Url
	ExecutionStatusURL        string `json:"executionStatusUrl,omitempty"`        // Execution Status Url
	ExecutionID               string `json:"executionId,omitempty"`               // Execution Id
	SiteID                    string `json:"siteId,omitempty"`                    // Site Id
}
type ResponseSdaDeleteVnFromSdaFabric struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaGetVnFromSdaFabric struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaAddVnInSdaFabric struct {
	Status                    string   `json:"status,omitempty"`                    // Status
	Description               string   `json:"description,omitempty"`               // Description
	Name                      string   `json:"name,omitempty"`                      // Name
	Roles                     []string `json:"roles,omitempty"`                     // Roles
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             // Site Hierarchy
}
type ResponseSdaGetIPPoolFromSdaVirtualNetwork struct {
	Status                   string `json:"status,omitempty"`                   // Status
	Description              string `json:"description,omitempty"`              // Description
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       // Virtual Network Name
	IPPoolName               string `json:"ipPoolName,omitempty"`               // Ip Pool Name
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` // Authentication Policy Name
	TrafficType              string `json:"trafficType,omitempty"`              // Traffic Type
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        // Scalable Group Name
	IsL2FloodingEnabled      *bool  `json:"isL2FloodingEnabled,omitempty"`      // Is L2 Flooding Enabled
	IsThisCriticalPool       *bool  `json:"isThisCriticalPool,omitempty"`       // Is This Critical Pool
}
type ResponseSdaDeleteIPPoolFromSdaVirtualNetwork struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaAddIPPoolInSdaVirtualNetwork struct {
	Status             string `json:"status,omitempty"`             // Status
	Description        string `json:"description,omitempty"`        // Description
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
}
type ResponseSdaAddVirtualNetworkWithScalableGroups struct {
	Status        string `json:"status,omitempty"`        // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description   string `json:"description,omitempty"`   // gives the api status description and reason for its success or failure.
	TaskID        string `json:"taskId,omitempty"`        // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID   string `json:"executionId,omitempty"`   // /dna/intent/api/v1/execution-status/<executionId>
}
type ResponseSdaDeleteVirtualNetworkWithScalableGroups struct {
	Status             string `json:"status,omitempty"`             // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description        string `json:"description,omitempty"`        // gives the api status description and reason for its success or failure.
	TaskID             string `json:"taskId,omitempty"`             // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL      string `json:"taskStatusUrl,omitempty"`      // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID        string `json:"executionId,omitempty"`        // DNA-Center executionId that carried out the API execution. It will be provided if execution was generated by api. For failed status, executionId may not be available
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // url to track the api progression via executionId /dna/intent/api/v1/execution-statu/<executionId>
}
type ResponseSdaGetVirtualNetworkWithScalableGroups struct {
	VirtualNetworkName    string   `json:"virtualNetworkName,omitempty"`    // Virtual Network Name
	IsGuestVirtualNetwork *bool    `json:"isGuestVirtualNetwork,omitempty"` // Is Guest Virtual Network
	ScalableGroupNames    []string `json:"scalableGroupNames,omitempty"`    // Scalable Group Names
	Status                string   `json:"status,omitempty"`                // Status
	Description           string   `json:"description,omitempty"`           // Description
	TaskID                string   `json:"taskId,omitempty"`                // Task Id
	TaskStatusURL         string   `json:"taskStatusUrl,omitempty"`         // Task Status Url
	ExecutionStatusURL    string   `json:"executionStatusUrl,omitempty"`    // Execution Status Url
	ExecutionID           string   `json:"executionId,omitempty"`           // Execution Id
}
type ResponseSdaUpdateVirtualNetworkWithScalableGroups struct {
	Status        string `json:"status,omitempty"`        // represents return status of api. status=success when api competed successfully, status=failed when api failed and has not completed the user request, status=pending when api execution is still in progression and user needs to track its further progress via taskId field.
	Description   string `json:"description,omitempty"`   // gives the api status description and reason for its success or failure.
	TaskID        string `json:"taskId,omitempty"`        // DNA-Center taskId that carried out the API execution. It will be provided if task was generated by api. For failed status, taskId may not be available
	TaskStatusURL string `json:"taskStatusUrl,omitempty"` // url to track the api progression via taskId /dna/intent/api/v1/task/<taskId>
	ExecutionID   string `json:"executionId,omitempty"`   // /dna/intent/api/v1/execution-status/<executionId>
}
type RequestSdaAddDefaultAuthenticationProfileInSdaFabric []RequestItemSdaAddDefaultAuthenticationProfileInSdaFabric // Array of RequestSdaAddDefaultAuthenticationProfileInSDAFabric
type RequestItemSdaAddDefaultAuthenticationProfileInSdaFabric struct {
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        // Site Name Hierarchy
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
}
type RequestSdaUpdateDefaultAuthenticationProfileInSdaFabric []RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric // Array of RequestSdaUpdateDefaultAuthenticationProfileInSDAFabric
type RequestItemSdaUpdateDefaultAuthenticationProfileInSdaFabric struct {
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        // Site Name Hierarchy
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` // Authenticate Template Name
}
type RequestSdaAddsBorderDeviceInSdaFabric []RequestItemSdaAddsBorderDeviceInSdaFabric // Array of RequestSdaAddsBorderDeviceInSDAFabric
type RequestItemSdaAddsBorderDeviceInSdaFabric struct {
	DeviceManagementIPAddress         string                                                                   `json:"deviceManagementIpAddress,omitempty"`         // Device Management Ip Address
	SiteNameHierarchy                 string                                                                   `json:"siteNameHierarchy,omitempty"`                 // Site Name Hierarchy
	ExternalDomainRoutingProtocolName string                                                                   `json:"externalDomainRoutingProtocolName,omitempty"` // External Domain Routing Protocol Name
	ExternalConnectivityIPPoolName    string                                                                   `json:"externalConnectivityIpPoolName,omitempty"`    // External Connectivity Ip Pool Name
	InternalAutonomouSystemNumber     string                                                                   `json:"internalAutonomouSystemNumber,omitempty"`     // Internal Autonomou System Number
	BorderSessionType                 string                                                                   `json:"borderSessionType,omitempty"`                 // Border Session Type
	ConnectedToInternet               *bool                                                                    `json:"connectedToInternet,omitempty"`               // Connected To Internet
	ExternalConnectivitySettings      *[]RequestItemSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettings `json:"externalConnectivitySettings,omitempty"`      //
}
type RequestItemSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettings struct {
	InterfaceName                 string                                                                            `json:"interfaceName,omitempty"`                 // Interface Name
	ExternalAutonomouSystemNumber string                                                                            `json:"externalAutonomouSystemNumber,omitempty"` // External Autonomou System Number
	L3Handoff                     *[]RequestItemSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"`                     //
}
type RequestItemSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff struct {
	VirtualNetwork *RequestItemSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"` //
}
type RequestItemSdaAddsBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name
}
type RequestSdaAddControlPlaneDeviceInSdaFabric []RequestItemSdaAddControlPlaneDeviceInSdaFabric // Array of RequestSdaAddControlPlaneDeviceInSDAFabric
type RequestItemSdaAddControlPlaneDeviceInSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
}
type RequestSdaAddEdgeDeviceInSdaFabric []RequestItemSdaAddEdgeDeviceInSdaFabric // Array of RequestSdaAddEdgeDeviceInSDAFabric
type RequestItemSdaAddEdgeDeviceInSdaFabric struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
}
type RequestSdaAddFabric []RequestItemSdaAddFabric // Array of RequestSdaAddFabric
type RequestItemSdaAddFabric struct {
	FabricName string `json:"fabricName,omitempty"` // Fabric Name
}
type RequestSdaAddSiteInSdaFabric []RequestItemSdaAddSiteInSdaFabric // Array of RequestSdaAddSiteInSDAFabric
type RequestItemSdaAddSiteInSdaFabric struct {
	FabricName        string `json:"fabricName,omitempty"`        // Fabric Name
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type RequestSdaAddPortAssignmentForAccessPointInSdaFabric []RequestItemSdaAddPortAssignmentForAccessPointInSdaFabric // Array of RequestSdaAddPortAssignmentForAccessPointInSDAFabric
type RequestItemSdaAddPortAssignmentForAccessPointInSdaFabric struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Data Ip Address Pool Name
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    // Voice Ip Address Pool Name
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
}
type RequestSdaAddPortAssignmentForUserDeviceInSdaFabric []RequestItemSdaAddPortAssignmentForUserDeviceInSdaFabric // Array of RequestSdaAddPortAssignmentForUserDeviceInSDAFabric
type RequestItemSdaAddPortAssignmentForUserDeviceInSdaFabric struct {
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	InterfaceName             string `json:"interfaceName,omitempty"`             // Interface Name
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     // Data Ip Address Pool Name
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    // Voice Ip Address Pool Name
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  // Authenticate Template Name
}
type RequestSdaAddMulticastInSdaFabric struct {
	FabricSiteNameHierarchy string                                            `json:"fabricSiteNameHierarchy,omitempty"` // fabricSiteNameHierarchy
	MulticastMethod         string                                            `json:"multicastMethod,omitempty"`         // multicast methods
	MuticastType            string                                            `json:"muticastType,omitempty"`            // muticast type
	MulticastVnInfo         *RequestSdaAddMulticastInSdaFabricMulticastVnInfo `json:"multicastVnInfo,omitempty"`         //
}
type RequestSdaAddMulticastInSdaFabricMulticastVnInfo struct {
	VirtualNetworkName  string                                                   `json:"virtualNetworkName,omitempty"`  // virtual network name in fabricSiteNameHierarchy
	IPPoolName          string                                                   `json:"ipPoolName,omitempty"`          // Ip Pool Name in fabricSiteNameHierarchy
	ExternalRpIPAddress string                                                   `json:"externalRpIpAddress,omitempty"` // External Rp Ip Address, required for muticastType=asm_with_external_rp
	SsmInfo             *RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo `json:"ssmInfo,omitempty"`             // Source-specific multicast information, required if muticastType=ssm
	SsmGroupRange       string                                                   `json:"ssmGroupRange,omitempty"`       // valid ssm group range ip address
	SsmWildcardMask     string                                                   `json:"ssmWildcardMask,omitempty"`     // valid ssm Wildcard Mask ip address
}
type RequestSdaAddMulticastInSdaFabricMulticastVnInfoSsmInfo interface{}
type RequestSdaProvisionWiredDevice struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` // Device Management Ip Address
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         // Site Name Hierarchy
}
type RequestSdaAddVnInSdaFabric []RequestItemSdaAddVnInSdaFabric // Array of RequestSdaAddVNInSDAFabric
type RequestItemSdaAddVnInSdaFabric struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` // Virtual Network Name
	SiteNameHierarchy  string `json:"siteNameHierarchy,omitempty"`  // Site Name Hierarchy
}
type RequestSdaAddIPPoolInSdaVirtualNetwork []RequestItemSdaAddIPPoolInSdaVirtualNetwork // Array of RequestSdaAddIPPoolInSDAVirtualNetwork
type RequestItemSdaAddIPPoolInSdaVirtualNetwork struct {
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       // Virtual Network Name
	IPPoolName               string `json:"ipPoolName,omitempty"`               // Ip Pool Name
	TrafficType              string `json:"trafficType,omitempty"`              // Traffic Type
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` // Authentication Policy Name
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        // Scalable Group Name
	IsL2FloodingEnabled      *bool  `json:"isL2FloodingEnabled,omitempty"`      // Is L2 Flooding Enabled
	IsThisCriticalPool       *bool  `json:"isThisCriticalPool,omitempty"`       // Is This Critical Pool
	PoolType                 string `json:"poolType,omitempty"`                 // Pool Type
}
type RequestSdaAddVirtualNetworkWithScalableGroups struct {
	VirtualNetworkName    string   `json:"virtualNetworkName,omitempty"`    // Virtual Network Name
	VirtualNetworkType    string   `json:"virtualNetworkType,omitempty"`    // Virtual Network Type
	ScalableGroupNames    []string `json:"scalableGroupNames,omitempty"`    // Scalable Group Names
	IsGuestVirtualNetwork *bool    `json:"isGuestVirtualNetwork,omitempty"` //
}
type RequestSdaUpdateVirtualNetworkWithScalableGroups struct {
	VirtualNetworkName    string   `json:"virtualNetworkName,omitempty"`    // Virtual Network Name
	VirtualNetworkType    string   `json:"virtualNetworkType,omitempty"`    // Virtual Network Type
	ScalableGroupNames    []string `json:"scalableGroupNames,omitempty"`    // Scalable Group Names
	IsGuestVirtualNetwork *bool    `json:"isGuestVirtualNetwork,omitempty"` // Is Guest Virtual Network
}

//GetDefaultAuthenticationProfileFromSdaFabric Get default authentication profile from SDA Fabric - 8b90-8a4e-4c5a-9a23
/* Get default authentication profile from SDA Fabric


@param GetDefaultAuthenticationProfileFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetDefaultAuthenticationProfileFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric)
	return result, response, err

}

//GetsBorderDeviceDetailFromSdaFabric Gets border device detail from SDA Fabric - 98a3-9bf4-485a-9871
/* Gets border device detail from SDA Fabric


@param GetsBorderDeviceDetailFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) GetsBorderDeviceDetailFromSdaFabric(GetsBorderDeviceDetailFromSDAFabricQueryParams *GetsBorderDeviceDetailFromSdaFabricQueryParams) (*ResponseSdaGetsBorderDeviceDetailFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(GetsBorderDeviceDetailFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetsBorderDeviceDetailFromSdaFabric{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetsBorderDeviceDetailFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetsBorderDeviceDetailFromSdaFabric)
	return result, response, err

}

//GetControlPlaneDeviceFromSdaFabric Get control plane device from SDA Fabric - aba4-991d-4e9b-8747
/* Get control plane device from SDA Fabric


@param GetControlPlaneDeviceFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetControlPlaneDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetControlPlaneDeviceFromSdaFabric)
	return result, response, err

}

//GetDeviceInfoFromSdaFabric Get device info from SDA Fabric - 1385-18e1-4069-ab5f
/* Get device info from SDA Fabric


@param GetDeviceInfoFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetDeviceInfoFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetDeviceInfoFromSdaFabric)
	return result, response, err

}

//GetDeviceRoleInSdaFabric Get device role in SDA Fabric - 8a92-d87c-416a-8e83
/* Get device role in SDA Fabric


@param GetDeviceRoleInSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetDeviceRoleInSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetDeviceRoleInSdaFabric)
	return result, response, err

}

//GetEdgeDeviceFromSdaFabric Get edge device from SDA Fabric - 7683-f90b-4efa-b090
/* Get edge device from SDA Fabric


@param GetEdgeDeviceFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetEdgeDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetEdgeDeviceFromSdaFabric)
	return result, response, err

}

//GetSdaFabricInfo Get SDA Fabric Info - 16a1-bb5d-48cb-873d
/* Get SDA Fabric Info


@param GetSDAFabricInfoQueryParams Filtering parameter
*/
func (s *SdaService) GetSdaFabricInfo(GetSDAFabricInfoQueryParams *GetSdaFabricInfoQueryParams) (*ResponseSdaGetSdaFabricInfo, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric"

	queryString, _ := query.Values(GetSDAFabricInfoQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaGetSdaFabricInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSdaFabricInfo")
	}

	result := response.Result().(*ResponseSdaGetSdaFabricInfo)
	return result, response, err

}

//GetSiteFromSdaFabric Get Site from SDA Fabric - 80b7-f8e6-406a-8701
/* Get Site info from SDA Fabric


@param GetSiteFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetSiteFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetSiteFromSdaFabric)
	return result, response, err

}

//GetSdaFabricCount Get SDA Fabric Count - 6fa0-f8d5-4d29-857a
/* Get SDA Fabric Count


 */
func (s *SdaService) GetSdaFabricCount() (*ResponseSdaGetSdaFabricCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSdaGetSdaFabricCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSdaFabricCount")
	}

	result := response.Result().(*ResponseSdaGetSdaFabricCount)
	return result, response, err

}

//GetPortAssignmentForAccessPointInSdaFabric Get Port assignment for access point in SDA Fabric - 5097-f8d4-45f9-8f51
/* Get Port assignment for access point in SDA Fabric


@param GetPortAssignmentForAccessPointInSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentForAccessPointInSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentForAccessPointInSdaFabric)
	return result, response, err

}

//GetPortAssignmentForUserDeviceInSdaFabric Get Port assignment for user device in SDA Fabric - a4a1-e8ed-41cb-9653
/* Get Port assignment for user device in SDA Fabric.


@param GetPortAssignmentForUserDeviceInSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetPortAssignmentForUserDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetPortAssignmentForUserDeviceInSdaFabric)
	return result, response, err

}

//GetMulticastDetailsFromSdaFabric Get multicast details from SDA fabric - c286-f98b-47ba-9ab4
/* Get multicast details from SDA fabric


@param GetMulticastDetailsFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetMulticastDetailsFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetMulticastDetailsFromSdaFabric)
	return result, response, err

}

//GetProvisionedWiredDevice Get Provisioned Wired Device - dfbf-2ae2-42ca-a449
/* Get Provisioned Wired Device


@param GetProvisionedWiredDeviceQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetProvisionedWiredDevice")
	}

	result := response.Result().(*ResponseSdaGetProvisionedWiredDevice)
	return result, response, err

}

//GetVnFromSdaFabric Get VN from SDA Fabric - 2eb1-fa1e-49ca-a2b4
/* Get virtual network (VN) from SDA Fabric


@param GetVNFromSDAFabricQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetVnFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaGetVnFromSdaFabric)
	return result, response, err

}

//GetIPPoolFromSdaVirtualNetwork Get IP Pool from SDA Virtual Network - fa92-19bf-45c8-b43b
/* Get IP Pool from SDA Virtual Network


@param GetIPPoolFromSDAVirtualNetworkQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetIpPoolFromSdaVirtualNetwork")
	}

	result := response.Result().(*ResponseSdaGetIPPoolFromSdaVirtualNetwork)
	return result, response, err

}

//GetVirtualNetworkWithScalableGroups Get virtual network with scalable groups - ec8a-1ab5-4eba-bca7
/* Get virtual network with scalable groups


@param GetVirtualNetworkWithScalableGroupsQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaGetVirtualNetworkWithScalableGroups)
	return result, response, err

}

//AddDefaultAuthenticationProfileInSdaFabric Add default authentication profile in SDA Fabric - bca3-39d8-44c8-a3c0
/* Add default authentication profile in SDA Fabric


 */
func (s *SdaService) AddDefaultAuthenticationProfileInSdaFabric(requestSdaAddDefaultAuthenticationProfileInSDAFabric *RequestSdaAddDefaultAuthenticationProfileInSdaFabric) (*ResponseSdaAddDefaultAuthenticationProfileInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddDefaultAuthenticationProfileInSDAFabric).
		SetResult(&ResponseSdaAddDefaultAuthenticationProfileInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddDefaultAuthenticationProfileInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddDefaultAuthenticationProfileInSdaFabric)
	return result, response, err

}

//AddsBorderDeviceInSdaFabric Adds border device in SDA Fabric - bead-7b34-43b9-96a7
/* Adds border device in SDA Fabric


 */
func (s *SdaService) AddsBorderDeviceInSdaFabric(requestSdaAddsBorderDeviceInSDAFabric *RequestSdaAddsBorderDeviceInSdaFabric) (*ResponseSdaAddsBorderDeviceInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddsBorderDeviceInSDAFabric).
		SetResult(&ResponseSdaAddsBorderDeviceInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddsBorderDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddsBorderDeviceInSdaFabric)
	return result, response, err

}

//AddControlPlaneDeviceInSdaFabric Add control plane device in SDA Fabric - dd85-c910-4248-9a3f
/* Add control plane device in SDA Fabric


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
		return nil, response, fmt.Errorf("error with operation AddControlPlaneDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddControlPlaneDeviceInSdaFabric)
	return result, response, err

}

//AddEdgeDeviceInSdaFabric Add edge device in SDA Fabric - 87a8-ba44-4ce9-bc59
/* Add edge device in SDA Fabric


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
		return nil, response, fmt.Errorf("error with operation AddEdgeDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddEdgeDeviceInSdaFabric)
	return result, response, err

}

//AddFabric Add Fabric - 6db9-292d-4f28-a26b
/* Add SDA Fabric


 */
func (s *SdaService) AddFabric(requestSdaAddFabric *RequestSdaAddFabric) (*ResponseSdaAddFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddFabric).
		SetResult(&ResponseSdaAddFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddFabric")
	}

	result := response.Result().(*ResponseSdaAddFabric)
	return result, response, err

}

//AddSiteInSdaFabric Add Site in SDA Fabric - d2b4-d9d0-4a4b-884c
/* Add Site in SDA Fabric


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
		return nil, response, fmt.Errorf("error with operation AddSiteInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddSiteInSdaFabric)
	return result, response, err

}

//AddPortAssignmentForAccessPointInSdaFabric Add Port assignment for access point in SDA Fabric - c2a4-3ad2-4098-baa7
/* Add Port assignment for access point in SDA Fabric


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
		return nil, response, fmt.Errorf("error with operation AddPortAssignmentForAccessPointInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentForAccessPointInSdaFabric)
	return result, response, err

}

//AddPortAssignmentForUserDeviceInSdaFabric Add Port assignment for user device in SDA Fabric - 9582-ab82-4ce8-b29d
/* Add Port assignment for user device in SDA Fabric.


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
		return nil, response, fmt.Errorf("error with operation AddPortAssignmentForUserDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddPortAssignmentForUserDeviceInSdaFabric)
	return result, response, err

}

//AddMulticastInSdaFabric Add multicast in SDA fabric - ff85-3826-472a-98fb
/* Add multicast in SDA fabric


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
		return nil, response, fmt.Errorf("error with operation AddMulticastInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddMulticastInSdaFabric)
	return result, response, err

}

//ProvisionWiredDevice Provision Wired Device - cf9a-5843-45fa-9399
/* Provision Wired Device


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
		return nil, response, fmt.Errorf("error with operation ProvisionWiredDevice")
	}

	result := response.Result().(*ResponseSdaProvisionWiredDevice)
	return result, response, err

}

//AddVnInSdaFabric Add VN in SDA Fabric - 518c-59cd-441a-a9fc
/* Add virtual network (VN) in SDA Fabric


 */
func (s *SdaService) AddVnInSdaFabric(requestSdaAddVNInSDAFabric *RequestSdaAddVnInSdaFabric) (*ResponseSdaAddVnInSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/virtual-network"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSdaAddVNInSDAFabric).
		SetResult(&ResponseSdaAddVnInSdaFabric{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddVnInSdaFabric")
	}

	result := response.Result().(*ResponseSdaAddVnInSdaFabric)
	return result, response, err

}

//AddIPPoolInSdaVirtualNetwork Add IP Pool in SDA Virtual Network - 2085-79ea-4ed9-8f4f
/* Add IP Pool in SDA Virtual Network


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
		return nil, response, fmt.Errorf("error with operation AddIpPoolInSdaVirtualNetwork")
	}

	result := response.Result().(*ResponseSdaAddIPPoolInSdaVirtualNetwork)
	return result, response, err

}

//AddVirtualNetworkWithScalableGroups Add virtual network with scalable groups - e3a8-5b19-406a-9f4e
/* Add virtual network with scalable groups at global level


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
		return nil, response, fmt.Errorf("error with operation UpdateDefaultAuthenticationProfileInSdaFabric")
	}

	result := response.Result().(*ResponseSdaUpdateDefaultAuthenticationProfileInSdaFabric)
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
		return nil, response, fmt.Errorf("error with operation UpdateVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaUpdateVirtualNetworkWithScalableGroups)
	return result, response, err

}

//DeleteDefaultAuthenticationProfileFromSdaFabric Delete default authentication profile from SDA Fabric - 3ebc-da3e-4acb-afb7
/* Add default authentication profile in SDA Fabric


@param DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteDefaultAuthenticationProfileFromSdaFabric(DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams *DeleteDefaultAuthenticationProfileFromSdaFabricQueryParams) (*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteDefaultAuthenticationProfileFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteDefaultAuthenticationProfileFromSdaFabric)
	return result, response, err

}

//DeletesBorderDeviceFromSdaFabric Deletes border device from SDA Fabric - cb81-b935-40ba-aab0
/* Deletes border device from SDA Fabric


@param DeletesBorderDeviceFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeletesBorderDeviceFromSdaFabric(DeletesBorderDeviceFromSDAFabricQueryParams *DeletesBorderDeviceFromSdaFabricQueryParams) (*ResponseSdaDeletesBorderDeviceFromSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(DeletesBorderDeviceFromSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeletesBorderDeviceFromSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeletesBorderDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeletesBorderDeviceFromSdaFabric)
	return result, response, err

}

//DeleteControlPlaneDeviceInSdaFabric Delete control plane device in SDA Fabric - f6bd-6bf6-4e68-90be
/* Delete control plane device in SDA Fabric


@param DeleteControlPlaneDeviceInSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteControlPlaneDeviceInSdaFabric(DeleteControlPlaneDeviceInSDAFabricQueryParams *DeleteControlPlaneDeviceInSdaFabricQueryParams) (*ResponseSdaDeleteControlPlaneDeviceInSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteControlPlaneDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteControlPlaneDeviceInSdaFabric)
	return result, response, err

}

//DeleteEdgeDeviceFromSdaFabric Delete edge device from SDA Fabric - 1fb8-f9f2-4c99-8133
/* Delete edge device from SDA Fabric.


@param DeleteEdgeDeviceFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteEdgeDeviceFromSdaFabric(DeleteEdgeDeviceFromSDAFabricQueryParams *DeleteEdgeDeviceFromSdaFabricQueryParams) (*ResponseSdaDeleteEdgeDeviceFromSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteEdgeDeviceFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteEdgeDeviceFromSdaFabric)
	return result, response, err

}

//DeleteSdaFabric Delete SDA Fabric - d0aa-fa69-4f4b-9d7b
/* Delete SDA Fabric


@param DeleteSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteSdaFabric(DeleteSDAFabricQueryParams *DeleteSdaFabricQueryParams) (*ResponseSdaDeleteSdaFabric, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/fabric"

	queryString, _ := query.Values(DeleteSDAFabricQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSdaDeleteSdaFabric{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteSdaFabric)
	return result, response, err

}

//DeleteSiteFromSdaFabric Delete Site from SDA Fabric - 5086-4acf-4ad8-b54d
/* Delete Site from SDA Fabric


@param DeleteSiteFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteSiteFromSdaFabric(DeleteSiteFromSDAFabricQueryParams *DeleteSiteFromSdaFabricQueryParams) (*ResponseSdaDeleteSiteFromSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteSiteFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteSiteFromSdaFabric)
	return result, response, err

}

//DeletePortAssignmentForAccessPointInSdaFabric Delete Port assignment for access point in SDA Fabric - 0787-4a4c-4c9a-abd9
/* Delete Port assignment for access point in SDA Fabric


@param DeletePortAssignmentForAccessPointInSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeletePortAssignmentForAccessPointInSdaFabric(DeletePortAssignmentForAccessPointInSDAFabricQueryParams *DeletePortAssignmentForAccessPointInSdaFabricQueryParams) (*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentForAccessPointInSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentForAccessPointInSdaFabric)
	return result, response, err

}

//DeletePortAssignmentForUserDeviceInSdaFabric Delete Port assignment for user device in SDA Fabric - cba5-b8b1-4edb-81f4
/* Delete Port assignment for user device in SDA Fabric.


@param DeletePortAssignmentForUserDeviceInSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeletePortAssignmentForUserDeviceInSdaFabric(DeletePortAssignmentForUserDeviceInSDAFabricQueryParams *DeletePortAssignmentForUserDeviceInSdaFabricQueryParams) (*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeletePortAssignmentForUserDeviceInSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeletePortAssignmentForUserDeviceInSdaFabric)
	return result, response, err

}

//DeleteMulticastFromSdaFabric Delete multicast from SDA fabric - 2bb0-0be5-45cb-bc99
/* Delete multicast from SDA fabric


@param DeleteMulticastFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteMulticastFromSdaFabric(DeleteMulticastFromSDAFabricQueryParams *DeleteMulticastFromSdaFabricQueryParams) (*ResponseSdaDeleteMulticastFromSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteMulticastFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteMulticastFromSdaFabric)
	return result, response, err

}

//DeleteProvisionedWiredDevice Delete provisioned Wired Device - e495-b94e-463b-ae04
/* Delete provisioned Wired Device


@param DeleteProvisionedWiredDeviceQueryParams Filtering parameter
*/
func (s *SdaService) DeleteProvisionedWiredDevice(DeleteProvisionedWiredDeviceQueryParams *DeleteProvisionedWiredDeviceQueryParams) (*ResponseSdaDeleteProvisionedWiredDevice, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteProvisionedWiredDevice")
	}

	result := response.Result().(*ResponseSdaDeleteProvisionedWiredDevice)
	return result, response, err

}

//DeleteVnFromSdaFabric Delete VN from SDA Fabric - c78c-9ad2-45bb-9657
/* Delete virtual network (VN) from SDA Fabric


@param DeleteVNFromSDAFabricQueryParams Filtering parameter
*/
func (s *SdaService) DeleteVnFromSdaFabric(DeleteVNFromSDAFabricQueryParams *DeleteVnFromSdaFabricQueryParams) (*ResponseSdaDeleteVnFromSdaFabric, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteVnFromSdaFabric")
	}

	result := response.Result().(*ResponseSdaDeleteVnFromSdaFabric)
	return result, response, err

}

//DeleteIPPoolFromSdaVirtualNetwork Delete IP Pool from SDA Virtual Network - 549e-4aff-42bb-b52a
/* Delete IP Pool from SDA Virtual Network


@param DeleteIPPoolFromSDAVirtualNetworkQueryParams Filtering parameter
*/
func (s *SdaService) DeleteIPPoolFromSdaVirtualNetwork(DeleteIPPoolFromSDAVirtualNetworkQueryParams *DeleteIPPoolFromSdaVirtualNetworkQueryParams) (*ResponseSdaDeleteIPPoolFromSdaVirtualNetwork, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteIpPoolFromSdaVirtualNetwork")
	}

	result := response.Result().(*ResponseSdaDeleteIPPoolFromSdaVirtualNetwork)
	return result, response, err

}

//DeleteVirtualNetworkWithScalableGroups Delete virtual network with scalable groups - c8b6-0bc3-4808-8d56
/* Delete virtual network with scalable groups


@param DeleteVirtualNetworkWithScalableGroupsQueryParams Filtering parameter
*/
func (s *SdaService) DeleteVirtualNetworkWithScalableGroups(DeleteVirtualNetworkWithScalableGroupsQueryParams *DeleteVirtualNetworkWithScalableGroupsQueryParams) (*ResponseSdaDeleteVirtualNetworkWithScalableGroups, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteVirtualNetworkWithScalableGroups")
	}

	result := response.Result().(*ResponseSdaDeleteVirtualNetworkWithScalableGroups)
	return result, response, err

}
