package dnac

import (
	"fmt"
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
	Type              string   `json:"type,omitempty"`              // Type. (Example: BUSINESS_RELEVANCE.)
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
	Response *int   `json:"response,omitempty"` // Response
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
	GroupID []string `json:"groupId,omitempty"` // Group id
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
	GroupID []string `json:"groupId,omitempty"` // Group id
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
	Type                              string                                                                                                   `json:"type,omitempty"`                              // Type
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
	Type                              string                                                                                                   `json:"type,omitempty"`                              // Type
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
	Name                string                                                              `json:"name,omitempty"`                // Name
	NetworkApplications *[]RequestItemApplicationPolicyCreateApplicationNetworkApplications `json:"networkApplications,omitempty"` //
	NetworkIDentity     *[]RequestItemApplicationPolicyCreateApplicationNetworkIDentity     `json:"networkIdentity,omitempty"`     //
	ApplicationSet      *RequestItemApplicationPolicyCreateApplicationApplicationSet        `json:"applicationSet,omitempty"`      //
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

//GetApplicationPolicy Get Application Policy - 3d9f-6b17-4879-8e45
/* Get all existing application policies


@param GetApplicationPolicyQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicy")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicy)
	return result, response, err

}

//GetApplicationPolicyDefault Get Application Policy Default - 21a2-4b5d-4f98-a730
/* Get default application policy


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
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyDefault")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyDefault)
	return result, response, err

}

//GetApplicationPolicyQueuingProfile Get Application Policy Queuing Profile - 1698-39cc-4bdb-8ed9
/* Get all or by name, existing application policy queuing profiles


@param GetApplicationPolicyQueuingProfileQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyQueuingProfile)
	return result, response, err

}

//GetApplicationPolicyQueuingProfileCount Get Application Policy Queuing Profile Count - efa2-4afa-4578-88e9
/* Get the number of all existing  application policy queuing profile


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
		return nil, response, fmt.Errorf("error with operation GetApplicationPolicyQueuingProfileCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationPolicyQueuingProfileCount)
	return result, response, err

}

//GetApplicationSets Get Application Sets - cb86-8b21-4289-8159
/* Get appllication-sets by offset/limit or by name


@param GetApplicationSetsQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetApplicationSets")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSets)
	return result, response, err

}

//GetApplicationSetsCount Get Application Sets Count - cfa0-49a6-44bb-8a07
/* Get the number of existing application-sets


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
		return nil, response, fmt.Errorf("error with operation GetApplicationSetsCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationSetsCount)
	return result, response, err

}

//GetApplications Get Applications - 8893-b834-445b-b29c
/* Get applications by offset/limit or by name


@param GetApplicationsQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetApplications")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplications)
	return result, response, err

}

//GetApplicationsCount Get Applications Count - 039d-e8b1-47a9-8690
/* Get the number of all existing applications


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
		return nil, response, fmt.Errorf("error with operation GetApplicationsCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetApplicationsCount)
	return result, response, err

}

//GetQosDeviceInterfaceInfo Get Qos Device Interface Info - 42a6-e9eb-46bb-a197
/* Get all or by network device id, existing qos device interface infos


@param GetQosDeviceInterfaceInfoQueryParams Filtering parameter
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
		return nil, response, fmt.Errorf("error with operation GetQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyGetQosDeviceInterfaceInfo)
	return result, response, err

}

//GetQosDeviceInterfaceInfoCount Get Qos Device Interface Info Count - 729a-98e1-4a3b-91f2
/* Get the number of all existing qos device interface infos group by network device id


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
		return nil, response, fmt.Errorf("error with operation GetQosDeviceInterfaceInfoCount")
	}

	result := response.Result().(*ResponseApplicationPolicyGetQosDeviceInterfaceInfoCount)
	return result, response, err

}

//ApplicationPolicyIntent Application Policy Intent - aea4-bb7b-4329-bd06
/* Create/Update/Delete application policy


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
		return nil, response, fmt.Errorf("error with operation ApplicationPolicyIntent")
	}

	result := response.Result().(*ResponseApplicationPolicyApplicationPolicyIntent)
	return result, response, err

}

//CreateApplicationPolicyQueuingProfile Create Application Policy Queuing Profile - 78b7-ba71-4959-bbd2
/* Create new custom application queuing profile


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
		return nil, response, fmt.Errorf("error with operation CreateApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationPolicyQueuingProfile)
	return result, response, err

}

//CreateApplicationSet Create Application Set - 3e94-cb1b-485b-8b0e
/* Create new custom application-set/s


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
		return nil, response, fmt.Errorf("error with operation CreateApplicationSet")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplicationSet)
	return result, response, err

}

//CreateApplication Create Application - fb9b-f80f-491a-9851
/* Create new Custom application


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
		return nil, response, fmt.Errorf("error with operation CreateApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateApplication)
	return result, response, err

}

//CreateQosDeviceInterfaceInfo Create Qos Device Interface Info - 3889-59af-4cf8-9fde
/* Create qos device interface infos associate with network device id to allow the user to mark specific interfaces as WAN, to associate WAN interfaces with specific SP Profile and to be able to define a shaper on WAN interfaces


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
		return nil, response, fmt.Errorf("error with operation CreateQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyCreateQosDeviceInterfaceInfo)
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
		return nil, response, fmt.Errorf("error with operation UpdateQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyUpdateQosDeviceInterfaceInfo)
	return result, response, err

}

//DeleteApplicationPolicyQueuingProfile Delete Application Policy Queuing Profile - 09a0-482f-422b-b325
/* Delete existing custom application policy queuing profile by id


@param id id path parameter. Id of custom queuing profile to delete

*/
func (s *ApplicationPolicyService) DeleteApplicationPolicyQueuingProfile(id string) (*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfile, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteApplicationPolicyQueuingProfile")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationPolicyQueuingProfile)
	return result, response, err

}

//DeleteApplicationSet Delete Application Set - 70b6-f8e1-40b8-b784
/* Delete existing application-set by it's id


@param DeleteApplicationSetQueryParams Filtering parameter
*/
func (s *ApplicationPolicyService) DeleteApplicationSet(DeleteApplicationSetQueryParams *DeleteApplicationSetQueryParams) (*ResponseApplicationPolicyDeleteApplicationSet, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteApplicationSet")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplicationSet)
	return result, response, err

}

//DeleteApplication Delete Application - d49a-f9b8-4c6a-a8ea
/* Delete existing application by its id


@param DeleteApplicationQueryParams Filtering parameter
*/
func (s *ApplicationPolicyService) DeleteApplication(DeleteApplicationQueryParams *DeleteApplicationQueryParams) (*ResponseApplicationPolicyDeleteApplication, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteApplication")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteApplication)
	return result, response, err

}

//DeleteQosDeviceInterfaceInfo Delete Qos Device Interface Info - 51b1-b8bd-435b-9842
/* Delete all qos device interface infos associate with network device id


@param id id path parameter. Id of the qos device info, this object holds all qos device interface infos associate with network device id

*/
func (s *ApplicationPolicyService) DeleteQosDeviceInterfaceInfo(id string) (*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfo, *resty.Response, error) {
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
		return nil, response, fmt.Errorf("error with operation DeleteQosDeviceInterfaceInfo")
	}

	result := response.Result().(*ResponseApplicationPolicyDeleteQosDeviceInterfaceInfo)
	return result, response, err

}
