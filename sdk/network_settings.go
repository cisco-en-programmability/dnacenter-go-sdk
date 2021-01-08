package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// NetworkSettingsService is the service to communicate with the NetworkSettings API endpoint
type NetworkSettingsService service

// AssignCredentialToSiteRequest is the assignCredentialToSiteRequest definition
type AssignCredentialToSiteRequest struct {
	CliID         string `json:"cliId,omitempty"`         //
	HTTPRead      string `json:"httpRead,omitempty"`      //
	HTTPWrite     string `json:"httpWrite,omitempty"`     //
	SNMPV2ReadID  string `json:"snmpV2ReadId,omitempty"`  //
	SNMPV2WriteID string `json:"snmpV2WriteId,omitempty"` //
	SNMPV3ID      string `json:"snmpV3Id,omitempty"`      //
}

// CreateDeviceCredentialsRequest is the createDeviceCredentialsRequest definition
type CreateDeviceCredentialsRequest struct {
	Settings CreateDeviceCredentialsRequestSettings `json:"settings,omitempty"` //
}

// CreateDeviceCredentialsRequestSettings is the createDeviceCredentialsRequestSettings definition
type CreateDeviceCredentialsRequestSettings struct {
	CliCredential []CreateDeviceCredentialsRequestSettingsCliCredential `json:"cliCredential,omitempty"` //
	HTTPSRead     []CreateDeviceCredentialsRequestSettingsHTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    []CreateDeviceCredentialsRequestSettingsHTTPSWrite    `json:"httpsWrite,omitempty"`    //
	SNMPV2cRead   []CreateDeviceCredentialsRequestSettingsSNMPV2cRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2cWrite  []CreateDeviceCredentialsRequestSettingsSNMPV2cWrite  `json:"snmpV2cWrite,omitempty"`  //
	SNMPV3        []CreateDeviceCredentialsRequestSettingsSNMPV3        `json:"snmpV3,omitempty"`        //
}

// CreateDeviceCredentialsRequestSettingsCliCredential is the createDeviceCredentialsRequestSettingsCliCredential definition
type CreateDeviceCredentialsRequestSettingsCliCredential struct {
	Description    string `json:"description,omitempty"`    //
	EnablePassword string `json:"enablePassword,omitempty"` //
	Password       string `json:"password,omitempty"`       //
	Username       string `json:"username,omitempty"`       //
}

// CreateDeviceCredentialsRequestSettingsHTTPSRead is the createDeviceCredentialsRequestSettingsHTTPSRead definition
type CreateDeviceCredentialsRequestSettingsHTTPSRead struct {
	Name     string  `json:"name,omitempty"`     //
	Password string  `json:"password,omitempty"` //
	Port     float64 `json:"port,omitempty"`     //
	Username string  `json:"username,omitempty"` //
}

// CreateDeviceCredentialsRequestSettingsHTTPSWrite is the createDeviceCredentialsRequestSettingsHTTPSWrite definition
type CreateDeviceCredentialsRequestSettingsHTTPSWrite struct {
	Name     string  `json:"name,omitempty"`     //
	Password string  `json:"password,omitempty"` //
	Port     float64 `json:"port,omitempty"`     //
	Username string  `json:"username,omitempty"` //
}

// CreateDeviceCredentialsRequestSettingsSNMPV2cRead is the createDeviceCredentialsRequestSettingsSNMPV2cRead definition
type CreateDeviceCredentialsRequestSettingsSNMPV2cRead struct {
	Description   string `json:"description,omitempty"`   //
	ReadCommunity string `json:"readCommunity,omitempty"` //
}

// CreateDeviceCredentialsRequestSettingsSNMPV2cWrite is the createDeviceCredentialsRequestSettingsSNMPV2cWrite definition
type CreateDeviceCredentialsRequestSettingsSNMPV2cWrite struct {
	Description    string `json:"description,omitempty"`    //
	WriteCommunity string `json:"writeCommunity,omitempty"` //
}

