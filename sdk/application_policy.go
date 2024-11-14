package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ApplicationPolicyService service

type GetApplicationPolicyQueryParams struct {
	PolicyScope string `url:"policyScope,omitempty"` //policy scope name
}
type GetApplicationPolicyQueuingProfileQueryParams struct {
	Name string `url:"name,omitempty"` //queuing profile name
}
type GetApplicationSetsQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //
	Limit  float64 `url:"limit,omitempty"`  //
	Name   string  `url:"name,omitempty"`   //
}
type DeleteApplicationSetQueryParams struct {
	ID string `url:"id,omitempty"` //
}
type DeleteApplicationQueryParams struct {
	ID string `url:"id,omitempty"` //Application's Id
}
type GetApplicationsQueryParams struct {
	Offset float64 `url:"offset,omitempty"` //The offset of the first application to be returned
	Limit  float64 `url:"limit,omitempty"`  //The maximum number of applications to be returned
	Name   string  `url:"name,omitempty"`   //Application's name
}
type GetQosDeviceInterfaceInfoQueryParams struct {
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //network device id
}
type GetApplicationSetsV2QueryParams struct {
	Attributes string  `url:"attributes,omitempty"` //Attributes to retrieve, valid value applicationSet
	Name       string  `url:"name,omitempty"`       //Application set name
	Offset     float64 `url:"offset,omitempty"`     //The starting point or index from where the paginated results should begin.
	Limit      float64 `url:"limit,omitempty"`      //The limit which is the maximum number of items to include in a single page of results, max value 500
}
type GetApplicationSetCountV2QueryParams struct {
	ScalableGroupType string `url:"scalableGroupType,omitempty"` //Scalable group type to retrieve, valid value APPLICATION_GROUP
}
type GetApplicationsV2QueryParams struct {
	Attributes string  `url:"attributes,omitempty"` //Attributes to retrieve, valid value application
	Name       string  `url:"name,omitempty"`       //The application name
	Offset     float64 `url:"offset,omitempty"`     //The starting point or index from where the paginated results should begin.
	Limit      float64 `url:"limit,omitempty"`      //The limit which is the maximum number of items to include in a single page of results, max value 500
}
type GetApplicationCountV2QueryParams struct {
	ScalableGroupType string `url:"scalableGroupType,omitempty"` //scalable group type to retrieve, valid value APPLICATION
}

