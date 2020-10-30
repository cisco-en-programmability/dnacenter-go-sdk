package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SiteDesignService is the service to communicate with the SiteDesign API endpoint
type SiteDesignService service

// Area is the Area definition
type Area struct {
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// Asav is the Asav definition
type Asav struct {
	Var1 string `json:"var1,omitempty"` //
}

// Building is the Building definition
type Building struct {
	Address    string `json:"address,omitempty"`    //
	Latitude   int    `json:"latitude,omitempty"`   //
	Longitude  int    `json:"longitude,omitempty"`  //
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// CreateNFVProfileRequest is the CreateNFVProfileRequest definition
type CreateNFVProfileRequest struct {
	Device      []Device `json:"device,omitempty"`      //
	ProfileName string   `json:"profileName,omitempty"` //
}

// CustomNetworks is the CustomNetworks definition
type CustomNetworks struct {
	ConnectionType    string              `json:"connectionType,omitempty"`    //
	NetworkName       string              `json:"networkName,omitempty"`       //
	ServicesToConnect []ServicesToConnect `json:"servicesToConnect,omitempty"` //
	VlanId            int                 `json:"vlanId,omitempty"`            //
	VlanMode          string              `json:"vlanMode,omitempty"`          //
}

// CustomServices is the CustomServices definition
type CustomServices struct {
	ApplicationType string   `json:"applicationType,omitempty"` //
	ImageName       string   `json:"imageName,omitempty"`       //
	Name            string   `json:"name,omitempty"`            //
	Profile         string   `json:"profile,omitempty"`         //
	Topology        Topology `json:"topology,omitempty"`        //
}

// CustomTemplate is the CustomTemplate definition
type CustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   //
	Template     string `json:"template,omitempty"`     //
	TemplateType string `json:"templateType,omitempty"` //
}

// Device is the Device definition
type Device struct {
	CurrentDeviceTag                string           `json:"currentDeviceTag,omitempty"`                //
	CustomNetworks                  []CustomNetworks `json:"customNetworks,omitempty"`                  //
	CustomTemplate                  []CustomTemplate `json:"customTemplate,omitempty"`                  //
	DeviceTag                       string           `json:"deviceTag,omitempty"`                       //
	DirectInternetAccessForFirewall bool             `json:"directInternetAccessForFirewall,omitempty"` //
	Services                        []Services       `json:"services,omitempty"`                        //
	VlanForL2                       []VlanForL2      `json:"vlanForL2,omitempty"`                       //
}

// Floor is the Floor definition
type Floor struct {
	Height     int    `json:"height,omitempty"`     //
	Length     int    `json:"length,omitempty"`     //
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
	RfModel    string `json:"rfModel,omitempty"`    //
	Width      int    `json:"width,omitempty"`      //
}

// NFVProvisioningDetailRequest is the NFVProvisioningDetailRequest definition
type NFVProvisioningDetailRequest struct {
	DeviceIp string `json:"device_ip,omitempty"` //
}

// Nfvis is the Nfvis definition
type Nfvis struct {
	Var1 string `json:"var1,omitempty"` //
}

// ProvisionNFVRequest is the ProvisionNFVRequest definition
type ProvisionNFVRequest struct {
	Provisioning []Provisioning `json:"provisioning,omitempty"` //
	SiteProfile  []SiteProfile  `json:"siteProfile,omitempty"`  //
}

// Provisioning is the Provisioning definition
type Provisioning struct {
	Device []Device `json:"device,omitempty"` //
	Site   Site     `json:"site,omitempty"`   //
}

// ServiceProviderProfile is the ServiceProviderProfile definition
type ServiceProviderProfile struct {
	Connect                    bool   `json:"connect,omitempty"`                    //
	ConnectDefaultGatewayOnWan bool   `json:"connectDefaultGatewayOnWan,omitempty"` //
	LinkType                   string `json:"linkType,omitempty"`                   //
	ServiceProvider            string `json:"serviceProvider,omitempty"`            //
}

// ServiceProviders is the ServiceProviders definition
type ServiceProviders struct {
	Connect         bool   `json:"connect,omitempty"`         //
	DefaultGateway  bool   `json:"defaultGateway,omitempty"`  //
	LinkType        string `json:"linkType,omitempty"`        //
	ServiceProvider string `json:"serviceProvider,omitempty"` //
}

// Services is the Services definition
type Services struct {
	FirewallMode string        `json:"firewallMode,omitempty"` //
	ImageName    string        `json:"imageName,omitempty"`    //
	ProfileType  string        `json:"profileType,omitempty"`  //
	ServiceName  string        `json:"serviceName,omitempty"`  //
	ServiceType  string        `json:"serviceType,omitempty"`  //
	VNicMapping  []VNicMapping `json:"vNicMapping,omitempty"`  //
}

// ServicesToConnect is the ServicesToConnect definition
type ServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` //
}

// Site is the Site definition
type Site struct {
	Area            Area     `json:"area,omitempty"`            //
	Building        Building `json:"building,omitempty"`        //
	Floor           Floor    `json:"floor,omitempty"`           //
	SiteProfileName string   `json:"siteProfileName,omitempty"` //
}

// SiteProfile is the SiteProfile definition
type SiteProfile struct {
	Device          []Device `json:"device,omitempty"`          //
	SiteProfileName string   `json:"siteProfileName,omitempty"` //
}

// SubPools is the SubPools definition
type SubPools struct {
	Gateway        string `json:"gateway,omitempty"`        //
	IpSubnet       string `json:"ipSubnet,omitempty"`       //
	Name           string `json:"name,omitempty"`           //
	ParentPoolName string `json:"parentPoolName,omitempty"` //
	Type           string `json:"type,omitempty"`           //
}

// TemplateParam is the TemplateParam definition
type TemplateParam struct {
	Asav  Asav  `json:"asav,omitempty"`  //
	Nfvis Nfvis `json:"nfvis,omitempty"` //
}

// Topology is the Topology definition
type Topology struct {
	AssignIp string `json:"assignIp,omitempty"` //
	Name     string `json:"name,omitempty"`     //
	Type     string `json:"type,omitempty"`     //
}

// UpdateNFVProfileRequest is the UpdateNFVProfileRequest definition
type UpdateNFVProfileRequest struct {
	Device []Device `json:"device,omitempty"` //
}

// VNicMapping is the VNicMapping definition
type VNicMapping struct {
	AssignIpAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
	NetworkType              string `json:"networkType,omitempty"`              //
}

// Vlan is the Vlan definition
type Vlan struct {
	Id   string `json:"id,omitempty"`   //
	Type string `json:"type,omitempty"` //
}

// VlanForL2 is the VlanForL2 definition
type VlanForL2 struct {
	VlanDescription string `json:"vlanDescription,omitempty"` //
	VlanId          int    `json:"vlanId,omitempty"`          //
	VlanType        string `json:"vlanType,omitempty"`        //
}

// WanInterface is the WanInterface definition
type WanInterface struct {
	Bandwidth     string `json:"bandwidth,omitempty"`     //
	Gateway       string `json:"gateway,omitempty"`       //
	InterfaceName string `json:"interfaceName,omitempty"` //
	IpAddress     string `json:"ipAddress,omitempty"`     //
	Subnetmask    string `json:"subnetmask,omitempty"`    //
}

// CreateNFVProfileResponse is the CreateNFVProfileResponse definition
type CreateNFVProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CustomNetworks is the CustomNetworks definition
type CustomNetworks struct {
	ConnectionType    string              `json:"connectionType,omitempty"`    //
	NetworkName       string              `json:"networkName,omitempty"`       //
	ServicesToConnect []ServicesToConnect `json:"servicesToConnect,omitempty"` //
	VlanId            string              `json:"vlanId,omitempty"`            //
	VlanMode          string              `json:"vlanMode,omitempty"`          //
}

// CustomTemplate is the CustomTemplate definition
type CustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   //
	Template     string `json:"template,omitempty"`     //
	TemplateType string `json:"templateType,omitempty"` //
}

// Device is the Device definition
type Device struct {
	CustomNetworks                  []CustomNetworks         `json:"customNetworks,omitempty"`                  //
	CustomTemplate                  []CustomTemplate         `json:"customTemplate,omitempty"`                  //
	DeviceTag                       string                   `json:"deviceTag,omitempty"`                       //
	DeviceType                      string                   `json:"deviceType,omitempty"`                      //
	DirectInternetAccessForFirewall bool                     `json:"directInternetAccessForFirewall,omitempty"` //
	ServiceProviderProfile          []ServiceProviderProfile `json:"serviceProviderProfile,omitempty"`          //
	Services                        []Services               `json:"services,omitempty"`                        //
	VlanForL2                       []VlanForL2              `json:"vlanForL2,omitempty"`                       //
}

// GetDeviceDetailsByIPResponse is the GetDeviceDetailsByIPResponse definition
type GetDeviceDetailsByIPResponse struct {
	ProvisionDetails ProvisionDetails `json:"provisionDetails,omitempty"` //
}

// GetNFVProfileResponse is the GetNFVProfileResponse definition
type GetNFVProfileResponse struct {
	Response []Response `json:"response,omitempty"` //
}

// NFVProvisioningDetailResponse is the NFVProvisioningDetailResponse definition
type NFVProvisioningDetailResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// ProvisionDetails is the ProvisionDetails definition
type ProvisionDetails struct {
	BeginStep     string      `json:"beginStep,omitempty"`     //
	Duration      string      `json:"duration,omitempty"`      //
	EndTime       string      `json:"endTime,omitempty"`       //
	StartTime     string      `json:"startTime,omitempty"`     //
	Status        string      `json:"status,omitempty"`        //
	StatusMessage string      `json:"statusMessage,omitempty"` //
	TaskNodes     []TaskNodes `json:"taskNodes,omitempty"`     //
	Topology      string      `json:"topology,omitempty"`      //
}

// ProvisionNFVResponse is the ProvisionNFVResponse definition
type ProvisionNFVResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// Response is the Response definition
type Response struct {
	Device      []Device `json:"device,omitempty"`      //
	Id          string   `json:"id,omitempty"`          //
	ProfileName string   `json:"profileName,omitempty"` //
}

// ServiceProviderProfile is the ServiceProviderProfile definition
type ServiceProviderProfile struct {
	Connect                    bool   `json:"connect,omitempty"`                    //
	ConnectDefaultGatewayOnWan bool   `json:"connectDefaultGatewayOnWan,omitempty"` //
	LinkType                   string `json:"linkType,omitempty"`                   //
	ServiceProvider            string `json:"serviceProvider,omitempty"`            //
}

// Services is the Services definition
type Services struct {
	FirewallMode string        `json:"firewallMode,omitempty"` //
	ImageName    string        `json:"imageName,omitempty"`    //
	ProfileType  string        `json:"profileType,omitempty"`  //
	ServiceName  string        `json:"serviceName,omitempty"`  //
	ServiceType  string        `json:"serviceType,omitempty"`  //
	VNicMapping  []VNicMapping `json:"vNicMapping,omitempty"`  //
}

// ServicesToConnect is the ServicesToConnect definition
type ServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` //
}

