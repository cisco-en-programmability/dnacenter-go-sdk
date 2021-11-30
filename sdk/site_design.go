package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SiteDesignService is the service to communicate with the SiteDesign API endpoint
type SiteDesignService service

// CreateNFVProfileRequest is the createNFVProfileRequest definition
type CreateNFVProfileRequest struct {
	Device      []CreateNFVProfileRequestDevice `json:"device,omitempty"`      //
	ProfileName string                          `json:"profileName,omitempty"` //
}

// CreateNFVProfileRequestDevice is the createNFVProfileRequestDevice definition
type CreateNFVProfileRequestDevice struct {
	CustomNetworks                  []CreateNFVProfileRequestDeviceCustomNetworks         `json:"customNetworks,omitempty"`                  //
	CustomTemplate                  []CreateNFVProfileRequestDeviceCustomTemplate         `json:"customTemplate,omitempty"`                  //
	DeviceTag                       string                                                `json:"deviceTag,omitempty"`                       //
	DeviceType                      string                                                `json:"deviceType,omitempty"`                      //
	DirectInternetAccessForFirewall bool                                                  `json:"directInternetAccessForFirewall,omitempty"` //
	ServiceProviderProfile          []CreateNFVProfileRequestDeviceServiceProviderProfile `json:"serviceProviderProfile,omitempty"`          //
	Services                        []CreateNFVProfileRequestDeviceServices               `json:"services,omitempty"`                        //
	VLANForL2                       []CreateNFVProfileRequestDeviceVLANForL2              `json:"vlanForL2,omitempty"`                       //
}

// CreateNFVProfileRequestDeviceCustomNetworks is the createNFVProfileRequestDeviceCustomNetworks definition
type CreateNFVProfileRequestDeviceCustomNetworks struct {
	ConnectionType    string                                                         `json:"connectionType,omitempty"`    //
	NetworkName       string                                                         `json:"networkName,omitempty"`       //
	ServicesToConnect []CreateNFVProfileRequestDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	VLANID            float64                                                        `json:"vlanId,omitempty"`            //
	VLANMode          string                                                         `json:"vlanMode,omitempty"`          //
}

