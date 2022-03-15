package main

import (
	"fmt"

	dnac "dnacenter-go-sdk/sdk"
)

// Client is DNA Center API client
var Client *dnac.Client

func main() {
	fmt.Println("Authenticating")
	var err error
	Client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Getting task")
	resourcePath := "/api/v1/image/task"
	query := make(map[string]string)
	query["taskUuid"] = "2bdff4eb-a05e-494e-9715-6ebb3d874ad2"
	resty, err := Client.CustomCall.GetCustomCall(resourcePath, &query)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resty.String())

}
