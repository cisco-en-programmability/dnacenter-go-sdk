package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SensorsService service

type ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams struct {
	Type      string  `url:"type,omitempty"`      //Capture Type
	ClientMac string  `url:"clientMac,omitempty"` //The macAddress of client
	ApMac     string  `url:"apMac,omitempty"`     //The base radio macAddress of the access point
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	Limit     float64 `url:"limit,omitempty"`     //Maximum number of records to return
	Offset    float64 `url:"offset,omitempty"`    //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy    string  `url:"sortBy,omitempty"`    //A field within the response to sort by.
	Order     string  `url:"order,omitempty"`     //The sort order of the field ascending or descending.
}
type ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams struct {
	Type      string  `url:"type,omitempty"`      //Capture Type
	ClientMac string  `url:"clientMac,omitempty"` //The macAddress of client
	ApMac     string  `url:"apMac,omitempty"`     //The base radio macAddress of the access point
	StartTime float64 `url:"startTime,omitempty"` //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime   float64 `url:"endTime,omitempty"`   //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
}
type RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesDetailsOfASpecificICapPacketCaptureFileHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type DownloadsASpecificICapPacketCaptureFileHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacQueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ApMac         string  `url:"apMac,omitempty"`         //The base ethernet macAddress of the access point
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	TimeSortOrder string  `url:"timeSortOrder,omitempty"` //The sort order of the field ascending or descending.
}
type RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacQueryParams struct {
	StartTime     float64 `url:"startTime,omitempty"`     //Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	EndTime       float64 `url:"endTime,omitempty"`       //End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
	ApMac         string  `url:"apMac,omitempty"`         //The base ethernet macAddress of the access point
	DataType      float64 `url:"dataType,omitempty"`      //Data type reported by the sensor |Data Type | Description | | --- | --- | | `0` | Duty Cycle | | `1` | Max Power | | `2` | Average Power | | `3` | Max Power in dBm with adjusted base of +48 | | `4` | Average Power in dBm with adjusted base of +48 |
	Limit         float64 `url:"limit,omitempty"`         //Maximum number of records to return
	Offset        float64 `url:"offset,omitempty"`        //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	TimeSortOrder string  `url:"timeSortOrder,omitempty"` //The sort order of the field ascending or descending.
}
type RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringQueryParams struct {
	CaptureStatus string  `url:"captureStatus,omitempty"` //Catalyst Center ICAP status
	CaptureType   string  `url:"captureType,omitempty"`   //Catalyst Center ICAP type
	ClientMac     string  `url:"clientMac,omitempty"`     //The client device MAC address in ICAP configuration
	APID          string  `url:"apId,omitempty"`          //The AP device's UUID.
	WlcID         string  `url:"wlcId,omitempty"`         //The wireless controller device's UUID
	Offset        float64 `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page.
}
type CreatesAnICapConfigurationIntentForPreviewApproveQueryParams struct {
	PreviewDescription string `url:"previewDescription,omitempty"` //The ICAP intent's preview-deploy description string
}
type RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringQueryParams struct {
	CaptureType   string `url:"captureType,omitempty"`   //Catalyst Center ICAP type
	CaptureStatus string `url:"captureStatus,omitempty"` //Catalyst Center ICAP status
	ClientMac     string `url:"clientMac,omitempty"`     //The client device MAC address in ICAP configuration
	APID          string `url:"apId,omitempty"`          //The AP device's UUID.
	WlcID         string `url:"wlcId,omitempty"`         //The wireless controller device's UUID
}
type DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveQueryParams struct {
	PreviewDescription string `url:"previewDescription,omitempty"` //The ICAP intent's preview-deploy description string
}
type GetDeviceDeploymentStatusQueryParams struct {
	DeployActivityID string  `url:"deployActivityId,omitempty"` //activity from the /deploy task response
	NetworkDeviceIDs string  `url:"networkDeviceIds,omitempty"` //device ids, retrievable from the id attribute in intent/api/v1/network-device
	Offset           float64 `url:"offset,omitempty"`           //The first record to show for this page; the first record is numbered 1.
	Limit            float64 `url:"limit,omitempty"`            //The number of records to show for this page.
	SortBy           string  `url:"sortBy,omitempty"`           //A property within the response to sort by.
	Order            string  `url:"order,omitempty"`            //Whether ascending or descending order should be used to sort the response.
}
type GetDeviceDeploymentStatusCountQueryParams struct {
	DeployActivityID string `url:"deployActivityId,omitempty"` //activity from the /deploy task response
	NetworkDeviceIDs string `url:"networkDeviceIds,omitempty"` //device ids, retrievable from the id attribute in intent/api/v1/network-device
}
type DeleteSensorTestQueryParams struct {
	TemplateName string `url:"templateName,omitempty"` //
}
type SensorsQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //
}

type ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteria struct {
	Response *[]ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaResponse `json:"response,omitempty"` //

	Page *ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaResponse struct {
	ID string `json:"id,omitempty"` // Id

	FileName string `json:"fileName,omitempty"` // File Name

	FileSize *int `json:"fileSize,omitempty"` // File Size

	Type string `json:"type,omitempty"` // Type

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	FileCreationTimestamp *int `json:"fileCreationTimestamp,omitempty"` // File Creation Timestamp

	LastUpdatedTimestamp *int `json:"lastUpdatedTimestamp,omitempty"` // Last Updated Timestamp
}
type ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteriaPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	SortBy string `json:"sortBy,omitempty"` // Sort By

	Order string `json:"order,omitempty"` // Order
}
type ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria struct {
	Response *ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFile struct {
	Response *ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFileResponse struct {
	ID string `json:"id,omitempty"` // Id

	FileName string `json:"fileName,omitempty"` // File Name

	FileSize *int `json:"fileSize,omitempty"` // File Size

	Type string `json:"type,omitempty"` // Type

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	FileCreationTimestamp *int `json:"fileCreationTimestamp,omitempty"` // File Creation Timestamp

	LastUpdatedTimestamp *int `json:"lastUpdatedTimestamp,omitempty"` // Last Updated Timestamp
}
type ResponseSensorsDownloadsASpecificICapPacketCaptureFile struct {
	object string `json:"object,omitempty"` // object
}
type ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime struct {
	Response *[]ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeResponse `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeResponse struct {
	ID string `json:"id,omitempty"` // Id

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	RadioID *int `json:"radioId,omitempty"` // Radio Id

	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Band string `json:"band,omitempty"` // Band

	SSID string `json:"ssid,omitempty"` // Ssid

	Rssi *int `json:"rssi,omitempty"` // Rssi

	Snr *int `json:"snr,omitempty"` // Snr

	TxBytes *int `json:"txBytes,omitempty"` // Tx Bytes

	RxBytes *int `json:"rxBytes,omitempty"` // Rx Bytes

	RxPackets *int `json:"rxPackets,omitempty"` // Rx Packets

	TxPackets *int `json:"txPackets,omitempty"` // Tx Packets

	RxMgmtPackets *int `json:"rxMgmtPackets,omitempty"` // Rx Mgmt Packets

	TxMgmtPackets *int `json:"txMgmtPackets,omitempty"` // Tx Mgmt Packets

	RxDataPackets *int `json:"rxDataPackets,omitempty"` // Rx Data Packets

	TxDataPackets *int `json:"txDataPackets,omitempty"` // Tx Data Packets

	TxUnicastDataPackets *float64 `json:"txUnicastDataPackets,omitempty"` // Tx Unicast Data Packets

	RxCtrlPackets *int `json:"rxCtrlPackets,omitempty"` // Rx Ctrl Packets

	TxCtrlPackets *int `json:"txCtrlPackets,omitempty"` // Tx Ctrl Packets

	RxRetries *int `json:"rxRetries,omitempty"` // Rx Retries

	RxRate *int `json:"rxRate,omitempty"` // Rx Rate

	TxRate *int `json:"txRate,omitempty"` // Tx Rate

	ClientIP string `json:"clientIp,omitempty"` // Client Ip
}
type ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime struct {
	Response *[]ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeResponse `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeResponse struct {
	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	RadioID *int `json:"radioId,omitempty"` // Radio Id

	Band string `json:"band,omitempty"` // Band

	Utilization *int `json:"utilization,omitempty"` // Utilization

	NonWifiUtilization *float64 `json:"nonWifiUtilization,omitempty"` // Non Wifi Utilization

	RxOtherBSSUtilization *int `json:"rxOtherBSSUtilization,omitempty"` // Rx Other B S S Utilization

	RxInBSSUtilization *float64 `json:"rxInBSSUtilization,omitempty"` // Rx In B S S Utilization

	TxUtilization *int `json:"txUtilization,omitempty"` // Tx Utilization

	NoiseFloor *int `json:"noiseFloor,omitempty"` // Noise Floor

	Channel *int `json:"channel,omitempty"` // Channel

	ChannelWidth *int `json:"channelWidth,omitempty"` // Channel Width

	TxPower *int `json:"txPower,omitempty"` // Tx Power

	MaxTxPower *float64 `json:"maxTxPower,omitempty"` // Max Tx Power

	TxBytes *int `json:"txBytes,omitempty"` // Tx Bytes

	RxBytes *int `json:"rxBytes,omitempty"` // Rx Bytes

	RxPackets *int `json:"rxPackets,omitempty"` // Rx Packets

	TxPackets *int `json:"txPackets,omitempty"` // Tx Packets

	RxMgmtPackets *int `json:"rxMgmtPackets,omitempty"` // Rx Mgmt Packets

	TxMgmtPackets *int `json:"txMgmtPackets,omitempty"` // Tx Mgmt Packets

	RxErrors *int `json:"rxErrors,omitempty"` // Rx Errors

	TxErrors *int `json:"txErrors,omitempty"` // Tx Errors
}
type ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac struct {
	Response *[]ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacResponse `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacResponse struct {
	ID string `json:"id,omitempty"` // Id

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	CentralFrequencyKHz *int `json:"centralFrequencyKHz,omitempty"` // Central Frequency K Hz

	BandWidthKHz *int `json:"bandWidthKHz,omitempty"` // Band Width K Hz

	LowEndFrequencyKHz *int `json:"lowEndFrequencyKHz,omitempty"` // Low End Frequency K Hz

	HighEndFrequencyKHz *int `json:"highEndFrequencyKHz,omitempty"` // High End Frequency K Hz

	PowerDbm *float64 `json:"powerDbm,omitempty"` // Power Dbm

	Band string `json:"band,omitempty"` // Band

	DutyCycle *int `json:"dutyCycle,omitempty"` // Duty Cycle

	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	SeverityIndex *int `json:"severityIndex,omitempty"` // Severity Index

	DetectedChannels *[]int `json:"detectedChannels,omitempty"` // Detected Channels
}
type ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac struct {
	Response *[]ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacResponse `json:"response,omitempty"` //

	Page *ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacPage `json:"page,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacResponse struct {
	ID string `json:"id,omitempty"` // Id

	SpanKHz *int `json:"spanKHz,omitempty"` // Span K Hz

	DataType *int `json:"dataType,omitempty"` // Data Type

	ApMac string `json:"apMac,omitempty"` // Ap Mac

	DataAvg *float64 `json:"dataAvg,omitempty"` // Data Avg

	DataMin *float64 `json:"dataMin,omitempty"` // Data Min

	DataMax *float64 `json:"dataMax,omitempty"` // Data Max

	DataUnits string `json:"dataUnits,omitempty"` // Data Units

	CentralFrequencyKHz *int `json:"centralFrequencyKHz,omitempty"` // Central Frequency K Hz

	Band string `json:"band,omitempty"` // Band

	Timestamp *int `json:"timestamp,omitempty"` // Timestamp

	Data *[]float64 `json:"data,omitempty"` // Data

	DataSize *int `json:"dataSize,omitempty"` // Data Size

	Channels *[]int `json:"channels,omitempty"` // Channels
}
type ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacPage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	Count *int `json:"count,omitempty"` // Count

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
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
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering struct {
	Response *[]ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringResponse struct {
	Name string `json:"name,omitempty"` // Name

	Slots *[]float64 `json:"slots,omitempty"` // Slots

	OtaBand string `json:"otaBand,omitempty"` // Ota Band

	OtaChannel *int `json:"otaChannel,omitempty"` // Ota Channel

	OtaChannelWidth *int `json:"otaChannelWidth,omitempty"` // Ota Channel Width

	ID string `json:"id,omitempty"` // Id

	DeployedID string `json:"deployedId,omitempty"` // Deployed Id

	DisableActivityID string `json:"disableActivityId,omitempty"` // Disable Activity Id

	ActivityID string `json:"activityId,omitempty"` // Activity Id

	CaptureType string `json:"captureType,omitempty"` // Capture Type

	APID string `json:"apId,omitempty"` // Ap Id

	WlcID string `json:"wlcId,omitempty"` // Wlc Id

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	CreateTime *int `json:"createTime,omitempty"` // Create Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	DurationInMins *int `json:"durationInMins,omitempty"` // Duration In Mins

	Status string `json:"status,omitempty"` // Status
}
type ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApprove struct {
	Response *ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice struct {
	Response *ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsDiscardsTheICapConfigurationIntentByActivityID struct {
	Response *ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsDeploysTheICapConfigurationIntentByActivityID struct {
	Response *ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDeploysTheICapConfigurationIntentByActivityIDResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsGetICapConfigurationStatusPerNetworkDevice struct {
	Response *[]ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	Status string `json:"status,omitempty"` // Status
}
type ResponseSensorsRetrievesTheDevicesClisOfTheICapintent struct {
	Response *ResponseSensorsRetrievesTheDevicesClisOfTheICapintentResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRetrievesTheDevicesClisOfTheICapintentResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	PreviewItems *[]ResponseSensorsRetrievesTheDevicesClisOfTheICapintentResponsePreviewItems `json:"previewItems,omitempty"` //

	Status string `json:"status,omitempty"` // Status
}
type ResponseSensorsRetrievesTheDevicesClisOfTheICapintentResponsePreviewItems struct {
	ConfigPreview string `json:"configPreview,omitempty"` // Config Preview

	ConfigType string `json:"configType,omitempty"` // Config Type

	ErrorMessages []string `json:"errorMessages,omitempty"` // Error Messages

	Name string `json:"name,omitempty"` // Name
}
type ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent struct {
	Response *ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntentResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering struct {
	Response *float64 `json:"response,omitempty"` // Response

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove struct {
	Response *ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreview struct {
	Response *ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreviewResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id

	URL string `json:"url,omitempty"` // Url
}
type ResponseSensorsGetDeviceDeploymentStatus struct {
	Response *[]ResponseSensorsGetDeviceDeploymentStatusResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGetDeviceDeploymentStatusResponse struct {
	DeployActivityID string `json:"deployActivityId,omitempty"` // Deploy Activity Id

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network Device Id

	ConfigGroupName string `json:"configGroupName,omitempty"` // Config Group Name

	ConfigGroupVersion *int `json:"configGroupVersion,omitempty"` // Config Group Version

	Status string `json:"status,omitempty"` // Status

	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Error *ResponseSensorsGetDeviceDeploymentStatusResponseError `json:"error,omitempty"` // Error
}
type ResponseSensorsGetDeviceDeploymentStatusResponseError interface{}
type ResponseSensorsGetDeviceDeploymentStatusCount struct {
	Response *ResponseSensorsGetDeviceDeploymentStatusCountResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSensorsGetDeviceDeploymentStatusCountResponse struct {
	Count *float64 `json:"count,omitempty"` // Count
}
type ResponseSensorsCreateSensorTestTemplate struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseSensorsCreateSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponse struct {
	Name   string `json:"name,omitempty"` // The sensor test template name
	TypeID string `json:"_id,omitempty"`  // (Used in edit only) The sensor test template unique identifier

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	StartTime *int `json:"startTime,omitempty"` // Start time

	LastModifiedTime *int `json:"lastModifiedTime,omitempty"` // Last modify time

	NumAssociatedSensor *int `json:"numAssociatedSensor,omitempty"` // Number of associated sensor

	Location string `json:"location,omitempty"` // Location string

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy

	Status string `json:"status,omitempty"` // Status of the test (RUNNING, NOTRUNNING)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	ActionInProgress string `json:"actionInProgress,omitempty"` // Indication of inprogress action

	Frequency *ResponseSensorsCreateSensorTestTemplateResponseFrequency `json:"frequency,omitempty"` //

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold

	NumNeighborApThreshold *int `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold

	ScheduleInDays *int `json:"scheduleInDays,omitempty"` // Bit-wise value of scheduled test days

	WLANs []string `json:"wlans,omitempty"` // WLANs list

	SSIDs *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDs `json:"ssids,omitempty"` //

	Profiles *[]ResponseSensorsCreateSensorTestTemplateResponseProfiles `json:"profiles,omitempty"` //

	TestScheduleMode string `json:"testScheduleMode,omitempty"` // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)

	ShowWlcUpgradeBanner *bool `json:"showWlcUpgradeBanner,omitempty"` // Show WLC upgrade banner

	RadioAsSensorRemoved *bool `json:"radioAsSensorRemoved,omitempty"` // Radio as sensor removed

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]ResponseSensorsCreateSensorTestTemplateResponseSensors `json:"sensors,omitempty"` //

	ApCoverage *[]ResponseSensorsCreateSensorTestTemplateResponseApCoverage `json:"apCoverage,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseFrequency struct {
	Value *int `json:"value,omitempty"` // Value of the unit

	Unit string `json:"unit,omitempty"` // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	NumAps *int `json:"numAps,omitempty"` // Number of APs in the test

	NumSensors *int `json:"numSensors,omitempty"` // Number of Sensors in the test

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty `json:"thirdParty,omitempty"` //

	ID *int `json:"id,omitempty"` // Identification number

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ValidFrom *int `json:"validFrom,omitempty"` // Valid From UTC timestamp

	ValidTo *int `json:"validTo,omitempty"` // Valid To UTC timestamp

	Status string `json:"status,omitempty"` // WLAN status: ENABLED or DISABLED

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests `json:"tests,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateResponseProfiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsCreateSensorTestTemplateResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsCreateSensorTestTemplateResponseProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsCreateSensorTestTemplateResponseSensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *ResponseSensorsCreateSensorTestTemplateResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *ResponseSensorsCreateSensorTestTemplateResponseSensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type ResponseSensorsCreateSensorTestTemplateResponseSensorsTestMacAddresses interface{}
type ResponseSensorsCreateSensorTestTemplateResponseSensorsIPerfInfo interface{}
type ResponseSensorsCreateSensorTestTemplateResponseApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type ResponseSensorsDeleteSensorTest struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseSensorsDeleteSensorTestResponse `json:"response,omitempty"` //
}
type ResponseSensorsDeleteSensorTestResponse struct {
	TemplateName string `json:"templateName,omitempty"` // Test template name to be delete

	Status string `json:"status,omitempty"` // Status of the DELETE action
}
type ResponseSensorsSensors struct {
	Version string `json:"version,omitempty"` // Version string of this API

	Response *[]ResponseSensorsSensorsResponse `json:"response,omitempty"` //
}
type ResponseSensorsSensorsResponse struct {
	Name string `json:"name,omitempty"` // The sensor device name

	Status string `json:"status,omitempty"` // Status of sensor device (REACHABLE, UNREACHABLE, DELETED, RUNNING, IDLE, UCLAIMED)

	RadioMacAddress string `json:"radioMacAddress,omitempty"` // Sensor device's radio MAC address

	EthernetMacAddress string `json:"ethernetMacAddress,omitempty"` // Sensor device's ethernet MAC address

	Location string `json:"location,omitempty"` // Site name in hierarchy form

	BackhaulType string `json:"backhaulType,omitempty"` // Backhall type: WIRED, WIRELESS

	SerialNumber string `json:"serialNumber,omitempty"` // Serial number

	IPAddress string `json:"ipAddress,omitempty"` // IP Address

	Version string `json:"version,omitempty"` // Sensor version

	LastSeen *int `json:"lastSeen,omitempty"` // Last seen timestamp

	Type string `json:"type,omitempty"` // Type

	SSH *ResponseSensorsSensorsResponseSSH `json:"ssh,omitempty"` //

	Led *bool `json:"led,omitempty"` // Is LED Enabled
}
type ResponseSensorsSensorsResponseSSH struct {
	SSHState string `json:"sshState,omitempty"` // SSH state

	SSHUserName string `json:"sshUserName,omitempty"` // SSH user name

	SSHPassword string `json:"sshPassword,omitempty"` // SSH password

	EnablePassword string `json:"enablePassword,omitempty"` // Enable password
}
type ResponseSensorsDuplicateSensorTestTemplate struct {
	Version string `json:"version,omitempty"` // Version

	Response *ResponseSensorsDuplicateSensorTestTemplateResponse `json:"response,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponse struct {
	Name   string `json:"name,omitempty"` // The sensor test template name
	TypeID string `json:"_id,omitempty"`  // The sensor test template unique identifier

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	StartTime *int `json:"startTime,omitempty"` // Start time

	LastModifiedTime *int `json:"lastModifiedTime,omitempty"` // Last modify time

	NumAssociatedSensor *int `json:"numAssociatedSensor,omitempty"` // Number of associated sensor

	Location string `json:"location,omitempty"` // Location string

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy

	Status string `json:"status,omitempty"` // Status of the test (RUNNING, NOTRUNNING)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	ActionInProgress string `json:"actionInProgress,omitempty"` // Indication of inprogress action

	Frequency *ResponseSensorsDuplicateSensorTestTemplateResponseFrequency `json:"frequency,omitempty"` //

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold

	NumNeighborApThreshold *int `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold

	ScheduleInDays *int `json:"scheduleInDays,omitempty"` // Bit-wise value of scheduled test days

	WLANs []string `json:"wlans,omitempty"` // WLANs list

	SSIDs *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs `json:"ssids,omitempty"` //

	Profiles *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfiles `json:"profiles,omitempty"` //

	TestScheduleMode string `json:"testScheduleMode,omitempty"` // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)

	ShowWlcUpgradeBanner *bool `json:"showWlcUpgradeBanner,omitempty"` // Show WLC upgrade banner

	RadioAsSensorRemoved *bool `json:"radioAsSensorRemoved,omitempty"` // Radio as sensor removed

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]ResponseSensorsDuplicateSensorTestTemplateResponseSensors `json:"sensors,omitempty"` //

	ApCoverage *[]ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage `json:"apCoverage,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseFrequency struct {
	Value *int `json:"value,omitempty"` // Value of the unit

	Unit string `json:"unit,omitempty"` // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	NumAps *int `json:"numAps,omitempty"` // Number of APs in the test

	NumSensors *int `json:"numSensors,omitempty"` // Number of Sensors in the test

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty `json:"thirdParty,omitempty"` //

	ID *int `json:"id,omitempty"` // Identification number

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ValidFrom *int `json:"validFrom,omitempty"` // Valid From UTC timestamp

	ValidTo *int `json:"validTo,omitempty"` // Valid To UTC timestamp

	Status string `json:"status,omitempty"` // WLAN status: ENABLED or DISABLED

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests `json:"tests,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTestsConfig `json:"config,omitempty"` //
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type ResponseSensorsDuplicateSensorTestTemplateResponseProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type ResponseSensorsDuplicateSensorTestTemplateResponseLocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *ResponseSensorsDuplicateSensorTestTemplateResponseSensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *ResponseSensorsDuplicateSensorTestTemplateResponseSensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensorsTestMacAddresses interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseSensorsIPerfInfo interface{}
type ResponseSensorsDuplicateSensorTestTemplateResponseApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters `json:"filters,omitempty"` //

	Page *RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage `json:"page,omitempty"` //
}
type RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value *int `json:"value,omitempty"` // Value
}
type RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime struct {
	StartTime *int `json:"startTime,omitempty"` // Start Time

	EndTime *int `json:"endTime,omitempty"` // End Time

	Filters *[]RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters `json:"filters,omitempty"` //

	Page *RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage `json:"page,omitempty"` //
}
type RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeFilters struct {
	Key string `json:"key,omitempty"` // Key

	Operator string `json:"operator,omitempty"` // Operator

	Value *int `json:"value,omitempty"` // Value
}
type RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimePage struct {
	Limit *int `json:"limit,omitempty"` // Limit

	Offset *int `json:"offset,omitempty"` // Offset

	TimeSortOrder string `json:"timeSortOrder,omitempty"` // Time Sort Order
}
type RequestSensorsEditSensorTestTemplate struct {
	TemplateName string `json:"templateName,omitempty"` // The test template name that is to be edited

	Name   string `json:"name,omitempty"` // The sensor test template name, which is the same as in 'templateName'
	TypeID string `json:"_id,omitempty"`  // The sensor test template unique identifier, generated at test creation time

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	StartTime *int `json:"startTime,omitempty"` // Start time

	LastModifiedTime *int `json:"lastModifiedTime,omitempty"` // Last modify time

	NumAssociatedSensor *int `json:"numAssociatedSensor,omitempty"` // Number of associated sensor

	Location string `json:"location,omitempty"` // Location string

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site hierarchy

	Status string `json:"status,omitempty"` // Status of the test (RUNNING, NOTRUNNING)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	ActionInProgress string `json:"actionInProgress,omitempty"` // Indication of inprogress action

	Frequency *RequestSensorsEditSensorTestTemplateFrequency `json:"frequency,omitempty"` //

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold

	NumNeighborApThreshold *int `json:"numNeighborAPThreshold,omitempty"` // Number of neighboring AP threshold

	ScheduleInDays *int `json:"scheduleInDays,omitempty"` // Bit-wise value of scheduled test days

	WLANs []string `json:"wlans,omitempty"` // WLANs list

	SSIDs *[]RequestSensorsEditSensorTestTemplateSSIDs `json:"ssids,omitempty"` //

	Profiles *[]RequestSensorsEditSensorTestTemplateProfiles `json:"profiles,omitempty"` //

	TestScheduleMode string `json:"testScheduleMode,omitempty"` // Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)

	ShowWlcUpgradeBanner *bool `json:"showWlcUpgradeBanner,omitempty"` // Show WLC upgrade banner

	RadioAsSensorRemoved *bool `json:"radioAsSensorRemoved,omitempty"` // Radio as sensor removed

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]RequestSensorsEditSensorTestTemplateLocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]RequestSensorsEditSensorTestTemplateSensors `json:"sensors,omitempty"` //

	ApCoverage *[]RequestSensorsEditSensorTestTemplateApCoverage `json:"apCoverage,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateFrequency struct {
	Value *int `json:"value,omitempty"` // Value of the unit

	Unit string `json:"unit,omitempty"` // Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
}
type RequestSensorsEditSensorTestTemplateSSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	NumAps *int `json:"numAps,omitempty"` // Number of APs in the test

	NumSensors *int `json:"numSensors,omitempty"` // Number of Sensors in the test

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *RequestSensorsEditSensorTestTemplateSSIDsThirdParty `json:"thirdParty,omitempty"` //

	ID *int `json:"id,omitempty"` // Identification number

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ValidFrom *int `json:"validFrom,omitempty"` // Valid From UTC timestamp

	ValidTo *int `json:"validTo,omitempty"` // Valid To UTC timestamp

	Status string `json:"status,omitempty"` // WLAN status: ENABLED or DISABLED

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsEditSensorTestTemplateSSIDsTests `json:"tests,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsEditSensorTestTemplateSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateSSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsEditSensorTestTemplateSSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateSSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsEditSensorTestTemplateProfiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EAP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsEditSensorTestTemplateProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]RequestSensorsEditSensorTestTemplateProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsEditSensorTestTemplateProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsEditSensorTestTemplateProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsEditSensorTestTemplateProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsEditSensorTestTemplateProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type RequestSensorsEditSensorTestTemplateLocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsEditSensorTestTemplateSensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *RequestSensorsEditSensorTestTemplateSensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *RequestSensorsEditSensorTestTemplateSensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type RequestSensorsEditSensorTestTemplateSensorsTestMacAddresses interface{}
type RequestSensorsEditSensorTestTemplateSensorsIPerfInfo interface{}
type RequestSensorsEditSensorTestTemplateApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type RequestSensorsCreatesAnICapConfigurationIntentForPreviewApprove []RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApprove // Array of RequestSensorsCreatesAnICAPConfigurationIntentForPreviewApprove
type RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApprove struct {
	CaptureType string `json:"captureType,omitempty"` // Capture Type

	DurationInMins *int `json:"durationInMins,omitempty"` // Duration In Mins

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	WlcID string `json:"wlcId,omitempty"` // Wlc Id

	APID string `json:"apId,omitempty"` // Ap Id

	Slot *[]float64 `json:"slot,omitempty"` // Slot

	OtaBand string `json:"otaBand,omitempty"` // Ota Band

	OtaChannel *int `json:"otaChannel,omitempty"` // Ota Channel

	OtaChannelWidth *int `json:"otaChannelWidth,omitempty"` // Ota Channel Width
}
type RequestSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsDeploysTheICapConfigurationIntentByActivityID struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove []RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove // Array of RequestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApprove
type RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove struct {
	CaptureType string `json:"captureType,omitempty"` // Capture Type

	DurationInMins *int `json:"durationInMins,omitempty"` // Duration In Mins

	ClientMac string `json:"clientMac,omitempty"` // Client Mac

	WlcID string `json:"wlcId,omitempty"` // Wlc Id

	APID string `json:"apId,omitempty"` // Ap Id

	Slot *[]float64 `json:"slot,omitempty"` // Slot

	OtaBand string `json:"otaBand,omitempty"` // Ota Band

	OtaChannel *int `json:"otaChannel,omitempty"` // Ota Channel

	OtaChannelWidth *int `json:"otaChannelWidth,omitempty"` // Ota Channel Width
}
type RequestSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreview struct {
	object string `json:"object,omitempty"` // object
}
type RequestSensorsCreateSensorTestTemplate struct {
	Name string `json:"name,omitempty"` // The sensor test template name

	Version *int `json:"version,omitempty"` // The sensor test template version (must be 2)

	ModelVersion *int `json:"modelVersion,omitempty"` // Test template object model version (must be 2)

	Connection string `json:"connection,omitempty"` // connection type of test: WIRED, WIRELESS, BOTH

	SSIDs *[]RequestSensorsCreateSensorTestTemplateSSIDs `json:"ssids,omitempty"` //

	Profiles *[]RequestSensorsCreateSensorTestTemplateProfiles `json:"profiles,omitempty"` //

	EncryptionMode string `json:"encryptionMode,omitempty"` // Encryption mode

	RunNow string `json:"runNow,omitempty"` // Run now (YES, NO)

	LocationInfoList *[]RequestSensorsCreateSensorTestTemplateLocationInfoList `json:"locationInfoList,omitempty"` //

	Sensors *[]RequestSensorsCreateSensorTestTemplateSensors `json:"sensors,omitempty"` //

	ApCoverage *[]RequestSensorsCreateSensorTestTemplateApCoverage `json:"apCoverage,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateSSIDs struct {
	Bands string `json:"bands,omitempty"` // WIFI bands: 2.4GHz or 5GHz

	SSID string `json:"ssid,omitempty"` // The SSID string

	ProfileName string `json:"profileName,omitempty"` // The SSID profile name string

	Layer3WebAuthsecurity string `json:"layer3webAuthsecurity,omitempty"` // Layer 3 WEB Auth security

	Layer3WebAuthuserName string `json:"layer3webAuthuserName,omitempty"` // Layer 3 WEB Auth user name

	Layer3WebAuthpassword string `json:"layer3webAuthpassword,omitempty"` // Layer 3 WEB Auth password

	Layer3WebAuthEmailAddress string `json:"layer3webAuthEmailAddress,omitempty"` // Layer 3 WEB Auth email address

	ThirdParty *RequestSensorsCreateSensorTestTemplateSSIDsThirdParty `json:"thirdParty,omitempty"` //

	WLANID *int `json:"wlanId,omitempty"` // WLAN ID

	Wlc string `json:"wlc,omitempty"` // WLC IP addres

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server for onboarding SSID

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy server port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy server user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy server password

	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsCreateSensorTestTemplateSSIDsTests `json:"tests,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateSSIDsThirdParty struct {
	Selected *bool `json:"selected,omitempty"` // true: the SSID is third party
}
type RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateSSIDsTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsCreateSensorTestTemplateProfiles struct {
	AuthType string `json:"authType,omitempty"` // Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER

	Psk string `json:"psk,omitempty"` // Password of SSID when passwordType is ASCII

	Username string `json:"username,omitempty"` // User name string for onboarding SSID

	Password string `json:"password,omitempty"` // Password string for onboarding SSID

	PasswordType string `json:"passwordType,omitempty"` // SSID password type: ASCII or HEX

	EapMethod string `json:"eapMethod,omitempty"` // WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC

	Scep *bool `json:"scep,omitempty"` // Secure certificate enrollment protocol: true or false or null for not applicable

	AuthProtocol string `json:"authProtocol,omitempty"` // Auth protocol

	Certfilename string `json:"certfilename,omitempty"` // Auth certificate file name

	Certxferprotocol string `json:"certxferprotocol,omitempty"` // Certificate transfering protocol: HTTP or HTTPS

	Certstatus string `json:"certstatus,omitempty"` // Certificate status: INACTIVE or ACTIVE

	Certpassphrase string `json:"certpassphrase,omitempty"` // Certificate password phrase

	Certdownloadurl string `json:"certdownloadurl,omitempty"` // Certificate download URL

	ExtWebAuthVirtualIP string `json:"extWebAuthVirtualIp,omitempty"` // External WEB Auth virtual IP

	ExtWebAuth *bool `json:"extWebAuth,omitempty"` // Indication of using external WEB Auth

	WhiteList *bool `json:"whiteList,omitempty"` // Indication of being on allowed list

	ExtWebAuthPortal string `json:"extWebAuthPortal,omitempty"` // External authentication portal

	ExtWebAuthAccessURL string `json:"extWebAuthAccessUrl,omitempty"` // External WEB Auth access URL

	ExtWebAuthHTMLTag *[]RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag `json:"extWebAuthHtmlTag,omitempty"` //

	QosPolicy string `json:"qosPolicy,omitempty"` // QoS policy: PlATINUM, GOLD, SILVER, BRONZE

	Tests *[]RequestSensorsCreateSensorTestTemplateProfilesTests `json:"tests,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` // Profile name

	DeviceType string `json:"deviceType,omitempty"` // Device Type

	VLAN string `json:"vlan,omitempty"` // VLAN

	LocationVLANList *[]RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList `json:"locationVlanList,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag struct {
	Label string `json:"label,omitempty"` // Label

	Tag string `json:"tag,omitempty"` // Tag

	Value string `json:"value,omitempty"` // Value
}
type RequestSensorsCreateSensorTestTemplateProfilesTests struct {
	Name string `json:"name,omitempty"` // Name of the test

	Config *[]RequestSensorsCreateSensorTestTemplateProfilesTestsConfig `json:"config,omitempty"` //
}
type RequestSensorsCreateSensorTestTemplateProfilesTestsConfig struct {
	Domains []string `json:"domains,omitempty"` // DNS domain name

	Server string `json:"server,omitempty"` // Ping, file transfer, mail, radius, ssh, or telnet server

	UserName string `json:"userName,omitempty"` // User name

	Password string `json:"password,omitempty"` // Password

	URL string `json:"url,omitempty"` // URL

	Port *int `json:"port,omitempty"` // Radius or WEB server port

	Protocol string `json:"protocol,omitempty"` // Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)

	Servers []string `json:"servers,omitempty"` // IPerf server list

	Direction string `json:"direction,omitempty"` // IPerf direction (UPLOAD, DOWNLOAD, BOTH)

	StartPort *int `json:"startPort,omitempty"` // IPerf start port

	EndPort *int `json:"endPort,omitempty"` // IPerf end port

	UDPBandwidth *int `json:"udpBandwidth,omitempty"` // IPerf UDP bandwidth

	ProbeType string `json:"probeType,omitempty"` // Probe type

	NumPackets *int `json:"numPackets,omitempty"` // Number of packets

	PathToDownload string `json:"pathToDownload,omitempty"` // File path for file transfer

	TransferType string `json:"transferType,omitempty"` // File transfer type (UPLOAD, DOWNLOAD, BOTH)

	SharedSecret string `json:"sharedSecret,omitempty"` // Shared secret

	NdtServer string `json:"ndtServer,omitempty"` // NDT server

	NdtServerPort string `json:"ndtServerPort,omitempty"` // NDT server port

	NdtServerPath string `json:"ndtServerPath,omitempty"` // NDT server path

	UplinkTest *bool `json:"uplinkTest,omitempty"` // Uplink test

	DownlinkTest *bool `json:"downlinkTest,omitempty"` // Downlink test

	ProxyServer string `json:"proxyServer,omitempty"` // Proxy server

	ProxyPort string `json:"proxyPort,omitempty"` // Proxy port

	ProxyUserName string `json:"proxyUserName,omitempty"` // Proxy user name

	ProxyPassword string `json:"proxyPassword,omitempty"` // Proxy password

	UserNamePrompt string `json:"userNamePrompt,omitempty"` // User name prompt

	PasswordPrompt string `json:"passwordPrompt,omitempty"` // Password prompt

	ExitCommand string `json:"exitCommand,omitempty"` // Exit command

	FinalPrompt string `json:"finalPrompt,omitempty"` // Final prompt
}
type RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	VLANs []string `json:"vlans,omitempty"` // Array of VLANs
}
type RequestSensorsCreateSensorTestTemplateLocationInfoList struct {
	LocationID string `json:"locationId,omitempty"` // Site UUID

	LocationType string `json:"locationType,omitempty"` // Site type

	AllSensors *bool `json:"allSensors,omitempty"` // Use all sensors in the site for test

	SiteHierarchy string `json:"siteHierarchy,omitempty"` // Site name hierarhy

	MacAddressList []string `json:"macAddressList,omitempty"` // MAC addresses

	ManagementVLAN string `json:"managementVlan,omitempty"` // Management VLAN

	CustomManagementVLAN *bool `json:"customManagementVlan,omitempty"` // Custom Management VLAN
}
type RequestSensorsCreateSensorTestTemplateSensors struct {
	Name string `json:"name,omitempty"` // Sensor name

	MacAddress string `json:"macAddress,omitempty"` // MAC address

	SwitchMac string `json:"switchMac,omitempty"` // Switch MAC address

	SwitchUUID string `json:"switchUuid,omitempty"` // Switch device UUID

	SwitchSerialNumber string `json:"switchSerialNumber,omitempty"` // Switch serial number

	MarkedForUninstall *bool `json:"markedForUninstall,omitempty"` // Is marked for uninstall

	IPAddress string `json:"ipAddress,omitempty"` // IP address

	HostName string `json:"hostName,omitempty"` // Host name

	WiredApplicationStatus string `json:"wiredApplicationStatus,omitempty"` // Wired application status

	WiredApplicationMessage string `json:"wiredApplicationMessage,omitempty"` // Wired application message

	Assigned *bool `json:"assigned,omitempty"` // Is assigned

	Status string `json:"status,omitempty"` // Sensor device status: UP, DOWN, REBOOT

	XorSensor *bool `json:"xorSensor,omitempty"` // Is XOR sensor

	TargetAPs []string `json:"targetAPs,omitempty"` // Array of target APs

	RunNow string `json:"runNow,omitempty"` // Run now: YES, NO

	LocationID string `json:"locationId,omitempty"` // Site UUID

	AllSensorAddition *bool `json:"allSensorAddition,omitempty"` // Is all sensor addition

	ConfigUpdated string `json:"configUpdated,omitempty"` // Configuration updated: YES, NO

	SensorType string `json:"sensorType,omitempty"` // Sensor type

	TestMacAddresses *RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses `json:"testMacAddresses,omitempty"` // A string-string test MAC address

	ID string `json:"id,omitempty"` // Sensor ID

	ServicePolicy string `json:"servicePolicy,omitempty"` // Service policy

	IPerfInfo *RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo `json:"iPerfInfo,omitempty"` // A string-stringList iPerf information
}
type RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses interface{}
type RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo interface{}
type RequestSensorsCreateSensorTestTemplateApCoverage struct {
	Bands string `json:"bands,omitempty"` // The WIFI bands

	NumberOfApsToTest *int `json:"numberOfApsToTest,omitempty"` // Number of APs to test

	RssiThreshold *int `json:"rssiThreshold,omitempty"` // RSSI threshold
}
type RequestSensorsRunNowSensorTest struct {
	TemplateName string `json:"templateName,omitempty"` // Template Name
}
type RequestSensorsDuplicateSensorTestTemplate struct {
	TemplateName string `json:"templateName,omitempty"` // Source test template name

	NewTemplateName string `json:"newTemplateName,omitempty"` // Destination test template name
}

//ListsICapPacketCaptureFilesMatchingSpecifiedCriteria Lists ICAP packet capture files matching specified criteria - b88d-fbd6-4639-9b2f
/* Lists the ICAP packet capture (pcap) files matching the specified criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams Custom header parameters
@param ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!lists-i-cap-packet-capture-files-matching-specified-criteria
*/
func (s *SensorsService) ListsICapPacketCaptureFilesMatchingSpecifiedCriteria(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams *ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams, ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams *ListsICapPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams) (*ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteria, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles"

	queryString, _ := query.Values(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams != nil {

		if ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteria{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ListsICapPacketCaptureFilesMatchingSpecifiedCriteria(ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams, ListsICAPPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ListsICapPacketCaptureFilesMatchingSpecifiedCriteria")
	}

	result := response.Result().(*ResponseSensorsListsICapPacketCaptureFilesMatchingSpecifiedCriteria)
	return result, response, err

}

//RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria Retrieves the total number of packet capture files matching specified criteria - 5483-e972-4b1a-98c6
/* Retrieves the total number of packet capture files matching the specified criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams Custom header parameters
@param RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-total-number-of-packet-capture-files-matching-specified-criteria
*/
func (s *SensorsService) RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams *RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams, RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams *RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams) (*ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles/count"

	queryString, _ := query.Values(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams != nil {

		if RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria(RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaHeaderParams, RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteriaQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheTotalNumberOfPacketCaptureFilesMatchingSpecifiedCriteria)
	return result, response, err

}

//RetrievesDetailsOfASpecificICapPacketCaptureFile Retrieves details of a specific ICAP packet capture file - 7bb7-5afe-4cc8-be27
/* Retrieves details of a specific ICAP packet capture file. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.

@param RetrievesDetailsOfASpecificICAPPacketCaptureFileHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-details-of-a-specific-i-cap-packet-capture-file
*/
func (s *SensorsService) RetrievesDetailsOfASpecificICapPacketCaptureFile(id string, RetrievesDetailsOfASpecificICAPPacketCaptureFileHeaderParams *RetrievesDetailsOfASpecificICapPacketCaptureFileHeaderParams) (*ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFile, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesDetailsOfASpecificICAPPacketCaptureFileHeaderParams != nil {

		if RetrievesDetailsOfASpecificICAPPacketCaptureFileHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesDetailsOfASpecificICAPPacketCaptureFileHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesDetailsOfASpecificICapPacketCaptureFile(id, RetrievesDetailsOfASpecificICAPPacketCaptureFileHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesDetailsOfASpecificICapPacketCaptureFile")
	}

	result := response.Result().(*ResponseSensorsRetrievesDetailsOfASpecificICapPacketCaptureFile)
	return result, response, err

}

//DownloadsASpecificICapPacketCaptureFile Downloads a specific ICAP packet capture file - 47b2-db4f-4468-ba5d
/* Downloads a specific ICAP packet capture file. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. The name of the packet capture file, as given by the GET /captureFiles API response.

@param DownloadsASpecificICAPPacketCaptureFileHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!downloads-a-specific-i-cap-packet-capture-file
*/
func (s *SensorsService) DownloadsASpecificICapPacketCaptureFile(id string, DownloadsASpecificICAPPacketCaptureFileHeaderParams *DownloadsASpecificICapPacketCaptureFileHeaderParams) (*ResponseSensorsDownloadsASpecificICapPacketCaptureFile, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/captureFiles/{id}/download"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if DownloadsASpecificICAPPacketCaptureFileHeaderParams != nil {

		if DownloadsASpecificICAPPacketCaptureFileHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", DownloadsASpecificICAPPacketCaptureFileHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseSensorsDownloadsASpecificICapPacketCaptureFile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadsASpecificICapPacketCaptureFile(id, DownloadsASpecificICAPPacketCaptureFileHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation DownloadsASpecificICapPacketCaptureFile")
	}

	result := response.Result().(*ResponseSensorsDownloadsASpecificICapPacketCaptureFile)
	return result, response, err

}

//RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac Retrieves the spectrum interference devices reports sent by WLC for provided AP Mac - 78a6-8aec-4278-8cdc
/* Retrieves the spectrum interference devices reports sent by WLC for provided AP Mac. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacHeaderParams Custom header parameters
@param RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-spectrum-interference-devices-reports-sent-by-w-l-c-for-provided-ap-mac
*/
func (s *SensorsService) RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacHeaderParams *RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacHeaderParams, RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacQueryParams *RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMacQueryParams) (*ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/spectrumInterferenceDeviceReports"

	queryString, _ := query.Values(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacHeaderParams != nil {

		if RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac(RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacHeaderParams, RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedAPMacQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheSpectrumInterferenceDevicesReportsSentByWLCForProvidedApMac)
	return result, response, err

}

//RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac Retrieves the spectrum sensor reports sent by WLC for provided AP Mac - f583-d956-4c1a-8cc3
/* Retrieves the spectrum sensor reports sent by WLC for provided AP Mac. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacHeaderParams Custom header parameters
@param RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-spectrum-sensor-reports-sent-by-w-l-c-for-provided-ap-mac
*/
func (s *SensorsService) RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacHeaderParams *RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacHeaderParams, RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacQueryParams *RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMacQueryParams) (*ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/spectrumSensorReports"

	queryString, _ := query.Values(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacHeaderParams != nil {

		if RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac(RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacHeaderParams, RetrievesTheSpectrumSensorReportsSentByWLCForProvidedAPMacQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheSpectrumSensorReportsSentByWLCForProvidedApMac)
	return result, response, err

}

//RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering Retrieves deployed ICAP configurations while supporting basic filtering. - c780-8b33-44c8-858c
/* Retrieves deployed ICAP configurations while supporting basic filtering. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-deployed-i-cap-configurations-while-supporting-basic-filtering
*/
func (s *SensorsService) RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams *RetrievesDeployedICapConfigurationsWhileSupportingBasicFilteringQueryParams) (*ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings"

	queryString, _ := query.Values(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering(RetrievesDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering")
	}

	result := response.Result().(*ResponseSensorsRetrievesDeployedICapConfigurationsWhileSupportingBasicFiltering)
	return result, response, err

}

//GetICapConfigurationStatusPerNetworkDevice Get ICAP configuration status per network device. - 5291-5b76-46ab-98eb
/* Get ICAP configuration status per network device using the activity ID, which was returned in property "taskId" of the TaskResponse of the POST. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-i-cap-configuration-status-per-network-device
*/
func (s *SensorsService) GetICapConfigurationStatusPerNetworkDevice(previewActivityID string) (*ResponseSensorsGetICapConfigurationStatusPerNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/networkDeviceStatusDetails"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorsGetICapConfigurationStatusPerNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetICapConfigurationStatusPerNetworkDevice(previewActivityID)
		}
		return nil, response, fmt.Errorf("error with operation GetICapConfigurationStatusPerNetworkDevice")
	}

	result := response.Result().(*ResponseSensorsGetICapConfigurationStatusPerNetworkDevice)
	return result, response, err

}

//RetrievesTheDevicesClisOfTheICapintent Retrieves the device's CLIs of the ICAP intent. - 4288-bbf8-4b59-8fcf
/* Returns the device's CLIs of the ICAP intent. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response

@param networkDeviceID networkDeviceId path parameter. device id from intent/api/v1/network-device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-devices-clis-of-the-i-capintent
*/
func (s *SensorsService) RetrievesTheDevicesClisOfTheICapintent(previewActivityID string, networkDeviceID string) (*ResponseSensorsRetrievesTheDevicesClisOfTheICapintent, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorsRetrievesTheDevicesClisOfTheICapintent{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheDevicesClisOfTheICapintent(previewActivityID, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheDevicesClisOfTheICapintent")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheDevicesClisOfTheICapintent)
	return result, response, err

}

//RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering Retrieves the count of deployed ICAP configurations while supporting basic filtering. - 46b9-694b-41ca-8155
/* Retrieves the count of deployed ICAP configurations while supporting basic filtering. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-deployed-i-cap-configurations-while-supporting-basic-filtering
*/
func (s *SensorsService) RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams *RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFilteringQueryParams) (*ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/count"

	queryString, _ := query.Values(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering(RetrievesTheCountOfDeployedICAPConfigurationsWhileSupportingBasicFilteringQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering")
	}

	result := response.Result().(*ResponseSensorsRetrievesTheCountOfDeployedICapConfigurationsWhileSupportingBasicFiltering)
	return result, response, err

}

//GetDeviceDeploymentStatus Get device deployment status. - 5aa1-1ac0-4d28-af86
/* Retrieves ICAP configuration deployment status(s) per device based on filter criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param GetDeviceDeploymentStatusQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-deployment-status
*/
func (s *SensorsService) GetDeviceDeploymentStatus(GetDeviceDeploymentStatusQueryParams *GetDeviceDeploymentStatusQueryParams) (*ResponseSensorsGetDeviceDeploymentStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deviceDeployments"

	queryString, _ := query.Values(GetDeviceDeploymentStatusQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsGetDeviceDeploymentStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDeploymentStatus(GetDeviceDeploymentStatusQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDeploymentStatus")
	}

	result := response.Result().(*ResponseSensorsGetDeviceDeploymentStatus)
	return result, response, err

}

//GetDeviceDeploymentStatusCount Get device deployment status count. - 6d80-6a7b-4459-a1a3
/* Returns the count of device deployment status(s) based on filter criteria. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param GetDeviceDeploymentStatusCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-deployment-status-count
*/
func (s *SensorsService) GetDeviceDeploymentStatusCount(GetDeviceDeploymentStatusCountQueryParams *GetDeviceDeploymentStatusCountQueryParams) (*ResponseSensorsGetDeviceDeploymentStatusCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deviceDeployments/count"

	queryString, _ := query.Values(GetDeviceDeploymentStatusCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSensorsGetDeviceDeploymentStatusCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDeploymentStatusCount(GetDeviceDeploymentStatusCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDeploymentStatusCount")
	}

	result := response.Result().(*ResponseSensorsGetDeviceDeploymentStatusCount)
	return result, response, err

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

//RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime Retrieves specific client statistics over specified period of time. - 8e8d-29bf-4019-95fb
/* Retrieves the time series statistics of a specific client by applying complex filters. If startTime and endTime are not provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. id is the client mac address. It can be specified in one of the notational conventions 01:23:45:67:89:AB or 01-23-45-67-89-AB or 0123.4567.89AB and is case insensitive

@param RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-client-statistics-over-specified-period-of-time
*/
func (s *SensorsService) RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime(id string, requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime *RequestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime, RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams *RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams) (*ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/clients/{id}/stats"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams != nil {

		if RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime).
		SetResult(&ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime(id, requestSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime, RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTimeHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime")
	}

	result := response.Result().(*ResponseSensorsRetrievesSpecificClientStatisticsOverSpecifiedPeriodOfTime)
	return result, response, err

}

//RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime Retrieves specific radio statistics over specified period of time. - 9fac-b838-4c29-b3b5
/* Retrieves the time series statistics of a specific radio by applying complex filters. If startTime and endTime are not provided, the API defaults to the last 24 hours. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-icap-1.0.0-resolved.yaml


@param id id path parameter. id is the composite key made of AP Base Ethernet macAddress and Radio Slot Id. Format apMac_RadioId

@param RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-specific-radio-statistics-over-specified-period-of-time
*/
func (s *SensorsService) RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime(id string, requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime *RequestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime, RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams *RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams) (*ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime, *resty.Response, error) {
	path := "/dna/data/api/v1/icap/radios/{id}/stats"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams != nil {

		if RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime).
		SetResult(&ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime(id, requestSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime, RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTimeHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime")
	}

	result := response.Result().(*ResponseSensorsRetrievesSpecificRadioStatisticsOverSpecifiedPeriodOfTime)
	return result, response, err

}

//CreatesAnICapConfigurationIntentForPreviewApprove Creates an ICAP configuration intent for preview-approve. - be85-b906-4ea8-acfa
/* This creates an ICAP configuration intent for preview approval. The intent is not deployed to the device until further preview-approve APIs are applied. This API is the first step in the preview-approve workflow, which consists of several APIs. Skipping any API in the process is not recommended for a complete preview-approve use case. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param CreatesAnICAPConfigurationIntentForPreviewApproveQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-an-i-cap-configuration-intent-for-preview-approve
*/
func (s *SensorsService) CreatesAnICapConfigurationIntentForPreviewApprove(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApprove *RequestSensorsCreatesAnICapConfigurationIntentForPreviewApprove, CreatesAnICAPConfigurationIntentForPreviewApproveQueryParams *CreatesAnICapConfigurationIntentForPreviewApproveQueryParams) (*ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApprove, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels"

	queryString, _ := query.Values(CreatesAnICAPConfigurationIntentForPreviewApproveQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApprove).
		SetResult(&ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApprove{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAnICapConfigurationIntentForPreviewApprove(requestSensorsCreatesAnICAPConfigurationIntentForPreviewApprove, CreatesAnICAPConfigurationIntentForPreviewApproveQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAnICapConfigurationIntentForPreviewApprove")
	}

	result := response.Result().(*ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApprove)
	return result, response, err

}

//CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice Creates a ICAP configuration workflow for ICAP intent to remove the ICAP configuration on the device. - e681-4857-4298-9079
/* Creates a ICAP configuration intent to remove the ICAP RFSTATS or ANOMALY configuration from the device. The task has not been applied to the device yet. Subsequent preview-approve workflow APIs must be used to complete the preview-approve process.  The path parameter 'id' can be retrieved from **GET /dna/intent/api/v1/icapSettings** API. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param id id path parameter. A unique ID of the deployed ICAP object, which can be obtained from **GET /dna/intent/api/v1/icapSettings**


Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-i-cap-configuration-workflow-for-i-capintent-to-remove-the-i-cap-configuration-on-the-device
*/
func (s *SensorsService) CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice(id string, requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDevice *RequestSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice) (*ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{id}/deleteDeploy"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDevice).
		SetResult(&ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice(id, requestSensorsCreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDevice)
		}

		return nil, response, fmt.Errorf("error with operation CreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice")
	}

	result := response.Result().(*ResponseSensorsCreatesAICapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice)
	return result, response, err

}

//DeploysTheICapConfigurationIntentByActivityID Deploys the ICAP configuration intent by activity ID. - f58e-6b46-42ba-ae2f
/* Deploys the ICAP configuration intent by activity ID, which was returned in property "taskId" of the TaskResponse of the POST.  POST'ing the intent prior to generating the intent CLI for preview-approve has the same effect as direct-deploy'ing the intent to the device.
Generating of device's CLIs for preview-approve is not available for this activity ID after using this POST API. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploys-the-i-cap-configuration-intent-by-activity-id
*/
func (s *SensorsService) DeploysTheICapConfigurationIntentByActivityID(previewActivityID string, requestSensorsDeploysTheICAPConfigurationIntentByActivityID *RequestSensorsDeploysTheICapConfigurationIntentByActivityID) (*ResponseSensorsDeploysTheICapConfigurationIntentByActivityID, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/deploy"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsDeploysTheICAPConfigurationIntentByActivityID).
		SetResult(&ResponseSensorsDeploysTheICapConfigurationIntentByActivityID{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeploysTheICapConfigurationIntentByActivityID(previewActivityID, requestSensorsDeploysTheICAPConfigurationIntentByActivityID)
		}

		return nil, response, fmt.Errorf("error with operation DeploysTheICapConfigurationIntentByActivityId")
	}

	result := response.Result().(*ResponseSensorsDeploysTheICapConfigurationIntentByActivityID)
	return result, response, err

}

//GeneratesTheDevicesClisOfTheICapConfigurationIntent Generates the device's CLIs of the ICAP configuration intent. - e1b2-aaf0-4358-9325
/* Generates the device's CLIs of the ICAP intent for preview and approve prior to deploying the ICAP configuration intent to the device.  After deploying the configuration intent, generating intent CLIs will not be available for preview.


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response

@param networkDeviceID networkDeviceId path parameter. device id from intent/api/v1/network-device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!generates-the-devices-clis-of-the-i-cap-configuration-intent
*/
func (s *SensorsService) GeneratesTheDevicesClisOfTheICapConfigurationIntent(previewActivityID string, networkDeviceID string, requestSensorsGeneratesTheDevicesCLIsOfTheICAPConfigurationIntent *RequestSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent) (*ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsGeneratesTheDevicesCLIsOfTheICAPConfigurationIntent).
		SetResult(&ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GeneratesTheDevicesClisOfTheICapConfigurationIntent(previewActivityID, networkDeviceID, requestSensorsGeneratesTheDevicesCLIsOfTheICAPConfigurationIntent)
		}

		return nil, response, fmt.Errorf("error with operation GeneratesTheDevicesClisOfTheICapConfigurationIntent")
	}

	result := response.Result().(*ResponseSensorsGeneratesTheDevicesClisOfTheICapConfigurationIntent)
	return result, response, err

}

//DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove Deploys the given ICAP configuration intent without preview and approve. - 78ba-0a27-4808-b722
/* Deploys the given ICAP intent without preview and approval. The response body contains a task object with a taskId and a URL for more information about the task. The deployment status of this ICAP intent can be found in the output of the URL.  For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploys-the-given-i-cap-configuration-intent-without-preview-and-approve
*/
func (s *SensorsService) DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApprove *RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove, DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveQueryParams *DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveQueryParams) (*ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deploy"

	queryString, _ := query.Values(DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApprove).
		SetResult(&ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove(requestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApprove, DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove")
	}

	result := response.Result().(*ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApprove)
	return result, response, err

}

//RemoveTheICapConfigurationOnTheDeviceWithoutPreview Remove the ICAP configuration on the device without preview - 779f-7b3d-40e9-b736
/* Remove the ICAP configuration from the device by *id* without preview-deploy. The path parameter *id* can be retrieved from the **GET /dna/intent/api/v1/icapSettings** API. The response body contains a task object with a taskId and a URL. Use the URL to check the task status. ICAP FULL, ONBOARDING, OTA, and SPECTRUM configurations have a durationInMins field. A disable task is scheduled to remove the configuration from the device. Removing the ICAP intent should be done after the pre-scheduled disable task has been deployed. For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param id id path parameter. A unique ID of the deployed ICAP object, which can be obtained from **GET /dna/intent/api/v1/icapSettings**


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-the-i-cap-configuration-on-the-device-without-preview
*/
func (s *SensorsService) RemoveTheICapConfigurationOnTheDeviceWithoutPreview(id string, requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreview *RequestSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreview) (*ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreview, *resty.Response, error) {
	path := "/dna/intent/api/v1/icapSettings/deploy/{id}/deleteDeploy"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreview).
		SetResult(&ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreview{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveTheICapConfigurationOnTheDeviceWithoutPreview(id, requestSensorsRemoveTheICAPConfigurationOnTheDeviceWithoutPreview)
		}

		return nil, response, fmt.Errorf("error with operation RemoveTheICapConfigurationOnTheDeviceWithoutPreview")
	}

	result := response.Result().(*ResponseSensorsRemoveTheICapConfigurationOnTheDeviceWithoutPreview)
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

//DiscardsTheICapConfigurationIntentByActivityID Discards the ICAP configuration intent by activity ID. - 02b1-1ba1-4afb-9155
/* Discard the ICAP configuration intent by activity ID, which was returned in TaskResponse's property "taskId" at the beginning of the preview-approve workflow.  Discarding the intent can only be applied to intent activities that have not been deployed.
Note that ICAP type FULL, ONBOARDING, OTA, and SPECTRUM for the scheduled-disabled task cannot be discarded or cancelled because they have already deployed.  The feature can only be disabled by sending in a direct-deploy DELETE with API /dna/intent/api/v1/icapSettings/deploy/deployedId/{icapDeployedId} For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml


@param previewActivityID previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response


Documentation Link: https://developer.cisco.com/docs/dna-center/#!discards-the-i-cap-configuration-intent-by-activity-id
*/
func (s *SensorsService) DiscardsTheICapConfigurationIntentByActivityID(previewActivityID string) (*ResponseSensorsDiscardsTheICapConfigurationIntentByActivityID, *resty.Response, error) {
	//previewActivityID string
	path := "/dna/intent/api/v1/icapSettings/configurationModels/{previewActivityId}"
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSensorsDiscardsTheICapConfigurationIntentByActivityID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DiscardsTheICapConfigurationIntentByActivityID(previewActivityID)
		}
		return nil, response, fmt.Errorf("error with operation DiscardsTheICapConfigurationIntentByActivityId")
	}

	result := response.Result().(*ResponseSensorsDiscardsTheICapConfigurationIntentByActivityID)
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
