package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// TagService is the service to communicate with the Tag API endpoint
type TagService service

// DynamicRules is the DynamicRules definition
type DynamicRules struct {
	MemberType string `json:"memberType,omitempty"` //
	Rules      Rules  `json:"rules,omitempty"`      //
}

// MemberToTags is the MemberToTags definition
type MemberToTags struct {
	Key []string `json:"key,omitempty"` //
}

// Rules is the Rules definition
type Rules struct {
	Items     []string `json:"items,omitempty"`     //
	Name      string   `json:"name,omitempty"`      //
	Operation string   `json:"operation,omitempty"` //
	Value     string   `json:"value,omitempty"`     //
	Values    []string `json:"values,omitempty"`    //
}

// TagDTO is the TagDTO definition
type TagDTO struct {
	Description      string         `json:"description,omitempty"`      //
	DynamicRules     []DynamicRules `json:"dynamicRules,omitempty"`     //
	Id               string         `json:"id,omitempty"`               //
	InstanceTenantId string         `json:"instanceTenantId,omitempty"` //
	Name             string         `json:"name,omitempty"`             //
	SystemTag        bool           `json:"systemTag,omitempty"`        //
}

// TagMemberDTO is the TagMemberDTO definition
type TagMemberDTO struct {
	MemberToTags []MemberToTags `json:"memberToTags,omitempty"` //
	MemberType   string         `json:"memberType,omitempty"`   //
}

// CountResult is the CountResult definition
type CountResult struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// DynamicRules is the DynamicRules definition
type DynamicRules struct {
	MemberType string `json:"memberType,omitempty"` //
	Rules      Rules  `json:"rules,omitempty"`      //
}

// Response is the Response definition
type Response struct {
	TaskId string `json:"taskId,omitempty"` //
	Url    string `json:"url,omitempty"`    //
}

// Rules is the Rules definition
type Rules struct {
	Items     []string `json:"items,omitempty"`     //
	Name      string   `json:"name,omitempty"`      //
	Operation string   `json:"operation,omitempty"` //
	Value     string   `json:"value,omitempty"`     //
	Values    []string `json:"values,omitempty"`    //
}

