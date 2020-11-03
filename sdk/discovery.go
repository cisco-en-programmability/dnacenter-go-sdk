package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DiscoveryService is the service to communicate with the Discovery API endpoint
type DiscoveryService service

// CreateCLICredentialsRequest is the CreateCLICredentialsRequest definition
type CreateCLICredentialsRequest struct {
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

// CreateHTTPReadCredentialsRequest is the CreateHTTPReadCredentialsRequest definition
type CreateHTTPReadCredentialsRequest struct {
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

// CreateHTTPWriteCredentialsRequest is the CreateHTTPWriteCredentialsRequest definition
type CreateHTTPWriteCredentialsRequest struct {
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

// CreateNetconfCredentialsRequest is the CreateNetconfCredentialsRequest definition
type CreateNetconfCredentialsRequest struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	NetconfPort      string `json:"netconfPort,omitempty"`      //
}

// CreateSNMPReadCommunityRequest is the CreateSNMPReadCommunityRequest definition
type CreateSNMPReadCommunityRequest struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	ReadCommunity    string `json:"readCommunity,omitempty"`    //
}

// CreateSNMPWriteCommunityRequest is the CreateSNMPWriteCommunityRequest definition
type CreateSNMPWriteCommunityRequest struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	WriteCommunity   string `json:"writeCommunity,omitempty"`   //
}

