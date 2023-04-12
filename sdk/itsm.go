package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ItsmService service

type GetCmdbSyncStatusQueryParams struct {
	Status string `url:"status,omitempty"` //Supported values are "Success","Failed" and "Unknown". Providing other values will result in all the available sync job status.
	Date   string `url:"date,omitempty"`   //Provide date in "YYYY-MM-DD" format
}
type GetFailedItsmEventsQueryParams struct {
	InstanceID string `url:"instanceId,omitempty"` //Instance Id of the failed event as in the Runtime Dashboard
}

type ResponseItsmGetCmdbSyncStatus []ResponseItemItsmGetCmdbSyncStatus // Array of ResponseItsmGetCMDBSyncStatus
type ResponseItemItsmGetCmdbSyncStatus struct {
	SuccessCount      string                                      `json:"successCount,omitempty"`      // Successfully synchronized device count
	FailureCount      string                                      `json:"failureCount,omitempty"`      // Failed device count
	Devices           *[]ResponseItemItsmGetCmdbSyncStatusDevices `json:"devices,omitempty"`           //
	UnknownErrorCount string                                      `json:"unknownErrorCount,omitempty"` // Unknown error count
	Message           string                                      `json:"message,omitempty"`           // Message
	SyncTime          string                                      `json:"syncTime,omitempty"`          // Synchronization Time
}
type ResponseItemItsmGetCmdbSyncStatusDevices struct {
	DeviceID string `json:"deviceId,omitempty"` // Device Id
	Status   string `json:"status,omitempty"`   // Status "Success" or "Failed"
}
type ResponseItsmGetFailedItsmEvents []ResponseItemItsmGetFailedItsmEvents // Array of ResponseItsmGetFailedITSMEvents
type ResponseItemItsmGetFailedItsmEvents struct {
	InstanceID     string                                             `json:"instanceId,omitempty"`     // Instance Id
	EventID        string                                             `json:"eventId,omitempty"`        // Event Id
	Name           string                                             `json:"name,omitempty"`           // Name
	Type           string                                             `json:"type,omitempty"`           // Type
	Category       string                                             `json:"category,omitempty"`       // Category
	Domain         string                                             `json:"domain,omitempty"`         // Domain
	SubDomain      string                                             `json:"subDomain,omitempty"`      // Sub Domain
	Severity       string                                             `json:"severity,omitempty"`       // Severity
	Source         string                                             `json:"source,omitempty"`         // Source
	Timestamp      *int                                               `json:"timestamp,omitempty"`      // Timestamp
	EnrichmentInfo *ResponseItemItsmGetFailedItsmEventsEnrichmentInfo `json:"enrichmentInfo,omitempty"` //
	Description    string                                             `json:"description,omitempty"`    // Description
}
type ResponseItemItsmGetFailedItsmEventsEnrichmentInfo struct {
	EventStatus                    string                                                                           `json:"eventStatus,omitempty"`                    // Event Status
	ErrorCode                      string                                                                           `json:"errorCode,omitempty"`                      // Error Code
	ErrorDescription               string                                                                           `json:"errorDescription,omitempty"`               // Error Description
	ResponseReceivedFromITSmsystem *ResponseItemItsmGetFailedItsmEventsEnrichmentInfoResponseReceivedFromITSmsystem `json:"responseReceivedFromITSMSystem,omitempty"` // Response Received From ITSMSystem
}
type ResponseItemItsmGetFailedItsmEventsEnrichmentInfoResponseReceivedFromITSmsystem interface{}
type ResponseItsmRetryIntegrationEvents struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type RequestItsmRetryIntegrationEvents []string // Array of RequestItsmRetryIntegrationEvents

//GetCmdbSyncStatus Get CMDB Sync Status - a492-8993-4948-b86c
/* This API allows to retrieve the detail of CMDB sync status.It accepts two query parameter "status","date".The supported values for status field are "Success","Failed","Unknown" and date field should be in "YYYY-MM-DD" format. By default all the cmdb sync status will be send as response and based on the query parameter filtered detail will be send as response.


@param GetCMDBSyncStatusQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-c-m-d-b-sync-status
*/
func (s *ItsmService) GetCmdbSyncStatus(GetCMDBSyncStatusQueryParams *GetCmdbSyncStatusQueryParams) (*ResponseItsmGetCmdbSyncStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/cmdb-sync/detail"

	queryString, _ := query.Values(GetCMDBSyncStatusQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseItsmGetCmdbSyncStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetCmdbSyncStatus")
	}

	result := response.Result().(*ResponseItsmGetCmdbSyncStatus)
	return result, response, err

}

//GetFailedItsmEvents Get Failed ITSM Events - a293-b82a-42a8-ab15
/* Used to retrieve the list of integration events that failed to create tickets in ITSM


@param GetFailedITSMEventsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-failed-itsm-events
*/
func (s *ItsmService) GetFailedItsmEvents(GetFailedITSMEventsQueryParams *GetFailedItsmEventsQueryParams) (*ResponseItsmGetFailedItsmEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration/events"

	queryString, _ := query.Values(GetFailedITSMEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseItsmGetFailedItsmEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetFailedItsmEvents")
	}

	result := response.Result().(*ResponseItsmGetFailedItsmEvents)
	return result, response, err

}

//RetryIntegrationEvents Retry Integration Events - fa9a-9817-4129-af50
/* Allows retry of multiple failed ITSM event instances. The retry request payload can be given as a list of strings: ["instance1","instance2","instance3",..] A minimum of one instance Id is mandatory. The list of failed event instance Ids can be retrieved using the 'Get Failed ITSM Events' API in the 'instanceId' attribute.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retry-integration-events
*/
func (s *ItsmService) RetryIntegrationEvents(requestItsmRetryIntegrationEvents *RequestItsmRetryIntegrationEvents) (*ResponseItsmRetryIntegrationEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/integration/events"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestItsmRetryIntegrationEvents).
		SetResult(&ResponseItsmRetryIntegrationEvents{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RetryIntegrationEvents")
	}

	result := response.Result().(*ResponseItsmRetryIntegrationEvents)
	return result, response, err

}
