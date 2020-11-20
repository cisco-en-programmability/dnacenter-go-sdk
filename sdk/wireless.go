package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// WirelessService is the service to communicate with the Wireless API endpoint
type WirelessService service

// APProvisionRequest is the aPProvisionRequest definition
type APProvisionRequest struct {
	CustomApGroupName   string   `json:"customApGroupName,omitempty"`   //
	CustomFlexGroupName []string `json:"customFlexGroupName,omitempty"` //
	DeviceName          string   `json:"deviceName,omitempty"`          //
	RfProfile           string   `json:"rfProfile,omitempty"`           //
	SiteID              string   `json:"siteId,omitempty"`              //
	Type                string   `json:"type,omitempty"`                //
}

// APProvisionRequestCustomFlexGroupName is the aPProvisionRequestCustomFlexGroupName definition
type APProvisionRequestCustomFlexGroupName []string

// CreateAndProvisionSSIDRequest is the createAndProvisionSSIDRequest definition
type CreateAndProvisionSSIDRequest struct {
	EnableFabric       bool                                     `json:"enableFabric,omitempty"`       //
	FlexConnect        CreateAndProvisionSSIDRequestFlexConnect `json:"flexConnect,omitempty"`        //
	ManagedAPLocations []string                                 `json:"managedAPLocations,omitempty"` //
	SSIDDetails        CreateAndProvisionSSIDRequestSSIDDetails `json:"ssidDetails,omitempty"`        //
	SSIDType           string                                   `json:"ssidType,omitempty"`           //
}

// CreateAndProvisionSSIDRequestFlexConnect is the createAndProvisionSSIDRequestFlexConnect definition
type CreateAndProvisionSSIDRequestFlexConnect struct {
	EnableFlexConnect bool `json:"enableFlexConnect,omitempty"` //
	LocalToVLAN       int  `json:"localToVlan,omitempty"`       //
}

// CreateAndProvisionSSIDRequestManagedAPLocations is the createAndProvisionSSIDRequestManagedAPLocations definition
type CreateAndProvisionSSIDRequestManagedAPLocations []string

// CreateAndProvisionSSIDRequestSSIDDetails is the createAndProvisionSSIDRequestSSIDDetails definition
type CreateAndProvisionSSIDRequestSSIDDetails struct {
	EnableBroadcastSSID bool   `json:"enableBroadcastSSID,omitempty"` //
	EnableFastLane      bool   `json:"enableFastLane,omitempty"`      //
	EnableMACFiltering  bool   `json:"enableMACFiltering,omitempty"`  //
	FastTransition      string `json:"fastTransition,omitempty"`      //
	Name                string `json:"name,omitempty"`                //
	Passphrase          string `json:"passphrase,omitempty"`          //
	RadioPolicy         string `json:"radioPolicy,omitempty"`         //
	SecurityLevel       string `json:"securityLevel,omitempty"`       //
	TrafficType         string `json:"trafficType,omitempty"`         //
	WebAuthURL          string `json:"webAuthURL,omitempty"`          //
}

// CreateEnterpriseSSIDRequest is the createEnterpriseSSIDRequest definition
type CreateEnterpriseSSIDRequest struct {
	EnableBroadcastSSID bool   `json:"enableBroadcastSSID,omitempty"` //
	EnableFastLane      bool   `json:"enableFastLane,omitempty"`      //
	EnableMACFiltering  bool   `json:"enableMACFiltering,omitempty"`  //
	FastTransition      string `json:"fastTransition,omitempty"`      //
	Name                string `json:"name,omitempty"`                //
	Passphrase          string `json:"passphrase,omitempty"`          //
	RadioPolicy         string `json:"radioPolicy,omitempty"`         //
	SecurityLevel       string `json:"securityLevel,omitempty"`       //
	TrafficType         string `json:"trafficType,omitempty"`         //
}

// CreateOrUpdateRFProfileRequest is the createOrUpdateRFProfileRequest definition
type CreateOrUpdateRFProfileRequest struct {
	ChannelWidth         string                                             `json:"channelWidth,omitempty"`         //
	DefaultRfProfile     bool                                               `json:"defaultRfProfile,omitempty"`     //
	EnableBrownField     bool                                               `json:"enableBrownField,omitempty"`     //
	EnableCustom         bool                                               `json:"enableCustom,omitempty"`         //
	EnableRadioTypeA     bool                                               `json:"enableRadioTypeA,omitempty"`     //
	EnableRadioTypeB     bool                                               `json:"enableRadioTypeB,omitempty"`     //
	Name                 string                                             `json:"name,omitempty"`                 //
	RadioTypeAProperties CreateOrUpdateRFProfileRequestRadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //
	RadioTypeBProperties CreateOrUpdateRFProfileRequestRadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //
}

