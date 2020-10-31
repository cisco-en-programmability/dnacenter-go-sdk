package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SensorsService is the service to communicate with the Sensors API endpoint
type SensorsService service

// ApCoverage is the ApCoverage definition
type ApCoverage struct {
	Bands             string `json:"bands,omitempty"`          //
	NumberOfApsToTest string `json:"intOfApsToTest,omitempty"` //
	RssiThreshold     string `json:"rssiThreshold,omitempty"`  //
}

// CreateSensorTestTemplateRequest is the CreateSensorTestTemplateRequest definition
type CreateSensorTestTemplateRequest struct {
	ApCoverage   []ApCoverage `json:"apCoverage,omitempty"`   //
	Connection   string       `json:"connection,omitempty"`   //
	ModelVersion int          `json:"modelVersion,omitempty"` //
	Name         string       `json:"name,omitempty"`         //
	SSIDs        []SSIDs      `json:"ssids,omitempty"`        //
}

// DuplicateSensorTestTemplateRequest is the DuplicateSensorTestTemplateRequest definition
type DuplicateSensorTestTemplateRequest struct {
	NewTemplateName string `json:"newTemplateName,omitempty"` //
	TemplateName    string `json:"templateName,omitempty"`    //
}

// EditSensorTestTemplateRequest is the EditSensorTestTemplateRequest definition
type EditSensorTestTemplateRequest struct {
	LocationInfoList []LocationInfoList `json:"locationInfoList,omitempty"` //
	Schedule         Schedule           `json:"schedule,omitempty"`         //
	TemplateName     string             `json:"templateName,omitempty"`     //
}

// Frequency is the Frequency definition
type Frequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// LocationInfoList is the LocationInfoList definition
type LocationInfoList struct {
	AllSensors    bool   `json:"allSensors,omitempty"`    //
	LocationID    string `json:"locationId,omitempty"`    //
	LocationType  string `json:"locationType,omitempty"`  //
	SiteHierarchy string `json:"siteHierarchy,omitempty"` //
}

// RunNowSensorTestRequest is the RunNowSensorTestRequest definition
type RunNowSensorTestRequest struct {
	TemplateName string `json:"templateName,omitempty"` //
}

// Schedule is the Schedule definition
type Schedule struct {
	Frequency        Frequency       `json:"frequency,omitempty"`        //
	ScheduleRange    []ScheduleRange `json:"scheduleRange,omitempty"`    //
	TestScheduleMode string          `json:"testScheduleMode,omitempty"` //
}

// ScheduleRange is the ScheduleRange definition
type ScheduleRange struct {
	Day       string      `json:"day,omitempty"`       //
	TimeRange []TimeRange `json:"timeRange,omitempty"` //
}

// SSIDs is the SSIDs definition
// type SSIDs struct {
// 	AuthType    string     `json:"authType,omitempty"`    //
// 	Categories  []string   `json:"categories,omitempty"`  //
// 	ProfileName string     `json:"profileName,omitempty"` //
// 	Psk         string     `json:"psk,omitempty"`         //
// 	QosPolicy   string     `json:"qosPolicy,omitempty"`   //
// 	SSID        string     `json:"ssid,omitempty"`        //
// 	Tests       []Tests    `json:"tests,omitempty"`       //
// 	ThirdParty  ThirdParty `json:"thirdParty,omitempty"`  //
// }

// Tests is the Tests definition
type Tests struct {
	Config []string `json:"config,omitempty"` //
	Name   string   `json:"name,omitempty"`   //
}

// ThirdParty is the ThirdParty definition
type ThirdParty struct {
	Selected bool `json:"selected,omitempty"` //
}

// TimeRange is the TimeRange definition
type TimeRange struct {
	Frequency Frequency `json:"frequency,omitempty"` //
	From      string    `json:"from,omitempty"`      //
	To        string    `json:"to,omitempty"`        //
}

// CreateSensorTestTemplateResponse is the CreateSensorTestTemplateResponse definition
type CreateSensorTestTemplateResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DeleteSensorTestResponse is the DeleteSensorTestResponse definition
type DeleteSensorTestResponse struct {
	Response DeleteSensorResponse `json:"response,omitempty"` //
	Version  string               `json:"version,omitempty"`  //
}

// DuplicateSensorTestTemplateResponse is the DuplicateSensorTestTemplateResponse definition
type DuplicateSensorTestTemplateResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// EditSensorTestTemplateResponse is the EditSensorTestTemplateResponse definition
type EditSensorTestTemplateResponse struct {
	Response Response `json:"response,omitempty"` //
	Version  string   `json:"version,omitempty"`  //
}

