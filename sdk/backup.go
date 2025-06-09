package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type BackupService service

type GetBackupAndRestoreExecutionsQueryParams struct {
	BackupID string `url:"backupId,omitempty"` //The `backupId` of the backup execution to be retrieved.Obtain the `backupId` from the id attribute in the response of the `/dna/system/api/v1/backups` API.
	JobType  string `url:"jobType,omitempty"`  //The `jobType` of the backup execution to be retrieved.
	Status   string `url:"status,omitempty"`   //The `status` of the backup execution to be retrieved.
	Offset   int    `url:"offset,omitempty"`   //The first record to show for this page.
	Limit    int    `url:"limit,omitempty"`    //The number of records to show for this page.
	SortBy   string `url:"sortBy,omitempty"`   //A property within the response to sort by.
	Order    string `url:"order,omitempty"`    //Whether ascending or descending order should be used to sort the response.
}
type GetBackupStoragesQueryParams struct {
	StorageType string `url:"storageType,omitempty"` //The `storageType` of the backup storage to be retrieved.
}
type GetAllBackupQueryParams struct {
	Query  string `url:"query,omitempty"`  //Filter based on the provided text on predefined fields
	Offset int    `url:"offset,omitempty"` //The first record to show for this page.
	Limit  int    `url:"limit,omitempty"`  //The number of records to show for this page.
	SortBy string `url:"sortBy,omitempty"` //A property within the response to sort by.
	Order  string `url:"order,omitempty"`  //Whether ascending or descending order should be used to sort the response.
}

