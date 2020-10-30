package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ConfigurationTemplatesService is the service to communicate with the ConfigurationTemplates API endpoint
type ConfigurationTemplatesService service

// ContainingTemplates is the ContainingTemplates definition
type ContainingTemplates struct {
	Composite bool   `json:"composite,omitempty"` //
	Id        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      //
	Version   string `json:"version,omitempty"`   //
}

// DeviceTypes is the DeviceTypes definition
type DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` //
	ProductSeries string `json:"productSeries,omitempty"` //
	ProductType   string `json:"productType,omitempty"`   //
}

// ProjectDTO is the ProjectDTO definition
type ProjectDTO struct {
	CreateTime     int      `json:"createTime,omitempty"`     //
	Description    string   `json:"description,omitempty"`    //
	Id             string   `json:"id,omitempty"`             //
	LastUpdateTime int      `json:"lastUpdateTime,omitempty"` //
	Name           string   `json:"name,omitempty"`           //
	Tags           []string `json:"tags,omitempty"`           //
	Templates      string   `json:"templates,omitempty"`      //
}

// Range is the Range definition
type Range struct {
	Id       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// RollbackTemplateParams is the RollbackTemplateParams definition
type RollbackTemplateParams struct {
	Binding         string    `json:"binding,omitempty"`         //
	DataType        string    `json:"dataType,omitempty"`        //
	DefaultValue    string    `json:"defaultValue,omitempty"`    //
	Description     string    `json:"description,omitempty"`     //
	DisplayName     string    `json:"displayName,omitempty"`     //
	Group           string    `json:"group,omitempty"`           //
	Id              string    `json:"id,omitempty"`              //
	InstructionText string    `json:"instructionText,omitempty"` //
	Key             string    `json:"key,omitempty"`             //
	NotParam        bool      `json:"notParam,omitempty"`        //
	Order           int       `json:"order,omitempty"`           //
	ParamArray      bool      `json:"paramArray,omitempty"`      //
	ParameterName   string    `json:"parameterName,omitempty"`   //
	Provider        string    `json:"provider,omitempty"`        //
	Range           []Range   `json:"range,omitempty"`           //
	Required        bool      `json:"required,omitempty"`        //
	Selection       Selection `json:"selection,omitempty"`       //
}

// Selection is the Selection definition
type Selection struct {
	Id              string `json:"id,omitempty"`              //
	SelectionType   string `json:"selectionType,omitempty"`   //
	SelectionValues string `json:"selectionValues,omitempty"` //
}

// TargetInfo is the TargetInfo definition
type TargetInfo struct {
	HostName string `json:"hostName,omitempty"` //
	Id       string `json:"id,omitempty"`       //
	Params   string `json:"params,omitempty"`   //
	Type     string `json:"type,omitempty"`     //
}

// TemplateDTO is the TemplateDTO definition
type TemplateDTO struct {
	Author                  string                   `json:"author,omitempty"`                  //
	Composite               bool                     `json:"composite,omitempty"`               //
	ContainingTemplates     []ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              int                      `json:"createTime,omitempty"`              //
	Description             string                   `json:"description,omitempty"`             //
	DeviceTypes             []DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                   `json:"failurePolicy,omitempty"`           //
	Id                      string                   `json:"id,omitempty"`                      //
	LastUpdateTime          int                      `json:"lastUpdateTime,omitempty"`          //
	Name                    string                   `json:"name,omitempty"`                    //
	ParentTemplateId        string                   `json:"parentTemplateId,omitempty"`        //
	ProjectId               string                   `json:"projectId,omitempty"`               //
	ProjectName             string                   `json:"projectName,omitempty"`             //
	RollbackTemplateContent string                   `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                   `json:"softwareType,omitempty"`            //
	SoftwareVariant         string                   `json:"softwareVariant,omitempty"`         //
	SoftwareVersion         string                   `json:"softwareVersion,omitempty"`         //
	Tags                    []string                 `json:"tags,omitempty"`                    //
	TemplateContent         string                   `json:"templateContent,omitempty"`         //
	TemplateParams          []TemplateParams         `json:"templateParams,omitempty"`          //
	Version                 string                   `json:"version,omitempty"`                 //
}

