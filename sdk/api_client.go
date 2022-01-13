package dnac

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// RestyClient is the REST Client

var apiURL = "https://sandboxdnac.cisco.com"

const DNAC_BASE_URL = "DNAC_BASE_URL"
const DNAC_USERNAME = "DNAC_USERNAME"
const DNAC_PASSWORD = "DNAC_PASSWORD"
const DNAC_DEBUG = "DNAC_DEBUG"
const DNAC_SSL_VERIFY = "DNAC_SSL_VERIFY"

// Client manages communication with the Cisco DNA Center API v2.1.2
// In most cases there should be only one, shared, APIClient.
type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	Authentication              *AuthenticationService
	Sites                       *SitesService
	Topology                    *TopologyService
	Devices                     *DevicesService
	Clients                     *ClientsService
	Users                       *UsersService
	Issues                      *IssuesService
	SiteDesign                  *SiteDesignService
	NetworkSettings             *NetworkSettingsService
	SoftwareImageManagementSWIM *SoftwareImageManagementSWIMService
	DeviceOnboardingPnP         *DeviceOnboardingPnPService
	ConfigurationTemplates      *ConfigurationTemplatesService
	ConfigurationArchive        *ConfigurationArchiveService
	SDA                         *SDAService
	Sensors                     *SensorsService
	Wireless                    *WirelessService
	CommandRunner               *CommandRunnerService
	Discovery                   *DiscoveryService
	PathTrace                   *PathTraceService
	File                        *FileService
	Task                        *TaskService
	Tag                         *TagService
	ApplicationPolicy           *ApplicationPolicyService
	Applications                *ApplicationsService
	ITSM                        *ITSMService
	EventManagement             *EventManagementService
	DeviceReplacement           *DeviceReplacementService
}

type service struct {
	client *resty.Client
}

// SetAuthToken defines the X-Auth-Token token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	s.common.client.SetHeader("X-Auth-Token", accessToken)
}

// Error indicates an error from the invocation of a Cisco DNA Center API.
type Error struct {
	Error string `json:"error,omitempty"` // Error message
}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() (*Client, error) {
	var err error
	c, err := NewClientNoAuth()
	if err != nil {
		return nil, err
	}

	err = c.AuthClient()
	if err != nil {
		return c, err
	}

	return c, nil
}

//NewClientWithOptions is the client with options passed with parameters
func NewClientWithOptions(baseURL string, username string, password string, debug string, sslVerify string) (*Client, error) {
	var err error

	err = SetOptions(baseURL, username, password, debug, sslVerify)
	if err != nil {
		return nil, err
	}

	return NewClient()
}

//SetOptions sets the environment variables
func SetOptions(baseURL string, username string, password string, debug string, sslVerify string) error {
	var err error
	err = os.Setenv(DNAC_BASE_URL, baseURL)
	if err != nil {
		return err
	}
	err = os.Setenv(DNAC_USERNAME, username)
	if err != nil {
		return err
	}
	err = os.Setenv(DNAC_PASSWORD, password)
	if err != nil {
		return err
	}
	err = os.Setenv(DNAC_DEBUG, debug)
	if err != nil {
		return err
	}
	err = os.Setenv(DNAC_SSL_VERIFY, sslVerify)
	if err != nil {
		return err
	}
	return nil
}

//NewClientNoAuth returns the client object without trying to authenticate
func NewClientNoAuth() (*Client, error) {
	var err error

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
		client.SetHostURL(os.Getenv(DNAC_BASE_URL))
	} else {
		err = fmt.Errorf("enviroment variable %s was not defined", DNAC_BASE_URL)
	}

	if err != nil {
		return nil, err
	}
	// API Services
	c.Authentication = (*AuthenticationService)(&c.common)
	c.Sites = (*SitesService)(&c.common)
	c.Topology = (*TopologyService)(&c.common)
	c.Devices = (*DevicesService)(&c.common)
	c.Clients = (*ClientsService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Issues = (*IssuesService)(&c.common)
	c.SiteDesign = (*SiteDesignService)(&c.common)
	c.NetworkSettings = (*NetworkSettingsService)(&c.common)
	c.SoftwareImageManagementSWIM = (*SoftwareImageManagementSWIMService)(&c.common)
	c.DeviceOnboardingPnP = (*DeviceOnboardingPnPService)(&c.common)
	c.ConfigurationTemplates = (*ConfigurationTemplatesService)(&c.common)
	c.ConfigurationArchive = (*ConfigurationArchiveService)(&c.common)
	c.SDA = (*SDAService)(&c.common)
	c.Sensors = (*SensorsService)(&c.common)
	c.Wireless = (*WirelessService)(&c.common)
	c.CommandRunner = (*CommandRunnerService)(&c.common)
	c.Discovery = (*DiscoveryService)(&c.common)
	c.PathTrace = (*PathTraceService)(&c.common)
	c.File = (*FileService)(&c.common)
	c.Task = (*TaskService)(&c.common)
	c.Tag = (*TagService)(&c.common)
	c.ApplicationPolicy = (*ApplicationPolicyService)(&c.common)
	c.Applications = (*ApplicationsService)(&c.common)
	c.ITSM = (*ITSMService)(&c.common)
	c.EventManagement = (*EventManagementService)(&c.common)
	c.DeviceReplacement = (*DeviceReplacementService)(&c.common)

	return c, nil
}

//NewClientWithOptionsNoAuth returns the client object without trying to authenticate and sets environment variables
func NewClientWithOptionsNoAuth(baseURL string, username string, password string, debug string, sslVerify string) (*Client, error) {
	var err error

	err = SetOptions(baseURL, username, password, debug, sslVerify)
	if err != nil {
		return nil, err
	}

	return NewClientNoAuth()
}

func (s *Client) AuthClient() error {
	var err error

	var username = ""
	var password = ""

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
		return err
	}

	result, response, err := s.Authentication.AuthenticationAPI(username, password)
	if err != nil {
		return err
	}
	if response.StatusCode() > 399 {
		error := response.Error()
		return fmt.Errorf("%s", error)
	}
	s.SetAuthToken(result.Token)

	return nil
}

// RestyClient returns the resty.Client used by the sdk
func (s *Client) RestyClient() *resty.Client {
	return s.common.client
}
