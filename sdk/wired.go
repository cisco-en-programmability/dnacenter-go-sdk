package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type WiredService service

type GetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceQueryParams struct {
	Feature string `url:"feature,omitempty"` //Name of the feature to configure. The API /data/intent/api/wired/networkDevices/{id}/configFeatures/supported/layer2 can be used to get the list of features supported on a device.
}

type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice struct {
	Response *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                         `json:"version,omitempty"`  // Version
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponse struct {
	Features *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeatures `json:"features,omitempty"` //
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeatures struct {
	CdpGlobalConfig             *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesCdpGlobalConfig             `json:"cdpGlobalConfig,omitempty"`             //
	CdpInterfaceConfig          *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesCdpInterfaceConfig          `json:"cdpInterfaceConfig,omitempty"`          //
	DhcpSnoopingGlobalConfig    *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDhcpSnoopingGlobalConfig    `json:"dhcpSnoopingGlobalConfig,omitempty"`    //
	DhcpSnoopingInterfaceConfig *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDhcpSnoopingInterfaceConfig `json:"dhcpSnoopingInterfaceConfig,omitempty"` //
	Dot1XGlobalConfig           *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDot1XGlobalConfig           `json:"dot1xGlobalConfig,omitempty"`           //
	Dot1XInterfaceConfig        *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDot1XInterfaceConfig        `json:"dot1xInterfaceConfig,omitempty"`        //
	IgmpSnoopingGlobalConfig    *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesIgmpSnoopingGlobalConfig    `json:"igmpSnoopingGlobalConfig,omitempty"`    //
	LldpGlobalConfig            *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesLldpGlobalConfig            `json:"lldpGlobalConfig,omitempty"`            //
	LldpInterfaceConfig         *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesLldpInterfaceConfig         `json:"lldpInterfaceConfig,omitempty"`         //
	MabInterfaceConfig          *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesMabInterfaceConfig          `json:"mabInterfaceConfig,omitempty"`          //
	MldSnoopingGlobalConfig     *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesMldSnoopingGlobalConfig     `json:"mldSnoopingGlobalConfig,omitempty"`     //
	PortchannelConfig           *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesPortchannelConfig           `json:"portchannelConfig,omitempty"`           //
	StpGlobalConfig             *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesStpGlobalConfig             `json:"stpGlobalConfig,omitempty"`             //
	StpInterfaceConfig          *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesStpInterfaceConfig          `json:"stpInterfaceConfig,omitempty"`          //
	SwitchportInterfaceConfig   *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesSwitchportInterfaceConfig   `json:"switchportInterfaceConfig,omitempty"`   //
	TrunkInterfaceConfig        *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesTrunkInterfaceConfig        `json:"trunkInterfaceConfig,omitempty"`        //
	VLANConfig                  *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesVLANConfig                  `json:"vlanConfig,omitempty"`                  //
	VtpGlobalConfig             *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesVtpGlobalConfig             `json:"vtpGlobalConfig,omitempty"`             //
	VtpInterfaceConfig          *ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesVtpInterfaceConfig          `json:"vtpInterfaceConfig,omitempty"`          //
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesCdpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesCdpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDhcpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDhcpSnoopingInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDot1XGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesDot1XInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesIgmpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesLldpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesLldpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesMabInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesMldSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesPortchannelConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesStpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesStpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesSwitchportInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesTrunkInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesVLANConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesVtpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponseFeaturesVtpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` // Items
}
type ResponseWiredDeployTheConfigurationModelOnTheNetworkDevice struct {
	Response *ResponseWiredDeployTheConfigurationModelOnTheNetworkDeviceResponse `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  // Version
}
type ResponseWiredDeployTheConfigurationModelOnTheNetworkDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredGetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseWiredCreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice struct {
	Response *ResponseWiredCreateConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseWiredCreateConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredUpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice struct {
	Response *ResponseWiredUpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                            `json:"version,omitempty"`  // Version
}
type ResponseWiredUpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	CdpGlobalConfig             *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpGlobalConfig             `json:"cdpGlobalConfig,omitempty"`             //
	CdpInterfaceConfig          *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpInterfaceConfig          `json:"cdpInterfaceConfig,omitempty"`          //
	DhcpSnoopingInterfaceConfig *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingInterfaceConfig `json:"dhcpSnoopingInterfaceConfig,omitempty"` //
	DhcpSnoopingGlobalConfig    *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingGlobalConfig    `json:"dhcpSnoopingGlobalConfig,omitempty"`    //
	Dot1XInterfaceConfig        *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XInterfaceConfig        `json:"dot1xInterfaceConfig,omitempty"`        //
	Dot1XGlobalConfig           *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XGlobalConfig           `json:"dot1xGlobalConfig,omitempty"`           //
	LldpGlobalConfig            *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpGlobalConfig            `json:"lldpGlobalConfig,omitempty"`            //
	LldpInterfaceConfig         *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpInterfaceConfig         `json:"lldpInterfaceConfig,omitempty"`         //
	MabInterfaceConfig          *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMabInterfaceConfig          `json:"mabInterfaceConfig,omitempty"`          //
	MldSnoopingGlobalConfig     *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMldSnoopingGlobalConfig     `json:"mldSnoopingGlobalConfig,omitempty"`     //
	IgmpSnoopingGlobalConfig    *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceIgmpSnoopingGlobalConfig    `json:"igmpSnoopingGlobalConfig,omitempty"`    //
	StpGlobalConfig             *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpGlobalConfig             `json:"stpGlobalConfig,omitempty"`             //
	StpInterfaceConfig          *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfig          `json:"stpInterfaceConfig,omitempty"`          //
	TrunkInterfaceConfig        *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceTrunkInterfaceConfig        `json:"trunkInterfaceConfig,omitempty"`        //
	VtpGlobalConfig             *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpGlobalConfig             `json:"vtpGlobalConfig,omitempty"`             //
	VtpInterfaceConfig          *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpInterfaceConfig          `json:"vtpInterfaceConfig,omitempty"`          //
	VLANConfig                  *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVLANConfig                  `json:"vlanConfig,omitempty"`                  //
	PortChannelConfig           *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevicePortChannelConfig           `json:"portChannelConfig,omitempty"`           //
	SwitchportInterfaceConfig   *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceSwitchportInterfaceConfig   `json:"switchportInterfaceConfig,omitempty"`   //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMabInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMldSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceIgmpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfig struct {
	Items *[]ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItems `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItems struct {
	ConfigType               string                                                                                                                 `json:"configType,omitempty"`               // Type of network functionality under a feature. Config type STP_INTERFACE is for configuring STP on an interface.
	InterfaceName            string                                                                                                                 `json:"interfaceName,omitempty"`            // Interface name. The API /intent/api/v1/interface/network-device/{deviceId} can be used to get the list of interfaces on a device.example: GigabitEthernet1/0/1
	GuardMode                string                                                                                                                 `json:"guardMode,omitempty"`                // Enable loop guard or root guard on the interface. Corresponding CLI - spanning-tree guard [loop | none | root]. LOOP - Enables loop guard on the interface. ROOT - Enables root guard on the interface
	BpduFilter               string                                                                                                                 `json:"bpduFilter,omitempty"`               // Configure BPDU filtering on a port. Enabling BPDU filtering on PortFast edge-enabled port keeps the port from sending or receiving BPDUs. Corresponding CLI - spanning-tree bpduguard [enable | disable].
	BpduGuard                string                                                                                                                 `json:"bpduGuard,omitempty"`                // Configures BPDU guard. When BPDU filtering is enabled on PortFast edge-enabled port, spanning tree shuts down the port if any BPDU is received on it.
	PathCost                 *int                                                                                                                   `json:"pathCost,omitempty"`                 // Configure the cost for an interface. Corresponding CLI - spanning-tree cost <1-200000000>.
	PortFastMode             string                                                                                                                 `json:"portFastMode,omitempty"`             // Configure the portFast mode for an interface. Corresponding CLI - spanning-tree portfast [disable | trunk | network | edge | edge trunk]. DISABLE - Disable portFast for this interface. EDGE - Enable edge behavior on a Layer 2 access port connected to an end workstation or server. EDGE_TRUNK - Enable edge behavior on a trunk port. Use this keyword if the link is a trunk. NETWORK - Configure the port as a network port. Ports that are connected to Layer 2 switches and bridges can be configured as network ports. If Bridge Assurance is enabled globally, it automatically gets enabled on a spanning tree network port. TRUNK - Enable portfast on a trunk port.
	Priority                 *int                                                                                                                   `json:"priority,omitempty"`                 // Configures port priority for an interface. If a loop occurs, spanning tree uses port priority when selecting an interface to put into the forwarding state. Assign higher priority values (lower numerical values) to interfaces that you want selected first and lower priority values (higher numerical values) that you want selected last. If all interfaces have the same priority value, spanning tree puts the interface with the lowest interface number in the forwarding state and blocks the other interfaces. Corresponding CLI - spanning-tree port-priority <0-240>.multipleOf: 16
	PortVLANCostSettings     *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettings     `json:"portVlanCostSettings,omitempty"`     //
	PortVLANPrioritySettings *ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettings `json:"portVlanPrioritySettings,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettings struct {
	ConfigType string                                                                                                                    `json:"configType,omitempty"` // Type of STP Cost Settings
	Items      *[]ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettingsItems `json:"items,omitempty"`      //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettingsItems struct {
	ConfigType string `json:"configType,omitempty"` // Type of network functionality under a feature. Config type STP_INTERFACE_VLAN_COST is for configuring per VLAN cost on an interface.
	Cost       *int   `json:"cost,omitempty"`       // Configure the cost for the VLANs. Corresponding CLI - spanning-tree vlan cost <1-200000000>.
	VLANs      string `json:"vlans,omitempty"`      // VLANs can be comma separated and/or a range like '2,4-5,7,10-20'.
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettings struct {
	ConfigType string                                                                                                                        `json:"configType,omitempty"` // Type of STP Priority Settings
	Items      *[]ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettingsItems `json:"items,omitempty"`      //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettingsItems struct {
	ConfigType string `json:"configType,omitempty"` // configType
	Priority   *int   `json:"priority,omitempty"`   // Configures the port priority for the VLANs. Corresponding CLI - spanning-tree vlan port-priority <0-240>.multipleOf: 16
	VLANs      string `json:"vlans,omitempty"`      // VLANs can be comma separated and/or a range like '2,4-5,7,10-20'.
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceTrunkInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVLANConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevicePortChannelConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceSwitchportInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type ResponseWiredDeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	Response *ResponseWiredDeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseWiredDeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	Response *ResponseWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	Response *ResponseWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                             `json:"version,omitempty"`  // Version
}
type ResponseWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredGetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	Response *int   `json:"response,omitempty"` // Response
	Version  string `json:"version,omitempty"`  // Version
}
type ResponseWiredGetTheSupportedLayer2FeaturesOnAWiredDevice struct {
	Response *[]ResponseWiredGetTheSupportedLayer2FeaturesOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                              `json:"version,omitempty"`  // Version
}
type ResponseWiredGetTheSupportedLayer2FeaturesOnAWiredDeviceResponse struct {
	Name string `json:"name,omitempty"` // Name
}
type ResponseWiredCreateAConfigurationModelForTheIntendedConfigsForAWiredDevice struct {
	Response *ResponseWiredCreateAConfigurationModelForTheIntendedConfigsForAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                              `json:"version,omitempty"`  // Version
}
type ResponseWiredCreateAConfigurationModelForTheIntendedConfigsForAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredDeleteTheConfigurationModel struct {
	Response *ResponseWiredDeleteTheConfigurationModelResponse `json:"response,omitempty"` //
	Version  string                                            `json:"version,omitempty"`  // Version
}
type ResponseWiredDeleteTheConfigurationModelResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredGetsTheDeviceConfigForTheConfigurationModel struct {
	Response *ResponseWiredGetsTheDeviceConfigForTheConfigurationModelResponse `json:"response,omitempty"` //
	Version  string                                                            `json:"version,omitempty"`  // version
}
type ResponseWiredGetsTheDeviceConfigForTheConfigurationModelResponse struct {
	NetworkDeviceID string                                                                          `json:"networkDeviceId,omitempty"` // Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.
	Status          string                                                                          `json:"status,omitempty"`          // Preview generation status: 'Not Started' - preview generation has not started 'In Progress' - preview generation is in progress 'Success' - preview generation completed and was succesful 'Failed' - preview generation completed with failure 'Warning' - preview generation completed with warning
	PreviewItems    *[]ResponseWiredGetsTheDeviceConfigForTheConfigurationModelResponsePreviewItems `json:"previewItems,omitempty"`    //
}
type ResponseWiredGetsTheDeviceConfigForTheConfigurationModelResponsePreviewItems struct {
	Name          string   `json:"name,omitempty"`          // This is a system-generated description of the preview item.
	ConfigType    string   `json:"configType,omitempty"`    // configType
	ConfigPreview string   `json:"configPreview,omitempty"` // The actual preview item. This is what will be configured on the device if the user decides to go ahead with the provisioning.example: interface FortyGigabitEthernet1/1/2 device-tracking attach-policy IPDT_POLICY
	ErrorMessages []string `json:"errorMessages,omitempty"` // Error messages during preview generation.
}
type ResponseWiredDeployTheIntendedConfigurationFeaturesOnAWiredDevice struct {
	Response *ResponseWiredDeployTheIntendedConfigurationFeaturesOnAWiredDeviceResponse `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  // Version
}
type ResponseWiredDeployTheIntendedConfigurationFeaturesOnAWiredDeviceResponse struct {
	TaskID string `json:"taskId,omitempty"` // Task Id
	URL    string `json:"url,omitempty"`    // Url
}
type ResponseWiredGetDeviceDeploymentStatusConnectivity struct {
	Response *[]ResponseWiredGetDeviceDeploymentStatusConnectivityResponse `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  // version
}
type ResponseWiredGetDeviceDeploymentStatusConnectivityResponse struct {
	ActivityID         string                                                           `json:"activityId,omitempty"`         // Activity id from intent/api/v1/activity.example: 7f422eeb-effe-4938-9371-ccf6dc2fe15e
	ConfigGroupName    string                                                           `json:"configGroupName,omitempty"`    // Name of the config group that has been provisioned. example: application.telemetry
	ConfigGroupVersion *int                                                             `json:"configGroupVersion,omitempty"` // Version of the config group that has been provisioned. The version increments everytime for a provisioning operation.
	Status             string                                                           `json:"status,omitempty"`             // Provisioning status on the device.
	NetworkDeviceID    string                                                           `json:"networkDeviceId,omitempty"`    // Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.
	CreateTime         *int                                                             `json:"createTime,omitempty"`         // Create time of the service deployment status record; as measured in Unix epoch time in milliseconds.
	LastUpdateTime     *int                                                             `json:"lastUpdateTime,omitempty"`     // Last update time of the service deployment status record; as measured in Unix epoch time in milliseconds.
	StartTime          *int                                                             `json:"startTime,omitempty"`          // Service provisioning task start time; as measured in Unix epoch time in milliseconds.
	EndTime            *int                                                             `json:"endTime,omitempty"`            // Service provisioning task end time; as measured in Unix epoch time in milliseconds.
	Error              *ResponseWiredGetDeviceDeploymentStatusConnectivityResponseError `json:"error,omitempty"`              //
}
type ResponseWiredGetDeviceDeploymentStatusConnectivityResponseError struct {
	Message string `json:"message,omitempty"` // A brief message about the error.example: Connection to device timedout
	Remedy  string `json:"remedy,omitempty"`  // A brief message to suggest remedy to the failure. example: Upgrade the IOS version
}
type ResponseWiredGetServiceDeploymentStatus struct {
	Response *[]ResponseWiredGetServiceDeploymentStatusResponse `json:"response,omitempty"` //
	Version  string                                             `json:"version,omitempty"`  // version
}
type ResponseWiredGetServiceDeploymentStatusResponse struct {
	ActivityID              string `json:"activityId,omitempty"`              // Activity id from intent/api/v1/activity.example: 7f422eeb-effe-4938-9371-ccf6dc2fe15e
	ConfigGroupName         string `json:"configGroupName,omitempty"`         // Name of the config group that has been provisioned. example: application.telemetry
	ConfigGroupVersion      *int   `json:"configGroupVersion,omitempty"`      // Version of the config group that has been provisioned. The version increments everytime for a provisioning operation.
	Status                  string `json:"status,omitempty"`                  // Provisioning status.
	ProvisioningDescription string `json:"provisioningDescription,omitempty"` // Description of the provisioning operation. example: Update device config
	CreateTime              *int   `json:"createTime,omitempty"`              // Create time of the service deployment status record; as measured in Unix epoch time in milliseconds.
	LastUpdateTime          *int   `json:"lastUpdateTime,omitempty"`          // Last update time of the service deployment status record; as measured in Unix epoch time in milliseconds.
	StartTime               *int   `json:"startTime,omitempty"`               // Service provisioning task start time; as measured in Unix epoch time in milliseconds.
	EndTime                 *int   `json:"endTime,omitempty"`                 // Service provisioning task end time; as measured in Unix epoch time in milliseconds.
	ErrorCode               string `json:"errorCode,omitempty"`               // If the provisioning operation. example: NCND80300 fails, this is the error code to indicate the cause of failure.
	FailureReason           string `json:"failureReason,omitempty"`           // If the provisioning operation fails, this is the textual description to indicate the cause of failure.
	ProvisioningCompleted   *bool  `json:"provisioningCompleted,omitempty"`   // Indicates if the service provisioning operation has completed or it is still ongoing.
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	CdpGlobalConfig             *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpGlobalConfig             `json:"cdpGlobalConfig,omitempty"`             //
	CdpInterfaceConfig          *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpInterfaceConfig          `json:"cdpInterfaceConfig,omitempty"`          //
	DhcpSnoopingInterfaceConfig *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingInterfaceConfig `json:"dhcpSnoopingInterfaceConfig,omitempty"` //
	DhcpSnoopingGlobalConfig    *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingGlobalConfig    `json:"dhcpSnoopingGlobalConfig,omitempty"`    //
	Dot1XInterfaceConfig        *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XInterfaceConfig        `json:"dot1xInterfaceConfig,omitempty"`        //
	Dot1XGlobalConfig           *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XGlobalConfig           `json:"dot1xGlobalConfig,omitempty"`           //
	LldpGlobalConfig            *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpGlobalConfig            `json:"lldpGlobalConfig,omitempty"`            //
	LldpInterfaceConfig         *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpInterfaceConfig         `json:"lldpInterfaceConfig,omitempty"`         //
	MabInterfaceConfig          *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMabInterfaceConfig          `json:"mabInterfaceConfig,omitempty"`          //
	MldSnoopingGlobalConfig     *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMldSnoopingGlobalConfig     `json:"mldSnoopingGlobalConfig,omitempty"`     //
	IgmpSnoopingGlobalConfig    *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceIgmpSnoopingGlobalConfig    `json:"igmpSnoopingGlobalConfig,omitempty"`    //
	StpGlobalConfig             *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpGlobalConfig             `json:"stpGlobalConfig,omitempty"`             //
	StpInterfaceConfig          *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfig          `json:"stpInterfaceConfig,omitempty"`          //
	TrunkInterfaceConfig        *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceTrunkInterfaceConfig        `json:"trunkInterfaceConfig,omitempty"`        //
	VtpGlobalConfig             *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpGlobalConfig             `json:"vtpGlobalConfig,omitempty"`             //
	VtpInterfaceConfig          *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpInterfaceConfig          `json:"vtpInterfaceConfig,omitempty"`          //
	VLANConfig                  *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVLANConfig                  `json:"vlanConfig,omitempty"`                  //
	PortChannelConfig           *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevicePortChannelConfig           `json:"portChannelConfig,omitempty"`           //
	SwitchportInterfaceConfig   *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceSwitchportInterfaceConfig   `json:"switchportInterfaceConfig,omitempty"`   //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMabInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMldSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceIgmpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfig struct {
	Items *[]RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItems `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItems struct {
	ConfigType               string                                                                                                                   `json:"configType,omitempty"`               // Type of network functionality under a feature. Config type STP_INTERFACE is for configuring STP on an interface.
	InterfaceName            string                                                                                                                   `json:"interfaceName,omitempty"`            // Interface name. The API /intent/api/v1/interface/network-device/{deviceId} can be used to get the list of interfaces on a device.example: GigabitEthernet1/0/1
	GuardMode                string                                                                                                                   `json:"guardMode,omitempty"`                // Enable loop guard or root guard on the interface. Corresponding CLI - spanning-tree guard [loop | none | root]. LOOP - Enables loop guard on the interface. ROOT - Enables root guard on the interface
	BpduFilter               string                                                                                                                   `json:"bpduFilter,omitempty"`               // Configure BPDU filtering on a port. Enabling BPDU filtering on PortFast edge-enabled port keeps the port from sending or receiving BPDUs. Corresponding CLI - spanning-tree bpduguard [enable | disable].
	BpduGuard                string                                                                                                                   `json:"bpduGuard,omitempty"`                // Configures BPDU guard. When BPDU filtering is enabled on PortFast edge-enabled port, spanning tree shuts down the port if any BPDU is received on it.
	PathCost                 *int                                                                                                                     `json:"pathCost,omitempty"`                 // Configure the cost for an interface. Corresponding CLI - spanning-tree cost <1-200000000>.
	PortFastMode             string                                                                                                                   `json:"portFastMode,omitempty"`             // Configure the portFast mode for an interface. Corresponding CLI - spanning-tree portfast [disable | trunk | network | edge | edge trunk]. DISABLE - Disable portFast for this interface. EDGE - Enable edge behavior on a Layer 2 access port connected to an end workstation or server. EDGE_TRUNK - Enable edge behavior on a trunk port. Use this keyword if the link is a trunk. NETWORK - Configure the port as a network port. Ports that are connected to Layer 2 switches and bridges can be configured as network ports. If Bridge Assurance is enabled globally, it automatically gets enabled on a spanning tree network port. TRUNK - Enable portfast on a trunk port.
	Priority                 *int                                                                                                                     `json:"priority,omitempty"`                 // Configures port priority for an interface. If a loop occurs, spanning tree uses port priority when selecting an interface to put into the forwarding state. Assign higher priority values (lower numerical values) to interfaces that you want selected first and lower priority values (higher numerical values) that you want selected last. If all interfaces have the same priority value, spanning tree puts the interface with the lowest interface number in the forwarding state and blocks the other interfaces. Corresponding CLI - spanning-tree port-priority <0-240>.multipleOf: 16
	PortVLANCostSettings     *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettings     `json:"portVlanCostSettings,omitempty"`     //
	PortVLANPrioritySettings *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettings `json:"portVlanPrioritySettings,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettings struct {
	ConfigType string                                                                                                                      `json:"configType,omitempty"` // Type of STP Cost Settings
	Items      *[]RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettingsItems `json:"items,omitempty"`      //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettingsItems struct {
	ConfigType string `json:"configType,omitempty"` // Type of network functionality under a feature. Config type STP_INTERFACE_VLAN_COST is for configuring per VLAN cost on an interface.
	Cost       *int   `json:"cost,omitempty"`       // Configure the cost for the VLANs. Corresponding CLI - spanning-tree vlan cost <1-200000000>.
	VLANs      string `json:"vlans,omitempty"`      // VLANs can be comma separated and/or a range like '2,4-5,7,10-20'.
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettings struct {
	ConfigType string                                                                                                                          `json:"configType,omitempty"` // Type of STP Priority Settings
	Items      *[]RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettingsItems `json:"items,omitempty"`      //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettingsItems struct {
	ConfigType string `json:"configType,omitempty"` // configType
	Priority   *int   `json:"priority,omitempty"`   // Configures the port priority for the VLANs. Corresponding CLI - spanning-tree vlan port-priority <0-240>.multipleOf: 16
	VLANs      string `json:"vlans,omitempty"`      // VLANs can be comma separated and/or a range like '2,4-5,7,10-20'.
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceTrunkInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVLANConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevicePortChannelConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceSwitchportInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice struct {
	CdpGlobalConfig             *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpGlobalConfig             `json:"cdpGlobalConfig,omitempty"`             //
	CdpInterfaceConfig          *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpInterfaceConfig          `json:"cdpInterfaceConfig,omitempty"`          //
	DhcpSnoopingInterfaceConfig *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingInterfaceConfig `json:"dhcpSnoopingInterfaceConfig,omitempty"` //
	DhcpSnoopingGlobalConfig    *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingGlobalConfig    `json:"dhcpSnoopingGlobalConfig,omitempty"`    //
	Dot1XInterfaceConfig        *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XInterfaceConfig        `json:"dot1xInterfaceConfig,omitempty"`        //
	Dot1XGlobalConfig           *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XGlobalConfig           `json:"dot1xGlobalConfig,omitempty"`           //
	LldpGlobalConfig            *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpGlobalConfig            `json:"lldpGlobalConfig,omitempty"`            //
	LldpInterfaceConfig         *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpInterfaceConfig         `json:"lldpInterfaceConfig,omitempty"`         //
	MabInterfaceConfig          *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMabInterfaceConfig          `json:"mabInterfaceConfig,omitempty"`          //
	MldSnoopingGlobalConfig     *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMldSnoopingGlobalConfig     `json:"mldSnoopingGlobalConfig,omitempty"`     //
	IgmpSnoopingGlobalConfig    *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceIgmpSnoopingGlobalConfig    `json:"igmpSnoopingGlobalConfig,omitempty"`    //
	StpGlobalConfig             *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpGlobalConfig             `json:"stpGlobalConfig,omitempty"`             //
	StpInterfaceConfig          *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfig          `json:"stpInterfaceConfig,omitempty"`          //
	TrunkInterfaceConfig        *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceTrunkInterfaceConfig        `json:"trunkInterfaceConfig,omitempty"`        //
	VtpGlobalConfig             *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpGlobalConfig             `json:"vtpGlobalConfig,omitempty"`             //
	VtpInterfaceConfig          *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpInterfaceConfig          `json:"vtpInterfaceConfig,omitempty"`          //
	VLANConfig                  *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVLANConfig                  `json:"vlanConfig,omitempty"`                  //
	PortChannelConfig           *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevicePortChannelConfig           `json:"portChannelConfig,omitempty"`           //
	SwitchportInterfaceConfig   *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceSwitchportInterfaceConfig   `json:"switchportInterfaceConfig,omitempty"`   //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceCdpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDhcpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceDot1XGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceLldpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMabInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceMldSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceIgmpSnoopingGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfig struct {
	Items *[]RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItems `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItems struct {
	ConfigType               string                                                                                                                   `json:"configType,omitempty"`               // Type of network functionality under a feature. Config type STP_INTERFACE is for configuring STP on an interface.
	InterfaceName            string                                                                                                                   `json:"interfaceName,omitempty"`            // Interface name. The API /intent/api/v1/interface/network-device/{deviceId} can be used to get the list of interfaces on a device.example: GigabitEthernet1/0/1
	GuardMode                string                                                                                                                   `json:"guardMode,omitempty"`                // Enable loop guard or root guard on the interface. Corresponding CLI - spanning-tree guard [loop | none | root]. LOOP - Enables loop guard on the interface. ROOT - Enables root guard on the interface
	BpduFilter               string                                                                                                                   `json:"bpduFilter,omitempty"`               // Configure BPDU filtering on a port. Enabling BPDU filtering on PortFast edge-enabled port keeps the port from sending or receiving BPDUs. Corresponding CLI - spanning-tree bpduguard [enable | disable].
	BpduGuard                string                                                                                                                   `json:"bpduGuard,omitempty"`                // Configures BPDU guard. When BPDU filtering is enabled on PortFast edge-enabled port, spanning tree shuts down the port if any BPDU is received on it.
	PathCost                 *int                                                                                                                     `json:"pathCost,omitempty"`                 // Configure the cost for an interface. Corresponding CLI - spanning-tree cost <1-200000000>.
	PortFastMode             string                                                                                                                   `json:"portFastMode,omitempty"`             // Configure the portFast mode for an interface. Corresponding CLI - spanning-tree portfast [disable | trunk | network | edge | edge trunk]. DISABLE - Disable portFast for this interface. EDGE - Enable edge behavior on a Layer 2 access port connected to an end workstation or server. EDGE_TRUNK - Enable edge behavior on a trunk port. Use this keyword if the link is a trunk. NETWORK - Configure the port as a network port. Ports that are connected to Layer 2 switches and bridges can be configured as network ports. If Bridge Assurance is enabled globally, it automatically gets enabled on a spanning tree network port. TRUNK - Enable portfast on a trunk port.
	Priority                 *int                                                                                                                     `json:"priority,omitempty"`                 // Configures port priority for an interface. If a loop occurs, spanning tree uses port priority when selecting an interface to put into the forwarding state. Assign higher priority values (lower numerical values) to interfaces that you want selected first and lower priority values (higher numerical values) that you want selected last. If all interfaces have the same priority value, spanning tree puts the interface with the lowest interface number in the forwarding state and blocks the other interfaces. Corresponding CLI - spanning-tree port-priority <0-240>.multipleOf: 16
	PortVLANCostSettings     *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettings     `json:"portVlanCostSettings,omitempty"`     //
	PortVLANPrioritySettings *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettings `json:"portVlanPrioritySettings,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettings struct {
	ConfigType string                                                                                                                      `json:"configType,omitempty"` // Type of STP Cost Settings
	Items      *[]RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettingsItems `json:"items,omitempty"`      //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANCostSettingsItems struct {
	ConfigType string `json:"configType,omitempty"` // Type of network functionality under a feature. Config type STP_INTERFACE_VLAN_COST is for configuring per VLAN cost on an interface.
	Cost       *int   `json:"cost,omitempty"`       // Configure the cost for the VLANs. Corresponding CLI - spanning-tree vlan cost <1-200000000>.
	VLANs      string `json:"vlans,omitempty"`      // VLANs can be comma separated and/or a range like '2,4-5,7,10-20'.
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettings struct {
	ConfigType string                                                                                                                          `json:"configType,omitempty"` // Type of STP Priority Settings
	Items      *[]RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettingsItems `json:"items,omitempty"`      //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceStpInterfaceConfigItemsPortVLANPrioritySettingsItems struct {
	ConfigType string `json:"configType,omitempty"` // configType
	Priority   *int   `json:"priority,omitempty"`   // Configures the port priority for the VLANs. Corresponding CLI - spanning-tree vlan port-priority <0-240>.multipleOf: 16
	VLANs      string `json:"vlans,omitempty"`      // VLANs can be comma separated and/or a range like '2,4-5,7,10-20'.
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceTrunkInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpGlobalConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVtpInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceVLANConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevicePortChannelConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}
type RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDeviceSwitchportInterfaceConfig struct {
	Items *[][]string `json:"items,omitempty"` //
}

//GetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice Get configurations for intended layer 2 features on a wired device. - f1a7-b844-4cba-833c
/* This API returns the configurations for the intended layer 2 features on a wired device. Even after the intended configurations are deployed using the API /intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy, they continue to be a part of the intended features on the device.


@param id id path parameter. Network device ID of the wired device to configure.

@param GetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configurations-for-intended-layer2-features-on-a-wired-device
*/
func (s *WiredService) GetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice(id string, GetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceQueryParams *GetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceQueryParams) (*ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	queryString, _ := query.Values(GetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice(id, GetConfigurationsForIntendedLayer2FeaturesOnAWiredDeviceQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredGetConfigurationsForIntendedLayer2FeaturesOnAWiredDevice)
	return result, response, err

}

//GetConfigurationsForADeployedLayer2FeatureOnAWiredDevice Get configurations for a deployed layer 2 feature on a wired device. - 42a6-794b-4dc8-87ed
/* The API returns configurations for a deployed layer 2 feature on a wired device.


@param id id path parameter. Network device ID of the wired device to retrieve configuration for.

@param feature feature path parameter. Name of the feature to retrieve Layer 2 configuration for. The API /dna/intent/api/v1/networkDevices/{id}/configFeatures/supported/layer2 can be used to get the list of features supported on a device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configurations-for-a-deployed-layer2-feature-on-a-wired-device
*/
func (s *WiredService) GetConfigurationsForADeployedLayer2FeatureOnAWiredDevice(id string, feature string) (*resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/deployed/layer2/{feature}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationsForADeployedLayer2FeatureOnAWiredDevice(id, feature)
		}
		return response, fmt.Errorf("error with operation GetConfigurationsForADeployedLayer2FeatureOnAWiredDevice")
	}

	return response, err

}

//GetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice Get number of configurations for a deployed layer 2 feature on a wired device. - 4297-4ba7-43b9-bca8
/* The API returns the number of configurations for a deployed layer 2 feature on a wired device.


@param id id path parameter. Network device ID of the wired device to retrieve configuration for.

@param feature feature path parameter. Name of the feature to retrieve configuration for. The API /dna/intent/api/v1/networkDevices/{id}/configFeatures/supported/layer2 can be used to get the list of features supported on a device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-number-of-configurations-for-a-deployed-layer2-feature-on-a-wired-device
*/
func (s *WiredService) GetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice(id string, feature string) (*ResponseWiredGetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/deployed/layer2/{feature}/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice(id, feature)
		}
		return nil, response, fmt.Errorf("error with operation GetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredGetNumberOfConfigurationsForADeployedLayer2FeatureOnAWiredDevice)
	return result, response, err

}

//GetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice Get configurations for an intended layer 2 feature on a wired device. - 8aac-8a1f-421a-9228
/* This API returns the configurations for an intended layer 2 feature on a wired device. Even after the intended configurations are deployed using the API /dna/intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy, they continue to be a part of the intended features on the device.


@param id id path parameter. Network device ID of the wired device to configure.

@param feature feature path parameter. The name of the feature to be retrieved.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-configurations-for-an-intended-layer2-feature-on-a-wired-device
*/
func (s *WiredService) GetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id string, feature string) (*ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2/{feature}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id, feature)
		}
		return nil, response, fmt.Errorf("error with operation GetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredGetConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
	return result, response, err

}

//GetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice Get number of configurations for an intended layer 2 feature on a wired device. - a3ac-6a96-4eea-92e3
/* This API returns the count of the instances of the configurations for an intended layer 2 feature on a wired device.


@param id id path parameter. Network device ID of the wired device to configure.

@param feature feature path parameter. Name of the feature to configure. The API /dna/intent/api/v1/networkDevices/{id}/configFeatures/supported/layer2 can be used to get the list of features supported on a device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-number-of-configurations-for-an-intended-layer2-feature-on-a-wired-device
*/
func (s *WiredService) GetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id string, feature string) (*ResponseWiredGetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2/{feature}/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id, feature)
		}
		return nil, response, fmt.Errorf("error with operation GetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredGetNumberOfConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
	return result, response, err

}

