package dnac

import (
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

// SensorsService is the service to communicate with the Sensors API endpoint
type SensorsService service

// CreateSensorTestTemplateRequest is the createSensorTestTemplateRequest definition
type CreateSensorTestTemplateRequest struct {
	ApCoverage   []CreateSensorTestTemplateRequestApCoverage `json:"apCoverage,omitempty"`   //
	Connection   string                                      `json:"connection,omitempty"`   //
	ModelVersion int                                         `json:"modelVersion,omitempty"` //
	Name         string                                      `json:"name,omitempty"`         //
	SSIDs        []CreateSensorTestTemplateRequestSSIDs      `json:"ssids,omitempty"`        //
}

// CreateSensorTestTemplateRequestApCoverage is the createSensorTestTemplateRequestApCoverage definition
type CreateSensorTestTemplateRequestApCoverage struct {
	Bands             string `json:"bands,omitempty"`             //
	NumberOfApsToTest string `json:"numberOfApsToTest,omitempty"` //
	RssiThreshold     string `json:"rssiThreshold,omitempty"`     //
}

// CreateSensorTestTemplateRequestSSIDs is the createSensorTestTemplateRequestSSIDs definition
type CreateSensorTestTemplateRequestSSIDs struct {
	AuthType    string                                         `json:"authType,omitempty"`    //
	Categories  []string                                       `json:"categories,omitempty"`  //
	ProfileName string                                         `json:"profileName,omitempty"` //
	Psk         string                                         `json:"psk,omitempty"`         //
	QosPolicy   string                                         `json:"qosPolicy,omitempty"`   //
	SSID        string                                         `json:"ssid,omitempty"`        //
	Tests       []CreateSensorTestTemplateRequestSSIDsTests    `json:"tests,omitempty"`       //
	ThirdParty  CreateSensorTestTemplateRequestSSIDsThirdParty `json:"thirdParty,omitempty"`  //
}

// CreateSensorTestTemplateRequestSSIDsCategories is the createSensorTestTemplateRequestSSIDsCategories definition
type CreateSensorTestTemplateRequestSSIDsCategories []string

// CreateSensorTestTemplateRequestSSIDsTests is the createSensorTestTemplateRequestSSIDsTests definition
type CreateSensorTestTemplateRequestSSIDsTests struct {
	Config []string `json:"config,omitempty"` //
	Name   string   `json:"name,omitempty"`   //
}

// CreateSensorTestTemplateRequestSSIDsTestsConfig is the createSensorTestTemplateRequestSSIDsTestsConfig definition
type CreateSensorTestTemplateRequestSSIDsTestsConfig []string

// CreateSensorTestTemplateRequestSSIDsThirdParty is the createSensorTestTemplateRequestSSIDsThirdParty definition
type CreateSensorTestTemplateRequestSSIDsThirdParty struct {
	Selected bool `json:"selected,omitempty"` //
}

// DuplicateSensorTestTemplateRequest is the duplicateSensorTestTemplateRequest definition
type DuplicateSensorTestTemplateRequest struct {
	NewTemplateName string `json:"newTemplateName,omitempty"` //
	TemplateName    string `json:"templateName,omitempty"`    //
}

// EditSensorTestTemplateRequest is the editSensorTestTemplateRequest definition
type EditSensorTestTemplateRequest struct {
	LocationInfoList []EditSensorTestTemplateRequestLocationInfoList `json:"locationInfoList,omitempty"` //
	Schedule         EditSensorTestTemplateRequestSchedule           `json:"schedule,omitempty"`         //
	TemplateName     string                                          `json:"templateName,omitempty"`     //
}

// EditSensorTestTemplateRequestLocationInfoList is the editSensorTestTemplateRequestLocationInfoList definition
type EditSensorTestTemplateRequestLocationInfoList struct {
	AllSensors    bool   `json:"allSensors,omitempty"`    //
	LocationID    string `json:"locationId,omitempty"`    //
	LocationType  string `json:"locationType,omitempty"`  //
	SiteHierarchy string `json:"siteHierarchy,omitempty"` //
}

// EditSensorTestTemplateRequestSchedule is the editSensorTestTemplateRequestSchedule definition
type EditSensorTestTemplateRequestSchedule struct {
	Frequency        EditSensorTestTemplateRequestScheduleFrequency       `json:"frequency,omitempty"`        //
	ScheduleRange    []EditSensorTestTemplateRequestScheduleScheduleRange `json:"scheduleRange,omitempty"`    //
	TestScheduleMode string                                               `json:"testScheduleMode,omitempty"` //
}

// EditSensorTestTemplateRequestScheduleFrequency is the editSensorTestTemplateRequestScheduleFrequency definition
type EditSensorTestTemplateRequestScheduleFrequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// EditSensorTestTemplateRequestScheduleScheduleRange is the editSensorTestTemplateRequestScheduleScheduleRange definition
type EditSensorTestTemplateRequestScheduleScheduleRange struct {
	Day       string                                                        `json:"day,omitempty"`       //
	TimeRange []EditSensorTestTemplateRequestScheduleScheduleRangeTimeRange `json:"timeRange,omitempty"` //
}

// EditSensorTestTemplateRequestScheduleScheduleRangeTimeRange is the editSensorTestTemplateRequestScheduleScheduleRangeTimeRange definition
type EditSensorTestTemplateRequestScheduleScheduleRangeTimeRange struct {
	Frequency EditSensorTestTemplateRequestScheduleScheduleRangeTimeRangeFrequency `json:"frequency,omitempty"` //
	From      string                                                               `json:"from,omitempty"`      //
	To        string                                                               `json:"to,omitempty"`        //
}

// EditSensorTestTemplateRequestScheduleScheduleRangeTimeRangeFrequency is the editSensorTestTemplateRequestScheduleScheduleRangeTimeRangeFrequency definition
type EditSensorTestTemplateRequestScheduleScheduleRangeTimeRangeFrequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// RunNowSensorTestRequest is the runNowSensorTestRequest definition
type RunNowSensorTestRequest struct {
	TemplateName string `json:"templateName,omitempty"` //
}

// CreateSensorTestTemplateResponse is the createSensorTestTemplateResponse definition
type CreateSensorTestTemplateResponse struct {
	Response CreateSensorTestTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                                   `json:"version,omitempty"`  //
}

// CreateSensorTestTemplateResponseResponse is the createSensorTestTemplateResponseResponse definition
type CreateSensorTestTemplateResponseResponse struct {
	TypeID                 string                                               `json:"_id,omitempty"`                    //
	ApCoverage             []CreateSensorTestTemplateResponseResponseApCoverage `json:"apCoverage,omitempty"`             //
	Connection             string                                               `json:"connection,omitempty"`             //
	EncryptionMode         string                                               `json:"encryptionMode,omitempty"`         //
	Frequency              string                                               `json:"frequency,omitempty"`              //
	LastModifiedTime       int                                                  `json:"lastModifiedTime,omitempty"`       //
	LegacyTestSuite        bool                                                 `json:"legacyTestSuite,omitempty"`        //
	Location               string                                               `json:"location,omitempty"`               //
	LocationInfoList       []string                                             `json:"locationInfoList,omitempty"`       //
	ModelVersion           int                                                  `json:"modelVersion,omitempty"`           //
	Name                   string                                               `json:"name,omitempty"`                   //
	NumAssociatedSensor    int                                                  `json:"numAssociatedSensor,omitempty"`    //
	NumNeighborAPThreshold int                                                  `json:"numNeighborAPThreshold,omitempty"` //
	RadioAsSensorRemoved   bool                                                 `json:"radioAsSensorRemoved,omitempty"`   //
	RssiThreshold          int                                                  `json:"rssiThreshold,omitempty"`          //
	RunNow                 string                                               `json:"runNow,omitempty"`                 //
	Schedule               string                                               `json:"schedule,omitempty"`               //
	ScheduleInDays         int                                                  `json:"scheduleInDays,omitempty"`         //
	Sensors                []string                                             `json:"sensors,omitempty"`                //
	ShowWlcUpgradeBanner   bool                                                 `json:"showWlcUpgradeBanner,omitempty"`   //
	SiteHierarchy          string                                               `json:"siteHierarchy,omitempty"`          //
	SSIDs                  []CreateSensorTestTemplateResponseResponseSSIDs      `json:"ssids,omitempty"`                  //
	StartTime              int                                                  `json:"startTime,omitempty"`              //
	Status                 string                                               `json:"status,omitempty"`                 //
	TenantID               string                                               `json:"tenantId,omitempty"`               //
	TestDurationEstimate   int                                                  `json:"testDurationEstimate,omitempty"`   //
	TestScheduleMode       string                                               `json:"testScheduleMode,omitempty"`       //
	TestTemplate           bool                                                 `json:"testTemplate,omitempty"`           //
	Tests                  string                                               `json:"tests,omitempty"`                  //
	Version                int                                                  `json:"version,omitempty"`                //
	WLANs                  []string                                             `json:"wlans,omitempty"`                  //
}

// CreateSensorTestTemplateResponseResponseApCoverage is the createSensorTestTemplateResponseResponseApCoverage definition
type CreateSensorTestTemplateResponseResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             //
	NumberOfApsToTest int    `json:"numberOfApsToTest,omitempty"` //
	RssiThreshold     int    `json:"rssiThreshold,omitempty"`     //
}

