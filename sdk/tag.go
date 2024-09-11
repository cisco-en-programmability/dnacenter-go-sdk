package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TagService service

type GetTagQueryParams struct {
	Name                     string  `url:"name,omitempty"`                      //Tag name is mandatory when filter operation is used.
	AdditionalInfonameSpace  string  `url:"additionalInfo.nameSpace,omitempty"`  //nameSpace
	AdditionalInfoattributes string  `url:"additionalInfo.attributes,omitempty"` //attributeName
	Level                    string  `url:"level,omitempty"`                     //levelArg
	Offset                   float64 `url:"offset,omitempty"`                    //offset
	Limit                    float64 `url:"limit,omitempty"`                     //limit
	Size                     string  `url:"size,omitempty"`                      //size in kilobytes(KB)
	Field                    string  `url:"field,omitempty"`                     //Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
	SortBy                   string  `url:"sortBy,omitempty"`                    //Only supported attribute is name. SortyBy is mandatory when order is used.
	Order                    string  `url:"order,omitempty"`                     //Available values are asc and des
	SystemTag                string  `url:"systemTag,omitempty"`                 //systemTag
}
type GetTagCountQueryParams struct {
	Name          string `url:"name,omitempty"`          //tagName
	NameSpace     string `url:"nameSpace,omitempty"`     //nameSpace
	AttributeName string `url:"attributeName,omitempty"` //attributeName
	Size          string `url:"size,omitempty"`          //size in kilobytes(KB)
	SystemTag     string `url:"systemTag,omitempty"`     //systemTag
}
type GetTagMembersByIDQueryParams struct {
	MemberType            string  `url:"memberType,omitempty"`            //Entity type of the member. Possible values can be retrieved by using /tag/member/type API
	Offset                float64 `url:"offset,omitempty"`                //Used for pagination. It indicates the starting row number out of available member records
	Limit                 float64 `url:"limit,omitempty"`                 //Used to Number of maximum members to return in the result
	MemberAssociationType string  `url:"memberAssociationType,omitempty"` //Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
	Level                 string  `url:"level,omitempty"`                 //level
}
type GetTagMemberCountQueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            //memberType
	MemberAssociationType string `url:"memberAssociationType,omitempty"` //memberAssociationType
}
type RetrieveTagsAssociatedWithTheInterfacesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. minimum: 1
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. minimum: 1, maximum: 500
}
type RetrieveTagsAssociatedWithNetworkDevicesQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1. minimum: 1
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page. minimum: 1, maximum: 500
}

