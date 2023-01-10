package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TaskService service

type GetTasksQueryParams struct {
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
	Offset        int    `url:"offset,omitempty"`        //offset
	Limit         int    `url:"limit,omitempty"`         //limit
	SortBy        string `url:"sortBy,omitempty"`        //Sort results by this field
	Order         string `url:"order,omitempty"`         //Sort order - asc or dsc
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

type ResponseTaskGetBusinessAPIExecutionDetails struct {
	BapiKey           string `json:"bapiKey,omitempty"`           // Business API Key (UUID)
	BapiName          string `json:"bapiName,omitempty"`          // Name of the Business API
	BapiExecutionID   string `json:"bapiExecutionId,omitempty"`   // Execution Id of the Business API (UUID)
	StartTime         string `json:"startTime,omitempty"`         // Execution Start Time of the Business API (Date Time Format)
	StartTimeEpoch    *int   `json:"startTimeEpoch,omitempty"`    // Execution Start Time of the Business API (Epoch Milliseconds)
	EndTime           string `json:"endTime,omitempty"`           // Execution End Time of the Business API (Date Time Format)
	EndTimeEpoch      *int   `json:"endTimeEpoch,omitempty"`      // Execution End Time of the Business API (Epoch Milliseconds)
	TimeDuration      *int   `json:"timeDuration,omitempty"`      // Time taken for Business API Execution (Milliseconds)
	Status            string `json:"status,omitempty"`            // Execution status of the Business API
	BapiError         string `json:"bapiError,omitempty"`         // Bapi Error
	RuntimeInstanceID string `json:"runtimeInstanceId,omitempty"` // Pod Id in which the Business API is executed
}
type ResponseTaskGetTasks struct {
	Response *[]ResponseTaskGetTasksResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}
type ResponseTaskGetTasksResponse struct {
	AdditionalStatusURL string                                       `json:"additionalStatusURL,omitempty"` //
	Data                string                                       `json:"data,omitempty"`                //
	EndTime             *int                                         `json:"endTime,omitempty"`             //
	ErrorCode           string                                       `json:"errorCode,omitempty"`           //
	ErrorKey            string                                       `json:"errorKey,omitempty"`            //
	FailureReason       string                                       `json:"failureReason,omitempty"`       //
	ID                  string                                       `json:"id,omitempty"`                  //
	InstanceTenantID    string                                       `json:"instanceTenantId,omitempty"`    //
	IsError             *bool                                        `json:"isError,omitempty"`             //
	LastUpdate          string                                       `json:"lastUpdate,omitempty"`          //
	OperationIDList     *ResponseTaskGetTasksResponseOperationIDList `json:"operationIdList,omitempty"`     //
	ParentID            string                                       `json:"parentId,omitempty"`            //
	Progress            string                                       `json:"progress,omitempty"`            //
	RootID              string                                       `json:"rootId,omitempty"`              //
	ServiceType         string                                       `json:"serviceType,omitempty"`         //
	StartTime           *int                                         `json:"startTime,omitempty"`           //
	Username            string                                       `json:"username,omitempty"`            //
	Version             *int                                         `json:"version,omitempty"`             //
}
type ResponseTaskGetTasksResponseOperationIDList interface{}
type ResponseTaskGetTaskCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByOperationID struct {
	Response *[]ResponseTaskGetTaskByOperationIDResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByOperationIDResponse struct {
	AdditionalStatusURL string                                                   `json:"additionalStatusURL,omitempty"` //
	Data                string                                                   `json:"data,omitempty"`                //
	EndTime             string                                                   `json:"endTime,omitempty"`             //
	ErrorCode           string                                                   `json:"errorCode,omitempty"`           //
	ErrorKey            string                                                   `json:"errorKey,omitempty"`            //
	FailureReason       string                                                   `json:"failureReason,omitempty"`       //
	ID                  string                                                   `json:"id,omitempty"`                  //
	InstanceTenantID    string                                                   `json:"instanceTenantId,omitempty"`    //
	IsError             *bool                                                    `json:"isError,omitempty"`             //
	LastUpdate          string                                                   `json:"lastUpdate,omitempty"`          //
	OperationIDList     *ResponseTaskGetTaskByOperationIDResponseOperationIDList `json:"operationIdList,omitempty"`     //
	ParentID            string                                                   `json:"parentId,omitempty"`            //
	Progress            string                                                   `json:"progress,omitempty"`            //
	RootID              string                                                   `json:"rootId,omitempty"`              //
	ServiceType         string                                                   `json:"serviceType,omitempty"`         //
	StartTime           string                                                   `json:"startTime,omitempty"`           //
	Username            string                                                   `json:"username,omitempty"`            //
	Version             *int                                                     `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskByOperationIDResponseOperationIDList interface{}
type ResponseTaskGetTaskByID struct {
	Response *ResponseTaskGetTaskByIDResponse `json:"response,omitempty"` //
	Version  string                           `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskByIDResponse struct {
	AdditionalStatusURL string                                          `json:"additionalStatusURL,omitempty"` //
	Data                string                                          `json:"data,omitempty"`                //
	EndTime             *int                                            `json:"endTime,omitempty"`             //
	ErrorCode           string                                          `json:"errorCode,omitempty"`           //
	ErrorKey            string                                          `json:"errorKey,omitempty"`            //
	FailureReason       string                                          `json:"failureReason,omitempty"`       //
	ID                  string                                          `json:"id,omitempty"`                  //
	InstanceTenantID    string                                          `json:"instanceTenantId,omitempty"`    //
	IsError             *bool                                           `json:"isError,omitempty"`             //
	LastUpdate          *int                                            `json:"lastUpdate,omitempty"`          //
	OperationIDList     *ResponseTaskGetTaskByIDResponseOperationIDList `json:"operationIdList,omitempty"`     //
	ParentID            string                                          `json:"parentId,omitempty"`            //
	Progress            string                                          `json:"progress,omitempty"`            //
	RootID              string                                          `json:"rootId,omitempty"`              //
	ServiceType         string                                          `json:"serviceType,omitempty"`         //
	StartTime           *int                                            `json:"startTime,omitempty"`           //
	Username            string                                          `json:"username,omitempty"`            //
	Version             *int                                            `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskByIDResponseOperationIDList interface{}
type ResponseTaskGetTaskTree struct {
	Response *[]ResponseTaskGetTaskTreeResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}
type ResponseTaskGetTaskTreeResponse struct {
	AdditionalStatusURL string                                          `json:"additionalStatusURL,omitempty"` //
	Data                string                                          `json:"data,omitempty"`                //
	EndTime             string                                          `json:"endTime,omitempty"`             //
	ErrorCode           string                                          `json:"errorCode,omitempty"`           //
	ErrorKey            string                                          `json:"errorKey,omitempty"`            //
	FailureReason       string                                          `json:"failureReason,omitempty"`       //
	ID                  string                                          `json:"id,omitempty"`                  //
	InstanceTenantID    string                                          `json:"instanceTenantId,omitempty"`    //
	IsError             *bool                                           `json:"isError,omitempty"`             //
	LastUpdate          string                                          `json:"lastUpdate,omitempty"`          //
	OperationIDList     *ResponseTaskGetTaskTreeResponseOperationIDList `json:"operationIdList,omitempty"`     //
	ParentID            string                                          `json:"parentId,omitempty"`            //
	Progress            string                                          `json:"progress,omitempty"`            //
	RootID              string                                          `json:"rootId,omitempty"`              //
	ServiceType         string                                          `json:"serviceType,omitempty"`         //
	StartTime           string                                          `json:"startTime,omitempty"`           //
	Username            string                                          `json:"username,omitempty"`            //
	Version             *int                                            `json:"version,omitempty"`             //
}
type ResponseTaskGetTaskTreeResponseOperationIDList interface{}

//GetBusinessAPIExecutionDetails Get Business API Execution Details - c1bc-a8c1-41fb-9f75
/* Retrieves the execution details of a Business API


@param executionID executionId path parameter. Execution Id of API

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
		return nil, response, fmt.Errorf("error with operation GetBusinessApiExecutionDetails")
	}

	result := response.Result().(*ResponseTaskGetBusinessAPIExecutionDetails)
	return result, response, err

}

//GetTasks Get tasks - e78b-b8a2-449b-9eed
/* Returns task(s) based on filter criteria


@param GetTasksQueryParams Filtering parameter
*/
func (s *TaskService) GetTasks(GetTasksQueryParams *GetTasksQueryParams) (*ResponseTaskGetTasks, *resty.Response, error) {
	path := "/dna/intent/api/v1/task"

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
		return nil, response, fmt.Errorf("error with operation GetTasks")
	}

	result := response.Result().(*ResponseTaskGetTasks)
	return result, response, err

}

//GetTaskCount Get task count - 26b4-4ab0-4649-a183
/* Returns Task count


@param GetTaskCountQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetTaskByOperationId")
	}

	result := response.Result().(*ResponseTaskGetTaskByOperationID)
	return result, response, err

}

//GetTaskByID Get task by Id - a1a9-3873-46ba-92b1
/* Returns a task by specified id


@param taskID taskId path parameter. UUID of the Task

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
		return nil, response, fmt.Errorf("error with operation GetTaskById")
	}

	result := response.Result().(*ResponseTaskGetTaskByID)
	return result, response, err

}

//GetTaskTree Get task tree - f5a2-69c4-4f2a-95fa
/* Returns a task with its children tasks by based on their id


@param taskID taskId path parameter. UUID of the Task

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
		return nil, response, fmt.Errorf("error with operation GetTaskTree")
	}

	result := response.Result().(*ResponseTaskGetTaskTree)
	return result, response, err

}