// CreateSensorTestTemplateResponseResponseLocationInfoList is the createSensorTestTemplateResponseResponseLocationInfoList definition
type CreateSensorTestTemplateResponseResponseLocationInfoList []string

// CreateSensorTestTemplateResponseResponseSSIDs is the createSensorTestTemplateResponseResponseSSIDs definition
type CreateSensorTestTemplateResponseResponseSSIDs struct {
	AuthProtocol              string                                                  `json:"authProtocol,omitempty"`              //
	AuthType                  string                                                  `json:"authType,omitempty"`                  //
	AuthTypeRcvd              string                                                  `json:"authTypeRcvd,omitempty"`              //
	Bands                     string                                                  `json:"bands,omitempty"`                     //
	Certdownloadurl           string                                                  `json:"certdownloadurl,omitempty"`           //
	Certfilename              string                                                  `json:"certfilename,omitempty"`              //
	Certpassphrase            string                                                  `json:"certpassphrase,omitempty"`            //
	Certstatus                string                                                  `json:"certstatus,omitempty"`                //
	Certxferprotocol          string                                                  `json:"certxferprotocol,omitempty"`          //
	EapMethod                 string                                                  `json:"eapMethod,omitempty"`                 //
	ExtWebAuth                bool                                                    `json:"extWebAuth,omitempty"`                //
	ExtWebAuthAccessURL       string                                                  `json:"extWebAuthAccessUrl,omitempty"`       //
	ExtWebAuthHTMLTag         []string                                                `json:"extWebAuthHtmlTag,omitempty"`         //
	ExtWebAuthPortal          string                                                  `json:"extWebAuthPortal,omitempty"`          //
	ExtWebAuthVirtualIP       string                                                  `json:"extWebAuthVirtualIp,omitempty"`       //
	ID                        int                                                     `json:"id,omitempty"`                        //
	Layer3webAuthEmailAddress string                                                  `json:"layer3webAuthEmailAddress,omitempty"` //
	Layer3webAuthpassword     string                                                  `json:"layer3webAuthpassword,omitempty"`     //
	Layer3webAuthsecurity     string                                                  `json:"layer3webAuthsecurity,omitempty"`     //
	Layer3webAuthuserName     string                                                  `json:"layer3webAuthuserName,omitempty"`     //
	NumAps                    int                                                     `json:"numAps,omitempty"`                    //
	NumSensors                int                                                     `json:"numSensors,omitempty"`                //
	Password                  string                                                  `json:"password,omitempty"`                  //
	ProfileName               string                                                  `json:"profileName,omitempty"`               //
	Psk                       string                                                  `json:"psk,omitempty"`                       //
	QosPolicy                 string                                                  `json:"qosPolicy,omitempty"`                 //
	Scep                      bool                                                    `json:"scep,omitempty"`                      //
	SSID                      string                                                  `json:"ssid,omitempty"`                      //
	Status                    string                                                  `json:"status,omitempty"`                    //
	Tests                     []CreateSensorTestTemplateResponseResponseSSIDsTests    `json:"tests,omitempty"`                     //
	ThirdParty                CreateSensorTestTemplateResponseResponseSSIDsThirdParty `json:"thirdParty,omitempty"`                //
	Username                  string                                                  `json:"username,omitempty"`                  //
	ValidFrom                 int                                                     `json:"validFrom,omitempty"`                 //
	ValidTo                   int                                                     `json:"validTo,omitempty"`                   //
	WhiteList                 bool                                                    `json:"whiteList,omitempty"`                 //
	WLANID                    int                                                     `json:"wlanId,omitempty"`                    //
	Wlc                       string                                                  `json:"wlc,omitempty"`                       //
}

// CreateSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag is the createSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag definition
type CreateSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag []string

// CreateSensorTestTemplateResponseResponseSSIDsTests is the createSensorTestTemplateResponseResponseSSIDsTests definition
type CreateSensorTestTemplateResponseResponseSSIDsTests struct {
	Config []string `json:"config,omitempty"` //
	Name   string   `json:"name,omitempty"`   //
}

// CreateSensorTestTemplateResponseResponseSSIDsTestsConfig is the createSensorTestTemplateResponseResponseSSIDsTestsConfig definition
type CreateSensorTestTemplateResponseResponseSSIDsTestsConfig []string

// CreateSensorTestTemplateResponseResponseSSIDsThirdParty is the createSensorTestTemplateResponseResponseSSIDsThirdParty definition
type CreateSensorTestTemplateResponseResponseSSIDsThirdParty struct {
	Selected bool `json:"selected,omitempty"` //
}

// CreateSensorTestTemplateResponseResponseSensors is the createSensorTestTemplateResponseResponseSensors definition
type CreateSensorTestTemplateResponseResponseSensors []string

// CreateSensorTestTemplateResponseResponseWLANs is the createSensorTestTemplateResponseResponseWLANs definition
type CreateSensorTestTemplateResponseResponseWLANs []string

// DeleteSensorTestResponse is the deleteSensorTestResponse definition
type DeleteSensorTestResponse struct {
	Response DeleteSensorTestResponseResponse `json:"response,omitempty"` //
	Version  string                           `json:"version,omitempty"`  //
}

