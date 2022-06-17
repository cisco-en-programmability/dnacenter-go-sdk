package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type EventManagementService service

type GetAuditLogParentRecordsQueryParams struct {
	InstanceID     string  `url:"instanceId,omitempty"`     //InstanceID of the Audit Log.
	Name           string  `url:"name,omitempty"`           //Audit Log notification event name.
	EventID        string  `url:"eventId,omitempty"`        //Audit Log notification's event ID.
	Category       string  `url:"category,omitempty"`       //Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
	Severity       string  `url:"severity,omitempty"`       //Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
	Domain         string  `url:"domain,omitempty"`         //Audit Log notification's event domain.
	SubDomain      string  `url:"subDomain,omitempty"`      //Audit Log notification's event sub-domain.
	Source         string  `url:"source,omitempty"`         //Audit Log notification's event source.
	UserID         string  `url:"userId,omitempty"`         //Audit Log notification's event userId.
	Context        string  `url:"context,omitempty"`        //Audit Log notification's event correlationId.
	EventHierarchy string  `url:"eventHierarchy,omitempty"` //Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" - Delimiter for hierarchy separation is ".".
	SiteID         string  `url:"siteId,omitempty"`         //Audit Log notification's siteId.
	DeviceID       string  `url:"deviceId,omitempty"`       //Audit Log notification's deviceId.
	IsSystemEvents bool    `url:"isSystemEvents,omitempty"` //Parameter to filter system generated audit-logs.
	Description    string  `url:"description,omitempty"`    //String full/partial search - (Provided input string is case insensitively matched for records).
	Offset         float64 `url:"offset,omitempty"`         //Position of a particular Audit Log record in the data.
	Limit          float64 `url:"limit,omitempty"`          //Number of Audit Log records to be returned per page.
	StartTime      float64 `url:"startTime,omitempty"`      //Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
	EndTime        float64 `url:"endTime,omitempty"`        //End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
	SortBy         string  `url:"sortBy,omitempty"`         //Sort the Audit Logs by certain fields. Supported values are event notification header attributes.
	Order          string  `url:"order,omitempty"`          //Order of the sorted Audit Log records. Default value is desc by timestamp. Supported values: asc, desc.
}
type GetAuditLogSummaryQueryParams struct {
	ParentInstanceID string  `url:"parentInstanceId,omitempty"` //Parent Audit Log record's instanceID.
	IsParentOnly     bool    `url:"isParentOnly,omitempty"`     //Parameter to filter parent only audit-logs.
	InstanceID       string  `url:"instanceId,omitempty"`       //InstanceID of the Audit Log.
	Name             string  `url:"name,omitempty"`             //Audit Log notification event name.
	EventID          string  `url:"eventId,omitempty"`          //Audit Log notification's event ID.
	Category         string  `url:"category,omitempty"`         //Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
	Severity         string  `url:"severity,omitempty"`         //Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
	Domain           string  `url:"domain,omitempty"`           //Audit Log notification's event domain.
	SubDomain        string  `url:"subDomain,omitempty"`        //Audit Log notification's event sub-domain.
	Source           string  `url:"source,omitempty"`           //Audit Log notification's event source.
	UserID           string  `url:"userId,omitempty"`           //Audit Log notification's event userId.
	Context          string  `url:"context,omitempty"`          //Audit Log notification's event correlationId.
	EventHierarchy   string  `url:"eventHierarchy,omitempty"`   //Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" - Delimiter for hierarchy separation is ".".
	SiteID           string  `url:"siteId,omitempty"`           //Audit Log notification's siteId.
	DeviceID         string  `url:"deviceId,omitempty"`         //Audit Log notification's deviceId.
	IsSystemEvents   bool    `url:"isSystemEvents,omitempty"`   //Parameter to filter system generated audit-logs.
	Description      string  `url:"description,omitempty"`      //String full/partial search - (Provided input string is case insensitively matched for records).
	StartTime        float64 `url:"startTime,omitempty"`        //Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
	EndTime          float64 `url:"endTime,omitempty"`          //End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
}
type GetAuditLogRecordsQueryParams struct {
	ParentInstanceID string  `url:"parentInstanceId,omitempty"` //Parent Audit Log record's instanceID.
	InstanceID       string  `url:"instanceId,omitempty"`       //InstanceID of the Audit Log.
	Name             string  `url:"name,omitempty"`             //Audit Log notification event name.
	EventID          string  `url:"eventId,omitempty"`          //Audit Log notification's event ID.
	Category         string  `url:"category,omitempty"`         //Audit Log notification's event category. Supported values: INFO, WARN, ERROR, ALERT, TASK_PROGRESS, TASK_FAILURE, TASK_COMPLETE, COMMAND, QUERY, CONVERSATION
	Severity         string  `url:"severity,omitempty"`         //Audit Log notification's event severity. Supported values: 1, 2, 3, 4, 5.
	Domain           string  `url:"domain,omitempty"`           //Audit Log notification's event domain.
	SubDomain        string  `url:"subDomain,omitempty"`        //Audit Log notification's event sub-domain.
	Source           string  `url:"source,omitempty"`           //Audit Log notification's event source.
	UserID           string  `url:"userId,omitempty"`           //Audit Log notification's event userId.
	Context          string  `url:"context,omitempty"`          //Audit Log notification's event correlationId.
	EventHierarchy   string  `url:"eventHierarchy,omitempty"`   //Audit Log notification's event eventHierarchy. Example: "US.CA.San Jose" OR "US.CA" OR "CA.San Jose" - Delimiter for hierarchy separation is ".".
	SiteID           string  `url:"siteId,omitempty"`           //Audit Log notification's siteId.
	DeviceID         string  `url:"deviceId,omitempty"`         //Audit Log notification's deviceId.
	IsSystemEvents   bool    `url:"isSystemEvents,omitempty"`   //Parameter to filter system generated audit-logs.
	Description      string  `url:"description,omitempty"`      //String full/partial search - (Provided input string is case insensitively matched for records).
	Offset           float64 `url:"offset,omitempty"`           //Position of a particular Audit Log record in the data.
	Limit            float64 `url:"limit,omitempty"`            //Number of Audit Log records to be returned per page.
	StartTime        float64 `url:"startTime,omitempty"`        //Start Time in milliseconds since Epoch Eg. 1597950637211 (when provided endTime is mandatory)
	EndTime          float64 `url:"endTime,omitempty"`          //End Time in milliseconds since Epoch Eg. 1597961437211 (when provided startTime is mandatory)
	SortBy           string  `url:"sortBy,omitempty"`           //Sort the Audit Logs by certain fields. Supported values are event notification header attributes.
	Order            string  `url:"order,omitempty"`            //Order of the sorted Audit Log records. Default value is desc by timestamp. Supported values: asc, desc.
}
type GetNotificationsQueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //The registered EventId should be provided
	StartTime float64 `url:"startTime,omitempty"` //Start Time in milliseconds
	EndTime   float64 `url:"endTime,omitempty"`   //End Time in milliseconds
	Category  string  `url:"category,omitempty"`  //Category
	Type      string  `url:"type,omitempty"`      //Type
	Severity  string  `url:"severity,omitempty"`  //Severity
	Domain    string  `url:"domain,omitempty"`    //Domain
	SubDomain string  `url:"subDomain,omitempty"` //Sub Domain
	Source    string  `url:"source,omitempty"`    //Source
	Offset    float64 `url:"offset,omitempty"`    //Start Offset
	Limit     float64 `url:"limit,omitempty"`     //# of records
	SortBy    string  `url:"sortBy,omitempty"`    //Sort By column
	Order     string  `url:"order,omitempty"`     //Ascending/Descending order [asc/desc]
	Tags      string  `url:"tags,omitempty"`      //Tags
	Namespace string  `url:"namespace,omitempty"` //Namespace
	SiteID    string  `url:"siteId,omitempty"`    //Site Id
}
type CountOfNotificationsQueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //The registered EventId should be provided
	StartTime float64 `url:"startTime,omitempty"` //Start Time in milliseconds
	EndTime   float64 `url:"endTime,omitempty"`   //End Time in milliseconds
	Category  string  `url:"category,omitempty"`  //Category
	Type      string  `url:"type,omitempty"`      //Type
	Severity  string  `url:"severity,omitempty"`  //Severity
	Domain    string  `url:"domain,omitempty"`    //Domain
	SubDomain string  `url:"subDomain,omitempty"` //Sub Domain
	Source    string  `url:"source,omitempty"`    //Source
}
type GetEventSubscriptionsQueryParams struct {
	EventIDs string  `url:"eventIds,omitempty"` //List of subscriptions related to the respective eventIds
	Offset   float64 `url:"offset,omitempty"`   //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit    float64 `url:"limit,omitempty"`    //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy   string  `url:"sortBy,omitempty"`   //SortBy field name
	Order    string  `url:"order,omitempty"`    //order(asc/desc)
}
type DeleteEventSubscriptionsQueryParams struct {
	Subscriptions string `url:"subscriptions,omitempty"` //List of EventSubscriptionId's for removal
}
type GetEmailSubscriptionDetailsQueryParams struct {
	Name       string  `url:"name,omitempty"`       //Name of the specific configuration
	InstanceID string  `url:"instanceId,omitempty"` //Instance Id of the specific configuration
	Offset     float64 `url:"offset,omitempty"`     //The number of Email Subscription detail's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of Email Subscription detail's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type GetRestWebhookSubscriptionDetailsQueryParams struct {
	Name       string  `url:"name,omitempty"`       //Name of the specific configuration
	InstanceID string  `url:"instanceId,omitempty"` //Instance Id of the specific configuration
	Offset     float64 `url:"offset,omitempty"`     //The number of Rest/Webhook Subscription detail's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of Rest/Webhook Subscription detail's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type GetSyslogSubscriptionDetailsQueryParams struct {
	Name       string  `url:"name,omitempty"`       //Name of the specific configuration
	InstanceID string  `url:"instanceId,omitempty"` //Instance Id of the specific configuration
	Offset     float64 `url:"offset,omitempty"`     //The number of Syslog Subscription detail's to offset in the resultset whose default value 0
	Limit      float64 `url:"limit,omitempty"`      //The number of Syslog Subscription detail's to limit in the resultset whose default value 10
	SortBy     string  `url:"sortBy,omitempty"`     //SortBy field name
	Order      string  `url:"order,omitempty"`      //order(asc/desc)
}
type CountOfEventSubscriptionsQueryParams struct {
	EventIDs string `url:"eventIds,omitempty"` //List of subscriptions related to the respective eventIds
}
type GetEmailEventSubscriptionsQueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //List of email subscriptions related to the respective eventIds (Comma separated event ids)
	Offset    float64 `url:"offset,omitempty"`    //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    //SortBy field name
	Order     string  `url:"order,omitempty"`     //order(asc/desc)
	Domain    string  `url:"domain,omitempty"`    //List of email subscriptions related to the respective domain
	SubDomain string  `url:"subDomain,omitempty"` //List of email subscriptions related to the respective sub-domain
	Category  string  `url:"category,omitempty"`  //List of email subscriptions related to the respective category
	Type      string  `url:"type,omitempty"`      //List of email subscriptions related to the respective type
	Name      string  `url:"name,omitempty"`      //List of email subscriptions related to the respective name
}
type GetRestWebhookEventSubscriptionsQueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //List of subscriptions related to the respective eventIds (Comma separated event ids)
	Offset    float64 `url:"offset,omitempty"`    //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    //SortBy field name
	Order     string  `url:"order,omitempty"`     //order(asc/desc)
	Domain    string  `url:"domain,omitempty"`    //List of subscriptions related to the respective domain
	SubDomain string  `url:"subDomain,omitempty"` //List of subscriptions related to the respective sub-domain
	Category  string  `url:"category,omitempty"`  //List of subscriptions related to the respective category
	Type      string  `url:"type,omitempty"`      //List of subscriptions related to the respective type
	Name      string  `url:"name,omitempty"`      //List of subscriptions related to the respective name
}
type GetSyslogEventSubscriptionsQueryParams struct {
	EventIDs  string  `url:"eventIds,omitempty"`  //List of subscriptions related to the respective eventIds (Comma separated event ids)
	Offset    float64 `url:"offset,omitempty"`    //The number of Subscriptions's to offset in the resultset whose default value 0
	Limit     float64 `url:"limit,omitempty"`     //The number of Subscriptions's to limit in the resultset whose default value 10
	SortBy    string  `url:"sortBy,omitempty"`    //SortBy field name
	Order     string  `url:"order,omitempty"`     //order(asc/desc)
	Domain    string  `url:"domain,omitempty"`    //List of subscriptions related to the respective domain
	SubDomain string  `url:"subDomain,omitempty"` //List of subscriptions related to the respective sub-domain
	Category  string  `url:"category,omitempty"`  //List of subscriptions related to the respective category
	Type      string  `url:"type,omitempty"`      //List of subscriptions related to the respective type
	Name      string  `url:"name,omitempty"`      //List of subscriptions related to the respective name
}
type GetEventsQueryParams struct {
	EventID string  `url:"eventId,omitempty"` //The registered EventId should be provided
	Tags    string  `url:"tags,omitempty"`    //The registered Tags should be provided
	Offset  float64 `url:"offset,omitempty"`  //The number of Registries to offset in the resultset whose default value 0
	Limit   float64 `url:"limit,omitempty"`   //The number of Registries to limit in the resultset whose default value 10
	SortBy  string  `url:"sortBy,omitempty"`  //SortBy field name
	Order   string  `url:"order,omitempty"`   //order(asc/desc)
}
type CountOfEventsQueryParams struct {
	EventID string `url:"eventId,omitempty"` //The registered EventId should be provided
	Tags    string `url:"tags,omitempty"`    //The registered Tags should be provided
}
type GetEventArtifactsQueryParams struct {
	EventIDs string  `url:"eventIds,omitempty"` //List of eventIds
	Tags     string  `url:"tags,omitempty"`     //Tags defined
	Offset   float64 `url:"offset,omitempty"`   //Record start offset
	Limit    float64 `url:"limit,omitempty"`    //# of records to return in result set
	SortBy   string  `url:"sortBy,omitempty"`   //Sort by field
	Order    string  `url:"order,omitempty"`    //sorting order (asc/desc)
	Search   string  `url:"search,omitempty"`   //findd matches in name, description, eventId, type, category
}

