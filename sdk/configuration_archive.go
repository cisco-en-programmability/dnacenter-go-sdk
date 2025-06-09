package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type ConfigurationArchiveService service

type GetConfigurationArchiveDetailsQueryParams struct {
	DeviceID    string  `url:"deviceId,omitempty"`    //comma separated device id for example cf35b0a1-407f-412f-b2f4-f0c3156695f9,aaa38191-0c22-4158-befd-779a09d7cec1 . if device id is not provided it will fetch for all devices
	FileType    string  `url:"fileType,omitempty"`    //Config File Type can be RUNNINGCONFIG or STARTUPCONFIG
	CreatedTime string  `url:"createdTime,omitempty"` //Supported with logical filters GT,GTE,LT,LTE & BT : time in milliseconds (epoc format)
	CreatedBy   string  `url:"createdBy,omitempty"`   //Comma separated values for createdBy - SCHEDULED, USER, CONFIG_CHANGE_EVENT, SCHEDULED_FIRST_TIME, DR_CALL_BACK, PRE_DEPLOY
	Offset      float64 `url:"offset,omitempty"`      //offset
	Limit       float64 `url:"limit,omitempty"`       //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type GetNetworkDeviceConfigurationFileDetailsQueryParams struct {
	ID              string  `url:"id,omitempty"`              //Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string  `url:"networkDeviceId,omitempty"` //Unique identifier (UUID) of the network devices. The number of networkDeviceId(s) must not exceed 5.
	FileType        string  `url:"fileType,omitempty"`        //Type of device configuration file.Available values : 'RUNNINGCONFIG', 'STARTUPCONFIG', 'VLAN'
	Offset          float64 `url:"offset,omitempty"`          //The first record to show for this page; the first record is numbered 1.
	Limit           float64 `url:"limit,omitempty"`           //The number of records to be retrieved defaults to 500 if not specified, with a maximum allowed limit of 500.
}
type CountOfNetworkDeviceConfigurationFilesQueryParams struct {
	ID              string `url:"id,omitempty"`              //Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string `url:"networkDeviceId,omitempty"` //Unique identifier (UUID) of the network devices. The number of networkDeviceId(s) must not exceed 5.
	FileType        string `url:"fileType,omitempty"`        //Type of device configuration file. Available values : 'RUNNINGCONFIG', 'STARTUPCONFIG', 'VLAN'
}