// TemplateDeploymentInfo is the TemplateDeploymentInfo definition
type TemplateDeploymentInfo struct {
	ForcePushTemplate            bool         `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  bool         `json:"isComposite,omitempty"`                  //
	MainTemplateId               string       `json:"mainTemplateId,omitempty"`               //
	MemberTemplateDeploymentInfo []string     `json:"memberTemplateDeploymentInfo,omitempty"` //
	TargetInfo                   []TargetInfo `json:"targetInfo,omitempty"`                   //
	TemplateId                   string       `json:"templateId,omitempty"`                   //
}

// TemplateParams is the TemplateParams definition
type TemplateParams struct {
	Binding         string    `json:"binding,omitempty"`         //
	DataType        string    `json:"dataType,omitempty"`        //
	DefaultValue    string    `json:"defaultValue,omitempty"`    //
	Description     string    `json:"description,omitempty"`     //
	DisplayName     string    `json:"displayName,omitempty"`     //
	Group           string    `json:"group,omitempty"`           //
	Id              string    `json:"id,omitempty"`              //
	InstructionText string    `json:"instructionText,omitempty"` //
	Key             string    `json:"key,omitempty"`             //
	NotParam        bool      `json:"notParam,omitempty"`        //
	Order           int       `json:"order,omitempty"`           //
	ParamArray      bool      `json:"paramArray,omitempty"`      //
	ParameterName   string    `json:"parameterName,omitempty"`   //
	Provider        string    `json:"provider,omitempty"`        //
	Range           []Range   `json:"range,omitempty"`           //
	Required        bool      `json:"required,omitempty"`        //
	Selection       Selection `json:"selection,omitempty"`       //
}

// TemplatePreviewRequestDTO is the TemplatePreviewRequestDTO definition
type TemplatePreviewRequestDTO struct {
	Params     string `json:"params,omitempty"`     //
	TemplateId string `json:"templateId,omitempty"` //
}

// TemplateVersionRequestDTO is the TemplateVersionRequestDTO definition
type TemplateVersionRequestDTO struct {
	Comments   string `json:"comments,omitempty"`   //
	TemplateId string `json:"templateId,omitempty"` //
}

// CollectionProjectDTO is the CollectionProjectDTO definition
type CollectionProjectDTO struct {
	Id        string      `json:"id,omitempty"`        //
	Name      string      `json:"name,omitempty"`      //
	Templates []Templates `json:"templates,omitempty"` //
}

// CollectionTemplateInfo is the CollectionTemplateInfo definition
type CollectionTemplateInfo struct {
	Composite    bool           `json:"composite,omitempty"`    //
	Name         string         `json:"name,omitempty"`         //
	ProjectId    string         `json:"projectId,omitempty"`    //
	ProjectName  string         `json:"projectName,omitempty"`  //
	TemplateId   string         `json:"templateId,omitempty"`   //
	VersionsInfo []VersionsInfo `json:"versionsInfo,omitempty"` //
}

// ContainingTemplates is the ContainingTemplates definition
type ContainingTemplates struct {
	Composite bool   `json:"composite,omitempty"` //
	Id        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      //
	Version   string `json:"version,omitempty"`   //
}

// DeviceTypes is the DeviceTypes definition
type DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` //
	ProductSeries string `json:"productSeries,omitempty"` //
	ProductType   string `json:"productType,omitempty"`   //
}

