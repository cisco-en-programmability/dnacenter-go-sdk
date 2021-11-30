package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// EventManagementService is the service to communicate with the EventManagement API endpoint
type EventManagementService service

// CreateEventSubscriptionsRequest is the createEventSubscriptionsRequest definition
type CreateEventSubscriptionsRequest struct {
	Description           string                                                 `json:"description,omitempty"`           //
	Filter                CreateEventSubscriptionsRequestFilter                  `json:"filter,omitempty"`                //
	Name                  string                                                 `json:"name,omitempty"`                  //
	SubscriptionEndpoints []CreateEventSubscriptionsRequestSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	SubscriptionID        string                                                 `json:"subscriptionId,omitempty"`        //
	Version               string                                                 `json:"version,omitempty"`               //
}

// CreateEventSubscriptionsRequestFilter is the createEventSubscriptionsRequestFilter definition
type CreateEventSubscriptionsRequestFilter struct {
	EventIDs []string `json:"eventIds,omitempty"` //
}

// CreateEventSubscriptionsRequestFilterEventIDs is the createEventSubscriptionsRequestFilterEventIDs definition
type CreateEventSubscriptionsRequestFilterEventIDs []string

// CreateEventSubscriptionsRequestSubscriptionEndpoints is the createEventSubscriptionsRequestSubscriptionEndpoints definition
type CreateEventSubscriptionsRequestSubscriptionEndpoints struct {
	InstanceID          string                                                                  `json:"instanceId,omitempty"`          //
	SubscriptionDetails CreateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}

// CreateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails is the createEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails definition
type CreateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` //
	Method        string `json:"method,omitempty"`        //
	Name          string `json:"name,omitempty"`          //
	URL           string `json:"url,omitempty"`           //
}

// UpdateEventSubscriptionsRequest is the updateEventSubscriptionsRequest definition
type UpdateEventSubscriptionsRequest struct {
	Description           string                                                 `json:"description,omitempty"`           //
	Filter                UpdateEventSubscriptionsRequestFilter                  `json:"filter,omitempty"`                //
	Name                  string                                                 `json:"name,omitempty"`                  //
	SubscriptionEndpoints []UpdateEventSubscriptionsRequestSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	SubscriptionID        string                                                 `json:"subscriptionId,omitempty"`        //
	Version               string                                                 `json:"version,omitempty"`               //
}

// UpdateEventSubscriptionsRequestFilter is the updateEventSubscriptionsRequestFilter definition
type UpdateEventSubscriptionsRequestFilter struct {
	EventIDs []string `json:"eventIds,omitempty"` //
}

// UpdateEventSubscriptionsRequestFilterEventIDs is the updateEventSubscriptionsRequestFilterEventIDs definition
type UpdateEventSubscriptionsRequestFilterEventIDs []string

// UpdateEventSubscriptionsRequestSubscriptionEndpoints is the updateEventSubscriptionsRequestSubscriptionEndpoints definition
type UpdateEventSubscriptionsRequestSubscriptionEndpoints struct {
	InstanceID          string                                                                  `json:"instanceId,omitempty"`          //
	SubscriptionDetails UpdateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}

// UpdateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails is the updateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails definition
type UpdateEventSubscriptionsRequestSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` //
	Method        string `json:"method,omitempty"`        //
	Name          string `json:"name,omitempty"`          //
	URL           string `json:"url,omitempty"`           //
}

// CountOfEventSubscriptionsResponse is the countOfEventSubscriptionsResponse definition
type CountOfEventSubscriptionsResponse struct {
	Response float64 `json:"response,omitempty"` //
}

// CountOfEventsResponse is the countOfEventsResponse definition
type CountOfEventsResponse struct {
	Response float64 `json:"response,omitempty"` //
}

// CountOfNotificationsResponse is the countOfNotificationsResponse definition
type CountOfNotificationsResponse struct {
	Response float64 `json:"response,omitempty"` //
}

// CreateEventSubscriptionsResponse is the createEventSubscriptionsResponse definition
type CreateEventSubscriptionsResponse struct {
	StatusURI string `json:"statusUri,omitempty"` //
}

// DeleteEventSubscriptionsResponse is the deleteEventSubscriptionsResponse definition
type DeleteEventSubscriptionsResponse struct {
	StatusURI string `json:"statusUri,omitempty"` //
}

