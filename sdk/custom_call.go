package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type CustomCallService service

func (s *CustomCallService) GetCustomCall(QueryParms *interface{}, ResourcePath string) (*resty.Response, error) {
	path := ResourcePath

	queryString, _ := query.Values(QueryParms)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationTemplatesGetsAListOfProjects{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetsAListOfProjects")
	}

	return response, err
}
