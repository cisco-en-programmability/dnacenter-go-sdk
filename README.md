# dnacenter-go-sdk

dnacenter-go-sdk is a Go client library for [DNA Center Platform](https://developer.cisco.com/dnacenter/).

## Usage

```go
import dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
Client = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
    "devnetuser", "Cisco123!",
    "false", "false")
devicesCount, _, err := Client.Devices.GetDeviceCount()
if err != nil {
    fmt.Println(err)
}
fmt.Println('Device Count:', devicesCount.Response)
getDeviceListQueryParams = &dnac.GetDeviceListQueryParams{
    PlatformID: []string{"C9300-24UX"},
}
fmt.Println("Printing device list  ... PlatformID is C9300-24UX")
devices, _, err = Client.Devices.GetDeviceList(getDeviceListQueryParams)
if err != nil {
    fmt.Println(err)
}
for id, device := range devices.Response {
    fmt.Println("GET:", id, device.ID, device.ManagementIPAddress, device.PlatformID)
}
```

## Documentation

https://godoc.org/github.com/cisco-en-programmability/dnacenter-go-sdk/sdk

## TODO

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
