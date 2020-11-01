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
	Device []struct {
		CustomNetworks []struct {
			ConnectionType    string `json:"connectionType,omitempty"` //
			NetworkName       string `json:"networkName,omitempty"`    //
			ServicesToConnect []struct {
				ServiceName string `json:"serviceName,omitempty"` //
			} `json:"servicesToConnect,omitempty"` //
			VlanID   int    `json:"vlanId,omitempty"`   //
			VlanMode string `json:"vlanMode,omitempty"` //
		} `json:"customNetworks,omitempty"` //
		CustomTemplate []struct {
			DeviceType   string `json:"deviceType,omitempty"`   //
			Template     string `json:"template,omitempty"`     //
			TemplateType string `json:"templateType,omitempty"` //
		} `json:"customTemplate,omitempty"` //
		DeviceTag                       string `json:"deviceTag,omitempty"`                       //
		DeviceType                      string `json:"deviceType,omitempty"`                      //
		DirectInternetAccessForFirewall bool   `json:"directInternetAccessForFirewall,omitempty"` //
		ServiceProviderProfile          []struct {
			Connect                    bool   `json:"connect,omitempty"`                    //
			ConnectDefaultGatewayOnWan bool   `json:"connectDefaultGatewayOnWan,omitempty"` //
			LinkType                   string `json:"linkType,omitempty"`                   //
			ServiceProvider            string `json:"serviceProvider,omitempty"`            //
		} `json:"serviceProviderProfile,omitempty"` //
		Services []struct {
			FirewallMode string `json:"firewallMode,omitempty"` //
			ImageName    string `json:"imageName,omitempty"`    //
			ProfileType  string `json:"profileType,omitempty"`  //
			ServiceName  string `json:"serviceName,omitempty"`  //
			ServiceType  string `json:"serviceType,omitempty"`  //
			VNicMapping  []struct {
				AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
				NetworkType              string `json:"networkType,omitempty"`              //
			} `json:"vNicMapping,omitempty"` //
		} `json:"services,omitempty"` //
		VlanForL2 []struct {
			VlanDescription string `json:"vlanDescription,omitempty"` //
			VlanID          int    `json:"vlanId,omitempty"`          //
			VlanType        string `json:"vlanType,omitempty"`        //
		} `json:"vlanForL2,omitempty"` //
	} `json:"device,omitempty"` //
	ProfileName string `json:"profileName,omitempty"` //
}

// CustomNetworks is the CustomNetworks definition
type CustomNetworks struct {
	ConnectionType    string `json:"connectionType,omitempty"` //
	NetworkName       string `json:"networkName,omitempty"`    //
	ServicesToConnect []struct {
		ServiceName string `json:"serviceName,omitempty"` //
	} `json:"servicesToConnect,omitempty"` //
	VlanID   int    `json:"vlanId,omitempty"`   //
	VlanMode string `json:"vlanMode,omitempty"` //
}

// CustomServices is the CustomServices definition
type CustomServices struct {
	ApplicationType string `json:"applicationType,omitempty"` //
	ImageName       string `json:"imageName,omitempty"`       //
	Name            string `json:"name,omitempty"`            //
	Profile         string `json:"profile,omitempty"`         //
	Topology        struct {
		AssignIP string `json:"assignIp,omitempty"` //
		Name     string `json:"name,omitempty"`     //
		Type     string `json:"type,omitempty"`     //
	} `json:"topology,omitempty"` //
}

// CustomTemplate is the CustomTemplate definition
type CustomTemplate struct {
	DeviceType   string `json:"deviceType,omitempty"`   //
	Template     string `json:"template,omitempty"`     //
	TemplateType string `json:"templateType,omitempty"` //
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
	DeviceIP string `json:"device_ip,omitempty"` //
}

// Nfvis is the Nfvis definition
type Nfvis struct {
	Var1 string `json:"var1,omitempty"` //
}