type ResponseTagUpdateTag struct {
	Version  string                        `json:"version,omitempty"`  //
	Response *ResponseTagUpdateTagResponse `json:"response,omitempty"` //
}
type ResponseTagUpdateTagResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTag struct {
	Version  string                       `json:"version,omitempty"`  //
	Response *[]ResponseTagGetTagResponse `json:"response,omitempty"` //
}
type ResponseTagGetTagResponse struct {
	SystemTag        *bool                                    `json:"systemTag,omitempty"`        //
	Description      string                                   `json:"description,omitempty"`      //
	DynamicRules     *[]ResponseTagGetTagResponseDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                                   `json:"name,omitempty"`             //
	ID               string                                   `json:"id,omitempty"`               //
	InstanceTenantID string                                   `json:"instanceTenantId,omitempty"` //
}
type ResponseTagGetTagResponseDynamicRules struct {
	MemberType string                                      `json:"memberType,omitempty"` //
	Rules      *ResponseTagGetTagResponseDynamicRulesRules `json:"rules,omitempty"`      //
}
type ResponseTagGetTagResponseDynamicRulesRules struct {
	Values    []string                                           `json:"values,omitempty"`    //
	Items     *[]ResponseTagGetTagResponseDynamicRulesRulesItems `json:"items,omitempty"`     //
	Operation string                                             `json:"operation,omitempty"` //
	Name      string                                             `json:"name,omitempty"`      //
	Value     string                                             `json:"value,omitempty"`     //
}
type ResponseTagGetTagResponseDynamicRulesRulesItems interface{}
type ResponseTagCreateTag struct {
	Version  string                        `json:"version,omitempty"`  //
	Response *ResponseTagCreateTagResponse `json:"response,omitempty"` //
}
type ResponseTagCreateTagResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagCount struct {
	Version  string `json:"version,omitempty"`  //
	Response *int   `json:"response,omitempty"` //
}
type ResponseTagUpdateTagMembership struct {
	Version  string                                  `json:"version,omitempty"`  //
	Response *ResponseTagUpdateTagMembershipResponse `json:"response,omitempty"` //
}
type ResponseTagUpdateTagMembershipResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagResourceTypes struct {
	Version  string   `json:"version,omitempty"`  //
	Response []string `json:"response,omitempty"` //
}
type ResponseTagDeleteTag struct {
	Version  string                        `json:"version,omitempty"`  //
	Response *ResponseTagDeleteTagResponse `json:"response,omitempty"` //
}
type ResponseTagDeleteTagResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagByID struct {
	Version  string                         `json:"version,omitempty"`  //
	Response *ResponseTagGetTagByIDResponse `json:"response,omitempty"` //
}
type ResponseTagGetTagByIDResponse struct {
	SystemTag        *bool                                        `json:"systemTag,omitempty"`        //
	Description      string                                       `json:"description,omitempty"`      //
	DynamicRules     *[]ResponseTagGetTagByIDResponseDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                                       `json:"name,omitempty"`             //
	ID               string                                       `json:"id,omitempty"`               //
	InstanceTenantID string                                       `json:"instanceTenantId,omitempty"` //
}
type ResponseTagGetTagByIDResponseDynamicRules struct {
	MemberType string                                          `json:"memberType,omitempty"` //
	Rules      *ResponseTagGetTagByIDResponseDynamicRulesRules `json:"rules,omitempty"`      //
}
type ResponseTagGetTagByIDResponseDynamicRulesRules struct {
	Values    []string                                               `json:"values,omitempty"`    //
	Items     *[]ResponseTagGetTagByIDResponseDynamicRulesRulesItems `json:"items,omitempty"`     //
	Operation string                                                 `json:"operation,omitempty"` //
	Name      string                                                 `json:"name,omitempty"`      //
	Value     string                                                 `json:"value,omitempty"`     //
}
type ResponseTagGetTagByIDResponseDynamicRulesRulesItems interface{}
type ResponseTagGetTagMembersByID struct {
	Version  string                                  `json:"version,omitempty"`  //
	Response *[]ResponseTagGetTagMembersByIDResponse `json:"response,omitempty"` //
}
type ResponseTagGetTagMembersByIDResponse struct {
	InstanceUUID string `json:"instanceUuid,omitempty"` //
}
type ResponseTagAddMembersToTheTag struct {
	Version  string                                 `json:"version,omitempty"`  //
	Response *ResponseTagAddMembersToTheTagResponse `json:"response,omitempty"` //
}
type ResponseTagAddMembersToTheTagResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagGetTagMemberCount struct {
	Version  string `json:"version,omitempty"`  //
	Response *int   `json:"response,omitempty"` //
}
type ResponseTagRemoveTagMember struct {
	Version  string                              `json:"version,omitempty"`  //
	Response *ResponseTagRemoveTagMemberResponse `json:"response,omitempty"` //
}
type ResponseTagRemoveTagMemberResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseTagRetrieveTagsAssociatedWithTheInterfaces struct {
	Response *[]ResponseTagRetrieveTagsAssociatedWithTheInterfacesResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTagsAssociatedWithTheInterfacesResponse struct {
	ID   string                                                            `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagRetrieveTagsAssociatedWithTheInterfacesResponseTags `json:"tags,omitempty"` //
}
type ResponseTagRetrieveTagsAssociatedWithTheInterfacesResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag struct {
	Response *ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTagQueryTheTagsAssociatedWithInterfaces struct {
	Response *[]ResponseTagQueryTheTagsAssociatedWithInterfacesResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagQueryTheTagsAssociatedWithInterfacesResponse struct {
	ID   string                                                         `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagQueryTheTagsAssociatedWithInterfacesResponseTags `json:"tags,omitempty"` //
}
type ResponseTagQueryTheTagsAssociatedWithInterfacesResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type ResponseTagRetrieveTagsAssociatedWithNetworkDevices struct {
	Response *[]ResponseTagRetrieveTagsAssociatedWithNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTagsAssociatedWithNetworkDevicesResponse struct {
	ID   string                                                             `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagRetrieveTagsAssociatedWithNetworkDevicesResponseTags `json:"tags,omitempty"` //
}
type ResponseTagRetrieveTagsAssociatedWithNetworkDevicesResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag struct {
	Response *ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagResponse `json:"response,omitempty"` //
	Version  string                                                                                 `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTagResponse struct {
	Count *int `json:"count,omitempty"` // The reported count.
}
type ResponseTagQueryTheTagsAssociatedWithNetworkDevices struct {
	Response *[]ResponseTagQueryTheTagsAssociatedWithNetworkDevicesResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // The version of the response.
}
type ResponseTagQueryTheTagsAssociatedWithNetworkDevicesResponse struct {
	ID   string                                                             `json:"id,omitempty"`   // Id of the member (network device or interface)
	Tags *[]ResponseTagQueryTheTagsAssociatedWithNetworkDevicesResponseTags `json:"tags,omitempty"` //
}
type ResponseTagQueryTheTagsAssociatedWithNetworkDevicesResponseTags struct {
	ID   string `json:"id,omitempty"`   // Tag id
	Name string `json:"name,omitempty"` // Tag name
}
type RequestTagUpdateTag struct {
	SystemTag        *bool                              `json:"systemTag,omitempty"`        // true for system created tags, false for user defined tags
	Description      string                             `json:"description,omitempty"`      // description of the tag.
	DynamicRules     *[]RequestTagUpdateTagDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                             `json:"name,omitempty"`             // name of the tag.
	ID               string                             `json:"id,omitempty"`               // mandatory instanceUuid of the tag that needs to be updated.
	InstanceTenantID string                             `json:"instanceTenantId,omitempty"` // instanceTenantId generated for the tag.
}
type RequestTagUpdateTagDynamicRules struct {
	MemberType string                                `json:"memberType,omitempty"` // memberType of the tag (e.g. networkdevice, interface)
	Rules      *RequestTagUpdateTagDynamicRulesRules `json:"rules,omitempty"`      //
}
type RequestTagUpdateTagDynamicRulesRules struct {
	Values    []string                                     `json:"values,omitempty"`    // values of the parameter,Only one of the value or values can be used for the given parameter. (for managementIpAddress e.g. ["10.197.124.90","10.197.124.91"])
	Items     *[]RequestTagUpdateTagDynamicRulesRulesItems `json:"items,omitempty"`     // items details,multiple rules can be defined by items(e.g. "items": [{"operation": "ILIKE", "name": "managementIpAddress", "value": "%10%"}, {"operation": "ILIKE", "name": "hostname", "value": "%NA%"} ])
	Operation string                                       `json:"operation,omitempty"` // opeartion used in the rules (e.g. OR,IN,EQ,LIKE,ILIKE,AND)
	Name      string                                       `json:"name,omitempty"`      // name of the parameter (e.g. for interface:portName,adminStatus,speed,status,description. for networkdevice:family,series,hostname,managementIpAddress,groupNameHierarchy,softwareVersion)
	Value     string                                       `json:"value,omitempty"`     // value of the parameter (e.g. for portName:1/0/1,for adminStatus,status:up/down, for speed: any integer value, for description: any valid string, for family:switches, for series:C3650, for managementIpAddress:10.197.124.90, groupNameHierarchy:Global, softwareVersion: 16.9.1)
}
type RequestTagUpdateTagDynamicRulesRulesItems interface{}
type RequestTagCreateTag struct {
	SystemTag        *bool                              `json:"systemTag,omitempty"`        // true for system created tags, false for user defined tags
	Description      string                             `json:"description,omitempty"`      // description of the tag.
	DynamicRules     *[]RequestTagCreateTagDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                             `json:"name,omitempty"`             // name of the tag.
	ID               string                             `json:"id,omitempty"`               // instanceUuid generated for the tag.
	InstanceTenantID string                             `json:"instanceTenantId,omitempty"` // instanceTenantId generated for the tag.
}
type RequestTagCreateTagDynamicRules struct {
	MemberType string                                `json:"memberType,omitempty"` // memberType of the tag (e.g. networkdevice, interface)
	Rules      *RequestTagCreateTagDynamicRulesRules `json:"rules,omitempty"`      //
}
type RequestTagCreateTagDynamicRulesRules struct {
	Values    []string                                     `json:"values,omitempty"`    // values of the parameter,Only one of the value or values can be used for the given parameter. (for managementIpAddress e.g. ["10.197.124.90","10.197.124.91"])
	Items     *[]RequestTagCreateTagDynamicRulesRulesItems `json:"items,omitempty"`     // items details,multiple rules can be defined by items(e.g. "items": [{"operation": "ILIKE", "name": "managementIpAddress", "value": "%10%"}, {"operation": "ILIKE", "name": "hostname", "value": "%NA%"} ])
	Operation string                                       `json:"operation,omitempty"` // opeartion used in the rules (e.g. OR,IN,EQ,LIKE,ILIKE,AND)
	Name      string                                       `json:"name,omitempty"`      // name of the parameter (e.g. for interface:portName,adminStatus,speed,status,description. for networkdevice:family,series,hostname,managementIpAddress,groupNameHierarchy,softwareVersion)
	Value     string                                       `json:"value,omitempty"`     // value of the parameter (e.g. for portName:1/0/1,for adminStatus,status:up/down, for speed: any integer value, for description: any valid string, for family:switches, for series:C3650, for managementIpAddress:10.197.124.90, groupNameHierarchy:Global, softwareVersion: 16.9.1)
}
type RequestTagCreateTagDynamicRulesRulesItems interface{}
type RequestTagUpdateTagMembership struct {
	MemberToTags map[string][]string `json:"memberToTags,omitempty"` //
	MemberType   string              `json:"memberType,omitempty"`   //
}
type RequestTagUpdateTagMembershipMemberToTags struct {
	Key []string `json:"key,omitempty"` //
}
type RequestTagAddMembersToTheTag map[string][]string
type RequestTagQueryTheTagsAssociatedWithInterfaces struct {
	IDs []string `json:"ids,omitempty"` // List of member ids (network device or interface), maximum 500 ids can be passed.
}
type RequestTagQueryTheTagsAssociatedWithNetworkDevices struct {
	IDs []string `json:"ids,omitempty"` // List of member ids (network device or interface), maximum 500 ids can be passed.
}

//GetTag Get Tag - ee9a-ab01-487a-8896
/* Returns the tags for given filter criteria


@param GetTagQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag
*/
func (s *TagService) GetTag(GetTagQueryParams *GetTagQueryParams) (*ResponseTagGetTag, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag"

	queryString, _ := query.Values(GetTagQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTag{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTag(GetTagQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTag")
	}

	result := response.Result().(*ResponseTagGetTag)
	return result, response, err

}

//GetTagCount Get Tag Count - 8091-a9b8-4bfb-a53b
/* Returns tag count


@param GetTagCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-count
*/
func (s *TagService) GetTagCount(GetTagCountQueryParams *GetTagCountQueryParams) (*ResponseTagGetTagCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/count"

	queryString, _ := query.Values(GetTagCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagCount(GetTagCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagCount")
	}

	result := response.Result().(*ResponseTagGetTagCount)
	return result, response, err

}

//GetTagResourceTypes Get Tag resource types - 4695-090d-403b-8eaa
/* Returns list of supported resource types



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-resource-types
*/
func (s *TagService) GetTagResourceTypes() (*ResponseTagGetTagResourceTypes, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/member/type"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagGetTagResourceTypes{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagResourceTypes()
		}
		return nil, response, fmt.Errorf("error with operation GetTagResourceTypes")
	}

	result := response.Result().(*ResponseTagGetTagResourceTypes)
	return result, response, err

}

//GetTagByID Get Tag by Id - c1a3-59b1-4c89-b573
/* Returns tag specified by Id


@param id id path parameter. Tag ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-by-id
*/
func (s *TagService) GetTagByID(id string) (*ResponseTagGetTagByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagGetTagByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTagById")
	}

	result := response.Result().(*ResponseTagGetTagByID)
	return result, response, err

}

//GetTagMembersByID Get Tag members by Id - eab7-abe0-48fb-99ad
/* Returns tag members specified by id


@param id id path parameter. Tag ID

@param GetTagMembersByIdQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-members-by-id
*/
func (s *TagService) GetTagMembersByID(id string, GetTagMembersByIdQueryParams *GetTagMembersByIDQueryParams) (*ResponseTagGetTagMembersByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTagMembersByIdQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagMembersByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagMembersByID(id, GetTagMembersByIdQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagMembersById")
	}

	result := response.Result().(*ResponseTagGetTagMembersByID)
	return result, response, err

}

//GetTagMemberCount Get Tag Member count - 2e9d-b858-40fb-b1cf
/* Returns the number of members in a given tag


@param id id path parameter. Tag ID

@param GetTagMemberCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-tag-member-count
*/
func (s *TagService) GetTagMemberCount(id string, GetTagMemberCountQueryParams *GetTagMemberCountQueryParams) (*ResponseTagGetTagMemberCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}/member/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetTagMemberCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagGetTagMemberCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTagMemberCount(id, GetTagMemberCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTagMemberCount")
	}

	result := response.Result().(*ResponseTagGetTagMemberCount)
	return result, response, err

}

//RetrieveTagsAssociatedWithTheInterfaces Retrieve tags associated with the interfaces. - b786-6abb-47f8-8c83
/* Fetches the tags associated with the interfaces. Interfaces that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it is called a member of the tag.


@param RetrieveTagsAssociatedWithTheInterfacesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-tags-associated-with-the-interfaces
*/
func (s *TagService) RetrieveTagsAssociatedWithTheInterfaces(RetrieveTagsAssociatedWithTheInterfacesQueryParams *RetrieveTagsAssociatedWithTheInterfacesQueryParams) (*ResponseTagRetrieveTagsAssociatedWithTheInterfaces, *resty.Response, error) {
	path := "/intent/api/v1/tags/interfaces/membersAssociations"

	queryString, _ := query.Values(RetrieveTagsAssociatedWithTheInterfacesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagRetrieveTagsAssociatedWithTheInterfaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTagsAssociatedWithTheInterfaces(RetrieveTagsAssociatedWithTheInterfacesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTagsAssociatedWithTheInterfaces")
	}

	result := response.Result().(*ResponseTagRetrieveTagsAssociatedWithTheInterfaces)
	return result, response, err

}

//RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag Retrieve the count of interfaces that are associated with at least one tag. - 3cb8-f8e6-4bfa-928c
/* Fetches the count of interfaces that are associated with at least one tag. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it is called a member of the tag.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-interfaces-that-are-associated-with-at-least-one-tag
*/
func (s *TagService) RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag() (*ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag, *resty.Response, error) {
	path := "/intent/api/v1/tags/interfaces/membersAssociations/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag")
	}

	result := response.Result().(*ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTag)
	return result, response, err

}

//RetrieveTagsAssociatedWithNetworkDevices Retrieve tags associated with network devices. - 0b84-9b56-4bda-bc68
/* Fetches the tags associated with network devices. Devices that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When a device is tagged, it is called a member of the tag.


@param RetrieveTagsAssociatedWithNetworkDevicesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-tags-associated-with-network-devices
*/
func (s *TagService) RetrieveTagsAssociatedWithNetworkDevices(RetrieveTagsAssociatedWithNetworkDevicesQueryParams *RetrieveTagsAssociatedWithNetworkDevicesQueryParams) (*ResponseTagRetrieveTagsAssociatedWithNetworkDevices, *resty.Response, error) {
	path := "/intent/api/v1/tags/networkDevices/membersAssociations"

	queryString, _ := query.Values(RetrieveTagsAssociatedWithNetworkDevicesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseTagRetrieveTagsAssociatedWithNetworkDevices{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTagsAssociatedWithNetworkDevices(RetrieveTagsAssociatedWithNetworkDevicesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTagsAssociatedWithNetworkDevices")
	}

	result := response.Result().(*ResponseTagRetrieveTagsAssociatedWithNetworkDevices)
	return result, response, err

}

//RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag Retrieve the count of network devices that are associated with at least one tag. - 63b1-69be-4919-9dc0
/* Fetches the count of network devices that are associated with at least one tag. A tag is a user-defined or system-defined construct to group resources. When a device is tagged, it is called a member of the tag.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-count-of-network-devices-that-are-associated-with-at-least-one-tag
*/
func (s *TagService) RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag() (*ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag, *resty.Response, error) {
	path := "/intent/api/v1/tags/networkDevices/membersAssociations/count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag")
	}

	result := response.Result().(*ResponseTagRetrieveTheCountOfNetworkDevicesThatAreAssociatedWithAtLeastOneTag)
	return result, response, err

}

//CreateTag Create Tag - 1399-891c-42a8-be64
/* Creates tag with specified tag attributes



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-tag
*/
func (s *TagService) CreateTag(requestTagCreateTag *RequestTagCreateTag) (*ResponseTagCreateTag, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagCreateTag).
		SetResult(&ResponseTagCreateTag{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateTag(requestTagCreateTag)
		}

		return nil, response, fmt.Errorf("error with operation CreateTag")
	}

	result := response.Result().(*ResponseTagCreateTag)
	return result, response, err

}