type ResponseConfigurationArchiveExportDeviceConfigurations struct {
	Response *ResponseConfigurationArchiveExportDeviceConfigurationsResponse `json:"response,omitempty"` //
	Version  string                                                          `json:"version,omitempty"`  // Version of API.
}
type ResponseConfigurationArchiveExportDeviceConfigurationsResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task.
	URL    string `json:"url,omitempty"`    // The path to the API endpoint to GET for information on the task.
}
type ResponseConfigurationArchiveGetConfigurationArchiveDetails []ResponseItemConfigurationArchiveGetConfigurationArchiveDetails // Array of ResponseConfigurationArchiveGetConfigurationArchiveDetails
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetails struct {
	IPAddress  string                                                                    `json:"ipAddress,omitempty"`  // IP address of the device.
	DeviceID   string                                                                    `json:"deviceId,omitempty"`   // UUID of the device.
	Versions   *[]ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersions `json:"versions,omitempty"`   //
	DeviceName string                                                                    `json:"deviceName,omitempty"` // Hostname of the device.
}
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersions struct {
	Files                *[]ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersionsFiles                `json:"files,omitempty"`                //
	CreatedBy            string                                                                                        `json:"createdBy,omitempty"`            // Reason for archive collection (CONFIG_CHANGE_EVENT - Syslog event based colletion, SCHEDULED - Weekly scheduled collection, SCHEDULED_FIRST_TIME-First Time Managed collection,DR_CALL_BACK- Post Disaster Recovery collection, USER- On Demand Trigger, PRE_DEPLOY- Predeploy collection)
	ConfigChangeType     string                                                                                        `json:"configChangeType,omitempty"`     // Source of configuration change (OUT_OF_BAND - Change was made outside Catalyst Center, IN_BAND - Change was made from Catalyst Center, NOT_APPLICABLE - System Triggered)
	SyslogConfigEventDto *[]ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersionsSyslogConfigEventDto `json:"syslogConfigEventDto,omitempty"` //
	CreatedTime          *float64                                                                                      `json:"createdTime,omitempty"`          // Source of configuration change (OUT_OF_BAND - Change was made outside Catalyst Center, IN_BAND - Change was made from Catalyst Center, NOT_APPLICABLE - System Triggered)
	StartupRunningStatus string                                                                                        `json:"startupRunningStatus,omitempty"` // Sync status of running and startup configurations (IN_SYNC - if both startup and running config are same excluding dynamic configurations, OUT_OF_SYNC - otherwise).
	ID                   string                                                                                        `json:"id,omitempty"`                   // Unique version ID.
	Tags                 []string                                                                                      `json:"tags,omitempty"`                 // Labels added to a configuration version.
	LastUpdatedTime      *float64                                                                                      `json:"lastUpdatedTime,omitempty"`      // Latest time stamp when the collected configuration is verified to be running on the device. LastUpdatedTime and createdTime will differ when verified configs are the same.
}
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersionsFiles struct {
	FileType     string `json:"fileType,omitempty"`     // Type of configuration file (RUNNINGCONFIG, STARTUPCONFIG, or VLAN).
	FileID       string `json:"fileId,omitempty"`       // Unique file ID.
	DownloadPath string `json:"downloadPath,omitempty"` // File download path (deprecated).
}
type ResponseItemConfigurationArchiveGetConfigurationArchiveDetailsVersionsSyslogConfigEventDto struct {
	UserName       string   `json:"userName,omitempty"`       // Name of the user who made the configuration change (if available in Syslog).
	DeviceUUID     string   `json:"deviceUuid,omitempty"`     // UUID of the device as recieved in syslog.
	OutOfBand      *bool    `json:"outOfBand,omitempty"`      // True if configuration changes were made from outside of the Catalist Center. False otherwise.
	ConfigMethod   string   `json:"configMethod,omitempty"`   // Connection mode used to do the changes pn the device (if available in Syslog).
	TerminalName   string   `json:"terminalName,omitempty"`   // Name of the terminal used to make changes on the device (if available in Syslog).
	LoginIPAddress string   `json:"loginIpAddress,omitempty"` // IP address from which the configuration changes were made (if available in Syslog).
	ProcessName    string   `json:"processName,omitempty"`    // Name of the process that made configuration change (only available when configuration got changed by a program such as YANG suite )
	SyslogTime     *float64 `json:"syslogTime,omitempty"`     // Time of configuration change as recorded in the syslog.
}
type ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetails struct {
	Response *[]ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsResponse `json:"response,omitempty"` //
	Version  string                                                                          `json:"version,omitempty"`  // The version of API.
}
type ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetailsResponse struct {
	ID              string `json:"id,omitempty"`              // Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Unique identifier (UUID) of the network devices.
	VersionID       string `json:"versionId,omitempty"`       // The version unique identifier triggered after any config change.
	FileType        string `json:"fileType,omitempty"`        // Type of configuration file. Config File Type can be 'RUNNINGCONFIG' or 'STARTUPCONFIG' or 'VLAN'.
	CreatedBy       string `json:"createdBy,omitempty"`       // The entity responsible for creating the configuration changes.
	CreatedTime     *int   `json:"createdTime,omitempty"`     // The UNIX epoch timestamp in milliseconds marking when the resource was created.
}
type ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFiles struct {
	Response *ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesResponse `json:"response,omitempty"` //
	Version  string                                                                      `json:"version,omitempty"`  // Version
}
type ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesResponse struct {
	Count *int `json:"count,omitempty"` // Count
}
type ResponseConfigurationArchiveGetConfigurationFileDetailsByID struct {
	Response *ResponseConfigurationArchiveGetConfigurationFileDetailsByIDResponse `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  // The version of API.
}
type ResponseConfigurationArchiveGetConfigurationFileDetailsByIDResponse struct {
	ID              string `json:"id,omitempty"`              // Unique identifier (UUID) of the configuration file.
	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Unique identifier (UUID) of the network devices.
	VersionID       string `json:"versionId,omitempty"`       // The version unique identifier triggered after any config change.
	FileType        string `json:"fileType,omitempty"`        // Type of configuration file. Config File Type can be 'RUNNINGCONFIG' or 'STARTUPCONFIG' or 'VLAN'.
	CreatedBy       string `json:"createdBy,omitempty"`       // The entity responsible for creating the configuration changes.
	CreatedTime     string `json:"createdTime,omitempty"`     // The UNIX epoch timestamp in milliseconds marking when the resource was created.
}
type ResponseConfigurationArchiveDownloadMaskedDeviceConfiguration interface{}
type ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP interface{}
type RequestConfigurationArchiveExportDeviceConfigurations struct {
	DeviceID []string `json:"deviceId,omitempty"` // UUIDs of the devices for which configurations need to be exported.
	Password string   `json:"password,omitempty"` // Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.
}
type RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP struct {
	Password string `json:"password,omitempty"` // Password for the zip file to protect exported configurations. Must contain, at minimum 8 characters, one lowercase letter, one uppercase letter, one number, one special character(-=[];,./~!@#$%^&*()_+{}|:?). It may not contain white space or the characters <>.
}

//GetConfigurationArchiveDetails Get configuration archive details - 3bba-48a9-422a-be1e
/* Returns the historical device configurations (running configuration , startup configuration , vlan if applicable) by specified criteria


@param GetConfigurationArchiveDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configuration-archive-details
*/
func (s *ConfigurationArchiveService) GetConfigurationArchiveDetails(GetConfigurationArchiveDetailsQueryParams *GetConfigurationArchiveDetailsQueryParams) (*ResponseConfigurationArchiveGetConfigurationArchiveDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-config"

	queryString, _ := query.Values(GetConfigurationArchiveDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationArchiveGetConfigurationArchiveDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationArchiveDetails(GetConfigurationArchiveDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigurationArchiveDetails")
	}

	result := response.Result().(*ResponseConfigurationArchiveGetConfigurationArchiveDetails)
	return result, response, err

}

//GetNetworkDeviceConfigurationFileDetails Get Network Device Configuration File Details - bd95-9a71-4b8a-9442
/* Retrieves the list of network device configuration file details, sorted by createdTime in descending order. Use /intent/api/v1/networkDeviceConfigFiles/{id}/downloadMasked to download masked configurations, or /intent/api/v1/networkDeviceConfigFiles/{id}/downloadUnmasked for unmasked configurations.


@param GetNetworkDeviceConfigurationFileDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-configuration-file-details
*/
func (s *ConfigurationArchiveService) GetNetworkDeviceConfigurationFileDetails(GetNetworkDeviceConfigurationFileDetailsQueryParams *GetNetworkDeviceConfigurationFileDetailsQueryParams) (*ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles"

	queryString, _ := query.Values(GetNetworkDeviceConfigurationFileDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceConfigurationFileDetails(GetNetworkDeviceConfigurationFileDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceConfigurationFileDetails")
	}

	result := response.Result().(*ResponseConfigurationArchiveGetNetworkDeviceConfigurationFileDetails)
	return result, response, err

}

//CountOfNetworkDeviceConfigurationFiles Count of Network Device Configuration Files - d296-cab3-4a6b-b826
/* Retrieves count the details of the network device configuration files.


@param CountOfNetworkDeviceConfigurationFilesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-network-device-configuration-files
*/
func (s *ConfigurationArchiveService) CountOfNetworkDeviceConfigurationFiles(CountOfNetworkDeviceConfigurationFilesQueryParams *CountOfNetworkDeviceConfigurationFilesQueryParams) (*ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/count"

	queryString, _ := query.Values(CountOfNetworkDeviceConfigurationFilesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNetworkDeviceConfigurationFiles(CountOfNetworkDeviceConfigurationFilesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNetworkDeviceConfigurationFiles")
	}

	result := response.Result().(*ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFiles)
	return result, response, err

}

//GetConfigurationFileDetailsByID Get Configuration File Details by ID - cc93-5822-44ab-b75f
/* Retrieves the details of a specific network device configuration file using the `id`.


@param id id path parameter. The value of `id` can be obtained from the response of API `/dna/intent/api/v1/networkDeviceConfigFiles`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configuration-file-details-by-id
*/
func (s *ConfigurationArchiveService) GetConfigurationFileDetailsByID(id string) (*ResponseConfigurationArchiveGetConfigurationFileDetailsByID, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseConfigurationArchiveGetConfigurationFileDetailsByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationFileDetailsByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigurationFileDetailsById")
	}

	result := response.Result().(*ResponseConfigurationArchiveGetConfigurationFileDetailsByID)
	return result, response, err

}

//ExportDeviceConfigurations Export Device configurations - 51a4-0aba-4c68-ac17
/* Export Device configuration for every device that is provided will be included in an encrypted zip file.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!export-device-configurations
*/
func (s *ConfigurationArchiveService) ExportDeviceConfigurations(requestConfigurationArchiveExportDeviceConfigurations *RequestConfigurationArchiveExportDeviceConfigurations) (*ResponseConfigurationArchiveExportDeviceConfigurations, *resty.Response, error) {
	path := "/dna/intent/api/v1/network-device-archive/cleartext"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationArchiveExportDeviceConfigurations).
		SetResult(&ResponseConfigurationArchiveExportDeviceConfigurations{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ExportDeviceConfigurations(requestConfigurationArchiveExportDeviceConfigurations)
		}

		return nil, response, fmt.Errorf("error with operation ExportDeviceConfigurations")
	}

	result := response.Result().(*ResponseConfigurationArchiveExportDeviceConfigurations)
	return result, response, err

}

//DownloadMaskedDeviceConfiguration Download masked device configuration - fe93-185d-4c58-a302
/* Download the masked (sanitized) device configuration by providing the file `id`.


@param id id path parameter. The value of `id` can be obtained from the response of API `/dna/intent/api/v1/networkDeviceConfigFiles`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-masked-device-configuration
*/
func (s *ConfigurationArchiveService) DownloadMaskedDeviceConfiguration(id string) (*ResponseConfigurationArchiveDownloadMaskedDeviceConfiguration, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/{id}/downloadMasked"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").

		// SetResult(&ResponseConfigurationArchiveDownloadMaskedDeviceConfiguration{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadMaskedDeviceConfiguration(id)
		}

		return nil, response, fmt.Errorf("error with operation DownloadMaskedDeviceConfiguration")
	}

	result := response.Result().(ResponseConfigurationArchiveDownloadMaskedDeviceConfiguration)

	return &result, response, err

}

//DownloadUnmaskedrawDeviceConfigurationAsZIP Download Unmasked (raw) Device Configuration as ZIP - 59a7-7a49-4e79-8fde
/* Download the unmasked (raw) device configuration by providing the file `id` and a `password`. The response will be a password-protected zip file containing the unmasked configuration. Password must contain a minimum of 8 characters, one lowercase letter, one uppercase letter, one number, one special character (`-=[];,./~!@#$%^&*()_+{}|:?`). It may not contain white space or the characters `<>`.


@param id id path parameter. The value of `id` can be obtained from the response of API `/dna/intent/api/v1/networkDeviceConfigFiles`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-unmaskedraw-device-configuration-as-z-ip
*/
func (s *ConfigurationArchiveService) DownloadUnmaskedrawDeviceConfigurationAsZIP(id string, requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP *RequestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP) (*ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceConfigFiles/{id}/downloadUnmasked"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP).
		// SetResult(&ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadUnmaskedrawDeviceConfigurationAsZIP(id, requestConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP)
		}

		return nil, response, fmt.Errorf("error with operation DownloadUnmaskedrawDeviceConfigurationAsZIp")
	}

	result := response.Result().(ResponseConfigurationArchiveDownloadUnmaskedrawDeviceConfigurationAsZIP)

	return &result, response, err

}
