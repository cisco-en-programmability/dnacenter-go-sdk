package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ComplianceService service

type GetComplianceStatusQueryParams struct {
	ComplianceStatus string  `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
	DeviceUUID       string  `url:"deviceUuid,omitempty"`       //Comma separated 'Device Ids'
	Offset           float64 `url:"offset,omitempty"`           //offset/starting row	number
	Limit            float64 `url:"limit,omitempty"`            //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type GetComplianceStatusCountQueryParams struct {
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
}
type GetComplianceDetailQueryParams struct {
	ComplianceType   string  `url:"complianceType,omitempty"`   //Specify "Compliance type(s)" in commas. The Compliance type can be 'NETWORK_PROFILE', 'IMAGE', 'FABRIC', 'APPLICATION_VISIBILITY', 'FABRIC', RUNNING_CONFIG', 'NETWORK_SETTINGS', 'WORKFLOW' , 'EOX'.
	ComplianceStatus string  `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" in commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
	DeviceUUID       string  `url:"deviceUuid,omitempty"`       //Comma separated "Device Id(s)"
	Offset           float64 `url:"offset,omitempty"`           //offset/starting row
	Limit            float64 `url:"limit,omitempty"`            //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type GetComplianceDetailCountQueryParams struct {
	ComplianceType   string `url:"complianceType,omitempty"`   //Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EOX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
}
type ComplianceDetailsOfDeviceQueryParams struct {
	Category             string `url:"category,omitempty"`             //category can have any value among 'INTENT', 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'DESIGN_OOD' , 'EOX' , 'NETWORK_SETTINGS'
	ComplianceType       string `url:"complianceType,omitempty"`       //Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EOX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	DiffList             bool   `url:"diffList,omitempty"`             //diff list [ pass true to fetch the diff list ]
	Status               string `url:"status,omitempty"`               //'COMPLIANT', 'NON_COMPLIANT', 'ERROR', 'IN_PROGRESS', 'NOT_APPLICABLE', 'NOT_AVAILABLE', 'WARNING', 'REMEDIATION_IN_PROGRESS' can be the value of the compliance 'status' parameter. [COMPLIANT: Device currently meets the compliance requirements.  NON_COMPLIANT: One of the compliance requirements like Software Image, PSIRT, Network Profile, Startup vs Running, etc. are not met. ERROR: Compliance is unable to compute status due to underlying errors. IN_PROGRESS: Compliance check is in progress for the device. NOT_APPLICABLE: Device is not supported for compliance, or minimum license requirement is not met. NOT_AVAILABLE: Compliance is not available for the device. COMPLIANT_WARNING: The device is compliant with warning if the last date of support is nearing. REMEDIATION_IN_PROGRESS: Compliance remediation is in progress for the device.]
	RemediationSupported bool   `url:"remediationSupported,omitempty"` //The 'remediationSupported' parameter can be set to 'true' or 'false'. The result will be a combination of both values if it is not provided.
}
type GetFieldNoticeNetworkDevicesQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
	NoticeCount     float64 `url:"noticeCount,omitempty"`     //Return network devices with noticeCount greater than this noticeCount
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticeNetworkDevicesQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
	NoticeCount     float64 `url:"noticeCount,omitempty"`     //Return network devices with noticeCount greater than this noticeCount
}
type GetFieldNoticesAffectingTheNetworkDeviceQueryParams struct {
	ID     string  `url:"id,omitempty"`     //Id of the field notice
	Type   string  `url:"type,omitempty"`   //Return field notices with this type. Available values : SOFTWARE, HARDWARE.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams struct {
	ID   string `url:"id,omitempty"`   //Id of the field notice
	Type string `url:"type,omitempty"` //Return field notices with this type. Available values : SOFTWARE, HARDWARE
}
type GetFieldNoticesQueryParams struct {
	ID          string  `url:"id,omitempty"`          //Id of the field notice
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return field notices with deviceCount greater than this deviceCount
	Type        string  `url:"type,omitempty"`        //Return field notices with this type. Available values : SOFTWARE, HARDWARE.
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy      string  `url:"sortBy,omitempty"`      //A property within the response to sort by.
	Order       string  `url:"order,omitempty"`       //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticesQueryParams struct {
	ID          string  `url:"id,omitempty"`          //Id of the field notice
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return field notices with deviceCount greater than this deviceCount
	Type        string  `url:"type,omitempty"`        //Return field notices with this type. Available values : SOFTWARE, HARDWARE
}
type GetFieldNoticeNetworkDevicesForTheNoticeQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //id of the network device
	ScanStatus      string `url:"scanStatus,omitempty"`      //status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
}
type GetFieldNoticesResultsTrendOverTimeQueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return field notices trend with scanTime greater than this scanTime
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
}
type GetCountOfFieldNoticesResultsTrendOverTimeQueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return field notices trend with scanTime greater than this scanTime
}
type TriggersAFieldNoticesScanForTheSupportedNetworkDevicesQueryParams struct {
	FailedDevicesOnly bool `url:"failedDevicesOnly,omitempty"` //Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
}
type GetConfigTaskDetailsQueryParams struct {
	ParentTaskID string `url:"parentTaskId,omitempty"` //task Id
}
type GetNetworkBugsQueryParams struct {
	ID          string  `url:"id,omitempty"`          //The id of the network bug
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return network bugs with deviceCount greater than this deviceCount
	Severity    string  `url:"severity,omitempty"`    //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. Default value is 1
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy      string  `url:"sortBy,omitempty"`      //A property within the response to sort by.
	Order       string  `url:"order,omitempty"`       //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc
}
type GetCountOfNetworkBugsQueryParams struct {
	ID          string  `url:"id,omitempty"`          //Id of the network bug
	DeviceCount float64 `url:"deviceCount,omitempty"` //Return network bugs with deviceCount greater than this deviceCount
	Severity    string  `url:"severity,omitempty"`    //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
}
type GetNetworkBugDevicesForTheBugQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc
}
type GetCountOfNetworkBugDevicesForTheBugQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
}
type GetNetworkBugDevicesQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
	BugCount        float64 `url:"bugCount,omitempty"`        //Return network devices with bugCount greater than this bugCount
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc
}
type GetCountOfNetworkBugDevicesQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK
	BugCount        float64 `url:"bugCount,omitempty"`        //Return network devices with bugCount greater than this bugCount
}
type GetBugsAffectingTheNetworkDeviceQueryParams struct {
	ID       string  `url:"id,omitempty"`       //Id of the network bug
	Severity string  `url:"severity,omitempty"` //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE.
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy   string  `url:"sortBy,omitempty"`   //A property within the response to sort by.
	Order    string  `url:"order,omitempty"`    //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfBugsAffectingTheNetworkDeviceQueryParams struct {
	ID       string `url:"id,omitempty"`       //Id of the network bug
	Severity string `url:"severity,omitempty"` //Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
}
type GetNetworkBugsResultsTrendOverTimeQueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return bugs trend with scanTime greater than this scanTime
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
}
type GetCountOfNetworkBugsResultsTrendOverTimeQueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return bugs trend with scanTime greater than this scanTime
}
type TriggersABugsScanForTheSupportedNetworkDevicesQueryParams struct {
	FailedDevicesOnly bool `url:"failedDevicesOnly,omitempty"` //Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
}
type GetSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams struct {
	ID                   string  `url:"id,omitempty"`                   //Id of the advisory
	DeviceCount          float64 `url:"deviceCount,omitempty"`          //Return advisories with deviceCount greater than this deviceCount
	CvssBaseScore        string  `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string  `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
	Offset               float64 `url:"offset,omitempty"`               //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit                float64 `url:"limit,omitempty"`                //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy               string  `url:"sortBy,omitempty"`               //A property within the response to sort by.
	Order                string  `url:"order,omitempty"`                //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams struct {
	ID                   string  `url:"id,omitempty"`                   //Id of the security advisory
	DeviceCount          float64 `url:"deviceCount,omitempty"`          //Return advisories with deviceCount greater than this deviceCount
	CvssBaseScore        string  `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string  `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
}
type GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
}
type GetSecurityAdvisoryNetworkDevicesQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
	AdvisoryCount   string  `url:"advisoryCount,omitempty"`   //Return network devices with advisoryCount greater than this advisoryCount
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy          string  `url:"sortBy,omitempty"`          //A property within the response to sort by.
	Order           string  `url:"order,omitempty"`           //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoryNetworkDevicesQueryParams struct {
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Id of the network device
	ScanMode        string  `url:"scanMode,omitempty"`        //Mode or the criteria using which the network device was scanned. Available values : ESSENTIALS, ADVANTAGE, CX_CLOUD, NOT_AVAILABLE
	ScanStatus      string  `url:"scanStatus,omitempty"`      //Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED, FALL_BACK.
	AdvisoryCount   float64 `url:"advisoryCount,omitempty"`   //Return network devices with advisoryCount greater than this advisoryCount
}
type GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams struct {
	ID                   string  `url:"id,omitempty"`                   //Id of the security advisory
	CvssBaseScore        string  `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string  `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
	Offset               float64 `url:"offset,omitempty"`               //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit                float64 `url:"limit,omitempty"`                //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
	SortBy               string  `url:"sortBy,omitempty"`               //A property within the response to sort by.
	Order                string  `url:"order,omitempty"`                //Whether ascending or descending order should be used to sort the response. Available values : asc, desc. Default value is asc.
}
type GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams struct {
	ID                   string `url:"id,omitempty"`                   //Id of the security advisory
	CvssBaseScore        string `url:"cvssBaseScore,omitempty"`        //Return advisories with cvssBaseScore greater than this cvssBaseScore. E.g. : 8.5
	SecurityImpactRating string `url:"securityImpactRating,omitempty"` //Return advisories with this securityImpactRating. Available values : CRITICAL, HIGH.
}
type GetSecurityAdvisoriesResultsTrendOverTimeQueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return advisories trend with scanTime greater than this scanTime
	Offset   float64 `url:"offset,omitempty"`   //The first record to show for this page; the first record is numbered 1. Default value is 1.
	Limit    float64 `url:"limit,omitempty"`    //The number of records to show for this page. Minimum value is 1. Maximum value is 500. Default value is 500.
}
type GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams struct {
	ScanTime float64 `url:"scanTime,omitempty"` //Return advisories trend with scanTime greater than this scanTime
}
type TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesQueryParams struct {
	FailedDevicesOnly bool `url:"failedDevicesOnly,omitempty"` //Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
}

