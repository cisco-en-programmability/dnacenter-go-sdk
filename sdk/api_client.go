package dnac

import (
	"crypto/tls"
	"fmt"
	"os"

	"io/ioutil"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

var apiURL = "https://sandboxdnac.cisco.com"

const DNAC_BASE_URL = "DNAC_BASE_URL"
const DNAC_USERNAME = "DNAC_USERNAME"
const DNAC_PASSWORD = "DNAC_PASSWORD"
const DNAC_DEBUG = "DNAC_DEBUG"
const DNAC_SSL_VERIFY = "DNAC_SSL_VERIFY"

type FileDownload struct {
	FileName string
	FileData []byte
}

func (f *FileDownload) SaveDownload(path string) error {
	fpath := filepath.Join(path, f.FileName)
	return ioutil.WriteFile(fpath, f.FileData, 0664)
}

// Client manages communication with the Cisco DNA Center API
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	Authentication              *AuthenticationService
	ApplicationPolicy           *ApplicationPolicyService
	Applications                *ApplicationsService
	AuthenticationManagement    *AuthenticationManagementService
	Clients                     *ClientsService
	CommandRunner               *CommandRunnerService
	Compliance                  *ComplianceService
	ConfigurationArchive        *ConfigurationArchiveService
	ConfigurationTemplates      *ConfigurationTemplatesService
	DeviceOnboardingPnp         *DeviceOnboardingPnpService
	DeviceReplacement           *DeviceReplacementService
	Devices                     *DevicesService
	DisasterRecovery            *DisasterRecoveryService
	Discovery                   *DiscoveryService
	EventManagement             *EventManagementService
	FabricWireless              *FabricWirelessService
	File                        *FileService
	HealthAndPerformance        *HealthAndPerformanceService
	Itsm                        *ItsmService
	Issues                      *IssuesService
	Licenses                    *LicensesService
	NetworkSettings             *NetworkSettingsService
	PathTrace                   *PathTraceService
	PlatformConfiguration       *PlatformConfigurationService
	Policy                      *PolicyService
	Reports                     *ReportsService
	Sda                         *SdaService
	SecurityAdvisories          *SecurityAdvisoriesService
	Sensors                     *SensorsService
	SiteDesign                  *SiteDesignService
	Sites                       *SitesService
	SoftwareImageManagementSwim *SoftwareImageManagementSwimService
	Tag                         *TagService
	Task                        *TaskService
	Topology                    *TopologyService
	Users                       *UsersService
	Wireless                    *WirelessService
}

type service struct {
	client *resty.Client
}

// SetAuthToken defines the Authorization token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	s.SetAuthToken(accessToken)
}

// Error indicates an error from the invocation of a Cisco DNA Center API.
var Error map[string]interface{}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() (*Client, error) {
	var err error
	var username = ""
	var password = ""
	client := resty.New()
	c := &Client{}
	c.common.client = client

	if os.Getenv(DNAC_DEBUG) == "true" {
		client.SetDebug(true)
	}

	if os.Getenv(DNAC_SSL_VERIFY) == "false" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if os.Getenv(DNAC_BASE_URL) != "" {
		client.SetHostURL(os.Getenv("DNAC_BASE_URL"))
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", DNAC_BASE_URL)
	}

	if os.Getenv(DNAC_USERNAME) != "" {
		username = os.Getenv(DNAC_USERNAME)
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", DNAC_USERNAME)
	}

	if os.Getenv(DNAC_PASSWORD) != "" {
		password = os.Getenv(DNAC_PASSWORD)
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", DNAC_PASSWORD)
	}

	if err != nil {
		return nil, err
	}

	c.Authentication = (*AuthenticationService)(&c.common)
	c.ApplicationPolicy = (*ApplicationPolicyService)(&c.common)
	c.Applications = (*ApplicationsService)(&c.common)
	c.AuthenticationManagement = (*AuthenticationManagementService)(&c.common)
	c.Clients = (*ClientsService)(&c.common)
	c.CommandRunner = (*CommandRunnerService)(&c.common)
	c.Compliance = (*ComplianceService)(&c.common)
	c.ConfigurationArchive = (*ConfigurationArchiveService)(&c.common)
	c.ConfigurationTemplates = (*ConfigurationTemplatesService)(&c.common)
	c.DeviceOnboardingPnp = (*DeviceOnboardingPnpService)(&c.common)
	c.DeviceReplacement = (*DeviceReplacementService)(&c.common)
	c.Devices = (*DevicesService)(&c.common)
	c.DisasterRecovery = (*DisasterRecoveryService)(&c.common)
	c.Discovery = (*DiscoveryService)(&c.common)
	c.EventManagement = (*EventManagementService)(&c.common)
	c.FabricWireless = (*FabricWirelessService)(&c.common)
	c.File = (*FileService)(&c.common)
	c.HealthAndPerformance = (*HealthAndPerformanceService)(&c.common)
	c.Itsm = (*ItsmService)(&c.common)
	c.Issues = (*IssuesService)(&c.common)
	c.Licenses = (*LicensesService)(&c.common)
	c.NetworkSettings = (*NetworkSettingsService)(&c.common)
	c.PathTrace = (*PathTraceService)(&c.common)
	c.PlatformConfiguration = (*PlatformConfigurationService)(&c.common)
	c.Policy = (*PolicyService)(&c.common)
	c.Reports = (*ReportsService)(&c.common)
	c.Sda = (*SdaService)(&c.common)
	c.SecurityAdvisories = (*SecurityAdvisoriesService)(&c.common)
	c.Sensors = (*SensorsService)(&c.common)
	c.SiteDesign = (*SiteDesignService)(&c.common)
	c.Sites = (*SitesService)(&c.common)
	c.SoftwareImageManagementSwim = (*SoftwareImageManagementSwimService)(&c.common)
	c.Tag = (*TagService)(&c.common)
	c.Task = (*TaskService)(&c.common)
	c.Topology = (*TopologyService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Wireless = (*WirelessService)(&c.common)

	result, response, err := c.Authentication.AuthenticationAPI(username, password)
	if err != nil {
		return c, err
	}
	if response.StatusCode() > 399 {
		error := response.Error()
		return c, fmt.Errorf("%s", error)
	}
	client.SetHeader("X-Auth-Token", result.Token)
	return c, nil
}

//NewClientWithOptions is the client with options passed with parameters
func NewClientWithOptions(baseURL string, username string, password string, debug string, sslVerify string) (*Client, error) {
	var err error
	err = os.Setenv("DNAC_BASE_URL", baseURL)
	if err != nil {
		return nil, err
	}
	err = os.Setenv("DNAC_USERNAME", username)
	if err != nil {
		return nil, err
	}
	err = os.Setenv("DNAC_PASSWORD", password)
	if err != nil {
		return nil, err
	}
	err = os.Setenv("DNAC_DEBUG", debug)
	if err != nil {
		return nil, err
	}
	err = os.Setenv("DNAC_SSL_VERIFY", sslVerify)
	if err != nil {
		return nil, err
	}
	return NewClient()
}

// RestyClient returns the resty.Client used by the sdk
func (s *Client) RestyClient() *resty.Client {
	return s.common.client
}
