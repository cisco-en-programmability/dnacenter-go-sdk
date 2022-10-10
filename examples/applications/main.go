package main

import (
	"fmt"

	dnac "dnacenter-go-sdk/sdk"
)

// Client is DNA Center API client
var client *dnac.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	queryParams := &dnac.ApplicationsQueryParams{
		SiteID: "1",
	}

	nResponse, _, err := client.Applications.Applications(queryParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse.Response)
}
