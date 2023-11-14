# dnacenter-go-sdk

dnacenter-go-sdk is a Go client library for [DNA Center Platform](https://developer.cisco.com/dnacenter/).

## Usage

```go
import dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"
```

## Introduction

The dnacenter-go-sdk makes it easier to work with the Cisco DNA Center Platform RESTFul APIs from Go.

It supports version 2.3.5.3, but it is backward compatible with other versions as long as those versions use the same URLs and options as version 2.3.5.3.

## Getting started

The first think you need to do is to generate an API client. There are two options to do it:

1. Parameters
2. Environment variables

### Parameters

The client could be generated with the following parameters:

- `baseURL`: The base URL, FQDN or IP, of the DNA instance.
- `username`: The username for the API authentication and authorization.
- `password`: The password for the API authentication and authorization.
- `debug`: Boolean to enable debugging
- `sslVerify`: Boolean to enable or disable SSL certificate verification.
- `waitTimeToManyRequest`: Time in minutes SDK will wait after API responds with 429 status code (to many requests) it defaults is in 1 minute.

```go
Client, err = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
    "devnetuser", "Cisco123!",
    "false", "false", nil)
devicesCount, _, err := Client.Devices.GetDeviceCount()
```

### Using environment variables

The client can be configured with the following environment variables:

- `DNAC_BASE_URL`: The base URL, FQDN or IP, of the DNA instance.
- `DNAC_USERNAME`: The username for the API authentication and authorization.
- `DNAC_PASSWORD`: The password for the API authentication and authorization.
- `DNAC_DEBUG`: Boolean to enable debugging
- `DNAC_SSL_VERIFY`: Boolean to enable or disable SSL certificate verification.
- `DNAC_WAIT_TIME`: Time in minutes SDK will wait after API responds with 429 status code (to many requests) it defaults is in 1 minute.

```go
Client, err = dnac.NewClient()
devicesCount, _, err := Client.Devices.GetDeviceCount()
```

## Examples

Here is an example of how we can generate a client, get a device count and then a list of devices filtering them using query params.

```go
Client, err = dnac.NewClientWithOptions("https://sandboxdnac.cisco.com",
    "devnetuser", "Cisco123!",
    "false", "false",nil)
devicesCount, _, err := Client.Devices.GetDeviceCount()
if err != nil {
    fmt.Println(err)
}
fmt.Println("Device Count:", devicesCount.Response)
getDeviceListQueryParams = &dnac.GetDeviceListQueryParams{
    PlatformID: []string{"C9300-24UX"},
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
```

## Documentation
### dnacenter-go-sdk/v5
[![PkgGoDev](https://pkg.go.dev/badge/github.com/)](https://pkg.go.dev/github.com/cisco-en-programmability/dnacenter-go-sdk/v5)
### dnacenter-go-sdk/v4
[![PkgGoDev](https://pkg.go.dev/badge/github.com/)](https://pkg.go.dev/github.com/cisco-en-programmability/dnacenter-go-sdk/v4)
### dnacenter-go-sdk/v3
[![PkgGoDev](https://pkg.go.dev/badge/github.com/)](https://pkg.go.dev/github.com/cisco-en-programmability/dnacenter-go-sdk/v3)
### dnacenter-go-sdk/v2
[![PkgGoDev](https://pkg.go.dev/badge/github.com/)](https://pkg.go.dev/github.com/cisco-en-programmability/dnacenter-go-sdk/v2)
### dnacenter-go-sdk
[![PkgGoDev](https://pkg.go.dev/badge/github.com/)](https://pkg.go.dev/github.com/cisco-en-programmability/dnacenter-go-sdk)
## Compatibility matrix

| SDK versions | Cisco DNA Center version supported |
|--------------|------------------------------------|
| 0.y.z        |  2.1.2                             |
| 1.y.z        |  2.1.2                             |
| 2.y.z        |  2.2.2.3                           |
| 3.y.z        |  2.2.3.3                           |
| 4.y.z        |  2.3.3.0                           |
| 5.y.z        |  2.3.5.3                           |

## Changelog

All notable changes to this project will be documented in the [CHANGELOG](https://github.com/cisco-en-programmability/dnacenter-go-sdk/blob/main/CHANGELOG.md) file.

The development team may make additional name changes as the library evolves with the Cisco DNA Center APIs.


## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
