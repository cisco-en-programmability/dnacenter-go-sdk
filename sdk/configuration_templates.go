package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ConfigurationTemplatesService service

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
type GetProjectsDetailsQueryParams struct {
	ID        string `url:"id,omitempty"`        //Id of project to be searched
	Name      string `url:"name,omitempty"`      //Name of project to be searched
	Offset    int    `url:"offset,omitempty"`    //Index of first result
	Limit     int    `url:"limit,omitempty"`     //Limits number of results
	SortOrder string `url:"sortOrder,omitempty"` //Sort Order Ascending (asc) or Descending (dsc)
}
type GetTemplatesDetailsQueryParams struct {
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
	Limit                      int      `url:"limit,omitempty"`                      //Limits number of results
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
	RollbackTemplateErrors *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                  `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                  `json:"templateVersion,omitempty"`        // Current version of template
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
	IsDeletable    *bool                                                                   `json:"isDeletable,omitempty"`    // Is deletable
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
	RollbackTemplateErrors *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                        `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                        `json:"templateVersion,omitempty"`        // Current version of template
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
	RollbackTemplateErrors *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                             `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                             `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesGetProjectsDetails struct {
	Response *[]ResponseConfigurationTemplatesGetProjectsDetailsResponse `json:"response,omitempty"` //
	Version  string                                                      `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponse struct {
	CreateTime     *int                                                                 `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                               `json:"description,omitempty"`    // Description of project
	ID             string                                                               `json:"id,omitempty"`             // UUID of project
	IsDeletable    *bool                                                                `json:"isDeletable,omitempty"`    // Flag to check if project is deletable or not(for internal use only)
	LastUpdateTime *int                                                                 `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                               `json:"name,omitempty"`           // Name of project
	Tags           *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTags      `json:"tags,omitempty"`           //
	Templates      *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplates struct {
	Tags                    *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                                     `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                                      `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                                       `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                                      `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                                     `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                                     `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                                     `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                                     `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                                       `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                                       `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                                     `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                                     `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                                     `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                                     `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                                     `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                                     `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                                     `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                                     `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                                     `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                                     `json:"version,omitempty"`                 // Current version of template
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplates struct {
	Tags                   *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                                         `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                        `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                        `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                        `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                        `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                        `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                                        `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                        `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                                               `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                                                 `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                                               `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                                               `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                                               `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                                               `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                                               `json:"group,omitempty"`           // group
	ID              string                                                                                                               `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                                               `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                                               `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                                                `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                                                 `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                                                `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                                               `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                                               `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParams struct {
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
	Range           *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParams struct {
	Binding         string                                                                                            `json:"binding,omitempty"`         // Bind to source
	CustomOrder     *int                                                                                              `json:"customOrder,omitempty"`     // CustomOrder of template param
	DataType        string                                                                                            `json:"dataType,omitempty"`        // Datatype of template param
	DefaultValue    string                                                                                            `json:"defaultValue,omitempty"`    // Default value of template param
	Description     string                                                                                            `json:"description,omitempty"`     // Description of template param
	DisplayName     string                                                                                            `json:"displayName,omitempty"`     // Display name of param
	Group           string                                                                                            `json:"group,omitempty"`           // group
	ID              string                                                                                            `json:"id,omitempty"`              // UUID of template param
	InstructionText string                                                                                            `json:"instructionText,omitempty"` // Instruction text for param
	Key             string                                                                                            `json:"key,omitempty"`             // key
	NotParam        *bool                                                                                             `json:"notParam,omitempty"`        // Is it not a variable
	Order           *int                                                                                              `json:"order,omitempty"`           // Order of template param
	ParamArray      *bool                                                                                             `json:"paramArray,omitempty"`      // Is it an array
	ParameterName   string                                                                                            `json:"parameterName,omitempty"`   // Name of template param
	Provider        string                                                                                            `json:"provider,omitempty"`        // provider
	Range           *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                             `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                         `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                           `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                           `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParams struct {
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
	Range           *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                     `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                 `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                   `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                   `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrors struct {
	RollbackTemplateErrors *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                     `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                     `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetails struct {
	Response *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponse `json:"response,omitempty"` //
	Version  string                                                       `json:"version,omitempty"`  //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponse struct {
	Author                  string                                                                             `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                              `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                               `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                              `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                             `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseDeviceTypes            `json:"deviceTypes,omitempty"`             //
	DocumentDatabase        *bool                                                                              `json:"documentDatabase,omitempty"`        // Document Database
	FailurePolicy           string                                                                             `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                             `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                             `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                               `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                               `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                             `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                             `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectAssociated       *bool                                                                              `json:"projectAssociated,omitempty"`       //
	ProjectID               string                                                                             `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                             `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                             `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                             `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                             `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                             `json:"softwareVersion,omitempty"`         // Applicable device software version
	Tags                    *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseTags                   `json:"tags,omitempty"`                    //
	TemplateContent         string                                                                             `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                             `json:"version,omitempty"`                 // Current version of template
	VersionsInfo            *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseVersionsInfo           `json:"versionsInfo,omitempty"`            //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplates struct {
	Composite              *bool                                                                                                 `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                                `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                                `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                                `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                                `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                                `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	Tags                   *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	TemplateContent        string                                                                                                `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                                `json:"version,omitempty"`                // Current version of template
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParams struct {
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
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                        `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                                    `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                      `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                      `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParams struct {
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
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                            `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                              `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                              `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParams struct {
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
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                     `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                 `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                   `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                   `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseRollbackTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParams struct {
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
	Range           *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                             `json:"required,omitempty"`        // Is param required
	Selection       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsSelection `json:"selection,omitempty"`       //
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                         `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                           `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                           `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseTemplateParamsSelectionSelectionValues interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrors struct {
	RollbackTemplateErrors *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                             `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                             `json:"templateVersion,omitempty"`        // Current version of template
}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrorsRollbackTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrorsTemplateErrors interface{}
type ResponseConfigurationTemplatesGetTemplatesDetailsResponseVersionsInfo struct {
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
type RequestConfigurationTemplatesCreateProject struct {
	Tags           *[]RequestConfigurationTemplatesCreateProjectTags      `json:"tags,omitempty"`           //
	CreateTime     *int                                                   `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                 `json:"description,omitempty"`    // Description of project
	ID             string                                                 `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                   `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                 `json:"name,omitempty"`           // Name of project
	Templates      *[]RequestConfigurationTemplatesCreateProjectTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type RequestConfigurationTemplatesCreateProjectTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateProjectTemplates struct {
	Tags                    *[]RequestConfigurationTemplatesCreateProjectTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                       `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                        `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                         `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                        `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                       `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                       `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                       `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                       `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                         `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                         `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                       `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                       `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                       `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                       `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                       `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                       `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                       `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                       `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                       `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                       `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesCreateProjectTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                           `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                          `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                          `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                          `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                          `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                          `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                          `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                          `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateProjectTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                               `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateProjectTemplatesTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesCreateProjectTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors struct {
	RollbackTemplateErrors *[]RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                       `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                       `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesCreateProjectTemplatesValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateProject struct {
	Tags           *[]RequestConfigurationTemplatesUpdateProjectTags      `json:"tags,omitempty"`           //
	CreateTime     *int                                                   `json:"createTime,omitempty"`     // Create time of project
	Description    string                                                 `json:"description,omitempty"`    // Description of project
	ID             string                                                 `json:"id,omitempty"`             // UUID of project
	LastUpdateTime *int                                                   `json:"lastUpdateTime,omitempty"` // Update time of project
	Name           string                                                 `json:"name,omitempty"`           // Name of project
	Templates      *[]RequestConfigurationTemplatesUpdateProjectTemplates `json:"templates,omitempty"`      // List of templates within the project
}
type RequestConfigurationTemplatesUpdateProjectTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateProjectTemplates struct {
	Tags                    *[]RequestConfigurationTemplatesUpdateProjectTemplatesTags                   `json:"tags,omitempty"`                    //
	Author                  string                                                                       `json:"author,omitempty"`                  // Author of template
	Composite               *bool                                                                        `json:"composite,omitempty"`               // Is it composite template
	ContainingTemplates     *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates    `json:"containingTemplates,omitempty"`     //
	CreateTime              *int                                                                         `json:"createTime,omitempty"`              // Create time of template
	CustomParamsOrder       *bool                                                                        `json:"customParamsOrder,omitempty"`       // Custom Params Order
	Description             string                                                                       `json:"description,omitempty"`             // Description of template
	DeviceTypes             *[]RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`             //
	FailurePolicy           string                                                                       `json:"failurePolicy,omitempty"`           // Define failure policy if template provisioning fails
	ID                      string                                                                       `json:"id,omitempty"`                      // UUID of template
	Language                string                                                                       `json:"language,omitempty"`                // Template language (JINJA or VELOCITY)
	LastUpdateTime          *int                                                                         `json:"lastUpdateTime,omitempty"`          // Update time of template
	LatestVersionTime       *int                                                                         `json:"latestVersionTime,omitempty"`       // Latest versioned template time
	Name                    string                                                                       `json:"name,omitempty"`                    // Name of template
	ParentTemplateID        string                                                                       `json:"parentTemplateId,omitempty"`        // Parent templateID
	ProjectID               string                                                                       `json:"projectId,omitempty"`               // Project UUID
	ProjectName             string                                                                       `json:"projectName,omitempty"`             // Project name
	RollbackTemplateContent string                                                                       `json:"rollbackTemplateContent,omitempty"` // Rollback template content
	RollbackTemplateParams  *[]RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"`  //
	SoftwareType            string                                                                       `json:"softwareType,omitempty"`            // Applicable device software type
	SoftwareVariant         string                                                                       `json:"softwareVariant,omitempty"`         // Applicable device software variant
	SoftwareVersion         string                                                                       `json:"softwareVersion,omitempty"`         // Applicable device software version
	TemplateContent         string                                                                       `json:"templateContent,omitempty"`         // Template content
	TemplateParams          *[]RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams         `json:"templateParams,omitempty"`          //
	ValidationErrors        *RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors         `json:"validationErrors,omitempty"`        //
	Version                 string                                                                       `json:"version,omitempty"`                 // Current version of template
}
type RequestConfigurationTemplatesUpdateProjectTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplates struct {
	Tags                   *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags                   `json:"tags,omitempty"`                   //
	Composite              *bool                                                                                           `json:"composite,omitempty"`              // Is it composite template
	Description            string                                                                                          `json:"description,omitempty"`            // Description of template
	DeviceTypes            *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes            `json:"deviceTypes,omitempty"`            //
	ID                     string                                                                                          `json:"id,omitempty"`                     // UUID of template
	Language               string                                                                                          `json:"language,omitempty"`               // Template language (JINJA or VELOCITY)
	Name                   string                                                                                          `json:"name,omitempty"`                   // Name of template
	ProjectName            string                                                                                          `json:"projectName,omitempty"`            // Project name
	RollbackTemplateParams *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams `json:"rollbackTemplateParams,omitempty"` //
	TemplateContent        string                                                                                          `json:"templateContent,omitempty"`        // Template content
	TemplateParams         *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams         `json:"templateParams,omitempty"`         //
	Version                string                                                                                          `json:"version,omitempty"`                // Current version of template
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTags struct {
	ID   string `json:"id,omitempty"`   // UUID of tag
	Name string `json:"name,omitempty"` // Name of tag
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                                  `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                              `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                                `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                                `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                                          `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                                      `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                                        `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                                        `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateProjectTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateProjectTemplatesDeviceTypes struct {
	ProductFamily string `json:"productFamily,omitempty"` // Device family
	ProductSeries string `json:"productSeries,omitempty"` // Device series
	ProductType   string `json:"productType,omitempty"`   // Device type
}
type RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                               `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                           `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                             `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                             `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateProjectTemplatesRollbackTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParams struct {
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
	Range           *[]RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange   `json:"range,omitempty"`           //
	Required        *bool                                                                       `json:"required,omitempty"`        // Is param required
	Selection       *RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelection `json:"selection,omitempty"`       //
}
type RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsRange struct {
	ID       string `json:"id,omitempty"`       // UUID of range
	MaxValue *int   `json:"maxValue,omitempty"` // Max value of range
	MinValue *int   `json:"minValue,omitempty"` // Min value of range
}
type RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelection struct {
	DefaultSelectedValues []string                                                                                   `json:"defaultSelectedValues,omitempty"` // Default selection values
	ID                    string                                                                                     `json:"id,omitempty"`                    // UUID of selection
	SelectionType         string                                                                                     `json:"selectionType,omitempty"`         // Type of selection(SINGLE_SELECT or MULTI_SELECT)
	SelectionValues       *RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelectionSelectionValues `json:"selectionValues,omitempty"`       // Selection values
}
type RequestConfigurationTemplatesUpdateProjectTemplatesTemplateParamsSelectionSelectionValues interface{}
type RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors struct {
	RollbackTemplateErrors *[]RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                       `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                       `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrorsTemplateErrors interface{}                                           // # Review unknown case
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
	RollbackTemplateErrors *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                                                `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                                                `json:"templateVersion,omitempty"`        // Current version of template
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
	RollbackTemplateErrors *[]RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                               `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                               `json:"templateVersion,omitempty"`        // Current version of template
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
	RollbackTemplateErrors *[]RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors `json:"rollbackTemplateErrors,omitempty"` // Validation or design conflicts errors of rollback template
	TemplateErrors         *[]RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors         `json:"templateErrors,omitempty"`         // Validation or design conflicts errors
	TemplateID             string                                                                               `json:"templateId,omitempty"`             // UUID of template
	TemplateVersion        string                                                                               `json:"templateVersion,omitempty"`        // Current version of template
}
type RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors interface{}
type RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors interface{}
type RequestConfigurationTemplatesDeployTemplate struct {
	ForcePushTemplate            *bool                                                                      `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  *bool                                                                      `json:"isComposite,omitempty"`                  // Composite template flag
	MainTemplateID               string                                                                     `json:"mainTemplateId,omitempty"`               // Main template UUID of versioned template
	MemberTemplateDeploymentInfo *[]RequestConfigurationTemplatesDeployTemplateMemberTemplateDeploymentInfo `json:"memberTemplateDeploymentInfo,omitempty"` // memberTemplateDeploymentInfo
	TargetInfo                   *[]RequestConfigurationTemplatesDeployTemplateTargetInfo                   `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                                                                     `json:"templateId,omitempty"`                   // UUID of template to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateMemberTemplateDeploymentInfo interface{}
type RequestConfigurationTemplatesDeployTemplateTargetInfo struct {
	HostName            string                                                                 `json:"hostName,omitempty"`            // Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
	ID                  string                                                                 `json:"id,omitempty"`                  // UUID of target is required if targetType is MANAGED_DEVICE_UUID
	Params              *RequestConfigurationTemplatesDeployTemplateTargetInfoParams           `json:"params,omitempty"`              // Template params/values to be provisioned
	ResourceParams      *[]RequestConfigurationTemplatesDeployTemplateTargetInfoResourceParams `json:"resourceParams,omitempty"`      // Resource params to be provisioned
	Type                string                                                                 `json:"type,omitempty"`                // Target type of device
	VersionedTemplateID string                                                                 `json:"versionedTemplateId,omitempty"` // Versioned templateUUID to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateTargetInfoParams map[string]interface{}
type RequestConfigurationTemplatesDeployTemplateTargetInfoResourceParams interface{}
type RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria []RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria // Array of RequestConfigurationTemplatesExportsTheTemplatesForAGivenCriteria
type RequestItemConfigurationTemplatesExportsTheTemplatesForAGivenCriteria interface{}
type RequestConfigurationTemplatesPreviewTemplate struct {
	DeviceID       string                                                        `json:"deviceId,omitempty"`       // UUID of device to get template preview
	Params         *RequestConfigurationTemplatesPreviewTemplateParams           `json:"params,omitempty"`         // Params to render preview
	ResourceParams *[]RequestConfigurationTemplatesPreviewTemplateResourceParams `json:"resourceParams,omitempty"` // Resource params to render preview
	TemplateID     string                                                        `json:"templateId,omitempty"`     // UUID of template to get template preview
}
type RequestConfigurationTemplatesPreviewTemplateParams interface{}
type RequestConfigurationTemplatesPreviewTemplateResourceParams interface{}
type RequestConfigurationTemplatesVersionTemplate struct {
	Comments   string `json:"comments,omitempty"`   // Template version comments
	TemplateID string `json:"templateId,omitempty"` // UUID of template
}
type RequestConfigurationTemplatesDeployTemplateV2 struct {
	ForcePushTemplate            *bool                                                                        `json:"forcePushTemplate,omitempty"`            //
	IsComposite                  *bool                                                                        `json:"isComposite,omitempty"`                  // Composite template flag
	MainTemplateID               string                                                                       `json:"mainTemplateId,omitempty"`               // Main template UUID of versioned template
	MemberTemplateDeploymentInfo *[]RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo `json:"memberTemplateDeploymentInfo,omitempty"` // memberTemplateDeploymentInfo
	TargetInfo                   *[]RequestConfigurationTemplatesDeployTemplateV2TargetInfo                   `json:"targetInfo,omitempty"`                   //
	TemplateID                   string                                                                       `json:"templateId,omitempty"`                   // UUID of template to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo interface{}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfo struct {
	HostName            string                                                                   `json:"hostName,omitempty"`            // Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
	ID                  string                                                                   `json:"id,omitempty"`                  // UUID of target is required if targetType is MANAGED_DEVICE_UUID
	Params              *RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams           `json:"params,omitempty"`              // Template params/values to be provisioned
	ResourceParams      *[]RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams `json:"resourceParams,omitempty"`      // Resource params to be provisioned
	Type                string                                                                   `json:"type,omitempty"`                // Target type of device
	VersionedTemplateID string                                                                   `json:"versionedTemplateId,omitempty"` // Versioned templateUUID to be provisioned
}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams map[string]interface{}
type RequestConfigurationTemplatesDeployTemplateV2TargetInfoResourceParams interface{}

//GetsAListOfProjects Gets a list of projects - 4f80-08c2-400b-98ee
/* List the projects


@param GetsAListOfProjectsQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetsAListOfProjects")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsAListOfProjects)
	return result, response, err

}

//GetsTheDetailsOfAGivenProject Gets the details of a given project. - dd91-a8c0-436a-82d9
/* Get the details of the given project by its id.


@param projectID projectId path parameter. projectId(UUID) of project to get project details

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
		return nil, response, fmt.Errorf("error with operation GetsTheDetailsOfAGivenProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject)
	return result, response, err

}

//GetsTheTemplatesAvailable Gets the templates available - e286-e848-47bb-a77e
/* List the templates available


@param GetsTheTemplatesAvailableQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetsTheTemplatesAvailable")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsTheTemplatesAvailable)
	return result, response, err

}

//StatusOfTemplateDeployment Status of template deployment - 078e-f800-49b8-80f1
/* API to retrieve the status of template deployment.


@param deploymentID deploymentId path parameter. UUID of deployment to retrieve template deployment status

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
		return nil, response, fmt.Errorf("error with operation StatusOfTemplateDeployment")
	}

	result := response.Result().(*ResponseConfigurationTemplatesStatusOfTemplateDeployment)
	return result, response, err

}

//GetsAllTheVersionsOfAGivenTemplate Gets all the versions of a given template - 0191-2b65-45fb-8891
/* Get all the versions of template by its id


@param templateID templateId path parameter. templateId(UUID) to get list of versioned templates

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
		return nil, response, fmt.Errorf("error with operation GetsAllTheVersionsOfAGivenTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsAllTheVersionsOfAGivenTemplate)
	return result, response, err

}

//GetsDetailsOfAGivenTemplate Gets details of a given template - 8c82-5900-40d9-8137
/* Details of the template by its id


@param templateID templateId path parameter. TemplateId(UUID) to get details of the template

@param GetsDetailsOfAGivenTemplateQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetsDetailsOfAGivenTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate)
	return result, response, err

}

//GetProjectsDetails Get project(s) details - 9a8c-aa6d-459b-a4a2
/* Get project(s) details


@param GetProjectsDetailsQueryParams Filtering parameter
*/
func (s *ConfigurationTemplatesService) GetProjectsDetails(GetProjectsDetailsQueryParams *GetProjectsDetailsQueryParams) (*ResponseConfigurationTemplatesGetProjectsDetails, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/project"

	queryString, _ := query.Values(GetProjectsDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetProjectsDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetProjectsDetails")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetProjectsDetails)
	return result, response, err

}

//GetTemplatesDetails Get template(s) details - b0b6-ba49-43c8-9f45
/* Get template(s) details


@param GetTemplatesDetailsQueryParams Filtering parameter
*/
func (s *ConfigurationTemplatesService) GetTemplatesDetails(GetTemplatesDetailsQueryParams *GetTemplatesDetailsQueryParams) (*ResponseConfigurationTemplatesGetTemplatesDetails, *resty.Response, error) {
	path := "/dna/intent/api/v2/template-programmer/template"

	queryString, _ := query.Values(GetTemplatesDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetTemplatesDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetTemplatesDetails")
	}

	result := response.Result().(*ResponseConfigurationTemplatesGetTemplatesDetails)
	return result, response, err

}

//CreatesACloneOfTheGivenTemplate Creates a clone of the given template - 0384-4a0a-4ee8-bfc2
/* API to clone template


@param name name path parameter. Template name to clone template(Name should be different than existing template name within same project)

@param templateID templateId path parameter. UUID of the template to clone it

@param projectID projectId path parameter.
@param CreatesACloneOfTheGivenTemplateQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation CreatesACloneOfTheGivenTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreatesACloneOfTheGivenTemplate)
	return result, response, err

}

//CreateProject Create Project - 5788-a8a8-4aa9-b97a
/* This API is used to create a new project.


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
		return nil, response, fmt.Errorf("error with operation CreateProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateProject)
	return result, response, err

}

//ImportsTheProjectsProvided Imports the Projects provided - f59f-a8ad-42d9-8b4f
/* Imports the Projects provided in the DTO


@param ImportsTheProjectsProvidedQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation ImportsTheProjectsProvided")
	}

	result := response.Result().(*ResponseConfigurationTemplatesImportsTheProjectsProvided)
	return result, response, err

}

//ExportsTheProjectsForAGivenCriteria Exports the projects for a given criteria. - 67bc-e964-45f8-b720
/* Exports the projects for given projectNames.


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
		return nil, response, fmt.Errorf("error with operation ExportsTheProjectsForAGivenCriteria")
	}

	result := response.Result().(*ResponseConfigurationTemplatesExportsTheProjectsForAGivenCriteria)
	return result, response, err

}

//ImportsTheTemplatesProvided Imports the templates provided - 4d86-f92a-4a7b-90bb
/* Imports the templates provided in the DTO by project Name


@param projectName projectName path parameter. Project name to create template under the project

@param ImportsTheTemplatesProvidedQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation ImportsTheTemplatesProvided")
	}

	result := response.Result().(*ResponseConfigurationTemplatesImportsTheTemplatesProvided)
	return result, response, err

}

//CreateTemplate Create Template - ab94-a88b-4b0b-8d3d
/* API to create a template by project id.


@param projectID projectId path parameter. UUID of the project in which the template needs to be created

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
		return nil, response, fmt.Errorf("error with operation CreateTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesCreateTemplate)
	return result, response, err

}

//DeployTemplate Deploy Template - 179f-09d8-430b-bee0
/* API to deploy a template.


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
		return nil, response, fmt.Errorf("error with operation DeployTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeployTemplate)
	return result, response, err

}

//ExportsTheTemplatesForAGivenCriteria Exports the templates for a given criteria. - a3a9-498f-4f48-a3c7
/* Exports the templates for given templateIds.


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
		return nil, response, fmt.Errorf("error with operation ExportsTheTemplatesForAGivenCriteria")
	}

	result := response.Result().(*ResponseConfigurationTemplatesExportsTheTemplatesForAGivenCriteria)
	return result, response, err

}

//VersionTemplate Version Template - f2a4-4a7e-4d5b-ab78
/* API to version the current contents of the template.


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
		return nil, response, fmt.Errorf("error with operation VersionTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesVersionTemplate)
	return result, response, err

}

//DeployTemplateV2 Deploy Template V2 - 02af-1bdf-4b48-9cbb
/* V2 API to deploy a template.


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
		return nil, response, fmt.Errorf("error with operation DeployTemplateV2")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeployTemplateV2)
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
		return nil, response, fmt.Errorf("error with operation PreviewTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesPreviewTemplate)
	return result, response, err

}

//DeletesTheProject Deletes the project - 8cbb-79f4-4259-82d4
/* Deletes the project by its id


@param projectID projectId path parameter. projectId(UUID) of project to be deleted

*/
func (s *ConfigurationTemplatesService) DeletesTheProject(projectID string) (*ResponseConfigurationTemplatesDeletesTheProject, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeletesTheProject")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeletesTheProject)
	return result, response, err

}

//DeletesTheTemplate Deletes the template - a2bb-4afd-4699-8965
/* Deletes the template by its id


@param templateID templateId path parameter. templateId(UUID) of template to be deleted

*/
func (s *ConfigurationTemplatesService) DeletesTheTemplate(templateID string) (*ResponseConfigurationTemplatesDeletesTheTemplate, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeletesTheTemplate")
	}

	result := response.Result().(*ResponseConfigurationTemplatesDeletesTheTemplate)
	return result, response, err

}
