package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
)

// Client is DNA Center API client
var Client *dnac.Client

func main() {
	fmt.Println("Authenticating")
	Client = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
		"devnetuser", "Cisco123!",
		"true", "false")
	fmt.Println("Getting device count")
	devicesCount, _, err := Client.Devices.GetDeviceCount()
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
	for id, device := range devices.Response {
		fmt.Println("GET:", id, device.ID, device.MacAddress, device.ManagementIPAddress, device.PlatformID)
	}

	getDeviceListQueryParams = &dnac.GetDeviceListQueryParams{
		PlatformID: []string{"C9300-24UX"},
	}

	fmt.Println("Printing device list  ... PlatformID is C9300-24UX")
	devices, _, err = Client.Devices.GetDeviceList(getDeviceListQueryParams)
	if err != nil {
		fmt.Println(err)
	}

	for id, device := range devices.Response {
		fmt.Println("GET:", id, device.ID, device.MacAddress, device.ManagementIPAddress, device.PlatformID)
	}

	fmt.Println("Printing device info by device id...")
	device, _, err := Client.Devices.GetDeviceByID(devices.Response[0].ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(device.ID, device.MacAddress, device.ManagementIPAddress, device.PlatformID)

}