// TaskNodes is the TaskNodes definition
type TaskNodes struct {
	CliTemplateUserMessageDTO string `json:"cliTemplateUserMessageDTO,omitempty"` //
	Duration                  string `json:"duration,omitempty"`                  //
	EndTime                   string `json:"endTime,omitempty"`                   //
	ErrorPayload              string `json:"errorPayload,omitempty"`              //
	Name                      string `json:"name,omitempty"`                      //
	NextTask                  string `json:"nextTask,omitempty"`                  //
	ParentTask                string `json:"parentTask,omitempty"`                //
	Payload                   string `json:"payload,omitempty"`                   //
	ProvisionedNames          string `json:"provisionedNames,omitempty"`          //
	StartTime                 string `json:"startTime,omitempty"`                 //
	Status                    string `json:"status,omitempty"`                    //
	StatusMessage             string `json:"statusMessage,omitempty"`             //
	StepRan                   string `json:"stepRan,omitempty"`                   //
	Target                    string `json:"target,omitempty"`                    //
}

// UpdateNFVProfileResponse is the UpdateNFVProfileResponse definition
type UpdateNFVProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// VNicMapping is the VNicMapping definition
type VNicMapping struct {
	AssignIpAddressToNetwork bool   `json:"assignIpAddressToNetwork,omitempty"` //
	NetworkType              string `json:"networkType,omitempty"`              //
}