// CreateSNMPv3CredentialsRequest is the CreateSNMPv3CredentialsRequest definition
type CreateSNMPv3CredentialsRequest struct {
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

// CreateUpdateSNMPPropertiesRequest is the CreateUpdateSNMPPropertiesRequest definition
type CreateUpdateSNMPPropertiesRequest struct {
	ID                 string `json:"id,omitempty"`                 //
	InstanceTenantID   string `json:"instanceTenantId,omitempty"`   //
	InstanceUUID       string `json:"instanceUuid,omitempty"`       //
	IntValue           int    `json:"intValue,omitempty"`           //
	SystemPropertyName string `json:"systemPropertyName,omitempty"` //
}

// StartDiscoveryRequest is the StartDiscoveryRequest definition
type StartDiscoveryRequest struct {
	CdpLevel               int      `json:"cdpLevel,omitempty"`               //
	DiscoveryType          string   `json:"discoveryType,omitempty"`          //
	EnablePasswordList     []string `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     struct {
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
	} `json:"httpReadCredential,omitempty"` //
	HTTPWriteCredential struct {
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
	} `json:"httpWriteCredential,omitempty"` //
	IPAddressList         string   `json:"ipAddressList,omitempty"`         //
	IPFilterList          []string `json:"ipFilterList,omitempty"`          //
	LldpLevel             int      `json:"lldpLevel,omitempty"`             //
	Name                  string   `json:"name,omitempty"`                  //
	NetconfPort           string   `json:"netconfPort,omitempty"`           //
	NoAddNewDevice        bool     `json:"noAddNewDevice,omitempty"`        //
	ParentDiscoveryID     string   `json:"parentDiscoveryId,omitempty"`     //
	PasswordList          []string `json:"passwordList,omitempty"`          //
	PreferredMgmtIPMethod string   `json:"preferredMgmtIPMethod,omitempty"` //
	ProtocolOrder         string   `json:"protocolOrder,omitempty"`         //
	ReDiscovery           bool     `json:"reDiscovery,omitempty"`           //
	Retry                 int      `json:"retry,omitempty"`                 //
	SNMPAuthPassphrase    string   `json:"snmpAuthPassphrase,omitempty"`    //
	SNMPAuthProtocol      string   `json:"snmpAuthProtocol,omitempty"`      //
	SNMPMode              string   `json:"snmpMode,omitempty"`              //
	SNMPPrivPassphrase    string   `json:"snmpPrivPassphrase,omitempty"`    //
	SNMPPrivProtocol      string   `json:"snmpPrivProtocol,omitempty"`      //
	SNMPROCommunity       string   `json:"snmpROCommunity,omitempty"`       //
	SNMPROCommunityDesc   string   `json:"snmpROCommunityDesc,omitempty"`   //
	SNMPRWCommunity       string   `json:"snmpRWCommunity,omitempty"`       //
	SNMPRWCommunityDesc   string   `json:"snmpRWCommunityDesc,omitempty"`   //
	SNMPUserName          string   `json:"snmpUserName,omitempty"`          //
	SNMPVersion           string   `json:"snmpVersion,omitempty"`           //
	Timeout               int      `json:"timeout,omitempty"`               //
	UpdateMgmtIP          bool     `json:"updateMgmtIp,omitempty"`          //
	UserNameList          []string `json:"userNameList,omitempty"`          //
}

// UpdateCLICredentialsRequest is the UpdateCLICredentialsRequest definition
type UpdateCLICredentialsRequest struct {
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

// UpdateGlobalCredentialsRequest is the UpdateGlobalCredentialsRequest definition
type UpdateGlobalCredentialsRequest struct {
	SiteUUIDs []string `json:"siteUuids,omitempty"` //
}

// UpdateHTTPReadCredentialRequest is the UpdateHTTPReadCredentialRequest definition
type UpdateHTTPReadCredentialRequest struct {
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

// UpdateHTTPWriteCredentialsRequest is the UpdateHTTPWriteCredentialsRequest definition
type UpdateHTTPWriteCredentialsRequest struct {
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

// UpdateNetconfCredentialsRequest is the UpdateNetconfCredentialsRequest definition
type UpdateNetconfCredentialsRequest struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	NetconfPort      string `json:"netconfPort,omitempty"`      //
}

// UpdateSNMPReadCommunityRequest is the UpdateSNMPReadCommunityRequest definition
type UpdateSNMPReadCommunityRequest struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	ReadCommunity    string `json:"readCommunity,omitempty"`    //
}

// UpdateSNMPWriteCommunityRequest is the UpdateSNMPWriteCommunityRequest definition
type UpdateSNMPWriteCommunityRequest struct {
	Comments         string `json:"comments,omitempty"`         //
	CredentialType   string `json:"credentialType,omitempty"`   //
	Description      string `json:"description,omitempty"`      //
	ID               string `json:"id,omitempty"`               //
	InstanceTenantID string `json:"instanceTenantId,omitempty"` //
	InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	WriteCommunity   string `json:"writeCommunity,omitempty"`   //
}

// UpdateSNMPv3CredentialsRequest is the UpdateSNMPv3CredentialsRequest definition
type UpdateSNMPv3CredentialsRequest struct {
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

// UpdatesAnExistingDiscoveryBySpecifiedIDRequest is the UpdatesAnExistingDiscoveryBySpecifiedIdRequest definition
type UpdatesAnExistingDiscoveryBySpecifiedIDRequest struct {
	AttributeInfo          string   `json:"attributeInfo,omitempty"`          //
	CdpLevel               int      `json:"cdpLevel,omitempty"`               //
	DeviceIDs              string   `json:"deviceIds,omitempty"`              //
	DiscoveryCondition     string   `json:"discoveryCondition,omitempty"`     //
	DiscoveryStatus        string   `json:"discoveryStatus,omitempty"`        //
	DiscoveryType          string   `json:"discoveryType,omitempty"`          //
	EnablePasswordList     string   `json:"enablePasswordList,omitempty"`     //
	GlobalCredentialIDList []string `json:"globalCredentialIdList,omitempty"` //
	HTTPReadCredential     struct {
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
	} `json:"httpReadCredential,omitempty"` //
	HTTPWriteCredential struct {
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
	} `json:"httpWriteCredential,omitempty"` //
	ID                    string `json:"id,omitempty"`                    //
	IPAddressList         string `json:"ipAddressList,omitempty"`         //
	IPFilterList          string `json:"ipFilterList,omitempty"`          //
	IsAutoCdp             bool   `json:"isAutoCdp,omitempty"`             //
	LldpLevel             int    `json:"lldpLevel,omitempty"`             //
	Name                  string `json:"name,omitempty"`                  //
	NetconfPort           string `json:"netconfPort,omitempty"`           //
	NumDevices            int    `json:"numDevices,omitempty"`            //
	ParentDiscoveryID     string `json:"parentDiscoveryId,omitempty"`     //
	PasswordList          string `json:"passwordList,omitempty"`          //
	PreferredMgmtIPMethod string `json:"preferredMgmtIPMethod,omitempty"` //
	ProtocolOrder         string `json:"protocolOrder,omitempty"`         //
	RetryCount            int    `json:"retryCount,omitempty"`            //
	SNMPAuthPassphrase    string `json:"snmpAuthPassphrase,omitempty"`    //
	SNMPAuthProtocol      string `json:"snmpAuthProtocol,omitempty"`      //
	SNMPMode              string `json:"snmpMode,omitempty"`              //
	SNMPPrivPassphrase    string `json:"snmpPrivPassphrase,omitempty"`    //
	SNMPPrivProtocol      string `json:"snmpPrivProtocol,omitempty"`      //
	SNMPRoCommunity       string `json:"snmpRoCommunity,omitempty"`       //
	SNMPRoCommunityDesc   string `json:"snmpRoCommunityDesc,omitempty"`   //
	SNMPRwCommunity       string `json:"snmpRwCommunity,omitempty"`       //
	SNMPRwCommunityDesc   string `json:"snmpRwCommunityDesc,omitempty"`   //
	SNMPUserName          string `json:"snmpUserName,omitempty"`          //
	TimeOut               int    `json:"timeOut,omitempty"`               //
	UpdateMgmtIP          bool   `json:"updateMgmtIp,omitempty"`          //
	UserNameList          string `json:"userNameList,omitempty"`          //
}

// CreateCLICredentialsResponse is the CreateCLICredentialsResponse definition
type CreateCLICredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateHTTPReadCredentialsResponse is the CreateHTTPReadCredentialsResponse definition
type CreateHTTPReadCredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateHTTPWriteCredentialsResponse is the CreateHTTPWriteCredentialsResponse definition
type CreateHTTPWriteCredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateNetconfCredentialsResponse is the CreateNetconfCredentialsResponse definition
type CreateNetconfCredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateSNMPReadCommunityResponse is the CreateSNMPReadCommunityResponse definition
type CreateSNMPReadCommunityResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateSNMPWriteCommunityResponse is the CreateSNMPWriteCommunityResponse definition
type CreateSNMPWriteCommunityResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateSNMPv3CredentialsResponse is the CreateSNMPv3CredentialsResponse definition
type CreateSNMPv3CredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateUpdateSNMPPropertiesResponse is the CreateUpdateSNMPPropertiesResponse definition
type CreateUpdateSNMPPropertiesResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteAllDiscoveryResponse is the DeleteAllDiscoveryResponse definition
type DeleteAllDiscoveryResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteDiscoveryByIDResponse is the DeleteDiscoveryByIdResponse definition
type DeleteDiscoveryByIDResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteDiscoveryBySpecifiedRangeResponse is the DeleteDiscoveryBySpecifiedRangeResponse definition
type DeleteDiscoveryBySpecifiedRangeResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteGlobalCredentialsByIDResponse is the DeleteGlobalCredentialsByIdResponse definition
type DeleteGlobalCredentialsByIDResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetCountOfAllDiscoveryJobsResponse is the GetCountOfAllDiscoveryJobsResponse definition
type GetCountOfAllDiscoveryJobsResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetCredentialSubTypeByCredentialIDResponse is the GetCredentialSubTypeByCredentialIdResponse definition
type GetCredentialSubTypeByCredentialIDResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDevicesDiscoveredByIDResponse is the GetDevicesDiscoveredByIdResponse definition
type GetDevicesDiscoveredByIDResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetDiscoveredDevicesByRangeResponse is the GetDiscoveredDevicesByRangeResponse definition
type GetDiscoveredDevicesByRangeResponse struct {
	Response []struct {
		AnchorWlcForAp              string `json:"anchorWlcForAp,omitempty"`              //
		AuthModelID                 string `json:"authModelId,omitempty"`                 //
		AvgUpdateFrequency          int    `json:"avgUpdateFrequency,omitempty"`          //
		BootDateTime                string `json:"bootDateTime,omitempty"`                //
		CLIStatus                   string `json:"cliStatus,omitempty"`                   //
		DuplicateDeviceID           string `json:"duplicateDeviceId,omitempty"`           //
		ErrorCode                   string `json:"errorCode,omitempty"`                   //
		ErrorDescription            string `json:"errorDescription,omitempty"`            //
		Family                      string `json:"family,omitempty"`                      //
		Hostname                    string `json:"hostname,omitempty"`                    //
		HTTPStatus                  string `json:"httpStatus,omitempty"`                  //
		ID                          string `json:"id,omitempty"`                          //
		ImageName                   string `json:"imageName,omitempty"`                   //
		IngressQueueConfig          string `json:"ingressQueueConfig,omitempty"`          //
		InterfaceCount              string `json:"interfaceCount,omitempty"`              //
		InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   //
		InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` //
		LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
		LineCardCount               string `json:"lineCardCount,omitempty"`               //
		LineCardID                  string `json:"lineCardId,omitempty"`                  //
		Location                    string `json:"location,omitempty"`                    //
		LocationName                string `json:"locationName,omitempty"`                //
		MacAddress                  string `json:"macAddress,omitempty"`                  //
		ManagementIPAddress         string `json:"managementIpAddress,omitempty"`         //
		MemorySize                  string `json:"memorySize,omitempty"`                  //
		NetconfStatus               string `json:"netconfStatus,omitempty"`               //
		NumUpdates                  int    `json:"numUpdates,omitempty"`                  //
		PingStatus                  string `json:"pingStatus,omitempty"`                  //
		PlatformID                  string `json:"platformId,omitempty"`                  //
		PortRange                   string `json:"portRange,omitempty"`                   //
		QosStatus                   string `json:"qosStatus,omitempty"`                   //
		ReachabilityFailureReason   string `json:"reachabilityFailureReason,omitempty"`   //
		ReachabilityStatus          string `json:"reachabilityStatus,omitempty"`          //
		Role                        string `json:"role,omitempty"`                        //
		RoleSource                  string `json:"roleSource,omitempty"`                  //
		SerialNumber                string `json:"serialNumber,omitempty"`                //
		SNMPContact                 string `json:"snmpContact,omitempty"`                 //
		SNMPLocation                string `json:"snmpLocation,omitempty"`                //
		SNMPStatus                  string `json:"snmpStatus,omitempty"`                  //
		SoftwareVersion             string `json:"softwareVersion,omitempty"`             //
		Tag                         string `json:"tag,omitempty"`                         //
		TagCount                    int    `json:"tagCount,omitempty"`                    //
		Type                        string `json:"type,omitempty"`                        //
		UpTime                      string `json:"upTime,omitempty"`                      //
		Vendor                      string `json:"vendor,omitempty"`                      //
		WlcApDeviceStatus           string `json:"wlcApDeviceStatus,omitempty"`           //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDiscoveredNetworkDevicesByDiscoveryIDResponse is the GetDiscoveredNetworkDevicesByDiscoveryIdResponse definition
type GetDiscoveredNetworkDevicesByDiscoveryIDResponse struct {
	Response []struct {
		AnchorWlcForAp              string `json:"anchorWlcForAp,omitempty"`              //
		AuthModelID                 string `json:"authModelId,omitempty"`                 //
		AvgUpdateFrequency          int    `json:"avgUpdateFrequency,omitempty"`          //
		BootDateTime                string `json:"bootDateTime,omitempty"`                //
		CLIStatus                   string `json:"cliStatus,omitempty"`                   //
		DuplicateDeviceID           string `json:"duplicateDeviceId,omitempty"`           //
		ErrorCode                   string `json:"errorCode,omitempty"`                   //
		ErrorDescription            string `json:"errorDescription,omitempty"`            //
		Family                      string `json:"family,omitempty"`                      //
		Hostname                    string `json:"hostname,omitempty"`                    //
		HTTPStatus                  string `json:"httpStatus,omitempty"`                  //
		ID                          string `json:"id,omitempty"`                          //
		ImageName                   string `json:"imageName,omitempty"`                   //
		IngressQueueConfig          string `json:"ingressQueueConfig,omitempty"`          //
		InterfaceCount              string `json:"interfaceCount,omitempty"`              //
		InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   //
		InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` //
		LastUpdated                 string `json:"lastUpdated,omitempty"`                 //
		LineCardCount               string `json:"lineCardCount,omitempty"`               //
		LineCardID                  string `json:"lineCardId,omitempty"`                  //
		Location                    string `json:"location,omitempty"`                    //
		LocationName                string `json:"locationName,omitempty"`                //
		MacAddress                  string `json:"macAddress,omitempty"`                  //
		ManagementIPAddress         string `json:"managementIpAddress,omitempty"`         //
		MemorySize                  string `json:"memorySize,omitempty"`                  //
		NetconfStatus               string `json:"netconfStatus,omitempty"`               //
		NumUpdates                  int    `json:"numUpdates,omitempty"`                  //
		PingStatus                  string `json:"pingStatus,omitempty"`                  //
		PlatformID                  string `json:"platformId,omitempty"`                  //
		PortRange                   string `json:"portRange,omitempty"`                   //
		QosStatus                   string `json:"qosStatus,omitempty"`                   //
		ReachabilityFailureReason   string `json:"reachabilityFailureReason,omitempty"`   //
		ReachabilityStatus          string `json:"reachabilityStatus,omitempty"`          //
		Role                        string `json:"role,omitempty"`                        //
		RoleSource                  string `json:"roleSource,omitempty"`                  //
		SerialNumber                string `json:"serialNumber,omitempty"`                //
		SNMPContact                 string `json:"snmpContact,omitempty"`                 //
		SNMPLocation                string `json:"snmpLocation,omitempty"`                //
		SNMPStatus                  string `json:"snmpStatus,omitempty"`                  //
		SoftwareVersion             string `json:"softwareVersion,omitempty"`             //
		Tag                         string `json:"tag,omitempty"`                         //
		TagCount                    int    `json:"tagCount,omitempty"`                    //
		Type                        string `json:"type,omitempty"`                        //
		UpTime                      string `json:"upTime,omitempty"`                      //
		Vendor                      string `json:"vendor,omitempty"`                      //
		WlcApDeviceStatus           string `json:"wlcApDeviceStatus,omitempty"`           //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDiscoveriesByRangeResponse is the GetDiscoveriesByRangeResponse definition
type GetDiscoveriesByRangeResponse struct {
	Response []struct {
		AttributeInfo          string   `json:"attributeInfo,omitempty"`          //
		CdpLevel               int      `json:"cdpLevel,omitempty"`               //
		DeviceIDs              string   `json:"deviceIds,omitempty"`              //
		DiscoveryCondition     string   `json:"discoveryCondition,omitempty"`     //
		DiscoveryStatus        string   `json:"discoveryStatus,omitempty"`        //
		DiscoveryType          string   `json:"discoveryType,omitempty"`          //
		EnablePasswordList     string   `json:"enablePasswordList,omitempty"`     //
		GlobalCredentialIDList []string `json:"globalCredentialIdList,omitempty"` //
		HTTPReadCredential     struct {
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
		} `json:"httpReadCredential,omitempty"` //
		HTTPWriteCredential struct {
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
		} `json:"httpWriteCredential,omitempty"` //
		ID                    string `json:"id,omitempty"`                    //
		IPAddressList         string `json:"ipAddressList,omitempty"`         //
		IPFilterList          string `json:"ipFilterList,omitempty"`          //
		IsAutoCdp             bool   `json:"isAutoCdp,omitempty"`             //
		LldpLevel             int    `json:"lldpLevel,omitempty"`             //
		Name                  string `json:"name,omitempty"`                  //
		NetconfPort           string `json:"netconfPort,omitempty"`           //
		NumDevices            int    `json:"numDevices,omitempty"`            //
		ParentDiscoveryID     string `json:"parentDiscoveryId,omitempty"`     //
		PasswordList          string `json:"passwordList,omitempty"`          //
		PreferredMgmtIPMethod string `json:"preferredMgmtIPMethod,omitempty"` //
		ProtocolOrder         string `json:"protocolOrder,omitempty"`         //
		RetryCount            int    `json:"retryCount,omitempty"`            //
		SNMPAuthPassphrase    string `json:"snmpAuthPassphrase,omitempty"`    //
		SNMPAuthProtocol      string `json:"snmpAuthProtocol,omitempty"`      //
		SNMPMode              string `json:"snmpMode,omitempty"`              //
		SNMPPrivPassphrase    string `json:"snmpPrivPassphrase,omitempty"`    //
		SNMPPrivProtocol      string `json:"snmpPrivProtocol,omitempty"`      //
		SNMPRoCommunity       string `json:"snmpRoCommunity,omitempty"`       //
		SNMPRoCommunityDesc   string `json:"snmpRoCommunityDesc,omitempty"`   //
		SNMPRwCommunity       string `json:"snmpRwCommunity,omitempty"`       //
		SNMPRwCommunityDesc   string `json:"snmpRwCommunityDesc,omitempty"`   //
		SNMPUserName          string `json:"snmpUserName,omitempty"`          //
		TimeOut               int    `json:"timeOut,omitempty"`               //
		UpdateMgmtIP          bool   `json:"updateMgmtIp,omitempty"`          //
		UserNameList          string `json:"userNameList,omitempty"`          //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDiscoveryByIDResponse is the GetDiscoveryByIdResponse definition
type GetDiscoveryByIDResponse struct {
	Response struct {
		AttributeInfo          string   `json:"attributeInfo,omitempty"`          //
		CdpLevel               int      `json:"cdpLevel,omitempty"`               //
		DeviceIDs              string   `json:"deviceIds,omitempty"`              //
		DiscoveryCondition     string   `json:"discoveryCondition,omitempty"`     //
		DiscoveryStatus        string   `json:"discoveryStatus,omitempty"`        //
		DiscoveryType          string   `json:"discoveryType,omitempty"`          //
		EnablePasswordList     string   `json:"enablePasswordList,omitempty"`     //
		GlobalCredentialIDList []string `json:"globalCredentialIdList,omitempty"` //
		HTTPReadCredential     struct {
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
		} `json:"httpReadCredential,omitempty"` //
		HTTPWriteCredential struct {
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
		} `json:"httpWriteCredential,omitempty"` //
		ID                    string `json:"id,omitempty"`                    //
		IPAddressList         string `json:"ipAddressList,omitempty"`         //
		IPFilterList          string `json:"ipFilterList,omitempty"`          //
		IsAutoCdp             bool   `json:"isAutoCdp,omitempty"`             //
		LldpLevel             int    `json:"lldpLevel,omitempty"`             //
		Name                  string `json:"name,omitempty"`                  //
		NetconfPort           string `json:"netconfPort,omitempty"`           //
		NumDevices            int    `json:"numDevices,omitempty"`            //
		ParentDiscoveryID     string `json:"parentDiscoveryId,omitempty"`     //
		PasswordList          string `json:"passwordList,omitempty"`          //
		PreferredMgmtIPMethod string `json:"preferredMgmtIPMethod,omitempty"` //
		ProtocolOrder         string `json:"protocolOrder,omitempty"`         //
		RetryCount            int    `json:"retryCount,omitempty"`            //
		SNMPAuthPassphrase    string `json:"snmpAuthPassphrase,omitempty"`    //
		SNMPAuthProtocol      string `json:"snmpAuthProtocol,omitempty"`      //
		SNMPMode              string `json:"snmpMode,omitempty"`              //
		SNMPPrivPassphrase    string `json:"snmpPrivPassphrase,omitempty"`    //
		SNMPPrivProtocol      string `json:"snmpPrivProtocol,omitempty"`      //
		SNMPRoCommunity       string `json:"snmpRoCommunity,omitempty"`       //
		SNMPRoCommunityDesc   string `json:"snmpRoCommunityDesc,omitempty"`   //
		SNMPRwCommunity       string `json:"snmpRwCommunity,omitempty"`       //
		SNMPRwCommunityDesc   string `json:"snmpRwCommunityDesc,omitempty"`   //
		SNMPUserName          string `json:"snmpUserName,omitempty"`          //
		TimeOut               int    `json:"timeOut,omitempty"`               //
		UpdateMgmtIP          bool   `json:"updateMgmtIp,omitempty"`          //
		UserNameList          string `json:"userNameList,omitempty"`          //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetDiscoveryJobsByIPResponse is the GetDiscoveryJobsByIPResponse definition
type GetDiscoveryJobsByIPResponse struct {
	Response []struct {
		AttributeInfo               string `json:"attributeInfo,omitempty"`               //
		CLIStatus                   string `json:"cliStatus,omitempty"`                   //
		DiscoveryStatus             string `json:"discoveryStatus,omitempty"`             //
		EndTime                     string `json:"endTime,omitempty"`                     //
		HTTPStatus                  string `json:"httpStatus,omitempty"`                  //
		ID                          string `json:"id,omitempty"`                          //
		InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   //
		InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` //
		IPAddress                   string `json:"ipAddress,omitempty"`                   //
		JobStatus                   string `json:"jobStatus,omitempty"`                   //
		Name                        string `json:"name,omitempty"`                        //
		NetconfStatus               string `json:"netconfStatus,omitempty"`               //
		PingStatus                  string `json:"pingStatus,omitempty"`                  //
		SNMPStatus                  string `json:"snmpStatus,omitempty"`                  //
		StartTime                   string `json:"startTime,omitempty"`                   //
		TaskID                      string `json:"taskId,omitempty"`                      //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetGlobalCredentialsResponse is the GetGlobalCredentialsResponse definition
type GetGlobalCredentialsResponse struct {
	Response []struct {
		Comments         string `json:"comments,omitempty"`         //
		CredentialType   string `json:"credentialType,omitempty"`   //
		Description      string `json:"description,omitempty"`      //
		ID               string `json:"id,omitempty"`               //
		InstanceTenantID string `json:"instanceTenantId,omitempty"` //
		InstanceUUID     string `json:"instanceUuid,omitempty"`     //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetListOfDiscoveriesByDiscoveryIDResponse is the GetListOfDiscoveriesByDiscoveryIdResponse definition
type GetListOfDiscoveriesByDiscoveryIDResponse struct {
	Response []struct {
		AttributeInfo               string `json:"attributeInfo,omitempty"`               //
		CLIStatus                   string `json:"cliStatus,omitempty"`                   //
		DiscoveryStatus             string `json:"discoveryStatus,omitempty"`             //
		EndTime                     string `json:"endTime,omitempty"`                     //
		HTTPStatus                  string `json:"httpStatus,omitempty"`                  //
		ID                          string `json:"id,omitempty"`                          //
		InventoryCollectionStatus   string `json:"inventoryCollectionStatus,omitempty"`   //
		InventoryReachabilityStatus string `json:"inventoryReachabilityStatus,omitempty"` //
		IPAddress                   string `json:"ipAddress,omitempty"`                   //
		JobStatus                   string `json:"jobStatus,omitempty"`                   //
		Name                        string `json:"name,omitempty"`                        //
		NetconfStatus               string `json:"netconfStatus,omitempty"`               //
		PingStatus                  string `json:"pingStatus,omitempty"`                  //
		SNMPStatus                  string `json:"snmpStatus,omitempty"`                  //
		StartTime                   string `json:"startTime,omitempty"`                   //
		TaskID                      string `json:"taskId,omitempty"`                      //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetNetworkDevicesFromDiscoveryResponse is the GetNetworkDevicesFromDiscoveryResponse definition
type GetNetworkDevicesFromDiscoveryResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetSNMPPropertiesResponse is the GetSNMPPropertiesResponse definition
type GetSNMPPropertiesResponse struct {
	Response []struct {
		ID                 string `json:"id,omitempty"`                 //
		InstanceTenantID   string `json:"instanceTenantId,omitempty"`   //
		InstanceUUID       string `json:"instanceUuid,omitempty"`       //
		IntValue           int    `json:"intValue,omitempty"`           //
		SystemPropertyName string `json:"systemPropertyName,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// StartDiscoveryResponse is the StartDiscoveryResponse definition
type StartDiscoveryResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateCLICredentialsResponse is the UpdateCLICredentialsResponse definition
type UpdateCLICredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateGlobalCredentialsResponse is the UpdateGlobalCredentialsResponse definition
type UpdateGlobalCredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateHTTPReadCredentialResponse is the UpdateHTTPReadCredentialResponse definition
type UpdateHTTPReadCredentialResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateHTTPWriteCredentialsResponse is the UpdateHTTPWriteCredentialsResponse definition
type UpdateHTTPWriteCredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateNetconfCredentialsResponse is the UpdateNetconfCredentialsResponse definition
type UpdateNetconfCredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateSNMPReadCommunityResponse is the UpdateSNMPReadCommunityResponse definition
type UpdateSNMPReadCommunityResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateSNMPWriteCommunityResponse is the UpdateSNMPWriteCommunityResponse definition
type UpdateSNMPWriteCommunityResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateSNMPv3CredentialsResponse is the UpdateSNMPv3CredentialsResponse definition
type UpdateSNMPv3CredentialsResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdatesAnExistingDiscoveryBySpecifiedIDResponse is the UpdatesAnExistingDiscoveryBySpecifiedIdResponse definition
type UpdatesAnExistingDiscoveryBySpecifiedIDResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateCLICredentials createCLICredentials
/* Adds global CLI credential
 */
func (s *DiscoveryService) CreateCLICredentials(createCLICredentialsRequest *CreateCLICredentialsRequest) (*CreateCLICredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/cli"

	response, err := RestyClient.R().
		SetBody(createCLICredentialsRequest).
		SetResult(&CreateCLICredentialsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateCLICredentialsResponse)
	return result, response, err
}

// CreateHTTPReadCredentials createHTTPReadCredentials
/* Adds HTTP read credentials
 */
func (s *DiscoveryService) CreateHTTPReadCredentials(createHTTPReadCredentialsRequest *CreateHTTPReadCredentialsRequest) (*CreateHTTPReadCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/http-read"

	response, err := RestyClient.R().
		SetBody(createHTTPReadCredentialsRequest).
		SetResult(&CreateHTTPReadCredentialsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateHTTPReadCredentialsResponse)
	return result, response, err
}

// CreateHTTPWriteCredentials createHTTPWriteCredentials
/* Adds global HTTP write credentials
 */
func (s *DiscoveryService) CreateHTTPWriteCredentials(createHTTPWriteCredentialsRequest *CreateHTTPWriteCredentialsRequest) (*CreateHTTPWriteCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/http-write"

	response, err := RestyClient.R().
		SetBody(createHTTPWriteCredentialsRequest).
		SetResult(&CreateHTTPWriteCredentialsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateHTTPWriteCredentialsResponse)
	return result, response, err
}

// CreateNetconfCredentials createNetconfCredentials
/* Adds global netconf credentials
 */
func (s *DiscoveryService) CreateNetconfCredentials(createNetconfCredentialsRequest *CreateNetconfCredentialsRequest) (*CreateNetconfCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/netconf"

	response, err := RestyClient.R().
		SetBody(createNetconfCredentialsRequest).
		SetResult(&CreateNetconfCredentialsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateNetconfCredentialsResponse)
	return result, response, err
}

// CreateSNMPReadCommunity createSNMPReadCommunity
/* Adds global SNMP read community
 */
func (s *DiscoveryService) CreateSNMPReadCommunity(createSNMPReadCommunityRequest *CreateSNMPReadCommunityRequest) (*CreateSNMPReadCommunityResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

	response, err := RestyClient.R().
		SetBody(createSNMPReadCommunityRequest).
		SetResult(&CreateSNMPReadCommunityResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateSNMPReadCommunityResponse)
	return result, response, err
}

// CreateSNMPWriteCommunity createSNMPWriteCommunity
/* Adds global SNMP write community
 */
func (s *DiscoveryService) CreateSNMPWriteCommunity(createSNMPWriteCommunityRequest *CreateSNMPWriteCommunityRequest) (*CreateSNMPWriteCommunityResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

	response, err := RestyClient.R().
		SetBody(createSNMPWriteCommunityRequest).
		SetResult(&CreateSNMPWriteCommunityResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateSNMPWriteCommunityResponse)
	return result, response, err
}

// CreateSNMPv3Credentials createSNMPv3Credentials
/* Adds global SNMPv3 credentials
 */
func (s *DiscoveryService) CreateSNMPv3Credentials(createSNMPv3CredentialsRequest *CreateSNMPv3CredentialsRequest) (*CreateSNMPv3CredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/snmpv3"

	response, err := RestyClient.R().
		SetBody(createSNMPv3CredentialsRequest).
		SetResult(&CreateSNMPv3CredentialsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateSNMPv3CredentialsResponse)
	return result, response, err
}

// CreateUpdateSNMPProperties createUpdateSNMPProperties
/* Adds SNMP properties
 */
func (s *DiscoveryService) CreateUpdateSNMPProperties(createUpdateSNMPPropertiesRequest *CreateUpdateSNMPPropertiesRequest) (*CreateUpdateSNMPPropertiesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/snmp-property"

	response, err := RestyClient.R().
		SetBody(createUpdateSNMPPropertiesRequest).
		SetResult(&CreateUpdateSNMPPropertiesResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateUpdateSNMPPropertiesResponse)
	return result, response, err
}

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

	path := "/dna/intent/api/v1/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{"+"globalCredentialId"+"}", fmt.Sprintf("%v", globalCredentialID), -1)

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
func (s *DiscoveryService) GetCountOfAllDiscoveryJobs() (*GetCountOfAllDiscoveryJobsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/count"

	response, err := RestyClient.R().
		SetResult(&GetCountOfAllDiscoveryJobsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetCountOfAllDiscoveryJobsResponse)
	return result, response, err
}

// GetCredentialSubTypeByCredentialID getCredentialSubTypeByCredentialId
/* Returns the credential sub type for the given Id
@param id Global Credential ID
*/
func (s *DiscoveryService) GetCredentialSubTypeByCredentialID(id string) (*GetCredentialSubTypeByCredentialIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetCredentialSubTypeByCredentialIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetCredentialSubTypeByCredentialIDResponse)
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
func (s *DiscoveryService) GetDevicesDiscoveredByID(id string, getDevicesDiscoveredByIDQueryParams *GetDevicesDiscoveredByIDQueryParams) (*GetDevicesDiscoveredByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/network-device/count"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getDevicesDiscoveredByIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDevicesDiscoveredByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDevicesDiscoveredByIDResponse)
	return result, response, err
}

// GetDiscoveredDevicesByRangeQueryParams defines the query parameters for this request
type GetDiscoveredDevicesByRangeQueryParams struct {
	TaskID string `url:"taskId,omitempty"` // taskId
}

// GetDiscoveredDevicesByRange getDiscoveredDevicesByRange
/* Returns the network devices discovered for the given discovery and for the given range. The maximum number of records that can be retrieved is 500. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
@param startIndex Start index
@param recordsToReturn Number of records to return
@param taskID taskId
*/
func (s *DiscoveryService) GetDiscoveredDevicesByRange(id string, startIndex int, recordsToReturn int, getDiscoveredDevicesByRangeQueryParams *GetDiscoveredDevicesByRangeQueryParams) (*GetDiscoveredDevicesByRangeResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/network-device/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	queryString, _ := query.Values(getDiscoveredDevicesByRangeQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDiscoveredDevicesByRangeResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDiscoveredDevicesByRangeResponse)
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
func (s *DiscoveryService) GetDiscoveredNetworkDevicesByDiscoveryID(id string, getDiscoveredNetworkDevicesByDiscoveryIDQueryParams *GetDiscoveredNetworkDevicesByDiscoveryIDQueryParams) (*GetDiscoveredNetworkDevicesByDiscoveryIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/network-device"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getDiscoveredNetworkDevicesByDiscoveryIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDiscoveredNetworkDevicesByDiscoveryIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDiscoveredNetworkDevicesByDiscoveryIDResponse)
	return result, response, err
}

// GetDiscoveriesByRange getDiscoveriesByRange
/* Returns the discovery by specified range
@param startIndex Start index
@param recordsToReturn Number of records to return
*/
func (s *DiscoveryService) GetDiscoveriesByRange(startIndex int, recordsToReturn int) (*GetDiscoveriesByRangeResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{startIndex}/{recordsToReturn}"
	path = strings.Replace(path, "{"+"startIndex"+"}", fmt.Sprintf("%v", startIndex), -1)
	path = strings.Replace(path, "{"+"recordsToReturn"+"}", fmt.Sprintf("%v", recordsToReturn), -1)

	response, err := RestyClient.R().
		SetResult(&GetDiscoveriesByRangeResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDiscoveriesByRangeResponse)
	return result, response, err
}

// GetDiscoveryByID getDiscoveryById
/* Returns discovery by Discovery ID. Discovery ID can be obtained using the "Get Discoveries by range" API.
@param id Discovery ID
*/
func (s *DiscoveryService) GetDiscoveryByID(id string) (*GetDiscoveryByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetDiscoveryByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDiscoveryByIDResponse)
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
func (s *DiscoveryService) GetDiscoveryJobsByIP(getDiscoveryJobsByIPQueryParams *GetDiscoveryJobsByIPQueryParams) (*GetDiscoveryJobsByIPResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/job"

	queryString, _ := query.Values(getDiscoveryJobsByIPQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDiscoveryJobsByIPResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetDiscoveryJobsByIPResponse)
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
func (s *DiscoveryService) GetGlobalCredentials(getGlobalCredentialsQueryParams *GetGlobalCredentialsQueryParams) (*GetGlobalCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential"

	queryString, _ := query.Values(getGlobalCredentialsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetGlobalCredentialsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetGlobalCredentialsResponse)
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
func (s *DiscoveryService) GetListOfDiscoveriesByDiscoveryID(id string, getListOfDiscoveriesByDiscoveryIDQueryParams *GetListOfDiscoveriesByDiscoveryIDQueryParams) (*GetListOfDiscoveriesByDiscoveryIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/job"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getListOfDiscoveriesByDiscoveryIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetListOfDiscoveriesByDiscoveryIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetListOfDiscoveriesByDiscoveryIDResponse)
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
	CLIStatus     []string `url:"cliStatus,omitempty"`     // cliStatus
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
func (s *DiscoveryService) GetNetworkDevicesFromDiscovery(id string, getNetworkDevicesFromDiscoveryQueryParams *GetNetworkDevicesFromDiscoveryQueryParams) (*GetNetworkDevicesFromDiscoveryResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery/{id}/summary"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getNetworkDevicesFromDiscoveryQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetNetworkDevicesFromDiscoveryResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetNetworkDevicesFromDiscoveryResponse)
	return result, response, err
}

// GetSNMPProperties getSNMPProperties
/* Returns SNMP properties
 */
func (s *DiscoveryService) GetSNMPProperties() (*GetSNMPPropertiesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/snmp-property"

	response, err := RestyClient.R().
		SetResult(&GetSNMPPropertiesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetSNMPPropertiesResponse)
	return result, response, err
}

// StartDiscovery startDiscovery
/* Initiates discovery with the given parameters
 */
func (s *DiscoveryService) StartDiscovery(startDiscoveryRequest *StartDiscoveryRequest) (*StartDiscoveryResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery"

	response, err := RestyClient.R().
		SetBody(startDiscoveryRequest).
		SetResult(&StartDiscoveryResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*StartDiscoveryResponse)
	return result, response, err
}

// UpdateCLICredentials updateCLICredentials
/* Updates global CLI credentials
 */
func (s *DiscoveryService) UpdateCLICredentials(updateCLICredentialsRequest *UpdateCLICredentialsRequest) (*UpdateCLICredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/cli"

	response, err := RestyClient.R().
		SetBody(updateCLICredentialsRequest).
		SetResult(&UpdateCLICredentialsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateCLICredentialsResponse)
	return result, response, err
}

// UpdateGlobalCredentials updateGlobalCredentials
/* Update global credential for network devices in site(s)
@param globalCredentialID Global credential Uuid
*/
func (s *DiscoveryService) UpdateGlobalCredentials(globalCredentialID string, updateGlobalCredentialsRequest *UpdateGlobalCredentialsRequest) (*UpdateGlobalCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/{globalCredentialId}"
	path = strings.Replace(path, "{"+"globalCredentialId"+"}", fmt.Sprintf("%v", globalCredentialID), -1)

	response, err := RestyClient.R().
		SetBody(updateGlobalCredentialsRequest).
		SetResult(&UpdateGlobalCredentialsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateGlobalCredentialsResponse)
	return result, response, err
}

// UpdateHTTPReadCredential updateHTTPReadCredential
/* Updates global HTTP Read credential
 */
func (s *DiscoveryService) UpdateHTTPReadCredential(updateHTTPReadCredentialRequest *UpdateHTTPReadCredentialRequest) (*UpdateHTTPReadCredentialResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/http-read"

	response, err := RestyClient.R().
		SetBody(updateHTTPReadCredentialRequest).
		SetResult(&UpdateHTTPReadCredentialResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateHTTPReadCredentialResponse)
	return result, response, err
}

// UpdateHTTPWriteCredentials updateHTTPWriteCredentials
/* Updates global HTTP write credentials
 */
func (s *DiscoveryService) UpdateHTTPWriteCredentials(updateHTTPWriteCredentialsRequest *UpdateHTTPWriteCredentialsRequest) (*UpdateHTTPWriteCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/http-write"

	response, err := RestyClient.R().
		SetBody(updateHTTPWriteCredentialsRequest).
		SetResult(&UpdateHTTPWriteCredentialsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateHTTPWriteCredentialsResponse)
	return result, response, err
}

// UpdateNetconfCredentials updateNetconfCredentials
/* Updates global netconf credentials
 */
func (s *DiscoveryService) UpdateNetconfCredentials(updateNetconfCredentialsRequest *UpdateNetconfCredentialsRequest) (*UpdateNetconfCredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/netconf"

	response, err := RestyClient.R().
		SetBody(updateNetconfCredentialsRequest).
		SetResult(&UpdateNetconfCredentialsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateNetconfCredentialsResponse)
	return result, response, err
}

// UpdateSNMPReadCommunity updateSNMPReadCommunity
/* Updates global SNMP read community
 */
func (s *DiscoveryService) UpdateSNMPReadCommunity(updateSNMPReadCommunityRequest *UpdateSNMPReadCommunityRequest) (*UpdateSNMPReadCommunityResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/snmpv2-read-community"

	response, err := RestyClient.R().
		SetBody(updateSNMPReadCommunityRequest).
		SetResult(&UpdateSNMPReadCommunityResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateSNMPReadCommunityResponse)
	return result, response, err
}

// UpdateSNMPWriteCommunity updateSNMPWriteCommunity
/* Updates global SNMP write community
 */
func (s *DiscoveryService) UpdateSNMPWriteCommunity(updateSNMPWriteCommunityRequest *UpdateSNMPWriteCommunityRequest) (*UpdateSNMPWriteCommunityResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/snmpv2-write-community"

	response, err := RestyClient.R().
		SetBody(updateSNMPWriteCommunityRequest).
		SetResult(&UpdateSNMPWriteCommunityResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateSNMPWriteCommunityResponse)
	return result, response, err
}

// UpdateSNMPv3Credentials updateSNMPv3Credentials
/* Updates global SNMPv3 credential
 */
func (s *DiscoveryService) UpdateSNMPv3Credentials(updateSNMPv3CredentialsRequest *UpdateSNMPv3CredentialsRequest) (*UpdateSNMPv3CredentialsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/global-credential/snmpv3"

	response, err := RestyClient.R().
		SetBody(updateSNMPv3CredentialsRequest).
		SetResult(&UpdateSNMPv3CredentialsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateSNMPv3CredentialsResponse)
	return result, response, err
}

// UpdatesAnExistingDiscoveryBySpecifiedID updatesAnExistingDiscoveryBySpecifiedId
/* Stops or starts an existing discovery
 */
func (s *DiscoveryService) UpdatesAnExistingDiscoveryBySpecifiedID(updatesAnExistingDiscoveryBySpecifiedIDRequest *UpdatesAnExistingDiscoveryBySpecifiedIDRequest) (*UpdatesAnExistingDiscoveryBySpecifiedIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/discovery"

	response, err := RestyClient.R().
		SetBody(updatesAnExistingDiscoveryBySpecifiedIDRequest).
		SetResult(&UpdatesAnExistingDiscoveryBySpecifiedIDResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdatesAnExistingDiscoveryBySpecifiedIDResponse)
	return result, response, err
}
