# dnacenter-go-sdk

dnacenter-go-sdk is a Go client library for [DNA Center Platform](https://developer.cisco.com/dnacenter/).

## Usage

```go
import dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
```

## Introduction

The dnacenter-go-sdk makes it easier to work with the Cisco DNA Center Platform RESTFul APIs from Go.

It supports version 2.1.2, but it is backward compatible with other versions as long as those versions use the same URLs and options as version 2.1.2.

## Getting started

The first think you need to do is to generate an API client. There are two options to do it: using parameters or environment variables

### Parameters

The client could be generated with parameters:

```go
Client = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
    "devnetuser", "Cisco123!",
    "false", "false")
devicesCount, _, err := Client.Devices.GetDeviceCount()
```

### Using environment variables

The client can be configured with the following environment variables:

- DNAC_BASE_URL: The base URL, FQDN or IP, of the DNA instance.
- DNAC_USERNAME: The username for the API authentication and authorization.
- DNAC_PASSWORD: The password for the API authentication and authorization.
- DNAC_DEBUG: Boolean to enable debugging
- DNAC_SSL_VERIFY: Boolean to enable or disable SSL certificate verification.

```go
Client = dnac.NewClient()
devicesCount, _, err := Client.Devices.GetDeviceCount()
```

## Examples

Here is an example of how we can generate a client, get a device count and then a list of devices filtering them using query params.

```go
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