type ResponseBackupGetBackupConfiguration struct {
	DataRetention                   *int   `json:"dataRetention,omitempty"`                   // Date retention policy of the backup
	ID                              string `json:"id,omitempty"`                              // The UUID of the backup configuration
	IsEncryptionPassPhraseAvailable *bool  `json:"isEncryptionPassPhraseAvailable,omitempty"` // Indicates whether an encryption pass phrase is available in the configuration
	MountPath                       string `json:"mountPath,omitempty"`                       // Backup storage mount path
	Type                            string `json:"type,omitempty"`                            // The storage type
}
type ResponseBackupCreateBackupConfiguration interface{}
type ResponseBackupCreateNFSConfiguration interface{}
type ResponseBackupGetAllNFSConfigurations struct {
	Response *[]ResponseBackupGetAllNFSConfigurationsResponse `json:"response,omitempty"` //
	Version  string                                           `json:"version,omitempty"`  // The version of the response
}
type ResponseBackupGetAllNFSConfigurationsResponse struct {
	ID     string                                               `json:"id,omitempty"`     // The UUID of the NFS configuration
	Spec   *ResponseBackupGetAllNFSConfigurationsResponseSpec   `json:"spec,omitempty"`   // The spec of NFS configuration
	Status *ResponseBackupGetAllNFSConfigurationsResponseStatus `json:"status,omitempty"` // The status of NFS configuration
}
type ResponseBackupGetAllNFSConfigurationsResponseSpec interface{}
type ResponseBackupGetAllNFSConfigurationsResponseStatus interface{}
type ResponseBackupDeleteNFSConfiguration interface{}
type ResponseBackupGetBackupAndRestoreExecutions struct {
	Filter   *ResponseBackupGetBackupAndRestoreExecutionsFilter     `json:"filter,omitempty"`   // Records are filtered based on the fields
	Page     *ResponseBackupGetBackupAndRestoreExecutionsPage       `json:"page,omitempty"`     // Contains pagination information.
	Response *[]ResponseBackupGetBackupAndRestoreExecutionsResponse `json:"response,omitempty"` //
	Version  string                                                 `json:"version,omitempty"`  // The version of the response
}
type ResponseBackupGetBackupAndRestoreExecutionsFilter interface{}
type ResponseBackupGetBackupAndRestoreExecutionsPage interface{}
type ResponseBackupGetBackupAndRestoreExecutionsResponse struct {
	Metadata            *ResponseBackupGetBackupAndRestoreExecutionsResponseMetadata         `json:"_metadata,omitempty"`           // This field holds additional information
	BackupID            string                                                               `json:"backupId,omitempty"`            // The UUID of the backup
	CompletedPercentage *float64                                                             `json:"completedPercentage,omitempty"` // Execution completion percentage
	CreatedBy           string                                                               `json:"createdBy,omitempty"`           // This field represents the user who triggered this execution
	Duration            *int                                                                 `json:"duration,omitempty"`            // Total time taken to complete this execution process
	EndDate             string                                                               `json:"endDate,omitempty"`             // The end data and time of the backup execution
	ErrorCode           string                                                               `json:"errorCode,omitempty"`           // Error code of the backup execution process
	ErrorMessage        string                                                               `json:"errorMessage,omitempty"`        // Error message of this backup execution process
	FailedTaskDetail    *ResponseBackupGetBackupAndRestoreExecutionsResponseFailedTaskDetail `json:"failedTaskDetail,omitempty"`    // Failed task detail of this execution
	ID                  string                                                               `json:"id,omitempty"`                  // The UUID of the execution process
	IsForceUpdate       *bool                                                                `json:"isForceUpdate,omitempty"`       // This states whether the execution is forcefully updated or not
	JobType             string                                                               `json:"jobType,omitempty"`             // The backup execution job type
	Scope               string                                                               `json:"scope,omitempty"`               // The backup scope states whether the backup is with assurance or without assurance data
	StartDate           string                                                               `json:"startDate,omitempty"`           // The start data and time of the backup execution
	Status              string                                                               `json:"status,omitempty"`              // The status of the backup execution
	SystemErrorMessage  string                                                               `json:"systemErrorMessage,omitempty"`  // System error message of this backup execution process
	UpdateMessage       string                                                               `json:"updateMessage,omitempty"`       // Force update message
}
type ResponseBackupGetBackupAndRestoreExecutionsResponseMetadata interface{}
type ResponseBackupGetBackupAndRestoreExecutionsResponseFailedTaskDetail interface{}
type ResponseBackupGetBackupAndRestoreExecution struct {
	Response *ResponseBackupGetBackupAndRestoreExecutionResponse `json:"response,omitempty"` // Contains information about backup execution details.
	Version  string                                              `json:"version,omitempty"`  // The version of the response
}
type ResponseBackupGetBackupAndRestoreExecutionResponse interface{}
type ResponseBackupGetBackupStorages struct {
	Response *[]ResponseBackupGetBackupStoragesResponse `json:"response,omitempty"` //
	Version  string                                     `json:"version,omitempty"`  // The version of the response
}
type ResponseBackupGetBackupStoragesResponse struct {
	DiskName      string   `json:"diskName,omitempty"`      // Mounted disk name
	FstypeR       string   `json:"fstype,omitempty"`        // Mounted disk fstype
	Label         string   `json:"label,omitempty"`         // The mounted volume label
	MountPoint    string   `json:"mountPoint,omitempty"`    // Mount path
	PartitionName string   `json:"partitionName,omitempty"` // Mounted disk partitionName
	PercentUsage  *float64 `json:"percentUsage,omitempty"`  // Used percentage
	SizeUnit      string   `json:"sizeUnit,omitempty"`      // Size representation
	TotalSize     *int     `json:"totalSize,omitempty"`     // Total size of the disk in bytes
	UsedSize      *int     `json:"usedSize,omitempty"`      // Used size of the disk in bytes
}
type ResponseBackupGetAllBackup struct {
	Filter   *ResponseBackupGetAllBackupFilter     `json:"filter,omitempty"`   // Records are filtered based on the fields
	Page     *ResponseBackupGetAllBackupPage       `json:"page,omitempty"`     // Contains pagination information.
	Response *[]ResponseBackupGetAllBackupResponse `json:"response,omitempty"` //
	Version  string                                `json:"version,omitempty"`  // The version of the response
}
type ResponseBackupGetAllBackupFilter interface{}
type ResponseBackupGetAllBackupPage interface{}
type ResponseBackupGetAllBackupResponse struct {
	CompatibilityError      *[]ResponseBackupGetAllBackupResponseCompatibilityError `json:"compatibilityError,omitempty"`      //
	Context                 *ResponseBackupGetAllBackupResponseContext              `json:"context,omitempty"`                 // This field states whether the backup is on-demand or schedule based
	CreatedBy               string                                                  `json:"createdBy,omitempty"`               // This field represents the user who triggered this backup
	CreatedDate             string                                                  `json:"createdDate,omitempty"`             // This field represents the backup creation date
	Duration                *int                                                    `json:"duration,omitempty"`                // Total time taken to create this backup
	EndDate                 string                                                  `json:"endDate,omitempty"`                 // This field represents the backup end date
	FipsEnabled             *bool                                                   `json:"fipsEnabled,omitempty"`             // This states whether FIPS is enabled or not
	ID                      string                                                  `json:"id,omitempty"`                      // The UUID for this backup
	InstalledPackages       *[]ResponseBackupGetAllBackupResponseInstalledPackages  `json:"installedPackages,omitempty"`       //
	InternetProtocolVersion string                                                  `json:"internetProtocolVersion,omitempty"` // The network protocol
	IsBackupAvailable       *bool                                                   `json:"isBackupAvailable,omitempty"`       // This field represents whether the backup is available in the configured storage
	IsCompatible            *bool                                                   `json:"isCompatible,omitempty"`            // This field represents whether the backup is compatible to restore
	Name                    string                                                  `json:"name,omitempty"`                    // The backup name
	NumberOfNodes           *int                                                    `json:"numberOfNodes,omitempty"`           // Total number of nodes
	ProductType             string                                                  `json:"productType,omitempty"`             // Installed product type
	ProductVersion          string                                                  `json:"productVersion,omitempty"`          // The product version
	ReleaseDisplayName      string                                                  `json:"releaseDisplayName,omitempty"`      // The release display name of the product installed
	ReleaseDisplayVersion   string                                                  `json:"releaseDisplayVersion,omitempty"`   // The release display version of the product installed
	ReleaseName             string                                                  `json:"releaseName,omitempty"`             // The release name of the product installed
	ReleaseVersion          string                                                  `json:"releaseVersion,omitempty"`          // The release version of the product installed
	Scope                   string                                                  `json:"scope,omitempty"`                   // The backup scope states whether the backup is with assurance or without assurance data
	Size                    *int                                                    `json:"size,omitempty"`                    // The total backup size
	Status                  string                                                  `json:"status,omitempty"`                  // Status of the backup
	Storage                 *ResponseBackupGetAllBackupResponseStorage              `json:"storage,omitempty"`                 // The storage detail of the backup
	Versions                []string                                                `json:"versions,omitempty"`                // List of packages available while creating the backup
}
type ResponseBackupGetAllBackupResponseCompatibilityError struct {
	EndDate     string                                                        `json:"endDate,omitempty"`     // Backup compatibility end date
	Namespace   string                                                        `json:"namespace,omitempty"`   // The namespace of the service
	Response    *ResponseBackupGetAllBackupResponseCompatibilityErrorResponse `json:"response,omitempty"`    // Response received from the service during compatibility check
	ServiceName string                                                        `json:"serviceName,omitempty"` // The service name on which the compatibility check executed
	StartDate   string                                                        `json:"startDate,omitempty"`   // Backup compatibility start date
}
type ResponseBackupGetAllBackupResponseCompatibilityErrorResponse interface{}
type ResponseBackupGetAllBackupResponseContext interface{}
type ResponseBackupGetAllBackupResponseInstalledPackages struct {
	DisplayName string `json:"displayName,omitempty"` // The package display name
	Name        string `json:"name,omitempty"`        // The package name
	Version     string `json:"version,omitempty"`     // The package version
}
type ResponseBackupGetAllBackupResponseStorage interface{}
type ResponseBackupGetBackupByID struct {
	Response *ResponseBackupGetBackupByIDResponse `json:"response,omitempty"` // Contains information about backup.
	Version  string                               `json:"version,omitempty"`  // The version of the response
}
type ResponseBackupGetBackupByIDResponse interface{}
type RequestBackupCreateBackupConfiguration struct {
	DataRetention        *int   `json:"dataRetention,omitempty"`        // Date retention policy of the backup
	EncryptionPassphrase string `json:"encryptionPassphrase,omitempty"` // Password to encrypt the backup information
	MountPath            string `json:"mountPath,omitempty"`            // Backup storage mount path
	Type                 string `json:"type,omitempty"`                 // The storage type
}
type RequestBackupCreateNFSConfiguration struct {
	NfsPort        *int   `json:"nfsPort,omitempty"`        // NFS Port
	NfsVersion     string `json:"nfsVersion,omitempty"`     // NFS version
	PortMapperPort *int   `json:"portMapperPort,omitempty"` // NFS port mapper port
	Server         string `json:"server,omitempty"`         // NFS server host
	SourcePath     string `json:"sourcePath,omitempty"`     // NFS server path
}
type RequestBackupCreateBackup struct {
	Name  string `json:"name,omitempty"`  // The backup name
	Scope string `json:"scope,omitempty"` // The backup scope states whether the backup is with assurance or without assurance data
}