// Devices is the Devices definition
type Devices struct {
	DeviceId  string `json:"deviceId,omitempty"`  //
	Duration  string `json:"duration,omitempty"`  //
	EndTime   string `json:"endTime,omitempty"`   //
	IpAddress string `json:"ipAddress,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	StartTime string `json:"startTime,omitempty"` //
	Status    string `json:"status,omitempty"`    //
}

// Range is the Range definition
type Range struct {
	Id       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// Response is the Response definition
type Response struct {
	TaskId string `json:"taskId,omitempty"` //
	Url    string `json:"url,omitempty"`    //
}

// RollbackTemplateParams is the RollbackTemplateParams definition
type RollbackTemplateParams struct {
	Binding         string    `json:"binding,omitempty"`         //
	DataType        string    `json:"dataType,omitempty"`        //
	DefaultValue    string    `json:"defaultValue,omitempty"`    //
	Description     string    `json:"description,omitempty"`     //
	DisplayName     string    `json:"displayName,omitempty"`     //
	Group           string    `json:"group,omitempty"`           //
	Id              string    `json:"id,omitempty"`              //
	InstructionText string    `json:"instructionText,omitempty"` //
	Key             string    `json:"key,omitempty"`             //
	NotParam        bool      `json:"notParam,omitempty"`        //
	Order           int       `json:"order,omitempty"`           //
	ParamArray      bool      `json:"paramArray,omitempty"`      //
	ParameterName   string    `json:"parameterName,omitempty"`   //
	Provider        string    `json:"provider,omitempty"`        //
	Range           []Range   `json:"range,omitempty"`           //
	Required        bool      `json:"required,omitempty"`        //
	Selection       Selection `json:"selection,omitempty"`       //
}

// Selection is the Selection definition
type Selection struct {
	Id              string `json:"id,omitempty"`              //
	SelectionType   string `json:"selectionType,omitempty"`   //
	SelectionValues string `json:"selectionValues,omitempty"` //
}

// TaskIdResult is the TaskIdResult definition
type TaskIdResult struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// TemplateDTO is the TemplateDTO definition
type TemplateDTO struct {
	Author                  string                   `json:"author,omitempty"`                  //
	Composite               bool                     `json:"composite,omitempty"`               //
	ContainingTemplates     []ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              int                      `json:"createTime,omitempty"`              //
	Description             string                   `json:"description,omitempty"`             //
	DeviceTypes             []DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                   `json:"failurePolicy,omitempty"`           //
	Id                      string                   `json:"id,omitempty"`                      //
	LastUpdateTime          int                      `json:"lastUpdateTime,omitempty"`          //
	Name                    string                   `json:"name,omitempty"`                    //
	ParentTemplateId        string                   `json:"parentTemplateId,omitempty"`        //
	ProjectId               string                   `json:"projectId,omitempty"`               //
	ProjectName             string                   `json:"projectName,omitempty"`             //
	RollbackTemplateContent string                   `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                   `json:"softwareType,omitempty"`            //
	SoftwareVariant         string                   `json:"softwareVariant,omitempty"`         //
	SoftwareVersion         string                   `json:"softwareVersion,omitempty"`         //
	Tags                    []string                 `json:"tags,omitempty"`                    //
	TemplateContent         string                   `json:"templateContent,omitempty"`         //
	TemplateParams          []TemplateParams         `json:"templateParams,omitempty"`          //
	Version                 string                   `json:"version,omitempty"`                 //
}

// TemplateDeploymentStatusDTO is the TemplateDeploymentStatusDTO definition
type TemplateDeploymentStatusDTO struct {
	DeploymentId    string    `json:"deploymentId,omitempty"`    //
	DeploymentName  string    `json:"deploymentName,omitempty"`  //
	Devices         []Devices `json:"devices,omitempty"`         //
	Duration        string    `json:"duration,omitempty"`        //
	EndTime         string    `json:"endTime,omitempty"`         //
	ProjectName     string    `json:"projectName,omitempty"`     //
	StartTime       string    `json:"startTime,omitempty"`       //
	Status          string    `json:"status,omitempty"`          //
	TemplateName    string    `json:"templateName,omitempty"`    //
	TemplateVersion string    `json:"templateVersion,omitempty"` //
}