// GetEventSubscriptionsResponse is the getEventSubscriptionsResponse definition
type GetEventSubscriptionsResponse struct {
	Description           string                                               `json:"description,omitempty"`           //
	Filter                GetEventSubscriptionsResponseFilter                  `json:"filter,omitempty"`                //
	Name                  string                                               `json:"name,omitempty"`                  //
	SubscriptionEndpoints []GetEventSubscriptionsResponseSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Version               string                                               `json:"version,omitempty"`               //
}

// GetEventSubscriptionsResponseFilter is the getEventSubscriptionsResponseFilter definition
type GetEventSubscriptionsResponseFilter struct {
	EventIDs []string `json:"eventIds,omitempty"` //
}

// GetEventSubscriptionsResponseFilterEventIDs is the getEventSubscriptionsResponseFilterEventIDs definition
type GetEventSubscriptionsResponseFilterEventIDs []string

// GetEventSubscriptionsResponseSubscriptionEndpoints is the getEventSubscriptionsResponseSubscriptionEndpoints definition
type GetEventSubscriptionsResponseSubscriptionEndpoints struct {
	ID                  string                                                                `json:"id,omitempty"`                  //
	InstanceID          string                                                                `json:"instanceId,omitempty"`          //
	SubscriptionDetails GetEventSubscriptionsResponseSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}

