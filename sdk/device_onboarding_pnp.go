package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DeviceOnboardingPnPService is the service to communicate with the DeviceOnboardingPnP API endpoint
type DeviceOnboardingPnPService service

// AaaCredentials is the AaaCredentials definition
type AaaCredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// AddnDetails is the AddnDetails definition
type AddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ClaimDeviceRequest is the ClaimDeviceRequest definition
type ClaimDeviceRequest struct {
	ConfigFileURL     string            `json:"configFileUrl,omitempty"`     //
	ConfigID          string            `json:"configId,omitempty"`          //
	DeviceClaimList   []DeviceClaimList `json:"deviceClaimList,omitempty"`   //
	FileServiceID     string            `json:"fileServiceId,omitempty"`     //
	ImageID           string            `json:"imageId,omitempty"`           //
	ImageURL          string            `json:"imageUrl,omitempty"`          //
	PopulateInventory bool              `json:"populateInventory,omitempty"` //
	ProjectID         string            `json:"projectId,omitempty"`         //
	WorkflowID        string            `json:"workflowId,omitempty"`        //
}

// ConfigList is the ConfigList definition
type ConfigList struct {
	ConfigID         string             `json:"configId,omitempty"`         //
	ConfigParameters []ConfigParameters `json:"configParameters,omitempty"` //
}

// ConfigParameters is the ConfigParameters definition
type ConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// DayZeroConfig is the DayZeroConfig definition
type DayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// DefaultProfile is the DefaultProfile definition
type DefaultProfile struct {
	Cert          string   `json:"cert,omitempty"`          //
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` //
	IPAddresses   []string `json:"ipAddresses,omitempty"`   //
	Port          int      `json:"port,omitempty"`          //
	Proxy         bool     `json:"proxy,omitempty"`         //
}

// Device is the Device definition
type Device struct {
	ID                   string              `json:"_id,omitempty"`                  //
	DayZeroConfig        DayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string              `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           DeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []RunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  SystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       SystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string              `json:"tenantId,omitempty"`             //
	Version              int                 `json:"version,omitempty"`              //
	Workflow             Workflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   WorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// DeviceClaimList is the DeviceClaimList definition
type DeviceClaimList struct {
	ConfigList             []ConfigList `json:"configList,omitempty"`             //
	DeviceID               string       `json:"deviceId,omitempty"`               //
	LicenseLevel           string       `json:"licenseLevel,omitempty"`           //
	LicenseType            string       `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string       `json:"topOfStackSerialNumber,omitempty"` //
}

// DeviceInfo is the DeviceInfo definition
type DeviceInfo struct {
	AaaCredentials            AaaCredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   int                    `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string               `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                 `json:"agentType,omitempty"`                 //
	AuthStatus                string                 `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                 `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                 `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string               `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                 `json:"cmState,omitempty"`                   //
	Description               string                 `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string               `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                 `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string               `json:"featuresSupported,omitempty"`         //
	FileSystemList            []FileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              int                    `json:"firstContact,omitempty"`              //
	Hostname                  string                 `json:"hostname,omitempty"`                  //
	HTTPHeaders               []HTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                 `json:"imageFile,omitempty"`                 //
	ImageVersion              string                 `json:"imageVersion,omitempty"`              //
	IPInterfaces              []IPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               int                    `json:"lastContact,omitempty"`               //
	LastSyncTime              int                    `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              int                    `json:"lastUpdateOn,omitempty"`              //
	Location                  Location               `json:"location,omitempty"`                  //
	MacAddress                string                 `json:"macAddress,omitempty"`                //
	Mode                      string                 `json:"mode,omitempty"`                      //
	Name                      string                 `json:"name,omitempty"`                      //
	NeighborLinks             []NeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                 `json:"onbState,omitempty"`                  //
	Pid                       string                 `json:"pid,omitempty"`                       //
	PnpProfileList            []PnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                   `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []PreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                 `json:"projectId,omitempty"`                 //
	ProjectName               string                 `json:"projectName,omitempty"`               //
	ReloadRequested           bool                   `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                 `json:"serialNumber,omitempty"`              //
	SiteID                    string                 `json:"siteId,omitempty"`                    //
	SiteName                  string                 `json:"siteName,omitempty"`                  //
	SmartAccountID            string                 `json:"smartAccountId,omitempty"`            //
	Source                    string                 `json:"source,omitempty"`                    //
	Stack                     bool                   `json:"stack,omitempty"`                     //
	StackInfo                 StackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                 `json:"state,omitempty"`                     //
	SudiRequired              bool                   `json:"sudiRequired,omitempty"`              //
	Tags                      string                 `json:"tags,omitempty"`                      //
	UserMicNumbers            []string               `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string               `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                 `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                 `json:"workflowId,omitempty"`                //
	WorkflowName              string                 `json:"workflowName,omitempty"`              //
}

// DeviceResetList is the DeviceResetList definition
type DeviceResetList struct {
	ConfigList             []ConfigList `json:"configList,omitempty"`             //
	DeviceID               string       `json:"deviceId,omitempty"`               //
	LicenseLevel           string       `json:"licenseLevel,omitempty"`           //
	LicenseType            string       `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string       `json:"topOfStackSerialNumber,omitempty"` //
}

// FileSystemList is the FileSystemList definition
type FileSystemList struct {
	Freespace int    `json:"freespace,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	Readable  bool   `json:"readable,omitempty"`  //
	Size      int    `json:"size,omitempty"`      //
	Type      string `json:"type,omitempty"`      //
	Writeable bool   `json:"writeable,omitempty"` //
}

// HistoryTaskInfo is the HistoryTaskInfo definition
type HistoryTaskInfo struct {
	AddnDetails  []AddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string         `json:"name,omitempty"`         //
	TimeTaken    int            `json:"timeTaken,omitempty"`    //
	Type         string         `json:"type,omitempty"`         //
	WorkItemList []WorkItemList `json:"workItemList,omitempty"` //
}

// HTTPHeaders is the HttpHeaders definition
type HTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// IPInterfaces is the IpInterfaces definition
type IPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// DeviceOnBoardingPnPLocation is the Location definition
type DeviceOnBoardingPnPLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// NeighborLinks is the NeighborLinks definition
type NeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       //
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          //
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  //
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         //
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      //
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         //
	RemotePlatform           string `json:"remotePlatform,omitempty"`           //
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` //
	RemoteVersion            string `json:"remoteVersion,omitempty"`            //
}

// PnpProfileList is the PnpProfileList definition
type PnpProfileList struct {
	CreatedBy         string            `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool              `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   PrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string            `json:"profileName,omitempty"`       //
	SecondaryEndpoint SecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// PreWorkflowCliOuputs is the PreWorkflowCliOuputs definition
type PreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// PrimaryEndpoint is the PrimaryEndpoint definition
type PrimaryEndpoint struct {
	Certificate string `json:"certificate,omitempty"` //
	Fqdn        string `json:"fqdn,omitempty"`        //
	IPv4Address string `json:"ipv4Address,omitempty"` //
	IPv6Address string `json:"ipv6Address,omitempty"` //
	Port        int    `json:"port,omitempty"`        //
	Protocol    string `json:"protocol,omitempty"`    //
}

// Profile is the Profile definition
type Profile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault bool   `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        int    `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       bool   `json:"proxy,omitempty"`       //
}

// ResetRequest is the ResetRequest definition
type ResetRequest struct {
	DeviceResetList []DeviceResetList `json:"deviceResetList,omitempty"` //
	ProjectID       string            `json:"projectId,omitempty"`       //
	WorkflowID      string            `json:"workflowId,omitempty"`      //
}

// RunSummaryList is the RunSummaryList definition
type RunSummaryList struct {
	Details         string          `json:"details,omitempty"`         //
	ErrorFlag       bool            `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo HistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       int             `json:"timestamp,omitempty"`       //
}