// DeleteSensorTestResponseResponse is the deleteSensorTestResponseResponse definition
type DeleteSensorTestResponseResponse struct {
	Status       string `json:"status,omitempty"`       //
	TemplateName string `json:"templateName,omitempty"` //
}

// DuplicateSensorTestTemplateResponse is the duplicateSensorTestTemplateResponse definition
type DuplicateSensorTestTemplateResponse struct {
	Response DuplicateSensorTestTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                                      `json:"version,omitempty"`  //
}

// DuplicateSensorTestTemplateResponseResponse is the duplicateSensorTestTemplateResponseResponse definition
type DuplicateSensorTestTemplateResponseResponse struct {
	TypeID                 string                                                        `json:"_id,omitempty"`                    //
	ApCoverage             []DuplicateSensorTestTemplateResponseResponseApCoverage       `json:"apCoverage,omitempty"`             //
	Connection             string                                                        `json:"connection,omitempty"`             //
	EncryptionMode         string                                                        `json:"encryptionMode,omitempty"`         //
	Frequency              string                                                        `json:"frequency,omitempty"`              //
	LastModifiedTime       int                                                           `json:"lastModifiedTime,omitempty"`       //
	LegacyTestSuite        bool                                                          `json:"legacyTestSuite,omitempty"`        //
	Location               string                                                        `json:"location,omitempty"`               //
	LocationInfoList       []DuplicateSensorTestTemplateResponseResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	ModelVersion           int                                                           `json:"modelVersion,omitempty"`           //
	Name                   string                                                        `json:"name,omitempty"`                   //
	NumAssociatedSensor    int                                                           `json:"numAssociatedSensor,omitempty"`    //
	NumNeighborAPThreshold int                                                           `json:"numNeighborAPThreshold,omitempty"` //
	RadioAsSensorRemoved   bool                                                          `json:"radioAsSensorRemoved,omitempty"`   //
	RssiThreshold          int                                                           `json:"rssiThreshold,omitempty"`          //
	RunNow                 string                                                        `json:"runNow,omitempty"`                 //
	Schedule               DuplicateSensorTestTemplateResponseResponseSchedule           `json:"schedule,omitempty"`               //
	ScheduleInDays         int                                                           `json:"scheduleInDays,omitempty"`         //
	Sensors                []string                                                      `json:"sensors,omitempty"`                //
	ShowWlcUpgradeBanner   bool                                                          `json:"showWlcUpgradeBanner,omitempty"`   //
	SiteHierarchy          string                                                        `json:"siteHierarchy,omitempty"`          //
	SSIDs                  []DuplicateSensorTestTemplateResponseResponseSSIDs            `json:"ssids,omitempty"`                  //
	StartTime              int                                                           `json:"startTime,omitempty"`              //
	Status                 string                                                        `json:"status,omitempty"`                 //
	TenantID               string                                                        `json:"tenantId,omitempty"`               //
	TestDurationEstimate   int                                                           `json:"testDurationEstimate,omitempty"`   //
	TestScheduleMode       string                                                        `json:"testScheduleMode,omitempty"`       //
	TestTemplate           bool                                                          `json:"testTemplate,omitempty"`           //
	Tests                  string                                                        `json:"tests,omitempty"`                  //
	Version                int                                                           `json:"version,omitempty"`                //
	WLANs                  []string                                                      `json:"wlans,omitempty"`                  //
}

// DuplicateSensorTestTemplateResponseResponseApCoverage is the duplicateSensorTestTemplateResponseResponseApCoverage definition
type DuplicateSensorTestTemplateResponseResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             //
	NumberOfApsToTest int    `json:"numberOfApsToTest,omitempty"` //
	RssiThreshold     int    `json:"rssiThreshold,omitempty"`     //
}

// DuplicateSensorTestTemplateResponseResponseLocationInfoList is the duplicateSensorTestTemplateResponseResponseLocationInfoList definition
type DuplicateSensorTestTemplateResponseResponseLocationInfoList struct {
	AllSensors     bool     `json:"allSensors,omitempty"`     //
	LocationID     string   `json:"locationId,omitempty"`     //
	LocationType   string   `json:"locationType,omitempty"`   //
	MacAddressList []string `json:"macAddressList,omitempty"` //
	SiteHierarchy  string   `json:"siteHierarchy,omitempty"`  //
}

// DuplicateSensorTestTemplateResponseResponseLocationInfoListMacAddressList is the duplicateSensorTestTemplateResponseResponseLocationInfoListMacAddressList definition
type DuplicateSensorTestTemplateResponseResponseLocationInfoListMacAddressList []string