// ProvisionNFVRequest is the ProvisionNFVRequest definition
type ProvisionNFVRequest struct {
	Provisioning []struct {
		Device []struct {
			CustomNetworks []struct {
				IPAddressPool string `json:"ipAddressPool,omitempty"` //
				Name          string `json:"name,omitempty"`          //
				Port          string `json:"port,omitempty"`          //
			} `json:"customNetworks,omitempty"` //
			DeviceSerialNumber string `json:"deviceSerialNumber,omitempty"` //
			IP                 string `json:"ip,omitempty"`                 //
			ServiceProviders   []struct {
				ServiceProvider string `json:"serviceProvider,omitempty"` //
				WanInterface    struct {
					Bandwidth     string `json:"bandwidth,omitempty"`     //
					Gateway       string `json:"gateway,omitempty"`       //
					InterfaceName string `json:"interfaceName,omitempty"` //
					IPAddress     string `json:"ipAddress,omitempty"`     //
					Subnetmask    string `json:"subnetmask,omitempty"`    //
				} `json:"wanInterface,omitempty"` //
			} `json:"serviceProviders,omitempty"` //
			Services []struct {
				AdminPasswordHash      string `json:"adminPasswordHash,omitempty"`      //
				CentralManagerIP       string `json:"centralManagerIP,omitempty"`       //
				CentralRegistrationKey string `json:"centralRegistrationKey,omitempty"` //
				CommonKey              string `json:"commonKey,omitempty"`              //
				Disk                   string `json:"disk,omitempty"`                   //
				Mode                   string `json:"mode,omitempty"`                   //
				SystemIP               string `json:"systemIp,omitempty"`               //
				Type                   string `json:"type,omitempty"`                   //
			} `json:"services,omitempty"` //
			SubPools []struct {
				Gateway        string `json:"gateway,omitempty"`        //
				IPSubnet       string `json:"ipSubnet,omitempty"`       //
				Name           string `json:"name,omitempty"`           //
				ParentPoolName string `json:"parentPoolName,omitempty"` //
				Type           string `json:"type,omitempty"`           //
			} `json:"subPools,omitempty"` //
			TagName       string `json:"tagName,omitempty"` //
			TemplateParam struct {
				Asav struct {
					Var1 string `json:"var1,omitempty"` //
				} `json:"asav,omitempty"` //
				Nfvis struct {
					Var1 string `json:"var1,omitempty"` //
				} `json:"nfvis,omitempty"` //
			} `json:"templateParam,omitempty"` //
			Vlan []struct {
				ID         string `json:"id,omitempty"`         //
				Interfaces string `json:"interfaces,omitempty"` //
				Network    string `json:"network,omitempty"`    //
				Type       string `json:"type,omitempty"`       //
			} `json:"vlan,omitempty"` //
		} `json:"device,omitempty"` //
		Site struct {
			Area struct {
				Name       string `json:"name,omitempty"`       //
				ParentName string `json:"parentName,omitempty"` //
			} `json:"area,omitempty"` //
			Building struct {
				Address    string `json:"address,omitempty"`    //
				Latitude   int    `json:"latitude,omitempty"`   //
				Longitude  int    `json:"longitude,omitempty"`  //
				Name       string `json:"name,omitempty"`       //
				ParentName string `json:"parentName,omitempty"` //
			} `json:"building,omitempty"` //
			Floor struct {
				Height     int    `json:"height,omitempty"`     //
				Length     int    `json:"length,omitempty"`     //
				Name       string `json:"name,omitempty"`       //
				ParentName string `json:"parentName,omitempty"` //
				RfModel    string `json:"rfModel,omitempty"`    //
				Width      int    `json:"width,omitempty"`      //
			} `json:"floor,omitempty"` //
			SiteProfileName string `json:"siteProfileName,omitempty"` //
		} `json:"site,omitempty"` //
	} `json:"provisioning,omitempty"` //
	SiteProfile []struct {
		Device []struct {
			CustomNetworks []struct {
				ConnectionType    string `json:"connectionType,omitempty"` //
				Name              string `json:"name,omitempty"`           //
				NetworkMode       string `json:"networkMode,omitempty"`    //
				ServicesToConnect []struct {
					Service string `json:"service,omitempty"` //
				} `json:"servicesToConnect,omitempty"` //
				Vlan string `json:"vlan,omitempty"` //
			} `json:"customNetworks,omitempty"` //
			CustomServices []struct {
				ApplicationType string `json:"applicationType,omitempty"` //
				ImageName       string `json:"imageName,omitempty"`       //
				Name            string `json:"name,omitempty"`            //
				Profile         string `json:"profile,omitempty"`         //
				Topology        struct {
					AssignIP string `json:"assignIp,omitempty"` //
					Name     string `json:"name,omitempty"`     //
					Type     string `json:"type,omitempty"`     //
				} `json:"topology,omitempty"` //
			} `json:"customServices,omitempty"` //
			CustomTemplate []struct {
				DeviceType string `json:"deviceType,omitempty"` //
				Template   string `json:"template,omitempty"`   //
			} `json:"customTemplate,omitempty"` //
			DeviceType       string `json:"deviceType,omitempty"` //
			Dia              bool   `json:"dia,omitempty"`        //
			ServiceProviders []struct {
				Connect         bool   `json:"connect,omitempty"`         //
				DefaultGateway  bool   `json:"defaultGateway,omitempty"`  //
				LinkType        string `json:"linkType,omitempty"`        //
				ServiceProvider string `json:"serviceProvider,omitempty"` //
			} `json:"serviceProviders,omitempty"` //
			Services []struct {
				ImageName string `json:"imageName,omitempty"` //
				Mode      string `json:"mode,omitempty"`      //
				Name      string `json:"name,omitempty"`      //
				Profile   string `json:"profile,omitempty"`   //
				Topology  struct {
					AssignIP string `json:"assignIp,omitempty"` //
					Name     string `json:"name,omitempty"`     //
					Type     string `json:"type,omitempty"`     //
				} `json:"topology,omitempty"` //
				Type string `json:"type,omitempty"` //
			} `json:"services,omitempty"` //
			TagName string `json:"tagName,omitempty"` //
			Vlan    []struct {
				ID   string `json:"id,omitempty"`   //
				Type string `json:"type,omitempty"` //
			} `json:"vlan,omitempty"` //
		} `json:"device,omitempty"` //
		SiteProfileName string `json:"siteProfileName,omitempty"` //
	} `json:"siteProfile,omitempty"` //
}

