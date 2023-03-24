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
type GetAccessPointRebootTaskResultQueryParams struct {
	ParentTaskID string `url:"parentTaskId,omitempty"` //task id of ap reboot request
}
type GetEnterpriseSSIDQueryParams struct {
	SSIDName string `url:"ssidName,omitempty"` //Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
}
type GetAccessPointConfigurationQueryParams struct {
	Key string `url:"key,omitempty"` //The ethernet MAC address of Access point
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
	RfProfileName string `url:"rf-profile-name,omitempty"` //RF Profile Name
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
	Name                        string                                                              `json:"name,omitempty"`                        // SSID Name
	WLANType                    string                                                              `json:"wlanType,omitempty"`                    // Wlan Type
	EnableFastLane              *bool                                                               `json:"enableFastLane,omitempty"`              // Enable Fast Lane
	SecurityLevel               string                                                              `json:"securityLevel,omitempty"`               // Security Level
	AuthServer                  string                                                              `json:"authServer,omitempty"`                  // Auth Server
	Passphrase                  string                                                              `json:"passphrase,omitempty"`                  // Passphrase
	TrafficType                 string                                                              `json:"trafficType,omitempty"`                 // Traffic Type
	EnableMacFiltering          *bool                                                               `json:"enableMACFiltering,omitempty"`          // Enable MAC Filtering
	IsEnabled                   *bool                                                               `json:"isEnabled,omitempty"`                   // Is Enabled
	IsFabric                    *bool                                                               `json:"isFabric,omitempty"`                    // Is Fabric
	FastTransition              string                                                              `json:"fastTransition,omitempty"`              // Fast Transition
	RadioPolicy                 string                                                              `json:"radioPolicy,omitempty"`                 // Radio Policy
	EnableBroadcastSSID         *bool                                                               `json:"enableBroadcastSSID,omitempty"`         // Enable Broadcast SSID
	NasOptions                  []string                                                            `json:"nasOptions,omitempty"`                  // Nas Options
	AAAOverride                 *bool                                                               `json:"aaaOverride,omitempty"`                 // Aaa Override
	CoverageHoleDetectionEnable *bool                                                               `json:"coverageHoleDetectionEnable,omitempty"` // Coverage Hole Detection Enable
	ProtectedManagementFrame    string                                                              `json:"protectedManagementFrame,omitempty"`    // Protected Management Frame
	MultipSKSettings            *[]ResponseItemWirelessGetEnterpriseSSIDSSIDDetailsMultipSKSettings `json:"multiPSKSettings,omitempty"`            //
	ClientRateLimit             *float64                                                            `json:"clientRateLimit,omitempty"`             // Client Rate Limit. (in bits per second)
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
type ResponseWirelessDeleteWirelessProfile struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseWirelessConfigureAccessPoints struct {
	Response *ResponseWirelessConfigureAccessPointsResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponseWirelessConfigureAccessPointsResponse struct {
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
type ResponseWirelessApProvision []ResponseItemWirelessApProvision // Array of ResponseWirelessAPProvision
type ResponseItemWirelessApProvision struct {
	ExecutionID  string `json:"executionId,omitempty"`  // Execution Id
	ExecutionURL string `json:"executionUrl,omitempty"` // Execution URL
	Message      string `json:"message,omitempty"`      // Response
}
type ResponseWirelessCreateUpdateDynamicInterface []ResponseItemWirelessCreateUpdateDynamicInterface // Array of ResponseWirelessCreateUpdateDynamicInterface
type ResponseItemWirelessCreateUpdateDynamicInterface struct {
	ExecutionID  string `json:"executionId,omitempty"`  // Execution Id
	ExecutionURL string `json:"executionUrl,omitempty"` // Execution URL
	Message      string `json:"message,omitempty"`      // Response
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
	Name                string `json:"name,omitempty"`                // Name
	ParentProfileA      string `json:"parentProfileA,omitempty"`      // Parent Profile A
	ParentProfileB      string `json:"parentProfileB,omitempty"`      // Parent Profile B
	EnableARadioType    *bool  `json:"enableARadioType,omitempty"`    // Enable ARadio Type
	EnableBRadioType    *bool  `json:"enableBRadioType,omitempty"`    // Enable BRadio Type
	EnableCRadioType    *bool  `json:"enableCRadioType,omitempty"`    // Enable CRadio Type
	ChannelWidth        string `json:"channelWidth,omitempty"`        // Channel Width
	ARadioChannels      string `json:"aRadioChannels,omitempty"`      // A Radio Channels
	BRadioChannels      string `json:"bRadioChannels,omitempty"`      // B Radio Channels
	CRadioChannels      string `json:"cRadioChannels,omitempty"`      // C Radio Channels
	DataRatesA          string `json:"dataRatesA,omitempty"`          // Data Rates A
	DataRatesB          string `json:"dataRatesB,omitempty"`          // Data Rates B
	DataRatesC          string `json:"dataRatesC,omitempty"`          // Data Rates C
	MandatoryDataRatesA string `json:"mandatoryDataRatesA,omitempty"` // Mandatory Data Rates A
	MandatoryDataRatesB string `json:"mandatoryDataRatesB,omitempty"` // Mandatory Data Rates B
	MandatoryDataRatesC string `json:"mandatoryDataRatesC,omitempty"` // Mandatory Data Rates C
	EnableCustom        *bool  `json:"enableCustom,omitempty"`        // Enable Custom
	MinPowerLevelA      string `json:"minPowerLevelA,omitempty"`      // Min Power Level A
	MinPowerLevelB      string `json:"minPowerLevelB,omitempty"`      // Min Power Level B
	MinPowerLevelC      string `json:"minPowerLevelC,omitempty"`      // Min Power Level C
	MaxPowerLevelA      string `json:"maxPowerLevelA,omitempty"`      // Max Power Level A
	MaxPowerLevelB      string `json:"maxPowerLevelB,omitempty"`      // Max Power Level B
	PowerThresholdV1A   *int   `json:"powerThresholdV1A,omitempty"`   // Power Threshold V1 A
	PowerThresholdV1B   *int   `json:"powerThresholdV1B,omitempty"`   // Power Threshold V1 B
	PowerThresholdV1C   *int   `json:"powerThresholdV1C,omitempty"`   // Power Threshold V1 C
	RxSopThresholdA     string `json:"rxSopThresholdA,omitempty"`     // Rx Sop Threshold A
	RxSopThresholdB     string `json:"rxSopThresholdB,omitempty"`     // Rx Sop Threshold B
	RxSopThresholdC     string `json:"rxSopThresholdC,omitempty"`     // Rx Sop Threshold C
	DefaultRfProfile    *bool  `json:"defaultRfProfile,omitempty"`    // Default Rf Profile
	EnableBrownField    *bool  `json:"enableBrownField,omitempty"`    // Enable Brown Field
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
	RadioPolicy         string `json:"radioPolicy,omitempty"`         // Radio Policy
	EnableMacFiltering  *bool  `json:"enableMACFiltering,omitempty"`  // Enable MAC Filtering
	FastTransition      string `json:"fastTransition,omitempty"`      // Fast Transition
	WebAuthURL          string `json:"webAuthURL,omitempty"`          // Web Auth URL
}
type RequestWirelessCreateAndProvisionSSIDFlexConnect struct {
	EnableFlexConnect *bool `json:"enableFlexConnect,omitempty"` // Enable Flex Connect
	LocalToVLAN       *int  `json:"localToVlan,omitempty"`       // Local To Vlan (range is 1 to 4094)
}
type RequestWirelessRebootAccessPoints struct {
	ApMacAddresses []string `json:"apMacAddresses,omitempty"` // The ethernet MAC address of the access point.
}
type RequestWirelessCreateEnterpriseSSID struct {
	Name                             string   `json:"name,omitempty"`                             // SSID NAME
	SecurityLevel                    string   `json:"securityLevel,omitempty"`                    // Security Level
	Passphrase                       string   `json:"passphrase,omitempty"`                       // Passphrase
	EnableFastLane                   *bool    `json:"enableFastLane,omitempty"`                   // Enable FastLane
	EnableMacFiltering               *bool    `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	TrafficType                      string   `json:"trafficType,omitempty"`                      // Traffic Type Enum (voicedata or data )
	RadioPolicy                      string   `json:"radioPolicy,omitempty"`                      // Radio Policy Enum (enum: Triple band operation (2.4GHz, 5GHz and 6GHz), Triple band operation with band select, 5GHz only, 2.4GHz only, 6GHz only)
	EnableBroadcastSSID              *bool    `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcase SSID
	FastTransition                   string   `json:"fastTransition,omitempty"`                   // Fast Transition
	EnableSessionTimeOut             *bool    `json:"enableSessionTimeOut,omitempty"`             // Enable Session Timeout
	SessionTimeOut                   *int     `json:"sessionTimeOut,omitempty"`                   // Session Time Out
	EnableClientExclusion            *bool    `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *int     `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool    `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *int     `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout
	EnableDirectedMulticastService   *bool    `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed Multicast Service
	EnableNeighborList               *bool    `json:"enableNeighborList,omitempty"`               // Enable Neighbor List
	MfpClientProtection              string   `json:"mfpClientProtection,omitempty"`              // Management Frame Protection Client
	NasOptions                       []string `json:"nasOptions,omitempty"`                       // Nas Options
}
type RequestWirelessUpdateEnterpriseSSID struct {
	Name                             string   `json:"name,omitempty"`                             // SSID NAME
	SecurityLevel                    string   `json:"securityLevel,omitempty"`                    // Security Level
	Passphrase                       string   `json:"passphrase,omitempty"`                       // Passphrase
	EnableFastLane                   *bool    `json:"enableFastLane,omitempty"`                   // Enable FastLane
	EnableMacFiltering               *bool    `json:"enableMACFiltering,omitempty"`               // Enable MAC Filtering
	TrafficType                      string   `json:"trafficType,omitempty"`                      // Traffic Type Enum (voicedata or data )
	RadioPolicy                      string   `json:"radioPolicy,omitempty"`                      // Radio Policy Enum (enum: Triple band operation (2.4GHz, 5GHz and 6GHz), Triple band operation with band select, 5GHz only, 2.4GHz only, 6GHz only)
	EnableBroadcastSSID              *bool    `json:"enableBroadcastSSID,omitempty"`              // Enable Broadcase SSID
	FastTransition                   string   `json:"fastTransition,omitempty"`                   // Fast Transition
	EnableSessionTimeOut             *bool    `json:"enableSessionTimeOut,omitempty"`             // Enable Session Timeout
	SessionTimeOut                   *int     `json:"sessionTimeOut,omitempty"`                   // Session Time Out
	EnableClientExclusion            *bool    `json:"enableClientExclusion,omitempty"`            // Enable Client Exclusion
	ClientExclusionTimeout           *int     `json:"clientExclusionTimeout,omitempty"`           // Client Exclusion Timeout
	EnableBasicServiceSetMaxIDle     *bool    `json:"enableBasicServiceSetMaxIdle,omitempty"`     // Enable Basic Service Set Max Idle
	BasicServiceSetClientIDleTimeout *int     `json:"basicServiceSetClientIdleTimeout,omitempty"` // Basic Service Set Client Idle Timeout
	EnableDirectedMulticastService   *bool    `json:"enableDirectedMulticastService,omitempty"`   // Enable Directed Multicast Service
	EnableNeighborList               *bool    `json:"enableNeighborList,omitempty"`               // Enable Neighbor List
	MfpClientProtection              string   `json:"mfpClientProtection,omitempty"`              // Management Frame Protection Client
	NasOptions                       []string `json:"nasOptions,omitempty"`                       // Nas Options
}
type RequestWirelessConfigureAccessPoints struct {
	ApList                      *[]RequestWirelessConfigureAccessPointsApList              `json:"apList,omitempty"`                      //
	ConfigureAdminStatus        *bool                                                      `json:"configureAdminStatus,omitempty"`        // To change the access point's admin status, set this parameter's value to "true".
	AdminStatus                 *bool                                                      `json:"adminStatus,omitempty"`                 // Configure the access point's admin status. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureApMode             *bool                                                      `json:"configureApMode,omitempty"`             // To change the access point's mode, set this parameter's value to "true".
	ApMode                      *int                                                       `json:"apMode,omitempty"`                      // Configure the access point's mode: for local/flexconnect mode, set "0"; for monitor mode, set "1"; for sniffer mode, set "4"; and for bridge/flex+bridge mode, set "5".
	ConfigureApHeight           *bool                                                      `json:"configureApHeight,omitempty"`           // To change the access point's height, set this parameter's value to "true".
	ApHeight                    *float64                                                   `json:"apHeight,omitempty"`                    // Configure the height of the access point by setting a value between 3 and height of the floor.
	ConfigureFailoverPriority   *bool                                                      `json:"configureFailoverPriority,omitempty"`   // To change the access point's failover priority, set this parameter's value to "true".
	FailoverPriority            *int                                                       `json:"failoverPriority,omitempty"`            // Configure the acess point's failover priority: for low, set "1"; for medium, set "2"; for high, set "3"; and for critical, set "4".
	ConfigureLedStatus          *bool                                                      `json:"configureLedStatus,omitempty"`          // To change the access point's LED status, set this parameter's value to "true".
	LedStatus                   *bool                                                      `json:"ledStatus,omitempty"`                   // Configure the access point's LED status. Set "true" to enable its status and "false" to disable it.
	ConfigureLedBrightnessLevel *bool                                                      `json:"configureLedBrightnessLevel,omitempty"` // To change the access point's LED brightness level, set this parameter's value to "true".
	LedBrightnessLevel          *int                                                       `json:"ledBrightnessLevel,omitempty"`          // Configure the access point's LED brightness level by setting a value between 1 and 8.
	ConfigureLocation           *bool                                                      `json:"configureLocation,omitempty"`           // To change the access point's location, set this parameter's value to "true".
	Location                    string                                                     `json:"location,omitempty"`                    // Configure the access point's location.
	ConfigureHAController       *bool                                                      `json:"configureHAController,omitempty"`       // To change the access point's HA controller, set this parameter's value to "true".
	PrimaryControllerName       string                                                     `json:"primaryControllerName,omitempty"`       // Configure the hostname for an access point's primary controller.
	PrimaryIPAddress            *RequestWirelessConfigureAccessPointsPrimaryIPAddress      `json:"primaryIpAddress,omitempty"`            //
	SecondaryControllerName     string                                                     `json:"secondaryControllerName,omitempty"`     // Configure the hostname for an access point's secondary controller.
	SecondaryIPAddress          *RequestWirelessConfigureAccessPointsSecondaryIPAddress    `json:"secondaryIpAddress,omitempty"`          //
	TertiaryControllerName      string                                                     `json:"tertiaryControllerName,omitempty"`      // Configure the hostname for an access point's tertiary controller.
	TertiaryIPAddress           *RequestWirelessConfigureAccessPointsTertiaryIPAddress     `json:"tertiaryIpAddress,omitempty"`           //
	RadioConfigurations         *[]RequestWirelessConfigureAccessPointsRadioConfigurations `json:"radioConfigurations,omitempty"`         //
}
type RequestWirelessConfigureAccessPointsApList struct {
	ApName     string `json:"apName,omitempty"`     // The current host name of the access point.
	MacAddress string `json:"macAddress,omitempty"` // The ethernet MAC address of the access point.
	ApNameNew  string `json:"apNameNew,omitempty"`  // The modified hostname of the access point.
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
	ConfigureRadioRoleAssignment *bool    `json:"configureRadioRoleAssignment,omitempty"` // To change the radio role on the specified radio for an access point, set this parameter's value to "true".
	RadioRoleAssignment          string   `json:"radioRoleAssignment,omitempty"`          // Configure one of the following roles on the specified radio for an access point: "auto", "serving", or "monitor".
	RadioBand                    string   `json:"radioBand,omitempty"`                    // Configure the band on the specified radio for an access point: for 2.4 GHz, set "RADIO24"; for 5 GHz, set "RADIO5".
	ConfigureAdminStatus         *bool    `json:"configureAdminStatus,omitempty"`         // To change the admin status on the specified radio for an access point, set this parameter's value to "true".
	AdminStatus                  *bool    `json:"adminStatus,omitempty"`                  // Configure the admin status on the specified radio for an access point. Set this parameter's value to "true" to enable it and "false" to disable it.
	ConfigureAntennaDegree       *bool    `json:"configureAntennaDegree,omitempty"`       // To change the antenna degree on the specified radio for an access point, set this parameter's value to "true".
	AntennaDegree                *int     `json:"antennaDegree,omitempty"`                // Configure the antenna degree on the specified radio for an access point.
	ConfigureElevAngleDegree     *bool    `json:"configureElevAngleDegree,omitempty"`     // To change the elevation angle degree on the specified radio for an access point, set this parameter's value to "true".
	AntennaElevAngleDegree       *int     `json:"antennaElevAngleDegree,omitempty"`       // Configure the antenna elevation angle on the specified radio for an access point.
	AntennaElevAngleSign         *int     `json:"antennaElevAngleSign,omitempty"`         // Configure the antenna elevation angle direction on the specified radio for an access point: for up, set "1"; for down, set "-1".
	ConfigureAntennaPatternName  *bool    `json:"configureAntennaPatternName,omitempty"`  // To change the antenna pattern name on the specified radio for an access point, set the value for this parameter to "true".
	AntennaPatternName           string   `json:"antennaPatternName,omitempty"`           // Configure the antenna pattern name on the specified radio for an access point. If antenna gain needs to be configured, set this parameter's value to "other".
	AntennaGain                  *int     `json:"antennaGain,omitempty"`                  // Configure the antenna gain on the specified radio for an access point by setting a decimal value (in dBi).
	ConfigureAntennaCable        *bool    `json:"configureAntennaCable,omitempty"`        // To change the antenna cable name on the specified radio for an access point, set this parameter's value to "true".
	AntennaCableName             string   `json:"antennaCableName,omitempty"`             // Configure the antenna cable name on the specified radio for an access point. If cable loss needs to be configured, set this parameter's value to "other".
	CableLoss                    *float64 `json:"cableLoss,omitempty"`                    // Configure the cable loss on the specified radio for an access point by setting a decimal value (in dBi).
	ConfigureChannel             *bool    `json:"configureChannel,omitempty"`             // To change the channel on the specified radio for an access point, set this parameter's value to "true".
	ChannelAssignmentMode        *int     `json:"channelAssignmentMode,omitempty"`        // Configure the channel assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
	ChannelNumber                *int     `json:"channelNumber,omitempty"`                // Configure the channel number on the specified radio for an access point.
	ConfigureChannelWidth        *bool    `json:"configureChannelWidth,omitempty"`        // To change the channel width on the specified radio for an access point, set this parameter's value to "true".
	ChannelWidth                 *int     `json:"channelWidth,omitempty"`                 // Configure the channel width on the specified radio for an access point: for 20 MHz, set "3"; for 40 MHz, set "4"; for 80 MHz, set "5"; and for 160 MHz, set "6".
	ConfigurePower               *bool    `json:"configurePower,omitempty"`               // To change the power assignment mode on the specified radio for an access point, set this parameter's value to "true".
	PowerAssignmentMode          *int     `json:"powerAssignmentMode,omitempty"`          // Configure the power assignment mode on the specified radio for an access point: for global mode, set "1"; and for custom mode, set "2".
	Powerlevel                   *int     `json:"powerlevel,omitempty"`                   // Configure the power level on the specified radio for an access point by setting a value between 1 and 8.
	ConfigureCleanAirSI          *bool    `json:"configureCleanAirSI,omitempty"`          // To enable or disable either CleanAir or Spectrum Intelligence on the specified radio for an access point, set this parameter's value to "true".
	CleanAirSI                   *int     `json:"cleanAirSI,omitempty"`                   // Configure CleanAir or Spectrum Intelligence on the specified radio for an access point. Set this parameter's value to "0" to disable the feature or "1" to enable it.
	RadioType                    *int     `json:"radioType,omitempty"`                    // Configure an access point's radio band: for 2.4 GHz, set "1"; for 5 GHz, set "2"; for XOR, set "3"; and for 6 GHz, set "6".
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
	Name              string                                                                    `json:"name,omitempty"`              // Ssid Name
	Type              string                                                                    `json:"type,omitempty"`              // Ssid Type(enum: Enterprise/Guest)
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
	Type              string                                                                    `json:"type,omitempty"`              // Ssid Type(enum: Enterprise/Guest)
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
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Rx Sop Threshold
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeBProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level
}
type RequestWirelessCreateOrUpdateRfProfileRadioTypeCProperties struct {
	ParentProfile      string   `json:"parentProfile,omitempty"`      // Parent Profile
	RadioChannels      string   `json:"radioChannels,omitempty"`      // Radio Channels
	DataRates          string   `json:"dataRates,omitempty"`          // Data Rates
	MandatoryDataRates string   `json:"mandatoryDataRates,omitempty"` // Mandatory Data Rates
	RxSopThreshold     string   `json:"rxSopThreshold,omitempty"`     // Rx Sop Threshold
	MinPowerLevel      *float64 `json:"minPowerLevel,omitempty"`      // Min Power Level
	MaxPowerLevel      *float64 `json:"maxPowerLevel,omitempty"`      // Max Power Level
	PowerThresholdV1   *float64 `json:"powerThresholdV1,omitempty"`   // Power Threshold V1
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

//GetAccessPointRebootTaskResult Get Access Point Reboot task result - c4b5-e9ce-460a-a8a3
/* Users can query the access point reboot status using this intent API


@param GetAccessPointRebootTaskResultQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetAccessPointRebootTaskResult")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointRebootTaskResult)
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

//GetAccessPointConfigurationTaskResult Get Access Point Configuration task result - fb90-69dc-4aeb-9afb
/* Users can query the access point configuration result using this intent API


@param taskTypeID task_id path parameter. task id information of ap config

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
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfigurationTaskResult")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfigurationTaskResult)
	return result, response, err

}

//GetAccessPointConfiguration Get Access Point Configuration - a191-f9f2-4cb8-9a55
/* Users can query the access point configuration information per device using the ethernet MAC address


@param GetAccessPointConfigurationQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetAccessPointConfiguration")
	}

	result := response.Result().(*ResponseWirelessGetAccessPointConfiguration)
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

//RebootAccessPoints Reboot Access Points - 6092-d8f1-468b-99ab
/* Users can reboot multiple access points up-to 200 at a time using this API


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
		return nil, response, fmt.Errorf("error with operation RebootAccessPoints")
	}

	result := response.Result().(*ResponseWirelessRebootAccessPoints)
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

//ConfigureAccessPoints Configure Access Points - 0081-cb89-4708-888f
/* User can configure multiple access points with required options using this intent API


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
		return nil, response, fmt.Errorf("error with operation ConfigureAccessPoints")
	}

	result := response.Result().(*ResponseWirelessConfigureAccessPoints)
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
/* Creates Wireless Network Profile on Cisco DNA Center and associates sites and SSIDs to it.


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
		return nil, response, fmt.Errorf("error with operation DeleteEnterpriseSsid")
	}

	result := response.Result().(*ResponseWirelessDeleteEnterpriseSSID)
	return result, response, err

}

//DeleteWirelessProfile Delete Wireless Profile - e395-88a5-4949-82c4
/* Delete the Wireless Profile from Cisco DNA Center whose name is provided.


@param wirelessProfileName wirelessProfileName path parameter. Wireless Profile Name

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
	//interfaceName string,DeleteDynamicInterfaceHeaderParams *DeleteDynamicInterfaceHeaderParams
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


@param rfProfileName rfProfileName path parameter. RF profile name to be deleted(required) *non-custom RF profile cannot be deleted

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
		return nil, response, fmt.Errorf("error with operation DeleteRfProfiles")
	}

	result := response.Result().(*ResponseWirelessDeleteRfProfiles)
	return result, response, err

}
