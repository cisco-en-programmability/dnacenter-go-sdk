package dnac

import (
	"os"

	"github.com/go-resty/resty/v2"
)

// RestyClient is the REST Client
var RestyClient *resty.Client

const apiURL = "https://{defaultHost}"

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
func NewClient() *Client {
	client := resty.New()
	c := &Client{}
	RestyClient = client
	RestyClient.SetHostURL(apiURL)
	if os.Getenv("WEBEX_TEAMS_ACCESS_TOKEN") != "" {
		RestyClient.SetAuthToken(os.Getenv("WEBEX_TEAMS_ACCESS_TOKEN"))
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

	return c
}

// Error indicates an error from the invocation of the API.
type Error struct {
	Message string `json:"message,omitempty"` // Error message
}
