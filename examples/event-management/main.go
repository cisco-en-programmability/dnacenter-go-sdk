package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v8/sdk"
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
	queryParams1 := dnac.GetSyslogEventSubscriptionsQueryParams{}
	queryParams1.Name = "Test Terraform 2"
	nResponse, _, err := client.EventManagement.GetSyslogEventSubscriptions(&queryParams1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse)

}