//AddMembersToTheTag Add members to the tag - 00a2-fa61-4608-9317
/* Adds members to the tag specified by id


@param id id path parameter. Tag ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-members-to-the-tag
*/
func (s *TagService) AddMembersToTheTag(id string, requestTagAddMembersToTheTag *RequestTagAddMembersToTheTag) (*ResponseTagAddMembersToTheTag, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagAddMembersToTheTag).
		SetResult(&ResponseTagAddMembersToTheTag{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddMembersToTheTag(id, requestTagAddMembersToTheTag)
		}

		return nil, response, fmt.Errorf("error with operation AddMembersToTheTag")
	}

	result := response.Result().(*ResponseTagAddMembersToTheTag)
	return result, response, err

}

//QueryTheTagsAssociatedWithInterfaces Query the tags associated with interfaces. - 87a2-4a4e-4109-becf
/* Fetches the tags associated with the given interface 'ids'. Interfaces that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it is called a member of the tag. 'ids' can be fetched via '/dna/intent/api/v1/interface' API.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-the-tags-associated-with-interfaces
*/
func (s *TagService) QueryTheTagsAssociatedWithInterfaces(requestTagQueryTheTagsAssociatedWithInterfaces *RequestTagQueryTheTagsAssociatedWithInterfaces) (*ResponseTagQueryTheTagsAssociatedWithInterfaces, *resty.Response, error) {
	path := "/intent/api/v1/tags/interfaces/membersAssociations/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagQueryTheTagsAssociatedWithInterfaces).
		SetResult(&ResponseTagQueryTheTagsAssociatedWithInterfaces{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryTheTagsAssociatedWithInterfaces(requestTagQueryTheTagsAssociatedWithInterfaces)
		}

		return nil, response, fmt.Errorf("error with operation QueryTheTagsAssociatedWithInterfaces")
	}

	result := response.Result().(*ResponseTagQueryTheTagsAssociatedWithInterfaces)
	return result, response, err

}

