package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// UsersService is the service to communicate with the Users API endpoint
type UsersService service

// ConnectedDevice is the ConnectedDevice definition
type ConnectedDevice struct {
	DeviceDetails DeviceDetails `json:"deviceDetails,omitempty"` //
}

// DeviceDetails is the DeviceDetails definition
type DeviceDetails struct {
	ApManagerInterfaceIp      string             `json:"apManagerInterfaceIp,omitempty"`      //
	AssociatedWlcIp           string             `json:"associatedWlcIp,omitempty"`           //
	BootDateTime              string             `json:"bootDateTime,omitempty"`              //
	CollectionInterval        string             `json:"collectionInterval,omitempty"`        //
	CollectionStatus          string             `json:"collectionStatus,omitempty"`          //
	ErrorCode                 string             `json:"errorCode,omitempty"`                 //
	ErrorDescription          string             `json:"errorDescription,omitempty"`          //
	Family                    string             `json:"family,omitempty"`                    //
	Hostname                  string             `json:"hostname,omitempty"`                  //
	Id                        string             `json:"id,omitempty"`                        //
	InstanceUuid              string             `json:"instanceUuid,omitempty"`              //
	InterfaceCount            string             `json:"interfaceCount,omitempty"`            //
	InventoryStatusDetail     string             `json:"inventoryStatusDetail,omitempty"`     //
	LastUpdateTime            int                `json:"lastUpdateTime,omitempty"`            //
	LastUpdated               string             `json:"lastUpdated,omitempty"`               //
	LineCardCount             string             `json:"lineCardCount,omitempty"`             //
	LineCardId                string             `json:"lineCardId,omitempty"`                //
	Location                  string             `json:"location,omitempty"`                  //
	LocationName              string             `json:"locationName,omitempty"`              //
	MacAddress                string             `json:"macAddress,omitempty"`                //
	ManagementIpAddress       string             `json:"managementIpAddress,omitempty"`       //
	MemorySize                string             `json:"memorySize,omitempty"`                //
	NeighborTopology          []NeighborTopology `json:"neighborTopology,omitempty"`          //
	PlatformId                string             `json:"platformId,omitempty"`                //
	ReachabilityFailureReason string             `json:"reachabilityFailureReason,omitempty"` //
	ReachabilityStatus        string             `json:"reachabilityStatus,omitempty"`        //
	Role                      string             `json:"role,omitempty"`                      //
	RoleSource                string             `json:"roleSource,omitempty"`                //
	SerialNumber              string             `json:"serialNumber,omitempty"`              //
	Series                    string             `json:"series,omitempty"`                    //
	SnmpContact               string             `json:"snmpContact,omitempty"`               //
	SnmpLocation              string             `json:"snmpLocation,omitempty"`              //
	SoftwareVersion           string             `json:"softwareVersion,omitempty"`           //
	TagCount                  string             `json:"tagCount,omitempty"`                  //
	TunnelUdpPort             string             `json:"tunnelUdpPort,omitempty"`             //
	Type                      string             `json:"type,omitempty"`                      //
	UpTime                    string             `json:"upTime,omitempty"`                    //
	WaasDeviceMode            string             `json:"waasDeviceMode,omitempty"`            //
}

// GetUserEnrichmentDetailsResponse is the GetUserEnrichmentDetailsResponse definition
type GetUserEnrichmentDetailsResponse struct {
	ConnectedDevice []ConnectedDevice `json:"connectedDevice,omitempty"` //
	UserDetails     UserDetails       `json:"userDetails,omitempty"`     //
}

// HealthScore is the HealthScore definition
type HealthScore struct {
	HealthType string `json:"healthType,omitempty"` //
	Reason     string `json:"reason,omitempty"`     //
	Score      int    `json:"score,omitempty"`      //
}

