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
	CliId         string `json:"cliId,omitempty"`         //
	HttpRead      string `json:"httpRead,omitempty"`      //
	HttpWrite     string `json:"httpWrite,omitempty"`     //
	SnmpV2ReadId  string `json:"snmpV2ReadId,omitempty"`  //
	SnmpV2WriteId string `json:"snmpV2WriteId,omitempty"` //
	SnmpV3Id      string `json:"snmpV3Id,omitempty"`      //
}

// CliCredential is the CliCredential definition
type CliCredential struct {
	Description    string `json:"description,omitempty"`    //
	EnablePassword string `json:"enablePassword,omitempty"` //
	Id             string `json:"id,omitempty"`             //
	Password       string `json:"password,omitempty"`       //
	Username       string `json:"username,omitempty"`       //
}

// ClientAndEndpoint_aaa is the ClientAndEndpoint_aaa definition
type ClientAndEndpoint_aaa struct {
	IpAddress    string `json:"ipAddress,omitempty"`    //
	Network      string `json:"network,omitempty"`      //
	Protocol     string `json:"protocol,omitempty"`     //
	Servers      string `json:"servers,omitempty"`      //
	SharedSecret string `json:"sharedSecret,omitempty"` //
}

// CreateDeviceCredentialsRequest is the CreateDeviceCredentialsRequest definition
type CreateDeviceCredentialsRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// CreateGlobalPoolRequest is the CreateGlobalPoolRequest definition
type CreateGlobalPoolRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// CreateNetworkRequest is the CreateNetworkRequest definition
type CreateNetworkRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// CreateSPProfileRequest is the CreateSPProfileRequest definition
type CreateSPProfileRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// DnsServer is the DnsServer definition
type DnsServer struct {
	DomainName         string `json:"domainName,omitempty"`         //
	PrimaryIpAddress   string `json:"primaryIpAddress,omitempty"`   //
	SecondaryIpAddress string `json:"secondaryIpAddress,omitempty"` //
}

