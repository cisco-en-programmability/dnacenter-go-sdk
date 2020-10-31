package dnac

// TaskResponse is the TaskResponse definition
type TaskResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// TaskIDResult is the TaskIdResult definition
type TaskIDResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// CountResult is the CountResult definition
type CountResult struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// Links is the Links definition
type Links struct {
	ID              string   `json:"id,omitempty"`              //
	Label           []string `json:"label,omitempty"`           //
	LinkStatus      string   `json:"linkStatus,omitempty"`      //
	PortUtilization string   `json:"portUtilization,omitempty"` //
	Source          string   `json:"source,omitempty"`          //
	Target          string   `json:"target,omitempty"`          //
}

// NeighborTopology is the NeighborTopology definition
type NeighborTopology struct {
	Links []Links `json:"links,omitempty"` //
	Nodes []Nodes `json:"nodes,omitempty"` //
}

// Nodes is the Nodes definition
type Nodes struct {
	Clients         int    `json:"clients,omitempty"`         //
	Count           string `json:"count,omitempty"`           //
	Description     string `json:"description,omitempty"`     //
	DeviceType      string `json:"deviceType,omitempty"`      //
	FabricGroup     string `json:"fabricGroup,omitempty"`     //
	Family          string `json:"family,omitempty"`          //
	HealthScore     string `json:"healthScore,omitempty"`     //
	ID              string `json:"id,omitempty"`              //
	IP              string `json:"ip,omitempty"`              //
	Level           int    `json:"level,omitempty"`           //
	Name            string `json:"name,omitempty"`            //
	NodeType        string `json:"nodeType,omitempty"`        //
	PlatformID      string `json:"platformId,omitempty"`      //
	RadioFrequency  string `json:"radioFrequency,omitempty"`  //
	Role            string `json:"role,omitempty"`            //
	SoftwareVersion string `json:"softwareVersion,omitempty"` //
	UserID          string `json:"userId,omitempty"`          //
}

// Issue is the Issue definition
type Issue struct {
	ImpactedHosts    []ImpactedHosts    `json:"impactedHosts,omitempty"`    //
	IssueCategory    string             `json:"issueCategory,omitempty"`    //
	IssueDescription string             `json:"issueDescription,omitempty"` //
	IssueEntity      string             `json:"issueEntity,omitempty"`      //
	IssueEntityValue string             `json:"issueEntityValue,omitempty"` //
	IssueID          string             `json:"issueId,omitempty"`          //
	IssueName        string             `json:"issueName,omitempty"`        //
	IssuePriority    string             `json:"issuePriority,omitempty"`    //
	IssueSeverity    string             `json:"issueSeverity,omitempty"`    //
	IssueSource      string             `json:"issueSource,omitempty"`      //
	IssueSummary     string             `json:"issueSummary,omitempty"`     //
	IssueTimestamp   int                `json:"issueTimestamp,omitempty"`   //
	SuggestedActions []SuggestedActions `json:"suggestedActions,omitempty"` //
}

// IssueDetails is the IssueDetails definition
type IssueDetails struct {
	Issue []Issue `json:"issue,omitempty"` //
}

// SuggestedActions is the SuggestedActions definition
type SuggestedActions struct {
	Message string   `json:"message,omitempty"` //
	Steps   []string `json:"steps,omitempty"`   //
}

// Topology is the Topology definition
type Topology struct {
	Links []Links `json:"links,omitempty"` //
	Nodes []Nodes `json:"nodes,omitempty"` //
}