//GetTheSupportedLayer2FeaturesOnAWiredDevice Get the supported layer 2 features on a wired device. - 8580-a81e-441a-a4de
/* The API returns the supported layer 2 features on a wired device.


@param id id path parameter. Network device ID of the wired device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-the-supported-layer2-features-on-a-wired-device
*/
func (s *WiredService) GetTheSupportedLayer2FeaturesOnAWiredDevice(id string) (*ResponseWiredGetTheSupportedLayer2FeaturesOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/supported/layer2"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetTheSupportedLayer2FeaturesOnAWiredDevice{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetTheSupportedLayer2FeaturesOnAWiredDevice(id)
		}
		return nil, response, fmt.Errorf("error with operation GetTheSupportedLayer2FeaturesOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredGetTheSupportedLayer2FeaturesOnAWiredDevice)
	return result, response, err

}

//GetsTheDeviceConfigForTheConfigurationModel Gets the device config for the configuration model. - bfbe-ebc3-4668-8962
/* Gets the device config for the configuration model. This API is 'Step 3' in the following workflow. Step 1Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels' to start the provision of intended features. The response has a taskId which is the previewActivityId in all subsequent APIs. The task must be successfully complete before proceeding to the next step. It is not recommended to proceed when there is any task failure in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 2Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to generate device CLIs for preview. The response has a task ID. The task must be successfully complete before using the GET API to view CLIs. It is not recommended to proceed when there is any task failure(s) in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 3Use 'GET /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to view the CLIs that will be applied to the device. Step 4Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/deploy' to deploy the intent to the device.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.

@param previewActivityID previewActivityId path parameter. Activity id is the taskId from Step 2- 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config


Documentation Link: https://developer.cisco.com/docs/dna-center/#!gets-the-device-config-for-the-configuration-model
*/
func (s *WiredService) GetsTheDeviceConfigForTheConfigurationModel(networkDeviceID string, previewActivityID string) (*ResponseWiredGetsTheDeviceConfigForTheConfigurationModel, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/config"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetsTheDeviceConfigForTheConfigurationModel{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetsTheDeviceConfigForTheConfigurationModel(networkDeviceID, previewActivityID)
		}
		return nil, response, fmt.Errorf("error with operation GetsTheDeviceConfigForTheConfigurationModel")
	}

	result := response.Result().(*ResponseWiredGetsTheDeviceConfigForTheConfigurationModel)
	return result, response, err

}