// TemplateParams is the TemplateParams definition
type TemplateParams struct {
	Binding         string    `json:"binding,omitempty"`         //
	DataType        string    `json:"dataType,omitempty"`        //
	DefaultValue    string    `json:"defaultValue,omitempty"`    //
	Description     string    `json:"description,omitempty"`     //
	DisplayName     string    `json:"displayName,omitempty"`     //
	Group           string    `json:"group,omitempty"`           //
	Id              string    `json:"id,omitempty"`              //
	InstructionText string    `json:"instructionText,omitempty"` //
	Key             string    `json:"key,omitempty"`             //
	NotParam        bool      `json:"notParam,omitempty"`        //
	Order           int       `json:"order,omitempty"`           //
	ParamArray      bool      `json:"paramArray,omitempty"`      //
	ParameterName   string    `json:"parameterName,omitempty"`   //
	Provider        string    `json:"provider,omitempty"`        //
	Range           []Range   `json:"range,omitempty"`           //
	Required        bool      `json:"required,omitempty"`        //
	Selection       Selection `json:"selection,omitempty"`       //
}

// TemplatePreviewResponseDTO is the TemplatePreviewResponseDTO definition
type TemplatePreviewResponseDTO struct {
	CliPreview       string `json:"cliPreview,omitempty"`       //
	TemplateId       string `json:"templateId,omitempty"`       //
	ValidationErrors string `json:"validationErrors,omitempty"` //
}

// Templates is the Templates definition
type Templates struct {
	Composite bool   `json:"composite,omitempty"` //
	Id        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      //
}

// VersionsInfo is the VersionsInfo definition
type VersionsInfo struct {
	Description string `json:"description,omitempty"` //
	Id          string `json:"id,omitempty"`          //
	VersionTime int    `json:"versionTime,omitempty"` //
}

// CreateProject createProject
/* Creates a new project
 */
