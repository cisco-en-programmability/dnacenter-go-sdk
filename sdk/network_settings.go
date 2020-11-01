package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// NetworkSettingsService is the service to communicate with the NetworkSettings API endpoint
type NetworkSettingsService service

// AssignCredentialToSiteRequest is the AssignCredentialToSiteRequest definition
type AssignCredentialToSiteRequest struct {
	CliID         string `json:"cliId,omitempty"`         //
	HTTPRead      string `json:"httpRead,omitempty"`      //
	HTTPWrite     string `json:"httpWrite,omitempty"`     //
	SNMPV2ReadID  string `json:"snmpV2ReadId,omitempty"`  //
	SNMPV2WriteID string `json:"snmpV2WriteId,omitempty"` //
	SNMPV3ID      string `json:"snmpV3Id,omitempty"`      //
}

// CliCredential is the CliCredential definition
type CliCredential struct {
	Description    string `json:"description,omitempty"`    //
	EnablePassword string `json:"enablePassword,omitempty"` //
	ID             string `json:"id,omitempty"`             //
	Password       string `json:"password,omitempty"`       //
	Username       string `json:"username,omitempty"`       //
}

// ClientAndEndpointAAA is the ClientAndEndpoint_aaa definition
type ClientAndEndpointAAA struct {
	IPAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// CreateDeviceCredentialsRequest is the CreateDeviceCredentialsRequest definition
type CreateDeviceCredentialsRequest struct {
	Settings struct {
		CliCredential []struct {
			Description    string `json:"description,omitempty"`    //
			EnablePassword string `json:"enablePassword,omitempty"` //
			Password       string `json:"password,omitempty"`       //
			Username       string `json:"username,omitempty"`       //
		} `json:"cliCredential,omitempty"` //
		HTTPSRead []struct {
			Name     string `json:"name,omitempty"`     //
			Password string `json:"password,omitempty"` //
			Port     int    `json:"port,omitempty"`     //
			Username string `json:"username,omitempty"` //
		} `json:"httpsRead,omitempty"` //
		HTTPSWrite []struct {
			Name     string `json:"name,omitempty"`     //
			Password string `json:"password,omitempty"` //
			Port     int    `json:"port,omitempty"`     //
			Username string `json:"username,omitempty"` //
		} `json:"httpsWrite,omitempty"` //
		SNMPV2cRead []struct {
			Description   string `json:"description,omitempty"`   //
			ReadCommunity string `json:"readCommunity,omitempty"` //
		} `json:"snmpV2cRead,omitempty"` //
		SNMPV2cWrite []struct {
			Description    string `json:"description,omitempty"`    //
			WriteCommunity string `json:"writeCommunity,omitempty"` //
		} `json:"snmpV2cWrite,omitempty"` //
		SNMPV3 []struct {
			AuthPassword    string `json:"authPassword,omitempty"`    //
			AuthType        string `json:"authType,omitempty"`        //
			Description     string `json:"description,omitempty"`     //
			PrivacyPassword string `json:"privacyPassword,omitempty"` //
			PrivacyType     string `json:"privacyType,omitempty"`     //
			SNMPMode        string `json:"snmpMode,omitempty"`        //
			Username        string `json:"username,omitempty"`        //
		} `json:"snmpV3,omitempty"` //
	} `json:"settings,omitempty"` //
}

// CreateGlobalPoolRequest is the CreateGlobalPoolRequest definition
type CreateGlobalPoolRequest struct {
	Settings struct {
		IPpool []struct {
			IPAddressSpace string   `json:"IpAddressSpace,omitempty"` //
			DhcpServerIPs  []string `json:"dhcpServerIps,omitempty"`  //
			DNSServerIPs   []string `json:"dnsServerIps,omitempty"`   //
			Gateway        string   `json:"gateway,omitempty"`        //
			IPPoolCidr     string   `json:"ipPoolCidr,omitempty"`     //
			IPPoolName     string   `json:"ipPoolName,omitempty"`     //
			Type           string   `json:"type,omitempty"`           //
		} `json:"ippool,omitempty"` //
	} `json:"settings,omitempty"` //
}

// CreateNetworkRequest is the CreateNetworkRequest definition
type CreateNetworkRequest struct {
	Settings struct {
		ClientAndEndpointAaa struct {
			IPAddress    string `json:"ipAddress,omitempty"`    //
			Network      string `json:"network,omitempty"`      //
			Protocol     string `json:"protocol,omitempty"`     //
			Servers      string `json:"servers,omitempty"`      //
			SharedSecret string `json:"sharedSecret,omitempty"` //
		} `json:"clientAndEndpoint_aaa,omitempty"` //
		DhcpServer []string `json:"dhcpServer,omitempty"` //
		DNSServer  struct {
			DomainName         string `json:"domainName,omitempty"`         //
			PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   //
			SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` //
		} `json:"dnsServer,omitempty"` //
		MessageOfTheday struct {
			BannerMessage        string `json:"bannerMessage,omitempty"`        //
			RetainExistingBanner bool   `json:"retainExistingBanner,omitempty"` //
		} `json:"messageOfTheday,omitempty"` //
		Netflowcollector struct {
			IPAddress string `json:"ipAddress,omitempty"` //
			Port      int    `json:"port,omitempty"`      //
		} `json:"netflowcollector,omitempty"` //
		NetworkAaa struct {
			IPAddress    string `json:"ipAddress,omitempty"`    //
			Network      string `json:"network,omitempty"`      //
			Protocol     string `json:"protocol,omitempty"`     //
			Servers      string `json:"servers,omitempty"`      //
			SharedSecret string `json:"sharedSecret,omitempty"` //
		} `json:"network_aaa,omitempty"` //
		NtpServer  []string `json:"ntpServer,omitempty"` //
		SNMPServer struct {
			ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
			IPAddresses     []string `json:"ipAddresses,omitempty"`     //
		} `json:"snmpServer,omitempty"` //
		SyslogServer struct {
			ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
			IPAddresses     []string `json:"ipAddresses,omitempty"`     //
		} `json:"syslogServer,omitempty"` //
		Timezone string `json:"timezone,omitempty"` //
	} `json:"settings,omitempty"` //
}

// CreateSPProfileRequest is the CreateSPProfileRequest definition
type CreateSPProfileRequest struct {
	Settings struct {
		Qos []struct {
			Model       string `json:"model,omitempty"`       //
			ProfileName string `json:"profileName,omitempty"` //
			WanProvider string `json:"wanProvider,omitempty"` //
		} `json:"qos,omitempty"` //
	} `json:"settings,omitempty"` //
}

// DNSServer is the DnsServer definition
type DNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         //
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   //
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` //
}

// HTTPSRead is the HttpsRead definition
type HTTPSRead struct {
	ID       string `json:"id,omitempty"`       //
	Name     string `json:"name,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Port     string `json:"port,omitempty"`     //
	Username string `json:"username,omitempty"` //
}

// HTTPSWrite is the HttpsWrite definition
type HTTPSWrite struct {
	ID       string `json:"id,omitempty"`       //
	Name     string `json:"name,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Port     string `json:"port,omitempty"`     //
	Username string `json:"username,omitempty"` //
}

// IPpool is the Ippool definition
type IPpool struct {
	DhcpServerIPs []string `json:"dhcpServerIps,omitempty"` //
	DNSServerIPs  []string `json:"dnsServerIps,omitempty"`  //
	Gateway       string   `json:"gateway,omitempty"`       //
	ID            string   `json:"id,omitempty"`            //
	IPPoolName    string   `json:"ipPoolName,omitempty"`    //
}

// MessageOfTheday is the MessageOfTheday definition
type MessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        //
	RetainExistingBanner bool   `json:"retainExistingBanner,omitempty"` //
}

// Netflowcollector is the Netflowcollector definition
type Netflowcollector struct {
	IPAddress string `json:"ipAddress,omitempty"` //
	Port      int    `json:"port,omitempty"`      //
}

// NetworkAAA is the Network_aaa definition
type NetworkAAA struct {
	IPAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// Qos is the Qos definition
type Qos struct {
	Model          string `json:"model,omitempty"`          //
	OldProfileName string `json:"oldProfileName,omitempty"` //
	ProfileName    string `json:"profileName,omitempty"`    //
	WanProvider    string `json:"wanProvider,omitempty"`    //
}

// SNMPServer is the SnmpServer definition
type SNMPServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// SNMPV2cRead is the SnmpV2cRead definition
type SNMPV2cRead struct {
	Description   string `json:"description,omitempty"`   //
	ID            string `json:"id,omitempty"`            //
	ReadCommunity string `json:"readCommunity,omitempty"` //
}

// SNMPV2cWrite is the SnmpV2cWrite definition
type SNMPV2cWrite struct {
	Description    string `json:"description,omitempty"`    //
	ID             string `json:"id,omitempty"`             //
	WriteCommunity string `json:"writeCommunity,omitempty"` //
}

// SNMPV3 is the SnmpV3 definition
type SNMPV3 struct {
	AuthPassword    string `json:"authPassword,omitempty"`    //
	AuthType        string `json:"authType,omitempty"`        //
	Description     string `json:"description,omitempty"`     //
	ID              string `json:"id,omitempty"`              //
	PrivacyPassword string `json:"privacyPassword,omitempty"` //
	PrivacyType     string `json:"privacyType,omitempty"`     //
	SNMPMode        string `json:"snmpMode,omitempty"`        //
	Username        string `json:"username,omitempty"`        //
}

// SyslogServer is the SyslogServer definition
type SyslogServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// UpdateDeviceCredentialsRequest is the UpdateDeviceCredentialsRequest definition
type UpdateDeviceCredentialsRequest struct {
	Settings struct {
		CliCredential struct {
			Description    string `json:"description,omitempty"`    //
			EnablePassword string `json:"enablePassword,omitempty"` //
			ID             string `json:"id,omitempty"`             //
			Password       string `json:"password,omitempty"`       //
			Username       string `json:"username,omitempty"`       //
		} `json:"cliCredential,omitempty"` //
		HTTPSRead struct {
			ID       string `json:"id,omitempty"`       //
			Name     string `json:"name,omitempty"`     //
			Password string `json:"password,omitempty"` //
			Port     string `json:"port,omitempty"`     //
			Username string `json:"username,omitempty"` //
		} `json:"httpsRead,omitempty"` //
		HTTPSWrite struct {
			ID       string `json:"id,omitempty"`       //
			Name     string `json:"name,omitempty"`     //
			Password string `json:"password,omitempty"` //
			Port     string `json:"port,omitempty"`     //
			Username string `json:"username,omitempty"` //
		} `json:"httpsWrite,omitempty"` //
		SNMPV2cRead struct {
			Description   string `json:"description,omitempty"`   //
			ID            string `json:"id,omitempty"`            //
			ReadCommunity string `json:"readCommunity,omitempty"` //
		} `json:"snmpV2cRead,omitempty"` //
		SNMPV2cWrite struct {
			Description    string `json:"description,omitempty"`    //
			ID             string `json:"id,omitempty"`             //
			WriteCommunity string `json:"writeCommunity,omitempty"` //
		} `json:"snmpV2cWrite,omitempty"` //
		SNMPV3 struct {
			AuthPassword    string `json:"authPassword,omitempty"`    //
			AuthType        string `json:"authType,omitempty"`        //
			Description     string `json:"description,omitempty"`     //
			ID              string `json:"id,omitempty"`              //
			PrivacyPassword string `json:"privacyPassword,omitempty"` //
			PrivacyType     string `json:"privacyType,omitempty"`     //
			SNMPMode        string `json:"snmpMode,omitempty"`        //
			Username        string `json:"username,omitempty"`        //
		} `json:"snmpV3,omitempty"` //
	} `json:"settings,omitempty"` //
}

// UpdateGlobalPoolRequest is the UpdateGlobalPoolRequest definition
type UpdateGlobalPoolRequest struct {
	Settings struct {
		IPpool []struct {
			DhcpServerIPs []string `json:"dhcpServerIps,omitempty"` //
			DNSServerIPs  []string `json:"dnsServerIps,omitempty"`  //
			Gateway       string   `json:"gateway,omitempty"`       //
			ID            string   `json:"id,omitempty"`            //
			IPPoolName    string   `json:"ipPoolName,omitempty"`    //
		} `json:"ippool,omitempty"` //
	} `json:"settings,omitempty"` //
}

// UpdateNetworkRequest is the UpdateNetworkRequest definition
type UpdateNetworkRequest struct {
	Settings struct {
		ClientAndEndpointAaa struct {
			IPAddress    string `json:"ipAddress,omitempty"`    //
			Network      string `json:"network,omitempty"`      //
			Protocol     string `json:"protocol,omitempty"`     //
			Servers      string `json:"servers,omitempty"`      //
			SharedSecret string `json:"sharedSecret,omitempty"` //
		} `json:"clientAndEndpoint_aaa,omitempty"` //
		DhcpServer []string `json:"dhcpServer,omitempty"` //
		DNSServer  struct {
			DomainName         string `json:"domainName,omitempty"`         //
			PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   //
			SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` //
		} `json:"dnsServer,omitempty"` //
		MessageOfTheday struct {
			BannerMessage        string `json:"bannerMessage,omitempty"`        //
			RetainExistingBanner bool   `json:"retainExistingBanner,omitempty"` //
		} `json:"messageOfTheday,omitempty"` //
		Netflowcollector struct {
			IPAddress string `json:"ipAddress,omitempty"` //
			Port      int    `json:"port,omitempty"`      //
		} `json:"netflowcollector,omitempty"` //
		NetworkAaa struct {
			IPAddress    string `json:"ipAddress,omitempty"`    //
			Network      string `json:"network,omitempty"`      //
			Protocol     string `json:"protocol,omitempty"`     //
			Servers      string `json:"servers,omitempty"`      //
			SharedSecret string `json:"sharedSecret,omitempty"` //
		} `json:"network_aaa,omitempty"` //
		NtpServer  []string `json:"ntpServer,omitempty"` //
		SNMPServer struct {
			ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
			IPAddresses     []string `json:"ipAddresses,omitempty"`     //
		} `json:"snmpServer,omitempty"` //
		SyslogServer struct {
			ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
			IPAddresses     []string `json:"ipAddresses,omitempty"`     //
		} `json:"syslogServer,omitempty"` //
		Timezone string `json:"timezone,omitempty"` //
	} `json:"settings,omitempty"` //
}

// UpdateSPProfileRequest is the UpdateSPProfileRequest definition
type UpdateSPProfileRequest struct {
	Settings struct {
		Qos []struct {
			Model          string `json:"model,omitempty"`          //
			OldProfileName string `json:"oldProfileName,omitempty"` //
			ProfileName    string `json:"profileName,omitempty"`    //
			WanProvider    string `json:"wanProvider,omitempty"`    //
		} `json:"qos,omitempty"` //
	} `json:"settings,omitempty"` //
}

// AssignCredentialToSiteResponse is the AssignCredentialToSiteResponse definition
type AssignCredentialToSiteResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// Cli is the Cli definition
type Cli struct {
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

// Context is the Context definition
type Context struct {
	ContextKey   string `json:"contextKey,omitempty"`   //
	ContextValue string `json:"contextValue,omitempty"` //
	Owner        string `json:"owner,omitempty"`        //
}

// CreateDeviceCredentialsResponse is the CreateDeviceCredentialsResponse definition
type CreateDeviceCredentialsResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateGlobalPoolResponse is the CreateGlobalPoolResponse definition
type CreateGlobalPoolResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateNetworkResponse is the CreateNetworkResponse definition
type CreateNetworkResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateSPProfileResponse is the CreateSPProfileResponse definition
type CreateSPProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteDeviceCredentialResponse is the DeleteDeviceCredentialResponse definition
type DeleteDeviceCredentialResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteGlobalIPPoolResponse is the DeleteGlobalIPPoolResponse definition
type DeleteGlobalIPPoolResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteSPProfileResponse is the DeleteSPProfileResponse definition
type DeleteSPProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetDeviceCredentialDetailsResponse is the GetDeviceCredentialDetailsResponse definition
type GetDeviceCredentialDetailsResponse struct {
	Cli []struct {
		Comments         string `json:"comments,omitempty"`         //
		CredentialType   string `json:"credentialType,omitempty"`   //
		Description      string `json:"description,omitempty"`      //
		EnablePassword   string `json:"enablePassword,omitempty"`   //
		ID               string `json:"id,omitempty"`               //
		InstanceTenantID string `json:"instanceTenantId,omitempty"` //
		InstanceUUID     string `json:"instanceUuid,omitempty"`     //
		Password         string `json:"password,omitempty"`         //
		Username         string `json:"username,omitempty"`         //
	} `json:"cli,omitempty"` //
	HTTPRead []struct {
		Comments         string `json:"comments,omitempty"`         //
		CredentialType   string `json:"credentialType,omitempty"`   //
		Description      string `json:"description,omitempty"`      //
		ID               string `json:"id,omitempty"`               //
		InstanceTenantID string `json:"instanceTenantId,omitempty"` //
		InstanceUUID     string `json:"instanceUuid,omitempty"`     //
		Password         string `json:"password,omitempty"`         //
		Port             string `json:"port,omitempty"`             //
		Secure           string `json:"secure,omitempty"`           //
		Username         string `json:"username,omitempty"`         //
	} `json:"http_read,omitempty"` //
	HTTPWrite []struct {
		Comments         string `json:"comments,omitempty"`         //
		CredentialType   string `json:"credentialType,omitempty"`   //
		Description      string `json:"description,omitempty"`      //
		ID               string `json:"id,omitempty"`               //
		InstanceTenantID string `json:"instanceTenantId,omitempty"` //
		InstanceUUID     string `json:"instanceUuid,omitempty"`     //
		Password         string `json:"password,omitempty"`         //
		Port             string `json:"port,omitempty"`             //
		Secure           string `json:"secure,omitempty"`           //
		Username         string `json:"username,omitempty"`         //
	} `json:"http_write,omitempty"` //
	SNMPV2Read []struct {
		Comments         string `json:"comments,omitempty"`         //
		CredentialType   string `json:"credentialType,omitempty"`   //
		Description      string `json:"description,omitempty"`      //
		ID               string `json:"id,omitempty"`               //
		InstanceTenantID string `json:"instanceTenantId,omitempty"` //
		InstanceUUID     string `json:"instanceUuid,omitempty"`     //
		ReadCommunity    string `json:"readCommunity,omitempty"`    //
	} `json:"snmp_v2_read,omitempty"` //
	SNMPV2Write []struct {
		Comments         string `json:"comments,omitempty"`         //
		CredentialType   string `json:"credentialType,omitempty"`   //
		Description      string `json:"description,omitempty"`      //
		ID               string `json:"id,omitempty"`               //
		InstanceTenantID string `json:"instanceTenantId,omitempty"` //
		InstanceUUID     string `json:"instanceUuid,omitempty"`     //
		WriteCommunity   string `json:"writeCommunity,omitempty"`   //
	} `json:"snmp_v2_write,omitempty"` //
	SNMPV3 []struct {
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
	} `json:"snmp_v3,omitempty"` //
}

// GetGlobalPoolResponse is the GetGlobalPoolResponse definition
type GetGlobalPoolResponse struct {
	Response []struct {
		ClientOptions         string `json:"clientOptions,omitempty"`         //
		ConfigureExternalDhcp string `json:"configureExternalDhcp,omitempty"` //
		Context               []struct {
			ContextKey   string `json:"contextKey,omitempty"`   //
			ContextValue string `json:"contextValue,omitempty"` //
			Owner        string `json:"owner,omitempty"`        //
		} `json:"context,omitempty"` //
		CreateTime          string   `json:"createTime,omitempty"`          //
		DhcpServerIPs       []string `json:"dhcpServerIps,omitempty"`       //
		DNSServerIPs        []string `json:"dnsServerIps,omitempty"`        //
		Gateways            []string `json:"gateways,omitempty"`            //
		ID                  string   `json:"id,omitempty"`                  //
		IPPoolCidr          string   `json:"ipPoolCidr,omitempty"`          //
		IPPoolName          string   `json:"ipPoolName,omitempty"`          //
		IPv6                string   `json:"ipv6,omitempty"`                //
		LastUpdateTime      int      `json:"lastUpdateTime,omitempty"`      //
		Overlapping         string   `json:"overlapping,omitempty"`         //
		Owner               string   `json:"owner,omitempty"`               //
		ParentUUID          string   `json:"parentUuid,omitempty"`          //
		Shared              string   `json:"shared,omitempty"`              //
		TotalIPAddressCount string   `json:"totalIpAddressCount,omitempty"` //
		UsedIPAddressCount  string   `json:"usedIpAddressCount,omitempty"`  //
		UsedPercentage      string   `json:"usedPercentage,omitempty"`      //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetNetworkResponse is the GetNetworkResponse definition
type GetNetworkResponse struct {
	Response []struct {
		GroupUUID          string `json:"groupUuid,omitempty"`          //
		InheritedGroupName string `json:"inheritedGroupName,omitempty"` //
		InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` //
		InstanceType       string `json:"instanceType,omitempty"`       //
		InstanceUUID       string `json:"instanceUuid,omitempty"`       //
		Key                string `json:"key,omitempty"`                //
		Namespace          string `json:"namespace,omitempty"`          //
		Type               string `json:"type,omitempty"`               //
		Value              []struct {
			ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
			IPAddresses     []string `json:"ipAddresses,omitempty"`     //
		} `json:"value,omitempty"` //
		Version int `json:"version,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetServiceProviderDetailsResponse is the GetServiceProviderDetailsResponse definition
type GetServiceProviderDetailsResponse struct {
	Response []struct {
		GroupUUID          string `json:"groupUuid,omitempty"`          //
		InheritedGroupName string `json:"inheritedGroupName,omitempty"` //
		InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` //
		InstanceType       string `json:"instanceType,omitempty"`       //
		InstanceUUID       string `json:"instanceUuid,omitempty"`       //
		Key                string `json:"key,omitempty"`                //
		Namespace          string `json:"namespace,omitempty"`          //
		Type               string `json:"type,omitempty"`               //
		Value              []struct {
			SLAProfileName string `json:"slaProfileName,omitempty"` //
			SpProfileName  string `json:"spProfileName,omitempty"`  //
			WanProvider    string `json:"wanProvider,omitempty"`    //
		} `json:"value,omitempty"` //
		Version string `json:"version,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// HTTPRead is the Http_read definition
type HTTPRead struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             string `json:"port,omitempty"`             //
	Secure           string `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}

// HTTPWrite is the Http_write definition
type HTTPWrite struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             string `json:"port,omitempty"`             //
	Secure           string `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}

// SNMPv2Read is the Snmp_v2_read definition
type SNMPv2Read struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	ReadCommunity    string `json:"readCommunity,omitempty"`    //
}

// SNMPv2Write is the Snmp_v2_write definition
type SNMPv2Write struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	WriteCommunity   string `json:"writeCommunity,omitempty"`   //
}

// SNMPv3 is the Snmp_v3 definition
type SNMPv3 struct {
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

// UpdateDeviceCredentialsResponse is the UpdateDeviceCredentialsResponse definition
type UpdateDeviceCredentialsResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateGlobalPoolResponse is the UpdateGlobalPoolResponse definition
type UpdateGlobalPoolResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateNetworkResponse is the UpdateNetworkResponse definition
type UpdateNetworkResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateSPProfileResponse is the UpdateSPProfileResponse definition
type UpdateSPProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// Value is the Value definition
type Value struct {
	SLAProfileName string `json:"slaProfileName,omitempty"` //
	SpProfileName  string `json:"spProfileName,omitempty"`  //
	WanProvider    string `json:"wanProvider,omitempty"`    //
}

// AssignCredentialToSite assignCredentialToSite
/* Assign Device Credential To Site
@param __persistbapioutput Persist bapi sync response
@param siteID site id to assign credential.
*/
func (s *NetworkSettingsService) AssignCredentialToSite(siteID string, assignCredentialToSiteRequest *AssignCredentialToSiteRequest) (*AssignCredentialToSiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/credential-to-site/{siteID}"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(assignCredentialToSiteRequest).
		SetResult(&AssignCredentialToSiteResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*AssignCredentialToSiteResponse)
	return result, response, err

}

// CreateDeviceCredentials createDeviceCredentials
/* API to create device credentials.
 */
func (s *NetworkSettingsService) CreateDeviceCredentials(createDeviceCredentialsRequest *CreateDeviceCredentialsRequest) (*CreateDeviceCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-credential"

	response, err := RestyClient.R().
		SetBody(createDeviceCredentialsRequest).
		SetResult(&CreateDeviceCredentialsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateDeviceCredentialsResponse)
	return result, response, err

}

// CreateGlobalPool createGlobalPool
/* API to create global pool.
@param __persistbapioutput Persist bapi sync response
*/
func (s *NetworkSettingsService) CreateGlobalPool(createGlobalPoolRequest *CreateGlobalPoolRequest) (*CreateGlobalPoolResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-pool"

	response, err := RestyClient.R().
		SetBody(createGlobalPoolRequest).
		SetResult(&CreateGlobalPoolResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateGlobalPoolResponse)
	return result, response, err

}

// CreateNetwork createNetwork
/* API to create a network for DHCP and DNS center server settings.
@param __persistbapioutput Persist bapi sync response
@param siteID Site id to which site details to associate with the network settings.
*/
func (s *NetworkSettingsService) CreateNetwork(siteID string, createNetworkRequest *CreateNetworkRequest) (*CreateNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network/{siteID}"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(createNetworkRequest).
		SetResult(&CreateNetworkResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateNetworkResponse)
	return result, response, err

}

// CreateSPProfile createSPProfile
/* API to create service provider profile(QOS).
 */
func (s *NetworkSettingsService) CreateSPProfile(createSPProfileRequest *CreateSPProfileRequest) (*CreateSPProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/service-provider"

	response, err := RestyClient.R().
		SetBody(createSPProfileRequest).
		SetResult(&CreateSPProfileResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateSPProfileResponse)
	return result, response, err

}

// DeleteDeviceCredential deleteDeviceCredential
/* Delete device credential.
@param id global credential id
*/
func (s *NetworkSettingsService) DeleteDeviceCredential(id string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/device-credential/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteGlobalIPPool deleteGlobalIPPool
/* API to delete global IP pool.
@param id global pool id
*/
func (s *NetworkSettingsService) DeleteGlobalIPPool(id string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/global-pool/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteSPProfile deleteSPProfile
/* API to delete Service Provider profile (QoS).
@param spProfileName sp profile name
*/
func (s *NetworkSettingsService) DeleteSPProfile(spProfileName string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/sp-profile/{spProfileName}"
	path = strings.Replace(path, "{"+"spProfileName"+"}", fmt.Sprintf("%v", spProfileName), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetDeviceCredentialDetailsQueryParams defines the query parameters for this request
type GetDeviceCredentialDetailsQueryParams struct {
	SiteID string `url:"siteId,omitempty"` // Site id to retrieve the credential details associated with the site.
}

// GetDeviceCredentialDetails getDeviceCredentialDetails
/* API to get device credential details.
@param siteID Site id to retrieve the credential details associated with the site.
*/
func (s *NetworkSettingsService) GetDeviceCredentialDetails(getDeviceCredentialDetailsQueryParams *GetDeviceCredentialDetailsQueryParams) (*GetDeviceCredentialDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-credential"

	queryString, _ := query.Values(getDeviceCredentialDetailsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceCredentialDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDeviceCredentialDetailsResponse)
	return result, response, err

}

// GetGlobalPoolQueryParams defines the query parameters for this request
type GetGlobalPoolQueryParams struct {
	Offset string `url:"offset,omitempty"` // offset/starting row
	Limit  string `url:"limit,omitempty"`  // No of Global Pools to be retrieved
}

// GetGlobalPool getGlobalPool
/* API to get global pool.
@param offset offset/starting row
@param limit No of Global Pools to be retrieved
*/
func (s *NetworkSettingsService) GetGlobalPool(getGlobalPoolQueryParams *GetGlobalPoolQueryParams) (*GetGlobalPoolResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-pool"

	queryString, _ := query.Values(getGlobalPoolQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetGlobalPoolResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetGlobalPoolResponse)
	return result, response, err

}

// GetNetworkQueryParams defines the query parameters for this request
type GetNetworkQueryParams struct {
	SiteID string `url:"siteId,omitempty"` // Site id to get the network settings associated with the site.
}

// GetNetwork getNetwork
/* API to get  DHCP and DNS center server details.
@param siteID Site id to get the network settings associated with the site.
*/
func (s *NetworkSettingsService) GetNetwork(getNetworkQueryParams *GetNetworkQueryParams) (*GetNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network"

	queryString, _ := query.Values(getNetworkQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetNetworkResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetNetworkResponse)
	return result, response, err

}

// GetServiceProviderDetails getServiceProviderDetails
/* API to get service provider details (QoS).
 */
func (s *NetworkSettingsService) GetServiceProviderDetails() (*GetServiceProviderDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/service-provider"

	response, err := RestyClient.R().
		SetResult(&GetServiceProviderDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetServiceProviderDetailsResponse)
	return result, response, err

}

// UpdateDeviceCredentials updateDeviceCredentials
/* API to update device credentials.
 */
func (s *NetworkSettingsService) UpdateDeviceCredentials(updateDeviceCredentialsRequest *UpdateDeviceCredentialsRequest) (*UpdateDeviceCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-credential"

	response, err := RestyClient.R().
		SetBody(updateDeviceCredentialsRequest).
		SetResult(&UpdateDeviceCredentialsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateDeviceCredentialsResponse)
	return result, response, err

}

// UpdateGlobalPool updateGlobalPool
/* API to update global pool
 */
func (s *NetworkSettingsService) UpdateGlobalPool(updateGlobalPoolRequest *UpdateGlobalPoolRequest) (*UpdateGlobalPoolResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-pool"

	response, err := RestyClient.R().
		SetBody(updateGlobalPoolRequest).
		SetResult(&UpdateGlobalPoolResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateGlobalPoolResponse)
	return result, response, err

}

// UpdateNetwork updateNetwork
/* API to update network for DHCP and DNS center server settings.
@param __persistbapioutput Persist bapi sync response
@param siteID Site id to update the network settings which is associated with the site
*/
func (s *NetworkSettingsService) UpdateNetwork(siteID string, updateNetworkRequest *UpdateNetworkRequest) (*UpdateNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network/{siteID}"
	path = strings.Replace(path, "{"+"siteID"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(updateNetworkRequest).
		SetResult(&UpdateNetworkResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateNetworkResponse)
	return result, response, err

}

// UpdateSPProfile updateSPProfile
/* API to update SP profile
 */
func (s *NetworkSettingsService) UpdateSPProfile(updateSPProfileRequest *UpdateSPProfileRequest) (*UpdateSPProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/service-provider"

	response, err := RestyClient.R().
		SetBody(updateSPProfileRequest).
		SetResult(&UpdateSPProfileResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateSPProfileResponse)
	return result, response, err

}