// Provisioning is the Provisioning definition
type Provisioning struct {
	Device []struct {
		CustomNetworks []struct {
			IPAddressPool string `json:"ipAddressPool,omitempty"` //
			Name          string `json:"name,omitempty"`          //
			Port          string `json:"port,omitempty"`          //
		} `json:"customNetworks,omitempty"` //
		DeviceSerialNumber string `json:"deviceSerialNumber,omitempty"` //
		IP                 string `json:"ip,omitempty"`                 //
		ServiceProviders   []struct {
			ServiceProvider string `json:"serviceProvider,omitempty"` //
			WanInterface    struct {
				Bandwidth     string `json:"bandwidth,omitempty"`     //
				Gateway       string `json:"gateway,omitempty"`       //
				InterfaceName string `json:"interfaceName,omitempty"` //
				IPAddress     string `json:"ipAddress,omitempty"`     //
				Subnetmask    string `json:"subnetmask,omitempty"`    //
			} `json:"wanInterface,omitempty"` //
		} `json:"serviceProviders,omitempty"` //
		Services []struct {
			AdminPasswordHash      string `json:"adminPasswordHash,omitempty"`      //
			CentralManagerIP       string `json:"centralManagerIP,omitempty"`       //
			CentralRegistrationKey string `json:"centralRegistrationKey,omitempty"` //
			CommonKey              string `json:"commonKey,omitempty"`              //
			Disk                   string `json:"disk,omitempty"`                   //
			Mode                   string `json:"mode,omitempty"`                   //
			SystemIP               string `json:"systemIp,omitempty"`               //
			Type                   string `json:"type,omitempty"`                   //
		} `json:"services,omitempty"` //
		SubPools []struct {
			Gateway        string `json:"gateway,omitempty"`        //
			IPSubnet       string `json:"ipSubnet,omitempty"`       //
			Name           string `json:"name,omitempty"`           //
			ParentPoolName string `json:"parentPoolName,omitempty"` //
			Type           string `json:"type,omitempty"`           //
		} `json:"subPools,omitempty"` //
		TagName       string `json:"tagName,omitempty"` //
		TemplateParam struct {
			Asav struct {
				Var1 string `json:"var1,omitempty"` //
			} `json:"asav,omitempty"` //
			Nfvis struct {
				Var1 string `json:"var1,omitempty"` //
			} `json:"nfvis,omitempty"` //
		} `json:"templateParam,omitempty"` //
		Vlan []struct {
			ID         string `json:"id,omitempty"`         //
			Interfaces string `json:"interfaces,omitempty"` //
			Network    string `json:"network,omitempty"`    //
			Type       string `json:"type,omitempty"`       //
		} `json:"vlan,omitempty"` //
	} `json:"device,omitempty"` //
	Site struct {
		Area struct {
			Name       string `json:"name,omitempty"`       //
			ParentName string `json:"parentName,omitempty"` //
		} `json:"area,omitempty"` //
		Building struct {
			Address    string `json:"address,omitempty"`    //
			Latitude   int    `json:"latitude,omitempty"`   //
			Longitude  int    `json:"longitude,omitempty"`  //
			Name       string `json:"name,omitempty"`       //
			ParentName string `json:"parentName,omitempty"` //
		} `json:"building,omitempty"` //
		Floor struct {
			Height     int    `json:"height,omitempty"`     //
			Length     int    `json:"length,omitempty"`     //
			Name       string `json:"name,omitempty"`       //
			ParentName string `json:"parentName,omitempty"` //
			RfModel    string `json:"rfModel,omitempty"`    //
			Width      int    `json:"width,omitempty"`      //
		} `json:"floor,omitempty"` //
		SiteProfileName string `json:"siteProfileName,omitempty"` //
	} `json:"site,omitempty"` //
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
	FirewallMode string `json:"firewallMode,omitempty"` //
	ImageName    string `json:"imageName,omitempty"`    //
	ProfileType  string `json:"profileType,omitempty"`  //
	ServiceName  string `json:"serviceName,omitempty"`  //
	ServiceType  string `json:"serviceType,omitempty"`  //
	VNicMapping  []struct {
		AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
		NetworkType              string `json:"networkType,omitempty"`              //
	} `json:"vNicMapping,omitempty"` //
}