// CreateNFVProfileRequestDeviceCustomNetworksServicesToConnect is the createNFVProfileRequestDeviceCustomNetworksServicesToConnect definition
type CreateNFVProfileRequestDeviceCustomNetworksServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` //
}

// CreateNFVProfileRequestDeviceCustomTemplate is the createNFVProfileRequestDeviceCustomTemplate definition
type CreateNFVProfileRequestDeviceCustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   //
	Template     string `json:"template,omitempty"`     //
	TemplateType string `json:"templateType,omitempty"` //
}

// CreateNFVProfileRequestDeviceServiceProviderProfile is the createNFVProfileRequestDeviceServiceProviderProfile definition
type CreateNFVProfileRequestDeviceServiceProviderProfile struct {
	Connect                    bool   `json:"connect,omitempty"`                    //
	ConnectDefaultGatewayOnWan bool   `json:"connectDefaultGatewayOnWan,omitempty"` //
	LinkType                   string `json:"linkType,omitempty"`                   //
	ServiceProvider            string `json:"serviceProvider,omitempty"`            //
}

// CreateNFVProfileRequestDeviceServices is the createNFVProfileRequestDeviceServices definition
type CreateNFVProfileRequestDeviceServices struct {
	FirewallMode string                                             `json:"firewallMode,omitempty"` //
	ImageName    string                                             `json:"imageName,omitempty"`    //
	ProfileType  string                                             `json:"profileType,omitempty"`  //
	ServiceName  string                                             `json:"serviceName,omitempty"`  //
	ServiceType  string                                             `json:"serviceType,omitempty"`  //
	VNicMapping  []CreateNFVProfileRequestDeviceServicesVNicMapping `json:"vNicMapping,omitempty"`  //
}

// CreateNFVProfileRequestDeviceServicesVNicMapping is the createNFVProfileRequestDeviceServicesVNicMapping definition
type CreateNFVProfileRequestDeviceServicesVNicMapping struct {
	AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
	NetworkType              string `json:"networkType,omitempty"`              //
}

// CreateNFVProfileRequestDeviceVLANForL2 is the createNFVProfileRequestDeviceVLANForL2 definition
type CreateNFVProfileRequestDeviceVLANForL2 struct {
	VLANDescription string  `json:"vlanDescription,omitempty"` //
	VLANID          float64 `json:"vlanId,omitempty"`          //
	VLANType        string  `json:"vlanType,omitempty"`        //
}

// NFVProvisioningDetailRequest is the nFVProvisioningDetailRequest definition
type NFVProvisioningDetailRequest struct {
	DeviceIP string `json:"device_ip,omitempty"` //
}

// ProvisionNFVRequest is the provisionNFVRequest definition
type ProvisionNFVRequest struct {
	Provisioning []ProvisionNFVRequestProvisioning `json:"provisioning,omitempty"` //
	SiteProfile  []ProvisionNFVRequestSiteProfile  `json:"siteProfile,omitempty"`  //
}

// ProvisionNFVRequestProvisioning is the provisionNFVRequestProvisioning definition
type ProvisionNFVRequestProvisioning struct {
	Device []ProvisionNFVRequestProvisioningDevice `json:"device,omitempty"` //
	Site   ProvisionNFVRequestProvisioningSite     `json:"site,omitempty"`   //
}

// ProvisionNFVRequestProvisioningDevice is the provisionNFVRequestProvisioningDevice definition
type ProvisionNFVRequestProvisioningDevice struct {
	CustomNetworks     []ProvisionNFVRequestProvisioningDeviceCustomNetworks   `json:"customNetworks,omitempty"`     //
	DeviceSerialNumber string                                                  `json:"deviceSerialNumber,omitempty"` //
	IP                 string                                                  `json:"ip,omitempty"`                 //
	ServiceProviders   []ProvisionNFVRequestProvisioningDeviceServiceProviders `json:"serviceProviders,omitempty"`   //
	Services           []ProvisionNFVRequestProvisioningDeviceServices         `json:"services,omitempty"`           //
	SubPools           []ProvisionNFVRequestProvisioningDeviceSubPools         `json:"subPools,omitempty"`           //
	TagName            string                                                  `json:"tagName,omitempty"`            //
	TemplateParam      ProvisionNFVRequestProvisioningDeviceTemplateParam      `json:"templateParam,omitempty"`      //
	VLAN               []ProvisionNFVRequestProvisioningDeviceVLAN             `json:"vlan,omitempty"`               //
}

// ProvisionNFVRequestProvisioningDeviceCustomNetworks is the provisionNFVRequestProvisioningDeviceCustomNetworks definition
type ProvisionNFVRequestProvisioningDeviceCustomNetworks struct {
	IPAddressPool string `json:"ipAddressPool,omitempty"` //
	Name          string `json:"name,omitempty"`          //
	Port          string `json:"port,omitempty"`          //
}

// ProvisionNFVRequestProvisioningDeviceServiceProviders is the provisionNFVRequestProvisioningDeviceServiceProviders definition
type ProvisionNFVRequestProvisioningDeviceServiceProviders struct {
	ServiceProvider string                                                            `json:"serviceProvider,omitempty"` //
	WanInterface    ProvisionNFVRequestProvisioningDeviceServiceProvidersWanInterface `json:"wanInterface,omitempty"`    //
}

// ProvisionNFVRequestProvisioningDeviceServiceProvidersWanInterface is the provisionNFVRequestProvisioningDeviceServiceProvidersWanInterface definition
type ProvisionNFVRequestProvisioningDeviceServiceProvidersWanInterface struct {
	Bandwidth     string `json:"bandwidth,omitempty"`     //
	Gateway       string `json:"gateway,omitempty"`       //
	InterfaceName string `json:"interfaceName,omitempty"` //
	IPAddress     string `json:"ipAddress,omitempty"`     //
	Subnetmask    string `json:"subnetmask,omitempty"`    //
}

// ProvisionNFVRequestProvisioningDeviceServices is the provisionNFVRequestProvisioningDeviceServices definition
type ProvisionNFVRequestProvisioningDeviceServices struct {
	AdminPasswordHash      string `json:"adminPasswordHash,omitempty"`      //
	CentralManagerIP       string `json:"centralManagerIP,omitempty"`       //
	CentralRegistrationKey string `json:"centralRegistrationKey,omitempty"` //
	CommonKey              string `json:"commonKey,omitempty"`              //
	Disk                   string `json:"disk,omitempty"`                   //
	Mode                   string `json:"mode,omitempty"`                   //
	SystemIP               string `json:"systemIp,omitempty"`               //
	Type                   string `json:"type,omitempty"`                   //
}

// ProvisionNFVRequestProvisioningDeviceSubPools is the provisionNFVRequestProvisioningDeviceSubPools definition
type ProvisionNFVRequestProvisioningDeviceSubPools struct {
	Gateway        string `json:"gateway,omitempty"`        //
	IPSubnet       string `json:"ipSubnet,omitempty"`       //
	Name           string `json:"name,omitempty"`           //
	ParentPoolName string `json:"parentPoolName,omitempty"` //
	Type           string `json:"type,omitempty"`           //
}

// ProvisionNFVRequestProvisioningDeviceTemplateParam is the provisionNFVRequestProvisioningDeviceTemplateParam definition
type ProvisionNFVRequestProvisioningDeviceTemplateParam struct {
	Asav  ProvisionNFVRequestProvisioningDeviceTemplateParamAsav  `json:"asav,omitempty"`  //
	Nfvis ProvisionNFVRequestProvisioningDeviceTemplateParamNfvis `json:"nfvis,omitempty"` //
}

// ProvisionNFVRequestProvisioningDeviceTemplateParamAsav is the provisionNFVRequestProvisioningDeviceTemplateParamAsav definition
type ProvisionNFVRequestProvisioningDeviceTemplateParamAsav struct {
	Var1 string `json:"var1,omitempty"` //
}

// ProvisionNFVRequestProvisioningDeviceTemplateParamNfvis is the provisionNFVRequestProvisioningDeviceTemplateParamNfvis definition
type ProvisionNFVRequestProvisioningDeviceTemplateParamNfvis struct {
	Var1 string `json:"var1,omitempty"` //
}

// ProvisionNFVRequestProvisioningDeviceVLAN is the provisionNFVRequestProvisioningDeviceVLAN definition
type ProvisionNFVRequestProvisioningDeviceVLAN struct {
	ID         string `json:"id,omitempty"`         //
	Interfaces string `json:"interfaces,omitempty"` //
	Network    string `json:"network,omitempty"`    //
	Type       string `json:"type,omitempty"`       //
}

// ProvisionNFVRequestProvisioningSite is the provisionNFVRequestProvisioningSite definition
type ProvisionNFVRequestProvisioningSite struct {
	Area            ProvisionNFVRequestProvisioningSiteArea     `json:"area,omitempty"`            //
	Building        ProvisionNFVRequestProvisioningSiteBuilding `json:"building,omitempty"`        //
	Floor           ProvisionNFVRequestProvisioningSiteFloor    `json:"floor,omitempty"`           //
	SiteProfileName string                                      `json:"siteProfileName,omitempty"` //
}

// ProvisionNFVRequestProvisioningSiteArea is the provisionNFVRequestProvisioningSiteArea definition
type ProvisionNFVRequestProvisioningSiteArea struct {
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// ProvisionNFVRequestProvisioningSiteBuilding is the provisionNFVRequestProvisioningSiteBuilding definition
type ProvisionNFVRequestProvisioningSiteBuilding struct {
	Address    string  `json:"address,omitempty"`    //
	Latitude   float64 `json:"latitude,omitempty"`   //
	Longitude  float64 `json:"longitude,omitempty"`  //
	Name       string  `json:"name,omitempty"`       //
	ParentName string  `json:"parentName,omitempty"` //
}

// ProvisionNFVRequestProvisioningSiteFloor is the provisionNFVRequestProvisioningSiteFloor definition
type ProvisionNFVRequestProvisioningSiteFloor struct {
	Height     float64 `json:"height,omitempty"`     //
	Length     float64 `json:"length,omitempty"`     //
	Name       string  `json:"name,omitempty"`       //
	ParentName string  `json:"parentName,omitempty"` //
	RfModel    string  `json:"rfModel,omitempty"`    //
	Width      float64 `json:"width,omitempty"`      //
}

// ProvisionNFVRequestSiteProfile is the provisionNFVRequestSiteProfile definition
type ProvisionNFVRequestSiteProfile struct {
	Device          []ProvisionNFVRequestSiteProfileDevice `json:"device,omitempty"`          //
	SiteProfileName string                                 `json:"siteProfileName,omitempty"` //
}

// ProvisionNFVRequestSiteProfileDevice is the provisionNFVRequestSiteProfileDevice definition
type ProvisionNFVRequestSiteProfileDevice struct {
	CustomNetworks   []ProvisionNFVRequestSiteProfileDeviceCustomNetworks   `json:"customNetworks,omitempty"`   //
	CustomServices   []ProvisionNFVRequestSiteProfileDeviceCustomServices   `json:"customServices,omitempty"`   //
	CustomTemplate   []ProvisionNFVRequestSiteProfileDeviceCustomTemplate   `json:"customTemplate,omitempty"`   //
	DeviceType       string                                                 `json:"deviceType,omitempty"`       //
	Dia              bool                                                   `json:"dia,omitempty"`              //
	ServiceProviders []ProvisionNFVRequestSiteProfileDeviceServiceProviders `json:"serviceProviders,omitempty"` //
	Services         []ProvisionNFVRequestSiteProfileDeviceServices         `json:"services,omitempty"`         //
	TagName          string                                                 `json:"tagName,omitempty"`          //
	VLAN             []ProvisionNFVRequestSiteProfileDeviceVLAN             `json:"vlan,omitempty"`             //
}

// ProvisionNFVRequestSiteProfileDeviceCustomNetworks is the provisionNFVRequestSiteProfileDeviceCustomNetworks definition
type ProvisionNFVRequestSiteProfileDeviceCustomNetworks struct {
	ConnectionType    string                                                                `json:"connectionType,omitempty"`    //
	Name              string                                                                `json:"name,omitempty"`              //
	NetworkMode       string                                                                `json:"networkMode,omitempty"`       //
	ServicesToConnect []ProvisionNFVRequestSiteProfileDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	VLAN              string                                                                `json:"vlan,omitempty"`              //
}

// ProvisionNFVRequestSiteProfileDeviceCustomNetworksServicesToConnect is the provisionNFVRequestSiteProfileDeviceCustomNetworksServicesToConnect definition
type ProvisionNFVRequestSiteProfileDeviceCustomNetworksServicesToConnect struct {
	Service string `json:"service,omitempty"` //
}

// ProvisionNFVRequestSiteProfileDeviceCustomServices is the provisionNFVRequestSiteProfileDeviceCustomServices definition
type ProvisionNFVRequestSiteProfileDeviceCustomServices struct {
	ApplicationType string                                                     `json:"applicationType,omitempty"` //
	ImageName       string                                                     `json:"imageName,omitempty"`       //
	Name            string                                                     `json:"name,omitempty"`            //
	Profile         string                                                     `json:"profile,omitempty"`         //
	Topology        ProvisionNFVRequestSiteProfileDeviceCustomServicesTopology `json:"topology,omitempty"`        //
}

// ProvisionNFVRequestSiteProfileDeviceCustomServicesTopology is the provisionNFVRequestSiteProfileDeviceCustomServicesTopology definition
type ProvisionNFVRequestSiteProfileDeviceCustomServicesTopology struct {
	AssignIP string `json:"assignIp,omitempty"` //
	Name     string `json:"name,omitempty"`     //
	Type     string `json:"type,omitempty"`     //
}

// ProvisionNFVRequestSiteProfileDeviceCustomTemplate is the provisionNFVRequestSiteProfileDeviceCustomTemplate definition
type ProvisionNFVRequestSiteProfileDeviceCustomTemplate struct {
	DeviceType string `json:"deviceType,omitempty"` //
	Template   string `json:"template,omitempty"`   //
}

// ProvisionNFVRequestSiteProfileDeviceServiceProviders is the provisionNFVRequestSiteProfileDeviceServiceProviders definition
type ProvisionNFVRequestSiteProfileDeviceServiceProviders struct {
	Connect         bool   `json:"connect,omitempty"`         //
	DefaultGateway  bool   `json:"defaultGateway,omitempty"`  //
	LinkType        string `json:"linkType,omitempty"`        //
	ServiceProvider string `json:"serviceProvider,omitempty"` //
}

// ProvisionNFVRequestSiteProfileDeviceServices is the provisionNFVRequestSiteProfileDeviceServices definition
type ProvisionNFVRequestSiteProfileDeviceServices struct {
	ImageName string                                               `json:"imageName,omitempty"` //
	Mode      string                                               `json:"mode,omitempty"`      //
	Name      string                                               `json:"name,omitempty"`      //
	Profile   string                                               `json:"profile,omitempty"`   //
	Topology  ProvisionNFVRequestSiteProfileDeviceServicesTopology `json:"topology,omitempty"`  //
	Type      string                                               `json:"type,omitempty"`      //
}

// ProvisionNFVRequestSiteProfileDeviceServicesTopology is the provisionNFVRequestSiteProfileDeviceServicesTopology definition
type ProvisionNFVRequestSiteProfileDeviceServicesTopology struct {
	AssignIP string `json:"assignIp,omitempty"` //
	Name     string `json:"name,omitempty"`     //
	Type     string `json:"type,omitempty"`     //
}

// ProvisionNFVRequestSiteProfileDeviceVLAN is the provisionNFVRequestSiteProfileDeviceVLAN definition
type ProvisionNFVRequestSiteProfileDeviceVLAN struct {
	ID   string `json:"id,omitempty"`   //
	Type string `json:"type,omitempty"` //
}

// UpdateNFVProfileRequest is the updateNFVProfileRequest definition
type UpdateNFVProfileRequest struct {
	Device []UpdateNFVProfileRequestDevice `json:"device,omitempty"` //
}

// UpdateNFVProfileRequestDevice is the updateNFVProfileRequestDevice definition
type UpdateNFVProfileRequestDevice struct {
	CurrentDeviceTag                string                                        `json:"currentDeviceTag,omitempty"`                //
	CustomNetworks                  []UpdateNFVProfileRequestDeviceCustomNetworks `json:"customNetworks,omitempty"`                  //
	CustomTemplate                  []UpdateNFVProfileRequestDeviceCustomTemplate `json:"customTemplate,omitempty"`                  //
	DeviceTag                       string                                        `json:"deviceTag,omitempty"`                       //
	DirectInternetAccessForFirewall bool                                          `json:"directInternetAccessForFirewall,omitempty"` //
	Services                        []UpdateNFVProfileRequestDeviceServices       `json:"services,omitempty"`                        //
	VLANForL2                       []UpdateNFVProfileRequestDeviceVLANForL2      `json:"vlanForL2,omitempty"`                       //
}

// UpdateNFVProfileRequestDeviceCustomNetworks is the updateNFVProfileRequestDeviceCustomNetworks definition
type UpdateNFVProfileRequestDeviceCustomNetworks struct {
	ConnectionType    string                                                         `json:"connectionType,omitempty"`    //
	NetworkName       string                                                         `json:"networkName,omitempty"`       //
	ServicesToConnect []UpdateNFVProfileRequestDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	VLANID            float64                                                        `json:"vlanId,omitempty"`            //
	VLANMode          string                                                         `json:"vlanMode,omitempty"`          //
}

// UpdateNFVProfileRequestDeviceCustomNetworksServicesToConnect is the updateNFVProfileRequestDeviceCustomNetworksServicesToConnect definition
type UpdateNFVProfileRequestDeviceCustomNetworksServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` //
}