// CreateDeviceCredentialsRequestSettingsSNMPV3 is the createDeviceCredentialsRequestSettingsSNMPV3 definition
type CreateDeviceCredentialsRequestSettingsSNMPV3 struct {
	AuthPassword    string `json:"authPassword,omitempty"`    //
	AuthType        string `json:"authType,omitempty"`        //
	Description     string `json:"description,omitempty"`     //
	PrivacyPassword string `json:"privacyPassword,omitempty"` //
	PrivacyType     string `json:"privacyType,omitempty"`     //
	SNMPMode        string `json:"snmpMode,omitempty"`        //
	Username        string `json:"username,omitempty"`        //
}

// CreateGlobalPoolRequest is the createGlobalPoolRequest definition
type CreateGlobalPoolRequest struct {
	Settings CreateGlobalPoolRequestSettings `json:"settings,omitempty"` //
}

// CreateGlobalPoolRequestSettings is the createGlobalPoolRequestSettings definition
type CreateGlobalPoolRequestSettings struct {
	IPpool []CreateGlobalPoolRequestSettingsIPpool `json:"ippool,omitempty"` //
}

// CreateGlobalPoolRequestSettingsIPpool is the createGlobalPoolRequestSettingsIPpool definition
type CreateGlobalPoolRequestSettingsIPpool struct {
	IPAddressSpace string   `json:"IpAddressSpace,omitempty"` //
	DhcpServerIPs  []string `json:"dhcpServerIps,omitempty"`  //
	DNSServerIPs   []string `json:"dnsServerIps,omitempty"`   //
	Gateway        string   `json:"gateway,omitempty"`        //
	IPPoolCidr     string   `json:"ipPoolCidr,omitempty"`     //
	IPPoolName     string   `json:"ipPoolName,omitempty"`     //
	Type           string   `json:"type,omitempty"`           //
}

// CreateGlobalPoolRequestSettingsIPpoolDNSServerIPs is the createGlobalPoolRequestSettingsIPpoolDNSServerIPs definition
type CreateGlobalPoolRequestSettingsIPpoolDNSServerIPs []string

// CreateGlobalPoolRequestSettingsIPpoolDhcpServerIPs is the createGlobalPoolRequestSettingsIPpoolDhcpServerIPs definition
type CreateGlobalPoolRequestSettingsIPpoolDhcpServerIPs []string

// CreateNetworkRequest is the createNetworkRequest definition
type CreateNetworkRequest struct {
	Settings CreateNetworkRequestSettings `json:"settings,omitempty"` //
}

// CreateNetworkRequestSettings is the createNetworkRequestSettings definition
type CreateNetworkRequestSettings struct {
	ClientAndEndpointAAA CreateNetworkRequestSettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
	DhcpServer           []string                                         `json:"dhcpServer,omitempty"`            //
	DNSServer            CreateNetworkRequestSettingsDNSServer            `json:"dnsServer,omitempty"`             //
	MessageOfTheday      CreateNetworkRequestSettingsMessageOfTheday      `json:"messageOfTheday,omitempty"`       //
	Netflowcollector     CreateNetworkRequestSettingsNetflowcollector     `json:"netflowcollector,omitempty"`      //
	NetworkAAA           CreateNetworkRequestSettingsNetworkAAA           `json:"network_aaa,omitempty"`           //
	NtpServer            []string                                         `json:"ntpServer,omitempty"`             //
	SNMPServer           CreateNetworkRequestSettingsSNMPServer           `json:"snmpServer,omitempty"`            //
	SyslogServer         CreateNetworkRequestSettingsSyslogServer         `json:"syslogServer,omitempty"`          //
	Timezone             string                                           `json:"timezone,omitempty"`              //
}

