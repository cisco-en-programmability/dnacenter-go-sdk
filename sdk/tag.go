package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// TagService is the service to communicate with the Tag API endpoint
type TagService service

// CreateTagRequest is the createTagRequest definition
type CreateTagRequest struct {
	Description      string                         `json:"description,omitempty"`      //
	DynamicRules     []CreateTagRequestDynamicRules `json:"dynamicRules,omitempty"`     //
	ID               string                         `json:"id,omitempty"`               //
	InstanceTenantID string                         `json:"instanceTenantId,omitempty"` //
	Name             string                         `json:"name,omitempty"`             //
	SystemTag        bool                           `json:"systemTag,omitempty"`        //
}

// CreateTagRequestDynamicRules is the createTagRequestDynamicRules definition
type CreateTagRequestDynamicRules struct {
	MemberType string                            `json:"memberType,omitempty"` //
	Rules      CreateTagRequestDynamicRulesRules `json:"rules,omitempty"`      //
}

// CreateTagRequestDynamicRulesRules is the createTagRequestDynamicRulesRules definition
type CreateTagRequestDynamicRulesRules struct {
	Items     []string `json:"items,omitempty"`     //
	Name      string   `json:"name,omitempty"`      //
	Operation string   `json:"operation,omitempty"` //
	Value     string   `json:"value,omitempty"`     //
	Values    []string `json:"values,omitempty"`    //
}

// CreateTagRequestDynamicRulesRulesItems is the createTagRequestDynamicRulesRulesItems definition
type CreateTagRequestDynamicRulesRulesItems []string

// CreateTagRequestDynamicRulesRulesValues is the createTagRequestDynamicRulesRulesValues definition
type CreateTagRequestDynamicRulesRulesValues []string

// UpdateTagRequest is the updateTagRequest definition
type UpdateTagRequest struct {
	Description      string                         `json:"description,omitempty"`      //
	DynamicRules     []UpdateTagRequestDynamicRules `json:"dynamicRules,omitempty"`     //
	ID               string                         `json:"id,omitempty"`               //
	InstanceTenantID string                         `json:"instanceTenantId,omitempty"` //
	Name             string                         `json:"name,omitempty"`             //
	SystemTag        bool                           `json:"systemTag,omitempty"`        //
}

// UpdateTagRequestDynamicRules is the updateTagRequestDynamicRules definition
type UpdateTagRequestDynamicRules struct {
	MemberType string                            `json:"memberType,omitempty"` //
	Rules      UpdateTagRequestDynamicRulesRules `json:"rules,omitempty"`      //
}

// UpdateTagRequestDynamicRulesRules is the updateTagRequestDynamicRulesRules definition
type UpdateTagRequestDynamicRulesRules struct {
	Items     []string `json:"items,omitempty"`     //
	Name      string   `json:"name,omitempty"`      //
	Operation string   `json:"operation,omitempty"` //
	Value     string   `json:"value,omitempty"`     //
	Values    []string `json:"values,omitempty"`    //
}

// UpdateTagRequestDynamicRulesRulesItems is the updateTagRequestDynamicRulesRulesItems definition
type UpdateTagRequestDynamicRulesRulesItems []string

// UpdateTagRequestDynamicRulesRulesValues is the updateTagRequestDynamicRulesRulesValues definition
type UpdateTagRequestDynamicRulesRulesValues []string

// UpdatesTagMembershipRequest is the updatesTagMembershipRequest definition
type UpdatesTagMembershipRequest struct {
	MemberToTags []UpdatesTagMembershipRequestMemberToTags `json:"memberToTags,omitempty"` //
	MemberType   string                                    `json:"memberType,omitempty"`   //
}

// UpdatesTagMembershipRequestMemberToTags is the updatesTagMembershipRequestMemberToTags definition
type UpdatesTagMembershipRequestMemberToTags struct {
	Key []string `json:"key,omitempty"` //
}

// UpdatesTagMembershipRequestMemberToTagsKey is the updatesTagMembershipRequestMemberToTagsKey definition
type UpdatesTagMembershipRequestMemberToTagsKey []string