// CreateOrUpdateRFProfileRequestRadioTypeAProperties is the createOrUpdateRFProfileRequestRadioTypeAProperties definition
type CreateOrUpdateRFProfileRequestRadioTypeAProperties struct {
	DataRates          string  `json:"dataRates,omitempty"`          //
	MandatoryDataRates string  `json:"mandatoryDataRates,omitempty"` //
	MaxPowerLevel      float64 `json:"maxPowerLevel,omitempty"`      //
	MinPowerLevel      float64 `json:"minPowerLevel,omitempty"`      //
	ParentProfile      string  `json:"parentProfile,omitempty"`      //
	PowerThresholdV1   float64 `json:"powerThresholdV1,omitempty"`   //
	RadioChannels      string  `json:"radioChannels,omitempty"`      //
	RxSopThreshold     string  `json:"rxSopThreshold,omitempty"`     //
}

// CreateOrUpdateRFProfileRequestRadioTypeBProperties is the createOrUpdateRFProfileRequestRadioTypeBProperties definition
type CreateOrUpdateRFProfileRequestRadioTypeBProperties struct {
	DataRates          string  `json:"dataRates,omitempty"`          //
	MandatoryDataRates string  `json:"mandatoryDataRates,omitempty"` //
	MaxPowerLevel      float64 `json:"maxPowerLevel,omitempty"`      //
	MinPowerLevel      float64 `json:"minPowerLevel,omitempty"`      //
	ParentProfile      string  `json:"parentProfile,omitempty"`      //
	PowerThresholdV1   float64 `json:"powerThresholdV1,omitempty"`   //
	RadioChannels      string  `json:"radioChannels,omitempty"`      //
	RxSopThreshold     string  `json:"rxSopThreshold,omitempty"`     //
}

// CreateWirelessProfileRequest is the createWirelessProfileRequest definition
type CreateWirelessProfileRequest struct {
	ProfileDetails CreateWirelessProfileRequestProfileDetails `json:"profileDetails,omitempty"` //
}

// CreateWirelessProfileRequestProfileDetails is the createWirelessProfileRequestProfileDetails definition
type CreateWirelessProfileRequestProfileDetails struct {
	Name        string                                                  `json:"name,omitempty"`        //
	Sites       []string                                                `json:"sites,omitempty"`       //
	SSIDDetails []CreateWirelessProfileRequestProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}

// CreateWirelessProfileRequestProfileDetailsSSIDDetails is the createWirelessProfileRequestProfileDetailsSSIDDetails definition
type CreateWirelessProfileRequestProfileDetailsSSIDDetails struct {
	EnableFabric  bool                                                             `json:"enableFabric,omitempty"`  //
	FlexConnect   CreateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string                                                           `json:"interfaceName,omitempty"` //
	Name          string                                                           `json:"name,omitempty"`          //
	Type          string                                                           `json:"type,omitempty"`          //
}

// CreateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect is the createWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect definition
type CreateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect bool `json:"enableFlexConnect,omitempty"` //
	LocalToVLAN       int  `json:"localToVlan,omitempty"`       //
}

// CreateWirelessProfileRequestProfileDetailsSites is the createWirelessProfileRequestProfileDetailsSites definition
type CreateWirelessProfileRequestProfileDetailsSites []string

// ProvisionRequest is the provisionRequest definition
type ProvisionRequest struct {
	DeviceName         string                              `json:"deviceName,omitempty"`         //
	DynamicInterfaces  []ProvisionRequestDynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
	ManagedAPLocations []string                            `json:"managedAPLocations,omitempty"` //
	Site               string                              `json:"site,omitempty"`               //
}

// ProvisionRequestDynamicInterfaces is the provisionRequestDynamicInterfaces definition
type ProvisionRequestDynamicInterfaces struct {
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       //
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     //
	InterfaceName          string `json:"interfaceName,omitempty"`          //
	InterfaceNetmaskInCIDR int    `json:"interfaceNetmaskInCIDR,omitempty"` //
	LagOrPortNumber        int    `json:"lagOrPortNumber,omitempty"`        //
	VLANID                 int    `json:"vlanId,omitempty"`                 //
}

// ProvisionRequestManagedAPLocations is the provisionRequestManagedAPLocations definition
type ProvisionRequestManagedAPLocations []string

// ProvisionUpdateRequest is the provisionUpdateRequest definition
type ProvisionUpdateRequest struct {
	DeviceName         string                                    `json:"deviceName,omitempty"`         //
	DynamicInterfaces  []ProvisionUpdateRequestDynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
	ManagedAPLocations []string                                  `json:"managedAPLocations,omitempty"` //
}

// ProvisionUpdateRequestDynamicInterfaces is the provisionUpdateRequestDynamicInterfaces definition
type ProvisionUpdateRequestDynamicInterfaces struct {
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       //
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     //
	InterfaceName          string `json:"interfaceName,omitempty"`          //
	InterfaceNetmaskInCIDR int    `json:"interfaceNetmaskInCIDR,omitempty"` //
	LagOrPortNumber        int    `json:"lagOrPortNumber,omitempty"`        //
	VLANID                 int    `json:"vlanId,omitempty"`                 //
}

