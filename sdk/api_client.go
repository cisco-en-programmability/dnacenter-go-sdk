package dnac

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// RestyClient is the REST Client
var RestyClient *resty.Client

var apiURL = "https://sandboxdnac.cisco.com"

// Client manages communication with the Webex Teams API API v1.0.0
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
	client *Client
}

// SetAuthToken defines the Authorization token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	RestyClient.SetAuthToken(accessToken)
}

// NewClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewClient() (*Client, error) {
	var username = ""
	var password = ""
	client := resty.New()
	c := &Client{}
	RestyClient = client

	if os.Getenv("DNAC_DEBUG") == "true" {
		client.SetDebug(true)
	}
	if os.Getenv("DNAC_SSL_VERIFY") == "false" {
		client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if os.Getenv("DNAC_BASE_URL") != "" {
		RestyClient.SetHostURL(os.Getenv("DNAC_BASE_URL"))
	}
	if os.Getenv("DNAC_USERNAME") != "" {
		username = os.Getenv("DNAC_USERNAME")
	}
	if os.Getenv("DNAC_PASSWORD") != "" {
		password = os.Getenv("DNAC_PASSWORD")
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

// Error indicates an error from the invocation of a Cisco DNA Center API.
type Error struct {
	Error string `json:"error,omitempty"` // Error message
}
