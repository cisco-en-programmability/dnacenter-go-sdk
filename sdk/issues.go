package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// IssuesService is the service to communicate with the Issues API endpoint
type IssuesService service

// GetIssueEnrichmentDetailsResponse is the getIssueEnrichmentDetailsResponse definition
type GetIssueEnrichmentDetailsResponse struct {
	IssueDetails GetIssueEnrichmentDetailsResponseIssueDetails `json:"issueDetails,omitempty"` //
}

// GetIssueEnrichmentDetailsResponseIssueDetails is the getIssueEnrichmentDetailsResponseIssueDetails definition
type GetIssueEnrichmentDetailsResponseIssueDetails struct {
	Issue []GetIssueEnrichmentDetailsResponseIssueDetailsIssue `json:"issue,omitempty"` //
}

// GetIssueEnrichmentDetailsResponseIssueDetailsIssue is the getIssueEnrichmentDetailsResponseIssueDetailsIssue definition
type GetIssueEnrichmentDetailsResponseIssueDetailsIssue struct {
	ImpactedHosts    []string                                                             `json:"impactedHosts,omitempty"`    //
	IssueCategory    string                                                               `json:"issueCategory,omitempty"`    //
	IssueDescription string                                                               `json:"issueDescription,omitempty"` //
	IssueEntity      string                                                               `json:"issueEntity,omitempty"`      //
	IssueEntityValue string                                                               `json:"issueEntityValue,omitempty"` //
	IssueID          string                                                               `json:"issueId,omitempty"`          //
	IssueName        string                                                               `json:"issueName,omitempty"`        //
	IssuePriority    string                                                               `json:"issuePriority,omitempty"`    //
	IssueSeverity    string                                                               `json:"issueSeverity,omitempty"`    //
	IssueSource      string                                                               `json:"issueSource,omitempty"`      //
	IssueSummary     string                                                               `json:"issueSummary,omitempty"`     //
	IssueTimestamp   int                                                                  `json:"issueTimestamp,omitempty"`   //
	SuggestedActions []GetIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions `json:"suggestedActions,omitempty"` //
}

// GetIssueEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts is the getIssueEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts definition
type GetIssueEnrichmentDetailsResponseIssueDetailsIssueImpactedHosts []string

// GetIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions is the getIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions definition
type GetIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActions struct {
	Message string   `json:"message,omitempty"` //
	Steps   []string `json:"steps,omitempty"`   //
}

// GetIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActionsSteps is the getIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActionsSteps definition
type GetIssueEnrichmentDetailsResponseIssueDetailsIssueSuggestedActionsSteps []string

// IssuesResponse is the issuesResponse definition
type IssuesResponse struct {
	Response   []IssuesResponseResponse `json:"response,omitempty"`   //
	TotalCount int                      `json:"totalCount,omitempty"` //
	Version    string                   `json:"version,omitempty"`    //
}

// IssuesResponseResponse is the issuesResponseResponse definition
type IssuesResponseResponse struct {
	AiDriven            bool   `json:"aiDriven,omitempty"`              //
	Category            string `json:"category,omitempty"`              //
	ClientMac           string `json:"clientMac,omitempty"`             //
	DeviceID            string `json:"deviceId,omitempty"`              //
	DeviceRole          string `json:"deviceRole,omitempty"`            //
	IssueID             string `json:"issueId,omitempty"`               //
	IssueOccurenceCount int    `json:"issue_occurence_count,omitempty"` //
	LastOccurenceTime   int    `json:"last_occurence_time,omitempty"`   //
	Name                string `json:"name,omitempty"`                  //
	Priority            string `json:"priority,omitempty"`              //
	SiteID              string `json:"siteId,omitempty"`                //
	Status              string `json:"status,omitempty"`                //
}

// GetIssueEnrichmentDetails getIssueEnrichmentDetails
/* Enriches a given network issue context (an issue id or end user’s Mac Address) with details about the issue(s), impacted hosts and suggested actions for remediation
@param entity_type Issue enrichment details can be fetched based on either Issue ID or Client MAC address. This parameter value must either be issue_id/mac_address
@param entity_value Contains the actual value for the entity type that has been defined
*/
func (s *IssuesService) GetIssueEnrichmentDetails() (*GetIssueEnrichmentDetailsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/issue-enrichment-details"

	response, err := RestyClient.R().
		SetResult(&GetIssueEnrichmentDetailsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation getIssueEnrichmentDetails")
	}

	result := response.Result().(*GetIssueEnrichmentDetailsResponse)
	return result, response, err
}

// IssuesQueryParams defines the query parameters for this request
type IssuesQueryParams struct {
	StartTime   float64 `url:"startTime,omitempty"`   // Starting epoch time in milliseconds of query time window
	EndTime     float64 `url:"endTime,omitempty"`     // Ending epoch time in milliseconds of query time window
	SiteID      string  `url:"siteId,omitempty"`      // Assurance UUID value of the site in the issue content
	DeviceID    string  `url:"deviceId,omitempty"`    // Assurance UUID value of the device in the issue content
	MacAddress  string  `url:"macAddress,omitempty"`  // Client's device MAC address of the issue (format xx:xx:xx:xx:xx:xx)
	Priority    string  `url:"priority,omitempty"`    // The issue's priority value (One of P1, P2, P3, or P4)(Use only when macAddress and deviceId are not provided)
	AiDriven    string  `url:"aiDriven,omitempty"`    // The issue's AI driven value (Yes or No)(Use only when macAddress and deviceId are not provided)
	IssueStatus string  `url:"issueStatus,omitempty"` // The issue's status value (One of ACTIVE, IGNORED, RESOLVED) (Use only when macAddress and deviceId are not provided)
}

// Issues issues
/* Intent API to get a list of global issues, issues for a specific device, or issue for a specific client device's MAC address.
@param startTime Starting epoch time in milliseconds of query time window
@param endTime Ending epoch time in milliseconds of query time window
@param siteID Assurance UUID value of the site in the issue content
@param deviceID Assurance UUID value of the device in the issue content
@param macAddress Client's device MAC address of the issue (format xx:xx:xx:xx:xx:xx)
@param priority The issue's priority value (One of P1, P2, P3, or P4)(Use only when macAddress and deviceId are not provided)
@param aiDriven The issue's AI driven value (Yes or No)(Use only when macAddress and deviceId are not provided)
@param issueStatus The issue's status value (One of ACTIVE, IGNORED, RESOLVED) (Use only when macAddress and deviceId are not provided)
*/
func (s *IssuesService) Issues(issuesQueryParams *IssuesQueryParams) (*IssuesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/issues"

	queryString, _ := query.Values(issuesQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&IssuesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation issues")
	}

	result := response.Result().(*IssuesResponse)
	return result, response, err
}