// ProvisionUpdateRequestManagedAPLocations is the provisionUpdateRequestManagedAPLocations definition
type ProvisionUpdateRequestManagedAPLocations []string

// UpdateWirelessProfileRequest is the updateWirelessProfileRequest definition
type UpdateWirelessProfileRequest struct {
	ProfileDetails UpdateWirelessProfileRequestProfileDetails `json:"profileDetails,omitempty"` //
}

// UpdateWirelessProfileRequestProfileDetails is the updateWirelessProfileRequestProfileDetails definition
type UpdateWirelessProfileRequestProfileDetails struct {
	Name        string                                                  `json:"name,omitempty"`        //
	Sites       []string                                                `json:"sites,omitempty"`       //
	SSIDDetails []UpdateWirelessProfileRequestProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}

// UpdateWirelessProfileRequestProfileDetailsSSIDDetails is the updateWirelessProfileRequestProfileDetailsSSIDDetails definition
type UpdateWirelessProfileRequestProfileDetailsSSIDDetails struct {
	EnableFabric  bool                                                             `json:"enableFabric,omitempty"`  //
	FlexConnect   UpdateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string                                                           `json:"interfaceName,omitempty"` //
	Name          string                                                           `json:"name,omitempty"`          //
	Type          string                                                           `json:"type,omitempty"`          //
}

// UpdateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect is the updateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect definition
type UpdateWirelessProfileRequestProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect bool `json:"enableFlexConnect,omitempty"` //
	LocalToVLAN       int  `json:"localToVlan,omitempty"`       //
}

// UpdateWirelessProfileRequestProfileDetailsSites is the updateWirelessProfileRequestProfileDetailsSites definition
type UpdateWirelessProfileRequestProfileDetailsSites []string

// APProvisionResponse is the aPProvisionResponse definition
type APProvisionResponse struct {
	ExecutionID    string                            `json:"executionId,omitempty"`    //
	ExecutionURL   string                            `json:"executionUrl,omitempty"`   //
	ProvisionTasks APProvisionResponseProvisionTasks `json:"provisionTasks,omitempty"` //
}

// APProvisionResponseProvisionTasks is the aPProvisionResponseProvisionTasks definition
type APProvisionResponseProvisionTasks struct {
	Failure APProvisionResponseProvisionTasksFailure   `json:"failure,omitempty"` //
	Success []APProvisionResponseProvisionTasksSuccess `json:"success,omitempty"` //
}

// APProvisionResponseProvisionTasksFailure is the aPProvisionResponseProvisionTasksFailure definition
type APProvisionResponseProvisionTasksFailure struct {
	FailureReason string `json:"failureReason,omitempty"` //
	TaskID        string `json:"taskId,omitempty"`        //
	TaskURL       string `json:"taskUrl,omitempty"`       //
}

// APProvisionResponseProvisionTasksSuccess is the aPProvisionResponseProvisionTasksSuccess definition
type APProvisionResponseProvisionTasksSuccess struct {
	TaskID  string `json:"taskId,omitempty"`  //
	TaskURL string `json:"taskUrl,omitempty"` //
}

// CreateAndProvisionSSIDResponse is the createAndProvisionSSIDResponse definition
type CreateAndProvisionSSIDResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateEnterpriseSSIDResponse is the createEnterpriseSSIDResponse definition
type CreateEnterpriseSSIDResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateOrUpdateRFProfileResponse is the createOrUpdateRFProfileResponse definition
type CreateOrUpdateRFProfileResponse struct {
	ExecutionID  string `json:"executionId,omitempty"`  //
	ExecutionURL string `json:"executionUrl,omitempty"` //
	Message      string `json:"message,omitempty"`      //
}

// CreateWirelessProfileResponse is the createWirelessProfileResponse definition
type CreateWirelessProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteEnterpriseSSIDResponse is the deleteEnterpriseSSIDResponse definition
type DeleteEnterpriseSSIDResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteRFProfilesResponse is the deleteRFProfilesResponse definition
type DeleteRFProfilesResponse struct {
	ExecutionID  string `json:"executionId,omitempty"`  //
	ExecutionURL string `json:"executionUrl,omitempty"` //
	Message      string `json:"message,omitempty"`      //
}

// DeleteSSIDAndProvisionItToDevicesResponse is the deleteSSIDAndProvisionItToDevicesResponse definition
type DeleteSSIDAndProvisionItToDevicesResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteWirelessProfileResponse is the deleteWirelessProfileResponse definition
type DeleteWirelessProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetEnterpriseSSIDResponse is the getEnterpriseSSIDResponse definition
type GetEnterpriseSSIDResponse struct {
	GroupUUID          string                                 `json:"groupUuid,omitempty"`          //
	InheritedGroupName string                                 `json:"inheritedGroupName,omitempty"` //
	InheritedGroupUUID string                                 `json:"inheritedGroupUuid,omitempty"` //
	InstanceUUID       string                                 `json:"instanceUuid,omitempty"`       //
	SSIDDetails        []GetEnterpriseSSIDResponseSSIDDetails `json:"ssidDetails,omitempty"`        //
	Version            int                                    `json:"version,omitempty"`            //
}