//GetDeviceDeploymentStatusConnectivity Get device deployment status - f2a8-a8b7-4c99-a0f9
/* The API returns device deployment status based on filter criteria.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-deployment-status-connectivity
*/
func (s *WiredService) GetDeviceDeploymentStatusConnectivity(networkDeviceID string) (*ResponseWiredGetDeviceDeploymentStatusConnectivity, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/deviceDeployments"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetDeviceDeploymentStatusConnectivity{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceDeploymentStatusConnectivity(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetDeviceDeploymentStatusConnectivity")
	}

	result := response.Result().(*ResponseWiredGetDeviceDeploymentStatusConnectivity)
	return result, response, err

}

//GetServiceDeploymentStatus Get service deployment status. - 82b5-08af-493a-9846
/* Returns service deployment status based on filter criteria.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-service-deployment-status
*/
func (s *WiredService) GetServiceDeploymentStatus(networkDeviceID string) (*ResponseWiredGetServiceDeploymentStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/serviceDeployments"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredGetServiceDeploymentStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetServiceDeploymentStatus(networkDeviceID)
		}
		return nil, response, fmt.Errorf("error with operation GetServiceDeploymentStatus")
	}

	result := response.Result().(*ResponseWiredGetServiceDeploymentStatus)
	return result, response, err

}