// ServicesToConnect is the ServicesToConnect definition
type ServicesToConnect struct {
	ServiceName string `json:"serviceName,omitempty"` //
}

// Site is the Site definition
type Site struct {
	Area struct {
		Name       string `json:"name,omitempty"`       //
		ParentName string `json:"parentName,omitempty"` //
	} `json:"area,omitempty"` //
	Building struct {
		Address    string `json:"address,omitempty"`    //
		Latitude   int    `json:"latitude,omitempty"`   //
		Longitude  int    `json:"longitude,omitempty"`  //
		Name       string `json:"name,omitempty"`       //
		ParentName string `json:"parentName,omitempty"` //
	} `json:"building,omitempty"` //
	Floor struct {
		Height     int    `json:"height,omitempty"`     //
		Length     int    `json:"length,omitempty"`     //
		Name       string `json:"name,omitempty"`       //
		ParentName string `json:"parentName,omitempty"` //
		RfModel    string `json:"rfModel,omitempty"`    //
		Width      int    `json:"width,omitempty"`      //
	} `json:"floor,omitempty"` //
	SiteProfileName string `json:"siteProfileName,omitempty"` //
}

// SiteProfile is the SiteProfile definition
type SiteProfile struct {
	Device []struct {
		CustomNetworks []struct {
			ConnectionType    string `json:"connectionType,omitempty"` //
			Name              string `json:"name,omitempty"`           //
			NetworkMode       string `json:"networkMode,omitempty"`    //
			ServicesToConnect []struct {
				Service string `json:"service,omitempty"` //
			} `json:"servicesToConnect,omitempty"` //
			Vlan string `json:"vlan,omitempty"` //
		} `json:"customNetworks,omitempty"` //
		CustomServices []struct {
			ApplicationType string `json:"applicationType,omitempty"` //
			ImageName       string `json:"imageName,omitempty"`       //
			Name            string `json:"name,omitempty"`            //
			Profile         string `json:"profile,omitempty"`         //
			Topology        struct {
				AssignIP string `json:"assignIp,omitempty"` //
				Name     string `json:"name,omitempty"`     //
				Type     string `json:"type,omitempty"`     //
			} `json:"topology,omitempty"` //
		} `json:"customServices,omitempty"` //
		CustomTemplate []struct {
			DeviceType string `json:"deviceType,omitempty"` //
			Template   string `json:"template,omitempty"`   //
		} `json:"customTemplate,omitempty"` //
		DeviceType       string `json:"deviceType,omitempty"` //
		Dia              bool   `json:"dia,omitempty"`        //
		ServiceProviders []struct {
			Connect         bool   `json:"connect,omitempty"`         //
			DefaultGateway  bool   `json:"defaultGateway,omitempty"`  //
			LinkType        string `json:"linkType,omitempty"`        //
			ServiceProvider string `json:"serviceProvider,omitempty"` //
		} `json:"serviceProviders,omitempty"` //
		Services []struct {
			ImageName string `json:"imageName,omitempty"` //
			Mode      string `json:"mode,omitempty"`      //
			Name      string `json:"name,omitempty"`      //
			Profile   string `json:"profile,omitempty"`   //
			Topology  struct {
				AssignIP string `json:"assignIp,omitempty"` //
				Name     string `json:"name,omitempty"`     //
				Type     string `json:"type,omitempty"`     //
			} `json:"topology,omitempty"` //
			Type string `json:"type,omitempty"` //
		} `json:"services,omitempty"` //
		TagName string `json:"tagName,omitempty"` //
		Vlan    []struct {
			ID   string `json:"id,omitempty"`   //
			Type string `json:"type,omitempty"` //
		} `json:"vlan,omitempty"` //
	} `json:"device,omitempty"` //
	SiteProfileName string `json:"siteProfileName,omitempty"` //
}

