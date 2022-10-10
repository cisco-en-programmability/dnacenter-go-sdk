package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"
)

// Client is DNA Center API client
var Client *dnac.Client

func main() {
	fmt.Println("Authenticating...")
	var err error
	Client, err = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
		"devnetuser", "Cisco123!",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing site topology...")
	topology, _, err := Client.Topology.GetSiteTopology()
	if err != nil {
		fmt.Println(err)
	}
	if topology.Response != nil && topology.Response.Sites != nil {
		for id, site := range *topology.Response.Sites {
			fmt.Println("GET:", id, site.ID, site.GroupNameHierarchy)
		}
	}

	fmt.Println("Printing physical topology...")

	getPhysicalTopologyQueryParams := &dnac.GetPhysicalTopologyQueryParams{
		NodeType: "",
	}
	physicalTopology, _, err := Client.Topology.GetPhysicalTopology(getPhysicalTopologyQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if physicalTopology.Response != nil && physicalTopology.Response.Nodes != nil {
		for id, nodes := range *physicalTopology.Response.Nodes {
			fmt.Println("GET:", id, nodes.ID, nodes.IP, nodes.Label, nodes.AdditionalInfo)
		}
	}

	fmt.Println("Printing VLAN Information...")
	vlanInformation, _, err := Client.Topology.GetVLANDetails()
	if err != nil {
		fmt.Println(err)
	}
	for id, name := range vlanInformation.Response {
		fmt.Println("GET:", id, name)
	}

	fmt.Println("Printing VLAN 1 Details ...")
	vlanDetails, _, err := Client.Topology.GetTopologyDetails("1")
	if err != nil {
		fmt.Println(err)
	}
	if vlanDetails.Response != nil && vlanDetails.Response.Links != nil {
		for id, link := range *vlanDetails.Response.Links {
			fmt.Println("GET:", id, link.ID, link.Source, link.Target, link.LinkStatus)
		}
	}

	fmt.Println("Printing Network Health...")
	getOverallNetworkHealthQueryParams := &dnac.GetOverallNetworkHealthQueryParams{
		Timestamp: "",
	}
	networkHealth, _, err := Client.Topology.GetOverallNetworkHealth(getOverallNetworkHealthQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(networkHealth.HealthDistirubution, networkHealth.LatestHealthScore, networkHealth.Response, networkHealth.MonitoredDevices)
}