//GetBackupConfiguration Get Backup Configuration - d180-a9a1-48f9-b173
/* This api is used to get the backup configuration



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-backup-configuration
*/
func (s *BackupService) GetBackupConfiguration() (*ResponseBackupGetBackupConfiguration, *resty.Response, error) {
	path := "/dna/system/api/v1/backupConfiguration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseBackupGetBackupConfiguration{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBackupConfiguration()
		}
		return nil, response, fmt.Errorf("error with operation GetBackupConfiguration")
	}

	result := response.Result().(*ResponseBackupGetBackupConfiguration)
	return result, response, err

}

//GetAllNFSConfigurations Get All NFS Configurations - 648e-59e0-41f9-99ec
/* This api is used to get all the configured NFS



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-n-f-s-configurations
*/
func (s *BackupService) GetAllNFSConfigurations() (*ResponseBackupGetAllNFSConfigurations, *resty.Response, error) {
	path := "/dna/system/api/v1/backupNfsConfigurations"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseBackupGetAllNFSConfigurations{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllNFSConfigurations()
		}
		return nil, response, fmt.Errorf("error with operation GetAllNFSConfigurations")
	}

	result := response.Result().(*ResponseBackupGetAllNFSConfigurations)
	return result, response, err

}

//GetBackupAndRestoreExecutions Get Backup And Restore Executions - a5a0-8bd1-4fa8-871f
/* This api is used to get all the backup and restore executions.


@param GetBackupAndRestoreExecutionsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-backup-and-restore-executions
*/
func (s *BackupService) GetBackupAndRestoreExecutions(GetBackupAndRestoreExecutionsQueryParams *GetBackupAndRestoreExecutionsQueryParams) (*ResponseBackupGetBackupAndRestoreExecutions, *resty.Response, error) {
	path := "/dna/system/api/v1/backupRestoreExecutions"

	queryString, _ := query.Values(GetBackupAndRestoreExecutionsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseBackupGetBackupAndRestoreExecutions{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBackupAndRestoreExecutions(GetBackupAndRestoreExecutionsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetBackupAndRestoreExecutions")
	}

	result := response.Result().(*ResponseBackupGetBackupAndRestoreExecutions)
	return result, response, err

}

//GetBackupAndRestoreExecution Get Backup And Restore Execution - 3eb8-dbe1-4828-9423
/* This api is used to get the execution detail of a specific backup or restore worflow process


@param id id path parameter. The `id` of the backup execution to be retrieved.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-backup-and-restore-execution
*/
func (s *BackupService) GetBackupAndRestoreExecution(id string) (*ResponseBackupGetBackupAndRestoreExecution, *resty.Response, error) {
	path := "/dna/system/api/v1/backupRestoreExecutions/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseBackupGetBackupAndRestoreExecution{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBackupAndRestoreExecution(id)
		}
		return nil, response, fmt.Errorf("error with operation GetBackupAndRestoreExecution")
	}

	result := response.Result().(*ResponseBackupGetBackupAndRestoreExecution)
	return result, response, err

}

