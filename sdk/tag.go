package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type TagService service

type GetTagQueryParams struct {
	Name                     string `url:"name,omitempty"`                      //Tag name is mandatory when filter operation is used.
	AdditionalInfonameSpace  string `url:"additionalInfo.nameSpace,omitempty"`  //nameSpace
	AdditionalInfoattributes string `url:"additionalInfo.attributes,omitempty"` //attributeName
	Level                    string `url:"level,omitempty"`                     //levelArg
	Offset                   int    `url:"offset,omitempty"`                    //offset
	Limit                    int    `url:"limit,omitempty"`                     //limit
	Size                     string `url:"size,omitempty"`                      //size in kilobytes(KB)
	Field                    string `url:"field,omitempty"`                     //Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
	SortBy                   string `url:"sortBy,omitempty"`                    //Only supported attribute is name. SortyBy is mandatory when order is used.
	Order                    string `url:"order,omitempty"`                     //Available values are asc and des
	SystemTag                string `url:"systemTag,omitempty"`                 //systemTag
}
type GetTagCountQueryParams struct {
	Name          string `url:"name,omitempty"`          //tagName
	NameSpace     string `url:"nameSpace,omitempty"`     //nameSpace
	AttributeName string `url:"attributeName,omitempty"` //attributeName
	Level         string `url:"level,omitempty"`         //levelArg
	Size          string `url:"size,omitempty"`          //size in kilobytes(KB)
	SystemTag     string `url:"systemTag,omitempty"`     //systemTag
}
type GetTagMembersByIDQueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            //Entity type of the member. Possible values can be retrieved by using /tag/member/type API
	Offset                string `url:"offset,omitempty"`                //Used for pagination. It indicates the starting row number out of available member records
	Limit                 string `url:"limit,omitempty"`                 //Used to Number of maximum members to return in the result
	MemberAssociationType string `url:"memberAssociationType,omitempty"` //Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
	Level                 string `url:"level,omitempty"`                 //level
}
type GetTagMemberCountQueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            //memberType
	MemberAssociationType string `url:"memberAssociationType,omitempty"` //memberAssociationType
	Level                 string `url:"level,omitempty"`                 //level
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
	Values    []string `json:"values,omitempty"`    //
	Items     string   `json:"items,omitempty"`     //
	Operation string   `json:"operation,omitempty"` //
	Name      string   `json:"name,omitempty"`      //
	Value     string   `json:"value,omitempty"`     //
}
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
type ResponseTagUpdatesTagMembership struct {
	Version  string                                   `json:"version,omitempty"`  //
	Response *ResponseTagUpdatesTagMembershipResponse `json:"response,omitempty"` //
}
type ResponseTagUpdatesTagMembershipResponse struct {
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
	Values    []string `json:"values,omitempty"`    //
	Items     string   `json:"items,omitempty"`     //
	Operation string   `json:"operation,omitempty"` //
	Name      string   `json:"name,omitempty"`      //
	Value     string   `json:"value,omitempty"`     //
}
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
type RequestTagUpdateTag struct {
	SystemTag        *bool                              `json:"systemTag,omitempty"`        //
	Description      string                             `json:"description,omitempty"`      //
	DynamicRules     *[]RequestTagUpdateTagDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                             `json:"name,omitempty"`             //
	ID               string                             `json:"id,omitempty"`               //
	InstanceTenantID string                             `json:"instanceTenantId,omitempty"` //
}
type RequestTagUpdateTagDynamicRules struct {
	MemberType string                                `json:"memberType,omitempty"` //
	Rules      *RequestTagUpdateTagDynamicRulesRules `json:"rules,omitempty"`      //
}
type RequestTagUpdateTagDynamicRulesRules struct {
	Values    []string `json:"values,omitempty"`    //
	Items     []string `json:"items,omitempty"`     //
	Operation string   `json:"operation,omitempty"` //
	Name      string   `json:"name,omitempty"`      //
	Value     string   `json:"value,omitempty"`     //
}
type RequestTagCreateTag struct {
	SystemTag        *bool                              `json:"systemTag,omitempty"`        //
	Description      string                             `json:"description,omitempty"`      //
	DynamicRules     *[]RequestTagCreateTagDynamicRules `json:"dynamicRules,omitempty"`     //
	Name             string                             `json:"name,omitempty"`             //
	ID               string                             `json:"id,omitempty"`               //
	InstanceTenantID string                             `json:"instanceTenantId,omitempty"` //
}
type RequestTagCreateTagDynamicRules struct {
	MemberType string                                `json:"memberType,omitempty"` //
	Rules      *RequestTagCreateTagDynamicRulesRules `json:"rules,omitempty"`      //
}
type RequestTagCreateTagDynamicRulesRules struct {
	Values    []string `json:"values,omitempty"`    //
	Items     string   `json:"items,omitempty"`     //
	Operation string   `json:"operation,omitempty"` //
	Name      string   `json:"name,omitempty"`      //
	Value     string   `json:"value,omitempty"`     //
}
type RequestTagUpdatesTagMembership struct {
	MemberToTags map[string][]string `json:"memberToTags,omitempty"` //
	MemberType   string              `json:"memberType,omitempty"`   //
}
type RequestTagAddMembersToTheTag map[string][]string

