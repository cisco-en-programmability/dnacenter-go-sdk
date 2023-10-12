package main

import (
	"fmt"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/v5/sdk"
)

// Client is DNA Center API client
var client *dnac.Client

func main() {
	var err error
	fmt.Println("Authenticating")
	client, err = dnac.NewClientWithOptions("https://192.168.196.2/",
		"altus", "Altus123",
		"true", "false", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	queryParams := &dnac.ApplicationsQueryParams{
		SiteID: "1",
		Offset: 1,
		Limit:  5,
	}

	client.SetAuthToken("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2MTY5Zjc3YzVlMzEyODY1YmMzZGQ1NTQiLCJhdXRoU291cmNlIjoiaW50ZXJuYWwiLCJ0ZW5hbnROYW1lIjoiVE5UMCIsInJvbGVzIjpbIjYxNjhiNzUxZTdhMjcwMWEzN2Q2NDUyOCJdLCJ0ZW5hbnRJZCI6IjYxNjhiNzUwZTdhMjcwMWEzN2Q2NDUyNiIsImV4cCI6MTY5NzEzNzIyMywiaWF0IjoxNjk3MTMzNjIzLCJqdGkiOiJkZGY1OGM4Ny0wYzMxLTQwMDctYTRiMC00Y2FjZTAzNmU1ZGEiLCJ1c2VybmFtZSI6ImFsdHVzIn0.dbJRqAJr27WhxA0f-ycK6vCuk8h-ooeOyk423gQKk8H0pSqwxngZK9UAcjNQHIlBQ9lGDdyb6Ri1eKzXv5md8FwQgEyK8pN27GA5-5rCEKbNC0sNnhu4qbHXH1O5eX_p1SD0HxJgZhcGg_PKel4Ir_OgUH8OoWgnb_rRTvMe5n-Abuaxib7Zrk0JinNNwS04g0_0u84DCT4W-hkn-GDZHFNlDYSUkAVg7unf-cBhycn1FwybEuFEOTkd2hXh9rf9N03GwiNT_ajVnD2YVMW_vhlz-sLAmPxwu59otTM0Ma6p8JRTnKzmT1zFFbSNTj6cvbKYZltKwO8kJ0xcwT8Hh1")

	nResponse, _, err := client.Applications.Applications(queryParams)
	if err != nil {
		fmt.Println(err)
		// return
	}
	if nResponse != nil {
		fmt.Println(nResponse.Response)
	}

	// nResponse, _, err = client.Applications.Applications(queryParams)
	// if err != nil {
	// 	fmt.Println(err)
	// 	// return
	// }
	// if nResponse != nil {
	// 	fmt.Println(nResponse.Response)
	// }

}
