package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type CustomCallService service

func (s *CustomCallService) GetCustomCall(ResourcePath string, QueryParms *map[string]string) (*resty.Response, error) {
	path := ResourcePath
	var response *resty.Response
	var err error

	if QueryParms != nil {
		response, err = s.client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetQueryParams(*QueryParms).
			SetError(&Error).
			Get(path)
	} else {
		response, err = s.client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json").
			SetError(&Error).
			Get(path)
	}

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with operation GetsAListOfProjects")
	}

	return response, err
}
