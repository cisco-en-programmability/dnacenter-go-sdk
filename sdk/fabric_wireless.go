package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type FabricWirelessService service

type AddSSIDToIPPoolMappingHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}
type GetSSIDToIPPoolMappingQueryParams struct {
	VLANName          string `url:"vlanName,omitempty"`          //VLAN Name
	SiteNameHierarchy string `url:"siteNameHierarchy,omitempty"` //Site Name Heirarchy
}
type RemoveWLCFromFabricDomainQueryParams struct {
	DeviceIPAddress string `url:"deviceIPAddress,omitempty"` //Device Management IP Address
}
type RemoveWLCFromFabricDomainHeaderParams struct {
	Persistbapioutput string `url:"__persistbapioutput,omitempty"` //Expects type string. Enable this parameter to execute the API and return a response asynchronously.
}

type ResponseFabricWirelessAddSSIDToIPPoolMapping struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Status of the job for wireless state change in fabric domain
	ExecutionStatusURL string `json:"executionStatusURL,omitempty"` // executionStatusURL
	Message            string `json:"message,omitempty"`            // message
}
type ResponseFabricWirelessUpdateSSIDToIPPoolMapping struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Status of the job for wireless state change in fabric domain
	ExecutionStatusURL string `json:"executionStatusURL,omitempty"` // executionStatusURL
	Message            string `json:"message,omitempty"`            // message
}
type ResponseFabricWirelessGetSSIDToIPPoolMapping struct {
	VLANName    string                                                     `json:"vlanName,omitempty"`    // VLAN Name
	SSIDDetails *[]ResponseFabricWirelessGetSSIDToIPPoolMappingSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseFabricWirelessGetSSIDToIPPoolMappingSSIDDetails struct {
	Name              string `json:"name,omitempty"`              // SSID Name
	ScalableGroupName string `json:"scalableGroupName,omitempty"` // Scalable Group Name
}
type ResponseFabricWirelessRemoveWLCFromFabricDomain struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseFabricWirelessAddWLCToFabricDomain struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type RequestFabricWirelessAddSSIDToIPPoolMapping struct {
	VLANName          string   `json:"vlanName,omitempty"`          // VLAN Name
	ScalableGroupName string   `json:"scalableGroupName,omitempty"` // Scalable Group Name
	SSIDNames         []string `json:"ssidNames,omitempty"`         // List of SSIDs
	SiteNameHierarchy string   `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type RequestFabricWirelessUpdateSSIDToIPPoolMapping struct {
	VLANName          string   `json:"vlanName,omitempty"`          // VLAN Name
	ScalableGroupName string   `json:"scalableGroupName,omitempty"` // Scalable Group Name
	SSIDNames         []string `json:"ssidNames,omitempty"`         // List of SSIDs
	SiteNameHierarchy string   `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}
type RequestFabricWirelessAddWLCToFabricDomain struct {
	DeviceName        string `json:"deviceName,omitempty"`        // EWLC Device Name
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Site Name Hierarchy
}

//GetSSIDToIPPoolMapping Get SSID to IP Pool Mapping - d891-8a44-4b6a-ad19
/* Get SSID to IP Pool Mapping


@param GetSSIDToIPPoolMappingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-to-ip-pool-mapping
*/
func (s *FabricWirelessService) GetSSIDToIPPoolMapping(GetSSIDToIPPoolMappingQueryParams *GetSSIDToIPPoolMappingQueryParams) (*ResponseFabricWirelessGetSSIDToIPPoolMapping, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	queryString, _ := query.Values(GetSSIDToIPPoolMappingQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessGetSSIDToIPPoolMapping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSSIDToIPPoolMapping(GetSSIDToIPPoolMappingQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSsidToIpPoolMapping")
	}

	result := response.Result().(*ResponseFabricWirelessGetSSIDToIPPoolMapping)
	return result, response, err

}

//AddSSIDToIPPoolMapping Add SSID to IP Pool Mapping - b783-49d9-463a-98dd
/* Add SSID to IP Pool Mapping.


@param AddSSIDToIPPoolMappingHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-ssid-to-ip-pool-mapping
*/
func (s *FabricWirelessService) AddSSIDToIPPoolMapping(requestFabricWirelessAddSSIDToIPPoolMapping *RequestFabricWirelessAddSSIDToIPPoolMapping, AddSSIDToIPPoolMappingHeaderParams *AddSSIDToIPPoolMappingHeaderParams) (*ResponseFabricWirelessAddSSIDToIPPoolMapping, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if AddSSIDToIPPoolMappingHeaderParams != nil {

		if AddSSIDToIPPoolMappingHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", AddSSIDToIPPoolMappingHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetBody(requestFabricWirelessAddSSIDToIPPoolMapping).
		SetResult(&ResponseFabricWirelessAddSSIDToIPPoolMapping{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddSSIDToIPPoolMapping(requestFabricWirelessAddSSIDToIPPoolMapping, AddSSIDToIPPoolMappingHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation AddSsidToIpPoolMapping")
	}

	result := response.Result().(*ResponseFabricWirelessAddSSIDToIPPoolMapping)
	return result, response, err

}

//AddWLCToFabricDomain Add WLC to Fabric Domain - f4ad-b85b-4f19-ae86
/* Add WLC to Fabric Domain



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-w-l-c-to-fabric-domain
*/
func (s *FabricWirelessService) AddWLCToFabricDomain(requestFabricWirelessAddWLCToFabricDomain *RequestFabricWirelessAddWLCToFabricDomain) (*ResponseFabricWirelessAddWLCToFabricDomain, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/wireless-controller"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessAddWLCToFabricDomain).
		SetResult(&ResponseFabricWirelessAddWLCToFabricDomain{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddWLCToFabricDomain(requestFabricWirelessAddWLCToFabricDomain)
		}

		return nil, response, fmt.Errorf("error with operation AddWLCToFabricDomain")
	}

	result := response.Result().(*ResponseFabricWirelessAddWLCToFabricDomain)
	return result, response, err

}

//UpdateSSIDToIPPoolMapping Update SSID to IP Pool Mapping - a7bf-1936-424b-91f0
/* Update SSID to IP Pool Mapping.


 */
func (s *FabricWirelessService) UpdateSSIDToIPPoolMapping(requestFabricWirelessUpdateSSIDToIPPoolMapping *RequestFabricWirelessUpdateSSIDToIPPoolMapping) (*ResponseFabricWirelessUpdateSSIDToIPPoolMapping, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessUpdateSSIDToIPPoolMapping).
		SetResult(&ResponseFabricWirelessUpdateSSIDToIPPoolMapping{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSSIDToIPPoolMapping(requestFabricWirelessUpdateSSIDToIPPoolMapping)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSsidToIpPoolMapping")
	}

	result := response.Result().(*ResponseFabricWirelessUpdateSSIDToIPPoolMapping)
	return result, response, err

}

//RemoveWLCFromFabricDomain Remove WLC from Fabric Domain - 10bb-1ae9-46e9-840f
/* Remove WLC from Fabric Domain


@param RemoveWLCFromFabricDomainHeaderParams Custom header parameters
@param RemoveWLCFromFabricDomainQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-w-l-c-from-fabric-domain
*/
func (s *FabricWirelessService) RemoveWLCFromFabricDomain(RemoveWLCFromFabricDomainHeaderParams *RemoveWLCFromFabricDomainHeaderParams, RemoveWLCFromFabricDomainQueryParams *RemoveWLCFromFabricDomainQueryParams) (*ResponseFabricWirelessRemoveWLCFromFabricDomain, *resty.Response, error) {
	//RemoveWLCFromFabricDomainHeaderParams *RemoveWLCFromFabricDomainHeaderParams,RemoveWLCFromFabricDomainQueryParams *RemoveWLCFromFabricDomainQueryParams
	path := "/dna/intent/api/v1/business/sda/wireless-controller"

	queryString, _ := query.Values(RemoveWLCFromFabricDomainQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if RemoveWLCFromFabricDomainHeaderParams != nil {

		if RemoveWLCFromFabricDomainHeaderParams.Persistbapioutput != "" {
			clientRequest = clientRequest.SetHeader("__persistbapioutput", RemoveWLCFromFabricDomainHeaderParams.Persistbapioutput)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessRemoveWLCFromFabricDomain{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveWLCFromFabricDomain(RemoveWLCFromFabricDomainHeaderParams, RemoveWLCFromFabricDomainQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RemoveWLCFromFabricDomain")
	}

	result := response.Result().(*ResponseFabricWirelessRemoveWLCFromFabricDomain)
	return result, response, err

}
