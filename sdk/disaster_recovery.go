package dnac

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type DisasterRecoveryService service

type ResponseDisasterRecoveryDisasterRecoveryOperationalStatus struct {
	Severity       string                                                               `json:"severity,omitempty"`       // Severity of the DR Event.
	Status         string                                                               `json:"status,omitempty"`         // Status of the DR Event.
	InitiatedBy    string                                                               `json:"initiated_by,omitempty"`   // Who initiated this event. Is it a system triggered one or user triggered one.
	IPconfig       *[]ResponseDisasterRecoveryDisasterRecoveryOperationalStatusIPconfig `json:"ipconfig,omitempty"`       //
	Tasks          *[]ResponseDisasterRecoveryDisasterRecoveryOperationalStatusTasks    `json:"tasks,omitempty"`          //
	Title          string                                                               `json:"title,omitempty"`          // DR Event Summary
	Site           string                                                               `json:"site,omitempty"`           // Site of the DR in which this event occurred.
	StartTimestamp string                                                               `json:"startTimestamp,omitempty"` // Starting timestamp of the DR event
	Message        string                                                               `json:"message,omitempty"`        // Detailed Description about the DR event
	EndTimestamp   string                                                               `json:"endTimestamp,omitempty"`   // End timestamp of the DR event
}
type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       string `json:"vip,omitempty"`       // Is this interface a Virtual IP address or not. This is true for Site VIP
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusTasks struct {
	Status         string                                                                    `json:"status,omitempty"`         // Status of the DR event.
	IPconfig       *[]ResponseDisasterRecoveryDisasterRecoveryOperationalStatusTasksIPconfig `json:"ipconfig,omitempty"`       //
	Title          string                                                                    `json:"title,omitempty"`          // DR Event Summary
	Site           string                                                                    `json:"site,omitempty"`           // Site of the DR in which this event occured
	StartTimestamp string                                                                    `json:"startTimestamp,omitempty"` // Starting timestamp of the DR event
	Message        string                                                                    `json:"message,omitempty"`        // Detailed description about the DR event
	EndTimestamp   string                                                                    `json:"endTimestamp,omitempty"`   // End timestamp of the DR event
}
type ResponseDisasterRecoveryDisasterRecoveryOperationalStatusTasksIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       string `json:"vip,omitempty"`       // Is this interface a Virtual IP address or not. This is true for Site VIP
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatus struct {
	IPconfig    *[]ResponseDisasterRecoveryDisasterRecoveryStatusIPconfig    `json:"ipconfig,omitempty"`     //
	Site        string                                                       `json:"site,omitempty"`         // Site of the disaster recovery system.
	Main        *ResponseDisasterRecoveryDisasterRecoveryStatusMain          `json:"main,omitempty"`         //
	Recovery    *ResponseDisasterRecoveryDisasterRecoveryStatusRecovery      `json:"recovery,omitempty"`     //
	Witness     *ResponseDisasterRecoveryDisasterRecoveryStatusWitness       `json:"witness,omitempty"`      //
	State       string                                                       `json:"state,omitempty"`        // State of the Disaster Recovery System.
	IPsecTunnel *[]ResponseDisasterRecoveryDisasterRecoveryStatusIPsecTunnel `json:"ipsec-tunnel,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is true for Global DR VIP
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusMain struct {
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusMainIPconfig `json:"ipconfig,omitempty"` //
	State    string                                                        `json:"state,omitempty"`    // State of the Main Site.
	Nodes    *[]ResponseDisasterRecoveryDisasterRecoveryStatusMainNodes    `json:"nodes,omitempty"`    //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusMainIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is true for cluster level.
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusMainNodes struct {
	Hostname    string                                                                `json:"hostname,omitempty"`    // Hostname of the node
	State       string                                                                `json:"state,omitempty"`       // State of the node
	IPaddresses *[]ResponseDisasterRecoveryDisasterRecoveryStatusMainNodesIPaddresses `json:"ipaddresses,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusMainNodesIPaddresses struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface a Virtual IP address or not. This is false for node level.
	IP        string `json:"ip,omitempty"`        // Node IP address
}
type ResponseDisasterRecoveryDisasterRecoveryStatusRecovery struct {
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryIPconfig `json:"ipconfig,omitempty"` //
	State    string                                                            `json:"state,omitempty"`    // State of the Recovery site
	Nodes    *[]ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryNodes    `json:"nodes,omitempty"`    //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is true for cluster level.
	IP        string `json:"ip,omitempty"`        // This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryNodes struct {
	Hostname string                                                                 `json:"hostname,omitempty"` // Hostname of the node
	State    string                                                                 `json:"state,omitempty"`    // State of the node
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryNodesIPconfig `json:"ipconfig,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusRecoveryNodesIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface a Virtual IP Address or not. This is false for node level.
	IP        string `json:"ip,omitempty"`        // Node IP Address
}
type ResponseDisasterRecoveryDisasterRecoveryStatusWitness struct {
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusWitnessIPconfig `json:"ipconfig,omitempty"` //
	State    string                                                           `json:"state,omitempty"`    // State of the Witness Site
	Nodes    *[]ResponseDisasterRecoveryDisasterRecoveryStatusWitnessNodes    `json:"nodes,omitempty"`    //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusWitnessIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is false for witness.
	IP        string `json:"ip,omitempty"`        // In case of witness, this is only an IP.
}
type ResponseDisasterRecoveryDisasterRecoveryStatusWitnessNodes struct {
	Hostname string                                                                `json:"hostname,omitempty"` // Hostname of the witness node
	State    string                                                                `json:"state,omitempty"`    // State of the node
	IPconfig *[]ResponseDisasterRecoveryDisasterRecoveryStatusWitnessNodesIPconfig `json:"ipconfig,omitempty"` //
}
type ResponseDisasterRecoveryDisasterRecoveryStatusWitnessNodesIPconfig struct {
	Interface string `json:"interface,omitempty"` // Enterprise or Management interface
	Vip       *bool  `json:"vip,omitempty"`       // Is this interface an Virtual IP address or not. This is false for Witness
	IP        string `json:"ip,omitempty"`        // In case of witness, this is only an IP
}
type ResponseDisasterRecoveryDisasterRecoveryStatusIPsecTunnel struct {
	SideA  string `json:"side_a,omitempty"` // A Side of the IPSec Tunnel
	SideB  string `json:"side_b,omitempty"` // Other Side of the IPSec Tunnel
	Status string `json:"status,omitempty"` // Status of this IPSec Tunnel
}

//DisasterRecoveryOperationalStatus Disaster Recovery Operational Status - b89c-dbd5-45da-a6e5
/* Returns the status of Disaster Recovery operation performed on the system.


 */
func (s *DisasterRecoveryService) DisasterRecoveryOperationalStatus() (*ResponseDisasterRecoveryDisasterRecoveryOperationalStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/disasterrecovery/system/operationstatus"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDisasterRecoveryDisasterRecoveryOperationalStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DisasterRecoveryOperationalStatus")
	}

	result := response.Result().(*ResponseDisasterRecoveryDisasterRecoveryOperationalStatus)
	return result, response, err

}

//DisasterRecoveryStatus Disaster Recovery Status - 0b83-9ba0-493a-bf2e
/* Detailed and Summarized status of DR components (Active, Standby and Witness system's health).


 */
func (s *DisasterRecoveryService) DisasterRecoveryStatus() (*ResponseDisasterRecoveryDisasterRecoveryStatus, *resty.Response, error) {
	path := "/dna/intent/api/v1/disasterrecovery/system/status"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseDisasterRecoveryDisasterRecoveryStatus{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation DisasterRecoveryStatus")
	}

	result := response.Result().(*ResponseDisasterRecoveryDisasterRecoveryStatus)
	return result, response, err

}
