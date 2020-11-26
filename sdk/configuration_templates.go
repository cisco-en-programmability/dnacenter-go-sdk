package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ConfigurationTemplatesService is the service to communicate with the ConfigurationTemplates API endpoint
type ConfigurationTemplatesService service

// CreateProjectRequest is the createProjectRequest definition
type CreateProjectRequest struct {
	IsDeletable    bool                            `json:"isDeletable,omitempty"`    //
	CreateTime     int                             `json:"createTime,omitempty"`     //
	Description    string                          `json:"description,omitempty"`    //
	ID             string                          `json:"id,omitempty"`             //
	LastUpdateTime int                             `json:"lastUpdateTime,omitempty"` //
	Name           string                          `json:"name,omitempty"`           //
	Tags           []string                        `json:"tags,omitempty"`           //
	Templates      []CreateProjectRequestTemplates `json:"templates,omitempty"`      //
}

// CreateProjectRequestTemplates is the createProjectRequestTemplates definition
type CreateProjectRequestTemplates struct {
	Name              string  `json:"name,omitempty"`              //
	Composite         bool    `json:"composite,omitempty"`         //
	Language          string  `json:"language,omitempty"`          //
	ID                string  `json:"id,omitempty"`                //
	CustomParamsOrder bool    `json:"customParamsOrder,omitempty"` //
	LastUpdateTime    float64 `json:"lastUpdateTime,omitempty"`    //
	LatestVersionTime float64 `json:"latestVersionTime,omitempty"` //
	ProjectAssociated bool    `json:"projectAssociated,omitempty"` //
	DocumentDatabase  bool    `json:"documentDatabase,omitempty"`  //
}

// CreateProjectRequestTags is the createProjectRequestTags definition
type CreateProjectRequestTags []string

// CreateTemplateRequest is the createTemplateRequest definition
type CreateTemplateRequest struct {
	Author                  string                                        `json:"author,omitempty"`                  //
	Composite               bool                                          `json:"composite,omitempty"`               //
	ContainingTemplates     []CreateTemplateRequestContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              int                                           `json:"createTime,omitempty"`              //
	Description             string                                        `json:"description,omitempty"`             //
	DeviceTypes             []CreateTemplateRequestDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                        `json:"failurePolicy,omitempty"`           //
	ID                      string                                        `json:"id,omitempty"`                      //
	LastUpdateTime          int                                           `json:"lastUpdateTime,omitempty"`          //
	Name                    string                                        `json:"name,omitempty"`                    //
	ParentTemplateID        string                                        `json:"parentTemplateId,omitempty"`        //
	ProjectID               string                                        `json:"projectId,omitempty"`               //
	ProjectName             string                                        `json:"projectName,omitempty"`             //
	RollbackTemplateContent string                                        `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []CreateTemplateRequestRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                        `json:"softwareType,omitempty"`            //
	SoftwareVariant         string                                        `json:"softwareVariant,omitempty"`         //
	SoftwareVersion         string                                        `json:"softwareVersion,omitempty"`         //
	Tags                    []string                                      `json:"tags,omitempty"`                    //
	TemplateContent         string                                        `json:"templateContent,omitempty"`         //
	TemplateParams          []CreateTemplateRequestTemplateParams         `json:"templateParams,omitempty"`          //
	Version                 string                                        `json:"version,omitempty"`                 //
}

// CreateTemplateRequestContainingTemplates is the createTemplateRequestContainingTemplates definition
type CreateTemplateRequestContainingTemplates struct {
	Composite bool   `json:"composite,omitempty"` //
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      //
	Version   string `json:"version,omitempty"`   //
}

// CreateTemplateRequestDeviceTypes is the createTemplateRequestDeviceTypes definition
type CreateTemplateRequestDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` //
	ProductSeries string `json:"productSeries,omitempty"` //
	ProductType   string `json:"productType,omitempty"`   //
}

