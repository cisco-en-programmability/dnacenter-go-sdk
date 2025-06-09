package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TaskService service

type RetrieveAListOfAssuranceTasksQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Maximum number of records to return
	Offset float64 `url:"offset,omitempty"` //Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
	SortBy string  `url:"sortBy,omitempty"` //A field within the response to sort by.
	Order  string  `url:"order,omitempty"`  //The sort order of the field ascending or descending.
	Status string  `url:"status,omitempty"` //used to get a subset of tasks by their status
}
type RetrieveAListOfAssuranceTasksHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams struct {
	Status string `url:"status,omitempty"` //used to get a subset of tasks by their status
}
type RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type RetrieveASpecificAssuranceTaskByIDHeaderParams struct {
	XCaLLERID string `url:"X-CALLER-ID,omitempty"` //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetActivitiesQueryParams struct {
	Description string  `url:"description,omitempty"` //The description of the activity
	Status      string  `url:"status,omitempty"`      //The status of the activity
	Type        string  `url:"type,omitempty"`        //The type of the activity
	Recurring   bool    `url:"recurring,omitempty"`   //If the activity is recurring
	StartTime   string  `url:"startTime,omitempty"`   //This is the epoch millisecond start time from which activities need to be fetched
	EndTime     string  `url:"endTime,omitempty"`     //This is the epoch millisecond end time upto which activities need to be fetched
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page.
	SortBy      string  `url:"sortBy,omitempty"`      //A property within the response to sort by.
	Order       string  `url:"order,omitempty"`       //Whether ascending or descending order should be used to sort the response.
}
type RetrievesTheCountOfActivitiesQueryParams struct {
	Description string `url:"description,omitempty"` //The description provided when creating the activity
	Status      string `url:"status,omitempty"`      //Status of the activity
	Type        string `url:"type,omitempty"`        //Type of the activity
	Recurring   bool   `url:"recurring,omitempty"`   //Denotes whether an activity is recurring or not
	StartTime   string `url:"startTime,omitempty"`   //This is the epoch millisecond start time from which activities need to be fetched
	EndTime     string `url:"endTime,omitempty"`     //This is the epoch millisecond end time upto which activities need to be fetched
}
type GetTriggeredJobsByActivityIDQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	SortBy string  `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string  `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
}
type GetTasksOperationalTasksQueryParams struct {
	StartTime     string  `url:"startTime,omitempty"`     //This is the epoch start time from which tasks need to be fetched
	EndTime       string  `url:"endTime,omitempty"`       //This is the epoch end time upto which audit records need to be fetched
	Data          string  `url:"data,omitempty"`          //Fetch tasks that contains this data
	ErrorCode     string  `url:"errorCode,omitempty"`     //Fetch tasks that have this error code
	ServiceType   string  `url:"serviceType,omitempty"`   //Fetch tasks with this service type
	Username      string  `url:"username,omitempty"`      //Fetch tasks with this username
	Progress      string  `url:"progress,omitempty"`      //Fetch tasks that contains this progress
	IsError       string  `url:"isError,omitempty"`       //Fetch tasks ended as success or failure. Valid values: true, false
	FailureReason string  `url:"failureReason,omitempty"` //Fetch tasks that contains this failure reason
	ParentID      string  `url:"parentId,omitempty"`      //Fetch tasks that have this parent Id
	Offset        float64 `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy        string  `url:"sortBy,omitempty"`        //Sort results by this field
	Order         string  `url:"order,omitempty"`         //Sort order - asc or dsc
}
type GetTaskCountQueryParams struct {
	StartTime     string `url:"startTime,omitempty"`     //This is the epoch start time from which tasks need to be fetched
	EndTime       string `url:"endTime,omitempty"`       //This is the epoch end time upto which audit records need to be fetched
	Data          string `url:"data,omitempty"`          //Fetch tasks that contains this data
	ErrorCode     string `url:"errorCode,omitempty"`     //Fetch tasks that have this error code
	ServiceType   string `url:"serviceType,omitempty"`   //Fetch tasks with this service type
	Username      string `url:"username,omitempty"`      //Fetch tasks with this username
	Progress      string `url:"progress,omitempty"`      //Fetch tasks that contains this progress
	IsError       string `url:"isError,omitempty"`       //Fetch tasks ended as success or failure. Valid values: true, false
	FailureReason string `url:"failureReason,omitempty"` //Fetch tasks that contains this failure reason
	ParentID      string `url:"parentId,omitempty"`      //Fetch tasks that have this parent Id
}
type GetTasksQueryParams struct {
	Offset    float64 `url:"offset,omitempty"`    //The first record to show for this page; the first record is numbered 1.
	Limit     float64 `url:"limit,omitempty"`     //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortBy    string  `url:"sortBy,omitempty"`    //A property within the response to sort by.
	Order     string  `url:"order,omitempty"`     //Whether ascending or descending order should be used to sort the response.
	StartTime int     `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched
	EndTime   int     `url:"endTime,omitempty"`   //This is the epoch millisecond end time upto which task records need to be fetched
	ParentID  string  `url:"parentId,omitempty"`  //Fetch tasks that have this parent Id
	RootID    string  `url:"rootId,omitempty"`    //Fetch tasks that have this root Id
	Status    string  `url:"status,omitempty"`    //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
}
type GetTasksCountQueryParams struct {
	StartTime int    `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched
	EndTime   int    `url:"endTime,omitempty"`   //This is the epoch millisecond end time upto which task records need to be fetched
	ParentID  string `url:"parentId,omitempty"`  //Fetch tasks that have this parent Id
	RootID    string `url:"rootId,omitempty"`    //Fetch tasks that have this root Id
	Status    string `url:"status,omitempty"`    //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
}

type ResponseTaskRetrieveAListOfAssuranceTasks struct {
	Response *[]ResponseTaskRetrieveAListOfAssuranceTasksResponse `json:"response,omitempty"` //
	Page     *ResponseTaskRetrieveAListOfAssuranceTasksPage       `json:"page,omitempty"`     //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseTaskRetrieveAListOfAssuranceTasksResponse struct {
	ID            string                                                 `json:"id,omitempty"`            // Id
	Status        string                                                 `json:"status,omitempty"`        // Status
	StartTime     *int                                                   `json:"startTime,omitempty"`     // Start Time
	EndTime       *int                                                   `json:"endTime,omitempty"`       // End Time
	UpdateTime    *int                                                   `json:"updateTime,omitempty"`    // Update Time
	Progress      string                                                 `json:"progress,omitempty"`      // Progress
	FailureReason string                                                 `json:"failureReason,omitempty"` // Failure Reason
	ErrorCode     string                                                 `json:"errorCode,omitempty"`     // Error Code
	RequestType   string                                                 `json:"requestType,omitempty"`   // Request Type
	Data          *ResponseTaskRetrieveAListOfAssuranceTasksResponseData `json:"data,omitempty"`          // Data
	ResultURL     string                                                 `json:"resultUrl,omitempty"`     // Result Url
}
type ResponseTaskRetrieveAListOfAssuranceTasksResponseData interface{}
type ResponseTaskRetrieveAListOfAssuranceTasksPage struct {
	Limit  *int                                                   `json:"limit,omitempty"`  // Limit
	Offset *int                                                   `json:"offset,omitempty"` // Offset
	Count  *int                                                   `json:"count,omitempty"`  // Count
	SortBy *[]ResponseTaskRetrieveAListOfAssuranceTasksPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseTaskRetrieveAListOfAssuranceTasksPageSortBy struct {
	Name  string `json:"name,omitempty"`  // Name
	Order string `json:"order,omitempty"` // Order
}
type ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist struct {
	Response *ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistResponse `json:"response,omitempty"` //
	Version  string                                                                           `json:"version,omitempty"`  // Version
}
type ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseTaskRetrieveASpecificAssuranceTaskByID struct {
	Response *ResponseTaskRetrieveASpecificAssuranceTaskByIDResponse `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  // Version
}
type ResponseTaskRetrieveASpecificAssuranceTaskByIDResponse struct {
	ID            string                                                      `json:"id,omitempty"`            // Id
	Status        string                                                      `json:"status,omitempty"`        // Status
	StartTime     *int                                                        `json:"startTime,omitempty"`     // Start Time
	EndTime       *int                                                        `json:"endTime,omitempty"`       // End Time
	UpdateTime    *int                                                        `json:"updateTime,omitempty"`    // Update Time
	Progress      string                                                      `json:"progress,omitempty"`      // Progress
	FailureReason string                                                      `json:"failureReason,omitempty"` // Failure Reason
	ErrorCode     string                                                      `json:"errorCode,omitempty"`     // Error Code
	RequestType   string                                                      `json:"requestType,omitempty"`   // Request Type
	Data          *ResponseTaskRetrieveASpecificAssuranceTaskByIDResponseData `json:"data,omitempty"`          // Data
	ResultURL     string                                                      `json:"resultUrl,omitempty"`     // Result Url
}
type ResponseTaskRetrieveASpecificAssuranceTaskByIDResponseData interface{}
type ResponseTaskGetActivities struct {
	Response *[]ResponseTaskGetActivitiesResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}
