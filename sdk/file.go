package dnac

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// FileService is the service to communicate with the File API endpoint
type FileService service

// FileObjectListResult is the FileObjectListResult definition
type FileObjectListResult struct {
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

// NameSpaceListResult is the NameSpaceListResult definition
type NameSpaceListResult struct {
	Response []string `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DownloadAFileByFileID downloadAFileByFileId
/* Downloads a file specified by fileID
@param fileID File Identification int
*/
func (s *FileService) DownloadAFileByFileID(fileID string) (*resty.Response, error) {

	path := "/dna/intent/api/v1/file/{fileID}"
	path = strings.Replace(path, "{"+"fileID"+"}", fmt.Sprintf("%v", fileID), -1)

	response, err := RestyClient.R().
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// GetListOfAvailableNamespaces getListOfAvailableNamespaces
/* Returns list of available namespaces
 */
func (s *FileService) GetListOfAvailableNamespaces() (*NameSpaceListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/file/namespace"

	response, err := RestyClient.R().
		SetResult(&NameSpaceListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*NameSpaceListResult)
	return result, response, err

}

// GetListOfFiles getListOfFiles
/* Returns list of files under a specific namespace
@param nameSpace A listing of fileID's
*/
func (s *FileService) GetListOfFiles(nameSpace string) (*FileObjectListResult, *resty.Response, error) {

	path := "/dna/intent/api/v1/file/namespace/{nameSpace}"
	path = strings.Replace(path, "{"+"nameSpace"+"}", fmt.Sprintf("%v", nameSpace), -1)

	response, err := RestyClient.R().
		SetResult(&FileObjectListResult{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*FileObjectListResult)
	return result, response, err

}
