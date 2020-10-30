package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// WirelessService is the service to communicate with the Wireless API endpoint
type WirelessService service

// APProvisionRequest is the APProvisionRequest definition
type APProvisionRequest struct {
	CustomApGroupName   string   `json:"customApGroupName,omitempty"`   //
	CustomFlexGroupName []string `json:"customFlexGroupName,omitempty"` //
	DeviceName          string   `json:"deviceName,omitempty"`          //
	RfProfile           string   `json:"rfProfile,omitempty"`           //
	SiteId              string   `json:"siteId,omitempty"`              //
	Type                string   `json:"type,omitempty"`                //
}

// CreateAndProvisionSSIDRequest is the CreateAndProvisionSSIDRequest definition
type CreateAndProvisionSSIDRequest struct {
	EnableFabric       bool        `json:"enableFabric,omitempty"`       //
	FlexConnect        FlexConnect `json:"flexConnect,omitempty"`        //
	ManagedAPLocations []string    `json:"managedAPLocations,omitempty"` //
	SsidDetails        SsidDetails `json:"ssidDetails,omitempty"`        //
	SsidType           string      `json:"ssidType,omitempty"`           //
}

// CreateEnterpriseSSIDRequest is the CreateEnterpriseSSIDRequest definition
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

// CreateOrUpdateRFProfileRequest is the CreateOrUpdateRFProfileRequest definition
type CreateOrUpdateRFProfileRequest struct {
	ChannelWidth         string               `json:"channelWidth,omitempty"`         //
	DefaultRfProfile     bool                 `json:"defaultRfProfile,omitempty"`     //
	EnableBrownField     bool                 `json:"enableBrownField,omitempty"`     //
	EnableCustom         bool                 `json:"enableCustom,omitempty"`         //
	EnableRadioTypeA     bool                 `json:"enableRadioTypeA,omitempty"`     //
	EnableRadioTypeB     bool                 `json:"enableRadioTypeB,omitempty"`     //
	Name                 string               `json:"name,omitempty"`                 //
	RadioTypeAProperties RadioTypeAProperties `json:"radioTypeAProperties,omitempty"` //
	RadioTypeBProperties RadioTypeBProperties `json:"radioTypeBProperties,omitempty"` //
}

// CreateWirelessProfileRequest is the CreateWirelessProfileRequest definition
type CreateWirelessProfileRequest struct {
	ProfileDetails ProfileDetails `json:"profileDetails,omitempty"` //
}

// DynamicInterfaces is the DynamicInterfaces definition
type DynamicInterfaces struct {
	InterfaceGateway       string `json:"interfaceGateway,omitempty"`       //
	InterfaceIPAddress     string `json:"interfaceIPAddress,omitempty"`     //
	InterfaceName          string `json:"interfaceName,omitempty"`          //
	InterfaceNetmaskInCIDR int    `json:"interfaceNetmaskInCIDR,omitempty"` //
	LagOrPortNumber        int    `json:"lagOrPortNumber,omitempty"`        //
	VlanId                 int    `json:"vlanId,omitempty"`                 //
}

// FlexConnect is the FlexConnect definition
type FlexConnect struct {
	EnableFlexConnect bool `json:"enableFlexConnect,omitempty"` //
	LocalToVlan       int  `json:"localToVlan,omitempty"`       //
}

// ProfileDetails is the ProfileDetails definition
type ProfileDetails struct {
	Name        string        `json:"name,omitempty"`        //
	Sites       []string      `json:"sites,omitempty"`       //
	SsidDetails []SsidDetails `json:"ssidDetails,omitempty"` //
}

// ProvisionRequest is the ProvisionRequest definition
type ProvisionRequest struct {
	DeviceName         string              `json:"deviceName,omitempty"`         //
	DynamicInterfaces  []DynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
	ManagedAPLocations []string            `json:"managedAPLocations,omitempty"` //
	Site               string              `json:"site,omitempty"`               //
}

