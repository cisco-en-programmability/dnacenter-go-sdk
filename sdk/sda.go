package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SDAService is the service to communicate with the SDA API endpoint
type SDAService service

// AddControlPlaneDeviceInSDAFabricRequest is the AddControlPlaneDeviceInSDAFabricRequest definition
type AddControlPlaneDeviceInSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// AddDefaultAuthenticationProfileInSDAFabricRequest is the AddDefaultAuthenticationProfileInSDAFabricRequest definition
type AddDefaultAuthenticationProfileInSDAFabricRequest struct {
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// AddEdgeDeviceInSDAFabricRequest is the AddEdgeDeviceInSDAFabricRequest definition
type AddEdgeDeviceInSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// AddFabricRequest is the AddFabricRequest definition
type AddFabricRequest struct {
	FabricName string `json:"fabricName,omitempty"` //
}

// AddIPPoolInSDAVirtualNetworkRequest is the AddIPPoolInSDAVirtualNetworkRequest definition
type AddIPPoolInSDAVirtualNetworkRequest struct {
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` //
	IPPoolName               string `json:"ipPoolName,omitempty"`               //
	IsL2FloodingEnabled      bool   `json:"isL2FloodingEnabled,omitempty"`      //
	IsThisCriticalPool       bool   `json:"isThisCriticalPool,omitempty"`       //
	PoolType                 string `json:"poolType,omitempty"`                 //
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        //
	TrafficType              string `json:"trafficType,omitempty"`              //
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       //
}

// AddPortAssignmentForAccessPointInSDAFabricRequest is the AddPortAssignmentForAccessPointInSDAFabricRequest definition
type AddPortAssignmentForAccessPointInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// AddPortAssignmentForUserDeviceInSDAFabricRequest is the AddPortAssignmentForUserDeviceInSDAFabricRequest definition
type AddPortAssignmentForUserDeviceInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// AddSiteInSDAFabricRequest is the AddSiteInSDAFabricRequest definition
type AddSiteInSDAFabricRequest struct {
	FabricName        string `json:"fabricName,omitempty"`        //
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` //
}

// AddVNInSDAFabricRequest is the AddVNInSDAFabricRequest definition
type AddVNInSDAFabricRequest struct {
	SiteNameHierarchy  string `json:"siteNameHierarchy,omitempty"`  //
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` //
}

// AddsBorderDeviceInSDAFabricRequest is the AddsBorderDeviceInSDAFabricRequest definition
type AddsBorderDeviceInSDAFabricRequest struct {
	BorderSessionType                 string                         `json:"borderSessionType,omitempty"`                 //
	ConnectedToInternet               bool                           `json:"connectedToInternet,omitempty"`               //
	DeviceManagementIPAddress         string                         `json:"deviceManagementIpAddress,omitempty"`         //
	ExternalConnectivityIPPoolName    string                         `json:"externalConnectivityIpPoolName,omitempty"`    //
	ExternalConnectivitySettings      []ExternalConnectivitySettings `json:"externalConnectivitySettings,omitempty"`      //
	ExternalDomainRoutingProtocolName string                         `json:"externalDomainRoutingProtocolName,omitempty"` //
	InternalAutonomouSystemNumber     string                         `json:"internalAutonomouSystemNumber,omitempty"`     //
	SiteNameHierarchy                 string                         `json:"siteNameHierarchy,omitempty"`                 //
}

// DeleteDefaultAuthenticationProfileFromSDAFabricRequest is the DeleteDefaultAuthenticationProfileFromSDAFabricRequest definition
type DeleteDefaultAuthenticationProfileFromSDAFabricRequest struct {
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// DeleteIPPoolFromSDAVirtualNetworkRequest is the DeleteIPPoolFromSDAVirtualNetworkRequest definition
type DeleteIPPoolFromSDAVirtualNetworkRequest struct {
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` //
	IPPoolName               string `json:"ipPoolName,omitempty"`               //
	IsL2FloodingEnabled      bool   `json:"isL2FloodingEnabled,omitempty"`      //
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
	TrafficType              string `json:"trafficType,omitempty"`              //
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       //
}

// DeletePortAssignmentForAccessPointInSDAFabricRequest is the DeletePortAssignmentForAccessPointInSDAFabricRequest definition
type DeletePortAssignmentForAccessPointInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// DeletePortAssignmentForUserDeviceInSDAFabricRequest is the DeletePortAssignmentForUserDeviceInSDAFabricRequest definition
type DeletePortAssignmentForUserDeviceInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// DeleteSDAFabricRequest is the DeleteSDAFabricRequest definition
type DeleteSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// DeleteSiteFromSDAFabricRequest is the DeleteSiteFromSDAFabricRequest definition
type DeleteSiteFromSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// ExternalConnectivitySettings is the ExternalConnectivitySettings definition
type ExternalConnectivitySettings struct {
	ExternalAutonomouSystemNumber string      `json:"externalAutonomouSystemNumber,omitempty"` //
	InterfaceName                 string      `json:"interfaceName,omitempty"`                 //
	L3Handoff                     []L3Handoff `json:"l3Handoff,omitempty"`                     //
}

// UpdateDefaultAuthenticationProfileInSDAFabricRequest is the UpdateDefaultAuthenticationProfileInSDAFabricRequest definition
type UpdateDefaultAuthenticationProfileInSDAFabricRequest struct {
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// VirtualNetwork is the VirtualNetwork definition
type VirtualNetwork struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` //
}

// AddControlPlaneDeviceInSDAFabricResponse is the AddControlPlaneDeviceInSDAFabricResponse definition
type AddControlPlaneDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddDefaultAuthenticationProfileInSDAFabricResponse is the AddDefaultAuthenticationProfileInSDAFabricResponse definition
type AddDefaultAuthenticationProfileInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddEdgeDeviceInSDAFabricResponse is the AddEdgeDeviceInSDAFabricResponse definition
type AddEdgeDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddFabricResponse is the AddFabricResponse definition
type AddFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddIPPoolInSDAVirtualNetworkResponse is the AddIPPoolInSDAVirtualNetworkResponse definition
type AddIPPoolInSDAVirtualNetworkResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddPortAssignmentForAccessPointInSDAFabricResponse is the AddPortAssignmentForAccessPointInSDAFabricResponse definition
type AddPortAssignmentForAccessPointInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddPortAssignmentForUserDeviceInSDAFabricResponse is the AddPortAssignmentForUserDeviceInSDAFabricResponse definition
type AddPortAssignmentForUserDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddSiteInSDAFabricResponse is the AddSiteInSDAFabricResponse definition
type AddSiteInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddVNInSDAFabricResponse is the AddVNInSDAFabricResponse definition
type AddVNInSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// AddsBorderDeviceInSDAFabricResponse is the AddsBorderDeviceInSDAFabricResponse definition
type AddsBorderDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteControlPlaneDeviceInSDAFabricResponse is the DeleteControlPlaneDeviceInSDAFabricResponse definition
type DeleteControlPlaneDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteDefaultAuthenticationProfileFromSDAFabricResponse is the DeleteDefaultAuthenticationProfileFromSDAFabricResponse definition
type DeleteDefaultAuthenticationProfileFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteEdgeDeviceFromSDAFabricResponse is the DeleteEdgeDeviceFromSDAFabricResponse definition
type DeleteEdgeDeviceFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteIPPoolFromSDAVirtualNetworkResponse is the DeleteIPPoolFromSDAVirtualNetworkResponse definition
type DeleteIPPoolFromSDAVirtualNetworkResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeletePortAssignmentForAccessPointInSDAFabricResponse is the DeletePortAssignmentForAccessPointInSDAFabricResponse definition
type DeletePortAssignmentForAccessPointInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeletePortAssignmentForUserDeviceInSDAFabricResponse is the DeletePortAssignmentForUserDeviceInSDAFabricResponse definition
type DeletePortAssignmentForUserDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteSDAFabricResponse is the DeleteSDAFabricResponse definition
type DeleteSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteSiteFromSDAFabricResponse is the DeleteSiteFromSDAFabricResponse definition
type DeleteSiteFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteVNFromSDAFabricResponse is the DeleteVNFromSDAFabricResponse definition
type DeleteVNFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// DeletesBorderDeviceFromSDAFabricResponse is the DeletesBorderDeviceFromSDAFabricResponse definition
type DeletesBorderDeviceFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeviceSettings is the DeviceSettings definition
type DeviceSettings struct {
	ConnectedTo                   []string                  `json:"connectedTo,omitempty"`                   //
	CPU                           int                       `json:"cpu,omitempty"`                           //
	DeployPending                 string                    `json:"deployPending,omitempty"`                 //
	DhcpEnabled                   bool                      `json:"dhcpEnabled,omitempty"`                   //
	DisplayName                   string                    `json:"displayName,omitempty"`                   //
	ExtConnectivitySettings       []ExtConnectivitySettings `json:"extConnectivitySettings,omitempty"`       //
	ExternalConnectivityIPPool    string                    `json:"externalConnectivityIpPool,omitempty"`    //
	ExternalDomainRoutingProtocol string                    `json:"externalDomainRoutingProtocol,omitempty"` //
	ID                            string                    `json:"id,omitempty"`                            //
	InstanceID                    int                       `json:"instanceId,omitempty"`                    //
	InstanceTenantID              string                    `json:"instanceTenantId,omitempty"`              //
	InstanceVersion               int                       `json:"instanceVersion,omitempty"`               //
	InternalDomainProtocolNumber  string                    `json:"internalDomainProtocolNumber,omitempty"`  //
	Memory                        int                       `json:"memory,omitempty"`                        //
	NodeType                      []string                  `json:"nodeType,omitempty"`                      //
	Storage                       int                       `json:"storage,omitempty"`                       //
}

// Dhcp is the Dhcp definition
type Dhcp struct {
	ID        string    `json:"id,omitempty"`        //
	IPAddress IPAddress `json:"ipAddress,omitempty"` //
}

// DNS is the Dns definition
type DNS struct {
	DomainName string `json:"domainName,omitempty"` //
	ID         string `json:"id,omitempty"`         //
	IP         IP     `json:"ip,omitempty"`         //
}

// ExtConnectivitySettings is the ExtConnectivitySettings definition
type ExtConnectivitySettings struct {
	DeployPending                string      `json:"deployPending,omitempty"`                //
	DisplayName                  string      `json:"displayName,omitempty"`                  //
	ExternalDomainProtocolNumber string      `json:"externalDomainProtocolNumber,omitempty"` //
	ID                           string      `json:"id,omitempty"`                           //
	InstanceID                   int         `json:"instanceId,omitempty"`                   //
	InstanceTenantID             string      `json:"instanceTenantId,omitempty"`             //
	InstanceVersion              int         `json:"instanceVersion,omitempty"`              //
	InterfaceUUID                string      `json:"interfaceUuid,omitempty"`                //
	L2Handoff                    []string    `json:"l2Handoff,omitempty"`                    //
	L3Handoff                    []L3Handoff `json:"l3Handoff,omitempty"`                    //
	PolicyPropagationEnabled     bool        `json:"policyPropagationEnabled,omitempty"`     //
	PolicySgtTag                 int         `json:"policySgtTag,omitempty"`                 //
}

// GetControlPlaneDeviceFromSDAFabricResponse is the GetControlPlaneDeviceFromSDAFabricResponse definition
type GetControlPlaneDeviceFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetDefaultAuthenticationProfileFromSDAFabricResponse is the GetDefaultAuthenticationProfileFromSDAFabricResponse definition
type GetDefaultAuthenticationProfileFromSDAFabricResponse struct {
	AuthenticateTemplateID   string `json:"authenticateTemplateId,omitempty"`   //
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// GetDeviceInfoFromSDAFabricResponse is the GetDeviceInfoFromSDAFabricResponse definition
type GetDeviceInfoFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetDeviceRoleInSDAFabricResponse is the GetDeviceRoleInSDAFabricResponse definition
type GetDeviceRoleInSDAFabricResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetEdgeDeviceFromSDAFabricResponse is the GetEdgeDeviceFromSDAFabricResponse definition
type GetEdgeDeviceFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetIPPoolFromSDAVirtualNetworkResponse is the GetIPPoolFromSDAVirtualNetworkResponse definition
type GetIPPoolFromSDAVirtualNetworkResponse struct {
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` //
	Description              string `json:"description,omitempty"`              //
	IPPoolName               string `json:"ipPoolName,omitempty"`               //
	IsL2FloodingEnabled      bool   `json:"isL2FloodingEnabled,omitempty"`      //
	IsThisCriticalPool       bool   `json:"isThisCriticalPool,omitempty"`       //
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        //
	Status                   string `json:"status,omitempty"`                   //
	TrafficType              string `json:"trafficType,omitempty"`              //
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       //
}

// GetPortAssignmentForAccessPointInSDAFabricResponse is the GetPortAssignmentForAccessPointInSDAFabricResponse definition
type GetPortAssignmentForAccessPointInSDAFabricResponse struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	Description               string `json:"description,omitempty"`               //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	Status                    string `json:"status,omitempty"`                    //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// GetPortAssignmentForUserDeviceInSDAFabricResponse is the GetPortAssignmentForUserDeviceInSDAFabricResponse definition
type GetPortAssignmentForUserDeviceInSDAFabricResponse struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	Description               string `json:"description,omitempty"`               //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	Status                    string `json:"status,omitempty"`                    //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// GetSDAFabricCountResponse is the GetSDAFabricCountResponse definition
type GetSDAFabricCountResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetSDAFabricInfoResponse is the GetSDAFabricInfoResponse definition
type GetSDAFabricInfoResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// GetSiteFromSDAFabricResponse is the GetSiteFromSDAFabricResponse definition
type GetSiteFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// GetVNFromSDAFabricResponse is the GetVNFromSDAFabricResponse definition
type GetVNFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetsBorderDeviceDetailFromSDAFabricResponse is the GetsBorderDeviceDetailFromSDAFabricResponse definition
type GetsBorderDeviceDetailFromSDAFabricResponse struct {
	Description string  `json:"description,omitempty"` //
	Payload     Payload `json:"payload,omitempty"`     //
	Status      string  `json:"status,omitempty"`      //
}

// IP is the Ip definition
type IP struct {
	Address       string `json:"address,omitempty"`       //
	AddressType   string `json:"addressType,omitempty"`   //
	ID            string `json:"id,omitempty"`            //
	PaddedAddress string `json:"paddedAddress,omitempty"` //
}

// IPAddress is the IpAddress definition
type IPAddress struct {
	Address       string `json:"address,omitempty"`       //
	AddressType   string `json:"addressType,omitempty"`   //
	ID            string `json:"id,omitempty"`            //
	PaddedAddress string `json:"paddedAddress,omitempty"` //
}

// L3Handoff is the L3Handoff definition
type L3Handoff struct {
	DeployPending    string         `json:"deployPending,omitempty"`    //
	DisplayName      string         `json:"displayName,omitempty"`      //
	ID               string         `json:"id,omitempty"`               //
	InstanceID       int            `json:"instanceId,omitempty"`       //
	InstanceTenantID string         `json:"instanceTenantId,omitempty"` //
	InstanceVersion  int            `json:"instanceVersion,omitempty"`  //
	LocalIPAddress   string         `json:"localIpAddress,omitempty"`   //
	RemoteIPAddress  string         `json:"remoteIpAddress,omitempty"`  //
	VirtualNetwork   VirtualNetwork `json:"virtualNetwork,omitempty"`   //
	VlanID           int            `json:"vlanId,omitempty"`           //
}

// NetworkWideSettings is the NetworkWideSettings definition
type NetworkWideSettings struct {
	Aaa              []string `json:"aaa,omitempty"`              //
	Cmx              []string `json:"cmx,omitempty"`              //
	DeployPending    string   `json:"deployPending,omitempty"`    //
	Dhcp             []DHCP   `json:"dhcp,omitempty"`             //
	DisplayName      string   `json:"displayName,omitempty"`      //
	DNS              []DNS    `json:"dns,omitempty"`              //
	ID               string   `json:"id,omitempty"`               //
	InstanceID       int      `json:"instanceId,omitempty"`       //
	InstanceTenantID string   `json:"instanceTenantId,omitempty"` //
	InstanceVersion  int      `json:"instanceVersion,omitempty"`  //
	Ldap             []string `json:"ldap,omitempty"`             //
	NativeVlan       []string `json:"nativeVlan,omitempty"`       //
	Netflow          []string `json:"netflow,omitempty"`          //
	Ntp              []string `json:"ntp,omitempty"`              //
	SNMP             []string `json:"snmp,omitempty"`             //
	Syslogs          []string `json:"syslogs,omitempty"`          //
}

// Payload is the Payload definition
type Payload struct {
	AkcSettingsCfs                 []string            `json:"akcSettingsCfs,omitempty"`                 //
	AuthEntityClass                int                 `json:"authEntityClass,omitempty"`                //
	AuthEntityID                   int                 `json:"authEntityId,omitempty"`                   //
	CfsChangeInfo                  []string            `json:"cfsChangeInfo,omitempty"`                  //
	Configs                        []string            `json:"configs,omitempty"`                        //
	CreateTime                     int                 `json:"createTime,omitempty"`                     //
	CustomProvisions               []string            `json:"customProvisions,omitempty"`               //
	DeployPending                  string              `json:"deployPending,omitempty"`                  //
	Deployed                       bool                `json:"deployed,omitempty"`                       //
	DeviceInterfaceInfo            []string            `json:"deviceInterfaceInfo,omitempty"`            //
	DeviceSettings                 DeviceSettings      `json:"deviceSettings,omitempty"`                 //
	DisplayName                    string              `json:"displayName,omitempty"`                    //
	ID                             string              `json:"id,omitempty"`                             //
	InstanceID                     int                 `json:"instanceId,omitempty"`                     //
	InstanceTenantID               string              `json:"instanceTenantId,omitempty"`               //
	InstanceVersion                int                 `json:"instanceVersion,omitempty"`                //
	IsSeeded                       bool                `json:"isSeeded,omitempty"`                       //
	IsStale                        bool                `json:"isStale,omitempty"`                        //
	LastUpdateTime                 int                 `json:"lastUpdateTime,omitempty"`                 //
	ManagedSites                   []string            `json:"managedSites,omitempty"`                   //
	Name                           string              `json:"name,omitempty"`                           //
	Namespace                      string              `json:"namespace,omitempty"`                      //
	NetworkDeviceID                string              `json:"networkDeviceId,omitempty"`                //
	NetworkWideSettings            NetworkWideSettings `json:"networkWideSettings,omitempty"`            //
	OtherDevice                    []string            `json:"otherDevice,omitempty"`                    //
	ProvisioningState              string              `json:"provisioningState,omitempty"`              //
	ResourceVersion                int                 `json:"resourceVersion,omitempty"`                //
	Roles                          []string            `json:"roles,omitempty"`                          //
	SaveWanConnectivityDetailsOnly bool                `json:"saveWanConnectivityDetailsOnly,omitempty"` //
	SiteID                         string              `json:"siteId,omitempty"`                         //
	TargetIDList                   []string            `json:"targetIdList,omitempty"`                   //
	TransitNetworks                []TransitNetworks   `json:"transitNetworks,omitempty"`                //
	Type                           string              `json:"type,omitempty"`                           //
	VirtualNetwork                 []string            `json:"virtualNetwork,omitempty"`                 //
	Wlan                           []string            `json:"wlan,omitempty"`                           //
}

// Response is the Response definition
// type Response struct {
// 	Description string `json:"description,omitempty"` //
// 	FabricCount string `json:"fabricCount,omitempty"` //
// 	Status      string `json:"status,omitempty"`      //
// }

// TransitNetworks is the TransitNetworks definition
type TransitNetworks struct {
	IDRef string `json:"idRef,omitempty"` //
}

// UpdateDefaultAuthenticationProfileInSDAFabricResponse is the UpdateDefaultAuthenticationProfileInSDAFabricResponse definition
type UpdateDefaultAuthenticationProfileInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// VirtualNetwork is the VirtualNetwork definition
// type VirtualNetwork struct {
// 	IDRef string `json:"idRef,omitempty"` //
// }

// AddControlPlaneDeviceInSDAFabric addControlPlaneDeviceInSDAFabric
/* Add control plane device in SDA Fabric
 */
func (s *SDAService) AddControlPlaneDeviceInSDAFabric(addControlPlaneDeviceInSDAFabricRequest *AddControlPlaneDeviceInSDAFabricRequest) (*AddControlPlaneDeviceInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	response, err := RestyClient.R().
		SetBody(addControlPlaneDeviceInSDAFabricRequest).
		SetResult(&AddControlPlaneDeviceInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddControlPlaneDeviceInSDAFabricResponse)
	return result, response, err

}

// AddDefaultAuthenticationProfileInSDAFabric addDefaultAuthenticationProfileInSDAFabric
/* Add default authentication profile in SDA Fabric
 */
func (s *SDAService) AddDefaultAuthenticationProfileInSDAFabric(addDefaultAuthenticationProfileInSDAFabricRequest *AddDefaultAuthenticationProfileInSDAFabricRequest) (*AddDefaultAuthenticationProfileInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := RestyClient.R().
		SetBody(addDefaultAuthenticationProfileInSDAFabricRequest).
		SetResult(&AddDefaultAuthenticationProfileInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddDefaultAuthenticationProfileInSDAFabricResponse)
	return result, response, err

}

// AddEdgeDeviceInSDAFabric addEdgeDeviceInSDAFabric
/* Add edge device in SDA Fabric
 */
func (s *SDAService) AddEdgeDeviceInSDAFabric(addEdgeDeviceInSDAFabricRequest *AddEdgeDeviceInSDAFabricRequest) (*AddEdgeDeviceInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/edge-device"

	response, err := RestyClient.R().
		SetBody(addEdgeDeviceInSDAFabricRequest).
		SetResult(&AddEdgeDeviceInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddEdgeDeviceInSDAFabricResponse)
	return result, response, err

}

// AddFabric addFabric
/* Add SDA Fabric
 */
func (s *SDAService) AddFabric(addFabricRequest *AddFabricRequest) (*AddFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric"

	response, err := RestyClient.R().
		SetBody(addFabricRequest).
		SetResult(&AddFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddFabricResponse)
	return result, response, err

}

// AddIPPoolInSDAVirtualNetwork addIPPoolInSDAVirtualNetwork
/* Add IP Pool in SDA Virtual Network
 */
func (s *SDAService) AddIPPoolInSDAVirtualNetwork(addIPPoolInSDAVirtualNetworkRequest *AddIPPoolInSDAVirtualNetworkRequest) (*AddIPPoolInSDAVirtualNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	response, err := RestyClient.R().
		SetBody(addIPPoolInSDAVirtualNetworkRequest).
		SetResult(&AddIPPoolInSDAVirtualNetworkResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddIPPoolInSDAVirtualNetworkResponse)
	return result, response, err

}

// AddPortAssignmentForAccessPointInSDAFabric addPortAssignmentForAccessPointInSDAFabric
/* Add Port assignment for access point in SDA Fabric
 */
func (s *SDAService) AddPortAssignmentForAccessPointInSDAFabric(addPortAssignmentForAccessPointInSDAFabricRequest *AddPortAssignmentForAccessPointInSDAFabricRequest) (*AddPortAssignmentForAccessPointInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	response, err := RestyClient.R().
		SetBody(addPortAssignmentForAccessPointInSDAFabricRequest).
		SetResult(&AddPortAssignmentForAccessPointInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddPortAssignmentForAccessPointInSDAFabricResponse)
	return result, response, err

}

// AddPortAssignmentForUserDeviceInSDAFabric addPortAssignmentForUserDeviceInSDAFabric
/* Add Port assignment for user device in SDA Fabric.
 */
func (s *SDAService) AddPortAssignmentForUserDeviceInSDAFabric(addPortAssignmentForUserDeviceInSDAFabricRequest *AddPortAssignmentForUserDeviceInSDAFabricRequest) (*AddPortAssignmentForUserDeviceInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	response, err := RestyClient.R().
		SetBody(addPortAssignmentForUserDeviceInSDAFabricRequest).
		SetResult(&AddPortAssignmentForUserDeviceInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddPortAssignmentForUserDeviceInSDAFabricResponse)
	return result, response, err

}

// AddSiteInSDAFabric addSiteInSDAFabric
/* Add Site in SDA Fabric
 */
func (s *SDAService) AddSiteInSDAFabric(addSiteInSDAFabricRequest *AddSiteInSDAFabricRequest) (*AddSiteInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric-site"

	response, err := RestyClient.R().
		SetBody(addSiteInSDAFabricRequest).
		SetResult(&AddSiteInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddSiteInSDAFabricResponse)
	return result, response, err

}

// AddVNInSDAFabric addVNInSDAFabric
/* Add virtual network (VN) in SDA Fabric
 */
func (s *SDAService) AddVNInSDAFabric(addVNInSDAFabricRequest *AddVNInSDAFabricRequest) (*AddVNInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/virtual-network"

	response, err := RestyClient.R().
		SetBody(addVNInSDAFabricRequest).
		SetResult(&AddVNInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddVNInSDAFabricResponse)
	return result, response, err

}

// AddsBorderDeviceInSDAFabric addsBorderDeviceInSDAFabric
/* Adds border device in SDA Fabric
 */
func (s *SDAService) AddsBorderDeviceInSDAFabric(addsBorderDeviceInSDAFabricRequest *AddsBorderDeviceInSDAFabricRequest) (*AddsBorderDeviceInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/border-device"

	response, err := RestyClient.R().
		SetBody(addsBorderDeviceInSDAFabricRequest).
		SetResult(&AddsBorderDeviceInSDAFabricResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AddsBorderDeviceInSDAFabricResponse)
	return result, response, err

}

// DeleteControlPlaneDeviceInSDAFabricQueryParams defines the query parameters for this request
type DeleteControlPlaneDeviceInSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// DeleteControlPlaneDeviceInSDAFabric deleteControlPlaneDeviceInSDAFabric
/* Delete control plane device in SDA Fabric
@param deviceIPAddress Device IP Address
*/
func (s *SDAService) DeleteControlPlaneDeviceInSDAFabric(deleteControlPlaneDeviceInSDAFabricQueryParams *DeleteControlPlaneDeviceInSDAFabricQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	queryString, _ := query.Values(deleteControlPlaneDeviceInSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams defines the query parameters for this request
type DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` // siteNameHierarchy
}

// DeleteDefaultAuthenticationProfileFromSDAFabric deleteDefaultAuthenticationProfileFromSDAFabric
/* Add default authentication profile in SDA Fabric
@param siteNameHierarchy siteNameHierarchy
*/
func (s *SDAService) DeleteDefaultAuthenticationProfileFromSDAFabric(deleteDefaultAuthenticationProfileFromSDAFabricQueryParams *DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams, deleteDefaultAuthenticationProfileFromSDAFabricRequest *DeleteDefaultAuthenticationProfileFromSDAFabricRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	queryString, _ := query.Values(deleteDefaultAuthenticationProfileFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deleteDefaultAuthenticationProfileFromSDAFabricRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteEdgeDeviceFromSDAFabricQueryParams defines the query parameters for this request
type DeleteEdgeDeviceFromSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// DeleteEdgeDeviceFromSDAFabric deleteEdgeDeviceFromSDAFabric
/* Delete edge device from SDA Fabric.
@param deviceIPAddress Device IP Address
*/
func (s *SDAService) DeleteEdgeDeviceFromSDAFabric(deleteEdgeDeviceFromSDAFabricQueryParams *DeleteEdgeDeviceFromSDAFabricQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/edge-device"

	queryString, _ := query.Values(deleteEdgeDeviceFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteIPPoolFromSDAVirtualNetworkQueryParams defines the query parameters for this request
type DeleteIPPoolFromSDAVirtualNetworkQueryParams struct {
	IPPoolName         string `url:"ipPoolName,omitempty"`         // ipPoolName
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` // virtualNetworkName
}

// DeleteIPPoolFromSDAVirtualNetwork deleteIPPoolFromSDAVirtualNetwork
/* Delete IP Pool from SDA Virtual Network
@param ipPoolName ipPoolName
@param virtualNetworkName virtualNetworkName
*/
func (s *SDAService) DeleteIPPoolFromSDAVirtualNetwork(deleteIPPoolFromSDAVirtualNetworkQueryParams *DeleteIPPoolFromSDAVirtualNetworkQueryParams, deleteIPPoolFromSDAVirtualNetworkRequest *DeleteIPPoolFromSDAVirtualNetworkRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	queryString, _ := query.Values(deleteIPPoolFromSDAVirtualNetworkQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deleteIPPoolFromSDAVirtualNetworkRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeletePortAssignmentForAccessPointInSDAFabricQueryParams defines the query parameters for this request
type DeletePortAssignmentForAccessPointInSDAFabricQueryParams struct {
	Devicp        string `url:"device-ip,omitempty"`     // device-ip
	InterfaceName string `url:"interfaceName,omitempty"` // interfaceName
}

// DeletePortAssignmentForAccessPointInSDAFabric deletePortAssignmentForAccessPointInSDAFabric
/* Delete Port assignment for access point in SDA Fabric
@param device-ip device-ip
@param interfaceName interfaceName
*/
func (s *SDAService) DeletePortAssignmentForAccessPointInSDAFabric(deletePortAssignmentForAccessPointInSDAFabricQueryParams *DeletePortAssignmentForAccessPointInSDAFabricQueryParams, deletePortAssignmentForAccessPointInSDAFabricRequest *DeletePortAssignmentForAccessPointInSDAFabricRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	queryString, _ := query.Values(deletePortAssignmentForAccessPointInSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deletePortAssignmentForAccessPointInSDAFabricRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeletePortAssignmentForUserDeviceInSDAFabricQueryParams defines the query parameters for this request
type DeletePortAssignmentForUserDeviceInSDAFabricQueryParams struct {
	Devicp        string `url:"device-ip,omitempty"`     // device-ip
	InterfaceName string `url:"interfaceName,omitempty"` // interfaceName
}

// DeletePortAssignmentForUserDeviceInSDAFabric deletePortAssignmentForUserDeviceInSDAFabric
/* Delete Port assignment for user device in SDA Fabric.
@param device-ip device-ip
@param interfaceName interfaceName
*/
func (s *SDAService) DeletePortAssignmentForUserDeviceInSDAFabric(deletePortAssignmentForUserDeviceInSDAFabricQueryParams *DeletePortAssignmentForUserDeviceInSDAFabricQueryParams, deletePortAssignmentForUserDeviceInSDAFabricRequest *DeletePortAssignmentForUserDeviceInSDAFabricRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	queryString, _ := query.Values(deletePortAssignmentForUserDeviceInSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deletePortAssignmentForUserDeviceInSDAFabricRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteSDAFabricQueryParams defines the query parameters for this request
type DeleteSDAFabricQueryParams struct {
	FabricName string `url:"fabricName,omitempty"` // Fabric Name
}

// DeleteSDAFabric deleteSDAFabric
/* Delete SDA Fabric
@param fabricName Fabric Name
*/
func (s *SDAService) DeleteSDAFabric(deleteSDAFabricQueryParams *DeleteSDAFabricQueryParams, deleteSDAFabricRequest *DeleteSDAFabricRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric"

	queryString, _ := query.Values(deleteSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deleteSDAFabricRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteSiteFromSDAFabricQueryParams defines the query parameters for this request
type DeleteSiteFromSDAFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}

// DeleteSiteFromSDAFabric deleteSiteFromSDAFabric
/* Delete Site from SDA Fabric
@param siteNameHierarchy Site Name Hierarchy
*/
func (s *SDAService) DeleteSiteFromSDAFabric(deleteSiteFromSDAFabricQueryParams *DeleteSiteFromSDAFabricQueryParams, deleteSiteFromSDAFabricRequest *DeleteSiteFromSDAFabricRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric-site"

	queryString, _ := query.Values(deleteSiteFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deleteSiteFromSDAFabricRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteVNFromSDAFabricQueryParams defines the query parameters for this request
type DeleteVNFromSDAFabricQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` // virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  // siteNameHierarchy
}

// DeleteVNFromSDAFabric deleteVNFromSDAFabric
/* Delete virtual network (VN) from SDA Fabric
@param virtualNetworkName virtualNetworkName
@param siteNameHierarchy siteNameHierarchy
*/
func (s *SDAService) DeleteVNFromSDAFabric(deleteVNFromSDAFabricQueryParams *DeleteVNFromSDAFabricQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/virtual-network"

	queryString, _ := query.Values(deleteVNFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeletesBorderDeviceFromSDAFabricQueryParams defines the query parameters for this request
type DeletesBorderDeviceFromSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// DeletesBorderDeviceFromSDAFabric deletesBorderDeviceFromSDAFabric
/* Deletes border device from SDA Fabric
@param deviceIPAddress Device IP Address
*/
// func (s *SDAService) DeletesBorderDeviceFromSDAFabric(deletesBorderDeviceFromSDAFabricQueryParams *DeletesBorderDeviceFromSDAFabricQueryParams, deletesBorderDeviceFromSDAFabricRequest *DeletesBorderDeviceFromSDAFabricRequest) (*resty.Response, error) {

// 	path := "/dna/intent/api/v1/business/sda/border-device"

// 	queryString, _ := query.Values(deletesBorderDeviceFromSDAFabricQueryParams)

// 	response, err := RestyClient.R().
// 		SetQueryString(queryString.Encode()).
// 		SetBody(deletesBorderDeviceFromSDAFabricRequest).
// 		SetError(&Error{}).
// 		Delete(path)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return response, err

// }

// GetControlPlaneDeviceFromSDAFabricQueryParams defines the query parameters for this request
type GetControlPlaneDeviceFromSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// GetControlPlaneDeviceFromSDAFabric getControlPlaneDeviceFromSDAFabric
/* Get control plane device from SDA Fabric
@param deviceIPAddress Device IP Address
*/
func (s *SDAService) GetControlPlaneDeviceFromSDAFabric(getControlPlaneDeviceFromSDAFabricQueryParams *GetControlPlaneDeviceFromSDAFabricQueryParams) (*GetControlPlaneDeviceFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/control-plane-device"

	queryString, _ := query.Values(getControlPlaneDeviceFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetControlPlaneDeviceFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetControlPlaneDeviceFromSDAFabricResponse)
	return result, response, err

}

// GetDefaultAuthenticationProfileFromSDAFabricQueryParams defines the query parameters for this request
type GetDefaultAuthenticationProfileFromSDAFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` // siteNameHierarchy
}

// GetDefaultAuthenticationProfileFromSDAFabric getDefaultAuthenticationProfileFromSDAFabric
/* Get default authentication profile from SDA Fabric
@param siteNameHierarchy siteNameHierarchy
*/
func (s *SDAService) GetDefaultAuthenticationProfileFromSDAFabric(getDefaultAuthenticationProfileFromSDAFabricQueryParams *GetDefaultAuthenticationProfileFromSDAFabricQueryParams) (*GetDefaultAuthenticationProfileFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	queryString, _ := query.Values(getDefaultAuthenticationProfileFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDefaultAuthenticationProfileFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDefaultAuthenticationProfileFromSDAFabricResponse)
	return result, response, err

}

// GetDeviceInfoFromSDAFabricQueryParams defines the query parameters for this request
type GetDeviceInfoFromSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// GetDeviceInfoFromSDAFabric getDeviceInfoFromSDAFabric
/* Get device info from SDA Fabric
@param deviceIPAddress Device IP Address
*/
func (s *SDAService) GetDeviceInfoFromSDAFabric(getDeviceInfoFromSDAFabricQueryParams *GetDeviceInfoFromSDAFabricQueryParams) (*GetDeviceInfoFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/device"

	queryString, _ := query.Values(getDeviceInfoFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceInfoFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDeviceInfoFromSDAFabricResponse)
	return result, response, err

}

// GetDeviceRoleInSDAFabricQueryParams defines the query parameters for this request
type GetDeviceRoleInSDAFabricQueryParams struct {
	DeviceManagementIPAddress string `url:"deviceManagementIpAddress,omitempty"` // Device Management IP Address
}

// GetDeviceRoleInSDAFabric getDeviceRoleInSDAFabric
/* Get device role in SDA Fabric
@param deviceManagementIPAddress Device Management IP Address
*/
func (s *SDAService) GetDeviceRoleInSDAFabric(getDeviceRoleInSDAFabricQueryParams *GetDeviceRoleInSDAFabricQueryParams) (*GetDeviceRoleInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/device/role"

	queryString, _ := query.Values(getDeviceRoleInSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceRoleInSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDeviceRoleInSDAFabricResponse)
	return result, response, err

}

// GetEdgeDeviceFromSDAFabricQueryParams defines the query parameters for this request
type GetEdgeDeviceFromSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// GetEdgeDeviceFromSDAFabric getEdgeDeviceFromSDAFabric
/* Get edge device from SDA Fabric
@param deviceIPAddress Device IP Address
*/
func (s *SDAService) GetEdgeDeviceFromSDAFabric(getEdgeDeviceFromSDAFabricQueryParams *GetEdgeDeviceFromSDAFabricQueryParams) (*GetEdgeDeviceFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/edge-device"

	queryString, _ := query.Values(getEdgeDeviceFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetEdgeDeviceFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetEdgeDeviceFromSDAFabricResponse)
	return result, response, err

}

// GetIPPoolFromSDAVirtualNetworkQueryParams defines the query parameters for this request
type GetIPPoolFromSDAVirtualNetworkQueryParams struct {
	IPPoolName         string `url:"ipPoolName,omitempty"`         // ipPoolName
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` // virtualNetworkName
}

// GetIPPoolFromSDAVirtualNetwork getIPPoolFromSDAVirtualNetwork
/* Get IP Pool from SDA Virtual Network
@param ipPoolName ipPoolName
@param virtualNetworkName virtualNetworkName
*/
func (s *SDAService) GetIPPoolFromSDAVirtualNetwork(getIPPoolFromSDAVirtualNetworkQueryParams *GetIPPoolFromSDAVirtualNetworkQueryParams) (*GetIPPoolFromSDAVirtualNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/virtualnetwork/ippool"

	queryString, _ := query.Values(getIPPoolFromSDAVirtualNetworkQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetIPPoolFromSDAVirtualNetworkResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetIPPoolFromSDAVirtualNetworkResponse)
	return result, response, err

}

// GetPortAssignmentForAccessPointInSDAFabricQueryParams defines the query parameters for this request
type GetPortAssignmentForAccessPointInSDAFabricQueryParams struct {
	Devicp        string `url:"device-ip,omitempty"`     // device-ip
	InterfaceName string `url:"interfaceName,omitempty"` // interfaceName
}

// GetPortAssignmentForAccessPointInSDAFabric getPortAssignmentForAccessPointInSDAFabric
/* Get Port assignment for access point in SDA Fabric
@param device-ip device-ip
@param interfaceName interfaceName
*/
func (s *SDAService) GetPortAssignmentForAccessPointInSDAFabric(getPortAssignmentForAccessPointInSDAFabricQueryParams *GetPortAssignmentForAccessPointInSDAFabricQueryParams) (*GetPortAssignmentForAccessPointInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/hostonboarding/access-point"

	queryString, _ := query.Values(getPortAssignmentForAccessPointInSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetPortAssignmentForAccessPointInSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetPortAssignmentForAccessPointInSDAFabricResponse)
	return result, response, err

}

// GetPortAssignmentForUserDeviceInSDAFabricQueryParams defines the query parameters for this request
type GetPortAssignmentForUserDeviceInSDAFabricQueryParams struct {
	Devicp        string `url:"device-ip,omitempty"`     // device-ip
	InterfaceName string `url:"interfaceName,omitempty"` // interfaceName
}

// GetPortAssignmentForUserDeviceInSDAFabric getPortAssignmentForUserDeviceInSDAFabric
/* Get Port assignment for user device in SDA Fabric.
@param device-ip device-ip
@param interfaceName interfaceName
*/
func (s *SDAService) GetPortAssignmentForUserDeviceInSDAFabric(getPortAssignmentForUserDeviceInSDAFabricQueryParams *GetPortAssignmentForUserDeviceInSDAFabricQueryParams) (*GetPortAssignmentForUserDeviceInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/hostonboarding/user-device"

	queryString, _ := query.Values(getPortAssignmentForUserDeviceInSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetPortAssignmentForUserDeviceInSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetPortAssignmentForUserDeviceInSDAFabricResponse)
	return result, response, err

}

// GetSDAFabricCount getSDAFabricCount
/* Get SDA Fabric Count
 */
func (s *SDAService) GetSDAFabricCount() (*GetSDAFabricCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric/count"

	response, err := RestyClient.R().
		SetResult(&GetSDAFabricCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSDAFabricCountResponse)
	return result, response, err

}

// GetSDAFabricInfoQueryParams defines the query parameters for this request
type GetSDAFabricInfoQueryParams struct {
	FabricName string `url:"fabricName,omitempty"` // Fabric Name
}

// GetSDAFabricInfo getSDAFabricInfo
/* Get SDA Fabric Info
@param fabricName Fabric Name
*/
func (s *SDAService) GetSDAFabricInfo(getSDAFabricInfoQueryParams *GetSDAFabricInfoQueryParams) (*GetSDAFabricInfoResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric"

	queryString, _ := query.Values(getSDAFabricInfoQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetSDAFabricInfoResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSDAFabricInfoResponse)
	return result, response, err

}

// GetSiteFromSDAFabricQueryParams defines the query parameters for this request
type GetSiteFromSDAFabricQueryParams struct {
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}

// GetSiteFromSDAFabric getSiteFromSDAFabric
/* Get Site info from SDA Fabric
@param siteNameHierarchy Site Name Hierarchy
*/
func (s *SDAService) GetSiteFromSDAFabric(getSiteFromSDAFabricQueryParams *GetSiteFromSDAFabricQueryParams) (*GetSiteFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/fabric-site"

	queryString, _ := query.Values(getSiteFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetSiteFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSiteFromSDAFabricResponse)
	return result, response, err

}

// GetVNFromSDAFabricQueryParams defines the query parameters for this request
type GetVNFromSDAFabricQueryParams struct {
	VirtualNetworkName string `url:"virtualNetworkName,omitempty"` // virtualNetworkName
	SiteNameHierarchy  string `url:"siteNameHierarchy,omitempty"`  // siteNameHierarchy
}

// GetVNFromSDAFabric getVNFromSDAFabric
/* Get virtual network (VN) from SDA Fabric
@param virtualNetworkName virtualNetworkName
@param siteNameHierarchy siteNameHierarchy
*/
func (s *SDAService) GetVNFromSDAFabric(getVNFromSDAFabricQueryParams *GetVNFromSDAFabricQueryParams) (*GetVNFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/virtual-network"

	queryString, _ := query.Values(getVNFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetVNFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetVNFromSDAFabricResponse)
	return result, response, err

}

// GetsBorderDeviceDetailFromSDAFabricQueryParams defines the query parameters for this request
type GetsBorderDeviceDetailFromSDAFabricQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` // Device IP Address
}

// GetsBorderDeviceDetailFromSDAFabric getsBorderDeviceDetailFromSDAFabric
/* Gets border device detail from SDA Fabric
@param deviceIPAddress Device IP Address
*/
func (s *SDAService) GetsBorderDeviceDetailFromSDAFabric(getsBorderDeviceDetailFromSDAFabricQueryParams *GetsBorderDeviceDetailFromSDAFabricQueryParams) (*GetsBorderDeviceDetailFromSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(getsBorderDeviceDetailFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetsBorderDeviceDetailFromSDAFabricResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetsBorderDeviceDetailFromSDAFabricResponse)
	return result, response, err

}

// UpdateDefaultAuthenticationProfileInSDAFabric updateDefaultAuthenticationProfileInSDAFabric
/* Update default authentication profile in SDA Fabric
 */
func (s *SDAService) UpdateDefaultAuthenticationProfileInSDAFabric(updateDefaultAuthenticationProfileInSDAFabricRequest *UpdateDefaultAuthenticationProfileInSDAFabricRequest) (*UpdateDefaultAuthenticationProfileInSDAFabricResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/authentication-profile"

	response, err := RestyClient.R().
		SetBody(updateDefaultAuthenticationProfileInSDAFabricRequest).
		SetResult(&UpdateDefaultAuthenticationProfileInSDAFabricResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateDefaultAuthenticationProfileInSDAFabricResponse)
	return result, response, err

}