//QueryTheTagsAssociatedWithNetworkDevices Query the tags associated with network devices. - 6480-fa01-417b-b397
/* Fetches the tags associated with the given network device 'ids'. Devices that don't have any tags associated will not be included in the response. A tag is a user-defined or system-defined construct to group resources. When a device is tagged, it is called a member of the tag. 'ids' can be fetched via '/dna/intent/api/v1/network-device' API.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-the-tags-associated-with-network-devices
*/
func (s *TagService) QueryTheTagsAssociatedWithNetworkDevices(requestTagQueryTheTagsAssociatedWithNetworkDevices *RequestTagQueryTheTagsAssociatedWithNetworkDevices) (*ResponseTagQueryTheTagsAssociatedWithNetworkDevices, *resty.Response, error) {
	path := "/intent/api/v1/tags/networkDevices/membersAssociations/query"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagQueryTheTagsAssociatedWithNetworkDevices).
		SetResult(&ResponseTagQueryTheTagsAssociatedWithNetworkDevices{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryTheTagsAssociatedWithNetworkDevices(requestTagQueryTheTagsAssociatedWithNetworkDevices)
		}

		return nil, response, fmt.Errorf("error with operation QueryTheTagsAssociatedWithNetworkDevices")
	}

	result := response.Result().(*ResponseTagQueryTheTagsAssociatedWithNetworkDevices)
	return result, response, err

}