//GetBackupStorages Get Backup Storages - 2398-8984-48d9-ae90
/* This api is used to get all the mounted backup storage information like mount point, disk size based on the provided storage type


@param GetBackupStoragesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-backup-storages
*/
func (s *BackupService) GetBackupStorages(GetBackupStoragesQueryParams *GetBackupStoragesQueryParams) (*ResponseBackupGetBackupStorages, *resty.Response, error) {
	path := "/dna/system/api/v1/backupStorages"

	queryString, _ := query.Values(GetBackupStoragesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseBackupGetBackupStorages{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBackupStorages(GetBackupStoragesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetBackupStorages")
	}

	result := response.Result().(*ResponseBackupGetBackupStorages)
	return result, response, err

}

//GetAllBackup Get All Backup - 6cbf-ea56-42c9-9b6d
/* This api is used to get all the backup available in the configured storage


@param GetAllBackupQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-all-backup
*/
func (s *BackupService) GetAllBackup(GetAllBackupQueryParams *GetAllBackupQueryParams) (*ResponseBackupGetAllBackup, *resty.Response, error) {
	path := "/dna/system/api/v1/backups"

	queryString, _ := query.Values(GetAllBackupQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseBackupGetAllBackup{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetAllBackup(GetAllBackupQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetAllBackup")
	}

	result := response.Result().(*ResponseBackupGetAllBackup)
	return result, response, err

}

//GetBackupByID Get Backup By Id - 26ab-9821-4979-b15b
/* This api is used to get a specific backup based on the provided `backup id`.


@param id id path parameter. The `id` of the backup to be retrieved.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-backup-by-id
*/
func (s *BackupService) GetBackupByID(id string) (*ResponseBackupGetBackupByID, *resty.Response, error) {
	path := "/dna/system/api/v1/backups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseBackupGetBackupByID{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetBackupByID(id)
		}
		return nil, response, fmt.Errorf("error with operation GetBackupById")
	}

	result := response.Result().(*ResponseBackupGetBackupByID)
	return result, response, err

}

//CreateBackupConfiguration Create Backup Configuration - c098-8833-4c9a-92f8
/* This api is used to create or update backup configuration. Obtain the `mountPath` value from the mountPoint attribute in the response of the `/dna/system/api/v1/backupStorages` API.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-backup-configuration
*/
func (s *BackupService) CreateBackupConfiguration(requestBackupCreateBackupConfiguration *RequestBackupCreateBackupConfiguration) (*ResponseBackupCreateBackupConfiguration, *resty.Response, error) {
	path := "/dna/system/api/v1/backupConfiguration"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupCreateBackupConfiguration).
		// SetResult(&ResponseBackupCreateBackupConfiguration{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateBackupConfiguration(requestBackupCreateBackupConfiguration)
		}

		return nil, response, fmt.Errorf("error with operation CreateBackupConfiguration")
	}

	result := response.Result().(ResponseBackupCreateBackupConfiguration)

	return &result, response, err

}