// HttpsRead is the HttpsRead definition
type HttpsRead struct {
	Id       string `json:"id,omitempty"`       //
	Name     string `json:"name,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Port     string `json:"port,omitempty"`     //
	Username string `json:"username,omitempty"` //
}

// HttpsWrite is the HttpsWrite definition
type HttpsWrite struct {
	Id       string `json:"id,omitempty"`       //
	Name     string `json:"name,omitempty"`     //
	Password string `json:"password,omitempty"` //
	Port     string `json:"port,omitempty"`     //
	Username string `json:"username,omitempty"` //
}

// Ippool is the Ippool definition
type Ippool struct {
	DhcpServerIps []string `json:"dhcpServerIps,omitempty"` //
	DnsServerIps  []string `json:"dnsServerIps,omitempty"`  //
	Gateway       string   `json:"gateway,omitempty"`       //
	Id            string   `json:"id,omitempty"`            //
	IpPoolName    string   `json:"ipPoolName,omitempty"`    //
}

// MessageOfTheday is the MessageOfTheday definition
type MessageOfTheday struct {
	BannerMessage        string `json:"bannerMessage,omitempty"`        //
	RetainExistingBanner bool   `json:"retainExistingBanner,omitempty"` //
}

// Netflowcollector is the Netflowcollector definition
type Netflowcollector struct {
	IpAddress string `json:"ipAddress,omitempty"` //
	Port      int    `json:"port,omitempty"`      //
}

// Network_aaa is the Network_aaa definition
type Network_aaa struct {
	IpAddress    string `json:"ipAddress,omitempty"`    //
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

// Settings is the Settings definition
type Settings struct {
	Qos []Qos `json:"qos,omitempty"` //
}

// SnmpServer is the SnmpServer definition
type SnmpServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IpAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// SnmpV2cRead is the SnmpV2cRead definition
type SnmpV2cRead struct {
	Description   string `json:"description,omitempty"`   //
	Id            string `json:"id,omitempty"`            //
	ReadCommunity string `json:"readCommunity,omitempty"` //
}

// SnmpV2cWrite is the SnmpV2cWrite definition
type SnmpV2cWrite struct {
	Description    string `json:"description,omitempty"`    //
	Id             string `json:"id,omitempty"`             //
	WriteCommunity string `json:"writeCommunity,omitempty"` //
}

// SnmpV3 is the SnmpV3 definition
type SnmpV3 struct {
	AuthPassword    string `json:"authPassword,omitempty"`    //
	AuthType        string `json:"authType,omitempty"`        //
	Description     string `json:"description,omitempty"`     //
	Id              string `json:"id,omitempty"`              //
	PrivacyPassword string `json:"privacyPassword,omitempty"` //
	PrivacyType     string `json:"privacyType,omitempty"`     //
	SnmpMode        string `json:"snmpMode,omitempty"`        //
	Username        string `json:"username,omitempty"`        //
}

// SyslogServer is the SyslogServer definition
type SyslogServer struct {
	ConfigureDnacIP bool     `json:"configureDnacIP,omitempty"` //
	IpAddresses     []string `json:"ipAddresses,omitempty"`     //
}

// UpdateDeviceCredentialsRequest is the UpdateDeviceCredentialsRequest definition
type UpdateDeviceCredentialsRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// UpdateGlobalPoolRequest is the UpdateGlobalPoolRequest definition
type UpdateGlobalPoolRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// UpdateNetworkRequest is the UpdateNetworkRequest definition
type UpdateNetworkRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// UpdateSPProfileRequest is the UpdateSPProfileRequest definition
type UpdateSPProfileRequest struct {
	Settings Settings `json:"settings,omitempty"` //
}

// AssignCredentialToSiteResponse is the AssignCredentialToSiteResponse definition
type AssignCredentialToSiteResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// Cli is the Cli definition
type Cli struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	EnablePassword   string `json:"enablePassword,omitempty"`   //
	Id               string `json:"id,omitempty"`               //
	InstanceTenantId string `json:"instanceTenantId,omitempty"` //
	InstanceUuid     string `json:"instanceUuid,omitempty"`     //
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
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateGlobalPoolResponse is the CreateGlobalPoolResponse definition
type CreateGlobalPoolResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateNetworkResponse is the CreateNetworkResponse definition
type CreateNetworkResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateSPProfileResponse is the CreateSPProfileResponse definition
type CreateSPProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteDeviceCredentialResponse is the DeleteDeviceCredentialResponse definition
type DeleteDeviceCredentialResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteGlobalIPPoolResponse is the DeleteGlobalIPPoolResponse definition
type DeleteGlobalIPPoolResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteSPProfileResponse is the DeleteSPProfileResponse definition
type DeleteSPProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetDeviceCredentialDetailsResponse is the GetDeviceCredentialDetailsResponse definition
type GetDeviceCredentialDetailsResponse struct {
	Cli         []Cli         `json:"cli,omitempty"`           //
	HttpRead    []HttpRead    `json:"http_read,omitempty"`     //
	HttpWrite   []HttpWrite   `json:"http_write,omitempty"`    //
	SnmpV2Read  []SnmpV2Read  `json:"snmp_v2_read,omitempty"`  //
	SnmpV2Write []SnmpV2Write `json:"snmp_v2_write,omitempty"` //
	SnmpV3      []SnmpV3      `json:"snmp_v3,omitempty"`       //
}

// GetGlobalPoolResponse is the GetGlobalPoolResponse definition
type GetGlobalPoolResponse struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// GetNetworkResponse is the GetNetworkResponse definition
type GetNetworkResponse struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// GetServiceProviderDetailsResponse is the GetServiceProviderDetailsResponse definition
type GetServiceProviderDetailsResponse struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// Http_read is the Http_read definition
type Http_read struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	Id               string `json:"id,omitempty"`               //
	InstanceTenantId string `json:"instanceTenantId,omitempty"` //
	InstanceUuid     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             string `json:"port,omitempty"`             //
	Secure           string `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}

// Http_write is the Http_write definition
type Http_write struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	Id               string `json:"id,omitempty"`               //
	InstanceTenantId string `json:"instanceTenantId,omitempty"` //
	InstanceUuid     string `json:"instanceUuid,omitempty"`     //
	Password         string `json:"password,omitempty"`         //
	Port             string `json:"port,omitempty"`             //
	Secure           string `json:"secure,omitempty"`           //
	Username         string `json:"username,omitempty"`         //
}