type ResponseComplianceGetComplianceStatus struct {
	Version  string                                           `json:"version,omitempty"`  // Version of the API.
	Response *[]ResponseComplianceGetComplianceStatusResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetComplianceStatusResponse struct {
	DeviceUUID       string   `json:"deviceUuid,omitempty"`       // UUID of the device.
	ComplianceStatus string   `json:"complianceStatus,omitempty"` // Current compliance status for the compliance type that will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	Message          string   `json:"message,omitempty"`          // Additional message of compliance status for the compliance type.
	ScheduleTime     *float64 `json:"scheduleTime,omitempty"`     // Timestamp when compliance is scheduled to run.
	LastUpdateTime   *float64 `json:"lastUpdateTime,omitempty"`   // Timestamp when the latest compliance checks ran.
}
type ResponseComplianceRunCompliance struct {
	Version  string                                   `json:"version,omitempty"`  // Version of the API.
	Response *ResponseComplianceRunComplianceResponse `json:"response,omitempty"` //
}
type ResponseComplianceRunComplianceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id.
	URL    string `json:"url,omitempty"`    // Additional url for task id.
}
type ResponseComplianceGetComplianceStatusCount struct {
	Version  string   `json:"version,omitempty"`  // Version of the API.
	Response *float64 `json:"response,omitempty"` // Returns count of compliant status
}
type ResponseComplianceGetComplianceDetail struct {
	Version  string                                           `json:"version,omitempty"`  // Version of the API.
	Response *[]ResponseComplianceGetComplianceDetailResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetComplianceDetailResponse struct {
	ComplianceType       string   `json:"complianceType,omitempty"`       // Compliance type corresponds to a tile on the UI. Will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EOX.
	LastSyncTime         *float64 `json:"lastSyncTime,omitempty"`         // Timestamp when the status changed from different value to the current value.
	DeviceUUID           string   `json:"deviceUuid,omitempty"`           // UUID of the device.
	DisplayName          string   `json:"displayName,omitempty"`          // User friendly name for the configuration.
	Status               string   `json:"status,omitempty"`               // Current status of compliance for the complianceType. Will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	Category             string   `json:"category,omitempty"`             // category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EOX' , 'NETWORK_SETTINGS'.
	LastUpdateTime       *float64 `json:"lastUpdateTime,omitempty"`       // Timestamp when the latest compliance checks ran.
	State                string   `json:"state,omitempty"`                // State of latest compliance check for the complianceType. Will be one of SUCCESS, FAILED, or IN_PROGRESS.
	RemediationSupported *bool    `json:"remediationSupported,omitempty"` // Indicates whether remediation is supported for this compliance type or not
}
type ResponseComplianceGetComplianceDetailCount struct {
	Version  string   `json:"version,omitempty"`  // Version of API.
	Response *float64 `json:"response,omitempty"` // Count of all devices or devices that match the query parameters.
}
type ResponseComplianceComplianceRemediation struct {
	Response *ResponseComplianceComplianceRemediationResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // Version of API.
}
type ResponseComplianceComplianceRemediationResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task.
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task.
}
type ResponseComplianceDeviceComplianceStatus struct {
	Response *ResponseComplianceDeviceComplianceStatusResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version of the API.
}
type ResponseComplianceDeviceComplianceStatusResponse struct {
	DeviceUUID       string   `json:"deviceUuid,omitempty"`       // UUID of the device.
	ComplianceStatus string   `json:"complianceStatus,omitempty"` // Current compliance status of the device that will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
	LastUpdateTime   *float64 `json:"lastUpdateTime,omitempty"`   // Timestamp when the latest compliance checks ran.
	ScheduleTime     string   `json:"scheduleTime,omitempty"`     // Timestamp when the next compliance checks will run.
}
type ResponseComplianceComplianceDetailsOfDevice struct {
	Response   *[]ResponseComplianceComplianceDetailsOfDeviceResponse `json:"response,omitempty"`   //
	DeviceUUID string                                                 `json:"deviceUuid,omitempty"` // UUID of the device.
}
type ResponseComplianceComplianceDetailsOfDeviceResponse struct {
	DeviceUUID           string                                                               `json:"deviceUuid,omitempty"`           // UUID of the device.
	ComplianceType       string                                                               `json:"complianceType,omitempty"`       // Compliance type corresponds to a tile on the UI that will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EOX.
	Status               string                                                               `json:"status,omitempty"`               // Status of compliance for the compliance type, will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, WARNING, REMEDIATION_IN_PROGRESS can be the value of the compliance status parameter. [COMPLIANT: Device currently meets the compliance requirements.  NON_COMPLIANT: One of the compliance requirements like Software Image, PSIRT, Network Profile, Startup vs Running, etc. are not met. ERROR: Compliance is unable to compute status due to underlying errors. IN_PROGRESS: Compliance check is in progress for the device. NOT_APPLICABLE: Device is not supported for compliance, or minimum license requirement is not met. NOT_AVAILABLE: Compliance is not available for the device. COMPLIANT_WARNING: The device is compliant with warning if the last date of support is nearing. REMEDIATION_IN_PROGRESS: Compliance remediation is in progress for the device.]
	State                string                                                               `json:"state,omitempty"`                // State of the compliance check for the compliance type, will be one of SUCCESS, FAILED, or IN_PROGRESS.
	LastSyncTime         *float64                                                             `json:"lastSyncTime,omitempty"`         // Timestamp when the status changed from a different value to the current value.
	LastUpdateTime       *float64                                                             `json:"lastUpdateTime,omitempty"`       // Timestamp of the latest compliance check that was run.
	SourceInfoList       *[]ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoList `json:"sourceInfoList,omitempty"`       //
	AckStatus            string                                                               `json:"ackStatus,omitempty"`            // Acknowledgment status of the compliance type. UNACKNOWLEDGED if none of the violations under the compliance type are acknowledged. Otherwise it will be ACKNOWLEDGED.
	Version              string                                                               `json:"version,omitempty"`              // Version of the API.
	RemediationSupported *bool                                                                `json:"remediationSupported,omitempty"` // Indicates whether remediation is supported for this compliance type or not.
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoList struct {
	Name                string                                                                        `json:"name,omitempty"`                // Name of the type of top level configuration.
	NameWithBusinessKey string                                                                        `json:"nameWithBusinessKey,omitempty"` // Name With Business Key
	SourceEnum          string                                                                        `json:"sourceEnum,omitempty"`          // Will be same as compliance type.
	Type                string                                                                        `json:"type,omitempty"`                // Type of the top level configuration.
	AppName             string                                                                        `json:"appName,omitempty"`             // Application name that is used to club the violations.
	Count               *float64                                                                      `json:"count,omitempty"`               // Number of violations present.
	AckStatus           string                                                                        `json:"ackStatus,omitempty"`           // Acknowledgment status of violations. UNACKNOWLEDGED if none of the violations are acknowledged. Otherwise it will be ACKNOWLEDGED.
	BusinessKey         *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKey `json:"businessKey,omitempty"`         //
	DiffList            *[]ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffList  `json:"diffList,omitempty"`            //
	DisplayName         string                                                                        `json:"displayName,omitempty"`         // Model display name.
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKey struct {
	ResourceName          string                                                                                             `json:"resourceName,omitempty"`          // Name of the top level resource. Same as name above.
	BusinessKeyAttributes *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyBusinessKeyAttributes `json:"businessKeyAttributes,omitempty"` // Attributes that together uniquely identify the configuration instance.
	OtherAttributes       *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributes       `json:"otherAttributes,omitempty"`       //
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyBusinessKeyAttributes interface{}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributes struct {
	Name          string                                                                                                    `json:"name,omitempty"`          // Name of the attributes.
	CfsAttributes *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributesCfsAttributes `json:"cfsAttributes,omitempty"` //
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributesCfsAttributes struct {
	DisplayName string `json:"displayName,omitempty"` // User friendly name for the configuration.
	AppName     string `json:"appName,omitempty"`     // Same as appName above.
	Description string `json:"description,omitempty"` // Description for the configuration, if available.
	Source      string `json:"source,omitempty"`      // Will be same as compliance type.
	Type        string `json:"type,omitempty"`        // The type of this attribute (for example, type can be Intent).
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffList struct {
	Op                 string                                                                                       `json:"op,omitempty"`                 // Type of change (add, remove, or update).
	ConfiguredValue    string                                                                                       `json:"configuredValue,omitempty"`    // Configured value i.e. running / current value. It will be empty for the template violations due to potentially large size of the template. Use a dedicated API to get the template data.
	IntendedValue      string                                                                                       `json:"intendedValue,omitempty"`      // Enable", Intended value. It will be empty for the template violations due to potentially large size of the template. Use a dedicated API to get the template data.
	MoveFromPath       string                                                                                       `json:"moveFromPath,omitempty"`       // Additional URI to fetch more details, if available.
	BusinessKey        string                                                                                       `json:"businessKey,omitempty"`        // The Unique key of the individual violation does not change after every compliance check, as long as the deployment data doesn't change.
	Path               string                                                                                       `json:"path,omitempty"`               // Path of the configuration relative to the top-level configuration. Use it along with a name to identify certain set of violations.
	ExtendedAttributes *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffListExtendedAttributes `json:"extendedAttributes,omitempty"` //
	AckStatus          string                                                                                       `json:"ackStatus,omitempty"`          // Acknowledgment status of the violation. ACKNOWLEDGED if the violation is acknowledged or at the top-level configuration. Otherwise it will be UNACKNOWLEDGED.
	InstanceUUID       string                                                                                       `json:"instanceUUID,omitempty"`       // UUID of the individual violation. Changes after every compliance check.
	DisplayName        string                                                                                       `json:"displayName,omitempty"`        // Display name for attribute in ui .If business key is null or of type owning entity type.
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffListExtendedAttributes struct {
	AttributeDisplayName string `json:"attributeDisplayName,omitempty"` // Display name for attribute in ui .if business key is null or only owning entity type.
	Path                 string `json:"path,omitempty"`                 // Path to be displayed on the UI, instead of the above path, if available.
	DataConverter        string `json:"dataConverter,omitempty"`        // Name of the converter used to display configurations in user-friendly format, if available.
	Type                 string `json:"type,omitempty"`                 // Type of this attribute.(example type can be Intent)
}
type ResponseComplianceGetFieldNoticeNetworkDevices struct {
	Response *[]ResponseComplianceGetFieldNoticeNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDevicesResponse struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevices struct {
	Version  string                                                         `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticeNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceID struct {
	Response *ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceIDResponse struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetFieldNoticesAffectingTheNetworkDevice struct {
	Response *[]ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticesAffectingTheNetworkDeviceResponse struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       *int   `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
	MatchConfidence      string `json:"matchConfidence,omitempty"`      // 'VULNERABLE' - network device is vulnerable to the field notice. 'POTENTIALLY_VULNERABLE' - network device is potentially vulnerable to the field notice. additional manual verifications are needed.
	MatchReason          string `json:"matchReason,omitempty"`          // If the MatchConfidence is POTENTIALLY_VULNERABLE, this gives more details such as what was matched and if additional manual verifications are needed.
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
}
type ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDevice struct {
	Version  string                                                                     `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDeviceResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID struct {
	Response *ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeIDResponse struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       string `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
	MatchConfidence      string `json:"matchConfidence,omitempty"`      // 'VULNERABLE' - network device is vulnerable to the field notice. 'POTENTIALLY_VULNERABLE' - network device is potentially vulnerable to the field notice. additional manual verifications are needed.
	MatchReason          string `json:"matchReason,omitempty"`          // If the MatchConfidence is POTENTIALLY_VULNERABLE, this gives more details such as what was matched and if additional manual verifications are needed.
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
}
type ResponseComplianceGetFieldNotices struct {
	Response *[]ResponseComplianceGetFieldNoticesResponse `json:"response,omitempty"` //
	Version  string                                       `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticesResponse struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       string `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
}
type ResponseComplianceGetCountOfFieldNotices struct {
	Version  string                                            `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeByID struct {
	Response *ResponseComplianceGetFieldNoticeByIDResponse `json:"response,omitempty"` //
	Version  string                                        `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeByIDResponse struct {
	ID                   string `json:"id,omitempty"`                   // Id of the field notice
	Name                 string `json:"name,omitempty"`                 // Name of the field notice
	PublicationURL       string `json:"publicationUrl,omitempty"`       // Url for getting field notice details on cisco website
	DeviceCount          *int   `json:"deviceCount,omitempty"`          // Number of devices which are vulnerable to this field notice
	PotentialDeviceCount *int   `json:"potentialDeviceCount,omitempty"` // Number of devices which are potentially vulnerable to this field notice
	Type                 string `json:"type,omitempty"`                 // 'SOFTWARE' - field notice is for the network device software. 'HARDWARE' - field notice is for the network device hardware
	FirstPublishDate     *int   `json:"firstPublishDate,omitempty"`     // Time at which the field notice was published
	LastUpdatedDate      *int   `json:"lastUpdatedDate,omitempty"`      // Time at which the field notice was last updated
}
type ResponseComplianceGetFieldNoticeNetworkDevicesForTheNotice struct {
	Response *[]ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDevicesForTheNoticeResponse struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNotice struct {
	Version  string                                                                     `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID struct {
	Response *ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDResponse `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceIDResponse struct {
	NetworkDeviceID      string `json:"networkDeviceId,omitempty"`      // Id of the device
	NoticeCount          *int   `json:"noticeCount,omitempty"`          // Number of field notices to which the network device is vulnerable
	PotentialNoticeCount *int   `json:"potentialNoticeCount,omitempty"` // Number of potential field notices to which the network device is vulnerable
	ScanStatus           string `json:"scanStatus,omitempty"`           // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed
	Comments             string `json:"comments,omitempty"`             // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime         *int   `json:"lastScanTime,omitempty"`         // Time at which the device was scanned
}
type ResponseComplianceGetFieldNoticesResultsTrendOverTime struct {
	Response *[]ResponseComplianceGetFieldNoticesResultsTrendOverTimeResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetFieldNoticesResultsTrendOverTimeResponse struct {
	ScanTime                  *int `json:"scanTime,omitempty"`                  // End time for the scan
	SoftwareFieldNoticesCount *int `json:"softwareFieldNoticesCount,omitempty"` // Number of field notices of type SOFTWARE
	HardwareFieldNoticesCount *int `json:"hardwareFieldNoticesCount,omitempty"` // Number of field notices of type HARDWARE
}
type ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTime struct {
	Version  string                                                                `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTimeResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevices struct {
	Version  string                                                                           `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevicesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevices struct {
	Version  string                                                                             `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesResponse struct {
	Type                     string `json:"type,omitempty"`                     // Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
	Feature                  string `json:"feature,omitempty"`                  // Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
	ContractLevel            string `json:"contractLevel,omitempty"`            // Contract level for which trial was created. this was used in older versions and exists only for compatibility.
	Active                   *bool  `json:"active,omitempty"`                   // Indicates if the trial is active
	StartTime                *int   `json:"startTime,omitempty"`                // Trial start time; as measured in Unix epoch time in milliseconds
	EndTime                  *int   `json:"endTime,omitempty"`                  // Trial end time; as measured in Unix epoch time in milliseconds
	SecondsRemainingToExpiry *int   `json:"secondsRemainingToExpiry,omitempty"` // Seconds remaining in the trial before it expires. for expired trials this will be 0.
	SecondsSinceExpired      *int   `json:"secondsSinceExpired,omitempty"`      // Seconds elapsed after the trial has expired. for active trials this will be 0.
}
type ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevices struct {
	Version  string                                                                            `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevicesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetConfigTaskDetails struct {
	Version  string                                            `json:"version,omitempty"`  // Version of the API.
	Response *[]ResponseComplianceGetConfigTaskDetailsResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetConfigTaskDetailsResponse struct {
	StartTime       *int   `json:"startTime,omitempty"`       // Timestamp when the task started.
	ErrorCode       string `json:"errorCode,omitempty"`       // Error code if the task failed.
	DeviceID        string `json:"deviceId,omitempty"`        // UUID of the device.
	TaskID          string `json:"taskId,omitempty"`          // UUID of the task.
	TaskStatus      string `json:"taskStatus,omitempty"`      // Status of the task.
	ParentTaskID    string `json:"parentTaskId,omitempty"`    // UUID of the parent task.
	DeviceIPAddress string `json:"deviceIpAddress,omitempty"` // IP address of the device.
	DetailMessage   string `json:"detailMessage,omitempty"`   // Details of the task, if available.
	FailureMessage  string `json:"failureMessage,omitempty"`  // Failure message, if the task failed.
	TaskType        string `json:"taskType,omitempty"`        // Task type can be 0,1,2 etc(ARCHIVE_RUNNING(0),ARCHIVE_STARTUP(1),ARCHIVE_VLAN(2),DEPLOY_RUNNING(3),DEPLOY_STARTUP(4),DEPLOY_VLAN(5),COPY_RUNNING_TO_STARTUP(6))
	CompletionTime  *int   `json:"completionTime,omitempty"`  // Timestamp when the task was completed.
	HostName        string `json:"hostName,omitempty"`        // Host name of the device.
}
type ResponseComplianceCommitDeviceConfiguration struct {
	Version  string                                               `json:"version,omitempty"`  // Version of the API.
	Response *ResponseComplianceCommitDeviceConfigurationResponse `json:"response,omitempty"` //
}
type ResponseComplianceCommitDeviceConfigurationResponse struct {
	URL    string `json:"url,omitempty"`    // Task Id url.
	TaskID string `json:"taskId,omitempty"` // Unique Id of task.
}
type ResponseComplianceGetNetworkBugs struct {
	Response *[]ResponseComplianceGetNetworkBugsResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugsResponse struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetCountOfNetworkBugs struct {
	Version  string                                           `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugsResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugsResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetNetworkBugByID struct {
	Version  string                                       `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetNetworkBugByIDResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetNetworkBugByIDResponse struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetNetworkBugDevicesForTheBug struct {
	Response *[]ResponseComplianceGetNetworkBugDevicesForTheBugResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugDevicesForTheBugResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfNetworkBugDevicesForTheBug struct {
	Version  string                                                          `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugDevicesForTheBugResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugDevicesForTheBugResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceID []ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceID // Array of ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceId
type ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceID struct {
	Response *ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseItemComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceIDResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetNetworkBugDevices struct {
	Response *[]ResponseComplianceGetNetworkBugDevicesResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugDevicesResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfNetworkBugDevices struct {
	Version  string                                                 `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugDevicesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetNetworkBugDeviceByDeviceID []ResponseItemComplianceGetNetworkBugDeviceByDeviceID // Array of ResponseComplianceGetNetworkBugDeviceByDeviceId
type ResponseItemComplianceGetNetworkBugDeviceByDeviceID struct {
	Response *ResponseItemComplianceGetNetworkBugDeviceByDeviceIDResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  // Version
}
type ResponseItemComplianceGetNetworkBugDeviceByDeviceIDResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	BugCount        *int   `json:"bugCount,omitempty"`        // Number of bugs to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetBugsAffectingTheNetworkDevice struct {
	Response *[]ResponseComplianceGetBugsAffectingTheNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetBugsAffectingTheNetworkDeviceResponse struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetCountOfBugsAffectingTheNetworkDevice struct {
	Version  string                                                             `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugID struct {
	Version  string                                                                       `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugIDResponse struct {
	ID                 string   `json:"id,omitempty"`                 // Id of the network bug
	Headline           string   `json:"headline,omitempty"`           // Title of the network bug
	PublicationURL     string   `json:"publicationUrl,omitempty"`     // Url for getting network bug details on cisco website
	DeviceCount        *int     `json:"deviceCount,omitempty"`        // Number of devices which are vulnerable to this network bug
	Severity           string   `json:"severity,omitempty"`           // 'CATASTROPHIC' - Reasonably common circumstances cause the entire system to fail, or a major subsystem to stop working. 'SEVERE' - Important functions are unusable. 'MODERATE' - Failures occur in unusual circumstances, or minor features do not work at all.
	HasWorkaround      *bool    `json:"hasWorkaround,omitempty"`      // Indicates if the network bug has a workaround
	Workaround         string   `json:"workaround,omitempty"`         // Workaround if any that exists for the network bug
	AffectedVersions   []string `json:"affectedVersions,omitempty"`   // Versions that are affected by the network bug
	IntegratedReleases []string `json:"integratedReleases,omitempty"` // Versions that have the fix for the network bug
}
type ResponseComplianceGetNetworkBugsResultsTrendOverTime struct {
	Response *[]ResponseComplianceGetNetworkBugsResultsTrendOverTimeResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetNetworkBugsResultsTrendOverTimeResponse struct {
	CatastrophicBugsCount *int `json:"catastrophicBugsCount,omitempty"` // Number of network bugs which have a severity of CATASTROPHIC
	SevereBugsCount       *int `json:"severeBugsCount,omitempty"`       // Number of network bugs which have a severity of SEVERE
	ModerateBugsCount     *int `json:"moderateBugsCount,omitempty"`     // Number of network bugs which have a severity of MODERATE
	ScanTime              *int `json:"scanTime,omitempty"`              // End time for the scan
}
type ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTime struct {
	Version  string                                                               `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevices struct {
	Version  string                                                                   `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevicesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevices struct {
	Version  string                                                                     `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevicesResponse struct {
	Type                     string `json:"type,omitempty"`                     // Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
	Feature                  string `json:"feature,omitempty"`                  // Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
	ContractLevel            string `json:"contractLevel,omitempty"`            // Contract level for which trial was created. this was used in older versions and exists only for compatibility.
	Active                   *bool  `json:"active,omitempty"`                   // Indicates if the trial is active
	StartTime                *int   `json:"startTime,omitempty"`                // Trial start time; as measured in Unix epoch time in milliseconds
	EndTime                  *int   `json:"endTime,omitempty"`                  // Trial end time; as measured in Unix epoch time in milliseconds
	SecondsRemainingToExpiry *int   `json:"secondsRemainingToExpiry,omitempty"` // Seconds remaining in the trial before it expires. for expired trials this will be 0.
	SecondsSinceExpired      *int   `json:"secondsSinceExpired,omitempty"`      // Seconds elapsed after the trial has expired. for active trials this will be 0.
}
type ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevices struct {
	Version  string                                                                    `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevices struct {
	Response *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesResponse struct {
	ID                     string                                                                                             `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                               `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                           `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                             `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                           `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                             `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevicesResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevices struct {
	Version  string                                                                            `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByID struct {
	Response *ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDResponse struct {
	ID                     string                                                                                               `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                                 `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                             `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                               `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                             `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                               `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByIDResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory struct {
	Response *[]ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory struct {
	Version  string                                                                                    `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID struct {
	Response *ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDResponse `json:"response,omitempty"` //
	Version  string                                                                                             `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceIDResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevices struct {
	Response *[]ResponseComplianceGetSecurityAdvisoryNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDevicesResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevices struct {
	Version  string                                                              `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceID struct {
	Response *ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceIDResponse struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Id of the device
	AdvisoryCount   *int   `json:"advisoryCount,omitempty"`   // Number of advisories to which the network device is vulnerable
	ScanMode        string `json:"scanMode,omitempty"`        // 'ESSENTIALS' - the device was scanned using a version based match criteria. 'ADVANTAGE' - the device was scanned using a version based match and user provided config match criteria. 'CX_CLOUD' - the device was scanned using CX cloud engine which uses advanced matching criteria which eliminates false positives.  ‘NOT_AVAILABLE’ - scan mode is not available. e.g. when the device is not scanned
	ScanStatus      string `json:"scanStatus,omitempty"`      // 'NOT_SCANNED' - the device was not scanned. 'IN_PROGRESS' - a scan is in progress for the device. 'SUCCESS' - device scan was successful. 'FAILED' - device scan failed. 'FALL_BACK' - the device was supposed to be scanned using CX_CLOUD but because of connectivity issues fell back to a ESSENTIALS or ADVANTAGE scan
	Comments        string `json:"comments,omitempty"`        // More details about the scan status. ie:- if the scan status is failed, comments will give the reason for failure
	LastScanTime    *int   `json:"lastScanTime,omitempty"`    // Time at which the device was scanned
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevice struct {
	Response *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceResponse struct {
	ID                     string                                                                                            `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                              `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                          `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                            `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                          `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                            `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDeviceResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevice struct {
	Version  string                                                                           `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID struct {
	Response *ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDResponse `json:"response,omitempty"` //
	Version  string                                                                                         `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDResponse struct {
	ID                     string                                                                                                                 `json:"id,omitempty"`                     // Id of the security advisory
	DeviceCount            *int                                                                                                                   `json:"deviceCount,omitempty"`            // Number of devices which are vulnerable to this advisory
	CveIDs                 []string                                                                                                               `json:"cveIds,omitempty"`                 // CVE (Common Vulnerabilities and Exposures) ID of the advisory
	PublicationURL         string                                                                                                                 `json:"publicationUrl,omitempty"`         // Url for getting advisory details on cisco website
	CvssBaseScore          *float64                                                                                                               `json:"cvssBaseScore,omitempty"`          // Common Vulnerability Scoring System(CVSS) base score
	SecurityImpactRating   string                                                                                                                 `json:"securityImpactRating,omitempty"`   // 'CRITICAL' - the advisory requires immediate mitigation. 'HIGH' - the advisory requires priority mitigation
	FirstFixedVersionsList *[]ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDResponseFirstFixedVersionsList `json:"firstFixedVersionsList,omitempty"` //
}
type ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryIDResponseFirstFixedVersionsList struct {
	VulnerableVersion string   `json:"vulnerableVersion,omitempty"` // Version that is vulnerable to the advisory
	FixedVersions     []string `json:"fixedVersions,omitempty"`     // First versions that have the fix for the advisory
}
type ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTime struct {
	Response *[]ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTimeResponse struct {
	CriticalSecurityImpactRatingAdvisoriesCount *int `json:"criticalSecurityImpactRatingAdvisoriesCount,omitempty"` // Number of advisories which have a security impact rating of critical
	HighSecurityImpactRatingAdvisoriesCount     *int `json:"highSecurityImpactRatingAdvisoriesCount,omitempty"`     // Number of advisories which have a security impact rating of high
	ScanTime                                    *int `json:"scanTime,omitempty"`                                    // End time for the scan
}
type ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTime struct {
	Version  string                                                                      `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTimeResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices struct {
	Version  string                                                                                   `json:"version,omitempty"`  // Version
	Response *ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevicesResponse struct {
	Type                     string `json:"type,omitempty"`                     // Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
	Feature                  string `json:"feature,omitempty"`                  // Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
	ContractLevel            string `json:"contractLevel,omitempty"`            // Contract level for which trial was created. this was used in older versions and exists only for compatibility.
	Active                   *bool  `json:"active,omitempty"`                   // Indicates if the trial is active
	StartTime                *int   `json:"startTime,omitempty"`                // Trial start time; as measured in Unix epoch time in milliseconds
	EndTime                  *int   `json:"endTime,omitempty"`                  // Trial end time; as measured in Unix epoch time in milliseconds
	SecondsRemainingToExpiry *int   `json:"secondsRemainingToExpiry,omitempty"` // Seconds remaining in the trial before it expires. for expired trials this will be 0.
	SecondsSinceExpired      *int   `json:"secondsSinceExpired,omitempty"`      // Seconds elapsed after the trial has expired. for active trials this will be 0.
}
type ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices struct {
	Version  string                                                                                 `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevicesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesResponse `json:"response,omitempty"` //
}
type ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type RequestComplianceRunCompliance struct {
	TriggerFull *bool    `json:"triggerFull,omitempty"` // if it is true then compliance will be triggered for all categories. If it is false then compliance will be triggered for categories mentioned in categories section .
	Categories  []string `json:"categories,omitempty"`  // Category can have any value among 'INTENT'(mapped to compliance types: NETWORK_SETTINGS,NETWORK_PROFILE,WORKFLOW,FABRIC,APPLICATION_VISIBILITY), 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'EOX'
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` // UUID of the device.
}
type RequestComplianceCommitDeviceConfiguration struct {
	DeviceID []string `json:"deviceId,omitempty"` // UUID of the device.
}

//GetComplianceStatus Get Compliance Status  - dda5-cb9a-49aa-aef6
/* Return compliance status of device(s).


@param GetComplianceStatusQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-status
*/
func (s *ComplianceService) GetComplianceStatus(GetComplianceStatusQueryParams *GetComplianceStatusQueryParams) (*ResponseComplianceGetComplianceStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance"

	queryString, _ := query.Values(GetComplianceStatusQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceStatus(GetComplianceStatusQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceStatus")
	}

	result := response.Result().(*ResponseComplianceGetComplianceStatus)
	return result, response, err

}

//GetComplianceStatusCount Get Compliance Status Count - db99-f919-424a-9f83
/* Return Compliance Status Count


@param GetComplianceStatusCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-status-count
*/
func (s *ComplianceService) GetComplianceStatusCount(GetComplianceStatusCountQueryParams *GetComplianceStatusCountQueryParams) (*ResponseComplianceGetComplianceStatusCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/count"

	queryString, _ := query.Values(GetComplianceStatusCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceStatusCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceStatusCount(GetComplianceStatusCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceStatusCount")
	}

	result := response.Result().(*ResponseComplianceGetComplianceStatusCount)
	return result, response, err

}

//GetComplianceDetail Get Compliance Detail  - dda1-1ae7-4788-9d49
/* Return Compliance Detail


@param GetComplianceDetailQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-detail
*/
func (s *ComplianceService) GetComplianceDetail(GetComplianceDetailQueryParams *GetComplianceDetailQueryParams) (*ResponseComplianceGetComplianceDetail, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/detail"

	queryString, _ := query.Values(GetComplianceDetailQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceDetail{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceDetail(GetComplianceDetailQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceDetail")
	}

	result := response.Result().(*ResponseComplianceGetComplianceDetail)
	return result, response, err

}

//GetComplianceDetailCount Get Compliance Detail Count - 3eb6-58c3-4549-94df
/* Return  Compliance Count Detail


@param GetComplianceDetailCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-compliance-detail-count
*/
func (s *ComplianceService) GetComplianceDetailCount(GetComplianceDetailCountQueryParams *GetComplianceDetailCountQueryParams) (*ResponseComplianceGetComplianceDetailCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/detail/count"

	queryString, _ := query.Values(GetComplianceDetailCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetComplianceDetailCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetComplianceDetailCount(GetComplianceDetailCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetComplianceDetailCount")
	}

	result := response.Result().(*ResponseComplianceGetComplianceDetailCount)
	return result, response, err

}

//DeviceComplianceStatus Device Compliance Status - 7aa8-5ad5-48ea-94a7
/* Return compliance status of a device.


@param deviceUUID deviceUuid path parameter. Device Id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!device-compliance-status
*/
func (s *ComplianceService) DeviceComplianceStatus(deviceUUID string) (*ResponseComplianceDeviceComplianceStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/{deviceUuid}"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceDeviceComplianceStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeviceComplianceStatus(deviceUUID)
		}
		return nil, response, fmt.Errorf("error with operation DeviceComplianceStatus")
	}

	result := response.Result().(*ResponseComplianceDeviceComplianceStatus)
	return result, response, err

}

//ComplianceDetailsOfDevice Compliance Details of Device - 52bf-e904-45aa-b017
/* Return compliance detailed report for a device.


@param deviceUUID deviceUuid path parameter. Device Id

@param ComplianceDetailsOfDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!compliance-details-of-device
*/
func (s *ComplianceService) ComplianceDetailsOfDevice(deviceUUID string, ComplianceDetailsOfDeviceQueryParams *ComplianceDetailsOfDeviceQueryParams) (*ResponseComplianceComplianceDetailsOfDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/{deviceUuid}/detail"
	path = strings.Replace(path, "{deviceUuid}", fmt.Sprintf("%v", deviceUUID), -1)

	queryString, _ := query.Values(ComplianceDetailsOfDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceComplianceDetailsOfDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ComplianceDetailsOfDevice(deviceUUID, ComplianceDetailsOfDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ComplianceDetailsOfDevice")
	}

	result := response.Result().(*ResponseComplianceComplianceDetailsOfDevice)
	return result, response, err

}

//GetFieldNoticeNetworkDevices Get field notice network devices - e8b3-68d9-483b-8e07
/* Get field notice network devices


@param GetFieldNoticeNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-devices
*/
func (s *ComplianceService) GetFieldNoticeNetworkDevices(GetFieldNoticeNetworkDevicesQueryParams *GetFieldNoticeNetworkDevicesQueryParams) (*ResponseComplianceGetFieldNoticeNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices"

	queryString, _ := query.Values(GetFieldNoticeNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticeNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDevices(GetFieldNoticeNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDevices)
	return result, response, err

}

//GetCountOfFieldNoticeNetworkDevices Get count of field notice network devices - 23bd-3911-4cc9-987c
/* Get count of field notice network devices


@param GetCountOfFieldNoticeNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notice-network-devices
*/
func (s *ComplianceService) GetCountOfFieldNoticeNetworkDevices(GetCountOfFieldNoticeNetworkDevicesQueryParams *GetCountOfFieldNoticeNetworkDevicesQueryParams) (*ResponseComplianceGetCountOfFieldNoticeNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/count"

	queryString, _ := query.Values(GetCountOfFieldNoticeNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticeNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticeNetworkDevices(GetCountOfFieldNoticeNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticeNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticeNetworkDevices)
	return result, response, err

}

//GetFieldNoticeNetworkDeviceByDeviceID Get field notice network device by device id - db80-68db-4b1b-976d
/* Get field notice network device by device id


@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-device-by-device-id
*/
func (s *ComplianceService) GetFieldNoticeNetworkDeviceByDeviceID(networkDeviceID string) (*ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDeviceByDeviceID(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDeviceByDeviceId")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDeviceByDeviceID)
	return result, response, err

}

//GetFieldNoticesAffectingTheNetworkDevice Get field notices affecting the network device - e5a6-3887-44c9-95d6
/* Get field notices affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetFieldNoticesAffectingTheNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notices-affecting-the-network-device
*/
func (s *ComplianceService) GetFieldNoticesAffectingTheNetworkDevice(networkDeviceID string, GetFieldNoticesAffectingTheNetworkDeviceQueryParams *GetFieldNoticesAffectingTheNetworkDeviceQueryParams) (*ResponseComplianceGetFieldNoticesAffectingTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}/notices"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetFieldNoticesAffectingTheNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticesAffectingTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticesAffectingTheNetworkDevice(networkDeviceID, GetFieldNoticesAffectingTheNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticesAffectingTheNetworkDevice")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticesAffectingTheNetworkDevice)
	return result, response, err

}

//GetCountOfFieldNoticesAffectingTheNetworkDevice Get count of field notices affecting the network device - 5494-098c-414b-bf34
/* Get count of field notices affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notices-affecting-the-network-device
*/
func (s *ComplianceService) GetCountOfFieldNoticesAffectingTheNetworkDevice(networkDeviceID string, GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams *GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams) (*ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}/notices/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticesAffectingTheNetworkDevice(networkDeviceID, GetCountOfFieldNoticesAffectingTheNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticesAffectingTheNetworkDevice")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticesAffectingTheNetworkDevice)
	return result, response, err

}

//GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID Get field notice affecting the network device by device Id and notice id - 86a9-2ad8-436a-9299
/* Get field notice affecting the network device by device Id and notice id


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param id id path parameter. Id of the field notice


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-affecting-the-network-device-by-device-id-and-notice-id
*/
func (s *ComplianceService) GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID(networkDeviceID string, id string) (*ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/networkDevices/{networkDeviceId}/notices/{id}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeAffectingTheNetworkDeviceByDeviceIdAndNoticeId")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeAffectingTheNetworkDeviceByDeviceIDAndNoticeID)
	return result, response, err

}

//GetFieldNotices Get field notices - 6989-39f3-4279-ae61
/* Get field notices


@param GetFieldNoticesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notices
*/
func (s *ComplianceService) GetFieldNotices(GetFieldNoticesQueryParams *GetFieldNoticesQueryParams) (*ResponseComplianceGetFieldNotices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices"

	queryString, _ := query.Values(GetFieldNoticesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNotices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNotices(GetFieldNoticesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNotices")
	}

	result := response.Result().(*ResponseComplianceGetFieldNotices)
	return result, response, err

}

//GetCountOfFieldNotices Get count of field notices - ba99-e9ba-40cb-99e1
/* Get count of field notices


@param GetCountOfFieldNoticesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notices
*/
func (s *ComplianceService) GetCountOfFieldNotices(GetCountOfFieldNoticesQueryParams *GetCountOfFieldNoticesQueryParams) (*ResponseComplianceGetCountOfFieldNotices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/count"

	queryString, _ := query.Values(GetCountOfFieldNoticesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNotices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNotices(GetCountOfFieldNoticesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNotices")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNotices)
	return result, response, err

}

//GetFieldNoticeByID Get field notice by Id - 7c90-9909-4a19-8d77
/* Get field notice by Id


@param id id path parameter. Id of the field notice


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-by-id
*/
func (s *ComplianceService) GetFieldNoticeByID(id string) (*ResponseComplianceGetFieldNoticeByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeById")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeByID)
	return result, response, err

}

//GetFieldNoticeNetworkDevicesForTheNotice Get field notice network devices for the notice - ddaa-9b91-4cfa-8943
/* Get field notice network devices for the notice


@param id id path parameter. Id of the field notice

@param GetFieldNoticeNetworkDevicesForTheNoticeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-devices-for-the-notice
*/
func (s *ComplianceService) GetFieldNoticeNetworkDevicesForTheNotice(id string, GetFieldNoticeNetworkDevicesForTheNoticeQueryParams *GetFieldNoticeNetworkDevicesForTheNoticeQueryParams) (*ResponseComplianceGetFieldNoticeNetworkDevicesForTheNotice, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}/networkDevices"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetFieldNoticeNetworkDevicesForTheNoticeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticeNetworkDevicesForTheNotice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDevicesForTheNotice(id, GetFieldNoticeNetworkDevicesForTheNoticeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDevicesForTheNotice")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDevicesForTheNotice)
	return result, response, err

}

//GetCountOfFieldNoticeNetworkDevicesForTheNotice Get count of field notice network devices for the notice - 4a9e-d86a-421b-890a
/* Get count of field notice network devices for the notice


@param id id path parameter. Id of the field notice

@param GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notice-network-devices-for-the-notice
*/
func (s *ComplianceService) GetCountOfFieldNoticeNetworkDevicesForTheNotice(id string, GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams *GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams) (*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNotice, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}/networkDevices/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNotice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticeNetworkDevicesForTheNotice(id, GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticeNetworkDevicesForTheNotice")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNotice)
	return result, response, err

}

//GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID Get field notice network device for the notice by network device id - e4b4-cb51-46cb-9925
/* Get field notice network device for the notice by network device id


@param id id path parameter. Id of the field notice

@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notice-network-device-for-the-notice-by-network-device-id
*/
func (s *ComplianceService) GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID(id string, networkDeviceID string) (*ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/results/notices/{id}/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID(id, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceId")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticeNetworkDeviceForTheNoticeByNetworkDeviceID)
	return result, response, err

}

//GetFieldNoticesResultsTrendOverTime Get field notices results trend over time - 6690-8bd1-4b58-ba8e
/* Get field notices results trend over time. The default sort is by scan time descending.


@param GetFieldNoticesResultsTrendOverTimeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-field-notices-results-trend-over-time
*/
func (s *ComplianceService) GetFieldNoticesResultsTrendOverTime(GetFieldNoticesResultsTrendOverTimeQueryParams *GetFieldNoticesResultsTrendOverTimeQueryParams) (*ResponseComplianceGetFieldNoticesResultsTrendOverTime, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/resultsTrend"

	queryString, _ := query.Values(GetFieldNoticesResultsTrendOverTimeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetFieldNoticesResultsTrendOverTime{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetFieldNoticesResultsTrendOverTime(GetFieldNoticesResultsTrendOverTimeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetFieldNoticesResultsTrendOverTime")
	}

	result := response.Result().(*ResponseComplianceGetFieldNoticesResultsTrendOverTime)
	return result, response, err

}

//GetCountOfFieldNoticesResultsTrendOverTime Get count of field notices results trend over time - d285-c901-46b9-a120
/* Get count of field notices results trend over time


@param GetCountOfFieldNoticesResultsTrendOverTimeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-field-notices-results-trend-over-time
*/
func (s *ComplianceService) GetCountOfFieldNoticesResultsTrendOverTime(GetCountOfFieldNoticesResultsTrendOverTimeQueryParams *GetCountOfFieldNoticesResultsTrendOverTimeQueryParams) (*ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTime, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/resultsTrend/count"

	queryString, _ := query.Values(GetCountOfFieldNoticesResultsTrendOverTimeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTime{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfFieldNoticesResultsTrendOverTime(GetCountOfFieldNoticesResultsTrendOverTimeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfFieldNoticesResultsTrendOverTime")
	}

	result := response.Result().(*ResponseComplianceGetCountOfFieldNoticesResultsTrendOverTime)
	return result, response, err

}

//GetTrialDetailsForFieldNoticesDetectionOnNetworkDevices Get trial details for field notices detection on network devices - 92b9-d9a8-4a09-ad05
/* Get trial details for field notices detection on network devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trial-details-for-field-notices-detection-on-network-devices
*/
func (s *ComplianceService) GetTrialDetailsForFieldNoticesDetectionOnNetworkDevices() (*ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrialDetailsForFieldNoticesDetectionOnNetworkDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetTrialDetailsForFieldNoticesDetectionOnNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevices)
	return result, response, err

}

//GetConfigTaskDetails Get config task details - 8183-1a90-4788-b8c5
/* Returns a config task result details by specified id


@param GetConfigTaskDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-config-task-details
*/
func (s *ComplianceService) GetConfigTaskDetails(GetConfigTaskDetailsQueryParams *GetConfigTaskDetailsQueryParams) (*ResponseComplianceGetConfigTaskDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-config/task"

	queryString, _ := query.Values(GetConfigTaskDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetConfigTaskDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigTaskDetails(GetConfigTaskDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigTaskDetails")
	}

	result := response.Result().(*ResponseComplianceGetConfigTaskDetails)
	return result, response, err

}

//GetNetworkBugs Get network bugs - 3e8a-0a51-423a-968f
/* Get network bugs


@param GetNetworkBugsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bugs
*/
func (s *ComplianceService) GetNetworkBugs(GetNetworkBugsQueryParams *GetNetworkBugsQueryParams) (*ResponseComplianceGetNetworkBugs, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs"

	queryString, _ := query.Values(GetNetworkBugsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugs(GetNetworkBugsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugs")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugs)
	return result, response, err

}

//GetCountOfNetworkBugs Get count of network bugs - 3cad-684d-4508-aa15
/* Get count of network bugs


@param GetCountOfNetworkBugsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bugs
*/
func (s *ComplianceService) GetCountOfNetworkBugs(GetCountOfNetworkBugsQueryParams *GetCountOfNetworkBugsQueryParams) (*ResponseComplianceGetCountOfNetworkBugs, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/count"

	queryString, _ := query.Values(GetCountOfNetworkBugsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugs{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugs(GetCountOfNetworkBugsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugs")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugs)
	return result, response, err

}

//GetNetworkBugByID Get network bug by Id - ec93-a9c6-48d9-9050
/* Get network bug by Id


@param id id path parameter. Id of the network bug


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-by-id
*/
func (s *ComplianceService) GetNetworkBugByID(id string) (*ResponseComplianceGetNetworkBugByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetNetworkBugByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugById")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugByID)
	return result, response, err

}

//GetNetworkBugDevicesForTheBug Get network bug devices for the bug - a18c-2be4-4a1b-bbf7
/* Get network bug devices for the bug


@param id id path parameter. Id of the network bug

@param GetNetworkBugDevicesForTheBugQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-devices-for-the-bug
*/
func (s *ComplianceService) GetNetworkBugDevicesForTheBug(id string, GetNetworkBugDevicesForTheBugQueryParams *GetNetworkBugDevicesForTheBugQueryParams) (*ResponseComplianceGetNetworkBugDevicesForTheBug, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}/networkDevices"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetNetworkBugDevicesForTheBugQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugDevicesForTheBug{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDevicesForTheBug(id, GetNetworkBugDevicesForTheBugQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDevicesForTheBug")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDevicesForTheBug)
	return result, response, err

}

//GetCountOfNetworkBugDevicesForTheBug Get count of network bug devices for the bug - 269a-d906-4e5b-aad1
/* Get count of network bug devices for the bug


@param id id path parameter. Id of the network bug

@param GetCountOfNetworkBugDevicesForTheBugQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bug-devices-for-the-bug
*/
func (s *ComplianceService) GetCountOfNetworkBugDevicesForTheBug(id string, GetCountOfNetworkBugDevicesForTheBugQueryParams *GetCountOfNetworkBugDevicesForTheBugQueryParams) (*ResponseComplianceGetCountOfNetworkBugDevicesForTheBug, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}/networkDevices/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetCountOfNetworkBugDevicesForTheBugQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugDevicesForTheBug{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugDevicesForTheBug(id, GetCountOfNetworkBugDevicesForTheBugQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugDevicesForTheBug")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugDevicesForTheBug)
	return result, response, err

}

//GetNetworkBugDeviceForTheBugByNetworkDeviceID Get network bug device for the bug by network device id - 7594-39b9-4d78-8144
/* Get network bug device for the bug by network device id


@param id id path parameter. Id of the network bug

@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-device-for-the-bug-by-network-device-id
*/
func (s *ComplianceService) GetNetworkBugDeviceForTheBugByNetworkDeviceID(id string, networkDeviceID string) (*ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/bugs/{id}/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDeviceForTheBugByNetworkDeviceID(id, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDeviceForTheBugByNetworkDeviceId")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDeviceForTheBugByNetworkDeviceID)
	return result, response, err

}

//GetNetworkBugDevices Get network bug devices - f9ad-f991-4c0a-9248
/* Get network bug devices


@param GetNetworkBugDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-devices
*/
func (s *ComplianceService) GetNetworkBugDevices(GetNetworkBugDevicesQueryParams *GetNetworkBugDevicesQueryParams) (*ResponseComplianceGetNetworkBugDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices"

	queryString, _ := query.Values(GetNetworkBugDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDevices(GetNetworkBugDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDevices")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDevices)
	return result, response, err

}

//GetCountOfNetworkBugDevices Get count of network bug devices - e583-ba39-4a4b-b8bd
/* Get count of network bug devices


@param GetCountOfNetworkBugDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bug-devices
*/
func (s *ComplianceService) GetCountOfNetworkBugDevices(GetCountOfNetworkBugDevicesQueryParams *GetCountOfNetworkBugDevicesQueryParams) (*ResponseComplianceGetCountOfNetworkBugDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/count"

	queryString, _ := query.Values(GetCountOfNetworkBugDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugDevices(GetCountOfNetworkBugDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugDevices")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugDevices)
	return result, response, err

}

//GetNetworkBugDeviceByDeviceID Get network bug device by device id - 3eae-5b2d-43eb-b61f
/* Get network bug device by device id


@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bug-device-by-device-id
*/
func (s *ComplianceService) GetNetworkBugDeviceByDeviceID(networkDeviceID string) (*ResponseComplianceGetNetworkBugDeviceByDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetNetworkBugDeviceByDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugDeviceByDeviceID(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugDeviceByDeviceId")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugDeviceByDeviceID)
	return result, response, err

}

//GetBugsAffectingTheNetworkDevice Get bugs affecting the network device - 199a-c80d-4138-8853
/* Get bugs affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetBugsAffectingTheNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-bugs-affecting-the-network-device
*/
func (s *ComplianceService) GetBugsAffectingTheNetworkDevice(networkDeviceID string, GetBugsAffectingTheNetworkDeviceQueryParams *GetBugsAffectingTheNetworkDeviceQueryParams) (*ResponseComplianceGetBugsAffectingTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}/bugs"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetBugsAffectingTheNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetBugsAffectingTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBugsAffectingTheNetworkDevice(networkDeviceID, GetBugsAffectingTheNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetBugsAffectingTheNetworkDevice")
	}

	result := response.Result().(*ResponseComplianceGetBugsAffectingTheNetworkDevice)
	return result, response, err

}

//GetCountOfBugsAffectingTheNetworkDevice Get count of bugs affecting the network device - 86a4-d898-466a-a7e0
/* Get count of bugs affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetCountOfBugsAffectingTheNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-bugs-affecting-the-network-device
*/
func (s *ComplianceService) GetCountOfBugsAffectingTheNetworkDevice(networkDeviceID string, GetCountOfBugsAffectingTheNetworkDeviceQueryParams *GetCountOfBugsAffectingTheNetworkDeviceQueryParams) (*ResponseComplianceGetCountOfBugsAffectingTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}/bugs/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetCountOfBugsAffectingTheNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfBugsAffectingTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfBugsAffectingTheNetworkDevice(networkDeviceID, GetCountOfBugsAffectingTheNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfBugsAffectingTheNetworkDevice")
	}

	result := response.Result().(*ResponseComplianceGetCountOfBugsAffectingTheNetworkDevice)
	return result, response, err

}

//GetBugAffectingTheNetworkDeviceByDeviceIDAndBugID Get bug affecting the network device by device Id and bug id - e293-6b08-49c9-b715
/* Get bug affecting the network device by device Id and bug id


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param id id path parameter. Id of the network bug


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-bug-affecting-the-network-device-by-device-id-and-bug-id
*/
func (s *ComplianceService) GetBugAffectingTheNetworkDeviceByDeviceIDAndBugID(networkDeviceID string, id string) (*ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugID, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/results/networkDevices/{networkDeviceId}/bugs/{id}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBugAffectingTheNetworkDeviceByDeviceIDAndBugID(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetBugAffectingTheNetworkDeviceByDeviceIdAndBugId")
	}

	result := response.Result().(*ResponseComplianceGetBugAffectingTheNetworkDeviceByDeviceIDAndBugID)
	return result, response, err

}

//GetNetworkBugsResultsTrendOverTime Get network bugs results trend over time - 708d-2bc8-42da-b062
/* Get network bugs results trend over time. The default sort is by scan time descending.


@param GetNetworkBugsResultsTrendOverTimeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-bugs-results-trend-over-time
*/
func (s *ComplianceService) GetNetworkBugsResultsTrendOverTime(GetNetworkBugsResultsTrendOverTimeQueryParams *GetNetworkBugsResultsTrendOverTimeQueryParams) (*ResponseComplianceGetNetworkBugsResultsTrendOverTime, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/resultsTrend"

	queryString, _ := query.Values(GetNetworkBugsResultsTrendOverTimeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetNetworkBugsResultsTrendOverTime{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkBugsResultsTrendOverTime(GetNetworkBugsResultsTrendOverTimeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkBugsResultsTrendOverTime")
	}

	result := response.Result().(*ResponseComplianceGetNetworkBugsResultsTrendOverTime)
	return result, response, err

}

//GetCountOfNetworkBugsResultsTrendOverTime Get count of network bugs results trend over time - 6791-696c-4199-86e1
/* Get count of network bugs results trend over time


@param GetCountOfNetworkBugsResultsTrendOverTimeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-network-bugs-results-trend-over-time
*/
func (s *ComplianceService) GetCountOfNetworkBugsResultsTrendOverTime(GetCountOfNetworkBugsResultsTrendOverTimeQueryParams *GetCountOfNetworkBugsResultsTrendOverTimeQueryParams) (*ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTime, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/resultsTrend/count"

	queryString, _ := query.Values(GetCountOfNetworkBugsResultsTrendOverTimeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTime{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfNetworkBugsResultsTrendOverTime(GetCountOfNetworkBugsResultsTrendOverTimeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfNetworkBugsResultsTrendOverTime")
	}

	result := response.Result().(*ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTime)
	return result, response, err

}

//GetTrialDetailsForBugsDetectionOnNetworkDevices Get trial details for bugs detection on network devices - 11a4-a89b-430b-93cd
/* Get trial details for bugs detection on network devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trial-details-for-bugs-detection-on-network-devices
*/
func (s *ComplianceService) GetTrialDetailsForBugsDetectionOnNetworkDevices() (*ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrialDetailsForBugsDetectionOnNetworkDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetTrialDetailsForBugsDetectionOnNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetTrialDetailsForBugsDetectionOnNetworkDevices)
	return result, response, err

}

//GetSecurityAdvisoriesAffectingTheNetworkDevices Get security advisories affecting the network devices - ef91-f8be-47d8-8fbf
/* Get security advisories affecting the network devices


@param GetSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisories-affecting-the-network-devices
*/
func (s *ComplianceService) GetSecurityAdvisoriesAffectingTheNetworkDevices(GetSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams *GetSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams) (*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories"

	queryString, _ := query.Values(GetSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoriesAffectingTheNetworkDevices(GetSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoriesAffectingTheNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevices)
	return result, response, err

}

//GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices Get count of security advisories affecting the network devices - 129c-9b1f-4dd8-9173
/* Get count of security advisories affecting the network devices


@param GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisories-affecting-the-network-devices
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams *GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/count"

	queryString, _ := query.Values(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices(GetCountOfSecurityAdvisoriesAffectingTheNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoriesAffectingTheNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevices)
	return result, response, err

}

//GetSecurityAdvisoryAffectingTheNetworkDevicesByID Get security advisory affecting the network devices by Id - 51aa-ea19-4c88-bea6
/* Get security advisory affecting the network devices by Id


@param id id path parameter. Id of the security advisory


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-affecting-the-network-devices-by-id
*/
func (s *ComplianceService) GetSecurityAdvisoryAffectingTheNetworkDevicesByID(id string) (*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryAffectingTheNetworkDevicesByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryAffectingTheNetworkDevicesById")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDevicesByID)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory Get security advisory network devices for the security advisory - ee81-e9ad-40bb-b3d1
/* Get security advisory network devices for the security advisory


@param id id path parameter. Id of the security advisory

@param GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-devices-for-the-security-advisory
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(id string, GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams *GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams) (*ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}/networkDevices"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(id, GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory)
	return result, response, err

}

//GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory Get count of security advisory network devices for the security advisory - 969b-bb96-404b-b905
/* Get count of security advisory network devices for the security advisory


@param id id path parameter. Id of the security advisory

@param GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisory-network-devices-for-the-security-advisory
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(id string, GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams *GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}/networkDevices/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory(id, GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisoryQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevicesForTheSecurityAdvisory)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID Get security advisory network device for the security advisory by network device id - 15ac-59b6-4668-a848
/* Get security advisory network device for the security advisory by network device id


@param id id path parameter. Id of the security advisory

@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-device-for-the-security-advisory-by-network-device-id
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID(id string, networkDeviceID string) (*ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/advisories/{id}/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID(id, networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceId")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDeviceForTheSecurityAdvisoryByNetworkDeviceID)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDevices Get security advisory network devices - af83-89a1-43da-9337
/* Get security advisory network devices


@param GetSecurityAdvisoryNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-devices
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDevices(GetSecurityAdvisoryNetworkDevicesQueryParams *GetSecurityAdvisoryNetworkDevicesQueryParams) (*ResponseComplianceGetSecurityAdvisoryNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices"

	queryString, _ := query.Values(GetSecurityAdvisoryNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDevices(GetSecurityAdvisoryNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDevices)
	return result, response, err

}

//GetCountOfSecurityAdvisoryNetworkDevices Get count of security advisory network devices - 93a6-8af1-438a-8f39
/* Get count of security advisory network devices


@param GetCountOfSecurityAdvisoryNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisory-network-devices
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoryNetworkDevices(GetCountOfSecurityAdvisoryNetworkDevicesQueryParams *GetCountOfSecurityAdvisoryNetworkDevicesQueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/count"

	queryString, _ := query.Values(GetCountOfSecurityAdvisoryNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoryNetworkDevices(GetCountOfSecurityAdvisoryNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoryNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoryNetworkDevices)
	return result, response, err

}

//GetSecurityAdvisoryNetworkDeviceByNetworkDeviceID Get security advisory network device by network device id - a5bb-ca1a-4abb-8a7f
/* Get security advisory network device by network device id


@param networkDeviceID networkDeviceId path parameter. Id of the network device


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-network-device-by-network-device-id
*/
func (s *ComplianceService) GetSecurityAdvisoryNetworkDeviceByNetworkDeviceID(networkDeviceID string) (*ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceID, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryNetworkDeviceByNetworkDeviceID(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryNetworkDeviceByNetworkDeviceId")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryNetworkDeviceByNetworkDeviceID)
	return result, response, err

}

//GetSecurityAdvisoriesAffectingTheNetworkDevice Get security advisories affecting the network device - 20a9-3b0d-4769-8091
/* Get security advisories affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisories-affecting-the-network-device
*/
func (s *ComplianceService) GetSecurityAdvisoriesAffectingTheNetworkDevice(networkDeviceID string, GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams *GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams) (*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}/advisories"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoriesAffectingTheNetworkDevice(networkDeviceID, GetSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoriesAffectingTheNetworkDevice")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoriesAffectingTheNetworkDevice)
	return result, response, err

}

//GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice Get count of security advisories affecting the network device - d4ba-db3a-4488-9d47
/* Get count of security advisories affecting the network device


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisories-affecting-the-network-device
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice(networkDeviceID string, GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams *GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}/advisories/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice(networkDeviceID, GetCountOfSecurityAdvisoriesAffectingTheNetworkDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoriesAffectingTheNetworkDevice")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoriesAffectingTheNetworkDevice)
	return result, response, err

}

//GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID Get security advisory affecting the network device by device Id and advisory id - c9a3-d93e-4fe8-959d
/* Get security advisory affecting the network device by device Id and advisory id


@param networkDeviceID networkDeviceId path parameter. Id of the network device

@param id id path parameter. Id of the security advisory


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisory-affecting-the-network-device-by-device-id-and-advisory-id
*/
func (s *ComplianceService) GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID(networkDeviceID string, id string) (*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/results/networkDevices/{networkDeviceId}/advisories/{id}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIdAndAdvisoryId")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoryAffectingTheNetworkDeviceByDeviceIDAndAdvisoryID)
	return result, response, err

}

//GetSecurityAdvisoriesResultsTrendOverTime Get security advisories results trend over time - b584-aa2b-4158-bc5a
/* Get security advisories results trend over time. The default sort is by scan time descending.


@param GetSecurityAdvisoriesResultsTrendOverTimeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-security-advisories-results-trend-over-time
*/
func (s *ComplianceService) GetSecurityAdvisoriesResultsTrendOverTime(GetSecurityAdvisoriesResultsTrendOverTimeQueryParams *GetSecurityAdvisoriesResultsTrendOverTimeQueryParams) (*ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTime, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/resultsTrend"

	queryString, _ := query.Values(GetSecurityAdvisoriesResultsTrendOverTimeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTime{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSecurityAdvisoriesResultsTrendOverTime(GetSecurityAdvisoriesResultsTrendOverTimeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSecurityAdvisoriesResultsTrendOverTime")
	}

	result := response.Result().(*ResponseComplianceGetSecurityAdvisoriesResultsTrendOverTime)
	return result, response, err

}

//GetCountOfSecurityAdvisoriesResultsTrendOverTime Get count of security advisories results trend over time - a9af-78ef-46aa-8534
/* Get count of security advisories results trend over time.


@param GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-security-advisories-results-trend-over-time
*/
func (s *ComplianceService) GetCountOfSecurityAdvisoriesResultsTrendOverTime(GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams *GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams) (*ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTime, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/resultsTrend/count"

	queryString, _ := query.Values(GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTime{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfSecurityAdvisoriesResultsTrendOverTime(GetCountOfSecurityAdvisoriesResultsTrendOverTimeQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfSecurityAdvisoriesResultsTrendOverTime")
	}

	result := response.Result().(*ResponseComplianceGetCountOfSecurityAdvisoriesResultsTrendOverTime)
	return result, response, err

}

//GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices Get trial details for security advisories detection on network devices - f6ba-8a34-4c4a-ba48
/* Get trial details for security advisories detection on network devices



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-trial-details-for-security-advisories-detection-on-network-devices
*/
func (s *ComplianceService) GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices() (*ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices()
		}
		return nil, response, fmt.Errorf("error with operation GetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceGetTrialDetailsForSecurityAdvisoriesDetectionOnNetworkDevices)
	return result, response, err

}

//RunCompliance Run Compliance - f6ae-c8a7-4428-a9ff
/* Run compliance check for device(s).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!run-compliance
*/
func (s *ComplianceService) RunCompliance(requestComplianceRunCompliance *RequestComplianceRunCompliance) (*ResponseComplianceRunCompliance, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestComplianceRunCompliance).
		SetResult(&ResponseComplianceRunCompliance{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RunCompliance(requestComplianceRunCompliance)
		}

		return nil, response, fmt.Errorf("error with operation RunCompliance")
	}

	result := response.Result().(*ResponseComplianceRunCompliance)
	return result, response, err

}

//ComplianceRemediation Compliance Remediation - 7d80-2867-4179-8488
/* Remediates configuration compliance issues. Compliance issues related to 'Routing', 'HA Remediation', 'Software Image', 'Securities Advisories', 'SD-Access Unsupported Configuration', 'Workflow', etc. will not be addressed by this API.
Warning: Fixing compliance mismatches could result in a possible network flap.


@param id id path parameter. Network device identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!compliance-remediation
*/
func (s *ComplianceService) ComplianceRemediation(id string) (*ResponseComplianceComplianceRemediation, *resty.Response, error) {
	path := "/dna/intent/api/v1/compliance/networkDevices/{id}/issues/remediation/provision"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceComplianceRemediation{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ComplianceRemediation(id)
		}

		return nil, response, fmt.Errorf("error with operation ComplianceRemediation")
	}

	result := response.Result().(*ResponseComplianceComplianceRemediation)
	return result, response, err

}

//CreatesATrialForFieldNoticesDetectionOnNetworkDevices Creates a trial for field notices detection on network devices - 3a9a-88e2-4c3a-9db8
/* Creates a trial for field notices detection on network devices. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-trial-for-field-notices-detection-on-network-devices
*/
func (s *ComplianceService) CreatesATrialForFieldNoticesDetectionOnNetworkDevices() (*ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesATrialForFieldNoticesDetectionOnNetworkDevices()
		}

		return nil, response, fmt.Errorf("error with operation CreatesATrialForFieldNoticesDetectionOnNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceCreatesATrialForFieldNoticesDetectionOnNetworkDevices)
	return result, response, err

}

//TriggersAFieldNoticesScanForTheSupportedNetworkDevices Triggers a field notices scan for the supported network devices - d4b4-5ae2-4e68-bb04
/* Triggers a field notices scan for the supported network devices. The supported devices are switches, routers and wireless controllers. If a device is not supported, the FieldNoticeNetworkDevice scanStatus will be Failed with appropriate comments. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.


@param TriggersAFieldNoticesScanForTheSupportedNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!triggers-a-field-notices-scan-for-the-supported-network-devices
*/
func (s *ComplianceService) TriggersAFieldNoticesScanForTheSupportedNetworkDevices(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesQueryParams *TriggersAFieldNoticesScanForTheSupportedNetworkDevicesQueryParams) (*ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/fieldNotices/triggerScan"

	queryString, _ := query.Values(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggersAFieldNoticesScanForTheSupportedNetworkDevices(TriggersAFieldNoticesScanForTheSupportedNetworkDevicesQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggersAFieldNoticesScanForTheSupportedNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceTriggersAFieldNoticesScanForTheSupportedNetworkDevices)
	return result, response, err

}

//CommitDeviceConfiguration Commit device configuration - 53a3-5a70-4e3b-87b5
/* This operation would commit device running configuration to startup by issuing "write memory" to device



Documentation Link: https://developer.cisco.com/docs/dna-center/#!commit-device-configuration
*/
func (s *ComplianceService) CommitDeviceConfiguration(requestComplianceCommitDeviceConfiguration *RequestComplianceCommitDeviceConfiguration) (*ResponseComplianceCommitDeviceConfiguration, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-config/write-memory"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestComplianceCommitDeviceConfiguration).
		SetResult(&ResponseComplianceCommitDeviceConfiguration{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CommitDeviceConfiguration(requestComplianceCommitDeviceConfiguration)
		}

		return nil, response, fmt.Errorf("error with operation CommitDeviceConfiguration")
	}

	result := response.Result().(*ResponseComplianceCommitDeviceConfiguration)
	return result, response, err

}

//CreatesATrialForBugsDetectionOnNetworkDevices Creates a trial for bugs detection on network devices - b080-6bcf-402b-ad8e
/* Creates a trial for bugs detection on network devices. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-trial-for-bugs-detection-on-network-devices
*/
func (s *ComplianceService) CreatesATrialForBugsDetectionOnNetworkDevices() (*ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesATrialForBugsDetectionOnNetworkDevices()
		}

		return nil, response, fmt.Errorf("error with operation CreatesATrialForBugsDetectionOnNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceCreatesATrialForBugsDetectionOnNetworkDevices)
	return result, response, err

}

//TriggersABugsScanForTheSupportedNetworkDevices Triggers a bugs scan for the supported network devices - 5296-db34-457b-b233
/* Triggers a bugs scan for the supported network devices. The supported devices are switches and routers. If a device is not supported, the NetworkBugsDevice scanStatus will be Failed with appropriate comments.


@param TriggersABugsScanForTheSupportedNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!triggers-a-bugs-scan-for-the-supported-network-devices
*/
func (s *ComplianceService) TriggersABugsScanForTheSupportedNetworkDevices(TriggersABugsScanForTheSupportedNetworkDevicesQueryParams *TriggersABugsScanForTheSupportedNetworkDevicesQueryParams) (*ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkBugs/triggerScan"

	queryString, _ := query.Values(TriggersABugsScanForTheSupportedNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggersABugsScanForTheSupportedNetworkDevices(TriggersABugsScanForTheSupportedNetworkDevicesQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggersABugsScanForTheSupportedNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevices)
	return result, response, err

}

//CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices Creates a trial for security advisories detection on network devices - 0190-1b7a-4edb-91d8
/* Creates a trial for security advisories detection on network devices. The consent to connect agreement must have been accepted in the UI for this to succeed. Please refer to the user guide at
 for more details on consent to connect.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-trial-for-security-advisories-detection-on-network-devices
*/
func (s *ComplianceService) CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices() (*ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/trials"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices()
		}

		return nil, response, fmt.Errorf("error with operation CreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceCreatesATrialForSecurityAdvisoriesDetectionOnNetworkDevices)
	return result, response, err

}

//TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices Triggers a security advisories scan for the supported network devices - a1a1-7b93-481b-9e03
/* Triggers a security advisories scan for the supported network devices. The supported devices are switches, routers and wireless controllers with IOS and IOS-XE. If a device is not supported, the SecurityAdvisoryNetworkDevice scanStatus will be Failed with appropriate comments.


@param TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!triggers-a-security-advisories-scan-for-the-supported-network-devices
*/
func (s *ComplianceService) TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesQueryParams *TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesQueryParams) (*ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices, *resty.Response, error) {
	path := "/dna/intent/api/v1/securityAdvisories/triggerScan"

	queryString, _ := query.Values(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices(TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevicesQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices")
	}

	result := response.Result().(*ResponseComplianceTriggersASecurityAdvisoriesScanForTheSupportedNetworkDevices)
	return result, response, err

}
