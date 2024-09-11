package dnac

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type WirelessService service

type SensorTestResultsQueryParams struct {
	SiteID        string  `url:"siteId,omitempty"`        //Assurance site UUID
	StartTime     float64 `url:"startTime,omitempty"`     //The epoch time in milliseconds
	EndTime       float64 `url:"endTime,omitempty"`       //The epoch time in milliseconds
	TestFailureBy string  `url:"testFailureBy,omitempty"` //Obtain failure statistics group by "area", "building", or "floor" (case insensitive)
}
type CreateAndProvisionSSIDHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type DeleteSSIDAndProvisionItToDevicesHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type GetAccessPointRebootTaskResultQueryParams struct {
	ParentTaskID string `url:"parentTaskId,omitempty"` //task id of ap reboot request
}
type GetEnterpriseSSIDQueryParams struct {
	SSIDName string `url:"ssidName,omitempty"` //Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
}
type GetSSIDBySiteQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Limit
	Offset float64 `url:"offset,omitempty"` //Offset
}
type GetAccessPointConfigurationQueryParams struct {
	Key string `url:"key,omitempty"` //The ethernet MAC address of Access point
}
type ApProvision2HeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string.
}
type DeleteDynamicInterfaceQueryParams struct {
	InterfaceName string `url:"interfaceName,omitempty"` //valid interface-name to be deleted
}
type DeleteDynamicInterfaceHeaderParams struct {
	Runsync string `url:"__runsync,omitempty"` //Expects type bool. Enable this parameter to execute the API and return a response synchronously
	Timeout string `url:"__timeout,omitempty"` //Expects type float64. If __runsync is set to ‘true’, this defines the maximum time before which if the API completes its execution, then a synchronous response is returned.  If the time taken for the API to complete the execution, exceeds this time, then an asynchronous response is returned with an execution id, that can be used to get the status and response associated with the API execution
}
type GetDynamicInterfaceQueryParams struct {
	InterfaceName string `url:"interface-name,omitempty"` //dynamic-interface name, if not specified all the existing dynamic interfaces will be retrieved
}
type GetWirelessProfileQueryParams struct {
	ProfileName string `url:"profileName,omitempty"` //Wireless Network Profile Name
}
type ProvisionUpdateHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type RetrieveRfProfilesQueryParams struct {
	RfProfileName string `url:"rf-profile-name,omitempty"` //RF Profile Name
}
type GetAccessPointsFactoryResetStatusQueryParams struct {
	TaskID string `url:"taskId,omitempty"` //provide the task id which is returned in the response of ap factory reset post api
}
type GetAllMobilityGroupsQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Employ this query parameter to obtain the details of the Mobility Group corresponding to the provided networkDeviceId. Obtain the network device ID value by using the API GET call /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
}
type GetAnchorManagedApLocationsForSpecificWirelessControllerQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetPrimaryManagedApLocationsForSpecificWirelessControllerQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetSecondaryManagedApLocationsForSpecificWirelessControllerQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetSSIDDetailsForSpecificWirelessControllerQueryParams struct {
	SSIDName    string  `url:"ssidName,omitempty"`    //Employ this query parameter to obtain the details of the SSID corresponding to the provided SSID name.
	AdminStatus bool    `url:"adminStatus,omitempty"` //Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed     bool    `url:"managed,omitempty"`     //If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page.
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1.
}
type GetSSIDCountForSpecificWirelessControllerQueryParams struct {
	AdminStatus bool `url:"adminStatus,omitempty"` //Utilize this query parameter to obtain the number of SSIDs according to their administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed     bool `url:"managed,omitempty"`     //If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
}
type GetWirelessProfilesQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Limit
	Offset float64 `url:"offset,omitempty"` //Offset
}
type GetAll80211BeProfilesQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Limit
	Offset float64 `url:"offset,omitempty"` //Offset
}
type GetInterfacesQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Limit
	Offset float64 `url:"offset,omitempty"` //Offset
}
type GetRfProfilesQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Limit
	Offset float64 `url:"offset,omitempty"` //Offset
}