type ResponseTaskGetActivitiesResponse struct {
	Description                   string   `json:"description,omitempty"`                   // The description provided when creating the activity
	EndTime                       *float64 `json:"endTime,omitempty"`                       // The time at which the activity ended
	ID                            string   `json:"id,omitempty"`                            // The id of the activity
	OriginatingWorkItemActivityID string   `json:"originatingWorkItemActivityId,omitempty"` // The id of the activity of type WORK_ITEM whose deployment created this TASK activity. a preview activity is of type WORK_ITEM. if the preview is deployed then a new activity of type TASK is created for deploying the preview and the originatingWorkItemActivityId is set to the preview activity id.
	Recurring                     *bool    `json:"recurring,omitempty"`                     // Denotes whether an activity is recurring or not
	StartTime                     *float64 `json:"startTime,omitempty"`                     // The time at which the activity started
	Status                        string   `json:"status,omitempty"`                        // The status of the activity
	Type                          string   `json:"type,omitempty"`                          // Type of activity:    * 'WORK_ITEM' - activity that requires user action when the status is READY or FAILED   * 'TASK' - activity that is not of type WORK_ITEM
}
type ResponseTaskRetrievesTheCountOfActivities struct {
	Response *ResponseTaskRetrievesTheCountOfActivitiesResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // The version of the response
}
type ResponseTaskRetrievesTheCountOfActivitiesResponse struct {
	Count *float64 `json:"count,omitempty"` // The reported count.
}
type ResponseTaskGetTriggeredJobsByActivityID struct {
	Response *[]ResponseTaskGetTriggeredJobsByActivityIDResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  //
}
type ResponseTaskGetTriggeredJobsByActivityIDResponse struct {
	Description string `json:"description,omitempty"` // The description of the triggered job
	ID          string `json:"id,omitempty"`          // The id of the triggered job
	Status      string `json:"status,omitempty"`      // Triggered job state:   * 'NOT_STARTED' - job hasn't started  * 'IN_PROGRESS' - job is running  * 'SUCCESS' - job has completed with success  * 'FAILURE' - job has completed with failure
	TaskID      string `json:"taskId,omitempty"`      // The id of the task for the job that was triggered
	TaskURL     string `json:"taskUrl,omitempty"`     // The path to the API endpoint to GET information on the task
}
type ResponseTaskRetrievesTheCountOfTriggeredJobsByActivityID struct {
	Response *ResponseTaskRetrievesTheCountOfTriggeredJobsByActivityIDResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // The version of the response
}
type ResponseTaskRetrievesTheCountOfTriggeredJobsByActivityIDResponse struct {
	Count *float64 `json:"count,omitempty"` // The reported count.
}
type ResponseTaskGetBusinessAPIExecutionDetails struct {
	BapiKey              string                                                          `json:"bapiKey,omitempty"`              // Business API Key (UUID)
	BapiName             string                                                          `json:"bapiName,omitempty"`             // Name of the Business API
	BapiExecutionID      string                                                          `json:"bapiExecutionId,omitempty"`      // Execution Id of the Business API (UUID)
	StartTime            string                                                          `json:"startTime,omitempty"`            // Execution Start Time of the Business API (Date Time Format)
	StartTimeEpoch       *int                                                            `json:"startTimeEpoch,omitempty"`       // Execution Start Time of the Business API (Epoch Milliseconds)
	EndTime              string                                                          `json:"endTime,omitempty"`              // Execution End Time of the Business API (Date Time Format)
	EndTimeEpoch         *int                                                            `json:"endTimeEpoch,omitempty"`         // Execution End Time of the Business API (Epoch Milliseconds)
	TimeDuration         *int                                                            `json:"timeDuration,omitempty"`         // Time taken for Business API Execution (Milliseconds)
	Status               string                                                          `json:"status,omitempty"`               // Execution status of the Business API
	BapiSyncResponse     string                                                          `json:"bapiSyncResponse,omitempty"`     // Returns the actual response of the original API as a string
	BapiSyncResponseJSON *ResponseTaskGetBusinessAPIExecutionDetailsBapiSyncResponseJSON `json:"bapiSyncResponseJson,omitempty"` // Returns the actual response of the original API  as a json
	RuntimeInstanceID    string                                                          `json:"runtimeInstanceId,omitempty"`    // Pod Id in which the Business API is executed
	BapiError            string                                                          `json:"bapiError,omitempty"`            // Returns the error response of the original API  as a string
}
type ResponseTaskGetBusinessAPIExecutionDetailsBapiSyncResponseJSON interface{}
type ResponseTaskGetTasksOperationalTasks struct {
	Response *[]ResponseTaskGetTasksOperationalTasksResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksOperationalTasksResponse struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     // Operation Id List
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByOperationID struct {
	Response *[]ResponseTaskGetTaskByOperationIDResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByOperationIDResponse struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     // Operation Id List
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskByID struct {
	Response *ResponseTaskGetTaskByIDResponse `json:"response,omitempty"` //
	Version  string                           `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByIDResponse struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     // Operation Id List
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskTree struct {
	Response *[]ResponseTaskGetTaskTreeResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskTreeResponse struct {
	AdditionalStatusURL string   `json:"additionalStatusURL,omitempty"` //
	Data                string   `json:"data,omitempty"`                //
	EndTime             *int     `json:"endTime,omitempty"`             //
	ErrorCode           string   `json:"errorCode,omitempty"`           //
	ErrorKey            string   `json:"errorKey,omitempty"`            //
	FailureReason       string   `json:"failureReason,omitempty"`       //
	ID                  string   `json:"id,omitempty"`                  //
	InstanceTenantID    string   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool    `json:"isError,omitempty"`             //
	LastUpdate          *int     `json:"lastUpdate,omitempty"`          //
	OperationIDList     []string `json:"operationIdList,omitempty"`     // Operation Id List
	ParentID            string   `json:"parentId,omitempty"`            //
	Progress            string   `json:"progress,omitempty"`            //
	RootID              string   `json:"rootId,omitempty"`              //
	ServiceType         string   `json:"serviceType,omitempty"`         //
	StartTime           *int     `json:"startTime,omitempty"`           //
	Username            string   `json:"username,omitempty"`            //
	Version             *int     `json:"version,omitempty"`             //
}
type ResponseTaskGetTasks struct {
	Response *[]ResponseTaskGetTasksResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksResponse struct {
	EndTime        *int   `json:"endTime,omitempty"`        // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds
	ID             string `json:"id,omitempty"`             // The ID of this task
	UpdatedTime    *int   `json:"updatedTime,omitempty"`    // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
	ParentID       string `json:"parentId,omitempty"`       // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.
	ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found
	RootID         string `json:"rootId,omitempty"`         // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.
	StartTime      *int   `json:"startTime,omitempty"`      // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds
	Status         string `json:"status,omitempty"`         //
}
type ResponseTaskGetTasksCount struct {
	Response *ResponseTaskGetTasksCountResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  // The version of the response
}
type ResponseTaskGetTasksCountResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTaskGetTasksByID struct {
	Response *ResponseTaskGetTasksByIDResponse `json:"response,omitempty"` //
	Version  string                            `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksByIDResponse struct {
	EndTime        *int   `json:"endTime,omitempty"`        // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds
	ID             string `json:"id,omitempty"`             // The ID of this task
	UpdatedTime    *int   `json:"updatedTime,omitempty"`    // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds
	ParentID       string `json:"parentId,omitempty"`       // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId.
	ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found
	RootID         string `json:"rootId,omitempty"`         // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task.
	StartTime      *int   `json:"startTime,omitempty"`      // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds
	Status         string `json:"status,omitempty"`         //
}
type ResponseTaskGetTaskDetailsByID struct {
	Response *ResponseTaskGetTaskDetailsByIDResponse `json:"response,omitempty"` //
	Version  string                                  `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskDetailsByIDResponse struct {
	Data          string `json:"data,omitempty"`          // Any data associated with this task; the value may vary significantly across different tasks
	Progress      string `json:"progress,omitempty"`      // A textual representation which indicates the progress of this task; the value may vary significantly across different tasks
	ErrorCode     string `json:"errorCode,omitempty"`     // An error code if in case this task has failed in its execution
	FailureReason string `json:"failureReason,omitempty"` // A textual description indicating the reason why a task has failed
}
type ResponseTaskGetActivityByID struct {
	Response *ResponseTaskGetActivityByIDResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}
type ResponseTaskGetActivityByIDResponse struct {
	Description                   string   `json:"description,omitempty"`                   // The description provided when creating the activity
	EndTime                       *float64 `json:"endTime,omitempty"`                       // The time at which the activity ended
	ID                            string   `json:"id,omitempty"`                            // The id of the activity
	OriginatingWorkItemActivityID string   `json:"originatingWorkItemActivityId,omitempty"` // The id of the activity of type WORK_ITEM whose deployment created this TASK activity. a preview activity is of type WORK_ITEM. if the preview is deployed then a new activity of type TASK is created for deploying the preview and the originatingWorkItemActivityId is set to the preview activity id.
	Recurring                     *bool    `json:"recurring,omitempty"`                     // Denotes whether an activity is recurring or not
	StartTime                     *float64 `json:"startTime,omitempty"`                     // The time at which the activity started
	Status                        string   `json:"status,omitempty"`                        // The status of the activity
	Type                          string   `json:"type,omitempty"`                          // Type of activity:    * 'WORK_ITEM' - activity that requires user action when the status is READY or FAILED   * 'TASK' - activity that is not of type WORK_ITEM
}

//RetrieveAListOfAssuranceTasks Retrieve a list of assurance tasks - ee8a-e874-40ca-8154
/* returns all existing tasks in a paginated list
default sorting of list is `startTime`, `asc`
valid field to sort by are [`startTime`,`endTime`,`updateTime`,`status`] For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml


@param RetrieveAListOfAssuranceTasksHeaderParams Custom header parameters
@param RetrieveAListOfAssuranceTasksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-list-of-assurance-tasks
*/
func (s *TaskService) RetrieveAListOfAssuranceTasks(RetrieveAListOfAssuranceTasksHeaderParams *RetrieveAListOfAssuranceTasksHeaderParams, RetrieveAListOfAssuranceTasksQueryParams *RetrieveAListOfAssuranceTasksQueryParams) (*ResponseTaskRetrieveAListOfAssuranceTasks, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceTasks"

	queryString, _ := query.Values(RetrieveAListOfAssuranceTasksQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrieveAListOfAssuranceTasksHeaderParams != nil {

		if RetrieveAListOfAssuranceTasksHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrieveAListOfAssuranceTasksHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskRetrieveAListOfAssuranceTasks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveAListOfAssuranceTasks(RetrieveAListOfAssuranceTasksHeaderParams, RetrieveAListOfAssuranceTasksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveAListOfAssuranceTasks")
	}

	result := response.Result().(*ResponseTaskRetrieveAListOfAssuranceTasks)
	return result, response, err

}

//RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist Retrieve a count of the number of assurance tasks that currently exist - b094-0b13-423b-bfb4
/* returns a count of the number of assurance tasks that are not expired For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml


@param RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams Custom header parameters
@param RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-count-of-the-number-of-assurance-tasks-that-currently-exist
*/
func (s *TaskService) RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams *RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams, RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams *RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams) (*ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceTasks/count"

	queryString, _ := query.Values(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams != nil {

		if RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist(RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistHeaderParams, RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExistQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist")
	}

	result := response.Result().(*ResponseTaskRetrieveACountOfTheNumberOfAssuranceTasksThatCurrentlyExist)
	return result, response, err

}

//RetrieveASpecificAssuranceTaskByID Retrieve a specific assurance task by id - 1e8f-7ae5-4798-88f1
/* returns a task given a specific task id For detailed information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml


@param id id path parameter. unique task id

@param RetrieveASpecificAssuranceTaskByIdHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-a-specific-assurance-task-by-id
*/
func (s *TaskService) RetrieveASpecificAssuranceTaskByID(id string, RetrieveASpecificAssuranceTaskByIdHeaderParams *RetrieveASpecificAssuranceTaskByIDHeaderParams) (*ResponseTaskRetrieveASpecificAssuranceTaskByID, *resty.Response, error) {
	path := "/dna/data/api/v1/assuranceTasks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RetrieveASpecificAssuranceTaskByIdHeaderParams != nil {

		if RetrieveASpecificAssuranceTaskByIdHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", RetrieveASpecificAssuranceTaskByIdHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseTaskRetrieveASpecificAssuranceTaskByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveASpecificAssuranceTaskByID(id, RetrieveASpecificAssuranceTaskByIdHeaderParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveASpecificAssuranceTaskById")
	}

	result := response.Result().(*ResponseTaskRetrieveASpecificAssuranceTaskByID)
	return result, response, err

}

//GetActivities Get activities - a6ad-1afe-4b98-a3f7
/* Returns activity(s) based on filter criteria


@param GetActivitiesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-activities
*/
func (s *TaskService) GetActivities(GetActivitiesQueryParams *GetActivitiesQueryParams) (*ResponseTaskGetActivities, *resty.Response, error) {
	path := "/dna/intent/api/v1/activities"

	queryString, _ := query.Values(GetActivitiesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetActivities{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetActivities(GetActivitiesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetActivities")
	}

	result := response.Result().(*ResponseTaskGetActivities)
	return result, response, err

}

//RetrievesTheCountOfActivities Retrieves the count of activities - a591-6b36-4238-a907
/* Retrieves the count of activities


@param RetrievesTheCountOfActivitiesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-activities
*/
func (s *TaskService) RetrievesTheCountOfActivities(RetrievesTheCountOfActivitiesQueryParams *RetrievesTheCountOfActivitiesQueryParams) (*ResponseTaskRetrievesTheCountOfActivities, *resty.Response, error) {
	path := "/dna/intent/api/v1/activities/count"

	queryString, _ := query.Values(RetrievesTheCountOfActivitiesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskRetrievesTheCountOfActivities{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfActivities(RetrievesTheCountOfActivitiesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfActivities")
	}

	result := response.Result().(*ResponseTaskRetrievesTheCountOfActivities)
	return result, response, err

}

//GetTriggeredJobsByActivityID Get triggered jobs by activity id. - 539b-2999-497b-8848
/* Returns the triggered jobs by the activity with the given activity id


@param activityID activityId path parameter. The id of the activity

@param GetTriggeredJobsByActivityIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-triggered-jobs-by-activity-id
*/
func (s *TaskService) GetTriggeredJobsByActivityID(activityID string, GetTriggeredJobsByActivityIdQueryParams *GetTriggeredJobsByActivityIDQueryParams) (*ResponseTaskGetTriggeredJobsByActivityID, *resty.Response, error) {
	path := "/dna/intent/api/v1/activities/{activityId}/triggeredJobs"
	path = strings.Replace(path, "{activityId}", fmt.Sprintf("%v", activityID), -1)

	queryString, _ := query.Values(GetTriggeredJobsByActivityIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTriggeredJobsByActivityID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTriggeredJobsByActivityID(activityID, GetTriggeredJobsByActivityIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTriggeredJobsByActivityId")
	}

	result := response.Result().(*ResponseTaskGetTriggeredJobsByActivityID)
	return result, response, err

}

//RetrievesTheCountOfTriggeredJobsByActivityID Retrieves the count of triggered jobs by activity id. - e8a1-cbfc-4ae9-a4ed
/* Retrieves the count of triggered jobs by activity id.


@param activityID activityId path parameter. The id of the activity


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-triggered-jobs-by-activity-id
*/
func (s *TaskService) RetrievesTheCountOfTriggeredJobsByActivityID(activityID string) (*ResponseTaskRetrievesTheCountOfTriggeredJobsByActivityID, *resty.Response, error) {
	path := "/dna/intent/api/v1/activities/{activityId}/triggeredJobs/count"
	path = strings.Replace(path, "{activityId}", fmt.Sprintf("%v", activityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskRetrievesTheCountOfTriggeredJobsByActivityID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfTriggeredJobsByActivityID(activityID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfTriggeredJobsByActivityId")
	}

	result := response.Result().(*ResponseTaskRetrievesTheCountOfTriggeredJobsByActivityID)
	return result, response, err

}

//GetBusinessAPIExecutionDetails Get Business API Execution Details - c1bc-a8c1-41fb-9f75
/* Retrieves the execution details of a Business API


@param executionID executionId path parameter. Execution Id of API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-business-api-execution-details
*/
func (s *TaskService) GetBusinessAPIExecutionDetails(executionID string) (*ResponseTaskGetBusinessAPIExecutionDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnacaap/management/execution-status/{executionId}"
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetBusinessAPIExecutionDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBusinessAPIExecutionDetails(executionID)
		}
		return nil, response, fmt.Errorf("error with operation GetBusinessApiExecutionDetails")
	}

	result := response.Result().(*ResponseTaskGetBusinessAPIExecutionDetails)
	return result, response, err

}

//GetTasksOperationalTasks Get tasks - e78b-b8a2-449b-9eed
/* Returns task(s) based on filter criteria


@param GetTasksOperationalTasksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-operational-tasks
*/
func (s *TaskService) GetTasksOperationalTasks(GetTasksOperationalTasksQueryParams *GetTasksOperationalTasksQueryParams) (*ResponseTaskGetTasksOperationalTasks, *resty.Response, error) {
	path := "/dna/intent/api/v1/task"

	queryString, _ := query.Values(GetTasksOperationalTasksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasksOperationalTasks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksOperationalTasks(GetTasksOperationalTasksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksOperationalTasks")
	}

	result := response.Result().(*ResponseTaskGetTasksOperationalTasks)
	return result, response, err

}

//GetTaskCount Get task count - 26b4-4ab0-4649-a183
/* Returns Task count


@param GetTaskCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-count
*/
func (s *TaskService) GetTaskCount(GetTaskCountQueryParams *GetTaskCountQueryParams) (*ResponseTaskGetTaskCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/count"

	queryString, _ := query.Values(GetTaskCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTaskCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskCount(GetTaskCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskCount")
	}

	result := response.Result().(*ResponseTaskGetTaskCount)
	return result, response, err

}

//GetTaskByOperationID Get task by OperationId - e487-f8d3-481b-94f2
/* Returns root tasks associated with an Operationid


@param operationID operationId path parameter.
@param offset offset path parameter. Index, minimum value is 0

@param limit limit path parameter. The maximum value of {limit} supported is 500. <br/> Base 1 indexing for {limit}, minimum value is 1


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-operation-id
*/
func (s *TaskService) GetTaskByOperationID(operationID string, offset int, limit int) (*ResponseTaskGetTaskByOperationID, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/operation/{operationId}/{offset}/{limit}"
	path = strings.Replace(path, "{operationId}", fmt.Sprintf("%v", operationID), -1)
	path = strings.Replace(path, "{offset}", fmt.Sprintf("%v", offset), -1)
	path = strings.Replace(path, "{limit}", fmt.Sprintf("%v", limit), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskByOperationID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskByOperationID(operationID, offset, limit)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskByOperationId")
	}

	result := response.Result().(*ResponseTaskGetTaskByOperationID)
	return result, response, err

}

//GetTaskByID Get task by Id - a1a9-3873-46ba-92b1
/* Returns a task by specified id


@param taskID taskId path parameter. UUID of the Task


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-id
*/
func (s *TaskService) GetTaskByID(taskID string) (*ResponseTaskGetTaskByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskByID(taskID)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskById")
	}

	result := response.Result().(*ResponseTaskGetTaskByID)
	return result, response, err

}

//GetTaskTree Get task tree - f5a2-69c4-4f2a-95fa
/* Returns a task with its children tasks by based on their id


@param taskID taskId path parameter. UUID of the Task


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-tree
*/
func (s *TaskService) GetTaskTree(taskID string) (*ResponseTaskGetTaskTree, *resty.Response, error) {
	path := "/dna/intent/api/v1/task/{taskId}/tree"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskTree{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskTree(taskID)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskTree")
	}

	result := response.Result().(*ResponseTaskGetTaskTree)
	return result, response, err

}

//GetTasks Get tasks - b7bf-c860-466b-aaa7
/* Returns task(s) based on filter criteria


@param GetTasksQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks
*/
func (s *TaskService) GetTasks(GetTasksQueryParams *GetTasksQueryParams) (*ResponseTaskGetTasks, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks"

	queryString, _ := query.Values(GetTasksQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasks{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasks(GetTasksQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTasks")
	}

	result := response.Result().(*ResponseTaskGetTasks)
	return result, response, err

}

//GetTasksCount Get tasks count - 6bb9-395b-4af9-8285
/* Returns the number of tasks that meet the filter criteria


@param GetTasksCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-count
*/
func (s *TaskService) GetTasksCount(GetTasksCountQueryParams *GetTasksCountQueryParams) (*ResponseTaskGetTasksCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks/count"

	queryString, _ := query.Values(GetTasksCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasksCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksCount(GetTasksCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksCount")
	}

	result := response.Result().(*ResponseTaskGetTasksCount)
	return result, response, err

}

//GetTasksByID Get tasks by ID - e493-ea85-4038-8183
/* Returns the task with the given ID


@param id id path parameter. the `id` of the task to retrieve


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-by-id
*/
func (s *TaskService) GetTasksByID(id string) (*ResponseTaskGetTasksByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTasksByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTasksByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTasksById")
	}

	result := response.Result().(*ResponseTaskGetTasksByID)
	return result, response, err

}

//GetTaskDetailsByID Get task details by ID - 408d-8acf-43fb-92c2
/* Returns the task details for the given task ID


@param id id path parameter. the `id` of the task to retrieve details for


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-details-by-id
*/
func (s *TaskService) GetTaskDetailsByID(id string) (*ResponseTaskGetTaskDetailsByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/tasks/{id}/detail"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetTaskDetailsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskDetailsByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskDetailsById")
	}

	result := response.Result().(*ResponseTaskGetTaskDetailsByID)
	return result, response, err

}

//GetActivityByID Get activity by ID. - 7f9e-892b-4749-9d3b
/* Returns the activity with the given ID


@param id id path parameter. The id of the activity to retrieve


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-activity-by-id
*/
func (s *TaskService) GetActivityByID(id string) (*ResponseTaskGetActivityByID, *resty.Response, error) {
	path := "/intent/api/v1/activities/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTaskGetActivityByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetActivityByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetActivityById")
	}

	result := response.Result().(*ResponseTaskGetActivityByID)
	return result, response, err

}
