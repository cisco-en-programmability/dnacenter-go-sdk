package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type PlatformService service

type ResponsePlatformCiscoCatalystCenterPackagesSummary struct {
	Response *[]ResponsePlatformCiscoCatalystCenterPackagesSummaryResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
}
type ResponsePlatformCiscoCatalystCenterPackagesSummaryResponse struct {
	Name    string `json:"name,omitempty"`    // Name of installed package
	Version string `json:"version,omitempty"` // Version of installed package
}
type ResponsePlatformCiscoCatalystCenterReleaseSummary struct {
	Version  string                                                     `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformCiscoCatalystCenterReleaseSummaryResponse `json:"response,omitempty"` //
}
type ResponsePlatformCiscoCatalystCenterReleaseSummaryResponse struct {
	CorePackages           []string                                                                           `json:"corePackages,omitempty"`           // The set of packages that are mandatory to be installed/upgraded with the release
	Packages               []string                                                                           `json:"packages,omitempty"`               // The set of packages upgrades available with this release that will not be upgraded unless selected for upgrade
	Name                   string                                                                             `json:"name,omitempty"`                   // Name of the release (example "dnac")
	InstalledVersion       string                                                                             `json:"installedVersion,omitempty"`       // The installed Cisco DNAC version
	SystemVersion          string                                                                             `json:"systemVersion,omitempty"`          // The MAGLEV-SYSTEM version
	SupportedDirectUpdates *[]ResponsePlatformCiscoCatalystCenterReleaseSummaryResponseSupportedDirectUpdates `json:"supportedDirectUpdates,omitempty"` // The list of earlier releases that can upgrade directly to the current release. If the supportedDirectUpdates value is empty, then no direct upgrades to the current release are allowed.
	TenantID               string                                                                             `json:"tenantId,omitempty"`               // Tenant ID (for multi tenant Cisco DNA Center)
}
type ResponsePlatformCiscoCatalystCenterReleaseSummaryResponseSupportedDirectUpdates interface{}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummary struct {
	Version  string                                                                `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponse `json:"response,omitempty"` //
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponse struct {
	Nodes *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodes `json:"nodes,omitempty"` //
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodes struct {
	Ntp      *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNtp       `json:"ntp,omitempty"`      //
	Network  *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetwork `json:"network,omitempty"`  //
	Proxy    *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesProxy     `json:"proxy,omitempty"`    //
	Platform *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesPlatform  `json:"platform,omitempty"` //
	ID       string                                                                              `json:"id,omitempty"`       // Cluster Identifier
	Name     string                                                                              `json:"name,omitempty"`     // Node name
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP server
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetwork struct {
	IntraClusterLink *bool                                                                                  `json:"intra_cluster_link,omitempty"` // Flag to indicate which interface is configured as the inter-cluster link
	LacpMode         *bool                                                                                  `json:"lacp_mode,omitempty"`          // LACP Mode configuration on NIC
	Inet             *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInet  `json:"inet,omitempty"`               //
	Interface        string                                                                                 `json:"interface,omitempty"`          // Interface name
	Inet6            *ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInet6 `json:"inet6,omitempty"`              //
	LacpSupported    *bool                                                                                  `json:"lacp_supported,omitempty"`     // LACP Support configuration on NIC
	SLAve            []string                                                                               `json:"slave,omitempty"`              // Physical interface name
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInet struct {
	Routes     *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInetRoutes     `json:"routes,omitempty"`      // Static route
	Gateway    string                                                                                            `json:"gateway,omitempty"`     // Default gateway
	DNSServers *[]ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInetDNSServers `json:"dns_servers,omitempty"` // DNS server
	Netmask    string                                                                                            `json:"netmask,omitempty"`     // Subnet mask
	HostIP     string                                                                                            `json:"host_ip,omitempty"`     // IP assigned
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInetRoutes interface{}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInetDNSServers interface{}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesNetworkInet6 struct {
	HostIP  string `json:"host_ip,omitempty"` // IP assigned to the host
	Netmask string `json:"netmask,omitempty"` // Subnet mask of the host
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesProxy struct {
	HTTPSProxy         string   `json:"https_proxy,omitempty"`          // Https Proxy Server
	NoProxy            []string `json:"no_proxy,omitempty"`             // Servers configured to explicitly use no proxy
	HTTPSProxyUsername string   `json:"https_proxy_username,omitempty"` // Configured Https proxy username
	HTTPProxy          string   `json:"http_proxy,omitempty"`           // Not Supported
	HTTPSProxyPassword string   `json:"https_proxy_password,omitempty"` // Configured Https excrypted proxy password.
}
type ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryResponseNodesPlatform struct {
	Vendor  string `json:"vendor,omitempty"`  // Product manufacturer
	Product string `json:"product,omitempty"` // Product Identifier
	Serial  string `json:"serial,omitempty"`  // Serial number of chassis
}

//CiscoCatalystCenterPackagesSummary Cisco Catalyst Center Packages Summary - f3aa-697a-453a-bba0
/* Provides information such as name, version of packages installed on the Catalyst center.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-catalyst-center-packages-summary
*/
func (s *PlatformService) CiscoCatalystCenterPackagesSummary() (*ResponsePlatformCiscoCatalystCenterPackagesSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-packages"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoCatalystCenterPackagesSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoCatalystCenterPackagesSummary()
		}
		return nil, response, fmt.Errorf("error with operation CiscoCatalystCenterPackagesSummary")
	}

	result := response.Result().(*ResponsePlatformCiscoCatalystCenterPackagesSummary)
	return result, response, err

}

//CiscoCatalystCenterReleaseSummary Cisco Catalyst Center Release Summary - 5b87-e929-418b-8550
/* Provides information such as API version, mandatory core packages for installation or upgrade, optional packages, Cisco Catalyst Center name and version, supported direct updates, and tenant ID.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-catalyst-center-release-summary
*/
func (s *PlatformService) CiscoCatalystCenterReleaseSummary() (*ResponsePlatformCiscoCatalystCenterReleaseSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-release"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoCatalystCenterReleaseSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoCatalystCenterReleaseSummary()
		}
		return nil, response, fmt.Errorf("error with operation CiscoCatalystCenterReleaseSummary")
	}

	result := response.Result().(*ResponsePlatformCiscoCatalystCenterReleaseSummary)
	return result, response, err

}

//CiscoCatalystCenterNodesConfigurationSummary Cisco Catalyst Center Nodes Configuration Summary - d8b0-fb13-4f08-a967
/* Provides details about the current Cisco Catalyst Center node configuration, such as API version, node name, NTP server, intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface information.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-catalyst-center-nodes-configuration-summary
*/
func (s *PlatformService) CiscoCatalystCenterNodesConfigurationSummary() (*ResponsePlatformCiscoCatalystCenterNodesConfigurationSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/nodes-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoCatalystCenterNodesConfigurationSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoCatalystCenterNodesConfigurationSummary()
		}
		return nil, response, fmt.Errorf("error with operation CiscoCatalystCenterNodesConfigurationSummary")
	}

	result := response.Result().(*ResponsePlatformCiscoCatalystCenterNodesConfigurationSummary)
	return result, response, err

}
