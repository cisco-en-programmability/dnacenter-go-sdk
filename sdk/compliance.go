package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ComplianceService service

type GetComplianceStatusQueryParams struct {
	ComplianceStatus string  `url:"complianceStatus,omitempty"` //Compliance status can be have value among 'COMPLIANT','NON_COMPLIANT','IN_PROGRESS', 'ERROR'
	DeviceUUID       string  `url:"deviceUuid,omitempty"`       //Comma separated deviceUuids
	Offset           float64 `url:"offset,omitempty"`           //offset/starting row
	Limit            float64 `url:"limit,omitempty"`            //Number of records to be retrieved
}
type GetComplianceStatusCountQueryParams struct {
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
}
type GetComplianceDetailQueryParams struct {
	ComplianceType   string `url:"complianceType,omitempty"`   //complianceType can have any value among 'NETWORK_PROFILE', 'IMAGE', 'APPLICATION_VISIBILITY', 'FABRIC', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
	DeviceUUID       string `url:"deviceUuid,omitempty"`       //Comma separated deviceUuids
	Offset           string `url:"offset,omitempty"`           //offset/starting row
	Limit            string `url:"limit,omitempty"`            //Number of records to be retrieved
}
type GetComplianceDetailCountQueryParams struct {
	ComplianceType   string `url:"complianceType,omitempty"`   //complianceType can have any value among 'NETWORK_PROFILE', 'IMAGE', 'APPLICATION_VISIBILITY', 'FABRIC', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
	ComplianceStatus string `url:"complianceStatus,omitempty"` //Compliance status can have value among 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'
}
type ComplianceDetailsOfDeviceQueryParams struct {
	Category       string `url:"category,omitempty"`       //complianceCategory can have any value among 'INTENT', 'RUNNING_CONFIG'
	ComplianceType string `url:"complianceType,omitempty"` //complianceType can have any value among 'NETWORK_DESIGN', 'NETWORK_PROFILE', 'FABRIC', 'POLICY', 'RUNNING_CONFIG'
	DiffList       bool   `url:"diffList,omitempty"`       //diff list [ pass true to fetch the diff list ]
	Key            string `url:"key,omitempty"`            //extended attribute key
	Value          string `url:"value,omitempty"`          //extended attribute value
}

