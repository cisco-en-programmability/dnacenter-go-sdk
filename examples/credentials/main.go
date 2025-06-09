package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v8/sdk"
)

// client is Catalyst Center API client
var client *dnac.Client

func main() {
	fmt.Println("Authenticating...")
	var err error
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"false", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Creating new SNMPv3 credentials...")
	snmpv3Credentials := &dnac.RequestDiscoveryCreateSNMPv3Credentials{
		dnac.RequestItemDiscoveryCreateSNMPv3Credentials{
			AuthType:        "SHA",
			AuthPassword:    "DNAC-2020",
			SNMPMode:        "AUTHPRIV",
			Username:        "dnac-guide",
			PrivacyType:     "AES128",
			PrivacyPassword: "DNAC-PRIV-2020",
		},
	}

	snmpv3CredentialsResponse, _, err := client.Discovery.CreateSNMPv3Credentials(snmpv3Credentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(snmpv3CredentialsResponse)

	fmt.Println("Printing SNMPv3 credentials...")
	getGlobalCredentialsQueryParams := &dnac.GetGlobalCredentialsQueryParams{
		CredentialSubType: "SNMPV3",
	}
	credentialsListResponse, _, err := client.Discovery.GetGlobalCredentials(getGlobalCredentialsQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if credentialsListResponse.Response != nil {
		for id, credential := range *credentialsListResponse.Response {
			fmt.Println("GET: ", id, credential.ID, credential.Description, credential.CredentialType)
			_, _, err := client.Discovery.DeleteGlobalCredentialsByID(credential.ID)
			if err != nil {
				continue
			}
		}
	}

	fmt.Println("Creating new HTTP Write credentials...")
	port := 443
	secure := true
	httpWriteCredentials := &dnac.RequestDiscoveryCreateHTTPWriteCredentials{
		dnac.RequestItemDiscoveryCreateHTTPWriteCredentials{
			Comments:    "Catalyst Center HTTP Credentials",
			Description: "HTTP Creds",
			Password:    "HTTP-cr3d$",
			Port:        &port,
			Username:    "dnac-http-user",
			Secure:      &secure,
		},
	}

	httpWriteCredentialsResponse, _, err := client.Discovery.CreateHTTPWriteCredentials(httpWriteCredentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(httpWriteCredentialsResponse)

	fmt.Println("Printing HTTP Write credentials...")
	getGlobalCredentialsQueryParams = &dnac.GetGlobalCredentialsQueryParams{
		CredentialSubType: "HTTP_WRITE",
	}
	credentialsListResponse, _, err = client.Discovery.GetGlobalCredentials(getGlobalCredentialsQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	if credentialsListResponse.Response != nil {
		for id, credential := range *credentialsListResponse.Response {
			fmt.Println("GET: ", id, credential.ID, credential.Description, credential.CredentialType)
			_, _, err := client.Discovery.DeleteGlobalCredentialsByID(credential.ID)
			if err != nil {
				continue
			}
		}
	}

}