type ResponseWirelessSensorTestResults struct {
	Version  string                                     `json:"version,omitempty"`  // Version
	Response *ResponseWirelessSensorTestResultsResponse `json:"response,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponse struct {
	Summary      *ResponseWirelessSensorTestResultsResponseSummary        `json:"summary,omitempty"`      //
	FailureStats *[]ResponseWirelessSensorTestResultsResponseFailureStats `json:"failureStats,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummary struct {
	TotalTestCount  *int                                                             `json:"totalTestCount,omitempty"`   // Total test count
	OnBoarding      *ResponseWirelessSensorTestResultsResponseSummaryOnBoarding      `json:"ONBOARDING,omitempty"`       //
	PERfORMAncE     *ResponseWirelessSensorTestResultsResponseSummaryPERfORMAncE     `json:"PERFORMANCE,omitempty"`      //
	NETWORKSERVICES *ResponseWirelessSensorTestResultsResponseSummaryNETWORKSERVICES `json:"NETWORK_SERVICES,omitempty"` //
	ApPCONNECTIVITY *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITY `json:"APP_CONNECTIVITY,omitempty"` //
	RfASSESSMENT    *ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENT    `json:"RF_ASSESSMENT,omitempty"`    //
	Email           *ResponseWirelessSensorTestResultsResponseSummaryEmail           `json:"EMAIL,omitempty"`            //
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoarding struct {
	Auth  *ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAuth  `json:"AUTH,omitempty"`  //
	DHCP  *ResponseWirelessSensorTestResultsResponseSummaryOnBoardingDHCP  `json:"DHCP,omitempty"`  //
	Assoc *ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAssoc `json:"ASSOC,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAuth struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count
	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoardingDHCP struct {
	PassCount *int     `json:"passCount,omitempty"` // Total passed test count
	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAssoc struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count
	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryPERfORMAncE struct {
	IPSLASENDER *ResponseWirelessSensorTestResultsResponseSummaryPERfORMAncEIPSLASENDER `json:"IPSLASENDER,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryPERfORMAncEIPSLASENDER struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count
	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryNETWORKSERVICES struct {
	DNS *ResponseWirelessSensorTestResultsResponseSummaryNETWORKSERVICESDNS `json:"DNS,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryNETWORKSERVICESDNS struct {
	PassCount *int     `json:"passCount,omitempty"` // Total passed test count
	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITY struct {
	HOSTREACHABILITY *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYHOSTREACHABILITY `json:"HOST_REACHABILITY,omitempty"` //
	WebServer        *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYWebServer        `json:"WEBSERVER,omitempty"`         //
	FileTransfer     *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYFileTransfer     `json:"FILETRANSFER,omitempty"`      //
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYHOSTREACHABILITY struct {
	PassCount *int     `json:"passCount,omitempty"` // Total passed test count
	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYWebServer struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count
	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYFileTransfer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Total passed test count
	FailCount *int     `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENT struct {
	DATARATE *ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTDATARATE `json:"DATA_RATE,omitempty"` //
	SNR      *ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTSNR      `json:"SNR,omitempty"`       //
}
type ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTDATARATE struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count
	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTSNR struct {
	PassCount *int     `json:"passCount,omitempty"` // Total passed test count
	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryEmail struct {
	MailServer *ResponseWirelessSensorTestResultsResponseSummaryEmailMailServer `json:"MAILSERVER,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryEmailMailServer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Total passed test count
	FailCount *int     `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseFailureStats struct {
	ErrorCode    *int   `json:"errorCode,omitempty"`    // The error code
	ErrorTitle   string `json:"errorTitle,omitempty"`   // The error title
	TestType     string `json:"testType,omitempty"`     // The test type
	TestCategory string `json:"testCategory,omitempty"` // The test category
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
type ResponseWirelessRebootAccessPoints struct {
	Response *ResponseWirelessRebootAccessPointsResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseWirelessRebootAccessPointsResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseWirelessGetAccessPointRebootTaskResult []ResponseItemWirelessGetAccessPointRebootTaskResult // Array of ResponseWirelessGetAccessPointRebootTaskResult
type ResponseItemWirelessGetAccessPointRebootTaskResult struct {
	WlcIP  string                                                      `json:"wlcIP,omitempty"`  //
	ApList *[]ResponseItemWirelessGetAccessPointRebootTaskResultApList `json:"apList,omitempty"` //
}
type ResponseItemWirelessGetAccessPointRebootTaskResultApList struct {
	ApName        string                                                                 `json:"apName,omitempty"`        //
	RebootStatus  string                                                                 `json:"rebootStatus,omitempty"`  //
	FailureReason *ResponseItemWirelessGetAccessPointRebootTaskResultApListFailureReason `json:"failureReason,omitempty"` //
}
type ResponseItemWirelessGetAccessPointRebootTaskResultApListFailureReason interface{}
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
	Name                             string                                                              `json:"name,omitempty"`                             // SSID Name
	WLANType                         string                                                              `json:"wlanType,omitempty"`                         // Wlan Type
	EnableFastLane                   *bool                                                               `json:"enableFastLane,omitempty"`                   // Enable Fast Lane
	SecurityLevel                    string                                                              `json:"securityLevel,omitempty"`                    // Security Level
	AuthServer                       string                                                              `json:"authServer,omitempty"`                       // Auth Server
	Passphrase                       string                                                              `json:"passphrase,omitempty"`                       // Passphrase
	TrafficType                      string                                                              `json:"trafficType,omitempty"`                      // Traffic Type
	EnableMacFiltering               *bool                                                               `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	IsEnabled                        *bool                                                               `json:"isEnabled,omitempty"`                        // Is Enabled
	IsFabric                         *bool                                                               `json:"isFabric,omitempty"`                         // Is Fabric
	FastTransition                   string                                                              `json:"fastTransition,omitempty"`                   // Fast Transition
	RadioPolicy                      string                                                              `json:"radioPolicy,omitempty"`                      // Radio Policy
	EnableBroadcastSSID              *bool                                                               `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcast SSID
	NasOptions                       []string                                                            `json:"nasOptions,omitempty"`                       // Nas Options
	AAAOverride                      *bool                                                               `json:"aaaOverride,omitempty"`                      // Aaa Override
	CoverageHoleDetectionEnable      *bool                                                               `json:"coverageHoleDetectionEnable,omitempty"`      // Coverage Hole Detection Enable
	ProtectedManagementFrame         string                                                              `json:"protectedManagementFrame,omitempty"`         // Protected Management Frame
	MultipSKSettings                 *[]ResponseItemWirelessGetEnterpriseSSIDSSIDDetailsMultipSKSettings `json:"multiPSKSettings,omitempty"`                 //
	ClientRateLimit                  *float64                                                            `json:"clientRateLimit,omitempty"`                  // Client Rate Limit. (in bits per second)
	EnableSessionTimeOut             *bool                                                               `json:"enableSessionTimeOut,omitempty"`             // Enable Session Time Out
	SessionTimeOut                   *float64                                                            `json:"sessionTimeOut,omitempty"`                   // sessionTimeOut
	EnableClientExclusion            *bool                                                               `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *float64                                                            `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool                                                               `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *float64                                                            `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set ClientIdle Timeout
	EnableDirectedMulticastService   *bool                                                               `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed MulticastService
	EnableNeighborList               *bool                                                               `json:"enableNeighborList,omitempty"`               // Enable NeighborList
	MfpClientProtection              string                                                              `json:"mfpClientProtection,omitempty"`              // Mfp Client Protection
}
type ResponseItemWirelessGetEnterpriseSSIDSSIDDetailsMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase
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
type ResponseWirelessCreateSSID struct {
	Response *ResponseWirelessCreateSSIDResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version
}
type ResponseWirelessCreateSSIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetSSIDBySite struct {
	Response *[]ResponseWirelessGetSSIDBySiteResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version of the response
}
type ResponseWirelessGetSSIDBySiteResponse struct {
	SSID                                         string                                                   `json:"ssid,omitempty"`                                         // Name of the SSID
	AuthType                                     string                                                   `json:"authType,omitempty"`                                     // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
	Passphrase                                   string                                                   `json:"passphrase,omitempty"`                                   // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
	IsFastLaneEnabled                            *bool                                                    `json:"isFastLaneEnabled,omitempty"`                            // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
	IsMacFilteringEnabled                        *bool                                                    `json:"isMacFilteringEnabled,omitempty"`                        // True if MAC Filtering is enabled, else False
	SSIDRadioType                                string                                                   `json:"ssidRadioType,omitempty"`                                // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
	IsBroadcastSSID                              *bool                                                    `json:"isBroadcastSSID,omitempty"`                              // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
	FastTransition                               string                                                   `json:"fastTransition,omitempty"`                               // Fast Transition
	SessionTimeOutEnable                         *bool                                                    `json:"sessionTimeOutEnable,omitempty"`                         // Turn on the feature that imposes a time limit on user sessions
	SessionTimeOut                               *int                                                     `json:"sessionTimeOut,omitempty"`                               // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity
	ClientExclusionEnable                        *bool                                                    `json:"clientExclusionEnable,omitempty"`                        // Activate the feature that allows for the exclusion of clients
	ClientExclusionTimeout                       *int                                                     `json:"clientExclusionTimeout,omitempty"`                       // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
	BasicServiceSetMaxIDleEnable                 *bool                                                    `json:"basicServiceSetMaxIdleEnable,omitempty"`                 // Activate the maximum idle feature for the Basic Service Set
	BasicServiceSetClientIDleTimeout             *int                                                     `json:"basicServiceSetClientIdleTimeout,omitempty"`             // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out
	DirectedMulticastServiceEnable               *bool                                                    `json:"directedMulticastServiceEnable,omitempty"`               // The Directed Multicast Service feature becomes operational when it is set to true
	NeighborListEnable                           *bool                                                    `json:"neighborListEnable,omitempty"`                           // The Neighbor List feature is enabled when it is set to true
	ManagementFrameProtectionClientprotection    string                                                   `json:"managementFrameProtectionClientprotection,omitempty"`    // Management Frame Protection Client
	NasOptions                                   []string                                                 `json:"nasOptions,omitempty"`                                   // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
	ProfileName                                  string                                                   `json:"profileName,omitempty"`                                  // WLAN Profile Name, if not passed autogenerated profile name will be assigned
	PolicyProfileName                            string                                                   `json:"policyProfileName,omitempty"`                            // Policy Profile Name. If not passed, profileName value will be used to populate this parameter
	AAAOverride                                  *bool                                                    `json:"aaaOverride,omitempty"`                                  // Activate the AAA Override feature when set to true
	CoverageHoleDetectionEnable                  *bool                                                    `json:"coverageHoleDetectionEnable,omitempty"`                  // Activate Coverage Hole Detection feature when set to true
	ProtectedManagementFrame                     string                                                   `json:"protectedManagementFrame,omitempty"`                     // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
	MultipSKSettings                             *[]ResponseWirelessGetSSIDBySiteResponseMultipSKSettings `json:"multiPSKSettings,omitempty"`                             //
	ClientRateLimit                              *int                                                     `json:"clientRateLimit,omitempty"`                              // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
	RsnCipherSuiteGcmp256                        *bool                                                    `json:"rsnCipherSuiteGcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
	RsnCipherSuiteCcmp256                        *bool                                                    `json:"rsnCipherSuiteCcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
	RsnCipherSuiteGcmp128                        *bool                                                    `json:"rsnCipherSuiteGcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated
	RsnCipherSuiteCcmp128                        *bool                                                    `json:"rsnCipherSuiteCcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
	Ghz6PolicyClientSteering                     *bool                                                    `json:"ghz6PolicyClientSteering,omitempty"`                     // True if 6 GHz Policy Client Steering is enabled, else False
	IsAuthKey8021X                               *bool                                                    `json:"isAuthKey8021x,omitempty"`                               // When set to true, the 802.1X authentication key is in use
	IsAuthKey8021XPlusFT                         *bool                                                    `json:"isAuthKey8021xPlusFT,omitempty"`                         // When set to true, the 802.1X-Plus-FT authentication key is in use
	IsAuthKey8021XSHA256                         *bool                                                    `json:"isAuthKey8021x_SHA256,omitempty"`                        // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
	IsAuthKeySae                                 *bool                                                    `json:"isAuthKeySae,omitempty"`                                 // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
	IsAuthKeySaePlusFT                           *bool                                                    `json:"isAuthKeySaePlusFT,omitempty"`                           // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
	IsAuthKeyPSK                                 *bool                                                    `json:"isAuthKeyPSK,omitempty"`                                 // When set to true, the Pre-shared Key (PSK) authentication feature is enabled
	IsAuthKeyPSKPlusFT                           *bool                                                    `json:"isAuthKeyPSKPlusFT,omitempty"`                           // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
	IsAuthKeyOWE                                 *bool                                                    `json:"isAuthKeyOWE,omitempty"`                                 // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
	IsAuthKeyEasyPSK                             *bool                                                    `json:"isAuthKeyEasyPSK,omitempty"`                             // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
	IsAuthKeyPSKSHA256                           *bool                                                    `json:"isAuthKeyPSKSHA256,omitempty"`                           // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
	OpenSSID                                     string                                                   `json:"openSsid,omitempty"`                                     // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
	IsCustomNasIDOptions                         *bool                                                    `json:"isCustomNasIdOptions,omitempty"`                         // Set to true if Custom NAS ID Options provided
	WLANBandSelectEnable                         *bool                                                    `json:"wlanBandSelectEnable,omitempty"`                         // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
	IsEnabled                                    *bool                                                    `json:"isEnabled,omitempty"`                                    // Set SSID's admin status as 'Enabled' when set to true
	AuthServers                                  []string                                                 `json:"authServers,omitempty"`                                  // List of Authentication/Authorization server IpAddresses
	AcctServers                                  []string                                                 `json:"acctServers,omitempty"`                                  // List of Accounting server IpAddresses
	EgressQos                                    string                                                   `json:"egressQos,omitempty"`                                    // Egress QOS
	IngressQos                                   string                                                   `json:"ingressQos,omitempty"`                                   // Ingress QOS
	InheritedSiteID                              string                                                   `json:"inheritedSiteId,omitempty"`                              // Site UUID from where the SSID is inherited
	InheritedSiteName                            string                                                   `json:"inheritedSiteName,omitempty"`                            // Site Name from where the SSID is inherited
	WLANType                                     string                                                   `json:"wlanType,omitempty"`                                     // Wlan Type
	L3AuthType                                   string                                                   `json:"l3AuthType,omitempty"`                                   // L3 Authentication Type
	AuthServer                                   string                                                   `json:"authServer,omitempty"`                                   // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth
	ExternalAuthIPAddress                        string                                                   `json:"externalAuthIpAddress,omitempty"`                        // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
	WebPassthrough                               *bool                                                    `json:"webPassthrough,omitempty"`                               // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
	SleepingClientEnable                         *bool                                                    `json:"sleepingClientEnable,omitempty"`                         // When set to true, this will activate the timeout settings that apply to clients in sleep mode
	SleepingClientTimeout                        *int                                                     `json:"sleepingClientTimeout,omitempty"`                        // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
	ACLName                                      string                                                   `json:"aclName,omitempty"`                                      // Pre-Auth Access Control List (ACL) Name
	IsPosturingEnabled                           *bool                                                    `json:"isPosturingEnabled,omitempty"`                           // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
	IsAuthKeySuiteB1X                            *bool                                                    `json:"isAuthKeySuiteB1x,omitempty"`                            // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
	IsAuthKeySuiteB1921X                         *bool                                                    `json:"isAuthKeySuiteB1921x,omitempty"`                         // When set to true, the SuiteB192-1x authentication key feature is enabled.
	IsAuthKeySaeExt                              *bool                                                    `json:"isAuthKeySaeExt,omitempty"`                              // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
	IsAuthKeySaeExtPlusFT                        *bool                                                    `json:"isAuthKeySaeExtPlusFT,omitempty"`                        // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
	IsApBeaconProtectionEnabled                  *bool                                                    `json:"isApBeaconProtectionEnabled,omitempty"`                  // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
	Ghz24Policy                                  string                                                   `json:"ghz24Policy,omitempty"`                                  // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
	CckmTsfTolerance                             *int                                                     `json:"cckmTsfTolerance,omitempty"`                             // Cckm TImestamp Tolerance(in milliseconds)
	IsCckmEnabled                                *bool                                                    `json:"isCckmEnabled,omitempty"`                                // True if CCKM is enabled, else False
	IsHex                                        *bool                                                    `json:"isHex,omitempty"`                                        // True if passphrase is in Hex format, else False.
	IsSensorPnp                                  *bool                                                    `json:"isSensorPnp,omitempty"`                                  // True if SSID is a sensor SSID
	ID                                           string                                                   `json:"id,omitempty"`                                           // SSID ID
	IsRandomMacFilterEnabled                     *bool                                                    `json:"isRandomMacFilterEnabled,omitempty"`                     // Deny clients using randomized MAC addresses when set to true
	FastTransitionOverTheDistributedSystemEnable *bool                                                    `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type ResponseWirelessGetSSIDBySiteResponseMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type ResponseWirelessGetSSIDCountBySite struct {
	Response *ResponseWirelessGetSSIDCountBySiteResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGetSSIDCountBySiteResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetSSIDByID struct {
	Response *ResponseWirelessGetSSIDByIDResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  // Version of the response
}
type ResponseWirelessGetSSIDByIDResponse struct {
	SSID                                         string                                                 `json:"ssid,omitempty"`                                         // Name of the SSID
	AuthType                                     string                                                 `json:"authType,omitempty"`                                     // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
	Passphrase                                   string                                                 `json:"passphrase,omitempty"`                                   // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
	IsFastLaneEnabled                            *bool                                                  `json:"isFastLaneEnabled,omitempty"`                            // True if FastLane is enabled, else False
	IsMacFilteringEnabled                        *bool                                                  `json:"isMacFilteringEnabled,omitempty"`                        // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
	SSIDRadioType                                string                                                 `json:"ssidRadioType,omitempty"`                                // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
	IsBroadcastSSID                              *bool                                                  `json:"isBroadcastSSID,omitempty"`                              // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
	FastTransition                               string                                                 `json:"fastTransition,omitempty"`                               // Fast Transition
	SessionTimeOutEnable                         *bool                                                  `json:"sessionTimeOutEnable,omitempty"`                         // Turn on the feature that imposes a time limit on user sessions
	SessionTimeOut                               *int                                                   `json:"sessionTimeOut,omitempty"`                               // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity
	ClientExclusionEnable                        *bool                                                  `json:"clientExclusionEnable,omitempty"`                        // Activate the feature that allows for the exclusion of clients
	ClientExclusionTimeout                       *int                                                   `json:"clientExclusionTimeout,omitempty"`                       // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
	BasicServiceSetMaxIDleEnable                 *bool                                                  `json:"basicServiceSetMaxIdleEnable,omitempty"`                 // Activate the maximum idle feature for the Basic Service Set
	BasicServiceSetClientIDleTimeout             *int                                                   `json:"basicServiceSetClientIdleTimeout,omitempty"`             // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out
	DirectedMulticastServiceEnable               *bool                                                  `json:"directedMulticastServiceEnable,omitempty"`               // The Directed Multicast Service feature becomes operational when it is set to true
	NeighborListEnable                           *bool                                                  `json:"neighborListEnable,omitempty"`                           // The Neighbor List feature is enabled when it is set to true
	ManagementFrameProtectionClientprotection    string                                                 `json:"managementFrameProtectionClientprotection,omitempty"`    // Management Frame Protection Client
	NasOptions                                   []string                                               `json:"nasOptions,omitempty"`                                   // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
	ProfileName                                  string                                                 `json:"profileName,omitempty"`                                  // WLAN Profile Name, if not passed autogenerated profile name will be assigned
	PolicyProfileName                            string                                                 `json:"policyProfileName,omitempty"`                            // Policy Profile Name. If not passed, profileName value will be used to populate this parameter
	AAAOverride                                  *bool                                                  `json:"aaaOverride,omitempty"`                                  // Activate the AAA Override feature when set to true
	CoverageHoleDetectionEnable                  *bool                                                  `json:"coverageHoleDetectionEnable,omitempty"`                  // Activate Coverage Hole Detection feature when set to true
	ProtectedManagementFrame                     string                                                 `json:"protectedManagementFrame,omitempty"`                     // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
	MultipSKSettings                             *[]ResponseWirelessGetSSIDByIDResponseMultipSKSettings `json:"multiPSKSettings,omitempty"`                             //
	ClientRateLimit                              *int                                                   `json:"clientRateLimit,omitempty"`                              // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
	RsnCipherSuiteGcmp256                        *bool                                                  `json:"rsnCipherSuiteGcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
	RsnCipherSuiteCcmp256                        *bool                                                  `json:"rsnCipherSuiteCcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
	RsnCipherSuiteGcmp128                        *bool                                                  `json:"rsnCipherSuiteGcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activatedTrue if RSN Cipher Suite GCMP128 is enabled, else False
	RsnCipherSuiteCcmp128                        *bool                                                  `json:"rsnCipherSuiteCcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
	Ghz6PolicyClientSteering                     *bool                                                  `json:"ghz6PolicyClientSteering,omitempty"`                     // True if 6 GHz Policy Client Steering is enabled, else False
	IsAuthKey8021X                               *bool                                                  `json:"isAuthKey8021x,omitempty"`                               // When set to true, the 802.1X authentication key is in use
	IsAuthKey8021XPlusFT                         *bool                                                  `json:"isAuthKey8021xPlusFT,omitempty"`                         // When set to true, the 802.1X-Plus-FT authentication key is in use
	IsAuthKey8021XSHA256                         *bool                                                  `json:"isAuthKey8021x_SHA256,omitempty"`                        // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
	IsAuthKeySae                                 *bool                                                  `json:"isAuthKeySae,omitempty"`                                 // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
	IsAuthKeySaePlusFT                           *bool                                                  `json:"isAuthKeySaePlusFT,omitempty"`                           // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
	IsAuthKeyPSK                                 *bool                                                  `json:"isAuthKeyPSK,omitempty"`                                 // When set to true, the Pre-shared Key (PSK) authentication feature is enabled
	IsAuthKeyPSKPlusFT                           *bool                                                  `json:"isAuthKeyPSKPlusFT,omitempty"`                           // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
	IsAuthKeyOWE                                 *bool                                                  `json:"isAuthKeyOWE,omitempty"`                                 // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
	IsAuthKeyEasyPSK                             *bool                                                  `json:"isAuthKeyEasyPSK,omitempty"`                             // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
	IsAuthKeyPSKSHA256                           *bool                                                  `json:"isAuthKeyPSKSHA256,omitempty"`                           // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
	OpenSSID                                     string                                                 `json:"openSsid,omitempty"`                                     // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
	IsCustomNasIDOptions                         *bool                                                  `json:"isCustomNasIdOptions,omitempty"`                         // Set to true if Custom NAS ID Options provided
	WLANBandSelectEnable                         *bool                                                  `json:"wlanBandSelectEnable,omitempty"`                         // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
	IsEnabled                                    *bool                                                  `json:"isEnabled,omitempty"`                                    // Set SSID's admin status as 'Enabled' when set to true
	AuthServers                                  []string                                               `json:"authServers,omitempty"`                                  // List of Authentication/Authorization server IpAddresses
	AcctServers                                  []string                                               `json:"acctServers,omitempty"`                                  // List of Accounting server IpAddresses
	EgressQos                                    string                                                 `json:"egressQos,omitempty"`                                    // Egress QOS
	IngressQos                                   string                                                 `json:"ingressQos,omitempty"`                                   // Ingress QOS
	InheritedSiteID                              string                                                 `json:"inheritedSiteId,omitempty"`                              // Site UUID from where the SSID is inherited
	InheritedSiteName                            string                                                 `json:"inheritedSiteName,omitempty"`                            // Site Name from where the SSID is inherited
	WLANType                                     string                                                 `json:"wlanType,omitempty"`                                     // Wlan Type
	L3AuthType                                   string                                                 `json:"l3AuthType,omitempty"`                                   // L3 Authentication Type
	AuthServer                                   string                                                 `json:"authServer,omitempty"`                                   // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth
	ExternalAuthIPAddress                        string                                                 `json:"externalAuthIpAddress,omitempty"`                        // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
	WebPassthrough                               *bool                                                  `json:"webPassthrough,omitempty"`                               // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
	SleepingClientEnable                         *bool                                                  `json:"sleepingClientEnable,omitempty"`                         // When set to true, this will activate the timeout settings that apply to clients in sleep mode
	SleepingClientTimeout                        *int                                                   `json:"sleepingClientTimeout,omitempty"`                        // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
	ACLName                                      string                                                 `json:"aclName,omitempty"`                                      // Pre-Auth Access Control List (ACL) Name
	IsPosturingEnabled                           *bool                                                  `json:"isPosturingEnabled,omitempty"`                           // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
	IsAuthKeySuiteB1X                            *bool                                                  `json:"isAuthKeySuiteB1x,omitempty"`                            // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
	IsAuthKeySuiteB1921X                         *bool                                                  `json:"isAuthKeySuiteB1921x,omitempty"`                         // When set to true, the SuiteB192-1x authentication key feature is enabled.
	IsAuthKeySaeExt                              *bool                                                  `json:"isAuthKeySaeExt,omitempty"`                              // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
	IsAuthKeySaeExtPlusFT                        *bool                                                  `json:"isAuthKeySaeExtPlusFT,omitempty"`                        // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
	IsApBeaconProtectionEnabled                  *bool                                                  `json:"isApBeaconProtectionEnabled,omitempty"`                  // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
	Ghz24Policy                                  string                                                 `json:"ghz24Policy,omitempty"`                                  // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
	CckmTsfTolerance                             *int                                                   `json:"cckmTsfTolerance,omitempty"`                             // Cckm TImestamp Tolerance(in milliseconds)
	IsCckmEnabled                                *bool                                                  `json:"isCckmEnabled,omitempty"`                                // True if CCKM is enabled, else False
	IsHex                                        *bool                                                  `json:"isHex,omitempty"`                                        // True if passphrase is in Hex format, else False.
	IsSensorPnp                                  *bool                                                  `json:"isSensorPnp,omitempty"`                                  // True if SSID is a sensor SSID
	ID                                           string                                                 `json:"id,omitempty"`                                           // SSID ID
	IsRandomMacFilterEnabled                     *bool                                                  `json:"isRandomMacFilterEnabled,omitempty"`                     // Deny clients using randomized MAC addresses when set to true
	FastTransitionOverTheDistributedSystemEnable *bool                                                  `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type ResponseWirelessGetSSIDByIDResponseMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type ResponseWirelessUpdateSSID struct {
	Response *ResponseWirelessUpdateSSIDResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version
}
type ResponseWirelessUpdateSSIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessDeleteSSID struct {
	Response *ResponseWirelessDeleteSSIDResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  // Version
}
type ResponseWirelessDeleteSSIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessDeleteWirelessProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessConfigureAccessPointsV1 struct {
	Response *ResponseWirelessConfigureAccessPointsV1Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseWirelessConfigureAccessPointsV1Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseWirelessGetAccessPointConfigurationTaskResult []ResponseItemWirelessGetAccessPointConfigurationTaskResult // Array of ResponseWirelessGetAccessPointConfigurationTaskResult
type ResponseItemWirelessGetAccessPointConfigurationTaskResult struct {
	InstanceUUID           *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUUID           `json:"instanceUuid,omitempty"`            //
	InstanceID             *float64                                                                         `json:"instanceId,omitempty"`              //
	AuthEntityID           *ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityID           `json:"authEntityId,omitempty"`            //
	DisplayName            string                                                                           `json:"displayName,omitempty"`             //
	AuthEntityClass        *ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityClass        `json:"authEntityClass,omitempty"`         //
	InstanceTenantID       string                                                                           `json:"instanceTenantId,omitempty"`        //
	OrderedListOEIndex     *float64                                                                         `json:"_orderedListOEIndex,omitempty"`     //
	OrderedListOEAssocName *ResponseItemWirelessGetAccessPointConfigurationTaskResultOrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //
	CreationOrderIndex     *float64                                                                         `json:"_creationOrderIndex,omitempty"`     //
	IsBeingChanged         *bool                                                                            `json:"_isBeingChanged,omitempty"`         //
	DeployPending          string                                                                           `json:"deployPending,omitempty"`           //
	InstanceCreatedOn      *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceCreatedOn      `json:"instanceCreatedOn,omitempty"`       //
	InstanceUpdatedOn      *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUpdatedOn      `json:"instanceUpdatedOn,omitempty"`       //
	ChangeLogList          *ResponseItemWirelessGetAccessPointConfigurationTaskResultChangeLogList          `json:"changeLogList,omitempty"`           //
	InstanceOrigin         *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceOrigin         `json:"instanceOrigin,omitempty"`          //
	LazyLoadedEntities     *ResponseItemWirelessGetAccessPointConfigurationTaskResultLazyLoadedEntities     `json:"lazyLoadedEntities,omitempty"`      //
	InstanceVersion        *float64                                                                         `json:"instanceVersion,omitempty"`         //
	ApName                 string                                                                           `json:"apName,omitempty"`                  //
	ControllerName         string                                                                           `json:"controllerName,omitempty"`          //
	LocationHeirarchy      string                                                                           `json:"locationHeirarchy,omitempty"`       //
	MacAddress             string                                                                           `json:"macAddress,omitempty"`              //
	Status                 string                                                                           `json:"status,omitempty"`                  //
	StatusDetails          string                                                                           `json:"statusDetails,omitempty"`           //
	InternalKey            *ResponseItemWirelessGetAccessPointConfigurationTaskResultInternalKey            `json:"internalKey,omitempty"`             //
}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUUID interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityID interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityClass interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultOrderedListOEAssocName interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceCreatedOn interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUpdatedOn interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultChangeLogList interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceOrigin interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultLazyLoadedEntities interface{}
type ResponseItemWirelessGetAccessPointConfigurationTaskResultInternalKey struct {
	Type     string   `json:"type,omitempty"`     //
	ID       *float64 `json:"id,omitempty"`       //
	LongType string   `json:"longType,omitempty"` //
	URL      string   `json:"url,omitempty"`      //
}
type ResponseWirelessGetAccessPointConfiguration struct {
	InstanceUUID            *ResponseWirelessGetAccessPointConfigurationInstanceUUID           `json:"instanceUuid,omitempty"`            //
	InstanceID              *float64                                                           `json:"instanceId,omitempty"`              //
	AuthEntityID            *ResponseWirelessGetAccessPointConfigurationAuthEntityID           `json:"authEntityId,omitempty"`            //
	DisplayName             string                                                             `json:"displayName,omitempty"`             //
	AuthEntityClass         *ResponseWirelessGetAccessPointConfigurationAuthEntityClass        `json:"authEntityClass,omitempty"`         //
	InstanceTenantID        string                                                             `json:"instanceTenantId,omitempty"`        //
	OrderedListOEIndex      *float64                                                           `json:"_orderedListOEIndex,omitempty"`     //
	OrderedListOEAssocName  *ResponseWirelessGetAccessPointConfigurationOrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //
	CreationOrderIndex      *float64                                                           `json:"_creationOrderIndex,omitempty"`     //
	IsBeingChanged          *bool                                                              `json:"_isBeingChanged,omitempty"`         //
	DeployPending           string                                                             `json:"deployPending,omitempty"`           //
	InstanceCreatedOn       *ResponseWirelessGetAccessPointConfigurationInstanceCreatedOn      `json:"instanceCreatedOn,omitempty"`       //
	InstanceUpdatedOn       *ResponseWirelessGetAccessPointConfigurationInstanceUpdatedOn      `json:"instanceUpdatedOn,omitempty"`       //
	ChangeLogList           *ResponseWirelessGetAccessPointConfigurationChangeLogList          `json:"changeLogList,omitempty"`           //
	InstanceOrigin          *ResponseWirelessGetAccessPointConfigurationInstanceOrigin         `json:"instanceOrigin,omitempty"`          //
	LazyLoadedEntities      *ResponseWirelessGetAccessPointConfigurationLazyLoadedEntities     `json:"lazyLoadedEntities,omitempty"`      //
	InstanceVersion         *float64                                                           `json:"instanceVersion,omitempty"`         //
	AdminStatus             string                                                             `json:"adminStatus,omitempty"`             //
	ApHeight                *float64                                                           `json:"apHeight,omitempty"`                //
	ApMode                  string                                                             `json:"apMode,omitempty"`                  //
	ApName                  string                                                             `json:"apName,omitempty"`                  //
	EthMac                  string                                                             `json:"ethMac,omitempty"`                  //
	FailoverPriority        string                                                             `json:"failoverPriority,omitempty"`        //
	LedBrightnessLevel      *int                                                               `json:"ledBrightnessLevel,omitempty"`      //
	LedStatus               string                                                             `json:"ledStatus,omitempty"`               //
	Location                string                                                             `json:"location,omitempty"`                //
	MacAddress              string                                                             `json:"macAddress,omitempty"`              //
	PrimaryControllerName   string                                                             `json:"primaryControllerName,omitempty"`   //
	PrimaryIPAddress        string                                                             `json:"primaryIpAddress,omitempty"`        //
	SecondaryControllerName string                                                             `json:"secondaryControllerName,omitempty"` //
	SecondaryIPAddress      string                                                             `json:"secondaryIpAddress,omitempty"`      //
	TertiaryControllerName  string                                                             `json:"tertiaryControllerName,omitempty"`  //
	TertiaryIPAddress       string                                                             `json:"tertiaryIpAddress,omitempty"`       //
	MeshDTOs                *[]ResponseWirelessGetAccessPointConfigurationMeshDTOs             `json:"meshDTOs,omitempty"`                //
	RadioDTOs               *[]ResponseWirelessGetAccessPointConfigurationRadioDTOs            `json:"radioDTOs,omitempty"`               //
	InternalKey             *ResponseWirelessGetAccessPointConfigurationInternalKey            `json:"internalKey,omitempty"`             //
}
type ResponseWirelessGetAccessPointConfigurationInstanceUUID interface{}
type ResponseWirelessGetAccessPointConfigurationAuthEntityID interface{}
type ResponseWirelessGetAccessPointConfigurationAuthEntityClass interface{}
type ResponseWirelessGetAccessPointConfigurationOrderedListOEAssocName interface{}
type ResponseWirelessGetAccessPointConfigurationInstanceCreatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationInstanceUpdatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationChangeLogList interface{}
type ResponseWirelessGetAccessPointConfigurationInstanceOrigin interface{}
type ResponseWirelessGetAccessPointConfigurationLazyLoadedEntities interface{}
type ResponseWirelessGetAccessPointConfigurationMeshDTOs interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOs struct {
	InstanceUUID           *ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceUUID           `json:"instanceUuid,omitempty"`            //
	InstanceID             *float64                                                                    `json:"instanceId,omitempty"`              //
	AuthEntityID           *ResponseWirelessGetAccessPointConfigurationRadioDTOsAuthEntityID           `json:"authEntityId,omitempty"`            //
	DisplayName            string                                                                      `json:"displayName,omitempty"`             //
	AuthEntityClass        *ResponseWirelessGetAccessPointConfigurationRadioDTOsAuthEntityClass        `json:"authEntityClass,omitempty"`         //
	InstanceTenantID       string                                                                      `json:"instanceTenantId,omitempty"`        //
	OrderedListOEIndex     *float64                                                                    `json:"_orderedListOEIndex,omitempty"`     //
	OrderedListOEAssocName *ResponseWirelessGetAccessPointConfigurationRadioDTOsOrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //
	CreationOrderIndex     *float64                                                                    `json:"_creationOrderIndex,omitempty"`     //
	IsBeingChanged         *bool                                                                       `json:"_isBeingChanged,omitempty"`         //
	DeployPending          string                                                                      `json:"deployPending,omitempty"`           //
	InstanceCreatedOn      *ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceCreatedOn      `json:"instanceCreatedOn,omitempty"`       //
	InstanceUpdatedOn      *ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceUpdatedOn      `json:"instanceUpdatedOn,omitempty"`       //
	ChangeLogList          *ResponseWirelessGetAccessPointConfigurationRadioDTOsChangeLogList          `json:"changeLogList,omitempty"`           //
	InstanceOrigin         *ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceOrigin         `json:"instanceOrigin,omitempty"`          //
	LazyLoadedEntities     *ResponseWirelessGetAccessPointConfigurationRadioDTOsLazyLoadedEntities     `json:"lazyLoadedEntities,omitempty"`      //
	InstanceVersion        *float64                                                                    `json:"instanceVersion,omitempty"`         //
	AdminStatus            string                                                                      `json:"adminStatus,omitempty"`             //
	AntennaAngle           *float64                                                                    `json:"antennaAngle,omitempty"`            //
	AntennaElevAngle       *float64                                                                    `json:"antennaElevAngle,omitempty"`        //
	AntennaGain            *int                                                                        `json:"antennaGain,omitempty"`             //
	AntennaPatternName     string                                                                      `json:"antennaPatternName,omitempty"`      //
	ChannelAssignmentMode  string                                                                      `json:"channelAssignmentMode,omitempty"`   //
	ChannelNumber          *int                                                                        `json:"channelNumber,omitempty"`           //
	ChannelWidth           string                                                                      `json:"channelWidth,omitempty"`            //
	CleanAirSI             string                                                                      `json:"cleanAirSI,omitempty"`              //
	IfType                 *int                                                                        `json:"ifType,omitempty"`                  //
	IfTypeValue            string                                                                      `json:"ifTypeValue,omitempty"`             //
	MacAddress             string                                                                      `json:"macAddress,omitempty"`              //
	PowerAssignmentMode    string                                                                      `json:"powerAssignmentMode,omitempty"`     //
	Powerlevel             *int                                                                        `json:"powerlevel,omitempty"`              //
	RadioBand              *ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioBand              `json:"radioBand,omitempty"`               //
	RadioRoleAssignment    *ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioRoleAssignment    `json:"radioRoleAssignment,omitempty"`     //
	SlotID                 *int                                                                        `json:"slotId,omitempty"`                  //
	InternalKey            *ResponseWirelessGetAccessPointConfigurationRadioDTOsInternalKey            `json:"internalKey,omitempty"`             //
}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceUUID interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsAuthEntityID interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsAuthEntityClass interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsOrderedListOEAssocName interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceCreatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceUpdatedOn interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsChangeLogList interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsInstanceOrigin interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsLazyLoadedEntities interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioBand interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioRoleAssignment interface{}
type ResponseWirelessGetAccessPointConfigurationRadioDTOsInternalKey struct {
	Type     string   `json:"type,omitempty"`     //
	ID       *float64 `json:"id,omitempty"`       //
	LongType string   `json:"longType,omitempty"` //
	URL      string   `json:"url,omitempty"`      //
}
type ResponseWirelessGetAccessPointConfigurationInternalKey struct {
	Type     string   `json:"type,omitempty"`     //
	ID       *float64 `json:"id,omitempty"`       //
	LongType string   `json:"longType,omitempty"` //
	URL      string   `json:"url,omitempty"`      //
}
type ResponseWirelessApProvision2 struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessDeleteDynamicInterface struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessCreateUpdateDynamicInterface struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL
	Message            string `json:"message,omitempty"`            // Message
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

func (r *ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails) UnmarshalJSON(data []byte) error {
	type Alias ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails
	aux := &struct {
		EnableFabric interface{} `json:"enableFabric"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	switch v := aux.EnableFabric.(type) {
	case bool:
		r.EnableFabric = &v
	case string:
		if v == "true" {
			r.EnableFabric = new(bool)
			*r.EnableFabric = true
		} else if v == "false" {
			r.EnableFabric = new(bool)
			*r.EnableFabric = false
		} else {
			r.EnableFabric = nil
		}
	case nil:
		r.EnableFabric = nil
	default:
		r.EnableFabric = nil
	}

	return nil
}

type ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To VLAN ID
}
type ResponseWirelessProvisionUpdate struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessProvision struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessPSKOverride struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessRetrieveRfProfiles struct {
	Name                 string                                                  `json:"name,omitempty"`                 // RF Profile Name
	DefaultRfProfile     *bool                                                   `json:"defaultRfProfile,omitempty"`     // is Default Rf Profile
	EnableRadioTypeA     *bool                                                   `json:"enableRadioTypeA,omitempty"`     // Enable Radio Type A
	EnableRadioTypeB     *bool                                                   `json:"enableRadioTypeB,omitempty"`     // Enable Radio Type B
	ChannelWidth         string                                                  `json:"channelWidth,omitempty"`         // Channel Width
	EnableCustom         *bool                                                   `json:"enableCustom,omitempty"`         // Enable Custom
	EnableBrownField     *bool                                                   `json:"enableBrownField,omitempty"`     // Enable Brown Field
	RadioTypeAProperties *ResponseWirelessRetrieveRfProfilesRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //
	RadioTypeBProperties *ResponseWirelessRetrieveRfProfilesRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //
	RadioTypeCProperties *ResponseWirelessRetrieveRfProfilesRadioTypeCProperties `json:"radioTypeCProperties,omitempty"` //
	EnableRadioTypeC     *bool                                                   `json:"enableRadioTypeC,omitempty"`     // Enable Radio Type C (6GHz)
}
type ResponseWirelessRetrieveRfProfilesRadioTypeAProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile (Default : CUSTOM)
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates (Default : "6,9,12,18,24,36,48,54")
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates (Default: "6,12,24")
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1 ( (Default: -70)
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold  (Default: "AUTO")
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Rx Sop Threshold  (Default: -10)
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level  (Default: 30)
}
type ResponseWirelessRetrieveRfProfilesRadioTypeBProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile (Default : CUSTOM)
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels (Default : "9,11,12,18,24,36,48,54")
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates  (Default: "9,11,12,18,24,36,48,54")
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "12")
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1  (Default: -70)
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold (Default: "AUTO")
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level  (Default: -10)
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level  (Default: 30)
}
type ResponseWirelessRetrieveRfProfilesRadioTypeCProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile (Default : CUSTOM)
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates  (Default: "6,9,12,18,24,36,48,54")
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "6,12,24")
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold  (Default: "AUTO")
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level  (Default: -10)
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level  (Default: 30)
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1  (Default: -70)
}
type ResponseWirelessCreateOrUpdateRfProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessDeleteRfProfiles struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessFactoryResetAccessPoints struct {
	Response *ResponseWirelessFactoryResetAccessPointsResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseWirelessFactoryResetAccessPointsResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWirelessGetAccessPointsFactoryResetStatus struct {
	Response *[]ResponseWirelessGetAccessPointsFactoryResetStatusResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetAccessPointsFactoryResetStatusResponse struct {
	WlcIP              string                                                                         `json:"wlcIP,omitempty"`              // Wireless Controller IP address
	WlcName            string                                                                         `json:"wlcName,omitempty"`            // Wireless Controller name
	ApResponseInfoList *[]ResponseWirelessGetAccessPointsFactoryResetStatusResponseApResponseInfoList `json:"apResponseInfoList,omitempty"` //
}
type ResponseWirelessGetAccessPointsFactoryResetStatusResponseApResponseInfoList struct {
	ApName               string `json:"apName,omitempty"`               // Access Point name
	ApFactoryResetStatus string `json:"apFactoryResetStatus,omitempty"` // AP factory reset status, "Success" or "Failure" or "In Progress"
	FailureReason        string `json:"failureReason,omitempty"`        // Reason for failure if the factory reset status is "Failure"
	RadioMacAddress      string `json:"radioMacAddress,omitempty"`      // AP Radio Mac Address
	EthernetMacAddress   string `json:"ethernetMacAddress,omitempty"`   // AP Ethernet Mac Address
}
type ResponseWirelessApProvision struct {
	Response *ResponseWirelessApProvisionResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  // Version
}
type ResponseWirelessApProvisionResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetAllMobilityGroups struct {
	Response *[]ResponseWirelessGetAllMobilityGroupsResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Response version.
}
type ResponseWirelessGetAllMobilityGroupsResponse struct {
	MobilityGroupName  string                                                       `json:"mobilityGroupName,omitempty"`  // Self device Group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
	MacAddress         string                                                       `json:"macAddress,omitempty"`         // Device mobility MAC Address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
	ManagementIP       string                                                       `json:"managementIp,omitempty"`       // Self device wireless Management IP.
	NetworkDeviceID    string                                                       `json:"networkDeviceId,omitempty"`    // Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
	DtlsHighCipher     *bool                                                        `json:"dtlsHighCipher,omitempty"`     // DTLS High Cipher.
	DataLinkEncryption *bool                                                        `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.
	MobilityPeers      *[]ResponseWirelessGetAllMobilityGroupsResponseMobilityPeers `json:"mobilityPeers,omitempty"`      //
}
type ResponseWirelessGetAllMobilityGroupsResponseMobilityPeers struct {
	MobilityGroupName   string `json:"mobilityGroupName,omitempty"`   // Peer device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.
	PeerNetworkDeviceID string `json:"peerNetworkDeviceId,omitempty"` // Peer device Id. The possible values are UNKNOWN or valid UUID of Network device ID. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device ID.
	MemberMacAddress    string `json:"memberMacAddress,omitempty"`    // Peer device mobility MAC Address.  Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
	DeviceSeries        string `json:"deviceSeries,omitempty"`        // Peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.
	DataLinkEncryption  *bool  `json:"dataLinkEncryption,omitempty"`  // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers.
	HashKey             string `json:"hashKey,omitempty"`             // SSC hash string must be 40 characters.
	Status              string `json:"status,omitempty"`              // Possible values are - Control and Data Path Down, Data Path Down, Control Path Down, UP.
	PeerIP              string `json:"peerIp,omitempty"`              // This indicates public IP address.
	PrivateIPAddress    string `json:"privateIpAddress,omitempty"`    // This indicates private/management IP address.
}
type ResponseWirelessGetMobilityGroupsCount struct {
	Response *ResponseWirelessGetMobilityGroupsCountResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Response version.
}
type ResponseWirelessGetMobilityGroupsCountResponse struct {
	Count *int `json:"count,omitempty"` // Total number of mobility groups available.
}
type ResponseWirelessMobilityProvision struct {
	Response *ResponseWirelessMobilityProvisionResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // Version
}
type ResponseWirelessMobilityProvisionResponse struct {
	TaskID string `json:"taskId,omitempty"` // Asynchronous Task Id
	URL    string `json:"url,omitempty"`    // Asynchronous Task URL for further tracking
}
type ResponseWirelessMobilityReset struct {
	Response *ResponseWirelessMobilityResetResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  // Version
}
type ResponseWirelessMobilityResetResponse struct {
	TaskID string `json:"taskId,omitempty"` // Asynchronous Task Id
	URL    string `json:"url,omitempty"`    // Asynchronous Task URL for further tracking
}
type ResponseWirelessAssignManagedApLocationsForWLC struct {
	Response *ResponseWirelessAssignManagedApLocationsForWLCResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseWirelessAssignManagedApLocationsForWLCResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessWirelessControllerProvision struct {
	Response *ResponseWirelessWirelessControllerProvisionResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseWirelessWirelessControllerProvisionResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponse `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponse struct {
	ManagedApLocations *[]ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations struct {
	SiteID            string `json:"siteId,omitempty"`            // The site id of the managed ap location.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetManagedApLocationsCountForSpecificWirelessController struct {
	Response *ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerResponse `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerResponse struct {
	PrimaryManagedApLocationsCount   *int `json:"primaryManagedApLocationsCount,omitempty"`   // The count of the Primary managed ap locations.
	SecondaryManagedApLocationsCount *int `json:"secondaryManagedApLocationsCount,omitempty"` // The count of the Secondary managed ap locations.
	AnchorManagedApLocationsCount    *int `json:"anchorManagedApLocationsCount,omitempty"`    // The count of the Anchor managed ap  locations.
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponse struct {
	ManagedApLocations *[]ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations struct {
	SiteID            string `json:"siteId,omitempty"`            // The site id of the managed ap location.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponse `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponse struct {
	ManagedApLocations *[]ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations struct {
	SiteID            string `json:"siteId,omitempty"`            // The site id of the managed ap location.
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetSSIDDetailsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerResponse struct {
	SSIDName        string `json:"ssidName,omitempty"`        // Name of the SSID.
	WLANID          *int   `json:"wlanId,omitempty"`          // WLAN ID.
	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name.
	L2Security      string `json:"l2Security,omitempty"`      // This represents the identifier for the Layer 2 authentication type. The authentication types supported include wpa2_enterprise, wpa2_personal, open, wpa3_enterprise, wpa3_personal, wpa2_wpa3_personal, wpa2_wpa3_enterprise, and open-secured.
	L3Security      string `json:"l3Security,omitempty"`      // This represents the identifier for the Layer 3 authentication type. The authentication types supported are 'open' and 'webauth'.
	RadioPolicy     string `json:"radioPolicy,omitempty"`     // This represents the identifier for the radio policy. The policies supported include 2.4GHz, 5GHz, and 6GHz.
	AdminStatus     *bool  `json:"adminStatus,omitempty"`     // Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed         *bool  `json:"managed,omitempty"`         // If the value is 'true,' the SSID is configured through design; if 'false,' it indicates out-of-band configuration on the Wireless LAN Controller.
}
type ResponseWirelessGetSSIDCountForSpecificWirelessController struct {
	Response *ResponseWirelessGetSSIDCountForSpecificWirelessControllerResponse `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetSSIDCountForSpecificWirelessControllerResponse struct {
	Count *int `json:"count,omitempty"` // The count of the SSIDs.
}
type ResponseWirelessGetWirelessProfiles struct {
	Response *[]ResponseWirelessGetWirelessProfilesResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetWirelessProfilesResponse struct {
	WirelessProfileName string                                                    `json:"wirelessProfileName,omitempty"` // Wireless Profile Name
	SSIDDetails         *[]ResponseWirelessGetWirelessProfilesResponseSSIDDetails `json:"ssidDetails,omitempty"`         //
	ID                  string                                                    `json:"id,omitempty"`                  // Id
}
type ResponseWirelessGetWirelessProfilesResponseSSIDDetails struct {
	SSIDName          string                                                             `json:"ssidName,omitempty"`          // SSID Name
	FlexConnect       *ResponseWirelessGetWirelessProfilesResponseSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`       //
	EnableFabric      *bool                                                              `json:"enableFabric,omitempty"`      // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	WLANProfileName   string                                                             `json:"wlanProfileName,omitempty"`   // WLAN Profile Name
	InterfaceName     string                                                             `json:"interfaceName,omitempty"`     // Interface Name
	PolicyProfileName string                                                             `json:"policyProfileName,omitempty"` // Policy Profile Name
	Dot11BeProfileID  string                                                             `json:"dot11beProfileId,omitempty"`  // 802.11be Profile ID
}
type ResponseWirelessGetWirelessProfilesResponseSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local to VLAN ID
}
type ResponseWirelessCreateWirelessProfile2 struct {
	Response *ResponseWirelessCreateWirelessProfile2Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseWirelessCreateWirelessProfile2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetWirelessProfilesCount struct {
	Response *ResponseWirelessGetWirelessProfilesCountResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGetWirelessProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessUpdateWirelessProfile2 struct {
	Response *ResponseWirelessUpdateWirelessProfile2Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseWirelessUpdateWirelessProfile2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetWirelessProfileByID struct {
	Response *ResponseWirelessGetWirelessProfileByIDResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetWirelessProfileByIDResponse struct {
	WirelessProfileName string                                                       `json:"wirelessProfileName,omitempty"` // Wireless Profile Name
	SSIDDetails         *[]ResponseWirelessGetWirelessProfileByIDResponseSSIDDetails `json:"ssidDetails,omitempty"`         //
	ID                  string                                                       `json:"id,omitempty"`                  // Id
}
type ResponseWirelessGetWirelessProfileByIDResponseSSIDDetails struct {
	SSIDName          string                                                                `json:"ssidName,omitempty"`          // SSID Name
	FlexConnect       *ResponseWirelessGetWirelessProfileByIDResponseSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`       //
	EnableFabric      *bool                                                                 `json:"enableFabric,omitempty"`      // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	WLANProfileName   string                                                                `json:"wlanProfileName,omitempty"`   // WLAN Profile Name
	InterfaceName     string                                                                `json:"interfaceName,omitempty"`     // Interface Name
	PolicyProfileName string                                                                `json:"policyProfileName,omitempty"` // Policy Profile Name
	Dot11BeProfileID  string                                                                `json:"dot11beProfileId,omitempty"`  // 802.11be Profile ID
}
type ResponseWirelessGetWirelessProfileByIDResponseSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local to VLAN ID
}
type ResponseWirelessDeleteWirelessProfile2 struct {
	Response *ResponseWirelessDeleteWirelessProfile2Response `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  // Version
}
type ResponseWirelessDeleteWirelessProfile2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetAll80211BeProfiles struct {
	Response *[]ResponseWirelessGetAll80211BeProfilesResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGetAll80211BeProfilesResponse struct {
	ID             string `json:"id,omitempty"`             // 802.11be Profile ID
	ProfileName    string `json:"profileName,omitempty"`    // 802.11be Profile Name
	OfdmaDownLink  *bool  `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool  `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoDownLink *bool  `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
	MuMimoUpLink   *bool  `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	OfdmaMultiRu   *bool  `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU
	Default        *bool  `json:"default,omitempty"`        // 802.11be Profile is marked default or custom (Read only field)
}
type ResponseWirelessCreateA80211BeProfile struct {
	Response *ResponseWirelessCreateA80211BeProfileResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseWirelessCreateA80211BeProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGet80211BeProfilesCount struct {
	Response *ResponseWirelessGet80211BeProfilesCountResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGet80211BeProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteA80211BeProfile struct {
	Response *ResponseWirelessDeleteA80211BeProfileResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Version
}
type ResponseWirelessDeleteA80211BeProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessUpdate80211BeProfile struct {
	Response *ResponseWirelessUpdate80211BeProfileResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version
}
type ResponseWirelessUpdate80211BeProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGet80211BeProfileByID struct {
	Response *ResponseWirelessGet80211BeProfileByIDResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGet80211BeProfileByIDResponse struct {
	ID             string `json:"id,omitempty"`             // 802.11be Profile ID
	ProfileName    string `json:"profileName,omitempty"`    // 802.11be Profile Name
	OfdmaDownLink  *bool  `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool  `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoDownLink *bool  `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
	MuMimoUpLink   *bool  `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	OfdmaMultiRu   *bool  `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU
	Default        *bool  `json:"default,omitempty"`        // Is 802.11be Profile marked as default in System . (Read only field)
}
type ResponseWirelessGetInterfaces struct {
	Response *[]ResponseWirelessGetInterfacesResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetInterfacesResponse struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name
	VLANID        *int   `json:"vlanId,omitempty"`        // VLAN ID
	ID            string `json:"id,omitempty"`            // Interface ID
}
type ResponseWirelessCreateInterface struct {
	Response *ResponseWirelessCreateInterfaceResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessCreateInterfaceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetInterfacesCount struct {
	Response *ResponseWirelessGetInterfacesCountResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGetInterfacesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetInterfaceByID struct {
	Response *ResponseWirelessGetInterfaceByIDResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetInterfaceByIDResponse struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name
	VLANID        *int   `json:"vlanId,omitempty"`        // VLAN ID
	ID            string `json:"id,omitempty"`            // Interface ID
}
type ResponseWirelessDeleteInterface struct {
	Response *ResponseWirelessDeleteInterfaceResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessDeleteInterfaceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessUpdateInterface struct {
	Response *ResponseWirelessUpdateInterfaceResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessUpdateInterfaceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessCreateRfProfile struct {
	Response *ResponseWirelessCreateRfProfileResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessCreateRfProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetRfProfiles struct {
	Response *[]ResponseWirelessGetRfProfilesResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetRfProfilesResponse struct {
	RfProfileName           string                                                        `json:"rfProfileName,omitempty"`           // RF Profile Name
	DefaultRfProfile        *bool                                                         `json:"defaultRfProfile,omitempty"`        // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
	EnableRadioTypeA        *bool                                                         `json:"enableRadioTypeA,omitempty"`        // True if 5 GHz radio band is enabled in the RF Profile, else False
	EnableRadioTypeB        *bool                                                         `json:"enableRadioTypeB,omitempty"`        // True if 2.4 GHz radio band is enabled in the RF Profile, else False
	EnableRadioType6GHz     *bool                                                         `json:"enableRadioType6GHz,omitempty"`     // True if 6 GHz radio band is enabled in the RF Profile, else False
	EnableCustom            *bool                                                         `json:"enableCustom,omitempty"`            // True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)
	RadioTypeAProperties    *ResponseWirelessGetRfProfilesResponseRadioTypeAProperties    `json:"radioTypeAProperties,omitempty"`    //
	RadioTypeBProperties    *ResponseWirelessGetRfProfilesResponseRadioTypeBProperties    `json:"radioTypeBProperties,omitempty"`    //
	RadioType6GHzProperties *ResponseWirelessGetRfProfilesResponseRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
	ID                      string                                                        `json:"id,omitempty"`                      // RF Profile ID
}
type ResponseWirelessGetRfProfilesResponseRadioTypeAProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 5 GHz radio band
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 5 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 5 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 5 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 5 GHz radio band
	ChannelWidth       string `json:"channelWidth,omitempty"`       // Channel Width
	PreamblePuncture   *bool  `json:"preamblePuncture,omitempty"`   // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
}
type ResponseWirelessGetRfProfilesResponseRadioTypeBProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 2.4 GHz radio band
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 2.4 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 2.4 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 2.4 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 2.4 GHz radio band
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzProperties struct {
	ParentProfile              string                                                                            `json:"parentProfile,omitempty"`              // Parent profile of 6 GHz radio band
	RadioChannels              string                                                                            `json:"radioChannels,omitempty"`              // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
	DataRates                  string                                                                            `json:"dataRates,omitempty"`                  // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates         string                                                                            `json:"mandatoryDataRates,omitempty"`         // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1           *int                                                                              `json:"powerThresholdV1,omitempty"`           // Power threshold of 6 GHz radio band
	RxSopThreshold             string                                                                            `json:"rxSopThreshold,omitempty"`             // RX-SOP threshold of 6 GHz radio band
	MinPowerLevel              *int                                                                              `json:"minPowerLevel,omitempty"`              // Minimum power level of 6 GHz radio band
	MaxPowerLevel              *int                                                                              `json:"maxPowerLevel,omitempty"`              // Maximum power level of 6 GHz radio band
	EnableStandardPowerService *bool                                                                             `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False
	MultiBssidProperties       *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"`       //
	PreamblePuncture           *bool                                                                             `json:"preamblePuncture,omitempty"`           // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
	MinDbsWidth                *int                                                                              `json:"minDbsWidth,omitempty"`                // Minimum DBS Width ( Permissible values : 20,40,80,160,320)
	MaxDbsWidth                *int                                                                              `json:"maxDbsWidth,omitempty"`                // Maximum DBS Width (Permissible Values: 20,40,80,160,320)
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters   *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"`   //
	Dot11BeParameters   *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"`   //
	TargetWakeTime      *bool                                                                                              `json:"targetWakeTime,omitempty"`      // Target Wake Time
	TwtBroadcastSupport *bool                                                                                              `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
	OfdmaMultiRu   *bool `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU
}
type ResponseWirelessGetRfProfilesCount struct {
	Response *ResponseWirelessGetRfProfilesCountResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Response Version
}
type ResponseWirelessGetRfProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteRfProfile struct {
	Response *ResponseWirelessDeleteRfProfileResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessDeleteRfProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessGetRfProfileByID struct {
	Response *ResponseWirelessGetRfProfileByIDResponse `json:"response,omitempty"` //
	Version  string                                    `json:"version,omitempty"`  // Version
}
type ResponseWirelessGetRfProfileByIDResponse struct {
	RfProfileName           string                                                           `json:"rfProfileName,omitempty"`           // RF Profile Name
	DefaultRfProfile        *bool                                                            `json:"defaultRfProfile,omitempty"`        // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
	EnableRadioTypeA        *bool                                                            `json:"enableRadioTypeA,omitempty"`        // True if 5 GHz radio band is enabled in the RF Profile, else False
	EnableRadioTypeB        *bool                                                            `json:"enableRadioTypeB,omitempty"`        // True if 2.4 GHz radio band is enabled in the RF Profile, else False
	EnableRadioType6GHz     *bool                                                            `json:"enableRadioType6GHz,omitempty"`     // True if 6 GHz radio band is enabled in the RF Profile, else False
	EnableCustom            *bool                                                            `json:"enableCustom,omitempty"`            // True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)
	RadioTypeAProperties    *ResponseWirelessGetRfProfileByIDResponseRadioTypeAProperties    `json:"radioTypeAProperties,omitempty"`    //
	RadioTypeBProperties    *ResponseWirelessGetRfProfileByIDResponseRadioTypeBProperties    `json:"radioTypeBProperties,omitempty"`    //
	RadioType6GHzProperties *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
	ID                      string                                                           `json:"id,omitempty"`                      // RF Profile ID
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeAProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 5 GHz radio band
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 5 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 5 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 5 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 5 GHz radio band
	ChannelWidth       string `json:"channelWidth,omitempty"`       // Channel Width
	PreamblePuncture   *bool  `json:"preamblePuncture,omitempty"`   // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeBProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 2.4 GHz radio band
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 2.4 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 2.4 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 2.4 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 2.4 GHz radio band
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzProperties struct {
	ParentProfile              string                                                                               `json:"parentProfile,omitempty"`              // Parent profile of 6 GHz radio band
	RadioChannels              string                                                                               `json:"radioChannels,omitempty"`              // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
	DataRates                  string                                                                               `json:"dataRates,omitempty"`                  // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates         string                                                                               `json:"mandatoryDataRates,omitempty"`         // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1           *int                                                                                 `json:"powerThresholdV1,omitempty"`           // Power threshold of 6 GHz radio band
	RxSopThreshold             string                                                                               `json:"rxSopThreshold,omitempty"`             // RX-SOP threshold of 6 GHz radio band
	MinPowerLevel              *int                                                                                 `json:"minPowerLevel,omitempty"`              // Minimum power level of 6 GHz radio band
	MaxPowerLevel              *int                                                                                 `json:"maxPowerLevel,omitempty"`              // Maximum power level of 6 GHz radio band
	EnableStandardPowerService *bool                                                                                `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False
	MultiBssidProperties       *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"`       //
	PreamblePuncture           *bool                                                                                `json:"preamblePuncture,omitempty"`           // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
	MinDbsWidth                *int                                                                                 `json:"minDbsWidth,omitempty"`                // Minimum DBS Width ( Permissible values : 20,40,80,160,320)
	MaxDbsWidth                *int                                                                                 `json:"maxDbsWidth,omitempty"`                // Maximum DBS Width (Permissible Values: 20,40,80,160,320)
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters   *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"`   //
	Dot11BeParameters   *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"`   //
	TargetWakeTime      *bool                                                                                                 `json:"targetWakeTime,omitempty"`      // Target Wake Time
	TwtBroadcastSupport *bool                                                                                                 `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
	OfdmaMultiRu   *bool `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU
}
type ResponseWirelessUpdateRfProfile struct {
	Response *ResponseWirelessUpdateRfProfileResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  // Version
}
type ResponseWirelessUpdateRfProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseWirelessConfigureAccessPointsV2 struct {
	Response *ResponseWirelessConfigureAccessPointsV2Response `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  //
}
type ResponseWirelessConfigureAccessPointsV2Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestWirelessCreateAndProvisionSSID struct {
	ManagedApLocations []string                                          `json:"managedAPLocations,omitempty"` // Managed AP Locations (Enter entire Site(s) hierarchy)
	SSIDDetails        *RequestWirelessCreateAndProvisionSSIDSSIDDetails `json:"ssidDetails,omitempty"`        //
	SSIDType           string                                            `json:"ssidType,omitempty"`           // SSID Type
	EnableFabric       *bool                                             `json:"enableFabric,omitempty"`       // Enable SSID for Fabric
	FlexConnect        *RequestWirelessCreateAndProvisionSSIDFlexConnect `json:"flexConnect,omitempty"`        //
}
type RequestWirelessCreateAndProvisionSSIDSSIDDetails struct {
	Name                     string   `json:"name,omitempty"`                     // SSID Name
	SecurityLevel            string   `json:"securityLevel,omitempty"`            // Security Level(For guest SSID OPEN/WEB_AUTH, For Enterprise SSID ENTERPRISE/PERSONAL/OPEN)
	EnableFastLane           *bool    `json:"enableFastLane,omitempty"`           // Enable Fast Lane
	Passphrase               string   `json:"passphrase,omitempty"`               // Pass Phrase ( Only applicable for SSID with PERSONAL auth type )
	TrafficType              string   `json:"trafficType,omitempty"`              // Traffic Type
	EnableBroadcastSSID      *bool    `json:"enableBroadcastSSID,omitempty"`      // Enable Broadcast SSID
	RadioPolicy              string   `json:"radioPolicy,omitempty"`              // Radio Policy
	EnableMacFiltering       *bool    `json:"enableMACFiltering,omitempty"`       // Enable MAC Filtering
	FastTransition           string   `json:"fastTransition,omitempty"`           // Fast Transition
	WebAuthURL               string   `json:"webAuthURL,omitempty"`               // Web Auth URL
	AuthKeyMgmt              []string `json:"authKeyMgmt,omitempty"`              // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft
	RsnCipherSuiteGcmp256    *bool    `json:"rsnCipherSuiteGcmp256,omitempty"`    // Rsn Cipher Suite Gcmp256
	RsnCipherSuiteGcmp128    *bool    `json:"rsnCipherSuiteGcmp128,omitempty"`    // Rsn Cipher Suite  Gcmp128
	RsnCipherSuiteCcmp256    *bool    `json:"rsnCipherSuiteCcmp256,omitempty"`    // Rsn Cipher Suite Ccmp256
	Ghz6PolicyClientSteering *bool    `json:"ghz6PolicyClientSteering,omitempty"` // 6 Ghz Client Steering
	Ghz24Policy              string   `json:"ghz24Policy,omitempty"`              // 2.4 GHz Policy
}
type RequestWirelessCreateAndProvisionSSIDFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // Enable Flex Connect
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan (range is 1 to 4094)
}
type RequestWirelessRebootAccessPoints struct {
	ApMacAddresses []string `json:"apMacAddresses,omitempty"` // The ethernet MAC address of the access point.
}
type RequestWirelessCreateEnterpriseSSID struct {
	Name                             string                                                 `json:"name,omitempty"`                             // SSID NAME
	SecurityLevel                    string                                                 `json:"securityLevel,omitempty"`                    // Security Level
	Passphrase                       string                                                 `json:"passphrase,omitempty"`                       // Passphrase
	EnableFastLane                   *bool                                                  `json:"enableFastLane,omitempty"`                   // Enable FastLane
	EnableMacFiltering               *bool                                                  `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	TrafficType                      string                                                 `json:"trafficType,omitempty"`                      // Traffic Type Enum (voicedata or data )
	RadioPolicy                      string                                                 `json:"radioPolicy,omitempty"`                      // Radio Policy Enum
	EnableBroadcastSSID              *bool                                                  `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcase SSID
	FastTransition                   string                                                 `json:"fastTransition,omitempty"`                   // Fast Transition
	EnableSessionTimeOut             *bool                                                  `json:"enableSessionTimeOut,omitempty"`             // Enable Session Timeout
	SessionTimeOut                   *int                                                   `json:"sessionTimeOut,omitempty"`                   // Session Time Out
	EnableClientExclusion            *bool                                                  `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *int                                                   `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool                                                  `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *int                                                   `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout
	EnableDirectedMulticastService   *bool                                                  `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed Multicast Service
	EnableNeighborList               *bool                                                  `json:"enableNeighborList,omitempty"`               // Enable Neighbor List
	MfpClientProtection              string                                                 `json:"mfpClientProtection,omitempty"`              // Management Frame Protection Client
	NasOptions                       []string                                               `json:"nasOptions,omitempty"`                       // Nas Options
	ProfileName                      string                                                 `json:"profileName,omitempty"`                      // Profile Name
	PolicyProfileName                string                                                 `json:"policyProfileName,omitempty"`                // Policy Profile Name
	AAAOverride                      *bool                                                  `json:"aaaOverride,omitempty"`                      // Aaa Override
	CoverageHoleDetectionEnable      *bool                                                  `json:"coverageHoleDetectionEnable,omitempty"`      // Coverage Hole Detection Enable
	ProtectedManagementFrame         string                                                 `json:"protectedManagementFrame,omitempty"`         // (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
	MultipSKSettings                 *[]RequestWirelessCreateEnterpriseSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"`                 //
	ClientRateLimit                  *float64                                               `json:"clientRateLimit,omitempty"`                  // Client Rate Limit (in bits per second)
	AuthKeyMgmt                      []string                                               `json:"authKeyMgmt,omitempty"`                      // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft
	RsnCipherSuiteGcmp256            *bool                                                  `json:"rsnCipherSuiteGcmp256,omitempty"`            // Rsn Cipher Suite Gcmp256
	RsnCipherSuiteCcmp256            *bool                                                  `json:"rsnCipherSuiteCcmp256,omitempty"`            // Rsn Cipher Suite Ccmp256
	RsnCipherSuiteGcmp128            *bool                                                  `json:"rsnCipherSuiteGcmp128,omitempty"`            // Rsn Cipher Suite Gcmp 128
	Ghz6PolicyClientSteering         *bool                                                  `json:"ghz6PolicyClientSteering,omitempty"`         // Ghz6 Policy Client Steering
	Ghz24Policy                      string                                                 `json:"ghz24Policy,omitempty"`                      // Ghz24 Policy
}
type RequestWirelessCreateEnterpriseSSIDMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase
}
type RequestWirelessUpdateEnterpriseSSID struct {
	Name                             string                                                 `json:"name,omitempty"`                             // SSID NAME
	SecurityLevel                    string                                                 `json:"securityLevel,omitempty"`                    // Security Level
	Passphrase                       string                                                 `json:"passphrase,omitempty"`                       // Passphrase
	EnableFastLane                   *bool                                                  `json:"enableFastLane,omitempty"`                   // Enable FastLane
	EnableMacFiltering               *bool                                                  `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	TrafficType                      string                                                 `json:"trafficType,omitempty"`                      // Traffic Type Enum (voicedata or data )
	RadioPolicy                      string                                                 `json:"radioPolicy,omitempty"`                      // Radio Policy Enum
	EnableBroadcastSSID              *bool                                                  `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcase SSID
	FastTransition                   string                                                 `json:"fastTransition,omitempty"`                   // Fast Transition
	EnableSessionTimeOut             *bool                                                  `json:"enableSessionTimeOut,omitempty"`             // Enable Session Timeout
	SessionTimeOut                   *int                                                   `json:"sessionTimeOut,omitempty"`                   // Session Time Out
	EnableClientExclusion            *bool                                                  `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *int                                                   `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool                                                  `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *int                                                   `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout
	EnableDirectedMulticastService   *bool                                                  `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed Multicast Service
	EnableNeighborList               *bool                                                  `json:"enableNeighborList,omitempty"`               // Enable Neighbor List
	MfpClientProtection              string                                                 `json:"mfpClientProtection,omitempty"`              // Management Frame Protection Client
	NasOptions                       []string                                               `json:"nasOptions,omitempty"`                       // Nas Options
	ProfileName                      string                                                 `json:"profileName,omitempty"`                      // Profile Name
	PolicyProfileName                string                                                 `json:"policyProfileName,omitempty"`                // Policy Profile Name
	AAAOverride                      *bool                                                  `json:"aaaOverride,omitempty"`                      // Aaa Override
	CoverageHoleDetectionEnable      *bool                                                  `json:"coverageHoleDetectionEnable,omitempty"`      // Coverage Hole Detection Enable
	ProtectedManagementFrame         string                                                 `json:"protectedManagementFrame,omitempty"`         // (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
	MultipSKSettings                 *[]RequestWirelessUpdateEnterpriseSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"`                 //
	ClientRateLimit                  *float64                                               `json:"clientRateLimit,omitempty"`                  // Client Rate Limit (in bits per second)
	AuthKeyMgmt                      []string                                               `json:"authKeyMgmt,omitempty"`                      // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft
	RsnCipherSuiteGcmp256            *bool                                                  `json:"rsnCipherSuiteGcmp256,omitempty"`            // Rsn Cipher Suite Gcmp256
	RsnCipherSuiteCcmp256            *bool                                                  `json:"rsnCipherSuiteCcmp256,omitempty"`            // Rsn Cipher Suite Ccmp256
	RsnCipherSuiteGcmp128            *bool                                                  `json:"rsnCipherSuiteGcmp128,omitempty"`            // Rsn Cipher Suite Gcmp 128
	Ghz6PolicyClientSteering         *bool                                                  `json:"ghz6PolicyClientSteering,omitempty"`         // Ghz6 Policy Client Steering
	Ghz24Policy                      string                                                 `json:"ghz24Policy,omitempty"`                      // Ghz24 Policy
}
type RequestWirelessUpdateEnterpriseSSIDMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase
}
type RequestWirelessCreateSSID struct {
	SSID                                         string                                       `json:"ssid,omitempty"`                                         // Name of the SSID
	AuthType                                     string                                       `json:"authType,omitempty"`                                     // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
	Passphrase                                   string                                       `json:"passphrase,omitempty"`                                   // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
	IsFastLaneEnabled                            *bool                                        `json:"isFastLaneEnabled,omitempty"`                            // True if FastLane is enabled, else False
	IsMacFilteringEnabled                        *bool                                        `json:"isMacFilteringEnabled,omitempty"`                        // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
	SSIDRadioType                                string                                       `json:"ssidRadioType,omitempty"`                                // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
	IsBroadcastSSID                              *bool                                        `json:"isBroadcastSSID,omitempty"`                              // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
	FastTransition                               string                                       `json:"fastTransition,omitempty"`                               // Fast Transition
	SessionTimeOutEnable                         *bool                                        `json:"sessionTimeOutEnable,omitempty"`                         // Turn on the feature that imposes a time limit on user sessions
	SessionTimeOut                               *int                                         `json:"sessionTimeOut,omitempty"`                               // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity
	ClientExclusionEnable                        *bool                                        `json:"clientExclusionEnable,omitempty"`                        // Activate the feature that allows for the exclusion of clients
	ClientExclusionTimeout                       *int                                         `json:"clientExclusionTimeout,omitempty"`                       // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
	BasicServiceSetMaxIDleEnable                 *bool                                        `json:"basicServiceSetMaxIdleEnable,omitempty"`                 // Activate the maximum idle feature for the Basic Service Set
	BasicServiceSetClientIDleTimeout             *int                                         `json:"basicServiceSetClientIdleTimeout,omitempty"`             // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out
	DirectedMulticastServiceEnable               *bool                                        `json:"directedMulticastServiceEnable,omitempty"`               // The Directed Multicast Service feature becomes operational when it is set to true
	NeighborListEnable                           *bool                                        `json:"neighborListEnable,omitempty"`                           // The Neighbor List feature is enabled when it is set to true
	ManagementFrameProtectionClientprotection    string                                       `json:"managementFrameProtectionClientprotection,omitempty"`    // Management Frame Protection Client
	NasOptions                                   []string                                     `json:"nasOptions,omitempty"`                                   // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
	ProfileName                                  string                                       `json:"profileName,omitempty"`                                  // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName
	AAAOverride                                  *bool                                        `json:"aaaOverride,omitempty"`                                  // Activate the AAA Override feature when set to true
	CoverageHoleDetectionEnable                  *bool                                        `json:"coverageHoleDetectionEnable,omitempty"`                  // Activate Coverage Hole Detection feature when set to true
	ProtectedManagementFrame                     string                                       `json:"protectedManagementFrame,omitempty"`                     // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
	MultipSKSettings                             *[]RequestWirelessCreateSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"`                             //
	ClientRateLimit                              *int                                         `json:"clientRateLimit,omitempty"`                              // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
	RsnCipherSuiteGcmp256                        *bool                                        `json:"rsnCipherSuiteGcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
	RsnCipherSuiteCcmp256                        *bool                                        `json:"rsnCipherSuiteCcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
	RsnCipherSuiteGcmp128                        *bool                                        `json:"rsnCipherSuiteGcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated
	RsnCipherSuiteCcmp128                        *bool                                        `json:"rsnCipherSuiteCcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
	Ghz6PolicyClientSteering                     *bool                                        `json:"ghz6PolicyClientSteering,omitempty"`                     // True if 6 GHz Policy Client Steering is enabled, else False
	IsAuthKey8021X                               *bool                                        `json:"isAuthKey8021x,omitempty"`                               // When set to true, the 802.1X authentication key is in use
	IsAuthKey8021XPlusFT                         *bool                                        `json:"isAuthKey8021xPlusFT,omitempty"`                         // When set to true, the 802.1X-Plus-FT authentication key is in use
	IsAuthKey8021XSHA256                         *bool                                        `json:"isAuthKey8021x_SHA256,omitempty"`                        // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
	IsAuthKeySae                                 *bool                                        `json:"isAuthKeySae,omitempty"`                                 // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
	IsAuthKeySaePlusFT                           *bool                                        `json:"isAuthKeySaePlusFT,omitempty"`                           // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
	IsAuthKeyPSK                                 *bool                                        `json:"isAuthKeyPSK,omitempty"`                                 // When set to true, the Pre-shared Key (PSK) authentication feature is enabled
	IsAuthKeyPSKPlusFT                           *bool                                        `json:"isAuthKeyPSKPlusFT,omitempty"`                           // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
	IsAuthKeyOWE                                 *bool                                        `json:"isAuthKeyOWE,omitempty"`                                 // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
	IsAuthKeyEasyPSK                             *bool                                        `json:"isAuthKeyEasyPSK,omitempty"`                             // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
	IsAuthKeyPSKSHA256                           *bool                                        `json:"isAuthKeyPSKSHA256,omitempty"`                           // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
	OpenSSID                                     string                                       `json:"openSsid,omitempty"`                                     // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
	WLANBandSelectEnable                         *bool                                        `json:"wlanBandSelectEnable,omitempty"`                         // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
	IsEnabled                                    *bool                                        `json:"isEnabled,omitempty"`                                    // Set SSID's admin status as 'Enabled' when set to true
	AuthServers                                  []string                                     `json:"authServers,omitempty"`                                  // List of Authentication/Authorization server IpAddresses
	AcctServers                                  []string                                     `json:"acctServers,omitempty"`                                  // List of Accounting server IpAddresses
	EgressQos                                    string                                       `json:"egressQos,omitempty"`                                    // Egress QOS
	IngressQos                                   string                                       `json:"ingressQos,omitempty"`                                   // Ingress QOS
	WLANType                                     string                                       `json:"wlanType,omitempty"`                                     // Wlan Type
	L3AuthType                                   string                                       `json:"l3AuthType,omitempty"`                                   // L3 Authentication Type
	AuthServer                                   string                                       `json:"authServer,omitempty"`                                   // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth
	ExternalAuthIPAddress                        string                                       `json:"externalAuthIpAddress,omitempty"`                        // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
	WebPassthrough                               *bool                                        `json:"webPassthrough,omitempty"`                               // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
	SleepingClientEnable                         *bool                                        `json:"sleepingClientEnable,omitempty"`                         // When set to true, this will activate the timeout settings that apply to clients in sleep mode
	SleepingClientTimeout                        *int                                         `json:"sleepingClientTimeout,omitempty"`                        // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
	ACLName                                      string                                       `json:"aclName,omitempty"`                                      // Pre-Auth Access Control List (ACL) Name
	IsPosturingEnabled                           *bool                                        `json:"isPosturingEnabled,omitempty"`                           // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
	IsAuthKeySuiteB1X                            *bool                                        `json:"isAuthKeySuiteB1x,omitempty"`                            // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
	IsAuthKeySuiteB1921X                         *bool                                        `json:"isAuthKeySuiteB1921x,omitempty"`                         // When set to true, the SuiteB192-1x authentication key feature is enabled.
	IsAuthKeySaeExt                              *bool                                        `json:"isAuthKeySaeExt,omitempty"`                              // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
	IsAuthKeySaeExtPlusFT                        *bool                                        `json:"isAuthKeySaeExtPlusFT,omitempty"`                        // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
	IsApBeaconProtectionEnabled                  *bool                                        `json:"isApBeaconProtectionEnabled,omitempty"`                  // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
	Ghz24Policy                                  string                                       `json:"ghz24Policy,omitempty"`                                  // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
	CckmTsfTolerance                             *int                                         `json:"cckmTsfTolerance,omitempty"`                             // Cckm TImestamp Tolerance(in milliseconds)
	IsCckmEnabled                                *bool                                        `json:"isCckmEnabled,omitempty"`                                // True if CCKM is enabled, else False
	IsHex                                        *bool                                        `json:"isHex,omitempty"`                                        // True if passphrase is in Hex format, else False.
	IsRandomMacFilterEnabled                     *bool                                        `json:"isRandomMacFilterEnabled,omitempty"`                     // Deny clients using randomized MAC addresses when set to true
	FastTransitionOverTheDistributedSystemEnable *bool                                        `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type RequestWirelessCreateSSIDMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessUpdateSSID struct {
	SSID                                         string                                       `json:"ssid,omitempty"`                                         // Name of the SSID
	AuthType                                     string                                       `json:"authType,omitempty"`                                     // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)
	Passphrase                                   string                                       `json:"passphrase,omitempty"`                                   // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
	IsFastLaneEnabled                            *bool                                        `json:"isFastLaneEnabled,omitempty"`                            // True if FastLane is enabled, else False
	IsMacFilteringEnabled                        *bool                                        `json:"isMacFilteringEnabled,omitempty"`                        // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device
	SSIDRadioType                                string                                       `json:"ssidRadioType,omitempty"`                                // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))
	IsBroadcastSSID                              *bool                                        `json:"isBroadcastSSID,omitempty"`                              // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks
	FastTransition                               string                                       `json:"fastTransition,omitempty"`                               // Fast Transition
	SessionTimeOutEnable                         *bool                                        `json:"sessionTimeOutEnable,omitempty"`                         // Turn on the feature that imposes a time limit on user sessions
	SessionTimeOut                               *int                                         `json:"sessionTimeOut,omitempty"`                               // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity
	ClientExclusionEnable                        *bool                                        `json:"clientExclusionEnable,omitempty"`                        // Activate the feature that allows for the exclusion of clients
	ClientExclusionTimeout                       *int                                         `json:"clientExclusionTimeout,omitempty"`                       // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts
	BasicServiceSetMaxIDleEnable                 *bool                                        `json:"basicServiceSetMaxIdleEnable,omitempty"`                 // Activate the maximum idle feature for the Basic Service Set
	BasicServiceSetClientIDleTimeout             *int                                         `json:"basicServiceSetClientIdleTimeout,omitempty"`             // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out
	DirectedMulticastServiceEnable               *bool                                        `json:"directedMulticastServiceEnable,omitempty"`               // The Directed Multicast Service feature becomes operational when it is set to true
	NeighborListEnable                           *bool                                        `json:"neighborListEnable,omitempty"`                           // The Neighbor List feature is enabled when it is set to true
	ManagementFrameProtectionClientprotection    string                                       `json:"managementFrameProtectionClientprotection,omitempty"`    // Management Frame Protection Client
	NasOptions                                   []string                                     `json:"nasOptions,omitempty"`                                   // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.
	ProfileName                                  string                                       `json:"profileName,omitempty"`                                  // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName
	AAAOverride                                  *bool                                        `json:"aaaOverride,omitempty"`                                  // Activate the AAA Override feature when set to true
	CoverageHoleDetectionEnable                  *bool                                        `json:"coverageHoleDetectionEnable,omitempty"`                  // Activate Coverage Hole Detection feature when set to true
	ProtectedManagementFrame                     string                                       `json:"protectedManagementFrame,omitempty"`                     // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)
	MultipSKSettings                             *[]RequestWirelessUpdateSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"`                             //
	ClientRateLimit                              *int                                         `json:"clientRateLimit,omitempty"`                              // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve
	RsnCipherSuiteGcmp256                        *bool                                        `json:"rsnCipherSuiteGcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated
	RsnCipherSuiteCcmp256                        *bool                                        `json:"rsnCipherSuiteCcmp256,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated
	RsnCipherSuiteGcmp128                        *bool                                        `json:"rsnCipherSuiteGcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated
	RsnCipherSuiteCcmp128                        *bool                                        `json:"rsnCipherSuiteCcmp128,omitempty"`                        // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated
	Ghz6PolicyClientSteering                     *bool                                        `json:"ghz6PolicyClientSteering,omitempty"`                     // True if 6 GHz Policy Client Steering is enabled, else False
	IsAuthKey8021X                               *bool                                        `json:"isAuthKey8021x,omitempty"`                               // When set to true, the 802.1X authentication key is in use
	IsAuthKey8021XPlusFT                         *bool                                        `json:"isAuthKey8021xPlusFT,omitempty"`                         // When set to true, the 802.1X-Plus-FT authentication key is in use
	IsAuthKey8021XSHA256                         *bool                                        `json:"isAuthKey8021x_SHA256,omitempty"`                        // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on
	IsAuthKeySae                                 *bool                                        `json:"isAuthKeySae,omitempty"`                                 // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated
	IsAuthKeySaePlusFT                           *bool                                        `json:"isAuthKeySaePlusFT,omitempty"`                           // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)
	IsAuthKeyPSK                                 *bool                                        `json:"isAuthKeyPSK,omitempty"`                                 // When set to true, the Pre-shared Key (PSK) authentication feature is enabled
	IsAuthKeyPSKPlusFT                           *bool                                        `json:"isAuthKeyPSKPlusFT,omitempty"`                           // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated
	IsAuthKeyOWE                                 *bool                                        `json:"isAuthKeyOWE,omitempty"`                                 // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on
	IsAuthKeyEasyPSK                             *bool                                        `json:"isAuthKeyEasyPSK,omitempty"`                             // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated
	IsAuthKeyPSKSHA256                           *bool                                        `json:"isAuthKeyPSKSHA256,omitempty"`                           // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true
	OpenSSID                                     string                                       `json:"openSsid,omitempty"`                                     // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID
	WLANBandSelectEnable                         *bool                                        `json:"wlanBandSelectEnable,omitempty"`                         // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band
	IsEnabled                                    *bool                                        `json:"isEnabled,omitempty"`                                    // Set SSID's admin status as 'Enabled' when set to true
	AuthServers                                  []string                                     `json:"authServers,omitempty"`                                  // List of Authentication/Authorization server IpAddresses
	AcctServers                                  []string                                     `json:"acctServers,omitempty"`                                  // List of Accounting server IpAddresses
	EgressQos                                    string                                       `json:"egressQos,omitempty"`                                    // Egress QOS
	IngressQos                                   string                                       `json:"ingressQos,omitempty"`                                   // Ingress QOS
	WLANType                                     string                                       `json:"wlanType,omitempty"`                                     // Wlan Type
	L3AuthType                                   string                                       `json:"l3AuthType,omitempty"`                                   // L3 Authentication Type
	AuthServer                                   string                                       `json:"authServer,omitempty"`                                   // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth
	ExternalAuthIPAddress                        string                                       `json:"externalAuthIpAddress,omitempty"`                        // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)
	WebPassthrough                               *bool                                        `json:"webPassthrough,omitempty"`                               // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements
	SleepingClientEnable                         *bool                                        `json:"sleepingClientEnable,omitempty"`                         // When set to true, this will activate the timeout settings that apply to clients in sleep mode
	SleepingClientTimeout                        *int                                         `json:"sleepingClientTimeout,omitempty"`                        // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network
	ACLName                                      string                                       `json:"aclName,omitempty"`                                      // Pre-Auth Access Control List (ACL) Name
	IsPosturingEnabled                           *bool                                        `json:"isPosturingEnabled,omitempty"`                           // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.
	IsAuthKeySuiteB1X                            *bool                                        `json:"isAuthKeySuiteB1x,omitempty"`                            // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.
	IsAuthKeySuiteB1921X                         *bool                                        `json:"isAuthKeySuiteB1921x,omitempty"`                         // When set to true, the SuiteB192-1x authentication key feature is enabled.
	IsAuthKeySaeExt                              *bool                                        `json:"isAuthKeySaeExt,omitempty"`                              // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.
	IsAuthKeySaeExtPlusFT                        *bool                                        `json:"isAuthKeySaeExtPlusFT,omitempty"`                        // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.
	IsApBeaconProtectionEnabled                  *bool                                        `json:"isApBeaconProtectionEnabled,omitempty"`                  // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.
	Ghz24Policy                                  string                                       `json:"ghz24Policy,omitempty"`                                  // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType
	CckmTsfTolerance                             *int                                         `json:"cckmTsfTolerance,omitempty"`                             // Cckm TImestamp Tolerance(in milliseconds)
	IsCckmEnabled                                *bool                                        `json:"isCckmEnabled,omitempty"`                                // True if CCKM is enabled, else False
	IsHex                                        *bool                                        `json:"isHex,omitempty"`                                        // True if passphrase is in Hex format, else False.
	IsRandomMacFilterEnabled                     *bool                                        `json:"isRandomMacFilterEnabled,omitempty"`                     // Deny clients using randomized MAC addresses when set to true
	FastTransitionOverTheDistributedSystemEnable *bool                                        `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type RequestWirelessUpdateSSIDMultipSKSettings struct {
	Priority       *int   `json:"priority,omitempty"`       // Priority
	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type
	Passphrase     string `json:"passphrase,omitempty"`     // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessConfigureAccessPointsV1 struct {
	ApList                      *[]RequestWirelessConfigureAccessPointsV1ApList              `json:"apList,omitempty"`                      //
	ConfigureAdminStatus        *bool                                                        `json:"configureAdminStatus,omitempty"`        // To change the access point's admin status, set this parameter's value to "true".
	AdminStatus                 *bool                                                        `json:"adminStatus,omitempty"`                 // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureApMode             *bool                                                        `json:"configureApMode,omitempty"`             // To change the access point's mode, set this parameter's value to "true".
	ApMode                      *int                                                         `json:"apMode,omitempty"`                      // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".
	ConfigureFailoverPriority   *bool                                                        `json:"configureFailoverPriority,omitempty"`   // To change the access point's failover priority, set this parameter's value to "true".
	FailoverPriority            *int                                                         `json:"failoverPriority,omitempty"`            // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".
	ConfigureLedStatus          *bool                                                        `json:"configureLedStatus,omitempty"`          // To change the access point's LED status, set this parameter's value to "true".
	LedStatus                   *bool                                                        `json:"ledStatus,omitempty"`                   // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.
	ConfigureLedBrightnessLevel *bool                                                        `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".
	LedBrightnessLevel          *int                                                         `json:"ledBrightnessLevel,omitempty"`          // Configure the access point's LED brightness level by setting a value between 1 and 8.
	ConfigureLocation           *bool                                                        `json:"configureLocation,omitempty"`           // To change the access point's location, set this parameter's value to "true".
	Location                    string                                                       `json:"location,omitempty"`                    // Configure the access point's location.
	ConfigureHAController       *bool                                                        `json:"configureHAController,omitempty"`       // To change the access point's HA controller, set this parameter's value to "true".
	PrimaryControllerName       string                                                       `json:"primaryControllerName,omitempty"`       // Configure the hostname for an access point's primary controller.
	PrimaryIPAddress            *RequestWirelessConfigureAccessPointsV1PrimaryIPAddress      `json:"primaryIpAddress,omitempty"`            //
	SecondaryControllerName     string                                                       `json:"secondaryControllerName,omitempty"`     // Configure the hostname for an access point's secondary controller.
	SecondaryIPAddress          *RequestWirelessConfigureAccessPointsV1SecondaryIPAddress    `json:"secondaryIpAddress,omitempty"`          //
	TertiaryControllerName      string                                                       `json:"tertiaryControllerName,omitempty"`      // Configure the hostname for an access point's tertiary controller.
	TertiaryIPAddress           *RequestWirelessConfigureAccessPointsV1TertiaryIPAddress     `json:"tertiaryIpAddress,omitempty"`           //
	RadioConfigurations         *[]RequestWirelessConfigureAccessPointsV1RadioConfigurations `json:"radioConfigurations,omitempty"`         //
	IsAssignedSiteAsLocation    *bool                                                        `json:"isAssignedSiteAsLocation,omitempty"`    // If AP is assigned to a site, then to assign AP location as the site name, set this parameter's value to "true".
}
type RequestWirelessConfigureAccessPointsV1ApList struct {
	ApName     string `json:"apName,omitempty"`     // The current host name of the access point.
	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.
	ApNameNew  string `json:"apNameNew,omitempty"`  // The modified hostname of the access point.
}
type RequestWirelessConfigureAccessPointsV1PrimaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's primary controller.
}
type RequestWirelessConfigureAccessPointsV1SecondaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's secondary controller.
}
type RequestWirelessConfigureAccessPointsV1TertiaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's tertiary controller.
}
type RequestWirelessConfigureAccessPointsV1RadioConfigurations struct {
	ConfigureRadioRoleAssignment *bool    `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".
	RadioRoleAssignment          string   `json:"radioRoleAssignment,omitempty"`          // Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string
	RadioBand                    string   `json:"radioBand,omitempty"`                    // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string
	ConfigureAdminStatus         *bool    `json:"configureAdminStatus,omitempty"`         // To change the admin status on the specified radio for an access point, set this parameter's value to "true".
	AdminStatus                  *bool    `json:"adminStatus,omitempty"`                  // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureAntennaPatternName  *bool    `json:"configureAntennaPatternName,omitempty"`  // To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".
	AntennaPatternName           string   `json:"antennaPatternName,omitempty"`           // Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.
	AntennaGain                  *int     `json:"antennaGain,omitempty"`                  // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".
	ConfigureAntennaCable        *bool    `json:"configureAntennaCable,omitempty"`        // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".
	AntennaCableName             string   `json:"antennaCableName,omitempty"`             // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".
	CableLoss                    *float64 `json:"cableLoss,omitempty"`                    // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).
	ConfigureChannel             *bool    `json:"configureChannel,omitempty"`             // To change the channel on the specified radio for an access point, set this parameter's value to "true".
	ChannelAssignmentMode        *int     `json:"channelAssignmentMode,omitempty"`        // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
	ChannelNumber                *int     `json:"channelNumber,omitempty"`                // Configure the channel number on the specified radio for an access point.
	ConfigureChannelWidth        *bool    `json:"configureChannelWidth,omitempty"`        // To change the channel width on the specified radio for an access point, set this parameter's value to "true".
	ChannelWidth                 *int     `json:"channelWidth,omitempty"`                 // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".
	ConfigurePower               *bool    `json:"configurePower,omitempty"`               // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".
	PowerAssignmentMode          *int     `json:"powerAssignmentMode,omitempty"`          // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
	Powerlevel                   *int     `json:"powerlevel,omitempty"`                   // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.
	ConfigureCleanAirSI          *bool    `json:"configureCleanAirSI,omitempty"`          // To enable or disable either CleanAir or Spectrum Intelligence on the specified radio for an access point, set this parameter's value to "true".
	CleanAirSI                   *int     `json:"cleanAirSI,omitempty"`                   // Configure CleanAir or Spectrum Intelligence on the specified radio for an access point. Set this parameter's value to "0" to disable the feature or "1" to enable it.
	RadioType                    *int     `json:"radioType,omitempty"`                    // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
}
type RequestWirelessApProvision2 []RequestItemWirelessApProvision2 // Array of RequestWirelessAPProvision2
type RequestItemWirelessApProvision2 struct {
	RfProfile           string   `json:"rfProfile,omitempty"`           // Radio frequency profile name
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
	Name              string                                                                    `json:"name,omitempty"`              // Ssid Name
	EnableFabric      *bool                                                                     `json:"enableFabric,omitempty"`      // true if ssid is fabric else false
	FlexConnect       *RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`       //
	InterfaceName     string                                                                    `json:"interfaceName,omitempty"`     // Interface Name
	WLANProfileName   string                                                                    `json:"wlanProfileName,omitempty"`   // WLAN Profile Name
	PolicyProfileName string                                                                    `json:"policyProfileName,omitempty"` // Policy Profile Name
}
type RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan Id
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
	Name              string                                                                    `json:"name,omitempty"`              // Ssid Name
	EnableFabric      *bool                                                                     `json:"enableFabric,omitempty"`      // true if ssid is fabric else false
	FlexConnect       *RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`       //
	InterfaceName     string                                                                    `json:"interfaceName,omitempty"`     // Interface Name
	WLANProfileName   string                                                                    `json:"wlanProfileName,omitempty"`   // WLAN Profile Name
	PolicyProfileName string                                                                    `json:"policyProfileName,omitempty"` // Policy Profile Name
}
type RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan Id
}
type RequestWirelessProvisionUpdate []RequestItemWirelessProvisionUpdate // Array of RequestWirelessProvisionUpdate
type RequestItemWirelessProvisionUpdate struct {
	DeviceName         string                                                 `json:"deviceName,omitempty"`         // Controller Name
	ManagedApLocations []string                                               `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)
	DynamicInterfaces  *[]RequestItemWirelessProvisionUpdateDynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
}
type RequestItemWirelessProvisionUpdateDynamicInterfaces struct {
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     // Interface IP Address. Required for AireOS.
	InterfaceNetmaskInCIDR *int   `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR. Required for AireOS.
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       // Interface Gateway. Required for AireOS.
	LagOrPortNumber        *int   `json:"lagOrPortNumber,omitempty"`        // Lag Or Port Number. Required for AireOS.
	VLANID                 *int   `json:"vlanId,omitempty"`                 // VLAN ID. Required for AireOS and EWLC.
	InterfaceName          string `json:"interfaceName,omitempty"`          // Interface Name. Required for AireOS and EWLC.
}
type RequestWirelessProvision []RequestItemWirelessProvision // Array of RequestWirelessProvision
type RequestItemWirelessProvision struct {
	DeviceName         string                                           `json:"deviceName,omitempty"`         // Controller Name
	Site               string                                           `json:"site,omitempty"`               // Full Site Hierarchy where device has to be assigned
	ManagedApLocations []string                                         `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)
	DynamicInterfaces  *[]RequestItemWirelessProvisionDynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
}
type RequestItemWirelessProvisionDynamicInterfaces struct {
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     // Interface IP Address. Required for AireOS.
	InterfaceNetmaskInCIDR *int   `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR. Required for AireOS.
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       // Interface Gateway.  Required for AireOS.
	LagOrPortNumber        *int   `json:"lagOrPortNumber,omitempty"`        // Lag Or Port Number.  Required for AireOS.
	VLANID                 *int   `json:"vlanId,omitempty"`                 // VLAN ID. Required for both AireOS and EWLC.
	InterfaceName          string `json:"interfaceName,omitempty"`          // Interface Name. Required for both AireOS and EWLC.
}
type RequestWirelessPSKOverride []RequestItemWirelessPSKOverride // Array of RequestWirelessPSKOverride
type RequestItemWirelessPSKOverride struct {
	SSID            string `json:"ssid,omitempty"`            // enterprise ssid name(already created/present)
	Site            string `json:"site,omitempty"`            // site name hierarchy (ex: Global/aaa/zzz/...)
	PassPhrase      string `json:"passPhrase,omitempty"`      // Pass phrase (create/update)
	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name
}
type RequestWirelessCreateOrUpdateRfProfile struct {
	Name                 string                                                      `json:"name,omitempty"`                 // RF Profile Name
	DefaultRfProfile     *bool                                                       `json:"defaultRfProfile,omitempty"`     // is Default Rf Profile
	EnableRadioTypeA     *bool                                                       `json:"enableRadioTypeA,omitempty"`     // Enable Radio Type A
	EnableRadioTypeB     *bool                                                       `json:"enableRadioTypeB,omitempty"`     // Enable Radio Type B
	ChannelWidth         string                                                      `json:"channelWidth,omitempty"`         // Channel Width
	EnableCustom         *bool                                                       `json:"enableCustom,omitempty"`         // Enable Custom
	EnableBrownField     *bool                                                       `json:"enableBrownField,omitempty"`     // Enable Brown Field
	RadioTypeAProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //
	RadioTypeBProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //
	RadioTypeCProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties `json:"radioTypeCProperties,omitempty"` //
	EnableRadioTypeC     *bool                                                       `json:"enableRadioTypeC,omitempty"`     // Enable Radio Type C (6GHz)
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile (Default : CUSTOM)
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates (Default : "6,9,12,18,24,36,48,54")
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates (Default: "6,12,24")
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1 ( (Default: -70)
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold  (Default: "AUTO")
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Rx Sop Threshold  (Default: -10)
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level  (Default: 30)
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile (Default : CUSTOM)
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels (Default : "9,11,12,18,24,36,48,54")
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates  (Default: "9,11,12,18,24,36,48,54")
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "12")
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1  (Default: -70)
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold (Default: "AUTO")
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level  (Default: -10)
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level  (Default: 30)
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile (Default : CUSTOM)
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates  (Default: "6,9,12,18,24,36,48,54")
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "6,12,24")
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold  (Default: "AUTO")
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level  (Default: -10)
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level  (Default: 30)
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1  (Default: -70)
}
type RequestWirelessFactoryResetAccessPoints struct {
	KeepStaticIPConfig *bool    `json:"keepStaticIPConfig,omitempty"` // Set the value of keepStaticIPConfig to false, to clear all configurations from Access Points and set the value of keepStaticIPConfig to true, to clear all configurations from Access Points without clearing static IP configuration.
	ApMacAddresses     []string `json:"apMacAddresses,omitempty"`     // List of Access Point's Ethernet MAC addresses, set maximum 100 ethernet MAC addresses per request.
}
type RequestWirelessApProvision struct {
	NetworkDevices *[]RequestWirelessApProvisionNetworkDevices `json:"networkDevices,omitempty"` //
	RfProfileName  string                                      `json:"rfProfileName,omitempty"`  // RF Profile Name. RF Profile is not allowed for custom AP Zones.
	ApZoneName     string                                      `json:"apZoneName,omitempty"`     // AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.
	SiteID         string                                      `json:"siteId,omitempty"`         // Site ID
}
type RequestWirelessApProvisionNetworkDevices struct {
	DeviceID string `json:"deviceId,omitempty"` // Network device ID of access points
	MeshRole string `json:"meshRole,omitempty"` // Mesh Role (Applicable only when AP is in Bridge Mode)
}
type RequestWirelessMobilityProvision struct {
	MobilityGroupName  string                                           `json:"mobilityGroupName,omitempty"`  // Self device Group Name. Must be alphanumeric without {!,<,space,?/'} <br/> and maximum of 31 characters.
	MacAddress         string                                           `json:"macAddress,omitempty"`         // Device mobility MAC Address. Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
	ManagementIP       string                                           `json:"managementIp,omitempty"`       // Self device wireless Management IP.
	NetworkDeviceID    string                                           `json:"networkDeviceId,omitempty"`    // Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
	DtlsHighCipher     *bool                                            `json:"dtlsHighCipher,omitempty"`     // DTLS High Cipher.
	DataLinkEncryption *bool                                            `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.
	MobilityPeers      *[]RequestWirelessMobilityProvisionMobilityPeers `json:"mobilityPeers,omitempty"`      //
}
type RequestWirelessMobilityProvisionMobilityPeers struct {
	PeerIP              string `json:"peerIp,omitempty"`              // This indicates public ip address.
	PrivateIPAddress    string `json:"privateIpAddress,omitempty"`    // This indicates private/management ip address.
	PeerDeviceName      string `json:"peerDeviceName,omitempty"`      // Peer device Host Name.
	PeerNetworkDeviceID string `json:"peerNetworkDeviceId,omitempty"` // The possible values are: UNKNOWN or valid UUID of Network device Id. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device id.
	MobilityGroupName   string `json:"mobilityGroupName,omitempty"`   // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} <br/> and maximum of 31 characters.
	MemberMacAddress    string `json:"memberMacAddress,omitempty"`    // Peer device mobility MAC Address.  Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11
	DeviceSeries        string `json:"deviceSeries,omitempty"`        // Indicates peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.
	HashKey             string `json:"hashKey,omitempty"`             // SSC hash string must be 40 characters.
}
type RequestWirelessMobilityReset struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device Id of Cisco wireless controller.Obtain the network device ID value by using the API call GET - /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
}
type RequestWirelessAssignManagedApLocationsForWLC struct {
	PrimaryManagedApLocationsSiteIDs   []string `json:"primaryManagedAPLocationsSiteIds,omitempty"`   // Site IDs of Primary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site
	SecondaryManagedApLocationsSiteIDs []string `json:"secondaryManagedAPLocationsSiteIds,omitempty"` // Site IDs of Secondary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site
}
type RequestWirelessWirelessControllerProvision struct {
	Interfaces       *[]RequestWirelessWirelessControllerProvisionInterfaces     `json:"interfaces,omitempty"`       //
	SkipApProvision  *bool                                                       `json:"skipApProvision,omitempty"`  // True if Skip AP Provision is enabled, else False
	RollingApUpgrade *RequestWirelessWirelessControllerProvisionRollingApUpgrade `json:"rollingApUpgrade,omitempty"` //
}
type RequestWirelessWirelessControllerProvisionInterfaces struct {
	InterfaceName          string `json:"interfaceName,omitempty"`          // Interface Name
	VLANID                 *int   `json:"vlanId,omitempty"`                 // VLAN ID range is 1 - 4094
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     // Interface IP Address
	InterfaceNetmaskInCIDR *int   `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR, range is 1-30
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       // Interface Gateway
	LagOrPortNumber        *int   `json:"lagOrPortNumber,omitempty"`        // Lag Or Port Number
}
type RequestWirelessWirelessControllerProvisionRollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // True if Rolling AP Upgrade is enabled, else False
	ApRebootPercentage     *int  `json:"apRebootPercentage,omitempty"`     // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type RequestWirelessCreateWirelessProfile2 struct {
	WirelessProfileName string                                              `json:"wirelessProfileName,omitempty"` // Wireless Network Profile Name
	SSIDDetails         *[]RequestWirelessCreateWirelessProfile2SSIDDetails `json:"ssidDetails,omitempty"`         //
}
type RequestWirelessCreateWirelessProfile2SSIDDetails struct {
	SSIDName         string                                                       `json:"ssidName,omitempty"`         // SSID Name
	FlexConnect      *RequestWirelessCreateWirelessProfile2SSIDDetailsFlexConnect `json:"flexConnect,omitempty"`      //
	EnableFabric     *bool                                                        `json:"enableFabric,omitempty"`     // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	WLANProfileName  string                                                       `json:"wlanProfileName,omitempty"`  // WLAN Profile Name
	InterfaceName    string                                                       `json:"interfaceName,omitempty"`    // Interface Name. Default Value: management
	Dot11BeProfileID string                                                       `json:"dot11beProfileId,omitempty"` // 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured
}
type RequestWirelessCreateWirelessProfile2SSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local to VLAN ID
}
type RequestWirelessUpdateWirelessProfile2 struct {
	WirelessProfileName string                                              `json:"wirelessProfileName,omitempty"` // Wireless Network Profile Name
	SSIDDetails         *[]RequestWirelessUpdateWirelessProfile2SSIDDetails `json:"ssidDetails,omitempty"`         //
}
type RequestWirelessUpdateWirelessProfile2SSIDDetails struct {
	SSIDName         string                                                       `json:"ssidName,omitempty"`         // SSID Name
	FlexConnect      *RequestWirelessUpdateWirelessProfile2SSIDDetailsFlexConnect `json:"flexConnect,omitempty"`      //
	EnableFabric     *bool                                                        `json:"enableFabric,omitempty"`     // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	WLANProfileName  string                                                       `json:"wlanProfileName,omitempty"`  // WLAN Profile Name
	InterfaceName    string                                                       `json:"interfaceName,omitempty"`    // Interface Name. Default Value: management
	Dot11BeProfileID string                                                       `json:"dot11beProfileId,omitempty"` // 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured
}
type RequestWirelessUpdateWirelessProfile2SSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local to VLAN ID
}
type RequestWirelessCreateA80211BeProfile struct {
	ProfileName    string `json:"profileName,omitempty"`    // 802.11be Profile Name
	OfdmaDownLink  *bool  `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink (Default: true)
	OfdmaUpLink    *bool  `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink (Default: true)
	MuMimoDownLink *bool  `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink (Default: false)
	MuMimoUpLink   *bool  `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink (Default: false)
	OfdmaMultiRu   *bool  `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU (Default: false)
}
type RequestWirelessUpdate80211BeProfile struct {
	ProfileName    string `json:"profileName,omitempty"`    // 802.11be Profile Name
	OfdmaDownLink  *bool  `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink (Default: true)
	OfdmaUpLink    *bool  `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink (Default: true)
	MuMimoDownLink *bool  `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink (Default: false)
	MuMimoUpLink   *bool  `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink (Default: false)
	OfdmaMultiRu   *bool  `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU (Default: false)
}
type RequestWirelessCreateInterface struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name
	VLANID        *int   `json:"vlanId,omitempty"`        // VLAN ID range is 1-4094
}
type RequestWirelessUpdateInterface struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name
	VLANID        *int   `json:"vlanId,omitempty"`        // VLAN ID range is 1-4094
}
type RequestWirelessCreateRfProfile struct {
	RfProfileName           string                                                 `json:"rfProfileName,omitempty"`           // RF Profile Name
	DefaultRfProfile        *bool                                                  `json:"defaultRfProfile,omitempty"`        // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
	EnableRadioTypeA        *bool                                                  `json:"enableRadioTypeA,omitempty"`        // True if 5 GHz radio band is enabled in the RF Profile, else False
	EnableRadioTypeB        *bool                                                  `json:"enableRadioTypeB,omitempty"`        // True if 2.4 GHz radio band is enabled in the RF Profile, else False
	EnableRadioType6GHz     *bool                                                  `json:"enableRadioType6GHz,omitempty"`     // True if 6 GHz radio band is enabled in the RF Profile, else False
	RadioTypeAProperties    *RequestWirelessCreateRfProfileRadioTypeAProperties    `json:"radioTypeAProperties,omitempty"`    //
	RadioTypeBProperties    *RequestWirelessCreateRfProfileRadioTypeBProperties    `json:"radioTypeBProperties,omitempty"`    //
	RadioType6GHzProperties *RequestWirelessCreateRfProfileRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileRadioTypeAProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 5 GHz radio band
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 5 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 5 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 5 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 5 GHz radio band
	ChannelWidth       string `json:"channelWidth,omitempty"`       // Channel Width
	PreamblePuncture   *bool  `json:"preamblePuncture,omitempty"`   // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
}
type RequestWirelessCreateRfProfileRadioTypeBProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 2.4 GHz radio band
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 2.4 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 2.4 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 2.4 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 2.4 GHz radio band
}
type RequestWirelessCreateRfProfileRadioType6GHzProperties struct {
	ParentProfile              string                                                                     `json:"parentProfile,omitempty"`              // Parent profile of 6 GHz radio band
	RadioChannels              string                                                                     `json:"radioChannels,omitempty"`              // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
	DataRates                  string                                                                     `json:"dataRates,omitempty"`                  // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates         string                                                                     `json:"mandatoryDataRates,omitempty"`         // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1           *int                                                                       `json:"powerThresholdV1,omitempty"`           // Power threshold of 6 GHz radio band
	RxSopThreshold             string                                                                     `json:"rxSopThreshold,omitempty"`             // RX-SOP threshold of 6 GHz radio band
	MinPowerLevel              *int                                                                       `json:"minPowerLevel,omitempty"`              // Minimum power level of 6 GHz radio band
	MaxPowerLevel              *int                                                                       `json:"maxPowerLevel,omitempty"`              // Maximum power level of 6 GHz radio band
	EnableStandardPowerService *bool                                                                      `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False
	MultiBssidProperties       *RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"`       //
	PreamblePuncture           *bool                                                                      `json:"preamblePuncture,omitempty"`           // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
	MinDbsWidth                *int                                                                       `json:"minDbsWidth,omitempty"`                // Minimum DBS Width (Permissible Values:20,40,80,160,320)
	MaxDbsWidth                *int                                                                       `json:"maxDbsWidth,omitempty"`                // Maximum DBS Width (Permissible Values:20,40,80,160,320)
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters   *RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"`   //
	Dot11BeParameters   *RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"`   //
	TargetWakeTime      *bool                                                                                       `json:"targetWakeTime,omitempty"`      // Target Wake Time
	TwtBroadcastSupport *bool                                                                                       `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
	OfdmaMultiRu   *bool `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU
}
type RequestWirelessUpdateRfProfile struct {
	RfProfileName           string                                                 `json:"rfProfileName,omitempty"`           // RF Profile Name
	DefaultRfProfile        *bool                                                  `json:"defaultRfProfile,omitempty"`        // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
	EnableRadioTypeA        *bool                                                  `json:"enableRadioTypeA,omitempty"`        // True if 5 GHz radio band is enabled in the RF Profile, else False
	EnableRadioTypeB        *bool                                                  `json:"enableRadioTypeB,omitempty"`        // True if 2.4 GHz radio band is enabled in the RF Profile, else False
	EnableRadioType6GHz     *bool                                                  `json:"enableRadioType6GHz,omitempty"`     // True if 6 GHz radio band is enabled in the RF Profile, else False
	RadioTypeAProperties    *RequestWirelessUpdateRfProfileRadioTypeAProperties    `json:"radioTypeAProperties,omitempty"`    //
	RadioTypeBProperties    *RequestWirelessUpdateRfProfileRadioTypeBProperties    `json:"radioTypeBProperties,omitempty"`    //
	RadioType6GHzProperties *RequestWirelessUpdateRfProfileRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileRadioTypeAProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 5 GHz radio band. In case of brownfield learnt RF Profile if the parent profile is GLOBAL, any change in RF Profile configurations will not be provisioned to device. Existing parent profile with values of HIGH, TYPICAL, LOW or CUSTOM cannot be modified to GLOBAL
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 5 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 5 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 5 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 5 GHz radio band
	ChannelWidth       string `json:"channelWidth,omitempty"`       // Channel Width
	PreamblePuncture   *bool  `json:"preamblePuncture,omitempty"`   // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
}
type RequestWirelessUpdateRfProfileRadioTypeBProperties struct {
	ParentProfile      string `json:"parentProfile,omitempty"`      // Parent profile of 2.4 GHz radio band. In case of brownfield learnt RF Profile if the parent profile is GLOBAL, any change in RF Profile configurations will not be provisioned to device. Existing parent profile with values of HIGH, TYPICAL, LOW or CUSTOM cannot be modified to GLOBAL
	RadioChannels      string `json:"radioChannels,omitempty"`      // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
	DataRates          string `json:"dataRates,omitempty"`          // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
	PowerThresholdV1   *int   `json:"powerThresholdV1,omitempty"`   // Power threshold of 2.4 GHz radio band
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     // RX-SOP threshold of 2.4 GHz radio band
	MinPowerLevel      *int   `json:"minPowerLevel,omitempty"`      // Minimum power level of 2.4 GHz radio band
	MaxPowerLevel      *int   `json:"maxPowerLevel,omitempty"`      // Maximum power level of 2.4 GHz radio band
}
type RequestWirelessUpdateRfProfileRadioType6GHzProperties struct {
	ParentProfile              string                                                                     `json:"parentProfile,omitempty"`              // Parent profile of 6 GHz radio band
	RadioChannels              string                                                                     `json:"radioChannels,omitempty"`              // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
	DataRates                  string                                                                     `json:"dataRates,omitempty"`                  // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	MandatoryDataRates         string                                                                     `json:"mandatoryDataRates,omitempty"`         // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
	PowerThresholdV1           *int                                                                       `json:"powerThresholdV1,omitempty"`           // Power threshold of 6 GHz radio band
	RxSopThreshold             string                                                                     `json:"rxSopThreshold,omitempty"`             // RX-SOP threshold of 6 GHz radio band
	MinPowerLevel              *int                                                                       `json:"minPowerLevel,omitempty"`              // Minimum power level of 6 GHz radio band
	MaxPowerLevel              *int                                                                       `json:"maxPowerLevel,omitempty"`              // Maximum power level of 6 GHz radio band
	EnableStandardPowerService *bool                                                                      `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False
	MultiBssidProperties       *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"`       //
	PreamblePuncture           *bool                                                                      `json:"preamblePuncture,omitempty"`           // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
	MinDbsWidth                *int                                                                       `json:"minDbsWidth,omitempty"`                // Minimum DBS Width (Permissible Values:20,40,80,160,320)
	MaxDbsWidth                *int                                                                       `json:"maxDbsWidth,omitempty"`                // Maximum DBS Width (Permissible Values:20,40,80,160,320)
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters   *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"`   //
	Dot11BeParameters   *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"`   //
	TargetWakeTime      *bool                                                                                       `json:"targetWakeTime,omitempty"`      // Target Wake Time
	TwtBroadcastSupport *bool                                                                                       `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink  *bool `json:"ofdmaDownLink,omitempty"`  // OFDMA Downlink
	OfdmaUpLink    *bool `json:"ofdmaUpLink,omitempty"`    // OFDMA Uplink
	MuMimoUpLink   *bool `json:"muMimoUpLink,omitempty"`   // MU-MIMO Uplink
	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
	OfdmaMultiRu   *bool `json:"ofdmaMultiRu,omitempty"`   // OFDMA Multi-RU
}
type RequestWirelessConfigureAccessPointsV2 struct {
	ApList                      *[]RequestWirelessConfigureAccessPointsV2ApList              `json:"apList,omitempty"`                      //
	ConfigureAdminStatus        *bool                                                        `json:"configureAdminStatus,omitempty"`        // To change the access point's admin status, set this parameter's value to "true".
	AdminStatus                 *bool                                                        `json:"adminStatus,omitempty"`                 // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureApMode             *bool                                                        `json:"configureApMode,omitempty"`             // To change the access point's mode, set this parameter's value to "true".
	ApMode                      *int                                                         `json:"apMode,omitempty"`                      // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".
	ConfigureFailoverPriority   *bool                                                        `json:"configureFailoverPriority,omitempty"`   // To change the access point's failover priority, set this parameter's value to "true".
	FailoverPriority            *int                                                         `json:"failoverPriority,omitempty"`            // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".
	ConfigureLedStatus          *bool                                                        `json:"configureLedStatus,omitempty"`          // To change the access point's LED status, set this parameter's value to "true".
	LedStatus                   *bool                                                        `json:"ledStatus,omitempty"`                   // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.
	ConfigureLedBrightnessLevel *bool                                                        `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".
	LedBrightnessLevel          *int                                                         `json:"ledBrightnessLevel,omitempty"`          // Configure the access point's LED brightness level by setting a value between 1 and 8.
	ConfigureLocation           *bool                                                        `json:"configureLocation,omitempty"`           // To change the access point's location, set this parameter's value to "true".
	Location                    string                                                       `json:"location,omitempty"`                    // Configure the access point's location.
	ConfigureHAController       *bool                                                        `json:"configureHAController,omitempty"`       // To change the access point's HA controller, set this parameter's value to "true".
	PrimaryControllerName       string                                                       `json:"primaryControllerName,omitempty"`       // Configure the hostname for an access point's primary controller.
	PrimaryIPAddress            *RequestWirelessConfigureAccessPointsV2PrimaryIPAddress      `json:"primaryIpAddress,omitempty"`            //
	SecondaryControllerName     string                                                       `json:"secondaryControllerName,omitempty"`     // Configure the hostname for an access point's secondary controller.
	SecondaryIPAddress          *RequestWirelessConfigureAccessPointsV2SecondaryIPAddress    `json:"secondaryIpAddress,omitempty"`          //
	TertiaryControllerName      string                                                       `json:"tertiaryControllerName,omitempty"`      // Configure the hostname for an access point's tertiary controller.
	TertiaryIPAddress           *RequestWirelessConfigureAccessPointsV2TertiaryIPAddress     `json:"tertiaryIpAddress,omitempty"`           //
	RadioConfigurations         *[]RequestWirelessConfigureAccessPointsV2RadioConfigurations `json:"radioConfigurations,omitempty"`         //
	ConfigureCleanAirSI24Ghz    *bool                                                        `json:"configureCleanAirSI24Ghz,omitempty"`    // To change the clean air status for radios that are in 2.4 Ghz band, set this parameter's value to "true".
	CleanAirSI24                *bool                                                        `json:"cleanAirSI24,omitempty"`                // Configure clean air status for radios that are in 2.4 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureCleanAirSI5Ghz     *bool                                                        `json:"configureCleanAirSI5Ghz,omitempty"`     // To change the clean air status for radios that are in 5 Ghz band, set this parameter's value to "true".
	CleanAirSI5                 *bool                                                        `json:"cleanAirSI5,omitempty"`                 // Configure clean air status for radios that are in 5 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureCleanAirSI6Ghz     *bool                                                        `json:"configureCleanAirSI6Ghz,omitempty"`     // To change the clean air status for radios that are in 6 Ghz band, set this parameter's value to "true".
	CleanAirSI6                 *bool                                                        `json:"cleanAirSI6,omitempty"`                 // Configure clean air status for radios that are in 6 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.
	IsAssignedSiteAsLocation    *bool                                                        `json:"isAssignedSiteAsLocation,omitempty"`    // To configure the access point's location as the site assigned to the access point, set this parameter's value to "true".
}
type RequestWirelessConfigureAccessPointsV2ApList struct {
	ApName     string `json:"apName,omitempty"`     // The current host name of the access point.
	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.
	ApNameNew  string `json:"apNameNew,omitempty"`  // The modified hostname of the access point.
}
type RequestWirelessConfigureAccessPointsV2PrimaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's primary controller.
}
type RequestWirelessConfigureAccessPointsV2SecondaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's secondary controller.
}
type RequestWirelessConfigureAccessPointsV2TertiaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's tertiary controller.
}
type RequestWirelessConfigureAccessPointsV2RadioConfigurations struct {
	ConfigureRadioRoleAssignment *bool    `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".
	RadioRoleAssignment          string   `json:"radioRoleAssignment,omitempty"`          // Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string
	RadioBand                    string   `json:"radioBand,omitempty"`                    // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string
	ConfigureAdminStatus         *bool    `json:"configureAdminStatus,omitempty"`         // To change the admin status on the specified radio for an access point, set this parameter's value to "true".
	AdminStatus                  *bool    `json:"adminStatus,omitempty"`                  // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureAntennaPatternName  *bool    `json:"configureAntennaPatternName,omitempty"`  // To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".
	AntennaPatternName           string   `json:"antennaPatternName,omitempty"`           // Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.
	AntennaGain                  *int     `json:"antennaGain,omitempty"`                  // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".
	ConfigureAntennaCable        *bool    `json:"configureAntennaCable,omitempty"`        // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".
	AntennaCableName             string   `json:"antennaCableName,omitempty"`             // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".
	CableLoss                    *float64 `json:"cableLoss,omitempty"`                    // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).
	ConfigureChannel             *bool    `json:"configureChannel,omitempty"`             // To change the channel on the specified radio for an access point, set this parameter's value to "true".
	ChannelAssignmentMode        *int     `json:"channelAssignmentMode,omitempty"`        // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
	ChannelNumber                *int     `json:"channelNumber,omitempty"`                // Configure the channel number on the specified radio for an access point.
	ConfigureChannelWidth        *bool    `json:"configureChannelWidth,omitempty"`        // To change the channel width on the specified radio for an access point, set this parameter's value to "true".
	ChannelWidth                 *int     `json:"channelWidth,omitempty"`                 // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".
	ConfigurePower               *bool    `json:"configurePower,omitempty"`               // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".
	PowerAssignmentMode          *int     `json:"powerAssignmentMode,omitempty"`          // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
	Powerlevel                   *int     `json:"powerlevel,omitempty"`                   // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.
	RadioType                    *int     `json:"radioType,omitempty"`                    // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
}

//SensorTestResults Sensor Test Results - 87ae-7b21-4f0b-a838
/* Intent API to get SENSOR test result summary


@param SensorTestResultsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!sensor-test-results
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SensorTestResults(SensorTestResultsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation SensorTestResults")
	}

	result := response.Result().(*ResponseWirelessSensorTestResults)
	return result, response, err

}

//GetAccessPointRebootTaskResult Get Access Point Reboot task result - c4b5-e9ce-460a-a8a3
/* Users can query the access point reboot status using this intent API


@param GetAccessPointRebootTaskResultQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-reboot-task-result
*/
func (s *WirelessService) GetAccessPointRebootTaskResult(GetAccessPointRebootTaskResultQueryParams *GetAccessPointRebootTaskResultQueryParams) (*ResponseWirelessGetAccessPointRebootTaskResult, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-reboot/apreboot/status"

	queryString, _ := query.Values(GetAccessPointRebootTaskResultQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointRebootTaskResult{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointRebootTaskResult(GetAccessPointRebootTaskResultQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointRebootTaskResult")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointRebootTaskResult)
	return result, response, err

}

//GetEnterpriseSSID Get Enterprise SSID - cca5-19ba-45eb-b423
/* Get Enterprise SSID


@param GetEnterpriseSSIDQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-enterprise-ssid
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEnterpriseSSID(GetEnterpriseSSIDQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessGetEnterpriseSSID)
	return result, response, err

}

//GetSSIDBySite Get SSID by Site - bb92-f946-4e19-a187
/* This API allows the user to get all SSIDs (Service Set Identifier) at the given site


@param siteID siteId path parameter. Site UUID

@param GetSSIDBySiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-by-site
*/
func (s *WirelessService) GetSSIDBySite(siteID string, GetSSIDBySiteQueryParams *GetSSIDBySiteQueryParams) (*ResponseWirelessGetSSIDBySite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	queryString, _ := query.Values(GetSSIDBySiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSSIDBySite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDBySite(siteID, GetSSIDBySiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidBySite")
	}

	result := response.Result().(*ResponseWirelessGetSSIDBySite)
	return result, response, err

}

//GetSSIDCountBySite Get SSID Count by Site - 52ae-589a-48ab-9116
/* This API allows the user to get count of all SSIDs (Service Set Identifier) present at global site.


@param siteID siteId path parameter. Site UUID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-count-by-site
*/
func (s *WirelessService) GetSSIDCountBySite(siteID string) (*ResponseWirelessGetSSIDCountBySite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/count"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetSSIDCountBySite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDCountBySite(siteID)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidCountBySite")
	}

	result := response.Result().(*ResponseWirelessGetSSIDCountBySite)
	return result, response, err

}

//GetSSIDByID Get SSID by ID - 78a1-2804-47a9-a6a8
/* This API allows the user to get an SSID (Service Set Identifier) by ID at the given site


@param siteID siteId path parameter. Site UUID

@param id id path parameter. SSID ID.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-by-id
*/
func (s *WirelessService) GetSSIDByID(siteID string, id string) (*ResponseWirelessGetSSIDByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetSSIDByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDByID(siteID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidById")
	}

	result := response.Result().(*ResponseWirelessGetSSIDByID)
	return result, response, err

}

//GetAccessPointConfigurationTaskResult Get Access Point Configuration task result - fb90-69dc-4aeb-9afb
/* Users can query the access point configuration result using this intent API


@param taskTypeID task_id path parameter. task id information of ap config


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-configuration-task-result
*/
func (s *WirelessService) GetAccessPointConfigurationTaskResult(taskTypeID string) (*ResponseWirelessGetAccessPointConfigurationTaskResult, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration/details/{task_id}"
	path = strings.Replace(path, "{task_id}", fmt.Sprintf("%v", taskTypeID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAccessPointConfigurationTaskResult{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointConfigurationTaskResult(taskTypeID)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfigurationTaskResult")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfigurationTaskResult)
	return result, response, err

}

//GetAccessPointConfiguration Get Access Point Configuration - a191-f9f2-4cb8-9a55
/* Users can query the access point configuration information per device using the ethernet MAC address


@param GetAccessPointConfigurationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-configuration
*/
func (s *WirelessService) GetAccessPointConfiguration(GetAccessPointConfigurationQueryParams *GetAccessPointConfigurationQueryParams) (*ResponseWirelessGetAccessPointConfiguration, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration/summary"

	queryString, _ := query.Values(GetAccessPointConfigurationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointConfiguration{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointConfiguration(GetAccessPointConfigurationQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfiguration")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfiguration)
	return result, response, err

}

//GetDynamicInterface Get dynamic interface - c5b0-c978-4dfb-90b4
/* Get one or all dynamic interface(s)


@param GetDynamicInterfaceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-dynamic-interface
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDynamicInterface(GetDynamicInterfaceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDynamicInterface")
	}

	result := response.Result().(*ResponseWirelessGetDynamicInterface)
	return result, response, err

}

//GetWirelessProfile Get Wireless Profile - b3a1-c880-4c8b-9b8b
/* Gets either one or all the wireless network profiles if no name is provided for network-profile.


@param GetWirelessProfileQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profile
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfile(GetWirelessProfileQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfile)
	return result, response, err

}

//RetrieveRfProfiles Retrieve RF profiles - 098c-ab91-41c9-a3fe
/* Retrieve all RF profiles


@param RetrieveRFProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-rf-profiles
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveRfProfiles(RetrieveRFProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveRfProfiles")
	}

	result := response.Result().(*ResponseWirelessRetrieveRfProfiles)
	return result, response, err

}

//GetAccessPointsFactoryResetStatus Get Access Point(s) Factory Reset status - 46bf-881b-45b8-a62f
/* This API returns each AP Factory Reset initiation status.


@param GetAccessPointsFactoryResetStatusQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-points-factory-reset-status
*/
func (s *WirelessService) GetAccessPointsFactoryResetStatus(GetAccessPointsFactoryResetStatusQueryParams *GetAccessPointsFactoryResetStatusQueryParams) (*ResponseWirelessGetAccessPointsFactoryResetStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessAccessPoints/factoryResetRequestStatus"

	queryString, _ := query.Values(GetAccessPointsFactoryResetStatusQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointsFactoryResetStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointsFactoryResetStatus(GetAccessPointsFactoryResetStatusQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointsFactoryResetStatus")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointsFactoryResetStatus)
	return result, response, err

}

//GetAllMobilityGroups Get All MobilityGroups - 628f-38bf-4f5a-a48c
/* Retrieve all configured mobility groups if no Network Device Id is provided as a query parameter. If a Network Device Id is given and a mobility group is configured for it, return the configured details; otherwise, return the default values from the device.


@param GetAllMobilityGroupsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-mobility-groups
*/
func (s *WirelessService) GetAllMobilityGroups(GetAllMobilityGroupsQueryParams *GetAllMobilityGroupsQueryParams) (*ResponseWirelessGetAllMobilityGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups"

	queryString, _ := query.Values(GetAllMobilityGroupsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAllMobilityGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllMobilityGroups(GetAllMobilityGroupsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllMobilityGroups")
	}

	result := response.Result().(*ResponseWirelessGetAllMobilityGroups)
	return result, response, err

}

//GetMobilityGroupsCount Get MobilityGroups Count - 29b2-08fb-420a-8970
/* Retrieves count of mobility groups configured



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mobility-groups-count
*/
func (s *WirelessService) GetMobilityGroupsCount() (*ResponseWirelessGetMobilityGroupsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetMobilityGroupsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMobilityGroupsCount()
		}
		return nil, response, fmt.Errorf("error with operation GetMobilityGroupsCount")
	}

	result := response.Result().(*ResponseWirelessGetMobilityGroupsCount)
	return result, response, err

}

//GetAnchorManagedApLocationsForSpecificWirelessController Get Anchor Managed AP Locations for specific Wireless Controller - 8dad-59b4-44b8-8995
/* Retrieves all the details of Anchor Managed AP locations associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetAnchorManagedAPLocationsForSpecificWirelessControllerQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-managed-ap-locations-for-specific-wireless-controller
*/
func (s *WirelessService) GetAnchorManagedApLocationsForSpecificWirelessController(networkDeviceID string, GetAnchorManagedAPLocationsForSpecificWirelessControllerQueryParams *GetAnchorManagedApLocationsForSpecificWirelessControllerQueryParams) (*ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessController, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/anchorManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetAnchorManagedAPLocationsForSpecificWirelessControllerQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessController{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorManagedApLocationsForSpecificWirelessController(networkDeviceID, GetAnchorManagedAPLocationsForSpecificWirelessControllerQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorManagedApLocationsForSpecificWirelessController")
	}

	result := response.Result().(*ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessController)
	return result, response, err

}

//GetManagedApLocationsCountForSpecificWirelessController Get Managed AP Locations Count for specific Wireless Controller - f490-6a9b-4c29-bc6a
/* Retrieves the count of Managed AP locations, including Primary Managed AP Locations, Secondary Managed AP Locations, and Anchor Managed AP Locations, associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-managed-ap-locations-count-for-specific-wireless-controller
*/
func (s *WirelessService) GetManagedApLocationsCountForSpecificWirelessController(networkDeviceID string) (*ResponseWirelessGetManagedApLocationsCountForSpecificWirelessController, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/managedApLocations/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetManagedApLocationsCountForSpecificWirelessController{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetManagedApLocationsCountForSpecificWirelessController(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetManagedApLocationsCountForSpecificWirelessController")
	}

	result := response.Result().(*ResponseWirelessGetManagedApLocationsCountForSpecificWirelessController)
	return result, response, err

}

//GetPrimaryManagedApLocationsForSpecificWirelessController Get Primary Managed AP Locations for specific Wireless Controller - 1dba-89f4-40ab-abda
/* Retrieves all the details of Primary Managed AP locations associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetPrimaryManagedAPLocationsForSpecificWirelessControllerQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-primary-managed-ap-locations-for-specific-wireless-controller
*/
func (s *WirelessService) GetPrimaryManagedApLocationsForSpecificWirelessController(networkDeviceID string, GetPrimaryManagedAPLocationsForSpecificWirelessControllerQueryParams *GetPrimaryManagedApLocationsForSpecificWirelessControllerQueryParams) (*ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessController, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/primaryManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetPrimaryManagedAPLocationsForSpecificWirelessControllerQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessController{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPrimaryManagedApLocationsForSpecificWirelessController(networkDeviceID, GetPrimaryManagedAPLocationsForSpecificWirelessControllerQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPrimaryManagedApLocationsForSpecificWirelessController")
	}

	result := response.Result().(*ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessController)
	return result, response, err

}

//GetSecondaryManagedApLocationsForSpecificWirelessController Get Secondary Managed AP Locations for specific Wireless Controller - b589-7bd6-4f1b-9efb
/* Retrieves all the details of Secondary Managed AP locations associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetSecondaryManagedAPLocationsForSpecificWirelessControllerQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-secondary-managed-ap-locations-for-specific-wireless-controller
*/
func (s *WirelessService) GetSecondaryManagedApLocationsForSpecificWirelessController(networkDeviceID string, GetSecondaryManagedAPLocationsForSpecificWirelessControllerQueryParams *GetSecondaryManagedApLocationsForSpecificWirelessControllerQueryParams) (*ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessController, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/secondaryManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSecondaryManagedAPLocationsForSpecificWirelessControllerQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessController{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecondaryManagedApLocationsForSpecificWirelessController(networkDeviceID, GetSecondaryManagedAPLocationsForSpecificWirelessControllerQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecondaryManagedApLocationsForSpecificWirelessController")
	}

	result := response.Result().(*ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessController)
	return result, response, err

}

//GetSSIDDetailsForSpecificWirelessController Get SSID Details for specific Wireless Controller - 70b6-393d-4899-ad4d
/* Retrieves all details of SSIDs associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetSSIDDetailsForSpecificWirelessControllerQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-details-for-specific-wireless-controller
*/
func (s *WirelessService) GetSSIDDetailsForSpecificWirelessController(networkDeviceID string, GetSSIDDetailsForSpecificWirelessControllerQueryParams *GetSSIDDetailsForSpecificWirelessControllerQueryParams) (*ResponseWirelessGetSSIDDetailsForSpecificWirelessController, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/ssidDetails"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSSIDDetailsForSpecificWirelessControllerQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSSIDDetailsForSpecificWirelessController{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDDetailsForSpecificWirelessController(networkDeviceID, GetSSIDDetailsForSpecificWirelessControllerQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidDetailsForSpecificWirelessController")
	}

	result := response.Result().(*ResponseWirelessGetSSIDDetailsForSpecificWirelessController)
	return result, response, err

}

//GetSSIDCountForSpecificWirelessController Get SSID Count for specific Wireless Controller - 3e98-c91d-42eb-a469
/* Retrieves the count of SSIDs associated with the specific Wireless Controller.


@param networkDeviceID networkDeviceId path parameter. Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

@param GetSSIDCountForSpecificWirelessControllerQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-count-for-specific-wireless-controller
*/
func (s *WirelessService) GetSSIDCountForSpecificWirelessController(networkDeviceID string, GetSSIDCountForSpecificWirelessControllerQueryParams *GetSSIDCountForSpecificWirelessControllerQueryParams) (*ResponseWirelessGetSSIDCountForSpecificWirelessController, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/ssidDetails/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSSIDCountForSpecificWirelessControllerQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetSSIDCountForSpecificWirelessController{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDCountForSpecificWirelessController(networkDeviceID, GetSSIDCountForSpecificWirelessControllerQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidCountForSpecificWirelessController")
	}

	result := response.Result().(*ResponseWirelessGetSSIDCountForSpecificWirelessController)
	return result, response, err

}

//GetWirelessProfiles Get Wireless Profiles - 7988-fac4-447b-8e3d
/* This API allows the user to get all Wireless Network Profiles


@param GetWirelessProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profiles
*/
func (s *WirelessService) GetWirelessProfiles(GetWirelessProfilesQueryParams *GetWirelessProfilesQueryParams) (*ResponseWirelessGetWirelessProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles"

	queryString, _ := query.Values(GetWirelessProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetWirelessProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfiles(GetWirelessProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfiles")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfiles)
	return result, response, err

}

//GetWirelessProfilesCount Get Wireless Profiles Count - 48a7-1883-48fb-93a5
/* This API allows the user to get count of all wireless profiles



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profiles-count
*/
func (s *WirelessService) GetWirelessProfilesCount() (*ResponseWirelessGetWirelessProfilesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetWirelessProfilesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfilesCount()
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfilesCount")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfilesCount)
	return result, response, err

}

//GetWirelessProfileByID Get Wireless Profile by ID - f5b9-fab9-4b79-b0f3
/* This API allows the user to get a Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-wireless-profile-by-id
*/
func (s *WirelessService) GetWirelessProfileByID(id string) (*ResponseWirelessGetWirelessProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetWirelessProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetWirelessProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetWirelessProfileById")
	}

	result := response.Result().(*ResponseWirelessGetWirelessProfileByID)
	return result, response, err

}

//GetAll80211BeProfiles Get all 802.11be Profiles - 1895-aac1-4428-bd0d
/* This API allows the user to get all 802.11be Profile(s) configured under Wireless Settings


@param GetAll80211beProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all80211be-profiles
*/
func (s *WirelessService) GetAll80211BeProfiles(GetAll80211beProfilesQueryParams *GetAll80211BeProfilesQueryParams) (*ResponseWirelessGetAll80211BeProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"

	queryString, _ := query.Values(GetAll80211beProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAll80211BeProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAll80211BeProfiles(GetAll80211beProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAll80211BeProfiles")
	}

	result := response.Result().(*ResponseWirelessGetAll80211BeProfiles)
	return result, response, err

}

//Get80211BeProfilesCount Get 802.11be Profiles Count - a0b7-da85-4faa-95b7
/* This API allows the user to get count of all 802.11be Profile(s)



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get80211be-profiles-count
*/
func (s *WirelessService) Get80211BeProfilesCount() (*ResponseWirelessGet80211BeProfilesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGet80211BeProfilesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Get80211BeProfilesCount()
		}
		return nil, response, fmt.Errorf("error with operation Get80211BeProfilesCount")
	}

	result := response.Result().(*ResponseWirelessGet80211BeProfilesCount)
	return result, response, err

}

//Get80211BeProfileByID Get 802.11be Profile by ID - fa93-88ce-49eb-a5d7
/* This API allows the user to get 802.11be Profile by ID


@param id id path parameter. 802.11be Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get80211be-profile-by-id
*/
func (s *WirelessService) Get80211BeProfileByID(id string) (*ResponseWirelessGet80211BeProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGet80211BeProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Get80211BeProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation Get80211BeProfileById")
	}

	result := response.Result().(*ResponseWirelessGet80211BeProfileByID)
	return result, response, err

}

//GetInterfaces Get Interfaces - 3793-ea73-438a-b243
/* This API allows the user to get all Interfaces


@param GetInterfacesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interfaces
*/
func (s *WirelessService) GetInterfaces(GetInterfacesQueryParams *GetInterfacesQueryParams) (*ResponseWirelessGetInterfaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces"

	queryString, _ := query.Values(GetInterfacesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaces(GetInterfacesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaces")
	}

	result := response.Result().(*ResponseWirelessGetInterfaces)
	return result, response, err

}

//GetInterfacesCount Get Interfaces Count - fd81-f950-424b-b992
/* This API allows the user to get count of all interfaces



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interfaces-count
*/
func (s *WirelessService) GetInterfacesCount() (*ResponseWirelessGetInterfacesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetInterfacesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfacesCount()
		}
		return nil, response, fmt.Errorf("error with operation GetInterfacesCount")
	}

	result := response.Result().(*ResponseWirelessGetInterfacesCount)
	return result, response, err

}

//GetInterfaceByID Get Interface by ID - 3fa4-19ab-482a-ad07
/* This API allows the user to get an interface by ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-interface-by-id
*/
func (s *WirelessService) GetInterfaceByID(id string) (*ResponseWirelessGetInterfaceByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetInterfaceByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetInterfaceByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetInterfaceById")
	}

	result := response.Result().(*ResponseWirelessGetInterfaceByID)
	return result, response, err

}

//GetRfProfiles Get RF Profiles - 15a6-e823-49ca-a9cc
/* This API allows the user to get all RF Profiles


@param GetRFProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rf-profiles
*/
func (s *WirelessService) GetRfProfiles(GetRFProfilesQueryParams *GetRfProfilesQueryParams) (*ResponseWirelessGetRfProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles"

	queryString, _ := query.Values(GetRFProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetRfProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRfProfiles(GetRFProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetRfProfiles")
	}

	result := response.Result().(*ResponseWirelessGetRfProfiles)
	return result, response, err

}

//GetRfProfilesCount Get RF Profiles Count - f996-2b80-477a-9de2
/* This API allows the user to get count of all RF profiles



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rf-profiles-count
*/
func (s *WirelessService) GetRfProfilesCount() (*ResponseWirelessGetRfProfilesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetRfProfilesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRfProfilesCount()
		}
		return nil, response, fmt.Errorf("error with operation GetRfProfilesCount")
	}

	result := response.Result().(*ResponseWirelessGetRfProfilesCount)
	return result, response, err

}

//GetRfProfileByID Get RF Profile by ID - 3298-aa56-4ec9-b510
/* This API allows the user to get a RF Profile by RF Profile ID


@param id id path parameter. RF Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-rf-profile-by-id
*/
func (s *WirelessService) GetRfProfileByID(id string) (*ResponseWirelessGetRfProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetRfProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetRfProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetRfProfileById")
	}

	result := response.Result().(*ResponseWirelessGetRfProfileByID)
	return result, response, err

}

//CreateAndProvisionSSID Create and Provision SSID - 1eb7-2ad3-4e09-8990
/* Creates SSID, updates the SSID to the corresponding site profiles and provision it to the devices matching the given sites


@param CreateAndProvisionSSIDHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-and-provision-ssid
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAndProvisionSSID(requestWirelessCreateAndProvisionSSID, CreateAndProvisionSSIDHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation CreateAndProvisionSsid")
	}

	result := response.Result().(*ResponseWirelessCreateAndProvisionSSID)
	return result, response, err

}

//RebootAccessPoints Reboot Access Points - 6092-d8f1-468b-99ab
/* Users can reboot multiple access points up-to 200 at a time using this API



Documentation Link: https://developer.cisco.com/docs/dna-center/#!reboot-access-points
*/
func (s *WirelessService) RebootAccessPoints(requestWirelessRebootAccessPoints *RequestWirelessRebootAccessPoints) (*ResponseWirelessRebootAccessPoints, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-reboot/apreboot"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessRebootAccessPoints).
		SetResult(&ResponseWirelessRebootAccessPoints{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RebootAccessPoints(requestWirelessRebootAccessPoints)
		}

		return nil, response, fmt.Errorf("error with operation RebootAccessPoints")
	}

	result := response.Result().(*ResponseWirelessRebootAccessPoints)
	return result, response, err

}

//CreateEnterpriseSSID Create Enterprise SSID - 8a96-fb95-4d09-a349
/* Creates enterprise SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-enterprise-ssid
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateEnterpriseSSID(requestWirelessCreateEnterpriseSSID)
		}

		return nil, response, fmt.Errorf("error with operation CreateEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessCreateEnterpriseSSID)
	return result, response, err

}

//CreateSSID Create SSID - 0193-8858-4789-9a53
/* This API allows the user to create an SSID (Service Set Identifier) at the Global site


@param siteID siteId path parameter. Site UUID of Global site


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-ssid
*/
func (s *WirelessService) CreateSSID(siteID string, requestWirelessCreateSSID *RequestWirelessCreateSSID) (*ResponseWirelessCreateSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateSSID).
		SetResult(&ResponseWirelessCreateSSID{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSSID(siteID, requestWirelessCreateSSID)
		}

		return nil, response, fmt.Errorf("error with operation CreateSsid")
	}

	result := response.Result().(*ResponseWirelessCreateSSID)
	return result, response, err

}

//ConfigureAccessPointsV1 Configure Access Points V1 - 0081-cb89-4708-888f
/* User can configure multiple access points with required options using this intent API. This API does not support configuration of CleanAir or SI for IOS-XE devices with version greater than or equal to 17.9



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-access-points-v1
*/
func (s *WirelessService) ConfigureAccessPointsV1(requestWirelessConfigureAccessPointsV1 *RequestWirelessConfigureAccessPointsV1) (*ResponseWirelessConfigureAccessPointsV1, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessConfigureAccessPointsV1).
		SetResult(&ResponseWirelessConfigureAccessPointsV1{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAccessPointsV1(requestWirelessConfigureAccessPointsV1)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAccessPointsV1")
	}

	result := response.Result().(*ResponseWirelessConfigureAccessPointsV1)
	return result, response, err

}

//ApProvision2 AP Provision - d897-19b8-47aa-a9c4
/* Access Point Provision and ReProvision


@param APProvision2HeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!ap-provision2
*/
func (s *WirelessService) ApProvision2(requestWirelessAPProvision2 *RequestWirelessApProvision2, APProvision2HeaderParams *ApProvision2HeaderParams) (*ResponseWirelessApProvision2, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/ap-provision"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if APProvision2HeaderParams != nil {

		if APProvision2HeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", APProvision2HeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessAPProvision2).
		SetResult(&ResponseWirelessApProvision2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApProvision2(requestWirelessAPProvision2, APProvision2HeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation ApProvision2")
	}

	result := response.Result().(*ResponseWirelessApProvision2)
	return result, response, err

}

//CreateUpdateDynamicInterface Create Update Dynamic interface - daa0-bb75-4e2a-8da6
/* API to create or update an dynamic interface



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-update-dynamic-interface
*/
func (s *WirelessService) CreateUpdateDynamicInterface(requestWirelessCreateUpdateDynamicInterface *RequestWirelessCreateUpdateDynamicInterface) (*ResponseWirelessCreateUpdateDynamicInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateUpdateDynamicInterface).
		SetResult(&ResponseWirelessCreateUpdateDynamicInterface{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateUpdateDynamicInterface(requestWirelessCreateUpdateDynamicInterface)
		}

		return nil, response, fmt.Errorf("error with operation CreateUpdateDynamicInterface")
	}

	result := response.Result().(*ResponseWirelessCreateUpdateDynamicInterface)
	return result, response, err

}

//CreateWirelessProfile Create Wireless Profile - 7097-6962-4bf9-88d5
/* Creates Wireless Network Profile on Cisco Catalyst Center and associates sites and SSIDs to it.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-wireless-profile
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateWirelessProfile(requestWirelessCreateWirelessProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessCreateWirelessProfile)
	return result, response, err

}

//Provision Provision - d09b-08a3-447a-a3b9
/* Provision wireless device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!provision
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.Provision(requestWirelessProvision)
		}

		return nil, response, fmt.Errorf("error with operation Provision")
	}

	result := response.Result().(*ResponseWirelessProvision)
	return result, response, err

}

//PSKOverride PSK override - 46ad-ab75-47c9-8762
/* Update/Override passphrase of SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!p-s-k-override
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.PSKOverride(requestWirelessPSKOverride)
		}

		return nil, response, fmt.Errorf("error with operation PSKOverride")
	}

	result := response.Result().(*ResponseWirelessPSKOverride)
	return result, response, err

}

//CreateOrUpdateRfProfile Create or Update RF profile - b783-2967-4878-b815
/* Create or Update RF profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-or-update-rf-profile
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateOrUpdateRfProfile(requestWirelessCreateOrUpdateRFProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateOrUpdateRfProfile")
	}

	result := response.Result().(*ResponseWirelessCreateOrUpdateRfProfile)
	return result, response, err

}

//FactoryResetAccessPoints Factory Reset Access Point(s) - b09d-4bbc-482b-aeb7
/* This API is used to factory reset Access Points. It is supported for maximum 100 Access Points per request. Factory reset clears all configurations from the Access Points. After factory reset the Access Point may become unreachable from the currently associated Wireless Controller and may or may not join back the same controller.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!factory-reset-access-points
*/
func (s *WirelessService) FactoryResetAccessPoints(requestWirelessFactoryResetAccessPoints *RequestWirelessFactoryResetAccessPoints) (*ResponseWirelessFactoryResetAccessPoints, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessAccessPoints/factoryResetRequest/provision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessFactoryResetAccessPoints).
		SetResult(&ResponseWirelessFactoryResetAccessPoints{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.FactoryResetAccessPoints(requestWirelessFactoryResetAccessPoints)
		}

		return nil, response, fmt.Errorf("error with operation FactoryResetAccessPoints")
	}

	result := response.Result().(*ResponseWirelessFactoryResetAccessPoints)
	return result, response, err

}

//ApProvision AP Provision - 11af-897a-413b-925a
/* This API is used to provision access points



Documentation Link: https://developer.cisco.com/docs/dna-center/#!ap-provision
*/
func (s *WirelessService) ApProvision(requestWirelessAPProvision *RequestWirelessApProvision) (*ResponseWirelessApProvision, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessAccessPoints/provision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAPProvision).
		SetResult(&ResponseWirelessApProvision{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApProvision(requestWirelessAPProvision)
		}

		return nil, response, fmt.Errorf("error with operation ApProvision")
	}

	result := response.Result().(*ResponseWirelessApProvision)
	return result, response, err

}

//MobilityProvision Mobility Provision - 6c8b-6bd5-40bb-ac31
/* This API is used to provision/deploy wireless mobility into Cisco wireless controllers.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mobility-provision
*/
func (s *WirelessService) MobilityProvision(requestWirelessMobilityProvision *RequestWirelessMobilityProvision) (*ResponseWirelessMobilityProvision, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups/mobilityProvision"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessMobilityProvision).
		SetResult(&ResponseWirelessMobilityProvision{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.MobilityProvision(requestWirelessMobilityProvision)
		}

		return nil, response, fmt.Errorf("error with operation MobilityProvision")
	}

	result := response.Result().(*ResponseWirelessMobilityProvision)
	return result, response, err

}

//MobilityReset Mobility Reset - e589-6baf-4caa-9bbc
/* This API is used to reset wireless mobility which in turn sets mobility group name as 'default'



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mobility-reset
*/
func (s *WirelessService) MobilityReset(requestWirelessMobilityReset *RequestWirelessMobilityReset) (*ResponseWirelessMobilityReset, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups/mobilityReset"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessMobilityReset).
		SetResult(&ResponseWirelessMobilityReset{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.MobilityReset(requestWirelessMobilityReset)
		}

		return nil, response, fmt.Errorf("error with operation MobilityReset")
	}

	result := response.Result().(*ResponseWirelessMobilityReset)
	return result, response, err

}

//AssignManagedApLocationsForWLC Assign Managed AP Locations For WLC - afbd-d880-488a-83e4
/* This API allows user to assign Managed AP Locations for WLC by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.


@param deviceID deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-managed-ap-locations-for-w-l-c
*/
func (s *WirelessService) AssignManagedApLocationsForWLC(deviceID string, requestWirelessAssignManagedAPLocationsForWLC *RequestWirelessAssignManagedApLocationsForWLC) (*ResponseWirelessAssignManagedApLocationsForWLC, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{deviceId}/assignManagedApLocations"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAssignManagedAPLocationsForWLC).
		SetResult(&ResponseWirelessAssignManagedApLocationsForWLC{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignManagedApLocationsForWLC(deviceID, requestWirelessAssignManagedAPLocationsForWLC)
		}

		return nil, response, fmt.Errorf("error with operation AssignManagedApLocationsForWLC")
	}

	result := response.Result().(*ResponseWirelessAssignManagedApLocationsForWLC)
	return result, response, err

}

//WirelessControllerProvision Wireless Controller Provision - 9e9c-386b-4069-9e7c
/* This API is used to provision wireless controller


@param deviceID deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}


Documentation Link: https://developer.cisco.com/docs/dna-center/#!wireless-controller-provision
*/
func (s *WirelessService) WirelessControllerProvision(deviceID string, requestWirelessWirelessControllerProvision *RequestWirelessWirelessControllerProvision) (*ResponseWirelessWirelessControllerProvision, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{deviceId}/provision"
	path = strings.Replace(path, "{deviceId}", fmt.Sprintf("%v", deviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessWirelessControllerProvision).
		SetResult(&ResponseWirelessWirelessControllerProvision{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.WirelessControllerProvision(deviceID, requestWirelessWirelessControllerProvision)
		}

		return nil, response, fmt.Errorf("error with operation WirelessControllerProvision")
	}

	result := response.Result().(*ResponseWirelessWirelessControllerProvision)
	return result, response, err

}

//CreateWirelessProfile2 Create Wireless Profile - dd88-bb37-492a-888b
/* This API allows the user to create a Wireless Network Profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-wireless-profile2
*/
func (s *WirelessService) CreateWirelessProfile2(requestWirelessCreateWirelessProfile2 *RequestWirelessCreateWirelessProfile2) (*ResponseWirelessCreateWirelessProfile2, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateWirelessProfile2).
		SetResult(&ResponseWirelessCreateWirelessProfile2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateWirelessProfile2(requestWirelessCreateWirelessProfile2)
		}

		return nil, response, fmt.Errorf("error with operation CreateWirelessProfile2")
	}

	result := response.Result().(*ResponseWirelessCreateWirelessProfile2)
	return result, response, err

}

//CreateA80211BeProfile Create a 802.11be Profile - efab-bbaf-4388-a046
/* This API allows the user to create a 802.11be Profile.Catalyst Center will push this profile to device's "default-dot11be-profile”.Also please note , 802.11be Profile is supported only on IOS-XE controllers since device version 17.15



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a80211be-profile
*/
func (s *WirelessService) CreateA80211BeProfile(requestWirelessCreateA80211beProfile *RequestWirelessCreateA80211BeProfile) (*ResponseWirelessCreateA80211BeProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateA80211beProfile).
		SetResult(&ResponseWirelessCreateA80211BeProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateA80211BeProfile(requestWirelessCreateA80211beProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateA80211BeProfile")
	}

	result := response.Result().(*ResponseWirelessCreateA80211BeProfile)
	return result, response, err

}

//CreateInterface Create Interface - a098-6877-44e8-ba31
/* This API allows the user to create an interface



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-interface
*/
func (s *WirelessService) CreateInterface(requestWirelessCreateInterface *RequestWirelessCreateInterface) (*ResponseWirelessCreateInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateInterface).
		SetResult(&ResponseWirelessCreateInterface{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateInterface(requestWirelessCreateInterface)
		}

		return nil, response, fmt.Errorf("error with operation CreateInterface")
	}

	result := response.Result().(*ResponseWirelessCreateInterface)
	return result, response, err

}

//CreateRfProfile Create RF Profile - 3cb0-ca20-45d9-8d07
/* This API allows the user to create a custom RF Profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-rf-profile
*/
func (s *WirelessService) CreateRfProfile(requestWirelessCreateRFProfile *RequestWirelessCreateRfProfile) (*ResponseWirelessCreateRfProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateRFProfile).
		SetResult(&ResponseWirelessCreateRfProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateRfProfile(requestWirelessCreateRFProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateRfProfile")
	}

	result := response.Result().(*ResponseWirelessCreateRfProfile)
	return result, response, err

}

//ConfigureAccessPointsV2 Configure Access Points V2 - 5ca7-4a81-4329-9506
/* User can configure multiple access points with required options using this intent API



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-access-points-v2
*/
func (s *WirelessService) ConfigureAccessPointsV2(requestWirelessConfigureAccessPointsV2 *RequestWirelessConfigureAccessPointsV2) (*ResponseWirelessConfigureAccessPointsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/wireless/accesspoint-configuration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessConfigureAccessPointsV2).
		SetResult(&ResponseWirelessConfigureAccessPointsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAccessPointsV2(requestWirelessConfigureAccessPointsV2)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAccessPointsV2")
	}

	result := response.Result().(*ResponseWirelessConfigureAccessPointsV2)
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateEnterpriseSSID(requestWirelessUpdateEnterpriseSSID)
		}
		return nil, response, fmt.Errorf("error with operation UpdateEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessUpdateEnterpriseSSID)
	return result, response, err

}

//UpdateSSID Update SSID - 2496-7ad2-4b8a-913b
/* This API allows the user to update an SSID (Service Set Identifier) at the given site


@param siteID siteId path parameter. Site UUID

@param id id path parameter. SSID ID. Inputs containing special characters should be encoded

*/
func (s *WirelessService) UpdateSSID(siteID string, id string, requestWirelessUpdateSSID *RequestWirelessUpdateSSID) (*ResponseWirelessUpdateSSID, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateSSID).
		SetResult(&ResponseWirelessUpdateSSID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSSID(siteID, id, requestWirelessUpdateSSID)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSsid")
	}

	result := response.Result().(*ResponseWirelessUpdateSSID)
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWirelessProfile(requestWirelessUpdateWirelessProfile)
		}
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProvisionUpdate(requestWirelessProvisionUpdate, ProvisionUpdateHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation ProvisionUpdate")
	}

	result := response.Result().(*ResponseWirelessProvisionUpdate)
	return result, response, err

}

//UpdateWirelessProfile2 Update Wireless Profile - 4f88-d9a3-4ef8-8e2e
/* This API allows the user to update a Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id

*/
func (s *WirelessService) UpdateWirelessProfile2(id string, requestWirelessUpdateWirelessProfile2 *RequestWirelessUpdateWirelessProfile2) (*ResponseWirelessUpdateWirelessProfile2, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateWirelessProfile2).
		SetResult(&ResponseWirelessUpdateWirelessProfile2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWirelessProfile2(id, requestWirelessUpdateWirelessProfile2)
		}
		return nil, response, fmt.Errorf("error with operation UpdateWirelessProfile2")
	}

	result := response.Result().(*ResponseWirelessUpdateWirelessProfile2)
	return result, response, err

}

//Update80211BeProfile Update 802.11be Profile - 699b-b8e0-48bb-9b90
/* This API allows the user to update a 802.11be Profile


@param id id path parameter. 802.11be Profile ID

*/
func (s *WirelessService) Update80211BeProfile(id string, requestWirelessUpdate80211beProfile *RequestWirelessUpdate80211BeProfile) (*ResponseWirelessUpdate80211BeProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdate80211beProfile).
		SetResult(&ResponseWirelessUpdate80211BeProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Update80211BeProfile(id, requestWirelessUpdate80211beProfile)
		}
		return nil, response, fmt.Errorf("error with operation Update80211BeProfile")
	}

	result := response.Result().(*ResponseWirelessUpdate80211BeProfile)
	return result, response, err

}

//UpdateInterface Update Interface - 0f93-9868-454b-a943
/* This API allows the user to update an interface by ID


@param id id path parameter. Interface ID

*/
func (s *WirelessService) UpdateInterface(id string, requestWirelessUpdateInterface *RequestWirelessUpdateInterface) (*ResponseWirelessUpdateInterface, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateInterface).
		SetResult(&ResponseWirelessUpdateInterface{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateInterface(id, requestWirelessUpdateInterface)
		}
		return nil, response, fmt.Errorf("error with operation UpdateInterface")
	}

	result := response.Result().(*ResponseWirelessUpdateInterface)
	return result, response, err

}

//UpdateRfProfile Update RF Profile - 2984-b995-4ae9-b3c3
/* This API allows the user to update a custom RF Profile


@param id id path parameter. RF Profile ID

*/
func (s *WirelessService) UpdateRfProfile(id string, requestWirelessUpdateRFProfile *RequestWirelessUpdateRfProfile) (*ResponseWirelessUpdateRfProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateRFProfile).
		SetResult(&ResponseWirelessUpdateRfProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRfProfile(id, requestWirelessUpdateRFProfile)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRfProfile")
	}

	result := response.Result().(*ResponseWirelessUpdateRfProfile)
	return result, response, err

}

//DeleteSSIDAndProvisionItToDevices Delete SSID and provision it to devices - fc95-38fe-43d9-884d
/* Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from Catalyst Center


@param ssidName ssidName path parameter. SSID Name. This parameter needs to be encoded as per UTF-8 encoding.

@param managedAPLocations managedAPLocations path parameter. List of managed AP locations (Site Hierarchies). This parameter needs to be encoded as per UTF-8 encoding

@param DeleteSSIDAndProvisionItToDevicesHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ssid-and-provision-it-to-devices
*/
func (s *WirelessService) DeleteSSIDAndProvisionItToDevices(ssidName string, managedAPLocations string, DeleteSSIDAndProvisionItToDevicesHeaderParams *DeleteSSIDAndProvisionItToDevicesHeaderParams) (*ResponseWirelessDeleteSSIDAndProvisionItToDevices, *resty.Response, error) {
	//ssidName string,managedAPLocations string,DeleteSSIDAndProvisionItToDevicesHeaderParams *DeleteSSIDAndProvisionItToDevicesHeaderParams
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSSIDAndProvisionItToDevices(ssidName, managedAPLocations, DeleteSSIDAndProvisionItToDevicesHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSsidAndProvisionItToDevices")
	}

	result := response.Result().(*ResponseWirelessDeleteSSIDAndProvisionItToDevices)
	return result, response, err

}

//DeleteEnterpriseSSID Delete Enterprise SSID - c7a6-592b-4b98-a369
/* Deletes given enterprise SSID


@param ssidName ssidName path parameter. Enter the SSID name to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-enterprise-ssid
*/
func (s *WirelessService) DeleteEnterpriseSSID(ssidName string) (*ResponseWirelessDeleteEnterpriseSSID, *resty.Response, error) {
	//ssidName string
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteEnterpriseSSID(ssidName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessDeleteEnterpriseSSID)
	return result, response, err

}

//DeleteSSID Delete SSID - acbe-8b6f-4e8b-9f6a
/* This API allows the user to delete an SSID (Service Set Identifier) at the global level, if the SSID is not mapped to any Wireless Profile


@param siteID siteId path parameter. Site UUID where SSID is to be deleted

@param id id path parameter. SSID ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ssid
*/
func (s *WirelessService) DeleteSSID(siteID string, id string) (*ResponseWirelessDeleteSSID, *resty.Response, error) {
	//siteID string,id string
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteSSID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSSID(siteID, id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSsid")
	}

	result := response.Result().(*ResponseWirelessDeleteSSID)
	return result, response, err

}

//DeleteWirelessProfile Delete Wireless Profile - e395-88a5-4949-82c4
/* Delete the Wireless Profile whose name is provided.


@param wirelessProfileName wirelessProfileName path parameter. Wireless Profile Name


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-wireless-profile
*/
func (s *WirelessService) DeleteWirelessProfile(wirelessProfileName string) (*ResponseWirelessDeleteWirelessProfile, *resty.Response, error) {
	//wirelessProfileName string
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteWirelessProfile(wirelessProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessDeleteWirelessProfile)
	return result, response, err

}

//DeleteDynamicInterface Delete dynamic interface - ffb4-abf4-44fb-b70a
/* Delete a dynamic interface


@param DeleteDynamicInterfaceHeaderParams Custom header parameters
@param DeleteDynamicInterfaceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-dynamic-interface
*/
func (s *WirelessService) DeleteDynamicInterface(DeleteDynamicInterfaceHeaderParams *DeleteDynamicInterfaceHeaderParams, DeleteDynamicInterfaceQueryParams *DeleteDynamicInterfaceQueryParams) (*ResponseWirelessDeleteDynamicInterface, *resty.Response, error) {
	//DeleteDynamicInterfaceHeaderParams *DeleteDynamicInterfaceHeaderParams,DeleteDynamicInterfaceQueryParams *DeleteDynamicInterfaceQueryParams
	path := "/dna/intent/api/v1/wireless/dynamic-interface"

	queryString, _ := query.Values(DeleteDynamicInterfaceQueryParams)

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
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessDeleteDynamicInterface{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteDynamicInterface(DeleteDynamicInterfaceHeaderParams, DeleteDynamicInterfaceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteDynamicInterface")
	}

	result := response.Result().(*ResponseWirelessDeleteDynamicInterface)
	return result, response, err

}

//DeleteRfProfiles Delete RF profiles - 28b2-4a74-4a99-94be
/* Delete RF profile


@param rfProfileName rfProfileName path parameter. RF profile name to be deleted(required) *non-custom RF profile cannot be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-rf-profiles
*/
func (s *WirelessService) DeleteRfProfiles(rfProfileName string) (*ResponseWirelessDeleteRfProfiles, *resty.Response, error) {
	//rfProfileName string
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteRfProfiles(rfProfileName)
		}
		return nil, response, fmt.Errorf("error with operation DeleteRfProfiles")
	}

	result := response.Result().(*ResponseWirelessDeleteRfProfiles)
	return result, response, err

}

//DeleteWirelessProfile2 Delete Wireless Profile - 289c-f9f5-4f78-b84c
/* This API allows the user to delete Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-wireless-profile2
*/
func (s *WirelessService) DeleteWirelessProfile2(id string) (*ResponseWirelessDeleteWirelessProfile2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteWirelessProfile2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteWirelessProfile2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteWirelessProfile2")
	}

	result := response.Result().(*ResponseWirelessDeleteWirelessProfile2)
	return result, response, err

}

//DeleteA80211BeProfile Delete a 802.11be Profile - e9b0-98c2-4b49-8fe6
/* This API allows the user to delete a 802.11be Profile,if the 802.11be Profile is not mapped to any Wireless Network Profile


@param id id path parameter. 802.11be Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a80211be-profile
*/
func (s *WirelessService) DeleteA80211BeProfile(id string) (*ResponseWirelessDeleteA80211BeProfile, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteA80211BeProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteA80211BeProfile(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteA80211BeProfile")
	}

	result := response.Result().(*ResponseWirelessDeleteA80211BeProfile)
	return result, response, err

}

//DeleteInterface Delete Interface - 0999-c9cd-4159-a6a1
/* This API allows the user to delete an interface by ID


@param id id path parameter. Interface ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-interface
*/
func (s *WirelessService) DeleteInterface(id string) (*ResponseWirelessDeleteInterface, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/interfaces/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteInterface{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteInterface(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteInterface")
	}

	result := response.Result().(*ResponseWirelessDeleteInterface)
	return result, response, err

}

//DeleteRfProfile Delete RF Profile - 2f8a-799d-4fa9-ac0e
/* This API allows the user to delete a custom RF Profile


@param id id path parameter. RF Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-rf-profile
*/
func (s *WirelessService) DeleteRfProfile(id string) (*ResponseWirelessDeleteRfProfile, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/rfProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteRfProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteRfProfile(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteRfProfile")
	}

	result := response.Result().(*ResponseWirelessDeleteRfProfile)
	return result, response, err

}
