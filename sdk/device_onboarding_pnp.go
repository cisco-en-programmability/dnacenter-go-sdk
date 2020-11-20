package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DeviceOnboardingPnPService is the service to communicate with the DeviceOnboardingPnP API endpoint
type DeviceOnboardingPnPService service

// AddAWorkflowRequest is the addAWorkflowRequest definition
type AddAWorkflowRequest struct {
	TypeID         string                     `json:"_id,omitempty"`            //
	AddToInventory bool                       `json:"addToInventory,omitempty"` //
	AddedOn        int                        `json:"addedOn,omitempty"`        //
	ConfigID       string                     `json:"configId,omitempty"`       //
	CurrTaskIDx    int                        `json:"currTaskIdx,omitempty"`    //
	Description    string                     `json:"description,omitempty"`    //
	EndTime        int                        `json:"endTime,omitempty"`        //
	ExecTime       int                        `json:"execTime,omitempty"`       //
	ImageID        string                     `json:"imageId,omitempty"`        //
	InstanceType   string                     `json:"instanceType,omitempty"`   //
	LastupdateOn   int                        `json:"lastupdateOn,omitempty"`   //
	Name           string                     `json:"name,omitempty"`           //
	StartTime      int                        `json:"startTime,omitempty"`      //
	State          string                     `json:"state,omitempty"`          //
	Tasks          []AddAWorkflowRequestTasks `json:"tasks,omitempty"`          //
	TenantID       string                     `json:"tenantId,omitempty"`       //
	Type           string                     `json:"type,omitempty"`           //
	UseState       string                     `json:"useState,omitempty"`       //
	Version        int                        `json:"version,omitempty"`        //
}