//GetTag Get Tag - ee9a-ab01-487a-8896
/* Returns the tags for given filter criteria


@param GetTagQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetTag")
	}

	result := response.Result().(*ResponseTagGetTag)
	return result, response, err

}

//GetTagCount Get Tag Count - 8091-a9b8-4bfb-a53b
/* Returns tag count


@param GetTagCountQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetTagCount")
	}

	result := response.Result().(*ResponseTagGetTagCount)
	return result, response, err

}

//GetTagResourceTypes Get Tag resource types - 4695-090d-403b-8eaa
/* Returns list of supported resource types


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
		return nil, response, fmt.Errorf("error with operation GetTagResourceTypes")
	}

	result := response.Result().(*ResponseTagGetTagResourceTypes)
	return result, response, err

}

//GetTagByID Get Tag by Id - c1a3-59b1-4c89-b573
/* Returns tag specified by Id


@param id id path parameter. Tag ID

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
		return nil, response, fmt.Errorf("error with operation GetTagById")
	}

	result := response.Result().(*ResponseTagGetTagByID)
	return result, response, err

}

//GetTagMembersByID Get Tag members by Id - eab7-abe0-48fb-99ad
/* Returns tag members specified by id


@param id id path parameter. Tag ID

@param GetTagMembersByIdQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetTagMembersById")
	}

	result := response.Result().(*ResponseTagGetTagMembersByID)
	return result, response, err

}

//GetTagMemberCount Get Tag Member count - 2e9d-b858-40fb-b1cf
/* Returns the number of members in a given tag


@param id id path parameter. Tag ID

@param GetTagMemberCountQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetTagMemberCount")
	}

	result := response.Result().(*ResponseTagGetTagMemberCount)
	return result, response, err

}

//CreateTag Create Tag - 1399-891c-42a8-be64
/* Creates tag with specified tag attributes


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
		return nil, response, fmt.Errorf("error with operation CreateTag")
	}

	result := response.Result().(*ResponseTagCreateTag)
	return result, response, err

}

//AddMembersToTheTag Add members to the tag - 00a2-fa61-4608-9317
/* Adds members to the tag specified by id


@param id id path parameter. Tag ID

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
		return nil, response, fmt.Errorf("error with operation AddMembersToTheTag")
	}

	result := response.Result().(*ResponseTagAddMembersToTheTag)
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
		return nil, response, fmt.Errorf("error with operation UpdateTag")
	}

	result := response.Result().(*ResponseTagUpdateTag)
	return result, response, err

}

//UpdatesTagMembership Updates tag membership - 45bc-7a83-44a8-bc1e
/* Updates tag membership. As part of the request payload through this API, only the specified members are added / retained to the given input tags. Possible values of memberType attribute in the request payload can be queried by using the /tag/member/type API


 */
func (s *TagService) UpdatesTagMembership(requestTagUpdatesTagMembership *RequestTagUpdatesTagMembership) (*ResponseTagUpdatesTagMembership, *resty.Response, error) {
	path := "/dna/intent/api/v1/tag/member"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestTagUpdatesTagMembership).
		SetResult(&ResponseTagUpdatesTagMembership{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation UpdatesTagMembership")
	}

	result := response.Result().(*ResponseTagUpdatesTagMembership)
	return result, response, err

}

//DeleteTag Delete Tag - 429c-2815-4bda-a13d
/* Deletes a tag specified by id


@param id id path parameter. Tag ID

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
		return nil, response, fmt.Errorf("error with operation DeleteTag")
	}

	result := response.Result().(*ResponseTagDeleteTag)
	return result, response, err

}

//RemoveTagMember Remove Tag member - caa3-ea70-4d78-b37e
/* Removes Tag member from the tag specified by id


@param id id path parameter. Tag ID

@param memberID memberId path parameter. TagMember id to be removed from tag

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
		return nil, response, fmt.Errorf("error with operation RemoveTagMember")
	}

	result := response.Result().(*ResponseTagRemoveTagMember)
	return result, response, err

}
