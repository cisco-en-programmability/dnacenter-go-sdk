package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type IssuesService service

type GetIssueEnrichmentDetailsHeaderParams struct {
	EntityType  string `url:"entity_type,omitempty"`  //Expects type string. Issue enrichment details can be fetched based on either Issue ID or Client MAC address. This parameter value must either be issue_id/mac_address
	EntityValue string `url:"entity_value,omitempty"` //Expects type string. Contains the actual value for the entity type that has been defined
}
type IssuesQueryParams struct {
	StartTime   float64 `url:"startTime,omitempty"`   //Starting epoch time in milliseconds of query time window
	EndTime     float64 `url:"endTime,omitempty"`     //Ending epoch time in milliseconds of query time window
	SiteID      string  `url:"siteId,omitempty"`      //Assurance UUID value of the site in the issue content
	DeviceID    string  `url:"deviceId,omitempty"`    //Assurance UUID value of the device in the issue content
	MacAddress  string  `url:"macAddress,omitempty"`  //Client's device MAC address of the issue (format xx:xx:xx:xx:xx:xx)
	Priority    string  `url:"priority,omitempty"`    //The issue's priority value (One of P1, P2, P3, or P4)(Use only when macAddress and deviceId are not provided)
	AiDriven    string  `url:"aiDriven,omitempty"`    //The issue's AI driven value (Yes or No)(Use only when macAddress and deviceId are not provided)
	IssueStatus string  `url:"issueStatus,omitempty"` //The issue's status value (One of ACTIVE, IGNORED, RESOLVED)
}