// GetEnterpriseSSIDResponseSSIDDetails is the getEnterpriseSSIDResponseSSIDDetails definition
type GetEnterpriseSSIDResponseSSIDDetails struct {
	AuthServer          string `json:"authServer,omitempty"`          //
	EnableBroadcastSSID bool   `json:"enableBroadcastSSID,omitempty"` //
	EnableFastLane      bool   `json:"enableFastLane,omitempty"`      //
	EnableMACFiltering  bool   `json:"enableMACFiltering,omitempty"`  //
	FastTransition      string `json:"fastTransition,omitempty"`      //
	IsEnabled           bool   `json:"isEnabled,omitempty"`           //
	IsFabric            bool   `json:"isFabric,omitempty"`            //
	Name                string `json:"name,omitempty"`                //
	Passphrase          string `json:"passphrase,omitempty"`          //
	RadioPolicy         string `json:"radioPolicy,omitempty"`         //
	SecurityLevel       string `json:"securityLevel,omitempty"`       //
	TrafficType         string `json:"trafficType,omitempty"`         //
	WLANType            string `json:"wlanType,omitempty"`            //
}

// GetWirelessProfileResponse is the getWirelessProfileResponse definition
type GetWirelessProfileResponse struct {
	ProfileDetails GetWirelessProfileResponseProfileDetails `json:"profileDetails,omitempty"` //
}

// GetWirelessProfileResponseProfileDetails is the getWirelessProfileResponseProfileDetails definition
type GetWirelessProfileResponseProfileDetails struct {
	Name        string                                                `json:"name,omitempty"`        //
	Sites       []string                                              `json:"sites,omitempty"`       //
	SSIDDetails []GetWirelessProfileResponseProfileDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}

// GetWirelessProfileResponseProfileDetailsSSIDDetails is the getWirelessProfileResponseProfileDetailsSSIDDetails definition
type GetWirelessProfileResponseProfileDetailsSSIDDetails struct {
	EnableFabric  bool                                                           `json:"enableFabric,omitempty"`  //
	FlexConnect   GetWirelessProfileResponseProfileDetailsSSIDDetailsFlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string                                                         `json:"interfaceName,omitempty"` //
	Name          string                                                         `json:"name,omitempty"`          //
	Type          string                                                         `json:"type,omitempty"`          //
}

// GetWirelessProfileResponseProfileDetailsSSIDDetailsFlexConnect is the getWirelessProfileResponseProfileDetailsSSIDDetailsFlexConnect definition
type GetWirelessProfileResponseProfileDetailsSSIDDetailsFlexConnect struct {
	EnableFlexConnect bool `json:"enableFlexConnect,omitempty"` //
	LocalToVLAN       int  `json:"localToVlan,omitempty"`       //
}

// GetWirelessProfileResponseProfileDetailsSites is the getWirelessProfileResponseProfileDetailsSites definition
type GetWirelessProfileResponseProfileDetailsSites []string

// ProvisionResponse is the provisionResponse definition
type ProvisionResponse struct {
	ExecutionID       string                             `json:"executionId,omitempty"`       //
	ExecutionURL      string                             `json:"executionUrl,omitempty"`      //
	ProvisioningTasks ProvisionResponseProvisioningTasks `json:"provisioningTasks,omitempty"` //
}

// ProvisionResponseProvisioningTasks is the provisionResponseProvisioningTasks definition
type ProvisionResponseProvisioningTasks struct {
	Failed  []string `json:"failed,omitempty"`  //
	Success []string `json:"success,omitempty"` //
}

// ProvisionResponseProvisioningTasksFailed is the provisionResponseProvisioningTasksFailed definition
type ProvisionResponseProvisioningTasksFailed []string

// ProvisionResponseProvisioningTasksSuccess is the provisionResponseProvisioningTasksSuccess definition
type ProvisionResponseProvisioningTasksSuccess []string

// ProvisionUpdateResponse is the provisionUpdateResponse definition
type ProvisionUpdateResponse struct {
	ExecutionID       string                                   `json:"executionId,omitempty"`       //
	ExecutionURL      string                                   `json:"executionUrl,omitempty"`      //
	ProvisioningTasks ProvisionUpdateResponseProvisioningTasks `json:"provisioningTasks,omitempty"` //
}

// ProvisionUpdateResponseProvisioningTasks is the provisionUpdateResponseProvisioningTasks definition
type ProvisionUpdateResponseProvisioningTasks struct {
	Failed  []string `json:"failed,omitempty"`  //
	Success []string `json:"success,omitempty"` //
}

// ProvisionUpdateResponseProvisioningTasksFailed is the provisionUpdateResponseProvisioningTasksFailed definition
type ProvisionUpdateResponseProvisioningTasksFailed []string

