package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v7/sdk"
)

// Client is Catalyst Center API client
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

	queryParams := &dnac.GetGlobalPoolQueryParams{
		Offset: 1,
		Limit:  5,
	}

	nResponse, _, err := client.NetworkSettings.GetGlobalPool(queryParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse.Response)
}
