package dnac

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type PlatformService service

type ResponsePlatformCiscoDnaCenterPackagesSummary struct {
	Response *[]ResponsePlatformCiscoDnaCenterPackagesSummaryResponse `json:"response,omitempty"` //
	Version  string                                                   `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
}
type ResponsePlatformCiscoDnaCenterPackagesSummaryResponse struct {
	Name    string `json:"name,omitempty"`    // Name of installed package
	Version string `json:"version,omitempty"` // Version of installed package
}
type ResponsePlatformCiscoDnaCenterReleaseSummary struct {
	Version  string                                                `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformCiscoDnaCenterReleaseSummaryResponse `json:"response,omitempty"` //
}
type ResponsePlatformCiscoDnaCenterReleaseSummaryResponse struct {
	CorePackages []string `json:"corePackages,omitempty"` // The set of packages that are mandatory to be installed/upgraded with the release

	Packages []string `json:"packages,omitempty"` // The set of packages upgrades available with this release that will not be upgraded unless selected for upgrade

	Name string `json:"name,omitempty"` // Name of the release (example "dnac")

	InstalledVersion string `json:"installedVersion,omitempty"` // The installed Cisco Dna Center version

	SystemVersion string `json:"systemVersion,omitempty"` // The MAGLEV-SYSTEM version

	SupportedDirectUpdates *[]ResponsePlatformCiscoDnaCenterReleaseSummaryResponseSupportedDirectUpdates `json:"supportedDirectUpdates,omitempty"` // The list of earlier releases that can upgrade directly to the current release. If the supportedDirectUpdates value is empty, then no direct upgrades to the current release are allowed.

	TenantID string `json:"tenantId,omitempty"` // Tenant ID (for multi tenant Cisco Dna Center)
}
type ResponsePlatformCiscoDnaCenterReleaseSummaryResponseSupportedDirectUpdates interface{}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummary struct {
	Version string `json:"version,omitempty"` // The MAGLEV-API version (this field is for internal development purpose)

	Response *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponse `json:"response,omitempty"` //
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponse struct {
	Nodes *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodes `json:"nodes,omitempty"` //
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodes struct {
	Ntp *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNtp `json:"ntp,omitempty"` //

	Network *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetwork `json:"network,omitempty"` //

	Proxy *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesProxy `json:"proxy,omitempty"` //

	Platform *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesPlatform `json:"platform,omitempty"` //

	ID string `json:"id,omitempty"` // Cluster Identifier

	Name string `json:"name,omitempty"` // Node name
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP server
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetwork struct {
	IntraClusterLink *bool `json:"intra_cluster_link,omitempty"` // Flag to indicate which interface is configured as the inter-cluster link

	LacpMode *bool `json:"lacp_mode,omitempty"` // LACP Mode configuration on NIC

	Inet *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet `json:"inet,omitempty"` //

	Interface string `json:"interface,omitempty"` // Interface name

	Inet6 *ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet6 `json:"inet6,omitempty"` //

	LacpSupported *bool `json:"lacp_supported,omitempty"` // LACP Support configuration on NIC

	SLAve []string `json:"slave,omitempty"` // Physical interface name
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet struct {
	Routes *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetRoutes `json:"routes,omitempty"` // Static route

	Gateway string `json:"gateway,omitempty"` // Default gateway

	DNSServers *[]ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetDNSServers `json:"dns_servers,omitempty"` // DNS server

	Netmask string `json:"netmask,omitempty"` // Subnet mask

	HostIP string `json:"host_ip,omitempty"` // IP assigned
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetRoutes interface{}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetDNSServers interface{}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet6 struct {
	HostIP string `json:"host_ip,omitempty"` // IP assigned to the host

	Netmask string `json:"netmask,omitempty"` // Subnet mask of the host
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesProxy struct {
	HTTPSProxy string `json:"https_proxy,omitempty"` // Https Proxy Server

	NoProxy []string `json:"no_proxy,omitempty"` // Servers configured to explicitly use no proxy

	HTTPSProxyUsername string `json:"https_proxy_username,omitempty"` // Configured Https proxy username

	HTTPProxy string `json:"http_proxy,omitempty"` // Not Supported

	HTTPSProxyPassword string `json:"https_proxy_password,omitempty"` // Configured Https excrypted proxy password.
}
type ResponsePlatformCiscoDnaCenterNodesConfigurationSummaryResponseNodesPlatform struct {
	Vendor string `json:"vendor,omitempty"` // Product manufacturer

	Product string `json:"product,omitempty"` // Product Identifier

	Serial string `json:"serial,omitempty"` // Serial number of chassis
}

//CiscoDnaCenterPackagesSummary Cisco Dna Center Packages Summary - f3aa-697a-453a-bba0
/* Provides information such as name, version of packages installed on the Dna center.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-packages-summary
*/
func (s *PlatformService) CiscoDnaCenterPackagesSummary() (*ResponsePlatformCiscoDnaCenterPackagesSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-packages"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoDnaCenterPackagesSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoDnaCenterPackagesSummary()
		}
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterPackagesSummary")
	}

	result := response.Result().(*ResponsePlatformCiscoDnaCenterPackagesSummary)
	return result, response, err

}

//CiscoDnaCenterReleaseSummary Cisco Dna Center Release Summary - 5b87-e929-418b-8550
/* Provides information such as API version, mandatory core packages for installation or upgrade, optional packages, Cisco Dna Center name and version, supported direct updates, and tenant ID.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-release-summary
*/
func (s *PlatformService) CiscoDnaCenterReleaseSummary() (*ResponsePlatformCiscoDnaCenterReleaseSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-release"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoDnaCenterReleaseSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoDnaCenterReleaseSummary()
		}
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterReleaseSummary")
	}

	result := response.Result().(*ResponsePlatformCiscoDnaCenterReleaseSummary)
	return result, response, err

}

//CiscoDnaCenterNodesConfigurationSummary Cisco Dna Center Nodes Configuration Summary - d8b0-fb13-4f08-a967
/* Provides details about the current Cisco Dna Center node configuration, such as API version, node name, NTP server, intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface information.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-nodes-configuration-summary
*/
func (s *PlatformService) CiscoDnaCenterNodesConfigurationSummary() (*ResponsePlatformCiscoDnaCenterNodesConfigurationSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/nodes-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformCiscoDnaCenterNodesConfigurationSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CiscoDnaCenterNodesConfigurationSummary()
		}
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterNodesConfigurationSummary")
	}

	result := response.Result().(*ResponsePlatformCiscoDnaCenterNodesConfigurationSummary)
	return result, response, err

}
