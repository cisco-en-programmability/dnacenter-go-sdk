package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// PathTraceService is the service to communicate with the PathTrace API endpoint
type PathTraceService service

// FlowAnalysisRequest is the FlowAnalysisRequest definition
type FlowAnalysisRequest struct {
	ControlPath     bool     `json:"controlPath,omitempty"`     //
	DestIP          string   `json:"destIP,omitempty"`          //
	DestPort        string   `json:"destPort,omitempty"`        //
	Inclusions      []string `json:"inclusions,omitempty"`      //
	PeriodicRefresh bool     `json:"periodicRefresh,omitempty"` //
	Protocol        string   `json:"protocol,omitempty"`        //
	SourceIP        string   `json:"sourceIP,omitempty"`        //
	SourcePort      string   `json:"sourcePort,omitempty"`      //
}

// AccuracyList is the AccuracyList definition
type AccuracyList struct {
	Percent int    `json:"percent,omitempty"` //
	Reason  string `json:"reason,omitempty"`  //
}

// ACLAnalysis is the AclAnalysis definition
type ACLAnalysis struct {
	ACLName      string `json:"aclName,omitempty"` //
	MatchingAces []struct {
		Ace           string `json:"ace,omitempty"` //
		MatchingPorts []struct {
			Ports []struct {
				DestPorts   []string `json:"destPorts,omitempty"`   //
				SourcePorts []string `json:"sourcePorts,omitempty"` //
			} `json:"ports,omitempty"` //
			Protocol string `json:"protocol,omitempty"` //
		} `json:"matchingPorts,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"matchingAces,omitempty"` //
	Result string `json:"result,omitempty"` //
}

// CPUStatistics is the CpuStatistics definition
type CPUStatistics struct {
	FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
	FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
	OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
	RefreshedAt               int `json:"refreshedAt,omitempty"`               //
}

// DetailedStatus is the DetailedStatus definition
type DetailedStatus struct {
	ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
	ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
}

// DeviceStatistics is the DeviceStatistics definition
type DeviceStatistics struct {
	CPUStatistics struct {
		FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
		FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
		OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
		RefreshedAt               int `json:"refreshedAt,omitempty"`               //
	} `json:"cpuStatistics,omitempty"` //
	MemoryStatistics struct {
		MemoryUsage int `json:"memoryUsage,omitempty"` //
		RefreshedAt int `json:"refreshedAt,omitempty"` //
		TotalMemory int `json:"totalMemory,omitempty"` //
	} `json:"memoryStatistics,omitempty"` //
}

// EgressACLAnalysis is the EgressAclAnalysis definition
type EgressACLAnalysis struct {
	ACLName      string `json:"aclName,omitempty"` //
	MatchingAces []struct {
		Ace           string `json:"ace,omitempty"` //
		MatchingPorts []struct {
			Ports []struct {
				DestPorts   []string `json:"destPorts,omitempty"`   //
				SourcePorts []string `json:"sourcePorts,omitempty"` //
			} `json:"ports,omitempty"` //
			Protocol string `json:"protocol,omitempty"` //
		} `json:"matchingPorts,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"matchingAces,omitempty"` //
	Result string `json:"result,omitempty"` //
}

// EgressInterface is the EgressInterface definition
type EgressInterface struct {
	PhysicalInterface struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"physicalInterface,omitempty"` //
	VirtualInterface []struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"virtualInterface,omitempty"` //
}

// EgressPhysicalInterface is the EgressPhysicalInterface definition
type EgressPhysicalInterface struct {
	ACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"aclAnalysis,omitempty"` //
	ID                  string `json:"id,omitempty"` //
	InterfaceStatistics struct {
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
	} `json:"interfaceStatistics,omitempty"` //
	InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []struct {
		ControlPlane            string `json:"controlPlane,omitempty"`            //
		DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
		DestIP                  string `json:"destIp,omitempty"`                  //
		DestPort                string `json:"destPort,omitempty"`                //
		Protocol                string `json:"protocol,omitempty"`                //
		SourceIP                string `json:"sourceIp,omitempty"`                //
		SourcePort              string `json:"sourcePort,omitempty"`              //
		VxlanInfo               struct {
			Dscp string `json:"dscp,omitempty"` //
			Vnid string `json:"vnid,omitempty"` //
		} `json:"vxlanInfo,omitempty"` //
	} `json:"pathOverlayInfo,omitempty"` //
	QosStatistics []struct {
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
	} `json:"qosStatistics,omitempty"` //
	QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
	QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
	UsedVlan                        string `json:"usedVlan,omitempty"`                        //
	VrfName                         string `json:"vrfName,omitempty"`                         //
}

// EgressVirtualInterface is the EgressVirtualInterface definition
type EgressVirtualInterface struct {
	ACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"aclAnalysis,omitempty"` //
	ID                  string `json:"id,omitempty"` //
	InterfaceStatistics struct {
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
	} `json:"interfaceStatistics,omitempty"` //
	InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []struct {
		ControlPlane            string `json:"controlPlane,omitempty"`            //
		DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
		DestIP                  string `json:"destIp,omitempty"`                  //
		DestPort                string `json:"destPort,omitempty"`                //
		Protocol                string `json:"protocol,omitempty"`                //
		SourceIP                string `json:"sourceIp,omitempty"`                //
		SourcePort              string `json:"sourcePort,omitempty"`              //
		VxlanInfo               struct {
			Dscp string `json:"dscp,omitempty"` //
			Vnid string `json:"vnid,omitempty"` //
		} `json:"vxlanInfo,omitempty"` //
	} `json:"pathOverlayInfo,omitempty"` //
	QosStatistics []struct {
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
	} `json:"qosStatistics,omitempty"` //
	QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
	QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
	UsedVlan                        string `json:"usedVlan,omitempty"`                        //
	VrfName                         string `json:"vrfName,omitempty"`                         //
}

// FlexConnect is the FlexConnect definition
type FlexConnect struct {
	Authentication    string `json:"authentication,omitempty"` //
	DataSwitching     string `json:"dataSwitching,omitempty"`  //
	EgressACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"egressACLAnalysis,omitempty"` //
	IngressACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"ingressACLAnalysis,omitempty"` //
	WirelessLanControllerID   string `json:"wirelessLanControllerId,omitempty"`   //
	WirelessLanControllerName string `json:"wirelessLanControllerName,omitempty"` //
}

// FlowAnalysisListOutput is the FlowAnalysisListOutput definition
type FlowAnalysisListOutput struct {
	Response []struct {
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
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// FlowAnalysisRequestResultOutput is the FlowAnalysisRequestResultOutput definition
type FlowAnalysisRequestResultOutput struct {
	Response struct {
		FlowAnalysisID string `json:"flowAnalysisId,omitempty"` //
		TaskID         string `json:"taskId,omitempty"`         //
		URL            string `json:"url,omitempty"`            //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// IngressACLAnalysis is the IngressAclAnalysis definition
type IngressACLAnalysis struct {
	ACLName      string `json:"aclName,omitempty"` //
	MatchingAces []struct {
		Ace           string `json:"ace,omitempty"` //
		MatchingPorts []struct {
			Ports []struct {
				DestPorts   []string `json:"destPorts,omitempty"`   //
				SourcePorts []string `json:"sourcePorts,omitempty"` //
			} `json:"ports,omitempty"` //
			Protocol string `json:"protocol,omitempty"` //
		} `json:"matchingPorts,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"matchingAces,omitempty"` //
	Result string `json:"result,omitempty"` //
}

// IngressInterface is the IngressInterface definition
type IngressInterface struct {
	PhysicalInterface struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"physicalInterface,omitempty"` //
	VirtualInterface []struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"virtualInterface,omitempty"` //
}

// IngressPhysicalInterface is the IngressPhysicalInterface definition
type IngressPhysicalInterface struct {
	ACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"aclAnalysis,omitempty"` //
	ID                  string `json:"id,omitempty"` //
	InterfaceStatistics struct {
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
	} `json:"interfaceStatistics,omitempty"` //
	InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []struct {
		ControlPlane            string `json:"controlPlane,omitempty"`            //
		DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
		DestIP                  string `json:"destIp,omitempty"`                  //
		DestPort                string `json:"destPort,omitempty"`                //
		Protocol                string `json:"protocol,omitempty"`                //
		SourceIP                string `json:"sourceIp,omitempty"`                //
		SourcePort              string `json:"sourcePort,omitempty"`              //
		VxlanInfo               struct {
			Dscp string `json:"dscp,omitempty"` //
			Vnid string `json:"vnid,omitempty"` //
		} `json:"vxlanInfo,omitempty"` //
	} `json:"pathOverlayInfo,omitempty"` //
	QosStatistics []struct {
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
	} `json:"qosStatistics,omitempty"` //
	QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
	QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
	UsedVlan                        string `json:"usedVlan,omitempty"`                        //
	VrfName                         string `json:"vrfName,omitempty"`                         //
}

// IngressVirtualInterface is the IngressVirtualInterface definition
type IngressVirtualInterface struct {
	ACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"aclAnalysis,omitempty"` //
	ID                  string `json:"id,omitempty"` //
	InterfaceStatistics struct {
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
	} `json:"interfaceStatistics,omitempty"` //
	InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []struct {
		ControlPlane            string `json:"controlPlane,omitempty"`            //
		DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
		DestIP                  string `json:"destIp,omitempty"`                  //
		DestPort                string `json:"destPort,omitempty"`                //
		Protocol                string `json:"protocol,omitempty"`                //
		SourceIP                string `json:"sourceIp,omitempty"`                //
		SourcePort              string `json:"sourcePort,omitempty"`              //
		VxlanInfo               struct {
			Dscp string `json:"dscp,omitempty"` //
			Vnid string `json:"vnid,omitempty"` //
		} `json:"vxlanInfo,omitempty"` //
	} `json:"pathOverlayInfo,omitempty"` //
	QosStatistics []struct {
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
	} `json:"qosStatistics,omitempty"` //
	QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
	QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
	UsedVlan                        string `json:"usedVlan,omitempty"`                        //
	VrfName                         string `json:"vrfName,omitempty"`                         //
}

// InterfaceStatistics is the InterfaceStatistics definition
type InterfaceStatistics struct {
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

// MatchingAces is the MatchingAces definition
type MatchingAces struct {
	Ace           string `json:"ace,omitempty"` //
	MatchingPorts []struct {
		Ports []struct {
			DestPorts   []string `json:"destPorts,omitempty"`   //
			SourcePorts []string `json:"sourcePorts,omitempty"` //
		} `json:"ports,omitempty"` //
		Protocol string `json:"protocol,omitempty"` //
	} `json:"matchingPorts,omitempty"` //
	Result string `json:"result,omitempty"` //
}

// MatchingPorts is the MatchingPorts definition
type MatchingPorts struct {
	Ports []struct {
		DestPorts   []string `json:"destPorts,omitempty"`   //
		SourcePorts []string `json:"sourcePorts,omitempty"` //
	} `json:"ports,omitempty"` //
	Protocol string `json:"protocol,omitempty"` //
}

// MemoryStatistics is the MemoryStatistics definition
type MemoryStatistics struct {
	MemoryUsage int `json:"memoryUsage,omitempty"` //
	RefreshedAt int `json:"refreshedAt,omitempty"` //
	TotalMemory int `json:"totalMemory,omitempty"` //
}

// NetworkElements is the NetworkElements definition
type NetworkElements struct {
	AccuracyList []struct {
		Percent int    `json:"percent,omitempty"` //
		Reason  string `json:"reason,omitempty"`  //
	} `json:"accuracyList,omitempty"` //
	DetailedStatus struct {
		ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
		ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
	} `json:"detailedStatus,omitempty"` //
	DeviceStatistics struct {
		CPUStatistics struct {
			FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
			FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
			OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
			RefreshedAt               int `json:"refreshedAt,omitempty"`               //
		} `json:"cpuStatistics,omitempty"` //
		MemoryStatistics struct {
			MemoryUsage int `json:"memoryUsage,omitempty"` //
			RefreshedAt int `json:"refreshedAt,omitempty"` //
			TotalMemory int `json:"totalMemory,omitempty"` //
		} `json:"memoryStatistics,omitempty"` //
	} `json:"deviceStatistics,omitempty"` //
	DeviceStatsCollection              string `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressPhysicalInterface            struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"egressPhysicalInterface,omitempty"` //
	EgressVirtualInterface struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"egressVirtualInterface,omitempty"` //
	FlexConnect struct {
		Authentication    string `json:"authentication,omitempty"` //
		DataSwitching     string `json:"dataSwitching,omitempty"`  //
		EgressACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"egressACLAnalysis,omitempty"` //
		IngressACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"ingressACLAnalysis,omitempty"` //
		WirelessLanControllerID   string `json:"wirelessLanControllerId,omitempty"`   //
		WirelessLanControllerName string `json:"wirelessLanControllerName,omitempty"` //
	} `json:"flexConnect,omitempty"` //
	ID                       string `json:"id,omitempty"` //
	IngressPhysicalInterface struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"ingressPhysicalInterface,omitempty"` //
	IngressVirtualInterface struct {
		ACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"aclAnalysis,omitempty"` //
		ID                  string `json:"id,omitempty"` //
		InterfaceStatistics struct {
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
		} `json:"interfaceStatistics,omitempty"` //
		InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
		InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
		Name                                  string `json:"name,omitempty"`                                  //
		PathOverlayInfo                       []struct {
			ControlPlane            string `json:"controlPlane,omitempty"`            //
			DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
			DestIP                  string `json:"destIp,omitempty"`                  //
			DestPort                string `json:"destPort,omitempty"`                //
			Protocol                string `json:"protocol,omitempty"`                //
			SourceIP                string `json:"sourceIp,omitempty"`                //
			SourcePort              string `json:"sourcePort,omitempty"`              //
			VxlanInfo               struct {
				Dscp string `json:"dscp,omitempty"` //
				Vnid string `json:"vnid,omitempty"` //
			} `json:"vxlanInfo,omitempty"` //
		} `json:"pathOverlayInfo,omitempty"` //
		QosStatistics []struct {
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
		} `json:"qosStatistics,omitempty"` //
		QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
		QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
		UsedVlan                        string `json:"usedVlan,omitempty"`                        //
		VrfName                         string `json:"vrfName,omitempty"`                         //
	} `json:"ingressVirtualInterface,omitempty"` //
	IP                             string `json:"ip,omitempty"`                             //
	LinkInformationSource          string `json:"linkInformationSource,omitempty"`          //
	Name                           string `json:"name,omitempty"`                           //
	PerfMonCollection              string `json:"perfMonCollection,omitempty"`              //
	PerfMonCollectionFailureReason string `json:"perfMonCollectionFailureReason,omitempty"` //
	PerfMonStatistics              []struct {
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
	} `json:"perfMonStatistics,omitempty"` //
	Role    string   `json:"role,omitempty"`    //
	SSID    string   `json:"ssid,omitempty"`    //
	Tunnels []string `json:"tunnels,omitempty"` //
	Type    string   `json:"type,omitempty"`    //
	WlanID  string   `json:"wlanId,omitempty"`  //
}

// NetworkElementsInfo is the NetworkElementsInfo definition
type NetworkElementsInfo struct {
	AccuracyList []struct {
		Percent int    `json:"percent,omitempty"` //
		Reason  string `json:"reason,omitempty"`  //
	} `json:"accuracyList,omitempty"` //
	DetailedStatus struct {
		ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
		ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
	} `json:"detailedStatus,omitempty"` //
	DeviceStatistics struct {
		CPUStatistics struct {
			FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
			FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
			OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
			RefreshedAt               int `json:"refreshedAt,omitempty"`               //
		} `json:"cpuStatistics,omitempty"` //
		MemoryStatistics struct {
			MemoryUsage int `json:"memoryUsage,omitempty"` //
			RefreshedAt int `json:"refreshedAt,omitempty"` //
			TotalMemory int `json:"totalMemory,omitempty"` //
		} `json:"memoryStatistics,omitempty"` //
	} `json:"deviceStatistics,omitempty"` //
	DeviceStatsCollection              string `json:"deviceStatsCollection,omitempty"`              //
	DeviceStatsCollectionFailureReason string `json:"deviceStatsCollectionFailureReason,omitempty"` //
	EgressInterface                    struct {
		PhysicalInterface struct {
			ACLAnalysis struct {
				ACLName      string `json:"aclName,omitempty"` //
				MatchingAces []struct {
					Ace           string `json:"ace,omitempty"` //
					MatchingPorts []struct {
						Ports []struct {
							DestPorts   []string `json:"destPorts,omitempty"`   //
							SourcePorts []string `json:"sourcePorts,omitempty"` //
						} `json:"ports,omitempty"` //
						Protocol string `json:"protocol,omitempty"` //
					} `json:"matchingPorts,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"matchingAces,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"aclAnalysis,omitempty"` //
			ID                  string `json:"id,omitempty"` //
			InterfaceStatistics struct {
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
			} `json:"interfaceStatistics,omitempty"` //
			InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
			InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
			Name                                  string `json:"name,omitempty"`                                  //
			PathOverlayInfo                       []struct {
				ControlPlane            string `json:"controlPlane,omitempty"`            //
				DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
				DestIP                  string `json:"destIp,omitempty"`                  //
				DestPort                string `json:"destPort,omitempty"`                //
				Protocol                string `json:"protocol,omitempty"`                //
				SourceIP                string `json:"sourceIp,omitempty"`                //
				SourcePort              string `json:"sourcePort,omitempty"`              //
				VxlanInfo               struct {
					Dscp string `json:"dscp,omitempty"` //
					Vnid string `json:"vnid,omitempty"` //
				} `json:"vxlanInfo,omitempty"` //
			} `json:"pathOverlayInfo,omitempty"` //
			QosStatistics []struct {
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
			} `json:"qosStatistics,omitempty"` //
			QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
			QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
			UsedVlan                        string `json:"usedVlan,omitempty"`                        //
			VrfName                         string `json:"vrfName,omitempty"`                         //
		} `json:"physicalInterface,omitempty"` //
		VirtualInterface []struct {
			ACLAnalysis struct {
				ACLName      string `json:"aclName,omitempty"` //
				MatchingAces []struct {
					Ace           string `json:"ace,omitempty"` //
					MatchingPorts []struct {
						Ports []struct {
							DestPorts   []string `json:"destPorts,omitempty"`   //
							SourcePorts []string `json:"sourcePorts,omitempty"` //
						} `json:"ports,omitempty"` //
						Protocol string `json:"protocol,omitempty"` //
					} `json:"matchingPorts,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"matchingAces,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"aclAnalysis,omitempty"` //
			ID                  string `json:"id,omitempty"` //
			InterfaceStatistics struct {
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
			} `json:"interfaceStatistics,omitempty"` //
			InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
			InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
			Name                                  string `json:"name,omitempty"`                                  //
			PathOverlayInfo                       []struct {
				ControlPlane            string `json:"controlPlane,omitempty"`            //
				DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
				DestIP                  string `json:"destIp,omitempty"`                  //
				DestPort                string `json:"destPort,omitempty"`                //
				Protocol                string `json:"protocol,omitempty"`                //
				SourceIP                string `json:"sourceIp,omitempty"`                //
				SourcePort              string `json:"sourcePort,omitempty"`              //
				VxlanInfo               struct {
					Dscp string `json:"dscp,omitempty"` //
					Vnid string `json:"vnid,omitempty"` //
				} `json:"vxlanInfo,omitempty"` //
			} `json:"pathOverlayInfo,omitempty"` //
			QosStatistics []struct {
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
			} `json:"qosStatistics,omitempty"` //
			QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
			QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
			UsedVlan                        string `json:"usedVlan,omitempty"`                        //
			VrfName                         string `json:"vrfName,omitempty"`                         //
		} `json:"virtualInterface,omitempty"` //
	} `json:"egressInterface,omitempty"` //
	FlexConnect struct {
		Authentication    string `json:"authentication,omitempty"` //
		DataSwitching     string `json:"dataSwitching,omitempty"`  //
		EgressACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"egressACLAnalysis,omitempty"` //
		IngressACLAnalysis struct {
			ACLName      string `json:"aclName,omitempty"` //
			MatchingAces []struct {
				Ace           string `json:"ace,omitempty"` //
				MatchingPorts []struct {
					Ports []struct {
						DestPorts   []string `json:"destPorts,omitempty"`   //
						SourcePorts []string `json:"sourcePorts,omitempty"` //
					} `json:"ports,omitempty"` //
					Protocol string `json:"protocol,omitempty"` //
				} `json:"matchingPorts,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"matchingAces,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"ingressACLAnalysis,omitempty"` //
		WirelessLanControllerID   string `json:"wirelessLanControllerId,omitempty"`   //
		WirelessLanControllerName string `json:"wirelessLanControllerName,omitempty"` //
	} `json:"flexConnect,omitempty"` //
	ID               string `json:"id,omitempty"` //
	IngressInterface struct {
		PhysicalInterface struct {
			ACLAnalysis struct {
				ACLName      string `json:"aclName,omitempty"` //
				MatchingAces []struct {
					Ace           string `json:"ace,omitempty"` //
					MatchingPorts []struct {
						Ports []struct {
							DestPorts   []string `json:"destPorts,omitempty"`   //
							SourcePorts []string `json:"sourcePorts,omitempty"` //
						} `json:"ports,omitempty"` //
						Protocol string `json:"protocol,omitempty"` //
					} `json:"matchingPorts,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"matchingAces,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"aclAnalysis,omitempty"` //
			ID                  string `json:"id,omitempty"` //
			InterfaceStatistics struct {
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
			} `json:"interfaceStatistics,omitempty"` //
			InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
			InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
			Name                                  string `json:"name,omitempty"`                                  //
			PathOverlayInfo                       []struct {
				ControlPlane            string `json:"controlPlane,omitempty"`            //
				DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
				DestIP                  string `json:"destIp,omitempty"`                  //
				DestPort                string `json:"destPort,omitempty"`                //
				Protocol                string `json:"protocol,omitempty"`                //
				SourceIP                string `json:"sourceIp,omitempty"`                //
				SourcePort              string `json:"sourcePort,omitempty"`              //
				VxlanInfo               struct {
					Dscp string `json:"dscp,omitempty"` //
					Vnid string `json:"vnid,omitempty"` //
				} `json:"vxlanInfo,omitempty"` //
			} `json:"pathOverlayInfo,omitempty"` //
			QosStatistics []struct {
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
			} `json:"qosStatistics,omitempty"` //
			QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
			QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
			UsedVlan                        string `json:"usedVlan,omitempty"`                        //
			VrfName                         string `json:"vrfName,omitempty"`                         //
		} `json:"physicalInterface,omitempty"` //
		VirtualInterface []struct {
			ACLAnalysis struct {
				ACLName      string `json:"aclName,omitempty"` //
				MatchingAces []struct {
					Ace           string `json:"ace,omitempty"` //
					MatchingPorts []struct {
						Ports []struct {
							DestPorts   []string `json:"destPorts,omitempty"`   //
							SourcePorts []string `json:"sourcePorts,omitempty"` //
						} `json:"ports,omitempty"` //
						Protocol string `json:"protocol,omitempty"` //
					} `json:"matchingPorts,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"matchingAces,omitempty"` //
				Result string `json:"result,omitempty"` //
			} `json:"aclAnalysis,omitempty"` //
			ID                  string `json:"id,omitempty"` //
			InterfaceStatistics struct {
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
			} `json:"interfaceStatistics,omitempty"` //
			InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
			InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
			Name                                  string `json:"name,omitempty"`                                  //
			PathOverlayInfo                       []struct {
				ControlPlane            string `json:"controlPlane,omitempty"`            //
				DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
				DestIP                  string `json:"destIp,omitempty"`                  //
				DestPort                string `json:"destPort,omitempty"`                //
				Protocol                string `json:"protocol,omitempty"`                //
				SourceIP                string `json:"sourceIp,omitempty"`                //
				SourcePort              string `json:"sourcePort,omitempty"`              //
				VxlanInfo               struct {
					Dscp string `json:"dscp,omitempty"` //
					Vnid string `json:"vnid,omitempty"` //
				} `json:"vxlanInfo,omitempty"` //
			} `json:"pathOverlayInfo,omitempty"` //
			QosStatistics []struct {
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
			} `json:"qosStatistics,omitempty"` //
			QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
			QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
			UsedVlan                        string `json:"usedVlan,omitempty"`                        //
			VrfName                         string `json:"vrfName,omitempty"`                         //
		} `json:"virtualInterface,omitempty"` //
	} `json:"ingressInterface,omitempty"` //
	IP                             string `json:"ip,omitempty"`                             //
	LinkInformationSource          string `json:"linkInformationSource,omitempty"`          //
	Name                           string `json:"name,omitempty"`                           //
	PerfMonCollection              string `json:"perfMonCollection,omitempty"`              //
	PerfMonCollectionFailureReason string `json:"perfMonCollectionFailureReason,omitempty"` //
	PerfMonitorStatistics          []struct {
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
	} `json:"perfMonitorStatistics,omitempty"` //
	Role    string   `json:"role,omitempty"`    //
	SSID    string   `json:"ssid,omitempty"`    //
	Tunnels []string `json:"tunnels,omitempty"` //
	Type    string   `json:"type,omitempty"`    //
	WlanID  string   `json:"wlanId,omitempty"`  //
}

// PathOverlayInfo is the PathOverlayInfo definition
type PathOverlayInfo struct {
	ControlPlane            string `json:"controlPlane,omitempty"`            //
	DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
	DestIP                  string `json:"destIp,omitempty"`                  //
	DestPort                string `json:"destPort,omitempty"`                //
	Protocol                string `json:"protocol,omitempty"`                //
	SourceIP                string `json:"sourceIp,omitempty"`                //
	SourcePort              string `json:"sourcePort,omitempty"`              //
	VxlanInfo               struct {
		Dscp string `json:"dscp,omitempty"` //
		Vnid string `json:"vnid,omitempty"` //
	} `json:"vxlanInfo,omitempty"` //
}

// PathResponseResult is the PathResponseResult definition
type PathResponseResult struct {
	Response struct {
		DetailedStatus struct {
			ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
			ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
		} `json:"detailedStatus,omitempty"` //
		LastUpdate      string `json:"lastUpdate,omitempty"` //
		NetworkElements []struct {
			AccuracyList []struct {
				Percent int    `json:"percent,omitempty"` //
				Reason  string `json:"reason,omitempty"`  //
			} `json:"accuracyList,omitempty"` //
			DetailedStatus struct {
				ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
				ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
			} `json:"detailedStatus,omitempty"` //
			DeviceStatistics struct {
				CPUStatistics struct {
					FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
					FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
					OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
					RefreshedAt               int `json:"refreshedAt,omitempty"`               //
				} `json:"cpuStatistics,omitempty"` //
				MemoryStatistics struct {
					MemoryUsage int `json:"memoryUsage,omitempty"` //
					RefreshedAt int `json:"refreshedAt,omitempty"` //
					TotalMemory int `json:"totalMemory,omitempty"` //
				} `json:"memoryStatistics,omitempty"` //
			} `json:"deviceStatistics,omitempty"` //
			DeviceStatsCollection              string `json:"deviceStatsCollection,omitempty"`              //
			DeviceStatsCollectionFailureReason string `json:"deviceStatsCollectionFailureReason,omitempty"` //
			EgressPhysicalInterface            struct {
				ACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"aclAnalysis,omitempty"` //
				ID                  string `json:"id,omitempty"` //
				InterfaceStatistics struct {
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
				} `json:"interfaceStatistics,omitempty"` //
				InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
				InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
				Name                                  string `json:"name,omitempty"`                                  //
				PathOverlayInfo                       []struct {
					ControlPlane            string `json:"controlPlane,omitempty"`            //
					DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
					DestIP                  string `json:"destIp,omitempty"`                  //
					DestPort                string `json:"destPort,omitempty"`                //
					Protocol                string `json:"protocol,omitempty"`                //
					SourceIP                string `json:"sourceIp,omitempty"`                //
					SourcePort              string `json:"sourcePort,omitempty"`              //
					VxlanInfo               struct {
						Dscp string `json:"dscp,omitempty"` //
						Vnid string `json:"vnid,omitempty"` //
					} `json:"vxlanInfo,omitempty"` //
				} `json:"pathOverlayInfo,omitempty"` //
				QosStatistics []struct {
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
				} `json:"qosStatistics,omitempty"` //
				QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
				QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
				UsedVlan                        string `json:"usedVlan,omitempty"`                        //
				VrfName                         string `json:"vrfName,omitempty"`                         //
			} `json:"egressPhysicalInterface,omitempty"` //
			EgressVirtualInterface struct {
				ACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"aclAnalysis,omitempty"` //
				ID                  string `json:"id,omitempty"` //
				InterfaceStatistics struct {
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
				} `json:"interfaceStatistics,omitempty"` //
				InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
				InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
				Name                                  string `json:"name,omitempty"`                                  //
				PathOverlayInfo                       []struct {
					ControlPlane            string `json:"controlPlane,omitempty"`            //
					DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
					DestIP                  string `json:"destIp,omitempty"`                  //
					DestPort                string `json:"destPort,omitempty"`                //
					Protocol                string `json:"protocol,omitempty"`                //
					SourceIP                string `json:"sourceIp,omitempty"`                //
					SourcePort              string `json:"sourcePort,omitempty"`              //
					VxlanInfo               struct {
						Dscp string `json:"dscp,omitempty"` //
						Vnid string `json:"vnid,omitempty"` //
					} `json:"vxlanInfo,omitempty"` //
				} `json:"pathOverlayInfo,omitempty"` //
				QosStatistics []struct {
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
				} `json:"qosStatistics,omitempty"` //
				QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
				QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
				UsedVlan                        string `json:"usedVlan,omitempty"`                        //
				VrfName                         string `json:"vrfName,omitempty"`                         //
			} `json:"egressVirtualInterface,omitempty"` //
			FlexConnect struct {
				Authentication    string `json:"authentication,omitempty"` //
				DataSwitching     string `json:"dataSwitching,omitempty"`  //
				EgressACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"egressACLAnalysis,omitempty"` //
				IngressACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"ingressACLAnalysis,omitempty"` //
				WirelessLanControllerID   string `json:"wirelessLanControllerId,omitempty"`   //
				WirelessLanControllerName string `json:"wirelessLanControllerName,omitempty"` //
			} `json:"flexConnect,omitempty"` //
			ID                       string `json:"id,omitempty"` //
			IngressPhysicalInterface struct {
				ACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"aclAnalysis,omitempty"` //
				ID                  string `json:"id,omitempty"` //
				InterfaceStatistics struct {
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
				} `json:"interfaceStatistics,omitempty"` //
				InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
				InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
				Name                                  string `json:"name,omitempty"`                                  //
				PathOverlayInfo                       []struct {
					ControlPlane            string `json:"controlPlane,omitempty"`            //
					DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
					DestIP                  string `json:"destIp,omitempty"`                  //
					DestPort                string `json:"destPort,omitempty"`                //
					Protocol                string `json:"protocol,omitempty"`                //
					SourceIP                string `json:"sourceIp,omitempty"`                //
					SourcePort              string `json:"sourcePort,omitempty"`              //
					VxlanInfo               struct {
						Dscp string `json:"dscp,omitempty"` //
						Vnid string `json:"vnid,omitempty"` //
					} `json:"vxlanInfo,omitempty"` //
				} `json:"pathOverlayInfo,omitempty"` //
				QosStatistics []struct {
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
				} `json:"qosStatistics,omitempty"` //
				QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
				QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
				UsedVlan                        string `json:"usedVlan,omitempty"`                        //
				VrfName                         string `json:"vrfName,omitempty"`                         //
			} `json:"ingressPhysicalInterface,omitempty"` //
			IngressVirtualInterface struct {
				ACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"aclAnalysis,omitempty"` //
				ID                  string `json:"id,omitempty"` //
				InterfaceStatistics struct {
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
				} `json:"interfaceStatistics,omitempty"` //
				InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
				InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
				Name                                  string `json:"name,omitempty"`                                  //
				PathOverlayInfo                       []struct {
					ControlPlane            string `json:"controlPlane,omitempty"`            //
					DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
					DestIP                  string `json:"destIp,omitempty"`                  //
					DestPort                string `json:"destPort,omitempty"`                //
					Protocol                string `json:"protocol,omitempty"`                //
					SourceIP                string `json:"sourceIp,omitempty"`                //
					SourcePort              string `json:"sourcePort,omitempty"`              //
					VxlanInfo               struct {
						Dscp string `json:"dscp,omitempty"` //
						Vnid string `json:"vnid,omitempty"` //
					} `json:"vxlanInfo,omitempty"` //
				} `json:"pathOverlayInfo,omitempty"` //
				QosStatistics []struct {
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
				} `json:"qosStatistics,omitempty"` //
				QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
				QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
				UsedVlan                        string `json:"usedVlan,omitempty"`                        //
				VrfName                         string `json:"vrfName,omitempty"`                         //
			} `json:"ingressVirtualInterface,omitempty"` //
			IP                             string `json:"ip,omitempty"`                             //
			LinkInformationSource          string `json:"linkInformationSource,omitempty"`          //
			Name                           string `json:"name,omitempty"`                           //
			PerfMonCollection              string `json:"perfMonCollection,omitempty"`              //
			PerfMonCollectionFailureReason string `json:"perfMonCollectionFailureReason,omitempty"` //
			PerfMonStatistics              []struct {
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
			} `json:"perfMonStatistics,omitempty"` //
			Role    string   `json:"role,omitempty"`    //
			SSID    string   `json:"ssid,omitempty"`    //
			Tunnels []string `json:"tunnels,omitempty"` //
			Type    string   `json:"type,omitempty"`    //
			WlanID  string   `json:"wlanId,omitempty"`  //
		} `json:"networkElements,omitempty"` //
		NetworkElementsInfo []struct {
			AccuracyList []struct {
				Percent int    `json:"percent,omitempty"` //
				Reason  string `json:"reason,omitempty"`  //
			} `json:"accuracyList,omitempty"` //
			DetailedStatus struct {
				ACLTraceCalculation              string `json:"aclTraceCalculation,omitempty"`              //
				ACLTraceCalculationFailureReason string `json:"aclTraceCalculationFailureReason,omitempty"` //
			} `json:"detailedStatus,omitempty"` //
			DeviceStatistics struct {
				CPUStatistics struct {
					FiveMinUsageInPercentage  int `json:"fiveMinUsageInPercentage,omitempty"`  //
					FiveSecsUsageInPercentage int `json:"fiveSecsUsageInPercentage,omitempty"` //
					OneMinUsageInPercentage   int `json:"oneMinUsageInPercentage,omitempty"`   //
					RefreshedAt               int `json:"refreshedAt,omitempty"`               //
				} `json:"cpuStatistics,omitempty"` //
				MemoryStatistics struct {
					MemoryUsage int `json:"memoryUsage,omitempty"` //
					RefreshedAt int `json:"refreshedAt,omitempty"` //
					TotalMemory int `json:"totalMemory,omitempty"` //
				} `json:"memoryStatistics,omitempty"` //
			} `json:"deviceStatistics,omitempty"` //
			DeviceStatsCollection              string `json:"deviceStatsCollection,omitempty"`              //
			DeviceStatsCollectionFailureReason string `json:"deviceStatsCollectionFailureReason,omitempty"` //
			EgressInterface                    struct {
				PhysicalInterface struct {
					ACLAnalysis struct {
						ACLName      string `json:"aclName,omitempty"` //
						MatchingAces []struct {
							Ace           string `json:"ace,omitempty"` //
							MatchingPorts []struct {
								Ports []struct {
									DestPorts   []string `json:"destPorts,omitempty"`   //
									SourcePorts []string `json:"sourcePorts,omitempty"` //
								} `json:"ports,omitempty"` //
								Protocol string `json:"protocol,omitempty"` //
							} `json:"matchingPorts,omitempty"` //
							Result string `json:"result,omitempty"` //
						} `json:"matchingAces,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"aclAnalysis,omitempty"` //
					ID                  string `json:"id,omitempty"` //
					InterfaceStatistics struct {
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
					} `json:"interfaceStatistics,omitempty"` //
					InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
					InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
					Name                                  string `json:"name,omitempty"`                                  //
					PathOverlayInfo                       []struct {
						ControlPlane            string `json:"controlPlane,omitempty"`            //
						DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
						DestIP                  string `json:"destIp,omitempty"`                  //
						DestPort                string `json:"destPort,omitempty"`                //
						Protocol                string `json:"protocol,omitempty"`                //
						SourceIP                string `json:"sourceIp,omitempty"`                //
						SourcePort              string `json:"sourcePort,omitempty"`              //
						VxlanInfo               struct {
							Dscp string `json:"dscp,omitempty"` //
							Vnid string `json:"vnid,omitempty"` //
						} `json:"vxlanInfo,omitempty"` //
					} `json:"pathOverlayInfo,omitempty"` //
					QosStatistics []struct {
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
					} `json:"qosStatistics,omitempty"` //
					QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
					QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
					UsedVlan                        string `json:"usedVlan,omitempty"`                        //
					VrfName                         string `json:"vrfName,omitempty"`                         //
				} `json:"physicalInterface,omitempty"` //
				VirtualInterface []struct {
					ACLAnalysis struct {
						ACLName      string `json:"aclName,omitempty"` //
						MatchingAces []struct {
							Ace           string `json:"ace,omitempty"` //
							MatchingPorts []struct {
								Ports []struct {
									DestPorts   []string `json:"destPorts,omitempty"`   //
									SourcePorts []string `json:"sourcePorts,omitempty"` //
								} `json:"ports,omitempty"` //
								Protocol string `json:"protocol,omitempty"` //
							} `json:"matchingPorts,omitempty"` //
							Result string `json:"result,omitempty"` //
						} `json:"matchingAces,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"aclAnalysis,omitempty"` //
					ID                  string `json:"id,omitempty"` //
					InterfaceStatistics struct {
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
					} `json:"interfaceStatistics,omitempty"` //
					InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
					InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
					Name                                  string `json:"name,omitempty"`                                  //
					PathOverlayInfo                       []struct {
						ControlPlane            string `json:"controlPlane,omitempty"`            //
						DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
						DestIP                  string `json:"destIp,omitempty"`                  //
						DestPort                string `json:"destPort,omitempty"`                //
						Protocol                string `json:"protocol,omitempty"`                //
						SourceIP                string `json:"sourceIp,omitempty"`                //
						SourcePort              string `json:"sourcePort,omitempty"`              //
						VxlanInfo               struct {
							Dscp string `json:"dscp,omitempty"` //
							Vnid string `json:"vnid,omitempty"` //
						} `json:"vxlanInfo,omitempty"` //
					} `json:"pathOverlayInfo,omitempty"` //
					QosStatistics []struct {
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
					} `json:"qosStatistics,omitempty"` //
					QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
					QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
					UsedVlan                        string `json:"usedVlan,omitempty"`                        //
					VrfName                         string `json:"vrfName,omitempty"`                         //
				} `json:"virtualInterface,omitempty"` //
			} `json:"egressInterface,omitempty"` //
			FlexConnect struct {
				Authentication    string `json:"authentication,omitempty"` //
				DataSwitching     string `json:"dataSwitching,omitempty"`  //
				EgressACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"egressACLAnalysis,omitempty"` //
				IngressACLAnalysis struct {
					ACLName      string `json:"aclName,omitempty"` //
					MatchingAces []struct {
						Ace           string `json:"ace,omitempty"` //
						MatchingPorts []struct {
							Ports []struct {
								DestPorts   []string `json:"destPorts,omitempty"`   //
								SourcePorts []string `json:"sourcePorts,omitempty"` //
							} `json:"ports,omitempty"` //
							Protocol string `json:"protocol,omitempty"` //
						} `json:"matchingPorts,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"matchingAces,omitempty"` //
					Result string `json:"result,omitempty"` //
				} `json:"ingressACLAnalysis,omitempty"` //
				WirelessLanControllerID   string `json:"wirelessLanControllerId,omitempty"`   //
				WirelessLanControllerName string `json:"wirelessLanControllerName,omitempty"` //
			} `json:"flexConnect,omitempty"` //
			ID               string `json:"id,omitempty"` //
			IngressInterface struct {
				PhysicalInterface struct {
					ACLAnalysis struct {
						ACLName      string `json:"aclName,omitempty"` //
						MatchingAces []struct {
							Ace           string `json:"ace,omitempty"` //
							MatchingPorts []struct {
								Ports []struct {
									DestPorts   []string `json:"destPorts,omitempty"`   //
									SourcePorts []string `json:"sourcePorts,omitempty"` //
								} `json:"ports,omitempty"` //
								Protocol string `json:"protocol,omitempty"` //
							} `json:"matchingPorts,omitempty"` //
							Result string `json:"result,omitempty"` //
						} `json:"matchingAces,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"aclAnalysis,omitempty"` //
					ID                  string `json:"id,omitempty"` //
					InterfaceStatistics struct {
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
					} `json:"interfaceStatistics,omitempty"` //
					InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
					InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
					Name                                  string `json:"name,omitempty"`                                  //
					PathOverlayInfo                       []struct {
						ControlPlane            string `json:"controlPlane,omitempty"`            //
						DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
						DestIP                  string `json:"destIp,omitempty"`                  //
						DestPort                string `json:"destPort,omitempty"`                //
						Protocol                string `json:"protocol,omitempty"`                //
						SourceIP                string `json:"sourceIp,omitempty"`                //
						SourcePort              string `json:"sourcePort,omitempty"`              //
						VxlanInfo               struct {
							Dscp string `json:"dscp,omitempty"` //
							Vnid string `json:"vnid,omitempty"` //
						} `json:"vxlanInfo,omitempty"` //
					} `json:"pathOverlayInfo,omitempty"` //
					QosStatistics []struct {
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
					} `json:"qosStatistics,omitempty"` //
					QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
					QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
					UsedVlan                        string `json:"usedVlan,omitempty"`                        //
					VrfName                         string `json:"vrfName,omitempty"`                         //
				} `json:"physicalInterface,omitempty"` //
				VirtualInterface []struct {
					ACLAnalysis struct {
						ACLName      string `json:"aclName,omitempty"` //
						MatchingAces []struct {
							Ace           string `json:"ace,omitempty"` //
							MatchingPorts []struct {
								Ports []struct {
									DestPorts   []string `json:"destPorts,omitempty"`   //
									SourcePorts []string `json:"sourcePorts,omitempty"` //
								} `json:"ports,omitempty"` //
								Protocol string `json:"protocol,omitempty"` //
							} `json:"matchingPorts,omitempty"` //
							Result string `json:"result,omitempty"` //
						} `json:"matchingAces,omitempty"` //
						Result string `json:"result,omitempty"` //
					} `json:"aclAnalysis,omitempty"` //
					ID                  string `json:"id,omitempty"` //
					InterfaceStatistics struct {
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
					} `json:"interfaceStatistics,omitempty"` //
					InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
					InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
					Name                                  string `json:"name,omitempty"`                                  //
					PathOverlayInfo                       []struct {
						ControlPlane            string `json:"controlPlane,omitempty"`            //
						DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
						DestIP                  string `json:"destIp,omitempty"`                  //
						DestPort                string `json:"destPort,omitempty"`                //
						Protocol                string `json:"protocol,omitempty"`                //
						SourceIP                string `json:"sourceIp,omitempty"`                //
						SourcePort              string `json:"sourcePort,omitempty"`              //
						VxlanInfo               struct {
							Dscp string `json:"dscp,omitempty"` //
							Vnid string `json:"vnid,omitempty"` //
						} `json:"vxlanInfo,omitempty"` //
					} `json:"pathOverlayInfo,omitempty"` //
					QosStatistics []struct {
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
					} `json:"qosStatistics,omitempty"` //
					QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
					QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
					UsedVlan                        string `json:"usedVlan,omitempty"`                        //
					VrfName                         string `json:"vrfName,omitempty"`                         //
				} `json:"virtualInterface,omitempty"` //
			} `json:"ingressInterface,omitempty"` //
			IP                             string `json:"ip,omitempty"`                             //
			LinkInformationSource          string `json:"linkInformationSource,omitempty"`          //
			Name                           string `json:"name,omitempty"`                           //
			PerfMonCollection              string `json:"perfMonCollection,omitempty"`              //
			PerfMonCollectionFailureReason string `json:"perfMonCollectionFailureReason,omitempty"` //
			PerfMonitorStatistics          []struct {
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
			} `json:"perfMonitorStatistics,omitempty"` //
			Role    string   `json:"role,omitempty"`    //
			SSID    string   `json:"ssid,omitempty"`    //
			Tunnels []string `json:"tunnels,omitempty"` //
			Type    string   `json:"type,omitempty"`    //
			WlanID  string   `json:"wlanId,omitempty"`  //
		} `json:"networkElementsInfo,omitempty"` //
		Properties []string `json:"properties,omitempty"` //
		Request    struct {
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
		} `json:"request,omitempty"` //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// PerfMonStatistics is the PerfMonStatistics definition
type PerfMonStatistics struct {
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

// PerfMonitorStatistics is the PerfMonitorStatistics definition
type PerfMonitorStatistics struct {
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

// PhysicalInterface is the PhysicalInterface definition
type PhysicalInterface struct {
	ACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"aclAnalysis,omitempty"` //
	ID                  string `json:"id,omitempty"` //
	InterfaceStatistics struct {
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
	} `json:"interfaceStatistics,omitempty"` //
	InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []struct {
		ControlPlane            string `json:"controlPlane,omitempty"`            //
		DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
		DestIP                  string `json:"destIp,omitempty"`                  //
		DestPort                string `json:"destPort,omitempty"`                //
		Protocol                string `json:"protocol,omitempty"`                //
		SourceIP                string `json:"sourceIp,omitempty"`                //
		SourcePort              string `json:"sourcePort,omitempty"`              //
		VxlanInfo               struct {
			Dscp string `json:"dscp,omitempty"` //
			Vnid string `json:"vnid,omitempty"` //
		} `json:"vxlanInfo,omitempty"` //
	} `json:"pathOverlayInfo,omitempty"` //
	QosStatistics []struct {
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
	} `json:"qosStatistics,omitempty"` //
	QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
	QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
	UsedVlan                        string `json:"usedVlan,omitempty"`                        //
	VrfName                         string `json:"vrfName,omitempty"`                         //
}

// Ports is the Ports definition
type Ports struct {
	DestPorts   []string `json:"destPorts,omitempty"`   //
	SourcePorts []string `json:"sourcePorts,omitempty"` //
}

// QosStatistics is the QosStatistics definition
type QosStatistics struct {
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

// Request is the Request definition
type Request struct {
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

// VirtualInterface is the VirtualInterface definition
type VirtualInterface struct {
	ACLAnalysis struct {
		ACLName      string `json:"aclName,omitempty"` //
		MatchingAces []struct {
			Ace           string `json:"ace,omitempty"` //
			MatchingPorts []struct {
				Ports []struct {
					DestPorts   []string `json:"destPorts,omitempty"`   //
					SourcePorts []string `json:"sourcePorts,omitempty"` //
				} `json:"ports,omitempty"` //
				Protocol string `json:"protocol,omitempty"` //
			} `json:"matchingPorts,omitempty"` //
			Result string `json:"result,omitempty"` //
		} `json:"matchingAces,omitempty"` //
		Result string `json:"result,omitempty"` //
	} `json:"aclAnalysis,omitempty"` //
	ID                  string `json:"id,omitempty"` //
	InterfaceStatistics struct {
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
	} `json:"interfaceStatistics,omitempty"` //
	InterfaceStatsCollection              string `json:"interfaceStatsCollection,omitempty"`              //
	InterfaceStatsCollectionFailureReason string `json:"interfaceStatsCollectionFailureReason,omitempty"` //
	Name                                  string `json:"name,omitempty"`                                  //
	PathOverlayInfo                       []struct {
		ControlPlane            string `json:"controlPlane,omitempty"`            //
		DataPacketEncapsulation string `json:"dataPacketEncapsulation,omitempty"` //
		DestIP                  string `json:"destIp,omitempty"`                  //
		DestPort                string `json:"destPort,omitempty"`                //
		Protocol                string `json:"protocol,omitempty"`                //
		SourceIP                string `json:"sourceIp,omitempty"`                //
		SourcePort              string `json:"sourcePort,omitempty"`              //
		VxlanInfo               struct {
			Dscp string `json:"dscp,omitempty"` //
			Vnid string `json:"vnid,omitempty"` //
		} `json:"vxlanInfo,omitempty"` //
	} `json:"pathOverlayInfo,omitempty"` //
	QosStatistics []struct {
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
	} `json:"qosStatistics,omitempty"` //
	QosStatsCollection              string `json:"qosStatsCollection,omitempty"`              //
	QosStatsCollectionFailureReason string `json:"qosStatsCollectionFailureReason,omitempty"` //
	UsedVlan                        string `json:"usedVlan,omitempty"`                        //
	VrfName                         string `json:"vrfName,omitempty"`                         //
}

// VxlanInfo is the VxlanInfo definition
type VxlanInfo struct {
	Dscp string `json:"dscp,omitempty"` //
	Vnid string `json:"vnid,omitempty"` //
}

// DeletesPathtraceByID deletesPathtraceById
/* Deletes a flow analysis request by its id
@param flowAnalysisID Flow analysis request id
*/
func (s *PathTraceService) DeletesPathtraceByID(flowAnalysisID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisID}"
	path = strings.Replace(path, "{"+"flowAnalysisID"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// InitiateANewPathtrace initiateANewPathtrace
/* Initiates a new flow analysis with periodic refresh and stat collection options. Returns a request id and a task id to get results and follow progress.
 */
// func (s *PathTraceService) InitiateANewPathtrace(initiateANewPathtraceRequest *InitiateANewPathtraceRequest) (*FlowAnalysisRequestResultOutput, *resty.Response, error) {

// 	path := "/dna/intent/api/v1/flow-analysis"

// 	response, err := RestyClient.R().
// 		SetBody(initiateANewPathtraceRequest).
// 		SetResult(&FlowAnalysisRequestResultOutput{}).
// 		SetError(&Error{}).
// 		Post(path)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	result := response.Result().(*FlowAnalysisRequestResultOutput)
// 	return result, response, err

// }

// RetrievesPreviousPathtrace retrievesPreviousPathtrace
/* Returns result of a previously requested flow analysis by its Flow Analysis id
@param flowAnalysisID Flow analysis request id
*/
func (s *PathTraceService) RetrievesPreviousPathtrace(flowAnalysisID string) (*PathResponseResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisID}"
	path = strings.Replace(path, "{"+"flowAnalysisID"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)

	response, err := RestyClient.R().
		SetResult(&PathResponseResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*PathResponseResult)
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
	LastUpdateTime  int    `url:"lastUpdateTime,omitempty"`  // Last update time
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
func (s *PathTraceService) RetrivesAllPreviousPathtracesSummary(retrivesAllPreviousPathtracesSummaryQueryParams *RetrivesAllPreviousPathtracesSummaryQueryParams) (*FlowAnalysisListOutput, *resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis"

	queryString, _ := query.Values(retrivesAllPreviousPathtracesSummaryQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&FlowAnalysisListOutput{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*FlowAnalysisListOutput)
	return result, response, err

}
