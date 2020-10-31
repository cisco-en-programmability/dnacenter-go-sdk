package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DiscoveryService is the service to communicate with the Discovery API endpoint
type DiscoveryService service

// CLICredentialDTO is the CLICredentialDTO definition
type CLICredentialDTO struct {
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

// DiscoveryNIO is the DiscoveryNIO definition
type DiscoveryNIO struct {
	AttributeInfo          string              `json:"attributeInfo,omitempty"`          //
	CdpLevel               int                 `json:"cdpLevel,omitempty"`               //
	DeviceIDs              string              `json:"deviceIds,omitempty"`              //
	DiscoveryCondition     string              `json:"discoveryCondition,omitempty"`     //
	DiscoveryStatus        string              `json:"discoveryStatus,omitempty"`        //
	DiscoveryType          string              `json:"discoveryType,omitempty"`          //
	EnablePasswordList     string              `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string            `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     HTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    HTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	ID                     string              `json:"id,omitempty"`                     //
	IPAddressList          string              `json:"ipAddressList,omitempty"`          //
	IPFilterList           string              `json:"ipFilterList,omitempty"`           //
	IsAutoCdp              bool                `json:"isAutoCdp,omitempty"`              //
	LldpLevel              int                 `json:"lldpLevel,omitempty"`              //
	Name                   string              `json:"name,omitempty"`                   //
	NetconfPort            string              `json:"netconfPort,omitempty"`            //
	NumDevices             int                 `json:"numDevices,omitempty"`             //
	ParentDiscoveryID      string              `json:"parentDiscoveryId,omitempty"`      //
	PasswordList           string              `json:"passwordList,omitempty"`           //
	PreferredMgmtIPMethod  string              `json:"preferredMgmtIPMethod,omitempty"`  //
	ProtocolOrder          string              `json:"protocolOrder,omitempty"`          //
	RetryCount             int                 `json:"retryCount,omitempty"`             //
	SNMPAuthPassphrase     string              `json:"snmpAuthPassphrase,omitempty"`     //
	SNMPAuthProtocol       string              `json:"snmpAuthProtocol,omitempty"`       //
	SNMPMode               string              `json:"snmpMode,omitempty"`               //
	SNMPPrivPassphrase     string              `json:"snmpPrivPassphrase,omitempty"`     //
	SNMPPrivProtocol       string              `json:"snmpPrivProtocol,omitempty"`       //
	SNMPRoCommunity        string              `json:"snmpRoCommunity,omitempty"`        //
	SNMPRoCommunityDesc    string              `json:"snmpRoCommunityDesc,omitempty"`    //
	SNMPRwCommunity        string              `json:"snmpRwCommunity,omitempty"`        //
	SNMPRwCommunityDesc    string              `json:"snmpRwCommunityDesc,omitempty"`    //
	SNMPUserName           string              `json:"snmpUserName,omitempty"`           //
	TimeOut                int                 `json:"timeOut,omitempty"`                //
	UpdateMgmtIP           bool                `json:"updateMgmtIp,omitempty"`           //
	UserNameList           string              `json:"userNameList,omitempty"`           //
}

// HTTPWriteCredentialDTO is the HTTPWriteCredentialDTO definition
type HTTPWriteCredentialDTO struct {
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

// HTTPReadCredential is the HTTPReadCredential definition
type HTTPReadCredential struct {
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

// HTTPWriteCredential is the HTTPWriteCredential definition
type HTTPWriteCredential struct {
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

// InventoryRequest is the InventoryRequest definition
type InventoryRequest struct {
	CdpLevel               int                 `json:"cdpLevel,omitempty"`               //
	DiscoveryType          string              `json:"discoveryType,omitempty"`          //
	EnablePasswordList     []string            `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string            `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     HTTPReadCredential  `json:"httpReadCredential,omitempty"`     //
	HTTPWriteCredential    HTTPWriteCredential `json:"httpWriteCredential,omitempty"`    //
	IPAddressList          string              `json:"ipAddressList,omitempty"`          //
	IPFilterList           []string            `json:"ipFilterList,omitempty"`           //
	LldpLevel              int                 `json:"lldpLevel,omitempty"`              //
	Name                   string              `json:"name,omitempty"`                   //
	NetconfPort            string              `json:"netconfPort,omitempty"`            //
	NoAddNewDevice         bool                `json:"noAddNewDevice,omitempty"`         //
	ParentDiscoveryID      string              `json:"parentDiscoveryId,omitempty"`      //
	PasswordList           []string            `json:"passwordList,omitempty"`           //
	PreferredMgmtIPMethod  string              `json:"preferredMgmtIPMethod,omitempty"`  //
	ProtocolOrder          string              `json:"protocolOrder,omitempty"`          //
	ReDiscovery            bool                `json:"reDiscovery,omitempty"`            //
	Retry                  int                 `json:"retry,omitempty"`                  //
	SNMPAuthPassphrase     string              `json:"snmpAuthPassphrase,omitempty"`     //
	SNMPAuthProtocol       string              `json:"snmpAuthProtocol,omitempty"`       //
	SNMPMode               string              `json:"snmpMode,omitempty"`               //
	SNMPPrivPassphrase     string              `json:"snmpPrivPassphrase,omitempty"`     //
	SNMPPrivProtocol       string              `json:"snmpPrivProtocol,omitempty"`       //
	SNMPROCommunity        string              `json:"snmpROCommunity,omitempty"`        //
	SNMPROCommunityDesc    string              `json:"snmpROCommunityDesc,omitempty"`    //
	SNMPRWCommunity        string              `json:"snmpRWCommunity,omitempty"`        //
	SNMPRWCommunityDesc    string              `json:"snmpRWCommunityDesc,omitempty"`    //
	SNMPUserName           string              `json:"snmpUserName,omitempty"`           //
	SNMPVersion            string              `json:"snmpVersion,omitempty"`            //
	Timeout                int                 `json:"timeout,omitempty"`                //
	UpdateMgmtIP           bool                `json:"updateMgmtIp,omitempty"`           //
	UserNameList           []string            `json:"userNameList,omitempty"`           //
}

// NetconfCredentialDTO is the NetconfCredentialDTO definition
type NetconfCredentialDTO struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	NetconfPort      string `json:"netconfPort,omitempty"`      //
}

// SNMPvCredentialDTO is the SNMPvCredentialDTO definition
type SNMPvCredentialDTO struct {
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

// SNMPvReadCommunityDTO is the SNMPvReadCommunityDTO definition
type SNMPvReadCommunityDTO struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	ReadCommunity    string `json:"readCommunity,omitempty"`    //
}

// SNMPvWriteCommunityDTO is the SNMPvWriteCommunityDTO definition
type SNMPvWriteCommunityDTO struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	WriteCommunity   string `json:"writeCommunity,omitempty"`   //
}

// SitesInfoDTO is the SitesInfoDTO definition
type SitesInfoDTO struct {
	SiteUUIDs []string `json:"siteUuids,omitempty"` //
}

// SystemPropertyNameAndIntValueDTO is the SystemPropertyNameAndIntValueDTO definition
type SystemPropertyNameAndIntValueDTO struct {
	ID                 string `json:"id,omitempty"`                 //
	InstanceTenantID   string `json:"instanceTenantId,omitempty"`   //
	InstanceUUID       string `json:"instanceUuid,omitempty"`       //
	IntValue           int    `json:"intValue,omitempty"`           //
	SystemPropertyName string `json:"systemPropertyName,omitempty"` //
}

// DiscoveryJobNIOListResult is the DiscoveryJobNIOListResult definition
type DiscoveryJobNIOListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// DiscoveryNIOListResult is the DiscoveryNIOListResult definition
type DiscoveryNIOListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// DiscoveryNIOResult is the DiscoveryNIOResult definition
type DiscoveryNIOResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GlobalCredentialListResult is the GlobalCredentialListResult definition
type GlobalCredentialListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// GlobalCredentialSubTypeResult is the GlobalCredentialSubTypeResult definition
type GlobalCredentialSubTypeResult struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// NetworkDeviceNIOListResult is the NetworkDeviceNIOListResult definition
type NetworkDeviceNIOListResult struct {
	Response []TaskResponse `json:"response,omitempty"` //
	Version  string         `json:"version,omitempty"`  //
}

// SystemPropertyListResult is the SystemPropertyListResult definition
type SystemPropertyListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// CreateCLICredentials createCLICredentials
/* Adds global CLI credential
 */
// func (s *DiscoveryService) CreateCLICredentials(createCLICredentialsRequest *CreateCLICredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/cli"

// 	response, err := RestyClient.R().
// 		SetBody(createCLICredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// CreateHTTPReadCredentials createHTTPReadCredentials
/* Adds HTTP read credentials
 */
// func (s *DiscoveryService) CreateHTTPReadCredentials(createHTTPReadCredentialsRequest *CreateHTTPReadCredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/http-read"

// 	response, err := RestyClient.R().
// 		SetBody(createHTTPReadCredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // CreateHTTPWriteCredentials createHTTPWriteCredentials
// /* Adds global HTTP write credentials
//  */
// func (s *DiscoveryService) CreateHTTPWriteCredentials(createHTTPWriteCredentialsRequest *CreateHTTPWriteCredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/http-write"

// 	response, err := RestyClient.R().
// 		SetBody(createHTTPWriteCredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // CreateNetconfCredentials createNetconfCredentials
// /* Adds global netconf credentials
//  */
// func (s *DiscoveryService) CreateNetconfCredentials(createNetconfCredentialsRequest *CreateNetconfCredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/netconf"

// 	response, err := RestyClient.R().
// 		SetBody(createNetconfCredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // CreateSNMPReadCommunity createSNMPReadCommunity
// /* Adds global SNMP read community
//  */
// func (s *DiscoveryService) CreateSNMPReadCommunity(createSNMPReadCommunityRequest *CreateSNMPReadCommunityRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

// 	response, err := RestyClient.R().
// 		SetBody(createSNMPReadCommunityRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // CreateSNMPWriteCommunity createSNMPWriteCommunity
// /* Adds global SNMP write community
//  */
// func (s *DiscoveryService) CreateSNMPWriteCommunity(createSNMPWriteCommunityRequest *CreateSNMPWriteCommunityRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

// 	response, err := RestyClient.R().
// 		SetBody(createSNMPWriteCommunityRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // CreateSNMPv3Credentials createSNMPv3Credentials
// /* Adds global SNMPv3 credentials
//  */
// func (s *DiscoveryService) CreateSNMPv3Credentials(createSNMPv3CredentialsRequest *CreateSNMPv3CredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/snmpv3"

// 	response, err := RestyClient.R().
// 		SetBody(createSNMPv3CredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // CreateUpdateSNMPProperties createUpdateSNMPProperties
// /* Adds SNMP properties
//  */
// func (s *DiscoveryService) CreateUpdateSNMPProperties(createUpdateSNMPPropertiesRequest *CreateUpdateSNMPPropertiesRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/snmp-property"

// 	response, err := RestyClient.R().
// 		SetBody(createUpdateSNMPPropertiesRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// DeleteAllDiscovery deleteAllDiscovery
/* Stops all the discoveries and removes them
 */
func (s *DiscoveryService) DeleteAllDiscovery() (*resty.Response, error) {

	path := "/dna/intent/api/v1/discovery"

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteDiscoveryByID deleteDiscoveryById
/* Stops the discovery for the given Discovery ID and removes it. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
*/
func (s *DiscoveryService) DeleteDiscoveryByID(id string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteDiscoveryBySpecifiedRange deleteDiscoveryBySpecifiedRange
/* Stops discovery for the given range and removes them
@param startIndex Start index
@param recordsToDelete Number of records to delete
*/
func (s *DiscoveryService) DeleteDiscoveryBySpecifiedRange(startIndex int, recordsToDelete int) (*resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToDelete}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToDelete"+"}", fmt.Sprintf("%v", recordsToDelete), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteGlobalCredentialsByID deleteGlobalCredentialsById
/* Deletes global credential for the given ID
@param globalCredentialID ID of global-credential
*/
func (s *DiscoveryService) DeleteGlobalCredentialsByID(globalCredentialID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/{globalCredentialID}"
	path = strings.Replace(path, "{"+"globalCredentialID"+"}", fmt.Sprintf("%v", globalCredentialID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetCountOfAllDiscoveryJobs getCountOfAllDiscoveryJobs
/* Returns the count of all available discovery jobs
 */
func (s *DiscoveryService) GetCountOfAllDiscoveryJobs() (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/count"

	response, err := RestyClient.R().
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetCredentialSubTypeByCredentialID getCredentialSubTypeByCredentialId
/* Returns the credential sub type for the given ID
@param id Global Credential ID
*/
func (s *DiscoveryService) GetCredentialSubTypeByCredentialID(id string) (*GlobalCredentialSubTypeResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GlobalCredentialSubTypeResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GlobalCredentialSubTypeResult)
	return result, response, err

}

// GetDevicesDiscoveredByIDQueryParams defines the query parameters for this request
type GetDevicesDiscoveredByIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` // taskId
}

// GetDevicesDiscoveredByID getDevicesDiscoveredById
/* Returns the count of network devices discovered in the given discovery. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
@param taskID taskId
*/
func (s *DiscoveryService) GetDevicesDiscoveredByID(id string, getDevicesDiscoveredByIDQueryParams *GetDevicesDiscoveredByIDQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/network-device/count"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getDevicesDiscoveredByIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetDiscoveredDevicesByRangeQueryParams defines the query parameters for this request
type GetDiscoveredDevicesByRangeQueryParams struct {
	TaskID string `url:"taskId,omitempty"` // taskId
}

// GetDiscoveredDevicesByRange getDiscoveredDevicesByRange
/* Returns the network devices discovered for the given discovery and for the given range. The maximum int of records that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
@param startIndex Start index
@param recordsToReturn Number of records to return
@param taskID taskId
*/
func (s *DiscoveryService) GetDiscoveredDevicesByRange(id string, startIndex int, recordsToReturn int, getDiscoveredDevicesByRangeQueryParams *GetDiscoveredDevicesByRangeQueryParams) (*NetworkDeviceNIOListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	queryString, _ := query.Values(getDiscoveredDevicesByRangeQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&NetworkDeviceNIOListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceNIOListResult)
	return result, response, err

}

// GetDiscoveredNetworkDevicesByDiscoveryIDQueryParams defines the query parameters for this request
type GetDiscoveredNetworkDevicesByDiscoveryIDQueryParams struct {
	TaskID string `url:"taskId,omitempty"` // taskId
}

// GetDiscoveredNetworkDevicesByDiscoveryID getDiscoveredNetworkDevicesByDiscoveryId
/* Returns the network devices discovered for the given Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
@param taskID taskId
*/
func (s *DiscoveryService) GetDiscoveredNetworkDevicesByDiscoveryID(id string, getDiscoveredNetworkDevicesByDiscoveryIDQueryParams *GetDiscoveredNetworkDevicesByDiscoveryIDQueryParams) (*NetworkDeviceNIOListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/network-device"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getDiscoveredNetworkDevicesByDiscoveryIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&NetworkDeviceNIOListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NetworkDeviceNIOListResult)
	return result, response, err

}

// GetDiscoveriesByRange getDiscoveriesByRange
/* Returns the discovery by specified range
@param startIndex Start index
@param recordsToReturn Number of records to return
*/
func (s *DiscoveryService) GetDiscoveriesByRange(startIndex int, recordsToReturn int) (*DiscoveryNIOListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := RestyClient.R().
		SetResult(&DiscoveryNIOListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DiscoveryNIOListResult)
	return result, response, err

}

// GetDiscoveryByID getDiscoveryById
/* Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
*/
func (s *DiscoveryService) GetDiscoveryByID(id string) (*DiscoveryNIOResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DiscoveryNIOResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DiscoveryNIOResult)
	return result, response, err

}

// GetDiscoveryJobsByIPQueryParams defines the query parameters for this request
type GetDiscoveryJobsByIPQueryParams struct {
	Offset    int    `url:"offset,omitempty"`    // offset
	Limit     int    `url:"limit,omitempty"`     // limit
	IPAddress string `url:"ipAddress,omitempty"` // ipAddress
	Name      string `url:"name,omitempty"`      // name
}

// GetDiscoveryJobsByIP getDiscoveryJobsByIP
/* Returns the list of discovery jobs for the given IP
@param offset offset
@param limit limit
@param ipAddress ipAddress
@param name name
*/
func (s *DiscoveryService) GetDiscoveryJobsByIP(getDiscoveryJobsByIPQueryParams *GetDiscoveryJobsByIPQueryParams) (*DiscoveryJobNIOListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/job"

	queryString, _ := query.Values(getDiscoveryJobsByIPQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DiscoveryJobNIOListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DiscoveryJobNIOListResult)
	return result, response, err

}

// GetGlobalCredentialsQueryParams defines the query parameters for this request
type GetGlobalCredentialsQueryParams struct {
	CredentialSubType string `url:"credentialSubType,omitempty"` // Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3 / HTTP_WRITE / HTTP_READ / NETCONF
	SortBy            string `url:"sortBy,omitempty"`            // sortBy
	Order             string `url:"order,omitempty"`             // order
}

// GetGlobalCredentials getGlobalCredentials
/* Returns global credential for the given credential sub type
@param credentialSubType Credential type as CLI / SNMPV2_READ_COMMUNITY / SNMPV2_WRITE_COMMUNITY / SNMPV3 / HTTP_WRITE / HTTP_READ / NETCONF
@param sortBy sortBy
@param order order
*/
func (s *DiscoveryService) GetGlobalCredentials(getGlobalCredentialsQueryParams *GetGlobalCredentialsQueryParams) (*GlobalCredentialListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential"

	queryString, _ := query.Values(getGlobalCredentialsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GlobalCredentialListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GlobalCredentialListResult)
	return result, response, err

}

// GetListOfDiscoveriesByDiscoveryIDQueryParams defines the query parameters for this request
type GetListOfDiscoveriesByDiscoveryIDQueryParams struct {
	Offset    int    `url:"offset,omitempty"`    // offset
	Limit     int    `url:"limit,omitempty"`     // limit
	IPAddress string `url:"ipAddress,omitempty"` // ipAddress
}

// GetListOfDiscoveriesByDiscoveryID getListOfDiscoveriesByDiscoveryId
/* Returns the list of discovery jobs for the given Discovery ID. The results can be optionally filtered based on IP. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
@param offset offset
@param limit limit
@param ipAddress ipAddress
*/
func (s *DiscoveryService) GetListOfDiscoveriesByDiscoveryID(id string, getListOfDiscoveriesByDiscoveryIDQueryParams *GetListOfDiscoveriesByDiscoveryIDQueryParams) (*DiscoveryJobNIOListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/job"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getListOfDiscoveriesByDiscoveryIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DiscoveryJobNIOListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DiscoveryJobNIOListResult)
	return result, response, err

}

// GetNetworkDevicesFromDiscoveryQueryParams defines the query parameters for this request
type GetNetworkDevicesFromDiscoveryQueryParams struct {
	TaskID        string   `url:"taskId,omitempty"`        // taskId
	SortBy        string   `url:"sortBy,omitempty"`        // sortBy
	SortOrder     string   `url:"sortOrder,omitempty"`     // sortOrder
	IPAddress     []string `url:"ipAddress,omitempty"`     // ipAddress
	PingStatus    []string `url:"pingStatus,omitempty"`    // pingStatus
	SNMPStatus    []string `url:"snmpStatus,omitempty"`    // snmpStatus
	CliStatus     []string `url:"cliStatus,omitempty"`     // cliStatus
	NetconfStatus []string `url:"netconfStatus,omitempty"` // netconfStatus
	HTTPStatus    []string `url:"httpStatus,omitempty"`    // httpStatus
}

// GetNetworkDevicesFromDiscovery getNetworkDevicesFromDiscovery
/* Returns the network devices from a discovery job based on given filters. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
@param taskID taskId
@param sortBy sortBy
@param sortOrder sortOrder
@param ipAddress ipAddress
@param pingStatus pingStatus
@param snmpStatus snmpStatus
@param cliStatus cliStatus
@param netconfStatus netconfStatus
@param httpStatus httpStatus
*/
func (s *DiscoveryService) GetNetworkDevicesFromDiscovery(id string, getNetworkDevicesFromDiscoveryQueryParams *GetNetworkDevicesFromDiscoveryQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/summary"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getNetworkDevicesFromDiscoveryQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetSNMPProperties getSNMPProperties
/* Returns SNMP properties
 */
func (s *DiscoveryService) GetSNMPProperties() (*SystemPropertyListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/snmp-property"

	response, err := RestyClient.R().
		SetResult(&SystemPropertyListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*SystemPropertyListResult)
	return result, response, err

}

// // StartDiscovery startDiscovery
// /* Initiates discovery with the given parameters
//  */
// func (s *DiscoveryService) StartDiscovery(startDiscoveryRequest *StartDiscoveryRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/discovery"

// 	response, err := RestyClient.R().
// 		SetBody(startDiscoveryRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateCLICredentials updateCLICredentials
// /* Updates global CLI credentials
//  */
// func (s *DiscoveryService) UpdateCLICredentials(updateCLICredentialsRequest *UpdateCLICredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/cli"

// 	response, err := RestyClient.R().
// 		SetBody(updateCLICredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateGlobalCredentials updateGlobalCredentials
// /* Update global credential for network devices in site(s)
// @param globalCredentialID Global credential UUID
// */
// func (s *DiscoveryService) UpdateGlobalCredentials(globalCredentialID string, updateGlobalCredentialsRequest *UpdateGlobalCredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/{globalCredentialID}"
// 	path = strings.Replace(path, "{"+"globalCredentialID"+"}", fmt.Sprintf("%v", globalCredentialId), -1)

// 	response, err := RestyClient.R().
// 		SetBody(updateGlobalCredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateHTTPReadCredential updateHTTPReadCredential
// /* Updates global HTTP Read credential
//  */
// func (s *DiscoveryService) UpdateHTTPReadCredential(updateHTTPReadCredentialRequest *UpdateHTTPReadCredentialRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/http-read"

// 	response, err := RestyClient.R().
// 		SetBody(updateHTTPReadCredentialRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateHTTPWriteCredentials updateHTTPWriteCredentials
// /* Updates global HTTP write credentials
//  */
// func (s *DiscoveryService) UpdateHTTPWriteCredentials(updateHTTPWriteCredentialsRequest *UpdateHTTPWriteCredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/http-write"

// 	response, err := RestyClient.R().
// 		SetBody(updateHTTPWriteCredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateNetconfCredentials updateNetconfCredentials
// /* Updates global netconf credentials
//  */
// func (s *DiscoveryService) UpdateNetconfCredentials(updateNetconfCredentialsRequest *UpdateNetconfCredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/netconf"

// 	response, err := RestyClient.R().
// 		SetBody(updateNetconfCredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateSNMPReadCommunity updateSNMPReadCommunity
// /* Updates global SNMP read community
//  */
// func (s *DiscoveryService) UpdateSNMPReadCommunity(updateSNMPReadCommunityRequest *UpdateSNMPReadCommunityRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

// 	response, err := RestyClient.R().
// 		SetBody(updateSNMPReadCommunityRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateSNMPWriteCommunity updateSNMPWriteCommunity
// /* Updates global SNMP write community
//  */
// func (s *DiscoveryService) UpdateSNMPWriteCommunity(updateSNMPWriteCommunityRequest *UpdateSNMPWriteCommunityRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

// 	response, err := RestyClient.R().
// 		SetBody(updateSNMPWriteCommunityRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdateSNMPv3Credentials updateSNMPv3Credentials
// /* Updates global SNMPv3 credential
//  */
// func (s *DiscoveryService) UpdateSNMPv3Credentials(updateSNMPv3CredentialsRequest *UpdateSNMPv3CredentialsRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/global-credential/snmpv3"

// 	response, err := RestyClient.R().
// 		SetBody(updateSNMPv3CredentialsRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// // UpdatesAnExistingDiscoveryBySpecifiedID updatesAnExistingDiscoveryBySpecifiedId
// /* Stops or starts an existing discovery
//  */
// func (s *DiscoveryService) UpdatesAnExistingDiscoveryBySpecifiedID(updatesAnExistingDiscoveryBySpecifiedIDRequest *UpdatesAnExistingDiscoveryBySpecifiedIDRequest) (*TaskIdResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/discovery"

// 	response, err := RestyClient.R().
// 		SetBody(updatesAnExistingDiscoveryBySpecifiedIDRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }
