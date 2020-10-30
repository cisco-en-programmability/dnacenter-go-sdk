package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// ConfigurationArchiveService is the service to communicate with the ConfigurationArchive API endpoint
type ConfigurationArchiveService service

// ExportDeviceConfigurations exportDeviceConfigurations
/* Export Device configurations to an encrypted zip file.
@param Content-Type
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurations() (*resty.Response, error) {

	path := "/dna/intent/api/v1/network-device-archive/cleartext"

	response, err := RestyClient.R().
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	return response, err

}