// CreateTemplateRequestRollbackTemplateParams is the createTemplateRequestRollbackTemplateParams definition
type CreateTemplateRequestRollbackTemplateParams struct {
	Binding         string                                               `json:"binding,omitempty"`         //
	DataType        string                                               `json:"dataType,omitempty"`        //
	DefaultValue    string                                               `json:"defaultValue,omitempty"`    //
	Description     string                                               `json:"description,omitempty"`     //
	DisplayName     string                                               `json:"displayName,omitempty"`     //
	Group           string                                               `json:"group,omitempty"`           //
	ID              string                                               `json:"id,omitempty"`              //
	InstructionText string                                               `json:"instructionText,omitempty"` //
	Key             string                                               `json:"key,omitempty"`             //
	NotParam        bool                                                 `json:"notParam,omitempty"`        //
	Order           int                                                  `json:"order,omitempty"`           //
	ParamArray      bool                                                 `json:"paramArray,omitempty"`      //
	ParameterName   string                                               `json:"parameterName,omitempty"`   //
	Provider        string                                               `json:"provider,omitempty"`        //
	Range           []CreateTemplateRequestRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        bool                                                 `json:"required,omitempty"`        //
	Selection       CreateTemplateRequestRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}

// CreateTemplateRequestRollbackTemplateParamsRange is the createTemplateRequestRollbackTemplateParamsRange definition
type CreateTemplateRequestRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// CreateTemplateRequestRollbackTemplateParamsSelection is the createTemplateRequestRollbackTemplateParamsSelection definition
type CreateTemplateRequestRollbackTemplateParamsSelection struct {
	ID              string                 `json:"id,omitempty"`              //
	SelectionType   string                 `json:"selectionType,omitempty"`   //
	SelectionValues map[string]interface{} `json:"selectionValues,omitempty"` //
}

// CreateTemplateRequestTags is the createTemplateRequestTags definition
type CreateTemplateRequestTags []string

// CreateTemplateRequestTemplateParams is the createTemplateRequestTemplateParams definition
type CreateTemplateRequestTemplateParams struct {
	Binding         string                                       `json:"binding,omitempty"`         //
	DataType        string                                       `json:"dataType,omitempty"`        //
	DefaultValue    string                                       `json:"defaultValue,omitempty"`    //
	Description     string                                       `json:"description,omitempty"`     //
	DisplayName     string                                       `json:"displayName,omitempty"`     //
	Group           string                                       `json:"group,omitempty"`           //
	ID              string                                       `json:"id,omitempty"`              //
	InstructionText string                                       `json:"instructionText,omitempty"` //
	Key             string                                       `json:"key,omitempty"`             //
	NotParam        bool                                         `json:"notParam,omitempty"`        //
	Order           int                                          `json:"order,omitempty"`           //
	ParamArray      bool                                         `json:"paramArray,omitempty"`      //
	ParameterName   string                                       `json:"parameterName,omitempty"`   //
	Provider        string                                       `json:"provider,omitempty"`        //
	Range           []CreateTemplateRequestTemplateParamsRange   `json:"range,omitempty"`           //
	Required        bool                                         `json:"required,omitempty"`        //
	Selection       CreateTemplateRequestTemplateParamsSelection `json:"selection,omitempty"`       //
}

// CreateTemplateRequestTemplateParamsRange is the createTemplateRequestTemplateParamsRange definition
type CreateTemplateRequestTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// CreateTemplateRequestTemplateParamsSelection is the createTemplateRequestTemplateParamsSelection definition
type CreateTemplateRequestTemplateParamsSelection struct {
	ID              string                 `json:"id,omitempty"`              //
	SelectionType   string                 `json:"selectionType,omitempty"`   //
	SelectionValues map[string]interface{} `json:"selectionValues,omitempty"` //
}

// DeployTemplateRequest is the deployTemplateRequest definition
type DeployTemplateRequest struct {
	ForcePushTemplate            bool                              `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  bool                              `json:"isComposite,omitempty"`                  //
	MainTemplateID               string                            `json:"mainTemplateId,omitempty"`               //
	MemberTemplateDeploymentInfo *[]DeployTemplateRequest          `json:"memberTemplateDeploymentInfo,omitempty"` //
	TargetInfo                   []DeployTemplateRequestTargetInfo `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                            `json:"templateId,omitempty"`                   //
}

// DeployTemplateRequestTargetInfo is the deployTemplateRequestTargetInfo definition
type DeployTemplateRequestTargetInfo struct {
	HostName string                 `json:"hostName,omitempty"` //
	ID       string                 `json:"id,omitempty"`       //
	Params   map[string]interface{} `json:"params,omitempty"`   //
	Type     string                 `json:"type,omitempty"`     //
}

