package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DeviceOnboardingPnpService service

type GetDeviceList2QueryParams struct {
	Limit            int      `url:"limit,omitempty"`            //Limits number of results
	Offset           int      `url:"offset,omitempty"`           //Index of first result
	Sort             []string `url:"sort,omitempty"`             //Comma seperated list of fields to sort on
	SortOrder        string   `url:"sortOrder,omitempty"`        //Sort Order Ascending (asc) or Descending (des)
	SerialNumber     []string `url:"serialNumber,omitempty"`     //Device Serial Number
	State            []string `url:"state,omitempty"`            //Device State
	OnbState         []string `url:"onbState,omitempty"`         //Device Onboarding State
	CmState          []string `url:"cmState,omitempty"`          //Device Connection Manager State
	Name             []string `url:"name,omitempty"`             //Device Name
	Pid              []string `url:"pid,omitempty"`              //Device ProductId
	Source           []string `url:"source,omitempty"`           //Device Source
	ProjectID        []string `url:"projectId,omitempty"`        //Device Project Id
	WorkflowID       []string `url:"workflowId,omitempty"`       //Device Workflow Id
	ProjectName      []string `url:"projectName,omitempty"`      //Device Project Name
	WorkflowName     []string `url:"workflowName,omitempty"`     //Device Workflow Name
	SmartAccountID   []string `url:"smartAccountId,omitempty"`   //Device Smart Account
	VirtualAccountID []string `url:"virtualAccountId,omitempty"` //Device Virtual Account
	LastContact      bool     `url:"lastContact,omitempty"`      //Device Has Contacted lastContact > 0
	MacAddress       string   `url:"macAddress,omitempty"`       //Device Mac Address
	Hostname         string   `url:"hostname,omitempty"`         //Device Hostname
	SiteName         string   `url:"siteName,omitempty"`         //Device Site Name
}
type GetDeviceCountQueryParams struct {
	SerialNumber     []string `url:"serialNumber,omitempty"`     //Device Serial Number
	State            []string `url:"state,omitempty"`            //Device State
	OnbState         []string `url:"onbState,omitempty"`         //Device Onboarding State
	CmState          []string `url:"cmState,omitempty"`          //Device Connection Manager State
	Name             []string `url:"name,omitempty"`             //Device Name
	Pid              []string `url:"pid,omitempty"`              //Device ProductId
	Source           []string `url:"source,omitempty"`           //Device Source
	ProjectID        []string `url:"projectId,omitempty"`        //Device Project Id
	WorkflowID       []string `url:"workflowId,omitempty"`       //Device Workflow Id
	ProjectName      []string `url:"projectName,omitempty"`      //Device Project Name
	WorkflowName     []string `url:"workflowName,omitempty"`     //Device Workflow Name
	SmartAccountID   []string `url:"smartAccountId,omitempty"`   //Device Smart Account
	VirtualAccountID []string `url:"virtualAccountId,omitempty"` //Device Virtual Account
	LastContact      bool     `url:"lastContact,omitempty"`      //Device Has Contacted lastContact > 0
}
type GetDeviceHistoryQueryParams struct {
	SerialNumber string   `url:"serialNumber,omitempty"` //Device Serial Number
	Sort         []string `url:"sort,omitempty"`         //Comma seperated list of fields to sort on
	SortOrder    string   `url:"sortOrder,omitempty"`    //Sort Order Ascending (asc) or Descending (des)
}
type DeregisterVirtualAccountQueryParams struct {
	Domain string `url:"domain,omitempty"` //Smart Account Domain
	Name   string `url:"name,omitempty"`   //Virtual Account Name
}
type GetWorkflowsQueryParams struct {
	Limit     int      `url:"limit,omitempty"`     //Limits number of results
	Offset    int      `url:"offset,omitempty"`    //Index of first result
	Sort      []string `url:"sort,omitempty"`      //Comma seperated lost of fields to sort on
	SortOrder string   `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (des)
	Type      []string `url:"type,omitempty"`      //Workflow Type
	Name      []string `url:"name,omitempty"`      //Workflow Name
}
type GetWorkflowCountQueryParams struct {
	Name []string `url:"name,omitempty"` //Workflow Name
}

type ResponseDeviceOnboardingPnpAddDevice struct {
	TypeID               string                                                    `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpAddDeviceDeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpAddDeviceSystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpAddDeviceSystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpAddDeviceWorkflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpAddDeviceRunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpAddDeviceWorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpAddDeviceDayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpAddDeviceDayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                  `json:"version,omitempty"`              // Version
	TenantID             string                                                    `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfo struct {
	Source                    string                                                                `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                 `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                `json:"mode,omitempty"`                      // Mode
	State                     string                                                                `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                `json:"description,omitempty"`               // Description
	OnbState                  string                                                                `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                              `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                              `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                              `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                              `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                              `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                              `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                              `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                 `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                              `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                              `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                              `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                              `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                 `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                 `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList struct {
	ProfileName       string                                                                         `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                          `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                         `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                `json:"port,omitempty"`        // Port
	Protocol    string                                                                                  `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                  `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                  `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                  `json:"port,omitempty"`        // Port
	Protocol    string                                                                                    `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                    `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                    `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces struct {
	Status          string                                                                       `json:"status,omitempty"`          // Status
	MacAddress      string                                                                       `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                       `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                     `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                     `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                    `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                  `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                  `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpAddDeviceDeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpAddDeviceSystemResetWorkflow struct {
	TypeID         string                                                          `json:"_id,omitempty"`            // Id
	State          string                                                          `json:"state,omitempty"`          // State
	Type           string                                                          `json:"type,omitempty"`           // Type
	Description    string                                                          `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                        `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                          `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                        `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                        `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                           `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                          `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                        `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                        `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                        `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                          `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                          `json:"configId,omitempty"`       // Config Id
	Name           string                                                          `json:"name,omitempty"`           // Name
	Version        *float64                                                        `json:"version,omitempty"`        // Version
	TenantID       string                                                          `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks struct {
	State           string                                                                      `json:"state,omitempty"`           // State
	Type            string                                                                      `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                    `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                    `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                    `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                    `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                    `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                      `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceSystemWorkflow struct {
	TypeID         string                                                     `json:"_id,omitempty"`            // Id
	State          string                                                     `json:"state,omitempty"`          // State
	Type           string                                                     `json:"type,omitempty"`           // Type
	Description    string                                                     `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                   `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                     `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                   `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                   `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddDeviceSystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                      `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                     `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                   `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                   `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                   `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                     `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                     `json:"configId,omitempty"`       // Config Id
	Name           string                                                     `json:"name,omitempty"`           // Name
	Version        *float64                                                   `json:"version,omitempty"`        // Version
	TenantID       string                                                     `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceSystemWorkflowTasks struct {
	State           string                                                                 `json:"state,omitempty"`           // State
	Type            string                                                                 `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                               `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                               `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                               `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                               `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                               `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                 `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceWorkflow struct {
	TypeID         string                                               `json:"_id,omitempty"`            // Id
	State          string                                               `json:"state,omitempty"`          // State
	Type           string                                               `json:"type,omitempty"`           // Type
	Description    string                                               `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                             `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                               `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                             `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                             `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddDeviceWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                               `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                             `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                             `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                             `json:"startTime,omitempty"`      // Start Time
	UseState       string                                               `json:"useState,omitempty"`       // Use State
	ConfigID       string                                               `json:"configId,omitempty"`       // Config Id
	Name           string                                               `json:"name,omitempty"`           // Name
	Version        *float64                                             `json:"version,omitempty"`        // Version
	TenantID       string                                               `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddDeviceWorkflowTasks struct {
	State           string                                                           `json:"state,omitempty"`           // State
	Type            string                                                           `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                         `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                         `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                         `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                         `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                         `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                           `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceRunSummaryList struct {
	Details         string                                                             `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                              `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                           `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo struct {
	Type         string                                                                           `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                         `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                           `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpAddDeviceWorkflowParameters struct {
	TopOfStackSerialNumber string                                                              `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                              `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                              `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                              `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpAddDeviceDayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpAddDeviceDayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpGetDeviceList2 []ResponseItemDeviceOnboardingPnpGetDeviceList2 // Array of ResponseDeviceOnboardingPnpGetDeviceList2
type ResponseItemDeviceOnboardingPnpGetDeviceList2 struct {
	DeviceInfo           *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseItemDeviceOnboardingPnpGetDeviceList2Workflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseItemDeviceOnboardingPnpGetDeviceList2DayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseItemDeviceOnboardingPnpGetDeviceList2DayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                           `json:"version,omitempty"`              // Version
	TenantID             string                                                             `json:"tenantId,omitempty"`             // Tenant Id
	ID                   string                                                             `json:"id,omitempty"`                   // Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfo struct {
	Source                    string                                                                         `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                         `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                          `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                         `json:"mode,omitempty"`                      // Mode
	State                     string                                                                         `json:"state,omitempty"`                     // State
	Location                  *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                         `json:"description,omitempty"`               // Description
	OnbState                  string                                                                         `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                         `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                         `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                       `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                       `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                         `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                       `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                       `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                         `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                         `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                       `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                       `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                         `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                         `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                         `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                         `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                         `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                         `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                         `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                         `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                       `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                         `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                         `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                          `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                       `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                         `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                       `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                       `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                       `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                          `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                         `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                         `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                          `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                         `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                         `json:"name,omitempty"`                      // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileList struct {
	ProfileName       string                                                                                  `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                   `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                                  `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                         `json:"port,omitempty"`        // Port
	Protocol    string                                                                                           `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                           `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                           `json:"certificate,omitempty"` // Certificate
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                           `json:"port,omitempty"`        // Port
	Protocol    string                                                                                             `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                             `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                             `json:"certificate,omitempty"` // Certificate
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfaces struct {
	Status          string                                                                                `json:"status,omitempty"`          // Status
	MacAddress      string                                                                                `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                                `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfacesIPv4Address interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                              `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                              `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                             `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                           `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                           `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DeviceInfoTags interface{}
type ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflow struct {
	TypeID         string                                                                   `json:"_id,omitempty"`            // Id
	State          string                                                                   `json:"state,omitempty"`          // State
	Type           string                                                                   `json:"type,omitempty"`           // Type
	Description    string                                                                   `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                 `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                   `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                 `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                 `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                    `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                   `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                 `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                 `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                 `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                   `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                   `json:"configId,omitempty"`       // Config Id
	Name           string                                                                   `json:"name,omitempty"`           // Name
	Version        *float64                                                                 `json:"version,omitempty"`        // Version
	TenantID       string                                                                   `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflowTasks struct {
	State           string                                                                               `json:"state,omitempty"`           // State
	Type            string                                                                               `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                             `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                             `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                             `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                             `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                             `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                               `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2SystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflow struct {
	TypeID         string                                                              `json:"_id,omitempty"`            // Id
	State          string                                                              `json:"state,omitempty"`          // State
	Type           string                                                              `json:"type,omitempty"`           // Type
	Description    string                                                              `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                            `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                              `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                            `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                            `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                               `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                              `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                            `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                            `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                            `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                              `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                              `json:"configId,omitempty"`       // Config Id
	Name           string                                                              `json:"name,omitempty"`           // Name
	Version        *float64                                                            `json:"version,omitempty"`        // Version
	TenantID       string                                                              `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflowTasks struct {
	State           string                                                                          `json:"state,omitempty"`           // State
	Type            string                                                                          `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                        `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                        `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                        `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                        `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                        `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                          `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2SystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2Workflow struct {
	TypeID         string                                                        `json:"_id,omitempty"`            // Id
	State          string                                                        `json:"state,omitempty"`          // State
	Type           string                                                        `json:"type,omitempty"`           // Type
	Description    string                                                        `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                      `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                        `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                      `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                      `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                         `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                        `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                      `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                      `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                      `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                        `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                        `json:"configId,omitempty"`       // Config Id
	Name           string                                                        `json:"name,omitempty"`           // Name
	Version        *float64                                                      `json:"version,omitempty"`        // Version
	TenantID       string                                                        `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowTasks struct {
	State           string                                                                    `json:"state,omitempty"`           // State
	Type            string                                                                    `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                  `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                  `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                  `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                  `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                  `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                    `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryList struct {
	Details         string                                                                      `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                       `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                    `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                    `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                                  `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                    `json:"name,omitempty"`         // Name
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2RunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParameters struct {
	TopOfStackSerialNumber string                                                                       `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                       `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                       `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                       `json:"configId,omitempty"`         // Config Id
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2WorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseItemDeviceOnboardingPnpGetDeviceList2DayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpClaimDevice struct {
	JSONArrayResponse *[]ResponseDeviceOnboardingPnpClaimDeviceJSONArrayResponse `json:"jsonArrayResponse,omitempty"` // Json Array Response
	JSONResponse      *ResponseDeviceOnboardingPnpClaimDeviceJSONResponse        `json:"jsonResponse,omitempty"`      // Json Response
	Message           string                                                     `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                   `json:"statusCode,omitempty"`        // Status Code
}
type ResponseDeviceOnboardingPnpClaimDeviceJSONArrayResponse interface{}
type ResponseDeviceOnboardingPnpClaimDeviceJSONResponse interface{}
type ResponseDeviceOnboardingPnpGetDeviceCount struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseDeviceOnboardingPnpGetDeviceHistory struct {
	Response   *[]ResponseDeviceOnboardingPnpGetDeviceHistoryResponse `json:"response,omitempty"`   //
	StatusCode *float64                                               `json:"statusCode,omitempty"` // Status Code
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryResponse struct {
	Timestamp       *float64                                                            `json:"timestamp,omitempty"`       // Timestamp
	Details         string                                                              `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpGetDeviceHistoryResponseHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                               `json:"errorFlag,omitempty"`       // Error Flag
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryResponseHistoryTaskInfo struct {
	Name         string                                                                            `json:"name,omitempty"`         // Name
	Type         string                                                                            `json:"type,omitempty"`         // Type
	TimeTaken    *float64                                                                          `json:"timeTaken,omitempty"`    // Time Taken
	WorkItemList *[]ResponseDeviceOnboardingPnpGetDeviceHistoryResponseHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	AddnDetails  *[]ResponseDeviceOnboardingPnpGetDeviceHistoryResponseHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryResponseHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
}
type ResponseDeviceOnboardingPnpGetDeviceHistoryResponseHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulk struct {
	SuccessList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessList `json:"successList,omitempty"` //
	FailureList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkFailureList `json:"failureList,omitempty"` //
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessList struct {
	TypeID               string                                                                         `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                                       `json:"version,omitempty"`              // Version
	TenantID             string                                                                         `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfo struct {
	Source                    string                                                                                     `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                                     `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                                      `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                                     `json:"mode,omitempty"`                      // Mode
	State                     string                                                                                     `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                                     `json:"description,omitempty"`               // Description
	OnbState                  string                                                                                     `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                                     `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                                     `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                                   `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                                   `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                                     `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                                   `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                                   `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                                     `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                                     `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                                   `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                                   `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                                     `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                                     `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                                     `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                                     `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                                     `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                                     `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                                     `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                                     `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                                   `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                                     `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                                     `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                                      `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                                   `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                                     `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                                   `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                                   `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                                   `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                                      `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                                     `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                                     `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                                      `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                                     `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                                     `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileList struct {
	ProfileName       string                                                                                              `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                               `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                                              `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                                     `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                       `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                       `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                       `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                                       `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                         `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                         `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                         `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfaces struct {
	Status          string                                                                                            `json:"status,omitempty"`          // Status
	MacAddress      string                                                                                            `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                                            `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                                          `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                                          `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                         `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                                       `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                                       `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflow struct {
	TypeID         string                                                                               `json:"_id,omitempty"`            // Id
	State          string                                                                               `json:"state,omitempty"`          // State
	Type           string                                                                               `json:"type,omitempty"`           // Type
	Description    string                                                                               `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                             `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                               `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                             `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                             `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                                `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                               `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                             `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                             `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                             `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                               `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                               `json:"configId,omitempty"`       // Config Id
	Name           string                                                                               `json:"name,omitempty"`           // Name
	Version        *float64                                                                             `json:"version,omitempty"`        // Version
	TenantID       string                                                                               `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasks struct {
	State           string                                                                                           `json:"state,omitempty"`           // State
	Type            string                                                                                           `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                         `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                         `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                         `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                         `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                         `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                           `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflow struct {
	TypeID         string                                                                          `json:"_id,omitempty"`            // Id
	State          string                                                                          `json:"state,omitempty"`          // State
	Type           string                                                                          `json:"type,omitempty"`           // Type
	Description    string                                                                          `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                        `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                          `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                        `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                        `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                           `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                          `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                        `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                        `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                        `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                          `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                          `json:"configId,omitempty"`       // Config Id
	Name           string                                                                          `json:"name,omitempty"`           // Name
	Version        *float64                                                                        `json:"version,omitempty"`        // Version
	TenantID       string                                                                          `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasks struct {
	State           string                                                                                      `json:"state,omitempty"`           // State
	Type            string                                                                                      `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                    `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                    `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                    `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                    `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                    `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                      `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflow struct {
	TypeID         string                                                                    `json:"_id,omitempty"`            // Id
	State          string                                                                    `json:"state,omitempty"`          // State
	Type           string                                                                    `json:"type,omitempty"`           // Type
	Description    string                                                                    `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                  `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                    `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                  `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                  `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                     `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                    `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                  `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                  `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                  `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                    `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                    `json:"configId,omitempty"`       // Config Id
	Name           string                                                                    `json:"name,omitempty"`           // Name
	Version        *float64                                                                  `json:"version,omitempty"`        // Version
	TenantID       string                                                                    `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasks struct {
	State           string                                                                                `json:"state,omitempty"`           // State
	Type            string                                                                                `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                              `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                              `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                              `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                              `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                              `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryList struct {
	Details         string                                                                                  `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                                   `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                                `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                                `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                                              `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                                `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParameters struct {
	TopOfStackSerialNumber string                                                                                   `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                                   `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                                   `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                                   `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpImportDevicesInBulkFailureList struct {
	Index     *float64 `json:"index,omitempty"`     // Index
	SerialNum string   `json:"serialNum,omitempty"` // Serial Num
	ID        string   `json:"id,omitempty"`        // Id
	Msg       string   `json:"msg,omitempty"`       // Msg
}
type ResponseDeviceOnboardingPnpResetDevice struct {
	JSONArrayResponse *[]ResponseDeviceOnboardingPnpResetDeviceJSONArrayResponse `json:"jsonArrayResponse,omitempty"` // Json Array Response
	JSONResponse      *ResponseDeviceOnboardingPnpResetDeviceJSONResponse        `json:"jsonResponse,omitempty"`      // Json Response
	Message           string                                                     `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                   `json:"statusCode,omitempty"`        // Status Code
}
type ResponseDeviceOnboardingPnpResetDeviceJSONArrayResponse interface{}
type ResponseDeviceOnboardingPnpResetDeviceJSONResponse interface{}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccount struct {
	VirtualAccountID string                                                               `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                             `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                               `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountProfile    `json:"profile,omitempty"`          //
	CcoUser          string                                                               `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountSyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                               `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                             `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                             `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                               `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                               `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                             `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                               `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountProfile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                         `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccountSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpClaimADeviceToASite struct {
	Response string `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDeviceOnboardingPnpPreviewConfig struct {
	Response *ResponseDeviceOnboardingPnpPreviewConfigResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  //
}
type ResponseDeviceOnboardingPnpPreviewConfigResponse struct {
	Complete      *bool  `json:"complete,omitempty"`      //
	Config        string `json:"config,omitempty"`        //
	Error         *bool  `json:"error,omitempty"`         //
	ErrorMessage  string `json:"errorMessage,omitempty"`  //
	ExpiredTime   *int   `json:"expiredTime,omitempty"`   //
	RfProfile     string `json:"rfProfile,omitempty"`     //
	SensorProfile string `json:"sensorProfile,omitempty"` //
	SiteID        string `json:"siteId,omitempty"`        //
	StartTime     *int   `json:"startTime,omitempty"`     //
	TaskID        string `json:"taskId,omitempty"`        //
}
type ResponseDeviceOnboardingPnpUnClaimDevice struct {
	JSONArrayResponse *[]ResponseDeviceOnboardingPnpUnClaimDeviceJSONArrayResponse `json:"jsonArrayResponse,omitempty"` // Json Array Response
	JSONResponse      *ResponseDeviceOnboardingPnpUnClaimDeviceJSONResponse        `json:"jsonResponse,omitempty"`      // Json Response
	Message           string                                                       `json:"message,omitempty"`           // Message
	StatusCode        *float64                                                     `json:"statusCode,omitempty"`        // Status Code
}
type ResponseDeviceOnboardingPnpUnClaimDeviceJSONArrayResponse interface{}
type ResponseDeviceOnboardingPnpUnClaimDeviceJSONResponse interface{}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevices struct {
	VirtualAccountID string                                                          `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                        `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                          `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesProfile    `json:"profile,omitempty"`          //
	CcoUser          string                                                          `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                          `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                        `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                        `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                          `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                          `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                        `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                          `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesProfile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                    `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpUpdateDevice struct {
	TypeID               string                                                       `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpUpdateDeviceSystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpUpdateDeviceWorkflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpUpdateDeviceWorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpUpdateDeviceDayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpUpdateDeviceDayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                     `json:"version,omitempty"`              // Version
	TenantID             string                                                       `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfo struct {
	Source                    string                                                                   `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                   `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                    `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                   `json:"mode,omitempty"`                      // Mode
	State                     string                                                                   `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                   `json:"description,omitempty"`               // Description
	OnbState                  string                                                                   `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                   `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                   `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                 `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                 `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                   `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                 `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                 `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                   `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                   `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                 `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                 `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                   `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                   `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                   `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                   `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                   `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                   `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                   `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                   `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                 `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                   `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                   `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                    `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                 `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                   `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                 `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                 `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                 `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                    `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                   `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                   `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                    `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                   `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                   `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList struct {
	ProfileName       string                                                                            `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                             `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                            `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                   `json:"port,omitempty"`        // Port
	Protocol    string                                                                                     `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                     `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                     `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                     `json:"port,omitempty"`        // Port
	Protocol    string                                                                                       `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                       `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                       `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces struct {
	Status          string                                                                          `json:"status,omitempty"`          // Status
	MacAddress      string                                                                          `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                          `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                        `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                        `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                       `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                     `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                     `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpUpdateDeviceDeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow struct {
	TypeID         string                                                             `json:"_id,omitempty"`            // Id
	State          string                                                             `json:"state,omitempty"`          // State
	Type           string                                                             `json:"type,omitempty"`           // Type
	Description    string                                                             `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                           `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                             `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                           `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                           `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                              `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                             `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                           `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                           `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                           `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                             `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                             `json:"configId,omitempty"`       // Config Id
	Name           string                                                             `json:"name,omitempty"`           // Name
	Version        *float64                                                           `json:"version,omitempty"`        // Version
	TenantID       string                                                             `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks struct {
	State           string                                                                         `json:"state,omitempty"`           // State
	Type            string                                                                         `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                       `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                       `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                       `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                       `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                       `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                         `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceSystemWorkflow struct {
	TypeID         string                                                        `json:"_id,omitempty"`            // Id
	State          string                                                        `json:"state,omitempty"`          // State
	Type           string                                                        `json:"type,omitempty"`           // Type
	Description    string                                                        `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                      `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                        `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                      `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                      `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                         `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                        `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                      `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                      `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                      `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                        `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                        `json:"configId,omitempty"`       // Config Id
	Name           string                                                        `json:"name,omitempty"`           // Name
	Version        *float64                                                      `json:"version,omitempty"`        // Version
	TenantID       string                                                        `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks struct {
	State           string                                                                    `json:"state,omitempty"`           // State
	Type            string                                                                    `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                  `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                  `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                  `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                  `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                  `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                    `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceWorkflow struct {
	TypeID         string                                                  `json:"_id,omitempty"`            // Id
	State          string                                                  `json:"state,omitempty"`          // State
	Type           string                                                  `json:"type,omitempty"`           // Type
	Description    string                                                  `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                  `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateDeviceWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                   `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                  `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                  `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                  `json:"configId,omitempty"`       // Config Id
	Name           string                                                  `json:"name,omitempty"`           // Name
	Version        *float64                                                `json:"version,omitempty"`        // Version
	TenantID       string                                                  `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceWorkflowTasks struct {
	State           string                                                              `json:"state,omitempty"`           // State
	Type            string                                                              `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                            `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                            `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                            `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                            `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                            `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                              `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryList struct {
	Details         string                                                                `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                 `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                              `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo struct {
	Type         string                                                                              `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                            `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                              `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpUpdateDeviceWorkflowParameters struct {
	TopOfStackSerialNumber string                                                                 `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                 `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                 `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                 `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpUpdateDeviceDayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpUpdateDeviceDayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnp struct {
	TypeID               string                                                                  `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                                `json:"version,omitempty"`              // Version
	TenantID             string                                                                  `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfo struct {
	Source                    string                                                                              `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                              `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                               `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                              `json:"mode,omitempty"`                      // Mode
	State                     string                                                                              `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                              `json:"description,omitempty"`               // Description
	OnbState                  string                                                                              `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                              `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                              `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                            `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                            `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                              `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                            `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                            `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                              `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                              `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                            `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                            `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                              `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                              `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                              `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                              `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                              `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                              `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                              `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                              `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                            `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                              `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                              `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                               `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                            `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                              `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                            `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                            `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                            `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                               `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                              `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                              `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                               `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                              `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                              `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileList struct {
	ProfileName       string                                                                                       `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                                        `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                                       `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                              `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                                `json:"port,omitempty"`        // Port
	Protocol    string                                                                                                  `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                                  `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                                  `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoIPInterfaces struct {
	Status          string                                                                                     `json:"status,omitempty"`          // Status
	MacAddress      string                                                                                     `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                                     `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                                   `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                                   `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                  `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                                `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                                `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemResetWorkflow struct {
	TypeID         string                                                                        `json:"_id,omitempty"`            // Id
	State          string                                                                        `json:"state,omitempty"`          // State
	Type           string                                                                        `json:"type,omitempty"`           // Type
	Description    string                                                                        `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                      `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                        `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                      `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                      `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                         `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                        `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                      `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                      `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                      `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                        `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                        `json:"configId,omitempty"`       // Config Id
	Name           string                                                                        `json:"name,omitempty"`           // Name
	Version        *float64                                                                      `json:"version,omitempty"`        // Version
	TenantID       string                                                                        `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemResetWorkflowTasks struct {
	State           string                                                                                    `json:"state,omitempty"`           // State
	Type            string                                                                                    `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                                  `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                                  `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                                  `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                                  `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                                  `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                                    `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemWorkflow struct {
	TypeID         string                                                                   `json:"_id,omitempty"`            // Id
	State          string                                                                   `json:"state,omitempty"`          // State
	Type           string                                                                   `json:"type,omitempty"`           // Type
	Description    string                                                                   `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                                 `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                                   `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                                 `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                                 `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                                    `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                                   `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                                 `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                                 `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                                 `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                                   `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                                   `json:"configId,omitempty"`       // Config Id
	Name           string                                                                   `json:"name,omitempty"`           // Name
	Version        *float64                                                                 `json:"version,omitempty"`        // Version
	TenantID       string                                                                   `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemWorkflowTasks struct {
	State           string                                                                               `json:"state,omitempty"`           // State
	Type            string                                                                               `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                             `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                             `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                             `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                             `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                             `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                               `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpSystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflow struct {
	TypeID         string                                                             `json:"_id,omitempty"`            // Id
	State          string                                                             `json:"state,omitempty"`          // State
	Type           string                                                             `json:"type,omitempty"`           // Type
	Description    string                                                             `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                           `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                             `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                           `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                           `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                              `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                             `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                           `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                           `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                           `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                             `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                             `json:"configId,omitempty"`       // Config Id
	Name           string                                                             `json:"name,omitempty"`           // Name
	Version        *float64                                                           `json:"version,omitempty"`        // Version
	TenantID       string                                                             `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowTasks struct {
	State           string                                                                         `json:"state,omitempty"`           // State
	Type            string                                                                         `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                       `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                       `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                       `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                       `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                       `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                         `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryList struct {
	Details         string                                                                           `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                            `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                                         `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryListHistoryTaskInfo struct {
	Type         string                                                                                         `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                                       `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                         `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowParameters struct {
	TopOfStackSerialNumber string                                                                            `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                            `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                            `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                            `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnpDayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpGetDeviceByID struct {
	TypeID               string                                                        `json:"_id,omitempty"`                  // Id
	DeviceInfo           *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfo           `json:"deviceInfo,omitempty"`           //
	SystemResetWorkflow  *ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflow  `json:"systemResetWorkflow,omitempty"`  //
	SystemWorkflow       *ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflow       `json:"systemWorkflow,omitempty"`       //
	Workflow             *ResponseDeviceOnboardingPnpGetDeviceByIDWorkflow             `json:"workflow,omitempty"`             //
	RunSummaryList       *[]ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryList     `json:"runSummaryList,omitempty"`       //
	WorkflowParameters   *ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParameters   `json:"workflowParameters,omitempty"`   //
	DayZeroConfig        *ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfig        `json:"dayZeroConfig,omitempty"`        //
	DayZeroConfigPreview *ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfigPreview `json:"dayZeroConfigPreview,omitempty"` // Day Zero Config Preview
	Version              *float64                                                      `json:"version,omitempty"`              // Version
	TenantID             string                                                        `json:"tenantId,omitempty"`             // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfo struct {
	Source                    string                                                                    `json:"source,omitempty"`                    // Source
	SerialNumber              string                                                                    `json:"serialNumber,omitempty"`              // Serial Number
	Stack                     *bool                                                                     `json:"stack,omitempty"`                     // Stack
	Mode                      string                                                                    `json:"mode,omitempty"`                      // Mode
	State                     string                                                                    `json:"state,omitempty"`                     // State
	Location                  *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoLocation               `json:"location,omitempty"`                  //
	Description               string                                                                    `json:"description,omitempty"`               // Description
	OnbState                  string                                                                    `json:"onbState,omitempty"`                  // Onb State
	AuthenticatedMicNumber    string                                                                    `json:"authenticatedMicNumber,omitempty"`    // Authenticated Mic Number
	AuthenticatedSudiSerialNo string                                                                    `json:"authenticatedSudiSerialNo,omitempty"` // Authenticated Sudi Serial No
	CapabilitiesSupported     []string                                                                  `json:"capabilitiesSupported,omitempty"`     // Capabilities Supported
	FeaturesSupported         []string                                                                  `json:"featuresSupported,omitempty"`         // Features Supported
	CmState                   string                                                                    `json:"cmState,omitempty"`                   // Cm State
	FirstContact              *float64                                                                  `json:"firstContact,omitempty"`              // First Contact
	LastContact               *float64                                                                  `json:"lastContact,omitempty"`               // Last Contact
	MacAddress                string                                                                    `json:"macAddress,omitempty"`                // Mac Address
	Pid                       string                                                                    `json:"pid,omitempty"`                       // Pid
	DeviceSudiSerialNos       []string                                                                  `json:"deviceSudiSerialNos,omitempty"`       // Device Sudi Serial Nos
	LastUpdateOn              *float64                                                                  `json:"lastUpdateOn,omitempty"`              // Last Update On
	WorkflowID                string                                                                    `json:"workflowId,omitempty"`                // Workflow Id
	WorkflowName              string                                                                    `json:"workflowName,omitempty"`              // Workflow Name
	ProjectID                 string                                                                    `json:"projectId,omitempty"`                 // Project Id
	ProjectName               string                                                                    `json:"projectName,omitempty"`               // Project Name
	DeviceType                string                                                                    `json:"deviceType,omitempty"`                // Device Type
	AgentType                 string                                                                    `json:"agentType,omitempty"`                 // Agent Type
	ImageVersion              string                                                                    `json:"imageVersion,omitempty"`              // Image Version
	FileSystemList            *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	PnpProfileList            *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	ImageFile                 string                                                                    `json:"imageFile,omitempty"`                 // Image File
	HTTPHeaders               *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	NeighborLinks             *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	LastSyncTime              *float64                                                                  `json:"lastSyncTime,omitempty"`              // Last Sync Time
	IPInterfaces              *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	Hostname                  string                                                                    `json:"hostname,omitempty"`                  // Hostname
	AuthStatus                string                                                                    `json:"authStatus,omitempty"`                // Auth Status
	StackInfo                 *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	ReloadRequested           *bool                                                                     `json:"reloadRequested,omitempty"`           // Reload Requested
	AddedOn                   *float64                                                                  `json:"addedOn,omitempty"`                   // Added On
	SiteID                    string                                                                    `json:"siteId,omitempty"`                    // Site Id
	AAACredentials            *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	UserMicNumbers            []string                                                                  `json:"userMicNumbers,omitempty"`            // User Mic Numbers
	UserSudiSerialNos         []string                                                                  `json:"userSudiSerialNos,omitempty"`         // User Sudi Serial Nos
	AddnMacAddrs              []string                                                                  `json:"addnMacAddrs,omitempty"`              // Addn Mac Addrs
	PreWorkflowCliOuputs      *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	Tags                      *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoTags                   `json:"tags,omitempty"`                      // Tags
	SudiRequired              *bool                                                                     `json:"sudiRequired,omitempty"`              // Sudi Required
	SmartAccountID            string                                                                    `json:"smartAccountId,omitempty"`            // Smart Account Id
	VirtualAccountID          string                                                                    `json:"virtualAccountId,omitempty"`          // Virtual Account Id
	PopulateInventory         *bool                                                                     `json:"populateInventory,omitempty"`         // Populate Inventory
	SiteName                  string                                                                    `json:"siteName,omitempty"`                  // Site Name
	Name                      string                                                                    `json:"name,omitempty"`                      // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoLocation struct {
	SiteID    string `json:"siteId,omitempty"`    // Site Id
	Address   string `json:"address,omitempty"`   // Address
	Latitude  string `json:"latitude,omitempty"`  // Latitude
	Longitude string `json:"longitude,omitempty"` // Longitude
	Altitude  string `json:"altitude,omitempty"`  // Altitude
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoFileSystemList struct {
	Type      string   `json:"type,omitempty"`      // Type
	Writeable *bool    `json:"writeable,omitempty"` // Writeable
	Freespace *float64 `json:"freespace,omitempty"` // Freespace
	Name      string   `json:"name,omitempty"`      // Name
	Readable  *bool    `json:"readable,omitempty"`  // Readable
	Size      *float64 `json:"size,omitempty"`      // Size
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileList struct {
	ProfileName       string                                                                             `json:"profileName,omitempty"`       // Profile Name
	DiscoveryCreated  *bool                                                                              `json:"discoveryCreated,omitempty"`  // Discovery Created
	CreatedBy         string                                                                             `json:"createdBy,omitempty"`         // Created By
	PrimaryEndpoint   *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	SecondaryEndpoint *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Port        *float64                                                                                    `json:"port,omitempty"`        // Port
	Protocol    string                                                                                      `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                      `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                      `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Port        *float64                                                                                      `json:"port,omitempty"`        // Port
	Protocol    string                                                                                        `json:"protocol,omitempty"`    // Protocol
	IPv4Address *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` // Ipv4 Address
	IPv6Address *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` // Ipv6 Address
	Fqdn        string                                                                                        `json:"fqdn,omitempty"`        // Fqdn
	Certificate string                                                                                        `json:"certificate,omitempty"` // Certificate
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoNeighborLinks struct {
	LocalInterfaceName       string `json:"localInterfaceName,omitempty"`       // Local Interface Name
	LocalShortInterfaceName  string `json:"localShortInterfaceName,omitempty"`  // Local Short Interface Name
	LocalMacAddress          string `json:"localMacAddress,omitempty"`          // Local Mac Address
	RemoteInterfaceName      string `json:"remoteInterfaceName,omitempty"`      // Remote Interface Name
	RemoteShortInterfaceName string `json:"remoteShortInterfaceName,omitempty"` // Remote Short Interface Name
	RemoteMacAddress         string `json:"remoteMacAddress,omitempty"`         // Remote Mac Address
	RemoteDeviceName         string `json:"remoteDeviceName,omitempty"`         // Remote Device Name
	RemotePlatform           string `json:"remotePlatform,omitempty"`           // Remote Platform
	RemoteVersion            string `json:"remoteVersion,omitempty"`            // Remote Version
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfaces struct {
	Status          string                                                                           `json:"status,omitempty"`          // Status
	MacAddress      string                                                                           `json:"macAddress,omitempty"`      // Mac Address
	IPv4Address     *ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     // Ipv4 Address
	IPv6AddressList *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` // Ipv6 Address List
	Name            string                                                                           `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv4Address interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoIPInterfacesIPv6AddressList interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfo struct {
	SupportsStackWorkflows *bool                                                                         `json:"supportsStackWorkflows,omitempty"` // Supports Stack Workflows
	IsFullRing             *bool                                                                         `json:"isFullRing,omitempty"`             // Is Full Ring
	StackMemberList        *[]ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                        `json:"stackRingProtocol,omitempty"`      // Stack Ring Protocol
	ValidLicenseLevels     []string                                                                      `json:"validLicenseLevels,omitempty"`     // Valid License Levels
	TotalMemberCount       *float64                                                                      `json:"totalMemberCount,omitempty"`       // Total Member Count
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoStackInfoStackMemberList struct {
	SerialNumber     string   `json:"serialNumber,omitempty"`     // Serial Number
	State            string   `json:"state,omitempty"`            // State
	Role             string   `json:"role,omitempty"`             // Role
	MacAddress       string   `json:"macAddress,omitempty"`       // Mac Address
	Pid              string   `json:"pid,omitempty"`              // Pid
	LicenseLevel     string   `json:"licenseLevel,omitempty"`     // License Level
	LicenseType      string   `json:"licenseType,omitempty"`      // License Type
	SudiSerialNumber string   `json:"sudiSerialNumber,omitempty"` // Sudi Serial Number
	HardwareVersion  string   `json:"hardwareVersion,omitempty"`  // Hardware Version
	StackNumber      *float64 `json:"stackNumber,omitempty"`      // Stack Number
	SoftwareVersion  string   `json:"softwareVersion,omitempty"`  // Software Version
	Priority         *float64 `json:"priority,omitempty"`         // Priority
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       // Cli
	CliOutput string `json:"cliOutput,omitempty"` // Cli Output
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDeviceInfoTags interface{}
type ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflow struct {
	TypeID         string                                                              `json:"_id,omitempty"`            // Id
	State          string                                                              `json:"state,omitempty"`          // State
	Type           string                                                              `json:"type,omitempty"`           // Type
	Description    string                                                              `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                            `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                              `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                            `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                            `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                               `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                              `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                            `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                            `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                            `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                              `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                              `json:"configId,omitempty"`       // Config Id
	Name           string                                                              `json:"name,omitempty"`           // Name
	Version        *float64                                                            `json:"version,omitempty"`        // Version
	TenantID       string                                                              `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasks struct {
	State           string                                                                          `json:"state,omitempty"`           // State
	Type            string                                                                          `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                        `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                        `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                        `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                        `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                        `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                          `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDSystemResetWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflow struct {
	TypeID         string                                                         `json:"_id,omitempty"`            // Id
	State          string                                                         `json:"state,omitempty"`          // State
	Type           string                                                         `json:"type,omitempty"`           // Type
	Description    string                                                         `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                       `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                         `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                       `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                       `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                          `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                         `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                       `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                       `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                       `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                         `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                         `json:"configId,omitempty"`       // Config Id
	Name           string                                                         `json:"name,omitempty"`           // Name
	Version        *float64                                                       `json:"version,omitempty"`        // Version
	TenantID       string                                                         `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasks struct {
	State           string                                                                     `json:"state,omitempty"`           // State
	Type            string                                                                     `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                                   `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                                   `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                                   `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                                   `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                                   `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                                     `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDSystemWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDWorkflow struct {
	TypeID         string                                                   `json:"_id,omitempty"`            // Id
	State          string                                                   `json:"state,omitempty"`          // State
	Type           string                                                   `json:"type,omitempty"`           // Type
	Description    string                                                   `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                                 `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                   `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                                 `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                                 `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                    `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                   `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                                 `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                                 `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                                 `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                   `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                   `json:"configId,omitempty"`       // Config Id
	Name           string                                                   `json:"name,omitempty"`           // Name
	Version        *float64                                                 `json:"version,omitempty"`        // Version
	TenantID       string                                                   `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasks struct {
	State           string                                                               `json:"state,omitempty"`           // State
	Type            string                                                               `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                             `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                             `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                             `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                             `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                             `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                               `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryList struct {
	Details         string                                                                 `json:"details,omitempty"`         // Details
	HistoryTaskInfo *ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	ErrorFlag       *bool                                                                  `json:"errorFlag,omitempty"`       // Error Flag
	Timestamp       *float64                                                               `json:"timestamp,omitempty"`       // Timestamp
}
type ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfo struct {
	Type         string                                                                               `json:"type,omitempty"`         // Type
	WorkItemList *[]ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
	TimeTaken    *float64                                                                             `json:"timeTaken,omitempty"`    // Time Taken
	AddnDetails  *[]ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                               `json:"name,omitempty"`         // Name
}
type ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetDeviceByIDRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParameters struct {
	TopOfStackSerialNumber string                                                                  `json:"topOfStackSerialNumber,omitempty"` // Top Of Stack Serial Number
	LicenseLevel           string                                                                  `json:"licenseLevel,omitempty"`           // License Level
	LicenseType            string                                                                  `json:"licenseType,omitempty"`            // License Type
	ConfigList             *[]ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigList `json:"configList,omitempty"`             //
}
type ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigList struct {
	ConfigParameters *[]ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
	ConfigID         string                                                                                  `json:"configId,omitempty"`         // Config Id
}
type ResponseDeviceOnboardingPnpGetDeviceByIDWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   // Key
	Value string `json:"value,omitempty"` // Value
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfig struct {
	Config string `json:"config,omitempty"` // Config
}
type ResponseDeviceOnboardingPnpGetDeviceByIDDayZeroConfigPreview interface{}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettings struct {
	SavaMappingList *[]ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsTaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                                               `json:"tenantId,omitempty"`        // Tenant Id
	AAACredentials  *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsAAACredentials    `json:"aaaCredentials,omitempty"`  //
	DefaultProfile  *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsDefaultProfile    `json:"defaultProfile,omitempty"`  //
	AcceptEula      *bool                                                                `json:"acceptEula,omitempty"`      // Accept Eula
	ID              string                                                               `json:"id,omitempty"`              // Id
	TypeID          string                                                               `json:"_id,omitempty"`             // Id
	Version         *float64                                                             `json:"version,omitempty"`         // Version
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList struct {
	SyncStatus       string                                                                       `json:"syncStatus,omitempty"`       // Sync Status
	SyncStartTime    *float64                                                                     `json:"syncStartTime,omitempty"`    // Sync Start Time
	SyncResult       *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	LastSync         *float64                                                                     `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                                       `json:"tenantId,omitempty"`         // Tenant Id
	Profile          *ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListProfile    `json:"profile,omitempty"`          //
	Token            string                                                                       `json:"token,omitempty"`            // Token
	Expiry           *float64                                                                     `json:"expiry,omitempty"`           // Expiry
	CcoUser          string                                                                       `json:"ccoUser,omitempty"`          // Cco User
	SmartAccountID   string                                                                       `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                                       `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                                     `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                                       `json:"syncResultStr,omitempty"`    // Sync Result Str
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                                 `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListProfile struct {
	Port        *float64 `json:"port,omitempty"`        // Port
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Cert        string   `json:"cert,omitempty"`        // Cert
	Name        string   `json:"name,omitempty"`        // Name
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsTaskTimeOuts struct {
	ImageDownloadTimeOut *float64 `json:"imageDownloadTimeOut,omitempty"` // Image Download Time Out
	ConfigTimeOut        *float64 `json:"configTimeOut,omitempty"`        // Config Time Out
	GeneralTimeOut       *float64 `json:"generalTimeOut,omitempty"`       // General Time Out
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpUpdatePnpGlobalSettingsDefaultProfile struct {
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` // Fqdn Addresses
	Proxy         *bool    `json:"proxy,omitempty"`         // Proxy
	Cert          string   `json:"cert,omitempty"`          // Cert
	IPAddresses   []string `json:"ipAddresses,omitempty"`   // Ip Addresses
	Port          *float64 `json:"port,omitempty"`          // Port
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettings struct {
	SavaMappingList *[]ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsTaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                                            `json:"tenantId,omitempty"`        // Tenant Id
	AAACredentials  *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsAAACredentials    `json:"aaaCredentials,omitempty"`  //
	DefaultProfile  *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsDefaultProfile    `json:"defaultProfile,omitempty"`  //
	AcceptEula      *bool                                                             `json:"acceptEula,omitempty"`      // Accept Eula
	ID              string                                                            `json:"id,omitempty"`              // Id
	TypeID          string                                                            `json:"_id,omitempty"`             // Id
	Version         *float64                                                          `json:"version,omitempty"`         // Version
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingList struct {
	SyncStatus       string                                                                    `json:"syncStatus,omitempty"`       // Sync Status
	SyncStartTime    *float64                                                                  `json:"syncStartTime,omitempty"`    // Sync Start Time
	SyncResult       *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	LastSync         *float64                                                                  `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                                    `json:"tenantId,omitempty"`         // Tenant Id
	Profile          *ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListProfile    `json:"profile,omitempty"`          //
	Token            string                                                                    `json:"token,omitempty"`            // Token
	Expiry           *float64                                                                  `json:"expiry,omitempty"`           // Expiry
	CcoUser          string                                                                    `json:"ccoUser,omitempty"`          // Cco User
	SmartAccountID   string                                                                    `json:"smartAccountId,omitempty"`   // Smart Account Id
	VirtualAccountID string                                                                    `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                                  `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                                    `json:"syncResultStr,omitempty"`    // Sync Result Str
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                              `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsSavaMappingListProfile struct {
	Port        *float64 `json:"port,omitempty"`        // Port
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Cert        string   `json:"cert,omitempty"`        // Cert
	Name        string   `json:"name,omitempty"`        // Name
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsTaskTimeOuts struct {
	ImageDownloadTimeOut *float64 `json:"imageDownloadTimeOut,omitempty"` // Image Download Time Out
	ConfigTimeOut        *float64 `json:"configTimeOut,omitempty"`        // Config Time Out
	GeneralTimeOut       *float64 `json:"generalTimeOut,omitempty"`       // General Time Out
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsAAACredentials struct {
	Password string `json:"password,omitempty"` // Password
	Username string `json:"username,omitempty"` // Username
}
type ResponseDeviceOnboardingPnpGetPnpGlobalSettingsDefaultProfile struct {
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` // Fqdn Addresses
	Proxy         *bool    `json:"proxy,omitempty"`         // Proxy
	Cert          string   `json:"cert,omitempty"`          // Cert
	IPAddresses   []string `json:"ipAddresses,omitempty"`   // Ip Addresses
	Port          *float64 `json:"port,omitempty"`          // Port
}
type ResponseDeviceOnboardingPnpGetSmartAccountList []string   // Array of ResponseDeviceOnboardingPnpGetSmartAccountList
type ResponseDeviceOnboardingPnpGetVirtualAccountList []string // Array of ResponseDeviceOnboardingPnpGetVirtualAccountList
type ResponseDeviceOnboardingPnpAddVirtualAccount struct {
	VirtualAccountID string                                                  `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                  `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpAddVirtualAccountProfile    `json:"profile,omitempty"`          //
	CcoUser          string                                                  `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpAddVirtualAccountSyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                  `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                  `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                  `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                  `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpAddVirtualAccountProfile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpAddVirtualAccountSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpAddVirtualAccountSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                            `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpAddVirtualAccountSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfile struct {
	VirtualAccountID string                                                       `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                     `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                       `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpUpdatePnpServerProfileProfile    `json:"profile,omitempty"`          //
	CcoUser          string                                                       `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpUpdatePnpServerProfileSyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                       `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                     `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                     `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                       `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                       `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                     `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                       `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileProfile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpUpdatePnpServerProfileSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                 `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpUpdatePnpServerProfileSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpDeregisterVirtualAccount struct {
	VirtualAccountID string                                                         `json:"virtualAccountId,omitempty"` // Virtual Account Id
	AutoSyncPeriod   *float64                                                       `json:"autoSyncPeriod,omitempty"`   // Auto Sync Period
	SyncResultStr    string                                                         `json:"syncResultStr,omitempty"`    // Sync Result Str
	Profile          *ResponseDeviceOnboardingPnpDeregisterVirtualAccountProfile    `json:"profile,omitempty"`          //
	CcoUser          string                                                         `json:"ccoUser,omitempty"`          // Cco User
	SyncResult       *ResponseDeviceOnboardingPnpDeregisterVirtualAccountSyncResult `json:"syncResult,omitempty"`       //
	Token            string                                                         `json:"token,omitempty"`            // Token
	SyncStartTime    *float64                                                       `json:"syncStartTime,omitempty"`    // Sync Start Time
	LastSync         *float64                                                       `json:"lastSync,omitempty"`         // Last Sync
	TenantID         string                                                         `json:"tenantId,omitempty"`         // Tenant Id
	SmartAccountID   string                                                         `json:"smartAccountId,omitempty"`   // Smart Account Id
	Expiry           *float64                                                       `json:"expiry,omitempty"`           // Expiry
	SyncStatus       string                                                         `json:"syncStatus,omitempty"`       // Sync Status
}
type ResponseDeviceOnboardingPnpDeregisterVirtualAccountProfile struct {
	Proxy       *bool    `json:"proxy,omitempty"`       // Proxy
	MakeDefault *bool    `json:"makeDefault,omitempty"` // Make Default
	Port        *float64 `json:"port,omitempty"`        // Port
	ProfileID   string   `json:"profileId,omitempty"`   // Profile Id
	Name        string   `json:"name,omitempty"`        // Name
	AddressIPV4 string   `json:"addressIpV4,omitempty"` // Address Ip V4
	Cert        string   `json:"cert,omitempty"`        // Cert
	AddressFqdn string   `json:"addressFqdn,omitempty"` // Address Fqdn
}
type ResponseDeviceOnboardingPnpDeregisterVirtualAccountSyncResult struct {
	SyncList *[]ResponseDeviceOnboardingPnpDeregisterVirtualAccountSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                   `json:"syncMsg,omitempty"`  // Sync Msg
}
type ResponseDeviceOnboardingPnpDeregisterVirtualAccountSyncResultSyncList struct {
	SyncType     string   `json:"syncType,omitempty"`     // Sync Type
	DeviceSnList []string `json:"deviceSnList,omitempty"` // Device Sn List
}
type ResponseDeviceOnboardingPnpGetWorkflows []ResponseItemDeviceOnboardingPnpGetWorkflows // Array of ResponseDeviceOnboardingPnpGetWorkflows
type ResponseItemDeviceOnboardingPnpGetWorkflows struct {
	TypeID         string                                              `json:"_id,omitempty"`            // Id
	State          string                                              `json:"state,omitempty"`          // State
	Type           string                                              `json:"type,omitempty"`           // Type
	Description    string                                              `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                            `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                              `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                            `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                            `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseItemDeviceOnboardingPnpGetWorkflowsTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                               `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                              `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                            `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                            `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                            `json:"startTime,omitempty"`      // Start Time
	UseState       string                                              `json:"useState,omitempty"`       // Use State
	ConfigID       string                                              `json:"configId,omitempty"`       // Config Id
	Name           string                                              `json:"name,omitempty"`           // Name
	Version        *float64                                            `json:"version,omitempty"`        // Version
	TenantID       string                                              `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseItemDeviceOnboardingPnpGetWorkflowsTasks struct {
	State           string                                                          `json:"state,omitempty"`           // State
	Type            string                                                          `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                        `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                        `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                        `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                        `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseItemDeviceOnboardingPnpGetWorkflowsTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                        `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                          `json:"name,omitempty"`            // Name
}
type ResponseItemDeviceOnboardingPnpGetWorkflowsTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpAddAWorkflow struct {
	TypeID         string                                          `json:"_id,omitempty"`            // Id
	State          string                                          `json:"state,omitempty"`          // State
	Type           string                                          `json:"type,omitempty"`           // Type
	Description    string                                          `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                        `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                          `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                        `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                        `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpAddAWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                           `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                          `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                        `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                        `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                        `json:"startTime,omitempty"`      // Start Time
	UseState       string                                          `json:"useState,omitempty"`       // Use State
	ConfigID       string                                          `json:"configId,omitempty"`       // Config Id
	Name           string                                          `json:"name,omitempty"`           // Name
	Version        *float64                                        `json:"version,omitempty"`        // Version
	TenantID       string                                          `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpAddAWorkflowTasks struct {
	State           string                                                      `json:"state,omitempty"`           // State
	Type            string                                                      `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                    `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                    `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                    `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                    `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpAddAWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                    `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                      `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpAddAWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpGetWorkflowCount struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseDeviceOnboardingPnpGetWorkflowByID struct {
	TypeID         string                                             `json:"_id,omitempty"`            // Id
	State          string                                             `json:"state,omitempty"`          // State
	Type           string                                             `json:"type,omitempty"`           // Type
	Description    string                                             `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                           `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                             `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                           `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                           `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpGetWorkflowByIDTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                              `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                             `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                           `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                           `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                           `json:"startTime,omitempty"`      // Start Time
	UseState       string                                             `json:"useState,omitempty"`       // Use State
	ConfigID       string                                             `json:"configId,omitempty"`       // Config Id
	Name           string                                             `json:"name,omitempty"`           // Name
	Version        *float64                                           `json:"version,omitempty"`        // Version
	TenantID       string                                             `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDTasks struct {
	State           string                                                         `json:"state,omitempty"`           // State
	Type            string                                                         `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                       `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                       `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                       `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                       `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpGetWorkflowByIDTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                       `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                         `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpGetWorkflowByIDTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpDeleteWorkflowByID struct {
	TypeID         string                                                `json:"_id,omitempty"`            // Id
	State          string                                                `json:"state,omitempty"`          // State
	Type           string                                                `json:"type,omitempty"`           // Type
	Description    string                                                `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                              `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                                `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                              `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                              `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpDeleteWorkflowByIDTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                                 `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                                `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                              `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                              `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                              `json:"startTime,omitempty"`      // Start Time
	UseState       string                                                `json:"useState,omitempty"`       // Use State
	ConfigID       string                                                `json:"configId,omitempty"`       // Config Id
	Name           string                                                `json:"name,omitempty"`           // Name
	Version        *float64                                              `json:"version,omitempty"`        // Version
	TenantID       string                                                `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpDeleteWorkflowByIDTasks struct {
	State           string                                                            `json:"state,omitempty"`           // State
	Type            string                                                            `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                          `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                          `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                          `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                          `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpDeleteWorkflowByIDTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                          `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                            `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpDeleteWorkflowByIDTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type ResponseDeviceOnboardingPnpUpdateWorkflow struct {
	TypeID         string                                            `json:"_id,omitempty"`            // Id
	State          string                                            `json:"state,omitempty"`          // State
	Type           string                                            `json:"type,omitempty"`           // Type
	Description    string                                            `json:"description,omitempty"`    // Description
	LastupdateOn   *float64                                          `json:"lastupdateOn,omitempty"`   // Lastupdate On
	ImageID        string                                            `json:"imageId,omitempty"`        // Image Id
	CurrTaskIDx    *float64                                          `json:"currTaskIdx,omitempty"`    // Curr Task Idx
	AddedOn        *float64                                          `json:"addedOn,omitempty"`        // Added On
	Tasks          *[]ResponseDeviceOnboardingPnpUpdateWorkflowTasks `json:"tasks,omitempty"`          //
	AddToInventory *bool                                             `json:"addToInventory,omitempty"` // Add To Inventory
	InstanceType   string                                            `json:"instanceType,omitempty"`   // Instance Type
	EndTime        *float64                                          `json:"endTime,omitempty"`        // End Time
	ExecTime       *float64                                          `json:"execTime,omitempty"`       // Exec Time
	StartTime      *float64                                          `json:"startTime,omitempty"`      // Start Time
	UseState       string                                            `json:"useState,omitempty"`       // Use State
	ConfigID       string                                            `json:"configId,omitempty"`       // Config Id
	Name           string                                            `json:"name,omitempty"`           // Name
	Version        *float64                                          `json:"version,omitempty"`        // Version
	TenantID       string                                            `json:"tenantId,omitempty"`       // Tenant Id
}
type ResponseDeviceOnboardingPnpUpdateWorkflowTasks struct {
	State           string                                                        `json:"state,omitempty"`           // State
	Type            string                                                        `json:"type,omitempty"`            // Type
	CurrWorkItemIDx *float64                                                      `json:"currWorkItemIdx,omitempty"` // Curr Work Item Idx
	TaskSeqNo       *float64                                                      `json:"taskSeqNo,omitempty"`       // Task Seq No
	EndTime         *float64                                                      `json:"endTime,omitempty"`         // End Time
	StartTime       *float64                                                      `json:"startTime,omitempty"`       // Start Time
	WorkItemList    *[]ResponseDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
	TimeTaken       *float64                                                      `json:"timeTaken,omitempty"`       // Time Taken
	Name            string                                                        `json:"name,omitempty"`            // Name
}
type ResponseDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList struct {
	State     string   `json:"state,omitempty"`     // State
	Command   string   `json:"command,omitempty"`   // Command
	OutputStr string   `json:"outputStr,omitempty"` // Output Str
	EndTime   *float64 `json:"endTime,omitempty"`   // End Time
	StartTime *float64 `json:"startTime,omitempty"` // Start Time
	TimeTaken *float64 `json:"timeTaken,omitempty"` // Time Taken
}
type RequestDeviceOnboardingPnpAddDevice struct {
	TypeID              string                                                  `json:"_id,omitempty"`                 //
	DeviceInfo          *RequestDeviceOnboardingPnpAddDeviceDeviceInfo          `json:"deviceInfo,omitempty"`          //
	RunSummaryList      *[]RequestDeviceOnboardingPnpAddDeviceRunSummaryList    `json:"runSummaryList,omitempty"`      //
	SystemResetWorkflow *RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow `json:"systemResetWorkflow,omitempty"` //
	SystemWorkflow      *RequestDeviceOnboardingPnpAddDeviceSystemWorkflow      `json:"systemWorkflow,omitempty"`      //
	TenantID            string                                                  `json:"tenantId,omitempty"`            //
	Version             *int                                                    `json:"version,omitempty"`             //
	Workflow            *RequestDeviceOnboardingPnpAddDeviceWorkflow            `json:"workflow,omitempty"`            //
	WorkflowParameters  *RequestDeviceOnboardingPnpAddDeviceWorkflowParameters  `json:"workflowParameters,omitempty"`  //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfo struct {
	AAACredentials            *RequestDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   *int                                                                 `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                             `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                               `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                               `json:"authStatus,omitempty"`                //
	AuthenticatedSudiSerialNo string                                                               `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                             `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                               `json:"cmState,omitempty"`                   //
	Description               string                                                               `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                             `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                               `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                             `json:"featuresSupported,omitempty"`         //
	FileSystemList            *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              *int                                                                 `json:"firstContact,omitempty"`              //
	Hostname                  string                                                               `json:"hostname,omitempty"`                  //
	HTTPHeaders               *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                               `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                               `json:"imageVersion,omitempty"`              //
	IPInterfaces              *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               *int                                                                 `json:"lastContact,omitempty"`               //
	LastSyncTime              *int                                                                 `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              *int                                                                 `json:"lastUpdateOn,omitempty"`              //
	Location                  *RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                               `json:"macAddress,omitempty"`                //
	Mode                      string                                                               `json:"mode,omitempty"`                      //
	Name                      string                                                               `json:"name,omitempty"`                      //
	NeighborLinks             *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                               `json:"onbState,omitempty"`                  //
	Pid                       string                                                               `json:"pid,omitempty"`                       //
	PnpProfileList            *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         *bool                                                                `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                               `json:"projectId,omitempty"`                 //
	ProjectName               string                                                               `json:"projectName,omitempty"`               //
	ReloadRequested           *bool                                                                `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                               `json:"serialNumber,omitempty"`              //
	SmartAccountID            string                                                               `json:"smartAccountId,omitempty"`            //
	Source                    string                                                               `json:"source,omitempty"`                    //
	Stack                     *bool                                                                `json:"stack,omitempty"`                     //
	StackInfo                 *RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                               `json:"state,omitempty"`                     //
	SudiRequired              *bool                                                                `json:"sudiRequired,omitempty"`              //
	Tags                      *RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags                   `json:"tags,omitempty"`                      //
	UserSudiSerialNos         []string                                                             `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                               `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                               `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                               `json:"workflowName,omitempty"`              //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoFileSystemList struct {
	Freespace *int   `json:"freespace,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	Readable  *bool  `json:"readable,omitempty"`  //
	Size      *int   `json:"size,omitempty"`      //
	Type      string `json:"type,omitempty"`      //
	Writeable *bool  `json:"writeable,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfaces struct {
	IPv4Address     *RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     //
	IPv6AddressList *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` //
	MacAddress      string                                                                      `json:"macAddress,omitempty"`      //
	Name            string                                                                      `json:"name,omitempty"`            //
	Status          string                                                                      `json:"status,omitempty"`          //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv4Address interface{}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoIPInterfacesIPv6AddressList interface{}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoNeighborLinks struct {
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
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                        `json:"createdBy,omitempty"`         //
	DiscoveryCreated  *bool                                                                         `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   *RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                        `json:"profileName,omitempty"`       //
	SecondaryEndpoint *RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string                                                                                 `json:"certificate,omitempty"` //
	Fqdn        string                                                                                 `json:"fqdn,omitempty"`        //
	IPv4Address *RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` //
	IPv6Address *RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` //
	Port        *int                                                                                   `json:"port,omitempty"`        //
	Protocol    string                                                                                 `json:"protocol,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string                                                                                   `json:"certificate,omitempty"` //
	Fqdn        string                                                                                   `json:"fqdn,omitempty"`        //
	IPv4Address *RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` //
	IPv6Address *RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` //
	Port        *int                                                                                     `json:"port,omitempty"`        //
	Protocol    string                                                                                   `json:"protocol,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfo struct {
	IsFullRing             *bool                                                                    `json:"isFullRing,omitempty"`             //
	StackMemberList        *[]RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                   `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows *bool                                                                    `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       *int                                                                     `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                                 `json:"validLicenseLevels,omitempty"`     //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string `json:"licenseLevel,omitempty"`     //
	LicenseType      string `json:"licenseType,omitempty"`      //
	MacAddress       string `json:"macAddress,omitempty"`       //
	Pid              string `json:"pid,omitempty"`              //
	Priority         *int   `json:"priority,omitempty"`         //
	Role             string `json:"role,omitempty"`             //
	SerialNumber     string `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string `json:"softwareVersion,omitempty"`  //
	StackNumber      *int   `json:"stackNumber,omitempty"`      //
	State            string `json:"state,omitempty"`            //
	SudiSerialNumber string `json:"sudiSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceDeviceInfoTags interface{}
type RequestDeviceOnboardingPnpAddDeviceRunSummaryList struct {
	Details         string                                                            `json:"details,omitempty"`         //
	ErrorFlag       *bool                                                             `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo *RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       *int                                                              `json:"timestamp,omitempty"`       //
}
type RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfo struct {
	AddnDetails  *[]RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                          `json:"name,omitempty"`         //
	TimeTaken    *int                                                                            `json:"timeTaken,omitempty"`    //
	Type         string                                                                          `json:"type,omitempty"`         //
	WorkItemList *[]RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflow struct {
	TypeID         string                                                         `json:"_id,omitempty"`            //
	AddToInventory *bool                                                          `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                           `json:"addedOn,omitempty"`        //
	ConfigID       string                                                         `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                           `json:"currTaskIdx,omitempty"`    //
	Description    string                                                         `json:"description,omitempty"`    //
	EndTime        *int                                                           `json:"endTime,omitempty"`        //
	ExecTime       *int                                                           `json:"execTime,omitempty"`       //
	ImageID        string                                                         `json:"imageId,omitempty"`        //
	InstanceType   string                                                         `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                           `json:"lastupdateOn,omitempty"`   //
	Name           string                                                         `json:"name,omitempty"`           //
	StartTime      *int                                                           `json:"startTime,omitempty"`      //
	State          string                                                         `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                         `json:"tenantId,omitempty"`       //
	Type           string                                                         `json:"type,omitempty"`           //
	UseState       string                                                         `json:"useState,omitempty"`       //
	Version        *int                                                           `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                       `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                       `json:"endTime,omitempty"`         //
	Name            string                                                                     `json:"name,omitempty"`            //
	StartTime       *int                                                                       `json:"startTime,omitempty"`       //
	State           string                                                                     `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                       `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                       `json:"timeTaken,omitempty"`       //
	Type            string                                                                     `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddDeviceSystemResetWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceSystemWorkflow struct {
	TypeID         string                                                    `json:"_id,omitempty"`            //
	AddToInventory *bool                                                     `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                      `json:"addedOn,omitempty"`        //
	ConfigID       string                                                    `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                      `json:"currTaskIdx,omitempty"`    //
	Description    string                                                    `json:"description,omitempty"`    //
	EndTime        *int                                                      `json:"endTime,omitempty"`        //
	ExecTime       *int                                                      `json:"execTime,omitempty"`       //
	ImageID        string                                                    `json:"imageId,omitempty"`        //
	InstanceType   string                                                    `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                      `json:"lastupdateOn,omitempty"`   //
	Name           string                                                    `json:"name,omitempty"`           //
	StartTime      *int                                                      `json:"startTime,omitempty"`      //
	State          string                                                    `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                    `json:"tenantId,omitempty"`       //
	Type           string                                                    `json:"type,omitempty"`           //
	UseState       string                                                    `json:"useState,omitempty"`       //
	Version        *int                                                      `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                  `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                  `json:"endTime,omitempty"`         //
	Name            string                                                                `json:"name,omitempty"`            //
	StartTime       *int                                                                  `json:"startTime,omitempty"`       //
	State           string                                                                `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                  `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                  `json:"timeTaken,omitempty"`       //
	Type            string                                                                `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddDeviceSystemWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceWorkflow struct {
	TypeID         string                                              `json:"_id,omitempty"`            //
	AddToInventory *bool                                               `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                `json:"addedOn,omitempty"`        //
	ConfigID       string                                              `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                `json:"currTaskIdx,omitempty"`    //
	Description    string                                              `json:"description,omitempty"`    //
	EndTime        *int                                                `json:"endTime,omitempty"`        //
	ExecTime       *int                                                `json:"execTime,omitempty"`       //
	ImageID        string                                              `json:"imageId,omitempty"`        //
	InstanceType   string                                              `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                `json:"lastupdateOn,omitempty"`   //
	Name           string                                              `json:"name,omitempty"`           //
	StartTime      *int                                                `json:"startTime,omitempty"`      //
	State          string                                              `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpAddDeviceWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                              `json:"tenantId,omitempty"`       //
	Type           string                                              `json:"type,omitempty"`           //
	UseState       string                                              `json:"useState,omitempty"`       //
	Version        *int                                                `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpAddDeviceWorkflowTasks struct {
	CurrWorkItemIDx *int                                                            `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                            `json:"endTime,omitempty"`         //
	Name            string                                                          `json:"name,omitempty"`            //
	StartTime       *int                                                            `json:"startTime,omitempty"`       //
	State           string                                                          `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                            `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                            `json:"timeTaken,omitempty"`       //
	Type            string                                                          `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddDeviceWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceWorkflowParameters struct {
	ConfigList             *[]RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                             `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                             `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                             `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigList struct {
	ConfigID         string                                                                             `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpAddDeviceWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimDevice struct {
	ConfigFileURL     string                                                  `json:"configFileUrl,omitempty"`     //
	ConfigID          string                                                  `json:"configId,omitempty"`          //
	DeviceClaimList   *[]RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList `json:"deviceClaimList,omitempty"`   //
	FileServiceID     string                                                  `json:"fileServiceId,omitempty"`     //
	ImageID           string                                                  `json:"imageId,omitempty"`           //
	ImageURL          string                                                  `json:"imageUrl,omitempty"`          //
	PopulateInventory *bool                                                   `json:"populateInventory,omitempty"` //
	ProjectID         string                                                  `json:"projectId,omitempty"`         //
	WorkflowID        string                                                  `json:"workflowId,omitempty"`        //
}
type RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList struct {
	ConfigList             *[]RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList `json:"configList,omitempty"`             //
	DeviceID               string                                                            `json:"deviceId,omitempty"`               //
	LicenseLevel           string                                                            `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                            `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                            `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList struct {
	ConfigID         string                                                                            `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpImportDevicesInBulk []RequestItemDeviceOnboardingPnpImportDevicesInBulk // Array of RequestDeviceOnboardingPnpImportDevicesInBulk
type RequestItemDeviceOnboardingPnpImportDevicesInBulk struct {
	TypeID              string                                                                `json:"_id,omitempty"`                 //
	DeviceInfo          *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfo          `json:"deviceInfo,omitempty"`          //
	RunSummaryList      *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryList    `json:"runSummaryList,omitempty"`      //
	SystemResetWorkflow *RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflow `json:"systemResetWorkflow,omitempty"` //
	SystemWorkflow      *RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflow      `json:"systemWorkflow,omitempty"`      //
	TenantID            string                                                                `json:"tenantId,omitempty"`            //
	Version             *int                                                                  `json:"version,omitempty"`             //
	Workflow            *RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflow            `json:"workflow,omitempty"`            //
	WorkflowParameters  *RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParameters  `json:"workflowParameters,omitempty"`  //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfo struct {
	AAACredentials            *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   *int                                                                               `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                                           `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                                             `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                                             `json:"authStatus,omitempty"`                //
	AuthenticatedSudiSerialNo string                                                                             `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                                           `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                                             `json:"cmState,omitempty"`                   //
	Description               string                                                                             `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                                           `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                                             `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                                           `json:"featuresSupported,omitempty"`         //
	FileSystemList            *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              *int                                                                               `json:"firstContact,omitempty"`              //
	Hostname                  string                                                                             `json:"hostname,omitempty"`                  //
	HTTPHeaders               *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                                             `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                                             `json:"imageVersion,omitempty"`              //
	IPInterfaces              *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               *int                                                                               `json:"lastContact,omitempty"`               //
	LastSyncTime              *int                                                                               `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              *int                                                                               `json:"lastUpdateOn,omitempty"`              //
	Location                  *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                                             `json:"macAddress,omitempty"`                //
	Mode                      string                                                                             `json:"mode,omitempty"`                      //
	Name                      string                                                                             `json:"name,omitempty"`                      //
	NeighborLinks             *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                                             `json:"onbState,omitempty"`                  //
	Pid                       string                                                                             `json:"pid,omitempty"`                       //
	PnpProfileList            *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         *bool                                                                              `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                                             `json:"projectId,omitempty"`                 //
	ProjectName               string                                                                             `json:"projectName,omitempty"`               //
	ReloadRequested           *bool                                                                              `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                                             `json:"serialNumber,omitempty"`              //
	SmartAccountID            string                                                                             `json:"smartAccountId,omitempty"`            //
	Source                    string                                                                             `json:"source,omitempty"`                    //
	Stack                     *bool                                                                              `json:"stack,omitempty"`                     //
	StackInfo                 *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                                             `json:"state,omitempty"`                     //
	SudiRequired              *bool                                                                              `json:"sudiRequired,omitempty"`              //
	Tags                      *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoTags                   `json:"tags,omitempty"`                      //
	UserSudiSerialNos         []string                                                                           `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                                             `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                                             `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                                             `json:"workflowName,omitempty"`              //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoFileSystemList struct {
	Freespace *int   `json:"freespace,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	Readable  *bool  `json:"readable,omitempty"`  //
	Size      *int   `json:"size,omitempty"`      //
	Type      string `json:"type,omitempty"`      //
	Writeable *bool  `json:"writeable,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfaces struct {
	IPv4Address     *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     //
	IPv6AddressList *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` //
	MacAddress      string                                                                                    `json:"macAddress,omitempty"`      //
	Name            string                                                                                    `json:"name,omitempty"`            //
	Status          string                                                                                    `json:"status,omitempty"`          //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv4Address interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoIPInterfacesIPv6AddressList interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoNeighborLinks struct {
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
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                                      `json:"createdBy,omitempty"`         //
	DiscoveryCreated  *bool                                                                                       `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                                      `json:"profileName,omitempty"`       //
	SecondaryEndpoint *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string                                                                                               `json:"certificate,omitempty"` //
	Fqdn        string                                                                                               `json:"fqdn,omitempty"`        //
	IPv4Address *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` //
	IPv6Address *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` //
	Port        *int                                                                                                 `json:"port,omitempty"`        //
	Protocol    string                                                                                               `json:"protocol,omitempty"`    //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string                                                                                                 `json:"certificate,omitempty"` //
	Fqdn        string                                                                                                 `json:"fqdn,omitempty"`        //
	IPv4Address *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` //
	IPv6Address *RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` //
	Port        *int                                                                                                   `json:"port,omitempty"`        //
	Protocol    string                                                                                                 `json:"protocol,omitempty"`    //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfo struct {
	IsFullRing             *bool                                                                                  `json:"isFullRing,omitempty"`             //
	StackMemberList        *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                                 `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows *bool                                                                                  `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       *int                                                                                   `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                                               `json:"validLicenseLevels,omitempty"`     //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string `json:"licenseLevel,omitempty"`     //
	LicenseType      string `json:"licenseType,omitempty"`      //
	MacAddress       string `json:"macAddress,omitempty"`       //
	Pid              string `json:"pid,omitempty"`              //
	Priority         *int   `json:"priority,omitempty"`         //
	Role             string `json:"role,omitempty"`             //
	SerialNumber     string `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string `json:"softwareVersion,omitempty"`  //
	StackNumber      *int   `json:"stackNumber,omitempty"`      //
	State            string `json:"state,omitempty"`            //
	SudiSerialNumber string `json:"sudiSerialNumber,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoTags interface{}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryList struct {
	Details         string                                                                          `json:"details,omitempty"`         //
	ErrorFlag       *bool                                                                           `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo *RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       *int                                                                            `json:"timestamp,omitempty"`       //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfo struct {
	AddnDetails  *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                                        `json:"name,omitempty"`         //
	TimeTaken    *int                                                                                          `json:"timeTaken,omitempty"`    //
	Type         string                                                                                        `json:"type,omitempty"`         //
	WorkItemList *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflow struct {
	TypeID         string                                                                       `json:"_id,omitempty"`            //
	AddToInventory *bool                                                                        `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                                         `json:"addedOn,omitempty"`        //
	ConfigID       string                                                                       `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                                         `json:"currTaskIdx,omitempty"`    //
	Description    string                                                                       `json:"description,omitempty"`    //
	EndTime        *int                                                                         `json:"endTime,omitempty"`        //
	ExecTime       *int                                                                         `json:"execTime,omitempty"`       //
	ImageID        string                                                                       `json:"imageId,omitempty"`        //
	InstanceType   string                                                                       `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                                         `json:"lastupdateOn,omitempty"`   //
	Name           string                                                                       `json:"name,omitempty"`           //
	StartTime      *int                                                                         `json:"startTime,omitempty"`      //
	State          string                                                                       `json:"state,omitempty"`          //
	Tasks          *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                                       `json:"tenantId,omitempty"`       //
	Type           string                                                                       `json:"type,omitempty"`           //
	UseState       string                                                                       `json:"useState,omitempty"`       //
	Version        *int                                                                         `json:"version,omitempty"`        //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                                     `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                                     `json:"endTime,omitempty"`         //
	Name            string                                                                                   `json:"name,omitempty"`            //
	StartTime       *int                                                                                     `json:"startTime,omitempty"`       //
	State           string                                                                                   `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                                     `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                                     `json:"timeTaken,omitempty"`       //
	Type            string                                                                                   `json:"type,omitempty"`            //
	WorkItemList    *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemResetWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflow struct {
	TypeID         string                                                                  `json:"_id,omitempty"`            //
	AddToInventory *bool                                                                   `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                                    `json:"addedOn,omitempty"`        //
	ConfigID       string                                                                  `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                                    `json:"currTaskIdx,omitempty"`    //
	Description    string                                                                  `json:"description,omitempty"`    //
	EndTime        *int                                                                    `json:"endTime,omitempty"`        //
	ExecTime       *int                                                                    `json:"execTime,omitempty"`       //
	ImageID        string                                                                  `json:"imageId,omitempty"`        //
	InstanceType   string                                                                  `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                                    `json:"lastupdateOn,omitempty"`   //
	Name           string                                                                  `json:"name,omitempty"`           //
	StartTime      *int                                                                    `json:"startTime,omitempty"`      //
	State          string                                                                  `json:"state,omitempty"`          //
	Tasks          *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                                  `json:"tenantId,omitempty"`       //
	Type           string                                                                  `json:"type,omitempty"`           //
	UseState       string                                                                  `json:"useState,omitempty"`       //
	Version        *int                                                                    `json:"version,omitempty"`        //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                                `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                                `json:"endTime,omitempty"`         //
	Name            string                                                                              `json:"name,omitempty"`            //
	StartTime       *int                                                                                `json:"startTime,omitempty"`       //
	State           string                                                                              `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                                `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                                `json:"timeTaken,omitempty"`       //
	Type            string                                                                              `json:"type,omitempty"`            //
	WorkItemList    *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkSystemWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflow struct {
	TypeID         string                                                            `json:"_id,omitempty"`            //
	AddToInventory *bool                                                             `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                              `json:"addedOn,omitempty"`        //
	ConfigID       string                                                            `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                              `json:"currTaskIdx,omitempty"`    //
	Description    string                                                            `json:"description,omitempty"`    //
	EndTime        *int                                                              `json:"endTime,omitempty"`        //
	ExecTime       *int                                                              `json:"execTime,omitempty"`       //
	ImageID        string                                                            `json:"imageId,omitempty"`        //
	InstanceType   string                                                            `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                              `json:"lastupdateOn,omitempty"`   //
	Name           string                                                            `json:"name,omitempty"`           //
	StartTime      *int                                                              `json:"startTime,omitempty"`      //
	State          string                                                            `json:"state,omitempty"`          //
	Tasks          *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                            `json:"tenantId,omitempty"`       //
	Type           string                                                            `json:"type,omitempty"`           //
	UseState       string                                                            `json:"useState,omitempty"`       //
	Version        *int                                                              `json:"version,omitempty"`        //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                          `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                          `json:"endTime,omitempty"`         //
	Name            string                                                                        `json:"name,omitempty"`            //
	StartTime       *int                                                                          `json:"startTime,omitempty"`       //
	State           string                                                                        `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                          `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                          `json:"timeTaken,omitempty"`       //
	Type            string                                                                        `json:"type,omitempty"`            //
	WorkItemList    *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParameters struct {
	ConfigList             *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                                           `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                                           `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                                           `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigList struct {
	ConfigID         string                                                                                           `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestItemDeviceOnboardingPnpImportDevicesInBulkWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpResetDevice struct {
	DeviceResetList *[]RequestDeviceOnboardingPnpResetDeviceDeviceResetList `json:"deviceResetList,omitempty"` //
	ProjectID       string                                                  `json:"projectId,omitempty"`       //
	WorkflowID      string                                                  `json:"workflowId,omitempty"`      //
}
type RequestDeviceOnboardingPnpResetDeviceDeviceResetList struct {
	ConfigList             *[]RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigList `json:"configList,omitempty"`             //
	DeviceID               string                                                            `json:"deviceId,omitempty"`               //
	LicenseLevel           string                                                            `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                            `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                            `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigList struct {
	ConfigID         string                                                                            `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimADeviceToASite struct {
	DeviceID   string                                                   `json:"deviceId,omitempty"`   //
	SiteID     string                                                   `json:"siteId,omitempty"`     //
	Type       string                                                   `json:"type,omitempty"`       //
	ImageInfo  *RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo  `json:"imageInfo,omitempty"`  //
	ConfigInfo *RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo `json:"configInfo,omitempty"` //
	Hostname   string                                                   `json:"hostname,omitempty"`   //
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo struct {
	ImageID string `json:"imageId,omitempty"` //
	Skip    *bool  `json:"skip,omitempty"`    //
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo struct {
	ConfigID         string                                                                     `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpPreviewConfig struct {
	DeviceID string `json:"deviceId,omitempty"` //
	SiteID   string `json:"siteId,omitempty"`   //
	Type     string `json:"type,omitempty"`     //
}
type RequestDeviceOnboardingPnpUnClaimDevice struct {
	DeviceIDList []string `json:"deviceIdList,omitempty"` //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevices struct {
	AutoSyncPeriod   *int                                                           `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                         `json:"ccoUser,omitempty"`          //
	Expiry           *int                                                           `json:"expiry,omitempty"`           //
	LastSync         *int                                                           `json:"lastSync,omitempty"`         //
	Profile          *RequestDeviceOnboardingPnpSyncVirtualAccountDevicesProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                         `json:"smartAccountId,omitempty"`   //
	SyncResult       *RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                         `json:"syncResultStr,omitempty"`    //
	SyncStartTime    *int                                                           `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                         `json:"syncStatus,omitempty"`       //
	TenantID         string                                                         `json:"tenantId,omitempty"`         //
	Token            string                                                         `json:"token,omitempty"`            //
	VirtualAccountID string                                                         `json:"virtualAccountId,omitempty"` //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault *bool  `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        *int   `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       *bool  `json:"proxy,omitempty"`       //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult struct {
	SyncList *[]RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                   `json:"syncMsg,omitempty"`  //
}
type RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}
type RequestDeviceOnboardingPnpUpdateDevice struct {
	TypeID              string                                                     `json:"_id,omitempty"`                 //
	DeviceInfo          *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo          `json:"deviceInfo,omitempty"`          //
	RunSummaryList      *[]RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList    `json:"runSummaryList,omitempty"`      //
	SystemResetWorkflow *RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow `json:"systemResetWorkflow,omitempty"` //
	SystemWorkflow      *RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow      `json:"systemWorkflow,omitempty"`      //
	TenantID            string                                                     `json:"tenantId,omitempty"`            //
	Version             *int                                                       `json:"version,omitempty"`             //
	Workflow            *RequestDeviceOnboardingPnpUpdateDeviceWorkflow            `json:"workflow,omitempty"`            //
	WorkflowParameters  *RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters  `json:"workflowParameters,omitempty"`  //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfo struct {
	AAACredentials            *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials         `json:"aaaCredentials,omitempty"`            //
	AddedOn                   *int                                                                    `json:"addedOn,omitempty"`                   //
	AddnMacAddrs              []string                                                                `json:"addnMacAddrs,omitempty"`              //
	AgentType                 string                                                                  `json:"agentType,omitempty"`                 //
	AuthStatus                string                                                                  `json:"authStatus,omitempty"`                //
	AuthenticatedSudiSerialNo string                                                                  `json:"authenticatedSudiSerialNo,omitempty"` //
	CapabilitiesSupported     []string                                                                `json:"capabilitiesSupported,omitempty"`     //
	CmState                   string                                                                  `json:"cmState,omitempty"`                   //
	Description               string                                                                  `json:"description,omitempty"`               //
	DeviceSudiSerialNos       []string                                                                `json:"deviceSudiSerialNos,omitempty"`       //
	DeviceType                string                                                                  `json:"deviceType,omitempty"`                //
	FeaturesSupported         []string                                                                `json:"featuresSupported,omitempty"`         //
	FileSystemList            *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList       `json:"fileSystemList,omitempty"`            //
	FirstContact              *int                                                                    `json:"firstContact,omitempty"`              //
	Hostname                  string                                                                  `json:"hostname,omitempty"`                  //
	HTTPHeaders               *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders          `json:"httpHeaders,omitempty"`               //
	ImageFile                 string                                                                  `json:"imageFile,omitempty"`                 //
	ImageVersion              string                                                                  `json:"imageVersion,omitempty"`              //
	IPInterfaces              *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces         `json:"ipInterfaces,omitempty"`              //
	LastContact               *int                                                                    `json:"lastContact,omitempty"`               //
	LastSyncTime              *int                                                                    `json:"lastSyncTime,omitempty"`              //
	LastUpdateOn              *int                                                                    `json:"lastUpdateOn,omitempty"`              //
	Location                  *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation               `json:"location,omitempty"`                  //
	MacAddress                string                                                                  `json:"macAddress,omitempty"`                //
	Mode                      string                                                                  `json:"mode,omitempty"`                      //
	Name                      string                                                                  `json:"name,omitempty"`                      //
	NeighborLinks             *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks        `json:"neighborLinks,omitempty"`             //
	OnbState                  string                                                                  `json:"onbState,omitempty"`                  //
	Pid                       string                                                                  `json:"pid,omitempty"`                       //
	PnpProfileList            *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList       `json:"pnpProfileList,omitempty"`            //
	PopulateInventory         *bool                                                                   `json:"populateInventory,omitempty"`         //
	PreWorkflowCliOuputs      *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs `json:"preWorkflowCliOuputs,omitempty"`      //
	ProjectID                 string                                                                  `json:"projectId,omitempty"`                 //
	ProjectName               string                                                                  `json:"projectName,omitempty"`               //
	ReloadRequested           *bool                                                                   `json:"reloadRequested,omitempty"`           //
	SerialNumber              string                                                                  `json:"serialNumber,omitempty"`              //
	SmartAccountID            string                                                                  `json:"smartAccountId,omitempty"`            //
	Source                    string                                                                  `json:"source,omitempty"`                    //
	Stack                     *bool                                                                   `json:"stack,omitempty"`                     //
	StackInfo                 *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo              `json:"stackInfo,omitempty"`                 //
	State                     string                                                                  `json:"state,omitempty"`                     //
	SudiRequired              *bool                                                                   `json:"sudiRequired,omitempty"`              //
	Tags                      *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags                   `json:"tags,omitempty"`                      //
	UserSudiSerialNos         []string                                                                `json:"userSudiSerialNos,omitempty"`         //
	VirtualAccountID          string                                                                  `json:"virtualAccountId,omitempty"`          //
	WorkflowID                string                                                                  `json:"workflowId,omitempty"`                //
	WorkflowName              string                                                                  `json:"workflowName,omitempty"`              //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoFileSystemList struct {
	Freespace *int   `json:"freespace,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	Readable  *bool  `json:"readable,omitempty"`  //
	Size      *int   `json:"size,omitempty"`      //
	Type      string `json:"type,omitempty"`      //
	Writeable *bool  `json:"writeable,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoHTTPHeaders struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfaces struct {
	IPv4Address     *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address       `json:"ipv4Address,omitempty"`     //
	IPv6AddressList *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList `json:"ipv6AddressList,omitempty"` //
	MacAddress      string                                                                         `json:"macAddress,omitempty"`      //
	Name            string                                                                         `json:"name,omitempty"`            //
	Status          string                                                                         `json:"status,omitempty"`          //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv4Address interface{}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoIPInterfacesIPv6AddressList interface{}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoLocation struct {
	Address   string `json:"address,omitempty"`   //
	Altitude  string `json:"altitude,omitempty"`  //
	Latitude  string `json:"latitude,omitempty"`  //
	Longitude string `json:"longitude,omitempty"` //
	SiteID    string `json:"siteId,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoNeighborLinks struct {
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
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileList struct {
	CreatedBy         string                                                                           `json:"createdBy,omitempty"`         //
	DiscoveryCreated  *bool                                                                            `json:"discoveryCreated,omitempty"`  //
	PrimaryEndpoint   *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint   `json:"primaryEndpoint,omitempty"`   //
	ProfileName       string                                                                           `json:"profileName,omitempty"`       //
	SecondaryEndpoint *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint `json:"secondaryEndpoint,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpoint struct {
	Certificate string                                                                                    `json:"certificate,omitempty"` //
	Fqdn        string                                                                                    `json:"fqdn,omitempty"`        //
	IPv4Address *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address `json:"ipv4Address,omitempty"` //
	IPv6Address *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address `json:"ipv6Address,omitempty"` //
	Port        *int                                                                                      `json:"port,omitempty"`        //
	Protocol    string                                                                                    `json:"protocol,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv4Address interface{}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListPrimaryEndpointIPv6Address interface{}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpoint struct {
	Certificate string                                                                                      `json:"certificate,omitempty"` //
	Fqdn        string                                                                                      `json:"fqdn,omitempty"`        //
	IPv4Address *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address `json:"ipv4Address,omitempty"` //
	IPv6Address *RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address `json:"ipv6Address,omitempty"` //
	Port        *int                                                                                        `json:"port,omitempty"`        //
	Protocol    string                                                                                      `json:"protocol,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv4Address interface{}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPnpProfileListSecondaryEndpointIPv6Address interface{}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoPreWorkflowCliOuputs struct {
	Cli       string `json:"cli,omitempty"`       //
	CliOutput string `json:"cliOutput,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfo struct {
	IsFullRing             *bool                                                                       `json:"isFullRing,omitempty"`             //
	StackMemberList        *[]RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList `json:"stackMemberList,omitempty"`        //
	StackRingProtocol      string                                                                      `json:"stackRingProtocol,omitempty"`      //
	SupportsStackWorkflows *bool                                                                       `json:"supportsStackWorkflows,omitempty"` //
	TotalMemberCount       *int                                                                        `json:"totalMemberCount,omitempty"`       //
	ValidLicenseLevels     []string                                                                    `json:"validLicenseLevels,omitempty"`     //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoStackInfoStackMemberList struct {
	HardwareVersion  string `json:"hardwareVersion,omitempty"`  //
	LicenseLevel     string `json:"licenseLevel,omitempty"`     //
	LicenseType      string `json:"licenseType,omitempty"`      //
	MacAddress       string `json:"macAddress,omitempty"`       //
	Pid              string `json:"pid,omitempty"`              //
	Priority         *int   `json:"priority,omitempty"`         //
	Role             string `json:"role,omitempty"`             //
	SerialNumber     string `json:"serialNumber,omitempty"`     //
	SoftwareVersion  string `json:"softwareVersion,omitempty"`  //
	StackNumber      *int   `json:"stackNumber,omitempty"`      //
	State            string `json:"state,omitempty"`            //
	SudiSerialNumber string `json:"sudiSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceDeviceInfoTags interface{}
type RequestDeviceOnboardingPnpUpdateDeviceRunSummaryList struct {
	Details         string                                                               `json:"details,omitempty"`         //
	ErrorFlag       *bool                                                                `json:"errorFlag,omitempty"`       //
	HistoryTaskInfo *RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo `json:"historyTaskInfo,omitempty"` //
	Timestamp       *int                                                                 `json:"timestamp,omitempty"`       //
}
type RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfo struct {
	AddnDetails  *[]RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails  `json:"addnDetails,omitempty"`  //
	Name         string                                                                             `json:"name,omitempty"`         //
	TimeTaken    *int                                                                               `json:"timeTaken,omitempty"`    //
	Type         string                                                                             `json:"type,omitempty"`         //
	WorkItemList *[]RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList `json:"workItemList,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoAddnDetails struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceRunSummaryListHistoryTaskInfoWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflow struct {
	TypeID         string                                                            `json:"_id,omitempty"`            //
	AddToInventory *bool                                                             `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                              `json:"addedOn,omitempty"`        //
	ConfigID       string                                                            `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                              `json:"currTaskIdx,omitempty"`    //
	Description    string                                                            `json:"description,omitempty"`    //
	EndTime        *int                                                              `json:"endTime,omitempty"`        //
	ExecTime       *int                                                              `json:"execTime,omitempty"`       //
	ImageID        string                                                            `json:"imageId,omitempty"`        //
	InstanceType   string                                                            `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                              `json:"lastupdateOn,omitempty"`   //
	Name           string                                                            `json:"name,omitempty"`           //
	StartTime      *int                                                              `json:"startTime,omitempty"`      //
	State          string                                                            `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                            `json:"tenantId,omitempty"`       //
	Type           string                                                            `json:"type,omitempty"`           //
	UseState       string                                                            `json:"useState,omitempty"`       //
	Version        *int                                                              `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                          `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                          `json:"endTime,omitempty"`         //
	Name            string                                                                        `json:"name,omitempty"`            //
	StartTime       *int                                                                          `json:"startTime,omitempty"`       //
	State           string                                                                        `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                          `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                          `json:"timeTaken,omitempty"`       //
	Type            string                                                                        `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateDeviceSystemResetWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflow struct {
	TypeID         string                                                       `json:"_id,omitempty"`            //
	AddToInventory *bool                                                        `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                         `json:"addedOn,omitempty"`        //
	ConfigID       string                                                       `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                         `json:"currTaskIdx,omitempty"`    //
	Description    string                                                       `json:"description,omitempty"`    //
	EndTime        *int                                                         `json:"endTime,omitempty"`        //
	ExecTime       *int                                                         `json:"execTime,omitempty"`       //
	ImageID        string                                                       `json:"imageId,omitempty"`        //
	InstanceType   string                                                       `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                         `json:"lastupdateOn,omitempty"`   //
	Name           string                                                       `json:"name,omitempty"`           //
	StartTime      *int                                                         `json:"startTime,omitempty"`      //
	State          string                                                       `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                       `json:"tenantId,omitempty"`       //
	Type           string                                                       `json:"type,omitempty"`           //
	UseState       string                                                       `json:"useState,omitempty"`       //
	Version        *int                                                         `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasks struct {
	CurrWorkItemIDx *int                                                                     `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                                     `json:"endTime,omitempty"`         //
	Name            string                                                                   `json:"name,omitempty"`            //
	StartTime       *int                                                                     `json:"startTime,omitempty"`       //
	State           string                                                                   `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                                     `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                                     `json:"timeTaken,omitempty"`       //
	Type            string                                                                   `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateDeviceSystemWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceWorkflow struct {
	TypeID         string                                                 `json:"_id,omitempty"`            //
	AddToInventory *bool                                                  `json:"addToInventory,omitempty"` //
	AddedOn        *int                                                   `json:"addedOn,omitempty"`        //
	ConfigID       string                                                 `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                                   `json:"currTaskIdx,omitempty"`    //
	Description    string                                                 `json:"description,omitempty"`    //
	EndTime        *int                                                   `json:"endTime,omitempty"`        //
	ExecTime       *int                                                   `json:"execTime,omitempty"`       //
	ImageID        string                                                 `json:"imageId,omitempty"`        //
	InstanceType   string                                                 `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                                   `json:"lastupdateOn,omitempty"`   //
	Name           string                                                 `json:"name,omitempty"`           //
	StartTime      *int                                                   `json:"startTime,omitempty"`      //
	State          string                                                 `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                                 `json:"tenantId,omitempty"`       //
	Type           string                                                 `json:"type,omitempty"`           //
	UseState       string                                                 `json:"useState,omitempty"`       //
	Version        *int                                                   `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasks struct {
	CurrWorkItemIDx *int                                                               `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                               `json:"endTime,omitempty"`         //
	Name            string                                                             `json:"name,omitempty"`            //
	StartTime       *int                                                               `json:"startTime,omitempty"`       //
	State           string                                                             `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                               `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                               `json:"timeTaken,omitempty"`       //
	Type            string                                                             `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateDeviceWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceWorkflowParameters struct {
	ConfigList             *[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList `json:"configList,omitempty"`             //
	LicenseLevel           string                                                                `json:"licenseLevel,omitempty"`           //
	LicenseType            string                                                                `json:"licenseType,omitempty"`            //
	TopOfStackSerialNumber string                                                                `json:"topOfStackSerialNumber,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigList struct {
	ConfigID         string                                                                                `json:"configId,omitempty"`         //
	ConfigParameters *[]RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters `json:"configParameters,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateDeviceWorkflowParametersConfigListConfigParameters struct {
	Key   string `json:"key,omitempty"`   //
	Value string `json:"value,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettings struct {
	TypeID          string                                                              `json:"_id,omitempty"`             //
	AAACredentials  *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsAAACredentials    `json:"aaaCredentials,omitempty"`  //
	AcceptEula      *bool                                                               `json:"acceptEula,omitempty"`      //
	DefaultProfile  *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsDefaultProfile    `json:"defaultProfile,omitempty"`  //
	SavaMappingList *[]RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList `json:"savaMappingList,omitempty"` //
	TaskTimeOuts    *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsTaskTimeOuts      `json:"taskTimeOuts,omitempty"`    //
	TenantID        string                                                              `json:"tenantId,omitempty"`        //
	Version         *int                                                                `json:"version,omitempty"`         //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsAAACredentials struct {
	Password string `json:"password,omitempty"` //
	Username string `json:"username,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsDefaultProfile struct {
	Cert          string   `json:"cert,omitempty"`          //
	FqdnAddresses []string `json:"fqdnAddresses,omitempty"` //
	IPAddresses   []string `json:"ipAddresses,omitempty"`   //
	Port          *int     `json:"port,omitempty"`          //
	Proxy         *bool    `json:"proxy,omitempty"`         //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingList struct {
	AutoSyncPeriod   *int                                                                        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                                      `json:"ccoUser,omitempty"`          //
	Expiry           *int                                                                        `json:"expiry,omitempty"`           //
	LastSync         *int                                                                        `json:"lastSync,omitempty"`         //
	Profile          *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                                      `json:"smartAccountId,omitempty"`   //
	SyncResult       *RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                                      `json:"syncResultStr,omitempty"`    //
	SyncStartTime    *int                                                                        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                                      `json:"syncStatus,omitempty"`       //
	TenantID         string                                                                      `json:"tenantId,omitempty"`         //
	Token            string                                                                      `json:"token,omitempty"`            //
	VirtualAccountID string                                                                      `json:"virtualAccountId,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault *bool  `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        *int   `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       *bool  `json:"proxy,omitempty"`       //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResult struct {
	SyncList *[]RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                                `json:"syncMsg,omitempty"`  //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsSavaMappingListSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}
type RequestDeviceOnboardingPnpUpdatePnpGlobalSettingsTaskTimeOuts struct {
	ConfigTimeOut        *int `json:"configTimeOut,omitempty"`        //
	GeneralTimeOut       *int `json:"generalTimeOut,omitempty"`       //
	ImageDownloadTimeOut *int `json:"imageDownloadTimeOut,omitempty"` //
}
type RequestDeviceOnboardingPnpAddVirtualAccount struct {
	AutoSyncPeriod   *int                                                   `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                 `json:"ccoUser,omitempty"`          //
	Expiry           *int                                                   `json:"expiry,omitempty"`           //
	LastSync         *int                                                   `json:"lastSync,omitempty"`         //
	Profile          *RequestDeviceOnboardingPnpAddVirtualAccountProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                 `json:"smartAccountId,omitempty"`   //
	SyncResult       *RequestDeviceOnboardingPnpAddVirtualAccountSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                 `json:"syncResultStr,omitempty"`    //
	SyncStartTime    *int                                                   `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                 `json:"syncStatus,omitempty"`       //
	TenantID         string                                                 `json:"tenantId,omitempty"`         //
	Token            string                                                 `json:"token,omitempty"`            //
	VirtualAccountID string                                                 `json:"virtualAccountId,omitempty"` //
}
type RequestDeviceOnboardingPnpAddVirtualAccountProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault *bool  `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        *int   `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       *bool  `json:"proxy,omitempty"`       //
}
type RequestDeviceOnboardingPnpAddVirtualAccountSyncResult struct {
	SyncList *[]RequestDeviceOnboardingPnpAddVirtualAccountSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                           `json:"syncMsg,omitempty"`  //
}
type RequestDeviceOnboardingPnpAddVirtualAccountSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}
type RequestDeviceOnboardingPnpUpdatePnpServerProfile struct {
	AutoSyncPeriod   *int                                                        `json:"autoSyncPeriod,omitempty"`   //
	CcoUser          string                                                      `json:"ccoUser,omitempty"`          //
	Expiry           *int                                                        `json:"expiry,omitempty"`           //
	LastSync         *int                                                        `json:"lastSync,omitempty"`         //
	Profile          *RequestDeviceOnboardingPnpUpdatePnpServerProfileProfile    `json:"profile,omitempty"`          //
	SmartAccountID   string                                                      `json:"smartAccountId,omitempty"`   //
	SyncResult       *RequestDeviceOnboardingPnpUpdatePnpServerProfileSyncResult `json:"syncResult,omitempty"`       //
	SyncResultStr    string                                                      `json:"syncResultStr,omitempty"`    //
	SyncStartTime    *int                                                        `json:"syncStartTime,omitempty"`    //
	SyncStatus       string                                                      `json:"syncStatus,omitempty"`       //
	TenantID         string                                                      `json:"tenantId,omitempty"`         //
	Token            string                                                      `json:"token,omitempty"`            //
	VirtualAccountID string                                                      `json:"virtualAccountId,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdatePnpServerProfileProfile struct {
	AddressFqdn string `json:"addressFqdn,omitempty"` //
	AddressIPV4 string `json:"addressIpV4,omitempty"` //
	Cert        string `json:"cert,omitempty"`        //
	MakeDefault *bool  `json:"makeDefault,omitempty"` //
	Name        string `json:"name,omitempty"`        //
	Port        *int   `json:"port,omitempty"`        //
	ProfileID   string `json:"profileId,omitempty"`   //
	Proxy       *bool  `json:"proxy,omitempty"`       //
}
type RequestDeviceOnboardingPnpUpdatePnpServerProfileSyncResult struct {
	SyncList *[]RequestDeviceOnboardingPnpUpdatePnpServerProfileSyncResultSyncList `json:"syncList,omitempty"` //
	SyncMsg  string                                                                `json:"syncMsg,omitempty"`  //
}
type RequestDeviceOnboardingPnpUpdatePnpServerProfileSyncResultSyncList struct {
	DeviceSnList []string `json:"deviceSnList,omitempty"` //
	SyncType     string   `json:"syncType,omitempty"`     //
}
type RequestDeviceOnboardingPnpAddAWorkflow struct {
	TypeID         string                                         `json:"_id,omitempty"`            //
	AddToInventory *bool                                          `json:"addToInventory,omitempty"` //
	AddedOn        *int                                           `json:"addedOn,omitempty"`        //
	ConfigID       string                                         `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                           `json:"currTaskIdx,omitempty"`    //
	Description    string                                         `json:"description,omitempty"`    //
	EndTime        *int                                           `json:"endTime,omitempty"`        //
	ExecTime       *int                                           `json:"execTime,omitempty"`       //
	ImageID        string                                         `json:"imageId,omitempty"`        //
	InstanceType   string                                         `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                           `json:"lastupdateOn,omitempty"`   //
	Name           string                                         `json:"name,omitempty"`           //
	StartTime      *int                                           `json:"startTime,omitempty"`      //
	State          string                                         `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpAddAWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                         `json:"tenantId,omitempty"`       //
	Type           string                                         `json:"type,omitempty"`           //
	UseState       string                                         `json:"useState,omitempty"`       //
	Version        *int                                           `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpAddAWorkflowTasks struct {
	CurrWorkItemIDx *int                                                       `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                       `json:"endTime,omitempty"`         //
	Name            string                                                     `json:"name,omitempty"`            //
	StartTime       *int                                                       `json:"startTime,omitempty"`       //
	State           string                                                     `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                       `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                       `json:"timeTaken,omitempty"`       //
	Type            string                                                     `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpAddAWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}
type RequestDeviceOnboardingPnpUpdateWorkflow struct {
	TypeID         string                                           `json:"_id,omitempty"`            //
	AddToInventory *bool                                            `json:"addToInventory,omitempty"` //
	AddedOn        *int                                             `json:"addedOn,omitempty"`        //
	ConfigID       string                                           `json:"configId,omitempty"`       //
	CurrTaskIDx    *int                                             `json:"currTaskIdx,omitempty"`    //
	Description    string                                           `json:"description,omitempty"`    //
	EndTime        *int                                             `json:"endTime,omitempty"`        //
	ExecTime       *int                                             `json:"execTime,omitempty"`       //
	ImageID        string                                           `json:"imageId,omitempty"`        //
	InstanceType   string                                           `json:"instanceType,omitempty"`   //
	LastupdateOn   *int                                             `json:"lastupdateOn,omitempty"`   //
	Name           string                                           `json:"name,omitempty"`           //
	StartTime      *int                                             `json:"startTime,omitempty"`      //
	State          string                                           `json:"state,omitempty"`          //
	Tasks          *[]RequestDeviceOnboardingPnpUpdateWorkflowTasks `json:"tasks,omitempty"`          //
	TenantID       string                                           `json:"tenantId,omitempty"`       //
	Type           string                                           `json:"type,omitempty"`           //
	UseState       string                                           `json:"useState,omitempty"`       //
	Version        *int                                             `json:"version,omitempty"`        //
}
type RequestDeviceOnboardingPnpUpdateWorkflowTasks struct {
	CurrWorkItemIDx *int                                                         `json:"currWorkItemIdx,omitempty"` //
	EndTime         *int                                                         `json:"endTime,omitempty"`         //
	Name            string                                                       `json:"name,omitempty"`            //
	StartTime       *int                                                         `json:"startTime,omitempty"`       //
	State           string                                                       `json:"state,omitempty"`           //
	TaskSeqNo       *int                                                         `json:"taskSeqNo,omitempty"`       //
	TimeTaken       *int                                                         `json:"timeTaken,omitempty"`       //
	Type            string                                                       `json:"type,omitempty"`            //
	WorkItemList    *[]RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList `json:"workItemList,omitempty"`    //
}
type RequestDeviceOnboardingPnpUpdateWorkflowTasksWorkItemList struct {
	Command   string `json:"command,omitempty"`   //
	EndTime   *int   `json:"endTime,omitempty"`   //
	OutputStr string `json:"outputStr,omitempty"` //
	StartTime *int   `json:"startTime,omitempty"` //
	State     string `json:"state,omitempty"`     //
	TimeTaken *int   `json:"timeTaken,omitempty"` //
}

//GetDeviceList2 Get Device list - e6b3-db80-46c9-9654
/* Returns list of devices based on filter crieteria. If a limit is not specified, it will default to return 50 devices. Pagination and sorting are also supported by this endpoint


@param GetDeviceList2QueryParams Filtering parameter
*/
func (s *DeviceOnboardingPnpService) GetDeviceList2(GetDeviceList2QueryParams *GetDeviceList2QueryParams) (*ResponseDeviceOnboardingPnpGetDeviceList2, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device"

	queryString, _ := query.Values(GetDeviceList2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetDeviceList2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceList2")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceList2)
	return result, response, err

}

//GetDeviceCount Get Device Count - d9a1-fa9c-4068-b23c
/* Returns the device count based on filter criteria. This is useful for pagination


@param GetDeviceCountQueryParams Filtering parameter
*/
func (s *DeviceOnboardingPnpService) GetDeviceCount(GetDeviceCountQueryParams *GetDeviceCountQueryParams) (*ResponseDeviceOnboardingPnpGetDeviceCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/count"

	queryString, _ := query.Values(GetDeviceCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetDeviceCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceCount")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceCount)
	return result, response, err

}

//GetDeviceHistory Get Device History - f093-1967-4049-a7d4
/* Returns history for a specific device. Serial number is a required parameter


@param GetDeviceHistoryQueryParams Filtering parameter
*/
func (s *DeviceOnboardingPnpService) GetDeviceHistory(GetDeviceHistoryQueryParams *GetDeviceHistoryQueryParams) (*ResponseDeviceOnboardingPnpGetDeviceHistory, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/history"

	queryString, _ := query.Values(GetDeviceHistoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetDeviceHistory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceHistory")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceHistory)
	return result, response, err

}

//GetSyncResultForVirtualAccount Get Sync Result for Virtual Account - 0a9c-9884-45cb-91c8
/* Returns the summary of devices synced from the given smart account & virtual account with PnP


@param domain domain path parameter. Smart Account Domain

@param name name path parameter. Virtual Account Name

*/
func (s *DeviceOnboardingPnpService) GetSyncResultForVirtualAccount(domain string, name string) (*ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccount, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/sacct/{domain}/vacct/{name}/sync-result"
	path = strings.Replace(path, "{domain}", fmt.Sprintf("%v", domain), -1)
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSyncResultForVirtualAccount")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetSyncResultForVirtualAccount)
	return result, response, err

}

//GetDeviceByID Get Device by Id - bab6-c9e5-4408-85cc
/* Returns device details specified by device id


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) GetDeviceByID(id string) (*ResponseDeviceOnboardingPnpGetDeviceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetDeviceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceById")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetDeviceByID)
	return result, response, err

}

//GetPnpGlobalSettings Get PnP global settings - 7e92-f9eb-46db-8320
/* Returns global PnP settings of the user


 */
func (s *DeviceOnboardingPnpService) GetPnpGlobalSettings() (*ResponseDeviceOnboardingPnpGetPnpGlobalSettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetPnpGlobalSettings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetPnpGlobalSettings")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetPnpGlobalSettings)
	return result, response, err

}

//GetSmartAccountList Get Smart Account List - 3cb2-4acb-486b-89d2
/* Returns the list of Smart Account domains


 */
func (s *DeviceOnboardingPnpService) GetSmartAccountList() (*ResponseDeviceOnboardingPnpGetSmartAccountList, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetSmartAccountList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSmartAccountList")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetSmartAccountList)
	return result, response, err

}

//GetVirtualAccountList Get Virtual Account List - 70a4-79a6-462a-9496
/* Returns list of virtual accounts associated with the specified smart account


@param domain domain path parameter. Smart Account Domain

*/
func (s *DeviceOnboardingPnpService) GetVirtualAccountList(domain string) (*ResponseDeviceOnboardingPnpGetVirtualAccountList, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/sacct/{domain}/vacct"
	path = strings.Replace(path, "{domain}", fmt.Sprintf("%v", domain), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetVirtualAccountList{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetVirtualAccountList")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetVirtualAccountList)
	return result, response, err

}

//GetWorkflows Get Workflows - aeb4-dad0-4a99-bbe3
/* Returns the list of workflows based on filter criteria. If a limit is not specified, it will default to return 50 workflows. Pagination and sorting are also supported by this endpoint


@param GetWorkflowsQueryParams Filtering parameter
*/
func (s *DeviceOnboardingPnpService) GetWorkflows(GetWorkflowsQueryParams *GetWorkflowsQueryParams) (*ResponseDeviceOnboardingPnpGetWorkflows, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

	queryString, _ := query.Values(GetWorkflowsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetWorkflows{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetWorkflows")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetWorkflows)
	return result, response, err

}

//GetWorkflowCount Get Workflow Count - 7989-f868-46fa-af99
/* Returns the workflow count


@param GetWorkflowCountQueryParams Filtering parameter
*/
func (s *DeviceOnboardingPnpService) GetWorkflowCount(GetWorkflowCountQueryParams *GetWorkflowCountQueryParams) (*ResponseDeviceOnboardingPnpGetWorkflowCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/count"

	queryString, _ := query.Values(GetWorkflowCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpGetWorkflowCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetWorkflowCount")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetWorkflowCount)
	return result, response, err

}

//GetWorkflowByID Get Workflow by Id - 80ac-b88e-4ac9-ac6d
/* Returns a workflow specified by id


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) GetWorkflowByID(id string) (*ResponseDeviceOnboardingPnpGetWorkflowByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpGetWorkflowByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetWorkflowById")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpGetWorkflowByID)
	return result, response, err

}

//AddDevice Add Device - f3b2-6b55-44ca-bab9
/* Adds a device to the PnP database.


 */
func (s *DeviceOnboardingPnpService) AddDevice(requestDeviceOnboardingPnpAddDevice *RequestDeviceOnboardingPnpAddDevice) (*ResponseDeviceOnboardingPnpAddDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAddDevice).
		SetResult(&ResponseDeviceOnboardingPnpAddDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddDevice")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAddDevice)
	return result, response, err

}

//ClaimDevice Claim Device - d8a6-1997-4a8a-8c48
/* Claims one of more devices with specified workflow


 */
func (s *DeviceOnboardingPnpService) ClaimDevice(requestDeviceOnboardingPnpClaimDevice *RequestDeviceOnboardingPnpClaimDevice) (*ResponseDeviceOnboardingPnpClaimDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/claim"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpClaimDevice).
		SetResult(&ResponseDeviceOnboardingPnpClaimDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ClaimDevice")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpClaimDevice)
	return result, response, err

}

//ImportDevicesInBulk Import Devices in bulk - 21a6-db25-4029-8f55
/* Add devices to PnP in bulk


 */
func (s *DeviceOnboardingPnpService) ImportDevicesInBulk(requestDeviceOnboardingPnpImportDevicesInBulk *RequestDeviceOnboardingPnpImportDevicesInBulk) (*ResponseDeviceOnboardingPnpImportDevicesInBulk, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/import"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpImportDevicesInBulk).
		SetResult(&ResponseDeviceOnboardingPnpImportDevicesInBulk{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportDevicesInBulk")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpImportDevicesInBulk)
	return result, response, err

}

//ResetDevice Reset Device - 9e85-7b5a-4a0b-bcdb
/* Recovers a device from a Workflow Execution Error state


 */
func (s *DeviceOnboardingPnpService) ResetDevice(requestDeviceOnboardingPnpResetDevice *RequestDeviceOnboardingPnpResetDevice) (*ResponseDeviceOnboardingPnpResetDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/reset"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpResetDevice).
		SetResult(&ResponseDeviceOnboardingPnpResetDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ResetDevice")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpResetDevice)
	return result, response, err

}

//ClaimADeviceToASite Claim a Device to a Site - 5889-fb84-4939-a13b
/* Claim a device based on DNA-C Site based design process. Different parameters are required for different device platforms.


 */
func (s *DeviceOnboardingPnpService) ClaimADeviceToASite(requestDeviceOnboardingPnpClaimADeviceToASite *RequestDeviceOnboardingPnpClaimADeviceToASite) (*ResponseDeviceOnboardingPnpClaimADeviceToASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/site-claim"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpClaimADeviceToASite).
		SetResult(&ResponseDeviceOnboardingPnpClaimADeviceToASite{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ClaimADeviceToASite")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpClaimADeviceToASite)
	return result, response, err

}

//PreviewConfig Preview Config - cf94-1823-4d9a-b37e
/* Triggers a preview for site-based Day 0 Configuration


 */
func (s *DeviceOnboardingPnpService) PreviewConfig(requestDeviceOnboardingPnpPreviewConfig *RequestDeviceOnboardingPnpPreviewConfig) (*ResponseDeviceOnboardingPnpPreviewConfig, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/site-config-preview"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpPreviewConfig).
		SetResult(&ResponseDeviceOnboardingPnpPreviewConfig{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PreviewConfig")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpPreviewConfig)
	return result, response, err

}

//UnClaimDevice Un-Claim Device - 0b83-6b7b-4b6a-9fd5
/* Un-Claims one of more devices with specified workflow


 */
func (s *DeviceOnboardingPnpService) UnClaimDevice(requestDeviceOnboardingPnpUnClaimDevice *RequestDeviceOnboardingPnpUnClaimDevice) (*ResponseDeviceOnboardingPnpUnClaimDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/unclaim"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUnClaimDevice).
		SetResult(&ResponseDeviceOnboardingPnpUnClaimDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UnClaimDevice")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUnClaimDevice)
	return result, response, err

}

//SyncVirtualAccountDevices Sync Virtual Account Devices - a4b6-c87a-4ffb-9efa
/* Synchronizes the device info from the given smart account & virtual account with the PnP database. The response payload returns a list of synced devices


 */
func (s *DeviceOnboardingPnpService) SyncVirtualAccountDevices(requestDeviceOnboardingPnpSyncVirtualAccountDevices *RequestDeviceOnboardingPnpSyncVirtualAccountDevices) (*ResponseDeviceOnboardingPnpSyncVirtualAccountDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/vacct-sync"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpSyncVirtualAccountDevices).
		SetResult(&ResponseDeviceOnboardingPnpSyncVirtualAccountDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SyncVirtualAccountDevices")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpSyncVirtualAccountDevices)
	return result, response, err

}

//AddVirtualAccount Add Virtual Account - 1e96-2af3-45b8-b59f
/* Registers a Smart Account, Virtual Account and the relevant server profile info with the PnP System & database. The devices present in the registered virtual account are synced with the PnP database as well. The response payload returns the new profile


 */
func (s *DeviceOnboardingPnpService) AddVirtualAccount(requestDeviceOnboardingPnpAddVirtualAccount *RequestDeviceOnboardingPnpAddVirtualAccount) (*ResponseDeviceOnboardingPnpAddVirtualAccount, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAddVirtualAccount).
		SetResult(&ResponseDeviceOnboardingPnpAddVirtualAccount{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddVirtualAccount")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAddVirtualAccount)
	return result, response, err

}

//AddAWorkflow Add a Workflow - 848b-5a7b-4f9b-8c12
/* Adds a PnP Workflow along with the relevant tasks in the workflow into the PnP database


 */
func (s *DeviceOnboardingPnpService) AddAWorkflow(requestDeviceOnboardingPnpAddAWorkflow *RequestDeviceOnboardingPnpAddAWorkflow) (*ResponseDeviceOnboardingPnpAddAWorkflow, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpAddAWorkflow).
		SetResult(&ResponseDeviceOnboardingPnpAddAWorkflow{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation AddAWorkflow")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpAddAWorkflow)
	return result, response, err

}

//UpdateDevice Update Device - 09b0-f9ce-4239-ae10
/* Updates device details specified by device id in PnP database


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) UpdateDevice(id string, requestDeviceOnboardingPnpUpdateDevice *RequestDeviceOnboardingPnpUpdateDevice) (*ResponseDeviceOnboardingPnpUpdateDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdateDevice).
		SetResult(&ResponseDeviceOnboardingPnpUpdateDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateDevice")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdateDevice)
	return result, response, err

}

//UpdatePnpGlobalSettings Update PnP global settings - 8da0-3919-4708-8a5a
/* Updates the user's list of global PnP settings


 */
func (s *DeviceOnboardingPnpService) UpdatePnpGlobalSettings(requestDeviceOnboardingPnpUpdatePnPGlobalSettings *RequestDeviceOnboardingPnpUpdatePnpGlobalSettings) (*ResponseDeviceOnboardingPnpUpdatePnpGlobalSettings, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdatePnPGlobalSettings).
		SetResult(&ResponseDeviceOnboardingPnpUpdatePnpGlobalSettings{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatePnpGlobalSettings")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdatePnpGlobalSettings)
	return result, response, err

}

//UpdatePnpServerProfile Update PnP Server Profile - 6f98-19e8-4178-870c
/* Updates the PnP Server profile in a registered Virtual Account in the PnP database. The response payload returns the updated smart & virtual account info


 */
func (s *DeviceOnboardingPnpService) UpdatePnpServerProfile(requestDeviceOnboardingPnpUpdatePnPServerProfile *RequestDeviceOnboardingPnpUpdatePnpServerProfile) (*ResponseDeviceOnboardingPnpUpdatePnpServerProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/savacct"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdatePnPServerProfile).
		SetResult(&ResponseDeviceOnboardingPnpUpdatePnpServerProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatePnpServerProfile")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdatePnpServerProfile)
	return result, response, err

}

//UpdateWorkflow Update Workflow - 3086-c962-4f49-8b85
/* Updates an existing workflow


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) UpdateWorkflow(id string, requestDeviceOnboardingPnpUpdateWorkflow *RequestDeviceOnboardingPnpUpdateWorkflow) (*ResponseDeviceOnboardingPnpUpdateWorkflow, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceOnboardingPnpUpdateWorkflow).
		SetResult(&ResponseDeviceOnboardingPnpUpdateWorkflow{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateWorkflow")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpUpdateWorkflow)
	return result, response, err

}

//DeleteDeviceByIDFromPnp Delete Device by Id from PnP - cdab-9b47-4899-ae06
/* Deletes specified device from PnP database


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) DeleteDeviceByIDFromPnp(id string) (*ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnp, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-device/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnp{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteDeviceByIdFromPnp")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpDeleteDeviceByIDFromPnp)
	return result, response, err

}

//DeregisterVirtualAccount Deregister Virtual Account - 2499-e9ad-42e8-ae5b
/* Deregisters the specified smart account & virtual account info and the associated device information from the PnP System & database. The devices associated with the deregistered virtual account are removed from the PnP database as well. The response payload contains the deregistered smart & virtual account information


@param DeregisterVirtualAccountQueryParams Filtering parameter
*/
func (s *DeviceOnboardingPnpService) DeregisterVirtualAccount(DeregisterVirtualAccountQueryParams *DeregisterVirtualAccountQueryParams) (*ResponseDeviceOnboardingPnpDeregisterVirtualAccount, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-settings/vacct"

	queryString, _ := query.Values(DeregisterVirtualAccountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceOnboardingPnpDeregisterVirtualAccount{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeregisterVirtualAccount")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpDeregisterVirtualAccount)
	return result, response, err

}

//DeleteWorkflowByID Delete Workflow By Id - af8d-7b0e-470b-8ae2
/* Deletes a workflow specified by id


@param id id path parameter.
*/
func (s *DeviceOnboardingPnpService) DeleteWorkflowByID(id string) (*ResponseDeviceOnboardingPnpDeleteWorkflowByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/onboarding/pnp-workflow/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceOnboardingPnpDeleteWorkflowByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteWorkflowById")
	}

	result := response.Result().(*ResponseDeviceOnboardingPnpDeleteWorkflowByID)
	return result, response, err

}
