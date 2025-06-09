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
	Limit  float64 `url:"limit,omitempty"`  //Return only this many IP Pool to SSID Mapping. Default is 500 if not specified. Maximum allowed limit is 500.
	Offset float64 `url:"offset,omitempty"` //Number of records to skip for pagination
}
type RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams struct {
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
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
type ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement struct {
	Response *ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version of the response.
}
type ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementResponse struct {
	TaskID string `json:"taskId,omitempty"` // Unique identifier for the task.
	URL    string `json:"url,omitempty"`    // URL for the task.
}
type ResponseFabricWirelessGetSdaWirelessDetailsFromSwitches struct {
	Response *[]ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponse `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponse struct {
	ID               string                                                                           `json:"id,omitempty"`               // Network Device ID of the Wireless Capable Switch
	EnableWireless   *bool                                                                            `json:"enableWireless,omitempty"`   // Enable Wireless
	RollingApUpgrade *ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponseRollingApUpgrade `json:"rollingApUpgrade,omitempty"` //
}
type ResponseFabricWirelessGetSdaWirelessDetailsFromSwitchesResponseRollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // Enable Rolling Ap Upgrade
	ApRebootPercentage     *int  `json:"apRebootPercentage,omitempty"`     // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type ResponseFabricWirelessReloadSwitchForWirelessControllerCleanup struct {
	Response *ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
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
type ResponseFabricWirelessUpdateSdaWirelessMulticast struct {
	Response *ResponseFabricWirelessUpdateSdaWirelessMulticastResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessUpdateSdaWirelessMulticastResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task ID
	URL    string `json:"url,omitempty"`    // Task URL
}
type ResponseFabricWirelessGetSdaWirelessMulticast struct {
	Response *ResponseFabricWirelessGetSdaWirelessMulticastResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseFabricWirelessGetSdaWirelessMulticastResponse struct {
	MulticastEnabled *bool `json:"multicastEnabled,omitempty"` // The setting indicates whether multicast is enabled (true) or disabled (false).
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
type RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement struct {
	ID               string                                                                                   `json:"id,omitempty"`               // Network Device ID of the wireless capable switch
	EnableWireless   *bool                                                                                    `json:"enableWireless,omitempty"`   // Enable Wireless
	RollingApUpgrade *RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementRollingApUpgrade `json:"rollingApUpgrade,omitempty"` //
}
type RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagementRollingApUpgrade struct {
	EnableRollingApUpgrade *bool `json:"enableRollingApUpgrade,omitempty"` // Enable Rolling Ap Upgrade
	ApRebootPercentage     *int  `json:"apRebootPercentage,omitempty"`     // AP Reboot Percentage. Permissible values - 5, 15, 25
}
type RequestFabricWirelessReloadSwitchForWirelessControllerCleanup struct {
	DeviceID string `json:"deviceId,omitempty"` // Network Device ID
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
type RequestFabricWirelessUpdateSdaWirelessMulticast struct {
	MulticastEnabled *bool `json:"multicastEnabled,omitempty"` // Multicast Enabled
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

//ReturnsAllTheFabricSitesThatHaveVLANToSSIDMapping Returns all the Fabric Sites that have VLAN to SSID mapping. - 7a96-98ce-400a-99ce
/* It will return all vlan to SSID mapping across all the fabric site


@param ReturnsAllTheFabricSitesThatHaveVLANToSSIDMappingQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-all-the-fabric-sites-that-have-vlan-to-ssid-mapping
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



Documentation Link: https://developer.cisco.com/docs/dna-center/#!return-the-count-of-all-the-fabric-site-which-has-ssid-to-ip-pool-mapping
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

//GetSdaWirelessDetailsFromSwitches Get SDA Wireless details from Switches - e48c-0bd1-459b-ad8d
/* Get the SDA Wireless details from the switches on the fabric site that have wireless capability enabled. A maximum of two switches can have a wireless role in a fabric site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sda-wireless-details-from-switches
*/
func (s *FabricWirelessService) GetSdaWirelessDetailsFromSwitches(fabricID string) (*ResponseFabricWirelessGetSdaWirelessDetailsFromSwitches, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/switchWirelessSetting"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessGetSdaWirelessDetailsFromSwitches{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSdaWirelessDetailsFromSwitches(fabricID)
		}
		return nil, response, fmt.Errorf("error with operation GetSdaWirelessDetailsFromSwitches")
	}

	result := response.Result().(*ResponseFabricWirelessGetSdaWirelessDetailsFromSwitches)
	return result, response, err

}

//RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSite Retrieve the VLANs and SSIDs mapped to the VLAN within a Fabric Site. - edbe-baa6-46cb-9f5e
/* Retrieve the VLANs and SSIDs mapped to the VLAN, within a Fabric Site. The 'fabricId' represents the Fabric ID of a particular Fabric Site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site

@param RetrieveTheVLANsAndSSIDsMappedToTheVLANWithinAFabricSiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-vlans-and-ssids-mapped-to-the-vlan-within-a-fabric-site
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-the-count-of-vlans-mapped-to-ssids-in-a-fabric-site
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

//GetSdaWirelessMulticast Get SDA Wireless Multicast - a58d-4bd1-4f89-8424
/* Retrieves the current Software-Defined Access (SDA) Wireless Multicast setting for a specified fabric site. The setting indicates whether multicast is enabled (true) or disabled (false). For optimal performance, ensure wired multicast is also enabled.


@param fabricID fabricId path parameter. The unique identifier of the fabric site for which the multicast setting is being requested. The identifier should be in the format of a UUID. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-sda-wireless-multicast
*/
func (s *FabricWirelessService) GetSdaWirelessMulticast(fabricID string) (*ResponseFabricWirelessGetSdaWirelessMulticast, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/wirelessMulticast"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFabricWirelessGetSdaWirelessMulticast{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSdaWirelessMulticast(fabricID)
		}
		return nil, response, fmt.Errorf("error with operation GetSdaWirelessMulticast")
	}

	result := response.Result().(*ResponseFabricWirelessGetSdaWirelessMulticast)
	return result, response, err

}

//AddSSIDToIPPoolMapping Add SSID to IP Pool Mapping - b783-49d9-463a-98dd
/* Add SSID to IP Pool Mapping



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-ssid-to-ip-pool-mapping
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

//ReloadSwitchForWirelessControllerCleanup Reload Switch for Wireless Controller Cleanup - 7ab4-3994-429b-827c
/* This API is used to reload switches after disabling wireless to remove the wireless-controller configuration on the device. When wireless is disabled on a switch, all wireless configurations are removed except for the wireless-controller configuration. To completely remove the wireless-controller configuration, you can use this API. Please note that using this API will cause a reload of the device(s). This API should only be used for devices that have wireless disabled but still have the 'wireless-controller' configuration present. The reload payload can have a maximum of two switches as only two switches can have a wireless role in a fabric site.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851


Documentation Link: https://developer.cisco.com/docs/dna-center/#!reload-switch-for-wireless-controller-cleanup
*/
func (s *FabricWirelessService) ReloadSwitchForWirelessControllerCleanup(fabricID string, requestFabricWirelessReloadSwitchForWirelessControllerCleanup *RequestFabricWirelessReloadSwitchForWirelessControllerCleanup) (*ResponseFabricWirelessReloadSwitchForWirelessControllerCleanup, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/switchWirelessSetting/reload"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessReloadSwitchForWirelessControllerCleanup).
		SetResult(&ResponseFabricWirelessReloadSwitchForWirelessControllerCleanup{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReloadSwitchForWirelessControllerCleanup(fabricID, requestFabricWirelessReloadSwitchForWirelessControllerCleanup)
		}

		return nil, response, fmt.Errorf("error with operation ReloadSwitchForWirelessControllerCleanup")
	}

	result := response.Result().(*ResponseFabricWirelessReloadSwitchForWirelessControllerCleanup)
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

//SwitchWirelessSettingAndRollingApUpgradeManagement Switch Wireless Setting and Rolling AP Upgrade Management - 2d95-183a-46db-87ed
/* This API is used to enable or disable wireless capabilities on switch devices, along with configuring rolling AP upgrades on the fabric site. Reboot action is required to remove wireless configurations.


@param fabricID fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851

*/
func (s *FabricWirelessService) SwitchWirelessSettingAndRollingApUpgradeManagement(fabricID string, requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagement *RequestFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement) (*ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/switchWirelessSetting"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagement).
		SetResult(&ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.SwitchWirelessSettingAndRollingApUpgradeManagement(fabricID, requestFabricWirelessSwitchWirelessSettingAndRollingAPUpgradeManagement)
		}
		return nil, response, fmt.Errorf("error with operation SwitchWirelessSettingAndRollingApUpgradeManagement")
	}

	result := response.Result().(*ResponseFabricWirelessSwitchWirelessSettingAndRollingApUpgradeManagement)
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

//UpdateSdaWirelessMulticast Update SDA Wireless Multicast - 908b-d97a-4458-8352
/* Updates the Software-Defined Access (SDA) Wireless Multicast setting for a specified fabric site. This API allows you to enable or disable the multicast feature. For optimal performance, ensure wired multicast is also enabled.


@param fabricID fabricId path parameter. The unique identifier of the fabric site for which the multicast setting is being requested. The identifier should be in the format of a UUID. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.

*/
func (s *FabricWirelessService) UpdateSdaWirelessMulticast(fabricID string, requestFabricWirelessUpdateSDAWirelessMulticast *RequestFabricWirelessUpdateSdaWirelessMulticast) (*ResponseFabricWirelessUpdateSdaWirelessMulticast, *resty.Response, error) {
	path := "/dna/intent/api/v1/sda/fabrics/{fabricId}/wirelessMulticast"
	path = strings.Replace(path, "{fabricId}", fmt.Sprintf("%v", fabricID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestFabricWirelessUpdateSDAWirelessMulticast).
		SetResult(&ResponseFabricWirelessUpdateSdaWirelessMulticast{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateSdaWirelessMulticast(fabricID, requestFabricWirelessUpdateSDAWirelessMulticast)
		}
		return nil, response, fmt.Errorf("error with operation UpdateSdaWirelessMulticast")
	}

	result := response.Result().(*ResponseFabricWirelessUpdateSdaWirelessMulticast)
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
			return s.RemoveWLCFromFabricDomain(
				RemoveWLCFromFabricDomainHeaderParams, RemoveWLCFromFabricDomainQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RemoveWLCFromFabricDomain")
	}

	result := response.Result().(*ResponseFabricWirelessRemoveWLCFromFabricDomain)
	return result, response, err

}