// PreviewTemplateRequest is the previewTemplateRequest definition
type PreviewTemplateRequest struct {
	Params     map[string]interface{} `json:"params,omitempty"`     //
	TemplateID string                 `json:"templateId,omitempty"` //
}

// UpdateProjectRequest is the updateProjectRequest definition
type UpdateProjectRequest struct {
	CreateTime     int                             `json:"createTime,omitempty"`     //
	Description    string                          `json:"description,omitempty"`    //
	ID             string                          `json:"id,omitempty"`             //
	LastUpdateTime int                             `json:"lastUpdateTime,omitempty"` //
	Name           string                          `json:"name,omitempty"`           //
	Tags           []string                        `json:"tags,omitempty"`           //
	Templates      []UpdateProjectRequestTemplates `json:"templates,omitempty"`      //
}

// UpdateProjectRequestTemplates is the updateProjectRequestTemplates definition
type UpdateProjectRequestTemplates struct {
	Name              string  `json:"name,omitempty"`              //
	Composite         bool    `json:"composite,omitempty"`         //
	Language          string  `json:"language,omitempty"`          //
	ID                string  `json:"id,omitempty"`                //
	CustomParamsOrder bool    `json:"customParamsOrder,omitempty"` //
	LastUpdateTime    float64 `json:"lastUpdateTime,omitempty"`    //
	LatestVersionTime float64 `json:"latestVersionTime,omitempty"` //
	ProjectAssociated bool    `json:"projectAssociated,omitempty"` //
	DocumentDatabase  bool    `json:"documentDatabase,omitempty"`  //
}

// UpdateProjectRequestTags is the updateProjectRequestTags definition
type UpdateProjectRequestTags []string

// UpdateTemplateRequest is the updateTemplateRequest definition
type UpdateTemplateRequest struct {
	Author                  string                                        `json:"author,omitempty"`                  //
	Composite               bool                                          `json:"composite,omitempty"`               //
	ContainingTemplates     []UpdateTemplateRequestContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              int                                           `json:"createTime,omitempty"`              //
	Description             string                                        `json:"description,omitempty"`             //
	DeviceTypes             []UpdateTemplateRequestDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                        `json:"failurePolicy,omitempty"`           //
	ID                      string                                        `json:"id,omitempty"`                      //
	LastUpdateTime          int                                           `json:"lastUpdateTime,omitempty"`          //
	Name                    string                                        `json:"name,omitempty"`                    //
	ParentTemplateID        string                                        `json:"parentTemplateId,omitempty"`        //
	ProjectID               string                                        `json:"projectId,omitempty"`               //
	ProjectName             string                                        `json:"projectName,omitempty"`             //
	RollbackTemplateContent string                                        `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []UpdateTemplateRequestRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                        `json:"softwareType,omitempty"`            //
	SoftwareVariant         string                                        `json:"softwareVariant,omitempty"`         //
	SoftwareVersion         string                                        `json:"softwareVersion,omitempty"`         //
	Tags                    []string                                      `json:"tags,omitempty"`                    //
	TemplateContent         string                                        `json:"templateContent,omitempty"`         //
	TemplateParams          []UpdateTemplateRequestTemplateParams         `json:"templateParams,omitempty"`          //
	Version                 string                                        `json:"version,omitempty"`                 //
}

// UpdateTemplateRequestContainingTemplates is the updateTemplateRequestContainingTemplates definition
type UpdateTemplateRequestContainingTemplates struct {
	Composite bool   `json:"composite,omitempty"` //
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      //
	Version   string `json:"version,omitempty"`   //
}

// UpdateTemplateRequestDeviceTypes is the updateTemplateRequestDeviceTypes definition
type UpdateTemplateRequestDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` //
	ProductSeries string `json:"productSeries,omitempty"` //
	ProductType   string `json:"productType,omitempty"`   //
}

