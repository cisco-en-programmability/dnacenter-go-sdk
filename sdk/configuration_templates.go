package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ConfigurationTemplatesService is the service to communicate with the ConfigurationTemplates API endpoint
type ConfigurationTemplatesService service

// CreateProjectRequest is the CreateProjectRequest definition
type CreateProjectRequest struct {
	CreateTime     int      `json:"createTime,omitempty"`     //
	Description    string   `json:"description,omitempty"`    //
	ID             string   `json:"id,omitempty"`             //
	LastUpdateTime int      `json:"lastUpdateTime,omitempty"` //
	Name           string   `json:"name,omitempty"`           //
	Tags           []string `json:"tags,omitempty"`           //
	Templates      string   `json:"templates,omitempty"`      //
}

// CreateTemplateRequest is the CreateTemplateRequest definition
type CreateTemplateRequest struct {
	Author              string `json:"author,omitempty"`    //
	Composite           bool   `json:"composite,omitempty"` //
	ContainingTemplates []struct {
		Composite bool   `json:"composite,omitempty"` //
		ID        string `json:"id,omitempty"`        //
		Name      string `json:"name,omitempty"`      //
		Version   string `json:"version,omitempty"`   //
	} `json:"containingTemplates,omitempty"` //
	CreateTime  int    `json:"createTime,omitempty"`  //
	Description string `json:"description,omitempty"` //
	DeviceTypes []struct {
		ProductFamily string `json:"productFamily,omitempty"` //
		ProductSeries string `json:"productSeries,omitempty"` //
		ProductType   string `json:"productType,omitempty"`   //
	} `json:"deviceTypes,omitempty"` //
	FailurePolicy           string `json:"failurePolicy,omitempty"`           //
	ID                      string `json:"id,omitempty"`                      //
	LastUpdateTime          int    `json:"lastUpdateTime,omitempty"`          //
	Name                    string `json:"name,omitempty"`                    //
	ParentTemplateID        string `json:"parentTemplateId,omitempty"`        //
	ProjectID               string `json:"projectId,omitempty"`               //
	ProjectName             string `json:"projectName,omitempty"`             //
	RollbackTemplateContent string `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []struct {
		Binding         string `json:"binding,omitempty"`         //
		DataType        string `json:"dataType,omitempty"`        //
		DefaultValue    string `json:"defaultValue,omitempty"`    //
		Description     string `json:"description,omitempty"`     //
		DisplayName     string `json:"displayName,omitempty"`     //
		Group           string `json:"group,omitempty"`           //
		ID              string `json:"id,omitempty"`              //
		InstructionText string `json:"instructionText,omitempty"` //
		Key             string `json:"key,omitempty"`             //
		NotParam        bool   `json:"notParam,omitempty"`        //
		Order           int    `json:"order,omitempty"`           //
		ParamArray      bool   `json:"paramArray,omitempty"`      //
		ParameterName   string `json:"parameterName,omitempty"`   //
		Provider        string `json:"provider,omitempty"`        //
		Range           []struct {
			ID       string `json:"id,omitempty"`       //
			MaxValue int    `json:"maxValue,omitempty"` //
			MinValue int    `json:"minValue,omitempty"` //
		} `json:"range,omitempty"` //
		Required  bool `json:"required,omitempty"` //
		Selection struct {
			ID              string `json:"id,omitempty"`              //
			SelectionType   string `json:"selectionType,omitempty"`   //
			SelectionValues string `json:"selectionValues,omitempty"` //
		} `json:"selection,omitempty"` //
	} `json:"rollbackTemplateParams,omitempty"` //
	SoftwareType    string   `json:"softwareType,omitempty"`    //
	SoftwareVariant string   `json:"softwareVariant,omitempty"` //
	SoftwareVersion string   `json:"softwareVersion,omitempty"` //
	Tags            []string `json:"tags,omitempty"`            //
	TemplateContent string   `json:"templateContent,omitempty"` //
	TemplateParams  []struct {
		Binding         string `json:"binding,omitempty"`         //
		DataType        string `json:"dataType,omitempty"`        //
		DefaultValue    string `json:"defaultValue,omitempty"`    //
		Description     string `json:"description,omitempty"`     //
		DisplayName     string `json:"displayName,omitempty"`     //
		Group           string `json:"group,omitempty"`           //
		ID              string `json:"id,omitempty"`              //
		InstructionText string `json:"instructionText,omitempty"` //
		Key             string `json:"key,omitempty"`             //
		NotParam        bool   `json:"notParam,omitempty"`        //
		Order           int    `json:"order,omitempty"`           //
		ParamArray      bool   `json:"paramArray,omitempty"`      //
		ParameterName   string `json:"parameterName,omitempty"`   //
		Provider        string `json:"provider,omitempty"`        //
		Range           []struct {
			ID       string `json:"id,omitempty"`       //
			MaxValue int    `json:"maxValue,omitempty"` //
			MinValue int    `json:"minValue,omitempty"` //
		} `json:"range,omitempty"` //
		Required  bool `json:"required,omitempty"` //
		Selection struct {
			ID              string `json:"id,omitempty"`              //
			SelectionType   string `json:"selectionType,omitempty"`   //
			SelectionValues string `json:"selectionValues,omitempty"` //
		} `json:"selection,omitempty"` //
	} `json:"templateParams,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeployTemplateRequest is the DeployTemplateRequest definition
type DeployTemplateRequest struct {
	ForcePushTemplate            bool     `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  bool     `json:"isComposite,omitempty"`                  //
	MainTemplateID               string   `json:"mainTemplateId,omitempty"`               //
	MemberTemplateDeploymentInfo []string `json:"memberTemplateDeploymentInfo,omitempty"` //
	TargetInfo                   []struct {
		HostName string `json:"hostName,omitempty"` //
		ID       string `json:"id,omitempty"`       //
		Params   string `json:"params,omitempty"`   //
		Type     string `json:"type,omitempty"`     //
	} `json:"targetInfo,omitempty"` //
	TemplateID string `json:"templateId,omitempty"` //
}

// PreviewTemplateRequest is the PreviewTemplateRequest definition
type PreviewTemplateRequest struct {
	Params     string `json:"params,omitempty"`     //
	TemplateID string `json:"templateId,omitempty"` //
}

// UpdateProjectRequest is the UpdateProjectRequest definition
type UpdateProjectRequest struct {
	CreateTime     int      `json:"createTime,omitempty"`     //
	Description    string   `json:"description,omitempty"`    //
	ID             string   `json:"id,omitempty"`             //
	LastUpdateTime int      `json:"lastUpdateTime,omitempty"` //
	Name           string   `json:"name,omitempty"`           //
	Tags           []string `json:"tags,omitempty"`           //
	Templates      string   `json:"templates,omitempty"`      //
}

// UpdateTemplateRequest is the UpdateTemplateRequest definition
type UpdateTemplateRequest struct {
	Author              string `json:"author,omitempty"`    //
	Composite           bool   `json:"composite,omitempty"` //
	ContainingTemplates []struct {
		Composite bool   `json:"composite,omitempty"` //
		ID        string `json:"id,omitempty"`        //
		Name      string `json:"name,omitempty"`      //
		Version   string `json:"version,omitempty"`   //
	} `json:"containingTemplates,omitempty"` //
	CreateTime  int    `json:"createTime,omitempty"`  //
	Description string `json:"description,omitempty"` //
	DeviceTypes []struct {
		ProductFamily string `json:"productFamily,omitempty"` //
		ProductSeries string `json:"productSeries,omitempty"` //
		ProductType   string `json:"productType,omitempty"`   //
	} `json:"deviceTypes,omitempty"` //
	FailurePolicy           string `json:"failurePolicy,omitempty"`           //
	ID                      string `json:"id,omitempty"`                      //
	LastUpdateTime          int    `json:"lastUpdateTime,omitempty"`          //
	Name                    string `json:"name,omitempty"`                    //
	ParentTemplateID        string `json:"parentTemplateId,omitempty"`        //
	ProjectID               string `json:"projectId,omitempty"`               //
	ProjectName             string `json:"projectName,omitempty"`             //
	RollbackTemplateContent string `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []struct {
		Binding         string `json:"binding,omitempty"`         //
		DataType        string `json:"dataType,omitempty"`        //
		DefaultValue    string `json:"defaultValue,omitempty"`    //
		Description     string `json:"description,omitempty"`     //
		DisplayName     string `json:"displayName,omitempty"`     //
		Group           string `json:"group,omitempty"`           //
		ID              string `json:"id,omitempty"`              //
		InstructionText string `json:"instructionText,omitempty"` //
		Key             string `json:"key,omitempty"`             //
		NotParam        bool   `json:"notParam,omitempty"`        //
		Order           int    `json:"order,omitempty"`           //
		ParamArray      bool   `json:"paramArray,omitempty"`      //
		ParameterName   string `json:"parameterName,omitempty"`   //
		Provider        string `json:"provider,omitempty"`        //
		Range           []struct {
			ID       string `json:"id,omitempty"`       //
			MaxValue int    `json:"maxValue,omitempty"` //
			MinValue int    `json:"minValue,omitempty"` //
		} `json:"range,omitempty"` //
		Required  bool `json:"required,omitempty"` //
		Selection struct {
			ID              string `json:"id,omitempty"`              //
			SelectionType   string `json:"selectionType,omitempty"`   //
			SelectionValues string `json:"selectionValues,omitempty"` //
		} `json:"selection,omitempty"` //
	} `json:"rollbackTemplateParams,omitempty"` //
	SoftwareType    string   `json:"softwareType,omitempty"`    //
	SoftwareVariant string   `json:"softwareVariant,omitempty"` //
	SoftwareVersion string   `json:"softwareVersion,omitempty"` //
	Tags            []string `json:"tags,omitempty"`            //
	TemplateContent string   `json:"templateContent,omitempty"` //
	TemplateParams  []struct {
		Binding         string `json:"binding,omitempty"`         //
		DataType        string `json:"dataType,omitempty"`        //
		DefaultValue    string `json:"defaultValue,omitempty"`    //
		Description     string `json:"description,omitempty"`     //
		DisplayName     string `json:"displayName,omitempty"`     //
		Group           string `json:"group,omitempty"`           //
		ID              string `json:"id,omitempty"`              //
		InstructionText string `json:"instructionText,omitempty"` //
		Key             string `json:"key,omitempty"`             //
		NotParam        bool   `json:"notParam,omitempty"`        //
		Order           int    `json:"order,omitempty"`           //
		ParamArray      bool   `json:"paramArray,omitempty"`      //
		ParameterName   string `json:"parameterName,omitempty"`   //
		Provider        string `json:"provider,omitempty"`        //
		Range           []struct {
			ID       string `json:"id,omitempty"`       //
			MaxValue int    `json:"maxValue,omitempty"` //
			MinValue int    `json:"minValue,omitempty"` //
		} `json:"range,omitempty"` //
		Required  bool `json:"required,omitempty"` //
		Selection struct {
			ID              string `json:"id,omitempty"`              //
			SelectionType   string `json:"selectionType,omitempty"`   //
			SelectionValues string `json:"selectionValues,omitempty"` //
		} `json:"selection,omitempty"` //
	} `json:"templateParams,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// VersionTemplateRequest is the VersionTemplateRequest definition
type VersionTemplateRequest struct {
	Comments   string `json:"comments,omitempty"`   //
	TemplateID string `json:"templateId,omitempty"` //
}

// CreateProjectResponse is the CreateProjectResponse definition
type CreateProjectResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateTemplateResponse is the CreateTemplateResponse definition
type CreateTemplateResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteProjectResponse is the DeleteProjectResponse definition
type DeleteProjectResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeleteTemplateResponse is the DeleteTemplateResponse definition
type DeleteTemplateResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DeployTemplateResponse is the DeployTemplateResponse definition
type DeployTemplateResponse struct {
	DeploymentID   string `json:"deploymentId,omitempty"`   //
	DeploymentName string `json:"deploymentName,omitempty"` //
	Devices        []struct {
		DeviceID  string `json:"deviceId,omitempty"`  //
		Duration  string `json:"duration,omitempty"`  //
		EndTime   string `json:"endTime,omitempty"`   //
		IPAddress string `json:"ipAddress,omitempty"` //
		Name      string `json:"name,omitempty"`      //
		StartTime string `json:"startTime,omitempty"` //
		Status    string `json:"status,omitempty"`    //
	} `json:"devices,omitempty"` //
	Duration        string `json:"duration,omitempty"`        //
	EndTime         string `json:"endTime,omitempty"`         //
	ProjectName     string `json:"projectName,omitempty"`     //
	StartTime       string `json:"startTime,omitempty"`       //
	Status          string `json:"status,omitempty"`          //
	TemplateName    string `json:"templateName,omitempty"`    //
	TemplateVersion string `json:"templateVersion,omitempty"` //
}

// GetProjectsResponse is the GetProjectsResponse definition
type GetProjectsResponse struct {
	ID        string `json:"id,omitempty"`   //
	Name      string `json:"name,omitempty"` //
	Templates []struct {
		Composite bool   `json:"composite,omitempty"` //
		ID        string `json:"id,omitempty"`        //
		Name      string `json:"name,omitempty"`      //
	} `json:"templates,omitempty"` //
}

// GetTemplateDeploymentStatusResponse is the GetTemplateDeploymentStatusResponse definition
type GetTemplateDeploymentStatusResponse struct {
	DeploymentID   string `json:"deploymentId,omitempty"`   //
	DeploymentName string `json:"deploymentName,omitempty"` //
	Devices        []struct {
		DeviceID  string `json:"deviceId,omitempty"`  //
		Duration  string `json:"duration,omitempty"`  //
		EndTime   string `json:"endTime,omitempty"`   //
		IPAddress string `json:"ipAddress,omitempty"` //
		Name      string `json:"name,omitempty"`      //
		StartTime string `json:"startTime,omitempty"` //
		Status    string `json:"status,omitempty"`    //
	} `json:"devices,omitempty"` //
	Duration        string `json:"duration,omitempty"`        //
	EndTime         string `json:"endTime,omitempty"`         //
	ProjectName     string `json:"projectName,omitempty"`     //
	StartTime       string `json:"startTime,omitempty"`       //
	Status          string `json:"status,omitempty"`          //
	TemplateName    string `json:"templateName,omitempty"`    //
	TemplateVersion string `json:"templateVersion,omitempty"` //
}

// GetTemplateDetailsResponse is the GetTemplateDetailsResponse definition
type GetTemplateDetailsResponse struct {
	Author              string `json:"author,omitempty"`    //
	Composite           bool   `json:"composite,omitempty"` //
	ContainingTemplates []struct {
		Composite bool   `json:"composite,omitempty"` //
		ID        string `json:"id,omitempty"`        //
		Name      string `json:"name,omitempty"`      //
		Version   string `json:"version,omitempty"`   //
	} `json:"containingTemplates,omitempty"` //
	CreateTime  int    `json:"createTime,omitempty"`  //
	Description string `json:"description,omitempty"` //
	DeviceTypes []struct {
		ProductFamily string `json:"productFamily,omitempty"` //
		ProductSeries string `json:"productSeries,omitempty"` //
		ProductType   string `json:"productType,omitempty"`   //
	} `json:"deviceTypes,omitempty"` //
	FailurePolicy           string `json:"failurePolicy,omitempty"`           //
	ID                      string `json:"id,omitempty"`                      //
	LastUpdateTime          int    `json:"lastUpdateTime,omitempty"`          //
	Name                    string `json:"name,omitempty"`                    //
	ParentTemplateID        string `json:"parentTemplateId,omitempty"`        //
	ProjectID               string `json:"projectId,omitempty"`               //
	ProjectName             string `json:"projectName,omitempty"`             //
	RollbackTemplateContent string `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []struct {
		Binding         string `json:"binding,omitempty"`         //
		DataType        string `json:"dataType,omitempty"`        //
		DefaultValue    string `json:"defaultValue,omitempty"`    //
		Description     string `json:"description,omitempty"`     //
		DisplayName     string `json:"displayName,omitempty"`     //
		Group           string `json:"group,omitempty"`           //
		ID              string `json:"id,omitempty"`              //
		InstructionText string `json:"instructionText,omitempty"` //
		Key             string `json:"key,omitempty"`             //
		NotParam        bool   `json:"notParam,omitempty"`        //
		Order           int    `json:"order,omitempty"`           //
		ParamArray      bool   `json:"paramArray,omitempty"`      //
		ParameterName   string `json:"parameterName,omitempty"`   //
		Provider        string `json:"provider,omitempty"`        //
		Range           []struct {
			ID       string `json:"id,omitempty"`       //
			MaxValue int    `json:"maxValue,omitempty"` //
			MinValue int    `json:"minValue,omitempty"` //
		} `json:"range,omitempty"` //
		Required  bool `json:"required,omitempty"` //
		Selection struct {
			ID              string `json:"id,omitempty"`              //
			SelectionType   string `json:"selectionType,omitempty"`   //
			SelectionValues string `json:"selectionValues,omitempty"` //
		} `json:"selection,omitempty"` //
	} `json:"rollbackTemplateParams,omitempty"` //
	SoftwareType    string   `json:"softwareType,omitempty"`    //
	SoftwareVariant string   `json:"softwareVariant,omitempty"` //
	SoftwareVersion string   `json:"softwareVersion,omitempty"` //
	Tags            []string `json:"tags,omitempty"`            //
	TemplateContent string   `json:"templateContent,omitempty"` //
	TemplateParams  []struct {
		Binding         string `json:"binding,omitempty"`         //
		DataType        string `json:"dataType,omitempty"`        //
		DefaultValue    string `json:"defaultValue,omitempty"`    //
		Description     string `json:"description,omitempty"`     //
		DisplayName     string `json:"displayName,omitempty"`     //
		Group           string `json:"group,omitempty"`           //
		ID              string `json:"id,omitempty"`              //
		InstructionText string `json:"instructionText,omitempty"` //
		Key             string `json:"key,omitempty"`             //
		NotParam        bool   `json:"notParam,omitempty"`        //
		Order           int    `json:"order,omitempty"`           //
		ParamArray      bool   `json:"paramArray,omitempty"`      //
		ParameterName   string `json:"parameterName,omitempty"`   //
		Provider        string `json:"provider,omitempty"`        //
		Range           []struct {
			ID       string `json:"id,omitempty"`       //
			MaxValue int    `json:"maxValue,omitempty"` //
			MinValue int    `json:"minValue,omitempty"` //
		} `json:"range,omitempty"` //
		Required  bool `json:"required,omitempty"` //
		Selection struct {
			ID              string `json:"id,omitempty"`              //
			SelectionType   string `json:"selectionType,omitempty"`   //
			SelectionValues string `json:"selectionValues,omitempty"` //
		} `json:"selection,omitempty"` //
	} `json:"templateParams,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// GetTemplateVersionsResponse is the GetTemplateVersionsResponse definition
type GetTemplateVersionsResponse struct {
	Composite    bool   `json:"composite,omitempty"`   //
	Name         string `json:"name,omitempty"`        //
	ProjectID    string `json:"projectId,omitempty"`   //
	ProjectName  string `json:"projectName,omitempty"` //
	TemplateID   string `json:"templateId,omitempty"`  //
	VersionsInfo []struct {
		Description string `json:"description,omitempty"` //
		ID          string `json:"id,omitempty"`          //
		VersionTime int    `json:"versionTime,omitempty"` //
	} `json:"versionsInfo,omitempty"` //
}

// GetsTheTemplatesAvailableResponse is the GetsTheTemplatesAvailableResponse definition
type GetsTheTemplatesAvailableResponse struct {
	Composite    bool   `json:"composite,omitempty"`   //
	Name         string `json:"name,omitempty"`        //
	ProjectID    string `json:"projectId,omitempty"`   //
	ProjectName  string `json:"projectName,omitempty"` //
	TemplateID   string `json:"templateId,omitempty"`  //
	VersionsInfo []struct {
		Description string `json:"description,omitempty"` //
		ID          string `json:"id,omitempty"`          //
		VersionTime int    `json:"versionTime,omitempty"` //
	} `json:"versionsInfo,omitempty"` //
}

// PreviewTemplateResponse is the PreviewTemplateResponse definition
type PreviewTemplateResponse struct {
	CLIPreview       string `json:"cliPreview,omitempty"`       //
	TemplateID       string `json:"templateId,omitempty"`       //
	ValidationErrors string `json:"validationErrors,omitempty"` //
}

// UpdateProjectResponse is the UpdateProjectResponse definition
type UpdateProjectResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// UpdateTemplateResponse is the UpdateTemplateResponse definition
type UpdateTemplateResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// VersionTemplateResponse is the VersionTemplateResponse definition
type VersionTemplateResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// CreateProject createProject
/* Creates a new project
 */
func (s *ConfigurationTemplatesService) CreateProject(createProjectRequest *CreateProjectRequest) (*CreateProjectResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := RestyClient.R().
		SetBody(createProjectRequest).
		SetResult(&CreateProjectResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateProjectResponse)
	return result, response, err
}

// CreateTemplate createTemplate
/* Creates a new template
@param projectID projectId
*/
func (s *ConfigurationTemplatesService) CreateTemplate(projectID string, createTemplateRequest *CreateTemplateRequest) (*CreateTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project/{projectId}/template"
	path = strings.Replace(path, "{"+"projectId"+"}", fmt.Sprintf("%v", projectID), -1)

	response, err := RestyClient.R().
		SetBody(createTemplateRequest).
		SetResult(&CreateTemplateResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*CreateTemplateResponse)
	return result, response, err
}

// DeleteProject deleteProject
/* Deletes an existing Project
@param projectID projectId
*/
func (s *ConfigurationTemplatesService) DeleteProject(projectID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{"+"projectId"+"}", fmt.Sprintf("%v", projectID), -1)

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
@param templateID templateId
*/
func (s *ConfigurationTemplatesService) DeleteTemplate(templateID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateID), -1)

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
func (s *ConfigurationTemplatesService) DeployTemplate(deployTemplateRequest *DeployTemplateRequest) (*DeployTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/deploy"

	response, err := RestyClient.R().
		SetBody(deployTemplateRequest).
		SetResult(&DeployTemplateResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*DeployTemplateResponse)
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
@param deploymentID deploymentId
*/
func (s *ConfigurationTemplatesService) GetTemplateDeploymentStatus(deploymentID string) (*GetTemplateDeploymentStatusResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/deploy/status/{deploymentId}"
	path = strings.Replace(path, "{"+"deploymentId"+"}", fmt.Sprintf("%v", deploymentID), -1)

	response, err := RestyClient.R().
		SetResult(&GetTemplateDeploymentStatusResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetTemplateDeploymentStatusResponse)
	return result, response, err
}

// GetTemplateDetailsQueryParams defines the query parameters for this request
type GetTemplateDetailsQueryParams struct {
	LatestVersion bool `url:"latestVersion,omitempty"` // latestVersion
}

// GetTemplateDetails getTemplateDetails
/* Returns details of the specified template
@param templateID templateId
@param latestVersion latestVersion
*/
func (s *ConfigurationTemplatesService) GetTemplateDetails(templateID string, getTemplateDetailsQueryParams *GetTemplateDetailsQueryParams) (*GetTemplateDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateID), -1)

	queryString, _ := query.Values(getTemplateDetailsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&GetTemplateDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetTemplateDetailsResponse)
	return result, response, err
}

// GetTemplateVersions getTemplateVersions
/* Returns the versions of a specified template
@param templateID templateId
*/
func (s *ConfigurationTemplatesService) GetTemplateVersions(templateID string) (*GetTemplateVersionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/version/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateID), -1)

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
	ProjectID                  string `url:"projectId,omitempty"`                  // projectId
	SoftwareType               string `url:"softwareType,omitempty"`               // softwareType
	SoftwareVersion            string `url:"softwareVersion,omitempty"`            // softwareVersion
	ProductFamily              string `url:"productFamily,omitempty"`              // productFamily
	ProductSeries              string `url:"productSeries,omitempty"`              // productSeries
	ProductType                string `url:"productType,omitempty"`                // productType
	FilterConflictingTemplates bool   `url:"filterConflictingTemplates,omitempty"` // filterConflictingTemplates
}

// GetsTheTemplatesAvailable getsTheTemplatesAvailable
/* List the templates available
@param projectID projectId
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
func (s *ConfigurationTemplatesService) PreviewTemplate(previewTemplateRequest *PreviewTemplateRequest) (*PreviewTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/preview"

	response, err := RestyClient.R().
		SetBody(previewTemplateRequest).
		SetResult(&PreviewTemplateResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*PreviewTemplateResponse)
	return result, response, err
}

// UpdateProject updateProject
/* Updates an existing project
 */
func (s *ConfigurationTemplatesService) UpdateProject(updateProjectRequest *UpdateProjectRequest) (*UpdateProjectResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := RestyClient.R().
		SetBody(updateProjectRequest).
		SetResult(&UpdateProjectResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateProjectResponse)
	return result, response, err
}

// UpdateTemplate updateTemplate
/* Updates an existing template
 */
func (s *ConfigurationTemplatesService) UpdateTemplate(updateTemplateRequest *UpdateTemplateRequest) (*UpdateTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template"

	response, err := RestyClient.R().
		SetBody(updateTemplateRequest).
		SetResult(&UpdateTemplateResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*UpdateTemplateResponse)
	return result, response, err
}

// VersionTemplate versionTemplate
/* Creates Versioning for the current contents of the template
 */
func (s *ConfigurationTemplatesService) VersionTemplate(versionTemplateRequest *VersionTemplateRequest) (*VersionTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/version"

	response, err := RestyClient.R().
		SetBody(versionTemplateRequest).
		SetResult(&VersionTemplateResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*VersionTemplateResponse)
	return result, response, err
}