// UpdateNFVProfileRequestDeviceCustomTemplate is the updateNFVProfileRequestDeviceCustomTemplate definition
type UpdateNFVProfileRequestDeviceCustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   //
	Template     string `json:"template,omitempty"`     //
	TemplateType string `json:"templateType,omitempty"` //
}

// UpdateNFVProfileRequestDeviceServices is the updateNFVProfileRequestDeviceServices definition
type UpdateNFVProfileRequestDeviceServices struct {
	FirewallMode string                                             `json:"firewallMode,omitempty"` //
	ImageName    string                                             `json:"imageName,omitempty"`    //
	ProfileType  string                                             `json:"profileType,omitempty"`  //
	ServiceName  string                                             `json:"serviceName,omitempty"`  //
	ServiceType  string                                             `json:"serviceType,omitempty"`  //
	VNicMapping  []UpdateNFVProfileRequestDeviceServicesVNicMapping `json:"vNicMapping,omitempty"`  //
}

// UpdateNFVProfileRequestDeviceServicesVNicMapping is the updateNFVProfileRequestDeviceServicesVNicMapping definition
type UpdateNFVProfileRequestDeviceServicesVNicMapping struct {
	AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
	NetworkType              string `json:"networkType,omitempty"`              //
}

// UpdateNFVProfileRequestDeviceVLANForL2 is the updateNFVProfileRequestDeviceVLANForL2 definition
type UpdateNFVProfileRequestDeviceVLANForL2 struct {
	VLANDescription string  `json:"vlanDescription,omitempty"` //
	VLANID          float64 `json:"vlanId,omitempty"`          //
	VLANType        string  `json:"vlanType,omitempty"`        //
}