// ProvisionUpdateResponseProvisioningTasksSuccess is the provisionUpdateResponseProvisioningTasksSuccess definition
type ProvisionUpdateResponseProvisioningTasksSuccess []string

// RetrieveRFProfilesResponse is the retrieveRFProfilesResponse definition
type RetrieveRFProfilesResponse struct {
	Response []RetrieveRFProfilesResponseResponse `json:"response,omitempty"` //
}

// RetrieveRFProfilesResponseResponse is the retrieveRFProfilesResponseResponse definition
type RetrieveRFProfilesResponseResponse struct {
	ARadioChannels      string `json:"aRadioChannels,omitempty"`      //
	BRadioChannels      string `json:"bRadioChannels,omitempty"`      //
	ChannelWidth        string `json:"channelWidth,omitempty"`        //
	DataRatesA          string `json:"dataRatesA,omitempty"`          //
	DataRatesB          string `json:"dataRatesB,omitempty"`          //
	DefaultRfProfile    bool   `json:"defaultRfProfile,omitempty"`    //
	EnableARadioType    bool   `json:"enableARadioType,omitempty"`    //
	EnableBRadioType    bool   `json:"enableBRadioType,omitempty"`    //
	EnableBrownField    bool   `json:"enableBrownField,omitempty"`    //
	EnableCustom        bool   `json:"enableCustom,omitempty"`        //
	MandatoryDataRatesA string `json:"mandatoryDataRatesA,omitempty"` //
	MandatoryDataRatesB string `json:"mandatoryDataRatesB,omitempty"` //
	MaxPowerLevelA      string `json:"maxPowerLevelA,omitempty"`      //
	MaxPowerLevelB      string `json:"maxPowerLevelB,omitempty"`      //
	MinPowerLevelA      string `json:"minPowerLevelA,omitempty"`      //
	MinPowerLevelB      string `json:"minPowerLevelB,omitempty"`      //
	Name                string `json:"name,omitempty"`                //
	ParentProfileA      string `json:"parentProfileA,omitempty"`      //
	ParentProfileB      string `json:"parentProfileB,omitempty"`      //
	PowerThresholdV1A   int    `json:"powerThresholdV1A,omitempty"`   //
	PowerThresholdV1B   int    `json:"powerThresholdV1B,omitempty"`   //
	RxSopThresholdA     string `json:"rxSopThresholdA,omitempty"`     //
	RxSopThresholdB     string `json:"rxSopThresholdB,omitempty"`     //
}

// SensorTestResultsResponse is the sensorTestResultsResponse definition
type SensorTestResultsResponse struct {
	FailureStats []SensorTestResultsResponseFailureStats `json:"failureStats,omitempty"` //
	Summary      SensorTestResultsResponseSummary        `json:"summary,omitempty"`      //
}

// SensorTestResultsResponseFailureStats is the sensorTestResultsResponseFailureStats definition
type SensorTestResultsResponseFailureStats struct {
	ErrorCode    int    `json:"errorCode,omitempty"`    //
	ErrorTitle   string `json:"errorTitle,omitempty"`   //
	TestCategory string `json:"testCategory,omitempty"` //
	TestType     string `json:"testType,omitempty"`     //
}

// SensorTestResultsResponseSummary is the sensorTestResultsResponseSummary definition
type SensorTestResultsResponseSummary struct {
	AppConnectivity SensorTestResultsResponseSummaryAppConnectivity `json:"APP_CONNECTIVITY,omitempty"` //
	Email           SensorTestResultsResponseSummaryEmail           `json:"EMAIL,omitempty"`            //
	NetworkServices SensorTestResultsResponseSummaryNetworkServices `json:"NETWORK_SERVICES,omitempty"` //
	OnBoarding      SensorTestResultsResponseSummaryOnBoarding      `json:"ONBOARDING,omitempty"`       //
	Performance     SensorTestResultsResponseSummaryPerformance     `json:"PERFORMANCE,omitempty"`      //
	RFAssessment    SensorTestResultsResponseSummaryRFAssessment    `json:"RF_ASSESSMENT,omitempty"`    //
	TotalTestCount  int                                             `json:"totalTestCount,omitempty"`   //
}

// SensorTestResultsResponseSummaryAppConnectivity is the sensorTestResultsResponseSummaryAppConnectivity definition
type SensorTestResultsResponseSummaryAppConnectivity struct {
	FileTransfer     SensorTestResultsResponseSummaryAppConnectivityFileTransfer     `json:"FILETRANSFER,omitempty"`      //
	HostReachability SensorTestResultsResponseSummaryAppConnectivityHostReachability `json:"HOST_REACHABILITY,omitempty"` //
	WebServer        SensorTestResultsResponseSummaryAppConnectivityWebServer        `json:"WEBSERVER,omitempty"`         //
}

