package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// DeviceReplacementService is the service to communicate with the DeviceReplacement API endpoint
type DeviceReplacementService service

// DeviceReplacementDataDTO is the DeviceReplacementDataDTO definition
type DeviceReplacementDataDTO struct {
	CreationTime                  int    `json:"creationTime,omitempty"`                  //
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
	ReplacementTime               int    `json:"replacementTime,omitempty"`               //
	WorkflowID                    string `json:"workflowId,omitempty"`                    //
}

// DeviceReplacementWorkflowDTO is the DeviceReplacementWorkflowDTO definition
type DeviceReplacementWorkflowDTO struct {
	FaultyDeviceSerialNumber      string `json:"faultyDeviceSerialNumber,omitempty"`      //
	ReplacementDeviceSerialNumber string `json:"replacementDeviceSerialNumber,omitempty"` //
}

// DeviceReplacementResult is the DeviceReplacementResult definition
type DeviceReplacementResult struct {
	Response []TaskResponse `json:"response,omitempty"` //
	Version  string         `json:"version,omitempty"`  //
}

// DeployDeviceReplacementWorkflow deployDeviceReplacementWorkflow
/* API to trigger RMA workflow that will replace faulty device with replacement device with same configuration and images
 */
// func (s *DeviceReplacementService) DeployDeviceReplacementWorkflow(deployDeviceReplacementWorkflowRequest *DeployDeviceReplacementWorkflowRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/device-replacement/workflow"

// 	response, err := RestyClient.R().
// 		SetBody(deployDeviceReplacementWorkflowRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// MarkDeviceForReplacement markDeviceForReplacement
/* Marks device for replacement
 */
// func (s *DeviceReplacementService) MarkDeviceForReplacement(markDeviceForReplacementRequest *MarkDeviceForReplacementRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/device-replacement"

// 	response, err := RestyClient.R().
// 		SetBody(markDeviceForReplacementRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }

// ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams defines the query parameters for this request
type ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams struct {
	FaultyDeviceName              string   `url:"faultyDeviceName,omitempty"`              // Faulty Device Name
	FaultyDevicePlatform          string   `url:"faultyDevicePlatform,omitempty"`          // Faulty Device Platform
	ReplacementDevicePlatform     string   `url:"replacementDevicePlatform,omitempty"`     // Replacement Device Platform
	FaultyDeviceSerialNumber      string   `url:"faultyDeviceSerialNumber,omitempty"`      // Faulty Device Serial Number
	ReplacementDeviceSerialNumber string   `url:"replacementDeviceSerialNumber,omitempty"` // Replacement Device Serial Number
	ReplacementStatus             []string `url:"replacementStatus,omitempty"`             // Device Replacement status [READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED]
	Family                        []string `url:"family,omitempty"`                        // List of families[Routers, Switches and Hubs, AP]
	SortBy                        string   `url:"sortBy,omitempty"`                        // SortBy this field. SortBy is mandatory when order is used.
	SortOrder                     string   `url:"sortOrder,omitempty"`                     // Order on displayName[ASC,DESC]
	Offset                        int      `url:"offset,omitempty"`                        // offset
	Limit                         int      `url:"limit,omitempty"`                         // limit
}

// ReturnListOfReplacementDevicesWithReplacementDetails returnListOfReplacementDevicesWithReplacementDetails
/* Get list of replacement devices with replacement details and it can filter replacement devices based on Faulty Device Name,Faulty Device Platform, Replacement Device Platform, Faulty Device Serial Number,Replacement Device Serial Number, Device Replacement status, Product Family.
@param faultyDeviceName Faulty Device Name
@param faultyDevicePlatform Faulty Device Platform
@param replacementDevicePlatform Replacement Device Platform
@param faultyDeviceSerialNumber Faulty Device Serial Number
@param replacementDeviceSerialNumber Replacement Device Serial Number
@param replacementStatus Device Replacement status [READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR, NETWORK_READINESS_REQUESTED, NETWORK_READINESS_FAILED]
@param family List of families[Routers, Switches and Hubs, AP]
@param sortBy SortBy this field. SortBy is mandatory when order is used.
@param sortOrder Order on displayName[ASC,DESC]
@param offset offset
@param limit limit
*/
func (s *DeviceReplacementService) ReturnListOfReplacementDevicesWithReplacementDetails(returnListOfReplacementDevicesWithReplacementDetailsQueryParams *ReturnListOfReplacementDevicesWithReplacementDetailsQueryParams) (*DeviceReplacementResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-replacement"

	queryString, _ := query.Values(returnListOfReplacementDevicesWithReplacementDetailsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&DeviceReplacementResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DeviceReplacementResult)
	return result, response, err

}

// ReturnReplacementDevicesCountQueryParams defines the query parameters for this request
type ReturnReplacementDevicesCountQueryParams struct {
	ReplacementStatus []string `url:"replacementStatus,omitempty"` // Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]
}

// ReturnReplacementDevicesCount returnReplacementDevicesCount
/* Get replacement devices count
@param replacementStatus Device Replacement status list[READY-FOR-REPLACEMENT, REPLACEMENT-IN-PROGRESS, REPLACEMENT-SCHEDULED, REPLACED, ERROR]
*/
func (s *DeviceReplacementService) ReturnReplacementDevicesCount(returnReplacementDevicesCountQueryParams *ReturnReplacementDevicesCountQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/device-replacement/count"

	queryString, _ := query.Values(returnReplacementDevicesCountQueryParams)

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

// UnMarkDeviceForReplacement unMarkDeviceForReplacement
/* UnMarks device for replacement
 */
// func (s *DeviceReplacementService) UnMarkDeviceForReplacement(unMarkDeviceForReplacementRequest *UnMarkDeviceForReplacementRequest) (*TaskIDResult, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/device-replacement"

// 	response, err := RestyClient.R().
// 		SetBody(unMarkDeviceForReplacementRequest).
// 		SetResult(&TaskIDResult{}).
// 		SetError(&Error{}).
// 		Put(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*TaskIDResult)
// 	return result, response, err

// }
