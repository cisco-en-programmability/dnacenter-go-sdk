package dnac

import (
	"fmt"
	"strings"
	"net/http"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TaskService service


type GetTasks2QueryParams struct{
	StartTime string `url:"startTime,omitempty"` //This is the epoch start time from which tasks need to be fetched 
	EndTime string `url:"endTime,omitempty"` //This is the epoch end time upto which audit records need to be fetched 
	Data string `url:"data,omitempty"` //Fetch tasks that contains this data 
	ErrorCode string `url:"errorCode,omitempty"` //Fetch tasks that have this error code 
	ServiceType string `url:"serviceType,omitempty"` //Fetch tasks with this service type 
	Username string `url:"username,omitempty"` //Fetch tasks with this username 
	Progress string `url:"progress,omitempty"` //Fetch tasks that contains this progress 
	IsError string `url:"isError,omitempty"` //Fetch tasks ended as success or failure. Valid values: true, false 
	FailureReason string `url:"failureReason,omitempty"` //Fetch tasks that contains this failure reason 
	ParentID string `url:"parentId,omitempty"` //Fetch tasks that have this parent Id 
	Offset int `url:"offset,omitempty"` //offset 
	Limit int `url:"limit,omitempty"` //limit 
	SortBy string `url:"sortBy,omitempty"` //Sort results by this field 
	Order string `url:"order,omitempty"` //Sort order - asc or dsc 
}
type GetTaskCountQueryParams struct{
	StartTime string `url:"startTime,omitempty"` //This is the epoch start time from which tasks need to be fetched 
	EndTime string `url:"endTime,omitempty"` //This is the epoch end time upto which audit records need to be fetched 
	Data string `url:"data,omitempty"` //Fetch tasks that contains this data 
	ErrorCode string `url:"errorCode,omitempty"` //Fetch tasks that have this error code 
	ServiceType string `url:"serviceType,omitempty"` //Fetch tasks with this service type 
	Username string `url:"username,omitempty"` //Fetch tasks with this username 
	Progress string `url:"progress,omitempty"` //Fetch tasks that contains this progress 
	IsError string `url:"isError,omitempty"` //Fetch tasks ended as success or failure. Valid values: true, false 
	FailureReason string `url:"failureReason,omitempty"` //Fetch tasks that contains this failure reason 
	ParentID string `url:"parentId,omitempty"` //Fetch tasks that have this parent Id 
}
type GetTasksQueryParams struct{
	Offset int `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. 
	Limit int `url:"limit,omitempty"` //The number of records to show for this page. 
	SortBy string `url:"sortBy,omitempty"` //A property within the response to sort by. 
	Order string `url:"order,omitempty"` //Whether ascending or descending order should be used to sort the response. 
	StartTime int `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched 
	EndTime int `url:"endTime,omitempty"` //This is the epoch millisecond end time upto which task records need to be fetched 
	ParentID string `url:"parentId,omitempty"` //Fetch tasks that have this parent Id 
	RootID string `url:"rootId,omitempty"` //Fetch tasks that have this root Id 
	Status string `url:"status,omitempty"` //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS 
}
type GetTasksCountQueryParams struct{
	StartTime int `url:"startTime,omitempty"` //This is the epoch millisecond start time from which tasks need to be fetched 
	EndTime int `url:"endTime,omitempty"` //This is the epoch millisecond end time upto which task records need to be fetched 
	ParentID string `url:"parentId,omitempty"` //Fetch tasks that have this parent Id 
	RootID string `url:"rootId,omitempty"` //Fetch tasks that have this root Id 
	Status string `url:"status,omitempty"` //Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS 
}


type ResponseTaskGetBusinessAPIExecutionDetails struct{
		BapiKey string `json:"bapiKey,omitempty"` // Business API Key (UUID) 
		BapiName string `json:"bapiName,omitempty"` // Name of the Business API 
		BapiExecutionID string `json:"bapiExecutionId,omitempty"` // Execution Id of the Business API (UUID) 
		StartTime string `json:"startTime,omitempty"` // Execution Start Time of the Business API (Date Time Format) 
		StartTimeEpoch *int `json:"startTimeEpoch,omitempty"` // Execution Start Time of the Business API (Epoch Milliseconds) 
		EndTime string `json:"endTime,omitempty"` // Execution End Time of the Business API (Date Time Format) 
		EndTimeEpoch *int `json:"endTimeEpoch,omitempty"` // Execution End Time of the Business API (Epoch Milliseconds) 
		TimeDuration *int `json:"timeDuration,omitempty"` // Time taken for Business API Execution (Milliseconds) 
		Status string `json:"status,omitempty"` // Execution status of the Business API 
		RuntimeInstanceID string `json:"runtimeInstanceId,omitempty"` // Pod Id in which the Business API is executed 
		BapiError string `json:"bapiError,omitempty"` // Business API error message 
}
type ResponseTaskGetTasks2 struct{
		Response *[]ResponseTaskGetTasks2Response `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTasks2Response struct{
		AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` // 
		Data string `json:"data,omitempty"` // 
		EndTime *int `json:"endTime,omitempty"` // 
		ErrorCode string `json:"errorCode,omitempty"` // 
		ErrorKey string `json:"errorKey,omitempty"` // 
		FailureReason string `json:"failureReason,omitempty"` // 
		ID string `json:"id,omitempty"` // 
		InstanceTenantID string `json:"instanceTenantId,omitempty"` // 
		IsError *bool `json:"isError,omitempty"` // 
		LastUpdate *int `json:"lastUpdate,omitempty"` // 
		OperationIDList []string `json:"operationIdList,omitempty"` // 
		ParentID string `json:"parentId,omitempty"` // 
		Progress string `json:"progress,omitempty"` // 
		RootID string `json:"rootId,omitempty"` // 
		ServiceType string `json:"serviceType,omitempty"` // 
		StartTime *int `json:"startTime,omitempty"` // 
		Username string `json:"username,omitempty"` // 
		Version *int `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskCount struct{
		Response *int `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskByOperationID struct{
		Response *[]ResponseTaskGetTaskByOperationIDResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskByOperationIDResponse struct{
		AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` // 
		Data string `json:"data,omitempty"` // 
		EndTime *int `json:"endTime,omitempty"` // 
		ErrorCode string `json:"errorCode,omitempty"` // 
		ErrorKey string `json:"errorKey,omitempty"` // 
		FailureReason string `json:"failureReason,omitempty"` // 
		ID string `json:"id,omitempty"` // 
		InstanceTenantID string `json:"instanceTenantId,omitempty"` // 
		IsError *bool `json:"isError,omitempty"` // 
		LastUpdate *int `json:"lastUpdate,omitempty"` // 
		OperationIDList []string `json:"operationIdList,omitempty"` // 
		ParentID string `json:"parentId,omitempty"` // 
		Progress string `json:"progress,omitempty"` // 
		RootID string `json:"rootId,omitempty"` // 
		ServiceType string `json:"serviceType,omitempty"` // 
		StartTime *int `json:"startTime,omitempty"` // 
		Username string `json:"username,omitempty"` // 
		Version *int `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskByID struct{
		Response *ResponseTaskGetTaskByIDResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskByIDResponse struct{
		AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` // 
		Data string `json:"data,omitempty"` // 
		EndTime *int `json:"endTime,omitempty"` // 
		ErrorCode string `json:"errorCode,omitempty"` // 
		ErrorKey string `json:"errorKey,omitempty"` // 
		FailureReason string `json:"failureReason,omitempty"` // 
		ID string `json:"id,omitempty"` // 
		InstanceTenantID string `json:"instanceTenantId,omitempty"` // 
		IsError *bool `json:"isError,omitempty"` // 
		LastUpdate *int `json:"lastUpdate,omitempty"` // 
		OperationIDList []string `json:"operationIdList,omitempty"` // 
		ParentID string `json:"parentId,omitempty"` // 
		Progress string `json:"progress,omitempty"` // 
		RootID string `json:"rootId,omitempty"` // 
		ServiceType string `json:"serviceType,omitempty"` // 
		StartTime *int `json:"startTime,omitempty"` // 
		Username string `json:"username,omitempty"` // 
		Version *int `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskTree struct{
		Response *[]ResponseTaskGetTaskTreeResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskTreeResponse struct{
		AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` // 
		Data string `json:"data,omitempty"` // 
		EndTime *int `json:"endTime,omitempty"` // 
		ErrorCode string `json:"errorCode,omitempty"` // 
		ErrorKey string `json:"errorKey,omitempty"` // 
		FailureReason string `json:"failureReason,omitempty"` // 
		ID string `json:"id,omitempty"` // 
		InstanceTenantID string `json:"instanceTenantId,omitempty"` // 
		IsError *bool `json:"isError,omitempty"` // 
		LastUpdate *int `json:"lastUpdate,omitempty"` // 
		OperationIDList []string `json:"operationIdList,omitempty"` // 
		ParentID string `json:"parentId,omitempty"` // 
		Progress string `json:"progress,omitempty"` // 
		RootID string `json:"rootId,omitempty"` // 
		ServiceType string `json:"serviceType,omitempty"` // 
		StartTime *int `json:"startTime,omitempty"` // 
		Username string `json:"username,omitempty"` // 
		Version *int `json:"version,omitempty"` // 
}
type ResponseTaskGetTasks struct{
		Response *[]ResponseTaskGetTasksResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTasksResponse struct{
		EndTime *int `json:"endTime,omitempty"` // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds 
		ID string `json:"id,omitempty"` // The ID of this task 
		UpdatedTime *int `json:"updatedTime,omitempty"` // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds 
		ParentID string `json:"parentId,omitempty"` // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId. 
		ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found 
		RootID string `json:"rootId,omitempty"` // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task. 
		StartTime *int `json:"startTime,omitempty"` // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds 
		Status string `json:"status,omitempty"` // 
}
type ResponseTaskGetTasksCount struct{
		Response *ResponseTaskGetTasksCountResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // The version of the response 
}
type ResponseTaskGetTasksCountResponse struct{
		Count *int `json:"count,omitempty"` // The reported count. 
}
type ResponseTaskGetTasksByID struct{
		Response *ResponseTaskGetTasksByIDResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTasksByIDResponse struct{
		EndTime *int `json:"endTime,omitempty"` // An approximate time of when this task has been marked completed; as measured in Unix epoch time in milliseconds 
		ID string `json:"id,omitempty"` // The ID of this task 
		UpdatedTime *int `json:"updatedTime,omitempty"` // A timestamp of when this task was last updated; as measured in Unix epoch time in milliseconds 
		ParentID string `json:"parentId,omitempty"` // The ID of the parent task if this happens to be a subtask. In case this task is not a subtask, then the parentId is expected to be null.  To construct a task tree, this task will be the child of the task with the ID listed here, or the root of the tree if this task has no parentId. 
		ResultLocation string `json:"resultLocation,omitempty"` // A server-relative URL indicating where additional task-specific details may be found 
		RootID string `json:"rootId,omitempty"` // The ID of the task representing the root node of the tree which this task belongs to.  In some cases, this may be the same as the ID or null, which indicates that this task is the root task. 
		StartTime *int `json:"startTime,omitempty"` // An approximate time of when the task creation was triggered; as measured in Unix epoch time in milliseconds 
		Status string `json:"status,omitempty"` // 
}
type ResponseTaskGetTaskDetailsByID struct{
		Response *ResponseTaskGetTaskDetailsByIDResponse `json:"response,omitempty"` // 
		Version string `json:"version,omitempty"` // 
}
type ResponseTaskGetTaskDetailsByIDResponse struct{
		Data string `json:"data,omitempty"` // Any data associated with this task; the value may vary significantly across different tasks 
		Progress string `json:"progress,omitempty"` // A textual representation which indicates the progress of this task; the value may vary significantly across different tasks 
		ErrorCode string `json:"errorCode,omitempty"` // An error code if in case this task has failed in its execution 
		FailureReason string `json:"failureReason,omitempty"` // A textual description indicating the reason why a task has failed 
}
	
//GetBusinessAPIExecutionDetails Get Business API Execution Details - c1bc-a8c1-41fb-9f75
/* Retrieves the execution details of a Business API


@param executionID executionId path parameter. Execution Id of API


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-business-api-execution-details
*/
func (s *TaskService) GetBusinessAPIExecutionDetails(executionID string ) (*ResponseTaskGetBusinessAPIExecutionDetails, *resty.Response, error) {
	path:= "/dna/intent/api/v1/dnacaap/management/execution-status/{executionId}"
	path = strings.Replace(path, "{executionId}", fmt.Sprintf("%v", executionID), -1)
	
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			SetResult(&ResponseTaskGetBusinessAPIExecutionDetails{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetBusinessAPIExecutionDetails(executionID)
			}
		return  nil,  response, fmt.Errorf("error with operation GetBusinessApiExecutionDetails")
	}

			
	result := response.Result().(*ResponseTaskGetBusinessAPIExecutionDetails)
	return result, response, err
	
			

		
}
	
//GetTasks2 Get tasks - e78b-b8a2-449b-9eed
/* Returns task(s) based on filter criteria


@param GetTasks2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks2
*/
func (s *TaskService) GetTasks2(GetTasks2QueryParams *GetTasks2QueryParams) (*ResponseTaskGetTasks2, *resty.Response, error) {
	path:= "/dna/intent/api/v1/task"
	
	queryString, _ := query.Values(GetTasks2QueryParams)
		
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			
	SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasks2{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTasks2(GetTasks2QueryParams)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTasks2")
	}

			
	result := response.Result().(*ResponseTaskGetTasks2)
	return result, response, err
	
			

		
}
	
//GetTaskCount Get task count - 26b4-4ab0-4649-a183
/* Returns Task count


@param GetTaskCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-count
*/
func (s *TaskService) GetTaskCount(GetTaskCountQueryParams *GetTaskCountQueryParams) (*ResponseTaskGetTaskCount, *resty.Response, error) {
	path:= "/dna/intent/api/v1/task/count"
	
	queryString, _ := query.Values(GetTaskCountQueryParams)
		
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			
	SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTaskCount{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTaskCount(GetTaskCountQueryParams)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTaskCount")
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
func (s *TaskService) GetTaskByOperationID(operationID string,offset int,limit int ) (*ResponseTaskGetTaskByOperationID, *resty.Response, error) {
	path:= "/dna/intent/api/v1/task/operation/{operationId}/{offset}/{limit}"
	path = strings.Replace(path, "{operationId}", fmt.Sprintf("%v", operationID), -1)
	path = strings.Replace(path, "{offset}", fmt.Sprintf("%v", offset), -1)
	path = strings.Replace(path, "{limit}", fmt.Sprintf("%v", limit), -1)
	
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			SetResult(&ResponseTaskGetTaskByOperationID{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTaskByOperationID(operationID,offset,limit)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTaskByOperationId")
	}

			
	result := response.Result().(*ResponseTaskGetTaskByOperationID)
	return result, response, err
	
			

		
}
	
//GetTaskByID Get task by Id - a1a9-3873-46ba-92b1
/* Returns a task by specified id


@param taskID taskId path parameter. UUID of the Task


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-by-id
*/
func (s *TaskService) GetTaskByID(taskID string ) (*ResponseTaskGetTaskByID, *resty.Response, error) {
	path:= "/dna/intent/api/v1/task/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)
	
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			SetResult(&ResponseTaskGetTaskByID{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTaskByID(taskID)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTaskById")
	}

			
	result := response.Result().(*ResponseTaskGetTaskByID)
	return result, response, err
	
			

		
}
	
//GetTaskTree Get task tree - f5a2-69c4-4f2a-95fa
/* Returns a task with its children tasks by based on their id


@param taskID taskId path parameter. UUID of the Task


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-tree
*/
func (s *TaskService) GetTaskTree(taskID string ) (*ResponseTaskGetTaskTree, *resty.Response, error) {
	path:= "/dna/intent/api/v1/task/{taskId}/tree"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)
	
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			SetResult(&ResponseTaskGetTaskTree{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTaskTree(taskID)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTaskTree")
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
	path:= "/dna/intent/api/v1/tasks"
	
	queryString, _ := query.Values(GetTasksQueryParams)
		
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			
	SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasks{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTasks(GetTasksQueryParams)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTasks")
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
	path:= "/dna/intent/api/v1/tasks/count"
	
	queryString, _ := query.Values(GetTasksCountQueryParams)
		
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			
	SetQueryString(queryString.Encode()).SetResult(&ResponseTaskGetTasksCount{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTasksCount(GetTasksCountQueryParams)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTasksCount")
	}

			
	result := response.Result().(*ResponseTaskGetTasksCount)
	return result, response, err
	
			

		
}
	
//GetTasksByID Get tasks by ID - e493-ea85-4038-8183
/* Returns the task with the given ID


@param id id path parameter. the 'id' of the task to retrieve


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tasks-by-id
*/
func (s *TaskService) GetTasksByID(id string ) (*ResponseTaskGetTasksByID, *resty.Response, error) {
	path:= "/dna/intent/api/v1/tasks/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			SetResult(&ResponseTaskGetTasksByID{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTasksByID(id)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTasksById")
	}

			
	result := response.Result().(*ResponseTaskGetTasksByID)
	return result, response, err
	
			

		
}
	
//GetTaskDetailsByID Get task details by ID - 408d-8acf-43fb-92c2
/* Returns the task details for the given task ID


@param id id path parameter. the 'id' of the task to retrieve details for


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-details-by-id
*/
func (s *TaskService) GetTaskDetailsByID(id string ) (*ResponseTaskGetTaskDetailsByID, *resty.Response, error) {
	path:= "/dna/intent/api/v1/tasks/{id}/detail"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	
	response, err := s.client.R().
	SetHeader("Content-Type", "application/json").
SetHeader("Accept", "application/json").
			SetResult(&ResponseTaskGetTaskDetailsByID{}).
	SetError(&Error).
	Get(path)

		
	if err != nil {return nil, nil, err
			
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
				return s.GetTaskDetailsByID(id)
			}
		return  nil,  response, fmt.Errorf("error with operation GetTaskDetailsById")
	}

			
	result := response.Result().(*ResponseTaskGetTaskDetailsByID)
	return result, response, err
	
			

		
}

	
	
	
	
	
	
	
	
	
	

	
	
	
	
	
	
	
	
	
	

	
	
	
	
	
	
	
	
	
	

	
	
	
	
	
	
	
	
	
	
