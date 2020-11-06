package main

import (
	"fmt"
	"strconv"
	"time"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
)

// client is DNA Center API client
var client *dnac.Client

func main() {
	fmt.Println("Authenticating...")
	client = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"false", "false")

	fmt.Println("Printing site health...")
	now := time.Now() // current local time
	sec := now.UnixNano()
	getSiteHealthQueryParams := &dnac.GetSiteHealthQueryParams{
		Timestamp: strconv.Itoa(int(sec) / 1000000),
	}
	siteHealth, _, err := client.Sites.GetSiteHealth(getSiteHealthQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	for id, site := range siteHealth.Response {
		fmt.Println(fmt.Sprintf("Site --> ID: %d, Name: %s, Health: %d", id, site.SiteName, site.NetworkHealthAverage))
	}

	getOverallNetworkHealthQueryParams := &dnac.GetOverallNetworkHealthQueryParams{
		Timestamp: strconv.Itoa(int(sec) / 1000000),
	}

	networkHealth, _, err := client.Topology.GetOverallNetworkHealth(getOverallNetworkHealthQueryParams)
	if err != nil {
		fmt.Println(err)
	}

	for _, network := range networkHealth.Response {
		fmt.Println(fmt.Sprintf("Network Health --> Good Count: %d, Bad Count: %d, Health Score: %d", network.GoodCount, network.BadCount, network.HealthScore))
	}
}