// Area is the Area definition
type Area struct {
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// Building is the Building definition
type Building struct {
	Address    string `json:"address,omitempty"`    //
	Latitude   int    `json:"latitude,omitempty"`   //
	Longitude  int    `json:"longitude,omitempty"`  //
	Name       string `json:"name,omitempty"`       //
	ParentName string `json:"parentName,omitempty"` //
}

// Floor is the Floor definition
type Floor struct {
	Height  int    `json:"height,omitempty"`  //
	Length  int    `json:"length,omitempty"`  //
	Name    string `json:"name,omitempty"`    //
	RfModel string `json:"rfModel,omitempty"` //
	Width   int    `json:"width,omitempty"`   //
}

// DeviceDetails is the DeviceDetails definition
type DeviceDetails struct {
	ApManagerInterfaceIP      string             `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIP           string             `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string             `json:"bootDateTime,omitempty"`              //
	Cisco360view              string             `json:"cisco360view,omitempty"`              //
	CollectionInterval        string             `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string             `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string             `json:"errorCode,omitempty"`                 //
	ErrorDescription          string             `json:"errorDescription,omitempty"`          //
	Family                    string             `json:"family,omitempty"`                    //
	Hostname                  string             `json:"hostname,omitempty"`                  //
	ID                        string             `json:"id,omitempty"`                        //
	InstanceUUID              string             `json:"instanceUuid,omitempty"`              //
	InterfaceCount            string             `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string             `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            int                `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string             `json:"lastUpdated,omitempty"`               //
	LineCardCount             string             `json:"lineCardCount,omitempty"`             //
	LineCardID                string             `json:"lineCardId,omitempty"`                //
	Location                  string             `json:"location,omitempty"`                  //
	LocationName              string             `json:"locationName,omitempty"`              //
	MacAddress                string             `json:"macAddress,omitempty"`                //
	ManagementIPAddress       string             `json:"managementIpAddress,omitempty"`       //
	MemorySize                string             `json:"memorySize,omitempty"`                //
	NeighborTopology          []NeighborTopology `json:"neighborTopology,omitempty"`          //
	PlatformID                string             `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string             `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string             `json:"reachabilityStatus,omitempty"`        //
	Role                      string             `json:"role,omitempty"`                      //
	RoleSource                string             `json:"roleSource,omitempty"`                //
	SerialNumber              string             `json:"serialNumber,omitempty"`              //
	Series                    string             `json:"series,omitempty"`                    //
	SNMPContact               string             `json:"snmpContact,omitempty"`               //
	SNMPLocation              string             `json:"snmpLocation,omitempty"`              //
	SoftwareVersion           string             `json:"softwareVersion,omitempty"`           //
	TagCount                  string             `json:"tagCount,omitempty"`                  //
	TunnelUDPPort             string             `json:"tunnelUdpPort,omitempty"`             //
	Type                      string             `json:"type,omitempty"`                      //
	UpTime                    string             `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string             `json:"waasDeviceMode,omitempty"`            //
}

// ConnectedDevice is the ConnectedDevice definition
type ConnectedDevice struct {
	DeviceDetails DeviceDetails `json:"deviceDetails,omitempty"` //
}

// HealthScore is the HealthScore definition
type HealthScore struct {
	HealthType string `json:"healthType,omitempty"` //
	Reason     string `json:"reason,omitempty"`     //
	Score      int    `json:"score,omitempty"`      //
}

// UserDetails is the UserDetails definition
type UserDetails struct {
	ApGroup          string          `json:"apGroup,omitempty"`          //
	AuthType         string          `json:"authType,omitempty"`         //
	AvgRssi          string          `json:"avgRssi,omitempty"`          //
	AvgSnr           string          `json:"avgSnr,omitempty"`           //
	Channel          string          `json:"channel,omitempty"`          //
	ClientConnection string          `json:"clientConnection,omitempty"` //
	ConnectedDevice  []string        `json:"connectedDevice,omitempty"`  //
	ConnectionStatus string          `json:"connectionStatus,omitempty"` //
	DataRate         string          `json:"dataRate,omitempty"`         //
	DNSFailure       string          `json:"dnsFailure,omitempty"`       //
	DNSSuccess       string          `json:"dnsSuccess,omitempty"`       //
	Frequency        string          `json:"frequency,omitempty"`        //
	HealthScore      []HealthScore   `json:"healthScore,omitempty"`      //
	HostIPV4         string          `json:"hostIpV4,omitempty"`         //
	HostIPV6         []string        `json:"hostIpV6,omitempty"`         //
	HostMac          string          `json:"hostMac,omitempty"`          //
	HostName         string          `json:"hostName,omitempty"`         //
	HostOs           string          `json:"hostOs,omitempty"`           //
	HostType         string          `json:"hostType,omitempty"`         //
	HostVersion      string          `json:"hostVersion,omitempty"`      //
	ID               string          `json:"id,omitempty"`               //
	IssueCount       int             `json:"issueCount,omitempty"`       //
	LastUpdated      int             `json:"lastUpdated,omitempty"`      //
	Location         string          `json:"location,omitempty"`         //
	Onboarding       UsersOnboarding `json:"onboarding,omitempty"`       //
	OnboardingTime   string          `json:"onboardingTime,omitempty"`   //
	Port             string          `json:"port,omitempty"`             //
	Rssi             string          `json:"rssi,omitempty"`             //
	RxBytes          string          `json:"rxBytes,omitempty"`          //
	Snr              string          `json:"snr,omitempty"`              //
	SSID             string          `json:"ssid,omitempty"`             //
	SubType          string          `json:"subType,omitempty"`          //
	TxBytes          string          `json:"txBytes,omitempty"`          //
	UserID           string          `json:"userId,omitempty"`           //
	VlanID           string          `json:"vlanId,omitempty"`           //
}
