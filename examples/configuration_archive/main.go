package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"
)

// client is DNA Center API client
var client *dnac.Client

func main() {
	var err error
	fmt.Println("Authenticating...")
	deviceUUID := ""
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Printing ComplianceDetails")
	getComplianceDetailQueryParams := &dnac.GetComplianceDetailQueryParams{}
	respComplianceDetail, _, err := client.Compliance.GetComplianceDetail(getComplianceDetailQueryParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	if respComplianceDetail != nil {
		fmt.Println(respComplianceDetail)
		deviceUUID = (*respComplianceDetail.Response)[0].DeviceUUID
	} else {
		fmt.Println("There is no data on response")
		return
	}

	fmt.Println("Post ConfArchive")
	reqBody := &dnac.RequestConfigurationArchiveExportDeviceConfigurations{
		DeviceID: []string{deviceUUID},
		Password: "C1sco123!",
	}

	resp, _, err := client.ConfigurationArchive.ExportDeviceConfigurations(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp != nil {
		fmt.Println(resp)
	} else {
		fmt.Println("There is no data on response")
		return
	}

}
