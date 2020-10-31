package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// EventManagementService is the service to communicate with the EventManagement API endpoint
type EventManagementService service

// CreateEventSubscriptionsRequest is the CreateEventSubscriptionsRequest definition
type CreateEventSubscriptionsRequest struct {
	Description           string                  `json:"description,omitempty"`           //
	Filter                Filter                  `json:"filter,omitempty"`                //
	Name                  string                  `json:"name,omitempty"`                  //
	SubscriptionEndpoints []SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	SubscriptionID        string                  `json:"subscriptionId,omitempty"`        //
	Version               string                  `json:"version,omitempty"`               //
}

// Filter is the Filter definition
type Filter struct {
	EventIDs []string `json:"eventIds,omitempty"` //
}

// SubscriptionDetails is the SubscriptionDetails definition
type SubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` //
	Method        string `json:"method,omitempty"`        //
	Name          string `json:"name,omitempty"`          //
	URL           string `json:"url,omitempty"`           //
}

// SubscriptionEndpoints is the SubscriptionEndpoints definition
type SubscriptionEndpoints struct {
	InstanceID          string              `json:"instanceId,omitempty"`          //
	SubscriptionDetails SubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}

// UpdateEventSubscriptionsRequest is the UpdateEventSubscriptionsRequest definition
type UpdateEventSubscriptionsRequest struct {
	Description           string                  `json:"description,omitempty"`           //
	Filter                Filter                  `json:"filter,omitempty"`                //
	Name                  string                  `json:"name,omitempty"`                  //
	SubscriptionEndpoints []SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	SubscriptionID        string                  `json:"subscriptionId,omitempty"`        //
	Version               string                  `json:"version,omitempty"`               //
}

// CountOfEventSubscriptionsResponse is the CountOfEventSubscriptionsResponse definition
type CountOfEventSubscriptionsResponse struct {
	Response int `json:"response,omitempty"` //
}

// CountOfEventsResponse is the CountOfEventsResponse definition
type CountOfEventsResponse struct {
	Response int `json:"response,omitempty"` //
}

// CountOfNotificationsResponse is the CountOfNotificationsResponse definition
type CountOfNotificationsResponse struct {
	Response int `json:"response,omitempty"` //
}

// CreateEventSubscriptionsResponse is the CreateEventSubscriptionsResponse definition
type CreateEventSubscriptionsResponse struct {
	StatusURI string `json:"statusUri,omitempty"` //
}

// DeleteEventSubscriptionsResponse is the DeleteEventSubscriptionsResponse definition
type DeleteEventSubscriptionsResponse struct {
	StatusURI string `json:"statusUri,omitempty"` //
}

// GetEventSubscriptionsResponse is the GetEventSubscriptionsResponse definition
type GetEventSubscriptionsResponse struct {
	Description           string                  `json:"description,omitempty"`           //
	Filter                Filter                  `json:"filter,omitempty"`                //
	Name                  string                  `json:"name,omitempty"`                  //
	SubscriptionEndpoints []SubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Version               string                  `json:"version,omitempty"`               //
}

// GetEventsResponse is the GetEventsResponse definition
type GetEventsResponse struct {
	Category          string   `json:"category,omitempty"`          //
	Description       string   `json:"description,omitempty"`       //
	Details           string   `json:"details,omitempty"`           //
	Domain            string   `json:"domain,omitempty"`            //
	EventID           string   `json:"eventId,omitempty"`           //
	Name              string   `json:"name,omitempty"`              //
	NameSpace         string   `json:"nameSpace,omitempty"`         //
	Severity          int      `json:"severity,omitempty"`          //
	SubDomain         string   `json:"subDomain,omitempty"`         //
	SubscriptionTypes []string `json:"subscriptionTypes,omitempty"` //
	Tags              []string `json:"tags,omitempty"`              //
	Type              string   `json:"type,omitempty"`              //
	Version           string   `json:"version,omitempty"`           //
}

// GetNotificationsResponse is the GetNotificationsResponse definition
type GetNotificationsResponse struct {
	Category    string `json:"category,omitempty"`    //
	Context     string `json:"context,omitempty"`     //
	Description string `json:"description,omitempty"` //
	Details     string `json:"details,omitempty"`     //
	Domain      string `json:"domain,omitempty"`      //
	EventID     string `json:"eventId,omitempty"`     //
	InstanceID  string `json:"instanceId,omitempty"`  //
	Name        string `json:"name,omitempty"`        //
	Namespace   string `json:"namespace,omitempty"`   //
	Severity    int    `json:"severity,omitempty"`    //
	Source      string `json:"source,omitempty"`      //
	SubDomain   string `json:"subDomain,omitempty"`   //
	TenantID    string `json:"tenantId,omitempty"`    //
	Timestamp   int    `json:"timestamp,omitempty"`   //
	Type        string `json:"type,omitempty"`        //
}

// GetStatusAPIForEventsResponse is the GetStatusAPIForEventsResponse definition
type GetStatusAPIForEventsResponse struct {
	APIStatus     string `json:"apiStatus,omitempty"`     //
	ErrorMessage  string `json:"errorMessage,omitempty"`  //
	StatusMessage string `json:"statusMessage,omitempty"` //
}

// UpdateEventSubscriptionsResponse is the UpdateEventSubscriptionsResponse definition
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

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountOfEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
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
/* Get the count of registered events with provided eventIDs or tags as mandatory
@param eventID The registered EventId should be provided
@param tags The registered Tags should be provided
*/
func (s *EventManagementService) CountOfEvents(countOfEventsQueryParams *CountOfEventsQueryParams) (*CountOfEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/events/count"

	queryString, _ := query.Values(countOfEventsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountOfEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
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

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountOfNotificationsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountOfNotificationsResponse)
	return result, response, err

}

// CreateEventSubscriptions createEventSubscriptions
/* Subscribe SubscriptionEndpoint to list of registered events
 */
func (s *EventManagementService) CreateEventSubscriptions(createEventSubscriptionsRequest *CreateEventSubscriptionsRequest) (*CreateEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	response, err := RestyClient.R().
		SetBody(createEventSubscriptionsRequest).
		SetResult(&CreateEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateEventSubscriptionsResponse)
	return result, response, err

}

// DeleteEventSubscriptionsQueryParams defines the query parameters for this request
type DeleteEventSubscriptionsQueryParams struct {
	Subscriptions string `url:"subscriptions,omitempty"` // List of EventSubscriptionID's for removal
}

// DeleteEventSubscriptions deleteEventSubscriptions
/* Delete EventSubscriptions
@param Content-Type Content Type
@param subscriptions List of EventSubscriptionID's for removal
*/
func (s *EventManagementService) DeleteEventSubscriptions(deleteEventSubscriptionsQueryParams *DeleteEventSubscriptionsQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(deleteEventSubscriptionsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetEventSubscriptionsQueryParams defines the query parameters for this request
type GetEventSubscriptionsQueryParams struct {
	EventIDs string `url:"eventIds,omitempty"` // List of subscriptions related to the respective eventIds
	Offset   int    `url:"offset,omitempty"`   // The int of Subscriptions's to offset in the resultset whose default value 0
	Limit    int    `url:"limit,omitempty"`    // The int of Subscriptions's to limit in the resultset whose default value 10
	SortBy   string `url:"sortBy,omitempty"`   // SortBy field name
	Order    string `url:"order,omitempty"`    // order(asc/desc)
}

// GetEventSubscriptions getEventSubscriptions
/* Gets the list of Subscriptions's based on provided offset and limit
@param eventIDs List of subscriptions related to the respective eventIds
@param offset The int of Subscriptions's to offset in the resultset whose default value 0
@param limit The int of Subscriptions's to limit in the resultset whose default value 10
@param sortBy SortBy field name
@param order order(asc/desc)
*/
func (s *EventManagementService) GetEventSubscriptions(getEventSubscriptionsQueryParams *GetEventSubscriptionsQueryParams) (*GetEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(getEventSubscriptionsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetEventSubscriptionsResponse)
	return result, response, err

}

// GetEventsQueryParams defines the query parameters for this request
type GetEventsQueryParams struct {
	EventID string `url:"eventId,omitempty"` // The registered EventId should be provided
	Tags    string `url:"tags,omitempty"`    // The registered Tags should be provided
	Offset  int    `url:"offset,omitempty"`  // The int of Registries to offset in the resultset whose default value 0
	Limit   int    `url:"limit,omitempty"`   // The int of Registries to limit in the resultset whose default value 10
	SortBy  string `url:"sortBy,omitempty"`  // SortBy field name
	Order   string `url:"order,omitempty"`   // order(asc/desc)
}

// GetEvents getEvents
/* Gets the list of registered Events with provided eventIDs or tags as mandatory
@param eventID The registered EventId should be provided
@param tags The registered Tags should be provided
@param offset The int of Registries to offset in the resultset whose default value 0
@param limit The int of Registries to limit in the resultset whose default value 10
@param sortBy SortBy field name
@param order order(asc/desc)
*/
func (s *EventManagementService) GetEvents(getEventsQueryParams *GetEventsQueryParams) (*GetEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/events"

	queryString, _ := query.Values(getEventsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetEventsResponse)
	return result, response, err

}

// GetNotificationsQueryParams defines the query parameters for this request
type GetNotificationsQueryParams struct {
	EventIDs  string `url:"eventIds,omitempty"`  // The registered EventIds should be provided
	StartTime string `url:"startTime,omitempty"` // StartTime
	EndTime   string `url:"endTime,omitempty"`   // endTime
	Category  string `url:"category,omitempty"`  // category
	Type      string `url:"type,omitempty"`      // type
	Severity  string `url:"severity,omitempty"`  // severity
	Domain    string `url:"domain,omitempty"`    // domain
	SubDomain string `url:"subDomain,omitempty"` // subDomain
	Source    string `url:"source,omitempty"`    // source
	Offset    int    `url:"offset,omitempty"`    // Offset whose default value 0
	Limit     int    `url:"limit,omitempty"`     // Limit whose default value 10
	SortBy    string `url:"sortBy,omitempty"`    // SortBy field name
	Order     string `url:"order,omitempty"`     // order(asc/desc)
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

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetNotificationsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetNotificationsResponse)
	return result, response, err

}

// GetStatusAPIForEvents getStatusAPIForEvents
/* Get the Status of events API calls with provided executionID as mandatory path parameter
@param executionID Execution ID
*/
func (s *EventManagementService) GetStatusAPIForEvents(executionID string) (*GetStatusAPIForEventsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/api-status/{executionID}"
	path = strings.Replace(path, "{"+"executionID"+"}", fmt.Sprintf("%v", executionID), -1)

	response, err := RestyClient.R().
		SetResult(&GetStatusAPIForEventsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetStatusAPIForEventsResponse)
	return result, response, err

}

// UpdateEventSubscriptions updateEventSubscriptions
/* Update SubscriptionEndpoint to list of registered events
 */
func (s *EventManagementService) UpdateEventSubscriptions(updateEventSubscriptionsRequest *UpdateEventSubscriptionsRequest) (*UpdateEventSubscriptionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/event/subscription"

	response, err := RestyClient.R().
		SetBody(updateEventSubscriptionsRequest).
		SetResult(&UpdateEventSubscriptionsResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*UpdateEventSubscriptionsResponse)
	return result, response, err

}