// TagListResult is the TagListResult definition
type TagListResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// TagMembersResult is the TagMembersResult definition
type TagMembersResult struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// TagResult is the TagResult definition
type TagResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// TagTypesResult is the TagTypesResult definition
type TagTypesResult struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// TaskIdResult is the TaskIdResult definition
type TaskIdResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// AddMembersToTheTag addMembersToTheTag
/* Adds members to the tag specified by id
@param id Tag ID
*/
func (s *TagService) AddMembersToTheTag(id string, addMembersToTheTagRequest *AddMembersToTheTagRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetBody(addMembersToTheTagRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// CreateTag createTag
/* Creates tag with specified tag attributes
 */
func (s *TagService) CreateTag(createTagRequest *CreateTagRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag"

	response, err := RestyClient.R().
		SetBody(createTagRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// DeleteTag deleteTag
/* Deletes a tag specified by id
@param id Tag ID
*/
func (s *TagService) DeleteTag(id string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetTagQueryParams defines the query parameters for this request
type GetTagQueryParams struct {
	Name                   string `url:"name,omitempty"`                      // Tag name is mandatory when filter operation is used.
	AdditionalInfameSpace  string `url:"additionalInfo.nameSpace,omitempty"`  // nameSpace
	AdditionalInfttributes string `url:"additionalInfo.attributes,omitempty"` // attributeName
	Level                  string `url:"level,omitempty"`                     // levelArg
	Offset                 string `url:"offset,omitempty"`                    // offset
	Limit                  string `url:"limit,omitempty"`                     // limit
	Size                   string `url:"size,omitempty"`                      // size in kilobytes(KB)
	Field                  string `url:"field,omitempty"`                     // Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
	SortBy                 string `url:"sortBy,omitempty"`                    // Only supported attribute is name. SortyBy is mandatory when order is used.
	Order                  string `url:"order,omitempty"`                     // Available values are asc and des
	SystemTag              string `url:"systemTag,omitempty"`                 // systemTag
}

// GetTag getTag
/* Returns the tags for given filter criteria
@param name Tag name is mandatory when filter operation is used.
@param additionalInfo.nameSpace nameSpace
@param additionalInfo.attributes attributeName
@param level levelArg
@param offset offset
@param limit limit
@param size size in kilobytes(KB)
@param field Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
@param sortBy Only supported attribute is name. SortyBy is mandatory when order is used.
@param order Available values are asc and des
@param systemTag systemTag
*/
func (s *TagService) GetTag(getTagQueryParams *GetTagQueryParams) (*TagListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag"

	queryString, _ := query.Values(getTagQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&TagListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TagListResult)
	return result, response, err

}

// GetTagById getTagById
/* Returns tag specified by Id
@param id Tag ID
*/
func (s *TagService) GetTagById(id string) (*TagResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&TagResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TagResult)
	return result, response, err

}

// GetTagCountQueryParams defines the query parameters for this request
type GetTagCountQueryParams struct {
	Name          string `url:"name,omitempty"`          // tagName
	NameSpace     string `url:"nameSpace,omitempty"`     // nameSpace
	AttributeName string `url:"attributeName,omitempty"` // attributeName
	Level         string `url:"level,omitempty"`         // levelArg
	Size          string `url:"size,omitempty"`          // size in kilobytes(KB)
	SystemTag     string `url:"systemTag,omitempty"`     // systemTag
}

// GetTagCount getTagCount
/* Returns tag count
@param name tagName
@param nameSpace nameSpace
@param attributeName attributeName
@param level levelArg
@param size size in kilobytes(KB)
@param systemTag systemTag
*/
func (s *TagService) GetTagCount(getTagCountQueryParams *GetTagCountQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/count"

	queryString, _ := query.Values(getTagCountQueryParams)

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

// GetTagMemberCountQueryParams defines the query parameters for this request
type GetTagMemberCountQueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            // memberType
	MemberAssociationType string `url:"memberAssociationType,omitempty"` // memberAssociationType
	Level                 string `url:"level,omitempty"`                 // level
}

// GetTagMemberCount getTagMemberCount
/* Returns the number of members in a given tag
@param id Tag ID
@param memberType memberType
@param memberAssociationType memberAssociationType
@param level level
*/
func (s *TagService) GetTagMemberCount(id string, getTagMemberCountQueryParams *GetTagMemberCountQueryParams) (*CountResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member/count"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getTagMemberCountQueryParams)

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

// GetTagMembersByIdQueryParams defines the query parameters for this request
type GetTagMembersByIdQueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            // Entity type of the member. Possible values can be retrieved by using /tag/member/type API
	Offset                string `url:"offset,omitempty"`                // Used for pagination. It indicates the starting row number out of available member records
	Limit                 string `url:"limit,omitempty"`                 // Used to Number of maximum members to return in the result
	MemberAssociationType string `url:"memberAssociationType,omitempty"` // Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
	Level                 string `url:"level,omitempty"`                 // level
}

// GetTagMembersById getTagMembersById
/* Returns tag members specified by id
@param id Tag ID
@param memberType Entity type of the member. Possible values can be retrieved by using /tag/member/type API
@param offset Used for pagination. It indicates the starting row number out of available member records
@param limit Used to Number of maximum members to return in the result
@param memberAssociationType Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
@param level level
*/
func (s *TagService) GetTagMembersById(id string, getTagMembersByIdQueryParams *GetTagMembersByIdQueryParams) (*TagMembersResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getTagMembersByIdQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&TagMembersResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TagMembersResult)
	return result, response, err

}

// GetTagResourceTypes getTagResourceTypes
/* Returns list of supported resource types
 */
func (s *TagService) GetTagResourceTypes() (*TagTypesResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/member/type"

	response, err := RestyClient.R().
		SetResult(&TagTypesResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TagTypesResult)
	return result, response, err

}

// RemoveTagMember removeTagMember
/* Removes Tag member from the tag specified by id
@param id Tag ID
@param memberId TagMember id to be removed from tag
*/
func (s *TagService) RemoveTagMember(id string, memberId string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member/{memberId}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"memberId"+"}", fmt.Sprintf("%v", memberId), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// UpdateTag updateTag
/* Updates a tag specified by id
 */
func (s *TagService) UpdateTag(updateTagRequest *UpdateTagRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag"

	response, err := RestyClient.R().
		SetBody(updateTagRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// UpdatesTagMembership updatesTagMembership
/* Updates tag membership. As part of the request payload through this API, only the specified members are added / retained to the given input tags. Possible values of memberType attribute in the request payload can be queried by using the /tag/member/type API
 */
func (s *TagService) UpdatesTagMembership(updatesTagMembershipRequest *UpdatesTagMembershipRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/member"

	response, err := RestyClient.R().
		SetBody(updatesTagMembershipRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}
