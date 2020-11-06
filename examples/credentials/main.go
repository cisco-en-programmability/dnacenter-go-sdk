package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"
)

// client is DNA Center API client
var client *dnac.Client

func main() {
	fmt.Println("Authenticating...")
	client = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"false", "false")

	fmt.Println("Creating new SNMPv3 credentials...")
	snmpv3Credentials := &[]dnac.CreateSNMPv3CredentialsRequest{
		dnac.CreateSNMPv3CredentialsRequest{
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
	for id, credential := range credentialsListResponse.Response {
		fmt.Println("GET: ", id, credential.ID, credential.Description, credential.CredentialType)
		client.Discovery.DeleteGlobalCredentialsByID(credential.ID)
	}

	fmt.Println("Creating new HTTP Write credentials...")
	httpWriteCredentials := &[]dnac.CreateHTTPWriteCredentialsRequest{
		dnac.CreateHTTPWriteCredentialsRequest{
			Comments:    "DNA Center HTTP Credentials",
			Description: "HTTP Creds",
			Password:    "HTTP-cr3d$",
			Port:        443,
			Username:    "dnac-http-user",
			Secure:      true,
		},
	}

	httpWriteCredentialsResponse, _, err := client.Discovery.CreateHTTPWriteCredentials(httpWriteCredentials)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(httpWriteCredentialsResponse)

	fmt.Println("Printing SNMPv3 credentials...")
	getGlobalCredentialsQueryParams = &dnac.GetGlobalCredentialsQueryParams{
		CredentialSubType: "HTTP_WRITE",
	}
	credentialsListResponse, _, err = client.Discovery.GetGlobalCredentials(getGlobalCredentialsQueryParams)
	if err != nil {
		fmt.Println(err)
	}
	for id, credential := range credentialsListResponse.Response {
		fmt.Println("GET: ", id, credential.ID, credential.Description, credential.CredentialType)
		client.Discovery.DeleteGlobalCredentialsByID(credential.ID)
	}

}
