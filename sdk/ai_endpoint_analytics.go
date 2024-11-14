package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AIEndpointAnalyticsService service

type GetAIEndpointAnalyticsAttributeDictionariesQueryParams struct {
	IncludeAttributes bool `url:"includeAttributes,omitempty"` //Flag to indicate whether attribute list for each dictionary should be included in response.
}
type QueryTheEndpointsQueryParams struct {
	ProfilingStatus          string   `url:"profilingStatus,omitempty"`          //Profiling status of the endpoint. Possible values are 'profiled', 'partialProfiled', 'notProfiled'.
	MacAddress               string   `url:"macAddress,omitempty"`               //MAC address to search for. Partial string is allowed.
	MacAddresses             []string `url:"macAddresses,omitempty"`             //List of MAC addresses to filter on. Only exact matches will be returned.
	IP                       string   `url:"ip,omitempty"`                       //IP address to search for. Partial string is allowed.
	DeviceType               string   `url:"deviceType,omitempty"`               //Type of device to search for. Partial string is allowed.
	HardwareManufacturer     string   `url:"hardwareManufacturer,omitempty"`     //Hardware manufacturer to search for. Partial string is allowed.
	HardwareModel            string   `url:"hardwareModel,omitempty"`            //Hardware model to search for. Partial string is allowed.
	OperatingSystem          string   `url:"operatingSystem,omitempty"`          //Operating system to search for. Partial string is allowed.
	Registered               bool     `url:"registered,omitempty"`               //Flag to fetch manually registered or non-registered endpoints.
	RandomMac                bool     `url:"randomMac,omitempty"`                //Flag to fetch endpoints having randomized MAC or not.
	TrustScore               string   `url:"trustScore,omitempty"`               //Overall trust score of the endpoint. It can be provided either as a number value (e.g. 5), or as a range (e.g. 3-7). Provide value as '-' if you want to search for all endpoints where trust score is not assigned.
	AuthMethod               string   `url:"authMethod,omitempty"`               //Authentication method. Partial string is allowed.
	PostureStatus            string   `url:"postureStatus,omitempty"`            //Posture status.
	AiSpoofingTrustLevel     string   `url:"aiSpoofingTrustLevel,omitempty"`     //Trust level of the endpoint due to AI spoofing. Possible values are 'low', 'medium', 'high'.
	ChangedProfileTrustLevel string   `url:"changedProfileTrustLevel,omitempty"` //Trust level of the endpoint due to changing profile labels. Possible values are 'low', 'medium', 'high'.
	NatTrustLevel            string   `url:"natTrustLevel,omitempty"`            //Trust level of the endpoint due to NAT access. Possible values are 'low', 'medium', 'high'.
	ConcurrentMacTrustLevel  string   `url:"concurrentMacTrustLevel,omitempty"`  //Trust level of the endpoint due to concurrent MAC address. Possible values are 'low', 'medium', 'high'.
	IPBlocklistDetected      bool     `url:"ipBlocklistDetected,omitempty"`      //Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       bool     `url:"unauthPortDetected,omitempty"`       //Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         bool     `url:"weakCredDetected,omitempty"`         //Flag to fetch endpoints having weak credentials or not.
	AncPolicy                string   `url:"ancPolicy,omitempty"`                //ANC policy. Only exact match will be returned.
	Limit                    int      `url:"limit,omitempty"`                    //Maximum number of records to be fetched. If not provided, 50 records will be fetched by default. Maximum 1000 records can be fetched at a time. Use pagination if more records need to be fetched.
	Offset                   int      `url:"offset,omitempty"`                   //Record offset to start data fetch at. Offset starts at zero.
	SortBy                   string   `url:"sortBy,omitempty"`                   //Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested. Possible values are 'macAddress', 'ip'.
	Order                    string   `url:"order,omitempty"`                    //Order to be used for sorting. Possible values are 'asc', 'desc'.
	Include                  string   `url:"include,omitempty"`                  //The datasets that should be included in the response. By default, value of this parameter is blank, and the response will include only basic details of the endpoint. To include other datasets or dictionaries, send comma separated list of following values: 'ALL' - Include all attributes. 'CDP', 'DHCP', etc. - Include attributes from given dictionaries. To get full list of dictionaries, use corresponding GET API. 'ANC' - Include ANC policy related details. 'TRUST' - Include trust score details.
}
type FetchTheCountOfEndpointsQueryParams struct {
	ProfilingStatus          string   `url:"profilingStatus,omitempty"`          //Profiling status of the endpoint. Possible values are 'profiled', 'partialProfiled', 'notProfiled'.
	MacAddress               string   `url:"macAddress,omitempty"`               //MAC address to search for. Partial string is allowed.
	MacAddresses             []string `url:"macAddresses,omitempty"`             //List of MAC addresses to filter on. Only exact matches will be returned.
	IP                       string   `url:"ip,omitempty"`                       //IP address to search for. Partial string is allowed.
	DeviceType               string   `url:"deviceType,omitempty"`               //Type of device to search for. Partial string is allowed.
	HardwareManufacturer     string   `url:"hardwareManufacturer,omitempty"`     //Hardware manufacturer to search for. Partial string is allowed.
	HardwareModel            string   `url:"hardwareModel,omitempty"`            //Hardware model to search for. Partial string is allowed.
	OperatingSystem          string   `url:"operatingSystem,omitempty"`          //Operating system to search for. Partial string is allowed.
	Registered               bool     `url:"registered,omitempty"`               //Flag to fetch manually registered or non-registered endpoints.
	RandomMac                bool     `url:"randomMac,omitempty"`                //Flag to fetch endpoints having randomized MAC or not.
	TrustScore               string   `url:"trustScore,omitempty"`               //Overall trust score of the endpoint. It can be provided either as a number value (e.g. 5), or as a range (e.g. 3-7). Provide value as '-' if you want to search for all endpoints where trust score is not assigned.
	AuthMethod               string   `url:"authMethod,omitempty"`               //Authentication method. Partial string is allowed.
	PostureStatus            string   `url:"postureStatus,omitempty"`            //Posture status.
	AiSpoofingTrustLevel     string   `url:"aiSpoofingTrustLevel,omitempty"`     //Trust level of the endpoint due to AI spoofing. Possible values are 'low', 'medium', 'high'.
	ChangedProfileTrustLevel string   `url:"changedProfileTrustLevel,omitempty"` //Trust level of the endpoint due to changing profile labels. Possible values are 'low', 'medium', 'high'.
	NatTrustLevel            string   `url:"natTrustLevel,omitempty"`            //Trust level of the endpoint due to NAT access. Possible values are 'low', 'medium', 'high'.
	ConcurrentMacTrustLevel  string   `url:"concurrentMacTrustLevel,omitempty"`  //Trust level of the endpoint due to concurrent MAC address. Possible values are 'low', 'medium', 'high'.
	IPBlocklistDetected      bool     `url:"ipBlocklistDetected,omitempty"`      //Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       bool     `url:"unauthPortDetected,omitempty"`       //Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         bool     `url:"weakCredDetected,omitempty"`         //Flag to fetch endpoints having weak credentials or not.
	AncPolicy                string   `url:"ancPolicy,omitempty"`                //ANC policy. Only exact match will be returned.
}
type GetEndpointDetailsQueryParams struct {
	Include string `url:"include,omitempty"` //The datasets that should be included in the response. By default, value of this parameter is blank, and the response will include only basic details of the endpoint. To include other datasets or dictionaries, send comma separated list of following values: 'ALL' - Include all attributes. 'CDP', 'DHCP', etc. - Include attributes from given dictionaries. To get full list of dictionaries, use corresponding GET API. 'ANC' - Include ANC policy related details. 'TRUST' - Include trust score details.
}
type GetListOfProfilingRulesQueryParams struct {
	RuleType       string  `url:"ruleType,omitempty"`       //Use comma-separated list of rule types to filter the data. Defaults to 'Custom Rule'.
	IncludeDeleted bool    `url:"includeDeleted,omitempty"` //Flag to indicate whether deleted rules should be part of the records fetched.
	Limit          float64 `url:"limit,omitempty"`          //Maximum number of records to be fetched. If not provided, 500 records will be fetched by default. To fetch all the records in the system, provide a large value for this parameter.
	Offset         float64 `url:"offset,omitempty"`         //Record offset to start data fetch at. Offset starts at zero.
	SortBy         string  `url:"sortBy,omitempty"`         //Name of the column to sort the results on. Please note that fetch might take more time if sorting is requested.
	Order          string  `url:"order,omitempty"`          //Order to be used for sorting.
}
type GetCountOfProfilingRulesQueryParams struct {
	RuleType       string `url:"ruleType,omitempty"`       //Use comma-separated list of rule types to filter the data. Defaults to 'Custom Rule'.
	IncludeDeleted bool   `url:"includeDeleted,omitempty"` //Flag to indicate whether deleted rules should be part of the records fetched.
}