// CreateNFVProfileResponse is the createNFVProfileResponse definition
type CreateNFVProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetDeviceDetailsByIPResponse is the getDeviceDetailsByIPResponse definition
type GetDeviceDetailsByIPResponse struct {
	ProvisionDetails GetDeviceDetailsByIPResponseProvisionDetails `json:"provisionDetails,omitempty"` //
}

// GetDeviceDetailsByIPResponseProvisionDetails is the getDeviceDetailsByIPResponseProvisionDetails definition
type GetDeviceDetailsByIPResponseProvisionDetails struct {
	BeginStep     string                                                  `json:"beginStep,omitempty"`     //
	Duration      string                                                  `json:"duration,omitempty"`      //
	EndTime       int                                                     `json:"endTime,omitempty"`       //
	StartTime     int                                                     `json:"startTime,omitempty"`     //
	Status        string                                                  `json:"status,omitempty"`        //
	StatusMessage string                                                  `json:"statusMessage,omitempty"` //
	TaskNodes     []GetDeviceDetailsByIPResponseProvisionDetailsTaskNodes `json:"taskNodes,omitempty"`     //
	Topology      string                                                  `json:"topology,omitempty"`      //
}

// GetDeviceDetailsByIPResponseProvisionDetailsTaskNodes is the getDeviceDetailsByIPResponseProvisionDetailsTaskNodes definition
type GetDeviceDetailsByIPResponseProvisionDetailsTaskNodes struct {
	CliTemplateUserMessageDTO string `json:"cliTemplateUserMessageDTO,omitempty"` //
	Duration                  string `json:"duration,omitempty"`                  //
	EndTime                   int    `json:"endTime,omitempty"`                   //
	ErrorPayload              string `json:"errorPayload,omitempty"`              //
	Name                      string `json:"name,omitempty"`                      //
	NextTask                  string `json:"nextTask,omitempty"`                  //
	ParentTask                string `json:"parentTask,omitempty"`                //
	Payload                   string `json:"payload,omitempty"`                   //
	ProvisionedNames          string `json:"provisionedNames,omitempty"`          //
	StartTime                 int    `json:"startTime,omitempty"`                 //
	Status                    string `json:"status,omitempty"`                    //
	StatusMessage             string `json:"statusMessage,omitempty"`             //
	StepRan                   string `json:"stepRan,omitempty"`                   //
	Target                    string `json:"target,omitempty"`                    //
}

