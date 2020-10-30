package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// TaskService is the service to communicate with the Task API endpoint
type TaskService service

// CountResult is the CountResult definition
type CountResult struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// Response is the Response definition
type Response struct {
	AdditionalStatusURL string `json:"additionalStatusURL,omitempty"` //
	Data                string `json:"data,omitempty"`                //
	EndTime             string `json:"endTime,omitempty"`             //
	ErrorCode           string `json:"errorCode,omitempty"`           //
	ErrorKey            string `json:"errorKey,omitempty"`            //
	FailureReason       string `json:"failureReason,omitempty"`       //
	Id                  string `json:"id,omitempty"`                  //
	InstanceTenantId    string `json:"instanceTenantId,omitempty"`    //
	IsError             bool   `json:"isError,omitempty"`             //
	LastUpdate          string `json:"lastUpdate,omitempty"`          //
	OperationIdList     string `json:"operationIdList,omitempty"`     //
	ParentId            string `json:"parentId,omitempty"`            //
	Progress            string `json:"progress,omitempty"`            //
	RootId              string `json:"rootId,omitempty"`              //
	ServiceType         string `json:"serviceType,omitempty"`         //
	StartTime           string `json:"startTime,omitempty"`           //
	Username            string `json:"username,omitempty"`            //
	Version             int    `json:"version,omitempty"`             //
}

// TaskDTOListResponse is the TaskDTOListResponse definition
type TaskDTOListResponse struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// TaskDTOResponse is the TaskDTOResponse definition
type TaskDTOResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetTaskById getTaskById
/* Returns a task by specified id
@param taskId UUID of the Task
*/
func (s *TaskService) GetTaskById(taskId string) (*TaskDTOResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/task/{taskId}"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	response, err := RestyClient.R().
		SetResult(&TaskDTOResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskDTOResponse)
	return result, response, err

}

// GetTaskByOperationId getTaskByOperationId
/* Returns root tasks associated with an Operationid
@param operationId operationId
@param offset Index, minimum value is 0
@param limit The maximum value of {limit} supported is 500. <br/> Base 1 indexing for {limit}, minimum value is 1
*/
func (s *TaskService) GetTaskByOperationId(operationId string, offset int, limit int) (*TaskDTOListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/task/operation/{operationId}/{offset}/{limit}"
	path = strings.Replace(path, "{"+"operationId"+"}", fmt.Sprintf("%v", operationId), -1)
	path = strings.Replace(path, "{"+"offset"+"}", fmt.Sprintf("%v", offset), -1)
	path = strings.Replace(path, "{"+"limit"+"}", fmt.Sprintf("%v", limit), -1)

	response, err := RestyClient.R().
		SetResult(&TaskDTOListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskDTOListResponse)
	return result, response, err

}

// GetTaskCountQueryParams defines the query parameters for this request
type GetTaskCountQueryParams struct {
	StartTime     string `url:"startTime,omitempty"`     // This is the epoch start time from which tasks need to be fetched
	EndTime       string `url:"endTime,omitempty"`       // This is the epoch end time upto which audit records need to be fetched
	Data          string `url:"data,omitempty"`          // Fetch tasks that contains this data
	ErrorCode     string `url:"errorCode,omitempty"`     // Fetch tasks that have this error code
	ServiceType   string `url:"serviceType,omitempty"`   // Fetch tasks with this service type
	Username      string `url:"username,omitempty"`      // Fetch tasks with this username
	Progress      string `url:"progress,omitempty"`      // Fetch tasks that contains this progress
	IsError       string `url:"isError,omitempty"`       // Fetch tasks ended as success or failure. Valid values: true, false
	FailureReason string `url:"failureReason,omitempty"` // Fetch tasks that contains this failure reason
	ParentId      string `url:"parentId,omitempty"`      // Fetch tasks that have this parent Id
}

// GetTaskCount getTaskCount
/* Returns Task count
@param startTime This is the epoch start time from which tasks need to be fetched
@param endTime This is the epoch end time upto which audit records need to be fetched
@param data Fetch tasks that contains this data
@param errorCode Fetch tasks that have this error code
@param serviceType Fetch tasks with this service type
@param username Fetch tasks with this username
@param progress Fetch tasks that contains this progress
@param isError Fetch tasks ended as success or failure. Valid values: true, false
@param failureReason Fetch tasks that contains this failure reason
@param parentId Fetch tasks that have this parent Id
*/
func (s *TaskService) GetTaskCount(getTaskCountQueryParams *GetTaskCountQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/task/count"

	queryString, _ := query.Values(getTaskCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&CountResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CountResult)
	return result, response, err

}

// GetTaskTree getTaskTree
/* Returns a task with its children tasks by based on their id
@param taskId UUID of the Task
*/
func (s *TaskService) GetTaskTree(taskId string) (*TaskDTOListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/task/{taskId}/tree"
	path = strings.Replace(path, "{"+"taskId"+"}", fmt.Sprintf("%v", taskId), -1)

	response, err := RestyClient.R().
		SetResult(&TaskDTOListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskDTOListResponse)
	return result, response, err

}

// GetTasksQueryParams defines the query parameters for this request
type GetTasksQueryParams struct {
	StartTime     string `url:"startTime,omitempty"`     // This is the epoch start time from which tasks need to be fetched
	EndTime       string `url:"endTime,omitempty"`       // This is the epoch end time upto which audit records need to be fetched
	Data          string `url:"data,omitempty"`          // Fetch tasks that contains this data
	ErrorCode     string `url:"errorCode,omitempty"`     // Fetch tasks that have this error code
	ServiceType   string `url:"serviceType,omitempty"`   // Fetch tasks with this service type
	Username      string `url:"username,omitempty"`      // Fetch tasks with this username
	Progress      string `url:"progress,omitempty"`      // Fetch tasks that contains this progress
	IsError       string `url:"isError,omitempty"`       // Fetch tasks ended as success or failure. Valid values: true, false
	FailureReason string `url:"failureReason,omitempty"` // Fetch tasks that contains this failure reason
	ParentId      string `url:"parentId,omitempty"`      // Fetch tasks that have this parent Id
	Offset        string `url:"offset,omitempty"`        // offset
	Limit         string `url:"limit,omitempty"`         // limit
	SortBy        string `url:"sortBy,omitempty"`        // Sort results by this field
	Order         string `url:"order,omitempty"`         // Sort order - asc or dsc
}

// GetTasks getTasks
/* Returns task(s) based on filter criteria
@param startTime This is the epoch start time from which tasks need to be fetched
@param endTime This is the epoch end time upto which audit records need to be fetched
@param data Fetch tasks that contains this data
@param errorCode Fetch tasks that have this error code
@param serviceType Fetch tasks with this service type
@param username Fetch tasks with this username
@param progress Fetch tasks that contains this progress
@param isError Fetch tasks ended as success or failure. Valid values: true, false
@param failureReason Fetch tasks that contains this failure reason
@param parentId Fetch tasks that have this parent Id
@param offset offset
@param limit limit
@param sortBy Sort results by this field
@param order Sort order - asc or dsc
*/
func (s *TaskService) GetTasks(getTasksQueryParams *GetTasksQueryParams) (*TaskDTOListResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/task"

	queryString, _ := query.Values(getTasksQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&TaskDTOListResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskDTOListResponse)
	return result, response, err

}