// DuplicateSensorTestTemplateResponseResponseSSIDs is the duplicateSensorTestTemplateResponseResponseSSIDs definition
type DuplicateSensorTestTemplateResponseResponseSSIDs struct {
	AuthProtocol              string                                                     `json:"authProtocol,omitempty"`              //
	AuthType                  string                                                     `json:"authType,omitempty"`                  //
	AuthTypeRcvd              string                                                     `json:"authTypeRcvd,omitempty"`              //
	Bands                     string                                                     `json:"bands,omitempty"`                     //
	Certdownloadurl           string                                                     `json:"certdownloadurl,omitempty"`           //
	Certfilename              string                                                     `json:"certfilename,omitempty"`              //
	Certpassphrase            string                                                     `json:"certpassphrase,omitempty"`            //
	Certstatus                string                                                     `json:"certstatus,omitempty"`                //
	Certxferprotocol          string                                                     `json:"certxferprotocol,omitempty"`          //
	EapMethod                 string                                                     `json:"eapMethod,omitempty"`                 //
	ExtWebAuth                bool                                                       `json:"extWebAuth,omitempty"`                //
	ExtWebAuthAccessURL       string                                                     `json:"extWebAuthAccessUrl,omitempty"`       //
	ExtWebAuthHTMLTag         []string                                                   `json:"extWebAuthHtmlTag,omitempty"`         //
	ExtWebAuthPortal          string                                                     `json:"extWebAuthPortal,omitempty"`          //
	ExtWebAuthVirtualIP       string                                                     `json:"extWebAuthVirtualIp,omitempty"`       //
	ID                        int                                                        `json:"id,omitempty"`                        //
	Layer3webAuthEmailAddress string                                                     `json:"layer3webAuthEmailAddress,omitempty"` //
	Layer3webAuthpassword     string                                                     `json:"layer3webAuthpassword,omitempty"`     //
	Layer3webAuthsecurity     string                                                     `json:"layer3webAuthsecurity,omitempty"`     //
	Layer3webAuthuserName     string                                                     `json:"layer3webAuthuserName,omitempty"`     //
	NumAps                    int                                                        `json:"numAps,omitempty"`                    //
	NumSensors                int                                                        `json:"numSensors,omitempty"`                //
	Password                  string                                                     `json:"password,omitempty"`                  //
	ProfileName               string                                                     `json:"profileName,omitempty"`               //
	Psk                       string                                                     `json:"psk,omitempty"`                       //
	QosPolicy                 string                                                     `json:"qosPolicy,omitempty"`                 //
	Scep                      bool                                                       `json:"scep,omitempty"`                      //
	SSID                      string                                                     `json:"ssid,omitempty"`                      //
	Status                    string                                                     `json:"status,omitempty"`                    //
	Tests                     []DuplicateSensorTestTemplateResponseResponseSSIDsTests    `json:"tests,omitempty"`                     //
	ThirdParty                DuplicateSensorTestTemplateResponseResponseSSIDsThirdParty `json:"thirdParty,omitempty"`                //
	Username                  string                                                     `json:"username,omitempty"`                  //
	ValidFrom                 int                                                        `json:"validFrom,omitempty"`                 //
	ValidTo                   int                                                        `json:"validTo,omitempty"`                   //
	WhiteList                 bool                                                       `json:"whiteList,omitempty"`                 //
	WLANID                    int                                                        `json:"wlanId,omitempty"`                    //
	Wlc                       string                                                     `json:"wlc,omitempty"`                       //
}

// DuplicateSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag is the duplicateSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag definition
type DuplicateSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag []string

// DuplicateSensorTestTemplateResponseResponseSSIDsTests is the duplicateSensorTestTemplateResponseResponseSSIDsTests definition
type DuplicateSensorTestTemplateResponseResponseSSIDsTests struct {
	Config []string `json:"config,omitempty"` //
	Name   string   `json:"name,omitempty"`   //
}

// DuplicateSensorTestTemplateResponseResponseSSIDsTestsConfig is the duplicateSensorTestTemplateResponseResponseSSIDsTestsConfig definition
type DuplicateSensorTestTemplateResponseResponseSSIDsTestsConfig []string

// DuplicateSensorTestTemplateResponseResponseSSIDsThirdParty is the duplicateSensorTestTemplateResponseResponseSSIDsThirdParty definition
type DuplicateSensorTestTemplateResponseResponseSSIDsThirdParty struct {
	Selected bool `json:"selected,omitempty"` //
}

