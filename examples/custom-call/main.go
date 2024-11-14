package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v6/sdk"
)

// Client is Catalyst Center API client
var Client *dnac.Client

func main() {
	fmt.Println("Authenticating")
	var err error
	Client, err = dnac.NewClientWithOptions("https://100.119.103.190",
		"cloverhound_user", "LABchsys!23$",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resourcePath := fmt.Sprintf("%s/api/v1/siteprofile", Client.RestyClient().BaseURL)
	a, err := Client.CustomCall.GetCustomCall(resourcePath, nil)

	print(a.String())

}
