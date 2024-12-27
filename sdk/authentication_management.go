package dnac

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type AuthenticationManagementService service

type ImportCertificateQueryParams struct {
	PkPassword  string   `url:"pkPassword,omitempty"`  //Password for encrypted private key
	ListOfUsers []string `url:"listOfUsers,omitempty"` //Specify whether the certificate will be used for controller ("server"), disaster recovery ("ipsec") or both ("server, ipsec"). If no value is provided, the default value taken will be "server"
}

type ImportCertificateMultipartFields struct {
	PkFileUploadName   string
	PkFileUpload       io.Reader
	CertFileUploadName string
	CertFileUpload     io.Reader
}

type ImportCertificateP12QueryParams struct {
	P12Password string   `url:"p12Password,omitempty"` //The password for PKCS12 certificate bundle
	PkPassword  string   `url:"pkPassword,omitempty"`  //Password for encrypted private key
	ListOfUsers []string `url:"listOfUsers,omitempty"` //Specify whether the certificate will be used for controller ("server"), disaster recovery ("ipsec") or both ("server, ipsec"). If no value is provided, the default value taken will be "server"
}

type ImportCertificateP12MultipartFields struct {
	P12FileUpload     io.Reader
	P12FileUploadName string
}

type ResponseAuthenticationManagementImportCertificate struct {
	Response *ResponseAuthenticationManagementImportCertificateResponse `json:"response,omitempty"` //
	Version  string                                                     `json:"version,omitempty"`  //
}
type ResponseAuthenticationManagementImportCertificateResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseAuthenticationManagementImportCertificateP12 struct {
	Response *ResponseAuthenticationManagementImportCertificateP12Response `json:"response,omitempty"` //
	Version  string                                                        `json:"version,omitempty"`  //
}
type ResponseAuthenticationManagementImportCertificateP12Response struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseAuthenticationManagementAuthenticationAPI struct {
	Token string `json:"Token,omitempty"` // Token
}

//ImportCertificate importCertificate - 2a9e-c8a4-454a-b942
/* This API enables a user to import a PEM certificate and its key for the controller and/or disaster recovery.


@param ImportCertificateQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-certificate
*/
func (s *AuthenticationManagementService) ImportCertificate(ImportCertificateQueryParams *ImportCertificateQueryParams, ImportCertificateMultipartFields *ImportCertificateMultipartFields) (*ResponseAuthenticationManagementImportCertificate, *resty.Response, error) {
	path := "/dna/intent/api/v1/certificate"

	queryString, _ := query.Values(ImportCertificateQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ImportCertificateMultipartFields != nil {
		clientRequest = clientRequest.SetFileReader("pkFileUpload", ImportCertificateMultipartFields.PkFileUploadName, ImportCertificateMultipartFields.PkFileUpload)
		clientRequest = clientRequest.SetFileReader("certFileUpload", ImportCertificateMultipartFields.CertFileUploadName, ImportCertificateMultipartFields.CertFileUpload)
	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseAuthenticationManagementImportCertificate{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportCertificate(ImportCertificateQueryParams, ImportCertificateMultipartFields)
		}

		return nil, response, fmt.Errorf("error with operation ImportCertificate")
	}

	result := response.Result().(*ResponseAuthenticationManagementImportCertificate)
	return result, response, err

}

//ImportCertificateP12 importCertificateP12 - 368e-79cf-4329-b63f
/* This API enables a user to import a PKCS12 certificate bundle for the controller and/or disaster recovery.


@param ImportCertificateP12QueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-certificate-p12
*/
func (s *AuthenticationManagementService) ImportCertificateP12(ImportCertificateP12QueryParams *ImportCertificateP12QueryParams, ImportCertificateP12MultipartFields *ImportCertificateP12MultipartFields) (*ResponseAuthenticationManagementImportCertificateP12, *resty.Response, error) {
	path := "/dna/intent/api/v1/certificate-p12"

	queryString, _ := query.Values(ImportCertificateP12QueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ImportCertificateP12MultipartFields != nil {
		clientRequest = clientRequest.SetFileReader("p12FileUpload", ImportCertificateP12MultipartFields.P12FileUploadName, ImportCertificateP12MultipartFields.P12FileUpload)
	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseAuthenticationManagementImportCertificateP12{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportCertificateP12(ImportCertificateP12QueryParams, ImportCertificateP12MultipartFields)
		}

		return nil, response, fmt.Errorf("error with operation ImportCertificateP12")
	}

	result := response.Result().(*ResponseAuthenticationManagementImportCertificateP12)
	return result, response, err

}

//AuthenticationAPI Authentication API - ac8a-e94c-4e69-a09d
/* API to obtain an access token, which remains valid for 1 hour. The token obtained using this API is required to be set as value to the X-Auth-Token HTTP Header for all API calls to Cisco DNA Center.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!authentication-api
*/
func (s *AuthenticationManagementService) AuthenticationAPI() (*ResponseAuthenticationManagementAuthenticationAPI, *resty.Response, error) {
	path := "/dna/system/api/v1/auth/token"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseAuthenticationManagementAuthenticationAPI{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AuthenticationAPI()
		}

		return nil, response, fmt.Errorf("error with operation AuthenticationApi")
	}

	result := response.Result().(*ResponseAuthenticationManagementAuthenticationAPI)
	return result, response, err

}