//UpdateTag Update Tag - 4d86-a993-469a-9da9
/* Updates a tag specified by id


 */
func (s *TagService) UpdateTag(requestTagUpdateTag *RequestTagUpdateTag) (*ResponseTagUpdateTag, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagUpdateTag).
		SetResult(&ResponseTagUpdateTag{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTag(requestTagUpdateTag)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTag")
	}

	result := response.Result().(*ResponseTagUpdateTag)
	return result, response, err

}

//UpdateTagMembership Update tag membership - 45bc-7a83-44a8-bc1e
/* Update tag membership. As part of the request payload through this API, only the specified members are added / retained to the given input tags. Possible values of memberType attribute in the request payload can be queried by using the /tag/member/type API


 */
func (s *TagService) UpdateTagMembership(requestTagUpdateTagMembership *RequestTagUpdateTagMembership) (*ResponseTagUpdateTagMembership, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/member"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagUpdateTagMembership).
		SetResult(&ResponseTagUpdateTagMembership{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTagMembership(requestTagUpdateTagMembership)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTagMembership")
	}

	result := response.Result().(*ResponseTagUpdateTagMembership)
	return result, response, err

}

//DeleteTag Delete Tag - 429c-2815-4bda-a13d
/* Deletes a tag specified by id


@param id id path parameter. Tag ID


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-tag
*/
func (s *TagService) DeleteTag(id string) (*ResponseTagDeleteTag, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagDeleteTag{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTag(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTag")
	}

	result := response.Result().(*ResponseTagDeleteTag)
	return result, response, err

}

//RemoveTagMember Remove Tag member - caa3-ea70-4d78-b37e
/* Removes Tag member from the tag specified by id


@param id id path parameter. Tag ID

@param memberID memberId path parameter. TagMember id to be removed from tag


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-tag-member
*/
func (s *TagService) RemoveTagMember(id string, memberID string) (*ResponseTagRemoveTagMember, *resty.Response, error) {
	//id string,memberID string
	path := "/dna/intent/api/v1/tag/{id}/member/{memberId}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{memberId}", fmt.Sprintf("%v", memberID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseTagRemoveTagMember{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveTagMember(id, memberID)
		}
		return nil, response, fmt.Errorf("error with operation RemoveTagMember")
	}

	result := response.Result().(*ResponseTagRemoveTagMember)
	return result, response, err

}
