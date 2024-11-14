package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
)

// client is Catalyst Center API client
var client *dnac.Client

func main() {
	fmt.Println("Authenticating...")
	var err error
	client, err = dnac.NewClientWithOptions("https://100.119.103.218",
		"cloverhound_user", "LABchsys!23$",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// now := time.Now() // current local time
	// sec := now.UnixNano()
	vStack := false
	request1 := dnac.RequestDeviceOnboardingPnpAddDevice{
		DeviceInfo: &dnac.RequestDeviceOnboardingPnpAddDeviceDeviceInfo{
			SerialNumber: "FLM2213W05S",
			Stack:        &vStack,
			SudiRequired: &vStack,
			Hostname:     "FLM2213W05W",
		},
	}
	resp1, restyResp1, err := client.DeviceOnboardingPnp.AddDevice(&request1)
	if err != nil {
		fmt.Println(err)
	}
	print(resp1)
	print(restyResp1)

	// if siteHealth.Response != nil {
	// 	for id, site := range *siteHealth.Response {
	// 		fmt.Println(fmt.Sprintf("Site --> ID: %d, Name: %s, Health: %d", id, site.SiteName, site.NetworkHealthAverage))
	// 	}
	// }

}
