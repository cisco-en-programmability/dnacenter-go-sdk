package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v4/sdk"
)

// Client is DNA Center API client
var Client *dnac.Client

func main() {
	fmt.Println("Authenticating")
	var err error
	Client, err = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
		"devnetuser", "Cisco123!",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Getting device count")
	devicesCount, _, err := Client.Devices.GetDeviceCount2()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(devicesCount.Response)

	fmt.Println("Printing device list ...")

	getDeviceListQueryParams := &dnac.GetDeviceListQueryParams{}
	devices, _, err := Client.Devices.GetDeviceList(getDeviceListQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if devices.Response != nil {
		for id, device := range *devices.Response {
			fmt.Println("GET:", id, device.ID, device.MacAddress, device.ManagementIPAddress, device.PlatformID)
		}
	}

	getDeviceListQueryParams = &dnac.GetDeviceListQueryParams{
		//PlatformID: []string{"C9300-24UX"},
	}

	fmt.Println("Printing device list  ... PlatformID is C9300-24UX")
	devices, _, err = Client.Devices.GetDeviceList(getDeviceListQueryParams)
	if err != nil {
		fmt.Println(err)
	}

	if devices.Response != nil {
		for id, device := range *devices.Response {
			fmt.Println("GET:", id, device.ID, device.MacAddress, device.ManagementIPAddress, device.PlatformID)
		}
	}

	fmt.Println("Printing device info by device id...")
	if devices.Response != nil {
		device, _, err := Client.Devices.GetDeviceByID((*devices.Response)[0].ID)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(device.Response.ID, device.Response.MacAddress, device.Response.ManagementIPAddress, device.Response.PlatformID)
	}

}