// ProvisionUpdateRequest is the ProvisionUpdateRequest definition
type ProvisionUpdateRequest struct {
	DeviceName         string              `json:"deviceName,omitempty"`         //
	DynamicInterfaces  []DynamicInterfaces `json:"dynamicInterfaces,omitempty"`  //
	ManagedAPLocations []string            `json:"managedAPLocations,omitempty"` //
}

// RadioTypeAProperties is the RadioTypeAProperties definition
type RadioTypeAProperties struct {
	DataRates          string `json:"dataRates,omitempty"`          //
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` //
	MaxPowerLevel      int    `json:"maxPowerLevel,omitempty"`      //
	MinPowerLevel      int    `json:"minPowerLevel,omitempty"`      //
	ParentProfile      string `json:"parentProfile,omitempty"`      //
	PowerThresholdV1   int    `json:"powerThresholdV1,omitempty"`   //
	RadioChannels      string `json:"radioChannels,omitempty"`      //
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     //
}

// RadioTypeBProperties is the RadioTypeBProperties definition
type RadioTypeBProperties struct {
	DataRates          string `json:"dataRates,omitempty"`          //
	MandatoryDataRates string `json:"mandatoryDataRates,omitempty"` //
	MaxPowerLevel      int    `json:"maxPowerLevel,omitempty"`      //
	MinPowerLevel      int    `json:"minPowerLevel,omitempty"`      //
	ParentProfile      string `json:"parentProfile,omitempty"`      //
	PowerThresholdV1   int    `json:"powerThresholdV1,omitempty"`   //
	RadioChannels      string `json:"radioChannels,omitempty"`      //
	RxSopThreshold     string `json:"rxSopThreshold,omitempty"`     //
}

// SsidDetails is the SsidDetails definition
type SsidDetails struct {
	EnableFabric  bool        `json:"enableFabric,omitempty"`  //
	FlexConnect   FlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string      `json:"interfaceName,omitempty"` //
	Name          string      `json:"name,omitempty"`          //
	Type          string      `json:"type,omitempty"`          //
}

// UpdateWirelessProfileRequest is the UpdateWirelessProfileRequest definition
type UpdateWirelessProfileRequest struct {
	ProfileDetails ProfileDetails `json:"profileDetails,omitempty"` //
}

// APP_CONNECTIVITY is the APP_CONNECTIVITY definition
type APP_CONNECTIVITY struct {
	FILETRANSFER      FILETRANSFER      `json:"FILETRANSFER,omitempty"`      //
	HOST_REACHABILITY HOST_REACHABILITY `json:"HOST_REACHABILITY,omitempty"` //
	WEBSERVER         WEBSERVER         `json:"WEBSERVER,omitempty"`         //
}

// APProvisionResponse is the APProvisionResponse definition
type APProvisionResponse struct {
	ExecutionId    string         `json:"executionId,omitempty"`    //
	ExecutionUrl   string         `json:"executionUrl,omitempty"`   //
	ProvisionTasks ProvisionTasks `json:"provisionTasks,omitempty"` //
}

// ASSOC is the ASSOC definition
type ASSOC struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// AUTH is the AUTH definition
type AUTH struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// CreateAndProvisionSSIDResponse is the CreateAndProvisionSSIDResponse definition
type CreateAndProvisionSSIDResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateEnterpriseSSIDResponse is the CreateEnterpriseSSIDResponse definition
type CreateEnterpriseSSIDResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// CreateOrUpdateRFProfileResponse is the CreateOrUpdateRFProfileResponse definition
type CreateOrUpdateRFProfileResponse struct {
	ExecutionId  string `json:"executionId,omitempty"`  //
	ExecutionUrl string `json:"executionUrl,omitempty"` //
	Message      string `json:"message,omitempty"`      //
}

// CreateWirelessProfileRequest is the CreateWirelessProfileRequest definition
type CreateWirelessProfileRequest struct {
	ProfileDetails ProfileDetails `json:"profileDetails,omitempty"` //
}

// CreateWirelessProfileResponse is the CreateWirelessProfileResponse definition
type CreateWirelessProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DATA_RATE is the DATA_RATE definition
type DATA_RATE struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// DHCP is the DHCP definition
type DHCP struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// DNS is the DNS definition
type DNS struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// DeleteEnterpriseSSIDResponse is the DeleteEnterpriseSSIDResponse definition
type DeleteEnterpriseSSIDResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteRFProfilesResponse is the DeleteRFProfilesResponse definition
type DeleteRFProfilesResponse struct {
	ExecutionId  string `json:"executionId,omitempty"`  //
	ExecutionUrl string `json:"executionUrl,omitempty"` //
	Message      string `json:"message,omitempty"`      //
}

// DeleteSSIDAndProvisionItToDevicesResponse is the DeleteSSIDAndProvisionItToDevicesResponse definition
type DeleteSSIDAndProvisionItToDevicesResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// DeleteWirelessProfileResponse is the DeleteWirelessProfileResponse definition
type DeleteWirelessProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// EMAIL is the EMAIL definition
type EMAIL struct {
	MAILSERVER MAILSERVER `json:"MAILSERVER,omitempty"` //
}

// FILETRANSFER is the FILETRANSFER definition
type FILETRANSFER struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// Failure is the Failure definition
type Failure struct {
	FailureReason string `json:"failureReason,omitempty"` //
	TaskId        string `json:"taskId,omitempty"`        //
	TaskUrl       string `json:"taskUrl,omitempty"`       //
}

// FailureStats is the FailureStats definition
type FailureStats struct {
	ErrorCode    int    `json:"errorCode,omitempty"`    //
	ErrorTitle   string `json:"errorTitle,omitempty"`   //
	TestCategory string `json:"testCategory,omitempty"` //
	TestType     string `json:"testType,omitempty"`     //
}

// FlexConnect is the FlexConnect definition
type FlexConnect struct {
	EnableFlexConnect bool `json:"enableFlexConnect,omitempty"` //
	LocalToVlan       int  `json:"localToVlan,omitempty"`       //
}

// GetEnterpriseSSIDResponse is the GetEnterpriseSSIDResponse definition
type GetEnterpriseSSIDResponse struct {
	GroupUuid          string        `json:"groupUuid,omitempty"`          //
	InheritedGroupName string        `json:"inheritedGroupName,omitempty"` //
	InheritedGroupUuid string        `json:"inheritedGroupUuid,omitempty"` //
	InstanceUuid       string        `json:"instanceUuid,omitempty"`       //
	SsidDetails        []SsidDetails `json:"ssidDetails,omitempty"`        //
	Version            int           `json:"version,omitempty"`            //
}

// HOST_REACHABILITY is the HOST_REACHABILITY definition
type HOST_REACHABILITY struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// IPSLASENDER is the IPSLASENDER definition
type IPSLASENDER struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// MAILSERVER is the MAILSERVER definition
type MAILSERVER struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// NETWORK_SERVICES is the NETWORK_SERVICES definition
type NETWORK_SERVICES struct {
	DNS DNS `json:"DNS,omitempty"` //
}

// ONBOARDING is the ONBOARDING definition
type ONBOARDING struct {
	ASSOC ASSOC `json:"ASSOC,omitempty"` //
	AUTH  AUTH  `json:"AUTH,omitempty"`  //
	DHCP  DHCP  `json:"DHCP,omitempty"`  //
}

// PERFORMANCE is the PERFORMANCE definition
type PERFORMANCE struct {
	IPSLASENDER IPSLASENDER `json:"IPSLASENDER,omitempty"` //
}

// ProfileDetails is the ProfileDetails definition
type ProfileDetails struct {
	Name        string        `json:"name,omitempty"`        //
	Sites       []string      `json:"sites,omitempty"`       //
	SsidDetails []SsidDetails `json:"ssidDetails,omitempty"` //
}

// ProvisionResponse is the ProvisionResponse definition
type ProvisionResponse struct {
	ExecutionId       string            `json:"executionId,omitempty"`       //
	ExecutionUrl      string            `json:"executionUrl,omitempty"`      //
	ProvisioningTasks ProvisioningTasks `json:"provisioningTasks,omitempty"` //
}

// ProvisionTasks is the ProvisionTasks definition
type ProvisionTasks struct {
	Failure Failure   `json:"failure,omitempty"` //
	Success []Success `json:"success,omitempty"` //
}

// ProvisionUpdateResponse is the ProvisionUpdateResponse definition
type ProvisionUpdateResponse struct {
	ExecutionId       string            `json:"executionId,omitempty"`       //
	ExecutionUrl      string            `json:"executionUrl,omitempty"`      //
	ProvisioningTasks ProvisioningTasks `json:"provisioningTasks,omitempty"` //
}

// ProvisioningTasks is the ProvisioningTasks definition
type ProvisioningTasks struct {
	Failed  []string `json:"failed,omitempty"`  //
	Success []string `json:"success,omitempty"` //
}

// RF_ASSESSMENT is the RF_ASSESSMENT definition
type RF_ASSESSMENT struct {
	DATA_RATE DATA_RATE `json:"DATA_RATE,omitempty"` //
	SNR       SNR       `json:"SNR,omitempty"`       //
}

// Response is the Response definition
type Response struct {
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

// RetrieveRFProfilesResponse is the RetrieveRFProfilesResponse definition
type RetrieveRFProfilesResponse struct {
	Response []Response `json:"response,omitempty"` //
}

// SNR is the SNR definition
type SNR struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// SensorTestResultsResponse is the SensorTestResultsResponse definition
type SensorTestResultsResponse struct {
	FailureStats []FailureStats `json:"failureStats,omitempty"` //
	Summary      Summary        `json:"summary,omitempty"`      //
}

// SsidDetails is the SsidDetails definition
type SsidDetails struct {
	EnableFabric  bool        `json:"enableFabric,omitempty"`  //
	FlexConnect   FlexConnect `json:"flexConnect,omitempty"`   //
	InterfaceName string      `json:"interfaceName,omitempty"` //
	Name          string      `json:"name,omitempty"`          //
	Type          string      `json:"type,omitempty"`          //
}

// Success is the Success definition
type Success struct {
	TaskId  string `json:"taskId,omitempty"`  //
	TaskUrl string `json:"taskUrl,omitempty"` //
}

// Summary is the Summary definition
type Summary struct {
	APP_CONNECTIVITY APP_CONNECTIVITY `json:"APP_CONNECTIVITY,omitempty"` //
	EMAIL            EMAIL            `json:"EMAIL,omitempty"`            //
	NETWORK_SERVICES NETWORK_SERVICES `json:"NETWORK_SERVICES,omitempty"` //
	ONBOARDING       ONBOARDING       `json:"ONBOARDING,omitempty"`       //
	PERFORMANCE      PERFORMANCE      `json:"PERFORMANCE,omitempty"`      //
	RF_ASSESSMENT    RF_ASSESSMENT    `json:"RF_ASSESSMENT,omitempty"`    //
	TotalTestCount   int              `json:"totalTestCount,omitempty"`   //
}

// UpdateWirelessProfileResponse is the UpdateWirelessProfileResponse definition
type UpdateWirelessProfileResponse struct {
	ExecutionId        string `json:"executionId,omitempty"`        //
	ExecutionStatusUrl string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// WEBSERVER is the WEBSERVER definition
type WEBSERVER struct {
	FailCount int `json:"failCount,omitempty"` //
	PassCount int `json:"passCount,omitempty"` //
}

// APProvision aPProvision
/* Provision wireless Access points
@param __persistbapioutput
*/
func (s *WirelessService) APProvision(aPProvisionRequest *APProvisionRequest) (*APProvisionResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/ap-provision"

	response, err := RestyClient.R().
		SetBody(aPProvisionRequest).
		SetResult(&APProvisionResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
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

	result := response.Result().(*CreateWirelessProfileResponse)
	return result, response, err

}

// DeleteEnterpriseSSID deleteEnterpriseSSID
/* Deletes given enterprise SSID
@param ssidName Enter the SSID name to be deleted
*/
func (s *WirelessService) DeleteEnterpriseSSID(ssidName string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/enterprise-ssid/{ssidName}"
	path = strings.Replace(path, "{"+"ssidName"+"}", fmt.Sprintf("%v", ssidName), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteRFProfiles deleteRFProfiles
/* Delete RF profile(s)
@param rfProfileName
*/
func (s *WirelessService) DeleteRFProfiles(rfProfileName string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/rf-profile/{rfProfileName}"
	path = strings.Replace(path, "{"+"rfProfileName"+"}", fmt.Sprintf("%v", rfProfileName), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteSSIDAndProvisionItToDevices deleteSSIDAndProvisionItToDevices
/* Removes SSID or WLAN from the network profile, reprovision the device(s) and deletes the SSID or WLAN from DNA Center
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
@param ssidName
@param managedAPLocations
*/
func (s *WirelessService) DeleteSSIDAndProvisionItToDevices(ssidName string, managedAPLocations string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/business/ssid/{ssidName}/{managedAPLocations}"
	path = strings.Replace(path, "{"+"ssidName"+"}", fmt.Sprintf("%v", ssidName), -1)
	path = strings.Replace(path, "{"+"managedAPLocations"+"}", fmt.Sprintf("%v", managedAPLocations), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteWirelessProfile deleteWirelessProfile
/* Delete the Wireless Profile from DNAC whose name is provided.
@param wirelessProfileName
*/
func (s *WirelessService) DeleteWirelessProfile(wirelessProfileName string, deleteWirelessProfileRequest *DeleteWirelessProfileRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/wireless-profile/{wirelessProfileName}"
	path = strings.Replace(path, "{"+"wirelessProfileName"+"}", fmt.Sprintf("%v", wirelessProfileName), -1)

	response, err := RestyClient.R().
		SetBody(deleteWirelessProfileRequest).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetEnterpriseSSIDQueryParams defines the query parameters for this request
type GetEnterpriseSSIDQueryParams struct {
	SsidName string `url:"ssidName,omitempty"` // Enter the enterprise SSID name that needs to be retrieved. If not entered, all the enterprise SSIDs will be retrieved.
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

	result := response.Result().(*GetWirelessProfileResponse)
	return result, response, err

}

// Provision provision
/* Provision wireless devices
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
*/
func (s *WirelessService) Provision(provisionRequest *ProvisionRequest) (*ProvisionResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/provision"

	response, err := RestyClient.R().
		SetBody(provisionRequest).
		SetResult(&ProvisionResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*ProvisionResponse)
	return result, response, err

}

// ProvisionUpdate provisionUpdate
/* Updates wireless provisioning
@param __persistbapioutput Enable this parameter to execute the API and return a response asynchronously.
*/
func (s *WirelessService) ProvisionUpdate(provisionUpdateRequest *ProvisionUpdateRequest) (*ProvisionUpdateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/wireless/provision"

	response, err := RestyClient.R().
		SetBody(provisionUpdateRequest).
		SetResult(&ProvisionUpdateResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
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

	result := response.Result().(*RetrieveRFProfilesResponse)
	return result, response, err

}

// SensorTestResultsQueryParams defines the query parameters for this request
type SensorTestResultsQueryParams struct {
	SiteId        string `url:"siteId,omitempty"`        // Assurance site UUID
	StartTime     number `url:"startTime,omitempty"`     // The epoch time in milliseconds
	EndTime       number `url:"endTime,omitempty"`       // The epoch time in milliseconds
	TestFailureBy string `url:"testFailureBy,omitempty"` // Obtain failure statistics group by "area", "building", or "floor"
}

// SensorTestResults sensorTestResults
/* Intent API to get SENSOR test result summary
@param siteId Assurance site UUID
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

	result := response.Result().(*UpdateWirelessProfileResponse)
	return result, response, err

}