// Response is the Response definition
type Response struct {
	GroupUuid          string  `json:"groupUuid,omitempty"`          //
	InheritedGroupName string  `json:"inheritedGroupName,omitempty"` //
	InheritedGroupUuid string  `json:"inheritedGroupUuid,omitempty"` //
	InstanceType       string  `json:"instanceType,omitempty"`       //
	InstanceUuid       string  `json:"instanceUuid,omitempty"`       //
	Key                string  `json:"key,omitempty"`                //
	Namespace          string  `json:"namespace,omitempty"`          //
	Type               string  `json:"type,omitempty"`               //
	Value              []Value `json:"value,omitempty"`              //
	Version            string  `json:"version,omitempty"`            //
}

// Snmp_v2_read is the Snmp_v2_read definition
type Snmp_v2_read struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	Id               string `json:"id,omitempty"`               //
	InstanceTenantId string `json:"instanceTenantId,omitempty"` //
	InstanceUuid     string `json:"instanceUuid,omitempty"`     //
	ReadCommunity    string `json:"readCommunity,omitempty"`    //
}

// Snmp_v2_write is the Snmp_v2_write definition
type Snmp_v2_write struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	Id               string `json:"id,omitempty"`               //
	InstanceTenantId string `json:"instanceTenantId,omitempty"` //
	InstanceUuid     string `json:"instanceUuid,omitempty"`     //
	WriteCommunity   string `json:"writeCommunity,omitempty"`   //
}

// Snmp_v3 is the Snmp_v3 definition
type Snmp_v3 struct {
	AuthPassword     string `json:"authPassword,omitempty"`     //
	AuthType         string `json:"authType,omitempty"`         //
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	Id               string `json:"id,omitempty"`               //
	InstanceTenantId string `json:"instanceTenantId,omitempty"` //
	InstanceUuid     string `json:"instanceUuid,omitempty"`     //
	PrivacyPassword  string `json:"privacyPassword,omitempty"`  //
	PrivacyType      string `json:"privacyType,omitempty"`      //
	SnmpMode         string `json:"snmpMode,omitempty"`         //
	Username         string `json:"username,omitempty"`         //
}

// UpdateDeviceCredentialsResponse is the UpdateDeviceCredentialsResponse definition
type UpdateDeviceCredentialsResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateGlobalPoolResponse is the UpdateGlobalPoolResponse definition
type UpdateGlobalPoolResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateNetworkResponse is the UpdateNetworkResponse definition
type UpdateNetworkResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// UpdateSPProfileResponse is the UpdateSPProfileResponse definition
type UpdateSPProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// Value is the Value definition
type Value struct {
	SlaProfileName string `json:"slaProfileName,omitempty"` //
	SpProfileName  string `json:"spProfileName,omitempty"`  //
	WanProvider    string `json:"wanProvider,omitempty"`    //
}

// AssignCredentialToSite assignCredentialToSite
/* Assign Device Credential To Site
@param __persistbapioutput Persist bapi sync response
@param siteId site id to assign credential.
*/
func (s *NetworkSettingsService) AssignCredentialToSite(siteId string, assignCredentialToSiteRequest *AssignCredentialToSiteRequest) (*AssignCredentialToSiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/credential-to-site/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteId), -1)

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
@param siteId Site id to which site details to associate with the network settings.
*/
func (s *NetworkSettingsService) CreateNetwork(siteId string, createNetworkRequest *CreateNetworkRequest) (*CreateNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteId), -1)

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
	SiteId string `url:"siteId,omitempty"` // Site id to retrieve the credential details associated with the site.
}

// GetDeviceCredentialDetails getDeviceCredentialDetails
/* API to get device credential details.
@param siteId Site id to retrieve the credential details associated with the site.
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
	SiteId string `url:"siteId,omitempty"` // Site id to get the network settings associated with the site.
}

// GetNetwork getNetwork
/* API to get  DHCP and DNS center server details.
@param siteId Site id to get the network settings associated with the site.
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
@param siteId Site id to update the network settings which is associated with the site
*/
func (s *NetworkSettingsService) UpdateNetwork(siteId string, updateNetworkRequest *UpdateNetworkRequest) (*UpdateNetworkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/network/{siteId}"
	path = strings.Replace(path, "{"+"siteId"+"}", fmt.Sprintf("%v", siteId), -1)

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