//DeployTheConfigurationModelOnTheNetworkDevice Deploy the configuration model on the network device. - 8e98-f9f2-4868-beab
/* This API deploys the configuration model on the network device. This is the final step 'Step 4' of the following workflow. Step 1Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels' to start the provision of intended features. The response has a taskId which is the previewActivityId in all subsequent APIs. The task must be successfully complete before proceeding to the next step. It is not recommended to proceed when there is any task failure in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 2Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to generate device CLIs for preview. The response has a task ID. The task must be successfully complete before using the GET API to view CLIs. It is not recommended to proceed when there is any task failure(s) in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 3Use 'GET /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to view the CLIs that will be applied to the device. Step 4Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/deploy' to push the intent to the device.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.

@param previewActivityID previewActivityId path parameter. Activity id from intent/api/v1/activity.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-the-configuration-model-on-the-network-device
*/
func (s *WiredService) DeployTheConfigurationModelOnTheNetworkDevice(networkDeviceID string, previewActivityID string) (*ResponseWiredDeployTheConfigurationModelOnTheNetworkDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/deploy"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredDeployTheConfigurationModelOnTheNetworkDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployTheConfigurationModelOnTheNetworkDevice(networkDeviceID, previewActivityID)
		}

		return nil, response, fmt.Errorf("error with operation DeployTheConfigurationModelOnTheNetworkDevice")
	}

	result := response.Result().(*ResponseWiredDeployTheConfigurationModelOnTheNetworkDevice)
	return result, response, err

}

//CreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice Create configurations for intended layer 2 features on a wired device - ab82-7920-485b-944e
/* This API creates configurations for the intended features on a wired device, if none have been added earlier. Only the feature configurations to be changed need to be added to the intended features. When the intended features are deployed to a device using the API /intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy, they are applied on top of the existing configurations on the device. Any existing configurations on the device which are not included in the intended features, are retained on the device.


@param id id path parameter. Network device ID of the wired device to configure.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-configurations-for-intended-layer2-features-on-a-wired-device
*/
func (s *WiredService) CreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice(id string) (*ResponseWiredCreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredCreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice(id)
		}

		return nil, response, fmt.Errorf("error with operation CreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredCreateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice)
	return result, response, err

}

//CreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice Create configurations for an intended layer 2 feature on a wired device - a79a-8a2e-4598-8219
/* This API creates configurations for an intended feature on a wired device. Once all the updates to intended features are complete, they can be deployed to a device using the API /dna/intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy. When the intended features are deployed, they are applied on top of the existing configurations on the device. Any existing configurations on the device which are not included in the intended features, are retained on the device.


@param id id path parameter. Network device ID of the wired device to configure.

@param feature feature path parameter. Name of the feature to configure. The API /dna/intent/api/v1/networkDevices/{id}/configFeatures/supported/layer2 can be used to get the list of features supported on a device.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-configurations-for-an-intended-layer2-feature-on-a-wired-device
*/
func (s *WiredService) CreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id string, feature string, requestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice *RequestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice) (*ResponseWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2/{feature}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice).
		SetResult(&ResponseWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id, feature, requestWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
		}

		return nil, response, fmt.Errorf("error with operation CreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredCreateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
	return result, response, err

}