// UpdateTemplateRequestRollbackTemplateParams is the updateTemplateRequestRollbackTemplateParams definition
type UpdateTemplateRequestRollbackTemplateParams struct {
	Binding         string                                               `json:"binding,omitempty"`         //
	DataType        string                                               `json:"dataType,omitempty"`        //
	DefaultValue    string                                               `json:"defaultValue,omitempty"`    //
	Description     string                                               `json:"description,omitempty"`     //
	DisplayName     string                                               `json:"displayName,omitempty"`     //
	Group           string                                               `json:"group,omitempty"`           //
	ID              string                                               `json:"id,omitempty"`              //
	InstructionText string                                               `json:"instructionText,omitempty"` //
	Key             string                                               `json:"key,omitempty"`             //
	NotParam        bool                                                 `json:"notParam,omitempty"`        //
	Order           int                                                  `json:"order,omitempty"`           //
	ParamArray      bool                                                 `json:"paramArray,omitempty"`      //
	ParameterName   string                                               `json:"parameterName,omitempty"`   //
	Provider        string                                               `json:"provider,omitempty"`        //
	Range           []UpdateTemplateRequestRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        bool                                                 `json:"required,omitempty"`        //
	Selection       UpdateTemplateRequestRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}

// UpdateTemplateRequestRollbackTemplateParamsRange is the updateTemplateRequestRollbackTemplateParamsRange definition
type UpdateTemplateRequestRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// UpdateTemplateRequestRollbackTemplateParamsSelection is the updateTemplateRequestRollbackTemplateParamsSelection definition
type UpdateTemplateRequestRollbackTemplateParamsSelection struct {
	ID              string                 `json:"id,omitempty"`              //
	SelectionType   string                 `json:"selectionType,omitempty"`   //
	SelectionValues map[string]interface{} `json:"selectionValues,omitempty"` //
}

// UpdateTemplateRequestTags is the updateTemplateRequestTags definition
type UpdateTemplateRequestTags []string

// UpdateTemplateRequestTemplateParams is the updateTemplateRequestTemplateParams definition
type UpdateTemplateRequestTemplateParams struct {
	Binding         string                                       `json:"binding,omitempty"`         //
	DataType        string                                       `json:"dataType,omitempty"`        //
	DefaultValue    string                                       `json:"defaultValue,omitempty"`    //
	Description     string                                       `json:"description,omitempty"`     //
	DisplayName     string                                       `json:"displayName,omitempty"`     //
	Group           string                                       `json:"group,omitempty"`           //
	ID              string                                       `json:"id,omitempty"`              //
	InstructionText string                                       `json:"instructionText,omitempty"` //
	Key             string                                       `json:"key,omitempty"`             //
	NotParam        bool                                         `json:"notParam,omitempty"`        //
	Order           int                                          `json:"order,omitempty"`           //
	ParamArray      bool                                         `json:"paramArray,omitempty"`      //
	ParameterName   string                                       `json:"parameterName,omitempty"`   //
	Provider        string                                       `json:"provider,omitempty"`        //
	Range           []UpdateTemplateRequestTemplateParamsRange   `json:"range,omitempty"`           //
	Required        bool                                         `json:"required,omitempty"`        //
	Selection       UpdateTemplateRequestTemplateParamsSelection `json:"selection,omitempty"`       //
}

// UpdateTemplateRequestTemplateParamsRange is the updateTemplateRequestTemplateParamsRange definition
type UpdateTemplateRequestTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// UpdateTemplateRequestTemplateParamsSelection is the updateTemplateRequestTemplateParamsSelection definition
type UpdateTemplateRequestTemplateParamsSelection struct {
	ID              string                 `json:"id,omitempty"`              //
	SelectionType   string                 `json:"selectionType,omitempty"`   //
	SelectionValues map[string]interface{} `json:"selectionValues,omitempty"` //
}

// VersionTemplateRequest is the versionTemplateRequest definition
type VersionTemplateRequest struct {
	Comments   string `json:"comments,omitempty"`   //
	TemplateID string `json:"templateId,omitempty"` //
}

// CreateProjectResponse is the createProjectResponse definition
type CreateProjectResponse struct {
	Response CreateProjectResponseResponse `json:"response,omitempty"` //
	Version  string                        `json:"version,omitempty"`  //
}

// CreateProjectResponseResponse is the createProjectResponseResponse definition
type CreateProjectResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// CreateTemplateResponse is the createTemplateResponse definition
type CreateTemplateResponse struct {
	Response CreateTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                         `json:"version,omitempty"`  //
}

// CreateTemplateResponseResponse is the createTemplateResponseResponse definition
type CreateTemplateResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeleteProjectResponse is the deleteProjectResponse definition
type DeleteProjectResponse struct {
	Response DeleteProjectResponseResponse `json:"response,omitempty"` //
	Version  string                        `json:"version,omitempty"`  //
}

