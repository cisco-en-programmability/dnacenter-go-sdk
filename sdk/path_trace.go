package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// PathTraceService is the service to communicate with the PathTrace API endpoint
type PathTraceService service

// InitiateANewPathtraceRequest is the InitiateANewPathtraceRequest definition
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

// DeletesPathtraceByIDResponse is the DeletesPathtraceByIdResponse definition
type DeletesPathtraceByIDResponse struct {
	Response struct {
		TaskID string `json:"taskId,omitempty"` //
		URL    string `json:"url,omitempty"`    //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// InitiateANewPathtraceResponse is the InitiateANewPathtraceResponse definition
type InitiateANewPathtraceResponse struct {
	Response struct {
		FlowAnalysisID string `json:"flowAnalysisId,omitempty"` //
		TaskID         string `json:"taskId,omitempty"`         //
		URL            string `json:"url,omitempty"`            //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// RetrievesPreviousPathtraceResponse is the RetrievesPreviousPathtraceResponse definition
type RetrievesPreviousPathtraceResponse struct {
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
				UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
				UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
				} `json:"egressAclAnalysis,omitempty"` //
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
				} `json:"ingressAclAnalysis,omitempty"` //
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
				UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
				UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
			WLANID  string   `json:"wlanId,omitempty"`  //
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
					UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
					UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
				} `json:"egressAclAnalysis,omitempty"` //
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
				} `json:"ingressAclAnalysis,omitempty"` //
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
					UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
					UsedVLAN                        string `json:"usedVlan,omitempty"`                        //
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
			WLANID  string   `json:"wlanId,omitempty"`  //
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

// RetrivesAllPreviousPathtracesSummaryResponse is the RetrivesAllPreviousPathtracesSummaryResponse definition
type RetrivesAllPreviousPathtracesSummaryResponse struct {
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

// DeletesPathtraceByID deletesPathtraceById
/* Deletes a flow analysis request by its id
@param flowAnalysisID Flow analysis request id
*/
func (s *PathTraceService) DeletesPathtraceByID(flowAnalysisID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/flow-analysis/{flowAnalysisId}"
	path = strings.Replace(path, "{"+"flowAnalysisId"+"}", fmt.Sprintf("%v", flowAnalysisID), -1)

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
	result := response.Result().(*RetrivesAllPreviousPathtracesSummaryResponse)
	return result, response, err
}
