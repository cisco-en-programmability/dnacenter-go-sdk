package main

import (
	"fmt"
	"io"
	"os"

	dnac "dnacenter-go-sdk/sdk"
)

// Client is DNA Center API client
var client *dnac.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"false", "false")
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open("/Users/wilhelm32/Downloads/dnac_2-2-3-3.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	var r io.Reader
	r = f

	nResponse, _, err := client.SoftwareImageManagementSwim.ImportLocalSoftwareImage(
		&dnac.ImportLocalSoftwareImageQueryParams{},
		&dnac.ImportLocalSoftwareImageMultipartFields{
			File:     r,
			FileName: "dnac_2-2-3-3_v2.zip",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	if nResponse.Response != nil {

		fmt.Println(nResponse.Response.TaskID)
	}

}