// DeleteProjectResponseResponse is the deleteProjectResponseResponse definition
type DeleteProjectResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeleteTemplateResponse is the deleteTemplateResponse definition
type DeleteTemplateResponse struct {
	Response DeleteTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                         `json:"version,omitempty"`  //
}

// DeleteTemplateResponseResponse is the deleteTemplateResponseResponse definition
type DeleteTemplateResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeployTemplateResponse is the deployTemplateResponse definition
type DeployTemplateResponse struct {
	DeploymentID    string                          `json:"deploymentId,omitempty"`    //
	DeploymentName  string                          `json:"deploymentName,omitempty"`  //
	Devices         []DeployTemplateResponseDevices `json:"devices,omitempty"`         //
	Duration        string                          `json:"duration,omitempty"`        //
	EndTime         string                          `json:"endTime,omitempty"`         //
	ProjectName     string                          `json:"projectName,omitempty"`     //
	StartTime       string                          `json:"startTime,omitempty"`       //
	Status          string                          `json:"status,omitempty"`          //
	TemplateName    string                          `json:"templateName,omitempty"`    //
	TemplateVersion string                          `json:"templateVersion,omitempty"` //
}

// DeployTemplateResponseDevices is the deployTemplateResponseDevices definition
type DeployTemplateResponseDevices struct {
	DeviceID  string `json:"deviceId,omitempty"`  //
	Duration  string `json:"duration,omitempty"`  //
	EndTime   string `json:"endTime,omitempty"`   //
	IPAddress string `json:"ipAddress,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	StartTime string `json:"startTime,omitempty"` //
	Status    string `json:"status,omitempty"`    //
}

// GetProjectsResponse is the getProjectsResponse definition
type GetProjectsResponse struct {
	ID          string                         `json:"id,omitempty"`          //
	Name        string                         `json:"name,omitempty"`        //
	Templates   []GetProjectsResponseTemplates `json:"templates,omitempty"`   //
	IsDeletable bool                           `json:"isDeletable,omitempty"` //
}

// GetProjectsResponseTemplates is the getProjectsResponseTemplates definition
type GetProjectsResponseTemplates struct {
	Composite         bool    `json:"composite,omitempty"`         //
	ID                string  `json:"id,omitempty"`                //
	Name              string  `json:"name,omitempty"`              //
	Language          string  `json:"language,omitempty"`          //
	CustomParamsOrder bool    `json:"customParamsOrder,omitempty"` //
	LastUpdateTime    float64 `json:"lastUpdateTime,omitempty"`    //
	LatestVersionTime float64 `json:"latestVersionTime,omitempty"` //
	ProjectAssociated bool    `json:"projectAssociated,omitempty"` //
	DocumentDatabase  bool    `json:"documentDatabase,omitempty"`  //
}

// GetTemplateDeploymentStatusResponse is the getTemplateDeploymentStatusResponse definition
type GetTemplateDeploymentStatusResponse struct {
	DeploymentID    string                                       `json:"deploymentId,omitempty"`    //
	DeploymentName  string                                       `json:"deploymentName,omitempty"`  //
	Devices         []GetTemplateDeploymentStatusResponseDevices `json:"devices,omitempty"`         //
	Duration        string                                       `json:"duration,omitempty"`        //
	EndTime         string                                       `json:"endTime,omitempty"`         //
	ProjectName     string                                       `json:"projectName,omitempty"`     //
	StartTime       string                                       `json:"startTime,omitempty"`       //
	Status          string                                       `json:"status,omitempty"`          //
	TemplateName    string                                       `json:"templateName,omitempty"`    //
	TemplateVersion string                                       `json:"templateVersion,omitempty"` //
}

// GetTemplateDeploymentStatusResponseDevices is the getTemplateDeploymentStatusResponseDevices definition
type GetTemplateDeploymentStatusResponseDevices struct {
	DeviceID  string `json:"deviceId,omitempty"`  //
	Duration  string `json:"duration,omitempty"`  //
	EndTime   string `json:"endTime,omitempty"`   //
	IPAddress string `json:"ipAddress,omitempty"` //
	Name      string `json:"name,omitempty"`      //
	StartTime string `json:"startTime,omitempty"` //
	Status    string `json:"status,omitempty"`    //
}

// GetTemplateDetailsResponse is the getTemplateDetailsResponse definition
type GetTemplateDetailsResponse struct {
	Author                  string                                             `json:"author,omitempty"`                  //
	Composite               bool                                               `json:"composite,omitempty"`               //
	ContainingTemplates     []GetTemplateDetailsResponseContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              int                                                `json:"createTime,omitempty"`              //
	Description             string                                             `json:"description,omitempty"`             //
	DeviceTypes             []GetTemplateDetailsResponseDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                             `json:"failurePolicy,omitempty"`           //
	ID                      string                                             `json:"id,omitempty"`                      //
	LastUpdateTime          int                                                `json:"lastUpdateTime,omitempty"`          //
	Name                    string                                             `json:"name,omitempty"`                    //
	ParentTemplateID        string                                             `json:"parentTemplateId,omitempty"`        //
	ProjectID               string                                             `json:"projectId,omitempty"`               //
	ProjectName             string                                             `json:"projectName,omitempty"`             //
	RollbackTemplateContent string                                             `json:"rollbackTemplateContent,omitempty"` //
	RollbackTemplateParams  []GetTemplateDetailsResponseRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                             `json:"softwareType,omitempty"`            //
	SoftwareVariant         string                                             `json:"softwareVariant,omitempty"`         //
	SoftwareVersion         string                                             `json:"softwareVersion,omitempty"`         //
	Tags                    []GetTemplateDetailsResponseTags                   `json:"tags,omitempty"`                    //
	TemplateContent         string                                             `json:"templateContent,omitempty"`         //
	TemplateParams          []GetTemplateDetailsResponseTemplateParams         `json:"templateParams,omitempty"`          //
	Version                 string                                             `json:"version,omitempty"`                 //
}

// GetTemplateDetailsResponseContainingTemplates is the getTemplateDetailsResponseContainingTemplates definition
type GetTemplateDetailsResponseContainingTemplates struct {
	Composite bool   `json:"composite,omitempty"` //
	ID        string `json:"id,omitempty"`        //
	Name      string `json:"name,omitempty"`      //
	Version   string `json:"version,omitempty"`   //
}

// GetTemplateDetailsResponseDeviceTypes is the getTemplateDetailsResponseDeviceTypes definition
type GetTemplateDetailsResponseDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` //
	ProductSeries string `json:"productSeries,omitempty"` //
	ProductType   string `json:"productType,omitempty"`   //
}

