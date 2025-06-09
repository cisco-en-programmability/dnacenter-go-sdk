package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type UsersService service

type GetUserEnrichmentDetailsHeaderParams struct {
	EntityType        string `url:"entity_type,omitempty"`         //Expects type string. User enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
	EntityValue       string `url:"entity_value,omitempty"`        //Expects type string. Contains the actual value for the entity type that has been defined
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type bool. For the enrichment details to be made available as part of the API response, this header must be set to true. This header must be explicitly passed when called from client applications outside Catalyst Center
}

type ResponseUsersGetUserEnrichmentDetails []ResponseItemUsersGetUserEnrichmentDetails // Array of ResponseUsersGetUserEnrichmentDetails
type ResponseItemUsersGetUserEnrichmentDetails struct {
	UserDetails     *ResponseItemUsersGetUserEnrichmentDetailsUserDetails       `json:"userDetails,omitempty"`     //
	ConnectedDevice *[]ResponseItemUsersGetUserEnrichmentDetailsConnectedDevice `json:"connectedDevice,omitempty"` //
}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetails struct {
	ID               string                                                                 `json:"id,omitempty"`               // Id
	ConnectionStatus string                                                                 `json:"connectionStatus,omitempty"` // Connection Status
	HostType         string                                                                 `json:"hostType,omitempty"`         // Host Type
	UserID           *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsUserID            `json:"userId,omitempty"`           // User Id
	HostName         *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostName          `json:"hostName,omitempty"`         // Host Name
	HostOs           *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostOs            `json:"hostOs,omitempty"`           // Host Os
	HostVersion      *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostVersion       `json:"hostVersion,omitempty"`      // Host Version
	SubType          string                                                                 `json:"subType,omitempty"`          // Sub Type
	LastUpdated      *int                                                                   `json:"lastUpdated,omitempty"`      // Last Updated
	HealthScore      *[]ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHealthScore     `json:"healthScore,omitempty"`      //
	HostMac          string                                                                 `json:"hostMac,omitempty"`          // Host Mac
	HostIPV4         string                                                                 `json:"hostIpV4,omitempty"`         // Host Ip V4
	HostIPV6         *[]ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostIPV6        `json:"hostIpV6,omitempty"`         // Host Ip V6
	AuthType         *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsAuthType          `json:"authType,omitempty"`         // Auth Type
	VLANID           string                                                                 `json:"vlanId,omitempty"`           // Vlan Id
	SSID             *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsSSID              `json:"ssid,omitempty"`             // Ssid
	Frequency        *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsFrequency         `json:"frequency,omitempty"`        // Frequency
	Channel          *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsChannel           `json:"channel,omitempty"`          // Channel
	ApGroup          *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsApGroup           `json:"apGroup,omitempty"`          // Ap Group
	Location         *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsLocation          `json:"location,omitempty"`         // Location
	ClientConnection string                                                                 `json:"clientConnection,omitempty"` // Client Connection
	ConnectedDevice  *[]ResponseItemUsersGetUserEnrichmentDetailsUserDetailsConnectedDevice `json:"connectedDevice,omitempty"`  // Connected Device
	IssueCount       *float64                                                               `json:"issueCount,omitempty"`       // Issue Count
	Rssi             *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsRssi              `json:"rssi,omitempty"`             // Rssi
	AvgRssi          *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsAvgRssi           `json:"avgRssi,omitempty"`          // Avg Rssi
	Snr              *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsSnr               `json:"snr,omitempty"`              // Snr
	AvgSnr           *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsAvgSnr            `json:"avgSnr,omitempty"`           // Avg Snr
	DataRate         *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsDataRate          `json:"dataRate,omitempty"`         // Data Rate
	TxBytes          *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsTxBytes           `json:"txBytes,omitempty"`          // Tx Bytes
	RxBytes          *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsRxBytes           `json:"rxBytes,omitempty"`          // Rx Bytes
	DNSSuccess       *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsDNSSuccess        `json:"dnsSuccess,omitempty"`       // Dns Success
	DNSFailure       *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsDNSFailure        `json:"dnsFailure,omitempty"`       // Dns Failure
	Onboarding       *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboarding        `json:"onboarding,omitempty"`       //
	OnboardingTime   *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingTime    `json:"onboardingTime,omitempty"`   // Onboarding Time
	Port             *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsPort              `json:"port,omitempty"`             // Port
}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsUserID interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostName interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostOs interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostVersion interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHealthScore struct {
	HealthType string `json:"healthType,omitempty"` // Health Type
	Reason     string `json:"reason,omitempty"`     // Reason
	Score      *int   `json:"score,omitempty"`      // Score
}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsHostIPV6 interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsAuthType interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsSSID interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsFrequency interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsChannel interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsApGroup interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsLocation interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsConnectedDevice interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsRssi interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsAvgRssi interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsSnr interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsAvgSnr interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsDataRate interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsTxBytes interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsRxBytes interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsDNSSuccess interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsDNSFailure interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboarding struct {
	AverageRunDuration   *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageRunDuration   `json:"averageRunDuration,omitempty"`   // Average Run Duration
	MaxRunDuration       *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxRunDuration       `json:"maxRunDuration,omitempty"`       // Max Run Duration
	AverageAssocDuration *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageAssocDuration `json:"averageAssocDuration,omitempty"` // Average Assoc Duration
	MaxAssocDuration     *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxAssocDuration     `json:"maxAssocDuration,omitempty"`     // Max Assoc Duration
	AverageAuthDuration  *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageAuthDuration  `json:"averageAuthDuration,omitempty"`  // Average Auth Duration
	MaxAuthDuration      *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxAuthDuration      `json:"maxAuthDuration,omitempty"`      // Max Auth Duration
	AverageDhcpDuration  *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageDhcpDuration  `json:"averageDhcpDuration,omitempty"`  // Average Dhcp Duration
	MaxDhcpDuration      *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxDhcpDuration      `json:"maxDhcpDuration,omitempty"`      // Max Dhcp Duration
	AAAServerIP          *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAAAServerIP          `json:"aaaServerIp,omitempty"`          // Aaa Server Ip
	DhcpServerIP         *ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingDhcpServerIP         `json:"dhcpServerIp,omitempty"`         // Dhcp Server Ip
}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageRunDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxRunDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageAssocDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxAssocDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageAuthDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxAuthDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAverageDhcpDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingMaxDhcpDuration interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingAAAServerIP interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingDhcpServerIP interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsOnboardingTime interface{}
type ResponseItemUsersGetUserEnrichmentDetailsUserDetailsPort interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDevice struct {
	DeviceDetails *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetails `json:"deviceDetails,omitempty"` //
}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetails struct {
	Family                    string                                                                                   `json:"family,omitempty"`                    // Family
	Type                      string                                                                                   `json:"type,omitempty"`                      // Type
	Location                  *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsLocation           `json:"location,omitempty"`                  // Location
	ErrorCode                 *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsErrorCode          `json:"errorCode,omitempty"`                 // Error Code
	MacAddress                string                                                                                   `json:"macAddress,omitempty"`                // Mac Address
	Role                      string                                                                                   `json:"role,omitempty"`                      // Role
	ApManagerInterfaceIP      string                                                                                   `json:"apManagerInterfaceIp,omitempty"`      // Ap Manager Interface Ip
	AssociatedWlcIP           string                                                                                   `json:"associatedWlcIp,omitempty"`           // Associated Wlc Ip
	BootDateTime              string                                                                                   `json:"bootDateTime,omitempty"`              // Boot Date Time
	CollectionStatus          string                                                                                   `json:"collectionStatus,omitempty"`          // Collection Status
	InterfaceCount            string                                                                                   `json:"interfaceCount,omitempty"`            // Interface Count
	LineCardCount             string                                                                                   `json:"lineCardCount,omitempty"`             // Line Card Count
	LineCardID                string                                                                                   `json:"lineCardId,omitempty"`                // Line Card Id
	ManagementIPAddress       string                                                                                   `json:"managementIpAddress,omitempty"`       // Management Ip Address
	MemorySize                string                                                                                   `json:"memorySize,omitempty"`                // Memory Size
	PlatformID                string                                                                                   `json:"platformId,omitempty"`                // Platform Id
	ReachabilityFailureReason string                                                                                   `json:"reachabilityFailureReason,omitempty"` // Reachability Failure Reason
	ReachabilityStatus        string                                                                                   `json:"reachabilityStatus,omitempty"`        // Reachability Status
	SNMPContact               string                                                                                   `json:"snmpContact,omitempty"`               // Snmp Contact
	SNMPLocation              string                                                                                   `json:"snmpLocation,omitempty"`              // Snmp Location
	TunnelUDPPort             *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsTunnelUDPPort      `json:"tunnelUdpPort,omitempty"`             // Tunnel Udp Port
	WaasDeviceMode            *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode     `json:"waasDeviceMode,omitempty"`            // Waas Device Mode
	Series                    string                                                                                   `json:"series,omitempty"`                    // Series
	InventoryStatusDetail     string                                                                                   `json:"inventoryStatusDetail,omitempty"`     // Inventory Status Detail
	CollectionInterval        string                                                                                   `json:"collectionInterval,omitempty"`        // Collection Interval
	SerialNumber              string                                                                                   `json:"serialNumber,omitempty"`              // Serial Number
	SoftwareVersion           string                                                                                   `json:"softwareVersion,omitempty"`           // Software Version
	RoleSource                string                                                                                   `json:"roleSource,omitempty"`                // Role Source
	Hostname                  string                                                                                   `json:"hostname,omitempty"`                  // Hostname
	UpTime                    string                                                                                   `json:"upTime,omitempty"`                    // Up Time
	LastUpdateTime            *int                                                                                     `json:"lastUpdateTime,omitempty"`            // Last Update Time
	ErrorDescription          *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription   `json:"errorDescription,omitempty"`          // Error Description
	LocationName              *ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName       `json:"locationName,omitempty"`              // Location Name
	TagCount                  string                                                                                   `json:"tagCount,omitempty"`                  // Tag Count
	LastUpdated               string                                                                                   `json:"lastUpdated,omitempty"`               // Last Updated
	InstanceUUID              string                                                                                   `json:"instanceUuid,omitempty"`              // Instance Uuid
	ID                        string                                                                                   `json:"id,omitempty"`                        // Id
	NeighborTopology          *[]ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology `json:"neighborTopology,omitempty"`          //
}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsLocation interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsErrorCode interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsTunnelUDPPort interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsWaasDeviceMode interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsErrorDescription interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsLocationName interface{}
type ResponseItemUsersGetUserEnrichmentDetailsConnectedDeviceDeviceDetailsNeighborTopology struct {
	ErrorCode *int   `json:"errorCode,omitempty"` // Error Code
	Message   string `json:"message,omitempty"`   // Message
	Detail    string `json:"detail,omitempty"`    // Detail
}

//GetUserEnrichmentDetails Get User Enrichment Details - d7a6-3928-45e8-969d
/* Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the user and devices that the user is connected to


@param GetUserEnrichmentDetailsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-user-enrichment-details
*/
func (s *UsersService) GetUserEnrichmentDetails(GetUserEnrichmentDetailsHeaderParams *GetUserEnrichmentDetailsHeaderParams) (*ResponseUsersGetUserEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/user-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetUserEnrichmentDetailsHeaderParams != nil {

		if GetUserEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetUserEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetUserEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetUserEnrichmentDetailsHeaderParams.EntityValue)
		}

		if GetUserEnrichmentDetailsHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", GetUserEnrichmentDetailsHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseUsersGetUserEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetUserEnrichmentDetails(GetUserEnrichmentDetailsHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation GetUserEnrichmentDetails")
	}

	result := response.Result().(*ResponseUsersGetUserEnrichmentDetails)
	return result, response, err

}
