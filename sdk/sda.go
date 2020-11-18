package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SDAService is the service to communicate with the SDA API endpoint
type SDAService service

// AddControlPlaneDeviceInSDAFabricRequest is the addControlPlaneDeviceInSDAFabricRequest definition
type AddControlPlaneDeviceInSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// AddDefaultAuthenticationProfileInSDAFabricRequest is the addDefaultAuthenticationProfileInSDAFabricRequest definition
type AddDefaultAuthenticationProfileInSDAFabricRequest struct {
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// AddEdgeDeviceInSDAFabricRequest is the addEdgeDeviceInSDAFabricRequest definition
type AddEdgeDeviceInSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// AddFabricRequest is the addFabricRequest definition
type AddFabricRequest struct {
	FabricName string `json:"fabricName,omitempty"` //
}

// AddIPPoolInSDAVirtualNetworkRequest is the addIPPoolInSDAVirtualNetworkRequest definition
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

// AddPortAssignmentForAccessPointInSDAFabricRequest is the addPortAssignmentForAccessPointInSDAFabricRequest definition
type AddPortAssignmentForAccessPointInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// AddPortAssignmentForUserDeviceInSDAFabricRequest is the addPortAssignmentForUserDeviceInSDAFabricRequest definition
type AddPortAssignmentForUserDeviceInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// AddSiteInSDAFabricRequest is the addSiteInSDAFabricRequest definition
type AddSiteInSDAFabricRequest struct {
	FabricName        string `json:"fabricName,omitempty"`        //
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` //
}

// AddVNInSDAFabricRequest is the addVNInSDAFabricRequest definition
type AddVNInSDAFabricRequest struct {
	SiteNameHierarchy  string `json:"siteNameHierarchy,omitempty"`  //
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` //
}

// AddsBorderDeviceInSDAFabricRequest is the addsBorderDeviceInSDAFabricRequest definition
type AddsBorderDeviceInSDAFabricRequest struct {
	BorderSessionType                 string                                                           `json:"borderSessionType,omitempty"`                 //
	ConnectedToInternet               bool                                                             `json:"connectedToInternet,omitempty"`               //
	DeviceManagementIPAddress         string                                                           `json:"deviceManagementIpAddress,omitempty"`         //
	ExternalConnectivityIPPoolName    string                                                           `json:"externalConnectivityIpPoolName,omitempty"`    //
	ExternalConnectivitySettings      []AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettings `json:"externalConnectivitySettings,omitempty"`      //
	ExternalDomainRoutingProtocolName string                                                           `json:"externalDomainRoutingProtocolName,omitempty"` //
	InternalAutonomouSystemNumber     string                                                           `json:"internalAutonomouSystemNumber,omitempty"`     //
	SiteNameHierarchy                 string                                                           `json:"siteNameHierarchy,omitempty"`                 //
}

// AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettings is the addsBorderDeviceInSDAFabricRequestExternalConnectivitySettings definition
type AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettings struct {
	ExternalAutonomouSystemNumber string                                                                    `json:"externalAutonomouSystemNumber,omitempty"` //
	InterfaceName                 string                                                                    `json:"interfaceName,omitempty"`                 //
	L3Handoff                     []AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"`                     //
}

// AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff is the addsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff definition
type AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3Handoff struct {
	VirtualNetwork AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"` //
}

// AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3HandoffVirtualNetwork is the addsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3HandoffVirtualNetwork definition
type AddsBorderDeviceInSDAFabricRequestExternalConnectivitySettingsL3HandoffVirtualNetwork struct {
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"` //
}

// DeleteDefaultAuthenticationProfileFromSDAFabricRequest is the deleteDefaultAuthenticationProfileFromSDAFabricRequest definition
type DeleteDefaultAuthenticationProfileFromSDAFabricRequest struct {
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// DeleteIPPoolFromSDAVirtualNetworkRequest is the deleteIPPoolFromSDAVirtualNetworkRequest definition
type DeleteIPPoolFromSDAVirtualNetworkRequest struct {
	AuthenticationPolicyName string `json:"authenticationPolicyName,omitempty"` //
	IPPoolName               string `json:"ipPoolName,omitempty"`               //
	IsL2FloodingEnabled      bool   `json:"isL2FloodingEnabled,omitempty"`      //
	ScalableGroupName        string `json:"scalableGroupName,omitempty"`        //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
	TrafficType              string `json:"trafficType,omitempty"`              //
	VirtualNetworkName       string `json:"virtualNetworkName,omitempty"`       //
}

// DeletePortAssignmentForAccessPointInSDAFabricRequest is the deletePortAssignmentForAccessPointInSDAFabricRequest definition
type DeletePortAssignmentForAccessPointInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// DeletePortAssignmentForUserDeviceInSDAFabricRequest is the deletePortAssignmentForUserDeviceInSDAFabricRequest definition
type DeletePortAssignmentForUserDeviceInSDAFabricRequest struct {
	AuthenticateTemplateName  string `json:"authenticateTemplateName,omitempty"`  //
	DataIPAddressPoolName     string `json:"dataIpAddressPoolName,omitempty"`     //
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	InterfaceName             string `json:"interfaceName,omitempty"`             //
	ScalableGroupName         string `json:"scalableGroupName,omitempty"`         //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
	VoiceIPAddressPoolName    string `json:"voiceIpAddressPoolName,omitempty"`    //
}

// DeleteSDAFabricRequest is the deleteSDAFabricRequest definition
type DeleteSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// DeleteSiteFromSDAFabricRequest is the deleteSiteFromSDAFabricRequest definition
type DeleteSiteFromSDAFabricRequest struct {
	DeviceManagementIPAddress string `json:"deviceManagementIpAddress,omitempty"` //
	SiteNameHierarchy         string `json:"siteNameHierarchy,omitempty"`         //
}

// UpdateDefaultAuthenticationProfileInSDAFabricRequest is the updateDefaultAuthenticationProfileInSDAFabricRequest definition
type UpdateDefaultAuthenticationProfileInSDAFabricRequest struct {
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// AddControlPlaneDeviceInSDAFabricResponse is the addControlPlaneDeviceInSDAFabricResponse definition
type AddControlPlaneDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddDefaultAuthenticationProfileInSDAFabricResponse is the addDefaultAuthenticationProfileInSDAFabricResponse definition
type AddDefaultAuthenticationProfileInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddEdgeDeviceInSDAFabricResponse is the addEdgeDeviceInSDAFabricResponse definition
type AddEdgeDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddFabricResponse is the addFabricResponse definition
type AddFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddIPPoolInSDAVirtualNetworkResponse is the addIPPoolInSDAVirtualNetworkResponse definition
type AddIPPoolInSDAVirtualNetworkResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddPortAssignmentForAccessPointInSDAFabricResponse is the addPortAssignmentForAccessPointInSDAFabricResponse definition
type AddPortAssignmentForAccessPointInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddPortAssignmentForUserDeviceInSDAFabricResponse is the addPortAssignmentForUserDeviceInSDAFabricResponse definition
type AddPortAssignmentForUserDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddSiteInSDAFabricResponse is the addSiteInSDAFabricResponse definition
type AddSiteInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// AddVNInSDAFabricResponse is the addVNInSDAFabricResponse definition
type AddVNInSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// AddVNInSDAFabricResponseRoles is the addVNInSDAFabricResponseRoles definition
type AddVNInSDAFabricResponseRoles []string

// AddsBorderDeviceInSDAFabricResponse is the addsBorderDeviceInSDAFabricResponse definition
type AddsBorderDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteControlPlaneDeviceInSDAFabricResponse is the deleteControlPlaneDeviceInSDAFabricResponse definition
type DeleteControlPlaneDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteDefaultAuthenticationProfileFromSDAFabricResponse is the deleteDefaultAuthenticationProfileFromSDAFabricResponse definition
type DeleteDefaultAuthenticationProfileFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteEdgeDeviceFromSDAFabricResponse is the deleteEdgeDeviceFromSDAFabricResponse definition
type DeleteEdgeDeviceFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteIPPoolFromSDAVirtualNetworkResponse is the deleteIPPoolFromSDAVirtualNetworkResponse definition
type DeleteIPPoolFromSDAVirtualNetworkResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeletePortAssignmentForAccessPointInSDAFabricResponse is the deletePortAssignmentForAccessPointInSDAFabricResponse definition
type DeletePortAssignmentForAccessPointInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeletePortAssignmentForUserDeviceInSDAFabricResponse is the deletePortAssignmentForUserDeviceInSDAFabricResponse definition
type DeletePortAssignmentForUserDeviceInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteSDAFabricResponse is the deleteSDAFabricResponse definition
type DeleteSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteSiteFromSDAFabricResponse is the deleteSiteFromSDAFabricResponse definition
type DeleteSiteFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeleteVNFromSDAFabricResponse is the deleteVNFromSDAFabricResponse definition
type DeleteVNFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// DeleteVNFromSDAFabricResponseRoles is the deleteVNFromSDAFabricResponseRoles definition
type DeleteVNFromSDAFabricResponseRoles []string

// DeletesBorderDeviceFromSDAFabricResponse is the deletesBorderDeviceFromSDAFabricResponse definition
type DeletesBorderDeviceFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// GetControlPlaneDeviceFromSDAFabricResponse is the getControlPlaneDeviceFromSDAFabricResponse definition
type GetControlPlaneDeviceFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetControlPlaneDeviceFromSDAFabricResponseRoles is the getControlPlaneDeviceFromSDAFabricResponseRoles definition
type GetControlPlaneDeviceFromSDAFabricResponseRoles []string

// GetDefaultAuthenticationProfileFromSDAFabricResponse is the getDefaultAuthenticationProfileFromSDAFabricResponse definition
type GetDefaultAuthenticationProfileFromSDAFabricResponse struct {
	AuthenticateTemplateID   string `json:"authenticateTemplateId,omitempty"`   //
	AuthenticateTemplateName string `json:"authenticateTemplateName,omitempty"` //
	SiteNameHierarchy        string `json:"siteNameHierarchy,omitempty"`        //
}

// GetDeviceInfoFromSDAFabricResponse is the getDeviceInfoFromSDAFabricResponse definition
type GetDeviceInfoFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetDeviceInfoFromSDAFabricResponseRoles is the getDeviceInfoFromSDAFabricResponseRoles definition
type GetDeviceInfoFromSDAFabricResponseRoles []string

// GetDeviceRoleInSDAFabricResponse is the getDeviceRoleInSDAFabricResponse definition
type GetDeviceRoleInSDAFabricResponse struct {
	Response GetDeviceRoleInSDAFabricResponseResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}

// GetDeviceRoleInSDAFabricResponseResponse is the getDeviceRoleInSDAFabricResponseResponse definition
type GetDeviceRoleInSDAFabricResponseResponse struct {
	Description string   `json:"description,omitempty"` //
	Roles       []string `json:"roles,omitempty"`       //
	Status      string   `json:"status,omitempty"`      //
}

// GetDeviceRoleInSDAFabricResponseResponseRoles is the getDeviceRoleInSDAFabricResponseResponseRoles definition
type GetDeviceRoleInSDAFabricResponseResponseRoles []string

// GetEdgeDeviceFromSDAFabricResponse is the getEdgeDeviceFromSDAFabricResponse definition
type GetEdgeDeviceFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetEdgeDeviceFromSDAFabricResponseRoles is the getEdgeDeviceFromSDAFabricResponseRoles definition
type GetEdgeDeviceFromSDAFabricResponseRoles []string

// GetIPPoolFromSDAVirtualNetworkResponse is the getIPPoolFromSDAVirtualNetworkResponse definition
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

// GetPortAssignmentForAccessPointInSDAFabricResponse is the getPortAssignmentForAccessPointInSDAFabricResponse definition
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

// GetPortAssignmentForUserDeviceInSDAFabricResponse is the getPortAssignmentForUserDeviceInSDAFabricResponse definition
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

// GetSDAFabricCountResponse is the getSDAFabricCountResponse definition
type GetSDAFabricCountResponse struct {
	Response GetSDAFabricCountResponseResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  //
}

// GetSDAFabricCountResponseResponse is the getSDAFabricCountResponseResponse definition
type GetSDAFabricCountResponseResponse struct {
	Description string `json:"description,omitempty"` //
	FabricCount string `json:"fabricCount,omitempty"` //
	Status      string `json:"status,omitempty"`      //
}

// GetSDAFabricInfoResponse is the getSDAFabricInfoResponse definition
type GetSDAFabricInfoResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// GetSiteFromSDAFabricResponse is the getSiteFromSDAFabricResponse definition
type GetSiteFromSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// GetVNFromSDAFabricResponse is the getVNFromSDAFabricResponse definition
type GetVNFromSDAFabricResponse struct {
	Description               string   `json:"description,omitempty"`               //
	DeviceManagementIPAddress string   `json:"deviceManagementIpAddress,omitempty"` //
	Name                      string   `json:"name,omitempty"`                      //
	Roles                     []string `json:"roles,omitempty"`                     //
	SiteHierarchy             string   `json:"siteHierarchy,omitempty"`             //
	Status                    string   `json:"status,omitempty"`                    //
}

// GetVNFromSDAFabricResponseRoles is the getVNFromSDAFabricResponseRoles definition
type GetVNFromSDAFabricResponseRoles []string

// GetsBorderDeviceDetailFromSDAFabricResponse is the getsBorderDeviceDetailFromSDAFabricResponse definition
type GetsBorderDeviceDetailFromSDAFabricResponse struct {
	Description string                                             `json:"description,omitempty"` //
	Payload     GetsBorderDeviceDetailFromSDAFabricResponsePayload `json:"payload,omitempty"`     //
	Status      string                                             `json:"status,omitempty"`      //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayload is the getsBorderDeviceDetailFromSDAFabricResponsePayload definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayload struct {
	AkcSettingsCfs                 []string                                                              `json:"akcSettingsCfs,omitempty"`                 //
	AuthEntityClass                int                                                                   `json:"authEntityClass,omitempty"`                //
	AuthEntityID                   int                                                                   `json:"authEntityId,omitempty"`                   //
	CfsChangeInfo                  []string                                                              `json:"cfsChangeInfo,omitempty"`                  //
	Configs                        []string                                                              `json:"configs,omitempty"`                        //
	CreateTime                     int                                                                   `json:"createTime,omitempty"`                     //
	CustomProvisions               []string                                                              `json:"customProvisions,omitempty"`               //
	DeployPending                  string                                                                `json:"deployPending,omitempty"`                  //
	Deployed                       bool                                                                  `json:"deployed,omitempty"`                       //
	DeviceInterfaceInfo            []string                                                              `json:"deviceInterfaceInfo,omitempty"`            //
	DeviceSettings                 GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettings      `json:"deviceSettings,omitempty"`                 //
	DisplayName                    string                                                                `json:"displayName,omitempty"`                    //
	ID                             string                                                                `json:"id,omitempty"`                             //
	InstanceID                     int                                                                   `json:"instanceId,omitempty"`                     //
	InstanceTenantID               string                                                                `json:"instanceTenantId,omitempty"`               //
	InstanceVersion                int                                                                   `json:"instanceVersion,omitempty"`                //
	IsSeeded                       bool                                                                  `json:"isSeeded,omitempty"`                       //
	IsStale                        bool                                                                  `json:"isStale,omitempty"`                        //
	LastUpdateTime                 int                                                                   `json:"lastUpdateTime,omitempty"`                 //
	ManagedSites                   []string                                                              `json:"managedSites,omitempty"`                   //
	Name                           string                                                                `json:"name,omitempty"`                           //
	Namespace                      string                                                                `json:"namespace,omitempty"`                      //
	NetworkDeviceID                string                                                                `json:"networkDeviceId,omitempty"`                //
	NetworkWideSettings            GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettings `json:"networkWideSettings,omitempty"`            //
	OtherDevice                    []string                                                              `json:"otherDevice,omitempty"`                    //
	ProvisioningState              string                                                                `json:"provisioningState,omitempty"`              //
	ResourceVersion                int                                                                   `json:"resourceVersion,omitempty"`                //
	Roles                          []string                                                              `json:"roles,omitempty"`                          //
	SaveWanConnectivityDetailsOnly bool                                                                  `json:"saveWanConnectivityDetailsOnly,omitempty"` //
	SiteID                         string                                                                `json:"siteId,omitempty"`                         //
	TargetIDList                   []string                                                              `json:"targetIdList,omitempty"`                   //
	TransitNetworks                []GetsBorderDeviceDetailFromSDAFabricResponsePayloadTransitNetworks   `json:"transitNetworks,omitempty"`                //
	Type                           string                                                                `json:"type,omitempty"`                           //
	VirtualNetwork                 []string                                                              `json:"virtualNetwork,omitempty"`                 //
	WLAN                           []string                                                              `json:"wlan,omitempty"`                           //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadAkcSettingsCfs is the getsBorderDeviceDetailFromSDAFabricResponsePayloadAkcSettingsCfs definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadAkcSettingsCfs []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadCfsChangeInfo is the getsBorderDeviceDetailFromSDAFabricResponsePayloadCfsChangeInfo definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadCfsChangeInfo []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadConfigs is the getsBorderDeviceDetailFromSDAFabricResponsePayloadConfigs definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadConfigs []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadCustomProvisions is the getsBorderDeviceDetailFromSDAFabricResponsePayloadCustomProvisions definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadCustomProvisions []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceInterfaceInfo is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceInterfaceInfo definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceInterfaceInfo []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettings is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettings definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettings struct {
	ConnectedTo                   []string                                                                                  `json:"connectedTo,omitempty"`                   //
	CPU                           int                                                                                       `json:"cpu,omitempty"`                           //
	DeployPending                 string                                                                                    `json:"deployPending,omitempty"`                 //
	DhcpEnabled                   bool                                                                                      `json:"dhcpEnabled,omitempty"`                   //
	DisplayName                   string                                                                                    `json:"displayName,omitempty"`                   //
	ExtConnectivitySettings       []GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettings `json:"extConnectivitySettings,omitempty"`       //
	ExternalConnectivityIPPool    string                                                                                    `json:"externalConnectivityIpPool,omitempty"`    //
	ExternalDomainRoutingProtocol string                                                                                    `json:"externalDomainRoutingProtocol,omitempty"` //
	ID                            string                                                                                    `json:"id,omitempty"`                            //
	InstanceID                    int                                                                                       `json:"instanceId,omitempty"`                    //
	InstanceTenantID              string                                                                                    `json:"instanceTenantId,omitempty"`              //
	InstanceVersion               int                                                                                       `json:"instanceVersion,omitempty"`               //
	InternalDomainProtocolNumber  string                                                                                    `json:"internalDomainProtocolNumber,omitempty"`  //
	Memory                        int                                                                                       `json:"memory,omitempty"`                        //
	NodeType                      []string                                                                                  `json:"nodeType,omitempty"`                      //
	Storage                       int                                                                                       `json:"storage,omitempty"`                       //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsConnectedTo is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsConnectedTo definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsConnectedTo []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettings is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettings definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettings struct {
	DeployPending                string                                                                                             `json:"deployPending,omitempty"`                //
	DisplayName                  string                                                                                             `json:"displayName,omitempty"`                  //
	ExternalDomainProtocolNumber string                                                                                             `json:"externalDomainProtocolNumber,omitempty"` //
	ID                           string                                                                                             `json:"id,omitempty"`                           //
	InstanceID                   int                                                                                                `json:"instanceId,omitempty"`                   //
	InstanceTenantID             string                                                                                             `json:"instanceTenantId,omitempty"`             //
	InstanceVersion              int                                                                                                `json:"instanceVersion,omitempty"`              //
	InterfaceUUID                string                                                                                             `json:"interfaceUuid,omitempty"`                //
	L2Handoff                    []string                                                                                           `json:"l2Handoff,omitempty"`                    //
	L3Handoff                    []GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3Handoff `json:"l3Handoff,omitempty"`                    //
	PolicyPropagationEnabled     bool                                                                                               `json:"policyPropagationEnabled,omitempty"`     //
	PolicySgtTag                 int                                                                                                `json:"policySgtTag,omitempty"`                 //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL2Handoff is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL2Handoff definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL2Handoff []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3Handoff is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3Handoff definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3Handoff struct {
	DeployPending    string                                                                                                         `json:"deployPending,omitempty"`    //
	DisplayName      string                                                                                                         `json:"displayName,omitempty"`      //
	ID               string                                                                                                         `json:"id,omitempty"`               //
	InstanceID       int                                                                                                            `json:"instanceId,omitempty"`       //
	InstanceTenantID string                                                                                                         `json:"instanceTenantId,omitempty"` //
	InstanceVersion  int                                                                                                            `json:"instanceVersion,omitempty"`  //
	LocalIPAddress   string                                                                                                         `json:"localIpAddress,omitempty"`   //
	RemoteIPAddress  string                                                                                                         `json:"remoteIpAddress,omitempty"`  //
	VirtualNetwork   GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork `json:"virtualNetwork,omitempty"`   //
	VLANID           int                                                                                                            `json:"vlanId,omitempty"`           //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsExtConnectivitySettingsL3HandoffVirtualNetwork struct {
	IDRef string `json:"idRef,omitempty"` //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsNodeType is the getsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsNodeType definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadDeviceSettingsNodeType []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadManagedSites is the getsBorderDeviceDetailFromSDAFabricResponsePayloadManagedSites definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadManagedSites []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettings is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettings definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettings struct {
	AAA              []string                                                                    `json:"aaa,omitempty"`              //
	Cmx              []string                                                                    `json:"cmx,omitempty"`              //
	DeployPending    string                                                                      `json:"deployPending,omitempty"`    //
	Dhcp             []GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcp `json:"dhcp,omitempty"`             //
	DisplayName      string                                                                      `json:"displayName,omitempty"`      //
	DNS              []GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNS  `json:"dns,omitempty"`              //
	ID               string                                                                      `json:"id,omitempty"`               //
	InstanceID       int                                                                         `json:"instanceId,omitempty"`       //
	InstanceTenantID string                                                                      `json:"instanceTenantId,omitempty"` //
	InstanceVersion  int                                                                         `json:"instanceVersion,omitempty"`  //
	Ldap             []string                                                                    `json:"ldap,omitempty"`             //
	NativeVLAN       []string                                                                    `json:"nativeVlan,omitempty"`       //
	Netflow          []string                                                                    `json:"netflow,omitempty"`          //
	Ntp              []string                                                                    `json:"ntp,omitempty"`              //
	SNMP             []string                                                                    `json:"snmp,omitempty"`             //
	Syslogs          []string                                                                    `json:"syslogs,omitempty"`          //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsAAA is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsAAA definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsAAA []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsCmx is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsCmx definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsCmx []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNS is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNS definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNS struct {
	DomainName string                                                                     `json:"domainName,omitempty"` //
	ID         string                                                                     `json:"id,omitempty"`         //
	IP         GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNSIP `json:"ip,omitempty"`         //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNSIP is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNSIP definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDNSIP struct {
	Address       string `json:"address,omitempty"`       //
	AddressType   string `json:"addressType,omitempty"`   //
	ID            string `json:"id,omitempty"`            //
	PaddedAddress string `json:"paddedAddress,omitempty"` //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcp is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcp definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcp struct {
	ID        string                                                                             `json:"id,omitempty"`        //
	IPAddress GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcpIPAddress `json:"ipAddress,omitempty"` //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcpIPAddress is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcpIPAddress definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsDhcpIPAddress struct {
	Address       string `json:"address,omitempty"`       //
	AddressType   string `json:"addressType,omitempty"`   //
	ID            string `json:"id,omitempty"`            //
	PaddedAddress string `json:"paddedAddress,omitempty"` //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsLdap is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsLdap definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsLdap []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNativeVLAN is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNativeVLAN definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNativeVLAN []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNetflow is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNetflow definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNetflow []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNtp is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNtp definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsNtp []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsSNMP is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsSNMP definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsSNMP []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsSyslogs is the getsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsSyslogs definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadNetworkWideSettingsSyslogs []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadOtherDevice is the getsBorderDeviceDetailFromSDAFabricResponsePayloadOtherDevice definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadOtherDevice []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadRoles is the getsBorderDeviceDetailFromSDAFabricResponsePayloadRoles definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadRoles []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadTargetIDList is the getsBorderDeviceDetailFromSDAFabricResponsePayloadTargetIDList definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadTargetIDList []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadTransitNetworks is the getsBorderDeviceDetailFromSDAFabricResponsePayloadTransitNetworks definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadTransitNetworks struct {
	IDRef string `json:"idRef,omitempty"` //
}

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadVirtualNetwork is the getsBorderDeviceDetailFromSDAFabricResponsePayloadVirtualNetwork definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadVirtualNetwork []string

// GetsBorderDeviceDetailFromSDAFabricResponsePayloadWLAN is the getsBorderDeviceDetailFromSDAFabricResponsePayloadWLAN definition
type GetsBorderDeviceDetailFromSDAFabricResponsePayloadWLAN []string

// UpdateDefaultAuthenticationProfileInSDAFabricResponse is the updateDefaultAuthenticationProfileInSDAFabricResponse definition
type UpdateDefaultAuthenticationProfileInSDAFabricResponse struct {
	Description        string `json:"description,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Status             string `json:"status,omitempty"`             //
}

// DeletesBorderDeviceFromSDAFabricRequest is the DeletesBorderDeviceFromSDAFabricRequest definition
type DeletesBorderDeviceFromSDAFabricRequest []string

// AddControlPlaneDeviceInSDAFabric addControlPlaneDeviceInSDAFabric
/* Add control plane device in SDA Fabric
 */
func (s *SDAService) AddControlPlaneDeviceInSDAFabric(addControlPlaneDeviceInSDAFabricRequest *[]AddControlPlaneDeviceInSDAFabricRequest) (*AddControlPlaneDeviceInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddDefaultAuthenticationProfileInSDAFabric(addDefaultAuthenticationProfileInSDAFabricRequest *[]AddDefaultAuthenticationProfileInSDAFabricRequest) (*AddDefaultAuthenticationProfileInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddEdgeDeviceInSDAFabric(addEdgeDeviceInSDAFabricRequest *[]AddEdgeDeviceInSDAFabricRequest) (*AddEdgeDeviceInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddFabric(addFabricRequest *[]AddFabricRequest) (*AddFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddIPPoolInSDAVirtualNetwork(addIPPoolInSDAVirtualNetworkRequest *[]AddIPPoolInSDAVirtualNetworkRequest) (*AddIPPoolInSDAVirtualNetworkResponse, *resty.Response, error) {

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
func (s *SDAService) AddPortAssignmentForAccessPointInSDAFabric(addPortAssignmentForAccessPointInSDAFabricRequest *[]AddPortAssignmentForAccessPointInSDAFabricRequest) (*AddPortAssignmentForAccessPointInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddPortAssignmentForUserDeviceInSDAFabric(addPortAssignmentForUserDeviceInSDAFabricRequest *[]AddPortAssignmentForUserDeviceInSDAFabricRequest) (*AddPortAssignmentForUserDeviceInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddSiteInSDAFabric(addSiteInSDAFabricRequest *[]AddSiteInSDAFabricRequest) (*AddSiteInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddVNInSDAFabric(addVNInSDAFabricRequest *[]AddVNInSDAFabricRequest) (*AddVNInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) AddsBorderDeviceInSDAFabric(addsBorderDeviceInSDAFabricRequest *[]AddsBorderDeviceInSDAFabricRequest) (*AddsBorderDeviceInSDAFabricResponse, *resty.Response, error) {

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
func (s *SDAService) DeleteDefaultAuthenticationProfileFromSDAFabric(deleteDefaultAuthenticationProfileFromSDAFabricQueryParams *DeleteDefaultAuthenticationProfileFromSDAFabricQueryParams, deleteDefaultAuthenticationProfileFromSDAFabricRequest *[]DeleteDefaultAuthenticationProfileFromSDAFabricRequest) (*resty.Response, error) {

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
func (s *SDAService) DeleteIPPoolFromSDAVirtualNetwork(deleteIPPoolFromSDAVirtualNetworkQueryParams *DeleteIPPoolFromSDAVirtualNetworkQueryParams, deleteIPPoolFromSDAVirtualNetworkRequest *[]DeleteIPPoolFromSDAVirtualNetworkRequest) (*resty.Response, error) {

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
func (s *SDAService) DeletePortAssignmentForAccessPointInSDAFabric(deletePortAssignmentForAccessPointInSDAFabricQueryParams *DeletePortAssignmentForAccessPointInSDAFabricQueryParams, deletePortAssignmentForAccessPointInSDAFabricRequest *[]DeletePortAssignmentForAccessPointInSDAFabricRequest) (*resty.Response, error) {

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
func (s *SDAService) DeletePortAssignmentForUserDeviceInSDAFabric(deletePortAssignmentForUserDeviceInSDAFabricQueryParams *DeletePortAssignmentForUserDeviceInSDAFabricQueryParams, deletePortAssignmentForUserDeviceInSDAFabricRequest *[]DeletePortAssignmentForUserDeviceInSDAFabricRequest) (*resty.Response, error) {

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
func (s *SDAService) DeleteSDAFabric(deleteSDAFabricQueryParams *DeleteSDAFabricQueryParams, deleteSDAFabricRequest *[]DeleteSDAFabricRequest) (*resty.Response, error) {

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
func (s *SDAService) DeleteSiteFromSDAFabric(deleteSiteFromSDAFabricQueryParams *DeleteSiteFromSDAFabricQueryParams, deleteSiteFromSDAFabricRequest *[]DeleteSiteFromSDAFabricRequest) (*resty.Response, error) {

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
func (s *SDAService) DeletesBorderDeviceFromSDAFabric(deletesBorderDeviceFromSDAFabricQueryParams *DeletesBorderDeviceFromSDAFabricQueryParams, deletesBorderDeviceFromSDAFabricRequest *DeletesBorderDeviceFromSDAFabricRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/sda/border-device"

	queryString, _ := query.Values(deletesBorderDeviceFromSDAFabricQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(deletesBorderDeviceFromSDAFabricRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

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
func (s *SDAService) UpdateDefaultAuthenticationProfileInSDAFabric(updateDefaultAuthenticationProfileInSDAFabricRequest *[]UpdateDefaultAuthenticationProfileInSDAFabricRequest) (*UpdateDefaultAuthenticationProfileInSDAFabricResponse, *resty.Response, error) {

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