// GetTemplateDetailsResponseRollbackTemplateParams is the getTemplateDetailsResponseRollbackTemplateParams definition
type GetTemplateDetailsResponseRollbackTemplateParams struct {
	Binding         string                                                    `json:"binding,omitempty"`         //
	DataType        string                                                    `json:"dataType,omitempty"`        //
	DefaultValue    string                                                    `json:"defaultValue,omitempty"`    //
	Description     string                                                    `json:"description,omitempty"`     //
	DisplayName     string                                                    `json:"displayName,omitempty"`     //
	Group           string                                                    `json:"group,omitempty"`           //
	ID              string                                                    `json:"id,omitempty"`              //
	InstructionText string                                                    `json:"instructionText,omitempty"` //
	Key             string                                                    `json:"key,omitempty"`             //
	NotParam        bool                                                      `json:"notParam,omitempty"`        //
	Order           int                                                       `json:"order,omitempty"`           //
	ParamArray      bool                                                      `json:"paramArray,omitempty"`      //
	ParameterName   string                                                    `json:"parameterName,omitempty"`   //
	Provider        string                                                    `json:"provider,omitempty"`        //
	Range           []GetTemplateDetailsResponseRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        bool                                                      `json:"required,omitempty"`        //
	Selection       GetTemplateDetailsResponseRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}

// GetTemplateDetailsResponseRollbackTemplateParamsRange is the getTemplateDetailsResponseRollbackTemplateParamsRange definition
type GetTemplateDetailsResponseRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// GetTemplateDetailsResponseRollbackTemplateParamsSelection is the getTemplateDetailsResponseRollbackTemplateParamsSelection definition
type GetTemplateDetailsResponseRollbackTemplateParamsSelection struct {
	ID              string                 `json:"id,omitempty"`              //
	SelectionType   string                 `json:"selectionType,omitempty"`   //
	SelectionValues map[string]interface{} `json:"selectionValues,omitempty"` //
}