// AddMembersToTheTagResponse is the addMembersToTheTagResponse definition
type AddMembersToTheTagResponse struct {
	Response AddMembersToTheTagResponseResponse `json:"response,omitempty"` //
	Version  string                             `json:"version,omitempty"`  //
}

// AddMembersToTheTagResponseResponse is the addMembersToTheTagResponseResponse definition
type AddMembersToTheTagResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// CreateTagResponse is the createTagResponse definition
type CreateTagResponse struct {
	Response CreateTagResponseResponse `json:"response,omitempty"` //
	Version  string                    `json:"version,omitempty"`  //
}

// CreateTagResponseResponse is the createTagResponseResponse definition
type CreateTagResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeleteTagResponse is the deleteTagResponse definition
type DeleteTagResponse struct {
	Response DeleteTagResponseResponse `json:"response,omitempty"` //
	Version  string                    `json:"version,omitempty"`  //
}

// DeleteTagResponseResponse is the deleteTagResponseResponse definition
type DeleteTagResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// GetTagByIDResponseResponse is the getTagByIDResponseResponse definition
type GetTagByIDResponseResponse struct {
	Description      string                                   `json:"description,omitempty"`      //
	DynamicRules     []GetTagByIDResponseResponseDynamicRules `json:"dynamicRules,omitempty"`     //
	ID               string                                   `json:"id,omitempty"`               //
	InstanceTenantID string                                   `json:"instanceTenantId,omitempty"` //
	Name             string                                   `json:"name,omitempty"`             //
	SystemTag        bool                                     `json:"systemTag,omitempty"`        //
}

// GetTagByIDResponseResponseDynamicRules is the getTagByIDResponseResponseDynamicRules definition
type GetTagByIDResponseResponseDynamicRules struct {
	MemberType string                                      `json:"memberType,omitempty"` //
	Rules      GetTagByIDResponseResponseDynamicRulesRules `json:"rules,omitempty"`      //
}

// GetTagByIDResponseResponseDynamicRulesRules is the getTagByIDResponseResponseDynamicRulesRules definition
type GetTagByIDResponseResponseDynamicRulesRules struct {
	Items     []string `json:"items,omitempty"`     //
	Name      string   `json:"name,omitempty"`      //
	Operation string   `json:"operation,omitempty"` //
	Value     string   `json:"value,omitempty"`     //
	Values    []string `json:"values,omitempty"`    //
}

// GetTagByIDResponseResponseDynamicRulesRulesItems is the getTagByIDResponseResponseDynamicRulesRulesItems definition
type GetTagByIDResponseResponseDynamicRulesRulesItems []string

// GetTagByIDResponseResponseDynamicRulesRulesValues is the getTagByIDResponseResponseDynamicRulesRulesValues definition
type GetTagByIDResponseResponseDynamicRulesRulesValues []string

// GetTagByIDResponse is the getTagByIdResponse definition
type GetTagByIDResponse struct {
	Response GetTagByIDResponseResponse `json:"response,omitempty"` //
	Version  string                     `json:"version,omitempty"`  //
}

// GetTagCountResponse is the getTagCountResponse definition
type GetTagCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetTagMemberCountResponse is the getTagMemberCountResponse definition
type GetTagMemberCountResponse struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}

// GetTagMembersByIDResponseResponse is the getTagMembersByIDResponseResponse definition
type GetTagMembersByIDResponseResponse struct {
	InstanceUUID string `json:"instanceUuid,omitempty"` //
}

// GetTagMembersByIDResponse is the getTagMembersByIdResponse definition
type GetTagMembersByIDResponse struct {
	Response []GetTagMembersByIDResponseResponse `json:"response,omitempty"` //
	Version  string                              `json:"version,omitempty"`  //
}

// GetTagResourceTypesResponse is the getTagResourceTypesResponse definition
type GetTagResourceTypesResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetTagResourceTypesResponseResponse is the getTagResourceTypesResponseResponse definition
type GetTagResourceTypesResponseResponse []string

// GetTagResponse is the getTagResponse definition
type GetTagResponse struct {
	Response []GetTagResponseResponse `json:"response,omitempty"` //
	Version  string                   `json:"version,omitempty"`  //
}

