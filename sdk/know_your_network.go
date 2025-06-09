package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type KnowYourNetworkService service

type GetEnergySummaryAnalyticsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}
type GetEnergyTrendAnalyticsHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	XCaLLERID   string `url:"X-CALLER-ID,omitempty"`  //Expects type string. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
}

type ResponseKnowYourNetworkGetEnergySummaryAnalytics struct {
	Response *ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponse `json:"response,omitempty"` //
	Page     *ResponseKnowYourNetworkGetEnergySummaryAnalyticsPage     `json:"page,omitempty"`     //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponse struct {
	Groups *[]ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponseGroups `json:"groups,omitempty"` //
}
type ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponseGroups struct {
	ID                  string                                                                               `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseKnowYourNetworkGetEnergySummaryAnalyticsResponseGroupsAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseKnowYourNetworkGetEnergySummaryAnalyticsPage struct {
	Limit  *int                                                          `json:"limit,omitempty"`  // Limit
	Offset *int                                                          `json:"offset,omitempty"` // Offset
	Count  *int                                                          `json:"count,omitempty"`  // Count
	SortBy *[]ResponseKnowYourNetworkGetEnergySummaryAnalyticsPageSortBy `json:"sortBy,omitempty"` //
}
type ResponseKnowYourNetworkGetEnergySummaryAnalyticsPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type ResponseKnowYourNetworkGetEnergyTrendAnalytics struct {
	Response *[]ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponse `json:"response,omitempty"` //
	Page     *ResponseKnowYourNetworkGetEnergyTrendAnalyticsPage       `json:"page,omitempty"`     //
	Version  string                                                    `json:"version,omitempty"`  // Version
}
type ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponse struct {
	Groups    *[]ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponseGroups `json:"groups,omitempty"`    //
	Timestamp *int                                                            `json:"timestamp,omitempty"` // Timestamp
}
type ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponseGroups struct {
	ID                  string                                                                             `json:"id,omitempty"`                  // Id
	Attributes          *[]ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponseGroupsAttributes          `json:"attributes,omitempty"`          //
	AggregateAttributes *[]ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponseGroupsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
}
type ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponseGroupsAttributes struct {
	Name  string `json:"name,omitempty"`  // Name
	Value string `json:"value,omitempty"` // Value
}
type ResponseKnowYourNetworkGetEnergyTrendAnalyticsResponseGroupsAggregateAttributes struct {
	Name     string   `json:"name,omitempty"`     // Name
	Function string   `json:"function,omitempty"` // Function
	Value    *float64 `json:"value,omitempty"`    // Value
}
type ResponseKnowYourNetworkGetEnergyTrendAnalyticsPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	Count          *int   `json:"count,omitempty"`          // Count
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}
type RequestKnowYourNetworkGetEnergySummaryAnalytics struct {
	StartTime           *int                                                                  `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                  `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestKnowYourNetworkGetEnergySummaryAnalyticsFilters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                              `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                              `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestKnowYourNetworkGetEnergySummaryAnalyticsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestKnowYourNetworkGetEnergySummaryAnalyticsPage                  `json:"page,omitempty"`                //
}
type RequestKnowYourNetworkGetEnergySummaryAnalyticsFilters struct {
	LogicalOperator string                                                           `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestKnowYourNetworkGetEnergySummaryAnalyticsFiltersFilters `json:"filters,omitempty"`         //
}
type RequestKnowYourNetworkGetEnergySummaryAnalyticsFiltersFilters struct {
	Key      string   `json:"key,omitempty"`      // Key
	Operator string   `json:"operator,omitempty"` // Operator
	Value    []string `json:"value,omitempty"`    // Value
}
type RequestKnowYourNetworkGetEnergySummaryAnalyticsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestKnowYourNetworkGetEnergySummaryAnalyticsPage struct {
	Limit  *int                                                         `json:"limit,omitempty"`  // Limit
	Offset *int                                                         `json:"offset,omitempty"` // Offset
	SortBy *[]RequestKnowYourNetworkGetEnergySummaryAnalyticsPageSortBy `json:"sortBy,omitempty"` //
}
type RequestKnowYourNetworkGetEnergySummaryAnalyticsPageSortBy struct {
	Name     string `json:"name,omitempty"`     // Name
	Order    string `json:"order,omitempty"`    // Order
	Function string `json:"function,omitempty"` // Function
}
type RequestKnowYourNetworkGetEnergyTrendAnalytics struct {
	StartTime           *int                                                                `json:"startTime,omitempty"`           // Start Time
	EndTime             *int                                                                `json:"endTime,omitempty"`             // End Time
	Filters             *[]RequestKnowYourNetworkGetEnergyTrendAnalyticsFilters             `json:"filters,omitempty"`             //
	GroupBy             []string                                                            `json:"groupBy,omitempty"`             // Group By
	Attributes          []string                                                            `json:"attributes,omitempty"`          // Attributes
	AggregateAttributes *[]RequestKnowYourNetworkGetEnergyTrendAnalyticsAggregateAttributes `json:"aggregateAttributes,omitempty"` //
	Page                *RequestKnowYourNetworkGetEnergyTrendAnalyticsPage                  `json:"page,omitempty"`                //
}
type RequestKnowYourNetworkGetEnergyTrendAnalyticsFilters struct {
	LogicalOperator string                                                         `json:"logicalOperator,omitempty"` // Logical Operator
	Filters         *[]RequestKnowYourNetworkGetEnergyTrendAnalyticsFiltersFilters `json:"filters,omitempty"`         //
}
type RequestKnowYourNetworkGetEnergyTrendAnalyticsFiltersFilters struct {
	Key      string   `json:"key,omitempty"`      // Key
	Operator string   `json:"operator,omitempty"` // Operator
	Value    []string `json:"value,omitempty"`    // Value
}
type RequestKnowYourNetworkGetEnergyTrendAnalyticsAggregateAttributes struct {
	Name     string `json:"name,omitempty"`     // Name
	Function string `json:"function,omitempty"` // Function
}
type RequestKnowYourNetworkGetEnergyTrendAnalyticsPage struct {
	Limit          *int   `json:"limit,omitempty"`          // Limit
	Offset         *int   `json:"offset,omitempty"`         // Offset
	TimestampOrder string `json:"timestampOrder,omitempty"` // Timestamp Order
}