// CreateNetworkRequestSettingsClientAndEndpointAAA is the createNetworkRequestSettingsClientAndEndpointAAA definition
type CreateNetworkRequestSettingsClientAndEndpointAAA struct {
	IPAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// CreateNetworkRequestSettingsDNSServer is the createNetworkRequestSettingsDNSServer definition
type CreateNetworkRequestSettingsDNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         //
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   //
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` //
}

// CreateNetworkRequestSettingsDhcpServer is the createNetworkRequestSettingsDhcpServer definition
type CreateNetworkRequestSettingsDhcpServer []string

// CreateNetworkRequestSettingsMessageOfTheday is the createNetworkRequestSettingsMessageOfTheday definition
type CreateNetworkRequestSettingsMessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        //
	RetainExistingBanner bool   `json:"retainExistingBanner,omitempty"` //
}

// CreateNetworkRequestSettingsNetflowcollector is the createNetworkRequestSettingsNetflowcollector definition
type CreateNetworkRequestSettingsNetflowcollector struct {
	IPAddress string  `json:"ipAddress,omitempty"` //
	Port      float64 `json:"port,omitempty"`      //
}

// CreateNetworkRequestSettingsNetworkAAA is the createNetworkRequestSettingsNetworkAAA definition
type CreateNetworkRequestSettingsNetworkAAA struct {
	IPAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// CreateNetworkRequestSettingsNtpServer is the createNetworkRequestSettingsNtpServer definition
type CreateNetworkRequestSettingsNtpServer []string

// CreateNetworkRequestSettingsSNMPServer is the createNetworkRequestSettingsSNMPServer definition
type CreateNetworkRequestSettingsSNMPServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// CreateNetworkRequestSettingsSNMPServerIPAddresses is the createNetworkRequestSettingsSNMPServerIPAddresses definition
type CreateNetworkRequestSettingsSNMPServerIPAddresses []string

// CreateNetworkRequestSettingsSyslogServer is the createNetworkRequestSettingsSyslogServer definition
type CreateNetworkRequestSettingsSyslogServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// CreateNetworkRequestSettingsSyslogServerIPAddresses is the createNetworkRequestSettingsSyslogServerIPAddresses definition
type CreateNetworkRequestSettingsSyslogServerIPAddresses []string

// CreateSPProfileRequest is the createSPProfileRequest definition
type CreateSPProfileRequest struct {
	Settings CreateSPProfileRequestSettings `json:"settings,omitempty"` //
}

// CreateSPProfileRequestSettings is the createSPProfileRequestSettings definition
type CreateSPProfileRequestSettings struct {
	Qos []CreateSPProfileRequestSettingsQos `json:"qos,omitempty"` //
}

// CreateSPProfileRequestSettingsQos is the createSPProfileRequestSettingsQos definition
type CreateSPProfileRequestSettingsQos struct {
	Model       string `json:"model,omitempty"`       //
	ProfileName string `json:"profileName,omitempty"` //
	WanProvider string `json:"wanProvider,omitempty"` //
}

// UpdateDeviceCredentialsRequest is the updateDeviceCredentialsRequest definition
type UpdateDeviceCredentialsRequest struct {
	Settings UpdateDeviceCredentialsRequestSettings `json:"settings,omitempty"` //
}

// UpdateDeviceCredentialsRequestSettings is the updateDeviceCredentialsRequestSettings definition
type UpdateDeviceCredentialsRequestSettings struct {
	CliCredential UpdateDeviceCredentialsRequestSettingsCliCredential `json:"cliCredential,omitempty"` //
	HTTPSRead     UpdateDeviceCredentialsRequestSettingsHTTPSRead     `json:"httpsRead,omitempty"`     //
	HTTPSWrite    UpdateDeviceCredentialsRequestSettingsHTTPSWrite    `json:"httpsWrite,omitempty"`    //
	SNMPV2cRead   UpdateDeviceCredentialsRequestSettingsSNMPV2cRead   `json:"snmpV2cRead,omitempty"`   //
	SNMPV2cWrite  UpdateDeviceCredentialsRequestSettingsSNMPV2cWrite  `json:"snmpV2cWrite,omitempty"`  //
	SNMPV3        UpdateDeviceCredentialsRequestSettingsSNMPV3        `json:"snmpV3,omitempty"`        //
}

// UpdateDeviceCredentialsRequestSettingsCliCredential is the updateDeviceCredentialsRequestSettingsCliCredential definition
type UpdateDeviceCredentialsRequestSettingsCliCredential struct {
	Description    string `json:"description,omitempty"`    //
	EnablePassword string `json:"enablePassword,omitempty"` //
	ID             string `json:"id,omitempty"`             //
	Password       string `json:"password,omitempty"`       //
	Username       string `json:"username,omitempty"`       //
}

// UpdateDeviceCredentialsRequestSettingsHTTPSRead is the updateDeviceCredentialsRequestSettingsHTTPSRead definition
type UpdateDeviceCredentialsRequestSettingsHTTPSRead struct {
	ID       string `json:"id,omitempty"`       //
	Name     string `json:"name,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Port     string `json:"port,omitempty"`     //
	Username string `json:"username,omitempty"` //
}