// GetNFVProfileResponse is the getNFVProfileResponse definition
type GetNFVProfileResponse struct {
	Response []GetNFVProfileResponseResponse `json:"response,omitempty"` //
}

// GetNFVProfileResponseResponse is the getNFVProfileResponseResponse definition
type GetNFVProfileResponseResponse struct {
	Device      []GetNFVProfileResponseResponseDevice `json:"device,omitempty"`      //
	ID          string                                `json:"id,omitempty"`          //
	ProfileName string                                `json:"profileName,omitempty"` //
}

// GetNFVProfileResponseResponseDevice is the getNFVProfileResponseResponseDevice definition
type GetNFVProfileResponseResponseDevice struct {
	CustomNetworks                  []GetNFVProfileResponseResponseDeviceCustomNetworks         `json:"customNetworks,omitempty"`                  //
	CustomTemplate                  []GetNFVProfileResponseResponseDeviceCustomTemplate         `json:"customTemplate,omitempty"`                  //
	DeviceTag                       string                                                      `json:"deviceTag,omitempty"`                       //
	DeviceType                      string                                                      `json:"deviceType,omitempty"`                      //
	DirectInternetAccessForFirewall bool                                                        `json:"directInternetAccessForFirewall,omitempty"` //
	ServiceProviderProfile          []GetNFVProfileResponseResponseDeviceServiceProviderProfile `json:"serviceProviderProfile,omitempty"`          //
	Services                        []GetNFVProfileResponseResponseDeviceServices               `json:"services,omitempty"`                        //
	VLANForL2                       []GetNFVProfileResponseResponseDeviceVLANForL2              `json:"vlanForL2,omitempty"`                       //
}