// GetTemplateDetailsResponseTags is the getTemplateDetailsResponseTags definition
type GetTemplateDetailsResponseTags struct {
	ID   string `json:"id,omitempty"`   //
	Name string `json:"name,omitempty"` //
}

// GetTemplateDetailsResponseTemplateParams is the getTemplateDetailsResponseTemplateParams definition
type GetTemplateDetailsResponseTemplateParams struct {
	Binding         string                                            `json:"binding,omitempty"`         //
	DataType        string                                            `json:"dataType,omitempty"`        //
	DefaultValue    string                                            `json:"defaultValue,omitempty"`    //
	Description     string                                            `json:"description,omitempty"`     //
	DisplayName     string                                            `json:"displayName,omitempty"`     //
	Group           string                                            `json:"group,omitempty"`           //
	ID              string                                            `json:"id,omitempty"`              //
	InstructionText string                                            `json:"instructionText,omitempty"` //
	Key             string                                            `json:"key,omitempty"`             //
	NotParam        bool                                              `json:"notParam,omitempty"`        //
	Order           int                                               `json:"order,omitempty"`           //
	ParamArray      bool                                              `json:"paramArray,omitempty"`      //
	ParameterName   string                                            `json:"parameterName,omitempty"`   //
	Provider        string                                            `json:"provider,omitempty"`        //
	Range           []GetTemplateDetailsResponseTemplateParamsRange   `json:"range,omitempty"`           //
	Required        bool                                              `json:"required,omitempty"`        //
	Selection       GetTemplateDetailsResponseTemplateParamsSelection `json:"selection,omitempty"`       //
}

// GetTemplateDetailsResponseTemplateParamsRange is the getTemplateDetailsResponseTemplateParamsRange definition
type GetTemplateDetailsResponseTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       //
	MaxValue int    `json:"maxValue,omitempty"` //
	MinValue int    `json:"minValue,omitempty"` //
}

// GetTemplateDetailsResponseTemplateParamsSelection is the getTemplateDetailsResponseTemplateParamsSelection definition
type GetTemplateDetailsResponseTemplateParamsSelection struct {
	ID              string                 `json:"id,omitempty"`              //
	SelectionType   string                 `json:"selectionType,omitempty"`   //
	SelectionValues map[string]interface{} `json:"selectionValues,omitempty"` //
}

// GetTemplateVersionsResponse is the getTemplateVersionsResponse definition
type GetTemplateVersionsResponse struct {
	Composite    bool                                      `json:"composite,omitempty"`    //
	Name         string                                    `json:"name,omitempty"`         //
	ProjectID    string                                    `json:"projectId,omitempty"`    //
	ProjectName  string                                    `json:"projectName,omitempty"`  //
	TemplateID   string                                    `json:"templateId,omitempty"`   //
	VersionsInfo []GetTemplateVersionsResponseVersionsInfo `json:"versionsInfo,omitempty"` //
}

// GetTemplateVersionsResponseVersionsInfo is the getTemplateVersionsResponseVersionsInfo definition
type GetTemplateVersionsResponseVersionsInfo struct {
	Description    string  `json:"description,omitempty"`    //
	ID             string  `json:"id,omitempty"`             //
	VersionTime    float64 `json:"versionTime,omitempty"`    //
	Author         string  `json:"author,omitempty"`         //
	Version        string  `json:"version,omitempty"`        //
	VersionComment string  `json:"versionComment,omitempty"` //
}

// GetsTheTemplatesAvailableResponse is the getsTheTemplatesAvailableResponse definition
type GetsTheTemplatesAvailableResponse struct {
	Composite    bool                                            `json:"composite,omitempty"`    //
	Name         string                                          `json:"name,omitempty"`         //
	ProjectID    string                                          `json:"projectId,omitempty"`    //
	ProjectName  string                                          `json:"projectName,omitempty"`  //
	TemplateID   string                                          `json:"templateId,omitempty"`   //
	VersionsInfo []GetsTheTemplatesAvailableResponseVersionsInfo `json:"versionsInfo,omitempty"` //
}

// GetsTheTemplatesAvailableResponseVersionsInfo is the getsTheTemplatesAvailableResponseVersionsInfo definition
type GetsTheTemplatesAvailableResponseVersionsInfo struct {
	Description    string  `json:"description,omitempty"`    //
	ID             string  `json:"id,omitempty"`             //
	VersionTime    float64 `json:"versionTime,omitempty"`    //
	Author         string  `json:"author,omitempty"`         //
	Version        string  `json:"version,omitempty"`        //
	VersionComment string  `json:"versionComment,omitempty"` //
}

