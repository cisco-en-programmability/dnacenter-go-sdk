package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SensorsService service

type DeleteSensorTestQueryParams struct {
	TemplateName string `url:"templateName,omitempty"` //
}
type SensorsQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //
}

type ResponseSensorsEditSensorTestTemplate struct {
	Version  string                                         `json:"version,omitempty"`  // Version
	Response *ResponseSensorsEditSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateResponse struct {
	TypeID                 string                                                           `json:"_id,omitempty"`                    // Id
	Name                   string                                                           `json:"name,omitempty"`                   // Name
	Version                *float64                                                         `json:"version,omitempty"`                // Version
	ModelVersion           *int                                                             `json:"modelVersion,omitempty"`           // Model Version
	StartTime              *float64                                                         `json:"startTime,omitempty"`              // Start Time
	LastModifiedTime       *float64                                                         `json:"lastModifiedTime,omitempty"`       // Last Modified Time
	NumAssociatedSensor    *float64                                                         `json:"numAssociatedSensor,omitempty"`    // Num Associated Sensor
	Location               *ResponseSensorsEditSensorTestTemplateResponseLocation           `json:"location,omitempty"`               // Location
	SiteHierarchy          *ResponseSensorsEditSensorTestTemplateResponseSiteHierarchy      `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	Status                 string                                                           `json:"status,omitempty"`                 // Status
	Connection             string                                                           `json:"connection,omitempty"`             // Connection
	Frequency              *ResponseSensorsEditSensorTestTemplateResponseFrequency          `json:"frequency,omitempty"`              // Frequency
	RssiThreshold          *int                                                             `json:"rssiThreshold,omitempty"`          // Rssi Threshold
	NumNeighborApThreshold *int                                                             `json:"numNeighborAPThreshold,omitempty"` // Num Neighbor A P Threshold
	ScheduleInDays         *float64                                                         `json:"scheduleInDays,omitempty"`         // Schedule In Days
	WLANs                  *[]ResponseSensorsEditSensorTestTemplateResponseWLANs            `json:"wlans,omitempty"`                  // Wlans
	SSIDs                  *[]ResponseSensorsEditSensorTestTemplateResponseSSIDs            `json:"ssids,omitempty"`                  //
	TestScheduleMode       string                                                           `json:"testScheduleMode,omitempty"`       // Test Schedule Mode
	ShowWlcUpgradeBanner   *bool                                                            `json:"showWlcUpgradeBanner,omitempty"`   // Show Wlc Upgrade Banner
	RadioAsSensorRemoved   *bool                                                            `json:"radioAsSensorRemoved,omitempty"`   // Radio As Sensor Removed
	EncryptionMode         string                                                           `json:"encryptionMode,omitempty"`         // Encryption Mode
	RunNow                 string                                                           `json:"runNow,omitempty"`                 // Run Now
	LocationInfoList       *[]ResponseSensorsEditSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Schedule               *ResponseSensorsEditSensorTestTemplateResponseSchedule           `json:"schedule,omitempty"`               //
	Tests                  *ResponseSensorsEditSensorTestTemplateResponseTests              `json:"tests,omitempty"`                  // Tests
	Sensors                *[]ResponseSensorsEditSensorTestTemplateResponseSensors          `json:"sensors,omitempty"`                // Sensors
	ApCoverage             *[]ResponseSensorsEditSensorTestTemplateResponseApCoverage       `json:"apCoverage,omitempty"`             //
	TestDurationEstimate   *int                                                             `json:"testDurationEstimate,omitempty"`   // Test Duration Estimate
	TestTemplate           *bool                                                            `json:"testTemplate,omitempty"`           // Test Template
	LegacyTestSuite        *bool                                                            `json:"legacyTestSuite,omitempty"`        // Legacy Test Suite
	TenantID               string                                                           `json:"tenantId,omitempty"`               // Tenant Id
}
type ResponseSensorsEditSensorTestTemplateResponseLocation interface{}
type ResponseSensorsEditSensorTestTemplateResponseSiteHierarchy interface{}
type ResponseSensorsEditSensorTestTemplateResponseFrequency interface{}
type ResponseSensorsEditSensorTestTemplateResponseWLANs interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDs struct {
	Bands                     *ResponseSensorsEditSensorTestTemplateResponseSSIDsBands                     `json:"bands,omitempty"`                     // Bands
	SSID                      string                                                                       `json:"ssid,omitempty"`                      // Ssid
	ProfileName               string                                                                       `json:"profileName,omitempty"`               // Profile Name
	AuthType                  string                                                                       `json:"authType,omitempty"`                  // Auth Type
	AuthTypeRcvd              *ResponseSensorsEditSensorTestTemplateResponseSSIDsAuthTypeRcvd              `json:"authTypeRcvd,omitempty"`              // Auth Type Rcvd
	Psk                       string                                                                       `json:"psk,omitempty"`                       // Psk
	Username                  *ResponseSensorsEditSensorTestTemplateResponseSSIDsUsername                  `json:"username,omitempty"`                  // Username
	Password                  *ResponseSensorsEditSensorTestTemplateResponseSSIDsPassword                  `json:"password,omitempty"`                  // Password
	EapMethod                 *ResponseSensorsEditSensorTestTemplateResponseSSIDsEapMethod                 `json:"eapMethod,omitempty"`                 // Eap Method
	Scep                      *bool                                                                        `json:"scep,omitempty"`                      // Scep
	AuthProtocol              *ResponseSensorsEditSensorTestTemplateResponseSSIDsAuthProtocol              `json:"authProtocol,omitempty"`              // Auth Protocol
	Certfilename              *ResponseSensorsEditSensorTestTemplateResponseSSIDsCertfilename              `json:"certfilename,omitempty"`              // Certfilename
	Certxferprotocol          string                                                                       `json:"certxferprotocol,omitempty"`          // Certxferprotocol
	Certstatus                string                                                                       `json:"certstatus,omitempty"`                // Certstatus
	Certpassphrase            *ResponseSensorsEditSensorTestTemplateResponseSSIDsCertpassphrase            `json:"certpassphrase,omitempty"`            // Certpassphrase
	Certdownloadurl           *ResponseSensorsEditSensorTestTemplateResponseSSIDsCertdownloadurl           `json:"certdownloadurl,omitempty"`           // Certdownloadurl
	NumAps                    *float64                                                                     `json:"numAps,omitempty"`                    // Num Aps
	NumSensors                *float64                                                                     `json:"numSensors,omitempty"`                // Num Sensors
	Layer3WebAuthsecurity     *ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity     `json:"layer3webAuthsecurity,omitempty"`     // Layer3web Authsecurity
	Layer3WebAuthuserName     *ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthuserName     `json:"layer3webAuthuserName,omitempty"`     // Layer3web Authuser Name
	Layer3WebAuthpassword     *ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthpassword     `json:"layer3webAuthpassword,omitempty"`     // Layer3web Authpassword
	ExtWebAuthVirtualIP       *ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP       `json:"extWebAuthVirtualIp,omitempty"`       // Ext Web Auth Virtual Ip
	Layer3WebAuthEmailAddress *ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress `json:"layer3webAuthEmailAddress,omitempty"` // Layer3web Auth Email Address
	QosPolicy                 string                                                                       `json:"qosPolicy,omitempty"`                 // Qos Policy
	ExtWebAuth                *bool                                                                        `json:"extWebAuth,omitempty"`                // Ext Web Auth
	WhiteList                 *bool                                                                        `json:"whiteList,omitempty"`                 // White List
	ExtWebAuthPortal          *ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthPortal          `json:"extWebAuthPortal,omitempty"`          // Ext Web Auth Portal
	ExtWebAuthAccessURL       *ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthAccessURL       `json:"extWebAuthAccessUrl,omitempty"`       // Ext Web Auth Access Url
	ExtWebAuthHTMLTag         *[]ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag       `json:"extWebAuthHtmlTag,omitempty"`         // Ext Web Auth Html Tag
	ThirdParty                *ResponseSensorsEditSensorTestTemplateResponseSSIDsThirdParty                `json:"thirdParty,omitempty"`                //
	ID                        *float64                                                                     `json:"id,omitempty"`                        // Id
	WLANID                    *float64                                                                     `json:"wlanId,omitempty"`                    // Wlan Id
	Wlc                       *ResponseSensorsEditSensorTestTemplateResponseSSIDsWlc                       `json:"wlc,omitempty"`                       // Wlc
	ValidFrom                 *float64                                                                     `json:"validFrom,omitempty"`                 // Valid From
	ValidTo                   *float64                                                                     `json:"validTo,omitempty"`                   // Valid To
	Status                    string                                                                       `json:"status,omitempty"`                    // Status
	Tests                     *[]ResponseSensorsEditSensorTestTemplateResponseSSIDsTests                   `json:"tests,omitempty"`                     //
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsBands interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsAuthTypeRcvd interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsUsername interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsPassword interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsEapMethod interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsAuthProtocol interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsCertfilename interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsCertpassphrase interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsCertdownloadurl interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthuserName interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthpassword interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthPortal interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthAccessURL interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // Selected
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsWlc interface{}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsTests struct {
	Name   string                                                           `json:"name,omitempty"`   // Name
	Config *[]ResponseSensorsEditSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` // Config
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsTestsConfig interface{}
type ResponseSensorsEditSensorTestTemplateResponseLocationInfoList struct {
	LocationID     string                                                                         `json:"locationId,omitempty"`     // Location Id
	LocationType   string                                                                         `json:"locationType,omitempty"`   // Location Type
	AllSensors     *bool                                                                          `json:"allSensors,omitempty"`     // All Sensors
	SiteHierarchy  string                                                                         `json:"siteHierarchy,omitempty"`  // Site Hierarchy
	MacAddressList *[]ResponseSensorsEditSensorTestTemplateResponseLocationInfoListMacAddressList `json:"macAddressList,omitempty"` // Mac Address List
}
type ResponseSensorsEditSensorTestTemplateResponseLocationInfoListMacAddressList interface{}
type ResponseSensorsEditSensorTestTemplateResponseSchedule struct {
	TestScheduleMode string                                                                `json:"testScheduleMode,omitempty"` // Test Schedule Mode
	ScheduleRange    *[]ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRange `json:"scheduleRange,omitempty"`    //
	StartTime        *float64                                                              `json:"startTime,omitempty"`        // Start Time
	Frequency        *ResponseSensorsEditSensorTestTemplateResponseScheduleFrequency       `json:"frequency,omitempty"`        //
}
type ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRange struct {
	TimeRange *[]ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRangeTimeRange `json:"timeRange,omitempty"` //
	Day       string                                                                         `json:"day,omitempty"`       // Day
}
type ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRangeTimeRange struct {
	From      string                                                                                `json:"from,omitempty"`      // From
	To        string                                                                                `json:"to,omitempty"`        // To
	Frequency *ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRangeTimeRangeFrequency `json:"frequency,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateResponseScheduleScheduleRangeTimeRangeFrequency struct {
	Value *int   `json:"value,omitempty"` // Value
	Unit  string `json:"unit,omitempty"`  // Unit
}
type ResponseSensorsEditSensorTestTemplateResponseScheduleFrequency struct {
	Value *int   `json:"value,omitempty"` // Value
	Unit  string `json:"unit,omitempty"`  // Unit
}
type ResponseSensorsEditSensorTestTemplateResponseTests interface{}
type ResponseSensorsEditSensorTestTemplateResponseSensors interface{}
type ResponseSensorsEditSensorTestTemplateResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // Bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number Of Aps To Test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // Rssi Threshold
}
type ResponseSensorsCreateSensorTestTemplate struct {
	Version  string                                           `json:"version,omitempty"`  // Version
	Response *ResponseSensorsCreateSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponse struct {
	TypeID                 string                                                             `json:"_id,omitempty"`                    // Id
	Name                   string                                                             `json:"name,omitempty"`                   // Name
	Version                *float64                                                           `json:"version,omitempty"`                // Version
	ModelVersion           *int                                                               `json:"modelVersion,omitempty"`           // Model Version
	StartTime              *float64                                                           `json:"startTime,omitempty"`              // Start Time
	LastModifiedTime       *float64                                                           `json:"lastModifiedTime,omitempty"`       // Last Modified Time
	NumAssociatedSensor    *float64                                                           `json:"numAssociatedSensor,omitempty"`    // Num Associated Sensor
	Location               *ResponseSensorsCreateSensorTestTemplateResponseLocation           `json:"location,omitempty"`               // Location
	SiteHierarchy          *ResponseSensorsCreateSensorTestTemplateResponseSiteHierarchy      `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	Status                 string                                                             `json:"status,omitempty"`                 // Status
	Connection             string                                                             `json:"connection,omitempty"`             // Connection
	Frequency              *ResponseSensorsCreateSensorTestTemplateResponseFrequency          `json:"frequency,omitempty"`              // Frequency
	RssiThreshold          *int                                                               `json:"rssiThreshold,omitempty"`          // Rssi Threshold
	NumNeighborApThreshold *int                                                               `json:"numNeighborAPThreshold,omitempty"` // Num Neighbor A P Threshold
	ScheduleInDays         *float64                                                           `json:"scheduleInDays,omitempty"`         // Schedule In Days
	WLANs                  *[]ResponseSensorsCreateSensorTestTemplateResponseWLANs            `json:"wlans,omitempty"`                  // Wlans
	SSIDs                  *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDs            `json:"ssids,omitempty"`                  //
	TestScheduleMode       string                                                             `json:"testScheduleMode,omitempty"`       // Test Schedule Mode
	ShowWlcUpgradeBanner   *bool                                                              `json:"showWlcUpgradeBanner,omitempty"`   // Show Wlc Upgrade Banner
	RadioAsSensorRemoved   *bool                                                              `json:"radioAsSensorRemoved,omitempty"`   // Radio As Sensor Removed
	EncryptionMode         string                                                             `json:"encryptionMode,omitempty"`         // Encryption Mode
	RunNow                 string                                                             `json:"runNow,omitempty"`                 // Run Now
	LocationInfoList       *[]ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"`       // Location Info List
	Schedule               *ResponseSensorsCreateSensorTestTemplateResponseSchedule           `json:"schedule,omitempty"`               // Schedule
	Tests                  *ResponseSensorsCreateSensorTestTemplateResponseTests              `json:"tests,omitempty"`                  // Tests
	Sensors                *[]ResponseSensorsCreateSensorTestTemplateResponseSensors          `json:"sensors,omitempty"`                // Sensors
	ApCoverage             *[]ResponseSensorsCreateSensorTestTemplateResponseApCoverage       `json:"apCoverage,omitempty"`             //
	TestDurationEstimate   *int                                                               `json:"testDurationEstimate,omitempty"`   // Test Duration Estimate
	TestTemplate           *bool                                                              `json:"testTemplate,omitempty"`           // Test Template
	LegacyTestSuite        *bool                                                              `json:"legacyTestSuite,omitempty"`        // Legacy Test Suite
	TenantID               *ResponseSensorsCreateSensorTestTemplateResponseTenantID           `json:"tenantId,omitempty"`               // Tenant Id
}
type ResponseSensorsCreateSensorTestTemplateResponseLocation interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSiteHierarchy interface{}
type ResponseSensorsCreateSensorTestTemplateResponseFrequency interface{}
type ResponseSensorsCreateSensorTestTemplateResponseWLANs interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDs struct {
	Bands                     *ResponseSensorsCreateSensorTestTemplateResponseSSIDsBands                     `json:"bands,omitempty"`                     // Bands
	SSID                      string                                                                         `json:"ssid,omitempty"`                      // Ssid
	ProfileName               string                                                                         `json:"profileName,omitempty"`               // Profile Name
	AuthType                  string                                                                         `json:"authType,omitempty"`                  // Auth Type
	AuthTypeRcvd              *ResponseSensorsCreateSensorTestTemplateResponseSSIDsAuthTypeRcvd              `json:"authTypeRcvd,omitempty"`              // Auth Type Rcvd
	Psk                       string                                                                         `json:"psk,omitempty"`                       // Psk
	Username                  *ResponseSensorsCreateSensorTestTemplateResponseSSIDsUsername                  `json:"username,omitempty"`                  // Username
	Password                  *ResponseSensorsCreateSensorTestTemplateResponseSSIDsPassword                  `json:"password,omitempty"`                  // Password
	EapMethod                 *ResponseSensorsCreateSensorTestTemplateResponseSSIDsEapMethod                 `json:"eapMethod,omitempty"`                 // Eap Method
	Scep                      *bool                                                                          `json:"scep,omitempty"`                      // Scep
	AuthProtocol              *ResponseSensorsCreateSensorTestTemplateResponseSSIDsAuthProtocol              `json:"authProtocol,omitempty"`              // Auth Protocol
	Certfilename              *ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertfilename              `json:"certfilename,omitempty"`              // Certfilename
	Certxferprotocol          string                                                                         `json:"certxferprotocol,omitempty"`          // Certxferprotocol
	Certstatus                string                                                                         `json:"certstatus,omitempty"`                // Certstatus
	Certpassphrase            *ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertpassphrase            `json:"certpassphrase,omitempty"`            // Certpassphrase
	Certdownloadurl           *ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertdownloadurl           `json:"certdownloadurl,omitempty"`           // Certdownloadurl
	NumAps                    *float64                                                                       `json:"numAps,omitempty"`                    // Num Aps
	NumSensors                *float64                                                                       `json:"numSensors,omitempty"`                // Num Sensors
	Layer3WebAuthsecurity     *ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity     `json:"layer3webAuthsecurity,omitempty"`     // Layer3web Authsecurity
	Layer3WebAuthuserName     *ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthuserName     `json:"layer3webAuthuserName,omitempty"`     // Layer3web Authuser Name
	Layer3WebAuthpassword     *ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthpassword     `json:"layer3webAuthpassword,omitempty"`     // Layer3web Authpassword
	ExtWebAuthVirtualIP       *ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP       `json:"extWebAuthVirtualIp,omitempty"`       // Ext Web Auth Virtual Ip
	Layer3WebAuthEmailAddress *ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress `json:"layer3webAuthEmailAddress,omitempty"` // Layer3web Auth Email Address
	QosPolicy                 string                                                                         `json:"qosPolicy,omitempty"`                 // Qos Policy
	ExtWebAuth                *bool                                                                          `json:"extWebAuth,omitempty"`                // Ext Web Auth
	WhiteList                 *bool                                                                          `json:"whiteList,omitempty"`                 // White List
	ExtWebAuthPortal          *ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthPortal          `json:"extWebAuthPortal,omitempty"`          // Ext Web Auth Portal
	ExtWebAuthAccessURL       *ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthAccessURL       `json:"extWebAuthAccessUrl,omitempty"`       // Ext Web Auth Access Url
	ExtWebAuthHTMLTag         *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag       `json:"extWebAuthHtmlTag,omitempty"`         // Ext Web Auth Html Tag
	ThirdParty                *ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty                `json:"thirdParty,omitempty"`                //
	ID                        *float64                                                                       `json:"id,omitempty"`                        // Id
	WLANID                    *float64                                                                       `json:"wlanId,omitempty"`                    // Wlan Id
	Wlc                       *ResponseSensorsCreateSensorTestTemplateResponseSSIDsWlc                       `json:"wlc,omitempty"`                       // Wlc
	ValidFrom                 *float64                                                                       `json:"validFrom,omitempty"`                 // Valid From
	ValidTo                   *float64                                                                       `json:"validTo,omitempty"`                   // Valid To
	Status                    string                                                                         `json:"status,omitempty"`                    // Status
	Tests                     *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests                   `json:"tests,omitempty"`                     //
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsBands interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsAuthTypeRcvd interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsUsername interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsPassword interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsEapMethod interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsAuthProtocol interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertfilename interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertpassphrase interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsCertdownloadurl interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthuserName interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthpassword interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthPortal interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthAccessURL interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // Selected
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsWlc interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests struct {
	Name   string                                                             `json:"name,omitempty"`   // Name
	Config *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` // Config
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig interface{}
type ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSchedule interface{}
type ResponseSensorsCreateSensorTestTemplateResponseTests interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSensors interface{}
type ResponseSensorsCreateSensorTestTemplateResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // Bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number Of Aps To Test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // Rssi Threshold
}
type ResponseSensorsCreateSensorTestTemplateResponseTenantID interface{}
type ResponseSensorsDeleteSensorTest struct {
	Version  string                                   `json:"version,omitempty"`  // Version
	Response *ResponseSensorsDeleteSensorTestResponse `json:"response,omitempty"` //
}
type ResponseSensorsDeleteSensorTestResponse struct {
	TemplateName string `json:"templateName,omitempty"` // Template Name
	Status       string `json:"status,omitempty"`       // Status
}
type ResponseSensorsSensors struct {
	Version  string                            `json:"version,omitempty"`  // Version
	Response *[]ResponseSensorsSensorsResponse `json:"response,omitempty"` //
}
type ResponseSensorsSensorsResponse struct {
	Name               string                                   `json:"name,omitempty"`               // Name
	Status             string                                   `json:"status,omitempty"`             // Status
	RadioMacAddress    string                                   `json:"radioMacAddress,omitempty"`    // Radio Mac Address
	EthernetMacAddress string                                   `json:"ethernetMacAddress,omitempty"` // Ethernet Mac Address
	Location           string                                   `json:"location,omitempty"`           // Location
	BackhaulType       string                                   `json:"backhaulType,omitempty"`       // Backhaul Type
	SerialNumber       string                                   `json:"serialNumber,omitempty"`       // Serial Number
	IPAddress          string                                   `json:"ipAddress,omitempty"`          // Ip Address
	Version            string                                   `json:"version,omitempty"`            // Version
	LastSeen           *int                                     `json:"lastSeen,omitempty"`           // Last Seen
	Type               string                                   `json:"type,omitempty"`               // Type
	SSHConfig          *ResponseSensorsSensorsResponseSSHConfig `json:"sshConfig,omitempty"`          //
	IsLEDEnabled       *bool                                    `json:"isLEDEnabled,omitempty"`       // Is L E D Enabled
}
type ResponseSensorsSensorsResponseSSHConfig struct {
	SSHState       string `json:"sshState,omitempty"`       // Ssh State
	SSHUserName    string `json:"sshUserName,omitempty"`    // Ssh User Name
	SSHPassword    string `json:"sshPassword,omitempty"`    // Ssh Password
	EnablePassword string `json:"enablePassword,omitempty"` // Enable Password
}
type ResponseSensorsDuplicateSensorTestTemplate struct {
	Version  string                                              `json:"version,omitempty"`  // Version
	Response *ResponseSensorsDuplicateSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponse struct {
	TypeID                 string                                                                `json:"_id,omitempty"`                    // Id
	Name                   string                                                                `json:"name,omitempty"`                   // Name
	Version                *float64                                                              `json:"version,omitempty"`                // Version
	ModelVersion           *int                                                                  `json:"modelVersion,omitempty"`           // Model Version
	StartTime              *float64                                                              `json:"startTime,omitempty"`              // Start Time
	LastModifiedTime       *float64                                                              `json:"lastModifiedTime,omitempty"`       // Last Modified Time
	NumAssociatedSensor    *float64                                                              `json:"numAssociatedSensor,omitempty"`    // Num Associated Sensor
	Location               *ResponseSensorsDuplicateSensorTestTemplateResponseLocation           `json:"location,omitempty"`               // Location
	SiteHierarchy          *ResponseSensorsDuplicateSensorTestTemplateResponseSiteHierarchy      `json:"siteHierarchy,omitempty"`          // Site Hierarchy
	Status                 string                                                                `json:"status,omitempty"`                 // Status
	Connection             string                                                                `json:"connection,omitempty"`             // Connection
	Frequency              *ResponseSensorsDuplicateSensorTestTemplateResponseFrequency          `json:"frequency,omitempty"`              // Frequency
	RssiThreshold          *int                                                                  `json:"rssiThreshold,omitempty"`          // Rssi Threshold
	NumNeighborApThreshold *int                                                                  `json:"numNeighborAPThreshold,omitempty"` // Num Neighbor A P Threshold
	ScheduleInDays         *float64                                                              `json:"scheduleInDays,omitempty"`         // Schedule In Days
	WLANs                  *[]ResponseSensorsDuplicateSensorTestTemplateResponseWLANs            `json:"wlans,omitempty"`                  // Wlans
	SSIDs                  *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs            `json:"ssids,omitempty"`                  //
	TestScheduleMode       string                                                                `json:"testScheduleMode,omitempty"`       // Test Schedule Mode
	ShowWlcUpgradeBanner   *bool                                                                 `json:"showWlcUpgradeBanner,omitempty"`   // Show Wlc Upgrade Banner
	RadioAsSensorRemoved   *bool                                                                 `json:"radioAsSensorRemoved,omitempty"`   // Radio As Sensor Removed
	EncryptionMode         string                                                                `json:"encryptionMode,omitempty"`         // Encryption Mode
	RunNow                 string                                                                `json:"runNow,omitempty"`                 // Run Now
	LocationInfoList       *[]ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Schedule               *ResponseSensorsDuplicateSensorTestTemplateResponseSchedule           `json:"schedule,omitempty"`               //
	Tests                  *ResponseSensorsDuplicateSensorTestTemplateResponseTests              `json:"tests,omitempty"`                  // Tests
	Sensors                *[]ResponseSensorsDuplicateSensorTestTemplateResponseSensors          `json:"sensors,omitempty"`                // Sensors
	ApCoverage             *[]ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage       `json:"apCoverage,omitempty"`             //
	TestDurationEstimate   *int                                                                  `json:"testDurationEstimate,omitempty"`   // Test Duration Estimate
	TestTemplate           *bool                                                                 `json:"testTemplate,omitempty"`           // Test Template
	LegacyTestSuite        *bool                                                                 `json:"legacyTestSuite,omitempty"`        // Legacy Test Suite
	TenantID               *ResponseSensorsDuplicateSensorTestTemplateResponseTenantID           `json:"tenantId,omitempty"`               // Tenant Id
}
type ResponseSensorsDuplicateSensorTestTemplateResponseLocation interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSiteHierarchy interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseFrequency interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseWLANs interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs struct {
	Bands                     *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsBands                     `json:"bands,omitempty"`                     // Bands
	SSID                      string                                                                            `json:"ssid,omitempty"`                      // Ssid
	ProfileName               string                                                                            `json:"profileName,omitempty"`               // Profile Name
	AuthType                  string                                                                            `json:"authType,omitempty"`                  // Auth Type
	AuthTypeRcvd              *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsAuthTypeRcvd              `json:"authTypeRcvd,omitempty"`              // Auth Type Rcvd
	Psk                       string                                                                            `json:"psk,omitempty"`                       // Psk
	Username                  *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsUsername                  `json:"username,omitempty"`                  // Username
	Password                  *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsPassword                  `json:"password,omitempty"`                  // Password
	EapMethod                 *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsEapMethod                 `json:"eapMethod,omitempty"`                 // Eap Method
	Scep                      *bool                                                                             `json:"scep,omitempty"`                      // Scep
	AuthProtocol              *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsAuthProtocol              `json:"authProtocol,omitempty"`              // Auth Protocol
	Certfilename              *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertfilename              `json:"certfilename,omitempty"`              // Certfilename
	Certxferprotocol          string                                                                            `json:"certxferprotocol,omitempty"`          // Certxferprotocol
	Certstatus                string                                                                            `json:"certstatus,omitempty"`                // Certstatus
	Certpassphrase            *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertpassphrase            `json:"certpassphrase,omitempty"`            // Certpassphrase
	Certdownloadurl           *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertdownloadurl           `json:"certdownloadurl,omitempty"`           // Certdownloadurl
	NumAps                    *float64                                                                          `json:"numAps,omitempty"`                    // Num Aps
	NumSensors                *float64                                                                          `json:"numSensors,omitempty"`                // Num Sensors
	Layer3WebAuthsecurity     *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity     `json:"layer3webAuthsecurity,omitempty"`     // Layer3web Authsecurity
	Layer3WebAuthuserName     *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthuserName     `json:"layer3webAuthuserName,omitempty"`     // Layer3web Authuser Name
	Layer3WebAuthpassword     *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthpassword     `json:"layer3webAuthpassword,omitempty"`     // Layer3web Authpassword
	ExtWebAuthVirtualIP       *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP       `json:"extWebAuthVirtualIp,omitempty"`       // Ext Web Auth Virtual Ip
	Layer3WebAuthEmailAddress *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress `json:"layer3webAuthEmailAddress,omitempty"` // Layer3web Auth Email Address
	QosPolicy                 string                                                                            `json:"qosPolicy,omitempty"`                 // Qos Policy
	ExtWebAuth                *bool                                                                             `json:"extWebAuth,omitempty"`                // Ext Web Auth
	WhiteList                 *bool                                                                             `json:"whiteList,omitempty"`                 // White List
	ExtWebAuthPortal          *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthPortal          `json:"extWebAuthPortal,omitempty"`          // Ext Web Auth Portal
	ExtWebAuthAccessURL       *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthAccessURL       `json:"extWebAuthAccessUrl,omitempty"`       // Ext Web Auth Access Url
	ExtWebAuthHTMLTag         *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag       `json:"extWebAuthHtmlTag,omitempty"`         // Ext Web Auth Html Tag
	ThirdParty                *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty                `json:"thirdParty,omitempty"`                //
	ID                        *float64                                                                          `json:"id,omitempty"`                        // Id
	WLANID                    *float64                                                                          `json:"wlanId,omitempty"`                    // Wlan Id
	Wlc                       *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsWlc                       `json:"wlc,omitempty"`                       // Wlc
	ValidFrom                 *float64                                                                          `json:"validFrom,omitempty"`                 // Valid From
	ValidTo                   *float64                                                                          `json:"validTo,omitempty"`                   // Valid To
	Status                    string                                                                            `json:"status,omitempty"`                    // Status
	Tests                     *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests                   `json:"tests,omitempty"`                     //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsBands interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsAuthTypeRcvd interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsUsername interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsPassword interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsEapMethod interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsAuthProtocol interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertfilename interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertpassphrase interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsCertdownloadurl interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthsecurity interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthuserName interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthpassword interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthVirtualIP interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsLayer3WebAuthEmailAddress interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthPortal interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthAccessURL interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // Selected
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsWlc interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests struct {
	Name   string                                                                `json:"name,omitempty"`   // Name
	Config *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` // Config
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList struct {
	LocationID     string                                                                              `json:"locationId,omitempty"`     // Location Id
	LocationType   string                                                                              `json:"locationType,omitempty"`   // Location Type
	AllSensors     *bool                                                                               `json:"allSensors,omitempty"`     // All Sensors
	SiteHierarchy  string                                                                              `json:"siteHierarchy,omitempty"`  // Site Hierarchy
	MacAddressList *[]ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoListMacAddressList `json:"macAddressList,omitempty"` // Mac Address List
}
type ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoListMacAddressList interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSchedule struct {
	TestScheduleMode string                                                                     `json:"testScheduleMode,omitempty"` // Test Schedule Mode
	ScheduleRange    *[]ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRange `json:"scheduleRange,omitempty"`    //
	StartTime        *float64                                                                   `json:"startTime,omitempty"`        // Start Time
	Frequency        *ResponseSensorsDuplicateSensorTestTemplateResponseScheduleFrequency       `json:"frequency,omitempty"`        //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRange struct {
	TimeRange *[]ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRangeTimeRange `json:"timeRange,omitempty"` //
	Day       string                                                                              `json:"day,omitempty"`       // Day
}
type ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRangeTimeRange struct {
	From      string                                                                                     `json:"from,omitempty"`      // From
	To        string                                                                                     `json:"to,omitempty"`        // To
	Frequency *ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRangeTimeRangeFrequency `json:"frequency,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseScheduleScheduleRangeTimeRangeFrequency struct {
	Value *int   `json:"value,omitempty"` // Value
	Unit  string `json:"unit,omitempty"`  // Unit
}
type ResponseSensorsDuplicateSensorTestTemplateResponseScheduleFrequency struct {
	Value *int   `json:"value,omitempty"` // Value
	Unit  string `json:"unit,omitempty"`  // Unit
}
type ResponseSensorsDuplicateSensorTestTemplateResponseTests interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensors interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // Bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number Of Aps To Test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // Rssi Threshold
}
type ResponseSensorsDuplicateSensorTestTemplateResponseTenantID interface{}
type RequestSensorsEditSensorTestTemplate struct {
	TemplateName     string                                                  `json:"templateName,omitempty"`     // Template Name
	LocationInfoList *[]RequestSensorsEditSensorTestTemplateLocationInfoList `json:"locationInfoList,omitempty"` //
	Schedule         *RequestSensorsEditSensorTestTemplateSchedule           `json:"schedule,omitempty"`         //
}
type RequestSensorsEditSensorTestTemplateLocationInfoList struct {
	LocationID    string `json:"locationId,omitempty"`    // Location Id
	LocationType  string `json:"locationType,omitempty"`  // Location Type
	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site Hierarchy
	AllSensors    *bool  `json:"allSensors,omitempty"`    // All Sensors
}
type RequestSensorsEditSensorTestTemplateSchedule struct {
	TestScheduleMode string                                                       `json:"testScheduleMode,omitempty"` // Test Schedule Mode
	Frequency        *RequestSensorsEditSensorTestTemplateScheduleFrequency       `json:"frequency,omitempty"`        //
	ScheduleRange    *[]RequestSensorsEditSensorTestTemplateScheduleScheduleRange `json:"scheduleRange,omitempty"`    //
}
type RequestSensorsEditSensorTestTemplateScheduleFrequency struct {
	Unit  string `json:"unit,omitempty"`  // Unit
	Value *int   `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateScheduleScheduleRange struct {
	Day       string                                                                `json:"day,omitempty"`       // Day
	TimeRange *[]RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRange `json:"timeRange,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRange struct {
	From      string                                                                       `json:"from,omitempty"`      // From
	To        string                                                                       `json:"to,omitempty"`        // To
	Frequency *RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRangeFrequency `json:"frequency,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateScheduleScheduleRangeTimeRangeFrequency struct {
	Unit  string `json:"unit,omitempty"`  // Unit
	Value *int   `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplate struct {
	SSIDs        *[]RequestSensorsCreateSensorTestTemplateSSIDs      `json:"ssids,omitempty"`        //
	Name         string                                              `json:"name,omitempty"`         // Name
	Connection   string                                              `json:"connection,omitempty"`   // Connection
	ApCoverage   *[]RequestSensorsCreateSensorTestTemplateApCoverage `json:"apCoverage,omitempty"`   //
	ModelVersion *int                                                `json:"modelVersion,omitempty"` // Model Version
}
type RequestSensorsCreateSensorTestTemplateSSIDs struct {
	SSID        string                                                 `json:"ssid,omitempty"`        // Ssid
	ProfileName string                                                 `json:"profileName,omitempty"` // Profile Name
	AuthType    string                                                 `json:"authType,omitempty"`    // Auth Type
	ThirdParty  *RequestSensorsCreateSensorTestTemplateSSIDsThirdParty `json:"thirdParty,omitempty"`  //
	Psk         string                                                 `json:"psk,omitempty"`         // Psk
	Tests       *[]RequestSensorsCreateSensorTestTemplateSSIDsTests    `json:"tests,omitempty"`       //
	Categories  []string                                               `json:"categories,omitempty"`  // Categories
	QosPolicy   string                                                 `json:"qosPolicy,omitempty"`   // Qos Policy
}
type RequestSensorsCreateSensorTestTemplateSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // Selected
}
type RequestSensorsCreateSensorTestTemplateSSIDsTests struct {
	Name   string                                                    `json:"name,omitempty"`   // Name
	Config *[]RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig `json:"config,omitempty"` // Config
}
type RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig interface{}
type RequestSensorsCreateSensorTestTemplateApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // Bands
	NumberOfApsToTest string `json:"numberOfApsToTest,omitempty"` // Number Of Aps To Test
	RssiThreshold     string `json:"rssiThreshold,omitempty"`     // Rssi Threshold
}
type RequestSensorsRunNowSensorTest struct {
	TemplateName string `json:"templateName,omitempty"` // Template Name
}
type RequestSensorsDuplicateSensorTestTemplate struct {
	TemplateName    string `json:"templateName,omitempty"`    // Template Name
	NewTemplateName string `json:"newTemplateName,omitempty"` // New Template Name
}

//Sensors Sensors - 71a1-2bb7-4569-9cc5
/* Intent API to get a list of SENSOR devices


@param SensorsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!sensors
*/
func (s *SensorsService) Sensors(SensorsQueryParams *SensorsQueryParams) (*ResponseSensorsSensors, *resty.Response, error) {
	path := "/dna/intent/api/v1/sensor"

	queryString, _ := query.Values(SensorsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsSensors{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.Sensors(SensorsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation Sensors")
	}

	result := response.Result().(*ResponseSensorsSensors)
	return result, response, err

}

//CreateSensorTestTemplate Create sensor test template - 08bd-8883-4a68-a2e6
/* Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-sensor-test-template
*/
func (s *SensorsService) CreateSensorTestTemplate(requestSensorsCreateSensorTestTemplate *RequestSensorsCreateSensorTestTemplate) (*ResponseSensorsCreateSensorTestTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/sensor"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsCreateSensorTestTemplate).
		SetResult(&ResponseSensorsCreateSensorTestTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateSensorTestTemplate(requestSensorsCreateSensorTestTemplate)
		}

		return nil, response, fmt.Errorf("error with operation CreateSensorTestTemplate")
	}

	result := response.Result().(*ResponseSensorsCreateSensorTestTemplate)
	return result, response, err

}

//EditSensorTestTemplate Edit sensor test template - c085-eaf5-4f89-ba34
/* Intent API to deploy, schedule, or edit and existing SENSOR test template


 */
func (s *SensorsService) EditSensorTestTemplate(requestSensorsEditSensorTestTemplate *RequestSensorsEditSensorTestTemplate) (*ResponseSensorsEditSensorTestTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/AssuranceScheduleSensorTest"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsEditSensorTestTemplate).
		SetResult(&ResponseSensorsEditSensorTestTemplate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditSensorTestTemplate(requestSensorsEditSensorTestTemplate)
		}
		return nil, response, fmt.Errorf("error with operation EditSensorTestTemplate")
	}

	result := response.Result().(*ResponseSensorsEditSensorTestTemplate)
	return result, response, err

}

//RunNowSensorTest Run now sensor test - f1a7-a8e7-4cf9-9c8f
/* Intent API to run a deployed SENSOR test


 */
func (s *SensorsService) RunNowSensorTest(requestSensorsRunNowSensorTest *RequestSensorsRunNowSensorTest) (*resty.Response, error) {
	path := "/dna/intent/api/v1/sensor-run-now"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsRunNowSensorTest).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RunNowSensorTest(requestSensorsRunNowSensorTest)
		}
		return response, fmt.Errorf("error with operation RunNowSensorTest")
	}

	return response, err

}

//DuplicateSensorTestTemplate Duplicate sensor test template - 85a2-8837-4909-9021
/* Intent API to duplicate an existing SENSOR test template


 */
func (s *SensorsService) DuplicateSensorTestTemplate(requestSensorsDuplicateSensorTestTemplate *RequestSensorsDuplicateSensorTestTemplate) (*ResponseSensorsDuplicateSensorTestTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/sensorTestTemplate"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsDuplicateSensorTestTemplate).
		SetResult(&ResponseSensorsDuplicateSensorTestTemplate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DuplicateSensorTestTemplate(requestSensorsDuplicateSensorTestTemplate)
		}
		return nil, response, fmt.Errorf("error with operation DuplicateSensorTestTemplate")
	}

	result := response.Result().(*ResponseSensorsDuplicateSensorTestTemplate)
	return result, response, err

}

//DeleteSensorTest Delete sensor test - 5bbb-28ff-442a-825f
/* Intent API to delete an existing SENSOR test template


@param DeleteSensorTestQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-sensor-test
*/
func (s *SensorsService) DeleteSensorTest(DeleteSensorTestQueryParams *DeleteSensorTestQueryParams) (*ResponseSensorsDeleteSensorTest, *resty.Response, error) {
	//DeleteSensorTestQueryParams *DeleteSensorTestQueryParams
	path := "/dna/intent/api/v1/sensor"

	queryString, _ := query.Values(DeleteSensorTestQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsDeleteSensorTest{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteSensorTest(DeleteSensorTestQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteSensorTest")
	}

	result := response.Result().(*ResponseSensorsDeleteSensorTest)
	return result, response, err

}