// SAVAMapping is the SAVAMapping definition
type SAVAMapping struct {
	AutoSyncPeriod   int        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string     `json:"ccoUser,omitempty"`          //
	Expiry           int        `json:"expiry,omitempty"`           //
	LastSync         int        `json:"lastSync,omitempty"`         //
	Profile          Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string     `json:"syncStatus,omitempty"`       //
	TenantID         string     `json:"tenantId,omitempty"`         //
	Token            string     `json:"token,omitempty"`            //
	VirtualAccountID string     `json:"virtualAccountId,omitempty"` //
}

// SavaMappingList is the SavaMappingList definition
type SavaMappingList struct {
	AutoSyncPeriod   int        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string     `json:"ccoUser,omitempty"`          //
	Expiry           int        `json:"expiry,omitempty"`           //
	LastSync         int        `json:"lastSync,omitempty"`         //
	Profile          Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string     `json:"syncStatus,omitempty"`       //
	TenantID         string     `json:"tenantId,omitempty"`         //
	Token            string     `json:"token,omitempty"`            //
	VirtualAccountID string     `json:"virtualAccountId,omitempty"` //
}

// SecondaryEndpoint is the SecondaryEndpoint definition
type SecondaryEndpoint struct {
	Certificate string `json:"certificate,omitempty"` //
	Fqdn        string `json:"fqdn,omitempty"`        //
	IPv4Address string `json:"ipv4Address,omitempty"` //
	IPv6Address string `json:"ipv6Address,omitempty"` //
	Port        int    `json:"port,omitempty"`        //
	Protocol    string `json:"protocol,omitempty"`    //
}

// SiteProvisionRequest is the SiteProvisionRequest definition
type SiteProvisionRequest struct {
	DeviceID string `json:"deviceId,omitempty"` //
	SiteID   string `json:"siteId,omitempty"`   //
	Type     string `json:"type,omitempty"`     //
}

