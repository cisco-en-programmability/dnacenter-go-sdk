package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type PathTraceService service

type RetrivesAllPreviousPathtracesSummaryQueryParams struct {
	PeriodicRefresh bool   `url:"periodicRefresh,omitempty"` //Is analysis periodically refreshed?
	SourceIP        string `url:"sourceIP,omitempty"`        //Source IP address
	DestIP          string `url:"destIP,omitempty"`          //Destination IP adress
	SourcePort      string `url:"sourcePort,omitempty"`      //Source port
	DestPort        string `url:"destPort,omitempty"`        //Destination port
	GtCreateTime    string `url:"gtCreateTime,omitempty"`    //Analyses requested after this time
	LtCreateTime    string `url:"ltCreateTime,omitempty"`    //Analyses requested before this time
	Protocol        string `url:"protocol,omitempty"`        //Protocol
	Status          string `url:"status,omitempty"`          //Status
	TaskID          string `url:"taskId,omitempty"`          //Task ID
	LastUpdateTime  string `url:"lastUpdateTime,omitempty"`  //Last update time
	Limit           int    `url:"limit,omitempty"`           //Number of resources returned
	Offset          int    `url:"offset,omitempty"`          //Start index of resources returned (1-based)
	Order           string `url:"order,omitempty"`           //Order by this field
	SortBy          string `url:"sortBy,omitempty"`          //Sort by this field
}

