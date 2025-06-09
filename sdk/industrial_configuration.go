package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IndustrialConfigurationService service

type RetrievesTheListOfMRPRingsQueryParams struct {
	ID     float64 `url:"id,omitempty"`     //ID of the MRP ring.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}
type RetrievesTheListOfNetworkDevicesPartOfMRPRingQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
}

type ResponseIndustrialConfigurationConfigureAREPRingOnFABRICDeployment struct {
	Response *ResponseIndustrialConfigurationConfigureAREPRingOnFABRICDeploymentResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Version of the response.
}
type ResponseIndustrialConfigurationConfigureAREPRingOnFABRICDeploymentResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task.
	URL    string `json:"url,omitempty"`    // URL for the task.
}
type ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheFABRICDeployment struct {
	Response *ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheFABRICDeploymentResponse `json:"response,omitempty"` //
	Version  string                                                                               `json:"version,omitempty"`  // Version of the response.
}
type ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheFABRICDeploymentResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task.
	URL    string `json:"url,omitempty"`    // URL for the task.
}
type ResponseIndustrialConfigurationRetrievesTheListOfMRPRings []ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRings // Array of ResponseIndustrialConfigurationRetrievesTheListOfMRPRings
type ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRings struct {
	Response *[]ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRingsResponse `json:"response,omitempty"` //
	Version  *int                                                                     `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRingsResponse struct {
	ID              *int                                                                                  `json:"id,omitempty"`              // ID of the MRP ring.
	NetworkDeviceID string                                                                                `json:"networkDeviceId,omitempty"` // Network device ID of the MRP ring member.
	RingSize        *int                                                                                  `json:"ringSize,omitempty"`        // Number of members participating in MRP ring.
	DeviceDetails   *[]ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRingsResponseDeviceDetails `json:"deviceDetails,omitempty"`   //
}
type ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRingsResponseDeviceDetails struct {
	Priority                                  *int                                                                                       `json:"priority,omitempty"`                                  // Value used in the MRM voting process.
	RecoveryTimeProfileMilliseconds           string                                                                                     `json:"recoveryTimeProfileMilliseconds,omitempty"`           // Recovery time profile in Milliseconds.
	VLANID                                    *int                                                                                       `json:"vlanId,omitempty"`                                    // VLAN ID range is 1-4094.
	BestManagerPrority                        *int                                                                                       `json:"bestManagerPrority,omitempty"`                        // Priority value of the best manager in MRP ring
	BestManagerMacAddress                     string                                                                                     `json:"bestManagerMacAddress,omitempty"`                     // MAC address of the best manager in MRP ring.
	BestManagerHostName                       string                                                                                     `json:"bestManagerHostName,omitempty"`                       // Name of the ring member who is the best manager.
	MrpLicense                                string                                                                                     `json:"mrpLicense,omitempty"`                                // MRP license type.
	DomainID                                  string                                                                                     `json:"domainId,omitempty"`                                  // A unique domain ID that represents the MRP ring.
	NetworkStatus                             string                                                                                     `json:"networkStatus,omitempty"`                             // Network status of the configured MRP ring.
	ConfiguredFrom                            string                                                                                     `json:"configuredFrom,omitempty"`                            // The MRP ring's configured mode.
	TopologyChangeRequestIntervalMilliseconds string                                                                                     `json:"topologyChangeRequestIntervalMilliseconds,omitempty"` // Time interval of MRP topology change request.
	DomainName                                string                                                                                     `json:"domainName,omitempty"`                                // Logical name of the configured MRP domain ID.
	OperationMode                             string                                                                                     `json:"operationMode,omitempty"`                             // The operating mode of the MRP member.
	ConfigurationMode                         string                                                                                     `json:"configurationMode,omitempty"`                         // The configured mode of the MRP member.
	Ports                                     *[]ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRingsResponseDeviceDetailsPorts `json:"ports,omitempty"`                                     //
}
type ResponseItemIndustrialConfigurationRetrievesTheListOfMRPRingsResponseDeviceDetailsPorts struct {
	InterfaceName  string `json:"interfaceName,omitempty"`  // Inteface name of the member.
	PortMacAddress string `json:"portMacAddress,omitempty"` // MAC address of the interface.
	PortNumber     *int   `json:"portNumber,omitempty"`     // The identifier of the MRP interface.
	PortStatus     *int   `json:"portStatus,omitempty"`     // Status of the MRP configured interface.
}
type ResponseIndustrialConfigurationRetrievesTheCountOfMRPRings []ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRings // Array of ResponseIndustrialConfigurationRetrievesTheCountOfMRPRings
type ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRings struct {
	Response *[]ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRingsResponse `json:"response,omitempty"` //
	Version  *int                                                                      `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRingsResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing []ResponseItemIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing // Array of ResponseIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing
type ResponseItemIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing struct {
	Response *[]ResponseItemIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRingResponse `json:"response,omitempty"` //
	Version  *int                                                                                        `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRingResponse struct {
	NetworkDeviceID   string `json:"networkDeviceId,omitempty"`   // Network device ID of the MRP ring member.
	ID                *int   `json:"id,omitempty"`                // ID of the MRP ring.
	DeviceName        string `json:"deviceName,omitempty"`        // Name of the MRP ring member.
	DomainID          string `json:"domainId,omitempty"`          // A unique domain ID that represents the MRP ring.
	PortName1         string `json:"portName1,omitempty"`         // Port 1 interface name of the MRP ring member.
	PortName2         string `json:"portName2,omitempty"`         // Port 2 interface name of the MRP ring member.
	Port1Status       string `json:"port1Status,omitempty"`       // Status of the MRP configured port1 interface.
	Port2Status       string `json:"port2Status,omitempty"`       // Status of the MRP configured port2 interface.
	OperationMode     string `json:"operationMode,omitempty"`     // The operating mode of the MRP ring member.
	ConfigurationMode string `json:"configurationMode,omitempty"` // The configured mode of the MRP ring member.
}
type ResponseIndustrialConfigurationRetrievesTheCountOfMRPRingMembers []ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRingMembers // Array of ResponseIndustrialConfigurationRetrievesTheCountOfMRPRingMembers
type ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRingMembers struct {
	Response *[]ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRingMembersResponse `json:"response,omitempty"` //
	Version  *int                                                                            `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationRetrievesTheCountOfMRPRingMembersResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment struct {
	Response *ResponseIndustrialConfigurationConfigureAREPRingOnNONFABRICDeploymentResponse `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version of the response.
}
type ResponseIndustrialConfigurationConfigureAREPRingOnNONFABRICDeploymentResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task.
	URL    string `json:"url,omitempty"`    // URL for the task.
}
type ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheNONFABRICDeployment struct {
	Response *ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheNONFABRICDeploymentResponse `json:"response,omitempty"` //
	Version  string                                                                                  `json:"version,omitempty"`  // Version of the response.
}
type ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheNONFABRICDeploymentResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task.
	URL    string `json:"url,omitempty"`    // URL for the task.
}
type ResponseIndustrialConfigurationRetrievesTheListOfREPRings []ResponseItemIndustrialConfigurationRetrievesTheListOfREPRings // Array of ResponseIndustrialConfigurationRetrievesTheListOfREPRings
type ResponseItemIndustrialConfigurationRetrievesTheListOfREPRings struct {
	Response *[]ResponseItemIndustrialConfigurationRetrievesTheListOfREPRingsResponse `json:"response,omitempty"` //
	Version  *int                                                                     `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationRetrievesTheListOfREPRingsResponse struct {
	ID                            string                                                                              `json:"id,omitempty"`                            // ID of REP ring.
	NetworkDeviceID               string                                                                              `json:"networkDeviceId,omitempty"`               // Network device ID of the node in the REP ring.
	RootNetworkDeviceID           string                                                                              `json:"rootNetworkDeviceId,omitempty"`           // Network device ID of the root node in the REP ring.
	RootNeighbourNetworkDeviceIDs []string                                                                            `json:"rootNeighbourNetworkDeviceIds,omitempty"` // List of network device IDs of the immediate neighboring ring members of the root node.
	Status                        string                                                                              `json:"status,omitempty"`                        // Status of the REP ring.
	RepSegmentID                  *int                                                                                `json:"repSegmentId,omitempty"`                  // REP segment is a chain of ports connected to each other and configured with a segment ID.
	DeploymentMode                string                                                                              `json:"deploymentMode,omitempty"`                // Deployment mode of the configured REP ring.
	RingName                      string                                                                              `json:"ringName,omitempty"`                      // Unique name of REP ring configured.
	RingMembers                   *[]ResponseItemIndustrialConfigurationRetrievesTheListOfREPRingsResponseRingMembers `json:"ringMembers,omitempty"`                   //
}
type ResponseItemIndustrialConfigurationRetrievesTheListOfREPRingsResponseRingMembers struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the REP ring member.
	NodeName        string `json:"nodeName,omitempty"`        // Name of the ring member.
	PortName1       string `json:"portName1,omitempty"`       // Interface name of the REP ring member.
	PortName2       string `json:"portName2,omitempty"`       // Interface name of the REP ring member.
	RingOrder       *int   `json:"ringOrder,omitempty"`       // Order of the REP ring member in the ring topology.
}
type ResponseIndustrialConfigurationRetrievesTheCountOfREPRings []ResponseItemIndustrialConfigurationRetrievesTheCountOfREPRings // Array of ResponseIndustrialConfigurationRetrievesTheCountOfREPRings
type ResponseItemIndustrialConfigurationRetrievesTheCountOfREPRings struct {
	Response *[]ResponseItemIndustrialConfigurationRetrievesTheCountOfREPRingsResponse `json:"response,omitempty"` //
	Version  *int                                                                      `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationRetrievesTheCountOfREPRingsResponse struct {
	Count *int `json:"count,omitempty"` // Number of REP rings.
}
type ResponseIndustrialConfigurationGetTheREPRingBasedOnTheRingID []ResponseItemIndustrialConfigurationGetTheREPRingBasedOnTheRingID // Array of ResponseIndustrialConfigurationGetTheREPRingBasedOnTheRingId
type ResponseItemIndustrialConfigurationGetTheREPRingBasedOnTheRingID struct {
	Response *[]ResponseItemIndustrialConfigurationGetTheREPRingBasedOnTheRingIDResponse `json:"response,omitempty"` //
	Version  *int                                                                        `json:"version,omitempty"`  // Version of the response.
}
type ResponseItemIndustrialConfigurationGetTheREPRingBasedOnTheRingIDResponse struct {
	ID                            string                                                                                 `json:"id,omitempty"`                            // ID of REP ring.
	RootNetworkDeviceID           string                                                                                 `json:"rootNetworkDeviceId,omitempty"`           // Network device ID of the root node in the REP ring.
	RootNeighbourNetworkDeviceIDs []string                                                                               `json:"rootNeighbourNetworkDeviceIds,omitempty"` // List of network device IDs of the immediate neighboring ring members of the root node.
	Status                        string                                                                                 `json:"status,omitempty"`                        // Status of the REP ring.
	RepSegmentID                  *int                                                                                   `json:"repSegmentId,omitempty"`                  // REP segment is a chain of ports connected to each other and configured with a segment ID.
	DeploymentMode                string                                                                                 `json:"deploymentMode,omitempty"`                // Deployment mode of the configured REP ring.
	RingName                      string                                                                                 `json:"ringName,omitempty"`                      // Unique name of REP ring configured.
	RingMembers                   *[]ResponseItemIndustrialConfigurationGetTheREPRingBasedOnTheRingIDResponseRingMembers `json:"ringMembers,omitempty"`                   //
}
type ResponseItemIndustrialConfigurationGetTheREPRingBasedOnTheRingIDResponseRingMembers struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device ID of the REP ring member.
	NodeName        string `json:"nodeName,omitempty"`        // Name of the ring member.
	PortName1       string `json:"portName1,omitempty"`       // Interface name of the REP ring member.
	PortName2       string `json:"portName2,omitempty"`       // Interface name of the REP ring member.
	RingOrder       *int   `json:"ringOrder,omitempty"`       // Order of the REP ring member in the ring topology.
}
type RequestIndustrialConfigurationConfigureAREPRingOnFABRICDeployment struct {
	RingName                      string   `json:"ringName,omitempty"`                      // Unique name of REP ring to be configured.
	RootNetworkDeviceID           string   `json:"rootNetworkDeviceId,omitempty"`           // rootNetworkDeviceId  is the network device ID of the root node in the REP ring. API `/dna/intent/api/v1/networkDevices` can be used to get the rootNetworkDeviceId , `instanceUuid` attribute in the response contains rootNetworkDeviceId.
	RootNeighbourNetworkDeviceIDs []string `json:"rootNeighbourNetworkDeviceIds,omitempty"` // It contains the network device IDs of the immediate neighboring ring members of the root node. API `/dna/intent/api/v1/networkDevices` can be used to get the list of networkDeviceIds of the neighbors , `instanceUuid` attribute in the response contains rootNeighbourNetworkDeviceIds.
	DeploymentMode                string   `json:"deploymentMode,omitempty"`                // Deployment mode of the configured REP ring.
}
type RequestIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment struct {
	RingName                      string   `json:"ringName,omitempty"`                      // Unique name of REP ring to be configured.
	RootNetworkDeviceID           string   `json:"rootNetworkDeviceId,omitempty"`           // rootNetworkDeviceId  is the network device ID of the root node in the REP ring. API `/dna/intent/api/v1/networkDevices` can be used to get the rootNetworkDeviceId , `instanceUuid` attribute in the response contains rootNetworkDeviceId.
	RootNeighbourNetworkDeviceIDs []string `json:"rootNeighbourNetworkDeviceIds,omitempty"` // It contains the network device IDs of the immediate neighboring ring members of the root node. API `/dna/intent/api/v1/networkDevices` can be used to get the list of networkDeviceIds of the neighbors , `instanceUuid` attribute in the response contains rootNeighbourNetworkDeviceIds.
	DeploymentMode                string   `json:"deploymentMode,omitempty"`                // Deployment mode of the configured REP ring.
}
type RequestIndustrialConfigurationRetrievesTheListOfREPRings struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device id of the REP ring member. API `/dna/intent/api/v1/networkDevices` can be used to get the list of networkDeviceIds of the neighbors , `instanceUuid` attribute in the response contains networkDeviceId.
	DeploymentMode  string `json:"deploymentMode,omitempty"`  // Deployment mode of the configured REP ring.
	Limit           *int   `json:"limit,omitempty"`           // The number of records to show for this page.
	Offset          *int   `json:"offset,omitempty"`          // The first record to show for this page; the first record is numbered 1.
}
type RequestIndustrialConfigurationRetrievesTheCountOfREPRings struct {
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device id of the REP ring member. API `/dna/intent/api/v1/networkDevices` can be used to get the list of networkDeviceIds of the neighbors , `instanceUuid` attribute in the response contains networkDeviceId.
	DeploymentMode  string `json:"deploymentMode,omitempty"`  // Deployment mode of the configured REP ring.
}

//RetrievesTheListOfMRPRings Retrieves the list of MRP rings. - 50ac-4820-4cd9-9a70
/* This API returns the list of all the MRP rings configured on the Network device when Ring ID is not specified and returns the details of a single MRP ring when Ring ID is specified based on the given fields networkDeviceId (Network device ID of the MRP ring member. The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices) and id (ID of the MRP ring. The id of the configured MRP Ring can be retrieved using the API /dna/intent/api/v1/iot/networkDevices/${networkDeviceId}/mrpRings).


@param networkDeviceID networkDeviceId path parameter. Network device ID of the MRP ring member.

@param RetrievesTheListOfMRPRingsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-m-r-p-rings
*/
func (s *IndustrialConfigurationService) RetrievesTheListOfMRPRings(networkDeviceID string, RetrievesTheListOfMRPRingsQueryParams *RetrievesTheListOfMRPRingsQueryParams) (*ResponseIndustrialConfigurationRetrievesTheListOfMRPRings, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/networkDevices/{networkDeviceId}/mrpRings"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	queryString, _ := query.Values(RetrievesTheListOfMRPRingsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIndustrialConfigurationRetrievesTheListOfMRPRings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfMRPRings(networkDeviceID, RetrievesTheListOfMRPRingsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfMRPRings")
	}

	result := response.Result().(*ResponseIndustrialConfigurationRetrievesTheListOfMRPRings)
	return result, response, err

}

//RetrievesTheCountOfMRPRings Retrieves the count of MRP rings. - 1280-4888-4cd8-b2f9
/* This API returns the count of MRP rings for the given fields networkDeviceId (Network device ID of the MRP ring member. The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices).


@param networkDeviceID networkDeviceId path parameter. Network device ID of the MRP ring member.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-m-r-p-rings
*/
func (s *IndustrialConfigurationService) RetrievesTheCountOfMRPRings(networkDeviceID string) (*ResponseIndustrialConfigurationRetrievesTheCountOfMRPRings, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/networkDevices/{networkDeviceId}/mrpRings/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIndustrialConfigurationRetrievesTheCountOfMRPRings{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfMRPRings(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfMRPRings")
	}

	result := response.Result().(*ResponseIndustrialConfigurationRetrievesTheCountOfMRPRings)
	return result, response, err

}

//RetrievesTheListOfNetworkDevicesPartOfMRPRing Retrieves the list of network devices part of MRP ring. - 7ba3-7b6a-4fd8-958c
/* This API returns the list of MRP ring members for the given fields networkDeviceId (Network device ID of the MRP ring member. The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices) and id (ID of the MRP ring. The id of the configured MRP Ring can be retrieved using the API /dna/intent/api/v1/iot/networkDevices/${networkDeviceId}/mrpRings).


@param networkDeviceID networkDeviceId path parameter. Network device ID of the MRP ring member.

@param id id path parameter. ID of the MRP ring.

@param RetrievesTheListOfNetworkDevicesPartOfMRPRingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-devices-part-of-m-r-p-ring
*/
func (s *IndustrialConfigurationService) RetrievesTheListOfNetworkDevicesPartOfMRPRing(networkDeviceID string, id float64, RetrievesTheListOfNetworkDevicesPartOfMRPRingQueryParams *RetrievesTheListOfNetworkDevicesPartOfMRPRingQueryParams) (*ResponseIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/networkDevices/{networkDeviceId}/mrpRings/{id}/members"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(RetrievesTheListOfNetworkDevicesPartOfMRPRingQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkDevicesPartOfMRPRing(networkDeviceID, id, RetrievesTheListOfNetworkDevicesPartOfMRPRingQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkDevicesPartOfMRPRing")
	}

	result := response.Result().(*ResponseIndustrialConfigurationRetrievesTheListOfNetworkDevicesPartOfMRPRing)
	return result, response, err

}

//RetrievesTheCountOfMRPRingMembers Retrieves the count of MRP ring members. - 929c-0833-46ba-b848
/* This API returns the count of MRP ring members for the given fields networkDeviceId (Network device ID of the MRP ring member. The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices) and id (ID of the MRP ring. The id of the configured MRP Ring can be retrieved using the API /dna/intent/api/v1/iot/networkDevices/${networkDeviceId}/mrpRings).


@param networkDeviceID networkDeviceId path parameter. Network device ID of the MRP ring member.

@param id id path parameter. ID of the MRP ring.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-m-r-p-ring-members
*/
func (s *IndustrialConfigurationService) RetrievesTheCountOfMRPRingMembers(networkDeviceID string, id float64) (*ResponseIndustrialConfigurationRetrievesTheCountOfMRPRingMembers, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/networkDevices/{networkDeviceId}/mrpRings/{id}/members/count"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIndustrialConfigurationRetrievesTheCountOfMRPRingMembers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfMRPRingMembers(networkDeviceID, id)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfMRPRingMembers")
	}

	result := response.Result().(*ResponseIndustrialConfigurationRetrievesTheCountOfMRPRingMembers)
	return result, response, err

}

//GetTheREPRingBasedOnTheRingID Get the REP ring based on the ring id. - 21b0-c89c-4aaa-9609
/* This API returns REP ring for the given id (The id of configured REP ring can be retrieved using the API /dna/intent/api/v1/iot/repRings/query).


@param id id path parameter. Ring ID of configured REP ring can be fetched using the API `/dna/intent/api/v1/iot/repRings/query`.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-r-e-p-ring-based-on-the-ring-id
*/
func (s *IndustrialConfigurationService) GetTheREPRingBasedOnTheRingID(id string) (*ResponseIndustrialConfigurationGetTheREPRingBasedOnTheRingID, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/repRings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIndustrialConfigurationGetTheREPRingBasedOnTheRingID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheREPRingBasedOnTheRingID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTheREPRingBasedOnTheRingId")
	}

	result := response.Result().(*ResponseIndustrialConfigurationGetTheREPRingBasedOnTheRingID)
	return result, response, err

}

//ConfigureAREPRingOnFABRICDeployment Configure a REP Ring on FABRIC deployment. - 8c85-28ca-492a-8a27
/* This API configures a REP ring on FABRIC deployment. The input payload contains the following fields ringName (unique ring name) , rootNetworkDeviceId (Network device ID of the root node of the REP Ring) and rootNeighbourNetworkDeviceIds (Network device IDs of the two immediate neighbour devices of the root node of the REP Ring). The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-a-r-e-p-ring-on-f-a-b-r-i-c-deployment
*/
func (s *IndustrialConfigurationService) ConfigureAREPRingOnFABRICDeployment(requestIndustrialConfigurationConfigureAREPRingOnFABRICDeployment *RequestIndustrialConfigurationConfigureAREPRingOnFABRICDeployment) (*ResponseIndustrialConfigurationConfigureAREPRingOnFABRICDeployment, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/fabric/repRings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIndustrialConfigurationConfigureAREPRingOnFABRICDeployment).
		SetResult(&ResponseIndustrialConfigurationConfigureAREPRingOnFABRICDeployment{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAREPRingOnFABRICDeployment(requestIndustrialConfigurationConfigureAREPRingOnFABRICDeployment)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAREPRingOnFABRICDeployment")
	}

	result := response.Result().(*ResponseIndustrialConfigurationConfigureAREPRingOnFABRICDeployment)
	return result, response, err

}

//ConfigureAREPRingOnNONFABRICDeployment Configure a REP Ring on NON-FABRIC deployment. - 8c85-28ca-492a-8a28
/* This API configures a REP ring on NON-FABRIC deployment. The input payload contains the following fields ringName (unique ring name) , rootNetworkDeviceId (Network device ID of the root node of the REP Ring) and rootNeighbourNetworkDeviceIds (Network device IDs of the two immediate neighbour devices of the root node of the REP Ring). The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!configure-a-r-e-p-ring-on-n-o-n-f-a-b-r-i-c-deployment
*/
func (s *IndustrialConfigurationService) ConfigureAREPRingOnNONFABRICDeployment(requestIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment *RequestIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment) (*ResponseIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/nonFabric/repRings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment).
		SetResult(&ResponseIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ConfigureAREPRingOnNONFABRICDeployment(requestIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment)
		}

		return nil, response, fmt.Errorf("error with operation ConfigureAREPRingOnNONFABRICDeployment")
	}

	result := response.Result().(*ResponseIndustrialConfigurationConfigureAREPRingOnNONFABRICDeployment)
	return result, response, err

}

//RetrievesTheListOfREPRings Retrieves the list of REP rings. - 69b7-4bb5-4e4b-b9aa
/* This API returns the list of REP rings for the given fields networkDeviceId (Network device ID of the REP ring member. In case of successful REP ring creation, any of the REP ring member networkDeviceId can be provided. In case of failed REP ring creation, provide only root node networkDeviceId. The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevice) and deploymentMode (FABRIC/NON_FABRIC).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-r-e-p-rings
*/
func (s *IndustrialConfigurationService) RetrievesTheListOfREPRings(requestIndustrialConfigurationRetrievesTheListOfREPRings *RequestIndustrialConfigurationRetrievesTheListOfREPRings) (*ResponseIndustrialConfigurationRetrievesTheListOfREPRings, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/repRings/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIndustrialConfigurationRetrievesTheListOfREPRings).
		SetResult(&ResponseIndustrialConfigurationRetrievesTheListOfREPRings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfREPRings(requestIndustrialConfigurationRetrievesTheListOfREPRings)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfREPRings")
	}

	result := response.Result().(*ResponseIndustrialConfigurationRetrievesTheListOfREPRings)
	return result, response, err

}

//RetrievesTheCountOfREPRings Retrieves the count of REP rings. - bca3-aaac-4278-942f
/* This API returns the count of REP rings for the given fields networkDeviceId (Network device ID of the REP ring member. The networkDeviceId is the instanceUuid attribute in the response of API /dna/intent/api/v1/networkDevices) and deploymentMode (FABRIC/NON_FABRIC).



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-r-e-p-rings
*/
func (s *IndustrialConfigurationService) RetrievesTheCountOfREPRings(requestIndustrialConfigurationRetrievesTheCountOfREPRings *RequestIndustrialConfigurationRetrievesTheCountOfREPRings) (*ResponseIndustrialConfigurationRetrievesTheCountOfREPRings, *resty.Response, error) {
	path := "/dna/intent/api/v1/iot/repRings/query/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIndustrialConfigurationRetrievesTheCountOfREPRings).
		SetResult(&ResponseIndustrialConfigurationRetrievesTheCountOfREPRings{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfREPRings(requestIndustrialConfigurationRetrievesTheCountOfREPRings)
		}

		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfREPRings")
	}

	result := response.Result().(*ResponseIndustrialConfigurationRetrievesTheCountOfREPRings)
	return result, response, err

}

//DeleteREPRingConfiguredInTheFABRICDeployment Delete REP ring configured in the FABRIC deployment. - bdb5-29f0-42a8-b33d
/* This API deletes the REP ring configured in the FABRIC deployment for the given id. The id of configured REP ring can be retrieved using the API /dna/intent/api/v1/iot/repRings/query.The taskid returned can be used to monitor the status of delete operation using following API -/intent/api/v1/task/{taskId}.


@param id id path parameter. Ring ID of configured REP ring can be fetched using the API `/dna/intent/api/v1/iot/repRings/query`.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-r-e-p-ring-configured-in-the-f-a-b-r-i-c-deployment
*/
func (s *IndustrialConfigurationService) DeleteREPRingConfiguredInTheFABRICDeployment(id string) (*ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheFABRICDeployment, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/iot/fabric/repRings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheFABRICDeployment{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteREPRingConfiguredInTheFABRICDeployment(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteREPRingConfiguredInTheFABRICDeployment")
	}

	result := response.Result().(*ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheFABRICDeployment)
	return result, response, err

}

//DeleteREPRingConfiguredInTheNONFABRICDeployment Delete REP ring configured in the NON-FABRIC deployment. - bdb5-29f0-42a8-b33c
/* This API deletes the REP ring configured in the NON-FABRIC deployment for the given id. The id of configured REP ring can be retrieved using the API /dna/intent/api/v1/iot/repRings/query.The taskid returned can be used to monitor the status of delete operation using following API -/intent/api/v1/task/{taskId}.


@param id id path parameter. Ring ID of configured REP ring can be fetched using the API `/dna/intent/api/v1/iot/repRings/query`.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-r-e-p-ring-configured-in-the-n-o-n-f-a-b-r-i-c-deployment
*/
func (s *IndustrialConfigurationService) DeleteREPRingConfiguredInTheNONFABRICDeployment(id string) (*ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheNONFABRICDeployment, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/iot/nonFabric/repRings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheNONFABRICDeployment{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteREPRingConfiguredInTheNONFABRICDeployment(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteREPRingConfiguredInTheNONFABRICDeployment")
	}

	result := response.Result().(*ResponseIndustrialConfigurationDeleteREPRingConfiguredInTheNONFABRICDeployment)
	return result, response, err

}
