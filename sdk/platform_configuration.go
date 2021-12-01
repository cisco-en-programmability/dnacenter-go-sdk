package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PlatformConfigurationService service

type ResponsePlatformConfigurationCiscoDnaCenterReleaseSummary struct {
	Version  string                                                             `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformConfigurationCiscoDnaCenterReleaseSummaryResponse `json:"response,omitempty"` //
}
type ResponsePlatformConfigurationCiscoDnaCenterReleaseSummaryResponse struct {
	CorePackages           []string                                                                                   `json:"corePackages,omitempty"`           // The set of packages that are mandatory to be installed/upgraded with the release
	Packages               []string                                                                                   `json:"packages,omitempty"`               // The set of packages upgrades available with this release that will not be upgraded unless selected for upgrade
	Name                   string                                                                                     `json:"name,omitempty"`                   // Name of the release (example "dnac")
	InstalledVersion       string                                                                                     `json:"installedVersion,omitempty"`       // The installed Cisco DNAC version
	SystemVersion          string                                                                                     `json:"systemVersion,omitempty"`          // The MAGLEV-SYSTEM version
	SupportedDirectUpdates *[]ResponsePlatformConfigurationCiscoDnaCenterReleaseSummaryResponseSupportedDirectUpdates `json:"supportedDirectUpdates,omitempty"` // The list of earlier releases that can upgrade directly to the current release. If the supportedDirectUpdates value is empty, then no direct upgrades to the current release are allowed.
	TenantID               string                                                                                     `json:"tenantId,omitempty"`               // Tenant ID (for multi tenant Cisco DNA Center)
	Modified               *float64                                                                                   `json:"modified,omitempty"`               //
}
type ResponsePlatformConfigurationCiscoDnaCenterReleaseSummaryResponseSupportedDirectUpdates interface{}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummary struct {
	Version  string                                                                        `json:"version,omitempty"`  // The MAGLEV-API version (this field is for internal development purpose)
	Response *ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponse `json:"response,omitempty"` //
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponse struct {
	Nodes *[]ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodes `json:"nodes,omitempty"` //
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodes struct {
	Ntp      *ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNtp       `json:"ntp,omitempty"`      //
	Network  *[]ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetwork `json:"network,omitempty"`  //
	Proxy    *ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesProxy     `json:"proxy,omitempty"`    //
	Platform *ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesPlatform  `json:"platform,omitempty"` //
	ID       string                                                                                      `json:"id,omitempty"`       // Cluster Identifier
	Name     string                                                                                      `json:"name,omitempty"`     // Node name
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNtp struct {
	Servers []string `json:"servers,omitempty"` // NTP server
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetwork struct {
	IntraClusterLink *bool                                                                                          `json:"intra_cluster_link,omitempty"` // Flag to indicate which interface is configured as the inter-cluster link
	LacpMode         *bool                                                                                          `json:"lacp_mode,omitempty"`          // LACP Mode configuration on NIC
	Inet             *ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet  `json:"inet,omitempty"`               //
	Interface        string                                                                                         `json:"interface,omitempty"`          // Interface name
	Inet6            *ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet6 `json:"inet6,omitempty"`              //
	LacpSupported    *bool                                                                                          `json:"lacp_supported,omitempty"`     // LACP Support configuration on NIC
	SLAve            []string                                                                                       `json:"slave,omitempty"`              // Physical interface name
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet struct {
	Routes     *[]ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetRoutes     `json:"routes,omitempty"`      // Static route
	Gateway    string                                                                                                    `json:"gateway,omitempty"`     // Default gateway
	DNSServers *[]ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetDNSServers `json:"dns_servers,omitempty"` // DNS server
	Netmask    string                                                                                                    `json:"netmask,omitempty"`     // Subnet mask
	HostIP     string                                                                                                    `json:"host_ip,omitempty"`     // IP assigned
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetRoutes interface{}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInetDNSServers interface{}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesNetworkInet6 struct {
	HostIP  string `json:"host_ip,omitempty"` // IP assigned to the host
	Netmask string `json:"netmask,omitempty"` // Subnet mask of the host
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesProxy struct {
	HTTPSProxy         string   `json:"https_proxy,omitempty"`          // Https Proxy Server
	NoProxy            []string `json:"no_proxy,omitempty"`             // Servers configured to explicitly use no proxy
	HTTPSProxyUsername string   `json:"https_proxy_username,omitempty"` // Configured Https proxy username
	HTTPProxy          string   `json:"http_proxy,omitempty"`           // Not Supported
	HTTPSProxyPassword string   `json:"https_proxy_password,omitempty"` // Configured Https excrypted proxy password.
}
type ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummaryResponseNodesPlatform struct {
	Vendor  string `json:"vendor,omitempty"`  // Product manufacturer
	Product string `json:"product,omitempty"` // Product Identifier
	Serial  string `json:"serial,omitempty"`  // Serial number of chassis
}

//CiscoDnaCenterReleaseSummary Cisco DNA Center Release Summary - 5b87-e929-418b-8550
/* Provides information such as API version, mandatory core packages for installation or upgrade, optional packages, Cisco DNA Center name and version, supported direct updates, and tenant ID.


 */
func (s *PlatformConfigurationService) CiscoDnaCenterReleaseSummary() (*ResponsePlatformConfigurationCiscoDnaCenterReleaseSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/dnac-release"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformConfigurationCiscoDnaCenterReleaseSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterReleaseSummary")
	}

	result := response.Result().(*ResponsePlatformConfigurationCiscoDnaCenterReleaseSummary)
	return result, response, err

}

//CiscoDnaCenterNodesConfigurationSummary Cisco DNA Center Nodes Configuration Summary - d8b0-fb13-4f08-a967
/* Provides details about the current Cisco DNA Center node configuration, such as API version, node name, NTP server, intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and interface information.


 */
func (s *PlatformConfigurationService) CiscoDnaCenterNodesConfigurationSummary() (*ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/nodes-config"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation CiscoDnaCenterNodesConfigurationSummary")
	}

	result := response.Result().(*ResponsePlatformConfigurationCiscoDnaCenterNodesConfigurationSummary)
	return result, response, err

}