// GetTagResponseResponse is the getTagResponseResponse definition
type GetTagResponseResponse struct {
	Description      string                               `json:"description,omitempty"`      //
	DynamicRules     []GetTagResponseResponseDynamicRules `json:"dynamicRules,omitempty"`     //
	ID               string                               `json:"id,omitempty"`               //
	InstanceTenantID string                               `json:"instanceTenantId,omitempty"` //
	Name             string                               `json:"name,omitempty"`             //
	SystemTag        bool                                 `json:"systemTag,omitempty"`        //
}

// GetTagResponseResponseDynamicRules is the getTagResponseResponseDynamicRules definition
type GetTagResponseResponseDynamicRules struct {
	MemberType string                                  `json:"memberType,omitempty"` //
	Rules      GetTagResponseResponseDynamicRulesRules `json:"rules,omitempty"`      //
}

// GetTagResponseResponseDynamicRulesRules is the getTagResponseResponseDynamicRulesRules definition
type GetTagResponseResponseDynamicRulesRules struct {
	Items     []string `json:"items,omitempty"`     //
	Name      string   `json:"name,omitempty"`      //
	Operation string   `json:"operation,omitempty"` //
	Value     string   `json:"value,omitempty"`     //
	Values    []string `json:"values,omitempty"`    //
}

// GetTagResponseResponseDynamicRulesRulesItems is the getTagResponseResponseDynamicRulesRulesItems definition
type GetTagResponseResponseDynamicRulesRulesItems []string

// GetTagResponseResponseDynamicRulesRulesValues is the getTagResponseResponseDynamicRulesRulesValues definition
type GetTagResponseResponseDynamicRulesRulesValues []string

// RemoveTagMemberResponse is the removeTagMemberResponse definition
type RemoveTagMemberResponse struct {
	Response RemoveTagMemberResponseResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}

// RemoveTagMemberResponseResponse is the removeTagMemberResponseResponse definition
type RemoveTagMemberResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// UpdateTagResponse is the updateTagResponse definition
type UpdateTagResponse struct {
	Response UpdateTagResponseResponse `json:"response,omitempty"` //
	Version  string                    `json:"version,omitempty"`  //
}

