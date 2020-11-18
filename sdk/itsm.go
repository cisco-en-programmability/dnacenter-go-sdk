package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ITSMService is the service to communicate with the ITSM API endpoint
type ITSMService service

// RetryIntegrationEventsRequest is the retryIntegrationEventsRequest definition
type RetryIntegrationEventsRequest []string

// GetFailedITSMEventsResponse is the getFailedITSMEventsResponse definition
type GetFailedITSMEventsResponse struct {
	Category       string                                    `json:"category,omitempty"`       //
	Description    string                                    `json:"description,omitempty"`    //
	Domain         string                                    `json:"domain,omitempty"`         //
	EnrichmentInfo GetFailedITSMEventsResponseEnrichmentInfo `json:"enrichmentInfo,omitempty"` //
	EventID        string                                    `json:"eventId,omitempty"`        //
	InstanceID     string                                    `json:"instanceId,omitempty"`     //
	Name           string                                    `json:"name,omitempty"`           //
	Severity       string                                    `json:"severity,omitempty"`       //
	Source         string                                    `json:"source,omitempty"`         //
	SubDomain      string                                    `json:"subDomain,omitempty"`      //
	Timestamp      int                                       `json:"timestamp,omitempty"`      //
	Type           string                                    `json:"type,omitempty"`           //
}

// GetFailedITSMEventsResponseEnrichmentInfo is the getFailedITSMEventsResponseEnrichmentInfo definition
type GetFailedITSMEventsResponseEnrichmentInfo struct {
	ErrorCode                      string `json:"errorCode,omitempty"`                      //
	ErrorDescription               string `json:"errorDescription,omitempty"`               //
	EventStatus                    string `json:"eventStatus,omitempty"`                    //
	ResponseReceivedFromITSMSystem string `json:"responseReceivedFromITSMSystem,omitempty"` //
}

// RetryIntegrationEventsResponse is the retryIntegrationEventsResponse definition
type RetryIntegrationEventsResponse struct {
	ExecutionID        string `json:"executionId,omitempty"`        //
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` //
	Message            string `json:"message,omitempty"`            //
}

// GetFailedITSMEventsQueryParams defines the query parameters for this request
type GetFailedITSMEventsQueryParams struct {
	InstanceID string `url:"instanceId,omitempty"` // Instance Id of the failed event as in the Runtime Dashboard
}

// GetFailedITSMEvents getFailedITSMEvents
/* Used to retrieve the list of integration events that failed to create tickets in ITSM
@param instanceID Instance Id of the failed event as in the Runtime Dashboard
*/
func (s *ITSMService) GetFailedITSMEvents(getFailedITSMEventsQueryParams *GetFailedITSMEventsQueryParams) (*GetFailedITSMEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/integration/events"

	queryString, _ := query.Values(getFailedITSMEventsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetFailedITSMEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetFailedITSMEventsResponse)
	return result, response, err
}

// RetryIntegrationEvents retryIntegrationEvents
/* Allows retry of multiple failed ITSM event instances. The retry request payload can be given as a list of strings: ["instance1","instance2","instance3",..] A minimum of one instance Id is mandatory. The list of failed event instance Ids can be retrieved using the 'Get Failed ITSM Events' API in the 'instanceId' attribute.
 */
func (s *ITSMService) RetryIntegrationEvents(retryIntegrationEventsRequest *[]RetryIntegrationEventsRequest) (*RetryIntegrationEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/integration/events"

	response, err := RestyClient.R().
		SetBody(retryIntegrationEventsRequest).
		SetResult(&RetryIntegrationEventsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*RetryIntegrationEventsResponse)
	return result, response, err
}
