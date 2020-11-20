package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// PathTraceService is the service to communicate with the PathTrace API endpoint
type PathTraceService service

// InitiateANewPathtraceRequest is the initiateANewPathtraceRequest definition
type InitiateANewPathtraceRequest struct {
	ControlPath     bool     `json:"controlPath,omitempty"`     //
	DestIP          string   `json:"destIP,omitempty"`          //
	DestPort        string   `json:"destPort,omitempty"`        //
	Inclusions      []string `json:"inclusions,omitempty"`      //
	PeriodicRefresh bool     `json:"periodicRefresh,omitempty"` //
	Protocol        string   `json:"protocol,omitempty"`        //
	SourceIP        string   `json:"sourceIP,omitempty"`        //
	SourcePort      string   `json:"sourcePort,omitempty"`      //
}

// InitiateANewPathtraceRequestInclusions is the initiateANewPathtraceRequestInclusions definition
type InitiateANewPathtraceRequestInclusions []string

// DeletesPathtraceByIDResponseResponse is the deletesPathtraceByIDResponseResponse definition
type DeletesPathtraceByIDResponseResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// DeletesPathtraceByIDResponse is the deletesPathtraceByIdResponse definition
type DeletesPathtraceByIDResponse struct {
	Response DeletesPathtraceByIDResponseResponse `json:"response,omitempty"` //
	Version  string                               `json:"version,omitempty"`  //
}

// InitiateANewPathtraceResponse is the initiateANewPathtraceResponse definition
type InitiateANewPathtraceResponse struct {
	Response InitiateANewPathtraceResponseResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  //
}

// InitiateANewPathtraceResponseResponse is the initiateANewPathtraceResponseResponse definition
type InitiateANewPathtraceResponseResponse struct {
	FlowAnalysisID string `json:"flowAnalysisId,omitempty"` //
	TaskID         string `json:"taskId,omitempty"`         //
	URL            string `json:"url,omitempty"`            //
}

// RetrievesPreviousPathtraceResponse is the retrievesPreviousPathtraceResponse definition
type RetrievesPreviousPathtraceResponse struct {
	Response RetrievesPreviousPathtraceResponseResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  //
}

// RetrievesPreviousPathtraceResponseResponse is the retrievesPreviousPathtraceResponseResponse definition
type RetrievesPreviousPathtraceResponseResponse struct {
	DetailedStatus      RetrievesPreviousPathtraceResponseResponseDetailedStatus        `json:"detailedStatus,omitempty"`      //
	LastUpdate          int                                                          `json:"lastUpdate,omitempty"`          //
	NetworkElements     []RetrievesPreviousPathtraceResponseResponseNetworkElements     `json:"networkElements,omitempty"`     //
	NetworkElementsInfo []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfo `json:"networkElementsInfo,omitempty"` //
	Properties          []string                                                        `json:"properties,omitempty"`          //
	Request             RetrievesPreviousPathtraceResponseResponseRequest               `json:"request,omitempty"`             //
}