// UpdateTagResponseResponse is the updateTagResponseResponse definition
type UpdateTagResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// UpdatesTagMembershipResponse is the updatesTagMembershipResponse definition
type UpdatesTagMembershipResponse struct {
	Response UpdatesTagMembershipResponseResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

// UpdatesTagMembershipResponseResponse is the updatesTagMembershipResponseResponse definition
type UpdatesTagMembershipResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// AddMembersToTheTag addMembersToTheTag
/* Adds members to the tag specified by id
@param id Tag ID
*/
// func (s *TagService) AddMembersToTheTag(id string, addMembersToTheTagRequest *AddMembersToTheTagRequest) (*AddMembersToTheTagResponse, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/tag/{id}/member"
// 	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

// 	response, err := RestyClient.R().
// 		SetBody(addMembersToTheTagRequest).
// 		SetResult(&AddMembersToTheTagResponse{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	if response.IsError() {
// 		return nil, response, fmt.Errorf("Error with operation addMembersToTheTag")
// 	}

// 	result := response.Result().(*AddMembersToTheTagResponse)
// 	return result, response, err
//}

// CreateTag createTag
/* Creates tag with specified tag attributes
 */
func (s *TagService) CreateTag(createTagRequest *CreateTagRequest) (*CreateTagResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag"

	response, err := RestyClient.R().
		SetBody(createTagRequest).
		SetResult(&CreateTagResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createTag")
	}

	result := response.Result().(*CreateTagResponse)
	return result, response, err
}

// DeleteTag deleteTag
/* Deletes a tag specified by id
@param id Tag ID
*/
func (s *TagService) DeleteTag(id string) (*DeleteTagResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteTag")
	}

	result := response.Result().(*DeleteTagResponse)
	return result, response, err
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
func (s *TagService) GetTag(getTagQueryParams *GetTagQueryParams) (*GetTagResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag"

	queryString, _ := query.Values(getTagQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetTagResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTag")
	}

	result := response.Result().(*GetTagResponse)
	return result, response, err
}

// GetTagByID getTagById
/* Returns tag specified by Id
@param id Tag ID
*/
func (s *TagService) GetTagByID(id string) (*GetTagByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	response, err := RestyClient.R().
		SetResult(&GetTagByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTagById")
	}

	result := response.Result().(*GetTagByIDResponse)
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
func (s *TagService) GetTagCount(getTagCountQueryParams *GetTagCountQueryParams) (*GetTagCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/count"

	queryString, _ := query.Values(getTagCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetTagCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTagCount")
	}

	result := response.Result().(*GetTagCountResponse)
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
func (s *TagService) GetTagMemberCount(id string, getTagMemberCountQueryParams *GetTagMemberCountQueryParams) (*GetTagMemberCountResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member/count"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getTagMemberCountQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetTagMemberCountResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTagMemberCount")
	}

	result := response.Result().(*GetTagMemberCountResponse)
	return result, response, err
}

// GetTagMembersByIDQueryParams defines the query parameters for this request
type GetTagMembersByIDQueryParams struct {
	MemberType            string `url:"memberType,omitempty"`            // Entity type of the member. Possible values can be retrieved by using /tag/member/type API
	Offset                string `url:"offset,omitempty"`                // Used for pagination. It indicates the starting row number out of available member records
	Limit                 string `url:"limit,omitempty"`                 // Used to Number of maximum members to return in the result
	MemberAssociationType string `url:"memberAssociationType,omitempty"` // Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
	Level                 string `url:"level,omitempty"`                 // level
}

// GetTagMembersByID getTagMembersById
/* Returns tag members specified by id
@param id Tag ID
@param memberType Entity type of the member. Possible values can be retrieved by using /tag/member/type API
@param offset Used for pagination. It indicates the starting row number out of available member records
@param limit Used to Number of maximum members to return in the result
@param memberAssociationType Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
@param level level
*/
func (s *TagService) GetTagMembersByID(id string, getTagMembersByIDQueryParams *GetTagMembersByIDQueryParams) (*GetTagMembersByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(getTagMembersByIDQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetTagMembersByIDResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTagMembersById")
	}

	result := response.Result().(*GetTagMembersByIDResponse)
	return result, response, err
}

// GetTagResourceTypes getTagResourceTypes
/* Returns list of supported resource types
 */
func (s *TagService) GetTagResourceTypes() (*GetTagResourceTypesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/member/type"

	response, err := RestyClient.R().
		SetResult(&GetTagResourceTypesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTagResourceTypes")
	}

	result := response.Result().(*GetTagResourceTypesResponse)
	return result, response, err
}

// RemoveTagMember removeTagMember
/* Removes Tag member from the tag specified by id
@param id Tag ID
@param memberID TagMember id to be removed from tag
*/
func (s *TagService) RemoveTagMember(id string, memberID string) (*RemoveTagMemberResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/{id}/member/{memberId}"
	path = strings.Replace(path, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{"+"memberId"+"}", fmt.Sprintf("%v", memberID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation removeTagMember")
	}

	result := response.Result().(*RemoveTagMemberResponse)
	return result, response, err
}

// UpdateTag updateTag
/* Updates a tag specified by id
 */
func (s *TagService) UpdateTag(updateTagRequest *UpdateTagRequest) (*UpdateTagResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag"

	response, err := RestyClient.R().
		SetBody(updateTagRequest).
		SetResult(&UpdateTagResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateTag")
	}

	result := response.Result().(*UpdateTagResponse)
	return result, response, err
}

// UpdatesTagMembership updatesTagMembership
/* Updates tag membership. As part of the request payload through this API, only the specified members are added / retained to the given input tags. Possible values of memberType attribute in the request payload can be queried by using the /tag/member/type API
 */
func (s *TagService) UpdatesTagMembership(updatesTagMembershipRequest *UpdatesTagMembershipRequest) (*UpdatesTagMembershipResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/tag/member"

	response, err := RestyClient.R().
		SetBody(updatesTagMembershipRequest).
		SetResult(&UpdatesTagMembershipResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updatesTagMembership")
	}

	result := response.Result().(*UpdatesTagMembershipResponse)
	return result, response, err
}