// StackInfo is the StackInfo definition
type StackInfo struct {
	IsFullRing             bool              `json:"isFullRing,omitempty"`             //
	StackMemberList        []StackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string            `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool              `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       int               `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string          `json:"validLicenseLevels,omitempty"`     //
}

// StackMemberList is the StackMemberList definition
type StackMemberList struct {
	HardwareVersion  string `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string `json:"licenseLevel,omitempty"`     //
	LicenseType      string `json:"licenseType,omitempty"`      //
	MacAddress       string `json:"macAddress,omitempty"`       //
	Pid              string `json:"pid,omitempty"`              //
	Priority         int    `json:"priority,omitempty"`         //
	Role             string `json:"role,omitempty"`             //
	SerialNumber     string `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string `json:"softwareVersion,omitempty"`  //
	StackNumber      int    `json:"stackNumber,omitempty"`      //
	State            string `json:"state,omitempty"`            //
	SudiSerialNumber string `json:"sudiSerialNumber,omitempty"` //
}

// SyncList is the SyncList definition
type SyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// SyncResult is the SyncResult definition
type SyncResult struct {
	SyncList []SyncList `json:"syncList,omitempty"` //
	SyncMsg  string     `json:"syncMsg,omitempty"`  //
}

// SystemResetWorkflow is the SystemResetWorkflow definition
type SystemResetWorkflow struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// SystemWorkflow is the SystemWorkflow definition
type SystemWorkflow struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// TaskTimeOuts is the TaskTimeOuts definition
type TaskTimeOuts struct {
	ConfigTimeOut        int `json:"configTimeOut,omitempty"`        //
	GeneralTimeOut       int `json:"generalTimeOut,omitempty"`       //
	ImageDownloadTimeOut int `json:"imageDownloadTimeOut,omitempty"` //
}

// Tasks is the Tasks definition
type Tasks struct {
	CurrWorkItemIDx int            `json:"currWorkItemIdx,omitempty"` //
	EndTime         int            `json:"endTime,omitempty"`         //
	Name            string         `json:"name,omitempty"`            //
	StartTime       int            `json:"startTime,omitempty"`       //
	State           string         `json:"state,omitempty"`           //
	TaskSeqNo       int            `json:"taskSeqNo,omitempty"`       //
	TimeTaken       int            `json:"timeTaken,omitempty"`       //
	Type            string         `json:"type,omitempty"`            //
	WorkItemList    []WorkItemList `json:"workItemList,omitempty"`    //
}

// UnclaimRequest is the UnclaimRequest definition
type UnclaimRequest struct {
	DeviceIDList []string `json:"deviceIdList,omitempty"` //
}

// WorkItemList is the WorkItemList definition
type WorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   int    `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime int    `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken int    `json:"timeTaken,omitempty"` //
}

// Workflow is the Workflow definition
type Workflow struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// WorkflowParameters is the WorkflowParameters definition
type WorkflowParameters struct {
	ConfigList             []ConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string       `json:"licenseLevel,omitempty"`           //
	LicenseType            string       `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string       `json:"topOfStackSerialNumber,omitempty"` //
}

// AddAWorkflowResponse is the AddAWorkflowResponse definition
type AddAWorkflowResponse struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// AddVirtualAccountResponse is the AddVirtualAccountResponse definition
type AddVirtualAccountResponse struct {
	AutoSyncPeriod   int        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string     `json:"ccoUser,omitempty"`          //
	Expiry           int        `json:"expiry,omitempty"`           //
	LastSync         int        `json:"lastSync,omitempty"`         //
	Profile          Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string     `json:"syncStatus,omitempty"`       //
	TenantID         string     `json:"tenantId,omitempty"`         //
	Token            string     `json:"token,omitempty"`            //
	VirtualAccountID string     `json:"virtualAccountId,omitempty"` //
}

// ClaimDeviceResponse is the ClaimDeviceResponse definition
type ClaimDeviceResponse struct {
	JSONArrayResponse []string `json:"jsonArrayResponse,omitempty"` //
	JSONResponse      string   `json:"jsonResponse,omitempty"`      //
	Message           string   `json:"message,omitempty"`           //
	StatusCode        int      `json:"statusCode,omitempty"`        //
}

