package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type DeviceReplacementService service

type ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams struct {
	FaultyDeviceName              string `url:"faultyDeviceName,omitempty"`              //Faulty Device Name
	FaultyDevicePlatform          string `url:"faultyDevicePlatform,omitempty"`          //Faulty Device Platform
	ReplacementDevicePlatform     string `url:"replacementDevicePlatform,omitempty"`     //Replacement Device Platform
	FaultyDeviceSerialNumber      string `url:"faultyDeviceSerialNumber,omitempty"`      //Faulty Device Serial Number
	ReplacementDeviceSerialNumber string `url:"replacementDeviceSerialNumber,omitempty"` //Replacement Device Serial Number
	ReplacementStatus             string `url:"replacementStatus,omitempty"`             //Device Replacement status [READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED]
	Family                        string `url:"family,omitempty"`                        //List of families[Routers, Switches and Hubs, AP]
	SortBy                        string `url:"sortBy,omitempty"`                        //SortBy this field. SortBy is mandatory when order is used.
	SortOrder                     string `url:"sortOrder,omitempty"`                     //Order on displayName[ASC,DESC]
	Offset                        int    `url:"offset,omitempty"`                        //offset
	Limit                         int    `url:"limit,omitempty"`                         //limit
}
type ReturnReplacementDevicesCountQueryParams struct {
	ReplacementStatus string `url:"replacementStatus,omitempty"` //Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]
}

type ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails struct {
	Response *[]ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse `json:"response,omitempty"` //
	Version  string                                                                                   `json:"version,omitempty"`  //
}
type ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetailsResponse struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  //
	Family                        string `json:"family,omitempty"`                        //
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                //
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              //
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          //
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      //
	ID                            string `json:"id,omitempty"`                            //
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             //
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        //
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     //
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` //
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             //
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               //
	WorkflowID                    string `json:"workflowId,omitempty"`                    //
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
type RequestDeviceReplacementUnmarkDeviceForReplacement []RequestItemDeviceReplacementUnmarkDeviceForReplacement // Array of RequestDeviceReplacementUnMarkDeviceForReplacement
type RequestItemDeviceReplacementUnmarkDeviceForReplacement struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  //
	Family                        string `json:"family,omitempty"`                        //
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                //
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              //
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          //
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      //
	ID                            string `json:"id,omitempty"`                            //
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             //
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        //
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     //
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` //
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             //
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               //
	WorkflowID                    string `json:"workflowId,omitempty"`                    //
}
type RequestDeviceReplacementMarkDeviceForReplacement []RequestItemDeviceReplacementMarkDeviceForReplacement // Array of RequestDeviceReplacementMarkDeviceForReplacement
type RequestItemDeviceReplacementMarkDeviceForReplacement struct {
	CreationTime                  *int   `json:"creationTime,omitempty"`                  //
	Family                        string `json:"family,omitempty"`                        //
	FaultyDeviceID                string `json:"faultyDeviceId,omitempty"`                //
	FaultyDeviceName              string `json:"faultyDeviceName,omitempty"`              //
	FaultyDevicePlatform          string `json:"faultyDevicePlatform,omitempty"`          //
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      //
	ID                            string `json:"id,omitempty"`                            //
	NeighbourDeviceID             string `json:"neighbourDeviceId,omitempty"`             //
	NetworkReadinessTaskID        string `json:"networkReadinessTaskId,omitempty"`        //
	ReplacementDevicePlatform     string `json:"replacementDevicePlatform,omitempty"`     //
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` //
	ReplacementStatus             string `json:"replacementStatus,omitempty"`             //
	ReplacementTime               *int   `json:"replacementTime,omitempty"`               //
	WorkflowID                    string `json:"workflowId,omitempty"`                    //
}
type RequestDeviceReplacementDeployDeviceReplacementWorkflow struct {
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      //
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` //
}

//ReturnListOfReplacementDevicesWithReplacementDetails Return list of replacement devices with replacement details - 809c-2956-4bc9-97d0
/* Get list of replacement devices with replacement details and it can filter replacement devices based on Faulty Device Name,Faulty Device Platform, Replacement Device Platform, Faulty Device Serial Number,Replacement Device Serial Number, Device Replacement status, Product Family.


@param ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation ReturnListOfReplacementDevicesWithReplacementDetails")
	}

	result := response.Result().(*ResponseDeviceReplacementReturnListOfReplacementDevicesWithReplacementDetails)
	return result, response, err

}

//ReturnReplacementDevicesCount Return replacement devices count - 9eb8-4ba5-4929-a2a2
/* Get replacement devices count


@param ReturnReplacementDevicesCountQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation ReturnReplacementDevicesCount")
	}

	result := response.Result().(*ResponseDeviceReplacementReturnReplacementDevicesCount)
	return result, response, err

}

//MarkDeviceForReplacement Mark device for replacement - 64b9-dad0-403a-aca1
/* Marks device for replacement


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
		return nil, response, fmt.Errorf("error with operation MarkDeviceForReplacement")
	}

	result := response.Result().(*ResponseDeviceReplacementMarkDeviceForReplacement)
	return result, response, err

}

//DeployDeviceReplacementWorkflow Deploy device replacement workflow - 3faa-a994-4b49-bc9f
/* API to trigger RMA workflow that will replace faulty device with replacement device with same configuration and images


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
		return nil, response, fmt.Errorf("error with operation UnmarkDeviceForReplacement")
	}

	result := response.Result().(*ResponseDeviceReplacementUnmarkDeviceForReplacement)
	return result, response, err

}
