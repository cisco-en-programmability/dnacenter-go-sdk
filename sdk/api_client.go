package dnac

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

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
const DNAC_WAIT_TIME = "DNAC_WAIT_TIME"

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
	CiscoDnaCenterSystem        *CiscoDnaCenterSystemService
	Clients                     *ClientsService
	CommandRunner               *CommandRunnerService
	Compliance                  *ComplianceService
	ConfigurationArchive        *ConfigurationArchiveService
	ConfigurationTemplates      *ConfigurationTemplatesService
	DeviceOnboardingPnp         *DeviceOnboardingPnpService
	DeviceReplacement           *DeviceReplacementService
	Devices                     *DevicesService
	Discovery                   *DiscoveryService
	EoX                         *EoXService
	EventManagement             *EventManagementService
	FabricWireless              *FabricWirelessService
	File                        *FileService
	HealthAndPerformance        *HealthAndPerformanceService
	Itsm                        *ItsmService
	ItsmIntegration             *ItsmIntegrationService
	Issues                      *IssuesService
	LanAutomation               *LanAutomationService
	Licenses                    *LicensesService
	NetworkSettings             *NetworkSettingsService
	PathTrace                   *PathTraceService
	Platform                    *PlatformService
	Reports                     *ReportsService
	Sda                         *SdaService
	SecurityAdvisories          *SecurityAdvisoriesService
	Sensors                     *SensorsService
	SiteDesign                  *SiteDesignService
	Sites                       *SitesService
	SoftwareImageManagementSwim *SoftwareImageManagementSwimService
	SystemSettings              *SystemSettingsService
	Tag                         *TagService
	Task                        *TaskService
	Topology                    *TopologyService
	UserandRoles                *UserandRolesService
	Users                       *UsersService
	Wireless                    *WirelessService
	CustomCall                  *CustomCallService
}

type service struct {
	client *resty.Client
}

// SetAuthToken defines the X-Auth-Token token sent in the request
func (s *Client) SetAuthToken(accessToken string) {
	s.common.client.SetHeader("X-Auth-Token", accessToken)
}

// Error indicates an error from the invocation of a Cisco DNA Center API.
var Error map[string]interface{}

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

// NewClientWithOptions is the client with options passed with parameters
func NewClientWithOptions(baseURL string, username string, password string, debug string, sslVerify string, waitTimeToManyRequest *int) (*Client, error) {
	var err error

	err = SetOptions(baseURL, username, password, debug, sslVerify, waitTimeToManyRequest)
	if err != nil {
		return nil, err
	}

	return NewClient()
}

// SetOptions sets the environment variables
func SetOptions(baseURL string, username string, password string, debug string, sslVerify string, waitTimeToManyRequest *int) error {
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
	if waitTimeToManyRequest != nil {
		err = os.Setenv(DNAC_WAIT_TIME, strconv.Itoa(*waitTimeToManyRequest))
		if err != nil {
			return err
		}
	} else {
		err = os.Setenv(DNAC_WAIT_TIME, "1")
		if err != nil {
			return err
		}
	}
	return nil
}