func (s *ConfigurationTemplatesService) CreateProject(createProjectRequest *CreateProjectRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := RestyClient.R().
		SetBody(createProjectRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// CreateTemplate createTemplate
/* Creates a new template
@param projectId projectId
*/
func (s *ConfigurationTemplatesService) CreateTemplate(projectId string, createTemplateRequest *CreateTemplateRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project/{projectId}/template"
	path = strings.Replace(path, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	response, err := RestyClient.R().
		SetBody(createTemplateRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// DeleteProject deleteProject
/* Deletes an existing Project
@param projectId projectId
*/
func (s *ConfigurationTemplatesService) DeleteProject(projectId string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{"+"projectId"+"}", fmt.Sprintf("%v", projectId), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeleteTemplate deleteTemplate
/* Deletes an existing template
@param templateId templateId
*/
func (s *ConfigurationTemplatesService) DeleteTemplate(templateId string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateId), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DeployTemplate deployTemplate
/* Deploys a template
 */
func (s *ConfigurationTemplatesService) DeployTemplate(deployTemplateRequest *DeployTemplateRequest) (*TemplateDeploymentStatusDTO, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/deploy"

	response, err := RestyClient.R().
		SetBody(deployTemplateRequest).
		SetResult(&TemplateDeploymentStatusDTO{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TemplateDeploymentStatusDTO)
	return result, response, err

}

// GetProjectsQueryParams defines the query parameters for this request
type GetProjectsQueryParams struct {
	Name string `url:"name,omitempty"` // Name of project to be searched
}

// GetProjects getProjects
/* Returns the projects in the system
@param name Name of project to be searched
*/
func (s *ConfigurationTemplatesService) GetProjects(getProjectsQueryParams *GetProjectsQueryParams) (*GetProjectsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project"

	queryString, _ := query.Values(getProjectsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetProjectsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetProjectsResponse)
	return result, response, err

}

// GetTemplateDeploymentStatus getTemplateDeploymentStatus
/* Returns the status of a deployed template.
@param deploymentId deploymentId
*/
func (s *ConfigurationTemplatesService) GetTemplateDeploymentStatus(deploymentId string) (*TemplateDeploymentStatusDTO, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/deploy/status/{deploymentId}"
	path = strings.Replace(path, "{"+"deploymentId"+"}", fmt.Sprintf("%v", deploymentId), -1)

	response, err := RestyClient.R().
		SetResult(&TemplateDeploymentStatusDTO{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TemplateDeploymentStatusDTO)
	return result, response, err

}

// GetTemplateDetailsQueryParams defines the query parameters for this request
type GetTemplateDetailsQueryParams struct {
	LatestVersion boolean `url:"latestVersion,omitempty"` // latestVersion
}

// GetTemplateDetails getTemplateDetails
/* Returns details of the specified template
@param templateId templateId
@param latestVersion latestVersion
*/
func (s *ConfigurationTemplatesService) GetTemplateDetails(templateId string, getTemplateDetailsQueryParams *GetTemplateDetailsQueryParams) (*TemplateDTO, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateId), -1)

	queryString, _ := query.Values(getTemplateDetailsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&TemplateDTO{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TemplateDTO)
	return result, response, err

}

// GetTemplateVersions getTemplateVersions
/* Returns the versions of a specified template
@param templateId templateId
*/
func (s *ConfigurationTemplatesService) GetTemplateVersions(templateId string) (*GetTemplateVersionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/version/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateId), -1)

	response, err := RestyClient.R().
		SetResult(&GetTemplateVersionsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetTemplateVersionsResponse)
	return result, response, err

}

// GetsTheTemplatesAvailableQueryParams defines the query parameters for this request
type GetsTheTemplatesAvailableQueryParams struct {
	ProjectId                  string  `url:"projectId,omitempty"`                  // projectId
	SoftwareType               string  `url:"softwareType,omitempty"`               // softwareType
	SoftwareVersion            string  `url:"softwareVersion,omitempty"`            // softwareVersion
	ProductFamily              string  `url:"productFamily,omitempty"`              // productFamily
	ProductSeries              string  `url:"productSeries,omitempty"`              // productSeries
	ProductType                string  `url:"productType,omitempty"`                // productType
	FilterConflictingTemplates boolean `url:"filterConflictingTemplates,omitempty"` // filterConflictingTemplates
}

// GetsTheTemplatesAvailable getsTheTemplatesAvailable
/* List the templates available
@param projectId projectId
@param softwareType softwareType
@param softwareVersion softwareVersion
@param productFamily productFamily
@param productSeries productSeries
@param productType productType
@param filterConflictingTemplates filterConflictingTemplates
*/
func (s *ConfigurationTemplatesService) GetsTheTemplatesAvailable(getsTheTemplatesAvailableQueryParams *GetsTheTemplatesAvailableQueryParams) (*GetsTheTemplatesAvailableResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template"

	queryString, _ := query.Values(getsTheTemplatesAvailableQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetsTheTemplatesAvailableResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*GetsTheTemplatesAvailableResponse)
	return result, response, err

}

// PreviewTemplate previewTemplate
/* Previews an existing template
 */
func (s *ConfigurationTemplatesService) PreviewTemplate(previewTemplateRequest *PreviewTemplateRequest) (*TemplatePreviewResponseDTO, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/preview"

	response, err := RestyClient.R().
		SetBody(previewTemplateRequest).
		SetResult(&TemplatePreviewResponseDTO{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TemplatePreviewResponseDTO)
	return result, response, err

}

// UpdateProject updateProject
/* Updates an existing project
 */
func (s *ConfigurationTemplatesService) UpdateProject(updateProjectRequest *UpdateProjectRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := RestyClient.R().
		SetBody(updateProjectRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// UpdateTemplate updateTemplate
/* Updates an existing template
 */
func (s *ConfigurationTemplatesService) UpdateTemplate(updateTemplateRequest *UpdateTemplateRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template"

	response, err := RestyClient.R().
		SetBody(updateTemplateRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}

// VersionTemplate versionTemplate
/* Creates Versioning for the current contents of the template
 */
func (s *ConfigurationTemplatesService) VersionTemplate(versionTemplateRequest *VersionTemplateRequest) (*TaskIdResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/version"

	response, err := RestyClient.R().
		SetBody(versionTemplateRequest).
		SetResult(&TaskIdResult{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*TaskIdResult)
	return result, response, err

}
