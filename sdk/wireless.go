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
	Limit      float64 `url:"limit,omitempty"`      //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset     float64 `url:"offset,omitempty"`     //The first record to show for this page; the first record is numbered 1.
	SSID       string  `url:"ssid,omitempty"`       //SSID Name
	WLANType   string  `url:"wlanType,omitempty"`   //Wlan Type
	AuthType   string  `url:"authType,omitempty"`   //Auth Type
	L3AuthType string  `url:"l3authType,omitempty"` //L3 Auth Type
}
type GetSSIDCountBySiteQueryParams struct {
	Inherited bool `url:"_inherited,omitempty"` //This query parameter indicates whether the current SSID count at the given 'siteId' is of the SSID(s) it is inheriting or count of non-inheriting SSID(s)
}
type DeleteSSIDQueryParams struct {
	RemoveOverrideInHierarchy bool `url:"removeOverrideInHierarchy,omitempty"` //Remove override in hierarchy . Refer Feature tab for details
}
type GetAccessPointConfigurationCountQueryParams struct {
	WlcIPAddress string `url:"wlcIpAddress,omitempty"` //WLC IP Address
	ApMode       string `url:"apMode,omitempty"`       //AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
	ApModel      string `url:"apModel,omitempty"`      //AP Model
	MeshRole     string `url:"meshRole,omitempty"`     //Mesh Role. Allowed values are RAP or MAP
	Provisioned  string `url:"provisioned,omitempty"`  //Indicate whether AP provisioned or not. Allowed values are True or False
}
type GetAccessPointConfigurationQueryParams struct {
	Key          string  `url:"key,omitempty"`          //The ethernet MAC address of Access point
	WlcIPAddress string  `url:"wlcIpAddress,omitempty"` //WLC IP Address
	ApMode       string  `url:"apMode,omitempty"`       //AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
	ApModel      string  `url:"apModel,omitempty"`      //AP Model
	MeshRole     string  `url:"meshRole,omitempty"`     //Mesh Role. Allowed values are RAP or MAP
	Provisioned  string  `url:"provisioned,omitempty"`  //Indicate whether AP provisioned or not. Allowed values are True or False
	Limit        float64 `url:"limit,omitempty"`        //The number of records to show for this page. The default is 500 if not specified. The maximum allowed limit is 500.
	Offset       float64 `url:"offset,omitempty"`       //The first record to show for this page; the first record is numbered 1.
}
type ApProvisionConnectivityHeaderParams struct {
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
type GetMeshApNeighboursQueryParams struct {
	WlcIPAddress       string `url:"wlcIpAddress,omitempty"`       //Employ this query parameter to obtain the details of the Access points corresponding to the provided WLC IP address.
	ApName             string `url:"apName,omitempty"`             //Employ this query parameter to obtain the details of the Access points corresponding to the provided ap name.
	EthernetMacAddress string `url:"ethernetMacAddress,omitempty"` //Employ this query parameter to obtain the details of the Access points corresponding to the provided EthernetMacAddress.
}
type GetMobilityGroupsQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Employ this query parameter to obtain the details of the Mobility Group corresponding to the provided networkDeviceId. Obtain the network device ID value by using the API GET call /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
}
type GetAnchorManagedApLocationsForSpecificWirelessControllerQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetPrimaryManagedApLocationsForSpecificWirelessControllerQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetSecondaryManagedApLocationsForSpecificWirelessControllerQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetSSIDDetailsForSpecificWirelessControllerQueryParams struct {
	SSIDName    string  `url:"ssidName,omitempty"`    //Employ this query parameter to obtain the details of the SSID corresponding to the provided SSID name.
	AdminStatus bool    `url:"adminStatus,omitempty"` //Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed     bool    `url:"managed,omitempty"`     //If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1.
}
type GetSSIDCountForSpecificWirelessControllerQueryParams struct {
	AdminStatus bool `url:"adminStatus,omitempty"` //Utilize this query parameter to obtain the number of SSIDs according to their administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.
	Managed     bool `url:"managed,omitempty"`     //If value is 'true' means SSIDs are configured through design.If the value is 'false' means out of band configuration from the Wireless Controller.
}
type GetWirelessProfilesQueryParams struct {
	Limit               float64 `url:"limit,omitempty"`               //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
	Offset              float64 `url:"offset,omitempty"`              //The first record to show for this page; the first record is numbered 1
	WirelessProfileName string  `url:"wirelessProfileName,omitempty"` //Wireless Profile Name
}
type RetrieveAllPolicyTagsForAWirelessProfileQueryParams struct {
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset        float64 `url:"offset,omitempty"`        //Offset
	PolicyTagName string  `url:"policyTagName,omitempty"` //PolicyTagName
}
type RetrieveAllSiteTagsForAWirelessProfileQueryParams struct {
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset      float64 `url:"offset,omitempty"`      //Offset
	SiteTagName string  `url:"siteTagName,omitempty"` //SiteTagName
}
type GetApAuthorizationListsQueryParams struct {
	ApAuthorizationListName string `url:"apAuthorizationListName,omitempty"` //Employ this query parameter to obtain the details of the AP Authorization List corresponding to the provided apAuthorizationListName.
	Offset                  string `url:"offset,omitempty"`                  //The first record to show for this page. The first record is numbered 1.
	Limit                   string `url:"limit,omitempty"`                   //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
}
type GetApProfilesQueryParams struct {
	Limit         string `url:"limit,omitempty"`         //The number of records to show for this page. The default is 500 if not specified. The maximum allowed limit is 500.
	Offset        string `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
	ApProfileName string `url:"apProfileName,omitempty"` //Employ this query parameter to obtain the details of the apProfiles corresponding to the provided apProfileName.
}
type Get80211BeProfilesQueryParams struct {
	Limit            float64 `url:"limit,omitempty"`            //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset           float64 `url:"offset,omitempty"`           //The first record to show for this page, the first record is numbered 1
	ProfileName      string  `url:"profileName,omitempty"`      //Profile Name
	IsOfDmaDownLink  bool    `url:"isOfDmaDownLink,omitempty"`  //OFDMA Downlink
	IsOfDmaUpLink    bool    `url:"isOfDmaUpLink,omitempty"`    //OFDMA Uplink
	IsMuMimoUpLink   bool    `url:"isMuMimoUpLink,omitempty"`   //MU-MIMO Uplink
	IsMuMimoDownLink bool    `url:"isMuMimoDownLink,omitempty"` //MU-MIMO Downlink
	IsOfDmaMultiRu   bool    `url:"isOfDmaMultiRu,omitempty"`   //OFDMA Multi-RU
}
type GetInterfacesQueryParams struct {
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset        float64 `url:"offset,omitempty"`        //The first record to show for this page. The first record is numbered 1.
	InterfaceName string  `url:"interfaceName,omitempty"` //Interface Name
	VLANID        float64 `url:"vlanId,omitempty"`        //Vlan Id
}
type GetPowerProfilesQueryParams struct {
	Limit       float64 `url:"limit,omitempty"`       //Limit
	Offset      float64 `url:"offset,omitempty"`      //Offset
	ProfileName string  `url:"profileName,omitempty"` //Power Profile Name
}
type GetRfProfilesQueryParams struct {
	Limit               float64 `url:"limit,omitempty"`               //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
	Offset              float64 `url:"offset,omitempty"`              //The first record to show for this page; the first record is numbered 1
	RfProfileName       string  `url:"rfProfileName,omitempty"`       //RF Profile Name
	EnableRadioTypeA    bool    `url:"enableRadioTypeA,omitempty"`    //Enable Radio TypeA
	EnableRadioTypeB    bool    `url:"enableRadioTypeB,omitempty"`    //Enable Radio TypeB
	EnableRadioType6GHz bool    `url:"enableRadioType6GHz,omitempty"` //Enable Radio Type6GHz
}
type RetrieveSitesWithOverriddenSSIDsQueryParams struct {
	SiteID string  `url:"siteId,omitempty"` //Site UUID
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
}

type ResponseWirelessSensorTestResults struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseWirelessSensorTestResultsResponse `json:"response,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponse struct {
	Summary *ResponseWirelessSensorTestResultsResponseSummary `json:"summary,omitempty"` //

	FailureStats *[]ResponseWirelessSensorTestResultsResponseFailureStats `json:"failureStats,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummary struct {
	TotalTestCount *int `json:"totalTestCount,omitempty"` // Total test count

	OnBoarding *ResponseWirelessSensorTestResultsResponseSummaryOnBoarding `json:"ONBOARDING,omitempty"` //

	PERfORMAncE *ResponseWirelessSensorTestResultsResponseSummaryPERfORMAncE `json:"PERFORMANCE,omitempty"` //

	NETWORKSERVICES *ResponseWirelessSensorTestResultsResponseSummaryNETWORKSERVICES `json:"NETWORK_SERVICES,omitempty"` //

	ApPCONNECTIVITY *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITY `json:"APP_CONNECTIVITY,omitempty"` //

	RfASSESSMENT *ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENT `json:"RF_ASSESSMENT,omitempty"` //

	Email *ResponseWirelessSensorTestResultsResponseSummaryEmail `json:"EMAIL,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoarding struct {
	Auth *ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAuth `json:"AUTH,omitempty"` //

	DHCP *ResponseWirelessSensorTestResultsResponseSummaryOnBoardingDHCP `json:"DHCP,omitempty"` //

	Assoc *ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAssoc `json:"ASSOC,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoardingAuth struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryOnBoardingDHCP struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

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
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITY struct {
	HOSTREACHABILITY *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYHOSTREACHABILITY `json:"HOST_REACHABILITY,omitempty"` //

	WebServer *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYWebServer `json:"WEBSERVER,omitempty"` //

	FileTransfer *ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYFileTransfer `json:"FILETRANSFER,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYHOSTREACHABILITY struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYWebServer struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryApPCONNECTIVITYFileTransfer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENT struct {
	DATARATE *ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTDATARATE `json:"DATA_RATE,omitempty"` //

	SNR *ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTSNR `json:"SNR,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTDATARATE struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryRfASSESSMENTSNR struct {
	PassCount *int `json:"passCount,omitempty"` // Total passed test count

	FailCount *float64 `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseSummaryEmail struct {
	MailServer *ResponseWirelessSensorTestResultsResponseSummaryEmailMailServer `json:"MAILSERVER,omitempty"` //
}
type ResponseWirelessSensorTestResultsResponseSummaryEmailMailServer struct {
	PassCount *float64 `json:"passCount,omitempty"` // Total passed test count

	FailCount *int `json:"failCount,omitempty"` // Total failed test count
}
type ResponseWirelessSensorTestResultsResponseFailureStats struct {
	ErrorCode *int `json:"errorCode,omitempty"` // The error code

	ErrorTitle string `json:"errorTitle,omitempty"` // The error title

	TestType string `json:"testType,omitempty"` // The test type

	TestCategory string `json:"testCategory,omitempty"` // The test category
}
type ResponseWirelessCreateAndProvisionSSID struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteSSIDAndProvisionItToDevices struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessRebootAccessPoints struct {
	Response *ResponseWirelessRebootAccessPointsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseWirelessRebootAccessPointsResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseWirelessGetAccessPointRebootTaskResult []ResponseItemWirelessGetAccessPointRebootTaskResult // Array of ResponseWirelessGetAccessPointRebootTaskResult
type ResponseItemWirelessGetAccessPointRebootTaskResult struct {
	WlcIP string `json:"wlcIP,omitempty"` //

	ApList *[]ResponseItemWirelessGetAccessPointRebootTaskResultApList `json:"apList,omitempty"` //
}
type ResponseItemWirelessGetAccessPointRebootTaskResultApList struct {
	ApName string `json:"apName,omitempty"` //

	RebootStatus string `json:"rebootStatus,omitempty"` //

	FailureReason *ResponseItemWirelessGetAccessPointRebootTaskResultApListFailureReason `json:"failureReason,omitempty"` //
}
type ResponseItemWirelessGetAccessPointRebootTaskResultApListFailureReason interface{}
type ResponseWirelessGetEnterpriseSSID []ResponseItemWirelessGetEnterpriseSSID // Array of ResponseWirelessGetEnterpriseSSID
type ResponseItemWirelessGetEnterpriseSSID struct {
	InstanceUUID string `json:"instanceUuid,omitempty"` // Instance Uuid

	Version *int `json:"version,omitempty"` // Version

	SSIDDetails *[]ResponseItemWirelessGetEnterpriseSSIDSSIDDetails `json:"ssidDetails,omitempty"` //

	GroupUUID string `json:"groupUuid,omitempty"` // Group Uuid

	InheritedGroupUUID string `json:"inheritedGroupUuid,omitempty"` // Inherited Group Uuid

	InheritedGroupName string `json:"inheritedGroupName,omitempty"` // Inherited Group Name
}
type ResponseItemWirelessGetEnterpriseSSIDSSIDDetails struct {
	Name string `json:"name,omitempty"` // SSID Name

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable Fast Lane

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level

	AuthServer string `json:"authServer,omitempty"` // Auth Server

	Passphrase string `json:"passphrase,omitempty"` // Passphrase

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	IsEnabled *bool `json:"isEnabled,omitempty"` // Is Enabled

	IsFabric *bool `json:"isFabric,omitempty"` // Is Fabric

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcast SSID

	NasOptions []string `json:"nasOptions,omitempty"` // Nas Options

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Aaa Override

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // Protected Management Frame

	MultipSKSettings *[]ResponseItemWirelessGetEnterpriseSSIDSSIDDetailsMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *float64 `json:"clientRateLimit,omitempty"` // Client Rate Limit. (in bits per second)

	EnableSessionTimeOut *bool `json:"enableSessionTimeOut,omitempty"` // Enable Session Time Out

	SessionTimeOut *float64 `json:"sessionTimeOut,omitempty"` // sessionTimeOut

	EnableClientExclusion *bool `json:"enableClientExclusion,omitempty"` // Enable Client Exclusion

	ClientExclusionTimeout *float64 `json:"clientExclusionTimeout,omitempty"` // Client Exclusion Timeout

	EnableBasicServiceSetMaxIDle *bool `json:"enableBasicServiceSetMaxIdle,omitempty"` // Enable Basic Service Set Max Idle

	BasicServiceSetClientIDleTimeout *float64 `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set ClientIdle Timeout

	EnableDirectedMulticastService *bool `json:"enableDirectedMulticastService,omitempty"` // Enable Directed MulticastService

	EnableNeighborList *bool `json:"enableNeighborList,omitempty"` // Enable NeighborList

	MfpClientProtection string `json:"mfpClientProtection,omitempty"` // Mfp Client Protection
}
type ResponseItemWirelessGetEnterpriseSSIDSSIDDetailsMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}
type ResponseWirelessCreateEnterpriseSSID struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessUpdateEnterpriseSSID struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteEnterpriseSSID struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessCreateSSID struct {
	Response *ResponseWirelessCreateSSIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateSSIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetSSIDBySite struct {
	Response *[]ResponseWirelessGetSSIDBySiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseWirelessGetSSIDBySiteResponse struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // True if MAC Filtering is enabled, else False

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name. If not passed, profileName value will be used to populate this parameter

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]ResponseWirelessGetSSIDBySiteResponseMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	IsCustomNasIDOptions *bool `json:"isCustomNasIdOptions,omitempty"` // Set to true if Custom NAS ID Options provided

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Site UUID from where the SSID is inherited

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Site Name from where the SSID is inherited

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // L3 Authentication Type

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // Cckm TImestamp Tolerance(in milliseconds)

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsSensorPnp *bool `json:"isSensorPnp,omitempty"` // True if SSID is a sensor SSID

	ID string `json:"id,omitempty"` // SSID ID

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true

	InheritedSiteNameHierarchy string `json:"inheritedSiteNameHierarchy,omitempty"` // Inherited Site Name Hierarchy
}
type ResponseWirelessGetSSIDBySiteResponseMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type ResponseWirelessGetSSIDCountBySite struct {
	Response *ResponseWirelessGetSSIDCountBySiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetSSIDCountBySiteResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetSSIDByID struct {
	Response *ResponseWirelessGetSSIDByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseWirelessGetSSIDByIDResponse struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name. If not passed, profileName value will be used to populate this parameter

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]ResponseWirelessGetSSIDByIDResponseMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activatedTrue if RSN Cipher Suite GCMP128 is enabled, else False

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	IsCustomNasIDOptions *bool `json:"isCustomNasIdOptions,omitempty"` // Set to true if Custom NAS ID Options provided

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Site UUID from where the SSID is inherited

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Site Name from where the SSID is inherited

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // L3 Authentication Type

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // Cckm TImestamp Tolerance(in milliseconds)

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsSensorPnp *bool `json:"isSensorPnp,omitempty"` // True if SSID is a sensor SSID

	ID string `json:"id,omitempty"` // SSID ID

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true

	InheritedSiteNameHierarchy string `json:"inheritedSiteNameHierarchy,omitempty"` // Inherited Site Name Hierarchy

	InheritedSiteUUID string `json:"inheritedSiteUUID,omitempty"` // Inherited Site UUID
}
type ResponseWirelessGetSSIDByIDResponseMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type ResponseWirelessUpdateSSID struct {
	Response *ResponseWirelessUpdateSSIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateSSIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessDeleteSSID struct {
	Response *ResponseWirelessDeleteSSIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteSSIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateOrOverridessid struct {
	Response *ResponseWirelessUpdateOrOverridessidResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateOrOverridessidResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessDeleteWirelessProfile struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessConfigureAccessPoints struct {
	Response *ResponseWirelessConfigureAccessPointsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseWirelessConfigureAccessPointsResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseWirelessGetAccessPointConfigurationCount struct {
	Response *ResponseWirelessGetAccessPointConfigurationCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetAccessPointConfigurationCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetAccessPointConfigurationTaskResult []ResponseItemWirelessGetAccessPointConfigurationTaskResult // Array of ResponseWirelessGetAccessPointConfigurationTaskResult
type ResponseItemWirelessGetAccessPointConfigurationTaskResult struct {
	InstanceUUID *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUUID `json:"instanceUuid,omitempty"` //

	InstanceID *float64 `json:"instanceId,omitempty"` //

	AuthEntityID *ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityID `json:"authEntityId,omitempty"` //

	DisplayName string `json:"displayName,omitempty"` //

	AuthEntityClass *ResponseItemWirelessGetAccessPointConfigurationTaskResultAuthEntityClass `json:"authEntityClass,omitempty"` //

	InstanceTenantID string `json:"instanceTenantId,omitempty"` //

	OrderedListOEIndex *float64 `json:"_orderedListOEIndex,omitempty"` //

	OrderedListOEAssocName *ResponseItemWirelessGetAccessPointConfigurationTaskResultOrderedListOEAssocName `json:"_orderedListOEAssocName,omitempty"` //

	CreationOrderIndex *float64 `json:"_creationOrderIndex,omitempty"` //

	IsBeingChanged *bool `json:"_isBeingChanged,omitempty"` //

	DeployPending string `json:"deployPending,omitempty"` //

	InstanceCreatedOn *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceCreatedOn `json:"instanceCreatedOn,omitempty"` //

	InstanceUpdatedOn *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceUpdatedOn `json:"instanceUpdatedOn,omitempty"` //

	ChangeLogList *ResponseItemWirelessGetAccessPointConfigurationTaskResultChangeLogList `json:"changeLogList,omitempty"` //

	InstanceOrigin *ResponseItemWirelessGetAccessPointConfigurationTaskResultInstanceOrigin `json:"instanceOrigin,omitempty"` //

	LazyLoadedEntities *ResponseItemWirelessGetAccessPointConfigurationTaskResultLazyLoadedEntities `json:"lazyLoadedEntities,omitempty"` //

	InstanceVersion *float64 `json:"instanceVersion,omitempty"` //

	ApName string `json:"apName,omitempty"` //

	ControllerName string `json:"controllerName,omitempty"` //

	LocationHeirarchy string `json:"locationHeirarchy,omitempty"` //

	MacAddress string `json:"macAddress,omitempty"` //

	Status string `json:"status,omitempty"` //

	StatusDetails string `json:"statusDetails,omitempty"` //

	InternalKey *ResponseItemWirelessGetAccessPointConfigurationTaskResultInternalKey `json:"internalKey,omitempty"` //
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
	Type string `json:"type,omitempty"` //

	ID *float64 `json:"id,omitempty"` //

	LongType string `json:"longType,omitempty"` //

	URL string `json:"url,omitempty"` //
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
type ResponseWirelessApProvisionConnectivity struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteDynamicInterface struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessCreateUpdateDynamicInterface struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status URL

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessGetDynamicInterface []ResponseItemWirelessGetDynamicInterface // Array of ResponseWirelessGetDynamicInterface
type ResponseItemWirelessGetDynamicInterface struct {
	InterfaceName string `json:"interfaceName,omitempty"` // dynamic interface name

	VLANID *float64 `json:"vlanId,omitempty"` // Vlan id
}
type ResponseWirelessUpdateWirelessProfile struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessCreateWirelessProfile struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessGetWirelessProfile []ResponseItemWirelessGetWirelessProfile // Array of ResponseWirelessGetWirelessProfile
type ResponseItemWirelessGetWirelessProfile struct {
	ProfileDetails *ResponseItemWirelessGetWirelessProfileProfileDetails `json:"profileDetails,omitempty"` //
}
type ResponseItemWirelessGetWirelessProfileProfileDetails struct {
	Name string `json:"name,omitempty"` // Profile Name

	Sites []string `json:"sites,omitempty"` // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])

	SSIDDetails *[]ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetails struct {
	Name string `json:"name,omitempty"` // SSID Name

	Type string `json:"type,omitempty"` // SSID Type(enum: Enterprise/Guest)

	EnableFabric *bool `json:"enableFabric,omitempty"` // true if fabric is enabled else false

	FlexConnect *ResponseItemWirelessGetWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name
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

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local To VLAN ID
}
type ResponseWirelessProvisionUpdate struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessProvision struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessPSKOverride struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessRetrieveRfProfiles struct {
	Name string `json:"name,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // is Default Rf Profile

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // Enable Radio Type A

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // Enable Radio Type B

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	EnableCustom *bool `json:"enableCustom,omitempty"` // Enable Custom

	EnableBrownField *bool `json:"enableBrownField,omitempty"` // Enable Brown Field

	RadioTypeAProperties *ResponseWirelessRetrieveRfProfilesRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *ResponseWirelessRetrieveRfProfilesRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioTypeCProperties *ResponseWirelessRetrieveRfProfilesRadioTypeCProperties `json:"radioTypeCProperties,omitempty"` //

	EnableRadioTypeC *bool `json:"enableRadioTypeC,omitempty"` // Enable Radio Type C (6GHz)
}
type ResponseWirelessRetrieveRfProfilesRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")

	DataRates string `json:"dataRates,omitempty"` // Data Rates (Default : "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates (Default: "6,12,24")

	PowerThreshold *float64 `json:"powerThreshold,omitempty"` // Power Threshold  ( (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Rx Sop Threshold  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type ResponseWirelessRetrieveRfProfilesRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "9,11,12,18,24,36,48,54")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "9,11,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "12")

	PowerThreshold *float64 `json:"powerThreshold,omitempty"` // Power Threshold   (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type ResponseWirelessRetrieveRfProfilesRadioTypeCProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "6,12,24")

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)

	PowerThreshold *float64 `json:"powerThreshold,omitempty"` // Power Threshold   (Default: -70)
}
type ResponseWirelessCreateOrUpdateRfProfile struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessDeleteRfProfiles struct {
	ExecutionID string `json:"executionId,omitempty"` // Execution Id

	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url

	Message string `json:"message,omitempty"` // Message
}
type ResponseWirelessFactoryResetAccessPoints struct {
	Response *ResponseWirelessFactoryResetAccessPointsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessFactoryResetAccessPointsResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseWirelessGetAccessPointsFactoryResetStatus struct {
	Response *[]ResponseWirelessGetAccessPointsFactoryResetStatusResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetAccessPointsFactoryResetStatusResponse struct {
	WlcIP string `json:"wlcIP,omitempty"` // Wireless Controller IP address

	WlcName string `json:"wlcName,omitempty"` // Wireless Controller name

	ApResponseInfoList *[]ResponseWirelessGetAccessPointsFactoryResetStatusResponseApResponseInfoList `json:"apResponseInfoList,omitempty"` //
}
type ResponseWirelessGetAccessPointsFactoryResetStatusResponseApResponseInfoList struct {
	ApName string `json:"apName,omitempty"` // Access Point name

	ApFactoryResetStatus string `json:"apFactoryResetStatus,omitempty"` // AP factory reset status, "Success" or "Failure" or "In Progress"

	FailureReason string `json:"failureReason,omitempty"` // Reason for failure if the factory reset status is "Failure"

	RadioMacAddress string `json:"radioMacAddress,omitempty"` // AP Radio Mac Address

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"` // AP Ethernet Mac Address
}
type ResponseWirelessApProvision struct {
	Response *ResponseWirelessApProvisionResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessApProvisionResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetAnchorCapableDevices struct {
	DeviceIP string `json:"deviceIp,omitempty"` // Anchor Controller Ip

	DeviceName string `json:"deviceName,omitempty"` // Anchor Controller host name

	WirelessMgmtIP string `json:"wirelessMgmtIP,omitempty"` // Wireless management Ip Address
}
type ResponseWirelessGetMeshApNeighbours struct {
	ID string `json:"id,omitempty"` // id

	ApName string `json:"apName,omitempty"` // Name of the Wireless Access point

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"` // AP Ethernet MacAddress mac

	NeighbourMacAddress string `json:"neighbourMacAddress,omitempty"` // AP Base Radio MacAddress mac.

	WlcIPAddress string `json:"wlcIpAddress,omitempty"` // Device wireless Management IP

	NeighbourType string `json:"neighbourType,omitempty"` // Neighbour Type

	MeshRole string `json:"meshRole,omitempty"` // Mesh Role
}
type ResponseWirelessGetMeshApNeighboursCount struct {
	Response *ResponseWirelessGetMeshApNeighboursCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetMeshApNeighboursCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetMobilityGroups struct {
	Response *[]ResponseWirelessGetMobilityGroupsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version.
}
type ResponseWirelessGetMobilityGroupsResponse struct {
	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Self device Group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	MacAddress string `json:"macAddress,omitempty"` // Device mobility MAC Address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ManagementIP string `json:"managementIp,omitempty"` // Self device wireless Management IP.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

	DtlsHighCipher *bool `json:"dtlsHighCipher,omitempty"` // DTLS High Cipher.

	DataLinkEncryption *bool `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.

	MobilityPeers *[]ResponseWirelessGetMobilityGroupsResponseMobilityPeers `json:"mobilityPeers,omitempty"` //
}
type ResponseWirelessGetMobilityGroupsResponseMobilityPeers struct {
	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PeerNetworkDeviceID string `json:"peerNetworkDeviceId,omitempty"` // Peer device Id. The possible values are UNKNOWN or valid UUID of Network device ID. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device ID.

	MemberMacAddress string `json:"memberMacAddress,omitempty"` // Peer device mobility MAC Address.  Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	DeviceSeries string `json:"deviceSeries,omitempty"` // Peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.

	DataLinkEncryption *bool `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers.

	HashKey string `json:"hashKey,omitempty"` // SSC hash string must be 40 characters.

	Status string `json:"status,omitempty"` // Possible values are - Control and Data Path Down, Data Path Down, Control Path Down, UP.

	PeerIP string `json:"peerIp,omitempty"` // This indicates public IP address.

	PrivateIPAddress string `json:"privateIpAddress,omitempty"` // This indicates private/management IP address.
}
type ResponseWirelessGetMobilityGroupsCount struct {
	Response *ResponseWirelessGetMobilityGroupsCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version.
}
type ResponseWirelessGetMobilityGroupsCountResponse struct {
	Count *int `json:"count,omitempty"` // Total number of mobility groups available.
}
type ResponseWirelessMobilityProvision struct {
	Response *ResponseWirelessMobilityProvisionResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessMobilityProvisionResponse struct {
	TaskID string `json:"taskId,omitempty"` // Asynchronous Task Id

	URL string `json:"url,omitempty"` // Asynchronous Task URL for further tracking
}
type ResponseWirelessMobilityReset struct {
	Response *ResponseWirelessMobilityResetResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessMobilityResetResponse struct {
	TaskID string `json:"taskId,omitempty"` // Asynchronous Task Id

	URL string `json:"url,omitempty"` // Asynchronous Task URL for further tracking
}
type ResponseWirelessAssignManagedApLocationsForWLC struct {
	Response *ResponseWirelessAssignManagedApLocationsForWLCResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessAssignManagedApLocationsForWLCResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessWirelessControllerProvision struct {
	Response *ResponseWirelessWirelessControllerProvisionResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessWirelessControllerProvisionResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponse struct {
	ManagedApLocations *[]ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetAnchorManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations struct {
	SiteID string `json:"siteId,omitempty"` // The site id of the managed ap location.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceID struct {
	Response *ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // Ap Authorization List Name

	LocalAuthorization *ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponseLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponseRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponseLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // Ap Mac Entries

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // Ap Serial Number Entries
}
type ResponseWirelessGetApAuthorizationListByNetworkDeviceIDResponseRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // Aaa Servers

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // Authorize Ap With Mac

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // Authorize Ap With Serial Number
}
type ResponseWirelessGetManagedApLocationsCountForSpecificWirelessController struct {
	Response *ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetManagedApLocationsCountForSpecificWirelessControllerResponse struct {
	PrimaryManagedApLocationsCount *int `json:"primaryManagedApLocationsCount,omitempty"` // The count of the Primary managed ap locations.

	SecondaryManagedApLocationsCount *int `json:"secondaryManagedApLocationsCount,omitempty"` // The count of the Secondary managed ap locations.

	AnchorManagedApLocationsCount *int `json:"anchorManagedApLocationsCount,omitempty"` // The count of the Anchor managed ap  locations.
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponse struct {
	ManagedApLocations *[]ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetPrimaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations struct {
	SiteID string `json:"siteId,omitempty"` // The site id of the managed ap location.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponse struct {
	ManagedApLocations *[]ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations `json:"managedApLocations,omitempty"` //
}
type ResponseWirelessGetSecondaryManagedApLocationsForSpecificWirelessControllerResponseManagedApLocations struct {
	SiteID string `json:"siteId,omitempty"` // The site id of the managed ap location.

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // The site name hierarchy of the managed ap location.
}
type ResponseWirelessGetSSIDDetailsForSpecificWirelessController struct {
	Response *[]ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetSSIDDetailsForSpecificWirelessControllerResponse struct {
	SSIDName string `json:"ssidName,omitempty"` // Name of the SSID.

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID.

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name.

	L2Security string `json:"l2Security,omitempty"` // This represents the identifier for the Layer 2 authentication type. The authentication types supported include wpa2_enterprise, wpa2_personal, open, wpa3_enterprise, wpa3_personal, wpa2_wpa3_personal, wpa2_wpa3_enterprise, and open-secured.

	L3Security string `json:"l3Security,omitempty"` // This represents the identifier for the Layer 3 authentication type. The authentication types supported are 'open' and 'webauth'.

	RadioPolicy string `json:"radioPolicy,omitempty"` // This represents the identifier for the radio policy. The policies supported include 2.4GHz, 5GHz, and 6GHz.

	AdminStatus *bool `json:"adminStatus,omitempty"` // Utilize this query parameter to obtain the administrative status. A 'true' value signifies that the admin status of the SSID is enabled, while a 'false' value indicates that the admin status of the SSID is disabled.

	Managed *bool `json:"managed,omitempty"` // If the value is 'true,' the SSID is configured through design; if 'false,' it indicates out-of-band configuration on the Wireless LAN Controller.
}
type ResponseWirelessGetSSIDCountForSpecificWirelessController struct {
	Response *ResponseWirelessGetSSIDCountForSpecificWirelessControllerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetSSIDCountForSpecificWirelessControllerResponse struct {
	Count *int `json:"count,omitempty"` // The count of the SSIDs.
}
type ResponseWirelessGetWirelessProfiles struct {
	Response *[]ResponseWirelessGetWirelessProfilesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetWirelessProfilesResponse struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Profile Name

	SSIDDetails *[]ResponseWirelessGetWirelessProfilesResponseSSIDDetails `json:"ssidDetails,omitempty"` //

	ID string `json:"id,omitempty"` // Wireless Profile Id

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // Additional Interfaces

	ApZones *[]ResponseWirelessGetWirelessProfilesResponseApZones `json:"apZones,omitempty"` //
}
type ResponseWirelessGetWirelessProfilesResponseSSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *ResponseWirelessGetWirelessProfilesResponseSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type ResponseWirelessGetWirelessProfilesResponseSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type ResponseWirelessGetWirelessProfilesResponseApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type ResponseWirelessCreateWirelessProfileConnectivity struct {
	Response *ResponseWirelessCreateWirelessProfileConnectivityResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateWirelessProfileConnectivityResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetWirelessProfilesCount struct {
	Response *ResponseWirelessGetWirelessProfilesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetWirelessProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessUpdateWirelessProfileConnectivity struct {
	Response *ResponseWirelessUpdateWirelessProfileConnectivityResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateWirelessProfileConnectivityResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetWirelessProfileByID struct {
	Response *ResponseWirelessGetWirelessProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetWirelessProfileByIDResponse struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Profile Name

	SSIDDetails *[]ResponseWirelessGetWirelessProfileByIDResponseSSIDDetails `json:"ssidDetails,omitempty"` //

	ID string `json:"id,omitempty"` // Wireless Profile Id

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // Additional Interfaces

	ApZones *[]ResponseWirelessGetWirelessProfileByIDResponseApZones `json:"apZones,omitempty"` //
}
type ResponseWirelessGetWirelessProfileByIDResponseSSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *ResponseWirelessGetWirelessProfileByIDResponseSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type ResponseWirelessGetWirelessProfileByIDResponseSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type ResponseWirelessGetWirelessProfileByIDResponseApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type ResponseWirelessDeleteWirelessProfileConnectivity struct {
	Response *ResponseWirelessDeleteWirelessProfileConnectivityResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteWirelessProfileConnectivityResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfile struct {
	Response *[]ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfileResponse struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	PolicyTagName string `json:"policyTagName,omitempty"` // Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space.

	ApZones []string `json:"apZones,omitempty"` // Ap Zones

	PolicyTagID string `json:"policyTagId,omitempty"` // Policy Tag Id
}
type ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk struct {
	Response *ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulkResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfile struct {
	Response *ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfileResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfile struct {
	Response *ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfile struct {
	Response *ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfile struct {
	Response *ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfileResponse struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	PolicyTagName string `json:"policyTagName,omitempty"` // Policy Tag Name

	ApZones []string `json:"apZones,omitempty"` // Ap Zones

	PolicyTagID string `json:"policyTagId,omitempty"` // Policy Tag Id
}
type ResponseWirelessRetrieveAllSiteTagsForAWirelessProfile struct {
	Response *[]ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveAllSiteTagsForAWirelessProfileResponse struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	SiteTagName string `json:"siteTagName,omitempty"` // Site Tag Name

	FlexProfileName string `json:"flexProfileName,omitempty"` // Flex Profile Name

	ApProfileName string `json:"apProfileName,omitempty"` // Ap Profile Name

	SiteTagID string `json:"siteTagId,omitempty"` // Site Tag Id
}
type ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk struct {
	Response *ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulkResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfile struct {
	Response *ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessUpdateASpecificSiteTagForAWirelessProfile struct {
	Response *ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateASpecificSiteTagForAWirelessProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfile struct {
	Response *ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileResponse struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	SiteTagName string `json:"siteTagName,omitempty"` // Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space

	FlexProfileName string `json:"flexProfileName,omitempty"` // Flex Profile Name

	ApProfileName string `json:"apProfileName,omitempty"` // Ap Profile Name

	SiteTagID string `json:"siteTagId,omitempty"` // Site Tag Id
}
type ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfile struct {
	Response *ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessCreateAnchorGroup struct {
	Response *ResponseWirelessCreateAnchorGroupResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateAnchorGroupResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetAnchorGroups struct {
	ID string `json:"id,omitempty"` // Anchor Profile unique ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]ResponseWirelessGetAnchorGroupsMobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type ResponseWirelessGetAnchorGroupsMobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address
}
type ResponseWirelessGetCountOfAnchorGroups struct {
	Response *ResponseWirelessGetCountOfAnchorGroupsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetCountOfAnchorGroupsResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetAnchorGroupByID struct {
	ID string `json:"id,omitempty"` // Anchor Profile unique ID

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]ResponseWirelessGetAnchorGroupByIDMobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type ResponseWirelessGetAnchorGroupByIDMobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address
}
type ResponseWirelessDeleteAnchorGroupByID struct {
	Response *ResponseWirelessDeleteAnchorGroupByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteAnchorGroupByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateAnchorGroup struct {
	Response *ResponseWirelessUpdateAnchorGroupResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateAnchorGroupResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApAuthorizationLists struct {
	Response *ResponseWirelessGetApAuthorizationListsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApAuthorizationListsResponse struct {
	ID string `json:"id,omitempty"` // Id

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // Ap Authorization List Name

	LocalAuthorization *ResponseWirelessGetApAuthorizationListsResponseLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *ResponseWirelessGetApAuthorizationListsResponseRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type ResponseWirelessGetApAuthorizationListsResponseLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // AP Mac Addresses

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // AP Serial Number Entries
}
type ResponseWirelessGetApAuthorizationListsResponseRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // AAA Servers

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // Authorize AP With Mac

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // Authorize AP With Serial Number
}
type ResponseWirelessCreateApAuthorizationList struct {
	Response *ResponseWirelessCreateApAuthorizationListResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateApAuthorizationListResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApAuthorizationListCount struct {
	Response *ResponseWirelessGetApAuthorizationListCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetApAuthorizationListCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteApAuthorizationList struct {
	Response *ResponseWirelessDeleteApAuthorizationListResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteApAuthorizationListResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateApAuthorizationList struct {
	Response *ResponseWirelessUpdateApAuthorizationListResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateApAuthorizationListResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApAuthorizationListByID struct {
	Response *ResponseWirelessGetApAuthorizationListByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApAuthorizationListByIDResponse struct {
	ID string `json:"id,omitempty"` // Id

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // Ap Authorization List Name

	LocalAuthorization *ResponseWirelessGetApAuthorizationListByIDResponseLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *ResponseWirelessGetApAuthorizationListByIDResponseRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type ResponseWirelessGetApAuthorizationListByIDResponseLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // AP Mac Addresses

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // AP Serial Number Entries
}
type ResponseWirelessGetApAuthorizationListByIDResponseRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // AAA Servers

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // Authorize AP With Mac

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // Authorize AP With Serial Number
}
type ResponseWirelessCreateApProfile struct {
	Response *ResponseWirelessCreateApProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateApProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApProfiles struct {
	Response *[]ResponseWirelessGetApProfilesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApProfilesResponse struct {
	ID string `json:"id,omitempty"` // AP Profile unique ID

	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *ResponseWirelessGetApProfilesResponseManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *ResponseWirelessGetApProfilesResponseRogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *ResponseWirelessGetApProfilesResponseMeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *ResponseWirelessGetApProfilesResponseCalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // Time zone of the AP.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Hour 'Delta from Controller' for the time zone. The value should be between -12 and 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Minute 'Delta from Controller' for the time zone. Value should be between 0 to 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type ResponseWirelessGetApProfilesResponseManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type ResponseWirelessGetApProfilesResponseRogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type ResponseWirelessGetApProfilesResponseMeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used.
}
type ResponseWirelessGetApProfilesResponseCalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. The following API is used create AP power profile. API-/intent/api/v1/wirelessSettings/powerProfiles

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *ResponseWirelessGetApProfilesResponseCalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type ResponseWirelessGetApProfilesResponseCalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type ResponseWirelessGetApProfilesCount struct {
	Response *ResponseWirelessGetApProfilesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetApProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteApProfileByID struct {
	Response *ResponseWirelessDeleteApProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteApProfileByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateApProfileByID struct {
	Response *ResponseWirelessUpdateApProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateApProfileByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetApProfileByID struct {
	Response *[]ResponseWirelessGetApProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetApProfileByIDResponse struct {
	ID string `json:"id,omitempty"` // AP Profile unique ID

	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *ResponseWirelessGetApProfileByIDResponseManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *ResponseWirelessGetApProfileByIDResponseRogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *ResponseWirelessGetApProfileByIDResponseMeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *ResponseWirelessGetApProfileByIDResponseCalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Enter the hour value (HH). The valid range is from -12 through 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Enter the minute value (MM). The valid range is from 0 through 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type ResponseWirelessGetApProfileByIDResponseManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type ResponseWirelessGetApProfileByIDResponseRogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type ResponseWirelessGetApProfileByIDResponseMeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used.
}
type ResponseWirelessGetApProfileByIDResponseCalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *ResponseWirelessGetApProfileByIDResponseCalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type ResponseWirelessGetApProfileByIDResponseCalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type ResponseWirelessGet80211BeProfiles struct {
	Response *[]ResponseWirelessGet80211BeProfilesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGet80211BeProfilesResponse struct {
	ID string `json:"id,omitempty"` // 802.11be Profile ID

	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU

	Default *bool `json:"default,omitempty"` // 802.11be Profile is marked default or custom (Read only field)
}
type ResponseWirelessCreateA80211BeProfile struct {
	Response *ResponseWirelessCreateA80211BeProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateA80211BeProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGet80211BeProfilesCount struct {
	Response *ResponseWirelessGet80211BeProfilesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGet80211BeProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteA80211BeProfile struct {
	Response *ResponseWirelessDeleteA80211BeProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteA80211BeProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdate80211BeProfile struct {
	Response *ResponseWirelessUpdate80211BeProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdate80211BeProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGet80211BeProfileByID struct {
	Response *ResponseWirelessGet80211BeProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGet80211BeProfileByIDResponse struct {
	ID string `json:"id,omitempty"` // 802.11be Profile ID

	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU

	Default *bool `json:"default,omitempty"` // Is 802.11be Profile marked as default in System . (Read only field)
}
type ResponseWirelessGetInterfaces struct {
	Response *[]ResponseWirelessGetInterfacesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetInterfacesResponse struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID

	ID string `json:"id,omitempty"` // Interface ID
}
type ResponseWirelessCreateInterface struct {
	Response *ResponseWirelessCreateInterfaceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateInterfaceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetInterfacesCount struct {
	Response *ResponseWirelessGetInterfacesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetInterfacesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessGetInterfaceByID struct {
	Response *ResponseWirelessGetInterfaceByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetInterfaceByIDResponse struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID

	ID string `json:"id,omitempty"` // Interface ID
}
type ResponseWirelessDeleteInterface struct {
	Response *ResponseWirelessDeleteInterfaceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteInterfaceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdateInterface struct {
	Response *ResponseWirelessUpdateInterfaceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateInterfaceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessCreatePowerProfile struct {
	Response *ResponseWirelessCreatePowerProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreatePowerProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetPowerProfiles struct {
	Response *[]ResponseWirelessGetPowerProfilesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetPowerProfilesResponse struct {
	ID string `json:"id,omitempty"` // Unique Identifier of the power profile.

	ProfileName string `json:"profileName,omitempty"` // The Name of the Power Profile.

	Description string `json:"description,omitempty"` // The description of the Power Profile.

	Rules *[]ResponseWirelessGetPowerProfilesResponseRules `json:"rules,omitempty"` //
}
type ResponseWirelessGetPowerProfilesResponseRules struct {
	Sequence *int `json:"sequence,omitempty"` // The sequence of the power profile rule.

	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type for the rule.

	InterfaceID string `json:"interfaceId,omitempty"` // Interface Id for the rule.

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type for the rule.

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value for the rule.
}
type ResponseWirelessGetPowerProfilesCount struct {
	Response *ResponseWirelessGetPowerProfilesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetPowerProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeletePowerProfileByID struct {
	Response *ResponseWirelessDeletePowerProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeletePowerProfileByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessUpdatePowerProfileByID struct {
	Response *ResponseWirelessUpdatePowerProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdatePowerProfileByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetPowerProfileByID struct {
	Response *ResponseWirelessGetPowerProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetPowerProfileByIDResponse struct {
	ID string `json:"id,omitempty"` // Unique Identifier of the power profile.

	ProfileName string `json:"profileName,omitempty"` // The Name of the Power Profile.

	Description string `json:"description,omitempty"` // The description of the Power Profile.

	Rules *[]ResponseWirelessGetPowerProfileByIDResponseRules `json:"rules,omitempty"` //
}
type ResponseWirelessGetPowerProfileByIDResponseRules struct {
	Sequence *int `json:"sequence,omitempty"` // Sequential Ordered List of rules for Power Profile.

	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type for the rule.

	InterfaceID string `json:"interfaceId,omitempty"` // Interface Id for the rule.

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type for the rule.

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value for the rule.
}
type ResponseWirelessCreateRfProfile struct {
	Response *ResponseWirelessCreateRfProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessCreateRfProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetRfProfiles struct {
	Response *[]ResponseWirelessGetRfProfilesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetRfProfilesResponse struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	EnableCustom *bool `json:"enableCustom,omitempty"` // True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)

	RadioTypeAProperties *ResponseWirelessGetRfProfilesResponseRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *ResponseWirelessGetRfProfilesResponseRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *ResponseWirelessGetRfProfilesResponseRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //

	ID string `json:"id,omitempty"` // RF Profile ID
}
type ResponseWirelessGetRfProfilesResponseRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *ResponseWirelessGetRfProfilesResponseRadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfilesResponseRadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfilesResponseRadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfilesResponseRadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type ResponseWirelessGetRfProfilesResponseRadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfilesResponseRadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfilesResponseRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfilesResponseRadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfilesResponseRadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfilesResponseRadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfilesResponseRadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width ( Permissible values : 20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values: 20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfilesResponseRadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfilesCount struct {
	Response *ResponseWirelessGetRfProfilesCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response Version
}
type ResponseWirelessGetRfProfilesCountResponse struct {
	Count *int `json:"count,omitempty"` // Count of the requested resource
}
type ResponseWirelessDeleteRfProfile struct {
	Response *ResponseWirelessDeleteRfProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessDeleteRfProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessGetRfProfileByID struct {
	Response *ResponseWirelessGetRfProfileByIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessGetRfProfileByIDResponse struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	EnableCustom *bool `json:"enableCustom,omitempty"` // True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)

	RadioTypeAProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //

	ID string `json:"id,omitempty"` // RF Profile ID
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfileByIDResponseRadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfileByIDResponseRadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width ( Permissible values : 20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values: 20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type ResponseWirelessGetRfProfileByIDResponseRadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type ResponseWirelessUpdateRfProfile struct {
	Response *ResponseWirelessUpdateRfProfileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseWirelessUpdateRfProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID

	URL string `json:"url,omitempty"` // Task URL
}
type ResponseWirelessRetrieveSitesWithOverriddenSSIDs struct {
	Response *[]ResponseWirelessRetrieveSitesWithOverriddenSSIDsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseWirelessRetrieveSitesWithOverriddenSSIDsResponse struct {
	SiteID string `json:"siteId,omitempty"` // Site ID

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy

	SSIDs *[]ResponseWirelessRetrieveSitesWithOverriddenSSIDsResponseSSIDs `json:"ssids,omitempty"` //
}
type ResponseWirelessRetrieveSitesWithOverriddenSSIDsResponseSSIDs struct {
	ID string `json:"id,omitempty"` // SSID ID

	SSID string `json:"ssid,omitempty"` // SSID
}
type ResponseWirelessAssignAnchorManagedApLocationsForWLC struct {
	Response *ResponseWirelessAssignAnchorManagedApLocationsForWLCResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseWirelessAssignAnchorManagedApLocationsForWLCResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task

	URL string `json:"url,omitempty"` // URL for the task
}
type ResponseWirelessConfigureAccessPointsV2 struct {
	Response *ResponseWirelessConfigureAccessPointsV2Response `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseWirelessConfigureAccessPointsV2Response struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type RequestWirelessCreateAndProvisionSSID struct {
	ManagedApLocations []string `json:"managedAPLocations,omitempty"` // Managed AP Locations (Enter entire Site(s) hierarchy)

	SSIDDetails *RequestWirelessCreateAndProvisionSSIDSSIDDetails `json:"ssidDetails,omitempty"` //

	SSIDType string `json:"ssidType,omitempty"` // SSID Type

	EnableFabric *bool `json:"enableFabric,omitempty"` // Enable SSID for Fabric

	FlexConnect *RequestWirelessCreateAndProvisionSSIDFlexConnect `json:"flexConnect,omitempty"` //
}
type RequestWirelessCreateAndProvisionSSIDSSIDDetails struct {
	Name string `json:"name,omitempty"` // SSID Name

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level(For guest SSID OPEN/WEB_AUTH, For Enterprise SSID ENTERPRISE/PERSONAL/OPEN)

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable Fast Lane

	Passphrase string `json:"passphrase,omitempty"` // Pass Phrase ( Only applicable for SSID with PERSONAL auth type )

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcast SSID

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	WebAuthURL string `json:"webAuthURL,omitempty"` // Web Auth URL

	AuthKeyMgmt []string `json:"authKeyMgmt,omitempty"` // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // Rsn Cipher Suite Gcmp256

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // Rsn Cipher Suite  Gcmp128

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // Rsn Cipher Suite Ccmp256

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // 6 Ghz Client Steering

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 GHz Policy
}
type RequestWirelessCreateAndProvisionSSIDFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // Enable Flex Connect

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local To Vlan (range is 1 to 4094)
}
type RequestWirelessRebootAccessPoints struct {
	ApMacAddresses []string `json:"apMacAddresses,omitempty"` // The ethernet MAC address of the access point.
}
type RequestWirelessCreateEnterpriseSSID struct {
	Name string `json:"name,omitempty"` // SSID NAME

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level

	Passphrase string `json:"passphrase,omitempty"` // Passphrase

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable FastLane

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type Enum (voicedata or data )

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy Enum

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcase SSID

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	EnableSessionTimeOut *bool `json:"enableSessionTimeOut,omitempty"` // Enable Session Timeout

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // Session Time Out

	EnableClientExclusion *bool `json:"enableClientExclusion,omitempty"` // Enable Client Exclusion

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // Client Exclusion Timeout

	EnableBasicServiceSetMaxIDle *bool `json:"enableBasicServiceSetMaxIdle,omitempty"` // Enable Basic Service Set Max Idle

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout

	EnableDirectedMulticastService *bool `json:"enableDirectedMulticastService,omitempty"` // Enable Directed Multicast Service

	EnableNeighborList *bool `json:"enableNeighborList,omitempty"` // Enable Neighbor List

	MfpClientProtection string `json:"mfpClientProtection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Nas Options

	ProfileName string `json:"profileName,omitempty"` // Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Aaa Override

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessCreateEnterpriseSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *float64 `json:"clientRateLimit,omitempty"` // Client Rate Limit (in bits per second)

	AuthKeyMgmt []string `json:"authKeyMgmt,omitempty"` // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // Rsn Cipher Suite Gcmp256

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // Rsn Cipher Suite Ccmp256

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // Rsn Cipher Suite Gcmp 128

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // Ghz6 Policy Client Steering

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // Ghz24 Policy
}
type RequestWirelessCreateEnterpriseSSIDMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}
type RequestWirelessUpdateEnterpriseSSID struct {
	Name string `json:"name,omitempty"` // SSID NAME

	SecurityLevel string `json:"securityLevel,omitempty"` // Security Level

	Passphrase string `json:"passphrase,omitempty"` // Passphrase

	EnableFastLane *bool `json:"enableFastLane,omitempty"` // Enable FastLane

	EnableMacFiltering *bool `json:"enableMACFiltering,omitempty"` // Enable MAC Filtering

	TrafficType string `json:"trafficType,omitempty"` // Traffic Type Enum (voicedata or data )

	RadioPolicy string `json:"radioPolicy,omitempty"` // Radio Policy Enum

	EnableBroadcastSSID *bool `json:"enableBroadcastSSID,omitempty"` // Enable Broadcase SSID

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	EnableSessionTimeOut *bool `json:"enableSessionTimeOut,omitempty"` // Enable Session Timeout

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // Session Time Out

	EnableClientExclusion *bool `json:"enableClientExclusion,omitempty"` // Enable Client Exclusion

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // Client Exclusion Timeout

	EnableBasicServiceSetMaxIDle *bool `json:"enableBasicServiceSetMaxIdle,omitempty"` // Enable Basic Service Set Max Idle

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout

	EnableDirectedMulticastService *bool `json:"enableDirectedMulticastService,omitempty"` // Enable Directed Multicast Service

	EnableNeighborList *bool `json:"enableNeighborList,omitempty"` // Enable Neighbor List

	MfpClientProtection string `json:"mfpClientProtection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Nas Options

	ProfileName string `json:"profileName,omitempty"` // Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Aaa Override

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (Required applicable for Security Type WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (Optional, Required Applicable for Security Type WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessUpdateEnterpriseSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *float64 `json:"clientRateLimit,omitempty"` // Client Rate Limit (in bits per second)

	AuthKeyMgmt []string `json:"authKeyMgmt,omitempty"` // Takes string inputs for the AKMs that should be set true. Possible AKM values : dot1x,dot1x_ft, dot1x_sha, psk, psk_ft, psk_sha, owe, sae, sae_ft

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // Rsn Cipher Suite Gcmp256

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // Rsn Cipher Suite Ccmp256

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // Rsn Cipher Suite Gcmp 128

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // Ghz6 Policy Client Steering

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // Ghz24 Policy
}
type RequestWirelessUpdateEnterpriseSSIDMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type

	Passphrase string `json:"passphrase,omitempty"` // Passphrase
}
type RequestWirelessCreateSSID struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled). Default is L2 Authentication Type if exists else .

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts. Default is Client Exclusion Timeout if exists else 180.

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out. Default is Basic ServiceSet ClientIdle Timeout if exists else 300.

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Default is Management Frame Protection Client if exists else Optional.

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessCreateSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve. It should be in mutliples of 500 . Default is Client Rate Limit if exists else 0.

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band else false.

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // Default is L3 Authentication Type if exists else .

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network. Default is Sleeping Client Timeout if exists else 720.

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // he default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type RequestWirelessCreateSSIDMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type(default: ASCII)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessUpdateSSID struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled). Default is L2 Authentication Type if exists else .

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts. Default is Client Exclusion Timeout if exists else 180.

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out. Default is Basic ServiceSet ClientIdle Timeout if exists else 300.

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Default is Management Frame Protection Client if exists else Optional.

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessUpdateSSIDMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve. It should be in mutliples of 500 . Default is Client Rate Limit if exists else 0.

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band else false.

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // Default is L3 Authentication Type if exists else .

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network. Default is Sleeping Client Timeout if exists else 720.

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // he default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true
}
type RequestWirelessUpdateSSIDMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type(default: ASCII)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessUpdateOrOverridessid struct {
	SSID string `json:"ssid,omitempty"` // Name of the SSID

	AuthType string `json:"authType,omitempty"` // L2 Authentication Type (If authType is not open , then atleast one RSN Cipher Suite and corresponding valid AKM must be enabled)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase (Only applicable for SSID with PERSONAL security level). Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters

	IsFastLaneEnabled *bool `json:"isFastLaneEnabled,omitempty"` // True if FastLane is enabled, else False

	IsMacFilteringEnabled *bool `json:"isMacFilteringEnabled,omitempty"` // When set to true, MAC Filtering will be activated, allowing control over network access based on the MAC address of the device

	SSIDRadioType string `json:"ssidRadioType,omitempty"` // Radio Policy Enum (default: Triple band operation(2.4GHz, 5GHz and 6GHz))

	IsBroadcastSSID *bool `json:"isBroadcastSSID,omitempty"` // When activated by setting it to true, the Broadcast SSID feature will make the SSID publicly visible to wireless devices searching for available networks

	FastTransition string `json:"fastTransition,omitempty"` // Fast Transition

	SessionTimeOutEnable *bool `json:"sessionTimeOutEnable,omitempty"` // Turn on the feature that imposes a time limit on user sessions

	SessionTimeOut *int `json:"sessionTimeOut,omitempty"` // This denotes the allotted time span, expressed in seconds, before a session is automatically terminated due to inactivity. Default sessionTimeOut is 1800.

	ClientExclusionEnable *bool `json:"clientExclusionEnable,omitempty"` // Activate the feature that allows for the exclusion of clients

	ClientExclusionTimeout *int `json:"clientExclusionTimeout,omitempty"` // This refers to the length of time, in seconds, a client is excluded or blocked from accessing the network after a specified number of unsuccessful attempts

	BasicServiceSetMaxIDleEnable *bool `json:"basicServiceSetMaxIdleEnable,omitempty"` // Activate the maximum idle feature for the Basic Service Set

	BasicServiceSetClientIDleTimeout *int `json:"basicServiceSetClientIdleTimeout,omitempty"` // This refers to the duration of inactivity, measured in seconds, before a client connected to the Basic Service Set is considered idle and timed out

	DirectedMulticastServiceEnable *bool `json:"directedMulticastServiceEnable,omitempty"` // The Directed Multicast Service feature becomes operational when it is set to true

	NeighborListEnable *bool `json:"neighborListEnable,omitempty"` // The Neighbor List feature is enabled when it is set to true

	ManagementFrameProtectionClientprotection string `json:"managementFrameProtectionClientprotection,omitempty"` // Management Frame Protection Client

	NasOptions []string `json:"nasOptions,omitempty"` // Pre-Defined NAS Options : AP ETH Mac Address, AP IP address, AP Location , AP MAC Address, AP Name, AP Policy Tag, AP Site Tag, SSID, System IP Address, System MAC Address, System Name.

	ProfileName string `json:"profileName,omitempty"` // WLAN Profile Name, if not passed autogenerated profile name will be assigned. The same wlanProfileName will also be used for policyProfileName

	AAAOverride *bool `json:"aaaOverride,omitempty"` // Activate the AAA Override feature when set to true

	CoverageHoleDetectionEnable *bool `json:"coverageHoleDetectionEnable,omitempty"` // Activate Coverage Hole Detection feature when set to true

	ProtectedManagementFrame string `json:"protectedManagementFrame,omitempty"` // (REQUIRED is applicable for authType WPA3_PERSONAL, WPA3_ENTERPRISE, OPEN_SECURED) and (OPTIONAL/REQUIRED is applicable for authType WPA2_WPA3_PERSONAL and WPA2_WPA3_ENTERPRISE)

	MultipSKSettings *[]RequestWirelessUpdateOrOverridessidMultipSKSettings `json:"multiPSKSettings,omitempty"` //

	ClientRateLimit *int `json:"clientRateLimit,omitempty"` // This pertains to the maximum data transfer rate, specified in bits per second, that a client is permitted to achieve

	RsnCipherSuiteGcmp256 *bool `json:"rsnCipherSuiteGcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP256 encryption protocol is activated

	RsnCipherSuiteCcmp256 *bool `json:"rsnCipherSuiteCcmp256,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP256 encryption protocol is activated

	RsnCipherSuiteGcmp128 *bool `json:"rsnCipherSuiteGcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite GCMP128 encryption protocol is activated

	RsnCipherSuiteCcmp128 *bool `json:"rsnCipherSuiteCcmp128,omitempty"` // When set to true, the Robust Security Network (RSN) Cipher Suite CCMP128 encryption protocol is activated

	Ghz6PolicyClientSteering *bool `json:"ghz6PolicyClientSteering,omitempty"` // True if 6 GHz Policy Client Steering is enabled, else False

	IsAuthKey8021X *bool `json:"isAuthKey8021x,omitempty"` // When set to true, the 802.1X authentication key is in use

	IsAuthKey8021XPlusFT *bool `json:"isAuthKey8021xPlusFT,omitempty"` // When set to true, the 802.1X-Plus-FT authentication key is in use

	IsAuthKey8021XSHA256 *bool `json:"isAuthKey8021x_SHA256,omitempty"` // When set to true, the feature that enables 802.1X authentication using the SHA256 algorithm is turned on

	IsAuthKeySae *bool `json:"isAuthKeySae,omitempty"` // When set to true, the feature enabling the Simultaneous Authentication of Equals (SAE) authentication key is activated

	IsAuthKeySaePlusFT *bool `json:"isAuthKeySaePlusFT,omitempty"` // Activating this setting by switching it to true turns on the authentication key feature that supports both Simultaneous Authentication of Equals (SAE) and Fast Transition (FT)

	IsAuthKeyPSK *bool `json:"isAuthKeyPSK,omitempty"` // When set to true, the Pre-shared Key (PSK) authentication feature is enabled

	IsAuthKeyPSKPlusFT *bool `json:"isAuthKeyPSKPlusFT,omitempty"` // When set to true, the feature that enables the combination of Pre-shared Key (PSK) and Fast Transition (FT) authentication keys is activated

	IsAuthKeyOWE *bool `json:"isAuthKeyOWE,omitempty"` // When set to true, the Opportunistic Wireless Encryption (OWE) authentication key feature is turned on

	IsAuthKeyEasyPSK *bool `json:"isAuthKeyEasyPSK,omitempty"` // When set to true, the feature that enables the use of Easy Pre-shared Key (PSK) authentication is activated

	IsAuthKeyPSKSHA256 *bool `json:"isAuthKeyPSKSHA256,omitempty"` // The feature that allows the use of Pre-shared Key (PSK) authentication with the SHA256 algorithm is enabled when it is set to true

	OpenSSID string `json:"openSsid,omitempty"` // Open SSID which is already created in the design and not associated to any other OPEN-SECURED SSID

	WLANBandSelectEnable *bool `json:"wlanBandSelectEnable,omitempty"` // Band select is allowed only when band options selected contains at least 2.4 GHz and 5 GHz band

	IsEnabled *bool `json:"isEnabled,omitempty"` // Set SSID's admin status as 'Enabled' when set to true

	AuthServers []string `json:"authServers,omitempty"` // List of Authentication/Authorization server IpAddresses

	AcctServers []string `json:"acctServers,omitempty"` // List of Accounting server IpAddresses

	EgressQos string `json:"egressQos,omitempty"` // Egress QOS

	IngressQos string `json:"ingressQos,omitempty"` // Ingress QOS

	WLANType string `json:"wlanType,omitempty"` // Wlan Type

	L3AuthType string `json:"l3AuthType,omitempty"` // L3 Authentication Type

	AuthServer string `json:"authServer,omitempty"` // Authentication Server, Mandatory for Guest SSIDs with wlanType=Guest and l3AuthType=web_auth

	ExternalAuthIPAddress string `json:"externalAuthIpAddress,omitempty"` // External WebAuth URL (Mandatory for Guest SSIDs with wlanType = Guest, l3AuthType = web_auth and authServer = auth_external)

	WebPassthrough *bool `json:"webPassthrough,omitempty"` // When set to true, the Web-Passthrough feature will be activated for the Guest SSID, allowing guests to bypass certain login requirements

	SleepingClientEnable *bool `json:"sleepingClientEnable,omitempty"` // When set to true, this will activate the timeout settings that apply to clients in sleep mode

	SleepingClientTimeout *int `json:"sleepingClientTimeout,omitempty"` // This refers to the amount of time, measured in minutes, before a sleeping (inactive) client is timed out of the network

	ACLName string `json:"aclName,omitempty"` // Pre-Auth Access Control List (ACL) Name

	IsPosturingEnabled *bool `json:"isPosturingEnabled,omitempty"` // Applicable only for Enterprise SSIDs. When set to True, Posturing will enabled. Required to be set to True if ACL needs to be mapped for Enterprise SSID.

	IsAuthKeySuiteB1X *bool `json:"isAuthKeySuiteB1x,omitempty"` // When activated by setting it to true, the SuiteB-1x authentication key feature is engaged.

	IsAuthKeySuiteB1921X *bool `json:"isAuthKeySuiteB1921x,omitempty"` // When set to true, the SuiteB192-1x authentication key feature is enabled.

	IsAuthKeySaeExt *bool `json:"isAuthKeySaeExt,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) Extended Authentication key feature is turned on.

	IsAuthKeySaeExtPlusFT *bool `json:"isAuthKeySaeExtPlusFT,omitempty"` // When set to true, the Simultaneous Authentication of Equals (SAE) combined with Fast Transition (FT) Authentication Key feature is enabled.

	IsApBeaconProtectionEnabled *bool `json:"isApBeaconProtectionEnabled,omitempty"` // When set to true, the Access Point (AP) Beacon Protection feature is activated, enhancing the security of the network.

	Ghz24Policy string `json:"ghz24Policy,omitempty"` // 2.4 Ghz Band Policy value. Allowed only when 2.4 Radio Band is enabled in ssidRadioType

	CckmTsfTolerance *int `json:"cckmTsfTolerance,omitempty"` // The default value is the Cckm Timestamp Tolerance (in milliseconds, if specified); otherwise, it is 0.

	IsCckmEnabled *bool `json:"isCckmEnabled,omitempty"` // True if CCKM is enabled, else False

	IsHex *bool `json:"isHex,omitempty"` // True if passphrase is in Hex format, else False.

	IsRandomMacFilterEnabled *bool `json:"isRandomMacFilterEnabled,omitempty"` // Deny clients using randomized MAC addresses when set to true

	FastTransitionOverTheDistributedSystemEnable *bool `json:"fastTransitionOverTheDistributedSystemEnable,omitempty"` // Enable Fast Transition over the Distributed System when set to true

	IsRadiusProfilingEnabled *bool `json:"isRadiusProfilingEnabled,omitempty"` // Enable Radius Profiling. At least one AAA/PSN server is required to enable Radius Profiling on WLAN.
}
type RequestWirelessUpdateOrOverridessidMultipSKSettings struct {
	Priority *int `json:"priority,omitempty"` // Priority

	PassphraseType string `json:"passphraseType,omitempty"` // Passphrase Type(default:ASCII)

	Passphrase string `json:"passphrase,omitempty"` // Passphrase needs to be between 8 and 63 characters for ASCII type. HEX passphrase needs to be 64 characters
}
type RequestWirelessConfigureAccessPoints struct {
	ApList *[]RequestWirelessConfigureAccessPointsApList `json:"apList,omitempty"` //

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the access point's admin status, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureApMode *bool `json:"configureApMode,omitempty"` // To change the access point's mode, set this parameter's value to "true".

	ApMode *int `json:"apMode,omitempty"` // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".

	ConfigureFailoverPriority *bool `json:"configureFailoverPriority,omitempty"` // To change the access point's failover priority, set this parameter's value to "true".

	FailoverPriority *int `json:"failoverPriority,omitempty"` // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".

	ConfigureLedStatus *bool `json:"configureLedStatus,omitempty"` // To change the access point's LED status, set this parameter's value to "true".

	LedStatus *bool `json:"ledStatus,omitempty"` // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.

	ConfigureLedBrightnessLevel *bool `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".

	LedBrightnessLevel *int `json:"ledBrightnessLevel,omitempty"` // Configure the access point's LED brightness level by setting a value between 1 and 8.

	ConfigureLocation *bool `json:"configureLocation,omitempty"` // To change the access point's location, set this parameter's value to "true".

	Location string `json:"location,omitempty"` // Configure the access point's location.

	ConfigureHAController *bool `json:"configureHAController,omitempty"` // To change the access point's HA controller, set this parameter's value to "true".

	PrimaryControllerName string `json:"primaryControllerName,omitempty"` // Configure the hostname for an access point's primary controller.

	PrimaryIPAddress *RequestWirelessConfigureAccessPointsPrimaryIPAddress `json:"primaryIpAddress,omitempty"` //

	SecondaryControllerName string `json:"secondaryControllerName,omitempty"` // Configure the hostname for an access point's secondary controller.

	SecondaryIPAddress *RequestWirelessConfigureAccessPointsSecondaryIPAddress `json:"secondaryIpAddress,omitempty"` //

	TertiaryControllerName string `json:"tertiaryControllerName,omitempty"` // Configure the hostname for an access point's tertiary controller.

	TertiaryIPAddress *RequestWirelessConfigureAccessPointsTertiaryIPAddress `json:"tertiaryIpAddress,omitempty"` //

	RadioConfigurations *[]RequestWirelessConfigureAccessPointsRadioConfigurations `json:"radioConfigurations,omitempty"` //

	IsAssignedSiteAsLocation *bool `json:"isAssignedSiteAsLocation,omitempty"` // If AP is assigned to a site, then to assign AP location as the site name, set this parameter's value to "true".
}
type RequestWirelessConfigureAccessPointsApList struct {
	ApName string `json:"apName,omitempty"` // The current host name of the access point.

	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.

	ApNameNew string `json:"apNameNew,omitempty"` // The modified hostname of the access point.
}
type RequestWirelessConfigureAccessPointsPrimaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's primary controller.
}
type RequestWirelessConfigureAccessPointsSecondaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's secondary controller.
}
type RequestWirelessConfigureAccessPointsTertiaryIPAddress struct {
	Address string `json:"address,omitempty"` // Configure the IP address for an access point's tertiary controller.
}
type RequestWirelessConfigureAccessPointsRadioConfigurations struct {
	ConfigureRadioRoleAssignment *bool `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".

	RadioRoleAssignment string `json:"radioRoleAssignment,omitempty"` // Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string

	RadioBand string `json:"radioBand,omitempty"` // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the admin status on the specified radio for an access point, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureAntennaPatternName *bool `json:"configureAntennaPatternName,omitempty"` // To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".

	AntennaPatternName string `json:"antennaPatternName,omitempty"` // Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.

	AntennaGain *int `json:"antennaGain,omitempty"` // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".

	ConfigureAntennaCable *bool `json:"configureAntennaCable,omitempty"` // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".

	AntennaCableName string `json:"antennaCableName,omitempty"` // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".

	CableLoss *float64 `json:"cableLoss,omitempty"` // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).

	ConfigureChannel *bool `json:"configureChannel,omitempty"` // To change the channel on the specified radio for an access point, set this parameter's value to "true".

	ChannelAssignmentMode *int `json:"channelAssignmentMode,omitempty"` // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	ChannelNumber *int `json:"channelNumber,omitempty"` // Configure the channel number on the specified radio for an access point.

	ConfigureChannelWidth *bool `json:"configureChannelWidth,omitempty"` // To change the channel width on the specified radio for an access point, set this parameter's value to "true".

	ChannelWidth *int `json:"channelWidth,omitempty"` // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".

	ConfigurePower *bool `json:"configurePower,omitempty"` // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".

	PowerAssignmentMode *int `json:"powerAssignmentMode,omitempty"` // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	Powerlevel *int `json:"powerlevel,omitempty"` // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.

	ConfigureCleanAirSI *bool `json:"configureCleanAirSI,omitempty"` // To enable or disable either CleanAir or Spectrum Intelligence on the specified radio for an access point, set this parameter's value to "true".

	CleanAirSI *int `json:"cleanAirSI,omitempty"` // Configure CleanAir or Spectrum Intelligence on the specified radio for an access point. Set this parameter's value to "0" to disable the feature or "1" to enable it.

	RadioType *int `json:"radioType,omitempty"` // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
}
type RequestWirelessApProvisionConnectivity []RequestItemWirelessApProvisionConnectivity // Array of RequestWirelessAPProvisionConnectivity
type RequestItemWirelessApProvisionConnectivity struct {
	RfProfile string `json:"rfProfile,omitempty"` // Radio frequency profile name

	DeviceName string `json:"deviceName,omitempty"` // Device name

	CustomApGroupName string `json:"customApGroupName,omitempty"` // Custom AP group name

	CustomFlexGroupName []string `json:"customFlexGroupName,omitempty"` // ["Custom flex group name"]

	Type string `json:"type,omitempty"` // ApWirelessConfiguration

	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site name hierarchy(ex: Global/...)
}
type RequestWirelessCreateUpdateDynamicInterface struct {
	InterfaceName string `json:"interfaceName,omitempty"` // dynamic-interface name

	VLANID *float64 `json:"vlanId,omitempty"` // Vlan Id
}
type RequestWirelessUpdateWirelessProfile struct {
	ProfileDetails *RequestWirelessUpdateWirelessProfileProfileDetails `json:"profileDetails,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileProfileDetails struct {
	Name string `json:"name,omitempty"` // Profile Name

	Sites []string `json:"sites,omitempty"` // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])

	SSIDDetails *[]RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetails struct {
	Name string `json:"name,omitempty"` // Ssid Name

	EnableFabric *bool `json:"enableFabric,omitempty"` // true if ssid is fabric else false

	FlexConnect *RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name
}
type RequestWirelessUpdateWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID. Required if enableFlexConnect is true.
}
type RequestWirelessCreateWirelessProfile struct {
	ProfileDetails *RequestWirelessCreateWirelessProfileProfileDetails `json:"profileDetails,omitempty"` //
}
type RequestWirelessCreateWirelessProfileProfileDetails struct {
	Name string `json:"name,omitempty"` // Profile Name

	Sites []string `json:"sites,omitempty"` // array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])

	SSIDDetails *[]RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetails struct {
	Name string `json:"name,omitempty"` // Ssid Name

	EnableFabric *bool `json:"enableFabric,omitempty"` // true if ssid is fabric else false

	FlexConnect *RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	PolicyProfileName string `json:"policyProfileName,omitempty"` // Policy Profile Name
}
type RequestWirelessCreateWirelessProfileProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // true if flex connect is enabled else false

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID. Required if enableFlexConnect is true.
}
type RequestWirelessProvisionUpdate []RequestItemWirelessProvisionUpdate // Array of RequestWirelessProvisionUpdate
type RequestItemWirelessProvisionUpdate struct {
	DeviceName string `json:"deviceName,omitempty"` // Controller Name

	ManagedApLocations []string `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)

	DynamicInterfaces *[]RequestItemWirelessProvisionUpdateDynamicInterfaces `json:"dynamicInterfaces,omitempty"` //
}
type RequestItemWirelessProvisionUpdateDynamicInterfaces struct {
	InterfaceIPAddress string `json:"interfaceIPAddress,omitempty"` // Interface IP Address. Required for AireOS.

	InterfaceNetmaskInCIDR *int `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR. Required for AireOS.

	InterfaceGateway string `json:"interfaceGateway,omitempty"` // Interface Gateway. Required for AireOS.

	LagOrPortNumber *int `json:"lagOrPortNumber,omitempty"` // Lag Or Port Number. Required for AireOS.

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID. Required for AireOS and EWLC.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name. Required for AireOS and EWLC.
}
type RequestWirelessProvision []RequestItemWirelessProvision // Array of RequestWirelessProvision
type RequestItemWirelessProvision struct {
	DeviceName string `json:"deviceName,omitempty"` // Controller Name

	Site string `json:"site,omitempty"` // Full Site Hierarchy where device has to be assigned

	ManagedApLocations []string `json:"managedAPLocations,omitempty"` // List of managed AP locations (Site Hierarchies)

	DynamicInterfaces *[]RequestItemWirelessProvisionDynamicInterfaces `json:"dynamicInterfaces,omitempty"` //
}
type RequestItemWirelessProvisionDynamicInterfaces struct {
	InterfaceIPAddress string `json:"interfaceIPAddress,omitempty"` // Interface IP Address. Required for AireOS.

	InterfaceNetmaskInCIDR *int `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR. Required for AireOS.

	InterfaceGateway string `json:"interfaceGateway,omitempty"` // Interface Gateway.  Required for AireOS.

	LagOrPortNumber *int `json:"lagOrPortNumber,omitempty"` // Lag Or Port Number.  Required for AireOS.

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID. Required for both AireOS and EWLC.

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name. Required for both AireOS and EWLC.
}
type RequestWirelessPSKOverride []RequestItemWirelessPSKOverride // Array of RequestWirelessPSKOverride
type RequestItemWirelessPSKOverride struct {
	SSID string `json:"ssid,omitempty"` // enterprise ssid name(already created/present)

	Site string `json:"site,omitempty"` // site name hierarchy (ex: Global/aaa/zzz/...)

	PassPhrase string `json:"passPhrase,omitempty"` // Pass phrase (create/update)

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name
}
type RequestWirelessCreateOrUpdateRfProfile struct {
	Name string `json:"name,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // is Default Rf Profile

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // Enable Radio Type A

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // Enable Radio Type B

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	EnableCustom *bool `json:"enableCustom,omitempty"` // Enable Custom

	EnableBrownField *bool `json:"enableBrownField,omitempty"` // Enable Brown Field

	RadioTypeAProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioTypeCProperties *RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties `json:"radioTypeCProperties,omitempty"` //

	EnableRadioTypeC *bool `json:"enableRadioTypeC,omitempty"` // Enable Radio Type C (6GHz)
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "36,40,44,48,52,56,60,64,149,153,157,161")

	DataRates string `json:"dataRates,omitempty"` // Data Rates (Default : "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates (Default: "6,12,24")

	PowerThreshold *float64 `json:"powerThreshold,omitempty"` // Power Threshold  ( (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Rx Sop Threshold  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "9,11,12,18,24,36,48,54")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "9,11,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "12")

	PowerThreshold *float64 `json:"powerThreshold,omitempty"` // Power Threshold   (Default: -70)

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent Profile (Default : CUSTOM)

	RadioChannels string `json:"radioChannels,omitempty"` // Radio Channels (Default : "5,21,37,53,69,85,101,117,133,149,165,181,197,213,229")

	DataRates string `json:"dataRates,omitempty"` // Data Rates  (Default: "6,9,12,18,24,36,48,54")

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates  (Default: "6,12,24")

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // Rx Sop Threshold  (Default: "AUTO")

	MinPowerLevel *float64 `json:"minPowerLevel,omitempty"` // Min Power Level  (Default: -10)

	MaxPowerLevel *float64 `json:"maxPowerLevel,omitempty"` // Max Power Level  (Default: 30)

	PowerThreshold *float64 `json:"powerThreshold,omitempty"` // Power Threshold   (Default: -70)
}
type RequestWirelessFactoryResetAccessPoints struct {
	KeepStaticIPConfig *bool `json:"keepStaticIPConfig,omitempty"` // Set the value of keepStaticIPConfig to false, to clear all configurations from Access Points and set the value of keepStaticIPConfig to true, to clear all configurations from Access Points without clearing static IP configuration.

	ApMacAddresses []string `json:"apMacAddresses,omitempty"` // List of Access Point's Ethernet MAC addresses, set maximum 100 ethernet MAC addresses per request.
}
type RequestWirelessApProvision struct {
	NetworkDevices *[]RequestWirelessApProvisionNetworkDevices `json:"networkDevices,omitempty"` //

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name. RF Profile is not allowed for custom AP Zones.

	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.

	SiteID string `json:"siteId,omitempty"` // Site ID
}
type RequestWirelessApProvisionNetworkDevices struct {
	DeviceID string `json:"deviceId,omitempty"` // Network device ID of access points

	MeshRole string `json:"meshRole,omitempty"` // Mesh Role (Applicable only when AP is in Bridge Mode)
}
type RequestWirelessMobilityProvision struct {
	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Self device Group Name. Must be alphanumeric without {!,<,space,?/'}  and maximum of 31 characters.

	MacAddress string `json:"macAddress,omitempty"` // Device mobility MAC Address. Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ManagementIP string `json:"managementIp,omitempty"` // Self device wireless Management IP.

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Obtain the network device ID value by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}.

	DtlsHighCipher *bool `json:"dtlsHighCipher,omitempty"` // DTLS High Cipher.

	DataLinkEncryption *bool `json:"dataLinkEncryption,omitempty"` // A secure link in which data is encrypted using CAPWAP DTLS protocol can be established between two controllers. This value will be applied to all peers during POST operation.

	MobilityPeers *[]RequestWirelessMobilityProvisionMobilityPeers `json:"mobilityPeers,omitempty"` //
}
type RequestWirelessMobilityProvisionMobilityPeers struct {
	PeerIP string `json:"peerIp,omitempty"` // This indicates public ip address.

	PrivateIPAddress string `json:"privateIpAddress,omitempty"` // This indicates private/management ip address.

	PeerDeviceName string `json:"peerDeviceName,omitempty"` // Peer device Host Name.

	PeerNetworkDeviceID string `json:"peerNetworkDeviceId,omitempty"` // The possible values are: UNKNOWN or valid UUID of Network device Id. UNKNOWN represents out of band device which is not managed internally. Valid UUID represents WLC network device id.

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} <br/> and maximum of 31 characters.

	MemberMacAddress string `json:"memberMacAddress,omitempty"` // Peer device mobility MAC Address.  Allowed formats are: 0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	DeviceSeries string `json:"deviceSeries,omitempty"` // Indicates peer device mobility belongs to AireOS or IOX-XE family. 0 - indicates AireOS and 1 - indicates C9800.

	HashKey string `json:"hashKey,omitempty"` // SSC hash string must be 40 characters.
}
type RequestWirelessMobilityReset struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device Id of Cisco wireless controller. Obtain the network device ID value by using the API call GET - /dna/intent/api/v1/network-device/ip-address/${ipAddress}.
}
type RequestWirelessAssignManagedApLocationsForWLC struct {
	PrimaryManagedApLocationsSiteIDs []string `json:"primaryManagedAPLocationsSiteIds,omitempty"` // Site IDs of Primary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site

	SecondaryManagedApLocationsSiteIDs []string `json:"secondaryManagedAPLocationsSiteIds,omitempty"` // Site IDs of Secondary Managed AP Locations. These values can be obtained by using the API call GET: /dna/intent/api/v1/site
}
type RequestWirelessWirelessControllerProvision struct {
	Interfaces *[]RequestWirelessWirelessControllerProvisionInterfaces `json:"interfaces,omitempty"` //

	SkipApProvision *bool `json:"skipApProvision,omitempty"` // True if Skip AP Provision is enabled, else False

	RollingApUpgrade *RequestWirelessWirelessControllerProvisionRollingApUpgrade `json:"rollingApUpgrade,omitempty"` //

	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // AP Authorization List name. 'Obtain the AP Authorization List names by using the API call GET: /intent/api/v1/wirelessSettings/apAuthorizationLists. During re-provision, obtain the AP Authorization List configured for the given provisioned network device Id using the API call GET: /intent/api/v1/wireless/apAuthorizationLists/{networkDeviceId}'

	AuthorizeMeshAndNonMeshAccessPoints *bool `json:"authorizeMeshAndNonMeshAccessPoints,omitempty"` // True if AP Authorization List should  authorize against All Mesh/Non-Mesh APs, else false if AP Authorization List should only authorize against Mesh APs (Applicable only when Mesh is enabled on sites)
}
type RequestWirelessWirelessControllerProvisionInterfaces struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID range is 1 - 4094

	InterfaceIPAddress string `json:"interfaceIPAddress,omitempty"` // Interface IP Address

	InterfaceNetmaskInCIDR *int `json:"interfaceNetmaskInCIDR,omitempty"` // Interface Netmask In CIDR, range is 1-30

	InterfaceGateway string `json:"interfaceGateway,omitempty"` // Interface Gateway

	LagOrPortNumber *int `json:"lagOrPortNumber,omitempty"` // Lag Or Port Number
}
type RequestWirelessWirelessControllerProvisionRollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // True if Rolling AP Upgrade is enabled, else False

	ApRebootPercentage *int `json:"apRebootPercentage,omitempty"` // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type RequestWirelessCreateWirelessProfileConnectivity struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Network Profile Name

	SSIDDetails *[]RequestWirelessCreateWirelessProfileConnectivitySSIDDetails `json:"ssidDetails,omitempty"` //

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // These additional interfaces will be configured on the device as independent interfaces in addition to the interfaces mapped to SSIDs. Max Limit 4094

	ApZones *[]RequestWirelessCreateWirelessProfileConnectivityApZones `json:"apZones,omitempty"` //
}
type RequestWirelessCreateWirelessProfileConnectivitySSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *RequestWirelessCreateWirelessProfileConnectivitySSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name.

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type RequestWirelessCreateWirelessProfileConnectivitySSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type RequestWirelessCreateWirelessProfileConnectivityApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type RequestWirelessUpdateWirelessProfileConnectivity struct {
	WirelessProfileName string `json:"wirelessProfileName,omitempty"` // Wireless Network Profile Name

	SSIDDetails *[]RequestWirelessUpdateWirelessProfileConnectivitySSIDDetails `json:"ssidDetails,omitempty"` //

	AdditionalInterfaces []string `json:"additionalInterfaces,omitempty"` // These additional interfaces will be configured on the device as independent interfaces in addition to the interfaces mapped to SSIDs. Max Limit 4094

	ApZones *[]RequestWirelessUpdateWirelessProfileConnectivityApZones `json:"apZones,omitempty"` //
}
type RequestWirelessUpdateWirelessProfileConnectivitySSIDDetails struct {
	SSIDName string `json:"ssidName,omitempty"` // SSID Name

	FlexConnect *RequestWirelessUpdateWirelessProfileConnectivitySSIDDetailsFlexConnect `json:"flexConnect,omitempty"` //

	EnableFabric *bool `json:"enableFabric,omitempty"` // True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	WLANProfileName string `json:"wlanProfileName,omitempty"` // WLAN Profile Name

	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name.

	Dot11BeProfileID string `json:"dot11beProfileId,omitempty"` // 802.11be Profile Id. Applicable to IOS controllers with version 17.15 and higher. 802.11be Profiles if passed, should be same across all SSIDs in network profile being configured

	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name

	VLANGroupName string `json:"vlanGroupName,omitempty"` // VLAN Group Name
}
type RequestWirelessUpdateWirelessProfileConnectivitySSIDDetailsFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time

	LocalToVLAN *int `json:"localToVlan,omitempty"` // Local to VLAN ID
}
type RequestWirelessUpdateWirelessProfileConnectivityApZones struct {
	ApZoneName string `json:"apZoneName,omitempty"` // AP Zone Name

	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	SSIDs []string `json:"ssids,omitempty"` // ssids part of apZone
}
type RequestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type RequestWirelessUpdateASpecificPolicyTagForAWirelessProfile struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	PolicyTagName string `json:"policyTagName,omitempty"` // Policy Tag Name

	ApZones []string `json:"apZones,omitempty"` // Ap Zones
}
type RequestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk struct {
	Items *[][]string `json:"items,omitempty"` // Root
}
type RequestWirelessUpdateASpecificSiteTagForAWirelessProfile struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Site Ids

	SiteTagName string `json:"siteTagName,omitempty"` // Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space.

	FlexProfileName string `json:"flexProfileName,omitempty"` // Flex Profile Name

	ApProfileName string `json:"apProfileName,omitempty"` // Ap Profile Name
}
type RequestWirelessCreateAnchorGroup struct {
	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]RequestWirelessCreateAnchorGroupMobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type RequestWirelessCreateAnchorGroupMobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255
}
type RequestWirelessUpdateAnchorGroup struct {
	AnchorGroupName string `json:"anchorGroupName,omitempty"` // Anchor Group Name. Max length is 32 characters

	MobilityAnchors *[]RequestWirelessUpdateAnchorGroupMobilityAnchors `json:"mobilityAnchors,omitempty"` //
}
type RequestWirelessUpdateAnchorGroupMobilityAnchors struct {
	DeviceName string `json:"deviceName,omitempty"` // Peer Host Name

	IPAddress string `json:"ipAddress,omitempty"` // This indicates Mobility public IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255

	AnchorPriority string `json:"anchorPriority,omitempty"` // This indicates anchor priority.  Priority values range from 1 (high) to 3 (low). Primary, secondary or tertiary and defined priority is displayed with guest anchor. Only one priority value is allowed per anchor WLC.

	ManagedAnchorWlc *bool `json:"managedAnchorWlc,omitempty"` // This indicates whether the Wireless LAN Controller supporting Anchor is managed by the Network Controller or not. True means this is managed by Network Controller.

	PeerDeviceType string `json:"peerDeviceType,omitempty"` // Indicates peer device mobility belongs to AireOS or IOS-XE family.

	MacAddress string `json:"macAddress,omitempty"` // Peer Device mobility MAC address. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	MobilityGroupName string `json:"mobilityGroupName,omitempty"` // Peer Device mobility group Name. Must be alphanumeric without {!,<,space,?/'} and maximum of 31 characters.

	PrivateIP string `json:"privateIp,omitempty"` // This indicates private management IP address. Allowed formats are 192.168.0.1, 10.0.0.1, 255.255.255.255
}
type RequestWirelessCreateApAuthorizationList struct {
	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // AP Authorization List Name. For a AP Authorization List to be created successfully, either Local Authorization or Remote Authorization is mandatory.

	LocalAuthorization *RequestWirelessCreateApAuthorizationListLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *RequestWirelessCreateApAuthorizationListRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type RequestWirelessCreateApAuthorizationListLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // List of Access Point's Ethernet MAC addresses. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // List of Access Point's Serial Numbers.
}
type RequestWirelessCreateApAuthorizationListRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // List of Authorization server IpAddresses. Obtain the AAA servers by using the API GET call '/dna/intent/api/v1/authentication-policy-servers'.

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // True if AP Authorization List should authorise APs With MAC addresses, else False. (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // True if server IpAddresses are added and AP Authorization List should authorise APs With Serial Numbers, else False (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)
}
type RequestWirelessUpdateApAuthorizationList struct {
	ApAuthorizationListName string `json:"apAuthorizationListName,omitempty"` // AP Authorization List Name. For a AP Authorization List to be created successfully, either Local Authorization or Remote Authorization is mandatory.

	LocalAuthorization *RequestWirelessUpdateApAuthorizationListLocalAuthorization `json:"localAuthorization,omitempty"` //

	RemoteAuthorization *RequestWirelessUpdateApAuthorizationListRemoteAuthorization `json:"remoteAuthorization,omitempty"` //
}
type RequestWirelessUpdateApAuthorizationListLocalAuthorization struct {
	ApMacEntries []string `json:"apMacEntries,omitempty"` // List of Access Point's Ethernet MAC addresses. Allowed formats are:0a0b.0c01.0211, 0a0b0c010211, 0a:0b:0c:01:02:11

	ApSerialNumberEntries []string `json:"apSerialNumberEntries,omitempty"` // List of Access Point's Serial Numbers.
}
type RequestWirelessUpdateApAuthorizationListRemoteAuthorization struct {
	AAAServers []string `json:"aaaServers,omitempty"` // List of Authorization server IpAddresses. Obtain the AAA servers by using the API GET call '/dna/intent/api/v1/authentication-policy-servers'.

	AuthorizeApWithMac *bool `json:"authorizeApWithMac,omitempty"` // True if AP Authorization List should authorise APs With MAC addresses, else False. (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)

	AuthorizeApWithSerialNumber *bool `json:"authorizeApWithSerialNumber,omitempty"` // True if server IpAddresses are added and AP Authorization List should authorise APs With Serial Numbers, else False (For Non-Mesh Access Points, either of Authorize AP With MAC Address or Serial Number is required to be set to true)
}
type RequestWirelessCreateApProfile struct {
	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *RequestWirelessCreateApProfileManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *RequestWirelessCreateApProfileRogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *RequestWirelessCreateApProfileMeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *RequestWirelessCreateApProfileCalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Enter the hour value (HH). The valid range is from -12 through 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Enter the minute value (MM). The valid range is from 0 through 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type RequestWirelessCreateApProfileManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type RequestWirelessCreateApProfileRogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type RequestWirelessCreateApProfileMeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used.
}
type RequestWirelessCreateApProfileCalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *RequestWirelessCreateApProfileCalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type RequestWirelessCreateApProfileCalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type RequestWirelessUpdateApProfileByID struct {
	ApProfileName string `json:"apProfileName,omitempty"` // Name of the Access Point profile. Max length is 32 characters.

	Description string `json:"description,omitempty"` // Description of the AP profile. Max length is 241 characters

	RemoteWorkerEnabled *bool `json:"remoteWorkerEnabled,omitempty"` // Indicates if remote worker mode is enabled on the AP. Remote teleworker enabled profile cannot support security features like aWIPS,Forensic Capture Enablement, Rogue Detection and Rogue Containment.

	ManagementSetting *RequestWirelessUpdateApProfileByIDManagementSetting `json:"managementSetting,omitempty"` //

	AwipsEnabled *bool `json:"awipsEnabled,omitempty"` // Indicates if AWIPS is enabled on the AP.

	AwipsForensicEnabled *bool `json:"awipsForensicEnabled,omitempty"` // Indicates if AWIPS forensic is enabled on the AP. Forensic Capture is supported from IOS-XE version 17.4 and above. Forensic Capture can be activated only if aWIPS is enabled.

	RogueDetectionSetting *RequestWirelessUpdateApProfileByIDRogueDetectionSetting `json:"rogueDetectionSetting,omitempty"` //

	PmfDenialEnabled *bool `json:"pmfDenialEnabled,omitempty"` // Indicates if PMF denial is active on the AP. PMF Denial is supported from IOS-XE version 17.12 and above.

	MeshEnabled *bool `json:"meshEnabled,omitempty"` // This indicates whether mesh networking is enabled on the AP. For IOS-XE devices, when mesh networking is enabled, a custom mesh profile with the configured parameters will be created and mapped to the AP join profile on the device. When mesh networking is disabled, any existing custom mesh profile will be deleted from the device, and the AP join profile will be mapped to the default mesh profile on the device.

	MeshSetting *RequestWirelessUpdateApProfileByIDMeshSetting `json:"meshSetting,omitempty"` //

	ApPowerProfileName string `json:"apPowerProfileName,omitempty"` // Name of the existing AP power profile.

	CalendarPowerProfiles *RequestWirelessUpdateApProfileByIDCalendarPowerProfiles `json:"calendarPowerProfiles,omitempty"` //

	CountryCode string `json:"countryCode,omitempty"` // Country Code

	TimeZone string `json:"timeZone,omitempty"` // In the Time Zone area, choose one of the following options.             Not Configured - APs operate in the UTC time zone.             Controller - APs operate in the Cisco Wireless Controller time zone.             Delta from Controller - APs operate in the offset time from the wireless controller time zone.

	TimeZoneOffsetHour *int `json:"timeZoneOffsetHour,omitempty"` // Enter the hour value (HH). The valid range is from -12 through 14.

	TimeZoneOffsetMinutes *int `json:"timeZoneOffsetMinutes,omitempty"` // Enter the minute value (MM). The valid range is from 0 through 59.

	ClientLimit *int `json:"clientLimit,omitempty"` // Number of clients. Value should be between 0-1200.
}
type RequestWirelessUpdateApProfileByIDManagementSetting struct {
	AuthType string `json:"authType,omitempty"` // Authentication type used in the AP profile. These setting are applicable during PnP claim and for day-N authentication of AP. Changing these settings will be service impacting for the PnP onboarded APs and will need a factory-reset for those APs.

	Dot1XUsername string `json:"dot1xUsername,omitempty"` // Username for 802.1X authentication. dot1xUsername must have a minimum of 1 character and a maximum of 32 characters.

	Dot1XPassword string `json:"dot1xPassword,omitempty"` // Password for 802.1X authentication. AP dot1x password length should not exceed 120.

	SSHEnabled *bool `json:"sshEnabled,omitempty"` // Indicates if SSH is enabled on the AP. Enable SSH add credentials for device management.

	TelnetEnabled *bool `json:"telnetEnabled,omitempty"` // Indicates if Telnet is enabled on the AP. Enable Telnet to add credentials for device management.

	ManagementUserName string `json:"managementUserName,omitempty"` // Management username must have a minimum of 1 character and a maximum of 32 characters.

	ManagementPassword string `json:"managementPassword,omitempty"` // Management password for the AP. Length must be 8-120 characters.

	ManagementEnablePassword string `json:"managementEnablePassword,omitempty"` // Enable password for managing the AP. Length must be 8-120 characters.

	CdpState *bool `json:"cdpState,omitempty"` // Indicates if CDP is enabled on the AP. Enable CDP in order to make Cisco Access Points known to its neighboring devices and vice-versa.
}
type RequestWirelessUpdateApProfileByIDRogueDetectionSetting struct {
	RogueDetection *bool `json:"rogueDetection,omitempty"` // Indicates if rogue detection is enabled on the AP. Detect Access Points that have been installed on a secure network without explicit authorization from a system administrator and configure rogue general configuration parameters

	RogueDetectionMinRssi *int `json:"rogueDetectionMinRssi,omitempty"` // Minimum RSSI for rogue detection. Value should be in range -128 decibel milliwatts and -70 decibel milliwatts

	RogueDetectionTransientInterval *int `json:"rogueDetectionTransientInterval,omitempty"` // Transient interval for rogue detection. Value should be 0 or from 120 to 1800.

	RogueDetectionReportInterval *int `json:"rogueDetectionReportInterval,omitempty"` // Report interval for rogue detection. Value should be in range 10 and 300.
}
type RequestWirelessUpdateApProfileByIDMeshSetting struct {
	BridgeGroupName string `json:"bridgeGroupName,omitempty"` // Name of the bridge group for mesh settings. If not configured, 'Default' Bridge group name will be used in mesh profile.

	BackhaulClientAccess *bool `json:"backhaulClientAccess,omitempty"` // Indicates if backhaul client access is enabled on the AP.

	Range *int `json:"range,omitempty"` // Range of the mesh network. Value should be between 150-132000

	Ghz5BackhaulDataRates string `json:"ghz5BackhaulDataRates,omitempty"` // 5GHz backhaul data rates.

	Ghz24BackhaulDataRates string `json:"ghz24BackhaulDataRates,omitempty"` // 2.4GHz backhaul data rates.

	RapDownlinkBackhaul string `json:"rapDownlinkBackhaul,omitempty"` // Type of downlink backhaul used. Available values 5 GHz, 2.4 GHz.
}
type RequestWirelessUpdateApProfileByIDCalendarPowerProfiles struct {
	PowerProfileName string `json:"powerProfileName,omitempty"` // Name of the existing AP power profile to be mapped to the calendar power profile. API-/intent/api/v1/wirelessSettings/powerProfiles.

	SchedulerType string `json:"schedulerType,omitempty"` // Type of the scheduler.

	Duration *RequestWirelessUpdateApProfileByIDCalendarPowerProfilesDuration `json:"duration,omitempty"` //
}
type RequestWirelessUpdateApProfileByIDCalendarPowerProfilesDuration struct {
	SchedulerStartTime string `json:"schedulerStartTime,omitempty"` // Start time of the duration setting.

	SchedulerEndTime string `json:"schedulerEndTime,omitempty"` // End time of the duration setting.

	SchedulerDay string `json:"schedulerDay,omitempty"` // Applies every week on the selected days

	SchedulerDate string `json:"schedulerDate,omitempty"` // Start and End date of the duration setting, applicable for MONTHLY schedulers.
}
type RequestWirelessCreateA80211BeProfile struct {
	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink (Default: true)

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink (Default: true)

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink (Default: false)

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink (Default: false)

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU (Default: false)
}
type RequestWirelessUpdate80211BeProfile struct {
	ProfileName string `json:"profileName,omitempty"` // 802.11be Profile Name

	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink (Default: true)

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink (Default: true)

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink (Default: false)

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink (Default: false)

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU (Default: false)
}
type RequestWirelessCreateInterface struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID range is 1-4094
}
type RequestWirelessUpdateInterface struct {
	InterfaceName string `json:"interfaceName,omitempty"` // Interface Name

	VLANID *int `json:"vlanId,omitempty"` // VLAN ID range is 1-4094
}
type RequestWirelessCreatePowerProfile struct {
	ProfileName string `json:"profileName,omitempty"` // Name of the Power Profile. Max allowed characters is 128

	Description string `json:"description,omitempty"` // Description of the Power Profile. Max allowed characters is 128

	Rules *[]RequestWirelessCreatePowerProfileRules `json:"rules,omitempty"` //
}
type RequestWirelessCreatePowerProfileRules struct {
	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type for the rule.

	InterfaceID string `json:"interfaceId,omitempty"` // Interface Id for the rule.

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type for the rule.

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value for the rule.
}
type RequestWirelessUpdatePowerProfileByID struct {
	ProfileName string `json:"profileName,omitempty"` // Name of the Power Profile. Max length is 32 characters

	Description string `json:"description,omitempty"` // Description of the Power Profile. Max length is 32 characters

	Rules *[]RequestWirelessUpdatePowerProfileByIDRules `json:"rules,omitempty"` //
}
type RequestWirelessUpdatePowerProfileByIDRules struct {
	InterfaceType string `json:"interfaceType,omitempty"` // Interface Type

	InterfaceID string `json:"interfaceID,omitempty"` // Interface ID

	ParameterType string `json:"parameterType,omitempty"` // Parameter Type

	ParameterValue string `json:"parameterValue,omitempty"` // Parameter Value
}
type RequestWirelessCreateRfProfile struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	RadioTypeAProperties *RequestWirelessCreateRfProfileRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *RequestWirelessCreateRfProfileRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *RequestWirelessCreateRfProfileRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *RequestWirelessCreateRfProfileRadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessCreateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessCreateRfProfileRadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileRadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type RequestWirelessCreateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessCreateRfProfileRadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessCreateRfProfileRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *RequestWirelessCreateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessCreateRfProfileRadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessCreateRfProfileRadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessCreateRfProfileRadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width (Permissible Values:20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values:20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *RequestWirelessCreateRfProfileRadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessCreateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessCreateRfProfileRadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessCreateRfProfileRadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessUpdateRfProfile struct {
	RfProfileName string `json:"rfProfileName,omitempty"` // RF Profile Name

	DefaultRfProfile *bool `json:"defaultRfProfile,omitempty"` // True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time

	EnableRadioTypeA *bool `json:"enableRadioTypeA,omitempty"` // True if 5 GHz radio band is enabled in the RF Profile, else False

	EnableRadioTypeB *bool `json:"enableRadioTypeB,omitempty"` // True if 2.4 GHz radio band is enabled in the RF Profile, else False

	EnableRadioType6GHz *bool `json:"enableRadioType6GHz,omitempty"` // True if 6 GHz radio band is enabled in the RF Profile, else False

	RadioTypeAProperties *RequestWirelessUpdateRfProfileRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //

	RadioTypeBProperties *RequestWirelessUpdateRfProfileRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //

	RadioType6GHzProperties *RequestWirelessUpdateRfProfileRadioType6GHzProperties `json:"radioType6GHzProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileRadioTypeAProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 5 GHz radio band. In case of brownfield learnt RF Profile if the parent profile is GLOBAL, any change in RF Profile configurations will not be provisioned to device. Existing parent profile with values of HIGH, TYPICAL, LOW or CUSTOM cannot be modified to GLOBAL

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173

	DataRates string `json:"dataRates,omitempty"` // Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 5 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 5 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 5 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 5 GHz radio band

	ChannelWidth string `json:"channelWidth,omitempty"` // Channel Width

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	ZeroWaitDfsEnable *bool `json:"zeroWaitDfsEnable,omitempty"` // Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 5 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 5 GHz radio band

	FraProperties *RequestWirelessUpdateRfProfileRadioTypeAPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessUpdateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessUpdateRfProfileRadioTypeAPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileRadioTypeAPropertiesFraProperties struct {
	ClientAware *bool `json:"clientAware,omitempty"` // Client Aware of 5 GHz radio band

	ClientSelect *int `json:"clientSelect,omitempty"` // Client Select(%) of 5 GHz radio band

	ClientReset *int `json:"clientReset,omitempty"` // Client Reset(%) of 5 GHz radio band
}
type RequestWirelessUpdateRfProfileRadioTypeAPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessUpdateRfProfileRadioTypeAPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessUpdateRfProfileRadioTypeBProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 2.4 GHz radio band. In case of brownfield learnt RF Profile if the parent profile is GLOBAL, any change in RF Profile configurations will not be provisioned to device. Existing parent profile with values of HIGH, TYPICAL, LOW or CUSTOM cannot be modified to GLOBAL

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14

	DataRates string `json:"dataRates,omitempty"` // Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 2.4 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 2.4 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 2.4 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 2.4 GHz radio band

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 2.4 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 2.4 GHz radio band

	CoverageHoleDetectionProperties *RequestWirelessUpdateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessUpdateRfProfileRadioTypeBPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileRadioTypeBPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessUpdateRfProfileRadioTypeBPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessUpdateRfProfileRadioType6GHzProperties struct {
	ParentProfile string `json:"parentProfile,omitempty"` // Parent profile of 6 GHz radio band

	RadioChannels string `json:"radioChannels,omitempty"` // DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233

	DataRates string `json:"dataRates,omitempty"` // Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` // Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54

	PowerThreshold *int `json:"powerThreshold,omitempty"` // Power threshold of 6 GHz radio band

	RxSopThreshold string `json:"rxSopThreshold,omitempty"` // RX-SOP threshold of 6 GHz radio band

	MinPowerLevel *int `json:"minPowerLevel,omitempty"` // Minimum power level of 6 GHz radio band

	MaxPowerLevel *int `json:"maxPowerLevel,omitempty"` // Maximum power level of 6 GHz radio band

	EnableStandardPowerService *bool `json:"enableStandardPowerService,omitempty"` // True if Standard Power Service is enabled, else False

	MultiBssidProperties *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties `json:"multiBssidProperties,omitempty"` //

	PreamblePuncture *bool `json:"preamblePuncture,omitempty"` // Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher

	MinDbsWidth *int `json:"minDbsWidth,omitempty"` // Minimum DBS Width (Permissible Values:20,40,80,160,320)

	MaxDbsWidth *int `json:"maxDbsWidth,omitempty"` // Maximum DBS Width (Permissible Values:20,40,80,160,320)

	CustomRxSopThreshold *int `json:"customRxSopThreshold,omitempty"` // RX-SOP threshold custom configuration of 6 GHz radio band

	MaxRadioClients *int `json:"maxRadioClients,omitempty"` // Client Limit of 6 GHz radio band

	PscEnforcingEnabled *bool `json:"pscEnforcingEnabled,omitempty"` // PSC Enforcing Enable for 6 GHz radio band

	DiscoveryFrames6GHz string `json:"discoveryFrames6GHz,omitempty"` // Discovery Frames of 6 GHz radio band

	BroadcastProbeResponseInterval *int `json:"broadcastProbeResponseInterval,omitempty"` // Broadcast Probe Response Interval of 6 GHz radio band

	FraProperties *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesFraProperties `json:"fraProperties,omitempty"` //

	CoverageHoleDetectionProperties *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties `json:"coverageHoleDetectionProperties,omitempty"` //

	SpatialReuseProperties *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesSpatialReuseProperties `json:"spatialReuseProperties,omitempty"` //
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidProperties struct {
	Dot11AxParameters *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters `json:"dot11axParameters,omitempty"` //

	Dot11BeParameters *RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters `json:"dot11beParameters,omitempty"` //

	TargetWakeTime *bool `json:"targetWakeTime,omitempty"` // Target Wake Time

	TwtBroadcastSupport *bool `json:"twtBroadcastSupport,omitempty"` // TWT Broadcast Support
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters struct {
	OfdmaDownLink *bool `json:"ofdmaDownLink,omitempty"` // OFDMA Downlink

	OfdmaUpLink *bool `json:"ofdmaUpLink,omitempty"` // OFDMA Uplink

	MuMimoUpLink *bool `json:"muMimoUpLink,omitempty"` // MU-MIMO Uplink

	MuMimoDownLink *bool `json:"muMimoDownLink,omitempty"` // MU-MIMO Downlink

	OfdmaMultiRu *bool `json:"ofdmaMultiRu,omitempty"` // OFDMA Multi-RU
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesFraProperties struct {
	ClientResetCount *int `json:"clientResetCount,omitempty"` // Client Reset Count of 6 GHz radio band

	ClientUtilizationThreshold *int `json:"clientUtilizationThreshold,omitempty"` // Client Utilization Threshold of 6 GHz radio band
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesCoverageHoleDetectionProperties struct {
	ChdClientLevel *int `json:"chdClientLevel,omitempty"` // Coverage Hole Detection Client Level

	ChdDataRssiThreshold *int `json:"chdDataRssiThreshold,omitempty"` // Coverage Hole Detection Data Rssi Threshold

	ChdVoiceRssiThreshold *int `json:"chdVoiceRssiThreshold,omitempty"` // Coverage Hole Detection Voice Rssi Threshold

	ChdExceptionLevel *int `json:"chdExceptionLevel,omitempty"` // Coverage Hole Detection Exception Level(%)
}
type RequestWirelessUpdateRfProfileRadioType6GHzPropertiesSpatialReuseProperties struct {
	Dot11AxNonSrgObssPacketDetect *bool `json:"dot11axNonSrgObssPacketDetect,omitempty"` // Dot11ax Non SRG OBSS PD

	Dot11AxNonSrgObssPacketDetectMaxThreshold *int `json:"dot11axNonSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax Non SRG OBSS PD Max Threshold

	Dot11AxSrgObssPacketDetect *bool `json:"dot11axSrgObssPacketDetect,omitempty"` // Dot11ax SRG OBSS PD

	Dot11AxSrgObssPacketDetectMinThreshold *int `json:"dot11axSrgObssPacketDetectMinThreshold,omitempty"` // Dot11ax SRG OBSS PD Min Threshold

	Dot11AxSrgObssPacketDetectMaxThreshold *int `json:"dot11axSrgObssPacketDetectMaxThreshold,omitempty"` // Dot11ax SRG OBSS PD Max Threshold
}
type RequestWirelessAssignAnchorManagedApLocationsForWLC struct {
	AnchorManagedApLocationsSiteIDs []string `json:"anchorManagedAPLocationsSiteIds,omitempty"` // This API allows user to assign Anchor Managed AP Locations for WLC by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.               -  When anchor managed location array present then it will add the anchor managed locations.
}
type RequestWirelessConfigureAccessPointsV2 struct {
	ApList *[]RequestWirelessConfigureAccessPointsV2ApList `json:"apList,omitempty"` //

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the access point's admin status, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureApMode *bool `json:"configureApMode,omitempty"` // To change the access point's mode, set this parameter's value to "true".

	ApMode *int `json:"apMode,omitempty"` // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".

	ConfigureFailoverPriority *bool `json:"configureFailoverPriority,omitempty"` // To change the access point's failover priority, set this parameter's value to "true".

	FailoverPriority *int `json:"failoverPriority,omitempty"` // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".

	ConfigureLedStatus *bool `json:"configureLedStatus,omitempty"` // To change the access point's LED status, set this parameter's value to "true".

	LedStatus *bool `json:"ledStatus,omitempty"` // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.

	ConfigureLedBrightnessLevel *bool `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".

	LedBrightnessLevel *int `json:"ledBrightnessLevel,omitempty"` // Configure the access point's LED brightness level by setting a value between 1 and 8.

	ConfigureLocation *bool `json:"configureLocation,omitempty"` // To change the access point's location, set this parameter's value to "true".

	Location string `json:"location,omitempty"` // Configure the access point's location.

	ConfigureHAController *bool `json:"configureHAController,omitempty"` // To change the access point's HA controller, set this parameter's value to "true".

	PrimaryControllerName string `json:"primaryControllerName,omitempty"` // Configure the hostname for an access point's primary controller.

	PrimaryIPAddress *RequestWirelessConfigureAccessPointsV2PrimaryIPAddress `json:"primaryIpAddress,omitempty"` //

	SecondaryControllerName string `json:"secondaryControllerName,omitempty"` // Configure the hostname for an access point's secondary controller.

	SecondaryIPAddress *RequestWirelessConfigureAccessPointsV2SecondaryIPAddress `json:"secondaryIpAddress,omitempty"` //

	TertiaryControllerName string `json:"tertiaryControllerName,omitempty"` // Configure the hostname for an access point's tertiary controller.

	TertiaryIPAddress *RequestWirelessConfigureAccessPointsV2TertiaryIPAddress `json:"tertiaryIpAddress,omitempty"` //

	RadioConfigurations *[]RequestWirelessConfigureAccessPointsV2RadioConfigurations `json:"radioConfigurations,omitempty"` //

	ConfigureCleanAirSI24Ghz *bool `json:"configureCleanAirSI24Ghz,omitempty"` // To change the clean air status for radios that are in 2.4 Ghz band, set this parameter's value to "true".

	CleanAirSI24 *bool `json:"cleanAirSI24,omitempty"` // Configure clean air status for radios that are in 2.4 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureCleanAirSI5Ghz *bool `json:"configureCleanAirSI5Ghz,omitempty"` // To change the clean air status for radios that are in 5 Ghz band, set this parameter's value to "true".

	CleanAirSI5 *bool `json:"cleanAirSI5,omitempty"` // Configure clean air status for radios that are in 5 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureCleanAirSI6Ghz *bool `json:"configureCleanAirSI6Ghz,omitempty"` // To change the clean air status for radios that are in 6 Ghz band, set this parameter's value to "true".

	CleanAirSI6 *bool `json:"cleanAirSI6,omitempty"` // Configure clean air status for radios that are in 6 Ghz band. Set this parameter's value to "true" to enable it and "false" to disable it.

	IsAssignedSiteAsLocation *bool `json:"isAssignedSiteAsLocation,omitempty"` // To configure the access point's location as the site assigned to the access point, set this parameter's value to "true".
}
type RequestWirelessConfigureAccessPointsV2ApList struct {
	ApName string `json:"apName,omitempty"` // The current host name of the access point.

	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.

	ApNameNew string `json:"apNameNew,omitempty"` // The modified hostname of the access point.
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
	ConfigureRadioRoleAssignment *bool `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".

	RadioRoleAssignment string `json:"radioRoleAssignment,omitempty"` // Configure only one of the following roles on the specified radio for an access point as "AUTO", "SERVING", or "MONITOR". Any other string is invalid, including empty string

	RadioBand string `json:"radioBand,omitempty"` // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5". Any other string is invalid, including empty string

	ConfigureAdminStatus *bool `json:"configureAdminStatus,omitempty"` // To change the admin status on the specified radio for an access point, set this parameter's value to "true".

	AdminStatus *bool `json:"adminStatus,omitempty"` // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.

	ConfigureAntennaPatternName *bool `json:"configureAntennaPatternName,omitempty"` // To change the antenna gain on the specified radio for an access point, set the value for this parameter to "true".

	AntennaPatternName string `json:"antennaPatternName,omitempty"` // Specify the antenna name on the specified radio for an access point. The antenna name is used to calculate the gain on the radio slot.

	AntennaGain *int `json:"antennaGain,omitempty"` // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi). To configure "antennaGain", set "antennaPatternName" value to "other".

	ConfigureAntennaCable *bool `json:"configureAntennaCable,omitempty"` // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".

	AntennaCableName string `json:"antennaCableName,omitempty"` // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".

	CableLoss *float64 `json:"cableLoss,omitempty"` // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).

	ConfigureChannel *bool `json:"configureChannel,omitempty"` // To change the channel on the specified radio for an access point, set this parameter's value to "true".

	ChannelAssignmentMode *int `json:"channelAssignmentMode,omitempty"` // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	ChannelNumber *int `json:"channelNumber,omitempty"` // Configure the channel number on the specified radio for an access point.

	ConfigureChannelWidth *bool `json:"configureChannelWidth,omitempty"` // To change the channel width on the specified radio for an access point, set this parameter's value to "true".

	ChannelWidth *int `json:"channelWidth,omitempty"` // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; for 160 MHz, set "6", and for 320 MHz, set "7".

	ConfigurePower *bool `json:"configurePower,omitempty"` // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".

	PowerAssignmentMode *int `json:"powerAssignmentMode,omitempty"` // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".

	Powerlevel *int `json:"powerlevel,omitempty"` // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.

	RadioType *int `json:"radioType,omitempty"` // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
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
/* This API allows the user to get count of all SSIDs (Service Set Identifier) .


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

@param id id path parameter. SSID ID


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

//GetAccessPointConfigurationCount Get Access Point Configuration Count - 118b-2898-457b-8d47
/* Get Access Point Configuration Count


@param GetAccessPointConfigurationCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-access-point-configuration-count
*/
func (s *WirelessService) GetAccessPointConfigurationCount(GetAccessPointConfigurationCountQueryParams *GetAccessPointConfigurationCountQueryParams) (*ResponseWirelessGetAccessPointConfigurationCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration/count"

	queryString, _ := query.Values(GetAccessPointConfigurationCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetAccessPointConfigurationCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAccessPointConfigurationCount(GetAccessPointConfigurationCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfigurationCount")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfigurationCount)
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

//GetAnchorCapableDevices Get Anchor capable devices - a581-0a06-4acb-8f4c
/* This API allows the user to get Anchor capable devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-capable-devices
*/
func (s *WirelessService) GetAnchorCapableDevices() (*ResponseWirelessGetAnchorCapableDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/anchorCapableDevices"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAnchorCapableDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorCapableDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorCapableDevices")
	}

	result := response.Result().(*ResponseWirelessGetAnchorCapableDevices)
	return result, response, err

}

//GetMeshApNeighbours Get Mesh Ap Neighbours - 8a88-98f8-4eca-8300
/* Retrieves all Mesh accesspoint Neighbours details whether child, parent, etc.


@param GetMeshApNeighboursQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mesh-ap-neighbours
*/
func (s *WirelessService) GetMeshApNeighbours(GetMeshApNeighboursQueryParams *GetMeshApNeighboursQueryParams) (*ResponseWirelessGetMeshApNeighbours, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/meshApNeighbours"

	queryString, _ := query.Values(GetMeshApNeighboursQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetMeshApNeighbours{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMeshApNeighbours(GetMeshApNeighboursQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMeshApNeighbours")
	}

	result := response.Result().(*ResponseWirelessGetMeshApNeighbours)
	return result, response, err

}

//GetMeshApNeighboursCount Get Mesh Ap Neighbours Count - 54b9-09f2-4dd8-b94f
/* This API returns the total number of mesh Ap Neighbours available.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mesh-ap-neighbours-count
*/
func (s *WirelessService) GetMeshApNeighboursCount() (*ResponseWirelessGetMeshApNeighboursCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/meshApNeighbours/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetMeshApNeighboursCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMeshApNeighboursCount()
		}
		return nil, response, fmt.Errorf("error with operation GetMeshApNeighboursCount")
	}

	result := response.Result().(*ResponseWirelessGetMeshApNeighboursCount)
	return result, response, err

}

//GetMobilityGroups Get MobilityGroups - 628f-38bf-4f5a-a48c
/* Retrieve configured mobility groups if no Network Device Id is provided as a query parameter. If a Network Device Id is given and a mobility group is configured for it, return the configured details; otherwise, return the default values from the device.


@param GetMobilityGroupsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-mobility-groups
*/
func (s *WirelessService) GetMobilityGroups(GetMobilityGroupsQueryParams *GetMobilityGroupsQueryParams) (*ResponseWirelessGetMobilityGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/wirelessMobilityGroups"

	queryString, _ := query.Values(GetMobilityGroupsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetMobilityGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetMobilityGroups(GetMobilityGroupsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetMobilityGroups")
	}

	result := response.Result().(*ResponseWirelessGetMobilityGroups)
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

//GetApAuthorizationListByNetworkDeviceID Get AP Authorization List by network device Id - c689-88c5-4128-a366
/* This API allows the user to get an AP Authorization List details configured for the given provisioned network device Id. Obtain the network device ID value by using the API GET call '/dna/intent/api/v1/network-device/ip-address/${ipAddress}'.


@param networkDeviceID networkDeviceId path parameter. Network Device ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-list-by-network-device-id
*/
func (s *WirelessService) GetApAuthorizationListByNetworkDeviceID(networkDeviceID string) (*ResponseWirelessGetApAuthorizationListByNetworkDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessControllers/{networkDeviceId}/apAuthorizationLists"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApAuthorizationListByNetworkDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListByNetworkDeviceID(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListByNetworkDeviceId")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListByNetworkDeviceID)
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
/* Retrieves the count of SSIDs associated with the specific wireless controller.


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

//RetrieveAllPolicyTagsForAWirelessProfile Retrieve all Policy Tags for a Wireless Profile - 428d-d8f7-4fa9-a0ca
/* This endpoint retrieves a list of all `Policy Tags` associated with a specific `Wireless Profile`. This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless Profile Id

@param RetrieveAllPolicyTagsForAWirelessProfileQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-all-policy-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveAllPolicyTagsForAWirelessProfile(id string, RetrieveAllPolicyTagsForAWirelessProfileQueryParams *RetrieveAllPolicyTagsForAWirelessProfileQueryParams) (*ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveAllPolicyTagsForAWirelessProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAllPolicyTagsForAWirelessProfile(id, RetrieveAllPolicyTagsForAWirelessProfileQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAllPolicyTagsForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessRetrieveAllPolicyTagsForAWirelessProfile)
	return result, response, err

}

//RetrieveTheCountOfPolicyTagsForAWirelessProfile Retrieve the count of Policy Tags for a Wireless Profile - 26be-7947-4a69-be88
/* This endpoint retrieves the total count of `Policy Tags` associated with a specific `Wireless Profile`.This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-policy-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveTheCountOfPolicyTagsForAWirelessProfile(id string) (*ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfPolicyTagsForAWirelessProfile(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfPolicyTagsForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessRetrieveTheCountOfPolicyTagsForAWirelessProfile)
	return result, response, err

}

//RetrieveASpecificPolicyTagForAWirelessProfile Retrieve a specific Policy Tag for a Wireless Profile - 558d-4b8c-4149-b325
/* This endpoint retrieves the details of a specific `Policy Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `policyTagId` of the `Policy Tag`.


@param id id path parameter. Wireless Profile Id

@param policyTagID policyTagId path parameter. Policy Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-specific-policy-tag-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveASpecificPolicyTagForAWirelessProfile(id string, policyTagID string) (*ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/{policyTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{policyTagId}", fmt.Sprintf("%v", policyTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveASpecificPolicyTagForAWirelessProfile(id, policyTagID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveASpecificPolicyTagForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessRetrieveASpecificPolicyTagForAWirelessProfile)
	return result, response, err

}

//RetrieveAllSiteTagsForAWirelessProfile Retrieve all Site Tags for a Wireless Profile - bd9d-5a4b-4f5a-ac11
/* This endpoint retrieves a list of all `Site Tags` associated with a specific `Wireless Profile`. This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless profile id

@param RetrieveAllSiteTagsForAWirelessProfileQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-all-site-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveAllSiteTagsForAWirelessProfile(id string, RetrieveAllSiteTagsForAWirelessProfileQueryParams *RetrieveAllSiteTagsForAWirelessProfileQueryParams) (*ResponseWirelessRetrieveAllSiteTagsForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrieveAllSiteTagsForAWirelessProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveAllSiteTagsForAWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAllSiteTagsForAWirelessProfile(id, RetrieveAllSiteTagsForAWirelessProfileQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAllSiteTagsForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessRetrieveAllSiteTagsForAWirelessProfile)
	return result, response, err

}

//RetrieveTheCountOfSiteTagsForAWirelessProfile Retrieve the count of Site Tags for a Wireless Profile - bdab-4896-4fa8-bbcc
/* This endpoint retrieves the total count of `Site Tags` associated with a specific `Wireless Profile`.This API requires the `id` of the `Wireless Profile` to be provided as a path parameter.


@param id id path parameter. Wireless profile id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-site-tags-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveTheCountOfSiteTagsForAWirelessProfile(id string) (*ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfSiteTagsForAWirelessProfile(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfSiteTagsForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfile)
	return result, response, err

}

//RetrieveASpecificSiteTagForAWirelessProfile Retrieve a specific Site Tag for a Wireless Profile - 1fac-e966-4e6b-96ea
/* This endpoint retrieves the details of a specific `Site Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `siteTagId` of the `Site Tag`.


@param id id path parameter. Wireless Profile Id

@param siteTagID siteTagId path parameter. Site Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-specific-site-tag-for-a-wireless-profile
*/
func (s *WirelessService) RetrieveASpecificSiteTagForAWirelessProfile(id string, siteTagID string) (*ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/{siteTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteTagId}", fmt.Sprintf("%v", siteTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveASpecificSiteTagForAWirelessProfile(id, siteTagID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveASpecificSiteTagForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfile)
	return result, response, err

}

//GetAnchorGroups Get AnchorGroups - 32b3-aa83-46db-aae7
/* This API allows the user to get AnchorGroups that captured in wireless settings design.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-groups
*/
func (s *WirelessService) GetAnchorGroups() (*ResponseWirelessGetAnchorGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAnchorGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorGroups()
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorGroups")
	}

	result := response.Result().(*ResponseWirelessGetAnchorGroups)
	return result, response, err

}

//GetCountOfAnchorGroups Get count of AnchorGroups - 5581-4892-4dc8-a66c
/* This API allows the user to get count of all AnchorGroups



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-anchor-groups
*/
func (s *WirelessService) GetCountOfAnchorGroups() (*ResponseWirelessGetCountOfAnchorGroups, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetCountOfAnchorGroups{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfAnchorGroups()
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfAnchorGroups")
	}

	result := response.Result().(*ResponseWirelessGetCountOfAnchorGroups)
	return result, response, err

}

//GetAnchorGroupByID Get AnchorGroup by ID - 029f-4acf-420b-9df6
/* This API allows the user to get an AnchorGroup by AnchorGroup ID


@param id id path parameter. AnchorGroup ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anchor-group-by-id
*/
func (s *WirelessService) GetAnchorGroupByID(id string) (*ResponseWirelessGetAnchorGroupByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetAnchorGroupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAnchorGroupByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetAnchorGroupById")
	}

	result := response.Result().(*ResponseWirelessGetAnchorGroupByID)
	return result, response, err

}

//GetApAuthorizationLists Get AP Authorization Lists - f7b1-b801-4738-937d
/* Retrieves the AP Authorization Lists that are created in the Catalyst Centre network Design for wireless. If an AP Authorization List name is given as query parameter, then returns respective AP Authorization List details including Local and/or Remote authorization.


@param GetAPAuthorizationListsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-lists
*/
func (s *WirelessService) GetApAuthorizationLists(GetAPAuthorizationListsQueryParams *GetApAuthorizationListsQueryParams) (*ResponseWirelessGetApAuthorizationLists, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists"

	queryString, _ := query.Values(GetAPAuthorizationListsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetApAuthorizationLists{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationLists(GetAPAuthorizationListsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationLists")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationLists)
	return result, response, err

}

//GetApAuthorizationListCount Get AP Authorization List Count - 51be-d862-47c8-a51a
/* This API allows the user to get count of all AP Authorization lists.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-list-count
*/
func (s *WirelessService) GetApAuthorizationListCount() (*ResponseWirelessGetApAuthorizationListCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApAuthorizationListCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListCount()
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListCount")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListCount)
	return result, response, err

}

//GetApAuthorizationListByID Get AP Authorization List by ID - 0aa3-1b69-4a58-b0f5
/* This API allows the user to get an AP Authorization List by AP Authorization List ID that captured in wireless settings design.


@param id id path parameter. AP Authorization List ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-authorization-list-by-id
*/
func (s *WirelessService) GetApAuthorizationListByID(id string) (*ResponseWirelessGetApAuthorizationListByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApAuthorizationListByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApAuthorizationListByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetApAuthorizationListById")
	}

	result := response.Result().(*ResponseWirelessGetApAuthorizationListByID)
	return result, response, err

}

//GetApProfiles Get AP Profiles - edad-9bfa-4298-a4cb
/* This API allows the user to get AP profiles that are captured in wireless settings design.


@param GetAPProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-profiles
*/
func (s *WirelessService) GetApProfiles(GetAPProfilesQueryParams *GetApProfilesQueryParams) (*ResponseWirelessGetApProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles"

	queryString, _ := query.Values(GetAPProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetApProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApProfiles(GetAPProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApProfiles")
	}

	result := response.Result().(*ResponseWirelessGetApProfiles)
	return result, response, err

}

//GetApProfilesCount Get AP Profiles Count - a687-f85e-438a-8941
/* This API returns the total number of AP Profiles available.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-profiles-count
*/
func (s *WirelessService) GetApProfilesCount() (*ResponseWirelessGetApProfilesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApProfilesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApProfilesCount()
		}
		return nil, response, fmt.Errorf("error with operation GetApProfilesCount")
	}

	result := response.Result().(*ResponseWirelessGetApProfilesCount)
	return result, response, err

}

//GetApProfileByID Get AP Profile by ID - ba9f-5899-4c5b-87f2
/* This API allows the user to get a AP Profile by AP Profile ID that captured in wireless settings design


@param id id path parameter. Ap Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ap-profile-by-id
*/
func (s *WirelessService) GetApProfileByID(id string) (*ResponseWirelessGetApProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetApProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetApProfileById")
	}

	result := response.Result().(*ResponseWirelessGetApProfileByID)
	return result, response, err

}

//Get80211BeProfiles Get 802.11be Profiles - 1895-aac1-4428-bd0d
/* This API allows the user to get 802.11be Profile(s) configured under Wireless Settings


@param Get80211beProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get80211be-profiles
*/
func (s *WirelessService) Get80211BeProfiles(Get80211beProfilesQueryParams *Get80211BeProfilesQueryParams) (*ResponseWirelessGet80211BeProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/dot11beProfiles"

	queryString, _ := query.Values(Get80211beProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGet80211BeProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Get80211BeProfiles(Get80211beProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Get80211BeProfiles")
	}

	result := response.Result().(*ResponseWirelessGet80211BeProfiles)
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

//GetPowerProfiles Get Power Profiles - a9b3-d8c2-4b6b-9b01
/* This API allows the user to get Power Profiles that captured in wireless settings design.


@param GetPowerProfilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-power-profiles
*/
func (s *WirelessService) GetPowerProfiles(GetPowerProfilesQueryParams *GetPowerProfilesQueryParams) (*ResponseWirelessGetPowerProfiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles"

	queryString, _ := query.Values(GetPowerProfilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessGetPowerProfiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPowerProfiles(GetPowerProfilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetPowerProfiles")
	}

	result := response.Result().(*ResponseWirelessGetPowerProfiles)
	return result, response, err

}

//GetPowerProfilesCount Get Power Profiles Count - 7091-1b6d-4849-ab2f
/* This API returns the total number of Power Profiles available.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-power-profiles-count
*/
func (s *WirelessService) GetPowerProfilesCount() (*ResponseWirelessGetPowerProfilesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetPowerProfilesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPowerProfilesCount()
		}
		return nil, response, fmt.Errorf("error with operation GetPowerProfilesCount")
	}

	result := response.Result().(*ResponseWirelessGetPowerProfilesCount)
	return result, response, err

}

//GetPowerProfileByID Get Power Profile by ID - 6c93-cb96-45b8-a53b
/* This API allows the user to get a Power Profile by Power Profile ID that captured in wireless settings design


@param id id path parameter. Power Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-power-profile-by-id
*/
func (s *WirelessService) GetPowerProfileByID(id string) (*ResponseWirelessGetPowerProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessGetPowerProfileByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetPowerProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetPowerProfileById")
	}

	result := response.Result().(*ResponseWirelessGetPowerProfileByID)
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

//RetrieveSitesWithOverriddenSSIDs Retrieve sites with overridden SSIDs - 9a9a-8b8b-4029-a86e
/* Retrieve list of siteId(s) with information of SSID(s) which are overridden


@param RetrieveSitesWithOverriddenSSIDsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-sites-with-overridden-ssids
*/
func (s *WirelessService) RetrieveSitesWithOverriddenSSIDs(RetrieveSitesWithOverriddenSSIDsQueryParams *RetrieveSitesWithOverriddenSSIDsQueryParams) (*ResponseWirelessRetrieveSitesWithOverriddenSSIDs, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/ssids/overrideAtSites"

	queryString, _ := query.Values(RetrieveSitesWithOverriddenSSIDsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWirelessRetrieveSitesWithOverriddenSSIDs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveSitesWithOverriddenSSIDs(RetrieveSitesWithOverriddenSSIDsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveSitesWithOverriddenSsids")
	}

	result := response.Result().(*ResponseWirelessRetrieveSitesWithOverriddenSSIDs)
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

//UpdateOrOverridessid Update or Override SSID - 559d-88ff-43c9-9fe3
/* This API allows to either update SSID at global 'siteId' or override SSID at given non-global 'siteId'


@param siteID siteId path parameter. Site UUID

@param id id path parameter. SSID ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!update-or-overridessid
*/
func (s *WirelessService) UpdateOrOverridessid(siteID string, id string, requestWirelessUpdateOrOverrideSSID *RequestWirelessUpdateOrOverridessid) (*ResponseWirelessUpdateOrOverridessid, *resty.Response, error) {
	path := "/dna/intent/api/v1/sites/{siteId}/wirelessSettings/ssids/{id}/update"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateOrOverrideSSID).
		SetResult(&ResponseWirelessUpdateOrOverridessid{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateOrOverridessid(siteID, id, requestWirelessUpdateOrOverrideSSID)
		}

		return nil, response, fmt.Errorf("error with operation UpdateOrOverridessid")
	}

	result := response.Result().(*ResponseWirelessUpdateOrOverridessid)
	return result, response, err

}

//ConfigureAccessPoints Configure Access Points  - 0081-cb89-4708-888f
/* User can configure multiple access points with required options using this intent API. This API does not support configuration of CleanAir or SI for IOS-XE devices with version greater than or equal to 17.9



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-access-points
*/
func (s *WirelessService) ConfigureAccessPoints(requestWirelessConfigureAccessPoints *RequestWirelessConfigureAccessPoints) (*ResponseWirelessConfigureAccessPoints, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/accesspoint-configuration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessConfigureAccessPoints).
		SetResult(&ResponseWirelessConfigureAccessPoints{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAccessPoints(requestWirelessConfigureAccessPoints)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAccessPoints")
	}

	result := response.Result().(*ResponseWirelessConfigureAccessPoints)
	return result, response, err

}

//ApProvisionConnectivity AP Provision - d897-19b8-47aa-a9c4
/* Access Point Provision and ReProvision


@param APProvisionConnectivityHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!ap-provision-connectivity
*/
func (s *WirelessService) ApProvisionConnectivity(requestWirelessAPProvisionConnectivity *RequestWirelessApProvisionConnectivity, APProvisionConnectivityHeaderParams *ApProvisionConnectivityHeaderParams) (*ResponseWirelessApProvisionConnectivity, *resty.Response, error) {
	path := "/dna/intent/api/v1/wireless/ap-provision"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if APProvisionConnectivityHeaderParams != nil {

		if APProvisionConnectivityHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", APProvisionConnectivityHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestWirelessAPProvisionConnectivity).
		SetResult(&ResponseWirelessApProvisionConnectivity{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApProvisionConnectivity(requestWirelessAPProvisionConnectivity, APProvisionConnectivityHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation ApProvisionConnectivity")
	}

	result := response.Result().(*ResponseWirelessApProvisionConnectivity)
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
/* Creates Wireless Network Profile on Cisco DNA Center and associates sites and SSIDs to it.



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
/* This API is used to provision Access Points. Prerequisite: Access Point has to be assigned to the site using the API /dna/intent/api/v1/networkDevices/assignToSite/apply.



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
/* This API is used to reset wireless mobility which in turn sets mobility group name as 'default'.



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
/* This API allows user to assign Managed AP Locations for IOS-XE Wireless supported devices by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.


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

//CreateWirelessProfileConnectivity Create Wireless Profile - dd88-bb37-492a-888b
/* This API allows the user to create a Wireless Network Profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-wireless-profile-connectivity
*/
func (s *WirelessService) CreateWirelessProfileConnectivity(requestWirelessCreateWirelessProfileConnectivity *RequestWirelessCreateWirelessProfileConnectivity) (*ResponseWirelessCreateWirelessProfileConnectivity, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateWirelessProfileConnectivity).
		SetResult(&ResponseWirelessCreateWirelessProfileConnectivity{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateWirelessProfileConnectivity(requestWirelessCreateWirelessProfileConnectivity)
		}

		return nil, response, fmt.Errorf("error with operation CreateWirelessProfileConnectivity")
	}

	result := response.Result().(*ResponseWirelessCreateWirelessProfileConnectivity)
	return result, response, err

}

//CreateMultiplePolicyTagsForAWirelessProfileInBulk Create multiple Policy Tags for a Wireless Profile in bulk - 6bbe-ca2b-430a-8665
/* This endpoint allows the creation of multiple `Policy Tags` associated with a specific `Wireless Profile` in a single request. The `id` of the Wireless Profile must be provided as a path parameter, and a list of `Policy Tags` should be included in the request body. Note: Multiple Policy Tags (policyTag) can be configured for the same siteId only if they have different sets of AP Zones (apZones). If multiple Policy Tags are created with the same apZones for the same site or a parent site, only the last one will be saved, overriding the previous ones.


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-multiple-policy-tags-for-a-wireless-profile-in-bulk
*/
func (s *WirelessService) CreateMultiplePolicyTagsForAWirelessProfileInBulk(id string, requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk *RequestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk) (*ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/bulk"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk).
		SetResult(&ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateMultiplePolicyTagsForAWirelessProfileInBulk(id, requestWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk)
		}

		return nil, response, fmt.Errorf("error with operation CreateMultiplePolicyTagsForAWirelessProfileInBulk")
	}

	result := response.Result().(*ResponseWirelessCreateMultiplePolicyTagsForAWirelessProfileInBulk)
	return result, response, err

}

//CreateMultipleSiteTagsForAWirelessProfileInBulk Create multiple Site Tags for a Wireless Profile in bulk - 7094-e87b-4b2b-9617
/* This endpoint allows the creation of multiple `Site Tags` associated with a specific `Wireless Profile` in a single request. The `id` of the `Wireless Profile` must be provided as a path parameter, and a list of `Site Tags` should be included in the request body. Note: Only one Site Tag (siteTag) can be created per siteId. If multiple siteTags are specified for the same siteId within a request, only the last one will be saved, overriding any previously configured tags. When creating a Site Tag under a Flex-enabled Wireless Profile (i.e., a Wireless Profile with one or more Flex SSIDs), a non-default Flex Profile Name (flexProfileName) will be used. If no custom flexProfileName is defined, the System will automatically generate one and configure it in the controller.


@param id id path parameter. network profile id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-multiple-site-tags-for-a-wireless-profile-in-bulk
*/
func (s *WirelessService) CreateMultipleSiteTagsForAWirelessProfileInBulk(id string, requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk *RequestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk) (*ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/bulk"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk).
		SetResult(&ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateMultipleSiteTagsForAWirelessProfileInBulk(id, requestWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk)
		}

		return nil, response, fmt.Errorf("error with operation CreateMultipleSiteTagsForAWirelessProfileInBulk")
	}

	result := response.Result().(*ResponseWirelessCreateMultipleSiteTagsForAWirelessProfileInBulk)
	return result, response, err

}

//CreateAnchorGroup Create AnchorGroup - 3d85-68e5-4909-988a
/* This API allows the user to create an AnchorGroup



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-anchor-group
*/
func (s *WirelessService) CreateAnchorGroup(requestWirelessCreateAnchorGroup *RequestWirelessCreateAnchorGroup) (*ResponseWirelessCreateAnchorGroup, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateAnchorGroup).
		SetResult(&ResponseWirelessCreateAnchorGroup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAnchorGroup(requestWirelessCreateAnchorGroup)
		}

		return nil, response, fmt.Errorf("error with operation CreateAnchorGroup")
	}

	result := response.Result().(*ResponseWirelessCreateAnchorGroup)
	return result, response, err

}

//CreateApAuthorizationList Create AP Authorization List - 5e9a-4806-489a-91db
/* This API allows the user to create an AP Authorization List.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-ap-authorization-list
*/
func (s *WirelessService) CreateApAuthorizationList(requestWirelessCreateAPAuthorizationList *RequestWirelessCreateApAuthorizationList) (*ResponseWirelessCreateApAuthorizationList, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateAPAuthorizationList).
		SetResult(&ResponseWirelessCreateApAuthorizationList{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApAuthorizationList(requestWirelessCreateAPAuthorizationList)
		}

		return nil, response, fmt.Errorf("error with operation CreateApAuthorizationList")
	}

	result := response.Result().(*ResponseWirelessCreateApAuthorizationList)
	return result, response, err

}

//CreateApProfile Create AP Profile - 3697-68d5-4149-9f02
/* This API allows the user to create a custom AP Profile.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-ap-profile
*/
func (s *WirelessService) CreateApProfile(requestWirelessCreateAPProfile *RequestWirelessCreateApProfile) (*ResponseWirelessCreateApProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreateAPProfile).
		SetResult(&ResponseWirelessCreateApProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApProfile(requestWirelessCreateAPProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateApProfile")
	}

	result := response.Result().(*ResponseWirelessCreateApProfile)
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

//CreatePowerProfile Create Power Profile - 7bac-6bd5-4269-8a3c
/* This API allows the user to create a custom Power Profile.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-power-profile
*/
func (s *WirelessService) CreatePowerProfile(requestWirelessCreatePowerProfile *RequestWirelessCreatePowerProfile) (*ResponseWirelessCreatePowerProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessCreatePowerProfile).
		SetResult(&ResponseWirelessCreatePowerProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatePowerProfile(requestWirelessCreatePowerProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreatePowerProfile")
	}

	result := response.Result().(*ResponseWirelessCreatePowerProfile)
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

//AssignAnchorManagedApLocationsForWLC Assign Anchor Managed AP Locations For WLC - 55af-697b-4e28-8167
/* This API allows user to assign Anchor Managed AP Locations for WLC by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.

       When anchor managed location array present then it will add the anchor managed locations.


@param networkDeviceID networkDeviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-anchor-managed-ap-locations-for-w-l-c
*/
func (s *WirelessService) AssignAnchorManagedApLocationsForWLC(networkDeviceID string, requestWirelessAssignAnchorManagedAPLocationsForWLC *RequestWirelessAssignAnchorManagedApLocationsForWLC) (*ResponseWirelessAssignAnchorManagedApLocationsForWLC, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/{networkDeviceId}/assignAnchorManagedApLocations"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessAssignAnchorManagedAPLocationsForWLC).
		SetResult(&ResponseWirelessAssignAnchorManagedApLocationsForWLC{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignAnchorManagedApLocationsForWLC(networkDeviceID, requestWirelessAssignAnchorManagedAPLocationsForWLC)
		}

		return nil, response, fmt.Errorf("error with operation AssignAnchorManagedApLocationsForWLC")
	}

	result := response.Result().(*ResponseWirelessAssignAnchorManagedApLocationsForWLC)
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

@param id id path parameter. SSID ID

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

//UpdateWirelessProfileConnectivity Update Wireless Profile - 4f88-d9a3-4ef8-8e2e
/* This API allows the user to update a Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id

*/
func (s *WirelessService) UpdateWirelessProfileConnectivity(id string, requestWirelessUpdateWirelessProfileConnectivity *RequestWirelessUpdateWirelessProfileConnectivity) (*ResponseWirelessUpdateWirelessProfileConnectivity, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateWirelessProfileConnectivity).
		SetResult(&ResponseWirelessUpdateWirelessProfileConnectivity{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateWirelessProfileConnectivity(id, requestWirelessUpdateWirelessProfileConnectivity)
		}
		return nil, response, fmt.Errorf("error with operation UpdateWirelessProfileConnectivity")
	}

	result := response.Result().(*ResponseWirelessUpdateWirelessProfileConnectivity)
	return result, response, err

}

//UpdateASpecificPolicyTagForAWirelessProfile Update a specific Policy Tag for a Wireless Profile - eca1-caf0-49bb-aa9a
/* This endpoint allows updating the details of a specific `Policy Tag` associated with a given `Wireless Profile`. The `id` of the `Wireless Profile` and the `policyTagId` of the Policy Tag must be provided as path parameters, and the request body should contain the updated details of the `Policy Tag`. The `policyTagName` cannot be modified through this endpoint. Note: When updating a Policy Tag, if the same set of AP Zones (apZones) is used for the same site or its parent site, the existing Policy Tag will be overridden by the new one.


@param id id path parameter. Wireless Profile Id

@param policyTagID policyTagId path parameter. Policy Tag Id

*/
func (s *WirelessService) UpdateASpecificPolicyTagForAWirelessProfile(id string, policyTagID string, requestWirelessUpdateASpecificPolicyTagForAWirelessProfile *RequestWirelessUpdateASpecificPolicyTagForAWirelessProfile) (*ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/{policyTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{policyTagId}", fmt.Sprintf("%v", policyTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateASpecificPolicyTagForAWirelessProfile).
		SetResult(&ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateASpecificPolicyTagForAWirelessProfile(id, policyTagID, requestWirelessUpdateASpecificPolicyTagForAWirelessProfile)
		}
		return nil, response, fmt.Errorf("error with operation UpdateASpecificPolicyTagForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessUpdateASpecificPolicyTagForAWirelessProfile)
	return result, response, err

}

//UpdateASpecificSiteTagForAWirelessProfile Update a specific Site Tag for a Wireless Profile - 47b2-1a6c-4158-9527
/* This endpoint allows updating the details of a specific `Site Tag` associated with a given `Wireless Profile`. The `id` of the `Wireless Profile` and the `siteTagId` of the Site Tag must be provided as path parameters, and the request body should contain the updated `Site Tag` details.  The `siteTagName` cannot be modified through this endpoint. Note: When updating a Site Tag (siteTag), if the siteId already has an associated siteTag and the same siteId is included in the update request, the existing siteTag for that siteId will be overridden by the new one. For Flex-enabled Wireless Profiles (i.e., a Wireless Profile with one or more Flex SSIDs), a non-default Flex Profile Name (flexProfileName) will be used. If no custom flexProfileName is provided, the System will automatically generate one and configure it in the controller.


@param id id path parameter. Wireless Profile Id

@param siteTagID siteTagId path parameter. Site Tag Id

*/
func (s *WirelessService) UpdateASpecificSiteTagForAWirelessProfile(id string, siteTagID string, requestWirelessUpdateASpecificSiteTagForAWirelessProfile *RequestWirelessUpdateASpecificSiteTagForAWirelessProfile) (*ResponseWirelessUpdateASpecificSiteTagForAWirelessProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/{siteTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteTagId}", fmt.Sprintf("%v", siteTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateASpecificSiteTagForAWirelessProfile).
		SetResult(&ResponseWirelessUpdateASpecificSiteTagForAWirelessProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateASpecificSiteTagForAWirelessProfile(id, siteTagID, requestWirelessUpdateASpecificSiteTagForAWirelessProfile)
		}
		return nil, response, fmt.Errorf("error with operation UpdateASpecificSiteTagForAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessUpdateASpecificSiteTagForAWirelessProfile)
	return result, response, err

}

//UpdateAnchorGroup Update AnchorGroup - 4da3-bae4-484a-8448
/* This API allows the user to update an AnchorGroup


@param id id path parameter. AnchorGroup ID

*/
func (s *WirelessService) UpdateAnchorGroup(id string, requestWirelessUpdateAnchorGroup *RequestWirelessUpdateAnchorGroup) (*ResponseWirelessUpdateAnchorGroup, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateAnchorGroup).
		SetResult(&ResponseWirelessUpdateAnchorGroup{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAnchorGroup(id, requestWirelessUpdateAnchorGroup)
		}
		return nil, response, fmt.Errorf("error with operation UpdateAnchorGroup")
	}

	result := response.Result().(*ResponseWirelessUpdateAnchorGroup)
	return result, response, err

}

//UpdateApAuthorizationList Update AP Authorization List - 768b-0b99-4189-986e
/* This API allows the user to update an AP Authorization List.


@param id id path parameter. AP Authorization List ID

*/
func (s *WirelessService) UpdateApAuthorizationList(id string, requestWirelessUpdateAPAuthorizationList *RequestWirelessUpdateApAuthorizationList) (*ResponseWirelessUpdateApAuthorizationList, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateAPAuthorizationList).
		SetResult(&ResponseWirelessUpdateApAuthorizationList{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateApAuthorizationList(id, requestWirelessUpdateAPAuthorizationList)
		}
		return nil, response, fmt.Errorf("error with operation UpdateApAuthorizationList")
	}

	result := response.Result().(*ResponseWirelessUpdateApAuthorizationList)
	return result, response, err

}

//UpdateApProfileByID Update AP Profile by ID - 41b5-ea3b-43b8-9d57
/* This API allows the user to update a custom AP Profile


@param id id path parameter. Ap Profile ID

*/
func (s *WirelessService) UpdateApProfileByID(id string, requestWirelessUpdateAPProfileByID *RequestWirelessUpdateApProfileByID) (*ResponseWirelessUpdateApProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdateAPProfileByID).
		SetResult(&ResponseWirelessUpdateApProfileByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateApProfileByID(id, requestWirelessUpdateAPProfileByID)
		}
		return nil, response, fmt.Errorf("error with operation UpdateApProfileById")
	}

	result := response.Result().(*ResponseWirelessUpdateApProfileByID)
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

//UpdatePowerProfileByID Update Power Profile by ID - 6f86-19b1-478a-9a5e
/* This API allows the user to update a custom power Profile


@param id id path parameter. Power Profile Id

*/
func (s *WirelessService) UpdatePowerProfileByID(id string, requestWirelessUpdatePowerProfileByID *RequestWirelessUpdatePowerProfileByID) (*ResponseWirelessUpdatePowerProfileByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWirelessUpdatePowerProfileByID).
		SetResult(&ResponseWirelessUpdatePowerProfileByID{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdatePowerProfileByID(id, requestWirelessUpdatePowerProfileByID)
		}
		return nil, response, fmt.Errorf("error with operation UpdatePowerProfileById")
	}

	result := response.Result().(*ResponseWirelessUpdatePowerProfileByID)
	return result, response, err

}

//UpdateRfProfile Update RF Profile - 2984-b995-4ae9-b3c3
/* This API allows the user to update a custom RF Profile.


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
/* Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from DNA Center


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
/* This API allows the user to delete an SSID (Service Set Identifier) at the global level , if the SSID is not mapped to any Wireless Profile, Or remove override from given site Id .


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

//DeleteWirelessProfileConnectivity Delete Wireless Profile - 289c-f9f5-4f78-b84c
/* This API allows the user to delete Wireless Network Profile by ID


@param id id path parameter. Wireless Profile Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-wireless-profile-connectivity
*/
func (s *WirelessService) DeleteWirelessProfileConnectivity(id string) (*ResponseWirelessDeleteWirelessProfileConnectivity, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteWirelessProfileConnectivity{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteWirelessProfileConnectivity(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteWirelessProfileConnectivity")
	}

	result := response.Result().(*ResponseWirelessDeleteWirelessProfileConnectivity)
	return result, response, err

}

//DeleteASpecificPolicyTagFromAWirelessProfile Delete a specific Policy Tag from a Wireless Profile - 4da4-fa50-4b89-a098
/* This endpoint allows for the deletion of a specific `Policy Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `policyTagId` of the `Policy Tag` to be provided as path parameters.


@param id id path parameter. Wireless Profile Id

@param policyTagID policyTagId path parameter. Policy Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-specific-policy-tag-from-a-wireless-profile
*/
func (s *WirelessService) DeleteASpecificPolicyTagFromAWirelessProfile(id string, policyTagID string) (*ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfile, *resty.Response, error) {
	//id string,policyTagID string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/policyTags/{policyTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{policyTagId}", fmt.Sprintf("%v", policyTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteASpecificPolicyTagFromAWirelessProfile(id, policyTagID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteASpecificPolicyTagFromAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessDeleteASpecificPolicyTagFromAWirelessProfile)
	return result, response, err

}

//DeleteASpecificSiteTagFromAWirelessProfile Delete a specific Site Tag from a Wireless Profile - e090-8a9d-4e2b-858f
/* This endpoint enables the deletion of a specific `Site Tag` associated with a given `Wireless Profile`. This API requires the `id` of the `Wireless Profile` and the `siteTagId` of the `Site Tag` to be provided as path parameters.


@param id id path parameter. Wireless Profile id

@param siteTagID siteTagId path parameter. Site Tag Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-a-specific-site-tag-from-a-wireless-profile
*/
func (s *WirelessService) DeleteASpecificSiteTagFromAWirelessProfile(id string, siteTagID string) (*ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfile, *resty.Response, error) {
	//id string,siteTagID string
	path := "/dna/intent/api/v1/wirelessProfiles/{id}/siteTags/{siteTagId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteTagId}", fmt.Sprintf("%v", siteTagID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteASpecificSiteTagFromAWirelessProfile(id, siteTagID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteASpecificSiteTagFromAWirelessProfile")
	}

	result := response.Result().(*ResponseWirelessDeleteASpecificSiteTagFromAWirelessProfile)
	return result, response, err

}

//DeleteAnchorGroupByID Delete AnchorGroup by ID - 6d83-eb95-46eb-9276
/* This API allows the user to delete an AnchorGroup  by specifying the AnchorGroup ID


@param id id path parameter. AnchorGroup ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-anchor-group-by-id
*/
func (s *WirelessService) DeleteAnchorGroupByID(id string) (*ResponseWirelessDeleteAnchorGroupByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/anchorGroups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteAnchorGroupByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAnchorGroupByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteAnchorGroupById")
	}

	result := response.Result().(*ResponseWirelessDeleteAnchorGroupByID)
	return result, response, err

}

//DeleteApAuthorizationList Delete AP Authorization List - 96a5-e901-47a8-8a31
/* This API allows the user to delete an AP Authorization List.


@param id id path parameter. AP Authorization List ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ap-authorization-list
*/
func (s *WirelessService) DeleteApAuthorizationList(id string) (*ResponseWirelessDeleteApAuthorizationList, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/apAuthorizationLists/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteApAuthorizationList{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApAuthorizationList(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApAuthorizationList")
	}

	result := response.Result().(*ResponseWirelessDeleteApAuthorizationList)
	return result, response, err

}

//DeleteApProfileByID Delete AP Profile by ID - 5b95-78eb-452b-856b
/* This API allows the user to delete an AP Profile by specifying the AP Profile ID.


@param id id path parameter. AP Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-ap-profile-by-id
*/
func (s *WirelessService) DeleteApProfileByID(id string) (*ResponseWirelessDeleteApProfileByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/apProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeleteApProfileByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApProfileById")
	}

	result := response.Result().(*ResponseWirelessDeleteApProfileByID)
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

//DeletePowerProfileByID Delete Power Profile by ID - 0eb7-faa1-41a9-9490
/* This API allows the user to delete an Power Profile by specifying the Power Profile ID.


@param id id path parameter. Power Profile ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-power-profile-by-id
*/
func (s *WirelessService) DeletePowerProfileByID(id string) (*ResponseWirelessDeletePowerProfileByID, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/wirelessSettings/powerProfiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWirelessDeletePowerProfileByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletePowerProfileByID(id)
		}
		return nil, response, fmt.Errorf("error with operation DeletePowerProfileById")
	}

	result := response.Result().(*ResponseWirelessDeletePowerProfileByID)
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
