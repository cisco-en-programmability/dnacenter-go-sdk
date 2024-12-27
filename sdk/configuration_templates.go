package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ConfigurationTemplatesService service

type GetTemplateProjectsQueryParams struct {
	Name   string  `url:"name,omitempty"`   //Name of project to be searched
	Limit  float64 `url:"limit,omitempty"`  //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	Offset float64 `url:"offset,omitempty"` //The first record to show for this page; the first record is numbered 1.
}
type GetTemplateProjectCountQueryParams struct {
	Name string `url:"name,omitempty"` //Name of project to be searched
}
type CreatesACloneOfTheGivenTemplateQueryParams struct {
	ProjectID string `url:"projectId,omitempty"` //UUID of the project in which the template needs to be created
}
type GetsAListOfProjectsQueryParams struct {
	Name      string `url:"name,omitempty"`      //Name of project to be searched
	SortOrder string `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (des)
}
type ImportsTheProjectsProvidedQueryParams struct {
	DoVersion bool `url:"doVersion,omitempty"` //If this flag is true then it creates a new version of the template with the imported contents in case if the templates already exists. " If this flag is false and if template already exists, then operation fails with 'Template already exists' error
}
type ImportsTheTemplatesProvidedQueryParams struct {
	DoVersion bool `url:"doVersion,omitempty"` //If this flag is true then it creates a new version of the template with the imported contents in case if the templates already exists. " If this flag is false and if template already exists, then operation fails with 'Template already exists' error
}
type GetsTheTemplatesAvailableQueryParams struct {
	ProjectID                  string   `url:"projectId,omitempty"`                  //Filter template(s) based on project UUID
	SoftwareType               string   `url:"softwareType,omitempty"`               //Filter template(s) based software type
	SoftwareVersion            string   `url:"softwareVersion,omitempty"`            //Filter template(s) based softwareVersion
	ProductFamily              string   `url:"productFamily,omitempty"`              //Filter template(s) based on device family
	ProductSeries              string   `url:"productSeries,omitempty"`              //Filter template(s) based on device series
	ProductType                string   `url:"productType,omitempty"`                //Filter template(s) based on device type
	FilterConflictingTemplates bool     `url:"filterConflictingTemplates,omitempty"` //Filter template(s) based on confliting templates
	Tags                       []string `url:"tags,omitempty"`                       //Filter template(s) based on tags
	ProjectNames               []string `url:"projectNames,omitempty"`               //Filter template(s) based on project names
	UnCommitted                bool     `url:"unCommitted,omitempty"`                //Filter template(s) based on template commited or not
	SortOrder                  string   `url:"sortOrder,omitempty"`                  //Sort Order Ascending (asc) or Descending (des)
}
type GetsDetailsOfAGivenTemplateQueryParams struct {
	LatestVersion bool `url:"latestVersion,omitempty"` //latestVersion flag to get the latest versioned template
}
type DetachAListOfNetworkProfilesFromADayNCliTemplateQueryParams struct {
	ProfileID string `url:"profileId,omitempty"` //The id or ids of the network profile, retrievable from /dna/intent/api/v1/networkProfilesForSites. The maximum number of profile Ids allowed is 20.  A list of profile ids can be passed as a queryParameter in two ways:   a comma-separated string ( profileId=388a23e9-4739-4be7-a0aa-cc5a95d158dd,2726dc60-3a12-451e-947a-d972ebf58743), or...  as separate query parameters with the same name ( profileId=388a23e9-4739-4be7-a0aa-cc5a95d158dd&profileId=2726dc60-3a12-451e-947a-d972ebf58743
}
type GetTemplateVersionsQueryParams struct {
	VersionNumber int     `url:"versionNumber,omitempty"` //Filter response to only get the template version that matches this version number
	LatestVersion bool    `url:"latestVersion,omitempty"` //Filter response to only include the latest version of a template
	Order         string  `url:"order,omitempty"`         //Whether ascending or descending order should be used to sort the response.
	Limit         float64 `url:"limit,omitempty"`         //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	Offset        int     `url:"offset,omitempty"`        //The first record to show for this page; the first record is numbered 1.
}
type GetTemplateVersionsCountQueryParams struct {
	VersionNumber int  `url:"versionNumber,omitempty"` //Filter response to only get the template version that matches this version number
	LatestVersion bool `url:"latestVersion,omitempty"` //Filter response to only include the latest version of a template
}
type GetProjectsDetailsV2QueryParams struct {
	ID        string  `url:"id,omitempty"`        //Id of project to be searched
	Name      string  `url:"name,omitempty"`      //Name of project to be searched
	Offset    int     `url:"offset,omitempty"`    //Index of first result
	Limit     float64 `url:"limit,omitempty"`     //The number of records to show for this page;The minimum is 1, and the maximum is 500.
	SortOrder string  `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (dsc)
}
type GetTemplatesDetailsV2QueryParams struct {
	ID                         string   `url:"id,omitempty"`                         //Id of template to be searched
	Name                       string   `url:"name,omitempty"`                       //Name of template to be searched
	ProjectID                  string   `url:"projectId,omitempty"`                  //Filter template(s) based on project id
	ProjectName                string   `url:"projectName,omitempty"`                //Filter template(s) based on project name
	SoftwareType               string   `url:"softwareType,omitempty"`               //Filter template(s) based software type
	SoftwareVersion            string   `url:"softwareVersion,omitempty"`            //Filter template(s) based softwareVersion
	ProductFamily              string   `url:"productFamily,omitempty"`              //Filter template(s) based on device family
	ProductSeries              string   `url:"productSeries,omitempty"`              //Filter template(s) based on device series
	ProductType                string   `url:"productType,omitempty"`                //Filter template(s) based on device type
	FilterConflictingTemplates bool     `url:"filterConflictingTemplates,omitempty"` //Filter template(s) based on confliting templates
	Tags                       []string `url:"tags,omitempty"`                       //Filter template(s) based on tags
	UnCommitted                bool     `url:"unCommitted,omitempty"`                //Return uncommitted template
	SortOrder                  string   `url:"sortOrder,omitempty"`                  //Sort Order Ascending (asc) or Descending (dsc)
	AllTemplateAttributes      bool     `url:"allTemplateAttributes,omitempty"`      //Return all template attributes
	IncludeVersionDetails      bool     `url:"includeVersionDetails,omitempty"`      //Include template version details
	Offset                     int      `url:"offset,omitempty"`                     //Index of first result
	Limit                      float64  `url:"limit,omitempty"`                      //The number of records to show for this page;The minimum is 1, and the maximum is 500.
}

type ResponseConfigurationTemplatesCreateTemplateProject struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseConfigurationTemplatesCreateTemplateProjectResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesCreateTemplateProjectResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesGetTemplateProjects struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *[]ResponseConfigurationTemplatesGetTemplateProjectsResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateProjectsResponse struct {
	ProjectID string `json:"projectId,omitempty"` // UUID of the project

	Name string `json:"name,omitempty"` // Name of the project

	Description string `json:"description,omitempty"` // Description of the project

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Timestamp of when the project was updated or modified
}
type ResponseConfigurationTemplatesGetTemplateProjectCount struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseConfigurationTemplatesGetTemplateProjectCountResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateProjectCountResponse struct {
	Count *int `json:"count,omitempty"` // The reported count
}
type ResponseConfigurationTemplatesGetTemplateProject struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseConfigurationTemplatesGetTemplateProjectResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateProjectResponse struct {
	ProjectID string `json:"projectId,omitempty"` // UUID of the project

	Name string `json:"name,omitempty"` // Name of the project

	Description string `json:"description,omitempty"` // Description of the project

	LastUpdateTime *float64 `json:"lastUpdateTime,omitempty"` // Timestamp of when the project was updated or modified
}
type ResponseConfigurationTemplatesUpdateTemplateProject struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseConfigurationTemplatesUpdateTemplateProjectResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesUpdateTemplateProjectResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesDeleteTemplateProject struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseConfigurationTemplatesDeleteTemplateProjectResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesDeleteTemplateProjectResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplate struct {
	Response *ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplateResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesCreateProject struct {
	Response *ResponseConfigurationTemplatesCreateProjectResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesCreateProjectResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesUpdateProject struct {
	Response *ResponseConfigurationTemplatesUpdateProjectResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesUpdateProjectResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsAListOfProjects []ResponseItemConfigurationTemplatesGetsAListOfProjects // Array of ResponseConfigurationTemplatesGetsAListOfProjects
type ResponseItemConfigurationTemplatesGetsAListOfProjects struct {
	Tags           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTags      `json:"tags,omitempty"`           //
	CreateTime     *int                                                              `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                            `json:"description,omitempty"`    // Description of project
	ID             string                                                            `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                              `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                            `json:"name,omitempty"`           // Name of project
	Templates      *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates struct {
	Tags                    *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                  `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                   `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                    `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                   `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                  `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                  `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                  `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                  `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                    `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                    `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                  `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                  `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                  `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                  `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                  `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                  `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                  `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                  `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                  `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                  `json:"version,omitempty"`                 // Current version of template
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplates struct {
	Tags                   *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                      `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                     `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                     `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                     `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                     `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                     `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                     `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                     `json:"version,omitempty"`                // Current version of template
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                            `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                              `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                            `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                            `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                            `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                            `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                            `json:"group,omitempty"`           // group
	ID              string                                                                                                            `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                            `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                            `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                             `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                              `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                             `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                            `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                            `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                             `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                         `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                           `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                           `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                    `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                      `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                    `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                    `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                    `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                    `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                    `json:"group,omitempty"`           // group
	ID              string                                                                                                    `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                    `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                    `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                     `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                      `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                     `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                    `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                    `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                     `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                 `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                   `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                   `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParams struct {
	Binding         string                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors struct {
	RollbackTemplateErrors *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors interface{}
type ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesImportsTheProjectsProvided struct {
	Response *ResponseConfigurationTemplatesImportsTheProjectsProvidedResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesImportsTheProjectsProvidedResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteria struct {
	Response *ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaResponse `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteriaResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesImportsTheTemplatesProvided struct {
	Response *ResponseConfigurationTemplatesImportsTheTemplatesProvidedResponse `json:"response,omitempty"` //
	Version  string                                                             `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesImportsTheTemplatesProvidedResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject struct {
	Tags           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags      `json:"tags,omitempty"`           //
	CreateTime     *int                                                                    `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                                  `json:"description,omitempty"`    // Description of project
	ID             string                                                                  `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                                    `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                                  `json:"name,omitempty"`           // Name of project
	Templates      *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates struct {
	Tags                    *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                        `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                         `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                          `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                         `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                        `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                        `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                        `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                        `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                          `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                          `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                        `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                        `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                        `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                        `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                        `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                        `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                        `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                        `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                        `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                        `json:"version,omitempty"`                 // Current version of template
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplates struct {
	Tags                   *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                            `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                           `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                           `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                           `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                           `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                           `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                           `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                           `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                                  `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                                    `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                                  `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                                  `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                                  `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                                  `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                                  `json:"group,omitempty"`           // group
	ID              string                                                                                                                  `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                                  `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                                  `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                                   `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                                    `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                                   `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                                  `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                                  `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                                   `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                               `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                                 `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                                 `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                          `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                            `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                          `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                          `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                          `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                          `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                          `json:"group,omitempty"`           // group
	ID              string                                                                                                          `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                          `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                          `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                           `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                            `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                           `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                          `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                          `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                           `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                       `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                         `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                         `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                               `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                 `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                               `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                               `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                               `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                               `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                               `json:"group,omitempty"`           // group
	ID              string                                                                                               `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                               `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                               `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                 `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                               `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                               `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParams struct {
	Binding         string                                                                                       `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                         `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                       `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                       `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                       `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                       `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                       `json:"group,omitempty"`           // group
	ID              string                                                                                       `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                       `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                       `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                        `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                         `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                        `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                       `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                       `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors struct {
	RollbackTemplateErrors *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                      `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                      `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesDeletesTheProject struct {
	Response *ResponseConfigurationTemplatesDeletesTheProjectResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesDeletesTheProjectResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesCreateTemplate struct {
	Response *ResponseConfigurationTemplatesCreateTemplateResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesCreateTemplateResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsTheTemplatesAvailable []ResponseItemConfigurationTemplatesGetsTheTemplatesAvailable // Array of ResponseConfigurationTemplatesGetsTheTemplatesAvailable
type ResponseItemConfigurationTemplatesGetsTheTemplatesAvailable struct {
	Composite    *bool                                                                      `json:"composite,omitempty"`    // Is it composite template
	Name         string                                                                     `json:"name,omitempty"`         // Name of template
	ProjectID    string                                                                     `json:"projectId,omitempty"`    // UUID of project
	ProjectName  string                                                                     `json:"projectName,omitempty"`  // Name of project
	TemplateID   string                                                                     `json:"templateId,omitempty"`   // UUID of template
	VersionsInfo *[]ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableVersionsInfo `json:"versionsInfo,omitempty"` //
}
type ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableVersionsInfo struct {
	Author         string `json:"author,omitempty"`         // Author of version template
	Description    string `json:"description,omitempty"`    // Description of template
	ID             string `json:"id,omitempty"`             // UUID of template
	Version        string `json:"version,omitempty"`        // Current version of template
	VersionComment string `json:"versionComment,omitempty"` // Version comment
	VersionTime    *int   `json:"versionTime,omitempty"`    // Template version time
}
type ResponseConfigurationTemplatesUpdateTemplate struct {
	Response *ResponseConfigurationTemplatesUpdateTemplateResponse `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesUpdateTemplateResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesDeployTemplate struct {
	DeploymentID    string                                                 `json:"deploymentId,omitempty"`    // UUID of deployment
	DeploymentName  string                                                 `json:"deploymentName,omitempty"`  // Name of deployment
	Devices         *[]ResponseConfigurationTemplatesDeployTemplateDevices `json:"devices,omitempty"`         //
	Duration        string                                                 `json:"duration,omitempty"`        // Total deployment duration
	EndTime         string                                                 `json:"endTime,omitempty"`         // Deployment end time
	ProjectName     string                                                 `json:"projectName,omitempty"`     // Name of project
	StartTime       string                                                 `json:"startTime,omitempty"`       // Deployment start time
	Status          string                                                 `json:"status,omitempty"`          // Current status of deployment
	StatusMessage   string                                                 `json:"statusMessage,omitempty"`   // Status message of deployment
	TemplateName    string                                                 `json:"templateName,omitempty"`    // Name of template deployed
	TemplateVersion string                                                 `json:"templateVersion,omitempty"` // Version of template deployed
}
type ResponseConfigurationTemplatesDeployTemplateDevices struct {
	DetailedStatusMessage string `json:"detailedStatusMessage,omitempty"` // Device detailed status message
	DeviceID              string `json:"deviceId,omitempty"`              // UUID of device
	Duration              string `json:"duration,omitempty"`              // Total duration of deployment
	EndTime               string `json:"endTime,omitempty"`               // EndTime of deployment
	IDentifier            string `json:"identifier,omitempty"`            // Identifier of device based on the target type
	IPAddress             string `json:"ipAddress,omitempty"`             // Device IPAddress
	Name                  string `json:"name,omitempty"`                  // Name of device
	StartTime             string `json:"startTime,omitempty"`             // StartTime of deployment
	Status                string `json:"status,omitempty"`                // Current status of device
	TargetType            string `json:"targetType,omitempty"`            // Target type of device
}
type ResponseConfigurationTemplatesStatusOfTemplateDeployment struct {
	DeploymentID    string                                                             `json:"deploymentId,omitempty"`    // UUID of deployment
	DeploymentName  string                                                             `json:"deploymentName,omitempty"`  // Name of deployment
	Devices         *[]ResponseConfigurationTemplatesStatusOfTemplateDeploymentDevices `json:"devices,omitempty"`         //
	Duration        string                                                             `json:"duration,omitempty"`        // Total deployment duration
	EndTime         string                                                             `json:"endTime,omitempty"`         // Deployment end time
	ProjectName     string                                                             `json:"projectName,omitempty"`     // Name of project
	StartTime       string                                                             `json:"startTime,omitempty"`       // Deployment start time
	Status          string                                                             `json:"status,omitempty"`          // Current status of deployment
	StatusMessage   string                                                             `json:"statusMessage,omitempty"`   // Status message of deployment
	TemplateName    string                                                             `json:"templateName,omitempty"`    // Name of template deployed
	TemplateVersion string                                                             `json:"templateVersion,omitempty"` // Version of template deployed
}
type ResponseConfigurationTemplatesStatusOfTemplateDeploymentDevices struct {
	DetailedStatusMessage string `json:"detailedStatusMessage,omitempty"` // Device detailed status message
	DeviceID              string `json:"deviceId,omitempty"`              // UUID of device
	Duration              string `json:"duration,omitempty"`              // Total duration of deployment
	EndTime               string `json:"endTime,omitempty"`               // EndTime of deployment
	IDentifier            string `json:"identifier,omitempty"`            // Identifier of device based on the target type
	IPAddress             string `json:"ipAddress,omitempty"`             // Device IPAddress
	Name                  string `json:"name,omitempty"`                  // Name of device
	StartTime             string `json:"startTime,omitempty"`             // StartTime of deployment
	Status                string `json:"status,omitempty"`                // Current status of device
	TargetType            string `json:"targetType,omitempty"`            // Target type of device
}
type ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteria struct {
	Response *ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteriaResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesPreviewTemplate struct {
	CliPreview       string                                                         `json:"cliPreview,omitempty"`       // Generated template preview
	DeviceID         string                                                         `json:"deviceId,omitempty"`         // UUID of device
	TemplateID       string                                                         `json:"templateId,omitempty"`       // UUID of template
	ValidationErrors *ResponseConfigurationTemplatesPreviewTemplateValidationErrors `json:"validationErrors,omitempty"` // Validation error in template content if any
}
type ResponseConfigurationTemplatesPreviewTemplateValidationErrors interface{}
type ResponseConfigurationTemplatesVersionTemplate struct {
	Response *ResponseConfigurationTemplatesVersionTemplateResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesVersionTemplateResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate []ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate // Array of ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate
type ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate struct {
	Composite    *bool                                                                               `json:"composite,omitempty"`    // Is it composite template
	Name         string                                                                              `json:"name,omitempty"`         // Name of template
	ProjectID    string                                                                              `json:"projectId,omitempty"`    // UUID of project
	ProjectName  string                                                                              `json:"projectName,omitempty"`  // Name of project
	TemplateID   string                                                                              `json:"templateId,omitempty"`   // UUID of template
	VersionsInfo *[]ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateVersionsInfo `json:"versionsInfo,omitempty"` //
}
type ResponseItemConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplateVersionsInfo struct {
	Author         string `json:"author,omitempty"`         // Author of version template
	Description    string `json:"description,omitempty"`    // Description of template
	ID             string `json:"id,omitempty"`             // UUID of template
	Version        string `json:"version,omitempty"`        // Current version of template
	VersionComment string `json:"versionComment,omitempty"` // Version comment
	VersionTime    *int   `json:"versionTime,omitempty"`    // Template version time
}
type ResponseConfigurationTemplatesDeletesTheTemplate struct {
	Response *ResponseConfigurationTemplatesDeletesTheTemplateResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesDeletesTheTemplateResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate struct {
	Tags                    *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                             `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                              `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                               `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                              `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                             `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                             `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                             `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                             `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                               `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                               `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                             `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                             `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                             `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                             `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                             `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                             `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                             `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                             `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                             `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                             `json:"version,omitempty"`                 // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplates struct {
	Tags                   *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                 `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                       `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                         `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                       `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                       `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                       `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                       `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                       `json:"group,omitempty"`           // group
	ID              string                                                                                                       `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                       `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                       `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                        `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                         `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                        `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                       `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                       `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParams struct {
	Binding         string                                                                                               `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                 `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                               `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                               `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                               `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                               `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                               `json:"group,omitempty"`           // group
	ID              string                                                                                               `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                               `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                               `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                 `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                               `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                               `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParams struct {
	Binding         string                                                                                    `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                      `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                    `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                    `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                    `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                    `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                    `json:"group,omitempty"`           // group
	ID              string                                                                                    `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                    `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                    `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                     `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                      `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                     `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                    `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                    `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                     `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                 `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                   `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                   `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParams struct {
	Binding         string                                                                            `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                              `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                            `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                            `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                            `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                            `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                            `json:"group,omitempty"`           // group
	ID              string                                                                            `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                            `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                            `json:"key,omitempty"`             // key
	NotParam        *bool                                                                             `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                              `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                             `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                            `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                            `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                             `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                         `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                           `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                           `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrors struct {
	RollbackTemplateErrors *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                           `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                           `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesRetrieveTheNetworkProfilesAttachedToACLITemplate struct {
	object string `json:"object,omitempty"` // object
}
type ResponseConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate struct {
	Version  string                                                                        `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesAttachNetworkProfileToADayNCliTemplateResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesAttachNetworkProfileToADayNCliTemplateResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesDetachAListOfNetworkProfilesFromADayNCliTemplate struct {
	Version  string                                                                                  `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesDetachAListOfNetworkProfilesFromADayNCliTemplateResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesDetachAListOfNetworkProfilesFromADayNCliTemplateResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplate struct {
	Version  string                                                                                `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplateResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplateResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplate struct {
	Version  string                                                                                      `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplateResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplateResponse struct {
	Count *int `json:"count,omitempty"` // The reported count
}
type ResponseConfigurationTemplatesDetachANetworkProfileFromADayNCliTemplate struct {
	Version  string                                                                           `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesDetachANetworkProfileFromADayNCliTemplateResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesDetachANetworkProfileFromADayNCliTemplateResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesGetTemplateVersions struct {
	Version  string                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *[]ResponseConfigurationTemplatesGetTemplateVersionsResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateVersionsResponse struct {
	VersionID         string                                                                      `json:"versionId,omitempty"`         // The id of this version
	Version           *int                                                                        `json:"version,omitempty"`           // The version number of this version
	VersionTime       *float64                                                                    `json:"versionTime,omitempty"`       // Time at which this version was committed
	RegularTemplate   *ResponseConfigurationTemplatesGetTemplateVersionsResponseRegularTemplate   `json:"RegularTemplate,omitempty"`   //
	CompositeTemplate *ResponseConfigurationTemplatesGetTemplateVersionsResponseCompositeTemplate `json:"CompositeTemplate,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateVersionsResponseRegularTemplate struct {
	TemplateID      string                                                                              `json:"templateId,omitempty"`      // The id of the template
	Name            string                                                                              `json:"name,omitempty"`            // Name of the template
	ProjectID       string                                                                              `json:"projectId,omitempty"`       // Id of the project
	Description     string                                                                              `json:"description,omitempty"`     // Description of the template
	SoftwareFamily  string                                                                              `json:"softwareFamily,omitempty"`  // Software Family
	Author          string                                                                              `json:"author,omitempty"`          // Author of the template
	Products        *[]ResponseConfigurationTemplatesGetTemplateVersionsResponseRegularTemplateProducts `json:"products,omitempty"`        //
	LastUpdateTime  *float64                                                                            `json:"lastUpdateTime,omitempty"`  // Timestamp of when the template was updated or modified
	Type            string                                                                              `json:"type,omitempty"`            // The type of the template
	Language        string                                                                              `json:"language,omitempty"`        // Language of the template
	TemplateContent string                                                                              `json:"templateContent,omitempty"` // Template content (uses LF styling for line-breaks)
}
type ResponseConfigurationTemplatesGetTemplateVersionsResponseRegularTemplateProducts struct {
	ProductFamily string `json:"productFamily,omitempty"` // Family name of the product
	ProductSeries string `json:"productSeries,omitempty"` // Series name of the product
	ProductName   string `json:"productName,omitempty"`   // Name of the product
}
type ResponseConfigurationTemplatesGetTemplateVersionsResponseCompositeTemplate struct {
	TemplateID     string                                                                                `json:"templateId,omitempty"`     // The id of the template
	Name           string                                                                                `json:"name,omitempty"`           // Name of the template
	ProjectID      string                                                                                `json:"projectId,omitempty"`      // Id of the project
	Description    string                                                                                `json:"description,omitempty"`    // Description of the template
	SoftwareFamily string                                                                                `json:"softwareFamily,omitempty"` // Software Family
	Author         string                                                                                `json:"author,omitempty"`         // Author of the template
	Products       *[]ResponseConfigurationTemplatesGetTemplateVersionsResponseCompositeTemplateProducts `json:"products,omitempty"`       //
	LastUpdateTime *float64                                                                              `json:"lastUpdateTime,omitempty"` // Timestamp of when the template was updated or modified
	Type           string                                                                                `json:"type,omitempty"`           // The type of the template
	FailurePolicy  string                                                                                `json:"failurePolicy,omitempty"`  // Policy to handle failure only applicable for composite templates  CONTINUE_ON_ERROR: If a composed template fails while deploying a device, continue deploying the next composed template  ABORT_TARGET_ON_ERROR: If a composed template fails while deploying to a device, abort the subsequent composed templates to that device if there any remaining
}
type ResponseConfigurationTemplatesGetTemplateVersionsResponseCompositeTemplateProducts struct {
	ProductFamily string `json:"productFamily,omitempty"` // Family name of the product
	ProductSeries string `json:"productSeries,omitempty"` // Series name of the product
	ProductName   string `json:"productName,omitempty"`   // Name of the product
}
type ResponseConfigurationTemplatesCommitTemplateForANewVersion struct {
	Version  string                                                              `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesCommitTemplateForANewVersionResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesCommitTemplateForANewVersionResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseConfigurationTemplatesGetTemplateVersionsCount struct {
	Version  string                                                          `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesGetTemplateVersionsCountResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateVersionsCountResponse struct {
	Count *int `json:"count,omitempty"` // The reported count
}
type ResponseConfigurationTemplatesGetTemplateVersion struct {
	Version  string                                                    `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseConfigurationTemplatesGetTemplateVersionResponse `json:"response,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateVersionResponse struct {
	VersionID         string                                                                     `json:"versionId,omitempty"`         // The id of this version
	Version           *int                                                                       `json:"version,omitempty"`           // The version number of this version
	VersionTime       *float64                                                                   `json:"versionTime,omitempty"`       // Time at which this version was committed
	RegularTemplate   *ResponseConfigurationTemplatesGetTemplateVersionResponseRegularTemplate   `json:"RegularTemplate,omitempty"`   //
	CompositeTemplate *ResponseConfigurationTemplatesGetTemplateVersionResponseCompositeTemplate `json:"CompositeTemplate,omitempty"` //
}
type ResponseConfigurationTemplatesGetTemplateVersionResponseRegularTemplate struct {
	TemplateID      string                                                                             `json:"templateId,omitempty"`      // The id of the template
	Name            string                                                                             `json:"name,omitempty"`            // Name of the template
	ProjectID       string                                                                             `json:"projectId,omitempty"`       // Id of the project
	Description     string                                                                             `json:"description,omitempty"`     // Description of the template
	SoftwareFamily  string                                                                             `json:"softwareFamily,omitempty"`  // Software Family
	Author          string                                                                             `json:"author,omitempty"`          // Author of the template
	Products        *[]ResponseConfigurationTemplatesGetTemplateVersionResponseRegularTemplateProducts `json:"products,omitempty"`        //
	LastUpdateTime  *float64                                                                           `json:"lastUpdateTime,omitempty"`  // Timestamp of when the template was updated or modified
	Type            string                                                                             `json:"type,omitempty"`            // The type of the template
	Language        string                                                                             `json:"language,omitempty"`        // Language of the template
	TemplateContent string                                                                             `json:"templateContent,omitempty"` // Template content (uses LF styling for line-breaks)
}
type ResponseConfigurationTemplatesGetTemplateVersionResponseRegularTemplateProducts struct {
	ProductFamily string `json:"productFamily,omitempty"` // Family name of the product
	ProductSeries string `json:"productSeries,omitempty"` // Series name of the product
	ProductName   string `json:"productName,omitempty"`   // Name of the product
}
type ResponseConfigurationTemplatesGetTemplateVersionResponseCompositeTemplate struct {
	TemplateID     string                                                                               `json:"templateId,omitempty"`     // The id of the template
	Name           string                                                                               `json:"name,omitempty"`           // Name of the template
	ProjectID      string                                                                               `json:"projectId,omitempty"`      // Id of the project
	Description    string                                                                               `json:"description,omitempty"`    // Description of the template
	SoftwareFamily string                                                                               `json:"softwareFamily,omitempty"` // Software Family
	Author         string                                                                               `json:"author,omitempty"`         // Author of the template
	Products       *[]ResponseConfigurationTemplatesGetTemplateVersionResponseCompositeTemplateProducts `json:"products,omitempty"`       //
	LastUpdateTime *float64                                                                             `json:"lastUpdateTime,omitempty"` // Timestamp of when the template was updated or modified
	Type           string                                                                               `json:"type,omitempty"`           // The type of the template
	FailurePolicy  string                                                                               `json:"failurePolicy,omitempty"`  // Policy to handle failure only applicable for composite templates  CONTINUE_ON_ERROR: If a composed template fails while deploying a device, continue deploying the next composed template  ABORT_TARGET_ON_ERROR: If a composed template fails while deploying to a device, abort the subsequent composed templates to that device if there any remaining
}
type ResponseConfigurationTemplatesGetTemplateVersionResponseCompositeTemplateProducts struct {
	ProductFamily string `json:"productFamily,omitempty"` // Family name of the product
	ProductSeries string `json:"productSeries,omitempty"` // Series name of the product
	ProductName   string `json:"productName,omitempty"`   // Name of the product
}
type ResponseConfigurationTemplatesGetProjectsDetailsV2 struct {
	CreateTime     *int                                                         `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                       `json:"description,omitempty"`    // Description of project
	ID             string                                                       `json:"id,omitempty"`             // UUID of project
	IsDeletable    *bool                                                        `json:"isDeletable,omitempty"`    // Flag to check if project is deletable or not(for internal use only)
	LastUpdateTime *int                                                         `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                       `json:"name,omitempty"`           // Name of project
	Tags           *[]ResponseConfigurationTemplatesGetProjectsDetailsV2Tags    `json:"tags,omitempty"`           //
	Templates      *ResponseConfigurationTemplatesGetProjectsDetailsV2Templates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseConfigurationTemplatesGetProjectsDetailsV2Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetProjectsDetailsV2Templates interface{}
type ResponseConfigurationTemplatesGetTemplatesV2Details struct {
	Response []ResponseConfigurationTemplatesGetTemplatesDetailsV2 `json:"response,omitempty"` // Response
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2 struct {
	Author                  string                                                                       `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                        `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                         `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                        `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                       `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2DeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                       `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                       `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                       `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                         `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                         `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                       `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                       `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectAssociated       *bool                                                                        `json:"projectAssociated,omitempty"`       //
	ProjectID               string                                                                       `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                       `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                       `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                       `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                       `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                       `json:"softwareVersion,omitempty"`         // Applicable device software version
	Tags                    *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2Tags                   `json:"tags,omitempty"`                    //
	TemplateContent         string                                                                       `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                       `json:"version,omitempty"`                 // Current version of template
	VersionsInfo            *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2VersionsInfo           `json:"versionsInfo,omitempty"`            //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplates struct {
	Composite              *bool                                                                                           `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                          `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                          `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                          `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                          `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                          `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	Tags                   *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	TemplateContent        string                                                                                          `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                          `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2DeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParams struct {
	Binding         string                                                                              `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                              `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                              `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                              `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                              `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                              `json:"group,omitempty"`           // group
	ID              string                                                                              `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                              `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                              `json:"key,omitempty"`             // key
	NotParam        *bool                                                                               `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                               `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                              `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                              `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                               `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2Tags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParams struct {
	Binding         string                                                                      `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                        `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                      `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                      `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                      `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                      `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                      `json:"group,omitempty"`           // group
	ID              string                                                                      `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                      `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                      `json:"key,omitempty"`             // key
	NotParam        *bool                                                                       `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                        `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                       `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                      `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                      `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrors struct {
	RollbackTemplateErrors *ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                     `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                     `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsV2VersionsInfo struct {
	Author         string `json:"author,omitempty"`         //
	Description    string `json:"description,omitempty"`    //
	ID             string `json:"id,omitempty"`             //
	Version        string `json:"version,omitempty"`        //
	VersionComment string `json:"versionComment,omitempty"` //
	VersionTime    *int   `json:"versionTime,omitempty"`    //
}
type ResponseConfigurationTemplatesDeployTemplateV2 struct {
	Response *ResponseConfigurationTemplatesDeployTemplateV2Response `json:"response,omitempty"` //
	Version  string                                                  `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesDeployTemplateV2Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestConfigurationTemplatesCreateTemplateProject struct {
	Name string `json:"name,omitempty"` // Name of the project

	Description string `json:"description,omitempty"` // Description of the project
}
type RequestConfigurationTemplatesUpdateTemplateProject struct {
	Name string `json:"name,omitempty"` // Name of the project

	Description string `json:"description,omitempty"` // Description of the project
}
type RequestConfigurationTemplatesCreateProject struct {
	Tags           *[]RequestConfigurationTemplatesCreateProjectTags    `json:"tags,omitempty"`           //
	CreateTime     *int                                                 `json:"createTime,omitempty"`     // Create time of project
	Description    string                                               `json:"description,omitempty"`    // Description of project
	ID             string                                               `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                 `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                               `json:"name,omitempty"`           // Name of project
	Templates      *RequestConfigurationTemplatesCreateProjectTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type RequestConfigurationTemplatesCreateProjectTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateProjectTemplates interface{}
type RequestConfigurationTemplatesUpdateProject struct {
	Tags           *[]RequestConfigurationTemplatesUpdateProjectTags    `json:"tags,omitempty"`           //
	CreateTime     *int                                                 `json:"createTime,omitempty"`     // Create time of project
	Description    string                                               `json:"description,omitempty"`    // Description of project
	ID             string                                               `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                 `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                               `json:"name,omitempty"`           // Name of project
	Templates      *RequestConfigurationTemplatesUpdateProjectTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type RequestConfigurationTemplatesUpdateProjectTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateProjectTemplates interface{}                                                                         // # Review unknown case
type RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteria []RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteria // Array of RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteria
type RequestItemConfigurationTemplatesExportsTheProjectsForAGivenCriteria interface{}
type RequestConfigurationTemplatesImportsTheTemplatesProvided []RequestItemConfigurationTemplatesImportsTheTemplatesProvided // Array of RequestConfigurationTemplatesImportsTheTemplatesProvided
type RequestItemConfigurationTemplatesImportsTheTemplatesProvided struct {
	Tags                    *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                 `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                  `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                 `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                  `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                  `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                `json:"version,omitempty"`                 // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplates struct {
	Tags                   *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                    `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                   `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                   `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                   `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                   `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                   `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                   `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                   `json:"version,omitempty"`                // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                          `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                            `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                          `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                          `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                          `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                          `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                          `json:"group,omitempty"`           // group
	ID              string                                                                                                          `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                          `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                          `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                           `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                            `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                           `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                          `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                          `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                           `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                       `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                         `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                         `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParams struct {
	Binding         string                                                                                                  `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                    `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                  `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                  `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                  `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                  `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                  `json:"group,omitempty"`           // group
	ID              string                                                                                                  `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                  `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                  `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                   `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                    `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                   `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                  `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                  `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                   `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                               `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                 `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                 `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParams struct {
	Binding         string                                                                                       `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                         `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                       `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                       `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                       `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                       `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                       `json:"group,omitempty"`           // group
	ID              string                                                                                       `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                       `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                       `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                        `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                         `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                        `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                       `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                       `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}

type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParams struct {
	Binding         string                                                                               `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                 `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                               `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                               `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                               `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                               `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                               `json:"group,omitempty"`           // group
	ID              string                                                                               `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                               `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                               `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                 `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                               `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                               `json:"provider,omitempty"`        // provider
	Range           *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedTemplateParamsSelectionSelectionValues interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrors struct {
	RollbackTemplateErrors *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                              `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                              `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors interface{}
type RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesCreateTemplate struct {
	Tags                    *[]RequestConfigurationTemplatesCreateTemplateTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                               `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesCreateTemplateContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                 `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                               `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesCreateTemplateDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                               `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                               `json:"id,omitempty"`                      // UUID of template
	Language                string                                                               `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                 `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                 `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                               `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                               `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                               `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                               `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                               `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                               `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                               `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                               `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                               `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesCreateTemplateTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesCreateTemplateValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                               `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                   `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                  `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                  `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                  `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                  `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                  `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                  `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                  `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams struct {
	Binding         string                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams struct {
	Binding         string                                                                      `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                        `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                      `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                      `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                      `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                      `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                      `json:"group,omitempty"`           // group
	ID              string                                                                      `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                      `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                      `json:"key,omitempty"`             // key
	NotParam        *bool                                                                       `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                        `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                       `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                      `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                      `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateTemplateParams struct {
	Binding         string                                                              `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                              `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                              `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                              `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                              `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                              `json:"group,omitempty"`           // group
	ID              string                                                              `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                              `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                              `json:"key,omitempty"`             // key
	NotParam        *bool                                                               `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                               `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                              `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                              `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesCreateTemplateTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                               `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateTemplateTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateTemplateTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateTemplateTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateTemplateTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateTemplateTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateTemplateValidationErrors struct {
	RollbackTemplateErrors *RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                             `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                             `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateTemplate struct {
	Tags                    *[]RequestConfigurationTemplatesUpdateTemplateTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                               `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                 `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                               `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesUpdateTemplateDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                               `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                               `json:"id,omitempty"`                      // UUID of template
	Language                string                                                               `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                 `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                 `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                               `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                               `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                               `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                               `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                               `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                               `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                               `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                               `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                               `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesUpdateTemplateTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesUpdateTemplateValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                               `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                   `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                  `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                  `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                  `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                  `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                  `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                  `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                  `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                         `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                           `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                         `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                         `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                         `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                         `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                         `json:"group,omitempty"`           // group
	ID              string                                                                                         `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                         `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                         `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                          `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                           `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                          `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                         `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                         `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams struct {
	Binding         string                                                                                 `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                   `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                 `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                 `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                 `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                 `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                 `json:"group,omitempty"`           // group
	ID              string                                                                                 `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                 `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                 `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                  `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                   `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                  `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                 `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                 `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams struct {
	Binding         string                                                                      `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                        `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                      `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                      `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                      `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                      `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                      `json:"group,omitempty"`           // group
	ID              string                                                                      `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                      `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                      `json:"key,omitempty"`             // key
	NotParam        *bool                                                                       `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                        `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                       `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                      `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                      `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateTemplateParams struct {
	Binding         string                                                              `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                              `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                              `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                              `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                              `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                              `json:"group,omitempty"`           // group
	ID              string                                                              `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                              `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                              `json:"key,omitempty"`             // key
	NotParam        *bool                                                               `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                               `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                              `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                              `json:"provider,omitempty"`        // provider
	Range           *[]RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                               `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateTemplateValidationErrors struct {
	RollbackTemplateErrors *RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                             `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                             `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesDeployTemplate struct {
	ForcePushTemplate            *bool                                                    `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  *bool                                                    `json:"isComposite,omitempty"`                  // Composite template flag
	MainTemplateID               string                                                   `json:"mainTemplateId,omitempty"`               // Main template UUID of versioned template
	MemberTemplateDeploymentInfo []string                                                 `json:"memberTemplateDeploymentInfo,omitempty"` // memberTemplateDeploymentInfo
	TargetInfo                   *[]RequestConfigurationTemplatesDeployTemplateTargetInfo `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                                                   `json:"templateId,omitempty"`                   // UUID of template to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateTargetInfo struct {
	HostName            string                                                       `json:"hostName,omitempty"`            // Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
	ID                  string                                                       `json:"id,omitempty"`                  // UUID of target is required if targetType is MANAGED_DEVICE_UUID
	Params              *RequestConfigurationTemplatesDeployTemplateTargetInfoParams `json:"params,omitempty"`              // Template params/values to be provisioned
	ResourceParams      []string                                                     `json:"resourceParams,omitempty"`      // Resource params to be provisioned. Refer to features page for usage details
	Type                string                                                       `json:"type,omitempty"`                // Target type of device
	VersionedTemplateID string                                                       `json:"versionedTemplateId,omitempty"` // Versioned templateUUID to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateTargetInfoParams interface{}
type RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria []RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria // Array of RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria
type RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria interface{}
type RequestConfigurationTemplatesPreviewTemplate struct {
	DeviceID       string                                                      `json:"deviceId,omitempty"`       // UUID of device to get template preview
	Params         *RequestConfigurationTemplatesPreviewTemplateParams         `json:"params,omitempty"`         // Params to render preview
	ResourceParams *RequestConfigurationTemplatesPreviewTemplateResourceParams `json:"resourceParams,omitempty"` // Resource params to render preview
	TemplateID     string                                                      `json:"templateId,omitempty"`     // UUID of template to get template preview
}
type RequestConfigurationTemplatesPreviewTemplateParams interface{}
type RequestConfigurationTemplatesPreviewTemplateResourceParams interface{}
type RequestConfigurationTemplatesVersionTemplate struct {
	Comments   string `json:"comments,omitempty"`   // Template version comments
	TemplateID string `json:"templateId,omitempty"` // UUID of template
}
type RequestConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate struct {
	ProfileID string `json:"profileId,omitempty"` // The id of the network profile, retrievable from `/intent/api/v1/networkProfilesForSites`
}
type RequestConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplate struct {
	Items *[][]string `json:"items,omitempty"` // Root
}
type RequestConfigurationTemplatesCommitTemplateForANewVersion struct {
	CommitNote string `json:"commitNote,omitempty"` // A message to leave as a note with the commit of a template. The maximum length allowed is 255 characters.
}
type RequestConfigurationTemplatesDeployTemplateV2 struct {
	ForcePushTemplate            *bool                                                      `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  *bool                                                      `json:"isComposite,omitempty"`                  // Composite template flag
	MainTemplateID               string                                                     `json:"mainTemplateId,omitempty"`               // Main template UUID of versioned template
	MemberTemplateDeploymentInfo []string                                                   `json:"memberTemplateDeploymentInfo,omitempty"` // memberTemplateDeploymentInfo
	TargetInfo                   *[]RequestConfigurationTemplatesDeployTemplateV2TargetInfo `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                                                     `json:"templateId,omitempty"`                   // UUID of template to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfo struct {
	HostName            string                                                         `json:"hostName,omitempty"`            // Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
	ID                  string                                                         `json:"id,omitempty"`                  // UUID of target is required if targetType is MANAGED_DEVICE_UUID
	Params              *RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams `json:"params,omitempty"`              // Template params/values to be provisioned
	ResourceParams      []string                                                       `json:"resourceParams,omitempty"`      // Resource params to be provisioned. Refer to features page for usage details
	Type                string                                                         `json:"type,omitempty"`                // Target type of device
	VersionedTemplateID string                                                         `json:"versionedTemplateId,omitempty"` // Versioned templateUUID to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams map[string]interface{}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams interface{}

//GetTemplateProjects Get Template Projects - 03b6-793d-45e8-9ff1
/* Get all matching template projects based on the filters selected.


@param GetTemplateProjectsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-template-projects
*/
func (s *ConfigurationTemplatesService) GetTemplateProjects(GetTemplateProjectsQueryParams *GetTemplateProjectsQueryParams) (*ResponseConfigurationTemplatesGetTemplateProjects, *resty.Response, error) {
	path := "/dna/intent/api/v1/projects"

	queryString, _ := query.Values(GetTemplateProjectsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplateProjects{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplateProjects(GetTemplateProjectsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplateProjects")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplateProjects)
	return result, response, err

}

//GetTemplateProjectCount Get Template Project Count - d394-ab82-44f8-9435
/* Get the count of all template projects.


@param GetTemplateProjectCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-template-project-count
*/
func (s *ConfigurationTemplatesService) GetTemplateProjectCount(GetTemplateProjectCountQueryParams *GetTemplateProjectCountQueryParams) (*ResponseConfigurationTemplatesGetTemplateProjectCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/projects/count"

	queryString, _ := query.Values(GetTemplateProjectCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplateProjectCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplateProjectCount(GetTemplateProjectCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplateProjectCount")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplateProjectCount)
	return result, response, err

}

//GetTemplateProject Get Template Project - 90b7-68c6-41d8-a1bd
/* Get a template project by the project's ID.


@param projectID projectId path parameter. The id of the project to get, retrieveable from `GET /dna/intent/api/v1/projects`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-template-project
*/
func (s *ConfigurationTemplatesService) GetTemplateProject(projectID string) (*ResponseConfigurationTemplatesGetTemplateProject, *resty.Response, error) {
	path := "/dna/intent/api/v1/projects/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesGetTemplateProject{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplateProject(projectID)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplateProject)
	return result, response, err

}

//GetsAListOfProjects Gets a list of projects - 4f80-08c2-400b-98ee
/* List the projects


@param GetsAListOfProjectsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-a-list-of-projects
*/
func (s *ConfigurationTemplatesService) GetsAListOfProjects(GetsAListOfProjectsQueryParams *GetsAListOfProjectsQueryParams) (*ResponseConfigurationTemplatesGetsAListOfProjects, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project"

	queryString, _ := query.Values(GetsAListOfProjectsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsAListOfProjects{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAListOfProjects(GetsAListOfProjectsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsAListOfProjects")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsAListOfProjects)
	return result, response, err

}

//GetsTheDetailsOfAGivenProject Gets the details of a given project. - dd91-a8c0-436a-82d9
/* Get the details of the given project by its id.


@param projectID projectId path parameter. projectId(UUID) of project to get project details


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-details-of-a-given-project
*/
func (s *ConfigurationTemplatesService) GetsTheDetailsOfAGivenProject(projectID string) (*ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheDetailsOfAGivenProject(projectID)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheDetailsOfAGivenProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject)
	return result, response, err

}

//GetsTheTemplatesAvailable Gets the templates available - e286-e848-47bb-a77e
/* List the templates available


@param GetsTheTemplatesAvailableQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-templates-available
*/
func (s *ConfigurationTemplatesService) GetsTheTemplatesAvailable(GetsTheTemplatesAvailableQueryParams *GetsTheTemplatesAvailableQueryParams) (*ResponseConfigurationTemplatesGetsTheTemplatesAvailable, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template"

	queryString, _ := query.Values(GetsTheTemplatesAvailableQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsTheTemplatesAvailable{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheTemplatesAvailable(GetsTheTemplatesAvailableQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheTemplatesAvailable")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsTheTemplatesAvailable)
	return result, response, err

}

//StatusOfTemplateDeployment Status of template deployment - 078e-f800-49b8-80f1
/* API to retrieve the status of template deployment.


@param deploymentID deploymentId path parameter. UUID of deployment to retrieve template deployment status


Documentation Link: https://developer.cisco.com/docs/dna-center/#!status-of-template-deployment
*/
func (s *ConfigurationTemplatesService) StatusOfTemplateDeployment(deploymentID string) (*ResponseConfigurationTemplatesStatusOfTemplateDeployment, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/deploy/status/{deploymentId}"
	path = strings.Replace(path, "{deploymentId}", fmt.Sprintf("%v", deploymentID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesStatusOfTemplateDeployment{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.StatusOfTemplateDeployment(deploymentID)
		}
		return nil, response, fmt.Errorf("error with operation StatusOfTemplateDeployment")
	}

	result := response.Result().(*ResponseConfigurationTemplatesStatusOfTemplateDeployment)
	return result, response, err

}

//GetsAllTheVersionsOfAGivenTemplate Gets all the versions of a given template - 0191-2b65-45fb-8891
/* Get all the versions of template by its id


@param templateID templateId path parameter. templateId(UUID) to get list of versioned templates


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-all-the-versions-of-a-given-template
*/
func (s *ConfigurationTemplatesService) GetsAllTheVersionsOfAGivenTemplate(templateID string) (*ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/version/{templateId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsAllTheVersionsOfAGivenTemplate(templateID)
		}
		return nil, response, fmt.Errorf("error with operation GetsAllTheVersionsOfAGivenTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate)
	return result, response, err

}

//GetsDetailsOfAGivenTemplate Gets details of a given template - 8c82-5900-40d9-8137
/* Details of the template by its id


@param templateID templateId path parameter. TemplateId(UUID) to get details of the template

@param GetsDetailsOfAGivenTemplateQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-details-of-a-given-template
*/
func (s *ConfigurationTemplatesService) GetsDetailsOfAGivenTemplate(templateID string, GetsDetailsOfAGivenTemplateQueryParams *GetsDetailsOfAGivenTemplateQueryParams) (*ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	queryString, _ := query.Values(GetsDetailsOfAGivenTemplateQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsDetailsOfAGivenTemplate(templateID, GetsDetailsOfAGivenTemplateQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetsDetailsOfAGivenTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate)
	return result, response, err

}

//RetrieveTheNetworkProfilesAttachedToACLITemplate Retrieve the network profiles attached to a CLI template - 7887-b9c5-40d9-b707
/* Retrieves the list of network profiles that a CLI template is currently attached to by the template ID.


@param templateID templateId path parameter. The `id` of the template, retrievable from `GET /intent/api/v1/templates`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-the-network-profiles-attached-to-acl-i-template
*/
func (s *ConfigurationTemplatesService) RetrieveTheNetworkProfilesAttachedToACLITemplate(templateID string) (*ResponseConfigurationTemplatesRetrieveTheNetworkProfilesAttachedToACLITemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/networkProfilesForSites"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesRetrieveTheNetworkProfilesAttachedToACLITemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveTheNetworkProfilesAttachedToACLITemplate(templateID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveTheNetworkProfilesAttachedToAclITemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesRetrieveTheNetworkProfilesAttachedToACLITemplate)
	return result, response, err

}

//RetrieveCountOfNetworkProfilesAttachedToACLITemplate Retrieve count of network profiles attached to a CLI template - f1ab-98e1-426a-9a06
/* Retrieves the count of network profiles that a CLI template has been attached to by the template ID.


@param templateID templateId path parameter. The `id` of the template, retrievable from `GET /intent/api/v1/templates`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-count-of-network-profiles-attached-to-acl-i-template
*/
func (s *ConfigurationTemplatesService) RetrieveCountOfNetworkProfilesAttachedToACLITemplate(templateID string) (*ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/networkProfilesForSites/count"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplate{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveCountOfNetworkProfilesAttachedToACLITemplate(templateID)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveCountOfNetworkProfilesAttachedToAclITemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesRetrieveCountOfNetworkProfilesAttachedToACLITemplate)
	return result, response, err

}

//GetTemplateVersions Get Template Versions - b08a-5b15-4c99-beb3
/* Get a template's version information.


@param templateID templateId path parameter. The id of the template to get versions of, retrieveable from `GET /dna/intent/api/v1/templates`

@param GetTemplateVersionsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-template-versions
*/
func (s *ConfigurationTemplatesService) GetTemplateVersions(templateID string, GetTemplateVersionsQueryParams *GetTemplateVersionsQueryParams) (*ResponseConfigurationTemplatesGetTemplateVersions, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/versions"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	queryString, _ := query.Values(GetTemplateVersionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplateVersions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplateVersions(templateID, GetTemplateVersionsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplateVersions")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplateVersions)
	return result, response, err

}

//GetTemplateVersionsCount Get Template Versions Count - 57a5-d893-4d98-a66f
/* Get the count of a template's version information.


@param templateID templateId path parameter. The id of the template to get versions of, retrieveable from `GET /dna/intent/api/v1/templates`

@param GetTemplateVersionsCountQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-template-versions-count
*/
func (s *ConfigurationTemplatesService) GetTemplateVersionsCount(templateID string, GetTemplateVersionsCountQueryParams *GetTemplateVersionsCountQueryParams) (*ResponseConfigurationTemplatesGetTemplateVersionsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/versions/count"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	queryString, _ := query.Values(GetTemplateVersionsCountQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplateVersionsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplateVersionsCount(templateID, GetTemplateVersionsCountQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplateVersionsCount")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplateVersionsCount)
	return result, response, err

}

//GetTemplateVersion Get Template Version - 689b-e88b-4c18-8a8c
/* Get a template's version by the version ID.


@param templateID templateId path parameter. The id of the template to get versions of, retrieveable from `GET /dna/intent/api/v1/templates`

@param versionID versionId path parameter. The id of the versioned template to get versions of, retrieveable from `GET /dna/intent/api/v1/templates/{id}/versions`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-template-version
*/
func (s *ConfigurationTemplatesService) GetTemplateVersion(templateID string, versionID string) (*ResponseConfigurationTemplatesGetTemplateVersion, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/versions/{versionId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)
	path = strings.Replace(path, "{versionId}", fmt.Sprintf("%v", versionID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesGetTemplateVersion{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplateVersion(templateID, versionID)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplateVersion")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplateVersion)
	return result, response, err

}

//GetProjectsDetailsV2 Get project(s) details - 9a8c-aa6d-459b-a4a2
/* Get project(s) details


@param GetProjectsDetailsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-projects-details
*/
func (s *ConfigurationTemplatesService) GetProjectsDetailsV2(GetProjectsDetailsV2QueryParams *GetProjectsDetailsV2QueryParams) (*ResponseConfigurationTemplatesGetProjectsDetailsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/project"

	queryString, _ := query.Values(GetProjectsDetailsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetProjectsDetailsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetProjectsDetailsV2(GetProjectsDetailsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetProjectsDetailsV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetProjectsDetailsV2)
	return result, response, err

}

//GetTemplatesDetailsV2 Get template(s) details - b0b6-ba49-43c8-9f45
/* Get template(s) details


@param GetTemplatesDetailsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-templates-details-v2
*/
func (s *ConfigurationTemplatesService) GetTemplatesDetailsV2(GetTemplatesDetailsV2QueryParams *GetTemplatesDetailsV2QueryParams) (*ResponseConfigurationTemplatesGetTemplatesDetailsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/template"

	queryString, _ := query.Values(GetTemplatesDetailsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplatesDetailsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTemplatesDetailsV2(GetTemplatesDetailsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetTemplatesDetailsV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplatesDetailsV2)
	return result, response, err

}

//CreateTemplateProject Create Template Project - a280-e91c-498b-a571
/* Create a template project.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-template-project
*/
func (s *ConfigurationTemplatesService) CreateTemplateProject(requestConfigurationTemplatesCreateTemplateProject *RequestConfigurationTemplatesCreateTemplateProject) (*ResponseConfigurationTemplatesCreateTemplateProject, *resty.Response, error) {
	path := "/dna/intent/api/v1/projects"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesCreateTemplateProject).
		SetResult(&ResponseConfigurationTemplatesCreateTemplateProject{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateTemplateProject(requestConfigurationTemplatesCreateTemplateProject)
		}

		return nil, response, fmt.Errorf("error with operation CreateTemplateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateTemplateProject)
	return result, response, err

}

//CreatesACloneOfTheGivenTemplate Creates a clone of the given template - 0384-4a0a-4ee8-bfc2
/* API to clone template


@param name name path parameter. Template name to clone template(Name should be different than existing template name within same project)

@param templateID templateId path parameter. UUID of the template to clone it

@param projectID projectId path parameter.
@param CreatesACloneOfTheGivenTemplateQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!creates-a-clone-of-the-given-template
*/
func (s *ConfigurationTemplatesService) CreatesACloneOfTheGivenTemplate(name string, templateID string, projectID string, CreatesACloneOfTheGivenTemplateQueryParams *CreatesACloneOfTheGivenTemplateQueryParams) (*ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/clone/name/{name}/project/{projectId}/template/{templateId}"
	path = strings.Replace(path, "{name}", fmt.Sprintf("%v", name), -1)
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	queryString, _ := query.Values(CreatesACloneOfTheGivenTemplateQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreatesACloneOfTheGivenTemplate(name, templateID, projectID, CreatesACloneOfTheGivenTemplateQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation CreatesACloneOfTheGivenTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplate)
	return result, response, err

}

//CreateProject Create Project - 5788-a8a8-4aa9-b97a
/* This API is used to create a new project.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-project
*/
func (s *ConfigurationTemplatesService) CreateProject(requestConfigurationTemplatesCreateProject *RequestConfigurationTemplatesCreateProject) (*ResponseConfigurationTemplatesCreateProject, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesCreateProject).
		SetResult(&ResponseConfigurationTemplatesCreateProject{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateProject(requestConfigurationTemplatesCreateProject)
		}

		return nil, response, fmt.Errorf("error with operation CreateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateProject)
	return result, response, err

}

//ImportsTheProjectsProvided Imports the Projects provided - f59f-a8ad-42d9-8b4f
/* Imports the Projects provided in the DTO


@param ImportsTheProjectsProvidedQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!imports-the-projects-provided
*/
func (s *ConfigurationTemplatesService) ImportsTheProjectsProvided(ImportsTheProjectsProvidedQueryParams *ImportsTheProjectsProvidedQueryParams) (*ResponseConfigurationTemplatesImportsTheProjectsProvided, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/importprojects"

	queryString, _ := query.Values(ImportsTheProjectsProvidedQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseConfigurationTemplatesImportsTheProjectsProvided{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportsTheProjectsProvided(ImportsTheProjectsProvidedQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ImportsTheProjectsProvided")
	}

	result := response.Result().(*ResponseConfigurationTemplatesImportsTheProjectsProvided)
	return result, response, err

}

//ExportsTheProjectsForAGivenCriteria Exports the projects for a given criteria. - 67bc-e964-45f8-b720
/* Exports the projects for given projectNames.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!exports-the-projects-for-a-given-criteria
*/
func (s *ConfigurationTemplatesService) ExportsTheProjectsForAGivenCriteria(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteria *RequestConfigurationTemplatesExportsTheProjectsForAGivenCriteria) (*ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteria, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/name/exportprojects"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteria).
		SetResult(&ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteria{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportsTheProjectsForAGivenCriteria(requestConfigurationTemplatesExportsTheProjectsForAGivenCriteria)
		}

		return nil, response, fmt.Errorf("error with operation ExportsTheProjectsForAGivenCriteria")
	}

	result := response.Result().(*ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteria)
	return result, response, err

}

//ImportsTheTemplatesProvided Imports the templates provided - 4d86-f92a-4a7b-90bb
/* Imports the templates provided in the DTO by project Name


@param projectName projectName path parameter. Project name to create template under the project

@param ImportsTheTemplatesProvidedQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!imports-the-templates-provided
*/
func (s *ConfigurationTemplatesService) ImportsTheTemplatesProvided(projectName string, requestConfigurationTemplatesImportsTheTemplatesProvided *RequestConfigurationTemplatesImportsTheTemplatesProvided, ImportsTheTemplatesProvidedQueryParams *ImportsTheTemplatesProvidedQueryParams) (*ResponseConfigurationTemplatesImportsTheTemplatesProvided, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/name/{projectName}/template/importtemplates"
	path = strings.Replace(path, "{projectName}", fmt.Sprintf("%v", projectName), -1)

	queryString, _ := query.Values(ImportsTheTemplatesProvidedQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestConfigurationTemplatesImportsTheTemplatesProvided).
		SetResult(&ResponseConfigurationTemplatesImportsTheTemplatesProvided{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportsTheTemplatesProvided(projectName, requestConfigurationTemplatesImportsTheTemplatesProvided, ImportsTheTemplatesProvidedQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ImportsTheTemplatesProvided")
	}

	result := response.Result().(*ResponseConfigurationTemplatesImportsTheTemplatesProvided)
	return result, response, err

}

//CreateTemplate Create Template - ab94-a88b-4b0b-8d3d
/* API to create a template by project id.


@param projectID projectId path parameter. UUID of the project in which the template needs to be created


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-template
*/
func (s *ConfigurationTemplatesService) CreateTemplate(projectID string, requestConfigurationTemplatesCreateTemplate *RequestConfigurationTemplatesCreateTemplate) (*ResponseConfigurationTemplatesCreateTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project/{projectId}/template"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesCreateTemplate).
		SetResult(&ResponseConfigurationTemplatesCreateTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateTemplate(projectID, requestConfigurationTemplatesCreateTemplate)
		}

		return nil, response, fmt.Errorf("error with operation CreateTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateTemplate)
	return result, response, err

}

//DeployTemplate Deploy Template - 179f-09d8-430b-bee0
/* API to deploy a template.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-template
*/
func (s *ConfigurationTemplatesService) DeployTemplate(requestConfigurationTemplatesDeployTemplate *RequestConfigurationTemplatesDeployTemplate) (*ResponseConfigurationTemplatesDeployTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/deploy"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesDeployTemplate).
		SetResult(&ResponseConfigurationTemplatesDeployTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployTemplate(requestConfigurationTemplatesDeployTemplate)
		}

		return nil, response, fmt.Errorf("error with operation DeployTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeployTemplate)
	return result, response, err

}

//ExportsTheTemplatesForAGivenCriteria Exports the templates for a given criteria. - a3a9-498f-4f48-a3c7
/* Exports the templates for given templateIds.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!exports-the-templates-for-a-given-criteria
*/
func (s *ConfigurationTemplatesService) ExportsTheTemplatesForAGivenCriteria(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria *RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria) (*ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteria, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/exporttemplates"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria).
		SetResult(&ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteria{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportsTheTemplatesForAGivenCriteria(requestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria)
		}

		return nil, response, fmt.Errorf("error with operation ExportsTheTemplatesForAGivenCriteria")
	}

	result := response.Result().(*ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteria)
	return result, response, err

}

//VersionTemplate Version Template - f2a4-4a7e-4d5b-ab78
/* API to version the current contents of the template.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!version-template
*/
func (s *ConfigurationTemplatesService) VersionTemplate(requestConfigurationTemplatesVersionTemplate *RequestConfigurationTemplatesVersionTemplate) (*ResponseConfigurationTemplatesVersionTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/version"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesVersionTemplate).
		SetResult(&ResponseConfigurationTemplatesVersionTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.VersionTemplate(requestConfigurationTemplatesVersionTemplate)
		}

		return nil, response, fmt.Errorf("error with operation VersionTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesVersionTemplate)
	return result, response, err

}

//AttachNetworkProfileToADayNCliTemplate Attach network profile to a Day-N CLI template - ecba-b925-4ed9-b235
/* Attaches a network profile to a Day-N CLI template by passing the profile ID and template ID.


@param templateID templateId path parameter. The `id` of the template, retrievable from `GET /intent/api/v1/templates`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!attach-network-profile-to-a-day-n-cli-template
*/
func (s *ConfigurationTemplatesService) AttachNetworkProfileToADayNCliTemplate(templateID string, requestConfigurationTemplatesAttachNetworkProfileToADayNCLITemplate *RequestConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate) (*ResponseConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/networkProfilesForSites"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesAttachNetworkProfileToADayNCLITemplate).
		SetResult(&ResponseConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AttachNetworkProfileToADayNCliTemplate(templateID, requestConfigurationTemplatesAttachNetworkProfileToADayNCLITemplate)
		}

		return nil, response, fmt.Errorf("error with operation AttachNetworkProfileToADayNCliTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesAttachNetworkProfileToADayNCliTemplate)
	return result, response, err

}

//AttachAListOfNetworkProfilesToADayNCliTemplate Attach a list of network profiles to a Day-N CLI template - 26a5-cb41-46c8-8fe4
/* Attaches a list of network profiles to the Day-N CLI template by passing the profile IDs and template ID.


@param templateID templateId path parameter. The `id` of the template, retrievable from `GET /intent/api/v1/templates`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!attach-a-list-of-network-profiles-to-a-day-n-cli-template
*/
func (s *ConfigurationTemplatesService) AttachAListOfNetworkProfilesToADayNCliTemplate(templateID string, requestConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCLITemplate *RequestConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplate) (*ResponseConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/networkProfilesForSites/bulk"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCLITemplate).
		SetResult(&ResponseConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AttachAListOfNetworkProfilesToADayNCliTemplate(templateID, requestConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCLITemplate)
		}

		return nil, response, fmt.Errorf("error with operation AttachAListOfNetworkProfilesToADayNCliTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesAttachAListOfNetworkProfilesToADayNCliTemplate)
	return result, response, err

}

//CommitTemplateForANewVersion Commit Template For a New Version - c8bd-49fa-4999-b43b
/* Transitions the current draft of a template to a new committed version with a higher version number.


@param templateID templateId path parameter. The id of the template to commit a new version for, retrieveable from `GET /dna/intent/api/v1/templates`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!commit-template-for-a-new-version
*/
func (s *ConfigurationTemplatesService) CommitTemplateForANewVersion(templateID string, requestConfigurationTemplatesCommitTemplateForANewVersion *RequestConfigurationTemplatesCommitTemplateForANewVersion) (*ResponseConfigurationTemplatesCommitTemplateForANewVersion, *resty.Response, error) {
	path := "/dna/intent/api/v1/templates/{templateId}/versions/commit"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesCommitTemplateForANewVersion).
		SetResult(&ResponseConfigurationTemplatesCommitTemplateForANewVersion{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CommitTemplateForANewVersion(templateID, requestConfigurationTemplatesCommitTemplateForANewVersion)
		}

		return nil, response, fmt.Errorf("error with operation CommitTemplateForANewVersion")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCommitTemplateForANewVersion)
	return result, response, err

}

//DeployTemplateV2 Deploy Template V2 - 02af-1bdf-4b48-9cbb
/* V2 API to deploy a template.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-template-v2
*/
func (s *ConfigurationTemplatesService) DeployTemplateV2(requestConfigurationTemplatesDeployTemplateV2 *RequestConfigurationTemplatesDeployTemplateV2) (*ResponseConfigurationTemplatesDeployTemplateV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/template/deploy"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesDeployTemplateV2).
		SetResult(&ResponseConfigurationTemplatesDeployTemplateV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployTemplateV2(requestConfigurationTemplatesDeployTemplateV2)
		}

		return nil, response, fmt.Errorf("error with operation DeployTemplateV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeployTemplateV2)
	return result, response, err

}

//UpdateTemplateProject Update Template Project - 5891-e9e3-464a-8787
/* Update a template project by the project's ID.


@param projectID projectId path parameter. The id of the project to update, retrieveable from `GET /dna/intent/api/v1/projects`

*/
func (s *ConfigurationTemplatesService) UpdateTemplateProject(projectID string, requestConfigurationTemplatesUpdateTemplateProject *RequestConfigurationTemplatesUpdateTemplateProject) (*ResponseConfigurationTemplatesUpdateTemplateProject, *resty.Response, error) {
	path := "/dna/intent/api/v1/projects/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesUpdateTemplateProject).
		SetResult(&ResponseConfigurationTemplatesUpdateTemplateProject{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTemplateProject(projectID, requestConfigurationTemplatesUpdateTemplateProject)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTemplateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesUpdateTemplateProject)
	return result, response, err

}

//UpdateProject Update Project - ecb8-8b89-4318-ac8d
/* This API is used to update an existing project.


 */
func (s *ConfigurationTemplatesService) UpdateProject(requestConfigurationTemplatesUpdateProject *RequestConfigurationTemplatesUpdateProject) (*ResponseConfigurationTemplatesUpdateProject, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/project"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesUpdateProject).
		SetResult(&ResponseConfigurationTemplatesUpdateProject{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateProject(requestConfigurationTemplatesUpdateProject)
		}
		return nil, response, fmt.Errorf("error with operation UpdateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesUpdateProject)
	return result, response, err

}

//UpdateTemplate Update Template - 2a80-39f5-4aab-86be
/* API to update a template.


 */
func (s *ConfigurationTemplatesService) UpdateTemplate(requestConfigurationTemplatesUpdateTemplate *RequestConfigurationTemplatesUpdateTemplate) (*ResponseConfigurationTemplatesUpdateTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesUpdateTemplate).
		SetResult(&ResponseConfigurationTemplatesUpdateTemplate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTemplate(requestConfigurationTemplatesUpdateTemplate)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesUpdateTemplate)
	return result, response, err

}

//PreviewTemplate Preview Template - 41bc-aaa6-4669-853e
/* API to preview a template.


 */
func (s *ConfigurationTemplatesService) PreviewTemplate(requestConfigurationTemplatesPreviewTemplate *RequestConfigurationTemplatesPreviewTemplate) (*ResponseConfigurationTemplatesPreviewTemplate, *resty.Response, error) {
	path := "/dna/intent/api/v1/template-programmer/template/preview"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationTemplatesPreviewTemplate).
		SetResult(&ResponseConfigurationTemplatesPreviewTemplate{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.PreviewTemplate(requestConfigurationTemplatesPreviewTemplate)
		}
		return nil, response, fmt.Errorf("error with operation PreviewTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesPreviewTemplate)
	return result, response, err

}

//DeleteTemplateProject Delete Template Project - 9c8e-baae-413b-9dbf
/* Delete a template project by the project's ID.


@param projectID projectId path parameter. the id of the project to delete, retrieveable from `GET /dna/intent/api/v1/projects`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-template-project
*/
func (s *ConfigurationTemplatesService) DeleteTemplateProject(projectID string) (*ResponseConfigurationTemplatesDeleteTemplateProject, *resty.Response, error) {
	//projectID string
	path := "/dna/intent/api/v1/projects/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesDeleteTemplateProject{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTemplateProject(projectID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTemplateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeleteTemplateProject)
	return result, response, err

}

//DeletesTheProject Deletes the project - 8cbb-79f4-4259-82d4
/* Deletes the project by its id


@param projectID projectId path parameter. projectId(UUID) of project to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-the-project
*/
func (s *ConfigurationTemplatesService) DeletesTheProject(projectID string) (*ResponseConfigurationTemplatesDeletesTheProject, *resty.Response, error) {
	//projectID string
	path := "/dna/intent/api/v1/template-programmer/project/{projectId}"
	path = strings.Replace(path, "{projectId}", fmt.Sprintf("%v", projectID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesDeletesTheProject{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesTheProject(projectID)
		}
		return nil, response, fmt.Errorf("error with operation DeletesTheProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeletesTheProject)
	return result, response, err

}

//DeletesTheTemplate Deletes the template - a2bb-4afd-4699-8965
/* Deletes the template by its id


@param templateID templateId path parameter. templateId(UUID) of template to be deleted


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deletes-the-template
*/
func (s *ConfigurationTemplatesService) DeletesTheTemplate(templateID string) (*ResponseConfigurationTemplatesDeletesTheTemplate, *resty.Response, error) {
	//templateID string
	path := "/dna/intent/api/v1/template-programmer/template/{templateId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesDeletesTheTemplate{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeletesTheTemplate(templateID)
		}
		return nil, response, fmt.Errorf("error with operation DeletesTheTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeletesTheTemplate)
	return result, response, err

}

//DetachAListOfNetworkProfilesFromADayNCliTemplate Detach a list of network profiles from a Day-N CLI template - a487-0999-4c7a-a94d
/* Detach a list of network profiles from a Day-N CLI template with a list of profile IDs along with the template ID.


@param templateID templateId path parameter. The `id` of the template, retrievable from `GET /intent/api/v1/templates`

@param DetachAListOfNetworkProfilesFromADayNCLITemplateQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!detach-a-list-of-network-profiles-from-a-day-n-cli-template
*/
func (s *ConfigurationTemplatesService) DetachAListOfNetworkProfilesFromADayNCliTemplate(templateID string, DetachAListOfNetworkProfilesFromADayNCLITemplateQueryParams *DetachAListOfNetworkProfilesFromADayNCliTemplateQueryParams) (*ResponseConfigurationTemplatesDetachAListOfNetworkProfilesFromADayNCliTemplate, *resty.Response, error) {
	//templateID string,DetachAListOfNetworkProfilesFromADayNCLITemplateQueryParams *DetachAListOfNetworkProfilesFromADayNCliTemplateQueryParams
	path := "/dna/intent/api/v1/templates/{templateId}/networkProfilesForSites/bulk"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)

	queryString, _ := query.Values(DetachAListOfNetworkProfilesFromADayNCLITemplateQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesDetachAListOfNetworkProfilesFromADayNCliTemplate{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DetachAListOfNetworkProfilesFromADayNCliTemplate(templateID, DetachAListOfNetworkProfilesFromADayNCLITemplateQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DetachAListOfNetworkProfilesFromADayNCliTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDetachAListOfNetworkProfilesFromADayNCliTemplate)
	return result, response, err

}

//DetachANetworkProfileFromADayNCliTemplate Detach a network profile from a Day-N CLI template - d98b-2986-4d5b-b04c
/* Detach a network profile from a Day-N CLI template by the profile ID and template ID.


@param templateID templateId path parameter. The `id` of the template, retrievable from `GET /intent/api/v1/templates`

@param profileID profileId path parameter. The `id` of the network profile, retrievable from `GET /intent/api/v1/networkProfilesForSites`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!detach-a-network-profile-from-a-day-n-cli-template
*/
func (s *ConfigurationTemplatesService) DetachANetworkProfileFromADayNCliTemplate(templateID string, profileID string) (*ResponseConfigurationTemplatesDetachANetworkProfileFromADayNCliTemplate, *resty.Response, error) {
	//templateID string,profileID string
	path := "/dna/intent/api/v1/templates/{templateId}/networkProfilesForSites/{profileId}"
	path = strings.Replace(path, "{templateId}", fmt.Sprintf("%v", templateID), -1)
	path = strings.Replace(path, "{profileId}", fmt.Sprintf("%v", profileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationTemplatesDetachANetworkProfileFromADayNCliTemplate{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DetachANetworkProfileFromADayNCliTemplate(templateID, profileID)
		}
		return nil, response, fmt.Errorf("error with operation DetachANetworkProfileFromADayNCliTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDetachANetworkProfileFromADayNCliTemplate)
	return result, response, err

}