// DuplicateSensorTestTemplateResponseResponseSchedule is the duplicateSensorTestTemplateResponseResponseSchedule definition
type DuplicateSensorTestTemplateResponseResponseSchedule struct {
	Frequency        DuplicateSensorTestTemplateResponseResponseScheduleFrequency       `json:"frequency,omitempty"`        //
	ScheduleRange    []DuplicateSensorTestTemplateResponseResponseScheduleScheduleRange `json:"scheduleRange,omitempty"`    //
	StartTime        int                                                                `json:"startTime,omitempty"`        //
	TestScheduleMode string                                                             `json:"testScheduleMode,omitempty"` //
}

// DuplicateSensorTestTemplateResponseResponseScheduleFrequency is the duplicateSensorTestTemplateResponseResponseScheduleFrequency definition
type DuplicateSensorTestTemplateResponseResponseScheduleFrequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// DuplicateSensorTestTemplateResponseResponseScheduleScheduleRange is the duplicateSensorTestTemplateResponseResponseScheduleScheduleRange definition
type DuplicateSensorTestTemplateResponseResponseScheduleScheduleRange struct {
	Day       string                                                                      `json:"day,omitempty"`       //
	TimeRange []DuplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange `json:"timeRange,omitempty"` //
}

// DuplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange is the duplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange definition
type DuplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange struct {
	Frequency DuplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency `json:"frequency,omitempty"` //
	From      string                                                                             `json:"from,omitempty"`      //
	To        string                                                                             `json:"to,omitempty"`        //
}

// DuplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency is the duplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency definition
type DuplicateSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// DuplicateSensorTestTemplateResponseResponseSensors is the duplicateSensorTestTemplateResponseResponseSensors definition
type DuplicateSensorTestTemplateResponseResponseSensors []string

// DuplicateSensorTestTemplateResponseResponseWLANs is the duplicateSensorTestTemplateResponseResponseWLANs definition
type DuplicateSensorTestTemplateResponseResponseWLANs []string

// EditSensorTestTemplateResponse is the editSensorTestTemplateResponse definition
type EditSensorTestTemplateResponse struct {
	Response EditSensorTestTemplateResponseResponse `json:"response,omitempty"` //
	Version  string                                 `json:"version,omitempty"`  //
}

// EditSensorTestTemplateResponseResponse is the editSensorTestTemplateResponseResponse definition
type EditSensorTestTemplateResponseResponse struct {
	TypeID                 string                                                   `json:"_id,omitempty"`                    //
	ApCoverage             []EditSensorTestTemplateResponseResponseApCoverage       `json:"apCoverage,omitempty"`             //
	Connection             string                                                   `json:"connection,omitempty"`             //
	EncryptionMode         string                                                   `json:"encryptionMode,omitempty"`         //
	Frequency              string                                                   `json:"frequency,omitempty"`              //
	LastModifiedTime       int                                                      `json:"lastModifiedTime,omitempty"`       //
	LegacyTestSuite        bool                                                     `json:"legacyTestSuite,omitempty"`        //
	Location               string                                                   `json:"location,omitempty"`               //
	LocationInfoList       []EditSensorTestTemplateResponseResponseLocationInfoList `json:"locationInfoList,omitempty"`       //
	ModelVersion           int                                                      `json:"modelVersion,omitempty"`           //
	Name                   string                                                   `json:"name,omitempty"`                   //
	NumAssociatedSensor    int                                                      `json:"numAssociatedSensor,omitempty"`    //
	NumNeighborAPThreshold int                                                      `json:"numNeighborAPThreshold,omitempty"` //
	RadioAsSensorRemoved   bool                                                     `json:"radioAsSensorRemoved,omitempty"`   //
	RssiThreshold          int                                                      `json:"rssiThreshold,omitempty"`          //
	RunNow                 string                                                   `json:"runNow,omitempty"`                 //
	Schedule               EditSensorTestTemplateResponseResponseSchedule           `json:"schedule,omitempty"`               //
	ScheduleInDays         int                                                      `json:"scheduleInDays,omitempty"`         //
	Sensors                []string                                                 `json:"sensors,omitempty"`                //
	ShowWlcUpgradeBanner   bool                                                     `json:"showWlcUpgradeBanner,omitempty"`   //
	SiteHierarchy          string                                                   `json:"siteHierarchy,omitempty"`          //
	SSIDs                  []EditSensorTestTemplateResponseResponseSSIDs            `json:"ssids,omitempty"`                  //
	StartTime              int                                                      `json:"startTime,omitempty"`              //
	Status                 string                                                   `json:"status,omitempty"`                 //
	TenantID               string                                                   `json:"tenantId,omitempty"`               //
	TestDurationEstimate   int                                                      `json:"testDurationEstimate,omitempty"`   //
	TestScheduleMode       string                                                   `json:"testScheduleMode,omitempty"`       //
	TestTemplate           bool                                                     `json:"testTemplate,omitempty"`           //
	Tests                  string                                                   `json:"tests,omitempty"`                  //
	Version                int                                                      `json:"version,omitempty"`                //
	WLANs                  []string                                                 `json:"wlans,omitempty"`                  //
}