// NewClientNoAuth returns the client object without trying to authenticate
func NewClientNoAuth() (*Client, error) {
	var err error

	client := resty.New()
	c := &Client{}
	c.common.client = client
	waitTimeToManyRequest := 0
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
	if os.Getenv(DNAC_WAIT_TIME) != "" {
		waitTimeToManyRequest, err = strconv.Atoi(os.Getenv(DNAC_WAIT_TIME))
		if err != nil {
			waitTimeToManyRequest = 1
		}
	} else {
		waitTimeToManyRequest = 1
	}
	c.common.client.AddRetryCondition(
		// RetryConditionFunc type is for retry condition function
		// input: non-nil Response OR request execution error
		func(r *resty.Response, err error) bool {
			if r.StatusCode() == http.StatusUnauthorized {
				cl := resty.New()

				username := os.Getenv("DNAC_USERNAME")
				password := os.Getenv("DNAC_PASSWORD")
				baseUrl := os.Getenv("DNAC_BASE_URL")
				if os.Getenv(DNAC_SSL_VERIFY) == "false" {
					cl.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
				}

				cl.SetBaseURL(baseUrl)

				response, err := cl.R().
					SetBasicAuth(username, password).
					SetResult(&AuthenticationAPIResponse{}).
					Post("dna/system/api/v1/auth/token")
				if err != nil {
					log.Printf("Err: %s", err.Error())
					return false
				}
				result := response.Result().(*AuthenticationAPIResponse)
				// log.Printf("resty: %s", response.String())
				// request.SetHeader("X-auth-token", result.Token)
				// request.SetHeader("X-auth-token", result.Token)
				c.common.client.SetHeader("X-auth-token", result.Token)
				r.Request.SetHeader("X-auth-token", result.Token)
			}

			return r.StatusCode() == http.StatusTooManyRequests || r.StatusCode() == http.StatusUnauthorized
		},
	)
	c.common.client.
		// Set retry count to non zero to enable retries
		SetRetryCount(1).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(time.Duration(waitTimeToManyRequest) * time.Minute).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(time.Duration(waitTimeToManyRequest+1) * time.Minute)
	if err != nil {
		return nil, err
	}

	c.Authentication = (*AuthenticationService)(&c.common)
	c.ApplicationPolicy = (*ApplicationPolicyService)(&c.common)
	c.Applications = (*ApplicationsService)(&c.common)
	c.AuthenticationManagement = (*AuthenticationManagementService)(&c.common)
	c.CiscoDnaCenterSystem = (*CiscoDnaCenterSystemService)(&c.common)
	c.Clients = (*ClientsService)(&c.common)
	c.CommandRunner = (*CommandRunnerService)(&c.common)
	c.Compliance = (*ComplianceService)(&c.common)
	c.ConfigurationArchive = (*ConfigurationArchiveService)(&c.common)
	c.ConfigurationTemplates = (*ConfigurationTemplatesService)(&c.common)
	c.DeviceOnboardingPnp = (*DeviceOnboardingPnpService)(&c.common)
	c.DeviceReplacement = (*DeviceReplacementService)(&c.common)
	c.Devices = (*DevicesService)(&c.common)
	c.Discovery = (*DiscoveryService)(&c.common)
	c.EoX = (*EoXService)(&c.common)
	c.EventManagement = (*EventManagementService)(&c.common)
	c.FabricWireless = (*FabricWirelessService)(&c.common)
	c.File = (*FileService)(&c.common)
	c.HealthAndPerformance = (*HealthAndPerformanceService)(&c.common)
	c.Itsm = (*ItsmService)(&c.common)
	c.ItsmIntegration = (*ItsmIntegrationService)(&c.common)
	c.Issues = (*IssuesService)(&c.common)
	c.LanAutomation = (*LanAutomationService)(&c.common)
	c.Licenses = (*LicensesService)(&c.common)
	c.NetworkSettings = (*NetworkSettingsService)(&c.common)
	c.PathTrace = (*PathTraceService)(&c.common)
	c.Platform = (*PlatformService)(&c.common)
	c.Reports = (*ReportsService)(&c.common)
	c.Sda = (*SdaService)(&c.common)
	c.SecurityAdvisories = (*SecurityAdvisoriesService)(&c.common)
	c.Sensors = (*SensorsService)(&c.common)
	c.SiteDesign = (*SiteDesignService)(&c.common)
	c.Sites = (*SitesService)(&c.common)
	c.SoftwareImageManagementSwim = (*SoftwareImageManagementSwimService)(&c.common)
	c.SystemSettings = (*SystemSettingsService)(&c.common)
	c.Tag = (*TagService)(&c.common)
	c.Task = (*TaskService)(&c.common)
	c.Topology = (*TopologyService)(&c.common)
	c.UserandRoles = (*UserandRolesService)(&c.common)
	c.Users = (*UsersService)(&c.common)
	c.Wireless = (*WirelessService)(&c.common)
	c.CustomCall = (*CustomCallService)(&c.common)

	return c, nil
}

// NewClientWithOptionsNoAuth returns the client object without trying to authenticate and sets environment variables
func NewClientWithOptionsNoAuth(baseURL string, username string, password string, debug string, sslVerify string, waitTimeToManyRequest *int) (*Client, error) {
	var err error

	err = SetOptions(baseURL, username, password, debug, sslVerify, waitTimeToManyRequest)
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

func (s *Client) SetDNACWaitTimeToManyRequest(waitTimeToManyRequest int) error {

	err := os.Setenv(DNAC_WAIT_TIME, strconv.Itoa(waitTimeToManyRequest))
	if err != nil {
		return err
	}
	s.common.client.
		// Set retry count to non zero to enable retries
		SetRetryCount(1).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(time.Duration(waitTimeToManyRequest) * time.Minute).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(time.Duration(waitTimeToManyRequest+1) * time.Minute)
	return err
}
