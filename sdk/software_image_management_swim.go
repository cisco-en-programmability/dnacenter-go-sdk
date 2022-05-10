package dnac

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

type SoftwareImageManagementSwimService service

type TriggerSoftwareImageActivationQueryParams struct {
	ScheduleValidate bool `url:"scheduleValidate,omitempty"` //scheduleValidate, validates data before schedule (Optional)
}
type TriggerSoftwareImageActivationHeaderParams struct {
	ContentType string `url:"Content-Type,omitempty"` //Expects type string. Request body content type
	ClientType  string `url:"Client-Type,omitempty"`  //Expects type string. Client-type (Optional)
	ClientURL   string `url:"Client-Url,omitempty"`   //Expects type string. Client-url (Optional)
}
type GetSoftwareImageDetailsQueryParams struct {
	ImageUUID            string `url:"imageUuid,omitempty"`            //imageUuid
	Name                 string `url:"name,omitempty"`                 //name
	Family               string `url:"family,omitempty"`               //family
	ApplicationType      string `url:"applicationType,omitempty"`      //applicationType
	ImageIntegrityStatus string `url:"imageIntegrityStatus,omitempty"` //imageIntegrityStatus - FAILURE, UNKNOWN, VERIFIED
	Version              string `url:"version,omitempty"`              //software Image Version
	ImageSeries          string `url:"imageSeries,omitempty"`          //image Series
	ImageName            string `url:"imageName,omitempty"`            //image Name
	IsTaggedGolden       bool   `url:"isTaggedGolden,omitempty"`       //is Tagged Golden
	IsCCORecommended     bool   `url:"isCCORecommended,omitempty"`     //is recommended from cisco.com
	IsCCOLatest          bool   `url:"isCCOLatest,omitempty"`          //is latest from cisco.com
	CreatedTime          int    `url:"createdTime,omitempty"`          //time in milliseconds (epoch format)
	ImageSizeGreaterThan int    `url:"imageSizeGreaterThan,omitempty"` //size in bytes
	ImageSizeLesserThan  int    `url:"imageSizeLesserThan,omitempty"`  //size in bytes
	SortBy               string `url:"sortBy,omitempty"`               //sort results by this field
	SortOrder            string `url:"sortOrder,omitempty"`            //sort order - 'asc' or 'des'. Default is asc
	Limit                int    `url:"limit,omitempty"`                //limit
	Offset               int    `url:"offset,omitempty"`               //offset
}
type ImportLocalSoftwareImageQueryParams struct {
	IsThirdParty              bool   `url:"isThirdParty,omitempty"`              //Third party Image check
	ThirdPartyVendor          string `url:"thirdPartyVendor,omitempty"`          //Third Party Vendor
	ThirdPartyImageFamily     string `url:"thirdPartyImageFamily,omitempty"`     //Third Party image family
	ThirdPartyApplicationType string `url:"thirdPartyApplicationType,omitempty"` //Third Party Application Type
}

type ImportLocalSoftwareImageMultipartFields struct {
	File     io.Reader
	FileName string
}

type ImportSoftwareImageViaURLQueryParams struct {
	ScheduleAt     string `url:"scheduleAt,omitempty"`     //Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional)
	ScheduleDesc   string `url:"scheduleDesc,omitempty"`   //Custom Description (Optional)
	ScheduleOrigin string `url:"scheduleOrigin,omitempty"` //Originator of this call (Optional)
}

type ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivation struct {
	Response *ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationResponse `json:"response,omitempty"` //
	Version  string                                                                     `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistribution struct {
	Response *ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionResponse `json:"response,omitempty"` //
	Version  string                                                                       `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetails struct {
	Response *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse struct {
	ApplicableDevicesForImage *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseApplicableDevicesForImage `json:"applicableDevicesForImage,omitempty"` //
	ApplicationType           string                                                                                         `json:"applicationType,omitempty"`           //
	CreatedTime               string                                                                                         `json:"createdTime,omitempty"`               //
	ExtendedAttributes        *ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseExtendedAttributes          `json:"extendedAttributes,omitempty"`        //
	Family                    string                                                                                         `json:"family,omitempty"`                    //
	Feature                   string                                                                                         `json:"feature,omitempty"`                   //
	FileServiceID             string                                                                                         `json:"fileServiceId,omitempty"`             //
	FileSize                  string                                                                                         `json:"fileSize,omitempty"`                  //
	ImageIntegrityStatus      string                                                                                         `json:"imageIntegrityStatus,omitempty"`      //
	ImageName                 string                                                                                         `json:"imageName,omitempty"`                 //
	ImageSeries               []string                                                                                       `json:"imageSeries,omitempty"`               //
	ImageSource               string                                                                                         `json:"imageSource,omitempty"`               //
	ImageType                 string                                                                                         `json:"imageType,omitempty"`                 //
	ImageUUID                 string                                                                                         `json:"imageUuid,omitempty"`                 //
	ImportSourceType          string                                                                                         `json:"importSourceType,omitempty"`          //
	IsTaggedGolden            *bool                                                                                          `json:"isTaggedGolden,omitempty"`            //
	Md5Checksum               string                                                                                         `json:"md5Checksum,omitempty"`               //
	Name                      string                                                                                         `json:"name,omitempty"`                      //
	ProfileInfo               *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfo               `json:"profileInfo,omitempty"`               //
	ShaCheckSum               string                                                                                         `json:"shaCheckSum,omitempty"`               //
	Vendor                    string                                                                                         `json:"vendor,omitempty"`                    //
	Version                   string                                                                                         `json:"version,omitempty"`                   //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseApplicableDevicesForImage struct {
	MdfID       string   `json:"mdfId,omitempty"`       //
	ProductID   []string `json:"productId,omitempty"`   //
	ProductName string   `json:"productName,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseExtendedAttributes interface{}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfo struct {
	Description        string                                                                                           `json:"description,omitempty"`        //
	ExtendedAttributes *ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfoExtendedAttributes `json:"extendedAttributes,omitempty"` //
	Memory             *int                                                                                             `json:"memory,omitempty"`             //
	ProductType        string                                                                                           `json:"productType,omitempty"`        //
	ProfileName        string                                                                                           `json:"profileName,omitempty"`        //
	Shares             *int                                                                                             `json:"shares,omitempty"`             //
	VCPU               *int                                                                                             `json:"vCpu,omitempty"`               //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfoExtendedAttributes interface{}
type ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiers struct {
	Version  string                                                                   `json:"version,omitempty"`  // Response Version
	Response *[]ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersResponse struct {
	DeviceFamily           string `json:"deviceFamily,omitempty"`           // Device Family e.g. : Cisco Catalyst 6503 Switch-Catalyst 6500 Series Supervisor Engine 2T
	DeviceFamilyIDentifier string `json:"deviceFamilyIdentifier,omitempty"` // Device Family Identifier used for tagging an image golden for certain Device Family e.g. : 277696480-283933147
}
type ResponseSoftwareImageManagementSwimTagAsGoldenImage struct {
	Version  string                                                       `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSoftwareImageManagementSwimTagAsGoldenImageResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimTagAsGoldenImageResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage struct {
	Version  string                                                              `json:"version,omitempty"`  // Response Version e.g. : 1.0
	Response *ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageResponse struct {
	URL    string `json:"url,omitempty"`    // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImage struct {
	Version  string                                                                  `json:"version,omitempty"`  // Response Version. E.G. : 1.0
	Response *ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageResponse struct {
	DeviceRole        string `json:"deviceRole,omitempty"`        // Device Role. Possible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
	TaggedGolden      *bool  `json:"taggedGolden,omitempty"`      // Tagged Golden.
	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the name of the site from where the tag is inherited.
	InheritedSiteID   string `json:"inheritedSiteId,omitempty"`   // Inherited Site Id. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the uuid of the site from where the tag is inherited. In case the golden tag is inherited from the Global site the value will be "-1".
}
type ResponseSoftwareImageManagementSwimImportLocalSoftwareImage struct {
	Response *ResponseSoftwareImageManagementSwimImportLocalSoftwareImageResponse `json:"response,omitempty"` //
	Version  string                                                               `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimImportLocalSoftwareImageResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type ResponseSoftwareImageManagementSwimImportSoftwareImageViaURL struct {
	Response *ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLResponse `json:"response,omitempty"` //
	Version  string                                                                `json:"version,omitempty"`  //
}
type ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}
type RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation []RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation // Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation
type RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation struct {
	ActivateLowerImageVersion *bool    `json:"activateLowerImageVersion,omitempty"` //
	DeviceUpgradeMode         string   `json:"deviceUpgradeMode,omitempty"`         //
	DeviceUUID                string   `json:"deviceUuid,omitempty"`                //
	DistributeIfNeeded        *bool    `json:"distributeIfNeeded,omitempty"`        //
	ImageUUIDList             []string `json:"imageUuidList,omitempty"`             //
	SmuImageUUIDList          []string `json:"smuImageUuidList,omitempty"`          //
}
type RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution []RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution // Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution
type RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution struct {
	DeviceUUID string `json:"deviceUuid,omitempty"` //
	ImageUUID  string `json:"imageUuid,omitempty"`  //
}
type RequestSoftwareImageManagementSwimTagAsGoldenImage struct {
	ImageID                string `json:"imageId,omitempty"`                // imageId in uuid format.
	SiteID                 string `json:"siteId,omitempty"`                 // SiteId in uuid format. For Global Site "-1" to be used.
	DeviceRole             string `json:"deviceRole,omitempty"`             // Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.
	DeviceFamilyIDentifier string `json:"deviceFamilyIdentifier,omitempty"` // Device Family Identifier e.g. : 277696480-283933147, 277696480
}
type RequestSoftwareImageManagementSwimImportSoftwareImageViaURL []RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL // Array of RequestSoftwareImageManagementSwimImportSoftwareImageViaURL
type RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL struct {
	ApplicationType string `json:"applicationType,omitempty"` //
	ImageFamily     string `json:"imageFamily,omitempty"`     //
	SourceURL       string `json:"sourceURL,omitempty"`       //
	ThirdParty      *bool  `json:"thirdParty,omitempty"`      //
	Vendor          string `json:"vendor,omitempty"`          //
}

//GetSoftwareImageDetails Get software image details - 0c8f-7a0b-49b9-aedd
/* Returns software image list based on a filter criteria. For example: "filterbyName = cat3k%"


@param GetSoftwareImageDetailsQueryParams Filtering parameter
*/
func (s *SoftwareImageManagementSwimService) GetSoftwareImageDetails(GetSoftwareImageDetailsQueryParams *GetSoftwareImageDetailsQueryParams) (*ResponseSoftwareImageManagementSwimGetSoftwareImageDetails, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation"

	queryString, _ := query.Values(GetSoftwareImageDetailsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimGetSoftwareImageDetails{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetSoftwareImageDetails")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetSoftwareImageDetails)
	return result, response, err

}

//GetDeviceFamilyIDentifiers Get Device Family Identifiers - 35ae-1bec-4bd8-89fc
/* API to get Device Family Identifiers for all Device Families that can be used for tagging an image golden.


 */
func (s *SoftwareImageManagementSwimService) GetDeviceFamilyIDentifiers() (*ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiers, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/device-family-identifiers"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetDeviceFamilyIdentifiers")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiers)
	return result, response, err

}

//GetGoldenTagStatusOfAnImage Get Golden Tag Status of an Image. - 96bf-b9b4-419b-a6d0
/* Get golden tag status of an image. Set siteId as -1 for Global site.


@param siteID siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.

@param deviceFamilyIDentifier deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480

@param deviceRole deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

@param imageID imageId path parameter. Image Id in uuid format.

*/
func (s *SoftwareImageManagementSwimService) GetGoldenTagStatusOfAnImage(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/golden/site/{siteId}/family/{deviceFamilyIdentifier}/role/{deviceRole}/image/{imageId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{deviceFamilyIdentifier}", fmt.Sprintf("%v", deviceFamilyIDentifier), -1)
	path = strings.Replace(path, "{deviceRole}", fmt.Sprintf("%v", deviceRole), -1)
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation GetGoldenTagStatusOfAnImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImage)
	return result, response, err

}

//TriggerSoftwareImageActivation Trigger software image activation - fb9b-eb66-4f2a-ba4c
/* Activates a software image on a given device. Software image must be present in the device flash


@param TriggerSoftwareImageActivationHeaderParams Custom header parameters
@param TriggerSoftwareImageActivationQueryParams Filtering parameter
*/
func (s *SoftwareImageManagementSwimService) TriggerSoftwareImageActivation(requestSoftwareImageManagementSwimTriggerSoftwareImageActivation *RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation, TriggerSoftwareImageActivationHeaderParams *TriggerSoftwareImageActivationHeaderParams, TriggerSoftwareImageActivationQueryParams *TriggerSoftwareImageActivationQueryParams) (*ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivation, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/activation/device"

	queryString, _ := query.Values(TriggerSoftwareImageActivationQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if TriggerSoftwareImageActivationHeaderParams != nil {

		if TriggerSoftwareImageActivationHeaderParams.ClientType != "" {
			clientRequest = clientRequest.SetHeader("Client-Type", TriggerSoftwareImageActivationHeaderParams.ClientType)
		}

		if TriggerSoftwareImageActivationHeaderParams.ClientURL != "" {
			clientRequest = clientRequest.SetHeader("Client-Url", TriggerSoftwareImageActivationHeaderParams.ClientURL)
		}

	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).SetBody(requestSoftwareImageManagementSwimTriggerSoftwareImageActivation).
		SetResult(&ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivation{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation TriggerSoftwareImageActivation")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivation)
	return result, response, err

}

//TriggerSoftwareImageDistribution Trigger software image distribution - 8cb6-783b-4fab-a1f4
/* Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it can be distributed


 */
func (s *SoftwareImageManagementSwimService) TriggerSoftwareImageDistribution(requestSoftwareImageManagementSwimTriggerSoftwareImageDistribution *RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution) (*ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistribution, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/distribution"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimTriggerSoftwareImageDistribution).
		SetResult(&ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistribution{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation TriggerSoftwareImageDistribution")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistribution)
	return result, response, err

}

//TagAsGoldenImage Tag as Golden Image - 5c91-7a67-474b-a0e0
/* Golden Tag image. Set siteId as -1 for Global site.


 */
func (s *SoftwareImageManagementSwimService) TagAsGoldenImage(requestSoftwareImageManagementSwimTagAsGoldenImage *RequestSoftwareImageManagementSwimTagAsGoldenImage) (*ResponseSoftwareImageManagementSwimTagAsGoldenImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/golden"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimTagAsGoldenImage).
		SetResult(&ResponseSoftwareImageManagementSwimTagAsGoldenImage{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation TagAsGoldenImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTagAsGoldenImage)
	return result, response, err

}

//ImportLocalSoftwareImage Import local software image - 4dbe-3bc7-43a8-91bc
/* Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2


@param ImportLocalSoftwareImageQueryParams Filtering parameter
*/
func (s *SoftwareImageManagementSwimService) ImportLocalSoftwareImage(ImportLocalSoftwareImageQueryParams *ImportLocalSoftwareImageQueryParams, ImportLocalSoftwareImageMultipartFields *ImportLocalSoftwareImageMultipartFields) (*ResponseSoftwareImageManagementSwimImportLocalSoftwareImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/source/file"

	queryString, _ := query.Values(ImportLocalSoftwareImageQueryParams)

	var response *resty.Response
	var err error
	clientRequest := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	if ImportLocalSoftwareImageMultipartFields != nil {
		clientRequest = clientRequest.SetFileReader("file", ImportLocalSoftwareImageMultipartFields.FileName, ImportLocalSoftwareImageMultipartFields.File)
	}

	response, err = clientRequest.
		SetQueryString(queryString.Encode()).
		SetResult(&ResponseSoftwareImageManagementSwimImportLocalSoftwareImage{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportLocalSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimImportLocalSoftwareImage)
	return result, response, err

}

//ImportSoftwareImageViaURL Import software image via URL - bc8a-ab47-46ca-883d
/* Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2


@param ImportSoftwareImageViaURLQueryParams Filtering parameter
*/
func (s *SoftwareImageManagementSwimService) ImportSoftwareImageViaURL(requestSoftwareImageManagementSwimImportSoftwareImageViaURL *RequestSoftwareImageManagementSwimImportSoftwareImageViaURL, ImportSoftwareImageViaURLQueryParams *ImportSoftwareImageViaURLQueryParams) (*ResponseSoftwareImageManagementSwimImportSoftwareImageViaURL, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/source/url"

	queryString, _ := query.Values(ImportSoftwareImageViaURLQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetBody(requestSoftwareImageManagementSwimImportSoftwareImageViaURL).
		SetResult(&ResponseSoftwareImageManagementSwimImportSoftwareImageViaURL{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation ImportSoftwareImageViaUrl")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimImportSoftwareImageViaURL)
	return result, response, err

}

//RemoveGoldenTagForImage Remove Golden Tag for image. - f3b9-5978-4f6a-897a
/* Remove golden tag. Set siteId as -1 for Global site.


@param siteID siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.

@param deviceFamilyIDentifier deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480

@param deviceRole deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

@param imageID imageId path parameter. Image Id in uuid format.

*/
func (s *SoftwareImageManagementSwimService) RemoveGoldenTagForImage(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/image/importation/golden/site/{siteId}/family/{deviceFamilyIdentifier}/role/{deviceRole}/image/{imageId}"
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)
	path = strings.Replace(path, "{deviceFamilyIdentifier}", fmt.Sprintf("%v", deviceFamilyIDentifier), -1)
	path = strings.Replace(path, "{deviceRole}", fmt.Sprintf("%v", deviceRole), -1)
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		return nil, response, fmt.Errorf("error with operation RemoveGoldenTagForImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage)
	return result, response, err

}
