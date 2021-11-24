package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SiteDesignService service

type ProvisionNfvHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout           string `url:"__timeout,omitempty"`           //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type GetDeviceDetailsByIPQueryParams struct {
	DeviceIP string `url:"deviceIp,omitempty"` //Device to which the provisioning detail has to be retrieved
}
type NfvProvisioningDetailHeaderParams struct {
	Runsync           string `url:"__runsync,omitempty"`           //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Runsynctimeout    string `url:"__runsynctimeout,omitempty"`    //Expects type float64. During synchronous execution, this defines the maximum time to wait for a response, before the API execution is terminated
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. Persist bapi sync response
}
type UpdateNfvProfileQueryParams struct {
	Name string `url:"name,omitempty"` //Name of the profile to be updated
}
type GetNfvProfileQueryParams struct {
	Offset string `url:"offset,omitempty"` //offset/starting row
	Limit  string `url:"limit,omitempty"`  //Number of profile to be retrieved
	Name   string `url:"name,omitempty"`   //Name of network profile to be retrieved
}
type DeleteNfvProfileQueryParams struct {
	Name string `url:"name,omitempty"` //Nameof nfv network profile to delete.
}

type ResponseSiteDesignProvisionNfv struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSiteDesignGetDeviceDetailsByIP struct {
	ProvisionDetails *ResponseSiteDesignGetDeviceDetailsByIPProvisionDetails `json:"provisionDetails,omitempty"` //
}
type ResponseSiteDesignGetDeviceDetailsByIPProvisionDetails struct {
	StartTime     string                                                             `json:"startTime,omitempty"`     // Start Time
	EndTime       string                                                             `json:"endTime,omitempty"`       // End Time
	Duration      string                                                             `json:"duration,omitempty"`      // Duration
	StatusMessage string                                                             `json:"statusMessage,omitempty"` // Status Message
	Status        string                                                             `json:"status,omitempty"`        // Status
	TaskNodes     *[]ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodes `json:"taskNodes,omitempty"`     //
	Topology      string                                                             `json:"topology,omitempty"`      // Topology
	BeginStep     string                                                             `json:"beginStep,omitempty"`     // Begin Step
}
type ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodes struct {
	StartTime                 string                                                                                    `json:"startTime,omitempty"`                 // Start Time
	EndTime                   string                                                                                    `json:"endTime,omitempty"`                   // End Time
	Duration                  string                                                                                    `json:"duration,omitempty"`                  // Duration
	Status                    string                                                                                    `json:"status,omitempty"`                    // Status
	NextTask                  string                                                                                    `json:"nextTask,omitempty"`                  // Next Task
	Name                      string                                                                                    `json:"name,omitempty"`                      // Name
	Target                    string                                                                                    `json:"target,omitempty"`                    // Target
	StatusMessage             string                                                                                    `json:"statusMessage,omitempty"`             // Status Message
	Payload                   string                                                                                    `json:"payload,omitempty"`                   // Payload
	ProvisionedNames          *ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesProvisionedNames          `json:"provisionedNames,omitempty"`          // Provisioned Names
	ErrorPayload              *ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesErrorPayload              `json:"errorPayload,omitempty"`              // Error Payload
	ParentTask                *ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesParentTask                `json:"parentTask,omitempty"`                // Parent Task
	CliTemplateUserMessageDTO *ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesCliTemplateUserMessageDTO `json:"cliTemplateUserMessageDTO,omitempty"` // Cli Template User Message D T O
	StepRan                   string                                                                                    `json:"stepRan,omitempty"`                   // Step Ran
}
type ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesProvisionedNames interface{}
type ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesErrorPayload interface{}
type ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesParentTask interface{}
type ResponseSiteDesignGetDeviceDetailsByIPProvisionDetailsTaskNodesCliTemplateUserMessageDTO interface{}
type ResponseSiteDesignNfvProvisioningDetail struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSiteDesignCreateNfvProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSiteDesignUpdateNfvProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSiteDesignGetNfvProfile struct {
	Response *[]ResponseSiteDesignGetNfvProfileResponse `json:"response,omitempty"` //
}
type ResponseSiteDesignGetNfvProfileResponse struct {
	ProfileName string                                           `json:"profileName,omitempty"` // Name of the profile to create NFV profile( eg: Nfvis_profile)
	ID          string                                           `json:"id,omitempty"`          // Id of nfv created nfv profile
	Device      *[]ResponseSiteDesignGetNfvProfileResponseDevice `json:"device,omitempty"`      //
}
type ResponseSiteDesignGetNfvProfileResponseDevice struct {
	DeviceType                      string                                                                 `json:"deviceType,omitempty"`                      // Name of the device used in creating nfv profile(eg: Cisco 5400 Enterprise Network Compute System)
	DeviceTag                       string                                                                 `json:"deviceTag,omitempty"`                       // Device Tag name(eg: dev1)
	ServiceProviderProfile          *[]ResponseSiteDesignGetNfvProfileResponseDeviceServiceProviderProfile `json:"serviceProviderProfile,omitempty"`          //
	DirectInternetAccessForFirewall *bool                                                                  `json:"directInternetAccessForFirewall,omitempty"` // Direct internet access value should be boolean (eg: false)
	Services                        *[]ResponseSiteDesignGetNfvProfileResponseDeviceServices               `json:"services,omitempty"`                        //
	CustomNetworks                  *[]ResponseSiteDesignGetNfvProfileResponseDeviceCustomNetworks         `json:"customNetworks,omitempty"`                  //
	VLANForL2                       *[]ResponseSiteDesignGetNfvProfileResponseDeviceVLANForL2              `json:"vlanForL2,omitempty"`                       //
	CustomTemplate                  *[]ResponseSiteDesignGetNfvProfileResponseDeviceCustomTemplate         `json:"customTemplate,omitempty"`                  //
}
type ResponseSiteDesignGetNfvProfileResponseDeviceServiceProviderProfile struct {
	LinkType                   string `json:"linkType,omitempty"`                   // Name of connection type(eg: GigabitEthernet)
	Connect                    *bool  `json:"connect,omitempty"`                    // Connection of service provider and device value should be boolean (eg: true)
	ConnectDefaultGatewayOnWan *bool  `json:"connectDefaultGatewayOnWan,omitempty"` // Default gateway connect value as boolean (eg: true)
	ServiceProvider            string `json:"serviceProvider,omitempty"`            // Name of the service provider(eg: Airtel)
}
type ResponseSiteDesignGetNfvProfileResponseDeviceServices struct {
	ServiceType  string                                                              `json:"serviceType,omitempty"`  // Service type (eg: ISRV)
	ProfileType  string                                                              `json:"profileType,omitempty"`  // Profile type of service (eg: ISRv-mini)
	ServiceName  string                                                              `json:"serviceName,omitempty"`  // Name of service (eg: router-1)
	ImageName    string                                                              `json:"imageName,omitempty"`    // Service image name (eg: isrv-universalk9.16.12.01a.tar.gz)
	VNicMapping  *[]ResponseSiteDesignGetNfvProfileResponseDeviceServicesVnicMapping `json:"vNicMapping,omitempty"`  //
	FirewallMode string                                                              `json:"firewallMode,omitempty"` // Mode of firewall (eg: routed, transparent)
}
type ResponseSiteDesignGetNfvProfileResponseDeviceServicesVnicMapping struct {
	NetworkType              string `json:"networkType,omitempty"`              // Type of connection (eg:  wan, lan or internal)
	AssignIPAddressToNetwork *bool  `json:"assignIpAddressToNetwork,omitempty"` // Assign ip address to network (eg: true)
}
type ResponseSiteDesignGetNfvProfileResponseDeviceCustomNetworks struct {
	NetworkName       string                                                                          `json:"networkName,omitempty"`       // name of custom network (eg: cust-1)
	ServicesToConnect *[]ResponseSiteDesignGetNfvProfileResponseDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	ConnectionType    string                                                                          `json:"connectionType,omitempty"`    // Type of network connection from custom network (eg: lan)
	VLANMode          string                                                                          `json:"vlanMode,omitempty"`          // Vlan network mode (eg Access or Trunk)
	VLANID            string                                                                          `json:"vlanId,omitempty"`            // Vlan id for the custom network(eg: 4000)
}
type ResponseSiteDesignGetNfvProfileResponseDeviceCustomNetworksServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` // Name of service to be connected to the custom network (eg: router-1)
}
type ResponseSiteDesignGetNfvProfileResponseDeviceVLANForL2 struct {
	VLANType        string `json:"vlanType,omitempty"`        // Vlan type(eg. Access or Trunk)
	VLANID          string `json:"vlanId,omitempty"`          // Vlan id(eg.4018)
	VLANDescription string `json:"vlanDescription,omitempty"` // Vlan description(eg. Access 4018)
}
type ResponseSiteDesignGetNfvProfileResponseDeviceCustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   // Type of the device(eg: Cisco 5400 Enterprise Network Compute System)
	Template     string `json:"template,omitempty"`     // Name of the template(eg NFVIS template)
	TemplateType string `json:"templateType,omitempty"` // Name of the template to which template is associated (eg: Cloud DayN Templates)
}
type ResponseSiteDesignDeleteNfvProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseSiteDesignCreateFloormap interface{}
type ResponseSiteDesignListFloormaps struct {
	Total *int                                    `json:"total,omitempty"` // Total
	Items *[]ResponseSiteDesignListFloormapsItems `json:"items,omitempty"` //
}
type ResponseSiteDesignListFloormapsItems struct {
	ParentGroupUUID      string                                         `json:"parentGroupUuid,omitempty"`      // Parent Group Uuid
	SiteName             string                                         `json:"siteName,omitempty"`             // Site Name
	ImageInfo            *ResponseSiteDesignListFloormapsItemsImageInfo `json:"imageInfo,omitempty"`            //
	Type                 *int                                           `json:"type,omitempty"`                 //
	Contact              *ResponseSiteDesignListFloormapsItemsContact   `json:"contact,omitempty"`              // Contact
	WirelessClientsCount string                                         `json:"wirelessClientsCount,omitempty"` // Wireless Clients Count
	ID                   *int                                           `json:"id,omitempty"`                   // Id
	ApCount              *int                                           `json:"apCount,omitempty"`              //
	BuildingName         string                                         `json:"buildingName,omitempty"`         // Building Name
	CriticalRadioCount   string                                         `json:"criticalRadioCount,omitempty"`   // Critical Radio Count
	InstanceUUID         string                                         `json:"instanceUuid,omitempty"`         // Instance Uuid
	Name                 string                                         `json:"name,omitempty"`                 // Name
	Location             *ResponseSiteDesignListFloormapsItemsLocation  `json:"location,omitempty"`             //
	Geometry             *ResponseSiteDesignListFloormapsItemsGeometry  `json:"geometry,omitempty"`             //
	GroupInstanceUUID    string                                         `json:"groupInstanceUuid,omitempty"`    // Group Instance Uuid
	Metrics              *ResponseSiteDesignListFloormapsItemsMetrics   `json:"metrics,omitempty"`              // Metrics
	Items                []string                                       `json:"items,omitempty"`                //
	FloorIndex           *float64                                       `json:"floorIndex,omitempty"`           // Floor Index
	MaintainAspectRatio  *bool                                          `json:"maintainAspectRatio,omitempty"`  //
}
type ResponseSiteDesignListFloormapsItemsImageInfo struct {
	Image                *ResponseSiteDesignListFloormapsItemsImageInfoImage                `json:"image,omitempty"`                // Image
	Thumbnail            *ResponseSiteDesignListFloormapsItemsImageInfoThumbnail            `json:"thumbnail,omitempty"`            // Thumbnail
	GeneratedRasterImage *ResponseSiteDesignListFloormapsItemsImageInfoGeneratedRasterImage `json:"generatedRasterImage,omitempty"` // Generated Raster Image
	EnteredImageName     *ResponseSiteDesignListFloormapsItemsImageInfoEnteredImageName     `json:"enteredImageName,omitempty"`     // Entered Image Name
	Features             *ResponseSiteDesignListFloormapsItemsImageInfoFeatures             `json:"features,omitempty"`             // Features
	Shapes               *ResponseSiteDesignListFloormapsItemsImageInfoShapes               `json:"shapes,omitempty"`               // Shapes
	IsCadFile            *bool                                                              `json:"isCadFile,omitempty"`            // Is Cad File
	Categories           *ResponseSiteDesignListFloormapsItemsImageInfoCategories           `json:"categories,omitempty"`           // Categories
}
type ResponseSiteDesignListFloormapsItemsImageInfoImage interface{}
type ResponseSiteDesignListFloormapsItemsImageInfoThumbnail interface{}
type ResponseSiteDesignListFloormapsItemsImageInfoGeneratedRasterImage interface{}
type ResponseSiteDesignListFloormapsItemsImageInfoEnteredImageName interface{}
type ResponseSiteDesignListFloormapsItemsImageInfoFeatures interface{}
type ResponseSiteDesignListFloormapsItemsImageInfoShapes interface{}
type ResponseSiteDesignListFloormapsItemsImageInfoCategories interface{}
type ResponseSiteDesignListFloormapsItemsContact interface{}
type ResponseSiteDesignListFloormapsItemsLocation struct {
	Country string   `json:"country,omitempty"` // Country
	Address string   `json:"address,omitempty"` // Address
	Lon     *float64 `json:"lon,omitempty"`     // Lon
	Lat     *float64 `json:"lat,omitempty"`     // Lat
	Height  *float64 `json:"height,omitempty"`  // Height
}
type ResponseSiteDesignListFloormapsItemsGeometry struct {
	Length  *float64                                          `json:"length,omitempty"`  // Length
	Type    *ResponseSiteDesignListFloormapsItemsGeometryType `json:"type,omitempty"`    // Type
	OffsetX *float64                                          `json:"offsetX,omitempty"` // OffsetX
	OffsetY *float64                                          `json:"offsetY,omitempty"` // OffsetY
	Width   *float64                                          `json:"width,omitempty"`   // Width
	Height  *float64                                          `json:"height,omitempty"`  // Height
}
type ResponseSiteDesignListFloormapsItemsGeometryType interface{}
type ResponseSiteDesignListFloormapsItemsMetrics interface{}
type ResponseSiteDesignDeleteFloormap interface{}
type ResponseSiteDesignUpdateFloormap interface{}
type ResponseSiteDesignListSpecifiedFloormaps interface{}
type RequestSiteDesignProvisionNfv struct {
	SiteProfile  *[]RequestSiteDesignProvisionNfvSiteProfile  `json:"siteProfile,omitempty"`  //
	Provisioning *[]RequestSiteDesignProvisionNfvProvisioning `json:"provisioning,omitempty"` //
}
type RequestSiteDesignProvisionNfvSiteProfile struct {
	SiteProfileName string                                            `json:"siteProfileName,omitempty"` // Name of the profile to create site profile profile( eg: profile-1)
	Device          *[]RequestSiteDesignProvisionNfvSiteProfileDevice `json:"device,omitempty"`          //
}
type RequestSiteDesignProvisionNfvSiteProfileDevice struct {
	DeviceType       string                                                            `json:"deviceType,omitempty"`       // Name of the device used in creating nfv profile(eg: ENCS5400)
	TagName          string                                                            `json:"tagName,omitempty"`          // Device Tag name(eg: dev1)
	ServiceProviders *[]RequestSiteDesignProvisionNfvSiteProfileDeviceServiceProviders `json:"serviceProviders,omitempty"` //
	Dia              *bool                                                             `json:"dia,omitempty"`              // Direct internet access value should be boolean (eg: false)
	Services         *[]RequestSiteDesignProvisionNfvSiteProfileDeviceServices         `json:"services,omitempty"`         //
	CustomServices   *[]RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServices   `json:"customServices,omitempty"`   //
	CustomNetworks   *[]RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworks   `json:"customNetworks,omitempty"`   //
	VLAN             *[]RequestSiteDesignProvisionNfvSiteProfileDeviceVLAN             `json:"vlan,omitempty"`             //
	CustomTemplate   *[]RequestSiteDesignProvisionNfvSiteProfileDeviceCustomTemplate   `json:"customTemplate,omitempty"`   //
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceServiceProviders struct {
	ServiceProvider string `json:"serviceProvider,omitempty"` // Name of the service provider(eg: Airtel)
	LinkType        string `json:"linkType,omitempty"`        // Name of connection type(eg: GigabitEthernet)
	Connect         *bool  `json:"connect,omitempty"`         // Connection of service provider and device value should be boolean (eg: true)
	DefaultGateway  *bool  `json:"defaultGateway,omitempty"`  // Default gateway connect value as boolean (eg: true)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceServices struct {
	Type      string                                                          `json:"type,omitempty"`      // Service type (eg: ISRV)
	Profile   string                                                          `json:"profile,omitempty"`   // Profile type of service (eg: ISRv-mini)
	Mode      string                                                          `json:"mode,omitempty"`      // Mode of firewall (eg: routed, transparent)
	Name      string                                                          `json:"name,omitempty"`      // Name of the service (eg: isrv)
	ImageName string                                                          `json:"imageName,omitempty"` // Name of image (eg: isrv-universalk9.16.06.02.tar.gz)
	Topology  *RequestSiteDesignProvisionNfvSiteProfileDeviceServicesTopology `json:"topology,omitempty"`  //
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceServicesTopology struct {
	Type     string `json:"type,omitempty"`     // Type of connection (eg:  wan, lan or internal)
	Name     string `json:"name,omitempty"`     // Name of connection (eg: wan-net)
	AssignIP string `json:"assignIp,omitempty"` // Assign ip address to network (eg: true)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServices struct {
	Name            string                                                                `json:"name,omitempty"`            // Name of custom service (eg: LINUX-1)
	ApplicationType string                                                                `json:"applicationType,omitempty"` // Application type of custom service (eg: LINUX)
	Profile         string                                                                `json:"profile,omitempty"`         // Profile type of service (eg: rhel7-medium)
	Topology        *RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServicesTopology `json:"topology,omitempty"`        //
	ImageName       string                                                                `json:"imageName,omitempty"`       // Image name of custom service (eg: redhat7.tar.gz.tar.gz)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceCustomServicesTopology struct {
	Type     string `json:"type,omitempty"`     // Type of connection from custom service (eg:  wan, lan or internal)
	Name     string `json:"name,omitempty"`     // Name of connection from custom service(eg: wan-net)
	AssignIP string `json:"assignIp,omitempty"` // Assign ip to network (eg: true)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworks struct {
	Name              string                                                                           `json:"name,omitempty"`              // Name of custom network (eg: cust-1)
	ServicesToConnect *[]RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	ConnectionType    string                                                                           `json:"connectionType,omitempty"`    // Type of network connection from custom network (eg: lan)
	NetworkMode       string                                                                           `json:"networkMode,omitempty"`       // Network mode (eg Access or Trunk)
	VLAN              string                                                                           `json:"vlan,omitempty"`              // Vlan id for the custom network(eg: 4000)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceCustomNetworksServicesToConnect struct {
	Service string `json:"service,omitempty"` // Name of service to be connected to the custom network (eg: router-1)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceVLAN struct {
	Type string `json:"type,omitempty"` // Vlan type(eg. Access or Trunk)
	ID   string `json:"id,omitempty"`   // Vlan id(eg.4018)
}
type RequestSiteDesignProvisionNfvSiteProfileDeviceCustomTemplate struct {
	DeviceType string `json:"deviceType,omitempty"` // Type of the device(eg: NFVIS)
	Template   string `json:"template,omitempty"`   // Name of the template(eg NFVIS template)
}
type RequestSiteDesignProvisionNfvProvisioning struct {
	Site   *RequestSiteDesignProvisionNfvProvisioningSite     `json:"site,omitempty"`   //
	Device *[]RequestSiteDesignProvisionNfvProvisioningDevice `json:"device,omitempty"` //
}
type RequestSiteDesignProvisionNfvProvisioningSite struct {
	SiteProfileName string                                                 `json:"siteProfileName,omitempty"` // Name of site profile to be provision with device
	Area            *RequestSiteDesignProvisionNfvProvisioningSiteArea     `json:"area,omitempty"`            //
	Building        *RequestSiteDesignProvisionNfvProvisioningSiteBuilding `json:"building,omitempty"`        //
	Floor           *RequestSiteDesignProvisionNfvProvisioningSiteFloor    `json:"floor,omitempty"`           //
}
type RequestSiteDesignProvisionNfvProvisioningSiteArea struct {
	Name       string `json:"name,omitempty"`       // Name of the area (eg: Area1)
	ParentName string `json:"parentName,omitempty"` // Parent name of the area to be created
}
type RequestSiteDesignProvisionNfvProvisioningSiteBuilding struct {
	Name       string   `json:"name,omitempty"`       // Name of the building (eg: building1)
	Address    string   `json:"address,omitempty"`    // Address of the building to be created
	Latitude   *float64 `json:"latitude,omitempty"`   // Latitude coordinate of the building (eg:37.338)
	Longitude  *float64 `json:"longitude,omitempty"`  // Longitude coordinate of the building (eg:-121.832)
	ParentName string   `json:"parentName,omitempty"` // Address of the building to be created
}
type RequestSiteDesignProvisionNfvProvisioningSiteFloor struct {
	Name       string   `json:"name,omitempty"`       // Name of the floor (eg:floor-1)
	ParentName string   `json:"parentName,omitempty"` // Parent name of the floor to be created
	RfModel    string   `json:"rfModel,omitempty"`    // Type of floor (eg: Cubes And Walled Offices)
	Width      *float64 `json:"width,omitempty"`      // Width of the floor (eg:100)
	Length     *float64 `json:"length,omitempty"`     // Length of the floor (eg: 100)
	Height     *float64 `json:"height,omitempty"`     // Height of the floor (eg: 15)
}
type RequestSiteDesignProvisionNfvProvisioningDevice struct {
	IP                 string                                                             `json:"ip,omitempty"`                 // IP address of the device (eg: 172.20.126.90)
	DeviceSerialNumber string                                                             `json:"deviceSerialNumber,omitempty"` // Serial number of device (eg: FGL210710QY)
	TagName            string                                                             `json:"tagName,omitempty"`            // Name of device tag (eg: dev1)
	ServiceProviders   *[]RequestSiteDesignProvisionNfvProvisioningDeviceServiceProviders `json:"serviceProviders,omitempty"`   //
	Services           *[]RequestSiteDesignProvisionNfvProvisioningDeviceServices         `json:"services,omitempty"`           //
	VLAN               *[]RequestSiteDesignProvisionNfvProvisioningDeviceVLAN             `json:"vlan,omitempty"`               //
	SubPools           *[]RequestSiteDesignProvisionNfvProvisioningDeviceSubPools         `json:"subPools,omitempty"`           //
	CustomNetworks     *[]RequestSiteDesignProvisionNfvProvisioningDeviceCustomNetworks   `json:"customNetworks,omitempty"`     //
	TemplateParam      *RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParam      `json:"templateParam,omitempty"`      //
}
type RequestSiteDesignProvisionNfvProvisioningDeviceServiceProviders struct {
	ServiceProvider string                                                                       `json:"serviceProvider,omitempty"` // Name of the service provider (eg: Airtel)
	WanInterface    *RequestSiteDesignProvisionNfvProvisioningDeviceServiceProvidersWanInterface `json:"wanInterface,omitempty"`    //
}
type RequestSiteDesignProvisionNfvProvisioningDeviceServiceProvidersWanInterface struct {
	IPAddress     string `json:"ipAddress,omitempty"`     // IP address (eg: 175.175.190.205)
	InterfaceName string `json:"interfaceName,omitempty"` // Name of the interface (eg: GE0-0)
	Subnetmask    string `json:"subnetmask,omitempty"`    // Subnet mask (eg: 255.255.255.0)
	Bandwidth     string `json:"bandwidth,omitempty"`     // Bandwidth limit (eg: 100)
	Gateway       string `json:"gateway,omitempty"`       // Gateway (eg: 175.175.190.1)
}
type RequestSiteDesignProvisionNfvProvisioningDeviceServices struct {
	Type                   string `json:"type,omitempty"`                   // Type of service (eg: ISR)
	Mode                   string `json:"mode,omitempty"`                   // Mode of firewall (eg: transparent)
	SystemIP               string `json:"systemIp,omitempty"`               // System IP
	CentralManagerIP       string `json:"centralManagerIP,omitempty"`       // WAAS Package needs to be installed to populate Central Manager IP automatically.
	CentralRegistrationKey string `json:"centralRegistrationKey,omitempty"` // Central registration key
	CommonKey              string `json:"commonKey,omitempty"`              // Common key
	AdminPasswordHash      string `json:"adminPasswordHash,omitempty"`      // Admin password hash
	Disk                   string `json:"disk,omitempty"`                   // Name of disk type (eg: internal)
}
type RequestSiteDesignProvisionNfvProvisioningDeviceVLAN struct {
	Type       string `json:"type,omitempty"`       // Vlan type(eg. Access or Trunk)
	ID         string `json:"id,omitempty"`         // Vlan id(e: .4018)
	Interfaces string `json:"interfaces,omitempty"` // Interface (eg: GigabitEathernet1/0)
	Network    string `json:"network,omitempty"`    // Network name to connect (eg: lan-net)
}
type RequestSiteDesignProvisionNfvProvisioningDeviceSubPools struct {
	Type           string `json:"type,omitempty"`           // Tyep of ip sub pool (eg: Lan)
	Name           string `json:"name,omitempty"`           // Name of the ip sub pool (eg; Lan-65)
	IPSubnet       string `json:"ipSubnet,omitempty"`       // IP pool cidir (eg: 175.175.140.0)
	Gateway        string `json:"gateway,omitempty"`        // IP address for gate way (eg: 175.175.140.1)
	ParentPoolName string `json:"parentPoolName,omitempty"` // Name of parent pool (global pool name)
}
type RequestSiteDesignProvisionNfvProvisioningDeviceCustomNetworks struct {
	Name          string `json:"name,omitempty"`          // Name of custom network (eg: cust-1)
	Port          string `json:"port,omitempty"`          // Port for custom network (eg: 443)
	IPAddressPool string `json:"ipAddressPool,omitempty"` // IP address pool of sub pool (eg: 175.175.140.1)
}
type RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParam struct {
	Nfvis *RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamNfvis `json:"nfvis,omitempty"` //
	Asav  *RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamAsav  `json:"asav,omitempty"`  //
}
type RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamNfvis struct {
	Var1 string `json:"var1,omitempty"` // Variable for nfvis template (eg: "test":"Hello nfvis")
}
type RequestSiteDesignProvisionNfvProvisioningDeviceTemplateParamAsav struct {
	Var1 string `json:"var1,omitempty"` // Variable for asav template (eg: "test":"Hello asav")
}
type RequestSiteDesignNfvProvisioningDetail struct {
	DeviceIP string `json:"device_ip,omitempty"` // Device Ip
}
type RequestSiteDesignCreateNfvProfile struct {
	ProfileName string                                     `json:"profileName,omitempty"` // Name of the profile to create NFV profile
	Device      *[]RequestSiteDesignCreateNfvProfileDevice `json:"device,omitempty"`      //
}
type RequestSiteDesignCreateNfvProfileDevice struct {
	DeviceType                      string                                                           `json:"deviceType,omitempty"`                      // Name of the device used in creating nfv profile
	DeviceTag                       string                                                           `json:"deviceTag,omitempty"`                       // Device Tag name(eg: dev1)
	ServiceProviderProfile          *[]RequestSiteDesignCreateNfvProfileDeviceServiceProviderProfile `json:"serviceProviderProfile,omitempty"`          //
	DirectInternetAccessForFirewall *bool                                                            `json:"directInternetAccessForFirewall,omitempty"` // Direct internet access value should be boolean (eg: false or true)
	Services                        *[]RequestSiteDesignCreateNfvProfileDeviceServices               `json:"services,omitempty"`                        //
	CustomNetworks                  *[]RequestSiteDesignCreateNfvProfileDeviceCustomNetworks         `json:"customNetworks,omitempty"`                  //
	VLANForL2                       *[]RequestSiteDesignCreateNfvProfileDeviceVLANForL2              `json:"vlanForL2,omitempty"`                       //
	CustomTemplate                  *[]RequestSiteDesignCreateNfvProfileDeviceCustomTemplate         `json:"customTemplate,omitempty"`                  //
}
type RequestSiteDesignCreateNfvProfileDeviceServiceProviderProfile struct {
	ServiceProvider            string `json:"serviceProvider,omitempty"`            // Name of the service provider(eg: Airtel)
	LinkType                   string `json:"linkType,omitempty"`                   // Name of connection type(eg: GigabitEthernet)
	Connect                    *bool  `json:"connect,omitempty"`                    // Connection of service provider and device value should be boolean (eg: true)
	ConnectDefaultGatewayOnWan *bool  `json:"connectDefaultGatewayOnWan,omitempty"` // Connect default gateway connect value as boolean (eg: true)
}
type RequestSiteDesignCreateNfvProfileDeviceServices struct {
	ServiceType  string                                                        `json:"serviceType,omitempty"`  // Service type (eg: ISRV)
	ProfileType  string                                                        `json:"profileType,omitempty"`  // Profile type of service (eg: ISRv-mini)
	ServiceName  string                                                        `json:"serviceName,omitempty"`  // Name of the service (eg: Router-1)
	ImageName    string                                                        `json:"imageName,omitempty"`    // Service image name (eg: isrv-universalk9.16.12.01a.tar.gz)
	VNicMapping  *[]RequestSiteDesignCreateNfvProfileDeviceServicesVnicMapping `json:"vNicMapping,omitempty"`  //
	FirewallMode string                                                        `json:"firewallMode,omitempty"` // Firewall mode details example (routed, transparent)
}
type RequestSiteDesignCreateNfvProfileDeviceServicesVnicMapping struct {
	NetworkType              string `json:"networkType,omitempty"`              // Type of connection (eg:  wan, lan or internal)
	AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` // Assign ip address to network (eg: true or false)
}
type RequestSiteDesignCreateNfvProfileDeviceCustomNetworks struct {
	NetworkName       string                                                                    `json:"networkName,omitempty"`       // Name of custom network (eg: cust-1)
	ServicesToConnect *[]RequestSiteDesignCreateNfvProfileDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	ConnectionType    string                                                                    `json:"connectionType,omitempty"`    // Type of network connection from custom network (eg: lan)
	VLANMode          string                                                                    `json:"vlanMode,omitempty"`          // Network mode (eg Access or Trunk)
	VLANID            *float64                                                                  `json:"vlanId,omitempty"`            // Vlan id for the custom network(eg: 4000)
}
type RequestSiteDesignCreateNfvProfileDeviceCustomNetworksServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` // Name of service to be connected to the custom network (eg: router-1)
}
type RequestSiteDesignCreateNfvProfileDeviceVLANForL2 struct {
	VLANType        string   `json:"vlanType,omitempty"`        // Vlan type(eg: Access or Trunk)
	VLANID          *float64 `json:"vlanId,omitempty"`          // Vlan id (eg: 4018)
	VLANDescription string   `json:"vlanDescription,omitempty"` // Vlan description(eg: Access 4018)
}
type RequestSiteDesignCreateNfvProfileDeviceCustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   // Type of the device(eg: Cisco 5400 Enterprise Network Compute System)
	Template     string `json:"template,omitempty"`     // Name of the template(eg NFVIS template)
	TemplateType string `json:"templateType,omitempty"` // Name of the template type to which template is associated (eg: Cloud DayN Templates)
}
type RequestSiteDesignUpdateNfvProfile struct {
	Device *[]RequestSiteDesignUpdateNfvProfileDevice `json:"device,omitempty"` //
}
type RequestSiteDesignUpdateNfvProfileDevice struct {
	DeviceTag                       string                                                   `json:"deviceTag,omitempty"`                       // Device Tag name(eg: dev1)
	DirectInternetAccessForFirewall *bool                                                    `json:"directInternetAccessForFirewall,omitempty"` // Direct internet access value should be boolean (eg: false)
	Services                        *[]RequestSiteDesignUpdateNfvProfileDeviceServices       `json:"services,omitempty"`                        //
	CustomNetworks                  *[]RequestSiteDesignUpdateNfvProfileDeviceCustomNetworks `json:"customNetworks,omitempty"`                  //
	VLANForL2                       *[]RequestSiteDesignUpdateNfvProfileDeviceVLANForL2      `json:"vlanForL2,omitempty"`                       //
	CustomTemplate                  *[]RequestSiteDesignUpdateNfvProfileDeviceCustomTemplate `json:"customTemplate,omitempty"`                  //
	CurrentDeviceTag                string                                                   `json:"currentDeviceTag,omitempty"`                // Existing device tag name saved in the nfv profiles (eg: dev1)
}
type RequestSiteDesignUpdateNfvProfileDeviceServices struct {
	ServiceType  string                                                        `json:"serviceType,omitempty"`  // Service type (eg: ISRV)
	ProfileType  string                                                        `json:"profileType,omitempty"`  // Profile type of service (eg: ISRv-mini)
	ServiceName  string                                                        `json:"serviceName,omitempty"`  // Name of the service (eg: Router-1)
	ImageName    string                                                        `json:"imageName,omitempty"`    // Service image name (eg: isrv-universalk9.16.12.01a.tar.gz)
	VNicMapping  *[]RequestSiteDesignUpdateNfvProfileDeviceServicesVnicMapping `json:"vNicMapping,omitempty"`  //
	FirewallMode string                                                        `json:"firewallMode,omitempty"` // Mode of firewall (eg: routed, transparent)
}
type RequestSiteDesignUpdateNfvProfileDeviceServicesVnicMapping struct {
	NetworkType              string `json:"networkType,omitempty"`              // Type of connection (eg:  wan, lan or internal)
	AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` // Assign ip address to network (eg: true or false)
}
type RequestSiteDesignUpdateNfvProfileDeviceCustomNetworks struct {
	NetworkName       string                                                                    `json:"networkName,omitempty"`       // Name of custom network (eg: cust-1)
	ServicesToConnect *[]RequestSiteDesignUpdateNfvProfileDeviceCustomNetworksServicesToConnect `json:"servicesToConnect,omitempty"` //
	ConnectionType    string                                                                    `json:"connectionType,omitempty"`    // Type of network connection from custom network (eg: lan)
	VLANMode          string                                                                    `json:"vlanMode,omitempty"`          // Vlan network mode (eg Access or Trunk)
	VLANID            *float64                                                                  `json:"vlanId,omitempty"`            // Vlan id for the custom network(eg: 4000)
}
type RequestSiteDesignUpdateNfvProfileDeviceCustomNetworksServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` // Name of service to be connected to the custom network (eg: router-1)
}
type RequestSiteDesignUpdateNfvProfileDeviceVLANForL2 struct {
	VLANType        string   `json:"vlanType,omitempty"`        // Vlan type(eg. Access or Trunk)
	VLANID          *float64 `json:"vlanId,omitempty"`          // Vlan id(eg.4018)
	VLANDescription string   `json:"vlanDescription,omitempty"` // Vlan description(eg. Access 4018)
}
type RequestSiteDesignUpdateNfvProfileDeviceCustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   // Type of the device(eg: Cisco 5400 Enterprise Network Compute System)
	Template     string `json:"template,omitempty"`     // Name of the template(eg NFVIS template)
	TemplateType string `json:"templateType,omitempty"` // Name of the project to which template is associated (eg: Cloud DayN Templates)
}
type RequestSiteDesignCreateFloormap interface{}
type RequestSiteDesignUpdateFloormap interface{}

//GetDeviceDetailsByIP Get Device details by IP - 9cb2-cb3f-494a-824f
/* Returns provisioning device information for the specified IP address.


@param GetDeviceDetailsByIPQueryParams Filtering parameter
*/
func (s *SiteDesignService) GetDeviceDetailsByIP(GetDeviceDetailsByIPQueryParams *GetDeviceDetailsByIPQueryParams) (*ResponseSiteDesignGetDeviceDetailsByIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/nfv/provisioningDetail"

	queryString, _ := query.Values(GetDeviceDetailsByIPQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetDeviceDetailsByIP{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceDetailsByIp")
	}

	result := response.Result().(*ResponseSiteDesignGetDeviceDetailsByIP)
	return result, response, err

}

//GetNfvProfile Get NFV Profile - 1eb1-9887-457b-9616
/* API to get NFV network profile.


@param id id path parameter. ID of network profile to retrieve.

@param GetNFVProfileQueryParams Filtering parameter
*/
func (s *SiteDesignService) GetNfvProfile(id string, GetNFVProfileQueryParams *GetNfvProfileQueryParams) (*ResponseSiteDesignGetNfvProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/nfv/network-profile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetNFVProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignGetNfvProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNfvProfile")
	}

	result := response.Result().(*ResponseSiteDesignGetNfvProfile)
	return result, response, err

}

//ListFloormaps List Floormaps - d78f-8888-4b9a-a5fe
/* List all floor maps


 */
func (s *SiteDesignService) ListFloormaps() (*ResponseSiteDesignListFloormaps, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/floormap/all"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSiteDesignListFloormaps{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ListFloormaps")
	}

	result := response.Result().(*ResponseSiteDesignListFloormaps)
	return result, response, err

}

//ListSpecifiedFloormaps List Specified Floormap(s) - 5fb3-1896-43f8-9acf
/* List specified floor map(s)


@param floorID floorId path parameter. Group Id of the specified floormap

*/
func (s *SiteDesignService) ListSpecifiedFloormaps(floorID string) (*ResponseSiteDesignListSpecifiedFloormaps, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/floormap/{floorId}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		// SetResult(&ResponseSiteDesignListSpecifiedFloormaps{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ListSpecifiedFloormaps")
	}

	result := response.Result().(ResponseSiteDesignListSpecifiedFloormaps)

	return &result, response, err

}

//ProvisionNfv Provision NFV - 6f9c-da9a-4658-84b4
/* Design and Provision single/multi NFV device with given site/area/building/floor .


@param ProvisionNFVHeaderParams Custom header parameters
*/
func (s *SiteDesignService) ProvisionNfv(requestSiteDesignProvisionNFV *RequestSiteDesignProvisionNfv, ProvisionNFVHeaderParams *ProvisionNfvHeaderParams) (*ResponseSiteDesignProvisionNfv, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/nfv"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ProvisionNFVHeaderParams != nil {

		if ProvisionNFVHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", ProvisionNFVHeaderParams.Runsync)
		}

		if ProvisionNFVHeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", ProvisionNFVHeaderParams.Timeout)
		}

		if ProvisionNFVHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", ProvisionNFVHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSiteDesignProvisionNFV).
		SetResult(&ResponseSiteDesignProvisionNfv{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ProvisionNfv")
	}

	result := response.Result().(*ResponseSiteDesignProvisionNfv)
	return result, response, err

}

//NfvProvisioningDetail NFV Provisioning Detail - 2f97-e8fa-45f8-b2a3
/* Checks the provisioning detail of an ENCS device including log information.


@param NFVProvisioningDetailHeaderParams Custom header parameters
*/
func (s *SiteDesignService) NfvProvisioningDetail(requestSiteDesignNFVProvisioningDetail *RequestSiteDesignNfvProvisioningDetail, NFVProvisioningDetailHeaderParams *NfvProvisioningDetailHeaderParams) (*ResponseSiteDesignNfvProvisioningDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/nfv-provision-detail"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if NFVProvisioningDetailHeaderParams != nil {

		if NFVProvisioningDetailHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", NFVProvisioningDetailHeaderParams.Runsync)
		}

		if NFVProvisioningDetailHeaderParams.Runsynctimeout != "" {
			clientRequest = clientRequest.SetHeader("__runsynctimeout", NFVProvisioningDetailHeaderParams.Runsynctimeout)
		}

		if NFVProvisioningDetailHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", NFVProvisioningDetailHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestSiteDesignNFVProvisioningDetail).
		SetResult(&ResponseSiteDesignNfvProvisioningDetail{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation NfvProvisioningDetail")
	}

	result := response.Result().(*ResponseSiteDesignNfvProvisioningDetail)
	return result, response, err

}

//CreateNfvProfile Create NFV Profile - 6695-1aaa-407b-a89c
/* API to create network profile for different NFV topologies


 */
func (s *SiteDesignService) CreateNfvProfile(requestSiteDesignCreateNFVProfile *RequestSiteDesignCreateNfvProfile) (*ResponseSiteDesignCreateNfvProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/nfv/network-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreateNFVProfile).
		SetResult(&ResponseSiteDesignCreateNfvProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateNfvProfile")
	}

	result := response.Result().(*ResponseSiteDesignCreateNfvProfile)
	return result, response, err

}

//CreateFloormap Create Floormap - 2fad-58dc-4689-893d
/* Service to create a floor map with callback


 */
func (s *SiteDesignService) CreateFloormap(requestSiteDesignCreateFloormap *RequestSiteDesignCreateFloormap) (*ResponseSiteDesignCreateFloormap, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/floormap"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignCreateFloormap).
		// SetResult(&ResponseSiteDesignCreateFloormap{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateFloormap")
	}

	result := response.Result().(ResponseSiteDesignCreateFloormap)

	return &result, response, err

}

//UpdateNfvProfile Update NFV Profile - 0fa0-0adf-4869-8287
/* API to update a NFV Network profile


@param id id path parameter. Id of the NFV profile to be updated

@param UpdateNFVProfileQueryParams Filtering parameter
*/
func (s *SiteDesignService) UpdateNfvProfile(id string, requestSiteDesignUpdateNFVProfile *RequestSiteDesignUpdateNfvProfile, UpdateNFVProfileQueryParams *UpdateNfvProfileQueryParams) (*ResponseSiteDesignUpdateNfvProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/nfv/network-profile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(UpdateNFVProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSiteDesignUpdateNFVProfile).
		SetResult(&ResponseSiteDesignUpdateNfvProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateNfvProfile")
	}

	result := response.Result().(*ResponseSiteDesignUpdateNfvProfile)
	return result, response, err

}

//UpdateFloormap Update Floormap - df9a-1912-43c9-96da
/* Service to create a floor map with callback


@param floorID floorId path parameter. Group ID of the floor to be modified

*/
func (s *SiteDesignService) UpdateFloormap(floorID string, requestSiteDesignUpdateFloormap *RequestSiteDesignUpdateFloormap) (*ResponseSiteDesignUpdateFloormap, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/floormap/{floorId}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSiteDesignUpdateFloormap).
		// SetResult(&ResponseSiteDesignUpdateFloormap{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateFloormap")
	}

	result := response.Result().(ResponseSiteDesignUpdateFloormap)

	return &result, response, err

}

//DeleteNfvProfile Delete NFV Profile - 5ebb-fa25-41b8-b8a9
/* API to delete nfv network profile.


@param id id path parameter. Id of nfv network profile to delete.

@param DeleteNFVProfileQueryParams Filtering parameter
*/
func (s *SiteDesignService) DeleteNfvProfile(id string, DeleteNFVProfileQueryParams *DeleteNfvProfileQueryParams) (*ResponseSiteDesignDeleteNfvProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/nfv/network-profile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(DeleteNFVProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSiteDesignDeleteNfvProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteNfvProfile")
	}

	result := response.Result().(*ResponseSiteDesignDeleteNfvProfile)
	return result, response, err

}

//DeleteFloormap Delete Floormap - 9fb8-7a26-401a-a33e
/* Service to delete an (empty) floor map with callback


@param floorID floorId path parameter. Group ID of floor to be deleted

*/
func (s *SiteDesignService) DeleteFloormap(floorID string) (*ResponseSiteDesignDeleteFloormap, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/floormap/{floorId}"
	path = strings.Replace(path, "{floorId}", fmt.Sprintf("%v", floorID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		// SetResult(&ResponseSiteDesignDeleteFloormap{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteFloormap")
	}

	result := response.Result().(ResponseSiteDesignDeleteFloormap)

	return &result, response, err

}