// UpdateDeviceCredentialsRequestSettingsHTTPSWrite is the updateDeviceCredentialsRequestSettingsHTTPSWrite definition
type UpdateDeviceCredentialsRequestSettingsHTTPSWrite struct {
	ID       string `json:"id,omitempty"`       //
	Name     string `json:"name,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Port     string `json:"port,omitempty"`     //
	Username string `json:"username,omitempty"` //
}

// UpdateDeviceCredentialsRequestSettingsSNMPV2cRead is the updateDeviceCredentialsRequestSettingsSNMPV2cRead definition
type UpdateDeviceCredentialsRequestSettingsSNMPV2cRead struct {
	Description   string `json:"description,omitempty"`   //
	ID            string `json:"id,omitempty"`            //
	ReadCommunity string `json:"readCommunity,omitempty"` //
}

// UpdateDeviceCredentialsRequestSettingsSNMPV2cWrite is the updateDeviceCredentialsRequestSettingsSNMPV2cWrite definition
type UpdateDeviceCredentialsRequestSettingsSNMPV2cWrite struct {
	Description    string `json:"description,omitempty"`    //
	ID             string `json:"id,omitempty"`             //
	WriteCommunity string `json:"writeCommunity,omitempty"` //
}

// UpdateDeviceCredentialsRequestSettingsSNMPV3 is the updateDeviceCredentialsRequestSettingsSNMPV3 definition
type UpdateDeviceCredentialsRequestSettingsSNMPV3 struct {
	AuthPassword    string `json:"authPassword,omitempty"`    //
	AuthType        string `json:"authType,omitempty"`        //
	Description     string `json:"description,omitempty"`     //
	ID              string `json:"id,omitempty"`              //
	PrivacyPassword string `json:"privacyPassword,omitempty"` //
	PrivacyType     string `json:"privacyType,omitempty"`     //
	SNMPMode        string `json:"snmpMode,omitempty"`        //
	Username        string `json:"username,omitempty"`        //
}

// UpdateGlobalPoolRequest is the updateGlobalPoolRequest definition
type UpdateGlobalPoolRequest struct {
	Settings UpdateGlobalPoolRequestSettings `json:"settings,omitempty"` //
}

// UpdateGlobalPoolRequestSettings is the updateGlobalPoolRequestSettings definition
type UpdateGlobalPoolRequestSettings struct {
	IPpool []UpdateGlobalPoolRequestSettingsIPpool `json:"ippool,omitempty"` //
}

// UpdateGlobalPoolRequestSettingsIPpool is the updateGlobalPoolRequestSettingsIPpool definition
type UpdateGlobalPoolRequestSettingsIPpool struct {
	DhcpServerIPs *[]string `json:"dhcpServerIps,omitempty"` //
	DNSServerIPs  *[]string `json:"dnsServerIps,omitempty"`  //
	Gateway       string    `json:"gateway,omitempty"`       //
	ID            string    `json:"id,omitempty"`            //
	IPPoolName    string    `json:"ipPoolName,omitempty"`    //
}

// UpdateGlobalPoolRequestSettingsIPpoolDNSServerIPs is the updateGlobalPoolRequestSettingsIPpoolDNSServerIPs definition
type UpdateGlobalPoolRequestSettingsIPpoolDNSServerIPs []string

// UpdateGlobalPoolRequestSettingsIPpoolDhcpServerIPs is the updateGlobalPoolRequestSettingsIPpoolDhcpServerIPs definition
type UpdateGlobalPoolRequestSettingsIPpoolDhcpServerIPs []string

// UpdateNetworkRequest is the updateNetworkRequest definition
type UpdateNetworkRequest struct {
	Settings UpdateNetworkRequestSettings `json:"settings,omitempty"` //
}

// UpdateNetworkRequestSettings is the updateNetworkRequestSettings definition
type UpdateNetworkRequestSettings struct {
	ClientAndEndpointAAA *UpdateNetworkRequestSettingsClientAndEndpointAAA `json:"clientAndEndpoint_aaa,omitempty"` //
	DhcpServer           []string                                          `json:"dhcpServer,omitempty"`            //
	DNSServer            *UpdateNetworkRequestSettingsDNSServer            `json:"dnsServer,omitempty"`             //
	MessageOfTheday      *UpdateNetworkRequestSettingsMessageOfTheday      `json:"messageOfTheday,omitempty"`       //
	Netflowcollector     *UpdateNetworkRequestSettingsNetflowcollector     `json:"netflowcollector,omitempty"`      //
	NetworkAAA           *UpdateNetworkRequestSettingsNetworkAAA           `json:"network_aaa,omitempty"`           //
	NtpServer            []string                                          `json:"ntpServer,omitempty"`             //
	SNMPServer           *UpdateNetworkRequestSettingsSNMPServer           `json:"snmpServer,omitempty"`            //
	SyslogServer         *UpdateNetworkRequestSettingsSyslogServer         `json:"syslogServer,omitempty"`          //
	Timezone             string                                            `json:"timezone,omitempty"`              //
}

// UpdateNetworkRequestSettingsClientAndEndpointAAA is the updateNetworkRequestSettingsClientAndEndpointAAA definition
type UpdateNetworkRequestSettingsClientAndEndpointAAA struct {
	IPAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// UpdateNetworkRequestSettingsDNSServer is the updateNetworkRequestSettingsDNSServer definition
type UpdateNetworkRequestSettingsDNSServer struct {
	DomainName         string `json:"domainName,omitempty"`         //
	PrimaryIPAddress   string `json:"primaryIpAddress,omitempty"`   //
	SecondaryIPAddress string `json:"secondaryIpAddress,omitempty"` //
}

// UpdateNetworkRequestSettingsDhcpServer is the updateNetworkRequestSettingsDhcpServer definition
type UpdateNetworkRequestSettingsDhcpServer []string

// UpdateNetworkRequestSettingsMessageOfTheday is the updateNetworkRequestSettingsMessageOfTheday definition
type UpdateNetworkRequestSettingsMessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        //
	RetainExistingBanner bool   `json:"retainExistingBanner,omitempty"` //
}

// UpdateNetworkRequestSettingsNetflowcollector is the updateNetworkRequestSettingsNetflowcollector definition
type UpdateNetworkRequestSettingsNetflowcollector struct {
	IPAddress string  `json:"ipAddress,omitempty"` //
	Port      float64 `json:"port,omitempty"`      //
}

// UpdateNetworkRequestSettingsNetworkAAA is the updateNetworkRequestSettingsNetworkAAA definition
type UpdateNetworkRequestSettingsNetworkAAA struct {
	IPAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// UpdateNetworkRequestSettingsNtpServer is the updateNetworkRequestSettingsNtpServer definition
type UpdateNetworkRequestSettingsNtpServer []string

// UpdateNetworkRequestSettingsSNMPServer is the updateNetworkRequestSettingsSNMPServer definition
type UpdateNetworkRequestSettingsSNMPServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// UpdateNetworkRequestSettingsSNMPServerIPAddresses is the updateNetworkRequestSettingsSNMPServerIPAddresses definition
type UpdateNetworkRequestSettingsSNMPServerIPAddresses []string

// UpdateNetworkRequestSettingsSyslogServer is the updateNetworkRequestSettingsSyslogServer definition
type UpdateNetworkRequestSettingsSyslogServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// UpdateNetworkRequestSettingsSyslogServerIPAddresses is the updateNetworkRequestSettingsSyslogServerIPAddresses definition
type UpdateNetworkRequestSettingsSyslogServerIPAddresses []string

// UpdateSPProfileRequest is the updateSPProfileRequest definition
type UpdateSPProfileRequest struct {
	Settings UpdateSPProfileRequestSettings `json:"settings,omitempty"` //
}

// UpdateSPProfileRequestSettings is the updateSPProfileRequestSettings definition
type UpdateSPProfileRequestSettings struct {
	Qos []UpdateSPProfileRequestSettingsQos `json:"qos,omitempty"` //
}

// UpdateSPProfileRequestSettingsQos is the updateSPProfileRequestSettingsQos definition
type UpdateSPProfileRequestSettingsQos struct {
	Model          string `json:"model,omitempty"`          //
	OldProfileName string `json:"oldProfileName,omitempty"` //
	ProfileName    string `json:"profileName,omitempty"`    //
	WanProvider    string `json:"wanProvider,omitempty"`    //
}

// AssignCredentialToSiteResponse is the assignCredentialToSiteResponse definition
type AssignCredentialToSiteResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateDeviceCredentialsResponse is the createDeviceCredentialsResponse definition
type CreateDeviceCredentialsResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateGlobalPoolResponse is the createGlobalPoolResponse definition
type CreateGlobalPoolResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateNetworkResponse is the createNetworkResponse definition
type CreateNetworkResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateSPProfileResponse is the createSPProfileResponse definition
type CreateSPProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteDeviceCredentialResponse is the deleteDeviceCredentialResponse definition
type DeleteDeviceCredentialResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteGlobalIPPoolResponse is the deleteGlobalIPPoolResponse definition
type DeleteGlobalIPPoolResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteSPProfileResponse is the deleteSPProfileResponse definition
type DeleteSPProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetDeviceCredentialDetailsResponse is the getDeviceCredentialDetailsResponse definition
type GetDeviceCredentialDetailsResponse struct {
	Cli         []GetDeviceCredentialDetailsResponseCli         `json:"cli,omitempty"`           //
	HTTPRead    []GetDeviceCredentialDetailsResponseHTTPRead    `json:"http_read,omitempty"`     //
	HTTPWrite   []GetDeviceCredentialDetailsResponseHTTPWrite   `json:"http_write,omitempty"`    //
	SNMPV2Read  []GetDeviceCredentialDetailsResponseSNMPv2Read  `json:"snmp_v2_read,omitempty"`  //
	SNMPV2Write []GetDeviceCredentialDetailsResponseSNMPv2Write `json:"snmp_v2_write,omitempty"` //
	SNMPV3      []GetDeviceCredentialDetailsResponseSNMPv3      `json:"snmp_v3,omitempty"`       //
}

// GetDeviceCredentialDetailsResponseCli is the getDeviceCredentialDetailsResponseCli definition
type GetDeviceCredentialDetailsResponseCli struct {
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

// GetDeviceCredentialDetailsResponseHTTPRead is the getDeviceCredentialDetailsResponseHTTPRead definition
type GetDeviceCredentialDetailsResponseHTTPRead struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             int    `json:"port,omitempty"`             //
	Secure           bool   `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}