//CreateNFSConfiguration Create NFS Configuration - 2ea6-eb71-40aa-bc4b
/* This api is used to create NFS configuration.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-n-f-s-configuration
*/
func (s *BackupService) CreateNFSConfiguration(requestBackupCreateNFSConfiguration *RequestBackupCreateNFSConfiguration) (*ResponseBackupCreateNFSConfiguration, *resty.Response, error) {
	path := "/dna/system/api/v1/backupNfsConfigurations"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupCreateNFSConfiguration).
		// SetResult(&ResponseBackupCreateNFSConfiguration{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateNFSConfiguration(requestBackupCreateNFSConfiguration)
		}

		return nil, response, fmt.Errorf("error with operation CreateNFSConfiguration")
	}

	result := response.Result().(ResponseBackupCreateNFSConfiguration)

	return &result, response, err

}

//CreateBackup Create Backup - a890-fa3c-4069-bde1
/* This api is used to trigger a workflow to create an on demand backup. To monitor the progress and completion of the backup creation , please call `/dna/system/api/v1/backupRestoreExecutions/{id}` api , where id is the taskId attribute from the response of the curent endpoint.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!create-backup
*/
func (s *BackupService) CreateBackup(requestBackupCreateBackup *RequestBackupCreateBackup) (*resty.Response, error) {
	path := "/dna/system/api/v1/backups"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestBackupCreateBackup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.CreateBackup(requestBackupCreateBackup)
		}

		return response, fmt.Errorf("error with operation CreateBackup")
	}

	return response, err

}

//DeleteNFSConfiguration Delete NFS Configuration - 6199-d965-4ebb-8258
/* This api is used to delete the NFS configuration. Obtain the `id` from the id attribute in the response of the `/dna/system/api/v1/backupNfsConfigurations` API.


@param id id path parameter. The `id` of the NFS configuration to be deleted.Obtain the `id` from the id attribute in the response of the `/dna/system/api/v1/backupNfsConfigurations` API.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-n-f-s-configuration
*/
func (s *BackupService) DeleteNFSConfiguration(id string) (*ResponseBackupDeleteNFSConfiguration, *resty.Response, error) {
	//id string
	path := "/dna/system/api/v1/backupNfsConfigurations/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		// SetResult(&ResponseBackupDeleteNFSConfiguration{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteNFSConfiguration(
				id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteNFSConfiguration")
	}

	result := response.Result().(ResponseBackupDeleteNFSConfiguration)

	return &result, response, err

}

//DeleteBackup Delete Backup - 34a4-c9e6-4f4a-af31
/* This api is used to trigger delete workflow of a specific backup based on the provided `id` Obtain the `id` from the id attribute in the response of the `/dna/system/api/v1/backups` API. To monitor the progress and completion of the backup deletion , please call `/dna/system/api/v1/backupRestoreExecutions/{id}` api , where id is the taskId attribute from the response of the curent endpoint.


@param id id path parameter. The `id` of the backup to be deleted.Obtain the 'id' from the id attribute in the response of the `/dna/system/api/v1/backups` API.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-backup
*/
func (s *BackupService) DeleteBackup(id string) (*resty.Response, error) {
	//id string
	path := "/dna/system/api/v1/backups/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteBackup(
				id)
		}
		return response, fmt.Errorf("error with operation DeleteBackup")
	}

	return response, err

}