type ResponseApplicationPolicyGetApplicationPolicy struct {
	Response *[]ResponseApplicationPolicyGetApplicationPolicyResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationPolicyResponse struct {
	ID                  string                                                                    `json:"id,omitempty"`                  // Id of Group based policy
	InstanceID          *int                                                                      `json:"instanceId,omitempty"`          // Instance id
	DisplayName         string                                                                    `json:"displayName,omitempty"`         // Display name
	InstanceCreatedOn   *int                                                                      `json:"instanceCreatedOn,omitempty"`   // Instance created on
	InstanceUpdatedOn   *int                                                                      `json:"instanceUpdatedOn,omitempty"`   // Instance updated on
	InstanceVersion     *float64                                                                  `json:"instanceVersion,omitempty"`     // Instance version
	CreateTime          *int                                                                      `json:"createTime,omitempty"`          // Create time
	Deployed            *bool                                                                     `json:"deployed,omitempty"`            // Deployed
	IsSeeded            *bool                                                                     `json:"isSeeded,omitempty"`            // Is seeded
	IsStale             *bool                                                                     `json:"isStale,omitempty"`             // Is stale
	LastUpdateTime      *int                                                                      `json:"lastUpdateTime,omitempty"`      // Last update time
	Name                string                                                                    `json:"name,omitempty"`                // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	Namespace           string                                                                    `json:"namespace,omitempty"`           // Namespace
	ProvisioningState   string                                                                    `json:"provisioningState,omitempty"`   // Provisioning state
	Qualifier           string                                                                    `json:"qualifier,omitempty"`           // Qualifier
	ResourceVersion     *float64                                                                  `json:"resourceVersion,omitempty"`     // Resource version
	TargetIDList        *[]ResponseApplicationPolicyGetApplicationPolicyResponseTargetIDList      `json:"targetIdList,omitempty"`        // Target id list
	Type                string                                                                    `json:"type,omitempty"`                // Type
	CfsChangeInfo       *[]ResponseApplicationPolicyGetApplicationPolicyResponseCfsChangeInfo     `json:"cfsChangeInfo,omitempty"`       // Cfs change info
	CustomProvisions    *[]ResponseApplicationPolicyGetApplicationPolicyResponseCustomProvisions  `json:"customProvisions,omitempty"`    // Custom provisions
	DeletePolicyStatus  string                                                                    `json:"deletePolicyStatus,omitempty"`  // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	Internal            *bool                                                                     `json:"internal,omitempty"`            // Internal
	IsDeleted           *bool                                                                     `json:"isDeleted,omitempty"`           // Is deleted
	IsEnabled           *bool                                                                     `json:"isEnabled,omitempty"`           // Is enabled
	IsScopeStale        *bool                                                                     `json:"isScopeStale,omitempty"`        // Is scope stale
	IseReserved         *bool                                                                     `json:"iseReserved,omitempty"`         // Is reserved
	PolicyScope         string                                                                    `json:"policyScope,omitempty"`         // Policy name
	PolicyStatus        string                                                                    `json:"policyStatus,omitempty"`        // Policy status
	Priority            *int                                                                      `json:"priority,omitempty"`            // Priority
	Pushed              *bool                                                                     `json:"pushed,omitempty"`              // Pushed
	AdvancedPolicyScope *ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScope `json:"advancedPolicyScope,omitempty"` //
	ContractList        *[]ResponseApplicationPolicyGetApplicationPolicyResponseContractList      `json:"contractList,omitempty"`        // Contract list
	ExclusiveContract   *ResponseApplicationPolicyGetApplicationPolicyResponseExclusiveContract   `json:"exclusiveContract,omitempty"`   //
	IDentitySource      *ResponseApplicationPolicyGetApplicationPolicyResponseIDentitySource      `json:"identitySource,omitempty"`      //
	Producer            *ResponseApplicationPolicyGetApplicationPolicyResponseProducer            `json:"producer,omitempty"`            //
	Consumer            *ResponseApplicationPolicyGetApplicationPolicyResponseConsumer            `json:"consumer,omitempty"`            //
}
type ResponseApplicationPolicyGetApplicationPolicyResponseTargetIDList interface{}
type ResponseApplicationPolicyGetApplicationPolicyResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetApplicationPolicyResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScope struct {
	ID                         string                                                                                                `json:"id,omitempty"`                         // Id of Advanced policy scope
	InstanceID                 *int                                                                                                  `json:"instanceId,omitempty"`                 // Instance id
	DisplayName                string                                                                                                `json:"displayName,omitempty"`                // Display name
	InstanceCreatedOn          *int                                                                                                  `json:"instanceCreatedOn,omitempty"`          // Instance created on
	InstanceUpdatedOn          *int                                                                                                  `json:"instanceUpdatedOn,omitempty"`          // Instance updated on
	InstanceVersion            *float64                                                                                              `json:"instanceVersion,omitempty"`            // Instance version
	Name                       string                                                                                                `json:"name,omitempty"`                       // Policy name
	AdvancedPolicyScopeElement *[]ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScopeAdvancedPolicyScopeElement `json:"advancedPolicyScopeElement,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScopeAdvancedPolicyScopeElement struct {
	ID                string                                                                                                    `json:"id,omitempty"`                // Id of Advanced policy scope element
	InstanceID        *int                                                                                                      `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                                                    `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                                                      `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                                                      `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                                                  `json:"instanceVersion,omitempty"`   // Instance version
	GroupID           []string                                                                                                  `json:"groupId,omitempty"`           // Group id
	SSID              *[]ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScopeAdvancedPolicyScopeElementSSID `json:"ssid,omitempty"`              // Ssid
}
type ResponseApplicationPolicyGetApplicationPolicyResponseAdvancedPolicyScopeAdvancedPolicyScopeElementSSID interface{}
type ResponseApplicationPolicyGetApplicationPolicyResponseContractList interface{}
type ResponseApplicationPolicyGetApplicationPolicyResponseExclusiveContract struct {
	ID                string                                                                          `json:"id,omitempty"`                // Id of Exclusive contract
	InstanceID        *int                                                                            `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                          `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                            `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                            `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                        `json:"instanceVersion,omitempty"`   // Instance version
	Clause            *[]ResponseApplicationPolicyGetApplicationPolicyResponseExclusiveContractClause `json:"clause,omitempty"`            //
}
type ResponseApplicationPolicyGetApplicationPolicyResponseExclusiveContractClause struct {
	ID                    string   `json:"id,omitempty"`                    // Id of Business relevance or Application policy knobs clause
	InstanceID            *int     `json:"instanceId,omitempty"`            // Instance id
	DisplayName           string   `json:"displayName,omitempty"`           // Display name
	InstanceCreatedOn     *int     `json:"instanceCreatedOn,omitempty"`     // Instance created on
	InstanceUpdatedOn     *int     `json:"instanceUpdatedOn,omitempty"`     // Instance updated on
	InstanceVersion       *float64 `json:"instanceVersion,omitempty"`       // Instance version
	Priority              *int     `json:"priority,omitempty"`              // Priority
	Type                  string   `json:"type,omitempty"`                  // Type
	RelevanceLevel        string   `json:"relevanceLevel,omitempty"`        // Relevance level
	DeviceRemovalBehavior string   `json:"deviceRemovalBehavior,omitempty"` // Device removal behavior
	HostTrackingEnabled   *bool    `json:"hostTrackingEnabled,omitempty"`   // Host tracking enabled
}
type ResponseApplicationPolicyGetApplicationPolicyResponseIDentitySource struct {
	ID                string   `json:"id,omitempty"`                // Id of Identity source
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	State             string   `json:"state,omitempty"`             // State
	Type              string   `json:"type,omitempty"`              // Type
}
type ResponseApplicationPolicyGetApplicationPolicyResponseProducer struct {
	ID                string                                                                        `json:"id,omitempty"`                // Id of Producer
	InstanceID        *int                                                                          `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                        `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                          `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                          `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                      `json:"instanceVersion,omitempty"`   // Instance version
	ScalableGroup     *[]ResponseApplicationPolicyGetApplicationPolicyResponseProducerScalableGroup `json:"scalableGroup,omitempty"`     //
}
type ResponseApplicationPolicyGetApplicationPolicyResponseProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type ResponseApplicationPolicyGetApplicationPolicyResponseConsumer struct {
	ID                string                                                                        `json:"id,omitempty"`                // Id of Consumer
	InstanceID        *int                                                                          `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                        `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                          `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                          `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                      `json:"instanceVersion,omitempty"`   // Instance version
	ScalableGroup     *[]ResponseApplicationPolicyGetApplicationPolicyResponseConsumerScalableGroup `json:"scalableGroup,omitempty"`     //
}
type ResponseApplicationPolicyGetApplicationPolicyResponseConsumerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application Scalable group
}
type ResponseApplicationPolicyGetApplicationPolicyDefault struct {
	Response *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponse struct {
	ID                 string                                                                          `json:"id,omitempty"`                 // Id of Group based policy
	InstanceID         *int                                                                            `json:"instanceId,omitempty"`         // Instance id
	DisplayName        string                                                                          `json:"displayName,omitempty"`        // Display name
	InstanceCreatedOn  *int                                                                            `json:"instanceCreatedOn,omitempty"`  // Instance created on
	InstanceUpdatedOn  *int                                                                            `json:"instanceUpdatedOn,omitempty"`  // Instance updated on
	InstanceVersion    *float64                                                                        `json:"instanceVersion,omitempty"`    // Instance version
	CreateTime         *int                                                                            `json:"createTime,omitempty"`         // Create time
	Deployed           *bool                                                                           `json:"deployed,omitempty"`           // Deployed
	IsSeeded           *bool                                                                           `json:"isSeeded,omitempty"`           // Is seeded
	IsStale            *bool                                                                           `json:"isStale,omitempty"`            // Is stale
	LastUpdateTime     *int                                                                            `json:"lastUpdateTime,omitempty"`     // Last update time
	Name               string                                                                          `json:"name,omitempty"`               // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	Namespace          string                                                                          `json:"namespace,omitempty"`          // Namespace
	ProvisioningState  string                                                                          `json:"provisioningState,omitempty"`  // Provisioning state
	Qualifier          string                                                                          `json:"qualifier,omitempty"`          // Qualifier
	ResourceVersion    *float64                                                                        `json:"resourceVersion,omitempty"`    // Resource version
	TargetIDList       *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponseTargetIDList     `json:"targetIdList,omitempty"`       // Target id list
	Type               string                                                                          `json:"type,omitempty"`               // Type
	CfsChangeInfo      *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponseCfsChangeInfo    `json:"cfsChangeInfo,omitempty"`      // Cfs change info
	CustomProvisions   *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponseCustomProvisions `json:"customProvisions,omitempty"`   // Custom provisions
	DeletePolicyStatus string                                                                          `json:"deletePolicyStatus,omitempty"` // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	Internal           *bool                                                                           `json:"internal,omitempty"`           // Internal
	IsDeleted          *bool                                                                           `json:"isDeleted,omitempty"`          // Is deleted
	IsEnabled          *bool                                                                           `json:"isEnabled,omitempty"`          // Is enabled
	IsScopeStale       *bool                                                                           `json:"isScopeStale,omitempty"`       // Is scope stale
	IseReserved        *bool                                                                           `json:"iseReserved,omitempty"`        // Is reserved
	PolicyStatus       string                                                                          `json:"policyStatus,omitempty"`       // Policy status
	Priority           *int                                                                            `json:"priority,omitempty"`           // Priority
	Pushed             *bool                                                                           `json:"pushed,omitempty"`             // Pushed
	ContractList       *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponseContractList     `json:"contractList,omitempty"`       // Contract list
	ExclusiveContract  *ResponseApplicationPolicyGetApplicationPolicyDefaultResponseExclusiveContract  `json:"exclusiveContract,omitempty"`  //
	IDentitySource     *ResponseApplicationPolicyGetApplicationPolicyDefaultResponseIDentitySource     `json:"identitySource,omitempty"`     //
	Producer           *ResponseApplicationPolicyGetApplicationPolicyDefaultResponseProducer           `json:"producer,omitempty"`           //
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseTargetIDList interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseContractList interface{}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseExclusiveContract struct {
	ID                string                                                                                 `json:"id,omitempty"`                // Id of Exclusive contract
	InstanceID        *int                                                                                   `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                                 `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                                   `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                                   `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                               `json:"instanceVersion,omitempty"`   // Instance version
	Clause            *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponseExclusiveContractClause `json:"clause,omitempty"`            //
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseExclusiveContractClause struct {
	ID                string   `json:"id,omitempty"`                // Id of Business relevance clause
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	Priority          *int     `json:"priority,omitempty"`          // Priority
	Type              string   `json:"type,omitempty"`              // Type
	RelevanceLevel    string   `json:"relevanceLevel,omitempty"`    // Relevance level
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseIDentitySource struct {
	ID                string   `json:"id,omitempty"`                // Id of Identity source
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	State             string   `json:"state,omitempty"`             // State
	Type              string   `json:"type,omitempty"`              // Type
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseProducer struct {
	ID                string                                                                               `json:"id,omitempty"`                // Id of Producer
	InstanceID        *int                                                                                 `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string                                                                               `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int                                                                                 `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int                                                                                 `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64                                                                             `json:"instanceVersion,omitempty"`   // Instance version
	ScalableGroup     *[]ResponseApplicationPolicyGetApplicationPolicyDefaultResponseProducerScalableGroup `json:"scalableGroup,omitempty"`     //
}
type ResponseApplicationPolicyGetApplicationPolicyDefaultResponseProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type ResponseApplicationPolicyApplicationPolicyIntent struct {
	Response *ResponseApplicationPolicyApplicationPolicyIntentResponse `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyApplicationPolicyIntentResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfile struct {
	Response *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse `json:"response,omitempty"` //
	Version  string                                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponse struct {
	ID                 string                                                                                   `json:"id,omitempty"`                 // Id of Queueing profile
	InstanceID         *int                                                                                     `json:"instanceId,omitempty"`         // Instance id
	DisplayName        string                                                                                   `json:"displayName,omitempty"`        // Display name
	InstanceCreatedOn  *int                                                                                     `json:"instanceCreatedOn,omitempty"`  // Instance created on
	InstanceUpdatedOn  *int                                                                                     `json:"instanceUpdatedOn,omitempty"`  // Instance updated on
	InstanceVersion    *float64                                                                                 `json:"instanceVersion,omitempty"`    // Instance version
	CreateTime         *int                                                                                     `json:"createTime,omitempty"`         // Create time
	Deployed           *bool                                                                                    `json:"deployed,omitempty"`           // Deployed
	Description        string                                                                                   `json:"description,omitempty"`        // Free test description
	IsSeeded           *bool                                                                                    `json:"isSeeded,omitempty"`           // Is seeded
	IsStale            *bool                                                                                    `json:"isStale,omitempty"`            // Is stale
	LastUpdateTime     *int                                                                                     `json:"lastUpdateTime,omitempty"`     // Last update time
	Name               string                                                                                   `json:"name,omitempty"`               // Queueing profile name
	Namespace          string                                                                                   `json:"namespace,omitempty"`          // Namespace
	ProvisioningState  string                                                                                   `json:"provisioningState,omitempty"`  // Provisioning State
	Qualifier          string                                                                                   `json:"qualifier,omitempty"`          // Qualifier
	ResourceVersion    *float64                                                                                 `json:"resourceVersion,omitempty"`    // Resource version
	TargetIDList       *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseTargetIDList       `json:"targetIdList,omitempty"`       // Target id list
	Type               string                                                                                   `json:"type,omitempty"`               // Type
	CfsChangeInfo      *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseCfsChangeInfo      `json:"cfsChangeInfo,omitempty"`      // Cfs change info
	CustomProvisions   *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseCustomProvisions   `json:"customProvisions,omitempty"`   // Custom provisions
	GenID              *float64                                                                                 `json:"genId,omitempty"`              // Gen id
	Internal           *bool                                                                                    `json:"internal,omitempty"`           // Internal
	IsDeleted          *bool                                                                                    `json:"isDeleted,omitempty"`          // Is deleted
	IseReserved        *bool                                                                                    `json:"iseReserved,omitempty"`        // Is reserved
	Pushed             *bool                                                                                    `json:"pushed,omitempty"`             // Pushed
	Clause             *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClause             `json:"clause,omitempty"`             //
	ContractClassifier *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseContractClassifier `json:"contractClassifier,omitempty"` // Contract classifier
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseTargetIDList interface{}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClause struct {
	ID                                string                                                                                                     `json:"id,omitempty"`                                // Id
	InstanceID                        *int                                                                                                       `json:"instanceId,omitempty"`                        // Instance id
	DisplayName                       string                                                                                                     `json:"displayName,omitempty"`                       // Display name
	InstanceCreatedOn                 *int                                                                                                       `json:"instanceCreatedOn,omitempty"`                 // Instance created on
	InstanceUpdatedOn                 *int                                                                                                       `json:"instanceUpdatedOn,omitempty"`                 // Instance updated on
	InstanceVersion                   *float64                                                                                                   `json:"instanceVersion,omitempty"`                   // Instance version
	Priority                          *int                                                                                                       `json:"priority,omitempty"`                          // Priority
	Type                              string                                                                                                     `json:"type,omitempty"`                              // Type
	IsCommonBetweenAllInterfaceSpeeds *bool                                                                                                      `json:"isCommonBetweenAllInterfaceSpeeds,omitempty"` // Is common between all interface speeds
	InterfaceSpeedBandwidthClauses    *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseInterfaceSpeedBandwidthClauses `json:"interfaceSpeedBandwidthClauses,omitempty"`    //
	TcDscpSettings                    *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseTcDscpSettings                 `json:"tcDscpSettings,omitempty"`                    //
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseInterfaceSpeedBandwidthClauses struct {
	ID                  string                                                                                                                        `json:"id,omitempty"`                  // Id
	InstanceID          *int                                                                                                                          `json:"instanceId,omitempty"`          // Instance id
	DisplayName         string                                                                                                                        `json:"displayName,omitempty"`         // Display name
	InstanceCreatedOn   *int                                                                                                                          `json:"instanceCreatedOn,omitempty"`   // Instance created on
	InstanceUpdatedOn   *int                                                                                                                          `json:"instanceUpdatedOn,omitempty"`   // Instance updated on
	InstanceVersion     *float64                                                                                                                      `json:"instanceVersion,omitempty"`     // Instance version
	InterfaceSpeed      string                                                                                                                        `json:"interfaceSpeed,omitempty"`      // Interface speed
	TcBandwidthSettings *[]ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `json:"tcBandwidthSettings,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	ID                  string   `json:"id,omitempty"`                  // Id
	InstanceID          *int     `json:"instanceId,omitempty"`          // Instance id
	DisplayName         string   `json:"displayName,omitempty"`         // Display name
	InstanceCreatedOn   *int     `json:"instanceCreatedOn,omitempty"`   // Instance created on
	InstanceUpdatedOn   *int     `json:"instanceUpdatedOn,omitempty"`   // Instance updated on
	InstanceVersion     *float64 `json:"instanceVersion,omitempty"`     // Instance version
	BandwidthPercentage *int     `json:"bandwidthPercentage,omitempty"` // Bandwidth percentage
	TrafficClass        string   `json:"trafficClass,omitempty"`        // Traffic Class
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseClauseTcDscpSettings struct {
	ID                string   `json:"id,omitempty"`                // Id
	InstanceID        *int     `json:"instanceId,omitempty"`        // Instance id
	DisplayName       string   `json:"displayName,omitempty"`       // Display name
	InstanceCreatedOn *int     `json:"instanceCreatedOn,omitempty"` // Instance created on
	InstanceUpdatedOn *int     `json:"instanceUpdatedOn,omitempty"` // Instance updated on
	InstanceVersion   *float64 `json:"instanceVersion,omitempty"`   // Instance version
	Dscp              string   `json:"dscp,omitempty"`              // Dscp value
	TrafficClass      string   `json:"trafficClass,omitempty"`      // Traffic Class
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileResponseContractClassifier interface{}
type ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfile struct {
	Response *ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateApplicationPolicyQueuingProfile struct {
	Response *ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationPolicyQueuingProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCount struct {
	Response *int   `json:"response,omitempty"` // Total number of Queueing Profile
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfile struct {
	Response *ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileResponse `json:"response,omitempty"` //
	Version  string                                                                  `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfileResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationSets struct {
	Response *[]ResponseApplicationPolicyGetApplicationSetsResponse `json:"response,omitempty"` //
}
type ResponseApplicationPolicyGetApplicationSetsResponse struct {
	ID             string                                                             `json:"id,omitempty"`             // Id
	IDentitySource *ResponseApplicationPolicyGetApplicationSetsResponseIDentitySource `json:"identitySource,omitempty"` //
	Name           string                                                             `json:"name,omitempty"`           // Name
}
type ResponseApplicationPolicyGetApplicationSetsResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type
}
type ResponseApplicationPolicyDeleteApplicationSet struct {
	Response *ResponseApplicationPolicyDeleteApplicationSetResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyCreateApplicationSet struct {
	Response *ResponseApplicationPolicyCreateApplicationSetResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationSetResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyGetApplicationSetsCount struct {
	Response string `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplication struct {
	Response *ResponseApplicationPolicyCreateApplicationResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyEditApplication struct {
	Response *ResponseApplicationPolicyEditApplicationResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyEditApplicationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyDeleteApplication struct {
	Response *ResponseApplicationPolicyDeleteApplicationResponse `json:"response,omitempty"` //
	Version  string                                              `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseApplicationPolicyGetApplications []ResponseItemApplicationPolicyGetApplications // Array of ResponseApplicationPolicyGetApplications
type ResponseItemApplicationPolicyGetApplications struct {
	ID                  string                                                             `json:"id,omitempty"`                  // Id
	Name                string                                                             `json:"name,omitempty"`                // Name
	NetworkApplications *[]ResponseItemApplicationPolicyGetApplicationsNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]ResponseItemApplicationPolicyGetApplicationsNetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *ResponseItemApplicationPolicyGetApplicationsApplicationSet        `json:"applicationSet,omitempty"`      //
}
type ResponseItemApplicationPolicyGetApplicationsNetworkApplications struct {
	ID                 string `json:"id,omitempty"`                 // Id
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type ResponseItemApplicationPolicyGetApplicationsNetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // Id
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type ResponseItemApplicationPolicyGetApplicationsApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type ResponseApplicationPolicyGetApplicationsCount struct {
	Response string `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfo struct {
	Response *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse struct {
	ID                     string                                                                              `json:"id,omitempty"`                     // Id of Qos device info
	InstanceID             *int                                                                                `json:"instanceId,omitempty"`             // Instance id
	DisplayName            string                                                                              `json:"displayName,omitempty"`            // Display name
	InstanceCreatedOn      *int                                                                                `json:"instanceCreatedOn,omitempty"`      // Instance created on
	InstanceUpdatedOn      *int                                                                                `json:"instanceUpdatedOn,omitempty"`      // Instance updated on
	InstanceVersion        *int                                                                                `json:"instanceVersion,omitempty"`        // Instance version
	CreateTime             *int                                                                                `json:"createTime,omitempty"`             // Create time
	Deployed               *bool                                                                               `json:"deployed,omitempty"`               // Deployed
	IsSeeded               *bool                                                                               `json:"isSeeded,omitempty"`               // Is seeded
	IsStale                *bool                                                                               `json:"isStale,omitempty"`                // Is stale
	LastUpdateTime         *int                                                                                `json:"lastUpdateTime,omitempty"`         // Last update time
	Name                   string                                                                              `json:"name,omitempty"`                   // Device name
	Namespace              string                                                                              `json:"namespace,omitempty"`              // Namespace
	ProvisioningState      string                                                                              `json:"provisioningState,omitempty"`      // Provisioning state
	Qualifier              string                                                                              `json:"qualifier,omitempty"`              // Qualifier
	ResourceVersion        *int                                                                                `json:"resourceVersion,omitempty"`        // Resource version
	TargetIDList           *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseTargetIDList           `json:"targetIdList,omitempty"`           // Target id list
	Type                   string                                                                              `json:"type,omitempty"`                   // Type
	CfsChangeInfo          *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseCfsChangeInfo          `json:"cfsChangeInfo,omitempty"`          // Cfs change info
	CustomProvisions       *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseCustomProvisions       `json:"customProvisions,omitempty"`       // Custom provisions
	ExcludedInterfaces     []string                                                                            `json:"excludedInterfaces,omitempty"`     // excluded interfaces ids
	IsExcluded             *bool                                                                               `json:"isExcluded,omitempty"`             // Is excluded
	NetworkDeviceID        string                                                                              `json:"networkDeviceId,omitempty"`        // Network device id
	QosDeviceInterfaceInfo *[]ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseQosDeviceInterfaceInfo `json:"qosDeviceInterfaceInfo,omitempty"` //
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseTargetIDList interface{}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseCfsChangeInfo interface{}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseCustomProvisions interface{}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponseQosDeviceInterfaceInfo struct {
	ID                 string   `json:"id,omitempty"`                 // Id of Qos device interface info
	InstanceID         *int     `json:"instanceId,omitempty"`         // Instance id
	DisplayName        string   `json:"displayName,omitempty"`        // Display name
	InstanceCreatedOn  *int     `json:"instanceCreatedOn,omitempty"`  // Instance created on
	InstanceUpdatedOn  *int     `json:"instanceUpdatedOn,omitempty"`  // Instance updated on
	InstanceVersion    *float64 `json:"instanceVersion,omitempty"`    // Instance version
	DmvpnRemoteSitesBw *[]int   `json:"dmvpnRemoteSitesBw,omitempty"` // Dmvpn remote sites bandwidth
	DownloadBW         *float64 `json:"downloadBW,omitempty"`         // Download bandwidth
	InterfaceID        string   `json:"interfaceId,omitempty"`        // Interface id
	InterfaceName      string   `json:"interfaceName,omitempty"`      // Interface name
	Label              string   `json:"label,omitempty"`              // SP Profile name
	Role               string   `json:"role,omitempty"`               // Interface role
	UploadBW           *int     `json:"uploadBW,omitempty"`           // Upload bandwidth
}
type ResponseApplicationPolicyUpdateQosDeviceInterfaceInfo struct {
	Response *ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyUpdateQosDeviceInterfaceInfoResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateQosDeviceInterfaceInfo struct {
	Response *ResponseApplicationPolicyCreateQosDeviceInterfaceInfoResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateQosDeviceInterfaceInfoResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetQosDeviceInterfaceInfoCount struct {
	Response *int   `json:"response,omitempty"` // Total number of Qos Device Interface Info
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteQosDeviceInterfaceInfo struct {
	Response *ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoResponse `json:"response,omitempty"` //
	Version  string                                                         `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteQosDeviceInterfaceInfoResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateApplicationSetsV2 struct {
	Response *ResponseApplicationPolicyCreateApplicationSetsV2Response `json:"response,omitempty"` //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationSetsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationSetsV2 struct {
	Response *[]ResponseApplicationPolicyGetApplicationSetsV2Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationSetsV2Response struct {
	ID                          string                                                               `json:"id,omitempty"`                          // Id of Application Set
	InstanceID                  *int                                                                 `json:"instanceId,omitempty"`                  // Instance id
	DisplayName                 string                                                               `json:"displayName,omitempty"`                 // Display name
	InstanceVersion             *float64                                                             `json:"instanceVersion,omitempty"`             // Instance version
	DefaultBusinessRelevance    string                                                               `json:"defaultBusinessRelevance,omitempty"`    // Default business relevance
	IDentitySource              *ResponseApplicationPolicyGetApplicationSetsV2ResponseIDentitySource `json:"identitySource,omitempty"`              //
	Name                        string                                                               `json:"name,omitempty"`                        // Application Set name
	Namespace                   string                                                               `json:"namespace,omitempty"`                   // Namespace, valid value scalablegroup:application
	ScalableGroupExternalHandle string                                                               `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be equal to Application Set name
	ScalableGroupType           string                                                               `json:"scalableGroupType,omitempty"`           // Scalable group type, valid value APPLICATION_GROUP
	Type                        string                                                               `json:"type,omitempty"`                        // Type, valid value scalablegroup
}
type ResponseApplicationPolicyGetApplicationSetsV2ResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type of identify source. NBAR: build in Application Set, APIC_EM: custom Application Set
}
type ResponseApplicationPolicyGetApplicationSetCountV2 struct {
	Response *int   `json:"response,omitempty"` // Total number of Application Set
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetV2 struct {
	Response *ResponseApplicationPolicyDeleteApplicationSetV2Response `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationSetV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyEditApplicationsV2 struct {
	Response *ResponseApplicationPolicyEditApplicationsV2Response `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyEditApplicationsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyCreateApplicationsV2 struct {
	Response *ResponseApplicationPolicyCreateApplicationsV2Response `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyCreateApplicationsV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type ResponseApplicationPolicyGetApplicationsV2 struct {
	Response *[]ResponseApplicationPolicyGetApplicationsV2Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyGetApplicationsV2Response struct {
	ID                          string                                                                         `json:"id,omitempty"`                          // Id of Application
	InstanceID                  *int                                                                           `json:"instanceId,omitempty"`                  // Instance id
	DisplayName                 string                                                                         `json:"displayName,omitempty"`                 // Display name
	InstanceVersion             *float64                                                                       `json:"instanceVersion,omitempty"`             // Instance version
	IDentitySource              *ResponseApplicationPolicyGetApplicationsV2ResponseIDentitySource              `json:"identitySource,omitempty"`              //
	IndicativeNetworkIDentity   *[]ResponseApplicationPolicyGetApplicationsV2ResponseIndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"`   //
	Name                        string                                                                         `json:"name,omitempty"`                        // Application name
	Namespace                   string                                                                         `json:"namespace,omitempty"`                   // Namespace, valid value scalablegroup:application
	NetworkApplications         *[]ResponseApplicationPolicyGetApplicationsV2ResponseNetworkApplications       `json:"networkApplications,omitempty"`         //
	NetworkIDentity             *[]ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentity           `json:"networkIdentity,omitempty"`             //
	ParentScalableGroup         *ResponseApplicationPolicyGetApplicationsV2ResponseParentScalableGroup         `json:"parentScalableGroup,omitempty"`         //
	Qualifier                   string                                                                         `json:"qualifier,omitempty"`                   // Qualifier, valid value application
	ScalableGroupExternalHandle string                                                                         `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be equal to Application name
	ScalableGroupType           string                                                                         `json:"scalableGroupType,omitempty"`           // Scalable group type, valid value APPLICATION
	Type                        string                                                                         `json:"type,omitempty"`                        // Type, valid value scalablegroup
}
type ResponseApplicationPolicyGetApplicationsV2ResponseIDentitySource struct {
	ID   string `json:"id,omitempty"`   // Id
	Type string `json:"type,omitempty"` // Type of identify source. NBAR: build in Application, APIC_EM: custom Application
}
type ResponseApplicationPolicyGetApplicationsV2ResponseIndicativeNetworkIDentity struct {
	ID          string   `json:"id,omitempty"`          // Id
	DisplayName string   `json:"displayName,omitempty"` // Display name
	LowerPort   *float64 `json:"lowerPort,omitempty"`   // Lower port
	Ports       string   `json:"ports,omitempty"`       // Ports
	Protocol    string   `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64 `json:"upperPort,omitempty"`   // Upper port
}
type ResponseApplicationPolicyGetApplicationsV2ResponseNetworkApplications struct {
	ID                 string   `json:"id,omitempty"`                 // Id
	AppProtocol        string   `json:"appProtocol,omitempty"`        // App protocol
	ApplicationSubType string   `json:"applicationSubType,omitempty"` // Application sub type, LEARNED: discovered application, NONE: nbar and custom application
	ApplicationType    string   `json:"applicationType,omitempty"`    // Application type, DEFAULT: nbar application, DEFAULT_MODIFIED: nbar modified application, CUSTOM: custom application
	CategoryID         string   `json:"categoryId,omitempty"`         // Category id
	DisplayName        string   `json:"displayName,omitempty"`        // Display name
	Dscp               string   `json:"dscp,omitempty"`               // Dscp
	EngineID           string   `json:"engineId,omitempty"`           // Engine id
	HelpString         string   `json:"helpString,omitempty"`         // Help string
	LongDescription    string   `json:"longDescription,omitempty"`    // Long description
	Name               string   `json:"name,omitempty"`               // Application name
	Popularity         *float64 `json:"popularity,omitempty"`         // Popularity
	Rank               *int     `json:"rank,omitempty"`               // Rank, any value between 1 to 65535
	SelectorID         string   `json:"selectorId,omitempty"`         // Selector id
	ServerName         string   `json:"serverName,omitempty"`         // Server name
	URL                string   `json:"url,omitempty"`                // Url
	TrafficClass       string   `json:"trafficClass,omitempty"`       // Traffic class
}
type ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentity struct {
	ID          string                                                                         `json:"id,omitempty"`          // Id
	DisplayName string                                                                         `json:"displayName,omitempty"` // Display name
	IPv4Subnet  []string                                                                       `json:"ipv4Subnet,omitempty"`  // Ipv4 subnet
	IPv6Subnet  *[]ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentityIPv6Subnet `json:"ipv6Subnet,omitempty"`  // Ipv6 subnet
	LowerPort   *float64                                                                       `json:"lowerPort,omitempty"`   // Lower port
	Ports       string                                                                         `json:"ports,omitempty"`       // Ports
	Protocol    string                                                                         `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64                                                                       `json:"upperPort,omitempty"`   // Upper port
}
type ResponseApplicationPolicyGetApplicationsV2ResponseNetworkIDentityIPv6Subnet interface{}
type ResponseApplicationPolicyGetApplicationsV2ResponseParentScalableGroup struct {
	ID    string `json:"id,omitempty"`    // Id
	IDRef string `json:"idRef,omitempty"` // Id reference to parent application set
}
type ResponseApplicationPolicyGetApplicationCountV2 struct {
	Response *int   `json:"response,omitempty"` // Total number of Application
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationV2 struct {
	Response *ResponseApplicationPolicyDeleteApplicationV2Response `json:"response,omitempty"` //
	Version  string                                                `json:"version,omitempty"`  // Version
}
type ResponseApplicationPolicyDeleteApplicationV2Response struct {
	TaskID string `json:"taskId,omitempty"` // Task id
	URL    string `json:"url,omitempty"`    // Task url
}
type RequestApplicationPolicyApplicationPolicyIntent struct {
	CreateList *[]RequestApplicationPolicyApplicationPolicyIntentCreateList `json:"createList,omitempty"` //
	UpdateList *[]RequestApplicationPolicyApplicationPolicyIntentUpdateList `json:"updateList,omitempty"` //
	DeleteList []string                                                     `json:"deleteList,omitempty"` // Delete list of Group Based Policy ids
}
type RequestApplicationPolicyApplicationPolicyIntentCreateList struct {
	Name                string                                                                        `json:"name,omitempty"`                // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	DeletePolicyStatus  string                                                                        `json:"deletePolicyStatus,omitempty"`  // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	PolicyScope         string                                                                        `json:"policyScope,omitempty"`         // Policy name
	Priority            string                                                                        `json:"priority,omitempty"`            // Set to 4095 while producer refer to application Scalable group otherwise 100
	AdvancedPolicyScope *RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScope `json:"advancedPolicyScope,omitempty"` //
	ExclusiveContract   *RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContract   `json:"exclusiveContract,omitempty"`   //
	Contract            *RequestApplicationPolicyApplicationPolicyIntentCreateListContract            `json:"contract,omitempty"`            //
	Producer            *RequestApplicationPolicyApplicationPolicyIntentCreateListProducer            `json:"producer,omitempty"`            //
	Consumer            *RequestApplicationPolicyApplicationPolicyIntentCreateListConsumer            `json:"consumer,omitempty"`            //
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScope struct {
	Name                       string                                                                                                    `json:"name,omitempty"`                       // Policy name
	AdvancedPolicyScopeElement *[]RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement `json:"advancedPolicyScopeElement,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement struct {
	GroupID []string `json:"groupId,omitempty"` // The site(s) ID where the Application QoS Policy will be deployed.
	SSID    []string `json:"ssid,omitempty"`    // Ssid
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContract struct {
	Clause *[]RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause `json:"clause,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause struct {
	Type                  string `json:"type,omitempty"`                  // Type
	RelevanceLevel        string `json:"relevanceLevel,omitempty"`        // Relevance level
	DeviceRemovalBehavior string `json:"deviceRemovalBehavior,omitempty"` // Device eemoval behavior
	HostTrackingEnabled   *bool  `json:"hostTrackingEnabled,omitempty"`   // Is host tracking enabled
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListContract struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to Queueing profile
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListProducer struct {
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListConsumer struct {
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application Scalable group
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateList struct {
	ID                  string                                                                        `json:"id,omitempty"`                  // Id of Group based policy
	Name                string                                                                        `json:"name,omitempty"`                // Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
	DeletePolicyStatus  string                                                                        `json:"deletePolicyStatus,omitempty"`  // NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
	PolicyScope         string                                                                        `json:"policyScope,omitempty"`         // Policy name
	Priority            string                                                                        `json:"priority,omitempty"`            // Set to 4095 while producer refer to application Scalable group otherwise 100
	AdvancedPolicyScope *RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScope `json:"advancedPolicyScope,omitempty"` //
	ExclusiveContract   *RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContract   `json:"exclusiveContract,omitempty"`   //
	Contract            *RequestApplicationPolicyApplicationPolicyIntentUpdateListContract            `json:"contract,omitempty"`            //
	Producer            *RequestApplicationPolicyApplicationPolicyIntentUpdateListProducer            `json:"producer,omitempty"`            //
	Consumer            *RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumer            `json:"consumer,omitempty"`            //
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScope struct {
	ID                         string                                                                                                    `json:"id,omitempty"`                         // Id of Advance policy scope
	Name                       string                                                                                                    `json:"name,omitempty"`                       // Policy name
	AdvancedPolicyScopeElement *[]RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement `json:"advancedPolicyScopeElement,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement struct {
	ID      string   `json:"id,omitempty"`      // Id of Advance policy scope element
	GroupID []string `json:"groupId,omitempty"` // The site(s) ID where the Application QoS Policy will be deployed.
	SSID    []string `json:"ssid,omitempty"`    // Ssid
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContract struct {
	ID     string                                                                              `json:"id,omitempty"`     // Id of Exclusive contract
	Clause *[]RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause `json:"clause,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause struct {
	ID                    string `json:"id,omitempty"`                    // Id of Business relevance or Application policy knobs clause
	Type                  string `json:"type,omitempty"`                  // Type
	RelevanceLevel        string `json:"relevanceLevel,omitempty"`        // Relevance level
	DeviceRemovalBehavior string `json:"deviceRemovalBehavior,omitempty"` // Device removal behavior
	HostTrackingEnabled   *bool  `json:"hostTrackingEnabled,omitempty"`   // Host tracking enabled
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListContract struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to Queueing profile
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListProducer struct {
	ID            string                                                                            `json:"id,omitempty"`            // Id of Producer
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application-set or application Scalable group
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumer struct {
	ID            string                                                                            `json:"id,omitempty"`            // Id of Consumer
	ScalableGroup *[]RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup `json:"scalableGroup,omitempty"` //
}
type RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id ref to application Scalable group
}
type RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile []RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile // Array of RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfile struct {
	ID          string                                                                     `json:"id,omitempty"`          // Id of Queueing profile
	Description string                                                                     `json:"description,omitempty"` // Free test description
	Name        string                                                                     `json:"name,omitempty"`        // Queueing profile name
	Clause      *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause `json:"clause,omitempty"`      //
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClause struct {
	InstanceID                        *int                                                                                                     `json:"instanceId,omitempty"`                        // Instance id
	Type                              string                                                                                                   `json:"type,omitempty"`                              // The allowed clause types are: BANDWIDTH, DSCP_CUSTOMIZATION
	IsCommonBetweenAllInterfaceSpeeds *bool                                                                                                    `json:"isCommonBetweenAllInterfaceSpeeds,omitempty"` // Is common between all interface speeds
	InterfaceSpeedBandwidthClauses    *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses `json:"interfaceSpeedBandwidthClauses,omitempty"`    //
	TcDscpSettings                    *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings                 `json:"tcDscpSettings,omitempty"`                    //
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses struct {
	InstanceID          *int                                                                                                                        `json:"instanceId,omitempty"`          // Instance id
	InterfaceSpeed      string                                                                                                                      `json:"interfaceSpeed,omitempty"`      // Interface speed
	TcBandwidthSettings *[]RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `json:"tcBandwidthSettings,omitempty"` //
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	InstanceID          *int   `json:"instanceId,omitempty"`          // Instance id
	BandwidthPercentage *int   `json:"bandwidthPercentage,omitempty"` // Bandwidth percentage
	TrafficClass        string `json:"trafficClass,omitempty"`        // Traffic Class
}
type RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileClauseTcDscpSettings struct {
	InstanceID   *int   `json:"instanceId,omitempty"`   // Instance id
	Dscp         string `json:"dscp,omitempty"`         // Dscp value
	TrafficClass string `json:"trafficClass,omitempty"` // Traffic Class
}
type RequestApplicationPolicyCreateApplicationPolicyQueuingProfile []RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile // Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfile
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfile struct {
	Description string                                                                     `json:"description,omitempty"` // Free test description
	Name        string                                                                     `json:"name,omitempty"`        // Queueing profile name
	Clause      *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause `json:"clause,omitempty"`      //
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClause struct {
	Type                              string                                                                                                   `json:"type,omitempty"`                              // The allowed clause types are: BANDWIDTH, DSCP_CUSTOMIZATION
	IsCommonBetweenAllInterfaceSpeeds *bool                                                                                                    `json:"isCommonBetweenAllInterfaceSpeeds,omitempty"` // Is common between all interface speeds
	InterfaceSpeedBandwidthClauses    *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses `json:"interfaceSpeedBandwidthClauses,omitempty"`    //
	TcDscpSettings                    *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings                 `json:"tcDscpSettings,omitempty"`                    //
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClauses struct {
	InterfaceSpeed      string                                                                                                                      `json:"interfaceSpeed,omitempty"`      // Interface speed
	TcBandwidthSettings *[]RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings `json:"tcBandwidthSettings,omitempty"` //
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings struct {
	BandwidthPercentage *int   `json:"bandwidthPercentage,omitempty"` // Bandwidth percentage
	TrafficClass        string `json:"trafficClass,omitempty"`        // Traffic Class
}
type RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileClauseTcDscpSettings struct {
	Dscp         string `json:"dscp,omitempty"`         // Dscp value
	TrafficClass string `json:"trafficClass,omitempty"` // Traffic Class
}
type RequestApplicationPolicyCreateApplicationSet []RequestItemApplicationPolicyCreateApplicationSet // Array of RequestApplicationPolicyCreateApplicationSet
type RequestItemApplicationPolicyCreateApplicationSet struct {
	Name string `json:"name,omitempty"` // Name
}
type RequestApplicationPolicyCreateApplication []RequestItemApplicationPolicyCreateApplication // Array of RequestApplicationPolicyCreateApplication
type RequestItemApplicationPolicyCreateApplication struct {
	Name                      string                                                                    `json:"name,omitempty"`                      // Name
	NetworkApplications       *[]RequestItemApplicationPolicyCreateApplicationNetworkApplications       `json:"networkApplications,omitempty"`       //
	NetworkIDentity           *[]RequestItemApplicationPolicyCreateApplicationNetworkIDentity           `json:"networkIdentity,omitempty"`           //
	ApplicationSet            *RequestItemApplicationPolicyCreateApplicationApplicationSet              `json:"applicationSet,omitempty"`            //
	IndicativeNetworkIDentity *[]RequestItemApplicationPolicyCreateApplicationIndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"` //
}
type RequestItemApplicationPolicyCreateApplicationIndicativeNetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // id
	DisplayName string `json:"displayName,omitempty"` // displayName
	LowerPort   *int   `json:"lowerPort,omitempty"`   // lowerPort
	Ports       string `json:"ports,omitempty"`       // ports
	Protocol    string `json:"protocol,omitempty"`    // protocol
	UpperPort   *int   `json:"upperPort,omitempty"`   // upperPort
}
type RequestItemApplicationPolicyCreateApplicationNetworkApplications struct {
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type RequestItemApplicationPolicyCreateApplicationNetworkIDentity struct {
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type RequestItemApplicationPolicyCreateApplicationApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type RequestApplicationPolicyEditApplication []RequestItemApplicationPolicyEditApplication // Array of RequestApplicationPolicyEditApplication
type RequestItemApplicationPolicyEditApplication struct {
	ID                  string                                                            `json:"id,omitempty"`                  // Id
	Name                string                                                            `json:"name,omitempty"`                // Name
	NetworkApplications *[]RequestItemApplicationPolicyEditApplicationNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]RequestItemApplicationPolicyEditApplicationNetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *RequestItemApplicationPolicyEditApplicationApplicationSet        `json:"applicationSet,omitempty"`      //
}
type RequestItemApplicationPolicyEditApplicationNetworkApplications struct {
	ID                 string `json:"id,omitempty"`                 // Id
	AppProtocol        string `json:"appProtocol,omitempty"`        // App Protocol
	ApplicationSubType string `json:"applicationSubType,omitempty"` // Application Sub Type
	ApplicationType    string `json:"applicationType,omitempty"`    // Application Type
	CategoryID         string `json:"categoryId,omitempty"`         // Category Id
	DisplayName        string `json:"displayName,omitempty"`        // Display Name
	EngineID           string `json:"engineId,omitempty"`           // Engine Id
	HelpString         string `json:"helpString,omitempty"`         // Help String
	LongDescription    string `json:"longDescription,omitempty"`    // Long Description
	Name               string `json:"name,omitempty"`               // Name
	Popularity         string `json:"popularity,omitempty"`         // Popularity
	Rank               string `json:"rank,omitempty"`               // Rank
	TrafficClass       string `json:"trafficClass,omitempty"`       // Traffic Class
	ServerName         string `json:"serverName,omitempty"`         // Server Name
	URL                string `json:"url,omitempty"`                // Url
	Dscp               string `json:"dscp,omitempty"`               // Dscp
	IgnoreConflict     string `json:"ignoreConflict,omitempty"`     // Ignore Conflict
}
type RequestItemApplicationPolicyEditApplicationNetworkIDentity struct {
	ID          string `json:"id,omitempty"`          // Id
	DisplayName string `json:"displayName,omitempty"` // Display Name
	LowerPort   string `json:"lowerPort,omitempty"`   // Lower Port
	Ports       string `json:"ports,omitempty"`       // Ports
	Protocol    string `json:"protocol,omitempty"`    // Protocol
	UpperPort   string `json:"upperPort,omitempty"`   // Upper Port
}
type RequestItemApplicationPolicyEditApplicationApplicationSet struct {
	IDRef string `json:"idRef,omitempty"` // Id Ref
}
type RequestApplicationPolicyUpdateQosDeviceInterfaceInfo []RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo // Array of RequestApplicationPolicyUpdateQosDeviceInterfaceInfo
type RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo struct {
	ID                     string                                                                            `json:"id,omitempty"`                     // Id of Qos device info
	Name                   string                                                                            `json:"name,omitempty"`                   // Device name
	ExcludedInterfaces     []string                                                                          `json:"excludedInterfaces,omitempty"`     // Excluded interfaces ids
	NetworkDeviceID        string                                                                            `json:"networkDeviceId,omitempty"`        // Network device id
	QosDeviceInterfaceInfo *[]RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo `json:"qosDeviceInterfaceInfo,omitempty"` //
}
type RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo struct {
	InstanceID         *int   `json:"instanceId,omitempty"`         // Instance id
	DmvpnRemoteSitesBw *[]int `json:"dmvpnRemoteSitesBw,omitempty"` // Dmvpn remote sites bandwidth
	InterfaceID        string `json:"interfaceId,omitempty"`        // Interface id
	InterfaceName      string `json:"interfaceName,omitempty"`      // Interface name
	Label              string `json:"label,omitempty"`              // SP Profile name
	Role               string `json:"role,omitempty"`               // Interface role
	UploadBW           *int   `json:"uploadBW,omitempty"`           // Upload bandwidth
}
type RequestApplicationPolicyCreateQosDeviceInterfaceInfo []RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo // Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfo
type RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo struct {
	Name                   string                                                                            `json:"name,omitempty"`                   // Device name
	ExcludedInterfaces     []string                                                                          `json:"excludedInterfaces,omitempty"`     // Excluded interfaces ids
	NetworkDeviceID        string                                                                            `json:"networkDeviceId,omitempty"`        // Network device id
	QosDeviceInterfaceInfo *[]RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo `json:"qosDeviceInterfaceInfo,omitempty"` //
}
type RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo struct {
	DmvpnRemoteSitesBw *[]int `json:"dmvpnRemoteSitesBw,omitempty"` // Dmvpn remote sites bandwidth
	InterfaceID        string `json:"interfaceId,omitempty"`        // Interface id
	InterfaceName      string `json:"interfaceName,omitempty"`      // Interface name
	Label              string `json:"label,omitempty"`              // SP Profile name
	Role               string `json:"role,omitempty"`               // Interface role
	UploadBW           *int   `json:"uploadBW,omitempty"`           // Upload bandwidth
}
type RequestApplicationPolicyCreateApplicationSetsV2 []RequestItemApplicationPolicyCreateApplicationSetsV2 // Array of RequestApplicationPolicyCreateApplicationSetsV2
type RequestItemApplicationPolicyCreateApplicationSetsV2 struct {
	Name                        string `json:"name,omitempty"`                        // Application Set name
	ScalableGroupType           string `json:"scalableGroupType,omitempty"`           // Scalable group type, should be set to APPLICATION_GROUP
	DefaultBusinessRelevance    string `json:"defaultBusinessRelevance,omitempty"`    // Default business relevance
	Namespace                   string `json:"namespace,omitempty"`                   // Namespace, should be set to scalablegroup:application
	Qualifier                   string `json:"qualifier,omitempty"`                   // Qualifier, should be set to application
	Type                        string `json:"type,omitempty"`                        // Type, should be set to scalablegroup
	ScalableGroupExternalHandle string `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be set to application set name
}
type RequestApplicationPolicyEditApplicationsV2 []RequestItemApplicationPolicyEditApplicationsV2 // Array of RequestApplicationPolicyEditApplicationsV2
type RequestItemApplicationPolicyEditApplicationsV2 struct {
	ID                          string                                                                     `json:"id,omitempty"`                          // Application id
	InstanceID                  *int                                                                       `json:"instanceId,omitempty"`                  // Instance id
	DisplayName                 string                                                                     `json:"displayName,omitempty"`                 // Display name
	InstanceVersion             *float64                                                                   `json:"instanceVersion,omitempty"`             // Instance version
	IndicativeNetworkIDentity   *[]RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"`   //
	Name                        string                                                                     `json:"name,omitempty"`                        // Application name
	Namespace                   string                                                                     `json:"namespace,omitempty"`                   // Namespace
	NetworkApplications         *[]RequestItemApplicationPolicyEditApplicationsV2NetworkApplications       `json:"networkApplications,omitempty"`         //
	NetworkIDentity             *[]RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity           `json:"networkIdentity,omitempty"`             //
	ParentScalableGroup         *RequestItemApplicationPolicyEditApplicationsV2ParentScalableGroup         `json:"parentScalableGroup,omitempty"`         //
	Qualifier                   string                                                                     `json:"qualifier,omitempty"`                   // Qualifier, valid value application
	ScalableGroupExternalHandle string                                                                     `json:"scalableGroupExternalHandle,omitempty"` // Scalable group external handle, should be equal to Application name
	ScalableGroupType           string                                                                     `json:"scalableGroupType,omitempty"`           // Scalable group type, valid value APPLICATION
	Type                        string                                                                     `json:"type,omitempty"`                        // Type, valid value scalablegroup
}
type RequestItemApplicationPolicyEditApplicationsV2IndicativeNetworkIDentity struct {
	ID          string   `json:"id,omitempty"`          // Id
	DisplayName string   `json:"displayName,omitempty"` // Display name
	LowerPort   *float64 `json:"lowerPort,omitempty"`   // Lower port
	Ports       string   `json:"ports,omitempty"`       // Ports
	Protocol    string   `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64 `json:"upperPort,omitempty"`   // Upper port
}
type RequestItemApplicationPolicyEditApplicationsV2NetworkApplications struct {
	ID                 string   `json:"id,omitempty"`                 // Id
	ApplicationSubType string   `json:"applicationSubType,omitempty"` // Application sub type, LEARNED: discovered application, NONE: nbar and custom application
	ApplicationType    string   `json:"applicationType,omitempty"`    // Application type, DEFAULT: nbar application, DEFAULT_MODIFIED: nbar modified application, CUSTOM: custom application
	CategoryID         string   `json:"categoryId,omitempty"`         // Category id
	DisplayName        string   `json:"displayName,omitempty"`        // Display name
	EngineID           string   `json:"engineId,omitempty"`           // Engine id
	HelpString         string   `json:"helpString,omitempty"`         // Help string
	LongDescription    string   `json:"longDescription,omitempty"`    // Long description
	Name               string   `json:"name,omitempty"`               // Application name
	Popularity         *float64 `json:"popularity,omitempty"`         // Popularity
	Rank               *int     `json:"rank,omitempty"`               // Rank, any value between 1 to 65535
	SelectorID         string   `json:"selectorId,omitempty"`         // Selector id
	Dscp               string   `json:"dscp,omitempty"`               // Dscp
	AppProtocol        string   `json:"appProtocol,omitempty"`        // App protocol
	ServerName         string   `json:"serverName,omitempty"`         // Server name
	URL                string   `json:"url,omitempty"`                // Url
	TrafficClass       string   `json:"trafficClass,omitempty"`       // Traffic class
	IgnoreConflict     *bool    `json:"ignoreConflict,omitempty"`     // Ignore conflict, true or false
}
type RequestItemApplicationPolicyEditApplicationsV2NetworkIDentity struct {
	ID          string                                                                     `json:"id,omitempty"`          // Id
	DisplayName string                                                                     `json:"displayName,omitempty"` // Display name
	IPv4Subnet  []string                                                                   `json:"ipv4Subnet,omitempty"`  // Ipv4 subnet
	IPv6Subnet  *[]RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet `json:"ipv6Subnet,omitempty"`  // Ipv6 subnet
	LowerPort   *float64                                                                   `json:"lowerPort,omitempty"`   // Lower port
	Ports       string                                                                     `json:"ports,omitempty"`       // Ports
	Protocol    string                                                                     `json:"protocol,omitempty"`    // Protocol
	UpperPort   *float64                                                                   `json:"upperPort,omitempty"`   // Upper port
}
type RequestItemApplicationPolicyEditApplicationsV2NetworkIDentityIPv6Subnet interface{}
type RequestItemApplicationPolicyEditApplicationsV2ParentScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id reference to parent application set
}
type RequestApplicationPolicyCreateApplicationsV2 []RequestItemApplicationPolicyCreateApplicationsV2 // Array of RequestApplicationPolicyCreateApplicationsV2
type RequestItemApplicationPolicyCreateApplicationsV2 struct {
	Name                      string                                                                       `json:"name,omitempty"`                      // Application name
	NetworkApplications       *[]RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications       `json:"networkApplications,omitempty"`       //
	ParentScalableGroup       *RequestItemApplicationPolicyCreateApplicationsV2ParentScalableGroup         `json:"parentScalableGroup,omitempty"`       //
	NetworkIDentity           *[]RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity           `json:"networkIdentity,omitempty"`           //
	IndicativeNetworkIDentity *[]RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity `json:"indicativeNetworkIdentity,omitempty"` //
	ScalableGroupType         string                                                                       `json:"scalableGroupType,omitempty"`         // Scalable group type, valid value APPLICATION
	Type                      string                                                                       `json:"type,omitempty"`                      // Type, valid value scalablegroup
}
type RequestItemApplicationPolicyCreateApplicationsV2NetworkApplications struct {
	HelpString      string `json:"helpString,omitempty"`      // Help string
	ApplicationType string `json:"applicationType,omitempty"` // Application type
	Type            string `json:"type,omitempty"`            // Custom application type
	Dscp            string `json:"dscp,omitempty"`            // Dscp, valid only in case of _server-ip custom application type
	AppProtocol     string `json:"appProtocol,omitempty"`     // App protocol, in case of _servername should not be set, in case of _url should be set to TCP
	ServerName      string `json:"serverName,omitempty"`      // Server name, should be set only in case of _servername
	URL             string `json:"url,omitempty"`             // Url, should be set only in case of _url
	TrafficClass    string `json:"trafficClass,omitempty"`    // Traffic class
	CategoryID      string `json:"categoryId,omitempty"`      // Category id
	IgnoreConflict  *bool  `json:"ignoreConflict,omitempty"`  // Ignore conflict, true or false
	Rank            *int   `json:"rank,omitempty"`            // Rank, should be set to 1
	EngineID        *int   `json:"engineId,omitempty"`        // Engine id, should be set to 6
}
type RequestItemApplicationPolicyCreateApplicationsV2ParentScalableGroup struct {
	IDRef string `json:"idRef,omitempty"` // Id reference to parent application set
}
type RequestItemApplicationPolicyCreateApplicationsV2NetworkIDentity struct {
	Protocol   string   `json:"protocol,omitempty"`   // Protocol
	Ports      string   `json:"ports,omitempty"`      // Ports
	IPv4Subnet []string `json:"ipv4Subnet,omitempty"` // Ipv4 subnet
	LowerPort  *float64 `json:"lowerPort,omitempty"`  // Lower port
	UpperPort  *float64 `json:"upperPort,omitempty"`  // Upper port
}
type RequestItemApplicationPolicyCreateApplicationsV2IndicativeNetworkIDentity struct {
	Protocol   string   `json:"protocol,omitempty"`   // Protocol
	Ports      string   `json:"ports,omitempty"`      // Ports
	IPv4Subnet []string `json:"ipv4Subnet,omitempty"` // Ipv4 subnet
	IPv6Subnet []string `json:"ipv6Subnet,omitempty"` // Ipv6 subnet
	LowerPort  *float64 `json:"lowerPort,omitempty"`  // The minimum port when used as a port range. For single port number, ports attribute should be used.
	UpperPort  *float64 `json:"upperPort,omitempty"`  // The maximum port when used as a port range. For single port number, ports attribute should be used.
}

//GetApplicationPolicy Get Application Policy - 3d9f-6b17-4879-8e45
/* Get all existing application policies


@param GetApplicationPolicyQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicy(GetApplicationPolicyQueryParams *GetApplicationPolicyQueryParams) (*ResponseApplicationPolicyGetApplicationPolicy, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy"

	queryString, _ := query.Values(GetApplicationPolicyQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationPolicy{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicy(GetApplicationPolicyQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicy")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicy)
	return result, response, err

}

//GetApplicationPolicyDefault Get Application Policy Default - 21a2-4b5d-4f98-a730
/* Get default application policy



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-default-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyDefault() (*ResponseApplicationPolicyGetApplicationPolicyDefault, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-default"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationPolicyDefault{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyDefault()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyDefault")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyDefault)
	return result, response, err

}

//GetApplicationPolicyQueuingProfile Get Application Policy Queuing Profile - 1698-39cc-4bdb-8ed9
/* Get all or by name, existing application policy queuing profiles


@param GetApplicationPolicyQueuingProfileQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-queuing-profile-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyQueuingProfile(GetApplicationPolicyQueuingProfileQueryParams *GetApplicationPolicyQueuingProfileQueryParams) (*ResponseApplicationPolicyGetApplicationPolicyQueuingProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile"

	queryString, _ := query.Values(GetApplicationPolicyQueuingProfileQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationPolicyQueuingProfile{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyQueuingProfile(GetApplicationPolicyQueuingProfileQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyQueuingProfile)
	return result, response, err

}

//GetApplicationPolicyQueuingProfileCount Get Application Policy Queuing Profile Count - efa2-4afa-4578-88e9
/* Get the number of all existing  application policy queuing profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-policy-queuing-profile-count-v1
*/
func (s *ApplicationPolicyService) GetApplicationPolicyQueuingProfileCount() (*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationPolicyQueuingProfileCount()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyQueuingProfileCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCount)
	return result, response, err

}

//GetApplicationSets Get Application Sets - cb86-8b21-4289-8159
/* Get appllication-sets by offset/limit or by name


@param GetApplicationSetsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-sets-v1
*/
func (s *ApplicationPolicyService) GetApplicationSets(GetApplicationSetsQueryParams *GetApplicationSetsQueryParams) (*ResponseApplicationPolicyGetApplicationSets, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(GetApplicationSetsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSets{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSets(GetApplicationSetsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSets")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSets)
	return result, response, err

}

//GetApplicationSetsCount Get Application Sets Count - cfa0-49a6-44bb-8a07
/* Get the number of existing application-sets



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-sets-count-v1
*/
func (s *ApplicationPolicyService) GetApplicationSetsCount() (*ResponseApplicationPolicyGetApplicationSetsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationSetsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetsCount()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsCount)
	return result, response, err

}

//GetApplications Get Applications - 8893-b834-445b-b29c
/* Get applications by offset/limit or by name


@param GetApplicationsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-applications-v1
*/
func (s *ApplicationPolicyService) GetApplications(GetApplicationsQueryParams *GetApplicationsQueryParams) (*ResponseApplicationPolicyGetApplications, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(GetApplicationsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplications{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplications(GetApplicationsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplications")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplications)
	return result, response, err

}

//GetApplicationsCount Get Applications Count - 039d-e8b1-47a9-8690
/* Get the number of all existing applications



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-applications-count-v1
*/
func (s *ApplicationPolicyService) GetApplicationsCount() (*ResponseApplicationPolicyGetApplicationsCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetApplicationsCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationsCount()
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationsCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsCount)
	return result, response, err

}

//GetQosDeviceInterfaceInfo Get Qos Device Interface Info - 42a6-e9eb-46bb-a197
/* Get all or by network device id, existing qos device interface infos


@param GetQosDeviceInterfaceInfoQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-qos-device-interface-info-v1
*/
func (s *ApplicationPolicyService) GetQosDeviceInterfaceInfo(GetQosDeviceInterfaceInfoQueryParams *GetQosDeviceInterfaceInfoQueryParams) (*ResponseApplicationPolicyGetQosDeviceInterfaceInfo, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info"

	queryString, _ := query.Values(GetQosDeviceInterfaceInfoQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetQosDeviceInterfaceInfo{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetQosDeviceInterfaceInfo(GetQosDeviceInterfaceInfoQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyGetQosDeviceInterfaceInfo)
	return result, response, err

}

//GetQosDeviceInterfaceInfoCount Get Qos Device Interface Info Count - 729a-98e1-4a3b-91f2
/* Get the number of all existing qos device interface infos group by network device id



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-qos-device-interface-info-count-v1
*/
func (s *ApplicationPolicyService) GetQosDeviceInterfaceInfoCount() (*ResponseApplicationPolicyGetQosDeviceInterfaceInfoCount, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info-count"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyGetQosDeviceInterfaceInfoCount{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetQosDeviceInterfaceInfoCount()
		}
		return nil, response, fmt.Errorf("error with operation GetQosDeviceInterfaceInfoCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetQosDeviceInterfaceInfoCount)
	return result, response, err

}

//GetApplicationSetsV2 Get Application Set/s - 00ac-d849-43aa-bc75
/* Get application set/s by offset/limit or by name


@param GetApplicationSetsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-sets-v2
*/
func (s *ApplicationPolicyService) GetApplicationSetsV2(GetApplicationSetsV2QueryParams *GetApplicationSetsV2QueryParams) (*ResponseApplicationPolicyGetApplicationSetsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/application-policy-application-set"

	queryString, _ := query.Values(GetApplicationSetsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSetsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetsV2(GetApplicationSetsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsV2)
	return result, response, err

}

//GetApplicationSetCountV2 Get Application Set Count - 798c-fa61-432b-a67e
/* Get the number of all existing application sets


@param GetApplicationSetCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-set-count-v2
*/
func (s *ApplicationPolicyService) GetApplicationSetCountV2(GetApplicationSetCountV2QueryParams *GetApplicationSetCountV2QueryParams) (*ResponseApplicationPolicyGetApplicationSetCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/application-policy-application-set-count"

	queryString, _ := query.Values(GetApplicationSetCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationSetCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationSetCountV2(GetApplicationSetCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationSetCountV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetCountV2)
	return result, response, err

}

//GetApplicationsV2 Get Application/s - a395-8970-43bb-bedd
/* Get application/s by offset/limit or by name


@param GetApplicationsV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-applications-v2
*/
func (s *ApplicationPolicyService) GetApplicationsV2(GetApplicationsV2QueryParams *GetApplicationsV2QueryParams) (*ResponseApplicationPolicyGetApplicationsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications"

	queryString, _ := query.Values(GetApplicationsV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationsV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationsV2(GetApplicationsV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsV2)
	return result, response, err

}

//GetApplicationCountV2 Get Application Count - cfa3-7a5c-4c08-b24e
/* Get the number of all existing applications


@param GetApplicationCountV2QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-application-count-v2
*/
func (s *ApplicationPolicyService) GetApplicationCountV2(GetApplicationCountV2QueryParams *GetApplicationCountV2QueryParams) (*ResponseApplicationPolicyGetApplicationCountV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications-count"

	queryString, _ := query.Values(GetApplicationCountV2QueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyGetApplicationCountV2{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetApplicationCountV2(GetApplicationCountV2QueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetApplicationCountV2")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationCountV2)
	return result, response, err

}

//ApplicationPolicyIntent Application Policy Intent - aea4-bb7b-4329-bd06
/* Create/Update/Delete application policy



Documentation Link: https://developer.cisco.com/docs/dna-center/#!application-policy-intent-v1
*/
func (s *ApplicationPolicyService) ApplicationPolicyIntent(requestApplicationPolicyApplicationPolicyIntent *RequestApplicationPolicyApplicationPolicyIntent) (*ResponseApplicationPolicyApplicationPolicyIntent, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-intent"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyApplicationPolicyIntent).
		SetResult(&ResponseApplicationPolicyApplicationPolicyIntent{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApplicationPolicyIntent(requestApplicationPolicyApplicationPolicyIntent)
		}

		return nil, response, fmt.Errorf("error with operation ApplicationPolicyIntent")
	}

	result := response.Result().(*ResponseApplicationPolicyApplicationPolicyIntent)
	return result, response, err

}

//CreateApplicationPolicyQueuingProfile Create Application Policy Queuing Profile - 78b7-ba71-4959-bbd2
/* Create new custom application queuing profile



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-policy-queuing-profile-v1
*/
func (s *ApplicationPolicyService) CreateApplicationPolicyQueuingProfile(requestApplicationPolicyCreateApplicationPolicyQueuingProfile *RequestApplicationPolicyCreateApplicationPolicyQueuingProfile) (*ResponseApplicationPolicyCreateApplicationPolicyQueuingProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationPolicyQueuingProfile).
		SetResult(&ResponseApplicationPolicyCreateApplicationPolicyQueuingProfile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationPolicyQueuingProfile(requestApplicationPolicyCreateApplicationPolicyQueuingProfile)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationPolicyQueuingProfile)
	return result, response, err

}

//CreateApplicationSet Create Application Set - 3e94-cb1b-485b-8b0e
/* Create new custom application-set/s



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-set-v1
*/
func (s *ApplicationPolicyService) CreateApplicationSet(requestApplicationPolicyCreateApplicationSet *RequestApplicationPolicyCreateApplicationSet) (*ResponseApplicationPolicyCreateApplicationSet, *resty.Response, error) {
	path := "/dna/intent/api/v1/application-policy-application-set"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationSet).
		SetResult(&ResponseApplicationPolicyCreateApplicationSet{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationSet(requestApplicationPolicyCreateApplicationSet)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationSet")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationSet)
	return result, response, err

}

//CreateApplication Create Application - fb9b-f80f-491a-9851
/* Create new Custom application



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-v1
*/
func (s *ApplicationPolicyService) CreateApplication(requestApplicationPolicyCreateApplication *RequestApplicationPolicyCreateApplication) (*ResponseApplicationPolicyCreateApplication, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplication).
		SetResult(&ResponseApplicationPolicyCreateApplication{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplication(requestApplicationPolicyCreateApplication)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplication)
	return result, response, err

}

//CreateQosDeviceInterfaceInfo Create Qos Device Interface Info - 3889-59af-4cf8-9fde
/* Create qos device interface infos associate with network device id to allow the user to mark specific interfaces as WAN, to associate WAN interfaces with specific SP Profile and to be able to define a shaper on WAN interfaces



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-qos-device-interface-info-v1
*/
func (s *ApplicationPolicyService) CreateQosDeviceInterfaceInfo(requestApplicationPolicyCreateQosDeviceInterfaceInfo *RequestApplicationPolicyCreateQosDeviceInterfaceInfo) (*ResponseApplicationPolicyCreateQosDeviceInterfaceInfo, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateQosDeviceInterfaceInfo).
		SetResult(&ResponseApplicationPolicyCreateQosDeviceInterfaceInfo{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateQosDeviceInterfaceInfo(requestApplicationPolicyCreateQosDeviceInterfaceInfo)
		}

		return nil, response, fmt.Errorf("error with operation CreateQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateQosDeviceInterfaceInfo)
	return result, response, err

}

//CreateApplicationSetsV2 Create Application Set/s - e4bf-ca74-45f9-a374
/* Create new custom application set/s



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-application-sets-v2
*/
func (s *ApplicationPolicyService) CreateApplicationSetsV2(requestApplicationPolicyCreateApplicationSetsV2 *RequestApplicationPolicyCreateApplicationSetsV2) (*ResponseApplicationPolicyCreateApplicationSetsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/application-policy-application-set"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationSetsV2).
		SetResult(&ResponseApplicationPolicyCreateApplicationSetsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationSetsV2(requestApplicationPolicyCreateApplicationSetsV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationSetsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationSetsV2)
	return result, response, err

}

//CreateApplicationsV2 Create Application/s - b4a6-dae7-4b29-992c
/* Create new custom application/s



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-applications-v2
*/
func (s *ApplicationPolicyService) CreateApplicationsV2(requestApplicationPolicyCreateApplicationsV2 *RequestApplicationPolicyCreateApplicationsV2) (*ResponseApplicationPolicyCreateApplicationsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyCreateApplicationsV2).
		SetResult(&ResponseApplicationPolicyCreateApplicationsV2{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateApplicationsV2(requestApplicationPolicyCreateApplicationsV2)
		}

		return nil, response, fmt.Errorf("error with operation CreateApplicationsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationsV2)
	return result, response, err

}

//UpdateApplicationPolicyQueuingProfile Update Application Policy Queuing Profile - da98-38ea-44aa-8024
/* Update existing custom application queuing profile


 */
func (s *ApplicationPolicyService) UpdateApplicationPolicyQueuingProfile(requestApplicationPolicyUpdateApplicationPolicyQueuingProfile *RequestApplicationPolicyUpdateApplicationPolicyQueuingProfile) (*ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfile, *resty.Response, error) {
	path := "/dna/intent/api/v1/app-policy-queuing-profile"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyUpdateApplicationPolicyQueuingProfile).
		SetResult(&ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfile{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateApplicationPolicyQueuingProfile(requestApplicationPolicyUpdateApplicationPolicyQueuingProfile)
		}
		return nil, response, fmt.Errorf("error with operation UpdateApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyUpdateApplicationPolicyQueuingProfile)
	return result, response, err

}

//EditApplication Edit Application - 3986-6887-4439-a41d
/* Edit the attributes of an existing application


 */
func (s *ApplicationPolicyService) EditApplication(requestApplicationPolicyEditApplication *RequestApplicationPolicyEditApplication) (*ResponseApplicationPolicyEditApplication, *resty.Response, error) {
	path := "/dna/intent/api/v1/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyEditApplication).
		SetResult(&ResponseApplicationPolicyEditApplication{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditApplication(requestApplicationPolicyEditApplication)
		}
		return nil, response, fmt.Errorf("error with operation EditApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyEditApplication)
	return result, response, err

}

//UpdateQosDeviceInterfaceInfo Update Qos Device Interface Info - 71b7-ba8c-47b8-95b6
/* Update existing qos device interface infos associate with network device id


 */
func (s *ApplicationPolicyService) UpdateQosDeviceInterfaceInfo(requestApplicationPolicyUpdateQosDeviceInterfaceInfo *RequestApplicationPolicyUpdateQosDeviceInterfaceInfo) (*ResponseApplicationPolicyUpdateQosDeviceInterfaceInfo, *resty.Response, error) {
	path := "/dna/intent/api/v1/qos-device-interface-info"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyUpdateQosDeviceInterfaceInfo).
		SetResult(&ResponseApplicationPolicyUpdateQosDeviceInterfaceInfo{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateQosDeviceInterfaceInfo(requestApplicationPolicyUpdateQosDeviceInterfaceInfo)
		}
		return nil, response, fmt.Errorf("error with operation UpdateQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyUpdateQosDeviceInterfaceInfo)
	return result, response, err

}

//EditApplicationsV2 Edit Application/s - 6995-2aea-4f2b-a053
/* Edit the attributes of an existing application


 */
func (s *ApplicationPolicyService) EditApplicationsV2(requestApplicationPolicyEditApplicationsV2 *RequestApplicationPolicyEditApplicationsV2) (*ResponseApplicationPolicyEditApplicationsV2, *resty.Response, error) {
	path := "/dna/intent/api/v2/applications"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestApplicationPolicyEditApplicationsV2).
		SetResult(&ResponseApplicationPolicyEditApplicationsV2{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.EditApplicationsV2(requestApplicationPolicyEditApplicationsV2)
		}
		return nil, response, fmt.Errorf("error with operation EditApplicationsV2")
	}

	result := response.Result().(*ResponseApplicationPolicyEditApplicationsV2)
	return result, response, err

}

//DeleteApplicationPolicyQueuingProfile Delete Application Policy Queuing Profile - 09a0-482f-422b-b325
/* Delete existing custom application policy queuing profile by id


@param id id path parameter. Id of custom queuing profile to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-policy-queuing-profile-v1
*/
func (s *ApplicationPolicyService) DeleteApplicationPolicyQueuingProfile(id string) (*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfile, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/app-policy-queuing-profile/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfile{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationPolicyQueuingProfile(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfile)
	return result, response, err

}

//DeleteApplicationSet Delete Application Set - 70b6-f8e1-40b8-b784
/* Delete existing application-set by it's id


@param DeleteApplicationSetQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-set-v1
*/
func (s *ApplicationPolicyService) DeleteApplicationSet(DeleteApplicationSetQueryParams *DeleteApplicationSetQueryParams) (*ResponseApplicationPolicyDeleteApplicationSet, *resty.Response, error) {
	//DeleteApplicationSetQueryParams *DeleteApplicationSetQueryParams
	path := "/dna/intent/api/v1/application-policy-application-set"

	queryString, _ := query.Values(DeleteApplicationSetQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyDeleteApplicationSet{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationSet(DeleteApplicationSetQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationSet")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationSet)
	return result, response, err

}

//DeleteApplication Delete Application - d49a-f9b8-4c6a-a8ea
/* Delete existing application by its id


@param DeleteApplicationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-v1
*/
func (s *ApplicationPolicyService) DeleteApplication(DeleteApplicationQueryParams *DeleteApplicationQueryParams) (*ResponseApplicationPolicyDeleteApplication, *resty.Response, error) {
	//DeleteApplicationQueryParams *DeleteApplicationQueryParams
	path := "/dna/intent/api/v1/applications"

	queryString, _ := query.Values(DeleteApplicationQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseApplicationPolicyDeleteApplication{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplication(DeleteApplicationQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplication)
	return result, response, err

}

//DeleteQosDeviceInterfaceInfo Delete Qos Device Interface Info - 51b1-b8bd-435b-9842
/* Delete all qos device interface infos associate with network device id


@param id id path parameter. Id of the qos device info, this object holds all qos device interface infos associate with network device id


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-qos-device-interface-info-v1
*/
func (s *ApplicationPolicyService) DeleteQosDeviceInterfaceInfo(id string) (*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfo, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/qos-device-interface-info/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteQosDeviceInterfaceInfo{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteQosDeviceInterfaceInfo(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfo)
	return result, response, err

}

//DeleteApplicationSetV2 Delete Application Set - b2a9-08f1-4879-9f70
/* Delete existing custom application set by id


@param id id path parameter. Id of custom application set to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-set-v2
*/
func (s *ApplicationPolicyService) DeleteApplicationSetV2(id string) (*ResponseApplicationPolicyDeleteApplicationSetV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/application-policy-application-set/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteApplicationSetV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationSetV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationSetV2")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationSetV2)
	return result, response, err

}

//DeleteApplicationV2 Delete Application - 9098-8ada-4abb-bedc
/* Delete existing custom application by id


@param id id path parameter. Id of custom application to delete


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-application-v2
*/
func (s *ApplicationPolicyService) DeleteApplicationV2(id string) (*ResponseApplicationPolicyDeleteApplicationV2, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v2/applications/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseApplicationPolicyDeleteApplicationV2{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteApplicationV2(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteApplicationV2")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationV2)
	return result, response, err

}