// GetDeviceCredentialDetailsResponseHTTPWrite is the getDeviceCredentialDetailsResponseHTTPWrite definition
type GetDeviceCredentialDetailsResponseHTTPWrite struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             int    `json:"port,omitempty"`             //
	Secure           bool   `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}

// GetDeviceCredentialDetailsResponseSNMPv2Read is the getDeviceCredentialDetailsResponseSNMPv2Read definition
type GetDeviceCredentialDetailsResponseSNMPv2Read struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	ReadCommunity    string `json:"readCommunity,omitempty"`    //
}

// GetDeviceCredentialDetailsResponseSNMPv2Write is the getDeviceCredentialDetailsResponseSNMPv2Write definition
type GetDeviceCredentialDetailsResponseSNMPv2Write struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	WriteCommunity   string `json:"writeCommunity,omitempty"`   //
}

// GetDeviceCredentialDetailsResponseSNMPv3 is the getDeviceCredentialDetailsResponseSNMPv3 definition
type GetDeviceCredentialDetailsResponseSNMPv3 struct {
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

// GetGlobalPoolResponse is the getGlobalPoolResponse definition
type GetGlobalPoolResponse struct {
	Response []GetGlobalPoolResponseResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}

// GetGlobalPoolResponseResponse is the getGlobalPoolResponseResponse definition
type GetGlobalPoolResponseResponse struct {
	ClientOptions         interface{}                            `json:"clientOptions,omitempty"`         //
	ConfigureExternalDhcp bool                                   `json:"configureExternalDhcp,omitempty"` //
	Context               []GetGlobalPoolResponseResponseContext `json:"context,omitempty"`               //
	CreateTime            int                                    `json:"createTime,omitempty"`            //
	DhcpServerIPs         []string                               `json:"dhcpServerIps,omitempty"`         //
	DNSServerIPs          []string                               `json:"dnsServerIps,omitempty"`          //
	Gateways              []string                               `json:"gateways,omitempty"`              //
	ID                    string                                 `json:"id,omitempty"`                    //
	IPPoolCidr            string                                 `json:"ipPoolCidr,omitempty"`            //
	IPPoolName            string                                 `json:"ipPoolName,omitempty"`            //
	IPv6                  bool                                   `json:"ipv6,omitempty"`                  //
	LastUpdateTime        int                                    `json:"lastUpdateTime,omitempty"`        //
	Overlapping           bool                                   `json:"overlapping,omitempty"`           //
	Owner                 string                                 `json:"owner,omitempty"`                 //
	ParentUUID            string                                 `json:"parentUuid,omitempty"`            //
	Shared                bool                                   `json:"shared,omitempty"`                //
	TotalIPAddressCount   int                                    `json:"totalIpAddressCount,omitempty"`   //
	UsedIPAddressCount    int                                    `json:"usedIpAddressCount,omitempty"`    //
	UsedPercentage        string                                 `json:"usedPercentage,omitempty"`        //
}

// GetGlobalPoolResponseResponseContext is the getGlobalPoolResponseResponseContext definition
type GetGlobalPoolResponseResponseContext struct {
	ContextKey   string `json:"contextKey,omitempty"`   //
	ContextValue string `json:"contextValue,omitempty"` //
	Owner        string `json:"owner,omitempty"`        //
}

// GetGlobalPoolResponseResponseDNSServerIPs is the getGlobalPoolResponseResponseDNSServerIPs definition
type GetGlobalPoolResponseResponseDNSServerIPs []string

// GetGlobalPoolResponseResponseDhcpServerIPs is the getGlobalPoolResponseResponseDhcpServerIPs definition
type GetGlobalPoolResponseResponseDhcpServerIPs []string

// GetGlobalPoolResponseResponseGateways is the getGlobalPoolResponseResponseGateways definition
type GetGlobalPoolResponseResponseGateways []string

// GetNetworkResponse is the getNetworkResponse definition
type GetNetworkResponse struct {
	Response []GetNetworkResponseResponse `json:"response,omitempty"` //
	Version  string                       `json:"version,omitempty"`  //
}

// GetNetworkResponseResponse is the getNetworkResponseResponse definition
type GetNetworkResponseResponse struct {
	GroupUUID          string        `json:"groupUuid,omitempty"`          //
	InheritedGroupName string        `json:"inheritedGroupName,omitempty"` //
	InheritedGroupUUID string        `json:"inheritedGroupUuid,omitempty"` //
	InstanceType       string        `json:"instanceType,omitempty"`       //
	InstanceUUID       string        `json:"instanceUuid,omitempty"`       //
	Key                string        `json:"key,omitempty"`                //
	Namespace          string        `json:"namespace,omitempty"`          //
	Type               string        `json:"type,omitempty"`               //
	Value              []interface{} `json:"value,omitempty"`              //
	Version            int           `json:"version,omitempty"`            //
}

// GetNetworkResponseResponseValue is the getNetworkResponseResponseValue definition
type GetNetworkResponseResponseValue struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IPAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// GetNetworkResponseResponseValueIPAddresses is the getNetworkResponseResponseValueIPAddresses definition
type GetNetworkResponseResponseValueIPAddresses []string

// GetServiceProviderDetailsResponse is the getServiceProviderDetailsResponse definition
type GetServiceProviderDetailsResponse struct {
	Response []GetServiceProviderDetailsResponseResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}

// GetServiceProviderDetailsResponseResponse is the getServiceProviderDetailsResponseResponse definition
type GetServiceProviderDetailsResponseResponse struct {
	GroupUUID          string                                           `json:"groupUuid,omitempty"`          //
	InheritedGroupName string                                           `json:"inheritedGroupName,omitempty"` //
	InheritedGroupUUID string                                           `json:"inheritedGroupUuid,omitempty"` //
	InstanceType       string                                           `json:"instanceType,omitempty"`       //
	InstanceUUID       string                                           `json:"instanceUuid,omitempty"`       //
	Key                string                                           `json:"key,omitempty"`                //
	Namespace          string                                           `json:"namespace,omitempty"`          //
	Type               string                                           `json:"type,omitempty"`               //
	Value              []GetServiceProviderDetailsResponseResponseValue `json:"value,omitempty"`              //
	Version            int                                              `json:"version,omitempty"`            //
}

// GetServiceProviderDetailsResponseResponseValue is the getServiceProviderDetailsResponseResponseValue definition
type GetServiceProviderDetailsResponseResponseValue struct {
	SLAProfileName string `json:"slaProfileName,omitempty"` //
	SpProfileName  string `json:"spProfileName,omitempty"`  //
	WanProvider    string `json:"wanProvider,omitempty"`    //
}

// UpdateDeviceCredentialsResponse is the updateDeviceCredentialsResponse definition
type UpdateDeviceCredentialsResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateGlobalPoolResponse is the updateGlobalPoolResponse definition
type UpdateGlobalPoolResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateNetworkResponse is the updateNetworkResponse definition
type UpdateNetworkResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateSPProfileResponse is the updateSPProfileResponse definition
type UpdateSPProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// AssignCredentialToSite assignCredentialToSite
/* Assign Device Credential To Site
@param __persistbapioutput Persist bapi sync response
@param siteID site id to assign credential.
*/
func (s *NetworkSettingsService) AssignCredentialToSite(siteID string, assignCredentialToSiteRequest *AssignCredentialToSiteRequest) (*AssignCredentialToSiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/credential-to-site/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(assignCredentialToSiteRequest).
		SetResult(&AssignCredentialToSiteResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation assignCredentialToSite")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createDeviceCredentials")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createGlobalPool")
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

	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(createNetworkRequest).
		SetResult(&CreateNetworkResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createNetwork")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createSPProfile")
	}

	result := response.Result().(*CreateSPProfileResponse)
	return result, response, err
}

// DeleteDeviceCredential deleteDeviceCredential
/* Delete device credential.
@param id global credential id
*/
func (s *NetworkSettingsService) DeleteDeviceCredential(id string) (*DeleteDeviceCredentialResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-credential/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteDeviceCredentialResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteDeviceCredential")
	}

	result := response.Result().(*DeleteDeviceCredentialResponse)
	return result, response, err
}

// DeleteGlobalIPPool deleteGlobalIPPool
/* API to delete global IP pool.
@param id global pool id
*/
func (s *NetworkSettingsService) DeleteGlobalIPPool(id string) (*DeleteGlobalIPPoolResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-pool/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteGlobalIPPoolResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteGlobalIPPool")
	}

	result := response.Result().(*DeleteGlobalIPPoolResponse)
	return result, response, err
}

// DeleteSPProfile deleteSPProfile
/* API to delete Service Provider profile (QoS).
@param spProfileName sp profile name
*/
func (s *NetworkSettingsService) DeleteSPProfile(spProfileName string) (*DeleteSPProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/sp-profile/{spProfileName}"
	path = strings.Replace(path, "{"+"spProfileName"+"}", fmt.Sprintf("%v", spProfileName), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteSPProfileResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteSPProfile")
	}

	result := response.Result().(*DeleteSPProfileResponse)
	return result, response, err
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getDeviceCredentialDetails")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getGlobalPool")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getNetwork")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getServiceProviderDetails")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateDeviceCredentials")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateGlobalPool")
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

	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteID), -1)

	response, err := RestyClient.R().
		SetBody(updateNetworkRequest).
		SetResult(&UpdateNetworkResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateNetwork")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateSPProfile")
	}

	result := response.Result().(*UpdateSPProfileResponse)
	return result, response, err
}