// SubPools is the SubPools definition
type SubPools struct {
	Gateway        string `json:"gateway,omitempty"`        //
	IPSubnet       string `json:"ipSubnet,omitempty"`       //
	Name           string `json:"name,omitempty"`           //
	ParentPoolName string `json:"parentPoolName,omitempty"` //
	Type           string `json:"type,omitempty"`           //
}

// TemplateParam is the TemplateParam definition
type TemplateParam struct {
	Asav struct {
		Var1 string `json:"var1,omitempty"` //
	} `json:"asav,omitempty"` //
	Nfvis struct {
		Var1 string `json:"var1,omitempty"` //
	} `json:"nfvis,omitempty"` //
}

// UpdateNFVProfileRequest is the UpdateNFVProfileRequest definition
type UpdateNFVProfileRequest struct {
	Device []struct {
		CurrentDeviceTag string `json:"currentDeviceTag,omitempty"` //
		CustomNetworks   []struct {
			ConnectionType    string `json:"connectionType,omitempty"` //
			NetworkName       string `json:"networkName,omitempty"`    //
			ServicesToConnect []struct {
				ServiceName string `json:"serviceName,omitempty"` //
			} `json:"servicesToConnect,omitempty"` //
			VlanID   int    `json:"vlanId,omitempty"`   //
			VlanMode string `json:"vlanMode,omitempty"` //
		} `json:"customNetworks,omitempty"` //
		CustomTemplate []struct {
			DeviceType   string `json:"deviceType,omitempty"`   //
			Template     string `json:"template,omitempty"`     //
			TemplateType string `json:"templateType,omitempty"` //
		} `json:"customTemplate,omitempty"` //
		DeviceTag                       string `json:"deviceTag,omitempty"`                       //
		DirectInternetAccessForFirewall bool   `json:"directInternetAccessForFirewall,omitempty"` //
		Services                        []struct {
			FirewallMode string `json:"firewallMode,omitempty"` //
			ImageName    string `json:"imageName,omitempty"`    //
			ProfileType  string `json:"profileType,omitempty"`  //
			ServiceName  string `json:"serviceName,omitempty"`  //
			ServiceType  string `json:"serviceType,omitempty"`  //
			VNicMapping  []struct {
				AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
				NetworkType              string `json:"networkType,omitempty"`              //
			} `json:"vNicMapping,omitempty"` //
		} `json:"services,omitempty"` //
		VlanForL2 []struct {
			VlanDescription string `json:"vlanDescription,omitempty"` //
			VlanID          int    `json:"vlanId,omitempty"`          //
			VlanType        string `json:"vlanType,omitempty"`        //
		} `json:"vlanForL2,omitempty"` //
	} `json:"device,omitempty"` //
}

// VNicMapping is the VNicMapping definition
type VNicMapping struct {
	AssignIPAddressToNetwork string `json:"assignIpAddressToNetwork,omitempty"` //
	NetworkType              string `json:"networkType,omitempty"`              //
}

// Vlan is the Vlan definition
type Vlan struct {
	ID   string `json:"id,omitempty"`   //
	Type string `json:"type,omitempty"` //
}

// VlanForL2 is the VlanForL2 definition
type VlanForL2 struct {
	VlanDescription string `json:"vlanDescription,omitempty"` //
	VlanID          int    `json:"vlanId,omitempty"`          //
	VlanType        string `json:"vlanType,omitempty"`        //
}

// WanInterface is the WanInterface definition
type WanInterface struct {
	Bandwidth     string `json:"bandwidth,omitempty"`     //
	Gateway       string `json:"gateway,omitempty"`       //
	InterfaceName string `json:"interfaceName,omitempty"` //
	IPAddress     string `json:"ipAddress,omitempty"`     //
	Subnetmask    string `json:"subnetmask,omitempty"`    //
}