//CreateAConfigurationModelForTheIntendedConfigsForAWiredDevice Create a configuration model for the intended configs for a wired device. - 1f83-7838-4c1b-9130
/* Create a configuration model for the intended configs for a wired device. This is a pre-requisite to preview the generated device config for the provisioning intent. This is mandatory if the provisioning settings require Preview or ITSM Approval before deploying configurations on network devices. The API /intent/api/v1/provisioningSettings can be used to get or update provisioning settings. This API is 'Step 1' in the following workflow Step 1Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels' to start the provision of intended features. The response has a taskId which is the previewActivityId in all subsequent APIs. The task must be successfully complete before proceeding to the next step. It is not recommended to proceed when there is any task failure in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 2Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to generate device CLIs for preview. The response has a task ID. The task must be successfully complete before using the GET API to view CLIs. It is not recommended to proceed when there is any task failure(s) in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 3Use 'GET /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to view the CLIs that will be applied to the device. Step 4Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/deploy' to deploy the intent to the device.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-a-configuration-model-for-the-intended-configs-for-a-wired-device
*/
func (s *WiredService) CreateAConfigurationModelForTheIntendedConfigsForAWiredDevice(networkDeviceID string) (*ResponseWiredCreateAConfigurationModelForTheIntendedConfigsForAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredCreateAConfigurationModelForTheIntendedConfigsForAWiredDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateAConfigurationModelForTheIntendedConfigsForAWiredDevice(networkDeviceID)
		}

		return nil, response, fmt.Errorf("error with operation CreateAConfigurationModelForTheIntendedConfigsForAWiredDevice")
	}

	result := response.Result().(*ResponseWiredCreateAConfigurationModelForTheIntendedConfigsForAWiredDevice)
	return result, response, err

}