type ResponseComplianceGetComplianceStatus struct {
	Version  string                                           `json:"version,omitempty"`  // Version
	Response *[]ResponseComplianceGetComplianceStatusResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetComplianceStatusResponse struct {
	DeviceUUID       string   `json:"deviceUuid,omitempty"`       // Device Uuid
	ComplianceStatus string   `json:"complianceStatus,omitempty"` // Compliance Status
	Message          string   `json:"message,omitempty"`          // Message
	ScheduleTime     *float64 `json:"scheduleTime,omitempty"`     // Schedule Time
	LastUpdateTime   *float64 `json:"lastUpdateTime,omitempty"`   // Last Update Time
}
type ResponseComplianceRunCompliance struct {
	Version  string                                   `json:"version,omitempty"`  // Version
	Response *ResponseComplianceRunComplianceResponse `json:"response,omitempty"` //
}
type ResponseComplianceRunComplianceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseComplianceGetComplianceStatusCount struct {
	Version  string   `json:"version,omitempty"`  // Version
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseComplianceGetComplianceDetail struct {
	Version  string                                           `json:"version,omitempty"`  // Version
	Response *[]ResponseComplianceGetComplianceDetailResponse `json:"response,omitempty"` //
}
type ResponseComplianceGetComplianceDetailResponse struct {
	ComplianceType string `json:"complianceType,omitempty"` // Compliance Type
	LastSyncTime   *int   `json:"lastSyncTime,omitempty"`   // Last Sync Time
	DeviceUUID     string `json:"deviceUuid,omitempty"`     // Device Uuid
	DisplayName    string `json:"displayName,omitempty"`    // Display Name
	Status         string `json:"status,omitempty"`         // Status
	Category       string `json:"category,omitempty"`       // Category
	LastUpdateTime *int   `json:"lastUpdateTime,omitempty"` // Last Update Time
	State          string `json:"state,omitempty"`          // State
}
type ResponseComplianceGetComplianceDetailCount struct {
	Version  string   `json:"version,omitempty"`  // Version
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseComplianceDeviceComplianceStatus struct {
	Version  string                                            `json:"version,omitempty"`  // Version
	Response *ResponseComplianceDeviceComplianceStatusResponse `json:"response,omitempty"` //
}
type ResponseComplianceDeviceComplianceStatusResponse struct {
	DeviceUUID       string   `json:"deviceUuid,omitempty"`       // Device Uuid
	ComplianceStatus string   `json:"complianceStatus,omitempty"` // Compliance Status
	Message          string   `json:"message,omitempty"`          // Message
	ScheduleTime     *float64 `json:"scheduleTime,omitempty"`     // Schedule Time
	LastUpdateTime   *float64 `json:"lastUpdateTime,omitempty"`   // Last Update Time
}
type ResponseComplianceComplianceDetailsOfDevice struct {
	DeviceUUID string                                                 `json:"deviceUuid,omitempty"` // Device Uuid
	Version    string                                                 `json:"version,omitempty"`    // Version
	Response   *[]ResponseComplianceComplianceDetailsOfDeviceResponse `json:"response,omitempty"`   //
}
type ResponseComplianceComplianceDetailsOfDeviceResponse struct {
	DisplayName       string                                                               `json:"displayName,omitempty"`       // Display Name
	ComplianceType    string                                                               `json:"complianceType,omitempty"`    // Compliance Type
	LastSyncTime      *int                                                                 `json:"lastSyncTime,omitempty"`      // Last Sync Time
	AdditionalDataURL string                                                               `json:"additionalDataURL,omitempty"` // Additional Data U R L
	SourceInfoList    *[]ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoList `json:"sourceInfoList,omitempty"`    //
	DeviceUUID        string                                                               `json:"deviceUuid,omitempty"`        // Device Uuid
	Message           string                                                               `json:"message,omitempty"`           // Message
	State             string                                                               `json:"state,omitempty"`             // State
	Status            string                                                               `json:"status,omitempty"`            // Status
	Category          string                                                               `json:"category,omitempty"`          // Category
	LastUpdateTime    *int                                                                 `json:"lastUpdateTime,omitempty"`    // Last Update Time
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoList struct {
	Count               *float64                                                                      `json:"count,omitempty"`               // Count
	DisplayName         string                                                                        `json:"displayName,omitempty"`         // Display Name
	DiffList            *[]ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffList  `json:"diffList,omitempty"`            //
	SourceEnum          string                                                                        `json:"sourceEnum,omitempty"`          // Source Enum
	LicenseAppName      string                                                                        `json:"licenseAppName,omitempty"`      // License App Name
	ProvisioningArea    string                                                                        `json:"provisioningArea,omitempty"`    // Provisioning Area
	NetworkProfileName  string                                                                        `json:"networkProfileName,omitempty"`  // Network Profile Name
	NameWithBusinessKey string                                                                        `json:"nameWithBusinessKey,omitempty"` // Name With Business Key
	AppName             string                                                                        `json:"appName,omitempty"`             // App Name
	Name                string                                                                        `json:"name,omitempty"`                // Name
	Type                string                                                                        `json:"type,omitempty"`                // Type
	BusinessKey         *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKey `json:"businessKey,omitempty"`         //
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListDiffList struct {
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	MoveFromPath       string `json:"moveFromPath,omitempty"`       // Move From Path
	Op                 string `json:"op,omitempty"`                 // Op
	ConfiguredValue    string `json:"configuredValue,omitempty"`    // Configured Value
	IntendedValue      string `json:"intendedValue,omitempty"`      // Intended Value
	Path               string `json:"path,omitempty"`               // Path
	BusinessKey        string `json:"businessKey,omitempty"`        // Business Key
	ExtendedAttributes string `json:"extendedAttributes,omitempty"` // Extended Attributes
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKey struct {
	OtherAttributes       *ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributes `json:"otherAttributes,omitempty"`       //
	ResourceName          string                                                                                       `json:"resourceName,omitempty"`          // Resource Name
	BusinessKeyAttributes string                                                                                       `json:"businessKeyAttributes,omitempty"` // Business Key Attributes
}
type ResponseComplianceComplianceDetailsOfDeviceResponseSourceInfoListBusinessKeyOtherAttributes struct {
	CfsAttributes string `json:"cfsAttributes,omitempty"` // Cfs Attributes
	Name          string `json:"name,omitempty"`          // Name
}
type RequestComplianceRunCompliance struct {
	TriggerFull *bool    `json:"triggerFull,omitempty"` //
	Categories  []string `json:"categories,omitempty"`  //
	DeviceUUIDs []string `json:"deviceUuids,omitempty"` //
}

//GetComplianceStatus Get Compliance Status  - dda5-cb9a-49aa-aef6
/* Return compliance status of device(s).


@param GetComplianceStatusQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetComplianceStatus")
	}

	result := response.Result().(*ResponseComplianceGetComplianceStatus)
	return result, response, err

}

//GetComplianceStatusCount Get Compliance Status Count - db99-f919-424a-9f83
/* Return Compliance Status Count


@param GetComplianceStatusCountQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetComplianceStatusCount")
	}

	result := response.Result().(*ResponseComplianceGetComplianceStatusCount)
	return result, response, err

}

//GetComplianceDetail Get Compliance Detail  - dda1-1ae7-4788-9d49
/* Return Compliance Detail


@param GetComplianceDetailQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetComplianceDetail")
	}

	result := response.Result().(*ResponseComplianceGetComplianceDetail)
	return result, response, err

}

//GetComplianceDetailCount Get Compliance Detail Count - 3eb6-58c3-4549-94df
/* Return  Compliance Count Detail


@param GetComplianceDetailCountQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetComplianceDetailCount")
	}

	result := response.Result().(*ResponseComplianceGetComplianceDetailCount)
	return result, response, err

}

//DeviceComplianceStatus Device Compliance Status - 7aa8-5ad5-48ea-94a7
/* Return compliance status of a device.


@param deviceUUID deviceUuid path parameter.
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
		return nil, response, fmt.Errorf("error with operation DeviceComplianceStatus")
	}

	result := response.Result().(*ResponseComplianceDeviceComplianceStatus)
	return result, response, err

}

//ComplianceDetailsOfDevice Compliance Details of Device - 52bf-e904-45aa-b017
/* Return compliance detailed report for a device.


@param deviceUUID deviceUuid path parameter.
@param ComplianceDetailsOfDeviceQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation ComplianceDetailsOfDevice")
	}

	result := response.Result().(*ResponseComplianceComplianceDetailsOfDevice)
	return result, response, err

}

//RunCompliance Run Compliance - f6ae-c8a7-4428-a9ff
/* Run compliance check for device(s).


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
		return nil, response, fmt.Errorf("error with operation RunCompliance")
	}

	result := response.Result().(*ResponseComplianceRunCompliance)
	return result, response, err

}