// NeighborTopology is the NeighborTopology definition
type NeighborTopology struct {
	Detail    string `json:"detail,omitempty"`    //
	ErrorCode int    `json:"errorCode,omitempty"` //
	Message   string `json:"message,omitempty"`   //
}

// Onboarding is the Onboarding definition
type Onboarding struct {
	AaaServerIp          string `json:"aaaServerIp,omitempty"`          //
	AverageAssocDuration string `json:"averageAssocDuration,omitempty"` //
	AverageAuthDuration  string `json:"averageAuthDuration,omitempty"`  //
	AverageDhcpDuration  string `json:"averageDhcpDuration,omitempty"`  //
	AverageRunDuration   string `json:"averageRunDuration,omitempty"`   //
	DhcpServerIp         string `json:"dhcpServerIp,omitempty"`         //
	MaxAssocDuration     string `json:"maxAssocDuration,omitempty"`     //
	MaxAuthDuration      string `json:"maxAuthDuration,omitempty"`      //
	MaxDhcpDuration      string `json:"maxDhcpDuration,omitempty"`      //
	MaxRunDuration       string `json:"maxRunDuration,omitempty"`       //
}

// UserDetails is the UserDetails definition
type UserDetails struct {
	ApGroup          string        `json:"apGroup,omitempty"`          //
	AuthType         string        `json:"authType,omitempty"`         //
	AvgRssi          string        `json:"avgRssi,omitempty"`          //
	AvgSnr           string        `json:"avgSnr,omitempty"`           //
	Channel          string        `json:"channel,omitempty"`          //
	ClientConnection string        `json:"clientConnection,omitempty"` //
	ConnectedDevice  []string      `json:"connectedDevice,omitempty"`  //
	ConnectionStatus string        `json:"connectionStatus,omitempty"` //
	DataRate         string        `json:"dataRate,omitempty"`         //
	DnsFailure       string        `json:"dnsFailure,omitempty"`       //
	DnsSuccess       string        `json:"dnsSuccess,omitempty"`       //
	Frequency        string        `json:"frequency,omitempty"`        //
	HealthScore      []HealthScore `json:"healthScore,omitempty"`      //
	HostIpV4         string        `json:"hostIpV4,omitempty"`         //
	HostIpV6         []string      `json:"hostIpV6,omitempty"`         //
	HostMac          string        `json:"hostMac,omitempty"`          //
	HostName         string        `json:"hostName,omitempty"`         //
	HostOs           string        `json:"hostOs,omitempty"`           //
	HostType         string        `json:"hostType,omitempty"`         //
	HostVersion      string        `json:"hostVersion,omitempty"`      //
	Id               string        `json:"id,omitempty"`               //
	IssueCount       int           `json:"issueCount,omitempty"`       //
	LastUpdated      int           `json:"lastUpdated,omitempty"`      //
	Location         string        `json:"location,omitempty"`         //
	Onboarding       Onboarding    `json:"onboarding,omitempty"`       //
	OnboardingTime   string        `json:"onboardingTime,omitempty"`   //
	Port             string        `json:"port,omitempty"`             //
	Rssi             string        `json:"rssi,omitempty"`             //
	RxBytes          string        `json:"rxBytes,omitempty"`          //
	Snr              string        `json:"snr,omitempty"`              //
	Ssid             string        `json:"ssid,omitempty"`             //
	SubType          string        `json:"subType,omitempty"`          //
	TxBytes          string        `json:"txBytes,omitempty"`          //
	UserId           string        `json:"userId,omitempty"`           //
	VlanId           string        `json:"vlanId,omitempty"`           //
}

// GetUserEnrichmentDetails getUserEnrichmentDetails
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user and devices that the user is connected to
@param entity_type User enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
@param entity_value Contains the actual value for the entity type that has been defined
*/
func (s *UsersService) GetUserEnrichmentDetails() (*GetUserEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/user-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetUserEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetUserEnrichmentDetailsResponse)
	return result, response, err

}