//GenerateTheDeviceConfigForTheConfigurationModel Generate the device config for the configuration model. - 9995-b9cf-4768-b85d
/* Generates the device config for the configuration model. This API is 'Step 2' in the following workflow  Step 1Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels' to start the provision of intended features. The response has a taskId which is the previewActivityId in all subsequent APIs. The task must be successfully complete before proceeding to the next step. It is not recommended to proceed when there is any task failure in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 2Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to generate device CLIs for preview. The response has a task ID. The task must be successfully complete before using the GET API to view CLIs. It is not recommended to proceed when there is any task failure(s) in this step. The API 'DELETE /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}' can be used at any step to discard/cancel the provision of intended features. Step 3Use 'GET /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config' to view the CLIs that will be applied to the device. Step 4Use 'POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/deploy' to deploy the intent to the device.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.

@param previewActivityID previewActivityId path parameter. Activity id is taskId from Step 1- POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels'


Documentation Link: https://developer.cisco.com/docs/dna-center/#!generate-the-device-config-for-the-configuration-model
*/
func (s *WiredService) GenerateTheDeviceConfigForTheConfigurationModel(networkDeviceID string, previewActivityID string) (*resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/config"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.GenerateTheDeviceConfigForTheConfigurationModel(networkDeviceID, previewActivityID)
		}

		return response, fmt.Errorf("error with operation GenerateTheDeviceConfigForTheConfigurationModel")
	}

	return response, err

}

