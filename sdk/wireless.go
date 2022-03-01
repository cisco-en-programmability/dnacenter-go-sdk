package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type WirelessService service

type SensorTestResultsQueryParams struct {
	SiteID        string  `url:"siteId,omitempty"`        //Assurance site UUID
	StartTime     float64 `url:"startTime,omitempty"`     //The epoch time in milliseconds
	EndTime       float64 `url:"endTime,omitempty"`       //The epoch time in milliseconds
	TestFailureBy string  `url:"testFailureBy,omitempty"` //Obtain failure statistics group by "area", "building", or "floor"
}
type CreateAndProvisionSSIDHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type DeleteSSIDAndProvisionItToDevicesHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type GetEnterpriseSSIDQueryParams struct {
	SSIDName string `url:"ssidName,omitempty"` //Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
}
type ApProvisionHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string.
}
type CreateUpdateDynamicInterfaceHeaderParams struct {
	Runsync string `url:"__runsync,omitempty"` //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout string `url:"__timeout,omitempty"` //Expects type float64. If __runsync is set to ‘true’, this defines the maximum time before which if the API completes its execution, then a synchronous response is returned.  If the time taken for the API to complete the execution, exceeds this time, then an asynchronous response is returned with an execution id, that can be used to get the status and response associated with the API execution
}
type GetDynamicInterfaceQueryParams struct {
	InterfaceName string `url:"interface-name,omitempty"` //dynamic-interface name, if not specified all the existing dynamic interfaces will be retrieved
}
type DeleteDynamicInterfaceHeaderParams struct {
	Runsync string `url:"__runsync,omitempty"` //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout string `url:"__timeout,omitempty"` //Expects type float64. If __runsync is set to ‘true’, this defines the maximum time before which if the API completes its execution, then a synchronous response is returned.  If the time taken for the API to complete the execution, exceeds this time, then an asynchronous response is returned with an execution id, that can be used to get the status and response associated with the API execution
}
type GetWirelessProfileQueryParams struct {
	ProfileName string `url:"profileName,omitempty"` //Wireless Network Profile Name
}
type ProvisionUpdateHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type RetrieveRfProfilesQueryParams struct {
	RfProfileName string `url:"rf-profile-name,omitempty"` //
}