// SensorTestResultsResponseSummaryAppConnectivityFileTransfer is the sensorTestResultsResponseSummaryAppConnectivityFileTransfer definition
type SensorTestResultsResponseSummaryAppConnectivityFileTransfer struct {
	FailCount int     `json:"failCount,omitempty"` //
	PassCount float64 `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryAppConnectivityHostReachability is the sensorTestResultsResponseSummaryAppConnectivityHostReachability definition
type SensorTestResultsResponseSummaryAppConnectivityHostReachability struct {
	FailCount float64 `json:"failCount,omitempty"` //
	PassCount int     `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryAppConnectivityWebServer is the sensorTestResultsResponseSummaryAppConnectivityWebServer definition
type SensorTestResultsResponseSummaryAppConnectivityWebServer struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryEmail is the sensorTestResultsResponseSummaryEmail definition
type SensorTestResultsResponseSummaryEmail struct {
	MailServer SensorTestResultsResponseSummaryEmailMailServer `json:"MAILSERVER,omitempty"` //
}

// SensorTestResultsResponseSummaryEmailMailServer is the sensorTestResultsResponseSummaryEmailMailServer definition
type SensorTestResultsResponseSummaryEmailMailServer struct {
	FailCount int     `json:"failCount,omitempty"` //
	PassCount float64 `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryNetworkServices is the sensorTestResultsResponseSummaryNetworkServices definition
type SensorTestResultsResponseSummaryNetworkServices struct {
	DNS SensorTestResultsResponseSummaryNetworkServicesDNS `json:"DNS,omitempty"` //
}

// SensorTestResultsResponseSummaryNetworkServicesDNS is the sensorTestResultsResponseSummaryNetworkServicesDNS definition
type SensorTestResultsResponseSummaryNetworkServicesDNS struct {
	FailCount float64 `json:"failCount,omitempty"` //
	PassCount int     `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryOnBoarding is the sensorTestResultsResponseSummaryOnBoarding definition
type SensorTestResultsResponseSummaryOnBoarding struct {
	Assoc SensorTestResultsResponseSummaryOnBoardingAssoc `json:"ASSOC,omitempty"` //
	Auth  SensorTestResultsResponseSummaryOnBoardingAuth  `json:"AUTH,omitempty"`  //
	DHCP  SensorTestResultsResponseSummaryOnBoardingDHCP  `json:"DHCP,omitempty"`  //
}

// SensorTestResultsResponseSummaryOnBoardingAssoc is the sensorTestResultsResponseSummaryOnBoardingAssoc definition
type SensorTestResultsResponseSummaryOnBoardingAssoc struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryOnBoardingAuth is the sensorTestResultsResponseSummaryOnBoardingAuth definition
type SensorTestResultsResponseSummaryOnBoardingAuth struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryOnBoardingDHCP is the sensorTestResultsResponseSummaryOnBoardingDHCP definition
type SensorTestResultsResponseSummaryOnBoardingDHCP struct {
	FailCount float64 `json:"failCount,omitempty"` //
	PassCount int     `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryPerformance is the sensorTestResultsResponseSummaryPerformance definition
type SensorTestResultsResponseSummaryPerformance struct {
	IPSLASender SensorTestResultsResponseSummaryPerformanceIPSLASender `json:"IPSLASENDER,omitempty"` //
}

// SensorTestResultsResponseSummaryPerformanceIPSLASender is the sensorTestResultsResponseSummaryPerformanceIPSLASender definition
type SensorTestResultsResponseSummaryPerformanceIPSLASender struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryRFAssessment is the sensorTestResultsResponseSummaryRFAssessment definition
type SensorTestResultsResponseSummaryRFAssessment struct {
	DataRate SensorTestResultsResponseSummaryRFAssessmentDataRate `json:"DATA_RATE,omitempty"` //
	SNR      SensorTestResultsResponseSummaryRFAssessmentSNR      `json:"SNR,omitempty"`       //
}

// SensorTestResultsResponseSummaryRFAssessmentDataRate is the sensorTestResultsResponseSummaryRFAssessmentDataRate definition
type SensorTestResultsResponseSummaryRFAssessmentDataRate struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// SensorTestResultsResponseSummaryRFAssessmentSNR is the sensorTestResultsResponseSummaryRFAssessmentSNR definition
type SensorTestResultsResponseSummaryRFAssessmentSNR struct {
	FailCount float64 `json:"failCount,omitempty"` //
	PassCount int     `json:"passCount,omitempty"` //
}

// UpdateWirelessProfileResponse is the updateWirelessProfileResponse definition
type UpdateWirelessProfileResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// APProvision aPProvision
/* Provision wireless Access points
@param __persistbapioutput
*/
func (s *WirelessService) APProvision(aPProvisionRequest *[]APProvisionRequest) (*APProvisionResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/ap-provision"

	response, err := RestyClient.R().
		SetBody(aPProvisionRequest).
		SetResult(&APProvisionResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation aPProvision")
	}

	result := response.Result().(*APProvisionResponse)
	return result, response, err
}

// CreateAndProvisionSSID createAndProvisionSSID
/* Creates SSID, updates the SSID to the corresponding site profiles and provision it to the devices matching the given sites
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
*/
func (s *WirelessService) CreateAndProvisionSSID(createAndProvisionSSIDRequest *CreateAndProvisionSSIDRequest) (*CreateAndProvisionSSIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/ssid"

	response, err := RestyClient.R().
		SetBody(createAndProvisionSSIDRequest).
		SetResult(&CreateAndProvisionSSIDResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createAndProvisionSSID")
	}

	result := response.Result().(*CreateAndProvisionSSIDResponse)
	return result, response, err
}

// CreateEnterpriseSSID createEnterpriseSSID
/* Creates enterprise SSID
 */
func (s *WirelessService) CreateEnterpriseSSID(createEnterpriseSSIDRequest *CreateEnterpriseSSIDRequest) (*CreateEnterpriseSSIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/enterprise-ssid"

	response, err := RestyClient.R().
		SetBody(createEnterpriseSSIDRequest).
		SetResult(&CreateEnterpriseSSIDResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createEnterpriseSSID")
	}

	result := response.Result().(*CreateEnterpriseSSIDResponse)
	return result, response, err
}

// CreateOrUpdateRFProfile createOrUpdateRFProfile
/* Create or Update RF profile
 */
func (s *WirelessService) CreateOrUpdateRFProfile(createOrUpdateRFProfileRequest *CreateOrUpdateRFProfileRequest) (*CreateOrUpdateRFProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/rf-profile"

	response, err := RestyClient.R().
		SetBody(createOrUpdateRFProfileRequest).
		SetResult(&CreateOrUpdateRFProfileResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createOrUpdateRFProfile")
	}

	result := response.Result().(*CreateOrUpdateRFProfileResponse)
	return result, response, err
}

// CreateWirelessProfile createWirelessProfile
/* Creates Wireless Network Profile on DNAC and associates sites and SSIDs to it.
 */
func (s *WirelessService) CreateWirelessProfile(createWirelessProfileRequest *CreateWirelessProfileRequest) (*CreateWirelessProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/profile"

	response, err := RestyClient.R().
		SetBody(createWirelessProfileRequest).
		SetResult(&CreateWirelessProfileResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createWirelessProfile")
	}

	result := response.Result().(*CreateWirelessProfileResponse)
	return result, response, err
}

// DeleteEnterpriseSSID deleteEnterpriseSSID
/* Deletes given enterprise SSID
@param ssidName Enter the SSID name to be deleted
*/
func (s *WirelessService) DeleteEnterpriseSSID(ssidName string) (*DeleteEnterpriseSSIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/enterprise-ssid/{ssidName}"
	path = strings.Replace(path, "{"+"ssidName"+"}", fmt.Sprintf("%v", ssidName), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteEnterpriseSSID")
	}

	result := response.Result().(*DeleteEnterpriseSSIDResponse)
	return result, response, err
}

// DeleteRFProfiles deleteRFProfiles
/* Delete RF profile(s)
@param rfProfileName
*/
func (s *WirelessService) DeleteRFProfiles(rfProfileName string) (*DeleteRFProfilesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/rf-profile/{rfProfileName}"
	path = strings.Replace(path, "{"+"rfProfileName"+"}", fmt.Sprintf("%v", rfProfileName), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteRFProfiles")
	}

	result := response.Result().(*DeleteRFProfilesResponse)
	return result, response, err
}

// DeleteSSIDAndProvisionItToDevices deleteSSIDAndProvisionItToDevices
/* Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from DNA Center
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
@param ssidName
@param managedAPLocations
*/
func (s *WirelessService) DeleteSSIDAndProvisionItToDevices(ssidName string, managedAPLocations string) (*DeleteSSIDAndProvisionItToDevicesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/business/ssid/{ssidName}/{managedAPLocations}"
	path = strings.Replace(path, "{"+"ssidName"+"}", fmt.Sprintf("%v", ssidName), -1)
	path = strings.Replace(path, "{"+"managedAPLocations"+"}", fmt.Sprintf("%v", managedAPLocations), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteSSIDAndProvisionItToDevices")
	}

	result := response.Result().(*DeleteSSIDAndProvisionItToDevicesResponse)
	return result, response, err
}

// DeleteWirelessProfile deleteWirelessProfile
/* Delete the Wireless Profile from DNAC whose name is provided.
@param wirelessProfileName
*/
// func (s *WirelessService) DeleteWirelessProfile(wirelessProfileName string, deleteWirelessProfileRequest *DeleteWirelessProfileRequest) (*DeleteWirelessProfileResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/wireless-profile/{wirelessProfileName}"
// 	path = strings.Replace(path, "{"+"wirelessProfileName"+"}", fmt.Sprintf("%v", wirelessProfileName), -1)

// 	response, err := RestyClient.R().
// 		SetBody(deleteWirelessProfileRequest).
// 		SetError(&Error{}).
// 		Delete(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if response.IsError() {
// 		return nil, response, fmt.Errorf("Error with operation deleteWirelessProfile")
// 	}

// 	result := response.Result().(*DeleteWirelessProfileResponse)
// 	return result, response, err
// }

// GetEnterpriseSSIDQueryParams defines the query parameters for this request
type GetEnterpriseSSIDQueryParams struct {
	SSIDName string `url:"ssidName,omitempty"` // Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
}

// GetEnterpriseSSID getEnterpriseSSID
/* Gets either one or all the enterprise SSID
@param ssidName Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
*/
func (s *WirelessService) GetEnterpriseSSID(getEnterpriseSSIDQueryParams *GetEnterpriseSSIDQueryParams) (*GetEnterpriseSSIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/enterprise-ssid"

	queryString, _ := query.Values(getEnterpriseSSIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetEnterpriseSSIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getEnterpriseSSID")
	}

	result := response.Result().(*GetEnterpriseSSIDResponse)
	return result, response, err
}

// GetWirelessProfileQueryParams defines the query parameters for this request
type GetWirelessProfileQueryParams struct {
	ProfileName string `url:"profileName,omitempty"` //
}

// GetWirelessProfile getWirelessProfile
/* Gets either one or all the wireless network profiles if no name is provided for network-profile.
@param profileName
*/
func (s *WirelessService) GetWirelessProfile(getWirelessProfileQueryParams *GetWirelessProfileQueryParams) (*GetWirelessProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/profile"

	queryString, _ := query.Values(getWirelessProfileQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetWirelessProfileResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getWirelessProfile")
	}

	result := response.Result().(*GetWirelessProfileResponse)
	return result, response, err
}

// Provision provision
/* Provision wireless devices
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
*/
func (s *WirelessService) Provision(provisionRequest *[]ProvisionRequest) (*ProvisionResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/provision"

	response, err := RestyClient.R().
		SetBody(provisionRequest).
		SetResult(&ProvisionResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation provision")
	}

	result := response.Result().(*ProvisionResponse)
	return result, response, err
}

// ProvisionUpdate provisionUpdate
/* Updates wireless provisioning
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
*/
func (s *WirelessService) ProvisionUpdate(provisionUpdateRequest *[]ProvisionUpdateRequest) (*ProvisionUpdateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/provision"

	response, err := RestyClient.R().
		SetBody(provisionUpdateRequest).
		SetResult(&ProvisionUpdateResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation provisionUpdate")
	}

	result := response.Result().(*ProvisionUpdateResponse)
	return result, response, err
}

// RetrieveRFProfilesQueryParams defines the query parameters for this request
type RetrieveRFProfilesQueryParams struct {
	RfProfileName string `url:"rfProfileName,omitempty"` //
}

// RetrieveRFProfiles retrieveRFProfiles
/* Retrieve all RF profiles
@param rfProfileName
*/
func (s *WirelessService) RetrieveRFProfiles(retrieveRFProfilesQueryParams *RetrieveRFProfilesQueryParams) (*RetrieveRFProfilesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/rf-profile"

	queryString, _ := query.Values(retrieveRFProfilesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&RetrieveRFProfilesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation retrieveRFProfiles")
	}

	result := response.Result().(*RetrieveRFProfilesResponse)
	return result, response, err
}

// SensorTestResultsQueryParams defines the query parameters for this request
type SensorTestResultsQueryParams struct {
	SiteID        string  `url:"siteId,omitempty"`        // Assurance site UUID
	StartTime     float64 `url:"startTime,omitempty"`     // The epoch time in milliseconds
	EndTime       float64 `url:"endTime,omitempty"`       // The epoch time in milliseconds
	TestFailureBy string  `url:"testFailureBy,omitempty"` // Obtain failure statistics group by "area", "building", or "floor"
}

// SensorTestResults sensorTestResults
/* Intent API to get SENSOR test result summary
@param siteID Assurance site UUID
@param startTime The epoch time in milliseconds
@param endTime The epoch time in milliseconds
@param testFailureBy Obtain failure statistics group by "area", "building", or "floor"
*/
func (s *WirelessService) SensorTestResults(sensorTestResultsQueryParams *SensorTestResultsQueryParams) (*SensorTestResultsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/AssuranceGetSensorTestResults"

	queryString, _ := query.Values(sensorTestResultsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&SensorTestResultsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation sensorTestResults")
	}

	result := response.Result().(*SensorTestResultsResponse)
	return result, response, err
}

// UpdateWirelessProfile updateWirelessProfile
/* Updates the wireless Network Profile with updated details provided. All sites to be present in the network profile should be provided.
 */
func (s *WirelessService) UpdateWirelessProfile(updateWirelessProfileRequest *UpdateWirelessProfileRequest) (*UpdateWirelessProfileResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/profile"

	response, err := RestyClient.R().
		SetBody(updateWirelessProfileRequest).
		SetResult(&UpdateWirelessProfileResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateWirelessProfile")
	}

	result := response.Result().(*UpdateWirelessProfileResponse)
	return result, response, err
}