// VlanForL2 is the VlanForL2 definition
type VlanForL2 struct {
	VlanDescription string `json:"vlanDescription,omitempty"` //
	VlanId          string `json:"vlanId,omitempty"`          //
	VlanType        string `json:"vlanType,omitempty"`        //
}

// CreateNFVProfile createNFVProfile
/* API to create network profile for different NFV topologies
 */
func (s *SiteDesignService) CreateNFVProfile(createNFVProfileRequest *CreateNFVProfileRequest) (*CreateNFVProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/nfv/network-profile"

	response, err := RestyClient.R().
		SetBody(createNFVProfileRequest).
		SetResult(&CreateNFVProfileResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateNFVProfileResponse)
	return result, response, err

}

// GetDeviceDetailsByIPQueryParams defines the query parameters for this request
type GetDeviceDetailsByIPQueryParams struct {
	DeviceIp string `url:"deviceIp,omitempty"` // Device to which the provisioning detail has to be retrieved
}

// GetDeviceDetailsByIP getDeviceDetailsByIP
/* Returns provisioning device information for the specified IP address.
@param deviceIp Device to which the provisioning detail has to be retrieved
*/
func (s *SiteDesignService) GetDeviceDetailsByIP(getDeviceDetailsByIPQueryParams *GetDeviceDetailsByIPQueryParams) (*GetDeviceDetailsByIPResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/nfv/provisioningDetail"

	queryString, _ := query.Values(getDeviceDetailsByIPQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceDetailsByIPResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
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

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetNFVProfileResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
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

	response, err := RestyClient.R().
		SetBody(nFVProvisioningDetailRequest).
		SetResult(&NFVProvisioningDetailResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
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

	response, err := RestyClient.R().
		SetBody(provisionNFVRequest).
		SetResult(&ProvisionNFVResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
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

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetBody(updateNFVProfileRequest).
		SetResult(&UpdateNFVProfileResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateNFVProfileResponse)
	return result, response, err

}
