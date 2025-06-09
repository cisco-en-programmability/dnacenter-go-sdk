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
	nResponse, _, err := client.Task.GetBusinessAPIExecutionDetails("a919fe4c-70c2-4023-a063-404e2705c277")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nResponse)

}