type ResponseAIEndpointAnalyticsGetAncPolicies []ResponseItemAIEndpointAnalyticsGetAncPolicies // Array of ResponseAIEndpointAnalyticsGetANCPolicies
type ResponseItemAIEndpointAnalyticsGetAncPolicies struct {
	Name string `json:"name,omitempty"` // Name of the ANC policy.
}
type ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries []ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries // Array of ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries
type ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries struct {
	Name        string                                                                                  `json:"name,omitempty"`        // Name of the dictionary.
	Description string                                                                                  `json:"description,omitempty"` // Description of the dictionary.
	Attributes  *[]ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesAttributes `json:"attributes,omitempty"`  //
}
type ResponseItemAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionariesAttributes struct {
	Name        string `json:"name,omitempty"`        // Name of the attribute.
	Description string `json:"description,omitempty"` // Description of the attribute.
}
type ResponseAIEndpointAnalyticsQueryTheEndpoints struct {
	TotalResults   *int                                                 `json:"totalResults,omitempty"`   // Total number of records matching the given filter criteria.
	HasMoreResults *bool                                                `json:"hasMoreResults,omitempty"` // Flag to indicate whether more results are available than what is currently in the response.
	Items          *[]ResponseAIEndpointAnalyticsQueryTheEndpointsItems `json:"items,omitempty"`          //
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsItems struct {
	ID                           string                                                                `json:"id,omitempty"`                           // Unique identifier for the endpoint.
	Duid                         string                                                                `json:"duid,omitempty"`                         // Unique DUID.
	MacAddress                   string                                                                `json:"macAddress,omitempty"`                   // MAC address of the endpoint.
	DeviceType                   []string                                                              `json:"deviceType,omitempty"`                   // Type of the device represented by this endpoint.
	HardwareManufacturer         []string                                                              `json:"hardwareManufacturer,omitempty"`         // Hardware manufacturer for the endpoint.
	HardwareModel                []string                                                              `json:"hardwareModel,omitempty"`                // Hardware model of the endpoint.
	OperatingSystem              []string                                                              `json:"operatingSystem,omitempty"`              // Operating system of the endpoint.
	LastProbeCollectionTimestamp *int                                                                  `json:"lastProbeCollectionTimestamp,omitempty"` // Last probe collection timestamp in epoch milliseconds.
	RandomMac                    *bool                                                                 `json:"randomMac,omitempty"`                    // Flag to indicate whether MAC address is a randomized one or not.
	Registered                   *bool                                                                 `json:"registered,omitempty"`                   // Flag to indicate whether this is a manually registered endpoint or not.
	Attributes                   *ResponseAIEndpointAnalyticsQueryTheEndpointsItemsAttributes          `json:"attributes,omitempty"`                   // Various endpoint attributes grouped by dictionaries (e.g. IP, DHCP, etc).
	TrustData                    *ResponseAIEndpointAnalyticsQueryTheEndpointsItemsTrustData           `json:"trustData,omitempty"`                    //
	AncPolicy                    string                                                                `json:"ancPolicy,omitempty"`                    // ANC policy currently applied to the endpoint in ISE.
	GranularAncPolicy            *[]ResponseAIEndpointAnalyticsQueryTheEndpointsItemsGranularAncPolicy `json:"granularAncPolicy,omitempty"`            //
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsItemsAttributes interface{}
type ResponseAIEndpointAnalyticsQueryTheEndpointsItemsTrustData struct {
	TrustScore               *int   `json:"trustScore,omitempty"`               // Overall trust score of the endpoint.
	AuthMethod               string `json:"authMethod,omitempty"`               // Authentication method.
	PostureStatus            string `json:"postureStatus,omitempty"`            // Posture status.
	AiSpoofingTrustLevel     string `json:"aiSpoofingTrustLevel,omitempty"`     // Trust level of the endpoint due to AI spoofing.
	ChangedProfileTrustLevel string `json:"changedProfileTrustLevel,omitempty"` // Trust level of the endpoint due to changing profile labels.
	NatTrustLevel            string `json:"natTrustLevel,omitempty"`            // Trust level of the endpoint due to NAT access.
	ConcurrentMacTrustLevel  string `json:"concurrentMacTrustLevel,omitempty"`  // Trust level of the endpoint due to concurrent MAC address.
	IPBlocklistDetected      *bool  `json:"ipBlocklistDetected,omitempty"`      // Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       *bool  `json:"unauthPortDetected,omitempty"`       // Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         *bool  `json:"weakCredDetected,omitempty"`         // Flag to fetch endpoints having weak credentials or not.
}
type ResponseAIEndpointAnalyticsQueryTheEndpointsItemsGranularAncPolicy struct {
	Name         string `json:"name,omitempty"`         // Name of the granular ANC policy.
	NasIPAddress string `json:"nasIpAddress,omitempty"` // IP address of the network device where endpoint is attached.
}
type ResponseAIEndpointAnalyticsFetchTheCountOfEndpoints struct {
	Count *int `json:"count,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetEndpointDetails struct {
	ID                           string                                                            `json:"id,omitempty"`                           // Unique identifier for the endpoint.
	Duid                         string                                                            `json:"duid,omitempty"`                         // Unique DUID.
	MacAddress                   string                                                            `json:"macAddress,omitempty"`                   // MAC address of the endpoint.
	DeviceType                   []string                                                          `json:"deviceType,omitempty"`                   // Type of the device represented by this endpoint.
	HardwareManufacturer         []string                                                          `json:"hardwareManufacturer,omitempty"`         // Hardware manufacturer for the endpoint.
	HardwareModel                []string                                                          `json:"hardwareModel,omitempty"`                // Hardware model of the endpoint.
	OperatingSystem              []string                                                          `json:"operatingSystem,omitempty"`              // Operating system of the endpoint.
	LastProbeCollectionTimestamp *int                                                              `json:"lastProbeCollectionTimestamp,omitempty"` // Last probe collection timestamp in epoch milliseconds.
	RandomMac                    *bool                                                             `json:"randomMac,omitempty"`                    // Flag to indicate whether MAC address is a randomized one or not.
	Registered                   *bool                                                             `json:"registered,omitempty"`                   // Flag to indicate whether this is a manually registered endpoint or not.
	Attributes                   *ResponseAIEndpointAnalyticsGetEndpointDetailsAttributes          `json:"attributes,omitempty"`                   // Various endpoint attributes grouped by dictionaries (e.g. IP, DHCP, etc).
	TrustData                    *ResponseAIEndpointAnalyticsGetEndpointDetailsTrustData           `json:"trustData,omitempty"`                    //
	AncPolicy                    string                                                            `json:"ancPolicy,omitempty"`                    // ANC policy currently applied to the endpoint in ISE.
	GranularAncPolicy            *[]ResponseAIEndpointAnalyticsGetEndpointDetailsGranularAncPolicy `json:"granularAncPolicy,omitempty"`            //
}
type ResponseAIEndpointAnalyticsGetEndpointDetailsAttributes interface{}
type ResponseAIEndpointAnalyticsGetEndpointDetailsTrustData struct {
	TrustScore               *int   `json:"trustScore,omitempty"`               // Overall trust score of the endpoint.
	AuthMethod               string `json:"authMethod,omitempty"`               // Authentication method.
	PostureStatus            string `json:"postureStatus,omitempty"`            // Posture status.
	AiSpoofingTrustLevel     string `json:"aiSpoofingTrustLevel,omitempty"`     // Trust level of the endpoint due to AI spoofing.
	ChangedProfileTrustLevel string `json:"changedProfileTrustLevel,omitempty"` // Trust level of the endpoint due to changing profile labels.
	NatTrustLevel            string `json:"natTrustLevel,omitempty"`            // Trust level of the endpoint due to NAT access.
	ConcurrentMacTrustLevel  string `json:"concurrentMacTrustLevel,omitempty"`  // Trust level of the endpoint due to concurrent MAC address.
	IPBlocklistDetected      *bool  `json:"ipBlocklistDetected,omitempty"`      // Flag to fetch endpoints hitting IP blocklist or not.
	UnauthPortDetected       *bool  `json:"unauthPortDetected,omitempty"`       // Flag to fetch endpoints exposing unauthorized ports or not.
	WeakCredDetected         *bool  `json:"weakCredDetected,omitempty"`         // Flag to fetch endpoints having weak credentials or not.
}
type ResponseAIEndpointAnalyticsGetEndpointDetailsGranularAncPolicy struct {
	Name         string `json:"name,omitempty"`         // Name of the granular ANC policy.
	NasIPAddress string `json:"nasIpAddress,omitempty"` // IP address of the network device where endpoint is attached.
}
type ResponseAIEndpointAnalyticsCreateAProfilingRule struct {
	ID   string `json:"id,omitempty"`   // Unique identifier for the newly created resource.
	Link string `json:"link,omitempty"` // Link to the newly created resource.
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRules struct {
	ProfilingRules *[]ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRules `json:"profilingRules,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRules struct {
	RuleID          string                                                                           `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                           `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                           `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                             `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                             `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                             `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                            `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                           `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                             `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                           `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                           `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                            `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRulesResult          `json:"result,omitempty"`          //
	ConditionGroups *ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRulesConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                         `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRulesResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRulesConditionGroups struct {
	Type           string                                                                                    `json:"type,omitempty"`           //
	Condition      *ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRulesConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                                    `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                                  `json:"conditionGroup,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetListOfProfilingRulesProfilingRulesConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetCountOfProfilingRules struct {
	Count *int `json:"count,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRule struct {
	RuleID          string                                                                      `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                      `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                      `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                        `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                        `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                        `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                       `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                      `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                        `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                      `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                      `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                       `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleResult          `json:"result,omitempty"`          //
	ConditionGroups *ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                    `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleConditionGroups struct {
	Type           string                                                                               `json:"type,omitempty"`           //
	Condition      *ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                               `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                             `json:"conditionGroup,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRuleConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type ResponseAIEndpointAnalyticsGetTaskDetails struct {
	ID             string                                                   `json:"id,omitempty"`             // Unique identifier for the task.
	Name           string                                                   `json:"name,omitempty"`           // Name of the task.
	Status         string                                                   `json:"status,omitempty"`         // Status of the task.
	Errors         *[]ResponseAIEndpointAnalyticsGetTaskDetailsErrors       `json:"errors,omitempty"`         //
	AdditionalInfo *ResponseAIEndpointAnalyticsGetTaskDetailsAdditionalInfo `json:"additionalInfo,omitempty"` // Additional information about the task.
	CreatedBy      string                                                   `json:"createdBy,omitempty"`      // Name of the user that created the task.
	CreatedOn      *int                                                     `json:"createdOn,omitempty"`      // Task creation timestamp in epoch milliseconds.
	LastUpdatedOn  *int                                                     `json:"lastUpdatedOn,omitempty"`  // Last update timestamp in epoch milliseconds.
}
type ResponseAIEndpointAnalyticsGetTaskDetailsErrors struct {
	Index   *int   `json:"index,omitempty"`   // Index of the input records which had error during processing. In case the input is not an array, or the error is not record specific, this will be -1.
	Code    *int   `json:"code,omitempty"`    // Error code.
	Message string `json:"message,omitempty"` // Error message.
	Details string `json:"details,omitempty"` // Optional details about the error.
}
type ResponseAIEndpointAnalyticsGetTaskDetailsAdditionalInfo interface{}
type RequestAIEndpointAnalyticsProcessCmdbEndpoints []RequestItemAIEndpointAnalyticsProcessCmdbEndpoints // Array of RequestAIEndpointAnalyticsProcessCMDBEndpoints
type RequestItemAIEndpointAnalyticsProcessCmdbEndpoints struct {
	MacAddress          string `json:"macAddress,omitempty"`          // MAC address of the endpoint.
	SerialNumber        string `json:"serialNumber,omitempty"`        // Serial number of the endpoint.
	AssetTag            string `json:"assetTag,omitempty"`            // Asset tag.
	ModelCategory       string `json:"modelCategory,omitempty"`       // Category of the model.
	Model               string `json:"model,omitempty"`               // Asset model.
	DisplayName         string `json:"displayName,omitempty"`         // Display name of the asset.
	Department          string `json:"department,omitempty"`          // Department that asset belongs to.
	Location            string `json:"location,omitempty"`            // Location of the asset.
	ManagedBy           string `json:"managedBy,omitempty"`           // Asset managed by.
	LastUpdateTimestamp *int   `json:"lastUpdateTimestamp,omitempty"` // Last update timestamp in epoch milliseconds.
}
type RequestAIEndpointAnalyticsRegisterAnEndpoint struct {
	MacAddress           string `json:"macAddress,omitempty"`           // MAC address of the endpoint.
	DeviceType           string `json:"deviceType,omitempty"`           // Type of the device represented by this endpoint.
	HardwareManufacturer string `json:"hardwareManufacturer,omitempty"` // Hardware manufacturer for the endpoint.
	HardwareModel        string `json:"hardwareModel,omitempty"`        // Hardware model of the endpoint.
}
type RequestAIEndpointAnalyticsUpdateARegisteredEndpoint struct {
	DeviceType           string `json:"deviceType,omitempty"`           // Type of the device represented by this endpoint.
	HardwareManufacturer string `json:"hardwareManufacturer,omitempty"` // Hardware manufacturer for the endpoint.
	HardwareModel        string `json:"hardwareModel,omitempty"`        // Hardware model of the endpoint.
}
type RequestAIEndpointAnalyticsApplyAncPolicy struct {
	AncPolicy         string                                                       `json:"ancPolicy,omitempty"`         // ANC policy name.
	GranularAncPolicy *[]RequestAIEndpointAnalyticsApplyAncPolicyGranularAncPolicy `json:"granularAncPolicy,omitempty"` //
}
type RequestAIEndpointAnalyticsApplyAncPolicyGranularAncPolicy struct {
	Name         string `json:"name,omitempty"`         // Name of the granular ANC policy.
	NasIPAddress string `json:"nasIpAddress,omitempty"` // IP address of the network device where endpoint is attached.
}
type RequestAIEndpointAnalyticsCreateAProfilingRule struct {
	RuleID          string                                                         `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                         `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                         `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                           `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                           `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                           `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                          `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                         `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                           `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                         `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                         `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                          `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *RequestAIEndpointAnalyticsCreateAProfilingRuleResult          `json:"result,omitempty"`          //
	ConditionGroups *RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                       `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroups struct {
	Type           string                                                                  `json:"type,omitempty"`           //
	Condition      *RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                  `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                `json:"conditionGroup,omitempty"` //
}
type RequestAIEndpointAnalyticsCreateAProfilingRuleConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulk struct {
	ProfilingRules *[]RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRules `json:"profilingRules,omitempty"` //
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRules struct {
	RuleID          string                                                                             `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                             `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                             `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                               `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                               `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                               `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                              `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                             `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                               `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                             `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                             `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                              `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRulesResult          `json:"result,omitempty"`          //
	ConditionGroups *RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRulesConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                           `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRulesResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRulesConditionGroups struct {
	Type           string                                                                                      `json:"type,omitempty"`           //
	Condition      *RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRulesConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                                      `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                                    `json:"conditionGroup,omitempty"` //
}
type RequestAIEndpointAnalyticsImportProfilingRulesInBulkProfilingRulesConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRule struct {
	RuleID          string                                                                  `json:"ruleId,omitempty"`          // Unique identifier for the rule. This is normally generated by the system, and client does not need to provide it for rules that need to be newly created.
	RuleName        string                                                                  `json:"ruleName,omitempty"`        // Human readable name for the rule.
	RuleType        string                                                                  `json:"ruleType,omitempty"`        // Type of the rule.
	RuleVersion     *int                                                                    `json:"ruleVersion,omitempty"`     // Version of the rule.
	RulePriority    *int                                                                    `json:"rulePriority,omitempty"`    // Priority for the rule.
	SourcePriority  *int                                                                    `json:"sourcePriority,omitempty"`  // Source priority for the rule.
	IsDeleted       *bool                                                                   `json:"isDeleted,omitempty"`       // Flag to indicate whether the rule was deleted.
	LastModifiedBy  string                                                                  `json:"lastModifiedBy,omitempty"`  // User that last modified the rule. It is read-only, and is ignored if provided as part of input request.
	LastModifiedOn  *int                                                                    `json:"lastModifiedOn,omitempty"`  // Timestamp (in epoch milliseconds) of last modification. It is read-only, and is ignored if provided as part of input request.
	PluginID        string                                                                  `json:"pluginId,omitempty"`        // Plugin for the rule. Only applicable for 'Cisco Default' rules.
	ClusterID       string                                                                  `json:"clusterId,omitempty"`       // Unique identifier for ML cluster. Only applicable for 'ML Rule'.
	Rejected        *bool                                                                   `json:"rejected,omitempty"`        // Flag to indicate whether rule has been rejected by user or not. Only applicable for 'ML Rule'.
	Result          *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleResult          `json:"result,omitempty"`          //
	ConditionGroups *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroups `json:"conditionGroups,omitempty"` //
	UsedAttributes  []string                                                                `json:"usedAttributes,omitempty"`  // List of attributes used in the rule. Only applicable for 'Cisco Default' rules.
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleResult struct {
	DeviceType           []string `json:"deviceType,omitempty"`           // List of device types determined by the current rule.
	HardwareManufacturer []string `json:"hardwareManufacturer,omitempty"` // List of hardware manufacturers determined by the current rule.
	HardwareModel        []string `json:"hardwareModel,omitempty"`        // List of hardware models determined by the current rule.
	OperatingSystem      []string `json:"operatingSystem,omitempty"`      // List of operating systems determined by the current rule.
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroups struct {
	Type           string                                                                           `json:"type,omitempty"`           //
	Condition      *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroupsCondition `json:"condition,omitempty"`      //
	Operator       string                                                                           `json:"operator,omitempty"`       //
	ConditionGroup []string                                                                         `json:"conditionGroup,omitempty"` //
}
type RequestAIEndpointAnalyticsUpdateAnExistingProfilingRuleConditionGroupsCondition struct {
	Attribute           string `json:"attribute,omitempty"`           //
	Operator            string `json:"operator,omitempty"`            //
	Value               string `json:"value,omitempty"`               //
	AttributeDictionary string `json:"attributeDictionary,omitempty"` //
}

//GetAncPolicies Get ANC policies - ae9a-f945-47a8-871e
/* Fetches the list of ANC policies available in ISE.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-anc-policies-v1
*/
func (s *AIEndpointAnalyticsService) GetAncPolicies() (*ResponseAIEndpointAnalyticsGetAncPolicies, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/anc-policies"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAIEndpointAnalyticsGetAncPolicies{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAncPolicies()
		}
		return nil, response, fmt.Errorf("error with operation GetAncPolicies")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetAncPolicies)
	return result, response, err

}

//GetAIEndpointAnalyticsAttributeDictionaries Get AI Endpoint Analytics attribute dictionaries - 409f-1aff-482a-ae1e
/* Fetches the list of attribute dictionaries.


@param GetAIEndpointAnalyticsAttributeDictionariesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-a-i-endpoint-analytics-attribute-dictionaries-v1
*/
func (s *AIEndpointAnalyticsService) GetAIEndpointAnalyticsAttributeDictionaries(GetAIEndpointAnalyticsAttributeDictionariesQueryParams *GetAIEndpointAnalyticsAttributeDictionariesQueryParams) (*ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/dictionaries"

	queryString, _ := query.Values(GetAIEndpointAnalyticsAttributeDictionariesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAIEndpointAnalyticsAttributeDictionaries(GetAIEndpointAnalyticsAttributeDictionariesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAIEndpointAnalyticsAttributeDictionaries")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetAIEndpointAnalyticsAttributeDictionaries)
	return result, response, err

}

//QueryTheEndpoints Query the endpoints - aeb4-7a77-425b-b30f
/* Query the endpoints, optionally using various filter and pagination criteria. 'GET /endpoints/count' API can be used to find out the total number of endpoints matching the filter criteria.


@param QueryTheEndpointsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!query-the-endpoints-v1
*/
func (s *AIEndpointAnalyticsService) QueryTheEndpoints(QueryTheEndpointsQueryParams *QueryTheEndpointsQueryParams) (*ResponseAIEndpointAnalyticsQueryTheEndpoints, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints"

	queryString, _ := query.Values(QueryTheEndpointsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsQueryTheEndpoints{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.QueryTheEndpoints(QueryTheEndpointsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation QueryTheEndpoints")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsQueryTheEndpoints)
	return result, response, err

}

//FetchTheCountOfEndpoints Fetch the count of endpoints - 04b2-bbb0-472b-9ce0
/* Fetch the total count of endpoints that match the given filter criteria.


@param FetchTheCountOfEndpointsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!fetch-the-count-of-endpoints-v1
*/
func (s *AIEndpointAnalyticsService) FetchTheCountOfEndpoints(FetchTheCountOfEndpointsQueryParams *FetchTheCountOfEndpointsQueryParams) (*ResponseAIEndpointAnalyticsFetchTheCountOfEndpoints, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/count"

	queryString, _ := query.Values(FetchTheCountOfEndpointsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsFetchTheCountOfEndpoints{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.FetchTheCountOfEndpoints(FetchTheCountOfEndpointsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation FetchTheCountOfEndpoints")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsFetchTheCountOfEndpoints)
	return result, response, err

}

//GetEndpointDetails Get endpoint details - 5881-9a5e-41a8-8cce
/* Fetches details of the endpoint for the given unique identifier 'epId'.


@param epID epId path parameter. Unique identifier for the endpoint.

@param GetEndpointDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-endpoint-details-v1
*/
func (s *AIEndpointAnalyticsService) GetEndpointDetails(epID string, GetEndpointDetailsQueryParams *GetEndpointDetailsQueryParams) (*ResponseAIEndpointAnalyticsGetEndpointDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	queryString, _ := query.Values(GetEndpointDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetEndpointDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEndpointDetails(epID, GetEndpointDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetEndpointDetails")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetEndpointDetails)
	return result, response, err

}

//GetListOfProfilingRules Get list of profiling rules - 07b4-eb60-435a-bf90
/* This API fetches the list of profiling rules. It can be used to show profiling rules in client applications, or export those from an environment. 'POST /profiling-rules/bulk' API can be used to import such exported rules into another environment. If this API is used to export rules to be imported into another Cisco DNA Center system, then ensure that 'includeDeleted' parameter is 'true', so that deleted rules get synchronized correctly. Use query parameters to filter the data, as required. If no filter is provided, then it will include only rules of type 'Custom Rule' in the response. By default, the response is limited to 500 records. Use 'limit' parameter to fetch higher number of records, if required. 'GET /profiling-rules/count' API can be used to find out the total number of rules in the system.


@param GetListOfProfilingRulesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-profiling-rules-v1
*/
func (s *AIEndpointAnalyticsService) GetListOfProfilingRules(GetListOfProfilingRulesQueryParams *GetListOfProfilingRulesQueryParams) (*ResponseAIEndpointAnalyticsGetListOfProfilingRules, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules"

	queryString, _ := query.Values(GetListOfProfilingRulesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetListOfProfilingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfProfilingRules(GetListOfProfilingRulesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfProfilingRules")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetListOfProfilingRules)
	return result, response, err

}

//GetCountOfProfilingRules Get count of profiling rules - 4dad-ba2c-4968-b494
/* This API fetches the count of profiling rules based on the filter values provided in the query parameters. The filter parameters are same as that of 'GET /profiling-rules' API, excluding the pagination and sort parameters.


@param GetCountOfProfilingRulesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-count-of-profiling-rules-v1
*/
func (s *AIEndpointAnalyticsService) GetCountOfProfilingRules(GetCountOfProfilingRulesQueryParams *GetCountOfProfilingRulesQueryParams) (*ResponseAIEndpointAnalyticsGetCountOfProfilingRules, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/count"

	queryString, _ := query.Values(GetCountOfProfilingRulesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseAIEndpointAnalyticsGetCountOfProfilingRules{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetCountOfProfilingRules(GetCountOfProfilingRulesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetCountOfProfilingRules")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetCountOfProfilingRules)
	return result, response, err

}

//GetDetailsOfASingleProfilingRule Get details of a single profiling rule - 20bc-6a22-4a4b-bead
/* Fetches details of the profiling rule for the given 'ruleId'.


@param ruleID ruleId path parameter. Unique rule identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-details-of-a-single-profiling-rule-v1
*/
func (s *AIEndpointAnalyticsService) GetDetailsOfASingleProfilingRule(ruleID string) (*ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRule, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/{ruleId}"
	path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRule{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDetailsOfASingleProfilingRule(ruleID)
		}
		return nil, response, fmt.Errorf("error with operation GetDetailsOfASingleProfilingRule")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetDetailsOfASingleProfilingRule)
	return result, response, err

}

//GetTaskDetails Get task details - 2689-39c4-43fa-a2f2
/* Fetches the details of backend task. Task is typically created by making call to some other API that takes longer time to execute.


@param taskID taskId path parameter. Unique identifier for the task.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-task-details-v1
*/
func (s *AIEndpointAnalyticsService) GetTaskDetails(taskID string) (*ResponseAIEndpointAnalyticsGetTaskDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/tasks/{taskId}"
	path = strings.Replace(path, "{taskId}", fmt.Sprintf("%v", taskID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAIEndpointAnalyticsGetTaskDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTaskDetails(taskID)
		}
		return nil, response, fmt.Errorf("error with operation GetTaskDetails")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsGetTaskDetails)
	return result, response, err

}

//ProcessCmdbEndpoints Process CMDB endpoints - fa9f-f839-42fb-9e38
/* Processes incoming CMDB endpoints data and imports the same in AI Endpoint Analytics.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!process-cmdb-endpoints-v1
*/
func (s *AIEndpointAnalyticsService) ProcessCmdbEndpoints(requestAIEndpointAnalyticsProcessCMDBEndpoints *RequestAIEndpointAnalyticsProcessCmdbEndpoints) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/cmdb/endpoints"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsProcessCMDBEndpoints).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ProcessCmdbEndpoints(requestAIEndpointAnalyticsProcessCMDBEndpoints)
		}

		return response, fmt.Errorf("error with operation ProcessCmdbEndpoints")
	}

	return response, err

}

//RegisterAnEndpoint Register an endpoint - a895-f856-4089-92fd
/* Register a new endpoint in the system.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!register-an-endpoint-v1
*/
func (s *AIEndpointAnalyticsService) RegisterAnEndpoint(requestAIEndpointAnalyticsRegisterAnEndpoint *RequestAIEndpointAnalyticsRegisterAnEndpoint) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsRegisterAnEndpoint).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RegisterAnEndpoint(requestAIEndpointAnalyticsRegisterAnEndpoint)
		}

		return response, fmt.Errorf("error with operation RegisterAnEndpoint")
	}

	return response, err

}

//CreateAProfilingRule Create a profiling rule - 6cb9-98bb-47ea-90f6
/* Creates profiling rule from the request body.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a-profiling-rule-v1
*/
func (s *AIEndpointAnalyticsService) CreateAProfilingRule(requestAIEndpointAnalyticsCreateAProfilingRule *RequestAIEndpointAnalyticsCreateAProfilingRule) (*ResponseAIEndpointAnalyticsCreateAProfilingRule, *resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsCreateAProfilingRule).
		SetResult(&ResponseAIEndpointAnalyticsCreateAProfilingRule{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAProfilingRule(requestAIEndpointAnalyticsCreateAProfilingRule)
		}

		return nil, response, fmt.Errorf("error with operation CreateAProfilingRule")
	}

	result := response.Result().(*ResponseAIEndpointAnalyticsCreateAProfilingRule)
	return result, response, err

}

//ImportProfilingRulesInBulk Import profiling rules in bulk - 70bf-885f-408a-9c74
/* This API imports the given list of profiling rules. For each record, 1) If 'ruleType' for a record is not 'Custom Rule', then it is rejected. 2) If 'ruleId' is provided in the input record,
  2a) Record with same 'ruleId' exists in the system, then it is replaced with new data.
  2b) Record with same 'ruleId' does not exist, then it is inserted in the database.
3) If 'ruleId' is not provided in the input record, then new 'ruleId' is generated by the system, and record is inserted.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-profiling-rules-in-bulk-v1
*/
func (s *AIEndpointAnalyticsService) ImportProfilingRulesInBulk(requestAIEndpointAnalyticsImportProfilingRulesInBulk *RequestAIEndpointAnalyticsImportProfilingRulesInBulk) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/bulk"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsImportProfilingRulesInBulk).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportProfilingRulesInBulk(requestAIEndpointAnalyticsImportProfilingRulesInBulk)
		}

		return response, fmt.Errorf("error with operation ImportProfilingRulesInBulk")
	}

	return response, err

}

//UpdateARegisteredEndpoint Update a registered endpoint - e5af-892c-40e9-a2a1
/* Update attributes of a registered endpoint.


@param epID epId path parameter. Unique identifier for the endpoint.

*/
func (s *AIEndpointAnalyticsService) UpdateARegisteredEndpoint(epID string, requestAIEndpointAnalyticsUpdateARegisteredEndpoint *RequestAIEndpointAnalyticsUpdateARegisteredEndpoint) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsUpdateARegisteredEndpoint).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateARegisteredEndpoint(epID, requestAIEndpointAnalyticsUpdateARegisteredEndpoint)
		}
		return response, fmt.Errorf("error with operation UpdateARegisteredEndpoint")
	}

	return response, err

}

//ApplyAncPolicy Apply ANC Policy - 2ebb-79f2-4489-8b73
/* Applies given ANC policy to the endpoint.


@param epID epId path parameter. Unique identifier for the endpoint.

*/
func (s *AIEndpointAnalyticsService) ApplyAncPolicy(epID string, requestAIEndpointAnalyticsApplyANCPolicy *RequestAIEndpointAnalyticsApplyAncPolicy) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}/anc-policy"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsApplyANCPolicy).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ApplyAncPolicy(epID, requestAIEndpointAnalyticsApplyANCPolicy)
		}
		return response, fmt.Errorf("error with operation ApplyAncPolicy")
	}

	return response, err

}

//UpdateAnExistingProfilingRule Update an existing profiling rule - c197-6aa2-4fd9-82d7
/* Updates the profiling rule for the given 'ruleId'.


@param ruleID ruleId path parameter. Unique rule identifier

*/
func (s *AIEndpointAnalyticsService) UpdateAnExistingProfilingRule(ruleID string, requestAIEndpointAnalyticsUpdateAnExistingProfilingRule *RequestAIEndpointAnalyticsUpdateAnExistingProfilingRule) (*resty.Response, error) {
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/{ruleId}"
	path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestAIEndpointAnalyticsUpdateAnExistingProfilingRule).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateAnExistingProfilingRule(ruleID, requestAIEndpointAnalyticsUpdateAnExistingProfilingRule)
		}
		return response, fmt.Errorf("error with operation UpdateAnExistingProfilingRule")
	}

	return response, err

}

//DeleteAnEndpoint Delete an endpoint - 689d-e83b-442a-9435
/* Deletes the endpoint for the given unique identifier 'epId'.


@param epID epId path parameter. Unique identifier for the endpoint.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-an-endpoint-v1
*/
func (s *AIEndpointAnalyticsService) DeleteAnEndpoint(epID string) (*resty.Response, error) {
	//epID string
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAnEndpoint(
				epID)
		}
		return response, fmt.Errorf("error with operation DeleteAnEndpoint")
	}

	return response, err

}

//RevokeAncPolicy Revoke ANC policy - 8982-89f3-4e1b-b3dc
/* Revokes given ANC policy from the endpoint.


@param epID epId path parameter. Unique identifier for the endpoint.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!revoke-anc-policy-v1
*/
func (s *AIEndpointAnalyticsService) RevokeAncPolicy(epID string) (*resty.Response, error) {
	//epID string
	path := "/dna/intent/api/v1/endpoint-analytics/endpoints/{epId}/anc-policy"
	path = strings.Replace(path, "{epId}", fmt.Sprintf("%v", epID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RevokeAncPolicy(
				epID)
		}
		return response, fmt.Errorf("error with operation RevokeAncPolicy")
	}

	return response, err

}

//DeleteAnExistingProfilingRule Delete an existing profiling rule - 6f9f-98d4-4b0b-9e7c
/* Deletes the profiling rule for the given 'ruleId'.


@param ruleID ruleId path parameter. Unique rule identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-an-existing-profiling-rule-v1
*/
func (s *AIEndpointAnalyticsService) DeleteAnExistingProfilingRule(ruleID string) (*resty.Response, error) {
	//ruleID string
	path := "/dna/intent/api/v1/endpoint-analytics/profiling-rules/{ruleId}"
	path = strings.Replace(path, "{ruleId}", fmt.Sprintf("%v", ruleID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteAnExistingProfilingRule(
				ruleID)
		}
		return response, fmt.Errorf("error with operation DeleteAnExistingProfilingRule")
	}

	return response, err

}