type ResponseEventManagementGetAuditLogParentRecords []ResponseItemEventManagementGetAuditLogParentRecords // Array of ResponseEventManagementGetAuditLogParentRecords
type ResponseItemEventManagementGetAuditLogParentRecords struct {
	Version           string                                                                `json:"version,omitempty"`           // Version
	InstanceID        string                                                                `json:"instanceId,omitempty"`        // Instance Id
	EventID           string                                                                `json:"eventId,omitempty"`           // Event Id
	Namespace         string                                                                `json:"namespace,omitempty"`         // Namespace
	Name              string                                                                `json:"name,omitempty"`              // Name
	Description       string                                                                `json:"description,omitempty"`       // Description
	Type              string                                                                `json:"type,omitempty"`              // Type
	Category          string                                                                `json:"category,omitempty"`          // Category
	Domain            string                                                                `json:"domain,omitempty"`            // Domain
	SubDomain         string                                                                `json:"subDomain,omitempty"`         // Sub Domain
	Severity          *int                                                                  `json:"severity,omitempty"`          // Severity
	Source            string                                                                `json:"source,omitempty"`            // Source
	Timestamp         *int                                                                  `json:"timestamp,omitempty"`         // Timestamp
	Tags              *[]ResponseItemEventManagementGetAuditLogParentRecordsTags            `json:"tags,omitempty"`              // Tags
	Details           *ResponseItemEventManagementGetAuditLogParentRecordsDetails           `json:"details,omitempty"`           // Details
	CiscoDnaEventLink *ResponseItemEventManagementGetAuditLogParentRecordsCiscoDnaEventLink `json:"ciscoDnaEventLink,omitempty"` // Cisco Dna Event Link
	Note              *ResponseItemEventManagementGetAuditLogParentRecordsNote              `json:"note,omitempty"`              // Note
	TntID             string                                                                `json:"tntId,omitempty"`             // Tnt Id
	Context           string                                                                `json:"context,omitempty"`           // Context
	UserID            string                                                                `json:"userId,omitempty"`            // User Id
	I18N              string                                                                `json:"i18n,omitempty"`              // I18n
	EventHierarchy    *ResponseItemEventManagementGetAuditLogParentRecordsEventHierarchy    `json:"eventHierarchy,omitempty"`    // Event Hierarchy
	Message           string                                                                `json:"message,omitempty"`           // Message
	MessageParams     *ResponseItemEventManagementGetAuditLogParentRecordsMessageParams     `json:"messageParams,omitempty"`     // Message Params
	AdditionalDetails *ResponseItemEventManagementGetAuditLogParentRecordsAdditionalDetails `json:"additionalDetails,omitempty"` // Additional Details
	ParentInstanceID  *ResponseItemEventManagementGetAuditLogParentRecordsParentInstanceID  `json:"parentInstanceId,omitempty"`  // Parent Instance Id
	Network           *ResponseItemEventManagementGetAuditLogParentRecordsNetwork           `json:"network,omitempty"`           // Network
	ChildCount        *float64                                                              `json:"childCount,omitempty"`        // Child Count
	TenantID          string                                                                `json:"tenantId,omitempty"`          // Tenant Id
}
type ResponseItemEventManagementGetAuditLogParentRecordsTags interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsDetails interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsCiscoDnaEventLink interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsNote interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsEventHierarchy interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsMessageParams interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsAdditionalDetails interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsParentInstanceID interface{}
type ResponseItemEventManagementGetAuditLogParentRecordsNetwork interface{}
type ResponseEventManagementGetAuditLogSummary []ResponseItemEventManagementGetAuditLogSummary // Array of ResponseEventManagementGetAuditLogSummary
type ResponseItemEventManagementGetAuditLogSummary struct {
	Count        *int `json:"count,omitempty"`        // Count
	MaxTimestamp *int `json:"maxTimestamp,omitempty"` // Max Timestamp
	MinTimestamp *int `json:"minTimestamp,omitempty"` // Min Timestamp
}
type ResponseEventManagementGetAuditLogRecords []ResponseItemEventManagementGetAuditLogRecords // Array of ResponseEventManagementGetAuditLogRecords
type ResponseItemEventManagementGetAuditLogRecords struct {
	Version           string                                                          `json:"version,omitempty"`           // Version
	InstanceID        string                                                          `json:"instanceId,omitempty"`        // Instance Id
	EventID           string                                                          `json:"eventId,omitempty"`           // Event Id
	Namespace         string                                                          `json:"namespace,omitempty"`         // Namespace
	Name              string                                                          `json:"name,omitempty"`              // Name
	Description       string                                                          `json:"description,omitempty"`       // Description
	Type              string                                                          `json:"type,omitempty"`              // Type
	Category          string                                                          `json:"category,omitempty"`          // Category
	Domain            string                                                          `json:"domain,omitempty"`            // Domain
	SubDomain         string                                                          `json:"subDomain,omitempty"`         // Sub Domain
	Severity          *int                                                            `json:"severity,omitempty"`          // Severity
	Source            string                                                          `json:"source,omitempty"`            // Source
	Timestamp         *int                                                            `json:"timestamp,omitempty"`         // Timestamp
	Tags              *[]ResponseItemEventManagementGetAuditLogRecordsTags            `json:"tags,omitempty"`              // Tags
	Details           *ResponseItemEventManagementGetAuditLogRecordsDetails           `json:"details,omitempty"`           // Details
	CiscoDnaEventLink string                                                          `json:"ciscoDnaEventLink,omitempty"` // Cisco Dna Event Link
	Note              string                                                          `json:"note,omitempty"`              // Note
	TntID             string                                                          `json:"tntId,omitempty"`             // Tnt Id
	Context           string                                                          `json:"context,omitempty"`           // Context
	UserID            string                                                          `json:"userId,omitempty"`            // User Id
	I18N              string                                                          `json:"i18n,omitempty"`              // I18n
	EventHierarchy    string                                                          `json:"eventHierarchy,omitempty"`    // Event Hierarchy
	Message           string                                                          `json:"message,omitempty"`           // Message
	MessageParams     string                                                          `json:"messageParams,omitempty"`     // Message Params
	AdditionalDetails *ResponseItemEventManagementGetAuditLogRecordsAdditionalDetails `json:"additionalDetails,omitempty"` // Additional Details
	ParentInstanceID  string                                                          `json:"parentInstanceId,omitempty"`  // Parent Instance Id
	Network           string                                                          `json:"network,omitempty"`           // Network
	ChildCount        *float64                                                        `json:"childCount,omitempty"`        // Child Count
	TenantID          string                                                          `json:"tenantId,omitempty"`          // Tenant Id
}
type ResponseItemEventManagementGetAuditLogRecordsTags interface{}
type ResponseItemEventManagementGetAuditLogRecordsDetails interface{}
type ResponseItemEventManagementGetAuditLogRecordsAdditionalDetails interface{}
type ResponseEventManagementGetStatusAPIForEvents struct {
	ErrorMessage  *ResponseEventManagementGetStatusAPIForEventsErrorMessage `json:"errorMessage,omitempty"`  // Error Message
	APIStatus     string                                                    `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                    `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementGetStatusAPIForEventsErrorMessage interface{}
type ResponseEventManagementUpdateEmailDestination struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementCreateEmailDestination struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetNotifications []ResponseItemEventManagementGetNotifications // Array of ResponseEventManagementGetNotifications
type ResponseItemEventManagementGetNotifications struct {
	EventID        string                                              `json:"eventId,omitempty"`        // Event Id
	InstanceID     string                                              `json:"instanceId,omitempty"`     // Instance Id
	Namespace      string                                              `json:"namespace,omitempty"`      // Namespace
	Name           string                                              `json:"name,omitempty"`           // Name
	Description    string                                              `json:"description,omitempty"`    // Description
	Version        string                                              `json:"version,omitempty"`        // Version
	Category       string                                              `json:"category,omitempty"`       // Category
	Domain         string                                              `json:"domain,omitempty"`         // Domain
	SubDomain      string                                              `json:"subDomain,omitempty"`      // Sub Domain
	Type           string                                              `json:"type,omitempty"`           // Type
	Severity       string                                              `json:"severity,omitempty"`       // Severity
	Source         string                                              `json:"source,omitempty"`         // Source
	Timestamp      string                                              `json:"timestamp,omitempty"`      // Timestamp
	Details        string                                              `json:"details,omitempty"`        // Details
	EventHierarchy string                                              `json:"eventHierarchy,omitempty"` // Event Hierarchy
	Network        *ResponseItemEventManagementGetNotificationsNetwork `json:"network,omitempty"`        //
}
type ResponseItemEventManagementGetNotificationsNetwork struct {
	SiteID   string `json:"siteId,omitempty"`   // Site Id
	DeviceID string `json:"deviceId,omitempty"` // Device Id
}
type ResponseEventManagementCountOfNotifications struct {
	Response string `json:"response,omitempty"` // Response
}
type ResponseEventManagementGetEventSubscriptions []ResponseItemEventManagementGetEventSubscriptions // Array of ResponseEventManagementGetEventSubscriptions
type ResponseItemEventManagementGetEventSubscriptions struct {
	Version               string                                                                   `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                   `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                   `json:"name,omitempty"`                  // Name
	Description           string                                                                   `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetEventSubscriptionsFilter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                    `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                   `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpoints struct {
	InstanceID          string                                                                                    `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                    `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType  string                                                                                                 `json:"connectorType,omitempty"`  // Connector Type
	InstanceID     string                                                                                                 `json:"instanceId,omitempty"`     // Instance Id
	Name           string                                                                                                 `json:"name,omitempty"`           // Name
	Description    string                                                                                                 `json:"description,omitempty"`    // Description
	URL            string                                                                                                 `json:"url,omitempty"`            // Url
	BasePath       string                                                                                                 `json:"basePath,omitempty"`       // Base Path
	Resource       string                                                                                                 `json:"resource,omitempty"`       // Resource
	Method         string                                                                                                 `json:"method,omitempty"`         // Method
	TrustCert      string                                                                                                 `json:"trustCert,omitempty"`      // Trust Cert
	Headers        *[]ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsHeaders     `json:"headers,omitempty"`        //
	QueryParams    *[]ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsQueryParams `json:"queryParams,omitempty"`    //
	PathParams     *[]ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsPathParams  `json:"pathParams,omitempty"`     //
	Body           string                                                                                                 `json:"body,omitempty"`           // Body
	ConnectTimeout string                                                                                                 `json:"connectTimeout,omitempty"` // Connect Timeout
	ReadTimeout    string                                                                                                 `json:"readTimeout,omitempty"`    // Read Timeout
}
type ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsHeaders struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsQueryParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsPathParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetEventSubscriptionsFilter struct {
	EventIDs          []string                                                                   `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                   `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetEventSubscriptionsFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                   `json:"types,omitempty"`             // Types
	Categories        []string                                                                   `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                   `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                   `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                   `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetEventSubscriptionsFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseEventManagementDeleteEventSubscriptions struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementUpdateEventSubscriptions struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementCreateEventSubscriptions struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetEmailSubscriptionDetails []ResponseItemEventManagementGetEmailSubscriptionDetails // Array of ResponseEventManagementGetEmailSubscriptionDetails
type ResponseItemEventManagementGetEmailSubscriptionDetails struct {
	InstanceID       string   `json:"instanceId,omitempty"`       // Instance Id
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // From Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // To Email Addresses
	Subject          string   `json:"subject,omitempty"`          // Subject
}
type ResponseEventManagementGetRestWebhookSubscriptionDetails []ResponseItemEventManagementGetRestWebhookSubscriptionDetails // Array of ResponseEventManagementGetRestWebhookSubscriptionDetails
type ResponseItemEventManagementGetRestWebhookSubscriptionDetails struct {
	InstanceID    string                                                                 `json:"instanceId,omitempty"`    // Instance Id
	Name          string                                                                 `json:"name,omitempty"`          // Name
	Description   string                                                                 `json:"description,omitempty"`   // Description
	ConnectorType string                                                                 `json:"connectorType,omitempty"` // Connector Type
	URL           string                                                                 `json:"url,omitempty"`           // Url
	Method        string                                                                 `json:"method,omitempty"`        // Method
	TrustCert     string                                                                 `json:"trustCert,omitempty"`     // Trust Cert
	Headers       *[]ResponseItemEventManagementGetRestWebhookSubscriptionDetailsHeaders `json:"headers,omitempty"`       //
	QueryParams   []string                                                               `json:"queryParams,omitempty"`   // Query Params
	PathParams    []string                                                               `json:"pathParams,omitempty"`    // Path Params
}
type ResponseItemEventManagementGetRestWebhookSubscriptionDetailsHeaders struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseEventManagementGetSyslogSubscriptionDetails []ResponseItemEventManagementGetSyslogSubscriptionDetails // Array of ResponseEventManagementGetSyslogSubscriptionDetails
type ResponseItemEventManagementGetSyslogSubscriptionDetails struct {
	InstanceID    string                                                               `json:"instanceId,omitempty"`    // Instance Id
	Name          string                                                               `json:"name,omitempty"`          // Name
	Description   string                                                               `json:"description,omitempty"`   // Description
	ConnectorType string                                                               `json:"connectorType,omitempty"` // Connector Type
	SyslogConfig  *ResponseItemEventManagementGetSyslogSubscriptionDetailsSyslogConfig `json:"syslogConfig,omitempty"`  //
}
type ResponseItemEventManagementGetSyslogSubscriptionDetailsSyslogConfig struct {
	ConfigID    string `json:"configId,omitempty"`    // Config Id
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Port        string `json:"port,omitempty"`        // Port
	Protocol    string `json:"protocol,omitempty"`    // Protocol
}
type ResponseEventManagementCountOfEventSubscriptions struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseEventManagementCreateEmailEventSubscription struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementUpdateEmailEventSubscription struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetEmailEventSubscriptions []ResponseItemEventManagementGetEmailEventSubscriptions // Array of ResponseEventManagementGetEmailEventSubscriptions
type ResponseItemEventManagementGetEmailEventSubscriptions struct {
	Version               string                                                                        `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                        `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                        `json:"name,omitempty"`                  // Name
	Description           string                                                                        `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetEmailEventSubscriptionsSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetEmailEventSubscriptionsFilter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                         `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                        `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetEmailEventSubscriptionsSubscriptionEndpoints struct {
	InstanceID          string                                                                                         `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetEmailEventSubscriptionsSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                         `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetEmailEventSubscriptionsSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type
	InstanceID       string   `json:"instanceId,omitempty"`       // Instance Id
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // From Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // To Email Addresses
	Subject          string   `json:"subject,omitempty"`          // Subject
}
type ResponseItemEventManagementGetEmailEventSubscriptionsFilter struct {
	EventIDs          []string                                                                        `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                        `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetEmailEventSubscriptionsFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                        `json:"types,omitempty"`             // Types
	Categories        []string                                                                        `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                        `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                        `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                        `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetEmailEventSubscriptionsFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseEventManagementCreateRestWebhookEventSubscription struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetRestWebhookEventSubscriptions []ResponseItemEventManagementGetRestWebhookEventSubscriptions // Array of ResponseEventManagementGetRestWebhookEventSubscriptions
type ResponseItemEventManagementGetRestWebhookEventSubscriptions struct {
	Version               string                                                                              `json:"version,omitempty"`               // Version
	SubscriptionID        string                                                                              `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                              `json:"name,omitempty"`                  // Name
	Description           string                                                                              `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetRestWebhookEventSubscriptionsFilter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                               `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                              `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpoints struct {
	InstanceID          string                                                                                               `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                               `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType  string                                                                                                            `json:"connectorType,omitempty"`  // Connector Type
	InstanceID     string                                                                                                            `json:"instanceId,omitempty"`     // Instance Id
	Name           string                                                                                                            `json:"name,omitempty"`           // Name
	Description    string                                                                                                            `json:"description,omitempty"`    // Description
	URL            string                                                                                                            `json:"url,omitempty"`            // Url
	BasePath       string                                                                                                            `json:"basePath,omitempty"`       // Base Path
	Resource       string                                                                                                            `json:"resource,omitempty"`       // Resource
	Method         string                                                                                                            `json:"method,omitempty"`         // Method
	TrustCert      string                                                                                                            `json:"trustCert,omitempty"`      // Trust Cert
	Headers        *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsHeaders     `json:"headers,omitempty"`        //
	QueryParams    *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsQueryParams `json:"queryParams,omitempty"`    //
	PathParams     *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsPathParams  `json:"pathParams,omitempty"`     //
	Body           string                                                                                                            `json:"body,omitempty"`           // Body
	ConnectTimeout string                                                                                                            `json:"connectTimeout,omitempty"` // Connect Timeout
	ReadTimeout    string                                                                                                            `json:"readTimeout,omitempty"`    // Read Timeout
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsHeaders struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsQueryParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsPathParams struct {
	String string `json:"string,omitempty"` // String
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsFilter struct {
	EventIDs          []string                                                                              `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                              `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetRestWebhookEventSubscriptionsFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                              `json:"types,omitempty"`             // Types
	Categories        []string                                                                              `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                              `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                              `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                              `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetRestWebhookEventSubscriptionsFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseEventManagementUpdateRestWebhookEventSubscription struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementUpdateSyslogEventSubscription struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementCreateSyslogEventSubscription struct {
	StatusURI string `json:"statusUri,omitempty"` // Status Uri
}
type ResponseEventManagementGetSyslogEventSubscriptions []ResponseItemEventManagementGetSyslogEventSubscriptions // Array of ResponseEventManagementGetSyslogEventSubscriptions
type ResponseItemEventManagementGetSyslogEventSubscriptions struct {
	Version               *ResponseItemEventManagementGetSyslogEventSubscriptionsVersion                 `json:"version,omitempty"`               // Version
	SubscriptionID        *ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionID          `json:"subscriptionId,omitempty"`        // Subscription Id
	Name                  string                                                                         `json:"name,omitempty"`                  // Name
	Description           string                                                                         `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *ResponseItemEventManagementGetSyslogEventSubscriptionsFilter                  `json:"filter,omitempty"`                //
	IsPrivate             *bool                                                                          `json:"isPrivate,omitempty"`             // Is Private
	TenantID              string                                                                         `json:"tenantId,omitempty"`              // Tenant Id
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsVersion interface{}
type ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionID interface{}
type ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpoints struct {
	InstanceID          string                                                                                          `json:"instanceId,omitempty"`          // Instance Id
	SubscriptionDetails *ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
	ConnectorType       string                                                                                          `json:"connectorType,omitempty"`       // Connector Type
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string                                                                                                      `json:"connectorType,omitempty"` // Connector Type
	InstanceID    string                                                                                                      `json:"instanceId,omitempty"`    // Instance Id
	Name          string                                                                                                      `json:"name,omitempty"`          // Name
	Description   string                                                                                                      `json:"description,omitempty"`   // Description
	SyslogConfig  *ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsSyslogConfig `json:"syslogConfig,omitempty"`  //
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsSyslogConfig struct {
	Version     string `json:"version,omitempty"`     // Version
	TenantID    string `json:"tenantId,omitempty"`    // Tenant Id
	ConfigID    string `json:"configId,omitempty"`    // Config Id
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Port        *int   `json:"port,omitempty"`        // Port
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsFilter struct {
	EventIDs          []string                                                                         `json:"eventIds,omitempty"`          // Event Ids
	Others            []string                                                                         `json:"others,omitempty"`            // Others
	DomainsSubdomains *[]ResponseItemEventManagementGetSyslogEventSubscriptionsFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                         `json:"types,omitempty"`             // Types
	Categories        []string                                                                         `json:"categories,omitempty"`        // Categories
	Severities        *[]ResponseItemEventManagementGetSyslogEventSubscriptionsFilterSeverities        `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                         `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                         `json:"siteIds,omitempty"`           // Site Ids
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type ResponseItemEventManagementGetSyslogEventSubscriptionsFilterSeverities interface{}
type ResponseEventManagementUpdateSyslogDestination struct {
	ErrorMessage  *ResponseEventManagementUpdateSyslogDestinationErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                      `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                      `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementUpdateSyslogDestinationErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementCreateSyslogDestination struct {
	ErrorMessage  *ResponseEventManagementCreateSyslogDestinationErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                      `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                      `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementCreateSyslogDestinationErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementCreateWebhookDestination struct {
	ErrorMessage  *ResponseEventManagementCreateWebhookDestinationErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                       `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                       `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementCreateWebhookDestinationErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementUpdateWebhookDestination struct {
	ErrorMessage  *ResponseEventManagementUpdateWebhookDestinationErrorMessage `json:"errorMessage,omitempty"`  //
	APIStatus     string                                                       `json:"apiStatus,omitempty"`     // Api Status
	StatusMessage string                                                       `json:"statusMessage,omitempty"` // Status Message
}
type ResponseEventManagementUpdateWebhookDestinationErrorMessage struct {
	Errors []string `json:"errors,omitempty"` // Errors
}
type ResponseEventManagementGetEvents []ResponseItemEventManagementGetEvents // Array of ResponseEventManagementGetEvents
type ResponseItemEventManagementGetEvents struct {
	EventID           string                                       `json:"eventId,omitempty"`           // Event Id
	NameSpace         string                                       `json:"nameSpace,omitempty"`         // Name Space
	Name              string                                       `json:"name,omitempty"`              // Name
	Description       string                                       `json:"description,omitempty"`       // Description
	Version           string                                       `json:"version,omitempty"`           // Version
	Category          string                                       `json:"category,omitempty"`          // Category
	Domain            string                                       `json:"domain,omitempty"`            // Domain
	SubDomain         string                                       `json:"subDomain,omitempty"`         // Sub Domain
	Type              string                                       `json:"type,omitempty"`              // Type
	Tags              []string                                     `json:"tags,omitempty"`              // Tags
	Severity          *float64                                     `json:"severity,omitempty"`          // Severity
	Details           *ResponseItemEventManagementGetEventsDetails `json:"details,omitempty"`           // Details
	SubscriptionTypes []string                                     `json:"subscriptionTypes,omitempty"` // Subscription Types
}
type ResponseItemEventManagementGetEventsDetails interface{}
type ResponseEventManagementCountOfEvents struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseEventManagementGetEventArtifacts []ResponseItemEventManagementGetEventArtifacts // Array of ResponseEventManagementGetEventArtifacts
type ResponseItemEventManagementGetEventArtifacts struct {
	Version                 string                                                        `json:"version,omitempty"`                 // Version
	ArtifactID              string                                                        `json:"artifactId,omitempty"`              // Artifact Id
	Namespace               string                                                        `json:"namespace,omitempty"`               // Namespace
	Name                    string                                                        `json:"name,omitempty"`                    // Name
	Description             string                                                        `json:"description,omitempty"`             // Description
	Domain                  string                                                        `json:"domain,omitempty"`                  // Domain
	SubDomain               string                                                        `json:"subDomain,omitempty"`               // Sub Domain
	Tags                    []string                                                      `json:"tags,omitempty"`                    // Tags
	IsTemplateEnabled       *bool                                                         `json:"isTemplateEnabled,omitempty"`       // Is Template Enabled
	CiscoDnaEventLink       string                                                        `json:"ciscoDNAEventLink,omitempty"`       // Cisco D N A Event Link
	Note                    string                                                        `json:"note,omitempty"`                    // Note
	IsPrivate               *bool                                                         `json:"isPrivate,omitempty"`               // Is Private
	EventPayload            *ResponseItemEventManagementGetEventArtifactsEventPayload     `json:"eventPayload,omitempty"`            //
	EventTemplates          *[]ResponseItemEventManagementGetEventArtifactsEventTemplates `json:"eventTemplates,omitempty"`          // Event Templates
	IsTenantAware           *bool                                                         `json:"isTenantAware,omitempty"`           // Is Tenant Aware
	SupportedConnectorTypes []string                                                      `json:"supportedConnectorTypes,omitempty"` // Supported Connector Types
	TenantID                string                                                        `json:"tenantId,omitempty"`                // Tenant Id
}
type ResponseItemEventManagementGetEventArtifactsEventPayload struct {
	EventID           string                                                                     `json:"eventId,omitempty"`           // Event Id
	Version           string                                                                     `json:"version,omitempty"`           // Version
	Category          string                                                                     `json:"category,omitempty"`          // Category
	Type              string                                                                     `json:"type,omitempty"`              // Type
	Source            string                                                                     `json:"source,omitempty"`            // Source
	Severity          string                                                                     `json:"severity,omitempty"`          // Severity
	Details           *ResponseItemEventManagementGetEventArtifactsEventPayloadDetails           `json:"details,omitempty"`           //
	AdditionalDetails *ResponseItemEventManagementGetEventArtifactsEventPayloadAdditionalDetails `json:"additionalDetails,omitempty"` // Additional Details
}
type ResponseItemEventManagementGetEventArtifactsEventPayloadDetails struct {
	DeviceIP string `json:"device_ip,omitempty"` // Device Ip
	Message  string `json:"message,omitempty"`   // Message
}
type ResponseItemEventManagementGetEventArtifactsEventPayloadAdditionalDetails interface{}
type ResponseItemEventManagementGetEventArtifactsEventTemplates interface{}
type ResponseEventManagementEventArtifactCount struct {
	Response *float64 `json:"response,omitempty"` // Response
}
type ResponseEventManagementGetConnectorTypes []ResponseItemEventManagementGetConnectorTypes // Array of ResponseEventManagementGetConnectorTypes
type ResponseItemEventManagementGetConnectorTypes struct {
	ConnectorType      string `json:"connectorType,omitempty"`      // Connector Type
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	IsDefaultSupported *bool  `json:"isDefaultSupported,omitempty"` // Is Default Supported
	IsCustomConnector  *bool  `json:"isCustomConnector,omitempty"`  // Is Custom Connector
}
type RequestEventManagementUpdateEmailDestination struct {
	EmailConfigID       string                                                           `json:"emailConfigId,omitempty"`       // Required only for update email configuration
	PrimarySmtpConfig   *RequestEventManagementUpdateEmailDestinationPrimarySmtpConfig   `json:"primarySMTPConfig,omitempty"`   //
	SecondarySmtpConfig *RequestEventManagementUpdateEmailDestinationSecondarySmtpConfig `json:"secondarySMTPConfig,omitempty"` //
	FromEmail           string                                                           `json:"fromEmail,omitempty"`           // From Email
	ToEmail             string                                                           `json:"toEmail,omitempty"`             // To Email
	Subject             string                                                           `json:"subject,omitempty"`             // Subject
}
type RequestEventManagementUpdateEmailDestinationPrimarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
}
type RequestEventManagementUpdateEmailDestinationSecondarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
}
type RequestEventManagementCreateEmailDestination struct {
	EmailConfigID       string                                                           `json:"emailConfigId,omitempty"`       // Required only for update email configuration
	PrimarySmtpConfig   *RequestEventManagementCreateEmailDestinationPrimarySmtpConfig   `json:"primarySMTPConfig,omitempty"`   //
	SecondarySmtpConfig *RequestEventManagementCreateEmailDestinationSecondarySmtpConfig `json:"secondarySMTPConfig,omitempty"` //
	FromEmail           string                                                           `json:"fromEmail,omitempty"`           // From Email
	ToEmail             string                                                           `json:"toEmail,omitempty"`             // To Email
	Subject             string                                                           `json:"subject,omitempty"`             // Subject
}
type RequestEventManagementCreateEmailDestinationPrimarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
}
type RequestEventManagementCreateEmailDestinationSecondarySmtpConfig struct {
	HostName string `json:"hostName,omitempty"` // Host Name
	Port     string `json:"port,omitempty"`     // Port
	UserName string `json:"userName,omitempty"` // User Name
	Password string `json:"password,omitempty"` // Password
}
type RequestEventManagementUpdateEventSubscriptions []RequestItemEventManagementUpdateEventSubscriptions // Array of RequestEventManagementUpdateEventSubscriptions
type RequestItemEventManagementUpdateEventSubscriptions struct {
	SubscriptionID        string                                                                     `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                     `json:"version,omitempty"`               // Version
	Name                  string                                                                     `json:"name,omitempty"`                  // Name
	Description           string                                                                     `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateEventSubscriptionsFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpoints struct {
	InstanceID          string                                                                                      `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementUpdateEventSubscriptionsFilter struct {
	EventIDs          []string                                                                     `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementUpdateEventSubscriptionsFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                     `json:"types,omitempty"`             // Types
	Categories        []string                                                                     `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                     `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                     `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                     `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateEventSubscriptionsFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateEventSubscriptions []RequestItemEventManagementCreateEventSubscriptions // Array of RequestEventManagementCreateEventSubscriptions
type RequestItemEventManagementCreateEventSubscriptions struct {
	SubscriptionID        string                                                                     `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                     `json:"version,omitempty"`               // Version
	Name                  string                                                                     `json:"name,omitempty"`                  // Name
	Description           string                                                                     `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateEventSubscriptionsFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpoints struct {
	InstanceID          string                                                                                      `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateEventSubscriptionsSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementCreateEventSubscriptionsFilter struct {
	EventIDs          []string                                                                     `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementCreateEventSubscriptionsFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                     `json:"types,omitempty"`             // Types
	Categories        []string                                                                     `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                     `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                     `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                     `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateEventSubscriptionsFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateEmailEventSubscription []RequestItemEventManagementCreateEmailEventSubscription // Array of RequestEventManagementCreateEmailEventSubscription
type RequestItemEventManagementCreateEmailEventSubscription struct {
	SubscriptionID        string                                                                         `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                         `json:"version,omitempty"`               // Version
	Name                  string                                                                         `json:"name,omitempty"`                  // Name
	Description           string                                                                         `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateEmailEventSubscriptionFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpoints struct {
	InstanceID          string                                                                                          `json:"instanceId,omitempty"`          // (From Get Email Subscription Details --> pick InstanceId)
	SubscriptionDetails *RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type (Must be EMAIL)
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // Senders Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // Recipient's Email Addresses (Comma separated)
	Subject          string   `json:"subject,omitempty"`          // Email Subject
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
}
type RequestItemEventManagementCreateEmailEventSubscriptionFilter struct {
	EventIDs          []string                                                                         `json:"eventIds,omitempty"`          // Event Ids
	DomainsSubdomains *[]RequestItemEventManagementCreateEmailEventSubscriptionFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                         `json:"types,omitempty"`             // Types
	Categories        []string                                                                         `json:"categories,omitempty"`        // Categories
	Severities        *[]int                                                                           `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                         `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                         `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateEmailEventSubscriptionFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateEmailEventSubscription []RequestItemEventManagementUpdateEmailEventSubscription // Array of RequestEventManagementUpdateEmailEventSubscription
type RequestItemEventManagementUpdateEmailEventSubscription struct {
	SubscriptionID        string                                                                         `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                         `json:"version,omitempty"`               // Version
	Name                  string                                                                         `json:"name,omitempty"`                  // Name
	Description           string                                                                         `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateEmailEventSubscriptionFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpoints struct {
	InstanceID          string                                                                                          `json:"instanceId,omitempty"`          // (From Get Email Subscription Details --> pick InstanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateEmailEventSubscriptionSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType    string   `json:"connectorType,omitempty"`    // Connector Type (Must be EMAIL)
	FromEmailAddress string   `json:"fromEmailAddress,omitempty"` // Senders Email Address
	ToEmailAddresses []string `json:"toEmailAddresses,omitempty"` // Recipient's Email Addresses (Comma separated)
	Subject          string   `json:"subject,omitempty"`          // Email Subject
	Name             string   `json:"name,omitempty"`             // Name
	Description      string   `json:"description,omitempty"`      // Description
}
type RequestItemEventManagementUpdateEmailEventSubscriptionFilter struct {
	EventIDs          []string                                                                         `json:"eventIds,omitempty"`          // Event Ids
	DomainsSubdomains *[]RequestItemEventManagementUpdateEmailEventSubscriptionFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                         `json:"types,omitempty"`             // Types
	Categories        []string                                                                         `json:"categories,omitempty"`        // Categories
	Severities        *[]int                                                                           `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                         `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                         `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateEmailEventSubscriptionFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateRestWebhookEventSubscription []RequestItemEventManagementCreateRestWebhookEventSubscription // Array of RequestEventManagementCreateRestWebhookEventSubscription
type RequestItemEventManagementCreateRestWebhookEventSubscription struct {
	SubscriptionID        string                                                                               `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                               `json:"version,omitempty"`               // Version
	Name                  string                                                                               `json:"name,omitempty"`                  // Name
	Description           string                                                                               `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints struct {
	InstanceID          string                                                                                                `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter struct {
	EventIDs          []string                                                                               `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementCreateRestWebhookEventSubscriptionFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                               `json:"types,omitempty"`             // Types
	Categories        []string                                                                               `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                               `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                               `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                               `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateRestWebhookEventSubscriptionFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateRestWebhookEventSubscription []RequestItemEventManagementUpdateRestWebhookEventSubscription // Array of RequestEventManagementUpdateRestWebhookEventSubscription
type RequestItemEventManagementUpdateRestWebhookEventSubscription struct {
	SubscriptionID        string                                                                               `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                               `json:"version,omitempty"`               // Version
	Name                  string                                                                               `json:"name,omitempty"`                  // Name
	Description           string                                                                               `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints struct {
	InstanceID          string                                                                                                `json:"instanceId,omitempty"`          // (From 	Get Rest/Webhook Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be REST)
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter struct {
	EventIDs          []string                                                                               `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                               `json:"types,omitempty"`             // Types
	Categories        []string                                                                               `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                               `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                               `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                               `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateSyslogEventSubscription []RequestItemEventManagementUpdateSyslogEventSubscription // Array of RequestEventManagementUpdateSyslogEventSubscription
type RequestItemEventManagementUpdateSyslogEventSubscription struct {
	SubscriptionID        string                                                                          `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                          `json:"version,omitempty"`               // Version
	Name                  string                                                                          `json:"name,omitempty"`                  // Name
	Description           string                                                                          `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementUpdateSyslogEventSubscriptionFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpoints struct {
	InstanceID          string                                                                                           `json:"instanceId,omitempty"`          // (From Get Syslog Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be SYSLOG)
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionFilter struct {
	EventIDs          []string                                                                          `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementUpdateSyslogEventSubscriptionFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                          `json:"types,omitempty"`             // Types
	Categories        []string                                                                          `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                          `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                          `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                          `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementUpdateSyslogEventSubscriptionFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementCreateSyslogEventSubscription []RequestItemEventManagementCreateSyslogEventSubscription // Array of RequestEventManagementCreateSyslogEventSubscription
type RequestItemEventManagementCreateSyslogEventSubscription struct {
	SubscriptionID        string                                                                          `json:"subscriptionId,omitempty"`        // Subscription Id (Unique UUID)
	Version               string                                                                          `json:"version,omitempty"`               // Version
	Name                  string                                                                          `json:"name,omitempty"`                  // Name
	Description           string                                                                          `json:"description,omitempty"`           // Description
	SubscriptionEndpoints *[]RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpoints `json:"subscriptionEndpoints,omitempty"` //
	Filter                *RequestItemEventManagementCreateSyslogEventSubscriptionFilter                  `json:"filter,omitempty"`                //
}
type RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpoints struct {
	InstanceID          string                                                                                           `json:"instanceId,omitempty"`          // (From Get Syslog Subscription Details --> pick instanceId)
	SubscriptionDetails *RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails `json:"subscriptionDetails,omitempty"` //
}
type RequestItemEventManagementCreateSyslogEventSubscriptionSubscriptionEndpointsSubscriptionDetails struct {
	ConnectorType string `json:"connectorType,omitempty"` // Connector Type (Must be SYSLOG)
}
type RequestItemEventManagementCreateSyslogEventSubscriptionFilter struct {
	EventIDs          []string                                                                          `json:"eventIds,omitempty"`          // Event Ids (Comma separated event ids)
	DomainsSubdomains *[]RequestItemEventManagementCreateSyslogEventSubscriptionFilterDomainsSubdomains `json:"domainsSubdomains,omitempty"` //
	Types             []string                                                                          `json:"types,omitempty"`             // Types
	Categories        []string                                                                          `json:"categories,omitempty"`        // Categories
	Severities        []string                                                                          `json:"severities,omitempty"`        // Severities
	Sources           []string                                                                          `json:"sources,omitempty"`           // Sources
	SiteIDs           []string                                                                          `json:"siteIds,omitempty"`           // Site Ids
}
type RequestItemEventManagementCreateSyslogEventSubscriptionFilterDomainsSubdomains struct {
	Domain     string   `json:"domain,omitempty"`     // Domain
	SubDomains []string `json:"subDomains,omitempty"` // Sub Domains
}
type RequestEventManagementUpdateSyslogDestination struct {
	ConfigID    string `json:"configId,omitempty"`    // Required only for update syslog configuration
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	Port        string `json:"port,omitempty"`        // Port
}
type RequestEventManagementCreateSyslogDestination struct {
	ConfigID    string `json:"configId,omitempty"`    // Required only for update syslog configuration
	Name        string `json:"name,omitempty"`        // Name
	Description string `json:"description,omitempty"` // Description
	Host        string `json:"host,omitempty"`        // Host
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	Port        string `json:"port,omitempty"`        // Port
}
type RequestEventManagementCreateWebhookDestination struct {
	WebhookID   string                                                   `json:"webhookId,omitempty"`   // Required only for update webhook configuration
	Name        string                                                   `json:"name,omitempty"`        // Name
	Description string                                                   `json:"description,omitempty"` // Description
	URL         string                                                   `json:"url,omitempty"`         // Url
	Method      string                                                   `json:"method,omitempty"`      // Method
	TrustCert   *bool                                                    `json:"trustCert,omitempty"`   // Trust Cert
	Headers     *[]RequestEventManagementCreateWebhookDestinationHeaders `json:"headers,omitempty"`     //
}
type RequestEventManagementCreateWebhookDestinationHeaders struct {
	Name         string `json:"name,omitempty"`         // Name
	Value        string `json:"value,omitempty"`        // Value
	DefaultValue string `json:"defaultValue,omitempty"` // Default Value
	Encrypt      *bool  `json:"encrypt,omitempty"`      // Encrypt
}
type RequestEventManagementUpdateWebhookDestination struct {
	WebhookID   string                                                   `json:"webhookId,omitempty"`   // Required only for update webhook configuration
	Name        string                                                   `json:"name,omitempty"`        // Name
	Description string                                                   `json:"description,omitempty"` // Description
	URL         string                                                   `json:"url,omitempty"`         // Url
	Method      string                                                   `json:"method,omitempty"`      // Method
	TrustCert   *bool                                                    `json:"trustCert,omitempty"`   // Trust Cert
	Headers     *[]RequestEventManagementUpdateWebhookDestinationHeaders `json:"headers,omitempty"`     //
}
type RequestEventManagementUpdateWebhookDestinationHeaders struct {
	Name         string `json:"name,omitempty"`         // Name
	Value        string `json:"value,omitempty"`        // Value
	DefaultValue string `json:"defaultValue,omitempty"` // Default Value
	Encrypt      *bool  `json:"encrypt,omitempty"`      // Encrypt
}

//GetAuditLogParentRecords Get AuditLog Parent Records - 9590-7ae9-46ea-b1c6
/* Get Parent Audit Log Event instances from the Event-Hub


@param GetAuditLogParentRecordsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetAuditLogParentRecords(GetAuditLogParentRecordsQueryParams *GetAuditLogParentRecordsQueryParams) (*ResponseEventManagementGetAuditLogParentRecords, *resty.Response, error) {
	path := "/dna/data/api/v1/event/event-series/audit-log/parent-records"

	queryString, _ := query.Values(GetAuditLogParentRecordsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetAuditLogParentRecords{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuditLogParentRecords")
	}

	result := response.Result().(*ResponseEventManagementGetAuditLogParentRecords)
	return result, response, err

}

//GetAuditLogSummary Get AuditLog Summary - 4a87-484a-4df9-819e
/* Get Audit Log Summary from the Event-Hub


@param GetAuditLogSummaryQueryParams Filtering parameter
*/
func (s *EventManagementService) GetAuditLogSummary(GetAuditLogSummaryQueryParams *GetAuditLogSummaryQueryParams) (*ResponseEventManagementGetAuditLogSummary, *resty.Response, error) {
	path := "/dna/data/api/v1/event/event-series/audit-log/summary"

	queryString, _ := query.Values(GetAuditLogSummaryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetAuditLogSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuditLogSummary")
	}

	result := response.Result().(*ResponseEventManagementGetAuditLogSummary)
	return result, response, err

}

//GetAuditLogRecords Get AuditLog Records - 89a9-fafb-4d49-bd86
/* Get Audit Log Event instances from the Event-Hub


@param GetAuditLogRecordsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetAuditLogRecords(GetAuditLogRecordsQueryParams *GetAuditLogRecordsQueryParams) (*ResponseEventManagementGetAuditLogRecords, *resty.Response, error) {
	path := "/dna/data/api/v1/event/event-series/audit-logs"

	queryString, _ := query.Values(GetAuditLogRecordsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetAuditLogRecords{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetAuditLogRecords")
	}

	result := response.Result().(*ResponseEventManagementGetAuditLogRecords)
	return result, response, err

}

//GetStatusAPIForEvents Get Status API for Events - f9bd-99c7-4bba-8832
/* Get the Status of events API calls with provided executionId as mandatory path parameter


@param executionID executionId path parameter. Execution ID

*/
func (s *EventManagementService) GetStatusAPIForEvents(executionID string) (*ResponseEventManagementGetStatusAPIForEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/api-status/{executionId}"
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementGetStatusAPIForEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetStatusApiForEvents")
	}

	result := response.Result().(*ResponseEventManagementGetStatusAPIForEvents)
	return result, response, err

}

//GetNotifications Get Notifications - 8499-9b56-4afb-8657
/* Get the list of Published Notifications


@param GetNotificationsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetNotifications(GetNotificationsQueryParams *GetNotificationsQueryParams) (*ResponseEventManagementGetNotifications, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/event-series"

	queryString, _ := query.Values(GetNotificationsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetNotifications{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetNotifications")
	}

	result := response.Result().(*ResponseEventManagementGetNotifications)
	return result, response, err

}

//CountOfNotifications Count of Notifications - 0eb8-faf7-42aa-abb7
/* Get the Count of Published Notifications


@param CountOfNotificationsQueryParams Filtering parameter
*/
func (s *EventManagementService) CountOfNotifications(CountOfNotificationsQueryParams *CountOfNotificationsQueryParams) (*ResponseEventManagementCountOfNotifications, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/event-series/count"

	queryString, _ := query.Values(CountOfNotificationsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementCountOfNotifications{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CountOfNotifications")
	}

	result := response.Result().(*ResponseEventManagementCountOfNotifications)
	return result, response, err

}

//GetEventSubscriptions Get Event Subscriptions - dcaa-6bde-4feb-9152
/* Gets the list of Subscriptions's based on provided offset and limit (Deprecated)


@param GetEventSubscriptionsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetEventSubscriptions(GetEventSubscriptionsQueryParams *GetEventSubscriptionsQueryParams) (*ResponseEventManagementGetEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(GetEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEventSubscriptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementGetEventSubscriptions)
	return result, response, err

}

//GetEmailSubscriptionDetails Get Email Subscription Details - 339f-d9f5-4719-a410
/* Gets the list of subscription details for specified connectorType


@param GetEmailSubscriptionDetailsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetEmailSubscriptionDetails(GetEmailSubscriptionDetailsQueryParams *GetEmailSubscriptionDetailsQueryParams) (*ResponseEventManagementGetEmailSubscriptionDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription-details/email"

	queryString, _ := query.Values(GetEmailSubscriptionDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEmailSubscriptionDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEmailSubscriptionDetails")
	}

	result := response.Result().(*ResponseEventManagementGetEmailSubscriptionDetails)
	return result, response, err

}

//GetRestWebhookSubscriptionDetails Get Rest/Webhook Subscription Details - eeb6-8baf-4338-bb23
/* Gets the list of subscription details for specified connectorType


@param GetRestWebhookSubscriptionDetailsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetRestWebhookSubscriptionDetails(GetRestWebhookSubscriptionDetailsQueryParams *GetRestWebhookSubscriptionDetailsQueryParams) (*ResponseEventManagementGetRestWebhookSubscriptionDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription-details/rest"

	queryString, _ := query.Values(GetRestWebhookSubscriptionDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetRestWebhookSubscriptionDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRestWebhookSubscriptionDetails")
	}

	result := response.Result().(*ResponseEventManagementGetRestWebhookSubscriptionDetails)
	return result, response, err

}

//GetSyslogSubscriptionDetails Get Syslog Subscription Details - 1785-5b4e-4e69-a497
/* Gets the list of subscription details for specified connectorType


@param GetSyslogSubscriptionDetailsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetSyslogSubscriptionDetails(GetSyslogSubscriptionDetailsQueryParams *GetSyslogSubscriptionDetailsQueryParams) (*ResponseEventManagementGetSyslogSubscriptionDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription-details/syslog"

	queryString, _ := query.Values(GetSyslogSubscriptionDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetSyslogSubscriptionDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSyslogSubscriptionDetails")
	}

	result := response.Result().(*ResponseEventManagementGetSyslogSubscriptionDetails)
	return result, response, err

}

//CountOfEventSubscriptions Count of Event Subscriptions - 149b-7ba0-4e58-90b2
/* Returns the Count of EventSubscriptions


@param CountOfEventSubscriptionsQueryParams Filtering parameter
*/
func (s *EventManagementService) CountOfEventSubscriptions(CountOfEventSubscriptionsQueryParams *CountOfEventSubscriptionsQueryParams) (*ResponseEventManagementCountOfEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/count"

	queryString, _ := query.Values(CountOfEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementCountOfEventSubscriptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CountOfEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementCountOfEventSubscriptions)
	return result, response, err

}

//GetEmailEventSubscriptions Get Email Event Subscriptions - 39b2-0851-4b39-837e
/* Gets the list of email Subscriptions's based on provided query params


@param GetEmailEventSubscriptionsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetEmailEventSubscriptions(GetEmailEventSubscriptionsQueryParams *GetEmailEventSubscriptionsQueryParams) (*ResponseEventManagementGetEmailEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/email"

	queryString, _ := query.Values(GetEmailEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEmailEventSubscriptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEmailEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementGetEmailEventSubscriptions)
	return result, response, err

}

//GetRestWebhookEventSubscriptions Get Rest/Webhook Event Subscriptions - dcaa-6bde-4feb-9153
/* Gets the list of Rest/Webhook Subscriptions's based on provided query params


@param GetRestWebhookEventSubscriptionsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetRestWebhookEventSubscriptions(GetRestWebhookEventSubscriptionsQueryParams *GetRestWebhookEventSubscriptionsQueryParams) (*ResponseEventManagementGetRestWebhookEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/rest"

	queryString, _ := query.Values(GetRestWebhookEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetRestWebhookEventSubscriptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetRestWebhookEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementGetRestWebhookEventSubscriptions)
	return result, response, err

}

//GetSyslogEventSubscriptions Get Syslog Event Subscriptions - c5a9-2a5b-4e6a-852e
/* Gets the list of Syslog Subscriptions's based on provided offset and limit


@param GetSyslogEventSubscriptionsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetSyslogEventSubscriptions(GetSyslogEventSubscriptionsQueryParams *GetSyslogEventSubscriptionsQueryParams) (*ResponseEventManagementGetSyslogEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/syslog"

	queryString, _ := query.Values(GetSyslogEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetSyslogEventSubscriptions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSyslogEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementGetSyslogEventSubscriptions)
	return result, response, err

}

//GetEvents Get Events - 44a3-9a07-4a6a-82a2
/* Gets the list of registered Events with provided eventIds or tags as mandatory


@param GetEventsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetEvents(GetEventsQueryParams *GetEventsQueryParams) (*ResponseEventManagementGetEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/events"

	queryString, _ := query.Values(GetEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEvents")
	}

	result := response.Result().(*ResponseEventManagementGetEvents)
	return result, response, err

}

//CountOfEvents Count of Events - 6a9e-dac1-49ba-86cf
/* Get the count of registered events with provided eventIds or tags as mandatory


@param CountOfEventsQueryParams Filtering parameter
*/
func (s *EventManagementService) CountOfEvents(CountOfEventsQueryParams *CountOfEventsQueryParams) (*ResponseEventManagementCountOfEvents, *resty.Response, error) {
	path := "/dna/intent/api/v1/events/count"

	queryString, _ := query.Values(CountOfEventsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementCountOfEvents{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CountOfEvents")
	}

	result := response.Result().(*ResponseEventManagementCountOfEvents)
	return result, response, err

}

//GetEventArtifacts Get EventArtifacts - 73b1-d832-4c98-bc22
/* Gets the list of artifacts based on provided offset and limit


@param GetEventArtifactsQueryParams Filtering parameter
*/
func (s *EventManagementService) GetEventArtifacts(GetEventArtifactsQueryParams *GetEventArtifactsQueryParams) (*ResponseEventManagementGetEventArtifacts, *resty.Response, error) {
	path := "/dna/system/api/v1/event/artifact"

	queryString, _ := query.Values(GetEventArtifactsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementGetEventArtifacts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetEventArtifacts")
	}

	result := response.Result().(*ResponseEventManagementGetEventArtifacts)
	return result, response, err

}

//EventArtifactCount EventArtifact Count - b78e-9bf7-4f8a-8321
/* Get the count of registered event artifacts with provided eventIds or tags as mandatory


 */
func (s *EventManagementService) EventArtifactCount() (*ResponseEventManagementEventArtifactCount, *resty.Response, error) {
	path := "/dna/system/api/v1/event/artifact/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementEventArtifactCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation EventArtifactCount")
	}

	result := response.Result().(*ResponseEventManagementEventArtifactCount)
	return result, response, err

}

//GetConnectorTypes Get Connector Types - c0be-2a2b-4dc9-8bfb
/* Get the list of connector types


 */
func (s *EventManagementService) GetConnectorTypes() (*ResponseEventManagementGetConnectorTypes, *resty.Response, error) {
	path := "/dna/system/api/v1/event/config/connector-types"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseEventManagementGetConnectorTypes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetConnectorTypes")
	}

	result := response.Result().(*ResponseEventManagementGetConnectorTypes)
	return result, response, err

}

//CreateEmailDestination Create Email Destination - 84ab-8bff-4769-a7d6
/* Create Email Destination


 */
func (s *EventManagementService) CreateEmailDestination(requestEventManagementCreateEmailDestination *RequestEventManagementCreateEmailDestination) (*ResponseEventManagementCreateEmailDestination, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/email-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateEmailDestination).
		SetResult(&ResponseEventManagementCreateEmailDestination{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateEmailDestination")
	}

	result := response.Result().(*ResponseEventManagementCreateEmailDestination)
	return result, response, err

}

//CreateEventSubscriptions Create Event Subscriptions - 4f9f-7a7b-40f9-90de
/* Subscribe SubscriptionEndpoint to list of registered events (Deprecated)


 */
func (s *EventManagementService) CreateEventSubscriptions(requestEventManagementCreateEventSubscriptions *RequestEventManagementCreateEventSubscriptions) (*ResponseEventManagementCreateEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateEventSubscriptions).
		SetResult(&ResponseEventManagementCreateEventSubscriptions{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementCreateEventSubscriptions)
	return result, response, err

}

//CreateEmailEventSubscription Create Email Event Subscription - 7bbc-88c8-424a-840f
/* Create Email Subscription Endpoint for list of registered events.


 */
func (s *EventManagementService) CreateEmailEventSubscription(requestEventManagementCreateEmailEventSubscription *RequestEventManagementCreateEmailEventSubscription) (*ResponseEventManagementCreateEmailEventSubscription, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/email"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateEmailEventSubscription).
		SetResult(&ResponseEventManagementCreateEmailEventSubscription{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateEmailEventSubscription")
	}

	result := response.Result().(*ResponseEventManagementCreateEmailEventSubscription)
	return result, response, err

}

//CreateRestWebhookEventSubscription Create Rest/Webhook Event Subscription - 9584-d988-45eb-b4b0
/* Create Rest/Webhook Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) CreateRestWebhookEventSubscription(requestEventManagementCreateRestWebhookEventSubscription *RequestEventManagementCreateRestWebhookEventSubscription) (*ResponseEventManagementCreateRestWebhookEventSubscription, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/rest"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateRestWebhookEventSubscription).
		SetResult(&ResponseEventManagementCreateRestWebhookEventSubscription{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateRestWebhookEventSubscription")
	}

	result := response.Result().(*ResponseEventManagementCreateRestWebhookEventSubscription)
	return result, response, err

}

//CreateSyslogEventSubscription Create Syslog Event Subscription - 919a-8bb7-445a-88fe
/* Create Syslog Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) CreateSyslogEventSubscription(requestEventManagementCreateSyslogEventSubscription *RequestEventManagementCreateSyslogEventSubscription) (*ResponseEventManagementCreateSyslogEventSubscription, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/syslog"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateSyslogEventSubscription).
		SetResult(&ResponseEventManagementCreateSyslogEventSubscription{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSyslogEventSubscription")
	}

	result := response.Result().(*ResponseEventManagementCreateSyslogEventSubscription)
	return result, response, err

}

//CreateSyslogDestination Create Syslog Destination - 08ad-2bd5-47fa-9227
/* Create Syslog Destination


 */
func (s *EventManagementService) CreateSyslogDestination(requestEventManagementCreateSyslogDestination *RequestEventManagementCreateSyslogDestination) (*ResponseEventManagementCreateSyslogDestination, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/syslogConfig"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateSyslogDestination).
		SetResult(&ResponseEventManagementCreateSyslogDestination{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateSyslogDestination")
	}

	result := response.Result().(*ResponseEventManagementCreateSyslogDestination)
	return result, response, err

}

//CreateWebhookDestination Create Webhook Destination - 4bbd-580a-4bab-80b7
/* Create Webhook Destination


 */
func (s *EventManagementService) CreateWebhookDestination(requestEventManagementCreateWebhookDestination *RequestEventManagementCreateWebhookDestination) (*ResponseEventManagementCreateWebhookDestination, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/webhook"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementCreateWebhookDestination).
		SetResult(&ResponseEventManagementCreateWebhookDestination{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CreateWebhookDestination")
	}

	result := response.Result().(*ResponseEventManagementCreateWebhookDestination)
	return result, response, err

}

//UpdateEmailDestination Update Email Destination - 8285-3bf9-4aba-a65f
/* Update Email Destination


 */
func (s *EventManagementService) UpdateEmailDestination(requestEventManagementUpdateEmailDestination *RequestEventManagementUpdateEmailDestination) (*ResponseEventManagementUpdateEmailDestination, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/email-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateEmailDestination).
		SetResult(&ResponseEventManagementUpdateEmailDestination{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEmailDestination")
	}

	result := response.Result().(*ResponseEventManagementUpdateEmailDestination)
	return result, response, err

}

//UpdateEventSubscriptions Update Event Subscriptions - 579a-6a72-48cb-94cf
/* Update SubscriptionEndpoint to list of registered events(Deprecated)


 */
func (s *EventManagementService) UpdateEventSubscriptions(requestEventManagementUpdateEventSubscriptions *RequestEventManagementUpdateEventSubscriptions) (*ResponseEventManagementUpdateEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateEventSubscriptions).
		SetResult(&ResponseEventManagementUpdateEventSubscriptions{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementUpdateEventSubscriptions)
	return result, response, err

}

//UpdateEmailEventSubscription Update Email Event Subscription - 87b2-2b83-46bb-8983
/* Update Email Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) UpdateEmailEventSubscription(requestEventManagementUpdateEmailEventSubscription *RequestEventManagementUpdateEmailEventSubscription) (*ResponseEventManagementUpdateEmailEventSubscription, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/email"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateEmailEventSubscription).
		SetResult(&ResponseEventManagementUpdateEmailEventSubscription{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateEmailEventSubscription")
	}

	result := response.Result().(*ResponseEventManagementUpdateEmailEventSubscription)
	return result, response, err

}

//UpdateRestWebhookEventSubscription Update Rest/Webhook Event Subscription - ce81-f9c5-4fc8-b576
/* Update Rest/Webhook Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) UpdateRestWebhookEventSubscription(requestEventManagementUpdateRestWebhookEventSubscription *RequestEventManagementUpdateRestWebhookEventSubscription) (*ResponseEventManagementUpdateRestWebhookEventSubscription, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/rest"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateRestWebhookEventSubscription).
		SetResult(&ResponseEventManagementUpdateRestWebhookEventSubscription{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateRestWebhookEventSubscription")
	}

	result := response.Result().(*ResponseEventManagementUpdateRestWebhookEventSubscription)
	return result, response, err

}

//UpdateSyslogEventSubscription Update Syslog Event Subscription - 6285-cbc1-4039-9ace
/* Update Syslog Subscription Endpoint for list of registered events


 */
func (s *EventManagementService) UpdateSyslogEventSubscription(requestEventManagementUpdateSyslogEventSubscription *RequestEventManagementUpdateSyslogEventSubscription) (*ResponseEventManagementUpdateSyslogEventSubscription, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription/syslog"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateSyslogEventSubscription).
		SetResult(&ResponseEventManagementUpdateSyslogEventSubscription{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSyslogEventSubscription")
	}

	result := response.Result().(*ResponseEventManagementUpdateSyslogEventSubscription)
	return result, response, err

}

//UpdateSyslogDestination Update Syslog Destination - eeac-fbe9-4518-b3fb
/* Update Syslog Destination


 */
func (s *EventManagementService) UpdateSyslogDestination(requestEventManagementUpdateSyslogDestination *RequestEventManagementUpdateSyslogDestination) (*ResponseEventManagementUpdateSyslogDestination, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/syslogConfig"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateSyslogDestination).
		SetResult(&ResponseEventManagementUpdateSyslogDestination{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateSyslogDestination")
	}

	result := response.Result().(*ResponseEventManagementUpdateSyslogDestination)
	return result, response, err

}

//UpdateWebhookDestination Update Webhook Destination - 06b5-da28-423a-9bf6
/* Update Webhook Destination


 */
func (s *EventManagementService) UpdateWebhookDestination(requestEventManagementUpdateWebhookDestination *RequestEventManagementUpdateWebhookDestination) (*ResponseEventManagementUpdateWebhookDestination, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/webhook"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestEventManagementUpdateWebhookDestination).
		SetResult(&ResponseEventManagementUpdateWebhookDestination{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdateWebhookDestination")
	}

	result := response.Result().(*ResponseEventManagementUpdateWebhookDestination)
	return result, response, err

}

//DeleteEventSubscriptions Delete Event Subscriptions - 9398-1baa-4079-9483
/* Delete EventSubscriptions


@param DeleteEventSubscriptionsQueryParams Filtering parameter
*/
func (s *EventManagementService) DeleteEventSubscriptions(DeleteEventSubscriptionsQueryParams *DeleteEventSubscriptionsQueryParams) (*ResponseEventManagementDeleteEventSubscriptions, *resty.Response, error) {
	path := "/dna/intent/api/v1/event/subscription"

	queryString, _ := query.Values(DeleteEventSubscriptionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseEventManagementDeleteEventSubscriptions{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeleteEventSubscriptions")
	}

	result := response.Result().(*ResponseEventManagementDeleteEventSubscriptions)
	return result, response, err

}