type ResponseWirelessSensorTestResults struct {
	Summary      *ResponseWirelessSensorTestResultsSummary        `json:"summary,omitempty"`      //
	FailureStats *[]ResponseWirelessSensorTestResultsFailureStats `json:"failureStats,omitempty"` //
}
type ResponseWirelessSensorTestResultsSummary struct {
	TotalTestCount  *int                                                     `json:"totalTestCount,omitempty"`   // Total Test Count
	OnBoarding      *ResponseWirelessSensorTestResultsSummaryOnBoarding      `json:"ONBOARDING,omitempty"`       //
	PERfORMAncE     *ResponseWirelessSensorTestResultsSummaryPERfORMAncE     `json:"PERFORMANCE,omitempty"`      //
	NETWORKSERVICES *ResponseWirelessSensorTestResultsSummaryNETWORKSERVICES `json:"NETWORK_SERVICES,omitempty"` //
	ApPCONNECTIVITY *ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITY `json:"APP_CONNECTIVITY,omitempty"` //
	RfASSESSMENT    *ResponseWirelessSensorTestResultsSummaryRfASSESSMENT    `json:"RF_ASSESSMENT,omitempty"`    //
	Email           *ResponseWirelessSensorTestResultsSummaryEmail           `json:"EMAIL,omitempty"`            //
}
type ResponseWirelessSensorTestResultsSummaryOnBoarding struct {
	Auth  *ResponseWirelessSensorTestResultsSummaryOnBoardingAuth  `json:"AUTH,omitempty"`  //
	DHCP  *ResponseWirelessSensorTestResultsSummaryOnBoardingDHCP  `json:"DHCP,omitempty"`  //
	Assoc *ResponseWirelessSensorTestResultsSummaryOnBoardingAssoc `json:"ASSOC,omitempty"` //
}
type ResponseWirelessSensorTestResultsSummaryOnBoardingAuth struct {
	PassCount *int `json:"passCount,omitempty"` // Pass Count
	FailCount *int `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryOnBoardingDHCP struct {
	PassCount *int     `json:"passCount,omitempty"` // Pass Count
	FailCount *float64 `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryOnBoardingAssoc struct {
	PassCount *int `json:"passCount,omitempty"` // Pass Count
	FailCount *int `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryPERfORMAncE struct {
	IPSLASENDER *ResponseWirelessSensorTestResultsSummaryPERfORMAncEIPSLASENDER `json:"IPSLASENDER,omitempty"` //
}
type ResponseWirelessSensorTestResultsSummaryPERfORMAncEIPSLASENDER struct {
	PassCount *int `json:"passCount,omitempty"` // Pass Count
	FailCount *int `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryNETWORKSERVICES struct {
	DNS *ResponseWirelessSensorTestResultsSummaryNETWORKSERVICESDNS `json:"DNS,omitempty"` //
}
type ResponseWirelessSensorTestResultsSummaryNETWORKSERVICESDNS struct {
	PassCount *int     `json:"passCount,omitempty"` // Pass Count
	FailCount *float64 `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITY struct {
	HOSTREACHABILITY *ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYHOSTREACHABILITY `json:"HOST_REACHABILITY,omitempty"` //
	WebServer        *ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYWebServer        `json:"WEBSERVER,omitempty"`         //
	FileTransfer     *ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYFileTransfer     `json:"FILETRANSFER,omitempty"`      //
}
type ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYHOSTREACHABILITY struct {
	PassCount *int     `json:"passCount,omitempty"` // Pass Count
	FailCount *float64 `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYWebServer struct {
	PassCount *int `json:"passCount,omitempty"` // Pass Count
	FailCount *int `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryApPCONNECTIVITYFileTransfer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Pass Count
	FailCount *int     `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryRfASSESSMENT struct {
	DATARATE *ResponseWirelessSensorTestResultsSummaryRfASSESSMENTDATARATE `json:"DATA_RATE,omitempty"` //
	SNR      *ResponseWirelessSensorTestResultsSummaryRfASSESSMENTSNR      `json:"SNR,omitempty"`       //
}
type ResponseWirelessSensorTestResultsSummaryRfASSESSMENTDATARATE struct {
	PassCount *int `json:"passCount,omitempty"` // Pass Count
	FailCount *int `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryRfASSESSMENTSNR struct {
	PassCount *int     `json:"passCount,omitempty"` // Pass Count
	FailCount *float64 `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsSummaryEmail struct {
	MailServer *ResponseWirelessSensorTestResultsSummaryEmailMailServer `json:"MAILSERVER,omitempty"` //
}
type ResponseWirelessSensorTestResultsSummaryEmailMailServer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Pass Count
	FailCount *int     `json:"failCount,omitempty"` // Fail Count
}
type ResponseWirelessSensorTestResultsFailureStats struct {
	ErrorCode    *int   `json:"errorCode,omitempty"`    // Error Code
	ErrorTitle   string `json:"errorTitle,omitempty"`   // Error Title
	TestType     string `json:"testType,omitempty"`     // Test Type
	TestCategory string `json:"testCategory,omitempty"` // Test Category
}
type ResponseWirelessCreateAndProvisionSSID struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessDeleteSSIDAndProvisionItToDevices struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessGetEnterpriseSSID []ResponseItemWirelessGetEnterpriseSSID // Array of ResponseWirelessGetEnterpriseSSID
type ResponseItemWirelessGetEnterpriseSSID struct {
	InstanceUUID       string                                              `json:"instanceUuid,omitempty"`       // Instance Uuid
	Version            *int                                                `json:"version,omitempty"`            // Version
	SSIDDetails        *[]ResponseItemWirelessGetEnterpriseSSIDSSIDDetails `json:"ssidDetails,omitempty"`        //
	GroupUUID          string                                              `json:"groupUuid,omitempty"`          // Group Uuid
	InheritedGroupUUID string                                              `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid
	InheritedGroupName string                                              `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseItemWirelessGetEnterpriseSSIDSSIDDetails struct {
	Name                string `json:"name,omitempty"`                // SSID Name
	WLANType            string `json:"wlanType,omitempty"`            // Wlan Type
	EnableFastLane      *bool  `json:"enableFastLane,omitempty"`      // Enable Fast Lane
	SecurityLevel       string `json:"securityLevel,omitempty"`       // Security Level
	AuthServer          string `json:"authServer,omitempty"`          // Auth Server
	Passphrase          string `json:"passphrase,omitempty"`          // Passphrase
	TrafficType         string `json:"trafficType,omitempty"`         // Traffic Type
	EnableMacFiltering  *bool  `json:"enableMACFiltering,omitempty"`  // Enable MAC Filtering
	IsEnabled           *bool  `json:"isEnabled,omitempty"`           // Is Enabled
	IsFabric            *bool  `json:"isFabric,omitempty"`            // Is Fabric
	FastTransition      string `json:"fastTransition,omitempty"`      // Fast Transition
	RadioPolicy         string `json:"radioPolicy,omitempty"`         // Radio Policy
	EnableBroadcastSSID *bool  `json:"enableBroadcastSSID,omitempty"` // Enable Broadcast SSID
}
type ResponseWirelessCreateEnterpriseSSID struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessUpdateEnterpriseSSID struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessDeleteEnterpriseSSID struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessDeleteWirelessProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessApProvision struct {
	ExecutionID  string `json:"executionId,omitempty"`  // Execution Id
	ExecutionURL string `json:"executionUrl,omitempty"` // Execution Url
	Message      string `json:"message,omitempty"`      // Message
}
type ResponseWirelessCreateUpdateDynamicInterface struct {
	ExecutionID  string `json:"executionId,omitempty"`  // Execution Id
	ExecutionURL string `json:"executionUrl,omitempty"` // Execution Url
	Message      string `json:"message,omitempty"`      // Message
}
type ResponseWirelessGetDynamicInterface []ResponseItemWirelessGetDynamicInterface // Array of ResponseWirelessGetDynamicInterface
type ResponseItemWirelessGetDynamicInterface struct {
	InterfaceName string   `json:"interfaceName,omitempty"` // dynamic interface name
	VLANID        *float64 `json:"vlanId,omitempty"`        // Vlan id
}
type ResponseWirelessUpdateWirelessProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessCreateWirelessProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessGetWirelessProfile []ResponseItemWirelessGetWirelessProfile // Array of ResponseWirelessGetWirelessProfile
type ResponseItemWirelessGetWirelessProfile struct {
	ProfileDetails *ResponseItemWirelessGetWirelessProfileProfileDetails `json:"profileDetails,omitempty"` //
}
type ResponseItemWirelessGetWirelessProfileProfileDetails struct {
	Name        string                                                             `json:"name,omitempty"`        // Profile Name
	Sites       []string                                                           `json:"sites,omitempty"`       // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])
	SSIDDetails *[]ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails struct {
	Name          string                                                                      `json:"name,omitempty"`          // SSID Name
	Type          string                                                                      `json:"type,omitempty"`          // SSID Type(enum: Enterprise/Guest)
	EnableFabric  *bool                                                                       `json:"enableFabric,omitempty"`  // true if fabric is enabled else false
	FlexConnect   *ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string                                                                      `json:"interfaceName,omitempty"` // Interface Name
}
type ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To VLAN ID
}
type ResponseWirelessProvisionUpdate struct {
	ExecutionID       string                                            `json:"executionId,omitempty"`       // Execution Id
	ExecutionURL      string                                            `json:"executionUrl,omitempty"`      // Execution Url
	ProvisioningTasks *ResponseWirelessProvisionUpdateProvisioningTasks `json:"provisioningTasks,omitempty"` //
}
type ResponseWirelessProvisionUpdateProvisioningTasks struct {
	Success []string `json:"success,omitempty"` // Success
	Failed  []string `json:"failed,omitempty"`  // Failed
}
type ResponseWirelessProvision struct {
	ExecutionID       string                                      `json:"executionId,omitempty"`       // Execution Id
	ExecutionURL      string                                      `json:"executionUrl,omitempty"`      // Execution Url
	ProvisioningTasks *ResponseWirelessProvisionProvisioningTasks `json:"provisioningTasks,omitempty"` //
}
type ResponseWirelessProvisionProvisioningTasks struct {
	Success []string `json:"success,omitempty"` // Success
	Failed  []string `json:"failed,omitempty"`  // Failed
}
type ResponseWirelessPSKOverride struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessRetrieveRfProfiles struct {
	Response *[]ResponseWirelessRetrieveRfProfilesResponse `json:"response,omitempty"` //
}
type ResponseWirelessRetrieveRfProfilesResponse struct {
	Name                 string                                                          `json:"name,omitempty"`                 // radio profile name
	DefaultRfProfile     *bool                                                           `json:"defaultRfProfile,omitempty"`     // is default radio profile
	ChannelWidth         string                                                          `json:"channelWidth,omitempty"`         // Channel Width
	EnableBrownField     *bool                                                           `json:"enableBrownField,omitempty"`     // is brownfield enabled
	EnableCustom         *bool                                                           `json:"enableCustom,omitempty"`         // is Custom Enable
	EnableRadioTypeA     *bool                                                           `json:"enableRadioTypeA,omitempty"`     // Enable Radio Type A
	EnableRadioTypeB     *bool                                                           `json:"enableRadioTypeB,omitempty"`     // Enable Radio Type B
	RadioTypeAProperties *ResponseWirelessRetrieveRfProfilesResponseRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //
	RadioTypeBProperties *ResponseWirelessRetrieveRfProfilesResponseRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //
}
type ResponseWirelessRetrieveRfProfilesResponseRadioTypeAProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent Profile name
	RadioChannels      string `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Min Power Level
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Max Power Level
}
type ResponseWirelessRetrieveRfProfilesResponseRadioTypeBProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent Profile name
	RadioChannels      string `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Min Power Level
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Max Power Level
}
type ResponseWirelessCreateOrUpdateRfProfile struct {
	ExecutionID  string `json:"executionId,omitempty"`  // Execution Id
	ExecutionURL string `json:"executionUrl,omitempty"` // Execution Url
	Message      string `json:"message,omitempty"`      // Message
}
type ResponseWirelessDeleteRfProfiles struct {
	ExecutionID  string `json:"executionId,omitempty"`  // Execution Id
	ExecutionURL string `json:"executionUrl,omitempty"` // Execution Url
	Message      string `json:"message,omitempty"`      // Message
}
type RequestWirelessCreateAndProvisionSSID struct {
	ManagedApLocations []string                                          `json:"managedAPLocations,omitempty"` // Managed AP Locations (Enter entire Site(s) hierarchy)
	SSIDDetails        *RequestWirelessCreateAndProvisionSSIDSSIDDetails `json:"ssidDetails,omitempty"`        //
	SSIDType           string                                            `json:"ssidType,omitempty"`           // SSID Type
	EnableFabric       *bool                                             `json:"enableFabric,omitempty"`       // Enable SSID for Fabric
	FlexConnect        *RequestWirelessCreateAndProvisionSSIDFlexConnect `json:"flexConnect,omitempty"`        //
}
type RequestWirelessCreateAndProvisionSSIDSSIDDetails struct {
	Name                string `json:"name,omitempty"`                // SSID Name
	SecurityLevel       string `json:"securityLevel,omitempty"`       // Security Level(For guest SSID OPEN/WEB_AUTH, For Enterprise SSID ENTERPRISE/PERSONAL/OPEN)
	EnableFastLane      *bool  `json:"enableFastLane,omitempty"`      // Enable Fast Lane
	Passphrase          string `json:"passphrase,omitempty"`          // Pass Phrase ( Only applicable for SSID with PERSONAL auth type )
	TrafficType         string `json:"trafficType,omitempty"`         // Traffic Type
	EnableBroadcastSSID *bool  `json:"enableBroadcastSSID,omitempty"` // Enable Broadcast SSID
	RadioPolicy         string `json:"radioPolicy,omitempty"`         // Radio Policy. Allowed values are 'Dual band operation (2.4GHz and 5GHz)', 'Dual band operation with band select', '5GHz only', '2.4GHz only'.
	EnableMacFiltering  *bool  `json:"enableMACFiltering,omitempty"`  // Enable MAC Filtering
	FastTransition      string `json:"fastTransition,omitempty"`      // Fast Transition
	WebAuthURL          string `json:"webAuthURL,omitempty"`          // Web Auth URL
}
type RequestWirelessCreateAndProvisionSSIDFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // Enable Flex Connect
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan (range is 1 to 4094)
}
type RequestWirelessCreateEnterpriseSSID struct {
	Name                             string `json:"name,omitempty"`                             // Enter SSID Name
	SecurityLevel                    string `json:"securityLevel,omitempty"`                    // Security Level
	Passphrase                       string `json:"passphrase,omitempty"`                       // Pass Phrase (Only applicable for SSID with PERSONAL security level)
	EnableFastLane                   *bool  `json:"enableFastLane,omitempty"`                   // Enable Fast Lane
	EnableMacFiltering               *bool  `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	TrafficType                      string `json:"trafficType,omitempty"`                      // Traffic Type
	RadioPolicy                      string `json:"radioPolicy,omitempty"`                      // Radio Policy. Allowed values are 'Dual band operation (2.4GHz and 5GHz)', 'Dual band operation with band select', '5GHz only', '2.4GHz only'.
	EnableBroadcastSSID              *bool  `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcast SSID
	FastTransition                   string `json:"fastTransition,omitempty"`                   // Fast Transition
	EnableSessionTimeOut             *bool  `json:"enableSessionTimeOut,omitempty"`             // Enable Session Timeout
	SessionTimeOut                   *int   `json:"sessionTimeOut,omitempty"`                   // Session Time Out
	EnableClientExclusion            *bool  `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *int   `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool  `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *int   `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout
	EnableDirectedMulticastService   *bool  `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed Multicast Service
	EnableNeighborList               *bool  `json:"enableNeighborList,omitempty"`               // Enable Neighbor List
	MfpClientProtection              string `json:"mfpClientProtection,omitempty"`              // Management Frame Protection Client
}
type RequestWirelessUpdateEnterpriseSSID struct {
	Name                             string `json:"name,omitempty"`                             // Enter SSID Name
	SecurityLevel                    string `json:"securityLevel,omitempty"`                    // Security Level
	Passphrase                       string `json:"passphrase,omitempty"`                       // Pass Phrase (Only applicable for SSID with PERSONAL security level)
	EnableFastLane                   *bool  `json:"enableFastLane,omitempty"`                   // Enable Fast Lane
	EnableMacFiltering               *bool  `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	TrafficType                      string `json:"trafficType,omitempty"`                      // Traffic Type
	RadioPolicy                      string `json:"radioPolicy,omitempty"`                      // Radio Policy. Allowed values are 'Dual band operation (2.4GHz and 5GHz)', 'Dual band operation with band select', '5GHz only', '2.4GHz only'
	EnableBroadcastSSID              *bool  `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcast SSID
	FastTransition                   string `json:"fastTransition,omitempty"`                   // Fast Transition
	EnableSessionTimeOut             *bool  `json:"enableSessionTimeOut,omitempty"`             // Enable Session Timeout
	SessionTimeOut                   *int   `json:"sessionTimeOut,omitempty"`                   // Session Time Out
	EnableClientExclusion            *bool  `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *int   `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool  `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *int   `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout
	EnableDirectedMulticastService   *bool  `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed Multicast Service
	EnableNeighborList               *bool  `json:"enableNeighborList,omitempty"`               // Enable Neighbor List
	MfpClientProtection              string `json:"mfpClientProtection,omitempty"`              // Management Frame Protection Client
}
type RequestWirelessApProvision []RequestItemWirelessApProvision // Array of RequestWirelessAPProvision
type RequestItemWirelessApProvision struct {
	RfProfile           string   `json:"rfProfile,omitempty"`           // Radio frequency profile name
	SiteID              string   `json:"siteId,omitempty"`              // Site name hierarchy(ex: Global/...)
	DeviceName          string   `json:"deviceName,omitempty"`          // Device name
	CustomApGroupName   string   `json:"customApGroupName,omitempty"`   // Custom AP group name
	CustomFlexGroupName []string `json:"customFlexGroupName,omitempty"` // ["Custom flex group name"]
	Type                string   `json:"type,omitempty"`                // ApWirelessConfiguration
	SiteNameHierarchy   string   `json:"siteNameHierarchy,omitempty"`   // Site name hierarchy(ex: Global/...)
}
type RequestWirelessCreateUpdateDynamicInterface struct {
	InterfaceName string   `json:"interfaceName,omitempty"` // dynamic-interface name
	VLANID        *float64 `json:"vlanId,omitempty"`        // Vlan Id
}
type RequestWirelessUpdateWirelessProfile struct {
	ProfileDetails *RequestWirelessUpdateWirelessProfileProfileDetails `json:"profileDetails,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileProfileDetails struct {
	Name        string                                                           `json:"name,omitempty"`        // Profile Name
	Sites       []string                                                         `json:"sites,omitempty"`       // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])
	SSIDDetails *[]RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails struct {
	Name          string                                                                    `json:"name,omitempty"`          // Ssid Name
	Type          string                                                                    `json:"type,omitempty"`          // Ssid Type(enum: Enterprise/Guest)
	EnableFabric  *bool                                                                     `json:"enableFabric,omitempty"`  // true is ssid is fabric else false
	FlexConnect   *RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string                                                                    `json:"interfaceName,omitempty"` // Interface Name
}
type RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan
}
type RequestWirelessCreateWirelessProfile struct {
	ProfileDetails *RequestWirelessCreateWirelessProfileProfileDetails `json:"profileDetails,omitempty"` //
}
type RequestWirelessCreateWirelessProfileProfileDetails struct {
	Name        string                                                           `json:"name,omitempty"`        // Profile Name
	Sites       []string                                                         `json:"sites,omitempty"`       // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])
	SSIDDetails *[]RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails struct {
	Name          string                                                                    `json:"name,omitempty"`          // Ssid Name
	Type          string                                                                    `json:"type,omitempty"`          // Ssid Type(enum: Enterprise/Guest)
	EnableFabric  *bool                                                                     `json:"enableFabric,omitempty"`  // true is ssid is fabric else false
	FlexConnect   *RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string                                                                    `json:"interfaceName,omitempty"` // Interface Name
}
type RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan
}
type RequestWirelessProvisionUpdate []RequestItemWirelessProvisionUpdate // Array of RequestWirelessProvisionUpdate
type RequestItemWirelessProvisionUpdate struct {
	DeviceName         string                                                 `json:"deviceName,omitempty"`         // Device Name
	ManagedApLocations []string                                               `json:"managedAPLocations,omitempty"` // Managed APLocations
	DynamicInterfaces  *[]RequestItemWirelessProvisionUpdateDynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
}
type RequestItemWirelessProvisionUpdateDynamicInterfaces struct {
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     // Interface IPAddress
	InterfaceNetmaskInCIDR *int   `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       // Interface Gateway
	LagOrPortNumber        *int   `json:"lagOrPortNumber,omitempty"`        // Lag Or Port Number
	VLANID                 *int   `json:"vlanId,omitempty"`                 // Vlan Id
	InterfaceName          string `json:"interfaceName,omitempty"`          // Interface Name
}
type RequestWirelessProvision []RequestItemWirelessProvision // Array of RequestWirelessProvision
type RequestItemWirelessProvision struct {
	DeviceName         string                                           `json:"deviceName,omitempty"`         // Controller Name
	Site               string                                           `json:"site,omitempty"`               // Full Site Hierarchy where device has to be assigned
	ManagedApLocations []string                                         `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)
	DynamicInterfaces  *[]RequestItemWirelessProvisionDynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
}
type RequestItemWirelessProvisionDynamicInterfaces struct {
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     // Interface IP Address
	InterfaceNetmaskInCIDR *int   `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       // Interface Gateway
	LagOrPortNumber        *int   `json:"lagOrPortNumber,omitempty"`        // Lag Or Port Number
	VLANID                 *int   `json:"vlanId,omitempty"`                 // VLAN ID
	InterfaceName          string `json:"interfaceName,omitempty"`          // Interface Name
}
type RequestWirelessPSKOverride []RequestItemWirelessPSKOverride // Array of RequestWirelessPSKOverride
type RequestItemWirelessPSKOverride struct {
	SSID       string `json:"ssid,omitempty"`       // enterprise ssid name(already created/present)
	Site       string `json:"site,omitempty"`       // site name hierarchy (ex: Global/aaa/zzz/...)
	PassPhrase string `json:"passPhrase,omitempty"` // Pass phrase (create/update)
}
type RequestWirelessCreateOrUpdateRfProfile struct {
	Name                 string                                                      `json:"name,omitempty"`                 // custom RF profile name
	DefaultRfProfile     *bool                                                       `json:"defaultRfProfile,omitempty"`     // isDefault rf-profile
	EnableRadioTypeA     *bool                                                       `json:"enableRadioTypeA,omitempty"`     // tru if Enable Radio Type A else false
	EnableRadioTypeB     *bool                                                       `json:"enableRadioTypeB,omitempty"`     // true if Enable Radio Type B else false
	ChannelWidth         string                                                      `json:"channelWidth,omitempty"`         // rf-profile channel width
	EnableCustom         *bool                                                       `json:"enableCustom,omitempty"`         // true if enable custom rf-profile else false
	EnableBrownField     *bool                                                       `json:"enableBrownField,omitempty"`     // true if enable brown field for rf-profile else false
	RadioTypeAProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //
	RadioTypeBProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent rf-profile name
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent rf-profile name
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level
}

//SensorTestResults Sensor Test Results - 87ae-7b21-4f0b-a838
/* Intent API to get SENSOR test result summary


@param SensorTestResultsQueryParams Filtering parameter
*/
func (s *WirelessService) SensorTestResults(SensorTestResultsQueryParams *SensorTestResultsQueryParams) (*ResponseWirelessSensorTestResults, *resty.Response, error) {
	path := "/dna/intent/api/v1/AssuranceGetSensorTestResults"

	queryString, _ := query.Values(SensorTestResultsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessSensorTestResults{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation SensorTestResults")
	}

	result := response.Result().(*ResponseWirelessSensorTestResults)
	return result, response, err

}

//GetEnterpriseSSID Get Enterprise SSID - cca5-19ba-45eb-b423
/* Gets either one or all the enterprise SSID


@param GetEnterpriseSSIDQueryParams Filtering parameter
*/
func (s *WirelessService) GetEnterpriseSSID(GetEnterpriseSSIDQueryParams *GetEnterpriseSSIDQueryParams) (*ResponseWirelessGetEnterpriseSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid"

	queryString, _ := query.Values(GetEnterpriseSSIDQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetEnterpriseSSID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessGetEnterpriseSSID)
	return result, response, err

}

//GetDynamicInterface Get dynamic interface - c5b0-c978-4dfb-90b4
/* Get one or all dynamic interface(s)


@param GetDynamicInterfaceQueryParams Filtering parameter
*/
func (s *WirelessService) GetDynamicInterface(GetDynamicInterfaceQueryParams *GetDynamicInterfaceQueryParams) (*ResponseWirelessGetDynamicInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	queryString, _ := query.Values(GetDynamicInterfaceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetDynamicInterface{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDynamicInterface")
	}

	result := response.Result().(*ResponseWirelessGetDynamicInterface)
	return result, response, err

}

//GetWirelessProfile Get Wireless Profile - b3a1-c880-4c8b-9b8b
/* Gets either one or all the wireless network profiles if no name is provided for network-profile.


@param GetWirelessProfileQueryParams Filtering parameter
*/
func (s *WirelessService) GetWirelessProfile(GetWirelessProfileQueryParams *GetWirelessProfileQueryParams) (*ResponseWirelessGetWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/profile"

	queryString, _ := query.Values(GetWirelessProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfile)
	return result, response, err

}

//RetrieveRfProfiles Retrieve RF profiles - 098c-ab91-41c9-a3fe
/* Retrieve all RF profiles


@param RetrieveRFProfilesQueryParams Filtering parameter
*/
func (s *WirelessService) RetrieveRfProfiles(RetrieveRFProfilesQueryParams *RetrieveRfProfilesQueryParams) (*ResponseWirelessRetrieveRfProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/rf-profile"

	queryString, _ := query.Values(RetrieveRFProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveRfProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RetrieveRfProfiles")
	}

	result := response.Result().(*ResponseWirelessRetrieveRfProfiles)
	return result, response, err

}

//CreateAndProvisionSSID Create and Provision SSID - 1eb7-2ad3-4e09-8990
/* Creates SSID, updates the SSID to the corresponding site profiles and provision it to the devices matching the given sites


@param CreateAndProvisionSSIDHeaderParams Custom header parameters
*/
func (s *WirelessService) CreateAndProvisionSSID(requestWirelessCreateAndProvisionSSID *RequestWirelessCreateAndProvisionSSID, CreateAndProvisionSSIDHeaderParams *CreateAndProvisionSSIDHeaderParams) (*ResponseWirelessCreateAndProvisionSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/ssid"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateAndProvisionSSIDHeaderParams != nil {

		if CreateAndProvisionSSIDHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", CreateAndProvisionSSIDHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessCreateAndProvisionSSID).
		SetResult(&ResponseWirelessCreateAndProvisionSSID{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateAndProvisionSsid")
	}

	result := response.Result().(*ResponseWirelessCreateAndProvisionSSID)
	return result, response, err

}

//CreateEnterpriseSSID Create Enterprise SSID - 8a96-fb95-4d09-a349
/* Creates enterprise SSID


 */
func (s *WirelessService) CreateEnterpriseSSID(requestWirelessCreateEnterpriseSSID *RequestWirelessCreateEnterpriseSSID) (*ResponseWirelessCreateEnterpriseSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateEnterpriseSSID).
		SetResult(&ResponseWirelessCreateEnterpriseSSID{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessCreateEnterpriseSSID)
	return result, response, err

}

//ApProvision AP Provision - d897-19b8-47aa-a9c4
/* Access Point Provision and ReProvision


@param APProvisionHeaderParams Custom header parameters
*/
func (s *WirelessService) ApProvision(requestWirelessAPProvision *RequestWirelessApProvision, APProvisionHeaderParams *ApProvisionHeaderParams) (*ResponseWirelessApProvision, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/ap-provision"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if APProvisionHeaderParams != nil {

		if APProvisionHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", APProvisionHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessAPProvision).
		SetResult(&ResponseWirelessApProvision{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ApProvision")
	}

	result := response.Result().(*ResponseWirelessApProvision)
	return result, response, err

}

//CreateUpdateDynamicInterface Create Update Dynamic interface - daa0-bb75-4e2a-8da6
/* API to create or update an dynamic interface


@param CreateUpdateDynamicInterfaceHeaderParams Custom header parameters
*/
func (s *WirelessService) CreateUpdateDynamicInterface(requestWirelessCreateUpdateDynamicInterface *RequestWirelessCreateUpdateDynamicInterface, CreateUpdateDynamicInterfaceHeaderParams *CreateUpdateDynamicInterfaceHeaderParams) (*ResponseWirelessCreateUpdateDynamicInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if CreateUpdateDynamicInterfaceHeaderParams != nil {

		if CreateUpdateDynamicInterfaceHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", CreateUpdateDynamicInterfaceHeaderParams.Runsync)
		}

		if CreateUpdateDynamicInterfaceHeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", CreateUpdateDynamicInterfaceHeaderParams.Timeout)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessCreateUpdateDynamicInterface).
		SetResult(&ResponseWirelessCreateUpdateDynamicInterface{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateUpdateDynamicInterface")
	}

	result := response.Result().(*ResponseWirelessCreateUpdateDynamicInterface)
	return result, response, err

}

//CreateWirelessProfile Create Wireless Profile - 7097-6962-4bf9-88d5
/* Creates Wireless Network Profile on DNAC and associates sites and SSIDs to it.


 */
func (s *WirelessService) CreateWirelessProfile(requestWirelessCreateWirelessProfile *RequestWirelessCreateWirelessProfile) (*ResponseWirelessCreateWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateWirelessProfile).
		SetResult(&ResponseWirelessCreateWirelessProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessCreateWirelessProfile)
	return result, response, err

}

//Provision Provision - d09b-08a3-447a-a3b9
/* Provision wireless devices


 */
func (s *WirelessService) Provision(requestWirelessProvision *RequestWirelessProvision) (*ResponseWirelessProvision, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/provision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessProvision).
		SetResult(&ResponseWirelessProvision{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation Provision")
	}

	result := response.Result().(*ResponseWirelessProvision)
	return result, response, err

}

//PSKOverride PSK override - 46ad-ab75-47c9-8762
/* Update/override pass phrase of enterprise SSID


 */
func (s *WirelessService) PSKOverride(requestWirelessPSKOverride *RequestWirelessPSKOverride) (*ResponseWirelessPSKOverride, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/psk-override"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessPSKOverride).
		SetResult(&ResponseWirelessPSKOverride{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation PSKOverride")
	}

	result := response.Result().(*ResponseWirelessPSKOverride)
	return result, response, err

}

//CreateOrUpdateRfProfile Create or Update RF profile - b783-2967-4878-b815
/* Create or Update RF profile


 */
func (s *WirelessService) CreateOrUpdateRfProfile(requestWirelessCreateOrUpdateRFProfile *RequestWirelessCreateOrUpdateRfProfile) (*ResponseWirelessCreateOrUpdateRfProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/rf-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateOrUpdateRFProfile).
		SetResult(&ResponseWirelessCreateOrUpdateRfProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateOrUpdateRfProfile")
	}

	result := response.Result().(*ResponseWirelessCreateOrUpdateRfProfile)
	return result, response, err

}

//UpdateEnterpriseSSID Update Enterprise SSID - c493-991f-40ca-ba44
/* Update enterprise SSID


 */
func (s *WirelessService) UpdateEnterpriseSSID(requestWirelessUpdateEnterpriseSSID *RequestWirelessUpdateEnterpriseSSID) (*ResponseWirelessUpdateEnterpriseSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateEnterpriseSSID).
		SetResult(&ResponseWirelessUpdateEnterpriseSSID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessUpdateEnterpriseSSID)
	return result, response, err

}

//UpdateWirelessProfile Update Wireless Profile - cfbd-3870-405a-ad55
/* Updates the wireless Network Profile with updated details provided. All sites to be present in the network profile should be provided.


 */
func (s *WirelessService) UpdateWirelessProfile(requestWirelessUpdateWirelessProfile *RequestWirelessUpdateWirelessProfile) (*ResponseWirelessUpdateWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateWirelessProfile).
		SetResult(&ResponseWirelessUpdateWirelessProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessUpdateWirelessProfile)
	return result, response, err

}

//ProvisionUpdate Provision update - 87a5-ab04-4139-862d
/* Updates wireless provisioning


@param ProvisionUpdateHeaderParams Custom header parameters
*/
func (s *WirelessService) ProvisionUpdate(requestWirelessProvisionUpdate *RequestWirelessProvisionUpdate, ProvisionUpdateHeaderParams *ProvisionUpdateHeaderParams) (*ResponseWirelessProvisionUpdate, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/provision"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ProvisionUpdateHeaderParams != nil {

		if ProvisionUpdateHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", ProvisionUpdateHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessProvisionUpdate).
		SetResult(&ResponseWirelessProvisionUpdate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ProvisionUpdate")
	}

	result := response.Result().(*ResponseWirelessProvisionUpdate)
	return result, response, err

}

//DeleteSSIDAndProvisionItToDevices Delete SSID and provision it to devices - fc95-38fe-43d9-884d
/* Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from DNA Center


@param ssidName ssidName path parameter.
@param managedAPLocations managedAPLocations path parameter.
@param DeleteSSIDAndProvisionItToDevicesHeaderParams Custom header parameters
*/
func (s *WirelessService) DeleteSSIDAndProvisionItToDevices(ssidName string, managedAPLocations string, DeleteSSIDAndProvisionItToDevicesHeaderParams *DeleteSSIDAndProvisionItToDevicesHeaderParams) (*ResponseWirelessDeleteSSIDAndProvisionItToDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/ssid/{ssidName}/{managedAPLocations}"
	path = strings.Replace(path, "{ssidName}", fmt.Sprintf("%v", ssidName), -1)
	path = strings.Replace(path, "{managedAPLocations}", fmt.Sprintf("%v", managedAPLocations), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if DeleteSSIDAndProvisionItToDevicesHeaderParams != nil {

		if DeleteSSIDAndProvisionItToDevicesHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", DeleteSSIDAndProvisionItToDevicesHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseWirelessDeleteSSIDAndProvisionItToDevices{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteSsidAndProvisionItToDevices")
	}

	result := response.Result().(*ResponseWirelessDeleteSSIDAndProvisionItToDevices)
	return result, response, err

}

//DeleteEnterpriseSSID Delete Enterprise SSID - c7a6-592b-4b98-a369
/* Deletes given enterprise SSID


@param ssidName ssidName path parameter. Enter the SSID name to be deleted

*/
func (s *WirelessService) DeleteEnterpriseSSID(ssidName string) (*ResponseWirelessDeleteEnterpriseSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/enterprise-ssid/{ssidName}"
	path = strings.Replace(path, "{ssidName}", fmt.Sprintf("%v", ssidName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteEnterpriseSSID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessDeleteEnterpriseSSID)
	return result, response, err

}

//DeleteWirelessProfile Delete Wireless Profile - e395-88a5-4949-82c4
/* Delete the Wireless Profile from DNAC whose name is provided.


@param wirelessProfileName wirelessProfileName path parameter. Wireless Profile Name

*/
func (s *WirelessService) DeleteWirelessProfile(wirelessProfileName string) (*ResponseWirelessDeleteWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless-profile/{wirelessProfileName}"
	path = strings.Replace(path, "{wirelessProfileName}", fmt.Sprintf("%v", wirelessProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteWirelessProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessDeleteWirelessProfile)
	return result, response, err

}

//DeleteDynamicInterface Delete dynamic interface - ffb4-abf4-44fb-b70a
/* Delete a dynamic interface


@param interfaceName interfaceName path parameter. valid interface-name to be deleted

@param DeleteDynamicInterfaceHeaderParams Custom header parameters
*/
func (s *WirelessService) DeleteDynamicInterface(interfaceName string, DeleteDynamicInterfaceHeaderParams *DeleteDynamicInterfaceHeaderParams) (*resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/dynamic-interface/{interfaceName}"
	path = strings.Replace(path, "{interfaceName}", fmt.Sprintf("%v", interfaceName), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if DeleteDynamicInterfaceHeaderParams != nil {

		if DeleteDynamicInterfaceHeaderParams.Runsync != "" {
			clientRequest = clientRequest.SetHeader("__runsync", DeleteDynamicInterfaceHeaderParams.Runsync)
		}

		if DeleteDynamicInterfaceHeaderParams.Timeout != "" {
			clientRequest = clientRequest.SetHeader("__timeout", DeleteDynamicInterfaceHeaderParams.Timeout)
		}

	}

	response, err = clientRequest.
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation DeleteDynamicInterface")
	}

	return response, err

}

//DeleteRfProfiles Delete RF profiles - 28b2-4a74-4a99-94be
/* Delete RF profile(s)


@param rfProfileName rfProfileName path parameter. RF profile name to be deleted(required) non-custom RF profile cannot be deleted

*/
func (s *WirelessService) DeleteRfProfiles(rfProfileName string) (*ResponseWirelessDeleteRfProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/rf-profile/{rfProfileName}"
	path = strings.Replace(path, "{rfProfileName}", fmt.Sprintf("%v", rfProfileName), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteRfProfiles{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteRfProfiles")
	}

	result := response.Result().(*ResponseWirelessDeleteRfProfiles)
	return result, response, err

}
