package dnac

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type RestoreService service

type RequestReportsRestoreBackup struct {
	EncryptionPassphrase string `json:"encryptionPassphrase,omitempty"` // Passphrase to restore backup
}

//RestoreBackup Restore Backup - e1ae-d8ee-4bfa-b686
/* This api is used to trigger restore workflow of a specific backup. Obtain the `id` from the id attribute in the response of the `/dna/system/api/v1/backups` API. To monitor the progress and completion of the backup deletion , please call `/dna/system/api/v1/backupRestoreExecutions/{id}` api , where id is the taskId attribute from the response of the curent endpoint.


@param id id path parameter. The `id` of the backup to be restored.Obtain the `id` from the id attribute in the response of the `/dna/system/api/v1/backups` API.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!restore-backup
*/
func (s *RestoreService) RestoreBackup(id string, requestReportsRestoreBackup *RequestReportsRestoreBackup) (*resty.Response, error) {
	path := "/dna/system/api/v1/backups/{id}/restore"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestReportsRestoreBackup).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.RestoreBackup(id, requestReportsRestoreBackup)
		}

		return response, fmt.Errorf("error with operation RestoreBackup")
	}

	return response, err

}