type ResponseIssuesExecuteSuggestedActionsCommands []ResponseItemIssuesExecuteSuggestedActionsCommands // Array of ResponseIssuesExecuteSuggestedActionsCommands
type ResponseItemIssuesExecuteSuggestedActionsCommands struct {
	ActionInfo       string                                                          `json:"actionInfo,omitempty"`       // Actions Info
	StepsCount       *int                                                            `json:"stepsCount,omitempty"`       // Steps Count
	EntityID         string                                                          `json:"entityId,omitempty"`         // Entity Id
	Hostname         string                                                          `json:"hostname,omitempty"`         // Hostname
	StepsDescription string                                                          `json:"stepsDescription,omitempty"` // Steps Description
	Command          string                                                          `json:"command,omitempty"`          // Command
	CommandOutput    *ResponseItemIssuesExecuteSuggestedActionsCommandsCommandOutput `json:"commandOutput,omitempty"`    // Command Output
}
type ResponseItemIssuesExecuteSuggestedActionsCommandsCommandOutput interface{}
type ResponseIssuesGetIssueEnrichmentDetails struct {
	IssueDetails *ResponseIssuesGetIssueEnrichmentDetailsIssueDetails `json:"issueDetails,omitempty"` //
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetails struct {
	Issue *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssue `json:"issue,omitempty"` //
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssue struct {
	IssueID          string                                                                      `json:"issueId,omitempty"`          // Issue Id
	IssueSource      string                                                                      `json:"issueSource,omitempty"`      // Issue Source
	IssueCategory    string                                                                      `json:"issueCategory,omitempty"`    // Issue Category
	IssueName        string                                                                      `json:"issueName,omitempty"`        // Issue Name
	IssueDescription string                                                                      `json:"issueDescription,omitempty"` // Issue Description
	IssueEntity      string                                                                      `json:"issueEntity,omitempty"`      // Issue Entity
	IssueEntityValue string                                                                      `json:"issueEntityValue,omitempty"` // Issue Entity Value
	IssueSeverity    string                                                                      `json:"issueSeverity,omitempty"`    // Issue Severity
	IssuePriority    string                                                                      `json:"issuePriority,omitempty"`    // Issue Priority
	IssueSummary     string                                                                      `json:"issueSummary,omitempty"`     // Issue Summary
	IssueTimestamp   *int                                                                        `json:"issueTimestamp,omitempty"`   // Issue Timestamp
	SuggestedActions *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
	ImpactedHosts    *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueImpactedHosts    `json:"impactedHosts,omitempty"`    // Impacted Hosts
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActions struct {
	Message string                                                                           `json:"message,omitempty"` // Message
	Steps   *[]ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps `json:"steps,omitempty"`   // Steps
}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueSuggestedActionsSteps interface{}
type ResponseIssuesGetIssueEnrichmentDetailsIssueDetailsIssueImpactedHosts interface{}
type ResponseIssuesIssues struct {
	Version    string                          `json:"version,omitempty"`    // Version
	TotalCount *int                            `json:"totalCount,omitempty"` // Total Count
	Response   *[]ResponseIssuesIssuesResponse `json:"response,omitempty"`   //
}
type ResponseIssuesIssuesResponse struct {
	IssueID             string `json:"issueId,omitempty"`               // Issue Id
	Name                string `json:"name,omitempty"`                  // Name
	SiteID              string `json:"siteId,omitempty"`                // Site Id
	DeviceID            string `json:"deviceId,omitempty"`              // Device Id
	DeviceRole          string `json:"deviceRole,omitempty"`            // Device Role
	AiDriven            *bool  `json:"aiDriven,omitempty"`              // Ai Driven
	ClientMac           string `json:"clientMac,omitempty"`             // Client Mac
	IssueOccurenceCount *int   `json:"issue_occurence_count,omitempty"` // Issue Occurence Count
	Status              string `json:"status,omitempty"`                // Status
	Priority            string `json:"priority,omitempty"`              // Priority
	Category            string `json:"category,omitempty"`              // Category
	LastOccurenceTime   *int   `json:"last_occurence_time,omitempty"`   // Last Occurence Time
}
type RequestIssuesExecuteSuggestedActionsCommands struct {
	EntityType  string `json:"entity_type,omitempty"`  // Commands provided as part of the suggested actions for an issue can be executed based on issue id. The value here must be issue_id
	EntityValue string `json:"entity_value,omitempty"` // Contains the actual value for the entity type that has been defined
}

//GetIssueEnrichmentDetails Get Issue Enrichment Details - 8684-39bb-4e89-a6e4
/* Enriches a given network issue context (an issue id or end user’s Mac Address) with details about the issue(s), impacted hosts and suggested actions for remediation


@param GetIssueEnrichmentDetailsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-issue-enrichment-details
*/
func (s *IssuesService) GetIssueEnrichmentDetails(GetIssueEnrichmentDetailsHeaderParams *GetIssueEnrichmentDetailsHeaderParams) (*ResponseIssuesGetIssueEnrichmentDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/issue-enrichment-details"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetIssueEnrichmentDetailsHeaderParams != nil {

		if GetIssueEnrichmentDetailsHeaderParams.EntityType != "" {
			clientRequest = clientRequest.SetHeader("entity_type", GetIssueEnrichmentDetailsHeaderParams.EntityType)
		}

		if GetIssueEnrichmentDetailsHeaderParams.EntityValue != "" {
			clientRequest = clientRequest.SetHeader("entity_value", GetIssueEnrichmentDetailsHeaderParams.EntityValue)
		}

	}

	response, err = clientRequest.
		SetResult(&ResponseIssuesGetIssueEnrichmentDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetIssueEnrichmentDetails")
	}

	result := response.Result().(*ResponseIssuesGetIssueEnrichmentDetails)
	return result, response, err

}

//Issues Issues - 5e86-3b7b-4a4b-b2f9
/* Intent API to get a list of global issues, issues for a specific device, or issue for a specific client device's MAC address.


@param IssuesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!issues
*/
func (s *IssuesService) Issues(IssuesQueryParams *IssuesQueryParams) (*ResponseIssuesIssues, *resty.Response, error) {
	path := "/dna/intent/api/v1/issues"

	queryString, _ := query.Values(IssuesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseIssuesIssues{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation Issues")
	}

	result := response.Result().(*ResponseIssuesIssues)
	return result, response, err

}

//ExecuteSuggestedActionsCommands Execute Suggested Actions Commands - cfb2-ab10-4cea-bfbb
/* This API triggers the execution of the suggested actions for an issue, given the Issue Id. It will return an execution Id. At the completion of the execution, the output of the commands associated with the suggested actions will be provided
Invoking this API would provide the execution id. Execute the 'Get Business API Execution Details' API with this execution id, to receive the suggested actions commands output.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!execute-suggested-actions-commands
*/
func (s *IssuesService) ExecuteSuggestedActionsCommands(requestIssuesExecuteSuggestedActionsCommands *RequestIssuesExecuteSuggestedActionsCommands) (*ResponseIssuesExecuteSuggestedActionsCommands, *resty.Response, error) {
	path := "/dna/intent/api/v1/execute-suggested-actions-commands"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestIssuesExecuteSuggestedActionsCommands).
		SetResult(&ResponseIssuesExecuteSuggestedActionsCommands{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ExecuteSuggestedActionsCommands")
	}

	result := response.Result().(*ResponseIssuesExecuteSuggestedActionsCommands)
	return result, response, err

}