// CreateNFVProfileResponse is the CreateNFVProfileResponse definition
type CreateNFVProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetDeviceDetailsByIPResponse is the GetDeviceDetailsByIPResponse definition
type GetDeviceDetailsByIPResponse struct {
	ProvisionDetails struct {
		BeginStep     string `json:"beginStep,omitempty"`     //
		Duration      string `json:"duration,omitempty"`      //
		EndTime       string `json:"endTime,omitempty"`       //
		StartTime     string `json:"startTime,omitempty"`     //
		Status        string `json:"status,omitempty"`        //
		StatusMessage string `json:"statusMessage,omitempty"` //
		TaskNodes     []struct {
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
		} `json:"taskNodes,omitempty"` //
		Topology string `json:"topology,omitempty"` //
	} `json:"provisionDetails,omitempty"` //
}

// GetNFVProfileResponse is the GetNFVProfileResponse definition
type GetNFVProfileResponse struct {
	Response []struct {
		Device []struct {
			CustomNetworks []struct {
				ConnectionType    string `json:"connectionType,omitempty"` //
				NetworkName       string `json:"networkName,omitempty"`    //
				ServicesToConnect []struct {
					ServiceName string `json:"serviceName,omitempty"` //
				} `json:"servicesToConnect,omitempty"` //
				VlanID   string `json:"vlanId,omitempty"`   //
				VlanMode string `json:"vlanMode,omitempty"` //
			} `json:"customNetworks,omitempty"` //
			CustomTemplate []struct {
				DeviceType   string `json:"deviceType,omitempty"`   //
				Template     string `json:"template,omitempty"`     //
				TemplateType string `json:"templateType,omitempty"` //
			} `json:"customTemplate,omitempty"` //
			DeviceTag                       string `json:"deviceTag,omitempty"`                       //
			DeviceType                      string `json:"deviceType,omitempty"`                      //
			DirectInternetAccessForFirewall bool   `json:"directInternetAccessForFirewall,omitempty"` //
			ServiceProviderProfile          []struct {
				Connect                    bool   `json:"connect,omitempty"`                    //
				ConnectDefaultGatewayOnWan bool   `json:"connectDefaultGatewayOnWan,omitempty"` //
				LinkType                   string `json:"linkType,omitempty"`                   //
				ServiceProvider            string `json:"serviceProvider,omitempty"`            //
			} `json:"serviceProviderProfile,omitempty"` //
			Services []struct {
				FirewallMode string `json:"firewallMode,omitempty"` //
				ImageName    string `json:"imageName,omitempty"`    //
				ProfileType  string `json:"profileType,omitempty"`  //
				ServiceName  string `json:"serviceName,omitempty"`  //
				ServiceType  string `json:"serviceType,omitempty"`  //
				VNicMapping  []struct {
					AssignIPAddressToNetwork bool   `json:"assignIpAddressToNetwork,omitempty"` //
					NetworkType              string `json:"networkType,omitempty"`              //
				} `json:"vNicMapping,omitempty"` //
			} `json:"services,omitempty"` //
			VlanForL2 []struct {
				VlanDescription string `json:"vlanDescription,omitempty"` //
				VlanID          string `json:"vlanId,omitempty"`          //
				VlanType        string `json:"vlanType,omitempty"`        //
			} `json:"vlanForL2,omitempty"` //
		} `json:"device,omitempty"` //
		ID          string `json:"id,omitempty"`          //
		ProfileName string `json:"profileName,omitempty"` //
	} `json:"response,omitempty"` //
}

// NFVProvisioningDetailResponse is the NFVProvisioningDetailResponse definition
type NFVProvisioningDetailResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// ProvisionDetails is the ProvisionDetails definition
type ProvisionDetails struct {
	BeginStep     string `json:"beginStep,omitempty"`     //
	Duration      string `json:"duration,omitempty"`      //
	EndTime       string `json:"endTime,omitempty"`       //
	StartTime     string `json:"startTime,omitempty"`     //
	Status        string `json:"status,omitempty"`        //
	StatusMessage string `json:"statusMessage,omitempty"` //
	TaskNodes     []struct {
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
	} `json:"taskNodes,omitempty"` //
	Topology string `json:"topology,omitempty"` //
}

// ProvisionNFVResponse is the ProvisionNFVResponse definition
type ProvisionNFVResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
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
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
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
	DeviceIP string `url:"deviceIp,omitempty"` // Device to which the provisioning detail has to be retrieved
}

// GetDeviceDetailsByIP getDeviceDetailsByIP
/* Returns provisioning device information for the specified IP address.
@param deviceIP Device to which the provisioning detail has to be retrieved
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
@param id ID of the NFV profile to be updated
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
