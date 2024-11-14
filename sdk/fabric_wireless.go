package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type FabricWirelessService service

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
type ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //Return only this many IP Pool to SSID Mapping
	Offset float64 `url:"offset,omitempty"` //Number of records to skip for pagination
}
type RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}

type ResponseFabricWirelessAddSSIDToIPPoolMapping struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
}
type ResponseFabricWirelessUpdateSSIDToIPPoolMapping struct {
	ExecutionID        string `json:"executionId,omitempty"`        // Execution Id
	ExecutionStatusURL string `json:"executionStatusUrl,omitempty"` // Execution Status Url
	Message            string `json:"message,omitempty"`            // Message
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
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping struct {
	Response *[]ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponse struct {
	FabricID    string                                                                                        `json:"fabricId,omitempty"`    // Fabric Id
	VLANDetails *[]ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponseVLANDetails `json:"vlanDetails,omitempty"` //
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponseVLANDetails struct {
	VLANName    string                                                                                                   `json:"vlanName,omitempty"`    // Vlan Name
	SSIDDetails *[]ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponseVLANDetailsSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingResponseVLANDetailsSSIDDetails struct {
	Name             string `json:"name,omitempty"`             // Name of the SSID.
	SecurityGroupTag string `json:"securityGroupTag,omitempty"` // Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
}
type ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping struct {
	Response *ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingResponse `json:"response,omitempty"` //
	Version  string                                                                                     `json:"version,omitempty"`  // Response Version
}
type ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMappingResponse struct {
	Count *int `json:"count,omitempty"` // Return the count of all the fabric site which has SSID to IP Pool mapping
}
type ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN struct {
	Response *ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANResponse `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite struct {
	Response *[]ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponse `json:"response,omitempty"` //
	Version  string                                                                                    `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponse struct {
	VLANName    string                                                                                               `json:"vlanName,omitempty"`    // Vlan Name
	SSIDDetails *[]ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponseSSIDDetails `json:"ssidDetails,omitempty"` //
}
type ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteResponseSSIDDetails struct {
	Name             string `json:"name,omitempty"`             // Name of the SSID
	SecurityGroupTag string `json:"securityGroupTag,omitempty"` // Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
}
type ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite struct {
	Response *ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteResponse `json:"response,omitempty"` //
	Version  string                                                                          `json:"version,omitempty"`  // Response Version
}
type ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSiteResponse struct {
	Count *int `json:"count,omitempty"` // Returns the count of VLANs mapped to SSIDs in a Fabric Site
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
	DeviceName        string `json:"deviceName,omitempty"`        // WLC Device Name
	SiteNameHierarchy string `json:"siteNameHierarchy,omitempty"` // Fabric Site Name Hierarchy
}
type RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN []RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN // Array of RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN
type RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN struct {
	VLANName    string                                                                     `json:"vlanName,omitempty"`    // Vlan Name
	SSIDDetails *[]RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANSSIDDetails `json:"ssidDetails,omitempty"` //
}
type RequestItemFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLANSSIDDetails struct {
	Name             string `json:"name,omitempty"`             // Name of the SSID
	SecurityGroupTag string `json:"securityGroupTag,omitempty"` // Represents the name of the Security Group. Example: Auditors, BYOD, Developers, etc.
}

//GetSSIDToIPPoolMapping Get SSID to IP Pool Mapping - d891-8a44-4b6a-ad19
/* Get SSID to IP Pool Mapping


@param GetSSIDToIPPoolMappingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-ssid-to-ip-pool-mapping-v1
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

//ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping Returns all the Fabric Sites that have VLAN to SSID mapping. - 7a96-98ce-400a-99ce
/* It will return all vlan to SSID mapping across all the fabric site


@param ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-all-the-fabric-sites-that-have-vlan-to-ssid-mapping-v1
*/
func (s *FabricWirelessService) ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams *ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams) (*ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/vlanToSsids"

	queryString, _ := query.Values(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping(ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsAllTheFabricSitesThatHaveVlanToSsidMapping")
	}

	result := response.Result().(*ResponseFabricWirelessReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping)
	return result, response, err

}

//ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping Return the count of all the fabric site which has SSID to IP Pool mapping  - 36b0-0b14-44fa-8c4b
/* Return the count of all the fabric site which has SSID to IP Pool mapping



Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-the-count-of-all-the-fabric-site-which-has-ssid-to-ip-pool-mapping-v1
*/
func (s *FabricWirelessService) ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping() (*ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/vlanToSsids/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping()
		}
		return nil, response, fmt.Errorf("error with operation ReturnTheCountOfAllTheFabricSiteWhichHasSsidToIpPoolMapping")
	}

	result := response.Result().(*ResponseFabricWirelessReturnTheCountOfAllTheFabricSiteWhichHasSSIDToIPPoolMapping)
	return result, response, err

}

//RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite Retrieve the VLANs and SSIDs mapped to the VLAN within a Fabric Site. - edbe-baa6-46cb-9f5e
/* Retrieve the VLANs and SSIDs mapped to the VLAN, within a Fabric Site. The 'fabricId' represents the Fabric ID of a particular Fabric Site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

@param RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-vlans-and-ssids-mapped-to-the-vlan-within-a-fabric-site-v1
*/
func (s *FabricWirelessService) RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(fabricID string, RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams *RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams) (*ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/vlanToSsids"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	queryString, _ := query.Values(RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite(fabricID, RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheVlansAndSsidsMappedToTheVlanWithinAFabricSite")
	}

	result := response.Result().(*ResponseFabricWirelessRetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite)
	return result, response, err

}

//ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite Returns the count of VLANs mapped to SSIDs in a Fabric Site. - e0ab-88b3-4198-a152
/* Returns the count of VLANs mapped to SSIDs in a Fabric Site. The 'fabricId' represents the Fabric ID of a particular Fabric Site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site


Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-the-count-of-vlans-mapped-to-ssids-in-a-fabric-site-v1
*/
func (s *FabricWirelessService) ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite(fabricID string) (*ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/vlanToSsids/count"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite(fabricID)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsTheCountOfVlansMappedToSsidsInAFabricSite")
	}

	result := response.Result().(*ResponseFabricWirelessReturnsTheCountOfVLANsMappedToSSIDsInAFabricSite)
	return result, response, err

}