// GetNFVProfileResponseResponseDeviceCustomNetworks is the getNFVProfileResponseResponseDeviceCustomNetworks definition
type GetNFVProfileResponseResponseDeviceCustomNetworks struct {
	ConnectionType    string                                                               `json:"connectionType,omitempty"`    //
	NetworkName       string                                                               `json:"networkName,omitempty"`       //
	ServicesToConnect []GetNFVProfileResponseResponseDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	VLANID            string                                                               `json:"vlanId,omitempty"`            //
	VLANMode          string                                                               `json:"vlanMode,omitempty"`          //
}

// GetNFVProfileResponseResponseDeviceCustomNetworksServicesToConnect is the getNFVProfileResponseResponseDeviceCustomNetworksServicesToConnect definition
type GetNFVProfileResponseResponseDeviceCustomNetworksServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` //
}

// GetNFVProfileResponseResponseDeviceCustomTemplate is the getNFVProfileResponseResponseDeviceCustomTemplate definition
type GetNFVProfileResponseResponseDeviceCustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   //
	Template     string `json:"template,omitempty"`     //
	TemplateType string `json:"templateType,omitempty"` //
}

// GetNFVProfileResponseResponseDeviceServiceProviderProfile is the getNFVProfileResponseResponseDeviceServiceProviderProfile definition
type GetNFVProfileResponseResponseDeviceServiceProviderProfile struct {
	Connect                    bool   `json:"connect,omitempty"`                    //
	ConnectDefaultGatewayOnWan bool   `json:"connectDefaultGatewayOnWan,omitempty"` //
	LinkType                   string `json:"linkType,omitempty"`                   //
	ServiceProvider            string `json:"serviceProvider,omitempty"`            //
}

// GetNFVProfileResponseResponseDeviceServices is the getNFVProfileResponseResponseDeviceServices definition
type GetNFVProfileResponseResponseDeviceServices struct {
	FirewallMode string                                                   `json:"firewallMode,omitempty"` //
	ImageName    string                                                   `json:"imageName,omitempty"`    //
	ProfileType  string                                                   `json:"profileType,omitempty"`  //
	ServiceName  string                                                   `json:"serviceName,omitempty"`  //
	ServiceType  string                                                   `json:"serviceType,omitempty"`  //
	VNicMapping  []GetNFVProfileResponseResponseDeviceServicesVNicMapping `json:"vNicMapping,omitempty"`  //
}

// GetNFVProfileResponseResponseDeviceServicesVNicMapping is the getNFVProfileResponseResponseDeviceServicesVNicMapping definition
type GetNFVProfileResponseResponseDeviceServicesVNicMapping struct {
	AssignIPAddressToNetwork bool   `json:"assignIpAddressToNetwork,omitempty"` //
	NetworkType              string `json:"networkType,omitempty"`              //
}

// GetNFVProfileResponseResponseDeviceVLANForL2 is the getNFVProfileResponseResponseDeviceVLANForL2 definition
type GetNFVProfileResponseResponseDeviceVLANForL2 struct {
	VLANDescription string `json:"vlanDescription,omitempty"` //
	VLANID          string `json:"vlanId,omitempty"`          //
	VLANType        string `json:"vlanType,omitempty"`        //
}

// NFVProvisioningDetailResponse is the nFVProvisioningDetailResponse definition
type NFVProvisioningDetailResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// ProvisionNFVResponse is the provisionNFVResponse definition
type ProvisionNFVResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateNFVProfileResponse is the updateNFVProfileResponse definition
type UpdateNFVProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateNFVProfile createNFVProfile
/* API to create network profile for different NFV topologies
 */
func (s *SiteDesignService) CreateNFVProfile(createNFVProfileRequest *CreateNFVProfileRequest) (*CreateNFVProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/nfv/network-profile"

	response, err := s.client.R().
		SetBody(createNFVProfileRequest).
		SetResult(&CreateNFVProfileResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createNFVProfile")
	}

	result := response.Result().(*CreateNFVProfileResponse)
	return result, response, err
}

// GetDeviceDetailsByIPQueryParams defines the query parameters for this request
type GetDeviceDetailsByIPQueryParams struct {
	DeviceIP string `url:"deviceIp,omitempty"` // Device to which the provisioning detail has to be retrieved
}

// GetDeviceDetailsByIP getDeviceDetailsByIP
/* Returns provisioning device information for the specified IP address.
@param deviceIP Device to which the provisioning detail has to be retrieved
*/
func (s *SiteDesignService) GetDeviceDetailsByIP(getDeviceDetailsByIPQueryParams *GetDeviceDetailsByIPQueryParams) (*GetDeviceDetailsByIPResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/nfv/provisioningDetail"

	queryString, _ := query.Values(getDeviceDetailsByIPQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceDetailsByIPResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getDeviceDetailsByIP")
	}

	result := response.Result().(*GetDeviceDetailsByIPResponse)
	return result, response, err
}

// GetNFVProfileQueryParams defines the query parameters for this request
type GetNFVProfileQueryParams struct {
	Offset string `url:"offset,omitempty"` // offset/starting row
	Limit  string `url:"limit,omitempty"`  // Number of profile to be retrieved
	Name   string `url:"name,omitempty"`   // Name of network profile to be retrieved
}

// GetNFVProfile getNFVProfile
/* API to get NFV network profile.
@param id ID of network profile to retrieve.
@param offset offset/starting row
@param limit Number of profile to be retrieved
@param name Name of network profile to be retrieved
*/
func (s *SiteDesignService) GetNFVProfile(id string, getNFVProfileQueryParams *GetNFVProfileQueryParams) (*GetNFVProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/nfv/network-profile/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getNFVProfileQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetNFVProfileResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getNFVProfile")
	}

	result := response.Result().(*GetNFVProfileResponse)
	return result, response, err
}

// NFVProvisioningDetail nFVProvisioningDetail
/* Checks the provisioning detail of an ENCS device including log information.
@param __runsync Enable this parameter to execute the API and return a response synchronously
@param __runsynctimeout During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
@param __persistbapioutput Persist bapi sync response
*/
func (s *SiteDesignService) NFVProvisioningDetail(nFVProvisioningDetailRequest *NFVProvisioningDetailRequest) (*NFVProvisioningDetailResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/nfv-provision-detail"

	response, err := s.client.R().
		SetBody(nFVProvisioningDetailRequest).
		SetResult(&NFVProvisioningDetailResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation nFVProvisioningDetail")
	}

	result := response.Result().(*NFVProvisioningDetailResponse)
	return result, response, err
}

// ProvisionNFV provisionNFV
/* Design and Provision single/multi NFV device with given site/area/building/floor .
@param __runsync Enable this parameter to execute the API and return a response synchronously
@param __timeout During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
@param __persistbapioutput Persist bapi sync response
*/
func (s *SiteDesignService) ProvisionNFV(provisionNFVRequest *ProvisionNFVRequest) (*ProvisionNFVResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/nfv"

	response, err := s.client.R().
		SetBody(provisionNFVRequest).
		SetResult(&ProvisionNFVResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation provisionNFV")
	}

	result := response.Result().(*ProvisionNFVResponse)
	return result, response, err
}

// UpdateNFVProfileQueryParams defines the query parameters for this request
type UpdateNFVProfileQueryParams struct {
	Name string `url:"name,omitempty"` // Name of the profile to be updated
}

// UpdateNFVProfile updateNFVProfile
/* API to update a NFV Network profile
@param id Id of the NFV profile to be updated
@param name Name of the profile to be updated
*/
func (s *SiteDesignService) UpdateNFVProfile(id string, updateNFVProfileQueryParams *UpdateNFVProfileQueryParams, updateNFVProfileRequest *UpdateNFVProfileRequest) (*UpdateNFVProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/nfv/network-profile/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(updateNFVProfileQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetBody(updateNFVProfileRequest).
		SetResult(&UpdateNFVProfileResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateNFVProfile")
	}

	result := response.Result().(*UpdateNFVProfileResponse)
	return result, response, err
}