// RetrievesPreviousPathtraceResponseResponseDetailedStatus is the retrievesPreviousPathtraceResponseResponseDetailedStatus definition
type RetrievesPreviousPathtraceResponseResponseDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElements is the retrievesPreviousPathtraceResponseResponseNetworkElements definition
type RetrievesPreviousPathtraceResponseResponseNetworkElements struct {
	AccuracyList                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsAccuracyList           `json:"accuracyList,omitempty"`                       //
	DetailedStatus                     RetrievesPreviousPathtraceResponseResponseNetworkElementsDetailedStatus           `json:"detailedStatus,omitempty"`                     //
	DeviceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatistics         `json:"deviceStatistics,omitempty"`                   //
	DeviceStatsCollection              string                                                                            `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string                                                                            `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressPhysicalInterface            RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterface  `json:"egressPhysicalInterface,omitempty"`            //
	EgressVirtualInterface             RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterface   `json:"egressVirtualInterface,omitempty"`             //
	FlexConnect                        RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnect              `json:"flexConnect,omitempty"`                        //
	ID                                 string                                                                            `json:"id,omitempty"`                                 //
	IngressPhysicalInterface           RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterface `json:"ingressPhysicalInterface,omitempty"`           //
	IngressVirtualInterface            RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterface  `json:"ingressVirtualInterface,omitempty"`            //
	IP                                 string                                                                            `json:"ip,omitempty"`                                 //
	LinkInformationSource              string                                                                            `json:"linkInformationSource,omitempty"`              //
	Name                               string                                                                            `json:"name,omitempty"`                               //
	PerfMonCollection                  string                                                                            `json:"perfMonCollection,omitempty"`                  //
	PerfMonCollectionFailureReason     string                                                                            `json:"perfMonCollectionFailureReason,omitempty"`     //
	PerfMonStatistics                  []RetrievesPreviousPathtraceResponseResponseNetworkElementsPerfMonStatistics      `json:"perfMonStatistics,omitempty"`                  //
	Role                               string                                                                            `json:"role,omitempty"`                               //
	SSID                               string                                                                            `json:"ssid,omitempty"`                               //
	Tunnels                            []string                                                                          `json:"tunnels,omitempty"`                            //
	Type                               string                                                                            `json:"type,omitempty"`                               //
	WLANID                             string                                                                            `json:"wlanId,omitempty"`                             //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsAccuracyList is the retrievesPreviousPathtraceResponseResponseNetworkElementsAccuracyList definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsAccuracyList struct {
	Percent int    `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsDetailedStatus is the retrievesPreviousPathtraceResponseResponseNetworkElementsDetailedStatus definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatistics struct {
	CPUStatistics    RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsCPUStatistics    `json:"cpuStatistics,omitempty"`    //
	MemoryStatistics RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsMemoryStatistics `json:"memoryStatistics,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsCPUStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsCPUStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsCPUStatistics struct {
	FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               int `json:"refreshedAt,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsMemoryStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsMemoryStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsDeviceStatisticsMemoryStatistics struct {
	MemoryUsage int `json:"memoryUsage,omitempty"` //
	RefreshedAt int `json:"refreshedAt,omitempty"` //
	TotalMemory int `json:"totalMemory,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                              `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                              `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                              `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                              `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                              `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                              `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                              `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                              `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                    `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                    `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                 `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                 `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                      `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                   `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                   `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                   `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                   `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                   `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                   `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                   `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressPhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                             `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                             `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                             `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                             `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                             `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                             `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                             `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                             `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                   `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                   `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                     `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                  `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                  `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                  `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                  `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                  `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                  `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                  `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsEgressVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnect is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnect definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnect struct {
	Authentication            string                                                                                 `json:"authentication,omitempty"`            //
	DataSwitching             string                                                                                 `json:"dataSwitching,omitempty"`             //
	EgressACLAnalysis         RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysis  `json:"egressAclAnalysis,omitempty"`         //
	IngressACLAnalysis        RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysis `json:"ingressAclAnalysis,omitempty"`        //
	WirelessLanControllerID   string                                                                                 `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string                                                                                 `json:"wirelessLanControllerName,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysis struct {
	ACLName      string                                                                                              `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                              `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                           `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                           `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysis struct {
	ACLName      string                                                                                               `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                               `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                            `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                            `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                 `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfo struct {
	AccuracyList                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoAccuracyList          `json:"accuracyList,omitempty"`                       //
	DetailedStatus                     RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDetailedStatus          `json:"detailedStatus,omitempty"`                     //
	DeviceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatistics        `json:"deviceStatistics,omitempty"`                   //
	DeviceStatsCollection              string                                                                               `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string                                                                               `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressInterface                    RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterface         `json:"egressInterface,omitempty"`                    //
	FlexConnect                        RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnect             `json:"flexConnect,omitempty"`                        //
	ID                                 string                                                                               `json:"id,omitempty"`                                 //
	IngressInterface                   RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterface        `json:"ingressInterface,omitempty"`                   //
	IP                                 string                                                                               `json:"ip,omitempty"`                                 //
	LinkInformationSource              string                                                                               `json:"linkInformationSource,omitempty"`              //
	Name                               string                                                                               `json:"name,omitempty"`                               //
	PerfMonCollection                  string                                                                               `json:"perfMonCollection,omitempty"`                  //
	PerfMonCollectionFailureReason     string                                                                               `json:"perfMonCollectionFailureReason,omitempty"`     //
	PerfMonitorStatistics              []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoPerfMonitorStatistics `json:"perfMonitorStatistics,omitempty"`              //
	Role                               string                                                                               `json:"role,omitempty"`                               //
	SSID                               string                                                                               `json:"ssid,omitempty"`                               //
	Tunnels                            []string                                                                             `json:"tunnels,omitempty"`                            //
	Type                               string                                                                               `json:"type,omitempty"`                               //
	WLANID                             string                                                                               `json:"wlanId,omitempty"`                             //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoAccuracyList is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoAccuracyList definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoAccuracyList struct {
	Percent int    `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDetailedStatus is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoDetailedStatus definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatistics struct {
	CPUStatistics    RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsCPUStatistics    `json:"cpuStatistics,omitempty"`    //
	MemoryStatistics RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics `json:"memoryStatistics,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsCPUStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsCPUStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsCPUStatistics struct {
	FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               int `json:"refreshedAt,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics struct {
	MemoryUsage int `json:"memoryUsage,omitempty"` //
	RefreshedAt int `json:"refreshedAt,omitempty"` //
	TotalMemory int `json:"totalMemory,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterface struct {
	PhysicalInterface RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterface  `json:"physicalInterface,omitempty"` //
	VirtualInterface  []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterface `json:"virtualInterface,omitempty"`  //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                           `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                           `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                           `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                           `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                           `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                           `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                           `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                           `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                 `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                 `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                              `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                              `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                   `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                          `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                          `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                          `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                          `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                          `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                          `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                          `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                          `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                             `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                             `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                  `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                               `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                               `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                               `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                               `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                               `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                               `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                               `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnect is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnect definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnect struct {
	Authentication            string                                                                                     `json:"authentication,omitempty"`            //
	DataSwitching             string                                                                                     `json:"dataSwitching,omitempty"`             //
	EgressACLAnalysis         RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysis  `json:"egressAclAnalysis,omitempty"`         //
	IngressACLAnalysis        RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysis `json:"ingressAclAnalysis,omitempty"`        //
	WirelessLanControllerID   string                                                                                     `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string                                                                                     `json:"wirelessLanControllerName,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysis struct {
	ACLName      string                                                                                                  `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                  `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                               `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                               `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                    `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysis struct {
	ACLName      string                                                                                                   `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                   `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                     `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterface struct {
	PhysicalInterface RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterface  `json:"physicalInterface,omitempty"` //
	VirtualInterface  []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterface `json:"virtualInterface,omitempty"`  //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                            `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                            `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                            `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                            `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                            `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                            `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                            `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                            `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                  `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                  `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                               `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                               `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                    `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                 `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                 `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                 `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                 `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                 `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                 `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                 `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                                           `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                                           `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                                           `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                                           `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                                           `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                                           `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                                           `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                                           `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                                 `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                                 `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                              `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                              `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                                   `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                                `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                                `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                                `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                                `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                                `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                                `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                                `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoPerfMonitorStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoPerfMonitorStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoPerfMonitorStatistics struct {
	ByteRate             int    `json:"byteRate,omitempty"`             //
	DestIPAddress        string `json:"destIpAddress,omitempty"`        //
	DestPort             string `json:"destPort,omitempty"`             //
	InputInterface       string `json:"inputInterface,omitempty"`       //
	IPv4DSCP             string `json:"ipv4DSCP,omitempty"`             //
	IPv4TTL              int    `json:"ipv4TTL,omitempty"`              //
	OutputInterface      string `json:"outputInterface,omitempty"`      //
	PacketBytes          int    `json:"packetBytes,omitempty"`          //
	PacketCount          int    `json:"packetCount,omitempty"`          //
	PacketLoss           int    `json:"packetLoss,omitempty"`           //
	PacketLossPercentage int    `json:"packetLossPercentage,omitempty"` //
	Protocol             string `json:"protocol,omitempty"`             //
	RefreshedAt          int    `json:"refreshedAt,omitempty"`          //
	RtpJitterMax         int    `json:"rtpJitterMax,omitempty"`         //
	RtpJitterMean        int    `json:"rtpJitterMean,omitempty"`        //
	RtpJitterMin         int    `json:"rtpJitterMin,omitempty"`         //
	SourceIPAddress      string `json:"sourceIpAddress,omitempty"`      //
	SourcePort           string `json:"sourcePort,omitempty"`           //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoTunnels is the retrievesPreviousPathtraceResponseResponseNetworkElementsInfoTunnels definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsInfoTunnels []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                               `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                               `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                               `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                               `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                               `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                               `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                               `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                               `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis struct {
	ACLName      string                                                                                                     `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                     `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                  `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                  `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                       `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                    `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                    `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                    `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                    `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                    `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                    `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                    `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressPhysicalInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterface is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterface definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterface struct {
	ACLAnalysis                           RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysis         `json:"aclAnalysis,omitempty"`                           //
	ID                                    string                                                                                              `json:"id,omitempty"`                                    //
	InterfaceStatistics                   RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics `json:"interfaceStatistics,omitempty"`                   //
	InterfaceStatsCollection              string                                                                                              `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string                                                                                              `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string                                                                                              `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo   `json:"pathOverlayInfo,omitempty"`                       //
	QosStatistics                         []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceQosStatistics     `json:"qosStatistics,omitempty"`                         //
	QosStatsCollection                    string                                                                                              `json:"qosStatsCollection,omitempty"`                    //
	QosStatsCollectionFailureReason       string                                                                                              `json:"qosStatsCollectionFailureReason,omitempty"`       //
	UsedVLAN                              string                                                                                              `json:"usedVlan,omitempty"`                              //
	VrfName                               string                                                                                              `json:"vrfName,omitempty"`                               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysis is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysis definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysis struct {
	ACLName      string                                                                                                    `json:"aclName,omitempty"`      //
	MatchingAces []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces `json:"matchingAces,omitempty"` //
	Result       string                                                                                                    `json:"result,omitempty"`       //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces struct {
	Ace           string                                                                                                                 `json:"ace,omitempty"`           //
	MatchingPorts []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts `json:"matchingPorts,omitempty"` //
	Result        string                                                                                                                 `json:"result,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts struct {
	Ports    []RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts `json:"ports,omitempty"`    //
	Protocol string                                                                                                                      `json:"protocol,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsDestPorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPortsSourcePorts []string

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics struct {
	AdminStatus        string `json:"adminStatus,omitempty"`        //
	InputPackets       int    `json:"inputPackets,omitempty"`       //
	InputQueueCount    int    `json:"inputQueueCount,omitempty"`    //
	InputQueueDrops    int    `json:"inputQueueDrops,omitempty"`    //
	InputQueueFlushes  int    `json:"inputQueueFlushes,omitempty"`  //
	InputQueueMaxDepth int    `json:"inputQueueMaxDepth,omitempty"` //
	InputRatebps       int    `json:"inputRatebps,omitempty"`       //
	OperationalStatus  string `json:"operationalStatus,omitempty"`  //
	OutputDrop         int    `json:"outputDrop,omitempty"`         //
	OutputPackets      int    `json:"outputPackets,omitempty"`      //
	OutputQueueCount   int    `json:"outputQueueCount,omitempty"`   //
	OutputQueueDepth   int    `json:"outputQueueDepth,omitempty"`   //
	OutputRatebps      int    `json:"outputRatebps,omitempty"`      //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo struct {
	ControlPlane            string                                                                                                   `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string                                                                                                   `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string                                                                                                   `json:"destIp,omitempty"`                  //
	DestPort                string                                                                                                   `json:"destPort,omitempty"`                //
	Protocol                string                                                                                                   `json:"protocol,omitempty"`                //
	SourceIP                string                                                                                                   `json:"sourceIp,omitempty"`                //
	SourcePort              string                                                                                                   `json:"sourcePort,omitempty"`              //
	VxlanInfo               RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo `json:"vxlanInfo,omitempty"`               //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceQosStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceQosStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsIngressVirtualInterfaceQosStatistics struct {
	ClassMapName       string `json:"classMapName,omitempty"`       //
	DropRate           int    `json:"dropRate,omitempty"`           //
	NumBytes           int    `json:"numBytes,omitempty"`           //
	NumPackets         int    `json:"numPackets,omitempty"`         //
	OfferedRate        int    `json:"offeredRate,omitempty"`        //
	QueueBandwidthbps  string `json:"queueBandwidthbps,omitempty"`  //
	QueueDepth         int    `json:"queueDepth,omitempty"`         //
	QueueNoBufferDrops int    `json:"queueNoBufferDrops,omitempty"` //
	QueueTotalDrops    int    `json:"queueTotalDrops,omitempty"`    //
	RefreshedAt        int    `json:"refreshedAt,omitempty"`        //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsPerfMonStatistics is the retrievesPreviousPathtraceResponseResponseNetworkElementsPerfMonStatistics definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsPerfMonStatistics struct {
	ByteRate             int    `json:"byteRate,omitempty"`             //
	DestIPAddress        string `json:"destIpAddress,omitempty"`        //
	DestPort             string `json:"destPort,omitempty"`             //
	InputInterface       string `json:"inputInterface,omitempty"`       //
	IPv4DSCP             string `json:"ipv4DSCP,omitempty"`             //
	IPv4TTL              int    `json:"ipv4TTL,omitempty"`              //
	OutputInterface      string `json:"outputInterface,omitempty"`      //
	PacketBytes          int    `json:"packetBytes,omitempty"`          //
	PacketCount          int    `json:"packetCount,omitempty"`          //
	PacketLoss           int    `json:"packetLoss,omitempty"`           //
	PacketLossPercentage int    `json:"packetLossPercentage,omitempty"` //
	Protocol             string `json:"protocol,omitempty"`             //
	RefreshedAt          int    `json:"refreshedAt,omitempty"`          //
	RtpJitterMax         int    `json:"rtpJitterMax,omitempty"`         //
	RtpJitterMean        int    `json:"rtpJitterMean,omitempty"`        //
	RtpJitterMin         int    `json:"rtpJitterMin,omitempty"`         //
	SourceIPAddress      string `json:"sourceIpAddress,omitempty"`      //
	SourcePort           string `json:"sourcePort,omitempty"`           //
}

// RetrievesPreviousPathtraceResponseResponseNetworkElementsTunnels is the retrievesPreviousPathtraceResponseResponseNetworkElementsTunnels definition
type RetrievesPreviousPathtraceResponseResponseNetworkElementsTunnels []string

// RetrievesPreviousPathtraceResponseResponseProperties is the retrievesPreviousPathtraceResponseResponseProperties definition
type RetrievesPreviousPathtraceResponseResponseProperties []string

// RetrievesPreviousPathtraceResponseResponseRequest is the retrievesPreviousPathtraceResponseResponseRequest definition
type RetrievesPreviousPathtraceResponseResponseRequest struct {
	ControlPath     bool     `json:"controlPath,omitempty"`     //
	CreateTime      int      `json:"createTime,omitempty"`      //
	DestIP          string   `json:"destIP,omitempty"`          //
	DestPort        string   `json:"destPort,omitempty"`        //
	FailureReason   string   `json:"failureReason,omitempty"`   //
	ID              string   `json:"id,omitempty"`              //
	Inclusions      []string `json:"inclusions,omitempty"`      //
	LastUpdateTime  int      `json:"lastUpdateTime,omitempty"`  //
	PeriodicRefresh bool     `json:"periodicRefresh,omitempty"` //
	Protocol        string   `json:"protocol,omitempty"`        //
	SourceIP        string   `json:"sourceIP,omitempty"`        //
	SourcePort      string   `json:"sourcePort,omitempty"`      //
	Status          string   `json:"status,omitempty"`          //
}

// RetrievesPreviousPathtraceResponseResponseRequestInclusions is the retrievesPreviousPathtraceResponseResponseRequestInclusions definition
type RetrievesPreviousPathtraceResponseResponseRequestInclusions []string

// RetrivesAllPreviousPathtracesSummaryResponse is the retrivesAllPreviousPathtracesSummaryResponse definition
type RetrivesAllPreviousPathtracesSummaryResponse struct {
	Response []RetrivesAllPreviousPathtracesSummaryResponseResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  //
}

// RetrivesAllPreviousPathtracesSummaryResponseResponse is the retrivesAllPreviousPathtracesSummaryResponseResponse definition
type RetrivesAllPreviousPathtracesSummaryResponseResponse struct {
	ControlPath     bool     `json:"controlPath,omitempty"`     //
	CreateTime      int      `json:"createTime,omitempty"`      //
	DestIP          string   `json:"destIP,omitempty"`          //
	DestPort        string   `json:"destPort,omitempty"`        //
	FailureReason   string   `json:"failureReason,omitempty"`   //
	ID              string   `json:"id,omitempty"`              //
	Inclusions      []string `json:"inclusions,omitempty"`      //
	LastUpdateTime  int      `json:"lastUpdateTime,omitempty"`  //
	PeriodicRefresh bool     `json:"periodicRefresh,omitempty"` //
	Protocol        string   `json:"protocol,omitempty"`        //
	SourceIP        string   `json:"sourceIP,omitempty"`        //
	SourcePort      string   `json:"sourcePort,omitempty"`      //
	Status          string   `json:"status,omitempty"`          //
}

// RetrivesAllPreviousPathtracesSummaryResponseResponseInclusions is the retrivesAllPreviousPathtracesSummaryResponseResponseInclusions definition
type RetrivesAllPreviousPathtracesSummaryResponseResponseInclusions []string

// DeletesPathtraceByID deletesPathtraceById
/* Deletes a flow analysis request by its id
@param flowAnalysisID Flow analysis request id
*/
func (s *PathTraceService) DeletesPathtraceByID(flowAnalysisID string) (*DeletesPathtraceByIDResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{"+"flowAnalysisId"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := RestyClient.R().
		SetResult(&DeletesPathtraceByIDResponse{}).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation deletesPathtraceById")
	}

	result := response.Result().(*DeletesPathtraceByIDResponse)
	return result, response, err
}

// InitiateANewPathtrace initiateANewPathtrace
/* Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to get results and follow progress.
 */
func (s *PathTraceService) InitiateANewPathtrace(initiateANewPathtraceRequest *InitiateANewPathtraceRequest) (*InitiateANewPathtraceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis"

	response, err := RestyClient.R().
		SetBody(initiateANewPathtraceRequest).
		SetResult(&InitiateANewPathtraceResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation initiateANewPathtrace")
	}

	result := response.Result().(*InitiateANewPathtraceResponse)
	return result, response, err
}

// RetrievesPreviousPathtrace retrievesPreviousPathtrace
/* Returns result of a previously requested flow analysis by its Flow Analysis id
@param flowAnalysisID Flow analysis request id
*/
func (s *PathTraceService) RetrievesPreviousPathtrace(flowAnalysisID string) (*RetrievesPreviousPathtraceResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{"+"flowAnalysisId"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := RestyClient.R().
		SetResult(&RetrievesPreviousPathtraceResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation retrievesPreviousPathtrace")
	}

	result := response.Result().(*RetrievesPreviousPathtraceResponse)
	return result, response, err
}

// RetrivesAllPreviousPathtracesSummaryQueryParams defines the query parameters for this request
type RetrivesAllPreviousPathtracesSummaryQueryParams struct {
	PeriodicRefresh bool   `url:"periodicRefresh,omitempty"` // Is analysis periodically refreshed?
	SourceIP        string `url:"sourceIP,omitempty"`        // Source IP address
	DestIP          string `url:"destIP,omitempty"`          // Destination IP adress
	SourcePort      string `url:"sourcePort,omitempty"`      // Source port
	DestPort        string `url:"destPort,omitempty"`        // Destination port
	GtCreateTime    string `url:"gtCreateTime,omitempty"`    // Analyses requested after this time
	LtCreateTime    string `url:"ltCreateTime,omitempty"`    // Analyses requested before this time
	Protocol        string `url:"protocol,omitempty"`        // Protocol
	Status          string `url:"status,omitempty"`          // Status
	TaskID          string `url:"taskId,omitempty"`          // Task ID
	LastUpdateTime  string `url:"lastUpdateTime,omitempty"`  // Last update time
	Limit           string `url:"limit,omitempty"`           // Number of resources returned
	Offset          string `url:"offset,omitempty"`          // Start index of resources returned (1-based)
	Order           string `url:"order,omitempty"`           // Order by this field
	SortBy          string `url:"sortBy,omitempty"`          // Sort by this field
}

// RetrivesAllPreviousPathtracesSummary retrivesAllPreviousPathtracesSummary
/* Returns a summary of all flow analyses stored. Results can be filtered by specified parameters.
@param periodicRefresh Is analysis periodically refreshed?
@param sourceIP Source IP address
@param destIP Destination IP adress
@param sourcePort Source port
@param destPort Destination port
@param gtCreateTime Analyses requested after this time
@param ltCreateTime Analyses requested before this time
@param protocol Protocol
@param status Status
@param taskID Task ID
@param lastUpdateTime Last update time
@param limit Number of resources returned
@param offset Start index of resources returned (1-based)
@param order Order by this field
@param sortBy Sort by this field
*/
func (s *PathTraceService) RetrivesAllPreviousPathtracesSummary(retrivesAllPreviousPathtracesSummaryQueryParams *RetrivesAllPreviousPathtracesSummaryQueryParams) (*RetrivesAllPreviousPathtracesSummaryResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis"

	queryString, _ := query.Values(retrivesAllPreviousPathtracesSummaryQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&RetrivesAllPreviousPathtracesSummaryResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	if response.IsError() {
		return nil, response, fmt.Errorf("Error with operation retrivesAllPreviousPathtracesSummary")
	}

	result := response.Result().(*RetrivesAllPreviousPathtracesSummaryResponse)
	return result, response, err
}