// DeleteSensorResponse is the Response definition
type DeleteSensorResponse struct {
	ID                     string             `json:"_id,omitempty"`                    //
	ApCoverage             []ApCoverage       `json:"apCoverage,omitempty"`             //
	Connection             string             `json:"connection,omitempty"`             //
	EncryptionMode         string             `json:"encryptionMode,omitempty"`         //
	Frequency              string             `json:"frequency,omitempty"`              //
	LastModifiedTime       int                `json:"lastModifiedTime,omitempty"`       //
	LegacyTestSuite        bool               `json:"legacyTestSuite,omitempty"`        //
	Location               string             `json:"location,omitempty"`               //
	LocationInfoList       []LocationInfoList `json:"locationInfoList,omitempty"`       //
	ModelVersion           int                `json:"modelVersion,omitempty"`           //
	Name                   string             `json:"name,omitempty"`                   //
	NumAssociatedSensor    int                `json:"numAssociatedSensor,omitempty"`    //
	NumNeighborAPThreshold int                `json:"numNeighborAPThreshold,omitempty"` //
	RadioAsSensorRemoved   bool               `json:"radioAsSensorRemoved,omitempty"`   //
	RssiThreshold          int                `json:"rssiThreshold,omitempty"`          //
	RunNow                 string             `json:"runNow,omitempty"`                 //
	Schedule               Schedule           `json:"schedule,omitempty"`               //
	ScheduleInDays         int                `json:"scheduleInDays,omitempty"`         //
	Sensors                []string           `json:"sensors,omitempty"`                //
	ShowWlcUpgradeBanner   bool               `json:"showWlcUpgradeBanner,omitempty"`   //
	SiteHierarchy          string             `json:"siteHierarchy,omitempty"`          //
	SSIDs                  []SSIDs            `json:"ssids,omitempty"`                  //
	StartTime              int                `json:"startTime,omitempty"`              //
	Status                 string             `json:"status,omitempty"`                 //
	TenantID               string             `json:"tenantId,omitempty"`               //
	TestDurationEstimate   int                `json:"testDurationEstimate,omitempty"`   //
	TestScheduleMode       string             `json:"testScheduleMode,omitempty"`       //
	TestTemplate           bool               `json:"testTemplate,omitempty"`           //
	Tests                  string             `json:"tests,omitempty"`                  //
	Version                int                `json:"version,omitempty"`                //
	Wlans                  []string           `json:"wlans,omitempty"`                  //
}

// SensorsResponse is the SensorsResponse definition
type SensorsResponse struct {
	Response []Response `json:"response,omitempty"` //
	Version  string     `json:"version,omitempty"`  //
}

// SSHConfig is the SshConfig definition
type SSHConfig struct {
	EnablePassword string `json:"enablePassword,omitempty"` //
	SSHPassword    string `json:"sshPassword,omitempty"`    //
	SSHState       string `json:"sshState,omitempty"`       //
	SSHUserName    string `json:"sshUserName,omitempty"`    //
}

// SSIDs is the SSIDs definition
type SSIDs struct {
	AuthProtocol              string     `json:"authProtocol,omitempty"`              //
	AuthType                  string     `json:"authType,omitempty"`                  //
	AuthTypeRcvd              string     `json:"authTypeRcvd,omitempty"`              //
	Bands                     string     `json:"bands,omitempty"`                     //
	Certdownloadurl           string     `json:"certdownloadurl,omitempty"`           //
	Certfilename              string     `json:"certfilename,omitempty"`              //
	Certpassphrase            string     `json:"certpassphrase,omitempty"`            //
	Certstatus                string     `json:"certstatus,omitempty"`                //
	Certxferprotocol          string     `json:"certxferprotocol,omitempty"`          //
	EapMethod                 string     `json:"eapMethod,omitempty"`                 //
	ExtWebAuth                bool       `json:"extWebAuth,omitempty"`                //
	ExtWebAuthAccessURL       string     `json:"extWebAuthAccessUrl,omitempty"`       //
	ExtWebAuthHTMLTag         []string   `json:"extWebAuthHtmlTag,omitempty"`         //
	ExtWebAuthPortal          string     `json:"extWebAuthPortal,omitempty"`          //
	ExtWebAuthVirtualIP       string     `json:"extWebAuthVirtualIp,omitempty"`       //
	ID                        int        `json:"id,omitempty"`                        //
	Layer3webAuthEmailAddress string     `json:"layer3webAuthEmailAddress,omitempty"` //
	Layer3webAuthpassword     string     `json:"layer3webAuthpassword,omitempty"`     //
	Layer3webAuthsecurity     string     `json:"layer3webAuthsecurity,omitempty"`     //
	Layer3webAuthuserName     string     `json:"layer3webAuthuserName,omitempty"`     //
	NumAps                    int        `json:"numAps,omitempty"`                    //
	NumSensors                int        `json:"numSensors,omitempty"`                //
	Password                  string     `json:"password,omitempty"`                  //
	ProfileName               string     `json:"profileName,omitempty"`               //
	Psk                       string     `json:"psk,omitempty"`                       //
	QosPolicy                 string     `json:"qosPolicy,omitempty"`                 //
	Scep                      bool       `json:"scep,omitempty"`                      //
	SSID                      string     `json:"ssid,omitempty"`                      //
	Status                    string     `json:"status,omitempty"`                    //
	Tests                     []Tests    `json:"tests,omitempty"`                     //
	ThirdParty                ThirdParty `json:"thirdParty,omitempty"`                //
	Username                  string     `json:"username,omitempty"`                  //
	ValidFrom                 int        `json:"validFrom,omitempty"`                 //
	ValidTo                   int        `json:"validTo,omitempty"`                   //
	WhiteList                 bool       `json:"whiteList,omitempty"`                 //
	WlanID                    int        `json:"wlanId,omitempty"`                    //
	Wlc                       string     `json:"wlc,omitempty"`                       //
}