//GetEnergySummaryAnalytics Get energy summary analytics - f190-0901-439a-80b4
/* Retrieve the summary analytics data related to device energy consumption for all devices, including network devices and clients assigned to specific sites. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param GetEnergySummaryAnalyticsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-energy-summary-analytics
*/
func (s *KnowYourNetworkService) GetEnergySummaryAnalytics(requestKnowYourNetworkGetEnergySummaryAnalytics *RequestKnowYourNetworkGetEnergySummaryAnalytics, GetEnergySummaryAnalyticsHeaderParams *GetEnergySummaryAnalyticsHeaderParams) (*ResponseKnowYourNetworkGetEnergySummaryAnalytics, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/summaryAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetEnergySummaryAnalyticsHeaderParams != nil {

		if GetEnergySummaryAnalyticsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetEnergySummaryAnalyticsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestKnowYourNetworkGetEnergySummaryAnalytics).
		SetResult(&ResponseKnowYourNetworkGetEnergySummaryAnalytics{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEnergySummaryAnalytics(requestKnowYourNetworkGetEnergySummaryAnalytics, GetEnergySummaryAnalyticsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetEnergySummaryAnalytics")
	}

	result := response.Result().(*ResponseKnowYourNetworkGetEnergySummaryAnalytics)
	return result, response, err

}

//GetEnergyTrendAnalytics Get energy trend analytics - dba3-f890-4c09-a89c
/* Retrieve the energy trend analytics data related to device energy consumption for all devices, including network devices and clients assigned to specific sites. For detailed information about the usage of the API, please refer to the Open API specification document - https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-deviceEnergy_1.0-1.0.1-resolved.yaml


@param GetEnergyTrendAnalyticsHeaderParams Custom header parameters

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-energy-trend-analytics
*/
func (s *KnowYourNetworkService) GetEnergyTrendAnalytics(requestKnowYourNetworkGetEnergyTrendAnalytics *RequestKnowYourNetworkGetEnergyTrendAnalytics, GetEnergyTrendAnalyticsHeaderParams *GetEnergyTrendAnalyticsHeaderParams) (*ResponseKnowYourNetworkGetEnergyTrendAnalytics, *resty.Response, error) {
	path := "/dna/data/api/v1/energy/trendAnalytics"

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if GetEnergyTrendAnalyticsHeaderParams != nil {

		if GetEnergyTrendAnalyticsHeaderParams.XCaLLERID != "" {
			clientRequest = clientRequest.SetHeader("X-CALLER-ID", GetEnergyTrendAnalyticsHeaderParams.XCaLLERID)
		}

	}

	response, err = clientRequest.
		SetBody(requestKnowYourNetworkGetEnergyTrendAnalytics).
		SetResult(&ResponseKnowYourNetworkGetEnergyTrendAnalytics{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetEnergyTrendAnalytics(requestKnowYourNetworkGetEnergyTrendAnalytics, GetEnergyTrendAnalyticsHeaderParams)
		}

		return nil, response, fmt.Errorf("error with operation GetEnergyTrendAnalytics")
	}

	result := response.Result().(*ResponseKnowYourNetworkGetEnergyTrendAnalytics)
	return result, response, err

}