// EditSensorTestTemplateResponseResponseApCoverage is the editSensorTestTemplateResponseResponseApCoverage definition
type EditSensorTestTemplateResponseResponseApCoverage struct {
	Bands             string `json:"bands,omitempty"`             //
	NumberOfApsToTest int    `json:"numberOfApsToTest,omitempty"` //
	RssiThreshold     int    `json:"rssiThreshold,omitempty"`     //
}

// EditSensorTestTemplateResponseResponseLocationInfoList is the editSensorTestTemplateResponseResponseLocationInfoList definition
type EditSensorTestTemplateResponseResponseLocationInfoList struct {
	AllSensors     bool     `json:"allSensors,omitempty"`     //
	LocationID     string   `json:"locationId,omitempty"`     //
	LocationType   string   `json:"locationType,omitempty"`   //
	MacAddressList []string `json:"macAddressList,omitempty"` //
	SiteHierarchy  string   `json:"siteHierarchy,omitempty"`  //
}

// EditSensorTestTemplateResponseResponseLocationInfoListMacAddressList is the editSensorTestTemplateResponseResponseLocationInfoListMacAddressList definition
type EditSensorTestTemplateResponseResponseLocationInfoListMacAddressList []string

// EditSensorTestTemplateResponseResponseSSIDs is the editSensorTestTemplateResponseResponseSSIDs definition
type EditSensorTestTemplateResponseResponseSSIDs struct {
	AuthProtocol              string                                                `json:"authProtocol,omitempty"`              //
	AuthType                  string                                                `json:"authType,omitempty"`                  //
	AuthTypeRcvd              string                                                `json:"authTypeRcvd,omitempty"`              //
	Bands                     string                                                `json:"bands,omitempty"`                     //
	Certdownloadurl           string                                                `json:"certdownloadurl,omitempty"`           //
	Certfilename              string                                                `json:"certfilename,omitempty"`              //
	Certpassphrase            string                                                `json:"certpassphrase,omitempty"`            //
	Certstatus                string                                                `json:"certstatus,omitempty"`                //
	Certxferprotocol          string                                                `json:"certxferprotocol,omitempty"`          //
	EapMethod                 string                                                `json:"eapMethod,omitempty"`                 //
	ExtWebAuth                bool                                                  `json:"extWebAuth,omitempty"`                //
	ExtWebAuthAccessURL       string                                                `json:"extWebAuthAccessUrl,omitempty"`       //
	ExtWebAuthHTMLTag         []string                                              `json:"extWebAuthHtmlTag,omitempty"`         //
	ExtWebAuthPortal          string                                                `json:"extWebAuthPortal,omitempty"`          //
	ExtWebAuthVirtualIP       string                                                `json:"extWebAuthVirtualIp,omitempty"`       //
	ID                        int                                                   `json:"id,omitempty"`                        //
	Layer3webAuthEmailAddress string                                                `json:"layer3webAuthEmailAddress,omitempty"` //
	Layer3webAuthpassword     string                                                `json:"layer3webAuthpassword,omitempty"`     //
	Layer3webAuthsecurity     string                                                `json:"layer3webAuthsecurity,omitempty"`     //
	Layer3webAuthuserName     string                                                `json:"layer3webAuthuserName,omitempty"`     //
	NumAps                    int                                                   `json:"numAps,omitempty"`                    //
	NumSensors                int                                                   `json:"numSensors,omitempty"`                //
	Password                  string                                                `json:"password,omitempty"`                  //
	ProfileName               string                                                `json:"profileName,omitempty"`               //
	Psk                       string                                                `json:"psk,omitempty"`                       //
	QosPolicy                 string                                                `json:"qosPolicy,omitempty"`                 //
	Scep                      bool                                                  `json:"scep,omitempty"`                      //
	SSID                      string                                                `json:"ssid,omitempty"`                      //
	Status                    string                                                `json:"status,omitempty"`                    //
	Tests                     []EditSensorTestTemplateResponseResponseSSIDsTests    `json:"tests,omitempty"`                     //
	ThirdParty                EditSensorTestTemplateResponseResponseSSIDsThirdParty `json:"thirdParty,omitempty"`                //
	Username                  string                                                `json:"username,omitempty"`                  //
	ValidFrom                 int                                                   `json:"validFrom,omitempty"`                 //
	ValidTo                   int                                                   `json:"validTo,omitempty"`                   //
	WhiteList                 bool                                                  `json:"whiteList,omitempty"`                 //
	WLANID                    int                                                   `json:"wlanId,omitempty"`                    //
	Wlc                       string                                                `json:"wlc,omitempty"`                       //
}

// EditSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag is the editSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag definition
type EditSensorTestTemplateResponseResponseSSIDsExtWebAuthHTMLTag []string

// EditSensorTestTemplateResponseResponseSSIDsTests is the editSensorTestTemplateResponseResponseSSIDsTests definition
type EditSensorTestTemplateResponseResponseSSIDsTests struct {
	Config []string `json:"config,omitempty"` //
	Name   string   `json:"name,omitempty"`   //
}

// EditSensorTestTemplateResponseResponseSSIDsTestsConfig is the editSensorTestTemplateResponseResponseSSIDsTestsConfig definition
type EditSensorTestTemplateResponseResponseSSIDsTestsConfig []string

// EditSensorTestTemplateResponseResponseSSIDsThirdParty is the editSensorTestTemplateResponseResponseSSIDsThirdParty definition
type EditSensorTestTemplateResponseResponseSSIDsThirdParty struct {
	Selected bool `json:"selected,omitempty"` //
}

// EditSensorTestTemplateResponseResponseSchedule is the editSensorTestTemplateResponseResponseSchedule definition
type EditSensorTestTemplateResponseResponseSchedule struct {
	Frequency        EditSensorTestTemplateResponseResponseScheduleFrequency       `json:"frequency,omitempty"`        //
	ScheduleRange    []EditSensorTestTemplateResponseResponseScheduleScheduleRange `json:"scheduleRange,omitempty"`    //
	StartTime        int                                                           `json:"startTime,omitempty"`        //
	TestScheduleMode string                                                        `json:"testScheduleMode,omitempty"` //
}

// EditSensorTestTemplateResponseResponseScheduleFrequency is the editSensorTestTemplateResponseResponseScheduleFrequency definition
type EditSensorTestTemplateResponseResponseScheduleFrequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// EditSensorTestTemplateResponseResponseScheduleScheduleRange is the editSensorTestTemplateResponseResponseScheduleScheduleRange definition
type EditSensorTestTemplateResponseResponseScheduleScheduleRange struct {
	Day       string                                                                 `json:"day,omitempty"`       //
	TimeRange []EditSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange `json:"timeRange,omitempty"` //
}

// EditSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange is the editSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange definition
type EditSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRange struct {
	Frequency EditSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency `json:"frequency,omitempty"` //
	From      string                                                                        `json:"from,omitempty"`      //
	To        string                                                                        `json:"to,omitempty"`        //
}

// EditSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency is the editSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency definition
type EditSensorTestTemplateResponseResponseScheduleScheduleRangeTimeRangeFrequency struct {
	Unit  string `json:"unit,omitempty"`  //
	Value int    `json:"value,omitempty"` //
}

// EditSensorTestTemplateResponseResponseSensors is the editSensorTestTemplateResponseResponseSensors definition
type EditSensorTestTemplateResponseResponseSensors []string

// EditSensorTestTemplateResponseResponseWLANs is the editSensorTestTemplateResponseResponseWLANs definition
type EditSensorTestTemplateResponseResponseWLANs []string

// SensorsResponse is the sensorsResponse definition
type SensorsResponse struct {
	Response []SensorsResponseResponse `json:"response,omitempty"` //
	Version  string                    `json:"version,omitempty"`  //
}

// SensorsResponseResponse is the sensorsResponseResponse definition
type SensorsResponseResponse struct {
	BackhaulType       string                           `json:"backhaulType,omitempty"`       //
	EthernetMacAddress string                           `json:"ethernetMacAddress,omitempty"` //
	IPAddress          string                           `json:"ipAddress,omitempty"`          //
	IsLEDEnabled       bool                             `json:"isLEDEnabled,omitempty"`       //
	LastSeen           int                              `json:"lastSeen,omitempty"`           //
	Location           string                           `json:"location,omitempty"`           //
	Name               string                           `json:"name,omitempty"`               //
	RadioMacAddress    string                           `json:"radioMacAddress,omitempty"`    //
	SerialNumber       string                           `json:"serialNumber,omitempty"`       //
	SSHConfig          SensorsResponseResponseSSHConfig `json:"sshConfig,omitempty"`          //
	Status             string                           `json:"status,omitempty"`             //
	Type               string                           `json:"type,omitempty"`               //
	Version            string                           `json:"version,omitempty"`            //
}

// SensorsResponseResponseSSHConfig is the sensorsResponseResponseSSHConfig definition
type SensorsResponseResponseSSHConfig struct {
	EnablePassword string `json:"enablePassword,omitempty"` //
	SSHPassword    string `json:"sshPassword,omitempty"`    //
	SSHState       string `json:"sshState,omitempty"`       //
	SSHUserName    string `json:"sshUserName,omitempty"`    //
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
		return nil, nil
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