//DeployTheIntendedConfigurationFeaturesOnAWiredDevice Deploy the intended configuration features on a wired device. - 688a-0b32-45a8-b7a6
/* Deploy the intended configuration features on a wired device. This can be used only if the provisioning settings do not require Preview or ITSM Approval before deploying configurations on network devices. The API /intent/api/v1/provisioningSettings can be used to get or update provisioning settings.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!deploy-the-intended-configuration-features-on-a-wired-device
*/
func (s *WiredService) DeployTheIntendedConfigurationFeaturesOnAWiredDevice(networkDeviceID string) (*ResponseWiredDeployTheIntendedConfigurationFeaturesOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/deploy"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredDeployTheIntendedConfigurationFeaturesOnAWiredDevice{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeployTheIntendedConfigurationFeaturesOnAWiredDevice(networkDeviceID)
		}

		return nil, response, fmt.Errorf("error with operation DeployTheIntendedConfigurationFeaturesOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredDeployTheIntendedConfigurationFeaturesOnAWiredDevice)
	return result, response, err

}

//UpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice Update configurations for intended layer 2 features on a wired device. - 82aa-ba7e-4b6a-a34a
/* This API updates the configurations for the intended features on a wired device. Only the feature configurations to be changed need to be added to the intended features. Updates to intended features can be done over several iterations. Once the updates are complete, the intended features can be deployed to a device using the API /dna/intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy. When the intended features are deployed, they are applied on top of the existing configurations on the device. Any existing configurations on the device which are not included in the intended features, are retained on the device.


@param id id path parameter. Network device ID of the wired device to configure.

*/
func (s *WiredService) UpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice(id string) (*ResponseWiredUpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredUpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice(id)
		}
		return nil, response, fmt.Errorf("error with operation UpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredUpdateConfigurationsForIntendedLayer2FeaturesOnAWiredDevice)
	return result, response, err

}

//UpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice Update configurations for an intended layer 2 feature on a wired device. - 2dab-88c9-45e8-93fd
/* This API updates the configurations for an intended feature on a wired device. Updates to other intended features can be done over several iterations. Once all the updates to intended features are complete, they can be deployed to a device using the API /dna/intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy. When the intended features are deployed, they are applied on top of the existing configurations on the device. Any existing configurations on the device which are not included in the intended features, are retained on the device.


@param id id path parameter. Network device ID of the wired device to configure.

@param feature feature path parameter. Name of the feature to update configuration for. The feature must be already created.

*/
func (s *WiredService) UpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id string, feature string, requestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice *RequestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice) (*ResponseWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice, *resty.Response, error) {
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2/{feature}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice).
		SetResult(&ResponseWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id, feature, requestWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
		}
		return nil, response, fmt.Errorf("error with operation UpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredUpdateConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
	return result, response, err

}

//DeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice Delete configurations for an intended layer 2 feature on a wired device. - aa91-fb8a-4b3b-b2c3
/* This API deletes the configurations for an intended feature on a wired device. Once all the updates to intended features are complete, they can be deployed to a device using the API /dna/intent/api/v1/networkDevices/{id}/configFeatures/intended/deploy. When the intended features are deployed, they are applied on top of the existing configurations on the device. Any existing configurations on the device which are not included in the intended features, are retained on the device.


@param id id path parameter. Network device ID of the wired device to configure.

@param feature feature path parameter. Name of the feature to delete.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-configurations-for-an-intended-layer2-feature-on-a-wired-device
*/
func (s *WiredService) DeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(id string, feature string) (*ResponseWiredDeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice, *resty.Response, error) {
	//id string,feature string
	path := "/dna/intent/api/v1/wired/networkDevices/{id}/configFeatures/intended/layer2/{feature}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{feature}", fmt.Sprintf("%v", feature), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredDeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice(
				id, feature)
		}
		return nil, response, fmt.Errorf("error with operation DeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice")
	}

	result := response.Result().(*ResponseWiredDeleteConfigurationsForAnIntendedLayer2FeatureOnAWiredDevice)
	return result, response, err

}

//DeleteTheConfigurationModel Delete the configuration model. - d4a0-5a32-43a9-a670
/* Deletes the configuration model. The API can be used at any step to discard/cancel the provision of intended features.


@param networkDeviceID networkDeviceId path parameter. Network device ID of the wired device to provision. The API /intent/api/v1/network-device can be used to get the network device ID.

@param previewActivityID previewActivityId path parameter. Activity id from POST /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels or /intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}/networkDevices/{networkDeviceId}/config


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-the-configuration-model
*/
func (s *WiredService) DeleteTheConfigurationModel(networkDeviceID string, previewActivityID string) (*ResponseWiredDeleteTheConfigurationModel, *resty.Response, error) {
	//networkDeviceID string,previewActivityID string
	path := "/dna/intent/api/v1/wired/networkDevices/{networkDeviceId}/configFeatures/intended/configurationModels/{previewActivityId}"
	path = strings.Replace(path, "{networkDeviceId}", fmt.Sprintf("%v", networkDeviceID), -1)
	path = strings.Replace(path, "{previewActivityId}", fmt.Sprintf("%v", previewActivityID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseWiredDeleteTheConfigurationModel{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteTheConfigurationModel(
				networkDeviceID, previewActivityID)
		}
		return nil, response, fmt.Errorf("error with operation DeleteTheConfigurationModel")
	}

	result := response.Result().(*ResponseWiredDeleteTheConfigurationModel)
	return result, response, err

}