// CreateSensorTestTemplate createSensorTestTemplate
/* Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID
 */
func (s *SensorsService) CreateSensorTestTemplate(createSensorTestTemplateRequest *CreateSensorTestTemplateRequest) (*CreateSensorTestTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/sensor"

	response, err := RestyClient.R().
		SetBody(createSensorTestTemplateRequest).
		SetResult(&CreateSensorTestTemplateResponse{}).
		SetError(&Error{}).
		Post(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*CreateSensorTestTemplateResponse)
	return result, response, err

}

// DeleteSensorTestQueryParams defines the query parameters for this request
type DeleteSensorTestQueryParams struct {
	TemplateName string `url:"templateName,omitempty"` //
}

// DeleteSensorTest deleteSensorTest
/* Intent API to delete an existing SENSOR test template
@param templateName
*/
func (s *SensorsService) DeleteSensorTest(deleteSensorTestQueryParams *DeleteSensorTestQueryParams) (*resty.Response, error) {

	path := "/dna/intent/api/v1/sensor"

	queryString, _ := query.Values(deleteSensorTestQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetError(&Error{}).
		Delete(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// DuplicateSensorTestTemplate duplicateSensorTestTemplate
/* Intent API to duplicate an existing SENSOR test template
 */
func (s *SensorsService) DuplicateSensorTestTemplate(duplicateSensorTestTemplateRequest *DuplicateSensorTestTemplateRequest) (*DuplicateSensorTestTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/sensorTestTemplate"

	response, err := RestyClient.R().
		SetBody(duplicateSensorTestTemplateRequest).
		SetResult(&DuplicateSensorTestTemplateResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*DuplicateSensorTestTemplateResponse)
	return result, response, err

}

// EditSensorTestTemplate editSensorTestTemplate
/* Intent API to deploy, schedule, or edit and existing SENSOR test template
 */
func (s *SensorsService) EditSensorTestTemplate(editSensorTestTemplateRequest *EditSensorTestTemplateRequest) (*EditSensorTestTemplateResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/AssuranceScheduleSensorTest"

	response, err := RestyClient.R().
		SetBody(editSensorTestTemplateRequest).
		SetResult(&EditSensorTestTemplateResponse{}).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*EditSensorTestTemplateResponse)
	return result, response, err

}

// RunNowSensorTest runNowSensorTest
/* Intent API to run a deployed SENSOR test
 */
func (s *SensorsService) RunNowSensorTest(runNowSensorTestRequest *RunNowSensorTestRequest) (*resty.Response, error) {

	path := "/dna/intent/api/v1/sensor-run-now"

	response, err := RestyClient.R().
		SetBody(runNowSensorTestRequest).
		SetError(&Error{}).
		Put(path)

	if err != nil {
		return nil, err
	}

	return response, err

}

// SensorsQueryParams defines the query parameters for this request
type SensorsQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //
}

// Sensors sensors
/* Intent API to get a list of SENSOR devices
@param siteID
*/
func (s *SensorsService) Sensors(sensorsQueryParams *SensorsQueryParams) (*SensorsResponse, *resty.Response, error) {

	path := "/dna/intent/api/v1/sensor"

	queryString, _ := query.Values(sensorsQueryParams)

	response, err := RestyClient.R().
		SetQueryString(queryString.Encode()).
		SetResult(&SensorsResponse{}).
		SetError(&Error{}).
		Get(path)

	if err != nil {
		return nil, nil, err
	}

	result := response.Result().(*SensorsResponse)
	return result, response, err

}
