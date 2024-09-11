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
	Name                   string                                                           `json:"name,omitempty"`                   // The sensor test template name
	TypeID                 string                                                           `json:"_id,omitempty"`                    // The sensor test template unique identifier
	Version                *int                                                             `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                             `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                             `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                             `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                             `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                           `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                           `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                           `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                           `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                           `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *ResponseSensorsEditSensorTestTemplateResponseFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                             `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                             `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                             `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                         `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]ResponseSensorsEditSensorTestTemplateResponseSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]ResponseSensorsEditSensorTestTemplateResponseProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                           `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                            `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                            `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                           `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                           `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]ResponseSensorsEditSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]ResponseSensorsEditSensorTestTemplateResponseSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]ResponseSensorsEditSensorTestTemplateResponseApCoverage       `json:"apCoverage,omitempty"`             //
}
type ResponseSensorsEditSensorTestTemplateResponseFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDs struct {
	Bands                     string                                                                 `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                                 `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                                 `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                                   `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                                   `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                                 `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                                 `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                                 `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                                 `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *ResponseSensorsEditSensorTestTemplateResponseSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                                   `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                                   `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                                 `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                                   `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                                   `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                                 `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                                 `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                                 `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                                 `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                                 `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                                 `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                                 `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                                 `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                                 `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                                 `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                                 `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                                  `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                                 `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                                 `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                                 `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                                 `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                                 `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                                 `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                                 `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                                  `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                                  `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                                 `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                                 `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                                 `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]ResponseSensorsEditSensorTestTemplateResponseSSIDsTests             `json:"tests,omitempty"`                     //
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsTests struct {
	Name   string                                                           `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsEditSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateResponseSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsEditSensorTestTemplateResponseProfiles struct {
	AuthType            string                                                                    `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                                    `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                                    `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                                    `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                                    `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                                    `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                     `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                                    `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                                    `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                                    `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                                    `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                                    `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                                    `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                                    `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                     `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                     `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                                    `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                                    `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]ResponseSensorsEditSensorTestTemplateResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                                    `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]ResponseSensorsEditSensorTestTemplateResponseProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                                    `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                                    `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                                    `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]ResponseSensorsEditSensorTestTemplateResponseProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type ResponseSensorsEditSensorTestTemplateResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsEditSensorTestTemplateResponseProfilesTests struct {
	Name   string                                                              `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsEditSensorTestTemplateResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsEditSensorTestTemplateResponseProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsEditSensorTestTemplateResponseProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type ResponseSensorsEditSensorTestTemplateResponseLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsEditSensorTestTemplateResponseSensors struct {
	Name                    string                                                                `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                                `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                                `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                                `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                                `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                                 `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                                `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                                `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                                `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                                `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                                 `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                                `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                                 `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                              `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                                `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                                `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                                 `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                                `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                                `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *ResponseSensorsEditSensorTestTemplateResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                                `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                                `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *ResponseSensorsEditSensorTestTemplateResponseSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type ResponseSensorsEditSensorTestTemplateResponseSensorsTestMacAddresses interface{}
type ResponseSensorsEditSensorTestTemplateResponseSensorsIPerfInfo interface{}
type ResponseSensorsEditSensorTestTemplateResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type ResponseSensorsCreateSensorTestTemplate struct {
	Version  string                                           `json:"version,omitempty"`  // Version
	Response *ResponseSensorsCreateSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponse struct {
	Name                   string                                                             `json:"name,omitempty"`                   // The sensor test template name
	TypeID                 string                                                             `json:"_id,omitempty"`                    // (Used in edit only) The sensor test template unique identifier
	Version                *int                                                               `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                               `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                               `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                               `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                               `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                             `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                             `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                             `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                             `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                             `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *ResponseSensorsCreateSensorTestTemplateResponseFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                               `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                               `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                               `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                           `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]ResponseSensorsCreateSensorTestTemplateResponseProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                             `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                              `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                              `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                             `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                             `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]ResponseSensorsCreateSensorTestTemplateResponseSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]ResponseSensorsCreateSensorTestTemplateResponseApCoverage       `json:"apCoverage,omitempty"`             //
}
type ResponseSensorsCreateSensorTestTemplateResponseFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDs struct {
	Bands                     string                                                                   `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                                   `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                                   `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                                     `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                                     `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                                   `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                                   `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                                   `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                                   `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                                     `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                                     `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                                   `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                                     `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                                     `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                                   `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                                   `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                                   `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                                   `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                                   `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                                   `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                                   `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                                   `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                                   `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                                   `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                                   `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                                    `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                                   `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                                   `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                                   `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                                   `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                                   `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                                   `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                                   `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                                    `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                                    `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                                   `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                                   `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                                   `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests             `json:"tests,omitempty"`                     //
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests struct {
	Name   string                                                             `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateResponseProfiles struct {
	AuthType            string                                                                      `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                                      `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                                      `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                                      `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                                      `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                                      `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                       `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                                      `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                                      `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                                      `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                                      `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                                      `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                                      `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                                      `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                       `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                       `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                                      `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                                      `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                                      `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                                      `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                                      `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                                      `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesTests struct {
	Name   string                                                                `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsCreateSensorTestTemplateResponseSensors struct {
	Name                    string                                                                  `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                                  `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                                  `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                                  `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                                  `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                                   `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                                  `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                                  `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                                  `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                                  `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                                   `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                                  `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                                   `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                                `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                                  `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                                  `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                                   `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                                  `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                                  `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *ResponseSensorsCreateSensorTestTemplateResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                                  `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                                  `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *ResponseSensorsCreateSensorTestTemplateResponseSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type ResponseSensorsCreateSensorTestTemplateResponseSensorsTestMacAddresses interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSensorsIPerfInfo interface{}
type ResponseSensorsCreateSensorTestTemplateResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type ResponseSensorsDeleteSensorTest struct {
	Version  string                                   `json:"version,omitempty"`  // Version
	Response *ResponseSensorsDeleteSensorTestResponse `json:"response,omitempty"` //
}
type ResponseSensorsDeleteSensorTestResponse struct {
	TemplateName string `json:"templateName,omitempty"` // Test template name to be delete
	Status       string `json:"status,omitempty"`       // Status of the DELETE action
}
type ResponseSensorsSensors struct {
	Version  string                            `json:"version,omitempty"`  // Version string of this API
	Response *[]ResponseSensorsSensorsResponse `json:"response,omitempty"` //
}
type ResponseSensorsSensorsResponse struct {
	Name               string                             `json:"name,omitempty"`               // The sensor device name
	Status             string                             `json:"status,omitempty"`             // Status of sensor device (REACHABLE, UNREACHABLE, DELETED, RUNNING, IDLE, UCLAIMED)
	RadioMacAddress    string                             `json:"radioMacAddress,omitempty"`    // Sensor device's radio MAC address
	EthernetMacAddress string                             `json:"ethernetMacAddress,omitempty"` // Sensor device's ethernet MAC address
	Location           string                             `json:"location,omitempty"`           // Site name in hierarchy form
	BackhaulType       string                             `json:"backhaulType,omitempty"`       // Backhall type: WIRED, WIRELESS
	SerialNumber       string                             `json:"serialNumber,omitempty"`       // Serial number
	IPAddress          string                             `json:"ipAddress,omitempty"`          // IP Address
	Version            string                             `json:"version,omitempty"`            // Sensor version
	LastSeen           *int                               `json:"lastSeen,omitempty"`           // Last seen timestamp
	Type               string                             `json:"type,omitempty"`               // Type
	SSH                *ResponseSensorsSensorsResponseSSH `json:"ssh,omitempty"`                //
	Led                *bool                              `json:"led,omitempty"`                // Is LED Enabled
}
type ResponseSensorsSensorsResponseSSH struct {
	SSHState       string `json:"sshState,omitempty"`       // SSH state
	SSHUserName    string `json:"sshUserName,omitempty"`    // SSH user name
	SSHPassword    string `json:"sshPassword,omitempty"`    // SSH password
	EnablePassword string `json:"enablePassword,omitempty"` // Enable password
}
type ResponseSensorsDuplicateSensorTestTemplate struct {
	Version  string                                              `json:"version,omitempty"`  // Version
	Response *ResponseSensorsDuplicateSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponse struct {
	Name                   string                                                                `json:"name,omitempty"`                   // The sensor test template name
	TypeID                 string                                                                `json:"_id,omitempty"`                    // The sensor test template unique identifier
	Version                *int                                                                  `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                                  `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                                  `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                                  `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                                  `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                                `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                                `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                                `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                                `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                                `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *ResponseSensorsDuplicateSensorTestTemplateResponseFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                                  `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                                  `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                                  `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                              `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                                `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                                 `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                                 `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                                `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                                `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]ResponseSensorsDuplicateSensorTestTemplateResponseSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage       `json:"apCoverage,omitempty"`             //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs struct {
	Bands                     string                                                                      `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                                      `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                                      `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                                        `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                                        `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                                      `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                                      `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                                      `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                                      `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                                        `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                                        `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                                      `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                                        `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                                        `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                                      `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                                      `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                                      `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                                      `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                                      `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                                      `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                                      `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                                      `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                                      `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                                      `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                                      `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                                       `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                                      `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                                      `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                                      `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                                      `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                                      `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                                      `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                                      `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                                       `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                                       `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                                      `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                                      `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                                      `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests             `json:"tests,omitempty"`                     //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests struct {
	Name   string                                                                `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfiles struct {
	AuthType            string                                                                         `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                                         `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                                         `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                                         `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                                         `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                                         `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                                          `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                                         `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                                         `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                                         `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                                         `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                                         `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                                         `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                                         `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                                          `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                                          `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                                         `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                                         `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                                         `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                                         `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                                         `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                                         `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTests struct {
	Name   string                                                                   `json:"name,omitempty"`   // Name of the test
	Config *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensors struct {
	Name                    string                                                                     `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                                     `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                                     `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                                     `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                                     `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                                      `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                                     `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                                     `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                                     `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                                     `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                                      `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                                     `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                                      `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                                   `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                                     `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                                     `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                                      `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                                     `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                                     `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *ResponseSensorsDuplicateSensorTestTemplateResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                                     `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                                     `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *ResponseSensorsDuplicateSensorTestTemplateResponseSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensorsTestMacAddresses interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensorsIPerfInfo interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type RequestSensorsEditSensorTestTemplate struct {
	TemplateName           string                                                  `json:"templateName,omitempty"`           // The test template name that is to be edited
	Name                   string                                                  `json:"name,omitempty"`                   // The sensor test template name, which is the same as in 'templateName'
	TypeID                 string                                                  `json:"_id,omitempty"`                    // The sensor test template unique identifier, generated at test creation time
	Version                *int                                                    `json:"version,omitempty"`                // The sensor test template version (must be 2)
	ModelVersion           *int                                                    `json:"modelVersion,omitempty"`           // Test template object model version (must be 2)
	StartTime              *int                                                    `json:"startTime,omitempty"`              // Start time
	LastModifiedTime       *int                                                    `json:"lastModifiedTime,omitempty"`       // Last modify time
	NumAssociatedSensor    *int                                                    `json:"numAssociatedSensor,omitempty"`    // Number of associated sensor
	Location               string                                                  `json:"location,omitempty"`               // Location string
	SiteHierarchy          string                                                  `json:"siteHierarchy,omitempty"`          // Site hierarchy
	Status                 string                                                  `json:"status,omitempty"`                 // Status of the test (RUNNING, NOTRUNNING)
	Connection             string                                                  `json:"connection,omitempty"`             // connection type of test: WIRED, WIRELESS, BOTH
	ActionInProgress       string                                                  `json:"actionInProgress,omitempty"`       // Indication of inprogress action
	Frequency              *RequestSensorsEditSensorTestTemplateFrequency          `json:"frequency,omitempty"`              //
	RssiThreshold          *int                                                    `json:"rssiThreshold,omitempty"`          // RSSI threshold
	NumNeighborApThreshold *int                                                    `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold
	ScheduleInDays         *int                                                    `json:"scheduleInDays,omitempty"`         // Bit-wise value of scheduled test days
	WLANs                  []string                                                `json:"wlans,omitempty"`                  // WLANs list
	SSIDs                  *[]RequestSensorsEditSensorTestTemplateSSIDs            `json:"ssids,omitempty"`                  //
	Profiles               *[]RequestSensorsEditSensorTestTemplateProfiles         `json:"profiles,omitempty"`               //
	TestScheduleMode       string                                                  `json:"testScheduleMode,omitempty"`       // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
	ShowWlcUpgradeBanner   *bool                                                   `json:"showWlcUpgradeBanner,omitempty"`   // Show WLC upgrade banner
	RadioAsSensorRemoved   *bool                                                   `json:"radioAsSensorRemoved,omitempty"`   // Radio as sensor removed
	EncryptionMode         string                                                  `json:"encryptionMode,omitempty"`         // Encryption mode
	RunNow                 string                                                  `json:"runNow,omitempty"`                 // Run now (YES, NO)
	LocationInfoList       *[]RequestSensorsEditSensorTestTemplateLocationInfoList `json:"locationInfoList,omitempty"`       //
	Sensors                *[]RequestSensorsEditSensorTestTemplateSensors          `json:"sensors,omitempty"`                //
	ApCoverage             *[]RequestSensorsEditSensorTestTemplateApCoverage       `json:"apCoverage,omitempty"`             //
}
type RequestSensorsEditSensorTestTemplateFrequency struct {
	Value *int   `json:"value,omitempty"` // Value of the unit
	Unit  string `json:"unit,omitempty"`  // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type RequestSensorsEditSensorTestTemplateSSIDs struct {
	Bands                     string                                                        `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                        `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                        `json:"profileName,omitempty"`               // The SSID profile name string
	NumAps                    *int                                                          `json:"numAps,omitempty"`                    // Number of APs in the test
	NumSensors                *int                                                          `json:"numSensors,omitempty"`                // Number of Sensors in the test
	Layer3WebAuthsecurity     string                                                        `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                        `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                        `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                        `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *RequestSensorsEditSensorTestTemplateSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	ID                        *int                                                          `json:"id,omitempty"`                        // Identification number
	WLANID                    *int                                                          `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                        `json:"wlc,omitempty"`                       // WLC IP addres
	ValidFrom                 *int                                                          `json:"validFrom,omitempty"`                 // Valid From UTC timestamp
	ValidTo                   *int                                                          `json:"validTo,omitempty"`                   // Valid To UTC timestamp
	Status                    string                                                        `json:"status,omitempty"`                    // WLAN status: ENABLED or DISABLED
	ProxyServer               string                                                        `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                        `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                        `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                        `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                        `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                        `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                        `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                        `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                        `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                        `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                         `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                        `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                        `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                        `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                        `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                        `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                        `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                        `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                         `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                         `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                        `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                        `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                        `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]RequestSensorsEditSensorTestTemplateSSIDsTests             `json:"tests,omitempty"`                     //
}
type RequestSensorsEditSensorTestTemplateSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateSSIDsTests struct {
	Name   string                                                  `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsEditSensorTestTemplateSSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsEditSensorTestTemplateProfiles struct {
	AuthType            string                                                           `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                           `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                           `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                           `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                           `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                           `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                            `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                           `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                           `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                           `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                           `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                           `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                           `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                           `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                            `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                            `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                           `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                           `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                           `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]RequestSensorsEditSensorTestTemplateProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                           `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                           `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                           `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]RequestSensorsEditSensorTestTemplateProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateProfilesTests struct {
	Name   string                                                     `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsEditSensorTestTemplateProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsEditSensorTestTemplateProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type RequestSensorsEditSensorTestTemplateLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsEditSensorTestTemplateSensors struct {
	Name                    string                                                       `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                       `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                       `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                       `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                       `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                        `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                       `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                       `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                       `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                       `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                        `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                       `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                        `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                     `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                       `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                       `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                        `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                       `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                       `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *RequestSensorsEditSensorTestTemplateSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                       `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                       `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *RequestSensorsEditSensorTestTemplateSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type RequestSensorsEditSensorTestTemplateSensorsTestMacAddresses interface{}
type RequestSensorsEditSensorTestTemplateSensorsIPerfInfo interface{}
type RequestSensorsEditSensorTestTemplateApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type RequestSensorsCreateSensorTestTemplate struct {
	Name             string                                                    `json:"name,omitempty"`             // The sensor test template name
	Version          *int                                                      `json:"version,omitempty"`          // The sensor test template version (must be 2)
	ModelVersion     *int                                                      `json:"modelVersion,omitempty"`     // Test template object model version (must be 2)
	Connection       string                                                    `json:"connection,omitempty"`       // connection type of test: WIRED, WIRELESS, BOTH
	SSIDs            *[]RequestSensorsCreateSensorTestTemplateSSIDs            `json:"ssids,omitempty"`            //
	Profiles         *[]RequestSensorsCreateSensorTestTemplateProfiles         `json:"profiles,omitempty"`         //
	EncryptionMode   string                                                    `json:"encryptionMode,omitempty"`   // Encryption mode
	RunNow           string                                                    `json:"runNow,omitempty"`           // Run now (YES, NO)
	LocationInfoList *[]RequestSensorsCreateSensorTestTemplateLocationInfoList `json:"locationInfoList,omitempty"` //
	Sensors          *[]RequestSensorsCreateSensorTestTemplateSensors          `json:"sensors,omitempty"`          //
	ApCoverage       *[]RequestSensorsCreateSensorTestTemplateApCoverage       `json:"apCoverage,omitempty"`       //
}
type RequestSensorsCreateSensorTestTemplateSSIDs struct {
	Bands                     string                                                          `json:"bands,omitempty"`                     // WIFI bands: 2.4GHz or 5GHz
	SSID                      string                                                          `json:"ssid,omitempty"`                      // The SSID string
	ProfileName               string                                                          `json:"profileName,omitempty"`               // The SSID profile name string
	Layer3WebAuthsecurity     string                                                          `json:"layer3webAuthsecurity,omitempty"`     // Layer 3 WEB Auth security
	Layer3WebAuthuserName     string                                                          `json:"layer3webAuthuserName,omitempty"`     // Layer 3 WEB Auth user name
	Layer3WebAuthpassword     string                                                          `json:"layer3webAuthpassword,omitempty"`     // Layer 3 WEB Auth password
	Layer3WebAuthEmailAddress string                                                          `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address
	ThirdParty                *RequestSensorsCreateSensorTestTemplateSSIDsThirdParty          `json:"thirdParty,omitempty"`                //
	WLANID                    *int                                                            `json:"wlanId,omitempty"`                    // WLAN ID
	Wlc                       string                                                          `json:"wlc,omitempty"`                       // WLC IP addres
	ProxyServer               string                                                          `json:"proxyServer,omitempty"`               // Proxy server for onboarding SSID
	ProxyPort                 string                                                          `json:"proxyPort,omitempty"`                 // Proxy server port
	ProxyUserName             string                                                          `json:"proxyUserName,omitempty"`             // Proxy server user name
	ProxyPassword             string                                                          `json:"proxyPassword,omitempty"`             // Proxy server password
	AuthType                  string                                                          `json:"authType,omitempty"`                  // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                       string                                                          `json:"psk,omitempty"`                       // Password of SSID when passwordType is ASCII
	Username                  string                                                          `json:"username,omitempty"`                  // User name string for onboarding SSID
	Password                  string                                                          `json:"password,omitempty"`                  // Password string for onboarding SSID
	PasswordType              string                                                          `json:"passwordType,omitempty"`              // SSID password type: ASCII or HEX
	EapMethod                 string                                                          `json:"eapMethod,omitempty"`                 // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                      *bool                                                           `json:"scep,omitempty"`                      // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol              string                                                          `json:"authProtocol,omitempty"`              // Auth protocol
	Certfilename              string                                                          `json:"certfilename,omitempty"`              // Auth certificate file name
	Certxferprotocol          string                                                          `json:"certxferprotocol,omitempty"`          // Certificate transfering protocol: HTTP or HTTPS
	Certstatus                string                                                          `json:"certstatus,omitempty"`                // Certificate status: INACTIVE or ACTIVE
	Certpassphrase            string                                                          `json:"certpassphrase,omitempty"`            // Certificate password phrase
	Certdownloadurl           string                                                          `json:"certdownloadurl,omitempty"`           // Certificate download URL
	ExtWebAuthVirtualIP       string                                                          `json:"extWebAuthVirtualIp,omitempty"`       // External WEB Auth virtual IP
	ExtWebAuth                *bool                                                           `json:"extWebAuth,omitempty"`                // Indication of using external WEB Auth
	WhiteList                 *bool                                                           `json:"whiteList,omitempty"`                 // Indication of being on allowed list
	ExtWebAuthPortal          string                                                          `json:"extWebAuthPortal,omitempty"`          // External authentication portal
	ExtWebAuthAccessURL       string                                                          `json:"extWebAuthAccessUrl,omitempty"`       // External WEB Auth access URL
	ExtWebAuthHTMLTag         *[]RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`         //
	QosPolicy                 string                                                          `json:"qosPolicy,omitempty"`                 // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests                     *[]RequestSensorsCreateSensorTestTemplateSSIDsTests             `json:"tests,omitempty"`                     //
}
type RequestSensorsCreateSensorTestTemplateSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateSSIDsTests struct {
	Name   string                                                    `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsCreateSensorTestTemplateProfiles struct {
	AuthType            string                                                             `json:"authType,omitempty"`            // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
	Psk                 string                                                             `json:"psk,omitempty"`                 // Password of SSID when passwordType is ASCII
	Username            string                                                             `json:"username,omitempty"`            // User name string for onboarding SSID
	Password            string                                                             `json:"password,omitempty"`            // Password string for onboarding SSID
	PasswordType        string                                                             `json:"passwordType,omitempty"`        // SSID password type: ASCII or HEX
	EapMethod           string                                                             `json:"eapMethod,omitempty"`           // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
	Scep                *bool                                                              `json:"scep,omitempty"`                // Secure certificate enrollment protocol: true or false or null for not applicable
	AuthProtocol        string                                                             `json:"authProtocol,omitempty"`        // Auth protocol
	Certfilename        string                                                             `json:"certfilename,omitempty"`        // Auth certificate file name
	Certxferprotocol    string                                                             `json:"certxferprotocol,omitempty"`    // Certificate transfering protocol: HTTP or HTTPS
	Certstatus          string                                                             `json:"certstatus,omitempty"`          // Certificate status: INACTIVE or ACTIVE
	Certpassphrase      string                                                             `json:"certpassphrase,omitempty"`      // Certificate password phrase
	Certdownloadurl     string                                                             `json:"certdownloadurl,omitempty"`     // Certificate download URL
	ExtWebAuthVirtualIP string                                                             `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP
	ExtWebAuth          *bool                                                              `json:"extWebAuth,omitempty"`          // Indication of using external WEB Auth
	WhiteList           *bool                                                              `json:"whiteList,omitempty"`           // Indication of being on allowed list
	ExtWebAuthPortal    string                                                             `json:"extWebAuthPortal,omitempty"`    // External authentication portal
	ExtWebAuthAccessURL string                                                             `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL
	ExtWebAuthHTMLTag   *[]RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"`   //
	QosPolicy           string                                                             `json:"qosPolicy,omitempty"`           // QoS policy: PlATINUM, GOLD, SILVER, BRONZE
	Tests               *[]RequestSensorsCreateSensorTestTemplateProfilesTests             `json:"tests,omitempty"`               //
	ProfileName         string                                                             `json:"profileName,omitempty"`         // Profile name
	DeviceType          string                                                             `json:"deviceType,omitempty"`          // Device Type
	VLAN                string                                                             `json:"vlan,omitempty"`                // VLAN
	LocationVLANList    *[]RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList  `json:"locationVlanList,omitempty"`    //
}
type RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label
	Tag   string `json:"tag,omitempty"`   // Tag
	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateProfilesTests struct {
	Name   string                                                       `json:"name,omitempty"`   // Name of the test
	Config *[]RequestSensorsCreateSensorTestTemplateProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateProfilesTestsConfig struct {
	Domains        []string `json:"domains,omitempty"`        // DNS domain name
	Server         string   `json:"server,omitempty"`         // Ping, file transfer, mail, radius, ssh, or telnet server
	UserName       string   `json:"userName,omitempty"`       // User name
	Password       string   `json:"password,omitempty"`       // Password
	URL            string   `json:"url,omitempty"`            // URL
	Port           *int     `json:"port,omitempty"`           // Radius or WEB server port
	Protocol       string   `json:"protocol,omitempty"`       // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
	Servers        []string `json:"servers,omitempty"`        // IPerf server list
	Direction      string   `json:"direction,omitempty"`      // IPerf direction (UPLOAD, DOWNLOAD, BOTH)
	StartPort      *int     `json:"startPort,omitempty"`      // IPerf start port
	EndPort        *int     `json:"endPort,omitempty"`        // IPerf end port
	UDPBandwidth   *int     `json:"udpBandwidth,omitempty"`   // IPerf UDP bandwidth
	ProbeType      string   `json:"probeType,omitempty"`      // Probe type
	NumPackets     *int     `json:"numPackets,omitempty"`     // Number of packets
	PathToDownload string   `json:"pathToDownload,omitempty"` // File path for file transfer
	TransferType   string   `json:"transferType,omitempty"`   // File transfer type (UPLOAD, DOWNLOAD, BOTH)
	SharedSecret   string   `json:"sharedSecret,omitempty"`   // Shared secret
	NdtServer      string   `json:"ndtServer,omitempty"`      // NDT server
	NdtServerPort  string   `json:"ndtServerPort,omitempty"`  // NDT server port
	NdtServerPath  string   `json:"ndtServerPath,omitempty"`  // NDT server path
	UplinkTest     *bool    `json:"uplinkTest,omitempty"`     // Uplink test
	DownlinkTest   *bool    `json:"downlinkTest,omitempty"`   // Downlink test
	ProxyServer    string   `json:"proxyServer,omitempty"`    // Proxy server
	ProxyPort      string   `json:"proxyPort,omitempty"`      // Proxy port
	ProxyUserName  string   `json:"proxyUserName,omitempty"`  // Proxy user name
	ProxyPassword  string   `json:"proxyPassword,omitempty"`  // Proxy password
	UserNamePrompt string   `json:"userNamePrompt,omitempty"` // User name prompt
	PasswordPrompt string   `json:"passwordPrompt,omitempty"` // Password prompt
	ExitCommand    string   `json:"exitCommand,omitempty"`    // Exit command
	FinalPrompt    string   `json:"finalPrompt,omitempty"`    // Final prompt
}
type RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList struct {
	LocationID string   `json:"locationId,omitempty"` // Site UUID
	VLANs      []string `json:"vlans,omitempty"`      // Array of VLANs
}
type RequestSensorsCreateSensorTestTemplateLocationInfoList struct {
	LocationID           string   `json:"locationId,omitempty"`           // Site UUID
	LocationType         string   `json:"locationType,omitempty"`         // Site type
	AllSensors           *bool    `json:"allSensors,omitempty"`           // Use all sensors in the site for test
	SiteHierarchy        string   `json:"siteHierarchy,omitempty"`        // Site name hierarhy
	MacAddressList       []string `json:"macAddressList,omitempty"`       // MAC addresses
	ManagementVLAN       string   `json:"managementVlan,omitempty"`       // Management VLAN
	CustomManagementVLAN *bool    `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsCreateSensorTestTemplateSensors struct {
	Name                    string                                                         `json:"name,omitempty"`                    // Sensor name
	MacAddress              string                                                         `json:"macAddress,omitempty"`              // MAC address
	SwitchMac               string                                                         `json:"switchMac,omitempty"`               // Switch MAC address
	SwitchUUID              string                                                         `json:"switchUuid,omitempty"`              // Switch device UUID
	SwitchSerialNumber      string                                                         `json:"switchSerialNumber,omitempty"`      // Switch serial number
	MarkedForUninstall      *bool                                                          `json:"markedForUninstall,omitempty"`      // Is marked for uninstall
	IPAddress               string                                                         `json:"ipAddress,omitempty"`               // IP address
	HostName                string                                                         `json:"hostName,omitempty"`                // Host name
	WiredApplicationStatus  string                                                         `json:"wiredApplicationStatus,omitempty"`  // Wired application status
	WiredApplicationMessage string                                                         `json:"wiredApplicationMessage,omitempty"` // Wired application message
	Assigned                *bool                                                          `json:"assigned,omitempty"`                // Is assigned
	Status                  string                                                         `json:"status,omitempty"`                  // Sensor device status: UP, DOWN, REBOOT
	XorSensor               *bool                                                          `json:"xorSensor,omitempty"`               // Is XOR sensor
	TargetAPs               []string                                                       `json:"targetAPs,omitempty"`               // Array of target APs
	RunNow                  string                                                         `json:"runNow,omitempty"`                  // Run now: YES, NO
	LocationID              string                                                         `json:"locationId,omitempty"`              // Site UUID
	AllSensorAddition       *bool                                                          `json:"allSensorAddition,omitempty"`       // Is all sensor addition
	ConfigUpdated           string                                                         `json:"configUpdated,omitempty"`           // Configuration updated: YES, NO
	SensorType              string                                                         `json:"sensorType,omitempty"`              // Sensor type
	TestMacAddresses        *RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses `json:"testMacAddresses,omitempty"`        // A string-string test MAC address
	ID                      string                                                         `json:"id,omitempty"`                      // Sensor ID
	ServicePolicy           string                                                         `json:"servicePolicy,omitempty"`           // Service policy
	IPerfInfo               *RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo        `json:"iPerfInfo,omitempty"`               // A string-stringList iPerf information
}
type RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses interface{}
type RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo interface{}
type RequestSensorsCreateSensorTestTemplateApCoverage struct {
	Bands             string `json:"bands,omitempty"`             // The WIFI bands
	NumberOfApsToTest *int   `json:"numberOfApsToTest,omitempty"` // Number of APs to test
	RssiThreshold     *int   `json:"rssiThreshold,omitempty"`     // RSSI threshold
}
type RequestSensorsRunNowSensorTest struct {
	TemplateName string `json:"templateName,omitempty"` // Template Name
}
type RequestSensorsDuplicateSensorTestTemplate struct {
	TemplateName    string `json:"templateName,omitempty"`    // Source test template name
	NewTemplateName string `json:"newTemplateName,omitempty"` // Destination test template name
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