type ResponsePathTraceRetrivesAllPreviousPathtracesSummary struct {
	Response *[]ResponsePathTraceRetrivesAllPreviousPathtracesSummaryResponse `json:"response,omitempty"` //
	Version  string                                                           `json:"version,omitempty"`  //
}
type ResponsePathTraceRetrivesAllPreviousPathtracesSummaryResponse struct {
	ControlPath     *bool    `json:"controlPath,omitempty"`     //
	CreateTime      *int     `json:"createTime,omitempty"`      //
	DestIP          string   `json:"destIP,omitempty"`          //
	DestPort        string   `json:"destPort,omitempty"`        //
	FailureReason   string   `json:"failureReason,omitempty"`   //
	ID              string   `json:"id,omitempty"`              //
	Inclusions      []string `json:"inclusions,omitempty"`      //
	LastUpdateTime  *int     `json:"lastUpdateTime,omitempty"`  //
	PeriodicRefresh *bool    `json:"periodicRefresh,omitempty"` //
	Protocol        string   `json:"protocol,omitempty"`        //
	SourceIP        string   `json:"sourceIP,omitempty"`        //
	SourcePort      string   `json:"sourcePort,omitempty"`      //
	Status          string   `json:"status,omitempty"`          //
}
type ResponsePathTraceInitiateANewPathtrace struct {
	Response *ResponsePathTraceInitiateANewPathtraceResponse `json:"response,omitempty"` //
	Version  string                                          `json:"version,omitempty"`  //
}
type ResponsePathTraceInitiateANewPathtraceResponse struct {
	FlowAnalysisID string `json:"flowAnalysisId,omitempty"` //
	TaskID         string `json:"taskId,omitempty"`         //
	URL            string `json:"url,omitempty"`            //
}
type ResponsePathTraceRetrievesPreviousPathtrace struct {
	Response *ResponsePathTraceRetrievesPreviousPathtraceResponse `json:"response,omitempty"` //
	Version  string                                               `json:"version,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponse struct {
	DetailedStatus      *ResponsePathTraceRetrievesPreviousPathtraceResponseDetailedStatus        `json:"detailedStatus,omitempty"`      //
	LastUpdate          string                                                                    `json:"lastUpdate,omitempty"`          //
	NetworkElements     *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElements     `json:"networkElements,omitempty"`     //
	NetworkElementsInfo *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfo `json:"networkElementsInfo,omitempty"` //
	Properties          []string                                                                  `json:"properties,omitempty"`          //
	Request             *ResponsePathTraceRetrievesPreviousPathtraceResponseRequest               `json:"request,omitempty"`             //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElements struct {
	AccuracyList                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsAccuracyList           `json:"accuracyList,omitempty"`                       //
	DetailedStatus                     *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDetailedStatus           `json:"detailedStatus,omitempty"`                     //
	DeviceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatistics         `json:"deviceStatistics,omitempty"`                   //
	DeviceStatsCollection              string                                                                                      `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string                                                                                      `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressPhysicalInterface            *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterface  `json:"egressPhysicalInterface,omitempty"`            //
	EgressVirtualInterface             *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterface   `json:"egressVirtualInterface,omitempty"`             //
	FlexConnect                        *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnect              `json:"flexConnect,omitempty"`                        //
	ID                                 string                                                                                      `json:"id,omitempty"`                                 //
	IngressPhysicalInterface           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterface `json:"ingressPhysicalInterface,omitempty"`           //
	IngressVirtualInterface            *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterface  `json:"ingressVirtualInterface,omitempty"`            //
	IP                                 string                                                                                      `json:"ip,omitempty"`                                 //
	LinkInformationSource              string                                                                                      `json:"linkInformationSource,omitempty"`              //
	Name                               string                                                                                      `json:"name,omitempty"`                               //
	PerfMonCollection                  string                                                                                      `json:"perfMonCollection,omitempty"`                  //
	PerfMonCollectionFailureReason     string                                                                                      `json:"perfMonCollectionFailureReason,omitempty"`     //
	PerfMonStatistics                  *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsPerfMonStatistics      `json:"perfMonStatistics,omitempty"`                  //
	Role                               string                                                                                      `json:"role,omitempty"`                               //
	SSID                               string                                                                                      `json:"ssid,omitempty"`                               //
	Tunnels                            []string                                                                                    `json:"tunnels,omitempty"`                            //
	Type                               string                                                                                      `json:"type,omitempty"`                               //
	WLANID                             string                                                                                      `json:"wlanId,omitempty"`                             //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsAccuracyList struct {
	Percent *int   `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatistics struct {
	CPUStatistics    *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsCPUStatistics    `json:"cpuStatistics,omitempty"`    //
	MemoryStatistics *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsMemoryStatistics `json:"memoryStatistics,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsCPUStatistics struct {
	FiveMinUsageInPercentage  *float64 `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage *float64 `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   *float64 `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               *int     `json:"refreshedAt,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsMemoryStatistics struct {
	MemoryUsage *int `json:"memoryUsage,omitempty"` //
	RefreshedAt *int `json:"refreshedAt,omitempty"` //
	TotalMemory *int `json:"totalMemory,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                        `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                        `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                        `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                        `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                        `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                        `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                        `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                        `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                              `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                              `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                           `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                           `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                             `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                             `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                             `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                             `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                             `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                             `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                             `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                       `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                       `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                       `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                       `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                       `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                       `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                       `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                       `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                             `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                             `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                          `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                          `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                               `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                            `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                            `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                            `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                            `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                            `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                            `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                            `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnect struct {
	Authentication            string                                                                                           `json:"authentication,omitempty"`            //
	DataSwitching             string                                                                                           `json:"dataSwitching,omitempty"`             //
	EgressACLAnalysis         *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysis  `json:"egressAclAnalysis,omitempty"`         //
	IngressACLAnalysis        *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysis `json:"ingressAclAnalysis,omitempty"`        //
	WirelessLanControllerID   string                                                                                           `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string                                                                                           `json:"wirelessLanControllerName,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysis struct {
	ACLName      string                                                                                                        `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                        `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                     `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                     `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                          `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysis struct {
	ACLName      string                                                                                                         `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                         `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                      `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                      `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                           `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                         `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                         `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                         `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                         `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                         `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                         `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                         `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                         `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                               `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                               `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                            `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                            `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                 `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                              `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                              `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                              `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                              `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                              `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                              `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                              `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                        `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                        `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                        `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                        `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                        `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                        `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                        `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                        `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                              `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                              `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                           `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                           `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                             `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                             `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                             `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                             `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                             `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                             `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                             `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsPerfMonStatistics struct {
	ByteRate             *int     `json:"byteRate,omitempty"`             //
	DestIPAddress        string   `json:"destIpAddress,omitempty"`        //
	DestPort             string   `json:"destPort,omitempty"`             //
	InputInterface       string   `json:"inputInterface,omitempty"`       //
	IPv4DSCP             string   `json:"ipv4DSCP,omitempty"`             //
	IPv4TTL              *int     `json:"ipv4TTL,omitempty"`              //
	OutputInterface      string   `json:"outputInterface,omitempty"`      //
	PacketBytes          *int     `json:"packetBytes,omitempty"`          //
	PacketCount          *int     `json:"packetCount,omitempty"`          //
	PacketLoss           *int     `json:"packetLoss,omitempty"`           //
	PacketLossPercentage *float64 `json:"packetLossPercentage,omitempty"` //
	Protocol             string   `json:"protocol,omitempty"`             //
	RefreshedAt          *int     `json:"refreshedAt,omitempty"`          //
	RtpJitterMax         *int     `json:"rtpJitterMax,omitempty"`         //
	RtpJitterMean        *int     `json:"rtpJitterMean,omitempty"`        //
	RtpJitterMin         *int     `json:"rtpJitterMin,omitempty"`         //
	SourceIPAddress      string   `json:"sourceIpAddress,omitempty"`      //
	SourcePort           string   `json:"sourcePort,omitempty"`           //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfo struct {
	AccuracyList                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoAccuracyList          `json:"accuracyList,omitempty"`                       //
	DetailedStatus                     *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDetailedStatus          `json:"detailedStatus,omitempty"`                     //
	DeviceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatistics        `json:"deviceStatistics,omitempty"`                   //
	DeviceStatsCollection              string                                                                                         `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string                                                                                         `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressInterface                    *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterface         `json:"egressInterface,omitempty"`                    //
	FlexConnect                        *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnect             `json:"flexConnect,omitempty"`                        //
	ID                                 string                                                                                         `json:"id,omitempty"`                                 //
	IngressInterface                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterface        `json:"ingressInterface,omitempty"`                   //
	IP                                 string                                                                                         `json:"ip,omitempty"`                                 //
	LinkInformationSource              string                                                                                         `json:"linkInformationSource,omitempty"`              //
	Name                               string                                                                                         `json:"name,omitempty"`                               //
	PerfMonCollection                  string                                                                                         `json:"perfMonCollection,omitempty"`                  //
	PerfMonCollectionFailureReason     string                                                                                         `json:"perfMonCollectionFailureReason,omitempty"`     //
	PerfMonitorStatistics              *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoPerfMonitorStatistics `json:"perfMonitorStatistics,omitempty"`              //
	Role                               string                                                                                         `json:"role,omitempty"`                               //
	SSID                               string                                                                                         `json:"ssid,omitempty"`                               //
	Tunnels                            []string                                                                                       `json:"tunnels,omitempty"`                            //
	Type                               string                                                                                         `json:"type,omitempty"`                               //
	WLANID                             string                                                                                         `json:"wlanId,omitempty"`                             //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoAccuracyList struct {
	Percent *int   `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatistics struct {
	CPUStatistics    *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsCPUStatistics    `json:"cpuStatistics,omitempty"`    //
	MemoryStatistics *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics `json:"memoryStatistics,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsCPUStatistics struct {
	FiveMinUsageInPercentage  *float64 `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage *float64 `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   *float64 `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               *int     `json:"refreshedAt,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics struct {
	MemoryUsage *int `json:"memoryUsage,omitempty"` //
	RefreshedAt *int `json:"refreshedAt,omitempty"` //
	TotalMemory *int `json:"totalMemory,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterface struct {
	PhysicalInterface *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterface  `json:"physicalInterface,omitempty"` //
	VirtualInterface  *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterface `json:"virtualInterface,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                     `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                     `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                     `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                     `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                     `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                     `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                     `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                     `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                           `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                           `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                        `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                        `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                             `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                          `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                          `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                          `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                          `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                          `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                          `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                          `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                    `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                    `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                    `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                    `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                    `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                    `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                    `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                    `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                          `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                          `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                       `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                       `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                            `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                         `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                         `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                         `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                         `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                         `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                         `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                         `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnect struct {
	Authentication            string                                                                                               `json:"authentication,omitempty"`            //
	DataSwitching             string                                                                                               `json:"dataSwitching,omitempty"`             //
	EgressACLAnalysis         *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysis  `json:"egressAclAnalysis,omitempty"`         //
	IngressACLAnalysis        *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysis `json:"ingressAclAnalysis,omitempty"`        //
	WirelessLanControllerID   string                                                                                               `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string                                                                                               `json:"wirelessLanControllerName,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysis struct {
	ACLName      string                                                                                                            `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                            `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                         `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                         `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                              `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysis struct {
	ACLName      string                                                                                                             `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                             `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                          `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                          `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                               `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterface struct {
	PhysicalInterface *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterface  `json:"physicalInterface,omitempty"` //
	VirtualInterface  *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterface `json:"virtualInterface,omitempty"`  //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                      `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                      `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                      `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                      `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                      `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                      `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                      `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                      `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                            `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                            `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                         `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                         `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                              `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                           `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                           `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                           `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                           `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                           `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                           `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                           `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterface struct {
	ACLAnalysis                           *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                                     `json:"id,omitempty"`                                    //
	InterfaceStatistics                   *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                                     `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                                     `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                                     `json:"name,omitempty"`                                  //
	PathOverlayInfo                       *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                                     `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                                     `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                                     `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                                     `json:"vrfName,omitempty"`                               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                           `json:"aclName,omitempty"`      //
	MatchingAces *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                           `json:"result,omitempty"`       //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                                        `json:"ace,omitempty"`           //
	MatchingPorts *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                                        `json:"result,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    *[]ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                             `json:"protocol,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       *int   `json:"inputPackets,omitempty"`       //
	InputQueueCount    *int   `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    *int   `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  *int   `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth *int   `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       *int   `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         *int   `json:"outputDrop,omitempty"`         //
	OutputPackets      *int   `json:"outputPackets,omitempty"`      //
	OutputQueueCount   *int   `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   *int   `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      *int   `json:"outputRatebps,omitempty"`      //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                          `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                          `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                          `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                          `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                          `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                          `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                          `json:"sourcePort,omitempty"`              //
	VxlanInfo               *ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           *int   `json:"dropRate,omitempty"`           //
	NumBytes           *int   `json:"numBytes,omitempty"`           //
	NumPackets         *int   `json:"numPackets,omitempty"`         //
	OfferedRate        *int   `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         *int   `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops *int   `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    *int   `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        *int   `json:"refreshedAt,omitempty"`        //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoPerfMonitorStatistics struct {
	ByteRate             *int     `json:"byteRate,omitempty"`             //
	DestIPAddress        string   `json:"destIpAddress,omitempty"`        //
	DestPort             string   `json:"destPort,omitempty"`             //
	InputInterface       string   `json:"inputInterface,omitempty"`       //
	IPv4DSCP             string   `json:"ipv4DSCP,omitempty"`             //
	IPv4TTL              *int     `json:"ipv4TTL,omitempty"`              //
	OutputInterface      string   `json:"outputInterface,omitempty"`      //
	PacketBytes          *int     `json:"packetBytes,omitempty"`          //
	PacketCount          *int     `json:"packetCount,omitempty"`          //
	PacketLoss           *int     `json:"packetLoss,omitempty"`           //
	PacketLossPercentage *float64 `json:"packetLossPercentage,omitempty"` //
	Protocol             string   `json:"protocol,omitempty"`             //
	RefreshedAt          *int     `json:"refreshedAt,omitempty"`          //
	RtpJitterMax         *int     `json:"rtpJitterMax,omitempty"`         //
	RtpJitterMean        *int     `json:"rtpJitterMean,omitempty"`        //
	RtpJitterMin         *int     `json:"rtpJitterMin,omitempty"`         //
	SourceIPAddress      string   `json:"sourceIpAddress,omitempty"`      //
	SourcePort           string   `json:"sourcePort,omitempty"`           //
}
type ResponsePathTraceRetrievesPreviousPathtraceResponseRequest struct {
	ControlPath     *bool    `json:"controlPath,omitempty"`     //
	CreateTime      *int     `json:"createTime,omitempty"`      //
	DestIP          string   `json:"destIP,omitempty"`          //
	DestPort        string   `json:"destPort,omitempty"`        //
	FailureReason   string   `json:"failureReason,omitempty"`   //
	ID              string   `json:"id,omitempty"`              //
	Inclusions      []string `json:"inclusions,omitempty"`      //
	LastUpdateTime  *int     `json:"lastUpdateTime,omitempty"`  //
	PeriodicRefresh *bool    `json:"periodicRefresh,omitempty"` //
	Protocol        string   `json:"protocol,omitempty"`        //
	SourceIP        string   `json:"sourceIP,omitempty"`        //
	SourcePort      string   `json:"sourcePort,omitempty"`      //
	Status          string   `json:"status,omitempty"`          //
}
type ResponsePathTraceDeletesPathtraceByID struct {
	Response *ResponsePathTraceDeletesPathtraceByIDResponse `json:"response,omitempty"` //
	Version  string                                         `json:"version,omitempty"`  //
}
type ResponsePathTraceDeletesPathtraceByIDResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestPathTraceInitiateANewPathtrace struct {
	ControlPath     *bool    `json:"controlPath,omitempty"`     // Control path tracing
	DestIP          string   `json:"destIP,omitempty"`          // Destination IP address
	DestPort        string   `json:"destPort,omitempty"`        // Destination Port
	Inclusions      []string `json:"inclusions,omitempty"`      // Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
	PeriodicRefresh *bool    `json:"periodicRefresh,omitempty"` // Periodic refresh of path for every 30 sec
	Protocol        string   `json:"protocol,omitempty"`        // Protocol
	SourceIP        string   `json:"sourceIP,omitempty"`        // Source IP address
	SourcePort      string   `json:"sourcePort,omitempty"`      // Source Port
}

//RetrivesAllPreviousPathtracesSummary Retrives all previous Pathtraces summary - 55bc-3bf9-4e38-b6ff
/* Returns a summary of all flow analyses stored. Results can be filtered by specified parameters.


@param RetrivesAllPreviousPathtracesSummaryQueryParams Filtering parameter
*/
func (s *PathTraceService) RetrivesAllPreviousPathtracesSummary(RetrivesAllPreviousPathtracesSummaryQueryParams *RetrivesAllPreviousPathtracesSummaryQueryParams) (*ResponsePathTraceRetrivesAllPreviousPathtracesSummary, *resty.Response, error) {
	path := "/dna/intent/api/v1/flow-analysis"

	queryString, _ := query.Values(RetrivesAllPreviousPathtracesSummaryQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponsePathTraceRetrivesAllPreviousPathtracesSummary{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RetrivesAllPreviousPathtracesSummary")
	}

	result := response.Result().(*ResponsePathTraceRetrivesAllPreviousPathtracesSummary)
	return result, response, err

}

//RetrievesPreviousPathtrace Retrieves previous Pathtrace - 7ab9-a8bd-4f3b-86a4
/* Returns result of a previously requested flow analysis by its Flow Analysis id


@param flowAnalysisID flowAnalysisId path parameter. Flow analysis request id

*/
func (s *PathTraceService) RetrievesPreviousPathtrace(flowAnalysisID string) (*ResponsePathTraceRetrievesPreviousPathtrace, *resty.Response, error) {
	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{flowAnalysisId}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePathTraceRetrievesPreviousPathtrace{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RetrievesPreviousPathtrace")
	}

	result := response.Result().(*ResponsePathTraceRetrievesPreviousPathtrace)
	return result, response, err

}

//InitiateANewPathtrace Initiate a new Pathtrace - a395-fae6-44ca-899c
/* Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to get results and follow progress.


 */
func (s *PathTraceService) InitiateANewPathtrace(requestPathTraceInitiateANewPathtrace *RequestPathTraceInitiateANewPathtrace) (*ResponsePathTraceInitiateANewPathtrace, *resty.Response, error) {
	path := "/dna/intent/api/v1/flow-analysis"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestPathTraceInitiateANewPathtrace).
		SetResult(&ResponsePathTraceInitiateANewPathtrace{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation InitiateANewPathtrace")
	}

	result := response.Result().(*ResponsePathTraceInitiateANewPathtrace)
	return result, response, err

}

//DeletesPathtraceByID Deletes Pathtrace by Id - 8a9d-2b76-443b-914e
/* Deletes a flow analysis request by its id


@param flowAnalysisID flowAnalysisId path parameter. Flow analysis request id

*/
func (s *PathTraceService) DeletesPathtraceByID(flowAnalysisID string) (*ResponsePathTraceDeletesPathtraceByID, *resty.Response, error) {
	//flowAnalysisID string
	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{flowAnalysisId}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponsePathTraceDeletesPathtraceByID{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DeletesPathtraceById")
	}

	result := response.Result().(*ResponsePathTraceDeletesPathtraceByID)
	return result, response, err

}