//AddSSIDToIPPoolMapping Add SSID to IP Pool Mapping - b783-49d9-463a-98dd
/* Add SSID to IP Pool Mapping



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-ssid-to-ip-pool-mapping-v1
*/
func (s *FabricWirelessService) AddSSIDToIPPoolMapping(requestFabricWirelessAddSSIDToIPPoolMapping *RequestFabricWirelessAddSSIDToIPPoolMapping) (*ResponseFabricWirelessAddSSIDToIPPoolMapping, *resty.Response, error) {
	path := "/dna/intent/api/v1/business/sda/hostonboarding/ssid-ippool"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessAddSSIDToIPPoolMapping).
		SetResult(&ResponseFabricWirelessAddSSIDToIPPoolMapping{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddSSIDToIPPoolMapping(requestFabricWirelessAddSSIDToIPPoolMapping)
		}

		return nil, response, fmt.Errorf("error with operation AddSsidToIpPoolMapping")
	}

	result := response.Result().(*ResponseFabricWirelessAddSSIDToIPPoolMapping)
	return result, response, err

}

//AddWLCToFabricDomain Add WLC to Fabric Domain - f4ad-b85b-4f19-ae86
/* Add WLC to Fabric Domain



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-w-l-c-to-fabric-domain-v1
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
/* Update SSID to IP Pool Mapping


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

//AddUpdateOrRemoveSSIDMappingToAVLAN Add, Update or Remove SSID mapping to a VLAN - 07af-b879-4c2a-983b
/* Add, update, or remove SSID mappings to a VLAN. If the payload doesn't contain a 'vlanName' which has SSIDs mapping done earlier then all the mapped SSIDs of the 'vlanName' is cleared. The request must include all SSIDs currently mapped to a VLAN, as determined by the response from the GET operation for the same fabricId used in the request. If an already-mapped SSID is not included in the payload, its mapping will be removed by this API. Conversely, if a new SSID is provided, it will be added to the Mapping. Ensure that any new SSID added is a Fabric SSID. This API can also be used to add a VLAN and associate the relevant SSIDs with it. The 'vlanName' must be 'Fabric Wireless Enabled' and should be part of the Fabric Site representing 'Fabric ID' specified in the API request.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

*/
func (s *FabricWirelessService) AddUpdateOrRemoveSSIDMappingToAVLAN(fabricID string, requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN *RequestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN) (*ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/vlanToSsids"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN).
		SetResult(&ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddUpdateOrRemoveSSIDMappingToAVLAN(fabricID, requestFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN)
		}
		return nil, response, fmt.Errorf("error with operation AddUpdateOrRemoveSsidMappingToAVlan")
	}

	result := response.Result().(*ResponseFabricWirelessAddUpdateOrRemoveSSIDMappingToAVLAN)
	return result, response, err

}

//RemoveWLCFromFabricDomain Remove WLC from Fabric Domain - 10bb-1ae9-46e9-840f
/* Remove WLC from Fabric Domain


@param RemoveWLCFromFabricDomainHeaderParams Custom header parameters
@param RemoveWLCFromFabricDomainQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-w-l-c-from-fabric-domain-v1
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