// PreviewTemplateResponse is the previewTemplateResponse definition
type PreviewTemplateResponse struct {
	CliPreview       string                                    `json:"cliPreview,omitempty"`       //
	TemplateID       string                                    `json:"templateId,omitempty"`       //
	ValidationErrors []PreviewTemplateResponseValidationErrors `json:"validationErrors,omitempty"` //
}

// PreviewTemplateResponseValidationErrors is the previewTemplateResponseValidationErrors definition
type PreviewTemplateResponseValidationErrors struct {
	Type    string `json:"type,omitempty"`    //
	Message string `json:"message,omitempty"` //
}

// UpdateProjectResponse is the updateProjectResponse definition
type UpdateProjectResponse struct {
	Response UpdateProjectResponseResponse `json:"response,omitempty"` //
	Version  string                        `json:"version,omitempty"`  //
}

// UpdateProjectResponseResponse is the updateProjectResponseResponse definition
type UpdateProjectResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// UpdateTemplateResponse is the updateTemplateResponse definition
type UpdateTemplateResponse struct {
	Response UpdateTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                         `json:"version,omitempty"`  //
}

// UpdateTemplateResponseResponse is the updateTemplateResponseResponse definition
type UpdateTemplateResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// VersionTemplateResponse is the versionTemplateResponse definition
type VersionTemplateResponse struct {
	Response VersionTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                          `json:"version,omitempty"`  //
}

// VersionTemplateResponseResponse is the versionTemplateResponseResponse definition
type VersionTemplateResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createProject")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation createTemplate")
	}

	result := response.Result().(*CreateTemplateResponse)
	return result, response, err
}

// DeleteProject deleteProject
/* Deletes an existing Project
@param projectID projectId
*/
func (s *ConfigurationTemplatesService) DeleteProject(projectID string) (*DeleteProjectResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{"+"projectId"+"}", fmt.Sprintf("%v", projectID), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteProjectResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteProject")
	}

	result := response.Result().(*DeleteProjectResponse)
	return result, response, err
}

// DeleteTemplate deleteTemplate
/* Deletes an existing template
@param templateID templateId
*/
func (s *ConfigurationTemplatesService) DeleteTemplate(templateID string) (*DeleteTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateID), -1)

	response, err := RestyClient.R().
		SetResult(&DeleteTemplateResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deleteTemplate")
	}

	result := response.Result().(*DeleteTemplateResponse)
	return result, response, err
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deployTemplate")
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
func (s *ConfigurationTemplatesService) GetProjects(getProjectsQueryParams *GetProjectsQueryParams) (*[]GetProjectsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/project"

	queryString, _ := query.Values(getProjectsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&[]GetProjectsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getProjects")
	}

	result := response.Result().(*[]GetProjectsResponse)
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTemplateDeploymentStatus")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTemplateDetails")
	}

	result := response.Result().(*GetTemplateDetailsResponse)
	return result, response, err
}

// GetTemplateVersions getTemplateVersions
/* Returns the versions of a specified template
@param templateID templateId
*/
func (s *ConfigurationTemplatesService) GetTemplateVersions(templateID string) (*[]GetTemplateVersionsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template/version/{templateId}"
	path = strings.Replace(path, "{"+"templateId"+"}", fmt.Sprintf("%v", templateID), -1)

	response, err := RestyClient.R().
		SetResult(&[]GetTemplateVersionsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getTemplateVersions")
	}

	result := response.Result().(*[]GetTemplateVersionsResponse)
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
func (s *ConfigurationTemplatesService) GetsTheTemplatesAvailable(getsTheTemplatesAvailableQueryParams *GetsTheTemplatesAvailableQueryParams) (*[]GetsTheTemplatesAvailableResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/template-programmer/template"

	queryString, _ := query.Values(getsTheTemplatesAvailableQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&[]GetsTheTemplatesAvailableResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getsTheTemplatesAvailable")
	}

	result := response.Result().(*[]GetsTheTemplatesAvailableResponse)
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation previewTemplate")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateProject")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation updateTemplate")
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

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation versionTemplate")
	}

	result := response.Result().(*VersionTemplateResponse)
	return result, response, err
}
