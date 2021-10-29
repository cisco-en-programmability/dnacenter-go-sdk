package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
)

// Client is DNA Center API client
var client *dnac.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false")
	if err != nil {
		fmt.Println(err)
		return
	}

	getClientEnrichmentDetailsHeaderParams := &dnac.GetClientEnrichmentDetailsHeaderParams{
		EntityType:    "network_user_id",
		EntityValue:   "test",
		IssueCategory: "test",
	}
	nResponse, _, err := client.Clients.GetClientEnrichmentDetails(getClientEnrichmentDetailsHeaderParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse)

}
