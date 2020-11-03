package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// FileService is the service to communicate with the File API endpoint
type FileService service

// GetListOfAvailableNamespacesResponse is the GetListOfAvailableNamespacesResponse definition
type GetListOfAvailableNamespacesResponse struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// GetListOfFilesResponse is the GetListOfFilesResponse definition
type GetListOfFilesResponse struct {
	Response []struct {
		AttributeInfo  string   `json:"attributeInfo,omitempty"`  //
		DownloadPath   string   `json:"downloadPath,omitempty"`   //
		Encrypted      bool     `json:"encrypted,omitempty"`      //
		FileFormat     string   `json:"fileFormat,omitempty"`     //
		FileSize       string   `json:"fileSize,omitempty"`       //
		ID             string   `json:"id,omitempty"`             //
		Md5Checksum    string   `json:"md5Checksum,omitempty"`    //
		Name           string   `json:"name,omitempty"`           //
		NameSpace      string   `json:"nameSpace,omitempty"`      //
		SftpServerList []string `json:"sftpServerList,omitempty"` //
		Sha1Checksum   string   `json:"sha1Checksum,omitempty"`   //
		TaskID         string   `json:"taskId,omitempty"`         //
	} `json:"response,omitempty"` //
	Version string `json:"version,omitempty"` //
}

// DownloadAFileByFileID downloadAFileByFileId
/* Downloads a file specified by fileId
@param fileID File Identification number
*/
func (s *FileService) DownloadAFileByFileID(fileID string) (string, *resty.Response, error) {

	path := "/dna/intent/api/v1/file/{fileId}"
	path = strings.Replace(path, "{"+"fileId"+"}", fmt.Sprintf("%v", fileID), -1)

	var operationResult string
	response, err := RestyClient.R().
		SetResult(&operationResult).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return "", nil, err
	}
	return operationResult, response, err
}

// GetListOfAvailableNamespaces getListOfAvailableNamespaces
/* Returns list of available namespaces
 */
func (s *FileService) GetListOfAvailableNamespaces() (*GetListOfAvailableNamespacesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/file/namespace"

	response, err := RestyClient.R().
		SetResult(&GetListOfAvailableNamespacesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetListOfAvailableNamespacesResponse)
	return result, response, err
}

// GetListOfFiles getListOfFiles
/* Returns list of files under a specific namespace
@param nameSpace A listing of fileId's
*/
func (s *FileService) GetListOfFiles(nameSpace string) (*GetListOfFilesResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/file/namespace/{nameSpace}"
	path = strings.Replace(path, "{"+"nameSpace"+"}", fmt.Sprintf("%v", nameSpace), -1)

	response, err := RestyClient.R().
		SetResult(&GetListOfFilesResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}
	result := response.Result().(*GetListOfFilesResponse)
	return result, response, err
}
