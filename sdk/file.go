package dnac

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type FileService service

type UploadFileMultipartFields struct {
	File     io.Reader
	FileName string
}

type ResponseFileGetListOfAvailableNamespaces struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}
type ResponseFileGetListOfFiles struct {
	Response *[]ResponseFileGetListOfFilesResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  //
}
type ResponseFileGetListOfFilesResponse struct {
	AttributeInfo  *ResponseFileGetListOfFilesResponseAttributeInfo    `json:"attributeInfo,omitempty"`  //
	DownloadPath   string                                              `json:"downloadPath,omitempty"`   //
	Encrypted      *bool                                               `json:"encrypted,omitempty"`      //
	FileFormat     string                                              `json:"fileFormat,omitempty"`     //
	FileSize       string                                              `json:"fileSize,omitempty"`       //
	ID             string                                              `json:"id,omitempty"`             //
	Md5Checksum    string                                              `json:"md5Checksum,omitempty"`    //
	Name           string                                              `json:"name,omitempty"`           //
	NameSpace      string                                              `json:"nameSpace,omitempty"`      //
	SftpServerList *[]ResponseFileGetListOfFilesResponseSftpServerList `json:"sftpServerList,omitempty"` //
	Sha1Checksum   string                                              `json:"sha1Checksum,omitempty"`   //
	TaskID         string                                              `json:"taskId,omitempty"`         //
}
type ResponseFileGetListOfFilesResponseAttributeInfo interface{}
type ResponseFileGetListOfFilesResponseSftpServerList interface{}
type ResponseFileUploadFile interface{}

//GetListOfAvailableNamespaces Get list of available namespaces - 3f89-bbfc-4f6b-8b50
/* Returns list of available namespaces



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-available-namespaces
*/
func (s *FileService) GetListOfAvailableNamespaces() (*ResponseFileGetListOfAvailableNamespaces, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/namespace"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFileGetListOfAvailableNamespaces{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfAvailableNamespaces()
		}
		return nil, response, fmt.Errorf("error with operation GetListOfAvailableNamespaces")
	}

	result := response.Result().(*ResponseFileGetListOfAvailableNamespaces)
	return result, response, err

}

//GetListOfFiles Get list of files - 42b6-a86e-44b8-bdfc
/* Returns list of files under a specific namespace


@param nameSpace nameSpace path parameter. A listing of fileId's


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-list-of-files
*/
func (s *FileService) GetListOfFiles(nameSpace string) (*ResponseFileGetListOfFiles, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/namespace/{nameSpace}"
	path = strings.Replace(path, "{nameSpace}", fmt.Sprintf("%v", nameSpace), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseFileGetListOfFiles{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetListOfFiles(nameSpace)
		}
		return nil, response, fmt.Errorf("error with operation GetListOfFiles")
	}

	result := response.Result().(*ResponseFileGetListOfFiles)
	return result, response, err

}

//DownloadAFileByFileID Download a file by fileId - 9698-c8ec-4a0b-8c1a
/* Downloads a file specified by fileId


@param fileID fileId path parameter. File Identification number


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-a-file-by-file-id
*/
func (s *FileService) DownloadAFileByFileID(fileID string) (FileDownload, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/{fileId}"
	path = strings.Replace(path, "{fileId}", fmt.Sprintf("%v", fileID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Get(path)

	fdownload := FileDownload{}
	if err != nil {
		return fdownload, nil, err
	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadAFileByFileID(fileID)
		}
		return fdownload, response, fmt.Errorf("error with operation ExportTrustedCertificate")
	}

	fdownload.FileData = response.Body()
	headerVal := response.Header()["Content-Disposition"][0]
	fname := strings.SplitAfter(headerVal, "=")
	fdownload.FileName = strings.ReplaceAll(fname[1], "\"", "")

	return fdownload, response, err

}

//UploadFile uploadFile - 15bf-fb0f-44c8-98f2
/* Uploads a new file within a specific nameSpace


@param nameSpace nameSpace path parameter.

Documentation Link: https://developer.cisco.com/docs/dna-center/#!upload-file
*/
func (s *FileService) UploadFile(nameSpace string, UploadFileMultipartFields *UploadFileMultipartFields) (*ResponseFileUploadFile, *resty.Response, error) {
	path := "/dna/intent/api/v1/file/{nameSpace}"
	path = strings.Replace(path, "{nameSpace}", fmt.Sprintf("%v", nameSpace), -1)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if UploadFileMultipartFields != nil {
		clientRequest = clientRequest.SetFileReader("file", UploadFileMultipartFields.FileName, UploadFileMultipartFields.File)
	}

	response, err = clientRequest.

		// SetResult(&ResponseFileUploadFile{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UploadFile(nameSpace, UploadFileMultipartFields)
		}

		return nil, response, fmt.Errorf("error with operation UploadFile")
	}

	result := response.Result().(ResponseFileUploadFile)

	return &result, response, err

}
