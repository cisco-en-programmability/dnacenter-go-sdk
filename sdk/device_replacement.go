package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DeviceReplacementService service

type ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams struct {
	FaultyDeviceName              string   `url:"faultyDeviceName,omitempty"`              //Faulty Device Name
	FaultyDevicePlatform          string   `url:"faultyDevicePlatform,omitempty"`          //Faulty Device Platform
	ReplacementDevicePlatform     string   `url:"replacementDevicePlatform,omitempty"`     //Replacement Device Platform
	FaultyDeviceSerialNumber      string   `url:"faultyDeviceSerialNumber,omitempty"`      //Faulty Device Serial Number
	ReplacementDeviceSerialNumber string   `url:"replacementDeviceSerialNumber,omitempty"` //Replacement Device Serial Number
	ReplacementStatus             []string `url:"replacementStatus,omitempty"`             //Device Replacement status [READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED]
	Family                        []string `url:"family,omitempty"`                        //List of families[Routers, Switches and Hubs, AP]
	SortBy                        string   `url:"sortBy,omitempty"`                        //SortBy this field. SortBy is mandatory when order is used.
	SortOrder                     string   `url:"sortOrder,omitempty"`                     //Order on displayName[ASC,DESC]
	Offset                        int      `url:"offset,omitempty"`                        //offset
	Limit                         int      `url:"limit,omitempty"`                         //limit
}
type ReturnReplacementDevicesCountQueryParams struct {
	ReplacementStatus []string `url:"replacementStatus,omitempty"` //Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]
}
type RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams struct {
	Family                        string  `url:"family,omitempty"`                        //Faulty device family.
	FaultyDeviceName              string  `url:"faultyDeviceName,omitempty"`              //Faulty device name.
	FaultyDevicePlatform          string  `url:"faultyDevicePlatform,omitempty"`          //Faulty device platform.
	FaultyDeviceSerialNumber      string  `url:"faultyDeviceSerialNumber,omitempty"`      //Faulty device serial number.
	ReplacementDevicePlatform     string  `url:"replacementDevicePlatform,omitempty"`     //Replacement device platform.
	ReplacementDeviceSerialNumber string  `url:"replacementDeviceSerialNumber,omitempty"` //Replacement device serial number.
	ReplacementStatus             string  `url:"replacementStatus,omitempty"`             //Device replacement status. Available values : MARKED_FOR_REPLACEMENT, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED, READY_FOR_REPLACEMENT, REPLACEMENT_SCHEDULED, REPLACEMENT_IN_PROGRESS, REPLACED, ERROR. Replacement status: 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
	Offset                        float64 `url:"offset,omitempty"`                        //The first record to show for this page; the first record is numbered 1.
	Limit                         float64 `url:"limit,omitempty"`                         //The number of records to show for this page. Maximum value can be 500.
	SortBy                        string  `url:"sortBy,omitempty"`                        //A property within the response to sort by. Available values : id, creationTime, family, faultyDeviceId, fautyDeviceName, faultyDevicePlatform, faultyDeviceSerialNumber, replacementDevicePlatform, replacementDeviceSerialNumber, replacementTime.
	SortOrder                     string  `url:"sortOrder,omitempty"`                     //Whether ascending or descending order should be used to sort the response. Available values : ASC, DESC
}

type ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails struct {
	Response *[]ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  // Date and time of marking the device for replacement
	Family                        string `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                // Unique identifier of the faulty device
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             // Unique identifier of the neighbor device to create the DHCP server
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        // Unique identifier of network readiness task
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             // Device Replacement status
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               // Date and time of device replacement
	WorkflowID                    string `json:"workflowId,omitempty"`                    // Unique identifier of the device replacement workflow
	WorkflowFailedStep            string `json:"workflowFailedStep,omitempty"`            // Step in which the device replacement failed
	ReadinesscheckTaskID          string `json:"readinesscheckTaskId,omitempty"`          // Unique identifier of the readiness check task for the replacement device
}
type ResponseDeviceReplacementUnmarkDeviceForReplacement struct {
	Response *ResponseDeviceReplacementUnmarkDeviceForReplacementResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementUnmarkDeviceForReplacementResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDeviceReplacementMarkDeviceForReplacement struct {
	Response *ResponseDeviceReplacementMarkDeviceForReplacementResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementMarkDeviceForReplacementResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDeviceReplacementReturnReplacementDevicesCount struct {
	Response *int   `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementDeployDeviceReplacementWorkflow struct {
	Response *ResponseDeviceReplacementDeployDeviceReplacementWorkflowResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementDeployDeviceReplacementWorkflowResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflows struct {
	Response *[]ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // The version of the response.
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponse struct {
	CreationTime                  *int                                                                                          `json:"creationTime,omitempty"`                  // Time of marking the device for replacement in Unix epoch time in milliseconds
	Family                        string                                                                                        `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string                                                                                        `json:"faultyDeviceId,omitempty"`                // Faulty device id
	FaultyDeviceName              string                                                                                        `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string                                                                                        `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string                                                                                        `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string                                                                                        `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighborDeviceID              string                                                                                        `json:"neighborDeviceId,omitempty"`              // Unique identifier of the neighbor device to create the DHCP server
	ReplacementDevicePlatform     string                                                                                        `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string                                                                                        `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string                                                                                        `json:"replacementStatus,omitempty"`             // Device Replacement status. 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
	ReplacementTime               *int                                                                                          `json:"replacementTime,omitempty"`               // The Unix epoch time in milliseconds at which the device was replaced successfully
	Workflow                      *ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponseWorkflow `json:"workflow,omitempty"`                      //
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponseWorkflow struct {
	ID             string                                                                                               `json:"id,omitempty"`             // Workflow id
	Name           string                                                                                               `json:"name,omitempty"`           // Name of the workflow
	WorkflowStatus string                                                                                               `json:"workflowStatus,omitempty"` // Workflow status. 'RUNNING' - Workflow is currently in progress. 'SUCCESS' - Workflow completed successfully. 'FAILED' - Workflow completed with failure.
	StartTime      *int                                                                                                 `json:"startTime,omitempty"`      // Start time of the workflow in Unix epoch time in milliseconds
	EndTime        *int                                                                                                 `json:"endTime,omitempty"`        // Completion time of the workflow in Unix epoch time in milliseconds
	Steps          *[]ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponseWorkflowSteps `json:"steps,omitempty"`          //
}
type ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflowsResponseWorkflowSteps struct {
	Name          string `json:"name,omitempty"`          // Workflow step name
	Status        string `json:"status,omitempty"`        // Workflow step status. 'INIT' - Workflow step has not started execution. 'RUNNING' - Workflow step is currently in progress. 'SUCCESS' - Workflow step completed successfully. 'FAILED' - Workflow step completed with failure. 'ABORTED' - Workflow step aborted execution due to failure of the previous step. 'TIMEOUT' - Workflow step timedout to complete execution.
	StatusMessage string `json:"statusMessage,omitempty"` // Detailed status message for the step
	StartTime     *int   `json:"startTime,omitempty"`     // Start time of the workflow step in Unix epoch time in milliseconds
	EndTime       *int   `json:"endTime,omitempty"`       // Completion time of the workflow step in Unix epoch time in milliseconds
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice struct {
	Response *ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                                                                        `json:"version,omitempty"`  // The version of the response.
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponse struct {
	CreationTime                  *int                                                                                                                                  `json:"creationTime,omitempty"`                  // Time of marking the device for replacement in Unix epoch time in milliseconds
	Family                        string                                                                                                                                `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string                                                                                                                                `json:"faultyDeviceId,omitempty"`                // Faulty device id
	FaultyDeviceName              string                                                                                                                                `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string                                                                                                                                `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string                                                                                                                                `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string                                                                                                                                `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighborDeviceID              string                                                                                                                                `json:"neighborDeviceId,omitempty"`              // Unique identifier of the neighbor device to create the DHCP server
	ReplacementDevicePlatform     string                                                                                                                                `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string                                                                                                                                `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string                                                                                                                                `json:"replacementStatus,omitempty"`             // Device Replacement status. 'MARKED_FOR_REPLACEMENT' - The faulty device has been marked for replacement. 'NETWORK_READINESS_REQUESTED' - Initiated steps to shut down neighboring device interfaces and create a DHCP server on the uplink neighbor if the faulty device is part of a fabric setup. 'NETWORK_READINESS_FAILED' - Preparation of the network failed. Neighboring device interfaces were not shut down, and the DHCP server on the uplink neighbor was not created. 'READY_FOR_REPLACEMENT' - The network is prepared for the faulty device replacement. Neighboring device interfaces are shut down, and the DHCP server on the uplink neighbor is set up. 'REPLACEMENT_SCHEDULED' - Device replacement has been scheduled. 'REPLACEMENT_IN_PROGRESS' - Device replacement is currently in progress. 'REPLACED' - Device replacement was successful. 'ERROR' - Device replacement has failed.
	ReplacementTime               *int                                                                                                                                  `json:"replacementTime,omitempty"`               // The Unix epoch time in milliseconds at which the device was replaced successfully
	Workflow                      *ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponseWorkflow `json:"workflow,omitempty"`                      //
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponseWorkflow struct {
	ID             string                                                                                                                                       `json:"id,omitempty"`             // Workflow id
	Name           string                                                                                                                                       `json:"name,omitempty"`           // Name of the workflow
	WorkflowStatus string                                                                                                                                       `json:"workflowStatus,omitempty"` // Workflow status. 'RUNNING' - Workflow is currently in progress. 'SUCCESS' - Workflow completed successfully. 'FAILED' - Workflow completed with failure.
	StartTime      *int                                                                                                                                         `json:"startTime,omitempty"`      // Start time of the workflow in Unix epoch time in milliseconds
	EndTime        *int                                                                                                                                         `json:"endTime,omitempty"`        // Completion time of the workflow in Unix epoch time in milliseconds
	Steps          *[]ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponseWorkflowSteps `json:"steps,omitempty"`          //
}
type ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDeviceResponseWorkflowSteps struct {
	Name          string `json:"name,omitempty"`          // Workflow step name
	Status        string `json:"status,omitempty"`        // Workflow step status. 'INIT' - Workflow step has not started execution. 'RUNNING' - Workflow step is currently in progress. 'SUCCESS' - Workflow step completed successfully. 'FAILED' - Workflow step completed with failure. 'ABORTED' - Workflow step aborted execution due to failure of the previous step. 'TIMEOUT' - Workflow step timedout to complete execution.
	StatusMessage string `json:"statusMessage,omitempty"` // Detailed status message for the step
	StartTime     *int   `json:"startTime,omitempty"`     // Start time of the workflow step in Unix epoch time in milliseconds
	EndTime       *int   `json:"endTime,omitempty"`       // Completion time of the workflow step in Unix epoch time in milliseconds
}
type RequestDeviceReplacementUnmarkDeviceForReplacement []RequestItemDeviceReplacementUnmarkDeviceForReplacement // Array of RequestDeviceReplacementUnMarkDeviceForReplacement
type RequestItemDeviceReplacementUnmarkDeviceForReplacement struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  // Date and time of marking the device for replacement
	Family                        string `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                // Unique identifier of the faulty device
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             // Unique identifier of the neighbor device to create the DHCP server
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        // Unique identifier of network readiness task
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             // Device replacement status. Use NON-FAULTY to unmark the device for replacement.
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               // Date and time of device replacement
	WorkflowID                    string `json:"workflowId,omitempty"`                    // Unique identifier of the device replacement workflow
}
type RequestDeviceReplacementMarkDeviceForReplacement []RequestItemDeviceReplacementMarkDeviceForReplacement // Array of RequestDeviceReplacementMarkDeviceForReplacement
type RequestItemDeviceReplacementMarkDeviceForReplacement struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  // Date and time of marking the device for replacement
	Family                        string `json:"family,omitempty"`                        // Faulty device family
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                // Unique identifier of the faulty device
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              // Faulty device name
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          // Faulty device platform
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ID                            string `json:"id,omitempty"`                            // Unique identifier of the device replacement resource
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             // Unique identifier of the neighbor device to create the DHCP server
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        // Unique identifier of network readiness task
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     // Replacement device platform
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             // Device replacement status. Use MARKED-FOR-REPLACEMENT to mark the device for replacement.
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               // Date and time of device replacement
	WorkflowID                    string `json:"workflowId,omitempty"`                    // Unique identifier of the device replacement workflow
}
type RequestDeviceReplacementDeployDeviceReplacementWorkflow struct {
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      // Faulty device serial number
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` // Replacement device serial number
}

//ReturnListOfReplacementDevicesWithReplacementDetails Return list of replacement devices with replacement details - 809c-2956-4bc9-97d0
/* Get list of replacement devices with replacement details and it can filter replacement devices based on Faulty Device Name,Faulty Device Platform, Replacement Device Platform, Faulty Device Serial Number,Replacement Device Serial Number, Device Replacement status, Product Family.


@param ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-list-of-replacement-devices-with-replacement-details
*/
func (s *DeviceReplacementService) ReturnListOfReplacementDevicesWithReplacementDetails(ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams *ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams) (*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement"

	queryString, _ := query.Values(ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnListOfReplacementDevicesWithReplacementDetails(ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnListOfReplacementDevicesWithReplacementDetails")
	}

	result := response.Result().(*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails)
	return result, response, err

}

//ReturnReplacementDevicesCount Return replacement devices count - 9eb8-4ba5-4929-a2a2
/* Get replacement devices count


@param ReturnReplacementDevicesCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-replacement-devices-count
*/
func (s *DeviceReplacementService) ReturnReplacementDevicesCount(ReturnReplacementDevicesCountQueryParams *ReturnReplacementDevicesCountQueryParams) (*ResponseDeviceReplacementReturnReplacementDevicesCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement/count"

	queryString, _ := query.Values(ReturnReplacementDevicesCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceReplacementReturnReplacementDevicesCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnReplacementDevicesCount(ReturnReplacementDevicesCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnReplacementDevicesCount")
	}

	result := response.Result().(*ResponseDeviceReplacementReturnReplacementDevicesCount)
	return result, response, err

}

//RetrieveTheStatusOfAllTheDeviceReplacementWorkflows Retrieve the status of all the device replacement workflows. - e6b8-0a1a-4929-a7a9
/* Retrieve the list of device replacements with replacement details. Filters can be applied based on faulty device name, faulty device platform, faulty device serial number, replacement device platform, replacement device serial number, device replacement status, device family.


@param RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-status-of-all-the-device-replacement-workflows
*/
func (s *DeviceReplacementService) RetrieveTheStatusOfAllTheDeviceReplacementWorkflows(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams *RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams) (*ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflows, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceReplacements"

	queryString, _ := query.Values(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflows{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheStatusOfAllTheDeviceReplacementWorkflows(RetrieveTheStatusOfAllTheDeviceReplacementWorkflowsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheStatusOfAllTheDeviceReplacementWorkflows")
	}

	result := response.Result().(*ResponseDeviceReplacementRetrieveTheStatusOfAllTheDeviceReplacementWorkflows)
	return result, response, err

}

//RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice Retrieve the status of device replacement workflow that replaces a faulty device with a replacement device. - 92ba-aa03-43c8-9d62
/* Fetches the status of the device replacement workflow for a given device replacement `id`. Invoke the API `/dna/intent/api/v1/networkDeviceReplacements` to `GET` the list of all device replacements and use the `id` field data as input to this API.


@param id id path parameter. Instance UUID of the device replacement


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-status-of-device-replacement-workflow-that-replaces-a-faulty-device-with-a-replacement-device
*/
func (s *DeviceReplacementService) RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice(id string) (*ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceReplacements/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice")
	}

	result := response.Result().(*ResponseDeviceReplacementRetrieveTheStatusOfDeviceReplacementWorkflowThatReplacesAFaultyDeviceWithAReplacementDevice)
	return result, response, err

}

//MarkDeviceForReplacement Mark device for replacement - 64b9-dad0-403a-aca1
/* Marks device for replacement



Documentation Link: https://developer.cisco.com/docs/dna-center/#!mark-device-for-replacement
*/
func (s *DeviceReplacementService) MarkDeviceForReplacement(requestDeviceReplacementMarkDeviceForReplacement *RequestDeviceReplacementMarkDeviceForReplacement) (*ResponseDeviceReplacementMarkDeviceForReplacement, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceReplacementMarkDeviceForReplacement).
		SetResult(&ResponseDeviceReplacementMarkDeviceForReplacement{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.MarkDeviceForReplacement(requestDeviceReplacementMarkDeviceForReplacement)
		}

		return nil, response, fmt.Errorf("error with operation MarkDeviceForReplacement")
	}

	result := response.Result().(*ResponseDeviceReplacementMarkDeviceForReplacement)
	return result, response, err

}

//DeployDeviceReplacementWorkflow Deploy device replacement workflow - 3faa-a994-4b49-bc9f
/* API to trigger RMA workflow that will replace faulty device with replacement device with same configuration and images



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-device-replacement-workflow
*/
func (s *DeviceReplacementService) DeployDeviceReplacementWorkflow(requestDeviceReplacementDeployDeviceReplacementWorkflow *RequestDeviceReplacementDeployDeviceReplacementWorkflow) (*ResponseDeviceReplacementDeployDeviceReplacementWorkflow, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement/workflow"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceReplacementDeployDeviceReplacementWorkflow).
		SetResult(&ResponseDeviceReplacementDeployDeviceReplacementWorkflow{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployDeviceReplacementWorkflow(requestDeviceReplacementDeployDeviceReplacementWorkflow)
		}

		return nil, response, fmt.Errorf("error with operation DeployDeviceReplacementWorkflow")
	}

	result := response.Result().(*ResponseDeviceReplacementDeployDeviceReplacementWorkflow)
	return result, response, err

}

//UnmarkDeviceForReplacement UnMark device for replacement - 4aba-ba75-489a-b24b
/* UnMarks device for replacement


 */
func (s *DeviceReplacementService) UnmarkDeviceForReplacement(requestDeviceReplacementUnMarkDeviceForReplacement *RequestDeviceReplacementUnmarkDeviceForReplacement) (*ResponseDeviceReplacementUnmarkDeviceForReplacement, *resty.Response, error) {
	path := "/dna/intent/api/v1/device-replacement"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestDeviceReplacementUnMarkDeviceForReplacement).
		SetResult(&ResponseDeviceReplacementUnmarkDeviceForReplacement{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnmarkDeviceForReplacement(requestDeviceReplacementUnMarkDeviceForReplacement)
		}
		return nil, response, fmt.Errorf("error with operation UnmarkDeviceForReplacement")
	}

	result := response.Result().(*ResponseDeviceReplacementUnmarkDeviceForReplacement)
	return result, response, err

}