// GetEventSubscriptionsResponseSubscriptionEndpointsSubscriptionDetails is the getEventSubscriptionsResponseSubscriptionEndpointsSubscriptionDetails definition
type GetEventSubscriptionsResponseSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` //
	Method        string `json:"method,omitempty"`        //
	Name          string `json:"name,omitempty"`          //
	URL           string `json:"url,omitempty"`           //
}

// GetEventsResponse is the getEventsResponse definition
type GetEventsResponse struct {
	Category          string   `json:"category,omitempty"`          //
	Description       string   `json:"description,omitempty"`       //
	Details           string   `json:"details,omitempty"`           //
	Domain            string   `json:"domain,omitempty"`            //
	EventID           string   `json:"eventId,omitempty"`           //
	Name              string   `json:"name,omitempty"`              //
	NameSpace         string   `json:"nameSpace,omitempty"`         //
	Severity          float64  `json:"severity,omitempty"`          //
	SubDomain         string   `json:"subDomain,omitempty"`         //
	SubscriptionTypes []string `json:"subscriptionTypes,omitempty"` //
	Tags              []string `json:"tags,omitempty"`              //
	Type              string   `json:"type,omitempty"`              //
	Version           string   `json:"version,omitempty"`           //
}

// GetEventsResponseSubscriptionTypes is the getEventsResponseSubscriptionTypes definition
type GetEventsResponseSubscriptionTypes []string

// GetEventsResponseTags is the getEventsResponseTags definition
type GetEventsResponseTags []string

// GetNotificationsResponse is the getNotificationsResponse definition
type GetNotificationsResponse struct {
	Category    string  `json:"category,omitempty"`    //
	Context     string  `json:"context,omitempty"`     //
	Description string  `json:"description,omitempty"` //
	Details     string  `json:"details,omitempty"`     //
	Domain      string  `json:"domain,omitempty"`      //
	EventID     string  `json:"eventId,omitempty"`     //
	InstanceID  string  `json:"instanceId,omitempty"`  //
	Name        string  `json:"name,omitempty"`        //
	Namespace   string  `json:"namespace,omitempty"`   //
	Severity    float64 `json:"severity,omitempty"`    //
	Source      string  `json:"source,omitempty"`      //
	SubDomain   string  `json:"subDomain,omitempty"`   //
	TenantID    string  `json:"tenantId,omitempty"`    //
	Timestamp   float64 `json:"timestamp,omitempty"`   //
	Type        string  `json:"type,omitempty"`        //
}

// GetStatusAPIForEventsResponse is the getStatusAPIForEventsResponse definition
type GetStatusAPIForEventsResponse struct {
	APIStatus     string `json:"apiStatus,omitempty"`     //
	ErrorMessage  string `json:"errorMessage,omitempty"`  //
	StatusMessage string `json:"statusMessage,omitempty"` //
}

// UpdateEventSubscriptionsResponse is the updateEventSubscriptionsResponse definition
type UpdateEventSubscriptionsResponse struct {
	StatusURI string `json:"statusUri,omitempty"` //
}

// CountOfEventSubscriptionsQueryParams defines the query parameters for this request
type CountOfEventSubscriptionsQueryParams struct {
	EventIDs string `url:"eventIds,omitempty"` // List of subscriptions related to the respective eventIds
}

// CountOfEventSubscriptions countOfEventSubscriptions
/* Returns the Count of EventSubscriptions
@param eventIDs List of subscriptions related to the respective eventIds
*/
func (s *EventManagementService) CountOfEventSubscriptions(countOfEventSubscriptionsQueryParams *CountOfEventSubscriptionsQueryParams) (*CountOfEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription/count"

	queryString, _ := query.Values(countOfEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountOfEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation countOfEventSubscriptions")
	}

	result := response.Result().(*CountOfEventSubscriptionsResponse)
	return result, response, err
}

// CountOfEventsQueryParams defines the query parameters for this request
type CountOfEventsQueryParams struct {
	EventID string `url:"eventId,omitempty"` // The registered EventId should be provided
	Tags    string `url:"tags,omitempty"`    // The registered Tags should be provided
}

// CountOfEvents countOfEvents
/* Get the count of registered events with provided eventIds or tags as mandatory
@param eventID The registered EventId should be provided
@param tags The registered Tags should be provided
*/
func (s *EventManagementService) CountOfEvents(countOfEventsQueryParams *CountOfEventsQueryParams) (*CountOfEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/events/count"

	queryString, _ := query.Values(countOfEventsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountOfEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation countOfEvents")
	}

	result := response.Result().(*CountOfEventsResponse)
	return result, response, err
}

// CountOfNotificationsQueryParams defines the query parameters for this request
type CountOfNotificationsQueryParams struct {
	EventIDs  string `url:"eventIds,omitempty"`  // The registered EventIds should be provided
	StartTime string `url:"startTime,omitempty"` // StartTime
	EndTime   string `url:"endTime,omitempty"`   // endTime
	Category  string `url:"category,omitempty"`  // category
	Type      string `url:"type,omitempty"`      // type
	Severity  string `url:"severity,omitempty"`  // severity
	Domain    string `url:"domain,omitempty"`    // domain
	SubDomain string `url:"subDomain,omitempty"` // subDomain
	Source    string `url:"source,omitempty"`    // source
}

// CountOfNotifications countOfNotifications
/* Get the Count of Published Notifications
@param eventIDs The registered EventIds should be provided
@param startTime StartTime
@param endTime endTime
@param category category
@param type type
@param severity severity
@param domain domain
@param subDomain subDomain
@param source source
*/
func (s *EventManagementService) CountOfNotifications(countOfNotificationsQueryParams *CountOfNotificationsQueryParams) (*CountOfNotificationsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/event-series/count"

	queryString, _ := query.Values(countOfNotificationsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountOfNotificationsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation countOfNotifications")
	}

	result := response.Result().(*CountOfNotificationsResponse)
	return result, response, err
}

// CreateEventSubscriptions createEventSubscriptions
/* Subscribe SubscriptionEndpoint to list of registered events
 */
func (s *EventManagementService) CreateEventSubscriptions(createEventSubscriptionsRequest *[]CreateEventSubscriptionsRequest) (*CreateEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	response, err := s.client.R().
		SetBody(createEventSubscriptionsRequest).
		SetResult(&CreateEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createEventSubscriptions")
	}

	result := response.Result().(*CreateEventSubscriptionsResponse)
	return result, response, err
}

// DeleteEventSubscriptionsQueryParams defines the query parameters for this request
type DeleteEventSubscriptionsQueryParams struct {
	Subscriptions string `url:"subscriptions,omitempty"` // List of EventSubscriptionId's for removal
}

// DeleteEventSubscriptions deleteEventSubscriptions
/* Delete EventSubscriptions
@param Content-Type Content Type
@param subscriptions List of EventSubscriptionId's for removal
*/
func (s *EventManagementService) DeleteEventSubscriptions(deleteEventSubscriptionsQueryParams *DeleteEventSubscriptionsQueryParams) (*DeleteEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(deleteEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeleteEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteEventSubscriptions")
	}

	result := response.Result().(*DeleteEventSubscriptionsResponse)
	return result, response, err
}

// GetEventSubscriptionsQueryParams defines the query parameters for this request
type GetEventSubscriptionsQueryParams struct {
	EventIDs string  `url:"eventIds,omitempty"` // List of subscriptions related to the respective eventIds
	Offset   float64 `url:"offset,omitempty"`   // The number of Subscriptions's to offset in the resultset whose default value 0
	Limit    float64 `url:"limit,omitempty"`    // The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy   string  `url:"sortBy,omitempty"`   // SortBy field name
	Order    string  `url:"order,omitempty"`    // order(asc/desc)
}

// GetEventSubscriptions getEventSubscriptions
/* Gets the list of Subscriptions's based on provided offset and limit
@param eventIDs List of subscriptions related to the respective eventIds
@param offset The number of Subscriptions's to offset in the resultset whose default value 0
@param limit The number of Subscriptions's to limit in the resultset whose default value 10
@param sortBy SortBy field name
@param order order(asc/desc)
*/
func (s *EventManagementService) GetEventSubscriptions(getEventSubscriptionsQueryParams *GetEventSubscriptionsQueryParams) (*GetEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(getEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getEventSubscriptions")
	}

	result := response.Result().(*GetEventSubscriptionsResponse)
	return result, response, err
}

// GetEventsQueryParams defines the query parameters for this request
type GetEventsQueryParams struct {
	EventID string  `url:"eventId,omitempty"` // The registered EventId should be provided
	Tags    string  `url:"tags,omitempty"`    // The registered Tags should be provided
	Offset  float64 `url:"offset,omitempty"`  // The number of Registries to offset in the resultset whose default value 0
	Limit   float64 `url:"limit,omitempty"`   // The number of Registries to limit in the resultset whose default value 10
	SortBy  string  `url:"sortBy,omitempty"`  // SortBy field name
	Order   string  `url:"order,omitempty"`   // order(asc/desc)
}

// GetEvents getEvents
/* Gets the list of registered Events with provided eventIds or tags as mandatory
@param eventID The registered EventId should be provided
@param tags The registered Tags should be provided
@param offset The number of Registries to offset in the resultset whose default value 0
@param limit The number of Registries to limit in the resultset whose default value 10
@param sortBy SortBy field name
@param order order(asc/desc)
*/
func (s *EventManagementService) GetEvents(getEventsQueryParams *GetEventsQueryParams) (*GetEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/events"

	queryString, _ := query.Values(getEventsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getEvents")
	}

	result := response.Result().(*GetEventsResponse)
	return result, response, err
}

// GetNotificationsQueryParams defines the query parameters for this request
type GetNotificationsQueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  // The registered EventIds should be provided
	StartTime string  `url:"startTime,omitempty"` // StartTime
	EndTime   string  `url:"endTime,omitempty"`   // endTime
	Category  string  `url:"category,omitempty"`  // category
	Type      string  `url:"type,omitempty"`      // type
	Severity  string  `url:"severity,omitempty"`  // severity
	Domain    string  `url:"domain,omitempty"`    // domain
	SubDomain string  `url:"subDomain,omitempty"` // subDomain
	Source    string  `url:"source,omitempty"`    // source
	Offset    float64 `url:"offset,omitempty"`    // Offset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     // Limit whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    // SortBy field name
	Order     string  `url:"order,omitempty"`     // order(asc/desc)
}

// GetNotifications getNotifications
/* Get the list of Published Notifications
@param eventIDs The registered EventIds should be provided
@param startTime StartTime
@param endTime endTime
@param category category
@param type type
@param severity severity
@param domain domain
@param subDomain subDomain
@param source source
@param offset Offset whose default value 0
@param limit Limit whose default value 10
@param sortBy SortBy field name
@param order order(asc/desc)
*/
func (s *EventManagementService) GetNotifications(getNotificationsQueryParams *GetNotificationsQueryParams) (*GetNotificationsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/event-series"

	queryString, _ := query.Values(getNotificationsQueryParams)

	response, err := s.client.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetNotificationsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getNotifications")
	}

	result := response.Result().(*GetNotificationsResponse)
	return result, response, err
}

// GetStatusAPIForEvents getStatusAPIForEvents
/* Get the Status of events API calls with provided executionId as mandatory path parameter
@param executionID Execution ID
*/
func (s *EventManagementService) GetStatusAPIForEvents(executionID string) (*GetStatusAPIForEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/api-status/{executionId}"
	path = strings.Replace(path, "{"+"executionId"+"}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetResult(&GetStatusAPIForEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getStatusAPIForEvents")
	}

	result := response.Result().(*GetStatusAPIForEventsResponse)
	return result, response, err
}

// UpdateEventSubscriptions updateEventSubscriptions
/* Update SubscriptionEndpoint to list of registered events
 */
func (s *EventManagementService) UpdateEventSubscriptions(updateEventSubscriptionsRequest *[]UpdateEventSubscriptionsRequest) (*UpdateEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	response, err := s.client.R().
		SetBody(updateEventSubscriptionsRequest).
		SetResult(&UpdateEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateEventSubscriptions")
	}

	result := response.Result().(*UpdateEventSubscriptionsResponse)
	return result, response, err
}