// DayZeroConfigPreviewResult is the DayZeroConfigPreviewResult definition
type DayZeroConfigPreviewResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DeleteDeviceByIDFromPnPResponse is the DeleteDeviceByIdFromPnPResponse definition
type DeleteDeviceByIDFromPnPResponse struct {
	ID                   string              `json:"_id,omitempty"`                  //
	DayZeroConfig        DayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string              `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           DeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []RunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  SystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       SystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string              `json:"tenantId,omitempty"`             //
	Version              int                 `json:"version,omitempty"`              //
	Workflow             Workflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   WorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// DeleteWorkflowByIDResponse is the DeleteWorkflowByIdResponse definition
type DeleteWorkflowByIDResponse struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// DeregisterVirtualAccountResponse is the DeregisterVirtualAccountResponse definition
type DeregisterVirtualAccountResponse struct {
	AutoSyncPeriod   int        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string     `json:"ccoUser,omitempty"`          //
	Expiry           int        `json:"expiry,omitempty"`           //
	LastSync         int        `json:"lastSync,omitempty"`         //
	Profile          Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string     `json:"syncStatus,omitempty"`       //
	TenantID         string     `json:"tenantId,omitempty"`         //
	Token            string     `json:"token,omitempty"`            //
	VirtualAccountID string     `json:"virtualAccountId,omitempty"` //
}

// FailureList is the FailureList definition
type FailureList struct {
	ID        string `json:"id,omitempty"`        //
	Index     int    `json:"index,omitempty"`     //
	Msg       string `json:"msg,omitempty"`       //
	SerialNum string `json:"serialNum,omitempty"` //
}

// GetDeviceByIDResponse is the GetDeviceByIdResponse definition
type GetDeviceByIDResponse struct {
	ID                   string              `json:"_id,omitempty"`                  //
	DayZeroConfig        DayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string              `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           DeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []RunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  SystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       SystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string              `json:"tenantId,omitempty"`             //
	Version              int                 `json:"version,omitempty"`              //
	Workflow             Workflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   WorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// GetDeviceCountResponse is the GetDeviceCountResponse definition
type GetDeviceCountResponse struct {
	Response int `json:"response,omitempty"` //
}

// GetDeviceHistoryResponse is the GetDeviceHistoryResponse definition
type GetDeviceHistoryResponse struct {
	Response   []DeviceHistoryResponse `json:"response,omitempty"`   //
	StatusCode int                     `json:"statusCode,omitempty"` //
}

// GetPnPGlobalSettingsResponse is the GetPnPGlobalSettingsResponse definition
type GetPnPGlobalSettingsResponse struct {
	AaaCredentials  AaaCredentials    `json:"aaaCredentials,omitempty"`  //
	AcceptEula      bool              `json:"acceptEula,omitempty"`      //
	DefaultProfile  DefaultProfile    `json:"defaultProfile,omitempty"`  //
	ID              string            `json:"id,omitempty"`              //
	SavaMappingList []SavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    TaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string            `json:"tenantId,omitempty"`        //
	Version         int               `json:"version,omitempty"`         //
}

// GetSyncResultForVirtualAccountResponse is the GetSyncResultForVirtualAccountResponse definition
type GetSyncResultForVirtualAccountResponse struct {
	AutoSyncPeriod   int        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string     `json:"ccoUser,omitempty"`          //
	Expiry           int        `json:"expiry,omitempty"`           //
	LastSync         int        `json:"lastSync,omitempty"`         //
	Profile          Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string     `json:"syncStatus,omitempty"`       //
	TenantID         string     `json:"tenantId,omitempty"`         //
	Token            string     `json:"token,omitempty"`            //
	VirtualAccountID string     `json:"virtualAccountId,omitempty"` //
}

// GetWorkflowByIDResponse is the GetWorkflowByIdResponse definition
type GetWorkflowByIDResponse struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// GetWorkflowCountResponse is the GetWorkflowCountResponse definition
type GetWorkflowCountResponse struct {
	Response int `json:"response,omitempty"` //
}

// ImportDevicesInBulkResponse is the ImportDevicesInBulkResponse definition
type ImportDevicesInBulkResponse struct {
	FailureList []FailureList `json:"failureList,omitempty"` //
	SuccessList []SuccessList `json:"successList,omitempty"` //
}

// ResetDeviceResponse is the ResetDeviceResponse definition
type ResetDeviceResponse struct {
	JSONArrayResponse []string `json:"jsonArrayResponse,omitempty"` //
	JSONResponse      string   `json:"jsonResponse,omitempty"`      //
	Message           string   `json:"message,omitempty"`           //
	StatusCode        int      `json:"statusCode,omitempty"`        //
}

// DeviceHistoryResponse is the Response definition
type DeviceHistoryResponse struct {
	Complete      bool   `json:"complete,omitempty"`      //
	Config        string `json:"config,omitempty"`        //
	Error         bool   `json:"error,omitempty"`         //
	ErrorMessage  string `json:"errorMessage,omitempty"`  //
	ExpiredTime   int    `json:"expiredTime,omitempty"`   //
	RfProfile     string `json:"rfProfile,omitempty"`     //
	SensorProfile string `json:"sensorProfile,omitempty"` //
	SiteID        string `json:"siteId,omitempty"`        //
	StartTime     int    `json:"startTime,omitempty"`     //
	TaskID        string `json:"taskId,omitempty"`        //
}

// SiteClaimResponse is the SiteClaimResponse definition
type SiteClaimResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// SuccessList is the SuccessList definition
type SuccessList struct {
	ID                   string              `json:"_id,omitempty"`                  //
	DayZeroConfig        DayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string              `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           DeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []RunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  SystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       SystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string              `json:"tenantId,omitempty"`             //
	Version              int                 `json:"version,omitempty"`              //
	Workflow             Workflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   WorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// UnClaimDeviceResponse is the UnClaimDeviceResponse definition
type UnClaimDeviceResponse struct {
	JSONArrayResponse []string `json:"jsonArrayResponse,omitempty"` //
	JSONResponse      string   `json:"jsonResponse,omitempty"`      //
	Message           string   `json:"message,omitempty"`           //
	StatusCode        int      `json:"statusCode,omitempty"`        //
}

// UpdateDeviceResponse is the UpdateDeviceResponse definition
type UpdateDeviceResponse struct {
	ID                   string              `json:"_id,omitempty"`                  //
	DayZeroConfig        DayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string              `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           DeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []RunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  SystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       SystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string              `json:"tenantId,omitempty"`             //
	Version              int                 `json:"version,omitempty"`              //
	Workflow             Workflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   WorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// UpdatePnPGlobalSettingsResponse is the UpdatePnPGlobalSettingsResponse definition
type UpdatePnPGlobalSettingsResponse struct {
	AaaCredentials  AaaCredentials    `json:"aaaCredentials,omitempty"`  //
	AcceptEula      bool              `json:"acceptEula,omitempty"`      //
	DefaultProfile  DefaultProfile    `json:"defaultProfile,omitempty"`  //
	ID              string            `json:"id,omitempty"`              //
	SavaMappingList []SavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    TaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string            `json:"tenantId,omitempty"`        //
	Version         int               `json:"version,omitempty"`         //
}

// UpdatePnPServerProfileResponse is the UpdatePnPServerProfileResponse definition
type UpdatePnPServerProfileResponse struct {
	AutoSyncPeriod   int        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string     `json:"ccoUser,omitempty"`          //
	Expiry           int        `json:"expiry,omitempty"`           //
	LastSync         int        `json:"lastSync,omitempty"`         //
	Profile          Profile    `json:"profile,omitempty"`          //
	SmartAccountID   string     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string     `json:"syncStatus,omitempty"`       //
	TenantID         string     `json:"tenantId,omitempty"`         //
	Token            string     `json:"token,omitempty"`            //
	VirtualAccountID string     `json:"virtualAccountId,omitempty"` //
}

// UpdateWorkflowResponse is the UpdateWorkflowResponse definition
type UpdateWorkflowResponse struct {
	ID             string  `json:"_id,omitempty"`            //
	AddToInventory bool    `json:"addToInventory,omitempty"` //
	AddedOn        int     `json:"addedOn,omitempty"`        //
	ConfigID       string  `json:"configId,omitempty"`       //
	CurrTaskIDx    int     `json:"currTaskIdx,omitempty"`    //
	Description    string  `json:"description,omitempty"`    //
	EndTime        int     `json:"endTime,omitempty"`        //
	ExecTime       int     `json:"execTime,omitempty"`       //
	ImageID        string  `json:"imageId,omitempty"`        //
	InstanceType   string  `json:"instanceType,omitempty"`   //
	LastupdateOn   int     `json:"lastupdateOn,omitempty"`   //
	Name           string  `json:"name,omitempty"`           //
	StartTime      int     `json:"startTime,omitempty"`      //
	State          string  `json:"state,omitempty"`          //
	Tasks          []Tasks `json:"tasks,omitempty"`          //
	TenantID       string  `json:"tenantId,omitempty"`       //
	Type           string  `json:"type,omitempty"`           //
	UseState       string  `json:"useState,omitempty"`       //
	Version        int     `json:"version,omitempty"`        //
}

// AddAWorkflow addAWorkflow
/* Adds a PnP Workflow along with the relevant tasks in the workflow into the PnP database
 */
// func (s *DeviceOnboardingPnPService) AddAWorkflow(addAWorkflowRequest *AddAWorkflowRequest) (*AddAWorkflowResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

// 	response, err := RestyClient.R().
// 		SetBody(addAWorkflowRequest).
// 		SetResult(&AddAWorkflowResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*AddAWorkflowResponse)
// 	return result, response, err

//}

// AddDeviceToPnpDatabase addDeviceToPnpDatabase
/* Adds a device to the PnP database.
 */
// func (s *DeviceOnboardingPnPService) AddDeviceToPnpDatabase(addDeviceToPnpDatabaseRequest *AddDeviceToPnpDatabaseRequest) (*Device, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device"

// 	response, err := RestyClient.R().
// 		SetBody(addDeviceToPnpDatabaseRequest).
// 		SetResult(&Device{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*Device)
// 	return result, response, err

// }

// AddVirtualAccount addVirtualAccount
/* Registers a Smart Account, Virtual Account and the relevant server profile info with the PnP System & database. The devices present in the registered virtual account are synced with the PnP database as well. The response payload returns the new profile
 */
// func (s *DeviceOnboardingPnPService) AddVirtualAccount(addVirtualAccountRequest *AddVirtualAccountRequest) (*AddVirtualAccountResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

// 	response, err := RestyClient.R().
// 		SetBody(addVirtualAccountRequest).
// 		SetResult(&AddVirtualAccountResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*AddVirtualAccountResponse)
// 	return result, response, err

// }

// ClaimADeviceToASite claimADeviceToASite
/* Claim a device based on DNA-C Site based design process. Different parameters are required for different device platforms.
 */
// func (s *DeviceOnboardingPnPService) ClaimADeviceToASite(claimADeviceToASiteRequest *ClaimADeviceToASiteRequest) (*SiteClaimResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/site-claim"

// 	response, err := RestyClient.R().
// 		SetBody(claimADeviceToASiteRequest).
// 		SetResult(&SiteClaimResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*SiteClaimResponse)
// 	return result, response, err

// }

// ClaimDevice claimDevice
/* Claims one of more devices with specified workflow
 */
func (s *DeviceOnboardingPnPService) ClaimDevice(claimDeviceRequest *ClaimDeviceRequest) (*ClaimDeviceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/claim"

	response, err := RestyClient.R().
		SetBody(claimDeviceRequest).
		SetResult(&ClaimDeviceResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*ClaimDeviceResponse)
	return result, response, err

}

// DeleteDeviceByIDFromPnP deleteDeviceByIdFromPnP
/* Deletes specified device from PnP database
@param id id
*/
func (s *DeviceOnboardingPnPService) DeleteDeviceByIDFromPnP(id string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteWorkflowByID deleteWorkflowById
/* Deletes a workflow specified by id
@param id id
*/
func (s *DeviceOnboardingPnPService) DeleteWorkflowByID(id string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeregisterVirtualAccountQueryParams defines the query parameters for this request
type DeregisterVirtualAccountQueryParams struct {
	Domain string `url:"domain,omitempty"` // Smart Account Domain
	Name   string `url:"name,omitempty"`   // Virtual Account Name
}

// DeregisterVirtualAccount deregisterVirtualAccount
/* Deregisters the specified smart account & virtual account info and the associated device information from the PnP System & database. The devices associated with the deregistered virtual account are removed from the PnP database as well. The response payload contains the deregistered smart & virtual account information
@param domain Smart Account Domain
@param name Virtual Account Name
*/
func (s *DeviceOnboardingPnPService) DeregisterVirtualAccount(deregisterVirtualAccountQueryParams *DeregisterVirtualAccountQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings/vacct"

	queryString, _ := query.Values(deregisterVirtualAccountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetDeviceByID getDeviceById
/* Returns device details specified by device id
@param id id
*/
func (s *DeviceOnboardingPnPService) GetDeviceByID(id string) (*GetDeviceByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetDeviceByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDeviceByIDResponse)
	return result, response, err

}

// GetDeviceHistoryQueryParams defines the query parameters for this request
type GetDeviceHistoryQueryParams struct {
	SerialNumber string   `url:"serialNumber,omitempty"` // Device Serial Number
	Sort         []string `url:"sort,omitempty"`         // Comma seperated list of fields to sort on
	SortOrder    string   `url:"sortOrder,omitempty"`    // Sort Order Ascending (asc) or Descending (des)
}

// GetDeviceHistory getDeviceHistory
/* Returns history for a specific device. Serial int is a required parameter
@param serialNumber Device Serial Number
@param sort Comma seperated list of fields to sort on
@param sortOrder Sort Order Ascending (asc) or Descending (des)
*/
func (s *DeviceOnboardingPnPService) GetDeviceHistory(getDeviceHistoryQueryParams *GetDeviceHistoryQueryParams) (*GetDeviceHistoryResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/history"

	queryString, _ := query.Values(getDeviceHistoryQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceHistoryResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDeviceHistoryResponse)
	return result, response, err

}

// GetPnPGlobalSettings getPnPGlobalSettings
/* Returns global PnP settings of the user
 */
func (s *DeviceOnboardingPnPService) GetPnPGlobalSettings() (*GetPnPGlobalSettingsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings"

	response, err := RestyClient.R().
		SetResult(&GetPnPGlobalSettingsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetPnPGlobalSettingsResponse)
	return result, response, err

}

// GetPnpDeviceCountQueryParams defines the query parameters for this request
type GetPnpDeviceCountQueryParams struct {
	SerialNumber     []string `url:"serialNumber,omitempty"`     // Device Serial Number
	State            []string `url:"state,omitempty"`            // Device State
	OnbState         []string `url:"onbState,omitempty"`         // Device Onboarding State
	CmState          []string `url:"cmState,omitempty"`          // Device Connection Manager State
	Name             []string `url:"name,omitempty"`             // Device Name
	Pid              []string `url:"pid,omitempty"`              // Device ProductID
	Source           []string `url:"source,omitempty"`           // Device Source
	ProjectID        []string `url:"projectId,omitempty"`        // Device Project Id
	WorkflowID       []string `url:"workflowId,omitempty"`       // Device Workflow Id
	ProjectName      []string `url:"projectName,omitempty"`      // Device Project Name
	WorkflowName     []string `url:"workflowName,omitempty"`     // Device Workflow Name
	SmartAccountID   []string `url:"smartAccountId,omitempty"`   // Device Smart Account
	VirtualAccountID []string `url:"virtualAccountId,omitempty"` // Device Virtual Account
	LastContact      bool     `url:"lastContact,omitempty"`      // Device Has Contacted lastContact > 0
}

// GetPnpDeviceCount getPnpDeviceCount
/* Returns the device count based on filter criteria. This is useful for pagination
@param serialNumber Device Serial Number
@param state Device State
@param onbState Device Onboarding State
@param cmState Device Connection Manager State
@param name Device Name
@param pid Device ProductID
@param source Device Source
@param projectID Device Project Id
@param workflowID Device Workflow Id
@param projectName Device Project Name
@param workflowName Device Workflow Name
@param smartAccountID Device Smart Account
@param virtualAccountID Device Virtual Account
@param lastContact Device Has Contacted lastContact > 0
*/
func (s *DeviceOnboardingPnPService) GetPnpDeviceCount(getPnpDeviceCountQueryParams *GetPnpDeviceCountQueryParams) (*GetDeviceCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/count"

	queryString, _ := query.Values(getPnpDeviceCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetDeviceCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetDeviceCountResponse)
	return result, response, err

}

// GetPnpDeviceListQueryParams defines the query parameters for this request
type GetPnpDeviceListQueryParams struct {
	Limit            int      `url:"limit,omitempty"`            // Limits int of results
	Offset           int      `url:"offset,omitempty"`           // Index of first result
	Sort             []string `url:"sort,omitempty"`             // Comma seperated list of fields to sort on
	SortOrder        string   `url:"sortOrder,omitempty"`        // Sort Order Ascending (asc) or Descending (des)
	SerialNumber     []string `url:"serialNumber,omitempty"`     // Device Serial Number
	State            []string `url:"state,omitempty"`            // Device State
	OnbState         []string `url:"onbState,omitempty"`         // Device Onboarding State
	CmState          []string `url:"cmState,omitempty"`          // Device Connection Manager State
	Name             []string `url:"name,omitempty"`             // Device Name
	Pid              []string `url:"pid,omitempty"`              // Device ProductID
	Source           []string `url:"source,omitempty"`           // Device Source
	ProjectID        []string `url:"projectId,omitempty"`        // Device Project Id
	WorkflowID       []string `url:"workflowId,omitempty"`       // Device Workflow Id
	ProjectName      []string `url:"projectName,omitempty"`      // Device Project Name
	WorkflowName     []string `url:"workflowName,omitempty"`     // Device Workflow Name
	SmartAccountID   []string `url:"smartAccountId,omitempty"`   // Device Smart Account
	VirtualAccountID []string `url:"virtualAccountId,omitempty"` // Device Virtual Account
	LastContact      bool     `url:"lastContact,omitempty"`      // Device Has Contacted lastContact > 0
	MacAddress       string   `url:"macAddress,omitempty"`       // Device Mac Address
	Hostname         string   `url:"hostname,omitempty"`         // Device Hostname
	SiteName         string   `url:"siteName,omitempty"`         // Device Site Name
}

// GetPnpDeviceList getPnpDeviceList
/* Returns list of devices based on filter crieteria. If a limit is not specified, it will default to return 50 devices. Pagination and sorting are also supported by this endpoint
@param limit Limits int of results
@param offset Index of first result
@param sort Comma seperated list of fields to sort on
@param sortOrder Sort Order Ascending (asc) or Descending (des)
@param serialNumber Device Serial Number
@param state Device State
@param onbState Device Onboarding State
@param cmState Device Connection Manager State
@param name Device Name
@param pid Device ProductID
@param source Device Source
@param projectID Device Project Id
@param workflowID Device Workflow Id
@param projectName Device Project Name
@param workflowName Device Workflow Name
@param smartAccountID Device Smart Account
@param virtualAccountID Device Virtual Account
@param lastContact Device Has Contacted lastContact > 0
@param macAddress Device Mac Address
@param hostname Device Hostname
@param siteName Device Site Name
*/
func (s *DeviceOnboardingPnPService) GetPnpDeviceList(getPnpDeviceListQueryParams *GetPnpDeviceListQueryParams) (*Device, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device"

	queryString, _ := query.Values(getPnpDeviceListQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&Device{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*Device)
	return result, response, err

}

// GetSmartAccountList getSmartAccountList
/* Returns the list of Smart Account domains
 */
// func (s *DeviceOnboardingPnPService) GetSmartAccountList() (*GetSmartAccountListResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct"

// 	response, err := RestyClient.R().
// 		SetResult(&GetSmartAccountListResponse{}).
// 		SetError(&Error{}).
// 		Get(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*GetSmartAccountListResponse)
// 	return result, response, err

// }

// GetSyncResultForVirtualAccount getSyncResultForVirtualAccount
/* Returns the summary of devices synced from the given smart account & virtual account with PnP
@param domain Smart Account Domain
@param name Virtual Account Name
*/
func (s *DeviceOnboardingPnPService) GetSyncResultForVirtualAccount(domain string, name string) (*GetSyncResultForVirtualAccountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/sacct/{domain}/vacct/{name}/sync-result"
	path = strings.Replace(path, "{"+"domain"+"}", fmt.Sprintf("%v", domain), -1)
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	response, err := RestyClient.R().
		SetResult(&GetSyncResultForVirtualAccountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetSyncResultForVirtualAccountResponse)
	return result, response, err

}

// GetVirtualAccountList getVirtualAccountList
/* Returns list of virtual accounts associated with the specified smart account
@param domain Smart Account Domain
*/
// func (s *DeviceOnboardingPnPService) GetVirtualAccountList(domain string) (*GetVirtualAccountListResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct/{domain}/vacct"
// 	path = strings.Replace(path, "{"+"domain"+"}", fmt.Sprintf("%v", domain), -1)

// 	response, err := RestyClient.R().
// 		SetResult(&GetVirtualAccountListResponse{}).
// 		SetError(&Error{}).
// 		Get(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*GetVirtualAccountListResponse)
// 	return result, response, err

//}

// GetWorkflowByID getWorkflowById
/* Returns a workflow specified by id
@param id id
*/
// func (s *DeviceOnboardingPnPService) GetWorkflowByID(id string) (*GetWorkflowByIdResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
// 	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

// 	response, err := RestyClient.R().
// 		SetResult(&GetWorkflowByIDResponse{}).
// 		SetError(&Error{}).
// 		Get(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*GetWorkflowByIDResponse)
// 	return result, response, err

// }

// GetWorkflowCountQueryParams defines the query parameters for this request
type GetWorkflowCountQueryParams struct {
	Name []string `url:"name,omitempty"` // Workflow Name
}

// GetWorkflowCount getWorkflowCount
/* Returns the workflow count
@param name Workflow Name
*/
func (s *DeviceOnboardingPnPService) GetWorkflowCount(getWorkflowCountQueryParams *GetWorkflowCountQueryParams) (*GetWorkflowCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow/count"

	queryString, _ := query.Values(getWorkflowCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetWorkflowCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetWorkflowCountResponse)
	return result, response, err

}

// GetWorkflowsQueryParams defines the query parameters for this request
type GetWorkflowsQueryParams struct {
	Limit     int      `url:"limit,omitempty"`     // Limits int of results
	Offset    int      `url:"offset,omitempty"`    // Index of first result
	Sort      []string `url:"sort,omitempty"`      // Comma seperated lost of fields to sort on
	SortOrder string   `url:"sortOrder,omitempty"` // Sort Order Ascending (asc) or Descending (des)
	Type      []string `url:"type,omitempty"`      // Workflow Type
	Name      []string `url:"name,omitempty"`      // Workflow Name
}

// GetWorkflows getWorkflows
/* Returns the list of workflows based on filter criteria. If a limit is not specified, it will default to return 50 workflows. Pagination and sorting are also supported by this endpoint
@param limit Limits int of results
@param offset Index of first result
@param sort Comma seperated lost of fields to sort on
@param sortOrder Sort Order Ascending (asc) or Descending (des)
@param type Workflow Type
@param name Workflow Name
*/
// func (s *DeviceOnboardingPnPService) GetWorkflows(getWorkflowsQueryParams *GetWorkflowsQueryParams) (*GetWorkflowsResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

// 	queryString, _ := query.Values(getWorkflowsQueryParams)

// 	response, err := RestyClient.R().
// 		SetQueryString(queryString.Encode()).
// 		SetResult(&GetWorkflowsResponse{}).
// 		SetError(&Error{}).
// 		Get(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*GetWorkflowsResponse)
// 	return result, response, err

// }

// ImportDevicesInBulk importDevicesInBulk
/* Add devices to PnP in bulk
 */
// func (s *DeviceOnboardingPnPService) ImportDevicesInBulk(importDevicesInBulkRequest *ImportDevicesInBulkRequest) (*ImportDevicesInBulkResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/import"

// 	response, err := RestyClient.R().
// 		SetBody(importDevicesInBulkRequest).
// 		SetResult(&ImportDevicesInBulkResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*ImportDevicesInBulkResponse)
// 	return result, response, err

// }

// PreviewConfig previewConfig
/* Triggers a preview for site-based Day 0 Configuration
 */
// func (s *DeviceOnboardingPnPService) PreviewConfig(previewConfigRequest *PreviewConfigRequest) (*DayZeroConfigPreviewResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/site-config-preview"

// 	response, err := RestyClient.R().
// 		SetBody(previewConfigRequest).
// 		SetResult(&DayZeroConfigPreviewResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*DayZeroConfigPreviewResult)
// 	return result, response, err

// }

// ResetDevice resetDevice
/* Recovers a device from a Workflow Execution Error state
 */
// func (s *DeviceOnboardingPnPService) ResetDevice(resetDeviceRequest *ResetDeviceRequest) (*ResetDeviceResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/reset"

// 	response, err := RestyClient.R().
// 		SetBody(resetDeviceRequest).
// 		SetResult(&ResetDeviceResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*ResetDeviceResponse)
// 	return result, response, err

// }

// SyncVirtualAccountDevices syncVirtualAccountDevices
/* Synchronizes the device info from the given smart account & virtual account with the PnP database. The response payload returns a list of synced devices
 */
// func (s *DeviceOnboardingPnPService) SyncVirtualAccountDevices(syncVirtualAccountDevicesRequest *SyncVirtualAccountDevicesRequest) (*SyncVirtualAccountDevicesResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/vacct-sync"

// 	response, err := RestyClient.R().
// 		SetBody(syncVirtualAccountDevicesRequest).
// 		SetResult(&SyncVirtualAccountDevicesResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*SyncVirtualAccountDevicesResponse)
// 	return result, response, err

// }

// UnClaimDevice unClaimDevice
/* UnClaims one of more devices with specified workflow
 */
// func (s *DeviceOnboardingPnPService) UnClaimDevice(unClaimDeviceRequest *UnClaimDeviceRequest) (*UnClaimDeviceResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/unclaim"

// 	response, err := RestyClient.R().
// 		SetBody(unClaimDeviceRequest).
// 		SetResult(&UnClaimDeviceResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*UnClaimDeviceResponse)
// 	return result, response, err

// }

// UpdateDevice updateDevice
/* Updates device details specified by device id in PnP database
@param id id
*/
// func (s *DeviceOnboardingPnPService) UpdateDevice(id string, updateDeviceRequest *UpdateDeviceRequest) (*UpdateDeviceResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
// 	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

// 	response, err := RestyClient.R().
// 		SetBody(updateDeviceRequest).
// 		SetResult(&UpdateDeviceResponse{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*UpdateDeviceResponse)
// 	return result, response, err

// }

// UpdatePnPGlobalSettings updatePnPGlobalSettings
/* Updates the user's list of global PnP settings
 */
// func (s *DeviceOnboardingPnPService) UpdatePnPGlobalSettings(updatePnPGlobalSettingsRequest *UpdatePnPGlobalSettingsRequest) (*UpdatePnPGlobalSettingsResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-settings"

// 	response, err := RestyClient.R().
// 		SetBody(updatePnPGlobalSettingsRequest).
// 		SetResult(&UpdatePnPGlobalSettingsResponse{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*UpdatePnPGlobalSettingsResponse)
// 	return result, response, err

// }

// UpdatePnPServerProfile updatePnPServerProfile
/* Updates the PnP Server profile in a registered Virtual Account in the PnP database. The response payload returns the updated smart & virtual account info
 */
// func (s *DeviceOnboardingPnPService) UpdatePnPServerProfile(updatePnPServerProfileRequest *UpdatePnPServerProfileRequest) (*UpdatePnPServerProfileResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

// 	response, err := RestyClient.R().
// 		SetBody(updatePnPServerProfileRequest).
// 		SetResult(&UpdatePnPServerProfileResponse{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*UpdatePnPServerProfileResponse)
// 	return result, response, err

// }

// UpdateWorkflow updateWorkflow
/* Updates an existing workflow
@param id id
*/
// func (s *DeviceOnboardingPnPService) UpdateWorkflow(id string, updateWorkflowRequest *UpdateWorkflowRequest) (*UpdateWorkflowResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
// 	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

// 	response, err := RestyClient.R().
// 		SetBody(updateWorkflowRequest).
// 		SetResult(&UpdateWorkflowResponse{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*UpdateWorkflowResponse)
// 	return result, response, err

// }