// AddAWorkflowRequestTasks is the addAWorkflowRequestTasks definition
type AddAWorkflowRequestTasks struct {
	CurrWorkItemIDx int                                    `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                    `json:"endTime,omitempty"`         //
	Name            string                                 `json:"name,omitempty"`            //
	StartTime       int                                    `json:"startTime,omitempty"`       //
	State           string                                 `json:"state,omitempty"`           //
	TaskSeqNo       int                                    `json:"taskSeqNo,omitempty"`       //
	TimeTaken       int                                    `json:"timeTaken,omitempty"`       //
	Type            string                                 `json:"type,omitempty"`            //
	WorkItemList    []AddAWorkflowRequestTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddAWorkflowRequestTasksWorkItemList is the addAWorkflowRequestTasksWorkItemList definition
type AddAWorkflowRequestTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   int    `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime int    `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken int    `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseRequest is the addDeviceToPnpDatabaseRequest definition
type AddDeviceToPnpDatabaseRequest struct {
	TypeID               string                                           `json:"_id,omitempty"`                  //
	DayZeroConfig        AddDeviceToPnpDatabaseRequestDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                           `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           AddDeviceToPnpDatabaseRequestDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []AddDeviceToPnpDatabaseRequestRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  AddDeviceToPnpDatabaseRequestSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       AddDeviceToPnpDatabaseRequestSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                           `json:"tenantId,omitempty"`             //
	Version              float64                                          `json:"version,omitempty"`              //
	Workflow             AddDeviceToPnpDatabaseRequestWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   AddDeviceToPnpDatabaseRequestWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// AddDeviceToPnpDatabaseRequestDayZeroConfig is the addDeviceToPnpDatabaseRequestDayZeroConfig definition
type AddDeviceToPnpDatabaseRequestDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfo is the addDeviceToPnpDatabaseRequestDeviceInfo definition
type AddDeviceToPnpDatabaseRequestDeviceInfo struct {
	AAACredentials            AddDeviceToPnpDatabaseRequestDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                                       `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                      `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                        `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                        `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                        `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                        `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                      `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                        `json:"cmState,omitempty"`                   //
	Description               string                                                        `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                      `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                        `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                      `json:"featuresSupported,omitempty"`         //
	FileSystemList            []AddDeviceToPnpDatabaseRequestDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                                       `json:"firstContact,omitempty"`              //
	Hostname                  string                                                        `json:"hostname,omitempty"`                  //
	HTTPHeaders               []AddDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                        `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                        `json:"imageVersion,omitempty"`              //
	IPInterfaces              []AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                                       `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                                       `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                                       `json:"lastUpdateOn,omitempty"`              //
	Location                  AddDeviceToPnpDatabaseRequestDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                        `json:"macAddress,omitempty"`                //
	Mode                      string                                                        `json:"mode,omitempty"`                      //
	Name                      string                                                        `json:"name,omitempty"`                      //
	NeighborLinks             []AddDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                        `json:"onbState,omitempty"`                  //
	Pid                       string                                                        `json:"pid,omitempty"`                       //
	PnpProfileList            []AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                          `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []AddDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                        `json:"projectId,omitempty"`                 //
	ProjectName               string                                                        `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                          `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                        `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                        `json:"siteId,omitempty"`                    //
	SiteName                  string                                                        `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                        `json:"smartAccountId,omitempty"`            //
	Source                    string                                                        `json:"source,omitempty"`                    //
	Stack                     bool                                                          `json:"stack,omitempty"`                     //
	StackInfo                 AddDeviceToPnpDatabaseRequestDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                        `json:"state,omitempty"`                     //
	SudiRequired              bool                                                          `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                        `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                                      `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                                      `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                        `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                        `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                        `json:"workflowName,omitempty"`              //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoAAACredentials is the addDeviceToPnpDatabaseRequestDeviceInfoAAACredentials definition
type AddDeviceToPnpDatabaseRequestDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoAddnMacAddrs is the addDeviceToPnpDatabaseRequestDeviceInfoAddnMacAddrs definition
type AddDeviceToPnpDatabaseRequestDeviceInfoAddnMacAddrs []string

// AddDeviceToPnpDatabaseRequestDeviceInfoCapabilitiesSupported is the addDeviceToPnpDatabaseRequestDeviceInfoCapabilitiesSupported definition
type AddDeviceToPnpDatabaseRequestDeviceInfoCapabilitiesSupported []string

// AddDeviceToPnpDatabaseRequestDeviceInfoDeviceSudiSerialNos is the addDeviceToPnpDatabaseRequestDeviceInfoDeviceSudiSerialNos definition
type AddDeviceToPnpDatabaseRequestDeviceInfoDeviceSudiSerialNos []string

// AddDeviceToPnpDatabaseRequestDeviceInfoFeaturesSupported is the addDeviceToPnpDatabaseRequestDeviceInfoFeaturesSupported definition
type AddDeviceToPnpDatabaseRequestDeviceInfoFeaturesSupported []string

// AddDeviceToPnpDatabaseRequestDeviceInfoFileSystemList is the addDeviceToPnpDatabaseRequestDeviceInfoFileSystemList definition
type AddDeviceToPnpDatabaseRequestDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders is the addDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders definition
type AddDeviceToPnpDatabaseRequestDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces is the addDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces definition
type AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfacesIPv6AddressList is the addDeviceToPnpDatabaseRequestDeviceInfoIPInterfacesIPv6AddressList definition
type AddDeviceToPnpDatabaseRequestDeviceInfoIPInterfacesIPv6AddressList []string

// AddDeviceToPnpDatabaseRequestDeviceInfoLocation is the addDeviceToPnpDatabaseRequestDeviceInfoLocation definition
type AddDeviceToPnpDatabaseRequestDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks is the addDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks definition
type AddDeviceToPnpDatabaseRequestDeviceInfoNeighborLinks struct {
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

// AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList is the addDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList definition
type AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                 `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                                   `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                 `json:"profileName,omitempty"`       //
	SecondaryEndpoint AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListPrimaryEndpoint is the addDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListPrimaryEndpoint definition
type AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListSecondaryEndpoint is the addDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListSecondaryEndpoint definition
type AddDeviceToPnpDatabaseRequestDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs is the addDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs definition
type AddDeviceToPnpDatabaseRequestDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoStackInfo is the addDeviceToPnpDatabaseRequestDeviceInfoStackInfo definition
type AddDeviceToPnpDatabaseRequestDeviceInfoStackInfo struct {
	IsFullRing             bool                                                              `json:"isFullRing,omitempty"`             //
	StackMemberList        []AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                            `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                              `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                           `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                          `json:"validLicenseLevels,omitempty"`     //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList is the addDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList definition
type AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoValidLicenseLevels is the addDeviceToPnpDatabaseRequestDeviceInfoStackInfoValidLicenseLevels definition
type AddDeviceToPnpDatabaseRequestDeviceInfoStackInfoValidLicenseLevels []string

// AddDeviceToPnpDatabaseRequestDeviceInfoUserMicNumbers is the addDeviceToPnpDatabaseRequestDeviceInfoUserMicNumbers definition
type AddDeviceToPnpDatabaseRequestDeviceInfoUserMicNumbers []string

// AddDeviceToPnpDatabaseRequestDeviceInfoUserSudiSerialNos is the addDeviceToPnpDatabaseRequestDeviceInfoUserSudiSerialNos definition
type AddDeviceToPnpDatabaseRequestDeviceInfoUserSudiSerialNos []string

// AddDeviceToPnpDatabaseRequestRunSummaryList is the addDeviceToPnpDatabaseRequestRunSummaryList definition
type AddDeviceToPnpDatabaseRequestRunSummaryList struct {
	Details         string                                                     `json:"details,omitempty"`         //
	ErrorFlag       bool                                                       `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                                    `json:"timestamp,omitempty"`       //
}

// AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfo is the addDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfo definition
type AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                   `json:"name,omitempty"`         //
	TimeTaken    float64                                                                  `json:"timeTaken,omitempty"`    //
	Type         string                                                                   `json:"type,omitempty"`         //
	WorkItemList []AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails is the addDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails definition
type AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList is the addDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList definition
type AddDeviceToPnpDatabaseRequestRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestSystemResetWorkflow is the addDeviceToPnpDatabaseRequestSystemResetWorkflow definition
type AddDeviceToPnpDatabaseRequestSystemResetWorkflow struct {
	TypeID         string                                                  `json:"_id,omitempty"`            //
	AddToInventory bool                                                    `json:"addToInventory,omitempty"` //
	AddedOn        float64                                                 `json:"addedOn,omitempty"`        //
	ConfigID       string                                                  `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                                 `json:"currTaskIdx,omitempty"`    //
	Description    string                                                  `json:"description,omitempty"`    //
	EndTime        int                                                     `json:"endTime,omitempty"`        //
	ExecTime       float64                                                 `json:"execTime,omitempty"`       //
	ImageID        string                                                  `json:"imageId,omitempty"`        //
	InstanceType   string                                                  `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                                 `json:"lastupdateOn,omitempty"`   //
	Name           string                                                  `json:"name,omitempty"`           //
	StartTime      int                                                     `json:"startTime,omitempty"`      //
	State          string                                                  `json:"state,omitempty"`          //
	Tasks          []AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                  `json:"tenantId,omitempty"`       //
	Type           string                                                  `json:"type,omitempty"`           //
	UseState       string                                                  `json:"useState,omitempty"`       //
	Version        float64                                                 `json:"version,omitempty"`        //
}

// AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasks is the addDeviceToPnpDatabaseRequestSystemResetWorkflowTasks definition
type AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                             `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                                 `json:"endTime,omitempty"`         //
	Name            string                                                              `json:"name,omitempty"`            //
	StartTime       int                                                                 `json:"startTime,omitempty"`       //
	State           string                                                              `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                             `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                             `json:"timeTaken,omitempty"`       //
	Type            string                                                              `json:"type,omitempty"`            //
	WorkItemList    []AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList is the addDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList definition
type AddDeviceToPnpDatabaseRequestSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestSystemWorkflow is the addDeviceToPnpDatabaseRequestSystemWorkflow definition
type AddDeviceToPnpDatabaseRequestSystemWorkflow struct {
	TypeID         string                                             `json:"_id,omitempty"`            //
	AddToInventory bool                                               `json:"addToInventory,omitempty"` //
	AddedOn        float64                                            `json:"addedOn,omitempty"`        //
	ConfigID       string                                             `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                            `json:"currTaskIdx,omitempty"`    //
	Description    string                                             `json:"description,omitempty"`    //
	EndTime        int                                                `json:"endTime,omitempty"`        //
	ExecTime       float64                                            `json:"execTime,omitempty"`       //
	ImageID        string                                             `json:"imageId,omitempty"`        //
	InstanceType   string                                             `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                            `json:"lastupdateOn,omitempty"`   //
	Name           string                                             `json:"name,omitempty"`           //
	StartTime      int                                                `json:"startTime,omitempty"`      //
	State          string                                             `json:"state,omitempty"`          //
	Tasks          []AddDeviceToPnpDatabaseRequestSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                             `json:"tenantId,omitempty"`       //
	Type           string                                             `json:"type,omitempty"`           //
	UseState       string                                             `json:"useState,omitempty"`       //
	Version        float64                                            `json:"version,omitempty"`        //
}

// AddDeviceToPnpDatabaseRequestSystemWorkflowTasks is the addDeviceToPnpDatabaseRequestSystemWorkflowTasks definition
type AddDeviceToPnpDatabaseRequestSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                        `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                            `json:"endTime,omitempty"`         //
	Name            string                                                         `json:"name,omitempty"`            //
	StartTime       int                                                            `json:"startTime,omitempty"`       //
	State           string                                                         `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                        `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                        `json:"timeTaken,omitempty"`       //
	Type            string                                                         `json:"type,omitempty"`            //
	WorkItemList    []AddDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList is the addDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList definition
type AddDeviceToPnpDatabaseRequestSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestWorkflow is the addDeviceToPnpDatabaseRequestWorkflow definition
type AddDeviceToPnpDatabaseRequestWorkflow struct {
	TypeID         string                                       `json:"_id,omitempty"`            //
	AddToInventory bool                                         `json:"addToInventory,omitempty"` //
	AddedOn        float64                                      `json:"addedOn,omitempty"`        //
	ConfigID       string                                       `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                      `json:"currTaskIdx,omitempty"`    //
	Description    string                                       `json:"description,omitempty"`    //
	EndTime        int                                          `json:"endTime,omitempty"`        //
	ExecTime       float64                                      `json:"execTime,omitempty"`       //
	ImageID        string                                       `json:"imageId,omitempty"`        //
	InstanceType   string                                       `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                      `json:"lastupdateOn,omitempty"`   //
	Name           string                                       `json:"name,omitempty"`           //
	StartTime      int                                          `json:"startTime,omitempty"`      //
	State          string                                       `json:"state,omitempty"`          //
	Tasks          []AddDeviceToPnpDatabaseRequestWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                       `json:"tenantId,omitempty"`       //
	Type           string                                       `json:"type,omitempty"`           //
	UseState       string                                       `json:"useState,omitempty"`       //
	Version        float64                                      `json:"version,omitempty"`        //
}

// AddDeviceToPnpDatabaseRequestWorkflowParameters is the addDeviceToPnpDatabaseRequestWorkflowParameters definition
type AddDeviceToPnpDatabaseRequestWorkflowParameters struct {
	ConfigList             []AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                      `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                      `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                      `json:"topOfStackSerialNumber,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList is the addDeviceToPnpDatabaseRequestWorkflowParametersConfigList definition
type AddDeviceToPnpDatabaseRequestWorkflowParametersConfigList struct {
	ConfigID         string                                                                      `json:"configId,omitempty"`         //
	ConfigParameters []AddDeviceToPnpDatabaseRequestWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestWorkflowParametersConfigListConfigParameters is the addDeviceToPnpDatabaseRequestWorkflowParametersConfigListConfigParameters definition
type AddDeviceToPnpDatabaseRequestWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// AddDeviceToPnpDatabaseRequestWorkflowTasks is the addDeviceToPnpDatabaseRequestWorkflowTasks definition
type AddDeviceToPnpDatabaseRequestWorkflowTasks struct {
	CurrWorkItemIDx float64                                                  `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                      `json:"endTime,omitempty"`         //
	Name            string                                                   `json:"name,omitempty"`            //
	StartTime       int                                                      `json:"startTime,omitempty"`       //
	State           string                                                   `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                  `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                  `json:"timeTaken,omitempty"`       //
	Type            string                                                   `json:"type,omitempty"`            //
	WorkItemList    []AddDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList is the addDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList definition
type AddDeviceToPnpDatabaseRequestWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddVirtualAccountRequest is the addVirtualAccountRequest definition
type AddVirtualAccountRequest struct {
	AutoSyncPeriod   int                                `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                             `json:"ccoUser,omitempty"`          //
	Expiry           int                                `json:"expiry,omitempty"`           //
	LastSync         int                                `json:"lastSync,omitempty"`         //
	Profile          AddVirtualAccountRequestProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                             `json:"smartAccountId,omitempty"`   //
	SyncResult       AddVirtualAccountRequestSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                             `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int                                `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                             `json:"syncStatus,omitempty"`       //
	TenantID         string                             `json:"tenantId,omitempty"`         //
	Token            string                             `json:"token,omitempty"`            //
	VirtualAccountID string                             `json:"virtualAccountId,omitempty"` //
}

// AddVirtualAccountRequestProfile is the addVirtualAccountRequestProfile definition
type AddVirtualAccountRequestProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault bool   `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        int    `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       bool   `json:"proxy,omitempty"`       //
}

// AddVirtualAccountRequestSyncResult is the addVirtualAccountRequestSyncResult definition
type AddVirtualAccountRequestSyncResult struct {
	SyncList []AddVirtualAccountRequestSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                       `json:"syncMsg,omitempty"`  //
}

// AddVirtualAccountRequestSyncResultSyncList is the addVirtualAccountRequestSyncResultSyncList definition
type AddVirtualAccountRequestSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// AddVirtualAccountRequestSyncResultSyncListDeviceSnList is the addVirtualAccountRequestSyncResultSyncListDeviceSnList definition
type AddVirtualAccountRequestSyncResultSyncListDeviceSnList []string

// ClaimADeviceToASiteRequest is the claimADeviceToASiteRequest definition
type ClaimADeviceToASiteRequest struct {
	DeviceID string `json:"deviceId,omitempty"` //
	SiteID   string `json:"siteId,omitempty"`   //
	Type     string `json:"type,omitempty"`     //
}

// ClaimDeviceRequest is the claimDeviceRequest definition
type ClaimDeviceRequest struct {
	ConfigFileURL     string                              `json:"configFileUrl,omitempty"`     //
	ConfigID          string                              `json:"configId,omitempty"`          //
	DeviceClaimList   []ClaimDeviceRequestDeviceClaimList `json:"deviceClaimList,omitempty"`   //
	FileServiceID     string                              `json:"fileServiceId,omitempty"`     //
	ImageID           string                              `json:"imageId,omitempty"`           //
	ImageURL          string                              `json:"imageUrl,omitempty"`          //
	PopulateInventory bool                                `json:"populateInventory,omitempty"` //
	ProjectID         string                              `json:"projectId,omitempty"`         //
	WorkflowID        string                              `json:"workflowId,omitempty"`        //
}

// ClaimDeviceRequestDeviceClaimList is the claimDeviceRequestDeviceClaimList definition
type ClaimDeviceRequestDeviceClaimList struct {
	ConfigList             []ClaimDeviceRequestDeviceClaimListConfigList `json:"configList,omitempty"`             //
	DeviceID               string                                        `json:"deviceId,omitempty"`               //
	LicenseLevel           string                                        `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                        `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                        `json:"topOfStackSerialNumber,omitempty"` //
}

// ClaimDeviceRequestDeviceClaimListConfigList is the claimDeviceRequestDeviceClaimListConfigList definition
type ClaimDeviceRequestDeviceClaimListConfigList struct {
	ConfigID         string                                                        `json:"configId,omitempty"`         //
	ConfigParameters []ClaimDeviceRequestDeviceClaimListConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// ClaimDeviceRequestDeviceClaimListConfigListConfigParameters is the claimDeviceRequestDeviceClaimListConfigListConfigParameters definition
type ClaimDeviceRequestDeviceClaimListConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkRequest is the importDevicesInBulkRequest definition
type ImportDevicesInBulkRequest struct {
	TypeID               string                                        `json:"_id,omitempty"`                  //
	DayZeroConfig        ImportDevicesInBulkRequestDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                        `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           ImportDevicesInBulkRequestDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []ImportDevicesInBulkRequestRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  ImportDevicesInBulkRequestSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       ImportDevicesInBulkRequestSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                        `json:"tenantId,omitempty"`             //
	Version              float64                                       `json:"version,omitempty"`              //
	Workflow             ImportDevicesInBulkRequestWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   ImportDevicesInBulkRequestWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// ImportDevicesInBulkRequestDayZeroConfig is the importDevicesInBulkRequestDayZeroConfig definition
type ImportDevicesInBulkRequestDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfo is the importDevicesInBulkRequestDeviceInfo definition
type ImportDevicesInBulkRequestDeviceInfo struct {
	AAACredentials            ImportDevicesInBulkRequestDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                                    `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                   `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                     `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                     `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                     `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                     `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                   `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                     `json:"cmState,omitempty"`                   //
	Description               string                                                     `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                   `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                     `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                   `json:"featuresSupported,omitempty"`         //
	FileSystemList            []ImportDevicesInBulkRequestDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                                    `json:"firstContact,omitempty"`              //
	Hostname                  string                                                     `json:"hostname,omitempty"`                  //
	HTTPHeaders               []ImportDevicesInBulkRequestDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                     `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                     `json:"imageVersion,omitempty"`              //
	IPInterfaces              []ImportDevicesInBulkRequestDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                                    `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                                    `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                                    `json:"lastUpdateOn,omitempty"`              //
	Location                  ImportDevicesInBulkRequestDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                     `json:"macAddress,omitempty"`                //
	Mode                      string                                                     `json:"mode,omitempty"`                      //
	Name                      string                                                     `json:"name,omitempty"`                      //
	NeighborLinks             []ImportDevicesInBulkRequestDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                     `json:"onbState,omitempty"`                  //
	Pid                       string                                                     `json:"pid,omitempty"`                       //
	PnpProfileList            []ImportDevicesInBulkRequestDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                       `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []ImportDevicesInBulkRequestDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                     `json:"projectId,omitempty"`                 //
	ProjectName               string                                                     `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                       `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                     `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                     `json:"siteId,omitempty"`                    //
	SiteName                  string                                                     `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                     `json:"smartAccountId,omitempty"`            //
	Source                    string                                                     `json:"source,omitempty"`                    //
	Stack                     bool                                                       `json:"stack,omitempty"`                     //
	StackInfo                 ImportDevicesInBulkRequestDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                     `json:"state,omitempty"`                     //
	SudiRequired              bool                                                       `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                     `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                                   `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                                   `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                     `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                     `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                     `json:"workflowName,omitempty"`              //
}

// ImportDevicesInBulkRequestDeviceInfoAAACredentials is the importDevicesInBulkRequestDeviceInfoAAACredentials definition
type ImportDevicesInBulkRequestDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfoAddnMacAddrs is the importDevicesInBulkRequestDeviceInfoAddnMacAddrs definition
type ImportDevicesInBulkRequestDeviceInfoAddnMacAddrs []string

// ImportDevicesInBulkRequestDeviceInfoCapabilitiesSupported is the importDevicesInBulkRequestDeviceInfoCapabilitiesSupported definition
type ImportDevicesInBulkRequestDeviceInfoCapabilitiesSupported []string

// ImportDevicesInBulkRequestDeviceInfoDeviceSudiSerialNos is the importDevicesInBulkRequestDeviceInfoDeviceSudiSerialNos definition
type ImportDevicesInBulkRequestDeviceInfoDeviceSudiSerialNos []string

// ImportDevicesInBulkRequestDeviceInfoFeaturesSupported is the importDevicesInBulkRequestDeviceInfoFeaturesSupported definition
type ImportDevicesInBulkRequestDeviceInfoFeaturesSupported []string

// ImportDevicesInBulkRequestDeviceInfoFileSystemList is the importDevicesInBulkRequestDeviceInfoFileSystemList definition
type ImportDevicesInBulkRequestDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfoHTTPHeaders is the importDevicesInBulkRequestDeviceInfoHTTPHeaders definition
type ImportDevicesInBulkRequestDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfoIPInterfaces is the importDevicesInBulkRequestDeviceInfoIPInterfaces definition
type ImportDevicesInBulkRequestDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// ImportDevicesInBulkRequestDeviceInfoIPInterfacesIPv6AddressList is the importDevicesInBulkRequestDeviceInfoIPInterfacesIPv6AddressList definition
type ImportDevicesInBulkRequestDeviceInfoIPInterfacesIPv6AddressList []string

// ImportDevicesInBulkRequestDeviceInfoLocation is the importDevicesInBulkRequestDeviceInfoLocation definition
type ImportDevicesInBulkRequestDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// ImportDevicesInBulkRequestDeviceInfoNeighborLinks is the importDevicesInBulkRequestDeviceInfoNeighborLinks definition
type ImportDevicesInBulkRequestDeviceInfoNeighborLinks struct {
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

// ImportDevicesInBulkRequestDeviceInfoPnpProfileList is the importDevicesInBulkRequestDeviceInfoPnpProfileList definition
type ImportDevicesInBulkRequestDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                              `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                                `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   ImportDevicesInBulkRequestDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                              `json:"profileName,omitempty"`       //
	SecondaryEndpoint ImportDevicesInBulkRequestDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfoPnpProfileListPrimaryEndpoint is the importDevicesInBulkRequestDeviceInfoPnpProfileListPrimaryEndpoint definition
type ImportDevicesInBulkRequestDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// ImportDevicesInBulkRequestDeviceInfoPnpProfileListSecondaryEndpoint is the importDevicesInBulkRequestDeviceInfoPnpProfileListSecondaryEndpoint definition
type ImportDevicesInBulkRequestDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// ImportDevicesInBulkRequestDeviceInfoPreWorkflowCliOuputs is the importDevicesInBulkRequestDeviceInfoPreWorkflowCliOuputs definition
type ImportDevicesInBulkRequestDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfoStackInfo is the importDevicesInBulkRequestDeviceInfoStackInfo definition
type ImportDevicesInBulkRequestDeviceInfoStackInfo struct {
	IsFullRing             bool                                                           `json:"isFullRing,omitempty"`             //
	StackMemberList        []ImportDevicesInBulkRequestDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                         `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                           `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                        `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                       `json:"validLicenseLevels,omitempty"`     //
}

// ImportDevicesInBulkRequestDeviceInfoStackInfoStackMemberList is the importDevicesInBulkRequestDeviceInfoStackInfoStackMemberList definition
type ImportDevicesInBulkRequestDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// ImportDevicesInBulkRequestDeviceInfoStackInfoValidLicenseLevels is the importDevicesInBulkRequestDeviceInfoStackInfoValidLicenseLevels definition
type ImportDevicesInBulkRequestDeviceInfoStackInfoValidLicenseLevels []string

// ImportDevicesInBulkRequestDeviceInfoUserMicNumbers is the importDevicesInBulkRequestDeviceInfoUserMicNumbers definition
type ImportDevicesInBulkRequestDeviceInfoUserMicNumbers []string

// ImportDevicesInBulkRequestDeviceInfoUserSudiSerialNos is the importDevicesInBulkRequestDeviceInfoUserSudiSerialNos definition
type ImportDevicesInBulkRequestDeviceInfoUserSudiSerialNos []string

// ImportDevicesInBulkRequestRunSummaryList is the importDevicesInBulkRequestRunSummaryList definition
type ImportDevicesInBulkRequestRunSummaryList struct {
	Details         string                                                  `json:"details,omitempty"`         //
	ErrorFlag       bool                                                    `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                                 `json:"timestamp,omitempty"`       //
}

// ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfo is the importDevicesInBulkRequestRunSummaryListHistoryTaskInfo definition
type ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                `json:"name,omitempty"`         //
	TimeTaken    float64                                                               `json:"timeTaken,omitempty"`    //
	Type         string                                                                `json:"type,omitempty"`         //
	WorkItemList []ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfoAddnDetails is the importDevicesInBulkRequestRunSummaryListHistoryTaskInfoAddnDetails definition
type ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfoWorkItemList is the importDevicesInBulkRequestRunSummaryListHistoryTaskInfoWorkItemList definition
type ImportDevicesInBulkRequestRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkRequestSystemResetWorkflow is the importDevicesInBulkRequestSystemResetWorkflow definition
type ImportDevicesInBulkRequestSystemResetWorkflow struct {
	TypeID         string                                               `json:"_id,omitempty"`            //
	AddToInventory bool                                                 `json:"addToInventory,omitempty"` //
	AddedOn        float64                                              `json:"addedOn,omitempty"`        //
	ConfigID       string                                               `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                              `json:"currTaskIdx,omitempty"`    //
	Description    string                                               `json:"description,omitempty"`    //
	EndTime        int                                                  `json:"endTime,omitempty"`        //
	ExecTime       float64                                              `json:"execTime,omitempty"`       //
	ImageID        string                                               `json:"imageId,omitempty"`        //
	InstanceType   string                                               `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                              `json:"lastupdateOn,omitempty"`   //
	Name           string                                               `json:"name,omitempty"`           //
	StartTime      int                                                  `json:"startTime,omitempty"`      //
	State          string                                               `json:"state,omitempty"`          //
	Tasks          []ImportDevicesInBulkRequestSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                               `json:"tenantId,omitempty"`       //
	Type           string                                               `json:"type,omitempty"`           //
	UseState       string                                               `json:"useState,omitempty"`       //
	Version        float64                                              `json:"version,omitempty"`        //
}

// ImportDevicesInBulkRequestSystemResetWorkflowTasks is the importDevicesInBulkRequestSystemResetWorkflowTasks definition
type ImportDevicesInBulkRequestSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                          `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                              `json:"endTime,omitempty"`         //
	Name            string                                                           `json:"name,omitempty"`            //
	StartTime       int                                                              `json:"startTime,omitempty"`       //
	State           string                                                           `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                          `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                          `json:"timeTaken,omitempty"`       //
	Type            string                                                           `json:"type,omitempty"`            //
	WorkItemList    []ImportDevicesInBulkRequestSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// ImportDevicesInBulkRequestSystemResetWorkflowTasksWorkItemList is the importDevicesInBulkRequestSystemResetWorkflowTasksWorkItemList definition
type ImportDevicesInBulkRequestSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkRequestSystemWorkflow is the importDevicesInBulkRequestSystemWorkflow definition
type ImportDevicesInBulkRequestSystemWorkflow struct {
	TypeID         string                                          `json:"_id,omitempty"`            //
	AddToInventory bool                                            `json:"addToInventory,omitempty"` //
	AddedOn        float64                                         `json:"addedOn,omitempty"`        //
	ConfigID       string                                          `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                         `json:"currTaskIdx,omitempty"`    //
	Description    string                                          `json:"description,omitempty"`    //
	EndTime        int                                             `json:"endTime,omitempty"`        //
	ExecTime       float64                                         `json:"execTime,omitempty"`       //
	ImageID        string                                          `json:"imageId,omitempty"`        //
	InstanceType   string                                          `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                         `json:"lastupdateOn,omitempty"`   //
	Name           string                                          `json:"name,omitempty"`           //
	StartTime      int                                             `json:"startTime,omitempty"`      //
	State          string                                          `json:"state,omitempty"`          //
	Tasks          []ImportDevicesInBulkRequestSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                          `json:"tenantId,omitempty"`       //
	Type           string                                          `json:"type,omitempty"`           //
	UseState       string                                          `json:"useState,omitempty"`       //
	Version        float64                                         `json:"version,omitempty"`        //
}

// ImportDevicesInBulkRequestSystemWorkflowTasks is the importDevicesInBulkRequestSystemWorkflowTasks definition
type ImportDevicesInBulkRequestSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                     `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                         `json:"endTime,omitempty"`         //
	Name            string                                                      `json:"name,omitempty"`            //
	StartTime       int                                                         `json:"startTime,omitempty"`       //
	State           string                                                      `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                     `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                     `json:"timeTaken,omitempty"`       //
	Type            string                                                      `json:"type,omitempty"`            //
	WorkItemList    []ImportDevicesInBulkRequestSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// ImportDevicesInBulkRequestSystemWorkflowTasksWorkItemList is the importDevicesInBulkRequestSystemWorkflowTasksWorkItemList definition
type ImportDevicesInBulkRequestSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkRequestWorkflow is the importDevicesInBulkRequestWorkflow definition
type ImportDevicesInBulkRequestWorkflow struct {
	TypeID         string                                    `json:"_id,omitempty"`            //
	AddToInventory bool                                      `json:"addToInventory,omitempty"` //
	AddedOn        float64                                   `json:"addedOn,omitempty"`        //
	ConfigID       string                                    `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                   `json:"currTaskIdx,omitempty"`    //
	Description    string                                    `json:"description,omitempty"`    //
	EndTime        int                                       `json:"endTime,omitempty"`        //
	ExecTime       float64                                   `json:"execTime,omitempty"`       //
	ImageID        string                                    `json:"imageId,omitempty"`        //
	InstanceType   string                                    `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                   `json:"lastupdateOn,omitempty"`   //
	Name           string                                    `json:"name,omitempty"`           //
	StartTime      int                                       `json:"startTime,omitempty"`      //
	State          string                                    `json:"state,omitempty"`          //
	Tasks          []ImportDevicesInBulkRequestWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                    `json:"tenantId,omitempty"`       //
	Type           string                                    `json:"type,omitempty"`           //
	UseState       string                                    `json:"useState,omitempty"`       //
	Version        float64                                   `json:"version,omitempty"`        //
}

// ImportDevicesInBulkRequestWorkflowParameters is the importDevicesInBulkRequestWorkflowParameters definition
type ImportDevicesInBulkRequestWorkflowParameters struct {
	ConfigList             []ImportDevicesInBulkRequestWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                   `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                   `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                   `json:"topOfStackSerialNumber,omitempty"` //
}

// ImportDevicesInBulkRequestWorkflowParametersConfigList is the importDevicesInBulkRequestWorkflowParametersConfigList definition
type ImportDevicesInBulkRequestWorkflowParametersConfigList struct {
	ConfigID         string                                                                   `json:"configId,omitempty"`         //
	ConfigParameters []ImportDevicesInBulkRequestWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// ImportDevicesInBulkRequestWorkflowParametersConfigListConfigParameters is the importDevicesInBulkRequestWorkflowParametersConfigListConfigParameters definition
type ImportDevicesInBulkRequestWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkRequestWorkflowTasks is the importDevicesInBulkRequestWorkflowTasks definition
type ImportDevicesInBulkRequestWorkflowTasks struct {
	CurrWorkItemIDx float64                                               `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                   `json:"endTime,omitempty"`         //
	Name            string                                                `json:"name,omitempty"`            //
	StartTime       int                                                   `json:"startTime,omitempty"`       //
	State           string                                                `json:"state,omitempty"`           //
	TaskSeqNo       float64                                               `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                               `json:"timeTaken,omitempty"`       //
	Type            string                                                `json:"type,omitempty"`            //
	WorkItemList    []ImportDevicesInBulkRequestWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// ImportDevicesInBulkRequestWorkflowTasksWorkItemList is the importDevicesInBulkRequestWorkflowTasksWorkItemList definition
type ImportDevicesInBulkRequestWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// PreviewConfigRequest is the previewConfigRequest definition
type PreviewConfigRequest struct {
	DeviceID string `json:"deviceId,omitempty"` //
	SiteID   string `json:"siteId,omitempty"`   //
	Type     string `json:"type,omitempty"`     //
}

// ResetDeviceRequest is the resetDeviceRequest definition
type ResetDeviceRequest struct {
	DeviceResetList []ResetDeviceRequestDeviceResetList `json:"deviceResetList,omitempty"` //
	ProjectID       string                              `json:"projectId,omitempty"`       //
	WorkflowID      string                              `json:"workflowId,omitempty"`      //
}

// ResetDeviceRequestDeviceResetList is the resetDeviceRequestDeviceResetList definition
type ResetDeviceRequestDeviceResetList struct {
	ConfigList             []ResetDeviceRequestDeviceResetListConfigList `json:"configList,omitempty"`             //
	DeviceID               string                                        `json:"deviceId,omitempty"`               //
	LicenseLevel           string                                        `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                        `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                        `json:"topOfStackSerialNumber,omitempty"` //
}

// ResetDeviceRequestDeviceResetListConfigList is the resetDeviceRequestDeviceResetListConfigList definition
type ResetDeviceRequestDeviceResetListConfigList struct {
	ConfigID         string                                                        `json:"configId,omitempty"`         //
	ConfigParameters []ResetDeviceRequestDeviceResetListConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// ResetDeviceRequestDeviceResetListConfigListConfigParameters is the resetDeviceRequestDeviceResetListConfigListConfigParameters definition
type ResetDeviceRequestDeviceResetListConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// SyncVirtualAccountDevicesRequest is the syncVirtualAccountDevicesRequest definition
type SyncVirtualAccountDevicesRequest struct {
	AutoSyncPeriod   int                                        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                     `json:"ccoUser,omitempty"`          //
	Expiry           int                                        `json:"expiry,omitempty"`           //
	LastSync         int                                        `json:"lastSync,omitempty"`         //
	Profile          SyncVirtualAccountDevicesRequestProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                     `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncVirtualAccountDevicesRequestSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int                                        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                     `json:"syncStatus,omitempty"`       //
	TenantID         string                                     `json:"tenantId,omitempty"`         //
	Token            string                                     `json:"token,omitempty"`            //
	VirtualAccountID string                                     `json:"virtualAccountId,omitempty"` //
}

// SyncVirtualAccountDevicesRequestProfile is the syncVirtualAccountDevicesRequestProfile definition
type SyncVirtualAccountDevicesRequestProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault bool   `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        int    `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       bool   `json:"proxy,omitempty"`       //
}

// SyncVirtualAccountDevicesRequestSyncResult is the syncVirtualAccountDevicesRequestSyncResult definition
type SyncVirtualAccountDevicesRequestSyncResult struct {
	SyncList []SyncVirtualAccountDevicesRequestSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                               `json:"syncMsg,omitempty"`  //
}

// SyncVirtualAccountDevicesRequestSyncResultSyncList is the syncVirtualAccountDevicesRequestSyncResultSyncList definition
type SyncVirtualAccountDevicesRequestSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// SyncVirtualAccountDevicesRequestSyncResultSyncListDeviceSnList is the syncVirtualAccountDevicesRequestSyncResultSyncListDeviceSnList definition
type SyncVirtualAccountDevicesRequestSyncResultSyncListDeviceSnList []string

// UnClaimDeviceRequest is the unClaimDeviceRequest definition
type UnClaimDeviceRequest struct {
	DeviceIDList []string `json:"deviceIdList,omitempty"` //
}

// UnClaimDeviceRequestDeviceIDList is the unClaimDeviceRequestDeviceIDList definition
type UnClaimDeviceRequestDeviceIDList []string

// UpdateDeviceRequest is the updateDeviceRequest definition
type UpdateDeviceRequest struct {
	TypeID               string                                 `json:"_id,omitempty"`                  //
	DayZeroConfig        UpdateDeviceRequestDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                 `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           UpdateDeviceRequestDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []UpdateDeviceRequestRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  UpdateDeviceRequestSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       UpdateDeviceRequestSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                 `json:"tenantId,omitempty"`             //
	Version              float64                                `json:"version,omitempty"`              //
	Workflow             UpdateDeviceRequestWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   UpdateDeviceRequestWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// UpdateDeviceRequestDayZeroConfig is the updateDeviceRequestDayZeroConfig definition
type UpdateDeviceRequestDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// UpdateDeviceRequestDeviceInfo is the updateDeviceRequestDeviceInfo definition
type UpdateDeviceRequestDeviceInfo struct {
	AAACredentials            UpdateDeviceRequestDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                             `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                            `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                              `json:"agentType,omitempty"`                 //
	AuthStatus                string                                              `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                              `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                              `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                            `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                              `json:"cmState,omitempty"`                   //
	Description               string                                              `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                            `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                              `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                            `json:"featuresSupported,omitempty"`         //
	FileSystemList            []UpdateDeviceRequestDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                             `json:"firstContact,omitempty"`              //
	Hostname                  string                                              `json:"hostname,omitempty"`                  //
	HTTPHeaders               []UpdateDeviceRequestDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                              `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                              `json:"imageVersion,omitempty"`              //
	IPInterfaces              []UpdateDeviceRequestDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                             `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                             `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                             `json:"lastUpdateOn,omitempty"`              //
	Location                  UpdateDeviceRequestDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                              `json:"macAddress,omitempty"`                //
	Mode                      string                                              `json:"mode,omitempty"`                      //
	Name                      string                                              `json:"name,omitempty"`                      //
	NeighborLinks             []UpdateDeviceRequestDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                              `json:"onbState,omitempty"`                  //
	Pid                       string                                              `json:"pid,omitempty"`                       //
	PnpProfileList            []UpdateDeviceRequestDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []UpdateDeviceRequestDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                              `json:"projectId,omitempty"`                 //
	ProjectName               string                                              `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                              `json:"serialNumber,omitempty"`              //
	SiteID                    string                                              `json:"siteId,omitempty"`                    //
	SiteName                  string                                              `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                              `json:"smartAccountId,omitempty"`            //
	Source                    string                                              `json:"source,omitempty"`                    //
	Stack                     bool                                                `json:"stack,omitempty"`                     //
	StackInfo                 UpdateDeviceRequestDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                              `json:"state,omitempty"`                     //
	SudiRequired              bool                                                `json:"sudiRequired,omitempty"`              //
	Tags                      string                                              `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                            `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                            `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                              `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                              `json:"workflowId,omitempty"`                //
	WorkflowName              string                                              `json:"workflowName,omitempty"`              //
}

// UpdateDeviceRequestDeviceInfoAAACredentials is the updateDeviceRequestDeviceInfoAAACredentials definition
type UpdateDeviceRequestDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// UpdateDeviceRequestDeviceInfoAddnMacAddrs is the updateDeviceRequestDeviceInfoAddnMacAddrs definition
type UpdateDeviceRequestDeviceInfoAddnMacAddrs []string

// UpdateDeviceRequestDeviceInfoCapabilitiesSupported is the updateDeviceRequestDeviceInfoCapabilitiesSupported definition
type UpdateDeviceRequestDeviceInfoCapabilitiesSupported []string

// UpdateDeviceRequestDeviceInfoDeviceSudiSerialNos is the updateDeviceRequestDeviceInfoDeviceSudiSerialNos definition
type UpdateDeviceRequestDeviceInfoDeviceSudiSerialNos []string

// UpdateDeviceRequestDeviceInfoFeaturesSupported is the updateDeviceRequestDeviceInfoFeaturesSupported definition
type UpdateDeviceRequestDeviceInfoFeaturesSupported []string

// UpdateDeviceRequestDeviceInfoFileSystemList is the updateDeviceRequestDeviceInfoFileSystemList definition
type UpdateDeviceRequestDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// UpdateDeviceRequestDeviceInfoHTTPHeaders is the updateDeviceRequestDeviceInfoHTTPHeaders definition
type UpdateDeviceRequestDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// UpdateDeviceRequestDeviceInfoIPInterfaces is the updateDeviceRequestDeviceInfoIPInterfaces definition
type UpdateDeviceRequestDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// UpdateDeviceRequestDeviceInfoIPInterfacesIPv6AddressList is the updateDeviceRequestDeviceInfoIPInterfacesIPv6AddressList definition
type UpdateDeviceRequestDeviceInfoIPInterfacesIPv6AddressList []string

// UpdateDeviceRequestDeviceInfoLocation is the updateDeviceRequestDeviceInfoLocation definition
type UpdateDeviceRequestDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// UpdateDeviceRequestDeviceInfoNeighborLinks is the updateDeviceRequestDeviceInfoNeighborLinks definition
type UpdateDeviceRequestDeviceInfoNeighborLinks struct {
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

// UpdateDeviceRequestDeviceInfoPnpProfileList is the updateDeviceRequestDeviceInfoPnpProfileList definition
type UpdateDeviceRequestDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                       `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                         `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   UpdateDeviceRequestDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                       `json:"profileName,omitempty"`       //
	SecondaryEndpoint UpdateDeviceRequestDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// UpdateDeviceRequestDeviceInfoPnpProfileListPrimaryEndpoint is the updateDeviceRequestDeviceInfoPnpProfileListPrimaryEndpoint definition
type UpdateDeviceRequestDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// UpdateDeviceRequestDeviceInfoPnpProfileListSecondaryEndpoint is the updateDeviceRequestDeviceInfoPnpProfileListSecondaryEndpoint definition
type UpdateDeviceRequestDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// UpdateDeviceRequestDeviceInfoPreWorkflowCliOuputs is the updateDeviceRequestDeviceInfoPreWorkflowCliOuputs definition
type UpdateDeviceRequestDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// UpdateDeviceRequestDeviceInfoStackInfo is the updateDeviceRequestDeviceInfoStackInfo definition
type UpdateDeviceRequestDeviceInfoStackInfo struct {
	IsFullRing             bool                                                    `json:"isFullRing,omitempty"`             //
	StackMemberList        []UpdateDeviceRequestDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                  `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                    `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                 `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                `json:"validLicenseLevels,omitempty"`     //
}

// UpdateDeviceRequestDeviceInfoStackInfoStackMemberList is the updateDeviceRequestDeviceInfoStackInfoStackMemberList definition
type UpdateDeviceRequestDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// UpdateDeviceRequestDeviceInfoStackInfoValidLicenseLevels is the updateDeviceRequestDeviceInfoStackInfoValidLicenseLevels definition
type UpdateDeviceRequestDeviceInfoStackInfoValidLicenseLevels []string

// UpdateDeviceRequestDeviceInfoUserMicNumbers is the updateDeviceRequestDeviceInfoUserMicNumbers definition
type UpdateDeviceRequestDeviceInfoUserMicNumbers []string

// UpdateDeviceRequestDeviceInfoUserSudiSerialNos is the updateDeviceRequestDeviceInfoUserSudiSerialNos definition
type UpdateDeviceRequestDeviceInfoUserSudiSerialNos []string

// UpdateDeviceRequestRunSummaryList is the updateDeviceRequestRunSummaryList definition
type UpdateDeviceRequestRunSummaryList struct {
	Details         string                                           `json:"details,omitempty"`         //
	ErrorFlag       bool                                             `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo UpdateDeviceRequestRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                          `json:"timestamp,omitempty"`       //
}

// UpdateDeviceRequestRunSummaryListHistoryTaskInfo is the updateDeviceRequestRunSummaryListHistoryTaskInfo definition
type UpdateDeviceRequestRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []UpdateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                         `json:"name,omitempty"`         //
	TimeTaken    float64                                                        `json:"timeTaken,omitempty"`    //
	Type         string                                                         `json:"type,omitempty"`         //
	WorkItemList []UpdateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// UpdateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails is the updateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails definition
type UpdateDeviceRequestRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// UpdateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList is the updateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList definition
type UpdateDeviceRequestRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdateDeviceRequestSystemResetWorkflow is the updateDeviceRequestSystemResetWorkflow definition
type UpdateDeviceRequestSystemResetWorkflow struct {
	TypeID         string                                        `json:"_id,omitempty"`            //
	AddToInventory bool                                          `json:"addToInventory,omitempty"` //
	AddedOn        float64                                       `json:"addedOn,omitempty"`        //
	ConfigID       string                                        `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                       `json:"currTaskIdx,omitempty"`    //
	Description    string                                        `json:"description,omitempty"`    //
	EndTime        int                                           `json:"endTime,omitempty"`        //
	ExecTime       float64                                       `json:"execTime,omitempty"`       //
	ImageID        string                                        `json:"imageId,omitempty"`        //
	InstanceType   string                                        `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                       `json:"lastupdateOn,omitempty"`   //
	Name           string                                        `json:"name,omitempty"`           //
	StartTime      int                                           `json:"startTime,omitempty"`      //
	State          string                                        `json:"state,omitempty"`          //
	Tasks          []UpdateDeviceRequestSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                        `json:"tenantId,omitempty"`       //
	Type           string                                        `json:"type,omitempty"`           //
	UseState       string                                        `json:"useState,omitempty"`       //
	Version        float64                                       `json:"version,omitempty"`        //
}

// UpdateDeviceRequestSystemResetWorkflowTasks is the updateDeviceRequestSystemResetWorkflowTasks definition
type UpdateDeviceRequestSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                   `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                       `json:"endTime,omitempty"`         //
	Name            string                                                    `json:"name,omitempty"`            //
	StartTime       int                                                       `json:"startTime,omitempty"`       //
	State           string                                                    `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                   `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                   `json:"timeTaken,omitempty"`       //
	Type            string                                                    `json:"type,omitempty"`            //
	WorkItemList    []UpdateDeviceRequestSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateDeviceRequestSystemResetWorkflowTasksWorkItemList is the updateDeviceRequestSystemResetWorkflowTasksWorkItemList definition
type UpdateDeviceRequestSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdateDeviceRequestSystemWorkflow is the updateDeviceRequestSystemWorkflow definition
type UpdateDeviceRequestSystemWorkflow struct {
	TypeID         string                                   `json:"_id,omitempty"`            //
	AddToInventory bool                                     `json:"addToInventory,omitempty"` //
	AddedOn        float64                                  `json:"addedOn,omitempty"`        //
	ConfigID       string                                   `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                  `json:"currTaskIdx,omitempty"`    //
	Description    string                                   `json:"description,omitempty"`    //
	EndTime        int                                      `json:"endTime,omitempty"`        //
	ExecTime       float64                                  `json:"execTime,omitempty"`       //
	ImageID        string                                   `json:"imageId,omitempty"`        //
	InstanceType   string                                   `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                  `json:"lastupdateOn,omitempty"`   //
	Name           string                                   `json:"name,omitempty"`           //
	StartTime      int                                      `json:"startTime,omitempty"`      //
	State          string                                   `json:"state,omitempty"`          //
	Tasks          []UpdateDeviceRequestSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                   `json:"tenantId,omitempty"`       //
	Type           string                                   `json:"type,omitempty"`           //
	UseState       string                                   `json:"useState,omitempty"`       //
	Version        float64                                  `json:"version,omitempty"`        //
}

// UpdateDeviceRequestSystemWorkflowTasks is the updateDeviceRequestSystemWorkflowTasks definition
type UpdateDeviceRequestSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                              `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                  `json:"endTime,omitempty"`         //
	Name            string                                               `json:"name,omitempty"`            //
	StartTime       int                                                  `json:"startTime,omitempty"`       //
	State           string                                               `json:"state,omitempty"`           //
	TaskSeqNo       float64                                              `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                              `json:"timeTaken,omitempty"`       //
	Type            string                                               `json:"type,omitempty"`            //
	WorkItemList    []UpdateDeviceRequestSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateDeviceRequestSystemWorkflowTasksWorkItemList is the updateDeviceRequestSystemWorkflowTasksWorkItemList definition
type UpdateDeviceRequestSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdateDeviceRequestWorkflow is the updateDeviceRequestWorkflow definition
type UpdateDeviceRequestWorkflow struct {
	TypeID         string                             `json:"_id,omitempty"`            //
	AddToInventory bool                               `json:"addToInventory,omitempty"` //
	AddedOn        float64                            `json:"addedOn,omitempty"`        //
	ConfigID       string                             `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                            `json:"currTaskIdx,omitempty"`    //
	Description    string                             `json:"description,omitempty"`    //
	EndTime        int                                `json:"endTime,omitempty"`        //
	ExecTime       float64                            `json:"execTime,omitempty"`       //
	ImageID        string                             `json:"imageId,omitempty"`        //
	InstanceType   string                             `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                            `json:"lastupdateOn,omitempty"`   //
	Name           string                             `json:"name,omitempty"`           //
	StartTime      int                                `json:"startTime,omitempty"`      //
	State          string                             `json:"state,omitempty"`          //
	Tasks          []UpdateDeviceRequestWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                             `json:"tenantId,omitempty"`       //
	Type           string                             `json:"type,omitempty"`           //
	UseState       string                             `json:"useState,omitempty"`       //
	Version        float64                            `json:"version,omitempty"`        //
}

// UpdateDeviceRequestWorkflowParameters is the updateDeviceRequestWorkflowParameters definition
type UpdateDeviceRequestWorkflowParameters struct {
	ConfigList             []UpdateDeviceRequestWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                            `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                            `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                            `json:"topOfStackSerialNumber,omitempty"` //
}

// UpdateDeviceRequestWorkflowParametersConfigList is the updateDeviceRequestWorkflowParametersConfigList definition
type UpdateDeviceRequestWorkflowParametersConfigList struct {
	ConfigID         string                                                            `json:"configId,omitempty"`         //
	ConfigParameters []UpdateDeviceRequestWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// UpdateDeviceRequestWorkflowParametersConfigListConfigParameters is the updateDeviceRequestWorkflowParametersConfigListConfigParameters definition
type UpdateDeviceRequestWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// UpdateDeviceRequestWorkflowTasks is the updateDeviceRequestWorkflowTasks definition
type UpdateDeviceRequestWorkflowTasks struct {
	CurrWorkItemIDx float64                                        `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                            `json:"endTime,omitempty"`         //
	Name            string                                         `json:"name,omitempty"`            //
	StartTime       int                                            `json:"startTime,omitempty"`       //
	State           string                                         `json:"state,omitempty"`           //
	TaskSeqNo       float64                                        `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                        `json:"timeTaken,omitempty"`       //
	Type            string                                         `json:"type,omitempty"`            //
	WorkItemList    []UpdateDeviceRequestWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateDeviceRequestWorkflowTasksWorkItemList is the updateDeviceRequestWorkflowTasksWorkItemList definition
type UpdateDeviceRequestWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdatePnPGlobalSettingsRequest is the updatePnPGlobalSettingsRequest definition
type UpdatePnPGlobalSettingsRequest struct {
	TypeID          string                                          `json:"_id,omitempty"`             //
	AAACredentials  UpdatePnPGlobalSettingsRequestAAACredentials    `json:"aaaCredentials,omitempty"`  //
	AcceptEula      bool                                            `json:"acceptEula,omitempty"`      //
	DefaultProfile  UpdatePnPGlobalSettingsRequestDefaultProfile    `json:"defaultProfile,omitempty"`  //
	SavaMappingList []UpdatePnPGlobalSettingsRequestSavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    UpdatePnPGlobalSettingsRequestTaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                          `json:"tenantId,omitempty"`        //
	Version         int                                             `json:"version,omitempty"`         //
}

// UpdatePnPGlobalSettingsRequestAAACredentials is the updatePnPGlobalSettingsRequestAAACredentials definition
type UpdatePnPGlobalSettingsRequestAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// UpdatePnPGlobalSettingsRequestDefaultProfile is the updatePnPGlobalSettingsRequestDefaultProfile definition
type UpdatePnPGlobalSettingsRequestDefaultProfile struct {
	Cert          string   `json:"cert,omitempty"`          //
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` //
	IPAddresses   []string `json:"ipAddresses,omitempty"`   //
	Port          int      `json:"port,omitempty"`          //
	Proxy         bool     `json:"proxy,omitempty"`         //
}

// UpdatePnPGlobalSettingsRequestDefaultProfileFqdnAddresses is the updatePnPGlobalSettingsRequestDefaultProfileFqdnAddresses definition
type UpdatePnPGlobalSettingsRequestDefaultProfileFqdnAddresses []string

// UpdatePnPGlobalSettingsRequestDefaultProfileIPAddresses is the updatePnPGlobalSettingsRequestDefaultProfileIPAddresses definition
type UpdatePnPGlobalSettingsRequestDefaultProfileIPAddresses []string

// UpdatePnPGlobalSettingsRequestSavaMappingList is the updatePnPGlobalSettingsRequestSavaMappingList definition
type UpdatePnPGlobalSettingsRequestSavaMappingList struct {
	AutoSyncPeriod   int                                                     `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                  `json:"ccoUser,omitempty"`          //
	Expiry           int                                                     `json:"expiry,omitempty"`           //
	LastSync         int                                                     `json:"lastSync,omitempty"`         //
	Profile          UpdatePnPGlobalSettingsRequestSavaMappingListProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                  `json:"smartAccountId,omitempty"`   //
	SyncResult       UpdatePnPGlobalSettingsRequestSavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                  `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int                                                     `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                  `json:"syncStatus,omitempty"`       //
	TenantID         string                                                  `json:"tenantId,omitempty"`         //
	Token            string                                                  `json:"token,omitempty"`            //
	VirtualAccountID string                                                  `json:"virtualAccountId,omitempty"` //
}

// UpdatePnPGlobalSettingsRequestSavaMappingListProfile is the updatePnPGlobalSettingsRequestSavaMappingListProfile definition
type UpdatePnPGlobalSettingsRequestSavaMappingListProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault bool   `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        int    `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       bool   `json:"proxy,omitempty"`       //
}

// UpdatePnPGlobalSettingsRequestSavaMappingListSyncResult is the updatePnPGlobalSettingsRequestSavaMappingListSyncResult definition
type UpdatePnPGlobalSettingsRequestSavaMappingListSyncResult struct {
	SyncList []UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                            `json:"syncMsg,omitempty"`  //
}

// UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList is the updatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList definition
type UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncListDeviceSnList is the updatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncListDeviceSnList definition
type UpdatePnPGlobalSettingsRequestSavaMappingListSyncResultSyncListDeviceSnList []string

// UpdatePnPGlobalSettingsRequestTaskTimeOuts is the updatePnPGlobalSettingsRequestTaskTimeOuts definition
type UpdatePnPGlobalSettingsRequestTaskTimeOuts struct {
	ConfigTimeOut        int `json:"configTimeOut,omitempty"`        //
	GeneralTimeOut       int `json:"generalTimeOut,omitempty"`       //
	ImageDownloadTimeOut int `json:"imageDownloadTimeOut,omitempty"` //
}

// UpdatePnPServerProfileRequest is the updatePnPServerProfileRequest definition
type UpdatePnPServerProfileRequest struct {
	AutoSyncPeriod   int                                     `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                  `json:"ccoUser,omitempty"`          //
	Expiry           int                                     `json:"expiry,omitempty"`           //
	LastSync         int                                     `json:"lastSync,omitempty"`         //
	Profile          UpdatePnPServerProfileRequestProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                  `json:"smartAccountId,omitempty"`   //
	SyncResult       UpdatePnPServerProfileRequestSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                  `json:"syncResultStr,omitempty"`    //
	SyncStartTime    int                                     `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                  `json:"syncStatus,omitempty"`       //
	TenantID         string                                  `json:"tenantId,omitempty"`         //
	Token            string                                  `json:"token,omitempty"`            //
	VirtualAccountID string                                  `json:"virtualAccountId,omitempty"` //
}

// UpdatePnPServerProfileRequestProfile is the updatePnPServerProfileRequestProfile definition
type UpdatePnPServerProfileRequestProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault bool   `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        int    `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       bool   `json:"proxy,omitempty"`       //
}

// UpdatePnPServerProfileRequestSyncResult is the updatePnPServerProfileRequestSyncResult definition
type UpdatePnPServerProfileRequestSyncResult struct {
	SyncList []UpdatePnPServerProfileRequestSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                            `json:"syncMsg,omitempty"`  //
}

// UpdatePnPServerProfileRequestSyncResultSyncList is the updatePnPServerProfileRequestSyncResultSyncList definition
type UpdatePnPServerProfileRequestSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// UpdatePnPServerProfileRequestSyncResultSyncListDeviceSnList is the updatePnPServerProfileRequestSyncResultSyncListDeviceSnList definition
type UpdatePnPServerProfileRequestSyncResultSyncListDeviceSnList []string

// UpdateWorkflowRequest is the updateWorkflowRequest definition
type UpdateWorkflowRequest struct {
	TypeID         string                       `json:"_id,omitempty"`            //
	AddToInventory bool                         `json:"addToInventory,omitempty"` //
	AddedOn        int                          `json:"addedOn,omitempty"`        //
	ConfigID       string                       `json:"configId,omitempty"`       //
	CurrTaskIDx    int                          `json:"currTaskIdx,omitempty"`    //
	Description    string                       `json:"description,omitempty"`    //
	EndTime        int                          `json:"endTime,omitempty"`        //
	ExecTime       int                          `json:"execTime,omitempty"`       //
	ImageID        string                       `json:"imageId,omitempty"`        //
	InstanceType   string                       `json:"instanceType,omitempty"`   //
	LastupdateOn   int                          `json:"lastupdateOn,omitempty"`   //
	Name           string                       `json:"name,omitempty"`           //
	StartTime      int                          `json:"startTime,omitempty"`      //
	State          string                       `json:"state,omitempty"`          //
	Tasks          []UpdateWorkflowRequestTasks `json:"tasks,omitempty"`          //
	TenantID       string                       `json:"tenantId,omitempty"`       //
	Type           string                       `json:"type,omitempty"`           //
	UseState       string                       `json:"useState,omitempty"`       //
	Version        int                          `json:"version,omitempty"`        //
}

// UpdateWorkflowRequestTasks is the updateWorkflowRequestTasks definition
type UpdateWorkflowRequestTasks struct {
	CurrWorkItemIDx int                                      `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                      `json:"endTime,omitempty"`         //
	Name            string                                   `json:"name,omitempty"`            //
	StartTime       int                                      `json:"startTime,omitempty"`       //
	State           string                                   `json:"state,omitempty"`           //
	TaskSeqNo       int                                      `json:"taskSeqNo,omitempty"`       //
	TimeTaken       int                                      `json:"timeTaken,omitempty"`       //
	Type            string                                   `json:"type,omitempty"`            //
	WorkItemList    []UpdateWorkflowRequestTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateWorkflowRequestTasksWorkItemList is the updateWorkflowRequestTasksWorkItemList definition
type UpdateWorkflowRequestTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   int    `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime int    `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken int    `json:"timeTaken,omitempty"` //
}

// AddAWorkflowResponse is the addAWorkflowResponse definition
type AddAWorkflowResponse struct {
	TypeID         string                      `json:"_id,omitempty"`            //
	AddToInventory bool                        `json:"addToInventory,omitempty"` //
	AddedOn        float64                     `json:"addedOn,omitempty"`        //
	ConfigID       string                      `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                     `json:"currTaskIdx,omitempty"`    //
	Description    string                      `json:"description,omitempty"`    //
	EndTime        int                         `json:"endTime,omitempty"`        //
	ExecTime       float64                     `json:"execTime,omitempty"`       //
	ImageID        string                      `json:"imageId,omitempty"`        //
	InstanceType   string                      `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                     `json:"lastupdateOn,omitempty"`   //
	Name           string                      `json:"name,omitempty"`           //
	StartTime      int                         `json:"startTime,omitempty"`      //
	State          string                      `json:"state,omitempty"`          //
	Tasks          []AddAWorkflowResponseTasks `json:"tasks,omitempty"`          //
	TenantID       string                      `json:"tenantId,omitempty"`       //
	Type           string                      `json:"type,omitempty"`           //
	UseState       string                      `json:"useState,omitempty"`       //
	Version        float64                     `json:"version,omitempty"`        //
}

// AddAWorkflowResponseTasks is the addAWorkflowResponseTasks definition
type AddAWorkflowResponseTasks struct {
	CurrWorkItemIDx float64                                 `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                     `json:"endTime,omitempty"`         //
	Name            string                                  `json:"name,omitempty"`            //
	StartTime       int                                     `json:"startTime,omitempty"`       //
	State           string                                  `json:"state,omitempty"`           //
	TaskSeqNo       float64                                 `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                 `json:"timeTaken,omitempty"`       //
	Type            string                                  `json:"type,omitempty"`            //
	WorkItemList    []AddAWorkflowResponseTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddAWorkflowResponseTasksWorkItemList is the addAWorkflowResponseTasksWorkItemList definition
type AddAWorkflowResponseTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseResponse is the addDeviceToPnpDatabaseResponse definition
type AddDeviceToPnpDatabaseResponse struct {
	TypeID               string                                            `json:"_id,omitempty"`                  //
	DayZeroConfig        AddDeviceToPnpDatabaseResponseDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                            `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           AddDeviceToPnpDatabaseResponseDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []AddDeviceToPnpDatabaseResponseRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  AddDeviceToPnpDatabaseResponseSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       AddDeviceToPnpDatabaseResponseSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                            `json:"tenantId,omitempty"`             //
	Version              float64                                           `json:"version,omitempty"`              //
	Workflow             AddDeviceToPnpDatabaseResponseWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   AddDeviceToPnpDatabaseResponseWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// AddDeviceToPnpDatabaseResponseDayZeroConfig is the addDeviceToPnpDatabaseResponseDayZeroConfig definition
type AddDeviceToPnpDatabaseResponseDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfo is the addDeviceToPnpDatabaseResponseDeviceInfo definition
type AddDeviceToPnpDatabaseResponseDeviceInfo struct {
	AAACredentials            AddDeviceToPnpDatabaseResponseDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                                        `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                       `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                         `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                         `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                         `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                         `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                       `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                         `json:"cmState,omitempty"`                   //
	Description               string                                                         `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                       `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                         `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                       `json:"featuresSupported,omitempty"`         //
	FileSystemList            []AddDeviceToPnpDatabaseResponseDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                                        `json:"firstContact,omitempty"`              //
	Hostname                  string                                                         `json:"hostname,omitempty"`                  //
	HTTPHeaders               []AddDeviceToPnpDatabaseResponseDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                         `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                         `json:"imageVersion,omitempty"`              //
	IPInterfaces              []AddDeviceToPnpDatabaseResponseDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                                        `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                                        `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                                        `json:"lastUpdateOn,omitempty"`              //
	Location                  AddDeviceToPnpDatabaseResponseDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                         `json:"macAddress,omitempty"`                //
	Mode                      string                                                         `json:"mode,omitempty"`                      //
	Name                      string                                                         `json:"name,omitempty"`                      //
	NeighborLinks             []AddDeviceToPnpDatabaseResponseDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                         `json:"onbState,omitempty"`                  //
	Pid                       string                                                         `json:"pid,omitempty"`                       //
	PnpProfileList            []AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                           `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []AddDeviceToPnpDatabaseResponseDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                         `json:"projectId,omitempty"`                 //
	ProjectName               string                                                         `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                           `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                         `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                         `json:"siteId,omitempty"`                    //
	SiteName                  string                                                         `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                         `json:"smartAccountId,omitempty"`            //
	Source                    string                                                         `json:"source,omitempty"`                    //
	Stack                     bool                                                           `json:"stack,omitempty"`                     //
	StackInfo                 AddDeviceToPnpDatabaseResponseDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                         `json:"state,omitempty"`                     //
	SudiRequired              bool                                                           `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                         `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                                       `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                                       `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                         `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                         `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                         `json:"workflowName,omitempty"`              //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoAAACredentials is the addDeviceToPnpDatabaseResponseDeviceInfoAAACredentials definition
type AddDeviceToPnpDatabaseResponseDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoAddnMacAddrs is the addDeviceToPnpDatabaseResponseDeviceInfoAddnMacAddrs definition
type AddDeviceToPnpDatabaseResponseDeviceInfoAddnMacAddrs []string

// AddDeviceToPnpDatabaseResponseDeviceInfoCapabilitiesSupported is the addDeviceToPnpDatabaseResponseDeviceInfoCapabilitiesSupported definition
type AddDeviceToPnpDatabaseResponseDeviceInfoCapabilitiesSupported []string

// AddDeviceToPnpDatabaseResponseDeviceInfoDeviceSudiSerialNos is the addDeviceToPnpDatabaseResponseDeviceInfoDeviceSudiSerialNos definition
type AddDeviceToPnpDatabaseResponseDeviceInfoDeviceSudiSerialNos []string

// AddDeviceToPnpDatabaseResponseDeviceInfoFeaturesSupported is the addDeviceToPnpDatabaseResponseDeviceInfoFeaturesSupported definition
type AddDeviceToPnpDatabaseResponseDeviceInfoFeaturesSupported []string

// AddDeviceToPnpDatabaseResponseDeviceInfoFileSystemList is the addDeviceToPnpDatabaseResponseDeviceInfoFileSystemList definition
type AddDeviceToPnpDatabaseResponseDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoHTTPHeaders is the addDeviceToPnpDatabaseResponseDeviceInfoHTTPHeaders definition
type AddDeviceToPnpDatabaseResponseDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoIPInterfaces is the addDeviceToPnpDatabaseResponseDeviceInfoIPInterfaces definition
type AddDeviceToPnpDatabaseResponseDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoIPInterfacesIPv6AddressList is the addDeviceToPnpDatabaseResponseDeviceInfoIPInterfacesIPv6AddressList definition
type AddDeviceToPnpDatabaseResponseDeviceInfoIPInterfacesIPv6AddressList []string

// AddDeviceToPnpDatabaseResponseDeviceInfoLocation is the addDeviceToPnpDatabaseResponseDeviceInfoLocation definition
type AddDeviceToPnpDatabaseResponseDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoNeighborLinks is the addDeviceToPnpDatabaseResponseDeviceInfoNeighborLinks definition
type AddDeviceToPnpDatabaseResponseDeviceInfoNeighborLinks struct {
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

// AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileList is the addDeviceToPnpDatabaseResponseDeviceInfoPnpProfileList definition
type AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                  `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                                    `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                  `json:"profileName,omitempty"`       //
	SecondaryEndpoint AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListPrimaryEndpoint is the addDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListPrimaryEndpoint definition
type AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListSecondaryEndpoint is the addDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListSecondaryEndpoint definition
type AddDeviceToPnpDatabaseResponseDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoPreWorkflowCliOuputs is the addDeviceToPnpDatabaseResponseDeviceInfoPreWorkflowCliOuputs definition
type AddDeviceToPnpDatabaseResponseDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoStackInfo is the addDeviceToPnpDatabaseResponseDeviceInfoStackInfo definition
type AddDeviceToPnpDatabaseResponseDeviceInfoStackInfo struct {
	IsFullRing             bool                                                               `json:"isFullRing,omitempty"`             //
	StackMemberList        []AddDeviceToPnpDatabaseResponseDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                             `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                               `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                            `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                           `json:"validLicenseLevels,omitempty"`     //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoStackInfoStackMemberList is the addDeviceToPnpDatabaseResponseDeviceInfoStackInfoStackMemberList definition
type AddDeviceToPnpDatabaseResponseDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseDeviceInfoStackInfoValidLicenseLevels is the addDeviceToPnpDatabaseResponseDeviceInfoStackInfoValidLicenseLevels definition
type AddDeviceToPnpDatabaseResponseDeviceInfoStackInfoValidLicenseLevels []string

// AddDeviceToPnpDatabaseResponseDeviceInfoUserMicNumbers is the addDeviceToPnpDatabaseResponseDeviceInfoUserMicNumbers definition
type AddDeviceToPnpDatabaseResponseDeviceInfoUserMicNumbers []string

// AddDeviceToPnpDatabaseResponseDeviceInfoUserSudiSerialNos is the addDeviceToPnpDatabaseResponseDeviceInfoUserSudiSerialNos definition
type AddDeviceToPnpDatabaseResponseDeviceInfoUserSudiSerialNos []string

// AddDeviceToPnpDatabaseResponseRunSummaryList is the addDeviceToPnpDatabaseResponseRunSummaryList definition
type AddDeviceToPnpDatabaseResponseRunSummaryList struct {
	Details         string                                                      `json:"details,omitempty"`         //
	ErrorFlag       bool                                                        `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                                     `json:"timestamp,omitempty"`       //
}

// AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfo is the addDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfo definition
type AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                    `json:"name,omitempty"`         //
	TimeTaken    float64                                                                   `json:"timeTaken,omitempty"`    //
	Type         string                                                                    `json:"type,omitempty"`         //
	WorkItemList []AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoAddnDetails is the addDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoAddnDetails definition
type AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoWorkItemList is the addDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoWorkItemList definition
type AddDeviceToPnpDatabaseResponseRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseSystemResetWorkflow is the addDeviceToPnpDatabaseResponseSystemResetWorkflow definition
type AddDeviceToPnpDatabaseResponseSystemResetWorkflow struct {
	TypeID         string                                                   `json:"_id,omitempty"`            //
	AddToInventory bool                                                     `json:"addToInventory,omitempty"` //
	AddedOn        float64                                                  `json:"addedOn,omitempty"`        //
	ConfigID       string                                                   `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                                  `json:"currTaskIdx,omitempty"`    //
	Description    string                                                   `json:"description,omitempty"`    //
	EndTime        int                                                      `json:"endTime,omitempty"`        //
	ExecTime       float64                                                  `json:"execTime,omitempty"`       //
	ImageID        string                                                   `json:"imageId,omitempty"`        //
	InstanceType   string                                                   `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                                  `json:"lastupdateOn,omitempty"`   //
	Name           string                                                   `json:"name,omitempty"`           //
	StartTime      int                                                      `json:"startTime,omitempty"`      //
	State          string                                                   `json:"state,omitempty"`          //
	Tasks          []AddDeviceToPnpDatabaseResponseSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                   `json:"tenantId,omitempty"`       //
	Type           string                                                   `json:"type,omitempty"`           //
	UseState       string                                                   `json:"useState,omitempty"`       //
	Version        float64                                                  `json:"version,omitempty"`        //
}

// AddDeviceToPnpDatabaseResponseSystemResetWorkflowTasks is the addDeviceToPnpDatabaseResponseSystemResetWorkflowTasks definition
type AddDeviceToPnpDatabaseResponseSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                              `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                                  `json:"endTime,omitempty"`         //
	Name            string                                                               `json:"name,omitempty"`            //
	StartTime       int                                                                  `json:"startTime,omitempty"`       //
	State           string                                                               `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                              `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                              `json:"timeTaken,omitempty"`       //
	Type            string                                                               `json:"type,omitempty"`            //
	WorkItemList    []AddDeviceToPnpDatabaseResponseSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddDeviceToPnpDatabaseResponseSystemResetWorkflowTasksWorkItemList is the addDeviceToPnpDatabaseResponseSystemResetWorkflowTasksWorkItemList definition
type AddDeviceToPnpDatabaseResponseSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseSystemWorkflow is the addDeviceToPnpDatabaseResponseSystemWorkflow definition
type AddDeviceToPnpDatabaseResponseSystemWorkflow struct {
	TypeID         string                                              `json:"_id,omitempty"`            //
	AddToInventory bool                                                `json:"addToInventory,omitempty"` //
	AddedOn        float64                                             `json:"addedOn,omitempty"`        //
	ConfigID       string                                              `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                             `json:"currTaskIdx,omitempty"`    //
	Description    string                                              `json:"description,omitempty"`    //
	EndTime        int                                                 `json:"endTime,omitempty"`        //
	ExecTime       float64                                             `json:"execTime,omitempty"`       //
	ImageID        string                                              `json:"imageId,omitempty"`        //
	InstanceType   string                                              `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                             `json:"lastupdateOn,omitempty"`   //
	Name           string                                              `json:"name,omitempty"`           //
	StartTime      int                                                 `json:"startTime,omitempty"`      //
	State          string                                              `json:"state,omitempty"`          //
	Tasks          []AddDeviceToPnpDatabaseResponseSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                              `json:"tenantId,omitempty"`       //
	Type           string                                              `json:"type,omitempty"`           //
	UseState       string                                              `json:"useState,omitempty"`       //
	Version        float64                                             `json:"version,omitempty"`        //
}

// AddDeviceToPnpDatabaseResponseSystemWorkflowTasks is the addDeviceToPnpDatabaseResponseSystemWorkflowTasks definition
type AddDeviceToPnpDatabaseResponseSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                         `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                             `json:"endTime,omitempty"`         //
	Name            string                                                          `json:"name,omitempty"`            //
	StartTime       int                                                             `json:"startTime,omitempty"`       //
	State           string                                                          `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                         `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                         `json:"timeTaken,omitempty"`       //
	Type            string                                                          `json:"type,omitempty"`            //
	WorkItemList    []AddDeviceToPnpDatabaseResponseSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddDeviceToPnpDatabaseResponseSystemWorkflowTasksWorkItemList is the addDeviceToPnpDatabaseResponseSystemWorkflowTasksWorkItemList definition
type AddDeviceToPnpDatabaseResponseSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseWorkflow is the addDeviceToPnpDatabaseResponseWorkflow definition
type AddDeviceToPnpDatabaseResponseWorkflow struct {
	TypeID         string                                        `json:"_id,omitempty"`            //
	AddToInventory bool                                          `json:"addToInventory,omitempty"` //
	AddedOn        float64                                       `json:"addedOn,omitempty"`        //
	ConfigID       string                                        `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                       `json:"currTaskIdx,omitempty"`    //
	Description    string                                        `json:"description,omitempty"`    //
	EndTime        int                                           `json:"endTime,omitempty"`        //
	ExecTime       float64                                       `json:"execTime,omitempty"`       //
	ImageID        string                                        `json:"imageId,omitempty"`        //
	InstanceType   string                                        `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                       `json:"lastupdateOn,omitempty"`   //
	Name           string                                        `json:"name,omitempty"`           //
	StartTime      int                                           `json:"startTime,omitempty"`      //
	State          string                                        `json:"state,omitempty"`          //
	Tasks          []AddDeviceToPnpDatabaseResponseWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                        `json:"tenantId,omitempty"`       //
	Type           string                                        `json:"type,omitempty"`           //
	UseState       string                                        `json:"useState,omitempty"`       //
	Version        float64                                       `json:"version,omitempty"`        //
}

// AddDeviceToPnpDatabaseResponseWorkflowParameters is the addDeviceToPnpDatabaseResponseWorkflowParameters definition
type AddDeviceToPnpDatabaseResponseWorkflowParameters struct {
	ConfigList             []AddDeviceToPnpDatabaseResponseWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                       `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                       `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                       `json:"topOfStackSerialNumber,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseWorkflowParametersConfigList is the addDeviceToPnpDatabaseResponseWorkflowParametersConfigList definition
type AddDeviceToPnpDatabaseResponseWorkflowParametersConfigList struct {
	ConfigID         string                                                                       `json:"configId,omitempty"`         //
	ConfigParameters []AddDeviceToPnpDatabaseResponseWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseWorkflowParametersConfigListConfigParameters is the addDeviceToPnpDatabaseResponseWorkflowParametersConfigListConfigParameters definition
type AddDeviceToPnpDatabaseResponseWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// AddDeviceToPnpDatabaseResponseWorkflowTasks is the addDeviceToPnpDatabaseResponseWorkflowTasks definition
type AddDeviceToPnpDatabaseResponseWorkflowTasks struct {
	CurrWorkItemIDx float64                                                   `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                       `json:"endTime,omitempty"`         //
	Name            string                                                    `json:"name,omitempty"`            //
	StartTime       int                                                       `json:"startTime,omitempty"`       //
	State           string                                                    `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                   `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                   `json:"timeTaken,omitempty"`       //
	Type            string                                                    `json:"type,omitempty"`            //
	WorkItemList    []AddDeviceToPnpDatabaseResponseWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// AddDeviceToPnpDatabaseResponseWorkflowTasksWorkItemList is the addDeviceToPnpDatabaseResponseWorkflowTasksWorkItemList definition
type AddDeviceToPnpDatabaseResponseWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddVirtualAccountResponse is the addVirtualAccountResponse definition
type AddVirtualAccountResponse struct {
	AutoSyncPeriod   float64                             `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                              `json:"ccoUser,omitempty"`          //
	Expiry           float64                             `json:"expiry,omitempty"`           //
	LastSync         float64                             `json:"lastSync,omitempty"`         //
	Profile          AddVirtualAccountResponseProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                              `json:"smartAccountId,omitempty"`   //
	SyncResult       AddVirtualAccountResponseSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                              `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                             `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                              `json:"syncStatus,omitempty"`       //
	TenantID         string                              `json:"tenantId,omitempty"`         //
	Token            string                              `json:"token,omitempty"`            //
	VirtualAccountID string                              `json:"virtualAccountId,omitempty"` //
}

// AddVirtualAccountResponseProfile is the addVirtualAccountResponseProfile definition
type AddVirtualAccountResponseProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// AddVirtualAccountResponseSyncResult is the addVirtualAccountResponseSyncResult definition
type AddVirtualAccountResponseSyncResult struct {
	SyncList []AddVirtualAccountResponseSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                        `json:"syncMsg,omitempty"`  //
}

// AddVirtualAccountResponseSyncResultSyncList is the addVirtualAccountResponseSyncResultSyncList definition
type AddVirtualAccountResponseSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// AddVirtualAccountResponseSyncResultSyncListDeviceSnList is the addVirtualAccountResponseSyncResultSyncListDeviceSnList definition
type AddVirtualAccountResponseSyncResultSyncListDeviceSnList []string

// ClaimADeviceToASiteResponse is the claimADeviceToASiteResponse definition
type ClaimADeviceToASiteResponse struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// ClaimDeviceResponse is the claimDeviceResponse definition
type ClaimDeviceResponse struct {
	JSONArrayResponse []string `json:"jsonArrayResponse,omitempty"` //
	JSONResponse      string   `json:"jsonResponse,omitempty"`      //
	Message           string   `json:"message,omitempty"`           //
	StatusCode        float64  `json:"statusCode,omitempty"`        //
}

// ClaimDeviceResponseJSONArrayResponse is the claimDeviceResponseJSONArrayResponse definition
type ClaimDeviceResponseJSONArrayResponse []string

// DeleteDeviceByIDFromPnPResponseDayZeroConfig is the deleteDeviceByIDFromPnPResponseDayZeroConfig definition
type DeleteDeviceByIDFromPnPResponseDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfo is the deleteDeviceByIDFromPnPResponseDeviceInfo definition
type DeleteDeviceByIDFromPnPResponseDeviceInfo struct {
	AAACredentials            DeleteDeviceByIDFromPnPResponseDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                                         `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                        `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                          `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                          `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                          `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                          `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                        `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                          `json:"cmState,omitempty"`                   //
	Description               string                                                          `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                        `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                          `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                        `json:"featuresSupported,omitempty"`         //
	FileSystemList            []DeleteDeviceByIDFromPnPResponseDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                                         `json:"firstContact,omitempty"`              //
	Hostname                  string                                                          `json:"hostname,omitempty"`                  //
	HTTPHeaders               []DeleteDeviceByIDFromPnPResponseDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                          `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                          `json:"imageVersion,omitempty"`              //
	IPInterfaces              []DeleteDeviceByIDFromPnPResponseDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                                         `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                                         `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                                         `json:"lastUpdateOn,omitempty"`              //
	Location                  DeleteDeviceByIDFromPnPResponseDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                          `json:"macAddress,omitempty"`                //
	Mode                      string                                                          `json:"mode,omitempty"`                      //
	Name                      string                                                          `json:"name,omitempty"`                      //
	NeighborLinks             []DeleteDeviceByIDFromPnPResponseDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                          `json:"onbState,omitempty"`                  //
	Pid                       string                                                          `json:"pid,omitempty"`                       //
	PnpProfileList            []DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                            `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []DeleteDeviceByIDFromPnPResponseDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                          `json:"projectId,omitempty"`                 //
	ProjectName               string                                                          `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                            `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                          `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                          `json:"siteId,omitempty"`                    //
	SiteName                  string                                                          `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                          `json:"smartAccountId,omitempty"`            //
	Source                    string                                                          `json:"source,omitempty"`                    //
	Stack                     bool                                                            `json:"stack,omitempty"`                     //
	StackInfo                 DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                          `json:"state,omitempty"`                     //
	SudiRequired              bool                                                            `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                          `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                                        `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                                        `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                          `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                          `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                          `json:"workflowName,omitempty"`              //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoAAACredentials is the deleteDeviceByIDFromPnPResponseDeviceInfoAAACredentials definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoAddnMacAddrs is the deleteDeviceByIDFromPnPResponseDeviceInfoAddnMacAddrs definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoAddnMacAddrs []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoCapabilitiesSupported is the deleteDeviceByIDFromPnPResponseDeviceInfoCapabilitiesSupported definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoCapabilitiesSupported []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoDeviceSudiSerialNos is the deleteDeviceByIDFromPnPResponseDeviceInfoDeviceSudiSerialNos definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoDeviceSudiSerialNos []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoFeaturesSupported is the deleteDeviceByIDFromPnPResponseDeviceInfoFeaturesSupported definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoFeaturesSupported []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoFileSystemList is the deleteDeviceByIDFromPnPResponseDeviceInfoFileSystemList definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoHTTPHeaders is the deleteDeviceByIDFromPnPResponseDeviceInfoHTTPHeaders definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoIPInterfaces is the deleteDeviceByIDFromPnPResponseDeviceInfoIPInterfaces definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoIPInterfacesIPv6AddressList is the deleteDeviceByIDFromPnPResponseDeviceInfoIPInterfacesIPv6AddressList definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoIPInterfacesIPv6AddressList []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoLocation is the deleteDeviceByIDFromPnPResponseDeviceInfoLocation definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoNeighborLinks is the deleteDeviceByIDFromPnPResponseDeviceInfoNeighborLinks definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoNeighborLinks struct {
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

// DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileList is the deleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileList definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                   `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                                     `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                   `json:"profileName,omitempty"`       //
	SecondaryEndpoint DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListPrimaryEndpoint is the deleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListPrimaryEndpoint definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListSecondaryEndpoint is the deleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListSecondaryEndpoint definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoPreWorkflowCliOuputs is the deleteDeviceByIDFromPnPResponseDeviceInfoPreWorkflowCliOuputs definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfo is the deleteDeviceByIDFromPnPResponseDeviceInfoStackInfo definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfo struct {
	IsFullRing             bool                                                                `json:"isFullRing,omitempty"`             //
	StackMemberList        []DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                              `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                                `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                             `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                            `json:"validLicenseLevels,omitempty"`     //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfoStackMemberList is the deleteDeviceByIDFromPnPResponseDeviceInfoStackInfoStackMemberList definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfoValidLicenseLevels is the deleteDeviceByIDFromPnPResponseDeviceInfoStackInfoValidLicenseLevels definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoStackInfoValidLicenseLevels []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoUserMicNumbers is the deleteDeviceByIDFromPnPResponseDeviceInfoUserMicNumbers definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoUserMicNumbers []string

// DeleteDeviceByIDFromPnPResponseDeviceInfoUserSudiSerialNos is the deleteDeviceByIDFromPnPResponseDeviceInfoUserSudiSerialNos definition
type DeleteDeviceByIDFromPnPResponseDeviceInfoUserSudiSerialNos []string

// DeleteDeviceByIDFromPnPResponseRunSummaryList is the deleteDeviceByIDFromPnPResponseRunSummaryList definition
type DeleteDeviceByIDFromPnPResponseRunSummaryList struct {
	Details         string                                                       `json:"details,omitempty"`         //
	ErrorFlag       bool                                                         `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                                      `json:"timestamp,omitempty"`       //
}

// DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfo is the deleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfo definition
type DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                     `json:"name,omitempty"`         //
	TimeTaken    float64                                                                    `json:"timeTaken,omitempty"`    //
	Type         string                                                                     `json:"type,omitempty"`         //
	WorkItemList []DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoAddnDetails is the deleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoAddnDetails definition
type DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoWorkItemList is the deleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoWorkItemList definition
type DeleteDeviceByIDFromPnPResponseRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseSystemResetWorkflow is the deleteDeviceByIDFromPnPResponseSystemResetWorkflow definition
type DeleteDeviceByIDFromPnPResponseSystemResetWorkflow struct {
	TypeID         string                                                    `json:"_id,omitempty"`            //
	AddToInventory bool                                                      `json:"addToInventory,omitempty"` //
	AddedOn        float64                                                   `json:"addedOn,omitempty"`        //
	ConfigID       string                                                    `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                                   `json:"currTaskIdx,omitempty"`    //
	Description    string                                                    `json:"description,omitempty"`    //
	EndTime        int                                                       `json:"endTime,omitempty"`        //
	ExecTime       float64                                                   `json:"execTime,omitempty"`       //
	ImageID        string                                                    `json:"imageId,omitempty"`        //
	InstanceType   string                                                    `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                                   `json:"lastupdateOn,omitempty"`   //
	Name           string                                                    `json:"name,omitempty"`           //
	StartTime      int                                                       `json:"startTime,omitempty"`      //
	State          string                                                    `json:"state,omitempty"`          //
	Tasks          []DeleteDeviceByIDFromPnPResponseSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                    `json:"tenantId,omitempty"`       //
	Type           string                                                    `json:"type,omitempty"`           //
	UseState       string                                                    `json:"useState,omitempty"`       //
	Version        float64                                                   `json:"version,omitempty"`        //
}

// DeleteDeviceByIDFromPnPResponseSystemResetWorkflowTasks is the deleteDeviceByIDFromPnPResponseSystemResetWorkflowTasks definition
type DeleteDeviceByIDFromPnPResponseSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                               `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                                   `json:"endTime,omitempty"`         //
	Name            string                                                                `json:"name,omitempty"`            //
	StartTime       int                                                                   `json:"startTime,omitempty"`       //
	State           string                                                                `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                               `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                               `json:"timeTaken,omitempty"`       //
	Type            string                                                                `json:"type,omitempty"`            //
	WorkItemList    []DeleteDeviceByIDFromPnPResponseSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// DeleteDeviceByIDFromPnPResponseSystemResetWorkflowTasksWorkItemList is the deleteDeviceByIDFromPnPResponseSystemResetWorkflowTasksWorkItemList definition
type DeleteDeviceByIDFromPnPResponseSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseSystemWorkflow is the deleteDeviceByIDFromPnPResponseSystemWorkflow definition
type DeleteDeviceByIDFromPnPResponseSystemWorkflow struct {
	TypeID         string                                               `json:"_id,omitempty"`            //
	AddToInventory bool                                                 `json:"addToInventory,omitempty"` //
	AddedOn        float64                                              `json:"addedOn,omitempty"`        //
	ConfigID       string                                               `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                              `json:"currTaskIdx,omitempty"`    //
	Description    string                                               `json:"description,omitempty"`    //
	EndTime        int                                                  `json:"endTime,omitempty"`        //
	ExecTime       float64                                              `json:"execTime,omitempty"`       //
	ImageID        string                                               `json:"imageId,omitempty"`        //
	InstanceType   string                                               `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                              `json:"lastupdateOn,omitempty"`   //
	Name           string                                               `json:"name,omitempty"`           //
	StartTime      int                                                  `json:"startTime,omitempty"`      //
	State          string                                               `json:"state,omitempty"`          //
	Tasks          []DeleteDeviceByIDFromPnPResponseSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                               `json:"tenantId,omitempty"`       //
	Type           string                                               `json:"type,omitempty"`           //
	UseState       string                                               `json:"useState,omitempty"`       //
	Version        float64                                              `json:"version,omitempty"`        //
}

// DeleteDeviceByIDFromPnPResponseSystemWorkflowTasks is the deleteDeviceByIDFromPnPResponseSystemWorkflowTasks definition
type DeleteDeviceByIDFromPnPResponseSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                          `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                              `json:"endTime,omitempty"`         //
	Name            string                                                           `json:"name,omitempty"`            //
	StartTime       int                                                              `json:"startTime,omitempty"`       //
	State           string                                                           `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                          `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                          `json:"timeTaken,omitempty"`       //
	Type            string                                                           `json:"type,omitempty"`            //
	WorkItemList    []DeleteDeviceByIDFromPnPResponseSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// DeleteDeviceByIDFromPnPResponseSystemWorkflowTasksWorkItemList is the deleteDeviceByIDFromPnPResponseSystemWorkflowTasksWorkItemList definition
type DeleteDeviceByIDFromPnPResponseSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseWorkflow is the deleteDeviceByIDFromPnPResponseWorkflow definition
type DeleteDeviceByIDFromPnPResponseWorkflow struct {
	TypeID         string                                         `json:"_id,omitempty"`            //
	AddToInventory bool                                           `json:"addToInventory,omitempty"` //
	AddedOn        float64                                        `json:"addedOn,omitempty"`        //
	ConfigID       string                                         `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                        `json:"currTaskIdx,omitempty"`    //
	Description    string                                         `json:"description,omitempty"`    //
	EndTime        int                                            `json:"endTime,omitempty"`        //
	ExecTime       float64                                        `json:"execTime,omitempty"`       //
	ImageID        string                                         `json:"imageId,omitempty"`        //
	InstanceType   string                                         `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                        `json:"lastupdateOn,omitempty"`   //
	Name           string                                         `json:"name,omitempty"`           //
	StartTime      int                                            `json:"startTime,omitempty"`      //
	State          string                                         `json:"state,omitempty"`          //
	Tasks          []DeleteDeviceByIDFromPnPResponseWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                         `json:"tenantId,omitempty"`       //
	Type           string                                         `json:"type,omitempty"`           //
	UseState       string                                         `json:"useState,omitempty"`       //
	Version        float64                                        `json:"version,omitempty"`        //
}

// DeleteDeviceByIDFromPnPResponseWorkflowParameters is the deleteDeviceByIDFromPnPResponseWorkflowParameters definition
type DeleteDeviceByIDFromPnPResponseWorkflowParameters struct {
	ConfigList             []DeleteDeviceByIDFromPnPResponseWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                        `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                        `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                        `json:"topOfStackSerialNumber,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseWorkflowParametersConfigList is the deleteDeviceByIDFromPnPResponseWorkflowParametersConfigList definition
type DeleteDeviceByIDFromPnPResponseWorkflowParametersConfigList struct {
	ConfigID         string                                                                        `json:"configId,omitempty"`         //
	ConfigParameters []DeleteDeviceByIDFromPnPResponseWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseWorkflowParametersConfigListConfigParameters is the deleteDeviceByIDFromPnPResponseWorkflowParametersConfigListConfigParameters definition
type DeleteDeviceByIDFromPnPResponseWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponseWorkflowTasks is the deleteDeviceByIDFromPnPResponseWorkflowTasks definition
type DeleteDeviceByIDFromPnPResponseWorkflowTasks struct {
	CurrWorkItemIDx float64                                                    `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                        `json:"endTime,omitempty"`         //
	Name            string                                                     `json:"name,omitempty"`            //
	StartTime       int                                                        `json:"startTime,omitempty"`       //
	State           string                                                     `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                    `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                    `json:"timeTaken,omitempty"`       //
	Type            string                                                     `json:"type,omitempty"`            //
	WorkItemList    []DeleteDeviceByIDFromPnPResponseWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// DeleteDeviceByIDFromPnPResponseWorkflowTasksWorkItemList is the deleteDeviceByIDFromPnPResponseWorkflowTasksWorkItemList definition
type DeleteDeviceByIDFromPnPResponseWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// DeleteDeviceByIDFromPnPResponse is the deleteDeviceByIdFromPnPResponse definition
type DeleteDeviceByIDFromPnPResponse struct {
	TypeID               string                                             `json:"_id,omitempty"`                  //
	DayZeroConfig        DeleteDeviceByIDFromPnPResponseDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                             `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           DeleteDeviceByIDFromPnPResponseDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []DeleteDeviceByIDFromPnPResponseRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  DeleteDeviceByIDFromPnPResponseSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       DeleteDeviceByIDFromPnPResponseSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                             `json:"tenantId,omitempty"`             //
	Version              float64                                            `json:"version,omitempty"`              //
	Workflow             DeleteDeviceByIDFromPnPResponseWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   DeleteDeviceByIDFromPnPResponseWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// DeleteWorkflowByIDResponseTasks is the deleteWorkflowByIDResponseTasks definition
type DeleteWorkflowByIDResponseTasks struct {
	CurrWorkItemIDx float64                                       `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                           `json:"endTime,omitempty"`         //
	Name            string                                        `json:"name,omitempty"`            //
	StartTime       int                                           `json:"startTime,omitempty"`       //
	State           string                                        `json:"state,omitempty"`           //
	TaskSeqNo       float64                                       `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                       `json:"timeTaken,omitempty"`       //
	Type            string                                        `json:"type,omitempty"`            //
	WorkItemList    []DeleteWorkflowByIDResponseTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// DeleteWorkflowByIDResponseTasksWorkItemList is the deleteWorkflowByIDResponseTasksWorkItemList definition
type DeleteWorkflowByIDResponseTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// DeleteWorkflowByIDResponse is the deleteWorkflowByIdResponse definition
type DeleteWorkflowByIDResponse struct {
	TypeID         string                            `json:"_id,omitempty"`            //
	AddToInventory bool                              `json:"addToInventory,omitempty"` //
	AddedOn        float64                           `json:"addedOn,omitempty"`        //
	ConfigID       string                            `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                           `json:"currTaskIdx,omitempty"`    //
	Description    string                            `json:"description,omitempty"`    //
	EndTime        int                               `json:"endTime,omitempty"`        //
	ExecTime       float64                           `json:"execTime,omitempty"`       //
	ImageID        string                            `json:"imageId,omitempty"`        //
	InstanceType   string                            `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                           `json:"lastupdateOn,omitempty"`   //
	Name           string                            `json:"name,omitempty"`           //
	StartTime      int                               `json:"startTime,omitempty"`      //
	State          string                            `json:"state,omitempty"`          //
	Tasks          []DeleteWorkflowByIDResponseTasks `json:"tasks,omitempty"`          //
	TenantID       string                            `json:"tenantId,omitempty"`       //
	Type           string                            `json:"type,omitempty"`           //
	UseState       string                            `json:"useState,omitempty"`       //
	Version        float64                           `json:"version,omitempty"`        //
}

// DeregisterVirtualAccountResponse is the deregisterVirtualAccountResponse definition
type DeregisterVirtualAccountResponse struct {
	AutoSyncPeriod   float64                                    `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                     `json:"ccoUser,omitempty"`          //
	Expiry           float64                                    `json:"expiry,omitempty"`           //
	LastSync         float64                                    `json:"lastSync,omitempty"`         //
	Profile          DeregisterVirtualAccountResponseProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                     `json:"smartAccountId,omitempty"`   //
	SyncResult       DeregisterVirtualAccountResponseSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                     `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                                    `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                     `json:"syncStatus,omitempty"`       //
	TenantID         string                                     `json:"tenantId,omitempty"`         //
	Token            string                                     `json:"token,omitempty"`            //
	VirtualAccountID string                                     `json:"virtualAccountId,omitempty"` //
}

// DeregisterVirtualAccountResponseProfile is the deregisterVirtualAccountResponseProfile definition
type DeregisterVirtualAccountResponseProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// DeregisterVirtualAccountResponseSyncResult is the deregisterVirtualAccountResponseSyncResult definition
type DeregisterVirtualAccountResponseSyncResult struct {
	SyncList []DeregisterVirtualAccountResponseSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                               `json:"syncMsg,omitempty"`  //
}

// DeregisterVirtualAccountResponseSyncResultSyncList is the deregisterVirtualAccountResponseSyncResultSyncList definition
type DeregisterVirtualAccountResponseSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// DeregisterVirtualAccountResponseSyncResultSyncListDeviceSnList is the deregisterVirtualAccountResponseSyncResultSyncListDeviceSnList definition
type DeregisterVirtualAccountResponseSyncResultSyncListDeviceSnList []string

// GetDeviceByIDResponseDayZeroConfig is the getDeviceByIDResponseDayZeroConfig definition
type GetDeviceByIDResponseDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfo is the getDeviceByIDResponseDeviceInfo definition
type GetDeviceByIDResponseDeviceInfo struct {
	AAACredentials            GetDeviceByIDResponseDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                               `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                              `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                              `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                `json:"cmState,omitempty"`                   //
	Description               string                                                `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                              `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                              `json:"featuresSupported,omitempty"`         //
	FileSystemList            []GetDeviceByIDResponseDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                               `json:"firstContact,omitempty"`              //
	Hostname                  string                                                `json:"hostname,omitempty"`                  //
	HTTPHeaders               []GetDeviceByIDResponseDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                `json:"imageVersion,omitempty"`              //
	IPInterfaces              []GetDeviceByIDResponseDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                               `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                               `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                               `json:"lastUpdateOn,omitempty"`              //
	Location                  GetDeviceByIDResponseDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                `json:"macAddress,omitempty"`                //
	Mode                      string                                                `json:"mode,omitempty"`                      //
	Name                      string                                                `json:"name,omitempty"`                      //
	NeighborLinks             []GetDeviceByIDResponseDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                `json:"onbState,omitempty"`                  //
	Pid                       string                                                `json:"pid,omitempty"`                       //
	PnpProfileList            []GetDeviceByIDResponseDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                  `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []GetDeviceByIDResponseDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                `json:"projectId,omitempty"`                 //
	ProjectName               string                                                `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                  `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                `json:"siteId,omitempty"`                    //
	SiteName                  string                                                `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                `json:"smartAccountId,omitempty"`            //
	Source                    string                                                `json:"source,omitempty"`                    //
	Stack                     bool                                                  `json:"stack,omitempty"`                     //
	StackInfo                 GetDeviceByIDResponseDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                `json:"state,omitempty"`                     //
	SudiRequired              bool                                                  `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                              `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                              `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                `json:"workflowName,omitempty"`              //
}

// GetDeviceByIDResponseDeviceInfoAAACredentials is the getDeviceByIDResponseDeviceInfoAAACredentials definition
type GetDeviceByIDResponseDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfoAddnMacAddrs is the getDeviceByIDResponseDeviceInfoAddnMacAddrs definition
type GetDeviceByIDResponseDeviceInfoAddnMacAddrs []string

// GetDeviceByIDResponseDeviceInfoCapabilitiesSupported is the getDeviceByIDResponseDeviceInfoCapabilitiesSupported definition
type GetDeviceByIDResponseDeviceInfoCapabilitiesSupported []string

// GetDeviceByIDResponseDeviceInfoDeviceSudiSerialNos is the getDeviceByIDResponseDeviceInfoDeviceSudiSerialNos definition
type GetDeviceByIDResponseDeviceInfoDeviceSudiSerialNos []string

// GetDeviceByIDResponseDeviceInfoFeaturesSupported is the getDeviceByIDResponseDeviceInfoFeaturesSupported definition
type GetDeviceByIDResponseDeviceInfoFeaturesSupported []string

// GetDeviceByIDResponseDeviceInfoFileSystemList is the getDeviceByIDResponseDeviceInfoFileSystemList definition
type GetDeviceByIDResponseDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfoHTTPHeaders is the getDeviceByIDResponseDeviceInfoHTTPHeaders definition
type GetDeviceByIDResponseDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfoIPInterfaces is the getDeviceByIDResponseDeviceInfoIPInterfaces definition
type GetDeviceByIDResponseDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// GetDeviceByIDResponseDeviceInfoIPInterfacesIPv6AddressList is the getDeviceByIDResponseDeviceInfoIPInterfacesIPv6AddressList definition
type GetDeviceByIDResponseDeviceInfoIPInterfacesIPv6AddressList []string

// GetDeviceByIDResponseDeviceInfoLocation is the getDeviceByIDResponseDeviceInfoLocation definition
type GetDeviceByIDResponseDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// GetDeviceByIDResponseDeviceInfoNeighborLinks is the getDeviceByIDResponseDeviceInfoNeighborLinks definition
type GetDeviceByIDResponseDeviceInfoNeighborLinks struct {
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

// GetDeviceByIDResponseDeviceInfoPnpProfileList is the getDeviceByIDResponseDeviceInfoPnpProfileList definition
type GetDeviceByIDResponseDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                         `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                           `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   GetDeviceByIDResponseDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                         `json:"profileName,omitempty"`       //
	SecondaryEndpoint GetDeviceByIDResponseDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfoPnpProfileListPrimaryEndpoint is the getDeviceByIDResponseDeviceInfoPnpProfileListPrimaryEndpoint definition
type GetDeviceByIDResponseDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// GetDeviceByIDResponseDeviceInfoPnpProfileListSecondaryEndpoint is the getDeviceByIDResponseDeviceInfoPnpProfileListSecondaryEndpoint definition
type GetDeviceByIDResponseDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// GetDeviceByIDResponseDeviceInfoPreWorkflowCliOuputs is the getDeviceByIDResponseDeviceInfoPreWorkflowCliOuputs definition
type GetDeviceByIDResponseDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfoStackInfo is the getDeviceByIDResponseDeviceInfoStackInfo definition
type GetDeviceByIDResponseDeviceInfoStackInfo struct {
	IsFullRing             bool                                                      `json:"isFullRing,omitempty"`             //
	StackMemberList        []GetDeviceByIDResponseDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                    `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                      `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                   `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                  `json:"validLicenseLevels,omitempty"`     //
}

// GetDeviceByIDResponseDeviceInfoStackInfoStackMemberList is the getDeviceByIDResponseDeviceInfoStackInfoStackMemberList definition
type GetDeviceByIDResponseDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// GetDeviceByIDResponseDeviceInfoStackInfoValidLicenseLevels is the getDeviceByIDResponseDeviceInfoStackInfoValidLicenseLevels definition
type GetDeviceByIDResponseDeviceInfoStackInfoValidLicenseLevels []string

// GetDeviceByIDResponseDeviceInfoUserMicNumbers is the getDeviceByIDResponseDeviceInfoUserMicNumbers definition
type GetDeviceByIDResponseDeviceInfoUserMicNumbers []string

// GetDeviceByIDResponseDeviceInfoUserSudiSerialNos is the getDeviceByIDResponseDeviceInfoUserSudiSerialNos definition
type GetDeviceByIDResponseDeviceInfoUserSudiSerialNos []string

// GetDeviceByIDResponseRunSummaryList is the getDeviceByIDResponseRunSummaryList definition
type GetDeviceByIDResponseRunSummaryList struct {
	Details         string                                             `json:"details,omitempty"`         //
	ErrorFlag       bool                                               `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo GetDeviceByIDResponseRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                            `json:"timestamp,omitempty"`       //
}

// GetDeviceByIDResponseRunSummaryListHistoryTaskInfo is the getDeviceByIDResponseRunSummaryListHistoryTaskInfo definition
type GetDeviceByIDResponseRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []GetDeviceByIDResponseRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                           `json:"name,omitempty"`         //
	TimeTaken    float64                                                          `json:"timeTaken,omitempty"`    //
	Type         string                                                           `json:"type,omitempty"`         //
	WorkItemList []GetDeviceByIDResponseRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// GetDeviceByIDResponseRunSummaryListHistoryTaskInfoAddnDetails is the getDeviceByIDResponseRunSummaryListHistoryTaskInfoAddnDetails definition
type GetDeviceByIDResponseRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetDeviceByIDResponseRunSummaryListHistoryTaskInfoWorkItemList is the getDeviceByIDResponseRunSummaryListHistoryTaskInfoWorkItemList definition
type GetDeviceByIDResponseRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetDeviceByIDResponseSystemResetWorkflow is the getDeviceByIDResponseSystemResetWorkflow definition
type GetDeviceByIDResponseSystemResetWorkflow struct {
	TypeID         string                                          `json:"_id,omitempty"`            //
	AddToInventory bool                                            `json:"addToInventory,omitempty"` //
	AddedOn        float64                                         `json:"addedOn,omitempty"`        //
	ConfigID       string                                          `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                         `json:"currTaskIdx,omitempty"`    //
	Description    string                                          `json:"description,omitempty"`    //
	EndTime        int                                             `json:"endTime,omitempty"`        //
	ExecTime       float64                                         `json:"execTime,omitempty"`       //
	ImageID        string                                          `json:"imageId,omitempty"`        //
	InstanceType   string                                          `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                         `json:"lastupdateOn,omitempty"`   //
	Name           string                                          `json:"name,omitempty"`           //
	StartTime      int                                             `json:"startTime,omitempty"`      //
	State          string                                          `json:"state,omitempty"`          //
	Tasks          []GetDeviceByIDResponseSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                          `json:"tenantId,omitempty"`       //
	Type           string                                          `json:"type,omitempty"`           //
	UseState       string                                          `json:"useState,omitempty"`       //
	Version        float64                                         `json:"version,omitempty"`        //
}

// GetDeviceByIDResponseSystemResetWorkflowTasks is the getDeviceByIDResponseSystemResetWorkflowTasks definition
type GetDeviceByIDResponseSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                     `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                         `json:"endTime,omitempty"`         //
	Name            string                                                      `json:"name,omitempty"`            //
	StartTime       int                                                         `json:"startTime,omitempty"`       //
	State           string                                                      `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                     `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                     `json:"timeTaken,omitempty"`       //
	Type            string                                                      `json:"type,omitempty"`            //
	WorkItemList    []GetDeviceByIDResponseSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetDeviceByIDResponseSystemResetWorkflowTasksWorkItemList is the getDeviceByIDResponseSystemResetWorkflowTasksWorkItemList definition
type GetDeviceByIDResponseSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetDeviceByIDResponseSystemWorkflow is the getDeviceByIDResponseSystemWorkflow definition
type GetDeviceByIDResponseSystemWorkflow struct {
	TypeID         string                                     `json:"_id,omitempty"`            //
	AddToInventory bool                                       `json:"addToInventory,omitempty"` //
	AddedOn        float64                                    `json:"addedOn,omitempty"`        //
	ConfigID       string                                     `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                    `json:"currTaskIdx,omitempty"`    //
	Description    string                                     `json:"description,omitempty"`    //
	EndTime        int                                        `json:"endTime,omitempty"`        //
	ExecTime       float64                                    `json:"execTime,omitempty"`       //
	ImageID        string                                     `json:"imageId,omitempty"`        //
	InstanceType   string                                     `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                    `json:"lastupdateOn,omitempty"`   //
	Name           string                                     `json:"name,omitempty"`           //
	StartTime      int                                        `json:"startTime,omitempty"`      //
	State          string                                     `json:"state,omitempty"`          //
	Tasks          []GetDeviceByIDResponseSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                     `json:"tenantId,omitempty"`       //
	Type           string                                     `json:"type,omitempty"`           //
	UseState       string                                     `json:"useState,omitempty"`       //
	Version        float64                                    `json:"version,omitempty"`        //
}

// GetDeviceByIDResponseSystemWorkflowTasks is the getDeviceByIDResponseSystemWorkflowTasks definition
type GetDeviceByIDResponseSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                    `json:"endTime,omitempty"`         //
	Name            string                                                 `json:"name,omitempty"`            //
	StartTime       int                                                    `json:"startTime,omitempty"`       //
	State           string                                                 `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                `json:"timeTaken,omitempty"`       //
	Type            string                                                 `json:"type,omitempty"`            //
	WorkItemList    []GetDeviceByIDResponseSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetDeviceByIDResponseSystemWorkflowTasksWorkItemList is the getDeviceByIDResponseSystemWorkflowTasksWorkItemList definition
type GetDeviceByIDResponseSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetDeviceByIDResponseWorkflow is the getDeviceByIDResponseWorkflow definition
type GetDeviceByIDResponseWorkflow struct {
	TypeID         string                               `json:"_id,omitempty"`            //
	AddToInventory bool                                 `json:"addToInventory,omitempty"` //
	AddedOn        float64                              `json:"addedOn,omitempty"`        //
	ConfigID       string                               `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                              `json:"currTaskIdx,omitempty"`    //
	Description    string                               `json:"description,omitempty"`    //
	EndTime        int                                  `json:"endTime,omitempty"`        //
	ExecTime       float64                              `json:"execTime,omitempty"`       //
	ImageID        string                               `json:"imageId,omitempty"`        //
	InstanceType   string                               `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                              `json:"lastupdateOn,omitempty"`   //
	Name           string                               `json:"name,omitempty"`           //
	StartTime      int                                  `json:"startTime,omitempty"`      //
	State          string                               `json:"state,omitempty"`          //
	Tasks          []GetDeviceByIDResponseWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                               `json:"tenantId,omitempty"`       //
	Type           string                               `json:"type,omitempty"`           //
	UseState       string                               `json:"useState,omitempty"`       //
	Version        float64                              `json:"version,omitempty"`        //
}

// GetDeviceByIDResponseWorkflowParameters is the getDeviceByIDResponseWorkflowParameters definition
type GetDeviceByIDResponseWorkflowParameters struct {
	ConfigList             []GetDeviceByIDResponseWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                              `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                              `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                              `json:"topOfStackSerialNumber,omitempty"` //
}

// GetDeviceByIDResponseWorkflowParametersConfigList is the getDeviceByIDResponseWorkflowParametersConfigList definition
type GetDeviceByIDResponseWorkflowParametersConfigList struct {
	ConfigID         string                                                              `json:"configId,omitempty"`         //
	ConfigParameters []GetDeviceByIDResponseWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// GetDeviceByIDResponseWorkflowParametersConfigListConfigParameters is the getDeviceByIDResponseWorkflowParametersConfigListConfigParameters definition
type GetDeviceByIDResponseWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetDeviceByIDResponseWorkflowTasks is the getDeviceByIDResponseWorkflowTasks definition
type GetDeviceByIDResponseWorkflowTasks struct {
	CurrWorkItemIDx float64                                          `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                              `json:"endTime,omitempty"`         //
	Name            string                                           `json:"name,omitempty"`            //
	StartTime       int                                              `json:"startTime,omitempty"`       //
	State           string                                           `json:"state,omitempty"`           //
	TaskSeqNo       float64                                          `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                          `json:"timeTaken,omitempty"`       //
	Type            string                                           `json:"type,omitempty"`            //
	WorkItemList    []GetDeviceByIDResponseWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetDeviceByIDResponseWorkflowTasksWorkItemList is the getDeviceByIDResponseWorkflowTasksWorkItemList definition
type GetDeviceByIDResponseWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetDeviceByIDResponse is the getDeviceByIdResponse definition
type GetDeviceByIDResponse struct {
	TypeID               string                                   `json:"_id,omitempty"`                  //
	DayZeroConfig        GetDeviceByIDResponseDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                   `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           GetDeviceByIDResponseDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []GetDeviceByIDResponseRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  GetDeviceByIDResponseSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       GetDeviceByIDResponseSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                   `json:"tenantId,omitempty"`             //
	Version              float64                                  `json:"version,omitempty"`              //
	Workflow             GetDeviceByIDResponseWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   GetDeviceByIDResponseWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// GetDeviceHistoryResponse is the getDeviceHistoryResponse definition
type GetDeviceHistoryResponse struct {
	Response   []GetDeviceHistoryResponseResponse `json:"response,omitempty"`   //
	StatusCode float64                            `json:"statusCode,omitempty"` //
}

// GetDeviceHistoryResponseResponse is the getDeviceHistoryResponseResponse definition
type GetDeviceHistoryResponseResponse struct {
	Details         string                                          `json:"details,omitempty"`         //
	ErrorFlag       bool                                            `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo GetDeviceHistoryResponseResponseHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                         `json:"timestamp,omitempty"`       //
}

// GetDeviceHistoryResponseResponseHistoryTaskInfo is the getDeviceHistoryResponseResponseHistoryTaskInfo definition
type GetDeviceHistoryResponseResponseHistoryTaskInfo struct {
	AddnDetails  []GetDeviceHistoryResponseResponseHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                        `json:"name,omitempty"`         //
	TimeTaken    float64                                                       `json:"timeTaken,omitempty"`    //
	Type         string                                                        `json:"type,omitempty"`         //
	WorkItemList []GetDeviceHistoryResponseResponseHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// GetDeviceHistoryResponseResponseHistoryTaskInfoAddnDetails is the getDeviceHistoryResponseResponseHistoryTaskInfoAddnDetails definition
type GetDeviceHistoryResponseResponseHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetDeviceHistoryResponseResponseHistoryTaskInfoWorkItemList is the getDeviceHistoryResponseResponseHistoryTaskInfoWorkItemList definition
type GetDeviceHistoryResponseResponseHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetPnPGlobalSettingsResponse is the getPnPGlobalSettingsResponse definition
type GetPnPGlobalSettingsResponse struct {
	TypeID          string                                        `json:"_id,omitempty"`             //
	AAACredentials  GetPnPGlobalSettingsResponseAAACredentials    `json:"aaaCredentials,omitempty"`  //
	AcceptEula      bool                                          `json:"acceptEula,omitempty"`      //
	DefaultProfile  GetPnPGlobalSettingsResponseDefaultProfile    `json:"defaultProfile,omitempty"`  //
	ID              string                                        `json:"id,omitempty"`              //
	SavaMappingList []GetPnPGlobalSettingsResponseSavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    GetPnPGlobalSettingsResponseTaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                        `json:"tenantId,omitempty"`        //
	Version         float64                                       `json:"version,omitempty"`         //
}

// GetPnPGlobalSettingsResponseAAACredentials is the getPnPGlobalSettingsResponseAAACredentials definition
type GetPnPGlobalSettingsResponseAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// GetPnPGlobalSettingsResponseDefaultProfile is the getPnPGlobalSettingsResponseDefaultProfile definition
type GetPnPGlobalSettingsResponseDefaultProfile struct {
	Cert          string   `json:"cert,omitempty"`          //
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` //
	IPAddresses   []string `json:"ipAddresses,omitempty"`   //
	Port          float64  `json:"port,omitempty"`          //
	Proxy         bool     `json:"proxy,omitempty"`         //
}

// GetPnPGlobalSettingsResponseDefaultProfileFqdnAddresses is the getPnPGlobalSettingsResponseDefaultProfileFqdnAddresses definition
type GetPnPGlobalSettingsResponseDefaultProfileFqdnAddresses []string

// GetPnPGlobalSettingsResponseDefaultProfileIPAddresses is the getPnPGlobalSettingsResponseDefaultProfileIPAddresses definition
type GetPnPGlobalSettingsResponseDefaultProfileIPAddresses []string

// GetPnPGlobalSettingsResponseSavaMappingList is the getPnPGlobalSettingsResponseSavaMappingList definition
type GetPnPGlobalSettingsResponseSavaMappingList struct {
	AutoSyncPeriod   float64                                               `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                `json:"ccoUser,omitempty"`          //
	Expiry           float64                                               `json:"expiry,omitempty"`           //
	LastSync         float64                                               `json:"lastSync,omitempty"`         //
	Profile          GetPnPGlobalSettingsResponseSavaMappingListProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                `json:"smartAccountId,omitempty"`   //
	SyncResult       GetPnPGlobalSettingsResponseSavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                                               `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                `json:"syncStatus,omitempty"`       //
	TenantID         string                                                `json:"tenantId,omitempty"`         //
	Token            string                                                `json:"token,omitempty"`            //
	VirtualAccountID string                                                `json:"virtualAccountId,omitempty"` //
}

// GetPnPGlobalSettingsResponseSavaMappingListProfile is the getPnPGlobalSettingsResponseSavaMappingListProfile definition
type GetPnPGlobalSettingsResponseSavaMappingListProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// GetPnPGlobalSettingsResponseSavaMappingListSyncResult is the getPnPGlobalSettingsResponseSavaMappingListSyncResult definition
type GetPnPGlobalSettingsResponseSavaMappingListSyncResult struct {
	SyncList []GetPnPGlobalSettingsResponseSavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                          `json:"syncMsg,omitempty"`  //
}

// GetPnPGlobalSettingsResponseSavaMappingListSyncResultSyncList is the getPnPGlobalSettingsResponseSavaMappingListSyncResultSyncList definition
type GetPnPGlobalSettingsResponseSavaMappingListSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// GetPnPGlobalSettingsResponseSavaMappingListSyncResultSyncListDeviceSnList is the getPnPGlobalSettingsResponseSavaMappingListSyncResultSyncListDeviceSnList definition
type GetPnPGlobalSettingsResponseSavaMappingListSyncResultSyncListDeviceSnList []string

// GetPnPGlobalSettingsResponseTaskTimeOuts is the getPnPGlobalSettingsResponseTaskTimeOuts definition
type GetPnPGlobalSettingsResponseTaskTimeOuts struct {
	ConfigTimeOut        float64 `json:"configTimeOut,omitempty"`        //
	GeneralTimeOut       float64 `json:"generalTimeOut,omitempty"`       //
	ImageDownloadTimeOut float64 `json:"imageDownloadTimeOut,omitempty"` //
}

// GetPnpDeviceCountResponse is the getPnpDeviceCountResponse definition
type GetPnpDeviceCountResponse struct {
	Response float64 `json:"response,omitempty"` //
}

// GetPnpDeviceListResponse is the getPnpDeviceListResponse definition
type GetPnpDeviceListResponse struct {
	TypeID               string                                      `json:"_id,omitempty"`                  //
	DayZeroConfig        GetPnpDeviceListResponseDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                      `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           GetPnpDeviceListResponseDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []GetPnpDeviceListResponseRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  GetPnpDeviceListResponseSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       GetPnpDeviceListResponseSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                      `json:"tenantId,omitempty"`             //
	Version              float64                                     `json:"version,omitempty"`              //
	Workflow             GetPnpDeviceListResponseWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   GetPnpDeviceListResponseWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// GetPnpDeviceListResponseDayZeroConfig is the getPnpDeviceListResponseDayZeroConfig definition
type GetPnpDeviceListResponseDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfo is the getPnpDeviceListResponseDeviceInfo definition
type GetPnpDeviceListResponseDeviceInfo struct {
	AAACredentials            GetPnpDeviceListResponseDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                                  `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                 `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                   `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                   `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                   `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                   `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                 `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                   `json:"cmState,omitempty"`                   //
	Description               string                                                   `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                 `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                   `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                 `json:"featuresSupported,omitempty"`         //
	FileSystemList            []GetPnpDeviceListResponseDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                                  `json:"firstContact,omitempty"`              //
	Hostname                  string                                                   `json:"hostname,omitempty"`                  //
	HTTPHeaders               []GetPnpDeviceListResponseDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                   `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                   `json:"imageVersion,omitempty"`              //
	IPInterfaces              []GetPnpDeviceListResponseDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                                  `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                                  `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                                  `json:"lastUpdateOn,omitempty"`              //
	Location                  GetPnpDeviceListResponseDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                   `json:"macAddress,omitempty"`                //
	Mode                      string                                                   `json:"mode,omitempty"`                      //
	Name                      string                                                   `json:"name,omitempty"`                      //
	NeighborLinks             []GetPnpDeviceListResponseDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                   `json:"onbState,omitempty"`                  //
	Pid                       string                                                   `json:"pid,omitempty"`                       //
	PnpProfileList            []GetPnpDeviceListResponseDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                     `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []GetPnpDeviceListResponseDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                   `json:"projectId,omitempty"`                 //
	ProjectName               string                                                   `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                     `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                   `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                   `json:"siteId,omitempty"`                    //
	SiteName                  string                                                   `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                   `json:"smartAccountId,omitempty"`            //
	Source                    string                                                   `json:"source,omitempty"`                    //
	Stack                     bool                                                     `json:"stack,omitempty"`                     //
	StackInfo                 GetPnpDeviceListResponseDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                   `json:"state,omitempty"`                     //
	SudiRequired              bool                                                     `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                   `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                                 `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                                 `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                   `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                   `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                   `json:"workflowName,omitempty"`              //
}

// GetPnpDeviceListResponseDeviceInfoAAACredentials is the getPnpDeviceListResponseDeviceInfoAAACredentials definition
type GetPnpDeviceListResponseDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfoAddnMacAddrs is the getPnpDeviceListResponseDeviceInfoAddnMacAddrs definition
type GetPnpDeviceListResponseDeviceInfoAddnMacAddrs []string

// GetPnpDeviceListResponseDeviceInfoCapabilitiesSupported is the getPnpDeviceListResponseDeviceInfoCapabilitiesSupported definition
type GetPnpDeviceListResponseDeviceInfoCapabilitiesSupported []string

// GetPnpDeviceListResponseDeviceInfoDeviceSudiSerialNos is the getPnpDeviceListResponseDeviceInfoDeviceSudiSerialNos definition
type GetPnpDeviceListResponseDeviceInfoDeviceSudiSerialNos []string

// GetPnpDeviceListResponseDeviceInfoFeaturesSupported is the getPnpDeviceListResponseDeviceInfoFeaturesSupported definition
type GetPnpDeviceListResponseDeviceInfoFeaturesSupported []string

// GetPnpDeviceListResponseDeviceInfoFileSystemList is the getPnpDeviceListResponseDeviceInfoFileSystemList definition
type GetPnpDeviceListResponseDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfoHTTPHeaders is the getPnpDeviceListResponseDeviceInfoHTTPHeaders definition
type GetPnpDeviceListResponseDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfoIPInterfaces is the getPnpDeviceListResponseDeviceInfoIPInterfaces definition
type GetPnpDeviceListResponseDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// GetPnpDeviceListResponseDeviceInfoIPInterfacesIPv6AddressList is the getPnpDeviceListResponseDeviceInfoIPInterfacesIPv6AddressList definition
type GetPnpDeviceListResponseDeviceInfoIPInterfacesIPv6AddressList []string

// GetPnpDeviceListResponseDeviceInfoLocation is the getPnpDeviceListResponseDeviceInfoLocation definition
type GetPnpDeviceListResponseDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// GetPnpDeviceListResponseDeviceInfoNeighborLinks is the getPnpDeviceListResponseDeviceInfoNeighborLinks definition
type GetPnpDeviceListResponseDeviceInfoNeighborLinks struct {
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

// GetPnpDeviceListResponseDeviceInfoPnpProfileList is the getPnpDeviceListResponseDeviceInfoPnpProfileList definition
type GetPnpDeviceListResponseDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                            `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                              `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   GetPnpDeviceListResponseDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                            `json:"profileName,omitempty"`       //
	SecondaryEndpoint GetPnpDeviceListResponseDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfoPnpProfileListPrimaryEndpoint is the getPnpDeviceListResponseDeviceInfoPnpProfileListPrimaryEndpoint definition
type GetPnpDeviceListResponseDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// GetPnpDeviceListResponseDeviceInfoPnpProfileListSecondaryEndpoint is the getPnpDeviceListResponseDeviceInfoPnpProfileListSecondaryEndpoint definition
type GetPnpDeviceListResponseDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// GetPnpDeviceListResponseDeviceInfoPreWorkflowCliOuputs is the getPnpDeviceListResponseDeviceInfoPreWorkflowCliOuputs definition
type GetPnpDeviceListResponseDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfoStackInfo is the getPnpDeviceListResponseDeviceInfoStackInfo definition
type GetPnpDeviceListResponseDeviceInfoStackInfo struct {
	IsFullRing             bool                                                         `json:"isFullRing,omitempty"`             //
	StackMemberList        []GetPnpDeviceListResponseDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                       `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                         `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                      `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                     `json:"validLicenseLevels,omitempty"`     //
}

// GetPnpDeviceListResponseDeviceInfoStackInfoStackMemberList is the getPnpDeviceListResponseDeviceInfoStackInfoStackMemberList definition
type GetPnpDeviceListResponseDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// GetPnpDeviceListResponseDeviceInfoStackInfoValidLicenseLevels is the getPnpDeviceListResponseDeviceInfoStackInfoValidLicenseLevels definition
type GetPnpDeviceListResponseDeviceInfoStackInfoValidLicenseLevels []string

// GetPnpDeviceListResponseDeviceInfoUserMicNumbers is the getPnpDeviceListResponseDeviceInfoUserMicNumbers definition
type GetPnpDeviceListResponseDeviceInfoUserMicNumbers []string

// GetPnpDeviceListResponseDeviceInfoUserSudiSerialNos is the getPnpDeviceListResponseDeviceInfoUserSudiSerialNos definition
type GetPnpDeviceListResponseDeviceInfoUserSudiSerialNos []string

// GetPnpDeviceListResponseRunSummaryList is the getPnpDeviceListResponseRunSummaryList definition
type GetPnpDeviceListResponseRunSummaryList struct {
	Details         string                                                `json:"details,omitempty"`         //
	ErrorFlag       bool                                                  `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo GetPnpDeviceListResponseRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                               `json:"timestamp,omitempty"`       //
}

// GetPnpDeviceListResponseRunSummaryListHistoryTaskInfo is the getPnpDeviceListResponseRunSummaryListHistoryTaskInfo definition
type GetPnpDeviceListResponseRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                              `json:"name,omitempty"`         //
	TimeTaken    float64                                                             `json:"timeTaken,omitempty"`    //
	Type         string                                                              `json:"type,omitempty"`         //
	WorkItemList []GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoAddnDetails is the getPnpDeviceListResponseRunSummaryListHistoryTaskInfoAddnDetails definition
type GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoWorkItemList is the getPnpDeviceListResponseRunSummaryListHistoryTaskInfoWorkItemList definition
type GetPnpDeviceListResponseRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetPnpDeviceListResponseSystemResetWorkflow is the getPnpDeviceListResponseSystemResetWorkflow definition
type GetPnpDeviceListResponseSystemResetWorkflow struct {
	TypeID         string                                             `json:"_id,omitempty"`            //
	AddToInventory bool                                               `json:"addToInventory,omitempty"` //
	AddedOn        float64                                            `json:"addedOn,omitempty"`        //
	ConfigID       string                                             `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                            `json:"currTaskIdx,omitempty"`    //
	Description    string                                             `json:"description,omitempty"`    //
	EndTime        int                                                `json:"endTime,omitempty"`        //
	ExecTime       float64                                            `json:"execTime,omitempty"`       //
	ImageID        string                                             `json:"imageId,omitempty"`        //
	InstanceType   string                                             `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                            `json:"lastupdateOn,omitempty"`   //
	Name           string                                             `json:"name,omitempty"`           //
	StartTime      int                                                `json:"startTime,omitempty"`      //
	State          string                                             `json:"state,omitempty"`          //
	Tasks          []GetPnpDeviceListResponseSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                             `json:"tenantId,omitempty"`       //
	Type           string                                             `json:"type,omitempty"`           //
	UseState       string                                             `json:"useState,omitempty"`       //
	Version        float64                                            `json:"version,omitempty"`        //
}

// GetPnpDeviceListResponseSystemResetWorkflowTasks is the getPnpDeviceListResponseSystemResetWorkflowTasks definition
type GetPnpDeviceListResponseSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                        `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                            `json:"endTime,omitempty"`         //
	Name            string                                                         `json:"name,omitempty"`            //
	StartTime       int                                                            `json:"startTime,omitempty"`       //
	State           string                                                         `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                        `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                        `json:"timeTaken,omitempty"`       //
	Type            string                                                         `json:"type,omitempty"`            //
	WorkItemList    []GetPnpDeviceListResponseSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetPnpDeviceListResponseSystemResetWorkflowTasksWorkItemList is the getPnpDeviceListResponseSystemResetWorkflowTasksWorkItemList definition
type GetPnpDeviceListResponseSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetPnpDeviceListResponseSystemWorkflow is the getPnpDeviceListResponseSystemWorkflow definition
type GetPnpDeviceListResponseSystemWorkflow struct {
	TypeID         string                                        `json:"_id,omitempty"`            //
	AddToInventory bool                                          `json:"addToInventory,omitempty"` //
	AddedOn        float64                                       `json:"addedOn,omitempty"`        //
	ConfigID       string                                        `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                       `json:"currTaskIdx,omitempty"`    //
	Description    string                                        `json:"description,omitempty"`    //
	EndTime        int                                           `json:"endTime,omitempty"`        //
	ExecTime       float64                                       `json:"execTime,omitempty"`       //
	ImageID        string                                        `json:"imageId,omitempty"`        //
	InstanceType   string                                        `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                       `json:"lastupdateOn,omitempty"`   //
	Name           string                                        `json:"name,omitempty"`           //
	StartTime      int                                           `json:"startTime,omitempty"`      //
	State          string                                        `json:"state,omitempty"`          //
	Tasks          []GetPnpDeviceListResponseSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                        `json:"tenantId,omitempty"`       //
	Type           string                                        `json:"type,omitempty"`           //
	UseState       string                                        `json:"useState,omitempty"`       //
	Version        float64                                       `json:"version,omitempty"`        //
}

// GetPnpDeviceListResponseSystemWorkflowTasks is the getPnpDeviceListResponseSystemWorkflowTasks definition
type GetPnpDeviceListResponseSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                   `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                       `json:"endTime,omitempty"`         //
	Name            string                                                    `json:"name,omitempty"`            //
	StartTime       int                                                       `json:"startTime,omitempty"`       //
	State           string                                                    `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                   `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                   `json:"timeTaken,omitempty"`       //
	Type            string                                                    `json:"type,omitempty"`            //
	WorkItemList    []GetPnpDeviceListResponseSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetPnpDeviceListResponseSystemWorkflowTasksWorkItemList is the getPnpDeviceListResponseSystemWorkflowTasksWorkItemList definition
type GetPnpDeviceListResponseSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetPnpDeviceListResponseWorkflow is the getPnpDeviceListResponseWorkflow definition
type GetPnpDeviceListResponseWorkflow struct {
	TypeID         string                                  `json:"_id,omitempty"`            //
	AddToInventory bool                                    `json:"addToInventory,omitempty"` //
	AddedOn        float64                                 `json:"addedOn,omitempty"`        //
	ConfigID       string                                  `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                 `json:"currTaskIdx,omitempty"`    //
	Description    string                                  `json:"description,omitempty"`    //
	EndTime        int                                     `json:"endTime,omitempty"`        //
	ExecTime       float64                                 `json:"execTime,omitempty"`       //
	ImageID        string                                  `json:"imageId,omitempty"`        //
	InstanceType   string                                  `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                 `json:"lastupdateOn,omitempty"`   //
	Name           string                                  `json:"name,omitempty"`           //
	StartTime      int                                     `json:"startTime,omitempty"`      //
	State          string                                  `json:"state,omitempty"`          //
	Tasks          []GetPnpDeviceListResponseWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                  `json:"tenantId,omitempty"`       //
	Type           string                                  `json:"type,omitempty"`           //
	UseState       string                                  `json:"useState,omitempty"`       //
	Version        float64                                 `json:"version,omitempty"`        //
}

// GetPnpDeviceListResponseWorkflowParameters is the getPnpDeviceListResponseWorkflowParameters definition
type GetPnpDeviceListResponseWorkflowParameters struct {
	ConfigList             []GetPnpDeviceListResponseWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                 `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                 `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                 `json:"topOfStackSerialNumber,omitempty"` //
}

// GetPnpDeviceListResponseWorkflowParametersConfigList is the getPnpDeviceListResponseWorkflowParametersConfigList definition
type GetPnpDeviceListResponseWorkflowParametersConfigList struct {
	ConfigID         string                                                                 `json:"configId,omitempty"`         //
	ConfigParameters []GetPnpDeviceListResponseWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// GetPnpDeviceListResponseWorkflowParametersConfigListConfigParameters is the getPnpDeviceListResponseWorkflowParametersConfigListConfigParameters definition
type GetPnpDeviceListResponseWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// GetPnpDeviceListResponseWorkflowTasks is the getPnpDeviceListResponseWorkflowTasks definition
type GetPnpDeviceListResponseWorkflowTasks struct {
	CurrWorkItemIDx float64                                             `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                 `json:"endTime,omitempty"`         //
	Name            string                                              `json:"name,omitempty"`            //
	StartTime       int                                                 `json:"startTime,omitempty"`       //
	State           string                                              `json:"state,omitempty"`           //
	TaskSeqNo       float64                                             `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                             `json:"timeTaken,omitempty"`       //
	Type            string                                              `json:"type,omitempty"`            //
	WorkItemList    []GetPnpDeviceListResponseWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetPnpDeviceListResponseWorkflowTasksWorkItemList is the getPnpDeviceListResponseWorkflowTasksWorkItemList definition
type GetPnpDeviceListResponseWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetSmartAccountListResponse is the getSmartAccountListResponse definition
type GetSmartAccountListResponse []string

// GetSyncResultForVirtualAccountResponse is the getSyncResultForVirtualAccountResponse definition
type GetSyncResultForVirtualAccountResponse struct {
	AutoSyncPeriod   float64                                          `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                           `json:"ccoUser,omitempty"`          //
	Expiry           float64                                          `json:"expiry,omitempty"`           //
	LastSync         float64                                          `json:"lastSync,omitempty"`         //
	Profile          GetSyncResultForVirtualAccountResponseProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                           `json:"smartAccountId,omitempty"`   //
	SyncResult       GetSyncResultForVirtualAccountResponseSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                           `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                                          `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                           `json:"syncStatus,omitempty"`       //
	TenantID         string                                           `json:"tenantId,omitempty"`         //
	Token            string                                           `json:"token,omitempty"`            //
	VirtualAccountID string                                           `json:"virtualAccountId,omitempty"` //
}

// GetSyncResultForVirtualAccountResponseProfile is the getSyncResultForVirtualAccountResponseProfile definition
type GetSyncResultForVirtualAccountResponseProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// GetSyncResultForVirtualAccountResponseSyncResult is the getSyncResultForVirtualAccountResponseSyncResult definition
type GetSyncResultForVirtualAccountResponseSyncResult struct {
	SyncList []GetSyncResultForVirtualAccountResponseSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                     `json:"syncMsg,omitempty"`  //
}

// GetSyncResultForVirtualAccountResponseSyncResultSyncList is the getSyncResultForVirtualAccountResponseSyncResultSyncList definition
type GetSyncResultForVirtualAccountResponseSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// GetSyncResultForVirtualAccountResponseSyncResultSyncListDeviceSnList is the getSyncResultForVirtualAccountResponseSyncResultSyncListDeviceSnList definition
type GetSyncResultForVirtualAccountResponseSyncResultSyncListDeviceSnList []string

// GetVirtualAccountListResponse is the getVirtualAccountListResponse definition
type GetVirtualAccountListResponse []string

// GetWorkflowByIDResponseTasks is the getWorkflowByIDResponseTasks definition
type GetWorkflowByIDResponseTasks struct {
	CurrWorkItemIDx float64                                    `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                        `json:"endTime,omitempty"`         //
	Name            string                                     `json:"name,omitempty"`            //
	StartTime       int                                        `json:"startTime,omitempty"`       //
	State           string                                     `json:"state,omitempty"`           //
	TaskSeqNo       float64                                    `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                    `json:"timeTaken,omitempty"`       //
	Type            string                                     `json:"type,omitempty"`            //
	WorkItemList    []GetWorkflowByIDResponseTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetWorkflowByIDResponseTasksWorkItemList is the getWorkflowByIDResponseTasksWorkItemList definition
type GetWorkflowByIDResponseTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// GetWorkflowByIDResponse is the getWorkflowByIdResponse definition
type GetWorkflowByIDResponse struct {
	TypeID         string                         `json:"_id,omitempty"`            //
	AddToInventory bool                           `json:"addToInventory,omitempty"` //
	AddedOn        float64                        `json:"addedOn,omitempty"`        //
	ConfigID       string                         `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                        `json:"currTaskIdx,omitempty"`    //
	Description    string                         `json:"description,omitempty"`    //
	EndTime        int                            `json:"endTime,omitempty"`        //
	ExecTime       float64                        `json:"execTime,omitempty"`       //
	ImageID        string                         `json:"imageId,omitempty"`        //
	InstanceType   string                         `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                        `json:"lastupdateOn,omitempty"`   //
	Name           string                         `json:"name,omitempty"`           //
	StartTime      int                            `json:"startTime,omitempty"`      //
	State          string                         `json:"state,omitempty"`          //
	Tasks          []GetWorkflowByIDResponseTasks `json:"tasks,omitempty"`          //
	TenantID       string                         `json:"tenantId,omitempty"`       //
	Type           string                         `json:"type,omitempty"`           //
	UseState       string                         `json:"useState,omitempty"`       //
	Version        float64                        `json:"version,omitempty"`        //
}

// GetWorkflowCountResponse is the getWorkflowCountResponse definition
type GetWorkflowCountResponse struct {
	Response float64 `json:"response,omitempty"` //
}

// GetWorkflowsResponse is the getWorkflowsResponse definition
type GetWorkflowsResponse struct {
	TypeID         string                      `json:"_id,omitempty"`            //
	AddToInventory bool                        `json:"addToInventory,omitempty"` //
	AddedOn        float64                     `json:"addedOn,omitempty"`        //
	ConfigID       string                      `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                     `json:"currTaskIdx,omitempty"`    //
	Description    string                      `json:"description,omitempty"`    //
	EndTime        int                         `json:"endTime,omitempty"`        //
	ExecTime       float64                     `json:"execTime,omitempty"`       //
	ImageID        string                      `json:"imageId,omitempty"`        //
	InstanceType   string                      `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                     `json:"lastupdateOn,omitempty"`   //
	Name           string                      `json:"name,omitempty"`           //
	StartTime      int                         `json:"startTime,omitempty"`      //
	State          string                      `json:"state,omitempty"`          //
	Tasks          []GetWorkflowsResponseTasks `json:"tasks,omitempty"`          //
	TenantID       string                      `json:"tenantId,omitempty"`       //
	Type           string                      `json:"type,omitempty"`           //
	UseState       string                      `json:"useState,omitempty"`       //
	Version        float64                     `json:"version,omitempty"`        //
}

// GetWorkflowsResponseTasks is the getWorkflowsResponseTasks definition
type GetWorkflowsResponseTasks struct {
	CurrWorkItemIDx float64                                 `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                     `json:"endTime,omitempty"`         //
	Name            string                                  `json:"name,omitempty"`            //
	StartTime       int                                     `json:"startTime,omitempty"`       //
	State           string                                  `json:"state,omitempty"`           //
	TaskSeqNo       float64                                 `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                 `json:"timeTaken,omitempty"`       //
	Type            string                                  `json:"type,omitempty"`            //
	WorkItemList    []GetWorkflowsResponseTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// GetWorkflowsResponseTasksWorkItemList is the getWorkflowsResponseTasksWorkItemList definition
type GetWorkflowsResponseTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkResponse is the importDevicesInBulkResponse definition
type ImportDevicesInBulkResponse struct {
	FailureList []ImportDevicesInBulkResponseFailureList `json:"failureList,omitempty"` //
	SuccessList []ImportDevicesInBulkResponseSuccessList `json:"successList,omitempty"` //
}

// ImportDevicesInBulkResponseFailureList is the importDevicesInBulkResponseFailureList definition
type ImportDevicesInBulkResponseFailureList struct {
	ID        string  `json:"id,omitempty"`        //
	Index     float64 `json:"index,omitempty"`     //
	Msg       string  `json:"msg,omitempty"`       //
	SerialNum string  `json:"serialNum,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessList is the importDevicesInBulkResponseSuccessList definition
type ImportDevicesInBulkResponseSuccessList struct {
	TypeID               string                                                    `json:"_id,omitempty"`                  //
	DayZeroConfig        ImportDevicesInBulkResponseSuccessListDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                                    `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           ImportDevicesInBulkResponseSuccessListDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []ImportDevicesInBulkResponseSuccessListRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  ImportDevicesInBulkResponseSuccessListSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       ImportDevicesInBulkResponseSuccessListSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                                    `json:"tenantId,omitempty"`             //
	Version              float64                                                   `json:"version,omitempty"`              //
	Workflow             ImportDevicesInBulkResponseSuccessListWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   ImportDevicesInBulkResponseSuccessListWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// ImportDevicesInBulkResponseSuccessListDayZeroConfig is the importDevicesInBulkResponseSuccessListDayZeroConfig definition
type ImportDevicesInBulkResponseSuccessListDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfo is the importDevicesInBulkResponseSuccessListDeviceInfo definition
type ImportDevicesInBulkResponseSuccessListDeviceInfo struct {
	AAACredentials            ImportDevicesInBulkResponseSuccessListDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                                                `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                               `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                                 `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                                 `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                                                 `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                                                 `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                               `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                                 `json:"cmState,omitempty"`                   //
	Description               string                                                                 `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                               `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                                 `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                               `json:"featuresSupported,omitempty"`         //
	FileSystemList            []ImportDevicesInBulkResponseSuccessListDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                                                `json:"firstContact,omitempty"`              //
	Hostname                  string                                                                 `json:"hostname,omitempty"`                  //
	HTTPHeaders               []ImportDevicesInBulkResponseSuccessListDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                                 `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                                 `json:"imageVersion,omitempty"`              //
	IPInterfaces              []ImportDevicesInBulkResponseSuccessListDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                                                `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                                                `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                                                `json:"lastUpdateOn,omitempty"`              //
	Location                  ImportDevicesInBulkResponseSuccessListDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                                 `json:"macAddress,omitempty"`                //
	Mode                      string                                                                 `json:"mode,omitempty"`                      //
	Name                      string                                                                 `json:"name,omitempty"`                      //
	NeighborLinks             []ImportDevicesInBulkResponseSuccessListDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                                 `json:"onbState,omitempty"`                  //
	Pid                       string                                                                 `json:"pid,omitempty"`                       //
	PnpProfileList            []ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                                   `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []ImportDevicesInBulkResponseSuccessListDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                                 `json:"projectId,omitempty"`                 //
	ProjectName               string                                                                 `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                                   `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                                 `json:"serialNumber,omitempty"`              //
	SiteID                    string                                                                 `json:"siteId,omitempty"`                    //
	SiteName                  string                                                                 `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                                                 `json:"smartAccountId,omitempty"`            //
	Source                    string                                                                 `json:"source,omitempty"`                    //
	Stack                     bool                                                                   `json:"stack,omitempty"`                     //
	StackInfo                 ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                                 `json:"state,omitempty"`                     //
	SudiRequired              bool                                                                   `json:"sudiRequired,omitempty"`              //
	Tags                      string                                                                 `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                                               `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                                               `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                                 `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                                 `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                                 `json:"workflowName,omitempty"`              //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoAAACredentials is the importDevicesInBulkResponseSuccessListDeviceInfoAAACredentials definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoAddnMacAddrs is the importDevicesInBulkResponseSuccessListDeviceInfoAddnMacAddrs definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoAddnMacAddrs []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoCapabilitiesSupported is the importDevicesInBulkResponseSuccessListDeviceInfoCapabilitiesSupported definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoCapabilitiesSupported []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoDeviceSudiSerialNos is the importDevicesInBulkResponseSuccessListDeviceInfoDeviceSudiSerialNos definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoDeviceSudiSerialNos []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoFeaturesSupported is the importDevicesInBulkResponseSuccessListDeviceInfoFeaturesSupported definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoFeaturesSupported []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoFileSystemList is the importDevicesInBulkResponseSuccessListDeviceInfoFileSystemList definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoHTTPHeaders is the importDevicesInBulkResponseSuccessListDeviceInfoHTTPHeaders definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoIPInterfaces is the importDevicesInBulkResponseSuccessListDeviceInfoIPInterfaces definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoIPInterfacesIPv6AddressList is the importDevicesInBulkResponseSuccessListDeviceInfoIPInterfacesIPv6AddressList definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoIPInterfacesIPv6AddressList []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoLocation is the importDevicesInBulkResponseSuccessListDeviceInfoLocation definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoNeighborLinks is the importDevicesInBulkResponseSuccessListDeviceInfoNeighborLinks definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoNeighborLinks struct {
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

// ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileList is the importDevicesInBulkResponseSuccessListDeviceInfoPnpProfileList definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                          `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                                            `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                          `json:"profileName,omitempty"`       //
	SecondaryEndpoint ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListPrimaryEndpoint is the importDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListPrimaryEndpoint definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListSecondaryEndpoint is the importDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListSecondaryEndpoint definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoPreWorkflowCliOuputs is the importDevicesInBulkResponseSuccessListDeviceInfoPreWorkflowCliOuputs definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfo is the importDevicesInBulkResponseSuccessListDeviceInfoStackInfo definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfo struct {
	IsFullRing             bool                                                                       `json:"isFullRing,omitempty"`             //
	StackMemberList        []ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                     `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                                       `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                                    `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                                   `json:"validLicenseLevels,omitempty"`     //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfoStackMemberList is the importDevicesInBulkResponseSuccessListDeviceInfoStackInfoStackMemberList definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfoValidLicenseLevels is the importDevicesInBulkResponseSuccessListDeviceInfoStackInfoValidLicenseLevels definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoStackInfoValidLicenseLevels []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoUserMicNumbers is the importDevicesInBulkResponseSuccessListDeviceInfoUserMicNumbers definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoUserMicNumbers []string

// ImportDevicesInBulkResponseSuccessListDeviceInfoUserSudiSerialNos is the importDevicesInBulkResponseSuccessListDeviceInfoUserSudiSerialNos definition
type ImportDevicesInBulkResponseSuccessListDeviceInfoUserSudiSerialNos []string

// ImportDevicesInBulkResponseSuccessListRunSummaryList is the importDevicesInBulkResponseSuccessListRunSummaryList definition
type ImportDevicesInBulkResponseSuccessListRunSummaryList struct {
	Details         string                                                              `json:"details,omitempty"`         //
	ErrorFlag       bool                                                                `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                                             `json:"timestamp,omitempty"`       //
}

// ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfo is the importDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfo definition
type ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                            `json:"name,omitempty"`         //
	TimeTaken    float64                                                                           `json:"timeTaken,omitempty"`    //
	Type         string                                                                            `json:"type,omitempty"`         //
	WorkItemList []ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoAddnDetails is the importDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoAddnDetails definition
type ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoWorkItemList is the importDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoWorkItemList definition
type ImportDevicesInBulkResponseSuccessListRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListSystemResetWorkflow is the importDevicesInBulkResponseSuccessListSystemResetWorkflow definition
type ImportDevicesInBulkResponseSuccessListSystemResetWorkflow struct {
	TypeID         string                                                           `json:"_id,omitempty"`            //
	AddToInventory bool                                                             `json:"addToInventory,omitempty"` //
	AddedOn        float64                                                          `json:"addedOn,omitempty"`        //
	ConfigID       string                                                           `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                                          `json:"currTaskIdx,omitempty"`    //
	Description    string                                                           `json:"description,omitempty"`    //
	EndTime        int                                                              `json:"endTime,omitempty"`        //
	ExecTime       float64                                                          `json:"execTime,omitempty"`       //
	ImageID        string                                                           `json:"imageId,omitempty"`        //
	InstanceType   string                                                           `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                                          `json:"lastupdateOn,omitempty"`   //
	Name           string                                                           `json:"name,omitempty"`           //
	StartTime      int                                                              `json:"startTime,omitempty"`      //
	State          string                                                           `json:"state,omitempty"`          //
	Tasks          []ImportDevicesInBulkResponseSuccessListSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                           `json:"tenantId,omitempty"`       //
	Type           string                                                           `json:"type,omitempty"`           //
	UseState       string                                                           `json:"useState,omitempty"`       //
	Version        float64                                                          `json:"version,omitempty"`        //
}

// ImportDevicesInBulkResponseSuccessListSystemResetWorkflowTasks is the importDevicesInBulkResponseSuccessListSystemResetWorkflowTasks definition
type ImportDevicesInBulkResponseSuccessListSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                                      `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                                          `json:"endTime,omitempty"`         //
	Name            string                                                                       `json:"name,omitempty"`            //
	StartTime       int                                                                          `json:"startTime,omitempty"`       //
	State           string                                                                       `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                                      `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                                      `json:"timeTaken,omitempty"`       //
	Type            string                                                                       `json:"type,omitempty"`            //
	WorkItemList    []ImportDevicesInBulkResponseSuccessListSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// ImportDevicesInBulkResponseSuccessListSystemResetWorkflowTasksWorkItemList is the importDevicesInBulkResponseSuccessListSystemResetWorkflowTasksWorkItemList definition
type ImportDevicesInBulkResponseSuccessListSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListSystemWorkflow is the importDevicesInBulkResponseSuccessListSystemWorkflow definition
type ImportDevicesInBulkResponseSuccessListSystemWorkflow struct {
	TypeID         string                                                      `json:"_id,omitempty"`            //
	AddToInventory bool                                                        `json:"addToInventory,omitempty"` //
	AddedOn        float64                                                     `json:"addedOn,omitempty"`        //
	ConfigID       string                                                      `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                                     `json:"currTaskIdx,omitempty"`    //
	Description    string                                                      `json:"description,omitempty"`    //
	EndTime        int                                                         `json:"endTime,omitempty"`        //
	ExecTime       float64                                                     `json:"execTime,omitempty"`       //
	ImageID        string                                                      `json:"imageId,omitempty"`        //
	InstanceType   string                                                      `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                                     `json:"lastupdateOn,omitempty"`   //
	Name           string                                                      `json:"name,omitempty"`           //
	StartTime      int                                                         `json:"startTime,omitempty"`      //
	State          string                                                      `json:"state,omitempty"`          //
	Tasks          []ImportDevicesInBulkResponseSuccessListSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                      `json:"tenantId,omitempty"`       //
	Type           string                                                      `json:"type,omitempty"`           //
	UseState       string                                                      `json:"useState,omitempty"`       //
	Version        float64                                                     `json:"version,omitempty"`        //
}

// ImportDevicesInBulkResponseSuccessListSystemWorkflowTasks is the importDevicesInBulkResponseSuccessListSystemWorkflowTasks definition
type ImportDevicesInBulkResponseSuccessListSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                                                 `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                                     `json:"endTime,omitempty"`         //
	Name            string                                                                  `json:"name,omitempty"`            //
	StartTime       int                                                                     `json:"startTime,omitempty"`       //
	State           string                                                                  `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                                 `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                                 `json:"timeTaken,omitempty"`       //
	Type            string                                                                  `json:"type,omitempty"`            //
	WorkItemList    []ImportDevicesInBulkResponseSuccessListSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// ImportDevicesInBulkResponseSuccessListSystemWorkflowTasksWorkItemList is the importDevicesInBulkResponseSuccessListSystemWorkflowTasksWorkItemList definition
type ImportDevicesInBulkResponseSuccessListSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListWorkflow is the importDevicesInBulkResponseSuccessListWorkflow definition
type ImportDevicesInBulkResponseSuccessListWorkflow struct {
	TypeID         string                                                `json:"_id,omitempty"`            //
	AddToInventory bool                                                  `json:"addToInventory,omitempty"` //
	AddedOn        float64                                               `json:"addedOn,omitempty"`        //
	ConfigID       string                                                `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                               `json:"currTaskIdx,omitempty"`    //
	Description    string                                                `json:"description,omitempty"`    //
	EndTime        int                                                   `json:"endTime,omitempty"`        //
	ExecTime       float64                                               `json:"execTime,omitempty"`       //
	ImageID        string                                                `json:"imageId,omitempty"`        //
	InstanceType   string                                                `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                               `json:"lastupdateOn,omitempty"`   //
	Name           string                                                `json:"name,omitempty"`           //
	StartTime      int                                                   `json:"startTime,omitempty"`      //
	State          string                                                `json:"state,omitempty"`          //
	Tasks          []ImportDevicesInBulkResponseSuccessListWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                `json:"tenantId,omitempty"`       //
	Type           string                                                `json:"type,omitempty"`           //
	UseState       string                                                `json:"useState,omitempty"`       //
	Version        float64                                               `json:"version,omitempty"`        //
}

// ImportDevicesInBulkResponseSuccessListWorkflowParameters is the importDevicesInBulkResponseSuccessListWorkflowParameters definition
type ImportDevicesInBulkResponseSuccessListWorkflowParameters struct {
	ConfigList             []ImportDevicesInBulkResponseSuccessListWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                               `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                               `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                               `json:"topOfStackSerialNumber,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListWorkflowParametersConfigList is the importDevicesInBulkResponseSuccessListWorkflowParametersConfigList definition
type ImportDevicesInBulkResponseSuccessListWorkflowParametersConfigList struct {
	ConfigID         string                                                                               `json:"configId,omitempty"`         //
	ConfigParameters []ImportDevicesInBulkResponseSuccessListWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListWorkflowParametersConfigListConfigParameters is the importDevicesInBulkResponseSuccessListWorkflowParametersConfigListConfigParameters definition
type ImportDevicesInBulkResponseSuccessListWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// ImportDevicesInBulkResponseSuccessListWorkflowTasks is the importDevicesInBulkResponseSuccessListWorkflowTasks definition
type ImportDevicesInBulkResponseSuccessListWorkflowTasks struct {
	CurrWorkItemIDx float64                                                           `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                               `json:"endTime,omitempty"`         //
	Name            string                                                            `json:"name,omitempty"`            //
	StartTime       int                                                               `json:"startTime,omitempty"`       //
	State           string                                                            `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                           `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                           `json:"timeTaken,omitempty"`       //
	Type            string                                                            `json:"type,omitempty"`            //
	WorkItemList    []ImportDevicesInBulkResponseSuccessListWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// ImportDevicesInBulkResponseSuccessListWorkflowTasksWorkItemList is the importDevicesInBulkResponseSuccessListWorkflowTasksWorkItemList definition
type ImportDevicesInBulkResponseSuccessListWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// PreviewConfigResponse is the previewConfigResponse definition
type PreviewConfigResponse struct {
	Response PreviewConfigResponseResponse `json:"response,omitempty"` //
	Version  string                        `json:"version,omitempty"`  //
}

// PreviewConfigResponseResponse is the previewConfigResponseResponse definition
type PreviewConfigResponseResponse struct {
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

// ResetDeviceResponse is the resetDeviceResponse definition
type ResetDeviceResponse struct {
	JSONArrayResponse []string `json:"jsonArrayResponse,omitempty"` //
	JSONResponse      string   `json:"jsonResponse,omitempty"`      //
	Message           string   `json:"message,omitempty"`           //
	StatusCode        float64  `json:"statusCode,omitempty"`        //
}

// ResetDeviceResponseJSONArrayResponse is the resetDeviceResponseJSONArrayResponse definition
type ResetDeviceResponseJSONArrayResponse []string

// SyncVirtualAccountDevicesResponse is the syncVirtualAccountDevicesResponse definition
type SyncVirtualAccountDevicesResponse struct {
	AutoSyncPeriod   float64                                     `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                      `json:"ccoUser,omitempty"`          //
	Expiry           float64                                     `json:"expiry,omitempty"`           //
	LastSync         float64                                     `json:"lastSync,omitempty"`         //
	Profile          SyncVirtualAccountDevicesResponseProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                      `json:"smartAccountId,omitempty"`   //
	SyncResult       SyncVirtualAccountDevicesResponseSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                      `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                                     `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                      `json:"syncStatus,omitempty"`       //
	TenantID         string                                      `json:"tenantId,omitempty"`         //
	Token            string                                      `json:"token,omitempty"`            //
	VirtualAccountID string                                      `json:"virtualAccountId,omitempty"` //
}

// SyncVirtualAccountDevicesResponseProfile is the syncVirtualAccountDevicesResponseProfile definition
type SyncVirtualAccountDevicesResponseProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// SyncVirtualAccountDevicesResponseSyncResult is the syncVirtualAccountDevicesResponseSyncResult definition
type SyncVirtualAccountDevicesResponseSyncResult struct {
	SyncList []SyncVirtualAccountDevicesResponseSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                `json:"syncMsg,omitempty"`  //
}

// SyncVirtualAccountDevicesResponseSyncResultSyncList is the syncVirtualAccountDevicesResponseSyncResultSyncList definition
type SyncVirtualAccountDevicesResponseSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// SyncVirtualAccountDevicesResponseSyncResultSyncListDeviceSnList is the syncVirtualAccountDevicesResponseSyncResultSyncListDeviceSnList definition
type SyncVirtualAccountDevicesResponseSyncResultSyncListDeviceSnList []string

// UnClaimDeviceResponse is the unClaimDeviceResponse definition
type UnClaimDeviceResponse struct {
	JSONArrayResponse []string `json:"jsonArrayResponse,omitempty"` //
	JSONResponse      string   `json:"jsonResponse,omitempty"`      //
	Message           string   `json:"message,omitempty"`           //
	StatusCode        float64  `json:"statusCode,omitempty"`        //
}

// UnClaimDeviceResponseJSONArrayResponse is the unClaimDeviceResponseJSONArrayResponse definition
type UnClaimDeviceResponseJSONArrayResponse []string

// UpdateDeviceResponse is the updateDeviceResponse definition
type UpdateDeviceResponse struct {
	TypeID               string                                  `json:"_id,omitempty"`                  //
	DayZeroConfig        UpdateDeviceResponseDayZeroConfig       `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview string                                  `json:"dayZeroConfigPreview,omitempty"` //
	DeviceInfo           UpdateDeviceResponseDeviceInfo          `json:"deviceInfo,omitempty"`           //
	RunSummaryList       []UpdateDeviceResponseRunSummaryList    `json:"runSummaryList,omitempty"`       //
	SystemResetWorkflow  UpdateDeviceResponseSystemResetWorkflow `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       UpdateDeviceResponseSystemWorkflow      `json:"systemWorkflow,omitempty"`       //
	TenantID             string                                  `json:"tenantId,omitempty"`             //
	Version              float64                                 `json:"version,omitempty"`              //
	Workflow             UpdateDeviceResponseWorkflow            `json:"workflow,omitempty"`             //
	WorkflowParameters   UpdateDeviceResponseWorkflowParameters  `json:"workflowParameters,omitempty"`   //
}

// UpdateDeviceResponseDayZeroConfig is the updateDeviceResponseDayZeroConfig definition
type UpdateDeviceResponseDayZeroConfig struct {
	Config string `json:"config,omitempty"` //
}

// UpdateDeviceResponseDeviceInfo is the updateDeviceResponseDeviceInfo definition
type UpdateDeviceResponseDeviceInfo struct {
	AAACredentials            UpdateDeviceResponseDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   float64                                              `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                             `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                               `json:"agentType,omitempty"`                 //
	AuthStatus                string                                               `json:"authStatus,omitempty"`                //
	AuthenticatedMicNumber    string                                               `json:"authenticatedMicNumber,omitempty"`    //
	AuthenticatedSudiSerialNo string                                               `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                             `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                               `json:"cmState,omitempty"`                   //
	Description               string                                               `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                             `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                               `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                             `json:"featuresSupported,omitempty"`         //
	FileSystemList            []UpdateDeviceResponseDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              float64                                              `json:"firstContact,omitempty"`              //
	Hostname                  string                                               `json:"hostname,omitempty"`                  //
	HTTPHeaders               []UpdateDeviceResponseDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                               `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                               `json:"imageVersion,omitempty"`              //
	IPInterfaces              []UpdateDeviceResponseDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               float64                                              `json:"lastContact,omitempty"`               //
	LastSyncTime              float64                                              `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              float64                                              `json:"lastUpdateOn,omitempty"`              //
	Location                  UpdateDeviceResponseDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                               `json:"macAddress,omitempty"`                //
	Mode                      string                                               `json:"mode,omitempty"`                      //
	Name                      string                                               `json:"name,omitempty"`                      //
	NeighborLinks             []UpdateDeviceResponseDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                               `json:"onbState,omitempty"`                  //
	Pid                       string                                               `json:"pid,omitempty"`                       //
	PnpProfileList            []UpdateDeviceResponseDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         bool                                                 `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      []UpdateDeviceResponseDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                               `json:"projectId,omitempty"`                 //
	ProjectName               string                                               `json:"projectName,omitempty"`               //
	ReloadRequested           bool                                                 `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                               `json:"serialNumber,omitempty"`              //
	SiteID                    string                                               `json:"siteId,omitempty"`                    //
	SiteName                  string                                               `json:"siteName,omitempty"`                  //
	SmartAccountID            string                                               `json:"smartAccountId,omitempty"`            //
	Source                    string                                               `json:"source,omitempty"`                    //
	Stack                     bool                                                 `json:"stack,omitempty"`                     //
	StackInfo                 UpdateDeviceResponseDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                               `json:"state,omitempty"`                     //
	SudiRequired              bool                                                 `json:"sudiRequired,omitempty"`              //
	Tags                      string                                               `json:"tags,omitempty"`                      //
	UserMicNumbers            []string                                             `json:"userMicNumbers,omitempty"`            //
	UserSudiSerialNos         []string                                             `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                               `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                               `json:"workflowId,omitempty"`                //
	WorkflowName              string                                               `json:"workflowName,omitempty"`              //
}

// UpdateDeviceResponseDeviceInfoAAACredentials is the updateDeviceResponseDeviceInfoAAACredentials definition
type UpdateDeviceResponseDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// UpdateDeviceResponseDeviceInfoAddnMacAddrs is the updateDeviceResponseDeviceInfoAddnMacAddrs definition
type UpdateDeviceResponseDeviceInfoAddnMacAddrs []string

// UpdateDeviceResponseDeviceInfoCapabilitiesSupported is the updateDeviceResponseDeviceInfoCapabilitiesSupported definition
type UpdateDeviceResponseDeviceInfoCapabilitiesSupported []string

// UpdateDeviceResponseDeviceInfoDeviceSudiSerialNos is the updateDeviceResponseDeviceInfoDeviceSudiSerialNos definition
type UpdateDeviceResponseDeviceInfoDeviceSudiSerialNos []string

// UpdateDeviceResponseDeviceInfoFeaturesSupported is the updateDeviceResponseDeviceInfoFeaturesSupported definition
type UpdateDeviceResponseDeviceInfoFeaturesSupported []string

// UpdateDeviceResponseDeviceInfoFileSystemList is the updateDeviceResponseDeviceInfoFileSystemList definition
type UpdateDeviceResponseDeviceInfoFileSystemList struct {
	Freespace float64 `json:"freespace,omitempty"` //
	Name      string  `json:"name,omitempty"`      //
	Readable  bool    `json:"readable,omitempty"`  //
	Size      float64 `json:"size,omitempty"`      //
	Type      string  `json:"type,omitempty"`      //
	Writeable bool    `json:"writeable,omitempty"` //
}

// UpdateDeviceResponseDeviceInfoHTTPHeaders is the updateDeviceResponseDeviceInfoHTTPHeaders definition
type UpdateDeviceResponseDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// UpdateDeviceResponseDeviceInfoIPInterfaces is the updateDeviceResponseDeviceInfoIPInterfaces definition
type UpdateDeviceResponseDeviceInfoIPInterfaces struct {
	IPv4Address     string   `json:"ipv4Address,omitempty"`     //
	IPv6AddressList []string `json:"ipv6AddressList,omitempty"` //
	MacAddress      string   `json:"macAddress,omitempty"`      //
	Name            string   `json:"name,omitempty"`            //
	Status          string   `json:"status,omitempty"`          //
}

// UpdateDeviceResponseDeviceInfoIPInterfacesIPv6AddressList is the updateDeviceResponseDeviceInfoIPInterfacesIPv6AddressList definition
type UpdateDeviceResponseDeviceInfoIPInterfacesIPv6AddressList []string

// UpdateDeviceResponseDeviceInfoLocation is the updateDeviceResponseDeviceInfoLocation definition
type UpdateDeviceResponseDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}

// UpdateDeviceResponseDeviceInfoNeighborLinks is the updateDeviceResponseDeviceInfoNeighborLinks definition
type UpdateDeviceResponseDeviceInfoNeighborLinks struct {
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

// UpdateDeviceResponseDeviceInfoPnpProfileList is the updateDeviceResponseDeviceInfoPnpProfileList definition
type UpdateDeviceResponseDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                        `json:"createdBy,omitempty"`         //
	DiscoveryCreated  bool                                                          `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   UpdateDeviceResponseDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                        `json:"profileName,omitempty"`       //
	SecondaryEndpoint UpdateDeviceResponseDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}

// UpdateDeviceResponseDeviceInfoPnpProfileListPrimaryEndpoint is the updateDeviceResponseDeviceInfoPnpProfileListPrimaryEndpoint definition
type UpdateDeviceResponseDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// UpdateDeviceResponseDeviceInfoPnpProfileListSecondaryEndpoint is the updateDeviceResponseDeviceInfoPnpProfileListSecondaryEndpoint definition
type UpdateDeviceResponseDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string  `json:"certificate,omitempty"` //
	Fqdn        string  `json:"fqdn,omitempty"`        //
	IPv4Address string  `json:"ipv4Address,omitempty"` //
	IPv6Address string  `json:"ipv6Address,omitempty"` //
	Port        float64 `json:"port,omitempty"`        //
	Protocol    string  `json:"protocol,omitempty"`    //
}

// UpdateDeviceResponseDeviceInfoPreWorkflowCliOuputs is the updateDeviceResponseDeviceInfoPreWorkflowCliOuputs definition
type UpdateDeviceResponseDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}

// UpdateDeviceResponseDeviceInfoStackInfo is the updateDeviceResponseDeviceInfoStackInfo definition
type UpdateDeviceResponseDeviceInfoStackInfo struct {
	IsFullRing             bool                                                     `json:"isFullRing,omitempty"`             //
	StackMemberList        []UpdateDeviceResponseDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                   `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows bool                                                     `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       float64                                                  `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                 `json:"validLicenseLevels,omitempty"`     //
}

// UpdateDeviceResponseDeviceInfoStackInfoStackMemberList is the updateDeviceResponseDeviceInfoStackInfoStackMemberList definition
type UpdateDeviceResponseDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string  `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string  `json:"licenseLevel,omitempty"`     //
	LicenseType      string  `json:"licenseType,omitempty"`      //
	MacAddress       string  `json:"macAddress,omitempty"`       //
	Pid              string  `json:"pid,omitempty"`              //
	Priority         float64 `json:"priority,omitempty"`         //
	Role             string  `json:"role,omitempty"`             //
	SerialNumber     string  `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string  `json:"softwareVersion,omitempty"`  //
	StackNumber      float64 `json:"stackNumber,omitempty"`      //
	State            string  `json:"state,omitempty"`            //
	SudiSerialNumber string  `json:"sudiSerialNumber,omitempty"` //
}

// UpdateDeviceResponseDeviceInfoStackInfoValidLicenseLevels is the updateDeviceResponseDeviceInfoStackInfoValidLicenseLevels definition
type UpdateDeviceResponseDeviceInfoStackInfoValidLicenseLevels []string

// UpdateDeviceResponseDeviceInfoUserMicNumbers is the updateDeviceResponseDeviceInfoUserMicNumbers definition
type UpdateDeviceResponseDeviceInfoUserMicNumbers []string

// UpdateDeviceResponseDeviceInfoUserSudiSerialNos is the updateDeviceResponseDeviceInfoUserSudiSerialNos definition
type UpdateDeviceResponseDeviceInfoUserSudiSerialNos []string

// UpdateDeviceResponseRunSummaryList is the updateDeviceResponseRunSummaryList definition
type UpdateDeviceResponseRunSummaryList struct {
	Details         string                                            `json:"details,omitempty"`         //
	ErrorFlag       bool                                              `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo UpdateDeviceResponseRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       float64                                           `json:"timestamp,omitempty"`       //
}

// UpdateDeviceResponseRunSummaryListHistoryTaskInfo is the updateDeviceResponseRunSummaryListHistoryTaskInfo definition
type UpdateDeviceResponseRunSummaryListHistoryTaskInfo struct {
	AddnDetails  []UpdateDeviceResponseRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                          `json:"name,omitempty"`         //
	TimeTaken    float64                                                         `json:"timeTaken,omitempty"`    //
	Type         string                                                          `json:"type,omitempty"`         //
	WorkItemList []UpdateDeviceResponseRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}

// UpdateDeviceResponseRunSummaryListHistoryTaskInfoAddnDetails is the updateDeviceResponseRunSummaryListHistoryTaskInfoAddnDetails definition
type UpdateDeviceResponseRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// UpdateDeviceResponseRunSummaryListHistoryTaskInfoWorkItemList is the updateDeviceResponseRunSummaryListHistoryTaskInfoWorkItemList definition
type UpdateDeviceResponseRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdateDeviceResponseSystemResetWorkflow is the updateDeviceResponseSystemResetWorkflow definition
type UpdateDeviceResponseSystemResetWorkflow struct {
	TypeID         string                                         `json:"_id,omitempty"`            //
	AddToInventory bool                                           `json:"addToInventory,omitempty"` //
	AddedOn        float64                                        `json:"addedOn,omitempty"`        //
	ConfigID       string                                         `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                        `json:"currTaskIdx,omitempty"`    //
	Description    string                                         `json:"description,omitempty"`    //
	EndTime        int                                            `json:"endTime,omitempty"`        //
	ExecTime       float64                                        `json:"execTime,omitempty"`       //
	ImageID        string                                         `json:"imageId,omitempty"`        //
	InstanceType   string                                         `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                        `json:"lastupdateOn,omitempty"`   //
	Name           string                                         `json:"name,omitempty"`           //
	StartTime      int                                            `json:"startTime,omitempty"`      //
	State          string                                         `json:"state,omitempty"`          //
	Tasks          []UpdateDeviceResponseSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                         `json:"tenantId,omitempty"`       //
	Type           string                                         `json:"type,omitempty"`           //
	UseState       string                                         `json:"useState,omitempty"`       //
	Version        float64                                        `json:"version,omitempty"`        //
}

// UpdateDeviceResponseSystemResetWorkflowTasks is the updateDeviceResponseSystemResetWorkflowTasks definition
type UpdateDeviceResponseSystemResetWorkflowTasks struct {
	CurrWorkItemIDx float64                                                    `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                        `json:"endTime,omitempty"`         //
	Name            string                                                     `json:"name,omitempty"`            //
	StartTime       int                                                        `json:"startTime,omitempty"`       //
	State           string                                                     `json:"state,omitempty"`           //
	TaskSeqNo       float64                                                    `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                                    `json:"timeTaken,omitempty"`       //
	Type            string                                                     `json:"type,omitempty"`            //
	WorkItemList    []UpdateDeviceResponseSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateDeviceResponseSystemResetWorkflowTasksWorkItemList is the updateDeviceResponseSystemResetWorkflowTasksWorkItemList definition
type UpdateDeviceResponseSystemResetWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdateDeviceResponseSystemWorkflow is the updateDeviceResponseSystemWorkflow definition
type UpdateDeviceResponseSystemWorkflow struct {
	TypeID         string                                    `json:"_id,omitempty"`            //
	AddToInventory bool                                      `json:"addToInventory,omitempty"` //
	AddedOn        float64                                   `json:"addedOn,omitempty"`        //
	ConfigID       string                                    `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                                   `json:"currTaskIdx,omitempty"`    //
	Description    string                                    `json:"description,omitempty"`    //
	EndTime        int                                       `json:"endTime,omitempty"`        //
	ExecTime       float64                                   `json:"execTime,omitempty"`       //
	ImageID        string                                    `json:"imageId,omitempty"`        //
	InstanceType   string                                    `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                                   `json:"lastupdateOn,omitempty"`   //
	Name           string                                    `json:"name,omitempty"`           //
	StartTime      int                                       `json:"startTime,omitempty"`      //
	State          string                                    `json:"state,omitempty"`          //
	Tasks          []UpdateDeviceResponseSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                    `json:"tenantId,omitempty"`       //
	Type           string                                    `json:"type,omitempty"`           //
	UseState       string                                    `json:"useState,omitempty"`       //
	Version        float64                                   `json:"version,omitempty"`        //
}

// UpdateDeviceResponseSystemWorkflowTasks is the updateDeviceResponseSystemWorkflowTasks definition
type UpdateDeviceResponseSystemWorkflowTasks struct {
	CurrWorkItemIDx float64                                               `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                                   `json:"endTime,omitempty"`         //
	Name            string                                                `json:"name,omitempty"`            //
	StartTime       int                                                   `json:"startTime,omitempty"`       //
	State           string                                                `json:"state,omitempty"`           //
	TaskSeqNo       float64                                               `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                               `json:"timeTaken,omitempty"`       //
	Type            string                                                `json:"type,omitempty"`            //
	WorkItemList    []UpdateDeviceResponseSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateDeviceResponseSystemWorkflowTasksWorkItemList is the updateDeviceResponseSystemWorkflowTasksWorkItemList definition
type UpdateDeviceResponseSystemWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdateDeviceResponseWorkflow is the updateDeviceResponseWorkflow definition
type UpdateDeviceResponseWorkflow struct {
	TypeID         string                              `json:"_id,omitempty"`            //
	AddToInventory bool                                `json:"addToInventory,omitempty"` //
	AddedOn        float64                             `json:"addedOn,omitempty"`        //
	ConfigID       string                              `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                             `json:"currTaskIdx,omitempty"`    //
	Description    string                              `json:"description,omitempty"`    //
	EndTime        int                                 `json:"endTime,omitempty"`        //
	ExecTime       float64                             `json:"execTime,omitempty"`       //
	ImageID        string                              `json:"imageId,omitempty"`        //
	InstanceType   string                              `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                             `json:"lastupdateOn,omitempty"`   //
	Name           string                              `json:"name,omitempty"`           //
	StartTime      int                                 `json:"startTime,omitempty"`      //
	State          string                              `json:"state,omitempty"`          //
	Tasks          []UpdateDeviceResponseWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                              `json:"tenantId,omitempty"`       //
	Type           string                              `json:"type,omitempty"`           //
	UseState       string                              `json:"useState,omitempty"`       //
	Version        float64                             `json:"version,omitempty"`        //
}

// UpdateDeviceResponseWorkflowParameters is the updateDeviceResponseWorkflowParameters definition
type UpdateDeviceResponseWorkflowParameters struct {
	ConfigList             []UpdateDeviceResponseWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                             `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                             `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                             `json:"topOfStackSerialNumber,omitempty"` //
}

// UpdateDeviceResponseWorkflowParametersConfigList is the updateDeviceResponseWorkflowParametersConfigList definition
type UpdateDeviceResponseWorkflowParametersConfigList struct {
	ConfigID         string                                                             `json:"configId,omitempty"`         //
	ConfigParameters []UpdateDeviceResponseWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}

// UpdateDeviceResponseWorkflowParametersConfigListConfigParameters is the updateDeviceResponseWorkflowParametersConfigListConfigParameters definition
type UpdateDeviceResponseWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}

// UpdateDeviceResponseWorkflowTasks is the updateDeviceResponseWorkflowTasks definition
type UpdateDeviceResponseWorkflowTasks struct {
	CurrWorkItemIDx float64                                         `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                             `json:"endTime,omitempty"`         //
	Name            string                                          `json:"name,omitempty"`            //
	StartTime       int                                             `json:"startTime,omitempty"`       //
	State           string                                          `json:"state,omitempty"`           //
	TaskSeqNo       float64                                         `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                         `json:"timeTaken,omitempty"`       //
	Type            string                                          `json:"type,omitempty"`            //
	WorkItemList    []UpdateDeviceResponseWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateDeviceResponseWorkflowTasksWorkItemList is the updateDeviceResponseWorkflowTasksWorkItemList definition
type UpdateDeviceResponseWorkflowTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// UpdatePnPGlobalSettingsResponse is the updatePnPGlobalSettingsResponse definition
type UpdatePnPGlobalSettingsResponse struct {
	TypeID          string                                           `json:"_id,omitempty"`             //
	AAACredentials  UpdatePnPGlobalSettingsResponseAAACredentials    `json:"aaaCredentials,omitempty"`  //
	AcceptEula      bool                                             `json:"acceptEula,omitempty"`      //
	DefaultProfile  UpdatePnPGlobalSettingsResponseDefaultProfile    `json:"defaultProfile,omitempty"`  //
	ID              string                                           `json:"id,omitempty"`              //
	SavaMappingList []UpdatePnPGlobalSettingsResponseSavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    UpdatePnPGlobalSettingsResponseTaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                           `json:"tenantId,omitempty"`        //
	Version         float64                                          `json:"version,omitempty"`         //
}

// UpdatePnPGlobalSettingsResponseAAACredentials is the updatePnPGlobalSettingsResponseAAACredentials definition
type UpdatePnPGlobalSettingsResponseAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}

// UpdatePnPGlobalSettingsResponseDefaultProfile is the updatePnPGlobalSettingsResponseDefaultProfile definition
type UpdatePnPGlobalSettingsResponseDefaultProfile struct {
	Cert          string   `json:"cert,omitempty"`          //
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` //
	IPAddresses   []string `json:"ipAddresses,omitempty"`   //
	Port          float64  `json:"port,omitempty"`          //
	Proxy         bool     `json:"proxy,omitempty"`         //
}

// UpdatePnPGlobalSettingsResponseDefaultProfileFqdnAddresses is the updatePnPGlobalSettingsResponseDefaultProfileFqdnAddresses definition
type UpdatePnPGlobalSettingsResponseDefaultProfileFqdnAddresses []string

// UpdatePnPGlobalSettingsResponseDefaultProfileIPAddresses is the updatePnPGlobalSettingsResponseDefaultProfileIPAddresses definition
type UpdatePnPGlobalSettingsResponseDefaultProfileIPAddresses []string

// UpdatePnPGlobalSettingsResponseSavaMappingList is the updatePnPGlobalSettingsResponseSavaMappingList definition
type UpdatePnPGlobalSettingsResponseSavaMappingList struct {
	AutoSyncPeriod   float64                                                  `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                   `json:"ccoUser,omitempty"`          //
	Expiry           float64                                                  `json:"expiry,omitempty"`           //
	LastSync         float64                                                  `json:"lastSync,omitempty"`         //
	Profile          UpdatePnPGlobalSettingsResponseSavaMappingListProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                   `json:"smartAccountId,omitempty"`   //
	SyncResult       UpdatePnPGlobalSettingsResponseSavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                   `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                                                  `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                   `json:"syncStatus,omitempty"`       //
	TenantID         string                                                   `json:"tenantId,omitempty"`         //
	Token            string                                                   `json:"token,omitempty"`            //
	VirtualAccountID string                                                   `json:"virtualAccountId,omitempty"` //
}

// UpdatePnPGlobalSettingsResponseSavaMappingListProfile is the updatePnPGlobalSettingsResponseSavaMappingListProfile definition
type UpdatePnPGlobalSettingsResponseSavaMappingListProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// UpdatePnPGlobalSettingsResponseSavaMappingListSyncResult is the updatePnPGlobalSettingsResponseSavaMappingListSyncResult definition
type UpdatePnPGlobalSettingsResponseSavaMappingListSyncResult struct {
	SyncList []UpdatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                             `json:"syncMsg,omitempty"`  //
}

// UpdatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncList is the updatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncList definition
type UpdatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// UpdatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncListDeviceSnList is the updatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncListDeviceSnList definition
type UpdatePnPGlobalSettingsResponseSavaMappingListSyncResultSyncListDeviceSnList []string

// UpdatePnPGlobalSettingsResponseTaskTimeOuts is the updatePnPGlobalSettingsResponseTaskTimeOuts definition
type UpdatePnPGlobalSettingsResponseTaskTimeOuts struct {
	ConfigTimeOut        float64 `json:"configTimeOut,omitempty"`        //
	GeneralTimeOut       float64 `json:"generalTimeOut,omitempty"`       //
	ImageDownloadTimeOut float64 `json:"imageDownloadTimeOut,omitempty"` //
}

// UpdatePnPServerProfileResponse is the updatePnPServerProfileResponse definition
type UpdatePnPServerProfileResponse struct {
	AutoSyncPeriod   float64                                  `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                   `json:"ccoUser,omitempty"`          //
	Expiry           float64                                  `json:"expiry,omitempty"`           //
	LastSync         float64                                  `json:"lastSync,omitempty"`         //
	Profile          UpdatePnPServerProfileResponseProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                   `json:"smartAccountId,omitempty"`   //
	SyncResult       UpdatePnPServerProfileResponseSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                   `json:"syncResultStr,omitempty"`    //
	SyncStartTime    float64                                  `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                   `json:"syncStatus,omitempty"`       //
	TenantID         string                                   `json:"tenantId,omitempty"`         //
	Token            string                                   `json:"token,omitempty"`            //
	VirtualAccountID string                                   `json:"virtualAccountId,omitempty"` //
}

// UpdatePnPServerProfileResponseProfile is the updatePnPServerProfileResponseProfile definition
type UpdatePnPServerProfileResponseProfile struct {
	AddressFqdn string  `json:"addressFqdn,omitempty"` //
	AddressIPV4 string  `json:"addressIpV4,omitempty"` //
	Cert        string  `json:"cert,omitempty"`        //
	MakeDefault bool    `json:"makeDefault,omitempty"` //
	Name        string  `json:"name,omitempty"`        //
	Port        float64 `json:"port,omitempty"`        //
	ProfileID   string  `json:"profileId,omitempty"`   //
	Proxy       bool    `json:"proxy,omitempty"`       //
}

// UpdatePnPServerProfileResponseSyncResult is the updatePnPServerProfileResponseSyncResult definition
type UpdatePnPServerProfileResponseSyncResult struct {
	SyncList []UpdatePnPServerProfileResponseSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                             `json:"syncMsg,omitempty"`  //
}

// UpdatePnPServerProfileResponseSyncResultSyncList is the updatePnPServerProfileResponseSyncResultSyncList definition
type UpdatePnPServerProfileResponseSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}

// UpdatePnPServerProfileResponseSyncResultSyncListDeviceSnList is the updatePnPServerProfileResponseSyncResultSyncListDeviceSnList definition
type UpdatePnPServerProfileResponseSyncResultSyncListDeviceSnList []string

// UpdateWorkflowResponse is the updateWorkflowResponse definition
type UpdateWorkflowResponse struct {
	TypeID         string                        `json:"_id,omitempty"`            //
	AddToInventory bool                          `json:"addToInventory,omitempty"` //
	AddedOn        float64                       `json:"addedOn,omitempty"`        //
	ConfigID       string                        `json:"configId,omitempty"`       //
	CurrTaskIDx    float64                       `json:"currTaskIdx,omitempty"`    //
	Description    string                        `json:"description,omitempty"`    //
	EndTime        int                           `json:"endTime,omitempty"`        //
	ExecTime       float64                       `json:"execTime,omitempty"`       //
	ImageID        string                        `json:"imageId,omitempty"`        //
	InstanceType   string                        `json:"instanceType,omitempty"`   //
	LastupdateOn   float64                       `json:"lastupdateOn,omitempty"`   //
	Name           string                        `json:"name,omitempty"`           //
	StartTime      int                           `json:"startTime,omitempty"`      //
	State          string                        `json:"state,omitempty"`          //
	Tasks          []UpdateWorkflowResponseTasks `json:"tasks,omitempty"`          //
	TenantID       string                        `json:"tenantId,omitempty"`       //
	Type           string                        `json:"type,omitempty"`           //
	UseState       string                        `json:"useState,omitempty"`       //
	Version        float64                       `json:"version,omitempty"`        //
}

// UpdateWorkflowResponseTasks is the updateWorkflowResponseTasks definition
type UpdateWorkflowResponseTasks struct {
	CurrWorkItemIDx float64                                   `json:"currWorkItemIdx,omitempty"` //
	EndTime         int                                       `json:"endTime,omitempty"`         //
	Name            string                                    `json:"name,omitempty"`            //
	StartTime       int                                       `json:"startTime,omitempty"`       //
	State           string                                    `json:"state,omitempty"`           //
	TaskSeqNo       float64                                   `json:"taskSeqNo,omitempty"`       //
	TimeTaken       float64                                   `json:"timeTaken,omitempty"`       //
	Type            string                                    `json:"type,omitempty"`            //
	WorkItemList    []UpdateWorkflowResponseTasksWorkItemList `json:"workItemList,omitempty"`    //
}

// UpdateWorkflowResponseTasksWorkItemList is the updateWorkflowResponseTasksWorkItemList definition
type UpdateWorkflowResponseTasksWorkItemList struct {
	Command   string  `json:"command,omitempty"`   //
	EndTime   int     `json:"endTime,omitempty"`   //
	OutputStr string  `json:"outputStr,omitempty"` //
	StartTime int     `json:"startTime,omitempty"` //
	State     string  `json:"state,omitempty"`     //
	TimeTaken float64 `json:"timeTaken,omitempty"` //
}

// AddAWorkflow addAWorkflow
/* Adds a PnP Workflow along with the relevant tasks in the workflow into the PnP database
 */
func (s *DeviceOnboardingPnPService) AddAWorkflow(addAWorkflowRequest *AddAWorkflowRequest) (*AddAWorkflowResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

	response, err := RestyClient.R().
		SetBody(addAWorkflowRequest).
		SetResult(&AddAWorkflowResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation addAWorkflow")
	}

	result := response.Result().(*AddAWorkflowResponse)
	return result, response, err
}

// AddDeviceToPnpDatabase addDeviceToPnpDatabase
/* Adds a device to the PnP database.
 */
func (s *DeviceOnboardingPnPService) AddDeviceToPnpDatabase(addDeviceToPnpDatabaseRequest *AddDeviceToPnpDatabaseRequest) (*AddDeviceToPnpDatabaseResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device"

	response, err := RestyClient.R().
		SetBody(addDeviceToPnpDatabaseRequest).
		SetResult(&AddDeviceToPnpDatabaseResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation addDeviceToPnpDatabase")
	}

	result := response.Result().(*AddDeviceToPnpDatabaseResponse)
	return result, response, err
}

// AddVirtualAccount addVirtualAccount
/* Registers a Smart Account, Virtual Account and the relevant server profile info with the PnP System & database. The devices present in the registered virtual account are synced with the PnP database as well. The response payload returns the new profile
 */
func (s *DeviceOnboardingPnPService) AddVirtualAccount(addVirtualAccountRequest *AddVirtualAccountRequest) (*AddVirtualAccountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

	response, err := RestyClient.R().
		SetBody(addVirtualAccountRequest).
		SetResult(&AddVirtualAccountResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation addVirtualAccount")
	}

	result := response.Result().(*AddVirtualAccountResponse)
	return result, response, err
}

// ClaimADeviceToASite claimADeviceToASite
/* Claim a device based on DNA-C Site based design process. Different parameters are required for different device platforms.
 */
func (s *DeviceOnboardingPnPService) ClaimADeviceToASite(claimADeviceToASiteRequest *ClaimADeviceToASiteRequest) (*ClaimADeviceToASiteResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/site-claim"

	response, err := RestyClient.R().
		SetBody(claimADeviceToASiteRequest).
		SetResult(&ClaimADeviceToASiteResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation claimADeviceToASite")
	}

	result := response.Result().(*ClaimADeviceToASiteResponse)
	return result, response, err
}

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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation claimDevice")
	}

	result := response.Result().(*ClaimDeviceResponse)
	return result, response, err
}

// DeleteDeviceByIDFromPnP deleteDeviceByIdFromPnP
/* Deletes specified device from PnP database
@param id id
*/
func (s *DeviceOnboardingPnPService) DeleteDeviceByIDFromPnP(id string) (*DeleteDeviceByIDFromPnPResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteDeviceByIDFromPnPResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteDeviceByIdFromPnP")
	}

	result := response.Result().(*DeleteDeviceByIDFromPnPResponse)
	return result, response, err
}

// DeleteWorkflowByID deleteWorkflowById
/* Deletes a workflow specified by id
@param id id
*/
func (s *DeviceOnboardingPnPService) DeleteWorkflowByID(id string) (*DeleteWorkflowByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteWorkflowByIDResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteWorkflowById")
	}

	result := response.Result().(*DeleteWorkflowByIDResponse)
	return result, response, err
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
func (s *DeviceOnboardingPnPService) DeregisterVirtualAccount(deregisterVirtualAccountQueryParams *DeregisterVirtualAccountQueryParams) (*DeregisterVirtualAccountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings/vacct"

	queryString, _ := query.Values(deregisterVirtualAccountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeregisterVirtualAccountResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deregisterVirtualAccount")
	}

	result := response.Result().(*DeregisterVirtualAccountResponse)
	return result, response, err
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getDeviceById")
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
/* Returns history for a specific device. Serial number is a required parameter
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getDeviceHistory")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getPnPGlobalSettings")
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
	Pid              []string `url:"pid,omitempty"`              // Device ProductId
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
@param pid Device ProductId
@param source Device Source
@param projectID Device Project Id
@param workflowID Device Workflow Id
@param projectName Device Project Name
@param workflowName Device Workflow Name
@param smartAccountID Device Smart Account
@param virtualAccountID Device Virtual Account
@param lastContact Device Has Contacted lastContact > 0
*/
func (s *DeviceOnboardingPnPService) GetPnpDeviceCount(getPnpDeviceCountQueryParams *GetPnpDeviceCountQueryParams) (*GetPnpDeviceCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/count"

	queryString, _ := query.Values(getPnpDeviceCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetPnpDeviceCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getPnpDeviceCount")
	}

	result := response.Result().(*GetPnpDeviceCountResponse)
	return result, response, err
}

// GetPnpDeviceListQueryParams defines the query parameters for this request
type GetPnpDeviceListQueryParams struct {
	Limit            int      `url:"limit,omitempty"`            // Limits number of results
	Offset           int      `url:"offset,omitempty"`           // Index of first result
	Sort             []string `url:"sort,omitempty"`             // Comma seperated list of fields to sort on
	SortOrder        string   `url:"sortOrder,omitempty"`        // Sort Order Ascending (asc) or Descending (des)
	SerialNumber     []string `url:"serialNumber,omitempty"`     // Device Serial Number
	State            []string `url:"state,omitempty"`            // Device State
	OnbState         []string `url:"onbState,omitempty"`         // Device Onboarding State
	CmState          []string `url:"cmState,omitempty"`          // Device Connection Manager State
	Name             []string `url:"name,omitempty"`             // Device Name
	Pid              []string `url:"pid,omitempty"`              // Device ProductId
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
@param limit Limits number of results
@param offset Index of first result
@param sort Comma seperated list of fields to sort on
@param sortOrder Sort Order Ascending (asc) or Descending (des)
@param serialNumber Device Serial Number
@param state Device State
@param onbState Device Onboarding State
@param cmState Device Connection Manager State
@param name Device Name
@param pid Device ProductId
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
func (s *DeviceOnboardingPnPService) GetPnpDeviceList(getPnpDeviceListQueryParams *GetPnpDeviceListQueryParams) (*GetPnpDeviceListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device"

	queryString, _ := query.Values(getPnpDeviceListQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetPnpDeviceListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getPnpDeviceList")
	}

	result := response.Result().(*GetPnpDeviceListResponse)
	return result, response, err
}

// GetSmartAccountList getSmartAccountList
/* Returns the list of Smart Account domains
 */
func (s *DeviceOnboardingPnPService) GetSmartAccountList() (*GetSmartAccountListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct"

	response, err := RestyClient.R().
		SetResult(&GetSmartAccountListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getSmartAccountList")
	}

	result := response.Result().(*GetSmartAccountListResponse)
	return result, response, err
}

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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getSyncResultForVirtualAccount")
	}

	result := response.Result().(*GetSyncResultForVirtualAccountResponse)
	return result, response, err
}

// GetVirtualAccountList getVirtualAccountList
/* Returns list of virtual accounts associated with the specified smart account
@param domain Smart Account Domain
*/
func (s *DeviceOnboardingPnPService) GetVirtualAccountList(domain string) (*GetVirtualAccountListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct/{domain}/vacct"
	path = strings.Replace(path, "{"+"domain"+"}", fmt.Sprintf("%v", domain), -1)

	response, err := RestyClient.R().
		SetResult(&GetVirtualAccountListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getVirtualAccountList")
	}

	result := response.Result().(*GetVirtualAccountListResponse)
	return result, response, err
}

// GetWorkflowByID getWorkflowById
/* Returns a workflow specified by id
@param id id
*/
func (s *DeviceOnboardingPnPService) GetWorkflowByID(id string) (*GetWorkflowByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetWorkflowByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getWorkflowById")
	}

	result := response.Result().(*GetWorkflowByIDResponse)
	return result, response, err
}

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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getWorkflowCount")
	}

	result := response.Result().(*GetWorkflowCountResponse)
	return result, response, err
}

// GetWorkflowsQueryParams defines the query parameters for this request
type GetWorkflowsQueryParams struct {
	Limit     int      `url:"limit,omitempty"`     // Limits number of results
	Offset    int      `url:"offset,omitempty"`    // Index of first result
	Sort      []string `url:"sort,omitempty"`      // Comma seperated lost of fields to sort on
	SortOrder string   `url:"sortOrder,omitempty"` // Sort Order Ascending (asc) or Descending (des)
	Type      []string `url:"type,omitempty"`      // Workflow Type
	Name      []string `url:"name,omitempty"`      // Workflow Name
}

// GetWorkflows getWorkflows
/* Returns the list of workflows based on filter criteria. If a limit is not specified, it will default to return 50 workflows. Pagination and sorting are also supported by this endpoint
@param limit Limits number of results
@param offset Index of first result
@param sort Comma seperated lost of fields to sort on
@param sortOrder Sort Order Ascending (asc) or Descending (des)
@param type Workflow Type
@param name Workflow Name
*/
func (s *DeviceOnboardingPnPService) GetWorkflows(getWorkflowsQueryParams *GetWorkflowsQueryParams) (*GetWorkflowsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

	queryString, _ := query.Values(getWorkflowsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetWorkflowsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getWorkflows")
	}

	result := response.Result().(*GetWorkflowsResponse)
	return result, response, err
}

// ImportDevicesInBulk importDevicesInBulk
/* Add devices to PnP in bulk
 */
func (s *DeviceOnboardingPnPService) ImportDevicesInBulk(importDevicesInBulkRequest *ImportDevicesInBulkRequest) (*ImportDevicesInBulkResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/import"

	response, err := RestyClient.R().
		SetBody(importDevicesInBulkRequest).
		SetResult(&ImportDevicesInBulkResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation importDevicesInBulk")
	}

	result := response.Result().(*ImportDevicesInBulkResponse)
	return result, response, err
}

// PreviewConfig previewConfig
/* Triggers a preview for site-based Day 0 Configuration
 */
func (s *DeviceOnboardingPnPService) PreviewConfig(previewConfigRequest *PreviewConfigRequest) (*PreviewConfigResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/site-config-preview"

	response, err := RestyClient.R().
		SetBody(previewConfigRequest).
		SetResult(&PreviewConfigResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation previewConfig")
	}

	result := response.Result().(*PreviewConfigResponse)
	return result, response, err
}

// ResetDevice resetDevice
/* Recovers a device from a Workflow Execution Error state
 */
func (s *DeviceOnboardingPnPService) ResetDevice(resetDeviceRequest *ResetDeviceRequest) (*ResetDeviceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/reset"

	response, err := RestyClient.R().
		SetBody(resetDeviceRequest).
		SetResult(&ResetDeviceResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation resetDevice")
	}

	result := response.Result().(*ResetDeviceResponse)
	return result, response, err
}

// SyncVirtualAccountDevices syncVirtualAccountDevices
/* Synchronizes the device info from the given smart account & virtual account with the PnP database. The response payload returns a list of synced devices
 */
func (s *DeviceOnboardingPnPService) SyncVirtualAccountDevices(syncVirtualAccountDevicesRequest *SyncVirtualAccountDevicesRequest) (*SyncVirtualAccountDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/vacct-sync"

	response, err := RestyClient.R().
		SetBody(syncVirtualAccountDevicesRequest).
		SetResult(&SyncVirtualAccountDevicesResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation syncVirtualAccountDevices")
	}

	result := response.Result().(*SyncVirtualAccountDevicesResponse)
	return result, response, err
}

// UnClaimDevice unClaimDevice
/* UnClaims one of more devices with specified workflow
 */
func (s *DeviceOnboardingPnPService) UnClaimDevice(unClaimDeviceRequest *UnClaimDeviceRequest) (*UnClaimDeviceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/unclaim"

	response, err := RestyClient.R().
		SetBody(unClaimDeviceRequest).
		SetResult(&UnClaimDeviceResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation unClaimDevice")
	}

	result := response.Result().(*UnClaimDeviceResponse)
	return result, response, err
}

// UpdateDevice updateDevice
/* Updates device details specified by device id in PnP database
@param id id
*/
func (s *DeviceOnboardingPnPService) UpdateDevice(id string, updateDeviceRequest *UpdateDeviceRequest) (*UpdateDeviceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetBody(updateDeviceRequest).
		SetResult(&UpdateDeviceResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateDevice")
	}

	result := response.Result().(*UpdateDeviceResponse)
	return result, response, err
}

// UpdatePnPGlobalSettings updatePnPGlobalSettings
/* Updates the user's list of global PnP settings
 */
func (s *DeviceOnboardingPnPService) UpdatePnPGlobalSettings(updatePnPGlobalSettingsRequest *UpdatePnPGlobalSettingsRequest) (*UpdatePnPGlobalSettingsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings"

	response, err := RestyClient.R().
		SetBody(updatePnPGlobalSettingsRequest).
		SetResult(&UpdatePnPGlobalSettingsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updatePnPGlobalSettings")
	}

	result := response.Result().(*UpdatePnPGlobalSettingsResponse)
	return result, response, err
}

// UpdatePnPServerProfile updatePnPServerProfile
/* Updates the PnP Server profile in a registered Virtual Account in the PnP database. The response payload returns the updated smart & virtual account info
 */
func (s *DeviceOnboardingPnPService) UpdatePnPServerProfile(updatePnPServerProfileRequest *UpdatePnPServerProfileRequest) (*UpdatePnPServerProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

	response, err := RestyClient.R().
		SetBody(updatePnPServerProfileRequest).
		SetResult(&UpdatePnPServerProfileResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updatePnPServerProfile")
	}

	result := response.Result().(*UpdatePnPServerProfileResponse)
	return result, response, err
}

// UpdateWorkflow updateWorkflow
/* Updates an existing workflow
@param id id
*/
func (s *DeviceOnboardingPnPService) UpdateWorkflow(id string, updateWorkflowRequest *UpdateWorkflowRequest) (*UpdateWorkflowResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetBody(updateWorkflowRequest).
		SetResult(&UpdateWorkflowResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateWorkflow")
	}

	result := response.Result().(*UpdateWorkflowResponse)
	return result, response, err
}
